package html

import (
	"fmt"
)

type ThHTMLElement struct {
	*Element
}

func TH(children ...ElementBuilder) *ThHTMLElement {
	return &ThHTMLElement{
		Element: &Element{
			Tag:           elementTagTH,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *ThHTMLElement) Children(children ...ElementBuilder) *ThHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *ThHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ThHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *ThHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ThHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *ThHTMLElement) Text(text string) *ThHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *ThHTMLElement) TextF(format string, args ...any) *ThHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *ThHTMLElement) Escaped(text string) *ThHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *ThHTMLElement) EscapedF(format string, args ...any) *ThHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *ThHTMLElement) CustomData(key, value string) *ThHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ThHTMLElement) CustomDataRemove(key string) *ThHTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}

// ABBR sets the "abbr" attribute.
// Alternative label to use for the header cell when referencing the cell in other contexts
// Values values are constrained to:
//   - text
func (e *ThHTMLElement) ABBR(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeABBRKey] = v
	return e
}

func (e *ThHTMLElement) IfABBR(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ABBR(v)
}

func (e *ThHTMLElement) RemoveABBR(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeABBRKey)
	return e
}

// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//   - identical_to
//   - ordered_set_of_unique_space_separated_tokens
func (e *ThHTMLElement) ACCESSKEY(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *ThHTMLElement) IfACCESSKEY(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *ThHTMLElement) RemoveACCESSKEY(v string) *ThHTMLElement {
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
func (e *ThHTMLElement) AUTOCAPITALIZE(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *ThHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *ThHTMLElement) RemoveAUTOCAPITALIZE(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *ThHTMLElement) AUTOFOCUS() *ThHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *ThHTMLElement) IfAUTOFOCUS(cond bool) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *ThHTMLElement) RemoveAUTOFOCUS() *ThHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *ThHTMLElement) SetAUTOFOCUS(b bool) *ThHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *ThHTMLElement) CLASS(v string) *ThHTMLElement {
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

func (e *ThHTMLElement) IfCLASS(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *ThHTMLElement) SetCLASS(v string) *ThHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *ThHTMLElement) RemoveCLASS(v string) *ThHTMLElement {
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
func (e *ThHTMLElement) COLSPAN(v int) *ThHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeCOLSPANKey] = v
	return e
}

func (e *ThHTMLElement) IfCOLSPAN(cond bool, v int) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.COLSPAN(v)
}

func (e *ThHTMLElement) RemoveCOLSPAN(v int) *ThHTMLElement {
	delete(e.IntAttributes, attributeCOLSPANKey)
	return e
}

// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//   - false
//   - plaintext_only
//   - true
func (e *ThHTMLElement) CONTENTEDITABLE(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *ThHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *ThHTMLElement) RemoveCONTENTEDITABLE(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *ThHTMLElement) DIR(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *ThHTMLElement) IfDIR(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *ThHTMLElement) RemoveDIR(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *ThHTMLElement) DRAGGABLE(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *ThHTMLElement) IfDRAGGABLE(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *ThHTMLElement) RemoveDRAGGABLE(v string) *ThHTMLElement {
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
func (e *ThHTMLElement) ENTERKEYHINT(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *ThHTMLElement) IfENTERKEYHINT(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *ThHTMLElement) RemoveENTERKEYHINT(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HEADERS sets the "headers" attribute.
// The header cells for this cell
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ThHTMLElement) HEADERS(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHEADERSKey] = v
	return e
}

func (e *ThHTMLElement) IfHEADERS(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.HEADERS(v)
}

func (e *ThHTMLElement) RemoveHEADERS(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeHEADERSKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *ThHTMLElement) HIDDEN(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *ThHTMLElement) IfHIDDEN(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *ThHTMLElement) RemoveHIDDEN(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *ThHTMLElement) ID(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *ThHTMLElement) IfID(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *ThHTMLElement) RemoveID(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *ThHTMLElement) INERT() *ThHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *ThHTMLElement) IfINERT(cond bool) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *ThHTMLElement) RemoveINERT() *ThHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *ThHTMLElement) SetINERT(b bool) *ThHTMLElement {
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
func (e *ThHTMLElement) INPUTMODE(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *ThHTMLElement) IfINPUTMODE(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *ThHTMLElement) RemoveINPUTMODE(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *ThHTMLElement) IS(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *ThHTMLElement) IfIS(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *ThHTMLElement) RemoveIS(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *ThHTMLElement) ITEMID(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *ThHTMLElement) IfITEMID(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *ThHTMLElement) RemoveITEMID(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *ThHTMLElement) ITEMPROP(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *ThHTMLElement) IfITEMPROP(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *ThHTMLElement) RemoveITEMPROP(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ThHTMLElement) ITEMREF(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *ThHTMLElement) IfITEMREF(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *ThHTMLElement) RemoveITEMREF(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *ThHTMLElement) ITEMSCOPE() *ThHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *ThHTMLElement) IfITEMSCOPE(cond bool) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *ThHTMLElement) RemoveITEMSCOPE() *ThHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *ThHTMLElement) SetITEMSCOPE(b bool) *ThHTMLElement {
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
func (e *ThHTMLElement) ITEMTYPE(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *ThHTMLElement) IfITEMTYPE(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *ThHTMLElement) RemoveITEMTYPE(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ThHTMLElement) LANG(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *ThHTMLElement) IfLANG(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *ThHTMLElement) RemoveLANG(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *ThHTMLElement) NONCE(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *ThHTMLElement) IfNONCE(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *ThHTMLElement) RemoveNONCE(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *ThHTMLElement) POPOVER(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *ThHTMLElement) IfPOPOVER(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *ThHTMLElement) RemovePOPOVER(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// ROWSPAN sets the "rowspan" attribute.
// Number of rows that the cell is to span
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *ThHTMLElement) ROWSPAN(v int) *ThHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeROWSPANKey] = v
	return e
}

func (e *ThHTMLElement) IfROWSPAN(cond bool, v int) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.ROWSPAN(v)
}

func (e *ThHTMLElement) RemoveROWSPAN(v int) *ThHTMLElement {
	delete(e.IntAttributes, attributeROWSPANKey)
	return e
}

// SCOPE sets the "scope" attribute.
// Specifies which cells the header cell applies to
// Values values are constrained to:
//   - col
//   - colgroup
//   - row
//   - rowgroup
func (e *ThHTMLElement) SCOPE(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSCOPEKey] = v
	return e
}

func (e *ThHTMLElement) IfSCOPE(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.SCOPE(v)
}

func (e *ThHTMLElement) RemoveSCOPE(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeSCOPEKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *ThHTMLElement) SLOT(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *ThHTMLElement) IfSLOT(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *ThHTMLElement) RemoveSLOT(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *ThHTMLElement) SPELLCHECK(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *ThHTMLElement) IfSPELLCHECK(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *ThHTMLElement) RemoveSPELLCHECK(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ThHTMLElement) STYLE(k, v string) *ThHTMLElement {
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

func (e *ThHTMLElement) IfSTYLE(cond bool, k string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *ThHTMLElement) RemoveSTYLE(k string) *ThHTMLElement {
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
func (e *ThHTMLElement) TABINDEX(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *ThHTMLElement) IfTABINDEX(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *ThHTMLElement) RemoveTABINDEX(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *ThHTMLElement) TITLE(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *ThHTMLElement) IfTITLE(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *ThHTMLElement) RemoveTITLE(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *ThHTMLElement) TRANSLATE(v string) *ThHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *ThHTMLElement) IfTRANSLATE(cond bool, v string) *ThHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *ThHTMLElement) RemoveTRANSLATE(v string) *ThHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
