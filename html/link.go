package html

import (
	"fmt"
)

type LinkHTMLElement struct {
	*Element
}

func LINK(children ...ElementBuilder) *LinkHTMLElement {
	return &LinkHTMLElement{
		Element: &Element{
			Tag:           elementTagLINK,
			IsSelfClosing: true,
			Descendants:   children,
		},
	}
}

func (e *LinkHTMLElement) Children(children ...ElementBuilder) *LinkHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *LinkHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *LinkHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *LinkHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *LinkHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *LinkHTMLElement) Text(text string) *LinkHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *LinkHTMLElement) TextF(format string, args ...any) *LinkHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *LinkHTMLElement) Escaped(text string) *LinkHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *LinkHTMLElement) EscapedF(format string, args ...any) *LinkHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *LinkHTMLElement) CustomData(key, value string) *LinkHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *LinkHTMLElement) CustomDataRemove(key string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) ACCESSKEY(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *LinkHTMLElement) IfACCESSKEY(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *LinkHTMLElement) RemoveACCESSKEY(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeACCESSKEYKey)
	return e
}

// AS sets the "as" attribute.
// Potential destination for a preload request (for rel="preload" and rel="modulepreload")
// Values values are constrained to:
//   - modulepreload
//   - potential_destination
//   - preload
//   - rel
//   - script_like_destination
func (e *LinkHTMLElement) AS(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeASKey] = v
	return e
}

func (e *LinkHTMLElement) IfAS(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.AS(v)
}

func (e *LinkHTMLElement) RemoveAS(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeASKey)
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
func (e *LinkHTMLElement) AUTOCAPITALIZE(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *LinkHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *LinkHTMLElement) RemoveAUTOCAPITALIZE(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *LinkHTMLElement) AUTOFOCUS() *LinkHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *LinkHTMLElement) IfAUTOFOCUS(cond bool) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *LinkHTMLElement) RemoveAUTOFOCUS() *LinkHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *LinkHTMLElement) SetAUTOFOCUS(b bool) *LinkHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// BLOCKING sets the "blocking" attribute.
// Whether the element is potentially render-blocking
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *LinkHTMLElement) BLOCKING(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeBLOCKINGKey] = v
	return e
}

func (e *LinkHTMLElement) IfBLOCKING(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.BLOCKING(v)
}

func (e *LinkHTMLElement) RemoveBLOCKING(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeBLOCKINGKey)
	return e
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *LinkHTMLElement) CLASS(v string) *LinkHTMLElement {
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

func (e *LinkHTMLElement) IfCLASS(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *LinkHTMLElement) SetCLASS(v string) *LinkHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *LinkHTMLElement) RemoveCLASS(v string) *LinkHTMLElement {
	kv, ok := e.DelimitedStringAttributes[attributeCLASSKey]
	if !ok {
		return e
	}
	kv.Remove(v)
	return e
}

// COLOR sets the "color" attribute.
// Color to use when customizing a site's icon (for rel="mask-icon")
// Values values are constrained to:
//   - <color>
func (e *LinkHTMLElement) COLOR(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCOLORKey] = v
	return e
}

func (e *LinkHTMLElement) IfCOLOR(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.COLOR(v)
}

func (e *LinkHTMLElement) RemoveCOLOR(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeCOLORKey)
	return e
}

// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//   - false
//   - plaintext_only
//   - true
func (e *LinkHTMLElement) CONTENTEDITABLE(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *LinkHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *LinkHTMLElement) RemoveCONTENTEDITABLE(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeCONTENTEDITABLEKey)
	return e
}

// CROSSORIGIN sets the "crossorigin" attribute.
// How the element handles crossorigin requests
// Values values are constrained to:
//   - anonymous
//   - use_credentials
func (e *LinkHTMLElement) CROSSORIGIN(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCROSSORIGINKey] = v
	return e
}

func (e *LinkHTMLElement) IfCROSSORIGIN(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.CROSSORIGIN(v)
}

func (e *LinkHTMLElement) RemoveCROSSORIGIN(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeCROSSORIGINKey)
	return e
}

// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//   - ltr
//   - rtl
func (e *LinkHTMLElement) DIR(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *LinkHTMLElement) IfDIR(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *LinkHTMLElement) RemoveDIR(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//   - boolean_attribute
func (e *LinkHTMLElement) DISABLED() *LinkHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeDISABLEDKey] = struct{}{}
	return e
}

func (e *LinkHTMLElement) IfDISABLED(cond bool) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.DISABLED()
}

func (e *LinkHTMLElement) RemoveDISABLED() *LinkHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeDISABLEDKey)
	return e
}

func (e *LinkHTMLElement) SetDISABLED(b bool) *LinkHTMLElement {
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
func (e *LinkHTMLElement) DRAGGABLE(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *LinkHTMLElement) IfDRAGGABLE(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *LinkHTMLElement) RemoveDRAGGABLE(v string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) ENTERKEYHINT(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *LinkHTMLElement) IfENTERKEYHINT(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *LinkHTMLElement) RemoveENTERKEYHINT(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeENTERKEYHINTKey)
	return e
}

// FETCHPRIORITY sets the "fetchpriority" attribute.
// Sets the priority for fetches initiated by the element
// Values values are constrained to:
//   - auto
//   - high
//   - low
func (e *LinkHTMLElement) FETCHPRIORITY(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeFETCHPRIORITYKey] = v
	return e
}

func (e *LinkHTMLElement) IfFETCHPRIORITY(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.FETCHPRIORITY(v)
}

func (e *LinkHTMLElement) RemoveFETCHPRIORITY(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeFETCHPRIORITYKey)
	return e
}

// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//   - hidden
//   - until_found
func (e *LinkHTMLElement) HIDDEN(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *LinkHTMLElement) IfHIDDEN(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *LinkHTMLElement) RemoveHIDDEN(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// HREF sets the "href" attribute.
// Document base URL
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *LinkHTMLElement) HREF(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHREFKey] = v
	return e
}

func (e *LinkHTMLElement) IfHREF(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.HREF(v)
}

func (e *LinkHTMLElement) RemoveHREF(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeHREFKey)
	return e
}

// HREFLANG sets the "hreflang" attribute.
// Language of the linked resource
// Values values are constrained to:
func (e *LinkHTMLElement) HREFLANG(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHREFLANGKey] = v
	return e
}

func (e *LinkHTMLElement) IfHREFLANG(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.HREFLANG(v)
}

func (e *LinkHTMLElement) RemoveHREFLANG(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeHREFLANGKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *LinkHTMLElement) ID(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *LinkHTMLElement) IfID(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *LinkHTMLElement) RemoveID(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// IMAGESIZES sets the "imagesizes" attribute.
// Image sizes for different page layouts (for rel="preload")
// Values values are constrained to:
//   - valid_source_size_list
func (e *LinkHTMLElement) IMAGESIZES(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIMAGESIZESKey] = v
	return e
}

func (e *LinkHTMLElement) IfIMAGESIZES(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.IMAGESIZES(v)
}

func (e *LinkHTMLElement) RemoveIMAGESIZES(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeIMAGESIZESKey)
	return e
}

// IMAGESRCSET sets the "imagesrcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc. (for rel="preload")
// Values values are constrained to:
//   - image_candidate_strings
func (e *LinkHTMLElement) IMAGESRCSET(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIMAGESRCSETKey] = v
	return e
}

func (e *LinkHTMLElement) IfIMAGESRCSET(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.IMAGESRCSET(v)
}

func (e *LinkHTMLElement) RemoveIMAGESRCSET(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeIMAGESRCSETKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *LinkHTMLElement) INERT() *LinkHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *LinkHTMLElement) IfINERT(cond bool) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *LinkHTMLElement) RemoveINERT() *LinkHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *LinkHTMLElement) SetINERT(b bool) *LinkHTMLElement {
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
func (e *LinkHTMLElement) INPUTMODE(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *LinkHTMLElement) IfINPUTMODE(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *LinkHTMLElement) RemoveINPUTMODE(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// INTEGRITY sets the "integrity" attribute.
// Integrity metadata used in Subresource Integrity checks [SRI]
// Values values are constrained to:
//   - text
func (e *LinkHTMLElement) INTEGRITY(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINTEGRITYKey] = v
	return e
}

func (e *LinkHTMLElement) IfINTEGRITY(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.INTEGRITY(v)
}

func (e *LinkHTMLElement) RemoveINTEGRITY(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeINTEGRITYKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - customized_built_in_element
//   - valid_custom_element_name
func (e *LinkHTMLElement) IS(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *LinkHTMLElement) IfIS(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *LinkHTMLElement) RemoveIS(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *LinkHTMLElement) ITEMID(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *LinkHTMLElement) IfITEMID(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *LinkHTMLElement) RemoveITEMID(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - defined_property_names
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
func (e *LinkHTMLElement) ITEMPROP(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *LinkHTMLElement) IfITEMPROP(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *LinkHTMLElement) RemoveITEMPROP(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *LinkHTMLElement) ITEMREF(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *LinkHTMLElement) IfITEMREF(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *LinkHTMLElement) RemoveITEMREF(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *LinkHTMLElement) ITEMSCOPE() *LinkHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *LinkHTMLElement) IfITEMSCOPE(cond bool) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *LinkHTMLElement) RemoveITEMSCOPE() *LinkHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *LinkHTMLElement) SetITEMSCOPE(b bool) *LinkHTMLElement {
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
func (e *LinkHTMLElement) ITEMTYPE(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *LinkHTMLElement) IfITEMTYPE(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *LinkHTMLElement) RemoveITEMTYPE(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *LinkHTMLElement) LANG(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *LinkHTMLElement) IfLANG(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *LinkHTMLElement) RemoveLANG(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// MEDIA sets the "media" attribute.
// Applicable media
// Values values are constrained to:
//   - valid_media_query_list
func (e *LinkHTMLElement) MEDIA(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeMEDIAKey] = v
	return e
}

func (e *LinkHTMLElement) IfMEDIA(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.MEDIA(v)
}

func (e *LinkHTMLElement) RemoveMEDIA(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeMEDIAKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *LinkHTMLElement) NONCE(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *LinkHTMLElement) IfNONCE(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *LinkHTMLElement) RemoveNONCE(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeNONCEKey)
	return e
}

// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//   - auto
//   - manual
func (e *LinkHTMLElement) POPOVER(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *LinkHTMLElement) IfPOPOVER(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *LinkHTMLElement) RemovePOPOVER(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// REFERRERPOLICY sets the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Values values are constrained to:
//   - referrer_policy
func (e *LinkHTMLElement) REFERRERPOLICY(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeREFERRERPOLICYKey] = v
	return e
}

func (e *LinkHTMLElement) IfREFERRERPOLICY(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.REFERRERPOLICY(v)
}

func (e *LinkHTMLElement) RemoveREFERRERPOLICY(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeREFERRERPOLICYKey)
	return e
}

// REL sets the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *LinkHTMLElement) REL(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeRELKey] = v
	return e
}

func (e *LinkHTMLElement) IfREL(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.REL(v)
}

func (e *LinkHTMLElement) RemoveREL(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeRELKey)
	return e
}

// SIZES sets the "sizes" attribute.
// Image sizes for different page layouts
// Values values are constrained to:
//   - valid_source_size_list
func (e *LinkHTMLElement) SIZES(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSIZESKey] = v
	return e
}

func (e *LinkHTMLElement) IfSIZES(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.SIZES(v)
}

func (e *LinkHTMLElement) RemoveSIZES(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeSIZESKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *LinkHTMLElement) SLOT(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *LinkHTMLElement) IfSLOT(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *LinkHTMLElement) RemoveSLOT(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - false
//   - true
func (e *LinkHTMLElement) SPELLCHECK(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *LinkHTMLElement) IfSPELLCHECK(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *LinkHTMLElement) RemoveSPELLCHECK(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *LinkHTMLElement) STYLE(k, v string) *LinkHTMLElement {
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

func (e *LinkHTMLElement) IfSTYLE(cond bool, k string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *LinkHTMLElement) RemoveSTYLE(k string) *LinkHTMLElement {
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
func (e *LinkHTMLElement) TABINDEX(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *LinkHTMLElement) IfTABINDEX(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *LinkHTMLElement) RemoveTABINDEX(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *LinkHTMLElement) TITLE(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *LinkHTMLElement) IfTITLE(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *LinkHTMLElement) RemoveTITLE(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - no
//   - yes
func (e *LinkHTMLElement) TRANSLATE(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *LinkHTMLElement) IfTRANSLATE(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *LinkHTMLElement) RemoveTRANSLATE(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}

// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//   - java_script_mime_type_essence_match
//   - module
//   - valid_mime_type_string
func (e *LinkHTMLElement) TYPE(v string) *LinkHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTYPEKey] = v
	return e
}

func (e *LinkHTMLElement) IfTYPE(cond bool, v string) *LinkHTMLElement {
	if !cond {
		return e
	}
	return e.TYPE(v)
}

func (e *LinkHTMLElement) RemoveTYPE(v string) *LinkHTMLElement {
	delete(e.StringAttributes, attributeTYPEKey)
	return e
}
