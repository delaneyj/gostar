// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feSpecularLighting is generated from configuration file.
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

// The <feSpecularLighting> SVG filter primitive lights a source graphic using the
// alpha channel as a bump map
// The resulting image is an RGBA image based on the light color
// The lighting calculation follows the standard specular component of the Phong
// lighting model
// The resulting image depends on the light color, light position and surface
// geometry of the input bump map.
type SVGFESPECULARLIGHTINGElement struct {
	*Element
}

// Create a new SVGFESPECULARLIGHTINGElement element.
// This will create a new element with the tag
// "feSpecularLighting" during rendering.
func SVG_FESPECULARLIGHTING(children ...ElementRenderer) *SVGFESPECULARLIGHTINGElement {
	e := NewElement("feSpecularLighting", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFESPECULARLIGHTINGElement{Element: e}
}

func (e *SVGFESPECULARLIGHTINGElement) Children(children ...ElementRenderer) *SVGFESPECULARLIGHTINGElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) Attr(name string, value string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) Attrs(attrs ...string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) AttrsMap(attrs map[string]string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) Text(text string) *SVGFESPECULARLIGHTINGElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) TextF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfText(condition bool, text string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfTextF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) Escaped(text string) *SVGFESPECULARLIGHTINGElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfEscaped(condition bool, text string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) EscapedF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfEscapedF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) CustomData(key, value string) *SVGFESPECULARLIGHTINGElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfCustomData(condition bool, key, value string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) CustomDataF(key, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) CustomDataRemove(key string) *SVGFESPECULARLIGHTINGElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The input for this filter.
func (e *SVGFESPECULARLIGHTINGElement) IN(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) INF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfIN(condition bool, s string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfINF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFESPECULARLIGHTINGElement) INRemove(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in")
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) INRemoveF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// The 'surfaceScale' attribute indicates the height of the surface when the alpha
// channel is 1.0.
func (e *SVGFESPECULARLIGHTINGElement) SURFACE_SCALE(f float64) *SVGFESPECULARLIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("surfaceScale", f)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfSURFACE_SCALE(condition bool, f float64) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.SURFACE_SCALE(f)
	}
	return e
}

// The specularConstant attribute represents the diffuse reflection constant.
func (e *SVGFESPECULARLIGHTINGElement) SPECULAR_CONSTANT(f float64) *SVGFESPECULARLIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("specularConstant", f)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfSPECULAR_CONSTANT(condition bool, f float64) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.SPECULAR_CONSTANT(f)
	}
	return e
}

// The specularExponent attribute represents the specular reflection constant.
func (e *SVGFESPECULARLIGHTINGElement) SPECULAR_EXPONENT(f float64) *SVGFESPECULARLIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("specularExponent", f)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfSPECULAR_EXPONENT(condition bool, f float64) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.SPECULAR_EXPONENT(f)
	}
	return e
}

// The kernelUnitLength attribute defines the intended distance in current filter
// units (i.e., units as determined by the value of attribute 'primitiveUnits')
// for dx and dy in the surface normal calculation formulas.
func (e *SVGFESPECULARLIGHTINGElement) KERNEL_UNIT_LENGTH(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("kernelUnitLength", s)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) KERNEL_UNIT_LENGTHF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.KERNEL_UNIT_LENGTH(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfKERNEL_UNIT_LENGTH(condition bool, s string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.KERNEL_UNIT_LENGTH(s)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfKERNEL_UNIT_LENGTHF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.KERNEL_UNIT_LENGTH(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute KERNEL_UNIT_LENGTH from the element.
func (e *SVGFESPECULARLIGHTINGElement) KERNEL_UNIT_LENGTHRemove(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("kernelUnitLength")
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) KERNEL_UNIT_LENGTHRemoveF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.KERNEL_UNIT_LENGTHRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFESPECULARLIGHTINGElement) ID(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IDF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfID(condition bool, s string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfIDF(condition bool, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFESPECULARLIGHTINGElement) IDRemove(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IDRemoveF(format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFESPECULARLIGHTINGElement) CLASS(s ...string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) IfCLASS(condition bool, s ...string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFESPECULARLIGHTINGElement) CLASSRemove(s ...string) *SVGFESPECULARLIGHTINGElement {
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
func (e *SVGFESPECULARLIGHTINGElement) STYLEF(k string, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFESPECULARLIGHTINGElement) IfSTYLE(condition bool, k string, v string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) STYLE(k string, v string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFESPECULARLIGHTINGElement) STYLEMap(m map[string]string) *SVGFESPECULARLIGHTINGElement {
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
func (e *SVGFESPECULARLIGHTINGElement) STYLEPairs(pairs ...string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFESPECULARLIGHTINGElement) STYLERemove(keys ...string) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_MERGE_STORE(v any) *SVGFESPECULARLIGHTINGElement {
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

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_REF(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfDATASTAR_REF(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_REFRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_BIND(key string, expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_BINDRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_MODEL(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_MODELRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_TEXT(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_TEXTRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeSpecularLightingOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeSpecularLightingOnModDebounce(
	d time.Duration,
) SVGFeSpecularLightingOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeSpecularLightingOnModThrottle(
	d time.Duration,
) SVGFeSpecularLightingOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeSpecularLightingOnMod) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeSpecularLightingOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeSpecularLightingOnMod) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_ONRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_FOCUSSet(b bool) *SVGFESPECULARLIGHTINGElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_FOCUS() *SVGFESPECULARLIGHTINGElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_HEADER(key string, expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_HEADERRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_FETCH_INDICATORRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_SHOWSet(b bool) *SVGFESPECULARLIGHTINGElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_SHOW() *SVGFESPECULARLIGHTINGElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_INTERSECTS(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_INTERSECTSRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_TELEPORTSet(b bool) *SVGFESPECULARLIGHTINGElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_TELEPORT() *SVGFESPECULARLIGHTINGElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFESPECULARLIGHTINGElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFESPECULARLIGHTINGElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFESPECULARLIGHTINGElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
