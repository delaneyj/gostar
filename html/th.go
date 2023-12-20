package  html

import (
    "fmt"
)

type ThHTMLElement struct {
    *Element
}

func TH(children ...fmt.Stringer) *ThHTMLElement {
    return &ThHTMLElement{
        Element: &Element{
            Tag: "th",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *ThHTMLElement) Children(children ...fmt.Stringer) *ThHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *ThHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *ThHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *ThHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *ThHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *ThHTMLElement) Text(text string) *ThHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *ThHTMLElement) TextF(format string, args ...any) *ThHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *ThHTMLElement) Raw(text string) *ThHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *ThHTMLElement) RawF(format string, args ...any) *ThHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *ThHTMLElement) CustomData(key, value string) *ThHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ThHTMLElement) CustomDataRemove(key string) *ThHTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}


// ABBR sets the "abbr" attribute.
// Alternative label to use for the header cell when referencing the cell in other contexts
// Values values are constrained to:
//  * text
func (e *ThHTMLElement) ABBR(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["abbr"] = v
    return e
}

func (e *ThHTMLElement) RemoveABBR(v string) *ThHTMLElement {
    delete(e.StringAttributes, "abbr")
    return e
}
// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//  * ordered_set_of_unique_space_separated_tokens
//  * identical_to
func (e *ThHTMLElement) ACCESSKEY(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *ThHTMLElement) RemoveACCESSKEY(v string) *ThHTMLElement {
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
func (e *ThHTMLElement) AUTOCAPITALIZE(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *ThHTMLElement) RemoveAUTOCAPITALIZE(v string) *ThHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *ThHTMLElement) AUTOFOCUS() *ThHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *ThHTMLElement) RemoveAUTOFOCUS() *ThHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *ThHTMLElement) SetAUTOFOCUS(b bool) *ThHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *ThHTMLElement) CLASS(v string) *ThHTMLElement {
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

func (e *ThHTMLElement) SetCLASS(v string) *ThHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *ThHTMLElement) RemoveCLASS(v string) *ThHTMLElement {
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
func (e *ThHTMLElement) COLSPAN(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["colspan"] = v
    return e
}

func (e *ThHTMLElement) RemoveCOLSPAN(v string) *ThHTMLElement {
    delete(e.StringAttributes, "colspan")
    return e
}
// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//  * true
//  * plaintext_only
//  * false
func (e *ThHTMLElement) CONTENTEDITABLE(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *ThHTMLElement) RemoveCONTENTEDITABLE(v string) *ThHTMLElement {
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
func (e *ThHTMLElement) DIR(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *ThHTMLElement) RemoveDIR(v string) *ThHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *ThHTMLElement) DRAGGABLE(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *ThHTMLElement) RemoveDRAGGABLE(v string) *ThHTMLElement {
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
func (e *ThHTMLElement) ENTERKEYHINT(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *ThHTMLElement) RemoveENTERKEYHINT(v string) *ThHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HEADERS sets the "headers" attribute.
// The header cells for this cell
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ThHTMLElement) HEADERS(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["headers"] = v
    return e
}

func (e *ThHTMLElement) RemoveHEADERS(v string) *ThHTMLElement {
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
func (e *ThHTMLElement) HIDDEN(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *ThHTMLElement) RemoveHIDDEN(v string) *ThHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *ThHTMLElement) ID(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *ThHTMLElement) RemoveID(v string) *ThHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *ThHTMLElement) INERT() *ThHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *ThHTMLElement) RemoveINERT() *ThHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *ThHTMLElement) SetINERT(b bool) *ThHTMLElement {
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
func (e *ThHTMLElement) INPUTMODE(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *ThHTMLElement) RemoveINPUTMODE(v string) *ThHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *ThHTMLElement) IS(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *ThHTMLElement) RemoveIS(v string) *ThHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *ThHTMLElement) ITEMID(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *ThHTMLElement) RemoveITEMID(v string) *ThHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *ThHTMLElement) ITEMPROP(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *ThHTMLElement) RemoveITEMPROP(v string) *ThHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ThHTMLElement) ITEMREF(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *ThHTMLElement) RemoveITEMREF(v string) *ThHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *ThHTMLElement) ITEMSCOPE() *ThHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *ThHTMLElement) RemoveITEMSCOPE() *ThHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *ThHTMLElement) SetITEMSCOPE(b bool) *ThHTMLElement {
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
func (e *ThHTMLElement) ITEMTYPE(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *ThHTMLElement) RemoveITEMTYPE(v string) *ThHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ThHTMLElement) LANG(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *ThHTMLElement) RemoveLANG(v string) *ThHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *ThHTMLElement) NONCE(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *ThHTMLElement) RemoveNONCE(v string) *ThHTMLElement {
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
func (e *ThHTMLElement) POPOVER(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *ThHTMLElement) RemovePOPOVER(v string) *ThHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// ROWSPAN sets the "rowspan" attribute.
// Number of rows that the cell is to span
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *ThHTMLElement) ROWSPAN(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["rowspan"] = v
    return e
}

func (e *ThHTMLElement) RemoveROWSPAN(v string) *ThHTMLElement {
    delete(e.StringAttributes, "rowspan")
    return e
}
// SCOPE sets the "scope" attribute.
// Specifies which cells the header cell applies to
// Values values are constrained to:
//  * row
//  * row
//  * col
//  * col
//  * rowgroup
//  * rowgroup
//  * colgroup
//  * colgroup
func (e *ThHTMLElement) SCOPE(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["scope"] = v
    return e
}

func (e *ThHTMLElement) RemoveSCOPE(v string) *ThHTMLElement {
    delete(e.StringAttributes, "scope")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *ThHTMLElement) SLOT(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *ThHTMLElement) RemoveSLOT(v string) *ThHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *ThHTMLElement) SPELLCHECK(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *ThHTMLElement) RemoveSPELLCHECK(v string) *ThHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ThHTMLElement) STYLE(k,v string) *ThHTMLElement {
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

func (e *ThHTMLElement) RemoveSTYLE(k string) *ThHTMLElement {
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
func (e *ThHTMLElement) TABINDEX(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *ThHTMLElement) RemoveTABINDEX(v string) *ThHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *ThHTMLElement) TITLE(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *ThHTMLElement) RemoveTITLE(v string) *ThHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *ThHTMLElement) TRANSLATE(v string) *ThHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *ThHTMLElement) RemoveTRANSLATE(v string) *ThHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
