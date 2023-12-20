package html

import (
	"fmt"
)

type ElementsHTMLElement struct {
	*Element
}

func ELEMENTS(children ...ElementBuilder) *ElementsHTMLElement {
	return &ElementsHTMLElement{
		Element: &Element{
			Tag:           elementTagELEMENTS,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *ElementsHTMLElement) Children(children ...ElementBuilder) *ElementsHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *ElementsHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ElementsHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *ElementsHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ElementsHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *ElementsHTMLElement) Text(text string) *ElementsHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *ElementsHTMLElement) TextF(format string, args ...any) *ElementsHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *ElementsHTMLElement) Escaped(text string) *ElementsHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *ElementsHTMLElement) EscapedF(format string, args ...any) *ElementsHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *ElementsHTMLElement) CustomData(key, value string) *ElementsHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ElementsHTMLElement) CustomDataRemove(key string) *ElementsHTMLElement {
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
func (e *ElementsHTMLElement) ACCESSKEY(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *ElementsHTMLElement) IfACCESSKEY(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *ElementsHTMLElement) RemoveACCESSKEY(v string) *ElementsHTMLElement {
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
func (e *ElementsHTMLElement) AUTOCAPITALIZE(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *ElementsHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *ElementsHTMLElement) RemoveAUTOCAPITALIZE(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *ElementsHTMLElement) AUTOFOCUS() *ElementsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *ElementsHTMLElement) IfAUTOFOCUS(cond bool) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *ElementsHTMLElement) RemoveAUTOFOCUS() *ElementsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *ElementsHTMLElement) SetAUTOFOCUS(b bool) *ElementsHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *ElementsHTMLElement) CLASS(v string) *ElementsHTMLElement {
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

func (e *ElementsHTMLElement) IfCLASS(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *ElementsHTMLElement) SetCLASS(v string) *ElementsHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *ElementsHTMLElement) RemoveCLASS(v string) *ElementsHTMLElement {
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
func (e *ElementsHTMLElement) CONTENTEDITABLE(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *ElementsHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *ElementsHTMLElement) RemoveCONTENTEDITABLE(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *ElementsHTMLElement) DIR(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *ElementsHTMLElement) IfDIR(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *ElementsHTMLElement) RemoveDIR(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *ElementsHTMLElement) DRAGGABLE(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *ElementsHTMLElement) IfDRAGGABLE(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *ElementsHTMLElement) RemoveDRAGGABLE(v string) *ElementsHTMLElement {
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
func (e *ElementsHTMLElement) ENTERKEYHINT(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *ElementsHTMLElement) IfENTERKEYHINT(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *ElementsHTMLElement) RemoveENTERKEYHINT(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *ElementsHTMLElement) HIDDEN(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *ElementsHTMLElement) IfHIDDEN(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *ElementsHTMLElement) RemoveHIDDEN(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *ElementsHTMLElement) ID(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *ElementsHTMLElement) IfID(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *ElementsHTMLElement) RemoveID(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *ElementsHTMLElement) INERT() *ElementsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *ElementsHTMLElement) IfINERT(cond bool) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *ElementsHTMLElement) RemoveINERT() *ElementsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *ElementsHTMLElement) SetINERT(b bool) *ElementsHTMLElement {
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
func (e *ElementsHTMLElement) INPUTMODE(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *ElementsHTMLElement) IfINPUTMODE(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *ElementsHTMLElement) RemoveINPUTMODE(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *ElementsHTMLElement) IS(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *ElementsHTMLElement) IfIS(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *ElementsHTMLElement) RemoveIS(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *ElementsHTMLElement) ITEMID(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *ElementsHTMLElement) IfITEMID(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *ElementsHTMLElement) RemoveITEMID(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *ElementsHTMLElement) ITEMPROP(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *ElementsHTMLElement) IfITEMPROP(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *ElementsHTMLElement) RemoveITEMPROP(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ElementsHTMLElement) ITEMREF(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *ElementsHTMLElement) IfITEMREF(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *ElementsHTMLElement) RemoveITEMREF(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *ElementsHTMLElement) ITEMSCOPE() *ElementsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *ElementsHTMLElement) IfITEMSCOPE(cond bool) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *ElementsHTMLElement) RemoveITEMSCOPE() *ElementsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *ElementsHTMLElement) SetITEMSCOPE(b bool) *ElementsHTMLElement {
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
func (e *ElementsHTMLElement) ITEMTYPE(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *ElementsHTMLElement) IfITEMTYPE(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *ElementsHTMLElement) RemoveITEMTYPE(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ElementsHTMLElement) LANG(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *ElementsHTMLElement) IfLANG(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *ElementsHTMLElement) RemoveLANG(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *ElementsHTMLElement) NONCE(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *ElementsHTMLElement) IfNONCE(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *ElementsHTMLElement) RemoveNONCE(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *ElementsHTMLElement) POPOVER(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *ElementsHTMLElement) IfPOPOVER(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *ElementsHTMLElement) RemovePOPOVER(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *ElementsHTMLElement) SLOT(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *ElementsHTMLElement) IfSLOT(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *ElementsHTMLElement) RemoveSLOT(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *ElementsHTMLElement) SPELLCHECK(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *ElementsHTMLElement) IfSPELLCHECK(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *ElementsHTMLElement) RemoveSPELLCHECK(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ElementsHTMLElement) STYLE(k, v string) *ElementsHTMLElement {
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

func (e *ElementsHTMLElement) IfSTYLE(cond bool, k string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *ElementsHTMLElement) RemoveSTYLE(k string) *ElementsHTMLElement {
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
func (e *ElementsHTMLElement) TABINDEX(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *ElementsHTMLElement) IfTABINDEX(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *ElementsHTMLElement) RemoveTABINDEX(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *ElementsHTMLElement) TITLE(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *ElementsHTMLElement) IfTITLE(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *ElementsHTMLElement) RemoveTITLE(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *ElementsHTMLElement) TRANSLATE(v string) *ElementsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *ElementsHTMLElement) IfTRANSLATE(cond bool, v string) *ElementsHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *ElementsHTMLElement) RemoveTRANSLATE(v string) *ElementsHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
