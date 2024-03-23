// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package Shoelace Avatar is generated from configuration file.
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

// Avatars are used to represent a user, project, or other item.
type SLAVATARElement struct {
	*Element
}

// Create a new SLAVATARElement element.
// This will create a new element with the tag
// "sl-avatar" during rendering.
func SL_AVATAR(children ...ElementRenderer) *SLAVATARElement {
	e := NewElement("sl-avatar", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SLAVATARElement{Element: e}
}

func (e *SLAVATARElement) Children(children ...ElementRenderer) *SLAVATARElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SLAVATARElement) IfChildren(condition bool, children ...ElementRenderer) *SLAVATARElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SLAVATARElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SLAVATARElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SLAVATARElement) Attr(name string, value string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SLAVATARElement) Attrs(attrs ...string) *SLAVATARElement {
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

func (e *SLAVATARElement) AttrsMap(attrs map[string]string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SLAVATARElement) Text(text string) *SLAVATARElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SLAVATARElement) TextF(format string, args ...any) *SLAVATARElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SLAVATARElement) IfText(condition bool, text string) *SLAVATARElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SLAVATARElement) IfTextF(condition bool, format string, args ...any) *SLAVATARElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SLAVATARElement) Escaped(text string) *SLAVATARElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SLAVATARElement) IfEscaped(condition bool, text string) *SLAVATARElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SLAVATARElement) EscapedF(format string, args ...any) *SLAVATARElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SLAVATARElement) IfEscapedF(condition bool, format string, args ...any) *SLAVATARElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SLAVATARElement) CustomData(key, value string) *SLAVATARElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SLAVATARElement) IfCustomData(condition bool, key, value string) *SLAVATARElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SLAVATARElement) CustomDataF(key, format string, args ...any) *SLAVATARElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SLAVATARElement) IfCustomDataF(condition bool, key, format string, args ...any) *SLAVATARElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SLAVATARElement) CustomDataRemove(key string) *SLAVATARElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

func (e *SLAVATARElement) IMAGE(s string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("image", s)
	return e
}

func (e *SLAVATARElement) IMAGEF(format string, args ...any) *SLAVATARElement {
	return e.IMAGE(fmt.Sprintf(format, args...))
}

func (e *SLAVATARElement) IfIMAGE(condition bool, s string) *SLAVATARElement {
	if condition {
		e.IMAGE(s)
	}
	return e
}

func (e *SLAVATARElement) IfIMAGEF(condition bool, format string, args ...any) *SLAVATARElement {
	if condition {
		e.IMAGE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute IMAGE from the element.
func (e *SLAVATARElement) IMAGERemove(s string) *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("image")
	return e
}

func (e *SLAVATARElement) IMAGERemoveF(format string, args ...any) *SLAVATARElement {
	return e.IMAGERemove(fmt.Sprintf(format, args...))
}

func (e *SLAVATARElement) LABEL(s string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("label", s)
	return e
}

func (e *SLAVATARElement) LABELF(format string, args ...any) *SLAVATARElement {
	return e.LABEL(fmt.Sprintf(format, args...))
}

func (e *SLAVATARElement) IfLABEL(condition bool, s string) *SLAVATARElement {
	if condition {
		e.LABEL(s)
	}
	return e
}

func (e *SLAVATARElement) IfLABELF(condition bool, format string, args ...any) *SLAVATARElement {
	if condition {
		e.LABEL(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute LABEL from the element.
func (e *SLAVATARElement) LABELRemove(s string) *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("label")
	return e
}

func (e *SLAVATARElement) LABELRemoveF(format string, args ...any) *SLAVATARElement {
	return e.LABELRemove(fmt.Sprintf(format, args...))
}

func (e *SLAVATARElement) LOADING(c SLAvatarLoadingChoice) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("loading", string(c))
	return e
}

type SLAvatarLoadingChoice string

const (
	// Lazy
	SLAvatarLoading_lazy SLAvatarLoadingChoice = "lazy"
)

// Remove the attribute LOADING from the element.
func (e *SLAVATARElement) LOADINGRemove(c SLAvatarLoadingChoice) *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("loading")
	return e
}

func (e *SLAVATARElement) INITIALS(s string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("initials", s)
	return e
}

func (e *SLAVATARElement) INITIALSF(format string, args ...any) *SLAVATARElement {
	return e.INITIALS(fmt.Sprintf(format, args...))
}

func (e *SLAVATARElement) IfINITIALS(condition bool, s string) *SLAVATARElement {
	if condition {
		e.INITIALS(s)
	}
	return e
}

func (e *SLAVATARElement) IfINITIALSF(condition bool, format string, args ...any) *SLAVATARElement {
	if condition {
		e.INITIALS(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute INITIALS from the element.
func (e *SLAVATARElement) INITIALSRemove(s string) *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("initials")
	return e
}

func (e *SLAVATARElement) INITIALSRemoveF(format string, args ...any) *SLAVATARElement {
	return e.INITIALSRemove(fmt.Sprintf(format, args...))
}

func (e *SLAVATARElement) SHAPE(c SLAvatarShapeChoice) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("shape", string(c))
	return e
}

type SLAvatarShapeChoice string

const (
	// Circle
	SLAvatarShape_circle SLAvatarShapeChoice = "circle"
	// Rounded
	SLAvatarShape_rounded SLAvatarShapeChoice = "rounded"
	// Square
	SLAvatarShape_square SLAvatarShapeChoice = "square"
)

// Remove the attribute SHAPE from the element.
func (e *SLAVATARElement) SHAPERemove(c SLAvatarShapeChoice) *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("shape")
	return e
}

// Merges the store with the given object

func (e *SLAVATARElement) DATASTAR_MERGE_STORE(v any) *SLAVATARElement {
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

func (e *SLAVATARElement) DATASTAR_REF(expression string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLAVATARElement) IfDATASTAR_REF(condition bool, expression string) *SLAVATARElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SLAVATARElement) DATASTAR_REFRemove() *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SLAVATARElement) DATASTAR_BIND(key string, expression string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLAVATARElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SLAVATARElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SLAVATARElement) DATASTAR_BINDRemove() *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SLAVATARElement) DATASTAR_MODEL(expression string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLAVATARElement) IfDATASTAR_MODEL(condition bool, expression string) *SLAVATARElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SLAVATARElement) DATASTAR_MODELRemove() *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SLAVATARElement) DATASTAR_TEXT(expression string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLAVATARElement) IfDATASTAR_TEXT(condition bool, expression string) *SLAVATARElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SLAVATARElement) DATASTAR_TEXTRemove() *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SLAvatarOnMod customDataKeyModifier

// Debounces the event handler
func SLAvatarOnModDebounce(
	d time.Duration,
) SLAvatarOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SLAvatarOnModThrottle(
	d time.Duration,
) SLAvatarOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SLAVATARElement) DATASTAR_ON(key string, expression string, modifiers ...SLAvatarOnMod) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SLAvatarOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLAVATARElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SLAvatarOnMod) *SLAVATARElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SLAVATARElement) DATASTAR_ONRemove() *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SLAVATARElement) DATASTAR_FOCUSSet(b bool) *SLAVATARElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLAVATARElement) DATASTAR_FOCUS() *SLAVATARElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SLAVATARElement) DATASTAR_HEADER(key string, expression string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLAVATARElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SLAVATARElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SLAVATARElement) DATASTAR_HEADERRemove() *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SLAVATARElement) DATASTAR_FETCH_INDICATOR(expression string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLAVATARElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SLAVATARElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SLAVATARElement) DATASTAR_FETCH_INDICATORRemove() *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SLAVATARElement) DATASTAR_SHOWSet(b bool) *SLAVATARElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLAVATARElement) DATASTAR_SHOW() *SLAVATARElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SLAVATARElement) DATASTAR_INTERSECTS(expression string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLAVATARElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SLAVATARElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SLAVATARElement) DATASTAR_INTERSECTSRemove() *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SLAVATARElement) DATASTAR_TELEPORTSet(b bool) *SLAVATARElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLAVATARElement) DATASTAR_TELEPORT() *SLAVATARElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SLAVATARElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SLAVATARElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SLAVATARElement) DATASTAR_SCROLL_INTO_VIEW() *SLAVATARElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SLAVATARElement) DATASTAR_VIEW_TRANSITION(expression string) *SLAVATARElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SLAVATARElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SLAVATARElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SLAVATARElement) DATASTAR_VIEW_TRANSITIONRemove() *SLAVATARElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
