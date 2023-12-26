package cfg

import (
	pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"
)

var SVG = &pb.Namespace{
	Name:        "svg",
	Description: `Scalable Vector Graphics (SVG) is an XML-based markup language for describing two-dimensional based vector graphics. As such, it's a text-based, open Web standard for describing images that can be rendered cleanly at any size and are designed specifically to work well with other web standards including CSS, DOM, JavaScript, and SMIL. SVG is, essentially, to graphics what HTML is to text. SVG images and their related behaviors are defined in XML text files, which means they can be searched, indexed, scripted, and compressed. Additionally, this means they can be created and edited with any text editor or with drawing software. Compared to classic bitmapped image formats such as JPEG or PNG, SVG-format vector images can be rendered at any size without loss of quality and can be easily localized by updating the text within them, without the need of a graphical editor to do so. With proper libraries, SVG files can even be localized on-the-fly.`,
	Prefix:      "SVG",
	GlobalAttributes: []*pb.Attribute{
		{
			Key:         "id",
			Description: "Specifies a unique id for an element",
			Type:        AttributeTypeString(),
		},
		{
			Key:         "class",
			Description: "Specifies one or more classnames for an element (refers to a class in a style sheet)",
			Type:        AttributeTypeSpaceDelimited(),
		},
		{
			Key:         "style",
			Description: "Specifies an inline CSS style for an element",
			Type:        AttributeTypeKVColonSemicolon(),
		},
	},
	Elements: []*pb.Element{
		{
			Tag:         "a",
			Description: `The <a> SVG element creates a hyperlink to other web pages, files, locations in the same page, email addresses, or any other URL. It is very similar to HTML's <a> element. SVG's <a> element is a container, which means you can create a link around text (like in HTML) but also around any shape.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "download",
					Description: `Indicates that the hyperlink is to be used for downloading a resource. When used together with the download attribute, the value of the attribute is used as the file name of the downloaded file. There are no restrictions on allowed values, though / and \ will be converted to underscores and leading spaces in filenames will be removed.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "href",
					Description: `The URL of a linked resource.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "hreflang",
					Description: `Specifies the language of the linked resource.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "ping",
					Description: `A space-separated list of URLs. When the link is followed, the browser should send POST requests with the body PING to the URLs. Typically for tracking.`,
					Type:        AttributeTypeSpaceDelimited(),
				},
				{
					Key:         "referrerpolicy",
					Description: `Referrer policy to use when fetching the resource.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("no-referrer", `The Referer header will not be sent.`),
						AttributeTypeChoice("no-referrer-when-downgrade", `The Referer header will not be sent to origins without TLS (HTTPS).`),
						AttributeTypeChoice("origin", `The Referer header will contain the origin of the request.`),
						AttributeTypeChoice("origin-when-cross-origin", `The Referer header will contain the origin of the request, unless it is a cross-origin request, in which case it will be omitted entirely.`),
						AttributeTypeChoice("same-origin", `The Referer header will contain the origin of the request, unless it is a cross-origin request, in which case it will be omitted entirely.`),
						AttributeTypeChoice("strict-origin", `The Referer header will contain the origin of the request, unless it is a cross-origin request, in which case it will be omitted entirely.`),
						AttributeTypeChoice("strict-origin-when-cross-origin", `The Referer header will contain the origin of the request, unless it is a cross-origin request, in which case it will be omitted entirely.`),
						AttributeTypeChoice("unsafe-url", `The Referer header will contain the origin of the request, unless it is a cross-origin request, in which case it will be omitted entirely.`),
					),
				},
				{
					Key:         "rel",
					Description: `Specifies the relationship of the target object to the link object.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("alternate", `Links to an alternate version of the document (i.e. print page, translated or mirror).`),
						AttributeTypeChoice("author", `Links to the author of the document.`),
						AttributeTypeChoice("bookmark", `Permanent URL used for bookmarking.`),
						AttributeTypeChoice("canonical", `Helps prevent duplicate content issues.`),
						AttributeTypeChoice("dns-prefetch", `Specifies that the browser should preemptively perform DNS resolution for the target resource's origin.`),
						AttributeTypeChoice("external", `Links to an external resource (an external stylesheet).`),
						AttributeTypeChoice("help", `Links to a help document.`),
						AttributeTypeChoice("icon", `Imports an icon to represent the document.`),
						AttributeTypeChoice("license", `Links to a license associated with the document.`),
						AttributeTypeChoice("manifest", `Specifies the location of a manifest file or an entry point for a web application.`),
						AttributeTypeChoice("me", `Links to a resource that is the primary topic of the document.`),
						AttributeTypeChoice("modulepreload", `Specifies that the target resource should be preemptively fetched and cached by the browser for later use.`),
						AttributeTypeChoice("next", `The next document in a selection.`),
						AttributeTypeChoice("nofollow", `Links to an unendorsed document, like a paid link. ("nofollow" is used by Google, to specify that the Google search spider should not follow that link.)`),
						AttributeTypeChoice("noopener", `Specifies that the browser should not open the linked document in a new tab or window.`),
						AttributeTypeChoice("noreferrer", `Specifies that the browser should not send a HTTP referer header if the user follows the hyperlink.`),
						AttributeTypeChoice("opener", `Specifies that the target URL should be opened in a top-level browsing context (that is, in the current tab or window).`),
						AttributeTypeChoice("pingback", `Links to the Pingback server of the current document.`),
						AttributeTypeChoice("preload", `Specifies that the target resource should be loaded immediately.`),
						AttributeTypeChoice("prev", `The previous document in a selection.`),
						AttributeTypeChoice("privacy-policy", `Links to a privacy policy document associated with the current document.`),
						AttributeTypeChoice("search", `Links to a search tool for the document.`),
						AttributeTypeChoice("stylesheet", `Links to an external style sheet.`),
						AttributeTypeChoice("tag", `A tag (keyword) for the current document.`),
						AttributeTypeChoice("terms-of-service", `Links to the terms of service document for the current document.`),
					),
				},
				{
					Key:         "target",
					Description: `Specifies where to display the linked resource.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("_self", `Default. Opens the document in the same frame as it was clicked.`),
						AttributeTypeChoice("_blank", `Opens the document in a new window or tab.`),
						AttributeTypeChoice("_parent", `Opens the document in the parent frame.`),
						AttributeTypeChoice("_top", `Opens the document in the full body of the window.`),
						AttributeTypeChoice("framename", `Opens the document in a named frame.`),
					),
				},
				{
					Key:         "type",
					Description: `Specifies the MIME type of the linked resource.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "animate",
			Description: `The <animate> SVG element is used to animate an attribute or property of an element over time.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "accumulate",
					Description: `Controls whether or not the animation is cumulative.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("none", `The animation is not cumulative. Each iteration starts over from the beginning.`),
						AttributeTypeChoice("sum", `The animation is cumulative. Each iteration the animation picks up where it left off in the previous iteration.`),
					),
				},
				{
					Key:         "additive",
					Description: `Controls whether or not the animation is additive.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("replace", `The animation is not additive. The animation replaces the underlying value.`),
						AttributeTypeChoice("sum", `The animation is additive. The animation adds to the underlying value.`),
					),
				},
				{
					Key:         "attributeName",
					Description: `The name of the attribute to animate.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "attributeType",
					Description: `The namespace of the attribute to animate.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("auto", `If the attribute is a presentation attribute, the animation will use the target element's corresponding baseVal. If the attribute is not a presentation attribute, the animation will use the target element's corresponding animVal.`),
						AttributeTypeChoice("CSS", `The animation will use the CSS namespace.`),
						AttributeTypeChoice("XML", `The animation will use the XML namespace.`),
						AttributeTypeChoice("XMLID", `The animation will use the XML ID namespace.`),
						AttributeTypeChoice("XMLLANG", `The animation will use the XML LANG namespace.`),
						AttributeTypeChoice("XMLSPACE", `The animation will use the XML SPACE namespace.`),
					),
				},
				{
					Key:         "begin",
					Description: `Defines when the animation should begin.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "by",
					Description: `Defines a relative offset value for the animation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "calcMode",
					Description: `Defines the pacing of the animation.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("discrete", `The animation is not paced. Each iteration of the animation is displayed as fast as possible.`),
						AttributeTypeChoice("linear", `The animation is paced such that it takes the same amount of time to go from the start value to the end value throughout the animation.`),
						AttributeTypeChoice("paced", `The animation is paced according to a cubic function.`),
						AttributeTypeChoice("spline", `The animation is paced according to a cubic function, but with easing at both the start and end.`),
					),
				},
				{
					Key:         "dur",
					Description: `Defines the duration of the animation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "end",
					Description: `Defines when the animation should end.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "fill",
					Description: `Defines the fill behavior for the animation.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("freeze", `The animation will hold the attribute value when the animation ends.`),
						AttributeTypeChoice("remove", `The animation will remove the attribute value when the animation ends.`),
					),
				},
				{
					Key:         "from",
					Description: `Defines the initial value of the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "keySplines",
					Description: `Defines the values for a cubic Bézier function that controls interval pacing.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "keyTimes",
					Description: `Defines when the animation should take place in terms of time fractions.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "max",
					Description: `Defines the maximum value allowed for the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "min",
					Description: `Defines the minimum value allowed for the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "repeatCount",
					Description: `Defines the number of times the animation should repeat.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "repeatDur",
					Description: `Defines the duration for repeating an animation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "restart",
					Description: `Defines if an animation should restart after it completes.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("always", `The animation will restart indefinitely.`),
						AttributeTypeChoice("never", `The animation will not restart after it completes.`),
						AttributeTypeChoice("whenNotActive", `The animation will restart after it completes if the animation is not currently active.`),
					),
				},
				{
					Key:         "to",
					Description: `Defines the ending value of the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "values",
					Description: `Defines a list of discrete values to interpolate.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "animateMotion",
			Description: `The <animateMotion> SVG element is used to animate a transformation attribute on a target element, thereby allowing the animation of translation, rotation, and scaling.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "accumulate",
					Description: `Controls whether or not the animation is cumulative.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("none", `The animation is not cumulative. Each iteration starts over from the beginning.`),
						AttributeTypeChoice("sum", `The animation is cumulative. Each iteration the animation picks up where it left off in the previous iteration.`),
					),
				},
				{
					Key:         "additive",
					Description: `Controls whether or not the animation is additive.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("replace", `The animation is not additive. The animation replaces the underlying value.`),
						AttributeTypeChoice("sum", `The animation is additive. The animation adds to the underlying value.`),
					),
				},
				{
					Key:         "begin",
					Description: `Defines when the animation should begin.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "by",
					Description: `Defines a relative offset value for the animation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "calcMode",
					Description: `Defines the pacing of the animation.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("discrete", `The animation is not paced. Each iteration of the animation is displayed as fast as possible.`),
						AttributeTypeChoice("linear", `The animation is paced such that it takes the same amount of time to go from the start value to the end value throughout the animation.`),
						AttributeTypeChoice("paced", `The animation is paced according to a cubic function.`),
						AttributeTypeChoice("spline", `The animation is paced according to a cubic function, but with easing at both the start and end.`),
					),
				},
				{
					Key:         "dur",
					Description: `Defines the duration of the animation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "end",
					Description: `Defines when the animation should end.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "fill",
					Description: `Defines the fill behavior for the animation.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("freeze", `The animation will hold the attribute value when the animation ends.`),
						AttributeTypeChoice("remove", `The animation will remove the attribute value when the animation ends.`),
					),
				},
				{
					Key:         "from",
					Description: `Defines the initial value of the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "keySplines",
					Description: `Defines the values for a cubic Bézier function that controls interval pacing.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "keyTimes",
					Description: `Defines when the animation should take place in terms of time fractions.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "max",
					Description: `Defines the maximum value allowed for the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "min",
					Description: `Defines the minimum value allowed for the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "repeatCount",
					Description: `Defines the number of times the animation should repeat.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "repeatDur",
					Description: `Defines the duration for repeating an animation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "restart",
					Description: `Defines if an animation should restart after it completes.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("always", `The animation will restart indefinitely.`),
						AttributeTypeChoice("never", `The animation will not restart after it completes.`),
						AttributeTypeChoice("whenNotActive", `The animation will restart after it completes if the animation is not currently active.`),
					),
				},
				{
					Key:         "to",
					Description: `Defines the ending value of the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "values",
					Description: `Defines a list of discrete values to interpolate.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "animateTransform",
			Description: `The <animateTransform> SVG element animates a transformation attribute on a target element, thereby allowing animations to control translation, scaling, rotation and/or skewing.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "accumulate",
					Description: `Controls whether or not the animation is cumulative.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("none", `The animation is not cumulative. Each iteration starts over from the beginning.`),
						AttributeTypeChoice("sum", `The animation is cumulative. Each iteration the animation picks up where it left off in the previous iteration.`),
					),
				},
				{
					Key:         "additive",
					Description: `Controls whether or not the animation is additive.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("replace", `The animation is not additive. The animation replaces the underlying value.`),
						AttributeTypeChoice("sum", `The animation is additive. The animation adds to the underlying value.`),
					),
				},
				{
					Key:         "attributeName",
					Description: `The name of the attribute to animate.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "attributeType",
					Description: `The namespace of the attribute to animate.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("auto", `If the attribute is a presentation attribute, the animation will use the target element's corresponding baseVal. If the attribute is not a presentation attribute, the animation will use the target element's corresponding animVal.`),
						AttributeTypeChoice("CSS", `The animation will use the CSS namespace.`),
						AttributeTypeChoice("XML", `The animation will use the XML namespace.`),
						AttributeTypeChoice("XMLID", `The animation will use the XML ID namespace.`),
						AttributeTypeChoice("XMLLANG", `The animation will use the XML LANG namespace.`),
						AttributeTypeChoice("XMLSPACE", `The animation will use the XML SPACE namespace.`),
					),
				},
				{
					Key:         "begin",
					Description: `Defines when the animation should begin.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "by",
					Description: `Defines a relative offset value for the animation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "calcMode",
					Description: `Defines the pacing of the animation.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("discrete", `The animation is not paced. Each iteration of the animation is displayed as fast as possible.`),
						AttributeTypeChoice("linear", `The animation is paced such that it takes the same amount of time to go from the start value to the end value throughout the animation.`),
						AttributeTypeChoice("paced", `The animation is paced according to a cubic function.`),
						AttributeTypeChoice("spline", `The animation is paced according to a cubic function, but with easing at both the start and end.`),
					),
				},
				{
					Key:         "dur",
					Description: `Defines the duration of the animation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "end",
					Description: `Defines when the animation should end.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "fill",
					Description: `Defines the fill behavior for the animation.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("freeze", `The animation will hold the attribute value when the animation ends.`),
						AttributeTypeChoice("remove", `The animation will remove the attribute value when the animation ends.`),
					),
				},
				{
					Key:         "from",
					Description: `Defines the initial value of the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "keySplines",
					Description: `Defines the values for a cubic Bézier function that controls interval pacing.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "keyTimes",
					Description: `Defines when the animation should take place in terms of time fractions.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "max",
					Description: `Defines the maximum value allowed for the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "min",
					Description: `Defines the minimum value allowed for the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "repeatCount",
					Description: `Defines the number of times the animation should repeat.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "repeatDur",
					Description: `Defines the duration for repeating an animation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "restart",
					Description: `Defines if an animation should restart after it completes.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("always", `The animation will restart indefinitely.`),
						AttributeTypeChoice("never", `The animation will not restart after it completes.`),
						AttributeTypeChoice("whenNotActive", `The animation will restart after it completes if the animation is not currently active.`),
					),
				},
				{
					Key:         "to",
					Description: `Defines the ending value of the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "type",
					Description: `Defines which transform to use.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("rotate", `The animation will use the rotate transform.`),
						AttributeTypeChoice("scale", `The animation will use the scale transform.`),
						AttributeTypeChoice("translate", `The animation will use the translate transform.`),
					),
				},
				{
					Key:         "values",
					Description: `Defines a list of discrete values to interpolate.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "circle",
			Description: `The <circle> SVG element is an SVG basic shape, used to create circles based on a center point and a radius.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "cx",
					Description: `The x-axis coordinate of the center of the circle.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "cy",
					Description: `The y-axis coordinate of the center of the circle.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "r",
					Description: `The radius of the circle.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "clipPath",
			Description: `The <clipPath> SVG element defines a clipping path. A clipping path is used/referenced using the clip-path property.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "clipPathUnits",
					Description: `The coordinate system for the contents of the <clipPath> element.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The contents of the <clipPath> element represent values in the current user coordinate system.`),
						AttributeTypeChoice("objectBoundingBox", `The contents of the <clipPath> element represent values in the coordinate system that results from taking the current user coordinate system in place at the time when the <clipPath> element is referenced (i.e., the user coordinate system for the element referencing the <clipPath> element via a clip-path property).`),
					),
				},
			},
		},

		{
			Tag:         "defs",
			Description: `The <defs> SVG element is used to embed definitions that can be reused inside an SVG image.`,
		},

		{
			Tag:         "desc",
			Description: `The <desc> SVG element provides a description container for SVG content.`,
		},

		{
			Tag:         "ellipse",
			Description: `The <ellipse> SVG element is an SVG basic shape, used to create ellipses based on a center coordinate, and both their x and y radius.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "cx",
					Description: `The x-axis coordinate of the center of the ellipse.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "cy",
					Description: `The y-axis coordinate of the center of the ellipse.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "rx",
					Description: `The x-axis radius of the ellipse.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "ry",
					Description: `The y-axis radius of the ellipse.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feBlend",
			Description: `The <feBlend> SVG filter primitive composes two objects together ruled by a certain blending mode. This is similar to what is known from image editing software when blending two layers. The mode is defined by the mode attribute.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `Input for the blending.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "in2",
					Description: `Second input for the blending.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "mode",
					Description: `The mode used to blend the two inputs together.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("normal", `The input image is placed over the backdrop image, then the parts of the input image that are outside the backdrop are discarded.`),
						AttributeTypeChoice("multiply", `The input image is multiplied by the backdrop image.`),
						AttributeTypeChoice("screen", `Multiplies the complements of the backdrop and input image color values, then complements the result.`),
						AttributeTypeChoice("darken", `Selects the darker of the backdrop and input image pixels.`),
						AttributeTypeChoice("lighten", `Selects the lighter of the backdrop and input image pixels.`),
					),
				},
			},
		},

		{
			Tag:         "feColorMatrix",
			Description: `The <feColorMatrix> SVG filter element changes colors based on a transformation matrix. Every pixel's color value (represented by an [R,G,B,A] vector) is matrix multiplied to create a new color.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The input for this filter.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "type",
					Description: `The type of matrix operation.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("matrix", `The type of matrix operation.`),
						AttributeTypeChoice("saturate", `The type of matrix operation.`),
						AttributeTypeChoice("hueRotate", `The type of matrix operation.`),
						AttributeTypeChoice("luminanceToAlpha", `The type of matrix operation.`),
					),
				},
				{
					Key:         "values",
					Description: `The list of one or more numbers that represent the matrix.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "feComponentTransfer",
			Description: `The <feComponentTransfer> SVG filter primitive performs color-component-wise remapping of data for each pixel. It allows operations like brightness adjustment, contrast adjustment, color balance or thresholding.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The input for this filter.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "feComposite",
			Description: `The <feComposite> SVG filter primitive performs the combination of two input images pixel-wise in image space using one of the Porter-Duff compositing operations: over, in, atop, out, xor.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `Input for the compositing operation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "in2",
					Description: `Second input for the compositing operation.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "operator",
					Description: `The type of compositing operation.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("over", `The source image is composited over the destination image.`),
						AttributeTypeChoice("in", `The part of the source image that lies inside of the destination image is composited over the destination image.`),
						AttributeTypeChoice("out", `The part of the source image that lies outside of the destination image is composited over the destination image.`),
						AttributeTypeChoice("atop", `The part of the source image that lies inside of the destination image is composited over the destination image and replaces the destination image.`),
						AttributeTypeChoice("xor", `The part of the source image that lies outside of the destination image is composited over the destination image.`),
						AttributeTypeChoice("arithmetic", `A standard arithmetic operator is applied (`),
					),
				},
				{
					Key:         "k1",
					Description: `First value to use in the arithmetic operation.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "k2",
					Description: `Second value to use in the arithmetic operation.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "k3",
					Description: `Third value to use in the arithmetic operation.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "k4",
					Description: `Fourth value to use in the arithmetic operation.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feConvolveMatrix",
			Description: `The <feConvolveMatrix> SVG filter primitive applies a matrix convolution filter effect.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The input for this filter.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "order",
					Description: `The number of cells in each dimension for 'kernelMatrix'`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "kernelMatrix",
					Description: `A list of numbers that make up the kernel matrix for the convolution.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "divisor",
					Description: `The divisor attribute specifies the value by which to divide the result of applying the convolution operator.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "bias",
					Description: `The bias attribute shifts the range of the filter. After applying the matrix operation, this bias value is added to each component.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "targetX",
					Description: `The targetX attribute determines the positioning in X of the convolution matrix relative to a given target pixel in the input image.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "targetY",
					Description: `The targetY attribute determines the positioning in Y of the convolution matrix relative to a given target pixel in the input image.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "edgeMode",
					Description: `The edgeMode attribute determines how to extend the input image as necessary with color values so that the matrix operations can be applied when the kernel is positioned at or near the edge of the input image.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("duplicate", `The input image is extended along each of its borders as necessary by duplicating the color values at the given edge of the input image.`),
						AttributeTypeChoice("wrap", `The input image is extended by taking the component values from the opposite edge of the image.`),
						AttributeTypeChoice("none", `Any values outside the input image are assumed to be transparent black.`),
					),
				},
				{
					Key:         "kernelUnitLength",
					Description: `The kernelUnitLength attribute defines the intended distance in current filter units (i.e., units as determined by the value of attribute 'primitiveUnits') for dx and dy in the surface normal calculation formulas.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "preserveAlpha",
					Description: `The preserveAlpha attribute indicates how the convolution will handle the alpha channel of the input image.`,
					Type:        AttributeTypeBool(),
				},
			},
		},

		{
			Tag:         "feDiffuseLighting",
			Description: `The <feDiffuseLighting> SVG filter primitive lights an image using the alpha channel as a bump map. The resulting image, which is an RGBA opaque image, depends on the light color, light position and surface geometry of the input bump map.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The input for this filter.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "surfaceScale",
					Description: `The 'surfaceScale' attribute indicates the height of the surface when the alpha channel is 1.0.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "diffuseConstant",
					Description: `The diffuseConstant attribute represents the proportion of the light that is reflected by the surface.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "kernelUnitLength",
					Description: `The kernelUnitLength attribute defines the intended distance in current filter units (i.e., units as determined by the value of attribute 'primitiveUnits') for dx and dy in the surface normal calculation formulas.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "feDisplacementMap",
			Description: `The <feDisplacementMap> SVG filter primitive uses the pixel values from the image from in2 to spatially displace the image from in.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The input for this filter.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "in2",
					Description: `The displacement map. This attribute can take on the same values as the 'in' attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "scale",
					Description: `The scale attribute defines the maximum value for the in2 displacement. A value of 0 disables the effect of the displacement map.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "xChannelSelector",
					Description: `The xChannelSelector attribute indicates which color channel from in2 to use to displace the pixels in in the horizontal direction.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("R", `The red channel of in2 is used to displace the x coordinate of each pixel.`),
						AttributeTypeChoice("G", `The green channel of in2 is used to displace the x coordinate of each pixel.`),
						AttributeTypeChoice("B", `The blue channel of in2 is used to displace the x coordinate of each pixel.`),
						AttributeTypeChoice("A", `The alpha channel of in2 is used to displace the x coordinate of each pixel.`),
					),
				},
				{
					Key:         "yChannelSelector",
					Description: `The yChannelSelector attribute indicates which color channel from in2 to use to displace the pixels in in the vertical direction.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("R", `The red channel of in2 is used to displace the y coordinate of each pixel.`),
						AttributeTypeChoice("G", `The green channel of in2 is used to displace the y coordinate of each pixel.`),
						AttributeTypeChoice("B", `The blue channel of in2 is used to displace the y coordinate of each pixel.`),
						AttributeTypeChoice("A", `The alpha channel of in2 is used to displace the y coordinate of each pixel.`),
					),
				},
			},
		},

		{
			Tag:         "feDistantLight",
			Description: `The <feDistantLight> SVG filter primitive defines a distant light source that can be used within a lighting filter primitive: <feDiffuseLighting> or <feSpecularLighting>`,
			Attributes: []*pb.Attribute{
				{
					Key:         "azimuth",
					Description: `The azimuth attribute represent the direction vector of the light source in the XY plane (clockwise), in degrees from the x axis.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "elevation",
					Description: `The elevation attribute represent the direction vector of the light source perpendicular to the XY plane, in degrees from the XY plane towards the z axis (clockwise).`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feDropShadow",
			Description: `The <feDropShadow> filter primitive creates a drop shadow of the input image. It is a shorthand filter, and is defined in terms of the <feGaussianBlur> and <feOffset> filter primitives.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "dx",
					Description: `The amount of offset in the x direction. If the <length> is 0, the shadow is placed at the same position as the input.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "dy",
					Description: `The amount of offset in the y direction. If the <length> is 0, the shadow is placed at the same position as the input.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "stdDeviation",
					Description: `The standard deviation for the blur operation. If two <numbers> are provided, the first number represents a standard deviation value along the x-axis of the coordinate system established by attribute 'primitiveUnits' on the <filter> element. The second value represents a standard deviation in Y. If one number is provided, then that value is used for both X and Y. Negative values are not allowed. A value of zero disables the effect of the given filter primitive (i.e., the result is a transparent black image).`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "flood-color",
					Description: `The flood-color attribute indicates what color to use to flood the current filter primitive subregion defined through the <feFlood> element. If attribute 'flood-color' is not specified, then the effect is as if a value of black were specified.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "flood-opacity",
					Description: `The flood-opacity attribute indicates the opacity value to use across the current filter primitive subregion defined through the <feFlood> element.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feFlood",
			Description: `The <feFlood> SVG filter primitive fills the filter subregion with the color and opacity defined by flood-color and flood-opacity.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "flood-color",
					Description: `The flood-color attribute indicates what color to use to flood the current filter primitive subregion defined through the <feFlood> element. If attribute 'flood-color' is not specified, then the effect is as if a value of black were specified.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "flood-opacity",
					Description: `The flood-opacity attribute indicates the opacity value to use across the current filter primitive subregion defined through the <feFlood> element.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feFuncA",
			Description: `The <feFuncA> SVG filter primitive defines the transfer function for the alpha component of the input graphic of its parent <feComponentTransfer> element.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "type",
					Description: `The type of transfer function.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("identity", `The type of transfer function.`),
						AttributeTypeChoice("table", `The type of transfer function.`),
						AttributeTypeChoice("discrete", `The type of transfer function.`),
						AttributeTypeChoice("linear", `The type of transfer function.`),
						AttributeTypeChoice("gamma", `The type of transfer function.`),
					),
				},
				{
					Key:         "tableValues",
					Description: `Contains the list of <number>s that define the lookup table. Values must be in the 0-1 range and be equally spaced. There must be at least two values.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "slope",
					Description: `The slope attribute indicates the slope of the linear function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "intercept",
					Description: `The intercept attribute indicates the intercept of the linear function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "amplitude",
					Description: `The amplitude attribute indicates the amplitude of the cubic function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "exponent",
					Description: `The exponent attribute indicates the exponent of the exponential function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "offset",
					Description: `The offset attribute indicates the offset of the function.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feFuncB",
			Description: `The <feFuncB> SVG filter primitive defines the transfer function for the blue component of the input graphic of its parent <feComponentTransfer> element.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "type",
					Description: `The type of transfer function.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("identity", `The type of transfer function.`),
						AttributeTypeChoice("table", `The type of transfer function.`),
						AttributeTypeChoice("discrete", `The type of transfer function.`),
						AttributeTypeChoice("linear", `The type of transfer function.`),
						AttributeTypeChoice("gamma", `The type of transfer function.`),
					),
				},
				{
					Key:         "tableValues",
					Description: `Contains the list of <number>s that define the lookup table. Values must be in the 0-1 range and be equally spaced. There must be at least two values.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "slope",
					Description: `The slope attribute indicates the slope of the linear function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "intercept",
					Description: `The intercept attribute indicates the intercept of the linear function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "amplitude",
					Description: `The amplitude attribute indicates the amplitude of the cubic function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "exponent",
					Description: `The exponent attribute indicates the exponent of the exponential function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "offset",
					Description: `The offset attribute indicates the offset of the function.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feFuncG",
			Description: `The <feFuncG> SVG filter primitive defines the transfer function for the green component of the input graphic of its parent <feComponentTransfer> element.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "type",
					Description: `The type of transfer function.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("identity", `The type of transfer function.`),
						AttributeTypeChoice("table", `The type of transfer function.`),
						AttributeTypeChoice("discrete", `The type of transfer function.`),
						AttributeTypeChoice("linear", `The type of transfer function.`),
						AttributeTypeChoice("gamma", `The type of transfer function.`),
					),
				},
				{
					Key:         "tableValues",
					Description: `Contains the list of <number>s that define the lookup table. Values must be in the 0-1 range and be equally spaced. There must be at least two values.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "slope",
					Description: `The slope attribute indicates the slope of the linear function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "intercept",
					Description: `The intercept attribute indicates the intercept of the linear function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "amplitude",
					Description: `The amplitude attribute indicates the amplitude of the cubic function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "exponent",
					Description: `The exponent attribute indicates the exponent of the exponential function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "offset",
					Description: `The offset attribute indicates the offset of the function.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feFuncR",
			Description: `The <feFuncR> SVG filter primitive defines the transfer function for the red component of the input graphic of its parent <feComponentTransfer> element.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "type",
					Description: `The type of transfer function.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("identity", `The type of transfer function.`),
						AttributeTypeChoice("table", `The type of transfer function.`),
						AttributeTypeChoice("discrete", `The type of transfer function.`),
						AttributeTypeChoice("linear", `The type of transfer function.`),
						AttributeTypeChoice("gamma", `The type of transfer function.`),
					),
				},
				{
					Key:         "tableValues",
					Description: `Contains the list of <number>s that define the lookup table. Values must be in the 0-1 range and be equally spaced. There must be at least two values.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "slope",
					Description: `The slope attribute indicates the slope of the linear function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "intercept",
					Description: `The intercept attribute indicates the intercept of the linear function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "amplitude",
					Description: `The amplitude attribute indicates the amplitude of the cubic function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "exponent",
					Description: `The exponent attribute indicates the exponent of the exponential function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "offset",
					Description: `The offset attribute indicates the offset of the function.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feGaussianBlur",
			Description: `The <feGaussianBlur> SVG filter primitive blurs the input image by the amount specified in stdDeviation, which defines the bell-curve.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The input for this filter.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "stdDeviation",
					Description: `The standard deviation for the blur operation. If two <numbers> are provided, the first number represents a standard deviation value along the x-axis of the coordinate system established by attribute 'primitiveUnits' on the <filter> element. The second value represents a standard deviation in Y. If one number is provided, then that value is used for both X and Y. Negative values are not allowed. A value of zero disables the effect of the given filter primitive (i.e., the result is a transparent black image).`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feImage",
			Description: `The <feImage> SVG filter primitive fetches image data from an external source and provides the pixel data as output (meaning if the external source is an SVG image, it is rasterized.)`,
			Attributes: []*pb.Attribute{
				{
					Key:         "externalResourcesRequired",
					Description: `Indicates whether or not to force synchronous behavior.`,
					Type:        AttributeTypeBool(),
				},
				{
					Key:         "preserveAspectRatio",
					Description: `Indicates how the fetched image is fitted into the destination rectangle.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("none", `Do not force uniform scaling.`),
						AttributeTypeChoice("xMinYMin", `Scale the image to the smallest size such that both its width and its height can completely fit inside the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMin", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMaxYMin", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMinYMid", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMid", `Scale the image to the smallest size such that it can completely fit inside the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMaxYMid", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMinYMax", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMax", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMaxYMax", `Scale the image to the smallest size such that both its width and its height can completely fit inside the corresponding dimension of the viewPort.`),
					),
				},
				{
					Key:         "href",
					Description: `A URI reference to an external resource.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "feMerge",
			Description: `The <feMerge> SVG element allows filter effects to be applied concurrently instead of sequentially.`,
		},

		{
			Tag:         "feMergeNode",
			Description: `The <feMergeNode> SVG element allows a series of filter primitives to be connected together graphically. Incoming nodes are blended into the background via the defined compositing operator.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The identifier for the input SVGAnimatedString attribute on the given 'feMergeNode' element.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "feMorphology",
			Description: `The <feMorphology> SVG filter primitive is used to erode or dilate the input image. It's usefulness lies especially in fattening or thinning effects.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The input for this filter.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "operator",
					Description: `The operator attribute defines what type of operation is performed.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("erode", `The operator attribute defines what type of operation is performed.`),
						AttributeTypeChoice("dilate", `The operator attribute defines what type of operation is performed.`),
					),
				},
				{
					Key:         "radius",
					Description: `The radius attribute indicates the size of the matrix.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feOffset",
			Description: `The <feOffset> SVG filter primitive allows to offset the input image. The amount of offset can be controlled by attributes dx and dy.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "dx",
					Description: `The dx attribute indicates a shift along the x-axis on the kernel matrix.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "dy",
					Description: `The dy attribute indicates a shift along the y-axis on the kernel matrix.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "fePointLight",
			Description: `The <fePointLight> SVG filter primitive defines a light source at a point on the plane of the user coordinate system.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "x",
					Description: `The x attribute indicates the x location of the light source in the coordinate system established by attribute 'primitiveUnits' on the <filter> element.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y",
					Description: `The y attribute indicates the y location of the light source in the coordinate system established by attribute 'primitiveUnits' on the <filter> element.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "z",
					Description: `The z attribute indicates the z location of the light source in the coordinate system established by attribute 'primitiveUnits' on the <filter> element.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feSpecularLighting",
			Description: `The <feSpecularLighting> SVG filter primitive lights a source graphic using the alpha channel as a bump map. The resulting image is an RGBA image based on the light color. The lighting calculation follows the standard specular component of the Phong lighting model. The resulting image depends on the light color, light position and surface geometry of the input bump map.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The input for this filter.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "surfaceScale",
					Description: `The 'surfaceScale' attribute indicates the height of the surface when the alpha channel is 1.0.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "specularConstant",
					Description: `The specularConstant attribute represents the diffuse reflection constant.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "specularExponent",
					Description: `The specularExponent attribute represents the specular reflection constant.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "kernelUnitLength",
					Description: `The kernelUnitLength attribute defines the intended distance in current filter units (i.e., units as determined by the value of attribute 'primitiveUnits') for dx and dy in the surface normal calculation formulas.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "feSpotLight",
			Description: `The <feSpotLight> SVG filter primitive allows to create a light source placed at a point x, y, z.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "x",
					Description: `The x attribute indicates the x location of the light source in the coordinate system established by attribute 'primitiveUnits' on the <filter> element.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y",
					Description: `The y attribute indicates the y location of the light source in the coordinate system established by attribute 'primitiveUnits' on the <filter> element.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "z",
					Description: `The z attribute indicates the z location of the light source in the coordinate system established by attribute 'primitiveUnits' on the <filter> element.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "pointsAtX",
					Description: `The pointsAtX attribute indicates the x location in the coordinate system established by attribute 'primitiveUnits' on the <filter> element of the point at which the light source is pointing.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "pointsAtY",
					Description: `The pointsAtY attribute indicates the y location in the coordinate system established by attribute 'primitiveUnits' on the <filter> element of the point at which the light source is pointing.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "pointsAtZ",
					Description: `The pointsAtZ attribute indicates the z location in the coordinate system established by attribute 'primitiveUnits' on the <filter> element of the point at which the light source is pointing.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "specularExponent",
					Description: `The specularExponent attribute represents the specular reflection constant.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "limitingConeAngle",
					Description: `The limitingConeAngle attribute represents the angle in degrees between the spot light axis and the spot light cone.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "feTile",
			Description: `The <feTile> SVG filter primitive allows to fill a target rectangle with a repeated, tiled pattern of an input image. The effect is similar to the one of a <pattern> element, but <feTile> can use complex (i.e., filter) tree as input, and can be animated.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "in",
					Description: `The input for this filter.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "feTurbulence",
			Description: `The <feTurbulence> SVG filter primitive creates an image using the Perlin turbulence function. It allows the synthesis of artificial textures like clouds or marble.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "baseFrequency",
					Description: `The baseFrequency attribute represent the base frequencies in the X and Y directions of the turbulence function.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "numOctaves",
					Description: `The numOctaves attribute indicates the number of octaves to be used by the noise function.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "seed",
					Description: `The seed attribute indicates which number to use to seed the random number generator.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "stitchTiles",
					Description: `The stitchTiles attribute indicates how the Perlin noise function should be tiled. It is ignored if type is not set to 'turbulence'.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("noStitch", `The <feTurbulence> SVG filter primitive creates an image using the Perlin turbulence function. It allows the synthesis of artificial textures like clouds or marble.`),
						AttributeTypeChoice("stitch", `The <feTurbulence> SVG filter primitive creates an image using the Perlin turbulence function. It allows the synthesis of artificial textures like clouds or marble.`),
					),
				},
				{
					Key:         "type",
					Description: `The type of turbulence function.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("fractalNoise", `The type of turbulence function.`),
						AttributeTypeChoice("turbulence", `The type of turbulence function.`),
					),
				},
			},
		},

		{
			Tag:         "filter",
			Description: `The <filter> SVG element defines a custom filter effect by grouping atomic filter primitives. It is never rendered directly. A filter is referenced by using the filter attribute on the target SVG element or via the filter CSS property.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "filterUnits",
					Description: `The coordinate system for attributes x, y, width and height.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The coordinate system for attributes x, y, width and height.`),
						AttributeTypeChoice("objectBoundingBox", `The coordinate system for attributes x, y, width and height.`),
					),
				},
				{
					Key:         "primitiveUnits",
					Description: `The coordinate system for the various length values within the filter.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The coordinate system for the various length values within the filter.`),
						AttributeTypeChoice("objectBoundingBox", `The coordinate system for the various length values within the filter.`),
					),
				},
				{
					Key:         "x",
					Description: `The x attribute indicates where the left edge of the filter is placed.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "y",
					Description: `The y attribute indicates where the top edge of the filter is placed.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "width",
					Description: `The width attribute indicates the width of the filter primitive box.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "height",
					Description: `The height attribute indicates the height of the filter primitive box.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "foreignObject",
			Description: `The <foreignObject> SVG element allows for inclusion of a foreign XML namespace which has its graphical content drawn by a different user agent. The included foreign graphical content is subject to SVG transformations and compositing.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "x",
					Description: `The x-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "y",
					Description: `The y-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "width",
					Description: `The width of the rectangular region.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "height",
					Description: `The height of the rectangular region.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "requiredExtensions",
					Description: `A space-separated list of required extensions, indicating that the parent SVG document must include the specified extensions for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "requiredFeatures",
					Description: `A space-separated list of required features, indicating that the parent SVG document must include support for all of the specified features for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "systemLanguage",
					Description: `A space-separated list of language codes, indicating that the parent SVG document must include support for all of the specified languages for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "g",
			Description: `The <g> SVG element is a container used to group other SVG elements.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "requiredExtensions",
					Description: `A space-separated list of required extensions, indicating that the parent SVG document must include the specified extensions for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "requiredFeatures",
					Description: `A space-separated list of required features, indicating that the parent SVG document must include support for all of the specified features for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "systemLanguage",
					Description: `A space-separated list of language codes, indicating that the parent SVG document must include support for all of the specified languages for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "image",
			Description: `The <image> SVG element includes images inside SVG documents.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "preserveAspectRatio",
					Description: `Indicates how the fetched image is fitted into the destination rectangle.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("none", `Do not force uniform scaling.`),
						AttributeTypeChoice("xMinYMin", `Scale the image to the smallest size such that both its width and its height can completely fit inside the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMin", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMaxYMin", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMinYMid", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMid", `Scale the image to the smallest size such that it can completely fit inside the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMaxYMid", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMinYMax", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMax", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMaxYMax", `Scale the image to the smallest size such that both its width and its height can completely fit inside the corresponding dimension of the viewPort.`),
					),
				},
				{
					Key:         "x",
					Description: `The x-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y",
					Description: `The y-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "width",
					Description: `The width of the rectangular region.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "height",
					Description: `The height of the rectangular region.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "href",
					Description: `A URI reference to the image to embed.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "line",
			Description: `The <line> SVG element is an SVG basic shape, used to create a line connecting two points.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "x1",
					Description: `The x-axis coordinate of the starting point of the line.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y1",
					Description: `The y-axis coordinate of the starting point of the line.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "x2",
					Description: `The x-axis coordinate of the ending point of the line.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y2",
					Description: `The y-axis coordinate of the ending point of the line.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "linearGradient",
			Description: `The <linearGradient> SVG element lets authors define linear gradients to fill or stroke graphical elements.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "gradientUnits",
					Description: `The coordinate system for attributes x1, y1, x2 and y2.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The coordinate system for attributes x1, y1, x2 and y2.`),
						AttributeTypeChoice("objectBoundingBox", `The coordinate system for attributes x1, y1, x2 and y2.`),
					),
				},
				{
					Key:         "gradientTransform",
					Description: `The definition of how the gradient is applied, read about <a href="https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/gradientTransform">gradientTransform</a>.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "spreadMethod",
					Description: `The method by which to fill a shape.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("pad", `The method by which to fill a shape.`),
						AttributeTypeChoice("reflect", `The method by which to fill a shape.`),
						AttributeTypeChoice("repeat", `The method by which to fill a shape.`),
					),
				},
				{
					Key:         "x1",
					Description: `The x-axis coordinate of the start of the gradient.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y1",
					Description: `The y-axis coordinate of the start of the gradient.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "x2",
					Description: `The x-axis coordinate of the end of the gradient.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y2",
					Description: `The y-axis coordinate of the end of the gradient.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "marker",
			Description: `The <marker> SVG element defines the graphics that is to be used for drawing arrowheads or polymarkers on a given <path>, <line>, <polyline> or <polygon> element.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "refX",
					Description: `The x-axis coordinate of the reference point which is to be aligned exactly at the marker position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "refY",
					Description: `The y-axis coordinate of the reference point which is to be aligned exactly at the marker position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "markerWidth",
					Description: `The width of the marker viewport.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "markerHeight",
					Description: `The height of the marker viewport.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "orient",
					Description: `The orientation of the marker relative to the shape it is attached to.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("auto", `The orientation of the marker relative to the shape it is attached to.`),
						AttributeTypeChoice("auto-start-reverse", `The orientation of the marker relative to the shape it is attached to.`),
						AttributeTypeChoice("angle", `The orientation of the marker relative to the shape it is attached to.`),
					),
				},
				{
					Key:         "markerUnits",
					Description: `The coordinate system for the various length values within the marker.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The coordinate system for the various length values within the marker.`),
						AttributeTypeChoice("strokeWidth", `The coordinate system for the various length values within the marker.`),
					),
				},
				{
					Key:         "viewBox",
					Description: `The position and size of the marker viewport (the bounds of the marker).`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "mask",
			Description: `The <mask> SVG element hides portions of SVG elements for user display.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "maskContentUnits",
					Description: `The coordinate system for attributes x, y, width and height.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The coordinate system for attributes x, y, width and height.`),
						AttributeTypeChoice("objectBoundingBox", `The coordinate system for attributes x, y, width and height.`),
					),
				},
				{
					Key:         "maskUnits",
					Description: `The coordinate system for the various length values within the filter.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The coordinate system for the various length values within the filter.`),
						AttributeTypeChoice("objectBoundingBox", `The coordinate system for the various length values within the filter.`),
					),
				},
				{
					Key:         "x",
					Description: `The x-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "y",
					Description: `The y-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "width",
					Description: `The width of the rectangular region.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "height",
					Description: `The height of the rectangular region.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "metadata",
			Description: `The <metadata> SVG element allows to add metadata to SVG content. Metadata is structured information about data. In XML, metadata can be added to an element using for example attributes.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "requiredExtensions",
					Description: `A space-separated list of required extensions, indicating that the parent SVG document must include the specified extensions for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "requiredFeatures",
					Description: `A space-separated list of required features, indicating that the parent SVG document must include support for all of the specified features for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "systemLanguage",
					Description: `A space-separated list of language codes, indicating that the parent SVG document must include support for all of the specified languages for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "mpath",
			Description: `The <mpath> SVG element allows to use the functionality of <animateMotion> to animate the <startOffset> attribute of SVG <textPath> elements.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "href",
					Description: `A URI reference to the motion path definition.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "path",
			Description: `The <path> SVG element is the generic element to define a shape. All the basic shapes can be created with a path element.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "d",
					Description: `The definition of the outline of a shape.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "pathLength",
					Description: `The total length for the path, in user units.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "pattern",
			Description: `The <pattern> SVG element fills a region with a pattern defined by an SVG image.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "patternUnits",
					Description: `The coordinate system for attributes x, y, width and height.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The coordinate system for attributes x, y, width and height.`),
						AttributeTypeChoice("objectBoundingBox", `The coordinate system for attributes x, y, width and height.`),
					),
				},
				{
					Key:         "patternContentUnits",
					Description: `The coordinate system for the various length values within the filter.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The coordinate system for the various length values within the filter.`),
						AttributeTypeChoice("objectBoundingBox", `The coordinate system for the various length values within the filter.`),
					),
				},
				{
					Key:         "patternTransform",
					Description: `The definition of how the pattern is tiled, read about <a href="https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/patternTransform">patternTransform</a>.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "x",
					Description: `The x-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y",
					Description: `The y-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "width",
					Description: `The width of the rectangular region.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "height",
					Description: `The height of the rectangular region.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "href",
					Description: `A URI reference to the image to paint.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "polygon",
			Description: `The <polygon> SVG element is an SVG basic shape, used to create a vector-based polygonal shape.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "points",
					Description: `A list of points, each of which is a coordinate pair.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "polyline",
			Description: `The <polyline> SVG element is an SVG basic shape, used to create a series of straight lines connecting several points. Typically a polyline is used to create open shapes as the last point doesn't have to be connected to the first point. For closed shapes see the <polygon> element.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "points",
					Description: `A list of points, each of which is a coordinate pair.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "radialGradient",
			Description: `The <radialGradient> SVG element lets authors define radial gradients to fill or stroke graphical elements.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "gradientUnits",
					Description: `The coordinate system for attributes cx, cy and r.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("userSpaceOnUse", `The coordinate system for attributes cx, cy and r.`),
						AttributeTypeChoice("objectBoundingBox", `The coordinate system for attributes cx, cy and r.`),
					),
				},
				{
					Key:         "gradientTransform",
					Description: `The definition of how the gradient is applied, read about <a href="https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/gradientTransform">gradientTransform</a>.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "cx",
					Description: `The x-axis coordinate of the largest (i.e., outermost) circle for the radial gradient.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "cy",
					Description: `The y-axis coordinate of the largest (i.e., outermost) circle for the radial gradient.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "r",
					Description: `The radius of the largest (i.e., outermost) circle for the radial gradient.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "fx",
					Description: `The x-axis coordinate of the point at which the focal point of the radial gradient is placed.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "fy",
					Description: `The y-axis coordinate of the point at which the focal point of the radial gradient is placed.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "rect",
			Description: `The <rect> SVG element is a basic SVG shape that draws rectangles, defined by their position, width, and height. The shape is created by connecting a line from one point to the other three points.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "x",
					Description: `The x-axis coordinate of the side of the rectangle which has the smaller x-axis value.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y",
					Description: `The y-axis coordinate of the side of the rectangle which has the smaller y-axis value.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "width",
					Description: `The width of the rectangle.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "height",
					Description: `The height of the rectangle.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "rx",
					Description: `The x-axis radius of the ellipse used to round off the corners of the rectangle.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "ry",
					Description: `The y-axis radius of the ellipse used to round off the corners of the rectangle.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},
		{
			Tag:         "script",
			Description: `The <script> SVG element includes scripts, which can be used to trigger user interface events.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "type",
					Description: `The scripting language used for the given script element.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "href",
					Description: `A Uniform Resource Identifier (URI) reference that specifies the location of the script to execute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "crossorigin",
					Description: `How the element handles crossorigin requests.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("anonymous", `How the element handles crossorigin requests.`),
						AttributeTypeChoice("use-credentials", `How the element handles crossorigin requests.`),
					),
				},
			},
		},

		{
			Tag:         "set",
			Description: `The <set> SVG element provides a simple means of just setting the value of an attribute for a specified duration. It supports all attribute types, including those that cannot reasonably be interpolated, such as string and boolean values. The <set> element is non-additive. The additive and accumulate attributes are not allowed, and will be ignored if specified.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "to",
					Description: `The target attribute value to assign on end.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "attributeName",
					Description: `The name of the attribute to assign.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "attributeType",
					Description: `The namespace in which the target attribute and its associated values are defined.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("auto", `The namespace in which the target attribute and its associated values are defined.`),
						AttributeTypeChoice("CSS", `The namespace in which the target attribute and its associated values are defined.`),
						AttributeTypeChoice("XML", `The namespace in which the target attribute and its associated values are defined.`),
						AttributeTypeChoice("XMLNS", `The namespace in which the target attribute and its associated values are defined.`),
						AttributeTypeChoice("empty", `The namespace in which the target attribute and its associated values are defined.`),
					),
				},
				{
					Key:         "begin",
					Description: `The begin time for the element.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "dur",
					Description: `The simple duration for the element.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "end",
					Description: `The end for the element.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "min",
					Description: `The minimum value allowed for the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "max",
					Description: `The maximum value allowed for the attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "restart",
					Description: `Defines how the element is restarted.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("always", `Defines how the element is restarted.`),
						AttributeTypeChoice("whenNotActive", `Defines how the element is restarted.`),
						AttributeTypeChoice("never", `Defines how the element is restarted.`),
					),
				},
				{
					Key:         "repeatCount",
					Description: `Defines the number of times the element is repeated.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "repeatDur",
					Description: `Defines the duration for the element to repeat.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "fill",
					Description: `Defines the value the animation will have before the begin event.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("remove", `Defines the value the animation will have before the begin event.`),
						AttributeTypeChoice("freeze", `Defines the value the animation will have before the begin event.`),
					),
				},
			},
		},
		{
			Tag:         "stop",
			Description: `The <stop> SVG element defines the ramp of colors to use on a gradient, which is a child element to either the <linearGradient> or the <radialGradient> element.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "offset",
					Description: `The offset from the start of the gradient where the color first takes effect.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "stop-color",
					Description: `The color of the gradient stop.`,
					Type:        AttributeTypeString(),
				},
			},
		},
		{
			Tag:         "style",
			Description: `The <style> SVG element allows style sheets to be embedded directly within SVG content. SVG's style element has the same attributes as the corresponding element in HTML (see HTML's <style> element).`,
			Attributes: []*pb.Attribute{
				{
					Key:         "type",
					Description: `The style sheet language.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "media",
					Description: `The intended destination medium for style information.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "title",
					Description: `The advisory title.`,
					Type:        AttributeTypeString(),
				},
			},
		},
		{
			Tag:         "svg",
			Description: `The <svg> element is a container that defines a new coordinate system and viewport. It is used as the outermost element of SVG documents, but it can also be used to embed a SVG fragment inside an SVG or HTML document.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "x",
					Description: `The x-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "y",
					Description: `The y-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "width",
					Description: `The width of the rectangular region.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "height",
					Description: `The height of the rectangular region.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "viewBox",
					Description: `The position and size of the viewport (the viewBox) is defined by the viewBox attribute.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "preserveAspectRatio",
					Description: `Indicates how the viewport is fitted to the rectangle of the given width and height.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("none", `Do not force uniform scaling.`),
						AttributeTypeChoice("xMinYMin", `Scale the image to the smallest size such that both its width and its height can completely fit inside the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMin", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMaxYMin", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMinYMid", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMid", `Scale the image to the smallest size such that it can completely fit inside the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMaxYMid", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMinYMax", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMax", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMaxYMax", `Scale the image to the smallest size such that both its width and its height can completely fit inside the corresponding dimension of the viewPort.`),
					),
				},
			},
		},

		{
			Tag:         "switch",
			Description: `The <switch> SVG element evaluates the requiredFeatures, requiredExtensions and systemLanguage attributes on its direct child elements in order, and then processes and renders the first child for which these attributes evaluate to true. All others will be bypassed and therefore not rendered. If the child element is a container element such as a <g>, then the entire subtree is either processed/rendered or bypassed/not rendered.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "requiredFeatures",
					Description: `A space-separated list of required features, indicating that the parent SVG document must include support for all of the specified features for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "requiredExtensions",
					Description: `A space-separated list of required extensions, indicating that the parent SVG document must include the specified extensions for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "systemLanguage",
					Description: `A space-separated list of language codes, indicating that the parent SVG document must include support for all of the specified languages for this element to be valid.`,
					Type:        AttributeTypeString(),
				},
			},
		},

		{
			Tag:         "symbol",
			Description: `The <symbol> SVG element is used to define graphical template objects which can be instantiated by a <use> element. The use of symbol elements for graphics that are used multiple times in the same document adds structure and semantics. Documents that are rich in structure may be rendered graphically, as speech, or as Braille, and thus promote accessibility. note that a symbol element itself is not rendered. Only instances of a symbol element (i.e., a reference to a symbol by a <use> element) are rendered. To render a 'stand-alone' graphic that has been defined using a symbol, a reference to the symbol is referenced using a <use> element.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "preserveAspectRatio",
					Description: `Indicates how the fetched image is fitted into the destination rectangle.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("none", `Do not force uniform scaling.`),
						AttributeTypeChoice("xMinYMin", `Scale the image to the smallest size such that both its width and its height can completely fit inside the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMin", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMaxYMin", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMinYMid", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMid", `Scale the image to the smallest size such that it can completely fit inside the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMaxYMid", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMinYMax", `Align the image along the middle of the corresponding dimension of the viewPort.`),
						AttributeTypeChoice("xMidYMax", `Align the image with the corresponding side of the viewPort.`),
						AttributeTypeChoice("xMaxYMax", `Scale the image to the smallest size such that both its width and its height can completely fit inside the corresponding dimension of the viewPort.`),
					),
				},
			},
		},

		{
			Tag:         "text",
			Description: `The <text> SVG element renders the first character at the initial current text position. This position is modified by the lengthAdjust and textLength attributes.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "x",
					Description: `The x-axis coordinate of the initial current text position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y",
					Description: `The y-axis coordinate of the initial current text position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "dx",
					Description: `The x-axis coordinate of the current text position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "dy",
					Description: `The y-axis coordinate of the current text position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "rotate",
					Description: `The rotation angle about the current text position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "textLength",
					Description: `The total sum of all of the advance values from rendering all of the characters within this element, including the advance value on the glyph (horizontal or vertical), the effect of properties 'kerning', 'letter-spacing' and 'word-spacing' and adjustments due to attributes 'x' and 'y' on the <text> element.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "lengthAdjust",
					Description: `Indicates how the text is stretched or compressed to fit the textLength.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("spacing", `Indicates how the text is stretched or compressed to fit the textLength.`),
						AttributeTypeChoice("spacingAndGlyphs", `Indicates how the text is stretched or compressed to fit the textLength.`),
					),
				},
			},
		},

		{
			Tag:         "textPath",
			Description: `The <textPath> SVG element defines a set of glyphs that exactly fit along a curve.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "href",
					Description: `A URI reference to the path to render along.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "startOffset",
					Description: `Indicates an offset from the start of the path, where the first character is rendered.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "method",
					Description: `Indicates the method by which text should be rendered along the path.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("align", `Indicates the method by which text should be rendered along the path.`),
						AttributeTypeChoice("stretch", `Indicates the method by which text should be rendered along the path.`),
					),
				},
				{
					Key:         "spacing",
					Description: `Indicates the spacing behavior between characters.`,
					Type: AttributeTypeChoices(
						AttributeTypeChoice("auto", `Indicates the spacing behavior between characters.`),
						AttributeTypeChoice("exact", `Indicates the spacing behavior between characters.`),
					),
				},
			},
		},

		{
			Tag:         "title",
			Description: `The <title> SVG element represents a text container used for the context information. This element is usually nested inside a <desc> element.`,
		},

		{
			Tag:         "tspan",
			Description: `The <tspan> SVG element lets authors explicitly specify the location of a glyph along the given path via the attributes.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "x",
					Description: `The x-axis coordinate of the current text position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y",
					Description: `The y-axis coordinate of the current text position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "dx",
					Description: `The x-axis coordinate of the current text position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "dy",
					Description: `The y-axis coordinate of the current text position.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "rotate",
					Description: `The rotation angle about the current text position.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "use",
			Description: `The <use> SVG element includes a reference to a <symbol> element and attempts to display the referenced content. The reference is drawn exactly as it was defined. It can be reused as often as needed and can be programmatically manipulated.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "href",
					Description: `A URI reference to the symbol to use.`,
					Type:        AttributeTypeString(),
				},
				{
					Key:         "x",
					Description: `The x-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "y",
					Description: `The y-axis coordinate of the side of the rectangular region which is closest to the user.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "width",
					Description: `The width of the rectangular region.`,
					Type:        AttributeTypeNumber(),
				},
				{
					Key:         "height",
					Description: `The height of the rectangular region.`,
					Type:        AttributeTypeNumber(),
				},
			},
		},

		{
			Tag:         "view",
			Description: `The <view> SVG element is used to define a view into a &lt;svg&gt; element. It is partially deprecated in SVG 2.0 and should generally not be used.`,
			Attributes: []*pb.Attribute{
				{
					Key:         "viewBox",
					Description: `The position and size of the viewport (the viewBox) is defined by the viewBox attribute.`,
					Type:        AttributeTypeString(),
				},
			},
		},
	},
}
