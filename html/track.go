package html

import (
	"fmt"
)

type TrackHTMLElement struct {
	*Element
}

func TRACK(children ...ElementBuilder) *TrackHTMLElement {
	return &TrackHTMLElement{
		Element: &Element{
			Tag:           elementTagTRACK,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *TrackHTMLElement) Children(children ...ElementBuilder) *TrackHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *TrackHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TrackHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *TrackHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TrackHTMLElement {
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

func (e *TrackHTMLElement) Escaped(text string) *TrackHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *TrackHTMLElement) EscapedF(format string, args ...any) *TrackHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
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
//   - identical_to
//   - ordered_set_of_unique_space_separated_tokens
func (e *TrackHTMLElement) ACCESSKEY(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *TrackHTMLElement) IfACCESSKEY(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *TrackHTMLElement) RemoveACCESSKEY(v string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) AUTOCAPITALIZE(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *TrackHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *TrackHTMLElement) RemoveAUTOCAPITALIZE(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *TrackHTMLElement) AUTOFOCUS() *TrackHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *TrackHTMLElement) IfAUTOFOCUS(cond bool) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *TrackHTMLElement) RemoveAUTOFOCUS() *TrackHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
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
//   - set_of_space_separated_tokens
func (e *TrackHTMLElement) CLASS(v string) *TrackHTMLElement {
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

func (e *TrackHTMLElement) IfCLASS(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *TrackHTMLElement) SetCLASS(v string) *TrackHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *TrackHTMLElement) RemoveCLASS(v string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) CONTENTEDITABLE(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *TrackHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *TrackHTMLElement) RemoveCONTENTEDITABLE(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DEFAULT sets the "default" attribute.
// Enable the track if no other text track is more suitable
// Values values are constrained to:
//   - boolean_attribute
func (e *TrackHTMLElement) DEFAULT() *TrackHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeDEFAULTKey] = struct{}{}
	return e
}

func (e *TrackHTMLElement) IfDEFAULT(cond bool) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.DEFAULT()
}

func (e *TrackHTMLElement) RemoveDEFAULT() *TrackHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeDEFAULTKey)
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
//   - ltr
//   - rtl
func (e *TrackHTMLElement) DIR(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *TrackHTMLElement) IfDIR(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *TrackHTMLElement) RemoveDIR(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *TrackHTMLElement) DRAGGABLE(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *TrackHTMLElement) IfDRAGGABLE(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *TrackHTMLElement) RemoveDRAGGABLE(v string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) ENTERKEYHINT(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *TrackHTMLElement) IfENTERKEYHINT(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *TrackHTMLElement) RemoveENTERKEYHINT(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *TrackHTMLElement) HIDDEN(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *TrackHTMLElement) IfHIDDEN(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *TrackHTMLElement) RemoveHIDDEN(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *TrackHTMLElement) ID(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *TrackHTMLElement) IfID(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *TrackHTMLElement) RemoveID(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *TrackHTMLElement) INERT() *TrackHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *TrackHTMLElement) IfINERT(cond bool) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *TrackHTMLElement) RemoveINERT() *TrackHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
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
//   - decimal
//   - email
//   - none
//   - numeric
//   - search
//   - tel
//   - text
//   - url
func (e *TrackHTMLElement) INPUTMODE(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *TrackHTMLElement) IfINPUTMODE(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *TrackHTMLElement) RemoveINPUTMODE(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *TrackHTMLElement) IS(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *TrackHTMLElement) IfIS(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *TrackHTMLElement) RemoveIS(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *TrackHTMLElement) ITEMID(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *TrackHTMLElement) IfITEMID(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *TrackHTMLElement) RemoveITEMID(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *TrackHTMLElement) ITEMPROP(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *TrackHTMLElement) IfITEMPROP(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *TrackHTMLElement) RemoveITEMPROP(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TrackHTMLElement) ITEMREF(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *TrackHTMLElement) IfITEMREF(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *TrackHTMLElement) RemoveITEMREF(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *TrackHTMLElement) ITEMSCOPE() *TrackHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *TrackHTMLElement) IfITEMSCOPE(cond bool) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *TrackHTMLElement) RemoveITEMSCOPE() *TrackHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
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
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *TrackHTMLElement) ITEMTYPE(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *TrackHTMLElement) IfITEMTYPE(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *TrackHTMLElement) RemoveITEMTYPE(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// KIND sets the "kind" attribute.
// The type of text track
// Values values are constrained to:
//   - captions
//   - chapters
//   - descriptions
//   - metadata
//   - subtitles
func (e *TrackHTMLElement) KIND(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeKINDKey] = v
	return e
}

func (e *TrackHTMLElement) IfKIND(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.KIND(v)
}

func (e *TrackHTMLElement) RemoveKIND(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeKINDKey)
	return e
}

// LABEL sets the "label" attribute.
// User-visible label
// Values values are constrained to:
//   - text
func (e *TrackHTMLElement) LABEL(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLABELKey] = v
	return e
}

func (e *TrackHTMLElement) IfLABEL(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.LABEL(v)
}

func (e *TrackHTMLElement) RemoveLABEL(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeLABELKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TrackHTMLElement) LANG(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *TrackHTMLElement) IfLANG(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *TrackHTMLElement) RemoveLANG(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *TrackHTMLElement) NONCE(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *TrackHTMLElement) IfNONCE(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *TrackHTMLElement) RemoveNONCE(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *TrackHTMLElement) POPOVER(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *TrackHTMLElement) IfPOPOVER(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *TrackHTMLElement) RemovePOPOVER(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *TrackHTMLElement) SLOT(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *TrackHTMLElement) IfSLOT(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *TrackHTMLElement) RemoveSLOT(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *TrackHTMLElement) SPELLCHECK(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *TrackHTMLElement) IfSPELLCHECK(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *TrackHTMLElement) RemoveSPELLCHECK(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *TrackHTMLElement) SRC(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCKey] = v
	return e
}

func (e *TrackHTMLElement) IfSRC(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.SRC(v)
}

func (e *TrackHTMLElement) RemoveSRC(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeSRCKey)
	return e
}

// SRCLANG sets the "srclang" attribute.
// Language of the text track
// Values values are constrained to:
func (e *TrackHTMLElement) SRCLANG(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCLANGKey] = v
	return e
}

func (e *TrackHTMLElement) IfSRCLANG(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.SRCLANG(v)
}

func (e *TrackHTMLElement) RemoveSRCLANG(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeSRCLANGKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TrackHTMLElement) STYLE(k, v string) *TrackHTMLElement {
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

func (e *TrackHTMLElement) IfSTYLE(cond bool, k string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *TrackHTMLElement) RemoveSTYLE(k string) *TrackHTMLElement {
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
func (e *TrackHTMLElement) TABINDEX(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *TrackHTMLElement) IfTABINDEX(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *TrackHTMLElement) RemoveTABINDEX(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *TrackHTMLElement) TITLE(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *TrackHTMLElement) IfTITLE(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *TrackHTMLElement) RemoveTITLE(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *TrackHTMLElement) TRANSLATE(v string) *TrackHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *TrackHTMLElement) IfTRANSLATE(cond bool, v string) *TrackHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *TrackHTMLElement) RemoveTRANSLATE(v string) *TrackHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
