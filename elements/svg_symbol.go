// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg symbol is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <symbol> SVG element is used to define graphical template objects which can
// be instantiated by a <use> element
// The use of symbol elements for graphics that are used multiple times in the
// same document adds structure and semantics
// Documents that are rich in structure may be rendered graphically, as speech, or
// as Braille, and thus promote accessibility
// note that a symbol element itself is not rendered
// Only instances of a symbol element (i.e., a reference to a symbol by a <use>
// element) are rendered
// To render a 'stand-alone' graphic that has been defined using a symbol, a
// reference to the symbol is referenced using a <use> element.
type SVGSYMBOLElement struct {
	*Element
}

// Create a new SVGSYMBOLElement element.
// This will create a new element with the tag
// "symbol" during rendering.
func SVG_SYMBOL(children ...ElementRenderer) *SVGSYMBOLElement {
	e := NewElement("symbol", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSYMBOLElement{Element: e}
}

func (e *SVGSYMBOLElement) Children(children ...ElementRenderer) *SVGSYMBOLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSYMBOLElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSYMBOLElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSYMBOLElement) Attr(name string, value string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGSYMBOLElement) Attrs(attrs ...string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) AttrsMap(attrs map[string]string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSYMBOLElement) Text(text string) *SVGSYMBOLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSYMBOLElement) TextF(format string, args ...any) *SVGSYMBOLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfText(condition bool, text string) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSYMBOLElement) IfTextF(condition bool, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSYMBOLElement) Escaped(text string) *SVGSYMBOLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSYMBOLElement) IfEscaped(condition bool, text string) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSYMBOLElement) EscapedF(format string, args ...any) *SVGSYMBOLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfEscapedF(condition bool, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSYMBOLElement) CustomData(key, value string) *SVGSYMBOLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSYMBOLElement) IfCustomData(condition bool, key, value string) *SVGSYMBOLElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSYMBOLElement) CustomDataF(key, format string, args ...any) *SVGSYMBOLElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSYMBOLElement) CustomDataRemove(key string) *SVGSYMBOLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Indicates how the fetched image is fitted into the destination rectangle.
func (e *SVGSYMBOLElement) PRESERVE_ASPECT_RATIO(c SVGSymbolPreserveAspectRatioChoice) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("preserveAspectRatio", string(c))
	return e
}

type SVGSymbolPreserveAspectRatioChoice string

const (
	// Do not force uniform scaling.
	SVGSymbolPreserveAspectRatio_none SVGSymbolPreserveAspectRatioChoice = "none"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGSymbolPreserveAspectRatio_xMinYMin SVGSymbolPreserveAspectRatioChoice = "xMinYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSymbolPreserveAspectRatio_xMidYMin SVGSymbolPreserveAspectRatioChoice = "xMidYMin"
	// Align the image with the corresponding side of the viewPort.
	SVGSymbolPreserveAspectRatio_xMaxYMin SVGSymbolPreserveAspectRatioChoice = "xMaxYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSymbolPreserveAspectRatio_xMinYMid SVGSymbolPreserveAspectRatioChoice = "xMinYMid"
	// Scale the image to the smallest size such that it can completely fit inside the
	// corresponding dimension of the viewPort.
	SVGSymbolPreserveAspectRatio_xMidYMid SVGSymbolPreserveAspectRatioChoice = "xMidYMid"
	// Align the image with the corresponding side of the viewPort.
	SVGSymbolPreserveAspectRatio_xMaxYMid SVGSymbolPreserveAspectRatioChoice = "xMaxYMid"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSymbolPreserveAspectRatio_xMinYMax SVGSymbolPreserveAspectRatioChoice = "xMinYMax"
	// Align the image with the corresponding side of the viewPort.
	SVGSymbolPreserveAspectRatio_xMidYMax SVGSymbolPreserveAspectRatioChoice = "xMidYMax"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGSymbolPreserveAspectRatio_xMaxYMax SVGSymbolPreserveAspectRatioChoice = "xMaxYMax"
)

// Remove the attribute PRESERVE_ASPECT_RATIO from the element.
func (e *SVGSYMBOLElement) PRESERVE_ASPECT_RATIORemove(c SVGSymbolPreserveAspectRatioChoice) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("preserveAspectRatio")
	return e
}

// Specifies a unique id for an element
func (e *SVGSYMBOLElement) ID(s string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSYMBOLElement) IDF(format string, args ...any) *SVGSYMBOLElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfID(condition bool, s string) *SVGSYMBOLElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSYMBOLElement) IfIDF(condition bool, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSYMBOLElement) IDRemove(s string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGSYMBOLElement) IDRemoveF(format string, args ...any) *SVGSYMBOLElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSYMBOLElement) CLASS(s ...string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) IfCLASS(condition bool, s ...string) *SVGSYMBOLElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSYMBOLElement) CLASSRemove(s ...string) *SVGSYMBOLElement {
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
func (e *SVGSYMBOLElement) STYLEF(k string, format string, args ...any) *SVGSYMBOLElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSYMBOLElement) IfSTYLE(condition bool, k string, v string) *SVGSYMBOLElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSYMBOLElement) STYLE(k string, v string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSYMBOLElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSYMBOLElement) STYLEMap(m map[string]string) *SVGSYMBOLElement {
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
func (e *SVGSYMBOLElement) STYLEPairs(pairs ...string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSYMBOLElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSYMBOLElement) STYLERemove(keys ...string) *SVGSYMBOLElement {
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

func (e *SVGSYMBOLElement) DATASTAR_MERGE_STORE(v any) *SVGSYMBOLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("merge-store", string(b))
	return e
}

// Sets the reference of the element

func (e *SVGSYMBOLElement) DATASTAR_REF(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_REF(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGSYMBOLElement) DATASTAR_REFRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGSYMBOLElement) DATASTAR_BIND(key string, expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGSYMBOLElement) DATASTAR_BINDRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGSYMBOLElement) DATASTAR_MODEL(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGSYMBOLElement) DATASTAR_MODELRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGSYMBOLElement) DATASTAR_TEXT(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGSYMBOLElement) DATASTAR_TEXTRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGSymbolOnMod customDataKeyModifier

// Debounces the event handler
func SVGSymbolOnModDebounce(
	d time.Duration,
) SVGSymbolOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGSymbolOnModThrottle(
	d time.Duration,
) SVGSymbolOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGSYMBOLElement) DATASTAR_ON(key string, expression string, modifiers ...SVGSymbolOnMod) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGSymbolOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGSymbolOnMod) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGSYMBOLElement) DATASTAR_ONRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGSYMBOLElement) DATASTAR_FOCUSSet(b bool) *SVGSYMBOLElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSYMBOLElement) DATASTAR_FOCUS() *SVGSYMBOLElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGSYMBOLElement) DATASTAR_HEADER(key string, expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGSYMBOLElement) DATASTAR_HEADERRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGSYMBOLElement) DATASTAR_FETCH_URL(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGSYMBOLElement) DATASTAR_FETCH_URLRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGSYMBOLElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGSYMBOLElement) DATASTAR_FETCH_INDICATORRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGSYMBOLElement) DATASTAR_SHOWSet(b bool) *SVGSYMBOLElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSYMBOLElement) DATASTAR_SHOW() *SVGSYMBOLElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGSYMBOLElement) DATASTAR_INTERSECTS(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGSYMBOLElement) DATASTAR_INTERSECTSRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGSYMBOLElement) DATASTAR_TELEPORTSet(b bool) *SVGSYMBOLElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSYMBOLElement) DATASTAR_TELEPORT() *SVGSYMBOLElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGSYMBOLElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGSYMBOLElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSYMBOLElement) DATASTAR_SCROLL_INTO_VIEW() *SVGSYMBOLElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGSYMBOLElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSYMBOLElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGSYMBOLElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGSYMBOLElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGSYMBOLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
