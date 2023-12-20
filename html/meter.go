package  html

import (
    "fmt"
)

type MeterHTMLElement struct {
    *Element
}

func METER(children ...fmt.Stringer) *MeterHTMLElement {
    return &MeterHTMLElement{
        Element: &Element{
            Tag: "meter",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *MeterHTMLElement) Children(children ...fmt.Stringer) *MeterHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *MeterHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *MeterHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *MeterHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *MeterHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *MeterHTMLElement) Text(text string) *MeterHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *MeterHTMLElement) TextF(format string, args ...any) *MeterHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *MeterHTMLElement) Raw(text string) *MeterHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *MeterHTMLElement) RawF(format string, args ...any) *MeterHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *MeterHTMLElement) CustomData(key, value string) *MeterHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *MeterHTMLElement) CustomDataRemove(key string) *MeterHTMLElement {
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
func (e *MeterHTMLElement) ACCESSKEY(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *MeterHTMLElement) RemoveACCESSKEY(v string) *MeterHTMLElement {
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
func (e *MeterHTMLElement) AUTOCAPITALIZE(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *MeterHTMLElement) RemoveAUTOCAPITALIZE(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *MeterHTMLElement) AUTOFOCUS() *MeterHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *MeterHTMLElement) RemoveAUTOFOCUS() *MeterHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *MeterHTMLElement) SetAUTOFOCUS(b bool) *MeterHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *MeterHTMLElement) CLASS(v string) *MeterHTMLElement {
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

func (e *MeterHTMLElement) SetCLASS(v string) *MeterHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *MeterHTMLElement) RemoveCLASS(v string) *MeterHTMLElement {
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
func (e *MeterHTMLElement) CONTENTEDITABLE(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *MeterHTMLElement) RemoveCONTENTEDITABLE(v string) *MeterHTMLElement {
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
func (e *MeterHTMLElement) DIR(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *MeterHTMLElement) RemoveDIR(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *MeterHTMLElement) DRAGGABLE(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *MeterHTMLElement) RemoveDRAGGABLE(v string) *MeterHTMLElement {
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
func (e *MeterHTMLElement) ENTERKEYHINT(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *MeterHTMLElement) RemoveENTERKEYHINT(v string) *MeterHTMLElement {
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
func (e *MeterHTMLElement) HIDDEN(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *MeterHTMLElement) RemoveHIDDEN(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// HIGH sets the "high" attribute.
// Low limit of high range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterHTMLElement) HIGH(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["high"] = v
    return e
}

func (e *MeterHTMLElement) RemoveHIGH(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "high")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *MeterHTMLElement) ID(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *MeterHTMLElement) RemoveID(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *MeterHTMLElement) INERT() *MeterHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *MeterHTMLElement) RemoveINERT() *MeterHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *MeterHTMLElement) SetINERT(b bool) *MeterHTMLElement {
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
func (e *MeterHTMLElement) INPUTMODE(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *MeterHTMLElement) RemoveINPUTMODE(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *MeterHTMLElement) IS(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *MeterHTMLElement) RemoveIS(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *MeterHTMLElement) ITEMID(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *MeterHTMLElement) RemoveITEMID(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *MeterHTMLElement) ITEMPROP(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *MeterHTMLElement) RemoveITEMPROP(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *MeterHTMLElement) ITEMREF(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *MeterHTMLElement) RemoveITEMREF(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *MeterHTMLElement) ITEMSCOPE() *MeterHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *MeterHTMLElement) RemoveITEMSCOPE() *MeterHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *MeterHTMLElement) SetITEMSCOPE(b bool) *MeterHTMLElement {
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
func (e *MeterHTMLElement) ITEMTYPE(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *MeterHTMLElement) RemoveITEMTYPE(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *MeterHTMLElement) LANG(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *MeterHTMLElement) RemoveLANG(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LOW sets the "low" attribute.
// High limit of low range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterHTMLElement) LOW(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["low"] = v
    return e
}

func (e *MeterHTMLElement) RemoveLOW(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "low")
    return e
}
// MAX sets the "max" attribute.
// Upper bound of range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterHTMLElement) MAX(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["max"] = v
    return e
}

func (e *MeterHTMLElement) RemoveMAX(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "max")
    return e
}
// MIN sets the "min" attribute.
// Lower bound of range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterHTMLElement) MIN(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["min"] = v
    return e
}

func (e *MeterHTMLElement) RemoveMIN(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "min")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *MeterHTMLElement) NONCE(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *MeterHTMLElement) RemoveNONCE(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// OPTIMUM sets the "optimum" attribute.
// Optimum value in gauge
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterHTMLElement) OPTIMUM(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["optimum"] = v
    return e
}

func (e *MeterHTMLElement) RemoveOPTIMUM(v string) *MeterHTMLElement {
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
func (e *MeterHTMLElement) POPOVER(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *MeterHTMLElement) RemovePOPOVER(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *MeterHTMLElement) SLOT(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *MeterHTMLElement) RemoveSLOT(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *MeterHTMLElement) SPELLCHECK(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *MeterHTMLElement) RemoveSPELLCHECK(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *MeterHTMLElement) STYLE(k,v string) *MeterHTMLElement {
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

func (e *MeterHTMLElement) RemoveSTYLE(k string) *MeterHTMLElement {
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
func (e *MeterHTMLElement) TABINDEX(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *MeterHTMLElement) RemoveTABINDEX(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *MeterHTMLElement) TITLE(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *MeterHTMLElement) RemoveTITLE(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *MeterHTMLElement) TRANSLATE(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *MeterHTMLElement) RemoveTRANSLATE(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//  * valid_floating_point_number
func (e *MeterHTMLElement) VALUE(v string) *MeterHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["value"] = v
    return e
}

func (e *MeterHTMLElement) RemoveVALUE(v string) *MeterHTMLElement {
    delete(e.StringAttributes, "value")
    return e
}
