package html

import (
	"fmt"
)

type KbdHTMLElement struct {
	*Element
}

func KBD(children ...ElementBuilder) *KbdHTMLElement {
	return &KbdHTMLElement{
		Element: &Element{
			Tag:           elementTagKBD,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *KbdHTMLElement) Children(children ...ElementBuilder) *KbdHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *KbdHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *KbdHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *KbdHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *KbdHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *KbdHTMLElement) Text(text string) *KbdHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *KbdHTMLElement) TextF(format string, args ...any) *KbdHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *KbdHTMLElement) Escaped(text string) *KbdHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *KbdHTMLElement) EscapedF(format string, args ...any) *KbdHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *KbdHTMLElement) CustomData(key, value string) *KbdHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *KbdHTMLElement) CustomDataRemove(key string) *KbdHTMLElement {
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
func (e *KbdHTMLElement) ACCESSKEY(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *KbdHTMLElement) IfACCESSKEY(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *KbdHTMLElement) RemoveACCESSKEY(v string) *KbdHTMLElement {
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
func (e *KbdHTMLElement) AUTOCAPITALIZE(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *KbdHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *KbdHTMLElement) RemoveAUTOCAPITALIZE(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *KbdHTMLElement) AUTOFOCUS() *KbdHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *KbdHTMLElement) IfAUTOFOCUS(cond bool) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *KbdHTMLElement) RemoveAUTOFOCUS() *KbdHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *KbdHTMLElement) SetAUTOFOCUS(b bool) *KbdHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *KbdHTMLElement) CLASS(v string) *KbdHTMLElement {
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

func (e *KbdHTMLElement) IfCLASS(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *KbdHTMLElement) SetCLASS(v string) *KbdHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *KbdHTMLElement) RemoveCLASS(v string) *KbdHTMLElement {
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
func (e *KbdHTMLElement) CONTENTEDITABLE(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *KbdHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *KbdHTMLElement) RemoveCONTENTEDITABLE(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *KbdHTMLElement) DIR(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *KbdHTMLElement) IfDIR(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *KbdHTMLElement) RemoveDIR(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *KbdHTMLElement) DRAGGABLE(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *KbdHTMLElement) IfDRAGGABLE(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *KbdHTMLElement) RemoveDRAGGABLE(v string) *KbdHTMLElement {
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
func (e *KbdHTMLElement) ENTERKEYHINT(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *KbdHTMLElement) IfENTERKEYHINT(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *KbdHTMLElement) RemoveENTERKEYHINT(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *KbdHTMLElement) HIDDEN(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *KbdHTMLElement) IfHIDDEN(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *KbdHTMLElement) RemoveHIDDEN(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *KbdHTMLElement) ID(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *KbdHTMLElement) IfID(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *KbdHTMLElement) RemoveID(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *KbdHTMLElement) INERT() *KbdHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *KbdHTMLElement) IfINERT(cond bool) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *KbdHTMLElement) RemoveINERT() *KbdHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *KbdHTMLElement) SetINERT(b bool) *KbdHTMLElement {
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
func (e *KbdHTMLElement) INPUTMODE(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *KbdHTMLElement) IfINPUTMODE(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *KbdHTMLElement) RemoveINPUTMODE(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *KbdHTMLElement) IS(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *KbdHTMLElement) IfIS(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *KbdHTMLElement) RemoveIS(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *KbdHTMLElement) ITEMID(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *KbdHTMLElement) IfITEMID(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *KbdHTMLElement) RemoveITEMID(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *KbdHTMLElement) ITEMPROP(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *KbdHTMLElement) IfITEMPROP(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *KbdHTMLElement) RemoveITEMPROP(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *KbdHTMLElement) ITEMREF(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *KbdHTMLElement) IfITEMREF(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *KbdHTMLElement) RemoveITEMREF(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *KbdHTMLElement) ITEMSCOPE() *KbdHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *KbdHTMLElement) IfITEMSCOPE(cond bool) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *KbdHTMLElement) RemoveITEMSCOPE() *KbdHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *KbdHTMLElement) SetITEMSCOPE(b bool) *KbdHTMLElement {
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
func (e *KbdHTMLElement) ITEMTYPE(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *KbdHTMLElement) IfITEMTYPE(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *KbdHTMLElement) RemoveITEMTYPE(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *KbdHTMLElement) LANG(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *KbdHTMLElement) IfLANG(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *KbdHTMLElement) RemoveLANG(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *KbdHTMLElement) NONCE(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *KbdHTMLElement) IfNONCE(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *KbdHTMLElement) RemoveNONCE(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *KbdHTMLElement) POPOVER(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *KbdHTMLElement) IfPOPOVER(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *KbdHTMLElement) RemovePOPOVER(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *KbdHTMLElement) SLOT(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *KbdHTMLElement) IfSLOT(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *KbdHTMLElement) RemoveSLOT(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *KbdHTMLElement) SPELLCHECK(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *KbdHTMLElement) IfSPELLCHECK(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *KbdHTMLElement) RemoveSPELLCHECK(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *KbdHTMLElement) STYLE(k, v string) *KbdHTMLElement {
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

func (e *KbdHTMLElement) IfSTYLE(cond bool, k string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *KbdHTMLElement) RemoveSTYLE(k string) *KbdHTMLElement {
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
func (e *KbdHTMLElement) TABINDEX(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *KbdHTMLElement) IfTABINDEX(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *KbdHTMLElement) RemoveTABINDEX(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *KbdHTMLElement) TITLE(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *KbdHTMLElement) IfTITLE(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *KbdHTMLElement) RemoveTITLE(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *KbdHTMLElement) TRANSLATE(v string) *KbdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *KbdHTMLElement) IfTRANSLATE(cond bool, v string) *KbdHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *KbdHTMLElement) RemoveTRANSLATE(v string) *KbdHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
