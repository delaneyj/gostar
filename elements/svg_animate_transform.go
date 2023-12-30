// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg animateTransform is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <animateTransform> SVG element animates a transformation attribute on a
// target element, thereby allowing animations to control translation, scaling,
// rotation and/or skewing.
type SVGANIMATETRANSFORMElement struct {
	*Element
}

// Create a new SVGANIMATETRANSFORMElement element.
// This will create a new element with the tag
// "animateTransform" during rendering.
func SVG_ANIMATETRANSFORM(children ...ElementRenderer) *SVGANIMATETRANSFORMElement {
	e := NewElement("animateTransform", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGANIMATETRANSFORMElement{Element: e}
}

func (e *SVGANIMATETRANSFORMElement) Children(children ...ElementRenderer) *SVGANIMATETRANSFORMElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfChildren(condition bool, children ...ElementRenderer) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) Attr(name string, value string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGANIMATETRANSFORMElement) Attrs(attrs ...string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) AttrsMap(attrs map[string]string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) Text(text string) *SVGANIMATETRANSFORMElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGANIMATETRANSFORMElement) TextF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfText(condition bool, text string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfTextF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) Escaped(text string) *SVGANIMATETRANSFORMElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfEscaped(condition bool, text string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) EscapedF(format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfEscapedF(condition bool, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) CustomData(key, value string) *SVGANIMATETRANSFORMElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfCustomData(condition bool, key, value string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) CustomDataF(key, format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) CustomDataRemove(key string) *SVGANIMATETRANSFORMElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Controls whether or not the animation is cumulative.
func (e *SVGANIMATETRANSFORMElement) ACCUMULATE(c SVGAnimateTransformAccumulateChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("accumulate", string(c))
	return e
}

type SVGAnimateTransformAccumulateChoice string

const (
	// The animation is not cumulative
	// Each iteration starts over from the beginning.
	SVGAnimateTransformAccumulate_none SVGAnimateTransformAccumulateChoice = "none"
	// The animation is cumulative
	// Each iteration the animation picks up where it left off in the previous
	// iteration.
	SVGAnimateTransformAccumulate_sum SVGAnimateTransformAccumulateChoice = "sum"
)

// Remove the attribute ACCUMULATE from the element.
func (e *SVGANIMATETRANSFORMElement) ACCUMULATERemove(c SVGAnimateTransformAccumulateChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("accumulate")
	return e
}

// Controls whether or not the animation is additive.
func (e *SVGANIMATETRANSFORMElement) ADDITIVE(c SVGAnimateTransformAdditiveChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("additive", string(c))
	return e
}

type SVGAnimateTransformAdditiveChoice string

const (
	// The animation is not additive
	// The animation replaces the underlying value.
	SVGAnimateTransformAdditive_replace SVGAnimateTransformAdditiveChoice = "replace"
	// The animation is additive
	// The animation adds to the underlying value.
	SVGAnimateTransformAdditive_sum SVGAnimateTransformAdditiveChoice = "sum"
)

// Remove the attribute ADDITIVE from the element.
func (e *SVGANIMATETRANSFORMElement) ADDITIVERemove(c SVGAnimateTransformAdditiveChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("additive")
	return e
}

// The name of the attribute to animate.
func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_NAME(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("attributeName", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfATTRIBUTE_NAME(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.ATTRIBUTE_NAME(s)
	}
	return e
}

// Remove the attribute ATTRIBUTE_NAME from the element.
func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_NAMERemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("attributeName")
	return e
}

// The namespace of the attribute to animate.
func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_TYPE(c SVGAnimateTransformAttributeTypeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("attributeType", string(c))
	return e
}

type SVGAnimateTransformAttributeTypeChoice string

const (
	// If the attribute is a presentation attribute, the animation will use the target
	// element's corresponding baseVal
	// If the attribute is not a presentation attribute, the animation will use the
	// target element's corresponding animVal.
	SVGAnimateTransformAttributeType_auto SVGAnimateTransformAttributeTypeChoice = "auto"
	// The animation will use the CSS namespace.
	SVGAnimateTransformAttributeType_CSS SVGAnimateTransformAttributeTypeChoice = "CSS"
	// The animation will use the XML namespace.
	SVGAnimateTransformAttributeType_XML SVGAnimateTransformAttributeTypeChoice = "XML"
	// The animation will use the XML ID namespace.
	SVGAnimateTransformAttributeType_XMLID SVGAnimateTransformAttributeTypeChoice = "XMLID"
	// The animation will use the XML LANG namespace.
	SVGAnimateTransformAttributeType_XMLLANG SVGAnimateTransformAttributeTypeChoice = "XMLLANG"
	// The animation will use the XML SPACE namespace.
	SVGAnimateTransformAttributeType_XMLSPACE SVGAnimateTransformAttributeTypeChoice = "XMLSPACE"
)

// Remove the attribute ATTRIBUTE_TYPE from the element.
func (e *SVGANIMATETRANSFORMElement) ATTRIBUTE_TYPERemove(c SVGAnimateTransformAttributeTypeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("attributeType")
	return e
}

// Defines when the animation should begin.
func (e *SVGANIMATETRANSFORMElement) BEGIN(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("begin", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfBEGIN(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.BEGIN(s)
	}
	return e
}

// Remove the attribute BEGIN from the element.
func (e *SVGANIMATETRANSFORMElement) BEGINRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("begin")
	return e
}

// Defines a relative offset value for the animation.
func (e *SVGANIMATETRANSFORMElement) BY(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("by", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfBY(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.BY(s)
	}
	return e
}

// Remove the attribute BY from the element.
func (e *SVGANIMATETRANSFORMElement) BYRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("by")
	return e
}

// Defines the pacing of the animation.
func (e *SVGANIMATETRANSFORMElement) CALC_MODE(c SVGAnimateTransformCalcModeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("calcMode", string(c))
	return e
}

type SVGAnimateTransformCalcModeChoice string

const (
	// The animation is not paced
	// Each iteration of the animation is displayed as fast as possible.
	SVGAnimateTransformCalcMode_discrete SVGAnimateTransformCalcModeChoice = "discrete"
	// The animation is paced such that it takes the same amount of time to go from
	// the start value to the end value throughout the animation.
	SVGAnimateTransformCalcMode_linear SVGAnimateTransformCalcModeChoice = "linear"
	// The animation is paced according to a cubic function.
	SVGAnimateTransformCalcMode_paced SVGAnimateTransformCalcModeChoice = "paced"
	// The animation is paced according to a cubic function, but with easing at both
	// the start and end.
	SVGAnimateTransformCalcMode_spline SVGAnimateTransformCalcModeChoice = "spline"
)

// Remove the attribute CALC_MODE from the element.
func (e *SVGANIMATETRANSFORMElement) CALC_MODERemove(c SVGAnimateTransformCalcModeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("calcMode")
	return e
}

// Defines the duration of the animation.
func (e *SVGANIMATETRANSFORMElement) DUR(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dur", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDUR(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DUR(s)
	}
	return e
}

// Remove the attribute DUR from the element.
func (e *SVGANIMATETRANSFORMElement) DURRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dur")
	return e
}

// Defines when the animation should end.
func (e *SVGANIMATETRANSFORMElement) END(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("end", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfEND(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.END(s)
	}
	return e
}

// Remove the attribute END from the element.
func (e *SVGANIMATETRANSFORMElement) ENDRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("end")
	return e
}

// Defines the fill behavior for the animation.
func (e *SVGANIMATETRANSFORMElement) FILL(c SVGAnimateTransformFillChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("fill", string(c))
	return e
}

type SVGAnimateTransformFillChoice string

const (
	// The animation will hold the attribute value when the animation ends.
	SVGAnimateTransformFill_freeze SVGAnimateTransformFillChoice = "freeze"
	// The animation will remove the attribute value when the animation ends.
	SVGAnimateTransformFill_remove SVGAnimateTransformFillChoice = "remove"
)

// Remove the attribute FILL from the element.
func (e *SVGANIMATETRANSFORMElement) FILLRemove(c SVGAnimateTransformFillChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("fill")
	return e
}

// Defines the initial value of the attribute.
func (e *SVGANIMATETRANSFORMElement) FROM(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("from", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfFROM(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.FROM(s)
	}
	return e
}

// Remove the attribute FROM from the element.
func (e *SVGANIMATETRANSFORMElement) FROMRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("from")
	return e
}

// Defines the values for a cubic Bézier function that controls interval pacing.
func (e *SVGANIMATETRANSFORMElement) KEY_SPLINES(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("keySplines", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfKEY_SPLINES(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.KEY_SPLINES(s)
	}
	return e
}

// Remove the attribute KEY_SPLINES from the element.
func (e *SVGANIMATETRANSFORMElement) KEY_SPLINESRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("keySplines")
	return e
}

// Defines when the animation should take place in terms of time fractions.
func (e *SVGANIMATETRANSFORMElement) KEY_TIMES(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("keyTimes", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfKEY_TIMES(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.KEY_TIMES(s)
	}
	return e
}

// Remove the attribute KEY_TIMES from the element.
func (e *SVGANIMATETRANSFORMElement) KEY_TIMESRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("keyTimes")
	return e
}

// Defines the maximum value allowed for the attribute.
func (e *SVGANIMATETRANSFORMElement) MAX(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("max", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfMAX(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.MAX(s)
	}
	return e
}

// Remove the attribute MAX from the element.
func (e *SVGANIMATETRANSFORMElement) MAXRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("max")
	return e
}

// Defines the minimum value allowed for the attribute.
func (e *SVGANIMATETRANSFORMElement) MIN(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("min", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfMIN(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.MIN(s)
	}
	return e
}

// Remove the attribute MIN from the element.
func (e *SVGANIMATETRANSFORMElement) MINRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("min")
	return e
}

// Defines the number of times the animation should repeat.
func (e *SVGANIMATETRANSFORMElement) REPEAT_COUNT(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("repeatCount", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfREPEAT_COUNT(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.REPEAT_COUNT(s)
	}
	return e
}

// Remove the attribute REPEAT_COUNT from the element.
func (e *SVGANIMATETRANSFORMElement) REPEAT_COUNTRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("repeatCount")
	return e
}

// Defines the duration for repeating an animation.
func (e *SVGANIMATETRANSFORMElement) REPEAT_DUR(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("repeatDur", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfREPEAT_DUR(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.REPEAT_DUR(s)
	}
	return e
}

// Remove the attribute REPEAT_DUR from the element.
func (e *SVGANIMATETRANSFORMElement) REPEAT_DURRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("repeatDur")
	return e
}

// Defines if an animation should restart after it completes.
func (e *SVGANIMATETRANSFORMElement) RESTART(c SVGAnimateTransformRestartChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("restart", string(c))
	return e
}

type SVGAnimateTransformRestartChoice string

const (
	// The animation will restart indefinitely.
	SVGAnimateTransformRestart_always SVGAnimateTransformRestartChoice = "always"
	// The animation will not restart after it completes.
	SVGAnimateTransformRestart_never SVGAnimateTransformRestartChoice = "never"
	// The animation will restart after it completes if the animation is not currently
	// active.
	SVGAnimateTransformRestart_whenNotActive SVGAnimateTransformRestartChoice = "whenNotActive"
)

// Remove the attribute RESTART from the element.
func (e *SVGANIMATETRANSFORMElement) RESTARTRemove(c SVGAnimateTransformRestartChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("restart")
	return e
}

// Defines the ending value of the attribute.
func (e *SVGANIMATETRANSFORMElement) TO(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("to", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfTO(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.TO(s)
	}
	return e
}

// Remove the attribute TO from the element.
func (e *SVGANIMATETRANSFORMElement) TORemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("to")
	return e
}

// Defines which transform to use.
func (e *SVGANIMATETRANSFORMElement) TYPE(c SVGAnimateTransformTypeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGAnimateTransformTypeChoice string

const (
	// The animation will use the rotate transform.
	SVGAnimateTransformType_rotate SVGAnimateTransformTypeChoice = "rotate"
	// The animation will use the scale transform.
	SVGAnimateTransformType_scale SVGAnimateTransformTypeChoice = "scale"
	// The animation will use the translate transform.
	SVGAnimateTransformType_translate SVGAnimateTransformTypeChoice = "translate"
)

// Remove the attribute TYPE from the element.
func (e *SVGANIMATETRANSFORMElement) TYPERemove(c SVGAnimateTransformTypeChoice) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

// Defines a list of discrete values to interpolate.
func (e *SVGANIMATETRANSFORMElement) VALUES(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("values", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfVALUES(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.VALUES(s)
	}
	return e
}

// Remove the attribute VALUES from the element.
func (e *SVGANIMATETRANSFORMElement) VALUESRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("values")
	return e
}

// Specifies a unique id for an element
func (e *SVGANIMATETRANSFORMElement) ID(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfID(condition bool, s string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGANIMATETRANSFORMElement) IDRemove(s string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGANIMATETRANSFORMElement) CLASS(s ...string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) IfCLASS(condition bool, s ...string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGANIMATETRANSFORMElement) CLASSRemove(s ...string) *SVGANIMATETRANSFORMElement {
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
func (e *SVGANIMATETRANSFORMElement) STYLEF(k string, format string, args ...any) *SVGANIMATETRANSFORMElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGANIMATETRANSFORMElement) IfSTYLE(condition bool, k string, v string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGANIMATETRANSFORMElement) STYLE(k string, v string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGANIMATETRANSFORMElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGANIMATETRANSFORMElement) STYLEMap(m map[string]string) *SVGANIMATETRANSFORMElement {
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
func (e *SVGANIMATETRANSFORMElement) STYLEPairs(pairs ...string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGANIMATETRANSFORMElement) STYLERemove(keys ...string) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) DATASTAR_MERGE_STORE(v any) *SVGANIMATETRANSFORMElement {
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

func (e *SVGANIMATETRANSFORMElement) DATASTAR_REF(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDATASTAR_REF(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGANIMATETRANSFORMElement) DATASTAR_REFRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGANIMATETRANSFORMElement) DATASTAR_BIND(key string, expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGANIMATETRANSFORMElement) DATASTAR_BINDRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGANIMATETRANSFORMElement) DATASTAR_MODEL(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGANIMATETRANSFORMElement) DATASTAR_MODELRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGANIMATETRANSFORMElement) DATASTAR_TEXT(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGANIMATETRANSFORMElement) DATASTAR_TEXTRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGAnimateTransformDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGAnimateTransformDataOnModDebounce(
	d time.Duration,
) SVGAnimateTransformDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGAnimateTransformDataOnModThrottle(
	d time.Duration,
) SVGAnimateTransformDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGANIMATETRANSFORMElement) DATASTAR_ON(key string, expression string, modifiers ...SVGAnimateTransformDataOnMod) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGAnimateTransformDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGAnimateTransformDataOnMod) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGANIMATETRANSFORMElement) DATASTAR_ONRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGANIMATETRANSFORMElement) DATASTAR_FOCUSSet(b bool) *SVGANIMATETRANSFORMElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATETRANSFORMElement) DATASTAR_FOCUS() *SVGANIMATETRANSFORMElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGANIMATETRANSFORMElement) DATASTAR_HEADER(key string, expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGANIMATETRANSFORMElement) DATASTAR_HEADERRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGANIMATETRANSFORMElement) DATASTAR_FETCH_URL(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGANIMATETRANSFORMElement) DATASTAR_FETCH_URLRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGANIMATETRANSFORMElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGANIMATETRANSFORMElement) DATASTAR_FETCH_INDICATORRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGANIMATETRANSFORMElement) DATASTAR_SHOWSet(b bool) *SVGANIMATETRANSFORMElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATETRANSFORMElement) DATASTAR_SHOW() *SVGANIMATETRANSFORMElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGANIMATETRANSFORMElement) DATASTAR_INTERSECTSSet(b bool) *SVGANIMATETRANSFORMElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATETRANSFORMElement) DATASTAR_INTERSECTS() *SVGANIMATETRANSFORMElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *SVGANIMATETRANSFORMElement) DATASTAR_TELEPORTSet(b bool) *SVGANIMATETRANSFORMElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATETRANSFORMElement) DATASTAR_TELEPORT() *SVGANIMATETRANSFORMElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGANIMATETRANSFORMElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGANIMATETRANSFORMElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATETRANSFORMElement) DATASTAR_SCROLL_INTO_VIEW() *SVGANIMATETRANSFORMElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGANIMATETRANSFORMElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATETRANSFORMElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGANIMATETRANSFORMElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGANIMATETRANSFORMElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGANIMATETRANSFORMElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
