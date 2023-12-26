// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml annotation-xml is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// This element is used to include comments or annotations within a MathML
// expression
// It can be used to provide additional information about the expression, or to
// include comments for the author of the expression.
type MathMLANNOTATION_XMLElement struct {
	*Element
}

// Create a new MathMLANNOTATION_XMLElement element.
// This will create a new element with the tag
// "annotation-xml" during rendering.
func MathML_ANNOTATION_XML(children ...ElementRenderer) *MathMLANNOTATION_XMLElement {
	e := NewElement("annotation-xml", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLANNOTATION_XMLElement{Element: e}
}

func (e *MathMLANNOTATION_XMLElement) Children(children ...ElementRenderer) *MathMLANNOTATION_XMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLANNOTATION_XMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLANNOTATION_XMLElement) Text(text string) *MathMLANNOTATION_XMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLANNOTATION_XMLElement) TextF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) Escaped(text string) *MathMLANNOTATION_XMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLANNOTATION_XMLElement) EscapedF(format string, args ...any) *MathMLANNOTATION_XMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) CustomData(key, value string) *MathMLANNOTATION_XMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLANNOTATION_XMLElement) CustomDataRemove(key string) *MathMLANNOTATION_XMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// This attribute specifies the encoding used for the text content of the element
// Possible values are text/plain, text/html, and application/x-tex.
func (e *MathMLANNOTATION_XMLElement) ENCODING(c MathMLAnnotationXmlEncodingChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("encoding", string(c))
	return e
}

type MathMLAnnotationXmlEncodingChoice string

const (
	MathMLAnnotationXmlEncoding_text_plain MathMLAnnotationXmlEncodingChoice = "text/plain"

	MathMLAnnotationXmlEncoding_text_html MathMLAnnotationXmlEncodingChoice = "text/html"

	MathMLAnnotationXmlEncoding_application_x_tex MathMLAnnotationXmlEncodingChoice = "application/x-tex"
)

// Remove the attribute encoding from the element.
func (e *MathMLANNOTATION_XMLElement) ENCODINGRemove(c MathMLAnnotationXmlEncodingChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("encoding")
	return e
}

// This attribute specifies the name of the annotation.
func (e *MathMLANNOTATION_XMLElement) NAME(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("name", s)
	return e
}

// Remove the attribute name from the element.
func (e *MathMLANNOTATION_XMLElement) NAMERemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("name")
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLANNOTATION_XMLElement) CLASS(s ...string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) CLASSRemove(s ...string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) DIR(c MathMLAnnotationXmlDirChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLAnnotationXmlDirChoice string

const (
	// left-to-right
	MathMLAnnotationXmlDir_ltr MathMLAnnotationXmlDirChoice = "ltr"
	// right-to-left
	MathMLAnnotationXmlDir_rtl MathMLAnnotationXmlDirChoice = "rtl"
)

// Remove the attribute dir from the element.
func (e *MathMLANNOTATION_XMLElement) DIRRemove(c MathMLAnnotationXmlDirChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLANNOTATION_XMLElement) DISPLAYSTYLE(c MathMLAnnotationXmlDisplaystyleChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLAnnotationXmlDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLAnnotationXmlDisplaystyle_true MathMLAnnotationXmlDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLAnnotationXmlDisplaystyle_false MathMLAnnotationXmlDisplaystyleChoice = "false"
)

// Remove the attribute displaystyle from the element.
func (e *MathMLANNOTATION_XMLElement) DISPLAYSTYLERemove(c MathMLAnnotationXmlDisplaystyleChoice) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLANNOTATION_XMLElement) ID(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *MathMLANNOTATION_XMLElement) IDRemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLANNOTATION_XMLElement) MATHBACKGROUND(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

// Remove the attribute mathbackground from the element.
func (e *MathMLANNOTATION_XMLElement) MATHBACKGROUNDRemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLANNOTATION_XMLElement) MATHCOLOR(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

// Remove the attribute mathcolor from the element.
func (e *MathMLANNOTATION_XMLElement) MATHCOLORRemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLANNOTATION_XMLElement) MATHSIZESTR(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

// Remove the attribute mathsizeStr from the element.
func (e *MathMLANNOTATION_XMLElement) MATHSIZESTRRemove(s string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) NONCE(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

// Remove the attribute nonce from the element.
func (e *MathMLANNOTATION_XMLElement) NONCERemove(s string) *MathMLANNOTATION_XMLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLANNOTATION_XMLElement) SCRIPTLEVEL(i int) *MathMLANNOTATION_XMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

// Remove the attribute scriptlevel from the element.
func (e *MathMLANNOTATION_XMLElement) SCRIPTLEVELRemove(i int) *MathMLANNOTATION_XMLElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLANNOTATION_XMLElement) STYLE(k string, v string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) STYLEMap(m map[string]string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) STYLEPairs(pairs ...string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) STYLERemove(keys ...string) *MathMLANNOTATION_XMLElement {
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
func (e *MathMLANNOTATION_XMLElement) TABINDEX(i int) *MathMLANNOTATION_XMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

// Remove the attribute tabindex from the element.
func (e *MathMLANNOTATION_XMLElement) TABINDEXRemove(i int) *MathMLANNOTATION_XMLElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}