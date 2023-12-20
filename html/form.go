package html

import (
	"fmt"
)

type FormHTMLElement struct {
	*Element
}

func FORM(children ...ElementBuilder) *FormHTMLElement {
	return &FormHTMLElement{
		Element: &Element{
			Tag:           elementTagFORM,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *FormHTMLElement) Children(children ...ElementBuilder) *FormHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *FormHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *FormHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *FormHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *FormHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *FormHTMLElement) Text(text string) *FormHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *FormHTMLElement) TextF(format string, args ...any) *FormHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *FormHTMLElement) Escaped(text string) *FormHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *FormHTMLElement) EscapedF(format string, args ...any) *FormHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *FormHTMLElement) CustomData(key, value string) *FormHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *FormHTMLElement) CustomDataRemove(key string) *FormHTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}

// ACCEPT_CHARSET sets the "accept-charset" attribute.
// Character encodings to use for form submission
// Values values are constrained to:
//   - ascii_case_insensitive
//   - utf_8
func (e *FormHTMLElement) ACCEPT_CHARSET(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCEPT_CHARSETKey] = v
	return e
}

func (e *FormHTMLElement) IfACCEPT_CHARSET(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ACCEPT_CHARSET(v)
}

func (e *FormHTMLElement) RemoveACCEPT_CHARSET(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeACCEPT_CHARSETKey)
	return e
}

// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//   - identical_to
//   - ordered_set_of_unique_space_separated_tokens
func (e *FormHTMLElement) ACCESSKEY(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *FormHTMLElement) IfACCESSKEY(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *FormHTMLElement) RemoveACCESSKEY(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// ACTION sets the "action" attribute.
// URL to use for form submission
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *FormHTMLElement) ACTION(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACTIONKey] = v
	return e
}

func (e *FormHTMLElement) IfACTION(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ACTION(v)
}

func (e *FormHTMLElement) RemoveACTION(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeACTIONKey)
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
func (e *FormHTMLElement) AUTOCAPITALIZE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *FormHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *FormHTMLElement) RemoveAUTOCAPITALIZE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//   - autofill_field
func (e *FormHTMLElement) AUTOCOMPLETE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCOMPLETEKey] = v
	return e
}

func (e *FormHTMLElement) IfAUTOCOMPLETE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCOMPLETE(v)
}

func (e *FormHTMLElement) RemoveAUTOCOMPLETE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeAUTOCOMPLETEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *FormHTMLElement) AUTOFOCUS() *FormHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *FormHTMLElement) IfAUTOFOCUS(cond bool) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *FormHTMLElement) RemoveAUTOFOCUS() *FormHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *FormHTMLElement) SetAUTOFOCUS(b bool) *FormHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *FormHTMLElement) CLASS(v string) *FormHTMLElement {
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

func (e *FormHTMLElement) IfCLASS(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *FormHTMLElement) SetCLASS(v string) *FormHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *FormHTMLElement) RemoveCLASS(v string) *FormHTMLElement {
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
func (e *FormHTMLElement) CONTENTEDITABLE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *FormHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *FormHTMLElement) RemoveCONTENTEDITABLE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *FormHTMLElement) DIR(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *FormHTMLElement) IfDIR(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *FormHTMLElement) RemoveDIR(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *FormHTMLElement) DRAGGABLE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *FormHTMLElement) IfDRAGGABLE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *FormHTMLElement) RemoveDRAGGABLE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeDRAGGABLEKey)
	return e
}

// ENCTYPE sets the "enctype" attribute.
// Entry list encoding type to use for form submission
// Values values are constrained to:
//   - application/x_www_form_urlencoded
//   - multipart/form_data
//   - text/plain
func (e *FormHTMLElement) ENCTYPE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENCTYPEKey] = v
	return e
}

func (e *FormHTMLElement) IfENCTYPE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ENCTYPE(v)
}

func (e *FormHTMLElement) RemoveENCTYPE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeENCTYPEKey)
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
func (e *FormHTMLElement) ENTERKEYHINT(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *FormHTMLElement) IfENTERKEYHINT(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *FormHTMLElement) RemoveENTERKEYHINT(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *FormHTMLElement) HIDDEN(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *FormHTMLElement) IfHIDDEN(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *FormHTMLElement) RemoveHIDDEN(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *FormHTMLElement) ID(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *FormHTMLElement) IfID(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *FormHTMLElement) RemoveID(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *FormHTMLElement) INERT() *FormHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *FormHTMLElement) IfINERT(cond bool) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *FormHTMLElement) RemoveINERT() *FormHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *FormHTMLElement) SetINERT(b bool) *FormHTMLElement {
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
func (e *FormHTMLElement) INPUTMODE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *FormHTMLElement) IfINPUTMODE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *FormHTMLElement) RemoveINPUTMODE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *FormHTMLElement) IS(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *FormHTMLElement) IfIS(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *FormHTMLElement) RemoveIS(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *FormHTMLElement) ITEMID(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *FormHTMLElement) IfITEMID(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *FormHTMLElement) RemoveITEMID(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *FormHTMLElement) ITEMPROP(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *FormHTMLElement) IfITEMPROP(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *FormHTMLElement) RemoveITEMPROP(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *FormHTMLElement) ITEMREF(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *FormHTMLElement) IfITEMREF(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *FormHTMLElement) RemoveITEMREF(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *FormHTMLElement) ITEMSCOPE() *FormHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *FormHTMLElement) IfITEMSCOPE(cond bool) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *FormHTMLElement) RemoveITEMSCOPE() *FormHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *FormHTMLElement) SetITEMSCOPE(b bool) *FormHTMLElement {
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
func (e *FormHTMLElement) ITEMTYPE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *FormHTMLElement) IfITEMTYPE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *FormHTMLElement) RemoveITEMTYPE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *FormHTMLElement) LANG(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *FormHTMLElement) IfLANG(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *FormHTMLElement) RemoveLANG(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// METHOD sets the "method" attribute.
// Variant to use for form submission
// Values values are constrained to:
//   - dialog
//   - get
//   - post
func (e *FormHTMLElement) METHOD(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMETHODKey] = v
	return e
}

func (e *FormHTMLElement) IfMETHOD(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.METHOD(v)
}

func (e *FormHTMLElement) RemoveMETHOD(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeMETHODKey)
	return e
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *FormHTMLElement) NAME(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *FormHTMLElement) IfNAME(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *FormHTMLElement) RemoveNAME(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *FormHTMLElement) NONCE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *FormHTMLElement) IfNONCE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *FormHTMLElement) RemoveNONCE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// NOVALIDATE sets the "novalidate" attribute.
// Bypass form control validation for form submission
// Values values are constrained to:
//   - boolean_attribute
func (e *FormHTMLElement) NOVALIDATE() *FormHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeNOVALIDATEKey] = struct{}{}
	return e
}

func (e *FormHTMLElement) IfNOVALIDATE(cond bool) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.NOVALIDATE()
}

func (e *FormHTMLElement) RemoveNOVALIDATE() *FormHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeNOVALIDATEKey)
	return e
}

func (e *FormHTMLElement) SetNOVALIDATE(b bool) *FormHTMLElement {
	if b {
		return e.NOVALIDATE()
	}
	return e.RemoveNOVALIDATE()
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *FormHTMLElement) POPOVER(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *FormHTMLElement) IfPOPOVER(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *FormHTMLElement) RemovePOPOVER(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *FormHTMLElement) SLOT(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *FormHTMLElement) IfSLOT(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *FormHTMLElement) RemoveSLOT(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *FormHTMLElement) SPELLCHECK(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *FormHTMLElement) IfSPELLCHECK(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *FormHTMLElement) RemoveSPELLCHECK(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *FormHTMLElement) STYLE(k, v string) *FormHTMLElement {
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

func (e *FormHTMLElement) IfSTYLE(cond bool, k string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *FormHTMLElement) RemoveSTYLE(k string) *FormHTMLElement {
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
func (e *FormHTMLElement) TABINDEX(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *FormHTMLElement) IfTABINDEX(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *FormHTMLElement) RemoveTABINDEX(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TARGET sets the "target" attribute.
// Navigable for form submission
// Values values are constrained to:
//   - valid_navigable_target_name_or_keyword
func (e *FormHTMLElement) TARGET(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTARGETKey] = v
	return e
}

func (e *FormHTMLElement) IfTARGET(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.TARGET(v)
}

func (e *FormHTMLElement) RemoveTARGET(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeTARGETKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *FormHTMLElement) TITLE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *FormHTMLElement) IfTITLE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *FormHTMLElement) RemoveTITLE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *FormHTMLElement) TRANSLATE(v string) *FormHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *FormHTMLElement) IfTRANSLATE(cond bool, v string) *FormHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *FormHTMLElement) RemoveTRANSLATE(v string) *FormHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
