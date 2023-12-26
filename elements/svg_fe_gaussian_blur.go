// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feGaussianBlur is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <feGaussianBlur> SVG filter primitive blurs the input image by the amount
// specified in stdDeviation, which defines the bell-curve.
type SVGFEGAUSSIANBLURElement struct {
	*Element
}

// Create a new SVGFEGAUSSIANBLURElement element.
// This will create a new element with the tag
// "feGaussianBlur" during rendering.
func SVG_FEGAUSSIANBLUR(children ...ElementRenderer) *SVGFEGAUSSIANBLURElement {
	e := NewElement("feGaussianBlur", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEGAUSSIANBLURElement{Element: e}
}

func (e *SVGFEGAUSSIANBLURElement) Children(children ...ElementRenderer) *SVGFEGAUSSIANBLURElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEGAUSSIANBLURElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEGAUSSIANBLURElement) Text(text string) *SVGFEGAUSSIANBLURElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEGAUSSIANBLURElement) TextF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) Escaped(text string) *SVGFEGAUSSIANBLURElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEGAUSSIANBLURElement) EscapedF(format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) CustomData(key, value string) *SVGFEGAUSSIANBLURElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEGAUSSIANBLURElement) CustomDataF(key, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) CustomDataRemove(key string) *SVGFEGAUSSIANBLURElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The input for this filter.
func (e *SVGFEGAUSSIANBLURElement) IN(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

// Remove the attribute in from the element.
func (e *SVGFEGAUSSIANBLURElement) INRemove(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in")
	return e
}

// The standard deviation for the blur operation
// If two <numbers> are provided, the first number represents a standard deviation
// value along the x-axis of the coordinate system established by attribute
// 'primitiveUnits' on the <filter> element
// The second value represents a standard deviation in Y
// If one number is provided, then that value is used for both X and Y
// Negative values are not allowed
// A value of zero disables the effect of the given filter primitive (i.e., the
// result is a transparent black image).
func (e *SVGFEGAUSSIANBLURElement) STDDEVIATION(f float64) *SVGFEGAUSSIANBLURElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("stdDeviation", f)
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEGAUSSIANBLURElement) CLASS(s ...string) *SVGFEGAUSSIANBLURElement {
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

// Remove the attribute class from the element.
func (e *SVGFEGAUSSIANBLURElement) CLASSRemove(s ...string) *SVGFEGAUSSIANBLURElement {
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
func (e *SVGFEGAUSSIANBLURElement) ID(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGFEGAUSSIANBLURElement) IDRemove(s string) *SVGFEGAUSSIANBLURElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFEGAUSSIANBLURElement) STYLEF(k string, format string, args ...any) *SVGFEGAUSSIANBLURElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEGAUSSIANBLURElement) STYLE(k string, v string) *SVGFEGAUSSIANBLURElement {
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

// Add the attributes in the map to the element.
func (e *SVGFEGAUSSIANBLURElement) STYLEMap(m map[string]string) *SVGFEGAUSSIANBLURElement {
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
func (e *SVGFEGAUSSIANBLURElement) STYLEPairs(pairs ...string) *SVGFEGAUSSIANBLURElement {
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

// Remove the attribute style from the element.
func (e *SVGFEGAUSSIANBLURElement) STYLERemove(keys ...string) *SVGFEGAUSSIANBLURElement {
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
