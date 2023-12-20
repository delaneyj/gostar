package  html

import (
    "fmt"
)

type LegendHTMLElement struct {
    *Element
}

func LEGEND(children ...ElementBuilder) *LegendHTMLElement {
    return &LegendHTMLElement{
        Element: &Element{
            Tag: "legend",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *LegendHTMLElement) Children(children ...ElementBuilder) *LegendHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *LegendHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *LegendHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *LegendHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *LegendHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *LegendHTMLElement) Text(text string) *LegendHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *LegendHTMLElement) TextF(format string, args ...any) *LegendHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *LegendHTMLElement) Escaped(text string) *LegendHTMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *LegendHTMLElement) EscapedF(format string, args ...any) *LegendHTMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *LegendHTMLElement) CustomData(key, value string) *LegendHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *LegendHTMLElement) CustomDataRemove(key string) *LegendHTMLElement {
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
func (e *LegendHTMLElement) ACCESSKEY(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *LegendHTMLElement) IfACCESSKEY(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.ACCESSKEY(v)
}

func (e *LegendHTMLElement) RemoveACCESSKEY(v string) *LegendHTMLElement {
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
func (e *LegendHTMLElement) AUTOCAPITALIZE(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *LegendHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCAPITALIZE(v)
}

func (e *LegendHTMLElement) RemoveAUTOCAPITALIZE(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *LegendHTMLElement) AUTOFOCUS() *LegendHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *LegendHTMLElement) IfAUTOFOCUS(cond bool) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOFOCUS()
}

func (e *LegendHTMLElement) RemoveAUTOFOCUS() *LegendHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *LegendHTMLElement) SetAUTOFOCUS(b bool) *LegendHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *LegendHTMLElement) CLASS(v string) *LegendHTMLElement {
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

func (e *LegendHTMLElement) IfCLASS(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.CLASS(v)
}

func (e *LegendHTMLElement) SetCLASS(v string) *LegendHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *LegendHTMLElement) RemoveCLASS(v string) *LegendHTMLElement {
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
func (e *LegendHTMLElement) CONTENTEDITABLE(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *LegendHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.CONTENTEDITABLE(v)
}

func (e *LegendHTMLElement) RemoveCONTENTEDITABLE(v string) *LegendHTMLElement {
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
func (e *LegendHTMLElement) DIR(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *LegendHTMLElement) IfDIR(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.DIR(v)
}

func (e *LegendHTMLElement) RemoveDIR(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *LegendHTMLElement) DRAGGABLE(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *LegendHTMLElement) IfDRAGGABLE(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.DRAGGABLE(v)
}

func (e *LegendHTMLElement) RemoveDRAGGABLE(v string) *LegendHTMLElement {
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
func (e *LegendHTMLElement) ENTERKEYHINT(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *LegendHTMLElement) IfENTERKEYHINT(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.ENTERKEYHINT(v)
}

func (e *LegendHTMLElement) RemoveENTERKEYHINT(v string) *LegendHTMLElement {
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
func (e *LegendHTMLElement) HIDDEN(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *LegendHTMLElement) IfHIDDEN(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.HIDDEN(v)
}

func (e *LegendHTMLElement) RemoveHIDDEN(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *LegendHTMLElement) ID(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *LegendHTMLElement) IfID(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.ID(v)
}

func (e *LegendHTMLElement) RemoveID(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *LegendHTMLElement) INERT() *LegendHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *LegendHTMLElement) IfINERT(cond bool) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.INERT()
}

func (e *LegendHTMLElement) RemoveINERT() *LegendHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *LegendHTMLElement) SetINERT(b bool) *LegendHTMLElement {
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
func (e *LegendHTMLElement) INPUTMODE(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *LegendHTMLElement) IfINPUTMODE(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.INPUTMODE(v)
}

func (e *LegendHTMLElement) RemoveINPUTMODE(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *LegendHTMLElement) IS(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *LegendHTMLElement) IfIS(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.IS(v)
}

func (e *LegendHTMLElement) RemoveIS(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *LegendHTMLElement) ITEMID(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *LegendHTMLElement) IfITEMID(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMID(v)
}

func (e *LegendHTMLElement) RemoveITEMID(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *LegendHTMLElement) ITEMPROP(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *LegendHTMLElement) IfITEMPROP(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMPROP(v)
}

func (e *LegendHTMLElement) RemoveITEMPROP(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *LegendHTMLElement) ITEMREF(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *LegendHTMLElement) IfITEMREF(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMREF(v)
}

func (e *LegendHTMLElement) RemoveITEMREF(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *LegendHTMLElement) ITEMSCOPE() *LegendHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *LegendHTMLElement) IfITEMSCOPE(cond bool) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMSCOPE()
}

func (e *LegendHTMLElement) RemoveITEMSCOPE() *LegendHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *LegendHTMLElement) SetITEMSCOPE(b bool) *LegendHTMLElement {
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
func (e *LegendHTMLElement) ITEMTYPE(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *LegendHTMLElement) IfITEMTYPE(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMTYPE(v)
}

func (e *LegendHTMLElement) RemoveITEMTYPE(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *LegendHTMLElement) LANG(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *LegendHTMLElement) IfLANG(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.LANG(v)
}

func (e *LegendHTMLElement) RemoveLANG(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *LegendHTMLElement) NONCE(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *LegendHTMLElement) IfNONCE(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.NONCE(v)
}

func (e *LegendHTMLElement) RemoveNONCE(v string) *LegendHTMLElement {
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
func (e *LegendHTMLElement) POPOVER(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *LegendHTMLElement) IfPOPOVER(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.POPOVER(v)
}

func (e *LegendHTMLElement) RemovePOPOVER(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *LegendHTMLElement) SLOT(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *LegendHTMLElement) IfSLOT(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.SLOT(v)
}

func (e *LegendHTMLElement) RemoveSLOT(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *LegendHTMLElement) SPELLCHECK(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *LegendHTMLElement) IfSPELLCHECK(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.SPELLCHECK(v)
}

func (e *LegendHTMLElement) RemoveSPELLCHECK(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *LegendHTMLElement) STYLE(k,v string) *LegendHTMLElement {
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

func (e *LegendHTMLElement) IfSTYLE(cond bool, k string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.STYLE(k, "")
}

func (e *LegendHTMLElement) RemoveSTYLE(k string) *LegendHTMLElement {
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
func (e *LegendHTMLElement) TABINDEX(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *LegendHTMLElement) IfTABINDEX(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.TABINDEX(v)
}

func (e *LegendHTMLElement) RemoveTABINDEX(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *LegendHTMLElement) TITLE(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *LegendHTMLElement) IfTITLE(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.TITLE(v)
}

func (e *LegendHTMLElement) RemoveTITLE(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *LegendHTMLElement) TRANSLATE(v string) *LegendHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *LegendHTMLElement) IfTRANSLATE(cond bool, v string) *LegendHTMLElement {
    if !cond {
        return e
    }
    return e.TRANSLATE(v)
}

func (e *LegendHTMLElement) RemoveTRANSLATE(v string) *LegendHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
