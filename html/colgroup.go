package html

import (
	"fmt"
)

type ColgroupHTMLElement struct {
	*Element
}

func COLGROUP(children ...ElementBuilder) *ColgroupHTMLElement {
	return &ColgroupHTMLElement{
		Element: &Element{
			Tag:           elementTagCOLGROUP,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *ColgroupHTMLElement) Children(children ...ElementBuilder) *ColgroupHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *ColgroupHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ColgroupHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *ColgroupHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ColgroupHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *ColgroupHTMLElement) Text(text string) *ColgroupHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *ColgroupHTMLElement) TextF(format string, args ...any) *ColgroupHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *ColgroupHTMLElement) Escaped(text string) *ColgroupHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *ColgroupHTMLElement) EscapedF(format string, args ...any) *ColgroupHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *ColgroupHTMLElement) CustomData(key, value string) *ColgroupHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ColgroupHTMLElement) CustomDataRemove(key string) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) ACCESSKEY(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfACCESSKEY(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *ColgroupHTMLElement) RemoveACCESSKEY(v string) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) AUTOCAPITALIZE(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *ColgroupHTMLElement) RemoveAUTOCAPITALIZE(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *ColgroupHTMLElement) AUTOFOCUS() *ColgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *ColgroupHTMLElement) IfAUTOFOCUS(cond bool) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *ColgroupHTMLElement) RemoveAUTOFOCUS() *ColgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *ColgroupHTMLElement) SetAUTOFOCUS(b bool) *ColgroupHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *ColgroupHTMLElement) CLASS(v string) *ColgroupHTMLElement {
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

func (e *ColgroupHTMLElement) IfCLASS(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *ColgroupHTMLElement) SetCLASS(v string) *ColgroupHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *ColgroupHTMLElement) RemoveCLASS(v string) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) CONTENTEDITABLE(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *ColgroupHTMLElement) RemoveCONTENTEDITABLE(v string) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) DIR(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfDIR(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *ColgroupHTMLElement) RemoveDIR(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *ColgroupHTMLElement) DRAGGABLE(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfDRAGGABLE(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *ColgroupHTMLElement) RemoveDRAGGABLE(v string) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) ENTERKEYHINT(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfENTERKEYHINT(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *ColgroupHTMLElement) RemoveENTERKEYHINT(v string) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) HIDDEN(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfHIDDEN(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *ColgroupHTMLElement) RemoveHIDDEN(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *ColgroupHTMLElement) ID(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfID(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *ColgroupHTMLElement) RemoveID(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *ColgroupHTMLElement) INERT() *ColgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *ColgroupHTMLElement) IfINERT(cond bool) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *ColgroupHTMLElement) RemoveINERT() *ColgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *ColgroupHTMLElement) SetINERT(b bool) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) INPUTMODE(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfINPUTMODE(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *ColgroupHTMLElement) RemoveINPUTMODE(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *ColgroupHTMLElement) IS(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfIS(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *ColgroupHTMLElement) RemoveIS(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *ColgroupHTMLElement) ITEMID(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfITEMID(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *ColgroupHTMLElement) RemoveITEMID(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *ColgroupHTMLElement) ITEMPROP(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfITEMPROP(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *ColgroupHTMLElement) RemoveITEMPROP(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ColgroupHTMLElement) ITEMREF(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfITEMREF(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *ColgroupHTMLElement) RemoveITEMREF(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *ColgroupHTMLElement) ITEMSCOPE() *ColgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *ColgroupHTMLElement) IfITEMSCOPE(cond bool) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *ColgroupHTMLElement) RemoveITEMSCOPE() *ColgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *ColgroupHTMLElement) SetITEMSCOPE(b bool) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) ITEMTYPE(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfITEMTYPE(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *ColgroupHTMLElement) RemoveITEMTYPE(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ColgroupHTMLElement) LANG(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfLANG(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *ColgroupHTMLElement) RemoveLANG(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *ColgroupHTMLElement) NONCE(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfNONCE(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *ColgroupHTMLElement) RemoveNONCE(v string) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) POPOVER(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfPOPOVER(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *ColgroupHTMLElement) RemovePOPOVER(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *ColgroupHTMLElement) SLOT(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfSLOT(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *ColgroupHTMLElement) RemoveSLOT(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPAN sets the "span" attribute.
// Number of columns spanned by the element
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *ColgroupHTMLElement) SPAN(v int) *ColgroupHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeSPANKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfSPAN(cond bool, v int) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.SPAN(v)
}

func (e *ColgroupHTMLElement) RemoveSPAN(v int) *ColgroupHTMLElement {
	delete(e.IntAttributes, attributeSPANKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *ColgroupHTMLElement) SPELLCHECK(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfSPELLCHECK(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *ColgroupHTMLElement) RemoveSPELLCHECK(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ColgroupHTMLElement) STYLE(k, v string) *ColgroupHTMLElement {
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

func (e *ColgroupHTMLElement) IfSTYLE(cond bool, k string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *ColgroupHTMLElement) RemoveSTYLE(k string) *ColgroupHTMLElement {
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
func (e *ColgroupHTMLElement) TABINDEX(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfTABINDEX(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *ColgroupHTMLElement) RemoveTABINDEX(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *ColgroupHTMLElement) TITLE(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfTITLE(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *ColgroupHTMLElement) RemoveTITLE(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *ColgroupHTMLElement) TRANSLATE(v string) *ColgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *ColgroupHTMLElement) IfTRANSLATE(cond bool, v string) *ColgroupHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *ColgroupHTMLElement) RemoveTRANSLATE(v string) *ColgroupHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
