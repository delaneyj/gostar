package html

import (
	"fmt"
)

type PreHTMLElement struct {
	*Element
}

func PRE(children ...ElementBuilder) *PreHTMLElement {
	return &PreHTMLElement{
		Element: &Element{
			Tag:           elementTagPRE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *PreHTMLElement) Children(children ...ElementBuilder) *PreHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *PreHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *PreHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *PreHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *PreHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *PreHTMLElement) Text(text string) *PreHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *PreHTMLElement) TextF(format string, args ...any) *PreHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *PreHTMLElement) Escaped(text string) *PreHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *PreHTMLElement) EscapedF(format string, args ...any) *PreHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *PreHTMLElement) CustomData(key, value string) *PreHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *PreHTMLElement) CustomDataRemove(key string) *PreHTMLElement {
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
func (e *PreHTMLElement) ACCESSKEY(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *PreHTMLElement) IfACCESSKEY(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *PreHTMLElement) RemoveACCESSKEY(v string) *PreHTMLElement {
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
func (e *PreHTMLElement) AUTOCAPITALIZE(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *PreHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *PreHTMLElement) RemoveAUTOCAPITALIZE(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *PreHTMLElement) AUTOFOCUS() *PreHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *PreHTMLElement) IfAUTOFOCUS(cond bool) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *PreHTMLElement) RemoveAUTOFOCUS() *PreHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *PreHTMLElement) SetAUTOFOCUS(b bool) *PreHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *PreHTMLElement) CLASS(v string) *PreHTMLElement {
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

func (e *PreHTMLElement) IfCLASS(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *PreHTMLElement) SetCLASS(v string) *PreHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *PreHTMLElement) RemoveCLASS(v string) *PreHTMLElement {
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
func (e *PreHTMLElement) CONTENTEDITABLE(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *PreHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *PreHTMLElement) RemoveCONTENTEDITABLE(v string) *PreHTMLElement {
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
func (e *PreHTMLElement) DIR(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *PreHTMLElement) IfDIR(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *PreHTMLElement) RemoveDIR(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *PreHTMLElement) DRAGGABLE(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *PreHTMLElement) IfDRAGGABLE(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *PreHTMLElement) RemoveDRAGGABLE(v string) *PreHTMLElement {
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
func (e *PreHTMLElement) ENTERKEYHINT(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *PreHTMLElement) IfENTERKEYHINT(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *PreHTMLElement) RemoveENTERKEYHINT(v string) *PreHTMLElement {
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
func (e *PreHTMLElement) HIDDEN(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *PreHTMLElement) IfHIDDEN(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *PreHTMLElement) RemoveHIDDEN(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *PreHTMLElement) ID(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *PreHTMLElement) IfID(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *PreHTMLElement) RemoveID(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *PreHTMLElement) INERT() *PreHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *PreHTMLElement) IfINERT(cond bool) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *PreHTMLElement) RemoveINERT() *PreHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *PreHTMLElement) SetINERT(b bool) *PreHTMLElement {
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
func (e *PreHTMLElement) INPUTMODE(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *PreHTMLElement) IfINPUTMODE(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *PreHTMLElement) RemoveINPUTMODE(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *PreHTMLElement) IS(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *PreHTMLElement) IfIS(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *PreHTMLElement) RemoveIS(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *PreHTMLElement) ITEMID(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *PreHTMLElement) IfITEMID(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *PreHTMLElement) RemoveITEMID(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *PreHTMLElement) ITEMPROP(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *PreHTMLElement) IfITEMPROP(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *PreHTMLElement) RemoveITEMPROP(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *PreHTMLElement) ITEMREF(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *PreHTMLElement) IfITEMREF(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *PreHTMLElement) RemoveITEMREF(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *PreHTMLElement) ITEMSCOPE() *PreHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *PreHTMLElement) IfITEMSCOPE(cond bool) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *PreHTMLElement) RemoveITEMSCOPE() *PreHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *PreHTMLElement) SetITEMSCOPE(b bool) *PreHTMLElement {
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
func (e *PreHTMLElement) ITEMTYPE(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *PreHTMLElement) IfITEMTYPE(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *PreHTMLElement) RemoveITEMTYPE(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *PreHTMLElement) LANG(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *PreHTMLElement) IfLANG(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *PreHTMLElement) RemoveLANG(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *PreHTMLElement) NONCE(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *PreHTMLElement) IfNONCE(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *PreHTMLElement) RemoveNONCE(v string) *PreHTMLElement {
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
func (e *PreHTMLElement) POPOVER(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *PreHTMLElement) IfPOPOVER(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *PreHTMLElement) RemovePOPOVER(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *PreHTMLElement) SLOT(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *PreHTMLElement) IfSLOT(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *PreHTMLElement) RemoveSLOT(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *PreHTMLElement) SPELLCHECK(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *PreHTMLElement) IfSPELLCHECK(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *PreHTMLElement) RemoveSPELLCHECK(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *PreHTMLElement) STYLE(k, v string) *PreHTMLElement {
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

func (e *PreHTMLElement) IfSTYLE(cond bool, k string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *PreHTMLElement) RemoveSTYLE(k string) *PreHTMLElement {
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
func (e *PreHTMLElement) TABINDEX(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *PreHTMLElement) IfTABINDEX(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *PreHTMLElement) RemoveTABINDEX(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *PreHTMLElement) TITLE(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *PreHTMLElement) IfTITLE(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *PreHTMLElement) RemoveTITLE(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *PreHTMLElement) TRANSLATE(v string) *PreHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *PreHTMLElement) IfTRANSLATE(cond bool, v string) *PreHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *PreHTMLElement) RemoveTRANSLATE(v string) *PreHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
