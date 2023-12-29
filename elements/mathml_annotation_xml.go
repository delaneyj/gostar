// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml annotation-xml is generated from configuration file.
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
type MathMLANNOTATION_XMLElement struct {
    *Element
}

// Create a new MathMLANNOTATION_XMLElement element.
// This will create a new element with the tag
// "annotation-xml" during rendering.
func MathML_ANNOTATION_XML(children ...ElementRenderer) *MathMLANNOTATION_XMLElement {
    e := NewElement("annotation-xml", children...)
    e.IsSelfClosing = false
    e.Descendants = children

    return &MathMLANNOTATION_XMLElement{ Element: e }
}

func (e *MathMLANNOTATION_XMLElement) Children(children ...ElementRenderer) *MathMLANNOTATION_XMLElement {
    e.Descendants = append(e.Descendants, children...)
    return e
}

func(e *MathMLANNOTATION_XMLElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLANNOTATION_XMLElement {
    if condition {
        e.Descendants = append(e.Descendants, children...)
    }
    return e
}

func(e *MathMLANNOTATION_XMLElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLANNOTATION_XMLElement {
    if condition {
        e.Descendants = append(e.Descendants, trueChildren)
    } else {
        e.Descendants = append(e.Descendants, falseChildren)
    }
    return e
}

func (e *MathMLANNOTATION_XMLElement) Text(text string) *MathMLANNOTATION_XMLElement {
    e.Descendants = append(e.Descendants, Text(text))
    return e
}

func (e *MathMLANNOTATION_XMLElement) TextF(format string, args ...any) *MathMLANNOTATION_XMLElement {
    return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfText(condition bool, text string) *MathMLANNOTATION_XMLElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(text))
    }
    return e
}

func (e *MathMLANNOTATION_XMLElement) IfTextF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
    if condition {
        e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
    }
    return e
}

func (e *MathMLANNOTATION_XMLElement) Escaped(text string) *MathMLANNOTATION_XMLElement {
    e.Descendants = append(e.Descendants, Escaped(text))
    return e
}

func (e *MathMLANNOTATION_XMLElement) IfEscaped(condition bool, text string) *MathMLANNOTATION_XMLElement {
    if condition {
        e.Descendants = append(e.Descendants, Escaped(text))
    }
    return e
}

func (e *MathMLANNOTATION_XMLElement) EscapedF(format string, args ...any) *MathMLANNOTATION_XMLElement {
    return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfEscapedF(condition bool, format string, args ...any) *MathMLANNOTATION_XMLElement {
    if condition {
        e.Descendants = append(e.Descendants, EscapedF(format, args...))
    }
    return e
}

func (e *MathMLANNOTATION_XMLElement) CustomData(key, value string) *MathMLANNOTATION_XMLElement {
    if e.CustomDataAttributes == nil {
        e.CustomDataAttributes = treemap.New[string,string]()
    }
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLANNOTATION_XMLElement) IfCustomData(condition bool, key, value string) *MathMLANNOTATION_XMLElement {
    if condition {
        e.CustomData(key, value)
    }
    return e
}

func (e *MathMLANNOTATION_XMLElement) CustomDataF(key, format string, args ...any) *MathMLANNOTATION_XMLElement {
    return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLANNOTATION_XMLElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLANNOTATION_XMLElement {
    if condition {
        e.CustomData(key, fmt.Sprintf(format, args...))
    }
    return e
}

func (e *MathMLANNOTATION_XMLElement) CustomDataRemove(key string) *MathMLANNOTATION_XMLElement {
	if e.CustomDataAttributes == nil {
		return e
	}
    e.CustomDataAttributes.Del(key)
	return e
}


    // This attribute specifies the encoding used for the text content of the element 
// Possible values are text/plain, text/html, and application/x-tex. 
    func(e *MathMLANNOTATION_XMLElement) ENCODING(c MathMLAnnotationXmlEncodingChoice) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("encoding", string(c))
            return e
        }

        type MathMLAnnotationXmlEncodingChoice string
        const(
        
            MathMLAnnotationXmlEncoding_text_plain MathMLAnnotationXmlEncodingChoice = "text/plain"
        
            MathMLAnnotationXmlEncoding_text_html MathMLAnnotationXmlEncodingChoice = "text/html"
        
            MathMLAnnotationXmlEncoding_application_x_tex MathMLAnnotationXmlEncodingChoice = "application/x-tex"
        )

        // Remove the attribute ENCODING from the element.
        func(e *MathMLANNOTATION_XMLElement) ENCODINGRemove(c MathMLAnnotationXmlEncodingChoice) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("encoding")
            return e
        }
        

    // This attribute specifies the name of the annotation. 
    func(e *MathMLANNOTATION_XMLElement) NAME(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("name", s)
            return e
        }

        func(e *MathMLANNOTATION_XMLElement) IfNAME(condition bool, s string) *MathMLANNOTATION_XMLElement{
            if condition {
                e.NAME(s)
            }
            return e
        }

        // Remove the attribute NAME from the element.
        func(e *MathMLANNOTATION_XMLElement) NAMERemove(s string) *MathMLANNOTATION_XMLElement{
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
    func(e *MathMLANNOTATION_XMLElement) CLASS(s ...string) *MathMLANNOTATION_XMLElement{
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

        func(e *MathMLANNOTATION_XMLElement) IfCLASS(condition bool, s ...string) *MathMLANNOTATION_XMLElement{
            if condition {
                e.CLASS(s...)
            }
            return e
        }

        // Remove the attribute CLASS from the element.
        func(e *MathMLANNOTATION_XMLElement) CLASSRemove(s ...string) *MathMLANNOTATION_XMLElement{
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
    func(e *MathMLANNOTATION_XMLElement) DIR(c MathMLAnnotationXmlDirChoice) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("dir", string(c))
            return e
        }

        type MathMLAnnotationXmlDirChoice string
        const(
        // left-to-right 
            MathMLAnnotationXmlDir_ltr MathMLAnnotationXmlDirChoice = "ltr"
        // right-to-left 
            MathMLAnnotationXmlDir_rtl MathMLAnnotationXmlDirChoice = "rtl"
        )

        // Remove the attribute DIR from the element.
        func(e *MathMLANNOTATION_XMLElement) DIRRemove(c MathMLAnnotationXmlDirChoice) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("dir")
            return e
        }
        

    // This attribute specifies whether the element should be rendered using 
// displaystyle rules or not 
// Possible values are true and false. 
    func(e *MathMLANNOTATION_XMLElement) DISPLAYSTYLE(c MathMLAnnotationXmlDisplaystyleChoice) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("displaystyle", string(c))
            return e
        }

        type MathMLAnnotationXmlDisplaystyleChoice string
        const(
        // displaystyle rules 
            MathMLAnnotationXmlDisplaystyle_true MathMLAnnotationXmlDisplaystyleChoice = "true"
        // not displaystyle rules 
            MathMLAnnotationXmlDisplaystyle_false MathMLAnnotationXmlDisplaystyleChoice = "false"
        )

        // Remove the attribute DISPLAYSTYLE from the element.
        func(e *MathMLANNOTATION_XMLElement) DISPLAYSTYLERemove(c MathMLAnnotationXmlDisplaystyleChoice) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("displaystyle")
            return e
        }
        

    // This attribute assigns a name to an element 
// This name must be unique in a document. 
    func(e *MathMLANNOTATION_XMLElement) ID(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("id", s)
            return e
        }

        func(e *MathMLANNOTATION_XMLElement) IfID(condition bool, s string) *MathMLANNOTATION_XMLElement{
            if condition {
                e.ID(s)
            }
            return e
        }

        // Remove the attribute ID from the element.
        func(e *MathMLANNOTATION_XMLElement) IDRemove(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("id")
            return e
        }
    

    // This attribute specifies the background color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
    func(e *MathMLANNOTATION_XMLElement) MATHBACKGROUND(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("mathbackground", s)
            return e
        }

        func(e *MathMLANNOTATION_XMLElement) IfMATHBACKGROUND(condition bool, s string) *MathMLANNOTATION_XMLElement{
            if condition {
                e.MATHBACKGROUND(s)
            }
            return e
        }

        // Remove the attribute MATHBACKGROUND from the element.
        func(e *MathMLANNOTATION_XMLElement) MATHBACKGROUNDRemove(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("mathbackground")
            return e
        }
    

    // This attribute specifies the color of the element 
// Possible values are a color name or a color specification in the format defined 
// in the CSS3 Color Module [CSS3COLOR]. 
    func(e *MathMLANNOTATION_XMLElement) MATHCOLOR(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("mathcolor", s)
            return e
        }

        func(e *MathMLANNOTATION_XMLElement) IfMATHCOLOR(condition bool, s string) *MathMLANNOTATION_XMLElement{
            if condition {
                e.MATHCOLOR(s)
            }
            return e
        }

        // Remove the attribute MATHCOLOR from the element.
        func(e *MathMLANNOTATION_XMLElement) MATHCOLORRemove(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("mathcolor")
            return e
        }
    

    // This attribute specifies the size of the element 
// Possible values are a dimension or a dimensionless number. 
    func(e *MathMLANNOTATION_XMLElement) MATHSIZE_STR(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("mathsize", s)
            return e
        }

        func(e *MathMLANNOTATION_XMLElement) IfMATHSIZE_STR(condition bool, s string) *MathMLANNOTATION_XMLElement{
            if condition {
                e.MATHSIZE_STR(s)
            }
            return e
        }

        // Remove the attribute MATHSIZE_STR from the element.
        func(e *MathMLANNOTATION_XMLElement) MATHSIZE_STRRemove(s string) *MathMLANNOTATION_XMLElement{
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
    func(e *MathMLANNOTATION_XMLElement) NONCE(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                e.StringAttributes = treemap.New[string,string]()
            }
            e.StringAttributes.Set("nonce", s)
            return e
        }

        func(e *MathMLANNOTATION_XMLElement) IfNONCE(condition bool, s string) *MathMLANNOTATION_XMLElement{
            if condition {
                e.NONCE(s)
            }
            return e
        }

        // Remove the attribute NONCE from the element.
        func(e *MathMLANNOTATION_XMLElement) NONCERemove(s string) *MathMLANNOTATION_XMLElement{
            if e.StringAttributes == nil {
                return e
            }
            e.StringAttributes.Del("nonce")
            return e
        }
    

    // This attribute specifies the script level of the element 
// Possible values are an integer between 0 and 7, inclusive. 
    func(e *MathMLANNOTATION_XMLElement) SCRIPTLEVEL(i int) *MathMLANNOTATION_XMLElement{
            if e.IntAttributes == nil {
                e.IntAttributes = treemap.New[string,int]()
            }
            e.IntAttributes.Set("scriptlevel", i)
            return e
        }

        func (e *MathMLANNOTATION_XMLElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLANNOTATION_XMLElement {
            if condition {
                e.SCRIPTLEVEL(i)
            }
            return e
        }

        // Remove the attribute SCRIPTLEVEL from the element.
        func(e *MathMLANNOTATION_XMLElement) SCRIPTLEVELRemove(i int) *MathMLANNOTATION_XMLElement{
            if e.IntAttributes == nil {
                return e
            }
            e.IntAttributes.Del("scriptlevel")
            return e
        }
        

    // This attribute offers advisory information about the element for which it is 
// set. 
    func (e *MathMLANNOTATION_XMLElement) STYLEF(k string, format string, args ...any) *MathMLANNOTATION_XMLElement {
            return e.STYLE(k, fmt.Sprintf(format, args...))
        }

        func (e *MathMLANNOTATION_XMLElement) IfSTYLE(condition bool, k string, v string) *MathMLANNOTATION_XMLElement {
            if condition {
                e.STYLE(k, v)
            }
            return e
        }

        func (e *MathMLANNOTATION_XMLElement) STYLE(k string, v string) *MathMLANNOTATION_XMLElement {
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

        func (e *MathMLANNOTATION_XMLElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLANNOTATION_XMLElement {
            if condition {
                e.STYLE(k, fmt.Sprintf(format, args...))
            }
            return e
        }

        // Add the attributes in the map to the element.
        func (e *MathMLANNOTATION_XMLElement) STYLEMap(m map[string]string) *MathMLANNOTATION_XMLElement {
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
        func (e *MathMLANNOTATION_XMLElement) STYLEPairs(pairs ...string) *MathMLANNOTATION_XMLElement {
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

        func (e *MathMLANNOTATION_XMLElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLANNOTATION_XMLElement {
            if condition {
                e.STYLEPairs(pairs...)
            }
            return e
        }

        // Remove the attribute STYLE from the element.
        func (e *MathMLANNOTATION_XMLElement) STYLERemove(keys ...string) *MathMLANNOTATION_XMLElement {
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
    func(e *MathMLANNOTATION_XMLElement) TABINDEX(i int) *MathMLANNOTATION_XMLElement{
            if e.IntAttributes == nil {
                e.IntAttributes = treemap.New[string,int]()
            }
            e.IntAttributes.Set("tabindex", i)
            return e
        }

        func (e *MathMLANNOTATION_XMLElement) IfTABINDEX(condition bool, i int) *MathMLANNOTATION_XMLElement {
            if condition {
                e.TABINDEX(i)
            }
            return e
        }

        // Remove the attribute TABINDEX from the element.
        func(e *MathMLANNOTATION_XMLElement) TABINDEXRemove(i int) *MathMLANNOTATION_XMLElement{
            if e.IntAttributes == nil {
                return e
            }
            e.IntAttributes.Del("tabindex")
            return e
        }
        

    // Merges the store with the given object 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_MERGE_STORE(v any) *MathMLANNOTATION_XMLElement{
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
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_REF(expression string) *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-ref"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) IfDATASTAR_REF(condition bool, expression string) *MathMLANNOTATION_XMLElement{
                if condition {
                    e.DATASTAR_REF( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_REF from the element.
            func(e *MathMLANNOTATION_XMLElement) DATASTAR_REFRemove() *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-ref")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_BIND(key string, expression string) *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-bind-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLANNOTATION_XMLElement{
                if condition {
                    e.DATASTAR_BIND(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_BIND from the element.
            func(e *MathMLANNOTATION_XMLElement) DATASTAR_BINDRemove() *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-bind")
                return e
            }

        

    // Sets the value of the element 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_MODEL(expression string) *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-model"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLANNOTATION_XMLElement{
                if condition {
                    e.DATASTAR_MODEL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_MODEL from the element.
            func(e *MathMLANNOTATION_XMLElement) DATASTAR_MODELRemove() *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-model")
                return e
            }

        

    // Sets the textContent of the element 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_TEXT(expression string) *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-text"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLANNOTATION_XMLElement{
                if condition {
                    e.DATASTAR_TEXT( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_TEXT from the element.
            func(e *MathMLANNOTATION_XMLElement) DATASTAR_TEXTRemove() *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-text")
                return e
            }

        

    // Sets the event handler of the element 
    
        type MathMLAnnotationXmlDataOnMod customDataKeyModifier

            
            // Debounces the event handler 
            func MathMLAnnotationXmlDataOnModDebounce(
                    d time.Duration,
            ) MathMLAnnotationXmlDataOnMod {
                return func() string {return fmt.Sprintf("debounce_%dms", d.Milliseconds())
                }
            }
            
            // Throttles the event handler 
            func MathMLAnnotationXmlDataOnModThrottle(
                    d time.Duration,
            ) MathMLAnnotationXmlDataOnMod {
                return func() string {return fmt.Sprintf("throttle_%dms", d.Milliseconds())
                }
            }
            
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLAnnotationXmlDataOnMod) *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-on-%s", key)
                
                customMods := lo.Map(modifiers, func(m MathMLAnnotationXmlDataOnMod, i int) customDataKeyModifier  {
                    return customDataKeyModifier(m)
                })
                key = customDataKey(key, customMods...)
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLAnnotationXmlDataOnMod) *MathMLANNOTATION_XMLElement{
                if condition {
                    e.DATASTAR_ON(key,  expression,  modifiers...)
                }
                return e
            }

            // Remove the attribute DATASTAR_ON from the element.
            func(e *MathMLANNOTATION_XMLElement) DATASTAR_ONRemove() *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-on")
                return e
            }

        

    // Sets the focus of the element 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_FOCUSSet(b bool) *MathMLANNOTATION_XMLElement{
                key := "data-focus"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) DATASTAR_FOCUS() *MathMLANNOTATION_XMLElement{
                return e.DATASTAR_FOCUSSet(true)
            }
        

    // Sets the header of for fetch requests 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_HEADER(key string, expression string) *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-header-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLANNOTATION_XMLElement{
                if condition {
                    e.DATASTAR_HEADER(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_HEADER from the element.
            func(e *MathMLANNOTATION_XMLElement) DATASTAR_HEADERRemove() *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-header")
                return e
            }

        

    // Sets the URL for fetch requests 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_FETCH_URL(expression string) *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "data-fetch-url"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) IfDATASTAR_FETCH_URL(condition bool, expression string) *MathMLANNOTATION_XMLElement{
                if condition {
                    e.DATASTAR_FETCH_URL( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_URL from the element.
            func(e *MathMLANNOTATION_XMLElement) DATASTAR_FETCH_URLRemove() *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-fetch-url")
                return e
            }

        

    // Sets the indicator selector for fetch requests 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key := "DatastarFetchIndicator"
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLANNOTATION_XMLElement{
                if condition {
                    e.DATASTAR_FETCH_INDICATOR( expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
            func(e *MathMLANNOTATION_XMLElement) DATASTAR_FETCH_INDICATORRemove() *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("DatastarFetchIndicator")
                return e
            }

        

    // Sets the visibility of the element 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_SHOWSet(b bool) *MathMLANNOTATION_XMLElement{
                key := "data-show"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) DATASTAR_SHOW() *MathMLANNOTATION_XMLElement{
                return e.DATASTAR_SHOWSet(true)
            }
        

    // Triggers the callback when the element intersects the viewport 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_INTERSECTSSet(b bool) *MathMLANNOTATION_XMLElement{
                key := "data-intersects"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) DATASTAR_INTERSECTS() *MathMLANNOTATION_XMLElement{
                return e.DATASTAR_INTERSECTSSet(true)
            }
        

    // Teleports the element to the given selector 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_TELEPORTSet(b bool) *MathMLANNOTATION_XMLElement{
                key := "data-teleport"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) DATASTAR_TELEPORT() *MathMLANNOTATION_XMLElement{
                return e.DATASTAR_TELEPORTSet(true)
            }
        

    // Scrolls the element into view 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLANNOTATION_XMLElement{
                key := "data-scroll-into-view"
                e.BoolAttributes.Set(key, b)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLANNOTATION_XMLElement{
                return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
            }
        

    // Setup the ViewTransitionAPI for the element 
    
        func(e *MathMLANNOTATION_XMLElement) DATASTAR_VIEW_TRANSITION(key string, expression string) *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    e.StringAttributes = treemap.New[string,string]()
                }
                
                key = fmt.Sprintf("data-view-transition-%s", key)
                
                e.StringAttributes.Set(key, expression)
                return e
            }

            func(e *MathMLANNOTATION_XMLElement) IfDATASTAR_VIEW_TRANSITION(condition bool, key string, expression string) *MathMLANNOTATION_XMLElement{
                if condition {
                    e.DATASTAR_VIEW_TRANSITION(key,  expression, )
                }
                return e
            }

            // Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
            func(e *MathMLANNOTATION_XMLElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLANNOTATION_XMLElement{
                if e.StringAttributes == nil {
                    return e
                }
                e.StringAttributes.Del("data-view-transition")
                return e
            }

        



