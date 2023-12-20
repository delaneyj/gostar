package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type ScriptElement struct {
    *gostar.Element
}

func SCRIPT(children ...fmt.Stringer) *ScriptElement {
    return &ScriptElement{
        Element: &gostar.Element{
            Tag: "script",
            IsSelfClosing: true,
            Children: children,
        },
    }
}

func (e *ScriptElement) AddChildren(children ...fmt.Stringer) *ScriptElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *ScriptElement) TEXT(text string) *ScriptElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *ScriptElement) RAW(text string) *ScriptElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *ScriptElement) CustomData(key, value string) *ScriptElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ScriptElement) CustomDataRemove(key string) *ScriptElement {
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
func (e *ScriptElement) ACCESSKEY(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *ScriptElement) RemoveACCESSKEY(v string) *ScriptElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ASYNC sets the "async" attribute.
// Execute script when available, without blocking while fetching
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptElement) ASYNC() *ScriptElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["async"] = struct{}{}
    return e
}

func (e *ScriptElement) RemoveASYNC() *ScriptElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "async")
    return e
}

func (e *ScriptElement) SetASYNC(b bool) *ScriptElement {
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
func (e *ScriptElement) AUTOCAPITALIZE(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *ScriptElement) RemoveAUTOCAPITALIZE(v string) *ScriptElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptElement) AUTOFOCUS() *ScriptElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *ScriptElement) RemoveAUTOFOCUS() *ScriptElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *ScriptElement) SetAUTOFOCUS(b bool) *ScriptElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// BLOCKING sets the "blocking" attribute.
// Whether the element is potentially render-blocking
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ScriptElement) BLOCKING(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["blocking"] = v
    return e
}

func (e *ScriptElement) RemoveBLOCKING(v string) *ScriptElement {
    delete(e.StringAttributes, "blocking")
    return e
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *ScriptElement) CLASS(v string) *ScriptElement {
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

func (e *ScriptElement) SetCLASS(v string) *ScriptElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *ScriptElement) RemoveCLASS(v string) *ScriptElement {
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
func (e *ScriptElement) CONTENTEDITABLE(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *ScriptElement) RemoveCONTENTEDITABLE(v string) *ScriptElement {
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
func (e *ScriptElement) CROSSORIGIN(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *ScriptElement) RemoveCROSSORIGIN(v string) *ScriptElement {
    delete(e.StringAttributes, "crossorigin")
    return e
}
// DEFER sets the "defer" attribute.
// Defer script execution
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptElement) DEFER() *ScriptElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["defer"] = struct{}{}
    return e
}

func (e *ScriptElement) RemoveDEFER() *ScriptElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "defer")
    return e
}

func (e *ScriptElement) SetDEFER(b bool) *ScriptElement {
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
func (e *ScriptElement) DIR(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *ScriptElement) RemoveDIR(v string) *ScriptElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *ScriptElement) DRAGGABLE(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *ScriptElement) RemoveDRAGGABLE(v string) *ScriptElement {
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
func (e *ScriptElement) ENTERKEYHINT(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *ScriptElement) RemoveENTERKEYHINT(v string) *ScriptElement {
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
func (e *ScriptElement) FETCHPRIORITY(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["fetchpriority"] = v
    return e
}

func (e *ScriptElement) RemoveFETCHPRIORITY(v string) *ScriptElement {
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
func (e *ScriptElement) HIDDEN(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *ScriptElement) RemoveHIDDEN(v string) *ScriptElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *ScriptElement) ID(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *ScriptElement) RemoveID(v string) *ScriptElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptElement) INERT() *ScriptElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *ScriptElement) RemoveINERT() *ScriptElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *ScriptElement) SetINERT(b bool) *ScriptElement {
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
func (e *ScriptElement) INPUTMODE(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *ScriptElement) RemoveINPUTMODE(v string) *ScriptElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// INTEGRITY sets the "integrity" attribute.
// Integrity metadata used in Subresource Integrity checks [SRI]
// Values values are constrained to:
//  * text
func (e *ScriptElement) INTEGRITY(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["integrity"] = v
    return e
}

func (e *ScriptElement) RemoveINTEGRITY(v string) *ScriptElement {
    delete(e.StringAttributes, "integrity")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *ScriptElement) IS(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *ScriptElement) RemoveIS(v string) *ScriptElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *ScriptElement) ITEMID(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *ScriptElement) RemoveITEMID(v string) *ScriptElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *ScriptElement) ITEMPROP(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *ScriptElement) RemoveITEMPROP(v string) *ScriptElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ScriptElement) ITEMREF(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *ScriptElement) RemoveITEMREF(v string) *ScriptElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptElement) ITEMSCOPE() *ScriptElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *ScriptElement) RemoveITEMSCOPE() *ScriptElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *ScriptElement) SetITEMSCOPE(b bool) *ScriptElement {
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
func (e *ScriptElement) ITEMTYPE(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *ScriptElement) RemoveITEMTYPE(v string) *ScriptElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ScriptElement) LANG(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *ScriptElement) RemoveLANG(v string) *ScriptElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NOMODULE sets the "nomodule" attribute.
// Prevents execution in user agents that support module scripts
// Values values are constrained to:
//  * boolean_attribute
func (e *ScriptElement) NOMODULE() *ScriptElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["nomodule"] = struct{}{}
    return e
}

func (e *ScriptElement) RemoveNOMODULE() *ScriptElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "nomodule")
    return e
}

func (e *ScriptElement) SetNOMODULE(b bool) *ScriptElement {
    if b {
        return e.NOMODULE()
    }
    return e.RemoveNOMODULE()
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *ScriptElement) NONCE(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *ScriptElement) RemoveNONCE(v string) *ScriptElement {
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
func (e *ScriptElement) POPOVER(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *ScriptElement) RemovePOPOVER(v string) *ScriptElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *ScriptElement) REFERRERPOLICY(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *ScriptElement) RemoveREFERRERPOLICY(v string) *ScriptElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *ScriptElement) SLOT(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *ScriptElement) RemoveSLOT(v string) *ScriptElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *ScriptElement) SPELLCHECK(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *ScriptElement) RemoveSPELLCHECK(v string) *ScriptElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ScriptElement) SRC(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *ScriptElement) RemoveSRC(v string) *ScriptElement {
    delete(e.StringAttributes, "src")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ScriptElement) STYLE(k,v string) *ScriptElement {
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

func (e *ScriptElement) RemoveSTYLE(k string) *ScriptElement {
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
func (e *ScriptElement) TABINDEX(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *ScriptElement) RemoveTABINDEX(v string) *ScriptElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *ScriptElement) TITLE(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *ScriptElement) RemoveTITLE(v string) *ScriptElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *ScriptElement) TRANSLATE(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *ScriptElement) RemoveTRANSLATE(v string) *ScriptElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *ScriptElement) TYPE(v string) *ScriptElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *ScriptElement) RemoveTYPE(v string) *ScriptElement {
    delete(e.StringAttributes, "type")
    return e
}
