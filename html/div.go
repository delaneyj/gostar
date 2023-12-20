package html

import (
	"fmt"
)

type DivHTMLElement struct {
	*Element
}

func DIV(children ...ElementBuilder) *DivHTMLElement {
	return &DivHTMLElement{
		Element: &Element{
			Tag:           elementTagDIV,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DivHTMLElement) Children(children ...ElementBuilder) *DivHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DivHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DivHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DivHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DivHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DivHTMLElement) Text(text string) *DivHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DivHTMLElement) TextF(format string, args ...any) *DivHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DivHTMLElement) Escaped(text string) *DivHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DivHTMLElement) EscapedF(format string, args ...any) *DivHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DivHTMLElement) CustomData(key, value string) *DivHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DivHTMLElement) CustomDataRemove(key string) *DivHTMLElement {
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
func (e *DivHTMLElement) ACCESSKEY(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DivHTMLElement) IfACCESSKEY(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DivHTMLElement) RemoveACCESSKEY(v string) *DivHTMLElement {
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
func (e *DivHTMLElement) AUTOCAPITALIZE(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DivHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DivHTMLElement) RemoveAUTOCAPITALIZE(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DivHTMLElement) AUTOFOCUS() *DivHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DivHTMLElement) IfAUTOFOCUS(cond bool) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DivHTMLElement) RemoveAUTOFOCUS() *DivHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DivHTMLElement) SetAUTOFOCUS(b bool) *DivHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DivHTMLElement) CLASS(v string) *DivHTMLElement {
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

func (e *DivHTMLElement) IfCLASS(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DivHTMLElement) SetCLASS(v string) *DivHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DivHTMLElement) RemoveCLASS(v string) *DivHTMLElement {
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
func (e *DivHTMLElement) CONTENTEDITABLE(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DivHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DivHTMLElement) RemoveCONTENTEDITABLE(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *DivHTMLElement) DIR(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DivHTMLElement) IfDIR(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DivHTMLElement) RemoveDIR(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *DivHTMLElement) DRAGGABLE(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DivHTMLElement) IfDRAGGABLE(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DivHTMLElement) RemoveDRAGGABLE(v string) *DivHTMLElement {
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
func (e *DivHTMLElement) ENTERKEYHINT(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DivHTMLElement) IfENTERKEYHINT(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DivHTMLElement) RemoveENTERKEYHINT(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *DivHTMLElement) HIDDEN(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DivHTMLElement) IfHIDDEN(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DivHTMLElement) RemoveHIDDEN(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DivHTMLElement) ID(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DivHTMLElement) IfID(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DivHTMLElement) RemoveID(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DivHTMLElement) INERT() *DivHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DivHTMLElement) IfINERT(cond bool) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DivHTMLElement) RemoveINERT() *DivHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DivHTMLElement) SetINERT(b bool) *DivHTMLElement {
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
func (e *DivHTMLElement) INPUTMODE(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DivHTMLElement) IfINPUTMODE(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DivHTMLElement) RemoveINPUTMODE(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *DivHTMLElement) IS(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DivHTMLElement) IfIS(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DivHTMLElement) RemoveIS(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DivHTMLElement) ITEMID(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DivHTMLElement) IfITEMID(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DivHTMLElement) RemoveITEMID(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *DivHTMLElement) ITEMPROP(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DivHTMLElement) IfITEMPROP(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DivHTMLElement) RemoveITEMPROP(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DivHTMLElement) ITEMREF(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DivHTMLElement) IfITEMREF(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DivHTMLElement) RemoveITEMREF(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DivHTMLElement) ITEMSCOPE() *DivHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DivHTMLElement) IfITEMSCOPE(cond bool) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DivHTMLElement) RemoveITEMSCOPE() *DivHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DivHTMLElement) SetITEMSCOPE(b bool) *DivHTMLElement {
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
func (e *DivHTMLElement) ITEMTYPE(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DivHTMLElement) IfITEMTYPE(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DivHTMLElement) RemoveITEMTYPE(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DivHTMLElement) LANG(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DivHTMLElement) IfLANG(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DivHTMLElement) RemoveLANG(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DivHTMLElement) NONCE(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DivHTMLElement) IfNONCE(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DivHTMLElement) RemoveNONCE(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *DivHTMLElement) POPOVER(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DivHTMLElement) IfPOPOVER(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DivHTMLElement) RemovePOPOVER(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DivHTMLElement) SLOT(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DivHTMLElement) IfSLOT(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DivHTMLElement) RemoveSLOT(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *DivHTMLElement) SPELLCHECK(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DivHTMLElement) IfSPELLCHECK(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DivHTMLElement) RemoveSPELLCHECK(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DivHTMLElement) STYLE(k, v string) *DivHTMLElement {
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

func (e *DivHTMLElement) IfSTYLE(cond bool, k string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DivHTMLElement) RemoveSTYLE(k string) *DivHTMLElement {
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
func (e *DivHTMLElement) TABINDEX(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DivHTMLElement) IfTABINDEX(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DivHTMLElement) RemoveTABINDEX(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DivHTMLElement) TITLE(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DivHTMLElement) IfTITLE(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DivHTMLElement) RemoveTITLE(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *DivHTMLElement) TRANSLATE(v string) *DivHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DivHTMLElement) IfTRANSLATE(cond bool, v string) *DivHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DivHTMLElement) RemoveTRANSLATE(v string) *DivHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
