package html

import (
	"fmt"
)

type HeadHTMLElement struct {
	*Element
}

func HEAD(children ...ElementBuilder) *HeadHTMLElement {
	return &HeadHTMLElement{
		Element: &Element{
			Tag:           elementTagHEAD,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *HeadHTMLElement) Children(children ...ElementBuilder) *HeadHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *HeadHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *HeadHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *HeadHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *HeadHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *HeadHTMLElement) Text(text string) *HeadHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *HeadHTMLElement) TextF(format string, args ...any) *HeadHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *HeadHTMLElement) Escaped(text string) *HeadHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *HeadHTMLElement) EscapedF(format string, args ...any) *HeadHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *HeadHTMLElement) CustomData(key, value string) *HeadHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *HeadHTMLElement) CustomDataRemove(key string) *HeadHTMLElement {
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
func (e *HeadHTMLElement) ACCESSKEY(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *HeadHTMLElement) IfACCESSKEY(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *HeadHTMLElement) RemoveACCESSKEY(v string) *HeadHTMLElement {
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
func (e *HeadHTMLElement) AUTOCAPITALIZE(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *HeadHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *HeadHTMLElement) RemoveAUTOCAPITALIZE(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *HeadHTMLElement) AUTOFOCUS() *HeadHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *HeadHTMLElement) IfAUTOFOCUS(cond bool) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *HeadHTMLElement) RemoveAUTOFOCUS() *HeadHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *HeadHTMLElement) SetAUTOFOCUS(b bool) *HeadHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *HeadHTMLElement) CLASS(v string) *HeadHTMLElement {
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

func (e *HeadHTMLElement) IfCLASS(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *HeadHTMLElement) SetCLASS(v string) *HeadHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *HeadHTMLElement) RemoveCLASS(v string) *HeadHTMLElement {
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
func (e *HeadHTMLElement) CONTENTEDITABLE(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *HeadHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *HeadHTMLElement) RemoveCONTENTEDITABLE(v string) *HeadHTMLElement {
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
func (e *HeadHTMLElement) DIR(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *HeadHTMLElement) IfDIR(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *HeadHTMLElement) RemoveDIR(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *HeadHTMLElement) DRAGGABLE(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *HeadHTMLElement) IfDRAGGABLE(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *HeadHTMLElement) RemoveDRAGGABLE(v string) *HeadHTMLElement {
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
func (e *HeadHTMLElement) ENTERKEYHINT(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *HeadHTMLElement) IfENTERKEYHINT(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *HeadHTMLElement) RemoveENTERKEYHINT(v string) *HeadHTMLElement {
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
func (e *HeadHTMLElement) HIDDEN(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *HeadHTMLElement) IfHIDDEN(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *HeadHTMLElement) RemoveHIDDEN(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *HeadHTMLElement) ID(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *HeadHTMLElement) IfID(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *HeadHTMLElement) RemoveID(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *HeadHTMLElement) INERT() *HeadHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *HeadHTMLElement) IfINERT(cond bool) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *HeadHTMLElement) RemoveINERT() *HeadHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *HeadHTMLElement) SetINERT(b bool) *HeadHTMLElement {
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
func (e *HeadHTMLElement) INPUTMODE(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *HeadHTMLElement) IfINPUTMODE(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *HeadHTMLElement) RemoveINPUTMODE(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *HeadHTMLElement) IS(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *HeadHTMLElement) IfIS(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *HeadHTMLElement) RemoveIS(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *HeadHTMLElement) ITEMID(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *HeadHTMLElement) IfITEMID(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *HeadHTMLElement) RemoveITEMID(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *HeadHTMLElement) ITEMPROP(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *HeadHTMLElement) IfITEMPROP(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *HeadHTMLElement) RemoveITEMPROP(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *HeadHTMLElement) ITEMREF(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *HeadHTMLElement) IfITEMREF(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *HeadHTMLElement) RemoveITEMREF(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *HeadHTMLElement) ITEMSCOPE() *HeadHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *HeadHTMLElement) IfITEMSCOPE(cond bool) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *HeadHTMLElement) RemoveITEMSCOPE() *HeadHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *HeadHTMLElement) SetITEMSCOPE(b bool) *HeadHTMLElement {
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
func (e *HeadHTMLElement) ITEMTYPE(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *HeadHTMLElement) IfITEMTYPE(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *HeadHTMLElement) RemoveITEMTYPE(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *HeadHTMLElement) LANG(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *HeadHTMLElement) IfLANG(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *HeadHTMLElement) RemoveLANG(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *HeadHTMLElement) NONCE(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *HeadHTMLElement) IfNONCE(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *HeadHTMLElement) RemoveNONCE(v string) *HeadHTMLElement {
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
func (e *HeadHTMLElement) POPOVER(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *HeadHTMLElement) IfPOPOVER(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *HeadHTMLElement) RemovePOPOVER(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *HeadHTMLElement) SLOT(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *HeadHTMLElement) IfSLOT(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *HeadHTMLElement) RemoveSLOT(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *HeadHTMLElement) SPELLCHECK(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *HeadHTMLElement) IfSPELLCHECK(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *HeadHTMLElement) RemoveSPELLCHECK(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *HeadHTMLElement) STYLE(k, v string) *HeadHTMLElement {
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

func (e *HeadHTMLElement) IfSTYLE(cond bool, k string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *HeadHTMLElement) RemoveSTYLE(k string) *HeadHTMLElement {
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
func (e *HeadHTMLElement) TABINDEX(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *HeadHTMLElement) IfTABINDEX(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *HeadHTMLElement) RemoveTABINDEX(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *HeadHTMLElement) TITLE(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *HeadHTMLElement) IfTITLE(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *HeadHTMLElement) RemoveTITLE(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *HeadHTMLElement) TRANSLATE(v string) *HeadHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *HeadHTMLElement) IfTRANSLATE(cond bool, v string) *HeadHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *HeadHTMLElement) RemoveTRANSLATE(v string) *HeadHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
