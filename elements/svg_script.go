// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg script is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <script> SVG element includes scripts, which can be used to trigger user
// interface events.
type SVGSCRIPTElement struct {
	*Element
}

// Create a new SVGSCRIPTElement element.
// This will create a new element with the tag
// "script" during rendering.
func SVG_SCRIPT(children ...ElementRenderer) *SVGSCRIPTElement {
	e := NewElement("script", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGSCRIPTElement{Element: e}
}

func (e *SVGSCRIPTElement) Children(children ...ElementRenderer) *SVGSCRIPTElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGSCRIPTElement) IfChildren(condition bool, children ...ElementRenderer) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGSCRIPTElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGSCRIPTElement) Text(text string) *SVGSCRIPTElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGSCRIPTElement) TextF(format string, args ...any) *SVGSCRIPTElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfText(condition bool, text string) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGSCRIPTElement) IfTextF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGSCRIPTElement) Escaped(text string) *SVGSCRIPTElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGSCRIPTElement) IfEscaped(condition bool, text string) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGSCRIPTElement) EscapedF(format string, args ...any) *SVGSCRIPTElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfEscapedF(condition bool, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGSCRIPTElement) CustomData(key, value string) *SVGSCRIPTElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGSCRIPTElement) IfCustomData(condition bool, key, value string) *SVGSCRIPTElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGSCRIPTElement) CustomDataF(key, format string, args ...any) *SVGSCRIPTElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGSCRIPTElement) CustomDataRemove(key string) *SVGSCRIPTElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The scripting language used for the given script element.
func (e *SVGSCRIPTElement) TYPE(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", s)
	return e
}

func (e *SVGSCRIPTElement) IfTYPE(condition bool, s string) *SVGSCRIPTElement {
	if condition {
		e.TYPE(s)
	}
	return e
}

// Remove the attribute type from the element.
func (e *SVGSCRIPTElement) TYPERemove(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

// A Uniform Resource Identifier (URI) reference that specifies the location of
// the script to execute.
func (e *SVGSCRIPTElement) HREF(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("href", s)
	return e
}

func (e *SVGSCRIPTElement) IfHREF(condition bool, s string) *SVGSCRIPTElement {
	if condition {
		e.HREF(s)
	}
	return e
}

// Remove the attribute href from the element.
func (e *SVGSCRIPTElement) HREFRemove(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("href")
	return e
}

// How the element handles crossorigin requests.
func (e *SVGSCRIPTElement) CROSSORIGIN(c SVGScriptCrossoriginChoice) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("crossorigin", string(c))
	return e
}

type SVGScriptCrossoriginChoice string

const (
	// How the element handles crossorigin requests.
	SVGScriptCrossorigin_anonymous SVGScriptCrossoriginChoice = "anonymous"
	// How the element handles crossorigin requests.
	SVGScriptCrossorigin_use_credentials SVGScriptCrossoriginChoice = "use-credentials"
)

// Remove the attribute crossorigin from the element.
func (e *SVGSCRIPTElement) CROSSORIGINRemove(c SVGScriptCrossoriginChoice) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("crossorigin")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGSCRIPTElement) CLASS(s ...string) *SVGSCRIPTElement {
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

func (e *SVGSCRIPTElement) IfCLASS(condition bool, s ...string) *SVGSCRIPTElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute class from the element.
func (e *SVGSCRIPTElement) CLASSRemove(s ...string) *SVGSCRIPTElement {
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
func (e *SVGSCRIPTElement) ID(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGSCRIPTElement) IfID(condition bool, s string) *SVGSCRIPTElement {
	if condition {
		e.ID(s)
	}
	return e
}

// Remove the attribute id from the element.
func (e *SVGSCRIPTElement) IDRemove(s string) *SVGSCRIPTElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGSCRIPTElement) STYLEF(k string, format string, args ...any) *SVGSCRIPTElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGSCRIPTElement) IfSTYLE(condition bool, k string, v string) *SVGSCRIPTElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGSCRIPTElement) STYLE(k string, v string) *SVGSCRIPTElement {
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

func (e *SVGSCRIPTElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGSCRIPTElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGSCRIPTElement) STYLEMap(m map[string]string) *SVGSCRIPTElement {
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
func (e *SVGSCRIPTElement) STYLEPairs(pairs ...string) *SVGSCRIPTElement {
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

func (e *SVGSCRIPTElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGSCRIPTElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute style from the element.
func (e *SVGSCRIPTElement) STYLERemove(keys ...string) *SVGSCRIPTElement {
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
