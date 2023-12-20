package html

import (
	"fmt"
)

type MathHTMLElement struct {
	*Element
}

func MATH(children ...ElementBuilder) *MathHTMLElement {
	return &MathHTMLElement{
		Element: &Element{
			Tag:           elementTagMATH,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *MathHTMLElement) Children(children ...ElementBuilder) *MathHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *MathHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *MathHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathHTMLElement) Text(text string) *MathHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathHTMLElement) TextF(format string, args ...any) *MathHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathHTMLElement) Escaped(text string) *MathHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathHTMLElement) EscapedF(format string, args ...any) *MathHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathHTMLElement) CustomData(key, value string) *MathHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *MathHTMLElement) CustomDataRemove(key string) *MathHTMLElement {
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
func (e *MathHTMLElement) ACCESSKEY(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *MathHTMLElement) IfACCESSKEY(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *MathHTMLElement) RemoveACCESSKEY(v string) *MathHTMLElement {
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
func (e *MathHTMLElement) AUTOCAPITALIZE(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *MathHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *MathHTMLElement) RemoveAUTOCAPITALIZE(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *MathHTMLElement) AUTOFOCUS() *MathHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *MathHTMLElement) IfAUTOFOCUS(cond bool) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *MathHTMLElement) RemoveAUTOFOCUS() *MathHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *MathHTMLElement) SetAUTOFOCUS(b bool) *MathHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *MathHTMLElement) CLASS(v string) *MathHTMLElement {
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

func (e *MathHTMLElement) IfCLASS(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *MathHTMLElement) SetCLASS(v string) *MathHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *MathHTMLElement) RemoveCLASS(v string) *MathHTMLElement {
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
func (e *MathHTMLElement) CONTENTEDITABLE(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *MathHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *MathHTMLElement) RemoveCONTENTEDITABLE(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *MathHTMLElement) DIR(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *MathHTMLElement) IfDIR(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *MathHTMLElement) RemoveDIR(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *MathHTMLElement) DRAGGABLE(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *MathHTMLElement) IfDRAGGABLE(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *MathHTMLElement) RemoveDRAGGABLE(v string) *MathHTMLElement {
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
func (e *MathHTMLElement) ENTERKEYHINT(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *MathHTMLElement) IfENTERKEYHINT(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *MathHTMLElement) RemoveENTERKEYHINT(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *MathHTMLElement) HIDDEN(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *MathHTMLElement) IfHIDDEN(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *MathHTMLElement) RemoveHIDDEN(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *MathHTMLElement) ID(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *MathHTMLElement) IfID(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *MathHTMLElement) RemoveID(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *MathHTMLElement) INERT() *MathHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *MathHTMLElement) IfINERT(cond bool) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *MathHTMLElement) RemoveINERT() *MathHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *MathHTMLElement) SetINERT(b bool) *MathHTMLElement {
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
func (e *MathHTMLElement) INPUTMODE(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *MathHTMLElement) IfINPUTMODE(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *MathHTMLElement) RemoveINPUTMODE(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *MathHTMLElement) IS(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *MathHTMLElement) IfIS(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *MathHTMLElement) RemoveIS(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *MathHTMLElement) ITEMID(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *MathHTMLElement) IfITEMID(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *MathHTMLElement) RemoveITEMID(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *MathHTMLElement) ITEMPROP(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *MathHTMLElement) IfITEMPROP(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *MathHTMLElement) RemoveITEMPROP(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *MathHTMLElement) ITEMREF(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *MathHTMLElement) IfITEMREF(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *MathHTMLElement) RemoveITEMREF(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *MathHTMLElement) ITEMSCOPE() *MathHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *MathHTMLElement) IfITEMSCOPE(cond bool) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *MathHTMLElement) RemoveITEMSCOPE() *MathHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *MathHTMLElement) SetITEMSCOPE(b bool) *MathHTMLElement {
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
func (e *MathHTMLElement) ITEMTYPE(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *MathHTMLElement) IfITEMTYPE(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *MathHTMLElement) RemoveITEMTYPE(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *MathHTMLElement) LANG(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *MathHTMLElement) IfLANG(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *MathHTMLElement) RemoveLANG(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *MathHTMLElement) NONCE(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *MathHTMLElement) IfNONCE(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *MathHTMLElement) RemoveNONCE(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *MathHTMLElement) POPOVER(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *MathHTMLElement) IfPOPOVER(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *MathHTMLElement) RemovePOPOVER(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *MathHTMLElement) SLOT(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *MathHTMLElement) IfSLOT(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *MathHTMLElement) RemoveSLOT(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *MathHTMLElement) SPELLCHECK(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *MathHTMLElement) IfSPELLCHECK(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *MathHTMLElement) RemoveSPELLCHECK(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *MathHTMLElement) STYLE(k, v string) *MathHTMLElement {
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

func (e *MathHTMLElement) IfSTYLE(cond bool, k string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *MathHTMLElement) RemoveSTYLE(k string) *MathHTMLElement {
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
func (e *MathHTMLElement) TABINDEX(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *MathHTMLElement) IfTABINDEX(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *MathHTMLElement) RemoveTABINDEX(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *MathHTMLElement) TITLE(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *MathHTMLElement) IfTITLE(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *MathHTMLElement) RemoveTITLE(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *MathHTMLElement) TRANSLATE(v string) *MathHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *MathHTMLElement) IfTRANSLATE(cond bool, v string) *MathHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *MathHTMLElement) RemoveTRANSLATE(v string) *MathHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
