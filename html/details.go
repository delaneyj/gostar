package  html

import (
    "fmt"
)

type DetailsHTMLElement struct {
    *Element
}

func DETAILS(children ...ElementBuilder) *DetailsHTMLElement {
    return &DetailsHTMLElement{
        Element: &Element{
            Tag: "details",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *DetailsHTMLElement) Children(children ...ElementBuilder) *DetailsHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *DetailsHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DetailsHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *DetailsHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DetailsHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *DetailsHTMLElement) Text(text string) *DetailsHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *DetailsHTMLElement) TextF(format string, args ...any) *DetailsHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *DetailsHTMLElement) Raw(text string) *DetailsHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *DetailsHTMLElement) RawF(format string, args ...any) *DetailsHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *DetailsHTMLElement) CustomData(key, value string) *DetailsHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DetailsHTMLElement) CustomDataRemove(key string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) ACCESSKEY(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveACCESSKEY(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) AUTOCAPITALIZE(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveAUTOCAPITALIZE(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *DetailsHTMLElement) AUTOFOCUS() *DetailsHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *DetailsHTMLElement) RemoveAUTOFOCUS() *DetailsHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *DetailsHTMLElement) SetAUTOFOCUS(b bool) *DetailsHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *DetailsHTMLElement) CLASS(v string) *DetailsHTMLElement {
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

func (e *DetailsHTMLElement) SetCLASS(v string) *DetailsHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *DetailsHTMLElement) RemoveCLASS(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) CONTENTEDITABLE(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveCONTENTEDITABLE(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) DIR(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveDIR(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *DetailsHTMLElement) DRAGGABLE(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveDRAGGABLE(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) ENTERKEYHINT(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveENTERKEYHINT(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) HIDDEN(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveHIDDEN(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *DetailsHTMLElement) ID(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveID(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *DetailsHTMLElement) INERT() *DetailsHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *DetailsHTMLElement) RemoveINERT() *DetailsHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *DetailsHTMLElement) SetINERT(b bool) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) INPUTMODE(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveINPUTMODE(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *DetailsHTMLElement) IS(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveIS(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *DetailsHTMLElement) ITEMID(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveITEMID(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *DetailsHTMLElement) ITEMPROP(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveITEMPROP(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *DetailsHTMLElement) ITEMREF(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveITEMREF(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *DetailsHTMLElement) ITEMSCOPE() *DetailsHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *DetailsHTMLElement) RemoveITEMSCOPE() *DetailsHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *DetailsHTMLElement) SetITEMSCOPE(b bool) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) ITEMTYPE(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveITEMTYPE(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DetailsHTMLElement) LANG(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveLANG(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *DetailsHTMLElement) NAME(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveNAME(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *DetailsHTMLElement) NONCE(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveNONCE(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// OPEN sets the "open" attribute.
// Whether the dialog box is showing
// Values values are constrained to:
//  * boolean_attribute
func (e *DetailsHTMLElement) OPEN() *DetailsHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["open"] = struct{}{}
    return e
}

func (e *DetailsHTMLElement) RemoveOPEN() *DetailsHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "open")
    return e
}

func (e *DetailsHTMLElement) SetOPEN(b bool) *DetailsHTMLElement {
    if b {
        return e.OPEN()
    }
    return e.RemoveOPEN()
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *DetailsHTMLElement) POPOVER(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *DetailsHTMLElement) RemovePOPOVER(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *DetailsHTMLElement) SLOT(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveSLOT(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *DetailsHTMLElement) SPELLCHECK(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveSPELLCHECK(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DetailsHTMLElement) STYLE(k,v string) *DetailsHTMLElement {
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

func (e *DetailsHTMLElement) RemoveSTYLE(k string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) TABINDEX(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveTABINDEX(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *DetailsHTMLElement) TITLE(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveTITLE(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *DetailsHTMLElement) TRANSLATE(v string) *DetailsHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *DetailsHTMLElement) RemoveTRANSLATE(v string) *DetailsHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
