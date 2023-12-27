package cfg

import (
	pb "github.com/delaneyj/gostar/cfg/gen/specs/v1"
)

var MathML = &pb.Namespace{
	Name:        "mathml",
	Description: `This MathML Core specification intends to address these issues by being as accurate as possible on the visual rendering of mathematical formulas using additional rules from the TeXBook’s Appendix G [TEXBOOK] and from the Open Font Format [OPEN-FONT-FORMAT], [OPEN-TYPE-MATH-ILLUMINATED]. It also relies on modern browser implementations and web technologies [HTML] [SVG] [CSS2] [DOM], clarifying interactions with them when needed or introducing new low-level primitives to improve the web platform layering.`,
	Prefix:      "MathML",
	Attributes: []*pb.Attribute{
		{
			Key:         "class",
			Description: "Assigns a class name or set of class names to an element. You may assign the same class name or names to any number of elements. If you specify multiple class names, they must be separated by whitespace characters.",
			Type:        AttributeTypeSpaceDelimited(),
		},
		{
			Key:         "dir",
			Description: "This attribute specifies the text directionality of the element, merely indicating what direction the text flows when surrounded by text with inherent directionality (such as Arabic or Hebrew). Possible values are ltr (left-to-right) and rtl (right-to-left).",
			Type: AttributeTypeChoices(
				AttributeTypeChoice("ltr", "left-to-right"),
				AttributeTypeChoice("rtl", "right-to-left"),
			),
		},
		{
			Key:         "displaystyle",
			Description: "This attribute specifies whether the element should be rendered using displaystyle rules or not. Possible values are true and false.",
			Type: AttributeTypeChoices(
				AttributeTypeChoice("true", "displaystyle rules"),
				AttributeTypeChoice("false", "not displaystyle rules"),
			),
		},
		{
			Key:         "id",
			Description: "This attribute assigns a name to an element. This name must be unique in a document.",
			Type:        AttributeTypeString(),
		},
		{
			Key:         "mathbackground",
			Description: "This attribute specifies the background color of the element. Possible values are a color name or a color specification in the format defined in the CSS3 Color Module [CSS3COLOR].",
			Type:        AttributeTypeString(),
		},
		{
			Key:         "mathcolor",
			Description: "This attribute specifies the color of the element. Possible values are a color name or a color specification in the format defined in the CSS3 Color Module [CSS3COLOR].",
			Type:        AttributeTypeString(),
		},
		{
			Key:         "mathsize",
			Name:        "mathsizeStr",
			Description: "This attribute specifies the size of the element. Possible values are a dimension or a dimensionless number.",
			Type:        AttributeTypeString(),
		},
		{
			Key:         "mathsize",
			Description: "This attribute specifies the size of the element. Possible values are a dimension or a dimensionless number.",
			Type:        AttributeTypeInt(),
		},
		{
			Key:         "nonce",
			Description: "This attribute declares a cryptographic nonce (number used once) that should be used by the server processing the element’s submission, and the resulting resource must be delivered with a Content-Security-Policy nonce attribute matching the value of the nonce attribute.",
			Type:        AttributeTypeString(),
		},
		{
			Key:         "scriptlevel",
			Description: "This attribute specifies the script level of the element. Possible values are an integer between 0 and 7, inclusive.",
			Type:        AttributeTypeInt(),
		},
		{
			Key:         "style",
			Description: "This attribute offers advisory information about the element for which it is set.",
			Type:        AttributeTypeKVColonSemicolon(),
		},
		{
			Key:         "tabindex",
			Description: "This attribute specifies the position of the current element in the tabbing order for the current document. This value must be a number between 0 and 32767. User agents should ignore leading zeros.",
			Type:        AttributeTypeInt(),
		},
	},
	Elements: []*pb.Element{
		{
			Tag:         "annotation",
			Description: "This element is used to include comments or annotations within a MathML expression. It can be used to provide additional information about the expression, or to include comments for the author of the expression.",
			Attributes: []*pb.Attribute{
				{
					Key:         "encoding",
					Description: "This attribute specifies the encoding used for the text content of the element. Possible values are text/plain, text/html, and application/x-tex.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("text/plain", ""),
						AttributeTypeChoice("text/html", ""),
						AttributeTypeChoice("application/x-tex", ""),
					),
				},
				{
					Key:         "name",
					Description: "This attribute specifies the name of the annotation.",
					Type:        AttributeTypeString(),
				},
			},
		},
		{
			Tag:         "annotation-xml",
			Description: "This element is used to include comments or annotations within a MathML expression. It can be used to provide additional information about the expression, or to include comments for the author of the expression.",
			Attributes: []*pb.Attribute{
				{
					Key:         "encoding",
					Description: "This attribute specifies the encoding used for the text content of the element. Possible values are text/plain, text/html, and application/x-tex.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("text/plain", ""),
						AttributeTypeChoice("text/html", ""),
						AttributeTypeChoice("application/x-tex", ""),
					),
				},
				{
					Key:         "name",
					Description: "This attribute specifies the name of the annotation.",
					Type:        AttributeTypeString(),
				},
			},
		},
		{
			Tag:         "maction",
			Description: "This element is used to specify the behavior of a subexpression when it is activated by the user. The action can be to toggle the visibility of the subexpression, or to toggle the selection of the subexpression, or to execute a script.",
			Attributes: []*pb.Attribute{
				{
					Key:         "actiontype",
					Description: "This attribute specifies the type of action performed when the element is activated. Possible values are toggle, statusline, tooltip, and script.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("toggle", ""),
						AttributeTypeChoice("statusline", ""),
						AttributeTypeChoice("tooltip", ""),
						AttributeTypeChoice("script", ""),
					),
				},
				{
					Key:         "selection",
					Description: "This attribute specifies the type of selection performed when the element is activated. Possible values are none, highlight, and unhighlight.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("none", ""),
						AttributeTypeChoice("highlight", ""),
						AttributeTypeChoice("unhighlight", ""),
					),
				},
			},
		},
		{
			Tag:         "math",
			Description: "This element is the root element of a MathML expression. It is used to identify the document as a MathML document, and to specify the namespaces used in the document.",
			Attributes: []*pb.Attribute{
				{
					Key:         "xmlns",
					Description: "This attribute specifies the default namespace for elements and attributes in the document. Possible values are http://www.w3.org/1998/Math/MathML and http://www.w3.org/1999/xhtml.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("http://www.w3.org/1998/Math/MathML", ""),
						AttributeTypeChoice("http://www.w3.org/1999/xhtml", ""),
					),
				},
				{
					Key:         "xmlns:m",
					Description: "This attribute specifies the namespace for elements and attributes in the document whose names start with the letter m. Possible values are http://www.w3.org/1998/Math/MathML and http://www.w3.org/1999/xhtml.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("http://www.w3.org/1998/Math/MathML", ""),
						AttributeTypeChoice("http://www.w3.org/1999/xhtml", ""),
					),
				},
				{
					Key:         "xmlns:xlink",
					Description: "This attribute specifies the namespace for elements and attributes in the document whose names start with the letters xlink. Possible values are http://www.w3.org/1999/xlink and http://www.w3.org/1999/xhtml.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("http://www.w3.org/1999/xlink", ""),
						AttributeTypeChoice("http://www.w3.org/1999/xhtml", ""),
					),
				},
				{
					Key:         "xmlns:xml",
					Description: "This attribute specifies the namespace for elements and attributes in the document whose names start with the letters xml. Possible values are http://www.w3.org/XML/1998/namespace and http://www.w3.org/1999/xhtml.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("http://www.w3.org/XML/1998/namespace", ""),
						AttributeTypeChoice("http://www.w3.org/1999/xhtml", ""),
					),
				},
			},
		},
		{
			Tag:         "merror",
			Description: `This element is used to indicate that an error has occurred while processing a MathML expression. It can be used to display an error message, or to highlight the error in the expression.`,
		},
		{
			Tag:         "mfrac",
			Description: "This element is used to display a fraction.",
			Attributes: []*pb.Attribute{
				{
					Key:         "bevelled",
					Description: "This attribute specifies whether the fraction line is to be drawn straight or to beveled. Possible values are true and false.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("true", ""),
						AttributeTypeChoice("false", ""),
					),
				},
			},
		},
		{
			Tag:         "mi",
			Description: "This element is used to display a single identifier or a single operator.",
			Attributes: []*pb.Attribute{
				{
					Key:         "mathvariant",
					Description: "This attribute specifies the variant form of the character. Possible values are normal, bold, italic, bold-italic, double-struck, bold-fraktur, script, bold-script, fraktur, sans-serif, bold-sans-serif, sans-serif-italic, sans-serif-bold-italic, monospace, initial, and tailed.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("normal", ""),
						AttributeTypeChoice("bold", ""),
						AttributeTypeChoice("italic", ""),
						AttributeTypeChoice("bold-italic", ""),
						AttributeTypeChoice("double-struck", ""),
						AttributeTypeChoice("bold-fraktur", ""),
						AttributeTypeChoice("script", ""),
						AttributeTypeChoice("bold-script", ""),
						AttributeTypeChoice("fraktur", ""),
						AttributeTypeChoice("sans-serif", ""),
						AttributeTypeChoice("bold-sans-serif", ""),
						AttributeTypeChoice("sans-serif-italic", ""),
						AttributeTypeChoice("sans-serif-bold-italic", ""),
						AttributeTypeChoice("monospace", ""),
						AttributeTypeChoice("initial", ""),
						AttributeTypeChoice("tailed", ""),
					),
				},
			},
		},
		{
			Tag:         "mmultiscripts",
			Description: "This element is used to display a base expression with multiple subscripts and superscripts.",
		},
		{
			Tag:         "mn",
			Description: "This element is used to display a single number.",
		},
		{
			Tag:         "mo",
			Description: "This element is used to display a single operator.",
			Attributes: []*pb.Attribute{
				{
					Key:         "fence",
					Description: "This attribute specifies whether the operator is to be rendered as a fence. Possible values are true and false.",
					Type: AttributeTypeChoices(
						AttributeTypeChoice("true", ""),
						AttributeTypeChoice("false", ""),
					),
				},
				{
					Key:         "lspace",
					Description: "This attribute specifies the minimum amount of space that should be left on the left side of the operator. Possible values are a dimension or a dimensionless number.",
					Type:        AttributeTypeString(),
				},
				{
					Key:         "rspace",
					Description: "This attribute specifies the minimum amount of space that should be left on the right side of the operator. Possible values are a dimension or a dimensionless number.",
					Type:        AttributeTypeString(),
				},
			},
		},
		{
			Tag:         "mover",
			Description: "This element is used to display an expression with an overbar.",
		},
		{
			Tag:         "mpadded",
			Description: "This element is used to display an expression with additional spacing.",
		},
		{
			Tag:         "mphantom",
			Description: "This element is used to display an expression without rendering it.",
		},
		{
			Tag:         "mprescripts",
			Description: "This element is used to display a base expression with multiple subscripts and superscripts.",
		},
		{
			Tag:         "mroot",
			Description: "This element is used to display an expression with a radical.",
		},
		{
			Tag:         "mrow",
			Description: "This element is used to display a sequence of expressions.",
		},
		{
			Tag:         "ms",
			Description: "This element is used to display a single character or a single number.",
		},
		{
			Tag:         "mspace",
			Description: "This element is used to display a space.",
		},
		{
			Tag:         "msqrt",
			Description: "This element is used to display an expression with a radical.",
		},
		{
			Tag:         "mstyle",
			Description: "This element is used to specify the style of a subexpression.",
		},
		{
			Tag:         "msub",
			Description: "This element is used to display a subscript expression.",
		},
		{
			Tag:         "msubsup",
			Description: "This element is used to display a subscript expression with a superscript expression.",
		},
		{
			Tag:         "msup",
			Description: "This element is used to display a superscript expression.",
		},
		{
			Tag:         "mtable",
			Description: "This element is used to display a table of expressions.",
		},
		{
			Tag:         "mtd",
			Description: "This element is used to display a cell in a table.",
		},
		{
			Tag:         "mtext",
			Description: "This element is used to display text.",
		},
		{
			Tag:         "mtr",
			Description: "This element is used to display a row in a table.",
		},
		{
			Tag:         "munder",
			Description: "This element is used to display an expression with an underbar.",
		},
		{
			Tag:         "munderover",
			Description: "This element is used to display an expression with both an underbar and an overbar.",
		},
		{
			Tag:         "semantics",
			Description: "This element is used to specify the meaning of a MathML expression.",
		},
	},
}
