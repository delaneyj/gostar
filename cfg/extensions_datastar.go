package cfg

import pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"

var DatastarExtensions = []*pb.Attribute{
	// Core
	{
		Name:        "DatastarMergeStore",
		Key:         "data-merge-store",
		Description: "Merges the store with the given object",
		Type:        AttributeTypeCustom("merge-store", false, AttributeTypeJSON()),
	},
	{
		Name:        "DatastarRef",
		Key:         "data-ref",
		Description: "Sets the reference of the element",
		Type:        AttributeTypeCustom("ref", false, AttributeTypeString()),
	},

	// Attributes
	{
		Name:        "DatastarBind",
		Key:         "data-bind",
		Description: "Sets the value of the element",
		Type:        AttributeTypeCustom("bind", true, AttributeTypeString()),
	},
	{
		Name:        "DatastarModel",
		Key:         "data-model",
		Description: "Sets the value of the element",
		Type:        AttributeTypeCustom("model", false, AttributeTypeString()),
	},
	{
		Name:        "DatastarText",
		Key:         "data-text",
		Description: "Sets the textContent of the element",
		Type:        AttributeTypeCustom("text", false, AttributeTypeString()),
	},
	{
		Name:        "DatastarOn",
		Key:         "data-on",
		Description: "Sets the event handler of the element",
		Type: AttributeTypeCustom("on", true, AttributeTypeString(),
			&pb.Attribute_Custom_Modifier{
				Name:        "Debounce",
				Description: "Debounces the event handler",
				Type:        AttributeTypeCustom("debounce", false, AttributeTypeDuration()),
				Prefix:      "debounce_",
				Suffix:      "ms",
			},
			&pb.Attribute_Custom_Modifier{
				Name:        "Throttle",
				Description: "Throttles the event handler",
				Type:        AttributeTypeCustom("throttle", false, AttributeTypeDuration()),
				Prefix:      "throttle_",
				Suffix:      "ms",
			},
		),
	},
	{
		Name:        "DatastarFocus",
		Key:         "data-focus",
		Description: "Sets the focus of the element",
		Type:        AttributeTypeCustom("focus", true, AttributeTypeBool()),
	},

	// Backend
	{
		Name:        "DatastarHeader",
		Key:         "data-header",
		Description: "Sets the header of for fetch requests",
		Type:        AttributeTypeCustom("header", true, AttributeTypeString()),
	},
	{
		Name:        "DatastarFetchURL",
		Key:         "data-fetch-url",
		Description: "Sets the URL for fetch requests",
		Type:        AttributeTypeCustom("fetch-url", false, AttributeTypeString()),
	},
	{
		Name:        "DatastarFetchIndicator",
		Description: "Sets the indicator selector for fetch requests",
		Type:        AttributeTypeCustom("fetch-indicator", false, AttributeTypeString()),
	},

	// Visibility
	{
		Name:        "DatastarShow",
		Key:         "data-show",
		Description: "Sets the visibility of the element",
		Type:        AttributeTypeCustom("show", false, AttributeTypeBool()),
	},
	{
		Name:        "DatastarIntersects",
		Key:         "data-intersects",
		Description: "Triggers the callback when the element intersects the viewport",
		Type:        AttributeTypeCustom("intersects", false, AttributeTypeBool()),
	},
	{
		Name:        "DatastarTeleport",
		Key:         "data-teleport",
		Description: "Teleports the element to the given selector",
		Type:        AttributeTypeCustom("teleport", false, AttributeTypeBool()),
	},
	{
		Name:        "DatastarScrollIntoView",
		Key:         "data-scroll-into-view",
		Description: "Scrolls the element into view",
		Type:        AttributeTypeCustom("scroll-into-view", false, AttributeTypeBool()),
	},
	{
		Name:        "DatastarViewTransition",
		Key:         "data-view-transition",
		Description: "Setup the ViewTransitionAPI for the element",
		Type:        AttributeTypeCustom("view-transition", true, AttributeTypeString()),
	},
}
