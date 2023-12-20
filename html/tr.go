package html

import (
	"fmt"
)

type TrHTMLElement struct {
	*Element
}

func TR(children ...ElementBuilder) *TrHTMLElement {
	return &TrHTMLElement{
		Element: &Element{
			Tag:           elementTagTR,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *TrHTMLElement) Children(children ...ElementBuilder) *TrHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *TrHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *TrHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *TrHTMLElement) Text(text string) *TrHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *TrHTMLElement) TextF(format string, args ...any) *TrHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *TrHTMLElement) Escaped(text string) *TrHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *TrHTMLElement) EscapedF(format string, args ...any) *TrHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TrHTMLElement) CustomData(key, value string) *TrHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TrHTMLElement) CustomDataRemove(key string) *TrHTMLElement {
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
func (e *TrHTMLElement) ACCESSKEY(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *TrHTMLElement) IfACCESSKEY(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *TrHTMLElement) RemoveACCESSKEY(v string) *TrHTMLElement {
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
func (e *TrHTMLElement) AUTOCAPITALIZE(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *TrHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *TrHTMLElement) RemoveAUTOCAPITALIZE(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *TrHTMLElement) AUTOFOCUS() *TrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *TrHTMLElement) IfAUTOFOCUS(cond bool) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *TrHTMLElement) RemoveAUTOFOCUS() *TrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *TrHTMLElement) SetAUTOFOCUS(b bool) *TrHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *TrHTMLElement) CLASS(v string) *TrHTMLElement {
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

func (e *TrHTMLElement) IfCLASS(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *TrHTMLElement) SetCLASS(v string) *TrHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *TrHTMLElement) RemoveCLASS(v string) *TrHTMLElement {
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
func (e *TrHTMLElement) CONTENTEDITABLE(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *TrHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *TrHTMLElement) RemoveCONTENTEDITABLE(v string) *TrHTMLElement {
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
func (e *TrHTMLElement) DIR(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *TrHTMLElement) IfDIR(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *TrHTMLElement) RemoveDIR(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *TrHTMLElement) DRAGGABLE(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *TrHTMLElement) IfDRAGGABLE(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *TrHTMLElement) RemoveDRAGGABLE(v string) *TrHTMLElement {
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
func (e *TrHTMLElement) ENTERKEYHINT(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *TrHTMLElement) IfENTERKEYHINT(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *TrHTMLElement) RemoveENTERKEYHINT(v string) *TrHTMLElement {
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
func (e *TrHTMLElement) HIDDEN(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *TrHTMLElement) IfHIDDEN(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *TrHTMLElement) RemoveHIDDEN(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *TrHTMLElement) ID(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *TrHTMLElement) IfID(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *TrHTMLElement) RemoveID(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *TrHTMLElement) INERT() *TrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *TrHTMLElement) IfINERT(cond bool) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *TrHTMLElement) RemoveINERT() *TrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *TrHTMLElement) SetINERT(b bool) *TrHTMLElement {
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
func (e *TrHTMLElement) INPUTMODE(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *TrHTMLElement) IfINPUTMODE(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *TrHTMLElement) RemoveINPUTMODE(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *TrHTMLElement) IS(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *TrHTMLElement) IfIS(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *TrHTMLElement) RemoveIS(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *TrHTMLElement) ITEMID(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *TrHTMLElement) IfITEMID(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *TrHTMLElement) RemoveITEMID(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *TrHTMLElement) ITEMPROP(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *TrHTMLElement) IfITEMPROP(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *TrHTMLElement) RemoveITEMPROP(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TrHTMLElement) ITEMREF(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *TrHTMLElement) IfITEMREF(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *TrHTMLElement) RemoveITEMREF(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *TrHTMLElement) ITEMSCOPE() *TrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *TrHTMLElement) IfITEMSCOPE(cond bool) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *TrHTMLElement) RemoveITEMSCOPE() *TrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *TrHTMLElement) SetITEMSCOPE(b bool) *TrHTMLElement {
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
func (e *TrHTMLElement) ITEMTYPE(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *TrHTMLElement) IfITEMTYPE(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *TrHTMLElement) RemoveITEMTYPE(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TrHTMLElement) LANG(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *TrHTMLElement) IfLANG(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *TrHTMLElement) RemoveLANG(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *TrHTMLElement) NONCE(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *TrHTMLElement) IfNONCE(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *TrHTMLElement) RemoveNONCE(v string) *TrHTMLElement {
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
func (e *TrHTMLElement) POPOVER(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *TrHTMLElement) IfPOPOVER(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *TrHTMLElement) RemovePOPOVER(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *TrHTMLElement) SLOT(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *TrHTMLElement) IfSLOT(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *TrHTMLElement) RemoveSLOT(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *TrHTMLElement) SPELLCHECK(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *TrHTMLElement) IfSPELLCHECK(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *TrHTMLElement) RemoveSPELLCHECK(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TrHTMLElement) STYLE(k, v string) *TrHTMLElement {
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

func (e *TrHTMLElement) IfSTYLE(cond bool, k string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *TrHTMLElement) RemoveSTYLE(k string) *TrHTMLElement {
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
func (e *TrHTMLElement) TABINDEX(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *TrHTMLElement) IfTABINDEX(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *TrHTMLElement) RemoveTABINDEX(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *TrHTMLElement) TITLE(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *TrHTMLElement) IfTITLE(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *TrHTMLElement) RemoveTITLE(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *TrHTMLElement) TRANSLATE(v string) *TrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *TrHTMLElement) IfTRANSLATE(cond bool, v string) *TrHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *TrHTMLElement) RemoveTRANSLATE(v string) *TrHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
