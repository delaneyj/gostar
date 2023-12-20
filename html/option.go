package html

import (
	"fmt"
)

type OptionHTMLElement struct {
	*Element
}

func OPTION(children ...ElementBuilder) *OptionHTMLElement {
	return &OptionHTMLElement{
		Element: &Element{
			Tag:           elementTagOPTION,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *OptionHTMLElement) Children(children ...ElementBuilder) *OptionHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *OptionHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *OptionHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *OptionHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *OptionHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *OptionHTMLElement) Text(text string) *OptionHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *OptionHTMLElement) TextF(format string, args ...any) *OptionHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *OptionHTMLElement) Escaped(text string) *OptionHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *OptionHTMLElement) EscapedF(format string, args ...any) *OptionHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *OptionHTMLElement) CustomData(key, value string) *OptionHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *OptionHTMLElement) CustomDataRemove(key string) *OptionHTMLElement {
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
func (e *OptionHTMLElement) ACCESSKEY(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *OptionHTMLElement) IfACCESSKEY(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *OptionHTMLElement) RemoveACCESSKEY(v string) *OptionHTMLElement {
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
func (e *OptionHTMLElement) AUTOCAPITALIZE(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *OptionHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *OptionHTMLElement) RemoveAUTOCAPITALIZE(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *OptionHTMLElement) AUTOFOCUS() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *OptionHTMLElement) IfAUTOFOCUS(cond bool) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *OptionHTMLElement) RemoveAUTOFOCUS() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *OptionHTMLElement) SetAUTOFOCUS(b bool) *OptionHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *OptionHTMLElement) CLASS(v string) *OptionHTMLElement {
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

func (e *OptionHTMLElement) IfCLASS(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *OptionHTMLElement) SetCLASS(v string) *OptionHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *OptionHTMLElement) RemoveCLASS(v string) *OptionHTMLElement {
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
func (e *OptionHTMLElement) CONTENTEDITABLE(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *OptionHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *OptionHTMLElement) RemoveCONTENTEDITABLE(v string) *OptionHTMLElement {
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
func (e *OptionHTMLElement) DIR(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *OptionHTMLElement) IfDIR(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *OptionHTMLElement) RemoveDIR(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//   - boolean_attribute
func (e *OptionHTMLElement) DISABLED() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeDISABLEDKey] = struct{}{}
	return e
}

func (e *OptionHTMLElement) IfDISABLED(cond bool) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.DISABLED()
}

func (e *OptionHTMLElement) RemoveDISABLED() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeDISABLEDKey)
	return e
}

func (e *OptionHTMLElement) SetDISABLED(b bool) *OptionHTMLElement {
	if b {
		return e.DISABLED()
	}
	return e.RemoveDISABLED()
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *OptionHTMLElement) DRAGGABLE(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *OptionHTMLElement) IfDRAGGABLE(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *OptionHTMLElement) RemoveDRAGGABLE(v string) *OptionHTMLElement {
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
func (e *OptionHTMLElement) ENTERKEYHINT(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *OptionHTMLElement) IfENTERKEYHINT(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *OptionHTMLElement) RemoveENTERKEYHINT(v string) *OptionHTMLElement {
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
func (e *OptionHTMLElement) HIDDEN(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *OptionHTMLElement) IfHIDDEN(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *OptionHTMLElement) RemoveHIDDEN(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *OptionHTMLElement) ID(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *OptionHTMLElement) IfID(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *OptionHTMLElement) RemoveID(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *OptionHTMLElement) INERT() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *OptionHTMLElement) IfINERT(cond bool) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *OptionHTMLElement) RemoveINERT() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *OptionHTMLElement) SetINERT(b bool) *OptionHTMLElement {
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
func (e *OptionHTMLElement) INPUTMODE(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *OptionHTMLElement) IfINPUTMODE(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *OptionHTMLElement) RemoveINPUTMODE(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *OptionHTMLElement) IS(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *OptionHTMLElement) IfIS(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *OptionHTMLElement) RemoveIS(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *OptionHTMLElement) ITEMID(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *OptionHTMLElement) IfITEMID(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *OptionHTMLElement) RemoveITEMID(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *OptionHTMLElement) ITEMPROP(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *OptionHTMLElement) IfITEMPROP(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *OptionHTMLElement) RemoveITEMPROP(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *OptionHTMLElement) ITEMREF(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *OptionHTMLElement) IfITEMREF(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *OptionHTMLElement) RemoveITEMREF(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *OptionHTMLElement) ITEMSCOPE() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *OptionHTMLElement) IfITEMSCOPE(cond bool) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *OptionHTMLElement) RemoveITEMSCOPE() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *OptionHTMLElement) SetITEMSCOPE(b bool) *OptionHTMLElement {
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
func (e *OptionHTMLElement) ITEMTYPE(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *OptionHTMLElement) IfITEMTYPE(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *OptionHTMLElement) RemoveITEMTYPE(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LABEL sets the "label" attribute.
// User-visible label
// Values values are constrained to:
//   - text
func (e *OptionHTMLElement) LABEL(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLABELKey] = v
	return e
}

func (e *OptionHTMLElement) IfLABEL(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.LABEL(v)
}

func (e *OptionHTMLElement) RemoveLABEL(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeLABELKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *OptionHTMLElement) LANG(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *OptionHTMLElement) IfLANG(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *OptionHTMLElement) RemoveLANG(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *OptionHTMLElement) NONCE(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *OptionHTMLElement) IfNONCE(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *OptionHTMLElement) RemoveNONCE(v string) *OptionHTMLElement {
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
func (e *OptionHTMLElement) POPOVER(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *OptionHTMLElement) IfPOPOVER(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *OptionHTMLElement) RemovePOPOVER(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SELECTED sets the "selected" attribute.
// Whether the option is selected by default
// Values values are constrained to:
//   - boolean_attribute
func (e *OptionHTMLElement) SELECTED() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeSELECTEDKey] = struct{}{}
	return e
}

func (e *OptionHTMLElement) IfSELECTED(cond bool) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.SELECTED()
}

func (e *OptionHTMLElement) RemoveSELECTED() *OptionHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeSELECTEDKey)
	return e
}

func (e *OptionHTMLElement) SetSELECTED(b bool) *OptionHTMLElement {
	if b {
		return e.SELECTED()
	}
	return e.RemoveSELECTED()
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *OptionHTMLElement) SLOT(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *OptionHTMLElement) IfSLOT(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *OptionHTMLElement) RemoveSLOT(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *OptionHTMLElement) SPELLCHECK(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *OptionHTMLElement) IfSPELLCHECK(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *OptionHTMLElement) RemoveSPELLCHECK(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *OptionHTMLElement) STYLE(k, v string) *OptionHTMLElement {
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

func (e *OptionHTMLElement) IfSTYLE(cond bool, k string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *OptionHTMLElement) RemoveSTYLE(k string) *OptionHTMLElement {
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
func (e *OptionHTMLElement) TABINDEX(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *OptionHTMLElement) IfTABINDEX(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *OptionHTMLElement) RemoveTABINDEX(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *OptionHTMLElement) TITLE(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *OptionHTMLElement) IfTITLE(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *OptionHTMLElement) RemoveTITLE(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *OptionHTMLElement) TRANSLATE(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *OptionHTMLElement) IfTRANSLATE(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *OptionHTMLElement) RemoveTRANSLATE(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//   - valid_floating_point_number
func (e *OptionHTMLElement) VALUE(v string) *OptionHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeVALUEKey] = v
	return e
}

func (e *OptionHTMLElement) IfVALUE(cond bool, v string) *OptionHTMLElement {
	if !cond {
		return e
	}
	return e.VALUE(v)
}

func (e *OptionHTMLElement) RemoveVALUE(v string) *OptionHTMLElement {
	delete(e.StringAttributes, attributeVALUEKey)
	return e
}
