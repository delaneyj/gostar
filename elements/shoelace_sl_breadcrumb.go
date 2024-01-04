// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package Shoelace Breadcrumb is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

type SLBREADCRUMBElement struct {
	*Element
}

// Create a new SLBREADCRUMBElement element.
// This will create a new element with the tag
// "sl-breadcrumb" during rendering.
func SL_BREADCRUMB(children ...ElementRenderer) *SLBREADCRUMBElement {
	e := NewElement("sl-breadcrumb", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SLBREADCRUMBElement{Element: e}
}

func (e *SLBREADCRUMBElement) Children(children ...ElementRenderer) *SLBREADCRUMBElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SLBREADCRUMBElement) IfChildren(condition bool, children ...ElementRenderer) *SLBREADCRUMBElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SLBREADCRUMBElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SLBREADCRUMBElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SLBREADCRUMBElement) Attr(name string, value string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SLBREADCRUMBElement) Attrs(attrs ...string) *SLBREADCRUMBElement {
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

func (e *SLBREADCRUMBElement) AttrsMap(attrs map[string]string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLBREADCRUMBElement) Text(text string) *SLBREADCRUMBElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SLBREADCRUMBElement) TextF(format string, args ...any) *SLBREADCRUMBElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SLBREADCRUMBElement) IfText(condition bool, text string) *SLBREADCRUMBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SLBREADCRUMBElement) IfTextF(condition bool, format string, args ...any) *SLBREADCRUMBElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SLBREADCRUMBElement) Escaped(text string) *SLBREADCRUMBElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SLBREADCRUMBElement) IfEscaped(condition bool, text string) *SLBREADCRUMBElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SLBREADCRUMBElement) EscapedF(format string, args ...any) *SLBREADCRUMBElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SLBREADCRUMBElement) IfEscapedF(condition bool, format string, args ...any) *SLBREADCRUMBElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SLBREADCRUMBElement) CustomData(key, value string) *SLBREADCRUMBElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SLBREADCRUMBElement) IfCustomData(condition bool, key, value string) *SLBREADCRUMBElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SLBREADCRUMBElement) CustomDataF(key, format string, args ...any) *SLBREADCRUMBElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SLBREADCRUMBElement) IfCustomDataF(condition bool, key, format string, args ...any) *SLBREADCRUMBElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SLBREADCRUMBElement) CustomDataRemove(key string) *SLBREADCRUMBElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Merges the store with the given object

func (e *SLBREADCRUMBElement) DATASTAR_MERGE_STORE(v any) *SLBREADCRUMBElement {
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

func (e *SLBREADCRUMBElement) DATASTAR_REF(expression string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_REF(condition bool, expression string) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SLBREADCRUMBElement) DATASTAR_REFRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SLBREADCRUMBElement) DATASTAR_BIND(key string, expression string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SLBREADCRUMBElement) DATASTAR_BINDRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SLBREADCRUMBElement) DATASTAR_MODEL(expression string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_MODEL(condition bool, expression string) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SLBREADCRUMBElement) DATASTAR_MODELRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SLBREADCRUMBElement) DATASTAR_TEXT(expression string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_TEXT(condition bool, expression string) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SLBREADCRUMBElement) DATASTAR_TEXTRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SLBreadcrumbOnMod customDataKeyModifier

// Debounces the event handler
func SLBreadcrumbOnModDebounce(
	d time.Duration,
) SLBreadcrumbOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SLBreadcrumbOnModThrottle(
	d time.Duration,
) SLBreadcrumbOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SLBREADCRUMBElement) DATASTAR_ON(key string, expression string, modifiers ...SLBreadcrumbOnMod) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SLBreadcrumbOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SLBreadcrumbOnMod) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SLBREADCRUMBElement) DATASTAR_ONRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SLBREADCRUMBElement) DATASTAR_FOCUSSet(b bool) *SLBREADCRUMBElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLBREADCRUMBElement) DATASTAR_FOCUS() *SLBREADCRUMBElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SLBREADCRUMBElement) DATASTAR_HEADER(key string, expression string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SLBREADCRUMBElement) DATASTAR_HEADERRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SLBREADCRUMBElement) DATASTAR_FETCH_URL(expression string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SLBREADCRUMBElement) DATASTAR_FETCH_URLRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SLBREADCRUMBElement) DATASTAR_FETCH_INDICATOR(expression string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SLBREADCRUMBElement) DATASTAR_FETCH_INDICATORRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SLBREADCRUMBElement) DATASTAR_SHOWSet(b bool) *SLBREADCRUMBElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLBREADCRUMBElement) DATASTAR_SHOW() *SLBREADCRUMBElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SLBREADCRUMBElement) DATASTAR_INTERSECTS(expression string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SLBREADCRUMBElement) DATASTAR_INTERSECTSRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SLBREADCRUMBElement) DATASTAR_TELEPORTSet(b bool) *SLBREADCRUMBElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLBREADCRUMBElement) DATASTAR_TELEPORT() *SLBREADCRUMBElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SLBREADCRUMBElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SLBREADCRUMBElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLBREADCRUMBElement) DATASTAR_SCROLL_INTO_VIEW() *SLBREADCRUMBElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SLBREADCRUMBElement) DATASTAR_VIEW_TRANSITION(expression string) *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLBREADCRUMBElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SLBREADCRUMBElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SLBREADCRUMBElement) DATASTAR_VIEW_TRANSITIONRemove() *SLBREADCRUMBElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
