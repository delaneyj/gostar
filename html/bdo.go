package html

import (
	"fmt"
)

type BdoHTMLElement struct {
	*Element
}

func BDO(children ...ElementBuilder) *BdoHTMLElement {
	return &BdoHTMLElement{
		Element: &Element{
			Tag:           elementTagBDO,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *BdoHTMLElement) Children(children ...ElementBuilder) *BdoHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *BdoHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *BdoHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *BdoHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *BdoHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *BdoHTMLElement) Text(text string) *BdoHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *BdoHTMLElement) TextF(format string, args ...any) *BdoHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *BdoHTMLElement) Escaped(text string) *BdoHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *BdoHTMLElement) EscapedF(format string, args ...any) *BdoHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *BdoHTMLElement) CustomData(key, value string) *BdoHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *BdoHTMLElement) CustomDataRemove(key string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) ACCESSKEY(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *BdoHTMLElement) IfACCESSKEY(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *BdoHTMLElement) RemoveACCESSKEY(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) AUTOCAPITALIZE(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *BdoHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *BdoHTMLElement) RemoveAUTOCAPITALIZE(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *BdoHTMLElement) AUTOFOCUS() *BdoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *BdoHTMLElement) IfAUTOFOCUS(cond bool) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *BdoHTMLElement) RemoveAUTOFOCUS() *BdoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *BdoHTMLElement) SetAUTOFOCUS(b bool) *BdoHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *BdoHTMLElement) CLASS(v string) *BdoHTMLElement {
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

func (e *BdoHTMLElement) IfCLASS(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *BdoHTMLElement) SetCLASS(v string) *BdoHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *BdoHTMLElement) RemoveCLASS(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) CONTENTEDITABLE(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *BdoHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *BdoHTMLElement) RemoveCONTENTEDITABLE(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) DIR(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *BdoHTMLElement) IfDIR(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *BdoHTMLElement) RemoveDIR(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *BdoHTMLElement) DRAGGABLE(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *BdoHTMLElement) IfDRAGGABLE(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *BdoHTMLElement) RemoveDRAGGABLE(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) ENTERKEYHINT(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *BdoHTMLElement) IfENTERKEYHINT(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *BdoHTMLElement) RemoveENTERKEYHINT(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) HIDDEN(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *BdoHTMLElement) IfHIDDEN(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *BdoHTMLElement) RemoveHIDDEN(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *BdoHTMLElement) ID(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *BdoHTMLElement) IfID(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *BdoHTMLElement) RemoveID(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *BdoHTMLElement) INERT() *BdoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *BdoHTMLElement) IfINERT(cond bool) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *BdoHTMLElement) RemoveINERT() *BdoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *BdoHTMLElement) SetINERT(b bool) *BdoHTMLElement {
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
func (e *BdoHTMLElement) INPUTMODE(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *BdoHTMLElement) IfINPUTMODE(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *BdoHTMLElement) RemoveINPUTMODE(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *BdoHTMLElement) IS(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *BdoHTMLElement) IfIS(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *BdoHTMLElement) RemoveIS(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *BdoHTMLElement) ITEMID(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *BdoHTMLElement) IfITEMID(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *BdoHTMLElement) RemoveITEMID(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *BdoHTMLElement) ITEMPROP(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *BdoHTMLElement) IfITEMPROP(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *BdoHTMLElement) RemoveITEMPROP(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *BdoHTMLElement) ITEMREF(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *BdoHTMLElement) IfITEMREF(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *BdoHTMLElement) RemoveITEMREF(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *BdoHTMLElement) ITEMSCOPE() *BdoHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *BdoHTMLElement) IfITEMSCOPE(cond bool) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *BdoHTMLElement) RemoveITEMSCOPE() *BdoHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *BdoHTMLElement) SetITEMSCOPE(b bool) *BdoHTMLElement {
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
func (e *BdoHTMLElement) ITEMTYPE(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *BdoHTMLElement) IfITEMTYPE(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *BdoHTMLElement) RemoveITEMTYPE(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *BdoHTMLElement) LANG(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *BdoHTMLElement) IfLANG(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *BdoHTMLElement) RemoveLANG(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *BdoHTMLElement) NONCE(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *BdoHTMLElement) IfNONCE(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *BdoHTMLElement) RemoveNONCE(v string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) POPOVER(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *BdoHTMLElement) IfPOPOVER(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *BdoHTMLElement) RemovePOPOVER(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *BdoHTMLElement) SLOT(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *BdoHTMLElement) IfSLOT(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *BdoHTMLElement) RemoveSLOT(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *BdoHTMLElement) SPELLCHECK(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *BdoHTMLElement) IfSPELLCHECK(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *BdoHTMLElement) RemoveSPELLCHECK(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *BdoHTMLElement) STYLE(k, v string) *BdoHTMLElement {
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

func (e *BdoHTMLElement) IfSTYLE(cond bool, k string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *BdoHTMLElement) RemoveSTYLE(k string) *BdoHTMLElement {
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
func (e *BdoHTMLElement) TABINDEX(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *BdoHTMLElement) IfTABINDEX(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *BdoHTMLElement) RemoveTABINDEX(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *BdoHTMLElement) TITLE(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *BdoHTMLElement) IfTITLE(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *BdoHTMLElement) RemoveTITLE(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *BdoHTMLElement) TRANSLATE(v string) *BdoHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *BdoHTMLElement) IfTRANSLATE(cond bool, v string) *BdoHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *BdoHTMLElement) RemoveTRANSLATE(v string) *BdoHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
