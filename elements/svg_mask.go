// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg mask is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <mask> SVG element hides portions of SVG elements for user display.
type SVGMASKElement struct {
	*Element
}

// Create a new SVGMASKElement element.
// This will create a new element with the tag
// "mask" during rendering.
func SVG_MASK(children ...ElementRenderer) *SVGMASKElement {
	e := NewElement("mask", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGMASKElement{Element: e}
}

func (e *SVGMASKElement) Children(children ...ElementRenderer) *SVGMASKElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGMASKElement) IfChildren(condition bool, children ...ElementRenderer) *SVGMASKElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGMASKElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGMASKElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGMASKElement) Text(text string) *SVGMASKElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGMASKElement) TextF(format string, args ...any) *SVGMASKElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) Escaped(text string) *SVGMASKElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGMASKElement) EscapedF(format string, args ...any) *SVGMASKElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGMASKElement) CustomData(key, value string) *SVGMASKElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGMASKElement) CustomDataRemove(key string) *SVGMASKElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The coordinate system for attributes x, y, width and height.
func (e *SVGMASKElement) MASKCONTENTUNITS(c SVGMaskMaskContentUnitsChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("maskContentUnits", string(c))
	return e
}

type SVGMaskMaskContentUnitsChoice string

const (
	// The coordinate system for attributes x, y, width and height.
	SVGMaskMaskContentUnits_userSpaceOnUse SVGMaskMaskContentUnitsChoice = "userSpaceOnUse"
	// The coordinate system for attributes x, y, width and height.
	SVGMaskMaskContentUnits_objectBoundingBox SVGMaskMaskContentUnitsChoice = "objectBoundingBox"
)

// Remove the attribute maskContentUnits from the element.
func (e *SVGMASKElement) MASKCONTENTUNITSRemove(c SVGMaskMaskContentUnitsChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("maskContentUnits")
	return e
}

// The coordinate system for the various length values within the filter.
func (e *SVGMASKElement) MASKUNITS(c SVGMaskMaskUnitsChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("maskUnits", string(c))
	return e
}

type SVGMaskMaskUnitsChoice string

const (
	// The coordinate system for the various length values within the filter.
	SVGMaskMaskUnits_userSpaceOnUse SVGMaskMaskUnitsChoice = "userSpaceOnUse"
	// The coordinate system for the various length values within the filter.
	SVGMaskMaskUnits_objectBoundingBox SVGMaskMaskUnitsChoice = "objectBoundingBox"
)

// Remove the attribute maskUnits from the element.
func (e *SVGMASKElement) MASKUNITSRemove(c SVGMaskMaskUnitsChoice) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("maskUnits")
	return e
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGMASKElement) X(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("x", s)
	return e
}

// Remove the attribute x from the element.
func (e *SVGMASKElement) XRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("x")
	return e
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGMASKElement) Y(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("y", s)
	return e
}

// Remove the attribute y from the element.
func (e *SVGMASKElement) YRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("y")
	return e
}

// The width of the rectangular region.
func (e *SVGMASKElement) WIDTH(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("width", s)
	return e
}

// Remove the attribute width from the element.
func (e *SVGMASKElement) WIDTHRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("width")
	return e
}

// The height of the rectangular region.
func (e *SVGMASKElement) HEIGHT(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("height", s)
	return e
}

// Remove the attribute height from the element.
func (e *SVGMASKElement) HEIGHTRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("height")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGMASKElement) CLASS(s ...string) *SVGMASKElement {
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
func (e *SVGMASKElement) CLASSRemove(s ...string) *SVGMASKElement {
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
func (e *SVGMASKElement) ID(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGMASKElement) IDRemove(s string) *SVGMASKElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGMASKElement) STYLE(k string, v string) *SVGMASKElement {
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
func (e *SVGMASKElement) STYLEMap(m map[string]string) *SVGMASKElement {
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
func (e *SVGMASKElement) STYLEPairs(pairs ...string) *SVGMASKElement {
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
func (e *SVGMASKElement) STYLERemove(keys ...string) *SVGMASKElement {
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