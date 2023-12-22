package generator

// type FromSVGSpecArgs struct {
// 	OutputPath string
// 	TempDir    string
// }

// type SVGElement struct {
// 	PackageName string
// 	Name        toolbelt.CasedString
// 	Tag         string
// 	Description []string
// 	Attributes  []*SVGAttribute
// }

// type SVGAttribute struct {
// 	Key         string
// 	Name        toolbelt.CasedString
// 	Description []string
// 	Mode        string
// }

// const mdnBaseURL = "https://developer.mozilla.org/"

// var svgTagRegex = regexp.MustCompile(`Value type: ([^;]*);`)

// func svgValueType(s string) (string, error) {
// 	matches := svgTagRegex.FindStringSubmatch(s)
// 	if len(matches) >= 2 {
// 		return matches[1], nil
// 	}

// 	return "", fmt.Errorf("failed to get svg value type from %s", s)
// }

// func FromSVGSpec(ctx context.Context, args *FromSVGSpecArgs) error {
// 	if err := os.RemoveAll(args.OutputPath); err != nil {
// 		return err
// 	}
// 	if err := os.MkdirAll(args.OutputPath, 0755); err != nil {
// 		return err
// 	}

// 	const mdnElementsURL = mdnBaseURL + "/en-US/docs/Web/SVG/Element"
// 	doc, err := getDoc(mdnElementsURL)
// 	if err != nil {
// 		return fmt.Errorf("failed to get mdn elements doc: %w", err)
// 	}

// 	elementsSection := doc.Find("#sidebar-quicklinks > nav > div > div.sidebar-body > ol > li:nth-child(5)")

// 	refElementsTitle := elementsSection.Find(`details > summary`).Text()
// 	if refElementsTitle != "Elements" {
// 		return fmt.Errorf("unexpected elements title: %s", refElementsTitle)
// 	}

// 	elementsURLs := map[string]string{}

// 	elementsSection.Find("a").Each(func(i int, link *goquery.Selection) {
// 		href, ok := link.Attr("href")
// 		if !ok {
// 			log.Printf("no href found for link: %+v", link)
// 			return
// 		}
// 		text := link.Text()
// 		elementsURLs[text] = mdnBaseURL + href
// 	})
// 	elementEntries := lo.Entries(elementsURLs)
// 	slices.SortFunc(elementEntries, func(a, b lo.Entry[string, string]) int {
// 		return strings.Compare(a.Key, b.Key)
// 	})

// 	outputPathParts := strings.Split(args.OutputPath, string(os.PathSeparator))
// 	lastOutputPathPart := outputPathParts[len(outputPathParts)-1]
// 	packageName := toolbelt.ToCasedString(lastOutputPathPart)

// 	tmpls, err := template.ParseFS(templatesFS, "templates/*.tmpl")
// 	if err != nil {
// 		return fmt.Errorf("could not parse templates: %v", err)
// 	}

// 	goStarFileName := "gostar_elements.go"
// 	goStarFile, err := os.Create(filepath.Join(args.OutputPath, goStarFileName))
// 	if err != nil {
// 		return fmt.Errorf("could not create file %s: %v", goStarFileName, err)
// 	}
// 	defer goStarFile.Close()
// 	if err := tmpls.ExecuteTemplate(goStarFile, "html_gostar.tmpl", struct {
// 		PackageName string
// 	}{
// 		PackageName: packageName.Snake,
// 	}); err != nil {
// 		return fmt.Errorf("could not execute template: %v", err)
// 	}

// 	for _, entry := range elementEntries {
// 		elCtx, err := svgElementFromLink(entry.Key, entry.Value)
// 		if err != nil {
// 			return fmt.Errorf("failed to get svg element from link %s: %w", entry.Value, err)
// 		}
// 		elCtx.PackageName = packageName.Snake
// 		if err != nil {
// 			return fmt.Errorf("failed to get svg element %s from link %s: %w", entry.Key, entry.Value, err)
// 		}

// 		filename := fmt.Sprintf("%s.go", elCtx.Name.Snake)
// 		fullpath := filepath.Join(args.OutputPath, filename)
// 		f, err := os.Create(fullpath)
// 		if err != nil {
// 			return fmt.Errorf("failed to create file %s: %w", fullpath, err)
// 		}
// 		defer f.Close()

// 		if err := tmpls.ExecuteTemplate(f, "svg_elements.tmpl", elCtx); err != nil {
// 			return fmt.Errorf("failed to execute svg element template: %w", err)
// 		}
// 		break
// 	}
// 	return nil
// }

// func svgElementFromLink(tag, url string) (*SVGElement, error) {
// 	log.Printf("tag: %s, url: %s", tag, url)
// 	tag = tag[1 : len(tag)-1]

// 	doc, err := getDoc(url)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get doc: %w", err)
// 	}

// 	article := doc.Find("#content > article")
// 	description := article.Find("div.section-content").First().Text()

// 	el := &SVGElement{
// 		Tag:         tag,
// 		Name:        toolbelt.ToCasedString(tag),
// 		Description: strings.Split(description, "\n"),
// 		Attributes:  []*SVGAttribute{},
// 	}

// 	article.Find("#attributes").Parent().Find("dl").Children().Each(func(i int, c *goquery.Selection) {
// 		isEven := i%2 == 0
// 		text := strings.TrimSpace(c.Text())

// 		if isEven {
// 			a := &SVGAttribute{
// 				Key:  text,
// 				Name: toolbelt.ToCasedString(text),
// 			}
// 			el.Attributes = append(el.Attributes, a)
// 		} else {
// 			attrIdx := i / 2
// 			if attrIdx >= len(el.Attributes) {
// 				log.Printf("attrIdx %d is out of range for %d attributes", attrIdx, len(el.Attributes))
// 				return
// 			}
// 			a := el.Attributes[attrIdx]
// 			a.Description = strings.Split(text, "\n")

// 			a.Mode, err = svgValueType(text)
// 			if err != nil {
// 				log.Printf("failed to get svg value type for %s: %s", a.Key, err)
// 			}
// 		}
// 	})
// 	slices.SortFunc(el.Attributes, func(a, b *SVGAttribute) int {
// 		return strings.Compare(a.Key, b.Key)
// 	})

// 	el.Attributes = lo.Filter(el.Attributes, func(a *SVGAttribute, i int) bool {
// 		isExperimental := strings.Contains(a.Name.Upper, "EXPERIMENTAL")
// 		isDeprecated := strings.Contains(a.Name.Upper, "DEPRECATED")

// 		return !isExperimental && !isDeprecated
// 	})

// 	return el, nil
// }
