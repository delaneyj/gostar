package cfg

import pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"

var DatastarExtensions = []*pb.Attribute{
	// Core
	{
		Name: 	  "DatastarMergeStore",
		Description: "Merges the store with the given object",
		Type:        AttributeTypeCustom("merge-store", AttributeTypeJSON()),
	},
	{
		Name:        "DatastarRef",
		Description: "Sets the reference of the element",
		Type:        AttributeTypeCustom("ref", AttributeTypeString()),
	},

	// Attributes
	{
		Name:        "DatastarBind",
		Description: "Sets the value of the element",
		Type:        AttributeTypeCustom("bind", AttributeTypeString()),
	},
	{
		Name:        "DatastarModel",
		Description: "Sets the value of the element",
		Type:        AttributeTypeCustom("model", AttributeTypeString()),
	},
	{
		Name:        "DatastarText",
		Description: "Sets the textContent of the element",
		Type:        AttributeTypeCustom("text", AttributeTypeString()),
	},
	{
		Name:        "DatastarOn",
		Description: "Sets the event handler of the element",
		Type:        AttributeTypeCustom("on", AttributeTypeString()),
	},
	{
		Name:        "DatastarFocus",
		Description: "Sets the focus of the element",
		Key:         "data-focus",
		Type:        AttributeTypeCustom("focus", AttributeTypeBool()),
	},

	// Backend
	{
		Name:        "DatastarHeader",
		Description: "Sets the header of for fetch requests",
		Type:        AttributeTypeCustom("header", AttributeTypeString()),
	},
	{
		Name:        "DatastarFetchURL",
		Description: "Sets the URL for fetch requests",
		Type:        AttributeTypeCustom("fetch-url", AttributeTypeString()),
	},
	{
		Name:        "DatastarFetchIndicator",
		Description: "Sets the indicator selector for fetch requests",
		Type:        AttributeTypeCustom("fetch-indicator", AttributeTypeString()),
	},

	// Visibility
	{
		Name:        "DatastarShow",
		Description: "Sets the visibility of the element",
		Type:        AttributeTypeCustom("show", AttributeTypeBool()),
	},
	{
		Name:        "DatastarIntersects",
		Description: "Triggers the callback when the element intersects the viewport",
		Type:        AttributeTypeCustom("intersects", AttributeTypeBool()),
	},
	{
		Name:        "DatastarTeleport",
		Description: "Teleports the element to the given selector",
		Type:        AttributeTypeCustom("teleport", AttributeTypeBool()),
	},
	{
		Name:        "DatastarScrollIntoView",
		Description: "Scrolls the element into view",
		Type:        AttributeTypeCustom("scroll-into-view", AttributeTypeBool()),
	},
	{
		Name:        "DatastarViewTransition",
		Description: "Setup the ViewTransitionAPI for the element",
		Type:        AttributeTypeCustom("view-transition", AttributeTypeString()),
	},
}
