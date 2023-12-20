package html

import (
	"fmt"
)

type EmHTMLElement struct {
	*Element
}

func EM(children ...ElementBuilder) *EmHTMLElement {
	return &EmHTMLElement{
		Element: &Element{
			Tag:           elementTagEM,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *EmHTMLElement) Children(children ...ElementBuilder) *EmHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *EmHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *EmHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *EmHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *EmHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *EmHTMLElement) Text(text string) *EmHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *EmHTMLElement) TextF(format string, args ...any) *EmHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *EmHTMLElement) Escaped(text string) *EmHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *EmHTMLElement) EscapedF(format string, args ...any) *EmHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *EmHTMLElement) CustomData(key, value string) *EmHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *EmHTMLElement) CustomDataRemove(key string) *EmHTMLElement {
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
func (e *EmHTMLElement) ACCESSKEY(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *EmHTMLElement) IfACCESSKEY(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *EmHTMLElement) RemoveACCESSKEY(v string) *EmHTMLElement {
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
func (e *EmHTMLElement) AUTOCAPITALIZE(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *EmHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *EmHTMLElement) RemoveAUTOCAPITALIZE(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *EmHTMLElement) AUTOFOCUS() *EmHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *EmHTMLElement) IfAUTOFOCUS(cond bool) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *EmHTMLElement) RemoveAUTOFOCUS() *EmHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *EmHTMLElement) SetAUTOFOCUS(b bool) *EmHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *EmHTMLElement) CLASS(v string) *EmHTMLElement {
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

func (e *EmHTMLElement) IfCLASS(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *EmHTMLElement) SetCLASS(v string) *EmHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *EmHTMLElement) RemoveCLASS(v string) *EmHTMLElement {
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
func (e *EmHTMLElement) CONTENTEDITABLE(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *EmHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *EmHTMLElement) RemoveCONTENTEDITABLE(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *EmHTMLElement) DIR(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *EmHTMLElement) IfDIR(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *EmHTMLElement) RemoveDIR(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *EmHTMLElement) DRAGGABLE(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *EmHTMLElement) IfDRAGGABLE(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *EmHTMLElement) RemoveDRAGGABLE(v string) *EmHTMLElement {
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
func (e *EmHTMLElement) ENTERKEYHINT(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *EmHTMLElement) IfENTERKEYHINT(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *EmHTMLElement) RemoveENTERKEYHINT(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *EmHTMLElement) HIDDEN(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *EmHTMLElement) IfHIDDEN(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *EmHTMLElement) RemoveHIDDEN(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *EmHTMLElement) ID(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *EmHTMLElement) IfID(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *EmHTMLElement) RemoveID(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *EmHTMLElement) INERT() *EmHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *EmHTMLElement) IfINERT(cond bool) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *EmHTMLElement) RemoveINERT() *EmHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *EmHTMLElement) SetINERT(b bool) *EmHTMLElement {
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
func (e *EmHTMLElement) INPUTMODE(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *EmHTMLElement) IfINPUTMODE(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *EmHTMLElement) RemoveINPUTMODE(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *EmHTMLElement) IS(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *EmHTMLElement) IfIS(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *EmHTMLElement) RemoveIS(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *EmHTMLElement) ITEMID(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *EmHTMLElement) IfITEMID(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *EmHTMLElement) RemoveITEMID(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *EmHTMLElement) ITEMPROP(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *EmHTMLElement) IfITEMPROP(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *EmHTMLElement) RemoveITEMPROP(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *EmHTMLElement) ITEMREF(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *EmHTMLElement) IfITEMREF(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *EmHTMLElement) RemoveITEMREF(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *EmHTMLElement) ITEMSCOPE() *EmHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *EmHTMLElement) IfITEMSCOPE(cond bool) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *EmHTMLElement) RemoveITEMSCOPE() *EmHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *EmHTMLElement) SetITEMSCOPE(b bool) *EmHTMLElement {
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
func (e *EmHTMLElement) ITEMTYPE(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *EmHTMLElement) IfITEMTYPE(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *EmHTMLElement) RemoveITEMTYPE(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *EmHTMLElement) LANG(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *EmHTMLElement) IfLANG(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *EmHTMLElement) RemoveLANG(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *EmHTMLElement) NONCE(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *EmHTMLElement) IfNONCE(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *EmHTMLElement) RemoveNONCE(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *EmHTMLElement) POPOVER(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *EmHTMLElement) IfPOPOVER(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *EmHTMLElement) RemovePOPOVER(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *EmHTMLElement) SLOT(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *EmHTMLElement) IfSLOT(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *EmHTMLElement) RemoveSLOT(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *EmHTMLElement) SPELLCHECK(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *EmHTMLElement) IfSPELLCHECK(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *EmHTMLElement) RemoveSPELLCHECK(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *EmHTMLElement) STYLE(k, v string) *EmHTMLElement {
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

func (e *EmHTMLElement) IfSTYLE(cond bool, k string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *EmHTMLElement) RemoveSTYLE(k string) *EmHTMLElement {
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
func (e *EmHTMLElement) TABINDEX(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *EmHTMLElement) IfTABINDEX(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *EmHTMLElement) RemoveTABINDEX(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *EmHTMLElement) TITLE(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *EmHTMLElement) IfTITLE(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *EmHTMLElement) RemoveTITLE(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *EmHTMLElement) TRANSLATE(v string) *EmHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *EmHTMLElement) IfTRANSLATE(cond bool, v string) *EmHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *EmHTMLElement) RemoveTRANSLATE(v string) *EmHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
