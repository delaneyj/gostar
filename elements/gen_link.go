
/* cSpell:disable */

package  elements

type LinkElement struct {
    *Element
}

func LINK() *LinkElement {
    return &LinkElement{
        Element: &Element{
            Tag: "link",
            IsSelfClosing: true,
        },
    }
}

func (e *LinkElement) AddChildren(children ...*Element) *LinkElement {
    e.Children = append(e.Children, children...)
	return e
}


// Accesskey is the "accesskey" attribute.
// Keyboard shortcut to activate or focus element
// Valid values are constrained to the following:
//   * ordered_set_of_unique_space_separated_tokens
//   * identical_to
func (e *LinkElement) Accesskey(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["accesskey"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) AccesskeyAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["accesskey"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["accesskey"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) AccesskeyRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["accesskey"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// As is the "as" attribute.
// Potential destination for a preload request (for rel=&quot;preload&quot; and rel=&quot;modulepreload&quot;)
// Valid values are constrained to the following:
//   * potential_destination
//   * rel
//   * rel
//   * preload
//   * preload
//   * script_like_destination
//   * rel
//   * rel
//   * modulepreload
//   * modulepreload
func (e *LinkElement) As(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["as"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) AsAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["as"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["as"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) AsRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["as"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Autocapitalize is the "autocapitalize" attribute.
// Recommended autocapitalization behavior (for supported input methods)
// Valid values are constrained to the following:
//   * on
//   * on
//   * off
//   * off
//   * none
//   * none
//   * sentences
//   * sentences
//   * words
//   * words
//   * characters
//   * characters
func (e *LinkElement) Autocapitalize(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) AutocapitalizeAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocapitalize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["autocapitalize"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) AutocapitalizeRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["autocapitalize"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Autofocus is the "autofocus" attribute.
// Automatically focus the element when the page is loaded
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *LinkElement) Autofocus() *LinkElement {
    e.BoolAttributes["autofocus"] = struct{}{}
    return e
}

func (e *LinkElement) AutofocusRemove() *LinkElement {
    delete(e.BoolAttributes,"autofocus")
    return e
}

func (e *LinkElement) AutofocusSet(v bool) *LinkElement {
    if v {
        e.Autofocus()
    } else {
        e.AutofocusRemove()
    }
    return e
}


// Blocking is the "blocking" attribute.
// Whether the element is potentially render-blocking
// Valid values are constrained to the following:
//   * unordered_set_of_unique_space_separated_tokens
func (e *LinkElement) Blocking(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["blocking"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) BlockingAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["blocking"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["blocking"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) BlockingRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["blocking"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Class is the "class" attribute.
// Classes to which the element belongs
// Valid values are constrained to the following:
//   * set_of_space_separated_tokens
func (e *LinkElement) Class(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["class"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ClassAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["class"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["class"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ClassRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["class"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Color is the "color" attribute.
// Color to use when customizing a site&#39;s icon (for rel=&quot;mask-icon&quot;)
// Valid values are constrained to the following:
//   * &lt;color&gt;
func (e *LinkElement) Color(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["color"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ColorAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["color"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["color"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ColorRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["color"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Contenteditable is the "contenteditable" attribute.
// Whether the element is editable
// Valid values are constrained to the following:
//   * true
//   * plaintext_only
//   * false
func (e *LinkElement) Contenteditable(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ContenteditableAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["contenteditable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["contenteditable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ContenteditableRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["contenteditable"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Crossorigin is the "crossorigin" attribute.
// How the element handles crossorigin requests
// Valid values are constrained to the following:
//   * anonymous
//   * anonymous
//   * use_credentials
//   * use_credentials
func (e *LinkElement) Crossorigin(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["crossorigin"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) CrossoriginAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["crossorigin"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["crossorigin"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) CrossoriginRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["crossorigin"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Dir is the "dir" attribute.
// The text directionality of the element
// Valid values are constrained to the following:
//   * ltr
//   * ltr
//   * rtl
//   * rtl
func (e *LinkElement) Dir(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["dir"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) DirAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dir"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["dir"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) DirRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["dir"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Disabled is the "disabled" attribute.
// Whether the link is disabled
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *LinkElement) Disabled() *LinkElement {
    e.BoolAttributes["disabled"] = struct{}{}
    return e
}

func (e *LinkElement) DisabledRemove() *LinkElement {
    delete(e.BoolAttributes,"disabled")
    return e
}

func (e *LinkElement) DisabledSet(v bool) *LinkElement {
    if v {
        e.Disabled()
    } else {
        e.DisabledRemove()
    }
    return e
}


// Draggable is the "draggable" attribute.
// Whether the element is draggable
// Valid values are constrained to the following:
//   * true
//   * false
func (e *LinkElement) Draggable(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["draggable"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) DraggableAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["draggable"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["draggable"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) DraggableRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["draggable"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Enterkeyhint is the "enterkeyhint" attribute.
// Hint for selecting an enter key action
// Valid values are constrained to the following:
//   * enter
//   * enter
//   * done
//   * done
//   * go
//   * go
//   * next
//   * next
//   * previous
//   * previous
//   * search
//   * search
//   * send
//   * send
func (e *LinkElement) Enterkeyhint(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) EnterkeyhintAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["enterkeyhint"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["enterkeyhint"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) EnterkeyhintRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["enterkeyhint"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Fetchpriority is the "fetchpriority" attribute.
// Sets the priority for fetches initiated by the element
// Valid values are constrained to the following:
//   * auto
//   * auto
//   * high
//   * high
//   * low
//   * low
func (e *LinkElement) Fetchpriority(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["fetchpriority"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) FetchpriorityAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["fetchpriority"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["fetchpriority"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) FetchpriorityRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["fetchpriority"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Hidden is the "hidden" attribute.
// Whether the element is relevant
// Valid values are constrained to the following:
//   * until_found
//   * until_found
//   * hidden
//   * hidden
func (e *LinkElement) Hidden(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["hidden"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) HiddenAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["hidden"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["hidden"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) HiddenRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["hidden"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Href is the "href" attribute.
// Document base URL
// Valid values are constrained to the following:
//   * valid_url_potentially_surrounded_by_spaces
func (e *LinkElement) Href(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["href"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) HrefAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["href"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["href"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) HrefRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["href"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Hreflang is the "hreflang" attribute.
// Language of the linked resource
// Valid values are constrained to the following:
func (e *LinkElement) Hreflang(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["hreflang"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) HreflangAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["hreflang"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["hreflang"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) HreflangRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["hreflang"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Id is the "id" attribute.
// The element&#39;s ID
// Valid values are constrained to the following:
//   * text
func (e *LinkElement) Id(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["id"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) IdAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["id"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["id"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) IdRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["id"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Imagesizes is the "imagesizes" attribute.
// Image sizes for different page layouts (for rel=&quot;preload&quot;)
// Valid values are constrained to the following:
//   * valid_source_size_list
func (e *LinkElement) Imagesizes(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["imagesizes"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ImagesizesAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["imagesizes"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["imagesizes"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ImagesizesRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["imagesizes"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Imagesrcset is the "imagesrcset" attribute.
// Images to use in different situations, e.g., high-resolution displays, small monitors, etc. (for rel=&quot;preload&quot;)
// Valid values are constrained to the following:
//   * image_candidate_strings
func (e *LinkElement) Imagesrcset(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["imagesrcset"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ImagesrcsetAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["imagesrcset"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["imagesrcset"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ImagesrcsetRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["imagesrcset"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Inert is the "inert" attribute.
// Whether the element is inert.
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *LinkElement) Inert() *LinkElement {
    e.BoolAttributes["inert"] = struct{}{}
    return e
}

func (e *LinkElement) InertRemove() *LinkElement {
    delete(e.BoolAttributes,"inert")
    return e
}

func (e *LinkElement) InertSet(v bool) *LinkElement {
    if v {
        e.Inert()
    } else {
        e.InertRemove()
    }
    return e
}


// Inputmode is the "inputmode" attribute.
// Hint for selecting an input modality
// Valid values are constrained to the following:
//   * none
//   * none
//   * text
//   * text
//   * tel
//   * tel
//   * email
//   * email
//   * url
//   * url
//   * numeric
//   * numeric
//   * decimal
//   * decimal
//   * search
//   * search
func (e *LinkElement) Inputmode(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["inputmode"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) InputmodeAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["inputmode"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["inputmode"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) InputmodeRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["inputmode"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Integrity is the "integrity" attribute.
// Integrity metadata used in Subresource Integrity checks [SRI]
// Valid values are constrained to the following:
//   * text
func (e *LinkElement) Integrity(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["integrity"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) IntegrityAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["integrity"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["integrity"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) IntegrityRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["integrity"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Is is the "is" attribute.
// Creates a customized built-in element
// Valid values are constrained to the following:
//   * valid_custom_element_name
//   * customized_built_in_element
func (e *LinkElement) Is(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["is"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) IsAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["is"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["is"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) IsRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["is"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Itemid is the "itemid" attribute.
// Global identifier for a microdata item
// Valid values are constrained to the following:
//   * valid_url_potentially_surrounded_by_spaces
func (e *LinkElement) Itemid(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemid"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ItemidAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemid"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ItemidRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemid"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Itemprop is the "itemprop" attribute.
// Property names of a microdata item
// Valid values are constrained to the following:
//   * unordered_set_of_unique_space_separated_tokens
//   * valid_absolute_ur_ls
//   * defined_property_names
func (e *LinkElement) Itemprop(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemprop"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ItempropAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemprop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemprop"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ItempropRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemprop"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Itemref is the "itemref" attribute.
// Referenced elements
// Valid values are constrained to the following:
//   * unordered_set_of_unique_space_separated_tokens
func (e *LinkElement) Itemref(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemref"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ItemrefAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemref"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemref"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ItemrefRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemref"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Itemscope is the "itemscope" attribute.
// Introduces a microdata item
// Valid values are constrained to the following:
//   * boolean_attribute
func (e *LinkElement) Itemscope() *LinkElement {
    e.BoolAttributes["itemscope"] = struct{}{}
    return e
}

func (e *LinkElement) ItemscopeRemove() *LinkElement {
    delete(e.BoolAttributes,"itemscope")
    return e
}

func (e *LinkElement) ItemscopeSet(v bool) *LinkElement {
    if v {
        e.Itemscope()
    } else {
        e.ItemscopeRemove()
    }
    return e
}


// Itemtype is the "itemtype" attribute.
// Item types of a microdata item
// Valid values are constrained to the following:
//   * unordered_set_of_unique_space_separated_tokens
//   * valid_absolute_ur_ls
func (e *LinkElement) Itemtype(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["itemtype"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ItemtypeAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["itemtype"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ItemtypeRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["itemtype"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Lang is the "lang" attribute.
// Language of the element
// Valid values are constrained to the following:
func (e *LinkElement) Lang(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["lang"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) LangAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["lang"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["lang"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) LangRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["lang"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Media is the "media" attribute.
// Applicable media
// Valid values are constrained to the following:
//   * valid_media_query_list
func (e *LinkElement) Media(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["media"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) MediaAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["media"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["media"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) MediaRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["media"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Nonce is the "nonce" attribute.
// Cryptographic nonce used in Content Security Policy checks [CSP]
// Valid values are constrained to the following:
//   * text
func (e *LinkElement) Nonce(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["nonce"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) NonceAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["nonce"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["nonce"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) NonceRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["nonce"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Popover is the "popover" attribute.
// Makes the element a popover element
// Valid values are constrained to the following:
//   * auto
//   * auto
//   * manual
//   * manual
func (e *LinkElement) Popover(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["popover"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) PopoverAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["popover"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) PopoverRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["popover"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Referrerpolicy is the "referrerpolicy" attribute.
// Referrer policy for fetches initiated by the element
// Valid values are constrained to the following:
//   * referrer_policy
func (e *LinkElement) Referrerpolicy(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["referrerpolicy"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ReferrerpolicyAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["referrerpolicy"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["referrerpolicy"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) ReferrerpolicyRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["referrerpolicy"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Rel is the "rel" attribute.
// Relationship between the document containing the hyperlink and the destination resource
// Valid values are constrained to the following:
//   * unordered_set_of_unique_space_separated_tokens
func (e *LinkElement) Rel(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["rel"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) RelAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["rel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["rel"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) RelRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["rel"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Sizes is the "sizes" attribute.
// Image sizes for different page layouts
// Valid values are constrained to the following:
//   * valid_source_size_list
func (e *LinkElement) Sizes(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["sizes"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) SizesAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["sizes"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["sizes"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) SizesRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["sizes"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Slot is the "slot" attribute.
// The element&#39;s desired slot
// Valid values are constrained to the following:
//   * text
func (e *LinkElement) Slot(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["slot"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) SlotAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["slot"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["slot"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) SlotRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["slot"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Spellcheck is the "spellcheck" attribute.
// Whether the element is to have its spelling and grammar checked
// Valid values are constrained to the following:
//   * true
//   * false
func (e *LinkElement) Spellcheck(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) SpellcheckAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["spellcheck"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) SpellcheckRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["spellcheck"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Style is the "style" attribute.
// Presentational and formatting instructions
// Valid values are constrained to the following:
func (e *LinkElement) Style(k,v string) *LinkElement {
    htmlStringer,ok := e.DelimitedKVAttributes["style"]
    if !ok {
        htmlStringer = NewEqualSemicolonDelimitedKVString()
        e.DelimitedKVAttributes["style"] = htmlStringer
    }
    htmlStringer.Add(k,v)
    return e
}

func (e *LinkElement) StyleRemove(k string) *LinkElement {
    htmlStringer,ok := e.DelimitedKVAttributes["style"]
    if ok {
        htmlStringer.Remove(k)
    }
    return e
}


// Tabindex is the "tabindex" attribute.
// Whether the element is focusable and sequentially focusable, and       the relative order of the element for the purposes of sequential focus navigation
// Valid values are constrained to the following:
//   * valid_integer
func (e *LinkElement) Tabindex(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["tabindex"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) TabindexAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["tabindex"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["tabindex"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) TabindexRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["tabindex"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Title is the "title" attribute.
// CSS style sheet set name
// Valid values are constrained to the following:
//   * text
func (e *LinkElement) Title(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["title"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) TitleAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["title"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["title"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) TitleRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["title"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Translate is the "translate" attribute.
// Whether the element is to be translated when the page is localized
// Valid values are constrained to the following:
//   * yes
//   * no
func (e *LinkElement) Translate(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["translate"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) TranslateAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["translate"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) TranslateRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["translate"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}


// Type is the "type" attribute.
// Type of script
// Valid values are constrained to the following:
//   * module
//   * valid_mime_type_string
//   * java_script_mime_type_essence_match
func (e *LinkElement) Type(v string) *LinkElement {
    htmlStringer := NewSpaceDelimitedString()
    e.DelimitedStringAttributes["type"] = htmlStringer
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) TypeAdd(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["type"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["type"] = htmlStringer
    }
    htmlStringer.Add(v)
    return e
}

func (e *LinkElement) TypeRemove(v string) *LinkElement {
    htmlStringer,ok := e.DelimitedStringAttributes["type"]
    if ok {
        htmlStringer.Remove(v)
    }
    return e
}




func (e *LinkElement) OnAuxclick(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onauxclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnAuxclickRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onauxclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnBeforematch(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforematch"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnBeforematchRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforematch"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnBeforetoggle(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onbeforetoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnBeforetoggleRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onbeforetoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnBlur(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onblur"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnBlurRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onblur"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnCancel(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncancel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnCancelRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncancel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnCanplay(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnCanplayRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnCanplaythrough(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncanplaythrough"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnCanplaythroughRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncanplaythrough"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnChange(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnChangeRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnClick(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnClickRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnClose(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onclose"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnCloseRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onclose"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnContextlost(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextlost"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnContextlostRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextlost"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnContextmenu(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextmenu"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnContextmenuRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextmenu"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnContextrestored(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncontextrestored"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnContextrestoredRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncontextrestored"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnCopy(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncopy"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnCopyRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncopy"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnCuechange(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncuechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnCuechangeRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncuechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnCut(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oncut"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnCutRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oncut"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnDblclick(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondblclick"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnDblclickRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondblclick"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnDrag(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrag"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnDragRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrag"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnDragend(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnDragendRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnDragenter(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnDragenterRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnDragleave(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnDragleaveRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnDragover(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnDragoverRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnDragstart(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondragstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnDragstartRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondragstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnDrop(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondrop"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnDropRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondrop"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnDurationchange(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ondurationchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnDurationchangeRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ondurationchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnEmptied(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onemptied"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnEmptiedRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onemptied"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnEnded(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onended"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnEndedRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onended"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnError(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onerror"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnErrorRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onerror"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnFocus(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onfocus"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnFocusRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onfocus"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnFormdata(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onformdata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnFormdataRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onformdata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnInput(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninput"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnInputRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninput"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnInvalid(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["oninvalid"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnInvalidRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["oninvalid"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnKeydown(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeydown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnKeydownRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeydown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnKeypress(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeypress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnKeypressRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeypress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnKeyup(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onkeyup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnKeyupRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onkeyup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnLoad(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onload"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnLoadRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onload"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnLoadeddata(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadeddata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnLoadeddataRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadeddata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnLoadedmetadata(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadedmetadata"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnLoadedmetadataRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadedmetadata"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnLoadstart(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onloadstart"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnLoadstartRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onloadstart"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnMousedown(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousedown"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnMousedownRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousedown"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnMouseenter(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseenter"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnMouseenterRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseenter"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnMouseleave(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseleave"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnMouseleaveRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseleave"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnMousemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmousemove"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnMousemoveRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmousemove"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnMouseout(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseout"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnMouseoutRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseout"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnMouseover(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseover"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnMouseoverRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseover"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnMouseup(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onmouseup"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnMouseupRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onmouseup"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnPaste(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpaste"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnPasteRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpaste"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnPause(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onpause"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnPauseRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onpause"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnPlay(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplay"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnPlayRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplay"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnPlaying(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onplaying"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnPlayingRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onplaying"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnProgress(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onprogress"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnProgressRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onprogress"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnRatechange(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onratechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnRatechangeRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onratechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnReset(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onreset"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnResetRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onreset"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnResize(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onresize"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnResizeRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onresize"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnScroll(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscroll"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnScrollRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscroll"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnScrollend(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onscrollend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnScrollendRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onscrollend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnSecuritypolicyviolation(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsecuritypolicyviolation"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnSecuritypolicyviolationRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsecuritypolicyviolation"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnSeeked(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeked"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnSeekedRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeked"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnSeeking(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onseeking"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnSeekingRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onseeking"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnSelect(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onselect"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnSelectRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onselect"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnSlotchange(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onslotchange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnSlotchangeRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onslotchange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnStalled(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onstalled"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnStalledRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onstalled"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnSubmit(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsubmit"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnSubmitRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsubmit"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnSuspend(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onsuspend"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnSuspendRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onsuspend"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnTimeupdate(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontimeupdate"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnTimeupdateRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontimeupdate"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnToggle(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["ontoggle"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnToggleRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["ontoggle"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnVolumechange(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onvolumechange"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnVolumechangeRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onvolumechange"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnWaiting(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwaiting"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnWaitingRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwaiting"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}



func (e *LinkElement) OnWheel(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if !ok {
        htmlStringer = NewSpaceDelimitedString()
        e.DelimitedStringAttributes["onwheel"] = htmlStringer
    }
    htmlStringer.Add(evt)
    return e
}

func (e *LinkElement) OnWheelRemove(evt string) (*LinkElement) {
    htmlStringer,ok := e.DelimitedStringAttributes["onwheel"]
    if ok {
        htmlStringer.Remove(evt)
    }
    return e
}


