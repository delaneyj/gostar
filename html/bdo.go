package  html

import (
    "fmt"
)

type BdoHTMLElement struct {
    *Element
}

func BDO(children ...fmt.Stringer) *BdoHTMLElement {
    return &BdoHTMLElement{
        Element: &Element{
            Tag: "bdo",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *BdoHTMLElement) Children(children ...fmt.Stringer) *BdoHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *BdoHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *BdoHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *BdoHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *BdoHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *BdoHTMLElement) Text(text string) *BdoHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *BdoHTMLElement) TextF(format string, args ...any) *BdoHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *BdoHTMLElement) Raw(text string) *BdoHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *BdoHTMLElement) RawF(format string, args ...any) *BdoHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *BdoHTMLElement) CustomData(key, value string) *BdoHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *BdoHTMLElement) CustomDataRemove(key string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) ACCESSKEY(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *BdoHTMLElement) RemoveACCESSKEY(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) AUTOCAPITALIZE(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *BdoHTMLElement) RemoveAUTOCAPITALIZE(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *BdoHTMLElement) AUTOFOCUS() *BdoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *BdoHTMLElement) RemoveAUTOFOCUS() *BdoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *BdoHTMLElement) SetAUTOFOCUS(b bool) *BdoHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *BdoHTMLElement) CLASS(v string) *BdoHTMLElement {
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

func (e *BdoHTMLElement) SetCLASS(v string) *BdoHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *BdoHTMLElement) RemoveCLASS(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) CONTENTEDITABLE(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *BdoHTMLElement) RemoveCONTENTEDITABLE(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) DIR(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *BdoHTMLElement) RemoveDIR(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *BdoHTMLElement) DRAGGABLE(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *BdoHTMLElement) RemoveDRAGGABLE(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) ENTERKEYHINT(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *BdoHTMLElement) RemoveENTERKEYHINT(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) HIDDEN(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *BdoHTMLElement) RemoveHIDDEN(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *BdoHTMLElement) ID(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *BdoHTMLElement) RemoveID(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *BdoHTMLElement) INERT() *BdoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *BdoHTMLElement) RemoveINERT() *BdoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *BdoHTMLElement) SetINERT(b bool) *BdoHTMLElement {
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
func (e *BdoHTMLElement) INPUTMODE(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *BdoHTMLElement) RemoveINPUTMODE(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *BdoHTMLElement) IS(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *BdoHTMLElement) RemoveIS(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *BdoHTMLElement) ITEMID(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *BdoHTMLElement) RemoveITEMID(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *BdoHTMLElement) ITEMPROP(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *BdoHTMLElement) RemoveITEMPROP(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *BdoHTMLElement) ITEMREF(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *BdoHTMLElement) RemoveITEMREF(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *BdoHTMLElement) ITEMSCOPE() *BdoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *BdoHTMLElement) RemoveITEMSCOPE() *BdoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *BdoHTMLElement) SetITEMSCOPE(b bool) *BdoHTMLElement {
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
func (e *BdoHTMLElement) ITEMTYPE(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *BdoHTMLElement) RemoveITEMTYPE(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *BdoHTMLElement) LANG(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *BdoHTMLElement) RemoveLANG(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *BdoHTMLElement) NONCE(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *BdoHTMLElement) RemoveNONCE(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) POPOVER(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *BdoHTMLElement) RemovePOPOVER(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *BdoHTMLElement) SLOT(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *BdoHTMLElement) RemoveSLOT(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *BdoHTMLElement) SPELLCHECK(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *BdoHTMLElement) RemoveSPELLCHECK(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *BdoHTMLElement) STYLE(k,v string) *BdoHTMLElement {
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

func (e *BdoHTMLElement) RemoveSTYLE(k string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) TABINDEX(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *BdoHTMLElement) RemoveTABINDEX(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *BdoHTMLElement) TITLE(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *BdoHTMLElement) RemoveTITLE(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *BdoHTMLElement) TRANSLATE(v string) *BdoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *BdoHTMLElement) RemoveTRANSLATE(v string) *BdoHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
