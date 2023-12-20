package examples

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	. "github.com/delaneyj/gostar/html"
)

func TestSlinso(t *testing.T) {
	type User struct {
		FirstName      string
		Email          string
		FavoriteColors []string
		RawContent     string
		EscapedContent string
	}

	type Navigation struct {
		Item string
		Link string
	}

	testComplexUser := &User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
		RawContent:     "<div><p>Raw Content to be displayed</p></div>",
		EscapedContent: "<div><div><div>Escaped</div></div></div>",
	}

	testComplexNav := []*Navigation{{
		Item: "Link 1",
		Link: "http://www.mytest.com/"}, {
		Item: "Link 2",
		Link: "http://www.mytest.com/"}, {
		Item: "Link 3",
		Link: "http://www.mytest.com/"},
	}
	testComplexTitle := testComplexUser.FirstName

	expectedtComplexResultMinified := `<!DOCTYPE html><html><body><header><title>Bob's Home Page</title><div class="header">Page Header</div></header><nav><ul class="navigation"><li><a href="http://www.mytest.com/">Link 1</a></li><li><a href="http://www.mytest.com/">Link 2</a></li><li><a href="http://www.mytest.com/">Link 3</a></li></ul></nav><section><div class="content"><div class="welcome"><h4>Hello Bob</h4><div class="raw"><div><p>Raw Content to be displayed</p></div></div><div class="enc">&lt;div&gt;&lt;div&gt;&lt;div&gt;Escaped&lt;/div&gt;&lt;/div&gt;&lt;/div&gt;</div></div><p>Bob has 1 message</p><p>Bob has 2 messages</p><p>Bob has 3 messages</p><p>Bob has 4 messages</p><p>Bob has 5 messages</p></div></section><footer><div class="footer">copyright 2016</div></footer></body></html>`

	header := func(title string) ElementBuilder {
		return HEADER(
			TITLE().RawF("%s's Home Page", title),
			DIV().CLASS("header").Text("Page Header"),
		)
	}

	navigation := func(nav []*Navigation) ElementBuilder {
		return NAV(
			UL(
				Range(nav, func(n *Navigation) ElementBuilder {
					return LI(
						A().HREF(n.Link).Text(n.Item),
					)
				}),
			).CLASS("navigation"),
		)
	}

	footer := func() ElementBuilder {
		return FOOTER(DIV().CLASS("footer").Text("copyright 2016"))
	}

	index := func(u *User, nav []*Navigation, title string) ElementBuilder {
		return Group(
			Raw("<!DOCTYPE html>"),
			HTML(
				BODY(
					header(title),
					navigation(nav),
					SECTION(
						DIV(
							DIV(
								H4().TextF("Hello %s", u.FirstName),
								DIV().CLASS("raw").Raw(u.RawContent),
								DIV().CLASS("enc").Text(u.EscapedContent),
							).CLASS("welcome"),
							Range(lo.Range(5), func(i int) ElementBuilder {
								count := i + 1
								return Tern(
									count == 1,
									P().TextF("%s has %d message", u.FirstName, count),
									P().TextF("%s has %d messages", u.FirstName, count),
								)
							}),
						).CLASS("content"),
					),
					footer(),
				),
			),
		)
	}

	sb := &strings.Builder{}
	builder := index(testComplexUser, testComplexNav, testComplexTitle)
	if err := builder.Build(sb); err != nil {
		t.Error(err)
	}

	assert.Equal(t, expectedtComplexResultMinified, sb.String())

}
