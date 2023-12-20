package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type VideoElement struct {
    *gostar.Element
}

func VIDEO(children ...fmt.Stringer) *VideoElement {
    return &VideoElement{
        Element: &gostar.Element{
            Tag: "video",
            IsSelfClosing: false,
            Children: children,
        },
    }
}

func (e *VideoElement) AddChildren(children ...fmt.Stringer) *VideoElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *VideoElement) TEXT(text string) *VideoElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *VideoElement) RAW(text string) *VideoElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *VideoElement) CustomData(key, value string) *VideoElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *VideoElement) CustomDataRemove(key string) *VideoElement {
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
func (e *VideoElement) ACCESSKEY(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *VideoElement) RemoveACCESSKEY(v string) *VideoElement {
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
func (e *VideoElement) AUTOCAPITALIZE(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *VideoElement) RemoveAUTOCAPITALIZE(v string) *VideoElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoElement) AUTOFOCUS() *VideoElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *VideoElement) RemoveAUTOFOCUS() *VideoElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *VideoElement) SetAUTOFOCUS(b bool) *VideoElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// AUTOPLAY sets the "autoplay" attribute.
// Hint that the media resource can be started automatically when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoElement) AUTOPLAY() *VideoElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autoplay"] = struct{}{}
    return e
}

func (e *VideoElement) RemoveAUTOPLAY() *VideoElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autoplay")
    return e
}

func (e *VideoElement) SetAUTOPLAY(b bool) *VideoElement {
    if b {
        return e.AUTOPLAY()
    }
    return e.RemoveAUTOPLAY()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *VideoElement) CLASS(v string) *VideoElement {
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

func (e *VideoElement) SetCLASS(v string) *VideoElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *VideoElement) RemoveCLASS(v string) *VideoElement {
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
func (e *VideoElement) CONTENTEDITABLE(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *VideoElement) RemoveCONTENTEDITABLE(v string) *VideoElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// CONTROLS sets the "controls" attribute.
// Show user agent controls
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoElement) CONTROLS() *VideoElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["controls"] = struct{}{}
    return e
}

func (e *VideoElement) RemoveCONTROLS() *VideoElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "controls")
    return e
}

func (e *VideoElement) SetCONTROLS(b bool) *VideoElement {
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
func (e *VideoElement) CROSSORIGIN(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *VideoElement) RemoveCROSSORIGIN(v string) *VideoElement {
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
func (e *VideoElement) DIR(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *VideoElement) RemoveDIR(v string) *VideoElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *VideoElement) DRAGGABLE(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *VideoElement) RemoveDRAGGABLE(v string) *VideoElement {
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
func (e *VideoElement) ENTERKEYHINT(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *VideoElement) RemoveENTERKEYHINT(v string) *VideoElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *VideoElement) HEIGHT(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *VideoElement) RemoveHEIGHT(v string) *VideoElement {
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
func (e *VideoElement) HIDDEN(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *VideoElement) RemoveHIDDEN(v string) *VideoElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *VideoElement) ID(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *VideoElement) RemoveID(v string) *VideoElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoElement) INERT() *VideoElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *VideoElement) RemoveINERT() *VideoElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *VideoElement) SetINERT(b bool) *VideoElement {
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
func (e *VideoElement) INPUTMODE(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *VideoElement) RemoveINPUTMODE(v string) *VideoElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *VideoElement) IS(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *VideoElement) RemoveIS(v string) *VideoElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *VideoElement) ITEMID(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *VideoElement) RemoveITEMID(v string) *VideoElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *VideoElement) ITEMPROP(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *VideoElement) RemoveITEMPROP(v string) *VideoElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *VideoElement) ITEMREF(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *VideoElement) RemoveITEMREF(v string) *VideoElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoElement) ITEMSCOPE() *VideoElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *VideoElement) RemoveITEMSCOPE() *VideoElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *VideoElement) SetITEMSCOPE(b bool) *VideoElement {
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
func (e *VideoElement) ITEMTYPE(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *VideoElement) RemoveITEMTYPE(v string) *VideoElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *VideoElement) LANG(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *VideoElement) RemoveLANG(v string) *VideoElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LOOP sets the "loop" attribute.
// Whether to loop the media resource
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoElement) LOOP() *VideoElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["loop"] = struct{}{}
    return e
}

func (e *VideoElement) RemoveLOOP() *VideoElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "loop")
    return e
}

func (e *VideoElement) SetLOOP(b bool) *VideoElement {
    if b {
        return e.LOOP()
    }
    return e.RemoveLOOP()
}
// MUTED sets the "muted" attribute.
// Whether to mute the media resource by default
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoElement) MUTED() *VideoElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["muted"] = struct{}{}
    return e
}

func (e *VideoElement) RemoveMUTED() *VideoElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "muted")
    return e
}

func (e *VideoElement) SetMUTED(b bool) *VideoElement {
    if b {
        return e.MUTED()
    }
    return e.RemoveMUTED()
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *VideoElement) NONCE(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *VideoElement) RemoveNONCE(v string) *VideoElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// PLAYSINLINE sets the "playsinline" attribute.
// Encourage the user agent to display video content within the element's playback area
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoElement) PLAYSINLINE() *VideoElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["playsinline"] = struct{}{}
    return e
}

func (e *VideoElement) RemovePLAYSINLINE() *VideoElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "playsinline")
    return e
}

func (e *VideoElement) SetPLAYSINLINE(b bool) *VideoElement {
    if b {
        return e.PLAYSINLINE()
    }
    return e.RemovePLAYSINLINE()
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *VideoElement) POPOVER(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *VideoElement) RemovePOPOVER(v string) *VideoElement {
    delete(e.StringAttributes, "popover")
    return e
}
// POSTER sets the "poster" attribute.
// Poster frame to show prior to video playback
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *VideoElement) POSTER(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["poster"] = v
    return e
}

func (e *VideoElement) RemovePOSTER(v string) *VideoElement {
    delete(e.StringAttributes, "poster")
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
func (e *VideoElement) PRELOAD(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["preload"] = v
    return e
}

func (e *VideoElement) RemovePRELOAD(v string) *VideoElement {
    delete(e.StringAttributes, "preload")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *VideoElement) SLOT(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *VideoElement) RemoveSLOT(v string) *VideoElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *VideoElement) SPELLCHECK(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *VideoElement) RemoveSPELLCHECK(v string) *VideoElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *VideoElement) SRC(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *VideoElement) RemoveSRC(v string) *VideoElement {
    delete(e.StringAttributes, "src")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *VideoElement) STYLE(k,v string) *VideoElement {
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

func (e *VideoElement) RemoveSTYLE(k string) *VideoElement {
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
func (e *VideoElement) TABINDEX(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *VideoElement) RemoveTABINDEX(v string) *VideoElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *VideoElement) TITLE(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *VideoElement) RemoveTITLE(v string) *VideoElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *VideoElement) TRANSLATE(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *VideoElement) RemoveTRANSLATE(v string) *VideoElement {
    delete(e.StringAttributes, "translate")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *VideoElement) WIDTH(v string) *VideoElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *VideoElement) RemoveWIDTH(v string) *VideoElement {
    delete(e.StringAttributes, "width")
    return e
}
