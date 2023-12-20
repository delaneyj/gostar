package  html

import (
    "fmt"
)

type TextareaHTMLElement struct {
    *Element
}

func TEXTAREA(children ...ElementBuilder) *TextareaHTMLElement {
    return &TextareaHTMLElement{
        Element: &Element{
            Tag: "textarea",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *TextareaHTMLElement) Children(children ...ElementBuilder) *TextareaHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *TextareaHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TextareaHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *TextareaHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TextareaHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *TextareaHTMLElement) Text(text string) *TextareaHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *TextareaHTMLElement) TextF(format string, args ...any) *TextareaHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *TextareaHTMLElement) Escaped(text string) *TextareaHTMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *TextareaHTMLElement) EscapedF(format string, args ...any) *TextareaHTMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *TextareaHTMLElement) CustomData(key, value string) *TextareaHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TextareaHTMLElement) CustomDataRemove(key string) *TextareaHTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}


// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//  * ordered_set_of_unique_space_separated_tokens
//  * identical_to
func (e *TextareaHTMLElement) ACCESSKEY(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *TextareaHTMLElement) IfACCESSKEY(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.ACCESSKEY(v)
}

func (e *TextareaHTMLElement) RemoveACCESSKEY(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// AUTOCAPITALIZE sets the "autocapitalize" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Values values are constrained to:
//  * on
//  * on
//  * off
//  * off
//  * none
//  * none
//  * sentences
//  * sentences
//  * words
//  * words
//  * characters
//  * characters
func (e *TextareaHTMLElement) AUTOCAPITALIZE(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *TextareaHTMLElement) IfAUTOCAPITALIZE(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCAPITALIZE(v)
}

func (e *TextareaHTMLElement) RemoveAUTOCAPITALIZE(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//  * autofill_field
func (e *TextareaHTMLElement) AUTOCOMPLETE(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocomplete"] = v
    return e
}

func (e *TextareaHTMLElement) IfAUTOCOMPLETE(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOCOMPLETE(v)
}

func (e *TextareaHTMLElement) RemoveAUTOCOMPLETE(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "autocomplete")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaHTMLElement) AUTOFOCUS() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *TextareaHTMLElement) IfAUTOFOCUS(cond bool) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.AUTOFOCUS()
}

func (e *TextareaHTMLElement) RemoveAUTOFOCUS() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *TextareaHTMLElement) SetAUTOFOCUS(b bool) *TextareaHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *TextareaHTMLElement) CLASS(v string) *TextareaHTMLElement {
    if e.DelimitedStringAttributes == nil {
        e.DelimitedStringAttributes = map[string]*DelimitedString{}
    }
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        kv = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = kv
    }
    kv.Add(v)
    return e
}

func (e *TextareaHTMLElement) IfCLASS(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.CLASS(v)
}

func (e *TextareaHTMLElement) SetCLASS(v string) *TextareaHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *TextareaHTMLElement) RemoveCLASS(v string) *TextareaHTMLElement {
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        return e
    }
    kv.Remove(v)
    return e
}
// COLS sets the "cols" attribute.
// Maximum number of characters per line
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *TextareaHTMLElement) COLS(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["cols"] = v
    return e
}

func (e *TextareaHTMLElement) IfCOLS(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.COLS(v)
}

func (e *TextareaHTMLElement) RemoveCOLS(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "cols")
    return e
}
// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//  * true
//  * plaintext_only
//  * false
func (e *TextareaHTMLElement) CONTENTEDITABLE(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *TextareaHTMLElement) IfCONTENTEDITABLE(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.CONTENTEDITABLE(v)
}

func (e *TextareaHTMLElement) RemoveCONTENTEDITABLE(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "contenteditable")
    return e
}
// DIR sets the "dir" attribute.
// The text directionality of the element
// Values values are constrained to:
//  * ltr
//  * ltr
//  * rtl
//  * rtl
func (e *TextareaHTMLElement) DIR(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *TextareaHTMLElement) IfDIR(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.DIR(v)
}

func (e *TextareaHTMLElement) RemoveDIR(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DIRNAME sets the "dirname" attribute.
// Name of form control to use for sending the element's directionality in form submission
// Values values are constrained to:
//  * text
func (e *TextareaHTMLElement) DIRNAME(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dirname"] = v
    return e
}

func (e *TextareaHTMLElement) IfDIRNAME(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.DIRNAME(v)
}

func (e *TextareaHTMLElement) RemoveDIRNAME(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "dirname")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaHTMLElement) DISABLED() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *TextareaHTMLElement) IfDISABLED(cond bool) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.DISABLED()
}

func (e *TextareaHTMLElement) RemoveDISABLED() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *TextareaHTMLElement) SetDISABLED(b bool) *TextareaHTMLElement {
    if b {
        return e.DISABLED()
    }
    return e.RemoveDISABLED()
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *TextareaHTMLElement) DRAGGABLE(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *TextareaHTMLElement) IfDRAGGABLE(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.DRAGGABLE(v)
}

func (e *TextareaHTMLElement) RemoveDRAGGABLE(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "draggable")
    return e
}
// ENTERKEYHINT sets the "enterkeyhint" attribute.
// Hint for selecting an enter key action
// Values values are constrained to:
//  * enter
//  * enter
//  * done
//  * done
//  * go
//  * go
//  * next
//  * next
//  * previous
//  * previous
//  * search
//  * search
//  * send
//  * send
func (e *TextareaHTMLElement) ENTERKEYHINT(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *TextareaHTMLElement) IfENTERKEYHINT(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.ENTERKEYHINT(v)
}

func (e *TextareaHTMLElement) RemoveENTERKEYHINT(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//  * id
func (e *TextareaHTMLElement) FORM(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["form"] = v
    return e
}

func (e *TextareaHTMLElement) IfFORM(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.FORM(v)
}

func (e *TextareaHTMLElement) RemoveFORM(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "form")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *TextareaHTMLElement) HIDDEN(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *TextareaHTMLElement) IfHIDDEN(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.HIDDEN(v)
}

func (e *TextareaHTMLElement) RemoveHIDDEN(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *TextareaHTMLElement) ID(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *TextareaHTMLElement) IfID(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.ID(v)
}

func (e *TextareaHTMLElement) RemoveID(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaHTMLElement) INERT() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *TextareaHTMLElement) IfINERT(cond bool) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.INERT()
}

func (e *TextareaHTMLElement) RemoveINERT() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *TextareaHTMLElement) SetINERT(b bool) *TextareaHTMLElement {
    if b {
        return e.INERT()
    }
    return e.RemoveINERT()
}
// INPUTMODE sets the "inputmode" attribute.
// Hint for selecting an input modality
// Values values are constrained to:
//  * none
//  * none
//  * text
//  * text
//  * tel
//  * tel
//  * email
//  * email
//  * url
//  * url
//  * numeric
//  * numeric
//  * decimal
//  * decimal
//  * search
//  * search
func (e *TextareaHTMLElement) INPUTMODE(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *TextareaHTMLElement) IfINPUTMODE(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.INPUTMODE(v)
}

func (e *TextareaHTMLElement) RemoveINPUTMODE(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *TextareaHTMLElement) IS(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *TextareaHTMLElement) IfIS(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.IS(v)
}

func (e *TextareaHTMLElement) RemoveIS(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *TextareaHTMLElement) ITEMID(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *TextareaHTMLElement) IfITEMID(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMID(v)
}

func (e *TextareaHTMLElement) RemoveITEMID(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *TextareaHTMLElement) ITEMPROP(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *TextareaHTMLElement) IfITEMPROP(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMPROP(v)
}

func (e *TextareaHTMLElement) RemoveITEMPROP(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *TextareaHTMLElement) ITEMREF(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *TextareaHTMLElement) IfITEMREF(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMREF(v)
}

func (e *TextareaHTMLElement) RemoveITEMREF(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaHTMLElement) ITEMSCOPE() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *TextareaHTMLElement) IfITEMSCOPE(cond bool) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMSCOPE()
}

func (e *TextareaHTMLElement) RemoveITEMSCOPE() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *TextareaHTMLElement) SetITEMSCOPE(b bool) *TextareaHTMLElement {
    if b {
        return e.ITEMSCOPE()
    }
    return e.RemoveITEMSCOPE()
}
// ITEMTYPE sets the "itemtype" attribute.
// Item types of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
func (e *TextareaHTMLElement) ITEMTYPE(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *TextareaHTMLElement) IfITEMTYPE(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.ITEMTYPE(v)
}

func (e *TextareaHTMLElement) RemoveITEMTYPE(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TextareaHTMLElement) LANG(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *TextareaHTMLElement) IfLANG(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.LANG(v)
}

func (e *TextareaHTMLElement) RemoveLANG(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// MAXLENGTH sets the "maxlength" attribute.
// Maximum length of value
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *TextareaHTMLElement) MAXLENGTH(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["maxlength"] = v
    return e
}

func (e *TextareaHTMLElement) IfMAXLENGTH(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.MAXLENGTH(v)
}

func (e *TextareaHTMLElement) RemoveMAXLENGTH(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "maxlength")
    return e
}
// MINLENGTH sets the "minlength" attribute.
// Minimum length of value
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *TextareaHTMLElement) MINLENGTH(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["minlength"] = v
    return e
}

func (e *TextareaHTMLElement) IfMINLENGTH(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.MINLENGTH(v)
}

func (e *TextareaHTMLElement) RemoveMINLENGTH(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "minlength")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *TextareaHTMLElement) NAME(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *TextareaHTMLElement) IfNAME(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.NAME(v)
}

func (e *TextareaHTMLElement) RemoveNAME(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *TextareaHTMLElement) NONCE(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *TextareaHTMLElement) IfNONCE(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.NONCE(v)
}

func (e *TextareaHTMLElement) RemoveNONCE(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// PLACEHOLDER sets the "placeholder" attribute.
// User-visible label to be placed within the form control
// Values values are constrained to:
//  * text
func (e *TextareaHTMLElement) PLACEHOLDER(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["placeholder"] = v
    return e
}

func (e *TextareaHTMLElement) IfPLACEHOLDER(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.PLACEHOLDER(v)
}

func (e *TextareaHTMLElement) RemovePLACEHOLDER(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "placeholder")
    return e
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *TextareaHTMLElement) POPOVER(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *TextareaHTMLElement) IfPOPOVER(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.POPOVER(v)
}

func (e *TextareaHTMLElement) RemovePOPOVER(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// READONLY sets the "readonly" attribute.
// Affects willValidate, plus any behavior added by the custom element author
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaHTMLElement) READONLY() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["readonly"] = struct{}{}
    return e
}

func (e *TextareaHTMLElement) IfREADONLY(cond bool) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.READONLY()
}

func (e *TextareaHTMLElement) RemoveREADONLY() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "readonly")
    return e
}

func (e *TextareaHTMLElement) SetREADONLY(b bool) *TextareaHTMLElement {
    if b {
        return e.READONLY()
    }
    return e.RemoveREADONLY()
}
// REQUIRED sets the "required" attribute.
// Whether the control is required for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaHTMLElement) REQUIRED() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["required"] = struct{}{}
    return e
}

func (e *TextareaHTMLElement) IfREQUIRED(cond bool) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.REQUIRED()
}

func (e *TextareaHTMLElement) RemoveREQUIRED() *TextareaHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "required")
    return e
}

func (e *TextareaHTMLElement) SetREQUIRED(b bool) *TextareaHTMLElement {
    if b {
        return e.REQUIRED()
    }
    return e.RemoveREQUIRED()
}
// ROWS sets the "rows" attribute.
// Number of lines to show
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *TextareaHTMLElement) ROWS(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["rows"] = v
    return e
}

func (e *TextareaHTMLElement) IfROWS(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.ROWS(v)
}

func (e *TextareaHTMLElement) RemoveROWS(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "rows")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *TextareaHTMLElement) SLOT(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *TextareaHTMLElement) IfSLOT(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.SLOT(v)
}

func (e *TextareaHTMLElement) RemoveSLOT(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *TextareaHTMLElement) SPELLCHECK(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *TextareaHTMLElement) IfSPELLCHECK(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.SPELLCHECK(v)
}

func (e *TextareaHTMLElement) RemoveSPELLCHECK(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TextareaHTMLElement) STYLE(k,v string) *TextareaHTMLElement {
    if e.DelimitedKVAttributes == nil {
        e.DelimitedKVAttributes = map[string]*DelimitedKVString{}
    }
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        kv = NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = kv
    }
    kv.Add(k,v)
    return e
}

func (e *TextareaHTMLElement) IfSTYLE(cond bool, k string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.STYLE(k, "")
}

func (e *TextareaHTMLElement) RemoveSTYLE(k string) *TextareaHTMLElement {
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        return e
    }
    kv.Remove(k)
    return e
}
// TABINDEX sets the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and
//      the relative order of the element for the purposes of sequential focus navigation
// Values values are constrained to:
//  * valid_integer
func (e *TextareaHTMLElement) TABINDEX(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *TextareaHTMLElement) IfTABINDEX(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.TABINDEX(v)
}

func (e *TextareaHTMLElement) RemoveTABINDEX(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *TextareaHTMLElement) TITLE(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *TextareaHTMLElement) IfTITLE(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.TITLE(v)
}

func (e *TextareaHTMLElement) RemoveTITLE(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *TextareaHTMLElement) TRANSLATE(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *TextareaHTMLElement) IfTRANSLATE(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.TRANSLATE(v)
}

func (e *TextareaHTMLElement) RemoveTRANSLATE(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// WRAP sets the "wrap" attribute.
// How the value of the form control is to be wrapped for form submission
// Values values are constrained to:
//  * soft
//  * soft
//  * hard
//  * hard
func (e *TextareaHTMLElement) WRAP(v string) *TextareaHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["wrap"] = v
    return e
}

func (e *TextareaHTMLElement) IfWRAP(cond bool, v string) *TextareaHTMLElement {
    if !cond {
        return e
    }
    return e.WRAP(v)
}

func (e *TextareaHTMLElement) RemoveWRAP(v string) *TextareaHTMLElement {
    delete(e.StringAttributes, "wrap")
    return e
}
