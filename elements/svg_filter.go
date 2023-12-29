// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg filter is generated from configuration file.
// Description:
package elements

import(
    "fmt"
    "time"
    "github.com/igrmk/treemap/v2"
    "github.com/goccy/go-json"
    "github.com/samber/lo"
)

// The <filter> SVG element defines a custom filter effect by grouping atomic 
// filter primitives 
// It is never rendered directly 
// A filter is referenced by using the filter attribute on the target SVG element 
// or via the filter CSS property. 
type SVGFILTERElement struct {
    *Element
}

// Create a new SVGFILTERElement element.
// This will create a new element with the tag
// "filter" during rendering.
func SVG_FILTER(children ...ElementRenderer) *SVGFILTERElement {
    e := NewElement("filter", children...)
    e.IsSelfClosing = false
    e.Descendants = children

    return &SVGFILTERElement{ Element: e }
}

func (e *SVGFILTERElement) Children(children ...ElementRenderer) *SVGFILTERElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *SVGFILTERElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFILTERElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *SVGFILTERElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFILTERElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *SVGFILTERElement) Text(text string) *SVGFILTERElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *SVGFILTERElement) TextF(format string, args ...any) *SVGFILTERElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfText(condition bool, text string) *SVGFILTERElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(text))
    }
    return e
}

func (e *SVGFILTERElement) IfTextF(condition bool, format string, args ...any) *SVGFILTERElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
    }
    return e
}

func (e *SVGFILTERElement) Escaped(text string) *SVGFILTERElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *SVGFILTERElement) IfEscaped(condition bool, text string) *SVGFILTERElement {
    if condition {
        e.Descendants = append(e.Descendants, Escaped(text))
    }
    return e
}

func (e *SVGFILTERElement) EscapedF(format string, args ...any) *SVGFILTERElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfEscapedF(condition bool, format string, args ...any) *SVGFILTERElement {
    if condition {
        e.Descendants = append(e.Descendants, EscapedF(format, args...))
    }
    return e
}

func (e *SVGFILTERElement) CustomData(key, value string) *SVGFILTERElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFILTERElement) IfCustomData(condition bool, key, value string) *SVGFILTERElement {
    if condition {
        e.CustomData(key, value)
    }
    return e
}

func (e *SVGFILTERElement) CustomDataF(key, format string, args ...any) *SVGFILTERElement {
    return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFILTERElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFILTERElement {
    if condition {
        e.CustomData(key, fmt.Sprintf(format, args...))
    }
    return e
}

func (e *SVGFILTERElement) CustomDataRemove(key string) *SVGFILTERElement {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


    // The coordinate system for attributes x, y, width and height. 
    func(e *SVGFILTERElement) FILTER_UNITS(c SVGFilterFilterUnitsChoice) *SVGFILTERElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("filterUnits", string(c))
            return e
        }

        type SVGFilterFilterUnitsChoice string
        const(
        // The coordinate system for attributes x, y, width and height. 
            SVGFilterFilterUnits_userSpaceOnUse SVGFilterFilterUnitsChoice = "userSpaceOnUse"
        // The coordinate system for attributes x, y, width and height. 
            SVGFilterFilterUnits_objectBoundingBox SVGFilterFilterUnitsChoice = "objectBoundingBox"
        )

        // Remove the attribute FILTER_UNITS from the element.
        func(e *SVGFILTERElement) FILTER_UNITSRemove(c SVGFilterFilterUnitsChoice) *SVGFILTERElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("filterUnits")
            return e
        }
        

    // The coordinate system for the various length values within the filter. 
    func(e *SVGFILTERElement) PRIMITIVE_UNITS(c SVGFilterPrimitiveUnitsChoice) *SVGFILTERElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("primitiveUnits", string(c))
            return e
        }

        type SVGFilterPrimitiveUnitsChoice string
        const(
        // The coordinate system for the various length values within the filter. 
            SVGFilterPrimitiveUnits_userSpaceOnUse SVGFilterPrimitiveUnitsChoice = "userSpaceOnUse"
        // The coordinate system for the various length values within the filter. 
            SVGFilterPrimitiveUnits_objectBoundingBox SVGFilterPrimitiveUnitsChoice = "objectBoundingBox"
        )

        // Remove the attribute PRIMITIVE_UNITS from the element.
        func(e *SVGFILTERElement) PRIMITIVE_UNITSRemove(c SVGFilterPrimitiveUnitsChoice) *SVGFILTERElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("primitiveUnits")
            return e
        }
        

    // The x attribute indicates where the left edge of the filter is placed. 
    func(e *SVGFILTERElement) X(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("x", s)
            return e
        }

        func(e *SVGFILTERElement) IfX(condition bool, s string) *SVGFILTERElement{
            if condition {
                e.X(s)
            }
            return e
        }

        // Remove the attribute X from the element.
        func(e *SVGFILTERElement) XRemove(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("x")
            return e
        }
    

    // The y attribute indicates where the top edge of the filter is placed. 
    func(e *SVGFILTERElement) Y(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("y", s)
            return e
        }

        func(e *SVGFILTERElement) IfY(condition bool, s string) *SVGFILTERElement{
            if condition {
                e.Y(s)
            }
            return e
        }

        // Remove the attribute Y from the element.
        func(e *SVGFILTERElement) YRemove(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("y")
            return e
        }
    

    // The width attribute indicates the width of the filter primitive box. 
    func(e *SVGFILTERElement) WIDTH(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("width", s)
            return e
        }

        func(e *SVGFILTERElement) IfWIDTH(condition bool, s string) *SVGFILTERElement{
            if condition {
                e.WIDTH(s)
            }
            return e
        }

        // Remove the attribute WIDTH from the element.
        func(e *SVGFILTERElement) WIDTHRemove(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("width")
            return e
        }
    

    // The height attribute indicates the height of the filter primitive box. 
    func(e *SVGFILTERElement) HEIGHT(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("height", s)
            return e
        }

        func(e *SVGFILTERElement) IfHEIGHT(condition bool, s string) *SVGFILTERElement{
            if condition {
                e.HEIGHT(s)
            }
            return e
        }

        // Remove the attribute HEIGHT from the element.
        func(e *SVGFILTERElement) HEIGHTRemove(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("height")
            return e
        }
    

    // Specifies a unique id for an element 
    func(e *SVGFILTERElement) ID(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("id", s)
            return e
        }

        func(e *SVGFILTERElement) IfID(condition bool, s string) *SVGFILTERElement{
            if condition {
                e.ID(s)
            }
            return e
        }

        // Remove the attribute ID from the element.
        func(e *SVGFILTERElement) IDRemove(s string) *SVGFILTERElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("id")
            return e
        }
    

    // Specifies one or more classnames for an element (refers to a class in a style 
// sheet) 
    func(e *SVGFILTERElement) CLASS(s ...string) *SVGFILTERElement{
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

        func(e *SVGFILTERElement) IfCLASS(condition bool, s ...string) *SVGFILTERElement{
            if condition {
                e.CLASS(s...)
            }
            return e
        }

        // Remove the attribute CLASS from the element.
        func(e *SVGFILTERElement) CLASSRemove(s ...string) *SVGFILTERElement{
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
    func (e *SVGFILTERElement) STYLEF(k string, format string, args ...any) *SVGFILTERElement {
            return e.STYLE(k, fmt.Sprintf(format, args...))
        }

        func (e *SVGFILTERElement) IfSTYLE(condition bool, k string, v string) *SVGFILTERElement {
            if condition {
                e.STYLE(k, v)
            }
            return e
        }

        func (e *SVGFILTERElement) STYLE(k string, v string) *SVGFILTERElement {
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

        func (e *SVGFILTERElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFILTERElement {
            if condition {
                e.STYLE(k, fmt.Sprintf(format, args...))
            }
            return e
        }

        // Add the attributes in the map to the element.
        func (e *SVGFILTERElement) STYLEMap(m map[string]string) *SVGFILTERElement {
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
        func (e *SVGFILTERElement) STYLEPairs(pairs ...string) *SVGFILTERElement {
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

        func (e *SVGFILTERElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFILTERElement {
            if condition {
                e.STYLEPairs(pairs...)
            }
            return e
        }

        // Remove the attribute STYLE from the element.
        func (e *SVGFILTERElement) STYLERemove(keys ...string) *SVGFILTERElement {
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
    
        func(e *SVGFILTERElement) DATASTAR_MERGE_STORE(v any) *SVGFILTERElement{
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
    
        func(e *SVGFILTERElement) DATASTAR_REF(expression string) *SVGFILTERElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-ref"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFILTERElement) IfDATASTAR_REF(condition bool, expression string) *SVGFILTERElement{
                if condition {
                    e.DATASTAR_REF( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_REF from the element.
            func(e *SVGFILTERElement) DATASTAR_REFRemove() *SVGFILTERElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-ref")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *SVGFILTERElement) DATASTAR_BIND(key string, expression string) *SVGFILTERElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-bind-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFILTERElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFILTERElement{
                if condition {
                    e.DATASTAR_BIND(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_BIND from the element.
            func(e *SVGFILTERElement) DATASTAR_BINDRemove() *SVGFILTERElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-bind")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *SVGFILTERElement) DATASTAR_MODEL(expression string) *SVGFILTERElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-model"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFILTERElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFILTERElement{
                if condition {
                    e.DATASTAR_MODEL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_MODEL from the element.
            func(e *SVGFILTERElement) DATASTAR_MODELRemove() *SVGFILTERElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-model")
                return e
            }

        

    // Sets the textContent of the element 
    
        func(e *SVGFILTERElement) DATASTAR_TEXT(expression string) *SVGFILTERElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-text"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFILTERElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFILTERElement{
                if condition {
                    e.DATASTAR_TEXT( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_TEXT from the element.
            func(e *SVGFILTERElement) DATASTAR_TEXTRemove() *SVGFILTERElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-text")
                return e
            }

        

    // Sets the event handler of the element 
    
        type SVGFilterDataOnMod customDataKeyModifier

            
            // Debounces the event handler 
            func SVGFilterDataOnModDebounce(
                    d time.Duration,
            ) SVGFilterDataOnMod {
                return func() string {return fmt.Sprintf("debounce_%dms", d.Milliseconds())
                }
            }
            
            // Throttles the event handler 
            func SVGFilterDataOnModThrottle(
                    d time.Duration,
            ) SVGFilterDataOnMod {
                return func() string {return fmt.Sprintf("throttle_%dms", d.Milliseconds())
                }
            }
            
        func(e *SVGFILTERElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFilterDataOnMod) *SVGFILTERElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-on-%s", key)
                
                customMods := lo.Map(modifiers, func(m SVGFilterDataOnMod, i int) customDataKeyModifier  {
                    return customDataKeyModifier(m)
                })
                key = customDataKey(key, customMods...)
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFILTERElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFilterDataOnMod) *SVGFILTERElement{
                if condition {
                    e.DATASTAR_ON(key,  expression,  modifiers...)
                }
                return e
            }

            // Remove the attribute DATASTAR_ON from the element.
            func(e *SVGFILTERElement) DATASTAR_ONRemove() *SVGFILTERElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-on")
                return e
            }

        

    // Sets the focus of the element 
    
        func(e *SVGFILTERElement) DATASTAR_FOCUSSet(b bool) *SVGFILTERElement{
                key := "data-focus"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFILTERElement) DATASTAR_FOCUS() *SVGFILTERElement{
                return e.DATASTAR_FOCUSSet(true)
            }
        

    // Sets the header of for fetch requests 
    
        func(e *SVGFILTERElement) DATASTAR_HEADER(key string, expression string) *SVGFILTERElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-header-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFILTERElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFILTERElement{
                if condition {
                    e.DATASTAR_HEADER(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_HEADER from the element.
            func(e *SVGFILTERElement) DATASTAR_HEADERRemove() *SVGFILTERElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-header")
                return e
            }

        

    // Sets the URL for fetch requests 
    
        func(e *SVGFILTERElement) DATASTAR_FETCH_URL(expression string) *SVGFILTERElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-fetch-url"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFILTERElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGFILTERElement{
                if condition {
                    e.DATASTAR_FETCH_URL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_URL from the element.
            func(e *SVGFILTERElement) DATASTAR_FETCH_URLRemove() *SVGFILTERElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-fetch-url")
                return e
            }

        

    // Sets the indicator selector for fetch requests 
    
        func(e *SVGFILTERElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFILTERElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "DatastarFetchIndicator"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFILTERElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFILTERElement{
                if condition {
                    e.DATASTAR_FETCH_INDICATOR( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
            func(e *SVGFILTERElement) DATASTAR_FETCH_INDICATORRemove() *SVGFILTERElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("DatastarFetchIndicator")
                return e
            }

        

    // Sets the visibility of the element 
    
        func(e *SVGFILTERElement) DATASTAR_SHOWSet(b bool) *SVGFILTERElement{
                key := "data-show"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFILTERElement) DATASTAR_SHOW() *SVGFILTERElement{
                return e.DATASTAR_SHOWSet(true)
            }
        

    // Triggers the callback when the element intersects the viewport 
    
        func(e *SVGFILTERElement) DATASTAR_INTERSECTSSet(b bool) *SVGFILTERElement{
                key := "data-intersects"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFILTERElement) DATASTAR_INTERSECTS() *SVGFILTERElement{
                return e.DATASTAR_INTERSECTSSet(true)
            }
        

    // Teleports the element to the given selector 
    
        func(e *SVGFILTERElement) DATASTAR_TELEPORTSet(b bool) *SVGFILTERElement{
                key := "data-teleport"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFILTERElement) DATASTAR_TELEPORT() *SVGFILTERElement{
                return e.DATASTAR_TELEPORTSet(true)
            }
        

    // Scrolls the element into view 
    
        func(e *SVGFILTERElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFILTERElement{
                key := "data-scroll-into-view"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFILTERElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFILTERElement{
                return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
            }
        

    // Setup the ViewTransitionAPI for the element 
    
        func(e *SVGFILTERElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGFILTERElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-view-transition-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFILTERElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGFILTERElement{
                if condition {
                    e.DATASTAR_VIEW_TRANSITION(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
            func(e *SVGFILTERElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFILTERElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-view-transition")
                return e
            }

        



