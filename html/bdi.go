package  html

import (
    "fmt"
)

type BdiHTMLElement struct {
    *Element
}

func BDI(children ...ElementBuilder) *BdiHTMLElement {
    return &BdiHTMLElement{
        Element: &Element{
            Tag: "bdi",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *BdiHTMLElement) Children(children ...ElementBuilder) *BdiHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *BdiHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *BdiHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *BdiHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *BdiHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *BdiHTMLElement) Text(text string) *BdiHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *BdiHTMLElement) TextF(format string, args ...any) *BdiHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *BdiHTMLElement) Escaped(text string) *BdiHTMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *BdiHTMLElement) EscapedF(format string, args ...any) *BdiHTMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *BdiHTMLElement) CustomData(key, value string) *BdiHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *BdiHTMLElement) CustomDataRemove(key string) *BdiHTMLElement {
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
func (e *BdiHTMLElement) ACCESSKEY(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *BdiHTMLElement) IfACCESSKEY(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.ACCESSKEY(v)
}

func (e *BdiHTMLElement) RemoveACCESSKEY(v string) *BdiHTMLElement {
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
func (e *BdiHTMLElement) AUTOCAPITALIZE(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *BdiHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCAPITALIZE(v)
}

func (e *BdiHTMLElement) RemoveAUTOCAPITALIZE(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *BdiHTMLElement) AUTOFOCUS() *BdiHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *BdiHTMLElement) IfAUTOFOCUS(cond bool) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOFOCUS()
}

func (e *BdiHTMLElement) RemoveAUTOFOCUS() *BdiHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *BdiHTMLElement) SetAUTOFOCUS(b bool) *BdiHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *BdiHTMLElement) CLASS(v string) *BdiHTMLElement {
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

func (e *BdiHTMLElement) IfCLASS(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.CLASS(v)
}

func (e *BdiHTMLElement) SetCLASS(v string) *BdiHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *BdiHTMLElement) RemoveCLASS(v string) *BdiHTMLElement {
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
func (e *BdiHTMLElement) CONTENTEDITABLE(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *BdiHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.CONTENTEDITABLE(v)
}

func (e *BdiHTMLElement) RemoveCONTENTEDITABLE(v string) *BdiHTMLElement {
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
func (e *BdiHTMLElement) DIR(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *BdiHTMLElement) IfDIR(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.DIR(v)
}

func (e *BdiHTMLElement) RemoveDIR(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *BdiHTMLElement) DRAGGABLE(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *BdiHTMLElement) IfDRAGGABLE(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.DRAGGABLE(v)
}

func (e *BdiHTMLElement) RemoveDRAGGABLE(v string) *BdiHTMLElement {
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
func (e *BdiHTMLElement) ENTERKEYHINT(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *BdiHTMLElement) IfENTERKEYHINT(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.ENTERKEYHINT(v)
}

func (e *BdiHTMLElement) RemoveENTERKEYHINT(v string) *BdiHTMLElement {
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
func (e *BdiHTMLElement) HIDDEN(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *BdiHTMLElement) IfHIDDEN(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.HIDDEN(v)
}

func (e *BdiHTMLElement) RemoveHIDDEN(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *BdiHTMLElement) ID(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *BdiHTMLElement) IfID(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.ID(v)
}

func (e *BdiHTMLElement) RemoveID(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *BdiHTMLElement) INERT() *BdiHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *BdiHTMLElement) IfINERT(cond bool) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.INERT()
}

func (e *BdiHTMLElement) RemoveINERT() *BdiHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *BdiHTMLElement) SetINERT(b bool) *BdiHTMLElement {
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
func (e *BdiHTMLElement) INPUTMODE(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *BdiHTMLElement) IfINPUTMODE(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.INPUTMODE(v)
}

func (e *BdiHTMLElement) RemoveINPUTMODE(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *BdiHTMLElement) IS(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *BdiHTMLElement) IfIS(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.IS(v)
}

func (e *BdiHTMLElement) RemoveIS(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *BdiHTMLElement) ITEMID(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *BdiHTMLElement) IfITEMID(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMID(v)
}

func (e *BdiHTMLElement) RemoveITEMID(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *BdiHTMLElement) ITEMPROP(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *BdiHTMLElement) IfITEMPROP(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMPROP(v)
}

func (e *BdiHTMLElement) RemoveITEMPROP(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *BdiHTMLElement) ITEMREF(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *BdiHTMLElement) IfITEMREF(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMREF(v)
}

func (e *BdiHTMLElement) RemoveITEMREF(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *BdiHTMLElement) ITEMSCOPE() *BdiHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *BdiHTMLElement) IfITEMSCOPE(cond bool) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMSCOPE()
}

func (e *BdiHTMLElement) RemoveITEMSCOPE() *BdiHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *BdiHTMLElement) SetITEMSCOPE(b bool) *BdiHTMLElement {
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
func (e *BdiHTMLElement) ITEMTYPE(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *BdiHTMLElement) IfITEMTYPE(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMTYPE(v)
}

func (e *BdiHTMLElement) RemoveITEMTYPE(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *BdiHTMLElement) LANG(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *BdiHTMLElement) IfLANG(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.LANG(v)
}

func (e *BdiHTMLElement) RemoveLANG(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *BdiHTMLElement) NONCE(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *BdiHTMLElement) IfNONCE(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.NONCE(v)
}

func (e *BdiHTMLElement) RemoveNONCE(v string) *BdiHTMLElement {
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
func (e *BdiHTMLElement) POPOVER(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *BdiHTMLElement) IfPOPOVER(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.POPOVER(v)
}

func (e *BdiHTMLElement) RemovePOPOVER(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *BdiHTMLElement) SLOT(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *BdiHTMLElement) IfSLOT(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.SLOT(v)
}

func (e *BdiHTMLElement) RemoveSLOT(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *BdiHTMLElement) SPELLCHECK(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *BdiHTMLElement) IfSPELLCHECK(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.SPELLCHECK(v)
}

func (e *BdiHTMLElement) RemoveSPELLCHECK(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *BdiHTMLElement) STYLE(k,v string) *BdiHTMLElement {
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

func (e *BdiHTMLElement) IfSTYLE(cond bool, k string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.STYLE(k, "")
}

func (e *BdiHTMLElement) RemoveSTYLE(k string) *BdiHTMLElement {
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
func (e *BdiHTMLElement) TABINDEX(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *BdiHTMLElement) IfTABINDEX(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.TABINDEX(v)
}

func (e *BdiHTMLElement) RemoveTABINDEX(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *BdiHTMLElement) TITLE(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *BdiHTMLElement) IfTITLE(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.TITLE(v)
}

func (e *BdiHTMLElement) RemoveTITLE(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *BdiHTMLElement) TRANSLATE(v string) *BdiHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *BdiHTMLElement) IfTRANSLATE(cond bool, v string) *BdiHTMLElement {
    if !cond {
        return e
    }
    return e.TRANSLATE(v)
}

func (e *BdiHTMLElement) RemoveTRANSLATE(v string) *BdiHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
