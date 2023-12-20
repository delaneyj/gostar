package  html

import (
    "fmt"
)

type ButtonHTMLElement struct {
    *Element
}

func BUTTON(children ...ElementBuilder) *ButtonHTMLElement {
    return &ButtonHTMLElement{
        Element: &Element{
            Tag: "button",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *ButtonHTMLElement) Children(children ...ElementBuilder) *ButtonHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *ButtonHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *ButtonHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *ButtonHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *ButtonHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *ButtonHTMLElement) Text(text string) *ButtonHTMLElement {
    e.Descendants = append(e.Descendants, TEXT(text))
    return e
}

func (e *ButtonHTMLElement) TextF(format string, args ...any) *ButtonHTMLElement {
    TEXT(fmt.Sprintf(format, args...))
    return e
}

func (e *ButtonHTMLElement) Raw(text string) *ButtonHTMLElement {
    e.Descendants = append(e.Descendants, RAW(text))
    return e
}

func (e *ButtonHTMLElement) RawF(format string, args ...any) *ButtonHTMLElement {
    RAW(fmt.Sprintf(format, args...))
    return e
}

func (e *ButtonHTMLElement) CustomData(key, value string) *ButtonHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *ButtonHTMLElement) CustomDataRemove(key string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) ACCESSKEY(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveACCESSKEY(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) AUTOCAPITALIZE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveAUTOCAPITALIZE(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonHTMLElement) AUTOFOCUS() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *ButtonHTMLElement) RemoveAUTOFOCUS() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *ButtonHTMLElement) SetAUTOFOCUS(b bool) *ButtonHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *ButtonHTMLElement) CLASS(v string) *ButtonHTMLElement {
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

func (e *ButtonHTMLElement) SetCLASS(v string) *ButtonHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *ButtonHTMLElement) RemoveCLASS(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) CONTENTEDITABLE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveCONTENTEDITABLE(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) DIR(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveDIR(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DISABLED sets the "disabled" attribute.
// Whether the link is disabled
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonHTMLElement) DISABLED() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *ButtonHTMLElement) RemoveDISABLED() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "disabled")
    return e
}

func (e *ButtonHTMLElement) SetDISABLED(b bool) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) DRAGGABLE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveDRAGGABLE(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) ENTERKEYHINT(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveENTERKEYHINT(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// FORM sets the "form" attribute.
// Associates the element with a form element
// Values values are constrained to:
//  * id
func (e *ButtonHTMLElement) FORM(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["form"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveFORM(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "form")
    return e
}
// FORMACTION sets the "formaction" attribute.
// URL to use for form submission
// Values values are constrained to:
//  * valid_non_empty_url_potentially_surrounded_by_spaces
func (e *ButtonHTMLElement) FORMACTION(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formaction"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveFORMACTION(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) FORMENCTYPE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formenctype"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveFORMENCTYPE(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "formenctype")
    return e
}
// FORMMETHOD sets the "formmethod" attribute.
// Variant to use for form submission
// Values values are constrained to:
//  * get
//  * post
//  * dialog
func (e *ButtonHTMLElement) FORMMETHOD(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formmethod"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveFORMMETHOD(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "formmethod")
    return e
}
// FORMNOVALIDATE sets the "formnovalidate" attribute.
// Bypass form control validation for form submission
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonHTMLElement) FORMNOVALIDATE() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["formnovalidate"] = struct{}{}
    return e
}

func (e *ButtonHTMLElement) RemoveFORMNOVALIDATE() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "formnovalidate")
    return e
}

func (e *ButtonHTMLElement) SetFORMNOVALIDATE(b bool) *ButtonHTMLElement {
    if b {
        return e.FORMNOVALIDATE()
    }
    return e.RemoveFORMNOVALIDATE()
}
// FORMTARGET sets the "formtarget" attribute.
// Navigable for form submission
// Values values are constrained to:
//  * valid_navigable_target_name_or_keyword
func (e *ButtonHTMLElement) FORMTARGET(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["formtarget"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveFORMTARGET(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) HIDDEN(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveHIDDEN(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *ButtonHTMLElement) ID(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveID(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonHTMLElement) INERT() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *ButtonHTMLElement) RemoveINERT() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *ButtonHTMLElement) SetINERT(b bool) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) INPUTMODE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveINPUTMODE(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *ButtonHTMLElement) IS(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveIS(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *ButtonHTMLElement) ITEMID(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveITEMID(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *ButtonHTMLElement) ITEMPROP(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveITEMPROP(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *ButtonHTMLElement) ITEMREF(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveITEMREF(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *ButtonHTMLElement) ITEMSCOPE() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *ButtonHTMLElement) RemoveITEMSCOPE() *ButtonHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *ButtonHTMLElement) SetITEMSCOPE(b bool) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) ITEMTYPE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveITEMTYPE(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *ButtonHTMLElement) LANG(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveLANG(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NAME sets the "name" attribute.
// Name of shadow tree slot
// Values values are constrained to:
//  * text
func (e *ButtonHTMLElement) NAME(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["name"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveNAME(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "name")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *ButtonHTMLElement) NONCE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveNONCE(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) POPOVER(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *ButtonHTMLElement) RemovePOPOVER(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// POPOVERTARGET sets the "popovertarget" attribute.
// Targets a popover element to toggle, show, or hide
// Values values are constrained to:
//  * id
func (e *ButtonHTMLElement) POPOVERTARGET(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popovertarget"] = v
    return e
}

func (e *ButtonHTMLElement) RemovePOPOVERTARGET(v string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) POPOVERTARGETACTION(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popovertargetaction"] = v
    return e
}

func (e *ButtonHTMLElement) RemovePOPOVERTARGETACTION(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "popovertargetaction")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *ButtonHTMLElement) SLOT(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveSLOT(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *ButtonHTMLElement) SPELLCHECK(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveSPELLCHECK(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *ButtonHTMLElement) STYLE(k,v string) *ButtonHTMLElement {
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

func (e *ButtonHTMLElement) RemoveSTYLE(k string) *ButtonHTMLElement {
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
func (e *ButtonHTMLElement) TABINDEX(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveTABINDEX(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *ButtonHTMLElement) TITLE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveTITLE(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *ButtonHTMLElement) TRANSLATE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveTRANSLATE(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
// TYPE sets the "type" attribute.
// Type of script
// Values values are constrained to:
//  * module
//  * valid_mime_type_string
//  * java_script_mime_type_essence_match
func (e *ButtonHTMLElement) TYPE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["type"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveTYPE(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "type")
    return e
}
// VALUE sets the "value" attribute.
// Current value of the element
// Values values are constrained to:
//  * valid_floating_point_number
func (e *ButtonHTMLElement) VALUE(v string) *ButtonHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["value"] = v
    return e
}

func (e *ButtonHTMLElement) RemoveVALUE(v string) *ButtonHTMLElement {
    delete(e.StringAttributes, "value")
    return e
}
