package  html

import (
    "fmt"
)

type TdHTMLElement struct {
    *Element
}

func TD(children ...ElementBuilder) *TdHTMLElement {
    return &TdHTMLElement{
        Element: &Element{
            Tag: "td",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *TdHTMLElement) Children(children ...ElementBuilder) *TdHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *TdHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TdHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *TdHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TdHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *TdHTMLElement) Text(text string) *TdHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *TdHTMLElement) TextF(format string, args ...any) *TdHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *TdHTMLElement) Raw(text string) *TdHTMLElement {
    e.Descendants = append(e.Descendants, Raw(text))
    return e
}

func (e *TdHTMLElement) RawF(format string, args ...any) *TdHTMLElement {
    return e.Raw(fmt.Sprintf(format, args...))
}

func (e *TdHTMLElement) CustomData(key, value string) *TdHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TdHTMLElement) CustomDataRemove(key string) *TdHTMLElement {
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
func (e *TdHTMLElement) ACCESSKEY(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *TdHTMLElement) RemoveACCESSKEY(v string) *TdHTMLElement {
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
func (e *TdHTMLElement) AUTOCAPITALIZE(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *TdHTMLElement) RemoveAUTOCAPITALIZE(v string) *TdHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *TdHTMLElement) AUTOFOCUS() *TdHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *TdHTMLElement) RemoveAUTOFOCUS() *TdHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *TdHTMLElement) SetAUTOFOCUS(b bool) *TdHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *TdHTMLElement) CLASS(v string) *TdHTMLElement {
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

func (e *TdHTMLElement) SetCLASS(v string) *TdHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *TdHTMLElement) RemoveCLASS(v string) *TdHTMLElement {
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        return e
    }
    kv.Remove(v)
    return e
}
// COLSPAN sets the "colspan" attribute.
// Number of columns that the cell is to span
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *TdHTMLElement) COLSPAN(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["colspan"] = v
    return e
}

func (e *TdHTMLElement) RemoveCOLSPAN(v string) *TdHTMLElement {
    delete(e.StringAttributes, "colspan")
    return e
}
// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//  * true
//  * plaintext_only
//  * false
func (e *TdHTMLElement) CONTENTEDITABLE(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *TdHTMLElement) RemoveCONTENTEDITABLE(v string) *TdHTMLElement {
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
func (e *TdHTMLElement) DIR(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *TdHTMLElement) RemoveDIR(v string) *TdHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *TdHTMLElement) DRAGGABLE(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *TdHTMLElement) RemoveDRAGGABLE(v string) *TdHTMLElement {
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
func (e *TdHTMLElement) ENTERKEYHINT(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *TdHTMLElement) RemoveENTERKEYHINT(v string) *TdHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HEADERS sets the "headers" attribute.
// The header cells for this cell
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *TdHTMLElement) HEADERS(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["headers"] = v
    return e
}

func (e *TdHTMLElement) RemoveHEADERS(v string) *TdHTMLElement {
    delete(e.StringAttributes, "headers")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *TdHTMLElement) HIDDEN(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *TdHTMLElement) RemoveHIDDEN(v string) *TdHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *TdHTMLElement) ID(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *TdHTMLElement) RemoveID(v string) *TdHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *TdHTMLElement) INERT() *TdHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *TdHTMLElement) RemoveINERT() *TdHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *TdHTMLElement) SetINERT(b bool) *TdHTMLElement {
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
func (e *TdHTMLElement) INPUTMODE(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *TdHTMLElement) RemoveINPUTMODE(v string) *TdHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *TdHTMLElement) IS(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *TdHTMLElement) RemoveIS(v string) *TdHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *TdHTMLElement) ITEMID(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *TdHTMLElement) RemoveITEMID(v string) *TdHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *TdHTMLElement) ITEMPROP(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *TdHTMLElement) RemoveITEMPROP(v string) *TdHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *TdHTMLElement) ITEMREF(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *TdHTMLElement) RemoveITEMREF(v string) *TdHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *TdHTMLElement) ITEMSCOPE() *TdHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *TdHTMLElement) RemoveITEMSCOPE() *TdHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *TdHTMLElement) SetITEMSCOPE(b bool) *TdHTMLElement {
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
func (e *TdHTMLElement) ITEMTYPE(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *TdHTMLElement) RemoveITEMTYPE(v string) *TdHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TdHTMLElement) LANG(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *TdHTMLElement) RemoveLANG(v string) *TdHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *TdHTMLElement) NONCE(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *TdHTMLElement) RemoveNONCE(v string) *TdHTMLElement {
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
func (e *TdHTMLElement) POPOVER(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *TdHTMLElement) RemovePOPOVER(v string) *TdHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// ROWSPAN sets the "rowspan" attribute.
// Number of rows that the cell is to span
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *TdHTMLElement) ROWSPAN(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["rowspan"] = v
    return e
}

func (e *TdHTMLElement) RemoveROWSPAN(v string) *TdHTMLElement {
    delete(e.StringAttributes, "rowspan")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *TdHTMLElement) SLOT(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *TdHTMLElement) RemoveSLOT(v string) *TdHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *TdHTMLElement) SPELLCHECK(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *TdHTMLElement) RemoveSPELLCHECK(v string) *TdHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TdHTMLElement) STYLE(k,v string) *TdHTMLElement {
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

func (e *TdHTMLElement) RemoveSTYLE(k string) *TdHTMLElement {
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
func (e *TdHTMLElement) TABINDEX(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *TdHTMLElement) RemoveTABINDEX(v string) *TdHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *TdHTMLElement) TITLE(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *TdHTMLElement) RemoveTITLE(v string) *TdHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *TdHTMLElement) TRANSLATE(v string) *TdHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *TdHTMLElement) RemoveTRANSLATE(v string) *TdHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
