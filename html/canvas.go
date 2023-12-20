package  html

import (
    "fmt"
)

type CanvasHTMLElement struct {
    *Element
}

func CANVAS(children ...fmt.Stringer) *CanvasHTMLElement {
    return &CanvasHTMLElement{
        Element: &Element{
            Tag: "canvas",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *CanvasHTMLElement) Children(children ...fmt.Stringer) *CanvasHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *CanvasHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *CanvasHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *CanvasHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *CanvasHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *CanvasHTMLElement) Text(text string) *CanvasHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *CanvasHTMLElement) TextF(format string, args ...any) *CanvasHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *CanvasHTMLElement) Raw(text string) *CanvasHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *CanvasHTMLElement) RawF(format string, args ...any) *CanvasHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *CanvasHTMLElement) CustomData(key, value string) *CanvasHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *CanvasHTMLElement) CustomDataRemove(key string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) ACCESSKEY(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveACCESSKEY(v string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) AUTOCAPITALIZE(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveAUTOCAPITALIZE(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *CanvasHTMLElement) AUTOFOCUS() *CanvasHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *CanvasHTMLElement) RemoveAUTOFOCUS() *CanvasHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *CanvasHTMLElement) SetAUTOFOCUS(b bool) *CanvasHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *CanvasHTMLElement) CLASS(v string) *CanvasHTMLElement {
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

func (e *CanvasHTMLElement) SetCLASS(v string) *CanvasHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *CanvasHTMLElement) RemoveCLASS(v string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) CONTENTEDITABLE(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveCONTENTEDITABLE(v string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) DIR(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveDIR(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *CanvasHTMLElement) DRAGGABLE(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveDRAGGABLE(v string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) ENTERKEYHINT(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveENTERKEYHINT(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *CanvasHTMLElement) HEIGHT(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveHEIGHT(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "height")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *CanvasHTMLElement) HIDDEN(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveHIDDEN(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *CanvasHTMLElement) ID(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveID(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *CanvasHTMLElement) INERT() *CanvasHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *CanvasHTMLElement) RemoveINERT() *CanvasHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *CanvasHTMLElement) SetINERT(b bool) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) INPUTMODE(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveINPUTMODE(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *CanvasHTMLElement) IS(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveIS(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *CanvasHTMLElement) ITEMID(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveITEMID(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *CanvasHTMLElement) ITEMPROP(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveITEMPROP(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *CanvasHTMLElement) ITEMREF(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveITEMREF(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *CanvasHTMLElement) ITEMSCOPE() *CanvasHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *CanvasHTMLElement) RemoveITEMSCOPE() *CanvasHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *CanvasHTMLElement) SetITEMSCOPE(b bool) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) ITEMTYPE(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveITEMTYPE(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *CanvasHTMLElement) LANG(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveLANG(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *CanvasHTMLElement) NONCE(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveNONCE(v string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) POPOVER(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *CanvasHTMLElement) RemovePOPOVER(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *CanvasHTMLElement) SLOT(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveSLOT(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *CanvasHTMLElement) SPELLCHECK(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveSPELLCHECK(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *CanvasHTMLElement) STYLE(k,v string) *CanvasHTMLElement {
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

func (e *CanvasHTMLElement) RemoveSTYLE(k string) *CanvasHTMLElement {
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
func (e *CanvasHTMLElement) TABINDEX(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveTABINDEX(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *CanvasHTMLElement) TITLE(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveTITLE(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *CanvasHTMLElement) TRANSLATE(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveTRANSLATE(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *CanvasHTMLElement) WIDTH(v string) *CanvasHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *CanvasHTMLElement) RemoveWIDTH(v string) *CanvasHTMLElement {
    delete(e.StringAttributes, "width")
    return e
}
