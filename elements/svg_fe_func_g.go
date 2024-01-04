// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feFuncG is generated from configuration file.
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

// The <feFuncG> SVG filter primitive defines the transfer function for the green
// component of the input graphic of its parent <feComponentTransfer> element.
type SVGFEFUNCGElement struct {
	*Element
}

// Create a new SVGFEFUNCGElement element.
// This will create a new element with the tag
// "feFuncG" during rendering.
func SVG_FEFUNCG(children ...ElementRenderer) *SVGFEFUNCGElement {
	e := NewElement("feFuncG", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEFUNCGElement{Element: e}
}

func (e *SVGFEFUNCGElement) Children(children ...ElementRenderer) *SVGFEFUNCGElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEFUNCGElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEFUNCGElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEFUNCGElement) Attr(name string, value string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFEFUNCGElement) Attrs(attrs ...string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) AttrsMap(attrs map[string]string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFEFUNCGElement) Text(text string) *SVGFEFUNCGElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEFUNCGElement) TextF(format string, args ...any) *SVGFEFUNCGElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfText(condition bool, text string) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFEFUNCGElement) IfTextF(condition bool, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFEFUNCGElement) Escaped(text string) *SVGFEFUNCGElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEFUNCGElement) IfEscaped(condition bool, text string) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFEFUNCGElement) EscapedF(format string, args ...any) *SVGFEFUNCGElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFEFUNCGElement) CustomData(key, value string) *SVGFEFUNCGElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEFUNCGElement) IfCustomData(condition bool, key, value string) *SVGFEFUNCGElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFEFUNCGElement) CustomDataF(key, format string, args ...any) *SVGFEFUNCGElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFEFUNCGElement) CustomDataRemove(key string) *SVGFEFUNCGElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The type of transfer function.
func (e *SVGFEFUNCGElement) TYPE(c SVGFeFuncGTypeChoice) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeFuncGTypeChoice string

const (
	// The type of transfer function.
	SVGFeFuncGType_identity SVGFeFuncGTypeChoice = "identity"
	// The type of transfer function.
	SVGFeFuncGType_table SVGFeFuncGTypeChoice = "table"
	// The type of transfer function.
	SVGFeFuncGType_discrete SVGFeFuncGTypeChoice = "discrete"
	// The type of transfer function.
	SVGFeFuncGType_linear SVGFeFuncGTypeChoice = "linear"
	// The type of transfer function.
	SVGFeFuncGType_gamma SVGFeFuncGTypeChoice = "gamma"
)

// Remove the attribute TYPE from the element.
func (e *SVGFEFUNCGElement) TYPERemove(c SVGFeFuncGTypeChoice) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

// Contains the list of <number>s that define the lookup table
// Values must be in the 0-1 range and be equally spaced
// There must be at least two values.
func (e *SVGFEFUNCGElement) TABLE_VALUES(s string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("tableValues", s)
	return e
}

func (e *SVGFEFUNCGElement) TABLE_VALUESF(format string, args ...any) *SVGFEFUNCGElement {
	return e.TABLE_VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfTABLE_VALUES(condition bool, s string) *SVGFEFUNCGElement {
	if condition {
		e.TABLE_VALUES(s)
	}
	return e
}

func (e *SVGFEFUNCGElement) IfTABLE_VALUESF(condition bool, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.TABLE_VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TABLE_VALUES from the element.
func (e *SVGFEFUNCGElement) TABLE_VALUESRemove(s string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("tableValues")
	return e
}

func (e *SVGFEFUNCGElement) TABLE_VALUESRemoveF(format string, args ...any) *SVGFEFUNCGElement {
	return e.TABLE_VALUESRemove(fmt.Sprintf(format, args...))
}

// The slope attribute indicates the slope of the linear function.
func (e *SVGFEFUNCGElement) SLOPE(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("slope", f)
	return e
}

func (e *SVGFEFUNCGElement) IfSLOPE(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.SLOPE(f)
	}
	return e
}

// The intercept attribute indicates the intercept of the linear function.
func (e *SVGFEFUNCGElement) INTERCEPT(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("intercept", f)
	return e
}

func (e *SVGFEFUNCGElement) IfINTERCEPT(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.INTERCEPT(f)
	}
	return e
}

// The amplitude attribute indicates the amplitude of the cubic function.
func (e *SVGFEFUNCGElement) AMPLITUDE(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("amplitude", f)
	return e
}

func (e *SVGFEFUNCGElement) IfAMPLITUDE(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.AMPLITUDE(f)
	}
	return e
}

// The exponent attribute indicates the exponent of the exponential function.
func (e *SVGFEFUNCGElement) EXPONENT(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("exponent", f)
	return e
}

func (e *SVGFEFUNCGElement) IfEXPONENT(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.EXPONENT(f)
	}
	return e
}

// The offset attribute indicates the offset of the function.
func (e *SVGFEFUNCGElement) OFFSET(f float64) *SVGFEFUNCGElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("offset", f)
	return e
}

func (e *SVGFEFUNCGElement) IfOFFSET(condition bool, f float64) *SVGFEFUNCGElement {
	if condition {
		e.OFFSET(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGFEFUNCGElement) ID(s string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFEFUNCGElement) IDF(format string, args ...any) *SVGFEFUNCGElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfID(condition bool, s string) *SVGFEFUNCGElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFEFUNCGElement) IfIDF(condition bool, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFEFUNCGElement) IDRemove(s string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGFEFUNCGElement) IDRemoveF(format string, args ...any) *SVGFEFUNCGElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEFUNCGElement) CLASS(s ...string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) IfCLASS(condition bool, s ...string) *SVGFEFUNCGElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFEFUNCGElement) CLASSRemove(s ...string) *SVGFEFUNCGElement {
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

// Specifies an inline CSS style for an element
func (e *SVGFEFUNCGElement) STYLEF(k string, format string, args ...any) *SVGFEFUNCGElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFEFUNCGElement) IfSTYLE(condition bool, k string, v string) *SVGFEFUNCGElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFEFUNCGElement) STYLE(k string, v string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEFUNCGElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFEFUNCGElement) STYLEMap(m map[string]string) *SVGFEFUNCGElement {
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
func (e *SVGFEFUNCGElement) STYLEPairs(pairs ...string) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEFUNCGElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFEFUNCGElement) STYLERemove(keys ...string) *SVGFEFUNCGElement {
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

// Merges the store with the given object

func (e *SVGFEFUNCGElement) DATASTAR_MERGE_STORE(v any) *SVGFEFUNCGElement {
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

func (e *SVGFEFUNCGElement) DATASTAR_REF(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_REF(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFEFUNCGElement) DATASTAR_REFRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFEFUNCGElement) DATASTAR_BIND(key string, expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFEFUNCGElement) DATASTAR_BINDRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFEFUNCGElement) DATASTAR_MODEL(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFEFUNCGElement) DATASTAR_MODELRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFEFUNCGElement) DATASTAR_TEXT(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFEFUNCGElement) DATASTAR_TEXTRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeFuncGOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeFuncGOnModDebounce(
	d time.Duration,
) SVGFeFuncGOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeFuncGOnModThrottle(
	d time.Duration,
) SVGFeFuncGOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFEFUNCGElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeFuncGOnMod) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeFuncGOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeFuncGOnMod) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFEFUNCGElement) DATASTAR_ONRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFEFUNCGElement) DATASTAR_FOCUSSet(b bool) *SVGFEFUNCGElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEFUNCGElement) DATASTAR_FOCUS() *SVGFEFUNCGElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFEFUNCGElement) DATASTAR_HEADER(key string, expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFEFUNCGElement) DATASTAR_HEADERRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGFEFUNCGElement) DATASTAR_FETCH_URL(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGFEFUNCGElement) DATASTAR_FETCH_URLRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFEFUNCGElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFEFUNCGElement) DATASTAR_FETCH_INDICATORRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFEFUNCGElement) DATASTAR_SHOWSet(b bool) *SVGFEFUNCGElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEFUNCGElement) DATASTAR_SHOW() *SVGFEFUNCGElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFEFUNCGElement) DATASTAR_INTERSECTS(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFEFUNCGElement) DATASTAR_INTERSECTSRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFEFUNCGElement) DATASTAR_TELEPORTSet(b bool) *SVGFEFUNCGElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEFUNCGElement) DATASTAR_TELEPORT() *SVGFEFUNCGElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFEFUNCGElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFEFUNCGElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFEFUNCGElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFEFUNCGElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFEFUNCGElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFEFUNCGElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFEFUNCGElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFEFUNCGElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFEFUNCGElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
