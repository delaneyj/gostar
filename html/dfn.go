package html

import (
	"fmt"
)

type DfnHTMLElement struct {
	*Element
}

func DFN(children ...ElementBuilder) *DfnHTMLElement {
	return &DfnHTMLElement{
		Element: &Element{
			Tag:           elementTagDFN,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DfnHTMLElement) Children(children ...ElementBuilder) *DfnHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DfnHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DfnHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DfnHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DfnHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DfnHTMLElement) Text(text string) *DfnHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DfnHTMLElement) TextF(format string, args ...any) *DfnHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DfnHTMLElement) Escaped(text string) *DfnHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DfnHTMLElement) EscapedF(format string, args ...any) *DfnHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DfnHTMLElement) CustomData(key, value string) *DfnHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DfnHTMLElement) CustomDataRemove(key string) *DfnHTMLElement {
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
func (e *DfnHTMLElement) ACCESSKEY(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DfnHTMLElement) IfACCESSKEY(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DfnHTMLElement) RemoveACCESSKEY(v string) *DfnHTMLElement {
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
func (e *DfnHTMLElement) AUTOCAPITALIZE(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DfnHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DfnHTMLElement) RemoveAUTOCAPITALIZE(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DfnHTMLElement) AUTOFOCUS() *DfnHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DfnHTMLElement) IfAUTOFOCUS(cond bool) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DfnHTMLElement) RemoveAUTOFOCUS() *DfnHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DfnHTMLElement) SetAUTOFOCUS(b bool) *DfnHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DfnHTMLElement) CLASS(v string) *DfnHTMLElement {
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

func (e *DfnHTMLElement) IfCLASS(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DfnHTMLElement) SetCLASS(v string) *DfnHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DfnHTMLElement) RemoveCLASS(v string) *DfnHTMLElement {
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
func (e *DfnHTMLElement) CONTENTEDITABLE(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DfnHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DfnHTMLElement) RemoveCONTENTEDITABLE(v string) *DfnHTMLElement {
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
func (e *DfnHTMLElement) DIR(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DfnHTMLElement) IfDIR(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DfnHTMLElement) RemoveDIR(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *DfnHTMLElement) DRAGGABLE(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DfnHTMLElement) IfDRAGGABLE(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DfnHTMLElement) RemoveDRAGGABLE(v string) *DfnHTMLElement {
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
func (e *DfnHTMLElement) ENTERKEYHINT(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DfnHTMLElement) IfENTERKEYHINT(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DfnHTMLElement) RemoveENTERKEYHINT(v string) *DfnHTMLElement {
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
func (e *DfnHTMLElement) HIDDEN(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DfnHTMLElement) IfHIDDEN(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DfnHTMLElement) RemoveHIDDEN(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DfnHTMLElement) ID(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DfnHTMLElement) IfID(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DfnHTMLElement) RemoveID(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DfnHTMLElement) INERT() *DfnHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DfnHTMLElement) IfINERT(cond bool) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DfnHTMLElement) RemoveINERT() *DfnHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DfnHTMLElement) SetINERT(b bool) *DfnHTMLElement {
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
func (e *DfnHTMLElement) INPUTMODE(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DfnHTMLElement) IfINPUTMODE(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DfnHTMLElement) RemoveINPUTMODE(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *DfnHTMLElement) IS(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DfnHTMLElement) IfIS(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DfnHTMLElement) RemoveIS(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DfnHTMLElement) ITEMID(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DfnHTMLElement) IfITEMID(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DfnHTMLElement) RemoveITEMID(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *DfnHTMLElement) ITEMPROP(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DfnHTMLElement) IfITEMPROP(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DfnHTMLElement) RemoveITEMPROP(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DfnHTMLElement) ITEMREF(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DfnHTMLElement) IfITEMREF(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DfnHTMLElement) RemoveITEMREF(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DfnHTMLElement) ITEMSCOPE() *DfnHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DfnHTMLElement) IfITEMSCOPE(cond bool) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DfnHTMLElement) RemoveITEMSCOPE() *DfnHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DfnHTMLElement) SetITEMSCOPE(b bool) *DfnHTMLElement {
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
func (e *DfnHTMLElement) ITEMTYPE(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DfnHTMLElement) IfITEMTYPE(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DfnHTMLElement) RemoveITEMTYPE(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DfnHTMLElement) LANG(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DfnHTMLElement) IfLANG(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DfnHTMLElement) RemoveLANG(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DfnHTMLElement) NONCE(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DfnHTMLElement) IfNONCE(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DfnHTMLElement) RemoveNONCE(v string) *DfnHTMLElement {
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
func (e *DfnHTMLElement) POPOVER(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DfnHTMLElement) IfPOPOVER(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DfnHTMLElement) RemovePOPOVER(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DfnHTMLElement) SLOT(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DfnHTMLElement) IfSLOT(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DfnHTMLElement) RemoveSLOT(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *DfnHTMLElement) SPELLCHECK(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DfnHTMLElement) IfSPELLCHECK(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DfnHTMLElement) RemoveSPELLCHECK(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DfnHTMLElement) STYLE(k, v string) *DfnHTMLElement {
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

func (e *DfnHTMLElement) IfSTYLE(cond bool, k string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DfnHTMLElement) RemoveSTYLE(k string) *DfnHTMLElement {
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
func (e *DfnHTMLElement) TABINDEX(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DfnHTMLElement) IfTABINDEX(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DfnHTMLElement) RemoveTABINDEX(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DfnHTMLElement) TITLE(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DfnHTMLElement) IfTITLE(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DfnHTMLElement) RemoveTITLE(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *DfnHTMLElement) TRANSLATE(v string) *DfnHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DfnHTMLElement) IfTRANSLATE(cond bool, v string) *DfnHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DfnHTMLElement) RemoveTRANSLATE(v string) *DfnHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
