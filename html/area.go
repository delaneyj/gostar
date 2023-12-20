package  html

import (
    "fmt"
)

type AreaHTMLElement struct {
    *Element
}

func AREA(children ...fmt.Stringer) *AreaHTMLElement {
    return &AreaHTMLElement{
        Element: &Element{
            Tag: "area",
            IsSelfClosing: true,
            Descendants: children,
        },
    }
}

func (e *AreaHTMLElement) Children(children ...fmt.Stringer) *AreaHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *AreaHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *AreaHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *AreaHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *AreaHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *AreaHTMLElement) Text(text string) *AreaHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *AreaHTMLElement) TextF(format string, args ...any) *AreaHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *AreaHTMLElement) Raw(text string) *AreaHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *AreaHTMLElement) RawF(format string, args ...any) *AreaHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *AreaHTMLElement) CustomData(key, value string) *AreaHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AreaHTMLElement) CustomDataRemove(key string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) ACCESSKEY(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *AreaHTMLElement) RemoveACCESSKEY(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ALT sets the "alt" attribute.
// Replacement text for use when images are not available
// Values values are constrained to:
//  * text
func (e *AreaHTMLElement) ALT(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["alt"] = v
    return e
}

func (e *AreaHTMLElement) RemoveALT(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "alt")
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
func (e *AreaHTMLElement) AUTOCAPITALIZE(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *AreaHTMLElement) RemoveAUTOCAPITALIZE(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *AreaHTMLElement) AUTOFOCUS() *AreaHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *AreaHTMLElement) RemoveAUTOFOCUS() *AreaHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *AreaHTMLElement) SetAUTOFOCUS(b bool) *AreaHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *AreaHTMLElement) CLASS(v string) *AreaHTMLElement {
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

func (e *AreaHTMLElement) SetCLASS(v string) *AreaHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *AreaHTMLElement) RemoveCLASS(v string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) CONTENTEDITABLE(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *AreaHTMLElement) RemoveCONTENTEDITABLE(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// COORDS sets the "coords" attribute.
// Coordinates for the shape to be created in an image map
// Values values are constrained to:
//  * valid_list_of_floating_point_numbers
func (e *AreaHTMLElement) COORDS(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["coords"] = v
    return e
}

func (e *AreaHTMLElement) RemoveCOORDS(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "coords")
    return e
}
// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//  * ltr
//  * ltr
//  * rtl
//  * rtl
func (e *AreaHTMLElement) DIR(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *AreaHTMLElement) RemoveDIR(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DOWNLOAD sets the "download" attribute.
// Whether to download the resource instead of navigating to it, and its filename if so
// Values values are constrained to:
func (e *AreaHTMLElement) DOWNLOAD(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["download"] = v
    return e
}

func (e *AreaHTMLElement) RemoveDOWNLOAD(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "download")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *AreaHTMLElement) DRAGGABLE(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *AreaHTMLElement) RemoveDRAGGABLE(v string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) ENTERKEYHINT(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *AreaHTMLElement) RemoveENTERKEYHINT(v string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) HIDDEN(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *AreaHTMLElement) RemoveHIDDEN(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// HREF sets the "href" attribute.
// Document base URL
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *AreaHTMLElement) HREF(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["href"] = v
    return e
}

func (e *AreaHTMLElement) RemoveHREF(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "href")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *AreaHTMLElement) ID(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *AreaHTMLElement) RemoveID(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *AreaHTMLElement) INERT() *AreaHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *AreaHTMLElement) RemoveINERT() *AreaHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *AreaHTMLElement) SetINERT(b bool) *AreaHTMLElement {
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
func (e *AreaHTMLElement) INPUTMODE(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *AreaHTMLElement) RemoveINPUTMODE(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *AreaHTMLElement) IS(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *AreaHTMLElement) RemoveIS(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *AreaHTMLElement) ITEMID(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *AreaHTMLElement) RemoveITEMID(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *AreaHTMLElement) ITEMPROP(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *AreaHTMLElement) RemoveITEMPROP(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *AreaHTMLElement) ITEMREF(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *AreaHTMLElement) RemoveITEMREF(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *AreaHTMLElement) ITEMSCOPE() *AreaHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *AreaHTMLElement) RemoveITEMSCOPE() *AreaHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *AreaHTMLElement) SetITEMSCOPE(b bool) *AreaHTMLElement {
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
func (e *AreaHTMLElement) ITEMTYPE(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *AreaHTMLElement) RemoveITEMTYPE(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AreaHTMLElement) LANG(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *AreaHTMLElement) RemoveLANG(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *AreaHTMLElement) NONCE(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *AreaHTMLElement) RemoveNONCE(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// PING sets the "ping" attribute.
// URLs to ping
// Values values are constrained to:
//  * set_of_space_separated_tokens
//  * valid_non_empty_ur_ls
func (e *AreaHTMLElement) PING(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["ping"] = v
    return e
}

func (e *AreaHTMLElement) RemovePING(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "ping")
    return e
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *AreaHTMLElement) POPOVER(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *AreaHTMLElement) RemovePOPOVER(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *AreaHTMLElement) REFERRERPOLICY(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *AreaHTMLElement) RemoveREFERRERPOLICY(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// REL sets the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *AreaHTMLElement) REL(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["rel"] = v
    return e
}

func (e *AreaHTMLElement) RemoveREL(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "rel")
    return e
}
// SHAPE sets the "shape" attribute.
// The kind of shape to be created in an image map
// Values values are constrained to:
//  * circle
//  * circle
//  * default
//  * default
//  * poly
//  * poly
//  * rect
//  * rect
func (e *AreaHTMLElement) SHAPE(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["shape"] = v
    return e
}

func (e *AreaHTMLElement) RemoveSHAPE(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "shape")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *AreaHTMLElement) SLOT(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *AreaHTMLElement) RemoveSLOT(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *AreaHTMLElement) SPELLCHECK(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *AreaHTMLElement) RemoveSPELLCHECK(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AreaHTMLElement) STYLE(k,v string) *AreaHTMLElement {
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

func (e *AreaHTMLElement) RemoveSTYLE(k string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) TABINDEX(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *AreaHTMLElement) RemoveTABINDEX(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TARGET sets the "target" attribute.
// Navigable for form submission
// Values values are constrained to:
//  * valid_navigable_target_name_or_keyword
func (e *AreaHTMLElement) TARGET(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["target"] = v
    return e
}

func (e *AreaHTMLElement) RemoveTARGET(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "target")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *AreaHTMLElement) TITLE(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *AreaHTMLElement) RemoveTITLE(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *AreaHTMLElement) TRANSLATE(v string) *AreaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *AreaHTMLElement) RemoveTRANSLATE(v string) *AreaHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
