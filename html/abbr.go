package html

import (
	"fmt"
)

type AbbrHTMLElement struct {
	*Element
}

func ABBR(children ...ElementBuilder) *AbbrHTMLElement {
	return &AbbrHTMLElement{
		Element: &Element{
			Tag:           elementTagABBR,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *AbbrHTMLElement) Children(children ...ElementBuilder) *AbbrHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *AbbrHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *AbbrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *AbbrHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *AbbrHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *AbbrHTMLElement) Text(text string) *AbbrHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *AbbrHTMLElement) TextF(format string, args ...any) *AbbrHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *AbbrHTMLElement) Escaped(text string) *AbbrHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *AbbrHTMLElement) EscapedF(format string, args ...any) *AbbrHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *AbbrHTMLElement) CustomData(key, value string) *AbbrHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AbbrHTMLElement) CustomDataRemove(key string) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) ACCESSKEY(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *AbbrHTMLElement) IfACCESSKEY(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *AbbrHTMLElement) RemoveACCESSKEY(v string) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) AUTOCAPITALIZE(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *AbbrHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *AbbrHTMLElement) RemoveAUTOCAPITALIZE(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *AbbrHTMLElement) AUTOFOCUS() *AbbrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *AbbrHTMLElement) IfAUTOFOCUS(cond bool) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *AbbrHTMLElement) RemoveAUTOFOCUS() *AbbrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *AbbrHTMLElement) SetAUTOFOCUS(b bool) *AbbrHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *AbbrHTMLElement) CLASS(v string) *AbbrHTMLElement {
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

func (e *AbbrHTMLElement) IfCLASS(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *AbbrHTMLElement) SetCLASS(v string) *AbbrHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *AbbrHTMLElement) RemoveCLASS(v string) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) CONTENTEDITABLE(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *AbbrHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *AbbrHTMLElement) RemoveCONTENTEDITABLE(v string) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) DIR(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *AbbrHTMLElement) IfDIR(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *AbbrHTMLElement) RemoveDIR(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *AbbrHTMLElement) DRAGGABLE(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *AbbrHTMLElement) IfDRAGGABLE(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *AbbrHTMLElement) RemoveDRAGGABLE(v string) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) ENTERKEYHINT(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *AbbrHTMLElement) IfENTERKEYHINT(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *AbbrHTMLElement) RemoveENTERKEYHINT(v string) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) HIDDEN(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *AbbrHTMLElement) IfHIDDEN(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *AbbrHTMLElement) RemoveHIDDEN(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *AbbrHTMLElement) ID(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *AbbrHTMLElement) IfID(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *AbbrHTMLElement) RemoveID(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *AbbrHTMLElement) INERT() *AbbrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *AbbrHTMLElement) IfINERT(cond bool) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *AbbrHTMLElement) RemoveINERT() *AbbrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *AbbrHTMLElement) SetINERT(b bool) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) INPUTMODE(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *AbbrHTMLElement) IfINPUTMODE(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *AbbrHTMLElement) RemoveINPUTMODE(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *AbbrHTMLElement) IS(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *AbbrHTMLElement) IfIS(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *AbbrHTMLElement) RemoveIS(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *AbbrHTMLElement) ITEMID(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *AbbrHTMLElement) IfITEMID(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *AbbrHTMLElement) RemoveITEMID(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *AbbrHTMLElement) ITEMPROP(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *AbbrHTMLElement) IfITEMPROP(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *AbbrHTMLElement) RemoveITEMPROP(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *AbbrHTMLElement) ITEMREF(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *AbbrHTMLElement) IfITEMREF(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *AbbrHTMLElement) RemoveITEMREF(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *AbbrHTMLElement) ITEMSCOPE() *AbbrHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *AbbrHTMLElement) IfITEMSCOPE(cond bool) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *AbbrHTMLElement) RemoveITEMSCOPE() *AbbrHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *AbbrHTMLElement) SetITEMSCOPE(b bool) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) ITEMTYPE(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *AbbrHTMLElement) IfITEMTYPE(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *AbbrHTMLElement) RemoveITEMTYPE(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AbbrHTMLElement) LANG(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *AbbrHTMLElement) IfLANG(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *AbbrHTMLElement) RemoveLANG(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *AbbrHTMLElement) NONCE(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *AbbrHTMLElement) IfNONCE(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *AbbrHTMLElement) RemoveNONCE(v string) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) POPOVER(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *AbbrHTMLElement) IfPOPOVER(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *AbbrHTMLElement) RemovePOPOVER(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *AbbrHTMLElement) SLOT(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *AbbrHTMLElement) IfSLOT(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *AbbrHTMLElement) RemoveSLOT(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *AbbrHTMLElement) SPELLCHECK(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *AbbrHTMLElement) IfSPELLCHECK(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *AbbrHTMLElement) RemoveSPELLCHECK(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AbbrHTMLElement) STYLE(k, v string) *AbbrHTMLElement {
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

func (e *AbbrHTMLElement) IfSTYLE(cond bool, k string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *AbbrHTMLElement) RemoveSTYLE(k string) *AbbrHTMLElement {
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
func (e *AbbrHTMLElement) TABINDEX(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *AbbrHTMLElement) IfTABINDEX(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *AbbrHTMLElement) RemoveTABINDEX(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *AbbrHTMLElement) TITLE(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *AbbrHTMLElement) IfTITLE(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *AbbrHTMLElement) RemoveTITLE(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *AbbrHTMLElement) TRANSLATE(v string) *AbbrHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *AbbrHTMLElement) IfTRANSLATE(cond bool, v string) *AbbrHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *AbbrHTMLElement) RemoveTRANSLATE(v string) *AbbrHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
