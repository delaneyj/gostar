package html

import (
	"fmt"
)

type IHTMLElement struct {
	*Element
}

func I(children ...ElementBuilder) *IHTMLElement {
	return &IHTMLElement{
		Element: &Element{
			Tag:           elementTagI,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *IHTMLElement) Children(children ...ElementBuilder) *IHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *IHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *IHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *IHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *IHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *IHTMLElement) Text(text string) *IHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *IHTMLElement) TextF(format string, args ...any) *IHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *IHTMLElement) Escaped(text string) *IHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *IHTMLElement) EscapedF(format string, args ...any) *IHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *IHTMLElement) CustomData(key, value string) *IHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *IHTMLElement) CustomDataRemove(key string) *IHTMLElement {
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
func (e *IHTMLElement) ACCESSKEY(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *IHTMLElement) IfACCESSKEY(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *IHTMLElement) RemoveACCESSKEY(v string) *IHTMLElement {
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
func (e *IHTMLElement) AUTOCAPITALIZE(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *IHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *IHTMLElement) RemoveAUTOCAPITALIZE(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *IHTMLElement) AUTOFOCUS() *IHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *IHTMLElement) IfAUTOFOCUS(cond bool) *IHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *IHTMLElement) RemoveAUTOFOCUS() *IHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *IHTMLElement) SetAUTOFOCUS(b bool) *IHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *IHTMLElement) CLASS(v string) *IHTMLElement {
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

func (e *IHTMLElement) IfCLASS(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *IHTMLElement) SetCLASS(v string) *IHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *IHTMLElement) RemoveCLASS(v string) *IHTMLElement {
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
func (e *IHTMLElement) CONTENTEDITABLE(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *IHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *IHTMLElement) RemoveCONTENTEDITABLE(v string) *IHTMLElement {
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
func (e *IHTMLElement) DIR(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *IHTMLElement) IfDIR(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *IHTMLElement) RemoveDIR(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *IHTMLElement) DRAGGABLE(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *IHTMLElement) IfDRAGGABLE(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *IHTMLElement) RemoveDRAGGABLE(v string) *IHTMLElement {
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
func (e *IHTMLElement) ENTERKEYHINT(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *IHTMLElement) IfENTERKEYHINT(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *IHTMLElement) RemoveENTERKEYHINT(v string) *IHTMLElement {
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
func (e *IHTMLElement) HIDDEN(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *IHTMLElement) IfHIDDEN(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *IHTMLElement) RemoveHIDDEN(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *IHTMLElement) ID(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *IHTMLElement) IfID(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *IHTMLElement) RemoveID(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *IHTMLElement) INERT() *IHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *IHTMLElement) IfINERT(cond bool) *IHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *IHTMLElement) RemoveINERT() *IHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *IHTMLElement) SetINERT(b bool) *IHTMLElement {
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
func (e *IHTMLElement) INPUTMODE(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *IHTMLElement) IfINPUTMODE(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *IHTMLElement) RemoveINPUTMODE(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *IHTMLElement) IS(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *IHTMLElement) IfIS(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *IHTMLElement) RemoveIS(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *IHTMLElement) ITEMID(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *IHTMLElement) IfITEMID(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *IHTMLElement) RemoveITEMID(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *IHTMLElement) ITEMPROP(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *IHTMLElement) IfITEMPROP(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *IHTMLElement) RemoveITEMPROP(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *IHTMLElement) ITEMREF(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *IHTMLElement) IfITEMREF(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *IHTMLElement) RemoveITEMREF(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *IHTMLElement) ITEMSCOPE() *IHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *IHTMLElement) IfITEMSCOPE(cond bool) *IHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *IHTMLElement) RemoveITEMSCOPE() *IHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *IHTMLElement) SetITEMSCOPE(b bool) *IHTMLElement {
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
func (e *IHTMLElement) ITEMTYPE(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *IHTMLElement) IfITEMTYPE(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *IHTMLElement) RemoveITEMTYPE(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *IHTMLElement) LANG(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *IHTMLElement) IfLANG(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *IHTMLElement) RemoveLANG(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *IHTMLElement) NONCE(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *IHTMLElement) IfNONCE(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *IHTMLElement) RemoveNONCE(v string) *IHTMLElement {
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
func (e *IHTMLElement) POPOVER(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *IHTMLElement) IfPOPOVER(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *IHTMLElement) RemovePOPOVER(v string) *IHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *IHTMLElement) SLOT(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *IHTMLElement) IfSLOT(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *IHTMLElement) RemoveSLOT(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *IHTMLElement) SPELLCHECK(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *IHTMLElement) IfSPELLCHECK(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *IHTMLElement) RemoveSPELLCHECK(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *IHTMLElement) STYLE(k, v string) *IHTMLElement {
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

func (e *IHTMLElement) IfSTYLE(cond bool, k string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *IHTMLElement) RemoveSTYLE(k string) *IHTMLElement {
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
func (e *IHTMLElement) TABINDEX(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *IHTMLElement) IfTABINDEX(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *IHTMLElement) RemoveTABINDEX(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *IHTMLElement) TITLE(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *IHTMLElement) IfTITLE(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *IHTMLElement) RemoveTITLE(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *IHTMLElement) TRANSLATE(v string) *IHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *IHTMLElement) IfTRANSLATE(cond bool, v string) *IHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *IHTMLElement) RemoveTRANSLATE(v string) *IHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
