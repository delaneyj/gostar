// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg a is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <a> SVG element creates a hyperlink to other web pages, files, locations in
// the same page, email addresses, or any other URL
// It is very similar to HTML's <a> element
// SVG's <a> element is a container, which means you can create a link around text
// (like in HTML) but also around any shape.
type SVGAElement struct {
	*Element
}

// Create a new SVGAElement element.
// This will create a new element with the tag
// "a" during rendering.
func SVG_A(children ...ElementRenderer) *SVGAElement {
	e := NewElement("a", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGAElement{Element: e}
}

func (e *SVGAElement) Children(children ...ElementRenderer) *SVGAElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGAElement) IfChildren(condition bool, children ...ElementRenderer) *SVGAElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGAElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGAElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGAElement) Attr(name string, value string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGAElement) Attrs(attrs ...string) *SVGAElement {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGAElement) AttrsMap(attrs map[string]string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGAElement) Text(text string) *SVGAElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGAElement) TextF(format string, args ...any) *SVGAElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGAElement) IfText(condition bool, text string) *SVGAElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGAElement) IfTextF(condition bool, format string, args ...any) *SVGAElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGAElement) Escaped(text string) *SVGAElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGAElement) IfEscaped(condition bool, text string) *SVGAElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGAElement) EscapedF(format string, args ...any) *SVGAElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGAElement) IfEscapedF(condition bool, format string, args ...any) *SVGAElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGAElement) CustomData(key, value string) *SVGAElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGAElement) IfCustomData(condition bool, key, value string) *SVGAElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGAElement) CustomDataF(key, format string, args ...any) *SVGAElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGAElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGAElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGAElement) CustomDataRemove(key string) *SVGAElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Indicates that the hyperlink is to be used for downloading a resource
// When used together with the download attribute, the value of the attribute is
// used as the file name of the downloaded file
// There are no restrictions on allowed values, though / and \ will be converted
// to underscores and leading spaces in filenames will be removed.
func (e *SVGAElement) DOWNLOAD(s string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("download", s)
	return e
}

func (e *SVGAElement) DOWNLOADF(format string, args ...any) *SVGAElement {
	return e.DOWNLOAD(fmt.Sprintf(format, args...))
}

func (e *SVGAElement) IfDOWNLOAD(condition bool, s string) *SVGAElement {
	if condition {
		e.DOWNLOAD(s)
	}
	return e
}

func (e *SVGAElement) IfDOWNLOADF(condition bool, format string, args ...any) *SVGAElement {
	if condition {
		e.DOWNLOAD(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute DOWNLOAD from the element.
func (e *SVGAElement) DOWNLOADRemove(s string) *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("download")
	return e
}

func (e *SVGAElement) DOWNLOADRemoveF(format string, args ...any) *SVGAElement {
	return e.DOWNLOADRemove(fmt.Sprintf(format, args...))
}

// The URL of a linked resource.
func (e *SVGAElement) HREF(s string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGAElement) HREFF(format string, args ...any) *SVGAElement {
	return e.HREF(fmt.Sprintf(format, args...))
}

func (e *SVGAElement) IfHREF(condition bool, s string) *SVGAElement {
	if condition {
		e.HREF(s)
	}
	return e
}

func (e *SVGAElement) IfHREFF(condition bool, format string, args ...any) *SVGAElement {
	if condition {
		e.HREF(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREF from the element.
func (e *SVGAElement) HREFRemove(s string) *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("href")
	return e
}

func (e *SVGAElement) HREFRemoveF(format string, args ...any) *SVGAElement {
	return e.HREFRemove(fmt.Sprintf(format, args...))
}

// Specifies the language of the linked resource.
func (e *SVGAElement) HREFLANG(s string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("hreflang", s)
	return e
}

func (e *SVGAElement) HREFLANGF(format string, args ...any) *SVGAElement {
	return e.HREFLANG(fmt.Sprintf(format, args...))
}

func (e *SVGAElement) IfHREFLANG(condition bool, s string) *SVGAElement {
	if condition {
		e.HREFLANG(s)
	}
	return e
}

func (e *SVGAElement) IfHREFLANGF(condition bool, format string, args ...any) *SVGAElement {
	if condition {
		e.HREFLANG(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute HREFLANG from the element.
func (e *SVGAElement) HREFLANGRemove(s string) *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("hreflang")
	return e
}

func (e *SVGAElement) HREFLANGRemoveF(format string, args ...any) *SVGAElement {
	return e.HREFLANGRemove(fmt.Sprintf(format, args...))
}

// A space-separated list of URLs
// When the link is followed, the browser should send POST requests with the body
// PING to the URLs
// Typically for tracking.
func (e *SVGAElement) PING(s ...string) *SVGAElement {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = treemap.New[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("ping")
	if !ok {
		ds = NewDelimitedBuilder[string](" ")
		e.DelimitedStrings.Set("ping", ds)
	}
	ds.Add(s...)
	return e
}

func (e *SVGAElement) IfPING(condition bool, s ...string) *SVGAElement {
	if condition {
		e.PING(s...)
	}
	return e
}

// Remove the attribute PING from the element.
func (e *SVGAElement) PINGRemove(s ...string) *SVGAElement {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("ping")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// Referrer policy to use when fetching the resource.
func (e *SVGAElement) REFERRERPOLICY(c SVGAReferrerpolicyChoice) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("referrerpolicy", string(c))
	return e
}

type SVGAReferrerpolicyChoice string

const (
	// The Referer header will not be sent.
	SVGAReferrerpolicy_no_referrer SVGAReferrerpolicyChoice = "no-referrer"
	// The Referer header will not be sent to origins without TLS (HTTPS).
	SVGAReferrerpolicy_no_referrer_when_downgrade SVGAReferrerpolicyChoice = "no-referrer-when-downgrade"
	// The Referer header will contain the origin of the request.
	SVGAReferrerpolicy_origin SVGAReferrerpolicyChoice = "origin"
	// The Referer header will contain the origin of the request, unless it is a
	// cross-origin request, in which case it will be omitted entirely.
	SVGAReferrerpolicy_origin_when_cross_origin SVGAReferrerpolicyChoice = "origin-when-cross-origin"
	// The Referer header will contain the origin of the request, unless it is a
	// cross-origin request, in which case it will be omitted entirely.
	SVGAReferrerpolicy_same_origin SVGAReferrerpolicyChoice = "same-origin"
	// The Referer header will contain the origin of the request, unless it is a
	// cross-origin request, in which case it will be omitted entirely.
	SVGAReferrerpolicy_strict_origin SVGAReferrerpolicyChoice = "strict-origin"
	// The Referer header will contain the origin of the request, unless it is a
	// cross-origin request, in which case it will be omitted entirely.
	SVGAReferrerpolicy_strict_origin_when_cross_origin SVGAReferrerpolicyChoice = "strict-origin-when-cross-origin"
	// The Referer header will contain the origin of the request, unless it is a
	// cross-origin request, in which case it will be omitted entirely.
	SVGAReferrerpolicy_unsafe_url SVGAReferrerpolicyChoice = "unsafe-url"
)

// Remove the attribute REFERRERPOLICY from the element.
func (e *SVGAElement) REFERRERPOLICYRemove(c SVGAReferrerpolicyChoice) *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("referrerpolicy")
	return e
}

// Specifies the relationship of the target object to the link object.
func (e *SVGAElement) REL(c SVGARelChoice) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("rel", string(c))
	return e
}

type SVGARelChoice string

const (
	// Links to an alternate version of the document (i.e
	// print page, translated or mirror).
	SVGARel_alternate SVGARelChoice = "alternate"
	// Links to the author of the document.
	SVGARel_author SVGARelChoice = "author"
	// Permanent URL used for bookmarking.
	SVGARel_bookmark SVGARelChoice = "bookmark"
	// Helps prevent duplicate content issues.
	SVGARel_canonical SVGARelChoice = "canonical"
	// Specifies that the browser should preemptively perform DNS resolution for the
	// target resource's origin.
	SVGARel_dns_prefetch SVGARelChoice = "dns-prefetch"
	// Links to an external resource (an external stylesheet).
	SVGARel_external SVGARelChoice = "external"
	// Links to a help document.
	SVGARel_help SVGARelChoice = "help"
	// Imports an icon to represent the document.
	SVGARel_icon SVGARelChoice = "icon"
	// Links to a license associated with the document.
	SVGARel_license SVGARelChoice = "license"
	// Specifies the location of a manifest file or an entry point for a web
	// application.
	SVGARel_manifest SVGARelChoice = "manifest"
	// Links to a resource that is the primary topic of the document.
	SVGARel_me SVGARelChoice = "me"
	// Specifies that the target resource should be preemptively fetched and cached by
	// the browser for later use.
	SVGARel_modulepreload SVGARelChoice = "modulepreload"
	// The next document in a selection.
	SVGARel_next SVGARelChoice = "next"
	// Links to an unendorsed document, like a paid link
	// ("nofollow" is used by Google, to specify that the Google search spider should
	// not follow that link.)
	SVGARel_nofollow SVGARelChoice = "nofollow"
	// Specifies that the browser should not open the linked document in a new tab or
	// window.
	SVGARel_noopener SVGARelChoice = "noopener"
	// Specifies that the browser should not send a HTTP referer header if the user
	// follows the hyperlink.
	SVGARel_noreferrer SVGARelChoice = "noreferrer"
	// Specifies that the target URL should be opened in a top-level browsing context
	// (that is, in the current tab or window).
	SVGARel_opener SVGARelChoice = "opener"
	// Links to the Pingback server of the current document.
	SVGARel_pingback SVGARelChoice = "pingback"
	// Specifies that the target resource should be loaded immediately.
	SVGARel_preload SVGARelChoice = "preload"
	// The previous document in a selection.
	SVGARel_prev SVGARelChoice = "prev"
	// Links to a privacy policy document associated with the current document.
	SVGARel_privacy_policy SVGARelChoice = "privacy-policy"
	// Links to a search tool for the document.
	SVGARel_search SVGARelChoice = "search"
	// Links to an external style sheet.
	SVGARel_stylesheet SVGARelChoice = "stylesheet"
	// A tag (keyword) for the current document.
	SVGARel_tag SVGARelChoice = "tag"
	// Links to the terms of service document for the current document.
	SVGARel_terms_of_service SVGARelChoice = "terms-of-service"
)

// Remove the attribute REL from the element.
func (e *SVGAElement) RELRemove(c SVGARelChoice) *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("rel")
	return e
}

// Specifies where to display the linked resource.
func (e *SVGAElement) TARGET(c SVGATargetChoice) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("target", string(c))
	return e
}

type SVGATargetChoice string

const (
	// Default
	// Opens the document in the same frame as it was clicked.
	SVGATarget__self SVGATargetChoice = "_self"
	// Opens the document in a new window or tab.
	SVGATarget__blank SVGATargetChoice = "_blank"
	// Opens the document in the parent frame.
	SVGATarget__parent SVGATargetChoice = "_parent"
	// Opens the document in the full body of the window.
	SVGATarget__top SVGATargetChoice = "_top"
	// Opens the document in a named frame.
	SVGATarget_framename SVGATargetChoice = "framename"
)

// Remove the attribute TARGET from the element.
func (e *SVGAElement) TARGETRemove(c SVGATargetChoice) *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("target")
	return e
}

// Specifies the MIME type of the linked resource.
func (e *SVGAElement) TYPE(s string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", s)
	return e
}

func (e *SVGAElement) TYPEF(format string, args ...any) *SVGAElement {
	return e.TYPE(fmt.Sprintf(format, args...))
}

func (e *SVGAElement) IfTYPE(condition bool, s string) *SVGAElement {
	if condition {
		e.TYPE(s)
	}
	return e
}

func (e *SVGAElement) IfTYPEF(condition bool, format string, args ...any) *SVGAElement {
	if condition {
		e.TYPE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute TYPE from the element.
func (e *SVGAElement) TYPERemove(s string) *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

func (e *SVGAElement) TYPERemoveF(format string, args ...any) *SVGAElement {
	return e.TYPERemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGAElement) ID(s string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGAElement) IDF(format string, args ...any) *SVGAElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGAElement) IfID(condition bool, s string) *SVGAElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGAElement) IfIDF(condition bool, format string, args ...any) *SVGAElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGAElement) IDRemove(s string) *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGAElement) IDRemoveF(format string, args ...any) *SVGAElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGAElement) CLASS(s ...string) *SVGAElement {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = treemap.New[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		ds = NewDelimitedBuilder[string](" ")
		e.DelimitedStrings.Set("class", ds)
	}
	ds.Add(s...)
	return e
}

func (e *SVGAElement) IfCLASS(condition bool, s ...string) *SVGAElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGAElement) CLASSRemove(s ...string) *SVGAElement {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGAElement) STYLEF(k string, format string, args ...any) *SVGAElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGAElement) IfSTYLE(condition bool, k string, v string) *SVGAElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGAElement) STYLE(k string, v string) *SVGAElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	kv.Add(k, v)
	return e
}

func (e *SVGAElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGAElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGAElement) STYLEMap(m map[string]string) *SVGAElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	for k, v := range m {
		kv.Add(k, v)
	}
	return e
}

// Add pairs of attributes to the element.
func (e *SVGAElement) STYLEPairs(pairs ...string) *SVGAElement {
	if len(pairs)%2 != 0 {
		panic("Must have an even number of pairs")
	}
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}

	for i := 0; i < len(pairs); i += 2 {
		kv.Add(pairs[i], pairs[i+1])
	}

	return e
}

func (e *SVGAElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGAElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGAElement) STYLERemove(keys ...string) *SVGAElement {
	if e.KVStrings == nil {
		return e
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		return e
	}
	for _, k := range keys {
		kv.Remove(k)
	}
	return e
}

// Merges the store with the given object

func (e *SVGAElement) DATASTAR_MERGE_STORE(v any) *SVGAElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("merge-store", string(b))
	return e
}

// Sets the reference of the element

func (e *SVGAElement) DATASTAR_REF(expression string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_REF(condition bool, expression string) *SVGAElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGAElement) DATASTAR_REFRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGAElement) DATASTAR_BIND(key string, expression string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGAElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGAElement) DATASTAR_BINDRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGAElement) DATASTAR_MODEL(expression string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGAElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGAElement) DATASTAR_MODELRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGAElement) DATASTAR_TEXT(expression string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGAElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGAElement) DATASTAR_TEXTRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGAOnMod customDataKeyModifier

// Debounces the event handler
func SVGAOnModDebounce(
	d time.Duration,
) SVGAOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGAOnModThrottle(
	d time.Duration,
) SVGAOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGAElement) DATASTAR_ON(key string, expression string, modifiers ...SVGAOnMod) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGAOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGAOnMod) *SVGAElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGAElement) DATASTAR_ONRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGAElement) DATASTAR_FOCUSSet(b bool) *SVGAElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGAElement) DATASTAR_FOCUS() *SVGAElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGAElement) DATASTAR_HEADER(key string, expression string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGAElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGAElement) DATASTAR_HEADERRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the URL for fetch requests

func (e *SVGAElement) DATASTAR_FETCH_URL(expression string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-url"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGAElement {
	if condition {
		e.DATASTAR_FETCH_URL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_URL from the element.
func (e *SVGAElement) DATASTAR_FETCH_URLRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-url")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGAElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGAElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGAElement) DATASTAR_FETCH_INDICATORRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGAElement) DATASTAR_SHOWSet(b bool) *SVGAElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGAElement) DATASTAR_SHOW() *SVGAElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGAElement) DATASTAR_INTERSECTS(expression string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGAElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGAElement) DATASTAR_INTERSECTSRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGAElement) DATASTAR_TELEPORTSet(b bool) *SVGAElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGAElement) DATASTAR_TELEPORT() *SVGAElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGAElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGAElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGAElement) DATASTAR_SCROLL_INTO_VIEW() *SVGAElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGAElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGAElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGAElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGAElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
