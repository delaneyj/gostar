package html

import (
	"fmt"
)

type HeaderHTMLElement struct {
	*Element
}

func HEADER(children ...ElementBuilder) *HeaderHTMLElement {
	return &HeaderHTMLElement{
		Element: &Element{
			Tag:           elementTagHEADER,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *HeaderHTMLElement) Children(children ...ElementBuilder) *HeaderHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *HeaderHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *HeaderHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *HeaderHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *HeaderHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *HeaderHTMLElement) Text(text string) *HeaderHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *HeaderHTMLElement) TextF(format string, args ...any) *HeaderHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *HeaderHTMLElement) Escaped(text string) *HeaderHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *HeaderHTMLElement) EscapedF(format string, args ...any) *HeaderHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *HeaderHTMLElement) CustomData(key, value string) *HeaderHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *HeaderHTMLElement) CustomDataRemove(key string) *HeaderHTMLElement {
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
func (e *HeaderHTMLElement) ACCESSKEY(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *HeaderHTMLElement) IfACCESSKEY(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *HeaderHTMLElement) RemoveACCESSKEY(v string) *HeaderHTMLElement {
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
func (e *HeaderHTMLElement) AUTOCAPITALIZE(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *HeaderHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *HeaderHTMLElement) RemoveAUTOCAPITALIZE(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *HeaderHTMLElement) AUTOFOCUS() *HeaderHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *HeaderHTMLElement) IfAUTOFOCUS(cond bool) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *HeaderHTMLElement) RemoveAUTOFOCUS() *HeaderHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *HeaderHTMLElement) SetAUTOFOCUS(b bool) *HeaderHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *HeaderHTMLElement) CLASS(v string) *HeaderHTMLElement {
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

func (e *HeaderHTMLElement) IfCLASS(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *HeaderHTMLElement) SetCLASS(v string) *HeaderHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *HeaderHTMLElement) RemoveCLASS(v string) *HeaderHTMLElement {
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
func (e *HeaderHTMLElement) CONTENTEDITABLE(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *HeaderHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *HeaderHTMLElement) RemoveCONTENTEDITABLE(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *HeaderHTMLElement) DIR(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *HeaderHTMLElement) IfDIR(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *HeaderHTMLElement) RemoveDIR(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *HeaderHTMLElement) DRAGGABLE(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *HeaderHTMLElement) IfDRAGGABLE(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *HeaderHTMLElement) RemoveDRAGGABLE(v string) *HeaderHTMLElement {
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
func (e *HeaderHTMLElement) ENTERKEYHINT(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *HeaderHTMLElement) IfENTERKEYHINT(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *HeaderHTMLElement) RemoveENTERKEYHINT(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *HeaderHTMLElement) HIDDEN(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *HeaderHTMLElement) IfHIDDEN(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *HeaderHTMLElement) RemoveHIDDEN(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *HeaderHTMLElement) ID(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *HeaderHTMLElement) IfID(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *HeaderHTMLElement) RemoveID(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *HeaderHTMLElement) INERT() *HeaderHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *HeaderHTMLElement) IfINERT(cond bool) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *HeaderHTMLElement) RemoveINERT() *HeaderHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *HeaderHTMLElement) SetINERT(b bool) *HeaderHTMLElement {
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
func (e *HeaderHTMLElement) INPUTMODE(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *HeaderHTMLElement) IfINPUTMODE(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *HeaderHTMLElement) RemoveINPUTMODE(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *HeaderHTMLElement) IS(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *HeaderHTMLElement) IfIS(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *HeaderHTMLElement) RemoveIS(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *HeaderHTMLElement) ITEMID(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *HeaderHTMLElement) IfITEMID(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *HeaderHTMLElement) RemoveITEMID(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *HeaderHTMLElement) ITEMPROP(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *HeaderHTMLElement) IfITEMPROP(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *HeaderHTMLElement) RemoveITEMPROP(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *HeaderHTMLElement) ITEMREF(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *HeaderHTMLElement) IfITEMREF(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *HeaderHTMLElement) RemoveITEMREF(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *HeaderHTMLElement) ITEMSCOPE() *HeaderHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *HeaderHTMLElement) IfITEMSCOPE(cond bool) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *HeaderHTMLElement) RemoveITEMSCOPE() *HeaderHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *HeaderHTMLElement) SetITEMSCOPE(b bool) *HeaderHTMLElement {
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
func (e *HeaderHTMLElement) ITEMTYPE(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *HeaderHTMLElement) IfITEMTYPE(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *HeaderHTMLElement) RemoveITEMTYPE(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *HeaderHTMLElement) LANG(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *HeaderHTMLElement) IfLANG(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *HeaderHTMLElement) RemoveLANG(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *HeaderHTMLElement) NONCE(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *HeaderHTMLElement) IfNONCE(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *HeaderHTMLElement) RemoveNONCE(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *HeaderHTMLElement) POPOVER(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *HeaderHTMLElement) IfPOPOVER(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *HeaderHTMLElement) RemovePOPOVER(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *HeaderHTMLElement) SLOT(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *HeaderHTMLElement) IfSLOT(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *HeaderHTMLElement) RemoveSLOT(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *HeaderHTMLElement) SPELLCHECK(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *HeaderHTMLElement) IfSPELLCHECK(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *HeaderHTMLElement) RemoveSPELLCHECK(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *HeaderHTMLElement) STYLE(k, v string) *HeaderHTMLElement {
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

func (e *HeaderHTMLElement) IfSTYLE(cond bool, k string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *HeaderHTMLElement) RemoveSTYLE(k string) *HeaderHTMLElement {
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
func (e *HeaderHTMLElement) TABINDEX(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *HeaderHTMLElement) IfTABINDEX(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *HeaderHTMLElement) RemoveTABINDEX(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *HeaderHTMLElement) TITLE(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *HeaderHTMLElement) IfTITLE(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *HeaderHTMLElement) RemoveTITLE(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *HeaderHTMLElement) TRANSLATE(v string) *HeaderHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *HeaderHTMLElement) IfTRANSLATE(cond bool, v string) *HeaderHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *HeaderHTMLElement) RemoveTRANSLATE(v string) *HeaderHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
