// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feFuncB is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <feFuncB> SVG filter primitive defines the transfer function for the blue
// component of the input graphic of its parent <feComponentTransfer> element.
type SVGFEFUNCBElement struct {
	*Element
}

// Create a new SVGFEFUNCBElement element.
// This will create a new element with the tag
// "feFuncB" during rendering.
func SVG_FEFUNCB(children ...ElementRenderer) *SVGFEFUNCBElement {
	e := NewElement("feFuncB", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEFUNCBElement{Element: e}
}

func (e *SVGFEFUNCBElement) Children(children ...ElementRenderer) *SVGFEFUNCBElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEFUNCBElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEFUNCBElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEFUNCBElement) Text(text string) *SVGFEFUNCBElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEFUNCBElement) TextF(format string, args ...any) *SVGFEFUNCBElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfText(condition bool, text string) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEFUNCBElement) IfTextF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEFUNCBElement) Escaped(text string) *SVGFEFUNCBElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEFUNCBElement) IfEscaped(condition bool, text string) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEFUNCBElement) EscapedF(format string, args ...any) *SVGFEFUNCBElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEFUNCBElement) CustomData(key, value string) *SVGFEFUNCBElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEFUNCBElement) IfCustomData(condition bool, key, value string) *SVGFEFUNCBElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEFUNCBElement) CustomDataF(key, format string, args ...any) *SVGFEFUNCBElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEFUNCBElement) CustomDataRemove(key string) *SVGFEFUNCBElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The type of transfer function.
func (e *SVGFEFUNCBElement) TYPE(c SVGFeFuncBTypeChoice) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeFuncBTypeChoice string

const (
	// The type of transfer function.
	SVGFeFuncBType_identity SVGFeFuncBTypeChoice = "identity"
	// The type of transfer function.
	SVGFeFuncBType_table SVGFeFuncBTypeChoice = "table"
	// The type of transfer function.
	SVGFeFuncBType_discrete SVGFeFuncBTypeChoice = "discrete"
	// The type of transfer function.
	SVGFeFuncBType_linear SVGFeFuncBTypeChoice = "linear"
	// The type of transfer function.
	SVGFeFuncBType_gamma SVGFeFuncBTypeChoice = "gamma"
)

// Remove the attribute type from the element.
func (e *SVGFEFUNCBElement) TYPERemove(c SVGFeFuncBTypeChoice) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

// Contains the list of <number>s that define the lookup table
// Values must be in the 0-1 range and be equally spaced
// There must be at least two values.
func (e *SVGFEFUNCBElement) TABLEVALUES(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("tableValues", s)
	return e
}

func (e *SVGFEFUNCBElement) IfTABLEVALUES(condition bool, s string) *SVGFEFUNCBElement {
	if condition {
		e.TABLEVALUES(s)
	}
	return e
}

// Remove the attribute tableValues from the element.
func (e *SVGFEFUNCBElement) TABLEVALUESRemove(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("tableValues")
	return e
}

// The slope attribute indicates the slope of the linear function.
func (e *SVGFEFUNCBElement) SLOPE(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("slope", f)
	return e
}

func (e *SVGFEFUNCBElement) IfSLOPE(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.SLOPE(f)
	}
	return e
}

// The intercept attribute indicates the intercept of the linear function.
func (e *SVGFEFUNCBElement) INTERCEPT(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("intercept", f)
	return e
}

func (e *SVGFEFUNCBElement) IfINTERCEPT(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.INTERCEPT(f)
	}
	return e
}

// The amplitude attribute indicates the amplitude of the cubic function.
func (e *SVGFEFUNCBElement) AMPLITUDE(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("amplitude", f)
	return e
}

func (e *SVGFEFUNCBElement) IfAMPLITUDE(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.AMPLITUDE(f)
	}
	return e
}

// The exponent attribute indicates the exponent of the exponential function.
func (e *SVGFEFUNCBElement) EXPONENT(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("exponent", f)
	return e
}

func (e *SVGFEFUNCBElement) IfEXPONENT(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.EXPONENT(f)
	}
	return e
}

// The offset attribute indicates the offset of the function.
func (e *SVGFEFUNCBElement) OFFSET(f float64) *SVGFEFUNCBElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("offset", f)
	return e
}

func (e *SVGFEFUNCBElement) IfOFFSET(condition bool, f float64) *SVGFEFUNCBElement {
	if condition {
		e.OFFSET(f)
	}
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEFUNCBElement) CLASS(s ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) IfCLASS(condition bool, s ...string) *SVGFEFUNCBElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute class from the element.
func (e *SVGFEFUNCBElement) CLASSRemove(s ...string) *SVGFEFUNCBElement {
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
func (e *SVGFEFUNCBElement) ID(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEFUNCBElement) IfID(condition bool, s string) *SVGFEFUNCBElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute id from the element.
func (e *SVGFEFUNCBElement) IDRemove(s string) *SVGFEFUNCBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFEFUNCBElement) STYLEF(k string, format string, args ...any) *SVGFEFUNCBElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCBElement) IfSTYLE(condition bool, k string, v string) *SVGFEFUNCBElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEFUNCBElement) STYLE(k string, v string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEFUNCBElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEFUNCBElement) STYLEMap(m map[string]string) *SVGFEFUNCBElement {
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
func (e *SVGFEFUNCBElement) STYLEPairs(pairs ...string) *SVGFEFUNCBElement {
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

func (e *SVGFEFUNCBElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEFUNCBElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute style from the element.
func (e *SVGFEFUNCBElement) STYLERemove(keys ...string) *SVGFEFUNCBElement {
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
