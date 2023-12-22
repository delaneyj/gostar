// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg is generated from configuration file.
// Description:
// Scalable Vector Graphics (SVG) is an XML-based markup language for describing 
// two-dimensional based vector graphics 
// As such, it's a text-based, open Web standard for describing images that can be 
// rendered cleanly at any size and are designed specifically to work well with 
// other web standards including CSS, DOM, JavaScript, and SMIL 
// SVG is, essentially, to graphics what HTML is to text 
// SVG images and their related behaviors are defined in XML text files, which 
// means they can be searched, indexed, scripted, and compressed 
// Additionally, this means they can be created and edited with any text editor or 
// with drawing software 
// Compared to classic bitmapped image formats such as JPEG or PNG, SVG-format 
// vector images can be rendered at any size without loss of quality and can be 
// easily localized by updating the text within them, without the need of a 
// graphical editor to do so 
// With proper libraries, SVG files can even be localized on-the-fly. 
package svg

import(
    "fmt"
    "github.com/igrmk/treemap/v2"
)

// The <pattern> SVG element fills a region with a pattern defined by an SVG 
// image. 
type PATTERNElementBuilder struct {
    *ElementBuilder
}

// Create a new PATTERNElementBuilder element.
// This will create a new element with the tag
// "pattern" during rendering.
func PATTERN(children ...ElementRenderer) *PATTERNElementBuilder {
    return &PATTERNElementBuilder{
        ElementBuilder: &ElementBuilder{
            Tag: []byte("pattern"),
            IsSelfClosing: false,
            Descendants: children,
		},
    }
}

func (e *PATTERNElementBuilder) Children(children ...ElementRenderer) *PATTERNElementBuilder {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *PATTERNElementBuilder) IfChildren(condition bool, children ...ElementRenderer) *PATTERNElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *PATTERNElementBuilder) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *PATTERNElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *PATTERNElementBuilder) Text(text string) *PATTERNElementBuilder {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *PATTERNElementBuilder) TextF(format string, args ...any) *PATTERNElementBuilder {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *PATTERNElementBuilder) Escaped(text string) *PATTERNElementBuilder {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *PATTERNElementBuilder) EscapedF(format string, args ...any) *PATTERNElementBuilder {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *PATTERNElementBuilder) CustomData(key, value string) *PATTERNElementBuilder {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *PATTERNElementBuilder) CustomDataRemove(key string) *PATTERNElementBuilder {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


// The coordinate system for attributes x, y, width and height. 
func(e *PATTERNElementBuilder) PATTERNUNITS(c PatternPatternUnitsChoice) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("patternUnits", string(c))
    return e
}

type PatternPatternUnitsChoice string
const(
// The coordinate system for attributes x, y, width and height. 
    PatternPatternUnits_userSpaceOnUse PatternPatternUnitsChoice = "userSpaceOnUse"
// The coordinate system for attributes x, y, width and height. 
    PatternPatternUnits_objectBoundingBox PatternPatternUnitsChoice = "objectBoundingBox"
)

// Remove the attribute patternUnits from the element.
func(e *PATTERNElementBuilder) PATTERNUNITSRemove(c PatternPatternUnitsChoice) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("patternUnits")
    return e
}


// The coordinate system for the various length values within the filter. 
func(e *PATTERNElementBuilder) PATTERNCONTENTUNITS(c PatternPatternContentUnitsChoice) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("patternContentUnits", string(c))
    return e
}

type PatternPatternContentUnitsChoice string
const(
// The coordinate system for the various length values within the filter. 
    PatternPatternContentUnits_userSpaceOnUse PatternPatternContentUnitsChoice = "userSpaceOnUse"
// The coordinate system for the various length values within the filter. 
    PatternPatternContentUnits_objectBoundingBox PatternPatternContentUnitsChoice = "objectBoundingBox"
)

// Remove the attribute patternContentUnits from the element.
func(e *PATTERNElementBuilder) PATTERNCONTENTUNITSRemove(c PatternPatternContentUnitsChoice) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("patternContentUnits")
    return e
}


// The definition of how the pattern is tiled, read about <a 
// href="https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/patternTransform">patternTransform</a>. 
func(e *PATTERNElementBuilder) PATTERNTRANSFORM(s string) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("patternTransform", s)
    return e
}

// Remove the attribute patternTransform from the element.
func(e *PATTERNElementBuilder) PATTERNTRANSFORMRemove(s string) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("patternTransform")
    return e
}


// The x-axis coordinate of the side of the rectangular region which is closest to 
// the user. 
func(e *PATTERNElementBuilder) X(f float64) *PATTERNElementBuilder{
    if e.FloatAttributes == nil {
        e.FloatAttributes = treemap.New[string,float64]()
    }
    e.FloatAttributes.Set("x", f)
    return e
}



// The y-axis coordinate of the side of the rectangular region which is closest to 
// the user. 
func(e *PATTERNElementBuilder) Y(f float64) *PATTERNElementBuilder{
    if e.FloatAttributes == nil {
        e.FloatAttributes = treemap.New[string,float64]()
    }
    e.FloatAttributes.Set("y", f)
    return e
}



// The width of the rectangular region. 
func(e *PATTERNElementBuilder) WIDTH(f float64) *PATTERNElementBuilder{
    if e.FloatAttributes == nil {
        e.FloatAttributes = treemap.New[string,float64]()
    }
    e.FloatAttributes.Set("width", f)
    return e
}



// The height of the rectangular region. 
func(e *PATTERNElementBuilder) HEIGHT(f float64) *PATTERNElementBuilder{
    if e.FloatAttributes == nil {
        e.FloatAttributes = treemap.New[string,float64]()
    }
    e.FloatAttributes.Set("height", f)
    return e
}



// A URI reference to the image to paint. 
func(e *PATTERNElementBuilder) HREF(s string) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("href", s)
    return e
}

// Remove the attribute href from the element.
func(e *PATTERNElementBuilder) HREFRemove(s string) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("href")
    return e
}


// Specifies one or more classnames for an element (refers to a class in a style 
// sheet) 
func(e *PATTERNElementBuilder) CLASS(s ...string) *PATTERNElementBuilder{
    if e.DelimitedStrings == nil {
        e.DelimitedStrings = treemap.New[string,*DelimitedBuilder[string]]()
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
func(e *PATTERNElementBuilder) CLASSRemove(s ...string) *PATTERNElementBuilder{
    if e.DelimitedStrings == nil {
        return e
    }
    ds, ok := e.DelimitedStrings.Get("class")
    if !ok {
        return e
    }
    ds.Remove(s ...)
    return e
}



// Specifies a unique id for an element 
func(e *PATTERNElementBuilder) ID(s string) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("id", s)
    return e
}

// Remove the attribute id from the element.
func(e *PATTERNElementBuilder) IDRemove(s string) *PATTERNElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("id")
    return e
}


// Specifies an inline CSS style for an element 
func (e *PATTERNElementBuilder) STYLE(k string, v string) *PATTERNElementBuilder {
    if e.KVStrings == nil {
        e.KVStrings = treemap.New[string,*KVBuilder]()
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
func (e *PATTERNElementBuilder) STYLEMap(m map[string]string) *PATTERNElementBuilder {
    if e.KVStrings == nil {
        e.KVStrings = treemap.New[string,*KVBuilder]()
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
func (e *PATTERNElementBuilder) STYLEPairs(pairs ...string) *PATTERNElementBuilder {
    if len(pairs) % 2 != 0 {
        panic("Must have an even number of pairs")
    }
    if e.KVStrings == nil {
        e.KVStrings = treemap.New[string,*KVBuilder]()
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
func (e *PATTERNElementBuilder) STYLERemove(keys ...string) *PATTERNElementBuilder {
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


