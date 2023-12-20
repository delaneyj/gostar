package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type FormElement struct {
    *gostar.Element
}

func FORM(children ...fmt.Stringer) *FormElement {
    return &FormElement{
        Element: &gostar.Element{
            Tag: "form",
            IsSelfClosing: false,
            Children: children,
        },
    }
}

func (e *FormElement) AddChildren(children ...fmt.Stringer) *FormElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *FormElement) TEXT(text string) *FormElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *FormElement) RAW(text string) *FormElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *FormElement) CustomData(key, value string) *FormElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *FormElement) CustomDataRemove(key string) *FormElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}


// ACCEPT_CHARSET sets the "accept-charset" attribute.
// Character encodings to use for form submission
// Values values are constrained to:
//  * ascii_case_insensitive
//  * utf_8
func (e *FormElement) ACCEPT_CHARSET(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accept-charset"] = v
    return e
}

func (e *FormElement) RemoveACCEPT_CHARSET(v string) *FormElement {
    delete(e.StringAttributes, "accept-charset")
    return e
}
// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//  * ordered_set_of_unique_space_separated_tokens
//  * identical_to
func (e *FormElement) ACCESSKEY(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *FormElement) RemoveACCESSKEY(v string) *FormElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ACTION sets the "action" attribute.
// URL to use for form submission
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *FormElement) ACTION(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["action"] = v
    return e
}

func (e *FormElement) RemoveACTION(v string) *FormElement {
    delete(e.StringAttributes, "action")
    return e
}
// AUTOCAPITALIZE sets the "autocapitalize" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Values values are constrained to:
//  * on
//  * on
//  * off
//  * off
//  * none
//  * none
//  * sentences
//  * sentences
//  * words
//  * words
//  * characters
//  * characters
func (e *FormElement) AUTOCAPITALIZE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *FormElement) RemoveAUTOCAPITALIZE(v string) *FormElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//  * autofill_field
func (e *FormElement) AUTOCOMPLETE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocomplete"] = v
    return e
}

func (e *FormElement) RemoveAUTOCOMPLETE(v string) *FormElement {
    delete(e.StringAttributes, "autocomplete")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *FormElement) AUTOFOCUS() *FormElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *FormElement) RemoveAUTOFOCUS() *FormElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *FormElement) SetAUTOFOCUS(b bool) *FormElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *FormElement) CLASS(v string) *FormElement {
    if e.DelimitedStringAttributes == nil {
        e.DelimitedStringAttributes = map[string]*gostar.DelimitedString{}
    }
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        kv = gostar.NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = kv
    }
    kv.Add(v)
    return e
}

func (e *FormElement) SetCLASS(v string) *FormElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *FormElement) RemoveCLASS(v string) *FormElement {
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        return e
    }
    kv.Remove(v)
    return e
}
// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//  * true
//  * plaintext_only
//  * false
func (e *FormElement) CONTENTEDITABLE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *FormElement) RemoveCONTENTEDITABLE(v string) *FormElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//  * ltr
//  * ltr
//  * rtl
//  * rtl
func (e *FormElement) DIR(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *FormElement) RemoveDIR(v string) *FormElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *FormElement) DRAGGABLE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *FormElement) RemoveDRAGGABLE(v string) *FormElement {
    delete(e.StringAttributes, "draggable")
    return e
}
// ENCTYPE sets the "enctype" attribute.
// Entry list encoding type to use for form submission
// Values values are constrained to:
//  * application/x_www_form_urlencoded
//  * application/x_www_form_urlencoded
//  * multipart/form_data
//  * multipart/form_data
//  * text/plain
//  * text/plain
func (e *FormElement) ENCTYPE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enctype"] = v
    return e
}

func (e *FormElement) RemoveENCTYPE(v string) *FormElement {
    delete(e.StringAttributes, "enctype")
    return e
}
// ENTERKEYHINT sets the "enterkeyhint" attribute.
// Hint for selecting an enter key action
// Values values are constrained to:
//  * enter
//  * enter
//  * done
//  * done
//  * go
//  * go
//  * next
//  * next
//  * previous
//  * previous
//  * search
//  * search
//  * send
//  * send
func (e *FormElement) ENTERKEYHINT(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *FormElement) RemoveENTERKEYHINT(v string) *FormElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *FormElement) HIDDEN(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *FormElement) RemoveHIDDEN(v string) *FormElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *FormElement) ID(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *FormElement) RemoveID(v string) *FormElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *FormElement) INERT() *FormElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *FormElement) RemoveINERT() *FormElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *FormElement) SetINERT(b bool) *FormElement {
    if b {
        return e.INERT()
    }
    return e.RemoveINERT()
}
// INPUTMODE sets the "inputmode" attribute.
// Hint for selecting an input modality
// Values values are constrained to:
//  * none
//  * none
//  * text
//  * text
//  * tel
//  * tel
//  * email
//  * email
//  * url
//  * url
//  * numeric
//  * numeric
//  * decimal
//  * decimal
//  * search
//  * search
func (e *FormElement) INPUTMODE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *FormElement) RemoveINPUTMODE(v string) *FormElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *FormElement) IS(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *FormElement) RemoveIS(v string) *FormElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *FormElement) ITEMID(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *FormElement) RemoveITEMID(v string) *FormElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *FormElement) ITEMPROP(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *FormElement) RemoveITEMPROP(v string) *FormElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *FormElement) ITEMREF(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *FormElement) RemoveITEMREF(v string) *FormElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *FormElement) ITEMSCOPE() *FormElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *FormElement) RemoveITEMSCOPE() *FormElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *FormElement) SetITEMSCOPE(b bool) *FormElement {
    if b {
        return e.ITEMSCOPE()
    }
    return e.RemoveITEMSCOPE()
}
// ITEMTYPE sets the "itemtype" attribute.
// Item types of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
func (e *FormElement) ITEMTYPE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *FormElement) RemoveITEMTYPE(v string) *FormElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *FormElement) LANG(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *FormElement) RemoveLANG(v string) *FormElement {
    delete(e.StringAttributes, "lang")
    return e
}
// METHOD sets the "method" attribute.
// Variant to use for form submission
// Values values are constrained to:
//  * get
//  * get
//  * post
//  * post
//  * dialog
//  * dialog
func (e *FormElement) METHOD(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["method"] = v
    return e
}

func (e *FormElement) RemoveMETHOD(v string) *FormElement {
    delete(e.StringAttributes, "method")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *FormElement) NAME(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *FormElement) RemoveNAME(v string) *FormElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *FormElement) NONCE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *FormElement) RemoveNONCE(v string) *FormElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// NOVALIDATE sets the "novalidate" attribute.
// Bypass form control validation for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *FormElement) NOVALIDATE() *FormElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["novalidate"] = struct{}{}
    return e
}

func (e *FormElement) RemoveNOVALIDATE() *FormElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "novalidate")
    return e
}

func (e *FormElement) SetNOVALIDATE(b bool) *FormElement {
    if b {
        return e.NOVALIDATE()
    }
    return e.RemoveNOVALIDATE()
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *FormElement) POPOVER(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *FormElement) RemovePOPOVER(v string) *FormElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *FormElement) SLOT(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *FormElement) RemoveSLOT(v string) *FormElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *FormElement) SPELLCHECK(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *FormElement) RemoveSPELLCHECK(v string) *FormElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *FormElement) STYLE(k,v string) *FormElement {
    if e.DelimitedKVAttributes == nil {
        e.DelimitedKVAttributes = map[string]*gostar.DelimitedKVString{}
    }
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        kv = gostar.NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = kv
    }
    kv.Add(k,v)
    return e
}

func (e *FormElement) RemoveSTYLE(k string) *FormElement {
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        return e
    }
    kv.Remove(k)
    return e
}
// TABINDEX sets the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and
//      the relative order of the element for the purposes of sequential focus navigation
// Values values are constrained to:
//  * valid_integer
func (e *FormElement) TABINDEX(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *FormElement) RemoveTABINDEX(v string) *FormElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TARGET sets the "target" attribute.
// Navigable for form submission
// Values values are constrained to:
//  * valid_navigable_target_name_or_keyword
func (e *FormElement) TARGET(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["target"] = v
    return e
}

func (e *FormElement) RemoveTARGET(v string) *FormElement {
    delete(e.StringAttributes, "target")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *FormElement) TITLE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *FormElement) RemoveTITLE(v string) *FormElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *FormElement) TRANSLATE(v string) *FormElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *FormElement) RemoveTRANSLATE(v string) *FormElement {
    delete(e.StringAttributes, "translate")
    return e
}
