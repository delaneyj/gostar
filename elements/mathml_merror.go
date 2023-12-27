// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml merror is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
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

func (e *MathMLMERRORElement) IfText(condition bool, text string) *MathMLMERRORElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLMERRORElement) IfTextF(condition bool, format string, args ...any) *MathMLMERRORElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLMERRORElement) Escaped(text string) *MathMLMERRORElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLMERRORElement) IfEscaped(condition bool, text string) *MathMLMERRORElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLMERRORElement) EscapedF(format string, args ...any) *MathMLMERRORElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMERRORElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMERRORElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLMERRORElement) CustomData(key, value string) *MathMLMERRORElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMERRORElement) IfCustomData(condition bool, key, value string) *MathMLMERRORElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLMERRORElement) CustomDataF(key, format string, args ...any) *MathMLMERRORElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMERRORElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMERRORElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
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

func (e *MathMLMERRORElement) IfCLASS(condition bool, s ...string) *MathMLMERRORElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
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

// Remove the attribute DIR from the element.
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

// Remove the attribute DISPLAYSTYLE from the element.
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

func (e *MathMLMERRORElement) IfID(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
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

func (e *MathMLMERRORElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
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

func (e *MathMLMERRORElement) IfMATHCOLOR(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLMERRORElement) MATHCOLORRemove(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLMERRORElement) MATHSIZE_STR(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLMERRORElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLMERRORElement) MATHSIZE_STRRemove(s string) *MathMLMERRORElement {
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

func (e *MathMLMERRORElement) IfNONCE(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

// Remove the attribute NONCE from the element.
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

func (e *MathMLMERRORElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMERRORElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
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

func (e *MathMLMERRORElement) IfSTYLE(condition bool, k string, v string) *MathMLMERRORElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
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

func (e *MathMLMERRORElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMERRORElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
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

func (e *MathMLMERRORElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMERRORElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
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

func (e *MathMLMERRORElement) IfTABINDEX(condition bool, i int) *MathMLMERRORElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLMERRORElement) TABINDEXRemove(i int) *MathMLMERRORElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the store with the given object

func (e *MathMLMERRORElement) DATASTAR_MERGE_STORE(v any) *MathMLMERRORElement {
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

func (e *MathMLMERRORElement) DATASTAR_REF(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-ref"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *MathMLMERRORElement) IfDATASTAR_REF(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.DATASTAR_REF(s)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLMERRORElement) DATASTAR_REFRemove() *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLMERRORElement) DATASTAR_BIND(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-bind"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *MathMLMERRORElement) IfDATASTAR_BIND(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.DATASTAR_BIND(s)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLMERRORElement) DATASTAR_BINDRemove() *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLMERRORElement) DATASTAR_MODEL(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-model"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *MathMLMERRORElement) IfDATASTAR_MODEL(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.DATASTAR_MODEL(s)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLMERRORElement) DATASTAR_MODELRemove() *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLMERRORElement) DATASTAR_TEXT(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-text"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *MathMLMERRORElement) IfDATASTAR_TEXT(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.DATASTAR_TEXT(s)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLMERRORElement) DATASTAR_TEXTRemove() *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLMerrorDataOnMod customDataKeyModifier

// Debounces the event handler
func MathMLMerrorDataOnModDebounce(
	s string,
) MathMLMerrorDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%sms", s)
	}
}

// Throttles the event handler
func MathMLMerrorDataOnModThrottle(
	s string,
) MathMLMerrorDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%sms", s)
	}
}

func (e *MathMLMERRORElement) DATASTAR_ON(s string, modifiers ...MathMLMerrorDataOnMod) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	customMods := lo.Map(modifiers, func(m MathMLMerrorDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key := customDataKey("data-on", customMods...)
	e.StringAttributes.Set(key, s)
	return e
}

func (e *MathMLMERRORElement) IfDATASTAR_ON(condition bool, s string, modifiers ...MathMLMerrorDataOnMod) *MathMLMERRORElement {
	if condition {
		e.DATASTAR_ON(s, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLMERRORElement) DATASTAR_ONRemove() *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLMERRORElement) DATASTAR_FOCUSSet(b bool) *MathMLMERRORElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMERRORElement) DATASTAR_FOCUS() *MathMLMERRORElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLMERRORElement) DATASTAR_HEADER(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-header"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *MathMLMERRORElement) IfDATASTAR_HEADER(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.DATASTAR_HEADER(s)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLMERRORElement) DATASTAR_HEADERRemove() *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *MathMLMERRORElement) DATASTAR_FETCH_URL(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-fetch-url"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *MathMLMERRORElement) IfDATASTAR_FETCH_URL(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.DATASTAR_FETCH_URL(s)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *MathMLMERRORElement) DATASTAR_FETCH_URLRemove() *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLMERRORElement) DATASTAR_FETCH_INDICATOR(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "DatastarFetchIndicator"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *MathMLMERRORElement) IfDATASTAR_FETCH_INDICATOR(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(s)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLMERRORElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLMERRORElement) DATASTAR_SHOWSet(b bool) *MathMLMERRORElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMERRORElement) DATASTAR_SHOW() *MathMLMERRORElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLMERRORElement) DATASTAR_INTERSECTSSet(b bool) *MathMLMERRORElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMERRORElement) DATASTAR_INTERSECTS() *MathMLMERRORElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *MathMLMERRORElement) DATASTAR_TELEPORTSet(b bool) *MathMLMERRORElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMERRORElement) DATASTAR_TELEPORT() *MathMLMERRORElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLMERRORElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMERRORElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLMERRORElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMERRORElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLMERRORElement) DATASTAR_VIEW_TRANSITION(s string) *MathMLMERRORElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	key := "data-view-transition"
	e.StringAttributes.Set(key, s)
	return e
}

func (e *MathMLMERRORElement) IfDATASTAR_VIEW_TRANSITION(condition bool, s string) *MathMLMERRORElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(s)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLMERRORElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMERRORElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
