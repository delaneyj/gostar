package html

import (
	"fmt"
)

type BaseHTMLElement struct {
	*Element
}

func BASE(children ...ElementBuilder) *BaseHTMLElement {
	return &BaseHTMLElement{
		Element: &Element{
			Tag:           elementTagBASE,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *BaseHTMLElement) Children(children ...ElementBuilder) *BaseHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *BaseHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *BaseHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *BaseHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *BaseHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *BaseHTMLElement) Text(text string) *BaseHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *BaseHTMLElement) TextF(format string, args ...any) *BaseHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *BaseHTMLElement) Escaped(text string) *BaseHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *BaseHTMLElement) EscapedF(format string, args ...any) *BaseHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *BaseHTMLElement) CustomData(key, value string) *BaseHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *BaseHTMLElement) CustomDataRemove(key string) *BaseHTMLElement {
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
func (e *BaseHTMLElement) ACCESSKEY(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *BaseHTMLElement) IfACCESSKEY(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *BaseHTMLElement) RemoveACCESSKEY(v string) *BaseHTMLElement {
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
func (e *BaseHTMLElement) AUTOCAPITALIZE(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *BaseHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *BaseHTMLElement) RemoveAUTOCAPITALIZE(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *BaseHTMLElement) AUTOFOCUS() *BaseHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *BaseHTMLElement) IfAUTOFOCUS(cond bool) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *BaseHTMLElement) RemoveAUTOFOCUS() *BaseHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *BaseHTMLElement) SetAUTOFOCUS(b bool) *BaseHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *BaseHTMLElement) CLASS(v string) *BaseHTMLElement {
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

func (e *BaseHTMLElement) IfCLASS(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *BaseHTMLElement) SetCLASS(v string) *BaseHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *BaseHTMLElement) RemoveCLASS(v string) *BaseHTMLElement {
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
func (e *BaseHTMLElement) CONTENTEDITABLE(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *BaseHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *BaseHTMLElement) RemoveCONTENTEDITABLE(v string) *BaseHTMLElement {
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
func (e *BaseHTMLElement) DIR(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *BaseHTMLElement) IfDIR(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *BaseHTMLElement) RemoveDIR(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *BaseHTMLElement) DRAGGABLE(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *BaseHTMLElement) IfDRAGGABLE(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *BaseHTMLElement) RemoveDRAGGABLE(v string) *BaseHTMLElement {
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
func (e *BaseHTMLElement) ENTERKEYHINT(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *BaseHTMLElement) IfENTERKEYHINT(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *BaseHTMLElement) RemoveENTERKEYHINT(v string) *BaseHTMLElement {
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
func (e *BaseHTMLElement) HIDDEN(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *BaseHTMLElement) IfHIDDEN(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *BaseHTMLElement) RemoveHIDDEN(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// HREF sets the "href" attribute.
// Document base URL
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *BaseHTMLElement) HREF(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHREFKey] = v
	return e
}

func (e *BaseHTMLElement) IfHREF(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.HREF(v)
}

func (e *BaseHTMLElement) RemoveHREF(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeHREFKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *BaseHTMLElement) ID(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *BaseHTMLElement) IfID(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *BaseHTMLElement) RemoveID(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *BaseHTMLElement) INERT() *BaseHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *BaseHTMLElement) IfINERT(cond bool) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *BaseHTMLElement) RemoveINERT() *BaseHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *BaseHTMLElement) SetINERT(b bool) *BaseHTMLElement {
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
func (e *BaseHTMLElement) INPUTMODE(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *BaseHTMLElement) IfINPUTMODE(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *BaseHTMLElement) RemoveINPUTMODE(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *BaseHTMLElement) IS(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *BaseHTMLElement) IfIS(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *BaseHTMLElement) RemoveIS(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *BaseHTMLElement) ITEMID(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *BaseHTMLElement) IfITEMID(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *BaseHTMLElement) RemoveITEMID(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *BaseHTMLElement) ITEMPROP(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *BaseHTMLElement) IfITEMPROP(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *BaseHTMLElement) RemoveITEMPROP(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *BaseHTMLElement) ITEMREF(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *BaseHTMLElement) IfITEMREF(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *BaseHTMLElement) RemoveITEMREF(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *BaseHTMLElement) ITEMSCOPE() *BaseHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *BaseHTMLElement) IfITEMSCOPE(cond bool) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *BaseHTMLElement) RemoveITEMSCOPE() *BaseHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *BaseHTMLElement) SetITEMSCOPE(b bool) *BaseHTMLElement {
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
func (e *BaseHTMLElement) ITEMTYPE(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *BaseHTMLElement) IfITEMTYPE(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *BaseHTMLElement) RemoveITEMTYPE(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *BaseHTMLElement) LANG(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *BaseHTMLElement) IfLANG(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *BaseHTMLElement) RemoveLANG(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *BaseHTMLElement) NONCE(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *BaseHTMLElement) IfNONCE(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *BaseHTMLElement) RemoveNONCE(v string) *BaseHTMLElement {
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
func (e *BaseHTMLElement) POPOVER(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *BaseHTMLElement) IfPOPOVER(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *BaseHTMLElement) RemovePOPOVER(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *BaseHTMLElement) SLOT(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *BaseHTMLElement) IfSLOT(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *BaseHTMLElement) RemoveSLOT(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *BaseHTMLElement) SPELLCHECK(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *BaseHTMLElement) IfSPELLCHECK(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *BaseHTMLElement) RemoveSPELLCHECK(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *BaseHTMLElement) STYLE(k, v string) *BaseHTMLElement {
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

func (e *BaseHTMLElement) IfSTYLE(cond bool, k string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *BaseHTMLElement) RemoveSTYLE(k string) *BaseHTMLElement {
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
func (e *BaseHTMLElement) TABINDEX(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *BaseHTMLElement) IfTABINDEX(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *BaseHTMLElement) RemoveTABINDEX(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TARGET sets the "target" attribute.
// Navigable for form submission
// Values values are constrained to:
//   - valid_navigable_target_name_or_keyword
func (e *BaseHTMLElement) TARGET(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTARGETKey] = v
	return e
}

func (e *BaseHTMLElement) IfTARGET(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.TARGET(v)
}

func (e *BaseHTMLElement) RemoveTARGET(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeTARGETKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *BaseHTMLElement) TITLE(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *BaseHTMLElement) IfTITLE(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *BaseHTMLElement) RemoveTITLE(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *BaseHTMLElement) TRANSLATE(v string) *BaseHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *BaseHTMLElement) IfTRANSLATE(cond bool, v string) *BaseHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *BaseHTMLElement) RemoveTRANSLATE(v string) *BaseHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
