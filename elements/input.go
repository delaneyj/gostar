package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type InputElement struct {
    *gostar.Element
}

func INPUT(children ...fmt.Stringer) *InputElement {
    return &InputElement{
        Element: &gostar.Element{
            Tag: "input",
            IsSelfClosing: true,
            Children: children,
        },
    }
}

func (e *InputElement) AddChildren(children ...fmt.Stringer) *InputElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *InputElement) TEXT(text string) *InputElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *InputElement) RAW(text string) *InputElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *InputElement) CustomData(key, value string) *InputElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *InputElement) CustomDataRemove(key string) *InputElement {
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
func (e *InputElement) ACCEPT(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accept"] = v
    return e
}

func (e *InputElement) RemoveACCEPT(v string) *InputElement {
    delete(e.StringAttributes, "accept")
    return e
}
// ACCESSKEY sets the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Values values are constrained to:
//  * ordered_set_of_unique_space_separated_tokens
//  * identical_to
func (e *InputElement) ACCESSKEY(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *InputElement) RemoveACCESSKEY(v string) *InputElement {
    delete(e.StringAttributes, "accesskey")
    return e
}
// ALT sets the "alt" attribute.
// Replacement text for use when images are not available
// Values values are constrained to:
//  * text
func (e *InputElement) ALT(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["alt"] = v
    return e
}

func (e *InputElement) RemoveALT(v string) *InputElement {
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
func (e *InputElement) AUTOCAPITALIZE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *InputElement) RemoveAUTOCAPITALIZE(v string) *InputElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//  * autofill_field
func (e *InputElement) AUTOCOMPLETE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocomplete"] = v
    return e
}

func (e *InputElement) RemoveAUTOCOMPLETE(v string) *InputElement {
    delete(e.StringAttributes, "autocomplete")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *InputElement) AUTOFOCUS() *InputElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *InputElement) RemoveAUTOFOCUS() *InputElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *InputElement) SetAUTOFOCUS(b bool) *InputElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CHECKED sets the "checked" attribute.
// Whether the control is checked
// Values values are constrained to:
//  * boolean_attribute
func (e *InputElement) CHECKED() *InputElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["checked"] = struct{}{}
    return e
}

func (e *InputElement) RemoveCHECKED() *InputElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "checked")
    return e
}

func (e *InputElement) SetCHECKED(b bool) *InputElement {
    if b {
        return e.CHECKED()
    }
    return e.RemoveCHECKED()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *InputElement) CLASS(v string) *InputElement {
    if e.DelimitedStringAttributes == nil {
        e.DelimitedStringAttributes = map[string]*gostar.DelimitedString{}
    }
    kv, ok := e.DelimitedStringAttributes["class"]
    if !ok {
        kv = gostar.NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = kv
    }
    kv.Add(v)
    return e
}

func (e *InputElement) SetCLASS(v string) *InputElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *InputElement) RemoveCLASS(v string) *InputElement {
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
func (e *InputElement) CONTENTEDITABLE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *InputElement) RemoveCONTENTEDITABLE(v string) *InputElement {
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
func (e *InputElement) DIR(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *InputElement) RemoveDIR(v string) *InputElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DIRNAME sets the "dirname" attribute.
// Name of form control to use for sending the element's directionality in form submission
// Values values are constrained to:
//  * text
func (e *InputElement) DIRNAME(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dirname"] = v
    return e
}

func (e *InputElement) RemoveDIRNAME(v string) *InputElement {
    delete(e.StringAttributes, "dirname")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *InputElement) DISABLED() *InputElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *InputElement) RemoveDISABLED() *InputElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *InputElement) SetDISABLED(b bool) *InputElement {
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
func (e *InputElement) DRAGGABLE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *InputElement) RemoveDRAGGABLE(v string) *InputElement {
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
func (e *InputElement) ENTERKEYHINT(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *InputElement) RemoveENTERKEYHINT(v string) *InputElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//  * id
func (e *InputElement) FORM(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["form"] = v
    return e
}

func (e *InputElement) RemoveFORM(v string) *InputElement {
    delete(e.StringAttributes, "form")
    return e
}
// FORMACTION sets the "formaction" attribute.
// URL to use for form submission
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *InputElement) FORMACTION(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formaction"] = v
    return e
}

func (e *InputElement) RemoveFORMACTION(v string) *InputElement {
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
func (e *InputElement) FORMENCTYPE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formenctype"] = v
    return e
}

func (e *InputElement) RemoveFORMENCTYPE(v string) *InputElement {
    delete(e.StringAttributes, "formenctype")
    return e
}
// FORMMETHOD sets the "formmethod" attribute.
// Variant to use for form submission
// Values values are constrained to:
//  * get
//  * post
//  * dialog
func (e *InputElement) FORMMETHOD(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formmethod"] = v
    return e
}

func (e *InputElement) RemoveFORMMETHOD(v string) *InputElement {
    delete(e.StringAttributes, "formmethod")
    return e
}
// FORMNOVALIDATE sets the "formnovalidate" attribute.
// Bypass form control validation for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *InputElement) FORMNOVALIDATE() *InputElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["formnovalidate"] = struct{}{}
    return e
}

func (e *InputElement) RemoveFORMNOVALIDATE() *InputElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "formnovalidate")
    return e
}

func (e *InputElement) SetFORMNOVALIDATE(b bool) *InputElement {
    if b {
        return e.FORMNOVALIDATE()
    }
    return e.RemoveFORMNOVALIDATE()
}
// FORMTARGET sets the "formtarget" attribute.
// Navigable for form submission
// Values values are constrained to:
//  * valid_navigable_target_name_or_keyword
func (e *InputElement) FORMTARGET(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formtarget"] = v
    return e
}

func (e *InputElement) RemoveFORMTARGET(v string) *InputElement {
    delete(e.StringAttributes, "formtarget")
    return e
}
// HEIGHT sets the "height" attribute.
// Vertical dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputElement) HEIGHT(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["height"] = v
    return e
}

func (e *InputElement) RemoveHEIGHT(v string) *InputElement {
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
func (e *InputElement) HIDDEN(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *InputElement) RemoveHIDDEN(v string) *InputElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *InputElement) ID(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *InputElement) RemoveID(v string) *InputElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *InputElement) INERT() *InputElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *InputElement) RemoveINERT() *InputElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *InputElement) SetINERT(b bool) *InputElement {
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
func (e *InputElement) INPUTMODE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *InputElement) RemoveINPUTMODE(v string) *InputElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *InputElement) IS(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *InputElement) RemoveIS(v string) *InputElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *InputElement) ITEMID(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *InputElement) RemoveITEMID(v string) *InputElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *InputElement) ITEMPROP(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *InputElement) RemoveITEMPROP(v string) *InputElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *InputElement) ITEMREF(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *InputElement) RemoveITEMREF(v string) *InputElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *InputElement) ITEMSCOPE() *InputElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *InputElement) RemoveITEMSCOPE() *InputElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *InputElement) SetITEMSCOPE(b bool) *InputElement {
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
func (e *InputElement) ITEMTYPE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *InputElement) RemoveITEMTYPE(v string) *InputElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *InputElement) LANG(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *InputElement) RemoveLANG(v string) *InputElement {
    delete(e.StringAttributes, "lang")
    return e
}
// LIST sets the "list" attribute.
// List of autocomplete options
// Values values are constrained to:
//  * id
func (e *InputElement) LIST(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["list"] = v
    return e
}

func (e *InputElement) RemoveLIST(v string) *InputElement {
    delete(e.StringAttributes, "list")
    return e
}
// MAX sets the "max" attribute.
// Upper bound of range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *InputElement) MAX(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["max"] = v
    return e
}

func (e *InputElement) RemoveMAX(v string) *InputElement {
    delete(e.StringAttributes, "max")
    return e
}
// MAXLENGTH sets the "maxlength" attribute.
// Maximum length of value
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputElement) MAXLENGTH(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["maxlength"] = v
    return e
}

func (e *InputElement) RemoveMAXLENGTH(v string) *InputElement {
    delete(e.StringAttributes, "maxlength")
    return e
}
// MIN sets the "min" attribute.
// Lower bound of range
// Values values are constrained to:
//  * valid_floating_point_number
func (e *InputElement) MIN(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["min"] = v
    return e
}

func (e *InputElement) RemoveMIN(v string) *InputElement {
    delete(e.StringAttributes, "min")
    return e
}
// MINLENGTH sets the "minlength" attribute.
// Minimum length of value
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputElement) MINLENGTH(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["minlength"] = v
    return e
}

func (e *InputElement) RemoveMINLENGTH(v string) *InputElement {
    delete(e.StringAttributes, "minlength")
    return e
}
// MULTIPLE sets the "multiple" attribute.
// Whether to allow multiple values
// Values values are constrained to:
//  * boolean_attribute
func (e *InputElement) MULTIPLE() *InputElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["multiple"] = struct{}{}
    return e
}

func (e *InputElement) RemoveMULTIPLE() *InputElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "multiple")
    return e
}

func (e *InputElement) SetMULTIPLE(b bool) *InputElement {
    if b {
        return e.MULTIPLE()
    }
    return e.RemoveMULTIPLE()
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *InputElement) NAME(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *InputElement) RemoveNAME(v string) *InputElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *InputElement) NONCE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *InputElement) RemoveNONCE(v string) *InputElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// PATTERN sets the "pattern" attribute.
// Pattern to be matched by the form control's value
// Values values are constrained to:
//  * pattern
func (e *InputElement) PATTERN(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["pattern"] = v
    return e
}

func (e *InputElement) RemovePATTERN(v string) *InputElement {
    delete(e.StringAttributes, "pattern")
    return e
}
// PLACEHOLDER sets the "placeholder" attribute.
// User-visible label to be placed within the form control
// Values values are constrained to:
//  * text
func (e *InputElement) PLACEHOLDER(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["placeholder"] = v
    return e
}

func (e *InputElement) RemovePLACEHOLDER(v string) *InputElement {
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
func (e *InputElement) POPOVER(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *InputElement) RemovePOPOVER(v string) *InputElement {
    delete(e.StringAttributes, "popover")
    return e
}
// POPOVERTARGET sets the "popovertarget" attribute.
// Targets a popover element to toggle, show, or hide
// Values values are constrained to:
//  * id
func (e *InputElement) POPOVERTARGET(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popovertarget"] = v
    return e
}

func (e *InputElement) RemovePOPOVERTARGET(v string) *InputElement {
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
func (e *InputElement) POPOVERTARGETACTION(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popovertargetaction"] = v
    return e
}

func (e *InputElement) RemovePOPOVERTARGETACTION(v string) *InputElement {
    delete(e.StringAttributes, "popovertargetaction")
    return e
}
// READONLY sets the "readonly" attribute.
// Affects willValidate, plus any behavior added by the custom element author
// Values values are constrained to:
//  * boolean_attribute
func (e *InputElement) READONLY() *InputElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["readonly"] = struct{}{}
    return e
}

func (e *InputElement) RemoveREADONLY() *InputElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "readonly")
    return e
}

func (e *InputElement) SetREADONLY(b bool) *InputElement {
    if b {
        return e.READONLY()
    }
    return e.RemoveREADONLY()
}
// REQUIRED sets the "required" attribute.
// Whether the control is required for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *InputElement) REQUIRED() *InputElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["required"] = struct{}{}
    return e
}

func (e *InputElement) RemoveREQUIRED() *InputElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "required")
    return e
}

func (e *InputElement) SetREQUIRED(b bool) *InputElement {
    if b {
        return e.REQUIRED()
    }
    return e.RemoveREQUIRED()
}
// SIZE sets the "size" attribute.
// Size of the control
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputElement) SIZE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["size"] = v
    return e
}

func (e *InputElement) RemoveSIZE(v string) *InputElement {
    delete(e.StringAttributes, "size")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *InputElement) SLOT(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *InputElement) RemoveSLOT(v string) *InputElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *InputElement) SPELLCHECK(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *InputElement) RemoveSPELLCHECK(v string) *InputElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// SRC sets the "src" attribute.
// Address of the resource
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *InputElement) SRC(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["src"] = v
    return e
}

func (e *InputElement) RemoveSRC(v string) *InputElement {
    delete(e.StringAttributes, "src")
    return e
}
// STEP sets the "step" attribute.
// Granularity to be matched by the form control's value
// Values values are constrained to:
//  * valid_floating_point_number
//  * any
func (e *InputElement) STEP(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["step"] = v
    return e
}

func (e *InputElement) RemoveSTEP(v string) *InputElement {
    delete(e.StringAttributes, "step")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *InputElement) STYLE(k,v string) *InputElement {
    if e.DelimitedKVAttributes == nil {
        e.DelimitedKVAttributes = map[string]*gostar.DelimitedKVString{}
    }
    kv, ok := e.DelimitedKVAttributes["style"]
    if !ok {
        kv = gostar.NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = kv
    }
    kv.Add(k,v)
    return e
}

func (e *InputElement) RemoveSTYLE(k string) *InputElement {
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
func (e *InputElement) TABINDEX(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *InputElement) RemoveTABINDEX(v string) *InputElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *InputElement) TITLE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *InputElement) RemoveTITLE(v string) *InputElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *InputElement) TRANSLATE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *InputElement) RemoveTRANSLATE(v string) *InputElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *InputElement) TYPE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *InputElement) RemoveTYPE(v string) *InputElement {
    delete(e.StringAttributes, "type")
    return e
}
// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//  * valid_floating_point_number
func (e *InputElement) VALUE(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["value"] = v
    return e
}

func (e *InputElement) RemoveVALUE(v string) *InputElement {
    delete(e.StringAttributes, "value")
    return e
}
// WIDTH sets the "width" attribute.
// Horizontal dimension
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *InputElement) WIDTH(v string) *InputElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["width"] = v
    return e
}

func (e *InputElement) RemoveWIDTH(v string) *InputElement {
    delete(e.StringAttributes, "width")
    return e
}
