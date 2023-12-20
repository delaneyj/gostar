package html

import (
	"fmt"
)

type TbodyHTMLElement struct {
	*Element
}

func TBODY(children ...ElementBuilder) *TbodyHTMLElement {
	return &TbodyHTMLElement{
		Element: &Element{
			Tag:           elementTagTBODY,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *TbodyHTMLElement) Children(children ...ElementBuilder) *TbodyHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *TbodyHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TbodyHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *TbodyHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TbodyHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *TbodyHTMLElement) Text(text string) *TbodyHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *TbodyHTMLElement) TextF(format string, args ...any) *TbodyHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *TbodyHTMLElement) Escaped(text string) *TbodyHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *TbodyHTMLElement) EscapedF(format string, args ...any) *TbodyHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TbodyHTMLElement) CustomData(key, value string) *TbodyHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TbodyHTMLElement) CustomDataRemove(key string) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) ACCESSKEY(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *TbodyHTMLElement) IfACCESSKEY(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *TbodyHTMLElement) RemoveACCESSKEY(v string) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) AUTOCAPITALIZE(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *TbodyHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *TbodyHTMLElement) RemoveAUTOCAPITALIZE(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *TbodyHTMLElement) AUTOFOCUS() *TbodyHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *TbodyHTMLElement) IfAUTOFOCUS(cond bool) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *TbodyHTMLElement) RemoveAUTOFOCUS() *TbodyHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *TbodyHTMLElement) SetAUTOFOCUS(b bool) *TbodyHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *TbodyHTMLElement) CLASS(v string) *TbodyHTMLElement {
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

func (e *TbodyHTMLElement) IfCLASS(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *TbodyHTMLElement) SetCLASS(v string) *TbodyHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *TbodyHTMLElement) RemoveCLASS(v string) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) CONTENTEDITABLE(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *TbodyHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *TbodyHTMLElement) RemoveCONTENTEDITABLE(v string) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) DIR(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *TbodyHTMLElement) IfDIR(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *TbodyHTMLElement) RemoveDIR(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *TbodyHTMLElement) DRAGGABLE(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *TbodyHTMLElement) IfDRAGGABLE(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *TbodyHTMLElement) RemoveDRAGGABLE(v string) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) ENTERKEYHINT(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *TbodyHTMLElement) IfENTERKEYHINT(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *TbodyHTMLElement) RemoveENTERKEYHINT(v string) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) HIDDEN(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *TbodyHTMLElement) IfHIDDEN(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *TbodyHTMLElement) RemoveHIDDEN(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *TbodyHTMLElement) ID(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *TbodyHTMLElement) IfID(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *TbodyHTMLElement) RemoveID(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *TbodyHTMLElement) INERT() *TbodyHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *TbodyHTMLElement) IfINERT(cond bool) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *TbodyHTMLElement) RemoveINERT() *TbodyHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *TbodyHTMLElement) SetINERT(b bool) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) INPUTMODE(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *TbodyHTMLElement) IfINPUTMODE(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *TbodyHTMLElement) RemoveINPUTMODE(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *TbodyHTMLElement) IS(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *TbodyHTMLElement) IfIS(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *TbodyHTMLElement) RemoveIS(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *TbodyHTMLElement) ITEMID(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *TbodyHTMLElement) IfITEMID(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *TbodyHTMLElement) RemoveITEMID(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *TbodyHTMLElement) ITEMPROP(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *TbodyHTMLElement) IfITEMPROP(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *TbodyHTMLElement) RemoveITEMPROP(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TbodyHTMLElement) ITEMREF(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *TbodyHTMLElement) IfITEMREF(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *TbodyHTMLElement) RemoveITEMREF(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *TbodyHTMLElement) ITEMSCOPE() *TbodyHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *TbodyHTMLElement) IfITEMSCOPE(cond bool) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *TbodyHTMLElement) RemoveITEMSCOPE() *TbodyHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *TbodyHTMLElement) SetITEMSCOPE(b bool) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) ITEMTYPE(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *TbodyHTMLElement) IfITEMTYPE(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *TbodyHTMLElement) RemoveITEMTYPE(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TbodyHTMLElement) LANG(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *TbodyHTMLElement) IfLANG(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *TbodyHTMLElement) RemoveLANG(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *TbodyHTMLElement) NONCE(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *TbodyHTMLElement) IfNONCE(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *TbodyHTMLElement) RemoveNONCE(v string) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) POPOVER(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *TbodyHTMLElement) IfPOPOVER(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *TbodyHTMLElement) RemovePOPOVER(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *TbodyHTMLElement) SLOT(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *TbodyHTMLElement) IfSLOT(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *TbodyHTMLElement) RemoveSLOT(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *TbodyHTMLElement) SPELLCHECK(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *TbodyHTMLElement) IfSPELLCHECK(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *TbodyHTMLElement) RemoveSPELLCHECK(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TbodyHTMLElement) STYLE(k, v string) *TbodyHTMLElement {
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

func (e *TbodyHTMLElement) IfSTYLE(cond bool, k string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *TbodyHTMLElement) RemoveSTYLE(k string) *TbodyHTMLElement {
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
func (e *TbodyHTMLElement) TABINDEX(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *TbodyHTMLElement) IfTABINDEX(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *TbodyHTMLElement) RemoveTABINDEX(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *TbodyHTMLElement) TITLE(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *TbodyHTMLElement) IfTITLE(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *TbodyHTMLElement) RemoveTITLE(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *TbodyHTMLElement) TRANSLATE(v string) *TbodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *TbodyHTMLElement) IfTRANSLATE(cond bool, v string) *TbodyHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *TbodyHTMLElement) RemoveTRANSLATE(v string) *TbodyHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
