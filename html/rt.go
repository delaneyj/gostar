package html

import (
	"fmt"
)

type RtHTMLElement struct {
	*Element
}

func RT(children ...ElementBuilder) *RtHTMLElement {
	return &RtHTMLElement{
		Element: &Element{
			Tag:           elementTagRT,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *RtHTMLElement) Children(children ...ElementBuilder) *RtHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *RtHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *RtHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *RtHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *RtHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *RtHTMLElement) Text(text string) *RtHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *RtHTMLElement) TextF(format string, args ...any) *RtHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *RtHTMLElement) Escaped(text string) *RtHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *RtHTMLElement) EscapedF(format string, args ...any) *RtHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *RtHTMLElement) CustomData(key, value string) *RtHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *RtHTMLElement) CustomDataRemove(key string) *RtHTMLElement {
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
func (e *RtHTMLElement) ACCESSKEY(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *RtHTMLElement) IfACCESSKEY(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *RtHTMLElement) RemoveACCESSKEY(v string) *RtHTMLElement {
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
func (e *RtHTMLElement) AUTOCAPITALIZE(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *RtHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *RtHTMLElement) RemoveAUTOCAPITALIZE(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *RtHTMLElement) AUTOFOCUS() *RtHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *RtHTMLElement) IfAUTOFOCUS(cond bool) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *RtHTMLElement) RemoveAUTOFOCUS() *RtHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *RtHTMLElement) SetAUTOFOCUS(b bool) *RtHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *RtHTMLElement) CLASS(v string) *RtHTMLElement {
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

func (e *RtHTMLElement) IfCLASS(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *RtHTMLElement) SetCLASS(v string) *RtHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *RtHTMLElement) RemoveCLASS(v string) *RtHTMLElement {
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
func (e *RtHTMLElement) CONTENTEDITABLE(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *RtHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *RtHTMLElement) RemoveCONTENTEDITABLE(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *RtHTMLElement) DIR(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *RtHTMLElement) IfDIR(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *RtHTMLElement) RemoveDIR(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *RtHTMLElement) DRAGGABLE(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *RtHTMLElement) IfDRAGGABLE(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *RtHTMLElement) RemoveDRAGGABLE(v string) *RtHTMLElement {
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
func (e *RtHTMLElement) ENTERKEYHINT(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *RtHTMLElement) IfENTERKEYHINT(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *RtHTMLElement) RemoveENTERKEYHINT(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *RtHTMLElement) HIDDEN(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *RtHTMLElement) IfHIDDEN(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *RtHTMLElement) RemoveHIDDEN(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *RtHTMLElement) ID(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *RtHTMLElement) IfID(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *RtHTMLElement) RemoveID(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *RtHTMLElement) INERT() *RtHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *RtHTMLElement) IfINERT(cond bool) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *RtHTMLElement) RemoveINERT() *RtHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *RtHTMLElement) SetINERT(b bool) *RtHTMLElement {
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
func (e *RtHTMLElement) INPUTMODE(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *RtHTMLElement) IfINPUTMODE(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *RtHTMLElement) RemoveINPUTMODE(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *RtHTMLElement) IS(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *RtHTMLElement) IfIS(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *RtHTMLElement) RemoveIS(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *RtHTMLElement) ITEMID(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *RtHTMLElement) IfITEMID(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *RtHTMLElement) RemoveITEMID(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *RtHTMLElement) ITEMPROP(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *RtHTMLElement) IfITEMPROP(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *RtHTMLElement) RemoveITEMPROP(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *RtHTMLElement) ITEMREF(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *RtHTMLElement) IfITEMREF(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *RtHTMLElement) RemoveITEMREF(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *RtHTMLElement) ITEMSCOPE() *RtHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *RtHTMLElement) IfITEMSCOPE(cond bool) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *RtHTMLElement) RemoveITEMSCOPE() *RtHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *RtHTMLElement) SetITEMSCOPE(b bool) *RtHTMLElement {
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
func (e *RtHTMLElement) ITEMTYPE(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *RtHTMLElement) IfITEMTYPE(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *RtHTMLElement) RemoveITEMTYPE(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *RtHTMLElement) LANG(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *RtHTMLElement) IfLANG(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *RtHTMLElement) RemoveLANG(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *RtHTMLElement) NONCE(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *RtHTMLElement) IfNONCE(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *RtHTMLElement) RemoveNONCE(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *RtHTMLElement) POPOVER(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *RtHTMLElement) IfPOPOVER(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *RtHTMLElement) RemovePOPOVER(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *RtHTMLElement) SLOT(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *RtHTMLElement) IfSLOT(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *RtHTMLElement) RemoveSLOT(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *RtHTMLElement) SPELLCHECK(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *RtHTMLElement) IfSPELLCHECK(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *RtHTMLElement) RemoveSPELLCHECK(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *RtHTMLElement) STYLE(k, v string) *RtHTMLElement {
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

func (e *RtHTMLElement) IfSTYLE(cond bool, k string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *RtHTMLElement) RemoveSTYLE(k string) *RtHTMLElement {
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
func (e *RtHTMLElement) TABINDEX(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *RtHTMLElement) IfTABINDEX(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *RtHTMLElement) RemoveTABINDEX(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *RtHTMLElement) TITLE(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *RtHTMLElement) IfTITLE(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *RtHTMLElement) RemoveTITLE(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *RtHTMLElement) TRANSLATE(v string) *RtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *RtHTMLElement) IfTRANSLATE(cond bool, v string) *RtHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *RtHTMLElement) RemoveTRANSLATE(v string) *RtHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
