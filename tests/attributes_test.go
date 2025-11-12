package tests

import (
	"testing"

	. "github.com/delaneyj/gostar/elements"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/bytebufferpool"
)

func TestAttributes(t *testing.T) {
	type ExpectedResult struct {
		Expected string
		Actual   ElementRenderer
	}

	expectedResults := []ExpectedResult{
		{
			Expected: `<fieldset disabled></fieldset>`,
			Actual:   FIELDSET().DISABLED(),
		},
		{
			Expected: `<fieldset form="form"></fieldset>`,
			Actual:   FIELDSET().FORM("form"),
		},
		{
			Expected: `<fieldset name="fieldset"></fieldset>`,
			Actual:   FIELDSET().NAME("fieldset"),
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
