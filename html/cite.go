package html

import (
	"fmt"
)

type CiteHTMLElement struct {
	*Element
}

func CITE(children ...ElementBuilder) *CiteHTMLElement {
	return &CiteHTMLElement{
		Element: &Element{
			Tag:           elementTagCITE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *CiteHTMLElement) Children(children ...ElementBuilder) *CiteHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *CiteHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *CiteHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *CiteHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *CiteHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *CiteHTMLElement) Text(text string) *CiteHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *CiteHTMLElement) TextF(format string, args ...any) *CiteHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *CiteHTMLElement) Escaped(text string) *CiteHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *CiteHTMLElement) EscapedF(format string, args ...any) *CiteHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *CiteHTMLElement) CustomData(key, value string) *CiteHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *CiteHTMLElement) CustomDataRemove(key string) *CiteHTMLElement {
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
func (e *CiteHTMLElement) ACCESSKEY(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *CiteHTMLElement) IfACCESSKEY(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *CiteHTMLElement) RemoveACCESSKEY(v string) *CiteHTMLElement {
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
func (e *CiteHTMLElement) AUTOCAPITALIZE(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *CiteHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *CiteHTMLElement) RemoveAUTOCAPITALIZE(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *CiteHTMLElement) AUTOFOCUS() *CiteHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *CiteHTMLElement) IfAUTOFOCUS(cond bool) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *CiteHTMLElement) RemoveAUTOFOCUS() *CiteHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *CiteHTMLElement) SetAUTOFOCUS(b bool) *CiteHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *CiteHTMLElement) CLASS(v string) *CiteHTMLElement {
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

func (e *CiteHTMLElement) IfCLASS(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *CiteHTMLElement) SetCLASS(v string) *CiteHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *CiteHTMLElement) RemoveCLASS(v string) *CiteHTMLElement {
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
func (e *CiteHTMLElement) CONTENTEDITABLE(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *CiteHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *CiteHTMLElement) RemoveCONTENTEDITABLE(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *CiteHTMLElement) DIR(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *CiteHTMLElement) IfDIR(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *CiteHTMLElement) RemoveDIR(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *CiteHTMLElement) DRAGGABLE(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *CiteHTMLElement) IfDRAGGABLE(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *CiteHTMLElement) RemoveDRAGGABLE(v string) *CiteHTMLElement {
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
func (e *CiteHTMLElement) ENTERKEYHINT(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *CiteHTMLElement) IfENTERKEYHINT(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *CiteHTMLElement) RemoveENTERKEYHINT(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *CiteHTMLElement) HIDDEN(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *CiteHTMLElement) IfHIDDEN(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *CiteHTMLElement) RemoveHIDDEN(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *CiteHTMLElement) ID(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *CiteHTMLElement) IfID(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *CiteHTMLElement) RemoveID(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *CiteHTMLElement) INERT() *CiteHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *CiteHTMLElement) IfINERT(cond bool) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *CiteHTMLElement) RemoveINERT() *CiteHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *CiteHTMLElement) SetINERT(b bool) *CiteHTMLElement {
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
func (e *CiteHTMLElement) INPUTMODE(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *CiteHTMLElement) IfINPUTMODE(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *CiteHTMLElement) RemoveINPUTMODE(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *CiteHTMLElement) IS(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *CiteHTMLElement) IfIS(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *CiteHTMLElement) RemoveIS(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *CiteHTMLElement) ITEMID(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *CiteHTMLElement) IfITEMID(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *CiteHTMLElement) RemoveITEMID(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *CiteHTMLElement) ITEMPROP(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *CiteHTMLElement) IfITEMPROP(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *CiteHTMLElement) RemoveITEMPROP(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *CiteHTMLElement) ITEMREF(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *CiteHTMLElement) IfITEMREF(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *CiteHTMLElement) RemoveITEMREF(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *CiteHTMLElement) ITEMSCOPE() *CiteHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *CiteHTMLElement) IfITEMSCOPE(cond bool) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *CiteHTMLElement) RemoveITEMSCOPE() *CiteHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *CiteHTMLElement) SetITEMSCOPE(b bool) *CiteHTMLElement {
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
func (e *CiteHTMLElement) ITEMTYPE(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *CiteHTMLElement) IfITEMTYPE(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *CiteHTMLElement) RemoveITEMTYPE(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *CiteHTMLElement) LANG(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *CiteHTMLElement) IfLANG(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *CiteHTMLElement) RemoveLANG(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *CiteHTMLElement) NONCE(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *CiteHTMLElement) IfNONCE(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *CiteHTMLElement) RemoveNONCE(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *CiteHTMLElement) POPOVER(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *CiteHTMLElement) IfPOPOVER(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *CiteHTMLElement) RemovePOPOVER(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *CiteHTMLElement) SLOT(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *CiteHTMLElement) IfSLOT(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *CiteHTMLElement) RemoveSLOT(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *CiteHTMLElement) SPELLCHECK(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *CiteHTMLElement) IfSPELLCHECK(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *CiteHTMLElement) RemoveSPELLCHECK(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *CiteHTMLElement) STYLE(k, v string) *CiteHTMLElement {
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

func (e *CiteHTMLElement) IfSTYLE(cond bool, k string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *CiteHTMLElement) RemoveSTYLE(k string) *CiteHTMLElement {
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
func (e *CiteHTMLElement) TABINDEX(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *CiteHTMLElement) IfTABINDEX(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *CiteHTMLElement) RemoveTABINDEX(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *CiteHTMLElement) TITLE(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *CiteHTMLElement) IfTITLE(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *CiteHTMLElement) RemoveTITLE(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *CiteHTMLElement) TRANSLATE(v string) *CiteHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *CiteHTMLElement) IfTRANSLATE(cond bool, v string) *CiteHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *CiteHTMLElement) RemoveTRANSLATE(v string) *CiteHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
