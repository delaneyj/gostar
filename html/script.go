package  html

import (
    "fmt"
)

type ScriptHTMLElement struct {
    *Element
}

func SCRIPT(children ...ElementBuilder) *ScriptHTMLElement {
    return &ScriptHTMLElement{
        Element: &Element{
            Tag: "script",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *ScriptHTMLElement) Children(children ...ElementBuilder) *ScriptHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *ScriptHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ScriptHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *ScriptHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ScriptHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *ScriptHTMLElement) Text(text string) *ScriptHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *ScriptHTMLElement) TextF(format string, args ...any) *ScriptHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *ScriptHTMLElement) Raw(text string) *ScriptHTMLElement {
    e.Descendants = append(e.Descendants, Raw(text))
    return e
}

func (e *ScriptHTMLElement) RawF(format string, args ...any) *ScriptHTMLElement {
    return e.Raw(fmt.Sprintf(format, args...))
}

func (e *ScriptHTMLElement) CustomData(key, value string) *ScriptHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ScriptHTMLElement) CustomDataRemove(key string) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) ACCESSKEY(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveACCESSKEY(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ASYNC sets the "async" attribute.
// Execute script when available, without blocking while fetching
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptHTMLElement) ASYNC() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["async"] = struct{}{}
    return e
}

func (e *ScriptHTMLElement) RemoveASYNC() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "async")
    return e
}

func (e *ScriptHTMLElement) SetASYNC(b bool) *ScriptHTMLElement {
    if b {
        return e.ASYNC()
    }
    return e.RemoveASYNC()
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
func (e *ScriptHTMLElement) AUTOCAPITALIZE(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveAUTOCAPITALIZE(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptHTMLElement) AUTOFOCUS() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *ScriptHTMLElement) RemoveAUTOFOCUS() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *ScriptHTMLElement) SetAUTOFOCUS(b bool) *ScriptHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// BLOCKING sets the "blocking" attribute.
// Whether the element is potentially render-blocking
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ScriptHTMLElement) BLOCKING(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["blocking"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveBLOCKING(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "blocking")
    return e
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *ScriptHTMLElement) CLASS(v string) *ScriptHTMLElement {
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

func (e *ScriptHTMLElement) SetCLASS(v string) *ScriptHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *ScriptHTMLElement) RemoveCLASS(v string) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) CONTENTEDITABLE(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveCONTENTEDITABLE(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// CROSSORIGIN sets the "crossorigin" attribute.
// How the element handles crossorigin requests
// Values values are constrained to:
//  * anonymous
//  * anonymous
//  * use_credentials
//  * use_credentials
func (e *ScriptHTMLElement) CROSSORIGIN(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveCROSSORIGIN(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "crossorigin")
    return e
}
// DEFER sets the "defer" attribute.
// Defer script execution
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptHTMLElement) DEFER() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["defer"] = struct{}{}
    return e
}

func (e *ScriptHTMLElement) RemoveDEFER() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "defer")
    return e
}

func (e *ScriptHTMLElement) SetDEFER(b bool) *ScriptHTMLElement {
    if b {
        return e.DEFER()
    }
    return e.RemoveDEFER()
}
// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//  * ltr
//  * ltr
//  * rtl
//  * rtl
func (e *ScriptHTMLElement) DIR(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveDIR(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *ScriptHTMLElement) DRAGGABLE(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveDRAGGABLE(v string) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) ENTERKEYHINT(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveENTERKEYHINT(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FETCHPRIORITY sets the "fetchpriority" attribute.
// Sets the priority for fetches initiated by the element
// Values values are constrained to:
//  * auto
//  * auto
//  * high
//  * high
//  * low
//  * low
func (e *ScriptHTMLElement) FETCHPRIORITY(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["fetchpriority"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveFETCHPRIORITY(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "fetchpriority")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *ScriptHTMLElement) HIDDEN(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveHIDDEN(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *ScriptHTMLElement) ID(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveID(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptHTMLElement) INERT() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *ScriptHTMLElement) RemoveINERT() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *ScriptHTMLElement) SetINERT(b bool) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) INPUTMODE(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveINPUTMODE(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// INTEGRITY sets the "integrity" attribute.
// Integrity metadata used in Subresource Integrity checks [SRI]
// Values values are constrained to:
//  * text
func (e *ScriptHTMLElement) INTEGRITY(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["integrity"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveINTEGRITY(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "integrity")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *ScriptHTMLElement) IS(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveIS(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *ScriptHTMLElement) ITEMID(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveITEMID(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *ScriptHTMLElement) ITEMPROP(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveITEMPROP(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ScriptHTMLElement) ITEMREF(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveITEMREF(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptHTMLElement) ITEMSCOPE() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *ScriptHTMLElement) RemoveITEMSCOPE() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *ScriptHTMLElement) SetITEMSCOPE(b bool) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) ITEMTYPE(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveITEMTYPE(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ScriptHTMLElement) LANG(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveLANG(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NOMODULE sets the "nomodule" attribute.
// Prevents execution in user agents that support module scripts
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptHTMLElement) NOMODULE() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["nomodule"] = struct{}{}
    return e
}

func (e *ScriptHTMLElement) RemoveNOMODULE() *ScriptHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "nomodule")
    return e
}

func (e *ScriptHTMLElement) SetNOMODULE(b bool) *ScriptHTMLElement {
    if b {
        return e.NOMODULE()
    }
    return e.RemoveNOMODULE()
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *ScriptHTMLElement) NONCE(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveNONCE(v string) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) POPOVER(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *ScriptHTMLElement) RemovePOPOVER(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *ScriptHTMLElement) REFERRERPOLICY(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveREFERRERPOLICY(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *ScriptHTMLElement) SLOT(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveSLOT(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *ScriptHTMLElement) SPELLCHECK(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveSPELLCHECK(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ScriptHTMLElement) SRC(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveSRC(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "src")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ScriptHTMLElement) STYLE(k,v string) *ScriptHTMLElement {
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

func (e *ScriptHTMLElement) RemoveSTYLE(k string) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) TABINDEX(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveTABINDEX(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *ScriptHTMLElement) TITLE(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveTITLE(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *ScriptHTMLElement) TRANSLATE(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveTRANSLATE(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *ScriptHTMLElement) TYPE(v string) *ScriptHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *ScriptHTMLElement) RemoveTYPE(v string) *ScriptHTMLElement {
    delete(e.StringAttributes, "type")
    return e
}
