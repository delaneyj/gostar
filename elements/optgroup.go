package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type OptgroupElement struct {
    *gostar.Element
}

func OPTGROUP(children ...fmt.Stringer) *OptgroupElement {
    return &OptgroupElement{
        Element: &gostar.Element{
            Tag: "optgroup",
            IsSelfClosing: false,
            Children: children,
        },
    }
}

func (e *OptgroupElement) AddChildren(children ...fmt.Stringer) *OptgroupElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *OptgroupElement) TEXT(text string) *OptgroupElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *OptgroupElement) RAW(text string) *OptgroupElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *OptgroupElement) CustomData(key, value string) *OptgroupElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *OptgroupElement) CustomDataRemove(key string) *OptgroupElement {
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
func (e *OptgroupElement) ACCESSKEY(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *OptgroupElement) RemoveACCESSKEY(v string) *OptgroupElement {
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
func (e *OptgroupElement) AUTOCAPITALIZE(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *OptgroupElement) RemoveAUTOCAPITALIZE(v string) *OptgroupElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *OptgroupElement) AUTOFOCUS() *OptgroupElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *OptgroupElement) RemoveAUTOFOCUS() *OptgroupElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *OptgroupElement) SetAUTOFOCUS(b bool) *OptgroupElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *OptgroupElement) CLASS(v string) *OptgroupElement {
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

func (e *OptgroupElement) SetCLASS(v string) *OptgroupElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *OptgroupElement) RemoveCLASS(v string) *OptgroupElement {
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
func (e *OptgroupElement) CONTENTEDITABLE(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *OptgroupElement) RemoveCONTENTEDITABLE(v string) *OptgroupElement {
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
func (e *OptgroupElement) DIR(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *OptgroupElement) RemoveDIR(v string) *OptgroupElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *OptgroupElement) DISABLED() *OptgroupElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *OptgroupElement) RemoveDISABLED() *OptgroupElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *OptgroupElement) SetDISABLED(b bool) *OptgroupElement {
    if b {
        return e.DISABLED()
    }
    return e.RemoveDISABLED()
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *OptgroupElement) DRAGGABLE(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *OptgroupElement) RemoveDRAGGABLE(v string) *OptgroupElement {
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
func (e *OptgroupElement) ENTERKEYHINT(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *OptgroupElement) RemoveENTERKEYHINT(v string) *OptgroupElement {
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
func (e *OptgroupElement) HIDDEN(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *OptgroupElement) RemoveHIDDEN(v string) *OptgroupElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *OptgroupElement) ID(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *OptgroupElement) RemoveID(v string) *OptgroupElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *OptgroupElement) INERT() *OptgroupElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *OptgroupElement) RemoveINERT() *OptgroupElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *OptgroupElement) SetINERT(b bool) *OptgroupElement {
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
func (e *OptgroupElement) INPUTMODE(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *OptgroupElement) RemoveINPUTMODE(v string) *OptgroupElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *OptgroupElement) IS(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *OptgroupElement) RemoveIS(v string) *OptgroupElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *OptgroupElement) ITEMID(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *OptgroupElement) RemoveITEMID(v string) *OptgroupElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *OptgroupElement) ITEMPROP(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *OptgroupElement) RemoveITEMPROP(v string) *OptgroupElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *OptgroupElement) ITEMREF(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *OptgroupElement) RemoveITEMREF(v string) *OptgroupElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *OptgroupElement) ITEMSCOPE() *OptgroupElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *OptgroupElement) RemoveITEMSCOPE() *OptgroupElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *OptgroupElement) SetITEMSCOPE(b bool) *OptgroupElement {
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
func (e *OptgroupElement) ITEMTYPE(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *OptgroupElement) RemoveITEMTYPE(v string) *OptgroupElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LABEL sets the "label" attribute.
// User-visible label
// Values values are constrained to:
//  * text
func (e *OptgroupElement) LABEL(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["label"] = v
    return e
}

func (e *OptgroupElement) RemoveLABEL(v string) *OptgroupElement {
    delete(e.StringAttributes, "label")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *OptgroupElement) LANG(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *OptgroupElement) RemoveLANG(v string) *OptgroupElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *OptgroupElement) NONCE(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *OptgroupElement) RemoveNONCE(v string) *OptgroupElement {
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
func (e *OptgroupElement) POPOVER(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *OptgroupElement) RemovePOPOVER(v string) *OptgroupElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *OptgroupElement) SLOT(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *OptgroupElement) RemoveSLOT(v string) *OptgroupElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *OptgroupElement) SPELLCHECK(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *OptgroupElement) RemoveSPELLCHECK(v string) *OptgroupElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *OptgroupElement) STYLE(k,v string) *OptgroupElement {
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

func (e *OptgroupElement) RemoveSTYLE(k string) *OptgroupElement {
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
func (e *OptgroupElement) TABINDEX(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *OptgroupElement) RemoveTABINDEX(v string) *OptgroupElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *OptgroupElement) TITLE(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *OptgroupElement) RemoveTITLE(v string) *OptgroupElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *OptgroupElement) TRANSLATE(v string) *OptgroupElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *OptgroupElement) RemoveTRANSLATE(v string) *OptgroupElement {
    delete(e.StringAttributes, "translate")
    return e
}
