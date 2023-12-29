// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feDisplacementMap is generated from configuration file.
// Description:
package elements

import(
    "fmt"
    "time"
    "github.com/igrmk/treemap/v2"
    "github.com/goccy/go-json"
    "github.com/samber/lo"
)

// The <feDisplacementMap> SVG filter primitive uses the pixel values from the 
// image from in2 to spatially displace the image from in. 
type SVGFEDISPLACEMENTMAPElement struct {
    *Element
}

// Create a new SVGFEDISPLACEMENTMAPElement element.
// This will create a new element with the tag
// "feDisplacementMap" during rendering.
func SVG_FEDISPLACEMENTMAP(children ...ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
    e := NewElement("feDisplacementMap", children...)
    e.IsSelfClosing = false
    e.Descendants = children

    return &SVGFEDISPLACEMENTMAPElement{ Element: e }
}

func (e *SVGFEDISPLACEMENTMAPElement) Children(children ...ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *SVGFEDISPLACEMENTMAPElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *SVGFEDISPLACEMENTMAPElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFEDISPLACEMENTMAPElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *SVGFEDISPLACEMENTMAPElement) Text(text string) *SVGFEDISPLACEMENTMAPElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *SVGFEDISPLACEMENTMAPElement) TextF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfText(condition bool, text string) *SVGFEDISPLACEMENTMAPElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(text))
    }
    return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfTextF(condition bool, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
    }
    return e
}

func (e *SVGFEDISPLACEMENTMAPElement) Escaped(text string) *SVGFEDISPLACEMENTMAPElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfEscaped(condition bool, text string) *SVGFEDISPLACEMENTMAPElement {
    if condition {
        e.Descendants = append(e.Descendants, Escaped(text))
    }
    return e
}

func (e *SVGFEDISPLACEMENTMAPElement) EscapedF(format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfEscapedF(condition bool, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
    if condition {
        e.Descendants = append(e.Descendants, EscapedF(format, args...))
    }
    return e
}

func (e *SVGFEDISPLACEMENTMAPElement) CustomData(key, value string) *SVGFEDISPLACEMENTMAPElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFEDISPLACEMENTMAPElement) IfCustomData(condition bool, key, value string) *SVGFEDISPLACEMENTMAPElement {
    if condition {
        e.CustomData(key, value)
    }
    return e
}

func (e *SVGFEDISPLACEMENTMAPElement) CustomDataF(key, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
    return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFEDISPLACEMENTMAPElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
    if condition {
        e.CustomData(key, fmt.Sprintf(format, args...))
    }
    return e
}

func (e *SVGFEDISPLACEMENTMAPElement) CustomDataRemove(key string) *SVGFEDISPLACEMENTMAPElement {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


    // The input for this filter. 
    func(e *SVGFEDISPLACEMENTMAPElement) IN(s string) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("in", s)
            return e
        }

        func(e *SVGFEDISPLACEMENTMAPElement) IfIN(condition bool, s string) *SVGFEDISPLACEMENTMAPElement{
            if condition {
                e.IN(s)
            }
            return e
        }

        // Remove the attribute IN from the element.
        func(e *SVGFEDISPLACEMENTMAPElement) INRemove(s string) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("in")
            return e
        }
    

    // The displacement map 
// This attribute can take on the same values as the 'in' attribute. 
    func(e *SVGFEDISPLACEMENTMAPElement) IN_2(s string) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("in2", s)
            return e
        }

        func(e *SVGFEDISPLACEMENTMAPElement) IfIN_2(condition bool, s string) *SVGFEDISPLACEMENTMAPElement{
            if condition {
                e.IN_2(s)
            }
            return e
        }

        // Remove the attribute IN_2 from the element.
        func(e *SVGFEDISPLACEMENTMAPElement) IN_2Remove(s string) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("in2")
            return e
        }
    

    // The scale attribute defines the maximum value for the in2 displacement 
// A value of 0 disables the effect of the displacement map. 
    func(e *SVGFEDISPLACEMENTMAPElement) SCALE(f float64) *SVGFEDISPLACEMENTMAPElement{
            if e.FloatAttributes == nil {
                e.FloatAttributes = treemap.New[string,float64]()
            }
            e.FloatAttributes.Set("scale", f)
            return e
        }

        func (e *SVGFEDISPLACEMENTMAPElement) IfSCALE(condition bool, f float64) *SVGFEDISPLACEMENTMAPElement {
            if condition {
                e.SCALE(f)
            }
            return e
        }

    

    // The xChannelSelector attribute indicates which color channel from in2 to use to 
// displace the pixels in in the horizontal direction. 
    func(e *SVGFEDISPLACEMENTMAPElement) X_CHANNEL_SELECTOR(c SVGFeDisplacementMapXChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("xChannelSelector", string(c))
            return e
        }

        type SVGFeDisplacementMapXChannelSelectorChoice string
        const(
        // The red channel of in2 is used to displace the x coordinate of each pixel. 
            SVGFeDisplacementMapXChannelSelector_R SVGFeDisplacementMapXChannelSelectorChoice = "R"
        // The green channel of in2 is used to displace the x coordinate of each pixel. 
            SVGFeDisplacementMapXChannelSelector_G SVGFeDisplacementMapXChannelSelectorChoice = "G"
        // The blue channel of in2 is used to displace the x coordinate of each pixel. 
            SVGFeDisplacementMapXChannelSelector_B SVGFeDisplacementMapXChannelSelectorChoice = "B"
        // The alpha channel of in2 is used to displace the x coordinate of each pixel. 
            SVGFeDisplacementMapXChannelSelector_A SVGFeDisplacementMapXChannelSelectorChoice = "A"
        )

        // Remove the attribute X_CHANNEL_SELECTOR from the element.
        func(e *SVGFEDISPLACEMENTMAPElement) X_CHANNEL_SELECTORRemove(c SVGFeDisplacementMapXChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("xChannelSelector")
            return e
        }
        

    // The yChannelSelector attribute indicates which color channel from in2 to use to 
// displace the pixels in in the vertical direction. 
    func(e *SVGFEDISPLACEMENTMAPElement) Y_CHANNEL_SELECTOR(c SVGFeDisplacementMapYChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("yChannelSelector", string(c))
            return e
        }

        type SVGFeDisplacementMapYChannelSelectorChoice string
        const(
        // The red channel of in2 is used to displace the y coordinate of each pixel. 
            SVGFeDisplacementMapYChannelSelector_R SVGFeDisplacementMapYChannelSelectorChoice = "R"
        // The green channel of in2 is used to displace the y coordinate of each pixel. 
            SVGFeDisplacementMapYChannelSelector_G SVGFeDisplacementMapYChannelSelectorChoice = "G"
        // The blue channel of in2 is used to displace the y coordinate of each pixel. 
            SVGFeDisplacementMapYChannelSelector_B SVGFeDisplacementMapYChannelSelectorChoice = "B"
        // The alpha channel of in2 is used to displace the y coordinate of each pixel. 
            SVGFeDisplacementMapYChannelSelector_A SVGFeDisplacementMapYChannelSelectorChoice = "A"
        )

        // Remove the attribute Y_CHANNEL_SELECTOR from the element.
        func(e *SVGFEDISPLACEMENTMAPElement) Y_CHANNEL_SELECTORRemove(c SVGFeDisplacementMapYChannelSelectorChoice) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("yChannelSelector")
            return e
        }
        

    // Specifies a unique id for an element 
    func(e *SVGFEDISPLACEMENTMAPElement) ID(s string) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("id", s)
            return e
        }

        func(e *SVGFEDISPLACEMENTMAPElement) IfID(condition bool, s string) *SVGFEDISPLACEMENTMAPElement{
            if condition {
                e.ID(s)
            }
            return e
        }

        // Remove the attribute ID from the element.
        func(e *SVGFEDISPLACEMENTMAPElement) IDRemove(s string) *SVGFEDISPLACEMENTMAPElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("id")
            return e
        }
    

    // Specifies one or more classnames for an element (refers to a class in a style 
// sheet) 
    func(e *SVGFEDISPLACEMENTMAPElement) CLASS(s ...string) *SVGFEDISPLACEMENTMAPElement{
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

        func(e *SVGFEDISPLACEMENTMAPElement) IfCLASS(condition bool, s ...string) *SVGFEDISPLACEMENTMAPElement{
            if condition {
                e.CLASS(s...)
            }
            return e
        }

        // Remove the attribute CLASS from the element.
        func(e *SVGFEDISPLACEMENTMAPElement) CLASSRemove(s ...string) *SVGFEDISPLACEMENTMAPElement{
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

    

    // Specifies an inline CSS style for an element 
    func (e *SVGFEDISPLACEMENTMAPElement) STYLEF(k string, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
            return e.STYLE(k, fmt.Sprintf(format, args...))
        }

        func (e *SVGFEDISPLACEMENTMAPElement) IfSTYLE(condition bool, k string, v string) *SVGFEDISPLACEMENTMAPElement {
            if condition {
                e.STYLE(k, v)
            }
            return e
        }

        func (e *SVGFEDISPLACEMENTMAPElement) STYLE(k string, v string) *SVGFEDISPLACEMENTMAPElement {
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

        func (e *SVGFEDISPLACEMENTMAPElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFEDISPLACEMENTMAPElement {
            if condition {
                e.STYLE(k, fmt.Sprintf(format, args...))
            }
            return e
        }

        // Add the attributes in the map to the element.
        func (e *SVGFEDISPLACEMENTMAPElement) STYLEMap(m map[string]string) *SVGFEDISPLACEMENTMAPElement {
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
        func (e *SVGFEDISPLACEMENTMAPElement) STYLEPairs(pairs ...string) *SVGFEDISPLACEMENTMAPElement {
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

        func (e *SVGFEDISPLACEMENTMAPElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFEDISPLACEMENTMAPElement {
            if condition {
                e.STYLEPairs(pairs...)
            }
            return e
        }

        // Remove the attribute STYLE from the element.
        func (e *SVGFEDISPLACEMENTMAPElement) STYLERemove(keys ...string) *SVGFEDISPLACEMENTMAPElement {
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
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_MERGE_STORE(v any) *SVGFEDISPLACEMENTMAPElement{
                if e.CustomDataAttributes == nil {
                    e.CustomDataAttributes = treemap.New[string,string]()
                }
                b, err := json.Marshal(v)
                if err != nil {
                    panic(err)
                }
                e.CustomDataAttributes.Set("data-merge-store", string(b))
                return e
            }

        

    // Sets the reference of the element 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_REF(expression string) *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-ref"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) IfDATASTAR_REF(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement{
                if condition {
                    e.DATASTAR_REF( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_REF from the element.
            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_REFRemove() *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-ref")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_BIND(key string, expression string) *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-bind-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFEDISPLACEMENTMAPElement{
                if condition {
                    e.DATASTAR_BIND(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_BIND from the element.
            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_BINDRemove() *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-bind")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_MODEL(expression string) *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-model"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement{
                if condition {
                    e.DATASTAR_MODEL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_MODEL from the element.
            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_MODELRemove() *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-model")
                return e
            }

        

    // Sets the textContent of the element 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_TEXT(expression string) *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-text"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement{
                if condition {
                    e.DATASTAR_TEXT( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_TEXT from the element.
            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_TEXTRemove() *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-text")
                return e
            }

        

    // Sets the event handler of the element 
    
        type SVGFeDisplacementMapDataOnMod customDataKeyModifier

            
            // Debounces the event handler 
            func SVGFeDisplacementMapDataOnModDebounce(
                    d time.Duration,
            ) SVGFeDisplacementMapDataOnMod {
                return func() string {return fmt.Sprintf("debounce_%dms", d.Milliseconds())
                }
            }
            
            // Throttles the event handler 
            func SVGFeDisplacementMapDataOnModThrottle(
                    d time.Duration,
            ) SVGFeDisplacementMapDataOnMod {
                return func() string {return fmt.Sprintf("throttle_%dms", d.Milliseconds())
                }
            }
            
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeDisplacementMapDataOnMod) *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-on-%s", key)
                
                customMods := lo.Map(modifiers, func(m SVGFeDisplacementMapDataOnMod, i int) customDataKeyModifier  {
                    return customDataKeyModifier(m)
                })
                key = customDataKey(key, customMods...)
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeDisplacementMapDataOnMod) *SVGFEDISPLACEMENTMAPElement{
                if condition {
                    e.DATASTAR_ON(key,  expression,  modifiers...)
                }
                return e
            }

            // Remove the attribute DATASTAR_ON from the element.
            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_ONRemove() *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-on")
                return e
            }

        

    // Sets the focus of the element 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_FOCUSSet(b bool) *SVGFEDISPLACEMENTMAPElement{
                key := "data-focus"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_FOCUS() *SVGFEDISPLACEMENTMAPElement{
                return e.DATASTAR_FOCUSSet(true)
            }
        

    // Sets the header of for fetch requests 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_HEADER(key string, expression string) *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-header-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFEDISPLACEMENTMAPElement{
                if condition {
                    e.DATASTAR_HEADER(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_HEADER from the element.
            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_HEADERRemove() *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-header")
                return e
            }

        

    // Sets the URL for fetch requests 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_FETCH_URL(expression string) *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-fetch-url"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement{
                if condition {
                    e.DATASTAR_FETCH_URL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_URL from the element.
            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_FETCH_URLRemove() *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-fetch-url")
                return e
            }

        

    // Sets the indicator selector for fetch requests 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "DatastarFetchIndicator"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFEDISPLACEMENTMAPElement{
                if condition {
                    e.DATASTAR_FETCH_INDICATOR( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_FETCH_INDICATORRemove() *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("DatastarFetchIndicator")
                return e
            }

        

    // Sets the visibility of the element 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_SHOWSet(b bool) *SVGFEDISPLACEMENTMAPElement{
                key := "data-show"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_SHOW() *SVGFEDISPLACEMENTMAPElement{
                return e.DATASTAR_SHOWSet(true)
            }
        

    // Triggers the callback when the element intersects the viewport 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_INTERSECTSSet(b bool) *SVGFEDISPLACEMENTMAPElement{
                key := "data-intersects"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_INTERSECTS() *SVGFEDISPLACEMENTMAPElement{
                return e.DATASTAR_INTERSECTSSet(true)
            }
        

    // Teleports the element to the given selector 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_TELEPORTSet(b bool) *SVGFEDISPLACEMENTMAPElement{
                key := "data-teleport"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_TELEPORT() *SVGFEDISPLACEMENTMAPElement{
                return e.DATASTAR_TELEPORTSet(true)
            }
        

    // Scrolls the element into view 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFEDISPLACEMENTMAPElement{
                key := "data-scroll-into-view"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFEDISPLACEMENTMAPElement{
                return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
            }
        

    // Setup the ViewTransitionAPI for the element 
    
        func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-view-transition-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFEDISPLACEMENTMAPElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGFEDISPLACEMENTMAPElement{
                if condition {
                    e.DATASTAR_VIEW_TRANSITION(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
            func(e *SVGFEDISPLACEMENTMAPElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFEDISPLACEMENTMAPElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-view-transition")
                return e
            }

        



