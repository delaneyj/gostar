package  elements

import (
    "fmt"
	"github.com/delaneyj/gostar"
)

type ButtonElement struct {
    *gostar.Element
}

func BUTTON(children ...fmt.Stringer) *ButtonElement {
    return &ButtonElement{
        Element: &gostar.Element{
            Tag: "button",
            IsSelfClosing: false,
            Children: children,
        },
    }
}

func (e *ButtonElement) AddChildren(children ...fmt.Stringer) *ButtonElement {
    e.Children = append(e.Children, children...)
    return e
}

func (e *ButtonElement) TEXT(text string) *ButtonElement {
    e.Children = append(e.Children, gostar.TEXT(text))
    return e
}

func (e *ButtonElement) RAW(text string) *ButtonElement {
    e.Children = append(e.Children, gostar.RAW(text))
    return e
}

func (e *ButtonElement) CustomData(key, value string) *ButtonElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ButtonElement) CustomDataRemove(key string) *ButtonElement {
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
func (e *ButtonElement) ACCESSKEY(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *ButtonElement) RemoveACCESSKEY(v string) *ButtonElement {
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
func (e *ButtonElement) AUTOCAPITALIZE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *ButtonElement) RemoveAUTOCAPITALIZE(v string) *ButtonElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonElement) AUTOFOCUS() *ButtonElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *ButtonElement) RemoveAUTOFOCUS() *ButtonElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *ButtonElement) SetAUTOFOCUS(b bool) *ButtonElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *ButtonElement) CLASS(v string) *ButtonElement {
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

func (e *ButtonElement) SetCLASS(v string) *ButtonElement {
    kv := gostar.NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *ButtonElement) RemoveCLASS(v string) *ButtonElement {
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
func (e *ButtonElement) CONTENTEDITABLE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *ButtonElement) RemoveCONTENTEDITABLE(v string) *ButtonElement {
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
func (e *ButtonElement) DIR(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *ButtonElement) RemoveDIR(v string) *ButtonElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonElement) DISABLED() *ButtonElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *ButtonElement) RemoveDISABLED() *ButtonElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *ButtonElement) SetDISABLED(b bool) *ButtonElement {
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
func (e *ButtonElement) DRAGGABLE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *ButtonElement) RemoveDRAGGABLE(v string) *ButtonElement {
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
func (e *ButtonElement) ENTERKEYHINT(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *ButtonElement) RemoveENTERKEYHINT(v string) *ButtonElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//  * id
func (e *ButtonElement) FORM(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["form"] = v
    return e
}

func (e *ButtonElement) RemoveFORM(v string) *ButtonElement {
    delete(e.StringAttributes, "form")
    return e
}
// FORMACTION sets the "formaction" attribute.
// URL to use for form submission
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ButtonElement) FORMACTION(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formaction"] = v
    return e
}

func (e *ButtonElement) RemoveFORMACTION(v string) *ButtonElement {
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
func (e *ButtonElement) FORMENCTYPE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formenctype"] = v
    return e
}

func (e *ButtonElement) RemoveFORMENCTYPE(v string) *ButtonElement {
    delete(e.StringAttributes, "formenctype")
    return e
}
// FORMMETHOD sets the "formmethod" attribute.
// Variant to use for form submission
// Values values are constrained to:
//  * get
//  * post
//  * dialog
func (e *ButtonElement) FORMMETHOD(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formmethod"] = v
    return e
}

func (e *ButtonElement) RemoveFORMMETHOD(v string) *ButtonElement {
    delete(e.StringAttributes, "formmethod")
    return e
}
// FORMNOVALIDATE sets the "formnovalidate" attribute.
// Bypass form control validation for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonElement) FORMNOVALIDATE() *ButtonElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["formnovalidate"] = struct{}{}
    return e
}

func (e *ButtonElement) RemoveFORMNOVALIDATE() *ButtonElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "formnovalidate")
    return e
}

func (e *ButtonElement) SetFORMNOVALIDATE(b bool) *ButtonElement {
    if b {
        return e.FORMNOVALIDATE()
    }
    return e.RemoveFORMNOVALIDATE()
}
// FORMTARGET sets the "formtarget" attribute.
// Navigable for form submission
// Values values are constrained to:
//  * valid_navigable_target_name_or_keyword
func (e *ButtonElement) FORMTARGET(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formtarget"] = v
    return e
}

func (e *ButtonElement) RemoveFORMTARGET(v string) *ButtonElement {
    delete(e.StringAttributes, "formtarget")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *ButtonElement) HIDDEN(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *ButtonElement) RemoveHIDDEN(v string) *ButtonElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *ButtonElement) ID(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *ButtonElement) RemoveID(v string) *ButtonElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonElement) INERT() *ButtonElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *ButtonElement) RemoveINERT() *ButtonElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *ButtonElement) SetINERT(b bool) *ButtonElement {
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
func (e *ButtonElement) INPUTMODE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *ButtonElement) RemoveINPUTMODE(v string) *ButtonElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *ButtonElement) IS(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *ButtonElement) RemoveIS(v string) *ButtonElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *ButtonElement) ITEMID(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *ButtonElement) RemoveITEMID(v string) *ButtonElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *ButtonElement) ITEMPROP(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *ButtonElement) RemoveITEMPROP(v string) *ButtonElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ButtonElement) ITEMREF(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *ButtonElement) RemoveITEMREF(v string) *ButtonElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonElement) ITEMSCOPE() *ButtonElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *ButtonElement) RemoveITEMSCOPE() *ButtonElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *ButtonElement) SetITEMSCOPE(b bool) *ButtonElement {
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
func (e *ButtonElement) ITEMTYPE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *ButtonElement) RemoveITEMTYPE(v string) *ButtonElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ButtonElement) LANG(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *ButtonElement) RemoveLANG(v string) *ButtonElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *ButtonElement) NAME(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *ButtonElement) RemoveNAME(v string) *ButtonElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *ButtonElement) NONCE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *ButtonElement) RemoveNONCE(v string) *ButtonElement {
    delete(e.StringAttributes, "nonce")
    return e
}
// POPOVER sets the "popover" attribute.
// Makes the element a popover element
// Values values are constrained to:
//  * auto
//  * auto
//  * manual
//  * manual
func (e *ButtonElement) POPOVER(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *ButtonElement) RemovePOPOVER(v string) *ButtonElement {
    delete(e.StringAttributes, "popover")
    return e
}
// POPOVERTARGET sets the "popovertarget" attribute.
// Targets a popover element to toggle, show, or hide
// Values values are constrained to:
//  * id
func (e *ButtonElement) POPOVERTARGET(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popovertarget"] = v
    return e
}

func (e *ButtonElement) RemovePOPOVERTARGET(v string) *ButtonElement {
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
func (e *ButtonElement) POPOVERTARGETACTION(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popovertargetaction"] = v
    return e
}

func (e *ButtonElement) RemovePOPOVERTARGETACTION(v string) *ButtonElement {
    delete(e.StringAttributes, "popovertargetaction")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *ButtonElement) SLOT(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *ButtonElement) RemoveSLOT(v string) *ButtonElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *ButtonElement) SPELLCHECK(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *ButtonElement) RemoveSPELLCHECK(v string) *ButtonElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ButtonElement) STYLE(k,v string) *ButtonElement {
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

func (e *ButtonElement) RemoveSTYLE(k string) *ButtonElement {
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
func (e *ButtonElement) TABINDEX(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *ButtonElement) RemoveTABINDEX(v string) *ButtonElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *ButtonElement) TITLE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *ButtonElement) RemoveTITLE(v string) *ButtonElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *ButtonElement) TRANSLATE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *ButtonElement) RemoveTRANSLATE(v string) *ButtonElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *ButtonElement) TYPE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *ButtonElement) RemoveTYPE(v string) *ButtonElement {
    delete(e.StringAttributes, "type")
    return e
}
// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//  * valid_floating_point_number
func (e *ButtonElement) VALUE(v string) *ButtonElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["value"] = v
    return e
}

func (e *ButtonElement) RemoveVALUE(v string) *ButtonElement {
    delete(e.StringAttributes, "value")
    return e
}
