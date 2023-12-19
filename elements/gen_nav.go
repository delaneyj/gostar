
/* cSpell:disable */

package  elements

type NavElement struct {
    *Element
}

func NAV() *NavElement {
    return &NavElement{
        Element: &Element{
            Tag: "nav",
            IsSelfClosing: false,
        },
    }
}

func (e *NavElement) AddChildren(children ...*Element) *NavElement {
    e.Children = append(e.Children, children...)
	return e
}


// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   * ordered_set_of_unique_space_separated_tokens
//   * identical_to
func (e *NavElement) Accesskey(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["accesskey"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) AccesskeyAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["accesskey"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["accesskey"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) AccesskeyRemove(v string) *NavElement {
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
func (e *NavElement) Autocapitalize(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) AutocapitalizeAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocapitalize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) AutocapitalizeRemove(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocapitalize"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *NavElement) Autofocus() *NavElement {
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *NavElement) AutofocusRemove() *NavElement {
    delete(e.BoolAttributes,"autofocus")
    return e
}

func (e *NavElement) AutofocusSet(v bool) *NavElement {
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
func (e *NavElement) Class(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ClassAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["class"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ClassRemove(v string) *NavElement {
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
func (e *NavElement) Contenteditable(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ContenteditableAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["contenteditable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ContenteditableRemove(v string) *NavElement {
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
func (e *NavElement) Dir(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["dir"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) DirAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dir"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["dir"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) DirRemove(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dir"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   * true
//   * false
func (e *NavElement) Draggable(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["draggable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) DraggableAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["draggable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["draggable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) DraggableRemove(v string) *NavElement {
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
func (e *NavElement) Enterkeyhint(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) EnterkeyhintAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["enterkeyhint"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) EnterkeyhintRemove(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["enterkeyhint"]
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
func (e *NavElement) Hidden(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["hidden"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) HiddenAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["hidden"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["hidden"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) HiddenRemove(v string) *NavElement {
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
func (e *NavElement) Id(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["id"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) IdAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["id"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["id"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) IdRemove(v string) *NavElement {
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
func (e *NavElement) Inert() *NavElement {
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *NavElement) InertRemove() *NavElement {
    delete(e.BoolAttributes,"inert")
    return e
}

func (e *NavElement) InertSet(v bool) *NavElement {
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
func (e *NavElement) Inputmode(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["inputmode"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) InputmodeAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["inputmode"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["inputmode"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) InputmodeRemove(v string) *NavElement {
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
func (e *NavElement) Is(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["is"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) IsAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["is"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["is"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) IsRemove(v string) *NavElement {
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
func (e *NavElement) Itemid(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemid"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ItemidAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemid"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ItemidRemove(v string) *NavElement {
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
func (e *NavElement) Itemprop(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemprop"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ItempropAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemprop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemprop"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ItempropRemove(v string) *NavElement {
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
func (e *NavElement) Itemref(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemref"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ItemrefAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemref"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemref"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ItemrefRemove(v string) *NavElement {
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
func (e *NavElement) Itemscope() *NavElement {
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *NavElement) ItemscopeRemove() *NavElement {
    delete(e.BoolAttributes,"itemscope")
    return e
}

func (e *NavElement) ItemscopeSet(v bool) *NavElement {
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
func (e *NavElement) Itemtype(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemtype"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ItemtypeAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemtype"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) ItemtypeRemove(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (e *NavElement) Lang(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["lang"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) LangAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["lang"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["lang"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) LangRemove(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["lang"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   * text
func (e *NavElement) Nonce(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["nonce"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) NonceAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["nonce"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["nonce"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) NonceRemove(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["nonce"]
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
func (e *NavElement) Popover(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["popover"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) PopoverAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["popover"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) PopoverRemove(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popover"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   * text
func (e *NavElement) Slot(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["slot"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) SlotAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["slot"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["slot"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) SlotRemove(v string) *NavElement {
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
func (e *NavElement) Spellcheck(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) SpellcheckAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) SpellcheckRemove(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (e *NavElement) Style(k,v string) *NavElement {
    htmlStringer,ok := e.DelimitedKVAttributes["style"]
    if !ok {
        htmlStringer = NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = htmlStringer
    }
    htmlStringer.Add(k,v)
    return e
}

func (e *NavElement) StyleRemove(k string) *NavElement {
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
func (e *NavElement) Tabindex(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["tabindex"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) TabindexAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["tabindex"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["tabindex"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) TabindexRemove(v string) *NavElement {
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
func (e *NavElement) Title(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["title"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) TitleAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["title"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["title"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) TitleRemove(v string) *NavElement {
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
func (e *NavElement) Translate(v string) *NavElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["translate"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) TranslateAdd(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["translate"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *NavElement) TranslateRemove(v string) *NavElement {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}




func (e *NavElement) OnAuxclick(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onauxclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnAuxclickRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnBeforematch(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforematch"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnBeforematchRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnBeforetoggle(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforetoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnBeforetoggleRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnBlur(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onblur"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnBlurRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnCancel(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncancel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnCancelRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnCanplay(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnCanplayRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnCanplaythrough(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplaythrough"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnCanplaythroughRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnChange(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnChangeRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnClick(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnClickRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnClose(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclose"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnCloseRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnContextlost(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextlost"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnContextlostRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnContextmenu(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextmenu"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnContextmenuRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnContextrestored(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextrestored"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnContextrestoredRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnCopy(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncopy"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnCopyRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnCuechange(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncuechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnCuechangeRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnCut(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncut"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnCutRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnDblclick(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondblclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnDblclickRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnDrag(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrag"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnDragRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnDragend(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnDragendRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnDragenter(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnDragenterRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnDragleave(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnDragleaveRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnDragover(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnDragoverRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnDragstart(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnDragstartRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnDrop(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrop"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnDropRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnDurationchange(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondurationchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnDurationchangeRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnEmptied(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onemptied"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnEmptiedRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnEnded(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onended"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnEndedRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnError(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onerror"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnErrorRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnFocus(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onfocus"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnFocusRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnFormdata(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onformdata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnFormdataRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnInput(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninput"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnInputRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnInvalid(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninvalid"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnInvalidRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnKeydown(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeydown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnKeydownRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnKeypress(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeypress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnKeypressRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnKeyup(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeyup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnKeyupRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnLoad(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onload"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnLoadRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnLoadeddata(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadeddata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnLoadeddataRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnLoadedmetadata(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadedmetadata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnLoadedmetadataRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnLoadstart(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnLoadstartRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnMousedown(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousedown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnMousedownRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnMouseenter(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnMouseenterRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnMouseleave(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnMouseleaveRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnMousemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousemove"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnMousemoveRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnMouseout(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseout"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnMouseoutRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnMouseover(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnMouseoverRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnMouseup(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnMouseupRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnPaste(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpaste"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnPasteRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnPause(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpause"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnPauseRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnPlay(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnPlayRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnPlaying(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplaying"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnPlayingRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnProgress(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onprogress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnProgressRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnRatechange(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onratechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnRatechangeRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnReset(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onreset"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnResetRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnResize(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onresize"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnResizeRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnScroll(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscroll"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnScrollRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnScrollend(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscrollend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnScrollendRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnSecuritypolicyviolation(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsecuritypolicyviolation"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnSecuritypolicyviolationRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnSeeked(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeked"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnSeekedRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnSeeking(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeking"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnSeekingRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnSelect(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onselect"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnSelectRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnSlotchange(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onslotchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnSlotchangeRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnStalled(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onstalled"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnStalledRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnSubmit(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsubmit"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnSubmitRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnSuspend(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsuspend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnSuspendRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnTimeupdate(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontimeupdate"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnTimeupdateRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnToggle(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnToggleRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnVolumechange(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onvolumechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnVolumechangeRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnWaiting(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwaiting"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnWaitingRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *NavElement) OnWheel(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwheel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *NavElement) OnWheelRemove(evt string) (*NavElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}


