package html

import (
	"fmt"
)

type DlHTMLElement struct {
	*Element
}

func DL(children ...ElementBuilder) *DlHTMLElement {
	return &DlHTMLElement{
		Element: &Element{
			Tag:           elementTagDL,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DlHTMLElement) Children(children ...ElementBuilder) *DlHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DlHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DlHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DlHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DlHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DlHTMLElement) Text(text string) *DlHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DlHTMLElement) TextF(format string, args ...any) *DlHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DlHTMLElement) Escaped(text string) *DlHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DlHTMLElement) EscapedF(format string, args ...any) *DlHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DlHTMLElement) CustomData(key, value string) *DlHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DlHTMLElement) CustomDataRemove(key string) *DlHTMLElement {
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
func (e *DlHTMLElement) ACCESSKEY(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DlHTMLElement) IfACCESSKEY(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DlHTMLElement) RemoveACCESSKEY(v string) *DlHTMLElement {
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
func (e *DlHTMLElement) AUTOCAPITALIZE(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DlHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DlHTMLElement) RemoveAUTOCAPITALIZE(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DlHTMLElement) AUTOFOCUS() *DlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DlHTMLElement) IfAUTOFOCUS(cond bool) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DlHTMLElement) RemoveAUTOFOCUS() *DlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DlHTMLElement) SetAUTOFOCUS(b bool) *DlHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DlHTMLElement) CLASS(v string) *DlHTMLElement {
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

func (e *DlHTMLElement) IfCLASS(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DlHTMLElement) SetCLASS(v string) *DlHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DlHTMLElement) RemoveCLASS(v string) *DlHTMLElement {
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
func (e *DlHTMLElement) CONTENTEDITABLE(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DlHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DlHTMLElement) RemoveCONTENTEDITABLE(v string) *DlHTMLElement {
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
func (e *DlHTMLElement) DIR(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DlHTMLElement) IfDIR(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DlHTMLElement) RemoveDIR(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *DlHTMLElement) DRAGGABLE(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DlHTMLElement) IfDRAGGABLE(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DlHTMLElement) RemoveDRAGGABLE(v string) *DlHTMLElement {
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
func (e *DlHTMLElement) ENTERKEYHINT(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DlHTMLElement) IfENTERKEYHINT(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DlHTMLElement) RemoveENTERKEYHINT(v string) *DlHTMLElement {
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
func (e *DlHTMLElement) HIDDEN(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DlHTMLElement) IfHIDDEN(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DlHTMLElement) RemoveHIDDEN(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DlHTMLElement) ID(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DlHTMLElement) IfID(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DlHTMLElement) RemoveID(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DlHTMLElement) INERT() *DlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DlHTMLElement) IfINERT(cond bool) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DlHTMLElement) RemoveINERT() *DlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DlHTMLElement) SetINERT(b bool) *DlHTMLElement {
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
func (e *DlHTMLElement) INPUTMODE(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DlHTMLElement) IfINPUTMODE(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DlHTMLElement) RemoveINPUTMODE(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *DlHTMLElement) IS(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DlHTMLElement) IfIS(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DlHTMLElement) RemoveIS(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DlHTMLElement) ITEMID(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DlHTMLElement) IfITEMID(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DlHTMLElement) RemoveITEMID(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *DlHTMLElement) ITEMPROP(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DlHTMLElement) IfITEMPROP(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DlHTMLElement) RemoveITEMPROP(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DlHTMLElement) ITEMREF(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DlHTMLElement) IfITEMREF(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DlHTMLElement) RemoveITEMREF(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DlHTMLElement) ITEMSCOPE() *DlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DlHTMLElement) IfITEMSCOPE(cond bool) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DlHTMLElement) RemoveITEMSCOPE() *DlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DlHTMLElement) SetITEMSCOPE(b bool) *DlHTMLElement {
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
func (e *DlHTMLElement) ITEMTYPE(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DlHTMLElement) IfITEMTYPE(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DlHTMLElement) RemoveITEMTYPE(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DlHTMLElement) LANG(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DlHTMLElement) IfLANG(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DlHTMLElement) RemoveLANG(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DlHTMLElement) NONCE(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DlHTMLElement) IfNONCE(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DlHTMLElement) RemoveNONCE(v string) *DlHTMLElement {
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
func (e *DlHTMLElement) POPOVER(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DlHTMLElement) IfPOPOVER(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DlHTMLElement) RemovePOPOVER(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DlHTMLElement) SLOT(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DlHTMLElement) IfSLOT(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DlHTMLElement) RemoveSLOT(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *DlHTMLElement) SPELLCHECK(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DlHTMLElement) IfSPELLCHECK(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DlHTMLElement) RemoveSPELLCHECK(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DlHTMLElement) STYLE(k, v string) *DlHTMLElement {
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

func (e *DlHTMLElement) IfSTYLE(cond bool, k string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DlHTMLElement) RemoveSTYLE(k string) *DlHTMLElement {
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
func (e *DlHTMLElement) TABINDEX(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DlHTMLElement) IfTABINDEX(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DlHTMLElement) RemoveTABINDEX(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DlHTMLElement) TITLE(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DlHTMLElement) IfTITLE(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DlHTMLElement) RemoveTITLE(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *DlHTMLElement) TRANSLATE(v string) *DlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DlHTMLElement) IfTRANSLATE(cond bool, v string) *DlHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DlHTMLElement) RemoveTRANSLATE(v string) *DlHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
