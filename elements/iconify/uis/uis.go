package uis

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "4.0.8"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type UisIcon struct {
	*SVGSVGElement
}

func Airplay(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.8 13.4c-.1-.1-.1-.2-.2-.2c-.5-.3-1.1-.2-1.5.2l-4 6c0 .2-.1.4-.1.6c0 .6.4 1 1 1h8c.2 0 .4-.1.6-.2c.5-.3.6-.9.3-1.4zM19 3H5C3.3 3 2 4.3 2 6v9c0 1.7 1.3 3 3 3h.8c.6 0 1-.4 1-1s-.4-1-1-1H5c-.6 0-1-.4-1-1V6c0-.6.4-1 1-1h14c.6 0 1 .4 1 1v9c0 .6-.4 1-1 1h-.8c-.6 0-1 .4-1 1s.4 1 1 1h.8c1.7 0 3-1.3 3-3V6c0-1.7-1.3-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 15H3c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1m0-4H3c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1m0-4H3c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1m4-2h7c.6 0 1-.4 1-1s-.4-1-1-1h-7c-.6 0-1 .4-1 1s.4 1 1 1m-4 14H3c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1m11-4h-7c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1m0-8h-7c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1M10 3H7c-.6 0-1 .4-1 1s.4 1 1 1h3c.6 0 1-.4 1-1s-.4-1-1-1m11 8h-7c-.6 0-1 .4-1 1s.4 1 1 1h7c.6 0 1-.4 1-1s-.4-1-1-1m-4 8h-3c-.6 0-1 .4-1 1s.4 1 1 1h3c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 9c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1zM3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m14 10H7c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1m4-4H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterJustify(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 19H7c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1M3 5h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18 10H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m0-4H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m0-4H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18 2H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m0 4H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m0 4H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m0 4h14c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18 2H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m-4 4H3c-.6 0-1 .4-1 1s.4 1 1 1h14c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeftJustify(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 15H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m-6 4H3c-.6 0-1 .4-1 1s.4 1 1 1h12c.6 0 1-.4 1-1s-.4-1-1-1M3 5h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18 2H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m0 4H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLetterRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1M10 4h11c.6 0 1-.4 1-1s-.4-1-1-1H10c-.6 0-1 .4-1 1s.4 1 1 1m11 12H11c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1m0-6H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m0-4H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18 2H7c-.6 0-1 .4-1 1s.4 1 1 1h14c.6 0 1-.4 1-1s-.4-1-1-1m0 4H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m0 4H7c-.6 0-1 .4-1 1s.4 1 1 1h14c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRightJustify(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18 14H11c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1m0-8H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m0 4H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m0-8H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Analysis(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.7 7.3c-.4-.4-1-.4-1.4 0L14 13.6L9.7 9.3C9.5 9.1 9.3 9 9 9c-.3 0-.5.1-.7.3l-6 6c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3L9 11.4l4.3 4.3c.1.1.2.2.3.2c.1.1.3.1.4.1c.2 0 .5-.1.6-.3h.1l7-7c.4-.4.4-1 0-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Analytics(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1s1-.4 1-1V3c0-.6-.4-1-1-1M5 12c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1s1-.4 1-1v-8c0-.6-.4-1-1-1m10-4c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1s1-.4 1-1V9c0-.6-.4-1-1-1m5 8c-.6 0-1 .4-1 1v4c0 .6.4 1 1 1s1-.4 1-1v-4c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 13h-2c-.6 0-1 .4-1 1s.4 1 1 1h.9c-.4 2.5-2.4 4.5-4.9 4.9V11h1c.6 0 1-.4 1-1s-.4-1-1-1h-1V7.8c1.2-.4 2-1.5 2-2.8c0-1.7-1.3-3-3-3S9 3.3 9 5c0 1.3.8 2.4 2 2.8V9h-1c-.6 0-1 .4-1 1s.4 1 1 1h1v8.9c-2.5-.4-4.5-2.4-4.9-4.9H7c.6 0 1-.4 1-1s-.4-1-1-1H5c-.6 0-1 .4-1 1c0 4.4 3.6 8 8 8s8-3.6 8-8c0-.6-.4-1-1-1m-7-9c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleDown(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.3 11.5c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L12 9.3L9.7 7c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4zm3 1L12 14.8l-2.3-2.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l3 3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4c-.4-.3-1-.3-1.4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 14.3L14.7 12L17 9.7c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-3 3c-.4.4-.4 1 0 1.4l3 3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.3-.4.3-1 0-1.4M9.2 12l2.3-2.3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-3 3c-.4.4-.4 1 0 1.4l3 3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 8.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4L9.3 12L7 14.3c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4zm8.5 3l-3-3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2.3 2.3l-2.3 2.3c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l3-3c.3-.4.3-1 0-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleUp(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.7 12.5c-.4-.4-1-.4-1.4 0l-3 3c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0l2.3-2.3l2.3 2.3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4zm-3-1L12 9.2l2.3 2.3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-3-3c-.4-.4-1-.4-1.4 0l-3 3c-.4.4-.4 1 0 1.4c.4.3 1 .3 1.4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDown(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.9 9.2c-.4-.4-1-.4-1.4 0L12 12.7L8.5 9.2c-.4-.4-1-.4-1.4 0s-.4 1 0 1.4l4.2 4.2c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l4.2-4.2c.4-.4.4-1 0-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.3 12l3.5-3.5c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-4.2 4.2c-.4.4-.4 1 0 1.4l4.2 4.2c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.9 17.2c-.6 0-1-.4-1-1c0-.3.1-.5.3-.7l3.5-3.5l-3.5-3.5c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l4.2 4.2c.4.4.4 1 0 1.4l-4.2 4.2c-.2.2-.5.3-.7.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleRightB(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 11.3L9.9 5.6c-.4-.4-1-.4-1.4 0s-.4 1 0 1.4l4.9 4.9l-4.9 4.9c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l5.7-5.7c.3-.2.3-.8-.1-1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleUp(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.9 13.4l-4.2-4.2c-.4-.4-1-.4-1.4 0l-4.2 4.2c-.4.4-.4 1 0 1.4s1 .4 1.4 0l3.5-3.5l3.5 3.5c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apps(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2H3c-.6 0-1 .4-1 1v7c0 .6.4 1 1 1h7c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1m0 11H3c-.6 0-1 .4-1 1v7c0 .6.4 1 1 1h7c.6 0 1-.4 1-1v-7c0-.6-.4-1-1-1M21 2h-7c-.6 0-1 .4-1 1v7c0 .6.4 1 1 1h7c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1m0 11h-7c-.6 0-1 .4-1 1v7c0 .6.4 1 1 1h7c.6 0 1-.4 1-1v-7c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDown(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m3.7 10.7l-3 3c-.1.1-.2.2-.3.2c-.2.1-.5.1-.8 0c-.1 0-.2-.1-.3-.2l-3-3c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l1.3 1.3V9c0-.6.4-1 1-1s1 .4 1 1v3.6l1.3-1.3c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m3 11h-3.6l1.3 1.3c.4.4.4 1 0 1.4c-.4.4-1 .4-1.4 0l-3-3c-.4-.4-.4-1 0-1.4l3-3c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4L11.4 11H15c.6 0 1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m3.7 10.7l-3 3c-.4.4-1 .4-1.4 0c-.4-.4-.4-1 0-1.4l1.3-1.3H9c-.6 0-1-.4-1-1s.4-1 1-1h3.6l-1.3-1.3c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l3 3c.4.4.4 1 0 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUp(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m3.7 10.7c-.4.4-1 .4-1.4 0L13 11.4V15c0 .6-.4 1-1 1s-1-.4-1-1v-3.6l-1.3 1.3c-.4.4-1 .4-1.4 0c-.4-.4-.4-1 0-1.4l3-3c.4-.4 1-.4 1.4 0l3 3c.4.4.4 1 0 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 16H9.4l8.3-8.3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L8 14.6V7c0-.6-.4-1-1-1s-1 .4-1 1v10c0 .1 0 .3.1.4c.1.2.3.4.5.5c.1.1.3.1.4.1h10c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6c-.6 0-1 .4-1 1v7.6L7.7 6.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l8.3 8.3H7c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.4 8H17c.6 0 1-.4 1-1s-.4-1-1-1H7c-.6 0-1 .4-1 1v10c0 .6.4 1 1 1s1-.4 1-1V9.4l8.3 8.3c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6H7c-.6 0-1 .4-1 1s.4 1 1 1h7.6l-8.3 8.3c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0L16 9.4V17c0 .6.4 1 1 1s1-.4 1-1V7c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10c1.8 0 3.5-.5 5-1.3c.5-.3.6-.9.4-1.4c-.3-.5-.9-.6-1.4-.4c-3.8 2.2-8.7.9-10.9-2.9C2.9 12.2 4.2 7.3 8 5.1c3.8-2.2 8.7-.9 10.9 2.9c.7 1.2 1.1 2.6 1.1 4v.8c0 1-.8 1.8-1.8 1.8s-1.8-.8-1.8-1.8V8.5c0-.6-.4-1-1-1c-.5 0-.9.3-1 .8c-2-1.4-4.9-.9-6.3 1.1c-1.4 2-.9 4.9 1.1 6.3c1.9 1.3 4.4 1 5.9-.7c1.3 1.6 3.6 1.9 5.2.7c.9-.7 1.5-1.8 1.5-3V12C22 6.5 17.5 2 12 2m0 12.5c-1.4 0-2.5-1.1-2.5-2.5s1.1-2.5 2.5-2.5s2.5 1.1 2.5 2.5s-1.1 2.5-2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bag(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6h-3V5c0-1.1-.9-2-2-2h-4c-1.1 0-2 .9-2 2v1H5C3.3 6 2 7.3 2 9v9c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V9c0-1.7-1.3-3-3-3m-9-1h4v1h-4zm10 13c0 .6-.4 1-1 1H5c-.6 0-1-.4-1-1v-5.6L8.7 14H15c.1 0 .2 0 .3-.1l4.7-1.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bars(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 11H4c-.6 0-1 .4-1 1s.4 1 1 1h16c.6 0 1-.4 1-1s-.4-1-1-1M4 8h16c.6 0 1-.4 1-1s-.4-1-1-1H4c-.6 0-1 .4-1 1s.4 1 1 1m16 8H4c-.6 0-1 .4-1 1s.4 1 1 1h16c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryBolt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 7H4c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h6.7c-.6 0-1-.4-1-1c0-.2 0-.3.1-.5l1.4-2.5H8c-.1 0-.2 0-.3-.1h-.2l-.1-.1c-.1 0-.1-.1-.2-.1c-.1-.1-.1-.2-.2-.3V12c0-.1.1-.3.1-.4v-.1l2.3-4c.3-.5.9-.6 1.4-.4c.5.3.6.9.4 1.4L9.7 11h3.4c.1 0 .3.1.4.1h.1l.1.1c.1 0 .1.1.2.1c.1.1.1.2.2.3v.4c0 .1-.1.3-.1.4v.1l-2.3 4c-.3.3-.7.5-1 .5H17c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m4 3c-.6 0-1 .4-1 1v2c0 .6.4 1 1 1s1-.4 1-1v-2c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10c-.6 0-1 .4-1 1v2c0 .6.4 1 1 1s1-.4 1-1v-2c0-.6-.4-1-1-1m-4-3H4c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h13c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C6.3 2 5 3.3 5 5v16c0 .2 0 .3.1.5c.3.5.9.6 1.4.4l5.5-3.2l5.5 3.2c.2.1.3.1.5.1c.6 0 1-.4 1-1V5c0-1.7-1.3-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 3.5c0-.6-.4-1-1-1h-16c-.6 0-1 .4-1 1v16c0 .6.4 1 1 1s1-.4 1-1v-15h15c.6 0 1-.4 1-1m-1 7c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-12 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottom(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9.5c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0 4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0 4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0-12c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m8 0c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m-12 0c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m8 0c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0 8c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m-12-4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m16 6c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 12H4c-.6 0-1 .4-1 1s.4 1 1 1h16c.6 0 1-.4 1-1s-.4-1-1-1m-16-6c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0-8c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m4 8c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m-4 4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderClear(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-12c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1M4 15c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1M4 7c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m8 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1M8 19c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m8 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-8-8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m8 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-12c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-8 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-8 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m16 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-6c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderHorizontal(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 11H4c-.6 0-1 .4-1 1s.4 1 1 1h16c.6 0 1-.4 1-1s-.4-1-1-1m-8 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-10c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0-4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1M4 15c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1M4 9c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0-4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m4 0c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m8 0c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1M8 19c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m8 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-12c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-2c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInner(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 12c0-.6-.4-1-1-1h-7V4c0-.6-.4-1-1-1s-1 .4-1 1v7H4c-.6 0-1 .4-1 1s.4 1 1 1h7v7c0 .6.4 1 1 1s1-.4 1-1v-7h7c.5 0 1-.4 1-1M4 15c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1M4 9c.6 0 1-.4 1-1s-.5-1-1-1s-1 .4-1 1s.4 1 1 1m0-4c.6 0 1-.4 1-1s-.5-1-1-1s-1 .4-1 1s.4 1 1 1m4 0c.6 0 1-.4 1-1s-.5-1-1-1s-1 .4-1 1s.4 1 1 1m8 0c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1M8 19c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1m8 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1m4-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.5-1-1-1m0-10c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0-4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 3c-.6 0-1 .4-1 1v16c0 .6.4 1 1 1s1-.4 1-1V4c0-.6-.4-1-1-1m4 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-12-8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 2c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0 2c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-8-8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOut(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4c-.6 0-1 .4-1 1v16c0 .6.4 1 1 1h16c.6 0 1-.4 1-1V4c0-.6-.4-1-1-1m-1 16H5V5h14zm-7-6c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0 4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0-8c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m-4 4c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m8 0c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 3c-.6 0-1 .4-1 1v16c0 .6.4 1 1 1s1-.4 1-1V4c0-.6-.4-1-1-1m-4 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m12 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m8 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTop(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4.5h16c.6 0 1-.4 1-1s-.4-1-1-1H4c-.6 0-1 .4-1 1s.4 1 1 1m8 2c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m8-12c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-8 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m8-8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-8 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderVertical(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3c-.6 0-1 .4-1 1v16c0 .6.4 1 1 1s1-.4 1-1V4c0-.6-.4-1-1-1m-4 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m12 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1M7 3c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1M3 3c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m12 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 2c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1m0 2c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1M3 7c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0 8c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m12 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 0c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 15.5V17c0 .6-.4 1-1 1s-1-.4-1-1v-1.5H9V17c0 .6-.4 1-1 1s-1-.4-1-1v-1.5H5c-.7 0-1.4-.2-2-.5v4c0 1.7 1.3 3 3 3h12c1.7 0 3-1.3 3-3v-4c-.6.3-1.3.5-2 .5zM21 6h-4V5c0-1.7-1.3-3-3-3h-4C8.3 2 7 3.3 7 5v1H3c-.6 0-1 .4-1 1v4c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V7c0-.6-.4-1-1-1m-6 0H9V5c0-.6.4-1 1-1h4c.6 0 1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calender(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 19c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3v-8H2zM19 4h-2V3c0-.6-.4-1-1-1s-1 .4-1 1v1H9V3c0-.6-.4-1-1-1s-1 .4-1 1v1H5C3.3 4 2 5.3 2 7v2h20V7c0-1.7-1.3-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5C3.3 2 2 3.3 2 5v14c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V5c0-1.7-1.3-3-3-3M8 17c0 .6-.4 1-1 1s-1-.4-1-1v-4c0-.6.4-1 1-1s1 .4 1 1zm5 0c0 .6-.4 1-1 1s-1-.4-1-1V7c0-.6.4-1 1-1s1 .4 1 1zm5 0c0 .6-.4 1-1 1s-1-.4-1-1v-6c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m0 2c4 0 7.4 3 7.9 7H12zm4 14.9L12.6 13H20c-.4 2.5-1.8 4.7-4 5.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.7 7.2c-.4-.4-1-.4-1.4 0l-7.5 7.5l-3.1-3.1c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l3.8 3.8c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l8.2-8.2c.4-.4.4-1 0-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m4.2 8.3l-4.8 4.8c-.4.4-1 .4-1.4 0l-2.2-2.2c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l1.5 1.5l4.1-4.1c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h18c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1m-3.3 7.3l-6.8 6.8c-.4.4-1 .4-1.4 0l-3.2-3.2c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l2.5 2.5l6.1-6.1c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLayer(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 12c-2.8 0-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5s-2.2-5-5-5m4-5c-2.4 0-4.6 1.5-5.5 3.7c3.5-.9 7 1.3 7.8 4.7c.3 1 .3 2.1 0 3.1c3.1-1.3 4.5-4.8 3.2-7.8C15.6 8.5 13.4 7 11 7m10.2-1.2C20.1 3.5 17.6 2 15 2S9.9 3.5 8.8 5.8c4-1.2 8.2 1 9.4 4.9c.5 1.5.5 3 0 4.5c3.4-1.7 4.8-5.9 3-9.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClinicMedical(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.7 10.3l-9-8c-.4-.3-.9-.3-1.3 0l-9 8c-.4.4-.5 1-.1 1.4s1 .5 1.4.1l.3-.4V21c0 .6.4 1 1 1h14c.6 0 1-.4 1-1v-9.6l.3.3c.4.4 1 .3 1.4-.1c.4-.3.4-1 0-1.3M14 15h-1v1c0 .6-.4 1-1 1s-1-.4-1-1v-1h-1c-.6 0-1-.4-1-1s.4-1 1-1h1v-1c0-.6.4-1 1-1s1 .4 1 1v1h1c.6 0 1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m3.5 12c-.3.5-.9.6-1.4.4l-2.6-1.5c-.3-.2-.5-.5-.5-.9V7c0-.6.4-1 1-1s1 .4 1 1v4.4l2.1 1.2c.5.3.6.9.4 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockEight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m1 10c0 .4-.2.7-.5.9l-2.6 1.5c-.5.3-1.1.1-1.4-.4s-.1-1.1.4-1.4l2.1-1.2V7c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockFive(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m2 13.7c-.5.3-1.1.1-1.4-.4l-1.5-2.8c-.1-.2-.1-.3-.1-.5V7c0-.6.4-1 1-1s1 .4 1 1v4.7l1.4 2.6c.2.5.1 1.1-.4 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockNine(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m1 10c0 .6-.4 1-1 1H9c-.6 0-1-.4-1-1s.4-1 1-1h2V7c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSeven(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m1 10c0 .2-.1.3-.1.4l-1.5 2.8c-.3.5-.9.7-1.4.4c-.5-.3-.7-.9-.4-1.4l1.4-2.6V7c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTen(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m1 10c0 .6-.4 1-1 1c-.2 0-.3 0-.5-.1l-2.6-1.5c-.5-.3-.6-.9-.4-1.4c.3-.5.9-.6 1.4-.4l1.1.6V7c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockThree(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m3 11h-3c-.6 0-1-.4-1-1V7c0-.6.4-1 1-1s1 .4 1 1v4h2c.6 0 1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTwo(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m3.1 9.4l-2.6 1.5c-.5.3-1.1.1-1.4-.4c-.1-.2-.1-.3-.1-.5V7c0-.6.4-1 1-1s1 .4 1 1v3.3l1.1-.6c.5-.3 1.1-.1 1.4.4s.1 1-.4 1.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Columns(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3v18c0 .6.4 1 1 1h8V2H3c-.6 0-1 .4-1 1m19-1h-8v20h8c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDots(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12c0 2.3.8 4.5 2.3 6.3l-2 2c-.4.4-.4 1 0 1.4c.2.2.4.3.7.3h9c5.5 0 10-4.5 10-10S17.5 2 12 2M8 13c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m4 0c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m4 0c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15H3c-.6 0-1 .4-1 1s.4 1 1 1h4v4c0 .6.4 1 1 1s1-.4 1-1v-5c0-.6-.4-1-1-1M8 2c-.6 0-1 .4-1 1v4H3c-.6 0-1 .4-1 1s.4 1 1 1h5c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1m8 7h5c.6 0 1-.4 1-1s-.4-1-1-1h-4V3c0-.6-.4-1-1-1s-1 .4-1 1v5c0 .6.4 1 1 1m5 6h-5c-.6 0-1 .4-1 1v5c0 .6.4 1 1 1s1-.4 1-1v-4h4c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.7 2c-.6 0-1 .4-1 1v10.4c0 1.1-.9 2-2 2h-8l2.9-2.9c.4-.4.4-1 0-1.4s-1-.4-1.4 0l-4.6 4.6c-.4.4-.4 1 0 1.4l4.6 4.6c.2.2.5.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-2.9-2.9h8c2.2 0 4-1.8 4-4V3c0-.6-.5-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.4 15.7L14.8 11c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2.9 2.9h-7c-1.7 0-3-1.3-3-3V3c0-.6-.4-1-1-1s-1 .4-1 1v9.4c0 2.8 2.2 5 5 5h7l-2.9 2.9c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l4.6-4.6c.4-.4.4-1 0-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLeftDown(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4.3h-9.4c-2.8 0-5 2.2-5 5v7l-2.9-2.9c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l4.6 4.6c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l4.6-4.6c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-2.9 2.9v-7c0-1.7 1.3-3 3-3H21c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerRightDown(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.7 13.4c-.4-.4-1-.4-1.4 0l-2.9 2.9v-8c0-2.2-1.8-4-4-4H3c-.6 0-1 .4-1 1s.4 1 1 1h10.4c1.1 0 2 .9 2 2v8l-2.9-2.9c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l4.6 4.6c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l4.6-4.6c.4-.4.4-1.1 0-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.7 6.6h-7l2.9-2.9c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L4.6 6.9c-.4.4-.4 1 0 1.4L9.2 13c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-2.9-3h7c1.7 0 3 1.3 3 3V21c0 .6.4 1 1 1s1-.4 1-1v-9.4c0-2.7-2.3-5-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.4 6.9l-4.6-4.6c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2.9 2.9h-8c-2.2 0-4 1.8-4 4V21c0 .6.4 1 1 1s1-.4 1-1V10.6c0-1.1.9-2 2-2h8l-2.9 2.9c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l4.6-4.6c.4-.4.4-1 0-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coronavirus(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 11h-1.1c-.2-1.7-.9-3.3-1.9-4.6l.8-.8c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-.8.8c-1.3-1.1-2.9-1.7-4.6-1.9V2c0-.6-.4-1-1-1s-1 .4-1 1v1.1c-1.7.1-3.3.8-4.6 1.9l-.8-.8c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l.8.8C3.9 7.7 3.2 9.3 3.1 11H2c-.6 0-1 .4-1 1s.4 1 1 1h1.1c.2 1.7.9 3.3 1.9 4.6l-.8.8c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0l.8-.8c1.3 1.1 2.9 1.7 4.6 1.9V22c0 .6.4 1 1 1s1-.4 1-1v-1.1c1.7-.2 3.3-.9 4.6-1.9l.8.8c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4l-.8-.8c1.1-1.3 1.7-2.9 1.9-4.6H22c.6 0 1-.4 1-1s-.4-1-1-1M9 16c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m.5-4c-.8 0-1.5-.7-1.5-1.5S8.7 9 9.5 9s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5m5 3c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5m.5-5c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dialpad(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 9.2h-4c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8v-4c0-.4-.4-.8-.8-.8m0 7h-4c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8v-4c0-.4-.4-.8-.8-.8m-7-14H3c-.4 0-.8.4-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8V3c0-.4-.4-.8-.8-.8m0 7H3c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8v-4c0-.4-.4-.8-.8-.8m14-7h-4c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8V3c0-.4-.4-.8-.8-.8m-7 0h-4c-.4 0-.8.4-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8V3c0-.4-.4-.8-.8-.8m7 7h-4c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8v-4c0-.4-.4-.8-.8-.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Direction(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.3 14.8L12 17.1l-2.3-2.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l3 3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0m-4.6-4.6L12 7.9l2.3 2.3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-3-3c-.4-.4-1-.4-1.4 0l-3 3c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentLayoutCenter(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 8h2c.6 0 1-.4 1-1s-.4-1-1-1h-2c-.6 0-1 .4-1 1s.4 1 1 1m2 2h-2c-.6 0-1 .4-1 1s.4 1 1 1h2c.6 0 1-.4 1-1s-.4-1-1-1M3 8h2c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m0 4h2c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m6 0h6c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1H9c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1m12 2H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m-8 4H3c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentLayoutLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 8h8c.6 0 1-.4 1-1s-.4-1-1-1h-8c-.6 0-1 .4-1 1s.4 1 1 1m8 2h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1M3 12h6c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1H3c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1m18 2H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m-8 4H3c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentLayoutRight(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 18H3c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1M3 8h8c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m0 4h8c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18-8h-6c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h6c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1m0 10H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 9h-2c-.6 0-1 .4-1 1s.4 1 1 1h2c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1H6c-.6 0-1-.4-1-1v-7c0-.6.4-1 1-1h2c.6 0 1-.4 1-1s-.4-1-1-1H6c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h12c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3m-9.7 5.7l3 3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L13 14.6V3c0-.6-.4-1-1-1s-1 .4-1 1v11.6l-1.3-1.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisH(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2m-7 0c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2m14 0c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisV(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2m0-3c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2s.9 2 2 2m0 10c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircle(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m0 15c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m1-5c0 .6-.4 1-1 1s-1-.4-1-1V8c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationOctagon(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.7 7.6l-5.3-5.3c-.2-.2-.4-.3-.7-.3H8.3c-.3 0-.5.1-.7.3L2.3 7.6c-.2.2-.3.4-.3.7v7.5c0 .3.1.5.3.7l5.3 5.3c.2.1.4.2.7.2h7.5c.3 0 .5-.1.7-.3l5.3-5.3c.2-.2.3-.4.3-.7V8.3c-.1-.3-.2-.5-.4-.7M12 17c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m1-5c0 .6-.4 1-1 1s-1-.4-1-1V8c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangle(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.7 17.5l-8.1-14c-.8-1.4-2.7-1.9-4.1-1.1c-.5.3-.9.7-1.1 1.1l-8.1 14c-.8 1.4-.3 3.3 1.1 4.1c.5.3 1 .4 1.5.4H20c1.7 0 3-1.4 3-3c.1-.6-.1-1.1-.3-1.5M12 18c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m1-5c0 .6-.4 1-1 1s-1-.4-1-1V9c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Favorite(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.1c.1-.5-.3-1.1-.8-1.1l-5.7-.8L12.9 3c-.1-.2-.2-.3-.4-.4c-.5-.3-1.1-.1-1.4.4L8.6 8.2L2.9 9c-.3 0-.5.1-.6.3c-.4.4-.4 1 0 1.4l4.1 4l-1 5.7c0 .2 0 .4.1.6c.3.5.9.7 1.4.4l5.1-2.7l5.1 2.7c.1.1.3.1.5.1h.2c.5-.1.9-.6.8-1.2l-1-5.7l4.1-4c.2-.1.3-.3.3-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipH(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1m-4 4H7c-.3 0-.5.1-.7.3c-.4.4-.4 1 0 1.4l5 5c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l5-5c.2-.2.3-.4.3-.7c0-.6-.4-1-1-1m-1.6-7c0 .6.4 1 1 1h.6c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-.7-.7c-.3-.3-.8-.4-1.2-.2c-.5.2-.7.8-.4 1.3zm-3.8-3.2l.4-.4l.8.8c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-1.1-1.1c-.1-.1-.2-.1-.3-.2c0-.1-.1-.2-.2-.3c-.4-.4-1-.4-1.4 0l-1.1 1.1c-.4.4-.4 1 0 1.4s1.2.5 1.5.1M7.1 9c.3 0 .5-.1.7-.3l1.1-1.1c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-1.2 1c-.1.2-.2.5-.2.7c0 .6.4 1 1 1.1m3.8 0h1.5c.6 0 1-.4 1-1s-.4-1-1-1h-1.5c-.6 0-1 .4-1 1s.5 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipHalt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.4 7.2h.2l1.9-.5c.5-.1.9-.7.7-1.2c-.1-.5-.7-.9-1.2-.7l-1.9.4c-.4.1-.8.5-.8 1c.1.5.5 1 1.1 1M9.5 9h-2c-.6 0-1 .4-1 1s.4 1 1 1h2c.6 0 1-.4 1-1s-.4-1-1-1m3.7-3.3h.2l1.9-.5c.6-.2 1-.7.8-1.2c-.1-.5-.7-.9-1.2-.7l-1.9.4c-.4.1-.8.5-.8 1s.5 1 1 1m.3 3.3c-.6 0-1 .4-1 1s.4 1 1 1h2c.6 0 1-.4 1-1s-.4-1-1-1zM21 2.8c-.1-.5-.7-.9-1.2-.7l-1 .2c-.4.1-.8.5-.8 1c0 .6.4 1 1 1c.1.4.5.7 1 .7c.6 0 1-.4 1-1V3zM4 10.5c.6 0 1-.4 1-1v-2c0-.6-.4-1-1-1s-1 .5-1 1v2c0 .6.4 1 1 1M20 7c-.6 0-1 .4-1 1v1.1c-.3.2-.5.5-.5.9c0 .6.4 1 1 1h.5c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1m0 6H4c-.6 0-1 .4-1 1v3c0 .5.3.9.8 1l16 4h.2c.6 0 1-.4 1-1v-7c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipV(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1s1-.4 1-1V3c0-.6-.4-1-1-1m9.7 9.3l-5-5c-.2-.2-.4-.3-.7-.3c-.6 0-1 .4-1 1v10c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l5-5c.4-.4.4-1 0-1.4M3.8 9.8l-1.1 1.1c-.1.1-.1.2-.2.3c-.1 0-.2.1-.3.2c-.4.4-.4 1 0 1.4l1.1 1.1c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-.3-.5l.8-.8c.4-.4.4-1 0-1.4s-1-.4-1.4 0m3.8 5.4c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l1.1 1.1c.1.1.4.3.6.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4zm-.3-8.9l-.7.7c-.2.1-.3.4-.3.7c0 .6.4 1 1 1c.2 0 .3 0 .5-.1H8c.6 0 1-.4 1-1V7c0-.3-.1-.5-.3-.7c-.4-.4-1-.4-1.4 0m.7 4.3c-.6 0-1 .4-1 1v1.5c0 .6.4 1 1 1s1-.4 1-1v-1.5c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipValt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.9 12.2c-.5-.1-1.1.2-1.2.7l-.5 1.9c-.1.6.2 1.2.8 1.3h.2c.5 0 .9-.3 1-.8l.5-1.9c.1-.5-.2-1-.8-1.2M7.5 5h2c.6 0 1-.4 1-1s-.4-1-1-1h-2c-.6 0-1 .4-1 1s.5 1 1 1M4.2 19c0-.5-.3-.9-.8-1c-.5-.1-1.1.2-1.2.7l-.2 1v.2c0 .6.4 1 1 1h1c.5 0 .9-.3 1-.8c.1-.4-.2-.9-.8-1.1M6.4 6.4c-.5-.1-1.1.2-1.2.7l-.5 2v.2c0 .6.4 1 1 1c.5 0 .9-.3 1-.8l.5-1.9c.1-.5-.3-1-.8-1.2m3.6 6.1c-.6 0-1 .4-1 1v2c0 .6.4 1 1 1s1-.4 1-1v-2c0-.5-.4-1-1-1m0-6c-.6 0-1 .4-1 1v2c0 .6.4 1 1 1s1-.4 1-1v-2c0-.5-.4-1-1-1m.5 12.1c-.5-.2-1.1-.1-1.4.4H8c-.6 0-1 .4-1 1s.4 1 1 1h2c.6 0 1-.4 1-1v-.5c0-.3-.2-.7-.5-.9M22 19.8l-4-16c-.1-.5-.5-.8-1-.8h-3c-.6 0-1 .4-1 1v16c0 .6.4 1 1 1h7.2c.6-.2.9-.7.8-1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphBar(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 13H2c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1h4c.6 0 1-.4 1-1v-8c0-.6-.4-1-1-1m16-4h-4c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1h4c.6 0 1-.4 1-1V10c0-.6-.4-1-1-1m-8-8h-4c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h4c.6 0 1-.4 1-1V2c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v4h20V3c0-.6-.4-1-1-1M2 15h9V9H2zm0 6c0 .6.4 1 1 1h8v-5H2zm11-6h9V9h-9zm0 7h8c.6 0 1-.4 1-1v-4h-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grids(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3.9v16.2c0 1 .9 1.9 1.9 1.9H8V2H3.9C2.9 2 2 2.9 2 3.9M20.1 2H16v20h4.1c1 0 1.9-.9 1.9-1.9V3.9c0-1-.9-1.9-1.9-1.9M10 22h4V2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GripHorizontalLine(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 11h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18 2H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSide(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.2 2c-2-.1-4 .7-5.5 2.1c-1.4 1.4-2.2 3.4-2.2 5.4l-1.9 4c-.2.5 0 1.1.5 1.3c.1.2.3.2.4.2h1v2c0 1.1.9 2 2 2h1v2c0 .6.4 1 1 1h9.3c.5-.2.8-.7.7-1.2l-.9-3.2l1.9-7.3v-.5c0-4.1-3.2-7.6-7.3-7.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSideCough(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 16c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-4 1c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m3 3c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1M16.2 2c-4-.1-7.2 2.9-7.3 6.9v.2l-1.8 3.8c-.2.5 0 1.1.5 1.3c.1.1.3.1.4.1h.9v1.8c0 1.1.9 1.9 1.9 1.9h.9v1.8c0 .6.4 1 1 1h8.7c.5-.2.8-.7.7-1.2l-.9-3L23 9.8v-.5c0-3.9-3-7.1-6.8-7.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSideMask(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.2 2c-2-.1-4 .7-5.5 2.1c-1.4 1.4-2.2 3.3-2.2 5.4l-1.8 3c-.1.2-.2.3-.2.5c0 .1 0 .2.1.3L5 17.2v.1c.5 1 1.5 1.7 2.7 1.7h.8v2c0 .6.4 1 1 1s1-.4 1-1v-2h2c.1 0 .2 0 .3-.1l3.7-1.3v.1l1 3.5c.1.4.5.7 1 .7h.3c.5-.2.8-.7.7-1.2l-.9-3.2l1.9-7.3v-.4c0-4.1-3.2-7.6-7.3-7.8M17 15.4l-3.5 1.2v-2.9l4.3-1.4zm1.4-5.4h-.3l-5.8 2h-6l1.1-1.7c.1-.2.2-.4.1-.6v-.2C7.5 6.5 10 4 13 4h.2c3 .2 5.4 2.7 5.3 5.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3.3C13.1 1.1 8.3 1.8 5.1 4.8V3c0-.6-.4-1-1-1s-1 .4-1 1v4.5c0 .6.4 1 1 1h4.5c.6 0 1-.4 1-1s-.4-1-1-1H6.2C7.7 4.9 9.8 4 12 4c4.4 0 8 3.6 8 8s-3.6 8-8 8s-8-3.6-8-8c0-.6-.4-1-1-1s-1 .4-1 1c0 5.5 4.5 10 10 10c3.6 0 6.9-1.9 8.7-5c2.7-4.8 1.1-10.9-3.7-13.7M12 8c-.6 0-1 .4-1 1v3c0 .6.4 1 1 1h2c.6 0 1-.4 1-1s-.4-1-1-1h-1V9c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistoryAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.4 3.3C12.5 1.1 7.7 1.8 4.6 4.8V3c0-.6-.4-1-1-1s-1 .4-1 1v4.5c0 .6.4 1 1 1h4.5c.6 0 1-.4 1-1s-.4-1-1-1H5.7C7.1 4.9 9.2 4 11.5 4c4.4 0 8 3.6 8 8s-3.6 8-8 8c-.6 0-1 .4-1 1s.4 1 1 1c3.6 0 6.9-1.9 8.7-5c2.7-4.8 1-10.9-3.8-13.7m-5 4.7c-.6 0-1 .4-1 1v3c0 .6.4 1 1 1h2c.6 0 1-.4 1-1s-.4-1-1-1h-1V9c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalAlignLeft(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10h-5V7c0-.6-.4-1-1-1H4V3c0-.6-.4-1-1-1c-.5 0-1 .4-1 1v18c0 .6.4 1 1 1s1-.4 1-1v-3h17c.6 0 1-.4 1-1v-6c0-.6-.4-1-1-1m-7 0H4V8h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 6.5h-3v-4c0-.6-.4-1-1-1h-11c-.6 0-1 .4-1 1v4h-3c-.6 0-1 .4-1 1v14c0 .6.4 1 1 1h19c.6 0 1-.4 1-1v-14c0-.6-.4-1-1-1m-14 12h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1m0-4h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1m5 4h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1m0-4h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1m1-5.5H13v.5c0 .6-.4 1-1 1s-1-.4-1-1V9h-.5c-.6 0-1-.4-1-1s.4-1 1-1h.5v-.5c0-.6.4-1 1-1s1 .4 1 1V7h.5c.6 0 1 .4 1 1s-.4 1-1 1m4 9.5h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1m0-4h-1c-.6 0-1-.4-1-1s.4-1 1-1h1c.6 0 1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalSquareSign(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5C3.3 2 2 3.3 2 5v14c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V5c0-1.7-1.3-3-3-3m-3 15c0 .6-.4 1-1 1s-1-.4-1-1v-4h-4v4c0 .6-.4 1-1 1s-1-.4-1-1V7c0-.6.4-1 1-1s1 .4 1 1v4h4V7c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalSymbol(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m4 14c0 .6-.4 1-1 1s-1-.4-1-1v-3h-4v3c0 .6-.4 1-1 1s-1-.4-1-1V8c0-.6.4-1 1-1s1 .4 1 1v3h4V8c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseUser(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.7 10.3l-9-8c-.4-.3-.9-.3-1.3 0l-9 8c-.4.4-.5 1-.1 1.4s1 .5 1.4.1l.3-.4V21c0 .6.4 1 1 1h14c.6 0 1-.4 1-1v-9.6l.3.3c.4.4 1 .3 1.4-.1c.4-.3.4-1 0-1.3M12 11c1.5 0 2.7 1.2 2.7 2.7c0 1.5-1.2 2.7-2.7 2.7c-1.5 0-2.7-1.2-2.7-2.7S10.5 11 12 11m-5 9c0-.1 0-.1.1-.2c2.2-2.7 6.2-3.2 8.9-1c.4.3.7.6 1 1c0 0 0 .1.1.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageV(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5C3.3 2 2 3.3 2 5v14c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V5c0-1.7-1.3-3-3-3m1 11.9L18.1 12c-1.2-1.1-3.1-1.1-4.2 0l-.9.9l-2.9-2.9C8.9 8.9 7 8.9 5.9 10L4 11.9V5c0-.6.4-1 1-1h14c.6 0 1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySkeleton(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21 4.4l.7-.7c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L9.8 12.8C9 12.3 8 12 7 12c-2.8 0-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5c0-1-.3-2-.8-2.8l5.6-5.6l2.1 2.1c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4l-2.1-2.1l1.4-1.4l.7.7c.4.4 1 .4 1.4 0s.4-1 0-1.4zM7 20c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySkeletonAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.3 5.1l1.4-1.4c.4-.4.4-1 0-1.4s-1-.4-1.4 0L9.8 12.8C9 12.3 8 12 7 12c-2.8 0-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5c0-1-.3-2-.8-2.8l4.9-4.9l1.4 1.4c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4l-1.4-1.4l1.4-1.4l1.4 1.4c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4zM7 20c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyholeCircle(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m1.7 9c-.2.3-.4.6-.7.7V15c0 .6-.4 1-1 1s-1-.4-1-1v-3.3c-1-.6-1.3-1.8-.7-2.7S12 7.7 13 8.3c1 .5 1.3 1.7.7 2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyholeSquare(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5C3.3 2 2 3.3 2 5v14c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V5c0-1.7-1.3-3-3-3m-5.3 9c-.2.3-.4.6-.7.7V15c0 .6-.4 1-1 1s-1-.4-1-1v-3.3c-1-.6-1.3-1.8-.7-2.7S12 7.7 13 8.3c1 .5 1.3 1.7.7 2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyholeSquareFull(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h18c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1m-7.3 9c-.2.3-.4.6-.7.7V15c0 .6-.4 1-1 1s-1-.4-1-1v-3.3c-1-.6-1.3-1.8-.7-2.7S12 7.7 13 8.3c1 .5 1.3 1.7.7 2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayerGroup(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.5 8.9l9 5.2c.2.1.3.1.5.1s.3 0 .5-.1l9-5.2c.2-.1.3-.2.4-.4c.2-.5.1-1.1-.4-1.4l-9-5.2c-.3-.2-.7-.2-1 0l-9 5.2c-.2.1-.3.2-.4.4c-.2.5-.1 1.1.4 1.4m19 2.2l-.2-.1l-8.8 5.1c-.3.2-.7.2-1 0L2.7 11l-.2.1c-.5.3-.6.9-.4 1.4c.1.2.2.3.4.4l9 5.2c.3.2.7.2 1 0l9-5.2c.5-.3.6-.9.4-1.4c-.1-.2-.2-.3-.4-.4m0 4l-.2-.1l-8.8 5.1c-.3.2-.7.2-1 0L2.7 15l-.2.1c-.5.3-.6.9-.4 1.4c.1.2.2.3.4.4l9 5.2c.3.2.7.2 1 0l9-5.2c.5-.3.6-.9.4-1.4c-.1-.2-.2-.3-.4-.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7H6c-.6 0-1 .4-1 1v2.5h7.5c.6 0 1 .4 1 1V19H16c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1m-5 5H3c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1h8c.6 0 1-.4 1-1v-8c0-.6-.4-1-1-1M21 2H9c-.6 0-1 .4-1 1v2.5h9.5c.6 0 1 .4 1 1V16H21c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftIndent(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m0 4h10c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18 6H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1M3 15h10c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18.8-5.3c-.4-.4-1-.5-1.4-.1l-2 1.7l-.1.1c-.4.4-.3 1.1.1 1.4l2 1.7c.2.1.4.2.6.2c.3 0 .6-.1.8-.4c.4-.4.3-1.1-.1-1.4l-1.1-.9l1.1-.9c.4-.4.4-1 .1-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftIndentAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 17h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1m0-4h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1m-8-6h8c.6 0 1-.4 1-1s-.4-1-1-1h-8c-.6 0-1 .4-1 1s.4 1 1 1m8 2h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1M9 5c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1s1-.4 1-1V6c0-.6-.4-1-1-1m-4.4 7l1.1-.9c.4-.4.5-1 .1-1.4c-.4-.4-1-.5-1.4-.1l-2 1.7l-.1.1c-.4.4-.3 1.1.1 1.4l2 1.7c.2.1.4.2.6.2c.3 0 .6-.1.8-.4c.4-.4.3-1.1-.1-1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineSpacing(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8h11c.6 0 1-.4 1-1s-.4-1-1-1H10c-.6 0-1 .4-1 1s.4 1 1 1m-4.3 7.3V8.7c.2.2.4.3.6.3c.3 0 .5-.1.7-.2c.4-.4.5-1 .1-1.4l-1.7-2C5.2 5.1 5 5 4.7 5s-.6.1-.8.4l-1.7 2c-.3.4-.3 1 .2 1.4c.4.3.9.3 1.3 0v6.6c-.4-.3-.9-.4-1.3 0s-.5 1-.1 1.4l1.7 2c.1.1.4.2.7.2s.6-.1.8-.4l1.7-2c.4-.4.3-1.1-.1-1.4c-.5-.3-1.1-.3-1.4.1M21 11H10c-.6 0-1 .4-1 1s.4 1 1 1h11c.6 0 1-.4 1-1s-.4-1-1-1m0 5H10c-.6 0-1 .4-1 1s.4 1 1 1h11c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkH(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 15H7c-1.7 0-3-1.3-3-3s1.3-3 3-3h3c.6 0 1-.4 1-1s-.4-1-1-1H7c-2.8 0-5 2.2-5 5s2.2 5 5 5h3c.6 0 1-.4 1-1s-.4-1-1-1m7-8h-3c-.6 0-1 .4-1 1s.4 1 1 1h3c1.7 0 3 1.3 3 3s-1.3 3-3 3h-3c-.6 0-1 .4-1 1s.4 1 1 1h3c2.8 0 5-2.2 5-5s-2.2-5-5-5m-9 5c0 .6.4 1 1 1h6c.6 0 1-.4 1-1s-.4-1-1-1H9c-.6 0-1 .4-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUiAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 8h14c.6 0 1-.4 1-1s-.4-1-1-1h-14c-.6 0-1 .4-1 1s.4 1 1 1m14 3h-10c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1m0 5h-6c-.6 0-1 .4-1 1s.4 1 1 1h6c.6 0 1-.4 1-1s-.4-1-1-1M3.5 6c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 5c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m4 5c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUl(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8h14c.6 0 1-.4 1-1s-.4-1-1-1H7c-.6 0-1 .4-1 1s.4 1 1 1m14 3H7c-.6 0-1 .4-1 1s.4 1 1 1h14c.6 0 1-.4 1-1s-.4-1-1-1m0 5H7c-.6 0-1 .4-1 1s.4 1 1 1h14c.6 0 1-.4 1-1s-.4-1-1-1M3.7 6.3c-.1-.1-.2-.2-.3-.2c-.4-.2-.8-.1-1.1.2c-.1.1-.2.2-.2.3c-.1.3-.1.5 0 .8c.1.1.1.2.2.3c.1.1.2.2.3.2c.1.1.3.1.4.1c.3 0 .5-.1.7-.3c.1-.1.2-.2.2-.3c.1-.3.1-.5 0-.8c0-.1-.1-.2-.2-.3m0 5c-.3-.3-.7-.4-1.1-.2c-.1.1-.2.1-.3.2c-.1.1-.2.2-.2.3c-.1.2-.1.5 0 .8c.1.1.1.2.2.3c.1.1.2.2.3.2c.1.1.3.1.4.1c.1 0 .3 0 .4-.1c.1-.1.2-.1.3-.2c.1-.1.2-.2.2-.3c.1-.2.1-.5 0-.8c0-.1-.1-.2-.2-.3m0 5c-.1-.1-.2-.2-.3-.2c-.2-.1-.5-.1-.8 0c-.1 0-.2.1-.3.2c-.1.1-.2.2-.2.3c-.2.4-.1.8.2 1.1c.1.1.2.2.3.2c.1.1.3.1.4.1c.1 0 .3 0 .4-.1c.1-.1.2-.1.3-.2c.3-.3.4-.7.2-1.1c0-.1-.1-.2-.2-.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9V7c0-2.8-2.2-5-5-5S7 4.2 7 7v2c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3M9 7c0-1.7 1.3-3 3-3s3 1.3 3 3v2H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockAccess(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2h-6c-.6 0-1 .4-1 1s.4 1 1 1h5v5c0 .6.4 1 1 1s1-.4 1-1V3c0-.6-.4-1-1-1M3 10c.6 0 1-.4 1-1V4h5c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1m6 10H4v-5c0-.6-.4-1-1-1s-1 .4-1 1v6c0 .6.4 1 1 1h6c.6 0 1-.4 1-1s-.4-1-1-1m12-6c-.6 0-1 .4-1 1v5h-5c-.6 0-1 .4-1 1s.4 1 1 1h6c.6 0 1-.4 1-1v-6c0-.6-.4-1-1-1m-9-8c-1.7 0-3 1.3-3 3v1c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2h6c1.1 0 2-.9 2-2v-4c0-1.1-.9-2-2-2V9c0-1.7-1.3-3-3-3m1 4h-2V9c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9V7c0-2.8-2.2-5-5-5S7 4.2 7 7v2c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3M9 7c0-1.7 1.3-3 3-3s3 1.3 3 3v2H9zm4 10c0 .6-.4 1-1 1s-1-.4-1-1v-3c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpenAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9H9V7c0-.8.3-1.5.9-2.1c1.2-1.2 3.1-1.2 4.2 0c.4.4.6.9.8 1.4c.1.5.7.8 1.2.7c.5-.1.9-.7.7-1.2c-.2-.9-.7-1.7-1.3-2.3c-.9-1-2.2-1.5-3.5-1.5c-2.8 0-5 2.2-5 5v2c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3m-3.9 6.5l-.1.1V17c0 .6-.4 1-1 1s-1-.4-1-1v-1.4c-.6-.6-.7-1.5-.1-2.1c.6-.6 1.5-.7 2.1-.1c.6.5.7 1.5.1 2.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microscope(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21h-5.2c.1-.2.1-.4.1-.6c3-1.2 5-4.1 5-7.4c0-2-.8-3.9-2.1-5.4l.6-.6c.1-.1.2-.2.2-.4l.7-2.1c.1-.4 0-.8-.2-1L17 1.4c-.3-.3-.7-.4-1-.2l-2 .6c-.1 0-.3.1-.4.2L7.2 8.4c-.4.4-.4 1 0 1.4l-1.4 1.4c-.4.4-.4 1 0 1.4l2.1 2.1c.4.4 1 .4 1.4 0l1.4-1.4c.4.4 1 .4 1.4 0L16.5 9c1 1.1 1.5 2.5 1.5 4c0 2.3-1.4 4.4-3.5 5.4c-.9-1.4-2.7-1.9-4.1-1c-.4.3-.8.6-1 1.1c-.4-.2-.8-.3-1.1-.6c.4-.1.7-.5.7-.9c0-.6-.4-1-1-1H4c-.6 0-1 .4-1 1s.4 1 1 1h1.3c1 1.1 2.3 2 3.7 2.5c0 .2.1.4.1.5H4c-.6 0-1 .4-1 1s.4 1 1 1h16c.6 0 1-.4 1-1c0-.5-.4-1-1-1M8.6 12.6l-.7-.7l.7-.7l.7.7zM12 21c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareFull(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h18c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1m-6 11H9c-.6 0-1-.4-1-1s.4-1 1-1h6c.6 0 1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Multiply(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.4 12l6.3-6.3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L12 10.6L5.7 4.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l6.3 6.3l-6.3 6.3c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l6.3-6.3l6.3 6.3c.2.2.4.3.7.3s.5-.1.7-.3c.4-.4.4-1 0-1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectGroup(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18.3V5.7c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2c-.7 0-1.4.4-1.7 1H5.7c-.3-.6-1-1-1.7-1c-1.1 0-2 .9-2 2c0 .7.4 1.4 1 1.7v12.6c-.6.3-1 1-1 1.7c0 1.1.9 2 2 2c.7 0 1.4-.4 1.7-1h12.6c.3.6 1 1 1.7 1c1.1 0 2-.9 2-2c0-.7-.4-1.4-1-1.7m-2 0c-.3.2-.5.4-.7.7H5.7c-.2-.3-.4-.5-.7-.7V5.7c.3-.2.5-.4.7-.7h12.6c.2.3.4.5.7.7zM14 9V8c0-.6-.4-1-1-1H8c-.6 0-1 .4-1 1v5c0 .6.4 1 1 1h1v-3c0-1.1.9-2 2-2zm2 1h-5c-.6 0-1 .4-1 1v5c0 .6.4 1 1 1h5c.6 0 1-.4 1-1v-5c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectUngroup(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18.3v-6.6c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2c-.7 0-1.4.4-1.7 1H15V5.7c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2c-.7 0-1.4.4-1.7 1H5.7c-.3-.6-1-1-1.7-1c-1.1 0-2 .9-2 2c0 .7.4 1.4 1 1.7v6.6c-.6.3-1 1-1 1.7c0 1.1.9 2 2 2c.7 0 1.4-.4 1.7-1H9v3.3c-.6.3-1 1-1 1.7c0 1.1.9 2 2 2c.7 0 1.4-.4 1.7-1h6.6c.3.6 1 1 1.7 1c1.1 0 2-.9 2-2c0-.7-.4-1.4-1-1.7M5.7 13c-.2-.3-.4-.5-.7-.7V5.7c.3-.2.5-.4.7-.7h6.6c.2.3.4.5.7.7V9h-1.3c-.3-.6-1-1-1.7-1c-1.1 0-2 .9-2 2c0 .7.4 1.4 1 1.7V13zm7.3-.7c-.3.2-.5.4-.7.7H11v-1.3c.3-.2.5-.4.7-.7H13zm-.7 2.7c.3.6 1 1 1.7 1c1.1 0 2-.9 2-2c0-.7-.4-1.4-1-1.7V11h3.3c.2.3.4.5.7.7v6.6c-.3.2-.5.4-.7.7h-6.6c-.2-.3-.4-.5-.7-.7V15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Padlock(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9V7c0-2.8-2.2-5-5-5S7 4.2 7 7v2c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3M9 7c0-1.7 1.3-3 3-3s3 1.3 3 3v2H9zm4.1 8.5l-.1.1V17c0 .6-.4 1-1 1s-1-.4-1-1v-1.4c-.6-.6-.7-1.5-.1-2.1c.6-.6 1.5-.7 2.1-.1c.6.5.7 1.5.1 2.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.1 12.4l-6.2 6.2c-1.7 1.7-4.4 1.7-6 0c-1.7-1.7-1.7-4.4 0-6l8-8c1-.9 2.5-.9 3.5 0c1 1 1 2.6 0 3.5L10.5 15c-.3.3-.8.3-1.1 0c-.3-.3-.3-.8 0-1.1l5.1-5.1c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L8 12.6c-1.1 1.1-1.1 2.8 0 3.9c1.1 1 2.8 1 3.9 0l6.9-6.9c1.8-1.8 1.8-4.6 0-6.4c-1.8-1.8-4.6-1.8-6.4 0l-8 8c-1.2 1.2-1.8 2.8-1.8 4.4c0 3.5 2.8 6.2 6.3 6.2c1.7 0 3.2-.7 4.4-1.8l6.2-6.2c.4-.4.4-1 0-1.4s-1-.4-1.4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 13.5H3c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1m8-5H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pentagon(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.6 9.2l-9-6.5c-.4-.3-.8-.3-1.2 0l-9 6.5c-.3.2-.5.7-.4 1.1l3.4 10.6c.1.4.5.7 1 .7h11.1c.4 0 .8-.3 1-.7L22 10.3c.1-.4-.1-.9-.4-1.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Polygon(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.9 11.5l-4.5-7.8c-.2-.3-.5-.5-.9-.5h-9c-.4 0-.7.2-.9.5l-4.5 7.8c-.2.3-.2.7 0 1l4.5 7.8c.2.3.5.5.9.5h9c.4 0 .7-.2.9-.5l4.5-7.8c.1-.3.1-.7 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.7 15.3L13.4 12l3.3-3.3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-4 4c-.4.4-.4 1 0 1.4l4 4c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4M8 7c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1s1-.4 1-1V8c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Process(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 6H6.7C8.2 4.7 10 4 12 4c.3 0 .6 0 .9.1c.5.1 1-.3 1.1-.9c.1-.5-.3-1-.9-1.1c-.4-.1-.7-.1-1.1-.1c-2.4 0-4.7.9-6.5 2.4V3c0-.6-.4-1-1-1s-1 .4-1 1v4c0 .6.4 1 1 1h4c.6 0 1-.4 1-1s-.4-1-1-1M7 14.5c-.6 0-1 .4-1 1v1.8C4.7 15.8 4 14 4 12c0-.3 0-.6.1-.9c.1-.5-.3-1-.9-1.1c-.5-.1-1 .3-1.1.9c-.1.4-.1.7-.1 1.1c0 2.4.9 4.7 2.4 6.5H3c-.6 0-1 .4-1 1s.4 1 1 1h4c.3 0 .6-.2.8-.4c0-.1.1-.2.1-.3v-4.3c.1-.6-.3-1-.9-1m14-9c.6 0 1-.4 1-1s-.4-1-1-1h-4.2c-.1 0-.2.1-.3.1c-.1.1-.2.1-.2.2s-.1.2-.1.2v4.3c0 .6.4 1 1 1s1-.4 1-1V6.7c1.3 1.4 2 3.3 2 5.3c0 .3 0 .6-.1.9c-.1.5.3 1 .9 1.1h.1c.5 0 .9-.4 1-.9c0-.4.1-.7.1-1.1c0-2.4-.9-4.7-2.4-6.5zm-.7 11l-.3-.3c-.1-.1-.2-.1-.3-.1h-4.2c-.6 0-1 .4-1 1s.4 1 1 1h1.8c-1.4 1.3-3.3 2-5.3 2c-.3 0-.6 0-.9-.1c-.5-.1-1 .3-1.1.9s.3 1 .9 1.1c.4 0 .7.1 1.1.1c2.4 0 4.7-.9 6.5-2.4V21c0 .6.4 1 1 1s1-.4 1-1v-4c0-.2-.1-.4-.2-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecordAudio(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m0 16c-3.3 0-6-2.7-6-6s2.7-6 6-6s6 2.7 6 6s-2.7 6-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11c-.6 0-1 .4-1 1c0 2.9-1.5 5.5-4 6.9c-3.8 2.2-8.7.9-10.9-2.9C2.9 12.2 4.2 7.3 8 5.1c3.3-1.9 7.3-1.2 9.8 1.4h-2.4c-.6 0-1 .4-1 1s.4 1 1 1h4.5c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1s-1 .4-1 1v1.8C17 3 14.6 2 12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 12c0-1.7-1.3-3-3-3s-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3m2-8.7C13.1 1.1 8.3 1.8 5.1 4.7V3c0-.6-.4-1-1-1s-1 .4-1 1v4.5c0 .6.4 1 1 1h4.5c.6 0 1-.4 1-1s-.4-1-1-1H6.2C7.7 4.9 9.8 4 12 4c4.4 0 8 3.6 8 8c0 .6.4 1 1 1s1-.4 1-1c0-3.6-1.9-6.9-5-8.7m2.9 12.2h-4.5c-.6 0-1 .4-1 1s.4 1 1 1h2.4C16.3 19.1 14.2 20 12 20c-4.4 0-8-3.6-8-8c0-.6-.4-1-1-1s-1 .4-1 1c0 5.5 4.5 10 10 10c2.6 0 5-1 6.9-2.8V21c0 .6.4 1 1 1s1-.4 1-1v-4.5c0-.6-.5-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 17.5H4v-11h7.8l-.8.8c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l2.5-2.5c.4-.4.4-1 0-1.4l-2.5-2.5c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l.8.8H3c-.6 0-1 .4-1 1v13c0 .6.4 1 1 1h2.5c.6 0 1-.4 1-1s-.4-1-1-1M21 4.5h-2.5c-.6 0-1 .4-1 1s.4 1 1 1H20v11h-8.4l.8-.8c.4-.4.4-1 0-1.4s-1-.4-1.4 0l-2.5 2.5c-.4.4-.4 1 0 1.4l2.5 2.5c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-.8-.8H21c.6 0 1-.4 1-1v-13c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightIndent(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m0 4h10c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m18 6H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1M3 15h10c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m15-1.3c0 .6.4 1 1 1c.2 0 .5-.1.6-.2l2-1.7l.1-.1c.4-.4.3-1.1-.1-1.4l-2-1.7c-.4-.4-1.1-.3-1.4.1c-.4.4-.3 1.1.1 1.4l1.1.9l-1.1.9c-.2.2-.3.5-.3.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightIndentAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 17h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1m0-4h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1m-8-6h8c.6 0 1-.4 1-1s-.4-1-1-1h-8c-.6 0-1 .4-1 1s.4 1 1 1m8 2h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1M9 5c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1s1-.4 1-1V6c0-.6-.4-1-1-1M3.6 9.6c-.4-.4-1-.3-1.4.1c-.4.4-.3 1.1.1 1.4l1.1.9l-1.1.9c-.2.2-.3.5-.3.8c0 .6.4 1 1 1c.2 0 .5-.1.6-.2l2-1.7l.1-.1c.4-.4.3-1.1-.1-1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.6 2.1c-.1-.3-.4-.6-.7-.7c-4.2-1.1-8.5.2-11.4 3.6L9.4 6.3l-2.7-.6C5.4 5.2 4 5.8 3.4 7l-2.2 3.9c-.2.3-.2.6 0 .9c.1.3.4.5.7.6l3.1.7c-.3.8-.4 1.6-.6 2.4c0 .3.1.6.3.8l3.1 3.1c.2.2.4.3.7.3h.1c.9-.1 1.7-.2 2.5-.5l.6 3c.1.3.3.6.6.7c.1.1.3.1.4.1c.2 0 .3 0 .5-.1l3.9-2.2c1.1-.6 1.7-2 1.4-3.3l-.7-2.8l1.2-1.1c3.3-2.8 4.7-7.3 3.6-11.4M7.3 8.8c-.6.8-1.2 1.6-1.6 2.4l-2.1-.5L5.1 8c.2-.4.6-.5 1.1-.4l1.7.4zM16 18.9l-2.7 1.5l-.4-2c.9-.4 1.7-1 2.4-1.6l.7-.7l.4 1.7c.2.5-.1 1-.4 1.1m.7-10.1c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5s1.5.7 1.5 1.5s-.6 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.6 7.1l-5.7-5.7c-.4-.4-1-.4-1.4 0L1.4 15.5c-.4.4-.4 1 0 1.4l5.7 5.7c.4.4 1 .4 1.4 0l2.1-2.1L7.1 17c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l3.5 3.5l1.4-1.4l-2.1-2.1c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l2.1 2.1l1.4-1.4l-3.5-3.5c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l3.5 3.5l1.4-1.4l-2.1-2.1c-.4-.4-.4-1 0-1.4s1-.4 1.4 0l2.1 2.1l2.1-2.1c.5-.5.5-1.2.1-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RulerCombined(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h6c.6 0 1-.4 1-1v-3H7c-.6 0-1-.4-1-1s.4-1 1-1h3v-2H7c-.6 0-1-.4-1-1s.4-1 1-1h3v-2H7c-.6 0-1-.4-1-1s.4-1 1-1h1V7c0-.6.4-1 1-1s1 .4 1 1v3h2V7c0-.6.4-1 1-1s1 .4 1 1v3h2V7c0-.6.4-1 1-1s1 .4 1 1v3h3c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sanitizer(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.8 7.6L16 5.5V3h1c.6 0 1-.4 1-1s-.4-1-1-1H8.7c-1.4 0-2.6.5-3.6 1.5l-.8.8c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0l.8-.8c.6-.6 1.4-.9 2.2-.9H10v2.5L7.2 7.6C6.4 8.2 6 9.1 6 10v12c0 .6.4 1 1 1h12c.6 0 1-.4 1-1V10c0-.9-.4-1.8-1.2-2.4M12 3h2v2h-2zm1 14c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SanitizerAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8V5c0-.6-.4-1-1-1h-1V3h1c.6 0 1-.4 1-1s-.4-1-1-1h-4.8C8.5 1 6.9 2 6.1 3.6c-.2.4 0 1 .5 1.3c.5.2 1.1 0 1.3-.4c.4-.9 1.3-1.5 2.3-1.5H12v1h-1c-.6 0-1 .4-1 1v3c-1.7 0-3 1.3-3 3v9c0 1.7 1.3 3 3 3h6c1.7 0 3-1.3 3-3v-9c0-1.7-1.3-3-3-3m-4-2h2v2h-2zm2 11h-2c-.6 0-1-.4-1-1s.4-1 1-1h2c.6 0 1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scenery(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 6c-.8 0-1.5.7-1.5 1.5S12.7 9 13.5 9S15 8.3 15 7.5S14.3 6 13.5 6M19 2H5C3.3 2 2 3.3 2 5v14c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V5c0-1.7-1.3-3-3-3m1 11.9L18.1 12c-1.2-1.1-3.1-1.1-4.2 0l-.9.9l-2.9-2.9C8.9 8.9 7 8.9 5.9 10L4 11.9V5c0-.6.4-1 1-1h14c.6 0 1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Schedule(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-1V3c0-.6-.4-1-1-1s-1 .4-1 1v1H8V3c0-.6-.4-1-1-1s-1 .4-1 1v1H5C3.3 4 2 5.3 2 7v1h20V7c0-1.7-1.3-3-3-3M2 19c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3v-9H2zm15-7c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m0 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m-5-4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m0 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m-5-4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m0 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldPlus(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.6 3.6c-.2-.2-.5-.3-.8-.2c-2.2.5-4.4 0-6.2-1.3c-.3-.2-.8-.2-1.1 0c-1.9 1.3-4.1 1.8-6.3 1.3c-.5-.1-1.1.3-1.2.8v7.7c0 2.9 1.4 5.6 3.8 7.3l3.7 2.6c.3.2.8.2 1.2 0l3.7-2.6c2.4-1.7 3.8-4.4 3.8-7.3V4.4c-.2-.3-.3-.6-.6-.8M14 13h-1v1c0 .6-.4 1-1 1s-1-.4-1-1v-1h-1c-.6 0-1-.4-1-1s.4-1 1-1h1v-1c0-.6.4-1 1-1s1 .4 1 1v1h1c.6 0 1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 18c-.6 0-1 .4-1 1v2c0 .6.4 1 1 1s1-.4 1-1v-2c0-.6-.4-1-1-1m5-4c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1s1-.4 1-1v-6c0-.6-.4-1-1-1M20 2c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1s1-.4 1-1V3c0-.6-.4-1-1-1m-5 7c-.6 0-1 .4-1 1v11c0 .6.4 1 1 1s1-.4 1-1V10c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalAltThree(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 15H2c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h4c.6 0 1-.4 1-1v-6c0-.6-.4-1-1-1m8-6h-4c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1h4c.6 0 1-.4 1-1V10c0-.6-.4-1-1-1m8-8h-4c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h4c.6 0 1-.4 1-1V2c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signout(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7C5.3 2 4 3.3 4 5v6h8.6l-2.3-2.3c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l4 4c.4.4.4 1 0 1.4l-4 4c-.4.4-1 .4-1.4 0c-.4-.4-.4-1 0-1.4l2.3-2.3H4v6c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3V5c0-1.7-1.3-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancing(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.2 17.3l-2-2c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l.3.3h-2.6c-.6 0-1 .4-1 1s.4 1 1 1h2.6l-.3.3c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l2-2c.1-.1.2-.2.2-.3c0-.1.1-.2.1-.4c0-.1 0-.3-.1-.4c0-.1-.1-.2-.2-.3M8.5 17H5.9l.3-.3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-2 2c-.1.1-.2.2-.2.3c0 .1-.1.2-.1.4c0 .1 0 .3.1.4c.1.1.1.2.2.3l2 2c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-.3-.3h2.6c.6 0 1-.4 1-1s-.4-1-1-1M18 9.2c1.4 0 2.6-1.2 2.6-2.6S19.4 4 18 4c-1.4 0-2.6 1.2-2.6 2.6c0 1.4 1.2 2.6 2.6 2.6m4.7 3.7c-2.1-2.6-5.9-3-8.5-.9c-.3.3-.7.6-.9.9c-.4.6-.4 1.4.2 1.8c.2.2.5.3.8.3h1.9c.1-.3.3-.5.5-.8c1-1 2.6-1 3.5 0l.8.8h.6c.7 0 1.3-.6 1.3-1.3c.1-.3 0-.6-.2-.8m-20.5-1c-.3.3-.7.6-.9.9c-.4.6-.4 1.4.2 1.8c.2.3.5.4.8.4H3l.8-.8c1-1 2.6-1 3.5 0c.2.2.4.5.5.8h1.9c.7 0 1.3-.6 1.3-1.3c0-.3-.1-.6-.3-.8c-2.1-2.6-5.9-3-8.5-1M6 9.2c1.4 0 2.6-1.2 2.6-2.6S7.4 4 6 4C4.6 4 3.4 5.2 3.4 6.6C3.4 8 4.6 9.2 6 9.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sorting(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.7 17.8l-3-3c-.4-.4-1-.4-1.4 0l-3 3c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0l2.3-2.3l2.3 2.3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4m-4.4-7.6c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L12 8.1L9.7 5.8c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpaceKey(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 9c-.6 0-1 .4-1 1v3H4v-3c0-.6-.4-1-1-1s-1 .4-1 1v4c0 .6.4 1 1 1h18c.6 0 1-.4 1-1v-4c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareFull(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h18c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1m-1 18H4V4h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.1c.1-.5-.3-1.1-.8-1.1l-5.7-.8L12.9 3c-.1-.2-.2-.3-.4-.4c-.5-.3-1.1-.1-1.4.4L8.6 8.2L2.9 9c-.3 0-.5.1-.6.3c-.4.4-.4 1 0 1.4l4.1 4l-1 5.7c0 .2 0 .4.1.6c.3.5.9.7 1.4.4l5.1-2.7l5.1 2.7c.1.1.3.1.5.1h.2c.5-.1.9-.6.8-1.2l-1-5.7l4.1-4c.2-.1.3-.3.3-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.1c.1-.5-.3-1.1-.8-1.1l-5.7-.8L12.9 3c-.1-.2-.2-.3-.4-.4c-.5-.3-1.1-.1-1.4.4L8.6 8.2L2.9 9c-.3 0-.5.1-.6.3c-.4.4-.4 1 0 1.4l4.1 4l-1 5.7c0 .2 0 .4.1.6c.3.5.9.7 1.4.4l5.1-2.7l5.1 2.7c.1.1.3.1.5.1h.2c.5-.1.9-.6.8-1.2l-1-5.7l4.1-4c.2-.1.3-.3.3-.5m-6.2 3.5c-.2.2-.3.6-.3.9l.7 4.2l-3.8-2c-.1-.1-.3-.1-.5-.1V5.7l1.9 3.8c.1.3.4.5.8.5l4.2.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForward(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.7 7.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l3.3 3.3l-3.3 3.3c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l4-4c.4-.4.4-1 0-1.4zM16 7c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1s1-.4 1-1V8c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.8 10c-.4-1.2-1.6-2-2.8-2c-1.7 0-3 1.3-3 3c0 1.3.8 2.4 2 2.8v1.7c0 2.5-2 4.5-4.5 4.5S9 18 9 15.5v-.6c2.9-.5 5-3 5-5.9V3c0-.6-.4-1-1-1h-2c-.6 0-1 .4-1 1s.4 1 1 1h1v5c0 2.2-1.8 4-4 4s-4-1.8-4-4V4h1c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1v6c0 2.9 2.1 5.4 5 5.9v.6c0 3.6 2.9 6.5 6.5 6.5s6.5-2.9 6.5-6.5v-1.7c1.6-.5 2.4-2.2 1.8-3.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StethoscopeAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.8 10c-.4-1.2-1.6-2-2.8-2c-1.7 0-3 1.3-3 3c0 1.3.8 2.4 2 2.8v1.7c0 2.5-2 4.5-4.5 4.5S9 18 9 15.5v-1l3.1-2.5c1.2-1 1.9-2.4 1.9-3.9V3c0-.6-.4-1-1-1h-2c-.6 0-1 .4-1 1s.4 1 1 1h1v4.1c0 .9-.4 1.8-1.1 2.3L8 12.7l-2.9-2.3C4.4 9.9 4 9 4 8.1V4h1c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1v5.1c0 1.5.7 3 1.9 3.9L7 14.5v1c0 3.6 2.9 6.5 6.5 6.5s6.5-2.9 6.5-6.5v-1.7c1.6-.5 2.4-2.2 1.8-3.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4.5h4c.6 0 1-.4 1-1s-.4-1-1-1h-4c-.6 0-1 .4-1 1s.4 1 1 1m8.3 4.1l.9-.9c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-.9.9C14 4.9 10 4.9 7.1 7.2l-.9-.9c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l.9.9C3 12.1 3.6 17.1 7.1 19.8c3.5 2.7 8.5 2.1 11.2-1.4c2.3-2.9 2.3-6.9 0-9.8m-4.6 5.9c-.4.6-1 1-1.7 1c-1.1 0-2-.9-2-2c0-.7.4-1.4 1-1.7V9.5c0-.6.4-1 1-1s1 .4 1 1v2.3c1 .5 1.3 1.7.7 2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StoreSlash(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 10.6c.8.9 1.8 1.4 3 1.4v1.8l2 2v-4.3c1.2-.7 2-2 2-3.4c0-.1 0-.3-.1-.4l-2-5c-.1-.5-.5-.7-.9-.7H6.2zc-.1.1 0 0 0 0m7.7 10.7L20 18.6l-2-2l-4.8-4.8l-9.1-9.2l-1.4-1.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2 2l-1.2 3c-.1 0-.1.2-.1.3c0 1.4.8 2.7 2 3.4V21c0 .6.4 1 1 1h14c.4 0 .8-.3.9-.7l1.4 1.4c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4M10 14c-.6 0-1 .4-1 1v5H6v-8c1.2 0 2.2-.5 3-1.4c.3.3.6.6 1 .8l2.6 2.6zm8 6h-3v-3.6l3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subject(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m10 8H3c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1m8-5H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SyncExclamation(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13c.6 0 1-.4 1-1V9c0-.6-.4-1-1-1s-1 .4-1 1v3c0 .6.4 1 1 1m-1 2c0 .5.5 1 1 1s1-.5 1-1c0-.3-.1-.5-.3-.7c-.3-.3-.7-.4-1.1-.2c-.1 0-.2.1-.3.2c-.2.2-.3.4-.3.7m6-11.7C13.1 1.1 8.3 1.8 5.1 4.7V3c0-.6-.4-1-1-1s-1 .4-1 1v4.5c0 .1 0 .2.1.3v.1c.1.2.3.4.5.5c.2.1.3.1.4.1h4.5c.6 0 1-.4 1-1s-.4-1-1-1H6.2C7.7 4.9 9.8 4 12 4c4.4 0 8 3.6 8 8c0 .6.4 1 1 1s1-.4 1-1c0-3.6-1.9-6.9-5-8.7m2.9 12.2h-4.5c-.6 0-1 .4-1 1s.4 1 1 1h2.4C16.3 19.1 14.2 20 12 20c-4.4 0-8-3.6-8-8c0-.6-.4-1-1-1s-1 .4-1 1c0 5.5 4.5 10 10 10c2.6 0 5-1 6.9-2.8V21c0 .6.4 1 1 1s1-.4 1-1v-4.5c0-.6-.5-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SyncSlash(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.3 7.1l1.4-1.4l2-2c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-2 2c-3.6-3-8.9-3.1-12.6 0c-.2.1-.4.3-.6.5V3c0-.6-.4-1-1-1s-1 .4-1 1v4.5c0 .6.4 1 1 1h4.5c.6 0 1-.4 1-1s-.4-1-1-1H6.2c2.5-2.6 6.5-3.3 9.8-1.4l.9.6L5.7 16.9C4.6 15.5 4 13.8 4 12c0-.6-.4-1-1-1s-1 .4-1 1c0 2.3.8 4.6 2.3 6.3l-2 2c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l2-2l1.4-1.4zm1.6 8.4h-4.5c-.6 0-1 .4-1 1s.4 1 1 1h2.4c-2.2 2.4-5.8 3.2-8.9 1.9l-1.5 1.5c1.4.7 3 1.1 4.6 1.1c2.6 0 5-1 6.9-2.8V21c0 .6.4 1 1 1s1-.4 1-1v-4.5c0-.6-.5-1-1-1m-.5-6.6c.4 1 .6 2 .6 3.1c0 .6.4 1 1 1s1-.4 1-1c0-1.6-.4-3.2-1.1-4.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h18c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1M10 20v-4h4v4zm0-6v-4h4v4zm-6-4h4v4H4zm6-2V4h4v4zm6 2h4v4h-4zm4-2h-4V4h4zM8 4v4H4V4zM4 16h4v4H4zm12 4v-4h4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThLarge(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4c-.6 0-1 .4-1 1v16c0 .6.4 1 1 1h16c.6 0 1-.4 1-1V4c0-.6-.4-1-1-1m-1 8h-6V5h6zm-8-6v6H5V5zm-6 8h6v6H5zm8 6v-6h6v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesCircle(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m3.7 12.3c.4.4.4 1 0 1.4c-.4.4-1 .4-1.4 0L12 13.4l-2.3 2.3c-.4.4-1 .4-1.4 0c-.4-.4-.4-1 0-1.4l2.3-2.3l-2.3-2.3c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l2.3 2.3l2.3-2.3c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4L13.4 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOff(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 6.5h-9C4.5 6.5 2 9 2 12s2.5 5.5 5.5 5.5h9c3 0 5.5-2.5 5.5-5.5s-2.5-5.5-5.5-5.5m-9 8C6.1 14.5 5 13.4 5 12s1.1-2.5 2.5-2.5S10 10.6 10 12s-1.1 2.5-2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOn(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 6.5h-9C4.5 6.5 2 9 2 12s2.5 5.5 5.5 5.5h9c3 0 5.5-2.5 5.5-5.5s-2.5-5.5-5.5-5.5m0 8c-1.4 0-2.5-1.1-2.5-2.5s1.1-2.5 2.5-2.5S19 10.6 19 12s-1.1 2.5-2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToiletPaper(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.3 20.4c-1.2-1.4-1.8-3.1-1.8-4.9V8c0-3.3-2.2-6-5-6h-9c-2.8 0-5 2.7-5 6s2.2 6 5 6h3v1.5c0 2.3.8 4.5 2.2 6.2c.2.2.5.3.8.3h9c.6 0 1-.4 1-1c0-.2-.1-.5-.2-.6M6.5 9.2c-.6 0-1.1-.6-1-1.2c-.1-.6.4-1.2 1-1.2c.6.1 1.1.6 1 1.2c.1.6-.4 1.2-1 1.2M13 20c-1-1.3-1.5-2.9-1.5-4.6V7.9c0-1.4-.4-2.8-1.3-4h5.3c1.7 0 3 1.8 3 4v7.5c0 1.6.4 3.2 1.1 4.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.9 19.3l-9-15.6l-.3-.3c-.5-.3-1.1-.2-1.4.3l-9 15.6c-.2.1-.2.3-.2.5c0 .6.4 1 1 1h18c.2 0 .3 0 .5-.1c.5-.3.6-.9.4-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9H9V7c0-.8.3-1.5.9-2.1c1.2-1.2 3.1-1.2 4.2 0c.4.4.6.9.8 1.4c.1.5.7.8 1.2.7c.5-.1.9-.7.7-1.2c-.2-.9-.7-1.7-1.3-2.3c-.9-1-2.2-1.5-3.5-1.5c-2.8 0-5 2.2-5 5v2c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9H9V7c0-.8.3-1.5.9-2.1c1.2-1.2 3.1-1.2 4.2 0c.4.4.6.9.8 1.4c.1.5.7.8 1.2.7c.5-.1.9-.7.7-1.2c-.2-.9-.7-1.7-1.3-2.3c-.9-1-2.2-1.5-3.5-1.5c-2.8 0-5 2.2-5 5v2c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3m-4 8c0 .6-.4 1-1 1s-1-.4-1-1v-3c0-.6.4-1 1-1s1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 9h-2c-.6 0-1 .4-1 1s.4 1 1 1h2c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1H6c-.6 0-1-.4-1-1v-7c0-.6.4-1 1-1h2c.6 0 1-.4 1-1s-.4-1-1-1H6c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h12c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3M9.7 6.7L11 5.4V17c0 .6.4 1 1 1s1-.4 1-1V5.4l1.3 1.3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-3-3c-.4-.4-1-.4-1.4 0l-3 3c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserArrows(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.6 6.9c0 .1.1.2.2.3l2 2c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-.3-.3h4.2l-.3.3c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l2-2c.1-.1.2-.2.2-.3c0-.1.1-.2.1-.4c0-.1 0-.3-.1-.4c-.1-.1-.1-.2-.2-.3l-2-2c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l.3.3H9.9l.3-.3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-2 2c-.1.1-.2.2-.2.3c0 .1-.1.2-.1.4c0 .1 0 .3.1.4M6 14.7c1.4 0 2.6-1.2 2.6-2.6c0-1.4-1.2-2.6-2.6-2.6c-1.4 0-2.6 1.2-2.6 2.6c0 1.4 1.2 2.6 2.6 2.6m3.8 2.7c-2.6-2.1-6.4-1.7-8.5.9c-.2.3-.3.6-.3.9c0 .7.6 1.3 1.3 1.3h7.4c.5 0 1-.3 1.2-.7c.2-.4.2-1-.2-1.4c-.3-.4-.6-.7-.9-1m5.6-5.3c0 1.4 1.2 2.6 2.6 2.6s2.6-1.2 2.6-2.6c0-1.4-1.2-2.6-2.6-2.6c-1.4 0-2.6 1.2-2.6 2.6m7.3 6.3c-.3-.3-.6-.7-.9-.9c-2.6-2.1-6.4-1.7-8.5.9c-.2.2-.3.5-.3.8c0 .7.6 1.3 1.3 1.3h7.4c.5 0 1-.3 1.2-.7c.2-.5.1-1-.2-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMd(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.9 13.2c-.1 0-.1-.1-.2-.1C17.2 12 18 10.3 18 8.4v-.7l.3-2.4c.2-1.6-.9-3-2.4-3.3l-.9-.2c-2-.3-4-.3-6 0l-.8.2c-1.6.3-2.7 1.7-2.5 3.3L6 7.7v.7c0 1.8.8 3.6 2.3 4.7c-.1 0-.1.1-.2.1c-3.3 1.4-5.6 4.5-6 8.1c-.1.5.3 1 .9 1.1c.6.1 17.5 0 18 0h.1c.5-.1.9-.6.9-1.1c-.5-3.6-2.8-6.7-6.1-8.1M12 16.3l-1.7-1.7c1.1-.2 2.2-.2 3.3 0zm0-3.9c-2.2 0-3.9-1.7-4-3.9h8c-.1 2.2-1.8 3.9-4 3.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserNurse(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.2 12.9s-.1 0 0 0c-.2-.1-.3-.2-.5-.2c2.6-2 3.1-5.8 1-8.4s-5.8-3.1-8.4-1s-3.1 5.8-1 8.4c.3.4.6.7 1 1c-.1.1-.3.1-.4.2h-.1c-3.2 1.5-5.4 4.5-5.8 8c0 .5.4 1 1 1.1h18c.5-.1.9-.6.9-1.1c-.3-3.5-2.5-6.5-5.7-8M8 7.6c.2-2.2 2.2-3.8 4.3-3.6c1.9.2 3.4 1.7 3.6 3.6zm4 8.6l-1.9-1.9c1.3-.3 2.6-.3 3.9 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorSquare(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 16.2V7.8c1.2-.4 2-1.5 2-2.8c0-1.7-1.3-3-3-3c-1.3 0-2.4.8-2.8 2H7.8C7.4 2.8 6.3 2 5 2C3.3 2 2 3.3 2 5c0 1.3.8 2.4 2 2.8v8.4c-1.2.4-2 1.5-2 2.8c0 1.7 1.3 3 3 3c1.3 0 2.4-.8 2.8-2h8.4c.4 1.2 1.5 2 2.8 2c1.7 0 3-1.3 3-3c0-1.3-.8-2.4-2-2.8M16.2 18H7.8c-.3-.8-1-1.5-1.8-1.8V7.8c.8-.3 1.5-1 1.8-1.8h8.4c.3.8 1 1.5 1.8 1.8v8.4c-.8.3-1.5 1-1.8 1.8M19 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1M5 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m0 16c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m14 0c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorSquareAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18.3V5.7c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2c-.7 0-1.4.4-1.7 1H5.7c-.3-.6-1-1-1.7-1c-1.1 0-2 .9-2 2c0 .7.4 1.4 1 1.7v12.6c-.6.3-1 1-1 1.7c0 1.1.9 2 2 2c.7 0 1.4-.4 1.7-1h12.6c.3.6 1 1 1.7 1c1.1 0 2-.9 2-2c0-.7-.4-1.4-1-1.7m-2 0c-.3.2-.5.4-.7.7H5.7c-.2-.3-.4-.5-.7-.7V5.7c.3-.2.5-.4.7-.7h12.6c.2.3.4.5.7.7zM16 7H8c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1h8c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirusSlash(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.8 18.4l-4-4l-2.1-2.1l-8-8l-3-3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4L5 6.4C3.9 7.7 3.2 9.3 3.1 11H2c-.6 0-1 .4-1 1s.4 1 1 1h1.1c.2 1.7.9 3.3 1.9 4.6l-.8.8c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0l.8-.8c1.3 1.1 2.9 1.7 4.6 1.9V22c0 .6.4 1 1 1s1-.4 1-1v-1.1c1.7-.2 3.3-.9 4.6-1.9l3.7 3.7c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4zM9 16c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m.5-4c-.8 0-1.5-.7-1.5-1.5c0-.3.1-.6.3-.8l2.1 2.1c-.3.1-.6.2-.9.2m11.4 1H22c.6 0 1-.4 1-1s-.4-1-1-1h-1.1c-.2-1.7-.9-3.3-1.9-4.6l.8-.8c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-.8.8c-1.3-1.1-2.9-1.7-4.6-1.9V2c0-.6-.4-1-1-1s-1 .4-1 1v1.1c-1 .1-2 .4-2.9.8l12 12c.4-.9.7-1.9.8-2.9M15 10c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebGrid(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2h-5v20h5c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1M2 21c0 .6.4 1 1 1h11v-9H2zM2 3v8h12V2H3c-.6 0-1 .4-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebGridAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v5h20V3c0-.6-.4-1-1-1M2 21c0 .6.4 1 1 1h8V10H2zm11 1h8c.6 0 1-.4 1-1V10h-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebSection(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3v18c0 .6.4 1 1 1h11V2H3c-.6 0-1 .4-1 1m19-1h-5v20h5c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebSectionAlt(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3v18c0 .6.4 1 1 1h5V2H3c-.6 0-1 .4-1 1m19-1H10v20h11c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowGrid(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3v18c0 .6.4 1 1 1h5V2H3c-.6 0-1 .4-1 1m19-1H10v9h12V3c0-.6-.4-1-1-1M10 22h11c.6 0 1-.4 1-1v-8H10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMaximize(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v5h20V3c0-.6-.4-1-1-1M2 21c0 .6.4 1 1 1h18c.6 0 1-.4 1-1V10H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowSection(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 21c0 .6.4 1 1 1h5V10H2zm14 1h5c.6 0 1-.4 1-1V10h-6zm-6 0h4V10h-4zM21 2H3c-.6 0-1 .4-1 1v5h20V3c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrapText(children ...ElementRenderer) *UisIcon {
	return &UisIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7.2h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1m6 8H3c-.6 0-1 .4-1 1s.4 1 1 1h6c.6 0 1-.4 1-1s-.4-1-1-1m9.5-5H3c-.6 0-1 .4-1 1s.4 1 1 1h15.5c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5h-2.8c.3-.4.4-.9 0-1.3s-1-.5-1.4-.1l-2 1.7l-.1.1c-.4.4-.3 1.1.1 1.4l2 1.7c.2.1.4.2.6.2c.3 0 .6-.1.8-.4c.3-.4.3-.9 0-1.3h2.8c1.9 0 3.5-1.6 3.5-3.5s-1.6-3.5-3.5-3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
