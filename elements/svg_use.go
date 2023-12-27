// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg use is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <use> SVG element includes a reference to a <symbol> element and attempts
// to display the referenced content
// The reference is drawn exactly as it was defined
// It can be reused as often as needed and can be programmatically manipulated.
type SVGUSEElement struct {
	*Element
}

// Create a new SVGUSEElement element.
// This will create a new element with the tag
// "use" during rendering.
func SVG_USE(children ...ElementRenderer) *SVGUSEElement {
	e := NewElement("use", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGUSEElement{Element: e}
}

func (e *SVGUSEElement) Children(children ...ElementRenderer) *SVGUSEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGUSEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGUSEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGUSEElement) Text(text string) *SVGUSEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGUSEElement) TextF(format string, args ...any) *SVGUSEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfText(condition bool, text string) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGUSEElement) IfTextF(condition bool, format string, args ...any) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGUSEElement) Escaped(text string) *SVGUSEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGUSEElement) IfEscaped(condition bool, text string) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGUSEElement) EscapedF(format string, args ...any) *SVGUSEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfEscapedF(condition bool, format string, args ...any) *SVGUSEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGUSEElement) CustomData(key, value string) *SVGUSEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGUSEElement) IfCustomData(condition bool, key, value string) *SVGUSEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGUSEElement) CustomDataF(key, format string, args ...any) *SVGUSEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGUSEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGUSEElement) CustomDataRemove(key string) *SVGUSEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// A URI reference to the symbol to use.
func (e *SVGUSEElement) HREF(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGUSEElement) IfHREF(condition bool, s string) *SVGUSEElement {
	if condition {
		e.HREF(s)
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGUSEElement) HREFRemove(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("href")
	return e
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGUSEElement) X(f float64) *SVGUSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("x", f)
	return e
}

func (e *SVGUSEElement) IfX(condition bool, f float64) *SVGUSEElement {
	if condition {
		e.X(f)
	}
	return e
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGUSEElement) Y(f float64) *SVGUSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("y", f)
	return e
}

func (e *SVGUSEElement) IfY(condition bool, f float64) *SVGUSEElement {
	if condition {
		e.Y(f)
	}
	return e
}

// The width of the rectangular region.
func (e *SVGUSEElement) WIDTH(f float64) *SVGUSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("width", f)
	return e
}

func (e *SVGUSEElement) IfWIDTH(condition bool, f float64) *SVGUSEElement {
	if condition {
		e.WIDTH(f)
	}
	return e
}

// The height of the rectangular region.
func (e *SVGUSEElement) HEIGHT(f float64) *SVGUSEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("height", f)
	return e
}

func (e *SVGUSEElement) IfHEIGHT(condition bool, f float64) *SVGUSEElement {
	if condition {
		e.HEIGHT(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGUSEElement) ID(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGUSEElement) IfID(condition bool, s string) *SVGUSEElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGUSEElement) IDRemove(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGUSEElement) CLASS(s ...string) *SVGUSEElement {
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

func (e *SVGUSEElement) IfCLASS(condition bool, s ...string) *SVGUSEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGUSEElement) CLASSRemove(s ...string) *SVGUSEElement {
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
func (e *SVGUSEElement) STYLEF(k string, format string, args ...any) *SVGUSEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGUSEElement) IfSTYLE(condition bool, k string, v string) *SVGUSEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGUSEElement) STYLE(k string, v string) *SVGUSEElement {
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

func (e *SVGUSEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGUSEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGUSEElement) STYLEMap(m map[string]string) *SVGUSEElement {
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
func (e *SVGUSEElement) STYLEPairs(pairs ...string) *SVGUSEElement {
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

func (e *SVGUSEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGUSEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGUSEElement) STYLERemove(keys ...string) *SVGUSEElement {
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

func (e *SVGUSEElement) DATASTAR_MERGE_STORE(v any) *SVGUSEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("data-merge-store", string(b))
	return e
}

// Sets the reference of the element

func (e *SVGUSEElement) DATASTAR_REF(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-ref"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGUSEElement) IfDATASTAR_REF(condition bool, s string) *SVGUSEElement {
	if condition {
		e.DATASTAR_REF(s)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGUSEElement) DATASTAR_REFRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGUSEElement) DATASTAR_BIND(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-bind"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGUSEElement) IfDATASTAR_BIND(condition bool, s string) *SVGUSEElement {
	if condition {
		e.DATASTAR_BIND(s)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGUSEElement) DATASTAR_BINDRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGUSEElement) DATASTAR_MODEL(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-model"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGUSEElement) IfDATASTAR_MODEL(condition bool, s string) *SVGUSEElement {
	if condition {
		e.DATASTAR_MODEL(s)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGUSEElement) DATASTAR_MODELRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGUSEElement) DATASTAR_TEXT(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-text"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGUSEElement) IfDATASTAR_TEXT(condition bool, s string) *SVGUSEElement {
	if condition {
		e.DATASTAR_TEXT(s)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGUSEElement) DATASTAR_TEXTRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGUseDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGUseDataOnModDebounce(
	s string,
) SVGUseDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%sms", s)
	}
}

// Throttles the event handler
func SVGUseDataOnModThrottle(
	s string,
) SVGUseDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%sms", s)
	}
}

func (e *SVGUSEElement) DATASTAR_ON(s string, modifiers ...SVGUseDataOnMod) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	customMods := lo.Map(modifiers, func(m SVGUseDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key := customDataKey("data-on", customMods...)
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGUSEElement) IfDATASTAR_ON(condition bool, s string, modifiers ...SVGUseDataOnMod) *SVGUSEElement {
	if condition {
		e.DATASTAR_ON(s, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGUSEElement) DATASTAR_ONRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGUSEElement) DATASTAR_FOCUSSet(b bool) *SVGUSEElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGUSEElement) DATASTAR_FOCUS() *SVGUSEElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGUSEElement) DATASTAR_HEADER(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-header"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGUSEElement) IfDATASTAR_HEADER(condition bool, s string) *SVGUSEElement {
	if condition {
		e.DATASTAR_HEADER(s)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGUSEElement) DATASTAR_HEADERRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGUSEElement) DATASTAR_FETCH_URL(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-fetch-url"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGUSEElement) IfDATASTAR_FETCH_URL(condition bool, s string) *SVGUSEElement {
	if condition {
		e.DATASTAR_FETCH_URL(s)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGUSEElement) DATASTAR_FETCH_URLRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGUSEElement) DATASTAR_FETCH_INDICATOR(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "DatastarFetchIndicator"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGUSEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, s string) *SVGUSEElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(s)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGUSEElement) DATASTAR_FETCH_INDICATORRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGUSEElement) DATASTAR_SHOWSet(b bool) *SVGUSEElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGUSEElement) DATASTAR_SHOW() *SVGUSEElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGUSEElement) DATASTAR_INTERSECTSSet(b bool) *SVGUSEElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGUSEElement) DATASTAR_INTERSECTS() *SVGUSEElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *SVGUSEElement) DATASTAR_TELEPORTSet(b bool) *SVGUSEElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGUSEElement) DATASTAR_TELEPORT() *SVGUSEElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGUSEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGUSEElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGUSEElement) DATASTAR_SCROLL_INTO_VIEW() *SVGUSEElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGUSEElement) DATASTAR_VIEW_TRANSITION(s string) *SVGUSEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-view-transition"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGUSEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, s string) *SVGUSEElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(s)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGUSEElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGUSEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
