package html

import (
	"fmt"
	"html"
	"io"
	"slices"
	"strings"

	"github.com/samber/lo"
)

var (
	openBracket       = []byte("<")
	closeBracket      = []byte(">")
	spaceCloseBracket = []byte(" >")
	openSlash         = []byte("</")
	dataDash          = []byte(" data-")
	equalDblQuote     = []byte("=\"")
	dblQuote          = []byte("\"")
	space             = []byte(" ")
)

type ElementBuilder interface {
	Build(w io.Writer) error
}

type Element struct {
	Tag                       []byte
	IsSelfClosing             bool
	BoolAttributes            map[string]struct{}
	StringAttributes          map[string]string
	IntAttributes             map[string]int
	DelimitedStringAttributes map[string]*DelimitedString
	DelimitedKVAttributes     map[string]*DelimitedKVString
	CustomDataAttributes      map[string]string
	Descendants               []ElementBuilder
}

func (e *Element) Build(w io.Writer) error {
	w.Write(openBracket)
	w.Write(e.Tag)

	dataValues := lo.Entries(e.CustomDataAttributes)
	slices.SortFunc(dataValues, func(a, b lo.Entry[string, string]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range dataValues {
		w.Write(dataDash)
		w.Write([]byte(entry.Key))
		w.Write(equalDblQuote)
		w.Write([]byte(entry.Value))
		w.Write(dblQuote)
	}

	stringValues := lo.Entries(e.StringAttributes)
	slices.SortFunc(stringValues, func(a, b lo.Entry[string, string]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range stringValues {
		w.Write(space)
		w.Write([]byte(entry.Key))
		w.Write(equalDblQuote)
		w.Write([]byte(entry.Value))
		w.Write(dblQuote)
	}

	values := lo.Entries(e.DelimitedStringAttributes)
	slices.SortFunc(values, func(a, b lo.Entry[string, *DelimitedString]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range values {
		w.Write(space)
		w.Write([]byte(entry.Key))
		w.Write(equalDblQuote)
		if err := entry.Value.Build(w); err != nil {
			return fmt.Errorf("failed to build element: %w", err)
		}
		w.Write(dblQuote)
	}

	kvValues := lo.Entries(e.DelimitedKVAttributes)
	slices.SortFunc(kvValues, func(a, b lo.Entry[string, *DelimitedKVString]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range kvValues {
		w.Write(space)
		w.Write([]byte(entry.Key))
		w.Write(equalDblQuote)
		if err := entry.Value.Build(w); err != nil {
			return fmt.Errorf("failed to build element: %w", err)
		}
		w.Write(dblQuote)
	}

	boolStrings := lo.Keys(e.BoolAttributes)
	slices.Sort(boolStrings)
	for _, name := range boolStrings {
		w.Write(space)
		w.Write([]byte(name))
	}

	if e.IsSelfClosing {
		w.Write(spaceCloseBracket)
		return nil
	}

	w.Write(closeBracket)
	for _, child := range e.Descendants {
		if err := child.Build(w); err != nil {
			return fmt.Errorf("failed to build element: %w", err)
		}
	}

	w.Write(openSlash)
	w.Write(e.Tag)
	w.Write(closeBracket)

	return nil
}

type DelimitedString struct {
	delimiter string
	values    map[string]struct{}
}

func (ds *DelimitedString) Add(values ...string) {
	for _, value := range values {
		ds.values[value] = struct{}{}
	}
}

func (ds *DelimitedString) Remove(values ...string) {
	for _, value := range values {
		delete(ds.values, value)
	}
}

func (ds *DelimitedString) Build(w io.Writer) error {
	values := lo.Keys(ds.values)
	slices.Sort(values)
	_, err := w.Write([]byte(strings.Join(values, ds.delimiter)))
	return err
}

func NewDelimitedString(delimiter string) *DelimitedString {
	return &DelimitedString{
		delimiter: delimiter,
		values:    make(map[string]struct{}),
	}
}

func NewSpaceDelimitedString() *DelimitedString {
	return NewDelimitedString(" ")
}

type DelimitedKVString struct {
	associationDelimiter string
	delimiter            string
	values               map[string]string
}

func (ds *DelimitedKVString) Add(key, value string) {
	ds.values[key] = value
}

func (ds *DelimitedKVString) Remove(key string) {
	delete(ds.values, key)
}

func (ds *DelimitedKVString) Build(w io.Writer) error {
	values := make([]string, 0, len(ds.values))
	for k, v := range ds.values {
		values = append(values, k+ds.associationDelimiter+v)
	}
	slices.Sort(values)
	w.Write([]byte(strings.Join(values, ds.delimiter)))
	return nil
}

func NewDelimitedKVString(associationDelimiter, delimiter string) *DelimitedKVString {
	return &DelimitedKVString{
		delimiter:            delimiter,
		associationDelimiter: associationDelimiter,
		values:               make(map[string]string),
	}
}

func NewEqualSemicolonDelimitedKVString() *DelimitedKVString {
	return NewDelimitedKVString("=", ";")
}

type TextContent string

func (tc *TextContent) Build(w io.Writer) error {
	_, err := w.Write([]byte(*tc))
	return err
}

func Text(text string) *TextContent {
	return (*TextContent)(&text)
}

type EscapedContent string

func (ec *EscapedContent) Build(w io.Writer) error {
	_, err := w.Write([]byte(html.EscapeString(string(*ec))))
	return err
}

func Escaped(text string) *EscapedContent {
	return (*EscapedContent)(&text)
}

type Grouper struct {
	Children []ElementBuilder
}

func (g *Grouper) Build(w io.Writer) error {
	for _, child := range g.Children {
		if err := child.Build(w); err != nil {
			return fmt.Errorf("failed to build element: %w", err)
		}
	}
	return nil
}

func Group(children ...ElementBuilder) *Grouper {
	return &Grouper{
		Children: children,
	}
}

func If(condition bool, children ...ElementBuilder) ElementBuilder {
	if condition {
		return Group(children...)
	}
	return nil
}

func Tern(condition bool, trueChildren, falseChildren ElementBuilder) ElementBuilder {
	if condition {
		return trueChildren
	}
	return falseChildren
}

func Range[T any](values []T, cb func(T) ElementBuilder) ElementBuilder {
	children := make([]ElementBuilder, 0, len(values))
	for _, value := range values {
		children = append(children, cb(value))
	}
	return Group(children...)
}

func NewElement(tag string, children ...ElementBuilder) *Element {
	return &Element{
		Tag:         []byte(tag),
		Descendants: children,
	}
}

func (e *Element) Attr(name string, value string) *Element {
	return e.Attrs(name, value)
}

func (e *Element) Attrs(attrs ...string) *Element {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes[k] = v
	}
	return e
}

func (e *Element) AttrsMap(attrs map[string]string) *Element {
	if e.StringAttributes == nil {
		e.StringAttributes = map[string]string{}
	}
	for k, v := range attrs {
		e.StringAttributes[k] = v
	}
	return e
}
