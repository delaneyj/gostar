package html

import (
	"fmt"
)

type SvgHTMLElement struct {
	*Element
}

func SVG(children ...ElementBuilder) *SvgHTMLElement {
	return &SvgHTMLElement{
		Element: &Element{
			Tag:           elementTagSVG,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *SvgHTMLElement) Children(children ...ElementBuilder) *SvgHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SvgHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *SvgHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SvgHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *SvgHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SvgHTMLElement) Text(text string) *SvgHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SvgHTMLElement) TextF(format string, args ...any) *SvgHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SvgHTMLElement) Escaped(text string) *SvgHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SvgHTMLElement) EscapedF(format string, args ...any) *SvgHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SvgHTMLElement) CustomData(key, value string) *SvgHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *SvgHTMLElement) CustomDataRemove(key string) *SvgHTMLElement {
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
func (e *SvgHTMLElement) ACCESSKEY(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *SvgHTMLElement) IfACCESSKEY(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *SvgHTMLElement) RemoveACCESSKEY(v string) *SvgHTMLElement {
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
func (e *SvgHTMLElement) AUTOCAPITALIZE(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *SvgHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *SvgHTMLElement) RemoveAUTOCAPITALIZE(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *SvgHTMLElement) AUTOFOCUS() *SvgHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *SvgHTMLElement) IfAUTOFOCUS(cond bool) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *SvgHTMLElement) RemoveAUTOFOCUS() *SvgHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *SvgHTMLElement) SetAUTOFOCUS(b bool) *SvgHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *SvgHTMLElement) CLASS(v string) *SvgHTMLElement {
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

func (e *SvgHTMLElement) IfCLASS(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *SvgHTMLElement) SetCLASS(v string) *SvgHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *SvgHTMLElement) RemoveCLASS(v string) *SvgHTMLElement {
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
func (e *SvgHTMLElement) CONTENTEDITABLE(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *SvgHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *SvgHTMLElement) RemoveCONTENTEDITABLE(v string) *SvgHTMLElement {
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
func (e *SvgHTMLElement) DIR(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *SvgHTMLElement) IfDIR(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *SvgHTMLElement) RemoveDIR(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *SvgHTMLElement) DRAGGABLE(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *SvgHTMLElement) IfDRAGGABLE(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *SvgHTMLElement) RemoveDRAGGABLE(v string) *SvgHTMLElement {
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
func (e *SvgHTMLElement) ENTERKEYHINT(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *SvgHTMLElement) IfENTERKEYHINT(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *SvgHTMLElement) RemoveENTERKEYHINT(v string) *SvgHTMLElement {
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
func (e *SvgHTMLElement) HIDDEN(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *SvgHTMLElement) IfHIDDEN(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *SvgHTMLElement) RemoveHIDDEN(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *SvgHTMLElement) ID(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *SvgHTMLElement) IfID(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *SvgHTMLElement) RemoveID(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *SvgHTMLElement) INERT() *SvgHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *SvgHTMLElement) IfINERT(cond bool) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *SvgHTMLElement) RemoveINERT() *SvgHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *SvgHTMLElement) SetINERT(b bool) *SvgHTMLElement {
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
func (e *SvgHTMLElement) INPUTMODE(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *SvgHTMLElement) IfINPUTMODE(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *SvgHTMLElement) RemoveINPUTMODE(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *SvgHTMLElement) IS(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *SvgHTMLElement) IfIS(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *SvgHTMLElement) RemoveIS(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *SvgHTMLElement) ITEMID(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *SvgHTMLElement) IfITEMID(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *SvgHTMLElement) RemoveITEMID(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *SvgHTMLElement) ITEMPROP(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *SvgHTMLElement) IfITEMPROP(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *SvgHTMLElement) RemoveITEMPROP(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *SvgHTMLElement) ITEMREF(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *SvgHTMLElement) IfITEMREF(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *SvgHTMLElement) RemoveITEMREF(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *SvgHTMLElement) ITEMSCOPE() *SvgHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *SvgHTMLElement) IfITEMSCOPE(cond bool) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *SvgHTMLElement) RemoveITEMSCOPE() *SvgHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *SvgHTMLElement) SetITEMSCOPE(b bool) *SvgHTMLElement {
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
func (e *SvgHTMLElement) ITEMTYPE(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *SvgHTMLElement) IfITEMTYPE(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *SvgHTMLElement) RemoveITEMTYPE(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *SvgHTMLElement) LANG(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *SvgHTMLElement) IfLANG(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *SvgHTMLElement) RemoveLANG(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *SvgHTMLElement) NONCE(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *SvgHTMLElement) IfNONCE(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *SvgHTMLElement) RemoveNONCE(v string) *SvgHTMLElement {
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
func (e *SvgHTMLElement) POPOVER(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *SvgHTMLElement) IfPOPOVER(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *SvgHTMLElement) RemovePOPOVER(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *SvgHTMLElement) SLOT(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *SvgHTMLElement) IfSLOT(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *SvgHTMLElement) RemoveSLOT(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *SvgHTMLElement) SPELLCHECK(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *SvgHTMLElement) IfSPELLCHECK(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *SvgHTMLElement) RemoveSPELLCHECK(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *SvgHTMLElement) STYLE(k, v string) *SvgHTMLElement {
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

func (e *SvgHTMLElement) IfSTYLE(cond bool, k string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *SvgHTMLElement) RemoveSTYLE(k string) *SvgHTMLElement {
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
func (e *SvgHTMLElement) TABINDEX(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *SvgHTMLElement) IfTABINDEX(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *SvgHTMLElement) RemoveTABINDEX(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *SvgHTMLElement) TITLE(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *SvgHTMLElement) IfTITLE(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *SvgHTMLElement) RemoveTITLE(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *SvgHTMLElement) TRANSLATE(v string) *SvgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *SvgHTMLElement) IfTRANSLATE(cond bool, v string) *SvgHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *SvgHTMLElement) RemoveTRANSLATE(v string) *SvgHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
