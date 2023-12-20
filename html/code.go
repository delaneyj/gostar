package html

import (
	"fmt"
)

type CodeHTMLElement struct {
	*Element
}

func CODE(children ...ElementBuilder) *CodeHTMLElement {
	return &CodeHTMLElement{
		Element: &Element{
			Tag:           elementTagCODE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *CodeHTMLElement) Children(children ...ElementBuilder) *CodeHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *CodeHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *CodeHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *CodeHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *CodeHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *CodeHTMLElement) Text(text string) *CodeHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *CodeHTMLElement) TextF(format string, args ...any) *CodeHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *CodeHTMLElement) Escaped(text string) *CodeHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *CodeHTMLElement) EscapedF(format string, args ...any) *CodeHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *CodeHTMLElement) CustomData(key, value string) *CodeHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *CodeHTMLElement) CustomDataRemove(key string) *CodeHTMLElement {
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
func (e *CodeHTMLElement) ACCESSKEY(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *CodeHTMLElement) IfACCESSKEY(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *CodeHTMLElement) RemoveACCESSKEY(v string) *CodeHTMLElement {
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
func (e *CodeHTMLElement) AUTOCAPITALIZE(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *CodeHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *CodeHTMLElement) RemoveAUTOCAPITALIZE(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *CodeHTMLElement) AUTOFOCUS() *CodeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *CodeHTMLElement) IfAUTOFOCUS(cond bool) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *CodeHTMLElement) RemoveAUTOFOCUS() *CodeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *CodeHTMLElement) SetAUTOFOCUS(b bool) *CodeHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *CodeHTMLElement) CLASS(v string) *CodeHTMLElement {
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

func (e *CodeHTMLElement) IfCLASS(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *CodeHTMLElement) SetCLASS(v string) *CodeHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *CodeHTMLElement) RemoveCLASS(v string) *CodeHTMLElement {
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
func (e *CodeHTMLElement) CONTENTEDITABLE(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *CodeHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *CodeHTMLElement) RemoveCONTENTEDITABLE(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *CodeHTMLElement) DIR(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *CodeHTMLElement) IfDIR(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *CodeHTMLElement) RemoveDIR(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *CodeHTMLElement) DRAGGABLE(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *CodeHTMLElement) IfDRAGGABLE(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *CodeHTMLElement) RemoveDRAGGABLE(v string) *CodeHTMLElement {
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
func (e *CodeHTMLElement) ENTERKEYHINT(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *CodeHTMLElement) IfENTERKEYHINT(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *CodeHTMLElement) RemoveENTERKEYHINT(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *CodeHTMLElement) HIDDEN(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *CodeHTMLElement) IfHIDDEN(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *CodeHTMLElement) RemoveHIDDEN(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *CodeHTMLElement) ID(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *CodeHTMLElement) IfID(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *CodeHTMLElement) RemoveID(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *CodeHTMLElement) INERT() *CodeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *CodeHTMLElement) IfINERT(cond bool) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *CodeHTMLElement) RemoveINERT() *CodeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *CodeHTMLElement) SetINERT(b bool) *CodeHTMLElement {
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
func (e *CodeHTMLElement) INPUTMODE(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *CodeHTMLElement) IfINPUTMODE(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *CodeHTMLElement) RemoveINPUTMODE(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *CodeHTMLElement) IS(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *CodeHTMLElement) IfIS(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *CodeHTMLElement) RemoveIS(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *CodeHTMLElement) ITEMID(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *CodeHTMLElement) IfITEMID(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *CodeHTMLElement) RemoveITEMID(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *CodeHTMLElement) ITEMPROP(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *CodeHTMLElement) IfITEMPROP(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *CodeHTMLElement) RemoveITEMPROP(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *CodeHTMLElement) ITEMREF(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *CodeHTMLElement) IfITEMREF(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *CodeHTMLElement) RemoveITEMREF(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *CodeHTMLElement) ITEMSCOPE() *CodeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *CodeHTMLElement) IfITEMSCOPE(cond bool) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *CodeHTMLElement) RemoveITEMSCOPE() *CodeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *CodeHTMLElement) SetITEMSCOPE(b bool) *CodeHTMLElement {
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
func (e *CodeHTMLElement) ITEMTYPE(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *CodeHTMLElement) IfITEMTYPE(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *CodeHTMLElement) RemoveITEMTYPE(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *CodeHTMLElement) LANG(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *CodeHTMLElement) IfLANG(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *CodeHTMLElement) RemoveLANG(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *CodeHTMLElement) NONCE(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *CodeHTMLElement) IfNONCE(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *CodeHTMLElement) RemoveNONCE(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *CodeHTMLElement) POPOVER(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *CodeHTMLElement) IfPOPOVER(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *CodeHTMLElement) RemovePOPOVER(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *CodeHTMLElement) SLOT(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *CodeHTMLElement) IfSLOT(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *CodeHTMLElement) RemoveSLOT(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *CodeHTMLElement) SPELLCHECK(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *CodeHTMLElement) IfSPELLCHECK(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *CodeHTMLElement) RemoveSPELLCHECK(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *CodeHTMLElement) STYLE(k, v string) *CodeHTMLElement {
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

func (e *CodeHTMLElement) IfSTYLE(cond bool, k string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *CodeHTMLElement) RemoveSTYLE(k string) *CodeHTMLElement {
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
func (e *CodeHTMLElement) TABINDEX(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *CodeHTMLElement) IfTABINDEX(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *CodeHTMLElement) RemoveTABINDEX(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *CodeHTMLElement) TITLE(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *CodeHTMLElement) IfTITLE(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *CodeHTMLElement) RemoveTITLE(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *CodeHTMLElement) TRANSLATE(v string) *CodeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *CodeHTMLElement) IfTRANSLATE(cond bool, v string) *CodeHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *CodeHTMLElement) RemoveTRANSLATE(v string) *CodeHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
