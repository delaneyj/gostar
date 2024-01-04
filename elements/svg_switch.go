// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg switch is generated from configuration file.
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

// The <switch> SVG element evaluates the requiredFeatures, requiredExtensions and
// systemLanguage attributes on its direct child elements in order, and then
// processes and renders the first child for which these attributes evaluate to
// true
// All others will be bypassed and therefore not rendered
// If the child element is a container element such as a <g>, then the entire
// subtree is either processed/rendered or bypassed/not rendered.
type SVGSWITCHElement struct {
	*Element
}

// Create a new SVGSWITCHElement element.
// This will create a new element with the tag
// "switch" during rendering.
func SVG_SWITCH(children ...ElementRenderer) *SVGSWITCHElement {
	e := NewElement("switch", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSWITCHElement{Element: e}
}

func (e *SVGSWITCHElement) Children(children ...ElementRenderer) *SVGSWITCHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSWITCHElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSWITCHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSWITCHElement) Attr(name string, value string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGSWITCHElement) Attrs(attrs ...string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) AttrsMap(attrs map[string]string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGSWITCHElement) Text(text string) *SVGSWITCHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSWITCHElement) TextF(format string, args ...any) *SVGSWITCHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfText(condition bool, text string) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSWITCHElement) IfTextF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSWITCHElement) Escaped(text string) *SVGSWITCHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSWITCHElement) IfEscaped(condition bool, text string) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSWITCHElement) EscapedF(format string, args ...any) *SVGSWITCHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfEscapedF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSWITCHElement) CustomData(key, value string) *SVGSWITCHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSWITCHElement) IfCustomData(condition bool, key, value string) *SVGSWITCHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSWITCHElement) CustomDataF(key, format string, args ...any) *SVGSWITCHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSWITCHElement) CustomDataRemove(key string) *SVGSWITCHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// A space-separated list of required features, indicating that the parent SVG
// document must include support for all of the specified features for this
// element to be valid.
func (e *SVGSWITCHElement) REQUIRED_FEATURES(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("requiredFeatures", s)
	return e
}

func (e *SVGSWITCHElement) REQUIRED_FEATURESF(format string, args ...any) *SVGSWITCHElement {
	return e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfREQUIRED_FEATURES(condition bool, s string) *SVGSWITCHElement {
	if condition {
		e.REQUIRED_FEATURES(s)
	}
	return e
}

func (e *SVGSWITCHElement) IfREQUIRED_FEATURESF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.REQUIRED_FEATURES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_FEATURES from the element.
func (e *SVGSWITCHElement) REQUIRED_FEATURESRemove(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("requiredFeatures")
	return e
}

func (e *SVGSWITCHElement) REQUIRED_FEATURESRemoveF(format string, args ...any) *SVGSWITCHElement {
	return e.REQUIRED_FEATURESRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of required extensions, indicating that the parent SVG
// document must include the specified extensions for this element to be valid.
func (e *SVGSWITCHElement) REQUIRED_EXTENSIONS(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("requiredExtensions", s)
	return e
}

func (e *SVGSWITCHElement) REQUIRED_EXTENSIONSF(format string, args ...any) *SVGSWITCHElement {
	return e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfREQUIRED_EXTENSIONS(condition bool, s string) *SVGSWITCHElement {
	if condition {
		e.REQUIRED_EXTENSIONS(s)
	}
	return e
}

func (e *SVGSWITCHElement) IfREQUIRED_EXTENSIONSF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.REQUIRED_EXTENSIONS(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REQUIRED_EXTENSIONS from the element.
func (e *SVGSWITCHElement) REQUIRED_EXTENSIONSRemove(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("requiredExtensions")
	return e
}

func (e *SVGSWITCHElement) REQUIRED_EXTENSIONSRemoveF(format string, args ...any) *SVGSWITCHElement {
	return e.REQUIRED_EXTENSIONSRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of language codes, indicating that the parent SVG
// document must include support for all of the specified languages for this
// element to be valid.
func (e *SVGSWITCHElement) SYSTEM_LANGUAGE(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("systemLanguage", s)
	return e
}

func (e *SVGSWITCHElement) SYSTEM_LANGUAGEF(format string, args ...any) *SVGSWITCHElement {
	return e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfSYSTEM_LANGUAGE(condition bool, s string) *SVGSWITCHElement {
	if condition {
		e.SYSTEM_LANGUAGE(s)
	}
	return e
}

func (e *SVGSWITCHElement) IfSYSTEM_LANGUAGEF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.SYSTEM_LANGUAGE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute SYSTEM_LANGUAGE from the element.
func (e *SVGSWITCHElement) SYSTEM_LANGUAGERemove(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("systemLanguage")
	return e
}

func (e *SVGSWITCHElement) SYSTEM_LANGUAGERemoveF(format string, args ...any) *SVGSWITCHElement {
	return e.SYSTEM_LANGUAGERemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGSWITCHElement) ID(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSWITCHElement) IDF(format string, args ...any) *SVGSWITCHElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfID(condition bool, s string) *SVGSWITCHElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGSWITCHElement) IfIDF(condition bool, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGSWITCHElement) IDRemove(s string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGSWITCHElement) IDRemoveF(format string, args ...any) *SVGSWITCHElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSWITCHElement) CLASS(s ...string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) IfCLASS(condition bool, s ...string) *SVGSWITCHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGSWITCHElement) CLASSRemove(s ...string) *SVGSWITCHElement {
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
func (e *SVGSWITCHElement) STYLEF(k string, format string, args ...any) *SVGSWITCHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSWITCHElement) IfSTYLE(condition bool, k string, v string) *SVGSWITCHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSWITCHElement) STYLE(k string, v string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSWITCHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSWITCHElement) STYLEMap(m map[string]string) *SVGSWITCHElement {
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
func (e *SVGSWITCHElement) STYLEPairs(pairs ...string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSWITCHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGSWITCHElement) STYLERemove(keys ...string) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) DATASTAR_MERGE_STORE(v any) *SVGSWITCHElement {
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

func (e *SVGSWITCHElement) DATASTAR_REF(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_REF(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGSWITCHElement) DATASTAR_REFRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGSWITCHElement) DATASTAR_BIND(key string, expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGSWITCHElement) DATASTAR_BINDRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGSWITCHElement) DATASTAR_MODEL(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGSWITCHElement) DATASTAR_MODELRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGSWITCHElement) DATASTAR_TEXT(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGSWITCHElement) DATASTAR_TEXTRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGSwitchOnMod customDataKeyModifier

// Debounces the event handler
func SVGSwitchOnModDebounce(
	d time.Duration,
) SVGSwitchOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGSwitchOnModThrottle(
	d time.Duration,
) SVGSwitchOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGSWITCHElement) DATASTAR_ON(key string, expression string, modifiers ...SVGSwitchOnMod) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGSwitchOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGSwitchOnMod) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGSWITCHElement) DATASTAR_ONRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGSWITCHElement) DATASTAR_FOCUSSet(b bool) *SVGSWITCHElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSWITCHElement) DATASTAR_FOCUS() *SVGSWITCHElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGSWITCHElement) DATASTAR_HEADER(key string, expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGSWITCHElement) DATASTAR_HEADERRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGSWITCHElement) DATASTAR_FETCH_URL(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGSWITCHElement) DATASTAR_FETCH_URLRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGSWITCHElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGSWITCHElement) DATASTAR_FETCH_INDICATORRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGSWITCHElement) DATASTAR_SHOWSet(b bool) *SVGSWITCHElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSWITCHElement) DATASTAR_SHOW() *SVGSWITCHElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGSWITCHElement) DATASTAR_INTERSECTS(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGSWITCHElement) DATASTAR_INTERSECTSRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGSWITCHElement) DATASTAR_TELEPORTSet(b bool) *SVGSWITCHElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSWITCHElement) DATASTAR_TELEPORT() *SVGSWITCHElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGSWITCHElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGSWITCHElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGSWITCHElement) DATASTAR_SCROLL_INTO_VIEW() *SVGSWITCHElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGSWITCHElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGSWITCHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGSWITCHElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGSWITCHElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGSWITCHElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGSWITCHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
