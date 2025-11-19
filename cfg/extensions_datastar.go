package cfg

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"
)

var DatastarExtensions = []*pb.Attribute{
	{
		Name: "DatastarAttr",
		Key:  "attr",
		Description: doc(
			"Sets the value of any HTML attribute to an expression, and keeps it in sync.",
			"https://data-star.dev/reference/attributes#data-attr",
		),
		Type: AttributeTypeCustom(true, AttributeTypeString()),
	},
	{
		Name: "DatastarBind",
		Key:  "bind",
		Description: doc(
			"Creates a signal (if one doesn’t already exist) and sets up two-way data binding between it and an element’s value.",
			"https://data-star.dev/reference/attributes#data-bind",
		),
		Type: AttributeTypeCustom(true, AttributeTypeString()),
	},
	{
		Name:        "DatastarClass",
		Key:         "datastar-class",
		Description: "Adds or removes a class to or from an element based on an expression.",
		Type: AttributeTypeCustom(true, AttributeTypeString(),
			datastarModifiers.Case,
		),
	},
	{
		Name: "DatastarComputed",
		Key:  "computed",
		Description: doc(
			"Creates a signal that is computed based on an expression. The computed signal is read-only, and its value is automatically updated when any signals in the expression are updated.",
			"https://data-star.dev/reference/attributes#data-computed",
		),
		Type: AttributeTypeCustom(true, AttributeTypeString(),
			datastarModifiers.Case,
		),
	},
	{
		Name: "DatastarEffect",
		Key:  "effect",
		Description: doc(
			"Executes an expression on page load and whenever any signals in the expression change. This is useful for performing side effects, such as updating other signals, making requests to the backend, or manipulating the DOM.",
			"https://data-star.dev/reference/attributes#data-effect",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString()),
	},
	{
		Name: "DatastarIgnore",
		Key:  "ignore",
		Description: doc(
			"Datastar walks the entire DOM and applies plugins to each element it encounters. It's possible to tell Datastar to ignore an element and its descendants by placing a data-ignore attribute on it. This can be useful for preventing naming conflicts with third-party libraries, or when you are unable to escape user input.",
			"https://data-star.dev/reference/attributes#data-ignore",
		),
		Type: AttributeTypeCustom(false, AttributeTypeBool(),
			datastarModifiers.Self,
		),
	},
	{
		Name: "DatastarIgnoreMorph",
		Key:  "ignore-morph",
		Description: doc(
			"Similar to the data-ignore attribute, the data-ignore-morph attribute tells the PatchElements watcher to skip processing an element and its children when morphing elements. This can be useful for preventing conflicts with third-party libraries that manipulate the DOM, or when you are unable to escape user input.",
			"https://data-star.dev/reference/attributes#data-ignore-morph",
		),
		Type: AttributeTypeCustom(false, AttributeTypeBool()),
	},
	{
		Name: "DatastarIndicator",
		Key:  "indicator",
		Description: doc(
			"Creates a signal and sets its value to true while a fetch request is in flight, otherwise false. The signal can be used to show a loading indicator.",
			"https://data-star.dev/reference/attributes#data-indicator",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString(),
			datastarModifiers.Case,
		),
	},
	{
		Name: "DatastarInit",
		Key:  "init",
		Description: doc(
			"Runs an expression when the attribute is initialized. This can happen on page load, when an element is patched into the DOM, and any time the attribute is modified (via a backend action or otherwise).",
			"https://data-star.dev/reference/attributes#data-indicator",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString(),
			datastarModifiers.DelayMs,
			datastarModifiers.DelaySec,
			datastarModifiers.ViewTransition,
		),
	},
	{
		Name: "DatastarJSONSignals",
		Key:  "json-signals",
		Description: doc(
			"Sets the text content of an element to a reactive JSON stringified version of signals. Useful when troubleshooting an issue.",
			"https://data-star.dev/reference/attributes#data-json-signals",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString(),
			datastarModifiers.Terse,
		),
	},
	{
		Name: "DatastarOn",
		Key:  "on",
		Description: doc(
			"Attaches an event listener to an element, executing an expression whenever the event is triggered.",
			"https://data-star.dev/reference/attributes#data-on",
		),
		Type: AttributeTypeCustom(true, AttributeTypeString(),
			datastarModifiers.Once,
			datastarModifiers.Passive,
			datastarModifiers.Capture,
			datastarModifiers.Case,
			datastarModifiers.DelayMs,
			datastarModifiers.DelaySec,
			datastarModifiers.DebounceMs,
			datastarModifiers.DebounceMsLeading,
			datastarModifiers.DebounceMsNoTrailing,
			datastarModifiers.DebounceSec,
			datastarModifiers.DebounceSecLeading,
			datastarModifiers.DebounceSecNoTrailing,
			datastarModifiers.ThrottleMs,
			datastarModifiers.ThrottleMsNoLeading,
			datastarModifiers.ThrottleMsTrailing,
			datastarModifiers.ThrotlleSec,
			datastarModifiers.ThrotlleSecNoLeading,
			datastarModifiers.ThrotlleSecTrailing,
			datastarModifiers.ViewTransition,
			datastarModifiers.Window,
			datastarModifiers.Prevent,
			datastarModifiers.Outside,
			datastarModifiers.Stop,
		),
	},
	{
		Name:        "DatastarOnIntersect",
		Key:         "on-intersect",
		Description: "Runs an expression when the element intersects with the viewport.",
		Type: AttributeTypeCustom(false, AttributeTypeString(),
			datastarModifiers.Once,
			datastarModifiers.Half,
			datastarModifiers.Full,
			datastarModifiers.DelayMs,
			datastarModifiers.DelaySec,
			datastarModifiers.DebounceMs,
			datastarModifiers.DebounceMsLeading,
			datastarModifiers.DebounceMsNoTrailing,
			datastarModifiers.DebounceSec,
			datastarModifiers.DebounceSecLeading,
			datastarModifiers.DebounceSecNoTrailing,
			datastarModifiers.ThrottleMs,
			datastarModifiers.ThrottleMsNoLeading,
			datastarModifiers.ThrottleMsTrailing,
			datastarModifiers.ThrotlleSec,
			datastarModifiers.ThrotlleSecNoLeading,
			datastarModifiers.ThrotlleSecTrailing,
			datastarModifiers.ViewTransition,
		),
	},
	{
		Name: "DatastarOnInterval",
		Key:  "on-interval",
		Description: doc(
			"Runs an expression at a regular interval. The interval duration defaults to one second and can be modified using the '__duration' modifier.",
			"https://data-star.dev/reference/attributes#data-on-interval",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString(),
			datastarModifiers.DurationMs,
			datastarModifiers.DurationMsLeading,
			datastarModifiers.DurationSec,
			datastarModifiers.DurationSecLeading,
			datastarModifiers.ViewTransition,
		),
	},
	{
		Name: "DatastarOnSignalPatch",
		Key:  "on-signal-patch",
		Description: doc(
			"Runs an expression whenever any signals are patched. This is useful for tracking changes, updating computed values, or triggering side effects when data updates.",
			"https://data-star.dev/reference/attributes#data-on-signal-patch",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString(),
			datastarModifiers.DelayMs,
			datastarModifiers.DelaySec,
			datastarModifiers.DebounceMs,
			datastarModifiers.DebounceMsLeading,
			datastarModifiers.DebounceMsNoTrailing,
			datastarModifiers.DebounceSec,
			datastarModifiers.DebounceSecLeading,
			datastarModifiers.DebounceSecNoTrailing,
			datastarModifiers.ThrottleMs,
			datastarModifiers.ThrottleMsNoLeading,
			datastarModifiers.ThrottleMsTrailing,
			datastarModifiers.ThrotlleSec,
			datastarModifiers.ThrotlleSecNoLeading,
			datastarModifiers.ThrotlleSecTrailing,
		),
	},
	{
		Name: "DatastarOnSignalPatchFilter",
		Key:  "on-signal-patch-filter",
		Description: doc(
			"Filters which signals to watch when using the data-on-signal-patch attribute.\n\nThe data-on-signal-patch-filter attribute accepts an object with include and/or exclude properties that are regular expressions.",
			"https://data-star.dev/reference/attributes#data-on-signal-patch-filter",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString()),
	},
	{
		Name: "DatastarPreserveAttr",
		Key:  "preserve-attr",
		Description: doc(
			"Preserves the value of an attribute when morphing DOM elements.",
			"https://data-star.dev/reference/attributes#data-preserve-attr",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString()),
	},
	{
		Name: "DatastarRef",
		Key:  "ref",
		Description: doc(
			"Creates a new signal that is a reference to the element on which the data attribute is placed.",
			"https://data-star.dev/reference/attributes#data-ref",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString(),
			datastarModifiers.Case,
		),
	},
	{
		Name: "DatastarShow",
		Key:  "show",
		Description: doc("Shows or hides an element based on whether an expression evaluates to 'true' or 'false'. For anything with custom requirements, use 'data-class' instead.",
			"https://data-star.dev/reference/attributes#data-show",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString()),
	},
	{
		Name: "DatastarSignals",
		Key:  "signals",
		Description: doc(
			"Patches (adds, updates or removes) one or more signals into the existing signals. Values defined later in the DOM tree override those defined earlier.",
			"https://data-star.dev/reference/attributes#data-signals",
		),
		Type: AttributeTypeCustom(true, AttributeTypeString(),
			datastarModifiers.Case,
			datastarModifiers.IfMissing,
		),
	},
	{
		Name: "DatastarStyle",
		Key:  "datastar-style",
		Description: doc(
			"Sets the value of inline CSS styles on an element based on an expression, and keeps them in sync.",
			"https://data-star.dev/reference/attributes#data-style",
		),
		Type: AttributeTypeCustom(true, AttributeTypeString(),
			datastarModifiers.Case,
		),
	},
	{
		Name: "DatastarText",
		Key:  "text",
		Description: doc(
			"Binds the text content of an element to an expression.",
			"https://data-star.dev/reference/attributes#data-text",
		),
		Type: AttributeTypeCustom(false, AttributeTypeString()),
	},
}

func doc(description string, u_ string) string {
	// Check if documentation url is valid
	u, err := url.Parse(u_)
	if err != nil {
		desc := description[0:30] + "..."
		panic(desc + ": invalid URL")
	}

	resp, err := http.Get(u.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Failed to fetch url %s (statusCode=%d)", u.String(), resp.StatusCode)
		panic(err)
	}

	return description + "\n\nSee: " + u.String()
}

type attributeCustomModifiers struct {
	Capture               *pb.Attribute_Custom_Modifier
	Case                  *pb.Attribute_Custom_Modifier
	DebounceMs            *pb.Attribute_Custom_Modifier
	DebounceMsLeading     *pb.Attribute_Custom_Modifier
	DebounceMsNoTrailing  *pb.Attribute_Custom_Modifier
	DebounceSec           *pb.Attribute_Custom_Modifier
	DebounceSecLeading    *pb.Attribute_Custom_Modifier
	DebounceSecNoTrailing *pb.Attribute_Custom_Modifier
	DelayMs               *pb.Attribute_Custom_Modifier
	DelaySec              *pb.Attribute_Custom_Modifier
	DurationMs            *pb.Attribute_Custom_Modifier
	DurationMsLeading     *pb.Attribute_Custom_Modifier
	DurationSec           *pb.Attribute_Custom_Modifier
	DurationSecLeading    *pb.Attribute_Custom_Modifier
	Full                  *pb.Attribute_Custom_Modifier
	IfMissing             *pb.Attribute_Custom_Modifier
	Half                  *pb.Attribute_Custom_Modifier
	Once                  *pb.Attribute_Custom_Modifier
	Outside               *pb.Attribute_Custom_Modifier
	Passive               *pb.Attribute_Custom_Modifier
	Prevent               *pb.Attribute_Custom_Modifier
	Self                  *pb.Attribute_Custom_Modifier
	Stop                  *pb.Attribute_Custom_Modifier
	Terse                 *pb.Attribute_Custom_Modifier
	ThrottleMs            *pb.Attribute_Custom_Modifier
	ThrottleMsNoLeading   *pb.Attribute_Custom_Modifier
	ThrottleMsTrailing    *pb.Attribute_Custom_Modifier
	ThrotlleSec           *pb.Attribute_Custom_Modifier
	ThrotlleSecNoLeading  *pb.Attribute_Custom_Modifier
	ThrotlleSecTrailing   *pb.Attribute_Custom_Modifier
	ViewTransition        *pb.Attribute_Custom_Modifier
	Window                *pb.Attribute_Custom_Modifier
}

// Check if a modifier was not initialized
func (a attributeCustomModifiers) validate() attributeCustomModifiers {
	value := reflect.ValueOf(a)
	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		fieldName := typ.Field(i).Name
		v := value.Field(i)
		if v.IsNil() {
			panic("Modifier " + fieldName + " is nil")
		}
	}

	return a
}

var datastarModifiers = attributeCustomModifiers{
	Capture: &pb.Attribute_Custom_Modifier{
		Name:        "Capture",
		Description: "Use capture event listener. Only works with built-in events.",
		Type:        AttributeTypeCustomModifier("capture", false, AttributeTypeBool()),
		Prefix:      "capture",
	},
	Case: &pb.Attribute_Custom_Modifier{
		Name:        "Case",
		Description: "Converts the casing of the signal name.\n\t- 'camel' – Camel case: 'mySignal' (default)\n\t- 'kebab' – Kebab case: 'my-signal'\n\t- 'snake' – Snake case: 'my_signal'\n\t- 'pascal' – Pascal case: 'MySignal'",
		Type:        AttributeTypeCustomModifier("case", false, AttributeTypeString()),
		Prefix:      "case.",
	},
	DebounceMs: &pb.Attribute_Custom_Modifier{
		Name:        "DebounceMs",
		Description: "Debounces the event handler",
		Type:        AttributeTypeCustomModifier("debounce", false, AttributeTypeDurationMs()),
		Prefix:      "debounce.",
		Suffix:      "ms",
	},
	DebounceMsLeading: &pb.Attribute_Custom_Modifier{
		Name:        "DebounceMsLeading",
		Description: "Debounce the event listener in milliseconds with leading edge.",
		Type:        AttributeTypeCustomModifier("debounce", false, AttributeTypeDurationMs()),
		Prefix:      "debounce.",
		Suffix:      "ms.leading",
	},
	DebounceMsNoTrailing: &pb.Attribute_Custom_Modifier{
		Name:        "DebounceMsNoTrailing",
		Description: "Debounce the event listener in milliseconds without trailing edge.",
		Type:        AttributeTypeCustomModifier("debounce", false, AttributeTypeDurationMs()),
		Prefix:      "debounce.",
		Suffix:      "ms.notrailing",
	},
	DebounceSec: &pb.Attribute_Custom_Modifier{
		Name:        "DebounceSec",
		Description: "Debounces the event handler",
		Type:        AttributeTypeCustomModifier("debounce", false, AttributeTypeDurationSec()),
		Prefix:      "debounce.",
		Suffix:      "s",
	},
	DebounceSecLeading: &pb.Attribute_Custom_Modifier{
		Name:        "DebounceSecLeading",
		Description: "Debounce the event listener in seconds with leading edge.",
		Type:        AttributeTypeCustomModifier("debounce", false, AttributeTypeDurationSec()),
		Prefix:      "debounce.",
		Suffix:      "s.leading",
	},
	DebounceSecNoTrailing: &pb.Attribute_Custom_Modifier{
		Name:        "DebounceSecNoTrailing",
		Description: "Debounce the event listener in seconds without trailing edge.",
		Type:        AttributeTypeCustomModifier("debounce", false, AttributeTypeDurationSec()),
		Prefix:      "debounce.",
		Suffix:      "s.notrailing",
	},
	DelayMs: &pb.Attribute_Custom_Modifier{
		Name:        "DelayMs",
		Description: "Delay the event listener in milliseconds.",
		Type:        AttributeTypeCustomModifier("delay", false, AttributeTypeDurationMs()),
		Prefix:      "delay.",
		Suffix:      "ms",
	},
	DelaySec: &pb.Attribute_Custom_Modifier{
		Name:        "DelaySec",
		Description: "Delay the event listener in seconds.",
		Type:        AttributeTypeCustomModifier("delay", false, AttributeTypeDurationSec()),
		Prefix:      "delay.",
		Suffix:      "s",
	},
	DurationMs: &pb.Attribute_Custom_Modifier{
		Name:        "DurationMs",
		Description: "Sets the interval duration in milliseconds.",
		Type:        AttributeTypeCustomModifier("duration", false, AttributeTypeDurationMs()),
		Prefix:      "duration.",
		Suffix:      "ms",
	},
	DurationMsLeading: &pb.Attribute_Custom_Modifier{
		Name:        "DurationMsLeading",
		Description: "Sets the interval duration in milliseconds. Execute the first interval immediately.",
		Type:        AttributeTypeCustomModifier("duration", false, AttributeTypeDurationMs()),
		Prefix:      "duration.",
		Suffix:      "ms.leading",
	},
	DurationSec: &pb.Attribute_Custom_Modifier{
		Name:        "DurationSec",
		Description: "Sets the interval duration in seconds.",
		Type:        AttributeTypeCustomModifier("duration", false, AttributeTypeDurationSec()),
		Prefix:      "duration.",
		Suffix:      "s",
	},
	DurationSecLeading: &pb.Attribute_Custom_Modifier{
		Name:        "DurationSecLeading",
		Description: "Sets the interval duration in seconds. Execute the first interval immediately.",
		Type:        AttributeTypeCustomModifier("duration", false, AttributeTypeDurationSec()),
		Prefix:      "duration.",
		Suffix:      "s.leading",
	},
	Full: &pb.Attribute_Custom_Modifier{
		Name:        "Full",
		Description: "Trigger when the full element is visible.",
		Type:        AttributeTypeCustomModifier("full", false, AttributeTypeBool()),
		Prefix:      "full",
	},
	IfMissing: &pb.Attribute_Custom_Modifier{
		Name:        "IfMissing",
		Description: "Only patches signals if their keys do not already exist. This is useful for setting defaults without overwriting existing values.",
		Type:        AttributeTypeCustomModifier("ifmissing", false, AttributeTypeBool()),
		Prefix:      "ifmissing",
	},
	Half: &pb.Attribute_Custom_Modifier{
		Name:        "Half",
		Description: "Trigger when half of the element is visible.",
		Type:        AttributeTypeCustomModifier("half", false, AttributeTypeBool()),
		Prefix:      "half",
	},
	Once: &pb.Attribute_Custom_Modifier{
		Name:        "Once",
		Description: "Only run the expression once. Only works with built-in events.",
		Type:        AttributeTypeCustomModifier("once", false, AttributeTypeBool()),
		Prefix:      "once",
	},
	Outside: &pb.Attribute_Custom_Modifier{
		Name:        "Outside",
		Description: "Triggers when the event is outside the element.",
		Type:        AttributeTypeCustomModifier("outside", false, AttributeTypeBool()),
		Prefix:      "outside",
	},
	Passive: &pb.Attribute_Custom_Modifier{
		Name:        "Passive",
		Description: "Do not call preventDefault on the event listener. Only works with built-in events.",
		Type:        AttributeTypeCustomModifier("passive", false, AttributeTypeBool()),
		Prefix:      "passive",
	},
	Prevent: &pb.Attribute_Custom_Modifier{
		Name:        "Prevent",
		Description: "Calls 'preventDefault' on the event listener.",
		Type:        AttributeTypeCustomModifier("prevent", false, AttributeTypeBool()),
		Prefix:      "prevent",
	},
	Self: &pb.Attribute_Custom_Modifier{
		Name:        "Self",
		Description: "Only ignore the element itself, not its descendants.",
		Type:        AttributeTypeCustomModifier("self", false, AttributeTypeBool()),
		Prefix:      "self",
	},
	Stop: &pb.Attribute_Custom_Modifier{
		Name:        "Stop",
		Description: "Calls 'stopPropagation' on the event listener.",
		Type:        AttributeTypeCustomModifier("stop", false, AttributeTypeBool()),
		Prefix:      "stop",
	},
	Terse: &pb.Attribute_Custom_Modifier{
		Name:        "Terse",
		Description: "Outputs a more compact JSON format without extra whitespace. Useful for displaying filtered data inline.",
		Type:        AttributeTypeCustomModifier("terse", false, AttributeTypeBool()),
		Prefix:      "terse",
	},
	ThrottleMs: &pb.Attribute_Custom_Modifier{
		Name:        "ThrottleMs",
		Description: "Throttles the event handler",
		Type:        AttributeTypeCustomModifier("throttleMs", false, AttributeTypeDurationMs()),
		Prefix:      "throttle.",
		Suffix:      "ms",
	},
	ThrottleMsNoLeading: &pb.Attribute_Custom_Modifier{
		Name:        "ThrottleMsNoLeading",
		Description: "Throttle the event listener in milliseconds without leading edge.",
		Type:        AttributeTypeCustomModifier("throttle", false, AttributeTypeDurationMs()),
		Prefix:      "throttle.",
		Suffix:      "ms.noleading",
	},
	ThrottleMsTrailing: &pb.Attribute_Custom_Modifier{
		Name:        "ThrottleMsTrailing",
		Description: "Throttle the event listener in milliseconds with trailing edge.",
		Type:        AttributeTypeCustomModifier("throttle", false, AttributeTypeDurationMs()),
		Prefix:      "throttle.",
		Suffix:      "ms.trailing",
	},
	ThrotlleSec: &pb.Attribute_Custom_Modifier{
		Name:        "ThrottleSec",
		Description: "Throttles the event listener in seconds.",
		Type:        AttributeTypeCustomModifier("throtlleSec", false, AttributeTypeDurationSec()),
		Prefix:      "throttle.",
		Suffix:      "s",
	},
	ThrotlleSecNoLeading: &pb.Attribute_Custom_Modifier{
		Name:        "ThrottleSecNoLeading",
		Description: "Throttle the event listener in seconds without leading edge.",
		Type:        AttributeTypeCustomModifier("throttle", false, AttributeTypeDurationSec()),
		Prefix:      "throttle.",
		Suffix:      "s.noleading",
	},
	ThrotlleSecTrailing: &pb.Attribute_Custom_Modifier{
		Name:        "ThrottleSecTrailing",
		Description: "Throttle the event listener in seconds with trailing edge.",
		Type:        AttributeTypeCustomModifier("throttle", false, AttributeTypeDurationSec()),
		Prefix:      "throttle.",
		Suffix:      "s.trailing",
	},
	ViewTransition: &pb.Attribute_Custom_Modifier{
		Name:        "ViewTransition",
		Description: "Wraps the expression in 'document.startViewTransition()' when the View Transition API is available.",
		Type:        AttributeTypeCustomModifier("viewtransition", false, AttributeTypeBool()),
		Prefix:      "viewtransition",
	},
	Window: &pb.Attribute_Custom_Modifier{
		Name:        "Window",
		Description: "Attaches the event listener to the 'window' element.",
		Type:        AttributeTypeCustomModifier("window", false, AttributeTypeBool()),
		Prefix:      "window",
	},
}.validate()
