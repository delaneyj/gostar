package  html

import (
    "fmt"
)

type InsHTMLElement struct {
    *Element
}

func INS(children ...fmt.Stringer) *InsHTMLElement {
    return &InsHTMLElement{
        Element: &Element{
            Tag: "ins",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *InsHTMLElement) Children(children ...fmt.Stringer) *InsHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *InsHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *InsHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *InsHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *InsHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *InsHTMLElement) Text(text string) *InsHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *InsHTMLElement) TextF(format string, args ...any) *InsHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *InsHTMLElement) Raw(text string) *InsHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *InsHTMLElement) RawF(format string, args ...any) *InsHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *InsHTMLElement) CustomData(key, value string) *InsHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *InsHTMLElement) CustomDataRemove(key string) *InsHTMLElement {
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
func (e *InsHTMLElement) ACCESSKEY(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *InsHTMLElement) RemoveACCESSKEY(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) AUTOCAPITALIZE(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *InsHTMLElement) RemoveAUTOCAPITALIZE(v string) *InsHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *InsHTMLElement) AUTOFOCUS() *InsHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *InsHTMLElement) RemoveAUTOFOCUS() *InsHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *InsHTMLElement) SetAUTOFOCUS(b bool) *InsHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CITE sets the "cite" attribute.
// Link to the source of the quotation or more information about the edit
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *InsHTMLElement) CITE(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["cite"] = v
    return e
}

func (e *InsHTMLElement) RemoveCITE(v string) *InsHTMLElement {
    delete(e.StringAttributes, "cite")
    return e
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *InsHTMLElement) CLASS(v string) *InsHTMLElement {
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

func (e *InsHTMLElement) SetCLASS(v string) *InsHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *InsHTMLElement) RemoveCLASS(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) CONTENTEDITABLE(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *InsHTMLElement) RemoveCONTENTEDITABLE(v string) *InsHTMLElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// DATETIME sets the "datetime" attribute.
// Machine-readable value
// Values values are constrained to:
//  * valid_month_string
//  * valid_date_string
//  * valid_yearless_date_string
//  * valid_time_string
//  * valid_local_date_and_time_string
//  * valid_time_zone_offset_string
//  * valid_global_date_and_time_string
//  * valid_week_string
//  * valid_non_negative_integer
//  * valid_duration_string
func (e *InsHTMLElement) DATETIME(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["datetime"] = v
    return e
}

func (e *InsHTMLElement) RemoveDATETIME(v string) *InsHTMLElement {
    delete(e.StringAttributes, "datetime")
    return e
}
// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//  * ltr
//  * ltr
//  * rtl
//  * rtl
func (e *InsHTMLElement) DIR(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *InsHTMLElement) RemoveDIR(v string) *InsHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *InsHTMLElement) DRAGGABLE(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *InsHTMLElement) RemoveDRAGGABLE(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) ENTERKEYHINT(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *InsHTMLElement) RemoveENTERKEYHINT(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) HIDDEN(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *InsHTMLElement) RemoveHIDDEN(v string) *InsHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *InsHTMLElement) ID(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *InsHTMLElement) RemoveID(v string) *InsHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *InsHTMLElement) INERT() *InsHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *InsHTMLElement) RemoveINERT() *InsHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *InsHTMLElement) SetINERT(b bool) *InsHTMLElement {
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
func (e *InsHTMLElement) INPUTMODE(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *InsHTMLElement) RemoveINPUTMODE(v string) *InsHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *InsHTMLElement) IS(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *InsHTMLElement) RemoveIS(v string) *InsHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *InsHTMLElement) ITEMID(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *InsHTMLElement) RemoveITEMID(v string) *InsHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *InsHTMLElement) ITEMPROP(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *InsHTMLElement) RemoveITEMPROP(v string) *InsHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *InsHTMLElement) ITEMREF(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *InsHTMLElement) RemoveITEMREF(v string) *InsHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *InsHTMLElement) ITEMSCOPE() *InsHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *InsHTMLElement) RemoveITEMSCOPE() *InsHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *InsHTMLElement) SetITEMSCOPE(b bool) *InsHTMLElement {
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
func (e *InsHTMLElement) ITEMTYPE(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *InsHTMLElement) RemoveITEMTYPE(v string) *InsHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *InsHTMLElement) LANG(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *InsHTMLElement) RemoveLANG(v string) *InsHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *InsHTMLElement) NONCE(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *InsHTMLElement) RemoveNONCE(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) POPOVER(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *InsHTMLElement) RemovePOPOVER(v string) *InsHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *InsHTMLElement) SLOT(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *InsHTMLElement) RemoveSLOT(v string) *InsHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *InsHTMLElement) SPELLCHECK(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *InsHTMLElement) RemoveSPELLCHECK(v string) *InsHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *InsHTMLElement) STYLE(k,v string) *InsHTMLElement {
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

func (e *InsHTMLElement) RemoveSTYLE(k string) *InsHTMLElement {
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
func (e *InsHTMLElement) TABINDEX(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *InsHTMLElement) RemoveTABINDEX(v string) *InsHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *InsHTMLElement) TITLE(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *InsHTMLElement) RemoveTITLE(v string) *InsHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *InsHTMLElement) TRANSLATE(v string) *InsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *InsHTMLElement) RemoveTRANSLATE(v string) *InsHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
