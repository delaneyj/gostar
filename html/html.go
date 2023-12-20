package html

import (
	"fmt"
)

type HtmlHTMLElement struct {
	*Element
}

func HTML(children ...ElementBuilder) *HtmlHTMLElement {
	return &HtmlHTMLElement{
		Element: &Element{
			Tag:           elementTagHTML,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *HtmlHTMLElement) Children(children ...ElementBuilder) *HtmlHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *HtmlHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *HtmlHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *HtmlHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *HtmlHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *HtmlHTMLElement) Text(text string) *HtmlHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *HtmlHTMLElement) TextF(format string, args ...any) *HtmlHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *HtmlHTMLElement) Escaped(text string) *HtmlHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *HtmlHTMLElement) EscapedF(format string, args ...any) *HtmlHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *HtmlHTMLElement) CustomData(key, value string) *HtmlHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *HtmlHTMLElement) CustomDataRemove(key string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) ACCESSKEY(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *HtmlHTMLElement) IfACCESSKEY(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *HtmlHTMLElement) RemoveACCESSKEY(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) AUTOCAPITALIZE(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *HtmlHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *HtmlHTMLElement) RemoveAUTOCAPITALIZE(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *HtmlHTMLElement) AUTOFOCUS() *HtmlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *HtmlHTMLElement) IfAUTOFOCUS(cond bool) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *HtmlHTMLElement) RemoveAUTOFOCUS() *HtmlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *HtmlHTMLElement) SetAUTOFOCUS(b bool) *HtmlHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *HtmlHTMLElement) CLASS(v string) *HtmlHTMLElement {
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

func (e *HtmlHTMLElement) IfCLASS(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *HtmlHTMLElement) SetCLASS(v string) *HtmlHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *HtmlHTMLElement) RemoveCLASS(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) CONTENTEDITABLE(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *HtmlHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *HtmlHTMLElement) RemoveCONTENTEDITABLE(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) DIR(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *HtmlHTMLElement) IfDIR(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *HtmlHTMLElement) RemoveDIR(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *HtmlHTMLElement) DRAGGABLE(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *HtmlHTMLElement) IfDRAGGABLE(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *HtmlHTMLElement) RemoveDRAGGABLE(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) ENTERKEYHINT(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *HtmlHTMLElement) IfENTERKEYHINT(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *HtmlHTMLElement) RemoveENTERKEYHINT(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) HIDDEN(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *HtmlHTMLElement) IfHIDDEN(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *HtmlHTMLElement) RemoveHIDDEN(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *HtmlHTMLElement) ID(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *HtmlHTMLElement) IfID(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *HtmlHTMLElement) RemoveID(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *HtmlHTMLElement) INERT() *HtmlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *HtmlHTMLElement) IfINERT(cond bool) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *HtmlHTMLElement) RemoveINERT() *HtmlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *HtmlHTMLElement) SetINERT(b bool) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) INPUTMODE(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *HtmlHTMLElement) IfINPUTMODE(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *HtmlHTMLElement) RemoveINPUTMODE(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *HtmlHTMLElement) IS(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *HtmlHTMLElement) IfIS(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *HtmlHTMLElement) RemoveIS(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *HtmlHTMLElement) ITEMID(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *HtmlHTMLElement) IfITEMID(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *HtmlHTMLElement) RemoveITEMID(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *HtmlHTMLElement) ITEMPROP(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *HtmlHTMLElement) IfITEMPROP(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *HtmlHTMLElement) RemoveITEMPROP(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *HtmlHTMLElement) ITEMREF(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *HtmlHTMLElement) IfITEMREF(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *HtmlHTMLElement) RemoveITEMREF(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *HtmlHTMLElement) ITEMSCOPE() *HtmlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *HtmlHTMLElement) IfITEMSCOPE(cond bool) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *HtmlHTMLElement) RemoveITEMSCOPE() *HtmlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *HtmlHTMLElement) SetITEMSCOPE(b bool) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) ITEMTYPE(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *HtmlHTMLElement) IfITEMTYPE(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *HtmlHTMLElement) RemoveITEMTYPE(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *HtmlHTMLElement) LANG(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *HtmlHTMLElement) IfLANG(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *HtmlHTMLElement) RemoveLANG(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *HtmlHTMLElement) NONCE(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *HtmlHTMLElement) IfNONCE(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *HtmlHTMLElement) RemoveNONCE(v string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) POPOVER(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *HtmlHTMLElement) IfPOPOVER(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *HtmlHTMLElement) RemovePOPOVER(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *HtmlHTMLElement) SLOT(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *HtmlHTMLElement) IfSLOT(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *HtmlHTMLElement) RemoveSLOT(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *HtmlHTMLElement) SPELLCHECK(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *HtmlHTMLElement) IfSPELLCHECK(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *HtmlHTMLElement) RemoveSPELLCHECK(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *HtmlHTMLElement) STYLE(k, v string) *HtmlHTMLElement {
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

func (e *HtmlHTMLElement) IfSTYLE(cond bool, k string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *HtmlHTMLElement) RemoveSTYLE(k string) *HtmlHTMLElement {
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
func (e *HtmlHTMLElement) TABINDEX(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *HtmlHTMLElement) IfTABINDEX(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *HtmlHTMLElement) RemoveTABINDEX(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *HtmlHTMLElement) TITLE(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *HtmlHTMLElement) IfTITLE(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *HtmlHTMLElement) RemoveTITLE(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *HtmlHTMLElement) TRANSLATE(v string) *HtmlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *HtmlHTMLElement) IfTRANSLATE(cond bool, v string) *HtmlHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *HtmlHTMLElement) RemoveTRANSLATE(v string) *HtmlHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
