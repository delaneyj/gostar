
/* cSpell:disable */

package  elements

type InputElement struct {
    *Element
}

func INPUT() *InputElement {
    return &InputElement{
        Element: &Element{
            Tag: "input",
            IsSelfClosing: true,
        },
    }
}

func (e *InputElement) AddChildren(children ...*Element) *InputElement {
    e.Children = append(e.Children, children...)
	return e
}


// Accept is the "accept" attribute.
// Hint for expected file type in file upload controls
// Valid values are constrained to the following:
//   * set_of_comma_separated_tokens
//   * valid_mime_type_strings_with_no_parameters
//   * audio/*
//   * video/*
//   * image/*
func (e *InputElement) Accept(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["accept"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AcceptAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["accept"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["accept"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AcceptRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["accept"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   * ordered_set_of_unique_space_separated_tokens
//   * identical_to
func (e *InputElement) Accesskey(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["accesskey"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AccesskeyAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["accesskey"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["accesskey"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AccesskeyRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["accesskey"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Alt is the "alt" attribute.
// Replacement text for use when images are not available
// Valid values are constrained to the following:
//   * text
func (e *InputElement) Alt(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["alt"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AltAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["alt"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["alt"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AltRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["alt"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Autocapitalize is the "autocapitalize" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Valid values are constrained to the following:
//   * on
//   * on
//   * off
//   * off
//   * none
//   * none
//   * sentences
//   * sentences
//   * words
//   * words
//   * characters
//   * characters
func (e *InputElement) Autocapitalize(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AutocapitalizeAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocapitalize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AutocapitalizeRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocapitalize"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Autocomplete is the "autocomplete" attribute.
// Hint for form autofill feature
// Valid values are constrained to the following:
//   * autofill_field
func (e *InputElement) Autocomplete(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["autocomplete"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AutocompleteAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocomplete"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["autocomplete"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) AutocompleteRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocomplete"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *InputElement) Autofocus() *InputElement {
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *InputElement) AutofocusRemove() *InputElement {
    delete(e.BoolAttributes,"autofocus")
    return e
}

func (e *InputElement) AutofocusSet(v bool) *InputElement {
    if v {
        e.Autofocus()
    } else {
        e.AutofocusRemove()
    }
    return e
}


// Checked is the "checked" attribute.
// Whether the control is checked
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *InputElement) Checked() *InputElement {
    e.BoolAttributes["checked"] = struct{}{}
    return e
}

func (e *InputElement) CheckedRemove() *InputElement {
    delete(e.BoolAttributes,"checked")
    return e
}

func (e *InputElement) CheckedSet(v bool) *InputElement {
    if v {
        e.Checked()
    } else {
        e.CheckedRemove()
    }
    return e
}


// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   * set_of_space_separated_tokens
func (e *InputElement) Class(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ClassAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["class"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ClassRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["class"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   * true
//   * plaintext_only
//   * false
func (e *InputElement) Contenteditable(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ContenteditableAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["contenteditable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ContenteditableRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["contenteditable"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   * ltr
//   * ltr
//   * rtl
//   * rtl
func (e *InputElement) Dir(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["dir"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) DirAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dir"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["dir"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) DirRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dir"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Dirname is the "dirname" attribute.
// Name of form control to use for sending the element&#39;s directionality in form submission
// Valid values are constrained to the following:
//   * text
func (e *InputElement) Dirname(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["dirname"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) DirnameAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dirname"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["dirname"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) DirnameRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dirname"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Disabled is the "disabled" attribute.
// Whether the link is disabled
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *InputElement) Disabled() *InputElement {
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *InputElement) DisabledRemove() *InputElement {
    delete(e.BoolAttributes,"disabled")
    return e
}

func (e *InputElement) DisabledSet(v bool) *InputElement {
    if v {
        e.Disabled()
    } else {
        e.DisabledRemove()
    }
    return e
}


// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   * true
//   * false
func (e *InputElement) Draggable(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["draggable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) DraggableAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["draggable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["draggable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) DraggableRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["draggable"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Enterkeyhint is the "enterkeyhint" attribute.
// Hint for selecting an enter key action
// Valid values are constrained to the following:
//   * enter
//   * enter
//   * done
//   * done
//   * go
//   * go
//   * next
//   * next
//   * previous
//   * previous
//   * search
//   * search
//   * send
//   * send
func (e *InputElement) Enterkeyhint(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) EnterkeyhintAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["enterkeyhint"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) EnterkeyhintRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["enterkeyhint"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Form is the "form" attribute.
// Associates the element with a form element
// Valid values are constrained to the following:
//   * id
func (e *InputElement) Form(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["form"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["form"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["form"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["form"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Formaction is the "formaction" attribute.
// URL to use for form submission
// Valid values are constrained to the following:
//   * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *InputElement) Formaction(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["formaction"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormactionAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["formaction"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["formaction"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormactionRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["formaction"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Formenctype is the "formenctype" attribute.
// Entry list encoding type to use for form submission
// Valid values are constrained to the following:
//   * application/x_www_form_urlencoded
//   * application/x_www_form_urlencoded
//   * multipart/form_data
//   * multipart/form_data
//   * text/plain
//   * text/plain
func (e *InputElement) Formenctype(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["formenctype"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormenctypeAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["formenctype"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["formenctype"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormenctypeRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["formenctype"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Formmethod is the "formmethod" attribute.
// Variant to use for form submission
// Valid values are constrained to the following:
//   * get
//   * post
//   * dialog
func (e *InputElement) Formmethod(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["formmethod"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormmethodAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["formmethod"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["formmethod"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormmethodRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["formmethod"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Formnovalidate is the "formnovalidate" attribute.
// Bypass form control validation for form submission
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *InputElement) Formnovalidate() *InputElement {
    e.BoolAttributes["formnovalidate"] = struct{}{}
    return e
}

func (e *InputElement) FormnovalidateRemove() *InputElement {
    delete(e.BoolAttributes,"formnovalidate")
    return e
}

func (e *InputElement) FormnovalidateSet(v bool) *InputElement {
    if v {
        e.Formnovalidate()
    } else {
        e.FormnovalidateRemove()
    }
    return e
}


// Formtarget is the "formtarget" attribute.
// Navigable for form submission
// Valid values are constrained to the following:
//   * valid_navigable_target_name_or_keyword
func (e *InputElement) Formtarget(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["formtarget"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormtargetAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["formtarget"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["formtarget"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) FormtargetRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["formtarget"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Height is the "height" attribute.
// Vertical dimension
// Valid values are constrained to the following:
//   * valid_non_negative_integer
func (e *InputElement) Height(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["height"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) HeightAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["height"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["height"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) HeightRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["height"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   * until_found
//   * until_found
//   * hidden
//   * hidden
func (e *InputElement) Hidden(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["hidden"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) HiddenAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["hidden"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["hidden"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) HiddenRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["hidden"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   * text
func (e *InputElement) Id(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["id"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) IdAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["id"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["id"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) IdRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["id"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *InputElement) Inert() *InputElement {
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *InputElement) InertRemove() *InputElement {
    delete(e.BoolAttributes,"inert")
    return e
}

func (e *InputElement) InertSet(v bool) *InputElement {
    if v {
        e.Inert()
    } else {
        e.InertRemove()
    }
    return e
}


// Inputmode is the "inputmode" attribute.
// Hint for selecting an input modality
// Valid values are constrained to the following:
//   * none
//   * none
//   * text
//   * text
//   * tel
//   * tel
//   * email
//   * email
//   * url
//   * url
//   * numeric
//   * numeric
//   * decimal
//   * decimal
//   * search
//   * search
func (e *InputElement) Inputmode(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["inputmode"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) InputmodeAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["inputmode"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["inputmode"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) InputmodeRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["inputmode"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   * valid_custom_element_name
//   * customized_built_in_element
func (e *InputElement) Is(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["is"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) IsAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["is"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["is"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) IsRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["is"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   * valid_url_potentially_surrounded_by_spaces
func (e *InputElement) Itemid(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemid"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ItemidAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemid"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ItemidRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemid"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   * unordered_set_of_unique_space_separated_tokens
//   * valid_absolute_ur_ls
//   * defined_property_names
func (e *InputElement) Itemprop(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemprop"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ItempropAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemprop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemprop"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ItempropRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemprop"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   * unordered_set_of_unique_space_separated_tokens
func (e *InputElement) Itemref(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemref"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ItemrefAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemref"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemref"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ItemrefRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemref"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *InputElement) Itemscope() *InputElement {
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *InputElement) ItemscopeRemove() *InputElement {
    delete(e.BoolAttributes,"itemscope")
    return e
}

func (e *InputElement) ItemscopeSet(v bool) *InputElement {
    if v {
        e.Itemscope()
    } else {
        e.ItemscopeRemove()
    }
    return e
}


// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   * unordered_set_of_unique_space_separated_tokens
//   * valid_absolute_ur_ls
func (e *InputElement) Itemtype(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemtype"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ItemtypeAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemtype"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ItemtypeRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (e *InputElement) Lang(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["lang"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) LangAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["lang"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["lang"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) LangRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["lang"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// List is the "list" attribute.
// List of autocomplete options
// Valid values are constrained to the following:
//   * id
func (e *InputElement) List(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["list"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ListAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["list"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["list"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ListRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["list"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Max is the "max" attribute.
// Upper bound of range
// Valid values are constrained to the following:
//   * valid_floating_point_number
func (e *InputElement) Max(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["max"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) MaxAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["max"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["max"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) MaxRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["max"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Maxlength is the "maxlength" attribute.
// Maximum length of value
// Valid values are constrained to the following:
//   * valid_non_negative_integer
func (e *InputElement) Maxlength(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["maxlength"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) MaxlengthAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["maxlength"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["maxlength"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) MaxlengthRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["maxlength"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Min is the "min" attribute.
// Lower bound of range
// Valid values are constrained to the following:
//   * valid_floating_point_number
func (e *InputElement) Min(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["min"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) MinAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["min"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["min"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) MinRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["min"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Minlength is the "minlength" attribute.
// Minimum length of value
// Valid values are constrained to the following:
//   * valid_non_negative_integer
func (e *InputElement) Minlength(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["minlength"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) MinlengthAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["minlength"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["minlength"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) MinlengthRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["minlength"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Multiple is the "multiple" attribute.
// Whether to allow multiple values
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *InputElement) Multiple() *InputElement {
    e.BoolAttributes["multiple"] = struct{}{}
    return e
}

func (e *InputElement) MultipleRemove() *InputElement {
    delete(e.BoolAttributes,"multiple")
    return e
}

func (e *InputElement) MultipleSet(v bool) *InputElement {
    if v {
        e.Multiple()
    } else {
        e.MultipleRemove()
    }
    return e
}


// Name is the "name" attribute.
// Name of shadow tree slot
// Valid values are constrained to the following:
//   * text
func (e *InputElement) Name(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["name"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) NameAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["name"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["name"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) NameRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["name"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   * text
func (e *InputElement) Nonce(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["nonce"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) NonceAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["nonce"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["nonce"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) NonceRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["nonce"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Pattern is the "pattern" attribute.
// Pattern to be matched by the form control&#39;s value
// Valid values are constrained to the following:
//   * pattern
func (e *InputElement) Pattern(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["pattern"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PatternAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["pattern"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["pattern"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PatternRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["pattern"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Placeholder is the "placeholder" attribute.
// User-visible label to be placed within the form control
// Valid values are constrained to the following:
//   * text
func (e *InputElement) Placeholder(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["placeholder"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PlaceholderAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["placeholder"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["placeholder"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PlaceholderRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["placeholder"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   * auto
//   * auto
//   * manual
//   * manual
func (e *InputElement) Popover(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["popover"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PopoverAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["popover"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PopoverRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popover"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Popovertarget is the "popovertarget" attribute.
// Targets a popover element to toggle, show, or hide
// Valid values are constrained to the following:
//   * id
func (e *InputElement) Popovertarget(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["popovertarget"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PopovertargetAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popovertarget"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["popovertarget"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PopovertargetRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popovertarget"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Popovertargetaction is the "popovertargetaction" attribute.
// Indicates whether a targeted popover element is to be toggled, shown, or hidden
// Valid values are constrained to the following:
//   * toggle
//   * toggle
//   * show
//   * show
//   * hide
//   * hide
func (e *InputElement) Popovertargetaction(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["popovertargetaction"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PopovertargetactionAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popovertargetaction"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["popovertargetaction"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) PopovertargetactionRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popovertargetaction"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Readonly is the "readonly" attribute.
// Affects willValidate, plus any behavior added by the custom element author
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *InputElement) Readonly() *InputElement {
    e.BoolAttributes["readonly"] = struct{}{}
    return e
}

func (e *InputElement) ReadonlyRemove() *InputElement {
    delete(e.BoolAttributes,"readonly")
    return e
}

func (e *InputElement) ReadonlySet(v bool) *InputElement {
    if v {
        e.Readonly()
    } else {
        e.ReadonlyRemove()
    }
    return e
}


// Required is the "required" attribute.
// Whether the control is required for form submission
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *InputElement) Required() *InputElement {
    e.BoolAttributes["required"] = struct{}{}
    return e
}

func (e *InputElement) RequiredRemove() *InputElement {
    delete(e.BoolAttributes,"required")
    return e
}

func (e *InputElement) RequiredSet(v bool) *InputElement {
    if v {
        e.Required()
    } else {
        e.RequiredRemove()
    }
    return e
}


// Size is the "size" attribute.
// Size of the control
// Valid values are constrained to the following:
//   * valid_non_negative_integer
func (e *InputElement) Size(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["size"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) SizeAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["size"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["size"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) SizeRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["size"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   * text
func (e *InputElement) Slot(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["slot"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) SlotAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["slot"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["slot"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) SlotRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["slot"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   * true
//   * false
func (e *InputElement) Spellcheck(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) SpellcheckAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) SpellcheckRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Src is the "src" attribute.
// Address of the resource
// Valid values are constrained to the following:
//   * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *InputElement) Src(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["src"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) SrcAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["src"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["src"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) SrcRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["src"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Step is the "step" attribute.
// Granularity to be matched by the form control&#39;s value
// Valid values are constrained to the following:
//   * valid_floating_point_number
//   * any
func (e *InputElement) Step(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["step"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) StepAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["step"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["step"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) StepRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["step"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (e *InputElement) Style(k,v string) *InputElement {
    htmlStringer,ok := e.DelimitedKVAttributes["style"]
    if !ok {
        htmlStringer = NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = htmlStringer
    }
    htmlStringer.Add(k,v)
    return e
}

func (e *InputElement) StyleRemove(k string) *InputElement {
    htmlStringer,ok := e.DelimitedKVAttributes["style"]
    if ok {
        htmlStringer.Remove(k)
    }
    return e
}


// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   * valid_integer
func (e *InputElement) Tabindex(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["tabindex"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) TabindexAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["tabindex"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["tabindex"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) TabindexRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["tabindex"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   * text
func (e *InputElement) Title(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["title"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) TitleAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["title"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["title"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) TitleRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["title"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   * yes
//   * no
func (e *InputElement) Translate(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["translate"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) TranslateAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["translate"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) TranslateRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Type is the "type" attribute.
// Type of script
// Valid values are constrained to the following:
//   * module
//   * valid_mime_type_string
//   * java_script_mime_type_essence_match
func (e *InputElement) Type(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["type"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) TypeAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["type"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["type"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) TypeRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["type"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Value is the "value" attribute.
// Current value of the element
// Valid values are constrained to the following:
//   * valid_floating_point_number
func (e *InputElement) Value(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["value"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ValueAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["value"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["value"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) ValueRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["value"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Width is the "width" attribute.
// Horizontal dimension
// Valid values are constrained to the following:
//   * valid_non_negative_integer
func (e *InputElement) Width(v string) *InputElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["width"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) WidthAdd(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["width"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["width"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *InputElement) WidthRemove(v string) *InputElement {
    htmlStringer,ok := e.DelimitedStringAttributes["width"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}




func (e *InputElement) OnAuxclick(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onauxclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnAuxclickRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnBeforematch(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforematch"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnBeforematchRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnBeforetoggle(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforetoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnBeforetoggleRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnBlur(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onblur"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnBlurRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnCancel(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncancel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnCancelRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnCanplay(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnCanplayRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnCanplaythrough(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplaythrough"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnCanplaythroughRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnChange(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnChangeRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnClick(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnClickRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnClose(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclose"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnCloseRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnContextlost(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextlost"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnContextlostRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnContextmenu(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextmenu"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnContextmenuRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnContextrestored(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextrestored"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnContextrestoredRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnCopy(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncopy"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnCopyRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnCuechange(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncuechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnCuechangeRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnCut(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncut"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnCutRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnDblclick(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondblclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnDblclickRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnDrag(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrag"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnDragRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnDragend(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnDragendRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnDragenter(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnDragenterRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnDragleave(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnDragleaveRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnDragover(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnDragoverRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnDragstart(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnDragstartRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnDrop(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrop"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnDropRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnDurationchange(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondurationchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnDurationchangeRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnEmptied(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onemptied"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnEmptiedRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnEnded(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onended"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnEndedRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnError(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onerror"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnErrorRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnFocus(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onfocus"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnFocusRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnFormdata(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onformdata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnFormdataRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnInput(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninput"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnInputRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnInvalid(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninvalid"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnInvalidRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnKeydown(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeydown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnKeydownRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnKeypress(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeypress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnKeypressRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnKeyup(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeyup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnKeyupRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnLoad(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onload"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnLoadRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnLoadeddata(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadeddata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnLoadeddataRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnLoadedmetadata(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadedmetadata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnLoadedmetadataRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnLoadstart(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnLoadstartRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnMousedown(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousedown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnMousedownRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnMouseenter(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnMouseenterRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnMouseleave(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnMouseleaveRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnMousemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousemove"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnMousemoveRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnMouseout(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseout"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnMouseoutRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnMouseover(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnMouseoverRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnMouseup(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnMouseupRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnPaste(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpaste"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnPasteRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnPause(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpause"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnPauseRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnPlay(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnPlayRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnPlaying(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplaying"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnPlayingRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnProgress(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onprogress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnProgressRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnRatechange(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onratechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnRatechangeRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnReset(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onreset"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnResetRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnResize(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onresize"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnResizeRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnScroll(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscroll"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnScrollRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnScrollend(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscrollend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnScrollendRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnSecuritypolicyviolation(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsecuritypolicyviolation"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnSecuritypolicyviolationRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnSeeked(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeked"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnSeekedRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnSeeking(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeking"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnSeekingRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnSelect(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onselect"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnSelectRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnSlotchange(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onslotchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnSlotchangeRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnStalled(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onstalled"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnStalledRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnSubmit(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsubmit"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnSubmitRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnSuspend(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsuspend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnSuspendRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnTimeupdate(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontimeupdate"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnTimeupdateRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnToggle(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnToggleRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnVolumechange(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onvolumechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnVolumechangeRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnWaiting(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwaiting"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnWaitingRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *InputElement) OnWheel(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwheel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *InputElement) OnWheelRemove(evt string) (*InputElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}


