// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package Shoelace Card is generated from configuration file.
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

type SLCARDElement struct {
	*Element
}

// Create a new SLCARDElement element.
// This will create a new element with the tag
// "sl-card" during rendering.
func SL_CARD(children ...ElementRenderer) *SLCARDElement {
	e := NewElement("sl-card", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SLCARDElement{Element: e}
}

func (e *SLCARDElement) Children(children ...ElementRenderer) *SLCARDElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SLCARDElement) IfChildren(condition bool, children ...ElementRenderer) *SLCARDElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SLCARDElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SLCARDElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SLCARDElement) Attr(name string, value string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SLCARDElement) Attrs(attrs ...string) *SLCARDElement {
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

func (e *SLCARDElement) AttrsMap(attrs map[string]string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLCARDElement) Text(text string) *SLCARDElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SLCARDElement) TextF(format string, args ...any) *SLCARDElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SLCARDElement) IfText(condition bool, text string) *SLCARDElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SLCARDElement) IfTextF(condition bool, format string, args ...any) *SLCARDElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SLCARDElement) Escaped(text string) *SLCARDElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SLCARDElement) IfEscaped(condition bool, text string) *SLCARDElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SLCARDElement) EscapedF(format string, args ...any) *SLCARDElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SLCARDElement) IfEscapedF(condition bool, format string, args ...any) *SLCARDElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SLCARDElement) CustomData(key, value string) *SLCARDElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SLCARDElement) IfCustomData(condition bool, key, value string) *SLCARDElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SLCARDElement) CustomDataF(key, format string, args ...any) *SLCARDElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SLCARDElement) IfCustomDataF(condition bool, key, format string, args ...any) *SLCARDElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SLCARDElement) CustomDataRemove(key string) *SLCARDElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Merges the singleton store with the given object

func (e *SLCARDElement) DATASTA_STORE(v any) *SLCARDElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("store", html.EscapeString(string(b)))
	return e
}

// Sets the reference of the element

func (e *SLCARDElement) DATASTAR_REF(expression string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLCARDElement) IfDATASTAR_REF(condition bool, expression string) *SLCARDElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SLCARDElement) DATASTAR_REFRemove() *SLCARDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SLCARDElement) DATASTAR_BIND(key string, expression string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLCARDElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SLCARDElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SLCARDElement) DATASTAR_BINDRemove() *SLCARDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SLCARDElement) DATASTAR_MODEL(expression string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLCARDElement) IfDATASTAR_MODEL(condition bool, expression string) *SLCARDElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SLCARDElement) DATASTAR_MODELRemove() *SLCARDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SLCARDElement) DATASTAR_TEXT(expression string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLCARDElement) IfDATASTAR_TEXT(condition bool, expression string) *SLCARDElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SLCARDElement) DATASTAR_TEXTRemove() *SLCARDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SLCardOnMod customDataKeyModifier

// Debounces the event handler
func SLCardOnModDebounce(
	d time.Duration,
) SLCardOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SLCardOnModThrottle(
	d time.Duration,
) SLCardOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SLCARDElement) DATASTAR_ON(key string, expression string, modifiers ...SLCardOnMod) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SLCardOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLCARDElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SLCardOnMod) *SLCARDElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SLCARDElement) DATASTAR_ONRemove() *SLCARDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SLCARDElement) DATASTAR_FOCUSSet(b bool) *SLCARDElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLCARDElement) DATASTAR_FOCUS() *SLCARDElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SLCARDElement) DATASTAR_HEADER(key string, expression string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLCARDElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SLCARDElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SLCARDElement) DATASTAR_HEADERRemove() *SLCARDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SLCARDElement) DATASTAR_FETCH_INDICATOR(expression string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLCARDElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SLCARDElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SLCARDElement) DATASTAR_FETCH_INDICATORRemove() *SLCARDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SLCARDElement) DATASTAR_SHOWSet(b bool) *SLCARDElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLCARDElement) DATASTAR_SHOW() *SLCARDElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SLCARDElement) DATASTAR_INTERSECTS(expression string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLCARDElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SLCARDElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SLCARDElement) DATASTAR_INTERSECTSRemove() *SLCARDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SLCARDElement) DATASTAR_TELEPORTSet(b bool) *SLCARDElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLCARDElement) DATASTAR_TELEPORT() *SLCARDElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SLCARDElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SLCARDElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLCARDElement) DATASTAR_SCROLL_INTO_VIEW() *SLCARDElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SLCARDElement) DATASTAR_VIEW_TRANSITION(expression string) *SLCARDElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLCARDElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SLCARDElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SLCARDElement) DATASTAR_VIEW_TRANSITIONRemove() *SLCARDElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
