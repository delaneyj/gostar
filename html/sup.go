package html

import (
	"fmt"
)

type SupHTMLElement struct {
	*Element
}

func SUP(children ...ElementBuilder) *SupHTMLElement {
	return &SupHTMLElement{
		Element: &Element{
			Tag:           elementTagSUP,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SupHTMLElement) Children(children ...ElementBuilder) *SupHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SupHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SupHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SupHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SupHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SupHTMLElement) Text(text string) *SupHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SupHTMLElement) TextF(format string, args ...any) *SupHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SupHTMLElement) Escaped(text string) *SupHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SupHTMLElement) EscapedF(format string, args ...any) *SupHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SupHTMLElement) CustomData(key, value string) *SupHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SupHTMLElement) CustomDataRemove(key string) *SupHTMLElement {
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
func (e *SupHTMLElement) ACCESSKEY(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SupHTMLElement) IfACCESSKEY(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SupHTMLElement) RemoveACCESSKEY(v string) *SupHTMLElement {
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
func (e *SupHTMLElement) AUTOCAPITALIZE(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SupHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SupHTMLElement) RemoveAUTOCAPITALIZE(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SupHTMLElement) AUTOFOCUS() *SupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SupHTMLElement) IfAUTOFOCUS(cond bool) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SupHTMLElement) RemoveAUTOFOCUS() *SupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SupHTMLElement) SetAUTOFOCUS(b bool) *SupHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SupHTMLElement) CLASS(v string) *SupHTMLElement {
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

func (e *SupHTMLElement) IfCLASS(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SupHTMLElement) SetCLASS(v string) *SupHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SupHTMLElement) RemoveCLASS(v string) *SupHTMLElement {
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
func (e *SupHTMLElement) CONTENTEDITABLE(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SupHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SupHTMLElement) RemoveCONTENTEDITABLE(v string) *SupHTMLElement {
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
func (e *SupHTMLElement) DIR(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SupHTMLElement) IfDIR(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SupHTMLElement) RemoveDIR(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *SupHTMLElement) DRAGGABLE(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SupHTMLElement) IfDRAGGABLE(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SupHTMLElement) RemoveDRAGGABLE(v string) *SupHTMLElement {
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
func (e *SupHTMLElement) ENTERKEYHINT(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SupHTMLElement) IfENTERKEYHINT(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SupHTMLElement) RemoveENTERKEYHINT(v string) *SupHTMLElement {
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
func (e *SupHTMLElement) HIDDEN(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SupHTMLElement) IfHIDDEN(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SupHTMLElement) RemoveHIDDEN(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SupHTMLElement) ID(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SupHTMLElement) IfID(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SupHTMLElement) RemoveID(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SupHTMLElement) INERT() *SupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SupHTMLElement) IfINERT(cond bool) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SupHTMLElement) RemoveINERT() *SupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SupHTMLElement) SetINERT(b bool) *SupHTMLElement {
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
func (e *SupHTMLElement) INPUTMODE(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SupHTMLElement) IfINPUTMODE(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SupHTMLElement) RemoveINPUTMODE(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *SupHTMLElement) IS(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SupHTMLElement) IfIS(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SupHTMLElement) RemoveIS(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SupHTMLElement) ITEMID(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SupHTMLElement) IfITEMID(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SupHTMLElement) RemoveITEMID(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *SupHTMLElement) ITEMPROP(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SupHTMLElement) IfITEMPROP(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SupHTMLElement) RemoveITEMPROP(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SupHTMLElement) ITEMREF(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SupHTMLElement) IfITEMREF(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SupHTMLElement) RemoveITEMREF(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SupHTMLElement) ITEMSCOPE() *SupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SupHTMLElement) IfITEMSCOPE(cond bool) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SupHTMLElement) RemoveITEMSCOPE() *SupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SupHTMLElement) SetITEMSCOPE(b bool) *SupHTMLElement {
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
func (e *SupHTMLElement) ITEMTYPE(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SupHTMLElement) IfITEMTYPE(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SupHTMLElement) RemoveITEMTYPE(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SupHTMLElement) LANG(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SupHTMLElement) IfLANG(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SupHTMLElement) RemoveLANG(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SupHTMLElement) NONCE(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SupHTMLElement) IfNONCE(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SupHTMLElement) RemoveNONCE(v string) *SupHTMLElement {
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
func (e *SupHTMLElement) POPOVER(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SupHTMLElement) IfPOPOVER(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SupHTMLElement) RemovePOPOVER(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SupHTMLElement) SLOT(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SupHTMLElement) IfSLOT(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SupHTMLElement) RemoveSLOT(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *SupHTMLElement) SPELLCHECK(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SupHTMLElement) IfSPELLCHECK(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SupHTMLElement) RemoveSPELLCHECK(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SupHTMLElement) STYLE(k, v string) *SupHTMLElement {
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

func (e *SupHTMLElement) IfSTYLE(cond bool, k string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SupHTMLElement) RemoveSTYLE(k string) *SupHTMLElement {
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
func (e *SupHTMLElement) TABINDEX(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SupHTMLElement) IfTABINDEX(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SupHTMLElement) RemoveTABINDEX(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SupHTMLElement) TITLE(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SupHTMLElement) IfTITLE(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SupHTMLElement) RemoveTITLE(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *SupHTMLElement) TRANSLATE(v string) *SupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SupHTMLElement) IfTRANSLATE(cond bool, v string) *SupHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SupHTMLElement) RemoveTRANSLATE(v string) *SupHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
