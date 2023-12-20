package html

import (
	"fmt"
)

type TheadHTMLElement struct {
	*Element
}

func THEAD(children ...ElementBuilder) *TheadHTMLElement {
	return &TheadHTMLElement{
		Element: &Element{
			Tag:           elementTagTHEAD,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *TheadHTMLElement) Children(children ...ElementBuilder) *TheadHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *TheadHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TheadHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *TheadHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TheadHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *TheadHTMLElement) Text(text string) *TheadHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *TheadHTMLElement) TextF(format string, args ...any) *TheadHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *TheadHTMLElement) Escaped(text string) *TheadHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *TheadHTMLElement) EscapedF(format string, args ...any) *TheadHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TheadHTMLElement) CustomData(key, value string) *TheadHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TheadHTMLElement) CustomDataRemove(key string) *TheadHTMLElement {
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
func (e *TheadHTMLElement) ACCESSKEY(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *TheadHTMLElement) IfACCESSKEY(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *TheadHTMLElement) RemoveACCESSKEY(v string) *TheadHTMLElement {
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
func (e *TheadHTMLElement) AUTOCAPITALIZE(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *TheadHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *TheadHTMLElement) RemoveAUTOCAPITALIZE(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *TheadHTMLElement) AUTOFOCUS() *TheadHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *TheadHTMLElement) IfAUTOFOCUS(cond bool) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *TheadHTMLElement) RemoveAUTOFOCUS() *TheadHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *TheadHTMLElement) SetAUTOFOCUS(b bool) *TheadHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *TheadHTMLElement) CLASS(v string) *TheadHTMLElement {
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

func (e *TheadHTMLElement) IfCLASS(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *TheadHTMLElement) SetCLASS(v string) *TheadHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *TheadHTMLElement) RemoveCLASS(v string) *TheadHTMLElement {
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
func (e *TheadHTMLElement) CONTENTEDITABLE(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *TheadHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *TheadHTMLElement) RemoveCONTENTEDITABLE(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *TheadHTMLElement) DIR(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *TheadHTMLElement) IfDIR(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *TheadHTMLElement) RemoveDIR(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *TheadHTMLElement) DRAGGABLE(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *TheadHTMLElement) IfDRAGGABLE(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *TheadHTMLElement) RemoveDRAGGABLE(v string) *TheadHTMLElement {
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
func (e *TheadHTMLElement) ENTERKEYHINT(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *TheadHTMLElement) IfENTERKEYHINT(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *TheadHTMLElement) RemoveENTERKEYHINT(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *TheadHTMLElement) HIDDEN(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *TheadHTMLElement) IfHIDDEN(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *TheadHTMLElement) RemoveHIDDEN(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *TheadHTMLElement) ID(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *TheadHTMLElement) IfID(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *TheadHTMLElement) RemoveID(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *TheadHTMLElement) INERT() *TheadHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *TheadHTMLElement) IfINERT(cond bool) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *TheadHTMLElement) RemoveINERT() *TheadHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *TheadHTMLElement) SetINERT(b bool) *TheadHTMLElement {
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
func (e *TheadHTMLElement) INPUTMODE(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *TheadHTMLElement) IfINPUTMODE(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *TheadHTMLElement) RemoveINPUTMODE(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *TheadHTMLElement) IS(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *TheadHTMLElement) IfIS(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *TheadHTMLElement) RemoveIS(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *TheadHTMLElement) ITEMID(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *TheadHTMLElement) IfITEMID(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *TheadHTMLElement) RemoveITEMID(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *TheadHTMLElement) ITEMPROP(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *TheadHTMLElement) IfITEMPROP(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *TheadHTMLElement) RemoveITEMPROP(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TheadHTMLElement) ITEMREF(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *TheadHTMLElement) IfITEMREF(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *TheadHTMLElement) RemoveITEMREF(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *TheadHTMLElement) ITEMSCOPE() *TheadHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *TheadHTMLElement) IfITEMSCOPE(cond bool) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *TheadHTMLElement) RemoveITEMSCOPE() *TheadHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *TheadHTMLElement) SetITEMSCOPE(b bool) *TheadHTMLElement {
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
func (e *TheadHTMLElement) ITEMTYPE(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *TheadHTMLElement) IfITEMTYPE(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *TheadHTMLElement) RemoveITEMTYPE(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TheadHTMLElement) LANG(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *TheadHTMLElement) IfLANG(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *TheadHTMLElement) RemoveLANG(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *TheadHTMLElement) NONCE(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *TheadHTMLElement) IfNONCE(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *TheadHTMLElement) RemoveNONCE(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *TheadHTMLElement) POPOVER(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *TheadHTMLElement) IfPOPOVER(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *TheadHTMLElement) RemovePOPOVER(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *TheadHTMLElement) SLOT(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *TheadHTMLElement) IfSLOT(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *TheadHTMLElement) RemoveSLOT(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *TheadHTMLElement) SPELLCHECK(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *TheadHTMLElement) IfSPELLCHECK(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *TheadHTMLElement) RemoveSPELLCHECK(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TheadHTMLElement) STYLE(k, v string) *TheadHTMLElement {
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

func (e *TheadHTMLElement) IfSTYLE(cond bool, k string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *TheadHTMLElement) RemoveSTYLE(k string) *TheadHTMLElement {
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
func (e *TheadHTMLElement) TABINDEX(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *TheadHTMLElement) IfTABINDEX(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *TheadHTMLElement) RemoveTABINDEX(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *TheadHTMLElement) TITLE(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *TheadHTMLElement) IfTITLE(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *TheadHTMLElement) RemoveTITLE(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *TheadHTMLElement) TRANSLATE(v string) *TheadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *TheadHTMLElement) IfTRANSLATE(cond bool, v string) *TheadHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *TheadHTMLElement) RemoveTRANSLATE(v string) *TheadHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
