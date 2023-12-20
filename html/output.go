package  html

import (
    "fmt"
)

type OutputHTMLElement struct {
    *Element
}

func OUTPUT(children ...fmt.Stringer) *OutputHTMLElement {
    return &OutputHTMLElement{
        Element: &Element{
            Tag: "output",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *OutputHTMLElement) Children(children ...fmt.Stringer) *OutputHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *OutputHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *OutputHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *OutputHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *OutputHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *OutputHTMLElement) Text(text string) *OutputHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *OutputHTMLElement) TextF(format string, args ...any) *OutputHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *OutputHTMLElement) Raw(text string) *OutputHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *OutputHTMLElement) RawF(format string, args ...any) *OutputHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *OutputHTMLElement) CustomData(key, value string) *OutputHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *OutputHTMLElement) CustomDataRemove(key string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) ACCESSKEY(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *OutputHTMLElement) RemoveACCESSKEY(v string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) AUTOCAPITALIZE(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *OutputHTMLElement) RemoveAUTOCAPITALIZE(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *OutputHTMLElement) AUTOFOCUS() *OutputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *OutputHTMLElement) RemoveAUTOFOCUS() *OutputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *OutputHTMLElement) SetAUTOFOCUS(b bool) *OutputHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *OutputHTMLElement) CLASS(v string) *OutputHTMLElement {
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

func (e *OutputHTMLElement) SetCLASS(v string) *OutputHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *OutputHTMLElement) RemoveCLASS(v string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) CONTENTEDITABLE(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *OutputHTMLElement) RemoveCONTENTEDITABLE(v string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) DIR(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *OutputHTMLElement) RemoveDIR(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *OutputHTMLElement) DRAGGABLE(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *OutputHTMLElement) RemoveDRAGGABLE(v string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) ENTERKEYHINT(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *OutputHTMLElement) RemoveENTERKEYHINT(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FOR sets the "for" attribute.
// Specifies controls from which the output was calculated
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *OutputHTMLElement) FOR(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["for"] = v
    return e
}

func (e *OutputHTMLElement) RemoveFOR(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "for")
    return e
}
// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//  * id
func (e *OutputHTMLElement) FORM(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["form"] = v
    return e
}

func (e *OutputHTMLElement) RemoveFORM(v string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) HIDDEN(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *OutputHTMLElement) RemoveHIDDEN(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *OutputHTMLElement) ID(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *OutputHTMLElement) RemoveID(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *OutputHTMLElement) INERT() *OutputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *OutputHTMLElement) RemoveINERT() *OutputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *OutputHTMLElement) SetINERT(b bool) *OutputHTMLElement {
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
func (e *OutputHTMLElement) INPUTMODE(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *OutputHTMLElement) RemoveINPUTMODE(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *OutputHTMLElement) IS(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *OutputHTMLElement) RemoveIS(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *OutputHTMLElement) ITEMID(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *OutputHTMLElement) RemoveITEMID(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *OutputHTMLElement) ITEMPROP(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *OutputHTMLElement) RemoveITEMPROP(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *OutputHTMLElement) ITEMREF(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *OutputHTMLElement) RemoveITEMREF(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *OutputHTMLElement) ITEMSCOPE() *OutputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *OutputHTMLElement) RemoveITEMSCOPE() *OutputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *OutputHTMLElement) SetITEMSCOPE(b bool) *OutputHTMLElement {
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
func (e *OutputHTMLElement) ITEMTYPE(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *OutputHTMLElement) RemoveITEMTYPE(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *OutputHTMLElement) LANG(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *OutputHTMLElement) RemoveLANG(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *OutputHTMLElement) NAME(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *OutputHTMLElement) RemoveNAME(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *OutputHTMLElement) NONCE(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *OutputHTMLElement) RemoveNONCE(v string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) POPOVER(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *OutputHTMLElement) RemovePOPOVER(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *OutputHTMLElement) SLOT(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *OutputHTMLElement) RemoveSLOT(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *OutputHTMLElement) SPELLCHECK(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *OutputHTMLElement) RemoveSPELLCHECK(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *OutputHTMLElement) STYLE(k,v string) *OutputHTMLElement {
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

func (e *OutputHTMLElement) RemoveSTYLE(k string) *OutputHTMLElement {
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
func (e *OutputHTMLElement) TABINDEX(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *OutputHTMLElement) RemoveTABINDEX(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *OutputHTMLElement) TITLE(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *OutputHTMLElement) RemoveTITLE(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *OutputHTMLElement) TRANSLATE(v string) *OutputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *OutputHTMLElement) RemoveTRANSLATE(v string) *OutputHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
