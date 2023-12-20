package  html

import (
    "fmt"
)

type SubHTMLElement struct {
    *Element
}

func SUB(children ...ElementBuilder) *SubHTMLElement {
    return &SubHTMLElement{
        Element: &Element{
            Tag: "sub",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *SubHTMLElement) Children(children ...ElementBuilder) *SubHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *SubHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SubHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *SubHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SubHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *SubHTMLElement) Text(text string) *SubHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *SubHTMLElement) TextF(format string, args ...any) *SubHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *SubHTMLElement) Escaped(text string) *SubHTMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *SubHTMLElement) EscapedF(format string, args ...any) *SubHTMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SubHTMLElement) CustomData(key, value string) *SubHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SubHTMLElement) CustomDataRemove(key string) *SubHTMLElement {
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
func (e *SubHTMLElement) ACCESSKEY(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *SubHTMLElement) IfACCESSKEY(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.ACCESSKEY(v)
}

func (e *SubHTMLElement) RemoveACCESSKEY(v string) *SubHTMLElement {
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
func (e *SubHTMLElement) AUTOCAPITALIZE(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *SubHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCAPITALIZE(v)
}

func (e *SubHTMLElement) RemoveAUTOCAPITALIZE(v string) *SubHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *SubHTMLElement) AUTOFOCUS() *SubHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *SubHTMLElement) IfAUTOFOCUS(cond bool) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOFOCUS()
}

func (e *SubHTMLElement) RemoveAUTOFOCUS() *SubHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *SubHTMLElement) SetAUTOFOCUS(b bool) *SubHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *SubHTMLElement) CLASS(v string) *SubHTMLElement {
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

func (e *SubHTMLElement) IfCLASS(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.CLASS(v)
}

func (e *SubHTMLElement) SetCLASS(v string) *SubHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *SubHTMLElement) RemoveCLASS(v string) *SubHTMLElement {
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
func (e *SubHTMLElement) CONTENTEDITABLE(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *SubHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.CONTENTEDITABLE(v)
}

func (e *SubHTMLElement) RemoveCONTENTEDITABLE(v string) *SubHTMLElement {
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
func (e *SubHTMLElement) DIR(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *SubHTMLElement) IfDIR(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.DIR(v)
}

func (e *SubHTMLElement) RemoveDIR(v string) *SubHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *SubHTMLElement) DRAGGABLE(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *SubHTMLElement) IfDRAGGABLE(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.DRAGGABLE(v)
}

func (e *SubHTMLElement) RemoveDRAGGABLE(v string) *SubHTMLElement {
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
func (e *SubHTMLElement) ENTERKEYHINT(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *SubHTMLElement) IfENTERKEYHINT(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.ENTERKEYHINT(v)
}

func (e *SubHTMLElement) RemoveENTERKEYHINT(v string) *SubHTMLElement {
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
func (e *SubHTMLElement) HIDDEN(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *SubHTMLElement) IfHIDDEN(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.HIDDEN(v)
}

func (e *SubHTMLElement) RemoveHIDDEN(v string) *SubHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *SubHTMLElement) ID(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *SubHTMLElement) IfID(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.ID(v)
}

func (e *SubHTMLElement) RemoveID(v string) *SubHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *SubHTMLElement) INERT() *SubHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *SubHTMLElement) IfINERT(cond bool) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.INERT()
}

func (e *SubHTMLElement) RemoveINERT() *SubHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *SubHTMLElement) SetINERT(b bool) *SubHTMLElement {
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
func (e *SubHTMLElement) INPUTMODE(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *SubHTMLElement) IfINPUTMODE(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.INPUTMODE(v)
}

func (e *SubHTMLElement) RemoveINPUTMODE(v string) *SubHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *SubHTMLElement) IS(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *SubHTMLElement) IfIS(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.IS(v)
}

func (e *SubHTMLElement) RemoveIS(v string) *SubHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *SubHTMLElement) ITEMID(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *SubHTMLElement) IfITEMID(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMID(v)
}

func (e *SubHTMLElement) RemoveITEMID(v string) *SubHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *SubHTMLElement) ITEMPROP(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *SubHTMLElement) IfITEMPROP(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMPROP(v)
}

func (e *SubHTMLElement) RemoveITEMPROP(v string) *SubHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *SubHTMLElement) ITEMREF(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *SubHTMLElement) IfITEMREF(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMREF(v)
}

func (e *SubHTMLElement) RemoveITEMREF(v string) *SubHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *SubHTMLElement) ITEMSCOPE() *SubHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *SubHTMLElement) IfITEMSCOPE(cond bool) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMSCOPE()
}

func (e *SubHTMLElement) RemoveITEMSCOPE() *SubHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *SubHTMLElement) SetITEMSCOPE(b bool) *SubHTMLElement {
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
func (e *SubHTMLElement) ITEMTYPE(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *SubHTMLElement) IfITEMTYPE(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMTYPE(v)
}

func (e *SubHTMLElement) RemoveITEMTYPE(v string) *SubHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SubHTMLElement) LANG(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *SubHTMLElement) IfLANG(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.LANG(v)
}

func (e *SubHTMLElement) RemoveLANG(v string) *SubHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *SubHTMLElement) NONCE(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *SubHTMLElement) IfNONCE(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.NONCE(v)
}

func (e *SubHTMLElement) RemoveNONCE(v string) *SubHTMLElement {
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
func (e *SubHTMLElement) POPOVER(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *SubHTMLElement) IfPOPOVER(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.POPOVER(v)
}

func (e *SubHTMLElement) RemovePOPOVER(v string) *SubHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *SubHTMLElement) SLOT(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *SubHTMLElement) IfSLOT(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.SLOT(v)
}

func (e *SubHTMLElement) RemoveSLOT(v string) *SubHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *SubHTMLElement) SPELLCHECK(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *SubHTMLElement) IfSPELLCHECK(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.SPELLCHECK(v)
}

func (e *SubHTMLElement) RemoveSPELLCHECK(v string) *SubHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SubHTMLElement) STYLE(k,v string) *SubHTMLElement {
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

func (e *SubHTMLElement) IfSTYLE(cond bool, k string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.STYLE(k, "")
}

func (e *SubHTMLElement) RemoveSTYLE(k string) *SubHTMLElement {
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
func (e *SubHTMLElement) TABINDEX(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *SubHTMLElement) IfTABINDEX(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.TABINDEX(v)
}

func (e *SubHTMLElement) RemoveTABINDEX(v string) *SubHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *SubHTMLElement) TITLE(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *SubHTMLElement) IfTITLE(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.TITLE(v)
}

func (e *SubHTMLElement) RemoveTITLE(v string) *SubHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *SubHTMLElement) TRANSLATE(v string) *SubHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *SubHTMLElement) IfTRANSLATE(cond bool, v string) *SubHTMLElement {
    if !cond {
        return e
    }
    return e.TRANSLATE(v)
}

func (e *SubHTMLElement) RemoveTRANSLATE(v string) *SubHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
