package html

import (
	"fmt"
)

type StrongHTMLElement struct {
	*Element
}

func STRONG(children ...ElementBuilder) *StrongHTMLElement {
	return &StrongHTMLElement{
		Element: &Element{
			Tag:           elementTagSTRONG,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *StrongHTMLElement) Children(children ...ElementBuilder) *StrongHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *StrongHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *StrongHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *StrongHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *StrongHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *StrongHTMLElement) Text(text string) *StrongHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *StrongHTMLElement) TextF(format string, args ...any) *StrongHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *StrongHTMLElement) Escaped(text string) *StrongHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *StrongHTMLElement) EscapedF(format string, args ...any) *StrongHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *StrongHTMLElement) CustomData(key, value string) *StrongHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *StrongHTMLElement) CustomDataRemove(key string) *StrongHTMLElement {
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
func (e *StrongHTMLElement) ACCESSKEY(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *StrongHTMLElement) IfACCESSKEY(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *StrongHTMLElement) RemoveACCESSKEY(v string) *StrongHTMLElement {
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
func (e *StrongHTMLElement) AUTOCAPITALIZE(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *StrongHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *StrongHTMLElement) RemoveAUTOCAPITALIZE(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *StrongHTMLElement) AUTOFOCUS() *StrongHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *StrongHTMLElement) IfAUTOFOCUS(cond bool) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *StrongHTMLElement) RemoveAUTOFOCUS() *StrongHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *StrongHTMLElement) SetAUTOFOCUS(b bool) *StrongHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *StrongHTMLElement) CLASS(v string) *StrongHTMLElement {
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

func (e *StrongHTMLElement) IfCLASS(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *StrongHTMLElement) SetCLASS(v string) *StrongHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *StrongHTMLElement) RemoveCLASS(v string) *StrongHTMLElement {
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
func (e *StrongHTMLElement) CONTENTEDITABLE(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *StrongHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *StrongHTMLElement) RemoveCONTENTEDITABLE(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *StrongHTMLElement) DIR(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *StrongHTMLElement) IfDIR(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *StrongHTMLElement) RemoveDIR(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *StrongHTMLElement) DRAGGABLE(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *StrongHTMLElement) IfDRAGGABLE(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *StrongHTMLElement) RemoveDRAGGABLE(v string) *StrongHTMLElement {
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
func (e *StrongHTMLElement) ENTERKEYHINT(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *StrongHTMLElement) IfENTERKEYHINT(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *StrongHTMLElement) RemoveENTERKEYHINT(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *StrongHTMLElement) HIDDEN(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *StrongHTMLElement) IfHIDDEN(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *StrongHTMLElement) RemoveHIDDEN(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *StrongHTMLElement) ID(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *StrongHTMLElement) IfID(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *StrongHTMLElement) RemoveID(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *StrongHTMLElement) INERT() *StrongHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *StrongHTMLElement) IfINERT(cond bool) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *StrongHTMLElement) RemoveINERT() *StrongHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *StrongHTMLElement) SetINERT(b bool) *StrongHTMLElement {
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
func (e *StrongHTMLElement) INPUTMODE(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *StrongHTMLElement) IfINPUTMODE(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *StrongHTMLElement) RemoveINPUTMODE(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *StrongHTMLElement) IS(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *StrongHTMLElement) IfIS(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *StrongHTMLElement) RemoveIS(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *StrongHTMLElement) ITEMID(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *StrongHTMLElement) IfITEMID(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *StrongHTMLElement) RemoveITEMID(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *StrongHTMLElement) ITEMPROP(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *StrongHTMLElement) IfITEMPROP(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *StrongHTMLElement) RemoveITEMPROP(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *StrongHTMLElement) ITEMREF(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *StrongHTMLElement) IfITEMREF(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *StrongHTMLElement) RemoveITEMREF(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *StrongHTMLElement) ITEMSCOPE() *StrongHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *StrongHTMLElement) IfITEMSCOPE(cond bool) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *StrongHTMLElement) RemoveITEMSCOPE() *StrongHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *StrongHTMLElement) SetITEMSCOPE(b bool) *StrongHTMLElement {
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
func (e *StrongHTMLElement) ITEMTYPE(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *StrongHTMLElement) IfITEMTYPE(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *StrongHTMLElement) RemoveITEMTYPE(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *StrongHTMLElement) LANG(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *StrongHTMLElement) IfLANG(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *StrongHTMLElement) RemoveLANG(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *StrongHTMLElement) NONCE(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *StrongHTMLElement) IfNONCE(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *StrongHTMLElement) RemoveNONCE(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *StrongHTMLElement) POPOVER(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *StrongHTMLElement) IfPOPOVER(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *StrongHTMLElement) RemovePOPOVER(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *StrongHTMLElement) SLOT(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *StrongHTMLElement) IfSLOT(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *StrongHTMLElement) RemoveSLOT(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *StrongHTMLElement) SPELLCHECK(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *StrongHTMLElement) IfSPELLCHECK(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *StrongHTMLElement) RemoveSPELLCHECK(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *StrongHTMLElement) STYLE(k, v string) *StrongHTMLElement {
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

func (e *StrongHTMLElement) IfSTYLE(cond bool, k string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *StrongHTMLElement) RemoveSTYLE(k string) *StrongHTMLElement {
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
func (e *StrongHTMLElement) TABINDEX(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *StrongHTMLElement) IfTABINDEX(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *StrongHTMLElement) RemoveTABINDEX(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *StrongHTMLElement) TITLE(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *StrongHTMLElement) IfTITLE(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *StrongHTMLElement) RemoveTITLE(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *StrongHTMLElement) TRANSLATE(v string) *StrongHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *StrongHTMLElement) IfTRANSLATE(cond bool, v string) *StrongHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *StrongHTMLElement) RemoveTRANSLATE(v string) *StrongHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
