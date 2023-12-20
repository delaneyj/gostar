package  html

import (
    "fmt"
)

type InputHTMLElement struct {
    *Element
}

func INPUT(children ...fmt.Stringer) *InputHTMLElement {
    return &InputHTMLElement{
        Element: &Element{
            Tag: "input",
            IsSelfClosing: true,
            Descendants: children,
        },
    }
}

func (e *InputHTMLElement) Children(children ...fmt.Stringer) *InputHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *InputHTMLElement) IfChildren(condition bool, children ...fmt.Stringer) *InputHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *InputHTMLElement) TernChildren(condition bool, trueChildren, falseChildren fmt.Stringer) *InputHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *InputHTMLElement) Text(text string) *InputHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *InputHTMLElement) TextF(format string, args ...any) *InputHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *InputHTMLElement) Raw(text string) *InputHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *InputHTMLElement) RawF(format string, args ...any) *InputHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *InputHTMLElement) CustomData(key, value string) *InputHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *InputHTMLElement) CustomDataRemove(key string) *InputHTMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	delete(e.CustomDataAttributes, key)
	return e
}


// ACCEPT sets the "accept" attribute.
// Hint for expected file type in file upload controls
// Values values are constrained to:
//  * set_of_comma_separated_tokens
//  * valid_mime_type_strings_with_no_parameters
//  * audio/*
//  * video/*
//  * image/*
func (e *InputHTMLElement) ACCEPT(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accept"] = v
    return e
}

func (e *InputHTMLElement) RemoveACCEPT(v string) *InputHTMLElement {
    delete(e.StringAttributes, "accept")
    return e
}
// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//  * ordered_set_of_unique_space_separated_tokens
//  * identical_to
func (e *InputHTMLElement) ACCESSKEY(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *InputHTMLElement) RemoveACCESSKEY(v string) *InputHTMLElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ALT sets the "alt" attribute.
// Replacement text for use when images are not available
// Values values are constrained to:
//  * text
func (e *InputHTMLElement) ALT(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["alt"] = v
    return e
}

func (e *InputHTMLElement) RemoveALT(v string) *InputHTMLElement {
    delete(e.StringAttributes, "alt")
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
func (e *InputHTMLElement) AUTOCAPITALIZE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *InputHTMLElement) RemoveAUTOCAPITALIZE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//  * autofill_field
func (e *InputHTMLElement) AUTOCOMPLETE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocomplete"] = v
    return e
}

func (e *InputHTMLElement) RemoveAUTOCOMPLETE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "autocomplete")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *InputHTMLElement) AUTOFOCUS() *InputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *InputHTMLElement) RemoveAUTOFOCUS() *InputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *InputHTMLElement) SetAUTOFOCUS(b bool) *InputHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CHECKED sets the "checked" attribute.
// Whether the control is checked
// Values values are constrained to:
//  * boolean_attribute
func (e *InputHTMLElement) CHECKED() *InputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["checked"] = struct{}{}
    return e
}

func (e *InputHTMLElement) RemoveCHECKED() *InputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "checked")
    return e
}

func (e *InputHTMLElement) SetCHECKED(b bool) *InputHTMLElement {
    if b {
        return e.CHECKED()
    }
    return e.RemoveCHECKED()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *InputHTMLElement) CLASS(v string) *InputHTMLElement {
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

func (e *InputHTMLElement) SetCLASS(v string) *InputHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *InputHTMLElement) RemoveCLASS(v string) *InputHTMLElement {
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        return e
    }
    kv.Remove(v)
    return e
}
// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//  * true
//  * plaintext_only
//  * false
func (e *InputHTMLElement) CONTENTEDITABLE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *InputHTMLElement) RemoveCONTENTEDITABLE(v string) *InputHTMLElement {
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
func (e *InputHTMLElement) DIR(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *InputHTMLElement) RemoveDIR(v string) *InputHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DIRNAME sets the "dirname" attribute.
// Name of form control to use for sending the element's directionality in form submission
// Values values are constrained to:
//  * text
func (e *InputHTMLElement) DIRNAME(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dirname"] = v
    return e
}

func (e *InputHTMLElement) RemoveDIRNAME(v string) *InputHTMLElement {
    delete(e.StringAttributes, "dirname")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *InputHTMLElement) DISABLED() *InputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *InputHTMLElement) RemoveDISABLED() *InputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *InputHTMLElement) SetDISABLED(b bool) *InputHTMLElement {
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
func (e *InputHTMLElement) DRAGGABLE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *InputHTMLElement) RemoveDRAGGABLE(v string) *InputHTMLElement {
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
func (e *InputHTMLElement) ENTERKEYHINT(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *InputHTMLElement) RemoveENTERKEYHINT(v string) *InputHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//  * id
func (e *InputHTMLElement) FORM(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["form"] = v
    return e
}

func (e *InputHTMLElement) RemoveFORM(v string) *InputHTMLElement {
    delete(e.StringAttributes, "form")
    return e
}
// FORMACTION sets the "formaction" attribute.
// URL to use for form submission
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *InputHTMLElement) FORMACTION(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formaction"] = v
    return e
}

func (e *InputHTMLElement) RemoveFORMACTION(v string) *InputHTMLElement {
    delete(e.StringAttributes, "formaction")
    return e
}
// FORMENCTYPE sets the "formenctype" attribute.
// Entry list encoding type to use for form submission
// Values values are constrained to:
//  * application/x_www_form_urlencoded
//  * application/x_www_form_urlencoded
//  * multipart/form_data
//  * multipart/form_data
//  * text/plain
//  * text/plain
func (e *InputHTMLElement) FORMENCTYPE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formenctype"] = v
    return e
}

func (e *InputHTMLElement) RemoveFORMENCTYPE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "formenctype")
    return e
}
// FORMMETHOD sets the "formmethod" attribute.
// Variant to use for form submission
// Values values are constrained to:
//  * get
//  * post
//  * dialog
func (e *InputHTMLElement) FORMMETHOD(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formmethod"] = v
    return e
}

func (e *InputHTMLElement) RemoveFORMMETHOD(v string) *InputHTMLElement {
    delete(e.StringAttributes, "formmethod")
    return e
}
// FORMNOVALIDATE sets the "formnovalidate" attribute.
// Bypass form control validation for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *InputHTMLElement) FORMNOVALIDATE() *InputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["formnovalidate"] = struct{}{}
    return e
}

func (e *InputHTMLElement) RemoveFORMNOVALIDATE() *InputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "formnovalidate")
    return e
}

func (e *InputHTMLElement) SetFORMNOVALIDATE(b bool) *InputHTMLElement {
    if b {
        return e.FORMNOVALIDATE()
    }
    return e.RemoveFORMNOVALIDATE()
}
// FORMTARGET sets the "formtarget" attribute.
// Navigable for form submission
// Values values are constrained to:
//  * valid_navigable_target_name_or_keyword
func (e *InputHTMLElement) FORMTARGET(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formtarget"] = v
    return e
}

func (e *InputHTMLElement) RemoveFORMTARGET(v string) *InputHTMLElement {
    delete(e.StringAttributes, "formtarget")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputHTMLElement) HEIGHT(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *InputHTMLElement) RemoveHEIGHT(v string) *InputHTMLElement {
    delete(e.StringAttributes, "height")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *InputHTMLElement) HIDDEN(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *InputHTMLElement) RemoveHIDDEN(v string) *InputHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *InputHTMLElement) ID(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *InputHTMLElement) RemoveID(v string) *InputHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *InputHTMLElement) INERT() *InputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *InputHTMLElement) RemoveINERT() *InputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *InputHTMLElement) SetINERT(b bool) *InputHTMLElement {
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
func (e *InputHTMLElement) INPUTMODE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *InputHTMLElement) RemoveINPUTMODE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *InputHTMLElement) IS(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *InputHTMLElement) RemoveIS(v string) *InputHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *InputHTMLElement) ITEMID(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *InputHTMLElement) RemoveITEMID(v string) *InputHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *InputHTMLElement) ITEMPROP(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *InputHTMLElement) RemoveITEMPROP(v string) *InputHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *InputHTMLElement) ITEMREF(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *InputHTMLElement) RemoveITEMREF(v string) *InputHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *InputHTMLElement) ITEMSCOPE() *InputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *InputHTMLElement) RemoveITEMSCOPE() *InputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *InputHTMLElement) SetITEMSCOPE(b bool) *InputHTMLElement {
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
func (e *InputHTMLElement) ITEMTYPE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *InputHTMLElement) RemoveITEMTYPE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *InputHTMLElement) LANG(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *InputHTMLElement) RemoveLANG(v string) *InputHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LIST sets the "list" attribute.
// List of autocomplete options
// Values values are constrained to:
//  * id
func (e *InputHTMLElement) LIST(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["list"] = v
    return e
}

func (e *InputHTMLElement) RemoveLIST(v string) *InputHTMLElement {
    delete(e.StringAttributes, "list")
    return e
}
// MAX sets the "max" attribute.
// Upper bound of range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *InputHTMLElement) MAX(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["max"] = v
    return e
}

func (e *InputHTMLElement) RemoveMAX(v string) *InputHTMLElement {
    delete(e.StringAttributes, "max")
    return e
}
// MAXLENGTH sets the "maxlength" attribute.
// Maximum length of value
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputHTMLElement) MAXLENGTH(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["maxlength"] = v
    return e
}

func (e *InputHTMLElement) RemoveMAXLENGTH(v string) *InputHTMLElement {
    delete(e.StringAttributes, "maxlength")
    return e
}
// MIN sets the "min" attribute.
// Lower bound of range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *InputHTMLElement) MIN(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["min"] = v
    return e
}

func (e *InputHTMLElement) RemoveMIN(v string) *InputHTMLElement {
    delete(e.StringAttributes, "min")
    return e
}
// MINLENGTH sets the "minlength" attribute.
// Minimum length of value
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputHTMLElement) MINLENGTH(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["minlength"] = v
    return e
}

func (e *InputHTMLElement) RemoveMINLENGTH(v string) *InputHTMLElement {
    delete(e.StringAttributes, "minlength")
    return e
}
// MULTIPLE sets the "multiple" attribute.
// Whether to allow multiple values
// Values values are constrained to:
//  * boolean_attribute
func (e *InputHTMLElement) MULTIPLE() *InputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["multiple"] = struct{}{}
    return e
}

func (e *InputHTMLElement) RemoveMULTIPLE() *InputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "multiple")
    return e
}

func (e *InputHTMLElement) SetMULTIPLE(b bool) *InputHTMLElement {
    if b {
        return e.MULTIPLE()
    }
    return e.RemoveMULTIPLE()
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *InputHTMLElement) NAME(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *InputHTMLElement) RemoveNAME(v string) *InputHTMLElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *InputHTMLElement) NONCE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *InputHTMLElement) RemoveNONCE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// PATTERN sets the "pattern" attribute.
// Pattern to be matched by the form control's value
// Values values are constrained to:
//  * pattern
func (e *InputHTMLElement) PATTERN(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["pattern"] = v
    return e
}

func (e *InputHTMLElement) RemovePATTERN(v string) *InputHTMLElement {
    delete(e.StringAttributes, "pattern")
    return e
}
// PLACEHOLDER sets the "placeholder" attribute.
// User-visible label to be placed within the form control
// Values values are constrained to:
//  * text
func (e *InputHTMLElement) PLACEHOLDER(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["placeholder"] = v
    return e
}

func (e *InputHTMLElement) RemovePLACEHOLDER(v string) *InputHTMLElement {
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
func (e *InputHTMLElement) POPOVER(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *InputHTMLElement) RemovePOPOVER(v string) *InputHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// POPOVERTARGET sets the "popovertarget" attribute.
// Targets a popover element to toggle, show, or hide
// Values values are constrained to:
//  * id
func (e *InputHTMLElement) POPOVERTARGET(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popovertarget"] = v
    return e
}

func (e *InputHTMLElement) RemovePOPOVERTARGET(v string) *InputHTMLElement {
    delete(e.StringAttributes, "popovertarget")
    return e
}
// POPOVERTARGETACTION sets the "popovertargetaction" attribute.
// Indicates whether a targeted popover element is to be toggled, shown, or hidden
// Values values are constrained to:
//  * toggle
//  * toggle
//  * show
//  * show
//  * hide
//  * hide
func (e *InputHTMLElement) POPOVERTARGETACTION(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popovertargetaction"] = v
    return e
}

func (e *InputHTMLElement) RemovePOPOVERTARGETACTION(v string) *InputHTMLElement {
    delete(e.StringAttributes, "popovertargetaction")
    return e
}
// READONLY sets the "readonly" attribute.
// Affects willValidate, plus any behavior added by the custom element author
// Values values are constrained to:
//  * boolean_attribute
func (e *InputHTMLElement) READONLY() *InputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["readonly"] = struct{}{}
    return e
}

func (e *InputHTMLElement) RemoveREADONLY() *InputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "readonly")
    return e
}

func (e *InputHTMLElement) SetREADONLY(b bool) *InputHTMLElement {
    if b {
        return e.READONLY()
    }
    return e.RemoveREADONLY()
}
// REQUIRED sets the "required" attribute.
// Whether the control is required for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *InputHTMLElement) REQUIRED() *InputHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["required"] = struct{}{}
    return e
}

func (e *InputHTMLElement) RemoveREQUIRED() *InputHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "required")
    return e
}

func (e *InputHTMLElement) SetREQUIRED(b bool) *InputHTMLElement {
    if b {
        return e.REQUIRED()
    }
    return e.RemoveREQUIRED()
}
// SIZE sets the "size" attribute.
// Size of the control
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputHTMLElement) SIZE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["size"] = v
    return e
}

func (e *InputHTMLElement) RemoveSIZE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "size")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *InputHTMLElement) SLOT(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *InputHTMLElement) RemoveSLOT(v string) *InputHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *InputHTMLElement) SPELLCHECK(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *InputHTMLElement) RemoveSPELLCHECK(v string) *InputHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *InputHTMLElement) SRC(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *InputHTMLElement) RemoveSRC(v string) *InputHTMLElement {
    delete(e.StringAttributes, "src")
    return e
}
// STEP sets the "step" attribute.
// Granularity to be matched by the form control's value
// Values values are constrained to:
//  * valid_floating_point_number
//  * any
func (e *InputHTMLElement) STEP(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["step"] = v
    return e
}

func (e *InputHTMLElement) RemoveSTEP(v string) *InputHTMLElement {
    delete(e.StringAttributes, "step")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *InputHTMLElement) STYLE(k,v string) *InputHTMLElement {
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

func (e *InputHTMLElement) RemoveSTYLE(k string) *InputHTMLElement {
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
func (e *InputHTMLElement) TABINDEX(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *InputHTMLElement) RemoveTABINDEX(v string) *InputHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *InputHTMLElement) TITLE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *InputHTMLElement) RemoveTITLE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *InputHTMLElement) TRANSLATE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *InputHTMLElement) RemoveTRANSLATE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *InputHTMLElement) TYPE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *InputHTMLElement) RemoveTYPE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "type")
    return e
}
// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//  * valid_floating_point_number
func (e *InputHTMLElement) VALUE(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["value"] = v
    return e
}

func (e *InputHTMLElement) RemoveVALUE(v string) *InputHTMLElement {
    delete(e.StringAttributes, "value")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputHTMLElement) WIDTH(v string) *InputHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *InputHTMLElement) RemoveWIDTH(v string) *InputHTMLElement {
    delete(e.StringAttributes, "width")
    return e
}
