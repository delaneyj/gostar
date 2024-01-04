package tests

import (
	"testing"

	. "github.com/delaneyj/gostar/elements"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/bytebufferpool"
)

func TestDatastar(t *testing.T) {
	type ExpectedResult struct {
		Expected string
		Actual   ElementRenderer
	}

	expectedResults := []ExpectedResult{
		{
			Expected: "<div data-merge-store=\"{&#34;baz&#34;:1234,&#34;foo&#34;:&#34;bar&#34;}\"></div>",
			Actual: DIV().DATASTAR_MERGE_STORE(map[string]any{
				"foo": "bar",
				"baz": 1234,
			}),
		},
		{
			Expected: "<div data-merge-store=\"{&#34;label&#34;:&#34;HTML on whatever backend you like&#34;,&#34;v&#34;:1}\"></div>",
			Actual: DIV().DATASTAR_MERGE_STORE(map[string]any{
				"label": "HTML on whatever backend you like",
				"v":     1,
			}),
		},
		{
			Expected: "<div data-ref=\"foo\"></div>",
			Actual:   DIV().DATASTAR_REF("foo"),
		},
		{
			Expected: "<div data-bind-foo=\"bar\"></div>",
			Actual:   DIV().DATASTAR_BIND("foo", "bar"),
		},
	}

	for _, expectedResult := range expectedResults {
		buf := bytebufferpool.Get()
		e := expectedResult.Expected
		err := expectedResult.Actual.Render(buf)
		assert.NoError(t, err)
		a := buf.String()
		assert.Equal(t, e, a)
		bytebufferpool.Put(buf)
	}

}
