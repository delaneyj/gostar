package  html

import (
    "fmt"
)

type TitleHTMLElement struct {
    *Element
}

func TITLE(children ...ElementBuilder) *TitleHTMLElement {
    return &TitleHTMLElement{
        Element: &Element{
            Tag: "title",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *TitleHTMLElement) Children(children ...ElementBuilder) *TitleHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *TitleHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TitleHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *TitleHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TitleHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *TitleHTMLElement) Text(text string) *TitleHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *TitleHTMLElement) TextF(format string, args ...any) *TitleHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *TitleHTMLElement) Escaped(text string) *TitleHTMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *TitleHTMLElement) EscapedF(format string, args ...any) *TitleHTMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TitleHTMLElement) CustomData(key, value string) *TitleHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TitleHTMLElement) CustomDataRemove(key string) *TitleHTMLElement {
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
func (e *TitleHTMLElement) ACCESSKEY(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *TitleHTMLElement) IfACCESSKEY(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.ACCESSKEY(v)
}

func (e *TitleHTMLElement) RemoveACCESSKEY(v string) *TitleHTMLElement {
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
func (e *TitleHTMLElement) AUTOCAPITALIZE(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *TitleHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCAPITALIZE(v)
}

func (e *TitleHTMLElement) RemoveAUTOCAPITALIZE(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *TitleHTMLElement) AUTOFOCUS() *TitleHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *TitleHTMLElement) IfAUTOFOCUS(cond bool) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOFOCUS()
}

func (e *TitleHTMLElement) RemoveAUTOFOCUS() *TitleHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *TitleHTMLElement) SetAUTOFOCUS(b bool) *TitleHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *TitleHTMLElement) CLASS(v string) *TitleHTMLElement {
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

func (e *TitleHTMLElement) IfCLASS(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.CLASS(v)
}

func (e *TitleHTMLElement) SetCLASS(v string) *TitleHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *TitleHTMLElement) RemoveCLASS(v string) *TitleHTMLElement {
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
func (e *TitleHTMLElement) CONTENTEDITABLE(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *TitleHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.CONTENTEDITABLE(v)
}

func (e *TitleHTMLElement) RemoveCONTENTEDITABLE(v string) *TitleHTMLElement {
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
func (e *TitleHTMLElement) DIR(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *TitleHTMLElement) IfDIR(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.DIR(v)
}

func (e *TitleHTMLElement) RemoveDIR(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *TitleHTMLElement) DRAGGABLE(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *TitleHTMLElement) IfDRAGGABLE(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.DRAGGABLE(v)
}

func (e *TitleHTMLElement) RemoveDRAGGABLE(v string) *TitleHTMLElement {
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
func (e *TitleHTMLElement) ENTERKEYHINT(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *TitleHTMLElement) IfENTERKEYHINT(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.ENTERKEYHINT(v)
}

func (e *TitleHTMLElement) RemoveENTERKEYHINT(v string) *TitleHTMLElement {
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
func (e *TitleHTMLElement) HIDDEN(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *TitleHTMLElement) IfHIDDEN(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.HIDDEN(v)
}

func (e *TitleHTMLElement) RemoveHIDDEN(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *TitleHTMLElement) ID(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *TitleHTMLElement) IfID(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.ID(v)
}

func (e *TitleHTMLElement) RemoveID(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *TitleHTMLElement) INERT() *TitleHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *TitleHTMLElement) IfINERT(cond bool) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.INERT()
}

func (e *TitleHTMLElement) RemoveINERT() *TitleHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *TitleHTMLElement) SetINERT(b bool) *TitleHTMLElement {
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
func (e *TitleHTMLElement) INPUTMODE(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *TitleHTMLElement) IfINPUTMODE(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.INPUTMODE(v)
}

func (e *TitleHTMLElement) RemoveINPUTMODE(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *TitleHTMLElement) IS(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *TitleHTMLElement) IfIS(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.IS(v)
}

func (e *TitleHTMLElement) RemoveIS(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *TitleHTMLElement) ITEMID(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *TitleHTMLElement) IfITEMID(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMID(v)
}

func (e *TitleHTMLElement) RemoveITEMID(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *TitleHTMLElement) ITEMPROP(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *TitleHTMLElement) IfITEMPROP(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMPROP(v)
}

func (e *TitleHTMLElement) RemoveITEMPROP(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *TitleHTMLElement) ITEMREF(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *TitleHTMLElement) IfITEMREF(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMREF(v)
}

func (e *TitleHTMLElement) RemoveITEMREF(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *TitleHTMLElement) ITEMSCOPE() *TitleHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *TitleHTMLElement) IfITEMSCOPE(cond bool) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMSCOPE()
}

func (e *TitleHTMLElement) RemoveITEMSCOPE() *TitleHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *TitleHTMLElement) SetITEMSCOPE(b bool) *TitleHTMLElement {
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
func (e *TitleHTMLElement) ITEMTYPE(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *TitleHTMLElement) IfITEMTYPE(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMTYPE(v)
}

func (e *TitleHTMLElement) RemoveITEMTYPE(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TitleHTMLElement) LANG(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *TitleHTMLElement) IfLANG(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.LANG(v)
}

func (e *TitleHTMLElement) RemoveLANG(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *TitleHTMLElement) NONCE(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *TitleHTMLElement) IfNONCE(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.NONCE(v)
}

func (e *TitleHTMLElement) RemoveNONCE(v string) *TitleHTMLElement {
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
func (e *TitleHTMLElement) POPOVER(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *TitleHTMLElement) IfPOPOVER(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.POPOVER(v)
}

func (e *TitleHTMLElement) RemovePOPOVER(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *TitleHTMLElement) SLOT(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *TitleHTMLElement) IfSLOT(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.SLOT(v)
}

func (e *TitleHTMLElement) RemoveSLOT(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *TitleHTMLElement) SPELLCHECK(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *TitleHTMLElement) IfSPELLCHECK(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.SPELLCHECK(v)
}

func (e *TitleHTMLElement) RemoveSPELLCHECK(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TitleHTMLElement) STYLE(k,v string) *TitleHTMLElement {
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

func (e *TitleHTMLElement) IfSTYLE(cond bool, k string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.STYLE(k, "")
}

func (e *TitleHTMLElement) RemoveSTYLE(k string) *TitleHTMLElement {
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
func (e *TitleHTMLElement) TABINDEX(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *TitleHTMLElement) IfTABINDEX(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.TABINDEX(v)
}

func (e *TitleHTMLElement) RemoveTABINDEX(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *TitleHTMLElement) TITLE(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *TitleHTMLElement) IfTITLE(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.TITLE(v)
}

func (e *TitleHTMLElement) RemoveTITLE(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *TitleHTMLElement) TRANSLATE(v string) *TitleHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *TitleHTMLElement) IfTRANSLATE(cond bool, v string) *TitleHTMLElement {
    if !cond {
        return e
    }
    return e.TRANSLATE(v)
}

func (e *TitleHTMLElement) RemoveTRANSLATE(v string) *TitleHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
