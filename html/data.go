package html

import (
	"fmt"
)

type DataHTMLElement struct {
	*Element
}

func DATA(children ...ElementBuilder) *DataHTMLElement {
	return &DataHTMLElement{
		Element: &Element{
			Tag:           elementTagDATA,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DataHTMLElement) Children(children ...ElementBuilder) *DataHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DataHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DataHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DataHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DataHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DataHTMLElement) Text(text string) *DataHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DataHTMLElement) TextF(format string, args ...any) *DataHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DataHTMLElement) Escaped(text string) *DataHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DataHTMLElement) EscapedF(format string, args ...any) *DataHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DataHTMLElement) CustomData(key, value string) *DataHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DataHTMLElement) CustomDataRemove(key string) *DataHTMLElement {
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
func (e *DataHTMLElement) ACCESSKEY(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DataHTMLElement) IfACCESSKEY(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DataHTMLElement) RemoveACCESSKEY(v string) *DataHTMLElement {
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
func (e *DataHTMLElement) AUTOCAPITALIZE(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DataHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DataHTMLElement) RemoveAUTOCAPITALIZE(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DataHTMLElement) AUTOFOCUS() *DataHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DataHTMLElement) IfAUTOFOCUS(cond bool) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DataHTMLElement) RemoveAUTOFOCUS() *DataHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DataHTMLElement) SetAUTOFOCUS(b bool) *DataHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DataHTMLElement) CLASS(v string) *DataHTMLElement {
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

func (e *DataHTMLElement) IfCLASS(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DataHTMLElement) SetCLASS(v string) *DataHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DataHTMLElement) RemoveCLASS(v string) *DataHTMLElement {
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
func (e *DataHTMLElement) CONTENTEDITABLE(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DataHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DataHTMLElement) RemoveCONTENTEDITABLE(v string) *DataHTMLElement {
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
func (e *DataHTMLElement) DIR(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DataHTMLElement) IfDIR(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DataHTMLElement) RemoveDIR(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *DataHTMLElement) DRAGGABLE(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DataHTMLElement) IfDRAGGABLE(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DataHTMLElement) RemoveDRAGGABLE(v string) *DataHTMLElement {
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
func (e *DataHTMLElement) ENTERKEYHINT(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DataHTMLElement) IfENTERKEYHINT(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DataHTMLElement) RemoveENTERKEYHINT(v string) *DataHTMLElement {
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
func (e *DataHTMLElement) HIDDEN(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DataHTMLElement) IfHIDDEN(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DataHTMLElement) RemoveHIDDEN(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DataHTMLElement) ID(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DataHTMLElement) IfID(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DataHTMLElement) RemoveID(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DataHTMLElement) INERT() *DataHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DataHTMLElement) IfINERT(cond bool) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DataHTMLElement) RemoveINERT() *DataHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DataHTMLElement) SetINERT(b bool) *DataHTMLElement {
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
func (e *DataHTMLElement) INPUTMODE(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DataHTMLElement) IfINPUTMODE(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DataHTMLElement) RemoveINPUTMODE(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *DataHTMLElement) IS(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DataHTMLElement) IfIS(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DataHTMLElement) RemoveIS(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DataHTMLElement) ITEMID(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DataHTMLElement) IfITEMID(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DataHTMLElement) RemoveITEMID(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *DataHTMLElement) ITEMPROP(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DataHTMLElement) IfITEMPROP(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DataHTMLElement) RemoveITEMPROP(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DataHTMLElement) ITEMREF(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DataHTMLElement) IfITEMREF(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DataHTMLElement) RemoveITEMREF(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DataHTMLElement) ITEMSCOPE() *DataHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DataHTMLElement) IfITEMSCOPE(cond bool) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DataHTMLElement) RemoveITEMSCOPE() *DataHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DataHTMLElement) SetITEMSCOPE(b bool) *DataHTMLElement {
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
func (e *DataHTMLElement) ITEMTYPE(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DataHTMLElement) IfITEMTYPE(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DataHTMLElement) RemoveITEMTYPE(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DataHTMLElement) LANG(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DataHTMLElement) IfLANG(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DataHTMLElement) RemoveLANG(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DataHTMLElement) NONCE(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DataHTMLElement) IfNONCE(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DataHTMLElement) RemoveNONCE(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - auto
//   - manual
//   - manual
func (e *DataHTMLElement) POPOVER(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DataHTMLElement) IfPOPOVER(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DataHTMLElement) RemovePOPOVER(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DataHTMLElement) SLOT(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DataHTMLElement) IfSLOT(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DataHTMLElement) RemoveSLOT(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *DataHTMLElement) SPELLCHECK(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DataHTMLElement) IfSPELLCHECK(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DataHTMLElement) RemoveSPELLCHECK(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DataHTMLElement) STYLE(k, v string) *DataHTMLElement {
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

func (e *DataHTMLElement) IfSTYLE(cond bool, k string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DataHTMLElement) RemoveSTYLE(k string) *DataHTMLElement {
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
func (e *DataHTMLElement) TABINDEX(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DataHTMLElement) IfTABINDEX(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DataHTMLElement) RemoveTABINDEX(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DataHTMLElement) TITLE(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DataHTMLElement) IfTITLE(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DataHTMLElement) RemoveTITLE(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *DataHTMLElement) TRANSLATE(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DataHTMLElement) IfTRANSLATE(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DataHTMLElement) RemoveTRANSLATE(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//   - valid_floating_point_number
func (e *DataHTMLElement) VALUE(v string) *DataHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeVALUEKey] = v
	return e
}

func (e *DataHTMLElement) IfVALUE(cond bool, v string) *DataHTMLElement {
	if !cond {
		return e
	}
	return e.VALUE(v)
}

func (e *DataHTMLElement) RemoveVALUE(v string) *DataHTMLElement {
	delete(e.StringAttributes, attributeVALUEKey)
	return e
}
