// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feDistantLight is generated from configuration file.
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

// The <feDistantLight> SVG filter primitive defines a distant light source that
// can be used within a lighting filter primitive: <feDiffuseLighting> or
// <feSpecularLighting>
type SVGFEDISTANTLIGHTElement struct {
	*Element
}

// Create a new SVGFEDISTANTLIGHTElement element.
// This will create a new element with the tag
// "feDistantLight" during rendering.
func SVG_FEDISTANTLIGHT(children ...ElementRenderer) *SVGFEDISTANTLIGHTElement {
	e := NewElement("feDistantLight", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEDISTANTLIGHTElement{Element: e}
}

func (e *SVGFEDISTANTLIGHTElement) Children(children ...ElementRenderer) *SVGFEDISTANTLIGHTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) Attr(name string, value string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) Attrs(attrs ...string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) AttrsMap(attrs map[string]string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) Text(text string) *SVGFEDISTANTLIGHTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEDISTANTLIGHTElement) TextF(format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfText(condition bool, text string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfTextF(condition bool, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) Escaped(text string) *SVGFEDISTANTLIGHTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfEscaped(condition bool, text string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) EscapedF(format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) CustomData(key, value string) *SVGFEDISTANTLIGHTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfCustomData(condition bool, key, value string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) CustomDataF(key, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) CustomDataRemove(key string) *SVGFEDISTANTLIGHTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The azimuth attribute represent the direction vector of the light source in the
// XY plane (clockwise), in degrees from the x axis.
func (e *SVGFEDISTANTLIGHTElement) AZIMUTH(f float64) *SVGFEDISTANTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("azimuth", f)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfAZIMUTH(condition bool, f float64) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.AZIMUTH(f)
	}
	return e
}

// The elevation attribute represent the direction vector of the light source
// perpendicular to the XY plane, in degrees from the XY plane towards the z axis
// (clockwise).
func (e *SVGFEDISTANTLIGHTElement) ELEVATION(f float64) *SVGFEDISTANTLIGHTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("elevation", f)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfELEVATION(condition bool, f float64) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.ELEVATION(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEDISTANTLIGHTElement) ID(s string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IDF(format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfID(condition bool, s string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfIDF(condition bool, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEDISTANTLIGHTElement) IDRemove(s string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IDRemoveF(format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEDISTANTLIGHTElement) CLASS(s ...string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) IfCLASS(condition bool, s ...string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEDISTANTLIGHTElement) CLASSRemove(s ...string) *SVGFEDISTANTLIGHTElement {
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
func (e *SVGFEDISTANTLIGHTElement) STYLEF(k string, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEDISTANTLIGHTElement) IfSTYLE(condition bool, k string, v string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEDISTANTLIGHTElement) STYLE(k string, v string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEDISTANTLIGHTElement) STYLEMap(m map[string]string) *SVGFEDISTANTLIGHTElement {
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
func (e *SVGFEDISTANTLIGHTElement) STYLEPairs(pairs ...string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEDISTANTLIGHTElement) STYLERemove(keys ...string) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_MERGE_STORE(v any) *SVGFEDISTANTLIGHTElement {
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

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_REF(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_REF(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_REFRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_BIND(key string, expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_BINDRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_MODEL(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_MODELRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_TEXT(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_TEXTRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeDistantLightOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeDistantLightOnModDebounce(
	d time.Duration,
) SVGFeDistantLightOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeDistantLightOnModThrottle(
	d time.Duration,
) SVGFeDistantLightOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeDistantLightOnMod) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeDistantLightOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeDistantLightOnMod) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_ONRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_FOCUSSet(b bool) *SVGFEDISTANTLIGHTElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_FOCUS() *SVGFEDISTANTLIGHTElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_HEADER(key string, expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_HEADERRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_FETCH_URL(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_FETCH_URLRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_FETCH_INDICATORRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_SHOWSet(b bool) *SVGFEDISTANTLIGHTElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_SHOW() *SVGFEDISTANTLIGHTElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_INTERSECTS(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_INTERSECTSRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_TELEPORTSet(b bool) *SVGFEDISTANTLIGHTElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_TELEPORT() *SVGFEDISTANTLIGHTElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFEDISTANTLIGHTElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFEDISTANTLIGHTElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFEDISTANTLIGHTElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEDISTANTLIGHTElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFEDISTANTLIGHTElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFEDISTANTLIGHTElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFEDISTANTLIGHTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
