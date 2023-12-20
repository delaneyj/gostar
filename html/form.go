package  html

import (
    "fmt"
)

type FormHTMLElement struct {
    *Element
}

func FORM(children ...fmt.Stringer) *FormHTMLElement {
    return &FormHTMLElement{
        Element: &Element{
            Tag: "form",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *FormHTMLElement) Children(children ...fmt.Stringer) *FormHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *FormHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *FormHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *FormHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *FormHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *FormHTMLElement) Text(text string) *FormHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *FormHTMLElement) TextF(format string, args ...any) *FormHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *FormHTMLElement) Raw(text string) *FormHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *FormHTMLElement) RawF(format string, args ...any) *FormHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *FormHTMLElement) CustomData(key, value string) *FormHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *FormHTMLElement) CustomDataRemove(key string) *FormHTMLElement {
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
func (e *FormHTMLElement) ACCEPT_CHARSET(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accept-charset"] = v
    return e
}

func (e *FormHTMLElement) RemoveACCEPT_CHARSET(v string) *FormHTMLElement {
    delete(e.StringAttributes, "accept-charset")
    return e
}
// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//  * ordered_set_of_unique_space_separated_tokens
//  * identical_to
func (e *FormHTMLElement) ACCESSKEY(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *FormHTMLElement) RemoveACCESSKEY(v string) *FormHTMLElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ACTION sets the "action" attribute.
// URL to use for form submission
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *FormHTMLElement) ACTION(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["action"] = v
    return e
}

func (e *FormHTMLElement) RemoveACTION(v string) *FormHTMLElement {
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
func (e *FormHTMLElement) AUTOCAPITALIZE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *FormHTMLElement) RemoveAUTOCAPITALIZE(v string) *FormHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//  * autofill_field
func (e *FormHTMLElement) AUTOCOMPLETE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocomplete"] = v
    return e
}

func (e *FormHTMLElement) RemoveAUTOCOMPLETE(v string) *FormHTMLElement {
    delete(e.StringAttributes, "autocomplete")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *FormHTMLElement) AUTOFOCUS() *FormHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *FormHTMLElement) RemoveAUTOFOCUS() *FormHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *FormHTMLElement) SetAUTOFOCUS(b bool) *FormHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *FormHTMLElement) CLASS(v string) *FormHTMLElement {
    if e.DelimitedStringAttributes == nil {
        e.DelimitedStringAttributes = map[string]*DelimitedString{}
    }
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        kv = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = kv
    }
    kv.Add(v)
    return e
}

func (e *FormHTMLElement) SetCLASS(v string) *FormHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *FormHTMLElement) RemoveCLASS(v string) *FormHTMLElement {
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
func (e *FormHTMLElement) CONTENTEDITABLE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *FormHTMLElement) RemoveCONTENTEDITABLE(v string) *FormHTMLElement {
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
func (e *FormHTMLElement) DIR(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *FormHTMLElement) RemoveDIR(v string) *FormHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *FormHTMLElement) DRAGGABLE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *FormHTMLElement) RemoveDRAGGABLE(v string) *FormHTMLElement {
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
func (e *FormHTMLElement) ENCTYPE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enctype"] = v
    return e
}

func (e *FormHTMLElement) RemoveENCTYPE(v string) *FormHTMLElement {
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
func (e *FormHTMLElement) ENTERKEYHINT(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *FormHTMLElement) RemoveENTERKEYHINT(v string) *FormHTMLElement {
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
func (e *FormHTMLElement) HIDDEN(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *FormHTMLElement) RemoveHIDDEN(v string) *FormHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *FormHTMLElement) ID(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *FormHTMLElement) RemoveID(v string) *FormHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *FormHTMLElement) INERT() *FormHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *FormHTMLElement) RemoveINERT() *FormHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *FormHTMLElement) SetINERT(b bool) *FormHTMLElement {
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
func (e *FormHTMLElement) INPUTMODE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *FormHTMLElement) RemoveINPUTMODE(v string) *FormHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *FormHTMLElement) IS(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *FormHTMLElement) RemoveIS(v string) *FormHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *FormHTMLElement) ITEMID(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *FormHTMLElement) RemoveITEMID(v string) *FormHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *FormHTMLElement) ITEMPROP(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *FormHTMLElement) RemoveITEMPROP(v string) *FormHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *FormHTMLElement) ITEMREF(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *FormHTMLElement) RemoveITEMREF(v string) *FormHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *FormHTMLElement) ITEMSCOPE() *FormHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *FormHTMLElement) RemoveITEMSCOPE() *FormHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *FormHTMLElement) SetITEMSCOPE(b bool) *FormHTMLElement {
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
func (e *FormHTMLElement) ITEMTYPE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *FormHTMLElement) RemoveITEMTYPE(v string) *FormHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *FormHTMLElement) LANG(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *FormHTMLElement) RemoveLANG(v string) *FormHTMLElement {
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
func (e *FormHTMLElement) METHOD(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["method"] = v
    return e
}

func (e *FormHTMLElement) RemoveMETHOD(v string) *FormHTMLElement {
    delete(e.StringAttributes, "method")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *FormHTMLElement) NAME(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *FormHTMLElement) RemoveNAME(v string) *FormHTMLElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *FormHTMLElement) NONCE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *FormHTMLElement) RemoveNONCE(v string) *FormHTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// NOVALIDATE sets the "novalidate" attribute.
// Bypass form control validation for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *FormHTMLElement) NOVALIDATE() *FormHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["novalidate"] = struct{}{}
    return e
}

func (e *FormHTMLElement) RemoveNOVALIDATE() *FormHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "novalidate")
    return e
}

func (e *FormHTMLElement) SetNOVALIDATE(b bool) *FormHTMLElement {
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
func (e *FormHTMLElement) POPOVER(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *FormHTMLElement) RemovePOPOVER(v string) *FormHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *FormHTMLElement) SLOT(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *FormHTMLElement) RemoveSLOT(v string) *FormHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *FormHTMLElement) SPELLCHECK(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *FormHTMLElement) RemoveSPELLCHECK(v string) *FormHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *FormHTMLElement) STYLE(k,v string) *FormHTMLElement {
    if e.DelimitedKVAttributes == nil {
        e.DelimitedKVAttributes = map[string]*DelimitedKVString{}
    }
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        kv = NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = kv
    }
    kv.Add(k,v)
    return e
}

func (e *FormHTMLElement) RemoveSTYLE(k string) *FormHTMLElement {
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
func (e *FormHTMLElement) TABINDEX(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *FormHTMLElement) RemoveTABINDEX(v string) *FormHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TARGET sets the "target" attribute.
// Navigable for form submission
// Values values are constrained to:
//  * valid_navigable_target_name_or_keyword
func (e *FormHTMLElement) TARGET(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["target"] = v
    return e
}

func (e *FormHTMLElement) RemoveTARGET(v string) *FormHTMLElement {
    delete(e.StringAttributes, "target")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *FormHTMLElement) TITLE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *FormHTMLElement) RemoveTITLE(v string) *FormHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *FormHTMLElement) TRANSLATE(v string) *FormHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *FormHTMLElement) RemoveTRANSLATE(v string) *FormHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
