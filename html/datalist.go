package html

import (
	"fmt"
)

type DatalistHTMLElement struct {
	*Element
}

func DATALIST(children ...ElementBuilder) *DatalistHTMLElement {
	return &DatalistHTMLElement{
		Element: &Element{
			Tag:           elementTagDATALIST,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *DatalistHTMLElement) Children(children ...ElementBuilder) *DatalistHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *DatalistHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *DatalistHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *DatalistHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *DatalistHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *DatalistHTMLElement) Text(text string) *DatalistHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *DatalistHTMLElement) TextF(format string, args ...any) *DatalistHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *DatalistHTMLElement) Escaped(text string) *DatalistHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *DatalistHTMLElement) EscapedF(format string, args ...any) *DatalistHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *DatalistHTMLElement) CustomData(key, value string) *DatalistHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *DatalistHTMLElement) CustomDataRemove(key string) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) ACCESSKEY(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *DatalistHTMLElement) IfACCESSKEY(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *DatalistHTMLElement) RemoveACCESSKEY(v string) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) AUTOCAPITALIZE(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *DatalistHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *DatalistHTMLElement) RemoveAUTOCAPITALIZE(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *DatalistHTMLElement) AUTOFOCUS() *DatalistHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *DatalistHTMLElement) IfAUTOFOCUS(cond bool) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *DatalistHTMLElement) RemoveAUTOFOCUS() *DatalistHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *DatalistHTMLElement) SetAUTOFOCUS(b bool) *DatalistHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *DatalistHTMLElement) CLASS(v string) *DatalistHTMLElement {
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

func (e *DatalistHTMLElement) IfCLASS(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *DatalistHTMLElement) SetCLASS(v string) *DatalistHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *DatalistHTMLElement) RemoveCLASS(v string) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) CONTENTEDITABLE(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *DatalistHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *DatalistHTMLElement) RemoveCONTENTEDITABLE(v string) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) DIR(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *DatalistHTMLElement) IfDIR(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *DatalistHTMLElement) RemoveDIR(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *DatalistHTMLElement) DRAGGABLE(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *DatalistHTMLElement) IfDRAGGABLE(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *DatalistHTMLElement) RemoveDRAGGABLE(v string) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) ENTERKEYHINT(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *DatalistHTMLElement) IfENTERKEYHINT(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *DatalistHTMLElement) RemoveENTERKEYHINT(v string) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) HIDDEN(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *DatalistHTMLElement) IfHIDDEN(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *DatalistHTMLElement) RemoveHIDDEN(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *DatalistHTMLElement) ID(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *DatalistHTMLElement) IfID(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *DatalistHTMLElement) RemoveID(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *DatalistHTMLElement) INERT() *DatalistHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *DatalistHTMLElement) IfINERT(cond bool) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *DatalistHTMLElement) RemoveINERT() *DatalistHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *DatalistHTMLElement) SetINERT(b bool) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) INPUTMODE(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *DatalistHTMLElement) IfINPUTMODE(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *DatalistHTMLElement) RemoveINPUTMODE(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *DatalistHTMLElement) IS(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *DatalistHTMLElement) IfIS(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *DatalistHTMLElement) RemoveIS(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *DatalistHTMLElement) ITEMID(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *DatalistHTMLElement) IfITEMID(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *DatalistHTMLElement) RemoveITEMID(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *DatalistHTMLElement) ITEMPROP(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *DatalistHTMLElement) IfITEMPROP(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *DatalistHTMLElement) RemoveITEMPROP(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *DatalistHTMLElement) ITEMREF(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *DatalistHTMLElement) IfITEMREF(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *DatalistHTMLElement) RemoveITEMREF(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *DatalistHTMLElement) ITEMSCOPE() *DatalistHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *DatalistHTMLElement) IfITEMSCOPE(cond bool) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *DatalistHTMLElement) RemoveITEMSCOPE() *DatalistHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *DatalistHTMLElement) SetITEMSCOPE(b bool) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) ITEMTYPE(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *DatalistHTMLElement) IfITEMTYPE(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *DatalistHTMLElement) RemoveITEMTYPE(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *DatalistHTMLElement) LANG(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *DatalistHTMLElement) IfLANG(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *DatalistHTMLElement) RemoveLANG(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *DatalistHTMLElement) NONCE(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *DatalistHTMLElement) IfNONCE(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *DatalistHTMLElement) RemoveNONCE(v string) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) POPOVER(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *DatalistHTMLElement) IfPOPOVER(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *DatalistHTMLElement) RemovePOPOVER(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *DatalistHTMLElement) SLOT(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *DatalistHTMLElement) IfSLOT(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *DatalistHTMLElement) RemoveSLOT(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *DatalistHTMLElement) SPELLCHECK(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *DatalistHTMLElement) IfSPELLCHECK(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *DatalistHTMLElement) RemoveSPELLCHECK(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *DatalistHTMLElement) STYLE(k, v string) *DatalistHTMLElement {
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

func (e *DatalistHTMLElement) IfSTYLE(cond bool, k string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *DatalistHTMLElement) RemoveSTYLE(k string) *DatalistHTMLElement {
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
func (e *DatalistHTMLElement) TABINDEX(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *DatalistHTMLElement) IfTABINDEX(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *DatalistHTMLElement) RemoveTABINDEX(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *DatalistHTMLElement) TITLE(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *DatalistHTMLElement) IfTITLE(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *DatalistHTMLElement) RemoveTITLE(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *DatalistHTMLElement) TRANSLATE(v string) *DatalistHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *DatalistHTMLElement) IfTRANSLATE(cond bool, v string) *DatalistHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *DatalistHTMLElement) RemoveTRANSLATE(v string) *DatalistHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
