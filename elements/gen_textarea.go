
/* cSpell:disable */

package  elements

type TextareaElement struct {
    *Element
}

func TEXTAREA() *TextareaElement {
    return &TextareaElement{
        Element: &Element{
            Tag: "textarea",
            IsSelfClosing: false,
        },
    }
}

func (e *TextareaElement) AddChildren(children ...*Element) *TextareaElement {
    e.Children = append(e.Children, children...)
	return e
}


// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   * ordered_set_of_unique_space_separated_tokens
//   * identical_to
func (e *TextareaElement) Accesskey(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["accesskey"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) AccesskeyAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["accesskey"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["accesskey"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) AccesskeyRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["accesskey"]
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
func (e *TextareaElement) Autocapitalize(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) AutocapitalizeAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocapitalize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) AutocapitalizeRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Autocomplete(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["autocomplete"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) AutocompleteAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocomplete"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["autocomplete"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) AutocompleteRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Autofocus() *TextareaElement {
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *TextareaElement) AutofocusRemove() *TextareaElement {
    delete(e.BoolAttributes,"autofocus")
    return e
}

func (e *TextareaElement) AutofocusSet(v bool) *TextareaElement {
    if v {
        e.Autofocus()
    } else {
        e.AutofocusRemove()
    }
    return e
}


// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   * set_of_space_separated_tokens
func (e *TextareaElement) Class(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ClassAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["class"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ClassRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["class"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Cols is the "cols" attribute.
// Maximum number of characters per line
// Valid values are constrained to the following:
//   * valid_non_negative_integer
func (e *TextareaElement) Cols(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["cols"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ColsAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["cols"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["cols"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ColsRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["cols"]
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
func (e *TextareaElement) Contenteditable(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ContenteditableAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["contenteditable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ContenteditableRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Dir(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["dir"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) DirAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dir"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["dir"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) DirRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Dirname(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["dirname"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) DirnameAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dirname"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["dirname"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) DirnameRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Disabled() *TextareaElement {
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *TextareaElement) DisabledRemove() *TextareaElement {
    delete(e.BoolAttributes,"disabled")
    return e
}

func (e *TextareaElement) DisabledSet(v bool) *TextareaElement {
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
func (e *TextareaElement) Draggable(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["draggable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) DraggableAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["draggable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["draggable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) DraggableRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Enterkeyhint(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) EnterkeyhintAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["enterkeyhint"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) EnterkeyhintRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Form(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["form"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) FormAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["form"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["form"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) FormRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["form"]
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
func (e *TextareaElement) Hidden(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["hidden"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) HiddenAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["hidden"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["hidden"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) HiddenRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Id(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["id"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) IdAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["id"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["id"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) IdRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Inert() *TextareaElement {
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *TextareaElement) InertRemove() *TextareaElement {
    delete(e.BoolAttributes,"inert")
    return e
}

func (e *TextareaElement) InertSet(v bool) *TextareaElement {
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
func (e *TextareaElement) Inputmode(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["inputmode"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) InputmodeAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["inputmode"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["inputmode"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) InputmodeRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Is(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["is"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) IsAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["is"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["is"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) IsRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Itemid(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemid"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ItemidAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemid"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ItemidRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Itemprop(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemprop"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ItempropAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemprop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemprop"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ItempropRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Itemref(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemref"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ItemrefAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemref"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemref"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ItemrefRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Itemscope() *TextareaElement {
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *TextareaElement) ItemscopeRemove() *TextareaElement {
    delete(e.BoolAttributes,"itemscope")
    return e
}

func (e *TextareaElement) ItemscopeSet(v bool) *TextareaElement {
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
func (e *TextareaElement) Itemtype(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemtype"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ItemtypeAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemtype"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) ItemtypeRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (e *TextareaElement) Lang(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["lang"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) LangAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["lang"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["lang"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) LangRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["lang"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Maxlength is the "maxlength" attribute.
// Maximum length of value
// Valid values are constrained to the following:
//   * valid_non_negative_integer
func (e *TextareaElement) Maxlength(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["maxlength"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) MaxlengthAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["maxlength"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["maxlength"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) MaxlengthRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["maxlength"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Minlength is the "minlength" attribute.
// Minimum length of value
// Valid values are constrained to the following:
//   * valid_non_negative_integer
func (e *TextareaElement) Minlength(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["minlength"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) MinlengthAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["minlength"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["minlength"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) MinlengthRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["minlength"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Name is the "name" attribute.
// Name of shadow tree slot
// Valid values are constrained to the following:
//   * text
func (e *TextareaElement) Name(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["name"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) NameAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["name"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["name"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) NameRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Nonce(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["nonce"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) NonceAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["nonce"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["nonce"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) NonceRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["nonce"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Placeholder is the "placeholder" attribute.
// User-visible label to be placed within the form control
// Valid values are constrained to the following:
//   * text
func (e *TextareaElement) Placeholder(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["placeholder"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) PlaceholderAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["placeholder"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["placeholder"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) PlaceholderRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Popover(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["popover"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) PopoverAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["popover"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) PopoverRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popover"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Readonly is the "readonly" attribute.
// Affects willValidate, plus any behavior added by the custom element author
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *TextareaElement) Readonly() *TextareaElement {
    e.BoolAttributes["readonly"] = struct{}{}
    return e
}

func (e *TextareaElement) ReadonlyRemove() *TextareaElement {
    delete(e.BoolAttributes,"readonly")
    return e
}

func (e *TextareaElement) ReadonlySet(v bool) *TextareaElement {
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
func (e *TextareaElement) Required() *TextareaElement {
    e.BoolAttributes["required"] = struct{}{}
    return e
}

func (e *TextareaElement) RequiredRemove() *TextareaElement {
    delete(e.BoolAttributes,"required")
    return e
}

func (e *TextareaElement) RequiredSet(v bool) *TextareaElement {
    if v {
        e.Required()
    } else {
        e.RequiredRemove()
    }
    return e
}


// Rows is the "rows" attribute.
// Number of lines to show
// Valid values are constrained to the following:
//   * valid_non_negative_integer
func (e *TextareaElement) Rows(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["rows"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) RowsAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["rows"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["rows"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) RowsRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["rows"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   * text
func (e *TextareaElement) Slot(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["slot"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) SlotAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["slot"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["slot"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) SlotRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Spellcheck(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) SpellcheckAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) SpellcheckRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (e *TextareaElement) Style(k,v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedKVAttributes["style"]
    if !ok {
        htmlStringer = NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = htmlStringer
    }
    htmlStringer.Add(k,v)
    return e
}

func (e *TextareaElement) StyleRemove(k string) *TextareaElement {
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
func (e *TextareaElement) Tabindex(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["tabindex"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) TabindexAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["tabindex"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["tabindex"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) TabindexRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Title(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["title"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) TitleAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["title"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["title"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) TitleRemove(v string) *TextareaElement {
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
func (e *TextareaElement) Translate(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["translate"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) TranslateAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["translate"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) TranslateRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Wrap is the "wrap" attribute.
// How the value of the form control is to be wrapped for form submission
// Valid values are constrained to the following:
//   * soft
//   * soft
//   * hard
//   * hard
func (e *TextareaElement) Wrap(v string) *TextareaElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["wrap"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) WrapAdd(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["wrap"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["wrap"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *TextareaElement) WrapRemove(v string) *TextareaElement {
    htmlStringer,ok := e.DelimitedStringAttributes["wrap"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}




func (e *TextareaElement) OnAuxclick(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onauxclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnAuxclickRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnBeforematch(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforematch"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnBeforematchRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnBeforetoggle(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforetoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnBeforetoggleRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnBlur(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onblur"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnBlurRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnCancel(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncancel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnCancelRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnCanplay(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnCanplayRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnCanplaythrough(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplaythrough"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnCanplaythroughRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnChange(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnChangeRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnClick(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnClickRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnClose(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclose"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnCloseRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnContextlost(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextlost"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnContextlostRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnContextmenu(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextmenu"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnContextmenuRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnContextrestored(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextrestored"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnContextrestoredRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnCopy(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncopy"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnCopyRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnCuechange(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncuechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnCuechangeRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnCut(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncut"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnCutRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnDblclick(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondblclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnDblclickRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnDrag(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrag"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnDragRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnDragend(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnDragendRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnDragenter(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnDragenterRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnDragleave(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnDragleaveRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnDragover(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnDragoverRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnDragstart(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnDragstartRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnDrop(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrop"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnDropRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnDurationchange(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondurationchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnDurationchangeRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnEmptied(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onemptied"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnEmptiedRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnEnded(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onended"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnEndedRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnError(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onerror"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnErrorRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnFocus(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onfocus"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnFocusRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnFormdata(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onformdata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnFormdataRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnInput(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninput"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnInputRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnInvalid(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninvalid"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnInvalidRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnKeydown(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeydown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnKeydownRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnKeypress(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeypress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnKeypressRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnKeyup(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeyup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnKeyupRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnLoad(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onload"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnLoadRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnLoadeddata(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadeddata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnLoadeddataRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnLoadedmetadata(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadedmetadata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnLoadedmetadataRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnLoadstart(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnLoadstartRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnMousedown(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousedown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnMousedownRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnMouseenter(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnMouseenterRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnMouseleave(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnMouseleaveRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnMousemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousemove"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnMousemoveRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnMouseout(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseout"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnMouseoutRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnMouseover(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnMouseoverRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnMouseup(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnMouseupRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnPaste(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpaste"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnPasteRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnPause(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpause"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnPauseRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnPlay(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnPlayRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnPlaying(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplaying"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnPlayingRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnProgress(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onprogress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnProgressRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnRatechange(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onratechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnRatechangeRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnReset(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onreset"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnResetRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnResize(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onresize"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnResizeRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnScroll(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscroll"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnScrollRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnScrollend(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscrollend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnScrollendRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnSecuritypolicyviolation(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsecuritypolicyviolation"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnSecuritypolicyviolationRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnSeeked(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeked"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnSeekedRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnSeeking(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeking"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnSeekingRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnSelect(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onselect"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnSelectRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnSlotchange(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onslotchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnSlotchangeRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnStalled(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onstalled"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnStalledRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnSubmit(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsubmit"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnSubmitRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnSuspend(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsuspend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnSuspendRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnTimeupdate(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontimeupdate"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnTimeupdateRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnToggle(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnToggleRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnVolumechange(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onvolumechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnVolumechangeRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnWaiting(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwaiting"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnWaitingRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *TextareaElement) OnWheel(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwheel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *TextareaElement) OnWheelRemove(evt string) (*TextareaElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}


