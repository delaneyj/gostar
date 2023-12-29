// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml annotation is generated from configuration file.
// Description:
package elements

import(
    "fmt"
    "time"
    "github.com/igrmk/treemap/v2"
    "github.com/goccy/go-json"
    "github.com/samber/lo"
)

// This element is used to include comments or annotations within a MathML 
// expression 
// It can be used to provide additional information about the expression, or to 
// include comments for the author of the expression. 
type MathMLANNOTATIONElement struct {
    *Element
}

// Create a new MathMLANNOTATIONElement element.
// This will create a new element with the tag
// "annotation" during rendering.
func MathML_ANNOTATION(children ...ElementRenderer) *MathMLANNOTATIONElement {
    e := NewElement("annotation", children...)
    e.IsSelfClosing = false
    e.Descendants = children

    return &MathMLANNOTATIONElement{ Element: e }
}

func (e *MathMLANNOTATIONElement) Children(children ...ElementRenderer) *MathMLANNOTATIONElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *MathMLANNOTATIONElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLANNOTATIONElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *MathMLANNOTATIONElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLANNOTATIONElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *MathMLANNOTATIONElement) Text(text string) *MathMLANNOTATIONElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *MathMLANNOTATIONElement) TextF(format string, args ...any) *MathMLANNOTATIONElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATIONElement) IfText(condition bool, text string) *MathMLANNOTATIONElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(text))
    }
    return e
}

func (e *MathMLANNOTATIONElement) IfTextF(condition bool, format string, args ...any) *MathMLANNOTATIONElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
    }
    return e
}

func (e *MathMLANNOTATIONElement) Escaped(text string) *MathMLANNOTATIONElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *MathMLANNOTATIONElement) IfEscaped(condition bool, text string) *MathMLANNOTATIONElement {
    if condition {
        e.Descendants = append(e.Descendants, Escaped(text))
    }
    return e
}

func (e *MathMLANNOTATIONElement) EscapedF(format string, args ...any) *MathMLANNOTATIONElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATIONElement) IfEscapedF(condition bool, format string, args ...any) *MathMLANNOTATIONElement {
    if condition {
        e.Descendants = append(e.Descendants, EscapedF(format, args...))
    }
    return e
}

func (e *MathMLANNOTATIONElement) CustomData(key, value string) *MathMLANNOTATIONElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLANNOTATIONElement) IfCustomData(condition bool, key, value string) *MathMLANNOTATIONElement {
    if condition {
        e.CustomData(key, value)
    }
    return e
}

func (e *MathMLANNOTATIONElement) CustomDataF(key, format string, args ...any) *MathMLANNOTATIONElement {
    return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATIONElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLANNOTATIONElement {
    if condition {
        e.CustomData(key, fmt.Sprintf(format, args...))
    }
    return e
}

func (e *MathMLANNOTATIONElement) CustomDataRemove(key string) *MathMLANNOTATIONElement {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


    // This attribute specifies the encoding used for the text content of the element 
// Possible values are text/plain, text/html, and application/x-tex. 
    func(e *MathMLANNOTATIONElement) ENCODING(c MathMLAnnotationEncodingChoice) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("encoding", string(c))
            return e
        }

        type MathMLAnnotationEncodingChoice string
        const(
        
            MathMLAnnotationEncoding_text_plain MathMLAnnotationEncodingChoice = "text/plain"
        
            MathMLAnnotationEncoding_text_html MathMLAnnotationEncodingChoice = "text/html"
        
            MathMLAnnotationEncoding_application_x_tex MathMLAnnotationEncodingChoice = "application/x-tex"
        )

        // Remove the attribute ENCODING from the element.
        func(e *MathMLANNOTATIONElement) ENCODINGRemove(c MathMLAnnotationEncodingChoice) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("encoding")
            return e
        }
        

    // This attribute specifies the name of the annotation. 
    func(e *MathMLANNOTATIONElement) NAME(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("name", s)
            return e
        }

        func(e *MathMLANNOTATIONElement) IfNAME(condition bool, s string) *MathMLANNOTATIONElement{
            if condition {
                e.NAME(s)
            }
            return e
        }

        // Remove the attribute NAME from the element.
        func(e *MathMLANNOTATIONElement) NAMERemove(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("name")
            return e
        }
    

    // Assigns a class name or set of class names to an element 
// You may assign the same class name or names to any number of elements 
// If you specify multiple class names, they must be separated by whitespace 
// characters. 
    func(e *MathMLANNOTATIONElement) CLASS(s ...string) *MathMLANNOTATIONElement{
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

        func(e *MathMLANNOTATIONElement) IfCLASS(condition bool, s ...string) *MathMLANNOTATIONElement{
            if condition {
                e.CLASS(s...)
            }
            return e
        }

        // Remove the attribute CLASS from the element.
        func(e *MathMLANNOTATIONElement) CLASSRemove(s ...string) *MathMLANNOTATIONElement{
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
    func(e *MathMLANNOTATIONElement) DIR(c MathMLAnnotationDirChoice) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("dir", string(c))
            return e
        }

        type MathMLAnnotationDirChoice string
        const(
        // left-to-right 
            MathMLAnnotationDir_ltr MathMLAnnotationDirChoice = "ltr"
        // right-to-left 
            MathMLAnnotationDir_rtl MathMLAnnotationDirChoice = "rtl"
        )

        // Remove the attribute DIR from the element.
        func(e *MathMLANNOTATIONElement) DIRRemove(c MathMLAnnotationDirChoice) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("dir")
            return e
        }
        

    // This attribute specifies whether the element should be rendered using 
// displaystyle rules or not 
// Possible values are true and false. 
    func(e *MathMLANNOTATIONElement) DISPLAYSTYLE(c MathMLAnnotationDisplaystyleChoice) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("displaystyle", string(c))
            return e
        }

        type MathMLAnnotationDisplaystyleChoice string
        const(
        // displaystyle rules 
            MathMLAnnotationDisplaystyle_true MathMLAnnotationDisplaystyleChoice = "true"
        // not displaystyle rules 
            MathMLAnnotationDisplaystyle_false MathMLAnnotationDisplaystyleChoice = "false"
        )

        // Remove the attribute DISPLAYSTYLE from the element.
        func(e *MathMLANNOTATIONElement) DISPLAYSTYLERemove(c MathMLAnnotationDisplaystyleChoice) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("displaystyle")
            return e
        }
        

    // This attribute assigns a name to an element 
// This name must be unique in a document. 
    func(e *MathMLANNOTATIONElement) ID(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("id", s)
            return e
        }

        func(e *MathMLANNOTATIONElement) IfID(condition bool, s string) *MathMLANNOTATIONElement{
            if condition {
                e.ID(s)
            }
            return e
        }

        // Remove the attribute ID from the element.
        func(e *MathMLANNOTATIONElement) IDRemove(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("id")
            return e
        }
    

    // This attribute specifies the background color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
    func(e *MathMLANNOTATIONElement) MATHBACKGROUND(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("mathbackground", s)
            return e
        }

        func(e *MathMLANNOTATIONElement) IfMATHBACKGROUND(condition bool, s string) *MathMLANNOTATIONElement{
            if condition {
                e.MATHBACKGROUND(s)
            }
            return e
        }

        // Remove the attribute MATHBACKGROUND from the element.
        func(e *MathMLANNOTATIONElement) MATHBACKGROUNDRemove(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("mathbackground")
            return e
        }
    

    // This attribute specifies the color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
    func(e *MathMLANNOTATIONElement) MATHCOLOR(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("mathcolor", s)
            return e
        }

        func(e *MathMLANNOTATIONElement) IfMATHCOLOR(condition bool, s string) *MathMLANNOTATIONElement{
            if condition {
                e.MATHCOLOR(s)
            }
            return e
        }

        // Remove the attribute MATHCOLOR from the element.
        func(e *MathMLANNOTATIONElement) MATHCOLORRemove(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("mathcolor")
            return e
        }
    

    // This attribute specifies the size of the element 
// Possible values are a dimension or a dimensionless number. 
    func(e *MathMLANNOTATIONElement) MATHSIZE_STR(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("mathsize", s)
            return e
        }

        func(e *MathMLANNOTATIONElement) IfMATHSIZE_STR(condition bool, s string) *MathMLANNOTATIONElement{
            if condition {
                e.MATHSIZE_STR(s)
            }
            return e
        }

        // Remove the attribute MATHSIZE_STR from the element.
        func(e *MathMLANNOTATIONElement) MATHSIZE_STRRemove(s string) *MathMLANNOTATIONElement{
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
    func(e *MathMLANNOTATIONElement) NONCE(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("nonce", s)
            return e
        }

        func(e *MathMLANNOTATIONElement) IfNONCE(condition bool, s string) *MathMLANNOTATIONElement{
            if condition {
                e.NONCE(s)
            }
            return e
        }

        // Remove the attribute NONCE from the element.
        func(e *MathMLANNOTATIONElement) NONCERemove(s string) *MathMLANNOTATIONElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("nonce")
            return e
        }
    

    // This attribute specifies the script level of the element 
// Possible values are an integer between 0 and 7, inclusive. 
    func(e *MathMLANNOTATIONElement) SCRIPTLEVEL(i int) *MathMLANNOTATIONElement{
            if e.IntAttributes == nil {
                e.IntAttributes = treemap.New[string,int]()
            }
            e.IntAttributes.Set("scriptlevel", i)
            return e
        }

        func (e *MathMLANNOTATIONElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLANNOTATIONElement {
            if condition {
                e.SCRIPTLEVEL(i)
            }
            return e
        }

        // Remove the attribute SCRIPTLEVEL from the element.
        func(e *MathMLANNOTATIONElement) SCRIPTLEVELRemove(i int) *MathMLANNOTATIONElement{
            if e.IntAttributes == nil {
                return e
            }
            e.IntAttributes.Del("scriptlevel")
            return e
        }
        

    // This attribute offers advisory information about the element for which it is 
// set. 
    func (e *MathMLANNOTATIONElement) STYLEF(k string, format string, args ...any) *MathMLANNOTATIONElement {
            return e.STYLE(k, fmt.Sprintf(format, args...))
        }

        func (e *MathMLANNOTATIONElement) IfSTYLE(condition bool, k string, v string) *MathMLANNOTATIONElement {
            if condition {
                e.STYLE(k, v)
            }
            return e
        }

        func (e *MathMLANNOTATIONElement) STYLE(k string, v string) *MathMLANNOTATIONElement {
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

        func (e *MathMLANNOTATIONElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLANNOTATIONElement {
            if condition {
                e.STYLE(k, fmt.Sprintf(format, args...))
            }
            return e
        }

        // Add the attributes in the map to the element.
        func (e *MathMLANNOTATIONElement) STYLEMap(m map[string]string) *MathMLANNOTATIONElement {
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
        func (e *MathMLANNOTATIONElement) STYLEPairs(pairs ...string) *MathMLANNOTATIONElement {
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

        func (e *MathMLANNOTATIONElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLANNOTATIONElement {
            if condition {
                e.STYLEPairs(pairs...)
            }
            return e
        }

        // Remove the attribute STYLE from the element.
        func (e *MathMLANNOTATIONElement) STYLERemove(keys ...string) *MathMLANNOTATIONElement {
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
    func(e *MathMLANNOTATIONElement) TABINDEX(i int) *MathMLANNOTATIONElement{
            if e.IntAttributes == nil {
                e.IntAttributes = treemap.New[string,int]()
            }
            e.IntAttributes.Set("tabindex", i)
            return e
        }

        func (e *MathMLANNOTATIONElement) IfTABINDEX(condition bool, i int) *MathMLANNOTATIONElement {
            if condition {
                e.TABINDEX(i)
            }
            return e
        }

        // Remove the attribute TABINDEX from the element.
        func(e *MathMLANNOTATIONElement) TABINDEXRemove(i int) *MathMLANNOTATIONElement{
            if e.IntAttributes == nil {
                return e
            }
            e.IntAttributes.Del("tabindex")
            return e
        }
        

    // Merges the store with the given object 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_MERGE_STORE(v any) *MathMLANNOTATIONElement{
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
    
        func(e *MathMLANNOTATIONElement) DATASTAR_REF(expression string) *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-ref"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATIONElement) IfDATASTAR_REF(condition bool, expression string) *MathMLANNOTATIONElement{
                if condition {
                    e.DATASTAR_REF( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_REF from the element.
            func(e *MathMLANNOTATIONElement) DATASTAR_REFRemove() *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-ref")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_BIND(key string, expression string) *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-bind-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATIONElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLANNOTATIONElement{
                if condition {
                    e.DATASTAR_BIND(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_BIND from the element.
            func(e *MathMLANNOTATIONElement) DATASTAR_BINDRemove() *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-bind")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_MODEL(expression string) *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-model"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATIONElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLANNOTATIONElement{
                if condition {
                    e.DATASTAR_MODEL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_MODEL from the element.
            func(e *MathMLANNOTATIONElement) DATASTAR_MODELRemove() *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-model")
                return e
            }

        

    // Sets the textContent of the element 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_TEXT(expression string) *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-text"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATIONElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLANNOTATIONElement{
                if condition {
                    e.DATASTAR_TEXT( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_TEXT from the element.
            func(e *MathMLANNOTATIONElement) DATASTAR_TEXTRemove() *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-text")
                return e
            }

        

    // Sets the event handler of the element 
    
        type MathMLAnnotationDataOnMod customDataKeyModifier

            
            // Debounces the event handler 
            func MathMLAnnotationDataOnModDebounce(
                    d time.Duration,
            ) MathMLAnnotationDataOnMod {
                return func() string {return fmt.Sprintf("debounce_%dms", d.Milliseconds())
                }
            }
            
            // Throttles the event handler 
            func MathMLAnnotationDataOnModThrottle(
                    d time.Duration,
            ) MathMLAnnotationDataOnMod {
                return func() string {return fmt.Sprintf("throttle_%dms", d.Milliseconds())
                }
            }
            
        func(e *MathMLANNOTATIONElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLAnnotationDataOnMod) *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-on-%s", key)
                
                customMods := lo.Map(modifiers, func(m MathMLAnnotationDataOnMod, i int) customDataKeyModifier  {
                    return customDataKeyModifier(m)
                })
                key = customDataKey(key, customMods...)
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATIONElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLAnnotationDataOnMod) *MathMLANNOTATIONElement{
                if condition {
                    e.DATASTAR_ON(key,  expression,  modifiers...)
                }
                return e
            }

            // Remove the attribute DATASTAR_ON from the element.
            func(e *MathMLANNOTATIONElement) DATASTAR_ONRemove() *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-on")
                return e
            }

        

    // Sets the focus of the element 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_FOCUSSet(b bool) *MathMLANNOTATIONElement{
                key := "data-focus"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATIONElement) DATASTAR_FOCUS() *MathMLANNOTATIONElement{
                return e.DATASTAR_FOCUSSet(true)
            }
        

    // Sets the header of for fetch requests 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_HEADER(key string, expression string) *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-header-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATIONElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLANNOTATIONElement{
                if condition {
                    e.DATASTAR_HEADER(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_HEADER from the element.
            func(e *MathMLANNOTATIONElement) DATASTAR_HEADERRemove() *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-header")
                return e
            }

        

    // Sets the URL for fetch requests 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_FETCH_URL(expression string) *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-fetch-url"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATIONElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLANNOTATIONElement{
                if condition {
                    e.DATASTAR_FETCH_URL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_URL from the element.
            func(e *MathMLANNOTATIONElement) DATASTAR_FETCH_URLRemove() *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-fetch-url")
                return e
            }

        

    // Sets the indicator selector for fetch requests 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "DatastarFetchIndicator"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATIONElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLANNOTATIONElement{
                if condition {
                    e.DATASTAR_FETCH_INDICATOR( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
            func(e *MathMLANNOTATIONElement) DATASTAR_FETCH_INDICATORRemove() *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("DatastarFetchIndicator")
                return e
            }

        

    // Sets the visibility of the element 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_SHOWSet(b bool) *MathMLANNOTATIONElement{
                key := "data-show"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATIONElement) DATASTAR_SHOW() *MathMLANNOTATIONElement{
                return e.DATASTAR_SHOWSet(true)
            }
        

    // Triggers the callback when the element intersects the viewport 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_INTERSECTSSet(b bool) *MathMLANNOTATIONElement{
                key := "data-intersects"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATIONElement) DATASTAR_INTERSECTS() *MathMLANNOTATIONElement{
                return e.DATASTAR_INTERSECTSSet(true)
            }
        

    // Teleports the element to the given selector 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_TELEPORTSet(b bool) *MathMLANNOTATIONElement{
                key := "data-teleport"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATIONElement) DATASTAR_TELEPORT() *MathMLANNOTATIONElement{
                return e.DATASTAR_TELEPORTSet(true)
            }
        

    // Scrolls the element into view 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLANNOTATIONElement{
                key := "data-scroll-into-view"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATIONElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLANNOTATIONElement{
                return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
            }
        

    // Setup the ViewTransitionAPI for the element 
    
        func(e *MathMLANNOTATIONElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-view-transition-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATIONElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *MathMLANNOTATIONElement{
                if condition {
                    e.DATASTAR_VIEW_TRANSITION(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
            func(e *MathMLANNOTATIONElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLANNOTATIONElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-view-transition")
                return e
            }

        



