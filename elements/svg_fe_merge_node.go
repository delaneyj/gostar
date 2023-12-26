// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feMergeNode is generated from configuration file.
// Description:
package elements

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

// The <feMergeNode> SVG element allows a series of filter primitives to be
// connected together graphically
// Incoming nodes are blended into the background via the defined compositing
// operator.
type SVGFEMERGENODEElement struct {
	*Element
}

// Create a new SVGFEMERGENODEElement element.
// This will create a new element with the tag
// "feMergeNode" during rendering.
func SVG_FEMERGENODE(children ...ElementRenderer) *SVGFEMERGENODEElement {
	e := NewElement("feMergeNode", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFEMERGENODEElement{Element: e}
}

func (e *SVGFEMERGENODEElement) Children(children ...ElementRenderer) *SVGFEMERGENODEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFEMERGENODEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEMERGENODEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFEMERGENODEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEMERGENODEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFEMERGENODEElement) Text(text string) *SVGFEMERGENODEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFEMERGENODEElement) TextF(format string, args ...any) *SVGFEMERGENODEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGENODEElement) Escaped(text string) *SVGFEMERGENODEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFEMERGENODEElement) EscapedF(format string, args ...any) *SVGFEMERGENODEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEMERGENODEElement) CustomData(key, value string) *SVGFEMERGENODEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEMERGENODEElement) CustomDataRemove(key string) *SVGFEMERGENODEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The identifier for the input SVGAnimatedString attribute on the given
// 'feMergeNode' element.
func (e *SVGFEMERGENODEElement) IN(s string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("in", s)
	return e
}

// Remove the attribute in from the element.
func (e *SVGFEMERGENODEElement) INRemove(s string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("in")
	return e
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFEMERGENODEElement) CLASS(s ...string) *SVGFEMERGENODEElement {
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
func (e *SVGFEMERGENODEElement) CLASSRemove(s ...string) *SVGFEMERGENODEElement {
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
func (e *SVGFEMERGENODEElement) ID(s string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

// Remove the attribute id from the element.
func (e *SVGFEMERGENODEElement) IDRemove(s string) *SVGFEMERGENODEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFEMERGENODEElement) STYLE(k string, v string) *SVGFEMERGENODEElement {
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
func (e *SVGFEMERGENODEElement) STYLEMap(m map[string]string) *SVGFEMERGENODEElement {
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
func (e *SVGFEMERGENODEElement) STYLEPairs(pairs ...string) *SVGFEMERGENODEElement {
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
func (e *SVGFEMERGENODEElement) STYLERemove(keys ...string) *SVGFEMERGENODEElement {
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