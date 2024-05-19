package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/delaneyj/gostar/elements"
	"github.com/samber/lo"
)

type (
	Navigation struct {
		Link string
		Item string
	}

	User struct {
		FirstName      string
		RawContent     string
		EscapedContent string
	}
)

var (
	header = func(title string) ElementRenderer {
		return HEADER(
			TITLE().TextF("%s's Home Page", title),
			DIV().CLASS("header").Text("Page Header"),
		)
	}

	navigation = func(nav []*Navigation) ElementRenderer {
		return NAV(
			UL(
				Range(nav, func(n *Navigation) ElementRenderer {
					return LI(
						A().HREF(n.Link).Text(n.Item),
					)
				}),
			).CLASS("navigation"),
		)
	}

	footer = func() ElementRenderer {
		return FOOTER(DIV().CLASS("footer").Text("copyright 2016"))
	}

	index = func(u *User, nav []*Navigation, title string) ElementRenderer {
		return Group(
			Text("<!DOCTYPE html>"),
			HTML(
				BODY(
					header(title),
					navigation(nav),
					SECTION(
						DIV(
							DIV(
								H4().TextF("Hello %s", u.FirstName),
								DIV().CLASS("raw").Text(u.RawContent),
								DIV().CLASS("enc").Escaped(u.EscapedContent),
							).CLASS("welcome"),
							Range(lo.Range(5), func(i int) ElementRenderer {
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
)

// An index handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{
		FirstName:      "John",
		RawContent:     `<a href="http://john.com">This is some raw content</a>`,
		EscapedContent: `&lt;a href=&#34;http://john.com&#34;&gt;This is some encoded content&lt;/a&gt;`,
	}

	nav := []*Navigation{
		{Link: "/home", Item: "Home"},
		{Link: "/about", Item: "About"},
	}

	index(user, nav, "John").Render(w)
}

func TestIndex(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(indexHandler))
	defer s.Close()

	resp, err := http.Get(s.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	c, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"Page Header",
		"Hello John",
		"This is some raw content",
		"&amp;lt;a href=&amp;#34;http://john.com&amp;#34;&amp;gt;This is some encoded content&amp;lt;/a&amp;gt;",
	}

	for _, e := range expected {
		if !strings.Contains(string(c), e) {
			t.Errorf("expected '%s'", e)
		}
	}
}
