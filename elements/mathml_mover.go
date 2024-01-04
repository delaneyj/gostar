// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml mover is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// This element is used to display an expression with an overbar.
type MathMLMOVERElement struct {
	*Element
}

// Create a new MathMLMOVERElement element.
// This will create a new element with the tag
// "mover" during rendering.
func MathML_MOVER(children ...ElementRenderer) *MathMLMOVERElement {
	e := NewElement("mover", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMOVERElement{Element: e}
}

func (e *MathMLMOVERElement) Children(children ...ElementRenderer) *MathMLMOVERElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMOVERElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMOVERElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMOVERElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMOVERElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMOVERElement) Attr(name string, value string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMOVERElement) Attrs(attrs ...string) *MathMLMOVERElement {
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

func (e *MathMLMOVERElement) AttrsMap(attrs map[string]string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMOVERElement) Text(text string) *MathMLMOVERElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMOVERElement) TextF(format string, args ...any) *MathMLMOVERElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMOVERElement) IfText(condition bool, text string) *MathMLMOVERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMOVERElement) IfTextF(condition bool, format string, args ...any) *MathMLMOVERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMOVERElement) Escaped(text string) *MathMLMOVERElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMOVERElement) IfEscaped(condition bool, text string) *MathMLMOVERElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMOVERElement) EscapedF(format string, args ...any) *MathMLMOVERElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMOVERElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMOVERElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMOVERElement) CustomData(key, value string) *MathMLMOVERElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMOVERElement) IfCustomData(condition bool, key, value string) *MathMLMOVERElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMOVERElement) CustomDataF(key, format string, args ...any) *MathMLMOVERElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMOVERElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMOVERElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMOVERElement) CustomDataRemove(key string) *MathMLMOVERElement {
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
func (e *MathMLMOVERElement) CLASS(s ...string) *MathMLMOVERElement {
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

func (e *MathMLMOVERElement) IfCLASS(condition bool, s ...string) *MathMLMOVERElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMOVERElement) CLASSRemove(s ...string) *MathMLMOVERElement {
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
func (e *MathMLMOVERElement) DIR(c MathMLMoverDirChoice) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMoverDirChoice string

const (
	// left-to-right
	MathMLMoverDir_ltr MathMLMoverDirChoice = "ltr"
	// right-to-left
	MathMLMoverDir_rtl MathMLMoverDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMOVERElement) DIRRemove(c MathMLMoverDirChoice) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMOVERElement) DISPLAYSTYLE(c MathMLMoverDisplaystyleChoice) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMoverDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMoverDisplaystyle_true MathMLMoverDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMoverDisplaystyle_false MathMLMoverDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMOVERElement) DISPLAYSTYLERemove(c MathMLMoverDisplaystyleChoice) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMOVERElement) ID(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMOVERElement) IDF(format string, args ...any) *MathMLMOVERElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMOVERElement) IfID(condition bool, s string) *MathMLMOVERElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMOVERElement) IfIDF(condition bool, format string, args ...any) *MathMLMOVERElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMOVERElement) IDRemove(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *MathMLMOVERElement) IDRemoveF(format string, args ...any) *MathMLMOVERElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMOVERElement) MATHBACKGROUND(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMOVERElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMOVERElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMOVERElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMOVERElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMOVERElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMOVERElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMOVERElement) MATHBACKGROUNDRemove(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

func (e *MathMLMOVERElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMOVERElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMOVERElement) MATHCOLOR(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMOVERElement) MATHCOLORF(format string, args ...any) *MathMLMOVERElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMOVERElement) IfMATHCOLOR(condition bool, s string) *MathMLMOVERElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMOVERElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMOVERElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMOVERElement) MATHCOLORRemove(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

func (e *MathMLMOVERElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMOVERElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMOVERElement) MATHSIZE_STR(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMOVERElement) MATHSIZE_STRF(format string, args ...any) *MathMLMOVERElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMOVERElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMOVERElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMOVERElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMOVERElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMOVERElement) MATHSIZE_STRRemove(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathsize")
	return e
}

func (e *MathMLMOVERElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMOVERElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMOVERElement) NONCE(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMOVERElement) NONCEF(format string, args ...any) *MathMLMOVERElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMOVERElement) IfNONCE(condition bool, s string) *MathMLMOVERElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMOVERElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMOVERElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMOVERElement) NONCERemove(s string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

func (e *MathMLMOVERElement) NONCERemoveF(format string, args ...any) *MathMLMOVERElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMOVERElement) SCRIPTLEVEL(i int) *MathMLMOVERElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMOVERElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMOVERElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMOVERElement) SCRIPTLEVELRemove(i int) *MathMLMOVERElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMOVERElement) STYLEF(k string, format string, args ...any) *MathMLMOVERElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMOVERElement) IfSTYLE(condition bool, k string, v string) *MathMLMOVERElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMOVERElement) STYLE(k string, v string) *MathMLMOVERElement {
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

func (e *MathMLMOVERElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMOVERElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMOVERElement) STYLEMap(m map[string]string) *MathMLMOVERElement {
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
func (e *MathMLMOVERElement) STYLEPairs(pairs ...string) *MathMLMOVERElement {
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

func (e *MathMLMOVERElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMOVERElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMOVERElement) STYLERemove(keys ...string) *MathMLMOVERElement {
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
func (e *MathMLMOVERElement) TABINDEX(i int) *MathMLMOVERElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMOVERElement) IfTABINDEX(condition bool, i int) *MathMLMOVERElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMOVERElement) TABINDEXRemove(i int) *MathMLMOVERElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMOVERElement) DATASTAR_MERGE_STORE(v any) *MathMLMOVERElement {
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

func (e *MathMLMOVERElement) DATASTAR_REF(expression string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMOVERElement) DATASTAR_REFRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMOVERElement) DATASTAR_BIND(key string, expression string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMOVERElement) DATASTAR_BINDRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMOVERElement) DATASTAR_MODEL(expression string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMOVERElement) DATASTAR_MODELRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMOVERElement) DATASTAR_TEXT(expression string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMOVERElement) DATASTAR_TEXTRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMoverOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMoverOnModDebounce(
	d time.Duration,
) MathMLMoverOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMoverOnModThrottle(
	d time.Duration,
) MathMLMoverOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMOVERElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMoverOnMod) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMoverOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMoverOnMod) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMOVERElement) DATASTAR_ONRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMOVERElement) DATASTAR_FOCUSSet(b bool) *MathMLMOVERElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMOVERElement) DATASTAR_FOCUS() *MathMLMOVERElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMOVERElement) DATASTAR_HEADER(key string, expression string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMOVERElement) DATASTAR_HEADERRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMOVERElement) DATASTAR_FETCH_URL(expression string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMOVERElement) DATASTAR_FETCH_URLRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMOVERElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMOVERElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMOVERElement) DATASTAR_SHOWSet(b bool) *MathMLMOVERElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMOVERElement) DATASTAR_SHOW() *MathMLMOVERElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMOVERElement) DATASTAR_INTERSECTS(expression string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *MathMLMOVERElement) DATASTAR_INTERSECTSRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *MathMLMOVERElement) DATASTAR_TELEPORTSet(b bool) *MathMLMOVERElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMOVERElement) DATASTAR_TELEPORT() *MathMLMOVERElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMOVERElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMOVERElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMOVERElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMOVERElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMOVERElement) DATASTAR_VIEW_TRANSITION(expression string) *MathMLMOVERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMOVERElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *MathMLMOVERElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMOVERElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMOVERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
