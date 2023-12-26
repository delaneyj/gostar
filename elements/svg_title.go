// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg title is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <title> SVG element represents a text container used for the context
// information
// This element is usually nested inside a <desc> element.
type SVGTITLEElement struct {
	*Element
}

// Create a new SVGTITLEElement element.
// This will create a new element with the tag
// "title" during rendering.
func SVG_TITLE(children ...ElementRenderer) *SVGTITLEElement {
	e := NewElement("title", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGTITLEElement{Element: e}
}

func (e *SVGTITLEElement) Children(children ...ElementRenderer) *SVGTITLEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGTITLEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGTITLEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGTITLEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGTITLEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGTITLEElement) Text(text string) *SVGTITLEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGTITLEElement) TextF(format string, args ...any) *SVGTITLEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGTITLEElement) Escaped(text string) *SVGTITLEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGTITLEElement) EscapedF(format string, args ...any) *SVGTITLEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGTITLEElement) CustomData(key, value string) *SVGTITLEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGTITLEElement) CustomDataF(key, format string, args ...any) *SVGTITLEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGTITLEElement) CustomDataRemove(key string) *SVGTITLEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGTITLEElement) CLASS(s ...string) *SVGTITLEElement {
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
func (e *SVGTITLEElement) CLASSRemove(s ...string) *SVGTITLEElement {
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
func (e *SVGTITLEElement) ID(s string) *SVGTITLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGTITLEElement) IDRemove(s string) *SVGTITLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGTITLEElement) STYLEF(k string, format string, args ...any) *SVGTITLEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGTITLEElement) STYLE(k string, v string) *SVGTITLEElement {
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
func (e *SVGTITLEElement) STYLEMap(m map[string]string) *SVGTITLEElement {
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
func (e *SVGTITLEElement) STYLEPairs(pairs ...string) *SVGTITLEElement {
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
func (e *SVGTITLEElement) STYLERemove(keys ...string) *SVGTITLEElement {
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
