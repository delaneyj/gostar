package html

import (
	"fmt"
)

type MenuHTMLElement struct {
	*Element
}

func MENU(children ...ElementBuilder) *MenuHTMLElement {
	return &MenuHTMLElement{
		Element: &Element{
			Tag:           elementTagMENU,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *MenuHTMLElement) Children(children ...ElementBuilder) *MenuHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MenuHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *MenuHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MenuHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *MenuHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MenuHTMLElement) Text(text string) *MenuHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MenuHTMLElement) TextF(format string, args ...any) *MenuHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MenuHTMLElement) Escaped(text string) *MenuHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MenuHTMLElement) EscapedF(format string, args ...any) *MenuHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MenuHTMLElement) CustomData(key, value string) *MenuHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *MenuHTMLElement) CustomDataRemove(key string) *MenuHTMLElement {
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
func (e *MenuHTMLElement) ACCESSKEY(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *MenuHTMLElement) IfACCESSKEY(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *MenuHTMLElement) RemoveACCESSKEY(v string) *MenuHTMLElement {
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
func (e *MenuHTMLElement) AUTOCAPITALIZE(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *MenuHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *MenuHTMLElement) RemoveAUTOCAPITALIZE(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *MenuHTMLElement) AUTOFOCUS() *MenuHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *MenuHTMLElement) IfAUTOFOCUS(cond bool) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *MenuHTMLElement) RemoveAUTOFOCUS() *MenuHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *MenuHTMLElement) SetAUTOFOCUS(b bool) *MenuHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *MenuHTMLElement) CLASS(v string) *MenuHTMLElement {
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

func (e *MenuHTMLElement) IfCLASS(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *MenuHTMLElement) SetCLASS(v string) *MenuHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *MenuHTMLElement) RemoveCLASS(v string) *MenuHTMLElement {
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
func (e *MenuHTMLElement) CONTENTEDITABLE(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *MenuHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *MenuHTMLElement) RemoveCONTENTEDITABLE(v string) *MenuHTMLElement {
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
func (e *MenuHTMLElement) DIR(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *MenuHTMLElement) IfDIR(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *MenuHTMLElement) RemoveDIR(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *MenuHTMLElement) DRAGGABLE(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *MenuHTMLElement) IfDRAGGABLE(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *MenuHTMLElement) RemoveDRAGGABLE(v string) *MenuHTMLElement {
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
func (e *MenuHTMLElement) ENTERKEYHINT(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *MenuHTMLElement) IfENTERKEYHINT(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *MenuHTMLElement) RemoveENTERKEYHINT(v string) *MenuHTMLElement {
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
func (e *MenuHTMLElement) HIDDEN(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *MenuHTMLElement) IfHIDDEN(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *MenuHTMLElement) RemoveHIDDEN(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *MenuHTMLElement) ID(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *MenuHTMLElement) IfID(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *MenuHTMLElement) RemoveID(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *MenuHTMLElement) INERT() *MenuHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *MenuHTMLElement) IfINERT(cond bool) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *MenuHTMLElement) RemoveINERT() *MenuHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *MenuHTMLElement) SetINERT(b bool) *MenuHTMLElement {
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
func (e *MenuHTMLElement) INPUTMODE(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *MenuHTMLElement) IfINPUTMODE(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *MenuHTMLElement) RemoveINPUTMODE(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *MenuHTMLElement) IS(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *MenuHTMLElement) IfIS(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *MenuHTMLElement) RemoveIS(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *MenuHTMLElement) ITEMID(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *MenuHTMLElement) IfITEMID(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *MenuHTMLElement) RemoveITEMID(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *MenuHTMLElement) ITEMPROP(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *MenuHTMLElement) IfITEMPROP(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *MenuHTMLElement) RemoveITEMPROP(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *MenuHTMLElement) ITEMREF(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *MenuHTMLElement) IfITEMREF(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *MenuHTMLElement) RemoveITEMREF(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *MenuHTMLElement) ITEMSCOPE() *MenuHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *MenuHTMLElement) IfITEMSCOPE(cond bool) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *MenuHTMLElement) RemoveITEMSCOPE() *MenuHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *MenuHTMLElement) SetITEMSCOPE(b bool) *MenuHTMLElement {
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
func (e *MenuHTMLElement) ITEMTYPE(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *MenuHTMLElement) IfITEMTYPE(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *MenuHTMLElement) RemoveITEMTYPE(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *MenuHTMLElement) LANG(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *MenuHTMLElement) IfLANG(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *MenuHTMLElement) RemoveLANG(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *MenuHTMLElement) NONCE(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *MenuHTMLElement) IfNONCE(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *MenuHTMLElement) RemoveNONCE(v string) *MenuHTMLElement {
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
func (e *MenuHTMLElement) POPOVER(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *MenuHTMLElement) IfPOPOVER(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *MenuHTMLElement) RemovePOPOVER(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *MenuHTMLElement) SLOT(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *MenuHTMLElement) IfSLOT(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *MenuHTMLElement) RemoveSLOT(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *MenuHTMLElement) SPELLCHECK(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *MenuHTMLElement) IfSPELLCHECK(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *MenuHTMLElement) RemoveSPELLCHECK(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *MenuHTMLElement) STYLE(k, v string) *MenuHTMLElement {
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

func (e *MenuHTMLElement) IfSTYLE(cond bool, k string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *MenuHTMLElement) RemoveSTYLE(k string) *MenuHTMLElement {
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
func (e *MenuHTMLElement) TABINDEX(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *MenuHTMLElement) IfTABINDEX(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *MenuHTMLElement) RemoveTABINDEX(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *MenuHTMLElement) TITLE(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *MenuHTMLElement) IfTITLE(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *MenuHTMLElement) RemoveTITLE(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *MenuHTMLElement) TRANSLATE(v string) *MenuHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *MenuHTMLElement) IfTRANSLATE(cond bool, v string) *MenuHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *MenuHTMLElement) RemoveTRANSLATE(v string) *MenuHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
