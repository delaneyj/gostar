// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feDropShadow is generated from configuration file.
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

// The <feDropShadow> filter primitive creates a drop shadow of the input image
// It is a shorthand filter, and is defined in terms of the <feGaussianBlur> and
// <feOffset> filter primitives.
type SVGFEDROPSHADOWElement struct {
	*Element
}

// Create a new SVGFEDROPSHADOWElement element.
// This will create a new element with the tag
// "feDropShadow" during rendering.
func SVG_FEDROPSHADOW(children ...ElementRenderer) *SVGFEDROPSHADOWElement {
	e := NewElement("feDropShadow", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEDROPSHADOWElement{Element: e}
}

func (e *SVGFEDROPSHADOWElement) Children(children ...ElementRenderer) *SVGFEDROPSHADOWElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) Attr(name string, value string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFEDROPSHADOWElement) Attrs(attrs ...string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) AttrsMap(attrs map[string]string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) Text(text string) *SVGFEDROPSHADOWElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEDROPSHADOWElement) TextF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfText(condition bool, text string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) IfTextF(condition bool, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) Escaped(text string) *SVGFEDROPSHADOWElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEDROPSHADOWElement) IfEscaped(condition bool, text string) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) EscapedF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) CustomData(key, value string) *SVGFEDROPSHADOWElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfCustomData(condition bool, key, value string) *SVGFEDROPSHADOWElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) CustomDataF(key, format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) CustomDataRemove(key string) *SVGFEDROPSHADOWElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The amount of offset in the x direction
// If the <length> is 0, the shadow is placed at the same position as the input.
func (e *SVGFEDROPSHADOWElement) DX(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dx", f)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDX(condition bool, f float64) *SVGFEDROPSHADOWElement {
	if condition {
		e.DX(f)
	}
	return e
}

// The amount of offset in the y direction
// If the <length> is 0, the shadow is placed at the same position as the input.
func (e *SVGFEDROPSHADOWElement) DY(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dy", f)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDY(condition bool, f float64) *SVGFEDROPSHADOWElement {
	if condition {
		e.DY(f)
	}
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
func (e *SVGFEDROPSHADOWElement) STD_DEVIATION(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("stdDeviation", f)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfSTD_DEVIATION(condition bool, f float64) *SVGFEDROPSHADOWElement {
	if condition {
		e.STD_DEVIATION(f)
	}
	return e
}

// The flood-color attribute indicates what color to use to flood the current
// filter primitive subregion defined through the <feFlood> element
// If attribute 'flood-color' is not specified, then the effect is as if a value
// of black were specified.
func (e *SVGFEDROPSHADOWElement) FLOOD_COLOR(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("flood-color", s)
	return e
}

func (e *SVGFEDROPSHADOWElement) FLOOD_COLORF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.FLOOD_COLOR(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfFLOOD_COLOR(condition bool, s string) *SVGFEDROPSHADOWElement {
	if condition {
		e.FLOOD_COLOR(s)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) IfFLOOD_COLORF(condition bool, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.FLOOD_COLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute FLOOD_COLOR from the element.
func (e *SVGFEDROPSHADOWElement) FLOOD_COLORRemove(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("flood-color")
	return e
}

func (e *SVGFEDROPSHADOWElement) FLOOD_COLORRemoveF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.FLOOD_COLORRemove(fmt.Sprintf(format, args...))
}

// The flood-opacity attribute indicates the opacity value to use across the
// current filter primitive subregion defined through the <feFlood> element.
func (e *SVGFEDROPSHADOWElement) FLOOD_OPACITY(f float64) *SVGFEDROPSHADOWElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("flood-opacity", f)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfFLOOD_OPACITY(condition bool, f float64) *SVGFEDROPSHADOWElement {
	if condition {
		e.FLOOD_OPACITY(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEDROPSHADOWElement) ID(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEDROPSHADOWElement) IDF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfID(condition bool, s string) *SVGFEDROPSHADOWElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) IfIDF(condition bool, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEDROPSHADOWElement) IDRemove(s string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGFEDROPSHADOWElement) IDRemoveF(format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEDROPSHADOWElement) CLASS(s ...string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) IfCLASS(condition bool, s ...string) *SVGFEDROPSHADOWElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEDROPSHADOWElement) CLASSRemove(s ...string) *SVGFEDROPSHADOWElement {
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
func (e *SVGFEDROPSHADOWElement) STYLEF(k string, format string, args ...any) *SVGFEDROPSHADOWElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEDROPSHADOWElement) IfSTYLE(condition bool, k string, v string) *SVGFEDROPSHADOWElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEDROPSHADOWElement) STYLE(k string, v string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEDROPSHADOWElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEDROPSHADOWElement) STYLEMap(m map[string]string) *SVGFEDROPSHADOWElement {
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
func (e *SVGFEDROPSHADOWElement) STYLEPairs(pairs ...string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEDROPSHADOWElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEDROPSHADOWElement) STYLERemove(keys ...string) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) DATASTAR_MERGE_STORE(v any) *SVGFEDROPSHADOWElement {
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

func (e *SVGFEDROPSHADOWElement) DATASTAR_REF(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDATASTAR_REF(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFEDROPSHADOWElement) DATASTAR_REFRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFEDROPSHADOWElement) DATASTAR_BIND(key string, expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFEDROPSHADOWElement) DATASTAR_BINDRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFEDROPSHADOWElement) DATASTAR_MODEL(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFEDROPSHADOWElement) DATASTAR_MODELRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFEDROPSHADOWElement) DATASTAR_TEXT(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFEDROPSHADOWElement) DATASTAR_TEXTRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeDropShadowOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeDropShadowOnModDebounce(
	d time.Duration,
) SVGFeDropShadowOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeDropShadowOnModThrottle(
	d time.Duration,
) SVGFeDropShadowOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFEDROPSHADOWElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeDropShadowOnMod) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeDropShadowOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeDropShadowOnMod) *SVGFEDROPSHADOWElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFEDROPSHADOWElement) DATASTAR_ONRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFEDROPSHADOWElement) DATASTAR_FOCUSSet(b bool) *SVGFEDROPSHADOWElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEDROPSHADOWElement) DATASTAR_FOCUS() *SVGFEDROPSHADOWElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFEDROPSHADOWElement) DATASTAR_HEADER(key string, expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFEDROPSHADOWElement) DATASTAR_HEADERRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFEDROPSHADOWElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFEDROPSHADOWElement) DATASTAR_FETCH_INDICATORRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFEDROPSHADOWElement) DATASTAR_SHOWSet(b bool) *SVGFEDROPSHADOWElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEDROPSHADOWElement) DATASTAR_SHOW() *SVGFEDROPSHADOWElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFEDROPSHADOWElement) DATASTAR_INTERSECTS(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFEDROPSHADOWElement) DATASTAR_INTERSECTSRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFEDROPSHADOWElement) DATASTAR_TELEPORTSet(b bool) *SVGFEDROPSHADOWElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEDROPSHADOWElement) DATASTAR_TELEPORT() *SVGFEDROPSHADOWElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFEDROPSHADOWElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFEDROPSHADOWElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEDROPSHADOWElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFEDROPSHADOWElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFEDROPSHADOWElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDROPSHADOWElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFEDROPSHADOWElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFEDROPSHADOWElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFEDROPSHADOWElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
