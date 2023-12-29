// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg view is generated from configuration file.
// Description:
package elements

import(
    "fmt"
    "time"
    "github.com/igrmk/treemap/v2"
    "github.com/goccy/go-json"
    "github.com/samber/lo"
)

// The <view> SVG element is used to define a view into a &lt;svg&gt; element 
// It is partially deprecated in SVG 2.0 and should generally not be used. 
type SVGVIEWElement struct {
    *Element
}

// Create a new SVGVIEWElement element.
// This will create a new element with the tag
// "view" during rendering.
func SVG_VIEW(children ...ElementRenderer) *SVGVIEWElement {
    e := NewElement("view", children...)
    e.IsSelfClosing = false
    e.Descendants = children

    return &SVGVIEWElement{ Element: e }
}

func (e *SVGVIEWElement) Children(children ...ElementRenderer) *SVGVIEWElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *SVGVIEWElement) IfChildren(condition bool, children ...ElementRenderer) *SVGVIEWElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *SVGVIEWElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGVIEWElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *SVGVIEWElement) Text(text string) *SVGVIEWElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *SVGVIEWElement) TextF(format string, args ...any) *SVGVIEWElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGVIEWElement) IfText(condition bool, text string) *SVGVIEWElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(text))
    }
    return e
}

func (e *SVGVIEWElement) IfTextF(condition bool, format string, args ...any) *SVGVIEWElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
    }
    return e
}

func (e *SVGVIEWElement) Escaped(text string) *SVGVIEWElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *SVGVIEWElement) IfEscaped(condition bool, text string) *SVGVIEWElement {
    if condition {
        e.Descendants = append(e.Descendants, Escaped(text))
    }
    return e
}

func (e *SVGVIEWElement) EscapedF(format string, args ...any) *SVGVIEWElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGVIEWElement) IfEscapedF(condition bool, format string, args ...any) *SVGVIEWElement {
    if condition {
        e.Descendants = append(e.Descendants, EscapedF(format, args...))
    }
    return e
}

func (e *SVGVIEWElement) CustomData(key, value string) *SVGVIEWElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGVIEWElement) IfCustomData(condition bool, key, value string) *SVGVIEWElement {
    if condition {
        e.CustomData(key, value)
    }
    return e
}

func (e *SVGVIEWElement) CustomDataF(key, format string, args ...any) *SVGVIEWElement {
    return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGVIEWElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGVIEWElement {
    if condition {
        e.CustomData(key, fmt.Sprintf(format, args...))
    }
    return e
}

func (e *SVGVIEWElement) CustomDataRemove(key string) *SVGVIEWElement {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


    // The position and size of the viewport (the viewBox) is defined by the viewBox 
// attribute. 
    func(e *SVGVIEWElement) VIEW_BOX(s string) *SVGVIEWElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("viewBox", s)
            return e
        }

        func(e *SVGVIEWElement) IfVIEW_BOX(condition bool, s string) *SVGVIEWElement{
            if condition {
                e.VIEW_BOX(s)
            }
            return e
        }

        // Remove the attribute VIEW_BOX from the element.
        func(e *SVGVIEWElement) VIEW_BOXRemove(s string) *SVGVIEWElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("viewBox")
            return e
        }
    

    // Specifies a unique id for an element 
    func(e *SVGVIEWElement) ID(s string) *SVGVIEWElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("id", s)
            return e
        }

        func(e *SVGVIEWElement) IfID(condition bool, s string) *SVGVIEWElement{
            if condition {
                e.ID(s)
            }
            return e
        }

        // Remove the attribute ID from the element.
        func(e *SVGVIEWElement) IDRemove(s string) *SVGVIEWElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("id")
            return e
        }
    

    // Specifies one or more classnames for an element (refers to a class in a style 
// sheet) 
    func(e *SVGVIEWElement) CLASS(s ...string) *SVGVIEWElement{
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

        func(e *SVGVIEWElement) IfCLASS(condition bool, s ...string) *SVGVIEWElement{
            if condition {
                e.CLASS(s...)
            }
            return e
        }

        // Remove the attribute CLASS from the element.
        func(e *SVGVIEWElement) CLASSRemove(s ...string) *SVGVIEWElement{
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
    func (e *SVGVIEWElement) STYLEF(k string, format string, args ...any) *SVGVIEWElement {
            return e.STYLE(k, fmt.Sprintf(format, args...))
        }

        func (e *SVGVIEWElement) IfSTYLE(condition bool, k string, v string) *SVGVIEWElement {
            if condition {
                e.STYLE(k, v)
            }
            return e
        }

        func (e *SVGVIEWElement) STYLE(k string, v string) *SVGVIEWElement {
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

        func (e *SVGVIEWElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGVIEWElement {
            if condition {
                e.STYLE(k, fmt.Sprintf(format, args...))
            }
            return e
        }

        // Add the attributes in the map to the element.
        func (e *SVGVIEWElement) STYLEMap(m map[string]string) *SVGVIEWElement {
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
        func (e *SVGVIEWElement) STYLEPairs(pairs ...string) *SVGVIEWElement {
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

        func (e *SVGVIEWElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGVIEWElement {
            if condition {
                e.STYLEPairs(pairs...)
            }
            return e
        }

        // Remove the attribute STYLE from the element.
        func (e *SVGVIEWElement) STYLERemove(keys ...string) *SVGVIEWElement {
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
    
        func(e *SVGVIEWElement) DATASTAR_MERGE_STORE(v any) *SVGVIEWElement{
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
    
        func(e *SVGVIEWElement) DATASTAR_REF(expression string) *SVGVIEWElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-ref"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGVIEWElement) IfDATASTAR_REF(condition bool, expression string) *SVGVIEWElement{
                if condition {
                    e.DATASTAR_REF( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_REF from the element.
            func(e *SVGVIEWElement) DATASTAR_REFRemove() *SVGVIEWElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-ref")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *SVGVIEWElement) DATASTAR_BIND(key string, expression string) *SVGVIEWElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-bind-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGVIEWElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGVIEWElement{
                if condition {
                    e.DATASTAR_BIND(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_BIND from the element.
            func(e *SVGVIEWElement) DATASTAR_BINDRemove() *SVGVIEWElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-bind")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *SVGVIEWElement) DATASTAR_MODEL(expression string) *SVGVIEWElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-model"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGVIEWElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGVIEWElement{
                if condition {
                    e.DATASTAR_MODEL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_MODEL from the element.
            func(e *SVGVIEWElement) DATASTAR_MODELRemove() *SVGVIEWElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-model")
                return e
            }

        

    // Sets the textContent of the element 
    
        func(e *SVGVIEWElement) DATASTAR_TEXT(expression string) *SVGVIEWElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-text"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGVIEWElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGVIEWElement{
                if condition {
                    e.DATASTAR_TEXT( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_TEXT from the element.
            func(e *SVGVIEWElement) DATASTAR_TEXTRemove() *SVGVIEWElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-text")
                return e
            }

        

    // Sets the event handler of the element 
    
        type SVGViewDataOnMod customDataKeyModifier

            
            // Debounces the event handler 
            func SVGViewDataOnModDebounce(
                    d time.Duration,
            ) SVGViewDataOnMod {
                return func() string {return fmt.Sprintf("debounce_%dms", d.Milliseconds())
                }
            }
            
            // Throttles the event handler 
            func SVGViewDataOnModThrottle(
                    d time.Duration,
            ) SVGViewDataOnMod {
                return func() string {return fmt.Sprintf("throttle_%dms", d.Milliseconds())
                }
            }
            
        func(e *SVGVIEWElement) DATASTAR_ON(key string, expression string, modifiers ...SVGViewDataOnMod) *SVGVIEWElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-on-%s", key)
                
                customMods := lo.Map(modifiers, func(m SVGViewDataOnMod, i int) customDataKeyModifier  {
                    return customDataKeyModifier(m)
                })
                key = customDataKey(key, customMods...)
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGVIEWElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGViewDataOnMod) *SVGVIEWElement{
                if condition {
                    e.DATASTAR_ON(key,  expression,  modifiers...)
                }
                return e
            }

            // Remove the attribute DATASTAR_ON from the element.
            func(e *SVGVIEWElement) DATASTAR_ONRemove() *SVGVIEWElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-on")
                return e
            }

        

    // Sets the focus of the element 
    
        func(e *SVGVIEWElement) DATASTAR_FOCUSSet(b bool) *SVGVIEWElement{
                key := "data-focus"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGVIEWElement) DATASTAR_FOCUS() *SVGVIEWElement{
                return e.DATASTAR_FOCUSSet(true)
            }
        

    // Sets the header of for fetch requests 
    
        func(e *SVGVIEWElement) DATASTAR_HEADER(key string, expression string) *SVGVIEWElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-header-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGVIEWElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGVIEWElement{
                if condition {
                    e.DATASTAR_HEADER(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_HEADER from the element.
            func(e *SVGVIEWElement) DATASTAR_HEADERRemove() *SVGVIEWElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-header")
                return e
            }

        

    // Sets the URL for fetch requests 
    
        func(e *SVGVIEWElement) DATASTAR_FETCH_URL(expression string) *SVGVIEWElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-fetch-url"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGVIEWElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *SVGVIEWElement{
                if condition {
                    e.DATASTAR_FETCH_URL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_URL from the element.
            func(e *SVGVIEWElement) DATASTAR_FETCH_URLRemove() *SVGVIEWElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-fetch-url")
                return e
            }

        

    // Sets the indicator selector for fetch requests 
    
        func(e *SVGVIEWElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGVIEWElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "DatastarFetchIndicator"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGVIEWElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGVIEWElement{
                if condition {
                    e.DATASTAR_FETCH_INDICATOR( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
            func(e *SVGVIEWElement) DATASTAR_FETCH_INDICATORRemove() *SVGVIEWElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("DatastarFetchIndicator")
                return e
            }

        

    // Sets the visibility of the element 
    
        func(e *SVGVIEWElement) DATASTAR_SHOWSet(b bool) *SVGVIEWElement{
                key := "data-show"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGVIEWElement) DATASTAR_SHOW() *SVGVIEWElement{
                return e.DATASTAR_SHOWSet(true)
            }
        

    // Triggers the callback when the element intersects the viewport 
    
        func(e *SVGVIEWElement) DATASTAR_INTERSECTSSet(b bool) *SVGVIEWElement{
                key := "data-intersects"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGVIEWElement) DATASTAR_INTERSECTS() *SVGVIEWElement{
                return e.DATASTAR_INTERSECTSSet(true)
            }
        

    // Teleports the element to the given selector 
    
        func(e *SVGVIEWElement) DATASTAR_TELEPORTSet(b bool) *SVGVIEWElement{
                key := "data-teleport"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGVIEWElement) DATASTAR_TELEPORT() *SVGVIEWElement{
                return e.DATASTAR_TELEPORTSet(true)
            }
        

    // Scrolls the element into view 
    
        func(e *SVGVIEWElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGVIEWElement{
                key := "data-scroll-into-view"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *SVGVIEWElement) DATASTAR_SCROLL_INTO_VIEW() *SVGVIEWElement{
                return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
            }
        

    // Setup the ViewTransitionAPI for the element 
    
        func(e *SVGVIEWElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *SVGVIEWElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-view-transition-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *SVGVIEWElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *SVGVIEWElement{
                if condition {
                    e.DATASTAR_VIEW_TRANSITION(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
            func(e *SVGVIEWElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGVIEWElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-view-transition")
                return e
            }

        



