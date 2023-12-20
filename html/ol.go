package  html

import (
    "fmt"
)

type OlHTMLElement struct {
    *Element
}

func OL(children ...fmt.Stringer) *OlHTMLElement {
    return &OlHTMLElement{
        Element: &Element{
            Tag: "ol",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *OlHTMLElement) Children(children ...fmt.Stringer) *OlHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *OlHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *OlHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *OlHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *OlHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *OlHTMLElement) Text(text string) *OlHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *OlHTMLElement) TextF(format string, args ...any) *OlHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *OlHTMLElement) Raw(text string) *OlHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *OlHTMLElement) RawF(format string, args ...any) *OlHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *OlHTMLElement) CustomData(key, value string) *OlHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *OlHTMLElement) CustomDataRemove(key string) *OlHTMLElement {
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
func (e *OlHTMLElement) ACCESSKEY(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *OlHTMLElement) RemoveACCESSKEY(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) AUTOCAPITALIZE(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *OlHTMLElement) RemoveAUTOCAPITALIZE(v string) *OlHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *OlHTMLElement) AUTOFOCUS() *OlHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *OlHTMLElement) RemoveAUTOFOCUS() *OlHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *OlHTMLElement) SetAUTOFOCUS(b bool) *OlHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *OlHTMLElement) CLASS(v string) *OlHTMLElement {
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

func (e *OlHTMLElement) SetCLASS(v string) *OlHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *OlHTMLElement) RemoveCLASS(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) CONTENTEDITABLE(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *OlHTMLElement) RemoveCONTENTEDITABLE(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) DIR(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *OlHTMLElement) RemoveDIR(v string) *OlHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *OlHTMLElement) DRAGGABLE(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *OlHTMLElement) RemoveDRAGGABLE(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) ENTERKEYHINT(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *OlHTMLElement) RemoveENTERKEYHINT(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) HIDDEN(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *OlHTMLElement) RemoveHIDDEN(v string) *OlHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *OlHTMLElement) ID(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *OlHTMLElement) RemoveID(v string) *OlHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *OlHTMLElement) INERT() *OlHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *OlHTMLElement) RemoveINERT() *OlHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *OlHTMLElement) SetINERT(b bool) *OlHTMLElement {
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
func (e *OlHTMLElement) INPUTMODE(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *OlHTMLElement) RemoveINPUTMODE(v string) *OlHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *OlHTMLElement) IS(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *OlHTMLElement) RemoveIS(v string) *OlHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *OlHTMLElement) ITEMID(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *OlHTMLElement) RemoveITEMID(v string) *OlHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *OlHTMLElement) ITEMPROP(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *OlHTMLElement) RemoveITEMPROP(v string) *OlHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *OlHTMLElement) ITEMREF(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *OlHTMLElement) RemoveITEMREF(v string) *OlHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *OlHTMLElement) ITEMSCOPE() *OlHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *OlHTMLElement) RemoveITEMSCOPE() *OlHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *OlHTMLElement) SetITEMSCOPE(b bool) *OlHTMLElement {
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
func (e *OlHTMLElement) ITEMTYPE(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *OlHTMLElement) RemoveITEMTYPE(v string) *OlHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *OlHTMLElement) LANG(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *OlHTMLElement) RemoveLANG(v string) *OlHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *OlHTMLElement) NONCE(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *OlHTMLElement) RemoveNONCE(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) POPOVER(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *OlHTMLElement) RemovePOPOVER(v string) *OlHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REVERSED sets the "reversed" attribute.
// Number the list backwards
// Values values are constrained to:
//  * boolean_attribute
func (e *OlHTMLElement) REVERSED() *OlHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["reversed"] = struct{}{}
    return e
}

func (e *OlHTMLElement) RemoveREVERSED() *OlHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "reversed")
    return e
}

func (e *OlHTMLElement) SetREVERSED(b bool) *OlHTMLElement {
    if b {
        return e.REVERSED()
    }
    return e.RemoveREVERSED()
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *OlHTMLElement) SLOT(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *OlHTMLElement) RemoveSLOT(v string) *OlHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *OlHTMLElement) SPELLCHECK(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *OlHTMLElement) RemoveSPELLCHECK(v string) *OlHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// START sets the "start" attribute.
// Starting value of the list
// Values values are constrained to:
//  * valid_integer
func (e *OlHTMLElement) START(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["start"] = v
    return e
}

func (e *OlHTMLElement) RemoveSTART(v string) *OlHTMLElement {
    delete(e.StringAttributes, "start")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *OlHTMLElement) STYLE(k,v string) *OlHTMLElement {
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

func (e *OlHTMLElement) RemoveSTYLE(k string) *OlHTMLElement {
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
func (e *OlHTMLElement) TABINDEX(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *OlHTMLElement) RemoveTABINDEX(v string) *OlHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *OlHTMLElement) TITLE(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *OlHTMLElement) RemoveTITLE(v string) *OlHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *OlHTMLElement) TRANSLATE(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *OlHTMLElement) RemoveTRANSLATE(v string) *OlHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *OlHTMLElement) TYPE(v string) *OlHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *OlHTMLElement) RemoveTYPE(v string) *OlHTMLElement {
    delete(e.StringAttributes, "type")
    return e
}
