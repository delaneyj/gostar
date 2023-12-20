package html

import (
	"fmt"
)

type SearchHTMLElement struct {
	*Element
}

func SEARCH(children ...ElementBuilder) *SearchHTMLElement {
	return &SearchHTMLElement{
		Element: &Element{
			Tag:           elementTagSEARCH,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SearchHTMLElement) Children(children ...ElementBuilder) *SearchHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SearchHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SearchHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SearchHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SearchHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SearchHTMLElement) Text(text string) *SearchHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SearchHTMLElement) TextF(format string, args ...any) *SearchHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SearchHTMLElement) Escaped(text string) *SearchHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SearchHTMLElement) EscapedF(format string, args ...any) *SearchHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SearchHTMLElement) CustomData(key, value string) *SearchHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SearchHTMLElement) CustomDataRemove(key string) *SearchHTMLElement {
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
func (e *SearchHTMLElement) ACCESSKEY(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SearchHTMLElement) IfACCESSKEY(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SearchHTMLElement) RemoveACCESSKEY(v string) *SearchHTMLElement {
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
func (e *SearchHTMLElement) AUTOCAPITALIZE(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SearchHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SearchHTMLElement) RemoveAUTOCAPITALIZE(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SearchHTMLElement) AUTOFOCUS() *SearchHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SearchHTMLElement) IfAUTOFOCUS(cond bool) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SearchHTMLElement) RemoveAUTOFOCUS() *SearchHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SearchHTMLElement) SetAUTOFOCUS(b bool) *SearchHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SearchHTMLElement) CLASS(v string) *SearchHTMLElement {
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

func (e *SearchHTMLElement) IfCLASS(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SearchHTMLElement) SetCLASS(v string) *SearchHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SearchHTMLElement) RemoveCLASS(v string) *SearchHTMLElement {
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
func (e *SearchHTMLElement) CONTENTEDITABLE(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SearchHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SearchHTMLElement) RemoveCONTENTEDITABLE(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *SearchHTMLElement) DIR(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SearchHTMLElement) IfDIR(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SearchHTMLElement) RemoveDIR(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *SearchHTMLElement) DRAGGABLE(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SearchHTMLElement) IfDRAGGABLE(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SearchHTMLElement) RemoveDRAGGABLE(v string) *SearchHTMLElement {
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
func (e *SearchHTMLElement) ENTERKEYHINT(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SearchHTMLElement) IfENTERKEYHINT(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SearchHTMLElement) RemoveENTERKEYHINT(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *SearchHTMLElement) HIDDEN(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SearchHTMLElement) IfHIDDEN(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SearchHTMLElement) RemoveHIDDEN(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SearchHTMLElement) ID(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SearchHTMLElement) IfID(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SearchHTMLElement) RemoveID(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SearchHTMLElement) INERT() *SearchHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SearchHTMLElement) IfINERT(cond bool) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SearchHTMLElement) RemoveINERT() *SearchHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SearchHTMLElement) SetINERT(b bool) *SearchHTMLElement {
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
func (e *SearchHTMLElement) INPUTMODE(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SearchHTMLElement) IfINPUTMODE(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SearchHTMLElement) RemoveINPUTMODE(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *SearchHTMLElement) IS(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SearchHTMLElement) IfIS(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SearchHTMLElement) RemoveIS(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SearchHTMLElement) ITEMID(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SearchHTMLElement) IfITEMID(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SearchHTMLElement) RemoveITEMID(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *SearchHTMLElement) ITEMPROP(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SearchHTMLElement) IfITEMPROP(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SearchHTMLElement) RemoveITEMPROP(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SearchHTMLElement) ITEMREF(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SearchHTMLElement) IfITEMREF(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SearchHTMLElement) RemoveITEMREF(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SearchHTMLElement) ITEMSCOPE() *SearchHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SearchHTMLElement) IfITEMSCOPE(cond bool) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SearchHTMLElement) RemoveITEMSCOPE() *SearchHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SearchHTMLElement) SetITEMSCOPE(b bool) *SearchHTMLElement {
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
func (e *SearchHTMLElement) ITEMTYPE(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SearchHTMLElement) IfITEMTYPE(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SearchHTMLElement) RemoveITEMTYPE(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SearchHTMLElement) LANG(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SearchHTMLElement) IfLANG(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SearchHTMLElement) RemoveLANG(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SearchHTMLElement) NONCE(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SearchHTMLElement) IfNONCE(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SearchHTMLElement) RemoveNONCE(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *SearchHTMLElement) POPOVER(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SearchHTMLElement) IfPOPOVER(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SearchHTMLElement) RemovePOPOVER(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SearchHTMLElement) SLOT(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SearchHTMLElement) IfSLOT(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SearchHTMLElement) RemoveSLOT(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *SearchHTMLElement) SPELLCHECK(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SearchHTMLElement) IfSPELLCHECK(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SearchHTMLElement) RemoveSPELLCHECK(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SearchHTMLElement) STYLE(k, v string) *SearchHTMLElement {
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

func (e *SearchHTMLElement) IfSTYLE(cond bool, k string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SearchHTMLElement) RemoveSTYLE(k string) *SearchHTMLElement {
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
func (e *SearchHTMLElement) TABINDEX(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SearchHTMLElement) IfTABINDEX(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SearchHTMLElement) RemoveTABINDEX(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SearchHTMLElement) TITLE(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SearchHTMLElement) IfTITLE(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SearchHTMLElement) RemoveTITLE(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *SearchHTMLElement) TRANSLATE(v string) *SearchHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SearchHTMLElement) IfTRANSLATE(cond bool, v string) *SearchHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SearchHTMLElement) RemoveTRANSLATE(v string) *SearchHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
