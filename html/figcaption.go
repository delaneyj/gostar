package  html

import (
    "fmt"
)

type FigcaptionHTMLElement struct {
    *Element
}

func FIGCAPTION(children ...fmt.Stringer) *FigcaptionHTMLElement {
    return &FigcaptionHTMLElement{
        Element: &Element{
            Tag: "figcaption",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *FigcaptionHTMLElement) Children(children ...fmt.Stringer) *FigcaptionHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *FigcaptionHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *FigcaptionHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *FigcaptionHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *FigcaptionHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *FigcaptionHTMLElement) Text(text string) *FigcaptionHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *FigcaptionHTMLElement) TextF(format string, args ...any) *FigcaptionHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *FigcaptionHTMLElement) Raw(text string) *FigcaptionHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *FigcaptionHTMLElement) RawF(format string, args ...any) *FigcaptionHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *FigcaptionHTMLElement) CustomData(key, value string) *FigcaptionHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *FigcaptionHTMLElement) CustomDataRemove(key string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) ACCESSKEY(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveACCESSKEY(v string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) AUTOCAPITALIZE(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveAUTOCAPITALIZE(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *FigcaptionHTMLElement) AUTOFOCUS() *FigcaptionHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *FigcaptionHTMLElement) RemoveAUTOFOCUS() *FigcaptionHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *FigcaptionHTMLElement) SetAUTOFOCUS(b bool) *FigcaptionHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *FigcaptionHTMLElement) CLASS(v string) *FigcaptionHTMLElement {
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

func (e *FigcaptionHTMLElement) SetCLASS(v string) *FigcaptionHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *FigcaptionHTMLElement) RemoveCLASS(v string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) CONTENTEDITABLE(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveCONTENTEDITABLE(v string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) DIR(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveDIR(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *FigcaptionHTMLElement) DRAGGABLE(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveDRAGGABLE(v string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) ENTERKEYHINT(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveENTERKEYHINT(v string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) HIDDEN(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveHIDDEN(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *FigcaptionHTMLElement) ID(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveID(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *FigcaptionHTMLElement) INERT() *FigcaptionHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *FigcaptionHTMLElement) RemoveINERT() *FigcaptionHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *FigcaptionHTMLElement) SetINERT(b bool) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) INPUTMODE(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveINPUTMODE(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *FigcaptionHTMLElement) IS(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveIS(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *FigcaptionHTMLElement) ITEMID(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveITEMID(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *FigcaptionHTMLElement) ITEMPROP(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveITEMPROP(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *FigcaptionHTMLElement) ITEMREF(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveITEMREF(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *FigcaptionHTMLElement) ITEMSCOPE() *FigcaptionHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *FigcaptionHTMLElement) RemoveITEMSCOPE() *FigcaptionHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *FigcaptionHTMLElement) SetITEMSCOPE(b bool) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) ITEMTYPE(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveITEMTYPE(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *FigcaptionHTMLElement) LANG(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveLANG(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *FigcaptionHTMLElement) NONCE(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveNONCE(v string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) POPOVER(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemovePOPOVER(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *FigcaptionHTMLElement) SLOT(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveSLOT(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *FigcaptionHTMLElement) SPELLCHECK(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveSPELLCHECK(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *FigcaptionHTMLElement) STYLE(k,v string) *FigcaptionHTMLElement {
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

func (e *FigcaptionHTMLElement) RemoveSTYLE(k string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) TABINDEX(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveTABINDEX(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *FigcaptionHTMLElement) TITLE(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveTITLE(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *FigcaptionHTMLElement) TRANSLATE(v string) *FigcaptionHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *FigcaptionHTMLElement) RemoveTRANSLATE(v string) *FigcaptionHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
