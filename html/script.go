package html

import (
	"fmt"
)

type ScriptHTMLElement struct {
	*Element
}

func SCRIPT(children ...ElementBuilder) *ScriptHTMLElement {
	return &ScriptHTMLElement{
		Element: &Element{
			Tag:           elementTagSCRIPT,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *ScriptHTMLElement) Children(children ...ElementBuilder) *ScriptHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *ScriptHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ScriptHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *ScriptHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ScriptHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *ScriptHTMLElement) Text(text string) *ScriptHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *ScriptHTMLElement) TextF(format string, args ...any) *ScriptHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *ScriptHTMLElement) Escaped(text string) *ScriptHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *ScriptHTMLElement) EscapedF(format string, args ...any) *ScriptHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *ScriptHTMLElement) CustomData(key, value string) *ScriptHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ScriptHTMLElement) CustomDataRemove(key string) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) ACCESSKEY(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *ScriptHTMLElement) IfACCESSKEY(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *ScriptHTMLElement) RemoveACCESSKEY(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// ASYNC sets the "async" attribute.
// Execute script when available, without blocking while fetching
// Values values are constrained to:
//   - boolean_attribute
func (e *ScriptHTMLElement) ASYNC() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeASYNCKey] = struct{}{}
	return e
}

func (e *ScriptHTMLElement) IfASYNC(cond bool) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.ASYNC()
}

func (e *ScriptHTMLElement) RemoveASYNC() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeASYNCKey)
	return e
}

func (e *ScriptHTMLElement) SetASYNC(b bool) *ScriptHTMLElement {
	if b {
		return e.ASYNC()
	}
	return e.RemoveASYNC()
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
func (e *ScriptHTMLElement) AUTOCAPITALIZE(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *ScriptHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *ScriptHTMLElement) RemoveAUTOCAPITALIZE(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *ScriptHTMLElement) AUTOFOCUS() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *ScriptHTMLElement) IfAUTOFOCUS(cond bool) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *ScriptHTMLElement) RemoveAUTOFOCUS() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *ScriptHTMLElement) SetAUTOFOCUS(b bool) *ScriptHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// BLOCKING sets the "blocking" attribute.
// Whether the element is potentially render-blocking
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ScriptHTMLElement) BLOCKING(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeBLOCKINGKey] = v
	return e
}

func (e *ScriptHTMLElement) IfBLOCKING(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.BLOCKING(v)
}

func (e *ScriptHTMLElement) RemoveBLOCKING(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeBLOCKINGKey)
	return e
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *ScriptHTMLElement) CLASS(v string) *ScriptHTMLElement {
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

func (e *ScriptHTMLElement) IfCLASS(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *ScriptHTMLElement) SetCLASS(v string) *ScriptHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *ScriptHTMLElement) RemoveCLASS(v string) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) CONTENTEDITABLE(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *ScriptHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *ScriptHTMLElement) RemoveCONTENTEDITABLE(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// CROSSORIGIN sets the "crossorigin" attribute.
// How the element handles crossorigin requests
// Values values are constrained to:
//   - anonymous
//   - use_credentials
func (e *ScriptHTMLElement) CROSSORIGIN(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCROSSORIGINKey] = v
	return e
}

func (e *ScriptHTMLElement) IfCROSSORIGIN(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.CROSSORIGIN(v)
}

func (e *ScriptHTMLElement) RemoveCROSSORIGIN(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeCROSSORIGINKey)
	return e
}

// DEFER sets the "defer" attribute.
// Defer script execution
// Values values are constrained to:
//   - boolean_attribute
func (e *ScriptHTMLElement) DEFER() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeDEFERKey] = struct{}{}
	return e
}

func (e *ScriptHTMLElement) IfDEFER(cond bool) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.DEFER()
}

func (e *ScriptHTMLElement) RemoveDEFER() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeDEFERKey)
	return e
}

func (e *ScriptHTMLElement) SetDEFER(b bool) *ScriptHTMLElement {
	if b {
		return e.DEFER()
	}
	return e.RemoveDEFER()
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *ScriptHTMLElement) DIR(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *ScriptHTMLElement) IfDIR(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *ScriptHTMLElement) RemoveDIR(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *ScriptHTMLElement) DRAGGABLE(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *ScriptHTMLElement) IfDRAGGABLE(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *ScriptHTMLElement) RemoveDRAGGABLE(v string) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) ENTERKEYHINT(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *ScriptHTMLElement) IfENTERKEYHINT(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *ScriptHTMLElement) RemoveENTERKEYHINT(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FETCHPRIORITY sets the "fetchpriority" attribute.
// Sets the priority for fetches initiated by the element
// Values values are constrained to:
//   - auto
//   - high
//   - low
func (e *ScriptHTMLElement) FETCHPRIORITY(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFETCHPRIORITYKey] = v
	return e
}

func (e *ScriptHTMLElement) IfFETCHPRIORITY(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.FETCHPRIORITY(v)
}

func (e *ScriptHTMLElement) RemoveFETCHPRIORITY(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeFETCHPRIORITYKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *ScriptHTMLElement) HIDDEN(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *ScriptHTMLElement) IfHIDDEN(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *ScriptHTMLElement) RemoveHIDDEN(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *ScriptHTMLElement) ID(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *ScriptHTMLElement) IfID(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *ScriptHTMLElement) RemoveID(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *ScriptHTMLElement) INERT() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *ScriptHTMLElement) IfINERT(cond bool) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *ScriptHTMLElement) RemoveINERT() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *ScriptHTMLElement) SetINERT(b bool) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) INPUTMODE(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *ScriptHTMLElement) IfINPUTMODE(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *ScriptHTMLElement) RemoveINPUTMODE(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// INTEGRITY sets the "integrity" attribute.
// Integrity metadata used in Subresource Integrity checks [SRI]
// Values values are constrained to:
//   - text
func (e *ScriptHTMLElement) INTEGRITY(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINTEGRITYKey] = v
	return e
}

func (e *ScriptHTMLElement) IfINTEGRITY(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.INTEGRITY(v)
}

func (e *ScriptHTMLElement) RemoveINTEGRITY(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeINTEGRITYKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *ScriptHTMLElement) IS(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *ScriptHTMLElement) IfIS(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *ScriptHTMLElement) RemoveIS(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *ScriptHTMLElement) ITEMID(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *ScriptHTMLElement) IfITEMID(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *ScriptHTMLElement) RemoveITEMID(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *ScriptHTMLElement) ITEMPROP(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *ScriptHTMLElement) IfITEMPROP(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *ScriptHTMLElement) RemoveITEMPROP(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ScriptHTMLElement) ITEMREF(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *ScriptHTMLElement) IfITEMREF(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *ScriptHTMLElement) RemoveITEMREF(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *ScriptHTMLElement) ITEMSCOPE() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *ScriptHTMLElement) IfITEMSCOPE(cond bool) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *ScriptHTMLElement) RemoveITEMSCOPE() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *ScriptHTMLElement) SetITEMSCOPE(b bool) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) ITEMTYPE(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *ScriptHTMLElement) IfITEMTYPE(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *ScriptHTMLElement) RemoveITEMTYPE(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ScriptHTMLElement) LANG(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *ScriptHTMLElement) IfLANG(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *ScriptHTMLElement) RemoveLANG(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NOMODULE sets the "nomodule" attribute.
// Prevents execution in user agents that support module scripts
// Values values are constrained to:
//   - boolean_attribute
func (e *ScriptHTMLElement) NOMODULE() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeNOMODULEKey] = struct{}{}
	return e
}

func (e *ScriptHTMLElement) IfNOMODULE(cond bool) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.NOMODULE()
}

func (e *ScriptHTMLElement) RemoveNOMODULE() *ScriptHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeNOMODULEKey)
	return e
}

func (e *ScriptHTMLElement) SetNOMODULE(b bool) *ScriptHTMLElement {
	if b {
		return e.NOMODULE()
	}
	return e.RemoveNOMODULE()
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *ScriptHTMLElement) NONCE(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *ScriptHTMLElement) IfNONCE(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *ScriptHTMLElement) RemoveNONCE(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *ScriptHTMLElement) POPOVER(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *ScriptHTMLElement) IfPOPOVER(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *ScriptHTMLElement) RemovePOPOVER(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//   - referrer_policy
func (e *ScriptHTMLElement) REFERRERPOLICY(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeREFERRERPOLICYKey] = v
	return e
}

func (e *ScriptHTMLElement) IfREFERRERPOLICY(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.REFERRERPOLICY(v)
}

func (e *ScriptHTMLElement) RemoveREFERRERPOLICY(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeREFERRERPOLICYKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *ScriptHTMLElement) SLOT(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *ScriptHTMLElement) IfSLOT(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *ScriptHTMLElement) RemoveSLOT(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *ScriptHTMLElement) SPELLCHECK(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *ScriptHTMLElement) IfSPELLCHECK(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *ScriptHTMLElement) RemoveSPELLCHECK(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ScriptHTMLElement) SRC(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCKey] = v
	return e
}

func (e *ScriptHTMLElement) IfSRC(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.SRC(v)
}

func (e *ScriptHTMLElement) RemoveSRC(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeSRCKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ScriptHTMLElement) STYLE(k, v string) *ScriptHTMLElement {
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

func (e *ScriptHTMLElement) IfSTYLE(cond bool, k string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *ScriptHTMLElement) RemoveSTYLE(k string) *ScriptHTMLElement {
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
func (e *ScriptHTMLElement) TABINDEX(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *ScriptHTMLElement) IfTABINDEX(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *ScriptHTMLElement) RemoveTABINDEX(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *ScriptHTMLElement) TITLE(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *ScriptHTMLElement) IfTITLE(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *ScriptHTMLElement) RemoveTITLE(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *ScriptHTMLElement) TRANSLATE(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *ScriptHTMLElement) IfTRANSLATE(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *ScriptHTMLElement) RemoveTRANSLATE(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//   - java_script_mime_type_essence_match
//   - module
//   - valid_mime_type_string
func (e *ScriptHTMLElement) TYPE(v string) *ScriptHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTYPEKey] = v
	return e
}

func (e *ScriptHTMLElement) IfTYPE(cond bool, v string) *ScriptHTMLElement {
	if !cond {
		return e
	}
	return e.TYPE(v)
}

func (e *ScriptHTMLElement) RemoveTYPE(v string) *ScriptHTMLElement {
	delete(e.StringAttributes, attributeTYPEKey)
	return e
}
