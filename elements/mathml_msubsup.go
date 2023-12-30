// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml msubsup is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// This element is used to display a subscript expression with a superscript
// expression.
type MathMLMSUBSUPElement struct {
	*Element
}

// Create a new MathMLMSUBSUPElement element.
// This will create a new element with the tag
// "msubsup" during rendering.
func MathML_MSUBSUP(children ...ElementRenderer) *MathMLMSUBSUPElement {
	e := NewElement("msubsup", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMSUBSUPElement{Element: e}
}

func (e *MathMLMSUBSUPElement) Children(children ...ElementRenderer) *MathMLMSUBSUPElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMSUBSUPElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMSUBSUPElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMSUBSUPElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMSUBSUPElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMSUBSUPElement) Attr(name string, value string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMSUBSUPElement) Attrs(attrs ...string) *MathMLMSUBSUPElement {
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

func (e *MathMLMSUBSUPElement) AttrsMap(attrs map[string]string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMSUBSUPElement) Text(text string) *MathMLMSUBSUPElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMSUBSUPElement) TextF(format string, args ...any) *MathMLMSUBSUPElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBSUPElement) IfText(condition bool, text string) *MathMLMSUBSUPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMSUBSUPElement) IfTextF(condition bool, format string, args ...any) *MathMLMSUBSUPElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMSUBSUPElement) Escaped(text string) *MathMLMSUBSUPElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMSUBSUPElement) IfEscaped(condition bool, text string) *MathMLMSUBSUPElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMSUBSUPElement) EscapedF(format string, args ...any) *MathMLMSUBSUPElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBSUPElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMSUBSUPElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMSUBSUPElement) CustomData(key, value string) *MathMLMSUBSUPElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMSUBSUPElement) IfCustomData(condition bool, key, value string) *MathMLMSUBSUPElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMSUBSUPElement) CustomDataF(key, format string, args ...any) *MathMLMSUBSUPElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBSUPElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMSUBSUPElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMSUBSUPElement) CustomDataRemove(key string) *MathMLMSUBSUPElement {
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
func (e *MathMLMSUBSUPElement) CLASS(s ...string) *MathMLMSUBSUPElement {
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

func (e *MathMLMSUBSUPElement) IfCLASS(condition bool, s ...string) *MathMLMSUBSUPElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMSUBSUPElement) CLASSRemove(s ...string) *MathMLMSUBSUPElement {
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
func (e *MathMLMSUBSUPElement) DIR(c MathMLMsubsupDirChoice) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMsubsupDirChoice string

const (
	// left-to-right
	MathMLMsubsupDir_ltr MathMLMsubsupDirChoice = "ltr"
	// right-to-left
	MathMLMsubsupDir_rtl MathMLMsubsupDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMSUBSUPElement) DIRRemove(c MathMLMsubsupDirChoice) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMSUBSUPElement) DISPLAYSTYLE(c MathMLMsubsupDisplaystyleChoice) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMsubsupDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMsubsupDisplaystyle_true MathMLMsubsupDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMsubsupDisplaystyle_false MathMLMsubsupDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMSUBSUPElement) DISPLAYSTYLERemove(c MathMLMsubsupDisplaystyleChoice) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMSUBSUPElement) ID(s string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMSUBSUPElement) IfID(condition bool, s string) *MathMLMSUBSUPElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMSUBSUPElement) IDRemove(s string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMSUBSUPElement) MATHBACKGROUND(s string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMSUBSUPElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMSUBSUPElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMSUBSUPElement) MATHBACKGROUNDRemove(s string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMSUBSUPElement) MATHCOLOR(s string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMSUBSUPElement) IfMATHCOLOR(condition bool, s string) *MathMLMSUBSUPElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMSUBSUPElement) MATHCOLORRemove(s string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMSUBSUPElement) MATHSIZE_STR(s string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMSUBSUPElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMSUBSUPElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMSUBSUPElement) MATHSIZE_STRRemove(s string) *MathMLMSUBSUPElement {
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
func (e *MathMLMSUBSUPElement) NONCE(s string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMSUBSUPElement) IfNONCE(condition bool, s string) *MathMLMSUBSUPElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMSUBSUPElement) NONCERemove(s string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMSUBSUPElement) SCRIPTLEVEL(i int) *MathMLMSUBSUPElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMSUBSUPElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMSUBSUPElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMSUBSUPElement) SCRIPTLEVELRemove(i int) *MathMLMSUBSUPElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMSUBSUPElement) STYLEF(k string, format string, args ...any) *MathMLMSUBSUPElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBSUPElement) IfSTYLE(condition bool, k string, v string) *MathMLMSUBSUPElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMSUBSUPElement) STYLE(k string, v string) *MathMLMSUBSUPElement {
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

func (e *MathMLMSUBSUPElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMSUBSUPElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMSUBSUPElement) STYLEMap(m map[string]string) *MathMLMSUBSUPElement {
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
func (e *MathMLMSUBSUPElement) STYLEPairs(pairs ...string) *MathMLMSUBSUPElement {
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

func (e *MathMLMSUBSUPElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMSUBSUPElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMSUBSUPElement) STYLERemove(keys ...string) *MathMLMSUBSUPElement {
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
func (e *MathMLMSUBSUPElement) TABINDEX(i int) *MathMLMSUBSUPElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMSUBSUPElement) IfTABINDEX(condition bool, i int) *MathMLMSUBSUPElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMSUBSUPElement) TABINDEXRemove(i int) *MathMLMSUBSUPElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMSUBSUPElement) DATASTAR_MERGE_STORE(v any) *MathMLMSUBSUPElement {
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

func (e *MathMLMSUBSUPElement) DATASTAR_REF(expression string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBSUPElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMSUBSUPElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMSUBSUPElement) DATASTAR_REFRemove() *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMSUBSUPElement) DATASTAR_BIND(key string, expression string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBSUPElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMSUBSUPElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMSUBSUPElement) DATASTAR_BINDRemove() *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMSUBSUPElement) DATASTAR_MODEL(expression string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBSUPElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMSUBSUPElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMSUBSUPElement) DATASTAR_MODELRemove() *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMSUBSUPElement) DATASTAR_TEXT(expression string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBSUPElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMSUBSUPElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMSUBSUPElement) DATASTAR_TEXTRemove() *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMsubsupDataOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMsubsupDataOnModDebounce(
	d time.Duration,
) MathMLMsubsupDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMsubsupDataOnModThrottle(
	d time.Duration,
) MathMLMsubsupDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMSUBSUPElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMsubsupDataOnMod) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMsubsupDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBSUPElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMsubsupDataOnMod) *MathMLMSUBSUPElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMSUBSUPElement) DATASTAR_ONRemove() *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMSUBSUPElement) DATASTAR_FOCUSSet(b bool) *MathMLMSUBSUPElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSUBSUPElement) DATASTAR_FOCUS() *MathMLMSUBSUPElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMSUBSUPElement) DATASTAR_HEADER(key string, expression string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBSUPElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMSUBSUPElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMSUBSUPElement) DATASTAR_HEADERRemove() *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMSUBSUPElement) DATASTAR_FETCH_URL(expression string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBSUPElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMSUBSUPElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMSUBSUPElement) DATASTAR_FETCH_URLRemove() *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMSUBSUPElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBSUPElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMSUBSUPElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMSUBSUPElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMSUBSUPElement) DATASTAR_SHOWSet(b bool) *MathMLMSUBSUPElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSUBSUPElement) DATASTAR_SHOW() *MathMLMSUBSUPElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMSUBSUPElement) DATASTAR_INTERSECTSSet(b bool) *MathMLMSUBSUPElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSUBSUPElement) DATASTAR_INTERSECTS() *MathMLMSUBSUPElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *MathMLMSUBSUPElement) DATASTAR_TELEPORTSet(b bool) *MathMLMSUBSUPElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSUBSUPElement) DATASTAR_TELEPORT() *MathMLMSUBSUPElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMSUBSUPElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMSUBSUPElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSUBSUPElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMSUBSUPElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMSUBSUPElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSUBSUPElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *MathMLMSUBSUPElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMSUBSUPElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMSUBSUPElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
