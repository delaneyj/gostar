// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feTile is generated from configuration file.
// Description:
package elements

import(
    "fmt"
    "time"
    "github.com/igrmk/treemap/v2"
    "github.com/goccy/go-json"
    "github.com/samber/lo"
)

// The <feTile> SVG filter primitive allows to fill a target rectangle with a 
// repeated, tiled pattern of an input image 
// The effect is similar to the one of a <pattern> element, but <feTile> can use 
// complex (i.e., filter) tree as input, and can be animated. 
type SVGFETILEElement struct {
    *Element
}

// Create a new SVGFETILEElement element.
// This will create a new element with the tag
// "feTile" during rendering.
func SVG_FETILE(children ...ElementRenderer) *SVGFETILEElement {
    e := NewElement("feTile", children...)
    e.IsSelfClosing = false
    e.Descendants = children

    return &SVGFETILEElement{ Element: e }
}

func (e *SVGFETILEElement) Children(children ...ElementRenderer) *SVGFETILEElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *SVGFETILEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFETILEElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *SVGFETILEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFETILEElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *SVGFETILEElement) Text(text string) *SVGFETILEElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *SVGFETILEElement) TextF(format string, args ...any) *SVGFETILEElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFETILEElement) IfText(condition bool, text string) *SVGFETILEElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(text))
    }
    return e
}

func (e *SVGFETILEElement) IfTextF(condition bool, format string, args ...any) *SVGFETILEElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
    }
    return e
}

func (e *SVGFETILEElement) Escaped(text string) *SVGFETILEElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *SVGFETILEElement) IfEscaped(condition bool, text string) *SVGFETILEElement {
    if condition {
        e.Descendants = append(e.Descendants, Escaped(text))
    }
    return e
}

func (e *SVGFETILEElement) EscapedF(format string, args ...any) *SVGFETILEElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFETILEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFETILEElement {
    if condition {
        e.Descendants = append(e.Descendants, EscapedF(format, args...))
    }
    return e
}

func (e *SVGFETILEElement) CustomData(key, value string) *SVGFETILEElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFETILEElement) IfCustomData(condition bool, key, value string) *SVGFETILEElement {
    if condition {
        e.CustomData(key, value)
    }
    return e
}

func (e *SVGFETILEElement) CustomDataF(key, format string, args ...any) *SVGFETILEElement {
    return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFETILEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFETILEElement {
    if condition {
        e.CustomData(key, fmt.Sprintf(format, args...))
    }
    return e
}

func (e *SVGFETILEElement) CustomDataRemove(key string) *SVGFETILEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


    // The input for this filter. 
    func(e *SVGFETILEElement) IN(s string) *SVGFETILEElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("in", s)
            return e
        }

        func(e *SVGFETILEElement) IfIN(condition bool, s string) *SVGFETILEElement{
            if condition {
                e.IN(s)
            }
            return e
        }

        // Remove the attribute IN from the element.
        func(e *SVGFETILEElement) INRemove(s string) *SVGFETILEElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("in")
            return e
        }
    

    // Specifies a unique id for an element 
    func(e *SVGFETILEElement) ID(s string) *SVGFETILEElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("id", s)
            return e
        }

        func(e *SVGFETILEElement) IfID(condition bool, s string) *SVGFETILEElement{
            if condition {
                e.ID(s)
            }
            return e
        }

        // Remove the attribute ID from the element.
        func(e *SVGFETILEElement) IDRemove(s string) *SVGFETILEElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("id")
            return e
        }
    

    // Specifies one or more classnames for an element (refers to a class in a style 
// sheet) 
    func(e *SVGFETILEElement) CLASS(s ...string) *SVGFETILEElement{
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

        func(e *SVGFETILEElement) IfCLASS(condition bool, s ...string) *SVGFETILEElement{
            if condition {
                e.CLASS(s...)
            }
            return e
        }

        // Remove the attribute CLASS from the element.
        func(e *SVGFETILEElement) CLASSRemove(s ...string) *SVGFETILEElement{
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
    func (e *SVGFETILEElement) STYLEF(k string, format string, args ...any) *SVGFETILEElement {
            return e.STYLE(k, fmt.Sprintf(format, args...))
        }

        func (e *SVGFETILEElement) IfSTYLE(condition bool, k string, v string) *SVGFETILEElement {
            if condition {
                e.STYLE(k, v)
            }
            return e
        }

        func (e *SVGFETILEElement) STYLE(k string, v string) *SVGFETILEElement {
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

        func (e *SVGFETILEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFETILEElement {
            if condition {
                e.STYLE(k, fmt.Sprintf(format, args...))
            }
            return e
        }

        // Add the attributes in the map to the element.
        func (e *SVGFETILEElement) STYLEMap(m map[string]string) *SVGFETILEElement {
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
        func (e *SVGFETILEElement) STYLEPairs(pairs ...string) *SVGFETILEElement {
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

        func (e *SVGFETILEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFETILEElement {
            if condition {
                e.STYLEPairs(pairs...)
            }
            return e
        }

        // Remove the attribute STYLE from the element.
        func (e *SVGFETILEElement) STYLERemove(keys ...string) *SVGFETILEElement {
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
    
        func(e *SVGFETILEElement) DATASTAR_MERGE_STORE(v any) *SVGFETILEElement{
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
    
        func(e *SVGFETILEElement) DATASTAR_REF(expression string) *SVGFETILEElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-ref"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFETILEElement) IfDATASTAR_REF(condition bool, expression string) *SVGFETILEElement{
                if condition {
                    e.DATASTAR_REF( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_REF from the element.
            func(e *SVGFETILEElement) DATASTAR_REFRemove() *SVGFETILEElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-ref")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *SVGFETILEElement) DATASTAR_BIND(key string, expression string) *SVGFETILEElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-bind-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFETILEElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFETILEElement{
                if condition {
                    e.DATASTAR_BIND(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_BIND from the element.
            func(e *SVGFETILEElement) DATASTAR_BINDRemove() *SVGFETILEElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-bind")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *SVGFETILEElement) DATASTAR_MODEL(expression string) *SVGFETILEElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-model"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFETILEElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFETILEElement{
                if condition {
                    e.DATASTAR_MODEL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_MODEL from the element.
            func(e *SVGFETILEElement) DATASTAR_MODELRemove() *SVGFETILEElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-model")
                return e
            }

        

    // Sets the textContent of the element 
    
        func(e *SVGFETILEElement) DATASTAR_TEXT(expression string) *SVGFETILEElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-text"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFETILEElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFETILEElement{
                if condition {
                    e.DATASTAR_TEXT( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_TEXT from the element.
            func(e *SVGFETILEElement) DATASTAR_TEXTRemove() *SVGFETILEElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-text")
                return e
            }

        

    // Sets the event handler of the element 
    
        type SVGFeTileDataOnMod customDataKeyModifier

            
            // Debounces the event handler 
            func SVGFeTileDataOnModDebounce(
                    d time.Duration,
            ) SVGFeTileDataOnMod {
                return func() string {return fmt.Sprintf("debounce_%dms", d.Milliseconds())
                }
            }
            
            // Throttles the event handler 
            func SVGFeTileDataOnModThrottle(
                    d time.Duration,
            ) SVGFeTileDataOnMod {
                return func() string {return fmt.Sprintf("throttle_%dms", d.Milliseconds())
                }
            }
            
        func(e *SVGFETILEElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeTileDataOnMod) *SVGFETILEElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-on-%s", key)
                
                customMods := lo.Map(modifiers, func(m SVGFeTileDataOnMod, i int) customDataKeyModifier  {
                    return customDataKeyModifier(m)
                })
                key = customDataKey(key, customMods...)
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFETILEElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeTileDataOnMod) *SVGFETILEElement{
                if condition {
                    e.DATASTAR_ON(key,  expression,  modifiers...)
                }
                return e
            }

            // Remove the attribute DATASTAR_ON from the element.
            func(e *SVGFETILEElement) DATASTAR_ONRemove() *SVGFETILEElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-on")
                return e
            }

        

    // Sets the focus of the element 
    
        func(e *SVGFETILEElement) DATASTAR_FOCUSSet(b bool) *SVGFETILEElement{
                key := "data-focus"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFETILEElement) DATASTAR_FOCUS() *SVGFETILEElement{
                return e.DATASTAR_FOCUSSet(true)
            }
        

    // Sets the header of for fetch requests 
    
        func(e *SVGFETILEElement) DATASTAR_HEADER(key string, expression string) *SVGFETILEElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-header-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFETILEElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFETILEElement{
                if condition {
                    e.DATASTAR_HEADER(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_HEADER from the element.
            func(e *SVGFETILEElement) DATASTAR_HEADERRemove() *SVGFETILEElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-header")
                return e
            }

        

    // Sets the URL for fetch requests 
    
        func(e *SVGFETILEElement) DATASTAR_FETCH_URL(expression string) *SVGFETILEElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-fetch-url"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFETILEElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGFETILEElement{
                if condition {
                    e.DATASTAR_FETCH_URL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_URL from the element.
            func(e *SVGFETILEElement) DATASTAR_FETCH_URLRemove() *SVGFETILEElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-fetch-url")
                return e
            }

        

    // Sets the indicator selector for fetch requests 
    
        func(e *SVGFETILEElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFETILEElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "DatastarFetchIndicator"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFETILEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFETILEElement{
                if condition {
                    e.DATASTAR_FETCH_INDICATOR( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
            func(e *SVGFETILEElement) DATASTAR_FETCH_INDICATORRemove() *SVGFETILEElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("DatastarFetchIndicator")
                return e
            }

        

    // Sets the visibility of the element 
    
        func(e *SVGFETILEElement) DATASTAR_SHOWSet(b bool) *SVGFETILEElement{
                key := "data-show"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFETILEElement) DATASTAR_SHOW() *SVGFETILEElement{
                return e.DATASTAR_SHOWSet(true)
            }
        

    // Triggers the callback when the element intersects the viewport 
    
        func(e *SVGFETILEElement) DATASTAR_INTERSECTSSet(b bool) *SVGFETILEElement{
                key := "data-intersects"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFETILEElement) DATASTAR_INTERSECTS() *SVGFETILEElement{
                return e.DATASTAR_INTERSECTSSet(true)
            }
        

    // Teleports the element to the given selector 
    
        func(e *SVGFETILEElement) DATASTAR_TELEPORTSet(b bool) *SVGFETILEElement{
                key := "data-teleport"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFETILEElement) DATASTAR_TELEPORT() *SVGFETILEElement{
                return e.DATASTAR_TELEPORTSet(true)
            }
        

    // Scrolls the element into view 
    
        func(e *SVGFETILEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFETILEElement{
                key := "data-scroll-into-view"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGFETILEElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFETILEElement{
                return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
            }
        

    // Setup the ViewTransitionAPI for the element 
    
        func(e *SVGFETILEElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGFETILEElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-view-transition-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGFETILEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGFETILEElement{
                if condition {
                    e.DATASTAR_VIEW_TRANSITION(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
            func(e *SVGFETILEElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFETILEElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-view-transition")
                return e
            }

        



