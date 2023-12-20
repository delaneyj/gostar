package html

import (
	"fmt"
)

type DdHTMLElement struct {
	*Element
}

func DD(children ...ElementBuilder) *DdHTMLElement {
	return &DdHTMLElement{
		Element: &Element{
			Tag:           elementTagDD,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DdHTMLElement) Children(children ...ElementBuilder) *DdHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DdHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DdHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DdHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DdHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DdHTMLElement) Text(text string) *DdHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DdHTMLElement) TextF(format string, args ...any) *DdHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DdHTMLElement) Escaped(text string) *DdHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DdHTMLElement) EscapedF(format string, args ...any) *DdHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DdHTMLElement) CustomData(key, value string) *DdHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DdHTMLElement) CustomDataRemove(key string) *DdHTMLElement {
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
func (e *DdHTMLElement) ACCESSKEY(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DdHTMLElement) IfACCESSKEY(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DdHTMLElement) RemoveACCESSKEY(v string) *DdHTMLElement {
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
func (e *DdHTMLElement) AUTOCAPITALIZE(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DdHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DdHTMLElement) RemoveAUTOCAPITALIZE(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DdHTMLElement) AUTOFOCUS() *DdHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DdHTMLElement) IfAUTOFOCUS(cond bool) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DdHTMLElement) RemoveAUTOFOCUS() *DdHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DdHTMLElement) SetAUTOFOCUS(b bool) *DdHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DdHTMLElement) CLASS(v string) *DdHTMLElement {
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

func (e *DdHTMLElement) IfCLASS(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DdHTMLElement) SetCLASS(v string) *DdHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DdHTMLElement) RemoveCLASS(v string) *DdHTMLElement {
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
func (e *DdHTMLElement) CONTENTEDITABLE(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DdHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DdHTMLElement) RemoveCONTENTEDITABLE(v string) *DdHTMLElement {
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
func (e *DdHTMLElement) DIR(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DdHTMLElement) IfDIR(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DdHTMLElement) RemoveDIR(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *DdHTMLElement) DRAGGABLE(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DdHTMLElement) IfDRAGGABLE(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DdHTMLElement) RemoveDRAGGABLE(v string) *DdHTMLElement {
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
func (e *DdHTMLElement) ENTERKEYHINT(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DdHTMLElement) IfENTERKEYHINT(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DdHTMLElement) RemoveENTERKEYHINT(v string) *DdHTMLElement {
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
func (e *DdHTMLElement) HIDDEN(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DdHTMLElement) IfHIDDEN(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DdHTMLElement) RemoveHIDDEN(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DdHTMLElement) ID(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DdHTMLElement) IfID(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DdHTMLElement) RemoveID(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DdHTMLElement) INERT() *DdHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DdHTMLElement) IfINERT(cond bool) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DdHTMLElement) RemoveINERT() *DdHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DdHTMLElement) SetINERT(b bool) *DdHTMLElement {
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
func (e *DdHTMLElement) INPUTMODE(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DdHTMLElement) IfINPUTMODE(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DdHTMLElement) RemoveINPUTMODE(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *DdHTMLElement) IS(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DdHTMLElement) IfIS(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DdHTMLElement) RemoveIS(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DdHTMLElement) ITEMID(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DdHTMLElement) IfITEMID(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DdHTMLElement) RemoveITEMID(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *DdHTMLElement) ITEMPROP(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DdHTMLElement) IfITEMPROP(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DdHTMLElement) RemoveITEMPROP(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DdHTMLElement) ITEMREF(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DdHTMLElement) IfITEMREF(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DdHTMLElement) RemoveITEMREF(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DdHTMLElement) ITEMSCOPE() *DdHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DdHTMLElement) IfITEMSCOPE(cond bool) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DdHTMLElement) RemoveITEMSCOPE() *DdHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DdHTMLElement) SetITEMSCOPE(b bool) *DdHTMLElement {
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
func (e *DdHTMLElement) ITEMTYPE(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DdHTMLElement) IfITEMTYPE(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DdHTMLElement) RemoveITEMTYPE(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DdHTMLElement) LANG(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DdHTMLElement) IfLANG(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DdHTMLElement) RemoveLANG(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DdHTMLElement) NONCE(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DdHTMLElement) IfNONCE(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DdHTMLElement) RemoveNONCE(v string) *DdHTMLElement {
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
func (e *DdHTMLElement) POPOVER(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DdHTMLElement) IfPOPOVER(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DdHTMLElement) RemovePOPOVER(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DdHTMLElement) SLOT(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DdHTMLElement) IfSLOT(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DdHTMLElement) RemoveSLOT(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *DdHTMLElement) SPELLCHECK(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DdHTMLElement) IfSPELLCHECK(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DdHTMLElement) RemoveSPELLCHECK(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DdHTMLElement) STYLE(k, v string) *DdHTMLElement {
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

func (e *DdHTMLElement) IfSTYLE(cond bool, k string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DdHTMLElement) RemoveSTYLE(k string) *DdHTMLElement {
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
func (e *DdHTMLElement) TABINDEX(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DdHTMLElement) IfTABINDEX(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DdHTMLElement) RemoveTABINDEX(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DdHTMLElement) TITLE(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DdHTMLElement) IfTITLE(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DdHTMLElement) RemoveTITLE(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *DdHTMLElement) TRANSLATE(v string) *DdHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DdHTMLElement) IfTRANSLATE(cond bool, v string) *DdHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DdHTMLElement) RemoveTRANSLATE(v string) *DdHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
