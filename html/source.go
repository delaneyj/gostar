package html

import (
	"fmt"
)

type SourceHTMLElement struct {
	*Element
}

func SOURCE(children ...ElementBuilder) *SourceHTMLElement {
	return &SourceHTMLElement{
		Element: &Element{
			Tag:           elementTagSOURCE,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *SourceHTMLElement) Children(children ...ElementBuilder) *SourceHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SourceHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SourceHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SourceHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SourceHTMLElement {
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

func (e *SourceHTMLElement) Escaped(text string) *SourceHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SourceHTMLElement) EscapedF(format string, args ...any) *SourceHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
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
//   - identical_to
//   - ordered_set_of_unique_space_separated_tokens
func (e *SourceHTMLElement) ACCESSKEY(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SourceHTMLElement) IfACCESSKEY(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SourceHTMLElement) RemoveACCESSKEY(v string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) AUTOCAPITALIZE(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SourceHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SourceHTMLElement) RemoveAUTOCAPITALIZE(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SourceHTMLElement) AUTOFOCUS() *SourceHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SourceHTMLElement) IfAUTOFOCUS(cond bool) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SourceHTMLElement) RemoveAUTOFOCUS() *SourceHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
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
//   - set_of_space_separated_tokens
func (e *SourceHTMLElement) CLASS(v string) *SourceHTMLElement {
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

func (e *SourceHTMLElement) IfCLASS(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SourceHTMLElement) SetCLASS(v string) *SourceHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SourceHTMLElement) RemoveCLASS(v string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) CONTENTEDITABLE(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SourceHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SourceHTMLElement) RemoveCONTENTEDITABLE(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *SourceHTMLElement) DIR(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SourceHTMLElement) IfDIR(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SourceHTMLElement) RemoveDIR(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *SourceHTMLElement) DRAGGABLE(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SourceHTMLElement) IfDRAGGABLE(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SourceHTMLElement) RemoveDRAGGABLE(v string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) ENTERKEYHINT(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SourceHTMLElement) IfENTERKEYHINT(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SourceHTMLElement) RemoveENTERKEYHINT(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *SourceHTMLElement) HEIGHT(v int) *SourceHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeHEIGHTKey] = v
	return e
}

func (e *SourceHTMLElement) IfHEIGHT(cond bool, v int) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.HEIGHT(v)
}

func (e *SourceHTMLElement) RemoveHEIGHT(v int) *SourceHTMLElement {
	delete(e.IntAttributes, attributeHEIGHTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *SourceHTMLElement) HIDDEN(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SourceHTMLElement) IfHIDDEN(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SourceHTMLElement) RemoveHIDDEN(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SourceHTMLElement) ID(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SourceHTMLElement) IfID(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SourceHTMLElement) RemoveID(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SourceHTMLElement) INERT() *SourceHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SourceHTMLElement) IfINERT(cond bool) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SourceHTMLElement) RemoveINERT() *SourceHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
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
//   - decimal
//   - email
//   - none
//   - numeric
//   - search
//   - tel
//   - text
//   - url
func (e *SourceHTMLElement) INPUTMODE(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SourceHTMLElement) IfINPUTMODE(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SourceHTMLElement) RemoveINPUTMODE(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *SourceHTMLElement) IS(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SourceHTMLElement) IfIS(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SourceHTMLElement) RemoveIS(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SourceHTMLElement) ITEMID(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SourceHTMLElement) IfITEMID(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SourceHTMLElement) RemoveITEMID(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *SourceHTMLElement) ITEMPROP(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SourceHTMLElement) IfITEMPROP(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SourceHTMLElement) RemoveITEMPROP(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SourceHTMLElement) ITEMREF(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SourceHTMLElement) IfITEMREF(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SourceHTMLElement) RemoveITEMREF(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SourceHTMLElement) ITEMSCOPE() *SourceHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SourceHTMLElement) IfITEMSCOPE(cond bool) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SourceHTMLElement) RemoveITEMSCOPE() *SourceHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
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
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *SourceHTMLElement) ITEMTYPE(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SourceHTMLElement) IfITEMTYPE(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SourceHTMLElement) RemoveITEMTYPE(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SourceHTMLElement) LANG(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SourceHTMLElement) IfLANG(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SourceHTMLElement) RemoveLANG(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// MEDIA sets the "media" attribute.
// Applicable media
// Values values are constrained to:
//   - valid_media_query_list
func (e *SourceHTMLElement) MEDIA(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMEDIAKey] = v
	return e
}

func (e *SourceHTMLElement) IfMEDIA(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.MEDIA(v)
}

func (e *SourceHTMLElement) RemoveMEDIA(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeMEDIAKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SourceHTMLElement) NONCE(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SourceHTMLElement) IfNONCE(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SourceHTMLElement) RemoveNONCE(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *SourceHTMLElement) POPOVER(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SourceHTMLElement) IfPOPOVER(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SourceHTMLElement) RemovePOPOVER(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SIZES sets the "sizes" attribute.
// Image sizes for different page layouts
// Values values are constrained to:
//   - valid_source_size_list
func (e *SourceHTMLElement) SIZES(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSIZESKey] = v
	return e
}

func (e *SourceHTMLElement) IfSIZES(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.SIZES(v)
}

func (e *SourceHTMLElement) RemoveSIZES(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeSIZESKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SourceHTMLElement) SLOT(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SourceHTMLElement) IfSLOT(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SourceHTMLElement) RemoveSLOT(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *SourceHTMLElement) SPELLCHECK(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SourceHTMLElement) IfSPELLCHECK(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SourceHTMLElement) RemoveSPELLCHECK(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *SourceHTMLElement) SRC(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCKey] = v
	return e
}

func (e *SourceHTMLElement) IfSRC(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.SRC(v)
}

func (e *SourceHTMLElement) RemoveSRC(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeSRCKey)
	return e
}

// SRCSET sets the "srcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc.
// Values values are constrained to:
//   - image_candidate_strings
func (e *SourceHTMLElement) SRCSET(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCSETKey] = v
	return e
}

func (e *SourceHTMLElement) IfSRCSET(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.SRCSET(v)
}

func (e *SourceHTMLElement) RemoveSRCSET(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeSRCSETKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SourceHTMLElement) STYLE(k, v string) *SourceHTMLElement {
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

func (e *SourceHTMLElement) IfSTYLE(cond bool, k string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SourceHTMLElement) RemoveSTYLE(k string) *SourceHTMLElement {
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
func (e *SourceHTMLElement) TABINDEX(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SourceHTMLElement) IfTABINDEX(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SourceHTMLElement) RemoveTABINDEX(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SourceHTMLElement) TITLE(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SourceHTMLElement) IfTITLE(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SourceHTMLElement) RemoveTITLE(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *SourceHTMLElement) TRANSLATE(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SourceHTMLElement) IfTRANSLATE(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SourceHTMLElement) RemoveTRANSLATE(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//   - java_script_mime_type_essence_match
//   - module
//   - valid_mime_type_string
func (e *SourceHTMLElement) TYPE(v string) *SourceHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTYPEKey] = v
	return e
}

func (e *SourceHTMLElement) IfTYPE(cond bool, v string) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.TYPE(v)
}

func (e *SourceHTMLElement) RemoveTYPE(v string) *SourceHTMLElement {
	delete(e.StringAttributes, attributeTYPEKey)
	return e
}

// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *SourceHTMLElement) WIDTH(v int) *SourceHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeWIDTHKey] = v
	return e
}

func (e *SourceHTMLElement) IfWIDTH(cond bool, v int) *SourceHTMLElement {
	if !cond {
		return e
	}
	return e.WIDTH(v)
}

func (e *SourceHTMLElement) RemoveWIDTH(v int) *SourceHTMLElement {
	delete(e.IntAttributes, attributeWIDTHKey)
	return e
}
