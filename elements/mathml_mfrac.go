// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml mfrac is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"html"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// This element is used to display a fraction.
type MathMLMFRACElement struct {
	*Element
}

// Create a new MathMLMFRACElement element.
// This will create a new element with the tag
// "mfrac" during rendering.
func MathML_MFRAC(children ...ElementRenderer) *MathMLMFRACElement {
	e := NewElement("mfrac", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMFRACElement{Element: e}
}

func (e *MathMLMFRACElement) Children(children ...ElementRenderer) *MathMLMFRACElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMFRACElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMFRACElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMFRACElement) Attr(name string, value string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMFRACElement) Attrs(attrs ...string) *MathMLMFRACElement {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMFRACElement) AttrsMap(attrs map[string]string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMFRACElement) Text(text string) *MathMLMFRACElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMFRACElement) TextF(format string, args ...any) *MathMLMFRACElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfText(condition bool, text string) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMFRACElement) IfTextF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMFRACElement) Escaped(text string) *MathMLMFRACElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMFRACElement) IfEscaped(condition bool, text string) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMFRACElement) EscapedF(format string, args ...any) *MathMLMFRACElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMFRACElement) CustomData(key, value string) *MathMLMFRACElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMFRACElement) IfCustomData(condition bool, key, value string) *MathMLMFRACElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMFRACElement) CustomDataF(key, format string, args ...any) *MathMLMFRACElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMFRACElement) CustomDataRemove(key string) *MathMLMFRACElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// This attribute specifies whether the fraction line is to be drawn straight or
// to beveled
// Possible values are true and false.
func (e *MathMLMFRACElement) BEVELLED(c MathMLMfracBevelledChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("bevelled", string(c))
	return e
}

type MathMLMfracBevelledChoice string

const (
	MathMLMfracBevelled_true MathMLMfracBevelledChoice = "true"

	MathMLMfracBevelled_false MathMLMfracBevelledChoice = "false"
)

// Remove the attribute BEVELLED from the element.
func (e *MathMLMFRACElement) BEVELLEDRemove(c MathMLMfracBevelledChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("bevelled")
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLMFRACElement) CLASS(s ...string) *MathMLMFRACElement {
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

func (e *MathMLMFRACElement) IfCLASS(condition bool, s ...string) *MathMLMFRACElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMFRACElement) CLASSRemove(s ...string) *MathMLMFRACElement {
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
func (e *MathMLMFRACElement) DIR(c MathMLMfracDirChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMfracDirChoice string

const (
	// left-to-right
	MathMLMfracDir_ltr MathMLMfracDirChoice = "ltr"
	// right-to-left
	MathMLMfracDir_rtl MathMLMfracDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMFRACElement) DIRRemove(c MathMLMfracDirChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMFRACElement) DISPLAYSTYLE(c MathMLMfracDisplaystyleChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMfracDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMfracDisplaystyle_true MathMLMfracDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMfracDisplaystyle_false MathMLMfracDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMFRACElement) DISPLAYSTYLERemove(c MathMLMfracDisplaystyleChoice) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMFRACElement) ID(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMFRACElement) IDF(format string, args ...any) *MathMLMFRACElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfID(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfIDF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMFRACElement) IDRemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *MathMLMFRACElement) IDRemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMFRACElement) MATHBACKGROUND(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMFRACElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMFRACElement) MATHBACKGROUNDRemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

func (e *MathMLMFRACElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMFRACElement) MATHCOLOR(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMFRACElement) MATHCOLORF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfMATHCOLOR(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMFRACElement) MATHCOLORRemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

func (e *MathMLMFRACElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMFRACElement) MATHSIZE_STR(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMFRACElement) MATHSIZE_STRF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMFRACElement) MATHSIZE_STRRemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathsize")
	return e
}

func (e *MathMLMFRACElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMFRACElement) NONCE(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMFRACElement) NONCEF(format string, args ...any) *MathMLMFRACElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfNONCE(condition bool, s string) *MathMLMFRACElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMFRACElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMFRACElement) NONCERemove(s string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

func (e *MathMLMFRACElement) NONCERemoveF(format string, args ...any) *MathMLMFRACElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMFRACElement) SCRIPTLEVEL(i int) *MathMLMFRACElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMFRACElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMFRACElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMFRACElement) SCRIPTLEVELRemove(i int) *MathMLMFRACElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMFRACElement) STYLEF(k string, format string, args ...any) *MathMLMFRACElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMFRACElement) IfSTYLE(condition bool, k string, v string) *MathMLMFRACElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMFRACElement) STYLE(k string, v string) *MathMLMFRACElement {
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

func (e *MathMLMFRACElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMFRACElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMFRACElement) STYLEMap(m map[string]string) *MathMLMFRACElement {
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
func (e *MathMLMFRACElement) STYLEPairs(pairs ...string) *MathMLMFRACElement {
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

func (e *MathMLMFRACElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMFRACElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMFRACElement) STYLERemove(keys ...string) *MathMLMFRACElement {
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
func (e *MathMLMFRACElement) TABINDEX(i int) *MathMLMFRACElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMFRACElement) IfTABINDEX(condition bool, i int) *MathMLMFRACElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMFRACElement) TABINDEXRemove(i int) *MathMLMFRACElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMFRACElement) DATASTAR_MERGE_STORE(v any) *MathMLMFRACElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("merge-store", html.EscapeString(string(b)))
	return e
}

// Sets the reference of the element

func (e *MathMLMFRACElement) DATASTAR_REF(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMFRACElement) DATASTAR_REFRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMFRACElement) DATASTAR_BIND(key string, expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMFRACElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMFRACElement) DATASTAR_BINDRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMFRACElement) DATASTAR_MODEL(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMFRACElement) DATASTAR_MODELRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMFRACElement) DATASTAR_TEXT(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMFRACElement) DATASTAR_TEXTRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMfracOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMfracOnModDebounce(
	d time.Duration,
) MathMLMfracOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMfracOnModThrottle(
	d time.Duration,
) MathMLMfracOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMFRACElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMfracOnMod) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMfracOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMfracOnMod) *MathMLMFRACElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMFRACElement) DATASTAR_ONRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMFRACElement) DATASTAR_FOCUSSet(b bool) *MathMLMFRACElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMFRACElement) DATASTAR_FOCUS() *MathMLMFRACElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMFRACElement) DATASTAR_HEADER(key string, expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMFRACElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMFRACElement) DATASTAR_HEADERRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMFRACElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMFRACElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMFRACElement) DATASTAR_SHOWSet(b bool) *MathMLMFRACElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMFRACElement) DATASTAR_SHOW() *MathMLMFRACElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMFRACElement) DATASTAR_INTERSECTS(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *MathMLMFRACElement) DATASTAR_INTERSECTSRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *MathMLMFRACElement) DATASTAR_TELEPORTSet(b bool) *MathMLMFRACElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMFRACElement) DATASTAR_TELEPORT() *MathMLMFRACElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMFRACElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMFRACElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMFRACElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMFRACElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMFRACElement) DATASTAR_VIEW_TRANSITION(expression string) *MathMLMFRACElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMFRACElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *MathMLMFRACElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMFRACElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMFRACElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
