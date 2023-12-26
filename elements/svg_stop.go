// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg stop is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <stop> SVG element defines the ramp of colors to use on a gradient, which
// is a child element to either the <linearGradient> or the <radialGradient>
// element.
type SVGSTOPElement struct {
	*Element
}

// Create a new SVGSTOPElement element.
// This will create a new element with the tag
// "stop" during rendering.
func SVG_STOP(children ...ElementRenderer) *SVGSTOPElement {
	e := NewElement("stop", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSTOPElement{Element: e}
}

func (e *SVGSTOPElement) Children(children ...ElementRenderer) *SVGSTOPElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSTOPElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSTOPElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSTOPElement) Text(text string) *SVGSTOPElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSTOPElement) TextF(format string, args ...any) *SVGSTOPElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSTOPElement) IfText(condition bool, text string) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSTOPElement) IfTextF(condition bool, format string, args ...any) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSTOPElement) Escaped(text string) *SVGSTOPElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSTOPElement) IfEscaped(condition bool, text string) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSTOPElement) EscapedF(format string, args ...any) *SVGSTOPElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSTOPElement) IfEscapedF(condition bool, format string, args ...any) *SVGSTOPElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSTOPElement) CustomData(key, value string) *SVGSTOPElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSTOPElement) IfCustomData(condition bool, key, value string) *SVGSTOPElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSTOPElement) CustomDataF(key, format string, args ...any) *SVGSTOPElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSTOPElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSTOPElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSTOPElement) CustomDataRemove(key string) *SVGSTOPElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The offset from the start of the gradient where the color first takes effect.
func (e *SVGSTOPElement) OFFSET(f float64) *SVGSTOPElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("offset", f)
	return e
}

func (e *SVGSTOPElement) IfOFFSET(condition bool, f float64) *SVGSTOPElement {
	if condition {
		e.OFFSET(f)
	}
	return e
}

// The color of the gradient stop.
func (e *SVGSTOPElement) STOP_COLOR(s string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("stop-color", s)
	return e
}

func (e *SVGSTOPElement) IfSTOP_COLOR(condition bool, s string) *SVGSTOPElement {
	if condition {
		e.STOP_COLOR(s)
	}
	return e
}

// Remove the attribute stop-color from the element.
func (e *SVGSTOPElement) STOP_COLORRemove(s string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("stop-color")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSTOPElement) CLASS(s ...string) *SVGSTOPElement {
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

func (e *SVGSTOPElement) IfCLASS(condition bool, s ...string) *SVGSTOPElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute class from the element.
func (e *SVGSTOPElement) CLASSRemove(s ...string) *SVGSTOPElement {
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
func (e *SVGSTOPElement) ID(s string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSTOPElement) IfID(condition bool, s string) *SVGSTOPElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute id from the element.
func (e *SVGSTOPElement) IDRemove(s string) *SVGSTOPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGSTOPElement) STYLEF(k string, format string, args ...any) *SVGSTOPElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSTOPElement) IfSTYLE(condition bool, k string, v string) *SVGSTOPElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSTOPElement) STYLE(k string, v string) *SVGSTOPElement {
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

func (e *SVGSTOPElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSTOPElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSTOPElement) STYLEMap(m map[string]string) *SVGSTOPElement {
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
func (e *SVGSTOPElement) STYLEPairs(pairs ...string) *SVGSTOPElement {
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

func (e *SVGSTOPElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSTOPElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute style from the element.
func (e *SVGSTOPElement) STYLERemove(keys ...string) *SVGSTOPElement {
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
