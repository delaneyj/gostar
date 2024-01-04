// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml mstyle is generated from configuration file.
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

// This element is used to specify the style of a subexpression.
type MathMLMSTYLEElement struct {
	*Element
}

// Create a new MathMLMSTYLEElement element.
// This will create a new element with the tag
// "mstyle" during rendering.
func MathML_MSTYLE(children ...ElementRenderer) *MathMLMSTYLEElement {
	e := NewElement("mstyle", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMSTYLEElement{Element: e}
}

func (e *MathMLMSTYLEElement) Children(children ...ElementRenderer) *MathMLMSTYLEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMSTYLEElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMSTYLEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMSTYLEElement) Attr(name string, value string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMSTYLEElement) Attrs(attrs ...string) *MathMLMSTYLEElement {
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

func (e *MathMLMSTYLEElement) AttrsMap(attrs map[string]string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMSTYLEElement) Text(text string) *MathMLMSTYLEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMSTYLEElement) TextF(format string, args ...any) *MathMLMSTYLEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMSTYLEElement) IfText(condition bool, text string) *MathMLMSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMSTYLEElement) IfTextF(condition bool, format string, args ...any) *MathMLMSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMSTYLEElement) Escaped(text string) *MathMLMSTYLEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMSTYLEElement) IfEscaped(condition bool, text string) *MathMLMSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMSTYLEElement) EscapedF(format string, args ...any) *MathMLMSTYLEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMSTYLEElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMSTYLEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMSTYLEElement) CustomData(key, value string) *MathMLMSTYLEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMSTYLEElement) IfCustomData(condition bool, key, value string) *MathMLMSTYLEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMSTYLEElement) CustomDataF(key, format string, args ...any) *MathMLMSTYLEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMSTYLEElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMSTYLEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMSTYLEElement) CustomDataRemove(key string) *MathMLMSTYLEElement {
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
func (e *MathMLMSTYLEElement) CLASS(s ...string) *MathMLMSTYLEElement {
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

func (e *MathMLMSTYLEElement) IfCLASS(condition bool, s ...string) *MathMLMSTYLEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMSTYLEElement) CLASSRemove(s ...string) *MathMLMSTYLEElement {
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
func (e *MathMLMSTYLEElement) DIR(c MathMLMstyleDirChoice) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMstyleDirChoice string

const (
	// left-to-right
	MathMLMstyleDir_ltr MathMLMstyleDirChoice = "ltr"
	// right-to-left
	MathMLMstyleDir_rtl MathMLMstyleDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMSTYLEElement) DIRRemove(c MathMLMstyleDirChoice) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMSTYLEElement) DISPLAYSTYLE(c MathMLMstyleDisplaystyleChoice) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMstyleDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMstyleDisplaystyle_true MathMLMstyleDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMstyleDisplaystyle_false MathMLMstyleDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMSTYLEElement) DISPLAYSTYLERemove(c MathMLMstyleDisplaystyleChoice) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMSTYLEElement) ID(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMSTYLEElement) IDF(format string, args ...any) *MathMLMSTYLEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMSTYLEElement) IfID(condition bool, s string) *MathMLMSTYLEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMSTYLEElement) IfIDF(condition bool, format string, args ...any) *MathMLMSTYLEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMSTYLEElement) IDRemove(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *MathMLMSTYLEElement) IDRemoveF(format string, args ...any) *MathMLMSTYLEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMSTYLEElement) MATHBACKGROUND(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMSTYLEElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMSTYLEElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMSTYLEElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMSTYLEElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMSTYLEElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMSTYLEElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMSTYLEElement) MATHBACKGROUNDRemove(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

func (e *MathMLMSTYLEElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMSTYLEElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMSTYLEElement) MATHCOLOR(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMSTYLEElement) MATHCOLORF(format string, args ...any) *MathMLMSTYLEElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMSTYLEElement) IfMATHCOLOR(condition bool, s string) *MathMLMSTYLEElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMSTYLEElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMSTYLEElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMSTYLEElement) MATHCOLORRemove(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

func (e *MathMLMSTYLEElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMSTYLEElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMSTYLEElement) MATHSIZE_STR(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMSTYLEElement) MATHSIZE_STRF(format string, args ...any) *MathMLMSTYLEElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMSTYLEElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMSTYLEElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMSTYLEElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMSTYLEElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMSTYLEElement) MATHSIZE_STRRemove(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathsize")
	return e
}

func (e *MathMLMSTYLEElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMSTYLEElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMSTYLEElement) NONCE(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMSTYLEElement) NONCEF(format string, args ...any) *MathMLMSTYLEElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMSTYLEElement) IfNONCE(condition bool, s string) *MathMLMSTYLEElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMSTYLEElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMSTYLEElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMSTYLEElement) NONCERemove(s string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

func (e *MathMLMSTYLEElement) NONCERemoveF(format string, args ...any) *MathMLMSTYLEElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMSTYLEElement) SCRIPTLEVEL(i int) *MathMLMSTYLEElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMSTYLEElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMSTYLEElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMSTYLEElement) SCRIPTLEVELRemove(i int) *MathMLMSTYLEElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMSTYLEElement) STYLEF(k string, format string, args ...any) *MathMLMSTYLEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMSTYLEElement) IfSTYLE(condition bool, k string, v string) *MathMLMSTYLEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMSTYLEElement) STYLE(k string, v string) *MathMLMSTYLEElement {
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

func (e *MathMLMSTYLEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMSTYLEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMSTYLEElement) STYLEMap(m map[string]string) *MathMLMSTYLEElement {
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
func (e *MathMLMSTYLEElement) STYLEPairs(pairs ...string) *MathMLMSTYLEElement {
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

func (e *MathMLMSTYLEElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMSTYLEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMSTYLEElement) STYLERemove(keys ...string) *MathMLMSTYLEElement {
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
func (e *MathMLMSTYLEElement) TABINDEX(i int) *MathMLMSTYLEElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMSTYLEElement) IfTABINDEX(condition bool, i int) *MathMLMSTYLEElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMSTYLEElement) TABINDEXRemove(i int) *MathMLMSTYLEElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMSTYLEElement) DATASTAR_MERGE_STORE(v any) *MathMLMSTYLEElement {
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

func (e *MathMLMSTYLEElement) DATASTAR_REF(expression string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMSTYLEElement) DATASTAR_REFRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMSTYLEElement) DATASTAR_BIND(key string, expression string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMSTYLEElement) DATASTAR_BINDRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMSTYLEElement) DATASTAR_MODEL(expression string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMSTYLEElement) DATASTAR_MODELRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMSTYLEElement) DATASTAR_TEXT(expression string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMSTYLEElement) DATASTAR_TEXTRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMstyleOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMstyleOnModDebounce(
	d time.Duration,
) MathMLMstyleOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMstyleOnModThrottle(
	d time.Duration,
) MathMLMstyleOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMSTYLEElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMstyleOnMod) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMstyleOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMstyleOnMod) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMSTYLEElement) DATASTAR_ONRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMSTYLEElement) DATASTAR_FOCUSSet(b bool) *MathMLMSTYLEElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSTYLEElement) DATASTAR_FOCUS() *MathMLMSTYLEElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMSTYLEElement) DATASTAR_HEADER(key string, expression string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMSTYLEElement) DATASTAR_HEADERRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMSTYLEElement) DATASTAR_FETCH_URL(expression string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMSTYLEElement) DATASTAR_FETCH_URLRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMSTYLEElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMSTYLEElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMSTYLEElement) DATASTAR_SHOWSet(b bool) *MathMLMSTYLEElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSTYLEElement) DATASTAR_SHOW() *MathMLMSTYLEElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMSTYLEElement) DATASTAR_INTERSECTS(expression string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *MathMLMSTYLEElement) DATASTAR_INTERSECTSRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *MathMLMSTYLEElement) DATASTAR_TELEPORTSet(b bool) *MathMLMSTYLEElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSTYLEElement) DATASTAR_TELEPORT() *MathMLMSTYLEElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMSTYLEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMSTYLEElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSTYLEElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMSTYLEElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMSTYLEElement) DATASTAR_VIEW_TRANSITION(expression string) *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSTYLEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *MathMLMSTYLEElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMSTYLEElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMSTYLEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
