package generator

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"
	"github.com/delaneyj/toolbelt"
	sprig "github.com/go-task/slim-sprig/v3"
	"github.com/samber/lo"
)

//go:embed templates
var templatesFS embed.FS

var templs *template.Template

func GenerateAll(ctx context.Context, outPath string, namespaces *pb.Namespaces) (err error) {
	if len(namespaces.Namespaces) == 0 {
		return fmt.Errorf("no namespaces specified")
	}
	// pkgs = pkgs[:1]
	// pkgs[0].Elements = pkgs[0].Elements[:1]

	if err := os.RemoveAll(outPath); err != nil {
		return fmt.Errorf("failed to remove output path: %w", err)
	}
	if err := os.MkdirAll(outPath, 0o755); err != nil {
		return fmt.Errorf("failed to create output path: %w", err)
	}

	fm := sprig.FuncMap()
	fm["pascal"] = toolbelt.Pascal
	fm["snake"] = toolbelt.Snake
	fm["camel"] = toolbelt.Camel
	fm["comments"] = strToComments
	fm["attrIsString"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_String_:
			return true
		}
		return false
	}
	fm["attrIsDelimited"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_Delimited:
			return true
		}
		return false
	}
	fm["attrIsKV"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_Kv:
			return true
		}
		return false
	}
	fm["attrIsRune"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_Rune:
			return true
		}
		return false
	}
	fm["attrIsBool"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_Bool:
			return true
		}
		return false
	}
	fm["attrIsInt"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_Integer:
			return true
		}
		return false
	}
	fm["attrIsNumber"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_Number:
			return true
		}
		return false
	}
	fm["attrIsChoices"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_Choices:
			return true
		}
		return false
	}
	fm["attrIsCustom"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_Custom:
			return true
		}
		return false
	}
	fm["attrIsJson"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_Json:
			return true
		}
		return false
	}
	fm["attrIsDuration"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_DurationMs:
			return true
		case *pb.Attribute_Type_DurationSec:
			return true
		}

		return false
	}
	fm["attrIsDurationMs"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_DurationMs:
			return true
		}
		return false
	}
	fm["attrIsDurationSec"] = func(attr *pb.Attribute_Type) bool {
		switch attr.Type.(type) {
		case *pb.Attribute_Type_DurationSec:
			return true
		}
		return false
	}
	fm["TrimDatastarPrefix"] = func(s string) string {
		return strings.TrimPrefix(s, "datastar-")
	}

	templs, err = template.New("base").Funcs(fm).ParseFS(templatesFS, "templates/*.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}

	builderFile, err := os.Create(filepath.Join(outPath, "gostar_builder.go"))
	if err != nil {
		return fmt.Errorf("failed to create builder file: %w", err)
	}
	defer builderFile.Close()

	if err := templs.ExecuteTemplate(builderFile, "builder.tmpl", nil); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	for _, ns := range namespaces.Namespaces {
		ns.Attributes = append(ns.Attributes, namespaces.Attributes...)
		for _, element := range ns.Elements {
			if err := generateElement(ctx, outPath, namespaces.Attributes, ns, element); err != nil {
				return fmt.Errorf("failed to generate element %s: %w", element.Tag, err)
			}
		}
	}

	return nil
}

func generateElement(ctx context.Context, pkgPath string, globalAttributes []*pb.Attribute, ns *pb.Namespace, element *pb.Element) error {
	if element.Name == "" {
		element.Name = element.Tag
	}
	element.Attributes = lo.Map(
		lo.Flatten([][]*pb.Attribute{
			element.Attributes,
			ns.Attributes,
			globalAttributes,
		}),
		func(attr *pb.Attribute, i int) *pb.Attribute {
			if attr.Name == "" {
				attr.Name = attr.Key
			}
			if attr.Key == "" {
				attr.Key = attr.Name
			}
			return attr
		},
	)
	element.Attributes = lo.UniqBy(
		element.Attributes,
		func(attr *pb.Attribute) string {
			return attr.Key
		},
	)
	// slices.SortFunc(element.Attributes, func(i, j *pb.Attribute) int {
	// 	return strings.Compare(i.Name, j.Name)
	// })
	for _, attr := range element.Attributes {
		s := toolbelt.Snake(attr.Name)
		s = strings.ReplaceAll(s, "-", "_")
		s = strings.ReplaceAll(s, ":", "_")
		s = strings.ToUpper(s)
		attr.Name = s
	}

	filename := fmt.Sprintf("%s_%s.go", toolbelt.Snake(ns.Name), toolbelt.Snake(element.Tag))
	elementFilepath := filepath.Join(pkgPath, filename)

	f, err := os.Create(elementFilepath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", elementFilepath, err)
	}
	defer f.Close()

	if err := templs.ExecuteTemplate(f, "element.tmpl", map[string]any{
		"Namespace": ns,
		"Element":   element,
	}); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func strToComments(s string) string {
	maxLen := 80
	lines := []string{}

	for _, row := range strings.Split(s, "\n") {
		splitOnSentences := strings.Split(row, ". ")
		for _, sentence := range splitOnSentences {
			words := strings.Split(sentence, " ")
			line := ""
			for _, word := range words {
				if len(line)+len(word)+1 > maxLen {
					lines = append(lines, line)
					line = ""
				}
				line += word + " "
			}
			lines = append(lines, line)
		}
	}
	for i, line := range lines {
		lines[i] = "// " + line
	}

	return strings.Join(lines, "\n")
}
