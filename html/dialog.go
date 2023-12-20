package html

import (
	"fmt"
)

type DialogHTMLElement struct {
	*Element
}

func DIALOG(children ...ElementBuilder) *DialogHTMLElement {
	return &DialogHTMLElement{
		Element: &Element{
			Tag:           elementTagDIALOG,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DialogHTMLElement) Children(children ...ElementBuilder) *DialogHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DialogHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DialogHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DialogHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DialogHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DialogHTMLElement) Text(text string) *DialogHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DialogHTMLElement) TextF(format string, args ...any) *DialogHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DialogHTMLElement) Escaped(text string) *DialogHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DialogHTMLElement) EscapedF(format string, args ...any) *DialogHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DialogHTMLElement) CustomData(key, value string) *DialogHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DialogHTMLElement) CustomDataRemove(key string) *DialogHTMLElement {
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
func (e *DialogHTMLElement) ACCESSKEY(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DialogHTMLElement) IfACCESSKEY(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DialogHTMLElement) RemoveACCESSKEY(v string) *DialogHTMLElement {
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
func (e *DialogHTMLElement) AUTOCAPITALIZE(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DialogHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DialogHTMLElement) RemoveAUTOCAPITALIZE(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DialogHTMLElement) AUTOFOCUS() *DialogHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DialogHTMLElement) IfAUTOFOCUS(cond bool) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DialogHTMLElement) RemoveAUTOFOCUS() *DialogHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DialogHTMLElement) SetAUTOFOCUS(b bool) *DialogHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DialogHTMLElement) CLASS(v string) *DialogHTMLElement {
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

func (e *DialogHTMLElement) IfCLASS(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DialogHTMLElement) SetCLASS(v string) *DialogHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DialogHTMLElement) RemoveCLASS(v string) *DialogHTMLElement {
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
func (e *DialogHTMLElement) CONTENTEDITABLE(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DialogHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DialogHTMLElement) RemoveCONTENTEDITABLE(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *DialogHTMLElement) DIR(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DialogHTMLElement) IfDIR(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DialogHTMLElement) RemoveDIR(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *DialogHTMLElement) DRAGGABLE(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DialogHTMLElement) IfDRAGGABLE(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DialogHTMLElement) RemoveDRAGGABLE(v string) *DialogHTMLElement {
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
func (e *DialogHTMLElement) ENTERKEYHINT(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DialogHTMLElement) IfENTERKEYHINT(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DialogHTMLElement) RemoveENTERKEYHINT(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *DialogHTMLElement) HIDDEN(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DialogHTMLElement) IfHIDDEN(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DialogHTMLElement) RemoveHIDDEN(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DialogHTMLElement) ID(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DialogHTMLElement) IfID(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DialogHTMLElement) RemoveID(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DialogHTMLElement) INERT() *DialogHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DialogHTMLElement) IfINERT(cond bool) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DialogHTMLElement) RemoveINERT() *DialogHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DialogHTMLElement) SetINERT(b bool) *DialogHTMLElement {
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
func (e *DialogHTMLElement) INPUTMODE(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DialogHTMLElement) IfINPUTMODE(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DialogHTMLElement) RemoveINPUTMODE(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *DialogHTMLElement) IS(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DialogHTMLElement) IfIS(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DialogHTMLElement) RemoveIS(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DialogHTMLElement) ITEMID(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DialogHTMLElement) IfITEMID(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DialogHTMLElement) RemoveITEMID(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *DialogHTMLElement) ITEMPROP(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DialogHTMLElement) IfITEMPROP(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DialogHTMLElement) RemoveITEMPROP(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DialogHTMLElement) ITEMREF(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DialogHTMLElement) IfITEMREF(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DialogHTMLElement) RemoveITEMREF(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DialogHTMLElement) ITEMSCOPE() *DialogHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DialogHTMLElement) IfITEMSCOPE(cond bool) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DialogHTMLElement) RemoveITEMSCOPE() *DialogHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DialogHTMLElement) SetITEMSCOPE(b bool) *DialogHTMLElement {
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
func (e *DialogHTMLElement) ITEMTYPE(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DialogHTMLElement) IfITEMTYPE(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DialogHTMLElement) RemoveITEMTYPE(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DialogHTMLElement) LANG(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DialogHTMLElement) IfLANG(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DialogHTMLElement) RemoveLANG(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DialogHTMLElement) NONCE(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DialogHTMLElement) IfNONCE(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DialogHTMLElement) RemoveNONCE(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// OPEN sets the "open" attribute.
// Whether the dialog box is showing
// Values values are constrained to:
//   - boolean_attribute
func (e *DialogHTMLElement) OPEN() *DialogHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeOPENKey] = struct{}{}
	return e
}

func (e *DialogHTMLElement) IfOPEN(cond bool) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.OPEN()
}

func (e *DialogHTMLElement) RemoveOPEN() *DialogHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeOPENKey)
	return e
}

func (e *DialogHTMLElement) SetOPEN(b bool) *DialogHTMLElement {
	if b {
		return e.OPEN()
	}
	return e.RemoveOPEN()
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *DialogHTMLElement) POPOVER(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DialogHTMLElement) IfPOPOVER(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DialogHTMLElement) RemovePOPOVER(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DialogHTMLElement) SLOT(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DialogHTMLElement) IfSLOT(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DialogHTMLElement) RemoveSLOT(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *DialogHTMLElement) SPELLCHECK(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DialogHTMLElement) IfSPELLCHECK(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DialogHTMLElement) RemoveSPELLCHECK(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DialogHTMLElement) STYLE(k, v string) *DialogHTMLElement {
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

func (e *DialogHTMLElement) IfSTYLE(cond bool, k string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DialogHTMLElement) RemoveSTYLE(k string) *DialogHTMLElement {
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
func (e *DialogHTMLElement) TABINDEX(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DialogHTMLElement) IfTABINDEX(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DialogHTMLElement) RemoveTABINDEX(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DialogHTMLElement) TITLE(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DialogHTMLElement) IfTITLE(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DialogHTMLElement) RemoveTITLE(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *DialogHTMLElement) TRANSLATE(v string) *DialogHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DialogHTMLElement) IfTRANSLATE(cond bool, v string) *DialogHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DialogHTMLElement) RemoveTRANSLATE(v string) *DialogHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
