package html

import (
	"fmt"
)

type AddressHTMLElement struct {
	*Element
}

func ADDRESS(children ...ElementBuilder) *AddressHTMLElement {
	return &AddressHTMLElement{
		Element: &Element{
			Tag:           elementTagADDRESS,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *AddressHTMLElement) Children(children ...ElementBuilder) *AddressHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *AddressHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *AddressHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *AddressHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *AddressHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *AddressHTMLElement) Text(text string) *AddressHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *AddressHTMLElement) TextF(format string, args ...any) *AddressHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *AddressHTMLElement) Escaped(text string) *AddressHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *AddressHTMLElement) EscapedF(format string, args ...any) *AddressHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *AddressHTMLElement) CustomData(key, value string) *AddressHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AddressHTMLElement) CustomDataRemove(key string) *AddressHTMLElement {
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
func (e *AddressHTMLElement) ACCESSKEY(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *AddressHTMLElement) IfACCESSKEY(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *AddressHTMLElement) RemoveACCESSKEY(v string) *AddressHTMLElement {
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
func (e *AddressHTMLElement) AUTOCAPITALIZE(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *AddressHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *AddressHTMLElement) RemoveAUTOCAPITALIZE(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *AddressHTMLElement) AUTOFOCUS() *AddressHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *AddressHTMLElement) IfAUTOFOCUS(cond bool) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *AddressHTMLElement) RemoveAUTOFOCUS() *AddressHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *AddressHTMLElement) SetAUTOFOCUS(b bool) *AddressHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *AddressHTMLElement) CLASS(v string) *AddressHTMLElement {
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

func (e *AddressHTMLElement) IfCLASS(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *AddressHTMLElement) SetCLASS(v string) *AddressHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *AddressHTMLElement) RemoveCLASS(v string) *AddressHTMLElement {
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
func (e *AddressHTMLElement) CONTENTEDITABLE(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *AddressHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *AddressHTMLElement) RemoveCONTENTEDITABLE(v string) *AddressHTMLElement {
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
func (e *AddressHTMLElement) DIR(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *AddressHTMLElement) IfDIR(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *AddressHTMLElement) RemoveDIR(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *AddressHTMLElement) DRAGGABLE(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *AddressHTMLElement) IfDRAGGABLE(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *AddressHTMLElement) RemoveDRAGGABLE(v string) *AddressHTMLElement {
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
func (e *AddressHTMLElement) ENTERKEYHINT(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *AddressHTMLElement) IfENTERKEYHINT(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *AddressHTMLElement) RemoveENTERKEYHINT(v string) *AddressHTMLElement {
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
func (e *AddressHTMLElement) HIDDEN(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *AddressHTMLElement) IfHIDDEN(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *AddressHTMLElement) RemoveHIDDEN(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *AddressHTMLElement) ID(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *AddressHTMLElement) IfID(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *AddressHTMLElement) RemoveID(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *AddressHTMLElement) INERT() *AddressHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *AddressHTMLElement) IfINERT(cond bool) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *AddressHTMLElement) RemoveINERT() *AddressHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *AddressHTMLElement) SetINERT(b bool) *AddressHTMLElement {
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
func (e *AddressHTMLElement) INPUTMODE(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *AddressHTMLElement) IfINPUTMODE(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *AddressHTMLElement) RemoveINPUTMODE(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *AddressHTMLElement) IS(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *AddressHTMLElement) IfIS(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *AddressHTMLElement) RemoveIS(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *AddressHTMLElement) ITEMID(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *AddressHTMLElement) IfITEMID(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *AddressHTMLElement) RemoveITEMID(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *AddressHTMLElement) ITEMPROP(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *AddressHTMLElement) IfITEMPROP(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *AddressHTMLElement) RemoveITEMPROP(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *AddressHTMLElement) ITEMREF(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *AddressHTMLElement) IfITEMREF(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *AddressHTMLElement) RemoveITEMREF(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *AddressHTMLElement) ITEMSCOPE() *AddressHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *AddressHTMLElement) IfITEMSCOPE(cond bool) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *AddressHTMLElement) RemoveITEMSCOPE() *AddressHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *AddressHTMLElement) SetITEMSCOPE(b bool) *AddressHTMLElement {
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
func (e *AddressHTMLElement) ITEMTYPE(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *AddressHTMLElement) IfITEMTYPE(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *AddressHTMLElement) RemoveITEMTYPE(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AddressHTMLElement) LANG(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *AddressHTMLElement) IfLANG(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *AddressHTMLElement) RemoveLANG(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *AddressHTMLElement) NONCE(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *AddressHTMLElement) IfNONCE(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *AddressHTMLElement) RemoveNONCE(v string) *AddressHTMLElement {
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
func (e *AddressHTMLElement) POPOVER(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *AddressHTMLElement) IfPOPOVER(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *AddressHTMLElement) RemovePOPOVER(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *AddressHTMLElement) SLOT(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *AddressHTMLElement) IfSLOT(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *AddressHTMLElement) RemoveSLOT(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *AddressHTMLElement) SPELLCHECK(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *AddressHTMLElement) IfSPELLCHECK(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *AddressHTMLElement) RemoveSPELLCHECK(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AddressHTMLElement) STYLE(k, v string) *AddressHTMLElement {
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

func (e *AddressHTMLElement) IfSTYLE(cond bool, k string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *AddressHTMLElement) RemoveSTYLE(k string) *AddressHTMLElement {
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
func (e *AddressHTMLElement) TABINDEX(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *AddressHTMLElement) IfTABINDEX(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *AddressHTMLElement) RemoveTABINDEX(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *AddressHTMLElement) TITLE(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *AddressHTMLElement) IfTITLE(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *AddressHTMLElement) RemoveTITLE(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *AddressHTMLElement) TRANSLATE(v string) *AddressHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *AddressHTMLElement) IfTRANSLATE(cond bool, v string) *AddressHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *AddressHTMLElement) RemoveTRANSLATE(v string) *AddressHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
