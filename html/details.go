package html

import (
	"fmt"
)

type DetailsHTMLElement struct {
	*Element
}

func DETAILS(children ...ElementBuilder) *DetailsHTMLElement {
	return &DetailsHTMLElement{
		Element: &Element{
			Tag:           elementTagDETAILS,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DetailsHTMLElement) Children(children ...ElementBuilder) *DetailsHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DetailsHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DetailsHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DetailsHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DetailsHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DetailsHTMLElement) Text(text string) *DetailsHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DetailsHTMLElement) TextF(format string, args ...any) *DetailsHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DetailsHTMLElement) Escaped(text string) *DetailsHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DetailsHTMLElement) EscapedF(format string, args ...any) *DetailsHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DetailsHTMLElement) CustomData(key, value string) *DetailsHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DetailsHTMLElement) CustomDataRemove(key string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) ACCESSKEY(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DetailsHTMLElement) IfACCESSKEY(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DetailsHTMLElement) RemoveACCESSKEY(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) AUTOCAPITALIZE(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DetailsHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DetailsHTMLElement) RemoveAUTOCAPITALIZE(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DetailsHTMLElement) AUTOFOCUS() *DetailsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DetailsHTMLElement) IfAUTOFOCUS(cond bool) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DetailsHTMLElement) RemoveAUTOFOCUS() *DetailsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DetailsHTMLElement) SetAUTOFOCUS(b bool) *DetailsHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DetailsHTMLElement) CLASS(v string) *DetailsHTMLElement {
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

func (e *DetailsHTMLElement) IfCLASS(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DetailsHTMLElement) SetCLASS(v string) *DetailsHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DetailsHTMLElement) RemoveCLASS(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) CONTENTEDITABLE(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DetailsHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DetailsHTMLElement) RemoveCONTENTEDITABLE(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) DIR(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DetailsHTMLElement) IfDIR(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DetailsHTMLElement) RemoveDIR(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *DetailsHTMLElement) DRAGGABLE(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DetailsHTMLElement) IfDRAGGABLE(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DetailsHTMLElement) RemoveDRAGGABLE(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) ENTERKEYHINT(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DetailsHTMLElement) IfENTERKEYHINT(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DetailsHTMLElement) RemoveENTERKEYHINT(v string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) HIDDEN(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DetailsHTMLElement) IfHIDDEN(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DetailsHTMLElement) RemoveHIDDEN(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DetailsHTMLElement) ID(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DetailsHTMLElement) IfID(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DetailsHTMLElement) RemoveID(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DetailsHTMLElement) INERT() *DetailsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DetailsHTMLElement) IfINERT(cond bool) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DetailsHTMLElement) RemoveINERT() *DetailsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DetailsHTMLElement) SetINERT(b bool) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) INPUTMODE(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DetailsHTMLElement) IfINPUTMODE(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DetailsHTMLElement) RemoveINPUTMODE(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *DetailsHTMLElement) IS(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DetailsHTMLElement) IfIS(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DetailsHTMLElement) RemoveIS(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DetailsHTMLElement) ITEMID(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DetailsHTMLElement) IfITEMID(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DetailsHTMLElement) RemoveITEMID(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *DetailsHTMLElement) ITEMPROP(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DetailsHTMLElement) IfITEMPROP(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DetailsHTMLElement) RemoveITEMPROP(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DetailsHTMLElement) ITEMREF(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DetailsHTMLElement) IfITEMREF(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DetailsHTMLElement) RemoveITEMREF(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DetailsHTMLElement) ITEMSCOPE() *DetailsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DetailsHTMLElement) IfITEMSCOPE(cond bool) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DetailsHTMLElement) RemoveITEMSCOPE() *DetailsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DetailsHTMLElement) SetITEMSCOPE(b bool) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) ITEMTYPE(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DetailsHTMLElement) IfITEMTYPE(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DetailsHTMLElement) RemoveITEMTYPE(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DetailsHTMLElement) LANG(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DetailsHTMLElement) IfLANG(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DetailsHTMLElement) RemoveLANG(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *DetailsHTMLElement) NAME(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *DetailsHTMLElement) IfNAME(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *DetailsHTMLElement) RemoveNAME(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DetailsHTMLElement) NONCE(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DetailsHTMLElement) IfNONCE(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DetailsHTMLElement) RemoveNONCE(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// OPEN sets the "open" attribute.
// Whether the dialog box is showing
// Values values are constrained to:
//   - boolean_attribute
func (e *DetailsHTMLElement) OPEN() *DetailsHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeOPENKey] = struct{}{}
	return e
}

func (e *DetailsHTMLElement) IfOPEN(cond bool) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.OPEN()
}

func (e *DetailsHTMLElement) RemoveOPEN() *DetailsHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeOPENKey)
	return e
}

func (e *DetailsHTMLElement) SetOPEN(b bool) *DetailsHTMLElement {
	if b {
		return e.OPEN()
	}
	return e.RemoveOPEN()
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - auto
//   - manual
//   - manual
func (e *DetailsHTMLElement) POPOVER(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DetailsHTMLElement) IfPOPOVER(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DetailsHTMLElement) RemovePOPOVER(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DetailsHTMLElement) SLOT(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DetailsHTMLElement) IfSLOT(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DetailsHTMLElement) RemoveSLOT(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *DetailsHTMLElement) SPELLCHECK(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DetailsHTMLElement) IfSPELLCHECK(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DetailsHTMLElement) RemoveSPELLCHECK(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DetailsHTMLElement) STYLE(k, v string) *DetailsHTMLElement {
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

func (e *DetailsHTMLElement) IfSTYLE(cond bool, k string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DetailsHTMLElement) RemoveSTYLE(k string) *DetailsHTMLElement {
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
func (e *DetailsHTMLElement) TABINDEX(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DetailsHTMLElement) IfTABINDEX(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DetailsHTMLElement) RemoveTABINDEX(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DetailsHTMLElement) TITLE(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DetailsHTMLElement) IfTITLE(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DetailsHTMLElement) RemoveTITLE(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *DetailsHTMLElement) TRANSLATE(v string) *DetailsHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DetailsHTMLElement) IfTRANSLATE(cond bool, v string) *DetailsHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DetailsHTMLElement) RemoveTRANSLATE(v string) *DetailsHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
