package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type LinkElement struct {
    *gostar.Element
}

func LINK(children ...fmt.Stringer) *LinkElement {
    return &LinkElement{
        Element: &gostar.Element{
            Tag: "link",
            IsSelfClosing: true,
            Children: children,
        },
    }
}

func (e *LinkElement) AddChildren(children ...fmt.Stringer) *LinkElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *LinkElement) TEXT(text string) *LinkElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *LinkElement) RAW(text string) *LinkElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *LinkElement) CustomData(key, value string) *LinkElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *LinkElement) CustomDataRemove(key string) *LinkElement {
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
func (e *LinkElement) ACCESSKEY(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *LinkElement) RemoveACCESSKEY(v string) *LinkElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// AS sets the "as" attribute.
// Potential destination for a preload request (for rel="preload" and rel="modulepreload")
// Values values are constrained to:
//  * potential_destination
//  * rel
//  * rel
//  * preload
//  * preload
//  * script_like_destination
//  * rel
//  * rel
//  * modulepreload
//  * modulepreload
func (e *LinkElement) AS(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["as"] = v
    return e
}

func (e *LinkElement) RemoveAS(v string) *LinkElement {
    delete(e.StringAttributes, "as")
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
func (e *LinkElement) AUTOCAPITALIZE(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *LinkElement) RemoveAUTOCAPITALIZE(v string) *LinkElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *LinkElement) AUTOFOCUS() *LinkElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *LinkElement) RemoveAUTOFOCUS() *LinkElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *LinkElement) SetAUTOFOCUS(b bool) *LinkElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// BLOCKING sets the "blocking" attribute.
// Whether the element is potentially render-blocking
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *LinkElement) BLOCKING(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["blocking"] = v
    return e
}

func (e *LinkElement) RemoveBLOCKING(v string) *LinkElement {
    delete(e.StringAttributes, "blocking")
    return e
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *LinkElement) CLASS(v string) *LinkElement {
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

func (e *LinkElement) SetCLASS(v string) *LinkElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *LinkElement) RemoveCLASS(v string) *LinkElement {
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        return e
    }
    kv.Remove(v)
    return e
}
// COLOR sets the "color" attribute.
// Color to use when customizing a site's icon (for rel="mask-icon")
// Values values are constrained to:
//  * <color>
func (e *LinkElement) COLOR(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["color"] = v
    return e
}

func (e *LinkElement) RemoveCOLOR(v string) *LinkElement {
    delete(e.StringAttributes, "color")
    return e
}
// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//  * true
//  * plaintext_only
//  * false
func (e *LinkElement) CONTENTEDITABLE(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *LinkElement) RemoveCONTENTEDITABLE(v string) *LinkElement {
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
func (e *LinkElement) CROSSORIGIN(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *LinkElement) RemoveCROSSORIGIN(v string) *LinkElement {
    delete(e.StringAttributes, "crossorigin")
    return e
}
// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//  * ltr
//  * ltr
//  * rtl
//  * rtl
func (e *LinkElement) DIR(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *LinkElement) RemoveDIR(v string) *LinkElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *LinkElement) DISABLED() *LinkElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *LinkElement) RemoveDISABLED() *LinkElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *LinkElement) SetDISABLED(b bool) *LinkElement {
    if b {
        return e.DISABLED()
    }
    return e.RemoveDISABLED()
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *LinkElement) DRAGGABLE(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *LinkElement) RemoveDRAGGABLE(v string) *LinkElement {
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
func (e *LinkElement) ENTERKEYHINT(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *LinkElement) RemoveENTERKEYHINT(v string) *LinkElement {
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
func (e *LinkElement) FETCHPRIORITY(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["fetchpriority"] = v
    return e
}

func (e *LinkElement) RemoveFETCHPRIORITY(v string) *LinkElement {
    delete(e.StringAttributes, "fetchpriority")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *LinkElement) HIDDEN(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *LinkElement) RemoveHIDDEN(v string) *LinkElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// HREF sets the "href" attribute.
// Document base URL
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *LinkElement) HREF(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["href"] = v
    return e
}

func (e *LinkElement) RemoveHREF(v string) *LinkElement {
    delete(e.StringAttributes, "href")
    return e
}
// HREFLANG sets the "hreflang" attribute.
// Language of the linked resource
// Values values are constrained to:
func (e *LinkElement) HREFLANG(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hreflang"] = v
    return e
}

func (e *LinkElement) RemoveHREFLANG(v string) *LinkElement {
    delete(e.StringAttributes, "hreflang")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *LinkElement) ID(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *LinkElement) RemoveID(v string) *LinkElement {
    delete(e.StringAttributes, "id")
    return e
}
// IMAGESIZES sets the "imagesizes" attribute.
// Image sizes for different page layouts (for rel="preload")
// Values values are constrained to:
//  * valid_source_size_list
func (e *LinkElement) IMAGESIZES(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["imagesizes"] = v
    return e
}

func (e *LinkElement) RemoveIMAGESIZES(v string) *LinkElement {
    delete(e.StringAttributes, "imagesizes")
    return e
}
// IMAGESRCSET sets the "imagesrcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc. (for rel="preload")
// Values values are constrained to:
//  * image_candidate_strings
func (e *LinkElement) IMAGESRCSET(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["imagesrcset"] = v
    return e
}

func (e *LinkElement) RemoveIMAGESRCSET(v string) *LinkElement {
    delete(e.StringAttributes, "imagesrcset")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *LinkElement) INERT() *LinkElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *LinkElement) RemoveINERT() *LinkElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *LinkElement) SetINERT(b bool) *LinkElement {
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
func (e *LinkElement) INPUTMODE(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *LinkElement) RemoveINPUTMODE(v string) *LinkElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// INTEGRITY sets the "integrity" attribute.
// Integrity metadata used in Subresource Integrity checks [SRI]
// Values values are constrained to:
//  * text
func (e *LinkElement) INTEGRITY(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["integrity"] = v
    return e
}

func (e *LinkElement) RemoveINTEGRITY(v string) *LinkElement {
    delete(e.StringAttributes, "integrity")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *LinkElement) IS(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *LinkElement) RemoveIS(v string) *LinkElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *LinkElement) ITEMID(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *LinkElement) RemoveITEMID(v string) *LinkElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *LinkElement) ITEMPROP(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *LinkElement) RemoveITEMPROP(v string) *LinkElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *LinkElement) ITEMREF(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *LinkElement) RemoveITEMREF(v string) *LinkElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *LinkElement) ITEMSCOPE() *LinkElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *LinkElement) RemoveITEMSCOPE() *LinkElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *LinkElement) SetITEMSCOPE(b bool) *LinkElement {
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
func (e *LinkElement) ITEMTYPE(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *LinkElement) RemoveITEMTYPE(v string) *LinkElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *LinkElement) LANG(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *LinkElement) RemoveLANG(v string) *LinkElement {
    delete(e.StringAttributes, "lang")
    return e
}
// MEDIA sets the "media" attribute.
// Applicable media
// Values values are constrained to:
//  * valid_media_query_list
func (e *LinkElement) MEDIA(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["media"] = v
    return e
}

func (e *LinkElement) RemoveMEDIA(v string) *LinkElement {
    delete(e.StringAttributes, "media")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *LinkElement) NONCE(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *LinkElement) RemoveNONCE(v string) *LinkElement {
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
func (e *LinkElement) POPOVER(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *LinkElement) RemovePOPOVER(v string) *LinkElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *LinkElement) REFERRERPOLICY(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *LinkElement) RemoveREFERRERPOLICY(v string) *LinkElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// REL sets the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *LinkElement) REL(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["rel"] = v
    return e
}

func (e *LinkElement) RemoveREL(v string) *LinkElement {
    delete(e.StringAttributes, "rel")
    return e
}
// SIZES sets the "sizes" attribute.
// Image sizes for different page layouts
// Values values are constrained to:
//  * valid_source_size_list
func (e *LinkElement) SIZES(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["sizes"] = v
    return e
}

func (e *LinkElement) RemoveSIZES(v string) *LinkElement {
    delete(e.StringAttributes, "sizes")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *LinkElement) SLOT(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *LinkElement) RemoveSLOT(v string) *LinkElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *LinkElement) SPELLCHECK(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *LinkElement) RemoveSPELLCHECK(v string) *LinkElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *LinkElement) STYLE(k,v string) *LinkElement {
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

func (e *LinkElement) RemoveSTYLE(k string) *LinkElement {
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
func (e *LinkElement) TABINDEX(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *LinkElement) RemoveTABINDEX(v string) *LinkElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *LinkElement) TITLE(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *LinkElement) RemoveTITLE(v string) *LinkElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *LinkElement) TRANSLATE(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *LinkElement) RemoveTRANSLATE(v string) *LinkElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *LinkElement) TYPE(v string) *LinkElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *LinkElement) RemoveTYPE(v string) *LinkElement {
    delete(e.StringAttributes, "type")
    return e
}
