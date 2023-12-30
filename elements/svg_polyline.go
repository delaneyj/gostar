// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg polyline is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <polyline> SVG element is an SVG basic shape, used to create a series of
// straight lines connecting several points
// Typically a polyline is used to create open shapes as the last point doesn't
// have to be connected to the first point
// For closed shapes see the <polygon> element.
type SVGPOLYLINEElement struct {
	*Element
}

// Create a new SVGPOLYLINEElement element.
// This will create a new element with the tag
// "polyline" during rendering.
func SVG_POLYLINE(children ...ElementRenderer) *SVGPOLYLINEElement {
	e := NewElement("polyline", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGPOLYLINEElement{Element: e}
}

func (e *SVGPOLYLINEElement) Children(children ...ElementRenderer) *SVGPOLYLINEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGPOLYLINEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGPOLYLINEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGPOLYLINEElement) Attr(name string, value string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGPOLYLINEElement) Attrs(attrs ...string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) AttrsMap(attrs map[string]string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGPOLYLINEElement) Text(text string) *SVGPOLYLINEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGPOLYLINEElement) TextF(format string, args ...any) *SVGPOLYLINEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfText(condition bool, text string) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGPOLYLINEElement) IfTextF(condition bool, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGPOLYLINEElement) Escaped(text string) *SVGPOLYLINEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGPOLYLINEElement) IfEscaped(condition bool, text string) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGPOLYLINEElement) EscapedF(format string, args ...any) *SVGPOLYLINEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfEscapedF(condition bool, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGPOLYLINEElement) CustomData(key, value string) *SVGPOLYLINEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGPOLYLINEElement) IfCustomData(condition bool, key, value string) *SVGPOLYLINEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGPOLYLINEElement) CustomDataF(key, format string, args ...any) *SVGPOLYLINEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGPOLYLINEElement) CustomDataRemove(key string) *SVGPOLYLINEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// A list of points, each of which is a coordinate pair.
func (e *SVGPOLYLINEElement) POINTS(s string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("points", s)
	return e
}

func (e *SVGPOLYLINEElement) IfPOINTS(condition bool, s string) *SVGPOLYLINEElement {
	if condition {
		e.POINTS(s)
	}
	return e
}

// Remove the attribute POINTS from the element.
func (e *SVGPOLYLINEElement) POINTSRemove(s string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("points")
	return e
}

// Specifies a unique id for an element
func (e *SVGPOLYLINEElement) ID(s string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGPOLYLINEElement) IfID(condition bool, s string) *SVGPOLYLINEElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGPOLYLINEElement) IDRemove(s string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGPOLYLINEElement) CLASS(s ...string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) IfCLASS(condition bool, s ...string) *SVGPOLYLINEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGPOLYLINEElement) CLASSRemove(s ...string) *SVGPOLYLINEElement {
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
func (e *SVGPOLYLINEElement) STYLEF(k string, format string, args ...any) *SVGPOLYLINEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGPOLYLINEElement) IfSTYLE(condition bool, k string, v string) *SVGPOLYLINEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGPOLYLINEElement) STYLE(k string, v string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGPOLYLINEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGPOLYLINEElement) STYLEMap(m map[string]string) *SVGPOLYLINEElement {
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
func (e *SVGPOLYLINEElement) STYLEPairs(pairs ...string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGPOLYLINEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGPOLYLINEElement) STYLERemove(keys ...string) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) DATASTAR_MERGE_STORE(v any) *SVGPOLYLINEElement {
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

func (e *SVGPOLYLINEElement) DATASTAR_REF(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfDATASTAR_REF(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGPOLYLINEElement) DATASTAR_REFRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGPOLYLINEElement) DATASTAR_BIND(key string, expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGPOLYLINEElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGPOLYLINEElement) DATASTAR_BINDRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGPOLYLINEElement) DATASTAR_MODEL(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGPOLYLINEElement) DATASTAR_MODELRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGPOLYLINEElement) DATASTAR_TEXT(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGPOLYLINEElement) DATASTAR_TEXTRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGPolylineDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGPolylineDataOnModDebounce(
	d time.Duration,
) SVGPolylineDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGPolylineDataOnModThrottle(
	d time.Duration,
) SVGPolylineDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGPOLYLINEElement) DATASTAR_ON(key string, expression string, modifiers ...SVGPolylineDataOnMod) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGPolylineDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGPolylineDataOnMod) *SVGPOLYLINEElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGPOLYLINEElement) DATASTAR_ONRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGPOLYLINEElement) DATASTAR_FOCUSSet(b bool) *SVGPOLYLINEElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGPOLYLINEElement) DATASTAR_FOCUS() *SVGPOLYLINEElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGPOLYLINEElement) DATASTAR_HEADER(key string, expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGPOLYLINEElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGPOLYLINEElement) DATASTAR_HEADERRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGPOLYLINEElement) DATASTAR_FETCH_URL(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGPOLYLINEElement) DATASTAR_FETCH_URLRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGPOLYLINEElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGPOLYLINEElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGPOLYLINEElement) DATASTAR_FETCH_INDICATORRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGPOLYLINEElement) DATASTAR_SHOWSet(b bool) *SVGPOLYLINEElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGPOLYLINEElement) DATASTAR_SHOW() *SVGPOLYLINEElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGPOLYLINEElement) DATASTAR_INTERSECTSSet(b bool) *SVGPOLYLINEElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGPOLYLINEElement) DATASTAR_INTERSECTS() *SVGPOLYLINEElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *SVGPOLYLINEElement) DATASTAR_TELEPORTSet(b bool) *SVGPOLYLINEElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGPOLYLINEElement) DATASTAR_TELEPORT() *SVGPOLYLINEElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGPOLYLINEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGPOLYLINEElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGPOLYLINEElement) DATASTAR_SCROLL_INTO_VIEW() *SVGPOLYLINEElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGPOLYLINEElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGPOLYLINEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGPOLYLINEElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGPOLYLINEElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGPOLYLINEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
