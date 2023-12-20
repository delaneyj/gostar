package html

import (
	"fmt"
)

type ImgHTMLElement struct {
	*Element
}

func IMG(children ...ElementBuilder) *ImgHTMLElement {
	return &ImgHTMLElement{
		Element: &Element{
			Tag:           elementTagIMG,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *ImgHTMLElement) Children(children ...ElementBuilder) *ImgHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *ImgHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ImgHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *ImgHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ImgHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *ImgHTMLElement) Text(text string) *ImgHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *ImgHTMLElement) TextF(format string, args ...any) *ImgHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *ImgHTMLElement) Escaped(text string) *ImgHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *ImgHTMLElement) EscapedF(format string, args ...any) *ImgHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *ImgHTMLElement) CustomData(key, value string) *ImgHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ImgHTMLElement) CustomDataRemove(key string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) ACCESSKEY(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *ImgHTMLElement) IfACCESSKEY(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *ImgHTMLElement) RemoveACCESSKEY(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// ALT sets the "alt" attribute.
// Replacement text for use when images are not available
// Values values are constrained to:
//   - text
func (e *ImgHTMLElement) ALT(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeALTKey] = v
	return e
}

func (e *ImgHTMLElement) IfALT(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ALT(v)
}

func (e *ImgHTMLElement) RemoveALT(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeALTKey)
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
func (e *ImgHTMLElement) AUTOCAPITALIZE(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *ImgHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *ImgHTMLElement) RemoveAUTOCAPITALIZE(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *ImgHTMLElement) AUTOFOCUS() *ImgHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *ImgHTMLElement) IfAUTOFOCUS(cond bool) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *ImgHTMLElement) RemoveAUTOFOCUS() *ImgHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *ImgHTMLElement) SetAUTOFOCUS(b bool) *ImgHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *ImgHTMLElement) CLASS(v string) *ImgHTMLElement {
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

func (e *ImgHTMLElement) IfCLASS(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *ImgHTMLElement) SetCLASS(v string) *ImgHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *ImgHTMLElement) RemoveCLASS(v string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) CONTENTEDITABLE(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *ImgHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *ImgHTMLElement) RemoveCONTENTEDITABLE(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// CROSSORIGIN sets the "crossorigin" attribute.
// How the element handles crossorigin requests
// Values values are constrained to:
//   - anonymous
//   - use_credentials
func (e *ImgHTMLElement) CROSSORIGIN(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCROSSORIGINKey] = v
	return e
}

func (e *ImgHTMLElement) IfCROSSORIGIN(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.CROSSORIGIN(v)
}

func (e *ImgHTMLElement) RemoveCROSSORIGIN(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeCROSSORIGINKey)
	return e
}

// DECODING sets the "decoding" attribute.
// Decoding hint to use when processing this image for presentation
// Values values are constrained to:
//   - async
//   - auto
//   - sync
func (e *ImgHTMLElement) DECODING(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDECODINGKey] = v
	return e
}

func (e *ImgHTMLElement) IfDECODING(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.DECODING(v)
}

func (e *ImgHTMLElement) RemoveDECODING(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeDECODINGKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *ImgHTMLElement) DIR(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *ImgHTMLElement) IfDIR(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *ImgHTMLElement) RemoveDIR(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - false
//   - true
func (e *ImgHTMLElement) DRAGGABLE(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *ImgHTMLElement) IfDRAGGABLE(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *ImgHTMLElement) RemoveDRAGGABLE(v string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) ENTERKEYHINT(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *ImgHTMLElement) IfENTERKEYHINT(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *ImgHTMLElement) RemoveENTERKEYHINT(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FETCHPRIORITY sets the "fetchpriority" attribute.
// Sets the priority for fetches initiated by the element
// Values values are constrained to:
//   - auto
//   - high
//   - low
func (e *ImgHTMLElement) FETCHPRIORITY(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFETCHPRIORITYKey] = v
	return e
}

func (e *ImgHTMLElement) IfFETCHPRIORITY(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.FETCHPRIORITY(v)
}

func (e *ImgHTMLElement) RemoveFETCHPRIORITY(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeFETCHPRIORITYKey)
	return e
}

// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *ImgHTMLElement) HEIGHT(v int) *ImgHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeHEIGHTKey] = v
	return e
}

func (e *ImgHTMLElement) IfHEIGHT(cond bool, v int) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.HEIGHT(v)
}

func (e *ImgHTMLElement) RemoveHEIGHT(v int) *ImgHTMLElement {
	delete(e.IntAttributes, attributeHEIGHTKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *ImgHTMLElement) HIDDEN(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *ImgHTMLElement) IfHIDDEN(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *ImgHTMLElement) RemoveHIDDEN(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *ImgHTMLElement) ID(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *ImgHTMLElement) IfID(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *ImgHTMLElement) RemoveID(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *ImgHTMLElement) INERT() *ImgHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *ImgHTMLElement) IfINERT(cond bool) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *ImgHTMLElement) RemoveINERT() *ImgHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *ImgHTMLElement) SetINERT(b bool) *ImgHTMLElement {
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
func (e *ImgHTMLElement) INPUTMODE(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *ImgHTMLElement) IfINPUTMODE(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *ImgHTMLElement) RemoveINPUTMODE(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *ImgHTMLElement) IS(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *ImgHTMLElement) IfIS(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *ImgHTMLElement) RemoveIS(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ISMAP sets the "ismap" attribute.
// Whether the image is a server-side image map
// Values values are constrained to:
//   - boolean_attribute
func (e *ImgHTMLElement) ISMAP() *ImgHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeISMAPKey] = struct{}{}
	return e
}

func (e *ImgHTMLElement) IfISMAP(cond bool) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ISMAP()
}

func (e *ImgHTMLElement) RemoveISMAP() *ImgHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeISMAPKey)
	return e
}

func (e *ImgHTMLElement) SetISMAP(b bool) *ImgHTMLElement {
	if b {
		return e.ISMAP()
	}
	return e.RemoveISMAP()
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *ImgHTMLElement) ITEMID(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *ImgHTMLElement) IfITEMID(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *ImgHTMLElement) RemoveITEMID(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *ImgHTMLElement) ITEMPROP(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *ImgHTMLElement) IfITEMPROP(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *ImgHTMLElement) RemoveITEMPROP(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ImgHTMLElement) ITEMREF(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *ImgHTMLElement) IfITEMREF(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *ImgHTMLElement) RemoveITEMREF(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *ImgHTMLElement) ITEMSCOPE() *ImgHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *ImgHTMLElement) IfITEMSCOPE(cond bool) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *ImgHTMLElement) RemoveITEMSCOPE() *ImgHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *ImgHTMLElement) SetITEMSCOPE(b bool) *ImgHTMLElement {
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
func (e *ImgHTMLElement) ITEMTYPE(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *ImgHTMLElement) IfITEMTYPE(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *ImgHTMLElement) RemoveITEMTYPE(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ImgHTMLElement) LANG(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *ImgHTMLElement) IfLANG(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *ImgHTMLElement) RemoveLANG(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// LOADING sets the "loading" attribute.
// Used when determining loading deferral
// Values values are constrained to:
//   - eager
//   - lazy
func (e *ImgHTMLElement) LOADING(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLOADINGKey] = v
	return e
}

func (e *ImgHTMLElement) IfLOADING(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.LOADING(v)
}

func (e *ImgHTMLElement) RemoveLOADING(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeLOADINGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *ImgHTMLElement) NONCE(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *ImgHTMLElement) IfNONCE(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *ImgHTMLElement) RemoveNONCE(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *ImgHTMLElement) POPOVER(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *ImgHTMLElement) IfPOPOVER(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *ImgHTMLElement) RemovePOPOVER(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//   - referrer_policy
func (e *ImgHTMLElement) REFERRERPOLICY(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeREFERRERPOLICYKey] = v
	return e
}

func (e *ImgHTMLElement) IfREFERRERPOLICY(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.REFERRERPOLICY(v)
}

func (e *ImgHTMLElement) RemoveREFERRERPOLICY(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeREFERRERPOLICYKey)
	return e
}

// SIZES sets the "sizes" attribute.
// Image sizes for different page layouts
// Values values are constrained to:
//   - valid_source_size_list
func (e *ImgHTMLElement) SIZES(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSIZESKey] = v
	return e
}

func (e *ImgHTMLElement) IfSIZES(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.SIZES(v)
}

func (e *ImgHTMLElement) RemoveSIZES(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeSIZESKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *ImgHTMLElement) SLOT(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *ImgHTMLElement) IfSLOT(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *ImgHTMLElement) RemoveSLOT(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *ImgHTMLElement) SPELLCHECK(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *ImgHTMLElement) IfSPELLCHECK(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *ImgHTMLElement) RemoveSPELLCHECK(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//   - valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ImgHTMLElement) SRC(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCKey] = v
	return e
}

func (e *ImgHTMLElement) IfSRC(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.SRC(v)
}

func (e *ImgHTMLElement) RemoveSRC(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeSRCKey)
	return e
}

// SRCSET sets the "srcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc.
// Values values are constrained to:
//   - image_candidate_strings
func (e *ImgHTMLElement) SRCSET(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSRCSETKey] = v
	return e
}

func (e *ImgHTMLElement) IfSRCSET(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.SRCSET(v)
}

func (e *ImgHTMLElement) RemoveSRCSET(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeSRCSETKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ImgHTMLElement) STYLE(k, v string) *ImgHTMLElement {
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

func (e *ImgHTMLElement) IfSTYLE(cond bool, k string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *ImgHTMLElement) RemoveSTYLE(k string) *ImgHTMLElement {
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
func (e *ImgHTMLElement) TABINDEX(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *ImgHTMLElement) IfTABINDEX(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *ImgHTMLElement) RemoveTABINDEX(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *ImgHTMLElement) TITLE(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *ImgHTMLElement) IfTITLE(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *ImgHTMLElement) RemoveTITLE(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *ImgHTMLElement) TRANSLATE(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *ImgHTMLElement) IfTRANSLATE(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *ImgHTMLElement) RemoveTRANSLATE(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// USEMAP sets the "usemap" attribute.
// Name of image map to use
// Values values are constrained to:
//   - valid_hash_name_reference
func (e *ImgHTMLElement) USEMAP(v string) *ImgHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeUSEMAPKey] = v
	return e
}

func (e *ImgHTMLElement) IfUSEMAP(cond bool, v string) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.USEMAP(v)
}

func (e *ImgHTMLElement) RemoveUSEMAP(v string) *ImgHTMLElement {
	delete(e.StringAttributes, attributeUSEMAPKey)
	return e
}

// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//   - valid_non_negative_integer
func (e *ImgHTMLElement) WIDTH(v int) *ImgHTMLElement {
	if e.IntAttributes == nil {
		e.IntAttributes = map[string]int{}
	}
	e.IntAttributes[attributeWIDTHKey] = v
	return e
}

func (e *ImgHTMLElement) IfWIDTH(cond bool, v int) *ImgHTMLElement {
	if !cond {
		return e
	}
	return e.WIDTH(v)
}

func (e *ImgHTMLElement) RemoveWIDTH(v int) *ImgHTMLElement {
	delete(e.IntAttributes, attributeWIDTHKey)
	return e
}
