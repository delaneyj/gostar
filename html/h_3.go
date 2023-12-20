package html

import (
	"fmt"
)

type H3HTMLElement struct {
	*Element
}

func H3(children ...ElementBuilder) *H3HTMLElement {
	return &H3HTMLElement{
		Element: &Element{
			Tag:           elementTagH3,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *H3HTMLElement) Children(children ...ElementBuilder) *H3HTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *H3HTMLElement) IfChildren(condition bool, children ...ElementBuilder) *H3HTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *H3HTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *H3HTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *H3HTMLElement) Text(text string) *H3HTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *H3HTMLElement) TextF(format string, args ...any) *H3HTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *H3HTMLElement) Escaped(text string) *H3HTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *H3HTMLElement) EscapedF(format string, args ...any) *H3HTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *H3HTMLElement) CustomData(key, value string) *H3HTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *H3HTMLElement) CustomDataRemove(key string) *H3HTMLElement {
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
func (e *H3HTMLElement) ACCESSKEY(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *H3HTMLElement) IfACCESSKEY(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *H3HTMLElement) RemoveACCESSKEY(v string) *H3HTMLElement {
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
func (e *H3HTMLElement) AUTOCAPITALIZE(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *H3HTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *H3HTMLElement) RemoveAUTOCAPITALIZE(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *H3HTMLElement) AUTOFOCUS() *H3HTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *H3HTMLElement) IfAUTOFOCUS(cond bool) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *H3HTMLElement) RemoveAUTOFOCUS() *H3HTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *H3HTMLElement) SetAUTOFOCUS(b bool) *H3HTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *H3HTMLElement) CLASS(v string) *H3HTMLElement {
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

func (e *H3HTMLElement) IfCLASS(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *H3HTMLElement) SetCLASS(v string) *H3HTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *H3HTMLElement) RemoveCLASS(v string) *H3HTMLElement {
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
func (e *H3HTMLElement) CONTENTEDITABLE(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *H3HTMLElement) IfCONTENTEDITABLE(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *H3HTMLElement) RemoveCONTENTEDITABLE(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *H3HTMLElement) DIR(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *H3HTMLElement) IfDIR(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *H3HTMLElement) RemoveDIR(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *H3HTMLElement) DRAGGABLE(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *H3HTMLElement) IfDRAGGABLE(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *H3HTMLElement) RemoveDRAGGABLE(v string) *H3HTMLElement {
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
func (e *H3HTMLElement) ENTERKEYHINT(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *H3HTMLElement) IfENTERKEYHINT(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *H3HTMLElement) RemoveENTERKEYHINT(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *H3HTMLElement) HIDDEN(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *H3HTMLElement) IfHIDDEN(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *H3HTMLElement) RemoveHIDDEN(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *H3HTMLElement) ID(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *H3HTMLElement) IfID(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *H3HTMLElement) RemoveID(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *H3HTMLElement) INERT() *H3HTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *H3HTMLElement) IfINERT(cond bool) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *H3HTMLElement) RemoveINERT() *H3HTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *H3HTMLElement) SetINERT(b bool) *H3HTMLElement {
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
func (e *H3HTMLElement) INPUTMODE(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *H3HTMLElement) IfINPUTMODE(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *H3HTMLElement) RemoveINPUTMODE(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *H3HTMLElement) IS(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *H3HTMLElement) IfIS(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *H3HTMLElement) RemoveIS(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *H3HTMLElement) ITEMID(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *H3HTMLElement) IfITEMID(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *H3HTMLElement) RemoveITEMID(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *H3HTMLElement) ITEMPROP(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *H3HTMLElement) IfITEMPROP(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *H3HTMLElement) RemoveITEMPROP(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *H3HTMLElement) ITEMREF(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *H3HTMLElement) IfITEMREF(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *H3HTMLElement) RemoveITEMREF(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *H3HTMLElement) ITEMSCOPE() *H3HTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *H3HTMLElement) IfITEMSCOPE(cond bool) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *H3HTMLElement) RemoveITEMSCOPE() *H3HTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *H3HTMLElement) SetITEMSCOPE(b bool) *H3HTMLElement {
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
func (e *H3HTMLElement) ITEMTYPE(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *H3HTMLElement) IfITEMTYPE(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *H3HTMLElement) RemoveITEMTYPE(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *H3HTMLElement) LANG(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *H3HTMLElement) IfLANG(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *H3HTMLElement) RemoveLANG(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *H3HTMLElement) NONCE(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *H3HTMLElement) IfNONCE(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *H3HTMLElement) RemoveNONCE(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *H3HTMLElement) POPOVER(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *H3HTMLElement) IfPOPOVER(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *H3HTMLElement) RemovePOPOVER(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *H3HTMLElement) SLOT(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *H3HTMLElement) IfSLOT(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *H3HTMLElement) RemoveSLOT(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *H3HTMLElement) SPELLCHECK(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *H3HTMLElement) IfSPELLCHECK(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *H3HTMLElement) RemoveSPELLCHECK(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *H3HTMLElement) STYLE(k, v string) *H3HTMLElement {
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

func (e *H3HTMLElement) IfSTYLE(cond bool, k string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *H3HTMLElement) RemoveSTYLE(k string) *H3HTMLElement {
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
func (e *H3HTMLElement) TABINDEX(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *H3HTMLElement) IfTABINDEX(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *H3HTMLElement) RemoveTABINDEX(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *H3HTMLElement) TITLE(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *H3HTMLElement) IfTITLE(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *H3HTMLElement) RemoveTITLE(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *H3HTMLElement) TRANSLATE(v string) *H3HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *H3HTMLElement) IfTRANSLATE(cond bool, v string) *H3HTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *H3HTMLElement) RemoveTRANSLATE(v string) *H3HTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
