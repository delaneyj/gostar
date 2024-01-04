// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feImage is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"html"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <feImage> SVG filter primitive fetches image data from an external source
// and provides the pixel data as output (meaning if the external source is an SVG
// image, it is rasterized.)
type SVGFEIMAGEElement struct {
	*Element
}

// Create a new SVGFEIMAGEElement element.
// This will create a new element with the tag
// "feImage" during rendering.
func SVG_FEIMAGE(children ...ElementRenderer) *SVGFEIMAGEElement {
	e := NewElement("feImage", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEIMAGEElement{Element: e}
}

func (e *SVGFEIMAGEElement) Children(children ...ElementRenderer) *SVGFEIMAGEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEIMAGEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEIMAGEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEIMAGEElement) Attr(name string, value string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFEIMAGEElement) Attrs(attrs ...string) *SVGFEIMAGEElement {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEIMAGEElement) AttrsMap(attrs map[string]string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEIMAGEElement) Text(text string) *SVGFEIMAGEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEIMAGEElement) TextF(format string, args ...any) *SVGFEIMAGEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfText(condition bool, text string) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEIMAGEElement) IfTextF(condition bool, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEIMAGEElement) Escaped(text string) *SVGFEIMAGEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEIMAGEElement) IfEscaped(condition bool, text string) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEIMAGEElement) EscapedF(format string, args ...any) *SVGFEIMAGEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEIMAGEElement) CustomData(key, value string) *SVGFEIMAGEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEIMAGEElement) IfCustomData(condition bool, key, value string) *SVGFEIMAGEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEIMAGEElement) CustomDataF(key, format string, args ...any) *SVGFEIMAGEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEIMAGEElement) CustomDataRemove(key string) *SVGFEIMAGEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Indicates whether or not to force synchronous behavior.
func (e *SVGFEIMAGEElement) EXTERNAL_RESOURCES_REQUIRED() *SVGFEIMAGEElement {
	e.EXTERNAL_RESOURCES_REQUIREDSet(true)
	return e
}

func (e *SVGFEIMAGEElement) IfEXTERNAL_RESOURCES_REQUIRED(condition bool) *SVGFEIMAGEElement {
	if condition {
		e.EXTERNAL_RESOURCES_REQUIREDSet(true)
	}
	return e
}

// Set the attribute EXTERNAL_RESOURCES_REQUIRED to the value b explicitly.
func (e *SVGFEIMAGEElement) EXTERNAL_RESOURCES_REQUIREDSet(b bool) *SVGFEIMAGEElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("externalResourcesRequired", b)
	return e
}

func (e *SVGFEIMAGEElement) IfSetEXTERNAL_RESOURCES_REQUIRED(condition bool, b bool) *SVGFEIMAGEElement {
	if condition {
		e.EXTERNAL_RESOURCES_REQUIREDSet(b)
	}
	return e
}

// Remove the attribute EXTERNAL_RESOURCES_REQUIRED from the element.
func (e *SVGFEIMAGEElement) EXTERNAL_RESOURCES_REQUIREDRemove(b bool) *SVGFEIMAGEElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("externalResourcesRequired")
	return e
}

// Indicates how the fetched image is fitted into the destination rectangle.
func (e *SVGFEIMAGEElement) PRESERVE_ASPECT_RATIO(c SVGFeImagePreserveAspectRatioChoice) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("preserveAspectRatio", string(c))
	return e
}

type SVGFeImagePreserveAspectRatioChoice string

const (
	// Do not force uniform scaling.
	SVGFeImagePreserveAspectRatio_none SVGFeImagePreserveAspectRatioChoice = "none"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGFeImagePreserveAspectRatio_xMinYMin SVGFeImagePreserveAspectRatioChoice = "xMinYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGFeImagePreserveAspectRatio_xMidYMin SVGFeImagePreserveAspectRatioChoice = "xMidYMin"
	// Align the image with the corresponding side of the viewPort.
	SVGFeImagePreserveAspectRatio_xMaxYMin SVGFeImagePreserveAspectRatioChoice = "xMaxYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGFeImagePreserveAspectRatio_xMinYMid SVGFeImagePreserveAspectRatioChoice = "xMinYMid"
	// Scale the image to the smallest size such that it can completely fit inside the
	// corresponding dimension of the viewPort.
	SVGFeImagePreserveAspectRatio_xMidYMid SVGFeImagePreserveAspectRatioChoice = "xMidYMid"
	// Align the image with the corresponding side of the viewPort.
	SVGFeImagePreserveAspectRatio_xMaxYMid SVGFeImagePreserveAspectRatioChoice = "xMaxYMid"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGFeImagePreserveAspectRatio_xMinYMax SVGFeImagePreserveAspectRatioChoice = "xMinYMax"
	// Align the image with the corresponding side of the viewPort.
	SVGFeImagePreserveAspectRatio_xMidYMax SVGFeImagePreserveAspectRatioChoice = "xMidYMax"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGFeImagePreserveAspectRatio_xMaxYMax SVGFeImagePreserveAspectRatioChoice = "xMaxYMax"
)

// Remove the attribute PRESERVE_ASPECT_RATIO from the element.
func (e *SVGFEIMAGEElement) PRESERVE_ASPECT_RATIORemove(c SVGFeImagePreserveAspectRatioChoice) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("preserveAspectRatio")
	return e
}

// A URI reference to an external resource.
func (e *SVGFEIMAGEElement) HREF(s string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGFEIMAGEElement) HREFF(format string, args ...any) *SVGFEIMAGEElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfHREF(condition bool, s string) *SVGFEIMAGEElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGFEIMAGEElement) IfHREFF(condition bool, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGFEIMAGEElement) HREFRemove(s string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("href")
	return e
}

func (e *SVGFEIMAGEElement) HREFRemoveF(format string, args ...any) *SVGFEIMAGEElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFEIMAGEElement) ID(s string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEIMAGEElement) IDF(format string, args ...any) *SVGFEIMAGEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfID(condition bool, s string) *SVGFEIMAGEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEIMAGEElement) IfIDF(condition bool, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEIMAGEElement) IDRemove(s string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGFEIMAGEElement) IDRemoveF(format string, args ...any) *SVGFEIMAGEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEIMAGEElement) CLASS(s ...string) *SVGFEIMAGEElement {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = treemap.New[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		ds = NewDelimitedBuilder[string](" ")
		e.DelimitedStrings.Set("class", ds)
	}
	ds.Add(s...)
	return e
}

func (e *SVGFEIMAGEElement) IfCLASS(condition bool, s ...string) *SVGFEIMAGEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEIMAGEElement) CLASSRemove(s ...string) *SVGFEIMAGEElement {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFEIMAGEElement) STYLEF(k string, format string, args ...any) *SVGFEIMAGEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEIMAGEElement) IfSTYLE(condition bool, k string, v string) *SVGFEIMAGEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEIMAGEElement) STYLE(k string, v string) *SVGFEIMAGEElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	kv.Add(k, v)
	return e
}

func (e *SVGFEIMAGEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEIMAGEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEIMAGEElement) STYLEMap(m map[string]string) *SVGFEIMAGEElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	for k, v := range m {
		kv.Add(k, v)
	}
	return e
}

// Add pairs of attributes to the element.
func (e *SVGFEIMAGEElement) STYLEPairs(pairs ...string) *SVGFEIMAGEElement {
	if len(pairs)%2 != 0 {
		panic("Must have an even number of pairs")
	}
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}

	for i := 0; i < len(pairs); i += 2 {
		kv.Add(pairs[i], pairs[i+1])
	}

	return e
}

func (e *SVGFEIMAGEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEIMAGEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEIMAGEElement) STYLERemove(keys ...string) *SVGFEIMAGEElement {
	if e.KVStrings == nil {
		return e
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		return e
	}
	for _, k := range keys {
		kv.Remove(k)
	}
	return e
}

// Merges the store with the given object

func (e *SVGFEIMAGEElement) DATASTAR_MERGE_STORE(v any) *SVGFEIMAGEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("merge-store", html.EscapeString(string(b)))
	return e
}

// Sets the reference of the element

func (e *SVGFEIMAGEElement) DATASTAR_REF(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_REF(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFEIMAGEElement) DATASTAR_REFRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFEIMAGEElement) DATASTAR_BIND(key string, expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFEIMAGEElement) DATASTAR_BINDRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFEIMAGEElement) DATASTAR_MODEL(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFEIMAGEElement) DATASTAR_MODELRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFEIMAGEElement) DATASTAR_TEXT(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFEIMAGEElement) DATASTAR_TEXTRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeImageOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeImageOnModDebounce(
	d time.Duration,
) SVGFeImageOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeImageOnModThrottle(
	d time.Duration,
) SVGFeImageOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFEIMAGEElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeImageOnMod) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeImageOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeImageOnMod) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFEIMAGEElement) DATASTAR_ONRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFEIMAGEElement) DATASTAR_FOCUSSet(b bool) *SVGFEIMAGEElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEIMAGEElement) DATASTAR_FOCUS() *SVGFEIMAGEElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFEIMAGEElement) DATASTAR_HEADER(key string, expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFEIMAGEElement) DATASTAR_HEADERRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGFEIMAGEElement) DATASTAR_FETCH_URL(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGFEIMAGEElement) DATASTAR_FETCH_URLRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFEIMAGEElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFEIMAGEElement) DATASTAR_FETCH_INDICATORRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFEIMAGEElement) DATASTAR_SHOWSet(b bool) *SVGFEIMAGEElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEIMAGEElement) DATASTAR_SHOW() *SVGFEIMAGEElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFEIMAGEElement) DATASTAR_INTERSECTS(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFEIMAGEElement) DATASTAR_INTERSECTSRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFEIMAGEElement) DATASTAR_TELEPORTSet(b bool) *SVGFEIMAGEElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEIMAGEElement) DATASTAR_TELEPORT() *SVGFEIMAGEElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFEIMAGEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFEIMAGEElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEIMAGEElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFEIMAGEElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFEIMAGEElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEIMAGEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFEIMAGEElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFEIMAGEElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFEIMAGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
