package html

import (
	"fmt"
)

type SmallHTMLElement struct {
	*Element
}

func SMALL(children ...ElementBuilder) *SmallHTMLElement {
	return &SmallHTMLElement{
		Element: &Element{
			Tag:           elementTagSMALL,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SmallHTMLElement) Children(children ...ElementBuilder) *SmallHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SmallHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SmallHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SmallHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SmallHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SmallHTMLElement) Text(text string) *SmallHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SmallHTMLElement) TextF(format string, args ...any) *SmallHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SmallHTMLElement) Escaped(text string) *SmallHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SmallHTMLElement) EscapedF(format string, args ...any) *SmallHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SmallHTMLElement) CustomData(key, value string) *SmallHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SmallHTMLElement) CustomDataRemove(key string) *SmallHTMLElement {
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
func (e *SmallHTMLElement) ACCESSKEY(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SmallHTMLElement) IfACCESSKEY(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SmallHTMLElement) RemoveACCESSKEY(v string) *SmallHTMLElement {
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
func (e *SmallHTMLElement) AUTOCAPITALIZE(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SmallHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SmallHTMLElement) RemoveAUTOCAPITALIZE(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SmallHTMLElement) AUTOFOCUS() *SmallHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SmallHTMLElement) IfAUTOFOCUS(cond bool) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SmallHTMLElement) RemoveAUTOFOCUS() *SmallHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SmallHTMLElement) SetAUTOFOCUS(b bool) *SmallHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SmallHTMLElement) CLASS(v string) *SmallHTMLElement {
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

func (e *SmallHTMLElement) IfCLASS(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SmallHTMLElement) SetCLASS(v string) *SmallHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SmallHTMLElement) RemoveCLASS(v string) *SmallHTMLElement {
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
func (e *SmallHTMLElement) CONTENTEDITABLE(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SmallHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SmallHTMLElement) RemoveCONTENTEDITABLE(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *SmallHTMLElement) DIR(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SmallHTMLElement) IfDIR(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SmallHTMLElement) RemoveDIR(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *SmallHTMLElement) DRAGGABLE(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SmallHTMLElement) IfDRAGGABLE(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SmallHTMLElement) RemoveDRAGGABLE(v string) *SmallHTMLElement {
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
func (e *SmallHTMLElement) ENTERKEYHINT(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SmallHTMLElement) IfENTERKEYHINT(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SmallHTMLElement) RemoveENTERKEYHINT(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *SmallHTMLElement) HIDDEN(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SmallHTMLElement) IfHIDDEN(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SmallHTMLElement) RemoveHIDDEN(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SmallHTMLElement) ID(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SmallHTMLElement) IfID(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SmallHTMLElement) RemoveID(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SmallHTMLElement) INERT() *SmallHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SmallHTMLElement) IfINERT(cond bool) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SmallHTMLElement) RemoveINERT() *SmallHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SmallHTMLElement) SetINERT(b bool) *SmallHTMLElement {
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
func (e *SmallHTMLElement) INPUTMODE(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SmallHTMLElement) IfINPUTMODE(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SmallHTMLElement) RemoveINPUTMODE(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *SmallHTMLElement) IS(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SmallHTMLElement) IfIS(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SmallHTMLElement) RemoveIS(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SmallHTMLElement) ITEMID(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SmallHTMLElement) IfITEMID(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SmallHTMLElement) RemoveITEMID(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *SmallHTMLElement) ITEMPROP(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SmallHTMLElement) IfITEMPROP(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SmallHTMLElement) RemoveITEMPROP(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SmallHTMLElement) ITEMREF(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SmallHTMLElement) IfITEMREF(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SmallHTMLElement) RemoveITEMREF(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SmallHTMLElement) ITEMSCOPE() *SmallHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SmallHTMLElement) IfITEMSCOPE(cond bool) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SmallHTMLElement) RemoveITEMSCOPE() *SmallHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SmallHTMLElement) SetITEMSCOPE(b bool) *SmallHTMLElement {
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
func (e *SmallHTMLElement) ITEMTYPE(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SmallHTMLElement) IfITEMTYPE(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SmallHTMLElement) RemoveITEMTYPE(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SmallHTMLElement) LANG(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SmallHTMLElement) IfLANG(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SmallHTMLElement) RemoveLANG(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SmallHTMLElement) NONCE(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SmallHTMLElement) IfNONCE(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SmallHTMLElement) RemoveNONCE(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *SmallHTMLElement) POPOVER(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SmallHTMLElement) IfPOPOVER(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SmallHTMLElement) RemovePOPOVER(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SmallHTMLElement) SLOT(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SmallHTMLElement) IfSLOT(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SmallHTMLElement) RemoveSLOT(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *SmallHTMLElement) SPELLCHECK(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SmallHTMLElement) IfSPELLCHECK(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SmallHTMLElement) RemoveSPELLCHECK(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SmallHTMLElement) STYLE(k, v string) *SmallHTMLElement {
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

func (e *SmallHTMLElement) IfSTYLE(cond bool, k string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SmallHTMLElement) RemoveSTYLE(k string) *SmallHTMLElement {
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
func (e *SmallHTMLElement) TABINDEX(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SmallHTMLElement) IfTABINDEX(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SmallHTMLElement) RemoveTABINDEX(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SmallHTMLElement) TITLE(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SmallHTMLElement) IfTITLE(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SmallHTMLElement) RemoveTITLE(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *SmallHTMLElement) TRANSLATE(v string) *SmallHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SmallHTMLElement) IfTRANSLATE(cond bool, v string) *SmallHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SmallHTMLElement) RemoveTRANSLATE(v string) *SmallHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
