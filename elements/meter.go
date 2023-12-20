package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type MeterElement struct {
    *gostar.Element
}

func METER(children ...fmt.Stringer) *MeterElement {
    return &MeterElement{
        Element: &gostar.Element{
            Tag: "meter",
            IsSelfClosing: false,
            Children: children,
        },
    }
}

func (e *MeterElement) AddChildren(children ...fmt.Stringer) *MeterElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *MeterElement) TEXT(text string) *MeterElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *MeterElement) RAW(text string) *MeterElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *MeterElement) CustomData(key, value string) *MeterElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *MeterElement) CustomDataRemove(key string) *MeterElement {
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
func (e *MeterElement) ACCESSKEY(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *MeterElement) RemoveACCESSKEY(v string) *MeterElement {
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
func (e *MeterElement) AUTOCAPITALIZE(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *MeterElement) RemoveAUTOCAPITALIZE(v string) *MeterElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *MeterElement) AUTOFOCUS() *MeterElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *MeterElement) RemoveAUTOFOCUS() *MeterElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *MeterElement) SetAUTOFOCUS(b bool) *MeterElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *MeterElement) CLASS(v string) *MeterElement {
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

func (e *MeterElement) SetCLASS(v string) *MeterElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *MeterElement) RemoveCLASS(v string) *MeterElement {
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
func (e *MeterElement) CONTENTEDITABLE(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *MeterElement) RemoveCONTENTEDITABLE(v string) *MeterElement {
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
func (e *MeterElement) DIR(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *MeterElement) RemoveDIR(v string) *MeterElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *MeterElement) DRAGGABLE(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *MeterElement) RemoveDRAGGABLE(v string) *MeterElement {
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
func (e *MeterElement) ENTERKEYHINT(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *MeterElement) RemoveENTERKEYHINT(v string) *MeterElement {
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
func (e *MeterElement) HIDDEN(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *MeterElement) RemoveHIDDEN(v string) *MeterElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// HIGH sets the "high" attribute.
// Low limit of high range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterElement) HIGH(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["high"] = v
    return e
}

func (e *MeterElement) RemoveHIGH(v string) *MeterElement {
    delete(e.StringAttributes, "high")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *MeterElement) ID(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *MeterElement) RemoveID(v string) *MeterElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *MeterElement) INERT() *MeterElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *MeterElement) RemoveINERT() *MeterElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *MeterElement) SetINERT(b bool) *MeterElement {
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
func (e *MeterElement) INPUTMODE(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *MeterElement) RemoveINPUTMODE(v string) *MeterElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *MeterElement) IS(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *MeterElement) RemoveIS(v string) *MeterElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *MeterElement) ITEMID(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *MeterElement) RemoveITEMID(v string) *MeterElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *MeterElement) ITEMPROP(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *MeterElement) RemoveITEMPROP(v string) *MeterElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *MeterElement) ITEMREF(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *MeterElement) RemoveITEMREF(v string) *MeterElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *MeterElement) ITEMSCOPE() *MeterElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *MeterElement) RemoveITEMSCOPE() *MeterElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *MeterElement) SetITEMSCOPE(b bool) *MeterElement {
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
func (e *MeterElement) ITEMTYPE(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *MeterElement) RemoveITEMTYPE(v string) *MeterElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *MeterElement) LANG(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *MeterElement) RemoveLANG(v string) *MeterElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LOW sets the "low" attribute.
// High limit of low range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterElement) LOW(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["low"] = v
    return e
}

func (e *MeterElement) RemoveLOW(v string) *MeterElement {
    delete(e.StringAttributes, "low")
    return e
}
// MAX sets the "max" attribute.
// Upper bound of range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterElement) MAX(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["max"] = v
    return e
}

func (e *MeterElement) RemoveMAX(v string) *MeterElement {
    delete(e.StringAttributes, "max")
    return e
}
// MIN sets the "min" attribute.
// Lower bound of range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterElement) MIN(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["min"] = v
    return e
}

func (e *MeterElement) RemoveMIN(v string) *MeterElement {
    delete(e.StringAttributes, "min")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *MeterElement) NONCE(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *MeterElement) RemoveNONCE(v string) *MeterElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// OPTIMUM sets the "optimum" attribute.
// Optimum value in gauge
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterElement) OPTIMUM(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["optimum"] = v
    return e
}

func (e *MeterElement) RemoveOPTIMUM(v string) *MeterElement {
    delete(e.StringAttributes, "optimum")
    return e
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *MeterElement) POPOVER(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *MeterElement) RemovePOPOVER(v string) *MeterElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *MeterElement) SLOT(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *MeterElement) RemoveSLOT(v string) *MeterElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *MeterElement) SPELLCHECK(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *MeterElement) RemoveSPELLCHECK(v string) *MeterElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *MeterElement) STYLE(k,v string) *MeterElement {
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

func (e *MeterElement) RemoveSTYLE(k string) *MeterElement {
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
func (e *MeterElement) TABINDEX(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *MeterElement) RemoveTABINDEX(v string) *MeterElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *MeterElement) TITLE(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *MeterElement) RemoveTITLE(v string) *MeterElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *MeterElement) TRANSLATE(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *MeterElement) RemoveTRANSLATE(v string) *MeterElement {
    delete(e.StringAttributes, "translate")
    return e
}
// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterElement) VALUE(v string) *MeterElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["value"] = v
    return e
}

func (e *MeterElement) RemoveVALUE(v string) *MeterElement {
    delete(e.StringAttributes, "value")
    return e
}
