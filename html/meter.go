package html

import (
	"fmt"
)

type MeterHTMLElement struct {
	*Element
}

func METER(children ...ElementBuilder) *MeterHTMLElement {
	return &MeterHTMLElement{
		Element: &Element{
			Tag:           elementTagMETER,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *MeterHTMLElement) Children(children ...ElementBuilder) *MeterHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MeterHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *MeterHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MeterHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *MeterHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MeterHTMLElement) Text(text string) *MeterHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MeterHTMLElement) TextF(format string, args ...any) *MeterHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MeterHTMLElement) Escaped(text string) *MeterHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MeterHTMLElement) EscapedF(format string, args ...any) *MeterHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MeterHTMLElement) CustomData(key, value string) *MeterHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *MeterHTMLElement) CustomDataRemove(key string) *MeterHTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}

// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//   - ordered_set_of_unique_space_separated_tokens
//   - identical_to
func (e *MeterHTMLElement) ACCESSKEY(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *MeterHTMLElement) IfACCESSKEY(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *MeterHTMLElement) RemoveACCESSKEY(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// AUTOCAPITALIZE sets the "autocapitalize" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Values values are constrained to:
//   - on
//   - on
//   - off
//   - off
//   - none
//   - none
//   - sentences
//   - sentences
//   - words
//   - words
//   - characters
//   - characters
func (e *MeterHTMLElement) AUTOCAPITALIZE(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *MeterHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *MeterHTMLElement) RemoveAUTOCAPITALIZE(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *MeterHTMLElement) AUTOFOCUS() *MeterHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *MeterHTMLElement) IfAUTOFOCUS(cond bool) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *MeterHTMLElement) RemoveAUTOFOCUS() *MeterHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *MeterHTMLElement) SetAUTOFOCUS(b bool) *MeterHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *MeterHTMLElement) CLASS(v string) *MeterHTMLElement {
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

func (e *MeterHTMLElement) IfCLASS(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *MeterHTMLElement) SetCLASS(v string) *MeterHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *MeterHTMLElement) RemoveCLASS(v string) *MeterHTMLElement {
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
//   - true
//   - plaintext_only
//   - false
func (e *MeterHTMLElement) CONTENTEDITABLE(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *MeterHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *MeterHTMLElement) RemoveCONTENTEDITABLE(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (e *MeterHTMLElement) DIR(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *MeterHTMLElement) IfDIR(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *MeterHTMLElement) RemoveDIR(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *MeterHTMLElement) DRAGGABLE(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *MeterHTMLElement) IfDRAGGABLE(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *MeterHTMLElement) RemoveDRAGGABLE(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeDRAGGABLEKey)
	return e
}

// ENTERKEYHINT sets the "enterkeyhint" attribute.
// Hint for selecting an enter key action
// Values values are constrained to:
//   - enter
//   - enter
//   - done
//   - done
//   - go
//   - go
//   - next
//   - next
//   - previous
//   - previous
//   - search
//   - search
//   - send
//   - send
func (e *MeterHTMLElement) ENTERKEYHINT(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *MeterHTMLElement) IfENTERKEYHINT(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *MeterHTMLElement) RemoveENTERKEYHINT(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (e *MeterHTMLElement) HIDDEN(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *MeterHTMLElement) IfHIDDEN(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *MeterHTMLElement) RemoveHIDDEN(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// HIGH sets the "high" attribute.
// Low limit of high range
// Values values are constrained to:
//   - valid_floating_point_number
func (e *MeterHTMLElement) HIGH(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIGHKey] = v
	return e
}

func (e *MeterHTMLElement) IfHIGH(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.HIGH(v)
}

func (e *MeterHTMLElement) RemoveHIGH(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeHIGHKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *MeterHTMLElement) ID(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *MeterHTMLElement) IfID(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *MeterHTMLElement) RemoveID(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *MeterHTMLElement) INERT() *MeterHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *MeterHTMLElement) IfINERT(cond bool) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *MeterHTMLElement) RemoveINERT() *MeterHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *MeterHTMLElement) SetINERT(b bool) *MeterHTMLElement {
	if b {
		return e.INERT()
	}
	return e.RemoveINERT()
}

// INPUTMODE sets the "inputmode" attribute.
// Hint for selecting an input modality
// Values values are constrained to:
//   - none
//   - none
//   - text
//   - text
//   - tel
//   - tel
//   - email
//   - email
//   - url
//   - url
//   - numeric
//   - numeric
//   - decimal
//   - decimal
//   - search
//   - search
func (e *MeterHTMLElement) INPUTMODE(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *MeterHTMLElement) IfINPUTMODE(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *MeterHTMLElement) RemoveINPUTMODE(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *MeterHTMLElement) IS(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *MeterHTMLElement) IfIS(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *MeterHTMLElement) RemoveIS(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *MeterHTMLElement) ITEMID(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *MeterHTMLElement) IfITEMID(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *MeterHTMLElement) RemoveITEMID(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *MeterHTMLElement) ITEMPROP(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *MeterHTMLElement) IfITEMPROP(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *MeterHTMLElement) RemoveITEMPROP(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *MeterHTMLElement) ITEMREF(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *MeterHTMLElement) IfITEMREF(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *MeterHTMLElement) RemoveITEMREF(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *MeterHTMLElement) ITEMSCOPE() *MeterHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *MeterHTMLElement) IfITEMSCOPE(cond bool) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *MeterHTMLElement) RemoveITEMSCOPE() *MeterHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *MeterHTMLElement) SetITEMSCOPE(b bool) *MeterHTMLElement {
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
func (e *MeterHTMLElement) ITEMTYPE(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *MeterHTMLElement) IfITEMTYPE(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *MeterHTMLElement) RemoveITEMTYPE(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *MeterHTMLElement) LANG(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *MeterHTMLElement) IfLANG(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *MeterHTMLElement) RemoveLANG(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// LOW sets the "low" attribute.
// High limit of low range
// Values values are constrained to:
//   - valid_floating_point_number
func (e *MeterHTMLElement) LOW(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLOWKey] = v
	return e
}

func (e *MeterHTMLElement) IfLOW(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.LOW(v)
}

func (e *MeterHTMLElement) RemoveLOW(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeLOWKey)
	return e
}

// MAX sets the "max" attribute.
// Upper bound of range
// Values values are constrained to:
//   - valid_floating_point_number
func (e *MeterHTMLElement) MAX(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMAXKey] = v
	return e
}

func (e *MeterHTMLElement) IfMAX(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.MAX(v)
}

func (e *MeterHTMLElement) RemoveMAX(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeMAXKey)
	return e
}

// MIN sets the "min" attribute.
// Lower bound of range
// Values values are constrained to:
//   - valid_floating_point_number
func (e *MeterHTMLElement) MIN(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMINKey] = v
	return e
}

func (e *MeterHTMLElement) IfMIN(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.MIN(v)
}

func (e *MeterHTMLElement) RemoveMIN(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeMINKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *MeterHTMLElement) NONCE(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *MeterHTMLElement) IfNONCE(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *MeterHTMLElement) RemoveNONCE(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// OPTIMUM sets the "optimum" attribute.
// Optimum value in gauge
// Values values are constrained to:
//   - valid_floating_point_number
func (e *MeterHTMLElement) OPTIMUM(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeOPTIMUMKey] = v
	return e
}

func (e *MeterHTMLElement) IfOPTIMUM(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.OPTIMUM(v)
}

func (e *MeterHTMLElement) RemoveOPTIMUM(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeOPTIMUMKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - auto
//   - manual
//   - manual
func (e *MeterHTMLElement) POPOVER(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *MeterHTMLElement) IfPOPOVER(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *MeterHTMLElement) RemovePOPOVER(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *MeterHTMLElement) SLOT(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *MeterHTMLElement) IfSLOT(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *MeterHTMLElement) RemoveSLOT(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *MeterHTMLElement) SPELLCHECK(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *MeterHTMLElement) IfSPELLCHECK(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *MeterHTMLElement) RemoveSPELLCHECK(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *MeterHTMLElement) STYLE(k, v string) *MeterHTMLElement {
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

func (e *MeterHTMLElement) IfSTYLE(cond bool, k string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *MeterHTMLElement) RemoveSTYLE(k string) *MeterHTMLElement {
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
func (e *MeterHTMLElement) TABINDEX(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *MeterHTMLElement) IfTABINDEX(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *MeterHTMLElement) RemoveTABINDEX(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *MeterHTMLElement) TITLE(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *MeterHTMLElement) IfTITLE(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *MeterHTMLElement) RemoveTITLE(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *MeterHTMLElement) TRANSLATE(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *MeterHTMLElement) IfTRANSLATE(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *MeterHTMLElement) RemoveTRANSLATE(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//   - valid_floating_point_number
func (e *MeterHTMLElement) VALUE(v string) *MeterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeVALUEKey] = v
	return e
}

func (e *MeterHTMLElement) IfVALUE(cond bool, v string) *MeterHTMLElement {
	if !cond {
		return e
	}
	return e.VALUE(v)
}

func (e *MeterHTMLElement) RemoveVALUE(v string) *MeterHTMLElement {
	delete(e.StringAttributes, attributeVALUEKey)
	return e
}
