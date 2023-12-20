package html

import (
	"fmt"
)

type StyleHTMLElement struct {
	*Element
}

func STYLE(children ...ElementBuilder) *StyleHTMLElement {
	return &StyleHTMLElement{
		Element: &Element{
			Tag:           elementTagSTYLE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *StyleHTMLElement) Children(children ...ElementBuilder) *StyleHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *StyleHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *StyleHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *StyleHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *StyleHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *StyleHTMLElement) Text(text string) *StyleHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *StyleHTMLElement) TextF(format string, args ...any) *StyleHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *StyleHTMLElement) Escaped(text string) *StyleHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *StyleHTMLElement) EscapedF(format string, args ...any) *StyleHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *StyleHTMLElement) CustomData(key, value string) *StyleHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *StyleHTMLElement) CustomDataRemove(key string) *StyleHTMLElement {
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
func (e *StyleHTMLElement) ACCESSKEY(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *StyleHTMLElement) IfACCESSKEY(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *StyleHTMLElement) RemoveACCESSKEY(v string) *StyleHTMLElement {
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
func (e *StyleHTMLElement) AUTOCAPITALIZE(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *StyleHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *StyleHTMLElement) RemoveAUTOCAPITALIZE(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *StyleHTMLElement) AUTOFOCUS() *StyleHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *StyleHTMLElement) IfAUTOFOCUS(cond bool) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *StyleHTMLElement) RemoveAUTOFOCUS() *StyleHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *StyleHTMLElement) SetAUTOFOCUS(b bool) *StyleHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// BLOCKING sets the "blocking" attribute.
// Whether the element is potentially render-blocking
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *StyleHTMLElement) BLOCKING(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeBLOCKINGKey] = v
	return e
}

func (e *StyleHTMLElement) IfBLOCKING(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.BLOCKING(v)
}

func (e *StyleHTMLElement) RemoveBLOCKING(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeBLOCKINGKey)
	return e
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *StyleHTMLElement) CLASS(v string) *StyleHTMLElement {
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

func (e *StyleHTMLElement) IfCLASS(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *StyleHTMLElement) SetCLASS(v string) *StyleHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *StyleHTMLElement) RemoveCLASS(v string) *StyleHTMLElement {
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
func (e *StyleHTMLElement) CONTENTEDITABLE(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *StyleHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *StyleHTMLElement) RemoveCONTENTEDITABLE(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *StyleHTMLElement) DIR(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *StyleHTMLElement) IfDIR(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *StyleHTMLElement) RemoveDIR(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *StyleHTMLElement) DRAGGABLE(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *StyleHTMLElement) IfDRAGGABLE(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *StyleHTMLElement) RemoveDRAGGABLE(v string) *StyleHTMLElement {
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
func (e *StyleHTMLElement) ENTERKEYHINT(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *StyleHTMLElement) IfENTERKEYHINT(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *StyleHTMLElement) RemoveENTERKEYHINT(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *StyleHTMLElement) HIDDEN(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *StyleHTMLElement) IfHIDDEN(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *StyleHTMLElement) RemoveHIDDEN(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *StyleHTMLElement) ID(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *StyleHTMLElement) IfID(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *StyleHTMLElement) RemoveID(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *StyleHTMLElement) INERT() *StyleHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *StyleHTMLElement) IfINERT(cond bool) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *StyleHTMLElement) RemoveINERT() *StyleHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *StyleHTMLElement) SetINERT(b bool) *StyleHTMLElement {
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
func (e *StyleHTMLElement) INPUTMODE(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *StyleHTMLElement) IfINPUTMODE(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *StyleHTMLElement) RemoveINPUTMODE(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *StyleHTMLElement) IS(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *StyleHTMLElement) IfIS(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *StyleHTMLElement) RemoveIS(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *StyleHTMLElement) ITEMID(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *StyleHTMLElement) IfITEMID(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *StyleHTMLElement) RemoveITEMID(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *StyleHTMLElement) ITEMPROP(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *StyleHTMLElement) IfITEMPROP(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *StyleHTMLElement) RemoveITEMPROP(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *StyleHTMLElement) ITEMREF(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *StyleHTMLElement) IfITEMREF(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *StyleHTMLElement) RemoveITEMREF(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *StyleHTMLElement) ITEMSCOPE() *StyleHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *StyleHTMLElement) IfITEMSCOPE(cond bool) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *StyleHTMLElement) RemoveITEMSCOPE() *StyleHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *StyleHTMLElement) SetITEMSCOPE(b bool) *StyleHTMLElement {
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
func (e *StyleHTMLElement) ITEMTYPE(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *StyleHTMLElement) IfITEMTYPE(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *StyleHTMLElement) RemoveITEMTYPE(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *StyleHTMLElement) LANG(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *StyleHTMLElement) IfLANG(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *StyleHTMLElement) RemoveLANG(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// MEDIA sets the "media" attribute.
// Applicable media
// Values values are constrained to:
//   - valid_media_query_list
func (e *StyleHTMLElement) MEDIA(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMEDIAKey] = v
	return e
}

func (e *StyleHTMLElement) IfMEDIA(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.MEDIA(v)
}

func (e *StyleHTMLElement) RemoveMEDIA(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeMEDIAKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *StyleHTMLElement) NONCE(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *StyleHTMLElement) IfNONCE(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *StyleHTMLElement) RemoveNONCE(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *StyleHTMLElement) POPOVER(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *StyleHTMLElement) IfPOPOVER(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *StyleHTMLElement) RemovePOPOVER(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *StyleHTMLElement) SLOT(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *StyleHTMLElement) IfSLOT(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *StyleHTMLElement) RemoveSLOT(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *StyleHTMLElement) SPELLCHECK(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *StyleHTMLElement) IfSPELLCHECK(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *StyleHTMLElement) RemoveSPELLCHECK(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *StyleHTMLElement) STYLE(k, v string) *StyleHTMLElement {
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

func (e *StyleHTMLElement) IfSTYLE(cond bool, k string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *StyleHTMLElement) RemoveSTYLE(k string) *StyleHTMLElement {
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
func (e *StyleHTMLElement) TABINDEX(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *StyleHTMLElement) IfTABINDEX(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *StyleHTMLElement) RemoveTABINDEX(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *StyleHTMLElement) TITLE(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *StyleHTMLElement) IfTITLE(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *StyleHTMLElement) RemoveTITLE(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *StyleHTMLElement) TRANSLATE(v string) *StyleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *StyleHTMLElement) IfTRANSLATE(cond bool, v string) *StyleHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *StyleHTMLElement) RemoveTRANSLATE(v string) *StyleHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
