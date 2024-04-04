package cfg

import pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"

var DatastarExtensions = []*pb.Attribute{
	// Core
	{
		Name:        "DatastaStore",
		Key:         "store",
		Description: "Merges the singleton store with the given object",
		Type:        AttributeTypeCustom(false, AttributeTypeJSON()),
	},
	{
		Name:        "DatastarRef",
		Key:         "ref",
		Description: "Sets the reference of the element",
		Type:        AttributeTypeCustom(false, AttributeTypeString()),
	},

	// Attributes
	{
		Name:        "DatastarBind",
		Key:         "bind",
		Description: "Sets the value of the element",
		Type:        AttributeTypeCustom(true, AttributeTypeString()),
	},
	{
		Name:        "DatastarModel",
		Key:         "model",
		Description: "Sets the value of the element",
		Type:        AttributeTypeCustom(false, AttributeTypeString()),
	},
	{
		Name:        "DatastarText",
		Key:         "text",
		Description: "Sets the textContent of the element",
		Type:        AttributeTypeCustom(false, AttributeTypeString()),
	},
	{
		Name:        "DatastarOn",
		Key:         "on",
		Description: "Sets the event handler of the element",
		Type: AttributeTypeCustom(true, AttributeTypeString(),
			&pb.Attribute_Custom_Modifier{
				Name:        "Debounce",
				Description: "Debounces the event handler",
				Type:        AttributeTypeCustomModifier("debounce", false, AttributeTypeDuration()),
				Prefix:      "debounce_",
				Suffix:      "ms",
			},
			&pb.Attribute_Custom_Modifier{
				Name:        "Throttle",
				Description: "Throttles the event handler",
				Type:        AttributeTypeCustomModifier("throttle", false, AttributeTypeDuration()),
				Prefix:      "throttle_",
				Suffix:      "ms",
			},
		),
	},
	{
		Name:        "DatastarFocus",
		Key:         "focus",
		Description: "Sets the focus of the element",
		Type:        AttributeTypeCustom(true, AttributeTypeBool()),
	},

	// Backend
	{
		Name:        "DatastarHeader",
		Key:         "header",
		Description: "Sets the header of for fetch requests",
		Type:        AttributeTypeCustom(true, AttributeTypeString()),
	},
	{
		Name:        "DatastarFetchIndicator",
		Key:         "fetch-indicator",
		Description: "Sets the indicator selector for fetch requests",
		Type:        AttributeTypeCustom(false, AttributeTypeString()),
	},

	// Visibility
	{
		Name:        "DatastarShow",
		Key:         "show",
		Description: "Sets the visibility of the element",
		Type:        AttributeTypeCustom(false, AttributeTypeBool()),
	},
	{
		Name:        "DatastarIntersects",
		Key:         "intersects",
		Description: "Triggers the callback when the element intersects the viewport",
		Type:        AttributeTypeCustom(false, AttributeTypeString()),
	},
	{
		Name:        "DatastarTeleport",
		Key:         "teleport",
		Description: "Teleports the element to the given selector",
		Type:        AttributeTypeCustom(false, AttributeTypeBool()),
	},
	{
		Name:        "DatastarScrollIntoView",
		Key:         "scroll-into-view",
		Description: "Scrolls the element into view",
		Type:        AttributeTypeCustom(false, AttributeTypeBool()),
	},
	{
		Name:        "DatastarViewTransition",
		Key:         "view-transition",
		Description: "Setup the ViewTransitionAPI for the element",
		Type:        AttributeTypeCustom(false, AttributeTypeString()),
	},
}
