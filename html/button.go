package html

import (
	"fmt"
)

type ButtonHTMLElement struct {
	*Element
}

func BUTTON(children ...ElementBuilder) *ButtonHTMLElement {
	return &ButtonHTMLElement{
		Element: &Element{
			Tag:           elementTagBUTTON,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *ButtonHTMLElement) Children(children ...ElementBuilder) *ButtonHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *ButtonHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ButtonHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *ButtonHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ButtonHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *ButtonHTMLElement) Text(text string) *ButtonHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *ButtonHTMLElement) TextF(format string, args ...any) *ButtonHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *ButtonHTMLElement) Escaped(text string) *ButtonHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *ButtonHTMLElement) EscapedF(format string, args ...any) *ButtonHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *ButtonHTMLElement) CustomData(key, value string) *ButtonHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ButtonHTMLElement) CustomDataRemove(key string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) ACCESSKEY(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *ButtonHTMLElement) IfACCESSKEY(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *ButtonHTMLElement) RemoveACCESSKEY(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) AUTOCAPITALIZE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *ButtonHTMLElement) RemoveAUTOCAPITALIZE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *ButtonHTMLElement) AUTOFOCUS() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *ButtonHTMLElement) IfAUTOFOCUS(cond bool) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *ButtonHTMLElement) RemoveAUTOFOCUS() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *ButtonHTMLElement) SetAUTOFOCUS(b bool) *ButtonHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *ButtonHTMLElement) CLASS(v string) *ButtonHTMLElement {
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

func (e *ButtonHTMLElement) IfCLASS(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *ButtonHTMLElement) SetCLASS(v string) *ButtonHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *ButtonHTMLElement) RemoveCLASS(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) CONTENTEDITABLE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *ButtonHTMLElement) RemoveCONTENTEDITABLE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *ButtonHTMLElement) DIR(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *ButtonHTMLElement) IfDIR(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *ButtonHTMLElement) RemoveDIR(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//   - boolean_attribute
func (e *ButtonHTMLElement) DISABLED() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeDISABLEDKey] = struct{}{}
	return e
}

func (e *ButtonHTMLElement) IfDISABLED(cond bool) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.DISABLED()
}

func (e *ButtonHTMLElement) RemoveDISABLED() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeDISABLEDKey)
	return e
}

func (e *ButtonHTMLElement) SetDISABLED(b bool) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) DRAGGABLE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfDRAGGABLE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *ButtonHTMLElement) RemoveDRAGGABLE(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) ENTERKEYHINT(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *ButtonHTMLElement) IfENTERKEYHINT(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *ButtonHTMLElement) RemoveENTERKEYHINT(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//   - id
func (e *ButtonHTMLElement) FORM(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMKey] = v
	return e
}

func (e *ButtonHTMLElement) IfFORM(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.FORM(v)
}

func (e *ButtonHTMLElement) RemoveFORM(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeFORMKey)
	return e
}

// FORMACTION sets the "formaction" attribute.
// URL to use for form submission
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ButtonHTMLElement) FORMACTION(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMACTIONKey] = v
	return e
}

func (e *ButtonHTMLElement) IfFORMACTION(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.FORMACTION(v)
}

func (e *ButtonHTMLElement) RemoveFORMACTION(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeFORMACTIONKey)
	return e
}

// FORMENCTYPE sets the "formenctype" attribute.
// Entry list encoding type to use for form submission
// Values values are constrained to:
//   - application/x_www_form_urlencoded
//   - multipart/form_data
//   - text/plain
func (e *ButtonHTMLElement) FORMENCTYPE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMENCTYPEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfFORMENCTYPE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.FORMENCTYPE(v)
}

func (e *ButtonHTMLElement) RemoveFORMENCTYPE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeFORMENCTYPEKey)
	return e
}

// FORMMETHOD sets the "formmethod" attribute.
// Variant to use for form submission
// Values values are constrained to:
//   - dialog
//   - get
//   - post
func (e *ButtonHTMLElement) FORMMETHOD(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMMETHODKey] = v
	return e
}

func (e *ButtonHTMLElement) IfFORMMETHOD(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.FORMMETHOD(v)
}

func (e *ButtonHTMLElement) RemoveFORMMETHOD(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeFORMMETHODKey)
	return e
}

// FORMNOVALIDATE sets the "formnovalidate" attribute.
// Bypass form control validation for form submission
// Values values are constrained to:
//   - boolean_attribute
func (e *ButtonHTMLElement) FORMNOVALIDATE() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeFORMNOVALIDATEKey] = struct{}{}
	return e
}

func (e *ButtonHTMLElement) IfFORMNOVALIDATE(cond bool) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.FORMNOVALIDATE()
}

func (e *ButtonHTMLElement) RemoveFORMNOVALIDATE() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeFORMNOVALIDATEKey)
	return e
}

func (e *ButtonHTMLElement) SetFORMNOVALIDATE(b bool) *ButtonHTMLElement {
	if b {
		return e.FORMNOVALIDATE()
	}
	return e.RemoveFORMNOVALIDATE()
}

// FORMTARGET sets the "formtarget" attribute.
// Navigable for form submission
// Values values are constrained to:
//   - valid_navigable_target_name_or_keyword
func (e *ButtonHTMLElement) FORMTARGET(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMTARGETKey] = v
	return e
}

func (e *ButtonHTMLElement) IfFORMTARGET(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.FORMTARGET(v)
}

func (e *ButtonHTMLElement) RemoveFORMTARGET(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeFORMTARGETKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *ButtonHTMLElement) HIDDEN(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *ButtonHTMLElement) IfHIDDEN(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *ButtonHTMLElement) RemoveHIDDEN(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *ButtonHTMLElement) ID(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *ButtonHTMLElement) IfID(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *ButtonHTMLElement) RemoveID(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *ButtonHTMLElement) INERT() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *ButtonHTMLElement) IfINERT(cond bool) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *ButtonHTMLElement) RemoveINERT() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *ButtonHTMLElement) SetINERT(b bool) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) INPUTMODE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfINPUTMODE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *ButtonHTMLElement) RemoveINPUTMODE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *ButtonHTMLElement) IS(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *ButtonHTMLElement) IfIS(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *ButtonHTMLElement) RemoveIS(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *ButtonHTMLElement) ITEMID(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *ButtonHTMLElement) IfITEMID(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *ButtonHTMLElement) RemoveITEMID(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *ButtonHTMLElement) ITEMPROP(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *ButtonHTMLElement) IfITEMPROP(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *ButtonHTMLElement) RemoveITEMPROP(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ButtonHTMLElement) ITEMREF(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *ButtonHTMLElement) IfITEMREF(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *ButtonHTMLElement) RemoveITEMREF(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *ButtonHTMLElement) ITEMSCOPE() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *ButtonHTMLElement) IfITEMSCOPE(cond bool) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *ButtonHTMLElement) RemoveITEMSCOPE() *ButtonHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *ButtonHTMLElement) SetITEMSCOPE(b bool) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) ITEMTYPE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfITEMTYPE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *ButtonHTMLElement) RemoveITEMTYPE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ButtonHTMLElement) LANG(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *ButtonHTMLElement) IfLANG(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *ButtonHTMLElement) RemoveLANG(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *ButtonHTMLElement) NAME(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfNAME(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *ButtonHTMLElement) RemoveNAME(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *ButtonHTMLElement) NONCE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfNONCE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *ButtonHTMLElement) RemoveNONCE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *ButtonHTMLElement) POPOVER(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *ButtonHTMLElement) IfPOPOVER(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *ButtonHTMLElement) RemovePOPOVER(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// POPOVERTARGET sets the "popovertarget" attribute.
// Targets a popover element to toggle, show, or hide
// Values values are constrained to:
//   - id
func (e *ButtonHTMLElement) POPOVERTARGET(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERTARGETKey] = v
	return e
}

func (e *ButtonHTMLElement) IfPOPOVERTARGET(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVERTARGET(v)
}

func (e *ButtonHTMLElement) RemovePOPOVERTARGET(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributePOPOVERTARGETKey)
	return e
}

// POPOVERTARGETACTION sets the "popovertargetaction" attribute.
// Indicates whether a targeted popover element is to be toggled, shown, or hidden
// Values values are constrained to:
//   - hide
//   - show
//   - toggle
func (e *ButtonHTMLElement) POPOVERTARGETACTION(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERTARGETACTIONKey] = v
	return e
}

func (e *ButtonHTMLElement) IfPOPOVERTARGETACTION(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVERTARGETACTION(v)
}

func (e *ButtonHTMLElement) RemovePOPOVERTARGETACTION(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributePOPOVERTARGETACTIONKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *ButtonHTMLElement) SLOT(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *ButtonHTMLElement) IfSLOT(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *ButtonHTMLElement) RemoveSLOT(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *ButtonHTMLElement) SPELLCHECK(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *ButtonHTMLElement) IfSPELLCHECK(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *ButtonHTMLElement) RemoveSPELLCHECK(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ButtonHTMLElement) STYLE(k, v string) *ButtonHTMLElement {
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

func (e *ButtonHTMLElement) IfSTYLE(cond bool, k string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *ButtonHTMLElement) RemoveSTYLE(k string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) TABINDEX(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *ButtonHTMLElement) IfTABINDEX(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *ButtonHTMLElement) RemoveTABINDEX(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *ButtonHTMLElement) TITLE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfTITLE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *ButtonHTMLElement) RemoveTITLE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *ButtonHTMLElement) TRANSLATE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfTRANSLATE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *ButtonHTMLElement) RemoveTRANSLATE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//   - java_script_mime_type_essence_match
//   - module
//   - valid_mime_type_string
func (e *ButtonHTMLElement) TYPE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTYPEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfTYPE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.TYPE(v)
}

func (e *ButtonHTMLElement) RemoveTYPE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeTYPEKey)
	return e
}

// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//   - valid_floating_point_number
func (e *ButtonHTMLElement) VALUE(v string) *ButtonHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeVALUEKey] = v
	return e
}

func (e *ButtonHTMLElement) IfVALUE(cond bool, v string) *ButtonHTMLElement {
	if !cond {
		return e
	}
	return e.VALUE(v)
}

func (e *ButtonHTMLElement) RemoveVALUE(v string) *ButtonHTMLElement {
	delete(e.StringAttributes, attributeVALUEKey)
	return e
}
