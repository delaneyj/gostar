package html

import (
	"fmt"
)

type FieldsetHTMLElement struct {
	*Element
}

func FIELDSET(children ...ElementBuilder) *FieldsetHTMLElement {
	return &FieldsetHTMLElement{
		Element: &Element{
			Tag:           elementTagFIELDSET,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *FieldsetHTMLElement) Children(children ...ElementBuilder) *FieldsetHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *FieldsetHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *FieldsetHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *FieldsetHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *FieldsetHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *FieldsetHTMLElement) Text(text string) *FieldsetHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *FieldsetHTMLElement) TextF(format string, args ...any) *FieldsetHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *FieldsetHTMLElement) Escaped(text string) *FieldsetHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *FieldsetHTMLElement) EscapedF(format string, args ...any) *FieldsetHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *FieldsetHTMLElement) CustomData(key, value string) *FieldsetHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *FieldsetHTMLElement) CustomDataRemove(key string) *FieldsetHTMLElement {
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
func (e *FieldsetHTMLElement) ACCESSKEY(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfACCESSKEY(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *FieldsetHTMLElement) RemoveACCESSKEY(v string) *FieldsetHTMLElement {
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
func (e *FieldsetHTMLElement) AUTOCAPITALIZE(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *FieldsetHTMLElement) RemoveAUTOCAPITALIZE(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *FieldsetHTMLElement) AUTOFOCUS() *FieldsetHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *FieldsetHTMLElement) IfAUTOFOCUS(cond bool) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *FieldsetHTMLElement) RemoveAUTOFOCUS() *FieldsetHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *FieldsetHTMLElement) SetAUTOFOCUS(b bool) *FieldsetHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *FieldsetHTMLElement) CLASS(v string) *FieldsetHTMLElement {
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

func (e *FieldsetHTMLElement) IfCLASS(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *FieldsetHTMLElement) SetCLASS(v string) *FieldsetHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *FieldsetHTMLElement) RemoveCLASS(v string) *FieldsetHTMLElement {
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
func (e *FieldsetHTMLElement) CONTENTEDITABLE(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *FieldsetHTMLElement) RemoveCONTENTEDITABLE(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *FieldsetHTMLElement) DIR(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfDIR(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *FieldsetHTMLElement) RemoveDIR(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//   - boolean_attribute
func (e *FieldsetHTMLElement) DISABLED() *FieldsetHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeDISABLEDKey] = struct{}{}
	return e
}

func (e *FieldsetHTMLElement) IfDISABLED(cond bool) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.DISABLED()
}

func (e *FieldsetHTMLElement) RemoveDISABLED() *FieldsetHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeDISABLEDKey)
	return e
}

func (e *FieldsetHTMLElement) SetDISABLED(b bool) *FieldsetHTMLElement {
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
func (e *FieldsetHTMLElement) DRAGGABLE(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfDRAGGABLE(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *FieldsetHTMLElement) RemoveDRAGGABLE(v string) *FieldsetHTMLElement {
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
func (e *FieldsetHTMLElement) ENTERKEYHINT(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfENTERKEYHINT(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *FieldsetHTMLElement) RemoveENTERKEYHINT(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//   - id
func (e *FieldsetHTMLElement) FORM(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfFORM(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.FORM(v)
}

func (e *FieldsetHTMLElement) RemoveFORM(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeFORMKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *FieldsetHTMLElement) HIDDEN(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfHIDDEN(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *FieldsetHTMLElement) RemoveHIDDEN(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *FieldsetHTMLElement) ID(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfID(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *FieldsetHTMLElement) RemoveID(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *FieldsetHTMLElement) INERT() *FieldsetHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *FieldsetHTMLElement) IfINERT(cond bool) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *FieldsetHTMLElement) RemoveINERT() *FieldsetHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *FieldsetHTMLElement) SetINERT(b bool) *FieldsetHTMLElement {
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
func (e *FieldsetHTMLElement) INPUTMODE(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfINPUTMODE(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *FieldsetHTMLElement) RemoveINPUTMODE(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *FieldsetHTMLElement) IS(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfIS(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *FieldsetHTMLElement) RemoveIS(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *FieldsetHTMLElement) ITEMID(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfITEMID(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *FieldsetHTMLElement) RemoveITEMID(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *FieldsetHTMLElement) ITEMPROP(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfITEMPROP(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *FieldsetHTMLElement) RemoveITEMPROP(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *FieldsetHTMLElement) ITEMREF(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfITEMREF(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *FieldsetHTMLElement) RemoveITEMREF(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *FieldsetHTMLElement) ITEMSCOPE() *FieldsetHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *FieldsetHTMLElement) IfITEMSCOPE(cond bool) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *FieldsetHTMLElement) RemoveITEMSCOPE() *FieldsetHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *FieldsetHTMLElement) SetITEMSCOPE(b bool) *FieldsetHTMLElement {
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
func (e *FieldsetHTMLElement) ITEMTYPE(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfITEMTYPE(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *FieldsetHTMLElement) RemoveITEMTYPE(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *FieldsetHTMLElement) LANG(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfLANG(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *FieldsetHTMLElement) RemoveLANG(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *FieldsetHTMLElement) NAME(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfNAME(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *FieldsetHTMLElement) RemoveNAME(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *FieldsetHTMLElement) NONCE(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfNONCE(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *FieldsetHTMLElement) RemoveNONCE(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *FieldsetHTMLElement) POPOVER(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfPOPOVER(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *FieldsetHTMLElement) RemovePOPOVER(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *FieldsetHTMLElement) SLOT(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfSLOT(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *FieldsetHTMLElement) RemoveSLOT(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *FieldsetHTMLElement) SPELLCHECK(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfSPELLCHECK(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *FieldsetHTMLElement) RemoveSPELLCHECK(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *FieldsetHTMLElement) STYLE(k, v string) *FieldsetHTMLElement {
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

func (e *FieldsetHTMLElement) IfSTYLE(cond bool, k string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *FieldsetHTMLElement) RemoveSTYLE(k string) *FieldsetHTMLElement {
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
func (e *FieldsetHTMLElement) TABINDEX(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfTABINDEX(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *FieldsetHTMLElement) RemoveTABINDEX(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *FieldsetHTMLElement) TITLE(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfTITLE(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *FieldsetHTMLElement) RemoveTITLE(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *FieldsetHTMLElement) TRANSLATE(v string) *FieldsetHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *FieldsetHTMLElement) IfTRANSLATE(cond bool, v string) *FieldsetHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *FieldsetHTMLElement) RemoveTRANSLATE(v string) *FieldsetHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
