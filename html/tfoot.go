package html

import (
	"fmt"
)

type TfootHTMLElement struct {
	*Element
}

func TFOOT(children ...ElementBuilder) *TfootHTMLElement {
	return &TfootHTMLElement{
		Element: &Element{
			Tag:           elementTagTFOOT,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *TfootHTMLElement) Children(children ...ElementBuilder) *TfootHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *TfootHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TfootHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *TfootHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TfootHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *TfootHTMLElement) Text(text string) *TfootHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *TfootHTMLElement) TextF(format string, args ...any) *TfootHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *TfootHTMLElement) Escaped(text string) *TfootHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *TfootHTMLElement) EscapedF(format string, args ...any) *TfootHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TfootHTMLElement) CustomData(key, value string) *TfootHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TfootHTMLElement) CustomDataRemove(key string) *TfootHTMLElement {
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
func (e *TfootHTMLElement) ACCESSKEY(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *TfootHTMLElement) IfACCESSKEY(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *TfootHTMLElement) RemoveACCESSKEY(v string) *TfootHTMLElement {
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
func (e *TfootHTMLElement) AUTOCAPITALIZE(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *TfootHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *TfootHTMLElement) RemoveAUTOCAPITALIZE(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *TfootHTMLElement) AUTOFOCUS() *TfootHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *TfootHTMLElement) IfAUTOFOCUS(cond bool) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *TfootHTMLElement) RemoveAUTOFOCUS() *TfootHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *TfootHTMLElement) SetAUTOFOCUS(b bool) *TfootHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *TfootHTMLElement) CLASS(v string) *TfootHTMLElement {
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

func (e *TfootHTMLElement) IfCLASS(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *TfootHTMLElement) SetCLASS(v string) *TfootHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *TfootHTMLElement) RemoveCLASS(v string) *TfootHTMLElement {
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
func (e *TfootHTMLElement) CONTENTEDITABLE(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *TfootHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *TfootHTMLElement) RemoveCONTENTEDITABLE(v string) *TfootHTMLElement {
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
func (e *TfootHTMLElement) DIR(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *TfootHTMLElement) IfDIR(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *TfootHTMLElement) RemoveDIR(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *TfootHTMLElement) DRAGGABLE(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *TfootHTMLElement) IfDRAGGABLE(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *TfootHTMLElement) RemoveDRAGGABLE(v string) *TfootHTMLElement {
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
func (e *TfootHTMLElement) ENTERKEYHINT(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *TfootHTMLElement) IfENTERKEYHINT(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *TfootHTMLElement) RemoveENTERKEYHINT(v string) *TfootHTMLElement {
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
func (e *TfootHTMLElement) HIDDEN(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *TfootHTMLElement) IfHIDDEN(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *TfootHTMLElement) RemoveHIDDEN(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *TfootHTMLElement) ID(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *TfootHTMLElement) IfID(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *TfootHTMLElement) RemoveID(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *TfootHTMLElement) INERT() *TfootHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *TfootHTMLElement) IfINERT(cond bool) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *TfootHTMLElement) RemoveINERT() *TfootHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *TfootHTMLElement) SetINERT(b bool) *TfootHTMLElement {
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
func (e *TfootHTMLElement) INPUTMODE(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *TfootHTMLElement) IfINPUTMODE(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *TfootHTMLElement) RemoveINPUTMODE(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *TfootHTMLElement) IS(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *TfootHTMLElement) IfIS(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *TfootHTMLElement) RemoveIS(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *TfootHTMLElement) ITEMID(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *TfootHTMLElement) IfITEMID(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *TfootHTMLElement) RemoveITEMID(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *TfootHTMLElement) ITEMPROP(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *TfootHTMLElement) IfITEMPROP(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *TfootHTMLElement) RemoveITEMPROP(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TfootHTMLElement) ITEMREF(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *TfootHTMLElement) IfITEMREF(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *TfootHTMLElement) RemoveITEMREF(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *TfootHTMLElement) ITEMSCOPE() *TfootHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *TfootHTMLElement) IfITEMSCOPE(cond bool) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *TfootHTMLElement) RemoveITEMSCOPE() *TfootHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *TfootHTMLElement) SetITEMSCOPE(b bool) *TfootHTMLElement {
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
func (e *TfootHTMLElement) ITEMTYPE(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *TfootHTMLElement) IfITEMTYPE(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *TfootHTMLElement) RemoveITEMTYPE(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TfootHTMLElement) LANG(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *TfootHTMLElement) IfLANG(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *TfootHTMLElement) RemoveLANG(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *TfootHTMLElement) NONCE(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *TfootHTMLElement) IfNONCE(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *TfootHTMLElement) RemoveNONCE(v string) *TfootHTMLElement {
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
func (e *TfootHTMLElement) POPOVER(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *TfootHTMLElement) IfPOPOVER(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *TfootHTMLElement) RemovePOPOVER(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *TfootHTMLElement) SLOT(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *TfootHTMLElement) IfSLOT(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *TfootHTMLElement) RemoveSLOT(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *TfootHTMLElement) SPELLCHECK(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *TfootHTMLElement) IfSPELLCHECK(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *TfootHTMLElement) RemoveSPELLCHECK(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TfootHTMLElement) STYLE(k, v string) *TfootHTMLElement {
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

func (e *TfootHTMLElement) IfSTYLE(cond bool, k string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *TfootHTMLElement) RemoveSTYLE(k string) *TfootHTMLElement {
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
func (e *TfootHTMLElement) TABINDEX(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *TfootHTMLElement) IfTABINDEX(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *TfootHTMLElement) RemoveTABINDEX(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *TfootHTMLElement) TITLE(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *TfootHTMLElement) IfTITLE(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *TfootHTMLElement) RemoveTITLE(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *TfootHTMLElement) TRANSLATE(v string) *TfootHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *TfootHTMLElement) IfTRANSLATE(cond bool, v string) *TfootHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *TfootHTMLElement) RemoveTRANSLATE(v string) *TfootHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
