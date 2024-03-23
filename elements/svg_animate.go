// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg animate is generated from configuration file.
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

// The <animate> SVG element is used to animate an attribute or property of an
// element over time.
type SVGANIMATEElement struct {
	*Element
}

// Create a new SVGANIMATEElement element.
// This will create a new element with the tag
// "animate" during rendering.
func SVG_ANIMATE(children ...ElementRenderer) *SVGANIMATEElement {
	e := NewElement("animate", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGANIMATEElement{Element: e}
}

func (e *SVGANIMATEElement) Children(children ...ElementRenderer) *SVGANIMATEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGANIMATEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGANIMATEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGANIMATEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGANIMATEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGANIMATEElement) Attr(name string, value string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGANIMATEElement) Attrs(attrs ...string) *SVGANIMATEElement {
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

func (e *SVGANIMATEElement) AttrsMap(attrs map[string]string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGANIMATEElement) Text(text string) *SVGANIMATEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGANIMATEElement) TextF(format string, args ...any) *SVGANIMATEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfText(condition bool, text string) *SVGANIMATEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGANIMATEElement) IfTextF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGANIMATEElement) Escaped(text string) *SVGANIMATEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGANIMATEElement) IfEscaped(condition bool, text string) *SVGANIMATEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGANIMATEElement) EscapedF(format string, args ...any) *SVGANIMATEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfEscapedF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGANIMATEElement) CustomData(key, value string) *SVGANIMATEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGANIMATEElement) IfCustomData(condition bool, key, value string) *SVGANIMATEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGANIMATEElement) CustomDataF(key, format string, args ...any) *SVGANIMATEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGANIMATEElement) CustomDataRemove(key string) *SVGANIMATEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Controls whether or not the animation is cumulative.
func (e *SVGANIMATEElement) ACCUMULATE(c SVGAnimateAccumulateChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("accumulate", string(c))
	return e
}

type SVGAnimateAccumulateChoice string

const (
	// The animation is not cumulative
	// Each iteration starts over from the beginning.
	SVGAnimateAccumulate_none SVGAnimateAccumulateChoice = "none"
	// The animation is cumulative
	// Each iteration the animation picks up where it left off in the previous
	// iteration.
	SVGAnimateAccumulate_sum SVGAnimateAccumulateChoice = "sum"
)

// Remove the attribute ACCUMULATE from the element.
func (e *SVGANIMATEElement) ACCUMULATERemove(c SVGAnimateAccumulateChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("accumulate")
	return e
}

// Controls whether or not the animation is additive.
func (e *SVGANIMATEElement) ADDITIVE(c SVGAnimateAdditiveChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("additive", string(c))
	return e
}

type SVGAnimateAdditiveChoice string

const (
	// The animation is not additive
	// The animation replaces the underlying value.
	SVGAnimateAdditive_replace SVGAnimateAdditiveChoice = "replace"
	// The animation is additive
	// The animation adds to the underlying value.
	SVGAnimateAdditive_sum SVGAnimateAdditiveChoice = "sum"
)

// Remove the attribute ADDITIVE from the element.
func (e *SVGANIMATEElement) ADDITIVERemove(c SVGAnimateAdditiveChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("additive")
	return e
}

// The name of the attribute to animate.
func (e *SVGANIMATEElement) ATTRIBUTE_NAME(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("attributeName", s)
	return e
}

func (e *SVGANIMATEElement) ATTRIBUTE_NAMEF(format string, args ...any) *SVGANIMATEElement {
	return e.ATTRIBUTE_NAME(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfATTRIBUTE_NAME(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.ATTRIBUTE_NAME(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfATTRIBUTE_NAMEF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.ATTRIBUTE_NAME(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ATTRIBUTE_NAME from the element.
func (e *SVGANIMATEElement) ATTRIBUTE_NAMERemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("attributeName")
	return e
}

func (e *SVGANIMATEElement) ATTRIBUTE_NAMERemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.ATTRIBUTE_NAMERemove(fmt.Sprintf(format, args...))
}

// The namespace of the attribute to animate.
func (e *SVGANIMATEElement) ATTRIBUTE_TYPE(c SVGAnimateAttributeTypeChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("attributeType", string(c))
	return e
}

type SVGAnimateAttributeTypeChoice string

const (
	// If the attribute is a presentation attribute, the animation will use the target
	// element's corresponding baseVal
	// If the attribute is not a presentation attribute, the animation will use the
	// target element's corresponding animVal.
	SVGAnimateAttributeType_auto SVGAnimateAttributeTypeChoice = "auto"
	// The animation will use the CSS namespace.
	SVGAnimateAttributeType_CSS SVGAnimateAttributeTypeChoice = "CSS"
	// The animation will use the XML namespace.
	SVGAnimateAttributeType_XML SVGAnimateAttributeTypeChoice = "XML"
	// The animation will use the XML ID namespace.
	SVGAnimateAttributeType_XMLID SVGAnimateAttributeTypeChoice = "XMLID"
	// The animation will use the XML LANG namespace.
	SVGAnimateAttributeType_XMLLANG SVGAnimateAttributeTypeChoice = "XMLLANG"
	// The animation will use the XML SPACE namespace.
	SVGAnimateAttributeType_XMLSPACE SVGAnimateAttributeTypeChoice = "XMLSPACE"
)

// Remove the attribute ATTRIBUTE_TYPE from the element.
func (e *SVGANIMATEElement) ATTRIBUTE_TYPERemove(c SVGAnimateAttributeTypeChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("attributeType")
	return e
}

// Defines when the animation should begin.
func (e *SVGANIMATEElement) BEGIN(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("begin", s)
	return e
}

func (e *SVGANIMATEElement) BEGINF(format string, args ...any) *SVGANIMATEElement {
	return e.BEGIN(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfBEGIN(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.BEGIN(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfBEGINF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.BEGIN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute BEGIN from the element.
func (e *SVGANIMATEElement) BEGINRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("begin")
	return e
}

func (e *SVGANIMATEElement) BEGINRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.BEGINRemove(fmt.Sprintf(format, args...))
}

// Defines a relative offset value for the animation.
func (e *SVGANIMATEElement) BY(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("by", s)
	return e
}

func (e *SVGANIMATEElement) BYF(format string, args ...any) *SVGANIMATEElement {
	return e.BY(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfBY(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.BY(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfBYF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.BY(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute BY from the element.
func (e *SVGANIMATEElement) BYRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("by")
	return e
}

func (e *SVGANIMATEElement) BYRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.BYRemove(fmt.Sprintf(format, args...))
}

// Defines the pacing of the animation.
func (e *SVGANIMATEElement) CALC_MODE(c SVGAnimateCalcModeChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("calcMode", string(c))
	return e
}

type SVGAnimateCalcModeChoice string

const (
	// The animation is not paced
	// Each iteration of the animation is displayed as fast as possible.
	SVGAnimateCalcMode_discrete SVGAnimateCalcModeChoice = "discrete"
	// The animation is paced such that it takes the same amount of time to go from
	// the start value to the end value throughout the animation.
	SVGAnimateCalcMode_linear SVGAnimateCalcModeChoice = "linear"
	// The animation is paced according to a cubic function.
	SVGAnimateCalcMode_paced SVGAnimateCalcModeChoice = "paced"
	// The animation is paced according to a cubic function, but with easing at both
	// the start and end.
	SVGAnimateCalcMode_spline SVGAnimateCalcModeChoice = "spline"
)

// Remove the attribute CALC_MODE from the element.
func (e *SVGANIMATEElement) CALC_MODERemove(c SVGAnimateCalcModeChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("calcMode")
	return e
}

// Defines the duration of the animation.
func (e *SVGANIMATEElement) DUR(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dur", s)
	return e
}

func (e *SVGANIMATEElement) DURF(format string, args ...any) *SVGANIMATEElement {
	return e.DUR(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfDUR(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.DUR(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfDURF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.DUR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute DUR from the element.
func (e *SVGANIMATEElement) DURRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dur")
	return e
}

func (e *SVGANIMATEElement) DURRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.DURRemove(fmt.Sprintf(format, args...))
}

// Defines when the animation should end.
func (e *SVGANIMATEElement) END(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("end", s)
	return e
}

func (e *SVGANIMATEElement) ENDF(format string, args ...any) *SVGANIMATEElement {
	return e.END(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfEND(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.END(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfENDF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.END(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute END from the element.
func (e *SVGANIMATEElement) ENDRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("end")
	return e
}

func (e *SVGANIMATEElement) ENDRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.ENDRemove(fmt.Sprintf(format, args...))
}

// Defines the fill behavior for the animation.
func (e *SVGANIMATEElement) FILL(c SVGAnimateFillChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("fill", string(c))
	return e
}

type SVGAnimateFillChoice string

const (
	// The animation will hold the attribute value when the animation ends.
	SVGAnimateFill_freeze SVGAnimateFillChoice = "freeze"
	// The animation will remove the attribute value when the animation ends.
	SVGAnimateFill_remove SVGAnimateFillChoice = "remove"
)

// Remove the attribute FILL from the element.
func (e *SVGANIMATEElement) FILLRemove(c SVGAnimateFillChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("fill")
	return e
}

// Defines the initial value of the attribute.
func (e *SVGANIMATEElement) FROM(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("from", s)
	return e
}

func (e *SVGANIMATEElement) FROMF(format string, args ...any) *SVGANIMATEElement {
	return e.FROM(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfFROM(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.FROM(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfFROMF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.FROM(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute FROM from the element.
func (e *SVGANIMATEElement) FROMRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("from")
	return e
}

func (e *SVGANIMATEElement) FROMRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.FROMRemove(fmt.Sprintf(format, args...))
}

// Defines the values for a cubic Bézier function that controls interval pacing.
func (e *SVGANIMATEElement) KEY_SPLINES(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("keySplines", s)
	return e
}

func (e *SVGANIMATEElement) KEY_SPLINESF(format string, args ...any) *SVGANIMATEElement {
	return e.KEY_SPLINES(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfKEY_SPLINES(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.KEY_SPLINES(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfKEY_SPLINESF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.KEY_SPLINES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute KEY_SPLINES from the element.
func (e *SVGANIMATEElement) KEY_SPLINESRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("keySplines")
	return e
}

func (e *SVGANIMATEElement) KEY_SPLINESRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.KEY_SPLINESRemove(fmt.Sprintf(format, args...))
}

// Defines when the animation should take place in terms of time fractions.
func (e *SVGANIMATEElement) KEY_TIMES(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("keyTimes", s)
	return e
}

func (e *SVGANIMATEElement) KEY_TIMESF(format string, args ...any) *SVGANIMATEElement {
	return e.KEY_TIMES(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfKEY_TIMES(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.KEY_TIMES(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfKEY_TIMESF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.KEY_TIMES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute KEY_TIMES from the element.
func (e *SVGANIMATEElement) KEY_TIMESRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("keyTimes")
	return e
}

func (e *SVGANIMATEElement) KEY_TIMESRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.KEY_TIMESRemove(fmt.Sprintf(format, args...))
}

// Defines the maximum value allowed for the attribute.
func (e *SVGANIMATEElement) MAX(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("max", s)
	return e
}

func (e *SVGANIMATEElement) MAXF(format string, args ...any) *SVGANIMATEElement {
	return e.MAX(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfMAX(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.MAX(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfMAXF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.MAX(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MAX from the element.
func (e *SVGANIMATEElement) MAXRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("max")
	return e
}

func (e *SVGANIMATEElement) MAXRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.MAXRemove(fmt.Sprintf(format, args...))
}

// Defines the minimum value allowed for the attribute.
func (e *SVGANIMATEElement) MIN(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("min", s)
	return e
}

func (e *SVGANIMATEElement) MINF(format string, args ...any) *SVGANIMATEElement {
	return e.MIN(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfMIN(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.MIN(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfMINF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.MIN(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MIN from the element.
func (e *SVGANIMATEElement) MINRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("min")
	return e
}

func (e *SVGANIMATEElement) MINRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.MINRemove(fmt.Sprintf(format, args...))
}

// Defines the number of times the animation should repeat.
func (e *SVGANIMATEElement) REPEAT_COUNT(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("repeatCount", s)
	return e
}

func (e *SVGANIMATEElement) REPEAT_COUNTF(format string, args ...any) *SVGANIMATEElement {
	return e.REPEAT_COUNT(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfREPEAT_COUNT(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.REPEAT_COUNT(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfREPEAT_COUNTF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.REPEAT_COUNT(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REPEAT_COUNT from the element.
func (e *SVGANIMATEElement) REPEAT_COUNTRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("repeatCount")
	return e
}

func (e *SVGANIMATEElement) REPEAT_COUNTRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.REPEAT_COUNTRemove(fmt.Sprintf(format, args...))
}

// Defines the duration for repeating an animation.
func (e *SVGANIMATEElement) REPEAT_DUR(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("repeatDur", s)
	return e
}

func (e *SVGANIMATEElement) REPEAT_DURF(format string, args ...any) *SVGANIMATEElement {
	return e.REPEAT_DUR(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfREPEAT_DUR(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.REPEAT_DUR(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfREPEAT_DURF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.REPEAT_DUR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute REPEAT_DUR from the element.
func (e *SVGANIMATEElement) REPEAT_DURRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("repeatDur")
	return e
}

func (e *SVGANIMATEElement) REPEAT_DURRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.REPEAT_DURRemove(fmt.Sprintf(format, args...))
}

// Defines if an animation should restart after it completes.
func (e *SVGANIMATEElement) RESTART(c SVGAnimateRestartChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("restart", string(c))
	return e
}

type SVGAnimateRestartChoice string

const (
	// The animation will restart indefinitely.
	SVGAnimateRestart_always SVGAnimateRestartChoice = "always"
	// The animation will not restart after it completes.
	SVGAnimateRestart_never SVGAnimateRestartChoice = "never"
	// The animation will restart after it completes if the animation is not currently
	// active.
	SVGAnimateRestart_whenNotActive SVGAnimateRestartChoice = "whenNotActive"
)

// Remove the attribute RESTART from the element.
func (e *SVGANIMATEElement) RESTARTRemove(c SVGAnimateRestartChoice) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("restart")
	return e
}

// Defines the ending value of the attribute.
func (e *SVGANIMATEElement) TO(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("to", s)
	return e
}

func (e *SVGANIMATEElement) TOF(format string, args ...any) *SVGANIMATEElement {
	return e.TO(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfTO(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.TO(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfTOF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.TO(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TO from the element.
func (e *SVGANIMATEElement) TORemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("to")
	return e
}

func (e *SVGANIMATEElement) TORemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.TORemove(fmt.Sprintf(format, args...))
}

// Defines a list of discrete values to interpolate.
func (e *SVGANIMATEElement) VALUES(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("values", s)
	return e
}

func (e *SVGANIMATEElement) VALUESF(format string, args ...any) *SVGANIMATEElement {
	return e.VALUES(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfVALUES(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.VALUES(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfVALUESF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.VALUES(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VALUES from the element.
func (e *SVGANIMATEElement) VALUESRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("values")
	return e
}

func (e *SVGANIMATEElement) VALUESRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.VALUESRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGANIMATEElement) ID(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGANIMATEElement) IDF(format string, args ...any) *SVGANIMATEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfID(condition bool, s string) *SVGANIMATEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGANIMATEElement) IfIDF(condition bool, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGANIMATEElement) IDRemove(s string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGANIMATEElement) IDRemoveF(format string, args ...any) *SVGANIMATEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGANIMATEElement) CLASS(s ...string) *SVGANIMATEElement {
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

func (e *SVGANIMATEElement) IfCLASS(condition bool, s ...string) *SVGANIMATEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGANIMATEElement) CLASSRemove(s ...string) *SVGANIMATEElement {
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
func (e *SVGANIMATEElement) STYLEF(k string, format string, args ...any) *SVGANIMATEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEElement) IfSTYLE(condition bool, k string, v string) *SVGANIMATEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGANIMATEElement) STYLE(k string, v string) *SVGANIMATEElement {
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

func (e *SVGANIMATEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGANIMATEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGANIMATEElement) STYLEMap(m map[string]string) *SVGANIMATEElement {
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
func (e *SVGANIMATEElement) STYLEPairs(pairs ...string) *SVGANIMATEElement {
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

func (e *SVGANIMATEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGANIMATEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGANIMATEElement) STYLERemove(keys ...string) *SVGANIMATEElement {
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

func (e *SVGANIMATEElement) DATASTAR_MERGE_STORE(v any) *SVGANIMATEElement {
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

func (e *SVGANIMATEElement) DATASTAR_REF(expression string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEElement) IfDATASTAR_REF(condition bool, expression string) *SVGANIMATEElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGANIMATEElement) DATASTAR_REFRemove() *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGANIMATEElement) DATASTAR_BIND(key string, expression string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGANIMATEElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGANIMATEElement) DATASTAR_BINDRemove() *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGANIMATEElement) DATASTAR_MODEL(expression string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGANIMATEElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGANIMATEElement) DATASTAR_MODELRemove() *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGANIMATEElement) DATASTAR_TEXT(expression string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGANIMATEElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGANIMATEElement) DATASTAR_TEXTRemove() *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGAnimateOnMod customDataKeyModifier

// Debounces the event handler
func SVGAnimateOnModDebounce(
	d time.Duration,
) SVGAnimateOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGAnimateOnModThrottle(
	d time.Duration,
) SVGAnimateOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGANIMATEElement) DATASTAR_ON(key string, expression string, modifiers ...SVGAnimateOnMod) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGAnimateOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGAnimateOnMod) *SVGANIMATEElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGANIMATEElement) DATASTAR_ONRemove() *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGANIMATEElement) DATASTAR_FOCUSSet(b bool) *SVGANIMATEElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATEElement) DATASTAR_FOCUS() *SVGANIMATEElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGANIMATEElement) DATASTAR_HEADER(key string, expression string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGANIMATEElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGANIMATEElement) DATASTAR_HEADERRemove() *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGANIMATEElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGANIMATEElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGANIMATEElement) DATASTAR_FETCH_INDICATORRemove() *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGANIMATEElement) DATASTAR_SHOWSet(b bool) *SVGANIMATEElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATEElement) DATASTAR_SHOW() *SVGANIMATEElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGANIMATEElement) DATASTAR_INTERSECTS(expression string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGANIMATEElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGANIMATEElement) DATASTAR_INTERSECTSRemove() *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGANIMATEElement) DATASTAR_TELEPORTSet(b bool) *SVGANIMATEElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATEElement) DATASTAR_TELEPORT() *SVGANIMATEElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGANIMATEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGANIMATEElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATEElement) DATASTAR_SCROLL_INTO_VIEW() *SVGANIMATEElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGANIMATEElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGANIMATEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGANIMATEElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGANIMATEElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGANIMATEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
