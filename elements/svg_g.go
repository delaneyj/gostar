// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg g is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <g> SVG element is a container used to group other SVG elements.
type SVGGElement struct {
	*Element
}

// Create a new SVGGElement element.
// This will create a new element with the tag
// "g" during rendering.
func SVG_G(children ...ElementRenderer) *SVGGElement {
	e := NewElement("g", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGGElement{Element: e}
}

func (e *SVGGElement) Children(children ...ElementRenderer) *SVGGElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGGElement) IfChildren(condition bool, children ...ElementRenderer) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGGElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGGElement) Text(text string) *SVGGElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGGElement) TextF(format string, args ...any) *SVGGElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfText(condition bool, text string) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGGElement) IfTextF(condition bool, format string, args ...any) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGGElement) Escaped(text string) *SVGGElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGGElement) IfEscaped(condition bool, text string) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGGElement) EscapedF(format string, args ...any) *SVGGElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfEscapedF(condition bool, format string, args ...any) *SVGGElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGGElement) CustomData(key, value string) *SVGGElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGGElement) IfCustomData(condition bool, key, value string) *SVGGElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGGElement) CustomDataF(key, format string, args ...any) *SVGGElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGGElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGGElement) CustomDataRemove(key string) *SVGGElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// A space-separated list of required extensions, indicating that the parent SVG
// document must include the specified extensions for this element to be valid.
func (e *SVGGElement) REQUIRED_EXTENSIONS(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("requiredExtensions", s)
	return e
}

func (e *SVGGElement) IfREQUIRED_EXTENSIONS(condition bool, s string) *SVGGElement {
	if condition {
		e.REQUIRED_EXTENSIONS(s)
	}
	return e
}

// Remove the attribute REQUIRED_EXTENSIONS from the element.
func (e *SVGGElement) REQUIRED_EXTENSIONSRemove(s string) *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("requiredExtensions")
	return e
}

// A space-separated list of required features, indicating that the parent SVG
// document must include support for all of the specified features for this
// element to be valid.
func (e *SVGGElement) REQUIRED_FEATURES(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("requiredFeatures", s)
	return e
}

func (e *SVGGElement) IfREQUIRED_FEATURES(condition bool, s string) *SVGGElement {
	if condition {
		e.REQUIRED_FEATURES(s)
	}
	return e
}

// Remove the attribute REQUIRED_FEATURES from the element.
func (e *SVGGElement) REQUIRED_FEATURESRemove(s string) *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("requiredFeatures")
	return e
}

// A space-separated list of language codes, indicating that the parent SVG
// document must include support for all of the specified languages for this
// element to be valid.
func (e *SVGGElement) SYSTEM_LANGUAGE(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("systemLanguage", s)
	return e
}

func (e *SVGGElement) IfSYSTEM_LANGUAGE(condition bool, s string) *SVGGElement {
	if condition {
		e.SYSTEM_LANGUAGE(s)
	}
	return e
}

// Remove the attribute SYSTEM_LANGUAGE from the element.
func (e *SVGGElement) SYSTEM_LANGUAGERemove(s string) *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("systemLanguage")
	return e
}

// Specifies a unique id for an element
func (e *SVGGElement) ID(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGGElement) IfID(condition bool, s string) *SVGGElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGGElement) IDRemove(s string) *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGGElement) CLASS(s ...string) *SVGGElement {
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

func (e *SVGGElement) IfCLASS(condition bool, s ...string) *SVGGElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGGElement) CLASSRemove(s ...string) *SVGGElement {
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
func (e *SVGGElement) STYLEF(k string, format string, args ...any) *SVGGElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGGElement) IfSTYLE(condition bool, k string, v string) *SVGGElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGGElement) STYLE(k string, v string) *SVGGElement {
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

func (e *SVGGElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGGElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGGElement) STYLEMap(m map[string]string) *SVGGElement {
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
func (e *SVGGElement) STYLEPairs(pairs ...string) *SVGGElement {
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

func (e *SVGGElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGGElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGGElement) STYLERemove(keys ...string) *SVGGElement {
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

func (e *SVGGElement) DATASTAR_MERGE_STORE(v any) *SVGGElement {
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

func (e *SVGGElement) DATASTAR_REF(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-ref"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGGElement) IfDATASTAR_REF(condition bool, s string) *SVGGElement {
	if condition {
		e.DATASTAR_REF(s)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGGElement) DATASTAR_REFRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGGElement) DATASTAR_BIND(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-bind"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGGElement) IfDATASTAR_BIND(condition bool, s string) *SVGGElement {
	if condition {
		e.DATASTAR_BIND(s)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGGElement) DATASTAR_BINDRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGGElement) DATASTAR_MODEL(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-model"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGGElement) IfDATASTAR_MODEL(condition bool, s string) *SVGGElement {
	if condition {
		e.DATASTAR_MODEL(s)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGGElement) DATASTAR_MODELRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGGElement) DATASTAR_TEXT(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-text"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGGElement) IfDATASTAR_TEXT(condition bool, s string) *SVGGElement {
	if condition {
		e.DATASTAR_TEXT(s)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGGElement) DATASTAR_TEXTRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGGDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGGDataOnModDebounce(
	s string,
) SVGGDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%sms", s)
	}
}

// Throttles the event handler
func SVGGDataOnModThrottle(
	s string,
) SVGGDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%sms", s)
	}
}

func (e *SVGGElement) DATASTAR_ON(s string, modifiers ...SVGGDataOnMod) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	customMods := lo.Map(modifiers, func(m SVGGDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key := customDataKey("data-on", customMods...)
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGGElement) IfDATASTAR_ON(condition bool, s string, modifiers ...SVGGDataOnMod) *SVGGElement {
	if condition {
		e.DATASTAR_ON(s, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGGElement) DATASTAR_ONRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGGElement) DATASTAR_FOCUSSet(b bool) *SVGGElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGGElement) DATASTAR_FOCUS() *SVGGElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGGElement) DATASTAR_HEADER(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-header"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGGElement) IfDATASTAR_HEADER(condition bool, s string) *SVGGElement {
	if condition {
		e.DATASTAR_HEADER(s)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGGElement) DATASTAR_HEADERRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGGElement) DATASTAR_FETCH_URL(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-fetch-url"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGGElement) IfDATASTAR_FETCH_URL(condition bool, s string) *SVGGElement {
	if condition {
		e.DATASTAR_FETCH_URL(s)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGGElement) DATASTAR_FETCH_URLRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGGElement) DATASTAR_FETCH_INDICATOR(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "DatastarFetchIndicator"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGGElement) IfDATASTAR_FETCH_INDICATOR(condition bool, s string) *SVGGElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(s)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGGElement) DATASTAR_FETCH_INDICATORRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGGElement) DATASTAR_SHOWSet(b bool) *SVGGElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGGElement) DATASTAR_SHOW() *SVGGElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGGElement) DATASTAR_INTERSECTSSet(b bool) *SVGGElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGGElement) DATASTAR_INTERSECTS() *SVGGElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *SVGGElement) DATASTAR_TELEPORTSet(b bool) *SVGGElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGGElement) DATASTAR_TELEPORT() *SVGGElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGGElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGGElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGGElement) DATASTAR_SCROLL_INTO_VIEW() *SVGGElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGGElement) DATASTAR_VIEW_TRANSITION(s string) *SVGGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-view-transition"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *SVGGElement) IfDATASTAR_VIEW_TRANSITION(condition bool, s string) *SVGGElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(s)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGGElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
