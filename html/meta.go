package html

import (
	"fmt"
)

type MetaHTMLElement struct {
	*Element
}

func META(children ...ElementBuilder) *MetaHTMLElement {
	return &MetaHTMLElement{
		Element: &Element{
			Tag:           elementTagMETA,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *MetaHTMLElement) Children(children ...ElementBuilder) *MetaHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MetaHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *MetaHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MetaHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *MetaHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MetaHTMLElement) Text(text string) *MetaHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MetaHTMLElement) TextF(format string, args ...any) *MetaHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MetaHTMLElement) Escaped(text string) *MetaHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MetaHTMLElement) EscapedF(format string, args ...any) *MetaHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MetaHTMLElement) CustomData(key, value string) *MetaHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *MetaHTMLElement) CustomDataRemove(key string) *MetaHTMLElement {
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
func (e *MetaHTMLElement) ACCESSKEY(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *MetaHTMLElement) IfACCESSKEY(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *MetaHTMLElement) RemoveACCESSKEY(v string) *MetaHTMLElement {
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
func (e *MetaHTMLElement) AUTOCAPITALIZE(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *MetaHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *MetaHTMLElement) RemoveAUTOCAPITALIZE(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *MetaHTMLElement) AUTOFOCUS() *MetaHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *MetaHTMLElement) IfAUTOFOCUS(cond bool) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *MetaHTMLElement) RemoveAUTOFOCUS() *MetaHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *MetaHTMLElement) SetAUTOFOCUS(b bool) *MetaHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CHARSET sets the "charset" attribute.
// Character encoding declaration
// Values values are constrained to:
//   - utf_8
func (e *MetaHTMLElement) CHARSET(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCHARSETKey] = v
	return e
}

func (e *MetaHTMLElement) IfCHARSET(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.CHARSET(v)
}

func (e *MetaHTMLElement) RemoveCHARSET(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeCHARSETKey)
	return e
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *MetaHTMLElement) CLASS(v string) *MetaHTMLElement {
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

func (e *MetaHTMLElement) IfCLASS(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *MetaHTMLElement) SetCLASS(v string) *MetaHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *MetaHTMLElement) RemoveCLASS(v string) *MetaHTMLElement {
	kv, ok := e.DelimitedStringAttributes[attributeCLASSKey]
	if !ok {
		return e
	}
	kv.Remove(v)
	return e
}

// CONTENT sets the "content" attribute.
// Value of the element
// Values values are constrained to:
//   - text
func (e *MetaHTMLElement) CONTENT(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTKey] = v
	return e
}

func (e *MetaHTMLElement) IfCONTENT(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENT(v)
}

func (e *MetaHTMLElement) RemoveCONTENT(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeCONTENTKey)
	return e
}

// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//   - false
//   - plaintext_only
//   - true
func (e *MetaHTMLElement) CONTENTEDITABLE(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *MetaHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *MetaHTMLElement) RemoveCONTENTEDITABLE(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *MetaHTMLElement) DIR(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *MetaHTMLElement) IfDIR(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *MetaHTMLElement) RemoveDIR(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *MetaHTMLElement) DRAGGABLE(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *MetaHTMLElement) IfDRAGGABLE(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *MetaHTMLElement) RemoveDRAGGABLE(v string) *MetaHTMLElement {
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
func (e *MetaHTMLElement) ENTERKEYHINT(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *MetaHTMLElement) IfENTERKEYHINT(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *MetaHTMLElement) RemoveENTERKEYHINT(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *MetaHTMLElement) HIDDEN(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *MetaHTMLElement) IfHIDDEN(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *MetaHTMLElement) RemoveHIDDEN(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// HTTP_EQUIV sets the "http-equiv" attribute.
// Pragma directive
// Values values are constrained to:
//   - content_security_policy
//   - content_type
//   - default_style
//   - refresh
//   - x_ua_compatible
func (e *MetaHTMLElement) HTTP_EQUIV(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHTTP_EQUIVKey] = v
	return e
}

func (e *MetaHTMLElement) IfHTTP_EQUIV(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.HTTP_EQUIV(v)
}

func (e *MetaHTMLElement) RemoveHTTP_EQUIV(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeHTTP_EQUIVKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *MetaHTMLElement) ID(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *MetaHTMLElement) IfID(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *MetaHTMLElement) RemoveID(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *MetaHTMLElement) INERT() *MetaHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *MetaHTMLElement) IfINERT(cond bool) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *MetaHTMLElement) RemoveINERT() *MetaHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *MetaHTMLElement) SetINERT(b bool) *MetaHTMLElement {
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
func (e *MetaHTMLElement) INPUTMODE(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *MetaHTMLElement) IfINPUTMODE(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *MetaHTMLElement) RemoveINPUTMODE(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *MetaHTMLElement) IS(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *MetaHTMLElement) IfIS(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *MetaHTMLElement) RemoveIS(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *MetaHTMLElement) ITEMID(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *MetaHTMLElement) IfITEMID(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *MetaHTMLElement) RemoveITEMID(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *MetaHTMLElement) ITEMPROP(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *MetaHTMLElement) IfITEMPROP(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *MetaHTMLElement) RemoveITEMPROP(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *MetaHTMLElement) ITEMREF(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *MetaHTMLElement) IfITEMREF(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *MetaHTMLElement) RemoveITEMREF(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *MetaHTMLElement) ITEMSCOPE() *MetaHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *MetaHTMLElement) IfITEMSCOPE(cond bool) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *MetaHTMLElement) RemoveITEMSCOPE() *MetaHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *MetaHTMLElement) SetITEMSCOPE(b bool) *MetaHTMLElement {
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
func (e *MetaHTMLElement) ITEMTYPE(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *MetaHTMLElement) IfITEMTYPE(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *MetaHTMLElement) RemoveITEMTYPE(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *MetaHTMLElement) LANG(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *MetaHTMLElement) IfLANG(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *MetaHTMLElement) RemoveLANG(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// MEDIA sets the "media" attribute.
// Applicable media
// Values values are constrained to:
//   - valid_media_query_list
func (e *MetaHTMLElement) MEDIA(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMEDIAKey] = v
	return e
}

func (e *MetaHTMLElement) IfMEDIA(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.MEDIA(v)
}

func (e *MetaHTMLElement) RemoveMEDIA(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeMEDIAKey)
	return e
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *MetaHTMLElement) NAME(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *MetaHTMLElement) IfNAME(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *MetaHTMLElement) RemoveNAME(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *MetaHTMLElement) NONCE(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *MetaHTMLElement) IfNONCE(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *MetaHTMLElement) RemoveNONCE(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *MetaHTMLElement) POPOVER(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *MetaHTMLElement) IfPOPOVER(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *MetaHTMLElement) RemovePOPOVER(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *MetaHTMLElement) SLOT(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *MetaHTMLElement) IfSLOT(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *MetaHTMLElement) RemoveSLOT(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *MetaHTMLElement) SPELLCHECK(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *MetaHTMLElement) IfSPELLCHECK(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *MetaHTMLElement) RemoveSPELLCHECK(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *MetaHTMLElement) STYLE(k, v string) *MetaHTMLElement {
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

func (e *MetaHTMLElement) IfSTYLE(cond bool, k string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *MetaHTMLElement) RemoveSTYLE(k string) *MetaHTMLElement {
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
func (e *MetaHTMLElement) TABINDEX(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *MetaHTMLElement) IfTABINDEX(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *MetaHTMLElement) RemoveTABINDEX(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *MetaHTMLElement) TITLE(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *MetaHTMLElement) IfTITLE(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *MetaHTMLElement) RemoveTITLE(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *MetaHTMLElement) TRANSLATE(v string) *MetaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *MetaHTMLElement) IfTRANSLATE(cond bool, v string) *MetaHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *MetaHTMLElement) RemoveTRANSLATE(v string) *MetaHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
