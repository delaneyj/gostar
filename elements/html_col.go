// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package html col is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The HTML <col> element defines a column within a table and is used for defining
// common semantics on all common cells
// It is generally found within a <colgroup> element.
type COLElement struct {
	*Element
}

// Create a new COLElement element.
// This will create a new element with the tag
// "col" during rendering.
func COL() *COLElement {
	e := NewElement("col")
	e.IsSelfClosing = true

	return &COLElement{Element: e}
}

func (e *COLElement) Children(children ...ElementRenderer) *COLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *COLElement) IfChildren(condition bool, children ...ElementRenderer) *COLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *COLElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *COLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *COLElement) Text(text string) *COLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *COLElement) TextF(format string, args ...any) *COLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *COLElement) Escaped(text string) *COLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *COLElement) EscapedF(format string, args ...any) *COLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *COLElement) CustomData(key, value string) *COLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *COLElement) CustomDataF(key, format string, args ...any) *COLElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *COLElement) CustomDataRemove(key string) *COLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// How many columns this column element spans.
func (e *COLElement) SPAN(i int) *COLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("span", i)
	return e
}

// Remove the attribute span from the element.
func (e *COLElement) SPANRemove(i int) *COLElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("span")
	return e
}

// The accesskey global attribute provides a hint for generating a keyboard
// shortcut for the current element
// The attribute value must consist of a single printable character (which
// includes accented and other characters that can be generated by the keyboard).
func (e *COLElement) ACCESSKEY(r rune) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("accesskey", string(r))
	return e
}

// Remove the attribute accesskey from the element.
func (e *COLElement) ACCESSKEYRemove() *COLElement {
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
func (e *COLElement) AUTOCAPITALIZE(c ColAutocapitalizeChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("autocapitalize", string(c))
	return e
}

type ColAutocapitalizeChoice string

const (
	// Do not automatically capitalize any text.
	ColAutocapitalize_off ColAutocapitalizeChoice = "off"
	// Do not automatically capitalize any text.
	ColAutocapitalize_none ColAutocapitalizeChoice = "none"
	// Automatically capitalize the first character of each sentence.
	ColAutocapitalize_sentences ColAutocapitalizeChoice = "sentences"
	// Automatically capitalize the first character of each sentence.
	ColAutocapitalize_on ColAutocapitalizeChoice = "on"
	// Automatically capitalize the first character of each word.
	ColAutocapitalize_words ColAutocapitalizeChoice = "words"
	// Automatically capitalize all characters.
	ColAutocapitalize_characters ColAutocapitalizeChoice = "characters"
)

// Remove the attribute autocapitalize from the element.
func (e *COLElement) AUTOCAPITALIZERemove(c ColAutocapitalizeChoice) *COLElement {
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
func (e *COLElement) AUTOFOCUS() *COLElement {
	e.AUTOFOCUSSet(true)
	return e
}

// Set the attribute autofocus to the value b explicitly.
func (e *COLElement) AUTOFOCUSSet(b bool) *COLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("autofocus", b)
	return e
}

// Remove the attribute autofocus from the element.
func (e *COLElement) AUTOFOCUSRemove(b bool) *COLElement {
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
func (e *COLElement) CLASS(s ...string) *COLElement {
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
func (e *COLElement) CLASSRemove(s ...string) *COLElement {
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
func (e *COLElement) CONTENTEDITABLE(c ColContenteditableChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("contenteditable", string(c))
	return e
}

type ColContenteditableChoice string

const (
	// The element is editable.
	ColContenteditable_empty ColContenteditableChoice = ""
	// The element is editable.
	ColContenteditable_true ColContenteditableChoice = "true"
	// The element is not editable.
	ColContenteditable_false ColContenteditableChoice = "false"
	// which indicates that the element's raw text is editable, but rich text
	// formatting is disabled.
	ColContenteditable_plaintext_only ColContenteditableChoice = "plaintext-only"
)

// Remove the attribute contenteditable from the element.
func (e *COLElement) CONTENTEDITABLERemove(c ColContenteditableChoice) *COLElement {
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
func (e *COLElement) DIR(c ColDirChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type ColDirChoice string

const (
	// which means left to right and is to be used for languages that are written from
	// the left to the right (like English);
	ColDir_ltr ColDirChoice = "ltr"
	// which means right to left and is to be used for languages that are written from
	// the right to the left (like Arabic);
	ColDir_rtl ColDirChoice = "rtl"
	// which lets the user agent decide
	// It uses a basic algorithm as it parses the characters inside the element until
	// it finds a character with a strong directionality, then it applies that
	// directionality to the whole element.
	ColDir_auto ColDirChoice = "auto"
)

// Remove the attribute dir from the element.
func (e *COLElement) DIRRemove(c ColDirChoice) *COLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// The draggable global attribute is an enumerated attribute that indicates
// whether the element can be dragged, either with native browser behavior or the
// HTML Drag and Drop API.
func (e *COLElement) DRAGGABLE(c ColDraggableChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("draggable", string(c))
	return e
}

type ColDraggableChoice string

const (
	// The element is draggable.
	ColDraggable_true ColDraggableChoice = "true"
	// The element is not draggable.
	ColDraggable_false ColDraggableChoice = "false"
	// drag behavior is the default browser behavior: only text selections, images,
	// and links can be dragged
	// For other elements, the event ondragstart must be set for drag and drop to work
	ColDraggable_empty ColDraggableChoice = ""
	// drag behavior is the default browser behavior: only text selections, images,
	// and links can be dragged
	// For other elements, the event ondragstart must be set for drag and drop to work
	ColDraggable_auto ColDraggableChoice = "auto"
)

// Remove the attribute draggable from the element.
func (e *COLElement) DRAGGABLERemove(c ColDraggableChoice) *COLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("draggable")
	return e
}

// The enterkeyhint global attribute is an enumerated attribute defining what
// action label (or icon) to present for the enter key on virtual keyboards.
func (e *COLElement) ENTERKEYHINT(c ColEnterkeyhintChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("enterkeyhint", string(c))
	return e
}

type ColEnterkeyhintChoice string

const (
	// Typically inserting a new line.
	ColEnterkeyhint_enter ColEnterkeyhintChoice = "enter"
	// Typically meaning there is nothing more to input and the input method editor
	// (IME) will be closed.
	ColEnterkeyhint_done ColEnterkeyhintChoice = "done"
	// Typically meaning to take the user to the target of the text they typed.
	ColEnterkeyhint_go ColEnterkeyhintChoice = "go"
	// Typically meaning to take the user to the next field that will accept text.
	ColEnterkeyhint_next ColEnterkeyhintChoice = "next"
	// Typically meaning to take the user to the previous field that will accept text.
	ColEnterkeyhint_previous ColEnterkeyhintChoice = "previous"
	// Typically taking the user to the results of searching for the text they have
	// typed.
	ColEnterkeyhint_search ColEnterkeyhintChoice = "search"
	// Typically delivering the text to its target.
	ColEnterkeyhint_send ColEnterkeyhintChoice = "send"
)

// Remove the attribute enterkeyhint from the element.
func (e *COLElement) ENTERKEYHINTRemove(c ColEnterkeyhintChoice) *COLElement {
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
func (e *COLElement) EXPORTPARTS(s ...string) *COLElement {
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
func (e *COLElement) EXPORTPARTSRemove(s ...string) *COLElement {
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
func (e *COLElement) HIDDEN(c ColHiddenChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("hidden", string(c))
	return e
}

type ColHiddenChoice string

const (
	// set the element to the hidden state
	// Additionally, invalid values set the element to the hidden state.
	ColHidden_empty ColHiddenChoice = ""
	// set the element to the hidden state
	// Additionally, invalid values set the element to the hidden state.
	ColHidden_hidden ColHiddenChoice = "hidden"
	// the element is hidden but its content will be accessible to the browser's "find
	// in page" feature or to fragment navigation
	// When these features cause a scroll to an element in a hidden until found
	// subtree, the browser will fire a beforematch event on the hidden element remove
	// the hidden attribute from the element scroll to the element
	//
	ColHidden_until_found ColHiddenChoice = "until-found"
)

// Remove the attribute hidden from the element.
func (e *COLElement) HIDDENRemove(c ColHiddenChoice) *COLElement {
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
func (e *COLElement) ID(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *COLElement) IDRemove(s string) *COLElement {
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
func (e *COLElement) INERT() *COLElement {
	e.INERTSet(true)
	return e
}

// Set the attribute inert to the value b explicitly.
func (e *COLElement) INERTSet(b bool) *COLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("inert", b)
	return e
}

// Remove the attribute inert from the element.
func (e *COLElement) INERTRemove(b bool) *COLElement {
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
func (e *COLElement) INPUTMODE(c ColInputmodeChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("inputmode", string(c))
	return e
}

type ColInputmodeChoice string

const (
	// No virtual keyboard
	// For when the page implements its own keyboard input control.
	ColInputmode_none ColInputmodeChoice = "none"
	// Standard input keyboard for the user's current locale.
	ColInputmode_empty ColInputmodeChoice = ""
	// Standard input keyboard for the user's current locale.
	ColInputmode_text ColInputmodeChoice = "text"
	// Fractional numeric input keyboard containing the digits and decimal separator
	// for the user's locale (typically
	// or ,)
	// Devices may or may not show a minus key (-).
	ColInputmode_decimal ColInputmodeChoice = "decimal"
	// Numeric input keyboard, but only requires the digits 0–9
	// Devices may or may not show a minus key.
	ColInputmode_numeric ColInputmodeChoice = "numeric"
	// A telephone keypad input, including the digits 0–9, the asterisk (*), and the
	// pound (#) key
	// Inputs that *require* a telephone number should typically use <input
	// type="tel"> instead.
	ColInputmode_tel ColInputmodeChoice = "tel"
	// A virtual keyboard optimized for search input
	// For instance, the return/submit key may be labeled "Search", along with
	// possible other optimizations
	// Inputs that require a search query should typically use <input type="search">
	// instead.
	ColInputmode_search ColInputmodeChoice = "search"
	// A virtual keyboard optimized for entering email addresses
	// Typically includes the @character as well as other optimizations
	// Inputs that require email addresses should typically use <input type="email">
	// instead.
	ColInputmode_email ColInputmodeChoice = "email"
	// A keypad optimized for entering URLs
	// This may have the / key more prominent, for example
	// Enhanced features could include history access and so on
	// Inputs that require a URL should typically use <input type="url"> instead.
	ColInputmode_url ColInputmodeChoice = "url"
)

// Remove the attribute inputmode from the element.
func (e *COLElement) INPUTMODERemove(c ColInputmodeChoice) *COLElement {
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
func (e *COLElement) IS(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("is", s)
	return e
}

// Remove the attribute is from the element.
func (e *COLElement) ISRemove(s string) *COLElement {
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
func (e *COLElement) ITEMID(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("itemid", s)
	return e
}

// Remove the attribute itemid from the element.
func (e *COLElement) ITEMIDRemove(s string) *COLElement {
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
func (e *COLElement) ITEMPROP(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("itemprop", s)
	return e
}

// Remove the attribute itemprop from the element.
func (e *COLElement) ITEMPROPRemove(s string) *COLElement {
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
func (e *COLElement) ITEMREF(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("itemref", s)
	return e
}

// Remove the attribute itemref from the element.
func (e *COLElement) ITEMREFRemove(s string) *COLElement {
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
func (e *COLElement) ITEMSCOPE() *COLElement {
	e.ITEMSCOPESet(true)
	return e
}

// Set the attribute itemscope to the value b explicitly.
func (e *COLElement) ITEMSCOPESet(b bool) *COLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = treemap.New[string, bool]()
	}
	e.BoolAttributes.Set("itemscope", b)
	return e
}

// Remove the attribute itemscope from the element.
func (e *COLElement) ITEMSCOPERemove(b bool) *COLElement {
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
func (e *COLElement) ITEMTYPE(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("itemtype", s)
	return e
}

// Remove the attribute itemtype from the element.
func (e *COLElement) ITEMTYPERemove(s string) *COLElement {
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
func (e *COLElement) LANG(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("lang", s)
	return e
}

// Remove the attribute lang from the element.
func (e *COLElement) LANGRemove(s string) *COLElement {
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
func (e *COLElement) NONCE(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

// Remove the attribute nonce from the element.
func (e *COLElement) NONCERemove(s string) *COLElement {
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
func (e *COLElement) PART(s ...string) *COLElement {
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
func (e *COLElement) PARTRemove(s ...string) *COLElement {
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
func (e *COLElement) POPVER(c ColPopverChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("popver", string(c))
	return e
}

type ColPopverChoice string

const (
	// Popovers that have the auto state can be "light dismissed" by selecting outside
	// the popover area, and generally only allow one popover to be displayed
	// on-screen at a time.
	ColPopver_auto ColPopverChoice = "auto"
	// Popovers that have the auto state can be "light dismissed" by selecting outside
	// the popover area, and generally only allow one popover to be displayed
	// on-screen at a time.
	ColPopver_empty ColPopverChoice = ""
	// manual popovers must always be explicitly hidden, but allow for use cases such
	// as nested popovers in menus.
	ColPopver_manual ColPopverChoice = "manual"
)

// Remove the attribute popver from the element.
func (e *COLElement) POPVERRemove(c ColPopverChoice) *COLElement {
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
func (e *COLElement) SLOT(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("slot", s)
	return e
}

// Remove the attribute slot from the element.
func (e *COLElement) SLOTRemove(s string) *COLElement {
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
func (e *COLElement) SPELLCHECK(c ColSpellcheckChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("spellcheck", string(c))
	return e
}

type ColSpellcheckChoice string

const (
	// The element will be checked for spelling errors.
	ColSpellcheck_empty ColSpellcheckChoice = ""
	// The element will be checked for spelling errors.
	ColSpellcheck_true ColSpellcheckChoice = "true"
	// The element will not be checked for spelling errors.
	ColSpellcheck_false ColSpellcheckChoice = "false"
)

// Remove the attribute spellcheck from the element.
func (e *COLElement) SPELLCHECKRemove(c ColSpellcheckChoice) *COLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("spellcheck")
	return e
}

// The style global attribute is used to add styles to an element, such as color,
// font, size, and more
// Styles are written in CSS.
func (e *COLElement) STYLEF(k string, format string, args ...any) *COLElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *COLElement) STYLE(k string, v string) *COLElement {
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
func (e *COLElement) STYLEMap(m map[string]string) *COLElement {
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
func (e *COLElement) STYLEPairs(pairs ...string) *COLElement {
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
func (e *COLElement) STYLERemove(keys ...string) *COLElement {
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
func (e *COLElement) TABINDEX(i int) *COLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

// Remove the attribute tabindex from the element.
func (e *COLElement) TABINDEXRemove(i int) *COLElement {
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
func (e *COLElement) TITLE(s string) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("title", s)
	return e
}

// Remove the attribute title from the element.
func (e *COLElement) TITLERemove(s string) *COLElement {
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
func (e *COLElement) TRANSLATE(c ColTranslateChoice) *COLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("translate", string(c))
	return e
}

type ColTranslateChoice string

const (
	// indicates that the element should be translated when the page is localized.
	ColTranslate_empty ColTranslateChoice = ""
	// indicates that the element should be translated when the page is localized.
	ColTranslate_yes ColTranslateChoice = "yes"
	// indicates that the element must not be translated when the page is localized.
	ColTranslate_no ColTranslateChoice = "no"
)

// Remove the attribute translate from the element.
func (e *COLElement) TRANSLATERemove(c ColTranslateChoice) *COLElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("translate")
	return e
}
