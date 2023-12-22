[![GoDoc](https://godoc.org/github.com/delaneyj/gostar?status.svg)](http://godoc.org/github.com/delaneyj/gostar)
[![Go Report Card](https://goreportcard.com/badge/github.com/delaneyj/gostar)](https://goreportcard.com/report/github.com/delaneyj/gostar)

# Gostar

A fluent HTML builder for Go built directly from the [HTML Living Standard](https://html.spec.whatwg.org/).

![mascot](docs/mascot.png)

## What does it do?

Instead of creating HTML using string concatenation or templates, Gostar allows you to create HTML using a fluent API. This allows you to create HTML in a more natural way, and makes it easier to create HTML programmatically.  It does that by building directly from the [HTML Living Standard](https://html.spec.whatwg.org/) and introspecting to creating automated conversions for all HTML elements and attributes.



Just import the package.  I prefer dot imports so I don't have to type the package name over and over.  All HTML tags and attributes are uppercase to avoid collisions with Go keywords and make explicit what is a DSL.


```go
. "github.com/delaneyj/gostar/elements/html"
```
Since everything is `just` fluent functions you can compose them together to create more complex HTML structures.  You can also create your own fluent functions to create reusable HTML components.

```go
header := func(title string) ElementBuilder {
    return HEADER(
        TITLE().TextF("%s's Home Page", title),
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
```

## Performance

Testing shows it's faster than Go builtin template/html package but this is including creating and writing to a buffer.  It's still sub-microsecond for most pages and the convenience of the fluent API is worth it.  Avoiding the parse, setup template context, then execute loop mean way less work for the developer.
