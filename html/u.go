package html

import (
	"fmt"
)

type UHTMLElement struct {
	*Element
}

func U(children ...ElementBuilder) *UHTMLElement {
	return &UHTMLElement{
		Element: &Element{
			Tag:           elementTagU,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *UHTMLElement) Children(children ...ElementBuilder) *UHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *UHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *UHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *UHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *UHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *UHTMLElement) Text(text string) *UHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *UHTMLElement) TextF(format string, args ...any) *UHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *UHTMLElement) Escaped(text string) *UHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *UHTMLElement) EscapedF(format string, args ...any) *UHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *UHTMLElement) CustomData(key, value string) *UHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *UHTMLElement) CustomDataRemove(key string) *UHTMLElement {
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
func (e *UHTMLElement) ACCESSKEY(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *UHTMLElement) IfACCESSKEY(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *UHTMLElement) RemoveACCESSKEY(v string) *UHTMLElement {
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
func (e *UHTMLElement) AUTOCAPITALIZE(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *UHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *UHTMLElement) RemoveAUTOCAPITALIZE(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *UHTMLElement) AUTOFOCUS() *UHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *UHTMLElement) IfAUTOFOCUS(cond bool) *UHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *UHTMLElement) RemoveAUTOFOCUS() *UHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *UHTMLElement) SetAUTOFOCUS(b bool) *UHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *UHTMLElement) CLASS(v string) *UHTMLElement {
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

func (e *UHTMLElement) IfCLASS(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *UHTMLElement) SetCLASS(v string) *UHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *UHTMLElement) RemoveCLASS(v string) *UHTMLElement {
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
func (e *UHTMLElement) CONTENTEDITABLE(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *UHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *UHTMLElement) RemoveCONTENTEDITABLE(v string) *UHTMLElement {
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
func (e *UHTMLElement) DIR(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *UHTMLElement) IfDIR(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *UHTMLElement) RemoveDIR(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *UHTMLElement) DRAGGABLE(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *UHTMLElement) IfDRAGGABLE(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *UHTMLElement) RemoveDRAGGABLE(v string) *UHTMLElement {
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
func (e *UHTMLElement) ENTERKEYHINT(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *UHTMLElement) IfENTERKEYHINT(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *UHTMLElement) RemoveENTERKEYHINT(v string) *UHTMLElement {
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
func (e *UHTMLElement) HIDDEN(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *UHTMLElement) IfHIDDEN(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *UHTMLElement) RemoveHIDDEN(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *UHTMLElement) ID(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *UHTMLElement) IfID(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *UHTMLElement) RemoveID(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *UHTMLElement) INERT() *UHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *UHTMLElement) IfINERT(cond bool) *UHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *UHTMLElement) RemoveINERT() *UHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *UHTMLElement) SetINERT(b bool) *UHTMLElement {
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
func (e *UHTMLElement) INPUTMODE(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *UHTMLElement) IfINPUTMODE(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *UHTMLElement) RemoveINPUTMODE(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *UHTMLElement) IS(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *UHTMLElement) IfIS(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *UHTMLElement) RemoveIS(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *UHTMLElement) ITEMID(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *UHTMLElement) IfITEMID(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *UHTMLElement) RemoveITEMID(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *UHTMLElement) ITEMPROP(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *UHTMLElement) IfITEMPROP(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *UHTMLElement) RemoveITEMPROP(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *UHTMLElement) ITEMREF(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *UHTMLElement) IfITEMREF(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *UHTMLElement) RemoveITEMREF(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *UHTMLElement) ITEMSCOPE() *UHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *UHTMLElement) IfITEMSCOPE(cond bool) *UHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *UHTMLElement) RemoveITEMSCOPE() *UHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *UHTMLElement) SetITEMSCOPE(b bool) *UHTMLElement {
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
func (e *UHTMLElement) ITEMTYPE(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *UHTMLElement) IfITEMTYPE(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *UHTMLElement) RemoveITEMTYPE(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *UHTMLElement) LANG(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *UHTMLElement) IfLANG(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *UHTMLElement) RemoveLANG(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *UHTMLElement) NONCE(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *UHTMLElement) IfNONCE(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *UHTMLElement) RemoveNONCE(v string) *UHTMLElement {
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
func (e *UHTMLElement) POPOVER(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *UHTMLElement) IfPOPOVER(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *UHTMLElement) RemovePOPOVER(v string) *UHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *UHTMLElement) SLOT(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *UHTMLElement) IfSLOT(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *UHTMLElement) RemoveSLOT(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *UHTMLElement) SPELLCHECK(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *UHTMLElement) IfSPELLCHECK(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *UHTMLElement) RemoveSPELLCHECK(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *UHTMLElement) STYLE(k, v string) *UHTMLElement {
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

func (e *UHTMLElement) IfSTYLE(cond bool, k string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *UHTMLElement) RemoveSTYLE(k string) *UHTMLElement {
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
func (e *UHTMLElement) TABINDEX(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *UHTMLElement) IfTABINDEX(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *UHTMLElement) RemoveTABINDEX(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *UHTMLElement) TITLE(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *UHTMLElement) IfTITLE(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *UHTMLElement) RemoveTITLE(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *UHTMLElement) TRANSLATE(v string) *UHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *UHTMLElement) IfTRANSLATE(cond bool, v string) *UHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *UHTMLElement) RemoveTRANSLATE(v string) *UHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
