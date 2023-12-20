package html

import (
	"fmt"
)

type FigcaptionHTMLElement struct {
	*Element
}

func FIGCAPTION(children ...ElementBuilder) *FigcaptionHTMLElement {
	return &FigcaptionHTMLElement{
		Element: &Element{
			Tag:           elementTagFIGCAPTION,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *FigcaptionHTMLElement) Children(children ...ElementBuilder) *FigcaptionHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *FigcaptionHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *FigcaptionHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *FigcaptionHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *FigcaptionHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *FigcaptionHTMLElement) Text(text string) *FigcaptionHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *FigcaptionHTMLElement) TextF(format string, args ...any) *FigcaptionHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *FigcaptionHTMLElement) Escaped(text string) *FigcaptionHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *FigcaptionHTMLElement) EscapedF(format string, args ...any) *FigcaptionHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *FigcaptionHTMLElement) CustomData(key, value string) *FigcaptionHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *FigcaptionHTMLElement) CustomDataRemove(key string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) ACCESSKEY(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfACCESSKEY(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *FigcaptionHTMLElement) RemoveACCESSKEY(v string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) AUTOCAPITALIZE(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *FigcaptionHTMLElement) RemoveAUTOCAPITALIZE(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *FigcaptionHTMLElement) AUTOFOCUS() *FigcaptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *FigcaptionHTMLElement) IfAUTOFOCUS(cond bool) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *FigcaptionHTMLElement) RemoveAUTOFOCUS() *FigcaptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *FigcaptionHTMLElement) SetAUTOFOCUS(b bool) *FigcaptionHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *FigcaptionHTMLElement) CLASS(v string) *FigcaptionHTMLElement {
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

func (e *FigcaptionHTMLElement) IfCLASS(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *FigcaptionHTMLElement) SetCLASS(v string) *FigcaptionHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *FigcaptionHTMLElement) RemoveCLASS(v string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) CONTENTEDITABLE(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *FigcaptionHTMLElement) RemoveCONTENTEDITABLE(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *FigcaptionHTMLElement) DIR(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfDIR(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *FigcaptionHTMLElement) RemoveDIR(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *FigcaptionHTMLElement) DRAGGABLE(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfDRAGGABLE(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *FigcaptionHTMLElement) RemoveDRAGGABLE(v string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) ENTERKEYHINT(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfENTERKEYHINT(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *FigcaptionHTMLElement) RemoveENTERKEYHINT(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *FigcaptionHTMLElement) HIDDEN(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfHIDDEN(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *FigcaptionHTMLElement) RemoveHIDDEN(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *FigcaptionHTMLElement) ID(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfID(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *FigcaptionHTMLElement) RemoveID(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *FigcaptionHTMLElement) INERT() *FigcaptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *FigcaptionHTMLElement) IfINERT(cond bool) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *FigcaptionHTMLElement) RemoveINERT() *FigcaptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *FigcaptionHTMLElement) SetINERT(b bool) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) INPUTMODE(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfINPUTMODE(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *FigcaptionHTMLElement) RemoveINPUTMODE(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *FigcaptionHTMLElement) IS(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfIS(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *FigcaptionHTMLElement) RemoveIS(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *FigcaptionHTMLElement) ITEMID(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfITEMID(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *FigcaptionHTMLElement) RemoveITEMID(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *FigcaptionHTMLElement) ITEMPROP(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfITEMPROP(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *FigcaptionHTMLElement) RemoveITEMPROP(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *FigcaptionHTMLElement) ITEMREF(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfITEMREF(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *FigcaptionHTMLElement) RemoveITEMREF(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *FigcaptionHTMLElement) ITEMSCOPE() *FigcaptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *FigcaptionHTMLElement) IfITEMSCOPE(cond bool) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *FigcaptionHTMLElement) RemoveITEMSCOPE() *FigcaptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *FigcaptionHTMLElement) SetITEMSCOPE(b bool) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) ITEMTYPE(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfITEMTYPE(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *FigcaptionHTMLElement) RemoveITEMTYPE(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *FigcaptionHTMLElement) LANG(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfLANG(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *FigcaptionHTMLElement) RemoveLANG(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *FigcaptionHTMLElement) NONCE(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfNONCE(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *FigcaptionHTMLElement) RemoveNONCE(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *FigcaptionHTMLElement) POPOVER(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfPOPOVER(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *FigcaptionHTMLElement) RemovePOPOVER(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *FigcaptionHTMLElement) SLOT(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfSLOT(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *FigcaptionHTMLElement) RemoveSLOT(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *FigcaptionHTMLElement) SPELLCHECK(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfSPELLCHECK(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *FigcaptionHTMLElement) RemoveSPELLCHECK(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *FigcaptionHTMLElement) STYLE(k, v string) *FigcaptionHTMLElement {
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

func (e *FigcaptionHTMLElement) IfSTYLE(cond bool, k string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *FigcaptionHTMLElement) RemoveSTYLE(k string) *FigcaptionHTMLElement {
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
func (e *FigcaptionHTMLElement) TABINDEX(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfTABINDEX(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *FigcaptionHTMLElement) RemoveTABINDEX(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *FigcaptionHTMLElement) TITLE(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfTITLE(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *FigcaptionHTMLElement) RemoveTITLE(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *FigcaptionHTMLElement) TRANSLATE(v string) *FigcaptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *FigcaptionHTMLElement) IfTRANSLATE(cond bool, v string) *FigcaptionHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *FigcaptionHTMLElement) RemoveTRANSLATE(v string) *FigcaptionHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
