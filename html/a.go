package  html

import (
    "fmt"
)

type AHTMLElement struct {
    *Element
}

func A(children ...fmt.Stringer) *AHTMLElement {
    return &AHTMLElement{
        Element: &Element{
            Tag: "a",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *AHTMLElement) Children(children ...fmt.Stringer) *AHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *AHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *AHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *AHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *AHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *AHTMLElement) Text(text string) *AHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *AHTMLElement) TextF(format string, args ...any) *AHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *AHTMLElement) Raw(text string) *AHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *AHTMLElement) RawF(format string, args ...any) *AHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *AHTMLElement) CustomData(key, value string) *AHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AHTMLElement) CustomDataRemove(key string) *AHTMLElement {
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
func (e *AHTMLElement) ACCESSKEY(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *AHTMLElement) RemoveACCESSKEY(v string) *AHTMLElement {
    delete(e.StringAttributes, "accesskey")
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
func (e *AHTMLElement) AUTOCAPITALIZE(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *AHTMLElement) RemoveAUTOCAPITALIZE(v string) *AHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *AHTMLElement) AUTOFOCUS() *AHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *AHTMLElement) RemoveAUTOFOCUS() *AHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *AHTMLElement) SetAUTOFOCUS(b bool) *AHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *AHTMLElement) CLASS(v string) *AHTMLElement {
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

func (e *AHTMLElement) SetCLASS(v string) *AHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *AHTMLElement) RemoveCLASS(v string) *AHTMLElement {
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
func (e *AHTMLElement) CONTENTEDITABLE(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *AHTMLElement) RemoveCONTENTEDITABLE(v string) *AHTMLElement {
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
func (e *AHTMLElement) DIR(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *AHTMLElement) RemoveDIR(v string) *AHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DOWNLOAD sets the "download" attribute.
// Whether to download the resource instead of navigating to it, and its filename if so
// Values values are constrained to:
func (e *AHTMLElement) DOWNLOAD(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["download"] = v
    return e
}

func (e *AHTMLElement) RemoveDOWNLOAD(v string) *AHTMLElement {
    delete(e.StringAttributes, "download")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *AHTMLElement) DRAGGABLE(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *AHTMLElement) RemoveDRAGGABLE(v string) *AHTMLElement {
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
func (e *AHTMLElement) ENTERKEYHINT(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *AHTMLElement) RemoveENTERKEYHINT(v string) *AHTMLElement {
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
func (e *AHTMLElement) HIDDEN(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *AHTMLElement) RemoveHIDDEN(v string) *AHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// HREF sets the "href" attribute.
// Document base URL
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *AHTMLElement) HREF(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["href"] = v
    return e
}

func (e *AHTMLElement) RemoveHREF(v string) *AHTMLElement {
    delete(e.StringAttributes, "href")
    return e
}
// HREFLANG sets the "hreflang" attribute.
// Language of the linked resource
// Values values are constrained to:
func (e *AHTMLElement) HREFLANG(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hreflang"] = v
    return e
}

func (e *AHTMLElement) RemoveHREFLANG(v string) *AHTMLElement {
    delete(e.StringAttributes, "hreflang")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *AHTMLElement) ID(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *AHTMLElement) RemoveID(v string) *AHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *AHTMLElement) INERT() *AHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *AHTMLElement) RemoveINERT() *AHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *AHTMLElement) SetINERT(b bool) *AHTMLElement {
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
func (e *AHTMLElement) INPUTMODE(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *AHTMLElement) RemoveINPUTMODE(v string) *AHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *AHTMLElement) IS(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *AHTMLElement) RemoveIS(v string) *AHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *AHTMLElement) ITEMID(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *AHTMLElement) RemoveITEMID(v string) *AHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *AHTMLElement) ITEMPROP(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *AHTMLElement) RemoveITEMPROP(v string) *AHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *AHTMLElement) ITEMREF(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *AHTMLElement) RemoveITEMREF(v string) *AHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *AHTMLElement) ITEMSCOPE() *AHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *AHTMLElement) RemoveITEMSCOPE() *AHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *AHTMLElement) SetITEMSCOPE(b bool) *AHTMLElement {
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
func (e *AHTMLElement) ITEMTYPE(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *AHTMLElement) RemoveITEMTYPE(v string) *AHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AHTMLElement) LANG(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *AHTMLElement) RemoveLANG(v string) *AHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *AHTMLElement) NONCE(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *AHTMLElement) RemoveNONCE(v string) *AHTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// PING sets the "ping" attribute.
// URLs to ping
// Values values are constrained to:
//  * set_of_space_separated_tokens
//  * valid_non_empty_ur_ls
func (e *AHTMLElement) PING(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["ping"] = v
    return e
}

func (e *AHTMLElement) RemovePING(v string) *AHTMLElement {
    delete(e.StringAttributes, "ping")
    return e
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *AHTMLElement) POPOVER(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *AHTMLElement) RemovePOPOVER(v string) *AHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *AHTMLElement) REFERRERPOLICY(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *AHTMLElement) RemoveREFERRERPOLICY(v string) *AHTMLElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// REL sets the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *AHTMLElement) REL(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["rel"] = v
    return e
}

func (e *AHTMLElement) RemoveREL(v string) *AHTMLElement {
    delete(e.StringAttributes, "rel")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *AHTMLElement) SLOT(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *AHTMLElement) RemoveSLOT(v string) *AHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *AHTMLElement) SPELLCHECK(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *AHTMLElement) RemoveSPELLCHECK(v string) *AHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AHTMLElement) STYLE(k,v string) *AHTMLElement {
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

func (e *AHTMLElement) RemoveSTYLE(k string) *AHTMLElement {
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
func (e *AHTMLElement) TABINDEX(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *AHTMLElement) RemoveTABINDEX(v string) *AHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TARGET sets the "target" attribute.
// Navigable for form submission
// Values values are constrained to:
//  * valid_navigable_target_name_or_keyword
func (e *AHTMLElement) TARGET(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["target"] = v
    return e
}

func (e *AHTMLElement) RemoveTARGET(v string) *AHTMLElement {
    delete(e.StringAttributes, "target")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *AHTMLElement) TITLE(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *AHTMLElement) RemoveTITLE(v string) *AHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *AHTMLElement) TRANSLATE(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *AHTMLElement) RemoveTRANSLATE(v string) *AHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *AHTMLElement) TYPE(v string) *AHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *AHTMLElement) RemoveTYPE(v string) *AHTMLElement {
    delete(e.StringAttributes, "type")
    return e
}
