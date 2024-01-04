// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml mtr is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// This element is used to display a row in a table.
type MathMLMTRElement struct {
	*Element
}

// Create a new MathMLMTRElement element.
// This will create a new element with the tag
// "mtr" during rendering.
func MathML_MTR(children ...ElementRenderer) *MathMLMTRElement {
	e := NewElement("mtr", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMTRElement{Element: e}
}

func (e *MathMLMTRElement) Children(children ...ElementRenderer) *MathMLMTRElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMTRElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMTRElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMTRElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMTRElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMTRElement) Attr(name string, value string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMTRElement) Attrs(attrs ...string) *MathMLMTRElement {
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

func (e *MathMLMTRElement) AttrsMap(attrs map[string]string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMTRElement) Text(text string) *MathMLMTRElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMTRElement) TextF(format string, args ...any) *MathMLMTRElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMTRElement) IfText(condition bool, text string) *MathMLMTRElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMTRElement) IfTextF(condition bool, format string, args ...any) *MathMLMTRElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMTRElement) Escaped(text string) *MathMLMTRElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMTRElement) IfEscaped(condition bool, text string) *MathMLMTRElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMTRElement) EscapedF(format string, args ...any) *MathMLMTRElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMTRElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMTRElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMTRElement) CustomData(key, value string) *MathMLMTRElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMTRElement) IfCustomData(condition bool, key, value string) *MathMLMTRElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMTRElement) CustomDataF(key, format string, args ...any) *MathMLMTRElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMTRElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMTRElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMTRElement) CustomDataRemove(key string) *MathMLMTRElement {
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
func (e *MathMLMTRElement) CLASS(s ...string) *MathMLMTRElement {
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

func (e *MathMLMTRElement) IfCLASS(condition bool, s ...string) *MathMLMTRElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMTRElement) CLASSRemove(s ...string) *MathMLMTRElement {
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
func (e *MathMLMTRElement) DIR(c MathMLMtrDirChoice) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMtrDirChoice string

const (
	// left-to-right
	MathMLMtrDir_ltr MathMLMtrDirChoice = "ltr"
	// right-to-left
	MathMLMtrDir_rtl MathMLMtrDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMTRElement) DIRRemove(c MathMLMtrDirChoice) *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMTRElement) DISPLAYSTYLE(c MathMLMtrDisplaystyleChoice) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMtrDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMtrDisplaystyle_true MathMLMtrDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMtrDisplaystyle_false MathMLMtrDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMTRElement) DISPLAYSTYLERemove(c MathMLMtrDisplaystyleChoice) *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMTRElement) ID(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMTRElement) IDF(format string, args ...any) *MathMLMTRElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMTRElement) IfID(condition bool, s string) *MathMLMTRElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMTRElement) IfIDF(condition bool, format string, args ...any) *MathMLMTRElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMTRElement) IDRemove(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *MathMLMTRElement) IDRemoveF(format string, args ...any) *MathMLMTRElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMTRElement) MATHBACKGROUND(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMTRElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMTRElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMTRElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMTRElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMTRElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMTRElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMTRElement) MATHBACKGROUNDRemove(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

func (e *MathMLMTRElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMTRElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMTRElement) MATHCOLOR(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMTRElement) MATHCOLORF(format string, args ...any) *MathMLMTRElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMTRElement) IfMATHCOLOR(condition bool, s string) *MathMLMTRElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMTRElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMTRElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMTRElement) MATHCOLORRemove(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

func (e *MathMLMTRElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMTRElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMTRElement) MATHSIZE_STR(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMTRElement) MATHSIZE_STRF(format string, args ...any) *MathMLMTRElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMTRElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMTRElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMTRElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMTRElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMTRElement) MATHSIZE_STRRemove(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathsize")
	return e
}

func (e *MathMLMTRElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMTRElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMTRElement) NONCE(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMTRElement) NONCEF(format string, args ...any) *MathMLMTRElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMTRElement) IfNONCE(condition bool, s string) *MathMLMTRElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMTRElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMTRElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMTRElement) NONCERemove(s string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

func (e *MathMLMTRElement) NONCERemoveF(format string, args ...any) *MathMLMTRElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMTRElement) SCRIPTLEVEL(i int) *MathMLMTRElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMTRElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMTRElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMTRElement) SCRIPTLEVELRemove(i int) *MathMLMTRElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMTRElement) STYLEF(k string, format string, args ...any) *MathMLMTRElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMTRElement) IfSTYLE(condition bool, k string, v string) *MathMLMTRElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMTRElement) STYLE(k string, v string) *MathMLMTRElement {
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

func (e *MathMLMTRElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMTRElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMTRElement) STYLEMap(m map[string]string) *MathMLMTRElement {
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
func (e *MathMLMTRElement) STYLEPairs(pairs ...string) *MathMLMTRElement {
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

func (e *MathMLMTRElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMTRElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMTRElement) STYLERemove(keys ...string) *MathMLMTRElement {
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
func (e *MathMLMTRElement) TABINDEX(i int) *MathMLMTRElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMTRElement) IfTABINDEX(condition bool, i int) *MathMLMTRElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMTRElement) TABINDEXRemove(i int) *MathMLMTRElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMTRElement) DATASTAR_MERGE_STORE(v any) *MathMLMTRElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("merge-store", string(b))
	return e
}

// Sets the reference of the element

func (e *MathMLMTRElement) DATASTAR_REF(expression string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMTRElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMTRElement) DATASTAR_REFRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMTRElement) DATASTAR_BIND(key string, expression string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMTRElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMTRElement) DATASTAR_BINDRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMTRElement) DATASTAR_MODEL(expression string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMTRElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMTRElement) DATASTAR_MODELRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMTRElement) DATASTAR_TEXT(expression string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMTRElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMTRElement) DATASTAR_TEXTRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMtrOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMtrOnModDebounce(
	d time.Duration,
) MathMLMtrOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMtrOnModThrottle(
	d time.Duration,
) MathMLMtrOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMTRElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMtrOnMod) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMtrOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMtrOnMod) *MathMLMTRElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMTRElement) DATASTAR_ONRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMTRElement) DATASTAR_FOCUSSet(b bool) *MathMLMTRElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMTRElement) DATASTAR_FOCUS() *MathMLMTRElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMTRElement) DATASTAR_HEADER(key string, expression string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMTRElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMTRElement) DATASTAR_HEADERRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMTRElement) DATASTAR_FETCH_URL(expression string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMTRElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMTRElement) DATASTAR_FETCH_URLRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMTRElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMTRElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMTRElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMTRElement) DATASTAR_SHOWSet(b bool) *MathMLMTRElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMTRElement) DATASTAR_SHOW() *MathMLMTRElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMTRElement) DATASTAR_INTERSECTS(expression string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *MathMLMTRElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *MathMLMTRElement) DATASTAR_INTERSECTSRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *MathMLMTRElement) DATASTAR_TELEPORTSet(b bool) *MathMLMTRElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMTRElement) DATASTAR_TELEPORT() *MathMLMTRElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMTRElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMTRElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMTRElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMTRElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMTRElement) DATASTAR_VIEW_TRANSITION(expression string) *MathMLMTRElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMTRElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *MathMLMTRElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMTRElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMTRElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
