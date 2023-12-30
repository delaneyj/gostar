// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg animateMotion is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <animateMotion> SVG element is used to animate a transformation attribute
// on a target element, thereby allowing the animation of translation, rotation,
// and scaling.
type SVGANIMATEMOTIONElement struct {
	*Element
}

// Create a new SVGANIMATEMOTIONElement element.
// This will create a new element with the tag
// "animateMotion" during rendering.
func SVG_ANIMATEMOTION(children ...ElementRenderer) *SVGANIMATEMOTIONElement {
	e := NewElement("animateMotion", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGANIMATEMOTIONElement{Element: e}
}

func (e *SVGANIMATEMOTIONElement) Children(children ...ElementRenderer) *SVGANIMATEMOTIONElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfChildren(condition bool, children ...ElementRenderer) *SVGANIMATEMOTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGANIMATEMOTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) Attr(name string, value string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGANIMATEMOTIONElement) Attrs(attrs ...string) *SVGANIMATEMOTIONElement {
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

func (e *SVGANIMATEMOTIONElement) AttrsMap(attrs map[string]string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) Text(text string) *SVGANIMATEMOTIONElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGANIMATEMOTIONElement) TextF(format string, args ...any) *SVGANIMATEMOTIONElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEMOTIONElement) IfText(condition bool, text string) *SVGANIMATEMOTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) IfTextF(condition bool, format string, args ...any) *SVGANIMATEMOTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) Escaped(text string) *SVGANIMATEMOTIONElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGANIMATEMOTIONElement) IfEscaped(condition bool, text string) *SVGANIMATEMOTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) EscapedF(format string, args ...any) *SVGANIMATEMOTIONElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEMOTIONElement) IfEscapedF(condition bool, format string, args ...any) *SVGANIMATEMOTIONElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) CustomData(key, value string) *SVGANIMATEMOTIONElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfCustomData(condition bool, key, value string) *SVGANIMATEMOTIONElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) CustomDataF(key, format string, args ...any) *SVGANIMATEMOTIONElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEMOTIONElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGANIMATEMOTIONElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) CustomDataRemove(key string) *SVGANIMATEMOTIONElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Controls whether or not the animation is cumulative.
func (e *SVGANIMATEMOTIONElement) ACCUMULATE(c SVGAnimateMotionAccumulateChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("accumulate", string(c))
	return e
}

type SVGAnimateMotionAccumulateChoice string

const (
	// The animation is not cumulative
	// Each iteration starts over from the beginning.
	SVGAnimateMotionAccumulate_none SVGAnimateMotionAccumulateChoice = "none"
	// The animation is cumulative
	// Each iteration the animation picks up where it left off in the previous
	// iteration.
	SVGAnimateMotionAccumulate_sum SVGAnimateMotionAccumulateChoice = "sum"
)

// Remove the attribute ACCUMULATE from the element.
func (e *SVGANIMATEMOTIONElement) ACCUMULATERemove(c SVGAnimateMotionAccumulateChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("accumulate")
	return e
}

// Controls whether or not the animation is additive.
func (e *SVGANIMATEMOTIONElement) ADDITIVE(c SVGAnimateMotionAdditiveChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("additive", string(c))
	return e
}

type SVGAnimateMotionAdditiveChoice string

const (
	// The animation is not additive
	// The animation replaces the underlying value.
	SVGAnimateMotionAdditive_replace SVGAnimateMotionAdditiveChoice = "replace"
	// The animation is additive
	// The animation adds to the underlying value.
	SVGAnimateMotionAdditive_sum SVGAnimateMotionAdditiveChoice = "sum"
)

// Remove the attribute ADDITIVE from the element.
func (e *SVGANIMATEMOTIONElement) ADDITIVERemove(c SVGAnimateMotionAdditiveChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("additive")
	return e
}

// Defines when the animation should begin.
func (e *SVGANIMATEMOTIONElement) BEGIN(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("begin", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfBEGIN(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.BEGIN(s)
	}
	return e
}

// Remove the attribute BEGIN from the element.
func (e *SVGANIMATEMOTIONElement) BEGINRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("begin")
	return e
}

// Defines a relative offset value for the animation.
func (e *SVGANIMATEMOTIONElement) BY(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("by", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfBY(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.BY(s)
	}
	return e
}

// Remove the attribute BY from the element.
func (e *SVGANIMATEMOTIONElement) BYRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("by")
	return e
}

// Defines the pacing of the animation.
func (e *SVGANIMATEMOTIONElement) CALC_MODE(c SVGAnimateMotionCalcModeChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("calcMode", string(c))
	return e
}

type SVGAnimateMotionCalcModeChoice string

const (
	// The animation is not paced
	// Each iteration of the animation is displayed as fast as possible.
	SVGAnimateMotionCalcMode_discrete SVGAnimateMotionCalcModeChoice = "discrete"
	// The animation is paced such that it takes the same amount of time to go from
	// the start value to the end value throughout the animation.
	SVGAnimateMotionCalcMode_linear SVGAnimateMotionCalcModeChoice = "linear"
	// The animation is paced according to a cubic function.
	SVGAnimateMotionCalcMode_paced SVGAnimateMotionCalcModeChoice = "paced"
	// The animation is paced according to a cubic function, but with easing at both
	// the start and end.
	SVGAnimateMotionCalcMode_spline SVGAnimateMotionCalcModeChoice = "spline"
)

// Remove the attribute CALC_MODE from the element.
func (e *SVGANIMATEMOTIONElement) CALC_MODERemove(c SVGAnimateMotionCalcModeChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("calcMode")
	return e
}

// Defines the duration of the animation.
func (e *SVGANIMATEMOTIONElement) DUR(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dur", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDUR(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.DUR(s)
	}
	return e
}

// Remove the attribute DUR from the element.
func (e *SVGANIMATEMOTIONElement) DURRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dur")
	return e
}

// Defines when the animation should end.
func (e *SVGANIMATEMOTIONElement) END(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("end", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfEND(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.END(s)
	}
	return e
}

// Remove the attribute END from the element.
func (e *SVGANIMATEMOTIONElement) ENDRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("end")
	return e
}

// Defines the fill behavior for the animation.
func (e *SVGANIMATEMOTIONElement) FILL(c SVGAnimateMotionFillChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("fill", string(c))
	return e
}

type SVGAnimateMotionFillChoice string

const (
	// The animation will hold the attribute value when the animation ends.
	SVGAnimateMotionFill_freeze SVGAnimateMotionFillChoice = "freeze"
	// The animation will remove the attribute value when the animation ends.
	SVGAnimateMotionFill_remove SVGAnimateMotionFillChoice = "remove"
)

// Remove the attribute FILL from the element.
func (e *SVGANIMATEMOTIONElement) FILLRemove(c SVGAnimateMotionFillChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("fill")
	return e
}

// Defines the initial value of the attribute.
func (e *SVGANIMATEMOTIONElement) FROM(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("from", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfFROM(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.FROM(s)
	}
	return e
}

// Remove the attribute FROM from the element.
func (e *SVGANIMATEMOTIONElement) FROMRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("from")
	return e
}

// Defines the values for a cubic Bézier function that controls interval pacing.
func (e *SVGANIMATEMOTIONElement) KEY_SPLINES(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("keySplines", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfKEY_SPLINES(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.KEY_SPLINES(s)
	}
	return e
}

// Remove the attribute KEY_SPLINES from the element.
func (e *SVGANIMATEMOTIONElement) KEY_SPLINESRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("keySplines")
	return e
}

// Defines when the animation should take place in terms of time fractions.
func (e *SVGANIMATEMOTIONElement) KEY_TIMES(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("keyTimes", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfKEY_TIMES(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.KEY_TIMES(s)
	}
	return e
}

// Remove the attribute KEY_TIMES from the element.
func (e *SVGANIMATEMOTIONElement) KEY_TIMESRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("keyTimes")
	return e
}

// Defines the maximum value allowed for the attribute.
func (e *SVGANIMATEMOTIONElement) MAX(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("max", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfMAX(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.MAX(s)
	}
	return e
}

// Remove the attribute MAX from the element.
func (e *SVGANIMATEMOTIONElement) MAXRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("max")
	return e
}

// Defines the minimum value allowed for the attribute.
func (e *SVGANIMATEMOTIONElement) MIN(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("min", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfMIN(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.MIN(s)
	}
	return e
}

// Remove the attribute MIN from the element.
func (e *SVGANIMATEMOTIONElement) MINRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("min")
	return e
}

// Defines the number of times the animation should repeat.
func (e *SVGANIMATEMOTIONElement) REPEAT_COUNT(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("repeatCount", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfREPEAT_COUNT(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.REPEAT_COUNT(s)
	}
	return e
}

// Remove the attribute REPEAT_COUNT from the element.
func (e *SVGANIMATEMOTIONElement) REPEAT_COUNTRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("repeatCount")
	return e
}

// Defines the duration for repeating an animation.
func (e *SVGANIMATEMOTIONElement) REPEAT_DUR(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("repeatDur", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfREPEAT_DUR(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.REPEAT_DUR(s)
	}
	return e
}

// Remove the attribute REPEAT_DUR from the element.
func (e *SVGANIMATEMOTIONElement) REPEAT_DURRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("repeatDur")
	return e
}

// Defines if an animation should restart after it completes.
func (e *SVGANIMATEMOTIONElement) RESTART(c SVGAnimateMotionRestartChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("restart", string(c))
	return e
}

type SVGAnimateMotionRestartChoice string

const (
	// The animation will restart indefinitely.
	SVGAnimateMotionRestart_always SVGAnimateMotionRestartChoice = "always"
	// The animation will not restart after it completes.
	SVGAnimateMotionRestart_never SVGAnimateMotionRestartChoice = "never"
	// The animation will restart after it completes if the animation is not currently
	// active.
	SVGAnimateMotionRestart_whenNotActive SVGAnimateMotionRestartChoice = "whenNotActive"
)

// Remove the attribute RESTART from the element.
func (e *SVGANIMATEMOTIONElement) RESTARTRemove(c SVGAnimateMotionRestartChoice) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("restart")
	return e
}

// Defines the ending value of the attribute.
func (e *SVGANIMATEMOTIONElement) TO(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("to", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfTO(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.TO(s)
	}
	return e
}

// Remove the attribute TO from the element.
func (e *SVGANIMATEMOTIONElement) TORemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("to")
	return e
}

// Defines a list of discrete values to interpolate.
func (e *SVGANIMATEMOTIONElement) VALUES(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("values", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfVALUES(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.VALUES(s)
	}
	return e
}

// Remove the attribute VALUES from the element.
func (e *SVGANIMATEMOTIONElement) VALUESRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("values")
	return e
}

// Specifies a unique id for an element
func (e *SVGANIMATEMOTIONElement) ID(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfID(condition bool, s string) *SVGANIMATEMOTIONElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGANIMATEMOTIONElement) IDRemove(s string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGANIMATEMOTIONElement) CLASS(s ...string) *SVGANIMATEMOTIONElement {
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

func (e *SVGANIMATEMOTIONElement) IfCLASS(condition bool, s ...string) *SVGANIMATEMOTIONElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGANIMATEMOTIONElement) CLASSRemove(s ...string) *SVGANIMATEMOTIONElement {
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
func (e *SVGANIMATEMOTIONElement) STYLEF(k string, format string, args ...any) *SVGANIMATEMOTIONElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGANIMATEMOTIONElement) IfSTYLE(condition bool, k string, v string) *SVGANIMATEMOTIONElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGANIMATEMOTIONElement) STYLE(k string, v string) *SVGANIMATEMOTIONElement {
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

func (e *SVGANIMATEMOTIONElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGANIMATEMOTIONElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGANIMATEMOTIONElement) STYLEMap(m map[string]string) *SVGANIMATEMOTIONElement {
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
func (e *SVGANIMATEMOTIONElement) STYLEPairs(pairs ...string) *SVGANIMATEMOTIONElement {
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

func (e *SVGANIMATEMOTIONElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGANIMATEMOTIONElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGANIMATEMOTIONElement) STYLERemove(keys ...string) *SVGANIMATEMOTIONElement {
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

func (e *SVGANIMATEMOTIONElement) DATASTAR_MERGE_STORE(v any) *SVGANIMATEMOTIONElement {
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

func (e *SVGANIMATEMOTIONElement) DATASTAR_REF(expression string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDATASTAR_REF(condition bool, expression string) *SVGANIMATEMOTIONElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGANIMATEMOTIONElement) DATASTAR_REFRemove() *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGANIMATEMOTIONElement) DATASTAR_BIND(key string, expression string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGANIMATEMOTIONElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGANIMATEMOTIONElement) DATASTAR_BINDRemove() *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGANIMATEMOTIONElement) DATASTAR_MODEL(expression string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGANIMATEMOTIONElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGANIMATEMOTIONElement) DATASTAR_MODELRemove() *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGANIMATEMOTIONElement) DATASTAR_TEXT(expression string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGANIMATEMOTIONElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGANIMATEMOTIONElement) DATASTAR_TEXTRemove() *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGAnimateMotionDataOnMod customDataKeyModifier

// Debounces the event handler
func SVGAnimateMotionDataOnModDebounce(
	d time.Duration,
) SVGAnimateMotionDataOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGAnimateMotionDataOnModThrottle(
	d time.Duration,
) SVGAnimateMotionDataOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGANIMATEMOTIONElement) DATASTAR_ON(key string, expression string, modifiers ...SVGAnimateMotionDataOnMod) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGAnimateMotionDataOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGAnimateMotionDataOnMod) *SVGANIMATEMOTIONElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGANIMATEMOTIONElement) DATASTAR_ONRemove() *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGANIMATEMOTIONElement) DATASTAR_FOCUSSet(b bool) *SVGANIMATEMOTIONElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATEMOTIONElement) DATASTAR_FOCUS() *SVGANIMATEMOTIONElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGANIMATEMOTIONElement) DATASTAR_HEADER(key string, expression string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGANIMATEMOTIONElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGANIMATEMOTIONElement) DATASTAR_HEADERRemove() *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGANIMATEMOTIONElement) DATASTAR_FETCH_URL(expression string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGANIMATEMOTIONElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGANIMATEMOTIONElement) DATASTAR_FETCH_URLRemove() *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGANIMATEMOTIONElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "DatastarFetchIndicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGANIMATEMOTIONElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGANIMATEMOTIONElement) DATASTAR_FETCH_INDICATORRemove() *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("DatastarFetchIndicator")
	return e
}

// Sets the visibility of the element

func (e *SVGANIMATEMOTIONElement) DATASTAR_SHOWSet(b bool) *SVGANIMATEMOTIONElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATEMOTIONElement) DATASTAR_SHOW() *SVGANIMATEMOTIONElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGANIMATEMOTIONElement) DATASTAR_INTERSECTSSet(b bool) *SVGANIMATEMOTIONElement {
	key := "data-intersects"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATEMOTIONElement) DATASTAR_INTERSECTS() *SVGANIMATEMOTIONElement {
	return e.DATASTAR_INTERSECTSSet(true)
}

// Teleports the element to the given selector

func (e *SVGANIMATEMOTIONElement) DATASTAR_TELEPORTSet(b bool) *SVGANIMATEMOTIONElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATEMOTIONElement) DATASTAR_TELEPORT() *SVGANIMATEMOTIONElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGANIMATEMOTIONElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGANIMATEMOTIONElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGANIMATEMOTIONElement) DATASTAR_SCROLL_INTO_VIEW() *SVGANIMATEMOTIONElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGANIMATEMOTIONElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-view-transition-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGANIMATEMOTIONElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGANIMATEMOTIONElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGANIMATEMOTIONElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGANIMATEMOTIONElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
