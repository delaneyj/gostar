// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feGaussianBlur is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <feGaussianBlur> SVG filter primitive blurs the input image by the amount
// specified in stdDeviation, which defines the bell-curve.
type SVGFEGAUSSIANBLURElement struct {
	*Element
}

// Create a new SVGFEGAUSSIANBLURElement element.
// This will create a new element with the tag
// "feGaussianBlur" during rendering.
func SVG_FEGAUSSIANBLUR(children ...ElementRenderer) *SVGFEGAUSSIANBLURElement {
	e := NewElement("feGaussianBlur", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEGAUSSIANBLURElement{Element: e}
}

func (e *SVGFEGAUSSIANBLURElement) Children(children ...ElementRenderer) *SVGFEGAUSSIANBLURElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) Text(text string) *SVGFEGAUSSIANBLURElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEGAUSSIANBLURElement) TextF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfText(condition bool, text string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfTextF(condition bool, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) Escaped(text string) *SVGFEGAUSSIANBLURElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfEscaped(condition bool, text string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) EscapedF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) CustomData(key, value string) *SVGFEGAUSSIANBLURElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfCustomData(condition bool, key, value string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) CustomDataF(key, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) CustomDataRemove(key string) *SVGFEGAUSSIANBLURElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The input for this filter.
func (e *SVGFEGAUSSIANBLURElement) IN(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfIN(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.IN(s)
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFEGAUSSIANBLURElement) INRemove(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in")
	return e
}

// The standard deviation for the blur operation
// If two <numbers> are provided, the first number represents a standard deviation
// value along the x-axis of the coordinate system established by attribute
// 'primitiveUnits' on the <filter> element
// The second value represents a standard deviation in Y
// If one number is provided, then that value is used for both X and Y
// Negative values are not allowed
// A value of zero disables the effect of the given filter primitive (i.e., the
// result is a transparent black image).
func (e *SVGFEGAUSSIANBLURElement) STD_DEVIATION(f float64) *SVGFEGAUSSIANBLURElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("stdDeviation", f)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfSTD_DEVIATION(condition bool, f float64) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.STD_DEVIATION(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEGAUSSIANBLURElement) ID(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfID(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEGAUSSIANBLURElement) IDRemove(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEGAUSSIANBLURElement) CLASS(s ...string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) IfCLASS(condition bool, s ...string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEGAUSSIANBLURElement) CLASSRemove(s ...string) *SVGFEGAUSSIANBLURElement {
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
func (e *SVGFEGAUSSIANBLURElement) STYLEF(k string, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) IfSTYLE(condition bool, k string, v string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) STYLE(k string, v string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEGAUSSIANBLURElement) STYLEMap(m map[string]string) *SVGFEGAUSSIANBLURElement {
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
func (e *SVGFEGAUSSIANBLURElement) STYLEPairs(pairs ...string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEGAUSSIANBLURElement) STYLERemove(keys ...string) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_MERGE_STORE(v any) *SVGFEGAUSSIANBLURElement {
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

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_REF(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-ref"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfDATASTAR_REF(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.DATASTAR_REF(s)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFEGAUSSIANBLURElement) DATASTAR_REFRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_BIND(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-bind"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfDATASTAR_BIND(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.DATASTAR_BIND(s)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFEGAUSSIANBLURElement) DATASTAR_BINDRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_MODEL(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-model"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfDATASTAR_MODEL(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.DATASTAR_MODEL(s)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFEGAUSSIANBLURElement) DATASTAR_MODELRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_TEXT(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-text"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfDATASTAR_TEXT(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.DATASTAR_TEXT(s)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFEGAUSSIANBLURElement) DATASTAR_TEXTRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeGaussianBlurDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeGaussianBlurDataOnModDebounce(
	s string,
) SVGFeGaussianBlurDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%sms", s)
	}
}

// Throttles the event handler
func SVGFeGaussianBlurDataOnModThrottle(
	s string,
) SVGFeGaussianBlurDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%sms", s)
	}
}

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_ON(s string, modifiers ...SVGFeGaussianBlurDataOnMod) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	customMods := lo.Map(modifiers, func(m SVGFeGaussianBlurDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key := customDataKey("data-on", customMods...)
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfDATASTAR_ON(condition bool, s string, modifiers ...SVGFeGaussianBlurDataOnMod) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.DATASTAR_ON(s, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFEGAUSSIANBLURElement) DATASTAR_ONRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_FOCUSSet(b bool) *SVGFEGAUSSIANBLURElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_FOCUS() *SVGFEGAUSSIANBLURElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_HEADER(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-header"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfDATASTAR_HEADER(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.DATASTAR_HEADER(s)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFEGAUSSIANBLURElement) DATASTAR_HEADERRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_FETCH_URL(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-fetch-url"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfDATASTAR_FETCH_URL(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.DATASTAR_FETCH_URL(s)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGFEGAUSSIANBLURElement) DATASTAR_FETCH_URLRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_FETCH_INDICATOR(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "DatastarFetchIndicator"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfDATASTAR_FETCH_INDICATOR(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(s)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFEGAUSSIANBLURElement) DATASTAR_FETCH_INDICATORRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_SHOWSet(b bool) *SVGFEGAUSSIANBLURElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_SHOW() *SVGFEGAUSSIANBLURElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_INTERSECTSSet(b bool) *SVGFEGAUSSIANBLURElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_INTERSECTS() *SVGFEGAUSSIANBLURElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_TELEPORTSet(b bool) *SVGFEGAUSSIANBLURElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_TELEPORT() *SVGFEGAUSSIANBLURElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFEGAUSSIANBLURElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFEGAUSSIANBLURElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFEGAUSSIANBLURElement) DATASTAR_VIEW_TRANSITION(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-view-transition"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfDATASTAR_VIEW_TRANSITION(condition bool, s string) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(s)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFEGAUSSIANBLURElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
