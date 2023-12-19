package elements

import (
	"slices"
	"strings"

	"github.com/samber/lo"
)

type Element struct {
	Tag                       string
	IsSelfClosing             bool
	BoolAttributes            map[string]struct{}
	DelimitedStringAttributes map[string]*DelimitedString
	DelimitedKVAttributes     map[string]*DelimitedKVString
	Children                  []*Element
}

func (e *Element) String() string {
	sb := &strings.Builder{}

	sb.WriteRune('<')
	sb.WriteString(e.Tag)

	values := lo.Entries(e.DelimitedStringAttributes)
	slices.SortFunc(values, func(a, b lo.Entry[string, *DelimitedString]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range values {
		sb.WriteRune(' ')
		sb.WriteString(entry.Key)
		sb.WriteRune('=')
		sb.WriteRune('"')
		sb.WriteString(entry.Value.String())
		sb.WriteRune('"')
	}

	kvValues := lo.Entries(e.DelimitedKVAttributes)
	slices.SortFunc(kvValues, func(a, b lo.Entry[string, *DelimitedKVString]) int {
		return strings.Compare(a.Key, b.Key)
	})
	for _, entry := range kvValues {
		sb.WriteRune(' ')
		sb.WriteString(entry.Key)
		sb.WriteRune('=')
		sb.WriteRune('"')
		sb.WriteString(entry.Value.String())
		sb.WriteRune('"')
	}

	boolStrings := lo.Keys(e.BoolAttributes)
	slices.Sort(boolStrings)
	for _, name := range boolStrings {
		sb.WriteRune(' ')
		sb.WriteString(name)
	}

	if e.IsSelfClosing {
		sb.WriteString(" />")
		return sb.String()
	}

	sb.WriteRune('>')
	for _, child := range e.Children {
		sb.WriteString(child.String())
	}
	sb.WriteString("</")
	sb.WriteString(e.Tag)
	sb.WriteRune('>')

	return sb.String()
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

func (ds *DelimitedString) String() string {
	values := lo.Keys(ds.values)
	slices.Sort(values)
	return strings.Join(values, ds.delimiter)
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

func (ds *DelimitedKVString) String() string {
	values := make([]string, 0, len(ds.values))
	for k, v := range ds.values {
		values = append(values, k+ds.associationDelimiter+v)
	}
	slices.Sort(values)
	return strings.Join(values, ds.delimiter)
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
