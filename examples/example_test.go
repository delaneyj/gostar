package examples

import (
	"testing"

	"github.com/delaneyj/gostar/html"
	. "github.com/delaneyj/gostar/html"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/bytebufferpool"
)

func TestExample1(t *testing.T) {
	type ExpectedResult struct {
		Expected string
		Actual   html.ElementBuilder
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
			Expected: `<nav class="navbar"><ol><li><a href="/">Home</a></li><li><a href="/contact">Contact</a></li><li><a href="/about">About</a></li></ol></nav>`,
			Actual: NAV(
				OL(
					LI(A().HREF("/").Text("Home")),
					LI(A().HREF("/contact").Text("Contact")),
					LI(A().HREF("/about").Text("About")),
				),
			).CLASS("navbar"),
		},
		{
			Expected: `<nav class="navbar"><ol><li><a href="/">Home</a></li><li><a href="/contact">Contact</a></li><li><a href="/about">About</a></li></ol></nav>`,
			Actual: NAV().CLASS("navbar").Children(
				OL(
					LI(A().HREF("/").Text("Home")),
					LI(A().HREF("/contact").Text("Contact")),
					LI(A().HREF("/about").Text("About")),
				),
			),
		},
		{
			Expected: `<clipPath id="clip-path"><rect class="cls-1" height="300" id="Rectangle_73" width="300"></rect></clipPath>`,
			Actual: NewElement("clipPath",
				NewElement("rect").Attrs("class", "cls-1", "id", "Rectangle_73", "width", "300", "height", "300"),
			).Attrs("id", "clip-path"),
		},
		{
			Expected: `<linearGradient gradientUnits="objectBoundingBox" id="linear-gradient" x1="0.048" x2="0.963" y1="0.5" y2="0.5"><stop offset="0" stop-color="#000000"></stop><stop offset="1" stop-color="#0E67B4"></stop></linearGradient>`,
			Actual: NewElement("linearGradient",
				NewElement("stop").Attrs("offset", "0", "stop-color", "#000000"),
				NewElement("stop").Attrs("offset", "1", "stop-color", "#0E67B4"),
			).
				Attrs("id", "linear-gradient", "gradientUnits", "objectBoundingBox").
				Attrs("x1", "0.048", "y1", "0.5", "x2", "0.963", "y2", "0.5"),
		},
	}

	for _, expectedResult := range expectedResults {
		buf := bytebufferpool.Get()
		defer bytebufferpool.Put(buf)

		e := expectedResult.Expected
		err := expectedResult.Actual.Build(buf)
		assert.NoError(t, err)
		a := buf.String()
		assert.Equal(t, e, a)
	}
}
