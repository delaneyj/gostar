// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feSpecularLighting is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
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

func (e *SVGFESPECULARLIGHTINGElement) IfIN(condition bool, s string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.IN(s)
	}
	return e
}

// Remove the attribute in from the element.
func (e *SVGFESPECULARLIGHTINGElement) INRemove(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in")
	return e
}

// The 'surfaceScale' attribute indicates the height of the surface when the alpha
// channel is 1.0.
func (e *SVGFESPECULARLIGHTINGElement) SURFACESCALE(f float64) *SVGFESPECULARLIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("surfaceScale", f)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfSURFACESCALE(condition bool, f float64) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.SURFACESCALE(f)
	}
	return e
}

// The specularConstant attribute represents the diffuse reflection constant.
func (e *SVGFESPECULARLIGHTINGElement) SPECULARCONSTANT(f float64) *SVGFESPECULARLIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("specularConstant", f)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfSPECULARCONSTANT(condition bool, f float64) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.SPECULARCONSTANT(f)
	}
	return e
}

// The specularExponent attribute represents the specular reflection constant.
func (e *SVGFESPECULARLIGHTINGElement) SPECULAREXPONENT(f float64) *SVGFESPECULARLIGHTINGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("specularExponent", f)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfSPECULAREXPONENT(condition bool, f float64) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.SPECULAREXPONENT(f)
	}
	return e
}

// The kernelUnitLength attribute defines the intended distance in current filter
// units (i.e., units as determined by the value of attribute 'primitiveUnits')
// for dx and dy in the surface normal calculation formulas.
func (e *SVGFESPECULARLIGHTINGElement) KERNELUNITLENGTH(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("kernelUnitLength", s)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfKERNELUNITLENGTH(condition bool, s string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.KERNELUNITLENGTH(s)
	}
	return e
}

// Remove the attribute kernelUnitLength from the element.
func (e *SVGFESPECULARLIGHTINGElement) KERNELUNITLENGTHRemove(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("kernelUnitLength")
	return e
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

// Remove the attribute class from the element.
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

// Specifies a unique id for an element
func (e *SVGFESPECULARLIGHTINGElement) ID(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFESPECULARLIGHTINGElement) IfID(condition bool, s string) *SVGFESPECULARLIGHTINGElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute id from the element.
func (e *SVGFESPECULARLIGHTINGElement) IDRemove(s string) *SVGFESPECULARLIGHTINGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
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

// Remove the attribute style from the element.
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
