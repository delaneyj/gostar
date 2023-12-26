// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feDisplacementMap is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <feDisplacementMap> SVG filter primitive uses the pixel values from the
// image from in2 to spatially displace the image from in.
type SVGFEDISPLACEMENTMAPElement struct {
	*Element
}

// Create a new SVGFEDISPLACEMENTMAPElement element.
// This will create a new element with the tag
// "feDisplacementMap" during rendering.
func SVG_FEDISPLACEMENTMAP(children ...ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
	e := NewElement("feDisplacementMap", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEDISPLACEMENTMAPElement{Element: e}
}

func (e *SVGFEDISPLACEMENTMAPElement) Children(children ...ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) Text(text string) *SVGFEDISPLACEMENTMAPElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) TextF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) Escaped(text string) *SVGFEDISPLACEMENTMAPElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) EscapedF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) CustomData(key, value string) *SVGFEDISPLACEMENTMAPElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) CustomDataRemove(key string) *SVGFEDISPLACEMENTMAPElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The input for this filter.
func (e *SVGFEDISPLACEMENTMAPElement) IN(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

// Remove the attribute in from the element.
func (e *SVGFEDISPLACEMENTMAPElement) INRemove(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in")
	return e
}

// The displacement map
// This attribute can take on the same values as the 'in' attribute.
func (e *SVGFEDISPLACEMENTMAPElement) IN2(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in2", s)
	return e
}

// Remove the attribute in2 from the element.
func (e *SVGFEDISPLACEMENTMAPElement) IN2Remove(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in2")
	return e
}

// The scale attribute defines the maximum value for the in2 displacement
// A value of 0 disables the effect of the displacement map.
func (e *SVGFEDISPLACEMENTMAPElement) SCALE(f float64) *SVGFEDISPLACEMENTMAPElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("scale", f)
	return e
}

// The xChannelSelector attribute indicates which color channel from in2 to use to
// displace the pixels in in the horizontal direction.
func (e *SVGFEDISPLACEMENTMAPElement) XCHANNELSELECTOR(c SVGFeDisplacementMapXChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("xChannelSelector", string(c))
	return e
}

type SVGFeDisplacementMapXChannelSelectorChoice string

const (
	// The red channel of in2 is used to displace the x coordinate of each pixel.
	SVGFeDisplacementMapXChannelSelector_R SVGFeDisplacementMapXChannelSelectorChoice = "R"
	// The green channel of in2 is used to displace the x coordinate of each pixel.
	SVGFeDisplacementMapXChannelSelector_G SVGFeDisplacementMapXChannelSelectorChoice = "G"
	// The blue channel of in2 is used to displace the x coordinate of each pixel.
	SVGFeDisplacementMapXChannelSelector_B SVGFeDisplacementMapXChannelSelectorChoice = "B"
	// The alpha channel of in2 is used to displace the x coordinate of each pixel.
	SVGFeDisplacementMapXChannelSelector_A SVGFeDisplacementMapXChannelSelectorChoice = "A"
)

// Remove the attribute xChannelSelector from the element.
func (e *SVGFEDISPLACEMENTMAPElement) XCHANNELSELECTORRemove(c SVGFeDisplacementMapXChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("xChannelSelector")
	return e
}

// The yChannelSelector attribute indicates which color channel from in2 to use to
// displace the pixels in in the vertical direction.
func (e *SVGFEDISPLACEMENTMAPElement) YCHANNELSELECTOR(c SVGFeDisplacementMapYChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("yChannelSelector", string(c))
	return e
}

type SVGFeDisplacementMapYChannelSelectorChoice string

const (
	// The red channel of in2 is used to displace the y coordinate of each pixel.
	SVGFeDisplacementMapYChannelSelector_R SVGFeDisplacementMapYChannelSelectorChoice = "R"
	// The green channel of in2 is used to displace the y coordinate of each pixel.
	SVGFeDisplacementMapYChannelSelector_G SVGFeDisplacementMapYChannelSelectorChoice = "G"
	// The blue channel of in2 is used to displace the y coordinate of each pixel.
	SVGFeDisplacementMapYChannelSelector_B SVGFeDisplacementMapYChannelSelectorChoice = "B"
	// The alpha channel of in2 is used to displace the y coordinate of each pixel.
	SVGFeDisplacementMapYChannelSelector_A SVGFeDisplacementMapYChannelSelectorChoice = "A"
)

// Remove the attribute yChannelSelector from the element.
func (e *SVGFEDISPLACEMENTMAPElement) YCHANNELSELECTORRemove(c SVGFeDisplacementMapYChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("yChannelSelector")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEDISPLACEMENTMAPElement) CLASS(s ...string) *SVGFEDISPLACEMENTMAPElement {
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
func (e *SVGFEDISPLACEMENTMAPElement) CLASSRemove(s ...string) *SVGFEDISPLACEMENTMAPElement {
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
func (e *SVGFEDISPLACEMENTMAPElement) ID(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGFEDISPLACEMENTMAPElement) IDRemove(s string) *SVGFEDISPLACEMENTMAPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFEDISPLACEMENTMAPElement) STYLE(k string, v string) *SVGFEDISPLACEMENTMAPElement {
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
func (e *SVGFEDISPLACEMENTMAPElement) STYLEMap(m map[string]string) *SVGFEDISPLACEMENTMAPElement {
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
func (e *SVGFEDISPLACEMENTMAPElement) STYLEPairs(pairs ...string) *SVGFEDISPLACEMENTMAPElement {
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
func (e *SVGFEDISPLACEMENTMAPElement) STYLERemove(keys ...string) *SVGFEDISPLACEMENTMAPElement {
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