package  html

import (
    "fmt"
)

type VideoHTMLElement struct {
    *Element
}

func VIDEO(children ...ElementBuilder) *VideoHTMLElement {
    return &VideoHTMLElement{
        Element: &Element{
            Tag: "video",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *VideoHTMLElement) Children(children ...ElementBuilder) *VideoHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *VideoHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *VideoHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *VideoHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *VideoHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *VideoHTMLElement) Text(text string) *VideoHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *VideoHTMLElement) TextF(format string, args ...any) *VideoHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *VideoHTMLElement) Raw(text string) *VideoHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *VideoHTMLElement) RawF(format string, args ...any) *VideoHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *VideoHTMLElement) CustomData(key, value string) *VideoHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *VideoHTMLElement) CustomDataRemove(key string) *VideoHTMLElement {
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
func (e *VideoHTMLElement) ACCESSKEY(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *VideoHTMLElement) RemoveACCESSKEY(v string) *VideoHTMLElement {
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
func (e *VideoHTMLElement) AUTOCAPITALIZE(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *VideoHTMLElement) RemoveAUTOCAPITALIZE(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoHTMLElement) AUTOFOCUS() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *VideoHTMLElement) RemoveAUTOFOCUS() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *VideoHTMLElement) SetAUTOFOCUS(b bool) *VideoHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// AUTOPLAY sets the "autoplay" attribute.
// Hint that the media resource can be started automatically when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoHTMLElement) AUTOPLAY() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autoplay"] = struct{}{}
    return e
}

func (e *VideoHTMLElement) RemoveAUTOPLAY() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autoplay")
    return e
}

func (e *VideoHTMLElement) SetAUTOPLAY(b bool) *VideoHTMLElement {
    if b {
        return e.AUTOPLAY()
    }
    return e.RemoveAUTOPLAY()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *VideoHTMLElement) CLASS(v string) *VideoHTMLElement {
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

func (e *VideoHTMLElement) SetCLASS(v string) *VideoHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *VideoHTMLElement) RemoveCLASS(v string) *VideoHTMLElement {
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
func (e *VideoHTMLElement) CONTENTEDITABLE(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *VideoHTMLElement) RemoveCONTENTEDITABLE(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// CONTROLS sets the "controls" attribute.
// Show user agent controls
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoHTMLElement) CONTROLS() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["controls"] = struct{}{}
    return e
}

func (e *VideoHTMLElement) RemoveCONTROLS() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "controls")
    return e
}

func (e *VideoHTMLElement) SetCONTROLS(b bool) *VideoHTMLElement {
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
func (e *VideoHTMLElement) CROSSORIGIN(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *VideoHTMLElement) RemoveCROSSORIGIN(v string) *VideoHTMLElement {
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
func (e *VideoHTMLElement) DIR(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *VideoHTMLElement) RemoveDIR(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *VideoHTMLElement) DRAGGABLE(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *VideoHTMLElement) RemoveDRAGGABLE(v string) *VideoHTMLElement {
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
func (e *VideoHTMLElement) ENTERKEYHINT(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *VideoHTMLElement) RemoveENTERKEYHINT(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *VideoHTMLElement) HEIGHT(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *VideoHTMLElement) RemoveHEIGHT(v string) *VideoHTMLElement {
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
func (e *VideoHTMLElement) HIDDEN(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *VideoHTMLElement) RemoveHIDDEN(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *VideoHTMLElement) ID(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *VideoHTMLElement) RemoveID(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoHTMLElement) INERT() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *VideoHTMLElement) RemoveINERT() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *VideoHTMLElement) SetINERT(b bool) *VideoHTMLElement {
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
func (e *VideoHTMLElement) INPUTMODE(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *VideoHTMLElement) RemoveINPUTMODE(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *VideoHTMLElement) IS(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *VideoHTMLElement) RemoveIS(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *VideoHTMLElement) ITEMID(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *VideoHTMLElement) RemoveITEMID(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *VideoHTMLElement) ITEMPROP(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *VideoHTMLElement) RemoveITEMPROP(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *VideoHTMLElement) ITEMREF(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *VideoHTMLElement) RemoveITEMREF(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoHTMLElement) ITEMSCOPE() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *VideoHTMLElement) RemoveITEMSCOPE() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *VideoHTMLElement) SetITEMSCOPE(b bool) *VideoHTMLElement {
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
func (e *VideoHTMLElement) ITEMTYPE(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *VideoHTMLElement) RemoveITEMTYPE(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *VideoHTMLElement) LANG(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *VideoHTMLElement) RemoveLANG(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LOOP sets the "loop" attribute.
// Whether to loop the media resource
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoHTMLElement) LOOP() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["loop"] = struct{}{}
    return e
}

func (e *VideoHTMLElement) RemoveLOOP() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "loop")
    return e
}

func (e *VideoHTMLElement) SetLOOP(b bool) *VideoHTMLElement {
    if b {
        return e.LOOP()
    }
    return e.RemoveLOOP()
}
// MUTED sets the "muted" attribute.
// Whether to mute the media resource by default
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoHTMLElement) MUTED() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["muted"] = struct{}{}
    return e
}

func (e *VideoHTMLElement) RemoveMUTED() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "muted")
    return e
}

func (e *VideoHTMLElement) SetMUTED(b bool) *VideoHTMLElement {
    if b {
        return e.MUTED()
    }
    return e.RemoveMUTED()
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *VideoHTMLElement) NONCE(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *VideoHTMLElement) RemoveNONCE(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// PLAYSINLINE sets the "playsinline" attribute.
// Encourage the user agent to display video content within the element's playback area
// Values values are constrained to:
//  * boolean_attribute
func (e *VideoHTMLElement) PLAYSINLINE() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["playsinline"] = struct{}{}
    return e
}

func (e *VideoHTMLElement) RemovePLAYSINLINE() *VideoHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "playsinline")
    return e
}

func (e *VideoHTMLElement) SetPLAYSINLINE(b bool) *VideoHTMLElement {
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
func (e *VideoHTMLElement) POPOVER(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *VideoHTMLElement) RemovePOPOVER(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// POSTER sets the "poster" attribute.
// Poster frame to show prior to video playback
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *VideoHTMLElement) POSTER(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["poster"] = v
    return e
}

func (e *VideoHTMLElement) RemovePOSTER(v string) *VideoHTMLElement {
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
func (e *VideoHTMLElement) PRELOAD(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["preload"] = v
    return e
}

func (e *VideoHTMLElement) RemovePRELOAD(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "preload")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *VideoHTMLElement) SLOT(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *VideoHTMLElement) RemoveSLOT(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *VideoHTMLElement) SPELLCHECK(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *VideoHTMLElement) RemoveSPELLCHECK(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *VideoHTMLElement) SRC(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *VideoHTMLElement) RemoveSRC(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "src")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *VideoHTMLElement) STYLE(k,v string) *VideoHTMLElement {
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

func (e *VideoHTMLElement) RemoveSTYLE(k string) *VideoHTMLElement {
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
func (e *VideoHTMLElement) TABINDEX(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *VideoHTMLElement) RemoveTABINDEX(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *VideoHTMLElement) TITLE(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *VideoHTMLElement) RemoveTITLE(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *VideoHTMLElement) TRANSLATE(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *VideoHTMLElement) RemoveTRANSLATE(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *VideoHTMLElement) WIDTH(v string) *VideoHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *VideoHTMLElement) RemoveWIDTH(v string) *VideoHTMLElement {
    delete(e.StringAttributes, "width")
    return e
}
