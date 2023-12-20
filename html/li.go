package html

import (
	"fmt"
)

type LiHTMLElement struct {
	*Element
}

func LI(children ...ElementBuilder) *LiHTMLElement {
	return &LiHTMLElement{
		Element: &Element{
			Tag:           elementTagLI,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *LiHTMLElement) Children(children ...ElementBuilder) *LiHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *LiHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *LiHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *LiHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *LiHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *LiHTMLElement) Text(text string) *LiHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *LiHTMLElement) TextF(format string, args ...any) *LiHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *LiHTMLElement) Escaped(text string) *LiHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *LiHTMLElement) EscapedF(format string, args ...any) *LiHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *LiHTMLElement) CustomData(key, value string) *LiHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *LiHTMLElement) CustomDataRemove(key string) *LiHTMLElement {
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
func (e *LiHTMLElement) ACCESSKEY(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *LiHTMLElement) IfACCESSKEY(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *LiHTMLElement) RemoveACCESSKEY(v string) *LiHTMLElement {
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
func (e *LiHTMLElement) AUTOCAPITALIZE(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *LiHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *LiHTMLElement) RemoveAUTOCAPITALIZE(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *LiHTMLElement) AUTOFOCUS() *LiHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *LiHTMLElement) IfAUTOFOCUS(cond bool) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *LiHTMLElement) RemoveAUTOFOCUS() *LiHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *LiHTMLElement) SetAUTOFOCUS(b bool) *LiHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *LiHTMLElement) CLASS(v string) *LiHTMLElement {
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

func (e *LiHTMLElement) IfCLASS(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *LiHTMLElement) SetCLASS(v string) *LiHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *LiHTMLElement) RemoveCLASS(v string) *LiHTMLElement {
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
func (e *LiHTMLElement) CONTENTEDITABLE(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *LiHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *LiHTMLElement) RemoveCONTENTEDITABLE(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *LiHTMLElement) DIR(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *LiHTMLElement) IfDIR(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *LiHTMLElement) RemoveDIR(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *LiHTMLElement) DRAGGABLE(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *LiHTMLElement) IfDRAGGABLE(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *LiHTMLElement) RemoveDRAGGABLE(v string) *LiHTMLElement {
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
func (e *LiHTMLElement) ENTERKEYHINT(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *LiHTMLElement) IfENTERKEYHINT(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *LiHTMLElement) RemoveENTERKEYHINT(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *LiHTMLElement) HIDDEN(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *LiHTMLElement) IfHIDDEN(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *LiHTMLElement) RemoveHIDDEN(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *LiHTMLElement) ID(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *LiHTMLElement) IfID(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *LiHTMLElement) RemoveID(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *LiHTMLElement) INERT() *LiHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *LiHTMLElement) IfINERT(cond bool) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *LiHTMLElement) RemoveINERT() *LiHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *LiHTMLElement) SetINERT(b bool) *LiHTMLElement {
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
func (e *LiHTMLElement) INPUTMODE(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *LiHTMLElement) IfINPUTMODE(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *LiHTMLElement) RemoveINPUTMODE(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *LiHTMLElement) IS(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *LiHTMLElement) IfIS(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *LiHTMLElement) RemoveIS(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *LiHTMLElement) ITEMID(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *LiHTMLElement) IfITEMID(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *LiHTMLElement) RemoveITEMID(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *LiHTMLElement) ITEMPROP(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *LiHTMLElement) IfITEMPROP(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *LiHTMLElement) RemoveITEMPROP(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *LiHTMLElement) ITEMREF(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *LiHTMLElement) IfITEMREF(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *LiHTMLElement) RemoveITEMREF(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *LiHTMLElement) ITEMSCOPE() *LiHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *LiHTMLElement) IfITEMSCOPE(cond bool) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *LiHTMLElement) RemoveITEMSCOPE() *LiHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *LiHTMLElement) SetITEMSCOPE(b bool) *LiHTMLElement {
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
func (e *LiHTMLElement) ITEMTYPE(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *LiHTMLElement) IfITEMTYPE(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *LiHTMLElement) RemoveITEMTYPE(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *LiHTMLElement) LANG(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *LiHTMLElement) IfLANG(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *LiHTMLElement) RemoveLANG(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *LiHTMLElement) NONCE(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *LiHTMLElement) IfNONCE(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *LiHTMLElement) RemoveNONCE(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *LiHTMLElement) POPOVER(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *LiHTMLElement) IfPOPOVER(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *LiHTMLElement) RemovePOPOVER(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *LiHTMLElement) SLOT(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *LiHTMLElement) IfSLOT(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *LiHTMLElement) RemoveSLOT(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *LiHTMLElement) SPELLCHECK(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *LiHTMLElement) IfSPELLCHECK(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *LiHTMLElement) RemoveSPELLCHECK(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *LiHTMLElement) STYLE(k, v string) *LiHTMLElement {
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

func (e *LiHTMLElement) IfSTYLE(cond bool, k string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *LiHTMLElement) RemoveSTYLE(k string) *LiHTMLElement {
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
func (e *LiHTMLElement) TABINDEX(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *LiHTMLElement) IfTABINDEX(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *LiHTMLElement) RemoveTABINDEX(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *LiHTMLElement) TITLE(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *LiHTMLElement) IfTITLE(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *LiHTMLElement) RemoveTITLE(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *LiHTMLElement) TRANSLATE(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *LiHTMLElement) IfTRANSLATE(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *LiHTMLElement) RemoveTRANSLATE(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//   - valid_floating_point_number
func (e *LiHTMLElement) VALUE(v string) *LiHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeVALUEKey] = v
	return e
}

func (e *LiHTMLElement) IfVALUE(cond bool, v string) *LiHTMLElement {
	if !cond {
		return e
	}
	return e.VALUE(v)
}

func (e *LiHTMLElement) RemoveVALUE(v string) *LiHTMLElement {
	delete(e.StringAttributes, attributeVALUEKey)
	return e
}
