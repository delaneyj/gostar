package html

import (
	"fmt"
)

type SpanHTMLElement struct {
	*Element
}

func SPAN(children ...ElementBuilder) *SpanHTMLElement {
	return &SpanHTMLElement{
		Element: &Element{
			Tag:           elementTagSPAN,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SpanHTMLElement) Children(children ...ElementBuilder) *SpanHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SpanHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SpanHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SpanHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SpanHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SpanHTMLElement) Text(text string) *SpanHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SpanHTMLElement) TextF(format string, args ...any) *SpanHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SpanHTMLElement) Escaped(text string) *SpanHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SpanHTMLElement) EscapedF(format string, args ...any) *SpanHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SpanHTMLElement) CustomData(key, value string) *SpanHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SpanHTMLElement) CustomDataRemove(key string) *SpanHTMLElement {
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
func (e *SpanHTMLElement) ACCESSKEY(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SpanHTMLElement) IfACCESSKEY(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SpanHTMLElement) RemoveACCESSKEY(v string) *SpanHTMLElement {
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
func (e *SpanHTMLElement) AUTOCAPITALIZE(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SpanHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SpanHTMLElement) RemoveAUTOCAPITALIZE(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SpanHTMLElement) AUTOFOCUS() *SpanHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SpanHTMLElement) IfAUTOFOCUS(cond bool) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SpanHTMLElement) RemoveAUTOFOCUS() *SpanHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SpanHTMLElement) SetAUTOFOCUS(b bool) *SpanHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SpanHTMLElement) CLASS(v string) *SpanHTMLElement {
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

func (e *SpanHTMLElement) IfCLASS(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SpanHTMLElement) SetCLASS(v string) *SpanHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SpanHTMLElement) RemoveCLASS(v string) *SpanHTMLElement {
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
func (e *SpanHTMLElement) CONTENTEDITABLE(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SpanHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SpanHTMLElement) RemoveCONTENTEDITABLE(v string) *SpanHTMLElement {
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
func (e *SpanHTMLElement) DIR(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SpanHTMLElement) IfDIR(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SpanHTMLElement) RemoveDIR(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *SpanHTMLElement) DRAGGABLE(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SpanHTMLElement) IfDRAGGABLE(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SpanHTMLElement) RemoveDRAGGABLE(v string) *SpanHTMLElement {
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
func (e *SpanHTMLElement) ENTERKEYHINT(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SpanHTMLElement) IfENTERKEYHINT(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SpanHTMLElement) RemoveENTERKEYHINT(v string) *SpanHTMLElement {
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
func (e *SpanHTMLElement) HIDDEN(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SpanHTMLElement) IfHIDDEN(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SpanHTMLElement) RemoveHIDDEN(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SpanHTMLElement) ID(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SpanHTMLElement) IfID(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SpanHTMLElement) RemoveID(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SpanHTMLElement) INERT() *SpanHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SpanHTMLElement) IfINERT(cond bool) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SpanHTMLElement) RemoveINERT() *SpanHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SpanHTMLElement) SetINERT(b bool) *SpanHTMLElement {
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
func (e *SpanHTMLElement) INPUTMODE(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SpanHTMLElement) IfINPUTMODE(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SpanHTMLElement) RemoveINPUTMODE(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *SpanHTMLElement) IS(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SpanHTMLElement) IfIS(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SpanHTMLElement) RemoveIS(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SpanHTMLElement) ITEMID(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SpanHTMLElement) IfITEMID(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SpanHTMLElement) RemoveITEMID(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *SpanHTMLElement) ITEMPROP(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SpanHTMLElement) IfITEMPROP(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SpanHTMLElement) RemoveITEMPROP(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SpanHTMLElement) ITEMREF(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SpanHTMLElement) IfITEMREF(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SpanHTMLElement) RemoveITEMREF(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SpanHTMLElement) ITEMSCOPE() *SpanHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SpanHTMLElement) IfITEMSCOPE(cond bool) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SpanHTMLElement) RemoveITEMSCOPE() *SpanHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SpanHTMLElement) SetITEMSCOPE(b bool) *SpanHTMLElement {
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
func (e *SpanHTMLElement) ITEMTYPE(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SpanHTMLElement) IfITEMTYPE(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SpanHTMLElement) RemoveITEMTYPE(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SpanHTMLElement) LANG(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SpanHTMLElement) IfLANG(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SpanHTMLElement) RemoveLANG(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SpanHTMLElement) NONCE(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SpanHTMLElement) IfNONCE(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SpanHTMLElement) RemoveNONCE(v string) *SpanHTMLElement {
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
func (e *SpanHTMLElement) POPOVER(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SpanHTMLElement) IfPOPOVER(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SpanHTMLElement) RemovePOPOVER(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SpanHTMLElement) SLOT(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SpanHTMLElement) IfSLOT(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SpanHTMLElement) RemoveSLOT(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *SpanHTMLElement) SPELLCHECK(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SpanHTMLElement) IfSPELLCHECK(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SpanHTMLElement) RemoveSPELLCHECK(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SpanHTMLElement) STYLE(k, v string) *SpanHTMLElement {
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

func (e *SpanHTMLElement) IfSTYLE(cond bool, k string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SpanHTMLElement) RemoveSTYLE(k string) *SpanHTMLElement {
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
func (e *SpanHTMLElement) TABINDEX(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SpanHTMLElement) IfTABINDEX(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SpanHTMLElement) RemoveTABINDEX(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SpanHTMLElement) TITLE(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SpanHTMLElement) IfTITLE(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SpanHTMLElement) RemoveTITLE(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *SpanHTMLElement) TRANSLATE(v string) *SpanHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SpanHTMLElement) IfTRANSLATE(cond bool, v string) *SpanHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SpanHTMLElement) RemoveTRANSLATE(v string) *SpanHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
