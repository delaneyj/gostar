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
			Expected: "<div data-merge-store=\"{\"baz\":1234,\"foo\":\"bar\"}\"></div>",
			Actual: DIV().DATASTAR_MERGE_STORE(map[string]any{
				"foo": "bar",
				"baz": 1234,
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
