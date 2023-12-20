package html

import (
	"fmt"
)

type EmbedHTMLElement struct {
	*Element
}

func EMBED(children ...ElementBuilder) *EmbedHTMLElement {
	return &EmbedHTMLElement{
		Element: &Element{
			Tag:           elementTagEMBED,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *EmbedHTMLElement) Children(children ...ElementBuilder) *EmbedHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *EmbedHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *EmbedHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *EmbedHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *EmbedHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *EmbedHTMLElement) Text(text string) *EmbedHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *EmbedHTMLElement) TextF(format string, args ...any) *EmbedHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *EmbedHTMLElement) Escaped(text string) *EmbedHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *EmbedHTMLElement) EscapedF(format string, args ...any) *EmbedHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *EmbedHTMLElement) CustomData(key, value string) *EmbedHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *EmbedHTMLElement) CustomDataRemove(key string) *EmbedHTMLElement {
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
func (e *EmbedHTMLElement) ACCESSKEY(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *EmbedHTMLElement) IfACCESSKEY(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *EmbedHTMLElement) RemoveACCESSKEY(v string) *EmbedHTMLElement {
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
func (e *EmbedHTMLElement) AUTOCAPITALIZE(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *EmbedHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *EmbedHTMLElement) RemoveAUTOCAPITALIZE(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *EmbedHTMLElement) AUTOFOCUS() *EmbedHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *EmbedHTMLElement) IfAUTOFOCUS(cond bool) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *EmbedHTMLElement) RemoveAUTOFOCUS() *EmbedHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *EmbedHTMLElement) SetAUTOFOCUS(b bool) *EmbedHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *EmbedHTMLElement) CLASS(v string) *EmbedHTMLElement {
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

func (e *EmbedHTMLElement) IfCLASS(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *EmbedHTMLElement) SetCLASS(v string) *EmbedHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *EmbedHTMLElement) RemoveCLASS(v string) *EmbedHTMLElement {
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
func (e *EmbedHTMLElement) CONTENTEDITABLE(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *EmbedHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *EmbedHTMLElement) RemoveCONTENTEDITABLE(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *EmbedHTMLElement) DIR(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *EmbedHTMLElement) IfDIR(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *EmbedHTMLElement) RemoveDIR(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *EmbedHTMLElement) DRAGGABLE(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *EmbedHTMLElement) IfDRAGGABLE(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *EmbedHTMLElement) RemoveDRAGGABLE(v string) *EmbedHTMLElement {
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
func (e *EmbedHTMLElement) ENTERKEYHINT(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *EmbedHTMLElement) IfENTERKEYHINT(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *EmbedHTMLElement) RemoveENTERKEYHINT(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *EmbedHTMLElement) HEIGHT(v int) *EmbedHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeHEIGHTKey] = v
	return e
}

func (e *EmbedHTMLElement) IfHEIGHT(cond bool, v int) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.HEIGHT(v)
}

func (e *EmbedHTMLElement) RemoveHEIGHT(v int) *EmbedHTMLElement {
	delete(e.IntAttributes, attributeHEIGHTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *EmbedHTMLElement) HIDDEN(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *EmbedHTMLElement) IfHIDDEN(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *EmbedHTMLElement) RemoveHIDDEN(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *EmbedHTMLElement) ID(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *EmbedHTMLElement) IfID(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *EmbedHTMLElement) RemoveID(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *EmbedHTMLElement) INERT() *EmbedHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *EmbedHTMLElement) IfINERT(cond bool) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *EmbedHTMLElement) RemoveINERT() *EmbedHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *EmbedHTMLElement) SetINERT(b bool) *EmbedHTMLElement {
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
func (e *EmbedHTMLElement) INPUTMODE(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *EmbedHTMLElement) IfINPUTMODE(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *EmbedHTMLElement) RemoveINPUTMODE(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *EmbedHTMLElement) IS(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *EmbedHTMLElement) IfIS(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *EmbedHTMLElement) RemoveIS(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *EmbedHTMLElement) ITEMID(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *EmbedHTMLElement) IfITEMID(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *EmbedHTMLElement) RemoveITEMID(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *EmbedHTMLElement) ITEMPROP(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *EmbedHTMLElement) IfITEMPROP(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *EmbedHTMLElement) RemoveITEMPROP(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *EmbedHTMLElement) ITEMREF(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *EmbedHTMLElement) IfITEMREF(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *EmbedHTMLElement) RemoveITEMREF(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *EmbedHTMLElement) ITEMSCOPE() *EmbedHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *EmbedHTMLElement) IfITEMSCOPE(cond bool) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *EmbedHTMLElement) RemoveITEMSCOPE() *EmbedHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *EmbedHTMLElement) SetITEMSCOPE(b bool) *EmbedHTMLElement {
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
func (e *EmbedHTMLElement) ITEMTYPE(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *EmbedHTMLElement) IfITEMTYPE(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *EmbedHTMLElement) RemoveITEMTYPE(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *EmbedHTMLElement) LANG(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *EmbedHTMLElement) IfLANG(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *EmbedHTMLElement) RemoveLANG(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *EmbedHTMLElement) NONCE(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *EmbedHTMLElement) IfNONCE(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *EmbedHTMLElement) RemoveNONCE(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *EmbedHTMLElement) POPOVER(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *EmbedHTMLElement) IfPOPOVER(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *EmbedHTMLElement) RemovePOPOVER(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *EmbedHTMLElement) SLOT(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *EmbedHTMLElement) IfSLOT(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *EmbedHTMLElement) RemoveSLOT(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *EmbedHTMLElement) SPELLCHECK(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *EmbedHTMLElement) IfSPELLCHECK(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *EmbedHTMLElement) RemoveSPELLCHECK(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *EmbedHTMLElement) SRC(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCKey] = v
	return e
}

func (e *EmbedHTMLElement) IfSRC(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.SRC(v)
}

func (e *EmbedHTMLElement) RemoveSRC(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeSRCKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *EmbedHTMLElement) STYLE(k, v string) *EmbedHTMLElement {
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

func (e *EmbedHTMLElement) IfSTYLE(cond bool, k string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *EmbedHTMLElement) RemoveSTYLE(k string) *EmbedHTMLElement {
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
func (e *EmbedHTMLElement) TABINDEX(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *EmbedHTMLElement) IfTABINDEX(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *EmbedHTMLElement) RemoveTABINDEX(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *EmbedHTMLElement) TITLE(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *EmbedHTMLElement) IfTITLE(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *EmbedHTMLElement) RemoveTITLE(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *EmbedHTMLElement) TRANSLATE(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *EmbedHTMLElement) IfTRANSLATE(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *EmbedHTMLElement) RemoveTRANSLATE(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//   - java_script_mime_type_essence_match
//   - module
//   - valid_mime_type_string
func (e *EmbedHTMLElement) TYPE(v string) *EmbedHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTYPEKey] = v
	return e
}

func (e *EmbedHTMLElement) IfTYPE(cond bool, v string) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.TYPE(v)
}

func (e *EmbedHTMLElement) RemoveTYPE(v string) *EmbedHTMLElement {
	delete(e.StringAttributes, attributeTYPEKey)
	return e
}

// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *EmbedHTMLElement) WIDTH(v int) *EmbedHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeWIDTHKey] = v
	return e
}

func (e *EmbedHTMLElement) IfWIDTH(cond bool, v int) *EmbedHTMLElement {
	if !cond {
		return e
	}
	return e.WIDTH(v)
}

func (e *EmbedHTMLElement) RemoveWIDTH(v int) *EmbedHTMLElement {
	delete(e.IntAttributes, attributeWIDTHKey)
	return e
}
