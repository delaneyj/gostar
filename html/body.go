package html

import (
	"fmt"
)

type BodyHTMLElement struct {
	*Element
}

func BODY(children ...ElementBuilder) *BodyHTMLElement {
	return &BodyHTMLElement{
		Element: &Element{
			Tag:           elementTagBODY,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *BodyHTMLElement) Children(children ...ElementBuilder) *BodyHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *BodyHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *BodyHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *BodyHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *BodyHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *BodyHTMLElement) Text(text string) *BodyHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *BodyHTMLElement) TextF(format string, args ...any) *BodyHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *BodyHTMLElement) Escaped(text string) *BodyHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *BodyHTMLElement) EscapedF(format string, args ...any) *BodyHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *BodyHTMLElement) CustomData(key, value string) *BodyHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *BodyHTMLElement) CustomDataRemove(key string) *BodyHTMLElement {
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
func (e *BodyHTMLElement) ACCESSKEY(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *BodyHTMLElement) IfACCESSKEY(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *BodyHTMLElement) RemoveACCESSKEY(v string) *BodyHTMLElement {
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
func (e *BodyHTMLElement) AUTOCAPITALIZE(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *BodyHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *BodyHTMLElement) RemoveAUTOCAPITALIZE(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *BodyHTMLElement) AUTOFOCUS() *BodyHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *BodyHTMLElement) IfAUTOFOCUS(cond bool) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *BodyHTMLElement) RemoveAUTOFOCUS() *BodyHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *BodyHTMLElement) SetAUTOFOCUS(b bool) *BodyHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *BodyHTMLElement) CLASS(v string) *BodyHTMLElement {
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

func (e *BodyHTMLElement) IfCLASS(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *BodyHTMLElement) SetCLASS(v string) *BodyHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *BodyHTMLElement) RemoveCLASS(v string) *BodyHTMLElement {
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
func (e *BodyHTMLElement) CONTENTEDITABLE(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *BodyHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *BodyHTMLElement) RemoveCONTENTEDITABLE(v string) *BodyHTMLElement {
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
func (e *BodyHTMLElement) DIR(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *BodyHTMLElement) IfDIR(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *BodyHTMLElement) RemoveDIR(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *BodyHTMLElement) DRAGGABLE(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *BodyHTMLElement) IfDRAGGABLE(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *BodyHTMLElement) RemoveDRAGGABLE(v string) *BodyHTMLElement {
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
func (e *BodyHTMLElement) ENTERKEYHINT(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *BodyHTMLElement) IfENTERKEYHINT(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *BodyHTMLElement) RemoveENTERKEYHINT(v string) *BodyHTMLElement {
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
func (e *BodyHTMLElement) HIDDEN(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *BodyHTMLElement) IfHIDDEN(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *BodyHTMLElement) RemoveHIDDEN(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *BodyHTMLElement) ID(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *BodyHTMLElement) IfID(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *BodyHTMLElement) RemoveID(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *BodyHTMLElement) INERT() *BodyHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *BodyHTMLElement) IfINERT(cond bool) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *BodyHTMLElement) RemoveINERT() *BodyHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *BodyHTMLElement) SetINERT(b bool) *BodyHTMLElement {
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
func (e *BodyHTMLElement) INPUTMODE(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *BodyHTMLElement) IfINPUTMODE(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *BodyHTMLElement) RemoveINPUTMODE(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *BodyHTMLElement) IS(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *BodyHTMLElement) IfIS(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *BodyHTMLElement) RemoveIS(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *BodyHTMLElement) ITEMID(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *BodyHTMLElement) IfITEMID(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *BodyHTMLElement) RemoveITEMID(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *BodyHTMLElement) ITEMPROP(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *BodyHTMLElement) IfITEMPROP(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *BodyHTMLElement) RemoveITEMPROP(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *BodyHTMLElement) ITEMREF(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *BodyHTMLElement) IfITEMREF(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *BodyHTMLElement) RemoveITEMREF(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *BodyHTMLElement) ITEMSCOPE() *BodyHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *BodyHTMLElement) IfITEMSCOPE(cond bool) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *BodyHTMLElement) RemoveITEMSCOPE() *BodyHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *BodyHTMLElement) SetITEMSCOPE(b bool) *BodyHTMLElement {
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
func (e *BodyHTMLElement) ITEMTYPE(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *BodyHTMLElement) IfITEMTYPE(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *BodyHTMLElement) RemoveITEMTYPE(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *BodyHTMLElement) LANG(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *BodyHTMLElement) IfLANG(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *BodyHTMLElement) RemoveLANG(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *BodyHTMLElement) NONCE(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *BodyHTMLElement) IfNONCE(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *BodyHTMLElement) RemoveNONCE(v string) *BodyHTMLElement {
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
func (e *BodyHTMLElement) POPOVER(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *BodyHTMLElement) IfPOPOVER(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *BodyHTMLElement) RemovePOPOVER(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *BodyHTMLElement) SLOT(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *BodyHTMLElement) IfSLOT(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *BodyHTMLElement) RemoveSLOT(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *BodyHTMLElement) SPELLCHECK(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *BodyHTMLElement) IfSPELLCHECK(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *BodyHTMLElement) RemoveSPELLCHECK(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *BodyHTMLElement) STYLE(k, v string) *BodyHTMLElement {
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

func (e *BodyHTMLElement) IfSTYLE(cond bool, k string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *BodyHTMLElement) RemoveSTYLE(k string) *BodyHTMLElement {
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
func (e *BodyHTMLElement) TABINDEX(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *BodyHTMLElement) IfTABINDEX(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *BodyHTMLElement) RemoveTABINDEX(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *BodyHTMLElement) TITLE(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *BodyHTMLElement) IfTITLE(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *BodyHTMLElement) RemoveTITLE(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *BodyHTMLElement) TRANSLATE(v string) *BodyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *BodyHTMLElement) IfTRANSLATE(cond bool, v string) *BodyHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *BodyHTMLElement) RemoveTRANSLATE(v string) *BodyHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
