package html

import (
	"fmt"
)

type NavHTMLElement struct {
	*Element
}

func NAV(children ...ElementBuilder) *NavHTMLElement {
	return &NavHTMLElement{
		Element: &Element{
			Tag:           elementTagNAV,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *NavHTMLElement) Children(children ...ElementBuilder) *NavHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *NavHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *NavHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *NavHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *NavHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *NavHTMLElement) Text(text string) *NavHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *NavHTMLElement) TextF(format string, args ...any) *NavHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *NavHTMLElement) Escaped(text string) *NavHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *NavHTMLElement) EscapedF(format string, args ...any) *NavHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *NavHTMLElement) CustomData(key, value string) *NavHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *NavHTMLElement) CustomDataRemove(key string) *NavHTMLElement {
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
func (e *NavHTMLElement) ACCESSKEY(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *NavHTMLElement) IfACCESSKEY(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *NavHTMLElement) RemoveACCESSKEY(v string) *NavHTMLElement {
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
func (e *NavHTMLElement) AUTOCAPITALIZE(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *NavHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *NavHTMLElement) RemoveAUTOCAPITALIZE(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *NavHTMLElement) AUTOFOCUS() *NavHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *NavHTMLElement) IfAUTOFOCUS(cond bool) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *NavHTMLElement) RemoveAUTOFOCUS() *NavHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *NavHTMLElement) SetAUTOFOCUS(b bool) *NavHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *NavHTMLElement) CLASS(v string) *NavHTMLElement {
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

func (e *NavHTMLElement) IfCLASS(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *NavHTMLElement) SetCLASS(v string) *NavHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *NavHTMLElement) RemoveCLASS(v string) *NavHTMLElement {
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
func (e *NavHTMLElement) CONTENTEDITABLE(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *NavHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *NavHTMLElement) RemoveCONTENTEDITABLE(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *NavHTMLElement) DIR(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *NavHTMLElement) IfDIR(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *NavHTMLElement) RemoveDIR(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *NavHTMLElement) DRAGGABLE(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *NavHTMLElement) IfDRAGGABLE(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *NavHTMLElement) RemoveDRAGGABLE(v string) *NavHTMLElement {
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
func (e *NavHTMLElement) ENTERKEYHINT(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *NavHTMLElement) IfENTERKEYHINT(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *NavHTMLElement) RemoveENTERKEYHINT(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *NavHTMLElement) HIDDEN(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *NavHTMLElement) IfHIDDEN(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *NavHTMLElement) RemoveHIDDEN(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *NavHTMLElement) ID(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *NavHTMLElement) IfID(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *NavHTMLElement) RemoveID(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *NavHTMLElement) INERT() *NavHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *NavHTMLElement) IfINERT(cond bool) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *NavHTMLElement) RemoveINERT() *NavHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *NavHTMLElement) SetINERT(b bool) *NavHTMLElement {
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
func (e *NavHTMLElement) INPUTMODE(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *NavHTMLElement) IfINPUTMODE(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *NavHTMLElement) RemoveINPUTMODE(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *NavHTMLElement) IS(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *NavHTMLElement) IfIS(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *NavHTMLElement) RemoveIS(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *NavHTMLElement) ITEMID(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *NavHTMLElement) IfITEMID(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *NavHTMLElement) RemoveITEMID(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *NavHTMLElement) ITEMPROP(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *NavHTMLElement) IfITEMPROP(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *NavHTMLElement) RemoveITEMPROP(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *NavHTMLElement) ITEMREF(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *NavHTMLElement) IfITEMREF(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *NavHTMLElement) RemoveITEMREF(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *NavHTMLElement) ITEMSCOPE() *NavHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *NavHTMLElement) IfITEMSCOPE(cond bool) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *NavHTMLElement) RemoveITEMSCOPE() *NavHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *NavHTMLElement) SetITEMSCOPE(b bool) *NavHTMLElement {
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
func (e *NavHTMLElement) ITEMTYPE(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *NavHTMLElement) IfITEMTYPE(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *NavHTMLElement) RemoveITEMTYPE(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *NavHTMLElement) LANG(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *NavHTMLElement) IfLANG(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *NavHTMLElement) RemoveLANG(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *NavHTMLElement) NONCE(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *NavHTMLElement) IfNONCE(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *NavHTMLElement) RemoveNONCE(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *NavHTMLElement) POPOVER(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *NavHTMLElement) IfPOPOVER(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *NavHTMLElement) RemovePOPOVER(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *NavHTMLElement) SLOT(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *NavHTMLElement) IfSLOT(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *NavHTMLElement) RemoveSLOT(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *NavHTMLElement) SPELLCHECK(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *NavHTMLElement) IfSPELLCHECK(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *NavHTMLElement) RemoveSPELLCHECK(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *NavHTMLElement) STYLE(k, v string) *NavHTMLElement {
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

func (e *NavHTMLElement) IfSTYLE(cond bool, k string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *NavHTMLElement) RemoveSTYLE(k string) *NavHTMLElement {
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
func (e *NavHTMLElement) TABINDEX(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *NavHTMLElement) IfTABINDEX(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *NavHTMLElement) RemoveTABINDEX(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *NavHTMLElement) TITLE(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *NavHTMLElement) IfTITLE(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *NavHTMLElement) RemoveTITLE(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *NavHTMLElement) TRANSLATE(v string) *NavHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *NavHTMLElement) IfTRANSLATE(cond bool, v string) *NavHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *NavHTMLElement) RemoveTRANSLATE(v string) *NavHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
