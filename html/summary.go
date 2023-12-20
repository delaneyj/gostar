package html

import (
	"fmt"
)

type SummaryHTMLElement struct {
	*Element
}

func SUMMARY(children ...ElementBuilder) *SummaryHTMLElement {
	return &SummaryHTMLElement{
		Element: &Element{
			Tag:           elementTagSUMMARY,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SummaryHTMLElement) Children(children ...ElementBuilder) *SummaryHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SummaryHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SummaryHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SummaryHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SummaryHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SummaryHTMLElement) Text(text string) *SummaryHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SummaryHTMLElement) TextF(format string, args ...any) *SummaryHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SummaryHTMLElement) Escaped(text string) *SummaryHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SummaryHTMLElement) EscapedF(format string, args ...any) *SummaryHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SummaryHTMLElement) CustomData(key, value string) *SummaryHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SummaryHTMLElement) CustomDataRemove(key string) *SummaryHTMLElement {
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
func (e *SummaryHTMLElement) ACCESSKEY(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SummaryHTMLElement) IfACCESSKEY(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SummaryHTMLElement) RemoveACCESSKEY(v string) *SummaryHTMLElement {
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
func (e *SummaryHTMLElement) AUTOCAPITALIZE(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SummaryHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SummaryHTMLElement) RemoveAUTOCAPITALIZE(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SummaryHTMLElement) AUTOFOCUS() *SummaryHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SummaryHTMLElement) IfAUTOFOCUS(cond bool) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SummaryHTMLElement) RemoveAUTOFOCUS() *SummaryHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SummaryHTMLElement) SetAUTOFOCUS(b bool) *SummaryHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SummaryHTMLElement) CLASS(v string) *SummaryHTMLElement {
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

func (e *SummaryHTMLElement) IfCLASS(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SummaryHTMLElement) SetCLASS(v string) *SummaryHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SummaryHTMLElement) RemoveCLASS(v string) *SummaryHTMLElement {
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
func (e *SummaryHTMLElement) CONTENTEDITABLE(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SummaryHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SummaryHTMLElement) RemoveCONTENTEDITABLE(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *SummaryHTMLElement) DIR(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SummaryHTMLElement) IfDIR(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SummaryHTMLElement) RemoveDIR(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *SummaryHTMLElement) DRAGGABLE(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SummaryHTMLElement) IfDRAGGABLE(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SummaryHTMLElement) RemoveDRAGGABLE(v string) *SummaryHTMLElement {
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
func (e *SummaryHTMLElement) ENTERKEYHINT(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SummaryHTMLElement) IfENTERKEYHINT(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SummaryHTMLElement) RemoveENTERKEYHINT(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *SummaryHTMLElement) HIDDEN(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SummaryHTMLElement) IfHIDDEN(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SummaryHTMLElement) RemoveHIDDEN(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SummaryHTMLElement) ID(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SummaryHTMLElement) IfID(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SummaryHTMLElement) RemoveID(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SummaryHTMLElement) INERT() *SummaryHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SummaryHTMLElement) IfINERT(cond bool) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SummaryHTMLElement) RemoveINERT() *SummaryHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SummaryHTMLElement) SetINERT(b bool) *SummaryHTMLElement {
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
func (e *SummaryHTMLElement) INPUTMODE(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SummaryHTMLElement) IfINPUTMODE(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SummaryHTMLElement) RemoveINPUTMODE(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *SummaryHTMLElement) IS(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SummaryHTMLElement) IfIS(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SummaryHTMLElement) RemoveIS(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SummaryHTMLElement) ITEMID(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SummaryHTMLElement) IfITEMID(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SummaryHTMLElement) RemoveITEMID(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *SummaryHTMLElement) ITEMPROP(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SummaryHTMLElement) IfITEMPROP(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SummaryHTMLElement) RemoveITEMPROP(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SummaryHTMLElement) ITEMREF(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SummaryHTMLElement) IfITEMREF(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SummaryHTMLElement) RemoveITEMREF(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SummaryHTMLElement) ITEMSCOPE() *SummaryHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SummaryHTMLElement) IfITEMSCOPE(cond bool) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SummaryHTMLElement) RemoveITEMSCOPE() *SummaryHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SummaryHTMLElement) SetITEMSCOPE(b bool) *SummaryHTMLElement {
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
func (e *SummaryHTMLElement) ITEMTYPE(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SummaryHTMLElement) IfITEMTYPE(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SummaryHTMLElement) RemoveITEMTYPE(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SummaryHTMLElement) LANG(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SummaryHTMLElement) IfLANG(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SummaryHTMLElement) RemoveLANG(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SummaryHTMLElement) NONCE(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SummaryHTMLElement) IfNONCE(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SummaryHTMLElement) RemoveNONCE(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *SummaryHTMLElement) POPOVER(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SummaryHTMLElement) IfPOPOVER(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SummaryHTMLElement) RemovePOPOVER(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SummaryHTMLElement) SLOT(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SummaryHTMLElement) IfSLOT(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SummaryHTMLElement) RemoveSLOT(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *SummaryHTMLElement) SPELLCHECK(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SummaryHTMLElement) IfSPELLCHECK(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SummaryHTMLElement) RemoveSPELLCHECK(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SummaryHTMLElement) STYLE(k, v string) *SummaryHTMLElement {
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

func (e *SummaryHTMLElement) IfSTYLE(cond bool, k string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SummaryHTMLElement) RemoveSTYLE(k string) *SummaryHTMLElement {
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
func (e *SummaryHTMLElement) TABINDEX(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SummaryHTMLElement) IfTABINDEX(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SummaryHTMLElement) RemoveTABINDEX(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SummaryHTMLElement) TITLE(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SummaryHTMLElement) IfTITLE(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SummaryHTMLElement) RemoveTITLE(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *SummaryHTMLElement) TRANSLATE(v string) *SummaryHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SummaryHTMLElement) IfTRANSLATE(cond bool, v string) *SummaryHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SummaryHTMLElement) RemoveTRANSLATE(v string) *SummaryHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
