package html

import (
	"fmt"
)

type InsHTMLElement struct {
	*Element
}

func INS(children ...ElementBuilder) *InsHTMLElement {
	return &InsHTMLElement{
		Element: &Element{
			Tag:           elementTagINS,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *InsHTMLElement) Children(children ...ElementBuilder) *InsHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *InsHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *InsHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *InsHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *InsHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *InsHTMLElement) Text(text string) *InsHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *InsHTMLElement) TextF(format string, args ...any) *InsHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *InsHTMLElement) Escaped(text string) *InsHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *InsHTMLElement) EscapedF(format string, args ...any) *InsHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *InsHTMLElement) CustomData(key, value string) *InsHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *InsHTMLElement) CustomDataRemove(key string) *InsHTMLElement {
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
func (e *InsHTMLElement) ACCESSKEY(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *InsHTMLElement) IfACCESSKEY(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *InsHTMLElement) RemoveACCESSKEY(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) AUTOCAPITALIZE(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *InsHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *InsHTMLElement) RemoveAUTOCAPITALIZE(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *InsHTMLElement) AUTOFOCUS() *InsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *InsHTMLElement) IfAUTOFOCUS(cond bool) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *InsHTMLElement) RemoveAUTOFOCUS() *InsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *InsHTMLElement) SetAUTOFOCUS(b bool) *InsHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CITE sets the "cite" attribute.
// Link to the source of the quotation or more information about the edit
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *InsHTMLElement) CITE(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCITEKey] = v
	return e
}

func (e *InsHTMLElement) IfCITE(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.CITE(v)
}

func (e *InsHTMLElement) RemoveCITE(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeCITEKey)
	return e
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *InsHTMLElement) CLASS(v string) *InsHTMLElement {
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

func (e *InsHTMLElement) IfCLASS(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *InsHTMLElement) SetCLASS(v string) *InsHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *InsHTMLElement) RemoveCLASS(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) CONTENTEDITABLE(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *InsHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *InsHTMLElement) RemoveCONTENTEDITABLE(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DATETIME sets the "datetime" attribute.
// Machine-readable value
// Values values are constrained to:
//   - valid_month_string
//   - valid_date_string
//   - valid_yearless_date_string
//   - valid_time_string
//   - valid_local_date_and_time_string
//   - valid_time_zone_offset_string
//   - valid_global_date_and_time_string
//   - valid_week_string
//   - valid_non_negative_integer
//   - valid_duration_string
func (e *InsHTMLElement) DATETIME(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDATETIMEKey] = v
	return e
}

func (e *InsHTMLElement) IfDATETIME(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.DATETIME(v)
}

func (e *InsHTMLElement) RemoveDATETIME(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeDATETIMEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (e *InsHTMLElement) DIR(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *InsHTMLElement) IfDIR(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *InsHTMLElement) RemoveDIR(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *InsHTMLElement) DRAGGABLE(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *InsHTMLElement) IfDRAGGABLE(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *InsHTMLElement) RemoveDRAGGABLE(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) ENTERKEYHINT(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *InsHTMLElement) IfENTERKEYHINT(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *InsHTMLElement) RemoveENTERKEYHINT(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) HIDDEN(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *InsHTMLElement) IfHIDDEN(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *InsHTMLElement) RemoveHIDDEN(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *InsHTMLElement) ID(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *InsHTMLElement) IfID(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *InsHTMLElement) RemoveID(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *InsHTMLElement) INERT() *InsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *InsHTMLElement) IfINERT(cond bool) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *InsHTMLElement) RemoveINERT() *InsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *InsHTMLElement) SetINERT(b bool) *InsHTMLElement {
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
func (e *InsHTMLElement) INPUTMODE(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *InsHTMLElement) IfINPUTMODE(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *InsHTMLElement) RemoveINPUTMODE(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *InsHTMLElement) IS(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *InsHTMLElement) IfIS(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *InsHTMLElement) RemoveIS(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *InsHTMLElement) ITEMID(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *InsHTMLElement) IfITEMID(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *InsHTMLElement) RemoveITEMID(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *InsHTMLElement) ITEMPROP(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *InsHTMLElement) IfITEMPROP(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *InsHTMLElement) RemoveITEMPROP(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *InsHTMLElement) ITEMREF(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *InsHTMLElement) IfITEMREF(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *InsHTMLElement) RemoveITEMREF(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *InsHTMLElement) ITEMSCOPE() *InsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *InsHTMLElement) IfITEMSCOPE(cond bool) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *InsHTMLElement) RemoveITEMSCOPE() *InsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *InsHTMLElement) SetITEMSCOPE(b bool) *InsHTMLElement {
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
func (e *InsHTMLElement) ITEMTYPE(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *InsHTMLElement) IfITEMTYPE(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *InsHTMLElement) RemoveITEMTYPE(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *InsHTMLElement) LANG(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *InsHTMLElement) IfLANG(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *InsHTMLElement) RemoveLANG(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *InsHTMLElement) NONCE(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *InsHTMLElement) IfNONCE(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *InsHTMLElement) RemoveNONCE(v string) *InsHTMLElement {
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
func (e *InsHTMLElement) POPOVER(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *InsHTMLElement) IfPOPOVER(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *InsHTMLElement) RemovePOPOVER(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *InsHTMLElement) SLOT(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *InsHTMLElement) IfSLOT(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *InsHTMLElement) RemoveSLOT(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *InsHTMLElement) SPELLCHECK(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *InsHTMLElement) IfSPELLCHECK(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *InsHTMLElement) RemoveSPELLCHECK(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *InsHTMLElement) STYLE(k, v string) *InsHTMLElement {
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

func (e *InsHTMLElement) IfSTYLE(cond bool, k string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *InsHTMLElement) RemoveSTYLE(k string) *InsHTMLElement {
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
func (e *InsHTMLElement) TABINDEX(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *InsHTMLElement) IfTABINDEX(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *InsHTMLElement) RemoveTABINDEX(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *InsHTMLElement) TITLE(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *InsHTMLElement) IfTITLE(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *InsHTMLElement) RemoveTITLE(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *InsHTMLElement) TRANSLATE(v string) *InsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *InsHTMLElement) IfTRANSLATE(cond bool, v string) *InsHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *InsHTMLElement) RemoveTRANSLATE(v string) *InsHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
