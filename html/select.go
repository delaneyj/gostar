package  html

import (
    "fmt"
)

type SelectHTMLElement struct {
    *Element
}

func SELECT(children ...fmt.Stringer) *SelectHTMLElement {
    return &SelectHTMLElement{
        Element: &Element{
            Tag: "select",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *SelectHTMLElement) Children(children ...fmt.Stringer) *SelectHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *SelectHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *SelectHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *SelectHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *SelectHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *SelectHTMLElement) Text(text string) *SelectHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *SelectHTMLElement) TextF(format string, args ...any) *SelectHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *SelectHTMLElement) Raw(text string) *SelectHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *SelectHTMLElement) RawF(format string, args ...any) *SelectHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *SelectHTMLElement) CustomData(key, value string) *SelectHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SelectHTMLElement) CustomDataRemove(key string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) ACCESSKEY(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *SelectHTMLElement) RemoveACCESSKEY(v string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) AUTOCAPITALIZE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *SelectHTMLElement) RemoveAUTOCAPITALIZE(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//  * autofill_field
func (e *SelectHTMLElement) AUTOCOMPLETE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocomplete"] = v
    return e
}

func (e *SelectHTMLElement) RemoveAUTOCOMPLETE(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "autocomplete")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *SelectHTMLElement) AUTOFOCUS() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *SelectHTMLElement) RemoveAUTOFOCUS() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *SelectHTMLElement) SetAUTOFOCUS(b bool) *SelectHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *SelectHTMLElement) CLASS(v string) *SelectHTMLElement {
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

func (e *SelectHTMLElement) SetCLASS(v string) *SelectHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *SelectHTMLElement) RemoveCLASS(v string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) CONTENTEDITABLE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *SelectHTMLElement) RemoveCONTENTEDITABLE(v string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) DIR(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *SelectHTMLElement) RemoveDIR(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *SelectHTMLElement) DISABLED() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *SelectHTMLElement) RemoveDISABLED() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *SelectHTMLElement) SetDISABLED(b bool) *SelectHTMLElement {
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
func (e *SelectHTMLElement) DRAGGABLE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *SelectHTMLElement) RemoveDRAGGABLE(v string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) ENTERKEYHINT(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *SelectHTMLElement) RemoveENTERKEYHINT(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//  * id
func (e *SelectHTMLElement) FORM(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["form"] = v
    return e
}

func (e *SelectHTMLElement) RemoveFORM(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "form")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *SelectHTMLElement) HIDDEN(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *SelectHTMLElement) RemoveHIDDEN(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *SelectHTMLElement) ID(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *SelectHTMLElement) RemoveID(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *SelectHTMLElement) INERT() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *SelectHTMLElement) RemoveINERT() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *SelectHTMLElement) SetINERT(b bool) *SelectHTMLElement {
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
func (e *SelectHTMLElement) INPUTMODE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *SelectHTMLElement) RemoveINPUTMODE(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *SelectHTMLElement) IS(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *SelectHTMLElement) RemoveIS(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *SelectHTMLElement) ITEMID(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *SelectHTMLElement) RemoveITEMID(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *SelectHTMLElement) ITEMPROP(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *SelectHTMLElement) RemoveITEMPROP(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *SelectHTMLElement) ITEMREF(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *SelectHTMLElement) RemoveITEMREF(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *SelectHTMLElement) ITEMSCOPE() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *SelectHTMLElement) RemoveITEMSCOPE() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *SelectHTMLElement) SetITEMSCOPE(b bool) *SelectHTMLElement {
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
func (e *SelectHTMLElement) ITEMTYPE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *SelectHTMLElement) RemoveITEMTYPE(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SelectHTMLElement) LANG(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *SelectHTMLElement) RemoveLANG(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// MULTIPLE sets the "multiple" attribute.
// Whether to allow multiple values
// Values values are constrained to:
//  * boolean_attribute
func (e *SelectHTMLElement) MULTIPLE() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["multiple"] = struct{}{}
    return e
}

func (e *SelectHTMLElement) RemoveMULTIPLE() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "multiple")
    return e
}

func (e *SelectHTMLElement) SetMULTIPLE(b bool) *SelectHTMLElement {
    if b {
        return e.MULTIPLE()
    }
    return e.RemoveMULTIPLE()
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *SelectHTMLElement) NAME(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *SelectHTMLElement) RemoveNAME(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *SelectHTMLElement) NONCE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *SelectHTMLElement) RemoveNONCE(v string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) POPOVER(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *SelectHTMLElement) RemovePOPOVER(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REQUIRED sets the "required" attribute.
// Whether the control is required for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *SelectHTMLElement) REQUIRED() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["required"] = struct{}{}
    return e
}

func (e *SelectHTMLElement) RemoveREQUIRED() *SelectHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "required")
    return e
}

func (e *SelectHTMLElement) SetREQUIRED(b bool) *SelectHTMLElement {
    if b {
        return e.REQUIRED()
    }
    return e.RemoveREQUIRED()
}
// SIZE sets the "size" attribute.
// Size of the control
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *SelectHTMLElement) SIZE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["size"] = v
    return e
}

func (e *SelectHTMLElement) RemoveSIZE(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "size")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *SelectHTMLElement) SLOT(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *SelectHTMLElement) RemoveSLOT(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *SelectHTMLElement) SPELLCHECK(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *SelectHTMLElement) RemoveSPELLCHECK(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SelectHTMLElement) STYLE(k,v string) *SelectHTMLElement {
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

func (e *SelectHTMLElement) RemoveSTYLE(k string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) TABINDEX(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *SelectHTMLElement) RemoveTABINDEX(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *SelectHTMLElement) TITLE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *SelectHTMLElement) RemoveTITLE(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *SelectHTMLElement) TRANSLATE(v string) *SelectHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *SelectHTMLElement) RemoveTRANSLATE(v string) *SelectHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
