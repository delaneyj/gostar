package html

import (
	"fmt"
)

type SelectHTMLElement struct {
	*Element
}

func SELECT(children ...ElementBuilder) *SelectHTMLElement {
	return &SelectHTMLElement{
		Element: &Element{
			Tag:           elementTagSELECT,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SelectHTMLElement) Children(children ...ElementBuilder) *SelectHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SelectHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SelectHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SelectHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SelectHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SelectHTMLElement) Text(text string) *SelectHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SelectHTMLElement) TextF(format string, args ...any) *SelectHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SelectHTMLElement) Escaped(text string) *SelectHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SelectHTMLElement) EscapedF(format string, args ...any) *SelectHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SelectHTMLElement) CustomData(key, value string) *SelectHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SelectHTMLElement) CustomDataRemove(key string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) ACCESSKEY(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SelectHTMLElement) IfACCESSKEY(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SelectHTMLElement) RemoveACCESSKEY(v string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) AUTOCAPITALIZE(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SelectHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SelectHTMLElement) RemoveAUTOCAPITALIZE(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//   - autofill_field
func (e *SelectHTMLElement) AUTOCOMPLETE(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCOMPLETEKey] = v
	return e
}

func (e *SelectHTMLElement) IfAUTOCOMPLETE(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCOMPLETE(v)
}

func (e *SelectHTMLElement) RemoveAUTOCOMPLETE(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeAUTOCOMPLETEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SelectHTMLElement) AUTOFOCUS() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SelectHTMLElement) IfAUTOFOCUS(cond bool) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SelectHTMLElement) RemoveAUTOFOCUS() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SelectHTMLElement) SetAUTOFOCUS(b bool) *SelectHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SelectHTMLElement) CLASS(v string) *SelectHTMLElement {
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

func (e *SelectHTMLElement) IfCLASS(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SelectHTMLElement) SetCLASS(v string) *SelectHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SelectHTMLElement) RemoveCLASS(v string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) CONTENTEDITABLE(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SelectHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SelectHTMLElement) RemoveCONTENTEDITABLE(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *SelectHTMLElement) DIR(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SelectHTMLElement) IfDIR(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SelectHTMLElement) RemoveDIR(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//   - boolean_attribute
func (e *SelectHTMLElement) DISABLED() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeDISABLEDKey] = struct{}{}
	return e
}

func (e *SelectHTMLElement) IfDISABLED(cond bool) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.DISABLED()
}

func (e *SelectHTMLElement) RemoveDISABLED() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeDISABLEDKey)
	return e
}

func (e *SelectHTMLElement) SetDISABLED(b bool) *SelectHTMLElement {
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
func (e *SelectHTMLElement) DRAGGABLE(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SelectHTMLElement) IfDRAGGABLE(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SelectHTMLElement) RemoveDRAGGABLE(v string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) ENTERKEYHINT(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SelectHTMLElement) IfENTERKEYHINT(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SelectHTMLElement) RemoveENTERKEYHINT(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//   - id
func (e *SelectHTMLElement) FORM(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMKey] = v
	return e
}

func (e *SelectHTMLElement) IfFORM(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.FORM(v)
}

func (e *SelectHTMLElement) RemoveFORM(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeFORMKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *SelectHTMLElement) HIDDEN(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SelectHTMLElement) IfHIDDEN(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SelectHTMLElement) RemoveHIDDEN(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SelectHTMLElement) ID(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SelectHTMLElement) IfID(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SelectHTMLElement) RemoveID(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SelectHTMLElement) INERT() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SelectHTMLElement) IfINERT(cond bool) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SelectHTMLElement) RemoveINERT() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SelectHTMLElement) SetINERT(b bool) *SelectHTMLElement {
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
func (e *SelectHTMLElement) INPUTMODE(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SelectHTMLElement) IfINPUTMODE(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SelectHTMLElement) RemoveINPUTMODE(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *SelectHTMLElement) IS(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SelectHTMLElement) IfIS(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SelectHTMLElement) RemoveIS(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SelectHTMLElement) ITEMID(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SelectHTMLElement) IfITEMID(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SelectHTMLElement) RemoveITEMID(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *SelectHTMLElement) ITEMPROP(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SelectHTMLElement) IfITEMPROP(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SelectHTMLElement) RemoveITEMPROP(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SelectHTMLElement) ITEMREF(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SelectHTMLElement) IfITEMREF(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SelectHTMLElement) RemoveITEMREF(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SelectHTMLElement) ITEMSCOPE() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SelectHTMLElement) IfITEMSCOPE(cond bool) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SelectHTMLElement) RemoveITEMSCOPE() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SelectHTMLElement) SetITEMSCOPE(b bool) *SelectHTMLElement {
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
func (e *SelectHTMLElement) ITEMTYPE(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SelectHTMLElement) IfITEMTYPE(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SelectHTMLElement) RemoveITEMTYPE(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SelectHTMLElement) LANG(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SelectHTMLElement) IfLANG(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SelectHTMLElement) RemoveLANG(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// MULTIPLE sets the "multiple" attribute.
// Whether to allow multiple values
// Values values are constrained to:
//   - boolean_attribute
func (e *SelectHTMLElement) MULTIPLE() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeMULTIPLEKey] = struct{}{}
	return e
}

func (e *SelectHTMLElement) IfMULTIPLE(cond bool) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.MULTIPLE()
}

func (e *SelectHTMLElement) RemoveMULTIPLE() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeMULTIPLEKey)
	return e
}

func (e *SelectHTMLElement) SetMULTIPLE(b bool) *SelectHTMLElement {
	if b {
		return e.MULTIPLE()
	}
	return e.RemoveMULTIPLE()
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *SelectHTMLElement) NAME(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *SelectHTMLElement) IfNAME(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *SelectHTMLElement) RemoveNAME(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SelectHTMLElement) NONCE(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SelectHTMLElement) IfNONCE(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SelectHTMLElement) RemoveNONCE(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *SelectHTMLElement) POPOVER(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SelectHTMLElement) IfPOPOVER(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SelectHTMLElement) RemovePOPOVER(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// REQUIRED sets the "required" attribute.
// Whether the control is required for form submission
// Values values are constrained to:
//   - boolean_attribute
func (e *SelectHTMLElement) REQUIRED() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeREQUIREDKey] = struct{}{}
	return e
}

func (e *SelectHTMLElement) IfREQUIRED(cond bool) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.REQUIRED()
}

func (e *SelectHTMLElement) RemoveREQUIRED() *SelectHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeREQUIREDKey)
	return e
}

func (e *SelectHTMLElement) SetREQUIRED(b bool) *SelectHTMLElement {
	if b {
		return e.REQUIRED()
	}
	return e.RemoveREQUIRED()
}

// SIZE sets the "size" attribute.
// Size of the control
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *SelectHTMLElement) SIZE(v int) *SelectHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeSIZEKey] = v
	return e
}

func (e *SelectHTMLElement) IfSIZE(cond bool, v int) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.SIZE(v)
}

func (e *SelectHTMLElement) RemoveSIZE(v int) *SelectHTMLElement {
	delete(e.IntAttributes, attributeSIZEKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SelectHTMLElement) SLOT(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SelectHTMLElement) IfSLOT(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SelectHTMLElement) RemoveSLOT(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *SelectHTMLElement) SPELLCHECK(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SelectHTMLElement) IfSPELLCHECK(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SelectHTMLElement) RemoveSPELLCHECK(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SelectHTMLElement) STYLE(k, v string) *SelectHTMLElement {
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

func (e *SelectHTMLElement) IfSTYLE(cond bool, k string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SelectHTMLElement) RemoveSTYLE(k string) *SelectHTMLElement {
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
func (e *SelectHTMLElement) TABINDEX(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SelectHTMLElement) IfTABINDEX(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SelectHTMLElement) RemoveTABINDEX(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SelectHTMLElement) TITLE(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SelectHTMLElement) IfTITLE(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SelectHTMLElement) RemoveTITLE(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *SelectHTMLElement) TRANSLATE(v string) *SelectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SelectHTMLElement) IfTRANSLATE(cond bool, v string) *SelectHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SelectHTMLElement) RemoveTRANSLATE(v string) *SelectHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
