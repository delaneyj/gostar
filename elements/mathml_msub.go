// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml msub is generated from configuration file.
// Description:
package elements

import(
    "fmt"
    "time"
    "github.com/igrmk/treemap/v2"
    "github.com/goccy/go-json"
    "github.com/samber/lo"
)

// This element is used to display a subscript expression. 
type MathMLMSUBElement struct {
    *Element
}

// Create a new MathMLMSUBElement element.
// This will create a new element with the tag
// "msub" during rendering.
func MathML_MSUB(children ...ElementRenderer) *MathMLMSUBElement {
    e := NewElement("msub", children...)
    e.IsSelfClosing = false
    e.Descendants = children

    return &MathMLMSUBElement{ Element: e }
}

func (e *MathMLMSUBElement) Children(children ...ElementRenderer) *MathMLMSUBElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *MathMLMSUBElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLMSUBElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *MathMLMSUBElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLMSUBElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *MathMLMSUBElement) Text(text string) *MathMLMSUBElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *MathMLMSUBElement) TextF(format string, args ...any) *MathMLMSUBElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfText(condition bool, text string) *MathMLMSUBElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(text))
    }
    return e
}

func (e *MathMLMSUBElement) IfTextF(condition bool, format string, args ...any) *MathMLMSUBElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
    }
    return e
}

func (e *MathMLMSUBElement) Escaped(text string) *MathMLMSUBElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *MathMLMSUBElement) IfEscaped(condition bool, text string) *MathMLMSUBElement {
    if condition {
        e.Descendants = append(e.Descendants, Escaped(text))
    }
    return e
}

func (e *MathMLMSUBElement) EscapedF(format string, args ...any) *MathMLMSUBElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfEscapedF(condition bool, format string, args ...any) *MathMLMSUBElement {
    if condition {
        e.Descendants = append(e.Descendants, EscapedF(format, args...))
    }
    return e
}

func (e *MathMLMSUBElement) CustomData(key, value string) *MathMLMSUBElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLMSUBElement) IfCustomData(condition bool, key, value string) *MathMLMSUBElement {
    if condition {
        e.CustomData(key, value)
    }
    return e
}

func (e *MathMLMSUBElement) CustomDataF(key, format string, args ...any) *MathMLMSUBElement {
    return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLMSUBElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLMSUBElement {
    if condition {
        e.CustomData(key, fmt.Sprintf(format, args...))
    }
    return e
}

func (e *MathMLMSUBElement) CustomDataRemove(key string) *MathMLMSUBElement {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


    // Assigns a class name or set of class names to an element 
// You may assign the same class name or names to any number of elements 
// If you specify multiple class names, they must be separated by whitespace 
// characters. 
    func(e *MathMLMSUBElement) CLASS(s ...string) *MathMLMSUBElement{
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

        func(e *MathMLMSUBElement) IfCLASS(condition bool, s ...string) *MathMLMSUBElement{
            if condition {
                e.CLASS(s...)
            }
            return e
        }

        // Remove the attribute CLASS from the element.
        func(e *MathMLMSUBElement) CLASSRemove(s ...string) *MathMLMSUBElement{
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

    

    // This attribute specifies the text directionality of the element, merely 
// indicating what direction the text flows when surrounded by text with inherent 
// directionality (such as Arabic or Hebrew) 
// Possible values are ltr (left-to-right) and rtl (right-to-left). 
    func(e *MathMLMSUBElement) DIR(c MathMLMsubDirChoice) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("dir", string(c))
            return e
        }

        type MathMLMsubDirChoice string
        const(
        // left-to-right 
            MathMLMsubDir_ltr MathMLMsubDirChoice = "ltr"
        // right-to-left 
            MathMLMsubDir_rtl MathMLMsubDirChoice = "rtl"
        )

        // Remove the attribute DIR from the element.
        func(e *MathMLMSUBElement) DIRRemove(c MathMLMsubDirChoice) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("dir")
            return e
        }
        

    // This attribute specifies whether the element should be rendered using 
// displaystyle rules or not 
// Possible values are true and false. 
    func(e *MathMLMSUBElement) DISPLAYSTYLE(c MathMLMsubDisplaystyleChoice) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("displaystyle", string(c))
            return e
        }

        type MathMLMsubDisplaystyleChoice string
        const(
        // displaystyle rules 
            MathMLMsubDisplaystyle_true MathMLMsubDisplaystyleChoice = "true"
        // not displaystyle rules 
            MathMLMsubDisplaystyle_false MathMLMsubDisplaystyleChoice = "false"
        )

        // Remove the attribute DISPLAYSTYLE from the element.
        func(e *MathMLMSUBElement) DISPLAYSTYLERemove(c MathMLMsubDisplaystyleChoice) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("displaystyle")
            return e
        }
        

    // This attribute assigns a name to an element 
// This name must be unique in a document. 
    func(e *MathMLMSUBElement) ID(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("id", s)
            return e
        }

        func(e *MathMLMSUBElement) IfID(condition bool, s string) *MathMLMSUBElement{
            if condition {
                e.ID(s)
            }
            return e
        }

        // Remove the attribute ID from the element.
        func(e *MathMLMSUBElement) IDRemove(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("id")
            return e
        }
    

    // This attribute specifies the background color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
    func(e *MathMLMSUBElement) MATHBACKGROUND(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("mathbackground", s)
            return e
        }

        func(e *MathMLMSUBElement) IfMATHBACKGROUND(condition bool, s string) *MathMLMSUBElement{
            if condition {
                e.MATHBACKGROUND(s)
            }
            return e
        }

        // Remove the attribute MATHBACKGROUND from the element.
        func(e *MathMLMSUBElement) MATHBACKGROUNDRemove(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("mathbackground")
            return e
        }
    

    // This attribute specifies the color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
    func(e *MathMLMSUBElement) MATHCOLOR(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("mathcolor", s)
            return e
        }

        func(e *MathMLMSUBElement) IfMATHCOLOR(condition bool, s string) *MathMLMSUBElement{
            if condition {
                e.MATHCOLOR(s)
            }
            return e
        }

        // Remove the attribute MATHCOLOR from the element.
        func(e *MathMLMSUBElement) MATHCOLORRemove(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("mathcolor")
            return e
        }
    

    // This attribute specifies the size of the element 
// Possible values are a dimension or a dimensionless number. 
    func(e *MathMLMSUBElement) MATHSIZE_STR(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("mathsize", s)
            return e
        }

        func(e *MathMLMSUBElement) IfMATHSIZE_STR(condition bool, s string) *MathMLMSUBElement{
            if condition {
                e.MATHSIZE_STR(s)
            }
            return e
        }

        // Remove the attribute MATHSIZE_STR from the element.
        func(e *MathMLMSUBElement) MATHSIZE_STRRemove(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("mathsize")
            return e
        }
    

    // This attribute declares a cryptographic nonce (number used once) that should be 
// used by the server processing the element’s submission, and the resulting 
// resource must be delivered with a Content-Security-Policy nonce attribute 
// matching the value of the nonce attribute. 
    func(e *MathMLMSUBElement) NONCE(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("nonce", s)
            return e
        }

        func(e *MathMLMSUBElement) IfNONCE(condition bool, s string) *MathMLMSUBElement{
            if condition {
                e.NONCE(s)
            }
            return e
        }

        // Remove the attribute NONCE from the element.
        func(e *MathMLMSUBElement) NONCERemove(s string) *MathMLMSUBElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("nonce")
            return e
        }
    

    // This attribute specifies the script level of the element 
// Possible values are an integer between 0 and 7, inclusive. 
    func(e *MathMLMSUBElement) SCRIPTLEVEL(i int) *MathMLMSUBElement{
            if e.IntAttributes == nil {
                e.IntAttributes = treemap.New[string,int]()
            }
            e.IntAttributes.Set("scriptlevel", i)
            return e
        }

        func (e *MathMLMSUBElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLMSUBElement {
            if condition {
                e.SCRIPTLEVEL(i)
            }
            return e
        }

        // Remove the attribute SCRIPTLEVEL from the element.
        func(e *MathMLMSUBElement) SCRIPTLEVELRemove(i int) *MathMLMSUBElement{
            if e.IntAttributes == nil {
                return e
            }
            e.IntAttributes.Del("scriptlevel")
            return e
        }
        

    // This attribute offers advisory information about the element for which it is 
// set. 
    func (e *MathMLMSUBElement) STYLEF(k string, format string, args ...any) *MathMLMSUBElement {
            return e.STYLE(k, fmt.Sprintf(format, args...))
        }

        func (e *MathMLMSUBElement) IfSTYLE(condition bool, k string, v string) *MathMLMSUBElement {
            if condition {
                e.STYLE(k, v)
            }
            return e
        }

        func (e *MathMLMSUBElement) STYLE(k string, v string) *MathMLMSUBElement {
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

        func (e *MathMLMSUBElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLMSUBElement {
            if condition {
                e.STYLE(k, fmt.Sprintf(format, args...))
            }
            return e
        }

        // Add the attributes in the map to the element.
        func (e *MathMLMSUBElement) STYLEMap(m map[string]string) *MathMLMSUBElement {
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
        func (e *MathMLMSUBElement) STYLEPairs(pairs ...string) *MathMLMSUBElement {
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

        func (e *MathMLMSUBElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLMSUBElement {
            if condition {
                e.STYLEPairs(pairs...)
            }
            return e
        }

        // Remove the attribute STYLE from the element.
        func (e *MathMLMSUBElement) STYLERemove(keys ...string) *MathMLMSUBElement {
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

    

    // This attribute specifies the position of the current element in the tabbing 
// order for the current document 
// This value must be a number between 0 and 32767 
// User agents should ignore leading zeros. 
    func(e *MathMLMSUBElement) TABINDEX(i int) *MathMLMSUBElement{
            if e.IntAttributes == nil {
                e.IntAttributes = treemap.New[string,int]()
            }
            e.IntAttributes.Set("tabindex", i)
            return e
        }

        func (e *MathMLMSUBElement) IfTABINDEX(condition bool, i int) *MathMLMSUBElement {
            if condition {
                e.TABINDEX(i)
            }
            return e
        }

        // Remove the attribute TABINDEX from the element.
        func(e *MathMLMSUBElement) TABINDEXRemove(i int) *MathMLMSUBElement{
            if e.IntAttributes == nil {
                return e
            }
            e.IntAttributes.Del("tabindex")
            return e
        }
        

    // Merges the store with the given object 
    
        func(e *MathMLMSUBElement) DATASTAR_MERGE_STORE(v any) *MathMLMSUBElement{
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
    
        func(e *MathMLMSUBElement) DATASTAR_REF(expression string) *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-ref"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLMSUBElement) IfDATASTAR_REF(condition bool, expression string) *MathMLMSUBElement{
                if condition {
                    e.DATASTAR_REF( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_REF from the element.
            func(e *MathMLMSUBElement) DATASTAR_REFRemove() *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-ref")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *MathMLMSUBElement) DATASTAR_BIND(key string, expression string) *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-bind-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLMSUBElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLMSUBElement{
                if condition {
                    e.DATASTAR_BIND(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_BIND from the element.
            func(e *MathMLMSUBElement) DATASTAR_BINDRemove() *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-bind")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *MathMLMSUBElement) DATASTAR_MODEL(expression string) *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-model"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLMSUBElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLMSUBElement{
                if condition {
                    e.DATASTAR_MODEL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_MODEL from the element.
            func(e *MathMLMSUBElement) DATASTAR_MODELRemove() *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-model")
                return e
            }

        

    // Sets the textContent of the element 
    
        func(e *MathMLMSUBElement) DATASTAR_TEXT(expression string) *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-text"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLMSUBElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLMSUBElement{
                if condition {
                    e.DATASTAR_TEXT( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_TEXT from the element.
            func(e *MathMLMSUBElement) DATASTAR_TEXTRemove() *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-text")
                return e
            }

        

    // Sets the event handler of the element 
    
        type MathMLMsubDataOnMod customDataKeyModifier

            
            // Debounces the event handler 
            func MathMLMsubDataOnModDebounce(
                    d time.Duration,
            ) MathMLMsubDataOnMod {
                return func() string {return fmt.Sprintf("debounce_%dms", d.Milliseconds())
                }
            }
            
            // Throttles the event handler 
            func MathMLMsubDataOnModThrottle(
                    d time.Duration,
            ) MathMLMsubDataOnMod {
                return func() string {return fmt.Sprintf("throttle_%dms", d.Milliseconds())
                }
            }
            
        func(e *MathMLMSUBElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLMsubDataOnMod) *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-on-%s", key)
                
                customMods := lo.Map(modifiers, func(m MathMLMsubDataOnMod, i int) customDataKeyModifier  {
                    return customDataKeyModifier(m)
                })
                key = customDataKey(key, customMods...)
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLMSUBElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLMsubDataOnMod) *MathMLMSUBElement{
                if condition {
                    e.DATASTAR_ON(key,  expression,  modifiers...)
                }
                return e
            }

            // Remove the attribute DATASTAR_ON from the element.
            func(e *MathMLMSUBElement) DATASTAR_ONRemove() *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-on")
                return e
            }

        

    // Sets the focus of the element 
    
        func(e *MathMLMSUBElement) DATASTAR_FOCUSSet(b bool) *MathMLMSUBElement{
                key := "data-focus"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLMSUBElement) DATASTAR_FOCUS() *MathMLMSUBElement{
                return e.DATASTAR_FOCUSSet(true)
            }
        

    // Sets the header of for fetch requests 
    
        func(e *MathMLMSUBElement) DATASTAR_HEADER(key string, expression string) *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-header-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLMSUBElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLMSUBElement{
                if condition {
                    e.DATASTAR_HEADER(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_HEADER from the element.
            func(e *MathMLMSUBElement) DATASTAR_HEADERRemove() *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-header")
                return e
            }

        

    // Sets the URL for fetch requests 
    
        func(e *MathMLMSUBElement) DATASTAR_FETCH_URL(expression string) *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-fetch-url"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLMSUBElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLMSUBElement{
                if condition {
                    e.DATASTAR_FETCH_URL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_URL from the element.
            func(e *MathMLMSUBElement) DATASTAR_FETCH_URLRemove() *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-fetch-url")
                return e
            }

        

    // Sets the indicator selector for fetch requests 
    
        func(e *MathMLMSUBElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "DatastarFetchIndicator"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLMSUBElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLMSUBElement{
                if condition {
                    e.DATASTAR_FETCH_INDICATOR( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
            func(e *MathMLMSUBElement) DATASTAR_FETCH_INDICATORRemove() *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("DatastarFetchIndicator")
                return e
            }

        

    // Sets the visibility of the element 
    
        func(e *MathMLMSUBElement) DATASTAR_SHOWSet(b bool) *MathMLMSUBElement{
                key := "data-show"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLMSUBElement) DATASTAR_SHOW() *MathMLMSUBElement{
                return e.DATASTAR_SHOWSet(true)
            }
        

    // Triggers the callback when the element intersects the viewport 
    
        func(e *MathMLMSUBElement) DATASTAR_INTERSECTSSet(b bool) *MathMLMSUBElement{
                key := "data-intersects"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLMSUBElement) DATASTAR_INTERSECTS() *MathMLMSUBElement{
                return e.DATASTAR_INTERSECTSSet(true)
            }
        

    // Teleports the element to the given selector 
    
        func(e *MathMLMSUBElement) DATASTAR_TELEPORTSet(b bool) *MathMLMSUBElement{
                key := "data-teleport"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLMSUBElement) DATASTAR_TELEPORT() *MathMLMSUBElement{
                return e.DATASTAR_TELEPORTSet(true)
            }
        

    // Scrolls the element into view 
    
        func(e *MathMLMSUBElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLMSUBElement{
                key := "data-scroll-into-view"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLMSUBElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLMSUBElement{
                return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
            }
        

    // Setup the ViewTransitionAPI for the element 
    
        func(e *MathMLMSUBElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-view-transition-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLMSUBElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *MathMLMSUBElement{
                if condition {
                    e.DATASTAR_VIEW_TRANSITION(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
            func(e *MathMLMSUBElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLMSUBElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-view-transition")
                return e
            }

        



