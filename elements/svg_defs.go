// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg defs is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <defs> SVG element is used to embed definitions that can be reused inside
// an SVG image.
type SVGDEFSElement struct {
	*Element
}

// Create a new SVGDEFSElement element.
// This will create a new element with the tag
// "defs" during rendering.
func SVG_DEFS(children ...ElementRenderer) *SVGDEFSElement {
	e := NewElement("defs", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGDEFSElement{Element: e}
}

func (e *SVGDEFSElement) Children(children ...ElementRenderer) *SVGDEFSElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGDEFSElement) IfChildren(condition bool, children ...ElementRenderer) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGDEFSElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGDEFSElement) Text(text string) *SVGDEFSElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGDEFSElement) TextF(format string, args ...any) *SVGDEFSElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGDEFSElement) IfText(condition bool, text string) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGDEFSElement) IfTextF(condition bool, format string, args ...any) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGDEFSElement) Escaped(text string) *SVGDEFSElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGDEFSElement) IfEscaped(condition bool, text string) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGDEFSElement) EscapedF(format string, args ...any) *SVGDEFSElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGDEFSElement) IfEscapedF(condition bool, format string, args ...any) *SVGDEFSElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGDEFSElement) CustomData(key, value string) *SVGDEFSElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGDEFSElement) IfCustomData(condition bool, key, value string) *SVGDEFSElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGDEFSElement) CustomDataF(key, format string, args ...any) *SVGDEFSElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGDEFSElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGDEFSElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGDEFSElement) CustomDataRemove(key string) *SVGDEFSElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGDEFSElement) CLASS(s ...string) *SVGDEFSElement {
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

func (e *SVGDEFSElement) IfCLASS(condition bool, s ...string) *SVGDEFSElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute class from the element.
func (e *SVGDEFSElement) CLASSRemove(s ...string) *SVGDEFSElement {
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
func (e *SVGDEFSElement) ID(s string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGDEFSElement) IfID(condition bool, s string) *SVGDEFSElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute id from the element.
func (e *SVGDEFSElement) IDRemove(s string) *SVGDEFSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGDEFSElement) STYLEF(k string, format string, args ...any) *SVGDEFSElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGDEFSElement) IfSTYLE(condition bool, k string, v string) *SVGDEFSElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGDEFSElement) STYLE(k string, v string) *SVGDEFSElement {
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

func (e *SVGDEFSElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGDEFSElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGDEFSElement) STYLEMap(m map[string]string) *SVGDEFSElement {
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
func (e *SVGDEFSElement) STYLEPairs(pairs ...string) *SVGDEFSElement {
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

func (e *SVGDEFSElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGDEFSElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute style from the element.
func (e *SVGDEFSElement) STYLERemove(keys ...string) *SVGDEFSElement {
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
