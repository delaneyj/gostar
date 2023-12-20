package html

import (
	"fmt"
)

type AsideHTMLElement struct {
	*Element
}

func ASIDE(children ...ElementBuilder) *AsideHTMLElement {
	return &AsideHTMLElement{
		Element: &Element{
			Tag:           elementTagASIDE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *AsideHTMLElement) Children(children ...ElementBuilder) *AsideHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *AsideHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *AsideHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *AsideHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *AsideHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *AsideHTMLElement) Text(text string) *AsideHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *AsideHTMLElement) TextF(format string, args ...any) *AsideHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *AsideHTMLElement) Escaped(text string) *AsideHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *AsideHTMLElement) EscapedF(format string, args ...any) *AsideHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *AsideHTMLElement) CustomData(key, value string) *AsideHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AsideHTMLElement) CustomDataRemove(key string) *AsideHTMLElement {
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
func (e *AsideHTMLElement) ACCESSKEY(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *AsideHTMLElement) IfACCESSKEY(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *AsideHTMLElement) RemoveACCESSKEY(v string) *AsideHTMLElement {
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
func (e *AsideHTMLElement) AUTOCAPITALIZE(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *AsideHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *AsideHTMLElement) RemoveAUTOCAPITALIZE(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *AsideHTMLElement) AUTOFOCUS() *AsideHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *AsideHTMLElement) IfAUTOFOCUS(cond bool) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *AsideHTMLElement) RemoveAUTOFOCUS() *AsideHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *AsideHTMLElement) SetAUTOFOCUS(b bool) *AsideHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *AsideHTMLElement) CLASS(v string) *AsideHTMLElement {
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

func (e *AsideHTMLElement) IfCLASS(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *AsideHTMLElement) SetCLASS(v string) *AsideHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *AsideHTMLElement) RemoveCLASS(v string) *AsideHTMLElement {
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
func (e *AsideHTMLElement) CONTENTEDITABLE(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *AsideHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *AsideHTMLElement) RemoveCONTENTEDITABLE(v string) *AsideHTMLElement {
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
func (e *AsideHTMLElement) DIR(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *AsideHTMLElement) IfDIR(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *AsideHTMLElement) RemoveDIR(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *AsideHTMLElement) DRAGGABLE(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *AsideHTMLElement) IfDRAGGABLE(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *AsideHTMLElement) RemoveDRAGGABLE(v string) *AsideHTMLElement {
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
func (e *AsideHTMLElement) ENTERKEYHINT(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *AsideHTMLElement) IfENTERKEYHINT(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *AsideHTMLElement) RemoveENTERKEYHINT(v string) *AsideHTMLElement {
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
func (e *AsideHTMLElement) HIDDEN(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *AsideHTMLElement) IfHIDDEN(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *AsideHTMLElement) RemoveHIDDEN(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *AsideHTMLElement) ID(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *AsideHTMLElement) IfID(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *AsideHTMLElement) RemoveID(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *AsideHTMLElement) INERT() *AsideHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *AsideHTMLElement) IfINERT(cond bool) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *AsideHTMLElement) RemoveINERT() *AsideHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *AsideHTMLElement) SetINERT(b bool) *AsideHTMLElement {
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
func (e *AsideHTMLElement) INPUTMODE(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *AsideHTMLElement) IfINPUTMODE(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *AsideHTMLElement) RemoveINPUTMODE(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *AsideHTMLElement) IS(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *AsideHTMLElement) IfIS(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *AsideHTMLElement) RemoveIS(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *AsideHTMLElement) ITEMID(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *AsideHTMLElement) IfITEMID(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *AsideHTMLElement) RemoveITEMID(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *AsideHTMLElement) ITEMPROP(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *AsideHTMLElement) IfITEMPROP(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *AsideHTMLElement) RemoveITEMPROP(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *AsideHTMLElement) ITEMREF(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *AsideHTMLElement) IfITEMREF(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *AsideHTMLElement) RemoveITEMREF(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *AsideHTMLElement) ITEMSCOPE() *AsideHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *AsideHTMLElement) IfITEMSCOPE(cond bool) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *AsideHTMLElement) RemoveITEMSCOPE() *AsideHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *AsideHTMLElement) SetITEMSCOPE(b bool) *AsideHTMLElement {
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
func (e *AsideHTMLElement) ITEMTYPE(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *AsideHTMLElement) IfITEMTYPE(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *AsideHTMLElement) RemoveITEMTYPE(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AsideHTMLElement) LANG(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *AsideHTMLElement) IfLANG(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *AsideHTMLElement) RemoveLANG(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *AsideHTMLElement) NONCE(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *AsideHTMLElement) IfNONCE(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *AsideHTMLElement) RemoveNONCE(v string) *AsideHTMLElement {
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
func (e *AsideHTMLElement) POPOVER(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *AsideHTMLElement) IfPOPOVER(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *AsideHTMLElement) RemovePOPOVER(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *AsideHTMLElement) SLOT(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *AsideHTMLElement) IfSLOT(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *AsideHTMLElement) RemoveSLOT(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *AsideHTMLElement) SPELLCHECK(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *AsideHTMLElement) IfSPELLCHECK(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *AsideHTMLElement) RemoveSPELLCHECK(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AsideHTMLElement) STYLE(k, v string) *AsideHTMLElement {
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

func (e *AsideHTMLElement) IfSTYLE(cond bool, k string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *AsideHTMLElement) RemoveSTYLE(k string) *AsideHTMLElement {
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
func (e *AsideHTMLElement) TABINDEX(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *AsideHTMLElement) IfTABINDEX(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *AsideHTMLElement) RemoveTABINDEX(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *AsideHTMLElement) TITLE(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *AsideHTMLElement) IfTITLE(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *AsideHTMLElement) RemoveTITLE(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *AsideHTMLElement) TRANSLATE(v string) *AsideHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *AsideHTMLElement) IfTRANSLATE(cond bool, v string) *AsideHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *AsideHTMLElement) RemoveTRANSLATE(v string) *AsideHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
