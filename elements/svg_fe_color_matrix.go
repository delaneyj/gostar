// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feColorMatrix is generated from configuration file.
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

// The <feColorMatrix> SVG filter element changes colors based on a transformation
// matrix
// Every pixel's color value (represented by an [R,G,B,A] vector) is matrix
// multiplied to create a new color.
type SVGFECOLORMATRIXElement struct {
	*Element
}

// Create a new SVGFECOLORMATRIXElement element.
// This will create a new element with the tag
// "feColorMatrix" during rendering.
func SVG_FECOLORMATRIX(children ...ElementRenderer) *SVGFECOLORMATRIXElement {
	e := NewElement("feColorMatrix", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFECOLORMATRIXElement{Element: e}
}

func (e *SVGFECOLORMATRIXElement) Children(children ...ElementRenderer) *SVGFECOLORMATRIXElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) Attr(name string, value string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFECOLORMATRIXElement) Attrs(attrs ...string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) AttrsMap(attrs map[string]string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) Text(text string) *SVGFECOLORMATRIXElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFECOLORMATRIXElement) TextF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfText(condition bool, text string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) IfTextF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) Escaped(text string) *SVGFECOLORMATRIXElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFECOLORMATRIXElement) IfEscaped(condition bool, text string) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) EscapedF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfEscapedF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) CustomData(key, value string) *SVGFECOLORMATRIXElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfCustomData(condition bool, key, value string) *SVGFECOLORMATRIXElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) CustomDataF(key, format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) CustomDataRemove(key string) *SVGFECOLORMATRIXElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The input for this filter.
func (e *SVGFECOLORMATRIXElement) IN(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

func (e *SVGFECOLORMATRIXElement) INF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.IN(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfIN(condition bool, s string) *SVGFECOLORMATRIXElement {
	if condition {
		e.IN(s)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) IfINF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.IN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IN from the element.
func (e *SVGFECOLORMATRIXElement) INRemove(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in")
	return e
}

func (e *SVGFECOLORMATRIXElement) INRemoveF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.INRemove(fmt.Sprintf(format, args...))
}

// The type of matrix operation.
func (e *SVGFECOLORMATRIXElement) TYPE(c SVGFeColorMatrixTypeChoice) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeColorMatrixTypeChoice string

const (
	// The type of matrix operation.
	SVGFeColorMatrixType_matrix SVGFeColorMatrixTypeChoice = "matrix"
	// The type of matrix operation.
	SVGFeColorMatrixType_saturate SVGFeColorMatrixTypeChoice = "saturate"
	// The type of matrix operation.
	SVGFeColorMatrixType_hueRotate SVGFeColorMatrixTypeChoice = "hueRotate"
	// The type of matrix operation.
	SVGFeColorMatrixType_luminanceToAlpha SVGFeColorMatrixTypeChoice = "luminanceToAlpha"
)

// Remove the attribute TYPE from the element.
func (e *SVGFECOLORMATRIXElement) TYPERemove(c SVGFeColorMatrixTypeChoice) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

// The list of one or more numbers that represent the matrix.
func (e *SVGFECOLORMATRIXElement) VALUES(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("values", s)
	return e
}

func (e *SVGFECOLORMATRIXElement) VALUESF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfVALUES(condition bool, s string) *SVGFECOLORMATRIXElement {
	if condition {
		e.VALUES(s)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) IfVALUESF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VALUES from the element.
func (e *SVGFECOLORMATRIXElement) VALUESRemove(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("values")
	return e
}

func (e *SVGFECOLORMATRIXElement) VALUESRemoveF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.VALUESRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGFECOLORMATRIXElement) ID(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFECOLORMATRIXElement) IDF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfID(condition bool, s string) *SVGFECOLORMATRIXElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) IfIDF(condition bool, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFECOLORMATRIXElement) IDRemove(s string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGFECOLORMATRIXElement) IDRemoveF(format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFECOLORMATRIXElement) CLASS(s ...string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) IfCLASS(condition bool, s ...string) *SVGFECOLORMATRIXElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFECOLORMATRIXElement) CLASSRemove(s ...string) *SVGFECOLORMATRIXElement {
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
func (e *SVGFECOLORMATRIXElement) STYLEF(k string, format string, args ...any) *SVGFECOLORMATRIXElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFECOLORMATRIXElement) IfSTYLE(condition bool, k string, v string) *SVGFECOLORMATRIXElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFECOLORMATRIXElement) STYLE(k string, v string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFECOLORMATRIXElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFECOLORMATRIXElement) STYLEMap(m map[string]string) *SVGFECOLORMATRIXElement {
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
func (e *SVGFECOLORMATRIXElement) STYLEPairs(pairs ...string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFECOLORMATRIXElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFECOLORMATRIXElement) STYLERemove(keys ...string) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) DATASTAR_MERGE_STORE(v any) *SVGFECOLORMATRIXElement {
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

func (e *SVGFECOLORMATRIXElement) DATASTAR_REF(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfDATASTAR_REF(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFECOLORMATRIXElement) DATASTAR_REFRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFECOLORMATRIXElement) DATASTAR_BIND(key string, expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFECOLORMATRIXElement) DATASTAR_BINDRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFECOLORMATRIXElement) DATASTAR_MODEL(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFECOLORMATRIXElement) DATASTAR_MODELRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFECOLORMATRIXElement) DATASTAR_TEXT(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFECOLORMATRIXElement) DATASTAR_TEXTRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeColorMatrixOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeColorMatrixOnModDebounce(
	d time.Duration,
) SVGFeColorMatrixOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeColorMatrixOnModThrottle(
	d time.Duration,
) SVGFeColorMatrixOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFECOLORMATRIXElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeColorMatrixOnMod) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeColorMatrixOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeColorMatrixOnMod) *SVGFECOLORMATRIXElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFECOLORMATRIXElement) DATASTAR_ONRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFECOLORMATRIXElement) DATASTAR_FOCUSSet(b bool) *SVGFECOLORMATRIXElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFECOLORMATRIXElement) DATASTAR_FOCUS() *SVGFECOLORMATRIXElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFECOLORMATRIXElement) DATASTAR_HEADER(key string, expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFECOLORMATRIXElement) DATASTAR_HEADERRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFECOLORMATRIXElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFECOLORMATRIXElement) DATASTAR_FETCH_INDICATORRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFECOLORMATRIXElement) DATASTAR_SHOWSet(b bool) *SVGFECOLORMATRIXElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFECOLORMATRIXElement) DATASTAR_SHOW() *SVGFECOLORMATRIXElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFECOLORMATRIXElement) DATASTAR_INTERSECTS(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFECOLORMATRIXElement) DATASTAR_INTERSECTSRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFECOLORMATRIXElement) DATASTAR_TELEPORTSet(b bool) *SVGFECOLORMATRIXElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFECOLORMATRIXElement) DATASTAR_TELEPORT() *SVGFECOLORMATRIXElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFECOLORMATRIXElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFECOLORMATRIXElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFECOLORMATRIXElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFECOLORMATRIXElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFECOLORMATRIXElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFECOLORMATRIXElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFECOLORMATRIXElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFECOLORMATRIXElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFECOLORMATRIXElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
