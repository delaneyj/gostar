package html

import (
	"fmt"
)

type BHTMLElement struct {
	*Element
}

func B(children ...ElementBuilder) *BHTMLElement {
	return &BHTMLElement{
		Element: &Element{
			Tag:           elementTagB,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *BHTMLElement) Children(children ...ElementBuilder) *BHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *BHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *BHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *BHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *BHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *BHTMLElement) Text(text string) *BHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *BHTMLElement) TextF(format string, args ...any) *BHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *BHTMLElement) Escaped(text string) *BHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *BHTMLElement) EscapedF(format string, args ...any) *BHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *BHTMLElement) CustomData(key, value string) *BHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *BHTMLElement) CustomDataRemove(key string) *BHTMLElement {
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
func (e *BHTMLElement) ACCESSKEY(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *BHTMLElement) IfACCESSKEY(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *BHTMLElement) RemoveACCESSKEY(v string) *BHTMLElement {
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
func (e *BHTMLElement) AUTOCAPITALIZE(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *BHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *BHTMLElement) RemoveAUTOCAPITALIZE(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *BHTMLElement) AUTOFOCUS() *BHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *BHTMLElement) IfAUTOFOCUS(cond bool) *BHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *BHTMLElement) RemoveAUTOFOCUS() *BHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *BHTMLElement) SetAUTOFOCUS(b bool) *BHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *BHTMLElement) CLASS(v string) *BHTMLElement {
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

func (e *BHTMLElement) IfCLASS(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *BHTMLElement) SetCLASS(v string) *BHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *BHTMLElement) RemoveCLASS(v string) *BHTMLElement {
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
func (e *BHTMLElement) CONTENTEDITABLE(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *BHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *BHTMLElement) RemoveCONTENTEDITABLE(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *BHTMLElement) DIR(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *BHTMLElement) IfDIR(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *BHTMLElement) RemoveDIR(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *BHTMLElement) DRAGGABLE(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *BHTMLElement) IfDRAGGABLE(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *BHTMLElement) RemoveDRAGGABLE(v string) *BHTMLElement {
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
func (e *BHTMLElement) ENTERKEYHINT(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *BHTMLElement) IfENTERKEYHINT(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *BHTMLElement) RemoveENTERKEYHINT(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *BHTMLElement) HIDDEN(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *BHTMLElement) IfHIDDEN(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *BHTMLElement) RemoveHIDDEN(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *BHTMLElement) ID(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *BHTMLElement) IfID(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *BHTMLElement) RemoveID(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *BHTMLElement) INERT() *BHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *BHTMLElement) IfINERT(cond bool) *BHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *BHTMLElement) RemoveINERT() *BHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *BHTMLElement) SetINERT(b bool) *BHTMLElement {
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
func (e *BHTMLElement) INPUTMODE(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *BHTMLElement) IfINPUTMODE(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *BHTMLElement) RemoveINPUTMODE(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *BHTMLElement) IS(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *BHTMLElement) IfIS(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *BHTMLElement) RemoveIS(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *BHTMLElement) ITEMID(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *BHTMLElement) IfITEMID(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *BHTMLElement) RemoveITEMID(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *BHTMLElement) ITEMPROP(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *BHTMLElement) IfITEMPROP(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *BHTMLElement) RemoveITEMPROP(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *BHTMLElement) ITEMREF(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *BHTMLElement) IfITEMREF(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *BHTMLElement) RemoveITEMREF(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *BHTMLElement) ITEMSCOPE() *BHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *BHTMLElement) IfITEMSCOPE(cond bool) *BHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *BHTMLElement) RemoveITEMSCOPE() *BHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *BHTMLElement) SetITEMSCOPE(b bool) *BHTMLElement {
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
func (e *BHTMLElement) ITEMTYPE(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *BHTMLElement) IfITEMTYPE(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *BHTMLElement) RemoveITEMTYPE(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *BHTMLElement) LANG(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *BHTMLElement) IfLANG(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *BHTMLElement) RemoveLANG(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *BHTMLElement) NONCE(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *BHTMLElement) IfNONCE(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *BHTMLElement) RemoveNONCE(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *BHTMLElement) POPOVER(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *BHTMLElement) IfPOPOVER(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *BHTMLElement) RemovePOPOVER(v string) *BHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *BHTMLElement) SLOT(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *BHTMLElement) IfSLOT(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *BHTMLElement) RemoveSLOT(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *BHTMLElement) SPELLCHECK(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *BHTMLElement) IfSPELLCHECK(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *BHTMLElement) RemoveSPELLCHECK(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *BHTMLElement) STYLE(k, v string) *BHTMLElement {
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

func (e *BHTMLElement) IfSTYLE(cond bool, k string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *BHTMLElement) RemoveSTYLE(k string) *BHTMLElement {
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
func (e *BHTMLElement) TABINDEX(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *BHTMLElement) IfTABINDEX(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *BHTMLElement) RemoveTABINDEX(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *BHTMLElement) TITLE(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *BHTMLElement) IfTITLE(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *BHTMLElement) RemoveTITLE(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *BHTMLElement) TRANSLATE(v string) *BHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *BHTMLElement) IfTRANSLATE(cond bool, v string) *BHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *BHTMLElement) RemoveTRANSLATE(v string) *BHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
