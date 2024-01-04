// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg linearGradient is generated from configuration file.
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

// The <linearGradient> SVG element lets authors define linear gradients to fill
// or stroke graphical elements.
type SVGLINEARGRADIENTElement struct {
	*Element
}

// Create a new SVGLINEARGRADIENTElement element.
// This will create a new element with the tag
// "linearGradient" during rendering.
func SVG_LINEARGRADIENT(children ...ElementRenderer) *SVGLINEARGRADIENTElement {
	e := NewElement("linearGradient", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGLINEARGRADIENTElement{Element: e}
}

func (e *SVGLINEARGRADIENTElement) Children(children ...ElementRenderer) *SVGLINEARGRADIENTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) Attr(name string, value string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGLINEARGRADIENTElement) Attrs(attrs ...string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) AttrsMap(attrs map[string]string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) Text(text string) *SVGLINEARGRADIENTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGLINEARGRADIENTElement) TextF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfText(condition bool, text string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) IfTextF(condition bool, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) Escaped(text string) *SVGLINEARGRADIENTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGLINEARGRADIENTElement) IfEscaped(condition bool, text string) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) EscapedF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfEscapedF(condition bool, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) CustomData(key, value string) *SVGLINEARGRADIENTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfCustomData(condition bool, key, value string) *SVGLINEARGRADIENTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) CustomDataF(key, format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) CustomDataRemove(key string) *SVGLINEARGRADIENTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The coordinate system for attributes x1, y1, x2 and y2.
func (e *SVGLINEARGRADIENTElement) GRADIENT_UNITS(c SVGLinearGradientGradientUnitsChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("gradientUnits", string(c))
	return e
}

type SVGLinearGradientGradientUnitsChoice string

const (
	// The coordinate system for attributes x1, y1, x2 and y2.
	SVGLinearGradientGradientUnits_userSpaceOnUse SVGLinearGradientGradientUnitsChoice = "userSpaceOnUse"
	// The coordinate system for attributes x1, y1, x2 and y2.
	SVGLinearGradientGradientUnits_objectBoundingBox SVGLinearGradientGradientUnitsChoice = "objectBoundingBox"
)

// Remove the attribute GRADIENT_UNITS from the element.
func (e *SVGLINEARGRADIENTElement) GRADIENT_UNITSRemove(c SVGLinearGradientGradientUnitsChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("gradientUnits")
	return e
}

// The definition of how the gradient is applied, read about <a
// href="https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/gradientTransform">gradientTransform</a>.
func (e *SVGLINEARGRADIENTElement) GRADIENT_TRANSFORM(s string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("gradientTransform", s)
	return e
}

func (e *SVGLINEARGRADIENTElement) GRADIENT_TRANSFORMF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.GRADIENT_TRANSFORM(fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfGRADIENT_TRANSFORM(condition bool, s string) *SVGLINEARGRADIENTElement {
	if condition {
		e.GRADIENT_TRANSFORM(s)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) IfGRADIENT_TRANSFORMF(condition bool, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.GRADIENT_TRANSFORM(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute GRADIENT_TRANSFORM from the element.
func (e *SVGLINEARGRADIENTElement) GRADIENT_TRANSFORMRemove(s string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("gradientTransform")
	return e
}

func (e *SVGLINEARGRADIENTElement) GRADIENT_TRANSFORMRemoveF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.GRADIENT_TRANSFORMRemove(fmt.Sprintf(format, args...))
}

// The method by which to fill a shape.
func (e *SVGLINEARGRADIENTElement) SPREAD_METHOD(c SVGLinearGradientSpreadMethodChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("spreadMethod", string(c))
	return e
}

type SVGLinearGradientSpreadMethodChoice string

const (
	// The method by which to fill a shape.
	SVGLinearGradientSpreadMethod_pad SVGLinearGradientSpreadMethodChoice = "pad"
	// The method by which to fill a shape.
	SVGLinearGradientSpreadMethod_reflect SVGLinearGradientSpreadMethodChoice = "reflect"
	// The method by which to fill a shape.
	SVGLinearGradientSpreadMethod_repeat SVGLinearGradientSpreadMethodChoice = "repeat"
)

// Remove the attribute SPREAD_METHOD from the element.
func (e *SVGLINEARGRADIENTElement) SPREAD_METHODRemove(c SVGLinearGradientSpreadMethodChoice) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("spreadMethod")
	return e
}

// The x-axis coordinate of the start of the gradient.
func (e *SVGLINEARGRADIENTElement) X_1(f float64) *SVGLINEARGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("x1", f)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfX_1(condition bool, f float64) *SVGLINEARGRADIENTElement {
	if condition {
		e.X_1(f)
	}
	return e
}

// The y-axis coordinate of the start of the gradient.
func (e *SVGLINEARGRADIENTElement) Y_1(f float64) *SVGLINEARGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("y1", f)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfY_1(condition bool, f float64) *SVGLINEARGRADIENTElement {
	if condition {
		e.Y_1(f)
	}
	return e
}

// The x-axis coordinate of the end of the gradient.
func (e *SVGLINEARGRADIENTElement) X_2(f float64) *SVGLINEARGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("x2", f)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfX_2(condition bool, f float64) *SVGLINEARGRADIENTElement {
	if condition {
		e.X_2(f)
	}
	return e
}

// The y-axis coordinate of the end of the gradient.
func (e *SVGLINEARGRADIENTElement) Y_2(f float64) *SVGLINEARGRADIENTElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("y2", f)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfY_2(condition bool, f float64) *SVGLINEARGRADIENTElement {
	if condition {
		e.Y_2(f)
	}
	return e
}

// Specifies a unique id for an element
func (e *SVGLINEARGRADIENTElement) ID(s string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGLINEARGRADIENTElement) IDF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfID(condition bool, s string) *SVGLINEARGRADIENTElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) IfIDF(condition bool, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGLINEARGRADIENTElement) IDRemove(s string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGLINEARGRADIENTElement) IDRemoveF(format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGLINEARGRADIENTElement) CLASS(s ...string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) IfCLASS(condition bool, s ...string) *SVGLINEARGRADIENTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGLINEARGRADIENTElement) CLASSRemove(s ...string) *SVGLINEARGRADIENTElement {
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
func (e *SVGLINEARGRADIENTElement) STYLEF(k string, format string, args ...any) *SVGLINEARGRADIENTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGLINEARGRADIENTElement) IfSTYLE(condition bool, k string, v string) *SVGLINEARGRADIENTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGLINEARGRADIENTElement) STYLE(k string, v string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGLINEARGRADIENTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGLINEARGRADIENTElement) STYLEMap(m map[string]string) *SVGLINEARGRADIENTElement {
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
func (e *SVGLINEARGRADIENTElement) STYLEPairs(pairs ...string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGLINEARGRADIENTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGLINEARGRADIENTElement) STYLERemove(keys ...string) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) DATASTAR_MERGE_STORE(v any) *SVGLINEARGRADIENTElement {
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

func (e *SVGLINEARGRADIENTElement) DATASTAR_REF(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_REF(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_REFRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGLINEARGRADIENTElement) DATASTAR_BIND(key string, expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_BINDRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGLINEARGRADIENTElement) DATASTAR_MODEL(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_MODELRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGLINEARGRADIENTElement) DATASTAR_TEXT(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_TEXTRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGLinearGradientOnMod customDataKeyModifier

// Debounces the event handler
func SVGLinearGradientOnModDebounce(
	d time.Duration,
) SVGLinearGradientOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGLinearGradientOnModThrottle(
	d time.Duration,
) SVGLinearGradientOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGLINEARGRADIENTElement) DATASTAR_ON(key string, expression string, modifiers ...SVGLinearGradientOnMod) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGLinearGradientOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGLinearGradientOnMod) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_ONRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGLINEARGRADIENTElement) DATASTAR_FOCUSSet(b bool) *SVGLINEARGRADIENTElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGLINEARGRADIENTElement) DATASTAR_FOCUS() *SVGLINEARGRADIENTElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGLINEARGRADIENTElement) DATASTAR_HEADER(key string, expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_HEADERRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGLINEARGRADIENTElement) DATASTAR_FETCH_URL(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_FETCH_URLRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGLINEARGRADIENTElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_FETCH_INDICATORRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGLINEARGRADIENTElement) DATASTAR_SHOWSet(b bool) *SVGLINEARGRADIENTElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGLINEARGRADIENTElement) DATASTAR_SHOW() *SVGLINEARGRADIENTElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGLINEARGRADIENTElement) DATASTAR_INTERSECTS(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_INTERSECTSRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGLINEARGRADIENTElement) DATASTAR_TELEPORTSet(b bool) *SVGLINEARGRADIENTElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGLINEARGRADIENTElement) DATASTAR_TELEPORT() *SVGLINEARGRADIENTElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGLINEARGRADIENTElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGLINEARGRADIENTElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGLINEARGRADIENTElement) DATASTAR_SCROLL_INTO_VIEW() *SVGLINEARGRADIENTElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGLINEARGRADIENTElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGLINEARGRADIENTElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGLINEARGRADIENTElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGLINEARGRADIENTElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGLINEARGRADIENTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
