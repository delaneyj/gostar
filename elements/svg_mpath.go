// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg mpath is generated from configuration file.
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

// The <mpath> SVG element allows to use the functionality of <animateMotion> to
// animate the <startOffset> attribute of SVG <textPath> elements.
type SVGMPATHElement struct {
	*Element
}

// Create a new SVGMPATHElement element.
// This will create a new element with the tag
// "mpath" during rendering.
func SVG_MPATH(children ...ElementRenderer) *SVGMPATHElement {
	e := NewElement("mpath", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGMPATHElement{Element: e}
}

func (e *SVGMPATHElement) Children(children ...ElementRenderer) *SVGMPATHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGMPATHElement) IfChildren(condition bool, children ...ElementRenderer) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGMPATHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGMPATHElement) Attr(name string, value string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGMPATHElement) Attrs(attrs ...string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) AttrsMap(attrs map[string]string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGMPATHElement) Text(text string) *SVGMPATHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGMPATHElement) TextF(format string, args ...any) *SVGMPATHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfText(condition bool, text string) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGMPATHElement) IfTextF(condition bool, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGMPATHElement) Escaped(text string) *SVGMPATHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGMPATHElement) IfEscaped(condition bool, text string) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGMPATHElement) EscapedF(format string, args ...any) *SVGMPATHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfEscapedF(condition bool, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGMPATHElement) CustomData(key, value string) *SVGMPATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGMPATHElement) IfCustomData(condition bool, key, value string) *SVGMPATHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGMPATHElement) CustomDataF(key, format string, args ...any) *SVGMPATHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGMPATHElement) CustomDataRemove(key string) *SVGMPATHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// A URI reference to the motion path definition.
func (e *SVGMPATHElement) HREF(s string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGMPATHElement) HREFF(format string, args ...any) *SVGMPATHElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfHREF(condition bool, s string) *SVGMPATHElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGMPATHElement) IfHREFF(condition bool, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGMPATHElement) HREFRemove(s string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("href")
	return e
}

func (e *SVGMPATHElement) HREFRemoveF(format string, args ...any) *SVGMPATHElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGMPATHElement) ID(s string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGMPATHElement) IDF(format string, args ...any) *SVGMPATHElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfID(condition bool, s string) *SVGMPATHElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGMPATHElement) IfIDF(condition bool, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGMPATHElement) IDRemove(s string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGMPATHElement) IDRemoveF(format string, args ...any) *SVGMPATHElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGMPATHElement) CLASS(s ...string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) IfCLASS(condition bool, s ...string) *SVGMPATHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGMPATHElement) CLASSRemove(s ...string) *SVGMPATHElement {
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
func (e *SVGMPATHElement) STYLEF(k string, format string, args ...any) *SVGMPATHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGMPATHElement) IfSTYLE(condition bool, k string, v string) *SVGMPATHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGMPATHElement) STYLE(k string, v string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGMPATHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGMPATHElement) STYLEMap(m map[string]string) *SVGMPATHElement {
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
func (e *SVGMPATHElement) STYLEPairs(pairs ...string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGMPATHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGMPATHElement) STYLERemove(keys ...string) *SVGMPATHElement {
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

func (e *SVGMPATHElement) DATASTAR_MERGE_STORE(v any) *SVGMPATHElement {
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

func (e *SVGMPATHElement) DATASTAR_REF(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfDATASTAR_REF(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGMPATHElement) DATASTAR_REFRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGMPATHElement) DATASTAR_BIND(key string, expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGMPATHElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGMPATHElement) DATASTAR_BINDRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGMPATHElement) DATASTAR_MODEL(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGMPATHElement) DATASTAR_MODELRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGMPATHElement) DATASTAR_TEXT(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGMPATHElement) DATASTAR_TEXTRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGMpathOnMod customDataKeyModifier

// Debounces the event handler
func SVGMpathOnModDebounce(
	d time.Duration,
) SVGMpathOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGMpathOnModThrottle(
	d time.Duration,
) SVGMpathOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGMPATHElement) DATASTAR_ON(key string, expression string, modifiers ...SVGMpathOnMod) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGMpathOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGMpathOnMod) *SVGMPATHElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGMPATHElement) DATASTAR_ONRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGMPATHElement) DATASTAR_FOCUSSet(b bool) *SVGMPATHElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGMPATHElement) DATASTAR_FOCUS() *SVGMPATHElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGMPATHElement) DATASTAR_HEADER(key string, expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGMPATHElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGMPATHElement) DATASTAR_HEADERRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGMPATHElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGMPATHElement) DATASTAR_FETCH_INDICATORRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGMPATHElement) DATASTAR_SHOWSet(b bool) *SVGMPATHElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGMPATHElement) DATASTAR_SHOW() *SVGMPATHElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGMPATHElement) DATASTAR_INTERSECTS(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGMPATHElement) DATASTAR_INTERSECTSRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGMPATHElement) DATASTAR_TELEPORTSet(b bool) *SVGMPATHElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGMPATHElement) DATASTAR_TELEPORT() *SVGMPATHElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGMPATHElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGMPATHElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGMPATHElement) DATASTAR_SCROLL_INTO_VIEW() *SVGMPATHElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGMPATHElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGMPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMPATHElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGMPATHElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGMPATHElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGMPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
