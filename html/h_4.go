package html

import (
	"fmt"
)

type H4HTMLElement struct {
	*Element
}

func H4(children ...ElementBuilder) *H4HTMLElement {
	return &H4HTMLElement{
		Element: &Element{
			Tag:           elementTagH4,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *H4HTMLElement) Children(children ...ElementBuilder) *H4HTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *H4HTMLElement) IfChildren(condition bool, children ...ElementBuilder) *H4HTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *H4HTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *H4HTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *H4HTMLElement) Text(text string) *H4HTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *H4HTMLElement) TextF(format string, args ...any) *H4HTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *H4HTMLElement) Escaped(text string) *H4HTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *H4HTMLElement) EscapedF(format string, args ...any) *H4HTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *H4HTMLElement) CustomData(key, value string) *H4HTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *H4HTMLElement) CustomDataRemove(key string) *H4HTMLElement {
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
func (e *H4HTMLElement) ACCESSKEY(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *H4HTMLElement) IfACCESSKEY(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *H4HTMLElement) RemoveACCESSKEY(v string) *H4HTMLElement {
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
func (e *H4HTMLElement) AUTOCAPITALIZE(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *H4HTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *H4HTMLElement) RemoveAUTOCAPITALIZE(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *H4HTMLElement) AUTOFOCUS() *H4HTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *H4HTMLElement) IfAUTOFOCUS(cond bool) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *H4HTMLElement) RemoveAUTOFOCUS() *H4HTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *H4HTMLElement) SetAUTOFOCUS(b bool) *H4HTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *H4HTMLElement) CLASS(v string) *H4HTMLElement {
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

func (e *H4HTMLElement) IfCLASS(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *H4HTMLElement) SetCLASS(v string) *H4HTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *H4HTMLElement) RemoveCLASS(v string) *H4HTMLElement {
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
func (e *H4HTMLElement) CONTENTEDITABLE(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *H4HTMLElement) IfCONTENTEDITABLE(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *H4HTMLElement) RemoveCONTENTEDITABLE(v string) *H4HTMLElement {
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
func (e *H4HTMLElement) DIR(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *H4HTMLElement) IfDIR(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *H4HTMLElement) RemoveDIR(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *H4HTMLElement) DRAGGABLE(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *H4HTMLElement) IfDRAGGABLE(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *H4HTMLElement) RemoveDRAGGABLE(v string) *H4HTMLElement {
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
func (e *H4HTMLElement) ENTERKEYHINT(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *H4HTMLElement) IfENTERKEYHINT(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *H4HTMLElement) RemoveENTERKEYHINT(v string) *H4HTMLElement {
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
func (e *H4HTMLElement) HIDDEN(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *H4HTMLElement) IfHIDDEN(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *H4HTMLElement) RemoveHIDDEN(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *H4HTMLElement) ID(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *H4HTMLElement) IfID(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *H4HTMLElement) RemoveID(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *H4HTMLElement) INERT() *H4HTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *H4HTMLElement) IfINERT(cond bool) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *H4HTMLElement) RemoveINERT() *H4HTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *H4HTMLElement) SetINERT(b bool) *H4HTMLElement {
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
func (e *H4HTMLElement) INPUTMODE(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *H4HTMLElement) IfINPUTMODE(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *H4HTMLElement) RemoveINPUTMODE(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *H4HTMLElement) IS(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *H4HTMLElement) IfIS(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *H4HTMLElement) RemoveIS(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *H4HTMLElement) ITEMID(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *H4HTMLElement) IfITEMID(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *H4HTMLElement) RemoveITEMID(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *H4HTMLElement) ITEMPROP(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *H4HTMLElement) IfITEMPROP(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *H4HTMLElement) RemoveITEMPROP(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *H4HTMLElement) ITEMREF(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *H4HTMLElement) IfITEMREF(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *H4HTMLElement) RemoveITEMREF(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *H4HTMLElement) ITEMSCOPE() *H4HTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *H4HTMLElement) IfITEMSCOPE(cond bool) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *H4HTMLElement) RemoveITEMSCOPE() *H4HTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *H4HTMLElement) SetITEMSCOPE(b bool) *H4HTMLElement {
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
func (e *H4HTMLElement) ITEMTYPE(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *H4HTMLElement) IfITEMTYPE(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *H4HTMLElement) RemoveITEMTYPE(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *H4HTMLElement) LANG(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *H4HTMLElement) IfLANG(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *H4HTMLElement) RemoveLANG(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *H4HTMLElement) NONCE(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *H4HTMLElement) IfNONCE(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *H4HTMLElement) RemoveNONCE(v string) *H4HTMLElement {
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
func (e *H4HTMLElement) POPOVER(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *H4HTMLElement) IfPOPOVER(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *H4HTMLElement) RemovePOPOVER(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *H4HTMLElement) SLOT(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *H4HTMLElement) IfSLOT(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *H4HTMLElement) RemoveSLOT(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *H4HTMLElement) SPELLCHECK(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *H4HTMLElement) IfSPELLCHECK(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *H4HTMLElement) RemoveSPELLCHECK(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *H4HTMLElement) STYLE(k, v string) *H4HTMLElement {
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

func (e *H4HTMLElement) IfSTYLE(cond bool, k string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *H4HTMLElement) RemoveSTYLE(k string) *H4HTMLElement {
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
func (e *H4HTMLElement) TABINDEX(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *H4HTMLElement) IfTABINDEX(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *H4HTMLElement) RemoveTABINDEX(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *H4HTMLElement) TITLE(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *H4HTMLElement) IfTITLE(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *H4HTMLElement) RemoveTITLE(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *H4HTMLElement) TRANSLATE(v string) *H4HTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *H4HTMLElement) IfTRANSLATE(cond bool, v string) *H4HTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *H4HTMLElement) RemoveTRANSLATE(v string) *H4HTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
