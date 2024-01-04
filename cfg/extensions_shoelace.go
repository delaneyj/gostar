package cfg

import pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"

var ShoelaceExtensions = &pb.Namespace{
	Name:        "Shoelace",
	Description: `Shoelace is a collection of professionally designed, every day UI components built on a framework-agnostic technology. It's a great fit for building internal tools, prototypes, and public web apps.`,
	Prefix:      "SL",
	Elements: []*pb.Element{
		{
			Name:        "Alert",
			Tag:         "sl-alert",
			Description: `Alerts are used to communicate a state that affects a system, feature or page. They should draw attention and provide the user with relevant information.`,
			Attributes: []*pb.Attribute{
				shoelaceVariant,
				shoelaceSize,
				{
					Key:  "open",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "duration",
					Type: AttributeTypeDuration(),
				},
			},
		},
		{
			Name:        "Avatar",
			Tag:         "sl-avatar",
			Description: `Avatars are used to represent a user, project, or other item.`,
			Attributes: []*pb.Attribute{
				{
					Key:  "image",
					Type: AttributeTypeString(),
				},
				{
					Key:  "label",
					Type: AttributeTypeString(),
				},
				{
					Key: "loading",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("lazy", "Lazy"),
					),
				},
				{
					Key:  "initials",
					Type: AttributeTypeString(),
				},
				{
					Key: "shape",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("circle", "Circle"),
						AttributeTypeChoice("rounded", "Rounded"),
						AttributeTypeChoice("square", "Square"),
					),
				},
			},
		},
		{
			Name: "Breadcrumb",
			Tag:  "sl-breadcrumb",
		},
		{
			Name: "BreadcrumbItem",
			Tag:  "sl-breadcrumb-item",
			Attributes: []*pb.Attribute{
				{
					Key:  "href",
					Type: AttributeTypeString(),
				},
			},
		},
		{
			Name: "Button",
			Tag:  "sl-button",
			Attributes: []*pb.Attribute{
				shoelaceVariant,
				shoelaceSize,

				{
					Key:  "outline",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "pill",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "circle",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "href",
					Type: AttributeTypeString(),
				},
				{
					Key:  "caret",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "loading",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "disabled",
					Type: AttributeTypeBool(),
				},
			},
		},
		{
			Name: "ButtonGroup",
			Tag:  "sl-button-group",
			Attributes: []*pb.Attribute{
				{
					Key:  "label",
					Type: AttributeTypeString(),
				},
			},
		},
		{
			Name: "Card",
			Tag:  "sl-card",
		},
		{
			Name: "Checkbox",
			Tag:  "sl-checkbox",
			Attributes: []*pb.Attribute{
				shoelaceSize,
				{
					Key:  "name",
					Type: AttributeTypeString(),
				},
				{
					Key:  "disabled",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "indeterminate",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "checked",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "required",
					Type: AttributeTypeBool(),
				},
			},
		},
		{
			Name: "Input",
			Tag:  "sl-input",
			Attributes: []*pb.Attribute{
				shoelaceSize,
				{
					Key: "type",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("date", "Date"),
						AttributeTypeChoice("datetime-local", "Datetime Local"),
						AttributeTypeChoice("email", "Email"),
						AttributeTypeChoice("number", "Number"),
						AttributeTypeChoice("password", "Password"),
						AttributeTypeChoice("search", "Search"),
						AttributeTypeChoice("tel", "Tel"),
						AttributeTypeChoice("text", "Text"),
						AttributeTypeChoice("time", "Time"),
						AttributeTypeChoice("url", "URL"),
					),
				},
				{
					Key:  "name",
					Type: AttributeTypeString(),
				},
				{
					Key:  "value",
					Type: AttributeTypeString(),
				},
				{
					Key:  "defaultValue",
					Type: AttributeTypeString(),
				},
				{
					Key:  "filled",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "pill",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "label",
					Type: AttributeTypeString(),
				},
				{
					Key:  "help-text",
					Type: AttributeTypeString(),
				},
				{
					Key:  "readonly",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "clearable",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "disabled",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "placeholder",
					Type: AttributeTypeString(),
				},
				{
					Key:  "readonly",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "password-toggle",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "password-visible",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "no-spin-buttons",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "required",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "minlength",
					Type: AttributeTypeInt(),
				},
				{
					Key:  "maxlength",
					Type: AttributeTypeInt(),
				},
				{
					Key:  "min",
					Type: AttributeTypeNumber(),
				},
				{
					Key:  "max",
					Type: AttributeTypeNumber(),
				},
				{
					Key:  "step",
					Type: AttributeTypeNumber(),
				},
				{
					Key: "autocapitalize",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("off", "Off"),
						AttributeTypeChoice("none", "None"),
						AttributeTypeChoice("on", "On"),
						AttributeTypeChoice("sentences", "Sentences"),
						AttributeTypeChoice("words", "Words"),
						AttributeTypeChoice("characters", "Characters"),
					),
				},
				{
					Key: "autocorrect",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("off", "Off"),
						AttributeTypeChoice("none", "None"),
						AttributeTypeChoice("on", "On"),
					),
				},
				{
					Key:  "autocomplete",
					Type: AttributeTypeString(),
				},
				{
					Key:  "enter-key-hint",
					Type: AttributeTypeString(),
				},
				{
					Key:  "spellcheck",
					Type: AttributeTypeBool(),
				},
			},
		},
		{
			Name: "RadioGroup",
			Tag:  "sl-radio-group",
			Attributes: []*pb.Attribute{
				shoelaceSize,
				{
					Key:  "label",
					Type: AttributeTypeString(),
				},
				{
					Key:  "help-text",
					Type: AttributeTypeString(),
				},
				{
					Key:  "name",
					Type: AttributeTypeString(),
				},
				{
					Key:  "value",
					Type: AttributeTypeString(),
				},
				{
					Key:  "required",
					Type: AttributeTypeBool(),
				},
			},
		},
		{
			Name: "Radio",
			Tag:  "sl-radio",
			Attributes: []*pb.Attribute{
				shoelaceSize,
				{
					Key:  "value",
					Type: AttributeTypeString(),
				},
				{
					Key:  "disabled",
					Type: AttributeTypeBool(),
				},
			},
		},
		{
			Name: "RadioButton",
			Tag:  "sl-radio-button",
			Attributes: []*pb.Attribute{
				shoelaceSize,
				{
					Key:  "value",
					Type: AttributeTypeString(),
				},
				{
					Key:  "disabled",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "pill",
					Type: AttributeTypeBool(),
				},
			},
		},
		{
			Name: "Select",
			Tag:  "sl-select",
			Attributes: []*pb.Attribute{
				shoelaceSize,
				{
					Key:  "hoist",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "label",
					Type: AttributeTypeString(),
				},
				{
					Key:  "help-text",
					Type: AttributeTypeString(),
				},
				{
					Key:  "placeholder",
					Type: AttributeTypeString(),
				},
				{
					Key:  "clearable",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "filled",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "disabled",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "multiple",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "value",
					Type: AttributeTypeString(),
				},
				{
					Key:  "value",
					Name: "ValueI",
					Type: AttributeTypeInt(),
				},
				{
					Key: "placement",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("bottom", "Bottom"),
						AttributeTypeChoice("top", "Top"),
					),
				},
			},
		},
		{
			Name: "Option",
			Tag:  "sl-option",
			Attributes: []*pb.Attribute{
				{
					Key:  "value",
					Type: AttributeTypeString(),
				},
				{
					Key:  "value",
					Name: "ValueI",
					Type: AttributeTypeInt(),
				},
			},
		},
		{
			Name: "Textarea",
			Tag:  "sl-textarea",
			Attributes: []*pb.Attribute{
				{
					Key:  "name",
					Type: AttributeTypeString(),
				},
				{
					Key:  "value",
					Type: AttributeTypeString(),
				},
				shoelaceSize,
				{
					Key:  "isFilled",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "label",
					Type: AttributeTypeString(),
				},
				{
					Key:  "help-text",
					Type: AttributeTypeString(),
				},
				{
					Key:  "placeholder",
					Type: AttributeTypeString(),
				},
				{
					Key:  "rows",
					Type: AttributeTypeInt(),
				},
				{
					Key: "resize",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("", "Unset"),
						AttributeTypeChoice("none", "None"),
						AttributeTypeChoice("vertical", "Vertical"),
						AttributeTypeChoice("auto", "Auto"),
					),
				},
				{
					Key:  "disabled",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "readonly",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "required",
					Type: AttributeTypeBool(),
				},
				{
					Key:  "minlength",
					Type: AttributeTypeInt(),
				},
				{
					Key:  "maxlength",
					Type: AttributeTypeInt(),
				},
				{
					Key: "autocapitalize",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("off", "Off"),
						AttributeTypeChoice("none", "None"),
						AttributeTypeChoice("on", "On"),
						AttributeTypeChoice("sentences", "Sentences"),
						AttributeTypeChoice("words", "Words"),
						AttributeTypeChoice("characters", "Characters"),
					),
				},
				{
					Key: "autocorrect",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("off", "Off"),
						AttributeTypeChoice("none", "None"),
						AttributeTypeChoice("on", "On"),
					),
				},
				{
					Key:  "autocomplete",
					Type: AttributeTypeString(),
				},
				{
					Key:  "spellcheck",
					Type: AttributeTypeBool(),
				},
				{
					Key: "inputmode",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("", "Unset"),
						AttributeTypeChoice("none", "None"),
						AttributeTypeChoice("text", "Text"),
						AttributeTypeChoice("decimal", "Decimal"),
						AttributeTypeChoice("numeric", "Numeric"),
						AttributeTypeChoice("tel", "Tel"),
						AttributeTypeChoice("search", "Search"),
						AttributeTypeChoice("email", "Email"),
						AttributeTypeChoice("url", "URL"),
					),
				},
				{
					Key:  "defaultValue",
					Type: AttributeTypeString(),
				},
			},
		},
		{
			Name: "Tooltip",
			Tag:  "sl-tooltip",
			Attributes: []*pb.Attribute{
				{
					Key:  "content",
					Type: AttributeTypeString(),
				},
				{
					Key: "placement",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("top", "Top"),
						AttributeTypeChoice("top-start", "Top Start"),
						AttributeTypeChoice("top-end", "Top End"),
						AttributeTypeChoice("right", "Right"),
						AttributeTypeChoice("right-start", "Right Start"),
						AttributeTypeChoice("right-end", "Right End"),
						AttributeTypeChoice("bottom", "Bottom"),
						AttributeTypeChoice("bottom-start", "Bottom Start"),
						AttributeTypeChoice("bottom-end", "Bottom End"),
						AttributeTypeChoice("left", "Left"),
						AttributeTypeChoice("left-start", "Left Start"),
						AttributeTypeChoice("left-end", "Left End"),
					),
				},
			},
		},
	},
}

var (
	shoelaceVariant = &pb.Attribute{
		Key: "variant",
		Type: AttributeTypeChoices(
			AttributeTypeChoice("default", "Default"),
			AttributeTypeChoice("primary", "Primary"),
			AttributeTypeChoice("success", "Success"),
			AttributeTypeChoice("info", "Info"),
			AttributeTypeChoice("warning", "Warning"),
			AttributeTypeChoice("danger", "Danger"),
			AttributeTypeChoice("text", "Text"),
		),
	}
	shoelaceSize = &pb.Attribute{
		Key: "size",
		Type: AttributeTypeChoices(
			AttributeTypeChoice("small", "Small"),
			AttributeTypeChoice("medium", "Medium"),
			AttributeTypeChoice("large", "Large"),
		),
	}
)
