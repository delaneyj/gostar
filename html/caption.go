package html

import (
	"fmt"
)

type CaptionHTMLElement struct {
	*Element
}

func CAPTION(children ...ElementBuilder) *CaptionHTMLElement {
	return &CaptionHTMLElement{
		Element: &Element{
			Tag:           elementTagCAPTION,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *CaptionHTMLElement) Children(children ...ElementBuilder) *CaptionHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *CaptionHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *CaptionHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *CaptionHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *CaptionHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *CaptionHTMLElement) Text(text string) *CaptionHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *CaptionHTMLElement) TextF(format string, args ...any) *CaptionHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *CaptionHTMLElement) Escaped(text string) *CaptionHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *CaptionHTMLElement) EscapedF(format string, args ...any) *CaptionHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *CaptionHTMLElement) CustomData(key, value string) *CaptionHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *CaptionHTMLElement) CustomDataRemove(key string) *CaptionHTMLElement {
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
func (e *CaptionHTMLElement) ACCESSKEY(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *CaptionHTMLElement) IfACCESSKEY(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *CaptionHTMLElement) RemoveACCESSKEY(v string) *CaptionHTMLElement {
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
func (e *CaptionHTMLElement) AUTOCAPITALIZE(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *CaptionHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *CaptionHTMLElement) RemoveAUTOCAPITALIZE(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *CaptionHTMLElement) AUTOFOCUS() *CaptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *CaptionHTMLElement) IfAUTOFOCUS(cond bool) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *CaptionHTMLElement) RemoveAUTOFOCUS() *CaptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *CaptionHTMLElement) SetAUTOFOCUS(b bool) *CaptionHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *CaptionHTMLElement) CLASS(v string) *CaptionHTMLElement {
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

func (e *CaptionHTMLElement) IfCLASS(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *CaptionHTMLElement) SetCLASS(v string) *CaptionHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *CaptionHTMLElement) RemoveCLASS(v string) *CaptionHTMLElement {
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
func (e *CaptionHTMLElement) CONTENTEDITABLE(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *CaptionHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *CaptionHTMLElement) RemoveCONTENTEDITABLE(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *CaptionHTMLElement) DIR(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *CaptionHTMLElement) IfDIR(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *CaptionHTMLElement) RemoveDIR(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *CaptionHTMLElement) DRAGGABLE(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *CaptionHTMLElement) IfDRAGGABLE(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *CaptionHTMLElement) RemoveDRAGGABLE(v string) *CaptionHTMLElement {
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
func (e *CaptionHTMLElement) ENTERKEYHINT(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *CaptionHTMLElement) IfENTERKEYHINT(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *CaptionHTMLElement) RemoveENTERKEYHINT(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *CaptionHTMLElement) HIDDEN(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *CaptionHTMLElement) IfHIDDEN(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *CaptionHTMLElement) RemoveHIDDEN(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *CaptionHTMLElement) ID(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *CaptionHTMLElement) IfID(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *CaptionHTMLElement) RemoveID(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *CaptionHTMLElement) INERT() *CaptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *CaptionHTMLElement) IfINERT(cond bool) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *CaptionHTMLElement) RemoveINERT() *CaptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *CaptionHTMLElement) SetINERT(b bool) *CaptionHTMLElement {
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
func (e *CaptionHTMLElement) INPUTMODE(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *CaptionHTMLElement) IfINPUTMODE(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *CaptionHTMLElement) RemoveINPUTMODE(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *CaptionHTMLElement) IS(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *CaptionHTMLElement) IfIS(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *CaptionHTMLElement) RemoveIS(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *CaptionHTMLElement) ITEMID(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *CaptionHTMLElement) IfITEMID(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *CaptionHTMLElement) RemoveITEMID(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *CaptionHTMLElement) ITEMPROP(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *CaptionHTMLElement) IfITEMPROP(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *CaptionHTMLElement) RemoveITEMPROP(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *CaptionHTMLElement) ITEMREF(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *CaptionHTMLElement) IfITEMREF(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *CaptionHTMLElement) RemoveITEMREF(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *CaptionHTMLElement) ITEMSCOPE() *CaptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *CaptionHTMLElement) IfITEMSCOPE(cond bool) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *CaptionHTMLElement) RemoveITEMSCOPE() *CaptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *CaptionHTMLElement) SetITEMSCOPE(b bool) *CaptionHTMLElement {
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
func (e *CaptionHTMLElement) ITEMTYPE(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *CaptionHTMLElement) IfITEMTYPE(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *CaptionHTMLElement) RemoveITEMTYPE(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *CaptionHTMLElement) LANG(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *CaptionHTMLElement) IfLANG(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *CaptionHTMLElement) RemoveLANG(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *CaptionHTMLElement) NONCE(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *CaptionHTMLElement) IfNONCE(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *CaptionHTMLElement) RemoveNONCE(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *CaptionHTMLElement) POPOVER(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *CaptionHTMLElement) IfPOPOVER(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *CaptionHTMLElement) RemovePOPOVER(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *CaptionHTMLElement) SLOT(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *CaptionHTMLElement) IfSLOT(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *CaptionHTMLElement) RemoveSLOT(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *CaptionHTMLElement) SPELLCHECK(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *CaptionHTMLElement) IfSPELLCHECK(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *CaptionHTMLElement) RemoveSPELLCHECK(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *CaptionHTMLElement) STYLE(k, v string) *CaptionHTMLElement {
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

func (e *CaptionHTMLElement) IfSTYLE(cond bool, k string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *CaptionHTMLElement) RemoveSTYLE(k string) *CaptionHTMLElement {
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
func (e *CaptionHTMLElement) TABINDEX(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *CaptionHTMLElement) IfTABINDEX(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *CaptionHTMLElement) RemoveTABINDEX(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *CaptionHTMLElement) TITLE(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *CaptionHTMLElement) IfTITLE(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *CaptionHTMLElement) RemoveTITLE(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *CaptionHTMLElement) TRANSLATE(v string) *CaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *CaptionHTMLElement) IfTRANSLATE(cond bool, v string) *CaptionHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *CaptionHTMLElement) RemoveTRANSLATE(v string) *CaptionHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
