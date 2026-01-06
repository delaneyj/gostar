package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/bytebufferpool"

	. "github.com/delaneyj/gostar/elements"
)

func TestGrouper_Render(t *testing.T) {
	group := Group(
		Text("foo"),
		If(false, Text("nope")),
		Text("bar"),
	)

	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	var err error
	require.NotPanics(t, func() {
		err = group.Render(buf)
	})
	require.NoError(t, err)
	assert.Equal(t, "foobar", buf.String())
}
