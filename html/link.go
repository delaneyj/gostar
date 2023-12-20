package  html

import (
    "fmt"
)

type LinkHTMLElement struct {
    *Element
}

func LINK(children ...ElementBuilder) *LinkHTMLElement {
    return &LinkHTMLElement{
        Element: &Element{
            Tag: "link",
            IsSelfClosing: true,
            Descendants: children,
        },
    }
}

func (e *LinkHTMLElement) Children(children ...ElementBuilder) *LinkHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *LinkHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *LinkHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *LinkHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *LinkHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *LinkHTMLElement) Text(text string) *LinkHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *LinkHTMLElement) TextF(format string, args ...any) *LinkHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *LinkHTMLElement) Raw(text string) *LinkHTMLElement {
    e.Descendants = append(e.Descendants, Raw(text))
    return e
}

func (e *LinkHTMLElement) RawF(format string, args ...any) *LinkHTMLElement {
    return e.Raw(fmt.Sprintf(format, args...))
}

func (e *LinkHTMLElement) CustomData(key, value string) *LinkHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *LinkHTMLElement) CustomDataRemove(key string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) ACCESSKEY(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *LinkHTMLElement) RemoveACCESSKEY(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) AS(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["as"] = v
    return e
}

func (e *LinkHTMLElement) RemoveAS(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) AUTOCAPITALIZE(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *LinkHTMLElement) RemoveAUTOCAPITALIZE(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *LinkHTMLElement) AUTOFOCUS() *LinkHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *LinkHTMLElement) RemoveAUTOFOCUS() *LinkHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *LinkHTMLElement) SetAUTOFOCUS(b bool) *LinkHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// BLOCKING sets the "blocking" attribute.
// Whether the element is potentially render-blocking
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *LinkHTMLElement) BLOCKING(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["blocking"] = v
    return e
}

func (e *LinkHTMLElement) RemoveBLOCKING(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "blocking")
    return e
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *LinkHTMLElement) CLASS(v string) *LinkHTMLElement {
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

func (e *LinkHTMLElement) SetCLASS(v string) *LinkHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *LinkHTMLElement) RemoveCLASS(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) COLOR(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["color"] = v
    return e
}

func (e *LinkHTMLElement) RemoveCOLOR(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "color")
    return e
}
// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//  * true
//  * plaintext_only
//  * false
func (e *LinkHTMLElement) CONTENTEDITABLE(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *LinkHTMLElement) RemoveCONTENTEDITABLE(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) CROSSORIGIN(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *LinkHTMLElement) RemoveCROSSORIGIN(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) DIR(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *LinkHTMLElement) RemoveDIR(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *LinkHTMLElement) DISABLED() *LinkHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *LinkHTMLElement) RemoveDISABLED() *LinkHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *LinkHTMLElement) SetDISABLED(b bool) *LinkHTMLElement {
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
func (e *LinkHTMLElement) DRAGGABLE(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *LinkHTMLElement) RemoveDRAGGABLE(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) ENTERKEYHINT(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *LinkHTMLElement) RemoveENTERKEYHINT(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) FETCHPRIORITY(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["fetchpriority"] = v
    return e
}

func (e *LinkHTMLElement) RemoveFETCHPRIORITY(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) HIDDEN(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *LinkHTMLElement) RemoveHIDDEN(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// HREF sets the "href" attribute.
// Document base URL
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *LinkHTMLElement) HREF(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["href"] = v
    return e
}

func (e *LinkHTMLElement) RemoveHREF(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "href")
    return e
}
// HREFLANG sets the "hreflang" attribute.
// Language of the linked resource
// Values values are constrained to:
func (e *LinkHTMLElement) HREFLANG(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hreflang"] = v
    return e
}

func (e *LinkHTMLElement) RemoveHREFLANG(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "hreflang")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *LinkHTMLElement) ID(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *LinkHTMLElement) RemoveID(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// IMAGESIZES sets the "imagesizes" attribute.
// Image sizes for different page layouts (for rel="preload")
// Values values are constrained to:
//  * valid_source_size_list
func (e *LinkHTMLElement) IMAGESIZES(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["imagesizes"] = v
    return e
}

func (e *LinkHTMLElement) RemoveIMAGESIZES(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "imagesizes")
    return e
}
// IMAGESRCSET sets the "imagesrcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc. (for rel="preload")
// Values values are constrained to:
//  * image_candidate_strings
func (e *LinkHTMLElement) IMAGESRCSET(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["imagesrcset"] = v
    return e
}

func (e *LinkHTMLElement) RemoveIMAGESRCSET(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "imagesrcset")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *LinkHTMLElement) INERT() *LinkHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *LinkHTMLElement) RemoveINERT() *LinkHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *LinkHTMLElement) SetINERT(b bool) *LinkHTMLElement {
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
func (e *LinkHTMLElement) INPUTMODE(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *LinkHTMLElement) RemoveINPUTMODE(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// INTEGRITY sets the "integrity" attribute.
// Integrity metadata used in Subresource Integrity checks [SRI]
// Values values are constrained to:
//  * text
func (e *LinkHTMLElement) INTEGRITY(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["integrity"] = v
    return e
}

func (e *LinkHTMLElement) RemoveINTEGRITY(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "integrity")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *LinkHTMLElement) IS(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *LinkHTMLElement) RemoveIS(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *LinkHTMLElement) ITEMID(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *LinkHTMLElement) RemoveITEMID(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *LinkHTMLElement) ITEMPROP(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *LinkHTMLElement) RemoveITEMPROP(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *LinkHTMLElement) ITEMREF(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *LinkHTMLElement) RemoveITEMREF(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *LinkHTMLElement) ITEMSCOPE() *LinkHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *LinkHTMLElement) RemoveITEMSCOPE() *LinkHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *LinkHTMLElement) SetITEMSCOPE(b bool) *LinkHTMLElement {
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
func (e *LinkHTMLElement) ITEMTYPE(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *LinkHTMLElement) RemoveITEMTYPE(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *LinkHTMLElement) LANG(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *LinkHTMLElement) RemoveLANG(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// MEDIA sets the "media" attribute.
// Applicable media
// Values values are constrained to:
//  * valid_media_query_list
func (e *LinkHTMLElement) MEDIA(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["media"] = v
    return e
}

func (e *LinkHTMLElement) RemoveMEDIA(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "media")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *LinkHTMLElement) NONCE(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *LinkHTMLElement) RemoveNONCE(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) POPOVER(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *LinkHTMLElement) RemovePOPOVER(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//  * referrer_policy
func (e *LinkHTMLElement) REFERRERPOLICY(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["referrerpolicy"] = v
    return e
}

func (e *LinkHTMLElement) RemoveREFERRERPOLICY(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "referrerpolicy")
    return e
}
// REL sets the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *LinkHTMLElement) REL(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["rel"] = v
    return e
}

func (e *LinkHTMLElement) RemoveREL(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "rel")
    return e
}
// SIZES sets the "sizes" attribute.
// Image sizes for different page layouts
// Values values are constrained to:
//  * valid_source_size_list
func (e *LinkHTMLElement) SIZES(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["sizes"] = v
    return e
}

func (e *LinkHTMLElement) RemoveSIZES(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "sizes")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *LinkHTMLElement) SLOT(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *LinkHTMLElement) RemoveSLOT(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *LinkHTMLElement) SPELLCHECK(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *LinkHTMLElement) RemoveSPELLCHECK(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *LinkHTMLElement) STYLE(k,v string) *LinkHTMLElement {
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

func (e *LinkHTMLElement) RemoveSTYLE(k string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) TABINDEX(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *LinkHTMLElement) RemoveTABINDEX(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *LinkHTMLElement) TITLE(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *LinkHTMLElement) RemoveTITLE(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *LinkHTMLElement) TRANSLATE(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *LinkHTMLElement) RemoveTRANSLATE(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *LinkHTMLElement) TYPE(v string) *LinkHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *LinkHTMLElement) RemoveTYPE(v string) *LinkHTMLElement {
    delete(e.StringAttributes, "type")
    return e
}
