package html

import (
	"fmt"
)

type BrHTMLElement struct {
	*Element
}

func BR(children ...ElementBuilder) *BrHTMLElement {
	return &BrHTMLElement{
		Element: &Element{
			Tag:           elementTagBR,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *BrHTMLElement) Children(children ...ElementBuilder) *BrHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *BrHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *BrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *BrHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *BrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *BrHTMLElement) Text(text string) *BrHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *BrHTMLElement) TextF(format string, args ...any) *BrHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *BrHTMLElement) Escaped(text string) *BrHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *BrHTMLElement) EscapedF(format string, args ...any) *BrHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *BrHTMLElement) CustomData(key, value string) *BrHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *BrHTMLElement) CustomDataRemove(key string) *BrHTMLElement {
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
func (e *BrHTMLElement) ACCESSKEY(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *BrHTMLElement) IfACCESSKEY(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *BrHTMLElement) RemoveACCESSKEY(v string) *BrHTMLElement {
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
func (e *BrHTMLElement) AUTOCAPITALIZE(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *BrHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *BrHTMLElement) RemoveAUTOCAPITALIZE(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *BrHTMLElement) AUTOFOCUS() *BrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *BrHTMLElement) IfAUTOFOCUS(cond bool) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *BrHTMLElement) RemoveAUTOFOCUS() *BrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *BrHTMLElement) SetAUTOFOCUS(b bool) *BrHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *BrHTMLElement) CLASS(v string) *BrHTMLElement {
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

func (e *BrHTMLElement) IfCLASS(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *BrHTMLElement) SetCLASS(v string) *BrHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *BrHTMLElement) RemoveCLASS(v string) *BrHTMLElement {
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
func (e *BrHTMLElement) CONTENTEDITABLE(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *BrHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *BrHTMLElement) RemoveCONTENTEDITABLE(v string) *BrHTMLElement {
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
func (e *BrHTMLElement) DIR(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *BrHTMLElement) IfDIR(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *BrHTMLElement) RemoveDIR(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *BrHTMLElement) DRAGGABLE(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *BrHTMLElement) IfDRAGGABLE(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *BrHTMLElement) RemoveDRAGGABLE(v string) *BrHTMLElement {
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
func (e *BrHTMLElement) ENTERKEYHINT(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *BrHTMLElement) IfENTERKEYHINT(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *BrHTMLElement) RemoveENTERKEYHINT(v string) *BrHTMLElement {
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
func (e *BrHTMLElement) HIDDEN(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *BrHTMLElement) IfHIDDEN(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *BrHTMLElement) RemoveHIDDEN(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *BrHTMLElement) ID(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *BrHTMLElement) IfID(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *BrHTMLElement) RemoveID(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *BrHTMLElement) INERT() *BrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *BrHTMLElement) IfINERT(cond bool) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *BrHTMLElement) RemoveINERT() *BrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *BrHTMLElement) SetINERT(b bool) *BrHTMLElement {
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
func (e *BrHTMLElement) INPUTMODE(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *BrHTMLElement) IfINPUTMODE(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *BrHTMLElement) RemoveINPUTMODE(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *BrHTMLElement) IS(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *BrHTMLElement) IfIS(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *BrHTMLElement) RemoveIS(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *BrHTMLElement) ITEMID(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *BrHTMLElement) IfITEMID(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *BrHTMLElement) RemoveITEMID(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *BrHTMLElement) ITEMPROP(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *BrHTMLElement) IfITEMPROP(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *BrHTMLElement) RemoveITEMPROP(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *BrHTMLElement) ITEMREF(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *BrHTMLElement) IfITEMREF(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *BrHTMLElement) RemoveITEMREF(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *BrHTMLElement) ITEMSCOPE() *BrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *BrHTMLElement) IfITEMSCOPE(cond bool) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *BrHTMLElement) RemoveITEMSCOPE() *BrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *BrHTMLElement) SetITEMSCOPE(b bool) *BrHTMLElement {
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
func (e *BrHTMLElement) ITEMTYPE(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *BrHTMLElement) IfITEMTYPE(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *BrHTMLElement) RemoveITEMTYPE(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *BrHTMLElement) LANG(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *BrHTMLElement) IfLANG(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *BrHTMLElement) RemoveLANG(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *BrHTMLElement) NONCE(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *BrHTMLElement) IfNONCE(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *BrHTMLElement) RemoveNONCE(v string) *BrHTMLElement {
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
func (e *BrHTMLElement) POPOVER(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *BrHTMLElement) IfPOPOVER(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *BrHTMLElement) RemovePOPOVER(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *BrHTMLElement) SLOT(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *BrHTMLElement) IfSLOT(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *BrHTMLElement) RemoveSLOT(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *BrHTMLElement) SPELLCHECK(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *BrHTMLElement) IfSPELLCHECK(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *BrHTMLElement) RemoveSPELLCHECK(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *BrHTMLElement) STYLE(k, v string) *BrHTMLElement {
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

func (e *BrHTMLElement) IfSTYLE(cond bool, k string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *BrHTMLElement) RemoveSTYLE(k string) *BrHTMLElement {
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
func (e *BrHTMLElement) TABINDEX(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *BrHTMLElement) IfTABINDEX(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *BrHTMLElement) RemoveTABINDEX(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *BrHTMLElement) TITLE(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *BrHTMLElement) IfTITLE(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *BrHTMLElement) RemoveTITLE(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *BrHTMLElement) TRANSLATE(v string) *BrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *BrHTMLElement) IfTRANSLATE(cond bool, v string) *BrHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *BrHTMLElement) RemoveTRANSLATE(v string) *BrHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
