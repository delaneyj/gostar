// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package html h4 is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The HTML <h1>–<h6> elements represent six levels of section headings
// <h1> is the highest section level and <h6> is the lowest.
type H4Element struct {
	*Element
}

// Create a new H4Element element.
// This will create a new element with the tag
// "h4" during rendering.
func H4(children ...ElementRenderer) *H4Element {
	e := NewElement("h4", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &H4Element{Element: e}
}

func (e *H4Element) Children(children ...ElementRenderer) *H4Element {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *H4Element) IfChildren(condition bool, children ...ElementRenderer) *H4Element {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *H4Element) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *H4Element {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *H4Element) Text(text string) *H4Element {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *H4Element) TextF(format string, args ...any) *H4Element {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *H4Element) Escaped(text string) *H4Element {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *H4Element) EscapedF(format string, args ...any) *H4Element {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *H4Element) CustomData(key, value string) *H4Element {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *H4Element) CustomDataRemove(key string) *H4Element {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The accesskey global attribute provides a hint for generating a keyboard
// shortcut for the current element
// The attribute value must consist of a single printable character (which
// includes accented and other characters that can be generated by the keyboard).
func (e *H4Element) ACCESSKEY(r rune) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("accesskey", string(r))
	return e
}

// Remove the attribute accesskey from the element.
func (e *H4Element) ACCESSKEYRemove() *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("accesskey")
	return e
}

// The autocapitalize global attribute is an enumerated attribute that controls
// whether and how text input is automatically capitalized as it is entered/edited
// by the user
// autocapitalize can be set on <input> and <textarea> elements, and on their
// containing <form> elements
// When autocapitalize is set on a <form> element, it sets the autocapitalize
// behavior for all contained <input>s and <textarea>s, overriding any
// autocapitalize values set on contained elements
// autocapitalize has no effect on the url, email, or password <input> types,
// where autocapitalization is never enabled
// Where autocapitalize is not specified, the adopted default behavior varies
// between browsers
// For example: Chrome and Safari default to on/sentences Firefox defaults to
// off/none.
func (e *H4Element) AUTOCAPITALIZE(c H4AutocapitalizeChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("autocapitalize", string(c))
	return e
}

type H4AutocapitalizeChoice string

const (
	// Do not automatically capitalize any text.
	H4Autocapitalize_off H4AutocapitalizeChoice = "off"
	// Do not automatically capitalize any text.
	H4Autocapitalize_none H4AutocapitalizeChoice = "none"
	// Automatically capitalize the first character of each sentence.
	H4Autocapitalize_sentences H4AutocapitalizeChoice = "sentences"
	// Automatically capitalize the first character of each sentence.
	H4Autocapitalize_on H4AutocapitalizeChoice = "on"
	// Automatically capitalize the first character of each word.
	H4Autocapitalize_words H4AutocapitalizeChoice = "words"
	// Automatically capitalize all characters.
	H4Autocapitalize_characters H4AutocapitalizeChoice = "characters"
)

// Remove the attribute autocapitalize from the element.
func (e *H4Element) AUTOCAPITALIZERemove(c H4AutocapitalizeChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("autocapitalize")
	return e
}

// The autofocus global attribute is a Boolean attribute indicating that an
// element should be focused on page load, or when the <dialog> that it is part of
// is displayed.
//
//	Accessibility concerns Automatically focusing a form control can confuse
//
// visually-impaired people using screen-reading technology and people with
// cognitive impairments
// When autofocus is assigned, screen-readers "teleport" their user to the form
// control without warning them beforehand.
//
//	Use careful consideration for accessibility when applying the autofocus
//
// attribute
// Automatically focusing on a control can cause the page to scroll on load
// The focus can also cause dynamic keyboards to display on some touch devices
// While a screen reader will announce the label of the form control receiving
// focus, the screen reader will not announce anything before the label, and the
// sighted user on a small device will equally miss the context created by the
// preceding content.
func (e *H4Element) AUTOFOCUS() *H4Element {
	e.AUTOFOCUSSet(true)
	return e
}

// Set the attribute autofocus to the value b explicitly.
func (e *H4Element) AUTOFOCUSSet(b bool) *H4Element {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("autofocus", b)
	return e
}

// Remove the attribute autofocus from the element.
func (e *H4Element) AUTOFOCUSRemove(b bool) *H4Element {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("autofocus")
	return e
}

// The class global attribute is a space-separated list of the case-sensitive
// classes of the element
// Classes allow CSS and JavaScript to select and access specific elements via the
// class selectors or functions like the DOM method
// document.getElementsByClassName.
func (e *H4Element) CLASS(s ...string) *H4Element {
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

// Remove the attribute class from the element.
func (e *H4Element) CLASSRemove(s ...string) *H4Element {
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

// The contenteditable global attribute is an enumerated attribute indicating if
// the element should be editable by the user
// If so, the browser modifies its widget to allow editing.
func (e *H4Element) CONTENTEDITABLE(c H4ContenteditableChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("contenteditable", string(c))
	return e
}

type H4ContenteditableChoice string

const (
	// The element is editable.
	H4Contenteditable_empty H4ContenteditableChoice = ""
	// The element is editable.
	H4Contenteditable_true H4ContenteditableChoice = "true"
	// The element is not editable.
	H4Contenteditable_false H4ContenteditableChoice = "false"
	// which indicates that the element's raw text is editable, but rich text
	// formatting is disabled.
	H4Contenteditable_plaintext_only H4ContenteditableChoice = "plaintext-only"
)

// Remove the attribute contenteditable from the element.
func (e *H4Element) CONTENTEDITABLERemove(c H4ContenteditableChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("contenteditable")
	return e
}

// The dir global attribute is an enumerated attribute that indicates the
// directionality of the element's text
// Note: This attribute is mandatory for the <bdo> element where it has a
// different semantic meaning
// This attribute is not inherited by the <bdi> element
// If not set, its value is auto
// This attribute can be overridden by the CSS properties direction and
// unicode-bidi, if a CSS page is active and the element supports these properties
// As the directionality of the text is semantically related to its content and
// not to its presentation, it is recommended that web developers use this
// attribute instead of the related CSS properties when possible
// That way, the text will display correctly even on a browser that doesn't
// support CSS or has the CSS deactivated
// The auto value should be used for data with an unknown directionality, like
// data coming from user input, eventually stored in a database.
func (e *H4Element) DIR(c H4DirChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type H4DirChoice string

const (
	// which means left to right and is to be used for languages that are written from
	// the left to the right (like English);
	H4Dir_ltr H4DirChoice = "ltr"
	// which means right to left and is to be used for languages that are written from
	// the right to the left (like Arabic);
	H4Dir_rtl H4DirChoice = "rtl"
	// which lets the user agent decide
	// It uses a basic algorithm as it parses the characters inside the element until
	// it finds a character with a strong directionality, then it applies that
	// directionality to the whole element.
	H4Dir_auto H4DirChoice = "auto"
)

// Remove the attribute dir from the element.
func (e *H4Element) DIRRemove(c H4DirChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// The draggable global attribute is an enumerated attribute that indicates
// whether the element can be dragged, either with native browser behavior or the
// HTML Drag and Drop API.
func (e *H4Element) DRAGGABLE(c H4DraggableChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("draggable", string(c))
	return e
}

type H4DraggableChoice string

const (
	// The element is draggable.
	H4Draggable_true H4DraggableChoice = "true"
	// The element is not draggable.
	H4Draggable_false H4DraggableChoice = "false"
	// drag behavior is the default browser behavior: only text selections, images,
	// and links can be dragged
	// For other elements, the event ondragstart must be set for drag and drop to work
	H4Draggable_empty H4DraggableChoice = ""
	// drag behavior is the default browser behavior: only text selections, images,
	// and links can be dragged
	// For other elements, the event ondragstart must be set for drag and drop to work
	H4Draggable_auto H4DraggableChoice = "auto"
)

// Remove the attribute draggable from the element.
func (e *H4Element) DRAGGABLERemove(c H4DraggableChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("draggable")
	return e
}

// The enterkeyhint global attribute is an enumerated attribute defining what
// action label (or icon) to present for the enter key on virtual keyboards.
func (e *H4Element) ENTERKEYHINT(c H4EnterkeyhintChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("enterkeyhint", string(c))
	return e
}

type H4EnterkeyhintChoice string

const (
	// Typically inserting a new line.
	H4Enterkeyhint_enter H4EnterkeyhintChoice = "enter"
	// Typically meaning there is nothing more to input and the input method editor
	// (IME) will be closed.
	H4Enterkeyhint_done H4EnterkeyhintChoice = "done"
	// Typically meaning to take the user to the target of the text they typed.
	H4Enterkeyhint_go H4EnterkeyhintChoice = "go"
	// Typically meaning to take the user to the next field that will accept text.
	H4Enterkeyhint_next H4EnterkeyhintChoice = "next"
	// Typically meaning to take the user to the previous field that will accept text.
	H4Enterkeyhint_previous H4EnterkeyhintChoice = "previous"
	// Typically taking the user to the results of searching for the text they have
	// typed.
	H4Enterkeyhint_search H4EnterkeyhintChoice = "search"
	// Typically delivering the text to its target.
	H4Enterkeyhint_send H4EnterkeyhintChoice = "send"
)

// Remove the attribute enterkeyhint from the element.
func (e *H4Element) ENTERKEYHINTRemove(c H4EnterkeyhintChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("enterkeyhint")
	return e
}

// The exportparts global attribute allows you to select and style elements
// existing in nested shadow trees, by exporting their part names
// The shadow tree is an isolated structure where identifiers, classes, and styles
// cannot be reached by selectors or queries belonging to a regular DOM
// To apply a style to an element living in a shadow tree, by CSS rule created
// outside of it, part global attribute has to be used
// It has to be assigned to an element present in Shadow Tree, and its value
// should be some identifier
// Rules present outside of the shadow tree, must use the ::part pseudo-element,
// containing the same identifier as the argument
// The global attribute part makes the element visible on just a single level of
// depth
// When the shadow tree is nested, parts will be visible only to the parent of the
// shadow tree but not to its ancestor
// Exporting parts further down is exactly what exportparts attribute is for
// Attribute exportparts must be placed on a shadow Host, which is the element to
// which the shadow tree is attached
// The value of the attribute should be a comma-separated list of part names
// present in the shadow tree and which should be made available via a DOM outside
// of the current structure.
func (e *H4Element) EXPORTPARTS(s ...string) *H4Element {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = treemap.New[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("exportparts")
	if !ok {
		ds = NewDelimitedBuilder[string](",")
		e.DelimitedStrings.Set("exportparts", ds)
	}
	ds.Add(s...)
	return e
}

// Remove the attribute exportparts from the element.
func (e *H4Element) EXPORTPARTSRemove(s ...string) *H4Element {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("exportparts")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// The hidden global attribute is a Boolean attribute indicating that the element
// is not yet, or is no longer, relevant
// For example, it can be used to hide elements of the page that can't be used
// until the login process has been completed
// Note that browsers typically implement hidden until found using
// content-visibility: hidden
// This means that unlike elements in the hidden state, elements in the hidden
// until found state will have generated boxes, meaning that: the element will
// participate in page layout margin, borders, padding, and background for the
// element will be rendered
// Also, the element needs to be affected by layout containment in order to be
// revealed
// This means that if the element in the hidden until found state has a display
// value of none, contents, or inline, then the element will not be revealed by
// find in page or fragment navigation.
func (e *H4Element) HIDDEN(c H4HiddenChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("hidden", string(c))
	return e
}

type H4HiddenChoice string

const (
	// set the element to the hidden state
	// Additionally, invalid values set the element to the hidden state.
	H4Hidden_empty H4HiddenChoice = ""
	// set the element to the hidden state
	// Additionally, invalid values set the element to the hidden state.
	H4Hidden_hidden H4HiddenChoice = "hidden"
	// the element is hidden but its content will be accessible to the browser's "find
	// in page" feature or to fragment navigation
	// When these features cause a scroll to an element in a hidden until found
	// subtree, the browser will fire a beforematch event on the hidden element remove
	// the hidden attribute from the element scroll to the element
	//
	H4Hidden_until_found H4HiddenChoice = "until-found"
)

// Remove the attribute hidden from the element.
func (e *H4Element) HIDDENRemove(c H4HiddenChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("hidden")
	return e
}

// The id global attribute defines a unique identifier (ID) which must be unique
// in the whole document
// Its purpose is to identify the element when linking (using a fragment
// identifier), scripting, or styling (with CSS).
func (e *H4Element) ID(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *H4Element) IDRemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// The inert global attribute is a Boolean attribute indicating that the browser
// will ignore the element
// With the inert attribute, all of the element's flat tree descendants (such as
// modal <dialog>s) that don't otherwise escape inertness are ignored
// The inert attribute also makes the browser ignore input events sent by the
// user, including focus-related events and events from assistive technologies
// Specifically, inert does the following: Prevents the click event from being
// fired when the user clicks on the element
// Prevents the focus event from being raised by preventing the element from
// gaining focus
// Hides the element and its content from assistive technologies by excluding them
// from the accessibility tree.
func (e *H4Element) INERT() *H4Element {
	e.INERTSet(true)
	return e
}

// Set the attribute inert to the value b explicitly.
func (e *H4Element) INERTSet(b bool) *H4Element {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("inert", b)
	return e
}

// Remove the attribute inert from the element.
func (e *H4Element) INERTRemove(b bool) *H4Element {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("inert")
	return e
}

// The inputmode global attribute is an enumerated attribute that hints at the
// type of data that might be entered by the user while editing the element or its
// contents
// This allows a browser to display an appropriate virtual keyboard
// It is used primarily on <input> elements, but is usable on any element in
// contenteditable mode
// It's important to understand that the inputmode attribute doesn't cause any
// validity requirements to be enforced on input
// To require that input conforms to a particular data type, choose an appropriate
// <input> element type
// For specific guidance on choosing <input> types, see the Values section.
func (e *H4Element) INPUTMODE(c H4InputmodeChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("inputmode", string(c))
	return e
}

type H4InputmodeChoice string

const (
	// No virtual keyboard
	// For when the page implements its own keyboard input control.
	H4Inputmode_none H4InputmodeChoice = "none"
	// Standard input keyboard for the user's current locale.
	H4Inputmode_empty H4InputmodeChoice = ""
	// Standard input keyboard for the user's current locale.
	H4Inputmode_text H4InputmodeChoice = "text"
	// Fractional numeric input keyboard containing the digits and decimal separator
	// for the user's locale (typically
	// or ,)
	// Devices may or may not show a minus key (-).
	H4Inputmode_decimal H4InputmodeChoice = "decimal"
	// Numeric input keyboard, but only requires the digits 0–9
	// Devices may or may not show a minus key.
	H4Inputmode_numeric H4InputmodeChoice = "numeric"
	// A telephone keypad input, including the digits 0–9, the asterisk (*), and the
	// pound (#) key
	// Inputs that *require* a telephone number should typically use <input
	// type="tel"> instead.
	H4Inputmode_tel H4InputmodeChoice = "tel"
	// A virtual keyboard optimized for search input
	// For instance, the return/submit key may be labeled "Search", along with
	// possible other optimizations
	// Inputs that require a search query should typically use <input type="search">
	// instead.
	H4Inputmode_search H4InputmodeChoice = "search"
	// A virtual keyboard optimized for entering email addresses
	// Typically includes the @character as well as other optimizations
	// Inputs that require email addresses should typically use <input type="email">
	// instead.
	H4Inputmode_email H4InputmodeChoice = "email"
	// A keypad optimized for entering URLs
	// This may have the / key more prominent, for example
	// Enhanced features could include history access and so on
	// Inputs that require a URL should typically use <input type="url"> instead.
	H4Inputmode_url H4InputmodeChoice = "url"
)

// Remove the attribute inputmode from the element.
func (e *H4Element) INPUTMODERemove(c H4InputmodeChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("inputmode")
	return e
}

// The is global attribute allows you to specify that a standard HTML element
// should behave like a defined custom built-in element (see Using custom elements
// for more details)
// This attribute can only be used if the specified custom element name has been
// successfully defined in the current document, and extends the element type it
// is being applied to.
func (e *H4Element) IS(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("is", s)
	return e
}

// Remove the attribute is from the element.
func (e *H4Element) ISRemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("is")
	return e
}

// The itemid global attribute provides microdata in the form of a unique, global
// identifier of an item.
//
//	An itemid attribute can only be specified for an element that has both
//
// itemscope and itemtype attributes
// Also, itemid can only be specified on elements that possess an itemscope
// attribute whose corresponding itemtype refers to or defines a vocabulary that
// supports global identifiers
// The exact meaning of an itemtype's global identifier is provided by the
// definition of that identifier within the specified vocabulary
// The vocabulary defines whether several items with the same global identifier
// can coexist and, if so, how items with the same identifier are handled.
func (e *H4Element) ITEMID(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("itemid", s)
	return e
}

// Remove the attribute itemid from the element.
func (e *H4Element) ITEMIDRemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("itemid")
	return e
}

// The itemprop global attribute is used to add properties to an item
// Every HTML element can have an itemprop attribute specified, and an itemprop
// consists of a name-value pair
// Each name-value pair is called a property, and a group of one or more
// properties forms an item
// Property values are either a string or a URL and can be associated with a very
// wide range of elements including <audio>, <embed>, <iframe>, <img>, <link>,
// <object>, <source>, <track>, and <video>.
func (e *H4Element) ITEMPROP(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("itemprop", s)
	return e
}

// Remove the attribute itemprop from the element.
func (e *H4Element) ITEMPROPRemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("itemprop")
	return e
}

// Properties that are not descendants of an element with the itemscope attribute
// can be associated with an item using the global attribute itemref
// itemref provides a list of element IDs (not itemids) elsewhere in the document,
// with additional properties The itemref attribute can only be specified on
// elements that have an itemscope attribute specified.
func (e *H4Element) ITEMREF(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("itemref", s)
	return e
}

// Remove the attribute itemref from the element.
func (e *H4Element) ITEMREFRemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("itemref")
	return e
}

// The itemscope global attribute is used to add an item to a microdata DOM tree
// Every HTML element can have an itemscope attribute specified, and an itemscope
// consists of a name-value pair
// Each name-value pair is called a property, and a group of one or more
// properties forms an item
// Property values are either a string or a URL and can be associated with a very
// wide range of elements including <audio>, <embed>, <iframe>, <img>, <link>,
// <object>, <source>, <track>, and <video>.
func (e *H4Element) ITEMSCOPE() *H4Element {
	e.ITEMSCOPESet(true)
	return e
}

// Set the attribute itemscope to the value b explicitly.
func (e *H4Element) ITEMSCOPESet(b bool) *H4Element {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("itemscope", b)
	return e
}

// Remove the attribute itemscope from the element.
func (e *H4Element) ITEMSCOPERemove(b bool) *H4Element {
	if e.BoolAttributes == nil {
		return e
	}
	e.BoolAttributes.Del("itemscope")
	return e
}

// The itemtype global attribute is used to add types to an item
// Every HTML element can have an itemtype attribute specified, and an itemtype
// consists of a name-value pair
// Each name-value pair is called a property, and a group of one or more
// properties forms an item
// Property values are either a string or a URL and can be associated with a very
// wide range of elements including <audio>, <embed>, <iframe>, <img>, <link>,
// <object>, <source>, <track>, and <video>.
func (e *H4Element) ITEMTYPE(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("itemtype", s)
	return e
}

// Remove the attribute itemtype from the element.
func (e *H4Element) ITEMTYPERemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("itemtype")
	return e
}

// The lang global attribute helps define the language of an element: the language
// that non-editable elements are written in or the language that editable
// elements should be written in by the user
// The tag contains one single entry value in the format defines in the Tags for
// Identifying Languages (BCP47) IETF document
// xml:lang has priority over it.
func (e *H4Element) LANG(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("lang", s)
	return e
}

// Remove the attribute lang from the element.
func (e *H4Element) LANGRemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("lang")
	return e
}

// The nonce global attribute is a unique identifier used to declare inline
// scripts and style elements to be used in a specific document
// It is a cryptographic nonce (number used once) that is used by Content Security
// Policy to determine whether or not a given inline script is allowed to execute.
func (e *H4Element) NONCE(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

// Remove the attribute nonce from the element.
func (e *H4Element) NONCERemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

// The part global attribute contains a space-separated list of the part names of
// the element
// Part names allows CSS to select and style specific elements in a shadow tree
// via the ::part pseudo-element.
func (e *H4Element) PART(s ...string) *H4Element {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = treemap.New[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("part")
	if !ok {
		ds = NewDelimitedBuilder[string](" ")
		e.DelimitedStrings.Set("part", ds)
	}
	ds.Add(s...)
	return e
}

// Remove the attribute part from the element.
func (e *H4Element) PARTRemove(s ...string) *H4Element {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("part")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// The popover global attribute is used to designate an element as a popover
// element
// Popover elements are hidden via display: none until opened via an
// invoking/control element (i.e
// a <button> or <input type="button"> with a popovertarget attribute) or a
// HTMLElement.showPopover() call
// When open, popover elements will appear above all other elements in the top
// layer, and won't be influenced by parent elements' position or overflow
// styling.
func (e *H4Element) POPVER(c H4PopverChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("popver", string(c))
	return e
}

type H4PopverChoice string

const (
	// Popovers that have the auto state can be "light dismissed" by selecting outside
	// the popover area, and generally only allow one popover to be displayed
	// on-screen at a time.
	H4Popver_auto H4PopverChoice = "auto"
	// Popovers that have the auto state can be "light dismissed" by selecting outside
	// the popover area, and generally only allow one popover to be displayed
	// on-screen at a time.
	H4Popver_empty H4PopverChoice = ""
	// manual popovers must always be explicitly hidden, but allow for use cases such
	// as nested popovers in menus.
	H4Popver_manual H4PopverChoice = "manual"
)

// Remove the attribute popver from the element.
func (e *H4Element) POPVERRemove(c H4PopverChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("popver")
	return e
}

// The slot global attribute assigns a slot in a shadow DOM shadow tree to an
// element: An element with a slot attribute is assigned to the slot created by
// the <slot> element whose name attribute's value matches that slot attribute's
// value.
func (e *H4Element) SLOT(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("slot", s)
	return e
}

// Remove the attribute slot from the element.
func (e *H4Element) SLOTRemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("slot")
	return e
}

// The spellcheck global attribute is an enumerated attribute that defines whether
// the element may be checked for spelling errors
// If this attribute is not set, its default value is element-type and
// browser-defined
// This default value may also be inherited, which means that the element content
// will be checked for spelling errors only if its nearest ancestor has a
// spellcheck state of true
// Security and privacy concerns Using spellchecking can have consequences for
// users' security and privacy
// The specification does not regulate how spellchecking is done and the content
// of the element may be sent to a third party for spellchecking results (see
// enhanced spellchecking and "spell-jacking")
// You should consider setting spellcheck to false for elements that can contain
// sensitive information.
func (e *H4Element) SPELLCHECK(c H4SpellcheckChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("spellcheck", string(c))
	return e
}

type H4SpellcheckChoice string

const (
	// The element will be checked for spelling errors.
	H4Spellcheck_empty H4SpellcheckChoice = ""
	// The element will be checked for spelling errors.
	H4Spellcheck_true H4SpellcheckChoice = "true"
	// The element will not be checked for spelling errors.
	H4Spellcheck_false H4SpellcheckChoice = "false"
)

// Remove the attribute spellcheck from the element.
func (e *H4Element) SPELLCHECKRemove(c H4SpellcheckChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("spellcheck")
	return e
}

// The style global attribute is used to add styles to an element, such as color,
// font, size, and more
// Styles are written in CSS.
func (e *H4Element) STYLE(k string, v string) *H4Element {
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

// Add the attributes in the map to the element.
func (e *H4Element) STYLEMap(m map[string]string) *H4Element {
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
func (e *H4Element) STYLEPairs(pairs ...string) *H4Element {
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

// Remove the attribute style from the element.
func (e *H4Element) STYLERemove(keys ...string) *H4Element {
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

// The tabindex global attribute indicates if its element can be focused, and
// if/where it participates in sequential keyboard navigation (usually with the
// Tab key, hence the name)
// It accepts an integer as a value, with different results depending on the
// integer's value: a negative value (usually tabindex="-1") means that the
// element should be focusable, but should not be reachable via sequential
// keyboard navigation; a value of 0 (tabindex="0") means that the element should
// be focusable and reachable via sequential keyboard navigation, but its relative
// order is defined by the platform convention; a positive value means should be
// focusable and reachable via sequential keyboard navigation; its relative order
// is defined by the value of the attribute: the sequential follow the increasing
// number of the tabindex
// If several elements share the same tabindex, their relative order follows their
// relative position in the document.
func (e *H4Element) TABINDEX(i int) *H4Element {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

// Remove the attribute tabindex from the element.
func (e *H4Element) TABINDEXRemove(i int) *H4Element {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// The title global attribute contains text representing advisory information
// related to the element it belongs to
// Such information can typically, but not necessarily, be presented to the user
// as a tooltip
// The main use of the title attribute is to label <iframe> elements for assistive
// technology
// The title attribute may also be used to label controls in data tables
// The title attribute, when added to <link rel="stylesheet">, creates an
// alternate stylesheet
// When defining an alternative style sheet with <link rel="alternate"> the
// attribute is required and must be set to a non-empty string
// If included on the <abbr> opening tag, the title must be a full expansion of
// the abbreviation or acronym
// Instead of using title, when possible, provide an expansion of the abbreviation
// or acronym in plain text on first use, using the <abbr> to mark up the
// abbreviation
// This enables all users know what name or term the abbreviation or acronym
// shortens while providing a hint to user agents on how to announce the content
// While title can be used to provide a programmatically associated label for an
// <input> element, this is not good practice
// Use a <label> instead.
func (e *H4Element) TITLE(s string) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("title", s)
	return e
}

// Remove the attribute title from the element.
func (e *H4Element) TITLERemove(s string) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("title")
	return e
}

// The translate global attribute is an enumerated attribute that is used to
// specify whether an element's attribute values and the values of its Text node
// children are to be translated when the page is localized, or whether to leave
// them unchanged.
func (e *H4Element) TRANSLATE(c H4TranslateChoice) *H4Element {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("translate", string(c))
	return e
}

type H4TranslateChoice string

const (
	// indicates that the element should be translated when the page is localized.
	H4Translate_empty H4TranslateChoice = ""
	// indicates that the element should be translated when the page is localized.
	H4Translate_yes H4TranslateChoice = "yes"
	// indicates that the element must not be translated when the page is localized.
	H4Translate_no H4TranslateChoice = "no"
)

// Remove the attribute translate from the element.
func (e *H4Element) TRANSLATERemove(c H4TranslateChoice) *H4Element {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("translate")
	return e
}