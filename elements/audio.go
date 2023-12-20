package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type AudioElement struct {
    *gostar.Element
}

func AUDIO(children ...fmt.Stringer) *AudioElement {
    return &AudioElement{
        Element: &gostar.Element{
            Tag: "audio",
            IsSelfClosing: false,
            Children: children,
        },
    }
}

func (e *AudioElement) AddChildren(children ...fmt.Stringer) *AudioElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *AudioElement) TEXT(text string) *AudioElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *AudioElement) RAW(text string) *AudioElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *AudioElement) CustomData(key, value string) *AudioElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AudioElement) CustomDataRemove(key string) *AudioElement {
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
func (e *AudioElement) ACCESSKEY(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *AudioElement) RemoveACCESSKEY(v string) *AudioElement {
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
func (e *AudioElement) AUTOCAPITALIZE(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *AudioElement) RemoveAUTOCAPITALIZE(v string) *AudioElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioElement) AUTOFOCUS() *AudioElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *AudioElement) RemoveAUTOFOCUS() *AudioElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *AudioElement) SetAUTOFOCUS(b bool) *AudioElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// AUTOPLAY sets the "autoplay" attribute.
// Hint that the media resource can be started automatically when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioElement) AUTOPLAY() *AudioElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autoplay"] = struct{}{}
    return e
}

func (e *AudioElement) RemoveAUTOPLAY() *AudioElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autoplay")
    return e
}

func (e *AudioElement) SetAUTOPLAY(b bool) *AudioElement {
    if b {
        return e.AUTOPLAY()
    }
    return e.RemoveAUTOPLAY()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *AudioElement) CLASS(v string) *AudioElement {
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

func (e *AudioElement) SetCLASS(v string) *AudioElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *AudioElement) RemoveCLASS(v string) *AudioElement {
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
func (e *AudioElement) CONTENTEDITABLE(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *AudioElement) RemoveCONTENTEDITABLE(v string) *AudioElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// CONTROLS sets the "controls" attribute.
// Show user agent controls
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioElement) CONTROLS() *AudioElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["controls"] = struct{}{}
    return e
}

func (e *AudioElement) RemoveCONTROLS() *AudioElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "controls")
    return e
}

func (e *AudioElement) SetCONTROLS(b bool) *AudioElement {
    if b {
        return e.CONTROLS()
    }
    return e.RemoveCONTROLS()
}
// CROSSORIGIN sets the "crossorigin" attribute.
// How the element handles crossorigin requests
// Values values are constrained to:
//  * anonymous
//  * anonymous
//  * use_credentials
//  * use_credentials
func (e *AudioElement) CROSSORIGIN(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *AudioElement) RemoveCROSSORIGIN(v string) *AudioElement {
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
func (e *AudioElement) DIR(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *AudioElement) RemoveDIR(v string) *AudioElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *AudioElement) DRAGGABLE(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *AudioElement) RemoveDRAGGABLE(v string) *AudioElement {
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
func (e *AudioElement) ENTERKEYHINT(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *AudioElement) RemoveENTERKEYHINT(v string) *AudioElement {
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
func (e *AudioElement) HIDDEN(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *AudioElement) RemoveHIDDEN(v string) *AudioElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *AudioElement) ID(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *AudioElement) RemoveID(v string) *AudioElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioElement) INERT() *AudioElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *AudioElement) RemoveINERT() *AudioElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *AudioElement) SetINERT(b bool) *AudioElement {
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
func (e *AudioElement) INPUTMODE(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *AudioElement) RemoveINPUTMODE(v string) *AudioElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *AudioElement) IS(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *AudioElement) RemoveIS(v string) *AudioElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *AudioElement) ITEMID(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *AudioElement) RemoveITEMID(v string) *AudioElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *AudioElement) ITEMPROP(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *AudioElement) RemoveITEMPROP(v string) *AudioElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *AudioElement) ITEMREF(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *AudioElement) RemoveITEMREF(v string) *AudioElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioElement) ITEMSCOPE() *AudioElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *AudioElement) RemoveITEMSCOPE() *AudioElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *AudioElement) SetITEMSCOPE(b bool) *AudioElement {
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
func (e *AudioElement) ITEMTYPE(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *AudioElement) RemoveITEMTYPE(v string) *AudioElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AudioElement) LANG(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *AudioElement) RemoveLANG(v string) *AudioElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LOOP sets the "loop" attribute.
// Whether to loop the media resource
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioElement) LOOP() *AudioElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["loop"] = struct{}{}
    return e
}

func (e *AudioElement) RemoveLOOP() *AudioElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "loop")
    return e
}

func (e *AudioElement) SetLOOP(b bool) *AudioElement {
    if b {
        return e.LOOP()
    }
    return e.RemoveLOOP()
}
// MUTED sets the "muted" attribute.
// Whether to mute the media resource by default
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioElement) MUTED() *AudioElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["muted"] = struct{}{}
    return e
}

func (e *AudioElement) RemoveMUTED() *AudioElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "muted")
    return e
}

func (e *AudioElement) SetMUTED(b bool) *AudioElement {
    if b {
        return e.MUTED()
    }
    return e.RemoveMUTED()
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *AudioElement) NONCE(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *AudioElement) RemoveNONCE(v string) *AudioElement {
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
func (e *AudioElement) POPOVER(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *AudioElement) RemovePOPOVER(v string) *AudioElement {
    delete(e.StringAttributes, "popover")
    return e
}
// PRELOAD sets the "preload" attribute.
// Hints how much buffering the media resource will likely need
// Values values are constrained to:
//  * none
//  * none
//  * metadata
//  * metadata
//  * auto
//  * auto
func (e *AudioElement) PRELOAD(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["preload"] = v
    return e
}

func (e *AudioElement) RemovePRELOAD(v string) *AudioElement {
    delete(e.StringAttributes, "preload")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *AudioElement) SLOT(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *AudioElement) RemoveSLOT(v string) *AudioElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *AudioElement) SPELLCHECK(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *AudioElement) RemoveSPELLCHECK(v string) *AudioElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *AudioElement) SRC(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *AudioElement) RemoveSRC(v string) *AudioElement {
    delete(e.StringAttributes, "src")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AudioElement) STYLE(k,v string) *AudioElement {
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

func (e *AudioElement) RemoveSTYLE(k string) *AudioElement {
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
func (e *AudioElement) TABINDEX(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *AudioElement) RemoveTABINDEX(v string) *AudioElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *AudioElement) TITLE(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *AudioElement) RemoveTITLE(v string) *AudioElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *AudioElement) TRANSLATE(v string) *AudioElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *AudioElement) RemoveTRANSLATE(v string) *AudioElement {
    delete(e.StringAttributes, "translate")
    return e
}
