// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml mi is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
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

func (e *MathMLMIElement) Attr(name string, value string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMIElement) Attrs(attrs ...string) *MathMLMIElement {
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

func (e *MathMLMIElement) AttrsMap(attrs map[string]string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
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

func (e *MathMLMIElement) IfText(condition bool, text string) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMIElement) IfTextF(condition bool, format string, args ...any) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMIElement) Escaped(text string) *MathMLMIElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMIElement) IfEscaped(condition bool, text string) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMIElement) EscapedF(format string, args ...any) *MathMLMIElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMIElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMIElement) CustomData(key, value string) *MathMLMIElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMIElement) IfCustomData(condition bool, key, value string) *MathMLMIElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMIElement) CustomDataF(key, format string, args ...any) *MathMLMIElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMIElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMIElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
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

// Remove the attribute MATHVARIANT from the element.
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

func (e *MathMLMIElement) IfCLASS(condition bool, s ...string) *MathMLMIElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
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

// Remove the attribute DIR from the element.
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

// Remove the attribute DISPLAYSTYLE from the element.
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

func (e *MathMLMIElement) IfID(condition bool, s string) *MathMLMIElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
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

func (e *MathMLMIElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMIElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
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

func (e *MathMLMIElement) IfMATHCOLOR(condition bool, s string) *MathMLMIElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMIElement) MATHCOLORRemove(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMIElement) MATHSIZE_STR(s string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMIElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMIElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMIElement) MATHSIZE_STRRemove(s string) *MathMLMIElement {
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

func (e *MathMLMIElement) IfNONCE(condition bool, s string) *MathMLMIElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

// Remove the attribute NONCE from the element.
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

func (e *MathMLMIElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMIElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
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

func (e *MathMLMIElement) IfSTYLE(condition bool, k string, v string) *MathMLMIElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
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

func (e *MathMLMIElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMIElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
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

func (e *MathMLMIElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMIElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
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

func (e *MathMLMIElement) IfTABINDEX(condition bool, i int) *MathMLMIElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMIElement) TABINDEXRemove(i int) *MathMLMIElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMIElement) DATASTAR_MERGE_STORE(v any) *MathMLMIElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("data-merge-store", string(b))
	return e
}

// Sets the reference of the element

func (e *MathMLMIElement) DATASTAR_REF(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMIElement) DATASTAR_REFRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMIElement) DATASTAR_BIND(key string, expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMIElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMIElement) DATASTAR_BINDRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMIElement) DATASTAR_MODEL(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMIElement) DATASTAR_MODELRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMIElement) DATASTAR_TEXT(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMIElement) DATASTAR_TEXTRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMiDataOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMiDataOnModDebounce(
	d time.Duration,
) MathMLMiDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMiDataOnModThrottle(
	d time.Duration,
) MathMLMiDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMIElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMiDataOnMod) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMiDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMiDataOnMod) *MathMLMIElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMIElement) DATASTAR_ONRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMIElement) DATASTAR_FOCUSSet(b bool) *MathMLMIElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMIElement) DATASTAR_FOCUS() *MathMLMIElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMIElement) DATASTAR_HEADER(key string, expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMIElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMIElement) DATASTAR_HEADERRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMIElement) DATASTAR_FETCH_URL(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMIElement) DATASTAR_FETCH_URLRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMIElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMIElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMIElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMIElement) DATASTAR_SHOWSet(b bool) *MathMLMIElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMIElement) DATASTAR_SHOW() *MathMLMIElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMIElement) DATASTAR_INTERSECTSSet(b bool) *MathMLMIElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMIElement) DATASTAR_INTERSECTS() *MathMLMIElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *MathMLMIElement) DATASTAR_TELEPORTSet(b bool) *MathMLMIElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMIElement) DATASTAR_TELEPORT() *MathMLMIElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMIElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMIElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMIElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMIElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMIElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *MathMLMIElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMIElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *MathMLMIElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMIElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMIElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
