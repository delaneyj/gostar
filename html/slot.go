package html

import (
	"fmt"
)

type SlotHTMLElement struct {
	*Element
}

func SLOT(children ...ElementBuilder) *SlotHTMLElement {
	return &SlotHTMLElement{
		Element: &Element{
			Tag:           elementTagSLOT,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SlotHTMLElement) Children(children ...ElementBuilder) *SlotHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SlotHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SlotHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SlotHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SlotHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SlotHTMLElement) Text(text string) *SlotHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SlotHTMLElement) TextF(format string, args ...any) *SlotHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SlotHTMLElement) Escaped(text string) *SlotHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SlotHTMLElement) EscapedF(format string, args ...any) *SlotHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SlotHTMLElement) CustomData(key, value string) *SlotHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SlotHTMLElement) CustomDataRemove(key string) *SlotHTMLElement {
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
func (e *SlotHTMLElement) ACCESSKEY(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SlotHTMLElement) IfACCESSKEY(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SlotHTMLElement) RemoveACCESSKEY(v string) *SlotHTMLElement {
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
func (e *SlotHTMLElement) AUTOCAPITALIZE(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SlotHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SlotHTMLElement) RemoveAUTOCAPITALIZE(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SlotHTMLElement) AUTOFOCUS() *SlotHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SlotHTMLElement) IfAUTOFOCUS(cond bool) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SlotHTMLElement) RemoveAUTOFOCUS() *SlotHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SlotHTMLElement) SetAUTOFOCUS(b bool) *SlotHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SlotHTMLElement) CLASS(v string) *SlotHTMLElement {
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

func (e *SlotHTMLElement) IfCLASS(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SlotHTMLElement) SetCLASS(v string) *SlotHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SlotHTMLElement) RemoveCLASS(v string) *SlotHTMLElement {
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
func (e *SlotHTMLElement) CONTENTEDITABLE(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SlotHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SlotHTMLElement) RemoveCONTENTEDITABLE(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *SlotHTMLElement) DIR(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SlotHTMLElement) IfDIR(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SlotHTMLElement) RemoveDIR(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *SlotHTMLElement) DRAGGABLE(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SlotHTMLElement) IfDRAGGABLE(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SlotHTMLElement) RemoveDRAGGABLE(v string) *SlotHTMLElement {
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
func (e *SlotHTMLElement) ENTERKEYHINT(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SlotHTMLElement) IfENTERKEYHINT(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SlotHTMLElement) RemoveENTERKEYHINT(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *SlotHTMLElement) HIDDEN(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SlotHTMLElement) IfHIDDEN(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SlotHTMLElement) RemoveHIDDEN(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SlotHTMLElement) ID(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SlotHTMLElement) IfID(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SlotHTMLElement) RemoveID(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SlotHTMLElement) INERT() *SlotHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SlotHTMLElement) IfINERT(cond bool) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SlotHTMLElement) RemoveINERT() *SlotHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SlotHTMLElement) SetINERT(b bool) *SlotHTMLElement {
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
func (e *SlotHTMLElement) INPUTMODE(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SlotHTMLElement) IfINPUTMODE(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SlotHTMLElement) RemoveINPUTMODE(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *SlotHTMLElement) IS(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SlotHTMLElement) IfIS(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SlotHTMLElement) RemoveIS(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SlotHTMLElement) ITEMID(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SlotHTMLElement) IfITEMID(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SlotHTMLElement) RemoveITEMID(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *SlotHTMLElement) ITEMPROP(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SlotHTMLElement) IfITEMPROP(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SlotHTMLElement) RemoveITEMPROP(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SlotHTMLElement) ITEMREF(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SlotHTMLElement) IfITEMREF(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SlotHTMLElement) RemoveITEMREF(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SlotHTMLElement) ITEMSCOPE() *SlotHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SlotHTMLElement) IfITEMSCOPE(cond bool) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SlotHTMLElement) RemoveITEMSCOPE() *SlotHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SlotHTMLElement) SetITEMSCOPE(b bool) *SlotHTMLElement {
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
func (e *SlotHTMLElement) ITEMTYPE(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SlotHTMLElement) IfITEMTYPE(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SlotHTMLElement) RemoveITEMTYPE(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SlotHTMLElement) LANG(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SlotHTMLElement) IfLANG(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SlotHTMLElement) RemoveLANG(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *SlotHTMLElement) NAME(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *SlotHTMLElement) IfNAME(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *SlotHTMLElement) RemoveNAME(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SlotHTMLElement) NONCE(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SlotHTMLElement) IfNONCE(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SlotHTMLElement) RemoveNONCE(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *SlotHTMLElement) POPOVER(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SlotHTMLElement) IfPOPOVER(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SlotHTMLElement) RemovePOPOVER(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SlotHTMLElement) SLOT(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SlotHTMLElement) IfSLOT(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SlotHTMLElement) RemoveSLOT(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *SlotHTMLElement) SPELLCHECK(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SlotHTMLElement) IfSPELLCHECK(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SlotHTMLElement) RemoveSPELLCHECK(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SlotHTMLElement) STYLE(k, v string) *SlotHTMLElement {
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

func (e *SlotHTMLElement) IfSTYLE(cond bool, k string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SlotHTMLElement) RemoveSTYLE(k string) *SlotHTMLElement {
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
func (e *SlotHTMLElement) TABINDEX(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SlotHTMLElement) IfTABINDEX(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SlotHTMLElement) RemoveTABINDEX(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SlotHTMLElement) TITLE(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SlotHTMLElement) IfTITLE(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SlotHTMLElement) RemoveTITLE(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *SlotHTMLElement) TRANSLATE(v string) *SlotHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SlotHTMLElement) IfTRANSLATE(cond bool, v string) *SlotHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SlotHTMLElement) RemoveTRANSLATE(v string) *SlotHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
