package html

import (
	"fmt"
)

type PictureHTMLElement struct {
	*Element
}

func PICTURE(children ...ElementBuilder) *PictureHTMLElement {
	return &PictureHTMLElement{
		Element: &Element{
			Tag:           elementTagPICTURE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *PictureHTMLElement) Children(children ...ElementBuilder) *PictureHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *PictureHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *PictureHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *PictureHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *PictureHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *PictureHTMLElement) Text(text string) *PictureHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *PictureHTMLElement) TextF(format string, args ...any) *PictureHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *PictureHTMLElement) Escaped(text string) *PictureHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *PictureHTMLElement) EscapedF(format string, args ...any) *PictureHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *PictureHTMLElement) CustomData(key, value string) *PictureHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *PictureHTMLElement) CustomDataRemove(key string) *PictureHTMLElement {
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
func (e *PictureHTMLElement) ACCESSKEY(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *PictureHTMLElement) IfACCESSKEY(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *PictureHTMLElement) RemoveACCESSKEY(v string) *PictureHTMLElement {
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
func (e *PictureHTMLElement) AUTOCAPITALIZE(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *PictureHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *PictureHTMLElement) RemoveAUTOCAPITALIZE(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *PictureHTMLElement) AUTOFOCUS() *PictureHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *PictureHTMLElement) IfAUTOFOCUS(cond bool) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *PictureHTMLElement) RemoveAUTOFOCUS() *PictureHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *PictureHTMLElement) SetAUTOFOCUS(b bool) *PictureHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *PictureHTMLElement) CLASS(v string) *PictureHTMLElement {
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

func (e *PictureHTMLElement) IfCLASS(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *PictureHTMLElement) SetCLASS(v string) *PictureHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *PictureHTMLElement) RemoveCLASS(v string) *PictureHTMLElement {
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
func (e *PictureHTMLElement) CONTENTEDITABLE(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *PictureHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *PictureHTMLElement) RemoveCONTENTEDITABLE(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *PictureHTMLElement) DIR(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *PictureHTMLElement) IfDIR(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *PictureHTMLElement) RemoveDIR(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *PictureHTMLElement) DRAGGABLE(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *PictureHTMLElement) IfDRAGGABLE(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *PictureHTMLElement) RemoveDRAGGABLE(v string) *PictureHTMLElement {
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
func (e *PictureHTMLElement) ENTERKEYHINT(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *PictureHTMLElement) IfENTERKEYHINT(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *PictureHTMLElement) RemoveENTERKEYHINT(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *PictureHTMLElement) HEIGHT(v int) *PictureHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeHEIGHTKey] = v
	return e
}

func (e *PictureHTMLElement) IfHEIGHT(cond bool, v int) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.HEIGHT(v)
}

func (e *PictureHTMLElement) RemoveHEIGHT(v int) *PictureHTMLElement {
	delete(e.IntAttributes, attributeHEIGHTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *PictureHTMLElement) HIDDEN(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *PictureHTMLElement) IfHIDDEN(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *PictureHTMLElement) RemoveHIDDEN(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *PictureHTMLElement) ID(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *PictureHTMLElement) IfID(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *PictureHTMLElement) RemoveID(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *PictureHTMLElement) INERT() *PictureHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *PictureHTMLElement) IfINERT(cond bool) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *PictureHTMLElement) RemoveINERT() *PictureHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *PictureHTMLElement) SetINERT(b bool) *PictureHTMLElement {
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
func (e *PictureHTMLElement) INPUTMODE(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *PictureHTMLElement) IfINPUTMODE(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *PictureHTMLElement) RemoveINPUTMODE(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *PictureHTMLElement) IS(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *PictureHTMLElement) IfIS(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *PictureHTMLElement) RemoveIS(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *PictureHTMLElement) ITEMID(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *PictureHTMLElement) IfITEMID(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *PictureHTMLElement) RemoveITEMID(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *PictureHTMLElement) ITEMPROP(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *PictureHTMLElement) IfITEMPROP(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *PictureHTMLElement) RemoveITEMPROP(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *PictureHTMLElement) ITEMREF(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *PictureHTMLElement) IfITEMREF(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *PictureHTMLElement) RemoveITEMREF(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *PictureHTMLElement) ITEMSCOPE() *PictureHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *PictureHTMLElement) IfITEMSCOPE(cond bool) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *PictureHTMLElement) RemoveITEMSCOPE() *PictureHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *PictureHTMLElement) SetITEMSCOPE(b bool) *PictureHTMLElement {
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
func (e *PictureHTMLElement) ITEMTYPE(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *PictureHTMLElement) IfITEMTYPE(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *PictureHTMLElement) RemoveITEMTYPE(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *PictureHTMLElement) LANG(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *PictureHTMLElement) IfLANG(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *PictureHTMLElement) RemoveLANG(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *PictureHTMLElement) NONCE(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *PictureHTMLElement) IfNONCE(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *PictureHTMLElement) RemoveNONCE(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *PictureHTMLElement) POPOVER(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *PictureHTMLElement) IfPOPOVER(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *PictureHTMLElement) RemovePOPOVER(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *PictureHTMLElement) SLOT(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *PictureHTMLElement) IfSLOT(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *PictureHTMLElement) RemoveSLOT(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *PictureHTMLElement) SPELLCHECK(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *PictureHTMLElement) IfSPELLCHECK(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *PictureHTMLElement) RemoveSPELLCHECK(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *PictureHTMLElement) STYLE(k, v string) *PictureHTMLElement {
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

func (e *PictureHTMLElement) IfSTYLE(cond bool, k string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *PictureHTMLElement) RemoveSTYLE(k string) *PictureHTMLElement {
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
func (e *PictureHTMLElement) TABINDEX(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *PictureHTMLElement) IfTABINDEX(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *PictureHTMLElement) RemoveTABINDEX(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *PictureHTMLElement) TITLE(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *PictureHTMLElement) IfTITLE(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *PictureHTMLElement) RemoveTITLE(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *PictureHTMLElement) TRANSLATE(v string) *PictureHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *PictureHTMLElement) IfTRANSLATE(cond bool, v string) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *PictureHTMLElement) RemoveTRANSLATE(v string) *PictureHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *PictureHTMLElement) WIDTH(v int) *PictureHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeWIDTHKey] = v
	return e
}

func (e *PictureHTMLElement) IfWIDTH(cond bool, v int) *PictureHTMLElement {
	if !cond {
		return e
	}
	return e.WIDTH(v)
}

func (e *PictureHTMLElement) RemoveWIDTH(v int) *PictureHTMLElement {
	delete(e.IntAttributes, attributeWIDTHKey)
	return e
}
