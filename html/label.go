package html

import (
	"fmt"
)

type LabelHTMLElement struct {
	*Element
}

func LABEL(children ...ElementBuilder) *LabelHTMLElement {
	return &LabelHTMLElement{
		Element: &Element{
			Tag:           elementTagLABEL,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *LabelHTMLElement) Children(children ...ElementBuilder) *LabelHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *LabelHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *LabelHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *LabelHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *LabelHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *LabelHTMLElement) Text(text string) *LabelHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *LabelHTMLElement) TextF(format string, args ...any) *LabelHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *LabelHTMLElement) Escaped(text string) *LabelHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *LabelHTMLElement) EscapedF(format string, args ...any) *LabelHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *LabelHTMLElement) CustomData(key, value string) *LabelHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *LabelHTMLElement) CustomDataRemove(key string) *LabelHTMLElement {
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
func (e *LabelHTMLElement) ACCESSKEY(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *LabelHTMLElement) IfACCESSKEY(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *LabelHTMLElement) RemoveACCESSKEY(v string) *LabelHTMLElement {
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
func (e *LabelHTMLElement) AUTOCAPITALIZE(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *LabelHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *LabelHTMLElement) RemoveAUTOCAPITALIZE(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *LabelHTMLElement) AUTOFOCUS() *LabelHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *LabelHTMLElement) IfAUTOFOCUS(cond bool) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *LabelHTMLElement) RemoveAUTOFOCUS() *LabelHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *LabelHTMLElement) SetAUTOFOCUS(b bool) *LabelHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *LabelHTMLElement) CLASS(v string) *LabelHTMLElement {
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

func (e *LabelHTMLElement) IfCLASS(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *LabelHTMLElement) SetCLASS(v string) *LabelHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *LabelHTMLElement) RemoveCLASS(v string) *LabelHTMLElement {
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
func (e *LabelHTMLElement) CONTENTEDITABLE(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *LabelHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *LabelHTMLElement) RemoveCONTENTEDITABLE(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *LabelHTMLElement) DIR(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *LabelHTMLElement) IfDIR(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *LabelHTMLElement) RemoveDIR(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *LabelHTMLElement) DRAGGABLE(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *LabelHTMLElement) IfDRAGGABLE(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *LabelHTMLElement) RemoveDRAGGABLE(v string) *LabelHTMLElement {
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
func (e *LabelHTMLElement) ENTERKEYHINT(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *LabelHTMLElement) IfENTERKEYHINT(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *LabelHTMLElement) RemoveENTERKEYHINT(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FOR sets the "for" attribute.
// Specifies controls from which the output was calculated
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *LabelHTMLElement) FOR(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORKey] = v
	return e
}

func (e *LabelHTMLElement) IfFOR(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.FOR(v)
}

func (e *LabelHTMLElement) RemoveFOR(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeFORKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *LabelHTMLElement) HIDDEN(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *LabelHTMLElement) IfHIDDEN(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *LabelHTMLElement) RemoveHIDDEN(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *LabelHTMLElement) ID(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *LabelHTMLElement) IfID(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *LabelHTMLElement) RemoveID(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *LabelHTMLElement) INERT() *LabelHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *LabelHTMLElement) IfINERT(cond bool) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *LabelHTMLElement) RemoveINERT() *LabelHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *LabelHTMLElement) SetINERT(b bool) *LabelHTMLElement {
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
func (e *LabelHTMLElement) INPUTMODE(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *LabelHTMLElement) IfINPUTMODE(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *LabelHTMLElement) RemoveINPUTMODE(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *LabelHTMLElement) IS(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *LabelHTMLElement) IfIS(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *LabelHTMLElement) RemoveIS(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *LabelHTMLElement) ITEMID(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *LabelHTMLElement) IfITEMID(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *LabelHTMLElement) RemoveITEMID(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *LabelHTMLElement) ITEMPROP(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *LabelHTMLElement) IfITEMPROP(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *LabelHTMLElement) RemoveITEMPROP(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *LabelHTMLElement) ITEMREF(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *LabelHTMLElement) IfITEMREF(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *LabelHTMLElement) RemoveITEMREF(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *LabelHTMLElement) ITEMSCOPE() *LabelHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *LabelHTMLElement) IfITEMSCOPE(cond bool) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *LabelHTMLElement) RemoveITEMSCOPE() *LabelHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *LabelHTMLElement) SetITEMSCOPE(b bool) *LabelHTMLElement {
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
func (e *LabelHTMLElement) ITEMTYPE(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *LabelHTMLElement) IfITEMTYPE(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *LabelHTMLElement) RemoveITEMTYPE(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *LabelHTMLElement) LANG(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *LabelHTMLElement) IfLANG(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *LabelHTMLElement) RemoveLANG(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *LabelHTMLElement) NONCE(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *LabelHTMLElement) IfNONCE(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *LabelHTMLElement) RemoveNONCE(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *LabelHTMLElement) POPOVER(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *LabelHTMLElement) IfPOPOVER(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *LabelHTMLElement) RemovePOPOVER(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *LabelHTMLElement) SLOT(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *LabelHTMLElement) IfSLOT(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *LabelHTMLElement) RemoveSLOT(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *LabelHTMLElement) SPELLCHECK(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *LabelHTMLElement) IfSPELLCHECK(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *LabelHTMLElement) RemoveSPELLCHECK(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *LabelHTMLElement) STYLE(k, v string) *LabelHTMLElement {
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

func (e *LabelHTMLElement) IfSTYLE(cond bool, k string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *LabelHTMLElement) RemoveSTYLE(k string) *LabelHTMLElement {
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
func (e *LabelHTMLElement) TABINDEX(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *LabelHTMLElement) IfTABINDEX(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *LabelHTMLElement) RemoveTABINDEX(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *LabelHTMLElement) TITLE(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *LabelHTMLElement) IfTITLE(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *LabelHTMLElement) RemoveTITLE(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *LabelHTMLElement) TRANSLATE(v string) *LabelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *LabelHTMLElement) IfTRANSLATE(cond bool, v string) *LabelHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *LabelHTMLElement) RemoveTRANSLATE(v string) *LabelHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
