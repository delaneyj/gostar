package html

import (
	"fmt"
)

type ObjectHTMLElement struct {
	*Element
}

func OBJECT(children ...ElementBuilder) *ObjectHTMLElement {
	return &ObjectHTMLElement{
		Element: &Element{
			Tag:           elementTagOBJECT,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *ObjectHTMLElement) Children(children ...ElementBuilder) *ObjectHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *ObjectHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ObjectHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *ObjectHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ObjectHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *ObjectHTMLElement) Text(text string) *ObjectHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *ObjectHTMLElement) TextF(format string, args ...any) *ObjectHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *ObjectHTMLElement) Escaped(text string) *ObjectHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *ObjectHTMLElement) EscapedF(format string, args ...any) *ObjectHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *ObjectHTMLElement) CustomData(key, value string) *ObjectHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ObjectHTMLElement) CustomDataRemove(key string) *ObjectHTMLElement {
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
func (e *ObjectHTMLElement) ACCESSKEY(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *ObjectHTMLElement) IfACCESSKEY(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *ObjectHTMLElement) RemoveACCESSKEY(v string) *ObjectHTMLElement {
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
func (e *ObjectHTMLElement) AUTOCAPITALIZE(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *ObjectHTMLElement) RemoveAUTOCAPITALIZE(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *ObjectHTMLElement) AUTOFOCUS() *ObjectHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *ObjectHTMLElement) IfAUTOFOCUS(cond bool) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *ObjectHTMLElement) RemoveAUTOFOCUS() *ObjectHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *ObjectHTMLElement) SetAUTOFOCUS(b bool) *ObjectHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *ObjectHTMLElement) CLASS(v string) *ObjectHTMLElement {
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

func (e *ObjectHTMLElement) IfCLASS(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *ObjectHTMLElement) SetCLASS(v string) *ObjectHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *ObjectHTMLElement) RemoveCLASS(v string) *ObjectHTMLElement {
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
func (e *ObjectHTMLElement) CONTENTEDITABLE(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *ObjectHTMLElement) RemoveCONTENTEDITABLE(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DATA sets the "data" attribute.
// Address of the resource
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ObjectHTMLElement) DATA(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDATAKey] = v
	return e
}

func (e *ObjectHTMLElement) IfDATA(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.DATA(v)
}

func (e *ObjectHTMLElement) RemoveDATA(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeDATAKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (e *ObjectHTMLElement) DIR(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *ObjectHTMLElement) IfDIR(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *ObjectHTMLElement) RemoveDIR(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *ObjectHTMLElement) DRAGGABLE(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfDRAGGABLE(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *ObjectHTMLElement) RemoveDRAGGABLE(v string) *ObjectHTMLElement {
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
func (e *ObjectHTMLElement) ENTERKEYHINT(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *ObjectHTMLElement) IfENTERKEYHINT(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *ObjectHTMLElement) RemoveENTERKEYHINT(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//   - id
func (e *ObjectHTMLElement) FORM(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFORMKey] = v
	return e
}

func (e *ObjectHTMLElement) IfFORM(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.FORM(v)
}

func (e *ObjectHTMLElement) RemoveFORM(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeFORMKey)
	return e
}

// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *ObjectHTMLElement) HEIGHT(v int) *ObjectHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeHEIGHTKey] = v
	return e
}

func (e *ObjectHTMLElement) IfHEIGHT(cond bool, v int) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.HEIGHT(v)
}

func (e *ObjectHTMLElement) RemoveHEIGHT(v int) *ObjectHTMLElement {
	delete(e.IntAttributes, attributeHEIGHTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - until_found
//   - until_found
//   - hidden
//   - hidden
func (e *ObjectHTMLElement) HIDDEN(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *ObjectHTMLElement) IfHIDDEN(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *ObjectHTMLElement) RemoveHIDDEN(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *ObjectHTMLElement) ID(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *ObjectHTMLElement) IfID(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *ObjectHTMLElement) RemoveID(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *ObjectHTMLElement) INERT() *ObjectHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *ObjectHTMLElement) IfINERT(cond bool) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *ObjectHTMLElement) RemoveINERT() *ObjectHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *ObjectHTMLElement) SetINERT(b bool) *ObjectHTMLElement {
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
func (e *ObjectHTMLElement) INPUTMODE(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfINPUTMODE(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *ObjectHTMLElement) RemoveINPUTMODE(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *ObjectHTMLElement) IS(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *ObjectHTMLElement) IfIS(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *ObjectHTMLElement) RemoveIS(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *ObjectHTMLElement) ITEMID(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *ObjectHTMLElement) IfITEMID(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *ObjectHTMLElement) RemoveITEMID(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *ObjectHTMLElement) ITEMPROP(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *ObjectHTMLElement) IfITEMPROP(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *ObjectHTMLElement) RemoveITEMPROP(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ObjectHTMLElement) ITEMREF(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *ObjectHTMLElement) IfITEMREF(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *ObjectHTMLElement) RemoveITEMREF(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *ObjectHTMLElement) ITEMSCOPE() *ObjectHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *ObjectHTMLElement) IfITEMSCOPE(cond bool) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *ObjectHTMLElement) RemoveITEMSCOPE() *ObjectHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *ObjectHTMLElement) SetITEMSCOPE(b bool) *ObjectHTMLElement {
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
func (e *ObjectHTMLElement) ITEMTYPE(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfITEMTYPE(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *ObjectHTMLElement) RemoveITEMTYPE(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ObjectHTMLElement) LANG(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *ObjectHTMLElement) IfLANG(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *ObjectHTMLElement) RemoveLANG(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *ObjectHTMLElement) NAME(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfNAME(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *ObjectHTMLElement) RemoveNAME(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *ObjectHTMLElement) NONCE(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfNONCE(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *ObjectHTMLElement) RemoveNONCE(v string) *ObjectHTMLElement {
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
func (e *ObjectHTMLElement) POPOVER(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *ObjectHTMLElement) IfPOPOVER(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *ObjectHTMLElement) RemovePOPOVER(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *ObjectHTMLElement) SLOT(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *ObjectHTMLElement) IfSLOT(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *ObjectHTMLElement) RemoveSLOT(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *ObjectHTMLElement) SPELLCHECK(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *ObjectHTMLElement) IfSPELLCHECK(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *ObjectHTMLElement) RemoveSPELLCHECK(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ObjectHTMLElement) STYLE(k, v string) *ObjectHTMLElement {
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

func (e *ObjectHTMLElement) IfSTYLE(cond bool, k string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *ObjectHTMLElement) RemoveSTYLE(k string) *ObjectHTMLElement {
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
func (e *ObjectHTMLElement) TABINDEX(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *ObjectHTMLElement) IfTABINDEX(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *ObjectHTMLElement) RemoveTABINDEX(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *ObjectHTMLElement) TITLE(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfTITLE(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *ObjectHTMLElement) RemoveTITLE(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *ObjectHTMLElement) TRANSLATE(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfTRANSLATE(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *ObjectHTMLElement) RemoveTRANSLATE(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//   - module
//   - valid_mime_type_string
//   - java_script_mime_type_essence_match
func (e *ObjectHTMLElement) TYPE(v string) *ObjectHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTYPEKey] = v
	return e
}

func (e *ObjectHTMLElement) IfTYPE(cond bool, v string) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.TYPE(v)
}

func (e *ObjectHTMLElement) RemoveTYPE(v string) *ObjectHTMLElement {
	delete(e.StringAttributes, attributeTYPEKey)
	return e
}

// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *ObjectHTMLElement) WIDTH(v int) *ObjectHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeWIDTHKey] = v
	return e
}

func (e *ObjectHTMLElement) IfWIDTH(cond bool, v int) *ObjectHTMLElement {
	if !cond {
		return e
	}
	return e.WIDTH(v)
}

func (e *ObjectHTMLElement) RemoveWIDTH(v int) *ObjectHTMLElement {
	delete(e.IntAttributes, attributeWIDTHKey)
	return e
}
