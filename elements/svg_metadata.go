// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg metadata is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <metadata> SVG element allows to add metadata to SVG content
// Metadata is structured information about data
// In XML, metadata can be added to an element using for example attributes.
type SVGMETADATAElement struct {
	*Element
}

// Create a new SVGMETADATAElement element.
// This will create a new element with the tag
// "metadata" during rendering.
func SVG_METADATA(children ...ElementRenderer) *SVGMETADATAElement {
	e := NewElement("metadata", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGMETADATAElement{Element: e}
}

func (e *SVGMETADATAElement) Children(children ...ElementRenderer) *SVGMETADATAElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGMETADATAElement) IfChildren(condition bool, children ...ElementRenderer) *SVGMETADATAElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGMETADATAElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGMETADATAElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGMETADATAElement) Text(text string) *SVGMETADATAElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGMETADATAElement) TextF(format string, args ...any) *SVGMETADATAElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) Escaped(text string) *SVGMETADATAElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGMETADATAElement) EscapedF(format string, args ...any) *SVGMETADATAElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGMETADATAElement) CustomData(key, value string) *SVGMETADATAElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGMETADATAElement) CustomDataRemove(key string) *SVGMETADATAElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// A space-separated list of required extensions, indicating that the parent SVG
// document must include the specified extensions for this element to be valid.
func (e *SVGMETADATAElement) REQUIREDEXTENSIONS(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("requiredExtensions", s)
	return e
}

// Remove the attribute requiredExtensions from the element.
func (e *SVGMETADATAElement) REQUIREDEXTENSIONSRemove(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("requiredExtensions")
	return e
}

// A space-separated list of required features, indicating that the parent SVG
// document must include support for all of the specified features for this
// element to be valid.
func (e *SVGMETADATAElement) REQUIREDFEATURES(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("requiredFeatures", s)
	return e
}

// Remove the attribute requiredFeatures from the element.
func (e *SVGMETADATAElement) REQUIREDFEATURESRemove(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("requiredFeatures")
	return e
}

// A space-separated list of language codes, indicating that the parent SVG
// document must include support for all of the specified languages for this
// element to be valid.
func (e *SVGMETADATAElement) SYSTEMLANGUAGE(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("systemLanguage", s)
	return e
}

// Remove the attribute systemLanguage from the element.
func (e *SVGMETADATAElement) SYSTEMLANGUAGERemove(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("systemLanguage")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGMETADATAElement) CLASS(s ...string) *SVGMETADATAElement {
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

// Remove the attribute class from the element.
func (e *SVGMETADATAElement) CLASSRemove(s ...string) *SVGMETADATAElement {
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

// Specifies a unique id for an element
func (e *SVGMETADATAElement) ID(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGMETADATAElement) IDRemove(s string) *SVGMETADATAElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGMETADATAElement) STYLE(k string, v string) *SVGMETADATAElement {
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

// Add the attributes in the map to the element.
func (e *SVGMETADATAElement) STYLEMap(m map[string]string) *SVGMETADATAElement {
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
func (e *SVGMETADATAElement) STYLEPairs(pairs ...string) *SVGMETADATAElement {
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

// Remove the attribute style from the element.
func (e *SVGMETADATAElement) STYLERemove(keys ...string) *SVGMETADATAElement {
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