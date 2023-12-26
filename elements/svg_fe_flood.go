// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feFlood is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <feFlood> SVG filter primitive fills the filter subregion with the color
// and opacity defined by flood-color and flood-opacity.
type SVGFEFLOODElement struct {
	*Element
}

// Create a new SVGFEFLOODElement element.
// This will create a new element with the tag
// "feFlood" during rendering.
func SVG_FEFLOOD(children ...ElementRenderer) *SVGFEFLOODElement {
	e := NewElement("feFlood", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEFLOODElement{Element: e}
}

func (e *SVGFEFLOODElement) Children(children ...ElementRenderer) *SVGFEFLOODElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEFLOODElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEFLOODElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEFLOODElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEFLOODElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEFLOODElement) Text(text string) *SVGFEFLOODElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEFLOODElement) TextF(format string, args ...any) *SVGFEFLOODElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) Escaped(text string) *SVGFEFLOODElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEFLOODElement) EscapedF(format string, args ...any) *SVGFEFLOODElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) CustomData(key, value string) *SVGFEFLOODElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEFLOODElement) CustomDataF(key, format string, args ...any) *SVGFEFLOODElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) CustomDataRemove(key string) *SVGFEFLOODElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The flood-color attribute indicates what color to use to flood the current
// filter primitive subregion defined through the <feFlood> element
// If attribute 'flood-color' is not specified, then the effect is as if a value
// of black were specified.
func (e *SVGFEFLOODElement) FLOOD_COLOR(s string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("flood-color", s)
	return e
}

// Remove the attribute flood-color from the element.
func (e *SVGFEFLOODElement) FLOOD_COLORRemove(s string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("flood-color")
	return e
}

// The flood-opacity attribute indicates the opacity value to use across the
// current filter primitive subregion defined through the <feFlood> element.
func (e *SVGFEFLOODElement) FLOOD_OPACITY(f float64) *SVGFEFLOODElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("flood-opacity", f)
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEFLOODElement) CLASS(s ...string) *SVGFEFLOODElement {
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
func (e *SVGFEFLOODElement) CLASSRemove(s ...string) *SVGFEFLOODElement {
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
func (e *SVGFEFLOODElement) ID(s string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGFEFLOODElement) IDRemove(s string) *SVGFEFLOODElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFEFLOODElement) STYLEF(k string, format string, args ...any) *SVGFEFLOODElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEFLOODElement) STYLE(k string, v string) *SVGFEFLOODElement {
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
func (e *SVGFEFLOODElement) STYLEMap(m map[string]string) *SVGFEFLOODElement {
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
func (e *SVGFEFLOODElement) STYLEPairs(pairs ...string) *SVGFEFLOODElement {
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
func (e *SVGFEFLOODElement) STYLERemove(keys ...string) *SVGFEFLOODElement {
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
