// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg path is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <path> SVG element is the generic element to define a shape
// All the basic shapes can be created with a path element.
type SVGPATHElement struct {
	*Element
}

// Create a new SVGPATHElement element.
// This will create a new element with the tag
// "path" during rendering.
func SVG_PATH(children ...ElementRenderer) *SVGPATHElement {
	e := NewElement("path", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGPATHElement{Element: e}
}

func (e *SVGPATHElement) Children(children ...ElementRenderer) *SVGPATHElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGPATHElement) IfChildren(condition bool, children ...ElementRenderer) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGPATHElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGPATHElement) Text(text string) *SVGPATHElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGPATHElement) TextF(format string, args ...any) *SVGPATHElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfText(condition bool, text string) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGPATHElement) IfTextF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGPATHElement) Escaped(text string) *SVGPATHElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGPATHElement) IfEscaped(condition bool, text string) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGPATHElement) EscapedF(format string, args ...any) *SVGPATHElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfEscapedF(condition bool, format string, args ...any) *SVGPATHElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGPATHElement) CustomData(key, value string) *SVGPATHElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGPATHElement) IfCustomData(condition bool, key, value string) *SVGPATHElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGPATHElement) CustomDataF(key, format string, args ...any) *SVGPATHElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGPATHElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGPATHElement) CustomDataRemove(key string) *SVGPATHElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The definition of the outline of a shape.
func (e *SVGPATHElement) D(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("d", s)
	return e
}

func (e *SVGPATHElement) IfD(condition bool, s string) *SVGPATHElement {
	if condition {
		e.D(s)
	}
	return e
}

// Remove the attribute d from the element.
func (e *SVGPATHElement) DRemove(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("d")
	return e
}

// The total length for the path, in user units.
func (e *SVGPATHElement) PATHLENGTH(f float64) *SVGPATHElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("pathLength", f)
	return e
}

func (e *SVGPATHElement) IfPATHLENGTH(condition bool, f float64) *SVGPATHElement {
	if condition {
		e.PATHLENGTH(f)
	}
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGPATHElement) CLASS(s ...string) *SVGPATHElement {
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

func (e *SVGPATHElement) IfCLASS(condition bool, s ...string) *SVGPATHElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute class from the element.
func (e *SVGPATHElement) CLASSRemove(s ...string) *SVGPATHElement {
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
func (e *SVGPATHElement) ID(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGPATHElement) IfID(condition bool, s string) *SVGPATHElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute id from the element.
func (e *SVGPATHElement) IDRemove(s string) *SVGPATHElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGPATHElement) STYLEF(k string, format string, args ...any) *SVGPATHElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGPATHElement) IfSTYLE(condition bool, k string, v string) *SVGPATHElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGPATHElement) STYLE(k string, v string) *SVGPATHElement {
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

func (e *SVGPATHElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGPATHElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGPATHElement) STYLEMap(m map[string]string) *SVGPATHElement {
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
func (e *SVGPATHElement) STYLEPairs(pairs ...string) *SVGPATHElement {
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

func (e *SVGPATHElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGPATHElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute style from the element.
func (e *SVGPATHElement) STYLERemove(keys ...string) *SVGPATHElement {
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
