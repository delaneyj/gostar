package html

import (
	"fmt"
	"html"
	"slices"
	"strings"

	"github.com/delaneyj/toolbelt"
	"github.com/samber/lo"
)

type ElementBuilder interface{
	Build(sb *strings.Builder) error
}

type Element struct {
	Tag                       string
	IsSelfClosing             bool
	BoolAttributes            map[string]struct{}
	StringAttributes          map[string]string
	DelimitedStringAttributes map[string]*DelimitedString
	DelimitedKVAttributes     map[string]*DelimitedKVString
	CustomDataAttributes      map[string]string
	Descendants               []ElementBuilder
}

func(e *Element) String() string {
	sb := &strings.Builder{}
	e.Build(sb)
	return sb.String()
}

func (e *Element) Build(sb *strings.Builder) error {

	sb.WriteRune('<')
	sb.WriteString(e.Tag)

	dataValues := lo.Entries(e.CustomDataAttributes)
	slices.SortFunc(dataValues, func(a, b lo.Entry[string, string]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range dataValues {
		sb.WriteString(" data-")
		sb.WriteString(toolbelt.Kebab(entry.Key))
		sb.WriteString("=\"")
		sb.WriteString(entry.Value)
		sb.WriteRune('"')
	}

	stringValues := lo.Entries(e.StringAttributes)
	slices.SortFunc(stringValues, func(a, b lo.Entry[string, string]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range stringValues {
		sb.WriteRune(' ')
		sb.WriteString(entry.Key)
		sb.WriteRune('=')
		sb.WriteRune('"')
		sb.WriteString(entry.Value)
		sb.WriteRune('"')
	}

	values := lo.Entries(e.DelimitedStringAttributes)
	slices.SortFunc(values, func(a, b lo.Entry[string, *DelimitedString]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range values {
		sb.WriteRune(' ')
		sb.WriteString(entry.Key)
		sb.WriteRune('=')
		sb.WriteRune('"')
		if err := entry.Value.Build(sb); err != nil {
			return fmt.Errorf("failed to build element: %w", err)
		}
		sb.WriteRune('"')
	}

	kvValues := lo.Entries(e.DelimitedKVAttributes)
	slices.SortFunc(kvValues, func(a, b lo.Entry[string, *DelimitedKVString]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range kvValues {
		sb.WriteRune(' ')
		sb.WriteString(entry.Key)
		sb.WriteRune('=')
		sb.WriteRune('"')
		if err := entry.Value.Build(sb); err != nil {
			return fmt.Errorf("failed to build element: %w", err)
		}
		sb.WriteRune('"')
	}

	boolStrings := lo.Keys(e.BoolAttributes)
	slices.Sort(boolStrings)
	for _, name := range boolStrings {
		sb.WriteRune(' ')
		sb.WriteString(name)
	}

	if e.IsSelfClosing {
		sb.WriteString(" />")
		return nil
	}

	sb.WriteRune('>')
	for _, child := range e.Descendants {
		if err := child.Build(sb); err != nil {
			return fmt.Errorf("failed to build element: %w", err)
		}
	}
	sb.WriteString("</")
	sb.WriteString(e.Tag)
	sb.WriteRune('>')

	return nil
}

type DelimitedString struct {
	delimiter string
	values    map[string]struct{}
}

func (ds *DelimitedString) Add(values ...string) {
	for _, value := range values {
		ds.values[value] = struct{}{}
	}
}

func (ds *DelimitedString) Remove(values ...string) {
	for _, value := range values {
		delete(ds.values, value)
	}
}

func (ds *DelimitedString) Build(sb *strings.Builder) error {
	values := lo.Keys(ds.values)
	slices.Sort(values)
	_, err := sb.WriteString(strings.Join(values, ds.delimiter))
	return err
}

func NewDelimitedString(delimiter string) *DelimitedString {
	return &DelimitedString{
		delimiter: delimiter,
		values:    make(map[string]struct{}),
	}
}

func NewSpaceDelimitedString() *DelimitedString {
	return NewDelimitedString(" ")
}

type DelimitedKVString struct {
	associationDelimiter string
	delimiter            string
	values               map[string]string
}

func (ds *DelimitedKVString) Add(key, value string) {
	ds.values[key] = value
}

func (ds *DelimitedKVString) Remove(key string) {
	delete(ds.values, key)
}

func (ds *DelimitedKVString) Build(sb *strings.Builder) error {
	values := make([]string, 0, len(ds.values))
	for k, v := range ds.values {
		values = append(values, k+ds.associationDelimiter+v)
	}
	slices.Sort(values)
	sb.WriteString(strings.Join(values, ds.delimiter))
	return nil
}

func NewDelimitedKVString(associationDelimiter, delimiter string) *DelimitedKVString {
	return &DelimitedKVString{
		delimiter:            delimiter,
		associationDelimiter: associationDelimiter,
		values:               make(map[string]string),
	}
}

func NewEqualSemicolonDelimitedKVString() *DelimitedKVString {
	return NewDelimitedKVString("=", ";")
}

type TextContent struct {
	text string
}

func (tc *TextContent) Build(sb *strings.Builder) error {
	sb.WriteString(tc.text)
	return nil
}

func TEXT(text string) *TextContent {
	return &TextContent{
		text: html.EscapeString(text),
	}
}

type RawContent struct {
	raw string
}

func (rc *RawContent) Build(sb *strings.Builder) error {
	sb.WriteString(rc.raw)
	return nil
}

func RAW(raw string) *RawContent {
	return &RawContent{
		raw: raw,
	}
}

type Grouper struct {
	Children []ElementBuilder
}

func (g *Grouper) Build(sb *strings.Builder) error {
	for _, child := range g.Children {
		if err := child.Build(sb); err != nil {
			return fmt.Errorf("failed to build element: %w", err)
		}
	}
	return nil
}

func Group(children ...ElementBuilder) *Grouper {
	return &Grouper{
		Children: children,
	}
}


func If(condition bool, children ...ElementBuilder) ElementBuilder {
	if condition {
		return Group(children...)
	}
	return nil
}

func Tern(condition bool, trueChildren, falseChildren ElementBuilder) ElementBuilder {
	if condition {
		return trueChildren
	}
	return falseChildren
}

func Range[T any](values []T, cb func(T) ElementBuilder) ElementBuilder {
	children := make([]ElementBuilder, 0, len(values))
	for _, value := range values {
		children = append(children, cb(value))
	}
	return Group(children...)
}

func NewElement(tag string, children ...ElementBuilder) *Element {
	return &Element{
		Tag:      tag,
		Descendants: children,
	}
}

func (e *Element) Attr(name string, value string) *Element {
	return e.Attrs(name, value)
}

func (e *Element) Attrs(attrs ...string) *Element {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes[k] = v
	}
	return e
}

func (e *Element) AttrsMap(attrs map[string]string) *Element {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	for k, v := range attrs {
		e.StringAttributes[k] = v
	}
	return e
}