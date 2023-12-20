package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type AreaElement struct {
    *gostar.Element
}

func AREA(children ...fmt.Stringer) *AreaElement {
    return &AreaElement{
        Element: &gostar.Element{
            Tag: "area",
            IsSelfClosing: true,
            Children: children,
        },
    }
}

func (e *AreaElement) AddChildren(children ...fmt.Stringer) *AreaElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *AreaElement) TEXT(text string) *AreaElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *AreaElement) RAW(text string) *AreaElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *AreaElement) CustomData(key, value string) *AreaElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AreaElement) CustomDataRemove(key string) *AreaElement {
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
func (e *AreaElement) ACCESSKEY(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *AreaElement) RemoveACCESSKEY(v string) *AreaElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ALT sets the "alt" attribute.
// Replacement text for use when images are not available
// Values values are constrained to:
//  * text
func (e *AreaElement) ALT(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["alt"] = v
    return e
}

func (e *AreaElement) RemoveALT(v string) *AreaElement {
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
func (e *AreaElement) AUTOCAPITALIZE(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *AreaElement) RemoveAUTOCAPITALIZE(v string) *AreaElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *AreaElement) AUTOFOCUS() *AreaElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *AreaElement) RemoveAUTOFOCUS() *AreaElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *AreaElement) SetAUTOFOCUS(b bool) *AreaElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *AreaElement) CLASS(v string) *AreaElement {
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

func (e *AreaElement) SetCLASS(v string) *AreaElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *AreaElement) RemoveCLASS(v string) *AreaElement {
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
func (e *AreaElement) CONTENTEDITABLE(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *AreaElement) RemoveCONTENTEDITABLE(v string) *AreaElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// COORDS sets the "coords" attribute.
// Coordinates for the shape to be created in an image map
// Values values are constrained to:
//  * valid_list_of_floating_point_numbers
func (e *AreaElement) COORDS(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["coords"] = v
    return e
}

func (e *AreaElement) RemoveCOORDS(v string) *AreaElement {
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
func (e *AreaElement) DIR(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *AreaElement) RemoveDIR(v string) *AreaElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DOWNLOAD sets the "download" attribute.
// Whether to download the resource instead of navigating to it, and its filename if so
// Values values are constrained to:
func (e *AreaElement) DOWNLOAD(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["download"] = v
    return e
}

func (e *AreaElement) RemoveDOWNLOAD(v string) *AreaElement {
    delete(e.StringAttributes, "download")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *AreaElement) DRAGGABLE(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *AreaElement) RemoveDRAGGABLE(v string) *AreaElement {
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
func (e *AreaElement) ENTERKEYHINT(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *AreaElement) RemoveENTERKEYHINT(v string) *AreaElement {
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
func (e *AreaElement) HIDDEN(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *AreaElement) RemoveHIDDEN(v string) *AreaElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// HREF sets the "href" attribute.
// Document base URL
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *AreaElement) HREF(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["href"] = v
    return e
}

func (e *AreaElement) RemoveHREF(v string) *AreaElement {
    delete(e.StringAttributes, "href")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *AreaElement) ID(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *AreaElement) RemoveID(v string) *AreaElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *AreaElement) INERT() *AreaElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *AreaElement) RemoveINERT() *AreaElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *AreaElement) SetINERT(b bool) *AreaElement {
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
func (e *AreaElement) INPUTMODE(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *AreaElement) RemoveINPUTMODE(v string) *AreaElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *AreaElement) IS(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *AreaElement) RemoveIS(v string) *AreaElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *AreaElement) ITEMID(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *AreaElement) RemoveITEMID(v string) *AreaElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *AreaElement) ITEMPROP(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *AreaElement) RemoveITEMPROP(v string) *AreaElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *AreaElement) ITEMREF(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *AreaElement) RemoveITEMREF(v string) *AreaElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *AreaElement) ITEMSCOPE() *AreaElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *AreaElement) RemoveITEMSCOPE() *AreaElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *AreaElement) SetITEMSCOPE(b bool) *AreaElement {
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
func (e *AreaElement) ITEMTYPE(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *AreaElement) RemoveITEMTYPE(v string) *AreaElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AreaElement) LANG(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *AreaElement) RemoveLANG(v string) *AreaElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *AreaElement) NONCE(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *AreaElement) RemoveNONCE(v string) *AreaElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// PING sets the "ping" attribute.
// URLs to ping
// Values values are constrained to:
//  * set_of_space_separated_tokens
//  * valid_non_empty_ur_ls
func (e *AreaElement) PING(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["ping"] = v
    return e
}

func (e *AreaElement) RemovePING(v string) *AreaElement {
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
func (e *AreaElement) POPOVER(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *AreaElement) RemovePOPOVER(v string) *AreaElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *AreaElement) REFERRERPOLICY(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *AreaElement) RemoveREFERRERPOLICY(v string) *AreaElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// REL sets the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *AreaElement) REL(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["rel"] = v
    return e
}

func (e *AreaElement) RemoveREL(v string) *AreaElement {
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
func (e *AreaElement) SHAPE(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["shape"] = v
    return e
}

func (e *AreaElement) RemoveSHAPE(v string) *AreaElement {
    delete(e.StringAttributes, "shape")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *AreaElement) SLOT(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *AreaElement) RemoveSLOT(v string) *AreaElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *AreaElement) SPELLCHECK(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *AreaElement) RemoveSPELLCHECK(v string) *AreaElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AreaElement) STYLE(k,v string) *AreaElement {
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

func (e *AreaElement) RemoveSTYLE(k string) *AreaElement {
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
func (e *AreaElement) TABINDEX(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *AreaElement) RemoveTABINDEX(v string) *AreaElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TARGET sets the "target" attribute.
// Navigable for form submission
// Values values are constrained to:
//  * valid_navigable_target_name_or_keyword
func (e *AreaElement) TARGET(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["target"] = v
    return e
}

func (e *AreaElement) RemoveTARGET(v string) *AreaElement {
    delete(e.StringAttributes, "target")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *AreaElement) TITLE(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *AreaElement) RemoveTITLE(v string) *AreaElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *AreaElement) TRANSLATE(v string) *AreaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *AreaElement) RemoveTRANSLATE(v string) *AreaElement {
    delete(e.StringAttributes, "translate")
    return e
}
