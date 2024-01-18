package fad

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 256 256"
	fill           = "currentColor"
)

type FadIcon struct {
	*SVGSVGElement
}

type FadIconFn func(children ...ElementRenderer) *FadIcon

var IconLookup = map[string]FadIconFn{
	"adr":               Adr,
	"adsr":              Adsr,
	"ahdsr":             Ahdsr,
	"ar":                Ar,
	"armrecording":      Armrecording,
	"arpchord":          Arpchord,
	"arpdown":           Arpdown,
	"arpdownandup":      Arpdownandup,
	"arpdownup":         Arpdownup,
	"arpplayorder":      Arpplayorder,
	"arprandom":         Arprandom,
	"arpup":             Arpup,
	"arpupandown":       Arpupandown,
	"arpupdown":         Arpupdown,
	"arrowsHorz":        ArrowsHorz,
	"arrowsVert":        ArrowsVert,
	"automationFourP":   AutomationFourP,
	"automationThreeP":  AutomationThreeP,
	"automationTwoP":    AutomationTwoP,
	"backward":          Backward,
	"bluetooth":         Bluetooth,
	"caretDown":         CaretDown,
	"caretLeft":         CaretLeft,
	"caretRight":        CaretRight,
	"caretUp":           CaretUp,
	"close":             Close,
	"copy":              Copy,
	"cpu":               Cpu,
	"cutter":            Cutter,
	"digitalColon":      DigitalColon,
	"digitalDot":        DigitalDot,
	"digitalEight":      DigitalEight,
	"digitalFive":       DigitalFive,
	"digitalFour":       DigitalFour,
	"digitalNine":       DigitalNine,
	"digitalOne":        DigitalOne,
	"digitalSeven":      DigitalSeven,
	"digitalSix":        DigitalSix,
	"digitalThree":      DigitalThree,
	"digitalTwo":        DigitalTwo,
	"digitalZero":       DigitalZero,
	"diskio":            Diskio,
	"drumpad":           Drumpad,
	"duplicate":         Duplicate,
	"eraser":            Eraser,
	"ffwd":              Ffwd,
	"filterBandpass":    FilterBandpass,
	"filterBell":        FilterBell,
	"filterBypass":      FilterBypass,
	"filterHighpass":    FilterHighpass,
	"filterLowpass":     FilterLowpass,
	"filterNotch":       FilterNotch,
	"filterRezHighpass": FilterRezHighpass,
	"filterRezLowpass":  FilterRezLowpass,
	"filterShelvingHi":  FilterShelvingHi,
	"filterShelvingLo":  FilterShelvingLo,
	"foldback":          Foldback,
	"forward":           Forward,
	"hexpand":           Hexpand,
	"hardclip":          Hardclip,
	"hardclipcurve":     Hardclipcurve,
	"headphones":        Headphones,
	"keyboard":          Keyboard,
	"lock":              Lock,
	"logoAax":           LogoAax,
	"logoAbletonlink":   LogoAbletonlink,
	"logoAu":            LogoAu,
	"logoAudacity":      LogoAudacity,
	"logoAudiobus":      LogoAudiobus,
	"logoCubase":        LogoCubase,
	"logoFl":            LogoFl,
	"logoJuce":          LogoJuce,
	"logoLadspa":        LogoLadspa,
	"logoLive":          LogoLive,
	"logoLvTwo":         LogoLvTwo,
	"logoProtools":      LogoProtools,
	"logoRackext":       LogoRackext,
	"logoReaper":        LogoReaper,
	"logoReason":        LogoReason,
	"logoRewire":        LogoRewire,
	"logoStudioone":     LogoStudioone,
	"logoTracktion":     LogoTracktion,
	"logoVst":           LogoVst,
	"logoWaveform":      LogoWaveform,
	"loop":              Loop,
	"metronome":         Metronome,
	"microphone":        Microphone,
	"midiplug":          Midiplug,
	"modrandom":         Modrandom,
	"modsawdown":        Modsawdown,
	"modsawup":          Modsawup,
	"modsh":             Modsh,
	"modsine":           Modsine,
	"modsquare":         Modsquare,
	"modtri":            Modtri,
	"modularplug":       Modularplug,
	"mono":              Mono,
	"mute":              Mute,
	"next":              Next,
	"open":              Open,
	"paste":             Paste,
	"pause":             Pause,
	"pen":               Pen,
	"phase":             Phase,
	"play":              Play,
	"pointer":           Pointer,
	"powerswitch":       Powerswitch,
	"presetA":           PresetA,
	"presetAb":          PresetAb,
	"presetB":           PresetB,
	"presetBa":          PresetBa,
	"prev":              Prev,
	"punchIn":           PunchIn,
	"punchOut":          PunchOut,
	"ram":               Ram,
	"randomOneDice":     RandomOneDice,
	"randomTwoDice":     RandomTwoDice,
	"record":            Record,
	"redo":              Redo,
	"repeat":            Repeat,
	"repeatOne":         RepeatOne,
	"rew":               Rew,
	"roundswitchOff":    RoundswitchOff,
	"roundswitchOn":     RoundswitchOn,
	"save":              Save,
	"saveas":            Saveas,
	"scissors":          Scissors,
	"shuffle":           Shuffle,
	"sliderRoundOne":    SliderRoundOne,
	"sliderRoundThree":  SliderRoundThree,
	"sliderRoundTwo":    SliderRoundTwo,
	"sliderhandleOne":   SliderhandleOne,
	"sliderhandleTwo":   SliderhandleTwo,
	"softclip":          Softclip,
	"softclipcurve":     Softclipcurve,
	"solo":              Solo,
	"speaker":           Speaker,
	"squareswitchOff":   SquareswitchOff,
	"squareswitchOn":    SquareswitchOn,
	"stereo":            Stereo,
	"stop":              Stop,
	"thunderbolt":       Thunderbolt,
	"timeselect":        Timeselect,
	"undo":              Undo,
	"unlock":            Unlock,
	"usb":               Usb,
	"vexpand":           Vexpand,
	"vroundswitchOff":   VroundswitchOff,
	"vroundswitchOn":    VroundswitchOn,
	"vsquareswitchOff":  VsquareswitchOff,
	"vsquareswitchOn":   VsquareswitchOn,
	"waveform":          Waveform,
	"xlrplug":           Xlrplug,
	"zoomin":            Zoomin,
	"zoomout":           Zoomout,
}

func Adr(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25 184c-.47 2.68-.227 4.354 1 6s4.563 2.464 8 2c3.437-.464 5.078-.958 6-4c.922-3.042 39.818-102.423 39.818-102.423c1.205-3.08 3.034-3.022 4.08.115l16.331 48.974c1.049 3.144 4.353 6.781 7.386 8.127c0 0 108.132 48.015 111.214 48.404c3.083.389 5.414.105 7.19-1.197c1.68-1.23 3.532-3.733 4.124-6c.591-2.267-1.12-5.898-3.343-7.28c-2.223-1.38-109.625-46.556-109.625-46.556c-2.026-.853-4.148-3.274-4.736-5.389c0 0-12.137-43.757-14.135-49.484c-1.998-5.727-6.562-10.983-14.989-10.983c-8.426 0-13.23 2.387-15.315 7.692c-2.086 5.305-42.53 109.32-43 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adsr(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25 184c-.47 2.68-.227 4.354 1 6s4.563 2.464 8 2c3.437-.464 5.078-.958 6-4c.922-3.042 40.167-102.359 40.167-102.359c1.012-2.563 2.655-2.563 3.676.018l15.08 38.128c1.218 3.077 4.9 5.591 8.21 5.616L188 130s26.23 53.28 27.993 56.358c1.764 3.079 4.94 5.225 7.237 5.24c2.297.016 5.15-2.404 6.552-3.598c1.402-1.194 2.88-3.828 1.661-6.824c-1.218-2.995-29.623-58.905-31.318-61.992C198.43 116.096 196.266 114 192 114s-76.877-.318-76.877-.318c-2.207-.01-4.608-1.691-5.37-3.777c0 0-12.403-33.889-13.717-37.905c-1.314-4.016-6.912-7.827-14.036-7.827S70.086 66.695 68 72c-2.086 5.305-42.53 109.32-43 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ahdsr(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25 184c-.47 2.68-.227 4.354 1 6s4.563 2.464 8 2c3.437-.464 5.078-.958 6-4c.922-3.042 20.436-105.053 20.436-105.053c.311-1.627 1.916-2.947 3.57-2.947H99.64c1.105 0 2.262.847 2.59 1.914l13.007 42.35c.973 3.168 4.445 5.736 7.754 5.736h64.644s28.03 54.976 29.423 58c1.392 3.024 2.537 3.984 4.834 4c2.296.016 6.895-3.215 8.25-4.63c1.355-1.414 1.842-3.08.624-6.075c-1.219-2.996-28.947-59.024-30.642-62.111C198.43 116.096 196.266 114 192 114s-59.994.286-59.994.286c-2.213.011-4.485-1.708-5.076-3.835c0 0-10.546-38.069-11.84-41.724C113.795 65.07 111.369 64 106 64H56c-5.6 0-5.793.422-7.879 5.727C46.035 75.03 25.471 181.319 25 184"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ar(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25 184c-.47 2.68-.227 4.354 1 6s4.563 2.464 8 2c3.437-.464 5.078-.958 6-4c.922-3.042 40.911-105.217 40.911-105.217c.601-1.537 2.13-1.93 3.404-.886c0 0 134.305 110.087 136.602 110.103c1.194.008 3.15-.157 4.942-1.477c1.656-1.22 3.173-3.57 3.824-4.25c1.355-1.414 3.54-3.424.46-6.697C227.063 176.304 104.148 78.758 100 74c-4.148-4.758-10.631-10-16-10s-13.914 2.695-16 8c-2.086 5.305-42.53 109.32-43 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Armrecording(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="127" cy="129" r="81" fill="currentColor" fill-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arpchord(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M96 203.996A4.002 4.002 0 0 1 91.996 208H68.004A4.002 4.002 0 0 1 64 203.996v-23.992A4.002 4.002 0 0 1 68.004 176h23.992A4.002 4.002 0 0 1 96 180.004zm0-64A4.002 4.002 0 0 1 91.996 144H68.004A4.002 4.002 0 0 1 64 139.996v-23.992A4.002 4.002 0 0 1 68.004 112h23.992A4.002 4.002 0 0 1 96 116.004zm0-64A4.002 4.002 0 0 1 91.996 80H68.004A4.002 4.002 0 0 1 64 75.996V52.004A4.002 4.002 0 0 1 68.004 48h23.992A4.002 4.002 0 0 1 96 52.004zm64 0A4.002 4.002 0 0 0 164.004 80h23.992A4.002 4.002 0 0 0 192 75.996V52.004A4.002 4.002 0 0 0 187.996 48h-23.992A4.002 4.002 0 0 0 160 52.004zm0 64a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004v-23.992a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004zm0 64a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004v-23.992a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arpdown(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M80 52.004A4.002 4.002 0 0 0 75.996 48H52.004A4.002 4.002 0 0 0 48 52.004v23.992A4.002 4.002 0 0 0 52.004 80h23.992A4.002 4.002 0 0 0 80 75.996zm32 32A4.002 4.002 0 0 0 107.996 80H84.004A4.002 4.002 0 0 0 80 84.004v23.992A4.002 4.002 0 0 0 84.004 112h23.992a4.002 4.002 0 0 0 4.004-4.004zm32 32a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004v23.992a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004zm32 32a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004v23.992a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004zm32 32a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004v23.992a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arpdownandup(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M80 75.996A4.002 4.002 0 0 1 75.996 80H52.004A4.002 4.002 0 0 1 48 75.996V52.004A4.002 4.002 0 0 1 52.004 48h23.992A4.002 4.002 0 0 1 80 52.004zm16 64A4.002 4.002 0 0 1 91.996 144H68.004A4.002 4.002 0 0 1 64 139.996v-23.992A4.002 4.002 0 0 1 68.004 112h23.992A4.002 4.002 0 0 1 96 116.004zm16 64a4.002 4.002 0 0 1-4.004 4.004H84.004A4.002 4.002 0 0 1 80 203.996v-23.992A4.002 4.002 0 0 1 84.004 176h23.992a4.002 4.002 0 0 1 4.004 4.004zm64 0a4.002 4.002 0 0 1-4.004 4.004h-23.992a4.002 4.002 0 0 1-4.004-4.004v-23.992a4.002 4.002 0 0 1 4.004-4.004h23.992a4.002 4.002 0 0 1 4.004 4.004zm16-64a4.002 4.002 0 0 1-4.004 4.004h-23.992a4.002 4.002 0 0 1-4.004-4.004v-23.992a4.002 4.002 0 0 1 4.004-4.004h23.992a4.002 4.002 0 0 1 4.004 4.004zm16-64A4.002 4.002 0 0 1 203.996 80h-23.992A4.002 4.002 0 0 1 176 75.996V52.004A4.002 4.002 0 0 1 180.004 48h23.992A4.002 4.002 0 0 1 208 52.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arpdownup(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M80 52.004A4.002 4.002 0 0 0 75.996 48H52.004A4.002 4.002 0 0 0 48 52.004v23.992A4.002 4.002 0 0 0 52.004 80h23.992A4.002 4.002 0 0 0 80 75.996zm32 64a4.002 4.002 0 0 0-4.004-4.004H84.004A4.002 4.002 0 0 0 80 116.004v23.992A4.002 4.002 0 0 0 84.004 144h23.992a4.002 4.002 0 0 0 4.004-4.004zm32 64a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004v23.992a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004zm32-64a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004v23.992a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004zm32-64A4.002 4.002 0 0 0 203.996 48h-23.992A4.002 4.002 0 0 0 176 52.004v23.992A4.002 4.002 0 0 0 180.004 80h23.992A4.002 4.002 0 0 0 208 75.996z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arpplayorder(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M80 203.996A4.002 4.002 0 0 1 75.996 208H52.004A4.002 4.002 0 0 1 48 203.996v-23.992A4.002 4.002 0 0 1 52.004 176h23.992A4.002 4.002 0 0 1 80 180.004zm16-64A4.002 4.002 0 0 1 91.996 144H68.004A4.002 4.002 0 0 1 64 139.996v-23.992A4.002 4.002 0 0 1 68.004 112h23.992A4.002 4.002 0 0 1 96 116.004zm16-64A4.002 4.002 0 0 1 107.996 80H84.004A4.002 4.002 0 0 1 80 75.996V52.004A4.002 4.002 0 0 1 84.004 48h23.992A4.002 4.002 0 0 1 112 52.004zm64 0A4.002 4.002 0 0 0 180.004 80h23.992A4.002 4.002 0 0 0 208 75.996V52.004A4.002 4.002 0 0 0 203.996 48h-23.992A4.002 4.002 0 0 0 176 52.004zm-16 64a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004v-23.992a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004zm-16 64a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004v-23.992a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arprandom(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M80 123.996A4.002 4.002 0 0 1 75.996 128H52.004A4.002 4.002 0 0 1 48 123.996v-23.992A4.002 4.002 0 0 1 52.004 96h23.992A4.002 4.002 0 0 1 80 100.004zm32 80a4.002 4.002 0 0 1-4.004 4.004H84.004A4.002 4.002 0 0 1 80 203.996v-23.992A4.002 4.002 0 0 1 84.004 176h23.992a4.002 4.002 0 0 1 4.004 4.004zm96-64a4.002 4.002 0 0 1-4.004 4.004h-23.992a4.002 4.002 0 0 1-4.004-4.004v-23.992a4.002 4.002 0 0 1 4.004-4.004h23.992a4.002 4.002 0 0 1 4.004 4.004zm-96 16a4.002 4.002 0 0 0 4.004 4.004h23.992a4.002 4.002 0 0 0 4.004-4.004v-23.992a4.002 4.002 0 0 0-4.004-4.004h-23.992a4.002 4.002 0 0 0-4.004 4.004zm32-80A4.002 4.002 0 0 0 148.004 80h23.992A4.002 4.002 0 0 0 176 75.996V52.004A4.002 4.002 0 0 0 171.996 48h-23.992A4.002 4.002 0 0 0 144 52.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arpup(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M48 203.996A4.002 4.002 0 0 0 52.004 208h25.994c1.106 0 2.002-.89 2.002-2.002v-27.996c0-1.106-.89-2.002-2.002-2.002H50.002c-1.106 0-2.002.89-2.002 2.002zm32-29.998c0 1.106.89 2.002 2.002 2.002h27.996c1.106 0 2.002-.89 2.002-2.002v-27.996c0-1.106-.89-2.002-2.002-2.002H82.002c-1.106 0-2.002.89-2.002 2.002zm32-32c0 1.106.89 2.002 2.002 2.002h27.996c1.106 0 2.002-.89 2.002-2.002v-27.996c0-1.106-.89-2.002-2.002-2.002h-27.996c-1.106 0-2.002.89-2.002 2.002zm32-32c0 1.106.89 2.002 2.002 2.002h27.996c1.106 0 2.002-.89 2.002-2.002V82.002c0-1.106-.89-2.002-2.002-2.002h-27.996c-1.106 0-2.002.89-2.002 2.002zm32-32c0 1.106.89 2.002 2.002 2.002h27.996c1.106 0 2.002-.89 2.002-2.002V50.002c0-1.106-.89-2.002-2.002-2.002h-27.996c-1.106 0-2.002.89-2.002 2.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arpupandown(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M80 203.996A4.002 4.002 0 0 1 75.996 208H52.004A4.002 4.002 0 0 1 48 203.996v-23.992A4.002 4.002 0 0 1 52.004 176h23.992A4.002 4.002 0 0 1 80 180.004zm16-64A4.002 4.002 0 0 1 91.996 144H68.004A4.002 4.002 0 0 1 64 139.996v-23.992A4.002 4.002 0 0 1 68.004 112h23.992A4.002 4.002 0 0 1 96 116.004zm16-64A4.002 4.002 0 0 1 107.996 80H84.004A4.002 4.002 0 0 1 80 75.996V52.004A4.002 4.002 0 0 1 84.004 48h23.992A4.002 4.002 0 0 1 112 52.004zm64 0A4.002 4.002 0 0 1 171.996 80h-23.992A4.002 4.002 0 0 1 144 75.996V52.004A4.002 4.002 0 0 1 148.004 48h23.992A4.002 4.002 0 0 1 176 52.004zm16 64a4.002 4.002 0 0 1-4.004 4.004h-23.992a4.002 4.002 0 0 1-4.004-4.004v-23.992a4.002 4.002 0 0 1 4.004-4.004h23.992a4.002 4.002 0 0 1 4.004 4.004zm16 64a4.002 4.002 0 0 1-4.004 4.004h-23.992a4.002 4.002 0 0 1-4.004-4.004v-23.992a4.002 4.002 0 0 1 4.004-4.004h23.992a4.002 4.002 0 0 1 4.004 4.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arpupdown(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M80 203.996A4.002 4.002 0 0 1 75.996 208H52.004A4.002 4.002 0 0 1 48 203.996v-23.992A4.002 4.002 0 0 1 52.004 176h23.992A4.002 4.002 0 0 1 80 180.004zm32-64a4.002 4.002 0 0 1-4.004 4.004H84.004A4.002 4.002 0 0 1 80 139.996v-23.992A4.002 4.002 0 0 1 84.004 112h23.992a4.002 4.002 0 0 1 4.004 4.004zm32-64A4.002 4.002 0 0 1 139.996 80h-23.992A4.002 4.002 0 0 1 112 75.996V52.004A4.002 4.002 0 0 1 116.004 48h23.992A4.002 4.002 0 0 1 144 52.004zm32 64a4.002 4.002 0 0 1-4.004 4.004h-23.992a4.002 4.002 0 0 1-4.004-4.004v-23.992a4.002 4.002 0 0 1 4.004-4.004h23.992a4.002 4.002 0 0 1 4.004 4.004zm32 64a4.002 4.002 0 0 1-4.004 4.004h-23.992a4.002 4.002 0 0 1-4.004-4.004v-23.992a4.002 4.002 0 0 1 4.004-4.004h23.992a4.002 4.002 0 0 1 4.004 4.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsHorz(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M100.218 80.292c-2.205 0-2.746-1.292-1.214-2.883L128 47.313l29.24 30.11c1.538 1.585 1 2.87-1.209 2.87h-19.239l.636 95.811h18.596c2.21 0 2.746 1.28 1.201 2.854L128 208.746l-28.982-29.773c-1.542-1.584-1.004-2.869 1.2-2.869h19.54V80.292z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsVert(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M175.863 100.122c0-2.205 1.293-2.747 2.883-1.214l30.096 28.996l-30.11 29.24c-1.585 1.538-2.87 1-2.87-1.209v-19.24l-95.811.637v18.596c0 2.21-1.28 2.746-2.854 1.201l-29.788-29.225l29.774-28.982c1.584-1.542 2.868-1.004 2.868 1.2v19.54h95.812z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AutomationFourP(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M144.282 145.51A32.19 32.19 0 0 1 150 145c4.032 0 7.89.746 11.444 2.107L175.584 127l.54.306A31.88 31.88 0 0 1 168 106c0-17.673 14.327-32 32-32c17.673 0 32 14.327 32 32c0 17.673-14.327 32-32 32c-3.672 0-7.2-.619-10.485-1.757l-14.31 21.038A31.863 31.863 0 0 1 182 177c0 17.673-14.327 32-32 32c-17.673 0-32-14.327-32-32c0-9.767 4.376-18.512 11.274-24.382l-20.764-41.28A32.14 32.14 0 0 1 102 112a32.05 32.05 0 0 1-8.16-1.05l-14.716 25.93C85.21 142.705 89 150.91 89 160c0 17.673-14.327 32-32 32c-17.673 0-32-14.327-32-32c0-17.673 14.327-32 32-32c2.655 0 5.234.323 7.7.932l14.809-26.17C73.638 96.963 70 88.907 70 80c0-17.673 14.327-32 32-32c17.673 0 32 14.327 32 32c0 9.563-4.195 18.146-10.844 24.01zM200 122c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16M57 176c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16m45-80c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16m48 97c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AutomationThreeP(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M129.175 107.005A31.85 31.85 0 0 1 112 112a32.54 32.54 0 0 1-1.503-.035l.077.035l-12.804 37.383C106.35 155.123 112 164.902 112 176c0 17.673-14.327 32-32 32c-17.673 0-32-14.327-32-32c0-17.673 14.327-32 32-32c.87 0 1.731.035 2.583.103l12.14-37.163C85.867 101.248 80 91.31 80 80c0-17.673 14.327-32 32-32c17.673 0 32 14.327 32 32c0 5.51-1.393 10.696-3.846 15.224l19.42 21.308A31.851 31.851 0 0 1 176 112c17.673 0 32 14.327 32 32c0 17.673-14.327 32-32 32c-17.673 0-32-14.327-32-32a31.857 31.857 0 0 1 3.92-15.36zM80 192c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16m32-96c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16m64 64c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AutomationTwoP(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m95.847 148.193l52.206-52.594A31.855 31.855 0 0 1 144 80c0-17.673 14.327-32 32-32c17.673 0 32 14.327 32 32c0 17.673-14.327 32-32 32a31.85 31.85 0 0 1-16.89-4.815l-52.068 51.699A31.85 31.85 0 0 1 112 176c0 17.673-14.327 32-32 32c-17.673 0-32-14.327-32-32c0-17.673 14.327-32 32-32a31.854 31.854 0 0 1 15.847 4.193M176 96c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16m-96 96c8.837 0 16-7.163 16-16s-7.163-16-16-16s-16 7.163-16 16s7.163 16 16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M117.027 60.858c0 14.184.118 120.734.118 134.66c0 13.925-7.798 16.307-14.543 10.568c-6.745-5.74-46.666-60.003-52.014-67.201c-5.349-7.198-4.45-12.951.086-20.28c4.535-7.328 47.284-62.224 52.46-68.2c5.175-5.977 13.893-3.73 13.893 10.453M65.424 132.65l32.874 43.167c1.336 1.754 2.425 1.393 2.424-.814l-.005-92.006c0-2.197-1.072-2.562-2.395-.792l-32.927 44.057c-1.327 1.776-1.31 4.63.03 6.388zm142.603-71.792c0 14.184.118 120.734.118 134.66c0 13.925-7.798 16.307-14.543 10.568c-6.745-5.74-46.666-60.003-52.014-67.201c-5.349-7.198-4.45-12.951.086-20.28c4.535-7.328 47.284-62.224 52.46-68.2c5.175-5.977 13.893-3.73 13.893 10.453m-51.603 71.792l32.874 43.167c1.336 1.754 2.425 1.393 2.424-.814l-.005-92.006c0-2.197-1.072-2.562-2.395-.792l-32.927 44.057c-1.327 1.776-1.31 4.63.03 6.388z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M131.31 27.08c5.287 4.567 56.852 46.483 59.5 49.262c2.648 2.778 2.494 8.238-.488 11.384c-2.981 3.145-49.453 41.303-49.453 41.303s45.242 37.302 48.583 40.157c3.341 2.855 3.121 11.217.87 13.162c-2.25 1.945-57.84 47.267-60.01 48.492c-2.168 1.224-8.976-1.584-9.444-5.395c-.468-3.81-1.099-77.728-1.099-77.728c-.01-1.102-.699-1.425-1.54-.721c0 0-41.477 34.746-44.537 36.474c-3.06 1.727-4.94.411-6.66-.901c-1.72-1.313-2.613-4.435-3.117-6.674c-.505-2.24.08-4.673 1.278-5.746c1.198-1.073 46.222-38.744 46.222-38.744c1.692-1.416 1.684-3.71-.022-5.122c0 0-45.314-37.463-47.08-39.048c-1.767-1.584-.83-4.267.002-6.382c.832-2.116 2.777-6.545 4.03-7.415c1.255-.87 2.628-.61 4.818 1.025c2.19 1.636 45.684 36.613 45.684 36.613c.863.694 1.55.352 1.545-.755c0 0-.306-75.85-.306-77.786c0-1.935 5.936-10.022 11.224-5.455m4.124 26.561l-.335 56.21c-.007 1.106.682 1.439 1.538.742l33.03-26.877c1.278-1.04 1.287-2.746.01-3.801l-32.684-27c-.85-.701-1.552-.38-1.559.726m1.453 92.099c-.831-.72-1.51-.407-1.517.7l-.351 59.45c-.007 1.106.662 1.426 1.509.702l33.993-29.091c.84-.72.837-1.894.016-2.606z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M72.158 119.649c-11.578-11.362-2.954-31.242 12.729-31.005c15.682.238 71.904.238 87.357.238c15.453 0 23.578 19.299 12.17 30.815c-2.119 2.14-4.694 4.746-7.524 7.612c-12.403 12.561-29.725 30.102-35.25 35.125c-6.786 6.17-18.794 7.154-26.124-.31c-4.841-4.929-17.145-16.914-28.28-27.759c-5.727-5.579-11.145-10.856-15.078-14.716m14.225-13.763c-2.214-.011-2.735 1.24-1.173 2.794l40.888 40.324c1.555 1.547 2.705 1.547 4.283 0l39.596-39.868c1.578-1.548 1.065-2.812-1.145-2.823z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M136.916 184.407c11.363 11.578 31.242 2.954 31.005-12.728c-.237-15.682-.237-71.898-.237-87.355v-.003c0-15.453-17.987-23.243-29.502-11.835c-3.95 3.913-9.505 9.312-15.348 14.99c-11.192 10.879-23.441 22.784-27.494 27.242c-6.169 6.786-8.362 19.001-.898 26.332c4.938 4.85 16.962 17.194 27.823 28.345c5.555 5.703 10.806 11.094 14.651 15.012m13.777-13.569c.011 2.214-1.465 1.562-3.02 0v-.001l-40.04-39.91c-1.548-1.554-1.561-4.107-.013-5.684l41.117-41.082c1.548-1.578 1.944-.605 1.956 1.605z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M117.846 184.606c-11.362 11.579-29.419 2.755-29.182-12.927h-.001c.237-15.683.237-71.905.237-87.358c0-15.453 18.3-23.578 29.817-12.17c11.515 11.409 37.257 38.828 43.427 45.614c6.168 6.786 7.463 14.243 0 21.574c-4.855 4.769-17.331 17.589-28.586 29.155c-6.047 6.214-11.741 12.066-15.712 16.112m-12.381-15.368c-.012 2.217 1.246 2.732 2.808 1.162l39.283-39.46a4.055 4.055 0 0 0 .03-5.693l-38.874-39.606c-1.54-1.568-2.808-1.058-2.82 1.152z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M183.427 136.897c11.578 11.362 3.954 31.242-11.729 31.004c-15.682-.237-71.904-.237-87.357-.237c-15.453 0-23.261-18.467-11.853-29.982c2.767-2.793 6.33-6.412 10.236-10.382c12.203-12.397 27.762-28.205 32.902-32.878c6.786-6.169 18.112-7.464 25.444 0c7.331 7.464 30.779 31.112 42.357 42.475m-14.608 14.261c2.214.011 2.735-2.24 1.173-3.795v.001l-39.486-39.31c-1.554-1.548-4.107-1.561-5.684-.014l-39.597 39.868c-1.578 1.548-1.064 2.812 1.146 2.823z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M39.723 193.132c.283-8.428-.284-102.83-.284-112.335c0-9.505 8.156-17.14 17.993-17.327c9.837-.188 55.149.172 55.149.172c1.103.008 2.644.63 3.45 1.39l34.483 32.477c1.206 1.136 3.534 2.06 5.196 2.06h46.823c7.174 0 12.552 5.312 12 11.865c0 0 1.2 18.088 1 36.787c-.2 18.699 0 43.91 0 48.408c0 4.497-8.46 11.442-15.84 10.883c-7.38-.56-132.898.813-142.362.813c-9.464 0-17.892-6.765-17.608-15.193m18.053-110.71l.235 16.507a2.03 2.03 0 0 0 2.024 2.002l65.144-.08c2.208-.003 2.696-1.233 1.085-2.748l-17.935-16.859c-.808-.76-2.36-1.366-3.466-1.354l-45.114.505a1.995 1.995 0 0 0-1.973 2.027m4.046 109.177l135.996.73a1.952 1.952 0 0 0 1.98-1.98l-.922-71.163a2.01 2.01 0 0 0-2.018-1.977l-135.05.894a2.011 2.011 0 0 0-1.992 2.016v69.466c0 1.1.899 2.008 2.006 2.014"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M48.186 92.137c0-8.392 6.49-14.89 16.264-14.89s29.827-.225 29.827-.225s-.306-6.99-.306-15.88c0-8.888 7.954-14.96 17.49-14.96c9.538 0 56.786.401 61.422.401c4.636 0 8.397 1.719 13.594 5.67c5.196 3.953 13.052 10.56 16.942 14.962c3.89 4.402 5.532 6.972 5.532 10.604c0 3.633 0 76.856-.06 85.34c-.059 8.485-7.877 14.757-17.134 14.881c-9.257.124-29.135.124-29.135.124s.466 6.275.466 15.15s-8.106 15.811-17.317 16.056c-9.21.245-71.944-.49-80.884-.245c-8.94.245-16.975-6.794-16.975-15.422s.274-93.175.274-101.566m16.734 3.946l-1.152 92.853a3.96 3.96 0 0 0 3.958 4.012l73.913.22a3.865 3.865 0 0 0 3.91-3.978l-.218-8.892a1.988 1.988 0 0 0-2.046-1.953s-21.866.64-31.767.293c-9.902-.348-16.672-6.807-16.675-15.516c-.003-8.709.003-69.142.003-69.142a1.989 1.989 0 0 0-2.007-1.993l-23.871.082a4.077 4.077 0 0 0-4.048 4.014m106.508-35.258c-1.666-1.45-3.016-.84-3.016 1.372v17.255c0 1.106.894 2.007 1.997 2.013l20.868.101c2.204.011 2.641-1.156.976-2.606zm-57.606.847a2.002 2.002 0 0 0-2.02 1.988l-.626 96.291a2.968 2.968 0 0 0 2.978 2.997l75.2-.186a2.054 2.054 0 0 0 2.044-2.012l1.268-62.421a1.951 1.951 0 0 0-1.96-2.004s-26.172.042-30.783.042c-4.611 0-7.535-2.222-7.535-6.482S152.3 63.92 152.3 63.92a2.033 2.033 0 0 0-2.015-2.018z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cpu(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m76.21 64l.164-16.15h15.24V64l13.847-.443V47.85l14.64-.439v16.145h14.486V47.85h15.324v15.706l15.313.368V47.85h14.643l.14 16.074l4.656.078a7.99 7.99 0 0 1 7.855 8.127l-.069 3.991h16.113v15.764H192.56l.094 13.812h15.907L208 121.091h-15.744l.202 14.523l15.542.64v15.227h-15.625l.184 13.3l15.441.17v14.94h-15.66l-.126 3.94c-.142 4.416-3.837 8.054-8.256 8.127l-4.043.066l-.048 16.555h-14.643v-16.097H149.91v16.097h-15.324v-16.097h-14.485v16.097H106.26v-16.097H91.613v16.097H76.374l-.164-16.753h-4.067a7.996 7.996 0 0 1-7.996-7.994v-3.941h-16.36v-14.94h16.36l.236-13.47H47.787v-15.227h16.36V121.09h-16.36v-14.345h16.847V91.99H47.787V77.063h16.596v-4.92c0-4.416 3.58-8.04 8.003-8.095zm3.722 18.052l.579 91.96c.007 1.115.909 2 2.014 1.995l91.836-.375a2.01 2.01 0 0 0 2.005-2.008l.15-91.771a1.984 1.984 0 0 0-1.993-1.996l-92.609.187a1.99 1.99 0 0 0-1.982 2.008"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cutter(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M70.526 8.975L215.553 155l-90.48 89.973L68.86 187.58zM85.246 47l1.018 134.104L125.74 223l55.346-55.78l-54.06-54.22l11.872-12.917z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalColon(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M76.846 80a4 4 0 0 1 4.002 4.002V107.6a4.002 4.002 0 0 1-4.002 4.003H53.001A4 4 0 0 1 49 107.599V84.002A4.002 4.002 0 0 1 53.001 80zm0 64a4 4 0 0 1 4.002 4.002V171.6a4.002 4.002 0 0 1-4.002 4.003H53.001A4 4 0 0 1 49 171.599v-23.597A4.002 4.002 0 0 1 53.001 144z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalDot(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M75.846 176a4 4 0 0 1 4.002 4.002V203.6a4.002 4.002 0 0 1-4.002 4.003H52.001A4 4 0 0 1 48 203.599v-23.597A4.002 4.002 0 0 1 52.001 176z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalEight(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M163.819 107.882c.006 2.204 1.314 5.22 2.922 6.738l6.463 6.098c1.608 1.518 4.309 1.626 6.022.25l9.525-7.65c1.718-1.38 3.113-4.283 3.116-6.496l.078-71.27c.002-2.208-1.303-2.78-2.915-1.277l-22.432 20.91c-1.612 1.502-2.914 4.516-2.908 6.71zm-71.78 11.642c1.561-1.56 4.618-2.841 6.834-2.863l56.433-.554c2.213-.022 5.297 1.204 6.887 2.738l7.074 6.82c1.59 1.535 1.51 3.937-.175 5.362l-6.575 5.564c-1.687 1.427-4.85 2.589-7.047 2.595l-55.465.16c-2.205.006-5.367-1.141-7.062-2.564l-6.812-5.717c-1.695-1.423-1.805-3.837-.241-5.4zm-16.746 16.58c1.741-1.359 4.574-1.377 6.334-.035l7.282 5.549c1.757 1.34 3.179 4.212 3.175 6.417l-.08 45.714c-.003 2.205-1.319 5.196-2.95 6.692l-22.11 20.28c-1.626 1.491-2.944.913-2.944-1.301v-70.501c0-2.21 1.411-5.103 3.152-6.461zm81.946 66.067c2.211.003 5.325 1.214 6.95 2.7l24.4 22.31c1.627 1.488 1.156 2.695-1.053 2.697l-118.63.092c-2.21.001-2.693-1.232-1.088-2.747l23.697-22.386c1.608-1.519 4.699-2.748 6.916-2.745l58.808.08zm0-149.283c2.211-.003 5.325-1.213 6.95-2.699l24.4-22.31c1.627-1.488 1.157-2.694-1.05-2.694H68.996c-2.207 0-2.694 1.23-1.094 2.74l23.623 22.307c1.602 1.514 4.689 2.738 6.906 2.735zm7.127 95.528c.015-2.212 1.4-5.156 3.1-6.58l6.35-5.323c1.698-1.422 4.56-1.565 6.387-.322l8.302 5.646c1.83 1.244 3.326 4.042 3.342 6.258l.51 70.57c.016 2.21-1.32 2.83-2.992 1.374l-22.313-19.423c-1.668-1.452-3.009-4.412-2.994-6.635zM89.19 56.053c1.618 1.5 2.93 4.501 2.93 6.714v44.353c0 2.209-1.275 5.253-2.848 6.8l-7.356 7.232c-1.572 1.546-4.225 1.666-5.933.26l-8.897-7.325c-1.704-1.403-3.086-4.335-3.086-6.53v-70.85c0-2.205 1.319-2.77 2.93-1.277z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalFive(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M92.039 119.524c1.561-1.56 4.618-2.841 6.834-2.863l56.433-.554c2.213-.022 5.297 1.204 6.887 2.738l7.074 6.82c1.59 1.535 1.51 3.937-.175 5.362l-6.575 5.564c-1.687 1.427-4.85 2.589-7.047 2.595l-55.465.16c-2.205.006-5.367-1.141-7.062-2.564l-6.812-5.717c-1.695-1.423-1.805-3.837-.241-5.4zm65.2 82.647c2.211.003 5.325 1.214 6.95 2.7l24.4 22.31c1.627 1.488 1.156 2.695-1.053 2.697l-118.63.092c-2.21.001-2.693-1.232-1.088-2.747l23.697-22.386c1.608-1.519 4.699-2.748 6.916-2.745l58.808.08zm0-149.283c2.211-.003 5.325-1.213 6.95-2.699l24.4-22.31c1.627-1.488 1.157-2.694-1.05-2.694H68.996c-2.207 0-2.694 1.23-1.094 2.74l23.623 22.307c1.602 1.514 4.689 2.738 6.906 2.735zm7.127 95.528c.015-2.212 1.4-5.156 3.1-6.58l6.35-5.323c1.698-1.422 4.56-1.565 6.387-.322l8.302 5.646c1.83 1.244 3.326 4.042 3.342 6.258l.51 70.57c.016 2.21-1.32 2.83-2.992 1.374l-22.313-19.423c-1.668-1.452-3.009-4.412-2.994-6.635zM89.19 56.053c1.618 1.5 2.93 4.501 2.93 6.714v44.353c0 2.209-1.275 5.253-2.848 6.8l-7.356 7.232c-1.572 1.546-4.225 1.666-5.933.26l-8.897-7.325c-1.704-1.403-3.086-4.335-3.086-6.53v-70.85c0-2.205 1.319-2.77 2.93-1.277z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalFour(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M163.819 107.882c.006 2.204 1.314 5.22 2.922 6.738l6.463 6.098c1.608 1.518 4.309 1.626 6.022.25l9.525-7.65c1.718-1.38 3.113-4.283 3.116-6.496l.078-71.27c.002-2.208-1.303-2.78-2.915-1.277l-22.432 20.91c-1.612 1.502-2.914 4.516-2.908 6.71zm-71.78 11.642c1.561-1.56 4.618-2.841 6.834-2.863l56.433-.554c2.213-.022 5.297 1.204 6.887 2.738l7.074 6.82c1.59 1.535 1.51 3.937-.175 5.362l-6.575 5.564c-1.687 1.427-4.85 2.589-7.047 2.595l-55.465.16c-2.205.006-5.367-1.141-7.062-2.564l-6.812-5.717c-1.695-1.423-1.805-3.837-.241-5.4zm72.327 28.892c.015-2.212 1.4-5.156 3.1-6.58l6.35-5.323c1.698-1.422 4.56-1.565 6.387-.322l8.302 5.646c1.83 1.244 3.326 4.042 3.342 6.258l.51 70.57c.016 2.21-1.32 2.83-2.992 1.374l-22.313-19.423c-1.668-1.452-3.009-4.412-2.994-6.635zM89.19 56.053c1.618 1.5 2.93 4.501 2.93 6.714v44.353c0 2.209-1.275 5.253-2.848 6.8l-7.356 7.232c-1.572 1.546-4.225 1.666-5.933.26l-8.897-7.325c-1.704-1.403-3.086-4.335-3.086-6.53v-70.85c0-2.205 1.319-2.77 2.93-1.277z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalNine(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M163.819 107.882c.006 2.204 1.314 5.22 2.922 6.738l6.463 6.098c1.608 1.518 4.309 1.626 6.022.25l9.525-7.65c1.718-1.38 3.113-4.283 3.116-6.496l.078-71.27c.002-2.208-1.303-2.78-2.915-1.277l-22.432 20.91c-1.612 1.502-2.914 4.516-2.908 6.71zm-71.78 11.642c1.561-1.56 4.618-2.841 6.834-2.863l56.433-.554c2.213-.022 5.297 1.204 6.887 2.738l7.074 6.82c1.59 1.535 1.51 3.937-.175 5.362l-6.575 5.564c-1.687 1.427-4.85 2.589-7.047 2.595l-55.465.16c-2.205.006-5.367-1.141-7.062-2.564l-6.812-5.717c-1.695-1.423-1.805-3.837-.241-5.4zm65.2 82.647c2.211.003 5.325 1.214 6.95 2.7l24.4 22.31c1.627 1.488 1.156 2.695-1.053 2.697l-118.63.092c-2.21.001-2.693-1.232-1.088-2.747l23.697-22.386c1.608-1.519 4.699-2.748 6.916-2.745l58.808.08zm0-149.283c2.211-.003 5.325-1.213 6.95-2.699l24.4-22.31c1.627-1.488 1.157-2.694-1.05-2.694H68.996c-2.207 0-2.694 1.23-1.094 2.74l23.623 22.307c1.602 1.514 4.689 2.738 6.906 2.735zm7.127 95.528c.015-2.212 1.4-5.156 3.1-6.58l6.35-5.323c1.698-1.422 4.56-1.565 6.387-.322l8.302 5.646c1.83 1.244 3.326 4.042 3.342 6.258l.51 70.57c.016 2.21-1.32 2.83-2.992 1.374l-22.313-19.423c-1.668-1.452-3.009-4.412-2.994-6.635zM89.19 56.053c1.618 1.5 2.93 4.501 2.93 6.714v44.353c0 2.209-1.275 5.253-2.848 6.8l-7.356 7.232c-1.572 1.546-4.225 1.666-5.933.26l-8.897-7.325c-1.704-1.403-3.086-4.335-3.086-6.53v-70.85c0-2.205 1.319-2.77 2.93-1.277z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalOne(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M163.819 107.882c.006 2.204 1.314 5.22 2.922 6.738l6.463 6.098c1.608 1.518 4.309 1.626 6.022.25l9.525-7.65c1.718-1.38 3.113-4.283 3.116-6.496l.078-71.27c.002-2.208-1.303-2.78-2.915-1.277l-22.432 20.91c-1.612 1.502-2.914 4.516-2.908 6.71zm.547 40.534c.015-2.212 1.4-5.156 3.1-6.58l6.35-5.323c1.698-1.422 4.56-1.565 6.387-.322l8.302 5.646c1.83 1.244 3.326 4.042 3.342 6.258l.51 70.57c.016 2.21-1.32 2.83-2.992 1.374l-22.313-19.423c-1.668-1.452-3.009-4.412-2.994-6.635z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalSeven(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M163.819 107.882c.006 2.204 1.314 5.22 2.922 6.738l6.463 6.098c1.608 1.518 4.309 1.626 6.022.25l9.525-7.65c1.718-1.38 3.113-4.283 3.116-6.496l.078-71.27c.002-2.208-1.303-2.78-2.915-1.277l-22.432 20.91c-1.612 1.502-2.914 4.516-2.908 6.71zm-6.58-54.994c2.211-.003 5.325-1.213 6.95-2.699l24.4-22.31c1.627-1.488 1.157-2.694-1.05-2.694H68.996c-2.207 0-2.694 1.23-1.094 2.74l23.623 22.307c1.602 1.514 4.689 2.738 6.906 2.735zm7.127 95.528c.015-2.212 1.4-5.156 3.1-6.58l6.35-5.323c1.698-1.422 4.56-1.565 6.387-.322l8.302 5.646c1.83 1.244 3.326 4.042 3.342 6.258l.51 70.57c.016 2.21-1.32 2.83-2.992 1.374l-22.313-19.423c-1.668-1.452-3.009-4.412-2.994-6.635z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalSix(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M92.039 119.524c1.561-1.56 4.618-2.841 6.834-2.863l56.433-.554c2.213-.022 5.297 1.204 6.887 2.738l7.074 6.82c1.59 1.535 1.51 3.937-.175 5.362l-6.575 5.564c-1.687 1.427-4.85 2.589-7.047 2.595l-55.465.16c-2.205.006-5.367-1.141-7.062-2.564l-6.812-5.717c-1.695-1.423-1.805-3.837-.241-5.4zm-16.746 16.58c1.741-1.359 4.574-1.377 6.334-.035l7.282 5.549c1.757 1.34 3.179 4.212 3.175 6.417l-.08 45.714c-.003 2.205-1.319 5.196-2.95 6.692l-22.11 20.28c-1.626 1.491-2.944.913-2.944-1.301v-70.501c0-2.21 1.411-5.103 3.152-6.461zm81.946 66.067c2.211.003 5.325 1.214 6.95 2.7l24.4 22.31c1.627 1.488 1.156 2.695-1.053 2.697l-118.63.092c-2.21.001-2.693-1.232-1.088-2.747l23.697-22.386c1.608-1.519 4.699-2.748 6.916-2.745l58.808.08zm0-149.283c2.211-.003 5.325-1.213 6.95-2.699l24.4-22.31c1.627-1.488 1.157-2.694-1.05-2.694H68.996c-2.207 0-2.694 1.23-1.094 2.74l23.623 22.307c1.602 1.514 4.689 2.738 6.906 2.735zm7.127 95.528c.015-2.212 1.4-5.156 3.1-6.58l6.35-5.323c1.698-1.422 4.56-1.565 6.387-.322l8.302 5.646c1.83 1.244 3.326 4.042 3.342 6.258l.51 70.57c.016 2.21-1.32 2.83-2.992 1.374l-22.313-19.423c-1.668-1.452-3.009-4.412-2.994-6.635zM89.19 56.053c1.618 1.5 2.93 4.501 2.93 6.714v44.353c0 2.209-1.275 5.253-2.848 6.8l-7.356 7.232c-1.572 1.546-4.225 1.666-5.933.26l-8.897-7.325c-1.704-1.403-3.086-4.335-3.086-6.53v-70.85c0-2.205 1.319-2.77 2.93-1.277z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalThree(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M163.819 107.882c.006 2.204 1.314 5.22 2.922 6.738l6.463 6.098c1.608 1.518 4.309 1.626 6.022.25l9.525-7.65c1.718-1.38 3.113-4.283 3.116-6.496l.078-71.27c.002-2.208-1.303-2.78-2.915-1.277l-22.432 20.91c-1.612 1.502-2.914 4.516-2.908 6.71zm-71.78 11.642c1.561-1.56 4.618-2.841 6.834-2.863l56.433-.554c2.213-.022 5.297 1.204 6.887 2.738l7.074 6.82c1.59 1.535 1.51 3.937-.175 5.362l-6.575 5.564c-1.687 1.427-4.85 2.589-7.047 2.595l-55.465.16c-2.205.006-5.367-1.141-7.062-2.564l-6.812-5.717c-1.695-1.423-1.805-3.837-.241-5.4zm65.2 82.647c2.211.003 5.325 1.214 6.95 2.7l24.4 22.31c1.627 1.488 1.156 2.695-1.053 2.697l-118.63.092c-2.21.001-2.693-1.232-1.088-2.747l23.697-22.386c1.608-1.519 4.699-2.748 6.916-2.745l58.808.08zm0-149.283c2.211-.003 5.325-1.213 6.95-2.699l24.4-22.31c1.627-1.488 1.157-2.694-1.05-2.694H68.996c-2.207 0-2.694 1.23-1.094 2.74l23.623 22.307c1.602 1.514 4.689 2.738 6.906 2.735zm7.127 95.528c.015-2.212 1.4-5.156 3.1-6.58l6.35-5.323c1.698-1.422 4.56-1.565 6.387-.322l8.302 5.646c1.83 1.244 3.326 4.042 3.342 6.258l.51 70.57c.016 2.21-1.32 2.83-2.992 1.374l-22.313-19.423c-1.668-1.452-3.009-4.412-2.994-6.635z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalTwo(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M163.819 107.882c.006 2.204 1.314 5.22 2.922 6.738l6.463 6.098c1.608 1.518 4.309 1.626 6.022.25l9.525-7.65c1.718-1.38 3.113-4.283 3.116-6.496l.078-71.27c.002-2.208-1.303-2.78-2.915-1.277l-22.432 20.91c-1.612 1.502-2.914 4.516-2.908 6.71zm-71.78 11.642c1.561-1.56 4.618-2.841 6.834-2.863l56.433-.554c2.213-.022 5.297 1.204 6.887 2.738l7.074 6.82c1.59 1.535 1.51 3.937-.175 5.362l-6.575 5.564c-1.687 1.427-4.85 2.589-7.047 2.595l-55.465.16c-2.205.006-5.367-1.141-7.062-2.564l-6.812-5.717c-1.695-1.423-1.805-3.837-.241-5.4zm-16.746 16.58c1.741-1.359 4.574-1.377 6.334-.035l7.282 5.549c1.757 1.34 3.179 4.212 3.175 6.417l-.08 45.714c-.003 2.205-1.319 5.196-2.95 6.692l-22.11 20.28c-1.626 1.491-2.944.913-2.944-1.301v-70.501c0-2.21 1.411-5.103 3.152-6.461zm81.946 66.067c2.211.003 5.325 1.214 6.95 2.7l24.4 22.31c1.627 1.488 1.156 2.695-1.053 2.697l-118.63.092c-2.21.001-2.693-1.232-1.088-2.747l23.697-22.386c1.608-1.519 4.699-2.748 6.916-2.745l58.808.08zm0-149.283c2.211-.003 5.325-1.213 6.95-2.699l24.4-22.31c1.627-1.488 1.157-2.694-1.05-2.694H68.996c-2.207 0-2.694 1.23-1.094 2.74l23.623 22.307c1.602 1.514 4.689 2.738 6.906 2.735z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DigitalZero(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M163.819 107.882c.006 2.204 1.314 5.22 2.922 6.738l6.463 6.098c1.608 1.518 4.309 1.626 6.022.25l9.525-7.65c1.718-1.38 3.113-4.283 3.116-6.496l.078-71.27c.002-2.208-1.303-2.78-2.915-1.277l-22.432 20.91c-1.612 1.502-2.914 4.516-2.908 6.71zm-88.526 28.222c1.741-1.359 4.574-1.377 6.334-.035l7.282 5.549c1.757 1.34 3.179 4.212 3.175 6.417l-.08 45.714c-.003 2.205-1.319 5.196-2.95 6.692l-22.11 20.28c-1.626 1.491-2.944.913-2.944-1.301v-70.501c0-2.21 1.411-5.103 3.152-6.461zm81.946 66.067c2.211.003 5.325 1.214 6.95 2.7l24.4 22.31c1.627 1.488 1.156 2.695-1.053 2.697l-118.63.092c-2.21.001-2.693-1.232-1.088-2.747l23.697-22.386c1.608-1.519 4.699-2.748 6.916-2.745l58.808.08zm0-149.283c2.211-.003 5.325-1.213 6.95-2.699l24.4-22.31c1.627-1.488 1.157-2.694-1.05-2.694H68.996c-2.207 0-2.694 1.23-1.094 2.74l23.623 22.307c1.602 1.514 4.689 2.738 6.906 2.735zm7.127 95.528c.015-2.212 1.4-5.156 3.1-6.58l6.35-5.323c1.698-1.422 4.56-1.565 6.387-.322l8.302 5.646c1.83 1.244 3.326 4.042 3.342 6.258l.51 70.57c.016 2.21-1.32 2.83-2.992 1.374l-22.313-19.423c-1.668-1.452-3.009-4.412-2.994-6.635zM89.19 56.053c1.618 1.5 2.93 4.501 2.93 6.714v44.353c0 2.209-1.275 5.253-2.848 6.8l-7.356 7.232c-1.572 1.546-4.225 1.666-5.933.26l-8.897-7.325c-1.704-1.403-3.086-4.335-3.086-6.53v-70.85c0-2.205 1.319-2.77 2.93-1.277z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diskio(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M52.512 90.498a3.994 3.994 0 0 0-3.994 4v8.395c0 2.209 1.786 4 3.994 4H178.17c1.105 0 1.359.625.565 1.398l-12.907 12.555c-1.58 1.537-1.62 4.076-.084 5.675l5.174 5.39a3.865 3.865 0 0 0 5.58.037l31.342-31.919a2.032 2.032 0 0 0-.011-2.853l-31.315-31.463c-1.557-1.565-4.019-1.498-5.499.15l-5.368 5.975a4.037 4.037 0 0 0 .278 5.677l12.765 11.633c.818.746.589 1.35-.52 1.35zm151.738 60a3.994 3.994 0 0 1 3.994 4v8.395c0 2.209-1.786 4-3.994 4H78.592c-1.106 0-1.36.625-.565 1.398l12.907 12.555c1.58 1.537 1.619 4.076.084 5.675l-5.174 5.39a3.865 3.865 0 0 1-5.581.037l-31.34-31.919a2.032 2.032 0 0 1 .01-2.853l31.315-31.463c1.557-1.565 4.019-1.498 5.498.15l5.369 5.975a4.037 4.037 0 0 1-.279 5.677l-12.764 11.633c-.819.746-.589 1.35.52 1.35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drumpad(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M32 64.001C32 46.328 46.333 32 64.001 32H192c17.672 0 32 14.333 32 32.001V192c0 17.672-14.333 32-32.001 32H64c-17.672 0-32-14.333-32-32.001V64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Duplicate(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M47.81 91.725c0-8.328 6.539-15.315 15.568-15.33c9.03-.016 14.863.015 14.863.015s-.388-8.9-.388-15.978c0-7.08 6.227-14.165 15.262-14.165s92.802-.26 101.297.37c8.495.63 15.256 5.973 15.256 14.567c0 8.594-.054 93.807-.054 101.7c0 7.892-7.08 15.063-15.858 15.162c-8.778.1-14.727-.1-14.727-.1s.323 9.97.323 16.094c0 6.123-7.12 15.016-15.474 15.016s-93.117.542-101.205.542c-8.088 0-15.552-7.116-15.207-15.987c.345-8.871.345-93.58.345-101.906zm46.06-28.487l-.068 98.164c0 1.096.894 1.99 1.999 1.984l95.555-.51a2.007 2.007 0 0 0 1.998-2.01l-.064-97.283a2.01 2.01 0 0 0-2.01-2.007l-95.4-.326a1.99 1.99 0 0 0-2.01 1.988M63.268 95.795l.916 96.246a2.007 2.007 0 0 0 2.02 1.982l94.125-.715a3.976 3.976 0 0 0 3.953-4.026l-.137-11.137s-62.877.578-71.054.578s-15.438-7.74-15.438-16.45c0-8.71.588-68.7.588-68.7c.01-1.1-.874-1.99-1.976-1.975l-9.027.13a4.025 4.025 0 0 0-3.97 4.067"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m38.032 127.602l76.405-78.588c3.074-3.162 9.157-5.737 13.573-5.737h88.392a7.958 7.958 0 0 1 7.966 7.996l-.243 63.237c-.017 4.412-2.563 10.53-5.686 13.655l-74.118 74.176c-3.121 3.124-9.239 5.659-13.66 5.659H40.46c-4.425 0-8.005-3.577-8.005-7.99v-58.682c0-4.414 2.497-10.558 5.577-13.726M120.002 144h-63.98c-4.405 0-7.997 3.578-7.997 7.993v32.09a7.992 7.992 0 0 0 7.998 7.993h63.98c4.405 0 7.997-3.578 7.997-7.993v-32.09a7.992 7.992 0 0 0-7.998-7.993m29.638-6.036c-3.115 3.14-5.64 9.268-5.64 13.685v21.72c0 4.418 2.476 5.42 5.542 2.227l53.585-55.823c3.06-3.188 5.542-9.36 5.542-13.771V80.459c0-1.105-.63-1.365-1.41-.578zm-18.046-78.59c-1.66.03-3.966 1-5.143 2.16l-61.065 60.197c-3.146 3.102-2.114 5.616 2.306 5.616h60.119c4.42 0 10.559-2.506 13.71-5.596l63.519-62.277c.788-.773.525-1.384-.577-1.365z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ffwd(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M48.567 63.148c0-11.962 9.145-18.903 15.545-10.988c6.4 7.916 52.713 60.673 52.713 60.673c.728.83 1.32.599 1.322-.504c0 0 .085-39.029.085-50.228c0-11.2 9.93-18.233 16.073-10.137c6.143 8.095 56.836 65.487 56.836 65.487c.727.824 1.328.593 1.341-.499c0 0 .747-59.858.577-62.396c-.17-2.538 1.22-4.792 3.45-4.792h8.193c3.655 0 3.52.682 3.52 3.858c0 3.177-.394 147.133-.459 149.861c-.065 2.728-1.06 4.454-3.22 4.454c-2.158 0-6.252 0-8.33-.007c-2.078-.006-3.087-1.896-3.087-4.032s-.617-64.526-.617-64.526c-.01-1.105-.607-1.318-1.327-.48c0 0-52.535 61.05-56.586 65.857c-4.052 4.807-16.516 1.394-16.516-10.433c0-11.826.37-49.528.37-49.528c.01-1.103-.563-1.322-1.282-.488c0 0-48.168 55.896-52.355 60.404c-4.187 4.509-16.027 1.107-16.027-10.347c0-11.454-.219-119.248-.219-131.21zM63.416 78.97l.095 99.405c.001 1.102.585 1.314 1.304.474l42.11-49.234c.72-.842.712-2.195-.015-3.024L64.73 78.477c-.725-.828-1.315-.609-1.314.494zm71.046-1.025c-.721-.842-1.31-.622-1.314.467l-.409 100.404c-.004 1.1.581 1.315 1.31.477l42.64-48.986c.729-.836.735-2.196.013-3.04z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterBandpass(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25.344 180.07a4.008 4.008 0 0 1 3.997-4.01h16.996c6.631 0 14.517-4.753 17.611-10.614l47.246-89.476c9.282-17.579 24.376-17.602 33.72-.042l47.637 89.532c3.115 5.855 11.007 10.6 17.65 10.6h16.489a4.01 4.01 0 0 1 4.001 4.01v8.809c0 2.214-1.8 4.009-4.007 4.009h-24.49c-8.838 0-19.361-6.32-23.513-14.133L136.446 99.28c-4.665-8.778-12.228-8.772-16.887 0l-42.21 79.475c-4.145 7.805-14.667 14.133-23.508 14.133h-24.49a4.012 4.012 0 0 1-4.007-4.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterBell(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128.012 64.105c-6.857.003-13.712 3.567-18.938 10.69l-42.357 37.28c-5.36 3.36-11.368 7.187-17.19 7.03H32.04L29.01 119c-2.149.367-3.962 2.35-4.009 4.102v9.902c.62 2.571 1.898 3.651 4.01 4.101l3.03-.105h17.487c5.908.484 13.043 3.495 17.19 7.031l42.357 37.28c10.452 14.246 27.417 14.253 37.89.015l42.483-36.418c5.028-4.31 14.472-7.803 21.106-7.803c5.485.294 10.97.596 16.455.895c2.208 0 3.992-2.79 3.992-4.996v-9.902c0-2.207-1.784-4.997-3.992-4.997c-5.485.296-10.97.605-16.455.895c-6.634 0-16.078-3.495-21.106-7.805L146.965 74.78c-5.237-7.118-12.097-10.676-18.953-10.674zm.21 18.702c3.143-.005 6.289 1.136 8.692 3.42l41.824 38.753c1.11 1.055 2.407 2.085 3.824 3.073c-1.417.987-2.714 2.015-3.824 3.07l-41.824 39.756c-4.807 4.569-12.577 4.56-17.37-.031l-42.47-40.676a28.2 28.2 0 0 0-2.513-2.12a28.194 28.194 0 0 0 2.513-2.118l42.47-39.676c2.397-2.296 5.536-3.447 8.679-3.451z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterBypass(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25 67.995A4.002 4.002 0 0 1 29.004 64H227a4 4 0 0 1 4 3.995v10.01A3.99 3.99 0 0 1 227 82H29a4 4 0 0 1-4-3.995z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterHighpass(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M231.007 68.729c0-2.206-1.787-4.995-4.007-4.995h-85.499c-6.466 0-19.531 7.705-22.66 15.97l-55.92 85.647c-3.624 5.55-11.93 10.05-18.559 10.05H28.167c-2.206 0-3.994 2.787-3.994 5.007v8.985a4.005 4.005 0 0 0 3.998 4.007h22.713c8.832 0 20.495-8.703 23.588-16.987l56.167-84.189c3.68-5.517 12.04-9.99 18.668-9.99h77.695c2.212 0 4.005-2.797 4.005-4.994v-8.51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterLowpass(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M24.22 67.796a3.995 3.995 0 0 1 4.008-3.991h85.498c8.834 0 19.732 6.112 24.345 13.657l53.76 87.936c3.46 5.66 11.628 10.247 18.256 10.247h16.718a3.996 3.996 0 0 1 3.994 4.007v8.985a4.007 4.007 0 0 1-4.007 4.008h-24.7c-8.835 0-19.709-6.13-24.283-13.683l-52.324-86.4c-3.43-5.665-11.577-10.257-18.202-10.257H28.214a3.995 3.995 0 0 1-3.993-3.992V67.796z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterNotch(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25.101 77.628a4.008 4.008 0 0 0 3.997 4.01h16.996c6.632 0 13.927 5.01 16.3 11.202l52.724 85.231c7.115 18.564 18.693 18.571 25.857.025L193.91 92.84c2.39-6.187 9.693-11.202 16.336-11.202h16.49a4.01 4.01 0 0 0 4-4.01V68.82a4 4 0 0 0-3.994-4.009h-23.508c-8.835 0-18.547 6.702-21.69 14.962l-47.147 73.852c-3.533 9.287-9.217 9.262-12.694-.051L75.2 79.805C72.108 71.524 62.44 64.81 53.6 64.81H29.11a4.012 4.012 0 0 0-4.008 4.01v8.808z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterRezHighpass(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M230.704 99.2a4.004 4.004 0 0 0-4.01-3.995h-50.981c-2.215 0-5.212-1.327-6.693-2.964L155.289 77.08c-17.795-19.65-41.628-16.256-53.234 7.58l-38.736 79.557C60.42 170.172 52.705 175 46.077 175H29.359a3.996 3.996 0 0 0-3.994 3.995v10.01A4 4 0 0 0 29.372 193h24.7c8.835 0 19.208-6.395 23.174-14.293l43.645-86.914c3.964-7.894 12.233-9.228 18.473-2.974l17.184 17.219c3.123 3.13 9.242 5.667 13.647 5.667H226.7a4.005 4.005 0 0 0 4.004-3.994v-8.512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterRezLowpass(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25.365 99.2a4.004 4.004 0 0 1 4.01-3.995h50.981c2.214 0 5.28-1.248 6.86-2.802l17.551-17.253c18.907-18.584 42.923-13.987 53.634 10.248l34.746 78.62c2.68 6.065 10.217 10.982 16.845 10.982h16.718a3.996 3.996 0 0 1 3.994 3.995v10.01a4 4 0 0 1-4.007 3.995h-24.7c-8.835 0-19.03-6.489-22.775-14.497l-40.444-86.506c-3.744-8.007-12.11-9.714-18.687-3.815l-20.267 18.179c-3.291 2.951-9.545 5.344-13.95 5.344H29.368a4.005 4.005 0 0 1-4.004-3.994v-8.512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterShelvingHi(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M203.037 63c-8.835 0-21.563 4.505-28.433 10.063l-47.528 38.445c-3.5 2.83-9.024 5.241-14.289 6.492H29.38a4.01 4.01 0 0 0-3.99 3.705v11.643a4.01 4.01 0 0 0 3.99 3.705h83.408c5.265 1.251 10.79 3.661 14.29 6.492l47.527 38.445c6.87 5.558 19.598 10.063 28.433 10.063h23.51a4 4 0 0 0 3.994-4.008v-8.984a4.003 4.003 0 0 0-4.008-4.008H209.04c-6.617 0-16.202-3.341-21.389-7.463l-43.906-34.897c-1.846-1.713-3.845-3.666-5.682-5.168c1.972-1.599 3.762-3.375 5.682-5.166l43.906-34.896C192.837 83.34 202.422 80 209.04 80h17.493a4.003 4.003 0 0 0 4.008-4.008v-8.984A4 4 0 0 0 226.547 63z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterShelvingLo(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M52.893 63c8.835 0 21.563 4.505 28.433 10.063l47.527 38.445c3.5 2.83 9.024 5.241 14.29 6.492h83.408a4.01 4.01 0 0 1 3.99 3.705v11.643a4.01 4.01 0 0 1-3.99 3.705h-83.408c-5.266 1.251-10.79 3.661-14.29 6.492L81.326 181.99c-6.87 5.558-19.598 10.063-28.433 10.063h-23.51a4 4 0 0 1-3.994-4.008v-8.984a4.003 4.003 0 0 1 4.007-4.008h17.495c6.617 0 16.201-3.341 21.388-7.463l43.906-34.897c1.846-1.713 3.845-3.666 5.682-5.168c-1.97-1.599-3.762-3.375-5.682-5.166L68.28 87.463C63.092 83.34 53.508 80 46.891 80H29.396a4.003 4.003 0 0 1-4.007-4.008v-8.984A4 4 0 0 1 29.383 63z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foldback(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m26.893 131.912l4.46 4.69l2.446 1.474l5.38 6.5l5.79 9.75l5.078 11.62l1.93 5.485l7.14-16.037l7.304-23.997l6.195-22.475l7.541-25.745l6.6-18.934l16.86.197s2.362 8.456 5.281 15.743c2.92 7.287 3.8 9.827 6.805 15.421c3.005 5.595 2.535 7.704 6.853 10.01s8.505 3.917 13.7-.034c5.195-3.951 5.944-6.498 9.038-14.544c3.094-8.046 6.703-16.39 6.703-16.39l4.266-11.376l15.742.29l5.441 15.985l3.654 16.758l3.453 11.883l4.768 15.94l4.517 14.87l3.35 10.835l8.068 19.517l8.43-16.59l12.177-16.284l6.196-4.408v20.692l-5.068 8.06l-6.562 14.624l-4.41 13l-1.17 3.638l-16.463-.325s-3.706-7.796-7.227-16.116c-3.52-8.32-2.7-6.092-6.662-16.703c-3.962-10.611-4.306-10.983-7.521-21.282c-3.216-10.299-1.807-6.957-5.586-20.7c-3.78-13.744-7.664-27.476-7.664-27.476s-4.545 8.716-7.894 14.32c-3.349 5.604-3.732 6.807-7.28 11.07c-3.549 4.264-3.972 5.41-8.61 8.023c-4.639 2.612-5.317 3.693-11.734 3.435s-5.726.902-12.608-4.415c-6.883-5.317-9.84-12.928-9.84-12.928l-9.947-19.675l-6.46 18.273l-4.717 16.656l-4.047 15.298l-6.152 19.5l-6.043 16.71l-7.825 16.01l-17.727-.172l-4.82-16.194l-6.766-15.777l-4.363-6.854z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M138.033 60.858c0-14.183 8.718-16.43 13.894-10.454c5.176 5.977 47.925 60.873 52.46 68.201c4.535 7.329 5.434 13.082.086 20.28c-5.349 7.198-45.27 61.462-52.015 67.2c-6.745 5.74-14.543 3.358-14.543-10.567c0-13.926.118-120.476.118-134.66m51.265 71.792c1.34-1.758 1.357-4.612.03-6.388L156.4 82.205c-1.323-1.77-2.396-1.405-2.396.792L154 175.003c0 2.207 1.088 2.568 2.424.814zM47.033 60.858c0-14.183 8.718-16.43 13.894-10.454c5.176 5.977 47.925 60.873 52.46 68.201c4.535 7.329 5.434 13.082.086 20.28c-5.349 7.198-45.27 61.462-52.015 67.2c-6.745 5.74-14.543 3.358-14.543-10.567c0-13.926.118-120.476.118-134.66m51.265 71.792c1.34-1.758 1.357-4.612.03-6.388L65.4 82.205c-1.323-1.77-2.396-1.405-2.396.792L63 175.003c0 2.207 1.088 2.568 2.424.814z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hexpand(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M201.002 111c4.417 0 7.998 3.579 7.998 8.01v16.98c0 4.424-3.588 8.01-7.998 8.01H55.998C51.58 144 48 140.421 48 135.99v-16.98c0-4.424 3.588-8.01 7.998-8.01zm0 65c4.417 0 7.998 3.579 7.998 8.01v16.98c0 4.424-3.588 8.01-7.998 8.01H55.998C51.58 209 48 205.421 48 200.99v-16.98c0-4.424 3.588-8.01 7.998-8.01zm0-129c4.417 0 7.998 3.579 7.998 8.01v16.98c0 4.424-3.588 8.01-7.998 8.01H55.998C51.58 80 48 76.421 48 71.99V55.01c0-4.424 3.588-8.01 7.998-8.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hardclip(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m25.856 160.231l-.105 15.5l21.52.091s10.258.899 17.47-1.033c7.21-1.932 5.846.283 11.266-6.664s2.59-5.662 5.685-15.063c3.095-9.402 5.482-18.859 5.482-18.859l5.383-19.944l6.906-17.103l4.976-5.127l11.477-.617l16.25.1l18.06-.211l6.094.464l5.18 7.82l4.468 14.117l4.062 14.727l4.04 14.208l4.367 14.726s2.14 7.77 6.398 11.62c4.257 3.851 13.406 6.293 13.406 6.293l20.313.3l13.651-.105l.502-15.884l-16.709.405l-17.022-.72l-2.563-.717l-5.742-17.059l-6.713-23.695l-5.777-19.032s-1.753-7.91-6.973-13.517c-5.22-5.607-8.141-8.08-15.059-10.146c-6.917-2.066-10.042-.902-21.245-.803c-11.202.099-17.124.015-22.405.19c-5.281.174-10.457.896-10.457.896l-9.33 3.96l-8.1 11.07l-5.023 12.188l-5.23 18.891l-3.999 14.727l-4.511 13.608l-3.282 9.445l-17.84.793l-18.87.16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hardclipcurve(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M232 64.5h-54l-111.5 112H26V193h50L187.5 81H232z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M46.735 116.622c0-47.684 41.426-69.596 81.994-69.696c40.567-.1 80.92 24.207 80.92 69.292c0 45.086.004 70.408.004 70.408l-4.204 3.418l-9.406.045l-2.946-3.716s-.485-33.33-.485-70.36c0-37.031-27.957-52.448-64.552-52.481c-36.595-.033-64.821 15.18-64.821 52.613v71.228l-3.966 3.393H51.49l-4.755-3.393z"/><path d="M147.773 144.92c5.397 0 3.968-.253 15.42-.253c11.454 0 20.303 8.673 20.303 19.444v23.847c0 11.872-9.148 20.76-20.36 20.76c-11.214 0-9.847-.135-15.433-.068s-10.349-4.51-10.349-10.51s.176-37.786.176-43.587c0-5.8 4.845-9.632 10.243-9.632zm5.214 15.68l.107 32.639s7.836 0 10.124-.043c2.288-.042 5.123-2.588 5.123-5.242c0-2.655.532-18.77.532-21.468c0-2.698-2.388-5.886-5.484-5.992c-3.097-.107-10.402.106-10.402.106m-59.403-15.732c11.228-.048 10.137 0 15.342 0s10.303 4.385 10.385 9.819c.082 5.433 0 38.77-.082 44.005c-.081 5.235-4.267 9.885-10.303 9.885c-6.035 0-4.906-.196-15.551-.098s-20.162-9.532-20.162-20.48s.661-13.256.33-23.772c-.33-10.516 8.813-19.311 20.041-19.36zm-5.253 21.12c0 2.807.093 19.047.093 22.116c0 3.068 2.26 4.918 5.38 4.918s10.077-.045 10.077-.045v-32.65s-7.217-.124-9.964-.124c-2.748 0-5.586 2.98-5.586 5.786z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M48 51.995A4 4 0 0 1 51.996 48H65.52a3.987 3.987 0 0 1 3.984 3.998l-.23 74.785a1 1 0 0 0 1.003 1.01h13.729A3.997 3.997 0 0 1 88 131.796v72.198A4.006 4.006 0 0 1 84 208H52c-2.21 0-4-1.797-4-3.995zm160-.672c0-1.835-1.51-3.323-3.379-3.323h-13.914c-1.866 0-3.379 1.488-3.379 3.327l.266 75.636a.835.835 0 0 1-.845.83h-15.877c-1.862 0-3.372 1.495-3.372 3.318v73.57c0 1.833 1.514 3.319 3.383 3.319h33.734c1.868 0 3.383-1.482 3.383-3.323zm-91.735 1.101a4.083 4.083 0 0 1 4.005-4.073l15.441-.275a3.92 3.92 0 0 1 4.005 3.932v75.438c0 .553.452.987 1.005.968l7.29-.247a3.83 3.83 0 0 1 3.989 3.86v72.035a4 4 0 0 1-4 3.996h-39.646a4.017 4.017 0 0 1-4.019-3.996L104 132.027a3.974 3.974 0 0 1 3.987-3.995h7.273c.555 0 1.005-.448 1.005-1.007z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M80.573 123L80.1 98.915C79.513 69.071 91.715 45.5 120 45.5l13.944.014c30.548.588 41.343 23.594 41.93 53.447l.473 24.039h7.646c4.416 0 8.007 3.586 8.007 8.01v84.98c0 4.43-3.585 8.01-8.007 8.01H72.007c-4.416 0-8.007-3.586-8.007-8.01v-84.98c0-4.43 3.585-8.01 8.007-8.01zm16.39 0h62.657c.033-5.473.104-12.555-.05-20.33l-.078-4.024c-.414-21.015-2.663-34.678-25.194-35.11l-12.405-.358c-28.124-.56-25.578 26.864-25.397 36.052l.08 4.024zM176 143.316c0-2.384-1.547-4.316-3.46-4.316H83.46c-1.911 0-3.46 1.936-3.46 4.316v60.368c0 2.384 1.547 4.316 3.46 4.316h89.08c1.911 0 3.46-1.936 3.46-4.316z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoAax(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m71.922 208.45l14.08-37.047l9.293-.155l13.755 35.308l14.05-35.028l9.303.061l13.118 34.854l11.959-14.645l-11.002-20.113l10.135.035l6.978 12.248l9.23-12.257l9.074.122l-10.784 18.213l10.784 18.83h-9.44l-8.942-12.859l-7.695 12.425l-18.195-.054l-3.61-8.886l-12.541-.298l-3.955 10.43l-17.202-.758l-2.976-9.81l-13.327.25l-2.686 8.846zm13.24-14.775l10.664-.034l-5.284-13.918zm36.876-.299h10.377L128 179.416zM128 165.854c-32.585 0-59-26.606-59-59.427C69 73.607 95.415 47 128 47s59 26.606 59 59.427c0 32.82-26.415 59.427-59 59.427m-35.75-23.047l7.102 6.61l9.093 4.196l8.008 2.893l11.15 1.277l10.7-1.277l8.866-2.971l9.633-5.078l6.153-5.505l-2.474-1.74l-5.4-4.86l-5.534-6.032l-3.327-5.515l-3.89-5.904l-4.027-6.067l-4.183-5.136l-6.517-1.925l-7.057 1.925l-4.093 5.136l-4.87 8.407l-5.57 8.556l-6.125 7.368l-7.637 5.642zm-12.004-22.971c-.007 2.21.69 5.658 1.548 7.68l-.9-2.122c.862 2.031 3.27 4.196 4.062 4.436l-2.501-.759c1.381.42 3.823-.446 5.46-1.94l.106-.096c1.634-1.49 3.852-4.244 4.963-6.165l2.825-4.885a183.839 183.839 0 0 1 4.189-6.824l2.084-3.203c1.207-1.855 3.382-4.687 4.867-6.338l.918-1.019c1.481-1.645 4.225-3.89 6.127-5.013l1.52-.897c1.903-1.123 5.18-2.505 7.307-3.083l.091-.025c2.133-.58 5.647-.869 7.834-.647l-1.281-.13c2.194.222 5.677.954 7.782 1.635l1.01.327c2.103.68 5.192 2.371 6.904 3.78l.187.154c1.709 1.407 4.225 3.94 5.615 5.654l.168.206c1.392 1.715 3.495 4.608 4.698 6.46l.242.374c1.203 1.852 2.97 4.955 3.953 6.938l.064.128a160.26 160.26 0 0 0 3.765 7.06l.145.253c1.099 1.919 3.165 4.813 4.632 6.482l-1.22-1.39c1.459 1.662 4.395 3.347 6.56 3.766l-3.918-.758c2.164.419 4.555-.915 5.34-2.974l-.599 1.57c.785-2.06 1.67-5.507 1.978-7.701l.405-2.887c.307-2.192.342-5.752.08-7.94l-.13-1.08c-.262-2.193-1.053-5.663-1.764-7.747l-.806-2.361c-.712-2.086-2.354-5.216-3.672-6.997l1.117 1.51a400.037 400.037 0 0 1-4.681-6.485l1.121 1.594c-1.268-1.803-3.663-4.415-5.358-5.842l1.267 1.066a580.391 580.391 0 0 1-6.075-5.2l.82.714c-1.663-1.45-4.618-3.403-6.617-4.373l.284.138c-1.99-.966-5.286-2.39-7.345-3.174l1.326.505c-2.066-.788-5.493-1.791-7.655-2.243l1.935.404c-2.162-.45-5.698-.844-7.914-.879l-8.099-.127c-2.21-.035-5.753.321-7.915.794l-.436.096c-2.162.473-5.55 1.601-7.558 2.516l1.935-.882c-2.011.917-5.148 2.637-7.002 3.841l-3.946 2.561c-1.856 1.204-4.548 3.527-6.01 5.185l-.366.413a151.865 151.865 0 0 0-5.1 6.165l-2.55 3.29c-1.354 1.746-3.095 4.825-3.895 6.892l1.105-2.854c-.798 2.06-1.84 5.484-2.324 7.631l.254-1.128c-.486 2.154-.891 5.69-.905 7.905l-.025 3.772c-.014 2.212-.031 5.791-.038 8.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoAbletonlink(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M77.004 105A2.006 2.006 0 0 0 75 106.994v4.661c0 .826.664 1.495 1.497 1.495h44.006c.827 0 1.497-.667 1.497-1.495v-5.16c0-.826-.664-1.495-1.497-1.495zm-.507 12c-.827 0-1.497.667-1.497 1.495v5.16c0 .826.664 1.495 1.497 1.495h44.006c.827 0 1.497-.667 1.497-1.495v-5.16c0-.826-.664-1.495-1.497-1.495zm0 13c-.827 0-1.497.667-1.497 1.495v5.16c0 .826.664 1.495 1.497 1.495h44.006c.827 0 1.497-.667 1.497-1.495v-5.16c0-.826-.664-1.495-1.497-1.495zm0 13c-.827 0-1.497.667-1.497 1.495v5.16c0 .826.664 1.495 1.497 1.495h44.006c.827 0 1.497-.667 1.497-1.495v-5.16c0-.826-.664-1.495-1.497-1.495zM63 150.503c0 .827.667 1.497 1.495 1.497h5.16c.826 0 1.495-.664 1.495-1.497v-44.006c0-.827-.667-1.497-1.495-1.497h-5.16c-.826 0-1.495.664-1.495 1.497zm-27 0c0 .827.667 1.497 1.495 1.497h5.16c.826 0 1.495-.664 1.495-1.497v-44.006c0-.827-.667-1.497-1.495-1.497h-5.16c-.826 0-1.495.664-1.495 1.497zm-5.85 0c0 .827-.667 1.497-1.495 1.497h-5.16A1.492 1.492 0 0 1 22 150.503v-44.006c0-.827.667-1.497 1.495-1.497h5.16c.826 0 1.495.664 1.495 1.497zm18.85 0c0 .827.667 1.497 1.495 1.497h5.16c.826 0 1.495-.664 1.495-1.497v-44.006c0-.827-.667-1.497-1.495-1.497h-5.16c-.826 0-1.495.664-1.495 1.497zm92.62-45.154c0 1.015-.205 45.676-.205 45.676l20.818-.014l.303-7.176l-14.247-.416l.04-37.945s-6.71-1.141-6.71-.125zm29.114 1.851c-1.492.056-2.948 1.127-2.88 2.23c.067 1.105 1.691 2.19 3.148 2.133c1.457-.055 2.796-1.286 2.73-2.357c-.065-1.071-1.505-2.064-2.998-2.007zm-2.557 14.816l-.112 28.517l6.094-.006l.495-28.665zm11.063.208l-.425 28.885h8.024l-.411-20.065l3.351-4.204l4.774-.205l3.758 3.86l.202 19.702l8.023-.004l-1.53-24.054l-6.05-4.922l-8.201-.513l-5.434 3.453l-.008-4.101zm31.406-21.795l.146 50.145l6.297.305s1.209-15.156.604-14.344c-.604.813 9.243 13.711 9.243 13.711l9.445-.204l-11.992-16.081l9.952-12.906l-7.711-.308l-8.309 10.245l.088-30.412z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoAu(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M79.912 85.68c.783.78.947 2.176.366 3.116l-3.784 6.12a36.589 36.589 0 0 0-1.874 3.536l-3.136 6.995c-.454 1.012-.98 2.722-1.175 3.81l-1.55 8.697a30.277 30.277 0 0 0-.389 3.974l-.155 8.611c-.02 1.103.169 2.863.424 3.948l1.983 8.408c.254 1.077.846 2.754 1.327 3.755l3.084 6.417a37.89 37.89 0 0 0 1.955 3.484l2.728 4.21c.602.927.468 2.322-.296 3.112l-2.834 2.93a1.875 1.875 0 0 1-1.383.585a1.88 1.88 0 0 1-1.371-.615l-5.782-6.179c-.756-.808-1.68-2.301-2.06-3.331l-4.543-12.277c-.382-1.031-.84-2.754-1.023-3.84l-1.506-8.952c-.183-1.09-.312-2.859-.288-3.972l.22-10.16c.023-1.104.203-2.889.397-3.96l1.802-9.972c.196-1.083.71-2.79 1.142-3.802l2.585-6.042a45.244 45.244 0 0 1 1.769-3.59l3.526-6.288a57.938 57.938 0 0 1 2.097-3.402l.11-.162c.617-.913 1.752-1.02 2.535-.24z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M57.286 68.45c.658-.885 1.87-1.026 2.72-.303l4.153 3.526c.844.717 1.084 2.068.524 3.036l-3.54 6.12c-.555.96-1.447 2.505-2.002 3.471l-2.281 3.97a30.77 30.77 0 0 0-1.71 3.6l-2.1 5.485a60.02 60.02 0 0 0-1.275 3.778l-1.787 6.095c-.31 1.057-.64 2.8-.737 3.902l-.612 6.869a46.374 46.374 0 0 0-.148 3.98l.14 9.312c.016 1.1.16 2.881.322 3.97l1.29 8.706c.163 1.093.612 2.823.999 3.85l3.007 7.984a46.092 46.092 0 0 0 1.605 3.665l3.236 6.447a38.202 38.202 0 0 0 2.018 3.455l3.222 4.79c.617.917.458 2.263-.362 3.014l-3.783 3.461c-.817.748-1.974.607-2.583-.312l-8.193-12.378a32.003 32.003 0 0 1-1.947-3.483l-3.912-8.412c-.466-1.003-1.03-2.7-1.255-3.782l-2.133-10.203a61.031 61.031 0 0 1-.664-3.951l-1.223-9.584c-.14-1.096-.206-2.884-.148-3.977l.515-9.715c.058-1.1.276-2.865.489-3.951l2.035-10.385c.212-1.081.69-2.793 1.073-3.834l2.793-7.608a31.658 31.658 0 0 1 1.656-3.628l4.864-8.816a35.437 35.437 0 0 1 2.159-3.355z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M42.098 56.514c.635-.903 1.797-1.004 2.58-.238l3.885 3.795c.79.772.942 2.141.332 3.07l-4.027 6.128a65.063 65.063 0 0 0-2.066 3.413l-1.98 3.562a79.34 79.34 0 0 0-1.83 3.543l-3.227 6.74a37.21 37.21 0 0 0-1.488 3.712l-1.363 4.153a50 50 0 0 0-1.063 3.863l-1.62 7.233a28.698 28.698 0 0 0-.54 3.955l-.525 10.074a238.821 238.821 0 0 1-.249 4l-.377 5.214c-.08 1.104-.05 2.887.068 3.995l.779 7.32c.117 1.104.417 2.879.665 3.947l2.534 10.9c.25 1.076.814 2.78 1.247 3.778l4.035 9.294a50.19 50.19 0 0 0 1.76 3.576l3.96 7.163c.535.966 1.485 2.47 2.133 3.373l3.497 4.875c.644.897.47 2.184-.391 2.876l-3.387 2.723a2.01 2.01 0 0 1-2.218.195a2.015 2.015 0 0 1-.599-.503l-6.017-7.449c-.695-.861-1.682-2.358-2.198-3.335l-4.875-9.219a86.801 86.801 0 0 1-1.77-3.587l-2.76-6.007c-.458-1-1.073-2.664-1.377-3.733l-2.356-8.283a34.366 34.366 0 0 1-.82-3.91l-1.552-11.235a34.264 34.264 0 0 1-.268-3.994l.033-10.192c.003-1.107.084-2.889.182-3.997l1.217-13.835c.097-1.1.45-2.844.79-3.899l2.674-8.303c.412-1.265.845-2.523 1.298-3.774l2.952-8.056c.379-1.032 1.148-2.628 1.723-3.571l5.504-9.03a73.369 73.369 0 0 1 2.188-3.34zM177.006 85.68c-.784.78-.947 2.176-.366 3.116l3.783 6.12c.581.941 1.42 2.523 1.874 3.536l3.137 6.995c.454 1.012.98 2.722 1.174 3.81l1.55 8.697c.196 1.092.37 2.863.39 3.974l.154 8.611c.02 1.103-.168 2.863-.424 3.948l-1.983 8.408c-.254 1.077-.846 2.754-1.326 3.755l-3.084 6.417a37.758 37.758 0 0 1-1.956 3.484l-2.728 4.21c-.601.927-.468 2.322.296 3.112l2.834 2.93a1.875 1.875 0 0 0 1.384.585a1.881 1.881 0 0 0 1.37-.615l5.783-6.179c.756-.808 1.679-2.301 2.06-3.331l4.542-12.277c.382-1.031.84-2.754 1.023-3.84l1.506-8.952c.183-1.09.312-2.859.288-3.972l-.219-10.16a30.758 30.758 0 0 0-.398-3.96l-1.802-9.972c-.196-1.083-.71-2.79-1.142-3.802l-2.585-6.042a45.21 45.21 0 0 0-1.769-3.59l-3.526-6.288a58.045 58.045 0 0 0-2.097-3.402l-.11-.162c-.616-.913-1.752-1.02-2.535-.24z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M198.805 68.45c-.658-.885-1.87-1.026-2.72-.303l-4.152 3.526c-.844.717-1.084 2.068-.524 3.036l3.54 6.12c.555.96 1.446 2.505 2.002 3.471l2.281 3.969c.55.958 1.316 2.568 1.71 3.6l2.1 5.485c.395 1.03.967 2.729 1.275 3.778l1.786 6.095c.31 1.057.64 2.8.738 3.902l.612 6.869c.106 1.324.155 2.652.147 3.98l-.138 9.312a36.375 36.375 0 0 1-.323 3.97l-1.291 8.706c-.162 1.093-.611 2.823-.998 3.85l-3.007 7.984a46.201 46.201 0 0 1-1.605 3.665L197 165.912a38.192 38.192 0 0 1-2.017 3.455l-3.222 4.79c-.617.917-.459 2.263.362 3.014l3.783 3.461c.817.748 1.974.607 2.583-.312l8.192-12.378a31.734 31.734 0 0 0 1.948-3.483l3.912-8.412c.466-1.003 1.03-2.7 1.255-3.782l2.133-10.203c.227-1.086.524-2.855.664-3.951l1.223-9.584c.14-1.096.206-2.884.148-3.977l-.515-9.715a34.08 34.08 0 0 0-.49-3.951l-2.035-10.385c-.273-1.3-.631-2.58-1.072-3.834l-2.793-7.608a31.572 31.572 0 0 0-1.656-3.628l-4.865-8.816a35.453 35.453 0 0 0-2.158-3.355l-3.576-4.808z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M211.329 56.276c.783-.766 1.945-.665 2.581.238h-.001l4.908 6.98a74.05 74.05 0 0 1 2.188 3.34l5.503 9.03c.575.943 1.345 2.539 1.723 3.571l2.952 8.056c.374 1.02.959 2.722 1.299 3.774l2.673 8.303c.34 1.055.693 2.799.79 3.899l1.218 13.835c.108 1.33.168 2.663.182 3.997l.032 10.192a34.082 34.082 0 0 1-.268 3.994l-1.551 11.235a34.432 34.432 0 0 1-.821 3.91l-2.355 8.283a31.207 31.207 0 0 1-1.378 3.733l-2.759 6.007a85.8 85.8 0 0 1-1.77 3.587l-4.875 9.219c-.517.977-1.503 2.474-2.199 3.335l-6.016 7.449c-.7.866-1.957.998-2.817.308l-3.387-2.723a2.132 2.132 0 0 1-.392-2.876l3.497-4.875a41.178 41.178 0 0 0 2.133-3.373l3.961-7.163a49.978 49.978 0 0 0 1.76-3.576l4.034-9.294c.434-.998.997-2.702 1.247-3.778l2.534-10.9c.288-1.304.511-2.621.666-3.947l.778-7.32c.118-1.108.148-2.891.068-3.995l-.377-5.214c-.08-1.115-.191-2.895-.248-4l-.525-10.074c-.057-1.1-.3-2.877-.542-3.955l-1.618-7.233a50.096 50.096 0 0 0-1.064-3.863l-1.362-4.153a36.991 36.991 0 0 0-1.489-3.712l-3.226-6.74a78.713 78.713 0 0 0-1.83-3.543l-1.98-3.562a65.613 65.613 0 0 0-2.066-3.413l-4.028-6.128c-.61-.929-.457-2.298.333-3.07zM105.408 82.375c-.819.493-1.52 1.182-3.395 2.622l-.509.39c-3.975 3.049-6.281 4.817-10.554 10.992l-.346.5c-4.135 5.972-4.665 6.737-6.714 14.308c-1.513 5.588-1.435 7.135-1.271 10.425c.064 1.289.142 2.846.143 5.017c0 1.843-.03 3.181-.054 4.23c-.076 3.345-.085 3.763 1.192 8.282c1.384 4.899 1.855 5.551 3.843 8.303c.42.584.91 1.261 1.491 2.094c.326.466.626.901.909 1.309c2.6 3.753 3.668 5.297 8.528 9.461c.813.697 1.51 1.308 2.132 1.854c3.499 3.07 4.605 4.04 10.422 6.218c6.309 2.363 7.829 2.413 14.378 2.629l1.802.061l1.278.045c6.873.247 7.302.263 14.39-1.619c6.497-1.726 7.183-2.163 11.279-4.777c.634-.405 1.349-.861 2.181-1.383c3.29-2.063 4.689-2.937 5.889-4.029c1.063-.967 1.97-2.105 3.897-4.394l.394-.467c3.736-4.432 4.156-4.93 6.384-10.783c2.276-5.978 2.453-7.541 3.302-15.034l.032-.281c.481-4.242.721-5.764.653-7.272c-.056-1.242-.322-2.475-.834-5.213c-1.135-6.059-1.386-7.177-4.658-13.361c-2.885-5.453-3.405-5.989-6.708-9.393c-.443-.456-.936-.964-1.491-1.542c-2.479-2.583-3.531-3.73-4.747-4.665c-1.087-.836-2.305-1.502-4.791-2.875c-5.265-2.907-6.519-3.551-13.475-5.167c-4.437-1.032-5.758-.97-7.851-.87a47.113 47.113 0 0 1-5.019.01c-6.387-.303-7.96.135-14.81 2.043l-.756.21c-4.225 1.176-5.478 1.486-6.536 2.122M80.836 223.178c-.355 1.053.255 1.9 1.36 1.9h4.531c1.103 0 2.323-.836 2.721-1.866l3.704-9.586c.399-1.033 1.612-1.867 2.712-1.867h11.746c1.1 0 2.318.834 2.72 1.86l3.824 9.76c.402 1.027 1.623 1.845 2.726 1.825l5.106-.09c1.108-.02 1.698-.88 1.328-1.92l-16.418-46.253c-.368-1.037-1.567-1.885-2.676-1.885h-5.13c-1.113 0-2.294.851-2.648 1.9zm15.286-20.571c-.246 1.059.44 1.94 1.54 1.94h9.594l-4.561-18.797h-2.65zm40.353-26.893l8.486-.137l.716 35.126s.778 4.31 2.726 6.32c1.948 2.01 5.173 1.827 5.173 1.827s4.89.83 7.523-.299c2.632-1.128 3.205-1.82 4.531-4.867c1.326-3.046.836-7.462.836-7.462l-.266-30.162l10.322-.124l-.439 32.3s-.609 5.629-3.762 10.018c-3.154 4.39-2.916 4.47-8.14 6.55c-5.226 2.079-6.049 2.137-11.194 1.144c-5.145-.994-7.546-2.597-10.91-5.22c-3.365-2.621-4.31-5.86-4.31-5.86l-1.233-4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoAudacity(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M61.264 122.309c0-3.998.603-29.443.603-34.157c0-4.714 3.124-12.91 13-22.55c9.877-9.641 19.451-11.343 30.057-14.97c10.606-3.625 14.1-2.743 23.508-2.743c9.407 0 9.79-.342 23.76 3.533c13.97 3.875 23.794 10.105 31.89 17.047s7.9 13.254 9.522 19.46c1.62 6.208 1.62 21.313 1.62 34.157c12.556 11.611 14.366 17.133 14.366 38.821c0 21.689-11.056 37.037-18.988 42.863c-7.933 5.825-4.441 4.586-9.018 5.078c-4.577.491-9.639-3.075-9.639-3.075l-.183-80.164l7.312-4.511s.998-24.812.998-29.657c0-4.844-4.225-12.684-6.765-16.199c-2.54-3.514-10.434-6.773-19.471-9.64c-9.037-2.868-13.725-3.533-25.404-3.487c-11.68.047-15.175.015-27.09 3.487c-11.915 3.471-10.543 3.688-16.764 9.64c-6.22 5.952-8.713 11.33-8.56 16.506c.152 5.176.804 30.703.804 30.703l9.14 5.281l-1.384 75.297s-3.958 4.19-8.408 4.86c-4.45.67-7.615-1.718-14.906-6.995c-7.291-5.278-14.742-17.096-14.742-38.16c0-21.063.044-29.037 14.742-40.425"/><path d="m92.731 155.314l7.613-23.496l15.719 45.975l12.555-62.277l10.81 62.45l15.869-45.726l9.202 22.975l-.778 29.69l-7.743-28.322l-15.685 45.376l-11.98-45.711l-11.762 45.583l-15.436-45.686l-8.447 28.612z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoAudiobus(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M36 150.726c0 .703.982 1.274 2.201 1.274h7.598c1.216 0 2.201-.565 2.201-1.274v-37.452c0-.703-.982-1.274-2.201-1.274h-7.598c-1.216 0-2.201.565-2.201 1.274m20 55.273c0 1.355.982 2.453 2.201 2.453h7.598c1.216 0 2.201-1.088 2.201-2.453V96.453C68 95.098 67.018 94 65.799 94h-7.598C56.985 94 56 95.088 56 96.453m19 56.273c0 .703.982 1.274 2.201 1.274h7.598c1.216 0 2.201-.565 2.201-1.274v-37.452c0-.703-.982-1.274-2.201-1.274h-7.598c-1.216 0-2.201.565-2.201 1.274m19 30.77c0 .528.982.956 2.201.956h7.598c1.216 0 2.201-.424 2.201-.956v-28.088c0-.528-.982-.956-2.201-.956h-7.598c-1.216 0-2.201.424-2.201.956m20 54.304c0 1.513.982 2.74 2.201 2.74h7.598c1.216 0 2.201-1.216 2.201-2.74V91.74c0-1.513-.982-2.74-2.201-2.74h-7.598c-1.216 0-2.201 1.216-2.201 2.74m19 106.8c0 2.463.982 4.46 2.201 4.46h7.598c1.216 0 2.201-1.98 2.201-4.46V67.46c0-2.463-.982-4.46-2.201-4.46h-7.598c-1.216 0-2.201 1.98-2.201 4.46m19 116.036c0 1.935.982 3.504 2.201 3.504h7.598c1.216 0 2.201-1.555 2.201-3.504V80.504c0-1.935-.982-3.504-2.201-3.504h-7.598c-1.216 0-2.201 1.555-2.201 3.504m19 89.947c0 1.408.982 2.549 2.201 2.549h7.598c1.216 0 2.201-1.13 2.201-2.549V95.55c0-1.408-.982-2.549-2.201-2.549h-7.598c-1.216 0-2.201 1.13-2.201 2.549m19 62.698c0 .967.982 1.752 2.201 1.752h7.598c1.216 0 2.201-.777 2.201-1.752v-51.496c0-.967-.982-1.752-2.201-1.752h-7.598c-1.216 0-2.201.777-2.201 1.752m19 73.062c0 1.76.982 3.186 2.201 3.186h7.598c1.216 0 2.201-1.414 2.201-3.186V86.186c0-1.76-.982-3.186-2.201-3.186h-7.598c-1.216 0-2.201 1.414-2.201 3.186"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoCubase(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M58.446 127.855L140.401 46.3l32.106 26.083s-11.13.225-20.794 3.706c-9.663 3.48-14.486 8.596-19.58 14.04c-11.146 11.915-14.325 23.365-14.243 37.955c.08 14.59 4.462 22.448 12.94 33.133c2.87 3.619 5.696 7.099 10.508 9.728c1.64.897 8.563 5.947 13.122 7.36c4.559 1.413 16.789 1.949 16.789 1.949l-29.911 31.06z"/><circle cx="173" cy="126" r="22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoFl(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M90.643 97.44s-4.586 0-5.309-1.005c-.722-1.005-1.558-2.991-1.5-5.913c.058-2.921.557-3.22 1.733-5.823c1.177-2.604 1.592-3.593 3.717-6.239c2.124-2.646 2.953-3.544 6.23-6.153c3.28-2.608 3.441-3.037 7.174-4.512c3.734-1.475 5.6-2.2 11.255-2.768c5.655-.566 4.472-.401 8.768.761s5.042 1.18 9.55 4.196c4.508 3.015 8.397 7.81 8.397 7.81l4.82-9.237l7.266-9.53l11.205-11.582l4.332-.101l1.392 4.69l-5.677 5.34l-8.826 11.27l-6.59 12.584s5.965-.458 11.805.587s5.267.322 9.828 3.284c4.56 2.962 6.525 3.368 10.667 10.028c4.141 6.66 4.082 7.188 4.927 15.051c.845 7.863.594 8.612-1.781 14.235c-2.376 5.623-7.35 7.375-7.35 7.375l-6.042-3.86s-.155 3.837-1.171 7.152c-1.017 3.314-2.294 7.256-5.262 13.832c-2.968 6.576-3.867 8.752-8.447 16.545c-4.58 7.794-4.533 7.476-9.561 14.097c-5.028 6.622-5.52 7.248-11.461 13.59c-5.942 6.341-5.408 7.14-11.71 11.142c-6.303 4.001-15 5.816-15 5.816s-7.236-.109-13.338-3.998c-6.102-3.888-8.627-10.001-8.627-10.001s-2.973-7.293-4.894-15.16c-1.92-7.865-1.733-7.071-2.316-14.37c-.584-7.298-.45-8.579-.2-17.11c.251-8.532.012-8.558 1.192-16.743c1.18-8.185 1.214-9.874 3.917-18.698c2.702-8.825 6.887-16.582 6.887-16.582m6.757-1.13s-4.981 6.523-8.491 16.386c-3.51 9.862-4.865 21.143-4.865 21.143s-1.057 7.613-1.178 16.826c-.122 9.214-.433 9.227.714 18.455c1.147 9.228 2.148 12.589 3.83 18.112c1.683 5.523 4.79 10.185 4.79 10.185s3.491 5.343 8.21 7.325c4.72 1.981 6.69 1.61 10.285.44c3.594-1.17 10.828-5.365 16.947-10.386c6.12-5.02 11.044-12.58 11.044-12.58s7.887-10.361 12.657-18.117c4.77-7.756 5.593-9.93 8.637-16.503c3.043-6.573 4.245-8.834 5.366-13.738c1.12-4.904 1.12-10.98 1.12-10.98l-8.993-10.091l-8.248 9.205s-6.31 4.725-11.26 5.854c-4.95 1.129-5.917.617-10.26-.946s-5.084-1.915-8.154-5.68c-3.071-3.765-4.529-9.87-4.529-9.87s-.717-10.799-.777-13.16c-.06-2.36.322-4.717.322-4.717z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoJuce(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M128 226c-54.124 0-98-43.876-98-98s43.876-98 98-98s98 43.876 98 98s-43.876 98-98 98m0-7c50.258 0 91-40.742 91-91s-40.742-91-91-91s-91 40.742-91 91s40.742 91 91 91"/><path d="M52.163 94.493C56.357 84.463 62.375 78.633 64 76.984c1.625-1.648 3.932-.844 5.477 0c1.545.845 40.273 39.636 42.002 41.584c1.728 1.948 3.56 4.854-1.401 4.854c-4.962 0-53.708-.426-57.915-.716c-4.208-.291-5.45-2.745-5.45-4.841s1.255-13.341 5.45-23.372M99.015 51.34c10.059-4.127 17.23-3.856 19.545-3.872c2.315-.016 3.377 2.183 3.873 3.872c.495 1.69.855 56.835.7 59.434c-.156 2.6-.113 5.483-3.621 1.974c-3.509-3.508-37.677-38.278-40.446-41.458c-2.77-3.181-1.913-5.795-.43-7.277c1.482-1.482 10.32-8.545 20.38-12.673zm61.492.823c10.03 4.194 15.86 10.212 17.509 11.837c1.648 1.625.844 3.932 0 5.477c-.845 1.545-39.636 40.273-41.584 42.002c-1.948 1.728-4.854 3.56-4.854-1.401c0-4.962.426-53.708.716-57.915c.291-4.208 2.745-5.45 4.841-5.45s13.341 1.255 23.372 5.45m43.153 46.852c4.127 10.059 3.856 17.23 3.872 19.545c.016 2.315-2.183 3.377-3.872 3.873c-1.69.495-56.835.855-59.434.7c-2.6-.156-5.483-.113-1.974-3.621c3.508-3.509 38.278-37.677 41.458-40.446c3.181-2.77 5.795-1.913 7.277-.43c1.482 1.482 8.545 10.32 12.673 20.38zm-.823 62.492c-4.194 10.03-10.212 15.86-11.837 17.509c-1.625 1.648-3.932.844-5.477 0c-1.545-.845-40.273-39.636-42.002-41.584c-1.728-1.948-3.56-4.854 1.401-4.854c4.962 0 53.708.426 57.915.716c4.208.291 5.45 2.745 5.45 4.841s-1.255 13.341-5.45 23.372m-46.852 43.153c-10.059 4.127-17.23 3.856-19.545 3.872c-2.315.016-3.377-2.183-3.873-3.872c-.495-1.69-.855-56.835-.7-59.434c.156-2.6.113-5.483 3.621-1.974c3.509 3.508 37.677 38.278 40.446 41.458c2.77 3.181 1.913 5.795.43 7.277c-1.482 1.482-10.32 8.545-20.38 12.673zm-61.492-.823c-10.03-4.194-15.86-10.212-17.509-11.837c-1.648-1.625-.844-3.932 0-5.477c.845-1.545 39.636-40.273 41.584-42.002c1.948-1.728 4.854-3.56 4.854 1.401c0 4.962-.426 53.708-.716 57.915c-.291 4.208-2.745 5.45-4.841 5.45s-13.341-1.255-23.372-5.45M51.34 156.985c-4.127-10.059-3.856-17.23-3.872-19.545c-.016-2.315 2.183-3.377 3.872-3.873c1.69-.495 56.835-.855 59.434-.7c2.6.156 5.483.113 1.974 3.621c-3.508 3.509-38.278 37.677-41.458 40.446c-3.181 2.77-5.795 1.913-7.277.43c-1.482-1.482-8.545-10.32-12.673-20.38z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoLadspa(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.102 100.914h12.386v30.215H39.82V141H8.102zm62.015 33.469h-14.11L54.068 141h-12.66l15.066-40.086H69.98L85.047 141H72.086zm-2.597-8.668l-4.43-14.41l-4.403 14.41zm21.656-24.8h18.402c3.628 0 6.558.491 8.791 1.476c2.233.984 4.079 2.397 5.537 4.238c1.459 1.841 2.516 3.983 3.172 6.426c.656 2.442.984 5.03.984 7.765c0 4.284-.487 7.606-1.462 9.967c-.976 2.36-2.33 4.339-4.06 5.934c-1.733 1.595-3.592 2.657-5.58 3.185c-2.715.73-5.176 1.094-7.382 1.094H89.176v-40.086zm12.386 9.077v21.903h3.036c2.588 0 4.43-.288 5.523-.862c1.094-.574 1.95-1.577 2.57-3.008c.62-1.43.93-3.75.93-6.959c0-4.247-.693-7.155-2.078-8.722c-1.385-1.568-3.682-2.352-6.89-2.352h-3.09zm28.848 17.746l11.785-.738c.256 1.914.775 3.372 1.559 4.375c1.276 1.622 3.099 2.434 5.469 2.434c1.768 0 3.13-.415 4.088-1.245c.957-.829 1.435-1.79 1.435-2.884c0-1.04-.456-1.97-1.367-2.79c-.912-.82-3.026-1.594-6.344-2.324c-5.432-1.221-9.306-2.843-11.62-4.867c-2.334-2.023-3.5-4.603-3.5-7.738c0-2.06.596-4.006 1.79-5.838c1.194-1.832 2.99-3.272 5.387-4.32c2.397-1.048 5.683-1.573 9.857-1.573c5.123 0 9.028.953 11.717 2.858c2.689 1.905 4.288 4.935 4.799 9.092l-11.676.683c-.31-1.804-.962-3.117-1.955-3.937c-.994-.82-2.365-1.23-4.115-1.23c-1.44 0-2.525.305-3.254.915c-.73.611-1.094 1.354-1.094 2.229c0 .638.3 1.212.902 1.722c.584.53 1.97 1.021 4.157 1.477c5.414 1.167 9.292 2.347 11.634 3.541c2.343 1.194 4.047 2.675 5.114 4.443c1.066 1.769 1.6 3.747 1.6 5.934c0 2.57-.712 4.94-2.133 7.11c-1.422 2.169-3.41 3.814-5.961 4.935c-2.553 1.121-5.77 1.682-9.653 1.682c-6.817 0-11.539-1.313-14.164-3.938s-4.11-5.96-4.457-10.008m42.547-26.824h20.59c4.484 0 7.843 1.066 10.076 3.2c2.233 2.132 3.35 5.167 3.35 9.105c0 4.047-1.217 7.21-3.65 9.488c-2.434 2.279-6.148 3.418-11.143 3.418h-6.782V141h-12.441zm12.441 17.09h3.036c2.388 0 4.065-.415 5.03-1.244c.967-.83 1.45-1.892 1.45-3.186c0-1.258-.42-2.324-1.258-3.199c-.838-.875-2.415-1.313-4.73-1.313h-3.528zm47.961 16.379H219.25L217.309 141h-12.66l15.066-40.086h13.508L248.289 141h-12.96zm-2.597-8.668l-4.43-14.41l-4.402 14.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoLive(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M136.107 91c-1.793 0-3.247 1.46-3.247 3.224v7.538a2.415 2.415 0 0 0 2.426 2.417h71.288a2.42 2.42 0 0 0 2.426-2.417v-8.345A2.415 2.415 0 0 0 206.574 91zm-.821 19.404a2.42 2.42 0 0 0-2.426 2.417v8.345a2.415 2.415 0 0 0 2.426 2.417h71.288a2.42 2.42 0 0 0 2.426-2.417v-8.345a2.415 2.415 0 0 0-2.426-2.417zm0 21.022a2.42 2.42 0 0 0-2.426 2.417v8.344a2.415 2.415 0 0 0 2.426 2.417h71.288a2.42 2.42 0 0 0 2.426-2.417v-8.344a2.415 2.415 0 0 0-2.426-2.417zm0 21.021a2.42 2.42 0 0 0-2.426 2.417v8.344a2.415 2.415 0 0 0 2.426 2.418h71.288a2.42 2.42 0 0 0 2.426-2.418v-8.344a2.415 2.415 0 0 0-2.426-2.417zm-21.866 12.132a2.42 2.42 0 0 0 2.422 2.421h8.36a2.415 2.415 0 0 0 2.421-2.421V93.42a2.42 2.42 0 0 0-2.422-2.42h-8.36a2.415 2.415 0 0 0-2.421 2.421zm-43.74 0A2.42 2.42 0 0 0 72.102 167h8.36a2.415 2.415 0 0 0 2.421-2.421V93.42A2.42 2.42 0 0 0 80.461 91h-8.36a2.415 2.415 0 0 0-2.421 2.421zm-9.477 0A2.42 2.42 0 0 1 57.781 167h-8.36A2.415 2.415 0 0 1 47 164.579V93.42A2.42 2.42 0 0 1 49.422 91h8.36a2.415 2.415 0 0 1 2.421 2.421zm30.537 0A2.42 2.42 0 0 0 93.162 167h8.36a2.415 2.415 0 0 0 2.421-2.421V93.42a2.42 2.42 0 0 0-2.422-2.42h-8.36a2.415 2.415 0 0 0-2.421 2.421z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoLvTwo(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M51.213 47.589c-2.204 0-3.574 1.737-3.056 3.892l4.49 18.688c.517 2.15 1.356 5.638 1.87 7.782l30.377 126.417c.516 2.148 2.723 3.89 4.926 3.89h81.545a3.996 3.996 0 0 0 3.991-4.005v-25.054a4.007 4.007 0 0 0-4.003-4.004h-54.85c-2.211 0-4.459-1.732-5.02-3.872L80.002 51.46c-.562-2.138-2.816-3.871-5.008-3.871z"/><path d="M94.031 48.388c-2.204 0-3.518 1.728-2.933 3.868l28.34 103.69c.585 2.137 2.841 3.869 5.054 3.869h7.292c2.207 0 4.385-1.75 4.862-3.9L159.66 52.287c.479-2.154-.923-3.9-3.13-3.9h-4.974c-2.207 0-4.339 1.756-4.762 3.92l-14.028 71.771c-.423 2.166-2.566 3.921-4.77 3.921h1.292c-2.212 0-4.46-1.737-5.02-3.867l-18.886-71.878c-.56-2.136-2.801-3.867-5.006-3.867H94.03zm52.796 111.398c-2.212 0-3.592-1.733-3.076-3.897l24.79-104.006c.513-2.152 2.719-3.897 4.931-3.897h5.58c2.211 0 3.592 1.733 3.08 3.897l-24.666 104.006c-.51 2.153-2.718 3.897-4.93 3.897zm28.854 1.41c2.21 0 4.431-1.747 4.954-3.878L206.391 52.4c.526-2.143-.834-3.879-3.055-3.879h-6.793c-2.213 0-4.436 1.747-4.959 3.879l-25.738 104.917c-.525 2.142.833 3.878 3.05 3.878h6.785z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoProtools(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M128 220c-50.81 0-92-41.19-92-92s41.19-92 92-92s92 41.19 92 92s-41.19 92-92 92m-.5-139.23c27.094 0 36.41 18.143 43.998 31.187c10.798 18.56 11.931 31.154 35.889 40.868C210 140.22 210 141.57 210 127.5c0-45.563-36.937-82.5-82.5-82.5S45 81.937 45 127.5c0 11.61-1.603 13.593 2.549 25.325c20.955-6.887 30.758-29.192 34.67-39.144c5.526-14.056 18.187-32.91 45.281-32.91zm.5 24.615c-17.128 0-31.173 32.022-34.929 39.68c-3.756 7.659-16.584 25.445-34.306 29.902c14.782 22.217 40.049 36.858 68.735 36.858c28.448 0 53.532-14.398 68.364-36.305c-14.832-10.778-22.641-15.315-33.531-30.455c-8.994-12.504-17.205-39.68-34.333-39.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoRackext(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m33.697 220.076l-1.353-2.697l-1.822-3.49l-.06-167.247l1.13-4.875l3.334-5.129l4.868-2.742l176.147.113l5.172 2.48l4.94 4.171l1.59 3.999l-.184 168.706l-2.285 5.073l-3.301 3.74l-5.383 1.487l-173.347-.142l-5.637-1.165zm5.027-7.494l1.269 2.21l1.879 1.472l174.69-.408l1.473-1.384l1.523-2.512l-.162-166.463l-1.589-2.205l-1.59-1.23l-174.801.157l-2.05 1.219l-1.235 2.285z"/><path d="m54.563 86.647l50.023-.095l11.73 2.64l8.785 4.722l5.28 6.955l2.44 4.16l1.066 6.449l.406 8.044l-1.283 8.323l-4.355 5.89l-5.534 4.47l-5.077 3.755l-5.89 1.115l-.515 1.524l26.001 38.956l-26.918.089l-21.057-36.675l-12.919.453l-.241 36.256l-22.098-.049zm21.714 18.215l.397 25.225l23.404-.26s5.18-1.64 5.484-2c.305-.358 4.266-4.664 4.266-4.664l1.013-4.767l-.974-6.355l-3.947-5.125l-4.87-1.999zm68.863 19.452l7.362-6.918l8.48-4.102l11.376-.925l12.086 1.32l9.293 6.247l5.532 7.231l2.742 6.704s2.489 8.653 2.489 9.06c0 .406-.104 12.441-.104 12.441l-45.375.191l.508 4.215l3.351 5.18l4.926 3.453l9.667.319l6.246-3.64l2.59-4.152l18.495 6.24l-5.69 8.175l-6.253 5.28l-8.555 4.61l-10.354.606l-12.385-.414l-10.15-5.126l-7.326-7.228l-4.432-10.108l-1.742-10.123l.062-10.343l3.047-10.036zm14.234 16.835l24.116.7l.002-4.013l-1.88-4.773l-3.402-3.402l-8.023-1.066l-6.145 2.05l-3.046 4.307zM66.64 207h-2.62c-1.569-2.478-2.543-4-2.922-4.565a8.93 8.93 0 0 0-1.198-1.456c-.419-.405-.813-.66-1.182-.768c-.369-.107-.857-.161-1.466-.161h-2.406V207h-2.084v-15.748h6.971c.953 0 1.751.06 2.396.183a3.78 3.78 0 0 1 2.82 2.245a4.62 4.62 0 0 1 .392 1.88c0 1.188-.376 2.144-1.128 2.868c-.752.723-1.848 1.181-3.287 1.375v.043c1.053.465 2.04 1.421 2.965 2.868zm-11.794-8.82h4.468c.903 0 1.617-.076 2.144-.23a2.295 2.295 0 0 0 1.262-.897a2.52 2.52 0 0 0 .472-1.493c0-.717-.261-1.325-.784-1.827c-.523-.5-1.386-.752-2.589-.752h-4.973v5.2zM78.242 207h-2.03c-.186-.358-.319-.827-.398-1.407c-1.296 1.117-2.696 1.676-4.2 1.676c-1.203 0-2.152-.301-2.846-.903c-.695-.601-1.042-1.396-1.042-2.385c0-.945.334-1.724 1.004-2.336c.67-.612 1.832-1.019 3.486-1.22l1.805-.268c.673-.107 1.228-.24 1.665-.397c0-.58-.018-.99-.054-1.23a1.82 1.82 0 0 0-.317-.747c-.176-.258-.457-.465-.843-.623c-.387-.157-.906-.236-1.558-.236c-.83 0-1.488.154-1.971.462c-.484.308-.815.88-.994 1.718l-1.89-.247c.207-1.167.74-2.044 1.6-2.631c.86-.588 2.041-.881 3.545-.881c1.368 0 2.372.202 3.013.607c.641.404 1.035.893 1.182 1.466c.147.573.22 1.31.22 2.213v2.6c0 1.582.03 2.644.091 3.185a4.13 4.13 0 0 0 .532 1.584m-2.556-5.017v-.698c-1.01.337-2.246.604-3.707.8c-1.46.197-2.19.815-2.19 1.853c0 .523.196.955.59 1.295c.394.34.96.51 1.697.51c.953 0 1.792-.285 2.52-.854c.726-.57 1.09-1.538 1.09-2.906m14.244 1.075c-.251 1.453-.824 2.519-1.72 3.195a4.912 4.912 0 0 1-3.04 1.016c-1.653 0-2.932-.543-3.834-1.628c-.902-1.085-1.354-2.523-1.354-4.313c0-1.482.26-2.67.78-3.56c.518-.893 1.168-1.52 1.949-1.88c.78-.362 1.6-.543 2.46-.543c1.174 0 2.179.304 3.013.913c.834.609 1.348 1.504 1.542 2.685l-1.87.29c-.193-.766-.506-1.342-.94-1.73c-.433-.386-.986-.58-1.66-.58c-1.08 0-1.9.384-2.46 1.15c-.558.767-.837 1.837-.837 3.212c0 1.404.274 2.487.822 3.25c.548.762 1.337 1.144 2.368 1.144c.831 0 1.485-.238 1.96-.715c.477-.476.776-1.19.898-2.143zM101.037 207h-2.374l-3.792-5.833l-1.364 1.3V207h-1.934v-15.748h1.934v8.97l4.576-4.62h2.514l-4.362 4.222zm19.701 0h-11.762v-15.748h11.397v1.87h-9.313v4.812h8.711v1.869h-8.711v5.328h9.678zm12.02 0h-2.384l-2.965-4.512l-2.997 4.512h-2.352l4.178-5.94l-3.845-5.457h2.395c1.561 2.356 2.413 3.67 2.557 3.942l2.793-3.942h2.31l-3.943 5.35zm6.07-.021c-.573.114-1.056.171-1.45.171c-.845 0-1.461-.148-1.848-.445a1.813 1.813 0 0 1-.709-1.134c-.086-.458-.129-1.095-.129-1.912v-6.563h-1.407v-1.493h1.407v-2.836l1.934-1.171v4.007h1.934v1.493h-1.934v6.67c0 .574.06.98.177 1.22c.118.24.417.36.897.36c.28 0 .566-.029.86-.086zm11.57-5.189h-8.54c.078 1.26.438 2.224 1.079 2.89c.641.666 1.441.999 2.4.999c.739 0 1.358-.195 1.86-.586c.5-.39.884-.98 1.149-1.767l1.987.247c-.322 1.189-.913 2.102-1.772 2.74c-.86.637-1.934.956-3.223.956c-1.712 0-3.05-.525-4.012-1.574c-.964-1.05-1.445-2.483-1.445-4.302c0-1.798.467-3.255 1.402-4.372c.934-1.118 2.25-1.676 3.947-1.676c.831 0 1.631.182 2.401.548c.77.365 1.424.986 1.96 1.863c.538.878.806 2.222.806 4.034zm-1.988-1.59c-.079-1.16-.438-1.996-1.08-2.508c-.64-.512-1.34-.768-2.1-.768c-.909 0-1.657.304-2.245.913c-.587.609-.927 1.396-1.02 2.363zm13.632 6.8h-1.944v-6.929c0-1.167-.21-1.97-.629-2.406c-.419-.437-1-.655-1.745-.655c-.573 0-1.116.14-1.628.419a2.374 2.374 0 0 0-1.09 1.181c-.215.509-.322 1.229-.322 2.16V207h-1.934v-11.397h1.74v1.61h.043c.408-.622.911-1.09 1.51-1.401c.597-.312 1.29-.467 2.078-.467c.602 0 1.196.107 1.783.322c.588.215 1.048.537 1.38.967c.334.43.543.888.63 1.375c.085.487.128 1.146.128 1.976zm9.281-8.207c-.086-.623-.338-1.09-.757-1.402c-.419-.311-1-.467-1.746-.467c-.737 0-1.334.125-1.788.376c-.455.25-.682.616-.682 1.096c0 .45.18.773.542.966c.362.194 1.076.43 2.143.71c1.203.307 2.104.58 2.702.816c.598.236 1.063.567 1.396.993c.333.426.5 1.02.5 1.778c0 1.017-.416 1.873-1.246 2.568c-.831.694-1.941 1.042-3.33 1.042c-1.447 0-2.577-.307-3.39-.919c-.812-.612-1.312-1.53-1.498-2.755l1.923-.29c.1.795.395 1.389.886 1.783c.49.394 1.176.59 2.057.59c.845 0 1.493-.177 1.944-.531c.452-.355.677-.786.677-1.294c0-.344-.104-.62-.311-.828a1.955 1.955 0 0 0-.78-.467c-.31-.104-1.01-.288-2.1-.553c-1.618-.387-2.687-.856-3.206-1.407a2.783 2.783 0 0 1-.779-1.977c0-.952.38-1.737 1.14-2.352c.758-.616 1.779-.924 3.06-.924c1.361 0 2.42.266 3.175.8c.756.534 1.208 1.33 1.359 2.39zm6.574 8.207h-1.933v-11.397h1.933zm0-13.557h-1.933v-2.191h1.933zm7.595 1.902c1.54 0 2.816.501 3.83 1.504c1.013 1.002 1.52 2.427 1.52 4.275c0 2.242-.548 3.826-1.644 4.753c-1.095.928-2.33 1.392-3.706 1.392c-1.44 0-2.689-.477-3.749-1.43c-1.06-.952-1.59-2.463-1.59-4.532c0-2.006.512-3.5 1.536-4.485c1.025-.985 2.292-1.477 3.803-1.477m0 10.334c1.103 0 1.94-.408 2.514-1.225c.573-.816.86-1.89.86-3.223c0-1.425-.325-2.499-.973-3.222c-.648-.724-1.448-1.085-2.4-1.085c-.982 0-1.788.367-2.418 1.1c-.63.735-.945 1.829-.945 3.283c0 1.446.32 2.537.961 3.27c.641.735 1.442 1.102 2.401 1.102M202.4 207h-1.944v-6.929c0-1.167-.21-1.97-.628-2.406c-.42-.437-1.001-.655-1.746-.655a3.35 3.35 0 0 0-1.627.419a2.37 2.37 0 0 0-1.09 1.181c-.216.509-.323 1.229-.323 2.16V207h-1.934v-11.397h1.74v1.61h.044c.408-.622.91-1.09 1.509-1.401c.598-.312 1.29-.467 2.078-.467c.602 0 1.196.107 1.784.322c.587.215 1.047.537 1.38.967a3.1 3.1 0 0 1 .628 1.375c.086.487.13 1.146.13 1.976z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoReaper(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M50.455 102.855c0 25.177 39.647 105.629 77.265 105.629h.001c37.619 0 77.397-76.894 77.397-108.531c0-20.003-18.606-52.373-77.62-52.373c-59.013 0-77.043 35.758-77.043 55.275m50.039 5.634C86.402 93.277 77.656 77.857 77.656 77.857s14.41-7.175 23.824-9.091c.802-.163 1.62-.338 2.458-.516c8.997-1.916 20.189-4.298 36.328 1.496c33.579 12.056 47.896 34.606 47.896 34.606s-21.609-17.401-45.404-20.895c-13.707-2.013-39.203 6.152-39.203 6.152s8.862 9.38 16.075 23.114c1.137 2.165 2.245 4.147 3.306 6.045c5.668 10.137 9.974 17.838 9.974 37.93c0 14.659-11.08 36.443-11.08 36.443s6.223-28.383 1.204-50.825c-3.627-16.223-6.425-18.719-13.02-24.603c-2.528-2.255-5.615-5.008-9.52-9.224" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoReason(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><circle cx="128.5" cy="139.5" r="18.5"/><path d="m89.269 69.823l39.943-23.625l38.081 23.376l.243 44.513l-18.906 12.071s-11.12-10.516-21.235-10.518c-10.117-.002-19.2 10.51-19.2 10.51l-17.83-12.01zm4.64 40.261l32.733-19.52l-.34-38.05L92.692 72.5l1.217 37.585zm1.83 4.806c1.828 0 12 6.919 12 6.919s9.188-10.227 20.027-9.894c10.838.333 21.319 10.165 21.319 10.165l14.32-8.436l-34.5-18.498S93.911 114.89 95.74 114.89zm-8.066 94.924l-40.431-22.78l1.203-44.667L87.873 118.9l18.907 11.338s-3.414 14.388 1.642 23.15c5.057 8.762 18.569 11.871 18.569 11.871l-.486 21.446zm32.547-24.15l-33.272-18.586l-32.781 19.318l34.112 19.115l31.94-19.846zm3.247-3.987c-.915-1.583-.008-13.852-.008-13.852s-13.317-3.344-18.448-12.897c-5.13-9.552-1.99-23.045-1.99-23.045l-14.467-8.184l1.231 39.127s34.596 20.434 33.682 18.85zm84.415-40.856l.488 46.404l-39.284 21.292l-38.671-22.046l-1-21.41s13.532-2.87 18.592-11.63c5.06-8.76 2.631-23.383 2.631-23.383l18.316-10.436zm-37.187-16.112l.54 38.108l33.12 18.73l-.502-39.099zm-4.077.183l-11.992 6.932s2.129 14.571-3.579 23.79c-5.708 9.22-17.329 11.881-17.329 11.881l.146 14.62l32.27-19.63z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoRewire(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m32.063 44.286l2.843-4.92l4.57-3.282l4.05-1.846l175.136.464l4.774 2.64l5.078 4.366l.753 175.396l-3.18 3.453l-6.357 2.945l-177.641.744l-6.153-2.974l-4.41-5.332zm9.45 142.864l178.043-.12l.505-142.383l-177.91-.544zm37.308 25.472l23.651.21l.11-3.368l-19.709-.18l.196-3.668l19.354.08l.157-2.539l-19.377-.16s-.047-4.044.004-4.53c.051-.486 19.394.486 19.394.486l-.05-2.54l-23.532.094zm111-1l23.651.21l.11-3.368l-19.709-.18l.196-3.668l19.354.08l.157-2.539l-19.377-.16s-.047-4.044.004-4.53c.051-.486 19.394.486 19.394.486l-.05-2.54l-23.532.094zm-79.678-15.033l-.123 15.666l31.106.145l.022-15.746l-5.854-.08l.169 11.848l-7.801.145l-.158-11.88l-5.3-.181l.116 12.21l-7.603-.235l-.213-11.94l-4.36.048zm38.623-.09l-.372 16.185l4.356-.069s-.15-15.84-.277-16.093c-.127-.254-3.707-.022-3.707-.022zm-100.627.146l-.137 16.053l4.494-.104l-.009-6.305l9.838.127l3.846 5.393l6.25.048l-4.991-5.361l4.25-2.934l-.199-7.529zm112 0l-.137 16.053l4.494-.104l-.009-6.305l9.838.127l3.846 5.393l6.25.048l-4.991-5.361l4.25-2.934l-.199-7.529z"/><path d="m52.736 146.965l22.141-.315l14.828-2.873l14.016-5.64l5.784-3.383l6.845 4.542l8.328 4.062l-11.399 6.907l-11.072 5.586l-9.74 3.453l-11.501 1.717l-28.207.52zm-.381-24.11l12.475.095l7.313-1.742l7.312-3.179l9.19 4.548l9.446 4.57l-11.396 6.397l-11.894 2.742l-9.946 1.421l-13.021.2z"/><path d="m52.59 107.09l14.524.808l13.61 4.672l13.609 6.196l11.88 5.672l12.187 8.329l11.68 7.718l10.764 7.1l14.263 6.317l14.422 5.993l7.617 1.32l31.999-.007l.04-15.175l-25.697.81l-15.682-4.613l-16.706-7.278l-13.117-9.222l-14.026-9.326l-14.377-7.84l-13.726-7.586l-14.735-6.662l-11.606-2.67l-16.705.187z"/><path d="m52.826 83.443l21.509.051l11.781 2.336l12.492 5.18l11.477 5.89l14.625 8.735l11.781 8.328l15.336 9.445l13.508 6.805l13.914 5.484l10.36 1.929l20.76.096l-.009-15.067l-17.731.711l-16.501-5.638l-15.785-9.023l-17.83-10.253l-16.393-10.765l-15.373-8.921l-14.518-6.713l-9.729-2.666l-4.563-.927l-29.015.639z"/><path d="m210.412 106.312l.233-16.148l-18.243.192l-8.508 1.625l-11.176 3.352l-8.715 4.266l-3.281 1.828l9.128 4.552l9.14 4.977l10.054-2.872l7.719-1.435zm-1.159-23.402l.482-15.852l-19.888.097l-14.246.095l-13.124 3.555l-13.837 6.398l-10.66 7.814l.078 2.008l6.703 3.961l7.312 5.078l9.75-5.536l11.07-3.794l14.016-3.084zM53.36 200.003l-.102 3.555l11.943.092l.508-3.691zm112 0l-.102 3.555l11.943.092l.508-3.691z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoStudioone(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m31.75 129.068l7.588 18.351c.633 1.531 2.491 2.791 4.144 2.815l43.245.607c6.077.086 14.583-3.22 19.008-7.393l32.867-30.996c4.017-3.788 5.72-11.056 3.799-16.245l-11.922-32.21l-6.356.202l3.774 8.691c1.757 4.05.57 9.77-2.666 12.793l-40.618 37.924c-3.23 3.016-9.422 5.461-13.851 5.461zm77.538 30.036c-1.112-3.12-.01-7.442 2.465-9.652l38.921-34.763c1.233-1.101 3.578-2.004 5.227-2.016l61.154-.463c2.214-.017 4.792 1.58 5.76 3.573l3.456 7.12c.481.991-.026 1.808-1.13 1.824l-51.144.765c-2.206.033-5.342 1.252-6.993 2.711l-38.281 33.856c-2.067 1.827-3.076 5.451-2.26 8.076l6.324 20.337c.328 1.055-.31 1.896-1.41 1.877l-9.29-.156c-.556-.01-1.157-.439-1.341-.955z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoTracktion(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m100.48 155.294l-8.369-4.528l-8.095 7.8l-14.076-14.045l7.676-7.838l-5.419-12.217H61.43l.519-20.188h10.248l5.927-12.28l-8.184-7.813L84.016 70.62l7.498 6.625l12.487-5.475V60.133h19.822V71.77l13.973 4.462l6.693-6.376L158.21 82.75s3.558-2.262 11.376-7.114C161.768 64.946 137.241 48 114.5 48C77.22 48 47 77.325 47 113.5c0 27.587 17.575 51.19 42.452 60.842c6.006-10.863 11.027-19.048 11.027-19.048z"/><path d="M106.6 109.603s8.717-1.092 21.788-5.425c13.071-4.334 22.25-8.786 33.151-14.044c10.902-5.258 21.195-12.167 21.195-12.167l5.306-4.068l12.14.588l-13.502 27.666c-.485.993.026 1.784 1.129 1.767l19.049-.298c1.108-.018 1.686.794 1.285 1.827l-4.077 10.499c-.4 1.026-1.627 1.843-2.713 1.825l-19.849-.328c-1.099-.018-2.394.775-2.889 1.764l-33.948 67.893c-1.235 2.47.002 4.436 2.753 4.391l9.749-.158l-7.373 14.651s-14.555 4.754-29.15 3.585c-14.594-1.169-19.287-3.464-23.673-7.678c-4.386-4.214-4.323-3.904-4.94-8.368c-.616-4.463 4.11-13.158 4.11-13.158l31.522-62.832l-25.815.589z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoVst(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.163 61.754c-.733.832-1.328 2.406-1.33 3.506L5.7 175.087c-.002 1.105.417 2.792.936 3.769l.248.469c.518.977 1.836 1.778 2.946 1.79l61.273.652c1.109.012 2.68-.584 3.498-1.318l13.167-11.835c.823-.74 2.387-1.336 3.492-1.332l31.571.114c1.106.004 2 .902 1.996 2.008l-.003 1.136a1.996 1.996 0 0 1-2.01 1.993l-10.968-.045a2.019 2.019 0 0 0-2.027 1.996l-.202 17.714a1.967 1.967 0 0 0 1.976 2.002l11.295-.015a2.064 2.064 0 0 1 2.06 1.988l.091 3.03a1.972 1.972 0 0 1-1.936 2.032l-30.826.635c-1.103.023-2.658-.558-3.482-1.303l-14.45-13.076c-.82-.742-2.387-1.347-3.484-1.351l-61.595-.23c-1.104-.004-2.704-.548-3.58-1.22l-.238-.183c-.872-.67-2.034-1.986-2.594-2.939l-.76-1.291c-.56-.954-1.016-2.627-1.018-3.717L.819 64.42c-.003-1.099.407-2.796.903-3.766l.268-.524c.501-.981 1.512-2.434 2.263-3.25l.67-.729c.75-.814 2.257-1.475 3.358-1.477l238.073-.288c1.106-.001 2.684.568 3.538 1.283l.295.247c.848.71 2.01 2.051 2.59 2.986l.767 1.236c.582.94 1.055 2.598 1.055 3.708l.007 111.612c0 1.108-.275 2.853-.62 3.915l-.223.683c-.343 1.054-1.273 2.522-2.085 3.286l-.671.631c-.809.761-2.357 1.383-3.461 1.39l-62.734.387c-1.103.007-2.649.614-3.474 1.376l-14.194 13.1c-.816.754-2.373 1.374-3.486 1.386l-24.876.269a2.03 2.03 0 0 1-2.051-1.982l-.148-6.918a2.049 2.049 0 0 0-2.037-1.998l-11.886.032a1.995 1.995 0 0 1-2-1.988l-.023-9.327a2 2 0 0 1 2-1.995l11.949-.01a2.072 2.072 0 0 0 2.06-2.003l.176-6.249c.03-1.105.943-2 2.047-2l25.666-.003c1.1 0 2.659.593 3.483 1.327l13.022 11.591c.824.733 2.386 1.326 3.481 1.325l60.165-.09c1.099-.002 2.494-.748 3.114-1.664l.367-.542c.621-.917 1.13-2.546 1.135-3.665l.53-110.113c.005-1.107-.507-2.728-1.153-3.634l-.638-.894c-.642-.9-2.048-1.629-3.166-1.627l-233.949.3c-1.106 0-2.598.678-3.33 1.509z"/><path d="M60.188 154.692c-1.1.03-2.457-.709-3.03-1.648l-35.294-57.9c-.574-.941-1.833-2.104-2.83-2.606l.675.34c-.99-.498-2.679-.899-3.795-.895l-.592.001a1.96 1.96 0 0 1-1.977-1.99l-.007.527a2.037 2.037 0 0 1 2.036-2l40.727-.096a1.909 1.909 0 0 1 1.93 1.991l.032-.812c-.044 1.103-.968 2.055-2.083 2.13l-4.445.294c-1.106.074-2.725.67-3.615 1.33l1.448-1.075c-.89.66-1.959 2.029-2.379 3.034l.509-1.217c-.424 1.015-.314 2.618.242 3.576l19.195 33.076c.558.96 1.437.962 1.972-.011l16.841-30.599c.532-.966 1.241-2.6 1.585-3.65l-.225.688c.343-1.05.164-2.668-.404-3.62l.526.881c-.566-.949-1.925-1.761-3.023-1.813l-5.214-.247c-1.103-.052-2.049-.981-2.112-2.098l.028.493a1.966 1.966 0 0 1 1.884-2.089l12.987-.551c1.103-.047 2.445-.857 3-1.816l1.526-2.637c.554-.956 1.705-2.302 2.552-2.989l3.96-3.216c.855-.694 2.385-1.578 3.415-1.973l6.71-2.572c1.031-.395 2.743-.849 3.843-1.016l6.755-1.025a29.984 29.984 0 0 1 3.975-.282l7.329.067a72.15 72.15 0 0 1 3.993.169l5.639.424c1.102.083 2.858.415 3.905.734l5.185 1.585c1.055.323 2.8.582 3.917.579l6.022-.015c1.108-.003 2.446-.782 2.992-1.748l.487-.86c.545-.964 1.879-1.798 2.99-1.866l2.219-.136a1.866 1.866 0 0 1 2.008 1.88l.023 14.055a2.023 2.023 0 0 0 1.999 2.014l58.083.369a2 2 0 0 1 1.984 2.014l-.086 15.916c-.006 1.106-.9 1.901-1.99 1.778l-.048-.005c-1.093-.124-2.267-1.066-2.626-2.115l-.2-.587a29.464 29.464 0 0 0-1.592-3.652l-.572-1.067c-.522-.972-1.652-2.312-2.517-2.985l-1.868-1.455a37.889 37.889 0 0 0-3.285-2.24l-.454-.27c-.946-.56-2.608-1.03-3.709-1.05l-9.342-.16a1.95 1.95 0 0 0-1.988 1.966l.218 51.76c.005 1.105.758 2.494 1.675 3.097l.121.08c.92.605 2.534 1.35 3.587 1.658l-.54-.158c1.061.31 2.808.563 3.928.564l2.99.002c1.108 0 2.006.9 2.007 1.998v-1a2.023 2.023 0 0 1-1.998 2.024l-45.808.617c-1.104.015-1.915-.867-1.812-1.965l-.084.894c.103-1.1 1.083-2.058 2.182-2.139l4.384-.32c1.102-.081 2.78-.57 3.75-1.095l.466-.25c.97-.525 2.01-1.803 2.331-2.873c.319-1.062.568-2.822.558-3.923l-.466-49.068a1.998 1.998 0 0 0-2.012-1.98l-5.2.052c-1.1.011-2.06.916-2.144 2.013l-.06.779c-.083 1.1-1.05 2.009-2.147 2.028l-1.761.03c-1.102.019-2.412-.762-2.92-1.734l-2.411-4.611c-.51-.977-1.584-2.365-2.406-3.11l-6.008-5.439c-.818-.74-2.278-1.758-3.254-2.269l-2.363-1.237a53.479 53.479 0 0 0-3.622-1.692l-2.814-1.163c-1.02-.423-2.74-.878-3.823-1.015l-3.209-.405a34.366 34.366 0 0 0-3.97-.222l-4.426.063c-1.101.016-2.839.308-3.888.655l-2.882.954a32.782 32.782 0 0 0-3.68 1.523l-.98.493c-.985.495-2.242 1.671-2.803 2.621l-.024.041c-.562.953-1.464 2.508-2.003 3.454l.735-1.29c-.544.955-.949 2.614-.903 3.722l.004.096c.045 1.1.471 2.796.954 3.794l-.03-.062c.482.995 1.604 2.316 2.502 2.948l2.705 1.902c.9.633 2.48 1.447 3.517 1.813l7.06 2.495c1.042.368 2.74.937 3.802 1.274l7.767 2.463c1.058.335 2.775.882 3.813 1.214l3.962 1.268c1.048.335 2.755.865 3.814 1.182l-.634-.19c1.059.318 1.91-.316 1.9-1.422l-.115-15.004a1.882 1.882 0 0 1 1.989-1.913l1.543.065c1.107.046 2.469.844 3.051 1.799l2.572 4.21c.578.948.446 2.374-.307 3.198l-.065.071c-.748.819-1.59 2.339-1.884 3.409l-.86 3.13c-.293 1.064-1.332 2.317-2.323 2.8l1.364-.665c-.99.482-2.26 1.636-2.465 1.972l.672-1.1c-.371.608.149 1.448 1.178 1.885l5.023 2.13a29.642 29.642 0 0 1 3.549 1.857l2.88 1.82c.938.593 2.274 1.747 3.002 2.599l1.66 1.942a29.86 29.86 0 0 1 2.34 3.232l.338.556c.572.943 1.074 2.609 1.12 3.712l.162 3.856a50.253 50.253 0 0 1-.025 4.007l-.038.716c-.06 1.106-.63 2.718-1.286 3.618l-1.4 1.92c-.65.892-1.855 2.2-2.697 2.926l-3.172 2.736c-.84.724-2.302 1.757-3.26 2.304l-4.104 2.342c-.962.549-2.609 1.195-3.696 1.447l-8.431 1.956a67.81 67.81 0 0 1-3.94.766l-2.302.362a33.538 33.538 0 0 1-3.98.335l-12.688.148c-1.102.013-2.877-.14-3.954-.342l-5.11-.953a643.126 643.126 0 0 1-3.923-.747l-5.78-1.123a28.383 28.383 0 0 0-3.96-.426l-2.372-.053c-1.102-.024-2.774.393-3.74.936l-.02.011c-.964.542-1.841 1.867-1.938 2.754l.193-1.774c-.106.98-1.084 1.818-2.185 1.872l-1.13.056a1.875 1.875 0 0 1-1.99-1.897l.048-26.778a2.063 2.063 0 0 1 1.993-2.046l1.447-.037c1.1-.028 2.266.803 2.604 1.846l1.179 3.647c.339 1.047 1.085 2.649 1.674 3.59l1.723 2.75c.586.934 1.692 2.326 2.477 3.116l2.667 2.681c.782.787 2.184 1.891 3.12 2.46l4.913 2.99c.94.573 2.575 1.22 3.66 1.447l6.109 1.281c1.08.227 2.844.4 3.964.386l9.204-.11c1.11-.014 2.881-.207 3.965-.433l5.368-1.119c1.08-.225 2.569-1.055 3.336-1.866l3.389-3.582c.761-.806 1.632-2.31 1.948-3.378l.312-1.054c.314-1.06.27-2.771-.092-3.804l-.538-1.534c-.365-1.04-1.345-2.462-2.195-3.18l-.232-.197c-.848-.716-2.357-1.649-3.376-2.086l-4.005-1.718a62.486 62.486 0 0 0-3.736-1.434l-6.347-2.16c-1.046-.355-2.735-.926-3.788-1.28l-5.682-1.908a85.778 85.778 0 0 0-3.823-1.165l-3.71-1.017a31.447 31.447 0 0 1-3.758-1.348l-2.275-1.02a135.624 135.624 0 0 1-3.63-1.704l-3.919-1.926a29.627 29.627 0 0 1-3.432-2.05L100 109.38c-.9-.643-2.23-1.833-2.971-2.66l.4.446c-.741-.826-1.69-2.31-2.131-3.342l.643 1.504c-.436-1.02-.245-2.555.425-3.426l-.707.92c.67-.873 1.886-2.159 2.737-2.89l1.939-1.667c.84-.722 2.36-1.61 3.404-1.986l-.446.16a55.237 55.237 0 0 1 3.818-1.19c1.07-.282 1.87-1.415 1.79-2.512l-.025-.357c-.08-1.105-1.04-2.044-2.142-2.098l-.222-.011c-1.103-.054-2.396.707-2.67 1.258l.674-1.355c-.373.748-1.55 1.522-2.642 1.73l-.834.158c-1.087.207-2.648.958-3.481 1.672l-.591.507c-.836.717-2.005 2.04-2.622 2.97l-.397.598a55.39 55.39 0 0 0-2.065 3.43L63.786 152.86c-.528.97-1.85 1.782-2.948 1.812z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWaveform(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M128 209c-44.735 0-81-36.265-81-81s36.265-81 81-81s81 36.265 81 81s-36.265 81-81 81m22.53-141.14A64.379 64.379 0 0 0 128.5 64a64.352 64.352 0 0 0-22.917 4.19a1317.969 1317.969 0 0 0 21.079 22.104c.77.79 2.042.806 2.831.038c0 0 12.798-12.235 21.038-22.472zm0 121.635c-8.239-10.237-21.05-22.665-21.05-22.665a1.975 1.975 0 0 0-2.806.041s-10.22 10.605-21.09 22.294a64.352 64.352 0 0 0 22.916 4.19a64.379 64.379 0 0 0 22.03-3.86M89.908 76.531s-26.02 17.694-26.02 51.538c0 33.845 26.02 53.465 26.02 53.465l38.09-39.787l38.376 38.29s27.389-21.385 27.389-52.066c0-30.682-27.739-51.678-27.739-51.678l-25.672 26.32l13.39 13.837l13.112-11.68s6.936 10.757 7.035 23.153c.098 12.396-6.653 24.904-6.653 24.904l-39.191-38.937l-37.495 38.69s-7.684-11.61-7.65-25.002c.035-13.393 7.008-24.966 7.008-24.966l12.75 12.488l14.996-12.609z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loop(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M24.849 72.002a8 8 0 0 1 8.027-7.983l189.787.462c4.42.01 8.004 3.61 8.004 8.026v111.986c0 4.422-3.587 8.04-8.008 8.08l-42.661.39A4.04 4.04 0 0 0 176 197v22.002c0 2.208-1.39 2.878-3.115 1.488l-35.31-28.463c-5.157-4.156-5.11-10.846.099-14.935l35.174-27.616c1.74-1.367 3.152-.685 3.152 1.534v21.482a4 4 0 0 0 3.992 4.009h30.683a4.02 4.02 0 0 0 4.013-3.991l.437-83.012a7.98 7.98 0 0 0-7.952-8.02l-158.807-.454c-4.415-.013-8.008 3.56-8.025 7.975l-.31 79.504a7.962 7.962 0 0 0 7.967 7.998H108a4.003 4.003 0 0 1 3.999 3.994v8.512a3.968 3.968 0 0 1-4.001 3.971l-75.502-.431c-4.417-.026-7.987-3.634-7.974-8.048l.326-112.496z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Metronome(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M64.458 228.867c-.428 2.167 1.007 3.91 3.226 3.893l121.557-.938c2.21-.017 3.68-1.794 3.284-3.97l-11.838-64.913c-.397-2.175-1.626-2.393-2.747-.487l-9.156 15.582c-1.12 1.907-1.71 5.207-1.313 7.388l4.915 27.03c.395 2.175-1.072 3.937-3.288 3.937H88.611c-2.211 0-3.659-1.755-3.233-3.92L114.85 62.533l28.44-.49l11.786 44.43c.567 2.139 2.01 2.386 3.236.535l8.392-12.67c1.22-1.843 1.73-5.058 1.139-7.185l-9.596-34.5c-1.184-4.257-5.735-7.677-10.138-7.638l-39.391.349c-4.415.039-8.688 3.584-9.544 7.912z"/><path d="M118.116 198.935c-1.182 1.865-.347 3.377 1.867 3.377h12.392c2.214 0 4.968-1.524 6.143-3.39l64.55-102.463c1.18-1.871 3.906-3.697 6.076-4.074l9.581-1.667c2.177-.379 4.492-2.38 5.178-4.496l4.772-14.69c.683-2.104-.063-5.034-1.677-6.555L215.53 54.173c-1.609-1.517-4.482-1.862-6.4-.78l-11.799 6.655c-1.925 1.086-3.626 3.754-3.799 5.954l-.938 11.967c-.173 2.202-1.27 5.498-2.453 7.363l-72.026 113.603z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m55.618 178.111l75.645-94.377s.03-15.79 12.989-26.864c15.797-13.498 40.626-11.241 53.643 3.201c13.016 14.443 16.545 38.02-1.25 54.93c-11.078 10.526-23.552 10.425-23.552 10.425l-94.71 75.646l-9.344-9.14l-17.266 16.088s-1.429 1.188-2.844-.578c-1.415-1.765-.71-3.282-.71-3.282l17.039-16.203zm17.729-1.144l5.687 5.484l75.162-59.15s-8.01-4.416-12.813-9.469c-4.804-5.053-7.58-11.981-7.58-11.981zm75.827-105.668s-13.299 18.34 2.398 33.329c15.698 14.99 33.564 2.4 33.564 2.4s-16.071-3.727-24.79-12.39c-8.72-8.662-11.172-23.34-11.172-23.34zm5.393-4.513s3.702 14.109 12.601 23.073c8.899 8.965 22.874 12.662 22.874 12.662s12.885-19.52-3.271-34.318c-16.157-14.798-32.204-1.417-32.204-1.417"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Midiplug(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M128 193.901c-13.606 0-21.823 9.814-23.434 22.258C62.374 207.79 36 172.65 36 128c0-50.81 41.19-92 92-92s92 41.19 92 92c0 44.21-25.713 77.476-67.501 86.16c-1.346-8.684-11.008-20.259-24.499-20.259m-.244-18.45c16.601 0 29.657 10.87 32.244 17.732c31.34-6.861 42.826-42.019 42.826-65.183c0-40.149-36.718-76-74.826-76c-38.108 0-75.313 35.851-75.313 76c0 35.28 21.881 61.702 43.313 66.628c2.095-10.012 15.155-19.178 31.756-19.178z"/><circle cx="80" cy="125" r="11"/><circle cx="95" cy="92" r="11"/><circle cx="128" cy="79" r="11"/><circle cx="161" cy="92" r="11"/><circle cx="174" cy="124" r="11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modrandom(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M32 147.996a3.917 3.917 0 0 1 3.991-3.934l24.305.376c2.204.034 3.98 1.852 3.967 4.07l-.239 39.483A3.973 3.973 0 0 0 67.996 192h8.008A4 4 0 0 0 80 187.998V52.002A4.005 4.005 0 0 1 84.01 48h39.98a4 4 0 0 1 4.01 4.003v103.994a4 4 0 0 0 3.996 4.003h8.008a4.002 4.002 0 0 0 3.996-4.01v-39.98a3.93 3.93 0 0 1 4.004-3.947l23.992.374A3.93 3.93 0 0 0 176 108.5V84a4.005 4.005 0 0 1 4.01-4h39.98a4.006 4.006 0 0 1 4.01 3.998v56.504a3.999 3.999 0 0 1-3.992 3.998h-7.516a4 4 0 0 1-3.992-4.01v-39.98a4.003 4.003 0 0 0-3.996-4.01h-8.008a4.072 4.072 0 0 0-4.06 4.005l-.372 23.49a4.08 4.08 0 0 1-4.068 4.005h-23.992a4.006 4.006 0 0 0-4.004 4.01v39.98a4.01 4.01 0 0 1-4.01 4.01h-39.98a4 4 0 0 1-4.01-4.003V68.003A4 4 0 0 0 108.004 64h-8.008A4 4 0 0 0 96 68.002v135.996A4.005 4.005 0 0 1 91.99 208H52.01a4.01 4.01 0 0 1-4.01-4.01v-39.98a4.003 4.003 0 0 0-3.996-4.01h-8.008A3.996 3.996 0 0 1 32 156.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modsawdown(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M222.032 208c2.208 0 2.618-1.144.925-2.546L37.105 51.546c-1.698-1.406-3.074-.76-3.074 1.464v150.98a4.003 4.003 0 0 0 3.996 4.01h8.008a3.996 3.996 0 0 0 3.996-4.007V84.007c0-2.213 1.387-2.861 3.079-1.465l148.842 122.916c1.7 1.404 4.87 2.542 7.078 2.542z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modsawup(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M36 208c-2.21 0-2.619-1.144-.926-2.546L220.926 51.546c1.698-1.406 3.074-.76 3.074 1.464v150.98a4.003 4.003 0 0 1-3.996 4.01h-8.008a3.996 3.996 0 0 1-3.996-4.007V84.007c0-2.213-1.387-2.861-3.079-1.465L56.08 205.458C54.379 206.862 51.209 208 49 208z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modsh(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M32 195.996A4.002 4.002 0 0 1 36.01 192H80V52.002A4.005 4.005 0 0 1 84.01 48h39.98a4.022 4.022 0 0 1 4.024 4.007l.472 135.486a4.02 4.02 0 0 0 4.012 4.007h56.504a3.999 3.999 0 0 0 3.998-4.007V52.007A4.01 4.01 0 0 1 197.005 48h22.99c2.212 0 4.065 1.792 4.125 3.992c0 0 .327 7.542-.12 12.008c-.09.895-16 0-16 0v139.998a4 4 0 0 1-4.003 4.002h-87.994a4.004 4.004 0 0 1-4.003-4.002V68.002A3.999 3.999 0 0 0 108.004 64h-8.008A4 4 0 0 0 96 68.002V208H36.003A4 4 0 0 1 32 204.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modsine(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M48 128c-1.955-29.248 19.364-64 37.364-64c18 0 36.136 13.843 36.136 64.5s19.136 80.5 49.136 80.5c30 0 53.364-40.125 53.364-80.5c-8.182 0-7.273-.752-16 0c0 32.35-20.455 64.45-37.364 64.45s-33.909-13.542-33.909-64.45S120.273 48 85.364 48C50.454 48 32 88.626 32 127.748c6 0 8.364.252 16 .252"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modsquare(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M35.996 208A4 4 0 0 1 32 204.005V51.995A3.997 3.997 0 0 1 36.003 48h87.994a4.018 4.018 0 0 1 4.017 4.007l.472 135.486a4.05 4.05 0 0 0 4.01 4.032l71.507.45a3.971 3.971 0 0 0 3.997-3.977V52.002A3.999 3.999 0 0 1 211.996 48h8.008A4 4 0 0 1 224 51.995v152.01a3.992 3.992 0 0 1-4.003 3.995H116.003a4.004 4.004 0 0 1-4.003-4.002V68.002A4.004 4.004 0 0 0 107.997 64H52.003A4.005 4.005 0 0 0 48 68.002v135.996A3.999 3.999 0 0 1 44.004 208z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modtri(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M35.996 208c-2.207 0-3.077-1.532-1.935-3.434l91.878-153.132c1.138-1.896 2.98-1.902 4.122 0l91.878 153.132c1.138 1.896.272 3.434-1.935 3.434h-8.008c-2.207 0-4.922-1.54-6.054-3.421L130.058 78.42c-1.137-1.89-2.984-1.881-4.116 0L50.058 204.58c-1.138 1.889-3.848 3.42-6.054 3.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modularplug(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M128 226c-54.124 0-98-43.876-98-98s43.876-98 98-98s98 43.876 98 98s-43.876 98-98 98m-1-8c49.706 0 90-40.294 90-90s-40.294-90-90-90s-90 40.294-90 90s40.294 90 90 90"/><path d="M128 197c-38.108 0-69-30.668-69-68.5S89.892 60 128 60c38.108 0 69 30.668 69 68.5S166.108 197 128 197m0-6.29c34.22 0 61.96-27.54 61.96-61.511c0-33.971-27.74-61.51-61.96-61.51s-61.96 27.539-61.96 61.51s27.74 61.51 61.96 61.51z"/><path d="M128 170c-23.196 0-42-18.804-42-42s18.804-42 42-42s42 18.804 42 42s-18.804 42-42 42"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mono(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M127 210c-44.735 0-81-36.265-81-81s36.265-81 81-81s81 36.265 81 81s-36.265 81-81 81m1-21c34.794 0 63-27.087 63-60.5S162.794 68 128 68s-63 27.087-63 60.5S93.206 189 128 189"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mute(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M208.552 206.834h-25v-92.492l-43.987 87.835c-1.081 2.238-2.592 3.935-4.532 5.09c-1.94 1.155-4.024 1.733-6.25 1.733c-2.163 0-4.183-.578-6.06-1.733c-1.876-1.155-3.355-2.852-4.436-5.09l-44.179-87.835v92.492H49.3V63.548c0-3.25.843-6.137 2.528-8.665c1.686-2.527 3.896-4.223 6.632-5.09a11.647 11.647 0 0 1 4.007-.379a12.01 12.01 0 0 1 3.865.975c1.24.541 2.37 1.3 3.387 2.274c1.018.975 1.877 2.148 2.576 3.52l56.488 111.445L185.27 56.183c1.463-2.744 3.483-4.693 6.06-5.848c2.576-1.156 5.263-1.336 8.062-.542c2.672.867 4.867 2.563 6.584 5.09c1.718 2.528 2.576 5.416 2.576 8.665z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M80.033 60.858c0-14.183 8.718-16.43 13.894-10.454c5.176 5.977 47.925 60.873 52.46 68.201c4.535 7.329 5.434 13.082.086 20.28c-5.349 7.198-45.27 61.462-52.015 67.2c-6.745 5.74-14.543 3.358-14.543-10.567c0-13.926.118-120.476.118-134.66m51.265 71.792c1.34-1.758 1.357-4.612.03-6.388L98.4 82.205c-1.323-1.77-2.396-1.405-2.396.792L96 175.003c0 2.207 1.088 2.568 2.424.814zm44.866-78.858c.17-2.538-1.22-4.792-3.45-4.792h-8.193c-3.655 0-3.52.682-3.52 3.859c0 3.176.394 147.132.46 149.86c.064 2.729 1.06 4.454 3.219 4.454s6.252 0 8.33-.006c2.079-.007 3.088-1.897 3.088-4.033s-.104-146.803.066-149.342"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Open(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25.723 176.132c.283-8.428-.284-85.83-.284-95.335c0-9.505 8.156-17.14 17.993-17.327c9.837-.188 55.149.172 55.149.172c1.103.008 2.622.656 3.4 1.447L120.47 83.86c1.164 1.183 3.452 2.092 5.107 2.03c0 0 44.654-1.705 54.282-1.705c9.628 0 17.035 8.52 17.035 16.647c0 8.127.489 14.842.489 14.842a1.92 1.92 0 0 0 2.109 1.809s.86-.386 15.122-.1c14.263.288 21.126 13.712 14.263 24.258c-6.863 10.546-24.24 36.296-28.111 41.618c-3.872 5.322-7.694 8.812-15.074 8.253c-7.38-.56-132.897.813-142.36.813c-9.465 0-17.893-7.765-17.61-16.193zm18.006-93.708l-.71 75.569c-.01 1.113.454 1.25 1.038.316c0 0 18.173-29.148 21.523-33.172c3.35-4.025 7.78-7.653 14.48-7.653s99.329-.614 99.329-.614s1.173-14.187-.386-15.627c-1.559-1.44-61.818-.43-61.818-.43c-1.099.018-2.659-.584-3.468-1.344L94.33 81.244c-.808-.76-2.36-1.366-3.466-1.354l-45.114.505a2.054 2.054 0 0 0-2.02 2.03zm9.33 93.533l130.564.154c1.112.001 2.52-.736 3.154-1.647l27.435-39.435c.631-.907.251-1.65-.855-1.653l-132.038-.268c-1.097-.003-2.494.742-3.101 1.662l-26.066 39.518c-.609.923-.2 1.668.908 1.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paste(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M47.174 74.383c0-9.94 8.067-17.87 17.997-17.712l21.398.34s4.011-10.206 17.826-10.206c13.814 0 16.6 9.58 16.6 9.58h23.35c8.841 0 16.066 7.169 16.137 16.007l.123 15.5l11.998-.08c2.206-.015 5.354 1.143 7.02 2.577l27.06 23.28c1.671 1.439 3.026 4.403 3.026 6.608v85.158a3.984 3.984 0 0 1-4.005 3.985l-106.865-.5a4.98 4.98 0 0 1-4.953-5.025l.247-27.695l-30.963.195A15.87 15.87 0 0 1 47.174 160.5zM104 72c4.418 0 8-3.134 8-7s-3.582-7-8-7s-8 3.134-8 7s3.582 7 8 7m-39.622 2.883l1.193 82.848a2.004 2.004 0 0 0 2.037 1.973l23.907-.298a2.043 2.043 0 0 0 2.017-2.023l.23-53.884a15.692 15.692 0 0 1 16.06-15.678l32.068.64a1.975 1.975 0 0 0 2.014-1.966l.12-11.571a2.042 2.042 0 0 0-1.986-2.051l-20.593-.477s-4.914 10.667-17.418 10.667s-17.133-11.101-17.133-11.101l-20.544.84a2.067 2.067 0 0 0-1.972 2.08zm107.671 33.413c-1.674-1.437-2.943-.822-2.833 1.401l.727 14.677c.11 2.21 1.986 4.01 4.198 4.017l17.418.062c2.209.008 2.647-1.146.968-2.587zm-59.97-3.415l-.326 88.575l80.378.57V144.46h-32.189c-4.418 0-8.019-3.575-8.041-8.007l-.167-32.275z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M46.677 64.652c0-9.362 7.132-17.387 16.447-17.394c9.315-.007 24.677.007 34.55.007c9.875 0 17.138 7.594 17.138 16.998c0 9.403-.083 119.094-.083 127.82c0 8.726-7.58 16.895-16.554 16.837c-8.975-.058-25.349.115-34.963.058c-9.614-.058-16.646-7.74-16.646-17.254c0-9.515.11-117.71.11-127.072zm14.759.818s-.09 118.144-.09 123.691c0 5.547 3.124 5.315 6.481 5.832c3.358.518 21.454.47 24.402.47c2.947 0 7.085-1.658 7.167-6.14c.08-4.483-.082-119.507-.082-123.249c0-3.742-4.299-4.264-7.085-4.66c-2.787-.395-25.796 0-25.796 0zm76.664-.793c.027-9.804 7.518-17.541 17.125-17.689c9.606-.147 25.283.148 35.004.148c9.72 0 17.397 8.52 17.397 17.77s-.178 117.809-.178 127c0 9.192-7.664 17.12-16.323 17.072c-8.66-.05-26.354 0-34.991.048c-8.638.05-17.98-8.582-18.007-17.783c-.027-9.201-.055-116.763-.027-126.566m16.917.554s-.089 118.145-.089 123.692c0 5.547 3.123 5.314 6.48 5.832c3.359.518 21.455.47 24.402.47c2.948 0 7.086-1.659 7.167-6.141c.081-4.482-.08-119.506-.08-123.248c0-3.742-4.3-4.265-7.087-4.66c-2.786-.396-25.796 0-25.796 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M32 160L166.394 26.643a4.001 4.001 0 0 1 5.654.026l57.837 58.237a4.034 4.034 0 0 1-.007 5.676L97.348 223.59L32 224zm16.797 5.594V208h40.488l121.92-121.396L180.57 56.56L64.656 175.772a3.937 3.937 0 0 1-5.624.037z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phase(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m116.589 182.742l-7.405 20.346a4 4 0 0 1-5.125 2.396l-7.525-2.738a4 4 0 0 1-2.386-5.13l7.435-20.427C83.963 167.623 72 148.959 72 127.5C72 96.296 97.296 71 128.5 71c3.877 0 7.663.39 11.32 1.134l6.996-19.222a4 4 0 0 1 5.125-2.396l7.525 2.738a4 4 0 0 1 2.386 5.13l-6.968 19.142C172.796 87.002 185 105.826 185 127.5c0 31.204-25.296 56.5-56.5 56.5c-4.086 0-8.071-.434-11.911-1.258m5.173-14.213A41.32 41.32 0 0 0 128 169c22.644 0 41-18.356 41-41c0-14.855-7.9-27.864-19.727-35.056zm-15.035-5.473l27.51-75.585A41.32 41.32 0 0 0 128 87c-22.644 0-41 18.356-41 41c0 14.855 7.9 27.864 19.727 35.056"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M59 61.922c0-9.768 13.016-15.432 22.352-11.615c10.695 7.017 101.643 58.238 109.869 65.076c8.226 6.838 10.585 17.695-.559 25.77c-11.143 8.074-99.712 60.203-109.31 64.73c-9.6 4.526-21.952-1.632-22.352-13.088c-.4-11.456 0-121.106 0-130.873m13.437 8.48c0 2.494-.076 112.852-.216 115.122c-.23 3.723 3 7.464 7.5 5.245c4.5-2.22 97.522-57.704 101.216-59.141c3.695-1.438 3.45-5.1 0-7.388C177.488 121.952 82.77 67.76 80 65.38c-2.77-2.381-7.563 1.193-7.563 5.023z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pointer(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m56.402 54l157.203 70l-37.957 37l29.604 29.188c1.58 1.558 1.649 4.135.165 5.767l-6.73 7.405a3.756 3.756 0 0 1-5.5.122L163 173.092l-39 39.812zm32.593 30.615c-2.503-1.157-3.628-.05-2.514 2.471l41.64 94.202c.893 2.021 2.859 2.38 4.397.793l48.982-50.53c1.535-1.584 1.15-3.62-.847-4.545l-91.658-42.39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Powerswitch(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m118 40.5l5-5h9.5c-1 0 4.5 5.5 4.5 5.5v81.5l-5.5 4.5H123l-5-4.5z"/><path d="M89.887 58.547s-.476 4.792-1.43 6.23c-20.973 13.42-33.367 35.466-34.32 63.264c1.43 41.218 34.32 74.288 74.837 74.288c40.518 0 73.408-33.07 72.932-75.246c.476-25.881-13.347-48.886-32.891-62.306c-1.43-2.875-2.383-5.751-2.383-5.751l5.243-8.148s1.907-1.917 5.243-1.438c26.218 17.254 43.378 46.01 43.378 78.601c0 50.804-41.47 92.021-91.998 92.021c-50.528 0-92.475-41.217-91.998-92.02c0-33.07 17.637-61.827 42.9-79.56c2.86-1.917 5.72 3.355 5.72 3.355z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresetA(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M211.97 210h-29.199v-39.531H80.125V210h-29.2v-80.522c0-11.755 2.022-22.574 6.065-32.457c4.043-9.882 9.658-18.38 16.846-25.493c7.187-7.112 15.685-12.653 25.493-16.62c9.808-3.969 20.514-5.953 32.12-5.953h65.81c2.021 0 3.93.374 5.727 1.123a14.47 14.47 0 0 1 4.717 3.145a14.47 14.47 0 0 1 3.145 4.716c.748 1.797 1.123 3.706 1.123 5.728V210zM80.126 141.27h102.646V78.154h-51.323c-.898 0-2.789.131-5.671.393c-2.883.262-6.214.918-9.995 1.966c-3.781 1.048-7.768 2.658-11.96 4.829c-4.193 2.171-8.05 5.166-11.568 8.984c-3.519 3.819-6.42 8.591-8.704 14.319c-2.283 5.728-3.425 12.672-3.425 20.833z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresetAb(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M101 124.751c0-1.106.89-2.003 1.992-2.003h25.707L129 105l28 24.5l-28.301 24.5l.301-18.52h-26.009c-1.1 0-1.991-.895-1.991-2.003zM89.047 161H77.262v-15.955h-41.43V161H24.048v-32.5c0-4.744.816-9.11 2.448-13.1c1.632-3.989 3.898-7.418 6.8-10.29c2.9-2.87 6.33-5.106 10.289-6.708C47.542 96.801 51.864 96 56.547 96H83.11a5.95 5.95 0 0 1 2.312.453a5.84 5.84 0 0 1 3.173 3.173a5.95 5.95 0 0 1 .453 2.312V161zm-53.214-27.74h41.43v-25.475H56.546c-.362 0-1.125.053-2.289.159c-1.163.106-2.508.37-4.034.793c-1.526.423-3.135 1.073-4.827 1.95c-1.693.876-3.249 2.084-4.669 3.625c-1.42 1.542-2.591 3.468-3.513 5.78c-.922 2.311-1.382 5.114-1.382 8.408zm194.757 12.681c0 1.429-.246 2.96-.737 4.597c-.49 1.637-1.294 3.154-2.41 4.553c-1.115 1.398-2.566 2.566-4.351 3.503c-1.785.937-3.972 1.406-6.56 1.406H172.48a5.974 5.974 0 0 1-2.32-.446a5.673 5.673 0 0 1-1.853-1.228a5.673 5.673 0 0 1-1.227-1.852a5.974 5.974 0 0 1-.447-2.32v-52.307a5.935 5.935 0 0 1 1.674-4.15a5.504 5.504 0 0 1 1.852-1.25A5.974 5.974 0 0 1 172.48 96h38.828c1.429 0 2.96.253 4.597.759c1.637.506 3.162 1.316 4.575 2.432c1.413 1.116 2.589 2.566 3.526 4.352c.937 1.785 1.406 3.972 1.406 6.56v2.187c0 2.053-.343 4.262-1.027 6.628c-.684 2.365-1.785 4.604-3.303 6.717c1.31.803 2.537 1.77 3.682 2.9a18.572 18.572 0 0 1 3.013 3.928c.863 1.488 1.547 3.169 2.053 5.043c.506 1.875.759 3.943.759 6.204v2.231zm-11.604-2.231c0-1.518-.238-2.879-.714-4.084c-.476-1.205-1.146-2.239-2.009-3.102c-.862-.862-1.904-1.524-3.124-1.986c-1.22-.46-2.588-.691-4.106-.691H182.88v-11.694h20.932c1.517 0 2.886-.23 4.105-.691c1.22-.462 2.262-1.124 3.125-1.986c.862-.863 1.524-1.897 1.986-3.102c.46-1.205.691-2.566.691-4.084v-2.187c0-1.666-.803-2.5-2.41-2.5h-33.07v40.793h38.292c.208 0 .461-.015.759-.045c.297-.03.565-.119.803-.267c.238-.15.447-.402.625-.759c.179-.357.268-.848.268-1.473z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresetB(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M210.666 171.624c0 3.594-.618 7.45-1.853 11.567c-1.235 4.118-3.257 7.937-6.064 11.455c-2.808 3.52-6.458 6.458-10.95 8.816c-4.492 2.359-9.995 3.538-16.509 3.538H64.445c-2.096 0-4.043-.374-5.84-1.123c-1.796-.749-3.35-1.778-4.66-3.088s-2.34-2.864-3.089-4.661c-.748-1.797-1.123-3.744-1.123-5.84V60.667c0-2.022.375-3.93 1.123-5.728a14.934 14.934 0 0 1 3.089-4.716c1.31-1.348 2.864-2.396 4.66-3.145c1.797-.749 3.744-1.123 5.84-1.123h97.705c3.594 0 7.45.636 11.568 1.91c4.118 1.272 7.955 3.312 11.511 6.12c3.556 2.807 6.514 6.457 8.872 10.95c2.358 4.492 3.538 9.995 3.538 16.508v5.503c0 5.166-.861 10.725-2.583 16.678c-1.722 5.952-4.493 11.586-8.31 16.901c3.293 2.022 6.382 4.455 9.264 7.3c2.883 2.845 5.41 6.14 7.58 9.883c2.172 3.744 3.894 7.974 5.167 12.69c1.273 4.717 1.909 9.92 1.909 15.61zm-29.2-5.615c0-3.819-.598-7.244-1.796-10.276s-2.883-5.634-5.054-7.805c-2.171-2.172-4.791-3.837-7.861-4.998c-3.07-1.16-6.514-1.74-10.332-1.74h-65.81v-29.424h52.67c3.819 0 7.263-.58 10.332-1.741c3.07-1.16 5.69-2.826 7.862-4.998c2.17-2.17 3.837-4.773 4.997-7.805c1.16-3.032 1.74-6.457 1.74-10.276v-5.503c0-4.192-2.02-6.289-6.064-6.289H78.933v102.647h96.357c.524 0 1.16-.038 1.91-.113a4.627 4.627 0 0 0 2.02-.673c.6-.375 1.124-1.011 1.573-1.91c.45-.898.674-2.133.674-3.706v-5.39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresetBa(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M101 124.751c0-1.106.89-2.003 1.992-2.003h25.707L129 105l28 24.5l-28.301 24.5l.301-18.52h-26.009c-1.1 0-1.991-.895-1.991-2.003zM230.547 160h-11.604v-15.71h-40.792V160h-11.604v-32c0-4.671.804-8.97 2.41-12.898c1.607-3.928 3.839-7.305 6.695-10.131c2.856-2.827 6.233-5.029 10.131-6.606c3.898-1.577 8.153-2.365 12.764-2.365h26.154c.803 0 1.562.149 2.276.446a5.75 5.75 0 0 1 3.124 3.124c.298.715.446 1.473.446 2.277zm-52.396-27.314h40.792v-25.082h-20.396c-.357 0-1.108.052-2.253.156c-1.146.104-2.47.365-3.973.781c-1.502.417-3.087 1.056-4.753 1.92c-1.666.862-3.198 2.052-4.597 3.57c-1.398 1.517-2.551 3.414-3.459 5.69c-.907 2.276-1.36 5.036-1.36 8.279v4.686zM88.59 145.941c0 1.429-.246 2.96-.737 4.597c-.49 1.637-1.294 3.154-2.41 4.553c-1.115 1.398-2.566 2.566-4.351 3.503c-1.785.937-3.972 1.406-6.56 1.406H30.48a5.974 5.974 0 0 1-2.32-.446a5.673 5.673 0 0 1-1.853-1.228a5.673 5.673 0 0 1-1.227-1.852a5.974 5.974 0 0 1-.447-2.32v-52.307a5.935 5.935 0 0 1 1.674-4.15a5.504 5.504 0 0 1 1.852-1.25A5.974 5.974 0 0 1 30.481 96h38.828c1.429 0 2.96.253 4.597.759c1.637.506 3.162 1.316 4.575 2.432c1.413 1.116 2.589 2.566 3.526 4.352c.937 1.785 1.406 3.972 1.406 6.56v2.187c0 2.053-.343 4.262-1.027 6.628c-.684 2.365-1.785 4.604-3.303 6.717c1.31.803 2.537 1.77 3.682 2.9a18.572 18.572 0 0 1 3.013 3.928c.863 1.488 1.547 3.169 2.053 5.043c.506 1.875.759 3.943.759 6.204zm-11.604-2.231c0-1.518-.238-2.879-.714-4.084c-.476-1.205-1.146-2.239-2.009-3.102c-.862-.862-1.904-1.524-3.124-1.986c-1.22-.46-2.588-.691-4.106-.691H40.88v-11.694h20.932c1.517 0 2.886-.23 4.105-.691c1.22-.462 2.262-1.124 3.125-1.986c.862-.863 1.524-1.897 1.986-3.102c.46-1.205.691-2.566.691-4.084v-2.187c0-1.666-.803-2.5-2.41-2.5H36.24v40.793h38.29c.208 0 .461-.015.759-.045c.297-.03.565-.119.803-.267c.238-.15.447-.402.625-.759c.179-.357.268-.848.268-1.473v-2.142z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Prev(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M175.967 60.858c0 14.184.118 120.734.118 134.66c0 13.925-7.798 16.307-14.543 10.568c-6.745-5.74-46.666-60.003-52.015-67.201c-5.348-7.198-4.449-12.951.086-20.28c4.535-7.328 47.284-62.224 52.46-68.2c5.176-5.977 13.894-3.73 13.894 10.453m-51.265 71.792l32.874 43.167c1.336 1.754 2.424 1.393 2.424-.814l-.005-92.006c0-2.197-1.073-2.562-2.396-.792l-32.926 44.057c-1.328 1.776-1.31 4.63.029 6.388M79.836 53.792c-.17-2.538 1.22-4.792 3.45-4.792h8.193c3.655 0 3.52.682 3.52 3.859c0 3.176-.394 147.132-.46 149.86c-.064 2.729-1.06 4.454-3.219 4.454s-6.252 0-8.33-.006c-2.079-.007-3.088-1.897-3.088-4.033s.104-146.803-.066-149.342"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PunchIn(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M24.28 67.494a3.999 3.999 0 0 1 4.005-3.994H83.78c8.837 0 19.659 6.169 24.165 13.767l50.391 84.966c4.51 7.604 15.33 13.767 24.157 13.767H227.7a4.002 4.002 0 0 1 4.005 4.007v8.986a4.004 4.004 0 0 1-3.998 4.007h-58.708c-8.835 0-19.602-6.196-24.04-13.828L95.32 93.828C90.88 86.19 80.114 80 71.28 80h-43c-2.21 0-4-1.797-4-3.994z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PunchOut(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M233.172 67.494a3.999 3.999 0 0 0-4.006-3.994h-55.494c-8.837 0-18.43 6.728-21.436 15.054l-54.837 82.392C94.397 169.26 84.8 176 75.971 176H30.765a4.002 4.002 0 0 0-4.006 4.007v8.986A4.004 4.004 0 0 0 30.757 193h58.708c8.835 0 18.381-6.752 21.324-15.089l54.057-82.822C167.787 86.756 177.337 80 186.171 80h43c2.21 0 4-1.797 4-3.994v-8.512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ram(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M223.977 63.69a5.967 5.967 0 0 0-2.295-.446l.001-.001l-188.877.77a6.013 6.013 0 0 0-5.984 6.024l.226 98.86c.008 3.314 2.699 6.015 6.01 6.034l15.326.087v10.633c0 3.315 2.683 6.007 5.993 6.007h25.471s6.319-15.38 23.62-15.38c17.301 0 23.783 15.38 23.783 15.38h74.032a6.003 6.003 0 0 0 5.999-6.008v-8.16h14.388a6.004 6.004 0 0 0 6.01-6.003V69.222a5.976 5.976 0 0 0-1.755-4.237a5.978 5.978 0 0 0-1.948-1.294zM42.588 80.586A2.002 2.002 0 0 1 44.002 80v.001h44A1.995 1.995 0 0 1 90 82.002l-.13 69.985s-4.127.137-7.667 2.48S77.76 160 77.76 160H42V81.997a2 2 0 0 1 .588-1.412zm62.001-.001a2.001 2.001 0 0 1 1.413-.584h44a1.999 1.999 0 0 1 1.847 1.232c.1.243.151.503.151.765l-.141 76.006a2.01 2.01 0 0 1-2.013 1.997h-20.189s-4.827-4.457-10.622-6.472c-5.795-2.014-15.035-2.447-15.035-2.447V81.994a2 2 0 0 1 .589-1.41zm61.999 0a2.002 2.002 0 0 1 1.414-.584h44a1.999 1.999 0 0 1 1.847 1.232c.1.243.151.503.151.765l-.141 76.006A2.004 2.004 0 0 1 211.86 160H166V81.997a2 2 0 0 1 .588-1.412" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RandomOneDice(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M47.895 88.097c.001-4.416 3.064-9.837 6.854-12.117l66.257-39.858c3.785-2.277 9.915-2.277 13.707.008l66.28 39.934c3.786 2.28 6.853 7.703 6.852 12.138l-.028 79.603c-.001 4.423-3.069 9.865-6.848 12.154l-66.4 40.205c-3.781 2.29-9.903 2.289-13.69-.01l-66.167-40.185c-3.78-2.295-6.842-7.733-6.84-12.151zm13.936-6.474l65.834 36.759l62.766-36.278l-62.872-36.918L61.83 81.623zM57.585 93.52c0 1.628-1.065 71.86-1.065 71.86c-.034 2.206 1.467 4.917 3.367 6.064l61.612 37.182l.567-77.413s-64.48-39.322-64.48-37.693zm76.107 114.938l60.912-38.66c2.332-1.48 4.223-4.915 4.223-7.679V93.125l-65.135 37.513z"/><path d="M77.76 132.287c-4.782 2.762-11.122.735-14.16-4.526c-3.037-5.261-1.622-11.765 3.16-14.526c4.783-2.762 11.123-.735 14.16 4.526c3.038 5.261 1.623 11.765-3.16 14.526m32 21c-4.782 2.762-11.122.735-14.16-4.526c-3.037-5.261-1.622-11.765 3.16-14.526c4.783-2.762 11.123-.735 14.16 4.526c3.038 5.261 1.623 11.765-3.16 14.526m-32 16c-4.782 2.762-11.122.735-14.16-4.526c-3.037-5.261-1.622-11.765 3.16-14.526c4.783-2.762 11.123-.735 14.16 4.526c3.038 5.261 1.623 11.765-3.16 14.526m32 21c-4.782 2.762-11.122.735-14.16-4.526c-3.037-5.261-1.622-11.765 3.16-14.526c4.783-2.762 11.123-.735 14.16 4.526c3.038 5.261 1.623 11.765-3.16 14.526m78.238-78.052c-4.783-2.762-11.122-.735-14.16 4.526c-3.037 5.261-1.623 11.765 3.16 14.526c4.783 2.762 11.123.735 14.16-4.526c3.038-5.261 1.623-11.765-3.16-14.526m-16.238 29c-4.782-2.762-11.122-.735-14.16 4.526c-3.037 5.261-1.622 11.765 3.16 14.526c4.783 2.762 11.123.735 14.16-4.526c3.038-5.261 1.623-11.765-3.16-14.526m-17 28c-4.782-2.762-11.122-.735-14.16 4.526c-3.037 5.261-1.622 11.765 3.16 14.526c4.783 2.762 11.123.735 14.16-4.526c3.038-5.261 1.623-11.765-3.16-14.526M128.5 69c-6.351 0-11.5 4.925-11.5 11s5.149 11 11.5 11S140 86.075 140 80s-5.149-11-11.5-11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RandomTwoDice(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M24.898 100.907a7.97 7.97 0 0 1 8.035-7.935l80.011.623c4.419.034 8.209 3.635 8.466 8.042l.517 8.868l26.68-42.392a7.776 7.776 0 0 1 10.94-2.349l66.996 44.369a8.03 8.03 0 0 1 2.275 11.113l-43.766 66.506c-2.432 3.695-7.447 4.8-11.197 2.47l-51.928-32.265v26.49c0 4.419-3.583 8-7.993 8H32.498a7.949 7.949 0 0 1-7.959-7.998zm11.828 6.694l-.189 71.811l74.127.073l-.035-29.78l-5.954-4.119c-1.809-1.25-2.375-3.81-1.257-5.71L111 127l-.466-19.749zM156.483 79L118 138.79l60.965 38.32l37.612-58.539z"/><circle cx="138" cy="135" r="8"/><circle cx="165" cy="130" r="8"/><circle cx="193" cy="125" r="8"/><circle cx="50" cy="124" r="8"/><circle cx="73" cy="145" r="8"/><circle cx="95" cy="123" r="8"/><circle cx="51" cy="165" r="8"/><circle cx="95" cy="165" r="8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M128 209c-44.735 0-81-36.265-81-81s36.265-81 81-81s81 36.265 81 81s-36.265 81-81 81m.5-33c26.51 0 48-21.49 48-48s-21.49-48-48-48s-48 21.49-48 48s21.49 48 48 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M202.238 167.072c.974-1.973 3.388-2.796 5.372-1.847l7.893 3.775s-22.5 53.5-85.5 56c-60-1.5-96.627-48.626-97-96.5c-.373-47.874 37-95.5 95.5-96c57.5-1 79.555 45.004 79.555 45.004c1.074 1.93 1.945 1.698 1.945-.501V51.997a4 4 0 0 1 4-3.997h6.5c2.209 0 4 1.8 4 4.008v48.984a3.998 3.998 0 0 1-3.998 4.008H170a3.995 3.995 0 0 1-3.998-3.993v-6.014c0-2.205 1.789-4.02 4.007-4.053l25.485-.38c2.213-.033 3.223-1.679 2.182-3.628c0 0-18.174-41.932-68.674-41.432c-49 .5-82.751 41.929-82.5 83.242c3 55.258 45 82.258 83.5 81.258c54.5 0 72.235-42.928 72.235-42.928z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M109.533 197.602a1.887 1.887 0 0 1-.034 2.76l-7.583 7.066a4.095 4.095 0 0 1-5.714-.152l-32.918-34.095c-1.537-1.592-1.54-4.162-.002-5.746l33.1-34.092c1.536-1.581 4.11-1.658 5.74-.18l7.655 6.94c.82.743.833 1.952.02 2.708l-21.11 19.659s53.036.129 71.708.064c18.672-.064 33.437-16.973 33.437-34.7c0-7.214-5.578-17.64-5.578-17.64c-.498-.99-.273-2.444.483-3.229l8.61-8.94c.764-.794 1.772-.632 2.242.364c0 0 9.212 18.651 9.212 28.562c0 28.035-21.765 50.882-48.533 50.882c-26.769 0-70.921.201-70.921.201z"/><path d="M144.398 58.435a1.887 1.887 0 0 1 .034-2.76l7.583-7.066a4.095 4.095 0 0 1 5.714.152l32.918 34.095c1.537 1.592 1.54 4.162.002 5.746l-33.1 34.092c-1.536 1.581-4.11 1.658-5.74.18l-7.656-6.94c-.819-.743-.832-1.952-.02-2.708l21.111-19.659s-53.036-.129-71.708-.064c-18.672.064-33.437 16.973-33.437 34.7c0 7.214 5.578 17.64 5.578 17.64c.498.99.273 2.444-.483 3.229l-8.61 8.94c-.764.794-1.772.632-2.242-.364c0 0-9.212-18.65-9.212-28.562c0-28.035 21.765-50.882 48.533-50.882c26.769 0 70.921-.201 70.921-.201z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RepeatOne(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M109.533 197.602a1.887 1.887 0 0 1-.034 2.76l-7.583 7.066a4.095 4.095 0 0 1-5.714-.152l-32.918-34.095c-1.537-1.592-1.54-4.162-.002-5.746l33.1-34.092c1.536-1.581 4.11-1.658 5.74-.18l7.655 6.94c.82.743.833 1.952.02 2.708l-21.11 19.659s53.036.129 71.708.064c18.672-.064 33.437-16.973 33.437-34.7c0-7.214-5.578-17.64-5.578-17.64c-.498-.99-.273-2.444.483-3.229l8.61-8.94c.764-.794 1.772-.632 2.242.364c0 0 9.212 18.651 9.212 28.562c0 28.035-21.765 50.882-48.533 50.882c-26.769 0-70.921.201-70.921.201z"/><path d="M144.398 58.435a1.887 1.887 0 0 1 .034-2.76l7.583-7.066a4.095 4.095 0 0 1 5.714.152l32.918 34.095c1.537 1.592 1.54 4.162.002 5.746l-33.1 34.092c-1.536 1.581-4.11 1.658-5.74.18l-7.656-6.94c-.819-.743-.832-1.952-.02-2.708l21.111-19.659s-53.036-.129-71.708-.064c-18.672.064-33.437 16.973-33.437 34.7c0 7.214 5.578 17.64 5.578 17.64c.498.99.273 2.444-.483 3.229l-8.61 8.94c-.764.794-1.772.632-2.242-.364c0 0-9.212-18.65-9.212-28.562c0-28.035 21.765-50.882 48.533-50.882c26.769 0 70.921-.201 70.921-.201z"/><path d="m127.992 104.543l6.53.146c1.105.025 2.013.945 2.027 2.037l.398 30.313a1.97 1.97 0 0 0 2.032 1.94l4.104-.103a1.951 1.951 0 0 1 2.01 1.958l.01 4.838a2.015 2.015 0 0 1-1.99 2.024l-21.14.147a1.982 1.982 0 0 1-1.994-1.983l-.002-4.71c0-1.103.897-1.997 1.996-1.997h4.254a2.018 2.018 0 0 0 2.016-1.994l.169-16.966l-6.047 5.912l-6.118-7.501z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rew(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M208.433 63.148c0 11.961-.219 119.755-.219 131.209c0 11.454-11.84 14.856-16.027 10.347c-4.187-4.508-52.355-60.404-52.355-60.404c-.719-.834-1.293-.615-1.282.488c0 0 .37 37.702.37 49.528c0 11.827-12.464 15.24-16.516 10.433c-4.051-4.806-56.586-65.858-56.586-65.858c-.72-.837-1.316-.624-1.327.48c0 0-.617 62.39-.617 64.527c0 2.136-1.009 4.026-3.087 4.032c-2.078.007-6.172.007-8.33.007c-2.16 0-3.155-1.726-3.22-4.454c-.065-2.728-.46-146.684-.46-149.86c0-3.177-.134-3.86 3.521-3.86h8.192c2.23 0 3.62 2.255 3.45 4.793c-.17 2.538.578 62.396.578 62.396c.013 1.092.614 1.323 1.341.5c0 0 50.693-57.393 56.836-65.488c6.142-8.096 16.073-1.062 16.073 10.137c0 11.2.085 50.228.085 50.228c.002 1.103.594 1.333 1.322.504c0 0 46.313-52.757 52.713-60.673c6.4-7.915 15.545-.974 15.545 10.988M193.584 78.97c.001-1.103-.59-1.322-1.315-.494l-42.179 48.115c-.727.83-.735 2.182-.015 3.024l42.11 49.234c.72.84 1.303.628 1.304-.474zm-71.046-1.025l-42.24 49.323c-.722.843-.716 2.203.013 3.04l42.64 48.985c.729.838 1.314.623 1.31-.477l-.41-100.404c-.004-1.09-.592-1.309-1.313-.467"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundswitchOff(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M231 128c0 8.667-1.71 16.95-5.13 24.85c-3.42 7.9-7.996 14.7-13.73 20.4c-5.733 5.7-12.573 10.25-20.52 13.65c-7.946 3.4-16.278 5.1-24.995 5.1h-77.25c-8.717 0-17.05-1.7-24.996-5.1c-7.946-3.4-14.786-7.95-20.52-13.65c-5.733-5.7-10.31-12.5-13.73-20.4C26.71 144.95 25 136.667 25 128c0-8.667 1.71-16.95 5.13-24.85c3.42-7.9 7.996-14.7 13.73-20.4c5.733-5.7 12.573-10.25 20.52-13.65c7.946-3.4 16.278-5.1 24.995-5.1h77.25c8.717 0 17.05 1.7 24.996 5.1c7.946 3.4 14.786 7.95 20.52 13.65c5.733 5.7 10.31 12.5 13.73 20.4C229.29 111.05 231 119.333 231 128m-18 .064c0-6.508-1.36-12.72-4.08-18.634c-2.72-5.914-6.397-11.03-11.032-15.349c-4.634-4.318-10.125-7.744-16.472-10.279C175.07 81.267 168.404 80 161.42 80h-31.47c5.44 3.802 16.472 16.147 19.947 23.281c3.476 7.135 5.642 17.96 5.642 24.468c0 .783.022 1.578.044 2.398c.166 5.999.367 13.327-7.701 26.831c-5.639 9.439-13.944 16.357-17.909 19.15h-.628c-2.637 1.88-1.855 1.748.628 0h31.447c6.985 0 13.65-1.267 19.997-3.801c6.347-2.535 11.838-5.962 16.472-10.28c4.635-4.318 8.312-9.434 11.032-15.348a44.085 44.085 0 0 0 4.08-18.635zm-78.797-18.674C136.734 115.297 138 121.5 138 128s-1.266 12.703-3.797 18.609c-2.531 5.907-5.953 11.016-10.265 15.328c-4.313 4.313-9.422 7.735-15.328 10.266C102.703 174.734 96.5 176 90 176s-12.703-1.266-18.609-3.797c-5.907-2.531-11.016-5.953-15.328-10.266c-4.313-4.312-7.735-9.421-10.266-15.327C43.266 140.703 42 134.5 42 128s1.266-12.703 3.797-18.609c2.531-5.907 5.953-11.016 10.265-15.328c4.313-4.313 9.422-7.735 15.328-10.266C77.297 81.266 83.5 80 90 80s12.703 1.266 18.609 3.797c5.907 2.531 11.016 5.953 15.328 10.265c4.313 4.313 7.735 9.422 10.266 15.328" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundswitchOn(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25 128c0-8.667 1.71-16.95 5.13-24.85c3.42-7.9 7.996-14.7 13.73-20.4c5.733-5.7 12.573-10.25 20.52-13.65c7.946-3.4 16.278-5.1 24.995-5.1h77.25c8.717 0 17.05 1.7 24.996 5.1c7.946 3.4 14.786 7.95 20.52 13.65c5.733 5.7 10.31 12.5 13.73 20.4C229.29 111.05 231 119.333 231 128c0 8.667-1.71 16.95-5.13 24.85c-3.42 7.9-7.996 14.7-13.73 20.4c-5.733 5.7-12.573 10.25-20.52 13.65c-7.946 3.4-16.278 5.1-24.995 5.1h-77.25c-8.717 0-17.05-1.7-24.996-5.1c-7.946-3.4-14.786-7.95-20.52-13.65c-5.733-5.7-10.31-12.5-13.73-20.4C26.71 144.95 25 136.667 25 128m18-.064c0 6.508 1.36 12.72 4.08 18.634c2.72 5.914 6.397 11.03 11.032 15.349c4.634 4.318 10.125 7.744 16.472 10.279C80.93 174.733 87.596 176 94.58 176h31.47c-5.44-3.802-16.472-16.147-19.947-23.281c-3.476-7.135-5.642-17.96-5.642-24.468c0-.782-.022-1.577-.044-2.397v-.001c-.166-5.999-.367-13.327 7.701-26.831c5.639-9.44 13.944-16.357 17.909-19.15h.628c2.637-1.88 1.855-1.748-.628 0H94.58c-6.985 0-13.65 1.267-19.997 3.801c-6.347 2.535-11.838 5.962-16.472 10.28c-4.635 4.318-8.312 9.434-11.032 15.348A44.088 44.088 0 0 0 43 127.936m78.797 18.674C119.266 140.703 118 134.5 118 128s1.266-12.703 3.797-18.609c2.531-5.907 5.953-11.016 10.266-15.328c4.312-4.313 9.421-7.735 15.327-10.266C153.297 81.266 159.5 80 166 80s12.703 1.266 18.609 3.797c5.907 2.531 11.016 5.953 15.328 10.265c4.313 4.313 7.735 9.422 10.266 15.328C212.734 115.297 214 121.5 214 128s-1.266 12.703-3.797 18.609c-2.531 5.907-5.953 11.016-10.266 15.328c-4.312 4.313-9.421 7.735-15.327 10.266C178.703 174.734 172.5 176 166 176s-12.703-1.266-18.609-3.797c-5.907-2.531-11.016-5.953-15.328-10.266c-4.313-4.312-7.735-9.421-10.266-15.327" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M65.456 48.385c10.02 0 96.169-.355 96.169-.355c2.209-.009 5.593.749 7.563 1.693c0 0-1.283-1.379.517.485c1.613 1.67 35.572 36.71 36.236 37.416c.665.707.241.332.241.332c.924 2.007 1.539 5.48 1.539 7.691v95.612c0 7.083-8.478 16.618-16.575 16.618c-8.098 0-118.535-.331-126.622-.331c-8.087 0-16-6.27-16.356-16.1c-.356-9.832.356-118.263.356-126.8c0-8.536 6.912-16.261 16.932-16.261m-1.838 17.853l.15 121c.003 2.198 1.8 4.003 4.012 4.015l120.562.638a3.971 3.971 0 0 0 4-3.981l-.143-90.364c-.001-1.098-.649-2.616-1.445-3.388L161.52 65.841c-.801-.776-1.443-.503-1.443.601v35.142c0 3.339-4.635 9.14-8.833 9.14H90.846c-4.6 0-9.56-4.714-9.56-9.14s-.014-35.14-.014-35.14c0-1.104-.892-2.01-1.992-2.023l-13.674-.155a1.968 1.968 0 0 0-1.988 1.972m32.542.44v27.805c0 1.1.896 2.001 2 2.001h44.701c1.113 0 2-.896 2-2.001V66.679a2.004 2.004 0 0 0-2-2.002h-44.7c-1.114 0-2 .896-2 2.002z"/><path d="M127.802 119.893c16.176.255 31.833 14.428 31.833 31.728s-14.615 31.782-31.016 31.524c-16.401-.259-32.728-14.764-32.728-31.544s15.735-31.963 31.91-31.708zm-16.158 31.31c0 9.676 7.685 16.882 16.218 16.843c8.534-.039 15.769-7.128 15.812-16.69c.043-9.563-7.708-16.351-15.985-16.351c-8.276 0-16.045 6.52-16.045 16.197z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Saveas(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M93.598 47.311c8.113 0 77.093-.284 77.093-.284c2.206-.009 5.48.779 7.306 1.76c0 0-1.457-1.508 0 0c1.306 1.352 28.798 29.707 29.336 30.28c.539.571 0 0 0 0c.817 1.577 1.44 5.043 1.44 7.253v76.614c0 5.731-6.862 13.448-13.418 13.448h-8.596v25.038c0 4.737-1.55 8.014-8.762 8.014H57.102c-7.796 0-10.13-3.07-10.13-8.666V82.318c0-7.65 5.155-12.205 9.15-12.205h23.726s.043-8.475.043-9.642c0-6.908 5.595-13.16 13.707-13.16M77.704 84.577H67.88a3.992 3.992 0 0 0-3.993 3.996v99.795a3.957 3.957 0 0 0 3.996 3.969l103.782-.727v-15.182s-79.998.36-85.653.36c-5.655 0-5.655-6.424-5.655-6.424l-.633-83.793a2.026 2.026 0 0 0-2.019-1.994zm17.463-19.472v92.24a4.006 4.006 0 0 0 4 4.008l90.162.105a3.987 3.987 0 0 0 3.993-4l-.113-69.98c-.002-1.1-.62-2.641-1.38-3.44L171.58 62.755s-.379.448-.379 1.554v26.053c0 2.702-4.752 7.397-8.15 7.397h-48.898c-3.725 0-7.74-3.815-7.74-7.397V64.941a1.966 1.966 0 0 0-1.998-1.964l-7.252.095a2.04 2.04 0 0 0-1.997 2.033zm23.289-2.61v21.743c0 1.112.9 1.998 2.01 1.998h35.408c1.112 0 2.01-.894 2.01-1.998V62.494c0-1.112-.9-1.998-2.01-1.998h-35.409a2.003 2.003 0 0 0-2.009 1.998z"/><path d="M144.073 105.18c13.095.206 25.771 11.676 25.771 25.676c0 14-11.831 25.72-25.11 25.511c-13.278-.21-26.496-11.948-26.496-25.527c0-13.58 12.739-25.867 25.835-25.66m-13.082 25.337c0 7.831 6.222 13.663 13.13 13.631c6.909-.031 12.767-5.768 12.801-13.507c.035-7.738-6.24-13.231-12.94-13.231c-6.7 0-12.99 5.276-12.99 13.107z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M109.56 157.868a33.419 33.419 0 0 1 5.59 18.557c0 18.543-15.032 33.575-33.575 33.575C63.032 210 48 194.968 48 176.425c0-18.543 15.032-33.575 33.575-33.575a33.42 33.42 0 0 1 16.753 4.472l19.743-19.843l-17.873-17.963a33.419 33.419 0 0 1-18.623 5.634C63.032 115.15 48 100.118 48 81.575C48 63.032 63.032 48 81.575 48c18.543 0 33.575 15.032 33.575 33.575a33.41 33.41 0 0 1-4.43 16.68l18.309 18.21L198.144 47h9.974l-33.82 61.459c-1.062 1.93-3.714 3.39-5.918 3.249l-6.667-.426c-2.753-.176-6.585 1.261-8.545 3.21l-13.06 12.992l14.06 13.986c1.96 1.95 5.792 3.387 8.545 3.21l6.667-.425c2.204-.14 4.856 1.319 5.918 3.25l33.82 61.458h-9.974l-70.11-70.465zM81.602 195c9.72 0 17.601-7.89 17.601-17.623s-7.88-17.623-17.601-17.623c-9.72 0-17.601 7.89-17.601 17.623S71.88 195 81.601 195zm0-95c9.72 0 17.601-7.89 17.601-17.623s-7.88-17.623-17.601-17.623c-9.72 0-17.601 7.89-17.601 17.623S71.88 100 81.601 100z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m47.813 71.467l36.415-.012c2.211 0 5.129 1.393 6.52 3.118l22.522 27.923l-10.972 11.499c-.763.8-1.929.738-2.599-.131L82.43 91.477c-.672-.872-2.112-1.577-3.21-1.575l-29.479.055a1.98 1.98 0 0 1-1.986-1.993l.059-16.497z"/><path d="M49.2 166.329a1.904 1.904 0 0 0-1.958 1.963l.303 14.054a2.092 2.092 0 0 0 2.042 2.034l34.042.55c2.212.035 5.14-1.32 6.544-3.03l74.103-90.29h15.474l-22.104 26.93h19.1c1.66 0 3.902-.988 5.026-2.226l23.499-25.87c4.086-4.497 3.968-11.673-.268-16.033l-22.46-23.115c-1.543-1.587-4.583-2.874-6.796-2.874H160.65c-1.659 0-2.096.977-.964 2.195l20.064 21.587h-19.106c-2.208 0-5.129 1.394-6.523 3.111L81.023 165.38c-.693.854-2.16 1.528-3.255 1.507L49.2 166.33z"/><path d="M140.208 139.822c1.451-1.663 3.72-1.586 5.067.167l19.001 24.754h15.474l-19.973-26.931h16.983c1.651 0 3.877 1.007 4.975 2.253l22.661 25.729c4.017 4.56 3.902 11.835-.271 16.265l-21.626 22.953c-1.518 1.612-4.539 2.918-6.743 2.918h-12.989c-1.65 0-2.144-1.049-1.11-2.331l18.093-22.451h-20.096c-2.213 0-5.153-1.389-6.547-3.08l-22.06-26.75c-.704-.853-.68-2.226.038-3.048z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderRoundOne(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path fill-rule="nonzero" d="M128 226c-13.27 0-25.936-2.584-37.994-7.752c-12.059-5.168-22.49-12.154-31.295-20.959c-8.805-8.805-15.791-19.236-20.959-31.295C32.584 153.936 30 141.271 30 128c0-13.27 2.584-25.936 7.752-37.994c5.168-12.059 12.154-22.49 20.959-31.295c8.805-8.805 19.236-15.791 31.295-20.959C102.064 32.584 114.729 30 128 30c13.27 0 25.936 2.584 37.994 7.752c12.059 5.168 22.49 12.154 31.295 20.959c8.805 8.805 15.791 19.236 20.959 31.295C223.416 102.064 226 114.729 226 128c0 13.27-2.584 25.936-7.752 37.994c-5.168 12.059-12.154 22.49-20.959 31.295c-8.805 8.805-19.236 15.791-31.295 20.959C153.936 223.416 141.271 226 128 226m0-16c11.104 0 21.701-2.162 31.791-6.486c10.09-4.325 18.818-10.17 26.186-17.537c7.367-7.368 13.212-16.096 17.537-26.186C207.838 149.701 210 139.104 210 128s-2.162-21.701-6.486-31.791c-4.325-10.09-10.17-18.818-17.537-26.186c-7.368-7.367-16.096-13.212-26.186-17.537C149.701 48.162 139.104 46 128 46s-21.701 2.162-31.791 6.486c-10.09 4.325-18.818 10.17-26.186 17.537c-7.367 7.368-13.212 16.096-17.537 26.186C48.162 106.299 46 116.896 46 128s2.162 21.701 6.486 31.791c4.325 10.09 10.17 18.818 17.537 26.186c7.368 7.367 16.096 13.212 26.186 17.537C106.299 207.838 116.896 210 128 210"/><path d="M128 96c9.941 0 18-8.059 18-18s-8.059-18-18-18a17.932 17.932 0 0 0-11.945 4.534A17.956 17.956 0 0 0 110 78c0 9.941 8.059 18 18 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderRoundThree(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path fill-rule="nonzero" d="M128 226c-13.27 0-25.936-2.584-37.994-7.752c-12.059-5.168-22.49-12.154-31.295-20.959c-8.805-8.805-15.791-19.236-20.959-31.295C32.584 153.936 30 141.271 30 128c0-13.27 2.584-25.936 7.752-37.994c5.168-12.059 12.154-22.49 20.959-31.295c8.805-8.805 19.236-15.791 31.295-20.959C102.064 32.584 114.729 30 128 30c13.27 0 25.936 2.584 37.994 7.752c12.059 5.168 22.49 12.154 31.295 20.959c8.805 8.805 15.791 19.236 20.959 31.295C223.416 102.064 226 114.729 226 128c0 13.27-2.584 25.936-7.752 37.994c-5.168 12.059-12.154 22.49-20.959 31.295c-8.805 8.805-19.236 15.791-31.295 20.959C153.936 223.416 141.271 226 128 226m-.247-7.742c50.034 0 90.594-40.634 90.594-90.758c0-50.124-40.56-90.758-90.594-90.758c-50.034 0-90.594 40.634-90.594 90.758c0 50.124 40.56 90.758 90.594 90.758"/><path fill-rule="nonzero" d="M202.302 131.704c-.247 4.699-.74 9.645-1.975 14.343c-1.234 4.699-2.715 9.398-4.69 13.602c-6.171 2.225-11.602 5.935-16.045 11.375c-4.444 5.44-6.912 11.376-7.9 17.806c-7.652 5.688-16.538 9.892-25.919 12.117c-5.43-3.462-11.849-5.44-18.76-5.44c-6.912 0-13.33 1.73-18.761 4.946c-9.38-2.72-18.02-7.172-25.673-12.86c-.987-6.182-3.455-12.612-7.652-18.052c-4.196-5.688-9.627-9.398-15.551-11.87c-3.703-8.656-5.678-18.3-5.925-27.945c4.69-4.699 8.146-10.387 9.627-17.064c1.481-6.677.988-13.354-.74-19.536c4.69-8.656 10.861-16.322 18.267-22.257c6.418.742 13.083-.247 19.254-3.215c6.171-2.968 11.355-7.172 14.811-12.612c9.134-1.731 19.008-1.484 28.635.494c3.456 5.194 8.393 9.645 14.564 12.86c5.924 2.968 12.59 4.204 19.007 3.71c7.406 6.43 13.33 14.095 17.527 22.75c-1.975 6.183-2.222 12.86-.988 19.537c1.235 6.677 4.444 12.612 8.887 17.311m-25.278 7.941c6.625-27.292-10.17-54.798-37.51-61.435c-27.343-6.637-54.878 10.108-61.503 37.4c-6.625 27.292 10.169 54.798 37.51 61.435c27.342 6.637 54.878-10.108 61.503-37.4"/><ellipse cx="128.5" cy="96" rx="14.5" ry="14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderRoundTwo(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 226c-13.27 0-25.936-2.584-37.994-7.752c-12.059-5.168-22.49-12.154-31.295-20.959c-8.805-8.805-15.791-19.236-20.959-31.295C32.584 153.936 30 141.271 30 128c0-13.27 2.584-25.936 7.752-37.994c5.168-12.059 12.154-22.49 20.959-31.295c8.805-8.805 19.236-15.791 31.295-20.959C102.064 32.584 114.729 30 128 30c13.27 0 25.936 2.584 37.994 7.752c12.059 5.168 22.49 12.154 31.295 20.959c8.805 8.805 15.791 19.236 20.959 31.295C223.416 102.064 226 114.729 226 128c0 13.27-2.584 25.936-7.752 37.994c-5.168 12.059-12.154 22.49-20.959 31.295c-8.805 8.805-19.236 15.791-31.295 20.959C153.936 223.416 141.271 226 128 226m0-16c11.104 0 21.701-2.162 31.791-6.486c10.09-4.325 18.818-10.17 26.186-17.537c7.367-7.368 13.212-16.096 17.537-26.186C207.838 149.701 210 139.104 210 128s-2.162-21.701-6.486-31.791c-4.325-10.09-10.17-18.818-17.537-26.186c-7.368-7.367-16.096-13.212-26.186-17.537c-7.694-3.297-11.507-5.337-19.791-6.12v71.125a3.998 3.998 0 0 1-3.99 4.009h-16.02a3.995 3.995 0 0 1-3.99-4.008v-71.13c-8.2 0-12.082 2.82-19.791 6.124c-10.09 4.325-18.818 10.17-26.186 17.537c-7.367 7.368-13.212 16.096-17.537 26.186C48.162 106.299 46 116.896 46 128s2.162 21.701 6.486 31.791c4.325 10.09 10.17 18.818 17.537 26.186c7.368 7.367 16.096 13.212 26.186 17.537C106.299 207.838 116.896 210 128 210"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderhandleOne(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M64.254 32.888C63.56 28.532 66.584 25 70.996 25h113.008c4.416 0 7.477 3.545 6.875 7.913c0 0-6.879 43.782-6.879 95.218c0 51.435 6.875 94.961 6.875 94.961c.621 4.368-2.46 7.908-6.871 7.908H70.996c-4.416 0-7.402-3.537-6.714-7.899c0 0 7.718-42.667 7.718-94.474c0-51.808-7.746-95.739-7.746-95.739M80 40s8 44.261 8 88.521S80 217 80 217h96s-6.533-44.446-6.533-88.479C169.467 84.49 176 40 176 40z"/><path d="M95 57h65l-4 16H98zm5 125h55l5 19H96zm4-58h48v9h-48z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderhandleTwo(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M55.17 48.542h144.553a8 8 0 0 1 8.004 8.004v143.736a8.001 8.001 0 0 1-8.004 8.004H55.17a8 8 0 0 1-8.003-8.004V56.546a8.001 8.001 0 0 1 8.003-8.004m8.618 143.857l17.54-8.168h99.462a3.983 3.983 0 0 0 3.976-4l-.393-99.693l7.901-16.31H63.788z"/><path d="M92.995 112A3.995 3.995 0 0 0 89 115.996v8.008a3.974 3.974 0 0 0 3.998 3.976l71.534-.365c2.208-.011 4.052-1.81 4.118-4.016l.23-7.603a3.85 3.85 0 0 0-3.875-3.996z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Softclip(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m26.572 175.798l-.423 16.215s6.07-.61 16.657-3.59c10.586-2.978 10.222-.383 20.515-9.87c10.293-9.486 9.362-6.84 15.235-20.838c5.873-13.998 5.28-12.871 7.91-25.156c2.63-12.286 2.907-16.198 7.088-26.018c4.18-9.82 4.448-10.736 10.46-16.25c6.014-5.514 9.596-6.742 16.67-8.629c7.074-1.887 8.357-2.026 16.579.3c8.221 2.327 6.33.988 13.914 8.329c7.583 7.34 6.95 6.504 11.476 16.25c4.527 9.746 3.555 10.641 6.431 20.151s3.3 11.71 5.393 18.947c2.091 7.236 1.493 7.698 4.265 14.464c2.772 6.765 2.963 8.233 7.104 13.287c4.14 5.053 3.365 5.15 11.273 9.14c7.908 3.99 11.118 5.02 18.61 6.939c7.493 1.92 15.641 1.836 15.641 1.836l.339-16.327s-11.697.108-20.815-3.749c-9.12-3.856-8.863-4.217-15.464-14.17c-6.602-9.953-2.691-5.123-7.283-19.616c-4.59-14.492-4.386-16.372-8.116-27.993c-3.73-11.62-3.692-14.478-10.296-23.138c-6.603-8.66-9.682-10.42-18.784-15.001c-9.102-4.582-13.384-5.268-22.453-5.756c-9.069-.488-8.27-.513-18.477 3.552c-10.207 4.066-9.839 3.645-17.729 10.868c-7.89 7.222-7.031 6.35-11.785 16.148c-4.754 9.798-4.66 9.053-6.893 20.284c-2.232 11.232-2.47 12.362-5.41 22.469c-2.942 10.107-3.318 13.965-6.975 20.092c-3.656 6.126-2.394 6.64-11.667 11.143c-9.273 4.503-22.99 5.687-22.99 5.687"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Softclipcurve(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M233 64.5h-28.495c-18.104 0-32.517 4.04-49.695 18.089c-15.765 12.892-30.941 31.655-39.559 46.948c-12.478 22.144-33.858 39.953-43.54 43.463c-9.68 3.51-23.202 3.5-30.711 3.5H25V192h23.5c9.747 0 26.265-.681 39.867-7.61c18.496-9.42 33.507-35.51 47.578-54.853c9.879-13.579 21.773-27.756 32.732-36.034C182.775 82.853 196.637 80 216.5 80H233z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Solo(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M208.714 161.607c0 5.765-.73 10.95-2.19 15.555c-1.46 4.604-3.388 8.666-5.784 12.185s-5.203 6.513-8.423 8.984c-3.219 2.47-6.55 4.492-9.995 6.065c-3.444 1.572-6.944 2.732-10.5 3.481c-3.557.749-6.907 1.123-10.051 1.123H49.129v-29.2H161.77c5.615 0 9.976-1.646 13.083-4.94c3.107-3.295 4.66-7.712 4.66-13.253c0-2.695-.411-5.166-1.235-7.412c-.823-2.246-2.002-4.192-3.537-5.84c-1.535-1.647-3.407-2.92-5.616-3.818c-2.208-.898-4.66-1.348-7.355-1.348H94.612c-4.717 0-9.808-.842-15.273-2.526c-5.466-1.685-10.538-4.399-15.217-8.142c-4.68-3.744-8.573-8.648-11.68-14.712c-3.107-6.065-4.66-13.477-4.66-22.237c0-8.76 1.553-16.153 4.66-22.18c3.107-6.027 7-10.93 11.68-14.712c4.679-3.78 9.751-6.513 15.217-8.198c5.465-1.685 10.556-2.527 15.273-2.527h99.39v29.2h-99.39c-5.54 0-9.864 1.684-12.97 5.053c-3.108 3.37-4.662 7.824-4.662 13.364c0 5.616 1.554 10.052 4.661 13.308c3.107 3.257 7.43 4.886 12.971 4.886h67.383c3.145.074 6.476.505 9.995 1.291c3.52.786 7.02 2.003 10.5 3.65c3.482 1.647 6.795 3.725 9.94 6.233c3.144 2.508 5.933 5.522 8.366 9.04c2.434 3.52 4.362 7.562 5.784 12.13c1.423 4.566 2.134 9.732 2.134 15.497"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M192 41.993v174.014c0 8.827-7.164 15.993-16 15.993H80c-8.844 0-16-7.16-16-15.993V41.993C64 33.166 71.164 26 80 26h96c8.844 0 16 7.16 16 15.993m-16 9.015c0-4.423-3.588-8.008-8-8.008H88c-4.419 0-8 3.573-8 8.008v155.984c0 4.423 3.588 8.008 8 8.008h80c4.419 0 8-3.573 8-8.008z"/><path d="M128 192c-17.673 0-32-14.327-32-32c0-17.673 14.327-32 32-32c17.673 0 32 14.327 32 32c0 17.673-14.327 32-32 32m.5-16c8.56 0 15.5-7.163 15.5-16s-6.94-16-15.5-16c-8.56 0-15.5 7.163-15.5 16s6.94 16 15.5 16m-.5-72c-12.703 0-23-10.521-23-23.5S115.297 57 128 57s23 10.521 23 23.5s-10.297 23.5-23 23.5m0-14.769c4.694 0 8.953-.25 8.953-9.087s-4.225-9.48-8.92-9.48c-4.694 0-8.919.643-8.919 9.48s4.192 9.087 8.886 9.087"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareswitchOff(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M25 185.377V70.39c0-3.527 2.878-6.39 6.428-6.39h192.95c3.547 0 6.427 2.861 6.427 6.39v114.987c0 3.527-2.878 6.39-6.428 6.39H31.427c-3.547 0-6.427-2.86-6.427-6.39m14.667-15.382c0 3.316 2.84 6.005 6.354 6.005h163.625c3.51 0 6.354-2.693 6.354-6.005v-83.99c0-3.316-2.841-6.005-6.354-6.005H46.021c-3.51 0-6.354 2.693-6.354 6.005z"/><path d="M120.236 163.017c0 1.766 1.789 3.197 3.996 3.197h8.008c2.207 0 3.996-1.43 3.996-3.197V92.75c0-1.765-1.79-3.197-3.996-3.197h-8.008c-2.207 0-3.996 1.43-3.996 3.197zm-32.726 0c0 1.766 1.789 3.197 3.996 3.197h8.008c2.207 0 3.996-1.43 3.996-3.197V92.75c0-1.765-1.789-3.197-3.996-3.197h-8.008c-2.207 0-3.996 1.43-3.996 3.197zm-31.726 0c0 1.766 1.79 3.197 3.997 3.197h8.007c2.207 0 3.996-1.43 3.996-3.197V92.75c0-1.765-1.788-3.197-3.996-3.197h-8.007c-2.207 0-3.997 1.43-3.997 3.197z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareswitchOn(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M231 70.623V185.61c0 3.528-2.878 6.391-6.428 6.391H31.622c-3.547 0-6.427-2.861-6.427-6.39V70.622c0-3.527 2.878-6.39 6.428-6.39h192.95c3.547 0 6.427 2.86 6.427 6.39zm-14.667 15.382c0-3.316-2.84-6.005-6.354-6.005H46.354C42.844 80 40 82.693 40 86.005v83.99c0 3.316 2.841 6.005 6.354 6.005h163.625c3.51 0 6.354-2.693 6.354-6.005z"/><path d="M135.764 92.983c0-1.766-1.789-3.197-3.996-3.197h-8.008c-2.207 0-3.996 1.43-3.996 3.197v70.267c0 1.765 1.79 3.197 3.996 3.197h8.008c2.207 0 3.996-1.43 3.996-3.197zm32.726 0c0-1.766-1.789-3.197-3.996-3.197h-8.008c-2.207 0-3.996 1.43-3.996 3.197v70.267c0 1.765 1.789 3.197 3.996 3.197h8.008c2.207 0 3.996-1.43 3.996-3.197zm31.726 0c0-1.766-1.79-3.197-3.997-3.197h-8.007c-2.207 0-3.996 1.43-3.996 3.197v70.267c0 1.765 1.788 3.197 3.996 3.197h8.007c2.207 0 3.997-1.43 3.997-3.197z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stereo(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M128.802 95.03c-9.229-9.369-22.39-15.228-37-15.228c-27.92 0-50.555 21.402-50.555 47.803c0 26.4 22.634 47.802 50.555 47.802c14.711 0 27.954-5.94 37.193-15.423c-12.232-16.88-14.177-19.888-14.177-32.38c0-12.016 5.924-18.458 14.19-31.142c6.753 13.293 13.629 19.445 13.629 31.538c0 12.802-6.03 20.525-13.402 32.614c9.206 9.115 22.185 14.793 36.567 14.793c27.922 0 50.556-21.401 50.556-47.802c0-26.4-22.634-47.803-50.556-47.803c-14.608 0-27.77 5.86-37 15.228M128 75.374C138.501 68.202 151.252 64 165 64c35.899 0 65 28.654 65 64c0 35.346-29.101 64-65 64c-13.748 0-26.499-4.202-37-11.374C117.499 187.798 104.748 192 91 192c-35.899 0-65-28.654-65-64c0-35.346 29.101-64 65-64c13.748 0 26.499 4.202 37 11.374"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M48.227 65.473c0-9.183 7.096-16.997 16.762-17.51c9.666-.513 116.887-.487 125.094-.487c8.207 0 17.917 9.212 17.917 17.71c0 8.499.98 117.936.49 126.609c-.49 8.673-9.635 15.995-17.011 15.995c-7.377 0-117.127-.327-126.341-.327c-9.214 0-17.472-7.793-17.192-16.1c.28-8.306.28-116.708.28-125.89zm15.951 4.684c-.153 3.953 0 112.665 0 116.19c0 3.524 3.115 5.959 7.236 6.156c4.12.198 112.165.288 114.852 0c2.686-.287 5.811-2.073 5.932-5.456c.12-3.383-.609-113.865-.609-116.89c0-3.025-3.358-5.84-6.02-5.924c-2.662-.085-110.503 0-114.155 0c-3.652 0-7.083 1.972-7.236 5.924"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thunderbolt(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M78.424 228.77c-.178 1.087.505 1.617 1.525 1.186l51.823-21.958c1.02-.433 1.149-1.346.29-2.039l-13.124-10.58l63.836-86.902c.654-.89.282-1.612-.808-1.612H128l53.414-79.142c.616-.912.21-1.65-.883-1.65h-34.145c-1.103 0-2.423.794-2.939 1.754L92.4 123.004c-.52.97-.044 1.755 1.053 1.755h48.942c1.102 0 1.477.73.834 1.638l-41.005 57.873l-13.127-9.535c-.895-.65-1.764-.301-1.943.792l-8.729 53.243z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeselect(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M64.254 25s48.253.303 54.601.303c6.35 0 4.887 10.672 9.524 10.596c4.637-.075 3.773-10.899 9.127-10.899c5.353 0 55.326.303 55.326.303v16.55l-55.049-.522l-.017 172.48l54.066.187V231s-51.437-1.257-56.008-.506c-4.57.75-4.41-9.416-7.816-9.55c-3.406-.134-2.774 8.985-6.56 8.985c-3.788 0-57.194.193-57.194.193l-.047-16.957l53.453.833l1.41-172.24l-54.816.096z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M55.265 167.072c-.975-1.973-3.388-2.796-5.372-1.847L42 169s22.5 53.5 85.5 56c60-1.5 96.627-48.626 97-96.5c.373-47.874-37-95.5-95.5-96c-57.5-1-79.556 45.004-79.556 45.004c-1.073 1.93-1.944 1.698-1.944-.501V51.997a4 4 0 0 0-4-3.997H37c-2.209 0-4 1.8-4 4.008v48.984A3.998 3.998 0 0 0 36.998 105h50.504a3.995 3.995 0 0 0 3.998-3.993v-6.014c0-2.205-1.79-4.02-4.008-4.053l-25.484-.38c-2.214-.033-3.223-1.679-2.182-3.628C59.826 86.932 78 45 128.5 45.5c49 .5 82.751 41.929 82.5 83.242C208 184 166 211 127.5 210c-54.5 0-72.235-42.928-72.235-42.928"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m82.757 128l-8.662-15.546C59.565 86.38 59.273 59.84 84.248 46.56l12.318-6.533c27.248-13.823 47.58 1.422 62.114 27.505l17.307 31.062l-14.612 7.413c-.95.482-6.824-12.246-15.349-27.546l-1.959-3.516c-10.23-18.36-18.63-29.368-38.728-19.173l-11.12 5.508c-25.096 12.71-9.972 35.728-5.5 43.756l1.96 3.515L101.515 128h82.478c4.416 0 8.007 3.586 8.007 8.01v84.98c0 4.43-3.585 8.01-8.007 8.01H72.007c-4.416 0-8.007-3.586-8.007-8.01v-84.98c0-4.43 3.585-8.01 8.007-8.01zM176 148.316c0-2.384-1.547-4.316-3.46-4.316H83.46c-1.911 0-3.46 1.936-3.46 4.316v60.368c0 2.384 1.547 4.316 3.46 4.316h89.08c1.911 0 3.46-1.936 3.46-4.316z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M129.042 230.413c14.807.56 24.14-9.993 24.14-20.644c0-10.652-10.127-19.1-15.966-19.1c0-6.462 0-34.093-.022-39.406c5.368-1.812 8.29-1.81 17.39-5.088c9.097-3.278 8.624-2.793 15.03-8.938c6.407-6.144 7.441-24.457 7.441-30.875c3.484-.18 13.856.036 13.856.036l-.422-33.11l-37.406.393v32.078h11.25c0 3.828.86 12.614-1.518 18.645c-2.377 6.031-2.992 5.485-9.732 8.245c-6.74 2.759-15.94 4.323-15.94 4.323s-.36-73.992-.144-73.992l13.224.001c1.102 0 1.52-.767.942-1.7l-22.217-35.878l-22.728 36.321c-.586.936-.17 1.676.943 1.651l15.465-.339s.75 97.485.732 101.874c-4.272 0-13.788-.045-20.224-5.307c-6.436-5.263-7.105-8.643-8.932-14.428c-1.828-5.785-2.512-12.533-2.512-17.572c6.495 0 17.021-10.184 17.021-19.98c0-9.798-11.035-22.247-21.027-22.247c-9.993 0-23.23 8.282-23.23 20.986c0 12.705 10.022 21.24 15.59 21.24c0 3.907.606 13.36 2.102 19.03c1.496 5.673 2.405 5.92 3.447 8.371c2.277 5.357 2.599 9.327 13.277 15.643c10.678 6.315 20.318 6.101 24.488 6.101c.25 2.888 0 10.361 0 13.922c-7.875 0-17.2 7.9-17.2 19.29c0 11.387 8.075 19.894 22.882 20.454"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vexpand(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M145.5 200.502c0 4.417-3.579 7.998-8.01 7.998h-16.98c-4.424 0-8.01-3.588-8.01-7.998V55.498c0-4.417 3.579-7.998 8.01-7.998h16.98c4.424 0 8.01 3.588 8.01 7.998zm-65 0c0 4.417-3.579 7.998-8.01 7.998H55.51c-4.424 0-8.01-3.588-8.01-7.998V55.498c0-4.417 3.579-7.998 8.01-7.998h16.98c4.424 0 8.01 3.588 8.01 7.998zm129 0c0 4.417-3.579 7.998-8.01 7.998h-16.98c-4.424 0-8.01-3.588-8.01-7.998V55.498c0-4.417 3.579-7.998 8.01-7.998h16.98c4.424 0 8.01 3.588 8.01 7.998z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VroundswitchOff(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M128 25c8.667 0 16.95 1.71 24.85 5.13c7.9 3.42 14.7 7.996 20.4 13.73c5.7 5.733 10.25 12.573 13.65 20.52c3.4 7.946 5.1 16.278 5.1 24.995v77.25c0 8.717-1.7 17.05-5.1 24.996c-3.4 7.946-7.95 14.786-13.65 20.52c-5.7 5.733-12.5 10.31-20.4 13.73C144.95 229.29 136.667 231 128 231c-8.667 0-16.95-1.71-24.85-5.13c-7.9-3.42-14.7-7.996-20.4-13.73c-5.7-5.733-10.25-12.573-13.65-20.52c-3.4-7.946-5.1-16.278-5.1-24.995v-77.25c0-8.717 1.7-17.05 5.1-24.996c3.4-7.946 7.95-14.786 13.65-20.52c5.7-5.733 12.5-10.31 20.4-13.73C111.05 26.71 119.333 25 128 25m.064 18c-6.508 0-12.72 1.36-18.634 4.08c-5.914 2.72-11.03 6.397-15.349 11.032c-4.318 4.634-7.744 10.125-10.279 16.472C81.267 80.93 80 87.596 80 94.58v31.47c3.802-5.44 16.147-16.472 23.281-19.947c7.135-3.476 17.96-5.642 24.468-5.642c.782 0 1.577-.022 2.397-.044h.001c5.999-.166 13.327-.367 26.831 7.701c9.439 5.639 16.357 13.944 19.15 17.909v.628c1.88 2.637 1.748 1.855 0-.628V94.58c0-6.985-1.267-13.65-3.801-19.997c-2.535-6.347-5.962-11.838-10.28-16.472c-4.318-4.635-9.434-8.312-15.348-11.032A44.086 44.086 0 0 0 128.064 43m-18.674 78.797C115.297 119.266 121.5 118 128 118s12.703 1.266 18.609 3.797c5.907 2.531 11.016 5.953 15.328 10.266c4.313 4.312 7.735 9.421 10.266 15.327C174.734 153.297 176 159.5 176 166s-1.266 12.703-3.797 18.609c-2.531 5.907-5.953 11.016-10.266 15.328c-4.312 4.313-9.421 7.735-15.327 10.266C140.703 212.734 134.5 214 128 214s-12.703-1.266-18.609-3.797c-5.907-2.531-11.016-5.953-15.328-10.266c-4.313-4.312-7.735-9.421-10.266-15.327C81.266 178.703 80 172.5 80 166s1.266-12.703 3.797-18.609c2.531-5.907 5.953-11.016 10.265-15.328c4.313-4.313 9.422-7.735 15.328-10.266" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VroundswitchOn(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M128 231c-8.667 0-16.95-1.71-24.85-5.13c-7.9-3.42-14.7-7.996-20.4-13.73c-5.7-5.733-10.25-12.573-13.65-20.52c-3.4-7.946-5.1-16.278-5.1-24.995v-77.25c0-8.717 1.7-17.05 5.1-24.996c3.4-7.946 7.95-14.786 13.65-20.52c5.7-5.733 12.5-10.31 20.4-13.73C111.05 26.71 119.333 25 128 25c8.667 0 16.95 1.71 24.85 5.13c7.9 3.42 14.7 7.996 20.4 13.73c5.7 5.733 10.25 12.573 13.65 20.52c3.4 7.946 5.1 16.278 5.1 24.995v77.25c0 8.717-1.7 17.05-5.1 24.996c-3.4 7.946-7.95 14.786-13.65 20.52c-5.7 5.733-12.5 10.31-20.4 13.73C144.95 229.29 136.667 231 128 231m-.064-18c6.508 0 12.72-1.36 18.634-4.08c5.914-2.72 11.03-6.397 15.349-11.032c4.318-4.634 7.744-10.125 10.279-16.472c2.535-6.346 3.802-13.012 3.802-19.997v-31.47c-3.802 5.44-16.147 16.472-23.281 19.947c-7.135 3.476-17.96 5.642-24.468 5.642c-.783 0-1.578.022-2.398.044c-5.999.166-13.327.367-26.831-7.701c-9.44-5.639-16.357-13.944-19.15-17.909v-.628c-1.88-2.637-1.748-1.855 0 .628v31.447c0 6.985 1.267 13.65 3.801 19.997c2.535 6.347 5.962 11.838 10.28 16.472c4.318 4.635 9.434 8.312 15.348 11.032a44.085 44.085 0 0 0 18.635 4.08m18.674-78.797C140.703 136.734 134.5 138 128 138s-12.703-1.266-18.609-3.797c-5.907-2.531-11.016-5.953-15.328-10.265c-4.313-4.313-7.735-9.422-10.266-15.328C81.266 102.703 80 96.5 80 90s1.266-12.703 3.797-18.609c2.531-5.907 5.953-11.016 10.265-15.328c4.313-4.313 9.422-7.735 15.328-10.266C115.297 43.266 121.5 42 128 42s12.703 1.266 18.609 3.797c5.907 2.531 11.016 5.953 15.328 10.265c4.313 4.313 7.735 9.422 10.266 15.328C174.734 77.297 176 83.5 176 90s-1.266 12.703-3.797 18.609c-2.531 5.907-5.953 11.016-10.266 15.328c-4.312 4.313-9.421 7.735-15.327 10.266" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VsquareswitchOff(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M185.377 231H70.39c-3.528 0-6.391-2.878-6.391-6.428V31.622c0-3.547 2.861-6.427 6.39-6.427h114.987c3.527 0 6.39 2.878 6.39 6.428v192.95c0 3.547-2.86 6.427-6.39 6.427zm-15.382-14.667c3.316 0 6.005-2.84 6.005-6.354V46.354c0-3.51-2.693-6.354-6.005-6.354h-83.99C82.69 40 80 42.841 80 46.354v163.625c0 3.51 2.693 6.354 6.005 6.354z"/><path d="M163.017 135.764c1.766 0 3.197-1.789 3.197-3.996v-8.008c0-2.207-1.43-3.996-3.197-3.996H92.75c-1.765 0-3.197 1.79-3.197 3.996v8.008c0 2.207 1.43 3.996 3.197 3.996zm0 32.726c1.766 0 3.197-1.789 3.197-3.996v-8.008c0-2.207-1.43-3.996-3.197-3.996H92.75c-1.765 0-3.197 1.789-3.197 3.996v8.008c0 2.207 1.43 3.996 3.197 3.996zm0 31.726c1.766 0 3.197-1.79 3.197-3.997v-8.007c0-2.207-1.43-3.996-3.197-3.996H92.75c-1.765 0-3.197 1.788-3.197 3.996v8.007c0 2.207 1.43 3.997 3.197 3.997z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VsquareswitchOn(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M70.623 25H185.61c3.528 0 6.391 2.878 6.391 6.428v192.95c0 3.547-2.861 6.427-6.39 6.427H70.622c-3.527 0-6.39-2.878-6.39-6.428V31.427c0-3.547 2.86-6.427 6.39-6.427zm15.382 14.667c-3.316 0-6.005 2.84-6.005 6.354v163.625c0 3.51 2.693 6.354 6.005 6.354h83.99c3.316 0 6.005-2.841 6.005-6.354V46.021c0-3.51-2.693-6.354-6.005-6.354z"/><path d="M92.983 120.236c-1.766 0-3.197 1.789-3.197 3.996v8.008c0 2.207 1.43 3.996 3.197 3.996h70.267c1.765 0 3.197-1.79 3.197-3.996v-8.008c0-2.207-1.43-3.996-3.197-3.996zm0-32.726c-1.766 0-3.197 1.789-3.197 3.996v8.008c0 2.207 1.43 3.996 3.197 3.996h70.267c1.765 0 3.197-1.789 3.197-3.996v-8.008c0-2.207-1.43-3.996-3.197-3.996zm0-31.726c-1.766 0-3.197 1.79-3.197 3.997v8.007c0 2.207 1.43 3.996 3.197 3.996h70.267c1.765 0 3.197-1.788 3.197-3.996v-8.007c0-2.207-1.43-3.997-3.197-3.997z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Waveform(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M32 225h12.993A4.004 4.004 0 0 0 49 220.997V138.01c0-4.976.724-5.04 1.614-.16l12.167 66.708c.397 2.177 2.516 3.942 4.713 3.942h8.512a3.937 3.937 0 0 0 3.947-4S79 127.5 80 129s14.488 52.67 14.488 52.67c.559 2.115 2.8 3.83 5.008 3.83h8.008a3.993 3.993 0 0 0 3.996-3.995v-43.506c0-4.97 1.82-5.412 4.079-.965l10.608 20.895c1.001 1.972 3.604 3.571 5.806 3.571h9.514a3.999 3.999 0 0 0 3.993-4.001v-19.49c0-4.975 2.751-6.074 6.155-2.443l6.111 6.518c1.51 1.61 4.528 2.916 6.734 2.916h7c2.21 0 5.567-.855 7.52-1.92l9.46-5.16c1.944-1.06 5.309-1.92 7.524-1.92h23.992a4.002 4.002 0 0 0 4.004-3.992v-7.516a3.996 3.996 0 0 0-4.004-3.992h-23.992c-2.211 0-5.601.823-7.564 1.834l-4.932 2.54c-4.423 2.279-12.028 3.858-16.993 3.527l2.97.198c-4.962-.33-10.942-4.12-13.356-8.467l-11.19-20.14c-1.07-1.929-3.733-3.492-5.939-3.492h-7c-2.21 0-4 1.794-4 4.001v19.49c0 4.975-1.14 5.138-2.542.382l-12.827-43.535c-.625-2.12-2.92-3.838-5.127-3.838h-8.008c-2.207 0-3.916 1.784-3.817 4.005l1.92 42.998c.221 4.969-.489 5.068-1.585.224l-15.13-66.825c-.488-2.155-2.681-3.902-4.878-3.902h-8.512a3.937 3.937 0 0 0-3.947 4s.953 77-.047 75.5s-13.937-92.072-13.937-92.072C49.252 34.758 47.21 33 45 33H31.999"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xlrplug(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M128 220c-50.81 0-92-41.19-92-92c0-34.045 18.492-63.77 45.98-79.68c26.471 0 69.646.367 92.67.367C201.79 64.685 220 94.216 220 128c0 50.81-41.19 92-92 92m0-15c42.526 0 77-34.474 77-77c0-26.467-13.353-49.815-33.69-63.675c-21.734 0-65.127-.18-86.353-.18C64.47 77.98 51 101.418 51 128c0 42.526 34.474 77 77 77"/><circle cx="128" cy="176" r="16"/><circle cx="176" cy="134" r="16"/><circle cx="80" cy="135" r="16"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoomin(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M120.46 158.29c-30.166 8.65-61.631-8.792-70.281-38.957c-8.65-30.165 8.792-61.63 38.957-70.28c30.165-8.65 61.63 8.792 70.28 38.957c4.417 15.403-1.937 38.002-9.347 50.872c-.614 1.067 59.212 53.064 59.212 53.064l-17.427 17.63l-53.514-56.72s-10.233 3.241-17.88 5.434M104 144c22.091 0 40-17.909 40-40s-17.909-40-40-40s-40 17.909-40 40s17.909 40 40 40"/><path d="M111.912 80.047h-15.95v16.037H80v15.992h15.962V128h15.95v-15.924H128V96.084h-16.088z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoomout(children ...ElementRenderer) *FadIcon {
	return &FadIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M120.46 158.29c-30.166 8.65-61.631-8.792-70.281-38.957c-8.65-30.165 8.792-61.63 38.957-70.28c30.165-8.65 61.63 8.792 70.28 38.957c4.417 15.403-1.937 38.002-9.347 50.872c-.614 1.067 59.212 53.064 59.212 53.064l-17.427 17.63l-53.514-56.72s-10.233 3.241-17.88 5.434M104 144c22.091 0 40-17.909 40-40s-17.909-40-40-40s-40 17.909-40 40s17.909 40 40 40"/><path d="M80 96.084v15.992h48V96.084z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
