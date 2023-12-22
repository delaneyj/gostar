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

// The <feDiffuseLighting> SVG filter primitive lights an image using the alpha 
// channel as a bump map 
// The resulting image, which is an RGBA opaque image, depends on the light color, 
// light position and surface geometry of the input bump map. 
type FEDIFFUSELIGHTINGElementBuilder struct {
    *ElementBuilder
}

// Create a new FEDIFFUSELIGHTINGElementBuilder element.
// This will create a new element with the tag
// "feDiffuseLighting" during rendering.
func FEDIFFUSELIGHTING(children ...ElementRenderer) *FEDIFFUSELIGHTINGElementBuilder {
    return &FEDIFFUSELIGHTINGElementBuilder{
        ElementBuilder: &ElementBuilder{
            Tag: []byte("feDiffuseLighting"),
            IsSelfClosing: false,
            Descendants: children,
		},
    }
}

func (e *FEDIFFUSELIGHTINGElementBuilder) Children(children ...ElementRenderer) *FEDIFFUSELIGHTINGElementBuilder {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *FEDIFFUSELIGHTINGElementBuilder) IfChildren(condition bool, children ...ElementRenderer) *FEDIFFUSELIGHTINGElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *FEDIFFUSELIGHTINGElementBuilder) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *FEDIFFUSELIGHTINGElementBuilder {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *FEDIFFUSELIGHTINGElementBuilder) Text(text string) *FEDIFFUSELIGHTINGElementBuilder {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *FEDIFFUSELIGHTINGElementBuilder) TextF(format string, args ...any) *FEDIFFUSELIGHTINGElementBuilder {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *FEDIFFUSELIGHTINGElementBuilder) Escaped(text string) *FEDIFFUSELIGHTINGElementBuilder {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *FEDIFFUSELIGHTINGElementBuilder) EscapedF(format string, args ...any) *FEDIFFUSELIGHTINGElementBuilder {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *FEDIFFUSELIGHTINGElementBuilder) CustomData(key, value string) *FEDIFFUSELIGHTINGElementBuilder {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *FEDIFFUSELIGHTINGElementBuilder) CustomDataRemove(key string) *FEDIFFUSELIGHTINGElementBuilder {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


// The input for this filter. 
func(e *FEDIFFUSELIGHTINGElementBuilder) IN(s string) *FEDIFFUSELIGHTINGElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("in", s)
    return e
}

// Remove the attribute in from the element.
func(e *FEDIFFUSELIGHTINGElementBuilder) INRemove(s string) *FEDIFFUSELIGHTINGElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("in")
    return e
}


// The 'surfaceScale' attribute indicates the height of the surface when the alpha 
// channel is 1.0. 
func(e *FEDIFFUSELIGHTINGElementBuilder) SURFACESCALE(f float64) *FEDIFFUSELIGHTINGElementBuilder{
    if e.FloatAttributes == nil {
        e.FloatAttributes = treemap.New[string,float64]()
    }
    e.FloatAttributes.Set("surfaceScale", f)
    return e
}



// The diffuseConstant attribute represents the proportion of the light that is 
// reflected by the surface. 
func(e *FEDIFFUSELIGHTINGElementBuilder) DIFFUSECONSTANT(f float64) *FEDIFFUSELIGHTINGElementBuilder{
    if e.FloatAttributes == nil {
        e.FloatAttributes = treemap.New[string,float64]()
    }
    e.FloatAttributes.Set("diffuseConstant", f)
    return e
}



// The kernelUnitLength attribute defines the intended distance in current filter 
// units (i.e., units as determined by the value of attribute 'primitiveUnits') 
// for dx and dy in the surface normal calculation formulas. 
func(e *FEDIFFUSELIGHTINGElementBuilder) KERNELUNITLENGTH(s string) *FEDIFFUSELIGHTINGElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("kernelUnitLength", s)
    return e
}

// Remove the attribute kernelUnitLength from the element.
func(e *FEDIFFUSELIGHTINGElementBuilder) KERNELUNITLENGTHRemove(s string) *FEDIFFUSELIGHTINGElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("kernelUnitLength")
    return e
}


// Specifies one or more classnames for an element (refers to a class in a style 
// sheet) 
func(e *FEDIFFUSELIGHTINGElementBuilder) CLASS(s ...string) *FEDIFFUSELIGHTINGElementBuilder{
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
func(e *FEDIFFUSELIGHTINGElementBuilder) CLASSRemove(s ...string) *FEDIFFUSELIGHTINGElementBuilder{
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
func(e *FEDIFFUSELIGHTINGElementBuilder) ID(s string) *FEDIFFUSELIGHTINGElementBuilder{
    if e.StringAttributes == nil {
        e.StringAttributes = treemap.New[string,string]()
    }
    e.StringAttributes.Set("id", s)
    return e
}

// Remove the attribute id from the element.
func(e *FEDIFFUSELIGHTINGElementBuilder) IDRemove(s string) *FEDIFFUSELIGHTINGElementBuilder{
    if e.StringAttributes == nil {
        return e
    }
    e.StringAttributes.Del("id")
    return e
}


// Specifies an inline CSS style for an element 
func (e *FEDIFFUSELIGHTINGElementBuilder) STYLE(k string, v string) *FEDIFFUSELIGHTINGElementBuilder {
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
func (e *FEDIFFUSELIGHTINGElementBuilder) STYLEMap(m map[string]string) *FEDIFFUSELIGHTINGElementBuilder {
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
func (e *FEDIFFUSELIGHTINGElementBuilder) STYLEPairs(pairs ...string) *FEDIFFUSELIGHTINGElementBuilder {
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
func (e *FEDIFFUSELIGHTINGElementBuilder) STYLERemove(keys ...string) *FEDIFFUSELIGHTINGElementBuilder {
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


