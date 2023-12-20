package  html

import (
    "fmt"
)

type TemplateHTMLElement struct {
    *Element
}

func TEMPLATE(children ...ElementBuilder) *TemplateHTMLElement {
    return &TemplateHTMLElement{
        Element: &Element{
            Tag: "template",
            IsSelfClosing: false,
            Descendants: children,
        },
    }
}

func (e *TemplateHTMLElement) Children(children ...ElementBuilder) *TemplateHTMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *TemplateHTMLElement) IfChildren(condition bool, children ...ElementBuilder) *TemplateHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *TemplateHTMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementBuilder) *TemplateHTMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *TemplateHTMLElement) Text(text string) *TemplateHTMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *TemplateHTMLElement) TextF(format string, args ...any) *TemplateHTMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *TemplateHTMLElement) Raw(text string) *TemplateHTMLElement {
    e.Descendants = append(e.Descendants, Raw(text))
    return e
}

func (e *TemplateHTMLElement) RawF(format string, args ...any) *TemplateHTMLElement {
    return e.Raw(fmt.Sprintf(format, args...))
}

func (e *TemplateHTMLElement) CustomData(key, value string) *TemplateHTMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = map[string]string{}
    }
	e.CustomDataAttributes[key] = value
	return e
}

func (e *TemplateHTMLElement) CustomDataRemove(key string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) ACCESSKEY(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["accesskey"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveACCESSKEY(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) AUTOCAPITALIZE(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["autocapitalize"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveAUTOCAPITALIZE(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "autocapitalize")
    return e
}
// AUTOFOCUS sets the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Values values are constrained to:
//  * boolean_attribute
func (e *TemplateHTMLElement) AUTOFOCUS() *TemplateHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *TemplateHTMLElement) RemoveAUTOFOCUS() *TemplateHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "autofocus")
    return e
}

func (e *TemplateHTMLElement) SetAUTOFOCUS(b bool) *TemplateHTMLElement {
    if b {
        return e.AUTOFOCUS()
    }
    return e.RemoveAUTOFOCUS()
}
// CLASS sets the "class" attribute.
// Classes to which the element belongs
// Values values are constrained to:
//  * set_of_space_separated_tokens
func(e *TemplateHTMLElement) CLASS(v string) *TemplateHTMLElement {
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

func (e *TemplateHTMLElement) SetCLASS(v string) *TemplateHTMLElement {
    kv := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = kv
    kv.Add(v)
    return e
}

func (e *TemplateHTMLElement) RemoveCLASS(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) CONTENTEDITABLE(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["contenteditable"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveCONTENTEDITABLE(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) DIR(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["dir"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveDIR(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "dir")
    return e
}
// DRAGGABLE sets the "draggable" attribute.
// Whether the element is draggable
// Values values are constrained to:
//  * true
//  * false
func (e *TemplateHTMLElement) DRAGGABLE(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["draggable"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveDRAGGABLE(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) ENTERKEYHINT(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["enterkeyhint"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveENTERKEYHINT(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "enterkeyhint")
    return e
}
// HIDDEN sets the "hidden" attribute.
// Whether the element is relevant
// Values values are constrained to:
//  * until_found
//  * until_found
//  * hidden
//  * hidden
func (e *TemplateHTMLElement) HIDDEN(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["hidden"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveHIDDEN(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "hidden")
    return e
}
// ID sets the "id" attribute.
// The element's ID
// Values values are constrained to:
//  * text
func (e *TemplateHTMLElement) ID(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["id"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveID(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "id")
    return e
}
// INERT sets the "inert" attribute.
// Whether the element is inert.
// Values values are constrained to:
//  * boolean_attribute
func (e *TemplateHTMLElement) INERT() *TemplateHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *TemplateHTMLElement) RemoveINERT() *TemplateHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "inert")
    return e
}

func (e *TemplateHTMLElement) SetINERT(b bool) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) INPUTMODE(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["inputmode"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveINPUTMODE(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "inputmode")
    return e
}
// IS sets the "is" attribute.
// Creates a customized built-in element
// Values values are constrained to:
//  * valid_custom_element_name
//  * customized_built_in_element
func (e *TemplateHTMLElement) IS(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["is"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveIS(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "is")
    return e
}
// ITEMID sets the "itemid" attribute.
// Global identifier for a microdata item
// Values values are constrained to:
//  * valid_url_potentially_surrounded_by_spaces
func (e *TemplateHTMLElement) ITEMID(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemid"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveITEMID(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "itemid")
    return e
}
// ITEMPROP sets the "itemprop" attribute.
// Property names of a microdata item
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
//  * valid_absolute_ur_ls
//  * defined_property_names
func (e *TemplateHTMLElement) ITEMPROP(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemprop"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveITEMPROP(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "itemprop")
    return e
}
// ITEMREF sets the "itemref" attribute.
// Referenced elements
// Values values are constrained to:
//  * unordered_set_of_unique_space_separated_tokens
func (e *TemplateHTMLElement) ITEMREF(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemref"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveITEMREF(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "itemref")
    return e
}
// ITEMSCOPE sets the "itemscope" attribute.
// Introduces a microdata item
// Values values are constrained to:
//  * boolean_attribute
func (e *TemplateHTMLElement) ITEMSCOPE() *TemplateHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *TemplateHTMLElement) RemoveITEMSCOPE() *TemplateHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "itemscope")
    return e
}

func (e *TemplateHTMLElement) SetITEMSCOPE(b bool) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) ITEMTYPE(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["itemtype"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveITEMTYPE(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "itemtype")
    return e
}
// LANG sets the "lang" attribute.
// Language of the element
// Values values are constrained to:
func (e *TemplateHTMLElement) LANG(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["lang"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveLANG(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "lang")
    return e
}
// NONCE sets the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Values values are constrained to:
//  * text
func (e *TemplateHTMLElement) NONCE(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["nonce"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveNONCE(v string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) POPOVER(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["popover"] = v
    return e
}

func (e *TemplateHTMLElement) RemovePOPOVER(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "popover")
    return e
}
// SHADOWROOTDELEGATESFOCUS sets the "shadowrootdelegatesfocus" attribute.
// Sets delegates focus on a declarative shadow root
// Values values are constrained to:
//  * boolean_attribute
func (e *TemplateHTMLElement) SHADOWROOTDELEGATESFOCUS() *TemplateHTMLElement {
    if e.BoolAttributes == nil {
        e.BoolAttributes = map[string]struct{}{}
    }
    e.BoolAttributes["shadowrootdelegatesfocus"] = struct{}{}
    return e
}

func (e *TemplateHTMLElement) RemoveSHADOWROOTDELEGATESFOCUS() *TemplateHTMLElement {
    if e.BoolAttributes == nil {
        return e
    }
    delete(e.BoolAttributes, "shadowrootdelegatesfocus")
    return e
}

func (e *TemplateHTMLElement) SetSHADOWROOTDELEGATESFOCUS(b bool) *TemplateHTMLElement {
    if b {
        return e.SHADOWROOTDELEGATESFOCUS()
    }
    return e.RemoveSHADOWROOTDELEGATESFOCUS()
}
// SHADOWROOTMODE sets the "shadowrootmode" attribute.
// Enables streaming declarative shadow roots
// Values values are constrained to:
//  * open
//  * closed
func (e *TemplateHTMLElement) SHADOWROOTMODE(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["shadowrootmode"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveSHADOWROOTMODE(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "shadowrootmode")
    return e
}
// SLOT sets the "slot" attribute.
// The element's desired slot
// Values values are constrained to:
//  * text
func (e *TemplateHTMLElement) SLOT(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["slot"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveSLOT(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "slot")
    return e
}
// SPELLCHECK sets the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Values values are constrained to:
//  * true
//  * false
func (e *TemplateHTMLElement) SPELLCHECK(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["spellcheck"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveSPELLCHECK(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "spellcheck")
    return e
}
// STYLE sets the "style" attribute.
// Presentational and formatting instructions
// Values values are constrained to:
func (e *TemplateHTMLElement) STYLE(k,v string) *TemplateHTMLElement {
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

func (e *TemplateHTMLElement) RemoveSTYLE(k string) *TemplateHTMLElement {
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
func (e *TemplateHTMLElement) TABINDEX(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["tabindex"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveTABINDEX(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "tabindex")
    return e
}
// TITLE sets the "title" attribute.
// CSS style sheet set name
// Values values are constrained to:
//  * text
func (e *TemplateHTMLElement) TITLE(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["title"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveTITLE(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "title")
    return e
}
// TRANSLATE sets the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Values values are constrained to:
//  * yes
//  * no
func (e *TemplateHTMLElement) TRANSLATE(v string) *TemplateHTMLElement {
    if e.StringAttributes == nil {
        e.StringAttributes = map[string]string{}
    }
    e.StringAttributes["translate"] = v
    return e
}

func (e *TemplateHTMLElement) RemoveTRANSLATE(v string) *TemplateHTMLElement {
    delete(e.StringAttributes, "translate")
    return e
}
