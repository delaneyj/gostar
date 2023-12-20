package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type TrackElement struct {
    *gostar.Element
}

func TRACK(children ...fmt.Stringer) *TrackElement {
    return &TrackElement{
        Element: &gostar.Element{
            Tag: "track",
            IsSelfClosing: true,
            Children: children,
        },
    }
}

func (e *TrackElement) AddChildren(children ...fmt.Stringer) *TrackElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *TrackElement) TEXT(text string) *TrackElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *TrackElement) RAW(text string) *TrackElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *TrackElement) CustomData(key, value string) *TrackElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TrackElement) CustomDataRemove(key string) *TrackElement {
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
func (e *TrackElement) ACCESSKEY(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *TrackElement) RemoveACCESSKEY(v string) *TrackElement {
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
func (e *TrackElement) AUTOCAPITALIZE(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *TrackElement) RemoveAUTOCAPITALIZE(v string) *TrackElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *TrackElement) AUTOFOCUS() *TrackElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *TrackElement) RemoveAUTOFOCUS() *TrackElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *TrackElement) SetAUTOFOCUS(b bool) *TrackElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *TrackElement) CLASS(v string) *TrackElement {
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

func (e *TrackElement) SetCLASS(v string) *TrackElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *TrackElement) RemoveCLASS(v string) *TrackElement {
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
func (e *TrackElement) CONTENTEDITABLE(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *TrackElement) RemoveCONTENTEDITABLE(v string) *TrackElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// DEFAULT sets the "default" attribute.
// Enable the track if no other text track is more suitable
// Values values are constrained to:
//  * boolean_attribute
func (e *TrackElement) DEFAULT() *TrackElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["default"] = struct{}{}
    return e
}

func (e *TrackElement) RemoveDEFAULT() *TrackElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "default")
    return e
}

func (e *TrackElement) SetDEFAULT(b bool) *TrackElement {
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
func (e *TrackElement) DIR(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *TrackElement) RemoveDIR(v string) *TrackElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *TrackElement) DRAGGABLE(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *TrackElement) RemoveDRAGGABLE(v string) *TrackElement {
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
func (e *TrackElement) ENTERKEYHINT(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *TrackElement) RemoveENTERKEYHINT(v string) *TrackElement {
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
func (e *TrackElement) HIDDEN(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *TrackElement) RemoveHIDDEN(v string) *TrackElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *TrackElement) ID(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *TrackElement) RemoveID(v string) *TrackElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *TrackElement) INERT() *TrackElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *TrackElement) RemoveINERT() *TrackElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *TrackElement) SetINERT(b bool) *TrackElement {
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
func (e *TrackElement) INPUTMODE(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *TrackElement) RemoveINPUTMODE(v string) *TrackElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *TrackElement) IS(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *TrackElement) RemoveIS(v string) *TrackElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *TrackElement) ITEMID(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *TrackElement) RemoveITEMID(v string) *TrackElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *TrackElement) ITEMPROP(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *TrackElement) RemoveITEMPROP(v string) *TrackElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *TrackElement) ITEMREF(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *TrackElement) RemoveITEMREF(v string) *TrackElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *TrackElement) ITEMSCOPE() *TrackElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *TrackElement) RemoveITEMSCOPE() *TrackElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *TrackElement) SetITEMSCOPE(b bool) *TrackElement {
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
func (e *TrackElement) ITEMTYPE(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *TrackElement) RemoveITEMTYPE(v string) *TrackElement {
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
func (e *TrackElement) KIND(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["kind"] = v
    return e
}

func (e *TrackElement) RemoveKIND(v string) *TrackElement {
    delete(e.StringAttributes, "kind")
    return e
}
// LABEL sets the "label" attribute.
// User-visible label
// Values values are constrained to:
//  * text
func (e *TrackElement) LABEL(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["label"] = v
    return e
}

func (e *TrackElement) RemoveLABEL(v string) *TrackElement {
    delete(e.StringAttributes, "label")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TrackElement) LANG(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *TrackElement) RemoveLANG(v string) *TrackElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *TrackElement) NONCE(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *TrackElement) RemoveNONCE(v string) *TrackElement {
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
func (e *TrackElement) POPOVER(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *TrackElement) RemovePOPOVER(v string) *TrackElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *TrackElement) SLOT(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *TrackElement) RemoveSLOT(v string) *TrackElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *TrackElement) SPELLCHECK(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *TrackElement) RemoveSPELLCHECK(v string) *TrackElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *TrackElement) SRC(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *TrackElement) RemoveSRC(v string) *TrackElement {
    delete(e.StringAttributes, "src")
    return e
}
// SRCLANG sets the "srclang" attribute.
// Language of the text track
// Values values are constrained to:
func (e *TrackElement) SRCLANG(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["srclang"] = v
    return e
}

func (e *TrackElement) RemoveSRCLANG(v string) *TrackElement {
    delete(e.StringAttributes, "srclang")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TrackElement) STYLE(k,v string) *TrackElement {
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

func (e *TrackElement) RemoveSTYLE(k string) *TrackElement {
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
func (e *TrackElement) TABINDEX(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *TrackElement) RemoveTABINDEX(v string) *TrackElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *TrackElement) TITLE(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *TrackElement) RemoveTITLE(v string) *TrackElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *TrackElement) TRANSLATE(v string) *TrackElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *TrackElement) RemoveTRANSLATE(v string) *TrackElement {
    delete(e.StringAttributes, "translate")
    return e
}
