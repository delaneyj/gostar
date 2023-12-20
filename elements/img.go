package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type ImgElement struct {
    *gostar.Element
}

func IMG(children ...fmt.Stringer) *ImgElement {
    return &ImgElement{
        Element: &gostar.Element{
            Tag: "img",
            IsSelfClosing: true,
            Children: children,
        },
    }
}

func (e *ImgElement) AddChildren(children ...fmt.Stringer) *ImgElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *ImgElement) TEXT(text string) *ImgElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *ImgElement) RAW(text string) *ImgElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *ImgElement) CustomData(key, value string) *ImgElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ImgElement) CustomDataRemove(key string) *ImgElement {
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
func (e *ImgElement) ACCESSKEY(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *ImgElement) RemoveACCESSKEY(v string) *ImgElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ALT sets the "alt" attribute.
// Replacement text for use when images are not available
// Values values are constrained to:
//  * text
func (e *ImgElement) ALT(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["alt"] = v
    return e
}

func (e *ImgElement) RemoveALT(v string) *ImgElement {
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
func (e *ImgElement) AUTOCAPITALIZE(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *ImgElement) RemoveAUTOCAPITALIZE(v string) *ImgElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *ImgElement) AUTOFOCUS() *ImgElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *ImgElement) RemoveAUTOFOCUS() *ImgElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *ImgElement) SetAUTOFOCUS(b bool) *ImgElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *ImgElement) CLASS(v string) *ImgElement {
    if e.DelimitedStringAttributes == nil {
        e.DelimitedStringAttributes = map[string]*gostar.DelimitedString{}
    }
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        kv = gostar.NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = kv
    }
    kv.Add(v)
    return e
}

func (e *ImgElement) SetCLASS(v string) *ImgElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *ImgElement) RemoveCLASS(v string) *ImgElement {
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
func (e *ImgElement) CONTENTEDITABLE(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *ImgElement) RemoveCONTENTEDITABLE(v string) *ImgElement {
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
func (e *ImgElement) CROSSORIGIN(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *ImgElement) RemoveCROSSORIGIN(v string) *ImgElement {
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
func (e *ImgElement) DECODING(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["decoding"] = v
    return e
}

func (e *ImgElement) RemoveDECODING(v string) *ImgElement {
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
func (e *ImgElement) DIR(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *ImgElement) RemoveDIR(v string) *ImgElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *ImgElement) DRAGGABLE(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *ImgElement) RemoveDRAGGABLE(v string) *ImgElement {
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
func (e *ImgElement) ENTERKEYHINT(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *ImgElement) RemoveENTERKEYHINT(v string) *ImgElement {
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
func (e *ImgElement) FETCHPRIORITY(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["fetchpriority"] = v
    return e
}

func (e *ImgElement) RemoveFETCHPRIORITY(v string) *ImgElement {
    delete(e.StringAttributes, "fetchpriority")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *ImgElement) HEIGHT(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *ImgElement) RemoveHEIGHT(v string) *ImgElement {
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
func (e *ImgElement) HIDDEN(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *ImgElement) RemoveHIDDEN(v string) *ImgElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *ImgElement) ID(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *ImgElement) RemoveID(v string) *ImgElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *ImgElement) INERT() *ImgElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *ImgElement) RemoveINERT() *ImgElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *ImgElement) SetINERT(b bool) *ImgElement {
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
func (e *ImgElement) INPUTMODE(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *ImgElement) RemoveINPUTMODE(v string) *ImgElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *ImgElement) IS(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *ImgElement) RemoveIS(v string) *ImgElement {
    delete(e.StringAttributes, "is")
    return e
}
// ISMAP sets the "ismap" attribute.
// Whether the image is a server-side image map
// Values values are constrained to:
//  * boolean_attribute
func (e *ImgElement) ISMAP() *ImgElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["ismap"] = struct{}{}
    return e
}

func (e *ImgElement) RemoveISMAP() *ImgElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "ismap")
    return e
}

func (e *ImgElement) SetISMAP(b bool) *ImgElement {
    if b {
        return e.ISMAP()
    }
    return e.RemoveISMAP()
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *ImgElement) ITEMID(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *ImgElement) RemoveITEMID(v string) *ImgElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *ImgElement) ITEMPROP(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *ImgElement) RemoveITEMPROP(v string) *ImgElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ImgElement) ITEMREF(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *ImgElement) RemoveITEMREF(v string) *ImgElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *ImgElement) ITEMSCOPE() *ImgElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *ImgElement) RemoveITEMSCOPE() *ImgElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *ImgElement) SetITEMSCOPE(b bool) *ImgElement {
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
func (e *ImgElement) ITEMTYPE(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *ImgElement) RemoveITEMTYPE(v string) *ImgElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ImgElement) LANG(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *ImgElement) RemoveLANG(v string) *ImgElement {
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
func (e *ImgElement) LOADING(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["loading"] = v
    return e
}

func (e *ImgElement) RemoveLOADING(v string) *ImgElement {
    delete(e.StringAttributes, "loading")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *ImgElement) NONCE(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *ImgElement) RemoveNONCE(v string) *ImgElement {
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
func (e *ImgElement) POPOVER(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *ImgElement) RemovePOPOVER(v string) *ImgElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *ImgElement) REFERRERPOLICY(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *ImgElement) RemoveREFERRERPOLICY(v string) *ImgElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// SIZES sets the "sizes" attribute.
// Image sizes for different page layouts
// Values values are constrained to:
//  * valid_source_size_list
func (e *ImgElement) SIZES(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["sizes"] = v
    return e
}

func (e *ImgElement) RemoveSIZES(v string) *ImgElement {
    delete(e.StringAttributes, "sizes")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *ImgElement) SLOT(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *ImgElement) RemoveSLOT(v string) *ImgElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *ImgElement) SPELLCHECK(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *ImgElement) RemoveSPELLCHECK(v string) *ImgElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ImgElement) SRC(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *ImgElement) RemoveSRC(v string) *ImgElement {
    delete(e.StringAttributes, "src")
    return e
}
// SRCSET sets the "srcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc.
// Values values are constrained to:
//  * image_candidate_strings
func (e *ImgElement) SRCSET(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["srcset"] = v
    return e
}

func (e *ImgElement) RemoveSRCSET(v string) *ImgElement {
    delete(e.StringAttributes, "srcset")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ImgElement) STYLE(k,v string) *ImgElement {
    if e.DelimitedKVAttributes == nil {
        e.DelimitedKVAttributes = map[string]*gostar.DelimitedKVString{}
    }
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        kv = gostar.NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = kv
    }
    kv.Add(k,v)
    return e
}

func (e *ImgElement) RemoveSTYLE(k string) *ImgElement {
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
func (e *ImgElement) TABINDEX(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *ImgElement) RemoveTABINDEX(v string) *ImgElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *ImgElement) TITLE(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *ImgElement) RemoveTITLE(v string) *ImgElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *ImgElement) TRANSLATE(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *ImgElement) RemoveTRANSLATE(v string) *ImgElement {
    delete(e.StringAttributes, "translate")
    return e
}
// USEMAP sets the "usemap" attribute.
// Name of image map to use
// Values values are constrained to:
//  * valid_hash_name_reference
func (e *ImgElement) USEMAP(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["usemap"] = v
    return e
}

func (e *ImgElement) RemoveUSEMAP(v string) *ImgElement {
    delete(e.StringAttributes, "usemap")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *ImgElement) WIDTH(v string) *ImgElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *ImgElement) RemoveWIDTH(v string) *ImgElement {
    delete(e.StringAttributes, "width")
    return e
}
