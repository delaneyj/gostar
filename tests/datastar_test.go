package tests

import (
	"testing"
	"time"

	. "github.com/delaneyj/gostar/elements"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/bytebufferpool"
)

type result struct {
	Expected string
	Actual   ElementRenderer
}

func TestDatastarAttr(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-attr:title=\"$foo\"></div>",
			Actual:   DIV().DATASTAR_ATTR("title", "$foo"),
		},
		{
			Expected: "<div></div>",
			Actual:   DIV().DATASTAR_ATTR("title", "$foo").DATASTAR_ATTRRemove("title"),
		},
		{
			Expected: "<div data-attr=\"{title: $foo, disabled: $bar}\"></div>",
			Actual:   DIV().DATASTAR_ATTR("", "{title: $foo, disabled: $bar}"),
		},
		{
			Expected: "<div></div>",
			Actual:   DIV().DATASTAR_ATTR("", "{title: $foo, disabled: $bar}").DATASTAR_ATTRRemove(""),
		},
	})
}

func TestDatastarBind(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-bind:title=\"foo\"></div>",
			Actual:   DIV().DATASTAR_BIND("title", "foo"),
		},
		{
			Expected: "<div></div>",
			Actual:   DIV().IfDATASTAR_BIND(false, "title", "foo"),
		},
	})
}

func TestDatastarClass(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-class:hidden=\"$foo\"></div>",
			Actual:   DIV().DATASTAR_CLASS("hidden", "$foo"),
		},
		{
			Expected: "<div data-class:hidden__case.camel=\"$foo\"></div>",
			Actual:   DIV().DATASTAR_CLASS("hidden", "$foo", DivClassModCase("camel")),
		},
	})
}

func TestDatastarComputed(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-computed:foo=\"$bar + $baz\"></div>",
			Actual:   DIV().DATASTAR_COMPUTED("foo", "$bar + $baz"),
		},
		{
			Expected: "<div data-computed:foo__case.kebab=\"$bar + $baz\"></div>",
			Actual:   DIV().DATASTAR_COMPUTED("foo", "$bar + $baz", DivComputedModCase("kebab")),
		},
	})
}

func TestDatastarEffect(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-effect=\"$foo = $bar + $baz\"></div>",
			Actual:   DIV().DATASTAR_EFFECT("$foo = $bar + $baz"),
		},
		{
			Expected: "<div data-effect=\"$foo = $bar + $baz\"></div>",
			Actual:   DIV().IfDATASTAR_EFFECT(true, "$foo = $bar + $baz"),
		},
	})
}

func TestDatastarIgnore(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-ignore></div>",
			Actual:   DIV().DATASTAR_IGNORE(),
		},
		{
			Expected: "<div data-ignore__self></div>",
			Actual:   DIV().DATASTAR_IGNORE(DivIgnoreModSelf()),
		},
	})
}

func TestDatastarIgnoreMorph(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-ignore-morph></div>",
			Actual:   DIV().DATASTAR_IGNORE_MORPH(),
		},
		{
			Expected: "<div></div>",
			Actual:   DIV().DATASTAR_IGNORE_MORPHSet(false),
		},
	})
}

func TestDatastarIndicator(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-indicator=\"fetching\"></div>",
			Actual:   DIV().DATASTAR_INDICATOR("fetching"),
		},
		{
			Expected: "<div data-indicator__case.pascal=\"foo\"></div>",
			Actual:   DIV().DATASTAR_INDICATOR("foo", DivIndicatorModCase("pascal")),
		},
	})
}

func TestDatastarInit(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-init=\"$count = 1\"></div>",
			Actual:   DIV().DATASTAR_INIT("$count = 1"),
		},
		{
			Expected: "<div data-init__delay.200ms=\"$count = 1\"></div>",
			Actual:   DIV().DATASTAR_INIT("$count = 1", DivInitModDelayMs(200*time.Millisecond)),
		},
		{
			Expected: "<div data-init__delay.5s=\"$count = 1\"></div>",
			Actual:   DIV().DATASTAR_INIT("$count = 1", DivInitModDelaySec(5*time.Second)),
		},
		{
			Expected: "<div data-init__viewtransition=\"$count = 1\"></div>",
			Actual:   DIV().DATASTAR_INIT("$count = 1", DivInitModViewTransition()),
		},
	})
}

func TestDatastarJsonSignals(t *testing.T) {
	run(t, []result{
		{
			Expected: "<pre data-json-signals=\"{include: /^app/, exclude: /password/}\"></pre>",
			Actual:   PRE().DATASTAR_JSON_SIGNALS("{include: /^app/, exclude: /password/}"),
		},
		{
			Expected: "<pre data-json-signals__terse=\"{include: /^app/, exclude: /password/}\"></pre>",
			Actual:   PRE().DATASTAR_JSON_SIGNALS("{include: /^app/, exclude: /password/}", PreJsonSignalsModTerse()),
		},
	})
}

func TestDatastarOn(t *testing.T) {
	run(t, []result{
		{
			Expected: "<button data-on:click=\"$foo = ''\"></button>",
			Actual:   BUTTON().DATASTAR_ON("click", "$foo = ''"),
		},
		{
			Expected: "<button data-on:click__window=\"$foo = ''\"></button>",
			Actual:   BUTTON().DATASTAR_ON("click", "$foo = ''", ButtonOnModWindow()),
		},
		{
			Expected: "<div data-on:my-event__debounce.2s=\"$foo = 'bar'\"></div>",
			Actual:   DIV().DATASTAR_ON("my-event", "$foo = 'bar'", DivOnModDebounceSec(2*time.Second)),
		},
		{
			Expected: "<button data-on:click__debounce.500ms.leading=\"$foo = ''\"></button>",
			Actual:   BUTTON().DATASTAR_ON("click", "$foo = ''", ButtonOnModDebounceMsLeading(500*time.Millisecond)),
		},
		{
			Expected: "<button data-on:click__window__debounce.1s.notrailing=\"$foo = ''\"></button>",
			Actual:   BUTTON().DATASTAR_ON("click", "$foo = ''", ButtonOnModWindow(), ButtonOnModDebounceSecNoTrailing(1*time.Second)),
		},
		{
			Expected: "<span data-on:click__throttle.200ms=\"$foo = 'bar'\"></span>",
			Actual:   SPAN().DATASTAR_ON("click", "$foo = 'bar'", SpanOnModThrottleMs(200*time.Millisecond)),
		},
		{
			Expected: "<button data-on:click__throttle.500ms.trailing__window=\"$foo = ''\"></button>",
			Actual:   BUTTON().DATASTAR_ON("click", "$foo = ''", ButtonOnModThrottleMsTrailing(500*time.Millisecond), ButtonOnModWindow()),
		},
		{
			Expected: "<div data-on:my-event__throttle.3s.noleading__case.snake=\"$foo = 'evt.detail'\"></div>",
			Actual:   DIV().DATASTAR_ON("my-event", "$foo = 'evt.detail'", DivOnModThrottleSecNoLeading(3*time.Second), DivOnModCase("snake")),
		},
	})
}

func TestDatastarOnIntersect(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-on-intersect=\"$intersected = true\"></div>",
			Actual:   DIV().DATASTAR_ON_INTERSECT("$intersected = true"),
		},
		{
			Expected: "<div data-on-intersect__once__full=\"$fullyIntersected = true\"></div>",
			Actual:   DIV().DATASTAR_ON_INTERSECT("$fullyIntersected = true", DivOnIntersectModOnce(), DivOnIntersectModFull()),
		},
		{
			Expected: "<div data-on-intersect__half__debounce.2s=\"$halfIntersected = true\"></div>",
			Actual:   DIV().DATASTAR_ON_INTERSECT("$halfIntersected = true", DivOnIntersectModHalf(), DivOnIntersectModDebounceSec(2*time.Second)),
		},
		{
			Expected: "<button data-on-intersect__debounce.500ms.leading=\"$intersected = true\"></button>",
			Actual:   BUTTON().DATASTAR_ON_INTERSECT("$intersected = true", ButtonOnIntersectModDebounceMsLeading(500*time.Millisecond)),
		},
		{
			Expected: "<button data-on-intersect__viewtransition__debounce.1s.notrailing=\"$intersected = true\"></button>",
			Actual:   BUTTON().DATASTAR_ON_INTERSECT("$intersected = true", ButtonOnIntersectModViewTransition(), ButtonOnIntersectModDebounceSecNoTrailing(1*time.Second)),
		},
		{
			Expected: "<span data-on-intersect__throttle.200ms=\"$intersected = true\"></span>",
			Actual:   SPAN().DATASTAR_ON_INTERSECT("$intersected = true", SpanOnIntersectModThrottleMs(200*time.Millisecond)),
		},
		{
			Expected: "<button data-on-intersect__throttle.500ms.trailing__viewtransition=\"$intersected = true\"></button>",
			Actual:   BUTTON().DATASTAR_ON_INTERSECT("$intersected = true", ButtonOnIntersectModThrottleMsTrailing(500*time.Millisecond), ButtonOnIntersectModViewTransition()),
		},
		{
			Expected: "<div data-on-intersect__throttle.3s.noleading=\"$intersected = true\"></div>",
			Actual:   DIV().DATASTAR_ON_INTERSECT("$intersected = true", DivOnIntersectModThrottleSecNoLeading(3*time.Second)),
		},
	})
}

func TestDataOnInterval(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-on-interval__duration.500ms=\"$count++\"></div>",
			Actual:   DIV().DATASTAR_ON_INTERVAL("$count++", DivOnIntervalModDurationMs(500*time.Millisecond)),
		},
		{
			Expected: "<div data-on-interval__duration.2s.leading=\"$count++\"></div>",
			Actual:   DIV().DATASTAR_ON_INTERVAL("$count++", DivOnIntervalModDurationSecLeading(2*time.Second)),
		},
		{
			Expected: "<div data-on-interval__duration.2s__viewtransition=\"$count++\"></div>",
			Actual:   DIV().DATASTAR_ON_INTERVAL("$count++", DivOnIntervalModDurationSec(2*time.Second), DivOnIntervalModViewTransition()),
		},
	})
}

func TestDatastarOnSignalPatch(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-on-signal-patch=\"console.log('A signal changed!')\"></div>",
			Actual:   DIV().DATASTAR_ON_SIGNAL_PATCH("console.log('A signal changed!')"),
		},
		{
			Expected: "<div data-on-signal-patch__delay.500ms=\"console.log('A signal changed!')\"></div>",
			Actual:   DIV().DATASTAR_ON_SIGNAL_PATCH("console.log('A signal changed!')", DivOnSignalPatchModDelayMs(500*time.Millisecond)),
		},
		{
			Expected: "<div data-on-signal-patch__debounce.500ms.leading=\"console.log('A signal changed!')\"></div>",
			Actual:   DIV().DATASTAR_ON_SIGNAL_PATCH("console.log('A signal changed!')", DivOnSignalPatchModDebounceMsLeading(500*time.Millisecond)),
		},
	})
}

func TestDatastarOnSignalPatchFilter(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-on-signal-patch-filter=\"{include: /^counter$/}\"></div>",
			Actual:   DIV().DATASTAR_ON_SIGNAL_PATCH_FILTER("{include: /^counter$/}"),
		},
	})
}

func TestDatastarPreserveAttr(t *testing.T) {
	run(t, []result{
		{
			Expected: "<details data-preserve-attr=\"open class\" open></details>",
			Actual:   DETAILS().OPEN().DATASTAR_PRESERVE_ATTR("open class"),
		},
	})
}

func TestDatastarRef(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-ref=\"foo\"></div>",
			Actual:   DIV().DATASTAR_REF("foo"),
		},
		{
			Expected: "<div data-ref__case.camel=\"fooBar\"></div>",
			Actual:   DIV().DATASTAR_REF("fooBar", DivRefModCase("camel")),
		},
	})
}

func TestDatastarShow(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-show=\"$foo\"></div>",
			Actual:   DIV().DATASTAR_SHOW("$foo"),
		},
	})
}

func TestDatastaSignals(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-signals:foo=\"1\"></div>",
			Actual:   DIV().DATASTAR_SIGNALS("foo", "1"),
		},
		{
			Expected: "<div data-signals=\"{foo: {bar: 1, baz: 2}}\"></div>",
			Actual:   DIV().DATASTAR_SIGNALS("", "{foo: {bar: 1, baz: 2}}"),
		},
		{
			Expected: "<div data-signals__ifmissing=\"{foo: null}\"></div>",
			Actual:   DIV().DATASTAR_SIGNALS("", "{foo: null}", DivSignalsModIfMissing()),
		},
		{
			Expected: "<div data-signals:my-signal__case.kebab=\"1\"></div>",
			Actual:   DIV().DATASTAR_SIGNALS("my-signal", "1", DivSignalsModCase("kebab")),
		},
	})
}

func TestDatastarStyle(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-style:background-color=\"$usingRed ? 'red' : 'blue'\"></div>",
			Actual:   DIV().DATASTAR_STYLE("background-color", "$usingRed ? 'red' : 'blue'"),
		},
		{
			Expected: "<div data-style=\"{color: '$x && green'\" style=\"color:red\"></div>",
			Actual:   DIV().STYLE("color", "red").DATASTAR_STYLE("", "{color: '$x && green'"),
		},
	})
}

func TestDatastarText(t *testing.T) {
	run(t, []result{
		{
			Expected: "<div data-text=\"$foo\"></div>",
			Actual:   DIV().DATASTAR_TEXT("$foo"),
		},
	})
}

func run(t *testing.T, results []result) {
	for _, result := range results {
		buf := bytebufferpool.Get()
		e := result.Expected
		err := result.Actual.Render(buf)
		assert.NoError(t, err)
		a := buf.String()
		assert.Equal(t, e, a)
		bytebufferpool.Put(buf)
	}
}
