// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package Shoelace RadioButton is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

type SLRADIOBUTTONElement struct {
	*Element
}

// Create a new SLRADIOBUTTONElement element.
// This will create a new element with the tag
// "sl-radio-button" during rendering.
func SL_RADIOBUTTON(children ...ElementRenderer) *SLRADIOBUTTONElement {
	e := NewElement("sl-radio-button", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SLRADIOBUTTONElement{Element: e}
}

func (e *SLRADIOBUTTONElement) Children(children ...ElementRenderer) *SLRADIOBUTTONElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SLRADIOBUTTONElement) IfChildren(condition bool, children ...ElementRenderer) *SLRADIOBUTTONElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SLRADIOBUTTONElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SLRADIOBUTTONElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SLRADIOBUTTONElement) Attr(name string, value string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SLRADIOBUTTONElement) Attrs(attrs ...string) *SLRADIOBUTTONElement {
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

func (e *SLRADIOBUTTONElement) AttrsMap(attrs map[string]string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLRADIOBUTTONElement) Text(text string) *SLRADIOBUTTONElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SLRADIOBUTTONElement) TextF(format string, args ...any) *SLRADIOBUTTONElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SLRADIOBUTTONElement) IfText(condition bool, text string) *SLRADIOBUTTONElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SLRADIOBUTTONElement) IfTextF(condition bool, format string, args ...any) *SLRADIOBUTTONElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SLRADIOBUTTONElement) Escaped(text string) *SLRADIOBUTTONElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SLRADIOBUTTONElement) IfEscaped(condition bool, text string) *SLRADIOBUTTONElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SLRADIOBUTTONElement) EscapedF(format string, args ...any) *SLRADIOBUTTONElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SLRADIOBUTTONElement) IfEscapedF(condition bool, format string, args ...any) *SLRADIOBUTTONElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SLRADIOBUTTONElement) CustomData(key, value string) *SLRADIOBUTTONElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SLRADIOBUTTONElement) IfCustomData(condition bool, key, value string) *SLRADIOBUTTONElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SLRADIOBUTTONElement) CustomDataF(key, format string, args ...any) *SLRADIOBUTTONElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SLRADIOBUTTONElement) IfCustomDataF(condition bool, key, format string, args ...any) *SLRADIOBUTTONElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SLRADIOBUTTONElement) CustomDataRemove(key string) *SLRADIOBUTTONElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

func (e *SLRADIOBUTTONElement) SIZE(c SLRadioButtonSizeChoice) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("size", string(c))
	return e
}

type SLRadioButtonSizeChoice string

const (
	// Small
	SLRadioButtonSize_small SLRadioButtonSizeChoice = "small"
	// Medium
	SLRadioButtonSize_medium SLRadioButtonSizeChoice = "medium"
	// Large
	SLRadioButtonSize_large SLRadioButtonSizeChoice = "large"
)

// Remove the attribute SIZE from the element.
func (e *SLRADIOBUTTONElement) SIZERemove(c SLRadioButtonSizeChoice) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("size")
	return e
}

func (e *SLRADIOBUTTONElement) VALUE(s string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("value", s)
	return e
}

func (e *SLRADIOBUTTONElement) IfVALUE(condition bool, s string) *SLRADIOBUTTONElement {
	if condition {
		e.VALUE(s)
	}
	return e
}

// Remove the attribute VALUE from the element.
func (e *SLRADIOBUTTONElement) VALUERemove(s string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("value")
	return e
}

func (e *SLRADIOBUTTONElement) DISABLED() *SLRADIOBUTTONElement {
	e.DISABLEDSet(true)
	return e
}

func (e *SLRADIOBUTTONElement) IfDISABLED(condition bool) *SLRADIOBUTTONElement {
	if condition {
		e.DISABLEDSet(true)
	}
	return e
}

// Set the attribute DISABLED to the value b explicitly.
func (e *SLRADIOBUTTONElement) DISABLEDSet(b bool) *SLRADIOBUTTONElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("disabled", b)
	return e
}

func (e *SLRADIOBUTTONElement) IfSetDISABLED(condition bool, b bool) *SLRADIOBUTTONElement {
	if condition {
		e.DISABLEDSet(b)
	}
	return e
}

// Remove the attribute DISABLED from the element.
func (e *SLRADIOBUTTONElement) DISABLEDRemove(b bool) *SLRADIOBUTTONElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("disabled")
	return e
}

func (e *SLRADIOBUTTONElement) PILL() *SLRADIOBUTTONElement {
	e.PILLSet(true)
	return e
}

func (e *SLRADIOBUTTONElement) IfPILL(condition bool) *SLRADIOBUTTONElement {
	if condition {
		e.PILLSet(true)
	}
	return e
}

// Set the attribute PILL to the value b explicitly.
func (e *SLRADIOBUTTONElement) PILLSet(b bool) *SLRADIOBUTTONElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("pill", b)
	return e
}

func (e *SLRADIOBUTTONElement) IfSetPILL(condition bool, b bool) *SLRADIOBUTTONElement {
	if condition {
		e.PILLSet(b)
	}
	return e
}

// Remove the attribute PILL from the element.
func (e *SLRADIOBUTTONElement) PILLRemove(b bool) *SLRADIOBUTTONElement {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("pill")
	return e
}

// Merges the store with the given object

func (e *SLRADIOBUTTONElement) DATASTAR_MERGE_STORE(v any) *SLRADIOBUTTONElement {
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

func (e *SLRADIOBUTTONElement) DATASTAR_REF(expression string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_REF(condition bool, expression string) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_REFRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SLRADIOBUTTONElement) DATASTAR_BIND(key string, expression string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_BINDRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SLRADIOBUTTONElement) DATASTAR_MODEL(expression string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_MODEL(condition bool, expression string) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_MODELRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SLRADIOBUTTONElement) DATASTAR_TEXT(expression string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_TEXT(condition bool, expression string) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_TEXTRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SLRadioButtonDataOnMod customDataKeyModifier

// Debounces the event handler
func SLRadioButtonDataOnModDebounce(
	d time.Duration,
) SLRadioButtonDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SLRadioButtonDataOnModThrottle(
	d time.Duration,
) SLRadioButtonDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SLRADIOBUTTONElement) DATASTAR_ON(key string, expression string, modifiers ...SLRadioButtonDataOnMod) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SLRadioButtonDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SLRadioButtonDataOnMod) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_ONRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SLRADIOBUTTONElement) DATASTAR_FOCUSSet(b bool) *SLRADIOBUTTONElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLRADIOBUTTONElement) DATASTAR_FOCUS() *SLRADIOBUTTONElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SLRADIOBUTTONElement) DATASTAR_HEADER(key string, expression string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_HEADERRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SLRADIOBUTTONElement) DATASTAR_FETCH_URL(expression string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_FETCH_URLRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SLRADIOBUTTONElement) DATASTAR_FETCH_INDICATOR(expression string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_FETCH_INDICATORRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SLRADIOBUTTONElement) DATASTAR_SHOWSet(b bool) *SLRADIOBUTTONElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLRADIOBUTTONElement) DATASTAR_SHOW() *SLRADIOBUTTONElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SLRADIOBUTTONElement) DATASTAR_INTERSECTS(expression string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_INTERSECTSRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SLRADIOBUTTONElement) DATASTAR_TELEPORTSet(b bool) *SLRADIOBUTTONElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLRADIOBUTTONElement) DATASTAR_TELEPORT() *SLRADIOBUTTONElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SLRADIOBUTTONElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SLRADIOBUTTONElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLRADIOBUTTONElement) DATASTAR_SCROLL_INTO_VIEW() *SLRADIOBUTTONElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SLRADIOBUTTONElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLRADIOBUTTONElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SLRADIOBUTTONElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SLRADIOBUTTONElement) DATASTAR_VIEW_TRANSITIONRemove() *SLRADIOBUTTONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
