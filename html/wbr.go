package html

import (
	"fmt"
)

type WbrHTMLElement struct {
	*Element
}

func WBR(children ...ElementBuilder) *WbrHTMLElement {
	return &WbrHTMLElement{
		Element: &Element{
			Tag:           elementTagWBR,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *WbrHTMLElement) Children(children ...ElementBuilder) *WbrHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *WbrHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *WbrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *WbrHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *WbrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *WbrHTMLElement) Text(text string) *WbrHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *WbrHTMLElement) TextF(format string, args ...any) *WbrHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *WbrHTMLElement) Escaped(text string) *WbrHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *WbrHTMLElement) EscapedF(format string, args ...any) *WbrHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *WbrHTMLElement) CustomData(key, value string) *WbrHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *WbrHTMLElement) CustomDataRemove(key string) *WbrHTMLElement {
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
func (e *WbrHTMLElement) ACCESSKEY(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *WbrHTMLElement) IfACCESSKEY(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *WbrHTMLElement) RemoveACCESSKEY(v string) *WbrHTMLElement {
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
func (e *WbrHTMLElement) AUTOCAPITALIZE(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *WbrHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *WbrHTMLElement) RemoveAUTOCAPITALIZE(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *WbrHTMLElement) AUTOFOCUS() *WbrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *WbrHTMLElement) IfAUTOFOCUS(cond bool) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *WbrHTMLElement) RemoveAUTOFOCUS() *WbrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *WbrHTMLElement) SetAUTOFOCUS(b bool) *WbrHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *WbrHTMLElement) CLASS(v string) *WbrHTMLElement {
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

func (e *WbrHTMLElement) IfCLASS(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *WbrHTMLElement) SetCLASS(v string) *WbrHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *WbrHTMLElement) RemoveCLASS(v string) *WbrHTMLElement {
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
func (e *WbrHTMLElement) CONTENTEDITABLE(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *WbrHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *WbrHTMLElement) RemoveCONTENTEDITABLE(v string) *WbrHTMLElement {
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
func (e *WbrHTMLElement) DIR(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *WbrHTMLElement) IfDIR(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *WbrHTMLElement) RemoveDIR(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *WbrHTMLElement) DRAGGABLE(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *WbrHTMLElement) IfDRAGGABLE(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *WbrHTMLElement) RemoveDRAGGABLE(v string) *WbrHTMLElement {
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
func (e *WbrHTMLElement) ENTERKEYHINT(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *WbrHTMLElement) IfENTERKEYHINT(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *WbrHTMLElement) RemoveENTERKEYHINT(v string) *WbrHTMLElement {
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
func (e *WbrHTMLElement) HIDDEN(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *WbrHTMLElement) IfHIDDEN(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *WbrHTMLElement) RemoveHIDDEN(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *WbrHTMLElement) ID(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *WbrHTMLElement) IfID(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *WbrHTMLElement) RemoveID(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *WbrHTMLElement) INERT() *WbrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *WbrHTMLElement) IfINERT(cond bool) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *WbrHTMLElement) RemoveINERT() *WbrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *WbrHTMLElement) SetINERT(b bool) *WbrHTMLElement {
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
func (e *WbrHTMLElement) INPUTMODE(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *WbrHTMLElement) IfINPUTMODE(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *WbrHTMLElement) RemoveINPUTMODE(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *WbrHTMLElement) IS(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *WbrHTMLElement) IfIS(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *WbrHTMLElement) RemoveIS(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *WbrHTMLElement) ITEMID(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *WbrHTMLElement) IfITEMID(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *WbrHTMLElement) RemoveITEMID(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *WbrHTMLElement) ITEMPROP(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *WbrHTMLElement) IfITEMPROP(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *WbrHTMLElement) RemoveITEMPROP(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *WbrHTMLElement) ITEMREF(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *WbrHTMLElement) IfITEMREF(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *WbrHTMLElement) RemoveITEMREF(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *WbrHTMLElement) ITEMSCOPE() *WbrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *WbrHTMLElement) IfITEMSCOPE(cond bool) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *WbrHTMLElement) RemoveITEMSCOPE() *WbrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *WbrHTMLElement) SetITEMSCOPE(b bool) *WbrHTMLElement {
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
func (e *WbrHTMLElement) ITEMTYPE(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *WbrHTMLElement) IfITEMTYPE(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *WbrHTMLElement) RemoveITEMTYPE(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *WbrHTMLElement) LANG(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *WbrHTMLElement) IfLANG(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *WbrHTMLElement) RemoveLANG(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *WbrHTMLElement) NONCE(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *WbrHTMLElement) IfNONCE(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *WbrHTMLElement) RemoveNONCE(v string) *WbrHTMLElement {
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
func (e *WbrHTMLElement) POPOVER(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *WbrHTMLElement) IfPOPOVER(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *WbrHTMLElement) RemovePOPOVER(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *WbrHTMLElement) SLOT(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *WbrHTMLElement) IfSLOT(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *WbrHTMLElement) RemoveSLOT(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *WbrHTMLElement) SPELLCHECK(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *WbrHTMLElement) IfSPELLCHECK(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *WbrHTMLElement) RemoveSPELLCHECK(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *WbrHTMLElement) STYLE(k, v string) *WbrHTMLElement {
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

func (e *WbrHTMLElement) IfSTYLE(cond bool, k string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *WbrHTMLElement) RemoveSTYLE(k string) *WbrHTMLElement {
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
func (e *WbrHTMLElement) TABINDEX(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *WbrHTMLElement) IfTABINDEX(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *WbrHTMLElement) RemoveTABINDEX(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *WbrHTMLElement) TITLE(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *WbrHTMLElement) IfTITLE(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *WbrHTMLElement) RemoveTITLE(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *WbrHTMLElement) TRANSLATE(v string) *WbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *WbrHTMLElement) IfTRANSLATE(cond bool, v string) *WbrHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *WbrHTMLElement) RemoveTRANSLATE(v string) *WbrHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
