package html

import (
	"fmt"
)

type TimeHTMLElement struct {
	*Element
}

func TIME(children ...ElementBuilder) *TimeHTMLElement {
	return &TimeHTMLElement{
		Element: &Element{
			Tag:           elementTagTIME,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *TimeHTMLElement) Children(children ...ElementBuilder) *TimeHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *TimeHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TimeHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *TimeHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TimeHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *TimeHTMLElement) Text(text string) *TimeHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *TimeHTMLElement) TextF(format string, args ...any) *TimeHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *TimeHTMLElement) Escaped(text string) *TimeHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *TimeHTMLElement) EscapedF(format string, args ...any) *TimeHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TimeHTMLElement) CustomData(key, value string) *TimeHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TimeHTMLElement) CustomDataRemove(key string) *TimeHTMLElement {
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
func (e *TimeHTMLElement) ACCESSKEY(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *TimeHTMLElement) IfACCESSKEY(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *TimeHTMLElement) RemoveACCESSKEY(v string) *TimeHTMLElement {
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
func (e *TimeHTMLElement) AUTOCAPITALIZE(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *TimeHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *TimeHTMLElement) RemoveAUTOCAPITALIZE(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *TimeHTMLElement) AUTOFOCUS() *TimeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *TimeHTMLElement) IfAUTOFOCUS(cond bool) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *TimeHTMLElement) RemoveAUTOFOCUS() *TimeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *TimeHTMLElement) SetAUTOFOCUS(b bool) *TimeHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *TimeHTMLElement) CLASS(v string) *TimeHTMLElement {
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

func (e *TimeHTMLElement) IfCLASS(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *TimeHTMLElement) SetCLASS(v string) *TimeHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *TimeHTMLElement) RemoveCLASS(v string) *TimeHTMLElement {
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
func (e *TimeHTMLElement) CONTENTEDITABLE(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *TimeHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *TimeHTMLElement) RemoveCONTENTEDITABLE(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DATETIME sets the "datetime" attribute.
// Machine-readable value
// Values values are constrained to:
//   - valid_date_string
//   - valid_duration_string
//   - valid_global_date_and_time_string
//   - valid_local_date_and_time_string
//   - valid_month_string
//   - valid_non_negative_integer
//   - valid_time_string
//   - valid_time_zone_offset_string
//   - valid_week_string
//   - valid_yearless_date_string
func (e *TimeHTMLElement) DATETIME(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDATETIMEKey] = v
	return e
}

func (e *TimeHTMLElement) IfDATETIME(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.DATETIME(v)
}

func (e *TimeHTMLElement) RemoveDATETIME(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeDATETIMEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *TimeHTMLElement) DIR(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *TimeHTMLElement) IfDIR(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *TimeHTMLElement) RemoveDIR(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *TimeHTMLElement) DRAGGABLE(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *TimeHTMLElement) IfDRAGGABLE(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *TimeHTMLElement) RemoveDRAGGABLE(v string) *TimeHTMLElement {
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
func (e *TimeHTMLElement) ENTERKEYHINT(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *TimeHTMLElement) IfENTERKEYHINT(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *TimeHTMLElement) RemoveENTERKEYHINT(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *TimeHTMLElement) HIDDEN(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *TimeHTMLElement) IfHIDDEN(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *TimeHTMLElement) RemoveHIDDEN(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *TimeHTMLElement) ID(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *TimeHTMLElement) IfID(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *TimeHTMLElement) RemoveID(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *TimeHTMLElement) INERT() *TimeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *TimeHTMLElement) IfINERT(cond bool) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *TimeHTMLElement) RemoveINERT() *TimeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *TimeHTMLElement) SetINERT(b bool) *TimeHTMLElement {
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
func (e *TimeHTMLElement) INPUTMODE(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *TimeHTMLElement) IfINPUTMODE(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *TimeHTMLElement) RemoveINPUTMODE(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *TimeHTMLElement) IS(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *TimeHTMLElement) IfIS(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *TimeHTMLElement) RemoveIS(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *TimeHTMLElement) ITEMID(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *TimeHTMLElement) IfITEMID(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *TimeHTMLElement) RemoveITEMID(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *TimeHTMLElement) ITEMPROP(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *TimeHTMLElement) IfITEMPROP(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *TimeHTMLElement) RemoveITEMPROP(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TimeHTMLElement) ITEMREF(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *TimeHTMLElement) IfITEMREF(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *TimeHTMLElement) RemoveITEMREF(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *TimeHTMLElement) ITEMSCOPE() *TimeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *TimeHTMLElement) IfITEMSCOPE(cond bool) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *TimeHTMLElement) RemoveITEMSCOPE() *TimeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *TimeHTMLElement) SetITEMSCOPE(b bool) *TimeHTMLElement {
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
func (e *TimeHTMLElement) ITEMTYPE(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *TimeHTMLElement) IfITEMTYPE(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *TimeHTMLElement) RemoveITEMTYPE(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TimeHTMLElement) LANG(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *TimeHTMLElement) IfLANG(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *TimeHTMLElement) RemoveLANG(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *TimeHTMLElement) NONCE(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *TimeHTMLElement) IfNONCE(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *TimeHTMLElement) RemoveNONCE(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *TimeHTMLElement) POPOVER(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *TimeHTMLElement) IfPOPOVER(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *TimeHTMLElement) RemovePOPOVER(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *TimeHTMLElement) SLOT(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *TimeHTMLElement) IfSLOT(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *TimeHTMLElement) RemoveSLOT(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *TimeHTMLElement) SPELLCHECK(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *TimeHTMLElement) IfSPELLCHECK(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *TimeHTMLElement) RemoveSPELLCHECK(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TimeHTMLElement) STYLE(k, v string) *TimeHTMLElement {
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

func (e *TimeHTMLElement) IfSTYLE(cond bool, k string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *TimeHTMLElement) RemoveSTYLE(k string) *TimeHTMLElement {
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
func (e *TimeHTMLElement) TABINDEX(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *TimeHTMLElement) IfTABINDEX(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *TimeHTMLElement) RemoveTABINDEX(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *TimeHTMLElement) TITLE(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *TimeHTMLElement) IfTITLE(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *TimeHTMLElement) RemoveTITLE(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *TimeHTMLElement) TRANSLATE(v string) *TimeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *TimeHTMLElement) IfTRANSLATE(cond bool, v string) *TimeHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *TimeHTMLElement) RemoveTRANSLATE(v string) *TimeHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
