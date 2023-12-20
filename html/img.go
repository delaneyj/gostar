package  html

import (
    "fmt"
)

type ImgHTMLElement struct {
    *Element
}

func IMG(children ...ElementBuilder) *ImgHTMLElement {
    return &ImgHTMLElement{
        Element: &Element{
            Tag: "img",
            IsSelfClosing: true,
            Descendants: children,
        },
    }
}

func (e *ImgHTMLElement) Children(children ...ElementBuilder) *ImgHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *ImgHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ImgHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *ImgHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ImgHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *ImgHTMLElement) Text(text string) *ImgHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *ImgHTMLElement) TextF(format string, args ...any) *ImgHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *ImgHTMLElement) Raw(text string) *ImgHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *ImgHTMLElement) RawF(format string, args ...any) *ImgHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *ImgHTMLElement) CustomData(key, value string) *ImgHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ImgHTMLElement) CustomDataRemove(key string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) ACCESSKEY(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *ImgHTMLElement) RemoveACCESSKEY(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ALT sets the "alt" attribute.
// Replacement text for use when images are not available
// Values values are constrained to:
//  * text
func (e *ImgHTMLElement) ALT(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["alt"] = v
    return e
}

func (e *ImgHTMLElement) RemoveALT(v string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) AUTOCAPITALIZE(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *ImgHTMLElement) RemoveAUTOCAPITALIZE(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *ImgHTMLElement) AUTOFOCUS() *ImgHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *ImgHTMLElement) RemoveAUTOFOCUS() *ImgHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *ImgHTMLElement) SetAUTOFOCUS(b bool) *ImgHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *ImgHTMLElement) CLASS(v string) *ImgHTMLElement {
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

func (e *ImgHTMLElement) SetCLASS(v string) *ImgHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *ImgHTMLElement) RemoveCLASS(v string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) CONTENTEDITABLE(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *ImgHTMLElement) RemoveCONTENTEDITABLE(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// CROSSORIGIN sets the "crossorigin" attribute.
// How the element handles crossorigin requests
// Values values are constrained to:
//  * anonymous
//  * anonymous
//  * use_credentials
//  * use_credentials
func (e *ImgHTMLElement) CROSSORIGIN(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *ImgHTMLElement) RemoveCROSSORIGIN(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "crossorigin")
    return e
}
// DECODING sets the "decoding" attribute.
// Decoding hint to use when processing this image for presentation
// Values values are constrained to:
//  * sync
//  * sync
//  * async
//  * async
//  * auto
//  * auto
func (e *ImgHTMLElement) DECODING(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["decoding"] = v
    return e
}

func (e *ImgHTMLElement) RemoveDECODING(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "decoding")
    return e
}
// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//  * ltr
//  * ltr
//  * rtl
//  * rtl
func (e *ImgHTMLElement) DIR(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *ImgHTMLElement) RemoveDIR(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *ImgHTMLElement) DRAGGABLE(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *ImgHTMLElement) RemoveDRAGGABLE(v string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) ENTERKEYHINT(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *ImgHTMLElement) RemoveENTERKEYHINT(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FETCHPRIORITY sets the "fetchpriority" attribute.
// Sets the priority for fetches initiated by the element
// Values values are constrained to:
//  * auto
//  * auto
//  * high
//  * high
//  * low
//  * low
func (e *ImgHTMLElement) FETCHPRIORITY(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["fetchpriority"] = v
    return e
}

func (e *ImgHTMLElement) RemoveFETCHPRIORITY(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "fetchpriority")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *ImgHTMLElement) HEIGHT(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *ImgHTMLElement) RemoveHEIGHT(v string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) HIDDEN(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *ImgHTMLElement) RemoveHIDDEN(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *ImgHTMLElement) ID(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *ImgHTMLElement) RemoveID(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *ImgHTMLElement) INERT() *ImgHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *ImgHTMLElement) RemoveINERT() *ImgHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *ImgHTMLElement) SetINERT(b bool) *ImgHTMLElement {
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
func (e *ImgHTMLElement) INPUTMODE(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *ImgHTMLElement) RemoveINPUTMODE(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *ImgHTMLElement) IS(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *ImgHTMLElement) RemoveIS(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ISMAP sets the "ismap" attribute.
// Whether the image is a server-side image map
// Values values are constrained to:
//  * boolean_attribute
func (e *ImgHTMLElement) ISMAP() *ImgHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["ismap"] = struct{}{}
    return e
}

func (e *ImgHTMLElement) RemoveISMAP() *ImgHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "ismap")
    return e
}

func (e *ImgHTMLElement) SetISMAP(b bool) *ImgHTMLElement {
    if b {
        return e.ISMAP()
    }
    return e.RemoveISMAP()
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *ImgHTMLElement) ITEMID(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *ImgHTMLElement) RemoveITEMID(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *ImgHTMLElement) ITEMPROP(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *ImgHTMLElement) RemoveITEMPROP(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ImgHTMLElement) ITEMREF(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *ImgHTMLElement) RemoveITEMREF(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *ImgHTMLElement) ITEMSCOPE() *ImgHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *ImgHTMLElement) RemoveITEMSCOPE() *ImgHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *ImgHTMLElement) SetITEMSCOPE(b bool) *ImgHTMLElement {
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
func (e *ImgHTMLElement) ITEMTYPE(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *ImgHTMLElement) RemoveITEMTYPE(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ImgHTMLElement) LANG(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *ImgHTMLElement) RemoveLANG(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LOADING sets the "loading" attribute.
// Used when determining loading deferral
// Values values are constrained to:
//  * lazy
//  * lazy
//  * eager
//  * eager
func (e *ImgHTMLElement) LOADING(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["loading"] = v
    return e
}

func (e *ImgHTMLElement) RemoveLOADING(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "loading")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *ImgHTMLElement) NONCE(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *ImgHTMLElement) RemoveNONCE(v string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) POPOVER(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *ImgHTMLElement) RemovePOPOVER(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *ImgHTMLElement) REFERRERPOLICY(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *ImgHTMLElement) RemoveREFERRERPOLICY(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// SIZES sets the "sizes" attribute.
// Image sizes for different page layouts
// Values values are constrained to:
//  * valid_source_size_list
func (e *ImgHTMLElement) SIZES(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["sizes"] = v
    return e
}

func (e *ImgHTMLElement) RemoveSIZES(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "sizes")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *ImgHTMLElement) SLOT(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *ImgHTMLElement) RemoveSLOT(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *ImgHTMLElement) SPELLCHECK(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *ImgHTMLElement) RemoveSPELLCHECK(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ImgHTMLElement) SRC(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *ImgHTMLElement) RemoveSRC(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "src")
    return e
}
// SRCSET sets the "srcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc.
// Values values are constrained to:
//  * image_candidate_strings
func (e *ImgHTMLElement) SRCSET(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["srcset"] = v
    return e
}

func (e *ImgHTMLElement) RemoveSRCSET(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "srcset")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ImgHTMLElement) STYLE(k,v string) *ImgHTMLElement {
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

func (e *ImgHTMLElement) RemoveSTYLE(k string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) TABINDEX(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *ImgHTMLElement) RemoveTABINDEX(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *ImgHTMLElement) TITLE(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *ImgHTMLElement) RemoveTITLE(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *ImgHTMLElement) TRANSLATE(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *ImgHTMLElement) RemoveTRANSLATE(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// USEMAP sets the "usemap" attribute.
// Name of image map to use
// Values values are constrained to:
//  * valid_hash_name_reference
func (e *ImgHTMLElement) USEMAP(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["usemap"] = v
    return e
}

func (e *ImgHTMLElement) RemoveUSEMAP(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "usemap")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *ImgHTMLElement) WIDTH(v string) *ImgHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *ImgHTMLElement) RemoveWIDTH(v string) *ImgHTMLElement {
    delete(e.StringAttributes, "width")
    return e
}
