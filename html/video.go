package html

import (
	"fmt"
)

type VideoHTMLElement struct {
	*Element
}

func VIDEO(children ...ElementBuilder) *VideoHTMLElement {
	return &VideoHTMLElement{
		Element: &Element{
			Tag:           elementTagVIDEO,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *VideoHTMLElement) Children(children ...ElementBuilder) *VideoHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *VideoHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *VideoHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *VideoHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *VideoHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *VideoHTMLElement) Text(text string) *VideoHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *VideoHTMLElement) TextF(format string, args ...any) *VideoHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *VideoHTMLElement) Escaped(text string) *VideoHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *VideoHTMLElement) EscapedF(format string, args ...any) *VideoHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
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
//   - identical_to
//   - ordered_set_of_unique_space_separated_tokens
func (e *VideoHTMLElement) ACCESSKEY(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *VideoHTMLElement) IfACCESSKEY(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *VideoHTMLElement) RemoveACCESSKEY(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// AUTOCAPITALIZE sets the "autocapitalize" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Values values are constrained to:
//   - characters
//   - none
//   - off
//   - on
//   - sentences
//   - words
func (e *VideoHTMLElement) AUTOCAPITALIZE(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *VideoHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *VideoHTMLElement) RemoveAUTOCAPITALIZE(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *VideoHTMLElement) AUTOFOCUS() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *VideoHTMLElement) IfAUTOFOCUS(cond bool) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *VideoHTMLElement) RemoveAUTOFOCUS() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
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
//   - boolean_attribute
func (e *VideoHTMLElement) AUTOPLAY() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOPLAYKey] = struct{}{}
	return e
}

func (e *VideoHTMLElement) IfAUTOPLAY(cond bool) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOPLAY()
}

func (e *VideoHTMLElement) RemoveAUTOPLAY() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOPLAYKey)
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
//   - set_of_space_separated_tokens
func (e *VideoHTMLElement) CLASS(v string) *VideoHTMLElement {
	if e.DelimitedStringAttributes == nil {
		e.DelimitedStringAttributes = map[string]*DelimitedString{}
	}
	kv, ok := e.DelimitedStringAttributes[attributeCLASSKey]
	if !ok {
		kv = NewSpaceDelimitedString()
		e.DelimitedStringAttributes[attributeCLASSKey] = kv
	}
	kv.Add(v)
	return e
}

func (e *VideoHTMLElement) IfCLASS(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *VideoHTMLElement) SetCLASS(v string) *VideoHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *VideoHTMLElement) RemoveCLASS(v string) *VideoHTMLElement {
	kv, ok := e.DelimitedStringAttributes[attributeCLASSKey]
	if !ok {
		return e
	}
	kv.Remove(v)
	return e
}

// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//   - false
//   - plaintext_only
//   - true
func (e *VideoHTMLElement) CONTENTEDITABLE(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *VideoHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *VideoHTMLElement) RemoveCONTENTEDITABLE(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// CONTROLS sets the "controls" attribute.
// Show user agent controls
// Values values are constrained to:
//   - boolean_attribute
func (e *VideoHTMLElement) CONTROLS() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeCONTROLSKey] = struct{}{}
	return e
}

func (e *VideoHTMLElement) IfCONTROLS(cond bool) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.CONTROLS()
}

func (e *VideoHTMLElement) RemoveCONTROLS() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeCONTROLSKey)
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
//   - anonymous
//   - use_credentials
func (e *VideoHTMLElement) CROSSORIGIN(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCROSSORIGINKey] = v
	return e
}

func (e *VideoHTMLElement) IfCROSSORIGIN(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.CROSSORIGIN(v)
}

func (e *VideoHTMLElement) RemoveCROSSORIGIN(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeCROSSORIGINKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *VideoHTMLElement) DIR(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *VideoHTMLElement) IfDIR(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *VideoHTMLElement) RemoveDIR(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *VideoHTMLElement) DRAGGABLE(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *VideoHTMLElement) IfDRAGGABLE(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *VideoHTMLElement) RemoveDRAGGABLE(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeDRAGGABLEKey)
	return e
}

// ENTERKEYHINT sets the "enterkeyhint" attribute.
// Hint for selecting an enter key action
// Values values are constrained to:
//   - done
//   - enter
//   - go
//   - next
//   - previous
//   - search
//   - send
func (e *VideoHTMLElement) ENTERKEYHINT(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *VideoHTMLElement) IfENTERKEYHINT(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *VideoHTMLElement) RemoveENTERKEYHINT(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *VideoHTMLElement) HEIGHT(v int) *VideoHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeHEIGHTKey] = v
	return e
}

func (e *VideoHTMLElement) IfHEIGHT(cond bool, v int) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.HEIGHT(v)
}

func (e *VideoHTMLElement) RemoveHEIGHT(v int) *VideoHTMLElement {
	delete(e.IntAttributes, attributeHEIGHTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *VideoHTMLElement) HIDDEN(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *VideoHTMLElement) IfHIDDEN(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *VideoHTMLElement) RemoveHIDDEN(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *VideoHTMLElement) ID(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *VideoHTMLElement) IfID(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *VideoHTMLElement) RemoveID(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *VideoHTMLElement) INERT() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *VideoHTMLElement) IfINERT(cond bool) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *VideoHTMLElement) RemoveINERT() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
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
//   - decimal
//   - email
//   - none
//   - numeric
//   - search
//   - tel
//   - text
//   - url
func (e *VideoHTMLElement) INPUTMODE(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *VideoHTMLElement) IfINPUTMODE(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *VideoHTMLElement) RemoveINPUTMODE(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *VideoHTMLElement) IS(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *VideoHTMLElement) IfIS(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *VideoHTMLElement) RemoveIS(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *VideoHTMLElement) ITEMID(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *VideoHTMLElement) IfITEMID(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *VideoHTMLElement) RemoveITEMID(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *VideoHTMLElement) ITEMPROP(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *VideoHTMLElement) IfITEMPROP(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *VideoHTMLElement) RemoveITEMPROP(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *VideoHTMLElement) ITEMREF(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *VideoHTMLElement) IfITEMREF(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *VideoHTMLElement) RemoveITEMREF(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *VideoHTMLElement) ITEMSCOPE() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *VideoHTMLElement) IfITEMSCOPE(cond bool) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *VideoHTMLElement) RemoveITEMSCOPE() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
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
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *VideoHTMLElement) ITEMTYPE(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *VideoHTMLElement) IfITEMTYPE(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *VideoHTMLElement) RemoveITEMTYPE(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *VideoHTMLElement) LANG(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *VideoHTMLElement) IfLANG(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *VideoHTMLElement) RemoveLANG(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// LOOP sets the "loop" attribute.
// Whether to loop the media resource
// Values values are constrained to:
//   - boolean_attribute
func (e *VideoHTMLElement) LOOP() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeLOOPKey] = struct{}{}
	return e
}

func (e *VideoHTMLElement) IfLOOP(cond bool) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.LOOP()
}

func (e *VideoHTMLElement) RemoveLOOP() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeLOOPKey)
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
//   - boolean_attribute
func (e *VideoHTMLElement) MUTED() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeMUTEDKey] = struct{}{}
	return e
}

func (e *VideoHTMLElement) IfMUTED(cond bool) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.MUTED()
}

func (e *VideoHTMLElement) RemoveMUTED() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeMUTEDKey)
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
//   - text
func (e *VideoHTMLElement) NONCE(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *VideoHTMLElement) IfNONCE(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *VideoHTMLElement) RemoveNONCE(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// PLAYSINLINE sets the "playsinline" attribute.
// Encourage the user agent to display video content within the element's playback area
// Values values are constrained to:
//   - boolean_attribute
func (e *VideoHTMLElement) PLAYSINLINE() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributePLAYSINLINEKey] = struct{}{}
	return e
}

func (e *VideoHTMLElement) IfPLAYSINLINE(cond bool) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.PLAYSINLINE()
}

func (e *VideoHTMLElement) RemovePLAYSINLINE() *VideoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributePLAYSINLINEKey)
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
//   - auto
//   - manual
func (e *VideoHTMLElement) POPOVER(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *VideoHTMLElement) IfPOPOVER(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *VideoHTMLElement) RemovePOPOVER(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// POSTER sets the "poster" attribute.
// Poster frame to show prior to video playback
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *VideoHTMLElement) POSTER(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOSTERKey] = v
	return e
}

func (e *VideoHTMLElement) IfPOSTER(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.POSTER(v)
}

func (e *VideoHTMLElement) RemovePOSTER(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributePOSTERKey)
	return e
}

// PRELOAD sets the "preload" attribute.
// Hints how much buffering the media resource will likely need
// Values values are constrained to:
//   - auto
//   - metadata
//   - none
func (e *VideoHTMLElement) PRELOAD(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePRELOADKey] = v
	return e
}

func (e *VideoHTMLElement) IfPRELOAD(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.PRELOAD(v)
}

func (e *VideoHTMLElement) RemovePRELOAD(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributePRELOADKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *VideoHTMLElement) SLOT(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *VideoHTMLElement) IfSLOT(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *VideoHTMLElement) RemoveSLOT(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *VideoHTMLElement) SPELLCHECK(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *VideoHTMLElement) IfSPELLCHECK(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *VideoHTMLElement) RemoveSPELLCHECK(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *VideoHTMLElement) SRC(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCKey] = v
	return e
}

func (e *VideoHTMLElement) IfSRC(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.SRC(v)
}

func (e *VideoHTMLElement) RemoveSRC(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeSRCKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *VideoHTMLElement) STYLE(k, v string) *VideoHTMLElement {
	if e.DelimitedKVAttributes == nil {
		e.DelimitedKVAttributes = map[string]*DelimitedKVString{}
	}
	kv, ok := e.DelimitedKVAttributes[attributeSTYLEKey]
	if !ok {
		kv = NewEqualSemicolonDelimitedKVString()
		e.DelimitedKVAttributes[attributeSTYLEKey] = kv
	}
	kv.Add(k, v)
	return e
}

func (e *VideoHTMLElement) IfSTYLE(cond bool, k string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *VideoHTMLElement) RemoveSTYLE(k string) *VideoHTMLElement {
	kv, ok := e.DelimitedKVAttributes[attributeSTYLEKey]
	if !ok {
		return e
	}
	kv.Remove(k)
	return e
}

// TABINDEX sets the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and
//
//	the relative order of the element for the purposes of sequential focus navigation
//
// Values values are constrained to:
//   - valid_integer
func (e *VideoHTMLElement) TABINDEX(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *VideoHTMLElement) IfTABINDEX(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *VideoHTMLElement) RemoveTABINDEX(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *VideoHTMLElement) TITLE(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *VideoHTMLElement) IfTITLE(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *VideoHTMLElement) RemoveTITLE(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *VideoHTMLElement) TRANSLATE(v string) *VideoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *VideoHTMLElement) IfTRANSLATE(cond bool, v string) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *VideoHTMLElement) RemoveTRANSLATE(v string) *VideoHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *VideoHTMLElement) WIDTH(v int) *VideoHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeWIDTHKey] = v
	return e
}

func (e *VideoHTMLElement) IfWIDTH(cond bool, v int) *VideoHTMLElement {
	if !cond {
		return e
	}
	return e.WIDTH(v)
}

func (e *VideoHTMLElement) RemoveWIDTH(v int) *VideoHTMLElement {
	delete(e.IntAttributes, attributeWIDTHKey)
	return e
}
