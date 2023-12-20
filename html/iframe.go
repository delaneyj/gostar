package html

import (
	"fmt"
)

type IframeHTMLElement struct {
	*Element
}

func IFRAME(children ...ElementBuilder) *IframeHTMLElement {
	return &IframeHTMLElement{
		Element: &Element{
			Tag:           elementTagIFRAME,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *IframeHTMLElement) Children(children ...ElementBuilder) *IframeHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *IframeHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *IframeHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *IframeHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *IframeHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *IframeHTMLElement) Text(text string) *IframeHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *IframeHTMLElement) TextF(format string, args ...any) *IframeHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *IframeHTMLElement) Escaped(text string) *IframeHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *IframeHTMLElement) EscapedF(format string, args ...any) *IframeHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *IframeHTMLElement) CustomData(key, value string) *IframeHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *IframeHTMLElement) CustomDataRemove(key string) *IframeHTMLElement {
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
func (e *IframeHTMLElement) ACCESSKEY(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *IframeHTMLElement) IfACCESSKEY(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *IframeHTMLElement) RemoveACCESSKEY(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// ALLOW sets the "allow" attribute.
// Permissions policy to be applied to the iframe's contents
// Values values are constrained to:
//   - serialized_permissions_policy
func (e *IframeHTMLElement) ALLOW(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeALLOWKey] = v
	return e
}

func (e *IframeHTMLElement) IfALLOW(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ALLOW(v)
}

func (e *IframeHTMLElement) RemoveALLOW(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeALLOWKey)
	return e
}

// ALLOWFULLSCREEN sets the "allowfullscreen" attribute.
// Whether to allow the iframe's contents to use requestFullscreen()
// Values values are constrained to:
//   - boolean_attribute
func (e *IframeHTMLElement) ALLOWFULLSCREEN() *IframeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeALLOWFULLSCREENKey] = struct{}{}
	return e
}

func (e *IframeHTMLElement) IfALLOWFULLSCREEN(cond bool) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ALLOWFULLSCREEN()
}

func (e *IframeHTMLElement) RemoveALLOWFULLSCREEN() *IframeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeALLOWFULLSCREENKey)
	return e
}

func (e *IframeHTMLElement) SetALLOWFULLSCREEN(b bool) *IframeHTMLElement {
	if b {
		return e.ALLOWFULLSCREEN()
	}
	return e.RemoveALLOWFULLSCREEN()
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
func (e *IframeHTMLElement) AUTOCAPITALIZE(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *IframeHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *IframeHTMLElement) RemoveAUTOCAPITALIZE(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *IframeHTMLElement) AUTOFOCUS() *IframeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *IframeHTMLElement) IfAUTOFOCUS(cond bool) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *IframeHTMLElement) RemoveAUTOFOCUS() *IframeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *IframeHTMLElement) SetAUTOFOCUS(b bool) *IframeHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *IframeHTMLElement) CLASS(v string) *IframeHTMLElement {
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

func (e *IframeHTMLElement) IfCLASS(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *IframeHTMLElement) SetCLASS(v string) *IframeHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *IframeHTMLElement) RemoveCLASS(v string) *IframeHTMLElement {
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
func (e *IframeHTMLElement) CONTENTEDITABLE(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *IframeHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *IframeHTMLElement) RemoveCONTENTEDITABLE(v string) *IframeHTMLElement {
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
func (e *IframeHTMLElement) DIR(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *IframeHTMLElement) IfDIR(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *IframeHTMLElement) RemoveDIR(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *IframeHTMLElement) DRAGGABLE(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *IframeHTMLElement) IfDRAGGABLE(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *IframeHTMLElement) RemoveDRAGGABLE(v string) *IframeHTMLElement {
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
func (e *IframeHTMLElement) ENTERKEYHINT(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *IframeHTMLElement) IfENTERKEYHINT(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *IframeHTMLElement) RemoveENTERKEYHINT(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *IframeHTMLElement) HEIGHT(v int) *IframeHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeHEIGHTKey] = v
	return e
}

func (e *IframeHTMLElement) IfHEIGHT(cond bool, v int) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.HEIGHT(v)
}

func (e *IframeHTMLElement) RemoveHEIGHT(v int) *IframeHTMLElement {
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
func (e *IframeHTMLElement) HIDDEN(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *IframeHTMLElement) IfHIDDEN(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *IframeHTMLElement) RemoveHIDDEN(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *IframeHTMLElement) ID(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *IframeHTMLElement) IfID(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *IframeHTMLElement) RemoveID(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *IframeHTMLElement) INERT() *IframeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *IframeHTMLElement) IfINERT(cond bool) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *IframeHTMLElement) RemoveINERT() *IframeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *IframeHTMLElement) SetINERT(b bool) *IframeHTMLElement {
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
func (e *IframeHTMLElement) INPUTMODE(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *IframeHTMLElement) IfINPUTMODE(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *IframeHTMLElement) RemoveINPUTMODE(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *IframeHTMLElement) IS(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *IframeHTMLElement) IfIS(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *IframeHTMLElement) RemoveIS(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *IframeHTMLElement) ITEMID(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *IframeHTMLElement) IfITEMID(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *IframeHTMLElement) RemoveITEMID(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *IframeHTMLElement) ITEMPROP(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *IframeHTMLElement) IfITEMPROP(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *IframeHTMLElement) RemoveITEMPROP(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *IframeHTMLElement) ITEMREF(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *IframeHTMLElement) IfITEMREF(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *IframeHTMLElement) RemoveITEMREF(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *IframeHTMLElement) ITEMSCOPE() *IframeHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *IframeHTMLElement) IfITEMSCOPE(cond bool) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *IframeHTMLElement) RemoveITEMSCOPE() *IframeHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *IframeHTMLElement) SetITEMSCOPE(b bool) *IframeHTMLElement {
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
func (e *IframeHTMLElement) ITEMTYPE(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *IframeHTMLElement) IfITEMTYPE(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *IframeHTMLElement) RemoveITEMTYPE(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *IframeHTMLElement) LANG(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *IframeHTMLElement) IfLANG(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *IframeHTMLElement) RemoveLANG(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// LOADING sets the "loading" attribute.
// Used when determining loading deferral
// Values values are constrained to:
//   - lazy
//   - lazy
//   - eager
//   - eager
func (e *IframeHTMLElement) LOADING(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLOADINGKey] = v
	return e
}

func (e *IframeHTMLElement) IfLOADING(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.LOADING(v)
}

func (e *IframeHTMLElement) RemoveLOADING(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeLOADINGKey)
	return e
}

// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//   - text
func (e *IframeHTMLElement) NAME(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNAMEKey] = v
	return e
}

func (e *IframeHTMLElement) IfNAME(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.NAME(v)
}

func (e *IframeHTMLElement) RemoveNAME(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeNAMEKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *IframeHTMLElement) NONCE(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *IframeHTMLElement) IfNONCE(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *IframeHTMLElement) RemoveNONCE(v string) *IframeHTMLElement {
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
func (e *IframeHTMLElement) POPOVER(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *IframeHTMLElement) IfPOPOVER(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *IframeHTMLElement) RemovePOPOVER(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//   - referrer_policy
func (e *IframeHTMLElement) REFERRERPOLICY(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeREFERRERPOLICYKey] = v
	return e
}

func (e *IframeHTMLElement) IfREFERRERPOLICY(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.REFERRERPOLICY(v)
}

func (e *IframeHTMLElement) RemoveREFERRERPOLICY(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeREFERRERPOLICYKey)
	return e
}

// SANDBOX sets the "sandbox" attribute.
// Security rules for nested content
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - ascii_case_insensitive
//   - allow_downloads
//   - allow_downloads
//   - allow_forms
//   - allow_forms
//   - allow_modals
//   - allow_modals
//   - allow_orientation_lock
//   - allow_orientation_lock
//   - allow_pointer_lock
//   - allow_pointer_lock
//   - allow_popups
//   - allow_popups
//   - allow_popups_to_escape_sandbox
//   - allow_popups_to_escape_sandbox
//   - allow_presentation
//   - allow_presentation
//   - allow_same_origin
//   - allow_same_origin
//   - allow_scripts
//   - allow_scripts
//   - allow_top_navigation
//   - allow_top_navigation
//   - allow_top_navigation_by_user_activation
//   - allow_top_navigation_by_user_activation
//   - allow_top_navigation_to_custom_protocols
//   - allow_top_navigation_to_custom_protocols
func (e *IframeHTMLElement) SANDBOX(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSANDBOXKey] = v
	return e
}

func (e *IframeHTMLElement) IfSANDBOX(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.SANDBOX(v)
}

func (e *IframeHTMLElement) RemoveSANDBOX(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeSANDBOXKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *IframeHTMLElement) SLOT(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *IframeHTMLElement) IfSLOT(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *IframeHTMLElement) RemoveSLOT(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *IframeHTMLElement) SPELLCHECK(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *IframeHTMLElement) IfSPELLCHECK(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *IframeHTMLElement) RemoveSPELLCHECK(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *IframeHTMLElement) SRC(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCKey] = v
	return e
}

func (e *IframeHTMLElement) IfSRC(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.SRC(v)
}

func (e *IframeHTMLElement) RemoveSRC(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeSRCKey)
	return e
}

// SRCDOC sets the "srcdoc" attribute.
// A document to render in the iframe
// Values values are constrained to:
//   - an_iframe_srcdoc_document
//   - iframe
//   - srcdoc
func (e *IframeHTMLElement) SRCDOC(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCDOCKey] = v
	return e
}

func (e *IframeHTMLElement) IfSRCDOC(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.SRCDOC(v)
}

func (e *IframeHTMLElement) RemoveSRCDOC(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeSRCDOCKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *IframeHTMLElement) STYLE(k, v string) *IframeHTMLElement {
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

func (e *IframeHTMLElement) IfSTYLE(cond bool, k string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *IframeHTMLElement) RemoveSTYLE(k string) *IframeHTMLElement {
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
func (e *IframeHTMLElement) TABINDEX(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *IframeHTMLElement) IfTABINDEX(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *IframeHTMLElement) RemoveTABINDEX(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *IframeHTMLElement) TITLE(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *IframeHTMLElement) IfTITLE(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *IframeHTMLElement) RemoveTITLE(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *IframeHTMLElement) TRANSLATE(v string) *IframeHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *IframeHTMLElement) IfTRANSLATE(cond bool, v string) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *IframeHTMLElement) RemoveTRANSLATE(v string) *IframeHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *IframeHTMLElement) WIDTH(v int) *IframeHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeWIDTHKey] = v
	return e
}

func (e *IframeHTMLElement) IfWIDTH(cond bool, v int) *IframeHTMLElement {
	if !cond {
		return e
	}
	return e.WIDTH(v)
}

func (e *IframeHTMLElement) RemoveWIDTH(v int) *IframeHTMLElement {
	delete(e.IntAttributes, attributeWIDTHKey)
	return e
}
