package html

import (
	"fmt"
)

type ProgressHTMLElement struct {
	*Element
}

func PROGRESS(children ...ElementBuilder) *ProgressHTMLElement {
	return &ProgressHTMLElement{
		Element: &Element{
			Tag:           elementTagPROGRESS,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *ProgressHTMLElement) Children(children ...ElementBuilder) *ProgressHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *ProgressHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ProgressHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *ProgressHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ProgressHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *ProgressHTMLElement) Text(text string) *ProgressHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *ProgressHTMLElement) TextF(format string, args ...any) *ProgressHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *ProgressHTMLElement) Escaped(text string) *ProgressHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *ProgressHTMLElement) EscapedF(format string, args ...any) *ProgressHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *ProgressHTMLElement) CustomData(key, value string) *ProgressHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ProgressHTMLElement) CustomDataRemove(key string) *ProgressHTMLElement {
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
func (e *ProgressHTMLElement) ACCESSKEY(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *ProgressHTMLElement) IfACCESSKEY(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *ProgressHTMLElement) RemoveACCESSKEY(v string) *ProgressHTMLElement {
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
func (e *ProgressHTMLElement) AUTOCAPITALIZE(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *ProgressHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *ProgressHTMLElement) RemoveAUTOCAPITALIZE(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *ProgressHTMLElement) AUTOFOCUS() *ProgressHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *ProgressHTMLElement) IfAUTOFOCUS(cond bool) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *ProgressHTMLElement) RemoveAUTOFOCUS() *ProgressHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *ProgressHTMLElement) SetAUTOFOCUS(b bool) *ProgressHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *ProgressHTMLElement) CLASS(v string) *ProgressHTMLElement {
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

func (e *ProgressHTMLElement) IfCLASS(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *ProgressHTMLElement) SetCLASS(v string) *ProgressHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *ProgressHTMLElement) RemoveCLASS(v string) *ProgressHTMLElement {
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
func (e *ProgressHTMLElement) CONTENTEDITABLE(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *ProgressHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *ProgressHTMLElement) RemoveCONTENTEDITABLE(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *ProgressHTMLElement) DIR(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *ProgressHTMLElement) IfDIR(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *ProgressHTMLElement) RemoveDIR(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *ProgressHTMLElement) DRAGGABLE(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *ProgressHTMLElement) IfDRAGGABLE(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *ProgressHTMLElement) RemoveDRAGGABLE(v string) *ProgressHTMLElement {
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
func (e *ProgressHTMLElement) ENTERKEYHINT(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *ProgressHTMLElement) IfENTERKEYHINT(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *ProgressHTMLElement) RemoveENTERKEYHINT(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *ProgressHTMLElement) HIDDEN(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *ProgressHTMLElement) IfHIDDEN(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *ProgressHTMLElement) RemoveHIDDEN(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *ProgressHTMLElement) ID(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *ProgressHTMLElement) IfID(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *ProgressHTMLElement) RemoveID(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *ProgressHTMLElement) INERT() *ProgressHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *ProgressHTMLElement) IfINERT(cond bool) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *ProgressHTMLElement) RemoveINERT() *ProgressHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *ProgressHTMLElement) SetINERT(b bool) *ProgressHTMLElement {
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
func (e *ProgressHTMLElement) INPUTMODE(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *ProgressHTMLElement) IfINPUTMODE(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *ProgressHTMLElement) RemoveINPUTMODE(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *ProgressHTMLElement) IS(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *ProgressHTMLElement) IfIS(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *ProgressHTMLElement) RemoveIS(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *ProgressHTMLElement) ITEMID(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *ProgressHTMLElement) IfITEMID(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *ProgressHTMLElement) RemoveITEMID(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *ProgressHTMLElement) ITEMPROP(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *ProgressHTMLElement) IfITEMPROP(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *ProgressHTMLElement) RemoveITEMPROP(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ProgressHTMLElement) ITEMREF(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *ProgressHTMLElement) IfITEMREF(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *ProgressHTMLElement) RemoveITEMREF(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *ProgressHTMLElement) ITEMSCOPE() *ProgressHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *ProgressHTMLElement) IfITEMSCOPE(cond bool) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *ProgressHTMLElement) RemoveITEMSCOPE() *ProgressHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *ProgressHTMLElement) SetITEMSCOPE(b bool) *ProgressHTMLElement {
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
func (e *ProgressHTMLElement) ITEMTYPE(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *ProgressHTMLElement) IfITEMTYPE(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *ProgressHTMLElement) RemoveITEMTYPE(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ProgressHTMLElement) LANG(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *ProgressHTMLElement) IfLANG(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *ProgressHTMLElement) RemoveLANG(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// MAX sets the "max" attribute.
// Upper bound of range
// Values values are constrained to:
//   - valid_floating_point_number
func (e *ProgressHTMLElement) MAX(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMAXKey] = v
	return e
}

func (e *ProgressHTMLElement) IfMAX(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.MAX(v)
}

func (e *ProgressHTMLElement) RemoveMAX(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeMAXKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *ProgressHTMLElement) NONCE(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *ProgressHTMLElement) IfNONCE(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *ProgressHTMLElement) RemoveNONCE(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *ProgressHTMLElement) POPOVER(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *ProgressHTMLElement) IfPOPOVER(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *ProgressHTMLElement) RemovePOPOVER(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *ProgressHTMLElement) SLOT(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *ProgressHTMLElement) IfSLOT(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *ProgressHTMLElement) RemoveSLOT(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *ProgressHTMLElement) SPELLCHECK(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *ProgressHTMLElement) IfSPELLCHECK(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *ProgressHTMLElement) RemoveSPELLCHECK(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ProgressHTMLElement) STYLE(k, v string) *ProgressHTMLElement {
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

func (e *ProgressHTMLElement) IfSTYLE(cond bool, k string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *ProgressHTMLElement) RemoveSTYLE(k string) *ProgressHTMLElement {
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
func (e *ProgressHTMLElement) TABINDEX(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *ProgressHTMLElement) IfTABINDEX(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *ProgressHTMLElement) RemoveTABINDEX(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *ProgressHTMLElement) TITLE(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *ProgressHTMLElement) IfTITLE(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *ProgressHTMLElement) RemoveTITLE(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *ProgressHTMLElement) TRANSLATE(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *ProgressHTMLElement) IfTRANSLATE(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *ProgressHTMLElement) RemoveTRANSLATE(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//   - valid_floating_point_number
func (e *ProgressHTMLElement) VALUE(v string) *ProgressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeVALUEKey] = v
	return e
}

func (e *ProgressHTMLElement) IfVALUE(cond bool, v string) *ProgressHTMLElement {
	if !cond {
		return e
	}
	return e.VALUE(v)
}

func (e *ProgressHTMLElement) RemoveVALUE(v string) *ProgressHTMLElement {
	delete(e.StringAttributes, attributeVALUEKey)
	return e
}
