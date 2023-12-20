package html

import (
	"fmt"
)

type HrHTMLElement struct {
	*Element
}

func HR(children ...ElementBuilder) *HrHTMLElement {
	return &HrHTMLElement{
		Element: &Element{
			Tag:           elementTagHR,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *HrHTMLElement) Children(children ...ElementBuilder) *HrHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *HrHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *HrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *HrHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *HrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *HrHTMLElement) Text(text string) *HrHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *HrHTMLElement) TextF(format string, args ...any) *HrHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *HrHTMLElement) Escaped(text string) *HrHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *HrHTMLElement) EscapedF(format string, args ...any) *HrHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *HrHTMLElement) CustomData(key, value string) *HrHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *HrHTMLElement) CustomDataRemove(key string) *HrHTMLElement {
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
func (e *HrHTMLElement) ACCESSKEY(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *HrHTMLElement) IfACCESSKEY(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *HrHTMLElement) RemoveACCESSKEY(v string) *HrHTMLElement {
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
func (e *HrHTMLElement) AUTOCAPITALIZE(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *HrHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *HrHTMLElement) RemoveAUTOCAPITALIZE(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *HrHTMLElement) AUTOFOCUS() *HrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *HrHTMLElement) IfAUTOFOCUS(cond bool) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *HrHTMLElement) RemoveAUTOFOCUS() *HrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *HrHTMLElement) SetAUTOFOCUS(b bool) *HrHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *HrHTMLElement) CLASS(v string) *HrHTMLElement {
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

func (e *HrHTMLElement) IfCLASS(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *HrHTMLElement) SetCLASS(v string) *HrHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *HrHTMLElement) RemoveCLASS(v string) *HrHTMLElement {
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
func (e *HrHTMLElement) CONTENTEDITABLE(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *HrHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *HrHTMLElement) RemoveCONTENTEDITABLE(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *HrHTMLElement) DIR(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *HrHTMLElement) IfDIR(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *HrHTMLElement) RemoveDIR(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *HrHTMLElement) DRAGGABLE(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *HrHTMLElement) IfDRAGGABLE(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *HrHTMLElement) RemoveDRAGGABLE(v string) *HrHTMLElement {
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
func (e *HrHTMLElement) ENTERKEYHINT(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *HrHTMLElement) IfENTERKEYHINT(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *HrHTMLElement) RemoveENTERKEYHINT(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *HrHTMLElement) HIDDEN(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *HrHTMLElement) IfHIDDEN(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *HrHTMLElement) RemoveHIDDEN(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *HrHTMLElement) ID(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *HrHTMLElement) IfID(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *HrHTMLElement) RemoveID(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *HrHTMLElement) INERT() *HrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *HrHTMLElement) IfINERT(cond bool) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *HrHTMLElement) RemoveINERT() *HrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *HrHTMLElement) SetINERT(b bool) *HrHTMLElement {
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
func (e *HrHTMLElement) INPUTMODE(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *HrHTMLElement) IfINPUTMODE(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *HrHTMLElement) RemoveINPUTMODE(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *HrHTMLElement) IS(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *HrHTMLElement) IfIS(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *HrHTMLElement) RemoveIS(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *HrHTMLElement) ITEMID(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *HrHTMLElement) IfITEMID(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *HrHTMLElement) RemoveITEMID(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *HrHTMLElement) ITEMPROP(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *HrHTMLElement) IfITEMPROP(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *HrHTMLElement) RemoveITEMPROP(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *HrHTMLElement) ITEMREF(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *HrHTMLElement) IfITEMREF(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *HrHTMLElement) RemoveITEMREF(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *HrHTMLElement) ITEMSCOPE() *HrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *HrHTMLElement) IfITEMSCOPE(cond bool) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *HrHTMLElement) RemoveITEMSCOPE() *HrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *HrHTMLElement) SetITEMSCOPE(b bool) *HrHTMLElement {
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
func (e *HrHTMLElement) ITEMTYPE(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *HrHTMLElement) IfITEMTYPE(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *HrHTMLElement) RemoveITEMTYPE(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *HrHTMLElement) LANG(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *HrHTMLElement) IfLANG(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *HrHTMLElement) RemoveLANG(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *HrHTMLElement) NONCE(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *HrHTMLElement) IfNONCE(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *HrHTMLElement) RemoveNONCE(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *HrHTMLElement) POPOVER(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *HrHTMLElement) IfPOPOVER(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *HrHTMLElement) RemovePOPOVER(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *HrHTMLElement) SLOT(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *HrHTMLElement) IfSLOT(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *HrHTMLElement) RemoveSLOT(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *HrHTMLElement) SPELLCHECK(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *HrHTMLElement) IfSPELLCHECK(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *HrHTMLElement) RemoveSPELLCHECK(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *HrHTMLElement) STYLE(k, v string) *HrHTMLElement {
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

func (e *HrHTMLElement) IfSTYLE(cond bool, k string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *HrHTMLElement) RemoveSTYLE(k string) *HrHTMLElement {
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
func (e *HrHTMLElement) TABINDEX(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *HrHTMLElement) IfTABINDEX(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *HrHTMLElement) RemoveTABINDEX(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *HrHTMLElement) TITLE(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *HrHTMLElement) IfTITLE(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *HrHTMLElement) RemoveTITLE(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *HrHTMLElement) TRANSLATE(v string) *HrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *HrHTMLElement) IfTRANSLATE(cond bool, v string) *HrHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *HrHTMLElement) RemoveTRANSLATE(v string) *HrHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
