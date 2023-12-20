package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type TextareaElement struct {
    *gostar.Element
}

func TEXTAREA(children ...fmt.Stringer) *TextareaElement {
    return &TextareaElement{
        Element: &gostar.Element{
            Tag: "textarea",
            IsSelfClosing: false,
            Children: children,
        },
    }
}

func (e *TextareaElement) AddChildren(children ...fmt.Stringer) *TextareaElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *TextareaElement) TEXT(text string) *TextareaElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *TextareaElement) RAW(text string) *TextareaElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *TextareaElement) CustomData(key, value string) *TextareaElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TextareaElement) CustomDataRemove(key string) *TextareaElement {
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
func (e *TextareaElement) ACCESSKEY(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *TextareaElement) RemoveACCESSKEY(v string) *TextareaElement {
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
func (e *TextareaElement) AUTOCAPITALIZE(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *TextareaElement) RemoveAUTOCAPITALIZE(v string) *TextareaElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOCOMPLETE sets the "autocomplete" attribute.
// Hint for form autofill feature
// Values values are constrained to:
//  * autofill_field
func (e *TextareaElement) AUTOCOMPLETE(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocomplete"] = v
    return e
}

func (e *TextareaElement) RemoveAUTOCOMPLETE(v string) *TextareaElement {
    delete(e.StringAttributes, "autocomplete")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaElement) AUTOFOCUS() *TextareaElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *TextareaElement) RemoveAUTOFOCUS() *TextareaElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *TextareaElement) SetAUTOFOCUS(b bool) *TextareaElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *TextareaElement) CLASS(v string) *TextareaElement {
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

func (e *TextareaElement) SetCLASS(v string) *TextareaElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *TextareaElement) RemoveCLASS(v string) *TextareaElement {
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
func (e *TextareaElement) COLS(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["cols"] = v
    return e
}

func (e *TextareaElement) RemoveCOLS(v string) *TextareaElement {
    delete(e.StringAttributes, "cols")
    return e
}
// CONTENTEDITABLE sets the "contenteditable" attribute.
// Whether the element is editable
// Values values are constrained to:
//  * true
//  * plaintext_only
//  * false
func (e *TextareaElement) CONTENTEDITABLE(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *TextareaElement) RemoveCONTENTEDITABLE(v string) *TextareaElement {
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
func (e *TextareaElement) DIR(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *TextareaElement) RemoveDIR(v string) *TextareaElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DIRNAME sets the "dirname" attribute.
// Name of form control to use for sending the element's directionality in form submission
// Values values are constrained to:
//  * text
func (e *TextareaElement) DIRNAME(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dirname"] = v
    return e
}

func (e *TextareaElement) RemoveDIRNAME(v string) *TextareaElement {
    delete(e.StringAttributes, "dirname")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaElement) DISABLED() *TextareaElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *TextareaElement) RemoveDISABLED() *TextareaElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *TextareaElement) SetDISABLED(b bool) *TextareaElement {
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
func (e *TextareaElement) DRAGGABLE(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *TextareaElement) RemoveDRAGGABLE(v string) *TextareaElement {
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
func (e *TextareaElement) ENTERKEYHINT(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *TextareaElement) RemoveENTERKEYHINT(v string) *TextareaElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//  * id
func (e *TextareaElement) FORM(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["form"] = v
    return e
}

func (e *TextareaElement) RemoveFORM(v string) *TextareaElement {
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
func (e *TextareaElement) HIDDEN(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *TextareaElement) RemoveHIDDEN(v string) *TextareaElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *TextareaElement) ID(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *TextareaElement) RemoveID(v string) *TextareaElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaElement) INERT() *TextareaElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *TextareaElement) RemoveINERT() *TextareaElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *TextareaElement) SetINERT(b bool) *TextareaElement {
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
func (e *TextareaElement) INPUTMODE(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *TextareaElement) RemoveINPUTMODE(v string) *TextareaElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *TextareaElement) IS(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *TextareaElement) RemoveIS(v string) *TextareaElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *TextareaElement) ITEMID(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *TextareaElement) RemoveITEMID(v string) *TextareaElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *TextareaElement) ITEMPROP(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *TextareaElement) RemoveITEMPROP(v string) *TextareaElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *TextareaElement) ITEMREF(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *TextareaElement) RemoveITEMREF(v string) *TextareaElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaElement) ITEMSCOPE() *TextareaElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *TextareaElement) RemoveITEMSCOPE() *TextareaElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *TextareaElement) SetITEMSCOPE(b bool) *TextareaElement {
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
func (e *TextareaElement) ITEMTYPE(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *TextareaElement) RemoveITEMTYPE(v string) *TextareaElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TextareaElement) LANG(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *TextareaElement) RemoveLANG(v string) *TextareaElement {
    delete(e.StringAttributes, "lang")
    return e
}
// MAXLENGTH sets the "maxlength" attribute.
// Maximum length of value
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *TextareaElement) MAXLENGTH(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["maxlength"] = v
    return e
}

func (e *TextareaElement) RemoveMAXLENGTH(v string) *TextareaElement {
    delete(e.StringAttributes, "maxlength")
    return e
}
// MINLENGTH sets the "minlength" attribute.
// Minimum length of value
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *TextareaElement) MINLENGTH(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["minlength"] = v
    return e
}

func (e *TextareaElement) RemoveMINLENGTH(v string) *TextareaElement {
    delete(e.StringAttributes, "minlength")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *TextareaElement) NAME(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *TextareaElement) RemoveNAME(v string) *TextareaElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *TextareaElement) NONCE(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *TextareaElement) RemoveNONCE(v string) *TextareaElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// PLACEHOLDER sets the "placeholder" attribute.
// User-visible label to be placed within the form control
// Values values are constrained to:
//  * text
func (e *TextareaElement) PLACEHOLDER(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["placeholder"] = v
    return e
}

func (e *TextareaElement) RemovePLACEHOLDER(v string) *TextareaElement {
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
func (e *TextareaElement) POPOVER(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *TextareaElement) RemovePOPOVER(v string) *TextareaElement {
    delete(e.StringAttributes, "popover")
    return e
}
// READONLY sets the "readonly" attribute.
// Affects willValidate, plus any behavior added by the custom element author
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaElement) READONLY() *TextareaElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["readonly"] = struct{}{}
    return e
}

func (e *TextareaElement) RemoveREADONLY() *TextareaElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "readonly")
    return e
}

func (e *TextareaElement) SetREADONLY(b bool) *TextareaElement {
    if b {
        return e.READONLY()
    }
    return e.RemoveREADONLY()
}
// REQUIRED sets the "required" attribute.
// Whether the control is required for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *TextareaElement) REQUIRED() *TextareaElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["required"] = struct{}{}
    return e
}

func (e *TextareaElement) RemoveREQUIRED() *TextareaElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "required")
    return e
}

func (e *TextareaElement) SetREQUIRED(b bool) *TextareaElement {
    if b {
        return e.REQUIRED()
    }
    return e.RemoveREQUIRED()
}
// ROWS sets the "rows" attribute.
// Number of lines to show
// Values values are constrained to:
//  * valid_non_negative_integer
func (e *TextareaElement) ROWS(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["rows"] = v
    return e
}

func (e *TextareaElement) RemoveROWS(v string) *TextareaElement {
    delete(e.StringAttributes, "rows")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *TextareaElement) SLOT(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *TextareaElement) RemoveSLOT(v string) *TextareaElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *TextareaElement) SPELLCHECK(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *TextareaElement) RemoveSPELLCHECK(v string) *TextareaElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TextareaElement) STYLE(k,v string) *TextareaElement {
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

func (e *TextareaElement) RemoveSTYLE(k string) *TextareaElement {
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
func (e *TextareaElement) TABINDEX(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *TextareaElement) RemoveTABINDEX(v string) *TextareaElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *TextareaElement) TITLE(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *TextareaElement) RemoveTITLE(v string) *TextareaElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *TextareaElement) TRANSLATE(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *TextareaElement) RemoveTRANSLATE(v string) *TextareaElement {
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
func (e *TextareaElement) WRAP(v string) *TextareaElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["wrap"] = v
    return e
}

func (e *TextareaElement) RemoveWRAP(v string) *TextareaElement {
    delete(e.StringAttributes, "wrap")
    return e
}
