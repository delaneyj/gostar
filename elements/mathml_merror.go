// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml merror is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// This element is used to indicate that an error has occurred while processing a
// MathML expression
// It can be used to display an error message, or to highlight the error in the
// expression.
type MathMLMERRORElement struct {
	*Element
}

// Create a new MathMLMERRORElement element.
// This will create a new element with the tag
// "merror" during rendering.
func MathML_MERROR(children ...ElementRenderer) *MathMLMERRORElement {
	e := NewElement("merror", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMERRORElement{Element: e}
}

func (e *MathMLMERRORElement) Children(children ...ElementRenderer) *MathMLMERRORElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMERRORElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMERRORElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMERRORElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMERRORElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMERRORElement) Text(text string) *MathMLMERRORElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMERRORElement) TextF(format string, args ...any) *MathMLMERRORElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMERRORElement) Escaped(text string) *MathMLMERRORElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMERRORElement) EscapedF(format string, args ...any) *MathMLMERRORElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMERRORElement) CustomData(key, value string) *MathMLMERRORElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMERRORElement) CustomDataF(key, format string, args ...any) *MathMLMERRORElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMERRORElement) CustomDataRemove(key string) *MathMLMERRORElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMERRORElement) CLASS(s ...string) *MathMLMERRORElement {
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
func (e *MathMLMERRORElement) CLASSRemove(s ...string) *MathMLMERRORElement {
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
func (e *MathMLMERRORElement) DIR(c MathMLMerrorDirChoice) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMerrorDirChoice string

const (
	// left-to-right
	MathMLMerrorDir_ltr MathMLMerrorDirChoice = "ltr"
	// right-to-left
	MathMLMerrorDir_rtl MathMLMerrorDirChoice = "rtl"
)

// Remove the attribute dir from the element.
func (e *MathMLMERRORElement) DIRRemove(c MathMLMerrorDirChoice) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMERRORElement) DISPLAYSTYLE(c MathMLMerrorDisplaystyleChoice) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMerrorDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMerrorDisplaystyle_true MathMLMerrorDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMerrorDisplaystyle_false MathMLMerrorDisplaystyleChoice = "false"
)

// Remove the attribute displaystyle from the element.
func (e *MathMLMERRORElement) DISPLAYSTYLERemove(c MathMLMerrorDisplaystyleChoice) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMERRORElement) ID(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *MathMLMERRORElement) IDRemove(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMERRORElement) MATHBACKGROUND(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

// Remove the attribute mathbackground from the element.
func (e *MathMLMERRORElement) MATHBACKGROUNDRemove(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMERRORElement) MATHCOLOR(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

// Remove the attribute mathcolor from the element.
func (e *MathMLMERRORElement) MATHCOLORRemove(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMERRORElement) MATHSIZESTR(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

// Remove the attribute mathsizeStr from the element.
func (e *MathMLMERRORElement) MATHSIZESTRRemove(s string) *MathMLMERRORElement {
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
func (e *MathMLMERRORElement) NONCE(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

// Remove the attribute nonce from the element.
func (e *MathMLMERRORElement) NONCERemove(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMERRORElement) SCRIPTLEVEL(i int) *MathMLMERRORElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

// Remove the attribute scriptlevel from the element.
func (e *MathMLMERRORElement) SCRIPTLEVELRemove(i int) *MathMLMERRORElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMERRORElement) STYLEF(k string, format string, args ...any) *MathMLMERRORElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMERRORElement) STYLE(k string, v string) *MathMLMERRORElement {
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
func (e *MathMLMERRORElement) STYLEMap(m map[string]string) *MathMLMERRORElement {
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
func (e *MathMLMERRORElement) STYLEPairs(pairs ...string) *MathMLMERRORElement {
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
func (e *MathMLMERRORElement) STYLERemove(keys ...string) *MathMLMERRORElement {
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
func (e *MathMLMERRORElement) TABINDEX(i int) *MathMLMERRORElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

// Remove the attribute tabindex from the element.
func (e *MathMLMERRORElement) TABINDEXRemove(i int) *MathMLMERRORElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}
