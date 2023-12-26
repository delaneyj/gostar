// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feBlend is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <feBlend> SVG filter primitive composes two objects together ruled by a
// certain blending mode
// This is similar to what is known from image editing software when blending two
// layers
// The mode is defined by the mode attribute.
type SVGFEBLENDElement struct {
	*Element
}

// Create a new SVGFEBLENDElement element.
// This will create a new element with the tag
// "feBlend" during rendering.
func SVG_FEBLEND(children ...ElementRenderer) *SVGFEBLENDElement {
	e := NewElement("feBlend", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEBLENDElement{Element: e}
}

func (e *SVGFEBLENDElement) Children(children ...ElementRenderer) *SVGFEBLENDElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEBLENDElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEBLENDElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEBLENDElement) Text(text string) *SVGFEBLENDElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEBLENDElement) TextF(format string, args ...any) *SVGFEBLENDElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfText(condition bool, text string) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEBLENDElement) IfTextF(condition bool, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEBLENDElement) Escaped(text string) *SVGFEBLENDElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEBLENDElement) IfEscaped(condition bool, text string) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEBLENDElement) EscapedF(format string, args ...any) *SVGFEBLENDElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEBLENDElement) CustomData(key, value string) *SVGFEBLENDElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEBLENDElement) IfCustomData(condition bool, key, value string) *SVGFEBLENDElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEBLENDElement) CustomDataF(key, format string, args ...any) *SVGFEBLENDElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEBLENDElement) CustomDataRemove(key string) *SVGFEBLENDElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Input for the blending.
func (e *SVGFEBLENDElement) IN(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFEBLENDElement) IfIN(condition bool, s string) *SVGFEBLENDElement {
	if condition {
		e.IN(s)
	}
	return e
}

// Remove the attribute in from the element.
func (e *SVGFEBLENDElement) INRemove(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in")
	return e
}

// Second input for the blending.
func (e *SVGFEBLENDElement) IN2(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in2", s)
	return e
}

func (e *SVGFEBLENDElement) IfIN2(condition bool, s string) *SVGFEBLENDElement {
	if condition {
		e.IN2(s)
	}
	return e
}

// Remove the attribute in2 from the element.
func (e *SVGFEBLENDElement) IN2Remove(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in2")
	return e
}

// The mode used to blend the two inputs together.
func (e *SVGFEBLENDElement) MODE(c SVGFeBlendModeChoice) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mode", string(c))
	return e
}

type SVGFeBlendModeChoice string

const (
	// The input image is placed over the backdrop image, then the parts of the input
	// image that are outside the backdrop are discarded.
	SVGFeBlendMode_normal SVGFeBlendModeChoice = "normal"
	// The input image is multiplied by the backdrop image.
	SVGFeBlendMode_multiply SVGFeBlendModeChoice = "multiply"
	// Multiplies the complements of the backdrop and input image color values, then
	// complements the result.
	SVGFeBlendMode_screen SVGFeBlendModeChoice = "screen"
	// Selects the darker of the backdrop and input image pixels.
	SVGFeBlendMode_darken SVGFeBlendModeChoice = "darken"
	// Selects the lighter of the backdrop and input image pixels.
	SVGFeBlendMode_lighten SVGFeBlendModeChoice = "lighten"
)

// Remove the attribute mode from the element.
func (e *SVGFEBLENDElement) MODERemove(c SVGFeBlendModeChoice) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mode")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEBLENDElement) CLASS(s ...string) *SVGFEBLENDElement {
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

func (e *SVGFEBLENDElement) IfCLASS(condition bool, s ...string) *SVGFEBLENDElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute class from the element.
func (e *SVGFEBLENDElement) CLASSRemove(s ...string) *SVGFEBLENDElement {
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
func (e *SVGFEBLENDElement) ID(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEBLENDElement) IfID(condition bool, s string) *SVGFEBLENDElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute id from the element.
func (e *SVGFEBLENDElement) IDRemove(s string) *SVGFEBLENDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFEBLENDElement) STYLEF(k string, format string, args ...any) *SVGFEBLENDElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEBLENDElement) IfSTYLE(condition bool, k string, v string) *SVGFEBLENDElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEBLENDElement) STYLE(k string, v string) *SVGFEBLENDElement {
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

func (e *SVGFEBLENDElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEBLENDElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEBLENDElement) STYLEMap(m map[string]string) *SVGFEBLENDElement {
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
func (e *SVGFEBLENDElement) STYLEPairs(pairs ...string) *SVGFEBLENDElement {
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

func (e *SVGFEBLENDElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEBLENDElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute style from the element.
func (e *SVGFEBLENDElement) STYLERemove(keys ...string) *SVGFEBLENDElement {
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
