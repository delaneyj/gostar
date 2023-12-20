package html

import (
	"fmt"
)

type InputHTMLElement struct {
	*Element
}

func INPUT(children ...ElementBuilder) *InputHTMLElement {
	return &InputHTMLElement{
		Element: &Element{
			Tag:           elementTagINPUT,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *InputHTMLElement) Children(children ...ElementBuilder) *InputHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *InputHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *InputHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *InputHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *InputHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *InputHTMLElement) Text(text string) *InputHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *InputHTMLElement) TextF(format string, args ...any) *InputHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *InputHTMLElement) Escaped(text string) *InputHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *InputHTMLElement) EscapedF(format string, args ...any) *InputHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *InputHTMLElement) CustomData(key, value string) *InputHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *InputHTMLElement) CustomDataRemove(key string) *InputHTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}

// ACCEPT sets the "accept" attribute.
// Hint for expected file type in file upload controls
// Values values are constrained to:
//   - audio/*
//   - image/*
//   - set_of_comma_separated_tokens
//   - valid_mime_type_strings_with_no_parameters
//   - video/*
func (e *InputHTMLElement) ACCEPT(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCEPTKey] = v
	return e
}

func (e *InputHTMLElement) IfACCEPT(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ACCEPT(v)
}

func (e *InputHTMLElement) RemoveACCEPT(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeACCEPTKey)
	return e
}

// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//   - identical_to
//   - ordered_set_of_unique_space_separated_tokens
func (e *InputHTMLElement) ACCESSKEY(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *InputHTMLElement) IfACCESSKEY(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *InputHTMLElement) RemoveACCESSKEY(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// ALT sets the "alt" attribute.
// Replacement text for use when images are not available
// Values values are constrained to:
//   - text
func (e *InputHTMLElement) ALT(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeALTKey] = v
	return e
}

func (e *InputHTMLElement) IfALT(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ALT(v)
}

func (e *InputHTMLElement) RemoveALT(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeALTKey)
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
func (e *InputHTMLElement) AUTOCAPITALIZE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *InputHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *InputHTMLElement) RemoveAUTOCAPITALIZE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//   - autofill_field
func (e *InputHTMLElement) AUTOCOMPLETE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCOMPLETEKey] = v
	return e
}

func (e *InputHTMLElement) IfAUTOCOMPLETE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCOMPLETE(v)
}

func (e *InputHTMLElement) RemoveAUTOCOMPLETE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeAUTOCOMPLETEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *InputHTMLElement) AUTOFOCUS() *InputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *InputHTMLElement) IfAUTOFOCUS(cond bool) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *InputHTMLElement) RemoveAUTOFOCUS() *InputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *InputHTMLElement) SetAUTOFOCUS(b bool) *InputHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CHECKED sets the "checked" attribute.
// Whether the control is checked
// Values values are constrained to:
//   - boolean_attribute
func (e *InputHTMLElement) CHECKED() *InputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeCHECKEDKey] = struct{}{}
	return e
}

func (e *InputHTMLElement) IfCHECKED(cond bool) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.CHECKED()
}

func (e *InputHTMLElement) RemoveCHECKED() *InputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeCHECKEDKey)
	return e
}

func (e *InputHTMLElement) SetCHECKED(b bool) *InputHTMLElement {
	if b {
		return e.CHECKED()
	}
	return e.RemoveCHECKED()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *InputHTMLElement) CLASS(v string) *InputHTMLElement {
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

func (e *InputHTMLElement) IfCLASS(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *InputHTMLElement) SetCLASS(v string) *InputHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *InputHTMLElement) RemoveCLASS(v string) *InputHTMLElement {
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
func (e *InputHTMLElement) CONTENTEDITABLE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *InputHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *InputHTMLElement) RemoveCONTENTEDITABLE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *InputHTMLElement) DIR(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *InputHTMLElement) IfDIR(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *InputHTMLElement) RemoveDIR(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DIRNAME sets the "dirname" attribute.
// Name of form control to use for sending the element's directionality in form submission
// Values values are constrained to:
//   - text
func (e *InputHTMLElement) DIRNAME(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRNAMEKey] = v
	return e
}

func (e *InputHTMLElement) IfDIRNAME(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.DIRNAME(v)
}

func (e *InputHTMLElement) RemoveDIRNAME(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeDIRNAMEKey)
	return e
}

// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//   - boolean_attribute
func (e *InputHTMLElement) DISABLED() *InputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeDISABLEDKey] = struct{}{}
	return e
}

func (e *InputHTMLElement) IfDISABLED(cond bool) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.DISABLED()
}

func (e *InputHTMLElement) RemoveDISABLED() *InputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeDISABLEDKey)
	return e
}

func (e *InputHTMLElement) SetDISABLED(b bool) *InputHTMLElement {
	if b {
		return e.DISABLED()
	}
	return e.RemoveDISABLED()
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *InputHTMLElement) DRAGGABLE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *InputHTMLElement) IfDRAGGABLE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *InputHTMLElement) RemoveDRAGGABLE(v string) *InputHTMLElement {
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
func (e *InputHTMLElement) ENTERKEYHINT(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *InputHTMLElement) IfENTERKEYHINT(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *InputHTMLElement) RemoveENTERKEYHINT(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//   - id
func (e *InputHTMLElement) FORM(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMKey] = v
	return e
}

func (e *InputHTMLElement) IfFORM(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.FORM(v)
}

func (e *InputHTMLElement) RemoveFORM(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeFORMKey)
	return e
}

// FORMACTION sets the "formaction" attribute.
// URL to use for form submission
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *InputHTMLElement) FORMACTION(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMACTIONKey] = v
	return e
}

func (e *InputHTMLElement) IfFORMACTION(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.FORMACTION(v)
}

func (e *InputHTMLElement) RemoveFORMACTION(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeFORMACTIONKey)
	return e
}

// FORMENCTYPE sets the "formenctype" attribute.
// Entry list encoding type to use for form submission
// Values values are constrained to:
//   - application/x_www_form_urlencoded
//   - multipart/form_data
//   - text/plain
func (e *InputHTMLElement) FORMENCTYPE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMENCTYPEKey] = v
	return e
}

func (e *InputHTMLElement) IfFORMENCTYPE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.FORMENCTYPE(v)
}

func (e *InputHTMLElement) RemoveFORMENCTYPE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeFORMENCTYPEKey)
	return e
}

// FORMMETHOD sets the "formmethod" attribute.
// Variant to use for form submission
// Values values are constrained to:
//   - dialog
//   - get
//   - post
func (e *InputHTMLElement) FORMMETHOD(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMMETHODKey] = v
	return e
}

func (e *InputHTMLElement) IfFORMMETHOD(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.FORMMETHOD(v)
}

func (e *InputHTMLElement) RemoveFORMMETHOD(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeFORMMETHODKey)
	return e
}

// FORMNOVALIDATE sets the "formnovalidate" attribute.
// Bypass form control validation for form submission
// Values values are constrained to:
//   - boolean_attribute
func (e *InputHTMLElement) FORMNOVALIDATE() *InputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeFORMNOVALIDATEKey] = struct{}{}
	return e
}

func (e *InputHTMLElement) IfFORMNOVALIDATE(cond bool) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.FORMNOVALIDATE()
}

func (e *InputHTMLElement) RemoveFORMNOVALIDATE() *InputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeFORMNOVALIDATEKey)
	return e
}

func (e *InputHTMLElement) SetFORMNOVALIDATE(b bool) *InputHTMLElement {
	if b {
		return e.FORMNOVALIDATE()
	}
	return e.RemoveFORMNOVALIDATE()
}

// FORMTARGET sets the "formtarget" attribute.
// Navigable for form submission
// Values values are constrained to:
//   - valid_navigable_target_name_or_keyword
func (e *InputHTMLElement) FORMTARGET(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMTARGETKey] = v
	return e
}

func (e *InputHTMLElement) IfFORMTARGET(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.FORMTARGET(v)
}

func (e *InputHTMLElement) RemoveFORMTARGET(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeFORMTARGETKey)
	return e
}

// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *InputHTMLElement) HEIGHT(v int) *InputHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeHEIGHTKey] = v
	return e
}

func (e *InputHTMLElement) IfHEIGHT(cond bool, v int) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.HEIGHT(v)
}

func (e *InputHTMLElement) RemoveHEIGHT(v int) *InputHTMLElement {
	delete(e.IntAttributes, attributeHEIGHTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *InputHTMLElement) HIDDEN(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *InputHTMLElement) IfHIDDEN(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *InputHTMLElement) RemoveHIDDEN(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *InputHTMLElement) ID(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *InputHTMLElement) IfID(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *InputHTMLElement) RemoveID(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *InputHTMLElement) INERT() *InputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *InputHTMLElement) IfINERT(cond bool) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *InputHTMLElement) RemoveINERT() *InputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *InputHTMLElement) SetINERT(b bool) *InputHTMLElement {
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
func (e *InputHTMLElement) INPUTMODE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *InputHTMLElement) IfINPUTMODE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *InputHTMLElement) RemoveINPUTMODE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *InputHTMLElement) IS(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *InputHTMLElement) IfIS(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *InputHTMLElement) RemoveIS(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *InputHTMLElement) ITEMID(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *InputHTMLElement) IfITEMID(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *InputHTMLElement) RemoveITEMID(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *InputHTMLElement) ITEMPROP(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *InputHTMLElement) IfITEMPROP(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *InputHTMLElement) RemoveITEMPROP(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *InputHTMLElement) ITEMREF(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *InputHTMLElement) IfITEMREF(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *InputHTMLElement) RemoveITEMREF(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *InputHTMLElement) ITEMSCOPE() *InputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *InputHTMLElement) IfITEMSCOPE(cond bool) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *InputHTMLElement) RemoveITEMSCOPE() *InputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *InputHTMLElement) SetITEMSCOPE(b bool) *InputHTMLElement {
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
func (e *InputHTMLElement) ITEMTYPE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *InputHTMLElement) IfITEMTYPE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *InputHTMLElement) RemoveITEMTYPE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *InputHTMLElement) LANG(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *InputHTMLElement) IfLANG(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *InputHTMLElement) RemoveLANG(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// LIST sets the "list" attribute.
// List of autocomplete options
// Values values are constrained to:
//   - id
func (e *InputHTMLElement) LIST(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLISTKey] = v
	return e
}

func (e *InputHTMLElement) IfLIST(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.LIST(v)
}

func (e *InputHTMLElement) RemoveLIST(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeLISTKey)
	return e
}

// MAX sets the "max" attribute.
// Upper bound of range
// Values values are constrained to:
//   - valid_floating_point_number
func (e *InputHTMLElement) MAX(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMAXKey] = v
	return e
}

func (e *InputHTMLElement) IfMAX(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.MAX(v)
}

func (e *InputHTMLElement) RemoveMAX(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeMAXKey)
	return e
}

// MAXLENGTH sets the "maxlength" attribute.
// Maximum length of value
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *InputHTMLElement) MAXLENGTH(v int) *InputHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeMAXLENGTHKey] = v
	return e
}

func (e *InputHTMLElement) IfMAXLENGTH(cond bool, v int) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.MAXLENGTH(v)
}

func (e *InputHTMLElement) RemoveMAXLENGTH(v int) *InputHTMLElement {
	delete(e.IntAttributes, attributeMAXLENGTHKey)
	return e
}

// MIN sets the "min" attribute.
// Lower bound of range
// Values values are constrained to:
//   - valid_floating_point_number
func (e *InputHTMLElement) MIN(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMINKey] = v
	return e
}

func (e *InputHTMLElement) IfMIN(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.MIN(v)
}

func (e *InputHTMLElement) RemoveMIN(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeMINKey)
	return e
}

// MINLENGTH sets the "minlength" attribute.
// Minimum length of value
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *InputHTMLElement) MINLENGTH(v int) *InputHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeMINLENGTHKey] = v
	return e
}

func (e *InputHTMLElement) IfMINLENGTH(cond bool, v int) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.MINLENGTH(v)
}

func (e *InputHTMLElement) RemoveMINLENGTH(v int) *InputHTMLElement {
	delete(e.IntAttributes, attributeMINLENGTHKey)
	return e
}

// MULTIPLE sets the "multiple" attribute.
// Whether to allow multiple values
// Values values are constrained to:
//   - boolean_attribute
func (e *InputHTMLElement) MULTIPLE() *InputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeMULTIPLEKey] = struct{}{}
	return e
}

func (e *InputHTMLElement) IfMULTIPLE(cond bool) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.MULTIPLE()
}

func (e *InputHTMLElement) RemoveMULTIPLE() *InputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeMULTIPLEKey)
	return e
}

func (e *InputHTMLElement) SetMULTIPLE(b bool) *InputHTMLElement {
	if b {
		return e.MULTIPLE()
	}
	return e.RemoveMULTIPLE()
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *InputHTMLElement) NAME(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *InputHTMLElement) IfNAME(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *InputHTMLElement) RemoveNAME(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *InputHTMLElement) NONCE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *InputHTMLElement) IfNONCE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *InputHTMLElement) RemoveNONCE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// PATTERN sets the "pattern" attribute.
// Pattern to be matched by the form control's value
// Values values are constrained to:
//   - pattern
func (e *InputHTMLElement) PATTERN(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePATTERNKey] = v
	return e
}

func (e *InputHTMLElement) IfPATTERN(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.PATTERN(v)
}

func (e *InputHTMLElement) RemovePATTERN(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributePATTERNKey)
	return e
}

// PLACEHOLDER sets the "placeholder" attribute.
// User-visible label to be placed within the form control
// Values values are constrained to:
//   - text
func (e *InputHTMLElement) PLACEHOLDER(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePLACEHOLDERKey] = v
	return e
}

func (e *InputHTMLElement) IfPLACEHOLDER(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.PLACEHOLDER(v)
}

func (e *InputHTMLElement) RemovePLACEHOLDER(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributePLACEHOLDERKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *InputHTMLElement) POPOVER(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *InputHTMLElement) IfPOPOVER(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *InputHTMLElement) RemovePOPOVER(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// POPOVERTARGET sets the "popovertarget" attribute.
// Targets a popover element to toggle, show, or hide
// Values values are constrained to:
//   - id
func (e *InputHTMLElement) POPOVERTARGET(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERTARGETKey] = v
	return e
}

func (e *InputHTMLElement) IfPOPOVERTARGET(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVERTARGET(v)
}

func (e *InputHTMLElement) RemovePOPOVERTARGET(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributePOPOVERTARGETKey)
	return e
}

// POPOVERTARGETACTION sets the "popovertargetaction" attribute.
// Indicates whether a targeted popover element is to be toggled, shown, or hidden
// Values values are constrained to:
//   - hide
//   - show
//   - toggle
func (e *InputHTMLElement) POPOVERTARGETACTION(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERTARGETACTIONKey] = v
	return e
}

func (e *InputHTMLElement) IfPOPOVERTARGETACTION(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVERTARGETACTION(v)
}

func (e *InputHTMLElement) RemovePOPOVERTARGETACTION(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributePOPOVERTARGETACTIONKey)
	return e
}

// READONLY sets the "readonly" attribute.
// Affects willValidate, plus any behavior added by the custom element author
// Values values are constrained to:
//   - boolean_attribute
func (e *InputHTMLElement) READONLY() *InputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeREADONLYKey] = struct{}{}
	return e
}

func (e *InputHTMLElement) IfREADONLY(cond bool) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.READONLY()
}

func (e *InputHTMLElement) RemoveREADONLY() *InputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeREADONLYKey)
	return e
}

func (e *InputHTMLElement) SetREADONLY(b bool) *InputHTMLElement {
	if b {
		return e.READONLY()
	}
	return e.RemoveREADONLY()
}

// REQUIRED sets the "required" attribute.
// Whether the control is required for form submission
// Values values are constrained to:
//   - boolean_attribute
func (e *InputHTMLElement) REQUIRED() *InputHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeREQUIREDKey] = struct{}{}
	return e
}

func (e *InputHTMLElement) IfREQUIRED(cond bool) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.REQUIRED()
}

func (e *InputHTMLElement) RemoveREQUIRED() *InputHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeREQUIREDKey)
	return e
}

func (e *InputHTMLElement) SetREQUIRED(b bool) *InputHTMLElement {
	if b {
		return e.REQUIRED()
	}
	return e.RemoveREQUIRED()
}

// SIZE sets the "size" attribute.
// Size of the control
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *InputHTMLElement) SIZE(v int) *InputHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeSIZEKey] = v
	return e
}

func (e *InputHTMLElement) IfSIZE(cond bool, v int) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.SIZE(v)
}

func (e *InputHTMLElement) RemoveSIZE(v int) *InputHTMLElement {
	delete(e.IntAttributes, attributeSIZEKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *InputHTMLElement) SLOT(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *InputHTMLElement) IfSLOT(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *InputHTMLElement) RemoveSLOT(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *InputHTMLElement) SPELLCHECK(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *InputHTMLElement) IfSPELLCHECK(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *InputHTMLElement) RemoveSPELLCHECK(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *InputHTMLElement) SRC(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCKey] = v
	return e
}

func (e *InputHTMLElement) IfSRC(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.SRC(v)
}

func (e *InputHTMLElement) RemoveSRC(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeSRCKey)
	return e
}

// STEP sets the "step" attribute.
// Granularity to be matched by the form control's value
// Values values are constrained to:
//   - any
//   - valid_floating_point_number
func (e *InputHTMLElement) STEP(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSTEPKey] = v
	return e
}

func (e *InputHTMLElement) IfSTEP(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.STEP(v)
}

func (e *InputHTMLElement) RemoveSTEP(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeSTEPKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *InputHTMLElement) STYLE(k, v string) *InputHTMLElement {
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

func (e *InputHTMLElement) IfSTYLE(cond bool, k string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *InputHTMLElement) RemoveSTYLE(k string) *InputHTMLElement {
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
func (e *InputHTMLElement) TABINDEX(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *InputHTMLElement) IfTABINDEX(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *InputHTMLElement) RemoveTABINDEX(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *InputHTMLElement) TITLE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *InputHTMLElement) IfTITLE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *InputHTMLElement) RemoveTITLE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *InputHTMLElement) TRANSLATE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *InputHTMLElement) IfTRANSLATE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *InputHTMLElement) RemoveTRANSLATE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//   - java_script_mime_type_essence_match
//   - module
//   - valid_mime_type_string
func (e *InputHTMLElement) TYPE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTYPEKey] = v
	return e
}

func (e *InputHTMLElement) IfTYPE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.TYPE(v)
}

func (e *InputHTMLElement) RemoveTYPE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeTYPEKey)
	return e
}

// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//   - valid_floating_point_number
func (e *InputHTMLElement) VALUE(v string) *InputHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeVALUEKey] = v
	return e
}

func (e *InputHTMLElement) IfVALUE(cond bool, v string) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.VALUE(v)
}

func (e *InputHTMLElement) RemoveVALUE(v string) *InputHTMLElement {
	delete(e.StringAttributes, attributeVALUEKey)
	return e
}

// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *InputHTMLElement) WIDTH(v int) *InputHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeWIDTHKey] = v
	return e
}

func (e *InputHTMLElement) IfWIDTH(cond bool, v int) *InputHTMLElement {
	if !cond {
		return e
	}
	return e.WIDTH(v)
}

func (e *InputHTMLElement) RemoveWIDTH(v int) *InputHTMLElement {
	delete(e.IntAttributes, attributeWIDTHKey)
	return e
}
