package html

import (
	"fmt"
)

type ArticleHTMLElement struct {
	*Element
}

func ARTICLE(children ...ElementBuilder) *ArticleHTMLElement {
	return &ArticleHTMLElement{
		Element: &Element{
			Tag:           elementTagARTICLE,
			IsSelfClosing: false,
			Descendants:   children,
		},
	}
}

func (e *ArticleHTMLElement) Children(children ...ElementBuilder) *ArticleHTMLElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *ArticleHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ArticleHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *ArticleHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ArticleHTMLElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *ArticleHTMLElement) Text(text string) *ArticleHTMLElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *ArticleHTMLElement) TextF(format string, args ...any) *ArticleHTMLElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *ArticleHTMLElement) Escaped(text string) *ArticleHTMLElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *ArticleHTMLElement) EscapedF(format string, args ...any) *ArticleHTMLElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *ArticleHTMLElement) CustomData(key, value string) *ArticleHTMLElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = map[string]string{}
	}
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ArticleHTMLElement) CustomDataRemove(key string) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) ACCESSKEY(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeACCESSKEYKey] = v
	return e
}

func (e *ArticleHTMLElement) IfACCESSKEY(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.ACCESSKEY(v)
}

func (e *ArticleHTMLElement) RemoveACCESSKEY(v string) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) AUTOCAPITALIZE(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeAUTOCAPITALIZEKey] = v
	return e
}

func (e *ArticleHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOCAPITALIZE(v)
}

func (e *ArticleHTMLElement) RemoveAUTOCAPITALIZE(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeAUTOCAPITALIZEKey)
	return e
}

// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//   - boolean_attribute
func (e *ArticleHTMLElement) AUTOFOCUS() *ArticleHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeAUTOFOCUSKey] = struct{}{}
	return e
}

func (e *ArticleHTMLElement) IfAUTOFOCUS(cond bool) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.AUTOFOCUS()
}

func (e *ArticleHTMLElement) RemoveAUTOFOCUS() *ArticleHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeAUTOFOCUSKey)
	return e
}

func (e *ArticleHTMLElement) SetAUTOFOCUS(b bool) *ArticleHTMLElement {
	if b {
		return e.AUTOFOCUS()
	}
	return e.RemoveAUTOFOCUS()
}

// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//   - set_of_space_separated_tokens
func (e *ArticleHTMLElement) CLASS(v string) *ArticleHTMLElement {
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

func (e *ArticleHTMLElement) IfCLASS(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.CLASS(v)
}

func (e *ArticleHTMLElement) SetCLASS(v string) *ArticleHTMLElement {
	kv := NewSpaceDelimitedString()
	e.DelimitedStringAttributes[attributeCLASSKey] = kv
	kv.Add(v)
	return e
}

func (e *ArticleHTMLElement) RemoveCLASS(v string) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) CONTENTEDITABLE(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeCONTENTEDITABLEKey] = v
	return e
}

func (e *ArticleHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.CONTENTEDITABLE(v)
}

func (e *ArticleHTMLElement) RemoveCONTENTEDITABLE(v string) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) DIR(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDIRKey] = v
	return e
}

func (e *ArticleHTMLElement) IfDIR(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.DIR(v)
}

func (e *ArticleHTMLElement) RemoveDIR(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeDIRKey)
	return e
}

// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//   - true
//   - false
func (e *ArticleHTMLElement) DRAGGABLE(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeDRAGGABLEKey] = v
	return e
}

func (e *ArticleHTMLElement) IfDRAGGABLE(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.DRAGGABLE(v)
}

func (e *ArticleHTMLElement) RemoveDRAGGABLE(v string) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) ENTERKEYHINT(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeENTERKEYHINTKey] = v
	return e
}

func (e *ArticleHTMLElement) IfENTERKEYHINT(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.ENTERKEYHINT(v)
}

func (e *ArticleHTMLElement) RemoveENTERKEYHINT(v string) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) HIDDEN(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeHIDDENKey] = v
	return e
}

func (e *ArticleHTMLElement) IfHIDDEN(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.HIDDEN(v)
}

func (e *ArticleHTMLElement) RemoveHIDDEN(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeHIDDENKey)
	return e
}

// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//   - text
func (e *ArticleHTMLElement) ID(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeIDKey] = v
	return e
}

func (e *ArticleHTMLElement) IfID(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.ID(v)
}

func (e *ArticleHTMLElement) RemoveID(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeIDKey)
	return e
}

// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//   - boolean_attribute
func (e *ArticleHTMLElement) INERT() *ArticleHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeINERTKey] = struct{}{}
	return e
}

func (e *ArticleHTMLElement) IfINERT(cond bool) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.INERT()
}

func (e *ArticleHTMLElement) RemoveINERT() *ArticleHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeINERTKey)
	return e
}

func (e *ArticleHTMLElement) SetINERT(b bool) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) INPUTMODE(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeINPUTMODEKey] = v
	return e
}

func (e *ArticleHTMLElement) IfINPUTMODE(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.INPUTMODE(v)
}

func (e *ArticleHTMLElement) RemoveINPUTMODE(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeINPUTMODEKey)
	return e
}

// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//   - valid_custom_element_name
//   - customized_built_in_element
func (e *ArticleHTMLElement) IS(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeISKey] = v
	return e
}

func (e *ArticleHTMLElement) IfIS(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.IS(v)
}

func (e *ArticleHTMLElement) RemoveIS(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeISKey)
	return e
}

// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//   - valid_url_potentially_surrounded_by_spaces
func (e *ArticleHTMLElement) ITEMID(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMIDKey] = v
	return e
}

func (e *ArticleHTMLElement) IfITEMID(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMID(v)
}

func (e *ArticleHTMLElement) RemoveITEMID(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeITEMIDKey)
	return e
}

// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
//   - valid_absolute_ur_ls
//   - defined_property_names
func (e *ArticleHTMLElement) ITEMPROP(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMPROPKey] = v
	return e
}

func (e *ArticleHTMLElement) IfITEMPROP(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMPROP(v)
}

func (e *ArticleHTMLElement) RemoveITEMPROP(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeITEMPROPKey)
	return e
}

// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//   - unordered_set_of_unique_space_separated_tokens
func (e *ArticleHTMLElement) ITEMREF(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMREFKey] = v
	return e
}

func (e *ArticleHTMLElement) IfITEMREF(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMREF(v)
}

func (e *ArticleHTMLElement) RemoveITEMREF(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeITEMREFKey)
	return e
}

// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//   - boolean_attribute
func (e *ArticleHTMLElement) ITEMSCOPE() *ArticleHTMLElement {
	if e.BoolAttributes == nil {
		e.BoolAttributes = map[string]struct{}{}
	}
	e.BoolAttributes[attributeITEMSCOPEKey] = struct{}{}
	return e
}

func (e *ArticleHTMLElement) IfITEMSCOPE(cond bool) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMSCOPE()
}

func (e *ArticleHTMLElement) RemoveITEMSCOPE() *ArticleHTMLElement {
	if e.BoolAttributes == nil {
		return e
	}
	delete(e.BoolAttributes, attributeITEMSCOPEKey)
	return e
}

func (e *ArticleHTMLElement) SetITEMSCOPE(b bool) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) ITEMTYPE(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeITEMTYPEKey] = v
	return e
}

func (e *ArticleHTMLElement) IfITEMTYPE(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.ITEMTYPE(v)
}

func (e *ArticleHTMLElement) RemoveITEMTYPE(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeITEMTYPEKey)
	return e
}

// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ArticleHTMLElement) LANG(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeLANGKey] = v
	return e
}

func (e *ArticleHTMLElement) IfLANG(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.LANG(v)
}

func (e *ArticleHTMLElement) RemoveLANG(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeLANGKey)
	return e
}

// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//   - text
func (e *ArticleHTMLElement) NONCE(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeNONCEKey] = v
	return e
}

func (e *ArticleHTMLElement) IfNONCE(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.NONCE(v)
}

func (e *ArticleHTMLElement) RemoveNONCE(v string) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) POPOVER(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributePOPOVERKey] = v
	return e
}

func (e *ArticleHTMLElement) IfPOPOVER(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.POPOVER(v)
}

func (e *ArticleHTMLElement) RemovePOPOVER(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributePOPOVERKey)
	return e
}

// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//   - text
func (e *ArticleHTMLElement) SLOT(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSLOTKey] = v
	return e
}

func (e *ArticleHTMLElement) IfSLOT(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.SLOT(v)
}

func (e *ArticleHTMLElement) RemoveSLOT(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeSLOTKey)
	return e
}

// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//   - true
//   - false
func (e *ArticleHTMLElement) SPELLCHECK(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeSPELLCHECKKey] = v
	return e
}

func (e *ArticleHTMLElement) IfSPELLCHECK(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.SPELLCHECK(v)
}

func (e *ArticleHTMLElement) RemoveSPELLCHECK(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeSPELLCHECKKey)
	return e
}

// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ArticleHTMLElement) STYLE(k, v string) *ArticleHTMLElement {
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

func (e *ArticleHTMLElement) IfSTYLE(cond bool, k string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.STYLE(k, "")
}

func (e *ArticleHTMLElement) RemoveSTYLE(k string) *ArticleHTMLElement {
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
func (e *ArticleHTMLElement) TABINDEX(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTABINDEXKey] = v
	return e
}

func (e *ArticleHTMLElement) IfTABINDEX(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.TABINDEX(v)
}

func (e *ArticleHTMLElement) RemoveTABINDEX(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeTABINDEXKey)
	return e
}

// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//   - text
func (e *ArticleHTMLElement) TITLE(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTITLEKey] = v
	return e
}

func (e *ArticleHTMLElement) IfTITLE(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.TITLE(v)
}

func (e *ArticleHTMLElement) RemoveTITLE(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeTITLEKey)
	return e
}

// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//   - yes
//   - no
func (e *ArticleHTMLElement) TRANSLATE(v string) *ArticleHTMLElement {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	e.StringAttributes[attributeTRANSLATEKey] = v
	return e
}

func (e *ArticleHTMLElement) IfTRANSLATE(cond bool, v string) *ArticleHTMLElement {
	if !cond {
		return e
	}
	return e.TRANSLATE(v)
}

func (e *ArticleHTMLElement) RemoveTRANSLATE(v string) *ArticleHTMLElement {
	delete(e.StringAttributes, attributeTRANSLATEKey)
	return e
}
