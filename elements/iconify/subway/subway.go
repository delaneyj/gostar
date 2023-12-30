package subway

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 512 512"
	fill           = "currentColor"
)

type SubwayIcon struct {
	*SVGSVGElement
}

func Add(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m149.3 277.3c0 11.8-9.5 21.3-21.3 21.3h-85.3V384c0 11.8-9.5 21.3-21.3 21.3h-42.7c-11.8 0-21.3-9.6-21.3-21.3v-85.3H128c-11.8 0-21.3-9.6-21.3-21.3v-42.7c0-11.8 9.5-21.3 21.3-21.3h85.3V128c0-11.8 9.5-21.3 21.3-21.3h42.7c11.8 0 21.3 9.6 21.3 21.3v85.3H384c11.8 0 21.3 9.6 21.3 21.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M298.7 213.3V0h-85.4v213.3H0v85.4h213.3V512h85.4V298.7H512v-85.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddPlaylist(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M93.1 325.8V139.6H46.5C20.9 139.6 0 160.5 0 186.2v279.3C0 491.1 20.9 512 46.5 512h279.3c25.7 0 46.5-20.9 46.5-46.5V419H186.2c-51.4-.1-93.1-41.8-93.1-93.2M465.5 0H186.2c-25.7 0-46.5 20.9-46.5 46.5v279.3c0 25.7 20.9 46.5 46.5 46.5h279.3c25.7 0 46.5-20.9 46.5-46.5V46.5C512 20.9 491.1 0 465.5 0m-23.3 209.5h-93.1v93.1h-46.5v-93.1h-93.1V163h93.1V69.8h46.5v93.1h93.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Admin(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m448 362.7l-117.3-21.3C320 320 320 310.7 320 298.7c10.7-10.7 32-21.3 32-32c10.7-32 10.7-53.3 10.7-53.3c5.5-8 21.3-21.3 21.3-42.7s-21.3-42.7-21.3-53.3C362.7 32 319.2 0 256 0c-60.5 0-106.7 32-106.7 117.3c0 10.7-21.3 32-21.3 53.3s15.2 35.4 21.3 42.7c0 0 0 21.3 10.7 53.3c0 10.7 21.3 21.3 32 32c0 10.7 0 21.3-10.7 42.7L64 362.7C21.3 373.3 0 448 0 512h512c0-64-21.3-138.7-64-149.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdminOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.3 384c0-87 65.2-158.7 149.3-169.2v-1.5c5.5-8 21.3-21.3 21.3-42.7s-21.3-42.7-21.3-53.3C362.7 32 319.2 0 256 0c-60.5 0-106.7 32-106.7 117.3c0 10.7-21.3 32-21.3 53.3s15.2 35.4 21.3 42.7c0 0 0 21.3 10.7 53.3c0 10.7 21.3 21.3 32 32c0 10.7 0 21.3-10.7 42.7L64 362.7C21.3 373.3 0 448 0 512h271.4c-35.5-31.3-58.1-77-58.1-128M384 256c-70.7 0-128 57.3-128 128s57.3 128 128 128s128-57.3 128-128s-57.3-128-128-128m85.3 149.3h-64v64h-42.7v-64h-64v-42.7h64v-64h42.7v64h64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdminTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 256c-70.7 0-128 57.3-128 128s57.3 128 128 128s128-57.3 128-128s-57.3-128-128-128m85.3 149.3H298.7v-42.7h170.7v42.7zm-256-21.3c0-87 65.2-158.7 149.3-169.2v-1.5c5.5-8 21.3-21.3 21.3-42.7s-21.3-42.7-21.3-53.3C362.7 32 319.2 0 256 0c-60.5 0-106.7 32-106.7 117.3c0 10.7-21.3 32-21.3 53.3s15.2 35.4 21.3 42.7c0 0 0 21.3 10.7 53.3c0 10.7 21.3 21.3 32 32c0 10.7 0 21.3-10.7 42.7L64 362.7C21.3 373.3 0 448 0 512h271.4c-35.5-31.3-58.1-77-58.1-128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneMode(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M490.7 0c-21.3 0-42.7 0-53.3 10.7l-85.3 96H21.3L0 138.7l234.7 74.7l-64 85.3h-96l-32 32l64 21.3v32c.7 11.5 10.7 21.3 21.3 21.3h32l21.3 64l32-32v-96l85.3-64L373.3 512l32-21.3V160l96-85.3C512 64 512 42.7 512 21.3C512 10.7 501.3 0 490.7 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alam(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M251 85.3c-117.8 0-213.3 95.5-213.3 213.3C37.7 416.5 133.2 512 251 512c117.8 0 213.3-95.5 213.3-213.3c0-117.8-95.5-213.4-213.3-213.4M357.7 320h-128V149.3h42.7v128h85.3zm100-171.9c17.1-15.6 27.9-37.8 27.9-62.8c0-47.1-38.2-85.3-85.3-85.3c-35.6 0-66 21.8-78.8 52.7c55.5 15.9 103.1 50 136.2 95.4M180.5 52.7C167.6 21.8 137.2 0 101.7 0C54.5 0 16.3 38.2 16.3 85.3c0 24.9 10.9 47.2 27.9 62.8c33.2-45.4 80.8-79.5 136.3-95.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M444.4 60.1C399.4 20 340.6 0 267.8 0C189.2 0 125 25.8 75.1 77.4C25.1 129.1.2 192.7.2 268.5C.2 342.8 23.6 402 70.4 446c46.8 44 107.5 66 182 66c54.4 0 98.8-7.6 133.3-22.8v-59.6c-37.8 17.9-79.4 26.8-124.9 26.8c-59.6 0-106.3-16.6-140.2-50c-33.8-33.3-50.8-78.9-50.8-136.7c0-59.7 19.1-110.4 57.4-151.9s86.3-62.3 144.1-62.3c54 0 96 15.6 126.1 46.7s45.2 71 45.2 119.7c0 35.6-5.2 64.4-15.8 86.5c-10.5 22.1-23.8 33.1-40 33.1c-15.3 0-23.1-14.6-23.1-43.7c0-27.3 4.5-85.2 13.7-173.9h-68.2l-3 30.9H305c-7.7-24.4-26.6-36.5-56.7-36.5c-33.3 0-61.8 14.9-85.5 44.8c-23.6 29.9-35.5 68.7-35.5 116.3c0 36.9 8.5 65.7 25.6 86.5c17.1 20.8 39.5 31.2 67.4 31.2c37.2 0 62.8-19.9 77-59.6h1.8c1 17.7 7.8 32.1 20.3 43.1c12.5 11.1 28.8 16.6 48.8 16.6c40.6 0 74.7-16.3 102.1-48.9c27.5-32.6 41.2-75.8 41.2-129.7c.3-65.6-22.2-118.5-67.1-158.5M284.2 310.6c-10.9 21-25.9 31.6-44.9 31.6c-12.3 0-22.4-5.6-30.2-16.7c-7.8-11.1-11.7-26.7-11.7-46.6c0-28.1 5.7-52.8 17.1-74.2c11.4-21.3 26.2-32 44.4-32c12.5 0 22.6 4.8 30.2 14.4c7.6 9.6 11.4 22.7 11.4 39.4c.1 35-5.4 63.1-16.3 84.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m257 271.5l256 128v-256zm-256 0l256 128v-256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackwardOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472.8 90L298.5 192.5v-74c0-28.4-17.6-41.2-39.2-28.5L16.2 232.9c-21.6 12.7-21.6 33.4 0 46.1l243.1 143c21.6 12.7 39.2-.2 39.2-28.5v-74L472.8 422c21.6 12.7 39.2-.2 39.2-28.5v-275c0-28.4-17.6-41.2-39.2-28.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bag(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M421.2 128h-42.7v21.3c0 23.5-19.1 42.7-42.7 42.7c-23.5 0-42.7-19.1-42.7-42.7V128h-85.3v21.3c0 23.5-19.1 42.7-42.7 42.7c-23.5 0-42.7-19.1-42.7-42.7V128H79.9c0 213.3-21.3 384-21.3 384h384c-.1 0-21.4-170.7-21.4-384m-256 42.7c11.8 0 21.3-9.5 21.3-21.3v-42.7c0-35.4 28.6-64 64-64s64 28.6 64 64v42.7c0 11.8 9.5 21.3 21.3 21.3s21.3-9.5 21.3-21.3v-42.7C357.2 47.8 309.4 0 250.5 0c-58.9 0-106.7 47.8-106.7 106.7v42.7c.1 11.7 9.6 21.3 21.4 21.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basket(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405.3 192L320 0h-42.7l85.3 192zM234.7 0H192l-85.3 192h42.7zm-192 469.3c0 23.5 19.1 42.7 42.7 42.7h341.3c23.5 0 42.7-19.1 42.7-42.7l21.3-192H21.3zm320-149.3h42.7L384 469.3h-42.7zm-128 0h42.7v149.3h-42.7zm-85.4 0l21.3 149.3H128L106.7 320zm341.4-106.7H21.3C9.5 213.3 0 222.9 0 234.7V256h512v-21.3c0-11.8-9.5-21.4-21.3-21.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M419.7 340.3V169.6C419.7 75.4 343.3-1 249-1S78.3 75.4 78.3 169.7v170.7c0 42.7-42.7 85.3-42.7 85.3h426.7s-42.6-42.7-42.6-85.4M249 511c35.4 0 64-19.1 64-42.7H185c0 23.6 28.6 42.7 64 42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlackWhite(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 469.3V42.7c117.8 0 213.3 95.5 213.3 213.3c0 117.8-95.5 213.3-213.3 213.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M375.4 157.5L217.8 0v196.9L139 118.1l-39.4 39.4l98.5 98.5l-98.5 98.5l39.4 39.4l78.8-78.8V512l157.5-157.5l-98.4-98.5zm-118.2-59l59.1 59.1l-59.1 59.1zm59.1 256l-59.1 59.1V295.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blur(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M391.3 213.4C343.4 134.8 296.4 106.7 258.5 0c-38.1 106.7-84.9 134.8-132.8 213.4c-43.3 71-56.9 170.7 0 234.7c37.9 42.7 76 64.1 132.8 63.9c57.1.2 94.9-21.2 132.8-63.9c56.9-64 43.3-163.7 0-234.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 124.7L256 18L0 124.7l256 106.7zM256 274l-144.9-67.6L0 252.7l256 106.7l256-106.7l-111.1-46.3zm0 128l-139.6-69.8L0 380.7l256 106.7l256-106.7l-116.4-48.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M277.3 50.2v426.7c64-53.3 170.7-53.3 234.7 0V50.2c-64-53.3-170.7-53.3-234.7 0M0 50.2v426.7c64-53.3 170.7-53.3 234.7 0V50.2C170.7-3.1 64-3.1 0 50.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M275.2 512L480 409.6l20.5-307.2l-225.3 61.4zM29.5 409.6L234.3 512V163.8L9 102.4zM254.8 0L9 61.4l245.8 61.4l245.8-61.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M365.7 102.4v146.3H512V102.4zm109.7 109.7h-73.1V139h73.1zM0 431.6h146.3V285.3H0zm36.6-109.7h73.1V395H36.6zM0 248.7h146.3V102.4H0zM36.6 139h73.1v73.1H36.6zm146.3 109.7h146.3V102.4H182.9zM219.4 139h73.1v73.1h-73.1zm-36.5 292.6h146.3V285.3H182.9zm36.5-109.7h73.1V395h-73.1zm146.3 109.7H512V285.3H365.7zm36.6-109.7h73.1V395h-73.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brightest(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 192c-35.4 0-64 28.6-64 64s28.6 64 64 64s64-28.6 64-64s-28.6-64-64-64m0-42.7c23.5 0 42.7-19.1 42.7-42.7v-64C298.7 19.1 279.5 0 256 0c-23.5 0-42.7 19.1-42.7 42.7v64c0 23.5 19.2 42.6 42.7 42.6m0 213.4c-23.5 0-42.7 19.1-42.7 42.7v64c0 23.5 19.1 42.7 42.7 42.7c23.5 0 42.7-19.1 42.7-42.7v-64c0-23.6-19.2-42.7-42.7-42.7m213.3-149.4h-64c-23.5 0-42.7 19.1-42.7 42.7c0 23.5 19.1 42.7 42.7 42.7h64c23.5 0 42.7-19.1 42.7-42.7c0-23.5-19.1-42.7-42.7-42.7m-320 42.7c0-23.5-19.1-42.7-42.7-42.7h-64C19.1 213.3 0 232.5 0 256c0 23.5 19.1 42.7 42.7 42.7h64c23.5 0 42.6-19.2 42.6-42.7m242.5 75.4c-16.7-16.6-43.7-16.6-60.4 0c-16.6 16.7-16.6 43.7 0 60.4l45.3 45.3c16.6 16.6 43.7 16.6 60.3 0c16.6-16.7 16.6-43.7 0-60.3zM120.2 180.6c16.7 16.6 43.7 16.6 60.4 0c16.6-16.7 16.6-43.7 0-60.4L135.3 75C118.7 58.3 91.6 58.3 75 75c-16.6 16.7-16.6 43.7 0 60.3zm271.6 0l45.3-45.3c16.6-16.6 16.6-43.7 0-60.3c-16.7-16.6-43.7-16.6-60.3 0l-45.3 45.3c-16.6 16.7-16.6 43.7 0 60.4c16.6 16.5 43.6 16.5 60.3-.1M120.2 331.4L75 376.7c-16.6 16.6-16.6 43.7 0 60.3c16.7 16.6 43.7 16.6 60.3 0l45.3-45.3c16.6-16.7 16.6-43.7 0-60.4c-16.7-16.5-43.7-16.5-60.4.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79.1 372.5C59.4 431.9 49.5 451.7 0 451.7c168.2 69.3 197.9-49.5 197.9-49.5l-49.5-49.5s-49.5-9.8-69.3 19.8M504.6 46.1c-14-14-38-3.9-38-3.9c-35 7-229.1 201.8-229.1 201.8l69.3 69.3S501.6 119.2 508.6 84.2c-.1-.1 10-24.2-4-38.1M168.2 333l49.5 49.5L287 333l-69.3-69.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func C(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M69.8 0C31.3 0 0 31.2 0 69.8c0 38.6 31.3 69.8 69.8 69.8c38.6 0 69.8-31.3 69.8-69.8C139.6 31.2 108.4 0 69.8 0m0 104.7c-19.3 0-34.9-15.6-34.9-34.9s15.6-34.9 34.9-34.9c19.3 0 34.9 15.6 34.9 34.9s-15.6 34.9-34.9 34.9m325.8-11.6c40.9 0 84.4 15.1 116.4 34.9V23.3C480 11.1 442.3 0 395.6 0c-75.3 0-137.8 24.7-185 73.9c-47.2 49.2-71 114.7-71 193.8c0 74.3 21.2 132.5 63.2 177.1c42 44.6 96.3 67.3 169.5 67.3c56.2 0 104.8-5.8 139.6-23.3V384c-34.9 21.2-74.4 34.9-116.4 34.9c-44 0-78.7-18.1-104.7-46.5s-34.9-56.9-34.9-104.7c0-49.9 7.4-98.8 34.9-128c27.7-29.3 60-46.6 104.8-46.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cain(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.7 426.7s-42.7 42.7-64 42.7c-35.4 0-64-28.6-64-64c0-21.3 42.7-64 42.7-64l83.4-83.4c-6.3-1.2-12.7-2-19.4-2c-24.4 0-46.1 7.5-64 21.3l-64 64C8 359 0 381.5 0 405.3C0 464.2 47.8 512 106.7 512c24 0 46.3-7.9 64-21.3l64-64c13.6-17.7 21.3-39.9 21.3-64c0-6.7-.8-13.1-2-19.4zM405.3 0c-24 0-46.3 7.9-64 21.3l-64 64c-13.6 17.7-21.3 39.9-21.3 64c0 6.7.8 13.1 2 19.4l83.4-83.4s42.7-42.7 64-42.7c35.4 0 64 28.6 64 64c0 21.3-42.7 64-42.7 64L343.3 254c6.3 1.2 12.8 2 19.4 2c24.4 0 46.1-7.5 64-21.3l64-64c13.3-17.6 21.3-40.1 21.3-64C512 47.8 464.2 0 405.3 0m-256 330.7c0 17.7 14.3 32 32 32c12.6 0 23.4-7.4 28.6-18L344.7 210c10.6-5.2 18-16 18-28.6c0-17.7-14.3-32-32-32c-12.6 0-23.4 7.4-28.6 18L167.3 302c-10.6 5.3-18 16.1-18 28.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M362.7 320H320v42.7h42.7zm0-85.3H320v42.7h42.7zm0-85.4H320V192h42.7zm-85.4 85.4h-42.7v42.7h42.7zm0-85.4h-42.7V192h42.7zm0 170.7h-42.7v42.7h42.7zM448 149.3h-42.7V192H448zm0 85.4h-42.7v42.7H448zm0 85.3h-42.7v42.7H448zm-341.3 85.3H64V448h42.7zm0-85.3H64v42.7h42.7zm170.6 85.3h-42.7V448h42.7zm-85.3 0h-42.7V448H192zM0 0v512h512V0zm469.3 469.3H42.7V128h426.7v341.3zM106.7 234.7H64v42.7h42.7zM192 320h-42.7v42.7H192zm0-170.7h-42.7V192H192zm0 85.4h-42.7v42.7H192zm170.7 170.6H320V448h42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149.3 0h-42.7v64h42.7zm277.4 0v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0zm42.6 469.3H42.7V128h426.7v341.3zM405.3 0h-42.7v64h42.7zm-42.6 277.3H149.3V320h213.3v-42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149.3 0h-42.7v64h42.7zm256 0h-42.7v64h42.7zm21.4 0v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0zm42.6 469.3H42.7V128h426.7v341.3zm-234.6-64h42.7V320h85.3v-42.7h-85.3V192h-42.7v85.3h-85.3V320h85.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M362.7 320H320v42.7h42.7zm-85.4-170.7h-42.7V192h42.7zm85.4 256H320V448h42.7zm0-256H320V192h42.7zM149.3 0h-42.7v64h42.7zM448 149.3h-42.7V192H448zm-170.7 85.4h-42.7v42.7h42.7zm170.7 0h-42.7v42.7H448zm0 85.3h-42.7v42.7H448zm-85.3-85.3H320v42.7h42.7zm-256 0H64v42.7h42.7zm0 85.3H64v42.7h42.7zm320-320v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0zm42.6 469.3H42.7V128h426.7v341.3zm-362.6-64H64V448h42.7zM277.3 320h-42.7v42.7h42.7zM192 405.3h-42.7V448H192zm85.3 0h-42.7V448h42.7zm-85.3-256h-42.7V192H192zm0 85.4h-42.7v42.7H192zm0 85.3h-42.7v42.7H192zM405.3 0h-42.7v64h42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSix(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149.3 0h-42.7v64h42.7zm256 0h-42.7v64h42.7zM148.9 426.7h214.3c33 0 52.3-22.2 35.9-49.4L286 190.9c-16.5-27.2-43.4-27.2-59.9 0L113 377.3c-16.5 27.1 2.9 49.4 35.9 49.4m85.8-192h42.7v106.7h-42.7zm0 128h42.7v42.7h-42.7zM426.7 0v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0zm42.6 469.3H42.7V128h426.7v341.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405.3 0h-42.7v64h42.7zm-256 0h-42.7v64h42.7zm277.4 0v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0zm42.6 469.3H42.7V128h426.7v341.3zm-64-245.3l-42.7-42.7l-149.3 149.4l-64-64l-42.7 42.7L213.3 416z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M149.3 0h-42.7v64h42.7zm256 0h-42.7v64h42.7zM234.7 384H192V192l-85.3 21.3v32l42.7-10.7V384h-42.7v21.3h128zm192-384v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0zm42.6 469.3H42.7V128h426.7v341.3zm-149.3-64c53.3 0 85.3-42.7 85.3-106.7S373.3 192 320 192c-53.4 0-85.3 42.7-85.3 106.7s32 106.6 85.3 106.6m0-181.3c21.3 0 32 32 32 74.7s-10.7 74.7-32 74.7s-32-32-32-74.7s10.7-74.7 32-74.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Call(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m426.7 453.8l-38.1-79.1c-8.2-16.9-18.8-29.2-37.1-21.7l-36.1 13.4c-28.9 13.4-43.3 0-57.8-20.2l-65-147.9c-8.2-16.9-3.9-32.8 14.4-40.3l50.5-20.2c18.3-7.6 15.4-23.4 7.2-40.3l-43.3-80.6c-8.2-16.9-25-21-43.3-13.5c-36.6 15.1-66.9 38.8-86.6 73.9c-24 42.9-12 102.6-7.2 127.7c4.8 25.1 21.6 69.1 43.3 114.2c21.7 45.2 40.7 80.7 57.8 100.8c17 20.1 57.8 75.1 108.3 87.4c41.4 10 86.1 1.6 122.7-13.5c18.4-7.2 18.4-23.1 10.3-40.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M331 374.8c-8.2-16.9-18.8-29.2-37.1-21.7l-36.1 13.4c-28.9 13.4-43.3 0-57.8-20.2l-65-147.8c-8.2-16.9-3.9-32.8 14.4-40.3l50.5-20.2c18.3-7.6 15.4-23.4 7.2-40.3L164 17.1c-8.2-16.9-25-21-43.3-13.5c-36.7 15.1-67 38.8-86.6 73.9c-24 42.9-12 102.6-7.2 127.7c4.8 25.1 21.6 69.1 43.3 114.2c21.7 45.2 40.7 80.7 57.8 100.8c17 20.1 57.8 75.1 108.3 87.4c41.4 10 86.1 1.6 122.7-13.5c18.3-7.5 18.4-23.4 10.2-40.4zM305.3 256c16.3 0 29.6-13.2 29.6-29.6c0-16.3-13.2-29.6-29.6-29.6c-16.3 0-29.6 13.2-29.6 29.6c0 16.3 13.3 29.6 29.6 29.6m78.8-59.2c-16.3 0-29.6 13.2-29.6 29.6c0 16.3 13.2 29.6 29.6 29.6c16.3 0 29.6-13.2 29.6-29.6c-.1-16.3-13.3-29.6-29.6-29.6m78.8 0c-16.3 0-29.6 13.2-29.6 29.6c0 16.3 13.2 29.6 29.6 29.6c16.3 0 29.6-13.2 29.6-29.6c-.1-16.3-13.3-29.6-29.6-29.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M310.5 374.8c-8.2-16.9-18.8-29.2-37.1-21.7l-36.1 13.4c-28.9 13.4-43.3 0-57.8-20.2l-65-147.8c-8.2-16.9-3.9-32.8 14.4-40.3l50.5-20.2c18.3-7.6 15.4-23.4 7.2-40.3l-43.3-80.6c-8.2-16.9-25-21-43.3-13.5c-36.5 15.1-66.9 38.8-86.5 73.9c-24 42.9-12 102.6-7.2 127.7c4.8 25.1 21.6 69.1 43.3 114.2c21.7 45.2 40.7 80.7 57.7 100.8c17 20.1 57.8 75.1 108.3 87.4c41.4 10 86.1 1.6 122.7-13.5c18.4-7.5 18.4-23.4 10.2-40.4zm200.7-178H333.9l78.8-78.8h-59.1l-98.5 98.5l98.5 98.5h59.1l-78.8-78.8h177.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328 374.8c-8.2-16.9-18.8-29.2-37.1-21.7l-36.1 13.4c-28.9 13.4-43.3 0-57.8-20.2l-65-147.8c-8.1-16.9-3.9-32.8 14.4-40.3l50.5-20.2c18.3-7.6 15.4-23.4 7.2-40.3L161 17.1c-8.2-16.9-25-21-43.3-13.5c-36.7 15.1-67 38.8-86.6 73.9c-24 42.9-12 102.6-7.2 127.7c4.8 25.1 21.6 69.1 43.3 114.2c21.7 45.2 40.8 80.7 57.8 100.8c17 20.1 57.8 75.1 108.3 87.4c41.4 10 86.1 1.6 122.7-13.5c18.3-7.5 18.4-23.4 10.2-40.4zm161.4-237.1L450 98.3l-59.1 59.1l-59.1-59.1l-39.4 39.4l59.1 59.1l-59.1 59.2l39.4 39.4l59.1-59.1l59.1 59.1l39.4-39.4l-59.1-59.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M311.3 374.7c-8.2-16.9-18.8-29.2-37.1-21.7l-36.1 13.4c-28.9 13.4-43.3 0-57.7-20.1l-64.9-147.8c-8.1-16.9-3.9-32.8 14.4-40.3l50.5-20.1c18.3-7.5 15.4-23.4 7.2-40.3l-43.3-80.6c-8.2-16.9-25-21-43.3-13.5c-36.6 15.2-66.9 38.8-86.5 73.9c-24 42.9-12 102.5-7.2 127.6c4.8 25.1 21.6 69 43.3 114.2c21.7 45.2 40.7 80.7 57.7 100.8c17 20.1 57.7 75.1 108.2 87.3c41.4 10 86 1.6 122.6-13.5c18.3-7.5 18.4-23.4 10.2-40.4zm102.2-256.6h-59.1l78.8 78.8H256v39.4h177.2L354.5 315h59.1l98.5-98.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M475.4 121.8H359.6l-12.2-36.6c0-20.2-16.4-36.6-36.6-36.6H201.1c-20.2 0-36.6 16.4-36.6 36.6l-12.2 36.6H36.6C16.4 121.8 0 138.2 0 158.4v219.4c0 20.2 16.4 36.6 36.6 36.6h438.9c20.2 0 36.6-16.4 36.6-36.6V158.4c-.1-20.2-16.5-36.6-36.7-36.6M256 377.8c-60.6 0-109.7-49.1-109.7-109.7S195.4 158.4 256 158.4s109.7 49.1 109.7 109.7c0 60.5-49.1 109.7-109.7 109.7m219.4-182.9c0 10.1-8.2 18.3-18.3 18.3h-18.3c-10.1 0-18.3-8.2-18.3-18.3v-18.3c0-10.1 8.2-18.3 18.3-18.3h18.3c10.1 0 18.3 8.2 18.3 18.3zm-219.4 0c-40.4 0-73.1 32.7-73.1 73.1s32.7 73.1 73.1 73.1s73.1-32.7 73.1-73.1s-32.7-73.1-73.1-73.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CercleEight(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M263 0C121.6 0 7 114.6 7 256s114.6 256 256 256s256-114.6 256-256S404.4 0 263 0m0 472.6c-119.6 0-216.6-97-216.6-216.6S143.4 39.4 263 39.4s216.6 97 216.6 216.6S382.7 472.6 263 472.6m0-393.8c-97.9 0-177.2 79.3-177.2 177.2S165.2 433.2 263 433.2c97.9 0 177.2-79.3 177.2-177.2S360.9 78.8 263 78.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CercleFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256c70.7 0 134.7-28.6 181-75c46.3-46.4 75-110.4 75-181C512 114.6 397.4 0 256 0m0 256c-96.7-96.7-160.9-96-200.8-71.5C84.6 101.9 163.3 42.7 256 42.7c117.8 0 213.3 95.5 213.3 213.3c0 25.1-4.6 49.1-12.5 71.5C416.9 352 352.7 352.7 256 256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CercleFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m21.3 446.7v-84.1h-42.7v84c-89-9.9-159.5-80.4-169.4-169.4h84.1v-42.7h-84c9.9-89 80.4-159.5 169.4-169.4v84.1h42.7V65.3c89 9.9 159.5 80.5 169.4 169.4h-84v42.7h84c-9.9 88.9-80.5 159.5-169.5 169.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CercleOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 426.7c-94.3 0-170.7-76.4-170.7-170.7S161.7 85.3 256 85.3S426.7 161.7 426.7 256S350.3 426.7 256 426.7m0-256c-47.1 0-85.3 38.2-85.3 85.3s38.2 85.3 85.3 85.3s85.3-38.2 85.3-85.3s-38.2-85.3-85.3-85.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CercleSeven(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 64c35.4 0 64 28.6 64 64s-28.6 64-64 64s-64-28.6-64-64s28.6-64 64-64m85.3 384H170.7v-21.3H192V256h-21.3v-21.3H320v192h21.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CercleSix(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 42.7c35.4 0 64 28.6 64 64s-28.6 64-64 64s-64-28.6-64-64s28.6-64 64-64M53.3 213.3c0-35.4 28.6-64 64-64s64 28.6 64 64s-28.6 64-64 64s-64-28.6-64-64m117.4 224c-35.4 0-64-28.6-64-64s28.6-64 64-64s64 28.6 64 64s-28.7 64-64 64M213.3 256c0-23.5 19.1-42.7 42.7-42.7s42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7s-42.7-19.2-42.7-42.7m128 181.3c-35.4 0-64-28.6-64-64s28.6-64 64-64s64 28.6 64 64s-28.6 64-64 64m53.4-160c-35.4 0-64-28.6-64-64s28.6-64 64-64s64 28.6 64 64s-28.7 64-64 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CercleThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.5 1.5C115.4 1.5 1 115.9 1 257s114.4 255.5 255.5 255.5S512 398.1 512 257S397.6 1.5 256.5 1.5m149 259C394.2 328.7 381.3 406 320.7 406c-60.6 0-73.5-77.3-84.9-145.5c-8.2-49-18.3-110-42.9-110c-24.5 0-34.7 61-42.9 110c-9.1 54.4-19.2 114.4-53.6 136.6c-9.6-10.9-18.1-22.8-25.3-35.5c20.2-9.3 29.4-63.4 36.9-108.1c11.4-68.2 24.3-145.5 84.9-145.5c60.6 0 73.5 77.3 84.9 145.5c8.2 49 18.3 110 42.9 110c24.6 0 34.7-61 42.9-110c9-54.2 19.1-114 53.2-136.4c9.5 10.9 18 22.7 25.2 35.5c-19.9 9.8-29.1 63.5-36.5 107.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CercleTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M401.1 130.5c-22.1 28.2-30.3 77-37.8 122c-8.2 49.1-18.4 110.2-43 110.2s-34.8-61.1-43-110.2c-11.4-68.4-24.3-145.8-85-145.8s-73.6 77.5-85 145.8c-5.3 32-11.6 68.9-22.1 90.8c7.1 13.9 15.9 26.9 26.1 38.6c22.3-28.1 30.6-77.2 38.1-122.4c8.2-49.1 18.4-110.2 43-110.2s34.8 61.1 43 110.2c11.4 68.3 24.3 145.8 85 145.8s73.6-77.5 85-145.8c5.3-31.8 11.5-68.2 21.8-90.3c-7.2-14-16-27-26.1-38.7M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 469.3c-117.8 0-213.3-95.5-213.3-213.3c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 117.8-95.5 213.3-213.3 213.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="256" cy="256" r="256" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCornerArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 42.7L469.3 0l-128 128V0l-64 64l-.1 170.8l170.8-.1l64-64H384zM0 341.3h128L0 469.3L42.7 512l128-128v128l64-64l.1-170.8l-170.8.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCornerArrowTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.7 0v128L42.7 0L0 42.7l128 128H0l64 64l170.8.1l-.1-170.8zM512 341.3l-64-64l-170.8-.1l.1 170.8l64 64V384l128 128l42.7-42.7l-128-128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m96 373.3l-96 96L42.7 512l96-96l74.7 74.7v-192h-192zm394.7-74.6h-192v192l74.7-74.7l96 96l42.7-42.7l-96-96zM42.7 0L0 42.7l96 96l-74.7 74.7h192v-192L138.7 96zM416 138.7l96-96L469.3 0l-96 96l-74.7-74.7v192h192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloth(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 512l149.3-42.7V0L0 42.7zM362.7 42.7V512L512 469.3V0zm-192 426.6L341.3 512V42.7L170.7 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClothOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M362.7 42.7v74.9c3.2.9 6.3 1.8 9.5 2.7c41 11.7 83.5 23.8 104.3 42.4c24.2 21.7 8.1 50.4-6.2 75.7c-11 19.4-22.3 39.5-22.3 60.2h-21.3c0-26.3 13.3-49.9 25-70.7c15.6-27.7 21-40 10.5-49.3c-17.1-15.4-59-27.3-95.9-37.8c-1.3-.4-2.4-.7-3.7-1.1V512L512 469.3V0zm128 362.6l-21.3 21.3l-32-32l-32 32l-21.4-21.3l32-32l-32-32l21.3-21.3l32 32l32-32l21.3 21.3l-32 32zm-320-58.7c15.8-8.6 34.5-12.1 56.7-4.7c14.6 4.9 19-.8 20.8-3.2c8.1-10.7 10.1-18.7 7.8-18.7c7.2 0 14.9 2.2 21.3 0c0 10.7-4 20.9-12.2 31.6c-5.7 7.6-19.2 18.9-44.5 10.5c-19.6-6.5-35.4-1.2-50 10.3v97L341.3 512V133.5c-18.2-5.4-34.8-11-47.4-17.3c-21-10.5-51-6.7-64.1 8.1c-11.3 12.7-4.6 29 3.1 40.4c3.5 5.3 9.1 20.8 12.5 27.2c-7.2 1.4-15.2-3.6-21.3 0c-2.6-4.9-6-11.1-8.9-15.4c-16.4-24.6-16.9-48.8-1.3-66.4c19.6-22.1 59.8-28 89.6-13.1c9.9 4.9 23.2 9.5 37.9 14.1V42.7L170.6 0zM256 213.3c23.5 0 42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7c-23.5 0-42.7-19.1-42.7-42.7c0-23.5 19.2-42.7 42.7-42.7M0 512l149.3-42.7v-74.7c-4.1 5-8.3 9.3-12.5 14.8c-9.5 12.6-17 22.5-25.4 26.7c-4.9 2.5-10 3.6-15.1 3.6c-23.7 0-47.4-23.7-61.2-37.6c-17.6-17.6-3.4-39.4 9.2-58.5c9.2-14.1 19.7-30.1 19.7-45v-85.3h21.3v85.3c0 21.3-12.3 40.1-23.2 56.7C50 374 45.6 382.5 50.2 387.1c17.3 17.3 38.8 36.5 51.7 30c3.9-2 11.4-11.8 18-20.5c8.1-10.7 17.8-23.1 29.5-34V0L0 42.7zm74.7-384c17.7 0 32 14.3 32 32s-14.3 32-32 32s-32-14.3-32-32s14.3-32 32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M470.7 306.2c3-11.2 4.7-22.9 4.7-35c0-75.8-61.4-137.1-137.1-137.1c-19.5 0-38 4.1-54.7 11.4c-16.8-39-55.6-66.3-100.7-66.3c-60.6 0-109.7 49.1-109.7 109.7c0 4.1.8 7.9 1.2 11.9C30.5 221.1 0 265.3 0 316.9c0 70.7 57.3 128 128 128h310.9c40.4 0 73.1-32.7 73.1-73.1c0-29-16.9-53.7-41.3-65.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M470.7 280.2c3-11.2 4.7-22.9 4.7-35c0-75.8-61.4-137.1-137.1-137.1c-19.5 0-38 4.1-54.7 11.4c-16.8-39-55.6-66.3-100.7-66.3c-60.6 0-109.7 49.1-109.7 109.7c0 4.1.8 7.9 1.2 11.9C30.5 195.1 0 239.3 0 290.9c0 70.7 57.3 128 128 128h310.9c40.4 0 73.1-32.7 73.1-73.1c0-29-16.9-53.7-41.3-65.6M256 364l-91.4-91.4h54.9v-91.4h73.1v91.4h54.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudReload(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M470.7 280.2c2.9-11.2 4.7-22.9 4.7-35c0-75.8-61.4-137.1-137.1-137.1c-19.5 0-37.9 4.1-54.7 11.4c-16.8-39-55.6-66.3-100.7-66.3c-60.6 0-109.7 49.1-109.7 109.7c0 4.1.8 7.9 1.2 11.9C30.5 195.1 0 239.3 0 290.9c0 70.7 57.3 128 128 128h310.9c40.4 0 73.1-32.7 73.1-73.1c0-29-16.9-53.7-41.3-65.6m-105-25.9h-91.4l33.4-33.4c-13.2-13.2-31.5-21.4-51.7-21.4c-40.4 0-73.1 32.7-73.1 73.1s32.7 73.1 73.1 73.1c24.9 0 46.8-12.5 60-31.4l25.9 25.9c-20.1 25.5-50.9 42.2-85.8 42.2c-60.6 0-109.7-49.1-109.7-109.7S195.5 163 256.1 163c30.3 0 57.7 12.3 77.6 32.1l32.1-32.1v91.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M470.7 277.2c3-11.2 4.7-22.9 4.7-35c0-75.8-61.4-137.1-137.1-137.1c-19.5 0-38 4.1-54.7 11.4c-16.8-39-55.6-66.3-100.7-66.3c-60.6 0-109.7 49.1-109.7 109.7c0 4.1.8 7.9 1.2 11.9C30.5 192.1 0 236.3 0 287.9c0 70.7 57.3 128 128 128h310.9c40.4 0 73.1-32.7 73.1-73.1c0-29-16.9-53.7-41.3-65.6m-178.1-25.9v91.4h-73.1v-91.4h-54.9l91.4-91.4l91.4 91.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coin(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 243.2c141.4 0 256-47.8 256-106.7c0-58.9-114.6-106.7-256-106.7S0 77.6 0 136.5c0 58.9 114.6 106.7 256 106.7m0 170.6c-97.7 0-184.1-23.5-238.6-59.8C6.3 366 0 378.9 0 392.5c0 58.9 114.6 106.7 256 106.7s256-47.8 256-106.7c0-13.6-6.3-26.5-17.4-38.4c-54.5 36.2-140.9 59.7-238.6 59.7m0-128c-97.7 0-184.1-23.5-238.6-59.8C6.3 238 0 250.9 0 264.5c0 58.9 114.6 106.7 256 106.7s256-47.8 256-106.7c0-13.6-6.3-26.5-17.4-38.4c-54.5 36.2-140.9 59.7-238.6 59.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232.5 343.9c-10.4-18-19.2-36.5-27.2-55.2c-35.2 60.1-57.8 126.7-66.6 195.5c67.3 34.6 148.5 38.6 221.2 6.3c-50-39.3-93.7-88.2-127.4-146.6m26.2-27.5c34.5 60.5 80.8 113.4 136 155.4c63.6-41 107.7-109.4 116-188.4c-59 23.6-123.2 37-190.6 37c-20.8 0-41.2-1.6-61.4-4m-56.1-87.9c-69.7-.4-138.7 13.3-202.6 40.1c3.7 75.6 40.8 147.9 105.2 194.7c9-62.9 29.5-125.2 63.3-183.6c10.3-18 21.9-34.9 34.1-51.2m50.7-32.4c-34.5-60.5-80.8-113.5-136-155.4c-63.6 41-107.7 109.4-116 188.4c59-23.6 123.2-37 190.6-37c20.8 0 41.2 1.6 61.4 4m26.2-27.4c10.4 18 19.2 36.5 27.2 55.2c35.2-60.1 57.8-126.7 66.6-195.5C306-6.2 224.8-10.2 152.1 22.1c50 39.2 93.7 88.2 127.4 146.6M309.4 284c69.7.4 138.7-13.3 202.6-40.1c-3.7-75.6-40.8-147.9-105.1-194.7c-9 62.9-29.6 125.2-63.3 183.6c-10.4 18-22 34.9-34.2 51.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 472.6c-119.6 0-216.6-97-216.6-216.6S136.4 39.4 256 39.4s216.6 97 216.6 216.6s-97 216.6-216.6 216.6m-137.8-78.8l187.1-88.6l88.6-187.1l-187.1 88.6zm167.3-108.3l-118.2 59.1l59.1-118.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0M42.7 256c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 144-426.6 144-426.6 0m256 85.3L352 96L213.3 341.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M236.3 452.9h39.4v-39.4h-39.4zm39.4-393.8h-39.4v39.4h39.4zM59.1 275.7h39.4v-39.4H59.1zM256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 472.6c-119.6 0-216.6-97-216.6-216.6S136.4 39.4 256 39.4s216.6 97 216.6 216.6s-97 216.6-216.6 216.6M137.8 256c0 65.2 52.9 118.2 118.2 118.2c65.2 0 118.2-52.9 118.2-118.2V137.8H256c-65.2 0-118.2 53-118.2 118.2M256 196.9c-32.6 0-59.1 26.4-59.1 59.1h-19.7c0-43.5 35.3-78.8 78.8-78.8zm157.5 39.4v39.4h39.4v-39.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compose(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 28.3H0v469.3h85.3V263H64zm149.3 0h-64V263H128v234.7h106.7V263h-21.3V28.3zm234.7 0V263h-21.3v234.7H512V28.3zm-85.3 0h-64V263h-21.3v234.7H384V263h-21.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cover(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469.3 106.7H42.7v-64h298.7v42.7H384V42.7C384 19.1 364.9 0 341.3 0H42.7C19.1 0 0 19.1 0 42.7v426.7C0 492.9 19.1 512 42.7 512h426.7c23.6 0 42.7-19.1 42.7-42.7V384H405.3c-23.6 0-42.7-19.1-42.7-42.7v-64c0-23.5 19.1-42.7 42.7-42.7H512v-85.3c0-23.5-19.1-42.6-42.7-42.6m-64 234.6H512v-64H405.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 21.3V0h-21.3l-85.3 85.3H128V0H85.3v85.3H0V128h85.3v298.7H384V512h42.7v-85.3H512V384h-85.3V106.7zM128 128h234.7L128 362.7zm256 256H149.3L384 149.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crpss(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 76.8L435.2 0L256 179.2L76.8 0L0 76.8L179.2 256L0 435.2L76.8 512L256 332.8L435.2 512l76.8-76.8L332.8 256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DailPad(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M202.5 93.1h93.1V0h-93.1zM342.1 0v93.1h93.1V0zM62.8 93.1h93.1V0H62.8zm139.7 139.6h93.1v-93.1h-93.1zm139.6 0h93.1v-93.1h-93.1zm-279.3 0h93.1v-93.1H62.8zm139.7 139.7h93.1v-93.1h-93.1zm139.6 0h93.1v-93.1h-93.1zm-279.3 0h93.1v-93.1H62.8zM202.5 512h93.1v-93.1h-93.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M42.7 469.3c0 23.5 19.1 42.7 42.7 42.7h341.3c23.5 0 42.7-19.1 42.7-42.7V192H42.7zm320-213.3h42.7v192h-42.7zm-128 0h42.7v192h-42.7zm-128 0h42.7v192h-42.7zm384-170.7h-128V42.7C362.7 19.1 343.5 0 320 0H192c-23.5 0-42.7 19.1-42.7 42.7v42.7h-128C9.5 85.3 0 94.9 0 106.7V128c0 11.8 9.5 21.3 21.3 21.3h469.3c11.8 0 21.3-9.5 21.3-21.3v-21.3c.1-11.8-9.4-21.4-21.2-21.4m-170.7 0H192V42.7h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Divide(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 85.3c23.5 0 42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7c-23.5 0-42.7-19.1-42.7-42.7s19.2-42.7 42.7-42.7m0 341.4c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7s-19.2 42.7-42.7 42.7m149.3-149.4c0 11.8-9.5 21.3-21.3 21.3H128c-11.8 0-21.3-9.6-21.3-21.3v-42.7c0-11.8 9.5-21.3 21.3-21.3h256c11.8 0 21.3 9.6 21.3 21.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DivideOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 228.3v85.3h512v-85.3zM298.7 57.7h-85.3V143h85.3zm-85.4 426.6h85.3V399h-85.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M372.4 512L512 372.4H372.4zM0 0v512h349.1V349.1H512V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M372.4 512L512 372.4H372.4zM0 0v512h349.1V349.1H512V0zm395.6 279.3H279.3v116.4h-46.5V279.3H116.4v-46.5h116.4V116.4h46.5v116.4h116.4v46.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M372.4 512L512 372.4H372.4zM93.1 93.1V512h256V349.1H512v-256zM418.9 0H0v418.9h46.5V46.5h372.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M372.4 517.5L512 377.9H372.4zM0 5.5v512h349.1V354.6H512V5.5zm395.6 279.3H116.4v-46.5h279.3v46.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Down(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 512h139.6V372.4H0zm46.5-93.1H93v46.5H46.5zM0 139.6h139.6V0H0zm0 186.2h139.6V186.2H0zm46.5-93.1H93v46.5H46.5zM209.5 0c162.9 93.1 244.4 221.1 58.2 337.5L186.2 256v256h256l-81.5-81.5C605.1 186.2 372.4 0 209.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownArrow(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M289.7 341.3V0h-85.4v341.3L33.7 170.7v128L247 512l213.3-213.3v-128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M283.7 298.7V0h-85.4v298.7h-128L241 512l170.7-213.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M442.2 93.1L256 279.3L69.8 93.1L0 162.9l256 256l256-256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M302.5 93.1h-93.1v186.2H69.8L256 512l186.2-232.7H302.5zM0 0v46.5h512V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M315.1 4.1v229.2h98.5L256 410.5L98.5 233.3H197V4.1C84.1 30.8 0 132 0 253c0 141.4 114.6 256 256 256s256-114.6 256-256c0-121-84.1-222.2-196.9-248.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M442.2 186.2H302.5V0h-93.1v186.2H69.8L256 418.9zM0 465.5V512h512v-46.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M442.2 186.2H302.5V0h-93.1v186.2H69.8L256 418.9zm23.3 186.2v93.1h-419v-93.1H0V512h512V372.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DubleCornerArrowBlodTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m511.9 232.7l-93-93V349l-93-93l-69.9 69.8l93 93H139.7l93 93l279.3.2zM256 186.2l-93-93h209.4l-93-93L0 0l.1 279.3l93 93V163l93 93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DubleCornerArrowFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M209.5 23.3L314.2 128L256 186.2l69.8 69.8l58.2-58.2l104.7 104.7L512 0zM256 325.8L186.2 256L128 314.2L23.3 209.5L0 512l302.5-23.3L197.8 384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DubleCornerArrowFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.8.1L256 69.9h139.6L256 209.5l46.5 46.5l139.6-139.6V256l69.8-69.8L512 0zM209.5 256L69.9 395.6V256L.1 325.8L0 512l186.2-.1l69.8-69.8H116.4L256 302.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DubleCornerArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m232.7.1l-93 93H349l-93 93l69.8 69.8l93-93v209.4l93-93L512 0zM256 325.8L186.2 256l-93 93V139.7l-93 93L0 512l279.3-.1l93-93H163z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DubleCornerArrowSix(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M197.8 128L302.5 23.3L0 0l23.3 302.5L128 197.8l58.2 58.2l69.8-69.8zM512 512l-23.3-302.5L384 314.2L325.8 256L256 325.8l58.2 58.2l-104.7 104.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DubleCornerArrowThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M116.4 69.9H256L186.2.1L0 0l.1 186.2L69.9 256V116.4L209.5 256l46.5-46.5zm395.5 255.9L442.1 256v139.6L302.5 256L256 302.5l139.6 139.6H256l69.8 69.8l186.2.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equal(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m149.3 341.3c0 11.8-9.5 21.3-21.3 21.3H128c-11.8 0-21.3-9.6-21.3-21.3v-42.7c0-11.8 9.5-21.3 21.3-21.3h256c11.8 0 21.3 9.6 21.3 21.3zm0-128c0 11.8-9.5 21.3-21.3 21.3H128c-11.8 0-21.3-9.6-21.3-21.3v-42.7c0-11.8 9.5-21.3 21.3-21.3h256c11.8 0 21.3 9.6 21.3 21.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EqualOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 144v85.3h512V144zm0 256h512v-85.3H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equalizer(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M106.7 25.8H21.3v64h85.3v-64zM21.3 495.2h85.3V217.8H21.3zM490.7 25.8h-85.3v277.3h85.3zm-85.4 469.4h85.3v-64h-85.3zm-192 0h85.3V324.5h-85.3zm85.4-469.4h-85.3v170.7h85.3zM0 175.2h128v-42.7H0zm192 106.6h128v-42.7H192zm192 64v42.7h128v-42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EqualizerOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85.3 21.3H42.7V64h42.7V21.3zM42.7 490.7h42.7v-256H42.7zm384 0h42.7V448h-42.7zm42.6-469.4h-42.7v256h42.7zm-192 0h-42.7v149.3h42.7zm-42.6 469.4h42.7V341.3h-42.7zM0 192h128v-85.3H0zm192 106.7h128v-85.3H192zM384 320v85.3h128V320z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EqualizerTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M42.7 486.7h42.7v-256H42.7zM469.3 17.3h-42.7v256h42.7zm-384 0H42.7V60h42.7V17.3zm192 0h-42.7v149.3h42.7zM0 188h128v-85.3H0zm21.3-64h85.3v42.7H21.3zm213.4 362.7h42.7V337.3h-42.7zm192 0h42.7V444h-42.7zM384 316v85.3h128V316zm106.7 64h-85.3v-42.7h85.3zM192 294.7h128v-85.3H192zm21.3-64h85.3v42.7h-85.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorIcon(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0M64 256c0-106.1 86-192 192-192c42.1 0 81 13.7 112.6 36.7L100.7 368.6C77.7 337 64 298.1 64 256m192 192c-42.1 0-81-13.7-112.6-36.7l267.9-267.9c23 31.7 36.7 70.5 36.7 112.6c0 106.1-86 192-192 192"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M307.5 64.4c31.9 0 69.2 14 95.8 31.9V21.8C376.3 8.9 339.5.5 307.5.5C254.3.5 216.6 14.7 181 45.3c-35.6 30.6-63.1 71.4-75.8 125.5H62.7v63.9h42.6v42.6H62.7v63.9h42.6c21.3 117.1 95.8 170.3 202.3 170.3c42.6 0 62.8-4.2 95.8-21.3v-74.5c-32.6 20-65.9 31.9-95.8 31.9c-64.2 0-103-34-119.5-106.3l172.7-.1v-63.9H190.4v-42.6h170.3v-63.9H190.4c8.9-33.8 23.8-59.3 46.1-79.4c22.4-20.1 39.1-27 71-27"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exit(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469.3 469.3H42.7V42.7H256L298.7 0h-256C19.1 0 0 19.1 0 42.7v426.7C0 492.9 19.1 512 42.7 512h426.7c23.6 0 42.7-19.1 42.7-42.7V320l-42.7 42.7v106.6zm-384-42.6C149.1 255.7 234.7 256 362.7 256v85.3L512 192L362.7 42.7V128C85.3 128 85.1 341.1 85.3 426.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 226v32c128 192 384 192 512 0v-32C384 34 128 34 0 226m256 144c-70.7 0-128-57.3-128-128s57.3-128 128-128s128 57.3 128 128s-57.3 128-128 128m0-200c0-8.3 1.7-16.1 4.3-23.6c-1.5-.1-2.8-.4-4.3-.4c-53 0-96 43-96 96s43 96 96 96s96-43 96-96c0-1.5-.4-2.8-.4-4.3c-7.4 2.6-15.3 4.3-23.6 4.3c-39.8 0-72-32.2-72-72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func F(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M69.8 0C31.3 0 0 31.2 0 69.8c0 38.6 31.3 69.8 69.8 69.8c38.6 0 69.8-31.3 69.8-69.8C139.6 31.2 108.4 0 69.8 0m0 104.7c-19.3 0-34.9-15.6-34.9-34.9s15.6-34.9 34.9-34.9c19.3 0 34.9 15.6 34.9 34.9s-15.6 34.9-34.9 34.9M512 93.1V0H209.5v512h116.4V302.5h162.9v-93.1h-163V93.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feed(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M58.2 395.6C26 395.6 0 421.7 0 453.8S26 512 58.2 512c32.1 0 58.2-26 58.2-58.2s-26.1-58.2-58.2-58.2M0 0v93.1c231.4 0 418.9 187.5 418.9 418.9H512C512 229.2 282.8 0 0 0m0 186.2v93.1c128.5 0 232.7 104.2 232.7 232.7h93.1c0-180-145.9-325.8-325.8-325.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 0v128h128zm-21.3 0H64v512h384V149.3H298.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileEight(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M300.7 0H66v512h384V149.3H300.7zm10.6 245.3L354 288l-53.3 53.3l53.3 53.3l-42.7 42.7L258 384l-53.3 53.3l-42.7-42.6l53.3-53.3L162 288l42.7-42.7l53.3 53.3zM322 0v128h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileEleven(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M62.7 85.3H20V512h384v-42.7H62.7zM361.3 0v128h128zM340 0H105.3v426.7h384V149.3H340zm64 298.7h-64V384h-85.3v-85.3h-64L297.3 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M334.5 0v128h128zm-21.3 0H78.5v512h384V149.3H313.2zm-21.4 448h-42.7v-42.7h42.7zm0-64h-42.7l-21.3-149.3h85.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M305.7 0H71v512h384V149.3H305.7zm64 320v42.7H156.3V320zM327 0v128h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileNine(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 0v128h128zm-64 405.3c35.4 0 64-28.6 64-64c0-9.4-2.1-18.2-5.7-26.2l-84.5 84.5c8 3.6 16.8 5.7 26.2 5.7M298.7 0H64v512h384V149.3H298.7zm64 341.3c0 58.9-47.8 106.7-106.7 106.7c-58.9 0-106.7-47.7-106.7-106.7c0-58.9 47.8-106.7 106.7-106.7c58.9.1 106.7 47.8 106.7 106.7m-170.7 0c0 9.4 2.1 18.2 5.7 26.2l84.5-84.5c-8-3.6-16.9-5.8-26.2-5.8c-35.4.1-64 28.8-64 64.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M330 0v128h128zm-21.3 0H74v512h384V149.3H308.7zm64 320h-64v128h-85.3V320h-64L266 213.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSeven(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.5 0v128h128zm-21.3 0H69.5v512h384V149.3H304.2zm64 277.3c23.5 0 42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7m-213.4 85.4c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7m106.7 0c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7s-19.2 42.7-42.7 42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSix(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M322 0v128h128zm-21.3 0H66v288l64-64l85.3 85.3l85.3-85.3l85.3 85.3l64-64v-96H300.7zm0 309.3l-85.3 85.3l-85.4-85.3l-64 64V512h384V330.7l-64 64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTen(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M338 0H103.3v426.7h384V149.3H338zm21.3 0v128h128zM60.7 85.3H18V512h384v-42.7H60.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileThirteen(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 0v128h128zm-21.3 0H64v512h384V149.3H298.7zm-192 42.7H256v106.7H106.7zM256 405.3H106.7v-42.7H256zM405.3 320H106.7v-42.7h298.7V320zm0-128v42.7H106.7V192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M302.2 0H67.5v512h384V149.3H302.2zm64 320v42.7h-85.3V448h-42.7v-85.3h-85.3V320h85.3v-85.3h42.7V320zM323.5 0v128h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTwelve(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M362.3 0v128h128zM63.7 85.3H21V512h384v-42.7H63.7zM341 0H106.3v426.7h384V149.3H341zm64 277.3L298.3 384L191.7 277.3h64V192H341v85.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.5 0v128h128zm-21.3 0H69.5v512h384V149.3H304.2zm64 341.3L261.5 448L154.8 341.3h64v-128h85.3v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M490.7 96H277.3l-10.7-42.7c-1.5-9.2-12-21.3-21.3-21.3H96c-9.3 0-19.8 12.2-21.3 21.3L64 96H21.3C12.1 96 0 108.1 0 117.3v341.3C0 468 12.1 480 21.3 480h469.3c9.3 0 21.3-12 21.3-21.3V117.3C512 108.1 500 96 490.7 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 192l36.6 237.7c0 8 10.3 18.3 18.3 18.3h402.3c8 0 18.3-10.3 18.3-18.3L512 192zm475.4-54.9c0-8-10.3-18.3-18.3-18.3H274.3l-9.1-36.6c-1.3-7.8-10.3-18.3-18.3-18.3h-128c-7.9 0-17 10.4-18.3 18.3l-9.1 36.6H54.9c-7.9 0-18.3 10.3-18.3 18.3v18.3h438.9v-18.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M475.4 137.1c0-8-10.3-18.3-18.3-18.3H274.3l-9.1-36.6c-1.3-7.8-10.3-18.3-18.3-18.3h-128c-7.9 0-17 10.4-18.3 18.3l-9.1 36.6H54.9c-7.9 0-18.3 10.3-18.3 18.3v18.3h438.9v-18.3zM0 192l36.6 237.7c0 8 10.3 18.3 18.3 18.3h402.3c8 0 18.3-10.3 18.3-18.3L512 192zm292.6 128v73.1h-73.1V320h-54.9l91.4-91.4l91.4 91.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M475.4 137.1c0-8-10.3-18.3-18.3-18.3H274.3l-9.1-36.6c-1.3-7.8-10.3-18.3-18.3-18.3h-128c-7.9 0-17 10.4-18.3 18.3l-9.1 36.6H54.9c-7.9 0-18.3 10.3-18.3 18.3v18.3h438.9v-18.3zM0 192l36.6 237.7c0 8 10.3 18.3 18.3 18.3h402.3c8 0 18.3-10.3 18.3-18.3L512 192zm256 219.4L164.6 320h54.9v-73.1h73.1V320h54.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FotScreen(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M418.9 418.9V279.3l-46.5 46.5l-46.5-46.5l-46.5 46.5l46.5 46.5l-46.5 46.5h139.5zM0 0v512h512V0zm465.5 465.5h-419v-419h418.9v419zM186.2 232.7l46.5-46.5l-46.5-46.5l46.5-46.5H93.1v139.6l46.5-46.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourBox(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 512h232.7V279.3H0zm0-279.3h232.7V0H0zM279.3 512H512V279.3H279.3zm0-512v232.7H512V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Froward(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 382.5l256-128l-256-128zm512-128l-256-128v256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrowardOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M495.8 232.9L252.7 90c-21.6-12.7-39.2.1-39.2 28.5v74L39.2 90C17.6 77.3 0 90.1 0 118.5v275c0 28.3 17.6 41.2 39.2 28.5l174.3-102.5v74c0 28.3 17.6 41.2 39.2 28.5l243.1-143c21.6-12.6 21.6-33.4 0-46.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fullscreen(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m93.1 139.6l46.5-46.5l-46.5-46.6L139.6 0H0v139.6l46.5-46.5zm0 232.8l-46.5 46.5L0 372.4V512h139.6l-46.5-46.5l46.5-46.5zm279.3-232.8H139.6v232.7h232.7V139.6zm-46.6 186.2H186.2V186.2h139.6zM372.4 0l46.5 46.5L372.4 93l46.5 46.5L465.4 93l46.5 46.5V0zm46.5 372.4l-46.5 46.5l46.5 46.5l-46.5 46.6H512V372.4l-46.5 46.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glass(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M511.6 197.5c-34.9-7.1-69.4-20.3-105-20c-20.4.1-40.9 3.7-61.1 7c-22.7 3.7-44.6 12.2-67.3 15.3c-7.3 1-14.7 1.4-22.1 1.4c-7.4 0-14.8-.4-22.1-1.4c-22.7-3.2-44.6-11.6-67.3-15.3c-20.3-3.4-40.8-6.9-61.2-7c-35.6-.2-70.1 12.9-105 20c0 12-2.1 27.5 3.1 36c5.4 8.8 14.5 16.7 16.8 28.6c5.6 28.1 13.5 55.6 23.7 82c4.9 12.8 14.2 21.8 24.6 27.9c25 14.8 54.6 17 82.1 12.4c39.1-6.6 60.3-41.2 76.5-80.8c6.1-16.1 6.7-34.6 14.7-49.3c3-5.5 8.4-6.1 14-5.8c5.6-.3 11.1.3 14 5.8c8 14.7 8.6 33.3 14.7 49.3c16.2 39.6 37.4 74.1 76.5 80.8c27.4 4.7 57.1 2.5 82.1-12.4c10.4-6.1 19.7-15.1 24.6-27.9c10.1-26.4 18.1-53.9 23.7-82c2.3-11.9 11.5-19.8 16.8-28.6c5.3-8.5 3.2-23.9 3.2-36M204.3 310.2c-19.4 47.5-54.3 57.4-80.1 57.4c-10 0-20.5-1.5-31.2-4.5c-19.1-5.3-33.3-21.3-40-44.8c-9-31.4-12.7-56.7-11.7-79.4c1.4-29.4 27.9-35.4 57.3-39.3c13.7-1.8 24.7-2.7 34.6-2.8h.1c13.3 0 23.8 1.2 33.6 3.6c28.5 7.1 44.7 16.9 50.9 30.9c6.9 15.7 3 38.6-13.5 78.9m254.8 8.1c-6.7 23.5-21 39.4-40 44.8c-10.7 3-21.2 4.5-31.3 4.5c-25.8 0-60.7-10-80.1-57.4c-16.5-40.3-20.4-63.1-13.4-78.8c6.2-14 22.4-23.8 50.9-30.9c9.7-2.4 20.3-3.6 33.6-3.6h.1c9.9.1 20.9 1 34.6 2.8c29.4 3.9 55.9 9.8 57.3 39.3c1 22.6-2.7 47.8-11.7 79.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0L0 256l42.7 42.7l64-64V512h298.6V234.7l64 64L512 256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0L0 256l42.7 42.7l64-64V512h298.7V234.7l64 64L512 256zm-21.3 448h-64v-64h64zm0-106.7h-64v-64h64zM341.3 448h-64v-64h64zm0-106.7h-64v-64h64zM426.7 0h-64v42.7l64 64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M257 0L1 256l42.7 42.7L257 85.3l213.3 213.3L513 256zM107.7 298.7V512h298.7V298.7L257 149.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0L0 256l42.7 42.7l64-64V512h298.7V234.7l64 64L512 256zm106.7 469.3H149.3V192L256 85.3L362.7 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HulfOfCircleTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M159 0v512c141.4 0 256-114.6 256-256S300.4 0 159 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hurt(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 28.3c-64 0-96.2 27.6-128 64c-31.8-36.4-64-64-128-64S0 71 0 199c0 64 64 192 256 298.7C448 391 512 263 512 199c0-128-64-170.7-128-170.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HurtOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 29.8c-64 0-96.2 27.6-128 64c-31.8-36.4-64-64-128-64S0 72.5 0 200.5c0 64 64 192 256 298.7c192-106.7 256-234.7 256-298.7c0-128-64-170.7-128-170.7M256 450C81.7 346.6 42.7 235.7 42.7 200.5c0-58.4 14.8-128 85.3-128c44.8 0 66.6 15.9 95.9 49.4l32.1 35.9l32.1-35.9c29.3-33.5 51.1-49.4 95.9-49.4c70.5 0 85.3 69.6 85.3 128c0 35.2-39 146.1-213.3 249.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HurtThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m384 34.3l-128 128h106.7L149.3 375.7L234.7 205H128l89.6-143.4C194.9 45 167.5 34.3 128 34.3C64 34.3 0 77 0 205c0 64 64 192 256 298.7C448 397 512 269 512 205c0-128-64-170.7-128-170.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCard(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 394c0 23.5 19.1 42.7 42.7 42.7h426.7c23.5 0 42.7-19.1 42.7-42.7v-42.7H0zM469.3 95.3H42.7C19.1 95.3 0 114.5 0 138v170.7h512V138c0-23.5-19.1-42.7-42.7-42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCardOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426.7 64H384v85.3h42.7zM128 64H85.3v85.3H128zm341.3 42.7H448v64h-85.3v-64H149.3v64H64v-64H42.7C19.1 106.7 0 125.8 0 149.3v256C0 428.9 19.1 448 42.7 448h426.7c23.5 0 42.7-19.1 42.7-42.7v-256c-.1-23.5-19.2-42.6-42.8-42.6M192 384H42.7V234.7H192zm149.3-21.3H234.7V320h106.7v42.7zm128-64H234.7V256h234.7v42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m349.1 140.6l-69.8 139.6l-46.5-46.5l-46.5 46.5l-46.5-46.5l-46.7 139.7h325.8zM0 1v512h512V1zm465.5 418.9h-419V47.5h418.9v372.4zM139.6 187.2c25.7 0 46.5-20.9 46.5-46.5c0-25.7-20.9-46.5-46.5-46.5S93.1 115 93.1 140.6c0 25.7 20.8 46.6 46.5 46.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoinCornerArrowFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m512 512l-23.3-302.5L384 314.2L197.8 128L302.5 23.3L0 0l23.3 302.5L128 197.8L314.2 384L209.5 488.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoinCornerArrowFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M511.9 325.8L442.1 256v139.6L116.4 69.9H256L186.2.1L0 0l.1 186.2L69.9 256V116.4l325.7 325.7H256l69.8 69.8l186.2.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoinCornerArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m232.7.1l-93 93.1H349L93.2 349V139.7l-93.1 93L0 512l279.3-.1l93-93.1H163L418.8 163v209.3l93.1-93L512 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoinCornerArrowSix(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M209.5 23.3L314.2 128L128 314.2L23.3 209.5L0 512l302.5-23.3L197.8 384L384 197.8l104.7 104.7L512 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoinCornerArrowThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.8.1L256 69.9h139.6L69.9 395.6V256L.1 325.8L0 512l186.2-.1l69.8-69.8H116.4l325.7-325.7V256l69.8-69.8L512 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoinCornerArrowTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m511.9 232.7l-93.1-93V349L163 93.2h209.3L279.3.1L0 0l.1 279.3l93.1 93V163L349 418.8H139.7l93 93.1l279.3.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M373.3 0c-76.6 0-138.7 62.1-138.7 138.7c0 16 2.8 31.2 7.8 45.5L21.3 405.3L0 512h128v-42.7l21.3-21.3l21.3 21.3l42.7-42.7l-21.3-21.3l42.7-42.7L256 384l42.7-42.7l-21.4-21.3l50.5-50.5c14.3 5 29.6 7.8 45.5 7.8c76.6 0 138.7-62.1 138.7-138.7S449.9 0 373.3 0m32 149.3c-23.5 0-42.7-19.1-42.7-42.7S381.8 64 405.3 64c23.5 0 42.7 19.1 42.7 42.7s-19.1 42.6-42.7 42.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrow(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M513 216.6H158.5L316.1 59.1H197.9L1 256l196.9 196.9h118.2L158.5 295.4H513z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.3 205.3v-128L0 248l213.3 170.7v-128H512v-85.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftDownCornerArrow(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 69.8L442.2 0l-349 349V93.1l-93.1 93L0 512l325.9-.1l93-93.1H163z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftDownCornerArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 69.8L442.2 0L128 314.2L23.3 209.5L0 512l302.5-23.3L197.8 384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftUpCornerArrow(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M163 93.2h255.9L325.9.1L0 0l.1 325.9l93.1 93V163l349 349l69.8-69.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftUpCornerArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M197.8 128L302.5 23.3L0 0l23.3 302.5L128 197.8L442.2 512l69.8-69.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 209.1C1 221.1 0 233.3 0 245.7c0 109 59.7 203.9 148.1 254.2l16.5-34.8V227.4l-18.3-18.3zM512 227.4c0-18.3-18.3-36.6-36.6-36.6H329.1c9.1-36.6 18.3-73.1 18.3-91.4c0-36.6-18.3-73.1-27.4-82.3c-.2-.2-9.1-9.1-27.4-9.1c-27.4 0-27.4 27.4-27.4 27.4c0 .5-9.1 45.7-9.1 64s-36.6 91.4-54.9 109.7l-18.3 9.1v256l18.3 9.1h201.1c36.6 0 54.9-18.3 54.9-36.6s-18.3-36.6-36.6-36.6c36.6 0 54.9-18.3 54.9-36.6s-18.3-36.6-36.6-36.6c36.6 0 54.9-18.3 54.9-36.6s-18.3-36.6-36.6-36.6c36.5.3 54.8-18 54.8-36.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C149.3 0 64 85.3 64 192c0 36.9 11 65.4 30.1 94.3l141.7 215c4.3 6.5 11.7 10.7 20.2 10.7s16-4.3 20.2-10.7l141.7-215C437 257.4 448 228.9 448 192C448 85.3 362.7 0 256 0m0 298.6c-58.9 0-106.7-47.8-106.7-106.8S197.1 85 256 85c58.9 0 106.7 47.8 106.7 106.8S314.9 298.6 256 298.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C149.3 0 64 85.3 64 192c0 36.9 11 65.4 30.1 94.3l141.7 215c4.3 6.5 11.7 10.7 20.2 10.7s16-4.3 20.2-10.7l141.7-215C437 257.4 448 228.9 448 192C448 85.3 362.7 0 256 0m0 298.6c-58.9 0-106.7-47.8-106.7-106.8S197.1 85 256 85c58.9 0 106.7 47.8 106.7 106.8S314.9 298.6 256 298.6m0-170.6c-35.4 0-64 28.6-64 64s28.6 64 64 64s64-28.6 64-64s-28.6-64-64-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240.3 396.8c3.3 5.1 9.1 8.5 15.7 8.5s12.4-3.4 15.8-8.5L382 226.6c14.8-22.9 23.4-48.1 23.4-77.3C405.3 64.9 339 0 256 0S106.7 64.9 106.7 149.3c0 29.2 8.6 54.4 23.4 77.3zM256 64c47.1 0 85.3 38.2 85.3 85.3s-38.2 85.3-85.3 85.3s-85.3-38.2-85.3-85.3S208.9 64 256 64m109.4 259.5L256 469.3L146.6 323.5c-37.4 19.6-61.3 48.9-61.3 81.8C85.3 464.2 161.7 512 256 512s170.7-47.8 170.7-106.7c0-32.9-23.9-62.2-61.3-81.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M264 0C157.3 0 72 85.3 72 192c0 36.9 11 65.4 30.1 94.3l141.7 215c4.3 6.5 11.7 10.7 20.2 10.7s16-4.3 20.2-10.7l141.7-215C445 257.4 456 228.9 456 192C456 85.3 370.7 0 264 0m0 341.2c-82.5 0-149.3-66.9-149.3-149.5c0-82.5 66.9-149.5 149.3-149.5c82.5 0 149.3 66.9 149.3 149.5S346.5 341.2 264 341.2m64-170.5h-42.7v-64c0-11.8-9.5-21.3-21.3-21.3s-21.3 9.5-21.3 21.3V192c0 11.8 9.6 21.3 21.3 21.3h64c11.8 0 21.3-9.5 21.3-21.3s-9.5-21.3-21.3-21.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M418.9 232.7h-23.3V69.8c0-38.6-31.3-69.8-69.8-69.8H186.2c-38.6 0-69.8 31.3-69.8 69.8v162.9H93.1c-12.8 0-23.3 10.4-23.3 23.3v232.7c0 12.9 10.4 23.3 23.3 23.3h325.8c12.8 0 23.3-10.4 23.3-23.3V256c0-12.9-10.4-23.3-23.3-23.3m-69.8 0H162.9V69.8c0-12.9 10.4-23.3 23.3-23.3h139.6c12.8 0 23.3 10.4 23.3 23.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M418.4 232.7h-23.3v-93.1C395.1 62.5 332.6 0 255.5 0S115.9 62.5 115.9 139.6v93.1H92.6c-12.8 0-23.3 10.4-23.3 23.3v232.7c0 12.9 10.4 23.3 23.3 23.3h325.8c12.8 0 23.3-10.4 23.3-23.3V256c0-12.9-10.5-23.3-23.3-23.3m-69.8 0H162.4v-93.1c0-51.4 41.7-93.1 93.1-93.1s93.1 41.7 93.1 93.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M280.9 326V196.9h-78.8V326c-23.4 13.7-39.4 38.8-39.4 67.8c0 43.5 35.3 78.8 78.8 78.8s78.8-35.3 78.8-78.8c0-29-16-54.1-39.4-67.8m39.4-20V78.8C320.3 35.3 285 0 241.5 0s-78.8 35.3-78.8 78.8V306c-24.1 21.6-39.4 52.9-39.4 87.9c0 65.2 52.9 118.2 118.2 118.2c65.2 0 118.2-52.9 118.2-118.2c0-35-15.3-66.3-39.4-87.9m-78.8 186.3c-54.4 0-98.5-44.1-98.5-98.5c0-32.2 15.5-60.7 39.4-78.6V78.8c0-32.6 26.4-59.1 59.1-59.1c32.6 0 59.1 26.4 59.1 59.1v236.4c23.9 18 39.4 46.5 39.4 78.6c0 54.4-44.1 98.5-98.5 98.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magic(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M349.1 0v116.4L116.4 349.1H0V512h162.9V395.6l232.7-232.7H512V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailIcon(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 206.2L326.8 21v108.9C-10.9 129.9 0 391.4 0 500.3c76.2-217.9 174.3-217.9 326.8-217.9v108.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailIconOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M291.1 147.9V41.3L120.5 212l170.7 170.7V284c96.4 6.7 165 39.4 220.8 199c0-90.8 8.2-296.7-220.9-335.1m-120.4-26.3V41.3L0 212l170.7 170.7v-80.3L80.3 212z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailIconTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341.3 32.7V113l90.3 90.3l-90.3 90.3V374L512 203.3zm50.2 170.6L220.9 32.7v106.6C-8.2 177.6 0 383.5 0 474.4c55.9-159.6 124.5-192.3 220.8-199V374z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mark(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.3 0h-128C38.2 0 0 38.2 0 85.3v128L298.7 512L512 298.7zm-128 128c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7S128 61.8 128 85.3S108.9 128 85.3 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M418.9 0H93.1C80.2 0 69.8 10.4 69.8 23.3V512L256 325.8L442.2 512V23.3c0-12.9-10.4-23.3-23.3-23.3m-46.5 186.2H139.6v-46.5h232.7v46.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M213.3 0h-128C38.2 0 0 38.2 0 85.3v128L298.7 512L512 298.7zm-128 128c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7S128 61.8 128 85.3S108.9 128 85.3 128m85.4 192L320 170.7l42.7 42.7l-149.4 149.3zm85.3 85.3L405.3 256l42.7 42.7L298.7 448z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432.9 0H107.1C94.3 0 83.8 10.4 83.8 23.3V512L270 325.8L456.2 512V23.3c0-12.9-10.4-23.3-23.3-23.3m-46.5 186.2h-93.1v93.1h-46.5v-93.1h-93.1v-46.5h93.1V46.5h46.5v93.1h93.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M410.9 0H85.1C72.3 0 61.8 10.4 61.8 23.3V512L248 325.8L434.2 512V23.3c0-12.9-10.4-23.3-23.3-23.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Massage(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 59.1v393.8h512V59.1zm433.2 39.4L256 275.7L78.8 98.5zm39.4 315H39.4V118.2L256 334.8l216.6-216.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MassageOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78.8 56.1v78.8H0V450h433.2v-78.8H512V56.1zm354.4 39.4L295.4 233.3L157.5 95.5zm-39.4 315H39.4V193.9l39.4 39.4v137.8h315.1v39.4zm78.8-78.7H118.2V115.2l177.2 177.2l177.2-177.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Media(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M244.4 69.8L174.5 0h-58.2l69.8 69.8zm151.2 0L325.8 0h-58.2l69.8 69.8zM418.9 0l69.8 69.8H512V0zm0 162.9h-93.1l69.8-69.8h-58.2l-69.8 69.8h-93.1l69.8-69.8h-58.2l-69.8 69.8h-93l69.8-69.8H0v372.4C0 491.1 20.9 512 46.5 512h418.9c25.7 0 46.5-20.9 46.5-46.5V93.1h-23.3zM23.3 0H0v69.8h93.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MemoriCard(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416.7 0H118L32.7 85.3v384c0 23.5 19.1 42.7 42.7 42.7h341.3c23.6 0 42.7-19.1 42.7-42.7V42.7c0-23.6-19.2-42.7-42.7-42.7m-256 128H118V42.7h42.7zm64 0H182V42.7h42.7zm64 0H246V42.7h42.7zm64 0H310V42.7h42.7zm64 0H374V42.7h42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M186.2 139.6h139.6V0H186.2zM372.4 0v139.6H512V0zM0 139.6h139.6V0H0zm186.2 186.2h139.6V186.2H186.2zm186.2 0H512V186.2H372.4zM0 325.8h139.6V186.2H0zM186.2 512h139.6V372.4H186.2zm186.2 0H512V372.4H372.4zM0 512h139.6V372.4H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M155.7 320.3c2.8 23.3 24.9 42.4 49 42.4h87.8c24.1 0 46.2-19.1 49-42.4l11.7-96.6c2.8-23.3 2.8-61.4 0-84.8l-11.7-96.6C338.6 19 316.5 0 292.4 0h-87.8c-24.1 0-46.2 19.1-49 42.4L144 138.9c-2.8 23.3-2.8 61.5 0 84.8zm263.4-149.6h-42.6c.4 19.5-.3 40.2-2.2 55.6l-11.7 96.6c-4.1 34.3-34.9 61.1-70.2 61.1h-87.8c-35.2 0-66-26.9-70.2-61.1l-11.7-96.6c-1.9-15.4-2.6-36.1-2.2-55.6H77.9c-.4 21.3.4 43.6 2.5 60.7L92.1 328c6.7 55.3 56.1 98.6 112.5 98.6h22.6v42.7h-64V512h170.7v-42.7h-64v-42.7h22.6c56.5 0 105.9-43.4 112.5-98.7l11.7-96.7c2-17 2.8-39.3 2.4-60.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Missing(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M499.5 385.4L308.9 57.2c-31.8-52.9-74.1-52.9-105.9 0L12.5 385.4c-31.8 52.9 0 95.3 63.5 95.3h360c63.5 0 95.3-42.4 63.5-95.3m-201.1 52.9h-84.7v-84.7h84.7zm0-127h-84.7V120.7h84.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 256L405.3 149.3v85.4h-128v-128h85.4L256 0L149.3 106.7h85.4v128h-128v-85.4L0 256l106.7 106.7v-85.4h128v128h-85.4L256 512l106.7-106.7h-85.4v-128h128v85.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M234.7 213.3h42.7V106.7h85.3L256 0L149.3 106.7h85.3v106.6zM512 256L405.3 149.3v85.3H298.7v42.7h106.7v85.3zm-298.7-21.3H106.7v-85.3L0 256l106.7 106.7v-85.3h106.7v-42.7zm64 64h-42.7v106.7h-85.3L256 512l106.7-106.7h-85.3V298.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M209.5 0H0v209.5L81.5 128l104.7 104.7l46.5-46.5L128 81.5zm-23.3 279.3L81.5 384L0 302.5V512h209.5L128 430.5l104.7-104.7zM302.5 0L384 81.5L279.3 186.2l46.5 46.5L430.5 128l81.5 81.5V0zm23.3 279.3l-46.5 46.5L384 430.5L302.5 512H512V302.5L430.5 384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Movie(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v512h512V0zm63.9 490.7H21.3V448H64v42.7zm0-85.3H21.3v-42.7H64v42.7zm0-85.4H21.3v-42.7H64V320zm0-85.3H21.3V192H64v42.7zm0-85.4H21.3v-42.7H64v42.7zm0-85.4H21.3V21.3H64v42.6zm362.8 405.5H85.3V277.3h341.4zm0-234.7H85.3V42.6h341.4zm64 256H448V448h42.7zm0-85.3H448v-42.7h42.7zm0-85.4H448v-42.7h42.7zm0-85.3H448V192h42.7zm0-85.4H448v-42.7h42.7zm0-85.4H448V21.3h42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Multiply(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m121.3 316.5c8.4 8.4 8.4 22 0 30.4l-30.4 30.4c-8.4 8.4-22 8.4-30.4 0l-60.8-60.8l-60.8 60.8c-8.4 8.4-22 8.4-30.4 0L134 346.9c-8.4-8.4-8.4-22 0-30.4l60.8-60.8l-60.8-60.9c-8.4-8.4-8.4-22 0-30.4l30.4-30.4c8.4-8.4 22-8.4 30.4 0l60.8 60.8l60.8-60.8c8.4-8.4 22-8.4 30.4 0l30.4 30.4c8.4 8.4 8.4 22 0 30.4l-60.8 60.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplyOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M513 76.8L436.2 0L257 179.2L77.8 0L1 76.8L180.2 256L1 435.2L77.8 512L257 332.8L436.2 512l76.8-76.8L333.8 256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M165 85.3v325.3c-9.7-3.3-20.5-5.3-32-5.3c-41.2 0-74.7 23.9-74.7 53.3c0 29.5 33.4 53.3 74.7 53.3c41.2 0 74.7-23.9 74.7-53.3c0-3.6-.5-7.2-1.5-10.7h1.5V179.8L421 118.9v227.8c-9.7-3.3-20.5-5.3-32-5.3c-41.2 0-74.7 23.9-74.7 53.3c0 29.5 33.4 53.3 74.7 53.3c41.2 0 74.7-23.9 74.7-53.3c0-3.4-.5-6.7-1.4-10l1.4-.7V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Musk(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 377.8c21-.4 19.3-9.4 41.8-9.1c34.5.6 43.1-13.6 43.5-36.1c.5-26.5-31.6-43.8-34.9-51.6c-3.3-7.8-25.1-120.7-25.1-120.7s-3.1-8.6-25.3-8.6s-25.3 8.6-25.3 8.6s-21.9 112.9-25.1 120.7c-3.3 7.8-35.4 25.1-34.9 51.6c.4 22.5 9 36.7 43.5 36.1c22.5-.4 20.8 8.7 41.8 9.1M511.6 20c-35-7-69.4-20-105-19.8c-20.4.1-40.9 3.6-61.1 6.9c-22.7 3.7-44.7 12-67.3 15.1c-7.3 1-14.7 1.4-22.1 1.4c-7.4 0-14.8-.4-22.1-1.4c-22.7-3.1-44.6-11.4-67.3-15.1C146.5 3.8 125.9.3 105.6.2C69.8 0 35.3 13 .4 20c0 11.9-2.1 27.1 3.1 35.5C8.9 64.1 18 72 20.4 83.7c5.6 27.7 13.5 54.9 23.7 80.9c4.9 12.6 14.2 21.5 24.6 27.6c25 14.6 54.7 16.8 82.1 12.2c39.1-6.6 60.3-40.7 76.5-79.7c6-15.9 6.7-34.2 14.7-48.6c3-5.4 8.5-6 14-5.7c5.6-.3 11.1.2 14 5.7c8 14.5 8.6 32.8 14.7 48.6c16.2 39 37.4 73.2 76.5 79.7c27.5 4.6 57.1 2.5 82.1-12.2c10.4-6.1 19.7-14.9 24.6-27.6c10.1-26 18.1-53.2 23.6-80.9c2.3-11.7 11.4-19.6 16.8-28.2c5.4-8.3 3.3-23.6 3.3-35.5M204.3 131.2c-19.4 46.8-54.3 56.6-80.1 56.6c-10.1 0-20.6-1.5-31.3-4.5c-19.1-5.3-33.3-21-40-44.2c-9-31-12.7-55.9-11.6-78.3c1.4-29 27.9-34.9 57.3-38.7c13.7-1.8 24.7-2.7 34.6-2.7h.1c13.3 0 23.8 1.2 33.6 3.6c28.5 7 44.7 16.7 50.9 30.5c6.9 15.4 3 37.9-13.5 77.7m254.8 7.9c-6.7 23.2-21 38.9-40.1 44.2c-10.7 3-21.2 4.5-31.3 4.5c-25.8 0-60.7-9.8-80.1-56.6c-16.5-39.8-20.4-62.3-13.4-77.7c6.2-13.8 22.4-23.5 50.9-30.5c9.7-2.4 20.3-3.6 33.6-3.6h.1c9.9.1 20.9.9 34.6 2.7c29.4 3.9 55.9 9.7 57.3 38.7c1.1 22.4-2.6 47.3-11.6 78.3m-83.4 297c-68.4-38-110.1-12-119.8-4.8c-9.7-7.2-51.3-33.3-119.8 4.8C68.1 474 34 408.7 30 399.1c6.3 15.4 80.4 178 180.5 83.5c15.9-15 31.6-19.1 45.4-18.9c13.8-.1 29.5 4 45.4 18.9c100.1 94.5 174.2-68.2 180.5-83.5c-3.9 9.7-38 74.9-106.1 37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mute(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m159.8 320l64-64l-64-64l85.3-85.3l64 64V0L159.8 149.3H74.4v213.3h85.3L309.1 512V341.3l-64 64zm245.3-128l-32-32l-64 64l-64-64l-32 32l64 64l-64 64l32 32l64-64l64 64l32-32l-64-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Netwark(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 247.2c-57.1 0-106.1 34-128.2 82.9l41.9 41.9c7.5-40.9 43.2-72 86.3-72c43.1 0 78.8 31.1 86.3 72l41.9-41.9c-22.1-48.9-71.1-82.9-128.2-82.9m0-175.9C150.6 71.3 57.5 123 0 202.2l37.6 37.6C85.1 170 165.2 124.1 256 124.1c90.8 0 170.9 45.9 218.4 115.8l37.6-37.6c-57.5-79.3-150.6-131-256-131m0 88c-81.1 0-152.3 42.4-192.8 106.1l38.4 38.4C131.4 249.1 189.3 212 256 212s124.6 37.1 154.5 91.7l38.4-38.4c-40.6-63.6-111.8-106-192.9-106m0 175.8c-29.1 0-52.8 23.6-52.8 52.8c0 29.1 23.6 52.8 52.8 52.8c29.1 0 52.8-23.6 52.8-52.8c0-29.1-23.7-52.8-52.8-52.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 381.4l237.7-118.9L0 143.6zm237.7-118.9v118.9l237.7-118.9l-237.7-118.9zm237.7-118.9v237.8H512V143.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469.3 95.3H448c-23.5 0-42.7 19.1-42.7 42.7v52l-152.7-89.8c-21.6-12.7-39.2.1-39.2 28.5v73.9L39.2 100.1C17.6 87.4 0 100.2 0 128.6v274.8c0 28.3 17.6 41.2 39.2 28.5l174.2-102.5v74c0 28.3 17.6 41.2 39.2 28.5l152.7-89.8V394c0 23.5 19.1 42.7 42.7 42.7h21.3c23.5 0 42.7-19.1 42.7-42.7V138c0-23.6-19.1-42.7-42.7-42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 451h512v-64H0zm0-106.7h512v-64H0zm0-106.6h512v-64H0zM0 67v64h512V67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphEight(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 351.3h85.3v-64H0zM0 138h85.3V74H0zm0 320h85.3v-64H0zm0-213.3h85.3v-64H0zM170.7 458H512v-64H170.7zm0-384v64H512V74zm0 170.7H512v-64H170.7zm0 106.6H512v-64H170.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 462h256v-64H0zm0-106.7h512v-64H0zm0-106.6h512v-64H0zM0 78v64h512V78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 464h256v-64H128zM0 357.3h512v-64H0zm128-170.6v64h256v-64zM0 80v64h512V80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphNine(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 236.7h341.3v-64H0zm0 106.6h341.3v-64H0zM0 130h341.3V66H0zm0 320h341.3v-64H0zm426.7-213.3H512v-64h-85.3zm0 213.3H512v-64h-85.3zm0-384v64H512V66zm0 277.3H512v-64h-85.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphSeven(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 464h256v-64H128zM0 357.3h512v-64H0zm0-106.6h512v-64H0zM0 80v64h512V80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphSix(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 456.5h256v-64H256zM0 349.8h512v-64H0zm0-106.6h512v-64H0zM0 72.5v64h512v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 458.5h256v-64H256zM0 351.8h512v-64H0zm256-106.6h256v-64H256zM0 74.5v64h512v-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 462h256v-64H0zm0-106.7h512v-64H0zm256-170.6H0v64h256zM0 78v64h512V78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PartOfCircle(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78.2 0v512L453 164C359.5 63.4 226.4 0 78.2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PartOfCircleFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0v256L75 75C28.6 121.3 0 185.3 0 256c0 141.4 114.6 256 256 256s256-114.6 256-256S397.4 0 256 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PartOfCircleFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0v256H0c0 141.4 114.6 256 256 256s256-114.6 256-256S397.4 0 256 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PartOfCircleOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v512h512C512 229.2 282.8 0 0 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PartOfCircleThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M221.8 0v256l-181 181c46.4 46.3 110.4 75 181 75c141.4 0 256-114.6 256-256S363.2 0 221.8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PartOfCircleTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M119.5 0v299.9L331.6 512c54.3-54.3 87.8-129.2 87.8-212.1C419.5 134.3 285.2 0 119.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Passing(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M186.2 232.7h139.6v-46.5H186.2zm0 93.1h139.6v-46.5H186.2zm0 93.1h139.6v-46.5H186.2zm0 93.1h139.6v-46.5H186.2zm186.2-93.1H512v-46.5H372.4zm0 93.1H512v-46.5H372.4zm0-279.3H512v-46.5H372.4zm0-139.6v46.5H512V93.1zm0 232.7H512v-46.5H372.4zM0 232.7h139.6v-46.5H0zm0 93.1h139.6v-46.5H0zm0 93.1h139.6v-46.5H0zM0 46.5h139.6V0H0zm0 93.1h139.6V93.1H0zM0 512h139.6v-46.5H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 512h162.9V0H0zM349.1 0v512H512V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M139.6 0H46.5C20.9 0 0 20.9 0 46.5v418.9C0 491.1 20.9 512 46.5 512h93.1c25.7 0 46.5-20.9 46.5-46.5v-419C186.2 20.9 165.3 0 139.6 0m325.9 0h-93.1c-25.7 0-46.5 20.9-46.5 46.5v418.9c0 25.7 20.9 46.5 46.5 46.5h93.1c25.7 0 46.5-20.9 46.5-46.5V46.5C512 20.9 491.1 0 465.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m70.2 337.4l104.4 104.4L441.5 175L337 70.5zM.6 499.8c-2.3 9.3 2.3 13.9 11.6 11.6L151.4 465L47 360.6zM487.9 24.1c-46.3-46.4-92.8-11.6-92.8-11.6c-7.6 5.8-34.8 34.8-34.8 34.8l104.4 104.4s28.9-27.2 34.8-34.8c0 0 34.8-46.3-11.6-92.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M387 128c-11.8 0-21.3 9.5-21.3 21.3v213.3c0 58.9-47.7 106.7-106.7 106.7s-106.7-47.8-106.7-106.7v-256c0-35.3 28.6-64 64-64s64 28.7 64 64V320c0 11.8-9.6 21.3-21.3 21.3s-21.3-9.5-21.3-21.3V149.3c0-11.8-9.6-21.3-21.3-21.3c-11.8 0-21.3 9.5-21.3 21.3V320c0 35.4 28.6 64 64 64s64-28.6 64-64V106.7C323 47.8 275.2 0 216.3 0S109.7 47.8 109.7 106.7v256c0 82.5 66.9 149.3 149.3 149.3s149.3-66.9 149.3-149.3V149.3c0-11.8-9.5-21.3-21.3-21.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M59.1 236.3h393.8s-19.7-98.5-98.5-98.5V59.1L413.5 0h-315l59.1 59.1v78.8c-78.8-.1-98.5 98.4-98.5 98.4M256 512l39.4-256h-78.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v512l512-256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464.7 221.5L86.1 7.3C52.5-11.7 25 7.5 25 50v412c0 42.5 27.5 61.7 61.1 42.7l378.6-214.1c33.5-19.1 33.5-50.1 0-69.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pound(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192.9 447.7c42.6-21.3 53.3-53.3 53.3-106.6v-42.6h106.6v-64H246.2v-85.3c0-62.5 32.1-95.9 85.3-95.9c25.8 0 40.8 18.6 64 32V10.7C370.6 1.1 361.7 0 331.5 0C286.7 0 243 13.6 214.8 40.8c-28.1 27.1-32.5 66.1-32.5 108.5v85.3H97v64h85.3v64c0 42.6-42.6 85.3-85.3 85.3l.3 64.3H438l.1-64.3H192.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M439.6 0H204.9L55.4 256h149.5l-128 256l341.3-320H247.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerBatton(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M334.5 53.1v64c58.3 29 98.5 89 98.5 158.6c0 97.9-79.3 177.2-177.2 177.2S78.5 373.6 78.5 275.7c0-69.6 40.2-129.6 98.5-158.6v-64C85.2 85.6 19.4 172.8 19.4 275.7C19.4 406.2 125.2 512 255.7 512S492 406.2 492 275.7c0-102.9-65.8-190.1-157.5-222.6m-78.8 143.8c21.7 0 39.4-17.7 39.4-39.4V39.4c0-21.7-17.6-39.4-39.4-39.4s-39.4 17.7-39.4 39.4v118.2c0 21.7 17.7 39.3 39.4 39.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M274.3 262.5L512 381.4V143.6zm-237.7 0l237.7 118.9V143.6zM0 143.6v237.7h36.6V143.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PreviousOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472.8 90.1L298.6 192.5v-73.9c0-28.4-17.6-41.2-39.2-28.5l-152.8 89.8v-52c0-23.6-19.1-42.7-42.7-42.7H42.7C19.1 85.3 0 104.4 0 128v256c0 23.5 19.1 42.7 42.7 42.7H64c23.5 0 42.7-19.1 42.7-42.7v-51.9l152.8 89.8c21.5 12.7 39.2-.2 39.2-28.5v-74l174.2 102.5c21.5 12.7 39.2-.2 39.2-28.5V118.6c-.1-28.4-17.7-41.2-39.3-28.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 362.7H128V384h256zM106.7 21.3h192V128h106.7v42.7h21.3v-64L320 0H85.3v170.7h21.3V21.3zM448 192H64c-42.7 0-64 21.3-64 64v128h85.3v128h341.3V384H512V256c0-42.7-21.3-64-64-64M85.3 277.3H42.7v-42.7h42.7v42.7zm320 213.4H106.7V341.3h298.7v149.4zM384 405.3H128v21.3h256zm0 42.7H128v21.3h256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Random(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341.3 28.3v85.3H128c-70.7 0-128 57.3-128 128c0 21.5 5.8 41.4 15.2 59.2L68 263.2c-2.4-6.8-4-13.9-4-21.5c0-35.4 28.7-64 64-64h213.3V263L512 156.3V135zM444 262.8c2.4 6.8 4 13.9 4 21.5c0 35.4-28.6 64-64 64H170.7V263L0 369.7V391l170.7 106.7v-85.3H384c70.7 0 128-57.3 128-128c0-21.5-5.8-41.4-15.2-59.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rectangle(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v512h512V0zm465.5 465.5h-419v-419h418.9v419z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M47.5 280.3H1v93.1h46.5zM1 94.1h46.5V47.5H94V1H1zm46.5 325.8H1V513h93.1v-46.5H47.5zM373.4 1h-93.1v46.5h93.1zM47.5 140.6H1v93.1h46.5zM140.6 513h93.1v-46.5h-93.1zM419.9 1v46.5h46.5V94H513V1zM140.6 47.5h93.1V1h-93.1zm325.9 186.2H513v-93.1h-46.5zM280.3 513H513V280.3H280.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M46.5 279.3H0v93.1h46.5zM279.3 46.5h93.1V0h-93.1zm0 465.5h93.1v-46.5h-93.1zM46.5 418.9H0V512h93.1v-46.5H46.5zm93.1 93.1h93.1v-46.5h-93.1zM418.9 0v46.5h46.5V93H512V0zm46.6 232.7H512v-93.1h-46.5zm0 232.8H419V512h93v-93.1h-46.5zm0-93.1H512v-93.1h-46.5zM0 232.7h232.7V0H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M139.6 46.5h93.1V0h-93.1zM418.9 0v46.5h46.5V93H512V0zM46.5 46.5H93V0H0v93.1h46.5zM279.3 512h93.1v-46.5h-93.1zM46.5 139.6H0v93.1h46.5zm419 93.1H512v-93.1h-46.5zm0 139.7H512v-93.1h-46.5zm0 93.1H419V512h93v-93.1h-46.5zm-186.2-419h93.1V0h-93.1zM0 512h232.7V279.3H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M46.5 279.3H0v93.1h46.5zm0 139.6H0V512h93.1v-46.5H46.5zm93.1 93.1h93.1v-46.5h-93.1zM46.5 139.6H0v93.1h46.5zM0 93.1h46.5V46.5H93V0H0zm465.5 372.4H419V512h93v-93.1h-46.5zm0-93.1H512v-93.1h-46.5zM139.6 46.5h93.1V0h-93.1zM279.3 512h93.1v-46.5h-93.1zm0-512v232.7H512V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rectangular(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M139.6 46.5h93.1V0h-93.1zM46.5 418.9H0V512h93.1v-46.5H46.5zm93.1 93.1h93.1v-46.5h-93.1zM46.5 279.3H0v93.1h46.5zm0-139.7H0v93.1h46.5zM0 93.1h46.5V46.5H93V0H0zm465.5 139.6H512v-93.1h-46.5zM279.3 512h93.1v-46.5h-93.1zM418.9 0v46.5h46.5V93H512V0zm46.6 372.4H512v-93.1h-46.5zm0 93.1H419V512h93v-93.1h-46.5zm-186.2-419h93.1V0h-93.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M354.5 65.1H236.3l118.2 118.2H137.8C61.7 183.2 0 244.9 0 321.1c0 76.1 61.7 137.8 137.8 137.8v-78.8c-32.6 0-59.1-26.4-59.1-59.1c0-32.6 26.4-59.1 59.1-59.1h216.6L236.3 380.2h118.2L512 222.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedoIcon(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 192V21.3l-64.9 64.9C400.3 33.4 332.2 0 256 0C114.6 0 0 114.6 0 256s114.6 256 256 256c70.7 0 134.7-28.6 181-75l-45.3-45.2C357 426.5 309 448 256 448c-106 0-192-85.9-192-192S150 64 256 64c58.5 0 110.4 26.5 145.5 67.8L341.3 192z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedoOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M389.6 289.4c0 73.8-59.8 133.6-133.6 133.6c-73.7 0-133.6-59.8-133.6-133.6S182.2 155.8 256 155.8v66.8l113.9-111.3L256 0v66.8c-122.9 0-222.6 99.7-222.6 222.6C33.4 412.3 133.1 512 256 512c122.9 0 222.6-99.7 222.6-222.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshTime(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C179.9 0 111.7 33.4 64.9 86.2L0 21.3V192h170.7l-60.2-60.2C145.6 90.5 197.5 64 256 64c106 0 192 85.9 192 192s-86 192-192 192c-53 0-101-21.5-135.8-56.2L75 437c46.4 46.3 110.4 75 181 75c141.4 0 256-114.6 256-256S397.4 0 256 0m-21.3 106.7v170.7h128v-42.7h-85.3v-128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemovePlaylist(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M465.5 0H186.2c-25.7 0-46.5 20.9-46.5 46.5v279.3c0 25.7 20.9 46.5 46.5 46.5h279.3c25.7 0 46.5-20.9 46.5-46.5V46.5C512 20.9 491.1 0 465.5 0m-23.3 209.5H209.5V163h232.7zM93.1 325.8V139.6H46.5C20.9 139.6 0 160.5 0 186.2v279.3C0 491.1 20.9 512 46.5 512h279.3c25.7 0 46.5-20.9 46.5-46.5V419H186.2c-51.4-.1-93.1-41.8-93.1-93.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M185.2 128.6V19.7L0 204.9l185.2 185.2v-109c152.5 0 250.5 0 326.8 217.9c0-108.9 10.9-370.4-326.8-370.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightArrow(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M315.1 48.6H196.9l157.6 157.5H0v78.8h354.5L196.9 442.4h118.2L512 245.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 261.5L298.7 90.8v128H0v85.4h298.7v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightDownCornerArrow(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m511.9 186.1l-93.1-93V349L69.8 0L0 69.8l349 349H93.1l93 93.1l325.9.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightDownCornerArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m512 512l-23.3-302.5L384 314.2L69.8 0L0 69.8L314.2 384L209.5 488.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightUpCornerArrow(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m186.1.1l-93 93.1H349L0 442.2L69.8 512l349-349v255.9l93.1-93L512 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightUpCornerArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M209.5 23.3L314.2 128L0 442.2L69.8 512L384 197.8l104.7 104.7L512 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundArrowFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M447.1 86.2C400.3 33.4 332.2 0 256 0C114.6 0 0 114.6 0 256h64c0-106.1 85.9-192 192-192c58.5 0 110.4 26.5 145.5 67.8L341.3 192H512V21.3zM256 448c-58.5 0-110.4-26.5-145.5-67.8l60.2-60.2H0v170.7l64.9-64.9c46.8 52.8 115 86.2 191.1 86.2c141.4 0 256-114.6 256-256h-64c0 106.1-85.9 192-192 192m42.7-192c0-23.6-19.1-42.7-42.7-42.7s-42.7 19.1-42.7 42.7s19.1 42.7 42.7 42.7s42.7-19.1 42.7-42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundArrowFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M254.5 157.5h-95.9c24.8-24.2 58.5-39.4 95.9-39.4c69.4 0 126.2 51.4 135.9 118.2h79.8c-10-110.4-102.6-196.9-215.6-196.9c-62.4 0-118.1 26.8-157.5 69V0L37.9 59.1v157.5h157.5zm59.1 137.9l-59.1 59.1h95.9c-24.8 24.2-58.5 39.4-95.9 39.4c-69.4 0-126.2-51.4-135.9-118.2H38.9c10 110.4 102.6 196.9 215.6 196.9c62.4 0 118.1-26.8 157.5-69V512l59.1-59.1V295.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M447.1 86.2C400.3 33.4 332.2 0 256 0C114.6 0 0 114.6 0 256h64c0-106.1 85.9-192 192-192c58.5 0 110.4 26.5 145.5 67.8L341.3 192H512V21.3zM256 448c-58.5 0-110.4-26.5-145.5-67.8l60.2-60.2H0v170.7l64.9-64.9c46.8 52.8 115 86.2 191.1 86.2c141.4 0 256-114.6 256-256h-64c0 106.1-85.9 192-192 192"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundArrowSix(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M110.5 131.8C145.6 90.5 197.5 64 256 64c106.1 0 192 85.9 192 192h64C512 114.6 397.4 0 256 0C179.8 0 111.7 33.4 64.9 86.2L0 21.3V192h170.7zm291 248.4c-35.2 41.3-87 67.8-145.5 67.8c-106.1 0-192-85.9-192-192H0c0 141.4 114.6 256 256 256c76.2 0 144.3-33.4 191.1-86.2l64.9 64.9V320H341.3zM213.3 256c0 23.6 19.1 42.7 42.7 42.7s42.7-19.1 42.7-42.7s-19.1-42.7-42.7-42.7s-42.7 19.1-42.7 42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundArrowThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.6 393.8c-37.4 0-71.1-15.2-95.9-39.4h95.9l-59.1-59.1H31v157.5L90.1 512V403.6c39.5 42.2 95.2 69 157.5 69c113 0 205.7-86.5 215.6-196.9h-79.8c-9.6 66.7-66.4 118.1-135.8 118.1M405.2 0v108.4c-39.5-42.2-95.2-69-157.5-69C134.6 39.4 42 125.9 32 236.3h79.8c9.6-66.7 66.5-118.2 135.9-118.2c37.4 0 71.1 15.2 95.9 39.4h-95.9l59.1 59.1h157.5V59.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundArrowTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M110.5 131.8C145.6 90.5 197.5 64 256 64c106.1 0 192 85.9 192 192h64C512 114.6 397.4 0 256 0C179.8 0 111.7 33.4 64.9 86.2L0 21.3V192h170.7zm291 248.4c-35.2 41.3-87 67.8-145.5 67.8c-106.1 0-192-85.9-192-192H0c0 141.4 114.6 256 256 256c76.2 0 144.3-33.4 191.1-86.2l64.9 64.9V320H341.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C128 149.3 0 42.7 0 42.7V128c0 170.7 149.3 384 256 384s256-213.3 256-384V42.7S384 149.3 256 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.8 0C223 0 139.6 83.4 139.6 186.2c0 33.5 9 64.8 24.4 92L0 442.2l23.3 46.5L69.8 512l164-164c27.1 15.5 58.5 24.4 92 24.4C428.6 372.4 512 289 512 186.2S428.6 0 325.8 0m0 314.2c-70.7 0-128-57.3-128-128s57.3-128 128-128s128 57.3 128 128s-57.3 128-128 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settong(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m447.1 259.8l64.9-27.1l-19.8-83.2l-72.7 7.8c-7.3-12.4-15.9-23.7-25.4-33.8L421 58.1l-72.8-44.9l-45.4 56.5c-15.7-4.1-32-6.2-48.3-6.3L225.9 0l-82.7 22.2l9.7 70.2c-11.4 7.1-21.9 15.1-31.2 24L59 90.5l-44.9 72.8l54 43.7c-3.9 14.9-6.3 30.3-6.6 45.8L0 278.4l19.8 83.3l68.7-7.4c7.9 13.8 17.2 26.3 28 37.3L91.2 453l72.8 44.9l43.6-54.2c14.8 3.7 29.9 5.7 45.3 5.9l25.8 62.4l83.2-19.8l-7.5-70.8c12.9-7.8 24.7-16.8 34.9-27l63.8 26.3l44.9-72.9l-56.4-45.6c3.4-13.8 5.3-28.1 5.5-42.4M256 374.4c-65.2 0-118-53-118-118.5c0-65.4 52.8-118.5 118-118.5s118 53 118 118.5s-52.9 118.5-118 118.5m0-158c-21.9 0-39.6 17.7-39.6 39.6c0 21.9 17.7 39.6 39.6 39.6s39.6-17.7 39.6-39.6c0-21.9-17.8-39.6-39.6-39.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426.7 341.3c-29.1 0-54.8 14.7-70.2 36.9L167.4 278.1c1.9-7.1 3.3-14.4 3.3-22.1c0-7.7-1.3-15-3.3-22.1l189.1-100.1c15.4 22.3 41 36.9 70.2 36.9c47.1 0 85.3-38.2 85.3-85.3S473.8 0 426.7 0s-85.3 38.2-85.3 85.3c0 .8.2 1.6.2 2.4l-199 105.4c-15.2-13.8-35.1-22.4-57.2-22.4C38.2 170.7 0 208.9 0 256s38.2 85.3 85.3 85.3c22.1 0 42-8.6 57.2-22.4l199 105.4c0 .8-.2 1.6-.2 2.4c0 47.1 38.2 85.3 85.3 85.3s85.3-38.2 85.3-85.3s-38.1-85.4-85.2-85.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426.7 341.3c-29.1 0-54.8 14.6-70.2 36.9L167.4 278.1c1.9-7.1 3.3-14.4 3.3-22.1c0-7.7-1.3-15-3.3-22.1l189.1-100.1c15.4 22.3 41 36.9 70.2 36.9c47.1 0 85.3-38.2 85.3-85.3S473.8 0 426.7 0s-85.3 38.2-85.3 85.3c0 .8.2 1.6.2 2.4l-199 105.4c-15.1-13.8-35.1-22.4-57.2-22.4C38.2 170.6 0 208.9 0 256s38.2 85.3 85.3 85.3c22.1 0 42-8.6 57.2-22.4l199 105.4c0 .9-.2 1.6-.2 2.4c0 47.1 38.2 85.3 85.3 85.3s85.3-38.2 85.3-85.3s-38.1-85.4-85.2-85.4m0-298.6c23.5 0 42.7 19.1 42.7 42.7c0 23.5-19.1 42.7-42.7 42.7c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7m-341.4 256c-23.5 0-42.7-19.1-42.7-42.7c0-23.5 19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7c0 23.5-19.1 42.7-42.7 42.7m341.4 170.6c-23.5 0-42.7-19.1-42.7-42.7c0-23.5 19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7c-.1 23.6-19.2 42.7-42.7 42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sharing(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.8 0v186.2h69.8v46.5H116.4v-46.5h69.8V0H0v186.2h69.8v93.1h162.9v46.5h-69.8V512h186.2V325.8h-69.8v-46.5h162.9v-93.1H512V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SharingOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.8 0v186.2h69.8v46.5H116.4v-46.5h69.8V0H0v186.2h69.8v93.1h162.9v46.5h-69.8V512h186.2V325.8h-69.8v-46.5h162.9v-93.1H512V0zM46.5 139.6V46.5h93.1v93.1zm256 232.8v93.1h-93.1v-93.1zm163-232.8h-93.1V46.5h93.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffile(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341.3 0v85.3h-64c-70.7 0-128 57.3-128 128v85.3c0 35.4-28.6 64-64 64H0v64h85.3c70.7 0 128-57.3 128-128v-85.3c0-35.4 28.6-64 64-64h64v85.3L512 128v-21.3zM114 156.4l37.6-52.1c-19.4-11.8-42-19-66.3-19H0v64h85.3c10.4 0 20 2.7 28.7 7.1m227.3 206.3h-64c-10.4 0-20-2.7-28.7-7.1L211 407.7c19.4 11.8 42 19 66.3 19h64V512L512 405.3V384L341.3 277.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sms(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0H64C28.6 0 0 28.6 0 64v256c0 35.4 28.6 64 64 64h128l-42.7 128l192-128H448c35.4 0 64-28.6 64-64V64c0-35.4-28.6-64-64-64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmsEight(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 85.9 0 192c0 75 57.5 139.8 141.1 171.4L85.3 512l160.5-128.4c3.4.1 6.7.4 10.2.4c141.4 0 256-85.9 256-192S397.4 0 256 0M128 234.7c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7s-19.2 42.7-42.7 42.7m128 0c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7s-19.2 42.7-42.7 42.7m128 0c-23.5 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7c23.5 0 42.7 19.1 42.7 42.7s-19.2 42.7-42.7 42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmsFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 85.9 0 192c0 75 57.5 139.8 141.1 171.4L85.3 512l160.5-128.4c3.4.1 6.7.4 10.2.4c141.4 0 256-85.9 256-192S397.4 0 256 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmsFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0H64C28.6 0 0 28.6 0 64v256c0 35.4 28.6 64 64 64h128l-42.7 128l192-128H448c35.4 0 64-28.6 64-64V64c0-35.4-28.6-64-64-64m-74.7 256L320 309.3l-64-64l-64 64l-53.3-53.3l64-64l-64-64L192 74.7l64 64l64-64l53.3 53.3l-64 64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmsNine(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 85.9 0 192c0 75 57.5 139.8 141.1 171.4L85.3 512l160.5-128.4c3.4.1 6.7.4 10.2.4c141.4 0 256-85.9 256-192S397.4 0 256 0m117.3 256L320 309.3l-64-64l-64 64l-53.3-53.3l64-64l-64-64L192 74.7l64 64l64-64l53.3 53.3l-64 64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmsOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0H64C28.6 0 0 28.6 0 64v256c0 35.4 28.6 64 64 64h128l-42.7 128l192-128H448c35.4 0 64-28.6 64-64V64c0-35.4-28.6-64-64-64m-64 213.3H277.3V320h-42.7V213.3H128v-42.7h106.7V64h42.7v106.7H384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmsSeven(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 85.9 0 192c0 75 57.5 139.8 141.1 171.4L85.3 512l160.5-128.4c3.4.1 6.7.4 10.2.4c141.4 0 256-85.9 256-192S397.4 0 256 0m128 213.3H128v-42.7h256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmsSix(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M257 0C115.6 0 1 85.9 1 192c0 75 57.5 139.8 141.1 171.4L86.3 512l160.5-128.4c3.4.1 6.7.4 10.2.4c141.4 0 256-85.9 256-192S398.4 0 257 0m128 213.3H278.3V320h-42.7V213.3H129v-42.7h106.7V64h42.7v106.7H385z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmsThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0H64C28.6 0 0 28.6 0 64v256c0 35.4 28.6 64 64 64h128l-42.7 128l192-128H448c35.4 0 64-28.6 64-64V64c0-35.4-28.6-64-64-64M128 234.7c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7s42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7m128 0c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7s42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7m128 0c-23.6 0-42.7-19.1-42.7-42.7s19.1-42.7 42.7-42.7s42.7 19.1 42.7 42.7s-19.1 42.7-42.7 42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmsTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0H64C28.6 0 0 28.6 0 64v256c0 35.4 28.6 64 64 64h128l-42.7 128l192-128H448c35.4 0 64-28.6 64-64V64c0-35.4-28.6-64-64-64m-64 213.3H128v-42.7h256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sound(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426.7 256c0-71-43.4-131.8-105-157.5l-16.4 39.4C351.5 157.2 384 202.8 384 256c0 53.3-32.5 98.8-78.8 118.1l16.4 39.4C383.3 387.8 426.7 327 426.7 256m-85.4 0c0-35.5-21.7-65.9-52.5-78.7l-16.4 39.4c15.4 6.4 26.2 21.6 26.2 39.4c0 17.7-10.8 32.9-26.2 39.4l16.4 39.4c30.8-13 52.5-43.4 52.5-78.9m13.2-236.3L338 59.1C415.1 91.2 469.3 167.2 469.3 256c0 88.7-54.2 164.8-131.3 196.9l16.4 39.4C447 453.7 512 362.5 512 256S447 58.3 354.5 19.7M0 149.3v213.3h85.3L234.7 512V0L85.3 149.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M391 256c0-35.5-21.7-65.9-52.5-78.7l-16.4 39.4c15.4 6.4 26.2 21.6 26.2 39.4c0 17.7-10.8 32.9-26.2 39.4l16.4 39.4c30.8-13 52.5-43.4 52.5-78.9M371.3 98.5l-16.4 39.4c46.3 19.3 78.8 64.9 78.8 118.1c0 53.3-32.5 98.8-78.8 118.1l16.4 39.4c61.7-25.7 105-86.5 105-157.5S433 124.2 371.3 98.5M49.7 149.3v213.3H135L284.3 512V0L135 149.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m376.1 177.3l-16.4 39.4c15.4 6.4 26.2 21.6 26.2 39.4c0 17.7-10.8 32.9-26.2 39.4l16.4 39.4c30.8-12.9 52.5-43.3 52.5-78.8c.1-35.6-21.6-66-52.5-78.8m-288.8-28v213.3h85.3L322 512V0L172.7 149.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 207.9H315.1L256 11l-59.1 196.9H0l157.5 108.3l-59 187.1L256 404.8l157.5 98.5l-59-187.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 206.8H315.1L256 9.8l-59.1 196.9H0L157.5 315l-59 187.2L256 403.7l157.5 98.5l-59.1-187.1zM256 354.5l-88.6 59.1l39.4-108.3l-78.8-59.1h98.5l29.5-98.5l29.5 98.5H384l-78.8 59.1l39.4 108.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Step(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M217 39.5C95.5 49.5 0 151.1 0 275.2C0 406 106 512 236.8 512c124.1 0 225.7-95.5 235.8-217H217zM256.5 0v255.5H512C502.3 118.7 393.3 9.7 256.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M140.6 280.3H1v139.6h139.6zm186.2-139.7H187.2v279.3h139.6zM373.4 1v418.9H513V1zM1 513h512v-46.5H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.8 279.3H186.2v139.6h139.6zm46.6-139.7v279.3H512V139.6zM139.6 0H0v418.9h139.6zM0 512h512v-46.5H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h512v512H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M465.5 0h-419C20.9 0 0 20.9 0 46.5v418.9C0 491.1 20.9 512 46.5 512h418.9c25.7 0 46.5-20.9 46.5-46.5v-419C512 20.9 491.1 0 465.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subtraction(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m149.3 277.3c0 11.8-9.5 21.3-21.3 21.3H128c-11.8 0-21.3-9.6-21.3-21.3v-42.7c0-11.8 9.5-21.3 21.3-21.3h256c11.8 0 21.3 9.6 21.3 21.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubtractionOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 229.3h512v85.3H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Switch(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469.3 362.7H42.7C19.1 362.7 0 381.8 0 405.3V448c0 23.5 19.1 42.7 42.7 42.7h426.7c23.5 0 42.7-19.1 42.7-42.7v-42.7c-.1-23.5-19.2-42.6-42.8-42.6M85.3 448H42.7v-42.7h42.7V448zm298.7 0h-21.3v-42.7H384zm42.7 0h-21.3v-42.7h21.3zm42.6 0H448v-42.7h21.3zm0-256H42.7C19.1 192 0 211.1 0 234.7v42.7C0 300.9 19.1 320 42.7 320h426.7c23.5 0 42.7-19.1 42.7-42.7v-42.7c-.1-23.5-19.2-42.6-42.8-42.6m-384 85.3H42.7v-42.7h42.7v42.7zm298.7 0h-21.3v-42.7H384zm42.7 0h-21.3v-42.7h21.3zm42.6 0H448v-42.7h21.3zm0-256H42.7C19.1 21.3 0 40.5 0 64v42.7c0 23.5 19.1 42.7 42.7 42.7h426.7c23.5 0 42.7-19.1 42.7-42.7V64c-.1-23.5-19.2-42.7-42.8-42.7m-384 85.4H42.7V64h42.7v42.7zm298.7 0h-21.3V64H384zm42.7 0h-21.3V64h21.3zm42.6 0H448V64h21.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Symbol(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M387.1 64.2c-7.8-14.1-18.5-25.9-32-35.4c-13.6-9.5-29.4-16.6-47.5-21.5C289.4 2.4 270.2 0 249.8 0c-44.2 0-87.3 9.5-128.5 31.7v85.4c33.9-30.7 70-49.8 107.4-49.8c11.6 0 22.4 1.3 32.6 3.8c10.2 2.5 19.2 6.4 27.1 11.6c7.9 5.2 14.2 11.7 18.7 19.4s6.2 15.2 6.2 25.6c0 11.6-2.5 23.8-8.7 33.6c-6.2 9.7-13.8 19-23 27.7c-9.2 8.8-19 17.8-29.6 27.1c-10.8 9.3-20.8 18.9-30 28.8c-9.2 9.9-16.8 20.6-23 32.1s-9.2 24.2-9.2 38.3c0 9 .9 18.2 2.6 27.4c1.8 9.2 1.2 14.1 3.5 19.8h85.4c-7.5-15.1-10.7-24.8-10.7-42.7c0-11.8 4.1-33 10.7-42.7c6.5-9.7 8.9-10.8 18.8-19.8s20.7-18 32.2-27.1c11.3-9 21.9-19.3 31.8-31c9.9-11.7 18.2-24.5 24.7-38.3c6.5-13.8 9.8-26.1 9.8-43.9c.2-18.6-3.8-38.7-11.5-52.8M238.7 405.3c-29.5 0-53.4 23.9-53.4 53.4s23.9 53.4 53.4 53.4s53.4-23.9 53.4-53.4c-.1-29.5-24-53.4-53.4-53.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C138.2 0 42.7 95.5 42.7 213.3V512h128V213.3c0-47.1 38.2-85.3 85.3-85.3s85.3 38.2 85.3 85.3V512h128V213.3C469.4 95.5 373.8 0 256 0M128 469.3H85.4V384H128zm298.7 0H384V384h42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405.3 149.3C405.3 66.9 338.5 0 256 0c-82.5 0-149.3 66.9-149.3 149.3c0 51.7 26.3 97.3 66.3 124.1L106.7 512h298.7l-66.3-238.6c39.9-26.8 66.2-72.3 66.2-124.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tep(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469.3 77H42.7C19.1 77 0 96.1 0 119.7v298.7C0 441.9 19.1 461 42.7 461h426.7c23.6 0 42.7-19.1 42.7-42.7V119.7C512 96.1 492.9 77 469.3 77M384 354.3H128c-47.1 0-85.3-38.2-85.3-85.3s38.2-85.3 85.3-85.3s85.3 38.2 85.3 85.3c0 15.6-4.5 30.1-11.8 42.7h109c-7.3-12.6-11.8-27-11.8-42.7c0-47.1 38.2-85.3 85.3-85.3s85.3 38.2 85.3 85.3s-38.2 85.3-85.3 85.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tick(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M437.3 30L202.7 339.3L64 200.7l-64 64L213.3 478L512 94z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 469.3c-117.8 0-213.3-95.5-213.3-213.3c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 117.8-95.5 213.3-213.3 213.3m21.3-256v-128h-42.7V320l128-128l-32-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeFive(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 469.3c-117.8 0-213.3-95.5-213.3-213.3c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 117.8-95.5 213.3-213.3 213.3m-21.3-256L181.3 160l-32 32l128 128V85.3h-42.7v128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 469.3c-117.8 0-213.3-95.5-213.3-213.3c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 117.8-95.5 213.3-213.3 213.3m-21.3-234.6H128v42.7h149.3v-192h-42.7v149.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 469.3c-117.8 0-213.3-95.5-213.3-213.3c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 117.8-95.5 213.3-213.3 213.3m21.3-384h-42.7v192H384v-42.7H277.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 469.3c-117.8 0-213.3-95.5-213.3-213.3c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 117.8-95.5 213.3-213.3 213.3m-21.3-234.6L149.3 320l32 32l96-96V85.3h-42.7v149.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 469.3c-117.8 0-213.3-95.5-213.3-213.3c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 117.8-95.5 213.3-213.3 213.3m21.3-384h-42.7V256l96 96l32-32l-85.3-85.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Title(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M46.5 0v139.6h23.3c0-23.3 0-69.8 23.3-93.1c23.2-23.3 46.5-23.3 69.8-23.3h46.5v395.6c0 34.9-11.6 69.8-46.5 69.8h-22.8l-.5 23.2h232.7v-23.3H349c-34.9 0-46.5-34.9-46.5-69.8V23.3H349c23.3 0 46.5 0 69.8 23.3s23.3 69.8 23.3 93.1h23.3V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolBox(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 149.3V448c0 23.5 19.1 42.7 42.7 42.7H64v-384H42.7C19.1 106.7 0 125.8 0 149.3m469.3-42.6H448v384h21.3c23.5 0 42.7-19.1 42.7-42.7V149.3c0-23.5-19.1-42.6-42.7-42.6m-362.6 384h298.7v-384H106.7zM213.3 64h85.3v21.3h42.7V64c0-23.5-19.1-42.7-42.7-42.7h-85.3c-23.5 0-42.7 19.1-42.7 42.7v21.3h42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolBoxOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469.3 85.3H42.7C19.1 85.3 0 104.5 0 128v21.3h512V128c0-23.5-19.1-42.7-42.7-42.7m-256-42.6h85.3V64h42.7V42.7c0-23.5-19.1-42.7-42.7-42.7h-85.3c-23.5 0-42.7 19.1-42.7 42.7V64h42.7zM0 426.7c0 23.5 19.1 42.7 42.7 42.7h426.7c23.5 0 42.7-19.1 42.7-42.7v-21.3H0zm0-64h512V192H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M374.2 189.2H157.5L275.7 71.1H157.5L0 228.6l157.5 157.5h118.2L157.5 268h216.6c32.6 0 59.1 26.4 59.1 59.1c0 32.6-26.4 59.1-59.1 59.1V465c76.1 0 137.8-61.7 137.8-137.8c.1-76.3-61.6-138-137.7-138"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UndoIcon(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 1C179.8 1 111.7 34.4 64.9 87.2L0 22.3V193h170.7l-60.2-60.2C145.6 91.5 197.5 65 256 65c106.1 0 192 85.9 192 192s-85.9 192-192 192c-53 0-101-21.5-135.8-56.2L75 438c46.4 46.3 110.4 75 181 75c141.4 0 256-114.6 256-256S397.4 1 256 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UndoOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M250.6 66.8V0L136.7 111.3l113.9 111.3v-66.8c73.7 0 133.6 59.8 133.6 133.6S324.4 423 250.6 423C176.9 423 117 363.2 117 289.4H28C28 412.3 127.7 512 250.6 512c123 0 222.6-99.7 222.6-222.6c0-123-99.6-222.6-222.6-222.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlike(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M109.7 28.5c-36.6 0-54.9 18.3-54.9 36.6s18.3 36.6 36.6 36.6c-36.6 0-54.9 18.3-54.9 36.6s18.3 36.6 36.6 36.6c-36.6 0-54.9 18.3-54.9 36.6S36.6 248 54.9 248C18.3 248 0 266.2 0 284.5s18.3 36.6 36.6 36.6h146.3c-9.1 36.6-18.3 73.1-18.3 91.4c0 36.6 18.3 73.1 27.4 82.3c.2.2 9.1 9.1 27.4 9.1c27.4 0 27.4-27.4 27.4-27.4c0-.5 9.1-45.7 9.1-64c0-18.3 36.6-91.4 54.9-109.7l18.3-9.1v-256l-18.3-9.1H109.7zM363.9 12l-16.5 34.8v237.7l18.3 18.3h143.8c1.5-12 2.5-24.2 2.5-36.6C512 157.3 452.3 62.4 363.9 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M412.4 232.7h-256V69.8c0-12.9 10.4-23.3 23.3-23.3h139.6c12.8 0 23.3 10.4 23.3 23.3v93.1h46.5V69.8c0-38.6-31.3-69.8-69.8-69.8H179.7c-38.6 0-69.8 31.3-69.8 69.8v162.9H86.6c-12.8 0-23.3 10.4-23.3 23.3v232.7c0 12.9 10.4 23.3 23.3 23.3h325.8c12.8 0 23.3-10.4 23.3-23.3V256c0-12.9-10.5-23.3-23.3-23.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M418.9 232.7h-256v-93.1c0-51.4 41.7-93.1 93.1-93.1s93.1 41.7 93.1 93.1v23.3h46.5v-23.3C395.6 62.5 333.1 0 256 0S116.4 62.5 116.4 139.6v93.1H93.1c-12.8 0-23.3 10.4-23.3 23.3v232.7c0 12.9 10.4 23.3 23.3 23.3h325.8c12.8 0 23.3-10.4 23.3-23.3V256c0-12.9-10.4-23.3-23.3-23.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Up(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 512h139.6V372.4H0zm0-186.2h139.6V186.2H0zm46.5-93.1H93v46.5H46.5zM0 139.6h139.6V0H0zm46.5-93.1H93V93H46.5zm314.2 35L442.2 0h-256v256l81.5-81.5C453.8 290.9 372.4 418.9 209.5 512c162.9 0 395.6-186.2 151.2-430.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpArrow(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.5 0L34.2 213.3v128l170.6-170.6V512h85.4V170.7l170.6 170.6v-128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpArrowOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M245 0L74.3 213.3h128V512h85.4V213.3h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 63.6L0 319.6l69.8 69.8L256 203.2l186.2 186.2l69.8-69.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadFour(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M69.8 325.8h139.6V512h93.1V325.8h139.6L256 93.1zM0 0v46.5h512V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 4.1C114.6 4.1 0 118.7 0 260.1c0 121 84.1 222.2 196.9 248.9V279.8H98.5L256 102.6l157.5 177.2H315V509c112.9-26.7 197-127.9 197-248.9c0-141.4-114.6-256-256-256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadThree(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M209.5 418.9h93.1V232.7h139.6L256 0L69.8 232.7h139.6v186.2zM0 465.5V512h512v-46.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadTwo(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M210.5 419.9h93.1V233.7h139.6L257 1L70.8 233.7h139.6v186.2zm256-46.5v93.1h-419v-93.1H1V513h512V373.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usd(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M393 315.6c-6.3-12.5-14.8-23.9-25.5-34.2c-10.7-10.3-23.1-19.6-37.3-28.1c-14.1-8.5-29.2-16.6-45.1-24.3c-3.8-1.6-7.2-3-10.8-4.5v-129c2.7-.1 5.5-.4 8.2-.4c40.8 0 75.6 15.5 101.5 32.8v-73c-21.8-10.1-59.2-18.3-107-18.3c-.9 0-1.7.1-2.6.1V0h-36.6v39.8c-7.1 1.1-14.3 2.1-21.5 3.7c-20.2 4.6-38.2 11.8-54 21.3c-15.8 9.6-28.6 21.7-38.3 36.4c-9.7 14.7-14.5 32-14.5 52c0 16.2 2.8 30.3 8.4 42.3c5.6 12.1 13.4 22.9 23.3 32.4c9.9 9.6 21.5 18.3 34.9 26.2c13.4 7.9 27.8 15.7 43.3 23.4c6.3 2.7 12.4 5.3 18.3 7.9v131.2c-3.3.1-6.2.4-9.7.4c-9.1 0-19.1-.9-29.9-2.7c-10.8-1.8-21.7-4.3-32.5-7.6c-10.8-3.2-21.1-7.3-30.8-12s-18-10.1-25-16v74.6c6.4 3.8 14.9 5.1 25.3 8c10.4 3 21.4 5.5 33.1 7.6c11.7 2.1 23.1 3.6 34.2 4.7c11.1 1.1 25.1 1.6 32.7 1.6c.9 0 1.7-.1 2.6-.1V512h36.6v-39.2c7-1 14.1-2.1 21.6-3.6c20.6-4.2 38.8-10.8 54.7-20.1c15.8-9.3 28.4-21.3 37.8-36.1c9.4-14.8 14-33 14-54.5c0-16-3.1-30.3-9.4-42.9M214.3 194.7c-8.2-5.7-14.4-12-18.6-18.9c-4.3-6.9-6.4-15.1-6.4-24.6c0-10.1 2.8-18.6 8.3-25.7c5.5-7.1 12.7-12.9 21.5-17.5c5.7-2.9 12.1-5.1 18.8-7v107.3c-9-4.5-17.1-9-23.6-13.6m85.1 207.6c-6.7 4.2-15.2 7.4-24.9 9.8v-110c6.8 3.5 13.1 7 18.7 10.7c9.3 6.1 16.6 12.9 21.8 20.3c5.2 7.4 7.8 16.1 7.8 26.2c0 18.8-7.8 33.1-23.4 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M302.5 162.9h-256C20.9 162.9 0 183.8 0 209.5v256C0 491.1 20.9 512 46.5 512h256c25.7 0 46.5-20.9 46.5-46.5v-256c.1-25.7-20.8-46.6-46.5-46.6m69.9 116.4v116.4L512 465.5v-256zM69.8 139.6c38.6 0 69.8-31.3 69.8-69.8S108.4 0 69.8 0C31.3 0 0 31.3 0 69.8s31.3 69.8 69.8 69.8M279.3 0c-38.6 0-69.8 31.3-69.8 69.8s31.3 69.8 69.8 69.8c38.6 0 69.8-31.3 69.8-69.8S317.8 0 279.3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M395.6 69.8L325.8 0h-58.2l69.8 69.8zM23.3 0H0v69.8h93.1zm221.1 69.8L174.5 0h-58.2l69.8 69.8zm174.5 93.1h-93.1l69.8-69.8h-58.2l-69.8 69.8h-93.1l69.8-69.8h-58.2l-69.8 69.8h-93l69.8-69.8H0v372.4C0 491.1 20.9 512 46.5 512h418.9c25.7 0 46.5-20.9 46.5-46.5V93.1h-23.3zM186.2 442.2V232.7l186.2 104.7zM418.9 0l69.8 69.8H512V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Webcam(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 426.7c-58.9 0-112.6-21.9-153.8-57.7L56 512h384l-38.2-143.1c-41.1 35.9-94.9 57.8-153.8 57.8m0-149.4c47.1 0 85.3-38.2 85.3-85.3s-38.2-85.3-85.3-85.3s-85.3 38.2-85.3 85.3s38.2 85.3 85.3 85.3m-21.3-128c11.8 0 21.3 9.5 21.3 21.3s-9.6 21.3-21.3 21.3c-11.8 0-21.3-9.5-21.3-21.3c0-11.7 9.5-21.3 21.3-21.3M248 384c106.1 0 192-86 192-192S354.1 0 248 0S56 86 56 192s86 192 192 192m0-320c70.7 0 128 57.3 128 128s-57.3 128-128 128s-128-57.3-128-128S177.3 64 248 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func World(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M236.3 65.6c1.2 0 2.7.6 4 2.4c-5.8-8.3 6.5-4.7 7.2-12.6c7-5.2 4.5-8.6 6.5-8.8l2.2-1.8c-4.3 1.6-5.8 1.9-7 1.9h-2c-1.5 0-4 .3-9.9 2.1c-13.5 1-23.2 15.6-27.8 16c-4.7 2.5-4.5 3-3.4 3c.3 0 .7-.1 1.1-.1c.4 0 .7-.1 1-.1c1.1 0 .8.5-5.6 3.5c1.5.4-11.2 6-11.3 14.8c-.4 3.4 1 5.7 3.3 5.7c2 0 4.5-1.5 7.2-5.3c9-10.4 28-11.4 36.2-17.7c-5.1-.4-4.1-3-1.7-3m-45.7 3.5c-.9 0 .3-1.1 1.6-2.2c1.4-1.1 3-2.2 3-2.2l-.1.1c5.4-2.1 2.1-2.6 5-4h-.2c-.6 0 .5-.8 1.5-1.7c1-.8 1.9-1.7 1-1.7c-.5 0-1.8.3-4 1.2c.7-1 3.4-1.7 7.2-3.8c-5.5 1.8-9.8 2.9-10.3 2.9c-.4 0 1.9-.8 8.8-2.9c.3 0 .6.1.8.1c.7 0 1-.2.9-.4c-.1-.2-.5-.3-1.2-.3c-9.9.8-17.3 7-15.1 7.8c1.3-.4 2.2-.6 2.8-.6c1.3 0 .7 1.1-1.9 2.7c1.2 2.8-20 10-22.4 12.6c1.2-.7 2-.9 2.4-.9c1.6 0-.5 3.1-2.6 4.6c1.3 1.2 2.3 1.6 3.1 1.6c2.1 0 3-2.9 4.8-3.8c.4.8.8 1.2 1.3 1.2c1.7 0 3.8-4.2 5.5-6.4c.5.5 1.1.9 1.9.9c1.6 0 4-1.3 7.8-5c-.8.1-1.4.2-1.6.2m-23 18.4c-3.6 2.2-5.5 3.1-6.3 3.1c-1.5 0 .5-2.8 2.3-5.7c1.8-2.8 3.4-5.6 1.1-5.6c-1.1 0-3.1.6-6.3 2.2c-8.7 5.8-23.4 21.8-29.9 21.8h-.6c10.5-7.7 7.4-22.8 22.1-26.7c14.5-8.3 24.4-5.7 37.3-15.2c-3.8 1.9-9.4 4-10.7 4c-.8 0 .4-1 5.6-3.5c-.5.1-.9.2-1.2.2c-3.5 0 23-9.5 24.3-9.7c-10.9 1.5-28.5 10.6-30.9 10.6c-.3 0-.3-.1-.3-.3c.4-.3.4-.5 0-.5c-1.3 0-6.7 1.8-12.2 3.6c-5.4 1.8-10.9 3.7-12.1 3.7c-.4 0-.4-.2.2-.6c-47.2 26.8-86.6 72-98.9 125.6c5.1 11.7 1.6 34 10 40.5c9.6 8.1-8.2 31.6-3.5 46.1c4.8 26.2 25.5 44.1 27.7 70.9c3.8 18.4 17.7 40.9 23.5 52.6c4.4 4.4 16.5 17.1 18.6 17.1c.8 0 .3-1.9-2.9-7c-2-5.7-13.6-20.6-8-20.6c.3 0 .6 0 1 .1c-6.9-7 15.3-5.1-.2-14.6c-1.8-2.2-2.2-2.9-1.8-2.9c.5 0 1.9.8 3.5 1.6c1.7.8 3.7 1.6 5.3 1.6c2.5 0 4.2-1.7 3.2-7.9c.4.2.8.2 1.1.2c2.6 0 1.7-5.1 1.9-10.2c.2-5.1 1.5-10.2 8.5-10.2c.8 0 1.7.1 2.6.2c16.2-8.8 2.1-33.1 20.7-42c-.7-20.2-27.8-21.2-38-29.4c-1.7 1-3.3 1.3-4.9 1.3c-1.8 0-3.4-.4-4.8-.9c-1.4-.4-2.7-.9-3.6-.9c-.2 0-.4 0-.6.1c17.4-4.6 4.7-28.8-9.7-28.8h-1c-1.1-10.5-6.5-6.5-7.1-15.2c-.9.5-1.8.7-2.7.7c-2.4 0-4.6-1.7-6.5-3.3c-1.9-1.7-3.5-3.3-4.9-3.3c-1.1 0-2.1 1.1-3 4.3c1-4.9.7-6.5-.3-6.5s-2.8 1.7-4.8 3.4c-1.9 1.7-4.1 3.4-5.8 3.4h-.5c-11.7-8.3.1-23-6.1-34.3c4-4.3 5.6-11.6 3.8-11.6c-.9 0-2.7 1.8-5.5 6.7c-4-10.2 6.6-33.3 15.6-34.6c.3 0 .5-.1.8-.1c1.6 0 3.2.8 4.6 2.5c.8 4.7-1.2 14.6-.4 14.6c.4 0 1.2-2 3.2-7.2c2.6-11.9 20.8-20.5 22.7-27.8c.1.1.1.1.2.1c2.1 0 15.6-11.9 21.5-13.2c2.3-2.3 3.9-3.1 5.1-3.1c1.3 0 2 1 2.5 1.9c.5 1 .9 1.9 1.4 1.9c.4 0 .9-.6 1.8-2.3c-3.1-6.1 4.9-11.9 2.7-11.9c-.8 0-3.2.8-8.1 2.8c-.8.4-1.2.5-1.4.5c-.6 0 1.8-1.6 5.3-3.2s7.9-3.2 11.6-3.2c1.8 0 3.4.4 4.6 1.4c15.5-3.7 7-8.8 7.8-8.8h.2c4.4-2.6 3.1-12.4 9.6-18.5m-69.3 40.9c-.2 0 .7-1.3 3.7-5c6.2-4.6 11.2-7.7 5.5-10.7c2.5-1.1 4.9-2.4 7.2-3.8c-.1 4.9-4.2 16.8-6.4 16.8c-.5 0-1-.7-1.2-2.2c1.3-2.3 1.4-3.2.9-3.2c-.8 0-3.1 2-5.3 4.1c-2.2 1.9-4.2 4-4.4 4m5.7 2.7c-2.7 0-2.2-1.7 6.1-3.4c3.6-.5.8-.3 5.1-1.3c-4.2 3.4-8.9 4.7-11.2 4.7M259.5 121c1.2-.2 2.1-.2 2.8-.2c2.9 0 1.6 1.2-.1 2.4c-1.7 1.2-3.8 2.4-2.6 2.4c.7 0 2.6-.4 6.5-1.5c19.7-1.4-7.7-18.1-3.3-23.7l-1.2-.4c-8.1 10.3 5.4 10.5-2.1 21m-8.2.4c1.4 0 3.2-1.6 5.1-6.1c2.6-3.1 1.5-3.5-1.3-5c-7.7 2-7 11.1-3.8 11.1m69 31.7c-.3 0-.5.1-.7.2c.8.2 1.6.4 2.4.4c-.8-.4-1.3-.6-1.7-.6M454.6 178c1.7 3.7 2.6 5.2 2.7 5.2c.7 0-12.6-32.4-18.1-38.2c-28.7-47.2-75.9-83.3-129.7-96.8h-.1c-1.4 0 .2.7 2.1 1.4c1.9.7 4.2 1.4 4.2 1.4s-1.5-.5-6.1-1.8c-5.1-1.5-10.3-2.6-15.6-3c-.9.3 21.9 9.3 37.2 14c-5.1-1.6-10.4-3-11.7-3c-.9 0 0 .6 3.7 2.2c-4.5-1.3-6.4-1.9-6.6-1.9c-.3 0 6 2.1 11.9 4.2c6 2.1 11.7 4.2 10.1 4.2c-.4 0-1.5-.2-3.2-.5c5.1 2.8 6.9 3.9 6.6 3.9c-.5 0-6.8-3-13.3-6c-6.4-3-12.9-6-13.8-6c-.5 0 .8 1 5.2 3.6c6 2.6 12 5.2 10.3 5.2c-.8 0-3.1-.5-7.6-1.8c10 5.1-7.1 3.1-1.3 10c-3.2-3-4.5-4-4.8-4s.5 1.2 1.3 2.4s1.4 2.4.7 2.4h-.2c7 7.6-4.9.8 3.6 7.6c-6.4-2.9-17.3-7.7-12.4-7.7c1.2 0 3.5.3 7.1 1.1c-6.5-7.6-24.9-7.9-25.3-8.8c-4.6.3-4.7 3.8-3.9 10.2c-.1 10.1-10 10.3-8.7 16.1c0-.1.1-.1.1-.1c.2 0 .3 2 .9 4.1c.7 2 1.9 4.1 4.5 4.1c1.4 0 3.2-.6 5.4-2.1c3.1 5.6 5.9 7.6 8.1 7.6c3.3 0 5.3-4.5 5.1-8.6c-3.6-2-8.3-18.6-3.1-18.6c.8 0 1.8.4 3.1 1.3c-8.1 19.9 32 6.6 9.3 15.7c5.1 6.7-3.2 6.4-.8 14.5c-3.9 1.1-8.2 2.1-12 2.1c-5.6 0-9.9-2.4-9.4-10.8c-7 3.7 7 16.2-6.9 16.2h-.8c-6.7 12.3-31.8 8.1-12.1 21.1c1.1 5.7-2.5 6.7-7.1 6.7c-1.2 0-2.5-.1-3.8-.2c-1.3-.1-2.6-.1-3.8-.1c-5.2 0-8.9 1.3-5.7 9.4c-2.3 8.4 3.1 12.1 9.7 12.1c8 0 17.8-5.2 18.5-13.8c4.3-6.5 9.8-9.2 15.5-9.2c9 0 18.5 6.8 24.7 15.2c.5.8.9 1.1 1.1 1.1c.5 0 0-2-.2-4c-.2-1.6-.2-3.2.4-3.7c-8.9-2.5-15.7-13.3-12.7-13.3c1 0 3.1 1.2 6.5 4.4h.5c5.6 0 10 4.4 13.4 8.8c3.5 4.4 6.1 8.8 8.4 8.8c.9 0 1.8-.8 2.6-2.6c-.4-.6-1.4-2.3-1.9-4c-.6-2.1-.6-4.2 2.1-4.2c1.1 0 2.6.4 4.7 1.2c6.8-1.8-4.3-19.5 4.7-19.5c.9 0 2 .2 3.4.6c-.1 3.1 1.3 4.2 2.6 4.2c2.6 0 5.2-3.7-1.1-4.8c2.4-2.3 3.8-3.2 4.8-3.2c3.5 0 .7 11.6 12.8 11.6c4.1 9.7-35.6.7-22.9 19.3c2.9 2.1 6.5 2.5 10 2.5c1.1 0 2.1 0 3.1-.1c1 0 2-.1 2.9-.1c6.2 0 10.6 1.5 7.9 14.3c-3.3 2.4-7 3.2-10.8 3.2c-3.4 0-6.8-.6-10.1-1.2s-6.5-1.2-9.4-1.2c-4.5 0-8.1 1.5-10.3 6.7c-11.7-4.8-27.1-5-25.6-19.5c-20.4 2.9-41.7 1.4-56.7 13.9c-.9 18.3-32.3 24.2-24.7 45c-9.1 19.2 11 46.1 31.5 46.1c.7 0 1.5 0 2.2-.1c11.2-.1 22.5-5.1 32.4-5.1c5.1 0 9.7 1.3 13.7 5.3c.9-.1 1.7-.2 2.4-.2c15.3 0-4.9 24.3 13.4 28.4c16.6 17.2-10.2 34.9 2.9 51.4c1.4 12.9 7.8 23.7 7.8 36.3c1.7.2 3.3.3 5 .3c21.1 0 34.3-18.4 46.7-32.3c-3.8-20.8 29.5-24.9 18.1-48.2c-5.3-25.5 25.1-40.6 23.4-66.1c.3-3.6-.7-4.7-2.2-4.7c-1.5 0-3.7 1.1-5.8 2.1c-2.2 1.1-4.6 2.1-6.5 2.1c-2.5 0-4.4-1.7-4.6-7.3c-14.7-13.7-21.4-32.3-35.1-47.4c17.8 8.6 27.9 28 36.1 45.6c.7.2 1.4.3 2.2.3c14.5 0 34.7-34.7 12.7-43.1c-1.2 3.3-2.8 4.5-4.7 4.5c-3.3 0-7.1-4-9.8-7.9c-2.7-4-4.1-8-2.6-8c1 0 3.2 1.6 7 5.8c4.3 2.7 8.5 3.4 12.9 3.4c4.3 0 8.7-.7 13.3-1.1c5.8 2.3 8.3 9.7 10.1 9.7c.3 0 .7-.3 1-1c4.6 11 9.9 30.6 14.7 35.9c-2.5-16.8-3.1-34.6-7.4-51.1m-55.8-23.1c-11.3-2.9-15-16.9-23.9-25c-.1-2.2 3.2-1.8 2.4-5.5c12 3.7 8.1 7.4 7 8.1c4.4 7.3 16.7 12.3 14.5 22.4M243.5 78.8c-10.9 3.2-12.4-.3-10.5 6.4c.8.6 2.3.8 3.8.8c5.3.1 12.5-3 6.7-7.2m160.1 275.5c14.3-12.1 16.1-26.5 16.8-41.1c-6.4 12.5-25.4 26.3-16.8 41.1m-65.1-195c-.1.1-.1.1-.1.2c.1.2.2.2.2.2s0-.1-.1-.4M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0m0 492.3C125.5 492.3 19.7 386.5 19.7 256S125.5 19.7 256 19.7S492.3 125.5 492.3 256S386.5 492.3 256 492.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WorldOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M93.8 114.6c-4.7 1.1-1.7.9-5.6 1.4c-16.7 3.6-4.3 6.8 5.6-1.4m293.7 6.7c1.2-.8 5.4-4.9-7.7-8.9c.8 4.1-2.7 3.7-2.7 6c9.7 8.8 13.7 24.1 26.1 27.3c2.4-11-11-16.4-15.7-24.4m-302.6-9.9c1.5 8.9 8.2-9.4 8.3-15.9c-2.6 1.5-5.2 3-7.9 4.2c6.3 3.2.8 6.6-6 11.7c-13.8 17.2 12.9-13.4 5.6 0M256 0C114.6 0 0 114.6 0 256c0 141.3 114.6 256 256 256s256-114.7 256-256C512 114.6 397.4 0 256 0m6.8 85.8l1.2.4c-4.8 6.2 25 24.3 3.6 25.8c-20 5.7 8.4-5.2-7.1-3.3c8.2-11.4-6.5-11.6 2.3-22.9m-121.4 16.4c-7.2-6-29.8 8.2-21.9 4.8c19.6-7.7 1.3.8 5.9 10c-4.2 8.7-1.4-8.6-11.8 1.7c-7.5 1.7-25.9 18.7-23.6 13.5c-.6 8.1-21.9 17.7-24.8 31.2c-7 18.7-1.7-.7-3-8c-10-12.7-28.2 21.5-22.8 35c9.1-16 8.4-1.7 1.8 5.4c6.7 12.3-6.1 28.3 6.6 37.4c5.6 1.3 16.8-18.8 11.9 2.1c3.4-18.1 9.4 4.3 19.1-.7c.6 9.5 6.5 5.1 7.8 16.6c16.2-1.2 31 26.2 11.7 31.4c2.9-.8 8.6 4.3 15.2.4c11.2 8.9 40.7 10 41.5 32c-20.3 9.7-5 36.3-22.6 45.8c-20.2-3-6.9 24.9-15.4 21.7c3.4 20.1-20.4-2.6-11.2 8.5c16.9 10.4-7.4 8.3.2 15.9c-8.5-1.8 5.3 15.8 7.6 22.3c12.2 19.8-10.5-4.4-17.2-11c-6.4-12.8-21.5-37.3-25.7-57.4c-2.4-29.2-25-48.8-30.2-77.3c-5.2-15.9 14.3-41.4 3.8-50.3c-9.1-7.1-5.4-31.4-10.8-44.2C47 130.5 89.9 81.2 141.4 52c-5.3 3.9 30.3-10.1 26.2-6.7c-1.1 2.5 20.8-9.5 34-11.3c-1.4.2-34.3 12-25.2 10.4c-14.1 6.9-1.4 3 5.6-.5c-14 10.3-24.8 7.4-40.7 16.5c-16 4.2-12.7 20.8-24.1 29.1c6.7 1.2 23.5-17.3 33.3-23.8c22.5-10.9-11.4 19.8 10 6.6c-7.2 6.7-5.7 17.4-10.1 20.4c-2.2-.6 8.7 5.2-9 9.5m35-46c-2.3 3.1-5.5 9.8-7.4 5.7c-2.6 1.3-3.6 6.9-8.5 2.4c2.9-2.1 5.9-7.1.2-4c2.6-2.8 25.8-10.7 24.5-13.7c4.1-2.6 3.7-3.9-1-2.3c-2.4-.8 5.7-7.6 16.5-8.5c1.5 0 2.1 1-.6.7c-16.3 5-9.3 3.6 1.7 0c-4.2 2.4-7.1 3.1-7.8 4.2c11-4-.6 2.9 1.9 2.4c-3.1 1.6.5 2.1-5.5 4.4c1.1-.9-9.8 6.5-3.3 4.3c-6.3 6-9.1 6.1-10.7 4.4m9.6 14.3c.2-9.6 14-15.7 12.3-16.2c17-8-5.9.3 7.5-6.9c5-.5 15.6-16.5 30.3-17.5c16.2-4.9 8.7.3 20.7-4.3l-2.4 2c-2.1.3.5 4-7.1 9.6c-.8 8.7-14.5 4.7-7.7 14c-4.4-6.3-11-.2-2.7.4c-8.9 6.8-29.6 8-39.5 19.3c-6.4 9.2-12.3 6.3-11.4-.4m71.1 32c-6.8 16.4-13.4-2.4-1.4-5.4c3 1.6 4.2 2.1 1.4 5.4m-25.6-32.8c-2-7.4-.4-3.5 11.5-7c8.2 5.9-7.3 9.8-11.5 7m186.2 293.5c-9.4-16.2 11.4-31.2 18.4-44.8c-.9 15.9-2.9 31.6-18.4 44.8m35.4-185.1c-10.2.8-19.4 3.2-28.6-2.6c-21.2-23.2 3.9 26.2 10.9 6c25.2 9.6-.4 51-16.3 46.7c-8.9-19.2-19.9-40.3-39.3-49.7c14.9 16.5 22.3 36.8 38.3 51.7c1.1 20.8 22.2-7.6 20.9 8.5c2 27.7-31.3 44.3-25.5 72.1c12.4 25.3-23.9 29.9-19.8 52.6c-14.6 16.3-30.2 38.3-56.4 34.8c0-13.8-7-25.5-8.6-39.7c-14.2-18 15-37.3-3.1-56.1c-20.9-4.7 4.3-33.5-17.2-30.8c-12.9-12.9-31.8-.4-50.3-.3c-23.2 2.2-47.1-28.5-36.8-50.2c-8.2-22.6 26-29.2 26.9-49.1c16.4-13.7 39.7-12 61.9-15.2c-1.6 15.9 15.2 16 27.9 21.3c7.1-17.2 29.2 2.8 44.3-8.1c5.2-25.4-14.7-10.1-26.1-18.2c-13.8-20.2 29.5-10.4 25-21c-16.8-.1-7.3-20.7-19.2-9.2c10.7 1.9-1.9 10.3-1.6.7c-16.2-4.7-.6 18.4-8.8 20.6c-12.5-5.2-6.6 5.9-5.3 7.6c-5.4 11.7-12-17.2-27.3-16.4c-15.2-13.9-6 6.3 7.2 9.6c-2.8.8 1.6 12.3-1.9 7.4c-10.9-15-31.6-25-43.9-6.6c-1.3 17.2-36.3 22.1-30.7 2c-8.2-20.8 25.4-.6 22.3-17.2c-21.6-14.3 5.9-9.7 13.2-23.1c16.6.5.7-13.6 8.5-17.7c-.8 15.3 12.7 12.4 23.4 9.5c-2.6-8.8 6.4-8.5.9-15.8c24.8-9.9-18.9 4.6-10.1-17.1c-10.7-7.4-4.5 16.3 0 18.8c.3 7.3-5.9 16.3-14.4 1c-12.4 8.1-11.1-8.2-11.9-6.5c-1.4-6.3 9.4-6.6 9.5-17.6c-.9-7-.7-10.7 4.3-11.1c.4 1 20.5 1.3 27.6 9.6c-19.4-3.9-2.9 3.2 5.8 7.2c-9.3-7.3 3.7 0-3.9-8.3c3 .6-8.3-11.4 3.3-.9c-6.3-7.5 12.3-5.3 1.3-10.9c16.1 4.5 6.6.4-2.9-3.7c-26.2-15.6 46.3 21.1 16.7 4.8c18.9 4.1-40.4-14.6-13.4-6.4c-10.3-4.5-.3-2 9 .9c-16.7-5.2-41.7-14.9-40.7-15.3c5.8.4 11.5 1.7 17 3.3c17.1 5.1-4.9-1.2-.2-1.1c58.8 15.1 110.3 54.5 141.6 106c7.3 7.7 27.2 58.6 16.8 36c4.7 18 5.4 37.4 7.9 55.8c-5.2-5.8-11-27.2-16-39.1c-2.1 4.6-4.5-6.5-12.2-9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Write(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 439.5h170.7V268.8H0zm42.7-128H128v85.3H42.7zm170.6 106.7H512v-42.7H213.3zM0 226.2h170.7V55.5H0zm42.7-128H128v85.3H42.7zm170.6-21.4v42.7H512V76.8zm0 256H512v-42.7H213.3zm0-128H512v-42.7H213.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WriteOne(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 242.9h170.7V72.3H0zM213.3 93.6v42.7H512V93.6zm0 128H512v-42.7H213.3zm0 128H512v-42.7H213.3zm0 85.3H512v-42.7H213.3zM0 456.3h170.7V285.6H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zip(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M413.4 0H114.7C91.1 0 72 19.1 72 42.7v426.7c0 23.5 19.1 42.7 42.7 42.7h298.7c23.5 0 42.7-19.1 42.7-42.7V42.7C456 19.1 436.9 0 413.4 0m-192 469.3L242.7 320h42.7l21.3 149.3zM328 128h-64v42.7h64v42.7h-64V256h64v42.7h-64V256h-64v-42.7h64v-42.7h-64V128h64V85.3h-64V42.7h64v42.7h64zm-74.6 277.3L242.7 448h42.7l-10.7-42.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomMinus(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.8 0C223 0 139.6 83.4 139.6 186.2c0 33.5 9 64.8 24.4 92L0 442.2l23.3 46.5L69.8 512l164-164c27.1 15.5 58.5 24.4 92 24.4C428.6 372.4 512 289 512 186.2S428.6 0 325.8 0m0 314.2c-70.7 0-128-57.3-128-128s57.3-128 128-128s128 57.3 128 128s-57.3 128-128 128M256 209.5h139.6V163H256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomPlus(children ...ElementRenderer) *SubwayIcon {
	return &SubwayIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325.8 0C223 0 139.6 83.4 139.6 186.2c0 33.5 9 64.8 24.4 92L0 442.2l23.3 46.5L69.8 512l164-164c27.2 15.5 58.5 24.4 92 24.4C428.6 372.4 512 289 512 186.2S428.6 0 325.8 0m0 314.2c-70.7 0-128-57.3-128-128s57.3-128 128-128s128 57.3 128 128s-57.3 128-128 128m23.3-197.8h-46.5v46.5H256v46.5h46.5V256H349v-46.5h46.5V163H349v-46.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
