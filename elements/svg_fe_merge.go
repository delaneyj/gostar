// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feMerge is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <feMerge> SVG element allows filter effects to be applied concurrently
// instead of sequentially.
type SVGFEMERGEElement struct {
	*Element
}

// Create a new SVGFEMERGEElement element.
// This will create a new element with the tag
// "feMerge" during rendering.
func SVG_FEMERGE(children ...ElementRenderer) *SVGFEMERGEElement {
	e := NewElement("feMerge", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEMERGEElement{Element: e}
}

func (e *SVGFEMERGEElement) Children(children ...ElementRenderer) *SVGFEMERGEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEMERGEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEMERGEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEMERGEElement) Attr(name string, value string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFEMERGEElement) Attrs(attrs ...string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) AttrsMap(attrs map[string]string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEMERGEElement) Text(text string) *SVGFEMERGEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEMERGEElement) TextF(format string, args ...any) *SVGFEMERGEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfText(condition bool, text string) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEMERGEElement) IfTextF(condition bool, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEMERGEElement) Escaped(text string) *SVGFEMERGEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEMERGEElement) IfEscaped(condition bool, text string) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEMERGEElement) EscapedF(format string, args ...any) *SVGFEMERGEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEMERGEElement) CustomData(key, value string) *SVGFEMERGEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEMERGEElement) IfCustomData(condition bool, key, value string) *SVGFEMERGEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEMERGEElement) CustomDataF(key, format string, args ...any) *SVGFEMERGEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEMERGEElement) CustomDataRemove(key string) *SVGFEMERGEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Specifies a unique id for an element
func (e *SVGFEMERGEElement) ID(s string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEMERGEElement) IDF(format string, args ...any) *SVGFEMERGEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfID(condition bool, s string) *SVGFEMERGEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEMERGEElement) IfIDF(condition bool, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEMERGEElement) IDRemove(s string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGFEMERGEElement) IDRemoveF(format string, args ...any) *SVGFEMERGEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEMERGEElement) CLASS(s ...string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) IfCLASS(condition bool, s ...string) *SVGFEMERGEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEMERGEElement) CLASSRemove(s ...string) *SVGFEMERGEElement {
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
func (e *SVGFEMERGEElement) STYLEF(k string, format string, args ...any) *SVGFEMERGEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGEElement) IfSTYLE(condition bool, k string, v string) *SVGFEMERGEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEMERGEElement) STYLE(k string, v string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEMERGEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEMERGEElement) STYLEMap(m map[string]string) *SVGFEMERGEElement {
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
func (e *SVGFEMERGEElement) STYLEPairs(pairs ...string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEMERGEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEMERGEElement) STYLERemove(keys ...string) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) DATASTAR_MERGE_STORE(v any) *SVGFEMERGEElement {
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

func (e *SVGFEMERGEElement) DATASTAR_REF(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_REF(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFEMERGEElement) DATASTAR_REFRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFEMERGEElement) DATASTAR_BIND(key string, expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFEMERGEElement) DATASTAR_BINDRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFEMERGEElement) DATASTAR_MODEL(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFEMERGEElement) DATASTAR_MODELRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFEMERGEElement) DATASTAR_TEXT(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFEMERGEElement) DATASTAR_TEXTRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeMergeOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeMergeOnModDebounce(
	d time.Duration,
) SVGFeMergeOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeMergeOnModThrottle(
	d time.Duration,
) SVGFeMergeOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFEMERGEElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeMergeOnMod) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeMergeOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeMergeOnMod) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFEMERGEElement) DATASTAR_ONRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFEMERGEElement) DATASTAR_FOCUSSet(b bool) *SVGFEMERGEElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEMERGEElement) DATASTAR_FOCUS() *SVGFEMERGEElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFEMERGEElement) DATASTAR_HEADER(key string, expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFEMERGEElement) DATASTAR_HEADERRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGFEMERGEElement) DATASTAR_FETCH_URL(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGFEMERGEElement) DATASTAR_FETCH_URLRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFEMERGEElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFEMERGEElement) DATASTAR_FETCH_INDICATORRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFEMERGEElement) DATASTAR_SHOWSet(b bool) *SVGFEMERGEElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEMERGEElement) DATASTAR_SHOW() *SVGFEMERGEElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFEMERGEElement) DATASTAR_INTERSECTS(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFEMERGEElement) DATASTAR_INTERSECTSRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFEMERGEElement) DATASTAR_TELEPORTSet(b bool) *SVGFEMERGEElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEMERGEElement) DATASTAR_TELEPORT() *SVGFEMERGEElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFEMERGEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFEMERGEElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEMERGEElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFEMERGEElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFEMERGEElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEMERGEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFEMERGEElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFEMERGEElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFEMERGEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
