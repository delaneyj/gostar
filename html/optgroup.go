package html

import (
	"fmt"
)

type OptgroupHTMLElement struct {
	*Element
}

func OPTGROUP(children ...ElementBuilder) *OptgroupHTMLElement {
	return &OptgroupHTMLElement{
		Element: &Element{
			Tag:           elementTagOPTGROUP,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *OptgroupHTMLElement) Children(children ...ElementBuilder) *OptgroupHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *OptgroupHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *OptgroupHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *OptgroupHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *OptgroupHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *OptgroupHTMLElement) Text(text string) *OptgroupHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *OptgroupHTMLElement) TextF(format string, args ...any) *OptgroupHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *OptgroupHTMLElement) Escaped(text string) *OptgroupHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *OptgroupHTMLElement) EscapedF(format string, args ...any) *OptgroupHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *OptgroupHTMLElement) CustomData(key, value string) *OptgroupHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *OptgroupHTMLElement) CustomDataRemove(key string) *OptgroupHTMLElement {
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
func (e *OptgroupHTMLElement) ACCESSKEY(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfACCESSKEY(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *OptgroupHTMLElement) RemoveACCESSKEY(v string) *OptgroupHTMLElement {
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
func (e *OptgroupHTMLElement) AUTOCAPITALIZE(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *OptgroupHTMLElement) RemoveAUTOCAPITALIZE(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *OptgroupHTMLElement) AUTOFOCUS() *OptgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *OptgroupHTMLElement) IfAUTOFOCUS(cond bool) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *OptgroupHTMLElement) RemoveAUTOFOCUS() *OptgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *OptgroupHTMLElement) SetAUTOFOCUS(b bool) *OptgroupHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *OptgroupHTMLElement) CLASS(v string) *OptgroupHTMLElement {
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

func (e *OptgroupHTMLElement) IfCLASS(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *OptgroupHTMLElement) SetCLASS(v string) *OptgroupHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *OptgroupHTMLElement) RemoveCLASS(v string) *OptgroupHTMLElement {
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
func (e *OptgroupHTMLElement) CONTENTEDITABLE(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *OptgroupHTMLElement) RemoveCONTENTEDITABLE(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *OptgroupHTMLElement) DIR(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfDIR(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *OptgroupHTMLElement) RemoveDIR(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//   - boolean_attribute
func (e *OptgroupHTMLElement) DISABLED() *OptgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeDISABLEDKey] = struct{}{}
	return e
}

func (e *OptgroupHTMLElement) IfDISABLED(cond bool) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.DISABLED()
}

func (e *OptgroupHTMLElement) RemoveDISABLED() *OptgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeDISABLEDKey)
	return e
}

func (e *OptgroupHTMLElement) SetDISABLED(b bool) *OptgroupHTMLElement {
	if b {
		return e.DISABLED()
	}
	return e.RemoveDISABLED()
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *OptgroupHTMLElement) DRAGGABLE(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfDRAGGABLE(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *OptgroupHTMLElement) RemoveDRAGGABLE(v string) *OptgroupHTMLElement {
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
func (e *OptgroupHTMLElement) ENTERKEYHINT(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfENTERKEYHINT(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *OptgroupHTMLElement) RemoveENTERKEYHINT(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *OptgroupHTMLElement) HIDDEN(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfHIDDEN(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *OptgroupHTMLElement) RemoveHIDDEN(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *OptgroupHTMLElement) ID(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfID(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *OptgroupHTMLElement) RemoveID(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *OptgroupHTMLElement) INERT() *OptgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *OptgroupHTMLElement) IfINERT(cond bool) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *OptgroupHTMLElement) RemoveINERT() *OptgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *OptgroupHTMLElement) SetINERT(b bool) *OptgroupHTMLElement {
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
func (e *OptgroupHTMLElement) INPUTMODE(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfINPUTMODE(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *OptgroupHTMLElement) RemoveINPUTMODE(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *OptgroupHTMLElement) IS(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfIS(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *OptgroupHTMLElement) RemoveIS(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *OptgroupHTMLElement) ITEMID(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfITEMID(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *OptgroupHTMLElement) RemoveITEMID(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *OptgroupHTMLElement) ITEMPROP(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfITEMPROP(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *OptgroupHTMLElement) RemoveITEMPROP(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *OptgroupHTMLElement) ITEMREF(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfITEMREF(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *OptgroupHTMLElement) RemoveITEMREF(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *OptgroupHTMLElement) ITEMSCOPE() *OptgroupHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *OptgroupHTMLElement) IfITEMSCOPE(cond bool) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *OptgroupHTMLElement) RemoveITEMSCOPE() *OptgroupHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *OptgroupHTMLElement) SetITEMSCOPE(b bool) *OptgroupHTMLElement {
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
func (e *OptgroupHTMLElement) ITEMTYPE(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfITEMTYPE(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *OptgroupHTMLElement) RemoveITEMTYPE(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LABEL sets the "label" attribute.
// User-visible label
// Values values are constrained to:
//   - text
func (e *OptgroupHTMLElement) LABEL(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLABELKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfLABEL(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.LABEL(v)
}

func (e *OptgroupHTMLElement) RemoveLABEL(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeLABELKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *OptgroupHTMLElement) LANG(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfLANG(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *OptgroupHTMLElement) RemoveLANG(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *OptgroupHTMLElement) NONCE(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfNONCE(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *OptgroupHTMLElement) RemoveNONCE(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *OptgroupHTMLElement) POPOVER(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfPOPOVER(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *OptgroupHTMLElement) RemovePOPOVER(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *OptgroupHTMLElement) SLOT(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfSLOT(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *OptgroupHTMLElement) RemoveSLOT(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *OptgroupHTMLElement) SPELLCHECK(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfSPELLCHECK(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *OptgroupHTMLElement) RemoveSPELLCHECK(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *OptgroupHTMLElement) STYLE(k, v string) *OptgroupHTMLElement {
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

func (e *OptgroupHTMLElement) IfSTYLE(cond bool, k string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *OptgroupHTMLElement) RemoveSTYLE(k string) *OptgroupHTMLElement {
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
func (e *OptgroupHTMLElement) TABINDEX(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfTABINDEX(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *OptgroupHTMLElement) RemoveTABINDEX(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *OptgroupHTMLElement) TITLE(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfTITLE(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *OptgroupHTMLElement) RemoveTITLE(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *OptgroupHTMLElement) TRANSLATE(v string) *OptgroupHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *OptgroupHTMLElement) IfTRANSLATE(cond bool, v string) *OptgroupHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *OptgroupHTMLElement) RemoveTRANSLATE(v string) *OptgroupHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
