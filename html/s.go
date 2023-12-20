package html

import (
	"fmt"
)

type SHTMLElement struct {
	*Element
}

func S(children ...ElementBuilder) *SHTMLElement {
	return &SHTMLElement{
		Element: &Element{
			Tag:           elementTagS,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SHTMLElement) Children(children ...ElementBuilder) *SHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SHTMLElement) Text(text string) *SHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SHTMLElement) TextF(format string, args ...any) *SHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SHTMLElement) Escaped(text string) *SHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SHTMLElement) EscapedF(format string, args ...any) *SHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SHTMLElement) CustomData(key, value string) *SHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SHTMLElement) CustomDataRemove(key string) *SHTMLElement {
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
func (e *SHTMLElement) ACCESSKEY(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SHTMLElement) IfACCESSKEY(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SHTMLElement) RemoveACCESSKEY(v string) *SHTMLElement {
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
func (e *SHTMLElement) AUTOCAPITALIZE(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SHTMLElement) RemoveAUTOCAPITALIZE(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SHTMLElement) AUTOFOCUS() *SHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SHTMLElement) IfAUTOFOCUS(cond bool) *SHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SHTMLElement) RemoveAUTOFOCUS() *SHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SHTMLElement) SetAUTOFOCUS(b bool) *SHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SHTMLElement) CLASS(v string) *SHTMLElement {
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

func (e *SHTMLElement) IfCLASS(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SHTMLElement) SetCLASS(v string) *SHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SHTMLElement) RemoveCLASS(v string) *SHTMLElement {
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
func (e *SHTMLElement) CONTENTEDITABLE(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SHTMLElement) RemoveCONTENTEDITABLE(v string) *SHTMLElement {
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
func (e *SHTMLElement) DIR(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SHTMLElement) IfDIR(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SHTMLElement) RemoveDIR(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *SHTMLElement) DRAGGABLE(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SHTMLElement) IfDRAGGABLE(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SHTMLElement) RemoveDRAGGABLE(v string) *SHTMLElement {
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
func (e *SHTMLElement) ENTERKEYHINT(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SHTMLElement) IfENTERKEYHINT(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SHTMLElement) RemoveENTERKEYHINT(v string) *SHTMLElement {
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
func (e *SHTMLElement) HIDDEN(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SHTMLElement) IfHIDDEN(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SHTMLElement) RemoveHIDDEN(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SHTMLElement) ID(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SHTMLElement) IfID(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SHTMLElement) RemoveID(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SHTMLElement) INERT() *SHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SHTMLElement) IfINERT(cond bool) *SHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SHTMLElement) RemoveINERT() *SHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SHTMLElement) SetINERT(b bool) *SHTMLElement {
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
func (e *SHTMLElement) INPUTMODE(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SHTMLElement) IfINPUTMODE(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SHTMLElement) RemoveINPUTMODE(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *SHTMLElement) IS(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SHTMLElement) IfIS(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SHTMLElement) RemoveIS(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SHTMLElement) ITEMID(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SHTMLElement) IfITEMID(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SHTMLElement) RemoveITEMID(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *SHTMLElement) ITEMPROP(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SHTMLElement) IfITEMPROP(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SHTMLElement) RemoveITEMPROP(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SHTMLElement) ITEMREF(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SHTMLElement) IfITEMREF(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SHTMLElement) RemoveITEMREF(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SHTMLElement) ITEMSCOPE() *SHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SHTMLElement) IfITEMSCOPE(cond bool) *SHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SHTMLElement) RemoveITEMSCOPE() *SHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SHTMLElement) SetITEMSCOPE(b bool) *SHTMLElement {
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
func (e *SHTMLElement) ITEMTYPE(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SHTMLElement) IfITEMTYPE(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SHTMLElement) RemoveITEMTYPE(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SHTMLElement) LANG(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SHTMLElement) IfLANG(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SHTMLElement) RemoveLANG(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SHTMLElement) NONCE(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SHTMLElement) IfNONCE(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SHTMLElement) RemoveNONCE(v string) *SHTMLElement {
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
func (e *SHTMLElement) POPOVER(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SHTMLElement) IfPOPOVER(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SHTMLElement) RemovePOPOVER(v string) *SHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SHTMLElement) SLOT(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SHTMLElement) IfSLOT(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SHTMLElement) RemoveSLOT(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *SHTMLElement) SPELLCHECK(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SHTMLElement) IfSPELLCHECK(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SHTMLElement) RemoveSPELLCHECK(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SHTMLElement) STYLE(k, v string) *SHTMLElement {
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

func (e *SHTMLElement) IfSTYLE(cond bool, k string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SHTMLElement) RemoveSTYLE(k string) *SHTMLElement {
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
func (e *SHTMLElement) TABINDEX(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SHTMLElement) IfTABINDEX(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SHTMLElement) RemoveTABINDEX(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SHTMLElement) TITLE(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SHTMLElement) IfTITLE(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SHTMLElement) RemoveTITLE(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *SHTMLElement) TRANSLATE(v string) *SHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SHTMLElement) IfTRANSLATE(cond bool, v string) *SHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SHTMLElement) RemoveTRANSLATE(v string) *SHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
