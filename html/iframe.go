package  html

import (
    "fmt"
)

type IframeHTMLElement struct {
    *Element
}

func IFRAME(children ...ElementBuilder) *IframeHTMLElement {
    return &IframeHTMLElement{
        Element: &Element{
            Tag: "iframe",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *IframeHTMLElement) Children(children ...ElementBuilder) *IframeHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *IframeHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *IframeHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *IframeHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *IframeHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *IframeHTMLElement) Text(text string) *IframeHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *IframeHTMLElement) TextF(format string, args ...any) *IframeHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *IframeHTMLElement) Raw(text string) *IframeHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *IframeHTMLElement) RawF(format string, args ...any) *IframeHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *IframeHTMLElement) CustomData(key, value string) *IframeHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *IframeHTMLElement) CustomDataRemove(key string) *IframeHTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}


// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//  * ordered_set_of_unique_space_separated_tokens
//  * identical_to
func (e *IframeHTMLElement) ACCESSKEY(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *IframeHTMLElement) RemoveACCESSKEY(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ALLOW sets the "allow" attribute.
// Permissions policy to be applied to the iframe's contents
// Values values are constrained to:
//  * serialized_permissions_policy
func (e *IframeHTMLElement) ALLOW(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["allow"] = v
    return e
}

func (e *IframeHTMLElement) RemoveALLOW(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "allow")
    return e
}
// ALLOWFULLSCREEN sets the "allowfullscreen" attribute.
// Whether to allow the iframe's contents to use requestFullscreen()
// Values values are constrained to:
//  * boolean_attribute
func (e *IframeHTMLElement) ALLOWFULLSCREEN() *IframeHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["allowfullscreen"] = struct{}{}
    return e
}

func (e *IframeHTMLElement) RemoveALLOWFULLSCREEN() *IframeHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "allowfullscreen")
    return e
}

func (e *IframeHTMLElement) SetALLOWFULLSCREEN(b bool) *IframeHTMLElement {
    if b {
        return e.ALLOWFULLSCREEN()
    }
    return e.RemoveALLOWFULLSCREEN()
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
func (e *IframeHTMLElement) AUTOCAPITALIZE(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *IframeHTMLElement) RemoveAUTOCAPITALIZE(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *IframeHTMLElement) AUTOFOCUS() *IframeHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *IframeHTMLElement) RemoveAUTOFOCUS() *IframeHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *IframeHTMLElement) SetAUTOFOCUS(b bool) *IframeHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *IframeHTMLElement) CLASS(v string) *IframeHTMLElement {
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

func (e *IframeHTMLElement) SetCLASS(v string) *IframeHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *IframeHTMLElement) RemoveCLASS(v string) *IframeHTMLElement {
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
func (e *IframeHTMLElement) CONTENTEDITABLE(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *IframeHTMLElement) RemoveCONTENTEDITABLE(v string) *IframeHTMLElement {
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
func (e *IframeHTMLElement) DIR(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *IframeHTMLElement) RemoveDIR(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *IframeHTMLElement) DRAGGABLE(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *IframeHTMLElement) RemoveDRAGGABLE(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "draggable")
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
func (e *IframeHTMLElement) ENTERKEYHINT(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *IframeHTMLElement) RemoveENTERKEYHINT(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *IframeHTMLElement) HEIGHT(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *IframeHTMLElement) RemoveHEIGHT(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "height")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *IframeHTMLElement) HIDDEN(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *IframeHTMLElement) RemoveHIDDEN(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *IframeHTMLElement) ID(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *IframeHTMLElement) RemoveID(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *IframeHTMLElement) INERT() *IframeHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *IframeHTMLElement) RemoveINERT() *IframeHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *IframeHTMLElement) SetINERT(b bool) *IframeHTMLElement {
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
func (e *IframeHTMLElement) INPUTMODE(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *IframeHTMLElement) RemoveINPUTMODE(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *IframeHTMLElement) IS(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *IframeHTMLElement) RemoveIS(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *IframeHTMLElement) ITEMID(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *IframeHTMLElement) RemoveITEMID(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *IframeHTMLElement) ITEMPROP(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *IframeHTMLElement) RemoveITEMPROP(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *IframeHTMLElement) ITEMREF(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *IframeHTMLElement) RemoveITEMREF(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *IframeHTMLElement) ITEMSCOPE() *IframeHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *IframeHTMLElement) RemoveITEMSCOPE() *IframeHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *IframeHTMLElement) SetITEMSCOPE(b bool) *IframeHTMLElement {
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
func (e *IframeHTMLElement) ITEMTYPE(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *IframeHTMLElement) RemoveITEMTYPE(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *IframeHTMLElement) LANG(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *IframeHTMLElement) RemoveLANG(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LOADING sets the "loading" attribute.
// Used when determining loading deferral
// Values values are constrained to:
//  * lazy
//  * lazy
//  * eager
//  * eager
func (e *IframeHTMLElement) LOADING(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["loading"] = v
    return e
}

func (e *IframeHTMLElement) RemoveLOADING(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "loading")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *IframeHTMLElement) NAME(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *IframeHTMLElement) RemoveNAME(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *IframeHTMLElement) NONCE(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *IframeHTMLElement) RemoveNONCE(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *IframeHTMLElement) POPOVER(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *IframeHTMLElement) RemovePOPOVER(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *IframeHTMLElement) REFERRERPOLICY(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *IframeHTMLElement) RemoveREFERRERPOLICY(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// SANDBOX sets the "sandbox" attribute.
// Security rules for nested content
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * ascii_case_insensitive
//  * allow_downloads
//  * allow_downloads
//  * allow_forms
//  * allow_forms
//  * allow_modals
//  * allow_modals
//  * allow_orientation_lock
//  * allow_orientation_lock
//  * allow_pointer_lock
//  * allow_pointer_lock
//  * allow_popups
//  * allow_popups
//  * allow_popups_to_escape_sandbox
//  * allow_popups_to_escape_sandbox
//  * allow_presentation
//  * allow_presentation
//  * allow_same_origin
//  * allow_same_origin
//  * allow_scripts
//  * allow_scripts
//  * allow_top_navigation
//  * allow_top_navigation
//  * allow_top_navigation_by_user_activation
//  * allow_top_navigation_by_user_activation
//  * allow_top_navigation_to_custom_protocols
//  * allow_top_navigation_to_custom_protocols
func (e *IframeHTMLElement) SANDBOX(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["sandbox"] = v
    return e
}

func (e *IframeHTMLElement) RemoveSANDBOX(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "sandbox")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *IframeHTMLElement) SLOT(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *IframeHTMLElement) RemoveSLOT(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *IframeHTMLElement) SPELLCHECK(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *IframeHTMLElement) RemoveSPELLCHECK(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *IframeHTMLElement) SRC(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *IframeHTMLElement) RemoveSRC(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "src")
    return e
}
// SRCDOC sets the "srcdoc" attribute.
// A document to render in the iframe
// Values values are constrained to:
//  * an_iframe_srcdoc_document
//  * iframe
//  * srcdoc
func (e *IframeHTMLElement) SRCDOC(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["srcdoc"] = v
    return e
}

func (e *IframeHTMLElement) RemoveSRCDOC(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "srcdoc")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *IframeHTMLElement) STYLE(k,v string) *IframeHTMLElement {
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

func (e *IframeHTMLElement) RemoveSTYLE(k string) *IframeHTMLElement {
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
func (e *IframeHTMLElement) TABINDEX(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *IframeHTMLElement) RemoveTABINDEX(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *IframeHTMLElement) TITLE(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *IframeHTMLElement) RemoveTITLE(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *IframeHTMLElement) TRANSLATE(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *IframeHTMLElement) RemoveTRANSLATE(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *IframeHTMLElement) WIDTH(v string) *IframeHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *IframeHTMLElement) RemoveWIDTH(v string) *IframeHTMLElement {
    delete(e.StringAttributes, "width")
    return e
}
