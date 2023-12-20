package html

import (
	"fmt"
)

type UlHTMLElement struct {
	*Element
}

func UL(children ...ElementBuilder) *UlHTMLElement {
	return &UlHTMLElement{
		Element: &Element{
			Tag:           elementTagUL,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *UlHTMLElement) Children(children ...ElementBuilder) *UlHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *UlHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *UlHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *UlHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *UlHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *UlHTMLElement) Text(text string) *UlHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *UlHTMLElement) TextF(format string, args ...any) *UlHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *UlHTMLElement) Escaped(text string) *UlHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *UlHTMLElement) EscapedF(format string, args ...any) *UlHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *UlHTMLElement) CustomData(key, value string) *UlHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *UlHTMLElement) CustomDataRemove(key string) *UlHTMLElement {
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
func (e *UlHTMLElement) ACCESSKEY(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *UlHTMLElement) IfACCESSKEY(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *UlHTMLElement) RemoveACCESSKEY(v string) *UlHTMLElement {
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
func (e *UlHTMLElement) AUTOCAPITALIZE(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *UlHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *UlHTMLElement) RemoveAUTOCAPITALIZE(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *UlHTMLElement) AUTOFOCUS() *UlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *UlHTMLElement) IfAUTOFOCUS(cond bool) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *UlHTMLElement) RemoveAUTOFOCUS() *UlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *UlHTMLElement) SetAUTOFOCUS(b bool) *UlHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *UlHTMLElement) CLASS(v string) *UlHTMLElement {
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

func (e *UlHTMLElement) IfCLASS(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *UlHTMLElement) SetCLASS(v string) *UlHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *UlHTMLElement) RemoveCLASS(v string) *UlHTMLElement {
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
func (e *UlHTMLElement) CONTENTEDITABLE(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *UlHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *UlHTMLElement) RemoveCONTENTEDITABLE(v string) *UlHTMLElement {
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
func (e *UlHTMLElement) DIR(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *UlHTMLElement) IfDIR(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *UlHTMLElement) RemoveDIR(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *UlHTMLElement) DRAGGABLE(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *UlHTMLElement) IfDRAGGABLE(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *UlHTMLElement) RemoveDRAGGABLE(v string) *UlHTMLElement {
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
func (e *UlHTMLElement) ENTERKEYHINT(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *UlHTMLElement) IfENTERKEYHINT(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *UlHTMLElement) RemoveENTERKEYHINT(v string) *UlHTMLElement {
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
func (e *UlHTMLElement) HIDDEN(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *UlHTMLElement) IfHIDDEN(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *UlHTMLElement) RemoveHIDDEN(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *UlHTMLElement) ID(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *UlHTMLElement) IfID(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *UlHTMLElement) RemoveID(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *UlHTMLElement) INERT() *UlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *UlHTMLElement) IfINERT(cond bool) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *UlHTMLElement) RemoveINERT() *UlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *UlHTMLElement) SetINERT(b bool) *UlHTMLElement {
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
func (e *UlHTMLElement) INPUTMODE(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *UlHTMLElement) IfINPUTMODE(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *UlHTMLElement) RemoveINPUTMODE(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *UlHTMLElement) IS(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *UlHTMLElement) IfIS(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *UlHTMLElement) RemoveIS(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *UlHTMLElement) ITEMID(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *UlHTMLElement) IfITEMID(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *UlHTMLElement) RemoveITEMID(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *UlHTMLElement) ITEMPROP(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *UlHTMLElement) IfITEMPROP(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *UlHTMLElement) RemoveITEMPROP(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *UlHTMLElement) ITEMREF(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *UlHTMLElement) IfITEMREF(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *UlHTMLElement) RemoveITEMREF(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *UlHTMLElement) ITEMSCOPE() *UlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *UlHTMLElement) IfITEMSCOPE(cond bool) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *UlHTMLElement) RemoveITEMSCOPE() *UlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *UlHTMLElement) SetITEMSCOPE(b bool) *UlHTMLElement {
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
func (e *UlHTMLElement) ITEMTYPE(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *UlHTMLElement) IfITEMTYPE(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *UlHTMLElement) RemoveITEMTYPE(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *UlHTMLElement) LANG(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *UlHTMLElement) IfLANG(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *UlHTMLElement) RemoveLANG(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *UlHTMLElement) NONCE(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *UlHTMLElement) IfNONCE(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *UlHTMLElement) RemoveNONCE(v string) *UlHTMLElement {
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
func (e *UlHTMLElement) POPOVER(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *UlHTMLElement) IfPOPOVER(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *UlHTMLElement) RemovePOPOVER(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *UlHTMLElement) SLOT(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *UlHTMLElement) IfSLOT(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *UlHTMLElement) RemoveSLOT(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *UlHTMLElement) SPELLCHECK(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *UlHTMLElement) IfSPELLCHECK(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *UlHTMLElement) RemoveSPELLCHECK(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *UlHTMLElement) STYLE(k, v string) *UlHTMLElement {
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

func (e *UlHTMLElement) IfSTYLE(cond bool, k string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *UlHTMLElement) RemoveSTYLE(k string) *UlHTMLElement {
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
func (e *UlHTMLElement) TABINDEX(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *UlHTMLElement) IfTABINDEX(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *UlHTMLElement) RemoveTABINDEX(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *UlHTMLElement) TITLE(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *UlHTMLElement) IfTITLE(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *UlHTMLElement) RemoveTITLE(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *UlHTMLElement) TRANSLATE(v string) *UlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *UlHTMLElement) IfTRANSLATE(cond bool, v string) *UlHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *UlHTMLElement) RemoveTRANSLATE(v string) *UlHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
