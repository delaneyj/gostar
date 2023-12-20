package html

import (
	"fmt"
)

type TdHTMLElement struct {
	*Element
}

func TD(children ...ElementBuilder) *TdHTMLElement {
	return &TdHTMLElement{
		Element: &Element{
			Tag:           elementTagTD,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *TdHTMLElement) Children(children ...ElementBuilder) *TdHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *TdHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TdHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *TdHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TdHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *TdHTMLElement) Text(text string) *TdHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *TdHTMLElement) TextF(format string, args ...any) *TdHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *TdHTMLElement) Escaped(text string) *TdHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *TdHTMLElement) EscapedF(format string, args ...any) *TdHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TdHTMLElement) CustomData(key, value string) *TdHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TdHTMLElement) CustomDataRemove(key string) *TdHTMLElement {
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
func (e *TdHTMLElement) ACCESSKEY(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *TdHTMLElement) IfACCESSKEY(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *TdHTMLElement) RemoveACCESSKEY(v string) *TdHTMLElement {
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
func (e *TdHTMLElement) AUTOCAPITALIZE(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *TdHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *TdHTMLElement) RemoveAUTOCAPITALIZE(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *TdHTMLElement) AUTOFOCUS() *TdHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *TdHTMLElement) IfAUTOFOCUS(cond bool) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *TdHTMLElement) RemoveAUTOFOCUS() *TdHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *TdHTMLElement) SetAUTOFOCUS(b bool) *TdHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *TdHTMLElement) CLASS(v string) *TdHTMLElement {
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

func (e *TdHTMLElement) IfCLASS(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *TdHTMLElement) SetCLASS(v string) *TdHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *TdHTMLElement) RemoveCLASS(v string) *TdHTMLElement {
	kv, ok := e.DelimitedStringAttributes[attributeCLASSKey]
	if !ok {
		return e
	}
	kv.Remove(v)
	return e
}

// COLSPAN sets the "colspan" attribute.
// Number of columns that the cell is to span
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *TdHTMLElement) COLSPAN(v int) *TdHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeCOLSPANKey] = v
	return e
}

func (e *TdHTMLElement) IfCOLSPAN(cond bool, v int) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.COLSPAN(v)
}

func (e *TdHTMLElement) RemoveCOLSPAN(v int) *TdHTMLElement {
	delete(e.IntAttributes, attributeCOLSPANKey)
	return e
}

// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//   - true
//   - plaintext_only
//   - false
func (e *TdHTMLElement) CONTENTEDITABLE(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *TdHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *TdHTMLElement) RemoveCONTENTEDITABLE(v string) *TdHTMLElement {
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
func (e *TdHTMLElement) DIR(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *TdHTMLElement) IfDIR(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *TdHTMLElement) RemoveDIR(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *TdHTMLElement) DRAGGABLE(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *TdHTMLElement) IfDRAGGABLE(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *TdHTMLElement) RemoveDRAGGABLE(v string) *TdHTMLElement {
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
func (e *TdHTMLElement) ENTERKEYHINT(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *TdHTMLElement) IfENTERKEYHINT(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *TdHTMLElement) RemoveENTERKEYHINT(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HEADERS sets the "headers" attribute.
// The header cells for this cell
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TdHTMLElement) HEADERS(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHEADERSKey] = v
	return e
}

func (e *TdHTMLElement) IfHEADERS(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.HEADERS(v)
}

func (e *TdHTMLElement) RemoveHEADERS(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeHEADERSKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (e *TdHTMLElement) HIDDEN(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *TdHTMLElement) IfHIDDEN(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *TdHTMLElement) RemoveHIDDEN(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *TdHTMLElement) ID(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *TdHTMLElement) IfID(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *TdHTMLElement) RemoveID(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *TdHTMLElement) INERT() *TdHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *TdHTMLElement) IfINERT(cond bool) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *TdHTMLElement) RemoveINERT() *TdHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *TdHTMLElement) SetINERT(b bool) *TdHTMLElement {
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
func (e *TdHTMLElement) INPUTMODE(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *TdHTMLElement) IfINPUTMODE(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *TdHTMLElement) RemoveINPUTMODE(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *TdHTMLElement) IS(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *TdHTMLElement) IfIS(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *TdHTMLElement) RemoveIS(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *TdHTMLElement) ITEMID(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *TdHTMLElement) IfITEMID(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *TdHTMLElement) RemoveITEMID(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *TdHTMLElement) ITEMPROP(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *TdHTMLElement) IfITEMPROP(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *TdHTMLElement) RemoveITEMPROP(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TdHTMLElement) ITEMREF(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *TdHTMLElement) IfITEMREF(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *TdHTMLElement) RemoveITEMREF(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *TdHTMLElement) ITEMSCOPE() *TdHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *TdHTMLElement) IfITEMSCOPE(cond bool) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *TdHTMLElement) RemoveITEMSCOPE() *TdHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *TdHTMLElement) SetITEMSCOPE(b bool) *TdHTMLElement {
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
func (e *TdHTMLElement) ITEMTYPE(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *TdHTMLElement) IfITEMTYPE(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *TdHTMLElement) RemoveITEMTYPE(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TdHTMLElement) LANG(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *TdHTMLElement) IfLANG(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *TdHTMLElement) RemoveLANG(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *TdHTMLElement) NONCE(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *TdHTMLElement) IfNONCE(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *TdHTMLElement) RemoveNONCE(v string) *TdHTMLElement {
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
func (e *TdHTMLElement) POPOVER(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *TdHTMLElement) IfPOPOVER(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *TdHTMLElement) RemovePOPOVER(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// ROWSPAN sets the "rowspan" attribute.
// Number of rows that the cell is to span
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *TdHTMLElement) ROWSPAN(v int) *TdHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeROWSPANKey] = v
	return e
}

func (e *TdHTMLElement) IfROWSPAN(cond bool, v int) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.ROWSPAN(v)
}

func (e *TdHTMLElement) RemoveROWSPAN(v int) *TdHTMLElement {
	delete(e.IntAttributes, attributeROWSPANKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *TdHTMLElement) SLOT(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *TdHTMLElement) IfSLOT(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *TdHTMLElement) RemoveSLOT(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *TdHTMLElement) SPELLCHECK(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *TdHTMLElement) IfSPELLCHECK(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *TdHTMLElement) RemoveSPELLCHECK(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TdHTMLElement) STYLE(k, v string) *TdHTMLElement {
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

func (e *TdHTMLElement) IfSTYLE(cond bool, k string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *TdHTMLElement) RemoveSTYLE(k string) *TdHTMLElement {
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
func (e *TdHTMLElement) TABINDEX(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *TdHTMLElement) IfTABINDEX(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *TdHTMLElement) RemoveTABINDEX(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *TdHTMLElement) TITLE(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *TdHTMLElement) IfTITLE(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *TdHTMLElement) RemoveTITLE(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *TdHTMLElement) TRANSLATE(v string) *TdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *TdHTMLElement) IfTRANSLATE(cond bool, v string) *TdHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *TdHTMLElement) RemoveTRANSLATE(v string) *TdHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
