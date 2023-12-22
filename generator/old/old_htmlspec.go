package generator

// type CustomAttribute struct {
// 	Name        string
// 	AllowedTags []string
// 	Valuer      fmt.Stringer
// }

// type FromHTMLSpecArgs struct {
// 	TempDir                  string
// 	OutputPath               string
// 	KvAttributes             []string
// 	SpaceDelimitedAttributes []string
// 	CustomizationTemplate    string
// }

// func NewGenerateElementArgs() *FromHTMLSpecArgs {
// 	return &FromHTMLSpecArgs{
// 		TempDir:      "tmp",
// 		OutputPath:   "elements",
// 		KvAttributes: []string{},
// 	}
// }

// func (gea *FromHTMLSpecArgs) WithTempDir(tmpDir string) *FromHTMLSpecArgs {
// 	gea.TempDir = tmpDir
// 	return gea
// }

// func (gea *FromHTMLSpecArgs) WithOutputPath(outputPath string) *FromHTMLSpecArgs {
// 	gea.OutputPath = outputPath
// 	return gea
// }

// func (gea *FromHTMLSpecArgs) WithKvAttributes(kvAttributes ...string) *FromHTMLSpecArgs {
// 	gea.KvAttributes = append(gea.KvAttributes, kvAttributes...)
// 	return gea
// }

// func FromHTMLSpec(ctx context.Context, args *FromHTMLSpecArgs) (err error) {
// 	args.KvAttributes = append(args.KvAttributes, "style")
// 	args.KvAttributes = lo.Uniq(args.KvAttributes)

// 	args.SpaceDelimitedAttributes = append(args.SpaceDelimitedAttributes, "class")
// 	args.SpaceDelimitedAttributes = lo.Uniq(args.SpaceDelimitedAttributes)

// 	cleanAndUpsertFolder := func(path string, shouldClear bool) (string, error) {
// 		path, err = filepath.Abs(path)
// 		if err != nil {
// 			return "", fmt.Errorf("could not get absolute path for %s: %v", path, err)
// 		}
// 		if shouldClear {
// 			if err := os.RemoveAll(path); err != nil {
// 				return "", fmt.Errorf("could not remove %s: %v", path, err)
// 			}
// 		}
// 		if err := os.MkdirAll(path, 0755); err != nil {
// 			return "", fmt.Errorf("could not create %s: %v", path, err)
// 		}
// 		return path, nil
// 	}

// 	args.TempDir, err = cleanAndUpsertFolder(args.TempDir, false)
// 	if err != nil {
// 		return fmt.Errorf("could not clean and upsert %s: %v", args.TempDir, err)
// 	}

// 	args.OutputPath, err = cleanAndUpsertFolder(args.OutputPath, true)
// 	if err != nil {
// 		return fmt.Errorf("could not clean and upsert %s: %v", args.OutputPath, err)
// 	}

// 	elements, err := scrapeHTMLSpec(ctx, args)
// 	if err != nil {
// 		return fmt.Errorf("could not scrape HTML spec: %v", err)
// 	}

// 	log.Printf("validating elements")
// 	categoryElements := map[string][]*Element{}
// 	for _, element := range elements {
// 		for _, category := range element.Categories {
// 			categoryElements[category] = append(categoryElements[category], element)
// 		}
// 	}
// 	validChildElements := map[string][]string{}
// 	for _, element := range elements {
// 		for _, child := range element.ChildElementsOrCategories {
// 			if _, ok := categoryElements[child]; ok {
// 				childElements := categoryElements[child]

// 				for _, childElement := range childElements {
// 					validChildElements[element.Tag] = append(validChildElements[element.Tag], childElement.Tag)
// 				}
// 			}
// 		}
// 	}

// 	tmpls, err := template.ParseFS(templatesFS, "templates/*.tmpl")
// 	if err != nil {
// 		return fmt.Errorf("could not parse templates: %v", err)
// 	}

// 	log.Printf("generating %d elements", len(elements))
// 	outputPathParts := strings.Split(args.OutputPath, string(os.PathSeparator))
// 	lastOutputPathPart := outputPathParts[len(outputPathParts)-1]

// 	packageName := toolbelt.ToCasedString(lastOutputPathPart)

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

// 	elementCtxs := []*ElementCtx{}
// 	for _, element := range elements {
// 		elName := toolbelt.ToCasedString(element.Tag)
// 		elCtx := ElementCtx{
// 			PackageName:   packageName.Snake,
// 			ElementName:   fmt.Sprintf("%sHTMLElement", elName.Pascal),
// 			NewElement:    elName.Upper,
// 			Tag:           element.Tag,
// 			IsSelfClosing: element.IsVoidElement,
// 		}

// 		for _, attribute := range element.Attributes {
// 			attrCtx := AttributeCtx{
// 				Name:            attribute.Name,
// 				Key:             attribute.Key,
// 				KeyConst:        fmt.Sprintf("attribute%sKey", attribute.Name),
// 				ValidValueTypes: lo.Uniq(attribute.ValidValueTypes),
// 				Mode:            attribute.Mode,
// 				Description:     strings.Split(attribute.Description, "\n"),
// 			}
// 			slices.Sort(attrCtx.ValidValueTypes)
// 			elCtx.Attributes = append(elCtx.Attributes, attrCtx)
// 		}
// 		elementCtxs = append(elementCtxs, &elCtx)
// 	}
// 	slices.SortFunc(elementCtxs, func(a, b *ElementCtx) int {
// 		return strings.Compare(a.Tag, b.Tag)
// 	})

// 	constFileName := "gostar_consts.go"
// 	constFile, err := os.Create(filepath.Join(args.OutputPath, constFileName))
// 	if err != nil {
// 		return fmt.Errorf("could not create file %s: %v", constFileName, err)
// 	}
// 	defer constFile.Close()

// 	attributeKeys := map[string]string{}
// 	for _, element := range elementCtxs {
// 		for _, attribute := range element.Attributes {
// 			k := fmt.Sprintf("attribute%sKey", attribute.Name)
// 			v := attribute.Key
// 			attributeKeys[k] = v
// 		}
// 	}
// 	attributeEntries := lo.Entries(attributeKeys)
// 	slices.SortFunc(attributeEntries, func(a, b lo.Entry[string, string]) int {
// 		return strings.Compare(a.Key, b.Key)
// 	})

// 	if err := tmpls.ExecuteTemplate(constFile, "html_consts.tmpl", struct {
// 		PackageName          string
// 		ByteSliceConversions []lo.Entry[string, string]
// 		AttributeKeys        []lo.Entry[string, string]
// 	}{
// 		PackageName: packageName.Snake,
// 		ByteSliceConversions: lo.Map(elementCtxs, func(el *ElementCtx, i int) lo.Entry[string, string] {
// 			return lo.Entry[string, string]{
// 				Key:   "elementTag" + el.NewElement,
// 				Value: el.Tag,
// 			}
// 		}),
// 		AttributeKeys: attributeEntries,
// 	}); err != nil {
// 		return fmt.Errorf("could not execute template: %v", err)
// 	}

// 	wrappedCustomizationTemplate := fmt.Sprintf(`{{define "CUSTOMIZE_ELEMENT"}}
// %s
// {{end}}`, args.CustomizationTemplate)
// 	tmpls, err = tmpls.Parse(wrappedCustomizationTemplate)
// 	if err != nil {
// 		return fmt.Errorf("could not parse templates: %v", err)
// 	}
// 	for _, elCtx := range elementCtxs {
// 		fName := fmt.Sprintf("%s.go", toolbelt.Lower(toolbelt.Snake(elCtx.NewElement)))
// 		f, err := os.Create(filepath.Join(args.OutputPath, fName))
// 		if err != nil {
// 			return fmt.Errorf("could not create file %s: %v", fName, err)
// 		}

// 		if err := tmpls.ExecuteTemplate(f, "html_element.tmpl", elCtx); err != nil {
// 			return fmt.Errorf("could not execute template: %v", err)
// 		}
// 	}

// 	return nil
// }

// // Scrapes the W3C HTML spec for information about HTML elements and attributes.
// func scrapeHTMLSpec(ctx context.Context, args *FromHTMLSpecArgs) (elements []*Element, err error) {
// 	isCachedStr := "Is HTML spec cached? "

// 	cachedJSON := filepath.Join(args.TempDir, "html_spec.json")

// 	if _, err := os.Stat(cachedJSON); err == nil {
// 		log.Println(isCachedStr + "yes")

// 		b, err := os.ReadFile(cachedJSON)
// 		if err != nil {
// 			return nil, fmt.Errorf("could not read cached JSON: %v", err)
// 		}

// 		elements = []*Element{}
// 		if err := json.Unmarshal(b, &elements); err != nil {
// 			// If the cached JSON is invalid, just regenerate it.
// 			if err := os.Remove(cachedJSON); err != nil {
// 				return nil, fmt.Errorf("could not remove cached JSON: %v", err)
// 			}
// 		}
// 	} else {

// 		log.Println(isCachedStr + "no")

// 		const specURL = "https://html.spec.whatwg.org"
// 		log.Printf("getting %s", specURL)
// 		htmlSpecDoc, err := getDoc(specURL)
// 		if err != nil {
// 			return nil, fmt.Errorf("could not get HTML spec: %v", err)
// 		}

// 		log.Printf("reading elements")
// 		elementsMap, err := readElements(htmlSpecDoc)
// 		if err != nil {
// 			return nil, fmt.Errorf("could not read elements: %v", err)
// 		}

// 		log.Printf("applying attributes")
// 		if err := applyAttributes(htmlSpecDoc, elementsMap, args); err != nil {
// 			return nil, fmt.Errorf("could not read attributes: %v", err)
// 		}

// 		log.Printf("applying event handlers")
// 		if err := applyEventHandlers(htmlSpecDoc, elementsMap); err != nil {
// 			return nil, fmt.Errorf("could not read event handlers: %v", err)
// 		}

// 		elements = lo.Values(elementsMap)
// 		slices.SortFunc(elements, func(a, b *Element) int {
// 			return strings.Compare(a.Tag, b.Tag)
// 		})

// 		log.Printf("caching JSON")
// 		b, err := json.MarshalIndent(elements, "", "  ")
// 		if err != nil {
// 			return nil, fmt.Errorf("could not marshal elements: %v", err)
// 		}
// 		if err := os.WriteFile(cachedJSON, b, 0644); err != nil {
// 			return nil, fmt.Errorf("could not write cached JSON: %v", err)
// 		}
// 	}

// 	return elements, nil
// }

// func readElements(htmlSpecDoc *goquery.Document) (map[string]*Element, error) {
// 	log.Printf("reading void elements")
// 	voidElements := []string{}
// 	htmlSpecDoc.Find("#void-elements").First().Parent().Parent().Find("dd").First().Find("a").Each(func(i int, veLink *goquery.Selection) {
// 		voidElements = append(voidElements, veLink.Text())
// 	})

// 	expectedColumns := 7
// 	elements := map[string]*Element{}
// 	var err error
// 	htmlSpecDoc.Find("#elements-3 ~ table").First().Find("tbody > tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
// 		trChildren := s.Children()
// 		columnCount := trChildren.Length()
// 		if columnCount != expectedColumns {
// 			var html string
// 			html, err = s.Html()
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			err = fmt.Errorf("unexpected column count for row %d: %d, %s", i, columnCount, html)
// 			return false
// 		}

// 		possibleTags := trChildren.Eq(0).Find("a")

// 		tags := []string{}
// 		possibleTags.Each(func(i int, s *goquery.Selection) {
// 			tagRaw := s.Text()
// 			parts := strings.Split(tagRaw, " ")
// 			last := parts[len(parts)-1]

// 			tags = append(tags, last)
// 		})

// 		description := trChildren.Eq(1).Text()
// 		interfaceName := trChildren.Eq(6).Children().Filter("a").First().Text()

// 		categories := []string{}
// 		childElementsOrCategories := []string{}
// 		trChildren.Eq(2).Children().Filter("a").Each(func(i int, s *goquery.Selection) {
// 			c := strcase.ToSnake(s.Text())
// 			categories = append(categories, c)
// 		})
// 		trChildren.Eq(4).Children().Filter("a").Each(func(i int, s *goquery.Selection) {
// 			childElementsOrCategories = append(childElementsOrCategories, s.Text())
// 		})

// 		for _, tag := range tags {
// 			el := &Element{
// 				Tag:                       tag,
// 				Description:               description,
// 				Interface:                 interfaceName,
// 				Categories:                categories,
// 				ChildElementsOrCategories: childElementsOrCategories,
// 				IsVoidElement:             lo.Contains(voidElements, tag),
// 			}
// 			elements[tag] = el
// 		}

// 		return true
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("could not read elements: %v", err)
// 	}

// 	return elements, nil
// }

// func applyAttributes(htmlSpecDoc *goquery.Document, elements map[string]*Element, args *FromHTMLSpecArgs) error {
// 	if len(elements) == 0 {
// 		return fmt.Errorf("no elements")
// 	}
// 	expectedColumns := 4

// 	const validAllHTMLElementsText = "HTML elements"

// 	categoriesElements := map[string]map[string]*Element{}
// 	for _, element := range elements {
// 		for _, category := range element.Categories {
// 			if _, ok := categoriesElements[category]; !ok {
// 				categoriesElements[category] = map[string]*Element{}
// 			}
// 			categoriesElements[category][element.Tag] = element
// 		}
// 	}
// 	attributes := map[string]*Attribute{}
// 	elementsAttributeNames := map[string]sets.Set[string]{}
// 	for _, element := range elements {
// 		elementsAttributeNames[element.Tag] = sets.New[string]()
// 	}

// 	var (
// 		err                error
// 		knownBadCategories = sets.NewString(
// 			"form_associated_custom_elements",
// 		)
// 	)
// 	htmlSpecDoc.Find("#attributes-1 > tbody > tr").EachWithBreak(func(i int, rows *goquery.Selection) bool {
// 		columns := rows.Children()
// 		if columns.Length() != expectedColumns {
// 			err = fmt.Errorf("expected %d columns, got %d", expectedColumns, columns.Length())
// 			return false
// 		}

// 		name := toolbelt.ToCasedString(columns.Eq(0).Find("code").Text())

// 		attr := &Attribute{
// 			Name:        toolbelt.Upper(name.Snake),
// 			Key:         name.Kebab,
// 			Description: strings.TrimSpace(columns.Eq(2).Text()),
// 			Mode:        AttributeModeString,
// 		}

// 		columns.Eq(3).Find("a, code").Each(func(i int, a *goquery.Selection) {
// 			attr.ValidValueTypes = append(attr.ValidValueTypes, strcase.ToSnake(a.Text()))
// 		})

// 		if lo.Contains(args.KvAttributes, name.Lower) {
// 			attr.Mode = AttributeModeKV
// 		} else if lo.Contains(args.SpaceDelimitedAttributes, name.Lower) {
// 			attr.Mode = AttributeModeSpaceDelimited
// 		} else if len(attr.ValidValueTypes) == 1 && attr.ValidValueTypes[0] == "boolean_attribute" {
// 			attr.Mode = AttributeModeBool
// 		} else if len(attr.ValidValueTypes) == 1 && attr.ValidValueTypes[0] == "valid_non_negative_integer" {
// 			attr.Mode = AttributeModeInt
// 		}

// 		attributes[attr.Name] = attr

// 		columns.Eq(1).Find("a").EachWithBreak(func(i int, a *goquery.Selection) bool {
// 			text := a.Text()
// 			if text == validAllHTMLElementsText {
// 				for _, set := range elementsAttributeNames {
// 					set.Insert(attr.Name)
// 				}
// 			} else {
// 				element, ok := elements[text]
// 				if !ok {
// 					// not element, but maybe an category
// 					c := strcase.ToSnake(text)
// 					categoryElements, ok := categoriesElements[c]
// 					if !ok {
// 						if knownBadCategories.Has(c) {
// 							return true
// 						}

// 						// if it's not a category, then it's an error
// 						err = fmt.Errorf("could not category element %s", text)
// 						return false
// 					}

// 					for _, element := range categoryElements {
// 						elementsAttributeNames[element.Tag].Insert(attr.Name)
// 					}
// 				}

// 				elementsAttributeNames[element.Tag].Insert(attr.Name)
// 			}

// 			return true
// 		})
// 		return err == nil
// 	})
// 	if err != nil {
// 		return fmt.Errorf("could not read attributes: %v", err)
// 	}

// 	for elementTag, attributeNames := range elementsAttributeNames {
// 		e, ok := elements[elementTag]
// 		if !ok {
// 			return fmt.Errorf("could not find element %s", elementTag)
// 		}

// 		attrNames := attributeNames.UnsortedList()
// 		slices.Sort(attrNames)

// 		for _, attributeName := range attrNames {
// 			a, ok := attributes[attributeName]
// 			if !ok {
// 				return fmt.Errorf("could not find attribute %s", attributeName)
// 			}
// 			e.Attributes = append(e.Attributes, a)
// 		}
// 	}

// 	return nil
// }

// func applyEventHandlers(htmlSpecDoc *goquery.Document, elements map[string]*Element) (err error) {
// 	const domSpecURL = "https://dom.spec.whatwg.org"
// 	domSpecDoc, err := getDoc(domSpecURL)
// 	if err != nil {
// 		return fmt.Errorf("could not get dom spec: %v", err)
// 	}

// 	const pointerSpecURL = "https://w3c.github.io/pointerevents" // hot garbage
// 	const workingSpecURLWorking = "https://www.w3.org/TR/pointerevents"
// 	pointerSpecDoc, err := getDoc(workingSpecURLWorking)
// 	if err != nil {
// 		return fmt.Errorf("could not get pointer spec: %v", err)
// 	}

// 	attributeInfoRegex := regexp.MustCompile(`attribute (?P<type>[\w]*)(?P<optional>\?*) (?P<name>\w*);( // (?P<comment>(.*)))?`)
// 	pointerAttributeInfoRegex := regexp.MustCompile(`attribute (?P<type>\w*)\s*(?P<name>\w*)`)

// 	events := map[string]*Event{}
// 	htmlSpecDoc.Find("#events-2 ~ table").First().Find("tbody > tr").EachWithBreak(func(i int, rows *goquery.Selection) bool {
// 		columns := rows.Children()

// 		event := &Event{
// 			Name:        columns.Eq(0).Find("code").Text(),
// 			Description: strings.TrimSpace(columns.Eq(3).Text()),
// 			Interface:   columns.Eq(1).Find("code").Text(),
// 		}

// 		link, exists := columns.Eq(1).Find("a").Attr("href")
// 		if !exists {
// 			return true
// 		}
// 		relativeLink := link[strings.Index(link, "#"):]

// 		var (
// 			idlBlock string
// 			idlRegex = attributeInfoRegex
// 		)
// 		switch {
// 		case strings.HasPrefix(link, domSpecURL):
// 			idlBlock = domSpecDoc.Find(relativeLink + " ~ pre.idl").First().Text()
// 		case strings.HasPrefix(link, "#"):
// 			idlBlock = htmlSpecDoc.Find(link).Parent().Parent().Text()
// 		case strings.HasPrefix(link, pointerSpecURL):
// 			idlBlock = pointerSpecDoc.Find(relativeLink).Text()
// 			idlRegex = pointerAttributeInfoRegex
// 		default:
// 			err = fmt.Errorf("expected link to be relative to spec, got %s", link)
// 			return false
// 		}
// 		if idlBlock == "" {
// 			err = fmt.Errorf("could not find IDL block for %s", link)
// 			return false
// 		}

// 		matches := idlRegex.FindAllStringSubmatch(idlBlock, -1)
// 		for _, match := range matches {
// 			evtAttribute := EventAttribute{
// 				Name: match[idlRegex.SubexpIndex("name")],
// 				Type: match[idlRegex.SubexpIndex("type")],
// 			}
// 			event.Attributes = append(event.Attributes, evtAttribute)

// 		}
// 		events[event.Name] = event
// 		return true
// 	})
// 	if err != nil {
// 		return fmt.Errorf("could not read events: %v", err)
// 	}

// 	htmlSpecDoc.Find("#ix-event-handlers > tbody > tr").EachWithBreak(func(i int, rows *goquery.Selection) bool {
// 		columns := rows.Children()
// 		if columns.Length() != 4 {
// 			err = fmt.Errorf("expected %d columns, got %d", 4, columns.Length())
// 			return false
// 		}

// 		evtHandler := &EventHandler{
// 			Name:        columns.Eq(0).Find("code").Text(),
// 			Description: strings.TrimSpace(columns.Eq(2).Text()),
// 		}

// 		target := columns.Eq(1).Find("a").Text()
// 		targetsBody := strings.Contains(target, "body")
// 		if targetsBody {
// 			elements["body"].EventHandlers = append(elements["body"].EventHandlers, evtHandler)
// 		} else {
// 			for _, element := range elements {
// 				element.EventHandlers = append(element.EventHandlers, evtHandler)
// 			}
// 		}

// 		// TODO: Can't get this in a clean way from the spec, so we'll just hardcode it
// 		// https://html.spec.whatwg.org/multipage/indices.html#events-2

// 		return true
// 	})
// 	if err != nil {
// 		return fmt.Errorf("could not read event handlers: %v", err)
// 	}

// 	for _, element := range elements {
// 		slices.SortFunc(element.EventHandlers, func(a, b *EventHandler) int {
// 			return strings.Compare(a.Name, b.Name)
// 		})
// 	}
// 	return nil
// }
