// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml mi is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// This element is used to display a single identifier or a single operator.
type MathMLMIElement struct {
	*Element
}

// Create a new MathMLMIElement element.
// This will create a new element with the tag
// "mi" during rendering.
func MathML_MI(children ...ElementRenderer) *MathMLMIElement {
	e := NewElement("mi", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMIElement{Element: e}
}

func (e *MathMLMIElement) Children(children ...ElementRenderer) *MathMLMIElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMIElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMIElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMIElement) Text(text string) *MathMLMIElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMIElement) TextF(format string, args ...any) *MathMLMIElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) Escaped(text string) *MathMLMIElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMIElement) EscapedF(format string, args ...any) *MathMLMIElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) CustomData(key, value string) *MathMLMIElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMIElement) CustomDataF(key, format string, args ...any) *MathMLMIElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) CustomDataRemove(key string) *MathMLMIElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// This attribute specifies the variant form of the character
// Possible values are normal, bold, italic, bold-italic, double-struck,
// bold-fraktur, script, bold-script, fraktur, sans-serif, bold-sans-serif,
// sans-serif-italic, sans-serif-bold-italic, monospace, initial, and tailed.
func (e *MathMLMIElement) MATHVARIANT(c MathMLMiMathvariantChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathvariant", string(c))
	return e
}

type MathMLMiMathvariantChoice string

const (
	MathMLMiMathvariant_normal MathMLMiMathvariantChoice = "normal"

	MathMLMiMathvariant_bold MathMLMiMathvariantChoice = "bold"

	MathMLMiMathvariant_italic MathMLMiMathvariantChoice = "italic"

	MathMLMiMathvariant_bold_italic MathMLMiMathvariantChoice = "bold-italic"

	MathMLMiMathvariant_double_struck MathMLMiMathvariantChoice = "double-struck"

	MathMLMiMathvariant_bold_fraktur MathMLMiMathvariantChoice = "bold-fraktur"

	MathMLMiMathvariant_script MathMLMiMathvariantChoice = "script"

	MathMLMiMathvariant_bold_script MathMLMiMathvariantChoice = "bold-script"

	MathMLMiMathvariant_fraktur MathMLMiMathvariantChoice = "fraktur"

	MathMLMiMathvariant_sans_serif MathMLMiMathvariantChoice = "sans-serif"

	MathMLMiMathvariant_bold_sans_serif MathMLMiMathvariantChoice = "bold-sans-serif"

	MathMLMiMathvariant_sans_serif_italic MathMLMiMathvariantChoice = "sans-serif-italic"

	MathMLMiMathvariant_sans_serif_bold_italic MathMLMiMathvariantChoice = "sans-serif-bold-italic"

	MathMLMiMathvariant_monospace MathMLMiMathvariantChoice = "monospace"

	MathMLMiMathvariant_initial MathMLMiMathvariantChoice = "initial"

	MathMLMiMathvariant_tailed MathMLMiMathvariantChoice = "tailed"
)

// Remove the attribute mathvariant from the element.
func (e *MathMLMIElement) MATHVARIANTRemove(c MathMLMiMathvariantChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathvariant")
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMIElement) CLASS(s ...string) *MathMLMIElement {
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
func (e *MathMLMIElement) CLASSRemove(s ...string) *MathMLMIElement {
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
func (e *MathMLMIElement) DIR(c MathMLMiDirChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMiDirChoice string

const (
	// left-to-right
	MathMLMiDir_ltr MathMLMiDirChoice = "ltr"
	// right-to-left
	MathMLMiDir_rtl MathMLMiDirChoice = "rtl"
)

// Remove the attribute dir from the element.
func (e *MathMLMIElement) DIRRemove(c MathMLMiDirChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMIElement) DISPLAYSTYLE(c MathMLMiDisplaystyleChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMiDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMiDisplaystyle_true MathMLMiDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMiDisplaystyle_false MathMLMiDisplaystyleChoice = "false"
)

// Remove the attribute displaystyle from the element.
func (e *MathMLMIElement) DISPLAYSTYLERemove(c MathMLMiDisplaystyleChoice) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMIElement) ID(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *MathMLMIElement) IDRemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMIElement) MATHBACKGROUND(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

// Remove the attribute mathbackground from the element.
func (e *MathMLMIElement) MATHBACKGROUNDRemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMIElement) MATHCOLOR(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

// Remove the attribute mathcolor from the element.
func (e *MathMLMIElement) MATHCOLORRemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMIElement) MATHSIZESTR(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

// Remove the attribute mathsizeStr from the element.
func (e *MathMLMIElement) MATHSIZESTRRemove(s string) *MathMLMIElement {
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
func (e *MathMLMIElement) NONCE(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

// Remove the attribute nonce from the element.
func (e *MathMLMIElement) NONCERemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMIElement) SCRIPTLEVEL(i int) *MathMLMIElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

// Remove the attribute scriptlevel from the element.
func (e *MathMLMIElement) SCRIPTLEVELRemove(i int) *MathMLMIElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMIElement) STYLEF(k string, format string, args ...any) *MathMLMIElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) STYLE(k string, v string) *MathMLMIElement {
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
func (e *MathMLMIElement) STYLEMap(m map[string]string) *MathMLMIElement {
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
func (e *MathMLMIElement) STYLEPairs(pairs ...string) *MathMLMIElement {
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
func (e *MathMLMIElement) STYLERemove(keys ...string) *MathMLMIElement {
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
func (e *MathMLMIElement) TABINDEX(i int) *MathMLMIElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

// Remove the attribute tabindex from the element.
func (e *MathMLMIElement) TABINDEXRemove(i int) *MathMLMIElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}
