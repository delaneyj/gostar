package html

import (
	"fmt"
)

type H2HTMLElement struct {
	*Element
}

func H2(children ...ElementBuilder) *H2HTMLElement {
	return &H2HTMLElement{
		Element: &Element{
			Tag:           elementTagH2,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *H2HTMLElement) Children(children ...ElementBuilder) *H2HTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *H2HTMLElement) IfChildren(condition bool, children ...ElementBuilder) *H2HTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *H2HTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *H2HTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *H2HTMLElement) Text(text string) *H2HTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *H2HTMLElement) TextF(format string, args ...any) *H2HTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *H2HTMLElement) Escaped(text string) *H2HTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *H2HTMLElement) EscapedF(format string, args ...any) *H2HTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *H2HTMLElement) CustomData(key, value string) *H2HTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *H2HTMLElement) CustomDataRemove(key string) *H2HTMLElement {
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
func (e *H2HTMLElement) ACCESSKEY(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *H2HTMLElement) IfACCESSKEY(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *H2HTMLElement) RemoveACCESSKEY(v string) *H2HTMLElement {
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
func (e *H2HTMLElement) AUTOCAPITALIZE(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *H2HTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *H2HTMLElement) RemoveAUTOCAPITALIZE(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *H2HTMLElement) AUTOFOCUS() *H2HTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *H2HTMLElement) IfAUTOFOCUS(cond bool) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *H2HTMLElement) RemoveAUTOFOCUS() *H2HTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *H2HTMLElement) SetAUTOFOCUS(b bool) *H2HTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *H2HTMLElement) CLASS(v string) *H2HTMLElement {
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

func (e *H2HTMLElement) IfCLASS(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *H2HTMLElement) SetCLASS(v string) *H2HTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *H2HTMLElement) RemoveCLASS(v string) *H2HTMLElement {
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
func (e *H2HTMLElement) CONTENTEDITABLE(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *H2HTMLElement) IfCONTENTEDITABLE(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *H2HTMLElement) RemoveCONTENTEDITABLE(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *H2HTMLElement) DIR(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *H2HTMLElement) IfDIR(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *H2HTMLElement) RemoveDIR(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *H2HTMLElement) DRAGGABLE(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *H2HTMLElement) IfDRAGGABLE(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *H2HTMLElement) RemoveDRAGGABLE(v string) *H2HTMLElement {
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
func (e *H2HTMLElement) ENTERKEYHINT(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *H2HTMLElement) IfENTERKEYHINT(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *H2HTMLElement) RemoveENTERKEYHINT(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *H2HTMLElement) HIDDEN(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *H2HTMLElement) IfHIDDEN(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *H2HTMLElement) RemoveHIDDEN(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *H2HTMLElement) ID(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *H2HTMLElement) IfID(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *H2HTMLElement) RemoveID(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *H2HTMLElement) INERT() *H2HTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *H2HTMLElement) IfINERT(cond bool) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *H2HTMLElement) RemoveINERT() *H2HTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *H2HTMLElement) SetINERT(b bool) *H2HTMLElement {
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
func (e *H2HTMLElement) INPUTMODE(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *H2HTMLElement) IfINPUTMODE(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *H2HTMLElement) RemoveINPUTMODE(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *H2HTMLElement) IS(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *H2HTMLElement) IfIS(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *H2HTMLElement) RemoveIS(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *H2HTMLElement) ITEMID(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *H2HTMLElement) IfITEMID(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *H2HTMLElement) RemoveITEMID(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *H2HTMLElement) ITEMPROP(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *H2HTMLElement) IfITEMPROP(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *H2HTMLElement) RemoveITEMPROP(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *H2HTMLElement) ITEMREF(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *H2HTMLElement) IfITEMREF(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *H2HTMLElement) RemoveITEMREF(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *H2HTMLElement) ITEMSCOPE() *H2HTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *H2HTMLElement) IfITEMSCOPE(cond bool) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *H2HTMLElement) RemoveITEMSCOPE() *H2HTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *H2HTMLElement) SetITEMSCOPE(b bool) *H2HTMLElement {
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
func (e *H2HTMLElement) ITEMTYPE(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *H2HTMLElement) IfITEMTYPE(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *H2HTMLElement) RemoveITEMTYPE(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *H2HTMLElement) LANG(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *H2HTMLElement) IfLANG(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *H2HTMLElement) RemoveLANG(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *H2HTMLElement) NONCE(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *H2HTMLElement) IfNONCE(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *H2HTMLElement) RemoveNONCE(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *H2HTMLElement) POPOVER(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *H2HTMLElement) IfPOPOVER(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *H2HTMLElement) RemovePOPOVER(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *H2HTMLElement) SLOT(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *H2HTMLElement) IfSLOT(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *H2HTMLElement) RemoveSLOT(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *H2HTMLElement) SPELLCHECK(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *H2HTMLElement) IfSPELLCHECK(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *H2HTMLElement) RemoveSPELLCHECK(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *H2HTMLElement) STYLE(k, v string) *H2HTMLElement {
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

func (e *H2HTMLElement) IfSTYLE(cond bool, k string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *H2HTMLElement) RemoveSTYLE(k string) *H2HTMLElement {
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
func (e *H2HTMLElement) TABINDEX(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *H2HTMLElement) IfTABINDEX(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *H2HTMLElement) RemoveTABINDEX(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *H2HTMLElement) TITLE(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *H2HTMLElement) IfTITLE(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *H2HTMLElement) RemoveTITLE(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *H2HTMLElement) TRANSLATE(v string) *H2HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *H2HTMLElement) IfTRANSLATE(cond bool, v string) *H2HTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *H2HTMLElement) RemoveTRANSLATE(v string) *H2HTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
