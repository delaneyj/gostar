// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml mo is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// This element is used to display a single operator.
type MathMLMOElement struct {
	*Element
}

// Create a new MathMLMOElement element.
// This will create a new element with the tag
// "mo" during rendering.
func MathML_MO(children ...ElementRenderer) *MathMLMOElement {
	e := NewElement("mo", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMOElement{Element: e}
}

func (e *MathMLMOElement) Children(children ...ElementRenderer) *MathMLMOElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMOElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMOElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMOElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMOElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMOElement) Text(text string) *MathMLMOElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMOElement) TextF(format string, args ...any) *MathMLMOElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMOElement) Escaped(text string) *MathMLMOElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMOElement) EscapedF(format string, args ...any) *MathMLMOElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMOElement) CustomData(key, value string) *MathMLMOElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMOElement) CustomDataRemove(key string) *MathMLMOElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// This attribute specifies whether the operator is to be rendered as a fence
// Possible values are true and false.
func (e *MathMLMOElement) FENCE(c MathMLMoFenceChoice) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("fence", string(c))
	return e
}

type MathMLMoFenceChoice string

const (
	MathMLMoFence_true MathMLMoFenceChoice = "true"

	MathMLMoFence_false MathMLMoFenceChoice = "false"
)

// Remove the attribute fence from the element.
func (e *MathMLMOElement) FENCERemove(c MathMLMoFenceChoice) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("fence")
	return e
}

// This attribute specifies the minimum amount of space that should be left on the
// left side of the operator
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMOElement) LSPACE(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("lspace", s)
	return e
}

// Remove the attribute lspace from the element.
func (e *MathMLMOElement) LSPACERemove(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("lspace")
	return e
}

// This attribute specifies the minimum amount of space that should be left on the
// right side of the operator
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMOElement) RSPACE(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("rspace", s)
	return e
}

// Remove the attribute rspace from the element.
func (e *MathMLMOElement) RSPACERemove(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("rspace")
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMOElement) CLASS(s ...string) *MathMLMOElement {
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
func (e *MathMLMOElement) CLASSRemove(s ...string) *MathMLMOElement {
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

// This attribute specifies the text directionality of the element, merely
// indicating what direction the text flows when surrounded by text with inherent
// directionality (such as Arabic or Hebrew)
// Possible values are ltr (left-to-right) and rtl (right-to-left).
func (e *MathMLMOElement) DIR(c MathMLMoDirChoice) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMoDirChoice string

const (
	// left-to-right
	MathMLMoDir_ltr MathMLMoDirChoice = "ltr"
	// right-to-left
	MathMLMoDir_rtl MathMLMoDirChoice = "rtl"
)

// Remove the attribute dir from the element.
func (e *MathMLMOElement) DIRRemove(c MathMLMoDirChoice) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMOElement) DISPLAYSTYLE(c MathMLMoDisplaystyleChoice) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMoDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMoDisplaystyle_true MathMLMoDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMoDisplaystyle_false MathMLMoDisplaystyleChoice = "false"
)

// Remove the attribute displaystyle from the element.
func (e *MathMLMOElement) DISPLAYSTYLERemove(c MathMLMoDisplaystyleChoice) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMOElement) ID(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *MathMLMOElement) IDRemove(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMOElement) MATHBACKGROUND(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

// Remove the attribute mathbackground from the element.
func (e *MathMLMOElement) MATHBACKGROUNDRemove(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMOElement) MATHCOLOR(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

// Remove the attribute mathcolor from the element.
func (e *MathMLMOElement) MATHCOLORRemove(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMOElement) MATHSIZESTR(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

// Remove the attribute mathsizeStr from the element.
func (e *MathMLMOElement) MATHSIZESTRRemove(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathsize")
	return e
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMOElement) NONCE(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

// Remove the attribute nonce from the element.
func (e *MathMLMOElement) NONCERemove(s string) *MathMLMOElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMOElement) SCRIPTLEVEL(i int) *MathMLMOElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

// Remove the attribute scriptlevel from the element.
func (e *MathMLMOElement) SCRIPTLEVELRemove(i int) *MathMLMOElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMOElement) STYLE(k string, v string) *MathMLMOElement {
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
func (e *MathMLMOElement) STYLEMap(m map[string]string) *MathMLMOElement {
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
func (e *MathMLMOElement) STYLEPairs(pairs ...string) *MathMLMOElement {
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
func (e *MathMLMOElement) STYLERemove(keys ...string) *MathMLMOElement {
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

// This attribute specifies the position of the current element in the tabbing
// order for the current document
// This value must be a number between 0 and 32767
// User agents should ignore leading zeros.
func (e *MathMLMOElement) TABINDEX(i int) *MathMLMOElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

// Remove the attribute tabindex from the element.
func (e *MathMLMOElement) TABINDEXRemove(i int) *MathMLMOElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}