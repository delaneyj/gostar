package html

import (
	"fmt"
)

type AHTMLElement struct {
	*Element
}

func A(children ...ElementBuilder) *AHTMLElement {
	return &AHTMLElement{
		Element: &Element{
			Tag:           elementTagA,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *AHTMLElement) Children(children ...ElementBuilder) *AHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *AHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *AHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *AHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *AHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *AHTMLElement) Text(text string) *AHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *AHTMLElement) TextF(format string, args ...any) *AHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *AHTMLElement) Escaped(text string) *AHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *AHTMLElement) EscapedF(format string, args ...any) *AHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *AHTMLElement) CustomData(key, value string) *AHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AHTMLElement) CustomDataRemove(key string) *AHTMLElement {
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
func (e *AHTMLElement) ACCESSKEY(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *AHTMLElement) IfACCESSKEY(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *AHTMLElement) RemoveACCESSKEY(v string) *AHTMLElement {
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
func (e *AHTMLElement) AUTOCAPITALIZE(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *AHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *AHTMLElement) RemoveAUTOCAPITALIZE(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *AHTMLElement) AUTOFOCUS() *AHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *AHTMLElement) IfAUTOFOCUS(cond bool) *AHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *AHTMLElement) RemoveAUTOFOCUS() *AHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *AHTMLElement) SetAUTOFOCUS(b bool) *AHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *AHTMLElement) CLASS(v string) *AHTMLElement {
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

func (e *AHTMLElement) IfCLASS(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *AHTMLElement) SetCLASS(v string) *AHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *AHTMLElement) RemoveCLASS(v string) *AHTMLElement {
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
func (e *AHTMLElement) CONTENTEDITABLE(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *AHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *AHTMLElement) RemoveCONTENTEDITABLE(v string) *AHTMLElement {
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
func (e *AHTMLElement) DIR(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *AHTMLElement) IfDIR(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *AHTMLElement) RemoveDIR(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DOWNLOAD sets the "download" attribute.
// Whether to download the resource instead of navigating to it, and its filename if so
// Values values are constrained to:
func (e *AHTMLElement) DOWNLOAD(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDOWNLOADKey] = v
	return e
}

func (e *AHTMLElement) IfDOWNLOAD(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.DOWNLOAD(v)
}

func (e *AHTMLElement) RemoveDOWNLOAD(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeDOWNLOADKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *AHTMLElement) DRAGGABLE(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *AHTMLElement) IfDRAGGABLE(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *AHTMLElement) RemoveDRAGGABLE(v string) *AHTMLElement {
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
func (e *AHTMLElement) ENTERKEYHINT(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *AHTMLElement) IfENTERKEYHINT(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *AHTMLElement) RemoveENTERKEYHINT(v string) *AHTMLElement {
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
func (e *AHTMLElement) HIDDEN(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *AHTMLElement) IfHIDDEN(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *AHTMLElement) RemoveHIDDEN(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// HREF sets the "href" attribute.
// Document base URL
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *AHTMLElement) HREF(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHREFKey] = v
	return e
}

func (e *AHTMLElement) IfHREF(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.HREF(v)
}

func (e *AHTMLElement) RemoveHREF(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeHREFKey)
	return e
}

// HREFLANG sets the "hreflang" attribute.
// Language of the linked resource
// Values values are constrained to:
func (e *AHTMLElement) HREFLANG(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHREFLANGKey] = v
	return e
}

func (e *AHTMLElement) IfHREFLANG(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.HREFLANG(v)
}

func (e *AHTMLElement) RemoveHREFLANG(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeHREFLANGKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *AHTMLElement) ID(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *AHTMLElement) IfID(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *AHTMLElement) RemoveID(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *AHTMLElement) INERT() *AHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *AHTMLElement) IfINERT(cond bool) *AHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *AHTMLElement) RemoveINERT() *AHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *AHTMLElement) SetINERT(b bool) *AHTMLElement {
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
func (e *AHTMLElement) INPUTMODE(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *AHTMLElement) IfINPUTMODE(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *AHTMLElement) RemoveINPUTMODE(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *AHTMLElement) IS(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *AHTMLElement) IfIS(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *AHTMLElement) RemoveIS(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *AHTMLElement) ITEMID(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *AHTMLElement) IfITEMID(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *AHTMLElement) RemoveITEMID(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *AHTMLElement) ITEMPROP(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *AHTMLElement) IfITEMPROP(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *AHTMLElement) RemoveITEMPROP(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *AHTMLElement) ITEMREF(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *AHTMLElement) IfITEMREF(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *AHTMLElement) RemoveITEMREF(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *AHTMLElement) ITEMSCOPE() *AHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *AHTMLElement) IfITEMSCOPE(cond bool) *AHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *AHTMLElement) RemoveITEMSCOPE() *AHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *AHTMLElement) SetITEMSCOPE(b bool) *AHTMLElement {
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
func (e *AHTMLElement) ITEMTYPE(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *AHTMLElement) IfITEMTYPE(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *AHTMLElement) RemoveITEMTYPE(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AHTMLElement) LANG(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *AHTMLElement) IfLANG(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *AHTMLElement) RemoveLANG(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *AHTMLElement) NONCE(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *AHTMLElement) IfNONCE(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *AHTMLElement) RemoveNONCE(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// PING sets the "ping" attribute.
// URLs to ping
// Values values are constrained to:
//   - set_of_space_separated_tokens
//   - valid_non_empty_ur_ls
func (e *AHTMLElement) PING(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePINGKey] = v
	return e
}

func (e *AHTMLElement) IfPING(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.PING(v)
}

func (e *AHTMLElement) RemovePING(v string) *AHTMLElement {
	delete(e.StringAttributes, attributePINGKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - auto
//   - manual
//   - manual
func (e *AHTMLElement) POPOVER(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *AHTMLElement) IfPOPOVER(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *AHTMLElement) RemovePOPOVER(v string) *AHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//   - referrer_policy
func (e *AHTMLElement) REFERRERPOLICY(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeREFERRERPOLICYKey] = v
	return e
}

func (e *AHTMLElement) IfREFERRERPOLICY(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.REFERRERPOLICY(v)
}

func (e *AHTMLElement) RemoveREFERRERPOLICY(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeREFERRERPOLICYKey)
	return e
}

// REL sets the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *AHTMLElement) REL(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeRELKey] = v
	return e
}

func (e *AHTMLElement) IfREL(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.REL(v)
}

func (e *AHTMLElement) RemoveREL(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeRELKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *AHTMLElement) SLOT(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *AHTMLElement) IfSLOT(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *AHTMLElement) RemoveSLOT(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *AHTMLElement) SPELLCHECK(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *AHTMLElement) IfSPELLCHECK(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *AHTMLElement) RemoveSPELLCHECK(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AHTMLElement) STYLE(k, v string) *AHTMLElement {
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

func (e *AHTMLElement) IfSTYLE(cond bool, k string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *AHTMLElement) RemoveSTYLE(k string) *AHTMLElement {
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
func (e *AHTMLElement) TABINDEX(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *AHTMLElement) IfTABINDEX(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *AHTMLElement) RemoveTABINDEX(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TARGET sets the "target" attribute.
// Navigable for form submission
// Values values are constrained to:
//   - valid_navigable_target_name_or_keyword
func (e *AHTMLElement) TARGET(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTARGETKey] = v
	return e
}

func (e *AHTMLElement) IfTARGET(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.TARGET(v)
}

func (e *AHTMLElement) RemoveTARGET(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeTARGETKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *AHTMLElement) TITLE(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *AHTMLElement) IfTITLE(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *AHTMLElement) RemoveTITLE(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *AHTMLElement) TRANSLATE(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *AHTMLElement) IfTRANSLATE(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *AHTMLElement) RemoveTRANSLATE(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//   - module
//   - valid_mime_type_string
//   - java_script_mime_type_essence_match
func (e *AHTMLElement) TYPE(v string) *AHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTYPEKey] = v
	return e
}

func (e *AHTMLElement) IfTYPE(cond bool, v string) *AHTMLElement {
	if !cond {
		return e
	}
	return e.TYPE(v)
}

func (e *AHTMLElement) RemoveTYPE(v string) *AHTMLElement {
	delete(e.StringAttributes, attributeTYPEKey)
	return e
}
