package html

import (
	"fmt"
)

type RubyHTMLElement struct {
	*Element
}

func RUBY(children ...ElementBuilder) *RubyHTMLElement {
	return &RubyHTMLElement{
		Element: &Element{
			Tag:           elementTagRUBY,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *RubyHTMLElement) Children(children ...ElementBuilder) *RubyHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *RubyHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *RubyHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *RubyHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *RubyHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *RubyHTMLElement) Text(text string) *RubyHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *RubyHTMLElement) TextF(format string, args ...any) *RubyHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *RubyHTMLElement) Escaped(text string) *RubyHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *RubyHTMLElement) EscapedF(format string, args ...any) *RubyHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *RubyHTMLElement) CustomData(key, value string) *RubyHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *RubyHTMLElement) CustomDataRemove(key string) *RubyHTMLElement {
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
func (e *RubyHTMLElement) ACCESSKEY(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *RubyHTMLElement) IfACCESSKEY(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *RubyHTMLElement) RemoveACCESSKEY(v string) *RubyHTMLElement {
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
func (e *RubyHTMLElement) AUTOCAPITALIZE(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *RubyHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *RubyHTMLElement) RemoveAUTOCAPITALIZE(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *RubyHTMLElement) AUTOFOCUS() *RubyHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *RubyHTMLElement) IfAUTOFOCUS(cond bool) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *RubyHTMLElement) RemoveAUTOFOCUS() *RubyHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *RubyHTMLElement) SetAUTOFOCUS(b bool) *RubyHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *RubyHTMLElement) CLASS(v string) *RubyHTMLElement {
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

func (e *RubyHTMLElement) IfCLASS(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *RubyHTMLElement) SetCLASS(v string) *RubyHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *RubyHTMLElement) RemoveCLASS(v string) *RubyHTMLElement {
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
func (e *RubyHTMLElement) CONTENTEDITABLE(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *RubyHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *RubyHTMLElement) RemoveCONTENTEDITABLE(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *RubyHTMLElement) DIR(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *RubyHTMLElement) IfDIR(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *RubyHTMLElement) RemoveDIR(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *RubyHTMLElement) DRAGGABLE(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *RubyHTMLElement) IfDRAGGABLE(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *RubyHTMLElement) RemoveDRAGGABLE(v string) *RubyHTMLElement {
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
func (e *RubyHTMLElement) ENTERKEYHINT(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *RubyHTMLElement) IfENTERKEYHINT(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *RubyHTMLElement) RemoveENTERKEYHINT(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *RubyHTMLElement) HIDDEN(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *RubyHTMLElement) IfHIDDEN(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *RubyHTMLElement) RemoveHIDDEN(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *RubyHTMLElement) ID(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *RubyHTMLElement) IfID(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *RubyHTMLElement) RemoveID(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *RubyHTMLElement) INERT() *RubyHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *RubyHTMLElement) IfINERT(cond bool) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *RubyHTMLElement) RemoveINERT() *RubyHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *RubyHTMLElement) SetINERT(b bool) *RubyHTMLElement {
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
func (e *RubyHTMLElement) INPUTMODE(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *RubyHTMLElement) IfINPUTMODE(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *RubyHTMLElement) RemoveINPUTMODE(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *RubyHTMLElement) IS(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *RubyHTMLElement) IfIS(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *RubyHTMLElement) RemoveIS(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *RubyHTMLElement) ITEMID(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *RubyHTMLElement) IfITEMID(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *RubyHTMLElement) RemoveITEMID(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *RubyHTMLElement) ITEMPROP(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *RubyHTMLElement) IfITEMPROP(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *RubyHTMLElement) RemoveITEMPROP(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *RubyHTMLElement) ITEMREF(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *RubyHTMLElement) IfITEMREF(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *RubyHTMLElement) RemoveITEMREF(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *RubyHTMLElement) ITEMSCOPE() *RubyHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *RubyHTMLElement) IfITEMSCOPE(cond bool) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *RubyHTMLElement) RemoveITEMSCOPE() *RubyHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *RubyHTMLElement) SetITEMSCOPE(b bool) *RubyHTMLElement {
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
func (e *RubyHTMLElement) ITEMTYPE(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *RubyHTMLElement) IfITEMTYPE(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *RubyHTMLElement) RemoveITEMTYPE(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *RubyHTMLElement) LANG(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *RubyHTMLElement) IfLANG(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *RubyHTMLElement) RemoveLANG(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *RubyHTMLElement) NONCE(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *RubyHTMLElement) IfNONCE(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *RubyHTMLElement) RemoveNONCE(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *RubyHTMLElement) POPOVER(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *RubyHTMLElement) IfPOPOVER(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *RubyHTMLElement) RemovePOPOVER(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *RubyHTMLElement) SLOT(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *RubyHTMLElement) IfSLOT(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *RubyHTMLElement) RemoveSLOT(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *RubyHTMLElement) SPELLCHECK(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *RubyHTMLElement) IfSPELLCHECK(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *RubyHTMLElement) RemoveSPELLCHECK(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *RubyHTMLElement) STYLE(k, v string) *RubyHTMLElement {
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

func (e *RubyHTMLElement) IfSTYLE(cond bool, k string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *RubyHTMLElement) RemoveSTYLE(k string) *RubyHTMLElement {
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
func (e *RubyHTMLElement) TABINDEX(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *RubyHTMLElement) IfTABINDEX(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *RubyHTMLElement) RemoveTABINDEX(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *RubyHTMLElement) TITLE(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *RubyHTMLElement) IfTITLE(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *RubyHTMLElement) RemoveTITLE(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *RubyHTMLElement) TRANSLATE(v string) *RubyHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *RubyHTMLElement) IfTRANSLATE(cond bool, v string) *RubyHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *RubyHTMLElement) RemoveTRANSLATE(v string) *RubyHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
