package html

import (
	"fmt"
)

type TemplateHTMLElement struct {
	*Element
}

func TEMPLATE(children ...ElementBuilder) *TemplateHTMLElement {
	return &TemplateHTMLElement{
		Element: &Element{
			Tag:           elementTagTEMPLATE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *TemplateHTMLElement) Children(children ...ElementBuilder) *TemplateHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *TemplateHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TemplateHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *TemplateHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TemplateHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *TemplateHTMLElement) Text(text string) *TemplateHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *TemplateHTMLElement) TextF(format string, args ...any) *TemplateHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *TemplateHTMLElement) Escaped(text string) *TemplateHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *TemplateHTMLElement) EscapedF(format string, args ...any) *TemplateHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TemplateHTMLElement) CustomData(key, value string) *TemplateHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TemplateHTMLElement) CustomDataRemove(key string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) ACCESSKEY(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *TemplateHTMLElement) IfACCESSKEY(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *TemplateHTMLElement) RemoveACCESSKEY(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) AUTOCAPITALIZE(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *TemplateHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *TemplateHTMLElement) RemoveAUTOCAPITALIZE(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *TemplateHTMLElement) AUTOFOCUS() *TemplateHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *TemplateHTMLElement) IfAUTOFOCUS(cond bool) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *TemplateHTMLElement) RemoveAUTOFOCUS() *TemplateHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *TemplateHTMLElement) SetAUTOFOCUS(b bool) *TemplateHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *TemplateHTMLElement) CLASS(v string) *TemplateHTMLElement {
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

func (e *TemplateHTMLElement) IfCLASS(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *TemplateHTMLElement) SetCLASS(v string) *TemplateHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *TemplateHTMLElement) RemoveCLASS(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) CONTENTEDITABLE(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *TemplateHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *TemplateHTMLElement) RemoveCONTENTEDITABLE(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) DIR(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *TemplateHTMLElement) IfDIR(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *TemplateHTMLElement) RemoveDIR(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *TemplateHTMLElement) DRAGGABLE(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *TemplateHTMLElement) IfDRAGGABLE(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *TemplateHTMLElement) RemoveDRAGGABLE(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) ENTERKEYHINT(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *TemplateHTMLElement) IfENTERKEYHINT(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *TemplateHTMLElement) RemoveENTERKEYHINT(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) HIDDEN(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *TemplateHTMLElement) IfHIDDEN(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *TemplateHTMLElement) RemoveHIDDEN(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *TemplateHTMLElement) ID(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *TemplateHTMLElement) IfID(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *TemplateHTMLElement) RemoveID(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *TemplateHTMLElement) INERT() *TemplateHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *TemplateHTMLElement) IfINERT(cond bool) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *TemplateHTMLElement) RemoveINERT() *TemplateHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *TemplateHTMLElement) SetINERT(b bool) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) INPUTMODE(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *TemplateHTMLElement) IfINPUTMODE(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *TemplateHTMLElement) RemoveINPUTMODE(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *TemplateHTMLElement) IS(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *TemplateHTMLElement) IfIS(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *TemplateHTMLElement) RemoveIS(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *TemplateHTMLElement) ITEMID(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *TemplateHTMLElement) IfITEMID(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *TemplateHTMLElement) RemoveITEMID(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *TemplateHTMLElement) ITEMPROP(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *TemplateHTMLElement) IfITEMPROP(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *TemplateHTMLElement) RemoveITEMPROP(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *TemplateHTMLElement) ITEMREF(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *TemplateHTMLElement) IfITEMREF(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *TemplateHTMLElement) RemoveITEMREF(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *TemplateHTMLElement) ITEMSCOPE() *TemplateHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *TemplateHTMLElement) IfITEMSCOPE(cond bool) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *TemplateHTMLElement) RemoveITEMSCOPE() *TemplateHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *TemplateHTMLElement) SetITEMSCOPE(b bool) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) ITEMTYPE(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *TemplateHTMLElement) IfITEMTYPE(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *TemplateHTMLElement) RemoveITEMTYPE(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TemplateHTMLElement) LANG(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *TemplateHTMLElement) IfLANG(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *TemplateHTMLElement) RemoveLANG(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *TemplateHTMLElement) NONCE(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *TemplateHTMLElement) IfNONCE(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *TemplateHTMLElement) RemoveNONCE(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) POPOVER(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *TemplateHTMLElement) IfPOPOVER(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *TemplateHTMLElement) RemovePOPOVER(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SHADOWROOTDELEGATESFOCUS sets the "shadowrootdelegatesfocus" attribute.
// Sets delegates focus on a declarative shadow root
// Values values are constrained to:
//   - boolean_attribute
func (e *TemplateHTMLElement) SHADOWROOTDELEGATESFOCUS() *TemplateHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeSHADOWROOTDELEGATESFOCUSKey] = struct{}{}
	return e
}

func (e *TemplateHTMLElement) IfSHADOWROOTDELEGATESFOCUS(cond bool) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.SHADOWROOTDELEGATESFOCUS()
}

func (e *TemplateHTMLElement) RemoveSHADOWROOTDELEGATESFOCUS() *TemplateHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeSHADOWROOTDELEGATESFOCUSKey)
	return e
}

func (e *TemplateHTMLElement) SetSHADOWROOTDELEGATESFOCUS(b bool) *TemplateHTMLElement {
	if b {
		return e.SHADOWROOTDELEGATESFOCUS()
	}
	return e.RemoveSHADOWROOTDELEGATESFOCUS()
}

// SHADOWROOTMODE sets the "shadowrootmode" attribute.
// Enables streaming declarative shadow roots
// Values values are constrained to:
//   - open
//   - closed
func (e *TemplateHTMLElement) SHADOWROOTMODE(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSHADOWROOTMODEKey] = v
	return e
}

func (e *TemplateHTMLElement) IfSHADOWROOTMODE(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.SHADOWROOTMODE(v)
}

func (e *TemplateHTMLElement) RemoveSHADOWROOTMODE(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeSHADOWROOTMODEKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *TemplateHTMLElement) SLOT(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *TemplateHTMLElement) IfSLOT(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *TemplateHTMLElement) RemoveSLOT(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *TemplateHTMLElement) SPELLCHECK(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *TemplateHTMLElement) IfSPELLCHECK(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *TemplateHTMLElement) RemoveSPELLCHECK(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TemplateHTMLElement) STYLE(k, v string) *TemplateHTMLElement {
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

func (e *TemplateHTMLElement) IfSTYLE(cond bool, k string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *TemplateHTMLElement) RemoveSTYLE(k string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) TABINDEX(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *TemplateHTMLElement) IfTABINDEX(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *TemplateHTMLElement) RemoveTABINDEX(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *TemplateHTMLElement) TITLE(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *TemplateHTMLElement) IfTITLE(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *TemplateHTMLElement) RemoveTITLE(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *TemplateHTMLElement) TRANSLATE(v string) *TemplateHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *TemplateHTMLElement) IfTRANSLATE(cond bool, v string) *TemplateHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *TemplateHTMLElement) RemoveTRANSLATE(v string) *TemplateHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
