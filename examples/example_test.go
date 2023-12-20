package examples

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/delaneyj/gostar/elements"
)

func TestExample1(t *testing.T) {
	type ExpectedResult struct {
		Expected string
		Actual   fmt.Stringer
	}

	expectedResults := []ExpectedResult{
		{
			Expected: `<div></div>`,
			Actual:   DIV(),
		},
		{
			Expected: `<div data-foo="bar"></div>`,
			Actual:   DIV().CustomData("foo", "bar"),
		},
		{
			Expected: `<div data-baz="qux" data-bind-foo="bar"></div>`,
			Actual:   DIV().CustomData("bind-foo", "bar").CustomData("baz", "qux"),
		},
		{
			Expected: `<nav class="navbar"><ol><li><a>Home</a></li><li><a>Contact</a></li><li><a>About</a></li></ol></nav>`,
			Actual: NAV(
				OL(
					LI(A().HREF("/").TEXT("Home")),
					LI(A().HREF("/contact").TEXT("Contact")),
					LI(A().HREF("/about").TEXT("About")),
				),
			).CLASS("navbar"),
		},
	}

	for _, expectedResult := range expectedResults {
		e := expectedResult.Expected
		a := expectedResult.Actual.String()
		assert.Equal(t, e, a)
	}
}
