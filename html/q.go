package html

import (
	"fmt"
)

type QHTMLElement struct {
	*Element
}

func Q(children ...ElementBuilder) *QHTMLElement {
	return &QHTMLElement{
		Element: &Element{
			Tag:           elementTagQ,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *QHTMLElement) Children(children ...ElementBuilder) *QHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *QHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *QHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *QHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *QHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *QHTMLElement) Text(text string) *QHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *QHTMLElement) TextF(format string, args ...any) *QHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *QHTMLElement) Escaped(text string) *QHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *QHTMLElement) EscapedF(format string, args ...any) *QHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *QHTMLElement) CustomData(key, value string) *QHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *QHTMLElement) CustomDataRemove(key string) *QHTMLElement {
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
func (e *QHTMLElement) ACCESSKEY(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *QHTMLElement) IfACCESSKEY(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *QHTMLElement) RemoveACCESSKEY(v string) *QHTMLElement {
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
func (e *QHTMLElement) AUTOCAPITALIZE(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *QHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *QHTMLElement) RemoveAUTOCAPITALIZE(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *QHTMLElement) AUTOFOCUS() *QHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *QHTMLElement) IfAUTOFOCUS(cond bool) *QHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *QHTMLElement) RemoveAUTOFOCUS() *QHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *QHTMLElement) SetAUTOFOCUS(b bool) *QHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CITE sets the "cite" attribute.
// Link to the source of the quotation or more information about the edit
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *QHTMLElement) CITE(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCITEKey] = v
	return e
}

func (e *QHTMLElement) IfCITE(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.CITE(v)
}

func (e *QHTMLElement) RemoveCITE(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeCITEKey)
	return e
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *QHTMLElement) CLASS(v string) *QHTMLElement {
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

func (e *QHTMLElement) IfCLASS(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *QHTMLElement) SetCLASS(v string) *QHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *QHTMLElement) RemoveCLASS(v string) *QHTMLElement {
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
func (e *QHTMLElement) CONTENTEDITABLE(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *QHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *QHTMLElement) RemoveCONTENTEDITABLE(v string) *QHTMLElement {
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
func (e *QHTMLElement) DIR(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *QHTMLElement) IfDIR(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *QHTMLElement) RemoveDIR(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *QHTMLElement) DRAGGABLE(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *QHTMLElement) IfDRAGGABLE(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *QHTMLElement) RemoveDRAGGABLE(v string) *QHTMLElement {
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
func (e *QHTMLElement) ENTERKEYHINT(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *QHTMLElement) IfENTERKEYHINT(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *QHTMLElement) RemoveENTERKEYHINT(v string) *QHTMLElement {
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
func (e *QHTMLElement) HIDDEN(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *QHTMLElement) IfHIDDEN(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *QHTMLElement) RemoveHIDDEN(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *QHTMLElement) ID(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *QHTMLElement) IfID(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *QHTMLElement) RemoveID(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *QHTMLElement) INERT() *QHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *QHTMLElement) IfINERT(cond bool) *QHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *QHTMLElement) RemoveINERT() *QHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *QHTMLElement) SetINERT(b bool) *QHTMLElement {
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
func (e *QHTMLElement) INPUTMODE(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *QHTMLElement) IfINPUTMODE(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *QHTMLElement) RemoveINPUTMODE(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *QHTMLElement) IS(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *QHTMLElement) IfIS(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *QHTMLElement) RemoveIS(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *QHTMLElement) ITEMID(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *QHTMLElement) IfITEMID(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *QHTMLElement) RemoveITEMID(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *QHTMLElement) ITEMPROP(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *QHTMLElement) IfITEMPROP(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *QHTMLElement) RemoveITEMPROP(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *QHTMLElement) ITEMREF(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *QHTMLElement) IfITEMREF(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *QHTMLElement) RemoveITEMREF(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *QHTMLElement) ITEMSCOPE() *QHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *QHTMLElement) IfITEMSCOPE(cond bool) *QHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *QHTMLElement) RemoveITEMSCOPE() *QHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *QHTMLElement) SetITEMSCOPE(b bool) *QHTMLElement {
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
func (e *QHTMLElement) ITEMTYPE(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *QHTMLElement) IfITEMTYPE(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *QHTMLElement) RemoveITEMTYPE(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *QHTMLElement) LANG(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *QHTMLElement) IfLANG(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *QHTMLElement) RemoveLANG(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *QHTMLElement) NONCE(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *QHTMLElement) IfNONCE(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *QHTMLElement) RemoveNONCE(v string) *QHTMLElement {
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
func (e *QHTMLElement) POPOVER(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *QHTMLElement) IfPOPOVER(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *QHTMLElement) RemovePOPOVER(v string) *QHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *QHTMLElement) SLOT(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *QHTMLElement) IfSLOT(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *QHTMLElement) RemoveSLOT(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *QHTMLElement) SPELLCHECK(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *QHTMLElement) IfSPELLCHECK(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *QHTMLElement) RemoveSPELLCHECK(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *QHTMLElement) STYLE(k, v string) *QHTMLElement {
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

func (e *QHTMLElement) IfSTYLE(cond bool, k string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *QHTMLElement) RemoveSTYLE(k string) *QHTMLElement {
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
func (e *QHTMLElement) TABINDEX(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *QHTMLElement) IfTABINDEX(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *QHTMLElement) RemoveTABINDEX(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *QHTMLElement) TITLE(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *QHTMLElement) IfTITLE(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *QHTMLElement) RemoveTITLE(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *QHTMLElement) TRANSLATE(v string) *QHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *QHTMLElement) IfTRANSLATE(cond bool, v string) *QHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *QHTMLElement) RemoveTRANSLATE(v string) *QHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
