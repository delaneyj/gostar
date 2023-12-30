// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg foreignObject is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <foreignObject> SVG element allows for inclusion of a foreign XML namespace
// which has its graphical content drawn by a different user agent
// The included foreign graphical content is subject to SVG transformations and
// compositing.
type SVGFOREIGNOBJECTElement struct {
	*Element
}

// Create a new SVGFOREIGNOBJECTElement element.
// This will create a new element with the tag
// "foreignObject" during rendering.
func SVG_FOREIGNOBJECT(children ...ElementRenderer) *SVGFOREIGNOBJECTElement {
	e := NewElement("foreignObject", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFOREIGNOBJECTElement{Element: e}
}

func (e *SVGFOREIGNOBJECTElement) Children(children ...ElementRenderer) *SVGFOREIGNOBJECTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) Attr(name string, value string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFOREIGNOBJECTElement) Attrs(attrs ...string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) AttrsMap(attrs map[string]string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) Text(text string) *SVGFOREIGNOBJECTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFOREIGNOBJECTElement) TextF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfText(condition bool, text string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfTextF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) Escaped(text string) *SVGFOREIGNOBJECTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfEscaped(condition bool, text string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) EscapedF(format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfEscapedF(condition bool, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) CustomData(key, value string) *SVGFOREIGNOBJECTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfCustomData(condition bool, key, value string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) CustomDataF(key, format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) CustomDataRemove(key string) *SVGFOREIGNOBJECTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGFOREIGNOBJECTElement) X(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("x", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfX(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.X(s)
	}
	return e
}

// Remove the attribute X from the element.
func (e *SVGFOREIGNOBJECTElement) XRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("x")
	return e
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGFOREIGNOBJECTElement) Y(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("y", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfY(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.Y(s)
	}
	return e
}

// Remove the attribute Y from the element.
func (e *SVGFOREIGNOBJECTElement) YRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("y")
	return e
}

// The width of the rectangular region.
func (e *SVGFOREIGNOBJECTElement) WIDTH(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("width", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfWIDTH(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.WIDTH(s)
	}
	return e
}

// Remove the attribute WIDTH from the element.
func (e *SVGFOREIGNOBJECTElement) WIDTHRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("width")
	return e
}

// The height of the rectangular region.
func (e *SVGFOREIGNOBJECTElement) HEIGHT(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("height", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfHEIGHT(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.HEIGHT(s)
	}
	return e
}

// Remove the attribute HEIGHT from the element.
func (e *SVGFOREIGNOBJECTElement) HEIGHTRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("height")
	return e
}

// A space-separated list of required extensions, indicating that the parent SVG
// document must include the specified extensions for this element to be valid.
func (e *SVGFOREIGNOBJECTElement) REQUIRED_EXTENSIONS(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("requiredExtensions", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfREQUIRED_EXTENSIONS(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.REQUIRED_EXTENSIONS(s)
	}
	return e
}

// Remove the attribute REQUIRED_EXTENSIONS from the element.
func (e *SVGFOREIGNOBJECTElement) REQUIRED_EXTENSIONSRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("requiredExtensions")
	return e
}

// A space-separated list of required features, indicating that the parent SVG
// document must include support for all of the specified features for this
// element to be valid.
func (e *SVGFOREIGNOBJECTElement) REQUIRED_FEATURES(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("requiredFeatures", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfREQUIRED_FEATURES(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.REQUIRED_FEATURES(s)
	}
	return e
}

// Remove the attribute REQUIRED_FEATURES from the element.
func (e *SVGFOREIGNOBJECTElement) REQUIRED_FEATURESRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("requiredFeatures")
	return e
}

// A space-separated list of language codes, indicating that the parent SVG
// document must include support for all of the specified languages for this
// element to be valid.
func (e *SVGFOREIGNOBJECTElement) SYSTEM_LANGUAGE(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("systemLanguage", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfSYSTEM_LANGUAGE(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.SYSTEM_LANGUAGE(s)
	}
	return e
}

// Remove the attribute SYSTEM_LANGUAGE from the element.
func (e *SVGFOREIGNOBJECTElement) SYSTEM_LANGUAGERemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("systemLanguage")
	return e
}

// Specifies a unique id for an element
func (e *SVGFOREIGNOBJECTElement) ID(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfID(condition bool, s string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFOREIGNOBJECTElement) IDRemove(s string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFOREIGNOBJECTElement) CLASS(s ...string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) IfCLASS(condition bool, s ...string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFOREIGNOBJECTElement) CLASSRemove(s ...string) *SVGFOREIGNOBJECTElement {
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
func (e *SVGFOREIGNOBJECTElement) STYLEF(k string, format string, args ...any) *SVGFOREIGNOBJECTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFOREIGNOBJECTElement) IfSTYLE(condition bool, k string, v string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFOREIGNOBJECTElement) STYLE(k string, v string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFOREIGNOBJECTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFOREIGNOBJECTElement) STYLEMap(m map[string]string) *SVGFOREIGNOBJECTElement {
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
func (e *SVGFOREIGNOBJECTElement) STYLEPairs(pairs ...string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFOREIGNOBJECTElement) STYLERemove(keys ...string) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) DATASTAR_MERGE_STORE(v any) *SVGFOREIGNOBJECTElement {
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

func (e *SVGFOREIGNOBJECTElement) DATASTAR_REF(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfDATASTAR_REF(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFOREIGNOBJECTElement) DATASTAR_REFRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFOREIGNOBJECTElement) DATASTAR_BIND(key string, expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFOREIGNOBJECTElement) DATASTAR_BINDRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFOREIGNOBJECTElement) DATASTAR_MODEL(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFOREIGNOBJECTElement) DATASTAR_MODELRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFOREIGNOBJECTElement) DATASTAR_TEXT(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFOREIGNOBJECTElement) DATASTAR_TEXTRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGForeignObjectDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGForeignObjectDataOnModDebounce(
	d time.Duration,
) SVGForeignObjectDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGForeignObjectDataOnModThrottle(
	d time.Duration,
) SVGForeignObjectDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFOREIGNOBJECTElement) DATASTAR_ON(key string, expression string, modifiers ...SVGForeignObjectDataOnMod) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGForeignObjectDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGForeignObjectDataOnMod) *SVGFOREIGNOBJECTElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFOREIGNOBJECTElement) DATASTAR_ONRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFOREIGNOBJECTElement) DATASTAR_FOCUSSet(b bool) *SVGFOREIGNOBJECTElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFOREIGNOBJECTElement) DATASTAR_FOCUS() *SVGFOREIGNOBJECTElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFOREIGNOBJECTElement) DATASTAR_HEADER(key string, expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFOREIGNOBJECTElement) DATASTAR_HEADERRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGFOREIGNOBJECTElement) DATASTAR_FETCH_URL(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGFOREIGNOBJECTElement) DATASTAR_FETCH_URLRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFOREIGNOBJECTElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFOREIGNOBJECTElement) DATASTAR_FETCH_INDICATORRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFOREIGNOBJECTElement) DATASTAR_SHOWSet(b bool) *SVGFOREIGNOBJECTElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFOREIGNOBJECTElement) DATASTAR_SHOW() *SVGFOREIGNOBJECTElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFOREIGNOBJECTElement) DATASTAR_INTERSECTSSet(b bool) *SVGFOREIGNOBJECTElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFOREIGNOBJECTElement) DATASTAR_INTERSECTS() *SVGFOREIGNOBJECTElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *SVGFOREIGNOBJECTElement) DATASTAR_TELEPORTSet(b bool) *SVGFOREIGNOBJECTElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFOREIGNOBJECTElement) DATASTAR_TELEPORT() *SVGFOREIGNOBJECTElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFOREIGNOBJECTElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFOREIGNOBJECTElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFOREIGNOBJECTElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFOREIGNOBJECTElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFOREIGNOBJECTElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFOREIGNOBJECTElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGFOREIGNOBJECTElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFOREIGNOBJECTElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFOREIGNOBJECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
