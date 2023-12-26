// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg svg is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <svg> element is a container that defines a new coordinate system and
// viewport
// It is used as the outermost element of SVG documents, but it can also be used
// to embed a SVG fragment inside an SVG or HTML document.
type SVGSVGElement struct {
	*Element
}

// Create a new SVGSVGElement element.
// This will create a new element with the tag
// "svg" during rendering.
func SVG_SVG(children ...ElementRenderer) *SVGSVGElement {
	e := NewElement("svg", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSVGElement{Element: e}
}

func (e *SVGSVGElement) Children(children ...ElementRenderer) *SVGSVGElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSVGElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSVGElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSVGElement) Text(text string) *SVGSVGElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSVGElement) TextF(format string, args ...any) *SVGSVGElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfText(condition bool, text string) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSVGElement) IfTextF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSVGElement) Escaped(text string) *SVGSVGElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSVGElement) IfEscaped(condition bool, text string) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSVGElement) EscapedF(format string, args ...any) *SVGSVGElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfEscapedF(condition bool, format string, args ...any) *SVGSVGElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSVGElement) CustomData(key, value string) *SVGSVGElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSVGElement) IfCustomData(condition bool, key, value string) *SVGSVGElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSVGElement) CustomDataF(key, format string, args ...any) *SVGSVGElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSVGElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSVGElement) CustomDataRemove(key string) *SVGSVGElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The x-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGSVGElement) X(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("x", s)
	return e
}

func (e *SVGSVGElement) IfX(condition bool, s string) *SVGSVGElement {
	if condition {
		e.X(s)
	}
	return e
}

// Remove the attribute x from the element.
func (e *SVGSVGElement) XRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("x")
	return e
}

// The y-axis coordinate of the side of the rectangular region which is closest to
// the user.
func (e *SVGSVGElement) Y(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("y", s)
	return e
}

func (e *SVGSVGElement) IfY(condition bool, s string) *SVGSVGElement {
	if condition {
		e.Y(s)
	}
	return e
}

// Remove the attribute y from the element.
func (e *SVGSVGElement) YRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("y")
	return e
}

// The width of the rectangular region.
func (e *SVGSVGElement) WIDTH(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("width", s)
	return e
}

func (e *SVGSVGElement) IfWIDTH(condition bool, s string) *SVGSVGElement {
	if condition {
		e.WIDTH(s)
	}
	return e
}

// Remove the attribute width from the element.
func (e *SVGSVGElement) WIDTHRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("width")
	return e
}

// The height of the rectangular region.
func (e *SVGSVGElement) HEIGHT(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("height", s)
	return e
}

func (e *SVGSVGElement) IfHEIGHT(condition bool, s string) *SVGSVGElement {
	if condition {
		e.HEIGHT(s)
	}
	return e
}

// Remove the attribute height from the element.
func (e *SVGSVGElement) HEIGHTRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("height")
	return e
}

// The position and size of the viewport (the viewBox) is defined by the viewBox
// attribute.
func (e *SVGSVGElement) VIEWBOX(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("viewBox", s)
	return e
}

func (e *SVGSVGElement) IfVIEWBOX(condition bool, s string) *SVGSVGElement {
	if condition {
		e.VIEWBOX(s)
	}
	return e
}

// Remove the attribute viewBox from the element.
func (e *SVGSVGElement) VIEWBOXRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("viewBox")
	return e
}

// Indicates how the viewport is fitted to the rectangle of the given width and
// height.
func (e *SVGSVGElement) PRESERVEASPECTRATIO(c SVGSvgPreserveAspectRatioChoice) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("preserveAspectRatio", string(c))
	return e
}

type SVGSvgPreserveAspectRatioChoice string

const (
	// Do not force uniform scaling.
	SVGSvgPreserveAspectRatio_none SVGSvgPreserveAspectRatioChoice = "none"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGSvgPreserveAspectRatio_xMinYMin SVGSvgPreserveAspectRatioChoice = "xMinYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSvgPreserveAspectRatio_xMidYMin SVGSvgPreserveAspectRatioChoice = "xMidYMin"
	// Align the image with the corresponding side of the viewPort.
	SVGSvgPreserveAspectRatio_xMaxYMin SVGSvgPreserveAspectRatioChoice = "xMaxYMin"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSvgPreserveAspectRatio_xMinYMid SVGSvgPreserveAspectRatioChoice = "xMinYMid"
	// Scale the image to the smallest size such that it can completely fit inside the
	// corresponding dimension of the viewPort.
	SVGSvgPreserveAspectRatio_xMidYMid SVGSvgPreserveAspectRatioChoice = "xMidYMid"
	// Align the image with the corresponding side of the viewPort.
	SVGSvgPreserveAspectRatio_xMaxYMid SVGSvgPreserveAspectRatioChoice = "xMaxYMid"
	// Align the image along the middle of the corresponding dimension of the
	// viewPort.
	SVGSvgPreserveAspectRatio_xMinYMax SVGSvgPreserveAspectRatioChoice = "xMinYMax"
	// Align the image with the corresponding side of the viewPort.
	SVGSvgPreserveAspectRatio_xMidYMax SVGSvgPreserveAspectRatioChoice = "xMidYMax"
	// Scale the image to the smallest size such that both its width and its height
	// can completely fit inside the corresponding dimension of the viewPort.
	SVGSvgPreserveAspectRatio_xMaxYMax SVGSvgPreserveAspectRatioChoice = "xMaxYMax"
)

// Remove the attribute preserveAspectRatio from the element.
func (e *SVGSVGElement) PRESERVEASPECTRATIORemove(c SVGSvgPreserveAspectRatioChoice) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("preserveAspectRatio")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSVGElement) CLASS(s ...string) *SVGSVGElement {
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

func (e *SVGSVGElement) IfCLASS(condition bool, s ...string) *SVGSVGElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute class from the element.
func (e *SVGSVGElement) CLASSRemove(s ...string) *SVGSVGElement {
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
func (e *SVGSVGElement) ID(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSVGElement) IfID(condition bool, s string) *SVGSVGElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute id from the element.
func (e *SVGSVGElement) IDRemove(s string) *SVGSVGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGSVGElement) STYLEF(k string, format string, args ...any) *SVGSVGElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSVGElement) IfSTYLE(condition bool, k string, v string) *SVGSVGElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSVGElement) STYLE(k string, v string) *SVGSVGElement {
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

func (e *SVGSVGElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSVGElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSVGElement) STYLEMap(m map[string]string) *SVGSVGElement {
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
func (e *SVGSVGElement) STYLEPairs(pairs ...string) *SVGSVGElement {
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

func (e *SVGSVGElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSVGElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute style from the element.
func (e *SVGSVGElement) STYLERemove(keys ...string) *SVGSVGElement {
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
