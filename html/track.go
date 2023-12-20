package  html

import (
    "fmt"
)

type TrackHTMLElement struct {
    *Element
}

func TRACK(children ...ElementBuilder) *TrackHTMLElement {
    return &TrackHTMLElement{
        Element: &Element{
            Tag: "track",
            IsSelfClosing: true,
            Descendants: children,
        },
    }
}

func (e *TrackHTMLElement) Children(children ...ElementBuilder) *TrackHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *TrackHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TrackHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *TrackHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TrackHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *TrackHTMLElement) Text(text string) *TrackHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *TrackHTMLElement) TextF(format string, args ...any) *TrackHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *TrackHTMLElement) Raw(text string) *TrackHTMLElement {
    e.Descendants = append(e.Descendants, Raw(text))
    return e
}

func (e *TrackHTMLElement) RawF(format string, args ...any) *TrackHTMLElement {
    return e.Raw(fmt.Sprintf(format, args...))
}

func (e *TrackHTMLElement) CustomData(key, value string) *TrackHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TrackHTMLElement) CustomDataRemove(key string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) ACCESSKEY(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *TrackHTMLElement) RemoveACCESSKEY(v string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) AUTOCAPITALIZE(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *TrackHTMLElement) RemoveAUTOCAPITALIZE(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *TrackHTMLElement) AUTOFOCUS() *TrackHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *TrackHTMLElement) RemoveAUTOFOCUS() *TrackHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *TrackHTMLElement) SetAUTOFOCUS(b bool) *TrackHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *TrackHTMLElement) CLASS(v string) *TrackHTMLElement {
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

func (e *TrackHTMLElement) SetCLASS(v string) *TrackHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *TrackHTMLElement) RemoveCLASS(v string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) CONTENTEDITABLE(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *TrackHTMLElement) RemoveCONTENTEDITABLE(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// DEFAULT sets the "default" attribute.
// Enable the track if no other text track is more suitable
// Values values are constrained to:
//  * boolean_attribute
func (e *TrackHTMLElement) DEFAULT() *TrackHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["default"] = struct{}{}
    return e
}

func (e *TrackHTMLElement) RemoveDEFAULT() *TrackHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "default")
    return e
}

func (e *TrackHTMLElement) SetDEFAULT(b bool) *TrackHTMLElement {
    if b {
        return e.DEFAULT()
    }
    return e.RemoveDEFAULT()
}
// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//  * ltr
//  * ltr
//  * rtl
//  * rtl
func (e *TrackHTMLElement) DIR(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *TrackHTMLElement) RemoveDIR(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *TrackHTMLElement) DRAGGABLE(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *TrackHTMLElement) RemoveDRAGGABLE(v string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) ENTERKEYHINT(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *TrackHTMLElement) RemoveENTERKEYHINT(v string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) HIDDEN(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *TrackHTMLElement) RemoveHIDDEN(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *TrackHTMLElement) ID(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *TrackHTMLElement) RemoveID(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *TrackHTMLElement) INERT() *TrackHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *TrackHTMLElement) RemoveINERT() *TrackHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *TrackHTMLElement) SetINERT(b bool) *TrackHTMLElement {
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
func (e *TrackHTMLElement) INPUTMODE(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *TrackHTMLElement) RemoveINPUTMODE(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *TrackHTMLElement) IS(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *TrackHTMLElement) RemoveIS(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *TrackHTMLElement) ITEMID(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *TrackHTMLElement) RemoveITEMID(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *TrackHTMLElement) ITEMPROP(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *TrackHTMLElement) RemoveITEMPROP(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *TrackHTMLElement) ITEMREF(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *TrackHTMLElement) RemoveITEMREF(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *TrackHTMLElement) ITEMSCOPE() *TrackHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *TrackHTMLElement) RemoveITEMSCOPE() *TrackHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *TrackHTMLElement) SetITEMSCOPE(b bool) *TrackHTMLElement {
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
func (e *TrackHTMLElement) ITEMTYPE(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *TrackHTMLElement) RemoveITEMTYPE(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// KIND sets the "kind" attribute.
// The type of text track
// Values values are constrained to:
//  * subtitles
//  * subtitles
//  * captions
//  * captions
//  * descriptions
//  * descriptions
//  * chapters
//  * chapters
//  * metadata
//  * metadata
func (e *TrackHTMLElement) KIND(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["kind"] = v
    return e
}

func (e *TrackHTMLElement) RemoveKIND(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "kind")
    return e
}
// LABEL sets the "label" attribute.
// User-visible label
// Values values are constrained to:
//  * text
func (e *TrackHTMLElement) LABEL(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["label"] = v
    return e
}

func (e *TrackHTMLElement) RemoveLABEL(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "label")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TrackHTMLElement) LANG(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *TrackHTMLElement) RemoveLANG(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *TrackHTMLElement) NONCE(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *TrackHTMLElement) RemoveNONCE(v string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) POPOVER(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *TrackHTMLElement) RemovePOPOVER(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *TrackHTMLElement) SLOT(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *TrackHTMLElement) RemoveSLOT(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *TrackHTMLElement) SPELLCHECK(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *TrackHTMLElement) RemoveSPELLCHECK(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *TrackHTMLElement) SRC(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *TrackHTMLElement) RemoveSRC(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "src")
    return e
}
// SRCLANG sets the "srclang" attribute.
// Language of the text track
// Values values are constrained to:
func (e *TrackHTMLElement) SRCLANG(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["srclang"] = v
    return e
}

func (e *TrackHTMLElement) RemoveSRCLANG(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "srclang")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TrackHTMLElement) STYLE(k,v string) *TrackHTMLElement {
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

func (e *TrackHTMLElement) RemoveSTYLE(k string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) TABINDEX(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *TrackHTMLElement) RemoveTABINDEX(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *TrackHTMLElement) TITLE(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *TrackHTMLElement) RemoveTITLE(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *TrackHTMLElement) TRANSLATE(v string) *TrackHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *TrackHTMLElement) RemoveTRANSLATE(v string) *TrackHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
