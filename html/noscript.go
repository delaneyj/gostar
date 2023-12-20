package html

import (
	"fmt"
)

type NoscriptHTMLElement struct {
	*Element
}

func NOSCRIPT(children ...ElementBuilder) *NoscriptHTMLElement {
	return &NoscriptHTMLElement{
		Element: &Element{
			Tag:           elementTagNOSCRIPT,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *NoscriptHTMLElement) Children(children ...ElementBuilder) *NoscriptHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *NoscriptHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *NoscriptHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *NoscriptHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *NoscriptHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *NoscriptHTMLElement) Text(text string) *NoscriptHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *NoscriptHTMLElement) TextF(format string, args ...any) *NoscriptHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *NoscriptHTMLElement) Escaped(text string) *NoscriptHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *NoscriptHTMLElement) EscapedF(format string, args ...any) *NoscriptHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *NoscriptHTMLElement) CustomData(key, value string) *NoscriptHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *NoscriptHTMLElement) CustomDataRemove(key string) *NoscriptHTMLElement {
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
func (e *NoscriptHTMLElement) ACCESSKEY(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfACCESSKEY(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *NoscriptHTMLElement) RemoveACCESSKEY(v string) *NoscriptHTMLElement {
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
func (e *NoscriptHTMLElement) AUTOCAPITALIZE(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *NoscriptHTMLElement) RemoveAUTOCAPITALIZE(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *NoscriptHTMLElement) AUTOFOCUS() *NoscriptHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *NoscriptHTMLElement) IfAUTOFOCUS(cond bool) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *NoscriptHTMLElement) RemoveAUTOFOCUS() *NoscriptHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *NoscriptHTMLElement) SetAUTOFOCUS(b bool) *NoscriptHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *NoscriptHTMLElement) CLASS(v string) *NoscriptHTMLElement {
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

func (e *NoscriptHTMLElement) IfCLASS(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *NoscriptHTMLElement) SetCLASS(v string) *NoscriptHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *NoscriptHTMLElement) RemoveCLASS(v string) *NoscriptHTMLElement {
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
func (e *NoscriptHTMLElement) CONTENTEDITABLE(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *NoscriptHTMLElement) RemoveCONTENTEDITABLE(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *NoscriptHTMLElement) DIR(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfDIR(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *NoscriptHTMLElement) RemoveDIR(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *NoscriptHTMLElement) DRAGGABLE(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfDRAGGABLE(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *NoscriptHTMLElement) RemoveDRAGGABLE(v string) *NoscriptHTMLElement {
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
func (e *NoscriptHTMLElement) ENTERKEYHINT(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfENTERKEYHINT(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *NoscriptHTMLElement) RemoveENTERKEYHINT(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *NoscriptHTMLElement) HIDDEN(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfHIDDEN(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *NoscriptHTMLElement) RemoveHIDDEN(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *NoscriptHTMLElement) ID(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfID(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *NoscriptHTMLElement) RemoveID(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *NoscriptHTMLElement) INERT() *NoscriptHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *NoscriptHTMLElement) IfINERT(cond bool) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *NoscriptHTMLElement) RemoveINERT() *NoscriptHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *NoscriptHTMLElement) SetINERT(b bool) *NoscriptHTMLElement {
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
func (e *NoscriptHTMLElement) INPUTMODE(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfINPUTMODE(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *NoscriptHTMLElement) RemoveINPUTMODE(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *NoscriptHTMLElement) IS(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfIS(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *NoscriptHTMLElement) RemoveIS(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *NoscriptHTMLElement) ITEMID(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfITEMID(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *NoscriptHTMLElement) RemoveITEMID(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *NoscriptHTMLElement) ITEMPROP(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfITEMPROP(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *NoscriptHTMLElement) RemoveITEMPROP(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *NoscriptHTMLElement) ITEMREF(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfITEMREF(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *NoscriptHTMLElement) RemoveITEMREF(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *NoscriptHTMLElement) ITEMSCOPE() *NoscriptHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *NoscriptHTMLElement) IfITEMSCOPE(cond bool) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *NoscriptHTMLElement) RemoveITEMSCOPE() *NoscriptHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *NoscriptHTMLElement) SetITEMSCOPE(b bool) *NoscriptHTMLElement {
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
func (e *NoscriptHTMLElement) ITEMTYPE(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfITEMTYPE(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *NoscriptHTMLElement) RemoveITEMTYPE(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *NoscriptHTMLElement) LANG(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfLANG(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *NoscriptHTMLElement) RemoveLANG(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *NoscriptHTMLElement) NONCE(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfNONCE(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *NoscriptHTMLElement) RemoveNONCE(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *NoscriptHTMLElement) POPOVER(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfPOPOVER(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *NoscriptHTMLElement) RemovePOPOVER(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *NoscriptHTMLElement) SLOT(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfSLOT(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *NoscriptHTMLElement) RemoveSLOT(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *NoscriptHTMLElement) SPELLCHECK(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfSPELLCHECK(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *NoscriptHTMLElement) RemoveSPELLCHECK(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *NoscriptHTMLElement) STYLE(k, v string) *NoscriptHTMLElement {
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

func (e *NoscriptHTMLElement) IfSTYLE(cond bool, k string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *NoscriptHTMLElement) RemoveSTYLE(k string) *NoscriptHTMLElement {
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
func (e *NoscriptHTMLElement) TABINDEX(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfTABINDEX(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *NoscriptHTMLElement) RemoveTABINDEX(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *NoscriptHTMLElement) TITLE(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfTITLE(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *NoscriptHTMLElement) RemoveTITLE(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *NoscriptHTMLElement) TRANSLATE(v string) *NoscriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *NoscriptHTMLElement) IfTRANSLATE(cond bool, v string) *NoscriptHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *NoscriptHTMLElement) RemoveTRANSLATE(v string) *NoscriptHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
