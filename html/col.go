package  html

import (
    "fmt"
)

type ColHTMLElement struct {
    *Element
}

func COL(children ...fmt.Stringer) *ColHTMLElement {
    return &ColHTMLElement{
        Element: &Element{
            Tag: "col",
            IsSelfClosing: true,
            Descendants: children,
        },
    }
}

func (e *ColHTMLElement) Children(children ...fmt.Stringer) *ColHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *ColHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *ColHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *ColHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *ColHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *ColHTMLElement) Text(text string) *ColHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *ColHTMLElement) TextF(format string, args ...any) *ColHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *ColHTMLElement) Raw(text string) *ColHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *ColHTMLElement) RawF(format string, args ...any) *ColHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *ColHTMLElement) CustomData(key, value string) *ColHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ColHTMLElement) CustomDataRemove(key string) *ColHTMLElement {
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
func (e *ColHTMLElement) ACCESSKEY(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *ColHTMLElement) RemoveACCESSKEY(v string) *ColHTMLElement {
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
func (e *ColHTMLElement) AUTOCAPITALIZE(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *ColHTMLElement) RemoveAUTOCAPITALIZE(v string) *ColHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *ColHTMLElement) AUTOFOCUS() *ColHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *ColHTMLElement) RemoveAUTOFOCUS() *ColHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *ColHTMLElement) SetAUTOFOCUS(b bool) *ColHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *ColHTMLElement) CLASS(v string) *ColHTMLElement {
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

func (e *ColHTMLElement) SetCLASS(v string) *ColHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *ColHTMLElement) RemoveCLASS(v string) *ColHTMLElement {
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
func (e *ColHTMLElement) CONTENTEDITABLE(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *ColHTMLElement) RemoveCONTENTEDITABLE(v string) *ColHTMLElement {
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
func (e *ColHTMLElement) DIR(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *ColHTMLElement) RemoveDIR(v string) *ColHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *ColHTMLElement) DRAGGABLE(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *ColHTMLElement) RemoveDRAGGABLE(v string) *ColHTMLElement {
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
func (e *ColHTMLElement) ENTERKEYHINT(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *ColHTMLElement) RemoveENTERKEYHINT(v string) *ColHTMLElement {
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
func (e *ColHTMLElement) HIDDEN(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *ColHTMLElement) RemoveHIDDEN(v string) *ColHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *ColHTMLElement) ID(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *ColHTMLElement) RemoveID(v string) *ColHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *ColHTMLElement) INERT() *ColHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *ColHTMLElement) RemoveINERT() *ColHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *ColHTMLElement) SetINERT(b bool) *ColHTMLElement {
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
func (e *ColHTMLElement) INPUTMODE(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *ColHTMLElement) RemoveINPUTMODE(v string) *ColHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *ColHTMLElement) IS(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *ColHTMLElement) RemoveIS(v string) *ColHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *ColHTMLElement) ITEMID(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *ColHTMLElement) RemoveITEMID(v string) *ColHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *ColHTMLElement) ITEMPROP(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *ColHTMLElement) RemoveITEMPROP(v string) *ColHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ColHTMLElement) ITEMREF(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *ColHTMLElement) RemoveITEMREF(v string) *ColHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *ColHTMLElement) ITEMSCOPE() *ColHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *ColHTMLElement) RemoveITEMSCOPE() *ColHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *ColHTMLElement) SetITEMSCOPE(b bool) *ColHTMLElement {
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
func (e *ColHTMLElement) ITEMTYPE(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *ColHTMLElement) RemoveITEMTYPE(v string) *ColHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ColHTMLElement) LANG(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *ColHTMLElement) RemoveLANG(v string) *ColHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *ColHTMLElement) NONCE(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *ColHTMLElement) RemoveNONCE(v string) *ColHTMLElement {
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
func (e *ColHTMLElement) POPOVER(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *ColHTMLElement) RemovePOPOVER(v string) *ColHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *ColHTMLElement) SLOT(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *ColHTMLElement) RemoveSLOT(v string) *ColHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPAN sets the "span" attribute.
// Number of columns spanned by the element
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *ColHTMLElement) SPAN(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["span"] = v
    return e
}

func (e *ColHTMLElement) RemoveSPAN(v string) *ColHTMLElement {
    delete(e.StringAttributes, "span")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *ColHTMLElement) SPELLCHECK(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *ColHTMLElement) RemoveSPELLCHECK(v string) *ColHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ColHTMLElement) STYLE(k,v string) *ColHTMLElement {
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

func (e *ColHTMLElement) RemoveSTYLE(k string) *ColHTMLElement {
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
func (e *ColHTMLElement) TABINDEX(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *ColHTMLElement) RemoveTABINDEX(v string) *ColHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *ColHTMLElement) TITLE(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *ColHTMLElement) RemoveTITLE(v string) *ColHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *ColHTMLElement) TRANSLATE(v string) *ColHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *ColHTMLElement) RemoveTRANSLATE(v string) *ColHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
