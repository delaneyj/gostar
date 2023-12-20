package  html

import (
    "fmt"
)

type AudioHTMLElement struct {
    *Element
}

func AUDIO(children ...ElementBuilder) *AudioHTMLElement {
    return &AudioHTMLElement{
        Element: &Element{
            Tag: "audio",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *AudioHTMLElement) Children(children ...ElementBuilder) *AudioHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *AudioHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *AudioHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *AudioHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *AudioHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *AudioHTMLElement) Text(text string) *AudioHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *AudioHTMLElement) TextF(format string, args ...any) *AudioHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *AudioHTMLElement) Escaped(text string) *AudioHTMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *AudioHTMLElement) EscapedF(format string, args ...any) *AudioHTMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *AudioHTMLElement) CustomData(key, value string) *AudioHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AudioHTMLElement) CustomDataRemove(key string) *AudioHTMLElement {
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
func (e *AudioHTMLElement) ACCESSKEY(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *AudioHTMLElement) IfACCESSKEY(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.ACCESSKEY(v)
}

func (e *AudioHTMLElement) RemoveACCESSKEY(v string) *AudioHTMLElement {
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
func (e *AudioHTMLElement) AUTOCAPITALIZE(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *AudioHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCAPITALIZE(v)
}

func (e *AudioHTMLElement) RemoveAUTOCAPITALIZE(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioHTMLElement) AUTOFOCUS() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *AudioHTMLElement) IfAUTOFOCUS(cond bool) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOFOCUS()
}

func (e *AudioHTMLElement) RemoveAUTOFOCUS() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *AudioHTMLElement) SetAUTOFOCUS(b bool) *AudioHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// AUTOPLAY sets the "autoplay" attribute.
// Hint that the media resource can be started automatically when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioHTMLElement) AUTOPLAY() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autoplay"] = struct{}{}
    return e
}

func (e *AudioHTMLElement) IfAUTOPLAY(cond bool) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOPLAY()
}

func (e *AudioHTMLElement) RemoveAUTOPLAY() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autoplay")
    return e
}

func (e *AudioHTMLElement) SetAUTOPLAY(b bool) *AudioHTMLElement {
    if b {
        return e.AUTOPLAY()
    }
    return e.RemoveAUTOPLAY()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *AudioHTMLElement) CLASS(v string) *AudioHTMLElement {
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

func (e *AudioHTMLElement) IfCLASS(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.CLASS(v)
}

func (e *AudioHTMLElement) SetCLASS(v string) *AudioHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *AudioHTMLElement) RemoveCLASS(v string) *AudioHTMLElement {
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
func (e *AudioHTMLElement) CONTENTEDITABLE(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *AudioHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.CONTENTEDITABLE(v)
}

func (e *AudioHTMLElement) RemoveCONTENTEDITABLE(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// CONTROLS sets the "controls" attribute.
// Show user agent controls
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioHTMLElement) CONTROLS() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["controls"] = struct{}{}
    return e
}

func (e *AudioHTMLElement) IfCONTROLS(cond bool) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.CONTROLS()
}

func (e *AudioHTMLElement) RemoveCONTROLS() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "controls")
    return e
}

func (e *AudioHTMLElement) SetCONTROLS(b bool) *AudioHTMLElement {
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
func (e *AudioHTMLElement) CROSSORIGIN(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["crossorigin"] = v
    return e
}

func (e *AudioHTMLElement) IfCROSSORIGIN(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.CROSSORIGIN(v)
}

func (e *AudioHTMLElement) RemoveCROSSORIGIN(v string) *AudioHTMLElement {
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
func (e *AudioHTMLElement) DIR(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *AudioHTMLElement) IfDIR(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.DIR(v)
}

func (e *AudioHTMLElement) RemoveDIR(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *AudioHTMLElement) DRAGGABLE(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *AudioHTMLElement) IfDRAGGABLE(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.DRAGGABLE(v)
}

func (e *AudioHTMLElement) RemoveDRAGGABLE(v string) *AudioHTMLElement {
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
func (e *AudioHTMLElement) ENTERKEYHINT(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *AudioHTMLElement) IfENTERKEYHINT(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.ENTERKEYHINT(v)
}

func (e *AudioHTMLElement) RemoveENTERKEYHINT(v string) *AudioHTMLElement {
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
func (e *AudioHTMLElement) HIDDEN(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *AudioHTMLElement) IfHIDDEN(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.HIDDEN(v)
}

func (e *AudioHTMLElement) RemoveHIDDEN(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *AudioHTMLElement) ID(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *AudioHTMLElement) IfID(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.ID(v)
}

func (e *AudioHTMLElement) RemoveID(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioHTMLElement) INERT() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *AudioHTMLElement) IfINERT(cond bool) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.INERT()
}

func (e *AudioHTMLElement) RemoveINERT() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *AudioHTMLElement) SetINERT(b bool) *AudioHTMLElement {
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
func (e *AudioHTMLElement) INPUTMODE(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *AudioHTMLElement) IfINPUTMODE(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.INPUTMODE(v)
}

func (e *AudioHTMLElement) RemoveINPUTMODE(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *AudioHTMLElement) IS(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *AudioHTMLElement) IfIS(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.IS(v)
}

func (e *AudioHTMLElement) RemoveIS(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *AudioHTMLElement) ITEMID(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *AudioHTMLElement) IfITEMID(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMID(v)
}

func (e *AudioHTMLElement) RemoveITEMID(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *AudioHTMLElement) ITEMPROP(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *AudioHTMLElement) IfITEMPROP(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMPROP(v)
}

func (e *AudioHTMLElement) RemoveITEMPROP(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *AudioHTMLElement) ITEMREF(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *AudioHTMLElement) IfITEMREF(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMREF(v)
}

func (e *AudioHTMLElement) RemoveITEMREF(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioHTMLElement) ITEMSCOPE() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *AudioHTMLElement) IfITEMSCOPE(cond bool) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMSCOPE()
}

func (e *AudioHTMLElement) RemoveITEMSCOPE() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *AudioHTMLElement) SetITEMSCOPE(b bool) *AudioHTMLElement {
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
func (e *AudioHTMLElement) ITEMTYPE(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *AudioHTMLElement) IfITEMTYPE(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMTYPE(v)
}

func (e *AudioHTMLElement) RemoveITEMTYPE(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AudioHTMLElement) LANG(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *AudioHTMLElement) IfLANG(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.LANG(v)
}

func (e *AudioHTMLElement) RemoveLANG(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LOOP sets the "loop" attribute.
// Whether to loop the media resource
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioHTMLElement) LOOP() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["loop"] = struct{}{}
    return e
}

func (e *AudioHTMLElement) IfLOOP(cond bool) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.LOOP()
}

func (e *AudioHTMLElement) RemoveLOOP() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "loop")
    return e
}

func (e *AudioHTMLElement) SetLOOP(b bool) *AudioHTMLElement {
    if b {
        return e.LOOP()
    }
    return e.RemoveLOOP()
}
// MUTED sets the "muted" attribute.
// Whether to mute the media resource by default
// Values values are constrained to:
//  * boolean_attribute
func (e *AudioHTMLElement) MUTED() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["muted"] = struct{}{}
    return e
}

func (e *AudioHTMLElement) IfMUTED(cond bool) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.MUTED()
}

func (e *AudioHTMLElement) RemoveMUTED() *AudioHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "muted")
    return e
}

func (e *AudioHTMLElement) SetMUTED(b bool) *AudioHTMLElement {
    if b {
        return e.MUTED()
    }
    return e.RemoveMUTED()
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *AudioHTMLElement) NONCE(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *AudioHTMLElement) IfNONCE(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.NONCE(v)
}

func (e *AudioHTMLElement) RemoveNONCE(v string) *AudioHTMLElement {
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
func (e *AudioHTMLElement) POPOVER(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *AudioHTMLElement) IfPOPOVER(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.POPOVER(v)
}

func (e *AudioHTMLElement) RemovePOPOVER(v string) *AudioHTMLElement {
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
func (e *AudioHTMLElement) PRELOAD(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["preload"] = v
    return e
}

func (e *AudioHTMLElement) IfPRELOAD(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.PRELOAD(v)
}

func (e *AudioHTMLElement) RemovePRELOAD(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "preload")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *AudioHTMLElement) SLOT(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *AudioHTMLElement) IfSLOT(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.SLOT(v)
}

func (e *AudioHTMLElement) RemoveSLOT(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *AudioHTMLElement) SPELLCHECK(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *AudioHTMLElement) IfSPELLCHECK(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.SPELLCHECK(v)
}

func (e *AudioHTMLElement) RemoveSPELLCHECK(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *AudioHTMLElement) SRC(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *AudioHTMLElement) IfSRC(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.SRC(v)
}

func (e *AudioHTMLElement) RemoveSRC(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "src")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AudioHTMLElement) STYLE(k,v string) *AudioHTMLElement {
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

func (e *AudioHTMLElement) IfSTYLE(cond bool, k string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.STYLE(k, "")
}

func (e *AudioHTMLElement) RemoveSTYLE(k string) *AudioHTMLElement {
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
func (e *AudioHTMLElement) TABINDEX(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *AudioHTMLElement) IfTABINDEX(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.TABINDEX(v)
}

func (e *AudioHTMLElement) RemoveTABINDEX(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *AudioHTMLElement) TITLE(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *AudioHTMLElement) IfTITLE(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.TITLE(v)
}

func (e *AudioHTMLElement) RemoveTITLE(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *AudioHTMLElement) TRANSLATE(v string) *AudioHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *AudioHTMLElement) IfTRANSLATE(cond bool, v string) *AudioHTMLElement {
    if !cond {
        return e
    }
    return e.TRANSLATE(v)
}

func (e *AudioHTMLElement) RemoveTRANSLATE(v string) *AudioHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
