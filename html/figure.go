package html

import (
	"fmt"
)

type FigureHTMLElement struct {
	*Element
}

func FIGURE(children ...ElementBuilder) *FigureHTMLElement {
	return &FigureHTMLElement{
		Element: &Element{
			Tag:           elementTagFIGURE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *FigureHTMLElement) Children(children ...ElementBuilder) *FigureHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *FigureHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *FigureHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *FigureHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *FigureHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *FigureHTMLElement) Text(text string) *FigureHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *FigureHTMLElement) TextF(format string, args ...any) *FigureHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *FigureHTMLElement) Escaped(text string) *FigureHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *FigureHTMLElement) EscapedF(format string, args ...any) *FigureHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *FigureHTMLElement) CustomData(key, value string) *FigureHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *FigureHTMLElement) CustomDataRemove(key string) *FigureHTMLElement {
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
func (e *FigureHTMLElement) ACCESSKEY(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *FigureHTMLElement) IfACCESSKEY(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *FigureHTMLElement) RemoveACCESSKEY(v string) *FigureHTMLElement {
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
func (e *FigureHTMLElement) AUTOCAPITALIZE(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *FigureHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *FigureHTMLElement) RemoveAUTOCAPITALIZE(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *FigureHTMLElement) AUTOFOCUS() *FigureHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *FigureHTMLElement) IfAUTOFOCUS(cond bool) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *FigureHTMLElement) RemoveAUTOFOCUS() *FigureHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *FigureHTMLElement) SetAUTOFOCUS(b bool) *FigureHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *FigureHTMLElement) CLASS(v string) *FigureHTMLElement {
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

func (e *FigureHTMLElement) IfCLASS(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *FigureHTMLElement) SetCLASS(v string) *FigureHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *FigureHTMLElement) RemoveCLASS(v string) *FigureHTMLElement {
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
func (e *FigureHTMLElement) CONTENTEDITABLE(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *FigureHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *FigureHTMLElement) RemoveCONTENTEDITABLE(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *FigureHTMLElement) DIR(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *FigureHTMLElement) IfDIR(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *FigureHTMLElement) RemoveDIR(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *FigureHTMLElement) DRAGGABLE(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *FigureHTMLElement) IfDRAGGABLE(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *FigureHTMLElement) RemoveDRAGGABLE(v string) *FigureHTMLElement {
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
func (e *FigureHTMLElement) ENTERKEYHINT(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *FigureHTMLElement) IfENTERKEYHINT(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *FigureHTMLElement) RemoveENTERKEYHINT(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *FigureHTMLElement) HIDDEN(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *FigureHTMLElement) IfHIDDEN(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *FigureHTMLElement) RemoveHIDDEN(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *FigureHTMLElement) ID(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *FigureHTMLElement) IfID(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *FigureHTMLElement) RemoveID(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *FigureHTMLElement) INERT() *FigureHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *FigureHTMLElement) IfINERT(cond bool) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *FigureHTMLElement) RemoveINERT() *FigureHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *FigureHTMLElement) SetINERT(b bool) *FigureHTMLElement {
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
func (e *FigureHTMLElement) INPUTMODE(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *FigureHTMLElement) IfINPUTMODE(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *FigureHTMLElement) RemoveINPUTMODE(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *FigureHTMLElement) IS(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *FigureHTMLElement) IfIS(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *FigureHTMLElement) RemoveIS(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *FigureHTMLElement) ITEMID(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *FigureHTMLElement) IfITEMID(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *FigureHTMLElement) RemoveITEMID(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *FigureHTMLElement) ITEMPROP(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *FigureHTMLElement) IfITEMPROP(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *FigureHTMLElement) RemoveITEMPROP(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *FigureHTMLElement) ITEMREF(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *FigureHTMLElement) IfITEMREF(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *FigureHTMLElement) RemoveITEMREF(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *FigureHTMLElement) ITEMSCOPE() *FigureHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *FigureHTMLElement) IfITEMSCOPE(cond bool) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *FigureHTMLElement) RemoveITEMSCOPE() *FigureHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *FigureHTMLElement) SetITEMSCOPE(b bool) *FigureHTMLElement {
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
func (e *FigureHTMLElement) ITEMTYPE(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *FigureHTMLElement) IfITEMTYPE(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *FigureHTMLElement) RemoveITEMTYPE(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *FigureHTMLElement) LANG(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *FigureHTMLElement) IfLANG(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *FigureHTMLElement) RemoveLANG(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *FigureHTMLElement) NONCE(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *FigureHTMLElement) IfNONCE(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *FigureHTMLElement) RemoveNONCE(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *FigureHTMLElement) POPOVER(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *FigureHTMLElement) IfPOPOVER(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *FigureHTMLElement) RemovePOPOVER(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *FigureHTMLElement) SLOT(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *FigureHTMLElement) IfSLOT(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *FigureHTMLElement) RemoveSLOT(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *FigureHTMLElement) SPELLCHECK(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *FigureHTMLElement) IfSPELLCHECK(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *FigureHTMLElement) RemoveSPELLCHECK(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *FigureHTMLElement) STYLE(k, v string) *FigureHTMLElement {
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

func (e *FigureHTMLElement) IfSTYLE(cond bool, k string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *FigureHTMLElement) RemoveSTYLE(k string) *FigureHTMLElement {
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
func (e *FigureHTMLElement) TABINDEX(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *FigureHTMLElement) IfTABINDEX(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *FigureHTMLElement) RemoveTABINDEX(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *FigureHTMLElement) TITLE(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *FigureHTMLElement) IfTITLE(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *FigureHTMLElement) RemoveTITLE(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *FigureHTMLElement) TRANSLATE(v string) *FigureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *FigureHTMLElement) IfTRANSLATE(cond bool, v string) *FigureHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *FigureHTMLElement) RemoveTRANSLATE(v string) *FigureHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
