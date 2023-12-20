package html

import (
	"fmt"
)

type TableHTMLElement struct {
	*Element
}

func TABLE(children ...ElementBuilder) *TableHTMLElement {
	return &TableHTMLElement{
		Element: &Element{
			Tag:           elementTagTABLE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *TableHTMLElement) Children(children ...ElementBuilder) *TableHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *TableHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TableHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *TableHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TableHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *TableHTMLElement) Text(text string) *TableHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *TableHTMLElement) TextF(format string, args ...any) *TableHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *TableHTMLElement) Escaped(text string) *TableHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *TableHTMLElement) EscapedF(format string, args ...any) *TableHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TableHTMLElement) CustomData(key, value string) *TableHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TableHTMLElement) CustomDataRemove(key string) *TableHTMLElement {
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
func (e *TableHTMLElement) ACCESSKEY(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *TableHTMLElement) IfACCESSKEY(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *TableHTMLElement) RemoveACCESSKEY(v string) *TableHTMLElement {
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
func (e *TableHTMLElement) AUTOCAPITALIZE(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *TableHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *TableHTMLElement) RemoveAUTOCAPITALIZE(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *TableHTMLElement) AUTOFOCUS() *TableHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *TableHTMLElement) IfAUTOFOCUS(cond bool) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *TableHTMLElement) RemoveAUTOFOCUS() *TableHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *TableHTMLElement) SetAUTOFOCUS(b bool) *TableHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *TableHTMLElement) CLASS(v string) *TableHTMLElement {
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

func (e *TableHTMLElement) IfCLASS(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *TableHTMLElement) SetCLASS(v string) *TableHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *TableHTMLElement) RemoveCLASS(v string) *TableHTMLElement {
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
func (e *TableHTMLElement) CONTENTEDITABLE(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *TableHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *TableHTMLElement) RemoveCONTENTEDITABLE(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *TableHTMLElement) DIR(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *TableHTMLElement) IfDIR(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *TableHTMLElement) RemoveDIR(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *TableHTMLElement) DRAGGABLE(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *TableHTMLElement) IfDRAGGABLE(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *TableHTMLElement) RemoveDRAGGABLE(v string) *TableHTMLElement {
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
func (e *TableHTMLElement) ENTERKEYHINT(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *TableHTMLElement) IfENTERKEYHINT(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *TableHTMLElement) RemoveENTERKEYHINT(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *TableHTMLElement) HIDDEN(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *TableHTMLElement) IfHIDDEN(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *TableHTMLElement) RemoveHIDDEN(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *TableHTMLElement) ID(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *TableHTMLElement) IfID(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *TableHTMLElement) RemoveID(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *TableHTMLElement) INERT() *TableHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *TableHTMLElement) IfINERT(cond bool) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *TableHTMLElement) RemoveINERT() *TableHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *TableHTMLElement) SetINERT(b bool) *TableHTMLElement {
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
func (e *TableHTMLElement) INPUTMODE(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *TableHTMLElement) IfINPUTMODE(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *TableHTMLElement) RemoveINPUTMODE(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *TableHTMLElement) IS(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *TableHTMLElement) IfIS(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *TableHTMLElement) RemoveIS(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *TableHTMLElement) ITEMID(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *TableHTMLElement) IfITEMID(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *TableHTMLElement) RemoveITEMID(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *TableHTMLElement) ITEMPROP(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *TableHTMLElement) IfITEMPROP(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *TableHTMLElement) RemoveITEMPROP(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TableHTMLElement) ITEMREF(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *TableHTMLElement) IfITEMREF(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *TableHTMLElement) RemoveITEMREF(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *TableHTMLElement) ITEMSCOPE() *TableHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *TableHTMLElement) IfITEMSCOPE(cond bool) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *TableHTMLElement) RemoveITEMSCOPE() *TableHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *TableHTMLElement) SetITEMSCOPE(b bool) *TableHTMLElement {
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
func (e *TableHTMLElement) ITEMTYPE(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *TableHTMLElement) IfITEMTYPE(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *TableHTMLElement) RemoveITEMTYPE(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TableHTMLElement) LANG(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *TableHTMLElement) IfLANG(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *TableHTMLElement) RemoveLANG(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *TableHTMLElement) NONCE(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *TableHTMLElement) IfNONCE(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *TableHTMLElement) RemoveNONCE(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *TableHTMLElement) POPOVER(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *TableHTMLElement) IfPOPOVER(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *TableHTMLElement) RemovePOPOVER(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *TableHTMLElement) SLOT(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *TableHTMLElement) IfSLOT(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *TableHTMLElement) RemoveSLOT(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *TableHTMLElement) SPELLCHECK(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *TableHTMLElement) IfSPELLCHECK(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *TableHTMLElement) RemoveSPELLCHECK(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TableHTMLElement) STYLE(k, v string) *TableHTMLElement {
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

func (e *TableHTMLElement) IfSTYLE(cond bool, k string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *TableHTMLElement) RemoveSTYLE(k string) *TableHTMLElement {
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
func (e *TableHTMLElement) TABINDEX(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *TableHTMLElement) IfTABINDEX(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *TableHTMLElement) RemoveTABINDEX(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *TableHTMLElement) TITLE(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *TableHTMLElement) IfTITLE(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *TableHTMLElement) RemoveTITLE(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *TableHTMLElement) TRANSLATE(v string) *TableHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *TableHTMLElement) IfTRANSLATE(cond bool, v string) *TableHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *TableHTMLElement) RemoveTRANSLATE(v string) *TableHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
