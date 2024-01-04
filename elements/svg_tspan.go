// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg tspan is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <tspan> SVG element lets authors explicitly specify the location of a glyph
// along the given path via the attributes.
type SVGTSPANElement struct {
	*Element
}

// Create a new SVGTSPANElement element.
// This will create a new element with the tag
// "tspan" during rendering.
func SVG_TSPAN(children ...ElementRenderer) *SVGTSPANElement {
	e := NewElement("tspan", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGTSPANElement{Element: e}
}

func (e *SVGTSPANElement) Children(children ...ElementRenderer) *SVGTSPANElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGTSPANElement) IfChildren(condition bool, children ...ElementRenderer) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGTSPANElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGTSPANElement) Attr(name string, value string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGTSPANElement) Attrs(attrs ...string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) AttrsMap(attrs map[string]string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGTSPANElement) Text(text string) *SVGTSPANElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGTSPANElement) TextF(format string, args ...any) *SVGTSPANElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfText(condition bool, text string) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGTSPANElement) IfTextF(condition bool, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGTSPANElement) Escaped(text string) *SVGTSPANElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGTSPANElement) IfEscaped(condition bool, text string) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGTSPANElement) EscapedF(format string, args ...any) *SVGTSPANElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfEscapedF(condition bool, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGTSPANElement) CustomData(key, value string) *SVGTSPANElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGTSPANElement) IfCustomData(condition bool, key, value string) *SVGTSPANElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGTSPANElement) CustomDataF(key, format string, args ...any) *SVGTSPANElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGTSPANElement) CustomDataRemove(key string) *SVGTSPANElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The x-axis coordinate of the current text position.
func (e *SVGTSPANElement) X(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("x", f)
	return e
}

func (e *SVGTSPANElement) IfX(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.X(f)
	}
	return e
}

// The y-axis coordinate of the current text position.
func (e *SVGTSPANElement) Y(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("y", f)
	return e
}

func (e *SVGTSPANElement) IfY(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.Y(f)
	}
	return e
}

// The x-axis coordinate of the current text position.
func (e *SVGTSPANElement) DX(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dx", f)
	return e
}

func (e *SVGTSPANElement) IfDX(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.DX(f)
	}
	return e
}

// The y-axis coordinate of the current text position.
func (e *SVGTSPANElement) DY(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("dy", f)
	return e
}

func (e *SVGTSPANElement) IfDY(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.DY(f)
	}
	return e
}

// The rotation angle about the current text position.
func (e *SVGTSPANElement) ROTATE(f float64) *SVGTSPANElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("rotate", f)
	return e
}

func (e *SVGTSPANElement) IfROTATE(condition bool, f float64) *SVGTSPANElement {
	if condition {
		e.ROTATE(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGTSPANElement) ID(s string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGTSPANElement) IDF(format string, args ...any) *SVGTSPANElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfID(condition bool, s string) *SVGTSPANElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGTSPANElement) IfIDF(condition bool, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGTSPANElement) IDRemove(s string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGTSPANElement) IDRemoveF(format string, args ...any) *SVGTSPANElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGTSPANElement) CLASS(s ...string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) IfCLASS(condition bool, s ...string) *SVGTSPANElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGTSPANElement) CLASSRemove(s ...string) *SVGTSPANElement {
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
func (e *SVGTSPANElement) STYLEF(k string, format string, args ...any) *SVGTSPANElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGTSPANElement) IfSTYLE(condition bool, k string, v string) *SVGTSPANElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGTSPANElement) STYLE(k string, v string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGTSPANElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGTSPANElement) STYLEMap(m map[string]string) *SVGTSPANElement {
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
func (e *SVGTSPANElement) STYLEPairs(pairs ...string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGTSPANElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGTSPANElement) STYLERemove(keys ...string) *SVGTSPANElement {
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

func (e *SVGTSPANElement) DATASTAR_MERGE_STORE(v any) *SVGTSPANElement {
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

func (e *SVGTSPANElement) DATASTAR_REF(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_REF(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGTSPANElement) DATASTAR_REFRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGTSPANElement) DATASTAR_BIND(key string, expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGTSPANElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGTSPANElement) DATASTAR_BINDRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGTSPANElement) DATASTAR_MODEL(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGTSPANElement) DATASTAR_MODELRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGTSPANElement) DATASTAR_TEXT(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGTSPANElement) DATASTAR_TEXTRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGTspanOnMod customDataKeyModifier

// Debounces the event handler
func SVGTspanOnModDebounce(
	d time.Duration,
) SVGTspanOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGTspanOnModThrottle(
	d time.Duration,
) SVGTspanOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGTSPANElement) DATASTAR_ON(key string, expression string, modifiers ...SVGTspanOnMod) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGTspanOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGTspanOnMod) *SVGTSPANElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGTSPANElement) DATASTAR_ONRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGTSPANElement) DATASTAR_FOCUSSet(b bool) *SVGTSPANElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGTSPANElement) DATASTAR_FOCUS() *SVGTSPANElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGTSPANElement) DATASTAR_HEADER(key string, expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGTSPANElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGTSPANElement) DATASTAR_HEADERRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGTSPANElement) DATASTAR_FETCH_URL(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGTSPANElement) DATASTAR_FETCH_URLRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGTSPANElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGTSPANElement) DATASTAR_FETCH_INDICATORRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGTSPANElement) DATASTAR_SHOWSet(b bool) *SVGTSPANElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGTSPANElement) DATASTAR_SHOW() *SVGTSPANElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGTSPANElement) DATASTAR_INTERSECTS(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGTSPANElement) DATASTAR_INTERSECTSRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGTSPANElement) DATASTAR_TELEPORTSet(b bool) *SVGTSPANElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGTSPANElement) DATASTAR_TELEPORT() *SVGTSPANElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGTSPANElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGTSPANElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGTSPANElement) DATASTAR_SCROLL_INTO_VIEW() *SVGTSPANElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGTSPANElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGTSPANElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGTSPANElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGTSPANElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGTSPANElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGTSPANElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
