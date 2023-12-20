package  html

import (
    "fmt"
)

type SourceHTMLElement struct {
    *Element
}

func SOURCE(children ...ElementBuilder) *SourceHTMLElement {
    return &SourceHTMLElement{
        Element: &Element{
            Tag: "source",
            IsSelfClosing: true,
            Descendants: children,
        },
    }
}

func (e *SourceHTMLElement) Children(children ...ElementBuilder) *SourceHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *SourceHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SourceHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *SourceHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SourceHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *SourceHTMLElement) Text(text string) *SourceHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *SourceHTMLElement) TextF(format string, args ...any) *SourceHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *SourceHTMLElement) Raw(text string) *SourceHTMLElement {
    e.Descendants = append(e.Descendants, Raw(text))
    return e
}

func (e *SourceHTMLElement) RawF(format string, args ...any) *SourceHTMLElement {
    return e.Raw(fmt.Sprintf(format, args...))
}

func (e *SourceHTMLElement) CustomData(key, value string) *SourceHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SourceHTMLElement) CustomDataRemove(key string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) ACCESSKEY(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *SourceHTMLElement) RemoveACCESSKEY(v string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) AUTOCAPITALIZE(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *SourceHTMLElement) RemoveAUTOCAPITALIZE(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *SourceHTMLElement) AUTOFOCUS() *SourceHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *SourceHTMLElement) RemoveAUTOFOCUS() *SourceHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *SourceHTMLElement) SetAUTOFOCUS(b bool) *SourceHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *SourceHTMLElement) CLASS(v string) *SourceHTMLElement {
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

func (e *SourceHTMLElement) SetCLASS(v string) *SourceHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *SourceHTMLElement) RemoveCLASS(v string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) CONTENTEDITABLE(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *SourceHTMLElement) RemoveCONTENTEDITABLE(v string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) DIR(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *SourceHTMLElement) RemoveDIR(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *SourceHTMLElement) DRAGGABLE(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *SourceHTMLElement) RemoveDRAGGABLE(v string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) ENTERKEYHINT(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *SourceHTMLElement) RemoveENTERKEYHINT(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *SourceHTMLElement) HEIGHT(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *SourceHTMLElement) RemoveHEIGHT(v string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) HIDDEN(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *SourceHTMLElement) RemoveHIDDEN(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *SourceHTMLElement) ID(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *SourceHTMLElement) RemoveID(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *SourceHTMLElement) INERT() *SourceHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *SourceHTMLElement) RemoveINERT() *SourceHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *SourceHTMLElement) SetINERT(b bool) *SourceHTMLElement {
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
func (e *SourceHTMLElement) INPUTMODE(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *SourceHTMLElement) RemoveINPUTMODE(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *SourceHTMLElement) IS(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *SourceHTMLElement) RemoveIS(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *SourceHTMLElement) ITEMID(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *SourceHTMLElement) RemoveITEMID(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *SourceHTMLElement) ITEMPROP(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *SourceHTMLElement) RemoveITEMPROP(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *SourceHTMLElement) ITEMREF(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *SourceHTMLElement) RemoveITEMREF(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *SourceHTMLElement) ITEMSCOPE() *SourceHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *SourceHTMLElement) RemoveITEMSCOPE() *SourceHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *SourceHTMLElement) SetITEMSCOPE(b bool) *SourceHTMLElement {
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
func (e *SourceHTMLElement) ITEMTYPE(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *SourceHTMLElement) RemoveITEMTYPE(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SourceHTMLElement) LANG(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *SourceHTMLElement) RemoveLANG(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// MEDIA sets the "media" attribute.
// Applicable media
// Values values are constrained to:
//  * valid_media_query_list
func (e *SourceHTMLElement) MEDIA(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["media"] = v
    return e
}

func (e *SourceHTMLElement) RemoveMEDIA(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "media")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *SourceHTMLElement) NONCE(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *SourceHTMLElement) RemoveNONCE(v string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) POPOVER(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *SourceHTMLElement) RemovePOPOVER(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SIZES sets the "sizes" attribute.
// Image sizes for different page layouts
// Values values are constrained to:
//  * valid_source_size_list
func (e *SourceHTMLElement) SIZES(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["sizes"] = v
    return e
}

func (e *SourceHTMLElement) RemoveSIZES(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "sizes")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *SourceHTMLElement) SLOT(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *SourceHTMLElement) RemoveSLOT(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *SourceHTMLElement) SPELLCHECK(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *SourceHTMLElement) RemoveSPELLCHECK(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *SourceHTMLElement) SRC(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *SourceHTMLElement) RemoveSRC(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "src")
    return e
}
// SRCSET sets the "srcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc.
// Values values are constrained to:
//  * image_candidate_strings
func (e *SourceHTMLElement) SRCSET(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["srcset"] = v
    return e
}

func (e *SourceHTMLElement) RemoveSRCSET(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "srcset")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SourceHTMLElement) STYLE(k,v string) *SourceHTMLElement {
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

func (e *SourceHTMLElement) RemoveSTYLE(k string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) TABINDEX(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *SourceHTMLElement) RemoveTABINDEX(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *SourceHTMLElement) TITLE(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *SourceHTMLElement) RemoveTITLE(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *SourceHTMLElement) TRANSLATE(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *SourceHTMLElement) RemoveTRANSLATE(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *SourceHTMLElement) TYPE(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *SourceHTMLElement) RemoveTYPE(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "type")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *SourceHTMLElement) WIDTH(v string) *SourceHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *SourceHTMLElement) RemoveWIDTH(v string) *SourceHTMLElement {
    delete(e.StringAttributes, "width")
    return e
}
