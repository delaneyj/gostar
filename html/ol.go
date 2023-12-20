package html

import (
	"fmt"
)

type OlHTMLElement struct {
	*Element
}

func OL(children ...ElementBuilder) *OlHTMLElement {
	return &OlHTMLElement{
		Element: &Element{
			Tag:           elementTagOL,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *OlHTMLElement) Children(children ...ElementBuilder) *OlHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *OlHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *OlHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *OlHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *OlHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *OlHTMLElement) Text(text string) *OlHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *OlHTMLElement) TextF(format string, args ...any) *OlHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *OlHTMLElement) Escaped(text string) *OlHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *OlHTMLElement) EscapedF(format string, args ...any) *OlHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *OlHTMLElement) CustomData(key, value string) *OlHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *OlHTMLElement) CustomDataRemove(key string) *OlHTMLElement {
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
func (e *OlHTMLElement) ACCESSKEY(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *OlHTMLElement) IfACCESSKEY(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *OlHTMLElement) RemoveACCESSKEY(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) AUTOCAPITALIZE(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *OlHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *OlHTMLElement) RemoveAUTOCAPITALIZE(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *OlHTMLElement) AUTOFOCUS() *OlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *OlHTMLElement) IfAUTOFOCUS(cond bool) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *OlHTMLElement) RemoveAUTOFOCUS() *OlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *OlHTMLElement) SetAUTOFOCUS(b bool) *OlHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *OlHTMLElement) CLASS(v string) *OlHTMLElement {
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

func (e *OlHTMLElement) IfCLASS(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *OlHTMLElement) SetCLASS(v string) *OlHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *OlHTMLElement) RemoveCLASS(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) CONTENTEDITABLE(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *OlHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *OlHTMLElement) RemoveCONTENTEDITABLE(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) DIR(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *OlHTMLElement) IfDIR(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *OlHTMLElement) RemoveDIR(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *OlHTMLElement) DRAGGABLE(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *OlHTMLElement) IfDRAGGABLE(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *OlHTMLElement) RemoveDRAGGABLE(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) ENTERKEYHINT(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *OlHTMLElement) IfENTERKEYHINT(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *OlHTMLElement) RemoveENTERKEYHINT(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) HIDDEN(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *OlHTMLElement) IfHIDDEN(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *OlHTMLElement) RemoveHIDDEN(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *OlHTMLElement) ID(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *OlHTMLElement) IfID(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *OlHTMLElement) RemoveID(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *OlHTMLElement) INERT() *OlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *OlHTMLElement) IfINERT(cond bool) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *OlHTMLElement) RemoveINERT() *OlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *OlHTMLElement) SetINERT(b bool) *OlHTMLElement {
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
func (e *OlHTMLElement) INPUTMODE(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *OlHTMLElement) IfINPUTMODE(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *OlHTMLElement) RemoveINPUTMODE(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *OlHTMLElement) IS(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *OlHTMLElement) IfIS(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *OlHTMLElement) RemoveIS(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *OlHTMLElement) ITEMID(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *OlHTMLElement) IfITEMID(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *OlHTMLElement) RemoveITEMID(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *OlHTMLElement) ITEMPROP(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *OlHTMLElement) IfITEMPROP(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *OlHTMLElement) RemoveITEMPROP(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *OlHTMLElement) ITEMREF(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *OlHTMLElement) IfITEMREF(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *OlHTMLElement) RemoveITEMREF(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *OlHTMLElement) ITEMSCOPE() *OlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *OlHTMLElement) IfITEMSCOPE(cond bool) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *OlHTMLElement) RemoveITEMSCOPE() *OlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *OlHTMLElement) SetITEMSCOPE(b bool) *OlHTMLElement {
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
func (e *OlHTMLElement) ITEMTYPE(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *OlHTMLElement) IfITEMTYPE(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *OlHTMLElement) RemoveITEMTYPE(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *OlHTMLElement) LANG(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *OlHTMLElement) IfLANG(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *OlHTMLElement) RemoveLANG(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *OlHTMLElement) NONCE(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *OlHTMLElement) IfNONCE(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *OlHTMLElement) RemoveNONCE(v string) *OlHTMLElement {
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
func (e *OlHTMLElement) POPOVER(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *OlHTMLElement) IfPOPOVER(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *OlHTMLElement) RemovePOPOVER(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// REVERSED sets the "reversed" attribute.
// Number the list backwards
// Values values are constrained to:
//   - boolean_attribute
func (e *OlHTMLElement) REVERSED() *OlHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeREVERSEDKey] = struct{}{}
	return e
}

func (e *OlHTMLElement) IfREVERSED(cond bool) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.REVERSED()
}

func (e *OlHTMLElement) RemoveREVERSED() *OlHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeREVERSEDKey)
	return e
}

func (e *OlHTMLElement) SetREVERSED(b bool) *OlHTMLElement {
	if b {
		return e.REVERSED()
	}
	return e.RemoveREVERSED()
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *OlHTMLElement) SLOT(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *OlHTMLElement) IfSLOT(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *OlHTMLElement) RemoveSLOT(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *OlHTMLElement) SPELLCHECK(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *OlHTMLElement) IfSPELLCHECK(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *OlHTMLElement) RemoveSPELLCHECK(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// START sets the "start" attribute.
// Starting value of the list
// Values values are constrained to:
//   - valid_integer
func (e *OlHTMLElement) START(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSTARTKey] = v
	return e
}

func (e *OlHTMLElement) IfSTART(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.START(v)
}

func (e *OlHTMLElement) RemoveSTART(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeSTARTKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *OlHTMLElement) STYLE(k, v string) *OlHTMLElement {
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

func (e *OlHTMLElement) IfSTYLE(cond bool, k string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *OlHTMLElement) RemoveSTYLE(k string) *OlHTMLElement {
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
func (e *OlHTMLElement) TABINDEX(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *OlHTMLElement) IfTABINDEX(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *OlHTMLElement) RemoveTABINDEX(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *OlHTMLElement) TITLE(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *OlHTMLElement) IfTITLE(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *OlHTMLElement) RemoveTITLE(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *OlHTMLElement) TRANSLATE(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *OlHTMLElement) IfTRANSLATE(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *OlHTMLElement) RemoveTRANSLATE(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//   - module
//   - valid_mime_type_string
//   - java_script_mime_type_essence_match
func (e *OlHTMLElement) TYPE(v string) *OlHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTYPEKey] = v
	return e
}

func (e *OlHTMLElement) IfTYPE(cond bool, v string) *OlHTMLElement {
	if !cond {
		return e
	}
	return e.TYPE(v)
}

func (e *OlHTMLElement) RemoveTYPE(v string) *OlHTMLElement {
	delete(e.StringAttributes, attributeTYPEKey)
	return e
}
