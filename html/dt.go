package html

import (
	"fmt"
)

type DtHTMLElement struct {
	*Element
}

func DT(children ...ElementBuilder) *DtHTMLElement {
	return &DtHTMLElement{
		Element: &Element{
			Tag:           elementTagDT,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DtHTMLElement) Children(children ...ElementBuilder) *DtHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DtHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DtHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DtHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DtHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DtHTMLElement) Text(text string) *DtHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DtHTMLElement) TextF(format string, args ...any) *DtHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DtHTMLElement) Escaped(text string) *DtHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DtHTMLElement) EscapedF(format string, args ...any) *DtHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DtHTMLElement) CustomData(key, value string) *DtHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DtHTMLElement) CustomDataRemove(key string) *DtHTMLElement {
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
func (e *DtHTMLElement) ACCESSKEY(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DtHTMLElement) IfACCESSKEY(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DtHTMLElement) RemoveACCESSKEY(v string) *DtHTMLElement {
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
func (e *DtHTMLElement) AUTOCAPITALIZE(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DtHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DtHTMLElement) RemoveAUTOCAPITALIZE(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DtHTMLElement) AUTOFOCUS() *DtHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DtHTMLElement) IfAUTOFOCUS(cond bool) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DtHTMLElement) RemoveAUTOFOCUS() *DtHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DtHTMLElement) SetAUTOFOCUS(b bool) *DtHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DtHTMLElement) CLASS(v string) *DtHTMLElement {
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

func (e *DtHTMLElement) IfCLASS(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DtHTMLElement) SetCLASS(v string) *DtHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DtHTMLElement) RemoveCLASS(v string) *DtHTMLElement {
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
func (e *DtHTMLElement) CONTENTEDITABLE(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DtHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DtHTMLElement) RemoveCONTENTEDITABLE(v string) *DtHTMLElement {
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
func (e *DtHTMLElement) DIR(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DtHTMLElement) IfDIR(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DtHTMLElement) RemoveDIR(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *DtHTMLElement) DRAGGABLE(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DtHTMLElement) IfDRAGGABLE(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DtHTMLElement) RemoveDRAGGABLE(v string) *DtHTMLElement {
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
func (e *DtHTMLElement) ENTERKEYHINT(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DtHTMLElement) IfENTERKEYHINT(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DtHTMLElement) RemoveENTERKEYHINT(v string) *DtHTMLElement {
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
func (e *DtHTMLElement) HIDDEN(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DtHTMLElement) IfHIDDEN(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DtHTMLElement) RemoveHIDDEN(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DtHTMLElement) ID(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DtHTMLElement) IfID(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DtHTMLElement) RemoveID(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DtHTMLElement) INERT() *DtHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DtHTMLElement) IfINERT(cond bool) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DtHTMLElement) RemoveINERT() *DtHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DtHTMLElement) SetINERT(b bool) *DtHTMLElement {
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
func (e *DtHTMLElement) INPUTMODE(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DtHTMLElement) IfINPUTMODE(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DtHTMLElement) RemoveINPUTMODE(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *DtHTMLElement) IS(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DtHTMLElement) IfIS(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DtHTMLElement) RemoveIS(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DtHTMLElement) ITEMID(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DtHTMLElement) IfITEMID(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DtHTMLElement) RemoveITEMID(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *DtHTMLElement) ITEMPROP(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DtHTMLElement) IfITEMPROP(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DtHTMLElement) RemoveITEMPROP(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DtHTMLElement) ITEMREF(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DtHTMLElement) IfITEMREF(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DtHTMLElement) RemoveITEMREF(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DtHTMLElement) ITEMSCOPE() *DtHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DtHTMLElement) IfITEMSCOPE(cond bool) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DtHTMLElement) RemoveITEMSCOPE() *DtHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DtHTMLElement) SetITEMSCOPE(b bool) *DtHTMLElement {
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
func (e *DtHTMLElement) ITEMTYPE(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DtHTMLElement) IfITEMTYPE(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DtHTMLElement) RemoveITEMTYPE(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DtHTMLElement) LANG(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DtHTMLElement) IfLANG(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DtHTMLElement) RemoveLANG(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DtHTMLElement) NONCE(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DtHTMLElement) IfNONCE(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DtHTMLElement) RemoveNONCE(v string) *DtHTMLElement {
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
func (e *DtHTMLElement) POPOVER(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DtHTMLElement) IfPOPOVER(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DtHTMLElement) RemovePOPOVER(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DtHTMLElement) SLOT(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DtHTMLElement) IfSLOT(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DtHTMLElement) RemoveSLOT(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *DtHTMLElement) SPELLCHECK(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DtHTMLElement) IfSPELLCHECK(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DtHTMLElement) RemoveSPELLCHECK(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DtHTMLElement) STYLE(k, v string) *DtHTMLElement {
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

func (e *DtHTMLElement) IfSTYLE(cond bool, k string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DtHTMLElement) RemoveSTYLE(k string) *DtHTMLElement {
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
func (e *DtHTMLElement) TABINDEX(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DtHTMLElement) IfTABINDEX(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DtHTMLElement) RemoveTABINDEX(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DtHTMLElement) TITLE(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DtHTMLElement) IfTITLE(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DtHTMLElement) RemoveTITLE(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *DtHTMLElement) TRANSLATE(v string) *DtHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DtHTMLElement) IfTRANSLATE(cond bool, v string) *DtHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DtHTMLElement) RemoveTRANSLATE(v string) *DtHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
