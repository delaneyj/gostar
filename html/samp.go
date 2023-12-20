package html

import (
	"fmt"
)

type SampHTMLElement struct {
	*Element
}

func SAMP(children ...ElementBuilder) *SampHTMLElement {
	return &SampHTMLElement{
		Element: &Element{
			Tag:           elementTagSAMP,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SampHTMLElement) Children(children ...ElementBuilder) *SampHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SampHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SampHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SampHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SampHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SampHTMLElement) Text(text string) *SampHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SampHTMLElement) TextF(format string, args ...any) *SampHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SampHTMLElement) Escaped(text string) *SampHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SampHTMLElement) EscapedF(format string, args ...any) *SampHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SampHTMLElement) CustomData(key, value string) *SampHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SampHTMLElement) CustomDataRemove(key string) *SampHTMLElement {
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
func (e *SampHTMLElement) ACCESSKEY(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SampHTMLElement) IfACCESSKEY(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SampHTMLElement) RemoveACCESSKEY(v string) *SampHTMLElement {
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
func (e *SampHTMLElement) AUTOCAPITALIZE(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SampHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SampHTMLElement) RemoveAUTOCAPITALIZE(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SampHTMLElement) AUTOFOCUS() *SampHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SampHTMLElement) IfAUTOFOCUS(cond bool) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SampHTMLElement) RemoveAUTOFOCUS() *SampHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SampHTMLElement) SetAUTOFOCUS(b bool) *SampHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SampHTMLElement) CLASS(v string) *SampHTMLElement {
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

func (e *SampHTMLElement) IfCLASS(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SampHTMLElement) SetCLASS(v string) *SampHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SampHTMLElement) RemoveCLASS(v string) *SampHTMLElement {
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
func (e *SampHTMLElement) CONTENTEDITABLE(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SampHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SampHTMLElement) RemoveCONTENTEDITABLE(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *SampHTMLElement) DIR(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SampHTMLElement) IfDIR(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SampHTMLElement) RemoveDIR(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *SampHTMLElement) DRAGGABLE(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SampHTMLElement) IfDRAGGABLE(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SampHTMLElement) RemoveDRAGGABLE(v string) *SampHTMLElement {
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
func (e *SampHTMLElement) ENTERKEYHINT(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SampHTMLElement) IfENTERKEYHINT(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SampHTMLElement) RemoveENTERKEYHINT(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *SampHTMLElement) HIDDEN(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SampHTMLElement) IfHIDDEN(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SampHTMLElement) RemoveHIDDEN(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SampHTMLElement) ID(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SampHTMLElement) IfID(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SampHTMLElement) RemoveID(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SampHTMLElement) INERT() *SampHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SampHTMLElement) IfINERT(cond bool) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SampHTMLElement) RemoveINERT() *SampHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SampHTMLElement) SetINERT(b bool) *SampHTMLElement {
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
func (e *SampHTMLElement) INPUTMODE(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SampHTMLElement) IfINPUTMODE(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SampHTMLElement) RemoveINPUTMODE(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *SampHTMLElement) IS(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SampHTMLElement) IfIS(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SampHTMLElement) RemoveIS(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SampHTMLElement) ITEMID(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SampHTMLElement) IfITEMID(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SampHTMLElement) RemoveITEMID(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *SampHTMLElement) ITEMPROP(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SampHTMLElement) IfITEMPROP(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SampHTMLElement) RemoveITEMPROP(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SampHTMLElement) ITEMREF(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SampHTMLElement) IfITEMREF(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SampHTMLElement) RemoveITEMREF(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SampHTMLElement) ITEMSCOPE() *SampHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SampHTMLElement) IfITEMSCOPE(cond bool) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SampHTMLElement) RemoveITEMSCOPE() *SampHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SampHTMLElement) SetITEMSCOPE(b bool) *SampHTMLElement {
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
func (e *SampHTMLElement) ITEMTYPE(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SampHTMLElement) IfITEMTYPE(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SampHTMLElement) RemoveITEMTYPE(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SampHTMLElement) LANG(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SampHTMLElement) IfLANG(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SampHTMLElement) RemoveLANG(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SampHTMLElement) NONCE(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SampHTMLElement) IfNONCE(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SampHTMLElement) RemoveNONCE(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *SampHTMLElement) POPOVER(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SampHTMLElement) IfPOPOVER(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SampHTMLElement) RemovePOPOVER(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SampHTMLElement) SLOT(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SampHTMLElement) IfSLOT(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SampHTMLElement) RemoveSLOT(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *SampHTMLElement) SPELLCHECK(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SampHTMLElement) IfSPELLCHECK(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SampHTMLElement) RemoveSPELLCHECK(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SampHTMLElement) STYLE(k, v string) *SampHTMLElement {
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

func (e *SampHTMLElement) IfSTYLE(cond bool, k string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SampHTMLElement) RemoveSTYLE(k string) *SampHTMLElement {
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
func (e *SampHTMLElement) TABINDEX(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SampHTMLElement) IfTABINDEX(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SampHTMLElement) RemoveTABINDEX(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SampHTMLElement) TITLE(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SampHTMLElement) IfTITLE(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SampHTMLElement) RemoveTITLE(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *SampHTMLElement) TRANSLATE(v string) *SampHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SampHTMLElement) IfTRANSLATE(cond bool, v string) *SampHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SampHTMLElement) RemoveTRANSLATE(v string) *SampHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
