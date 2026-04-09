package elements

import (
	"strings"
	"testing"
)

func TestGrouperRenderNilChildren(t *testing.T) {
	div := Group(
		Text("foo"),
		If(false, Text("nope")),
		Text("bar"),
	)
	var buf strings.Builder
	if err := div.Render(&buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := buf.String()
	if got != "foobar" {
		t.Errorf("expected %q, got %q", "foobar", got)
	}
}

func TestGrouperRenderNoPanicOnNil(t *testing.T) {
	div := Group(
		Text("foo"),
		If(false, Text("nope")),
		Text("bar"),
	)
	if panics(t, func() { div.Render(new(strings.Builder)) }) {
		t.Fatal("expected no panic, but panicked")
	}
}

func panics(t *testing.T, fn func()) (didPanic bool) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			didPanic = true
		}
	}()
	fn()
	return false
}
