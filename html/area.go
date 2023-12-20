package html

import (
	"fmt"
)

type AreaHTMLElement struct {
	*Element
}

func AREA(children ...ElementBuilder) *AreaHTMLElement {
	return &AreaHTMLElement{
		Element: &Element{
			Tag:           elementTagAREA,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *AreaHTMLElement) Children(children ...ElementBuilder) *AreaHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *AreaHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *AreaHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *AreaHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *AreaHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *AreaHTMLElement) Text(text string) *AreaHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *AreaHTMLElement) TextF(format string, args ...any) *AreaHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *AreaHTMLElement) Escaped(text string) *AreaHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *AreaHTMLElement) EscapedF(format string, args ...any) *AreaHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *AreaHTMLElement) CustomData(key, value string) *AreaHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *AreaHTMLElement) CustomDataRemove(key string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) ACCESSKEY(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *AreaHTMLElement) IfACCESSKEY(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *AreaHTMLElement) RemoveACCESSKEY(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// ALT sets the "alt" attribute.
// Replacement text for use when images are not available
// Values values are constrained to:
//   - text
func (e *AreaHTMLElement) ALT(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeALTKey] = v
	return e
}

func (e *AreaHTMLElement) IfALT(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.ALT(v)
}

func (e *AreaHTMLElement) RemoveALT(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeALTKey)
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
func (e *AreaHTMLElement) AUTOCAPITALIZE(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *AreaHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *AreaHTMLElement) RemoveAUTOCAPITALIZE(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *AreaHTMLElement) AUTOFOCUS() *AreaHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *AreaHTMLElement) IfAUTOFOCUS(cond bool) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *AreaHTMLElement) RemoveAUTOFOCUS() *AreaHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *AreaHTMLElement) SetAUTOFOCUS(b bool) *AreaHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *AreaHTMLElement) CLASS(v string) *AreaHTMLElement {
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

func (e *AreaHTMLElement) IfCLASS(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *AreaHTMLElement) SetCLASS(v string) *AreaHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *AreaHTMLElement) RemoveCLASS(v string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) CONTENTEDITABLE(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *AreaHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *AreaHTMLElement) RemoveCONTENTEDITABLE(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// COORDS sets the "coords" attribute.
// Coordinates for the shape to be created in an image map
// Values values are constrained to:
//   - valid_list_of_floating_point_numbers
func (e *AreaHTMLElement) COORDS(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCOORDSKey] = v
	return e
}

func (e *AreaHTMLElement) IfCOORDS(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.COORDS(v)
}

func (e *AreaHTMLElement) RemoveCOORDS(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeCOORDSKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - ltr
//   - rtl
//   - rtl
func (e *AreaHTMLElement) DIR(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *AreaHTMLElement) IfDIR(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *AreaHTMLElement) RemoveDIR(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DOWNLOAD sets the "download" attribute.
// Whether to download the resource instead of navigating to it, and its filename if so
// Values values are constrained to:
func (e *AreaHTMLElement) DOWNLOAD(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDOWNLOADKey] = v
	return e
}

func (e *AreaHTMLElement) IfDOWNLOAD(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.DOWNLOAD(v)
}

func (e *AreaHTMLElement) RemoveDOWNLOAD(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeDOWNLOADKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *AreaHTMLElement) DRAGGABLE(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *AreaHTMLElement) IfDRAGGABLE(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *AreaHTMLElement) RemoveDRAGGABLE(v string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) ENTERKEYHINT(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *AreaHTMLElement) IfENTERKEYHINT(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *AreaHTMLElement) RemoveENTERKEYHINT(v string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) HIDDEN(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *AreaHTMLElement) IfHIDDEN(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *AreaHTMLElement) RemoveHIDDEN(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// HREF sets the "href" attribute.
// Document base URL
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *AreaHTMLElement) HREF(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHREFKey] = v
	return e
}

func (e *AreaHTMLElement) IfHREF(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.HREF(v)
}

func (e *AreaHTMLElement) RemoveHREF(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeHREFKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *AreaHTMLElement) ID(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *AreaHTMLElement) IfID(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *AreaHTMLElement) RemoveID(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *AreaHTMLElement) INERT() *AreaHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *AreaHTMLElement) IfINERT(cond bool) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *AreaHTMLElement) RemoveINERT() *AreaHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *AreaHTMLElement) SetINERT(b bool) *AreaHTMLElement {
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
func (e *AreaHTMLElement) INPUTMODE(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *AreaHTMLElement) IfINPUTMODE(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *AreaHTMLElement) RemoveINPUTMODE(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *AreaHTMLElement) IS(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *AreaHTMLElement) IfIS(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *AreaHTMLElement) RemoveIS(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *AreaHTMLElement) ITEMID(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *AreaHTMLElement) IfITEMID(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *AreaHTMLElement) RemoveITEMID(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *AreaHTMLElement) ITEMPROP(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *AreaHTMLElement) IfITEMPROP(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *AreaHTMLElement) RemoveITEMPROP(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *AreaHTMLElement) ITEMREF(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *AreaHTMLElement) IfITEMREF(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *AreaHTMLElement) RemoveITEMREF(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *AreaHTMLElement) ITEMSCOPE() *AreaHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *AreaHTMLElement) IfITEMSCOPE(cond bool) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *AreaHTMLElement) RemoveITEMSCOPE() *AreaHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *AreaHTMLElement) SetITEMSCOPE(b bool) *AreaHTMLElement {
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
func (e *AreaHTMLElement) ITEMTYPE(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *AreaHTMLElement) IfITEMTYPE(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *AreaHTMLElement) RemoveITEMTYPE(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *AreaHTMLElement) LANG(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *AreaHTMLElement) IfLANG(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *AreaHTMLElement) RemoveLANG(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *AreaHTMLElement) NONCE(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *AreaHTMLElement) IfNONCE(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *AreaHTMLElement) RemoveNONCE(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// PING sets the "ping" attribute.
// URLs to ping
// Values values are constrained to:
//   - set_of_space_separated_tokens
//   - valid_non_empty_ur_ls
func (e *AreaHTMLElement) PING(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePINGKey] = v
	return e
}

func (e *AreaHTMLElement) IfPING(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.PING(v)
}

func (e *AreaHTMLElement) RemovePING(v string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) POPOVER(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *AreaHTMLElement) IfPOPOVER(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *AreaHTMLElement) RemovePOPOVER(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//   - referrer_policy
func (e *AreaHTMLElement) REFERRERPOLICY(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeREFERRERPOLICYKey] = v
	return e
}

func (e *AreaHTMLElement) IfREFERRERPOLICY(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.REFERRERPOLICY(v)
}

func (e *AreaHTMLElement) RemoveREFERRERPOLICY(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeREFERRERPOLICYKey)
	return e
}

// REL sets the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *AreaHTMLElement) REL(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeRELKey] = v
	return e
}

func (e *AreaHTMLElement) IfREL(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.REL(v)
}

func (e *AreaHTMLElement) RemoveREL(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeRELKey)
	return e
}

// SHAPE sets the "shape" attribute.
// The kind of shape to be created in an image map
// Values values are constrained to:
//   - circle
//   - circle
//   - default
//   - default
//   - poly
//   - poly
//   - rect
//   - rect
func (e *AreaHTMLElement) SHAPE(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSHAPEKey] = v
	return e
}

func (e *AreaHTMLElement) IfSHAPE(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.SHAPE(v)
}

func (e *AreaHTMLElement) RemoveSHAPE(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeSHAPEKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *AreaHTMLElement) SLOT(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *AreaHTMLElement) IfSLOT(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *AreaHTMLElement) RemoveSLOT(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *AreaHTMLElement) SPELLCHECK(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *AreaHTMLElement) IfSPELLCHECK(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *AreaHTMLElement) RemoveSPELLCHECK(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *AreaHTMLElement) STYLE(k, v string) *AreaHTMLElement {
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

func (e *AreaHTMLElement) IfSTYLE(cond bool, k string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *AreaHTMLElement) RemoveSTYLE(k string) *AreaHTMLElement {
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
func (e *AreaHTMLElement) TABINDEX(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *AreaHTMLElement) IfTABINDEX(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *AreaHTMLElement) RemoveTABINDEX(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TARGET sets the "target" attribute.
// Navigable for form submission
// Values values are constrained to:
//   - valid_navigable_target_name_or_keyword
func (e *AreaHTMLElement) TARGET(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTARGETKey] = v
	return e
}

func (e *AreaHTMLElement) IfTARGET(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.TARGET(v)
}

func (e *AreaHTMLElement) RemoveTARGET(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeTARGETKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *AreaHTMLElement) TITLE(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *AreaHTMLElement) IfTITLE(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *AreaHTMLElement) RemoveTITLE(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *AreaHTMLElement) TRANSLATE(v string) *AreaHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *AreaHTMLElement) IfTRANSLATE(cond bool, v string) *AreaHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *AreaHTMLElement) RemoveTRANSLATE(v string) *AreaHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
