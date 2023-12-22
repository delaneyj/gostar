package examples

import (
	"strings"
	"testing"

	. "github.com/delaneyj/gostar/elements/html"
	"github.com/delaneyj/gostar/elements/svg"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/bytebufferpool"
)

func TestExample1(t *testing.T) {
	type ExpectedResult struct {
		Expected string
		Actual   ElementRenderer
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
			Actual: svg.CLIPPATH(
				svg.RECT().CLASS("cls-1").ID("Rectangle_73").WIDTH(300).HEIGHT(300),
			).ID("clip-path"),
		},
		{
			Expected: `<linearGradient gradientUnits="objectBoundingBox" id="linear-gradient" x1="0.048" x2="0.963" y1="0.5" y2="0.5"><stop offset="0" stop-color="#000000"></stop><stop offset="1" stop-color="#0E67B4"></stop></linearGradient>`,
			Actual: svg.LINEARGRADIENT(
				svg.STOP().OFFSET(0).STOP_COLOR("#000000"),
				svg.STOP().OFFSET(1).STOP_COLOR("#0E67B4"),
			).
				ID("linear-gradient").
				GRADIENTUNITS("objectBoundingBox").
				X1(0.048).Y1(0.5).
				X2(0.963).Y2(0.5),
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

func TestExample2(t *testing.T) {

	x := DIV().STYLE("color", "rad").CLASS("foo", "aaaa").Text("bar")
	x.STYLEPairs("color", "rad", "font-size", "12px")
	x.STYLERemove("color")
	x.AUTOCAPITALIZE(DivAutocapitalize_off)

	sb := &strings.Builder{}
	HTML(
		BODY(
			DIV().CLASS("header").Text("Page Header"),
			x,
		),
	).Render(sb)
	expected := `<html><body><div class="header">Page Header</div><div autocapitalize="off" class="aaaa foo" style="font-size:12px">bar</div></body></html>`
	actual := sb.String()
	assert.Equal(t, expected, actual)
}
