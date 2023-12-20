package html

import (
	"fmt"
)

type DelHTMLElement struct {
	*Element
}

func DEL(children ...ElementBuilder) *DelHTMLElement {
	return &DelHTMLElement{
		Element: &Element{
			Tag:           elementTagDEL,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DelHTMLElement) Children(children ...ElementBuilder) *DelHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DelHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DelHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DelHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DelHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DelHTMLElement) Text(text string) *DelHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DelHTMLElement) TextF(format string, args ...any) *DelHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DelHTMLElement) Escaped(text string) *DelHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DelHTMLElement) EscapedF(format string, args ...any) *DelHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DelHTMLElement) CustomData(key, value string) *DelHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DelHTMLElement) CustomDataRemove(key string) *DelHTMLElement {
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
func (e *DelHTMLElement) ACCESSKEY(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DelHTMLElement) IfACCESSKEY(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DelHTMLElement) RemoveACCESSKEY(v string) *DelHTMLElement {
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
func (e *DelHTMLElement) AUTOCAPITALIZE(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DelHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DelHTMLElement) RemoveAUTOCAPITALIZE(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DelHTMLElement) AUTOFOCUS() *DelHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DelHTMLElement) IfAUTOFOCUS(cond bool) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DelHTMLElement) RemoveAUTOFOCUS() *DelHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DelHTMLElement) SetAUTOFOCUS(b bool) *DelHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CITE sets the "cite" attribute.
// Link to the source of the quotation or more information about the edit
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DelHTMLElement) CITE(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCITEKey] = v
	return e
}

func (e *DelHTMLElement) IfCITE(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.CITE(v)
}

func (e *DelHTMLElement) RemoveCITE(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeCITEKey)
	return e
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DelHTMLElement) CLASS(v string) *DelHTMLElement {
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

func (e *DelHTMLElement) IfCLASS(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DelHTMLElement) SetCLASS(v string) *DelHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DelHTMLElement) RemoveCLASS(v string) *DelHTMLElement {
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
func (e *DelHTMLElement) CONTENTEDITABLE(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DelHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DelHTMLElement) RemoveCONTENTEDITABLE(v string) *DelHTMLElement {
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
func (e *DelHTMLElement) DATETIME(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDATETIMEKey] = v
	return e
}

func (e *DelHTMLElement) IfDATETIME(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.DATETIME(v)
}

func (e *DelHTMLElement) RemoveDATETIME(v string) *DelHTMLElement {
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
func (e *DelHTMLElement) DIR(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DelHTMLElement) IfDIR(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DelHTMLElement) RemoveDIR(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *DelHTMLElement) DRAGGABLE(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DelHTMLElement) IfDRAGGABLE(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DelHTMLElement) RemoveDRAGGABLE(v string) *DelHTMLElement {
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
func (e *DelHTMLElement) ENTERKEYHINT(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DelHTMLElement) IfENTERKEYHINT(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DelHTMLElement) RemoveENTERKEYHINT(v string) *DelHTMLElement {
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
func (e *DelHTMLElement) HIDDEN(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DelHTMLElement) IfHIDDEN(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DelHTMLElement) RemoveHIDDEN(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DelHTMLElement) ID(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DelHTMLElement) IfID(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DelHTMLElement) RemoveID(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DelHTMLElement) INERT() *DelHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DelHTMLElement) IfINERT(cond bool) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DelHTMLElement) RemoveINERT() *DelHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DelHTMLElement) SetINERT(b bool) *DelHTMLElement {
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
func (e *DelHTMLElement) INPUTMODE(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DelHTMLElement) IfINPUTMODE(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DelHTMLElement) RemoveINPUTMODE(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *DelHTMLElement) IS(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DelHTMLElement) IfIS(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DelHTMLElement) RemoveIS(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DelHTMLElement) ITEMID(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DelHTMLElement) IfITEMID(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DelHTMLElement) RemoveITEMID(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *DelHTMLElement) ITEMPROP(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DelHTMLElement) IfITEMPROP(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DelHTMLElement) RemoveITEMPROP(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DelHTMLElement) ITEMREF(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DelHTMLElement) IfITEMREF(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DelHTMLElement) RemoveITEMREF(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DelHTMLElement) ITEMSCOPE() *DelHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DelHTMLElement) IfITEMSCOPE(cond bool) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DelHTMLElement) RemoveITEMSCOPE() *DelHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DelHTMLElement) SetITEMSCOPE(b bool) *DelHTMLElement {
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
func (e *DelHTMLElement) ITEMTYPE(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DelHTMLElement) IfITEMTYPE(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DelHTMLElement) RemoveITEMTYPE(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DelHTMLElement) LANG(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DelHTMLElement) IfLANG(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DelHTMLElement) RemoveLANG(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DelHTMLElement) NONCE(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DelHTMLElement) IfNONCE(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DelHTMLElement) RemoveNONCE(v string) *DelHTMLElement {
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
func (e *DelHTMLElement) POPOVER(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DelHTMLElement) IfPOPOVER(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DelHTMLElement) RemovePOPOVER(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DelHTMLElement) SLOT(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DelHTMLElement) IfSLOT(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DelHTMLElement) RemoveSLOT(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *DelHTMLElement) SPELLCHECK(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DelHTMLElement) IfSPELLCHECK(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DelHTMLElement) RemoveSPELLCHECK(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DelHTMLElement) STYLE(k, v string) *DelHTMLElement {
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

func (e *DelHTMLElement) IfSTYLE(cond bool, k string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DelHTMLElement) RemoveSTYLE(k string) *DelHTMLElement {
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
func (e *DelHTMLElement) TABINDEX(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DelHTMLElement) IfTABINDEX(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DelHTMLElement) RemoveTABINDEX(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DelHTMLElement) TITLE(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DelHTMLElement) IfTITLE(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DelHTMLElement) RemoveTITLE(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *DelHTMLElement) TRANSLATE(v string) *DelHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DelHTMLElement) IfTRANSLATE(cond bool, v string) *DelHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DelHTMLElement) RemoveTRANSLATE(v string) *DelHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
