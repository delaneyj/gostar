package html

import (
	"fmt"
)

type FooterHTMLElement struct {
	*Element
}

func FOOTER(children ...ElementBuilder) *FooterHTMLElement {
	return &FooterHTMLElement{
		Element: &Element{
			Tag:           elementTagFOOTER,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *FooterHTMLElement) Children(children ...ElementBuilder) *FooterHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *FooterHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *FooterHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *FooterHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *FooterHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *FooterHTMLElement) Text(text string) *FooterHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *FooterHTMLElement) TextF(format string, args ...any) *FooterHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *FooterHTMLElement) Escaped(text string) *FooterHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *FooterHTMLElement) EscapedF(format string, args ...any) *FooterHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *FooterHTMLElement) CustomData(key, value string) *FooterHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *FooterHTMLElement) CustomDataRemove(key string) *FooterHTMLElement {
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
func (e *FooterHTMLElement) ACCESSKEY(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *FooterHTMLElement) IfACCESSKEY(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *FooterHTMLElement) RemoveACCESSKEY(v string) *FooterHTMLElement {
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
func (e *FooterHTMLElement) AUTOCAPITALIZE(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *FooterHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *FooterHTMLElement) RemoveAUTOCAPITALIZE(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *FooterHTMLElement) AUTOFOCUS() *FooterHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *FooterHTMLElement) IfAUTOFOCUS(cond bool) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *FooterHTMLElement) RemoveAUTOFOCUS() *FooterHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *FooterHTMLElement) SetAUTOFOCUS(b bool) *FooterHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *FooterHTMLElement) CLASS(v string) *FooterHTMLElement {
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

func (e *FooterHTMLElement) IfCLASS(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *FooterHTMLElement) SetCLASS(v string) *FooterHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *FooterHTMLElement) RemoveCLASS(v string) *FooterHTMLElement {
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
func (e *FooterHTMLElement) CONTENTEDITABLE(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *FooterHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *FooterHTMLElement) RemoveCONTENTEDITABLE(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *FooterHTMLElement) DIR(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *FooterHTMLElement) IfDIR(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *FooterHTMLElement) RemoveDIR(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *FooterHTMLElement) DRAGGABLE(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *FooterHTMLElement) IfDRAGGABLE(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *FooterHTMLElement) RemoveDRAGGABLE(v string) *FooterHTMLElement {
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
func (e *FooterHTMLElement) ENTERKEYHINT(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *FooterHTMLElement) IfENTERKEYHINT(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *FooterHTMLElement) RemoveENTERKEYHINT(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *FooterHTMLElement) HIDDEN(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *FooterHTMLElement) IfHIDDEN(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *FooterHTMLElement) RemoveHIDDEN(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *FooterHTMLElement) ID(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *FooterHTMLElement) IfID(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *FooterHTMLElement) RemoveID(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *FooterHTMLElement) INERT() *FooterHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *FooterHTMLElement) IfINERT(cond bool) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *FooterHTMLElement) RemoveINERT() *FooterHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *FooterHTMLElement) SetINERT(b bool) *FooterHTMLElement {
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
func (e *FooterHTMLElement) INPUTMODE(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *FooterHTMLElement) IfINPUTMODE(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *FooterHTMLElement) RemoveINPUTMODE(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *FooterHTMLElement) IS(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *FooterHTMLElement) IfIS(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *FooterHTMLElement) RemoveIS(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *FooterHTMLElement) ITEMID(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *FooterHTMLElement) IfITEMID(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *FooterHTMLElement) RemoveITEMID(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *FooterHTMLElement) ITEMPROP(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *FooterHTMLElement) IfITEMPROP(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *FooterHTMLElement) RemoveITEMPROP(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *FooterHTMLElement) ITEMREF(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *FooterHTMLElement) IfITEMREF(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *FooterHTMLElement) RemoveITEMREF(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *FooterHTMLElement) ITEMSCOPE() *FooterHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *FooterHTMLElement) IfITEMSCOPE(cond bool) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *FooterHTMLElement) RemoveITEMSCOPE() *FooterHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *FooterHTMLElement) SetITEMSCOPE(b bool) *FooterHTMLElement {
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
func (e *FooterHTMLElement) ITEMTYPE(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *FooterHTMLElement) IfITEMTYPE(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *FooterHTMLElement) RemoveITEMTYPE(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *FooterHTMLElement) LANG(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *FooterHTMLElement) IfLANG(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *FooterHTMLElement) RemoveLANG(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *FooterHTMLElement) NONCE(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *FooterHTMLElement) IfNONCE(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *FooterHTMLElement) RemoveNONCE(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *FooterHTMLElement) POPOVER(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *FooterHTMLElement) IfPOPOVER(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *FooterHTMLElement) RemovePOPOVER(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *FooterHTMLElement) SLOT(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *FooterHTMLElement) IfSLOT(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *FooterHTMLElement) RemoveSLOT(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *FooterHTMLElement) SPELLCHECK(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *FooterHTMLElement) IfSPELLCHECK(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *FooterHTMLElement) RemoveSPELLCHECK(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *FooterHTMLElement) STYLE(k, v string) *FooterHTMLElement {
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

func (e *FooterHTMLElement) IfSTYLE(cond bool, k string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *FooterHTMLElement) RemoveSTYLE(k string) *FooterHTMLElement {
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
func (e *FooterHTMLElement) TABINDEX(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *FooterHTMLElement) IfTABINDEX(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *FooterHTMLElement) RemoveTABINDEX(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *FooterHTMLElement) TITLE(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *FooterHTMLElement) IfTITLE(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *FooterHTMLElement) RemoveTITLE(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *FooterHTMLElement) TRANSLATE(v string) *FooterHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *FooterHTMLElement) IfTRANSLATE(cond bool, v string) *FooterHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *FooterHTMLElement) RemoveTRANSLATE(v string) *FooterHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
