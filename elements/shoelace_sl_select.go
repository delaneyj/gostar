// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package Shoelace Select is generated from configuration file.
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

type SLSELECTElement struct {
	*Element
}

// Create a new SLSELECTElement element.
// This will create a new element with the tag
// "sl-select" during rendering.
func SL_SELECT(children ...ElementRenderer) *SLSELECTElement {
	e := NewElement("sl-select", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SLSELECTElement{Element: e}
}

func (e *SLSELECTElement) Children(children ...ElementRenderer) *SLSELECTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SLSELECTElement) IfChildren(condition bool, children ...ElementRenderer) *SLSELECTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SLSELECTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SLSELECTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SLSELECTElement) Attr(name string, value string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SLSELECTElement) Attrs(attrs ...string) *SLSELECTElement {
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

func (e *SLSELECTElement) AttrsMap(attrs map[string]string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLSELECTElement) Text(text string) *SLSELECTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SLSELECTElement) TextF(format string, args ...any) *SLSELECTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) IfText(condition bool, text string) *SLSELECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SLSELECTElement) IfTextF(condition bool, format string, args ...any) *SLSELECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SLSELECTElement) Escaped(text string) *SLSELECTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SLSELECTElement) IfEscaped(condition bool, text string) *SLSELECTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SLSELECTElement) EscapedF(format string, args ...any) *SLSELECTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) IfEscapedF(condition bool, format string, args ...any) *SLSELECTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SLSELECTElement) CustomData(key, value string) *SLSELECTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SLSELECTElement) IfCustomData(condition bool, key, value string) *SLSELECTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SLSELECTElement) CustomDataF(key, format string, args ...any) *SLSELECTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SLSELECTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SLSELECTElement) CustomDataRemove(key string) *SLSELECTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

func (e *SLSELECTElement) SIZE(c SLSelectSizeChoice) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("size", string(c))
	return e
}

type SLSelectSizeChoice string

const (
	// Small
	SLSelectSize_small SLSelectSizeChoice = "small"
	// Medium
	SLSelectSize_medium SLSelectSizeChoice = "medium"
	// Large
	SLSelectSize_large SLSelectSizeChoice = "large"
)

// Remove the attribute SIZE from the element.
func (e *SLSELECTElement) SIZERemove(c SLSelectSizeChoice) *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("size")
	return e
}

func (e *SLSELECTElement) HOIST() *SLSELECTElement {
	e.HOISTSet(true)
	return e
}

func (e *SLSELECTElement) IfHOIST(condition bool) *SLSELECTElement {
	if condition {
		e.HOISTSet(true)
	}
	return e
}

// Set the attribute HOIST to the value b explicitly.
func (e *SLSELECTElement) HOISTSet(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("hoist", b)
	return e
}

func (e *SLSELECTElement) IfSetHOIST(condition bool, b bool) *SLSELECTElement {
	if condition {
		e.HOISTSet(b)
	}
	return e
}

// Remove the attribute HOIST from the element.
func (e *SLSELECTElement) HOISTRemove(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("hoist")
	return e
}

func (e *SLSELECTElement) LABEL(s string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("label", s)
	return e
}

func (e *SLSELECTElement) LABELF(format string, args ...any) *SLSELECTElement {
	return e.LABEL(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) IfLABEL(condition bool, s string) *SLSELECTElement {
	if condition {
		e.LABEL(s)
	}
	return e
}

func (e *SLSELECTElement) IfLABELF(condition bool, format string, args ...any) *SLSELECTElement {
	if condition {
		e.LABEL(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute LABEL from the element.
func (e *SLSELECTElement) LABELRemove(s string) *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("label")
	return e
}

func (e *SLSELECTElement) LABELRemoveF(format string, args ...any) *SLSELECTElement {
	return e.LABELRemove(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) HELP_TEXT(s string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("help-text", s)
	return e
}

func (e *SLSELECTElement) HELP_TEXTF(format string, args ...any) *SLSELECTElement {
	return e.HELP_TEXT(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) IfHELP_TEXT(condition bool, s string) *SLSELECTElement {
	if condition {
		e.HELP_TEXT(s)
	}
	return e
}

func (e *SLSELECTElement) IfHELP_TEXTF(condition bool, format string, args ...any) *SLSELECTElement {
	if condition {
		e.HELP_TEXT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HELP_TEXT from the element.
func (e *SLSELECTElement) HELP_TEXTRemove(s string) *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("help-text")
	return e
}

func (e *SLSELECTElement) HELP_TEXTRemoveF(format string, args ...any) *SLSELECTElement {
	return e.HELP_TEXTRemove(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) PLACEHOLDER(s string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("placeholder", s)
	return e
}

func (e *SLSELECTElement) PLACEHOLDERF(format string, args ...any) *SLSELECTElement {
	return e.PLACEHOLDER(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) IfPLACEHOLDER(condition bool, s string) *SLSELECTElement {
	if condition {
		e.PLACEHOLDER(s)
	}
	return e
}

func (e *SLSELECTElement) IfPLACEHOLDERF(condition bool, format string, args ...any) *SLSELECTElement {
	if condition {
		e.PLACEHOLDER(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute PLACEHOLDER from the element.
func (e *SLSELECTElement) PLACEHOLDERRemove(s string) *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("placeholder")
	return e
}

func (e *SLSELECTElement) PLACEHOLDERRemoveF(format string, args ...any) *SLSELECTElement {
	return e.PLACEHOLDERRemove(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) CLEARABLE() *SLSELECTElement {
	e.CLEARABLESet(true)
	return e
}

func (e *SLSELECTElement) IfCLEARABLE(condition bool) *SLSELECTElement {
	if condition {
		e.CLEARABLESet(true)
	}
	return e
}

// Set the attribute CLEARABLE to the value b explicitly.
func (e *SLSELECTElement) CLEARABLESet(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("clearable", b)
	return e
}

func (e *SLSELECTElement) IfSetCLEARABLE(condition bool, b bool) *SLSELECTElement {
	if condition {
		e.CLEARABLESet(b)
	}
	return e
}

// Remove the attribute CLEARABLE from the element.
func (e *SLSELECTElement) CLEARABLERemove(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("clearable")
	return e
}

func (e *SLSELECTElement) FILLED() *SLSELECTElement {
	e.FILLEDSet(true)
	return e
}

func (e *SLSELECTElement) IfFILLED(condition bool) *SLSELECTElement {
	if condition {
		e.FILLEDSet(true)
	}
	return e
}

// Set the attribute FILLED to the value b explicitly.
func (e *SLSELECTElement) FILLEDSet(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("filled", b)
	return e
}

func (e *SLSELECTElement) IfSetFILLED(condition bool, b bool) *SLSELECTElement {
	if condition {
		e.FILLEDSet(b)
	}
	return e
}

// Remove the attribute FILLED from the element.
func (e *SLSELECTElement) FILLEDRemove(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("filled")
	return e
}

func (e *SLSELECTElement) DISABLED() *SLSELECTElement {
	e.DISABLEDSet(true)
	return e
}

func (e *SLSELECTElement) IfDISABLED(condition bool) *SLSELECTElement {
	if condition {
		e.DISABLEDSet(true)
	}
	return e
}

// Set the attribute DISABLED to the value b explicitly.
func (e *SLSELECTElement) DISABLEDSet(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("disabled", b)
	return e
}

func (e *SLSELECTElement) IfSetDISABLED(condition bool, b bool) *SLSELECTElement {
	if condition {
		e.DISABLEDSet(b)
	}
	return e
}

// Remove the attribute DISABLED from the element.
func (e *SLSELECTElement) DISABLEDRemove(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("disabled")
	return e
}

func (e *SLSELECTElement) MULTIPLE() *SLSELECTElement {
	e.MULTIPLESet(true)
	return e
}

func (e *SLSELECTElement) IfMULTIPLE(condition bool) *SLSELECTElement {
	if condition {
		e.MULTIPLESet(true)
	}
	return e
}

// Set the attribute MULTIPLE to the value b explicitly.
func (e *SLSELECTElement) MULTIPLESet(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("multiple", b)
	return e
}

func (e *SLSELECTElement) IfSetMULTIPLE(condition bool, b bool) *SLSELECTElement {
	if condition {
		e.MULTIPLESet(b)
	}
	return e
}

// Remove the attribute MULTIPLE from the element.
func (e *SLSELECTElement) MULTIPLERemove(b bool) *SLSELECTElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("multiple")
	return e
}

func (e *SLSELECTElement) VALUE(s string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("value", s)
	return e
}

func (e *SLSELECTElement) VALUEF(format string, args ...any) *SLSELECTElement {
	return e.VALUE(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) IfVALUE(condition bool, s string) *SLSELECTElement {
	if condition {
		e.VALUE(s)
	}
	return e
}

func (e *SLSELECTElement) IfVALUEF(condition bool, format string, args ...any) *SLSELECTElement {
	if condition {
		e.VALUE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VALUE from the element.
func (e *SLSELECTElement) VALUERemove(s string) *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("value")
	return e
}

func (e *SLSELECTElement) VALUERemoveF(format string, args ...any) *SLSELECTElement {
	return e.VALUERemove(fmt.Sprintf(format, args...))
}

func (e *SLSELECTElement) PLACEMENT(c SLSelectPlacementChoice) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("placement", string(c))
	return e
}

type SLSelectPlacementChoice string

const (
	// Bottom
	SLSelectPlacement_bottom SLSelectPlacementChoice = "bottom"
	// Top
	SLSelectPlacement_top SLSelectPlacementChoice = "top"
)

// Remove the attribute PLACEMENT from the element.
func (e *SLSELECTElement) PLACEMENTRemove(c SLSelectPlacementChoice) *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("placement")
	return e
}

// Merges the store with the given object

func (e *SLSELECTElement) DATASTAR_MERGE_STORE(v any) *SLSELECTElement {
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

func (e *SLSELECTElement) DATASTAR_REF(expression string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_REF(condition bool, expression string) *SLSELECTElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SLSELECTElement) DATASTAR_REFRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SLSELECTElement) DATASTAR_BIND(key string, expression string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SLSELECTElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SLSELECTElement) DATASTAR_BINDRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SLSELECTElement) DATASTAR_MODEL(expression string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_MODEL(condition bool, expression string) *SLSELECTElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SLSELECTElement) DATASTAR_MODELRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SLSELECTElement) DATASTAR_TEXT(expression string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_TEXT(condition bool, expression string) *SLSELECTElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SLSELECTElement) DATASTAR_TEXTRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SLSelectOnMod customDataKeyModifier

// Debounces the event handler
func SLSelectOnModDebounce(
	d time.Duration,
) SLSelectOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SLSelectOnModThrottle(
	d time.Duration,
) SLSelectOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SLSELECTElement) DATASTAR_ON(key string, expression string, modifiers ...SLSelectOnMod) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SLSelectOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SLSelectOnMod) *SLSELECTElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SLSELECTElement) DATASTAR_ONRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SLSELECTElement) DATASTAR_FOCUSSet(b bool) *SLSELECTElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLSELECTElement) DATASTAR_FOCUS() *SLSELECTElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SLSELECTElement) DATASTAR_HEADER(key string, expression string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SLSELECTElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SLSELECTElement) DATASTAR_HEADERRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SLSELECTElement) DATASTAR_FETCH_URL(expression string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SLSELECTElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SLSELECTElement) DATASTAR_FETCH_URLRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SLSELECTElement) DATASTAR_FETCH_INDICATOR(expression string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SLSELECTElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SLSELECTElement) DATASTAR_FETCH_INDICATORRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SLSELECTElement) DATASTAR_SHOWSet(b bool) *SLSELECTElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLSELECTElement) DATASTAR_SHOW() *SLSELECTElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SLSELECTElement) DATASTAR_INTERSECTS(expression string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SLSELECTElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SLSELECTElement) DATASTAR_INTERSECTSRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SLSELECTElement) DATASTAR_TELEPORTSet(b bool) *SLSELECTElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLSELECTElement) DATASTAR_TELEPORT() *SLSELECTElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SLSELECTElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SLSELECTElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLSELECTElement) DATASTAR_SCROLL_INTO_VIEW() *SLSELECTElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SLSELECTElement) DATASTAR_VIEW_TRANSITION(expression string) *SLSELECTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLSELECTElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SLSELECTElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SLSELECTElement) DATASTAR_VIEW_TRANSITIONRemove() *SLSELECTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
