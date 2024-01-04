// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml mspace is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// This element is used to display a space.
type MathMLMSPACEElement struct {
	*Element
}

// Create a new MathMLMSPACEElement element.
// This will create a new element with the tag
// "mspace" during rendering.
func MathML_MSPACE(children ...ElementRenderer) *MathMLMSPACEElement {
	e := NewElement("mspace", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLMSPACEElement{Element: e}
}

func (e *MathMLMSPACEElement) Children(children ...ElementRenderer) *MathMLMSPACEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLMSPACEElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMSPACEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLMSPACEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMSPACEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLMSPACEElement) Attr(name string, value string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLMSPACEElement) Attrs(attrs ...string) *MathMLMSPACEElement {
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

func (e *MathMLMSPACEElement) AttrsMap(attrs map[string]string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLMSPACEElement) Text(text string) *MathMLMSPACEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLMSPACEElement) TextF(format string, args ...any) *MathMLMSPACEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMSPACEElement) IfText(condition bool, text string) *MathMLMSPACEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMSPACEElement) IfTextF(condition bool, format string, args ...any) *MathMLMSPACEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMSPACEElement) Escaped(text string) *MathMLMSPACEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMSPACEElement) IfEscaped(condition bool, text string) *MathMLMSPACEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMSPACEElement) EscapedF(format string, args ...any) *MathMLMSPACEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMSPACEElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMSPACEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMSPACEElement) CustomData(key, value string) *MathMLMSPACEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMSPACEElement) IfCustomData(condition bool, key, value string) *MathMLMSPACEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMSPACEElement) CustomDataF(key, format string, args ...any) *MathMLMSPACEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMSPACEElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMSPACEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLMSPACEElement) CustomDataRemove(key string) *MathMLMSPACEElement {
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
func (e *MathMLMSPACEElement) CLASS(s ...string) *MathMLMSPACEElement {
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

func (e *MathMLMSPACEElement) IfCLASS(condition bool, s ...string) *MathMLMSPACEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLMSPACEElement) CLASSRemove(s ...string) *MathMLMSPACEElement {
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
func (e *MathMLMSPACEElement) DIR(c MathMLMspaceDirChoice) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLMspaceDirChoice string

const (
	// left-to-right
	MathMLMspaceDir_ltr MathMLMspaceDirChoice = "ltr"
	// right-to-left
	MathMLMspaceDir_rtl MathMLMspaceDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLMSPACEElement) DIRRemove(c MathMLMspaceDirChoice) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLMSPACEElement) DISPLAYSTYLE(c MathMLMspaceDisplaystyleChoice) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLMspaceDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLMspaceDisplaystyle_true MathMLMspaceDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLMspaceDisplaystyle_false MathMLMspaceDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLMSPACEElement) DISPLAYSTYLERemove(c MathMLMspaceDisplaystyleChoice) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLMSPACEElement) ID(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLMSPACEElement) IDF(format string, args ...any) *MathMLMSPACEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLMSPACEElement) IfID(condition bool, s string) *MathMLMSPACEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLMSPACEElement) IfIDF(condition bool, format string, args ...any) *MathMLMSPACEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLMSPACEElement) IDRemove(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *MathMLMSPACEElement) IDRemoveF(format string, args ...any) *MathMLMSPACEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMSPACEElement) MATHBACKGROUND(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLMSPACEElement) MATHBACKGROUNDF(format string, args ...any) *MathMLMSPACEElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLMSPACEElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMSPACEElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLMSPACEElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLMSPACEElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLMSPACEElement) MATHBACKGROUNDRemove(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

func (e *MathMLMSPACEElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLMSPACEElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLMSPACEElement) MATHCOLOR(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLMSPACEElement) MATHCOLORF(format string, args ...any) *MathMLMSPACEElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLMSPACEElement) IfMATHCOLOR(condition bool, s string) *MathMLMSPACEElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLMSPACEElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLMSPACEElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMSPACEElement) MATHCOLORRemove(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

func (e *MathMLMSPACEElement) MATHCOLORRemoveF(format string, args ...any) *MathMLMSPACEElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMSPACEElement) MATHSIZE_STR(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMSPACEElement) MATHSIZE_STRF(format string, args ...any) *MathMLMSPACEElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLMSPACEElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMSPACEElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLMSPACEElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLMSPACEElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMSPACEElement) MATHSIZE_STRRemove(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathsize")
	return e
}

func (e *MathMLMSPACEElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLMSPACEElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLMSPACEElement) NONCE(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLMSPACEElement) NONCEF(format string, args ...any) *MathMLMSPACEElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLMSPACEElement) IfNONCE(condition bool, s string) *MathMLMSPACEElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLMSPACEElement) IfNONCEF(condition bool, format string, args ...any) *MathMLMSPACEElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLMSPACEElement) NONCERemove(s string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

func (e *MathMLMSPACEElement) NONCERemoveF(format string, args ...any) *MathMLMSPACEElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLMSPACEElement) SCRIPTLEVEL(i int) *MathMLMSPACEElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLMSPACEElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMSPACEElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLMSPACEElement) SCRIPTLEVELRemove(i int) *MathMLMSPACEElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLMSPACEElement) STYLEF(k string, format string, args ...any) *MathMLMSPACEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLMSPACEElement) IfSTYLE(condition bool, k string, v string) *MathMLMSPACEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLMSPACEElement) STYLE(k string, v string) *MathMLMSPACEElement {
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

func (e *MathMLMSPACEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMSPACEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLMSPACEElement) STYLEMap(m map[string]string) *MathMLMSPACEElement {
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
func (e *MathMLMSPACEElement) STYLEPairs(pairs ...string) *MathMLMSPACEElement {
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

func (e *MathMLMSPACEElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMSPACEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLMSPACEElement) STYLERemove(keys ...string) *MathMLMSPACEElement {
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
func (e *MathMLMSPACEElement) TABINDEX(i int) *MathMLMSPACEElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLMSPACEElement) IfTABINDEX(condition bool, i int) *MathMLMSPACEElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMSPACEElement) TABINDEXRemove(i int) *MathMLMSPACEElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMSPACEElement) DATASTAR_MERGE_STORE(v any) *MathMLMSPACEElement {
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

func (e *MathMLMSPACEElement) DATASTAR_REF(expression string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMSPACEElement) DATASTAR_REFRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMSPACEElement) DATASTAR_BIND(key string, expression string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMSPACEElement) DATASTAR_BINDRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMSPACEElement) DATASTAR_MODEL(expression string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMSPACEElement) DATASTAR_MODELRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMSPACEElement) DATASTAR_TEXT(expression string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMSPACEElement) DATASTAR_TEXTRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMspaceOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMspaceOnModDebounce(
	d time.Duration,
) MathMLMspaceOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLMspaceOnModThrottle(
	d time.Duration,
) MathMLMspaceOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLMSPACEElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMspaceOnMod) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLMspaceOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMspaceOnMod) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMSPACEElement) DATASTAR_ONRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMSPACEElement) DATASTAR_FOCUSSet(b bool) *MathMLMSPACEElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSPACEElement) DATASTAR_FOCUS() *MathMLMSPACEElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMSPACEElement) DATASTAR_HEADER(key string, expression string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMSPACEElement) DATASTAR_HEADERRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMSPACEElement) DATASTAR_FETCH_URL(expression string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMSPACEElement) DATASTAR_FETCH_URLRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMSPACEElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMSPACEElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMSPACEElement) DATASTAR_SHOWSet(b bool) *MathMLMSPACEElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSPACEElement) DATASTAR_SHOW() *MathMLMSPACEElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMSPACEElement) DATASTAR_INTERSECTS(expression string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *MathMLMSPACEElement) DATASTAR_INTERSECTSRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *MathMLMSPACEElement) DATASTAR_TELEPORTSet(b bool) *MathMLMSPACEElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSPACEElement) DATASTAR_TELEPORT() *MathMLMSPACEElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMSPACEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMSPACEElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMSPACEElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMSPACEElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMSPACEElement) DATASTAR_VIEW_TRANSITION(expression string) *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLMSPACEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *MathMLMSPACEElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMSPACEElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMSPACEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
