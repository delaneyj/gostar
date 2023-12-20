package html

import (
	"fmt"
)

type HgroupHTMLElement struct {
	*Element
}

func HGROUP(children ...ElementBuilder) *HgroupHTMLElement {
	return &HgroupHTMLElement{
		Element: &Element{
			Tag:           elementTagHGROUP,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *HgroupHTMLElement) Children(children ...ElementBuilder) *HgroupHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *HgroupHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *HgroupHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *HgroupHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *HgroupHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *HgroupHTMLElement) Text(text string) *HgroupHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *HgroupHTMLElement) TextF(format string, args ...any) *HgroupHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *HgroupHTMLElement) Escaped(text string) *HgroupHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *HgroupHTMLElement) EscapedF(format string, args ...any) *HgroupHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *HgroupHTMLElement) CustomData(key, value string) *HgroupHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *HgroupHTMLElement) CustomDataRemove(key string) *HgroupHTMLElement {
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
func (e *HgroupHTMLElement) ACCESSKEY(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *HgroupHTMLElement) IfACCESSKEY(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *HgroupHTMLElement) RemoveACCESSKEY(v string) *HgroupHTMLElement {
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
func (e *HgroupHTMLElement) AUTOCAPITALIZE(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *HgroupHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *HgroupHTMLElement) RemoveAUTOCAPITALIZE(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *HgroupHTMLElement) AUTOFOCUS() *HgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *HgroupHTMLElement) IfAUTOFOCUS(cond bool) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *HgroupHTMLElement) RemoveAUTOFOCUS() *HgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *HgroupHTMLElement) SetAUTOFOCUS(b bool) *HgroupHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *HgroupHTMLElement) CLASS(v string) *HgroupHTMLElement {
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

func (e *HgroupHTMLElement) IfCLASS(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *HgroupHTMLElement) SetCLASS(v string) *HgroupHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *HgroupHTMLElement) RemoveCLASS(v string) *HgroupHTMLElement {
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
func (e *HgroupHTMLElement) CONTENTEDITABLE(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *HgroupHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *HgroupHTMLElement) RemoveCONTENTEDITABLE(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *HgroupHTMLElement) DIR(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *HgroupHTMLElement) IfDIR(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *HgroupHTMLElement) RemoveDIR(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *HgroupHTMLElement) DRAGGABLE(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *HgroupHTMLElement) IfDRAGGABLE(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *HgroupHTMLElement) RemoveDRAGGABLE(v string) *HgroupHTMLElement {
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
func (e *HgroupHTMLElement) ENTERKEYHINT(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *HgroupHTMLElement) IfENTERKEYHINT(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *HgroupHTMLElement) RemoveENTERKEYHINT(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *HgroupHTMLElement) HIDDEN(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *HgroupHTMLElement) IfHIDDEN(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *HgroupHTMLElement) RemoveHIDDEN(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *HgroupHTMLElement) ID(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *HgroupHTMLElement) IfID(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *HgroupHTMLElement) RemoveID(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *HgroupHTMLElement) INERT() *HgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *HgroupHTMLElement) IfINERT(cond bool) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *HgroupHTMLElement) RemoveINERT() *HgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *HgroupHTMLElement) SetINERT(b bool) *HgroupHTMLElement {
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
func (e *HgroupHTMLElement) INPUTMODE(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *HgroupHTMLElement) IfINPUTMODE(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *HgroupHTMLElement) RemoveINPUTMODE(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *HgroupHTMLElement) IS(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *HgroupHTMLElement) IfIS(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *HgroupHTMLElement) RemoveIS(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *HgroupHTMLElement) ITEMID(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *HgroupHTMLElement) IfITEMID(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *HgroupHTMLElement) RemoveITEMID(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *HgroupHTMLElement) ITEMPROP(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *HgroupHTMLElement) IfITEMPROP(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *HgroupHTMLElement) RemoveITEMPROP(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *HgroupHTMLElement) ITEMREF(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *HgroupHTMLElement) IfITEMREF(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *HgroupHTMLElement) RemoveITEMREF(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *HgroupHTMLElement) ITEMSCOPE() *HgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *HgroupHTMLElement) IfITEMSCOPE(cond bool) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *HgroupHTMLElement) RemoveITEMSCOPE() *HgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *HgroupHTMLElement) SetITEMSCOPE(b bool) *HgroupHTMLElement {
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
func (e *HgroupHTMLElement) ITEMTYPE(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *HgroupHTMLElement) IfITEMTYPE(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *HgroupHTMLElement) RemoveITEMTYPE(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *HgroupHTMLElement) LANG(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *HgroupHTMLElement) IfLANG(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *HgroupHTMLElement) RemoveLANG(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *HgroupHTMLElement) NONCE(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *HgroupHTMLElement) IfNONCE(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *HgroupHTMLElement) RemoveNONCE(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *HgroupHTMLElement) POPOVER(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *HgroupHTMLElement) IfPOPOVER(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *HgroupHTMLElement) RemovePOPOVER(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *HgroupHTMLElement) SLOT(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *HgroupHTMLElement) IfSLOT(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *HgroupHTMLElement) RemoveSLOT(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *HgroupHTMLElement) SPELLCHECK(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *HgroupHTMLElement) IfSPELLCHECK(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *HgroupHTMLElement) RemoveSPELLCHECK(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *HgroupHTMLElement) STYLE(k, v string) *HgroupHTMLElement {
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

func (e *HgroupHTMLElement) IfSTYLE(cond bool, k string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *HgroupHTMLElement) RemoveSTYLE(k string) *HgroupHTMLElement {
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
func (e *HgroupHTMLElement) TABINDEX(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *HgroupHTMLElement) IfTABINDEX(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *HgroupHTMLElement) RemoveTABINDEX(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *HgroupHTMLElement) TITLE(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *HgroupHTMLElement) IfTITLE(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *HgroupHTMLElement) RemoveTITLE(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *HgroupHTMLElement) TRANSLATE(v string) *HgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *HgroupHTMLElement) IfTRANSLATE(cond bool, v string) *HgroupHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *HgroupHTMLElement) RemoveTRANSLATE(v string) *HgroupHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
