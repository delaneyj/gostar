
/* cSpell:disable */

package  elements

type H1Element struct {
    *Element
}

func H1() *H1Element {
    return &H1Element{
        Element: &Element{
            Tag: "h1",
            IsSelfClosing: false,
        },
    }
}

func (e *H1Element) AddChildren(children ...*Element) *H1Element {
    e.Children = append(e.Children, children...)
	return e
}


// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   * ordered_set_of_unique_space_separated_tokens
//   * identical_to
func (e *H1Element) Accesskey(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["accesskey"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) AccesskeyAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["accesskey"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["accesskey"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) AccesskeyRemove(v string) *H1Element {
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
func (e *H1Element) Autocapitalize(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) AutocapitalizeAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["autocapitalize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) AutocapitalizeRemove(v string) *H1Element {
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
func (e *H1Element) Autofocus() *H1Element {
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *H1Element) AutofocusRemove() *H1Element {
    delete(e.BoolAttributes,"autofocus")
    return e
}

func (e *H1Element) AutofocusSet(v bool) *H1Element {
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
func (e *H1Element) Class(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ClassAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["class"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ClassRemove(v string) *H1Element {
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
func (e *H1Element) Contenteditable(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ContenteditableAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["contenteditable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ContenteditableRemove(v string) *H1Element {
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
func (e *H1Element) Dir(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["dir"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) DirAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["dir"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["dir"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) DirRemove(v string) *H1Element {
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
func (e *H1Element) Draggable(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["draggable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) DraggableAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["draggable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["draggable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) DraggableRemove(v string) *H1Element {
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
func (e *H1Element) Enterkeyhint(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) EnterkeyhintAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["enterkeyhint"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) EnterkeyhintRemove(v string) *H1Element {
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
func (e *H1Element) Hidden(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["hidden"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) HiddenAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["hidden"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["hidden"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) HiddenRemove(v string) *H1Element {
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
func (e *H1Element) Id(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["id"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) IdAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["id"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["id"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) IdRemove(v string) *H1Element {
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
func (e *H1Element) Inert() *H1Element {
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *H1Element) InertRemove() *H1Element {
    delete(e.BoolAttributes,"inert")
    return e
}

func (e *H1Element) InertSet(v bool) *H1Element {
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
func (e *H1Element) Inputmode(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["inputmode"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) InputmodeAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["inputmode"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["inputmode"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) InputmodeRemove(v string) *H1Element {
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
func (e *H1Element) Is(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["is"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) IsAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["is"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["is"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) IsRemove(v string) *H1Element {
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
func (e *H1Element) Itemid(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemid"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ItemidAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["itemid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemid"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ItemidRemove(v string) *H1Element {
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
func (e *H1Element) Itemprop(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemprop"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ItempropAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["itemprop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemprop"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ItempropRemove(v string) *H1Element {
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
func (e *H1Element) Itemref(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemref"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ItemrefAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["itemref"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemref"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ItemrefRemove(v string) *H1Element {
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
func (e *H1Element) Itemscope() *H1Element {
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *H1Element) ItemscopeRemove() *H1Element {
    delete(e.BoolAttributes,"itemscope")
    return e
}

func (e *H1Element) ItemscopeSet(v bool) *H1Element {
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
func (e *H1Element) Itemtype(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemtype"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ItemtypeAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemtype"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) ItemtypeRemove(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (e *H1Element) Lang(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["lang"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) LangAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["lang"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["lang"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) LangRemove(v string) *H1Element {
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
func (e *H1Element) Nonce(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["nonce"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) NonceAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["nonce"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["nonce"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) NonceRemove(v string) *H1Element {
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
func (e *H1Element) Popover(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["popover"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) PopoverAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["popover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["popover"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) PopoverRemove(v string) *H1Element {
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
func (e *H1Element) Slot(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["slot"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) SlotAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["slot"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["slot"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) SlotRemove(v string) *H1Element {
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
func (e *H1Element) Spellcheck(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) SpellcheckAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) SpellcheckRemove(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (e *H1Element) Style(k,v string) *H1Element {
    htmlStringer,ok := e.DelimitedKVAttributes["style"]
    if !ok {
        htmlStringer = NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = htmlStringer
    }
    htmlStringer.Add(k,v)
    return e
}

func (e *H1Element) StyleRemove(k string) *H1Element {
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
func (e *H1Element) Tabindex(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["tabindex"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) TabindexAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["tabindex"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["tabindex"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) TabindexRemove(v string) *H1Element {
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
func (e *H1Element) Title(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["title"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) TitleAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["title"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["title"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) TitleRemove(v string) *H1Element {
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
func (e *H1Element) Translate(v string) *H1Element {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["translate"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) TranslateAdd(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["translate"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *H1Element) TranslateRemove(v string) *H1Element {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}




func (e *H1Element) OnAuxclick(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onauxclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnAuxclickRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnBeforematch(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforematch"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnBeforematchRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnBeforetoggle(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforetoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnBeforetoggleRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnBlur(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onblur"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnBlurRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnCancel(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncancel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnCancelRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnCanplay(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnCanplayRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnCanplaythrough(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplaythrough"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnCanplaythroughRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnChange(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnChangeRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnClick(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnClickRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnClose(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclose"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnCloseRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnContextlost(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextlost"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnContextlostRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnContextmenu(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextmenu"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnContextmenuRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnContextrestored(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextrestored"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnContextrestoredRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnCopy(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncopy"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnCopyRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnCuechange(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncuechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnCuechangeRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnCut(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncut"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnCutRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnDblclick(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondblclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnDblclickRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnDrag(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrag"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnDragRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnDragend(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnDragendRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnDragenter(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnDragenterRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnDragleave(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnDragleaveRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnDragover(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnDragoverRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnDragstart(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnDragstartRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnDrop(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrop"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnDropRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnDurationchange(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondurationchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnDurationchangeRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnEmptied(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onemptied"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnEmptiedRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnEnded(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onended"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnEndedRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnError(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onerror"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnErrorRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnFocus(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onfocus"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnFocusRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnFormdata(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onformdata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnFormdataRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnInput(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninput"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnInputRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnInvalid(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninvalid"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnInvalidRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnKeydown(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeydown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnKeydownRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnKeypress(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeypress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnKeypressRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnKeyup(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeyup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnKeyupRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnLoad(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onload"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnLoadRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnLoadeddata(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadeddata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnLoadeddataRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnLoadedmetadata(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadedmetadata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnLoadedmetadataRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnLoadstart(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnLoadstartRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnMousedown(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousedown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnMousedownRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnMouseenter(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnMouseenterRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnMouseleave(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnMouseleaveRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnMousemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousemove"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnMousemoveRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnMouseout(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseout"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnMouseoutRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnMouseover(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnMouseoverRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnMouseup(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnMouseupRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnPaste(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpaste"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnPasteRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnPause(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpause"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnPauseRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnPlay(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnPlayRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnPlaying(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplaying"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnPlayingRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnProgress(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onprogress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnProgressRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnRatechange(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onratechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnRatechangeRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnReset(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onreset"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnResetRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnResize(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onresize"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnResizeRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnScroll(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscroll"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnScrollRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnScrollend(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscrollend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnScrollendRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnSecuritypolicyviolation(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsecuritypolicyviolation"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnSecuritypolicyviolationRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnSeeked(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeked"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnSeekedRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnSeeking(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeking"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnSeekingRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnSelect(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onselect"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnSelectRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnSlotchange(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onslotchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnSlotchangeRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnStalled(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onstalled"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnStalledRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnSubmit(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsubmit"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnSubmitRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnSuspend(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsuspend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnSuspendRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnTimeupdate(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontimeupdate"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnTimeupdateRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnToggle(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnToggleRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnVolumechange(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onvolumechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnVolumechangeRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnWaiting(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwaiting"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnWaitingRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *H1Element) OnWheel(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwheel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *H1Element) OnWheelRemove(evt string) (*H1Element) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}


