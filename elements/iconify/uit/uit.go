package uit

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "4.0.8"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type UitIcon struct {
	*SVGSVGElement
}

func AdobeAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.917 2.224A.5.5 0 0 0 9.5 2h-8a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .96.194l8-19a.504.504 0 0 0-.043-.47M2 19.024V3h6.747zM22.5 2h-8a.5.5 0 0 0-.46.694l8 19A.5.5 0 0 0 23 21.5v-19a.5.5 0 0 0-.5-.5M22 19.024L15.253 3H22zm-9.532-9.7A.498.498 0 0 0 12.003 9H12a.5.5 0 0 0-.466.318l-3.5 9A.5.5 0 0 0 8.5 19h3.191l1.362 2.724A.5.5 0 0 0 13.5 22h3a.5.5 0 0 0 .468-.676zM13.808 21l-1.36-2.724A.501.501 0 0 0 12 18H9.23l2.761-7.099L15.778 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplay(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 3h-15A2.502 2.502 0 0 0 2 5.5v10A2.502 2.502 0 0 0 4.5 18h1a.5.5 0 0 0 0-1h-1A1.5 1.5 0 0 1 3 15.5v-10A1.5 1.5 0 0 1 4.5 4h15A1.5 1.5 0 0 1 21 5.5v10a1.5 1.5 0 0 1-1.5 1.5h-1a.5.5 0 0 0 0 1h1a2.502 2.502 0 0 0 2.5-2.5v-10A2.502 2.502 0 0 0 19.5 3m-6.259 11.44a1.557 1.557 0 0 0-2.482-.002l-2.863 4.22A1.5 1.5 0 0 0 9.136 21h5.727a1.5 1.5 0 0 0 1.241-2.342zM14.863 20H9.137a.5.5 0 0 1-.413-.781L11.587 15a.494.494 0 0 1 .413-.219a.492.492 0 0 1 .413.22l2.863 4.218a.5.5 0 0 1-.413.781"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 12h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1m0 4h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1m3-11h8a.5.5 0 0 0 0-1h-8a.5.5 0 0 0 0 1m-3 15h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1m0-12h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1m7 12h-4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1m4-8h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1m-11-8h-4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1m11 4h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1m0 8h-8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 10a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1zm-4-3h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m15 11h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1m4-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 7h15a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1m17 4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m-2 5h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterJustify(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 20h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1m4-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1M2.5 5h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m19 3h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m0 4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 7h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m0 4h15a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1m15 7h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1m4-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeftJustify(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 20h-13a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1M2.5 5h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m19 7h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m0-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m0 8h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLetterRight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 4h12a.5.5 0 0 0 0-1h-12a.5.5 0 0 0 0 1m12 17h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m0-4h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1m0-6h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m0-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 18h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1M2.5 7h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m19 3h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1m0 4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRightJustify(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 20h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1M2.5 5h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m19 7h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m0 4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m0-8h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 13.5h-3a.5.5 0 0 0 0 1h2.482a7 7 0 0 1-6.482 6.482V11h2a.5.5 0 0 0 0-1h-2V7.95A2.998 2.998 0 0 0 15 5a3 3 0 1 0-6 0a2.994 2.994 0 0 0 2.5 2.95V10h-2a.5.5 0 0 0 0 1h2v9.974c-3.419-.241-6.23-2.956-6.482-6.474H7.5a.5.5 0 0 0 0-1h-3a.5.5 0 0 0-.5.5a8.01 8.01 0 0 0 8 8a8.01 8.01 0 0 0 8-8a.5.5 0 0 0-.5-.5M10 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 9a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 1 0v-5a.5.5 0 0 0-.5-.5m12.344-4.602l1.073-1.622a.5.5 0 0 0-.834-.552l-1.06 1.602a5.946 5.946 0 0 0-6.046 0l-1.06-1.602a.5.5 0 0 0-.834.552l1.073 1.622A5.988 5.988 0 0 0 6 9v8.5a.5.5 0 0 0 .5.5h3v3.5a.5.5 0 0 0 1 0V18h3v3.5a.5.5 0 0 0 1 0V18h3a.5.5 0 0 0 .5-.5V9a5.99 5.99 0 0 0-2.156-4.602M17 17H7v-7h10zM7 9a5 5 0 0 1 10 0zm13.5 0a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 1 0v-5a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleDown(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.646 10.854a.498.498 0 0 0 .707 0l3.5-3.5a.5.5 0 0 0-.707-.708L12 9.793L8.853 6.647a.5.5 0 0 0-.707.707zm3.5 1.792L12 15.793l-3.147-3.146a.5.5 0 0 0-.707.707l3.5 3.5a.498.498 0 0 0 .707 0l3.5-3.5a.5.5 0 0 0-.707-.708"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.207 12l3.146-3.146a.5.5 0 0 0-.707-.707l-3.5 3.5a.5.5 0 0 0 0 .707l3.5 3.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708zm9.147 3.146L14.207 12l3.146-3.146a.5.5 0 0 0-.707-.707l-3.5 3.5a.5.5 0 0 0 0 .707l3.5 3.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleRight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.353 8.147a.5.5 0 0 0-.707.707L9.793 12l-3.147 3.146a.5.5 0 1 0 .707.708l3.5-3.5a.5.5 0 0 0 0-.707zm9.5 3.5l-3.5-3.5a.5.5 0 0 0-.707.707L15.793 12l-3.147 3.146a.5.5 0 1 0 .707.708l3.5-3.5a.5.5 0 0 0 0-.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleUp(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.854 13.647l-4.5-4.5a.5.5 0 0 0-.707 0l-4.5 4.5a.5.5 0 0 0 .707.707L12 10.207l4.146 4.147a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ankh(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 13h-4.807A7.783 7.783 0 0 0 17 6.9a5.001 5.001 0 0 0-10 0a7.783 7.783 0 0 0 3.307 6.1H5.5a.5.5 0 0 0 0 1h6v7.5a.5.5 0 0 0 1 0V14h6a.5.5 0 0 0 0-1m-6.501-.07C11.132 12.436 8 10.399 8 6.9a4.001 4.001 0 0 1 8 0c0 3.488-3.134 5.533-4.001 6.03"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.326 16.207a3.892 3.892 0 0 1-2.368-3.587a4.005 4.005 0 0 1 1.925-3.375a.5.5 0 0 0 .15-.707a5.125 5.125 0 0 0-4.02-2.2a3.673 3.673 0 0 0-.518-.016a6.071 6.071 0 0 0 1.402-4.376a.485.485 0 0 0-.516-.446a6.165 6.165 0 0 0-5.55 5.335a5.234 5.234 0 0 0-2.006-.451A5.42 5.42 0 0 0 4.261 9.14c-1.785 3.1-.668 7.78 1.29 10.611c.76 1.099 1.901 2.748 3.5 2.748l.088-.001a4.258 4.258 0 0 0 1.616-.426a3.772 3.772 0 0 1 1.641-.393c.54.005 1.073.135 1.555.379a4.144 4.144 0 0 0 1.755.415c1.655-.03 2.66-1.497 3.466-2.675a11.476 11.476 0 0 0 1.447-2.978a.503.503 0 0 0-.293-.614M13.051 4.244a5.363 5.363 0 0 1 2.874-1.685a5.206 5.206 0 0 1-1.228 3.16a4.924 4.924 0 0 1-1.413 1.16l-.018.01l-.12.048a3.96 3.96 0 0 1-1.013.316a1.273 1.273 0 0 1-.327-.057a5.112 5.112 0 0 1 1.245-2.952m5.295 14.99c-.889 1.3-1.6 2.221-2.658 2.24a3.122 3.122 0 0 1-1.34-.333a4.48 4.48 0 0 0-1.952-.461a4.728 4.728 0 0 0-2.033.473a3.353 3.353 0 0 1-1.262.346l-.051.001c-.99 0-1.832-1.095-2.677-2.316c-1.527-2.21-2.924-6.63-1.245-9.544a4.417 4.417 0 0 1 3.714-2.256h.042a5.206 5.206 0 0 1 1.827.478c.132.052.262.104.39.152c.024.011.049.02.075.027c.304.125.628.196.957.212a4.12 4.12 0 0 0 1.38-.386a5.246 5.246 0 0 1 2.428-.531a4.196 4.196 0 0 1 2.963 1.376a4.913 4.913 0 0 0-1.946 3.918a4.853 4.853 0 0 0 2.57 4.299a10.585 10.585 0 0 1-1.182 2.305"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDown(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.146 11.646L12.5 15.293V7.5a.5.5 0 0 0-1 0v7.793l-3.647-3.646a.5.5 0 0 0-.707.707l4.5 4.5A.5.5 0 0 0 12 17c.011 0 .02-.005.03-.006a.498.498 0 0 0 .163-.033a.5.5 0 0 0 .162-.109l4.498-4.498a.5.5 0 0 0-.707-.708M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 11.5H8.707l3.646-3.646a.5.5 0 0 0-.707-.707l-4.5 4.5a.499.499 0 0 0-.145.35L7 12a.5.5 0 0 0 .146.354l4.5 4.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708L8.708 12.5H16.5a.5.5 0 0 0 0-1M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.961 11.809a.5.5 0 0 0-.108-.163l-4.5-4.5a.5.5 0 0 0-.707.708l3.647 3.646H7.5a.5.5 0 0 0 0 1h7.793l-3.647 3.646a.5.5 0 1 0 .707.708l4.5-4.5a.5.5 0 0 0 .145-.344L17 12l-.001-.007a.5.5 0 0 0-.038-.184M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUp(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.354 7.147a.5.5 0 0 0-.344-.145L12 7l-.007.001a.5.5 0 0 0-.347.146l-4.5 4.5a.5.5 0 0 0 .708.707L11.5 8.707v7.794a.5.5 0 1 0 1 0V8.706l3.646 3.647a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707zM12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.452 16.904H7.867l9.97-9.97a.546.546 0 1 0-.77-.772l-9.971 9.97V6.548a.548.548 0 0 0-1.096 0v10.904c0 .303.245.548.548.548h10.904a.548.548 0 0 0 0-1.096"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.453 6a.548.548 0 0 0-.549.548v9.585l-9.97-9.971a.546.546 0 0 0-.772.771l9.97 9.971H6.549a.548.548 0 0 0 0 1.096h10.904a.548.548 0 0 0 .548-.548V6.548A.548.548 0 0 0 17.453 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.868 7.096h9.584a.548.548 0 0 0 0-1.096H6.548A.548.548 0 0 0 6 6.548v10.904a.548.548 0 1 0 1.096 0V7.867l9.97 9.97a.543.543 0 0 0 .773 0a.545.545 0 0 0-.001-.77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.452 6H6.547a.548.548 0 0 0 0 1.096h9.585l-9.97 9.97a.545.545 0 1 0 .772.772l9.97-9.971v9.586a.548.548 0 0 0 1.096 0V6.546A.548.548 0 0 0 17.452 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.808 2.162a10.031 10.031 0 0 0-4.456.208C4.014 3.85.887 9.38 2.37 14.719a9.846 9.846 0 0 0 4.15 5.665a10.22 10.22 0 0 0 5.492 1.613A9.96 9.96 0 0 0 17 20.66a.5.5 0 1 0-.5-.865a9.084 9.084 0 0 1-9.444-.255a8.831 8.831 0 0 1-3.722-5.084A9.032 9.032 0 0 1 13.63 3.146A9.153 9.153 0 0 1 21 12.291v.209a2.5 2.5 0 1 1-5 0v-4a.5.5 0 0 0-1 0v.88A3.973 3.973 0 0 0 12 8a4 4 0 1 0 0 8a3.99 3.99 0 0 0 3.397-1.914A3.487 3.487 0 0 0 18.5 16a3.5 3.5 0 0 0 3.5-3.5v-.209a10.152 10.152 0 0 0-8.192-10.129M12 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bag(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 6H16V5a2.003 2.003 0 0 0-2-2h-4a2.003 2.003 0 0 0-2 2v1H4.5A2.502 2.502 0 0 0 2 8.5v10A2.502 2.502 0 0 0 4.5 21h15a2.502 2.502 0 0 0 2.5-2.5v-10A2.502 2.502 0 0 0 19.5 6M9 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1H9zm12 13.5a1.5 1.5 0 0 1-1.5 1.5h-15A1.5 1.5 0 0 1 3 18.5v-6.473l5.842 1.948A.51.51 0 0 0 9 14h6a.51.51 0 0 0 .158-.025L21 12.027zm0-7.494a.49.49 0 0 0-.158.02L14.919 13H9.081l-5.923-1.975a.49.49 0 0 0-.158-.02V8.5A1.5 1.5 0 0 1 4.5 7h15A1.5 1.5 0 0 1 21 8.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryBolt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 16h-4a.501.501 0 0 1-.5-.5v-7a.5.5 0 0 1 .5-.5h3a.5.5 0 0 0 0-1h-3A1.5 1.5 0 0 0 2 8.5v7A1.5 1.5 0 0 0 3.5 17h4a.5.5 0 0 0 0-1m5.966-4.167c-.01-.029-.025-.053-.04-.079a.482.482 0 0 0-.056-.083c-.022-.025-.048-.043-.074-.063c-.019-.014-.032-.034-.053-.046c-.008-.004-.018-.004-.027-.008c-.027-.014-.058-.02-.088-.028c-.035-.01-.069-.02-.104-.021c-.008 0-.016-.005-.024-.005H8.85l2.088-3.757a.5.5 0 1 0-.876-.486l-2.5 4.5c-.004.008-.004.018-.008.027c-.014.028-.02.058-.028.09c-.01.034-.02.067-.021.101c0 .009-.005.016-.005.025c0 .022.01.04.012.061c.005.036.01.07.022.105c.01.03.025.055.04.081a.477.477 0 0 0 .056.082c.022.025.048.043.074.063c.019.015.032.034.053.046c.007.003.016.002.023.005A.49.49 0 0 0 8 12.5h4.15l-2.088 3.757a.501.501 0 0 0 .876.486l2.5-4.5c.004-.008.004-.018.008-.027c.014-.027.02-.058.028-.088c.01-.035.02-.069.021-.104c0-.008.005-.016.005-.024c0-.021-.01-.039-.012-.06a.492.492 0 0 0-.022-.107M17.5 7h-4a.5.5 0 0 0 0 1h4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-3a.5.5 0 0 0 0 1h3a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 17.5 7m4 3a.5.5 0 0 0-.5.5v3a.5.5 0 1 0 1 0v-3a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 10a.5.5 0 0 0-.5.5v3a.5.5 0 1 0 1 0v-3a.5.5 0 0 0-.5-.5m-4-3h-14A1.5 1.5 0 0 0 2 8.5v7A1.5 1.5 0 0 0 3.5 17h14a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 17.5 7m.5 8.5a.5.5 0 0 1-.5.5h-14a.501.501 0 0 1-.5-.5v-7a.5.5 0 0 1 .5-.5h14a.5.5 0 0 1 .5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 6h5a.5.5 0 0 0 0-1h-5a.5.5 0 0 0 0 1m7.026 9.003a.494.494 0 0 0-.553.441A3.999 3.999 0 0 1 14 15v-2h8.47a.5.5 0 0 0 .497-.553A5.004 5.004 0 0 0 18 8a5.006 5.006 0 0 0-5 5v2a4.999 4.999 0 0 0 9.967.556a.5.5 0 0 0-.44-.553M18 9a4.015 4.015 0 0 1 3.874 3h-7.747A4.006 4.006 0 0 1 18 9m-8.604 2.434A3.495 3.495 0 0 0 7.5 5h-6a.5.5 0 0 0-.5.5v14a.5.5 0 0 0 .5.5h6a4.491 4.491 0 0 0 1.896-8.566M2 6h5.5a2.5 2.5 0 1 1 0 5H2zm5.5 13H2v-7h5.5a3.5 3.5 0 0 1 0 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.722 11.548A3.735 3.735 0 0 0 13.25 5H12V3.5a.5.5 0 0 0-1 0V5H8V3.5a.5.5 0 0 0-1 0V5H5.5a.5.5 0 0 0 0 1H7v12H5.5a.5.5 0 0 0 0 1H7v1.5a.5.5 0 0 0 1 0V19h3v1.5a.5.5 0 0 0 1 0V19h3.25a3.74 3.74 0 0 0 .472-7.452M8 6h5.25a2.75 2.75 0 1 1 0 5.5H8zm7.25 12H8v-5.5h7.25a2.75 2.75 0 1 1 0 5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BloggerAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 13.5h-4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1m-4-3h1a.5.5 0 0 0 0-1h-1a.5.5 0 0 0 0 1m6 0a.501.501 0 0 1-.5-.5V9A3.504 3.504 0 0 0 12 5.5h-2A4.505 4.505 0 0 0 5.5 10v4a4.505 4.505 0 0 0 4.5 4.5h4a4.505 4.505 0 0 0 4.5-4.5v-1a2.502 2.502 0 0 0-2.5-2.5m1.5 3.5a3.504 3.504 0 0 1-3.5 3.5h-4A3.504 3.504 0 0 1 6.5 14v-4A3.504 3.504 0 0 1 10 6.5h2A2.502 2.502 0 0 1 14.5 9v1c0 .828.672 1.5 1.5 1.5s1.5.672 1.5 1.5zM20 1H4a3.003 3.003 0 0 0-3 3v16a3.003 3.003 0 0 0 3 3h16a3.003 3.003 0 0 0 3-3V4a3.003 3.003 0 0 0-3-3m2 19a2.003 2.003 0 0 1-2 2H4a2.003 2.003 0 0 1-2-2V4a2.003 2.003 0 0 1 2-2h16a2.003 2.003 0 0 1 2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8a3.003 3.003 0 0 0-3 3v16.5a.5.5 0 0 0 .75.434l6.25-3.6l6.25 3.6A.5.5 0 0 0 19 21.5V5a3.003 3.003 0 0 0-3-3m2 18.635l-5.75-3.312a.51.51 0 0 0-.5 0L6 20.635V5a2.003 2.003 0 0 1 2-2h8a2.003 2.003 0 0 1 2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M21 3.5a.5.5 0 0 0-.5-.5h-17a.5.5 0 0 0-.5.5v17a.5.5 0 1 0 1 0V4h16.5a.5.5 0 0 0 .5-.5M11.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m8 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-8 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m8-10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-8-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottom(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 8m4 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 16 12m4-6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 20 4m-4 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 16 4m-4 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 4M4 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 4 4m8 10a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 12m9 8H3a.5.5 0 0 0 0 1h18a.5.5 0 0 0 0-1m-1-5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M4 18a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 4 16M8 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 8 4m4 14a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 16m-8-2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 4 12m0-2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 4 8m4 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 8 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderClear(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 7a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M4 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m16-7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 20 4m-4-1a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4-14a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M8 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 14a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M8 11a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4-10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderHorizontal(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 3.5m-8 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 4 3.5m12 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 16 3.5m4 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 20 3.5m-12 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 8 3.5m4 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 7.5m8 7a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-16-7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 4 7.5m12 11a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m5-9H3a.5.5 0 0 0 0 1h18a.5.5 0 0 0 0-1m-1-5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-16-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m8-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInner(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 8 3.5m13 8h-8.5V3a.5.5 0 0 0-1 0v8.5H3a.5.5 0 0 0 0 1h8.5V21a.5.5 0 1 0 1 0v-8.5H21a.5.5 0 0 0 0-1m-1-6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 20 3.5m0 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-3a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 16 3.5m-12 11a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m12 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-8-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m12-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M4 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 4 7.5m16 7a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M4 5.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 4 3.5m0 15a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 2.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-10A.5.5 0 0 0 3 3v18a.5.5 0 1 0 1 0V3a.5.5 0 0 0-.5-.5m16 3a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 19.5 3.5m-8 15a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m8-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-18a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-8 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m8-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-8-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOut(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 13.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 16 11.5M20.5 3h-17a.5.5 0 0 0-.5.5v17a.5.5 0 0 0 .5.5h17a.5.5 0 0 0 .5-.5v-17a.5.5 0 0 0-.5-.5M20 20H4V4h16zm-8-2.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 15.5m0-6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 7.5m-4 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 8 11.5m4 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 12 11.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTop(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M3 4h18a.5.5 0 0 0 0-1H3a.5.5 0 0 0 0 1m5 14a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M4 10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m16-10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-8 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M8 10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4-2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderVertical(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m16-7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 20 3.5m-16 15a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m12-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-14a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M8 2.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-18a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m8-2a.5.5 0 0 0-.5.5v18a.5.5 0 1 0 1 0V3a.5.5 0 0 0-.5-.5m4 16a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calender(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 4h-3V2.5a.5.5 0 0 0-1 0V4h-7V2.5a.5.5 0 0 0-1 0V4h-3A2.503 2.503 0 0 0 2 6.5v13A2.503 2.503 0 0 0 4.5 22h15a2.502 2.502 0 0 0 2.5-2.5v-13A2.502 2.502 0 0 0 19.5 4M21 19.5a1.5 1.5 0 0 1-1.5 1.5h-15A1.5 1.5 0 0 1 3 19.5V11h18zm0-9.5H3V6.5C3 5.672 3.67 5 4.5 5h3v1.5a.5.5 0 0 0 1 0V5h7v1.5a.5.5 0 0 0 1 0V5h3A1.5 1.5 0 0 1 21 6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartGrowth(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 21H20V4.5a.5.5 0 0 0-1 0V21h-3V8.5a.5.5 0 0 0-1 0V21h-3v-8.5a.5.5 0 0 0-1 0V21H8v-4.5a.5.5 0 0 0-1 0V21H3V2.5a.5.5 0 0 0-1 0v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m.5 1.03c4.56.252 8.215 3.923 8.46 8.486H12.5zM12 21c-4.963 0-9-4.037-9-9c0-4.794 3.77-8.713 8.5-8.975V12a.5.5 0 0 0 .067.25l4.488 7.774A8.933 8.933 0 0 1 12 21m4.917-1.482l-4.042-7.002h8.076a9.002 9.002 0 0 1-4.034 7.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.353 7.574a.5.5 0 0 0-.707-.008L9.84 15.373l-3.487-3.486a.5.5 0 0 0-.707.707l3.84 3.84a.498.498 0 0 0 .707 0l8.16-8.16a.5.5 0 0 0 0-.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.809 8.646l-5.647 5.647l-2.47-2.47l-.008-.007a.5.5 0 1 0-.7.714l2.825 2.824a.498.498 0 0 0 .707 0l6-6l.007-.008a.5.5 0 0 0-.714-.7M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.847 15.754a.5.5 0 0 0 .707 0l6.8-6.8a.5.5 0 0 0-.708-.708L10.2 14.693l-2.847-2.846a.5.5 0 0 0-.707.707zM21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5M21 21H3V3h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLayer(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.979 2.068c-3.73-.524-7.18 2-7.849 5.672a5.985 5.985 0 0 0-3.086 4.658A5 5 0 0 0 7 22a5.007 5.007 0 0 0 4.607-3.06a5.963 5.963 0 0 0 4.635-3.076a6.985 6.985 0 0 0 5.69-5.893a6.996 6.996 0 0 0-5.953-7.903M7 21a4 4 0 1 1 0-8a4 4 0 0 1 0 8m4.907-3.089c.055-.296.093-.599.093-.911a5 5 0 0 0-5-5c-.312 0-.615.037-.912.092a4.997 4.997 0 1 1 5.82 5.82m4.82-3.172l-.02.032a5.94 5.94 0 0 0 .288-1.529a5.997 5.997 0 0 0-5.746-6.237a5.973 5.973 0 0 0-1.985.258a5.996 5.996 0 1 1 7.464 7.476"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circuit(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m11.5 2a.5.5 0 0 0 0-1H19v-3h2.5a.5.5 0 0 0 0-1H19A2.502 2.502 0 0 0 16.5 5V2.5a.5.5 0 0 0-1 0V5h-3V2.5a.5.5 0 0 0-1 0V5h-3V2.5a.5.5 0 0 0-1 0V5A2.502 2.502 0 0 0 5 7.5H2.5a.5.5 0 0 0 0 1H5v3H2.5a.5.5 0 0 0 0 1H5v3H2.5a.5.5 0 0 0 0 1H5A2.502 2.502 0 0 0 7.5 19v2.5a.5.5 0 1 0 1 0V19h3v2.5a.5.5 0 1 0 1 0V19h3v2.5a.5.5 0 1 0 1 0V19a2.502 2.502 0 0 0 2.5-2.5h2.5a.5.5 0 0 0 0-1H19v-3zm-3.5 4a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 6 16.5v-9A1.5 1.5 0 0 1 7.5 6h9A1.5 1.5 0 0 1 18 7.5zm-4-8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClinicMedical(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 12v1.5H10c-.3 0-.5.2-.5.5s.2.5.5.5h1.5V16c0 .3.2.5.5.5s.5-.2.5-.5v-1.5H14c.3 0 .5-.2.5-.5s-.2-.5-.5-.5h-1.5V12c0-.3-.2-.5-.5-.5s-.5.2-.5.5m10.3-1.4l-9.5-8.4c-.2-.2-.5-.2-.7 0l-9.5 8.4c-.2.2-.2.5 0 .7c.2.2.5.2.7 0l1.2-1v11.2c0 .3.2.5.5.5h15c.3 0 .5-.2.5-.5V10.3l1.2 1c.1.1.2.1.3.1c.1 0 .3-.1.4-.2c.2-.1.1-.5-.1-.6M19 21H5V9.4l7-6.2l7 6.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.848 13.067L12.5 11.711V7a.5.5 0 0 0-1 0v5a.5.5 0 0 0 .25.433l2.598 1.5a.496.496 0 0 0 .682-.183a.5.5 0 0 0-.182-.683M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockEight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6.5a.5.5 0 0 0-.5.5v4.691l-2.706 1.363a.5.5 0 0 0 .45.892l2.98-1.5A.5.5 0 0 0 12.5 12V7a.5.5 0 0 0-.5-.5M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockFive(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 11.874V7a.5.5 0 0 0-1 0v5c0 .082.02.164.06.236l1.5 2.8a.5.5 0 0 0 .676.203h.001a.5.5 0 0 0 .203-.677zM12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockNine(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6a.5.5 0 0 0-.5.5v5h-3a.5.5 0 0 0 0 1H12a.5.5 0 0 0 .5-.5V6.5A.5.5 0 0 0 12 6m0-4C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSeven(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6.5a.5.5 0 0 0-.5.5v4.874l-1.44 2.688v.001a.5.5 0 1 0 .88.472l1.5-2.799A.502.502 0 0 0 12.5 12V7a.5.5 0 0 0-.5-.5M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTen(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6.5a.5.5 0 0 0-.5.5v4.134l-1.848-1.067a.5.5 0 1 0-.5.866l2.598 1.5A.5.5 0 0 0 12.5 12V7a.5.5 0 0 0-.5-.5M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockThree(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 11.5h-3v-5a.5.5 0 0 0-1 0V12a.5.5 0 0 0 .5.5h3.5a.5.5 0 0 0 0-1M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTwo(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.348 10.067L12.5 11.134V7a.5.5 0 0 0-1 0v5a.5.5 0 0 0 .75.433l2.598-1.5a.5.5 0 1 0-.5-.866M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Columns(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-10 19H3V3h8.5zm9.5 0h-8.5V3H21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDots(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11.25a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5m-3 0a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5m6 0a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5m4.415-5.96C15.71 1.195 9.385.88 5.29 4.584C1.195 8.289.88 14.614 4.584 18.709l-2.438 2.437A.5.5 0 0 0 2.5 22H12a10 10 0 0 0 6.709-2.585c4.096-3.705 4.412-10.03.706-14.125M12 21H3.707l1.929-1.929a.5.5 0 0 0 0-.707a8.999 8.999 0 0 1 6.362-15.362A8.999 8.999 0 0 1 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 16h-5a.5.5 0 0 0 0 1H7v4.5a.5.5 0 1 0 1 0v-5a.5.5 0 0 0-.5-.5m9-8h5a.5.5 0 0 0 0-1H17V2.5a.5.5 0 0 0-1 0v5a.5.5 0 0 0 .5.5m-9-6a.5.5 0 0 0-.5.5V7H2.5a.5.5 0 0 0 0 1h5a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5m14 14h-5a.5.5 0 0 0-.5.5v5a.5.5 0 1 0 1 0V17h4.5a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 3a.5.5 0 0 0-.5.5v10a2.502 2.502 0 0 1-2.5 2.5H6.707l3.646-3.646a.5.5 0 0 0-.707-.707l-4.5 4.5a.499.499 0 0 0-.145.35L5 16.5a.5.5 0 0 0 .146.354l4.5 4.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708L6.708 17H15.5a3.504 3.504 0 0 0 3.5-3.5v-10a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownRight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.961 16.309a.5.5 0 0 0-.108-.163l-4.5-4.5a.5.5 0 0 0-.707.708L17.293 16H8.5A2.502 2.502 0 0 1 6 13.5v-10a.5.5 0 0 0-1 0v10A3.504 3.504 0 0 0 8.5 17h8.793l-3.647 3.646a.5.5 0 1 0 .708.708l4.5-4.5a.499.499 0 0 0 .146-.352V16.5a.5.5 0 0 0-.039-.191"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLeftDown(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 5h-9.111A4.394 4.394 0 0 0 7 9.389v8.904l-4.147-4.146a.5.5 0 0 0-.707.707l5 5A.5.5 0 0 0 7.5 20c.011 0 .02-.005.03-.006a.498.498 0 0 0 .163-.033a.5.5 0 0 0 .162-.109l4.998-4.998a.5.5 0 0 0-.707-.708L8 18.293V9.389A3.393 3.393 0 0 1 11.389 6H20.5a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerRightDown(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.853 14.159a.5.5 0 0 0-.707-.013L17 18.293V9.389A4.394 4.394 0 0 0 12.611 5H3.5a.5.5 0 0 0 0 1h9.111A3.393 3.393 0 0 1 16 9.389v8.904l-4.147-4.146a.5.5 0 0 0-.707.707l5 5A.5.5 0 0 0 16.5 20c.011 0 .02-.005.03-.006a.498.498 0 0 0 .163-.033a.5.5 0 0 0 .162-.109l4.998-4.998a.5.5 0 0 0 0-.695"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.111 7.5H5.71l4.646-4.646a.5.5 0 0 0-.707-.707l-5.5 5.5a.5.5 0 0 0 0 .707l5.5 5.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708L5.71 8.5h9.402a3.393 3.393 0 0 1 3.389 3.389V21.5a.5.5 0 1 0 1 0v-9.612a4.394 4.394 0 0 0-4.39-4.388"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpRight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.852 7.647l-5.5-5.5a.5.5 0 0 0-.707.707L18.29 7.5H8.889A4.394 4.394 0 0 0 4.5 11.889V21.5a.5.5 0 1 0 1 0v-9.611A3.393 3.393 0 0 1 8.889 8.5h9.402l-4.646 4.646a.5.5 0 1 0 .707.708l5.5-5.5a.5.5 0 0 0 0-.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CovidNineteen(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.7 4.7c.9-.4 1.8-.6 2.8-.7v2.5c0 .3.2.5.5.5s.5-.2.5-.5V4c1.8.1 3.5.8 4.8 2l-.3.3c-.1.1-.1.2-.1.4c0 .3.2.5.5.5c.1 0 .3-.1.4-.1l.2-.4c1.1 1.3 1.9 3 2 4.8h-2.5c-.3 0-.5.2-.5.5s.2.5.5.5H20c-.1 1-.3 1.9-.7 2.8c-.1.3 0 .5.2.7h.2c.2 0 .4-.1.5-.3c.5-1 .7-2.1.8-3.2h1.5c.3 0 .5-.2.5-.5s-.2-.5-.5-.5H21c-.1-2.1-1-4-2.3-5.5l1-1c.2-.2.2-.5 0-.7c-.2-.2-.5-.2-.7 0l-1 1C16.5 4 14.6 3.1 12.5 3V1.5c0-.3-.2-.5-.5-.5s-.5.2-.5.5V3c-1.1.1-2.2.3-3.2.8c-.3.1-.4.4-.3.7c.2.2.5.3.7.2m.3 8.8c-.8 0-1.5.7-1.5 1.5s.7 1.5 1.5 1.5s1.5-.7 1.5-1.5s-.7-1.5-1.5-1.5m0 2c-.3 0-.5-.2-.5-.5s.2-.5.5-.5s.5.2.5.5s-.2.5-.5.5M1.9 1.1c-.2-.1-.6-.1-.7 0c-.2.2-.2.6-.1.8L5.3 6C3.9 7.5 3.1 9.5 3 11.5H1.5c-.3 0-.5.2-.5.5s.2.5.5.5H3c.1 2.1 1 4 2.3 5.5l-1 1c-.1.1-.1.2-.1.4c0 .3.2.5.5.5c.1 0 .3-.1.4-.1l1-1c1.5 1.3 3.4 2.2 5.5 2.3v1.5c0 .3.2.5.5.5s.5-.2.5-.5V21c2-.1 4-.9 5.5-2.3l4.2 4.2c.1.1.2.1.4.1c.1 0 .3-.1.4-.1c.2-.2.2-.5 0-.7zM9 9.7l1.4 1.4c-.2.3-.5.5-.8.5c-.6 0-1-.4-1-1c-.1-.4.1-.8.4-.9M12.5 20v-3.5c0-.3-.2-.5-.5-.5s-.5.2-.5.5V20c-1.7-.1-3.4-.8-4.8-2l.3-.3c.2-.2.2-.5 0-.7c-.2-.2-.5-.2-.7 0l-.3.3c-1.2-1.4-1.9-3.1-2-4.8h2.5c.3 0 .5-.2.5-.5s-.2-.5-.5-.5H4c.1-1.7.8-3.4 2-4.8L8.3 9c-.5.4-.8.9-.8 1.5c0 1.1.9 2 2 2c.6 0 1.2-.3 1.5-.8l6.2 6.2c-1.3 1.3-3 2-4.7 2.1m1-11c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5s-.7-1.5-1.5-1.5s-1.5.7-1.5 1.5m2 0c0 .3-.2.5-.5.5s-.5-.2-.5-.5s.2-.5.5-.5s.5.2.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreateDashboard(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 13h-7c-.3 0-.5.2-.5.5v7c0 .3.2.5.5.5h7c.3 0 .5-.2.5-.5v-7c0-.3-.2-.5-.5-.5m-.5 7H4v-6h6zm.5-17h-7c-.3 0-.5.2-.5.5v7c0 .3.2.5.5.5h7c.3 0 .5-.2.5-.5v-7c0-.3-.2-.5-.5-.5m-.5 7H4V4h6zm10.5-7h-7c-.3 0-.5.2-.5.5v7c0 .3.2.5.5.5h7c.3 0 .5-.2.5-.5v-7c0-.3-.2-.5-.5-.5m-.5 7h-6V4h6zm.5 6.5h-3v-3c0-.3-.2-.5-.5-.5s-.5.2-.5.5v3h-3c-.3 0-.5.2-.5.5s.2.5.5.5h3v3c0 .3.2.5.5.5s.5-.2.5-.5v-3h3c.3 0 .5-.2.5-.5s-.2-.5-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopAltSlash(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.8 4h10.7c.8 0 1.5.7 1.5 1.5V12h-3.2c-.3 0-.5.2-.5.5s.2.5.5.5H20v1.5c0 .4-.1.7-.3 1c-.1.1-.1.2-.1.3c0 .3.2.5.5.5c.1 0 .3-.1.4-.2c.4-.4.6-1 .6-1.6v-9C21 4.1 19.9 3 18.5 3H7.8c-.3 0-.5.2-.5.5s.2.5.5.5m14.1 17.1l-5-5l-4-4l-8.4-8.3l-.1-.1h-.1L2.9 2.1c-.2-.1-.5-.1-.7 0c-.2.2-.2.6-.1.8l1.3 1.3c-.3.3-.4.8-.4 1.3v9C3 15.9 4.1 17 5.5 17H9v3H5.5c-.3 0-.5.2-.5.5s.2.5.5.5h13c.3 0 .5-.2.5-.5s-.2-.5-.5-.5H15v-3h1.3l4.9 4.9c.1.1.2.1.4.1c.1 0 .3-.1.4-.1c0-.2 0-.6-.1-.8M4 5.5c0-.2.1-.4.2-.6l7.1 7.1H4zM14 20h-4v-3h4zm-8.5-4c-.8 0-1.5-.7-1.5-1.5V13h8.3l3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dialpad(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 10H3a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4A.5.5 0 0 0 7 10m-.5 4h-3v-3h3zM7 3H3a.5.5 0 0 0-.5.5v4A.5.5 0 0 0 3 8h4a.5.5 0 0 0 .5-.5v-4A.5.5 0 0 0 7 3m-.5 4h-3V4h3zM14 3h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4A.5.5 0 0 0 14 3m-.5 4h-3V4h3zM21 3h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4A.5.5 0 0 0 21 3m-.5 4h-3V4h3zM14 17h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5m-.5 4h-3v-3h3zM21 10h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5m-.5 4h-3v-3h3zM14 10h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5m-.5 4h-3v-3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Direction(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.854 9.854L12 6.707l3.146 3.147a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707l-3.5-3.5a.5.5 0 0 0-.706 0l-3.5 3.5a.5.5 0 0 0 .707.707m6.292 4.292L12 17.293l-3.147-3.146a.5.5 0 0 0-.707.707l3.5 3.5a.498.498 0 0 0 .707 0l3.5-3.5a.5.5 0 0 0-.707-.708"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentLayoutCenter(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 12h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5M9 5h6v6H9zm12.5 6h-3a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1m-19-3h3a.5.5 0 0 0 0-1h-3a.5.5 0 0 0 0 1m16 0h3a.5.5 0 0 0 0-1h-3a.5.5 0 0 0 0 1m3 7h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m-19-3h3a.5.5 0 0 0 0-1h-3a.5.5 0 0 0 0 1m11 7h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentLayoutLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 12h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5M3 5h6v6H3zm9.5 3h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1m9 7h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m-8 4h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1m8-8h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentLayoutRight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 12h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1m0-4h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1m11 11h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1m8-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1m0-11h-7a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5m-.5 7h-6V5h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.162 13.655a.5.5 0 0 0-.007.707l3.491 3.492a.498.498 0 0 0 .708 0l3.49-3.492a.5.5 0 0 0-.706-.707L12.5 16.293V2.5a.5.5 0 0 0-1 0v13.793l-2.638-2.638a.5.5 0 0 0-.7 0M18 9h-1.5a.5.5 0 0 0 0 1H18a2.003 2.003 0 0 1 2 2v7a2.003 2.003 0 0 1-2 2H6a2.003 2.003 0 0 1-2-2v-7a2.003 2.003 0 0 1 2-2h2.5a.5.5 0 0 0 0-1H6a3.003 3.003 0 0 0-3 3v7a3.003 3.003 0 0 0 3 3h12a3.003 3.003 0 0 0 3-3v-7a3.003 3.003 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.3 12.5l-3.6-3.3l3.6-2.9c.1 0 .1-.1.1-.1c.1-.2 0-.5-.2-.7l-5-2.7c-.2-.1-.4-.1-.5 0L12 5.6L8.3 2.8c-.2-.1-.4-.1-.5 0l-5 2.7c-.1 0-.1.1-.1.1c-.3.1-.2.5 0 .6l3.6 2.9l-3.6 3.3l-.1.1c-.1.2 0 .5.2.7l3.7 2v2.5c0 .2.1.3.2.4l5 3c.1 0 .2.1.3.1c.1 0 .2 0 .3-.1l5-3c.2-.1.2-.3.2-.4v-2.5l3.7-2l.1-.1c.3-.1.2-.4 0-.6M16 3.7l4.1 2.2L17 8.6l-4.1-2.4zm.1 5.5L12 11.9L7.9 9.2L12 6.8zM3.9 5.9L8 3.7l3.2 2.4L7 8.6zm0 6.8l3.2-3l4.1 2.7L8 15zm12.6 4.9L12 20.3l-4.5-2.7v-1.7l.3.1H8c.1 0 .2 0 .3-.1L12 13l3.7 2.9c.1.1.2.1.3.1c.1 0 .2 0 .2-.1l.3-.1zM16 15l-3.2-2.5l4.1-2.7l3.2 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisH(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m7-3a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m7-3a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisV(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0-3a1 1 0 1 1-.001 2.001A1 1 0 0 1 12 4m0 6a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircle(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14.375a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25M12 13a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5m0-11C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationOctagon(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5m9.854-5.288l-5.566-5.566A.5.5 0 0 0 15.935 2h-7.87a.5.5 0 0 0-.353.146L2.146 7.712A.5.5 0 0 0 2 8.065v7.87a.5.5 0 0 0 .146.353l5.566 5.566a.5.5 0 0 0 .353.146h7.87a.5.5 0 0 0 .353-.146l5.566-5.566a.5.5 0 0 0 .146-.353v-7.87a.5.5 0 0 0-.146-.353M21 15.728L15.728 21H8.272L3 15.728V8.272L8.272 3h7.456L21 8.272zm-9-1.353a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangle(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15.875a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25M12 14a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5m10.713 3.547L14.564 3.479a2.964 2.964 0 0 0-5.128 0l-8.15 14.07A2.965 2.965 0 0 0 3.852 22h16.294a2.966 2.966 0 0 0 2.567-4.453M20.146 21H3.852a1.965 1.965 0 0 1-1.7-2.95L10.3 3.98a1.943 1.943 0 0 1 1.699-.979a1.944 1.944 0 0 1 1.7.98l8.148 14.068A1.966 1.966 0 0 1 20.146 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookF(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.99 1.596a27.982 27.982 0 0 0-3.037-.156C11.59 1.44 9.5 3.582 9.5 7.03v2.341H6.675a.5.5 0 0 0-.5.5v3.85a.5.5 0 0 0 .5.5H9.5v7.72a.5.5 0 0 0 .5.5h3.978a.5.5 0 0 0 .5-.5v-7.72h2.816a.5.5 0 0 0 .496-.435l.497-3.85a.5.5 0 0 0-.496-.565h-3.313V7.412c0-.97.195-1.375 1.408-1.375h2.039a.5.5 0 0 0 .5-.5V2.092a.5.5 0 0 0-.435-.496m-.565 3.44l-1.54.001c-2.157 0-2.407 1.356-2.407 2.375v2.46a.5.5 0 0 0 .499.5h3.246l-.369 2.85h-2.876a.5.5 0 0 0-.5.5v7.718H10.5v-7.718a.5.5 0 0 0-.5-.5H7.176v-2.85H10a.5.5 0 0 0 .5-.5V7.03c0-2.874 1.665-4.59 4.453-4.59c1.009 0 1.92.055 2.472.103z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookMessengerAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.117 8.906l-2.589 3.086l-3.715-2.281a.499.499 0 0 0-.644.104l-3.052 3.636a.5.5 0 0 0 .766.643l2.774-3.306l3.715 2.281c.211.13.485.085.645-.104l2.866-3.416a.5.5 0 0 0-.766-.643M12 1C5.715 1 .975 5.594.975 11.686a10.433 10.433 0 0 0 3.471 7.905c.071.06.114.149.118.242l.058 1.867a1.343 1.343 0 0 0 1.883 1.187l2.088-.92a.33.33 0 0 1 .226-.018c1.037.283 2.107.425 3.181.422c6.285 0 11.025-4.594 11.025-10.685S18.285 1 12 1m0 20.371a10.959 10.959 0 0 1-2.916-.387a1.36 1.36 0 0 0-.894.068l-2.086.919a.346.346 0 0 1-.324-.026a.336.336 0 0 1-.158-.276l-.058-1.871a1.342 1.342 0 0 0-.45-.952a9.446 9.446 0 0 1-3.14-7.16C1.975 6.164 6.285 2 12 2s10.025 4.164 10.025 9.686S17.715 21.37 12 21.37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Favorite(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.919 10.127a1 1 0 0 0-.845-1.136l-5.651-.826l-2.526-5.147a1.037 1.037 0 0 0-1.795.001L8.577 8.165l-5.651.826a1 1 0 0 0-.556 1.704l4.093 4.013l-.966 5.664a1.002 1.002 0 0 0 1.453 1.052l5.05-2.67l5.049 2.669a1 1 0 0 0 1.454-1.05l-.966-5.665l4.094-4.014a1 1 0 0 0 .288-.567m-5.269 4.05a.502.502 0 0 0-.143.441l1.01 5.921l-5.284-2.793a.505.505 0 0 0-.466 0L6.483 20.54l1.01-5.922a.502.502 0 0 0-.143-.441L3.07 9.98l5.912-.864a.503.503 0 0 0 .377-.275L12 3.46l2.64 5.382a.503.503 0 0 0 .378.275l5.913.863z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipH(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.815 3.383L12 3.2l.185.182a.499.499 0 0 0 .707-.006a.5.5 0 0 0-.006-.707l-.535-.526a.5.5 0 0 0-.702 0l-.535.526l-.008.008a.5.5 0 1 0 .71.705m3.282 2.864a.499.499 0 1 0 .7-.713l-.793-.781a.5.5 0 0 0-.701.713zM11.155 9h1.69a.5.5 0 0 0 0-1h-1.69a.5.5 0 0 0 0 1m6.195 0h.75a.5.5 0 0 0 .35-.857l-.534-.526a.5.5 0 0 0-.839.464a.5.5 0 0 0 .273.919M8.553 6.39a.499.499 0 0 0 .35-.143l.794-.781a.5.5 0 0 0-.7-.713l-.795.781a.5.5 0 0 0 .35.857M21.5 11.5h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1M17.929 15H6.07a.5.5 0 0 0-.355.852l5.929 6a.5.5 0 0 0 .71 0l5.93-6a.5.5 0 0 0-.356-.852M12 20.79L7.269 16h9.462zM5.9 9h.75a.5.5 0 0 0 .273-.919a.5.5 0 0 0-.839-.464l-.534.527A.5.5 0 0 0 5.9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipHalt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 8.321a.5.5 0 0 0 .5-.5v-.366l.354-.095a.5.5 0 0 0-.26-.966l-.724.194A.5.5 0 0 0 3 7.07v.75a.5.5 0 0 0 .5.5M9.934 10h-1.55a.5.5 0 0 0 0 1h1.55a.5.5 0 0 0 0-1M8.372 6.262a.5.5 0 0 0 .13-.018l1.555-.418h.003a.5.5 0 1 0-.263-.965l-1.555.418a.5.5 0 0 0 .13.983M14.067 11h1.55a.5.5 0 0 0 0-1h-1.55a.5.5 0 0 0 0 1m.007-6.272a.502.502 0 0 0 .13-.018l1.555-.418h.002a.5.5 0 0 0-.262-.965l-1.555.418a.5.5 0 0 0 .13.983M20.5 13h-17a.5.5 0 0 0-.5.5v3.75a.5.5 0 0 0 .379.485l17 4.25c.04.01.08.015.121.015h.001a.5.5 0 0 0 .499-.5v-8a.5.5 0 0 0-.5-.5m-.5 7.86l-16-4V14h16zm.983-18.489a.5.5 0 0 0-.613-.354l-.724.194h-.002a.5.5 0 1 0 .261.967L20 3.15v.099a.5.5 0 0 0 1 0V2.5a.5.5 0 0 0-.017-.129M20.5 5.487a.5.5 0 0 0-.5.5v1.027a.5.5 0 1 0 1 0V5.987a.5.5 0 0 0-.5-.5m0 3.763a.5.5 0 0 0-.5.5V10h-.25a.5.5 0 0 0 0 1h.75a.5.5 0 0 0 .5-.5v-.75a.5.5 0 0 0-.5-.5M3.5 11h.75a.5.5 0 0 0 0-1H4v-.25a.5.5 0 0 0-1 0v.75a.5.5 0 0 0 .5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipV(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.466 14.302l-.009-.009a.5.5 0 1 0-.704.71l.781.794a.5.5 0 1 0 .713-.701zm.068-6.1l-.781.794a.5.5 0 1 0 .713.701l.781-.794a.5.5 0 1 0-.713-.701M3.201 12l.182-.185a.5.5 0 0 0-.713-.7l-.526.534a.5.5 0 0 0 0 .702l.526.535a.498.498 0 0 0 .707.006a.5.5 0 0 0 .006-.707zm4.88 5.077a.5.5 0 0 0-.464.839l.527.534A.5.5 0 0 0 9 18.1v-.75a.5.5 0 0 0-.919-.273m13.775-5.428l-6-6.1A.5.5 0 0 0 15 5.9v12.2a.5.5 0 0 0 .856.35l6-6.1a.5.5 0 0 0 0-.7M16 16.878V7.122L20.799 12zm-7.5-6.223a.5.5 0 0 0-.5.5v1.69a.5.5 0 1 0 1 0v-1.69a.5.5 0 0 0-.5-.5M12 2a.5.5 0 0 0-.5.5v19a.5.5 0 1 0 1 0v-19A.5.5 0 0 0 12 2M8.143 5.55l-.526.534a.5.5 0 0 0 .464.839A.5.5 0 0 0 9 6.65V5.9a.5.5 0 0 0-.857-.351"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipValt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.826 10.057l.418-1.555v-.003a.5.5 0 1 0-.965-.257l-.418 1.555a.5.5 0 0 0 .965.26M3.25 20h-.099l.027-.095a.502.502 0 0 0-.354-.613a.508.508 0 0 0-.613.354l-.194.724A.5.5 0 0 0 2.5 21h.75a.5.5 0 0 0 0-1M7.821 3h-.75a.5.5 0 0 0-.483.37l-.194.725a.502.502 0 0 0 .966.26L7.455 4h.366a.5.5 0 0 0 0-1M4.357 13.592a.5.5 0 0 0-.612.352L3.327 15.5a.5.5 0 0 0 .965.26l.418-1.555a.5.5 0 0 0-.353-.612M7.013 20H5.986a.5.5 0 0 0 0 1h1.027a.5.5 0 0 0 0-1m14.97.37l-4.572-17A.499.499 0 0 0 16.93 3H13.5a.5.5 0 0 0-.5.5v17a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .482-.63M14 20V4h2.545l4.303 16zM10.5 3h-.75a.5.5 0 0 0 0 1H10v.25a.5.5 0 0 0 1 0V3.5a.5.5 0 0 0-.5-.5m0 16.25a.5.5 0 0 0-.5.5V20h-.25a.5.5 0 0 0 0 1h.75a.5.5 0 0 0 .5-.5v-.75a.5.5 0 0 0-.5-.5m0-5.683a.5.5 0 0 0-.5.5v1.55a.5.5 0 1 0 1 0v-1.55a.5.5 0 0 0-.5-.5m0-5.683a.5.5 0 0 0-.5.5v1.55a.5.5 0 1 0 1 0v-1.55a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.653 20.865a.501.501 0 0 0-.59-.492c-1.59.295-3.608.299-4.194-1.347a5.622 5.622 0 0 0-1.679-2.249a1.46 1.46 0 0 1-.408-.31a.5.5 0 0 0-.491-.42h-.002a.5.5 0 0 0-.5.499c-.003.57.71.997.922 1.11c.53.475.947 1.064 1.222 1.721c.348.98 1.41 2.495 4.722 2.072c.003.365.008.553.012.697l.005.294a.5.5 0 0 0 1 0l-.005-.322a38.783 38.783 0 0 1-.014-1.253M20.921 5.224a7.357 7.357 0 0 0-.096-.119c.067-.21.12-.427.158-.645a6.033 6.033 0 0 0-.397-3.17a.5.5 0 0 0-.309-.29c-.141-.047-1.433-.395-4.13 1.382a13.973 13.973 0 0 0-6.884 0C6.567.626 5.283.955 5.14.997a.503.503 0 0 0-.316.292a6.038 6.038 0 0 0-.394 3.219c.038.2.087.397.145.592a3.584 3.584 0 0 0-.1.128A5.946 5.946 0 0 0 3.2 9.002c-.002.31.012.62.043.929c.34 4.664 3.349 5.962 5.947 6.405a3.807 3.807 0 0 0-.376.967a.5.5 0 0 0 .971.238c.093-.46.325-.88.665-1.202a.5.5 0 0 0-.272-.874C7.422 15.152 4.56 14.24 4.24 9.84a7.683 7.683 0 0 1-.039-.838a4.982 4.982 0 0 1 1.07-3.168c.076-.102.16-.194.243-.286a.501.501 0 0 0 .096-.516a4.05 4.05 0 0 1-.194-.695a4.955 4.955 0 0 1 .232-2.39a6.71 6.71 0 0 1 3.248 1.39a.497.497 0 0 0 .414.067a12.973 12.973 0 0 1 6.793 0a.5.5 0 0 0 .415-.067a6.555 6.555 0 0 1 3.242-1.398a4.94 4.94 0 0 1 .237 2.356a3.877 3.877 0 0 1-.206.737a.501.501 0 0 0 .097.516c.088.097.175.205.253.302a4.913 4.913 0 0 1 1.07 3.152c.002.286-.012.572-.042.856c-.317 4.381-3.19 5.292-5.957 5.607a.5.5 0 0 0-.273.874c.345.325.576.752.659 1.219c.085.329.125.668.118 1.008v2.46c-.01.675-.01 1.182-.01 1.414a.5.5 0 0 0 1 0c0-.23 0-.731.01-1.407v-2.467a4.633 4.633 0 0 0-.15-1.255a3.653 3.653 0 0 0-.367-.975c2.609-.442 5.63-1.74 5.966-6.385c.033-.315.049-.632.046-.949a5.892 5.892 0 0 0-1.29-3.778"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gold(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 11h9.1c.3 0 .5-.3.4-.6l-1-6c0-.2-.3-.4-.5-.4h-7c-.2 0-.5.2-.5.4l-1 6v.1c0 .3.2.5.5.5m1.4-6h6.2l.8 5H8.1zM22 13.4c0-.2-.2-.4-.5-.4h-7c-.2 0-.5.2-.5.4l-1 6v.1c0 .3.2.5.5.5h9.1c.3 0 .5-.3.4-.6zM14.1 19l.8-5h6.2l.8 5zm-4.6-6h-7c-.2 0-.5.2-.5.4l-1 6v.1c0 .3.2.5.5.5h9.1c.3 0 .5-.3.4-.6l-1-6c0-.2-.3-.4-.5-.4m-7.4 6l.8-5h6.2l.8 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.132 10.033a.5.5 0 0 0-.492-.41h-9.418a.5.5 0 0 0-.5.498v3.869a.5.5 0 0 0 .5.5h4.735a4.136 4.136 0 0 1-1.627 2.103a5.55 5.55 0 0 1-3.108.87A5.434 5.434 0 0 1 7.1 13.682v-.002a5.416 5.416 0 0 1 0-3.48v-.002a5.434 5.434 0 0 1 5.12-3.781a4.93 4.93 0 0 1 3.48 1.357a.5.5 0 0 0 .7-.007l2.868-2.869a.5.5 0 0 0-.013-.72a10.135 10.135 0 0 0-7.032-2.738A10.451 10.451 0 0 0 2.84 7.225a10.51 10.51 0 0 0 0 9.43a10.453 10.453 0 0 0 9.383 5.785a10.034 10.034 0 0 0 6.952-2.552l.005-.002a10.296 10.296 0 0 0 3.143-7.719c0-.716-.064-1.43-.19-2.134m-9.91-7.593a9.153 9.153 0 0 1 5.962 2.127l-2.16 2.16a5.937 5.937 0 0 0-3.802-1.31a6.407 6.407 0 0 0-5.817 3.818l-2.48-1.924a9.453 9.453 0 0 1 8.297-4.87m-9.5 9.5a9.427 9.427 0 0 1 .753-3.71l2.573 1.994a6.387 6.387 0 0 0 0 3.431L3.476 15.65a9.552 9.552 0 0 1-.754-3.71m9.5 9.5a9.454 9.454 0 0 1-8.298-4.871l2.481-1.924a6.406 6.406 0 0 0 5.817 3.818a6.67 6.67 0 0 0 3.355-.847l.11.085l2.363 1.836a9.19 9.19 0 0 1-5.828 1.903m6.582-2.584l-2.378-1.846a5.092 5.092 0 0 0 1.669-2.929a.5.5 0 0 0-.491-.59h-4.882v-2.869h8.492c.072.512.108 1.028.108 1.545a9.42 9.42 0 0 1-2.518 6.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDriveAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.732 15.138c-.002-.034-.011-.065-.02-.098c-.009-.033-.017-.064-.031-.094c-.005-.01-.005-.02-.01-.029L15.513 4.25A.5.5 0 0 0 15.08 4H8.92a.5.5 0 0 0-.249.067a.488.488 0 0 0-.165.166c-.005.007-.013.01-.018.017L2.33 14.917a.497.497 0 0 0 0 .5l3.08 5.333c.007.013.022.018.03.03a.48.48 0 0 0 .123.123c.019.013.034.026.054.036c.069.036.143.06.226.061H18.16a.5.5 0 0 0 .433-.25l3.079-5.333c.007-.012.005-.028.01-.04a.487.487 0 0 0 .046-.17c.001-.013.01-.025.01-.04c0-.01-.004-.019-.005-.03M14.79 5l5.581 9.667H15.37L9.787 5zm-.577 9.667H9.786L12 10.832zm-10.875.5L8.92 5.5l2.502 4.333l-5.58 9.667zM17.87 20H6.708l2.502-4.333h11.162z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleHangoutsAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.992 8a2 2 0 0 0 0 4c.358-.002.71-.101 1.016-.287V12a1 1 0 0 1-1 1a.5.5 0 0 0 0 1a2.003 2.003 0 0 0 2-2v-2a.49.49 0 0 0-.032-.156A1.997 1.997 0 0 0 14.992 8m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2M13.34 1.082A9.664 9.664 0 0 0 12 1C6.628.985 2.262 5.328 2.248 10.7a9.726 9.726 0 0 0 9.252 9.74v2.06a.5.5 0 0 0 .5.5c.04 0 .08-.005.12-.015a12.625 12.625 0 0 0 9.55-11.056c.695-5.296-3.035-10.152-8.33-10.847m7.337 10.74A11.628 11.628 0 0 1 12.5 21.843v-1.89a.5.5 0 0 0-.5-.5a8.788 8.788 0 1 1 8.676-7.633M8.992 8a2 2 0 0 0 0 4c.358-.002.71-.101 1.016-.287V12a1 1 0 0 1-1 1a.5.5 0 0 0 0 1a2.003 2.003 0 0 0 2-2v-2a.49.49 0 0 0-.032-.156A1.997 1.997 0 0 0 8.992 8m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlay(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.9 11.1c-.4-.2-4-2.3-4-2.3L4.6 1.7c-.2-.2-.6-.2-.9-.1c0 0-.1 0-.1.1c-.4.1-.6.5-.6.9v19c0 .3.2.7.4.8h.1c.1 0 .2.1.3.1c.2 0 .4-.1.6-.2c.4-.2 12.4-7.2 12.4-7.2l4-2.3c.4-.2.6-.5.6-.9c.1-.3-.1-.7-.5-.8m-5.1-1.8l-2 2l-7.4-7.5zM4 21V2.9l9.2 9.1zm2.4-1l7.4-7.4l2 1.9c-1.6 1.1-6.4 3.8-9.4 5.5m14-8l-3.7 2.1l-2.1-2.1l2.2-2.1c.8.4 3 1.7 3.6 2.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-10 19H3v-5.5h8.5zm0-6.5H3v-5h8.5zM21 21h-8.5v-5.5H21zm0-6.5h-8.5v-5H21zm0-6H3V3h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grids(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-13 19H3V3h5.5zm6 0h-5V3h5zm6.5 0h-5.5V3H21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GripHorizontalLine(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 11h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m19 3h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSide(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 9.4c0-4-3.2-7.3-7.2-7.4c-4.1-.1-7.5 3.2-7.6 7.2l-2.2 4.5v.2c0 .3.2.5.5.5h2V17c0 1.1.9 2 2 2h1v2.5c0 .3.2.5.5.5s.5-.2.5-.5V19h.5c.3 0 .5-.2.5-.5s-.2-.5-.5-.5H8c-.6 0-1-.4-1-1v-3c0-.3-.2-.5-.5-.5H4.8l1.9-3.9c0-.1.1-.1 0-.2C6.7 5.9 9.5 3 13 3s6.4 2.8 6.5 6.3l-2.1 7.9v.3l1.1 4.2c.1.2.3.4.5.4h.1c.3-.1.4-.3.4-.6l-1-4.1l2.1-7.8c-.1-.1-.1-.2-.1-.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSideCough(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 20.4c-.3 0-.6.3-.6.6s.3.6.6.6s.6-.3.6-.6s-.3-.6-.6-.6M23 9c0-3.8-3.1-6.9-6.9-7c-3.8-.1-7 3-7.1 6.9l-2.1 4.4v.2c0 .3.2.5.5.5H9v2c0 1.1.9 2 2 2h1v2.5c0 .3.2.5.5.5s.5-.2.5-.5V18h.5c.3 0 .5-.2.5-.5s-.2-.5-.5-.5H11c-.6 0-1-.4-1-1v-2.5c0-.3-.2-.5-.5-.5H8.2L10 9.2V9c0-3.3 2.6-6 5.9-6c3.3 0 6 2.6 6.1 5.9l-2 7.4v.2l1 4c.1.2.3.4.5.4h.1c.3-.1.4-.3.4-.6l-1-3.9zc0 .1 0 0 0 0M2 17.4c-.3 0-.6.3-.6.6s.3.6.6.6s.6-.3.6-.6s-.3-.6-.6-.6m4-1c-.3 0-.6.3-.6.6s.3.6.6.6s.6-.3.6-.6s-.3-.6-.6-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSideMask(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.2 2.5H13c-3.9 0-7 3.1-7 7v.1l-1.9 3.1c-.1.1-.1.2-.1.3v.2l1.5 3.9c.4.8 1.3 1.4 2.2 1.4H9V21c0 .3.2.5.5.5s.5-.2.5-.5v-2.5h2.7l4.5-1.5l-.2.4v.3l1 3.5c.1.2.3.4.5.4h.1c.3-.1.4-.4.3-.6l-1-3.3l2-7.4v-.5c.1-3.9-2.8-7.1-6.7-7.3m-1.2 15H7.7c-.6 0-1.1-.3-1.3-.8l-1.2-3.2H12zm5.4-1.7L13 17.3v-3.9l5.6-1.9zm1.5-5.6l-6.5 2.3h-7L6.9 10c.1-.1.1-.2.1-.3v-.2c0-1.6.7-3.2 1.8-4.3c1.2-1.1 2.7-1.8 4.4-1.7c3.3.2 5.9 3 5.8 6.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.727 3.18C12.31.81 6.915 2.103 4 6V3.5a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 0-1H4.522a8.954 8.954 0 0 1 7.411-4A8.967 8.967 0 1 1 3 12c0-.16 0-.312.009-.472A.5.5 0 0 0 2.52 11c-.27-.01-.5.2-.51.472C2 11.652 2 11.82 2 12c.006 5.52 4.48 9.994 10 10a10.005 10.005 0 0 0 8.81-5.273c2.614-4.868.786-10.933-4.083-13.547M12 8a.5.5 0 0 0-.5.5V12a.5.5 0 0 0 .5.5h2.5a.5.5 0 0 0 0-1h-2v-3A.5.5 0 0 0 12 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistoryAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.728 3.18C12.31.81 6.915 2.105 4 6V3.5a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 0-1H4.523a8.987 8.987 0 0 1 7.45-4A9 9 0 0 1 12 21a.5.5 0 0 0 0 1a10.005 10.005 0 0 0 8.81-5.272c2.614-4.868.786-10.934-4.082-13.547M12 8a.5.5 0 0 0-.5.5V12a.5.5 0 0 0 .5.5h2.5a.5.5 0 0 0 0-1h-2v-3A.5.5 0 0 0 12 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalAlignLeft(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 11H16V6.5a.5.5 0 0 0-.5-.5H3V2.5a.5.5 0 0 0-1 0v19a.5.5 0 1 0 1 0V18h18.5a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5M3 7h12v4H3zm18 10H3v-5h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 17h-1c-.3 0-.5.2-.5.5s.2.5.5.5h1c.3 0 .5-.2.5-.5s-.2-.5-.5-.5m-5-4h-1c-.3 0-.5.2-.5.5s.2.5.5.5h1c.3 0 .5-.2.5-.5s-.2-.5-.5-.5m5 0h-1c-.3 0-.5.2-.5.5s.2.5.5.5h1c.3 0 .5-.2.5-.5s-.2-.5-.5-.5m-5 4h-1c-.3 0-.5.2-.5.5s.2.5.5.5h1c.3 0 .5-.2.5-.5s-.2-.5-.5-.5m6-9.5h-1v-1c0-.3-.2-.5-.5-.5s-.5.2-.5.5v1h-1c-.3 0-.5.2-.5.5s.2.5.5.5h1v1c0 .3.2.5.5.5s.5-.2.5-.5v-1h1c.3 0 .5-.2.5-.5s-.2-.5-.5-.5m8-.5H18V2.5c0-.3-.2-.5-.5-.5h-11c-.3 0-.5.2-.5.5V7H2.5c-.3 0-.5.2-.5.5v14c0 .3.2.5.5.5h19c.3 0 .5-.2.5-.5v-14c0-.3-.2-.5-.5-.5M21 21H3V8h3.5c.3 0 .5-.2.5-.5V3h10v4.5c0 .3.2.5.5.5H21zm-3.5-4h-1c-.3 0-.5.2-.5.5s.2.5.5.5h1c.3 0 .5-.2.5-.5s-.2-.5-.5-.5m-1-4c-.3 0-.5.2-.5.5s.2.5.5.5h1c.3 0 .5-.2.5-.5s-.2-.5-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalSquareSign(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 6c-.3 0-.5.2-.5.5v5H9v-5c0-.3-.2-.5-.5-.5s-.5.2-.5.5v11c0 .3.2.5.5.5s.5-.2.5-.5v-5h6v5c0 .3.2.5.5.5s.5-.2.5-.5v-11c0-.3-.2-.5-.5-.5m4-4h-15C3.1 2 2 3.1 2 4.5v15C2 20.9 3.1 22 4.5 22h15c1.4 0 2.5-1.1 2.5-2.5v-15C22 3.1 20.9 2 19.5 2M21 19.5c0 .8-.7 1.5-1.5 1.5h-15c-.8 0-1.5-.7-1.5-1.5v-15C3 3.7 3.7 3 4.5 3h15c.8 0 1.5.7 1.5 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalSymbol(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 7.5c-.3 0-.5.2-.5.5v3.5h-5V8c0-.3-.2-.5-.5-.5s-.5.2-.5.5v8c0 .3.2.5.5.5s.5-.2.5-.5v-3.5h5V16c0 .3.2.5.5.5s.5-.2.5-.5V8c0-.3-.2-.5-.5-.5M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m0 19c-5 0-9-4-9-9s4-9 9-9s9 4 9 9s-4 9-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseUser(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.8 10.6l-2-1.8l-7.5-6.7c-.2-.2-.5-.2-.7 0L4.2 8.8l-2 1.8c-.2.2-.2.5 0 .7c.2.2.5.2.7 0l1.1-1v11.2c0 .3.3.5.5.5h14.9c.3 0 .5-.2.5-.5V10.3l1.2 1.1c.1.1.2.1.3.1c.1 0 .3-.1.4-.2c.3-.2.2-.5 0-.7M8.1 21c.4-1.6 1.8-2.8 3.4-3c2-.2 3.8 1.1 4.3 3zm1.4-6.5c0-1.4 1.1-2.5 2.5-2.5s2.5 1.1 2.5 2.5S13.4 17 12 17s-2.5-1.1-2.5-2.5M19 21h-2.1c-.3-1.6-1.4-3-2.9-3.6c.9-.6 1.6-1.7 1.6-2.9c0-1.9-1.6-3.5-3.5-3.5s-3.5 1.6-3.5 3.5c0 1.2.6 2.3 1.6 2.9c-1.6.6-2.7 1.9-3.1 3.6H5V9.4l7-6.2l7 6.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlThreeAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.87 2.162A.5.5 0 0 0 20.5 2h-17a.5.5 0 0 0-.498.544l1.5 17a.5.5 0 0 0 .36.437l6.99 2a.51.51 0 0 0 .275 0l7.01-2a.5.5 0 0 0 .361-.437l1.5-17a.5.5 0 0 0-.129-.382m-2.338 16.951L11.99 20.98l-6.522-1.866L4.046 3h15.908zM7.96 7.5h7.534l-.266 3H10a.5.5 0 0 0 0 1h5.139l-.345 3.883l-2.794.635l-2.794-.635l-.128-1.428a.513.513 0 0 0-.543-.453a.5.5 0 0 0-.453.543l.16 1.79a.5.5 0 0 0 .388.442l3.26.74a.487.487 0 0 0 .22 0l3.26-.74a.5.5 0 0 0 .388-.443l.426-4.803c0-.01.006-.02.006-.031l-.002-.012l.35-3.944a.5.5 0 0 0-.498-.544H7.96a.5.5 0 1 0 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageV(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 2h-15A2.502 2.502 0 0 0 2 4.5v15A2.502 2.502 0 0 0 4.5 22h15a2.502 2.502 0 0 0 2.5-2.5v-15A2.502 2.502 0 0 0 19.5 2m-15 19A1.5 1.5 0 0 1 3 19.5v-5.225l3.763-3.762a1.753 1.753 0 0 1 2.474 0l10.468 10.466c-.068.01-.135.02-.205.021zM21 19.5c0 .378-.145.72-.377.984l-6.916-6.916l1.056-1.055a1.753 1.753 0 0 1 2.474 0L21 16.275zm0-4.639l-3.056-3.055a2.753 2.753 0 0 0-3.888 0L13 12.86L9.944 9.806a2.815 2.815 0 0 0-3.888 0L3 12.86V4.5A1.5 1.5 0 0 1 4.5 3h15A1.5 1.5 0 0 1 21 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntercomAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 14.5a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-1 0v8a.5.5 0 0 0 .5.5m4 0a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-1 0v8a.5.5 0 0 0 .5.5M20 1H4a3.003 3.003 0 0 0-3 3v16a3.003 3.003 0 0 0 3 3h16a3.003 3.003 0 0 0 3-3V4a3.003 3.003 0 0 0-3-3m2 19a2.003 2.003 0 0 1-2 2H4a2.003 2.003 0 0 1-2-2V4a2.003 2.003 0 0 1 2-2h16a2.003 2.003 0 0 1 2 2zM6 12.5a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5m12-5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 1 0V8a.5.5 0 0 0-.5-.5m-.32 8.291A9.459 9.459 0 0 1 12 17.5a9.447 9.447 0 0 1-5.68-1.71a.5.5 0 0 0-.641.767A10.255 10.255 0 0 0 12 18.5a10.258 10.258 0 0 0 6.321-1.942a.5.5 0 0 0-.641-.767"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySkeleton(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.853 2.159a.5.5 0 0 0-.707-.013l-1.89 1.89c-.01.009-.022.012-.03.022c-.009.008-.012.02-.02.028l-9.911 9.912A4.457 4.457 0 0 0 6.5 13a4.5 4.5 0 1 0 4.5 4.5a4.457 4.457 0 0 0-.998-2.795l6.757-6.757l2.473 2.474a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707L17.466 7.24l2.121-2.121l1.06 1.059a.498.498 0 0 0 .706 0a.5.5 0 0 0 0-.707l-1.059-1.06l1.56-1.559a.5.5 0 0 0 0-.694M6.5 21a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySkeletonAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.853 2.159a.5.5 0 0 0-.707-.013L15.714 7.58c-.007.006-.017.009-.024.016c-.006.006-.008.015-.014.022l-6.382 6.38A4.458 4.458 0 0 0 6.5 13a4.5 4.5 0 1 0 4.5 4.5a4.457 4.457 0 0 0-.999-2.795l6.05-6.05l1.767 1.768a.5.5 0 0 0 .707-.707l-1.767-1.768l2.122-2.121l1.767 1.768a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707L19.587 5.12l2.266-2.267a.5.5 0 0 0 0-.694M6.5 21a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 15H20V7a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v8H2.5a.5.5 0 0 0-.5.5V17a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3v-1.5a.5.5 0 0 0-.5-.5M5 7a2.003 2.003 0 0 1 2-2h10a2.003 2.003 0 0 1 2 2v8H5zm16 10a2.003 2.003 0 0 1-2 2H5a2.003 2.003 0 0 1-2-2v-1h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayerGroup(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.25 8.452l9.5 5.48a.498.498 0 0 0 .5 0l9.5-5.48a.5.5 0 0 0 0-.866l-9.5-5.476a.507.507 0 0 0-.5 0l-9.5 5.476a.5.5 0 0 0 0 .866M12 3.122l8.499 4.898L12 12.923L3.501 8.02zm9.248 12.404L12 20.921l-9.248-5.395a.5.5 0 1 0-.504.864l9.5 5.542a.496.496 0 0 0 .504 0l9.5-5.542a.5.5 0 1 0-.504-.864m0-4L12 16.921l-9.248-5.395a.5.5 0 1 0-.504.864l9.5 5.542a.496.496 0 0 0 .504 0l9.5-5.542a.5.5 0 1 0-.504-.864"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-13a.5.5 0 0 0-.5.5V7H5.5a.5.5 0 0 0-.5.5V12H2.5a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V19h4.5a.5.5 0 0 0 .5-.5V16h4.5a.5.5 0 0 0 .5-.5v-13a.5.5 0 0 0-.5-.5M11 18.5V21H3v-8h8zm5-3V18h-4v-5.5a.5.5 0 0 0-.5-.5H6V8h10zm5-.5h-4V7.5a.5.5 0 0 0-.5-.5H9V3h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftIndent(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 10.5h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1m0-4h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m19.345 2.852a.666.666 0 0 0-.939-.086l-2.666 2.222a.665.665 0 0 0 0 1.023l2.666 2.223a.661.661 0 0 0 .938-.084a.667.667 0 0 0-.084-.94L19.708 12l2.052-1.71a.666.666 0 0 0 .085-.938M2.5 14.5h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1m19 3h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftIndentAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.884 9.18a.5.5 0 0 0-.704-.064l-3 2.5a.5.5 0 0 0 0 .768l3 2.5a.498.498 0 0 0 .704-.064a.5.5 0 0 0-.064-.704L3.281 12l2.54-2.116a.5.5 0 0 0 .063-.704M12.5 7h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1M9.045 5h-.003a.5.5 0 0 0-.5.497l-.084 13a.5.5 0 0 0 .497.503h.003a.5.5 0 0 0 .5-.497l.084-13A.5.5 0 0 0 9.045 5M21.5 18h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1m0-8h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1m0 4h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineSpacing(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 8h12a.5.5 0 0 0 0-1h-12a.5.5 0 0 0 0 1m-3.459 8.17l-.874 1V6.83l.874 1a.497.497 0 0 0 .705.046a.5.5 0 0 0 .047-.705l-1.75-2a.516.516 0 0 0-.752 0l-1.75 2a.5.5 0 1 0 .752.658l.874-.998v10.338l-.874-.998a.5.5 0 0 0-.752.658l1.75 2a.499.499 0 0 0 .752 0l1.75-2a.5.5 0 0 0-.752-.658M21.5 17h-12a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1m0-5h-12a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkBroken(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.242 10.218a.5.5 0 0 0 .13-.984l-2.726-.73a.509.509 0 0 0-.613.353a.501.501 0 0 0 .354.614l2.725.73c.043.01.086.016.13.017m8.565 5.169a.5.5 0 0 0-.707-.007l-3.89 3.888a3.241 3.241 0 0 1-4.478 0v-.001a3.166 3.166 0 0 1 0-4.478L8.62 10.9a.5.5 0 0 0-.707-.707l-3.888 3.889a4.167 4.167 0 1 0 5.893 5.893l3.889-3.888a.5.5 0 0 0 0-.7M5.883 6.59a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707L4.595 3.888a.5.5 0 0 0-.707.707zm3.351-1.218a.502.502 0 0 0 .614.354a.501.501 0 0 0 .353-.614l-.73-2.725a.506.506 0 0 0-.614-.354a.501.501 0 0 0-.353.613zm10.74-1.347a4.167 4.167 0 0 0-5.892 0l-3.889 3.888a.5.5 0 0 0 .707.707l3.89-3.888a3.241 3.241 0 0 1 4.478 0v.001a3.166 3.166 0 0 1 0 4.478L15.38 13.1a.5.5 0 0 0 .707.707l3.888-3.889a4.167 4.167 0 0 0 0-5.893m1.64 10.504l-2.726-.73a.508.508 0 0 0-.614.353a.501.501 0 0 0 .354.614l2.726.73a.5.5 0 1 0 .26-.967m-3.497 2.881l-.007-.007a.5.5 0 0 0-.7.714l1.995 1.995a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707zm-3.351 1.218a.5.5 0 0 0-.967.26l.73 2.725a.502.502 0 0 0 .614.354a.501.501 0 0 0 .353-.613z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkH(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 16H7a4 4 0 1 1 0-8h3.5a.5.5 0 0 0 0-1H7a5 5 0 0 0 0 10h3.5a.5.5 0 0 0 0-1M8 12a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 0-1h-7a.5.5 0 0 0-.5.5m9-5h-3.5a.5.5 0 0 0 0 1H17a4 4 0 1 1 0 8h-3.5a.5.5 0 0 0 0 1H17a5 5 0 0 0 0-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 9h-4a.5.5 0 0 0-.5.5v12a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-12a.5.5 0 0 0-.5-.5M7 21H4V10h3zM18 9c-1.085 0-2.14.358-3 1.019V9.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0-.5.5v12a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V16a1.5 1.5 0 1 1 3 0v5.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V14a5.006 5.006 0 0 0-5-5m4 12h-3v-5a2.5 2.5 0 1 0-5 0v5h-3V10h3v1.203a.5.5 0 0 0 .89.313A3.983 3.983 0 0 1 22 14zM5.868 2.002A2.73 2.73 0 0 0 5.515 2a2.74 2.74 0 0 0-2.926 2.729a2.71 2.71 0 0 0 2.869 2.728h.028a2.734 2.734 0 1 0 .382-5.455M5.833 6.46a1.75 1.75 0 0 1-.347-.003h-.028A1.736 1.736 0 1 1 5.515 3a1.737 1.737 0 0 1 .318 3.46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUiAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 0h14a.5.5 0 0 0 0-1h-14a.5.5 0 0 0 0 1m14 4h-10a.5.5 0 0 0 0 1h10a.5.5 0 0 0 0-1m0 5h-6a.5.5 0 0 0 0 1h6a.5.5 0 0 0 0-1m-14-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUl(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 12a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m3-4h15a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1m-3 9a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m18-5h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1m-18-5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m18 10h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MasterCard(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 5.204a6.67 6.67 0 0 0-3.25.843a6.67 6.67 0 0 0-3.25-.843a6.796 6.796 0 0 0 0 13.592a6.67 6.67 0 0 0 3.25-.843a6.67 6.67 0 0 0 3.25.843a6.796 6.796 0 0 0 0-13.592m-6.5 12.592a5.796 5.796 0 0 1 0-11.592c.792 0 1.575.166 2.298.487a6.805 6.805 0 0 0 0 10.618a5.676 5.676 0 0 1-2.298.487m3.25-1.02A5.805 5.805 0 0 1 9.5 12A5.805 5.805 0 0 1 12 7.223a5.813 5.813 0 0 1 0 9.554m3.25 1.02a5.676 5.676 0 0 1-2.298-.487a6.805 6.805 0 0 0 0-10.618a5.676 5.676 0 0 1 2.298-.487a5.796 5.796 0 0 1 0 11.592"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microscope(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 22H14c.3-.4.5-.9.5-1.5v-.1c1.2-.5 2.4-1.2 3.3-2.2c3-3.3 2.9-8.4-.2-11.6l.5-.5c.1-.1.2-.1.2-.2l.7-2.1c.1-.2 0-.4-.1-.5l-2.1-2.1c-.2-.2-.4-.2-.6-.2l-2.1.7c-.1 0-.1.1-.2.1L7.6 8.2c-.2.2-.2.5 0 .7l.4.4L6.1 11c-.2.2-.2.5 0 .7l2.1 2.1c.2.2.3.2.4.2c.1 0 .3 0 .4-.1l1.8-1.8l.4.4c.1.1.2.1.4.1c.1 0 .3-.1.4-.1l5.1-5.1c.8.8 1.4 1.8 1.7 2.9c1.2 3.8-.7 7.8-4.4 9.2c-.5-.9-1.4-1.5-2.4-1.5c-1.1 0-2.1.7-2.4 1.8c-1.2-.3-2.3-.9-3.2-1.8h1.1c.3 0 .5-.2.5-.5s-.2-.5-.5-.5h-4c-.3 0-.5.2-.5.5s.2.5.5.5H5c1.2 1.4 2.7 2.3 4.5 2.8c0 .5.2.9.5 1.2H3.5c-.3 0-.5.2-.5.5s.2.5.5.5h17c.3 0 .5-.2.5-.5s-.2-.5-.5-.5M8.6 12.8l-1.4-1.4L8.6 10l1.4 1.4zm8-6.5s-.1 0 0 0c-.1 0-.1 0 0 0l-5.1 5.1l-2.9-2.8l5.9-5.9l1.7-.6l1.7 1.7l-.6 1.7zM12 22c-.8 0-1.5-.7-1.5-1.5S11.2 19 12 19s1.5.7 1.5 1.5S12.8 22 12 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareFull(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 12.5h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5M21 21H3V3h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modem(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 18.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 17 16.5m4.987-3.064a.488.488 0 0 0-.02-.095a.478.478 0 0 0-.043-.09a.481.481 0 0 0-.05-.073a.488.488 0 0 0-.08-.072c-.016-.012-.027-.029-.044-.039l-15.58-9a.5.5 0 1 0-.5.866L19.637 13H3.5a.5.5 0 0 0-.5.5v5A2.502 2.502 0 0 0 5.5 21h14a2.502 2.502 0 0 0 2.5-2.5v-5c0-.023-.01-.043-.013-.064M21 18.5a1.5 1.5 0 0 1-1.5 1.5h-14A1.5 1.5 0 0 1 4 18.5V14h17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseAltTwo(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a7 7 0 0 0-7 7v6a7.008 7.008 0 0 0 7 7a7.008 7.008 0 0 0 7-7V9a7 7 0 0 0-7-7m6 13a6 6 0 1 1-12 0V9a6.004 6.004 0 0 1 5.5-5.974V12.5a.5.5 0 1 0 1 0V3.026A6.004 6.004 0 0 1 18 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Multiply(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.853 19.147L12.707 12l7.146-7.147a.5.5 0 0 0-.707-.707L12 11.293L4.853 4.147a.5.5 0 0 0-.707.707L11.293 12l-7.147 7.146a.5.5 0 1 0 .707.707L12 12.707l7.146 7.147a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectGroup(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 19.092V4.908A1.497 1.497 0 0 0 20.5 2c-.652 0-1.202.419-1.408 1H4.908A1.496 1.496 0 0 0 2 3.5c0 .652.419 1.202 1 1.408v14.184A1.496 1.496 0 1 0 4.908 21h14.184A1.496 1.496 0 0 0 22 20.5c0-.652-.419-1.202-1-1.408M20.5 3a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 20.5 3m-17 0a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 3.5 3m0 18a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m15.59-.994L19.086 20H4.914l-.004.006a1.495 1.495 0 0 0-.916-.916L4 19.086V4.914l-.006-.004c.428-.15.765-.488.916-.916L4.914 4h14.172l.004-.006c.15.429.487.766.916.916L20 4.914v14.172l.006.004c-.429.15-.766.487-.916.916M20.5 21a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-4-11H14V7.5a.5.5 0 0 0-.5-.5h-6a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5H10v2.5a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5m-6 0a.5.5 0 0 0-.5.5V13H8V8h5v2zm5.5 6h-5v-5h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectUngroup(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 19.092v-8.184c.58-.207 1-.757 1-1.408c0-.827-.673-1.5-1.5-1.5a1.5 1.5 0 0 0-1.408 1H15V4.908c.58-.207 1-.757 1-1.408c0-.827-.673-1.5-1.5-1.5a1.5 1.5 0 0 0-1.408 1H4.908A1.5 1.5 0 0 0 3.5 2C2.673 2 2 2.673 2 3.5c0 .651.42 1.2 1 1.408v8.184c-.58.207-1 .757-1 1.408c0 .827.673 1.5 1.5 1.5a1.5 1.5 0 0 0 1.408-1H9v4.092c-.58.207-1 .757-1 1.408c0 .827.673 1.5 1.5 1.5a1.5 1.5 0 0 0 1.408-1h8.184c.207.58.757 1 1.408 1c.827 0 1.5-.673 1.5-1.5a1.5 1.5 0 0 0-1-1.408M20.5 9a.501.501 0 0 1 0 1a.501.501 0 0 1 0-1m-6-6a.501.501 0 0 1 0 1a.501.501 0 0 1 0-1m-11 0a.501.501 0 0 1 0 1a.501.501 0 0 1 0-1m0 12a.501.501 0 0 1 0-1a.501.501 0 0 1 0 1m1.414-1l-.005.002a1.498 1.498 0 0 0-.912-.91L4 13.085V4.914l-.003-.005c.426-.15.76-.485.911-.91L4.914 4h8.172l.005-.003c.151.426.485.76.91.911L14 4.914V9h-3.092A1.5 1.5 0 0 0 9.5 8C8.673 8 8 8.673 8 9.5c0 .651.42 1.2 1 1.408V14zm9.586 0a.501.501 0 0 1 0 1a.501.501 0 0 1 0-1m-.498-.908a1.49 1.49 0 0 0-.91.91L13.085 14H10v-3.086l-.003-.005c.426-.151.76-.485.911-.91c.001-.002.006.002.006.001H14v3.086zM9.5 10a.501.501 0 0 1 0-1a.501.501 0 0 1 0 1m3.592 5c.207.58.757 1 1.408 1c.827 0 1.5-.673 1.5-1.5a1.5 1.5 0 0 0-1-1.408V10h4.086l.005-.004c.15.427.486.762.912.913l-.003.005v8.172l.003.005a1.498 1.498 0 0 0-.914.918L19.086 20h-8.172l-.004.008a1.497 1.497 0 0 0-.918-.918l.008-.004V15zM9.5 21a.501.501 0 0 1 0-1a.501.501 0 0 1 0 1m11 0a.501.501 0 0 1 0-1a.501.501 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OperaAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5.5c-3.1 0-3.1 4.7-3.1 6.4c0 1.8 0 6.6 3.2 6.6s3.2-4.8 3.2-6.6c-.2-4.2-1.2-6.4-3.3-6.4m0 12c-1.5 0-2.2-1.8-2.2-5.7c0-3.6.7-5.3 2.1-5.3c1.4 0 2.2 1.8 2.2 5.4c0 3.8-.6 5.6-2.1 5.6m0-16c-5.8 0-9.8 4.3-9.8 10.4c0 5.2 3.7 10.6 9.8 10.6c6.1 0 9.8-5.4 9.8-10.6c0-6.1-4-10.4-9.8-10.4m0 20c-5.5 0-8.8-4.9-8.8-9.6c0-5.6 3.5-9.4 8.8-9.4c5.3 0 8.8 3.8 8.8 9.4c0 4.7-3.3 9.6-8.8 9.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.188 13.379l-6.011 6.01a5 5 0 0 1-7.072-7.07l8.486-8.486a3 3 0 1 1 4.243 4.242l-7.778 7.779a1.024 1.024 0 0 1-1.414 0a1.001 1.001 0 0 1 0-1.415l5.302-5.303a.5.5 0 1 0-.707-.707l-5.302 5.303l-.001.001a2 2 0 0 0 0 2.828a2.048 2.048 0 0 0 2.829 0l7.778-7.779a4 4 0 0 0-5.657-5.656l-8.486 8.485a6 6 0 0 0 4.244 10.243a5.957 5.957 0 0 0 4.242-1.758l6.01-6.01a.5.5 0 1 0-.707-.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 14h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1m8-5h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.654 7.587a3.31 3.31 0 0 0-.683-.57a4.123 4.123 0 0 0-.953-3.478c-.938-1.06-2.576-1.599-4.869-1.599H7.265c-.66.002-1.22.481-1.324 1.132l-2.453 15.59a1.007 1.007 0 0 0 .994 1.172H7.53l-.16 1.014a.945.945 0 0 0 .93 1.092h3.064a1.23 1.23 0 0 0 1.219-1.032l.606-3.827l.041-.217a.236.236 0 0 1 .234-.202h.458c3.629 0 5.803-1.72 6.462-5.112a4.48 4.48 0 0 0-.729-3.963m-11.763 9.94l-.206 1.306l-3.21-.01l2.455-15.6a.338.338 0 0 1 .335-.283h5.884c1.995 0 3.38.424 4.117 1.26c.702.825.955 1.942.677 2.99h.002a9.544 9.544 0 0 1-.056.326l-.001.003c-.65 3.347-2.721 4.975-6.333 4.975H9.827a1.335 1.335 0 0 0-1.323 1.136zm11.51-6.168c-.563 2.896-2.355 4.303-5.48 4.303h-.458a1.228 1.228 0 0 0-1.219 1.032l-.615 3.873l-.032.171a.237.237 0 0 1-.235.202l-3.005.065l.605-3.849l.534-3.37h-.004v-.002a.34.34 0 0 1 .335-.29h1.728c3.96 0 6.39-1.807 7.227-5.375c.042.041.082.083.12.126c.65.901.834 2.056.498 3.114"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pentagon(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.795 9.456l-9.5-6.923a.504.504 0 0 0-.59 0l-9.5 6.923a.502.502 0 0 0-.18.559l3.628 11.202a.5.5 0 0 0 .476.345H17.87a.5.5 0 0 0 .476-.345l3.629-11.202a.502.502 0 0 0-.181-.559m-4.287 11.107H6.492L3.087 10.05L12 3.557l8.913 6.494z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Polygon(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.934 11.751l-4.75-8.255a.501.501 0 0 0-.434-.25h-9.5a.501.501 0 0 0-.434.25l-4.75 8.255a.498.498 0 0 0 0 .498l4.75 8.255c.09.155.255.25.434.25h9.5a.501.501 0 0 0 .434-.25l4.75-8.255a.498.498 0 0 0 0-.498m-5.473 8.004H7.539L3.077 12L7.54 4.245h8.922L20.923 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 7a.5.5 0 0 0-.5.5v9a.5.5 0 1 0 1 0v-9a.5.5 0 0 0-.5-.5m9.354 9.146L12.698 12l4.155-4.146a.5.5 0 0 0-.707-.707l-4.51 4.5a.5.5 0 0 0 0 .707l4.51 4.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 9.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M19.5 6H18V2.5a.5.5 0 0 0-.5-.5h-11a.5.5 0 0 0-.5.5V6H4.5A2.502 2.502 0 0 0 2 8.5V15a3.003 3.003 0 0 0 3 3h1v3.5a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5V18h1a3.003 3.003 0 0 0 3-3V8.5A2.502 2.502 0 0 0 19.5 6M7 3h10v3H7zm10 18H7v-6h10zm4-6a2.003 2.003 0 0 1-2 2h-1v-2.5a.5.5 0 0 0-.5-.5h-11a.5.5 0 0 0-.5.5V17H5a2.003 2.003 0 0 1-2-2V8.5A1.5 1.5 0 0 1 4.5 7h15A1.5 1.5 0 0 1 21 8.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Process(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 7H5.145a8.504 8.504 0 0 1 8.274-3.387a.5.5 0 0 0 .162-.986A9.849 9.849 0 0 0 12 2.5a9.52 9.52 0 0 0-7.5 3.677V2.5a.5.5 0 0 0-1 0v5A.5.5 0 0 0 4 8h5a.5.5 0 0 0 0-1m-1.5 7.5a.5.5 0 0 0-.5.5v3.855a8.504 8.504 0 0 1-3.387-8.274a.5.5 0 0 0-.986-.162a9.52 9.52 0 0 0 3.55 9.081H2.5a.5.5 0 0 0 0 1h5A.5.5 0 0 0 8 20v-5a.5.5 0 0 0-.5-.5M20 16h-5a.5.5 0 0 0 0 1h3.855a8.504 8.504 0 0 1-8.274 3.387a.5.5 0 0 0-.162.986A9.849 9.849 0 0 0 12 21.5a9.52 9.52 0 0 0 7.5-3.677V21.5a.5.5 0 0 0 1 0v-5a.5.5 0 0 0-.5-.5m1.5-12.5h-5a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 1 0V5.14a8.3 8.3 0 0 1 2.358 2.612A8.441 8.441 0 0 1 20.5 12c0 .475-.037.95-.113 1.419a.499.499 0 1 0 .986.162A9.855 9.855 0 0 0 21.5 12a9.443 9.443 0 0 0-1.275-4.747A9.294 9.294 0 0 0 17.828 4.5H21.5a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pump(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.682 6.975l-.828-.828l-2.975-2.976a.5.5 0 1 0-.707.707l1.747 1.748l-.935 2.492a2.506 2.506 0 0 0 .573 2.646L21 14.207V19c0 1.103-.897 2-2 2s-2-.897-2-2v-2.5c0-1.379-1.121-2.5-2.5-2.5H14V4.5C14 3.122 12.879 2 11.5 2h-7A2.503 2.503 0 0 0 2 4.5v15C2 20.879 3.122 22 4.5 22h7c1.379 0 2.5-1.121 2.5-2.5V15h.5c.827 0 1.5.673 1.5 1.5V19c0 1.654 1.346 3 3 3s3-1.346 3-3v-8.843c0-1.184-.48-2.344-1.318-3.182M13 19.5c0 .827-.673 1.5-1.5 1.5h-7c-.827 0-1.5-.673-1.5-1.5V11h10zm0-9.5H3V4.5C3 3.673 3.673 3 4.5 3h7c.827 0 1.5.673 1.5 1.5zm8 2.793l-2.736-2.735a1.508 1.508 0 0 1-.344-1.588l.775-2.067l1.28 1.28A3.526 3.526 0 0 1 21 10.156z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircle(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.596 8.105A2.498 2.498 0 0 0 9.55 9.897a.5.5 0 1 0 .969.25a1.5 1.5 0 1 1 1.926 1.796a1.507 1.507 0 0 0-.981 1.452v.628a.5.5 0 1 0 1 0v-.628a.517.517 0 0 1 .304-.504a2.498 2.498 0 0 0-.173-4.786m-.631 7.292a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecordAudio(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10m0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9m0-15a6 6 0 1 0 0 12a6.007 6.007 0 0 0 6-6a6 6 0 0 0-6-6m0 11a5 5 0 1 1 0-10a5 5 0 0 1 0 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditAlienAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.5 11.72a3.278 3.278 0 0 0-6.098-1.673a15.824 15.824 0 0 0-3.824-.588l1.027-6.853l3.022 1.027a2.49 2.49 0 0 0-.034.285a2.5 2.5 0 1 0 5 0v-.01a2.526 2.526 0 0 0-2.5-2.468a2.495 2.495 0 0 0-2.139 1.248l-3.593-1.221a.5.5 0 0 0-.655.4l-1.138 7.587a15.872 15.872 0 0 0-3.97.593a3.27 3.27 0 1 0-5.426 3.627A4.361 4.361 0 0 0 1.5 15.94c0 3.584 4.71 6.5 10.5 6.5s10.5-2.916 10.5-6.5a4.36 4.36 0 0 0-.671-2.265a3.28 3.28 0 0 0 .671-1.955m-4.907-7.791a1.5 1.5 0 1 1 0 .012zm2.964 5.929c.94.675 1.22 1.952.646 2.958a9.524 9.524 0 0 0-3.804-2.44a2.263 2.263 0 0 1 3.158-.518m-16.92-.12a2.272 2.272 0 0 1 2.964.638a9.523 9.523 0 0 0-3.804 2.44a2.256 2.256 0 0 1 .84-3.078M12 21.44c-5.238 0-9.5-2.467-9.5-5.5c0-2.095 2.036-3.92 5.022-4.848c.006-.002.012 0 .017-.003l.003-.002A15.185 15.185 0 0 1 12 10.44c1.51-.007 3.012.21 4.458.647l.003.002l.017.003c2.986.929 5.022 2.753 5.022 4.848c0 3.033-4.262 5.5-9.5 5.5m3.2-8.999a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m-6 3.001a1.5 1.5 0 1 0 .002-3a1.5 1.5 0 0 0-.003 3zm-.5-1.5a.5.5 0 1 1 .501.5a.501.501 0 0 1-.5-.5m6.064 3.22a3.764 3.764 0 0 1-2.722.778a3.76 3.76 0 0 1-2.72-.778a.5.5 0 0 0-.708.707a4.653 4.653 0 0 0 3.428 1.071a4.66 4.66 0 0 0 3.429-1.07a.5.5 0 0 0-.707-.708"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 11.5a.5.5 0 0 0-.5.5a9.02 9.02 0 1 1-1.502-5H16.5a.5.5 0 0 0 0 1h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-1 0v2.497A9.994 9.994 0 0 0 12.025 2C6.502 1.993 2.02 6.465 2.013 11.987C2.006 17.51 6.477 21.993 12 22c5.52-.006 9.994-4.48 10-10a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 16h-4a.5.5 0 0 0 0 1h2.976a8.996 8.996 0 0 1-7.459 4a9.009 9.009 0 0 1-8.904-10.419a.5.5 0 0 0-.986-.162A9.847 9.847 0 0 0 2 12c.006 5.52 4.48 9.994 10 10a10.009 10.009 0 0 0 8-3.999v2.5a.5.5 0 1 0 1 0v-4.002a.5.5 0 0 0-.5-.499M16.737 3.196C12.32.818 6.921 2.107 4 6.003V3.5a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 0-1H4.524a8.996 8.996 0 0 1 7.459-4a9.009 9.009 0 0 1 8.903 10.421a.499.499 0 1 0 .987.16A9.847 9.847 0 0 0 22 12a10.008 10.008 0 0 0-5.263-8.804M15 12a3 3 0 1 0-6 0a3 3 0 0 0 6 0m-5 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 5h-4a.5.5 0 0 0 0 1H21v12H10.707l1.646-1.646a.5.5 0 0 0-.707-.707l-2.5 2.5a.5.5 0 0 0 0 .707l2.5 2.5a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.708L10.708 19h10.794a.5.5 0 0 0 .499-.5v-13a.5.5 0 0 0-.5-.5m-16 13H3V6h10.293l-1.647 1.646a.5.5 0 1 0 .707.708l2.5-2.5a.5.5 0 0 0 0-.707l-2.5-2.5a.5.5 0 0 0-.707.707L13.293 5H2.499A.5.5 0 0 0 2 5.5v13a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightIndent(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.555 14a.5.5 0 0 0 .816.387l2.445-2a.5.5 0 0 0 0-.773l-2.445-2a.5.5 0 0 0-.633.773L20.71 12l-1.972 1.613a.5.5 0 0 0-.183.386M2.5 6.5h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m0 4h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1m0 4h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1m19 3h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightIndentAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.82 9.116a.5.5 0 0 0-.64.768L4.719 12l-2.54 2.116a.5.5 0 0 0 .641.768l3-2.5a.5.5 0 0 0 0-.768zM12.5 7h9a.5.5 0 0 0 0-1h-9a.5.5 0 0 0 0 1M9.045 5h-.003a.5.5 0 0 0-.5.497l-.084 13a.5.5 0 0 0 .497.503h.003a.5.5 0 0 0 .5-.497l.084-13A.5.5 0 0 0 9.045 5M21.5 10h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1m0 8h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1m0-4h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.6 2.7c0-.2-.2-.3-.4-.4c-3.8-1-7.9.3-10.4 3.3L9.5 7.1l-2.7-.7c-1.1-.4-2.2.1-2.7 1.1L2 11.2s0 .1-.1.1c-.1.3.1.5.4.6l3.4.7c-.3.9-.6 1.8-.7 2.7c0 .2 0 .3.1.4l3 2.9c.1.1.2.1.4.1c.9-.1 1.9-.3 2.8-.6l.7 3.3c0 .2.3.4.5.4c.1 0 .2 0 .2-.1l3.7-2.1c.9-.5 1.3-1.6 1.1-2.6l-.7-2.9l1.4-1.3c3.1-2.3 4.4-6.3 3.4-10.1M3.2 11.1L4.9 8c.3-.6.9-.8 1.5-.6l2.3.6l-1 1.2c-.6.8-1.2 1.6-1.6 2.5zM16 19l-3.1 1.8l-.6-2.9c.9-.4 1.7-1 2.5-1.6l1.3-1.2l.6 2.3c0 .6-.2 1.3-.7 1.6m1.6-6.7l-3.5 3.2c-1.5 1.3-3.4 2.1-5.4 2.3l-2.6-2.6c.3-2 1.1-3.9 2.4-5.4L10.1 8l.1-.1l1.4-1.6c2.2-2.6 5.8-3.8 9.1-3.1c.7 3.4-.4 6.9-3.1 9.1m-1.2-6.7c-1.1 0-1.9.9-1.9 1.9s.9 1.9 1.9 1.9c1.1 0 1.9-.9 1.9-1.9c0-1-.8-1.9-1.9-1.9m0 2.9c-.5 0-.9-.4-.9-.9s.4-.9.9-.9s.9.4.9.9s-.4.9-.9.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.96 7.404L16.596 1.04a.5.5 0 0 0-.707 0L1.04 15.889a.5.5 0 0 0 0 .707l6.364 6.364a.5.5 0 0 0 .707 0l3.18-3.18l.002-.002l2.827-2.827h.001v-.002l2.829-2.827v-.001l2.828-2.828l3.182-3.182a.5.5 0 0 0 0-.707m-3.535 2.828l-1.768-1.767a.5.5 0 1 0-.707.707l1.768 1.767l-2.122 2.122l-3.182-3.182a.5.5 0 1 0-.707.707l3.182 3.182l-2.121 2.121L12 14.121a.5.5 0 0 0-.707.707l1.767 1.768l-2.12 2.122l-3.183-3.183a.5.5 0 1 0-.707.707l3.182 3.183l-2.475 2.474l-5.656-5.657L16.242 2.101L21.9 7.758z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RulerCombined(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5V10h11.5a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5M21 9h-3.5V6.5a.5.5 0 0 0-1 0V9h-3V6.5a.5.5 0 0 0-1 0V9H10V6.5a.5.5 0 0 0-1 0V9H6.5a.5.5 0 0 0 0 1H9v2.5H6.5a.5.5 0 0 0 0 1H9v3H6.5a.5.5 0 0 0 0 1H9V21H3V3h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sanitizer(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 12.5c-1.4 0-2.5 1.1-2.5 2.5s1.1 2.5 2.5 2.5S15 16.4 15 15s-1.1-2.5-2.5-2.5m0 4c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5m5.1-9.3l-2.6-2V3h1.5c.3 0 .5-.2.5-.5s-.2-.5-.5-.5H7.7c-.9 0-1.8.4-2.4 1L4.1 4.1c0 .1-.1.3-.1.4c0 .3.2.5.5.5c.1 0 .3-.1.4-.1L6 3.7c.4-.4 1.1-.7 1.7-.7H10v2.2l-2.6 2C6.5 7.9 6 8.9 6 10v11.5c0 .3.2.5.5.5h12c.3 0 .5-.2.5-.5V10c0-1.1-.5-2.1-1.4-2.8M11 3h3v2h-3zm7 18H7V10c0-.8.4-1.5 1-2l2.7-2h3.7L17 8c.6.5 1 1.2 1 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SanitizerAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 16h-3c-.3 0-.5.2-.5.5s.2.5.5.5h3c.3 0 .5-.2.5-.5s-.2-.5-.5-.5m1.8-8H16V4.5c0-.3-.2-.5-.5-.5h-2V2h1c.3 0 .5-.2.5-.5s-.2-.5-.5-.5H9.7c-1.6 0-3 .9-3.7 2.3c-.1.2 0 .5.3.6c0 .1.1.1.2.1c.2 0 .4-.1.4-.3c.6-1 1.7-1.7 2.8-1.7h2.8v2h-2c-.3 0-.5.2-.5.5V8h-.3C8.2 8 7 9.2 7 10.7v9.6C7 21.8 8.2 23 9.7 23h6.6c1.5 0 2.7-1.2 2.7-2.7v-9.6C19 9.2 17.8 8 16.3 8M11 5h4v3h-4zm7 15.3c0 .9-.8 1.7-1.7 1.7H9.7c-.9 0-1.7-.8-1.7-1.7v-9.6C8 9.8 8.8 9 9.7 9h6.6c.9 0 1.7.8 1.7 1.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scenery(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14c.182 0 .362-.022.54-.055c.066-.012.129-.03.193-.047a2.971 2.971 0 0 0 .531-.187c.084-.04.165-.086.245-.133c.075-.044.151-.086.222-.136l.018-.011l.021-.018c.014-.011.03-.017.043-.03c.007-.006.01-.015.015-.022A2.987 2.987 0 0 0 22 19V5a3.003 3.003 0 0 0-3-3m.575 18.905A1.95 1.95 0 0 1 19 21H5a2.003 2.003 0 0 1-2-2v-4.725l3.763-3.762a1.753 1.753 0 0 1 2.474 0l3.405 3.404l.004.007l6.97 6.969zM21 19c0 .516-.202.982-.523 1.337l-6.769-6.768l1.056-1.055a1.787 1.787 0 0 1 2.472 0L21 16.277zm0-4.137l-3.057-3.056a2.75 2.75 0 0 0-3.886 0L13 12.862L9.944 9.806a2.753 2.753 0 0 0-3.888 0L3 12.86V5a2.003 2.003 0 0 1 2-2h14a2.003 2.003 0 0 1 2 2zM13.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.99 3.753a.5.5 0 0 0-.593-.387a9.106 9.106 0 0 1-7.11-1.454a.5.5 0 0 0-.574 0a9.107 9.107 0 0 1-7.11 1.454a.498.498 0 0 0-.603.49v8.018a9.131 9.131 0 0 0 3.799 7.407l3.91 2.804a.497.497 0 0 0 .582 0l3.91-2.804A9.131 9.131 0 0 0 20 11.874V3.855a.498.498 0 0 0-.01-.102M19 11.874a8.129 8.129 0 0 1-3.38 6.595L12 21.063L8.38 18.47A8.13 8.13 0 0 1 5 11.874v-7.42a10.082 10.082 0 0 0 7-1.53a10.085 10.085 0 0 0 7 1.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.896 9.946l-3.046 3.047l-1.247-1.247a.5.5 0 0 0-.707.707l1.6 1.6a.5.5 0 0 0 .707 0l3.408-3.407a.5.5 0 0 0-.715-.7m6.093-6.193a.5.5 0 0 0-.592-.387a9.107 9.107 0 0 1-7.11-1.454a.5.5 0 0 0-.574 0a9.106 9.106 0 0 1-7.11 1.454a.498.498 0 0 0-.603.49v8.018a9.131 9.131 0 0 0 3.799 7.407l3.91 2.804a.497.497 0 0 0 .582 0l3.91-2.804A9.131 9.131 0 0 0 20 11.874V3.855a.498.498 0 0 0-.01-.102M19 11.874a8.129 8.129 0 0 1-3.38 6.595L12 21.063L8.38 18.47A8.13 8.13 0 0 1 5 11.874v-7.42a10.082 10.082 0 0 0 7-1.53a10.087 10.087 0 0 0 7 1.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldExclamation(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-1 0v3a.5.5 0 0 0 .5.5m7.99-9.247a.5.5 0 0 0-.593-.387a9.106 9.106 0 0 1-7.11-1.454a.5.5 0 0 0-.574 0a9.107 9.107 0 0 1-7.11 1.454a.498.498 0 0 0-.603.49v8.018a9.131 9.131 0 0 0 3.799 7.407l3.91 2.804a.497.497 0 0 0 .582 0l3.91-2.804A9.131 9.131 0 0 0 20 11.874V3.855a.498.498 0 0 0-.01-.102M19 11.874a8.129 8.129 0 0 1-3.38 6.595L12 21.063L8.38 18.47A8.13 8.13 0 0 1 5 11.874v-7.42a10.082 10.082 0 0 0 7-1.53a10.084 10.084 0 0 0 7 1.53zm-7 2.001a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldPlus(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 11.5h-2v-2c0-.3-.2-.5-.5-.5s-.5.2-.5.5v2h-2c-.3 0-.5.2-.5.5s.2.5.5.5h2v2c0 .3.2.5.5.5s.5-.2.5-.5v-2h2c.3 0 .5-.2.5-.5s-.2-.5-.5-.5M20 3.8c-.1-.3-.3-.4-.6-.4c-2.5.5-5 0-7.1-1.5c-.2-.1-.4-.1-.6 0c-2.1 1.4-4.6 2-7.1 1.5h-.1c-.3 0-.5.2-.5.5v8c0 2.9 1.4 5.7 3.8 7.4l3.9 2.8c.1.1.2.1.3.1c.1 0 .2 0 .3-.1l3.9-2.8c2.4-1.7 3.8-4.5 3.8-7.4v-8zm-1 8.1c0 2.6-1.3 5.1-3.4 6.6L12 21.1l-3.6-2.6C6.3 17 5 14.5 5 11.9V4.5c2.4.4 4.9-.2 7-1.5c2.1 1.3 4.6 1.9 7 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldQuestion(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.596 8.083A2.498 2.498 0 0 0 9.55 9.875a.5.5 0 1 0 .969.25a1.5 1.5 0 1 1 1.926 1.796a1.507 1.507 0 0 0-.981 1.451v.629a.5.5 0 1 0 1 0v-.629a.524.524 0 0 1 .304-.504a2.498 2.498 0 0 0-.173-4.785m-.631 7.292a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25m8.024-11.622a.5.5 0 0 0-.592-.387a9.105 9.105 0 0 1-7.11-1.454a.5.5 0 0 0-.574 0a9.105 9.105 0 0 1-7.11 1.454a.498.498 0 0 0-.603.49v8.018a9.131 9.131 0 0 0 3.799 7.407l3.91 2.804a.497.497 0 0 0 .582 0l3.91-2.804a9.131 9.131 0 0 0 3.8-7.407V3.855a.498.498 0 0 0-.01-.102M19 11.874a8.129 8.129 0 0 1-3.38 6.595L12 21.063L8.38 18.47A8.13 8.13 0 0 1 5 11.874v-7.42a10.082 10.082 0 0 0 7-1.53a10.085 10.085 0 0 0 7 1.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldSlash(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.285 4.394A9.767 9.767 0 0 0 12 2.924a10.188 10.188 0 0 0 7 1.53v7.416a7.896 7.896 0 0 1-.47 2.7a.486.486 0 0 0-.03.17v.014a.496.496 0 0 0 .5.491a.501.501 0 0 0 .476-.35A8.91 8.91 0 0 0 20 11.87V3.86a.502.502 0 0 0-.6-.49a9.244 9.244 0 0 1-7.112-1.458a.496.496 0 0 0-.576 0a8.802 8.802 0 0 1-3.597 1.495l-.004.001a.5.5 0 1 0 .174.986m13.568 16.753l-19-19a.5.5 0 0 0-.707.707L4 4.707v7.163a9.123 9.123 0 0 0 3.798 7.406l3.91 2.81a.5.5 0 0 0 .584 0l3.911-2.81a8.75 8.75 0 0 0 1.255-1.111l3.688 3.689a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707m-6.235-2.683L12 21.064l-3.62-2.6A8.122 8.122 0 0 1 5 11.87V5.707l11.752 11.752a7.85 7.85 0 0 1-1.134 1.005"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 18a.5.5 0 0 0-.5.5v3a.5.5 0 1 0 1 0v-3a.5.5 0 0 0-.5-.5m5-4a.5.5 0 0 0-.5.5v7a.5.5 0 1 0 1 0v-7a.5.5 0 0 0-.5-.5m10-12a.5.5 0 0 0-.5.5v19a.5.5 0 1 0 1 0v-19a.5.5 0 0 0-.5-.5m-5 7a.5.5 0 0 0-.5.5v12a.5.5 0 1 0 1 0v-12a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalAltThree(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 16h-4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5M6 22H3v-5h3zM22.5 2h-4a.5.5 0 0 0-.5.5v20a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-20a.5.5 0 0 0-.5-.5M22 22h-3V3h3zm-7.5-12h-4a.5.5 0 0 0-.5.5v12a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-12a.5.5 0 0 0-.5-.5M14 22h-3V11h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signout(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 12a.5.5 0 0 0 .5.5h8.793l-2.647 2.646a.5.5 0 1 0 .707.708l3.5-3.5a.5.5 0 0 0 0-.707l-3.5-3.5a.5.5 0 0 0-.707.707l2.647 2.646H4.5a.5.5 0 0 0-.5.5M17.5 2h-11A2.502 2.502 0 0 0 4 4.5v4a.5.5 0 0 0 1 0v-4A1.5 1.5 0 0 1 6.5 3h11A1.5 1.5 0 0 1 19 4.5v15a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 5 19.5v-4a.5.5 0 0 0-1 0v4A2.502 2.502 0 0 0 6.5 22h11a2.502 2.502 0 0 0 2.5-2.5v-15A2.502 2.502 0 0 0 17.5 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimCard(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.98 4H7.01A3.014 3.014 0 0 0 4 7.01v9.97A3.025 3.025 0 0 0 7.02 20h9.97A3.014 3.014 0 0 0 20 16.99V7.02A3.025 3.025 0 0 0 16.98 4M10 5h4v4h-4zM9 19H7.02A2.023 2.023 0 0 1 5 16.98V15h4zm5 0h-4v-4h4zm5-2.01A2.012 2.012 0 0 1 16.99 19H15v-4h4zM19 14H5V7.01A2.01 2.01 0 0 1 7.01 5H9v4.5a.5.5 0 0 0 .5.5H19zm0-5h-4V5h1.98A2.023 2.023 0 0 1 19 7.02z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkypeAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.867 12.232a4.529 4.529 0 0 0-1.164-.534a17.92 17.92 0 0 0-.861-.234a18.5 18.5 0 0 0-1.036-.252a9.876 9.876 0 0 1-1.518-.45a1.986 1.986 0 0 1-.819-.602a1.346 1.346 0 0 1-.264-.872c0-.324.106-.64.304-.896c.241-.297.56-.52.922-.643a4.477 4.477 0 0 1 1.601-.249c.425-.006.849.052 1.257.17c.297.084.577.222.824.407c.183.134.337.303.453.497a.5.5 0 0 0 .881-.474a2.673 2.673 0 0 0-.71-.804a3.52 3.52 0 0 0-1.151-.581a5.286 5.286 0 0 0-1.554-.215a5.451 5.451 0 0 0-1.961.316a3.06 3.06 0 0 0-1.362.975c-.327.43-.504.956-.504 1.497a2.36 2.36 0 0 0 .475 1.486c.324.406.747.721 1.23.914c.411.158.833.29 1.262.393c.008.003.015.01.024.013c.183.054.478.128.9.22c.173.037.335.078.498.119c.03.008.07.017.098.026c.012.004.024.003.036.005c.24.062.471.127.684.191c.318.094.622.232.901.41c.21.136.387.32.515.536c.13.25.192.53.183.813c.018.373-.091.742-.31 1.044a2.225 2.225 0 0 1-1.008.749c-.536.21-1.109.31-1.685.293a4.174 4.174 0 0 1-1.934-.4a2.288 2.288 0 0 1-.801-.713a1.456 1.456 0 0 1-.286-.78a.5.5 0 1 0-1 0c.017.485.178.954.463 1.347c.3.433.699.788 1.164 1.033a5.134 5.134 0 0 0 2.394.513a5.228 5.228 0 0 0 2.07-.37a3.185 3.185 0 0 0 1.438-1.093a2.625 2.625 0 0 0 .495-1.641a2.453 2.453 0 0 0-1.144-2.164m6.974 1.536A9.993 9.993 0 0 0 10.233 2.16a6 6 0 0 0-8.074 8.072A9.992 9.992 0 0 0 13.768 21.84a5.999 5.999 0 0 0 8.073-8.072m-2.928 7.112a4.999 4.999 0 0 1-4.823 0a.505.505 0 0 0-.339-.053A8.994 8.994 0 0 1 3.173 10.25a.499.499 0 0 0-.053-.338a5 5 0 0 1 6.791-6.79a.482.482 0 0 0 .339.052A8.995 8.995 0 0 1 20.827 13.75a.5.5 0 0 0 .053.339a4.999 4.999 0 0 1-1.967 6.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.2 12c.8-.5 1.3-1.5 1.3-2.5c0-1.7-1.3-3-3-3c-.8 0-1.5.3-2 .8V4.5c0-1.7-1.3-3-3-3c-1 0-2 .5-2.5 1.3c-.5-.8-1.5-1.3-2.5-1.3c-1.7 0-3 1.3-3 3c0 .8.3 1.5.8 2H4.5c-1.7 0-3 1.3-3 3c0 1 .5 2 1.3 2.5c-.8.5-1.3 1.5-1.3 2.5c0 1.7 1.3 3 3 3c.8 0 1.5-.3 2-.8v2.8c0 1.7 1.3 3 3 3c1 0 2-.5 2.5-1.3c.5.8 1.5 1.3 2.5 1.3c1.7 0 3-1.3 3-3c0-.8-.3-1.5-.8-2h2.8c1.7 0 3-1.3 3-3c0-1-.5-2-1.3-2.5m-3.7-2.5c0-1.1.9-2 2-2s2 .9 2 2s-.9 2-2 2h-2zm-5-5c0-1.1.9-2 2-2s2 .9 2 2v5c0 1.1-.9 2-2 2s-2-.9-2-2zm-3-2c1.1 0 2 .9 2 2v2h-2c-1.1 0-2-.9-2-2s.9-2 2-2m-3 12c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2h2zm-2-3c-1.1 0-2-.9-2-2s.9-2 2-2h5c1.1 0 2 .9 2 2s-.9 2-2 2zm7 8c0 1.1-.9 2-2 2s-2-.9-2-2v-5c0-1.1.9-2 2-2s2 .9 2 2zm.5-6.7c-.2-.3-.5-.6-.8-.8c.3-.2.6-.5.8-.8c.2.3.5.6.8.8c-.3.2-.6.5-.8.8m2.5 8.7c-1.1 0-2-.9-2-2v-2h2c1.1 0 2 .9 2 2s-.9 2-2 2m5-5h-5c-1.1 0-2-.9-2-2s.9-2 2-2h5c1.1 0 2 .9 2 2s-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.9 11.2c.3.1.5-.1.6-.4c.1-.1.3-.2.4-.2c.4.2.8.4 1.2.4c.3 0 .7-.1.9-.4c.1-.1.2-.2.2-.4v-.6c-.2-1.2-.1-2.5.2-3.8C8.2 4.1 9.9 3 11.8 3h.4c1.9 0 3.6 1.1 4.4 2.8c.3 1.3.4 2.6.2 3.9v.5c0 .1.1.3.2.4c.2.2.6.4.9.4c.4 0 .8-.2 1.1-.4h.3c.1.1.2.1.2.1c0 .3.2.5.5.5s.5-.2.5-.5c-.1-.5-.4-.9-.9-1c-.3-.2-.7-.2-1.1 0c-.2.1-.5.2-.7.3v-.3c.2-1.4.1-2.9-.3-4.3c-.9-2.1-3-3.5-5.3-3.4h-.4C9.5 2 7.4 3.3 6.5 5.4c-.4 1.4-.5 2.9-.3 4.3v.3c-.3 0-.6-.1-.9-.3c-.1 0-.3-.1-.4-.1c-.6 0-1.2.4-1.4.9c0 .3.1.6.4.7m18 5.4c-1.4-.3-2.6-1.2-3.3-2.4c-.2-.2-.5-.3-.7-.1c-.2.2-.3.5-.1.7c.8 1.3 2 2.2 3.5 2.7c-.6.2-1.2.4-1.9.5c-.5.1-.6.6-.7 1v.3h-.2c-.7-.2-1.4-.2-2.1-.1c-.6.1-1.2.4-1.7.9c-.7.6-1.5.9-2.4 1H12c-.9 0-1.7-.4-2.4-1c-.5-.4-1.1-.7-1.7-.9c-.7-.1-1.4-.1-2.1.1h-.2V19c-.1-.4-.2-.9-.7-1c-.6-.1-1.3-.2-1.9-.5c1.1-.3 2.2-1 2.9-1.9c.2-.3.4-.5.6-.8c.1-.2.1-.5-.2-.7c-.2-.1-.5-.1-.7.2c-.1.2-.3.5-.5.7c-.7.9-1.7 1.5-2.8 1.7c-.4.1-.6.4-.6.7c0 .1 0 .2.1.3c.2.4.6.9 2.7 1.2c0 .1.1.2.1.3v.2c0 .1.1.2.1.3c.1.4.4.6.8.6c.2 0 .3 0 .5-.1c.6-.1 1.2-.1 1.8 0c.5.1.9.4 1.3.7c.7.7 1.7 1.1 2.9 1.1h.2c1.1 0 2.1-.4 3-1.2c.4-.3.8-.6 1.3-.7c.6-.1 1.2-.1 1.8.1c.2 0 .4.1.5.1c.4 0 .7-.2.8-.6c0-.1.1-.2.1-.3v-.2c0-.1 0-.2.1-.3c2.1-.4 2.6-.9 2.7-1.2c0-.1.1-.1.1-.2c0-.4-.3-.8-.7-.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancing(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 17.5H4.2l1.6-1.6c.2-.2.2-.5 0-.7c-.2-.2-.5-.2-.7 0l-2.5 2.5s-.1.1-.1.2v.4c0 .1.1.1.1.2L5.1 21c.1-.1.3 0 .4 0c.1 0 .3-.1.4-.1c.2-.2.2-.5 0-.7l-1.6-1.6H9c.3 0 .5-.2.5-.5s-.2-.6-.5-.6m12.4.1l-2.5-2.5c-.2-.2-.5-.2-.7 0c-.2.2-.2.5 0 .7l1.6 1.6H15c-.3 0-.5.2-.5.5s.2.5.5.5h4.8L18.2 20c-.1.1-.1.2-.1.4c0 .3.2.5.5.5c.1 0 .3-.1.4-.1l2.5-2.5s.1-.1.1-.2v-.4c-.2 0-.2 0-.2-.1M7.9 9.9c.8-.6 1.4-1.5 1.4-2.6C9.2 5.5 7.8 4 6 4S2.8 5.5 2.8 7.2c0 1.1.5 2 1.4 2.6c-1.9.8-3.2 2.6-3.2 4.7c0 .3.2.5.5.5s.5-.2.5-.5c0-2.2 1.8-4 4-4s4 1.8 4 4c0 .3.2.5.5.5s.5-.2.5-.5c0-2.1-1.3-3.9-3.1-4.6M6 9.5c-1.2 0-2.2-1-2.2-2.2S4.8 5 6 5c1.2 0 2.2 1 2.2 2.2c0 1.3-1 2.3-2.2 2.3m13.9.4c.8-.6 1.4-1.5 1.4-2.6C21.2 5.5 19.8 4 18 4s-3.2 1.5-3.2 3.2c0 1.1.5 2 1.4 2.6c-1.8.7-3.1 2.5-3.1 4.6c0 .3.2.5.5.5s.5-.2.5-.5c0-2.2 1.8-4 4-4s4 1.8 4 4c0 .3.2.5.5.5s.5-.2.5-.5c-.1-2-1.4-3.8-3.2-4.5M18 9.5c-1.2 0-2.2-1-2.2-2.2S16.8 5 18 5c1.2 0 2.2 1 2.2 2.2c0 1.3-1 2.3-2.2 2.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpaceKey(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 9a.5.5 0 0 0-.5.5V14H3V9.5a.5.5 0 0 0-1 0v5a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareFull(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5M21 21H3V3h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.995 9.933a.5.5 0 0 0-.423-.568l-6.304-.919l-2.82-5.73a.52.52 0 0 0-.896 0l-2.82 5.73l-6.304.92a.5.5 0 0 0-.278.852l4.563 4.46l-1.077 6.3a.5.5 0 0 0 .726.527L12 18.532l5.638 2.973a.506.506 0 0 0 .316.05a.5.5 0 0 0 .41-.576l-1.077-6.3l4.563-4.461a.5.5 0 0 0 .145-.285M16.4 14.147a.501.501 0 0 0-.143.442l.95 5.558l-4.974-2.623a.506.506 0 0 0-.466 0l-4.974 2.623l.95-5.558a.501.501 0 0 0-.143-.442L3.572 10.21l5.565-.81a.501.501 0 0 0 .376-.275L12 4.07l2.487 5.054a.5.5 0 0 0 .376.274l5.565.811z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.919 10.127a1 1 0 0 0-.845-1.136l-5.651-.826l-2.526-5.147a1.037 1.037 0 0 0-1.795.001L8.577 8.165l-5.651.826a1 1 0 0 0-.556 1.704l4.093 4.013l-.966 5.664a1.002 1.002 0 0 0 1.453 1.052l5.05-2.67l5.049 2.669a1 1 0 0 0 1.454-1.05l-.966-5.665l4.094-4.014a1 1 0 0 0 .288-.567M11.5 17.887L6.483 20.54l1.01-5.922a.502.502 0 0 0-.143-.441L3.07 9.98l5.912-.864a.503.503 0 0 0 .377-.275L11.5 4.479zm5.15-3.71a.502.502 0 0 0-.143.441l1.01 5.921l-5.017-2.652V4.479l2.14 4.363a.503.503 0 0 0 .378.275l5.913.863z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForward(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.853 7.147a.5.5 0 0 0-.707.707L11.302 12l-4.156 4.146a.5.5 0 1 0 .707.708l4.51-4.5a.5.5 0 0 0 0-.707zM16.5 7a.5.5 0 0 0-.5.5v9a.5.5 0 1 0 1 0v-9a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 8C18.1 8 17 9.1 17 10.5c0 1.2.9 2.2 2 2.4v2.6c0 3-2.5 5.5-5.5 5.5S8 18.5 8 15.5V15c2.8-.3 5-2.6 5-5.5v-7c0-.3-.2-.5-.5-.5h-2c-.3 0-.5.2-.5.5s.2.5.5.5H12v6.5C12 12 10 14 7.5 14S3 12 3 9.5V3h1.5c.3 0 .5-.2.5-.5S4.8 2 4.5 2h-2c-.3 0-.5.2-.5.5v7c0 2.9 2.2 5.2 5 5.5v.5c0 3.6 2.9 6.5 6.5 6.5s6.5-2.9 6.5-6.5v-2.6c1.1-.2 2-1.2 2-2.4C22 9.1 20.9 8 19.5 8m0 4c-.8 0-1.5-.7-1.5-1.5S18.7 9 19.5 9s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StethoscopeAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 8C18.1 8 17 9.1 17 10.5c0 1.2.9 2.2 2 2.4v2.6c0 3-2.5 5.5-5.5 5.5S8 18.5 8 15.5v-1.8l3.3-2.6C12.4 10.2 13 9 13 7.6V2.5c0-.3-.2-.5-.5-.5h-2c-.3 0-.5.2-.5.5s.2.5.5.5H12v4.6c0 1.1-.5 2.1-1.3 2.7l-3.2 2.5l-3.2-2.5C3.5 9.6 3 8.6 3 7.6V3h1.5c.3 0 .5-.2.5-.5S4.8 2 4.5 2h-2c-.3 0-.5.2-.5.5v5.1c0 1.4.6 2.7 1.7 3.5L7 13.7v1.8c0 3.6 2.9 6.5 6.5 6.5s6.5-2.9 6.5-6.5v-2.6c1.1-.2 2-1.2 2-2.4C22 9.1 20.9 8 19.5 8m0 4c-.8 0-1.5-.7-1.5-1.5S18.7 9 19.5 9s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1m7.993 4.713l1.36-1.36a.5.5 0 0 0-.707-.707l-1.36 1.36C15.876 6.762 14.03 6 12 6s-3.876.762-5.287 2.007l-1.36-1.36a.5.5 0 0 0-.707.707l1.36 1.36A7.96 7.96 0 0 0 4 14a8 8 0 0 0 8 8a8.01 8.01 0 0 0 8-8a7.96 7.96 0 0 0-2.007-5.287M12 21A7 7 0 0 1 7.037 9.065c.005-.005.012-.006.017-.011c.005-.005.006-.012.01-.017A6.977 6.977 0 0 1 12 7a6.986 6.986 0 0 1 4.943 2.049l.008.008A6.986 6.986 0 0 1 19 14a7 7 0 0 1-7 7m.5-8.408V10.5a.5.5 0 0 0-1 0v2.092A1.496 1.496 0 0 0 12 15.5a1.496 1.496 0 0 0 .5-2.908"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StoreSlash(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.8 4h10.9l1.8 4.6c0 1.3-1.1 2.4-2.5 2.4s-2.5-1.1-2.5-2.5c0-.3-.2-.5-.5-.5s-.5.2-.5.5c0 .5-.2 1-.4 1.4c-.1.1-.1.2-.1.3c0 .3.2.5.5.5c.2 0 .3-.1.4-.2c.1-.1.1-.1.1-.2c.6 1 1.7 1.7 3 1.7c.3 0 .7-.1 1-.2v3.4c0 .3.2.5.5.5s.5-.2.5-.5v-3.9c.9-.6 1.5-1.7 1.5-2.9v-.2l-2-5c-.1-.1-.3-.2-.5-.2H7.8c-.3 0-.6.2-.6.5s.3.5.6.5m15.1 18.1l-3-3l-5-5l-13-13c-.2-.1-.5-.1-.7 0c-.2.2-.2.6-.1.8L4 4.7L2.5 8.3v.2c0 1.2.6 2.2 1.5 2.9v10.1c0 .3.2.5.5.5h15c.3 0 .5-.2.5-.5v-.8l2.1 2.1c.1.1.2.1.4.1c.1 0 .3-.1.4-.1c.1-.1.1-.5 0-.7M3.5 8.6l1.2-3.1l3.7 3.7C8.1 10.2 7.1 11 6 11c-1.4 0-2.5-1.1-2.5-2.4m6 5.4c-.3 0-.5.2-.5.5V21H5v-9.2c.3.1.7.2 1 .2c1.3 0 2.4-.7 3-1.7c.5.8 1.3 1.4 2.2 1.6l2.1 2.1zm4.5 7h-4v-6h4zm5 0h-4v-5.3l4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subject(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 8h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1m11 9h-11a.5.5 0 0 0 0 1h11a.5.5 0 0 0 0-1m8-5h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SyncExclamation(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15.625a.625.625 0 1 0 0-1.25a.625.625 0 0 0 0 1.25m8.5.375h-4a.5.5 0 0 0 0 1h2.976a8.996 8.996 0 0 1-7.459 4a9.009 9.009 0 0 1-8.904-10.419a.5.5 0 0 0-.986-.162A9.847 9.847 0 0 0 2 12c.006 5.52 4.48 9.994 10 10a10.009 10.009 0 0 0 8-3.999v2.5a.5.5 0 1 0 1 0v-4.002a.5.5 0 0 0-.5-.499M12 13a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-1 0v3a.5.5 0 0 0 .5.5m4.737-9.804C12.32.818 6.921 2.107 4 6.003V3.5a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 0-1H4.524a8.996 8.996 0 0 1 7.459-4a9.009 9.009 0 0 1 8.903 10.421a.499.499 0 1 0 .987.16A9.847 9.847 0 0 0 22 12a10.008 10.008 0 0 0-5.263-8.804"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SyncSlash(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.853 2.159a.5.5 0 0 0-.707-.013L18.694 4.6a9.999 9.999 0 0 0-1.942-1.4C12.332.817 6.925 2.105 4 6.004V3.5a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 0-1H4.53c.322-.48.673-.94 1.083-1.351c3.398-3.407 8.837-3.518 12.381-.35L5.297 17.996A8.93 8.93 0 0 1 3 12c0-.475.037-.95.113-1.419a.5.5 0 0 0-.986-.162a9.981 9.981 0 0 0 2.467 8.28l-2.448 2.447a.5.5 0 1 0 .707.707l19-19a.5.5 0 0 0 0-.694M20.5 16h-4a.5.5 0 0 0 0 1h2.987c-2.587 3.867-7.75 5.164-11.882 2.854a.5.5 0 0 0-.49.872A9.99 9.99 0 0 0 12 22a10.008 10.008 0 0 0 8-3.999V20.5a.5.5 0 1 0 1 0v-4a.5.5 0 0 0-.5-.5m.226-8.885a.5.5 0 0 0-.872.49a8.995 8.995 0 0 1 1.032 5.816a.499.499 0 1 0 .987.16a9.99 9.99 0 0 0-1.147-6.466"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-13 19H3v-5.5h5.5zm0-6.5H3v-5h5.5zm0-6H3V3h5.5zm6 12.5h-5v-5.5h5zm0-6.5h-5v-5h5zm0-6h-5V3h5zM21 21h-5.5v-5.5H21zm0-6.5h-5.5v-5H21zm0-6h-5.5V3H21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelegramAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.945 2.765a1.552 1.552 0 0 0-1.572-.244L2.456 9.754a1.543 1.543 0 0 0 .078 2.884L6.4 13.98l2.095 6.926c.004.014.017.023.023.036a.486.486 0 0 0 .093.15a.49.49 0 0 0 .226.143c.01.004.017.013.027.015h.006l.003.001a.448.448 0 0 0 .233-.012c.008-.002.016-.002.025-.005a.495.495 0 0 0 .191-.122c.006-.007.016-.008.022-.014l3.013-3.326l4.397 3.405c.267.209.596.322.935.322c.734 0 1.367-.514 1.518-1.231L22.469 4.25a1.533 1.533 0 0 0-.524-1.486M9.588 15.295l-.707 3.437l-1.475-4.878l7.315-3.81l-4.997 4.998a.498.498 0 0 0-.136.253m8.639 4.772a.54.54 0 0 1-.347.399a.525.525 0 0 1-.514-.078l-4.763-3.689a.5.5 0 0 0-.676.06L9.83 19.07l.706-3.427l7.189-7.19a.5.5 0 0 0-.584-.797L6.778 13.054l-3.917-1.362A.526.526 0 0 1 2.5 11.2a.532.532 0 0 1 .334-.518l17.914-7.233a.536.536 0 0 1 .558.086a.523.523 0 0 1 .182.518z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Th(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-10 19H3v-8.5h8.5zm0-9.5H3V3h8.5zM21 21h-8.5v-8.5H21zm0-9.5h-8.5V3H21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThLarge(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.47 3H3.53a.53.53 0 0 0-.53.53v16.94c0 .293.237.53.53.53h16.94a.53.53 0 0 0 .53-.53V3.53a.53.53 0 0 0-.53-.53M11.5 19.941H4.059V12.5H11.5zm0-8.441H4.059V4.059H11.5zm8.441 8.441H12.5V12.5h7.441zm0-8.441H12.5V4.059h7.441z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesCircle(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.012 2.004C6.492 2 2.013 6.474 2.01 11.994S6.48 21.993 12 21.996a9.996 9.996 0 0 0 10.002-9.99c.003-5.52-4.47-9.999-9.99-10.002m0 18.992A8.996 8.996 0 1 1 12 3.004a8.996 8.996 0 0 1 .012 17.992M12.707 12l3.182-3.182a.5.5 0 0 0-.707-.707L12 11.293L8.818 8.111a.5.5 0 0 0-.707.707L11.293 12l-3.182 3.182a.5.5 0 0 0 .707.707L12 12.707l3.182 3.182a.5.5 0 0 0 .707-.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOff(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 9.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5m0 4a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m9-7h-9a5.5 5.5 0 0 0 0 11h9a5.5 5.5 0 0 0 0-11m0 10h-9a4.5 4.5 0 1 1 0-9h9a4.5 4.5 0 1 1 0 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOn(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 6.5h-9a5.5 5.5 0 0 0 0 11h9a5.5 5.5 0 0 0 0-11m0 10h-9a4.5 4.5 0 1 1 0-9h9a4.5 4.5 0 1 1 0 9m0-7a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5m0 4a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToiletPaper(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.4 21.2l-.6-.8c-1.1-1.3-1.7-3-1.7-4.8V7.5c0-3-2.5-5.5-5.5-5.5H6C3.2 2 1 4.7 1 8s2.2 6 5 6h4v1.6c0 2 .7 3.9 2 5.4l.6.8c.1.1.2.2.4.2h9c.1 0 .2 0 .3-.1c.2-.2.3-.5.1-.7M10 13H8.7c.5-.4.9-.9 1.3-1.4zm-4 0c-2.2 0-4-2.2-4-5c0-2.7 1.8-5 4-5h.1C8.2 3 10 5.3 10 8c0 2.8-1.8 5-4 5m7.2 8l-.5-.6c-1.1-1.3-1.7-3-1.7-4.8V8c0-2.1-.9-3.9-2.2-4.9h5.7C17 3 19 5 19 7.5v8.1c0 2 .7 3.9 1.9 5.4zM6 6.2c-.9.1-1.6.9-1.5 1.8c-.1.9.6 1.7 1.5 1.7S7.6 8.9 7.5 8c.1-.9-.6-1.7-1.5-1.8m0 2.5c-.3 0-.5-.3-.5-.7s.2-.8.5-.8s.5.4.5.8s-.2.7-.5.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrafficLight(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0 7a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2M23 4.5a.5.5 0 0 0-.5-.5h-4.591a2.65 2.65 0 0 0-2.561-2H8.652a2.65 2.65 0 0 0-2.561 2H1.5a.5.5 0 0 0-.5.5c0 1.862.929 3.505 2.345 4.5H1.5a.5.5 0 0 0-.5.5c0 1.862.929 3.505 2.345 4.5H1.5a.5.5 0 0 0-.5.5a5.498 5.498 0 0 0 5.084 5.479A2.65 2.65 0 0 0 8.652 22h6.696a2.65 2.65 0 0 0 2.568-2.021A5.498 5.498 0 0 0 23 14.5a.5.5 0 0 0-.5-.5h-1.845A5.492 5.492 0 0 0 23 9.5a.5.5 0 0 0-.5-.5h-1.845A5.492 5.492 0 0 0 23 4.5M6 18.972A4.5 4.5 0 0 1 2.028 15H6zm0-5A4.5 4.5 0 0 1 2.028 10H6zm0-5A4.5 4.5 0 0 1 2.028 5H6zM17 9.5v9.849A1.654 1.654 0 0 1 15.348 21H8.652A1.654 1.654 0 0 1 7 19.348V4.652A1.654 1.654 0 0 1 8.652 3h6.696c.912.001 1.651.74 1.652 1.652zm4.972 5.5A4.5 4.5 0 0 1 18 18.972V15zm0-5A4.5 4.5 0 0 1 18 13.972V10zM18 8.972V5h3.972A4.5 4.5 0 0 1 18 8.972M12 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.934 20.006l-9.5-16.51a.521.521 0 0 0-.868 0l-9.5 16.51a.5.5 0 0 0 .434.749h19a.5.5 0 0 0 .434-.75m-18.57-.251L12 4.748l8.636 15.007z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumblrAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.969 20.667l-.962-2.853a.5.5 0 0 0-.689-.29c-.45.177-.928.278-1.413.296a1.348 1.348 0 0 1-1.017-.303a1.685 1.685 0 0 1-.395-1.17V11.5h3.517a.5.5 0 0 0 .5-.5V7.096a.5.5 0 0 0-.5-.5h-3.5V2a.5.5 0 0 0-.5-.5H9.08a.65.65 0 0 0-.64.595a5.569 5.569 0 0 1-3.61 4.973a.5.5 0 0 0-.33.471v3.583a.5.5 0 0 0 .5.5h1.524v5.035a5.9 5.9 0 0 0 1.739 4.123a6.415 6.415 0 0 0 4.624 1.72l.154-.001c1.572-.027 3.38-.677 3.867-1.39a.5.5 0 0 0 .06-.442m-3.946.832a5.403 5.403 0 0 1-4.06-1.432a4.882 4.882 0 0 1-1.439-3.41v-5.535a.5.5 0 0 0-.499-.5H5.5V7.881A6.57 6.57 0 0 0 9.395 2.5h2.114v4.595a.5.5 0 0 0 .5.5h3.5V10.5h-3.516a.5.5 0 0 0-.5.5v5.356a2.594 2.594 0 0 0 .703 1.883a2.332 2.332 0 0 0 1.737.581a5.622 5.622 0 0 0 1.289-.205l.7 2.08c-.88.513-1.879.79-2.899.804"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.5 3.589a.5.5 0 0 0-.755-.43a7.938 7.938 0 0 1-2.266.912a4.662 4.662 0 0 0-3.238-1.29a4.731 4.731 0 0 0-4.707 5.135a11.527 11.527 0 0 1-7.717-4.18a.5.5 0 0 0-.82.067a4.777 4.777 0 0 0-.633 2.377a4.724 4.724 0 0 0 .762 2.579l-.06-.033a.504.504 0 0 0-.497.03a.543.543 0 0 0-.247.458c-.004.118.003.237.022.353a4.692 4.692 0 0 0 1.818 3.383a.5.5 0 0 0-.335.632a4.704 4.704 0 0 0 3.088 3.057a7.998 7.998 0 0 1-4.854.963a.5.5 0 0 0-.332.917A12.442 12.442 0 0 0 8.468 20.5a12.299 12.299 0 0 0 11.986-9.006c.339-1.137.512-2.318.514-3.505c0-.12 0-.245-.003-.372A5.37 5.37 0 0 0 22.5 3.59m-2.424 3.533a.498.498 0 0 0-.117.349a11.366 11.366 0 0 1-.464 3.741A11.174 11.174 0 0 1 8.468 19.5a11.45 11.45 0 0 1-4.443-.897a8.867 8.867 0 0 0 4.525-1.86a.5.5 0 0 0-.3-.893A3.71 3.71 0 0 1 5.1 14c.425.001.847-.057 1.254-.174a.5.5 0 0 0-.042-.97a3.706 3.706 0 0 1-2.905-2.898a4.72 4.72 0 0 0 1.313.228a.473.473 0 0 0 .492-.35a.5.5 0 0 0-.2-.567a3.696 3.696 0 0 1-1.648-3.09c0-.413.067-.823.2-1.213a12.515 12.515 0 0 0 8.54 3.995a.45.45 0 0 0 .409-.179a.5.5 0 0 0 .103-.434a3.642 3.642 0 0 1-.1-.842A3.73 3.73 0 0 1 16.24 3.78a3.68 3.68 0 0 1 2.71 1.179a.499.499 0 0 0 .462.148a8.94 8.94 0 0 0 2.055-.671a4.92 4.92 0 0 1-1.392 2.686"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 3.012V2.5a.5.5 0 0 0-1 0v.512C6.898 3.149 3.142 6.84 3 11.5a.5.5 0 0 0 .5.5h8v7.25a1.75 1.75 0 0 1-3.5 0a.5.5 0 0 0-1 0A2.753 2.753 0 0 0 9.75 22a2.753 2.753 0 0 0 2.75-2.75V12h8a.5.5 0 0 0 .5-.5c-.142-4.66-3.898-8.35-8.5-8.488M8.05 11H4.019a7.81 7.81 0 0 1 5.985-6.79C8.905 5.686 8.229 8.318 8.05 11m1.002 0C9.345 6.664 10.776 4 12 4s2.656 2.664 2.948 7zm6.898 0c-.179-2.682-.854-5.314-1.952-6.79A7.81 7.81 0 0 1 19.982 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 9h-1.5a.5.5 0 0 0 0 1H18a2.003 2.003 0 0 1 2 2v7a2.003 2.003 0 0 1-2 2H6a2.003 2.003 0 0 1-2-2v-7a2.003 2.003 0 0 1 2-2h2.5a.5.5 0 0 0 0-1H6a3.003 3.003 0 0 0-3 3v7a3.003 3.003 0 0 0 3 3h12a3.003 3.003 0 0 0 3-3v-7a3.003 3.003 0 0 0-3-3M8.862 6.345L11.5 3.707v13.794a.5.5 0 1 0 1 0V3.706l2.638 2.638a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707l-3.491-3.491a.5.5 0 0 0-.344-.145L12 2l-.007.001a.498.498 0 0 0-.347.146l-3.49 3.49a.5.5 0 1 0 .707.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserArrows(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.6 6.9l2.5 2.5c.1.1.2.1.4.1c.1 0 .3-.1.4-.1c.2-.2.2-.5 0-.7L8.2 7h7.6l-1.6 1.6c-.1.1-.2.3-.2.4c0 .3.2.5.5.5c.1 0 .3-.1.4-.1l2.5-2.5c.2-.2.2-.5 0-.7l-2.5-2.5c-.2-.2-.5-.2-.7 0c-.2.2-.2.5 0 .7L15.8 6H8.2l1.6-1.6c.2-.2.2-.5 0-.7c-.2-.2-.5-.2-.7 0L6.6 6.1s-.1.1-.1.2v.4c0 .1.1.1.1.2m1.3 9c.8-.6 1.4-1.5 1.4-2.6C9.2 11.5 7.8 10 6 10s-3.2 1.5-3.2 3.2c0 1.1.5 2 1.4 2.6c-1.9.8-3.2 2.6-3.2 4.7c0 .3.2.5.5.5s.5-.2.5-.5c0-2.2 1.8-4 4-4s4 1.8 4 4c0 .3.2.5.5.5s.5-.2.5-.5c0-2.1-1.3-3.9-3.1-4.6M6 15.5c-1.2 0-2.2-1-2.2-2.2S4.8 11 6 11c1.2 0 2.2 1 2.2 2.2c0 1.3-1 2.3-2.2 2.3m13.9.4c.8-.6 1.4-1.5 1.4-2.6c0-1.8-1.5-3.2-3.2-3.2s-3.2 1.5-3.2 3.2c0 1.1.5 2 1.4 2.6c-1.8.7-3.1 2.5-3.1 4.6c0 .3.2.5.5.5s.5-.2.5-.5c0-2.2 1.8-4 4-4s4 1.8 4 4c0 .3.2.5.5.5s.5-.2.5-.5c-.2-2.1-1.5-3.9-3.3-4.6m-1.9-.4c-1.2 0-2.2-1-2.2-2.2s1-2.2 2.2-2.2c1.2 0 2.2 1 2.2 2.2c0 1.2-1 2.2-2.2 2.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorSquare(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 16.05v-8.1A2.998 2.998 0 0 0 22 5a2.994 2.994 0 0 0-5.95-.5h-8.1A2.994 2.994 0 0 0 2 5a2.994 2.994 0 0 0 2.5 2.95v8.1A2.994 2.994 0 0 0 5 22a2.998 2.998 0 0 0 2.95-2.5h8.1A2.994 2.994 0 0 0 22 19a2.994 2.994 0 0 0-2.5-2.95M19 3a2 2 0 1 1 0 4a2 2 0 0 1 0-4M3 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0m2 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4m11.05-2.5h-8.1a2.99 2.99 0 0 0-2.45-2.45v-8.1A2.993 2.993 0 0 0 7.95 5.5h8.1a2.99 2.99 0 0 0 2.45 2.45v8.1a2.99 2.99 0 0 0-2.45 2.45M19 21a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorSquareAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 7h-9a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5m-.5 9H8V8h8zm5 3.092V4.908c.58-.207 1-.757 1-1.408c0-.827-.673-1.5-1.5-1.5a1.5 1.5 0 0 0-1.408 1H4.908A1.5 1.5 0 0 0 3.5 2C2.673 2 2 2.673 2 3.5c0 .651.42 1.2 1 1.408v14.184c-.58.207-1 .757-1 1.408c0 .827.673 1.5 1.5 1.5a1.5 1.5 0 0 0 1.408-1h14.184c.207.58.757 1 1.408 1c.827 0 1.5-.673 1.5-1.5a1.5 1.5 0 0 0-1-1.408M20.5 3a.501.501 0 0 1 0 1a.501.501 0 0 1 0-1m-17 0a.501.501 0 0 1 0 1a.501.501 0 0 1 0-1m0 18a.501.501 0 0 1 0-1a.501.501 0 0 1 0 1m15.59-.994L19.086 20H4.914l-.004.006a1.498 1.498 0 0 0-.916-.916L4 19.086V4.914l-.006-.004c.428-.15.765-.488.916-.916L4.914 4h14.172l.004-.006c.15.428.488.765.916.916L20 4.914v14.172l.006.004c-.428.15-.765.488-.916.916M20.5 21a.501.501 0 0 1 0-1a.501.501 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VkAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.375 13.867a5.034 5.034 0 0 0-1.552-1.13a7.363 7.363 0 0 0 3.315-4.627a.5.5 0 1 0-.975-.22a6.362 6.362 0 0 1-3.68 4.428c-.109-.018-.22-.023-.329-.034V8.002a.487.487 0 0 0-.026-.129a.526.526 0 0 0-.013-.064a.5.5 0 0 0-.095-.143c-.005-.005-.007-.012-.011-.016a.501.501 0 0 0-.12-.082c-.014-.008-.025-.02-.039-.026a.513.513 0 0 0-.103-.021a.48.48 0 0 0-.09-.019l-.655-.002H11a.5.5 0 0 0-.002 1h.152v5.46c0 .007.004.013.004.02v.998a11.792 11.792 0 0 1-4.007-7.059a.51.51 0 0 0-.576-.41a.5.5 0 0 0-.41.576a12.838 12.838 0 0 0 5.2 8.322c.01.007.022.004.031.01a.486.486 0 0 0 .262.085a.492.492 0 0 0 .212-.051l.016-.004c.01-.005.016-.018.026-.024a.49.49 0 0 0 .18-.186a.487.487 0 0 0 .066-.235v-2.74a4.06 4.06 0 0 1 2.485 1.281l1.647 1.797a.499.499 0 1 0 .736-.676zM15.073 1.5H8.938C3.029 1.5 1.5 3.027 1.5 8.927v6.136c0 5.908 1.527 7.437 7.427 7.437h6.136c5.908 0 7.437-1.527 7.437-7.427V8.938c0-5.909-1.527-7.438-7.427-7.438M21.5 15.073c0 5.346-1.083 6.427-6.437 6.427H8.927C3.58 21.5 2.5 20.417 2.5 15.063V8.927C2.5 3.58 3.583 2.5 8.938 2.5h6.135c5.346 0 6.427 1.083 6.427 6.438z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VuejsAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.934 3.256a.499.499 0 0 0-.434-.251l-4.99-.003L17.503 3l-4-.026H13.5a.498.498 0 0 0-.43.245l-1.072 1.805l-1.07-1.78A.498.498 0 0 0 10.505 3l-4-.027H6.5A.48.48 0 0 0 6.399 3H1.5a.5.5 0 0 0-.432.752l10.5 18a.5.5 0 0 0 .864 0l10.5-17.995a.5.5 0 0 0 .002-.501m-12.718.742l1.355 2.259A.5.5 0 0 0 12 6.5h.001a.5.5 0 0 0 .429-.245l1.353-2.28l2.83.02l-3.006 4.917L12 11.54L7.394 3.979zM12 20.508L2.37 4h3.85l5.353 8.76a.493.493 0 0 0 .147.142c.014.01.021.026.035.034a.5.5 0 0 0 .672-.175l5.353-8.759l3.85.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 7H18V6a3.003 3.003 0 0 0-3-3H4.5A2.502 2.502 0 0 0 2 5.5V18a3.003 3.003 0 0 0 3 3h14.5a2.502 2.502 0 0 0 2.5-2.5v-9A2.502 2.502 0 0 0 19.5 7m-15-3H15a2.003 2.003 0 0 1 2 2v1H4.5a1.5 1.5 0 1 1 0-3M21 16h-2a2 2 0 0 1 0-4h2zm0-5h-2a3 3 0 1 0 0 6h2v1.5a1.5 1.5 0 0 1-1.5 1.5H5a2.003 2.003 0 0 1-2-2V7.499c.432.326.959.502 1.5.501h15A1.5 1.5 0 0 1 21 9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebGrid(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-7 19H3v-8.5h11.5zm0-9.5H3V3h11.5zM21 21h-5.5V3H21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebGridAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-10 19H3V9.5h8.5zm9.5 0h-8.5V9.5H21zm0-12.5H3V3h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebSection(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-7 19H3V3h11.5zm6.5 0h-5.5V3H21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebSectionAlt(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-13 19H3V3h5.5zM21 21H9.5V3H21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowGrid(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-13 19H3V3h5.5zM21 21H9.5v-8.5H21zm0-9.5H9.5V3H21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMaximize(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5M21 21H3V9.5h18zm0-12.5H3V3h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowSection(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 2h-19a.5.5 0 0 0-.5.5v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5m-13 19H3V9.5h5.5zm6 0h-5V9.5h5zm6.5 0h-5.5V9.5H21zm0-12.5H3V3h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrapText(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 7h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1M19 11H2.5a.5.5 0 0 0 0 1H19a2 2 0 0 1 0 4h-5.074l1.386-1.11a.5.5 0 1 0-.624-.78l-2.5 2c-.011.008-.02.019-.03.028a.538.538 0 0 0-.02.02c-.009.01-.02.018-.028.029c-.014.018-.019.039-.03.058a.48.48 0 0 0-.047.09a.45.45 0 0 0-.02.102c-.003.021-.013.04-.013.063c0 .01.005.017.005.026c.002.036.012.068.02.102c.009.03.015.06.028.087c.013.027.033.05.051.076a.49.49 0 0 0 .065.078c.008.006.01.016.019.022l2.5 2a.502.502 0 0 0 .703-.079a.5.5 0 0 0-.078-.703L13.926 17H19a3 3 0 1 0 0-6m-9.5 5h-7a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *UitIcon {
	return &UitIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.792 10.775l-3.668-2.112A1.417 1.417 0 0 0 9 9.89v4.222c-.003.506.267.974.706 1.224a1.41 1.41 0 0 0 1.419.002l3.667-2.112a1.413 1.413 0 0 0 0-2.45m-.5 1.582l-3.666 2.113a.414.414 0 0 1-.419 0A.408.408 0 0 1 10 14.11V9.89a.408.408 0 0 1 .207-.359a.402.402 0 0 1 .418 0l3.667 2.113a.41.41 0 0 1 0 .714M19 4H5a3.003 3.003 0 0 0-3 3v10a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V7a3.003 3.003 0 0 0-3-3m2 13a2.003 2.003 0 0 1-2 2H5a2.003 2.003 0 0 1-2-2V7a2.003 2.003 0 0 1 2-2h14a2.003 2.003 0 0 1 2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
