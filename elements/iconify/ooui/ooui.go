package ooui

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "0.48.3"
	hAttr          = "1em"
	viewbox        = "0 0 20 20"
	fill           = "currentColor"
)

type OouiIcon struct {
	*SVGSVGElement
}

func Add(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 9V4H9v5H4v2h5v5h2v-5h5V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alert(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.53 2.3A1.85 1.85 0 0 0 10 1.21A1.85 1.85 0 0 0 8.48 2.3L.36 16.36C-.48 17.81.21 19 1.88 19h16.24c1.67 0 2.36-1.19 1.52-2.64zM11 16H9v-2h2zm0-4H9V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15h18v2H1zM1 3h18v2H1z"/><rect width="8" height="6" x="6" y="7" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7h7v2h-7zm0 4h7v2h-7zM1 15h18v2H1zM1 3h18v2H1z"/><rect width="8" height="6" x="1" y="7" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 7h7v2H1zm0 4h7v2H1zm0 4h18v2H1zM1 3h18v2H1z"/><rect width="8" height="6" x="11" y="7" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Appearance(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 14.5a4.25 4.25 0 0 0 4.247-4.092c.3-.21.706-.21 1.006 0a4.25 4.25 0 0 0 8.431.59L19 11a1 1 0 1 0-.182-1.984a4.252 4.252 0 0 0-7.896-.615a2.895 2.895 0 0 0-1.844 0a4.25 4.25 0 0 0-7.896.615a1 1 0 1 0-.116 1.981A4.251 4.251 0 0 0 5.25 14.5m0-2a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5M17 10.25a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowNextLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.6 3.4L14.2 9H2v2h12.2l-5.6 5.6L10 18l8-8l-8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowNextRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2 10l8 8l1.4-1.4L5.8 11H18V9H5.8l5.6-5.6L10 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPreviousLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.83 9l5.58-5.58L10 2l-8 8l8 8l1.41-1.41L5.83 11H18V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowPreviousRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 11h12.2l-5.6 5.6L10 18l8-8l-8-8l-1.4 1.4L14.2 9H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleAdd(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2zm10 10h-4v4H9v-4H5V9h4V5h2v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleCheck(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 17l-4.59-4.59L5.83 11L9 14.17l8-8V3a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleDisambiguationLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1H5c-1.1 0-2 .9-2 2v6h4.6l3.7-3.7L10 4h4v4l-1.3-1.3L9.4 10l3.3 3.3L14 12v4h-4l1.3-1.3L7.6 11H3v6c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleDisambiguationRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1c1.1 0 2 .9 2 2v14c0 1.1-.9 2-2 2H5c-1.1 0-2-.9-2-2V3c0-1.1.9-2 2-2zM8.998 4h-4v4l1.281-1.281L9.56 10l-3.28 3.281L4.997 12v4h4l-1.312-1.313L11.404 11h3.594V9h-3.594L7.717 5.281z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm0 3h5v1H5zm0 2h5v1H5zm0 2h5v1H5zm10 7H5v-1h10zm0-2H5v-1h10zm0-2H5v-1h10zm0-2h-4V4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleNotFoundLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2m-4 15H9v-2h2zm2.7-7.6a4.88 4.88 0 0 1-.3.7a2.65 2.65 0 0 1-.5.6l-.5.5a2.65 2.65 0 0 1-.6.5c-.2.2-.3.4-.5.6a1.91 1.91 0 0 0-.3.8a3.4 3.4 0 0 0-.1 1H9.1a4.87 4.87 0 0 1 .1-1.2a2.92 2.92 0 0 1 .2-.9a2.51 2.51 0 0 1 .4-.7l.6-.6a1.76 1.76 0 0 1 .5-.4c.2-.1.3-.3.4-.4l.3-.6a1.7 1.7 0 0 0 .1-.7a2.92 2.92 0 0 0-.2-.9a2.19 2.19 0 0 0-1-.9a.9.9 0 0 0-.5-.1a1.68 1.68 0 0 0-1.5.7A2.86 2.86 0 0 0 8 8.1H6.2a5.08 5.08 0 0 1 .3-1.7a3.53 3.53 0 0 1 .8-1.3a3.6 3.6 0 0 1 1.2-.8a5.08 5.08 0 0 1 1.7-.3a5.9 5.9 0 0 1 1.4.2a2.59 2.59 0 0 1 1.1.7a4.44 4.44 0 0 1 .8 1.1a4 4 0 0 1 .3 1.5a3.08 3.08 0 0 1-.1.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleNotFoundRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2m6 11h2v2H9zM6.2 7.5A4 4 0 0 1 6.5 6a4.44 4.44 0 0 1 .8-1.1a2.59 2.59 0 0 1 1.1-.7A5.9 5.9 0 0 1 9.8 4a5.08 5.08 0 0 1 1.7.3a3.6 3.6 0 0 1 1.2.8a3.53 3.53 0 0 1 .8 1.3a5.08 5.08 0 0 1 .3 1.7H12a2.86 2.86 0 0 0-.5-1.7a1.68 1.68 0 0 0-1.5-.7a.9.9 0 0 0-.5.1a2.19 2.19 0 0 0-1 .9a2.92 2.92 0 0 0-.2.9a1.7 1.7 0 0 0 .1.7l.3.6c.1.1.2.3.4.4a1.76 1.76 0 0 1 .5.4l.6.6a2.51 2.51 0 0 1 .4.7a2.92 2.92 0 0 1 .2.9a4.87 4.87 0 0 1 .1 1.2H9.1a3.4 3.4 0 0 0-.1-1a1.91 1.91 0 0 0-.3-.8c-.2-.2-.3-.4-.5-.6a2.65 2.65 0 0 1-.6-.5l-.5-.5a2.65 2.65 0 0 1-.5-.6a4.88 4.88 0 0 1-.3-.7a3.08 3.08 0 0 1-.1-.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleRedirectLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1a2 2 0 0 0-2 2v1c0 5 2 8 7 8V9l5 4l-5 4v-3c-3.18 0-5.51-.85-7-2.68V17a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleRedirectRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 17c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2v-5.7c-1.5 1.8-3.8 2.7-7 2.7v3l-5-4l5-4v3c5 0 7-3 7-8V3c0-1.1-.9-2-2-2H5c-1.1 0-2 .9-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 17a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2zM15 5h-5V4h5zm0 2h-5V6h5zm0 2h-5V8h5zM5 14h10v1H5zm0-2h10v1H5zm0-2h10v1H5zm0-6h4v5H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleSearch(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.43 14.34A5 5 0 0 1 10 15a5 5 0 1 1 3.95-2L17 16.09V3a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 1.45-.63z"/><circle cx="10" cy="10" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticlesLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0a2 2 0 0 0-2 2h9a2 2 0 0 1 2 2v12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2z"/><path fill="currentColor" d="M13 20a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2zM9 5h4v5H9zM4 5h4v1H4zm0 2h4v1H4zm0 2h4v1H4zm0 2h9v1H4zm0 2h9v1H4zm0 2h9v1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticlesRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0a2 2 0 0 1 2 2H6a2 2 0 0 0-2 2v12a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M18 5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2zm-7 5H7V5h4zm5-4h-4V5h4zm0 2h-4V7h4zm0 2h-4V9h4zm0 2H7v-1h9zm0 2H7v-1h9zm0 2H7v-1h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticlesSearchLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0a2 2 0 0 0-2 2h9a2 2 0 0 1 2 2v12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2z"/><path fill="currentColor" d="M10.8 15.6a4.6 4.7 0 0 1-2.3.6a4.6 4.7 0 1 1 3.7-1.9l2.8 3V4.9A1.9 1.9 0 0 0 13.1 3H4a1.9 1.9 0 0 0-2 1.9V18a1.9 1.9 0 0 0 1.9 2H13a1.9 1.9 0 0 0 1.4-.6z"/><circle cx="8.5" cy="11.5" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticlesSearchRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0a2 2 0 0 1 2 2H6a2 2 0 0 0-2 2v12a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M13.8 15.5a4.6 4.7 0 0 1-2.3.6a4.6 4.7 0 1 1 3.7-1.9l2.8 3V5a1.9 1.9 0 0 0-1.9-2H7a1.9 1.9 0 0 0-2 1.9V18a1.9 1.9 0 0 0 1.9 2H16a1.9 1.9 0 0 0 1.4-.6z"/><circle cx="11.5" cy="11.5" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 19.75a4.25 4.25 0 0 1-4.25-4.25V9a2.75 2.75 0 0 1 5.5 0v6h-1.5V9a1.25 1.25 0 0 0-2.5 0v6.5a2.75 2.75 0 0 0 5.5 0V4a2.25 2.25 0 0 0-4.5 0v1h-1.5V4a3.75 3.75 0 0 1 7.5 0v11.5a4.25 4.25 0 0 1-4.25 4.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7a5.38 5.38 0 0 0-4.46-4.85C11.6 1.46 11.53 0 10 0S8.4 1.46 8.46 2.15A5.38 5.38 0 0 0 4 7v6l-2 2v1h16v-1l-2-2zm-6 13a3 3 0 0 0 3-3H7a3 3 0 0 0 3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellOutline(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 2.19C14.09 2.86 16 5.2 16 8v6l2 2v1H2v-1l2-2V8c0-2.8 1.91-5.14 4.5-5.81V1.5C8.5.67 9.17 0 10 0s1.5.67 1.5 1.5zM10 4C7.79 4 6 5.79 6 8v7h8V8c0-2.21-1.79-4-4-4M8 18h4c0 1.1-.9 2-2 2s-2-.9-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bigger(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 18h-1.57a.66.66 0 0 1-.44-.13a.87.87 0 0 1-.25-.34l-1-2.77H5.3l-1 2.77a.83.83 0 0 1-.24.32a.65.65 0 0 1-.44.15H2L7 5.47h2zm-3.85-4.7L8.42 8.72A12.66 12.66 0 0 1 8 7.37q-.1.41-.21.75t-.21.6L5.85 13.3zM15 2l3 4h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Block(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 1a9 9 0 1 0 9 9a9 9 0 0 0-9-9m5 10H5V9h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldA(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 15h-7L5 19H1L8 1h4l7 18h-4Zm-6-3h5L10 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldArabAin(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.89 11.9c0 1.57 1.61 2.44 4.85 2.63l2.55-.04l.37.05c-.03.14-.29.4-.77.76l-.1.07A6.97 6.97 0 0 1 9.63 17a4.3 4.3 0 0 1-3.16-1.16a4.3 4.3 0 0 1-1.13-3.15c0-1.58.66-3 1.96-4.27v-.05l-.7-.64A1.11 1.11 0 0 1 6.33 7c0-.58.28-1.3.84-2.18C7.93 3.6 8.7 3 9.46 3c1.03 0 1.89.49 2.56 1.45c.38.56-.03.65-1.24.26c-.98-.38-1.78-.06-2.4.96l.02.1l1.31 1h.06c1.64-.57 2.82-.86 3.55-.84a5.5 5.5 0 0 0-.28.86a32.4 32.4 0 0 1-.36 1.14l-.14.44l-.45.05c-2.04.28-3.5.84-4.37 1.67a2.5 2.5 0 0 0-.83 1.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldArabDad(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 6.6c.4 0 .7-.2.9-.4a1.3 1.3 0 0 0 0-1.8c-.2-.3-.5-.4-.9-.4s-.7.1-1 .4a1.3 1.3 0 0 0 0 1.8c.3.2.6.4 1 .4M9.8 14l.4-1l.6.2a10.1 10.1 0 0 0 2.6.2c1 0 1.8 0 2.5-.2s1.3-.4 1.7-.7A3 3 0 0 0 19 10c0-.6-.1-1-.4-1.5a2 2 0 0 0-1-1c-.5-.3-1-.5-1.7-.5l-1.3.3c-.4 0-.8.3-1.2.5a8.7 8.7 0 0 0-2.3 2l-.7.8a.7.7 0 0 1-.2-.2l-.2-.3a27.5 27.5 0 0 0-.4-1.7l-.1-.3l-2.8.7l.1.3l.5 1.6l.2 1.3l-.1.6l-.4.5l-.6.2a4 4 0 0 1-1 .1l-1-.1l-.5-.5a1 1 0 0 1-.2-.6v-.4l.2-.4l.2-.8l.1-.3l-2.4-.8l-.1.3l-.5 1.2a6.7 6.7 0 0 0-.2 1.6c0 .6.1 1 .4 1.5c.2.5.5.8 1 1.1c.4.3.9.5 1.4.6a7.2 7.2 0 0 0 3.4 0a5 5 0 0 0 1.5-.7l1.1-1Zm3.4-3c.3-.5 1.1-1.4 2.3-1.5c1 0 1 .8.5 1c-.3.2-1.2.5-2.8.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldArabJeem(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.61 13.756c-.437 0-.81-.152-1.12-.457a1.497 1.497 0 0 1-.463-1.106c0-.445.154-.82.463-1.125c.31-.318.683-.477 1.12-.477c.439 0 .806.159 1.102.477c.309.305.463.68.463 1.125c0 .432-.154.8-.463 1.106a1.473 1.473 0 0 1-1.101.457ZM11.109 19a9.788 9.788 0 0 1-2.858-.4a6.695 6.695 0 0 1-2.26-1.183a5.195 5.195 0 0 1-1.468-1.887C4.174 14.78 4 13.922 4 12.956c0-.954.155-1.818.464-2.594a6.822 6.822 0 0 1 1.294-2.078A8.82 8.82 0 0 1 7.67 6.7c.734-.457 1.52-.832 2.357-1.125c.85-.292.914-.394 1.816-.534c-.4-.101-.838-.197-1.314-.286a24.347 24.347 0 0 0-1.487-.228a12.925 12.925 0 0 0-1.584-.096c-.515 0-1.018.026-1.507.076a8.809 8.809 0 0 0-1.236.21L4.232 1.38c.386-.101.805-.19 1.255-.267A8.47 8.47 0 0 1 6.878 1c.799 0 1.526.064 2.183.19c.67.128 1.281.287 1.835.477c.567.178 1.088.363 1.565.553c.489.191.946.35 1.371.477c.438.127.856.19 1.256.19h1.004l.387 3.662c-.4 0-.902.057-1.507.171a15.08 15.08 0 0 0-1.913.515c-.682.229-1.358.515-2.028.858a9.628 9.628 0 0 0-1.796 1.163a5.81 5.81 0 0 0-1.314 1.45a3.334 3.334 0 0 0-.483 1.754c0 .584.097 1.074.29 1.468c.206.407.483.73.83.972c.361.255.78.439 1.256.553c.49.115 1.018.172 1.584.172c.76 0 1.526-.076 2.299-.229a12.893 12.893 0 0 0 2.337-.705L17 17.684c-1.056.47-2.093.807-3.11 1.01c-1.017.204-1.945.306-2.782.306Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldArmnTo(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 18V9c0-1.34.29-2.61.86-3.64a6.13 6.13 0 0 1 2.55-2.45C5.54 2.3 7.38 2 9 2c2.19 0 3.47.43 4.73 1.28A5.3 5.3 0 0 1 16 7h3v3h-3c0 1.78-.27 2.73-1.16 3.69a4.55 4.55 0 0 1-3.48 1.44a5.5 5.5 0 0 1-2.5-.53a3.94 3.94 0 0 1-1.62-1.46a4.1 4.1 0 0 1-.58-2.17a3.6 3.6 0 0 1 1.26-2.94c.85-.7 2.1-1.06 3.74-1.06L13 7c0-.5-.76-1.03-1.47-1.4A5.2 5.2 0 0 0 9 5c-1.44 0-3 .04-4 1c-.76.74-1 1.61-1 3v9zm12-8.42h-1.31c-.53 0-.97.12-1.3.35c-.34.22-.51.56-.51 1.02c0 .44.1.82.3 1.13c.23.3.62.44 1.2.44c.51 0 .92-.2 1.23-.62c.32-.41.39-1.19.39-2.32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldB(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.9 1c2.4 0 4.2.3 5.4 1c1.3.7 1.9 1.9 1.9 3.6c0 1-.3 1.9-.7 2.6a3 3 0 0 1-2 1.3a4.8 4.8 0 0 1 1.6.7c.4.3.8.8 1.1 1.3a5 5 0 0 1 .4 2.3a4.6 4.6 0 0 1-1.7 3.8A7.6 7.6 0 0 1 10 19H3.3V1zm.4 7.1c1.1 0 1.9-.1 2.3-.5c.5-.3.7-.9.7-1.5c0-.7-.3-1.2-.8-1.5c-.5-.3-1.3-.5-2.4-.5h-2v4zm-2.2 3V16h2.5c1.1 0 2-.3 2.4-.7c.5-.5.7-1 .7-1.8a2 2 0 0 0-.7-1.6c-.5-.4-1.3-.6-2.5-.6H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldCyrlBe(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 19V1h12v3H8v4h1c2 0 3.62.14 4.72.61a5.27 5.27 0 0 1 2.48 1.95c.53.82.8 1.88.8 2.94c0 1.78-.78 3-2 4c-1.2.97-3.35 1.5-6 1.5zm5-3a6.7 6.7 0 0 0 3.01-.68c.7-.37.99-.9.99-1.82c0-.96-.33-1.71-1.07-2.01A7.61 7.61 0 0 0 9 11H8v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldCyrlPalochka(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 19H6v-2l2-1V4L6 3V1h8v2l-2 1v12l2 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldCyrlTe(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19H8V4H3V1h14v3h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldCyrlZhe(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 10L1 3h4l3 7V3h4v7l3-7h4l-3.5 7l3.5 7h-4l-3-7v7H8v-7l-3 7H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldF(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 19H5V1h10v3H9v5h5v3H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldG(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 9h7v9a20.41 20.41 0 0 1-6.5 1a9.67 9.67 0 0 1-4.64-1.03a6.8 6.8 0 0 1-2.89-3.04A11.21 11.21 0 0 1 2 10a9.9 9.9 0 0 1 1.1-4.78A7.7 7.7 0 0 1 6.35 2.1C7.77 1.37 9.95 1 12 1c.97 0 1.48.1 2.42.3c.93.19 1.81.37 2.58.7l-1 3c-.56-.28-1.5-.46-2.23-.64A9.24 9.24 0 0 0 11.53 4c-1.12 0-2.1.34-2.94.83A5.05 5.05 0 0 0 6.66 6.9A6.72 6.72 0 0 0 6 10a9 9 0 0 0 .48 3.09c.32.88.84 1.58 1.53 2.08c.7.5 1.62.83 2.75.83c.37 0 .8.02 1.08 0c.3-.03.66 0 1.16-.34V12h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldGeorMan(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 19a6.4 6.4 0 0 1-3.15-.74a5.13 5.13 0 0 1-2.12-2.15A6.2 6.2 0 0 1 4 13c0-2.02.46-3.84 1.36-4.87a4.54 4.54 0 0 1 3.58-1.55c.61 0 1.15.07 1.6.21a22 22 0 0 1 1.15.41C12.5 7.5 13 8 13 8V6c0-1-.5-1.5-.5-1.5s-.5-.5-.97-.73c0 0-.53-.27-1.53-.3c0 0-.33-.02-.78.13c-.42.14-.8.3-1.02.65c-.16.25-.2.44-.2.75H5c-.02-.15 0-.4 0-.53c0-.77.02-1.2.5-1.76c.5-.56 1.18-.98 2-1.27C8.35 1.14 8.97 1 10 1c1.18 0 2.5.21 3.34.65A4.54 4.54 0 0 1 16 6v7c0 2-.51 3.25-1.53 4.36c-1 1.1-2.57 1.64-4.47 1.64m0-3c.88 0 1.65-.06 2.07-.65c.42-.6.64-1.37.64-2.3v-.47c0-1-.2-1.8-.57-2.42c-.38-.61-1.07-.92-2.1-.92a2.1 2.1 0 0 0-1.83.94A5.44 5.44 0 0 0 7.58 13c0 1 .18 1.6.56 2.26c.4.65.93.74 1.86.74"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldL(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 19V1h4v15h7v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldN(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 19h-4L5 5v14H2V1h5l8 13c-.02-.84 0-1 0-2V1h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldV(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 19l6-18h-4l-4 14L6 1H2l6 18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2a7.65 7.65 0 0 0-5 2a7.65 7.65 0 0 0-5-2H1v15h4a7.65 7.65 0 0 1 5 2a7.65 7.65 0 0 1 5-2h4V2zm2.5 13.5H14a4.38 4.38 0 0 0-3 1V5s1-1.5 4-1.5h2.5z"/><path fill="currentColor" d="M9 3.5h2v1H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 2v15h4a7.65 7.65 0 0 1 5 2a7.65 7.65 0 0 1 5-2h4V2h-4a7.65 7.65 0 0 0-5 2a7.65 7.65 0 0 0-5-2zm1.5 1.5H5C8 3.5 9 5 9 5v11.5a4.38 4.38 0 0 0-3-1H2.5z"/><path fill="currentColor" d="M9 3.5h2v1H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1a2 2 0 0 0-2 2v16l7-5l7 5V3a2 2 0 0 0-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkOutline(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1a2 2 0 0 0-2 2v16l7-5l7 5V3a2 2 0 0 0-2-2zm10 14.25l-5-3.5l-5 3.5V3h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bright(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.07 7.07V2.93h-4.14L10 0L7.07 2.93H2.93v4.14L0 10l2.93 2.93v4.14h4.14L10 20l2.93-2.93h4.14v-4.14L20 10zM10 16a6 6 0 1 1 6-6a6 6 0 0 1-6 6"/><circle cx="10" cy="10" r="4.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrowserLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm2 1.5A1.5 1.5 0 1 1 2.5 5A1.5 1.5 0 0 1 4 3.5M18 16H2V8h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrowserRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 16c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2H2C.9 2 0 2.9 0 4zM17.5 5c0 .8-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5s1.5.7 1.5 1.5M2 8h16v8H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3V1h-2v2H7V1H5v2H2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm3 14H2V8h16zm-2-6h-4v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 16c-4.455 0-6.685-5.386-3.535-8.535C9.615 4.315 15 6.545 15 11a5 5 0 0 1-5 5M6.42 2.56l-.67.64c-.37.357-.865.808-1.38.81H2C.914 4 0 4.712 0 5.76v10.48C0 17.27 1 18 2 18h16c1 0 2-.716 2-1.76V5.76C20 4.723 19 4 18 4h-2.37c-.515-.002-1.01-.453-1.38-.81l-.67-.64A2 2 0 0 0 12.2 2H7.8a2 2 0 0 0-1.38.56"/><circle cx="10" cy="11" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0a10 10 0 1 0 10 10A10 10 0 0 0 10 0M2 10a8 8 0 0 1 1.69-4.9L14.9 16.31A8 8 0 0 1 2 10m14.31 4.9L5.1 3.69A8 8 0 0 1 16.31 14.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3H1v16h18v-2H3z"/><path fill="currentColor" d="M11 11L8 9l-4 4v3h14V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 14.2L2.8 10l-1.4 1.4L7 17L19 5l-1.4-1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckAll(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.29 12.71l1.42-1.42l2.22 2.22l8.3-10.14l1.54 1.26l-9.7 11.86zM12 10h5v2h-5zm-3 4h5v2H9zm6-8h5v2h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clear(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0a10 10 0 1 0 10 10A10 10 0 0 0 10 0m5.66 14.24l-1.41 1.41L10 11.41l-4.24 4.25l-1.42-1.42L8.59 10L4.34 5.76l1.42-1.42L10 8.59l4.24-4.24l1.41 1.41L11.41 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0a10 10 0 1 0 10 10A10 10 0 0 0 10 0m2.5 14.5L9 11V4h2v6l3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.3 2.9l12.8 12.8l-1.4 1.4L2.9 4.3z"/><path fill="currentColor" d="M17.1 4.3L4.3 17.1l-1.4-1.4L15.7 2.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path id="oouiCode0" fill="currentColor" d="M1 10.08V8.92h1.15c1.15 0 1.15 0 1.15-1.15V5a7.42 7.42 0 0 1 .09-1.3a2 2 0 0 1 .3-.7a1.84 1.84 0 0 1 .93-.68A6.44 6.44 0 0 1 6.74 2h1.18v1.15h-.86A1.32 1.32 0 0 0 6 3.62a1.71 1.71 0 0 0-.36 1.23V7a3.22 3.22 0 0 1-.28 1.72a2 2 0 0 1-1.26.77a2.15 2.15 0 0 1 1.26.79A3.26 3.26 0 0 1 5.62 12v3.15A1.67 1.67 0 0 0 6 16.37a1.31 1.31 0 0 0 1.08.47h.87V18H6.74a6.3 6.3 0 0 1-2.12-.29a1.82 1.82 0 0 1-.93-.71a1.94 1.94 0 0 1-.3-.72A7.46 7.46 0 0 1 3.31 15v-3.77c0-1.15 0-1.15-1.15-1.15zm18 0V8.92h-1.15c-1.15 0-1.15 0-1.15-1.15V5a7.42 7.42 0 0 0-.08-1.32a2 2 0 0 0-.3-.73a1.84 1.84 0 0 0-.93-.68A6.44 6.44 0 0 0 13.26 2h-1.18v1.15h.87a1.32 1.32 0 0 1 1.05.47a1.71 1.71 0 0 1 .36 1.23V7a3.22 3.22 0 0 0 .28 1.72a2 2 0 0 0 1.26.77a2.15 2.15 0 0 0-1.26.79a3.26 3.26 0 0 0-.26 1.72v3.15a1.67 1.67 0 0 1-.38 1.22a1.31 1.31 0 0 1-1.08.47h-.87V18h1.19a6.3 6.3 0 0 0 2.12-.29a1.82 1.82 0 0 0 .93-.68a1.94 1.94 0 0 0 .3-.72a7.46 7.46 0 0 0 .1-1.31v-3.77c0-1.15 0-1.15 1.15-1.15z"/><use href="#oouiCode0" transform="matrix(-1 0 0 1 20 0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collapse(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.5 15.25l7.5-7.5l7.5 7.5l1.5-1.5l-9-9l-9 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h8v2h2V3c0-1.1-.895-2-2-2H3c-1.1 0-2 .895-2 2v8c0 1.1.895 2 2 2h2v-2H3z"/><path fill="currentColor" d="M9 9h8v8H9zm0-2c-1.1 0-2 .895-2 2v8c0 1.1.895 2 2 2h8c1.1 0 2-.895 2-2V9c0-1.1-.895-2-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3H9v2H7V3c0-1.1.895-2 2-2h8c1.1 0 2 .895 2 2v8c0 1.1-.895 2-2 2h-2v-2h2z"/><path fill="currentColor" d="M3 9v8h8V9zm8-2c1.1 0 2 .895 2 2v8c0 1.1-.895 2-2 2H3c-1.1 0-2-.895-2-2V9c0-1.1.895-2 2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CutLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.97 0c-1.66 0-3 1.34-3 3v3l-2.4-1.5a3.003 3.003 0 0 0-3 5.2a3.003 3.003 0 0 0 4.452-2.051l.952.55v6.8h2v-5.65l4.01 2.32l.988-1.73l-5-2.94v-1.17a2.996 2.996 0 0 0 4-2.829c0-1.66-1.34-3-3-3zM9 3a1 1 0 0 1 2 0a1 1 0 0 1-2 0M2 7a1 1 0 0 1 2 0a1 1 0 0 1-2 0m15 12h-2v-2h2V9h-3V7h3c1.1 0 2 .895 2 2v8c0 1.1-.895 2-2 2m-6 0h2v-2h-2zm-4-2c0 1.1.895 2 2 2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CutRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c1.66 0 3 1.34 3 3v3l2.4-1.5a3.003 3.003 0 0 1 3 5.2a3.003 3.003 0 0 1-4.452-2.051l-.952.55v6.8h-2v-5.65l-4.01 2.32l-.988-1.73l5-2.94v-1.17a2.996 2.996 0 0 1-4-2.829c0-1.66 1.34-3 3-3zM9 3a1 1 0 0 0 2 0a1 1 0 0 0-2 0m7 4a1 1 0 0 0 2 0a1 1 0 0 0-2 0M2.97 19h2v-2h-2V9h3V7h-3c-1.1 0-2 .895-2 2v8c0 1.1.895 2 2 2m6 0h-2v-2h2zm4-2c0 1.1-.895 2-2 2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 1h12l2 2v2H2V3zM2 8h16v4H2zm16 9v-2H2v2l2 2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Die(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm2 16a2 2 0 1 1 2-2a2 2 0 0 1-2 2M5 7a2 2 0 1 1 2-2a2 2 0 0 1-2 2m5 5a2 2 0 1 1 2-2a2 2 0 0 1-2 2m5 5a2 2 0 1 1 2-2a2 2 0 0 1-2 2m0-10a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleChevronEndLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2L9.7 3.3l6.6 6.7l-6.6 6.7L11 18l8-8zM2.5 2L1 3.3L7.8 10l-6.7 6.7L2.5 18l8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleChevronEndRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 2l1.3 1.3L3.7 10l6.6 6.7L9 18l-8-8zm8.5 0L19 3.3L12.2 10l6.7 6.7l-1.4 1.3l-8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleChevronStartLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 2l1.3 1.3L3.7 10l6.6 6.7L9 18l-8-8zm8.5 0L19 3.3L12.2 10l6.7 6.7l-1.4 1.3l-8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleChevronStartRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2L9.7 3.3l6.6 6.7l-6.6 6.7L11 18l8-8zM2.5 2L1 3.3L7.8 10l-6.7 6.7L2.5 18l8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownTriangle(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 15L2 5h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 12v5H3v-5H1v5a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5z"/><path fill="currentColor" d="M15 9h-4V1H9v8H5l5 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Draggable(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 11h16v2H2zm0-4h16v2H2zm11 8H7l3 3zM7 5h6l-3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.77 8l1.94-2a1 1 0 0 0 0-1.41l-3.34-3.3a1 1 0 0 0-1.41 0L12 3.23zM1 14.25V19h4.75l9.96-9.96l-4.75-4.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditLock(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12a2 2 0 0 1-2-2V5.25l-9 9V19h4.75l7-7zm7-8h-.5V2.5a2.5 2.5 0 0 0-5 0V4H13a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1m-3 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m1.5-4h-3V2.75C14.5 2 14.5 1 16 1s1.5 1 1.5 1.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditUndoLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 14.25V19h4.75l8.33-8.33l-5.27-4.23zM13 2.86V0L8 4l5 4V5h.86c2.29 0 4 1.43 4 4.29H20a6.51 6.51 0 0 0-6.14-6.43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditUndoRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 15.25V20h4.75l8.33-8.33l-5.27-4.23z"/><path fill="currentColor" d="M13 2.86V0l5 4l-5 4V5h-.86c-2.28 0-4 1.43-4 4.29H6a6.51 6.51 0 0 1 6.14-6.43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ellipsis(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="10" cy="10" r="2" fill="currentColor"/><circle cx="3" cy="10" r="2" fill="currentColor"/><circle cx="17" cy="10" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorIcon(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.728 1H6.272L1 6.272v7.456L6.272 19h7.456L19 13.728V6.272zM11 15H9v-2h2zm0-4H9V5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExitFullscreen(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7V1H5v4H1v2zM5 19h2v-6H1v2h4zm10-4h4v-2h-6v6h2zm0-8h4V5h-4V1h-2v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.5 4.75l-7.5 7.5l-7.5-7.5L1 6.25l9 9l9-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 14.5a4.5 4.5 0 1 1 4.5-4.5a4.5 4.5 0 0 1-4.5 4.5M10 3C3 3 0 10 0 10s3 7 10 7s10-7 10-7s-3-7-10-7"/><circle cx="10" cy="10" r="2.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosed(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.49 9.94A2.5 2.5 0 0 0 10 7.5z"/><path fill="currentColor" d="M8.2 5.9a4.38 4.38 0 0 1 1.8-.4a4.5 4.5 0 0 1 4.5 4.5a4.34 4.34 0 0 1-.29 1.55L17 14.14A14 14 0 0 0 20 10s-3-7-10-7a9.63 9.63 0 0 0-4 .85zM2 2L1 3l2.55 2.4A13.89 13.89 0 0 0 0 10s3 7 10 7a9.67 9.67 0 0 0 4.64-1.16L18 19l1-1zm8 12.5A4.5 4.5 0 0 1 5.5 10a4.45 4.45 0 0 1 .6-2.2l1.53 1.44a2.47 2.47 0 0 0-.13.76a2.49 2.49 0 0 0 3.41 2.32l1.54 1.45a4.47 4.47 0 0 1-2.45.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FeedbackLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 16L2 12a3.83 3.83 0 0 1-1-2.5A3.83 3.83 0 0 1 2 7l17-4z"/><rect width="4" height="8" x="4" y="9" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FeedbackRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7c.6.7 1 1.6 1 2.5c0 .9-.4 1.8-1 2.5L1 16V3z"/><rect width="4" height="8" x="12" y="9" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6L3 1v18h2v-6.87z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3 6l14-5v18h-2v-6.87z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlaceholderLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2H2a2 2 0 0 0-2 2v2h12z"/><rect width="20" height="14" y="4" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlaceholderRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6h12V4a2 2 0 0 0-2-2h-6z"/><rect width="20" height="14" y="4" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullScreen(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 1v6h2V3h4V1zm2 12H1v6h6v-2H3zm14 4h-4v2h6v-6h-2zm0-16h-4v2h4v4h2V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Function(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.263 0a4 4 0 0 0-3.846 2.901L8.246 7H5v2h2.674L6.66 12.55A2 2 0 0 1 4.737 14h-.131a2 2 0 0 1-1.665-.89l-.664-.997l-1.664 1.11l.664.996A4 4 0 0 0 4.606 16h.131a4 4 0 0 0 3.846-2.901L9.754 9H13V7h-2.674l1.014-3.55A2 2 0 0 1 13.263 2h.132a2 2 0 0 1 1.664.89l.664.997l1.664-1.11l-.664-.996A4 4 0 0 0 13.395 0zm6.151 18l-2.5-2.5l2.5-2.5L18 11.586l-2.5 2.5l-2.5-2.5L11.586 13l2.5 2.5l-2.5 2.5L13 19.414l2.5-2.5l2.5 2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunctionArgumentLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10a8 8 0 0 1-14.93 4H.832C2.375 17.532 5.9 20 10 20c5.523 0 10-4.477 10-10S15.523 0 10 0C5.9 0 2.375 2.468.832 6H3.07A8 8 0 0 1 18 10M2.062 11H8v4l6-5l-6-5v4H0v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunctionArgumentRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 10a8 8 0 0 1 14.93-4h2.238A10.002 10.002 0 0 0 10 0C4.477 0 0 4.477 0 10s4.477 10 10 10c4.1 0 7.625-2.468 9.168-6H16.93A8 8 0 0 1 2 10m15.938-1H12V5l-6 5l6 5v-4h8V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnelLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 13L1 1h18z"/><path fill="currentColor" d="M8 9v8l4 2V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnelRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 1h18l-9 12z"/><path fill="currentColor" d="m8 19l4-2V9H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.2 17.94c1.26-2 2-4.45 2.14-7.06h3.86a8.26 8.26 0 0 1-6 7.06M1.8 10.88h3.86c.14 2.6.88 5.06 2.14 7.06a8.26 8.26 0 0 1-6-7.06m6-8.82c-1.26 2-2 4.45-2.14 7.07H1.8a8.26 8.26 0 0 1 6-7.07m4.79 8.82A12.5 12.5 0 0 1 10 18a12.51 12.51 0 0 1-2.59-7.13zM7.4 9.13A12.51 12.51 0 0 1 10 1.99a12.5 12.5 0 0 1 2.59 7.14zm10.8 0h-3.87a14.79 14.79 0 0 0-2.14-7.07a8.26 8.26 0 0 1 6 7.07M10 0a10 10 0 1 0 0 20a10 10 0 0 0 0-20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HalfBrightLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6.67V3h-4.2L9.87.07L6.94 3H3v3.67L.07 9.6L3 12.53V17h3.94l2.93 2.93L12.8 17H17v-4.47l2.93-2.93zm-7 8.93v-12a6.21 6.21 0 0 1 6 6a6.21 6.21 0 0 1-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HalfBrightRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.1 9.6L3 12.5V17h4.2l2.9 2.9L13 17h4v-4.5l2.9-2.9L17 6.7V3h-3.9L10.2.1L7.2 3H3v3.7zm3.9 0c.1-3.3 2.7-5.9 6-6v12c-3.3-.1-5.9-2.7-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HalfStarLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 7h-7L10 .5L7 7H0l5.46 5.47l-1.64 7l6.18-3.7l6.18 3.73l-1.63-7zm-10 6.9V4.6l1.9 3.9h4.6l-3.73 3.4l1 4.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HalfStarRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.4 12.5l-1.6 7l6.2-3.7l6.2 3.7l-1.6-7L20 7h-7L10 .5L7 7H0zm.8 3.7l1-4.3l-3.7-3.4h4.6L10 4.6v9.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hand(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4.6V17c0 1.9-.5 3-2.4 3H9.5c-.9 0-1.8-.4-2.4-1l-4.6-5l-.5-1c0-1 .5-1 .5-1c.3 0 .6 0 1 .2L7 14V3.3C7 2.6 7.3 2 8 2c.6 0 1 .7 1 1.4V9h1V1.2c0-.6.3-1.2 1-1.2s1 .6 1 1.3V9h1V2c0-.7.3-1.3 1-1.3s1 .6 1 1.3v7h1V4.6c0-.7.3-1.3 1-1.3s1 .6 1 1.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.75 1A5.24 5.24 0 0 0 10 4A5.24 5.24 0 0 0 0 6.25C0 11.75 10 19 10 19s10-7.25 10-12.75A5.25 5.25 0 0 0 14.75 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.06 1C13 1 15 2.89 15 5.53a4.59 4.59 0 0 1-2.29 4.08c-1.42.92-1.82 1.53-1.82 2.71V13H8.38v-.81a3.84 3.84 0 0 1 2-3.84c1.34-.9 1.79-1.53 1.79-2.71a2.1 2.1 0 0 0-2.08-2.14h-.17a2.3 2.3 0 0 0-2.38 2.22v.17H5A4.71 4.71 0 0 1 9.51 1a5 5 0 0 1 .55 0"/><circle cx="10" cy="17" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpNoticeLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0a10 10 0 1 0 10 10A10 10 0 0 0 10 0m1 16H9v-2h2zm2.71-7.6a2.64 2.64 0 0 1-.33.74a3.16 3.16 0 0 1-.48.55l-.54.48c-.21.18-.41.35-.58.52a2.54 2.54 0 0 0-.47.56A2.3 2.3 0 0 0 11 12a3.79 3.79 0 0 0-.11 1H9.08a8.9 8.9 0 0 1 .07-1.25a3.28 3.28 0 0 1 .25-.9a2.79 2.79 0 0 1 .41-.67a4 4 0 0 1 .58-.58c.17-.16.34-.3.51-.44a3 3 0 0 0 .43-.44a1.83 1.83 0 0 0 .3-.55a2 2 0 0 0 .11-.72a2.06 2.06 0 0 0-.17-.86a1.71 1.71 0 0 0-1-.9a1.7 1.7 0 0 0-.5-.1a1.77 1.77 0 0 0-1.53.68a3 3 0 0 0-.5 1.82H6.16a4.74 4.74 0 0 1 .28-1.68a3.56 3.56 0 0 1 .8-1.29a3.88 3.88 0 0 1 1.28-.83A4.59 4.59 0 0 1 10.18 4a4.44 4.44 0 0 1 1.44.23a3.51 3.51 0 0 1 1.15.65a3.08 3.08 0 0 1 .78 1.06a3.54 3.54 0 0 1 .29 1.45a3.39 3.39 0 0 1-.13 1.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpNoticeRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 10A10 10 0 1 0 10 0A10 10 0 0 0 0 10m9 4h2v2H9zM6.16 7.39a3.54 3.54 0 0 1 .29-1.45a3.08 3.08 0 0 1 .78-1.06a3.51 3.51 0 0 1 1.15-.65A4.44 4.44 0 0 1 9.82 4a4.59 4.59 0 0 1 1.66.29a3.88 3.88 0 0 1 1.28.83a3.56 3.56 0 0 1 .8 1.29a4.74 4.74 0 0 1 .28 1.68h-1.91a3 3 0 0 0-.5-1.82a1.77 1.77 0 0 0-1.53-.68a1.7 1.7 0 0 0-.5.1a1.71 1.71 0 0 0-1 .9a2.06 2.06 0 0 0-.17.86a2 2 0 0 0 .11.72a1.83 1.83 0 0 0 .3.55a3 3 0 0 0 .43.44c.17.14.34.28.51.44a4 4 0 0 1 .58.58a2.79 2.79 0 0 1 .41.67a3.28 3.28 0 0 1 .25.9a8.9 8.9 0 0 1 .1 1.25H9.11A3.79 3.79 0 0 0 9 12a2.3 2.3 0 0 0-.31-.73a2.54 2.54 0 0 0-.47-.56c-.17-.17-.37-.34-.58-.52l-.54-.5a3.16 3.16 0 0 1-.48-.55a2.64 2.64 0 0 1-.33-.74a3.39 3.39 0 0 1-.13-1.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.49 1.01A4.72 4.72 0 0 1 15 5.9h-2.54v-.17a2.31 2.31 0 0 0-2.38-2.22h-.19a2.1 2.1 0 0 0-2.06 2.14c0 1.18.45 1.81 1.79 2.71a3.86 3.86 0 0 1 2 3.84v.81H9.11v-.68c0-1.18-.4-1.79-1.82-2.71A4.59 4.59 0 0 1 5 5.54C5 2.9 7 1.01 9.94 1.01a5 5 0 0 1 .55 0"/><circle cx="10" cy="17" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hieroglyph(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 11h-3.75l2.55-3.4a4.75 4.75 0 1 0-7.6 0L8.75 11H5v2h4v7h2v-7h4zM7.54 3.52A2.75 2.75 0 1 1 12.2 6.4L10 9.33L7.8 6.4a2.69 2.69 0 0 1-.26-2.88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Highlight(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.14 2.27a1 1 0 0 0-1.41 0l-10 10a1 1 0 0 0 0 1.41L4 14l-3 4h5l1-1l.29.29a1 1 0 0 0 1.41 0l10-10a1 1 0 0 0 .03-1.43zM7 15l-2-2l9-9l2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 6v5h.06l2.48 2.47l1.41-1.41L11 10.11V6z"/><path fill="currentColor" d="M10 1a9 9 0 0 0-7.85 13.35L.5 16H6v-5.5l-2.38 2.38A7 7 0 1 1 10 17v2a9 9 0 0 0 0-18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 1L0 10h3v9h4v-4.6c0-1.47 1.31-2.66 3-2.66s3 1.19 3 2.66V19h4v-9h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm-.17 13l4.09-5.25l2.92 3.51L12.92 8l5.25 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageAddLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 17H2l3.5-4.5l2.5 3l3.5-4.5l.5.67V8H8V6H2a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-6h-5.75z"/><path fill="currentColor" d="M16 4V0h-2v4h-4v2h4v4h2V6h4V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageAddRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6v2H8v4H2v6a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zM3.83 17l3.55-4.5l2.52 3l3.55-4.5L18 17zM4 10h2V6h4V4H6V0H4v4H0v2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageBroken(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.67 9.47L20 12.13v4.09A2 2 0 0 1 17.78 18H2.22A2 2 0 0 1 0 16.22v-5.86L3.33 13l4.45-3.53L12.22 13z"/><path fill="currentColor" d="M20 9.64L16.67 7l-4.44 3.56L7.78 7l-4.45 3.53L0 7.87V3.78A2 2 0 0 1 2.22 2h15.56A2 2 0 0 1 20 3.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageGallery(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zm0 11l3.5-4.5l2.5 3l3.5-4.5l4.5 6zM16 2a2 2 0 0 1 2 2H2a2 2 0 0 1 2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageLayoutBasic(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3v14h18V3zm17 13H2V4h16z"/><path fill="currentColor" d="M8.58 14h.81l3.11-4l3 4H17l-4.5-6L9 12.51L6.5 9.5L3 14h1.56l1.94-2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageLayoutFrame(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm0 15a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M17 4H3v12h14zM5 13l2.5-3l2 2L12 9l3 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageLayoutFrameless(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H1v14h18zM3 14l3.5-4.5l2.5 3L12.5 8l4.5 6z"/><path fill="currentColor" d="M19 5H1V3h18zm0 12H1v-2h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageLayoutThumbnail(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm0 15a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M17 4H3v10h14zM5 12l2.5-3l2 2L12 8l3 4zm-1 3h12v1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageLockLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 17H2l3.5-4.5l2.5 3l3-3.81A2 2 0 0 1 10 10V6H2a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-6h-5.75z"/><path fill="currentColor" d="M19 4h-.5V2.5a2.5 2.5 0 0 0-5 0V4H13a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1m-3 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m1.5-4h-3V2.75C14.5 2 14.5 1 16 1s1.5 1 1.5 1.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageLockRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5a1 1 0 0 0-1-1h-.5V2.5A2.45 2.45 0 0 0 4 0a2.45 2.45 0 0 0-2.5 2.5V4H1a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1zM4 8a1 1 0 1 1 1-1a1 1 0 0 1-1 1m1.5-4h-3V2.75C2.5 2 2.5 1 4 1s1.5 1 1.5 1.75zM10 6v4a2 2 0 0 1-2 2H2v6a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zM4 17l3.54-4.5l2.53 3l3.54-4.5l4.56 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 16h18v2H1zm8-9h10v2H9zm0 4h10v2H9zM1 2h18v2H1zm5 8l-5 4V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 16h18v2H1zm0-9h10v2H1zm0 4h10v2H1zm0-9h18v2H1zm18 4v8l-5-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 10a6 6 0 1 0 12 0a6 6 0 0 0-12 0m6-8a8 8 0 1 1 0 16a8 8 0 0 1 0-16m1 7v5H9V9zm0-1V6H9v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoFilled(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0C4.477 0 0 4.477 0 10s4.477 10 10 10s10-4.477 10-10S15.523 0 10 0M9 5h2v2H9zm0 4h2v6H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstanceLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.703 1.845L4.641 7.65a4.115 4.115 0 0 0-.099.148L3 10l1.542 2.203c.032.05.065.1.1.148l2.914 4.164l1.147 1.64a2.002 2.002 0 0 1-2.783-.496l-4.558-6.512a2 2 0 0 1 0-2.294L5.92 2.341a2.002 2.002 0 0 1 2.783-.496m-2.341 9.302l4.558 6.512a2 2 0 0 0 3.277 0l4.559-6.512a2 2 0 0 0 0-2.294l-4.559-6.512a2 2 0 0 0-3.277 0L6.362 8.853a2 2 0 0 0 0 2.294m7.835-8.806L12.56 3.488L17.117 10l-4.558 6.512l-1.639 1.147l1.639-1.147L8 10l4.559-6.512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstanceRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.414 18.154l4.062-5.803l.1-.148L17.116 10l-1.542-2.203a4.21 4.21 0 0 0-.099-.148l-2.914-4.164l-1.148-1.64a2.002 2.002 0 0 1 2.783.496l4.559 6.512a2 2 0 0 1 0 2.294l-4.559 6.512a2.002 2.002 0 0 1-2.783.495Zm2.342-9.301L9.197 2.341a2 2 0 0 0-3.277 0L1.362 8.853a2 2 0 0 0 0 2.294l4.558 6.512a2 2 0 0 0 3.277 0l4.559-6.512a2 2 0 0 0 0-2.294M5.92 17.659l1.639-1.147L3 10l4.56-6.513l1.638-1.146L7.56 3.488L12.117 10L7.56 16.512L5.92 17.66Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicA(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3 19l9.992-18H16l2 18h-3l-.4-5H8.5L6 19zm7-8h4.5c-.051-.69-.483-6.429-.5-7c-.255.588-3.693 6.361-4 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicArabKehehJeem(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.7 11.8h-1l-1.1-.5a2 2 0 0 1-.7-1V8.8a10.7 10.7 0 0 0-3 1.5l-1 1a3 3 0 0 0-.6 1.3v1c.1.4.3.6.5.8l.9.4l1.1.1c.6 0 1 0 1.6-.2c.6 0 1.1-.2 1.7-.4l.3 1.8a11 11 0 0 1-4.1.9c-.7 0-1.4 0-2-.3c-.5-.1-1-.4-1.3-.8a2 2 0 0 1-.8-1.2C1 14 1 13.6 1 12.9a5 5 0 0 1 .6-1.7c.3-.6.7-1 1.1-1.5a9.6 9.6 0 0 1 4.7-2.3a25 25 0 0 0-2-.4a6.7 6.7 0 0 0-2 0h-.8V5a10.4 10.4 0 0 1 1.9-.2a6.7 6.7 0 0 1 2.5.4l1 .4l1 .3l.9.2h.6l-.2 2.2l-1 .1V9h.1c0 .3.2.5.4.6l.7.2H14a.8.8 0 0 0 .5-.2l.1-.3l-.1-.6l-.6-.9a30.5 30.5 0 0 0-1.4-1.4V6a4.7 4.7 0 0 1 .2-1l.2-.3l.9-.6a23 23 0 0 1 2.9-1.4l1.7-.7l.5 1.8a188.3 188.3 0 0 0-4 1.8a8.8 8.8 0 0 0 1.5 2c.2.4.4.7.4 1v1l-.3.7l-.7.8l-1.3.5l-1.8.2h-2.1Zm-5.4 1.6c.1.2.4.3.7.3c.3 0 .5-.1.8-.3l.4-.7c0-.3 0-.6-.2-.8a.8.8 0 0 0-.6-.3l-.9.3l-.4.8c0 .2 0 .5.2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicArabMeem(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.29 19a54.25 54.25 0 0 1-.73-2.88a30.96 30.96 0 0 1-.42-2.47A15.17 15.17 0 0 1 5 11.79c0-.97.18-1.78.54-2.43a4.22 4.22 0 0 1 1.43-1.57a6.03 6.03 0 0 1 1.99-.84a9.3 9.3 0 0 1 2.2-.22c.74.01 1.45.08 2.1.22a3.93 3.93 0 0 0-.12-1.3a2.8 2.8 0 0 0-.46-1.07a2.1 2.1 0 0 0-.75-.7a1.88 1.88 0 0 0-.99-.27a2.23 2.23 0 0 0-.96.22a3.02 3.02 0 0 0-.89.66c-.28.3-.54.67-.79 1.1L6.17 4.48c.42-.85.9-1.53 1.43-2.03c.54-.5 1.11-.86 1.7-1.09A4.9 4.9 0 0 1 11.1 1c.89 0 1.64.18 2.27.53a4.22 4.22 0 0 1 1.52 1.4c.39.55.67 1.16.84 1.8A6.48 6.48 0 0 1 16 6.58a6.82 6.82 0 0 1-.89 3.23a20.12 20.12 0 0 0-2.41-.48a9.8 9.8 0 0 0-2.04-.1a4.2 4.2 0 0 0-1.54.36a2.16 2.16 0 0 0-1 .86c-.24.39-.36.87-.36 1.46c0 .55.05 1.18.14 1.9c.1.73.23 1.48.4 2.26c.17.78.37 1.55.58 2.32l-2.6.62Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicArabTeh(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.5 16l.37-2.28c1.34 0 2.5-.04 3.5-.12c.99-.07 1.8-.2 2.45-.34a3.54 3.54 0 0 0 1.43-.62c.33-.27.49-.6.49-.99c0-.48-.09-1.04-.25-1.68a17.7 17.7 0 0 0-.62-2.01l2.12-.66a18.76 18.76 0 0 1 .85 2.9c.1.49.16.95.16 1.37a3.8 3.8 0 0 1-.52 2.03a3.79 3.79 0 0 1-1.74 1.37c-.8.35-1.88.61-3.25.78c-1.37.17-2.94.25-4.99.25m0 0c-1.3 0-2.57-.1-3.5-.29a6.9 6.9 0 0 1-2.26-.81c-.6-.36-1.03-.8-1.32-1.32A3.64 3.64 0 0 1 1 11.8a10.77 10.77 0 0 1 .35-2.3l.29-1.06l2.05.5l-.18.74a17.92 17.92 0 0 0-.17.77a4.84 4.84 0 0 0-.06.72c0 .54.17 1 .5 1.39c.35.39.93.68 1.74.87c.82.2 1.92.29 3.33.29l.39 1.66zm3.06-8.58a1.2 1.2 0 0 1-.86-.35a1.2 1.2 0 0 1-.34-.85a1.2 1.2 0 0 1 .34-.85c.24-.25.52-.37.86-.37c.32 0 .6.12.83.37a1.2 1.2 0 0 1 .34.85a1.2 1.2 0 0 1-.34.85a1.13 1.13 0 0 1-.84.35ZM8.47 7.4a1.22 1.22 0 0 1-.87-.35a1.14 1.14 0 0 1-.35-.83a1.2 1.2 0 0 1 .35-.85c.23-.25.52-.37.87-.37c.32 0 .6.12.83.37a1.2 1.2 0 0 1 .35.85c0 .32-.12.6-.35.83c-.23.23-.5.35-.83.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicArmnSha(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.37 1L17 3.183l-.297 2.158l-4.918-.859a14.225 14.225 0 0 0-2.718 2.305a10.685 10.685 0 0 0-1.83 2.723c-.428.964-.642 1.929-.642 2.893c0 1.357.346 2.436 1.038 3.237c.709.801 1.672 1.202 2.892 1.202c1.384 0 2.372-.376 2.966-1.128c.593-.752.89-1.766.89-3.041v-1.079H17v1.226c0 1.897-.56 3.4-1.68 4.512C14.198 18.444 12.6 19 10.524 19c-1.335 0-2.496-.27-3.485-.81a5.734 5.734 0 0 1-2.25-2.304c-.526-.998-.79-2.159-.79-3.483c0-1.013.214-2.043.643-3.09a11.567 11.567 0 0 1 1.754-2.967a11.915 11.915 0 0 1 2.521-2.379l-4.844-.833L4.371 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicC(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 19c-2.01 0-3.26-.6-4.35-1.81S3 14.32 3 12.25a14.05 14.05 0 0 1 1.47-6.23a10.92 10.92 0 0 1 1.87-2.58a8 8 0 0 1 2.65-1.79a10.16 10.16 0 0 1 5.92-.4a7.84 7.84 0 0 1 2.09.8L15.94 4.2a8.81 8.81 0 0 0-1.54-.66a6.3 6.3 0 0 0-1.94-.27a5.6 5.6 0 0 0-3 .78a6.6 6.6 0 0 0-2.09 2.11c-.54.87-.95 1.83-1.23 2.87a12.78 12.78 0 0 0-.39 3.12a5 5 0 0 0 1.03 3.36c.7.8 1.71 1.2 3.02 1.2c.67 0 1.34-.07 2-.22a16.39 16.39 0 0 0 1.93-.57v2.28a13.63 13.63 0 0 1-4.73.79Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicD(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 19L6 1h4.6c2.29 0 4.17.59 5.46 1.76C17.36 3.94 18 5.72 18 8.09c0 2.14-.42 4.03-1.28 5.68a9.36 9.36 0 0 1-3.74 3.85A11.98 11.98 0 0 1 7 19zm5.16-2.3c1.71 0 3.16-.38 4.33-1.15a7.37 7.37 0 0 0 2.66-3.1c.6-1.3.9-2.74.9-4.34c0-1.68-.43-2.9-1.28-3.66a4.86 4.86 0 0 0-3.4-1.16H8.33l-3 13.42h1.83Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicE(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 19L8 1h9l-.6 2H10L8.6 9H15l-.6 2H8.2L7 17h6.5l-.5 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicGeorKan(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 19c-1.08 0-2.1-.22-2.93-.65a4.9 4.9 0 0 1-1.91-1.76a4.77 4.77 0 0 1-.66-2.5c0-.42.04-.82.11-1.2A4.53 4.53 0 0 1 5.4 12h2.3a5.92 5.92 0 0 0-.3 1.84a3.8 3.8 0 0 0 .32 1.56c.21.48.54.87.98 1.17c.46.31 1.1.43 1.79.43c.71 0 1.22-.12 1.67-.43a2.74 2.74 0 0 0 1-1.29c.23-.52.35-1.12.35-1.78c0-.7-.13-1.33-.39-1.89A2.7 2.7 0 0 0 12 10.4a3.05 3.05 0 0 0-1.77-.4H9V8h1.33a3.2 3.2 0 0 0 1.64-.47c.4-.25.7-.57.89-.94a2.8 2.8 0 0 0-.03-2.44a2.3 2.3 0 0 0-.84-.9S11.52 3 11 3V1c.96 0 1.47.19 2.2.56c.73.36 1.3.85 1.7 1.47a3.49 3.49 0 0 1 .64 2.03c0 .96-.3 1.77-.88 2.43c-.58.64-1.2 1.3-2.16 1.51c.64.09 1.09.2 1.63.58c.56.38 1.01.9 1.35 1.58c.35.66.52 1.49.52 2.48a5.4 5.4 0 0 1-.7 2.81a4.63 4.63 0 0 1-1.94 1.87a5.7 5.7 0 0 1-2.86.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicI(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5 19l.33-1.51l2.17-.66l2.9-13.66l-1.9-.63L9 1h7l-.71 1.6l-2.29.57l-2.83 13.66l2.14.66L12 19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicK(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 19L6 1h3L7 9.5L15 1h3l-8.5 9l5.5 9h-3l-5-9l-2 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicS(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 19c-.836 0-1.578-.074-2.225-.22c-.647-.148-1.239-.486-1.775-.78v-2.36c.552.295 1.152.556 1.799.785c.663.213 1.404.319 2.225.319a5.14 5.14 0 0 0 1.775-.294a2.842 2.842 0 0 0 1.325-.932c.332-.425.497-.965.497-1.619s-.205-1.201-.615-1.643c-.41-.458-1.065-.964-1.965-1.52c-.568-.36-1.08-.744-1.538-1.153a5.134 5.134 0 0 1-1.089-1.447c-.268-.556-.402-1.218-.402-1.986c0-1.063.244-1.978.734-2.747A4.942 4.942 0 0 1 8.78 1.638C9.633 1.213 10.588 1 11.645 1c.884 0 1.673.098 2.367.294c.694.18 1.357.442 1.988.785l-.947 2.134a9.318 9.318 0 0 0-1.61-.638a5.61 5.61 0 0 0-1.798-.294c-.868 0-1.578.237-2.13.71c-.552.475-.829 1.129-.829 1.963c0 .719.19 1.275.568 1.667c.395.376.97.81 1.728 1.3c.663.425 1.239.858 1.728 1.3c.49.425.86.915 1.112 1.471c.269.54.403 1.194.403 1.962c0 1.177-.268 2.166-.805 2.967c-.52.785-1.254 1.382-2.201 1.79C10.288 18.804 9.215 19 8 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JournalLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 18.5A1.5 1.5 0 0 0 3.5 20H5V0H3.5A1.5 1.5 0 0 0 2 1.5zM6 0v20h10a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm7 8H8V7h5zm3-2H8V5h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JournalRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 0v20h1.5c.8 0 1.5-.7 1.5-1.5v-17c0-.8-.7-1.5-1.5-1.5zM2 18c0 1.1.9 2 2 2h10V0H4C2.9 0 2 .9 2 2zM4 5h8v1H4zm3 2h5v1H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6a1.54 1.54 0 0 1-1.5-1.5a1.5 1.5 0 0 1 3 0A1.54 1.54 0 0 1 15 6m-1.5-5A5.55 5.55 0 0 0 8 6.5a6.81 6.81 0 0 0 .7 2.8L1 17v2h4v-2h2v-2h2l3.2-3.2a5.85 5.85 0 0 0 1.3.2A5.55 5.55 0 0 0 19 6.5A5.55 5.55 0 0 0 13.5 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 15a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2zm9-9h2v2H9zm0 3h2v2H9zM6 6h2v2H6zm0 3h2v2H6zm-1 5H3v-2h2zm0-3H3V9h2zm0-3H3V6h2zm9 6H6v-2h8zm0-3h-2V9h2zm0-3h-2V6h2zm3 6h-2v-2h2zm0-3h-2V9h2zm0-3h-2V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabFlask(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 7.61V3h1V1H6v2h1v4.61l-5.86 9.88A1 1 0 0 0 2 19h16a1 1 0 0 0 .86-1.51zm-4.2.88a1 1 0 0 0 .2-.6V3h2v4.89a1 1 0 0 0 .14.51l2.14 3.6H6.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 18h-1.44a.61.61 0 0 1-.4-.12a.81.81 0 0 1-.23-.31L17 15h-5l-1 2.54a.77.77 0 0 1-.22.3a.59.59 0 0 1-.4.14H9l4.55-11.47h1.89zm-3.53-4.31L14.89 9.5a11.62 11.62 0 0 1-.39-1.24q-.09.37-.19.69l-.19.56l-1.58 4.19zm-6.3-1.58a13.43 13.43 0 0 1-2.91-1.41a11.46 11.46 0 0 0 2.81-5.37H12V4H7.31a4 4 0 0 0-.2-.56C6.87 2.79 6.6 2 6.6 2l-1.47.5s.4.89.6 1.5H0v1.33h2.15A11.23 11.23 0 0 0 5 10.7a17.19 17.19 0 0 1-5 2.1q.56.82.87 1.38a23.28 23.28 0 0 0 5.22-2.51a15.64 15.64 0 0 0 3.56 1.77zM3.63 5.33h4.91a8.11 8.11 0 0 1-2.45 4.45a9.11 9.11 0 0 1-2.46-4.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LargerText(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.66 18h-2a.85.85 0 0 1-.56-.17a1.11 1.11 0 0 1-.32-.43l-1.33-3.53h-6.9L5.22 17.4a1.06 1.06 0 0 1-.31.41a.83.83 0 0 1-.56.19h-2L8.68 2h2.63zm-4.92-6l-2.2-5.84A16.17 16.17 0 0 1 10 4.43q-.12.52-.27 1t-.27.77L7.26 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12V1H1v18h18v-7z"/><path fill="currentColor" d="M11 1v8h8V1zm6 6h-4V3h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 12v7h18V1h-7v11z"/><path fill="currentColor" d="M1 1v8h8V1zm2 2h4v4H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 19a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-1H8zm9-12a7 7 0 1 0-12 4.9S7 14 7 15v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-1c0-1 2-3.1 2-3.1A7 7 0 0 0 17 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.83 15h2.91a4.88 4.88 0 0 1-1.55-2H5a3 3 0 1 1 0-6h3a3 3 0 0 1 2.82 4h2.1a4.82 4.82 0 0 0 .08-.83v-.34A4.83 4.83 0 0 0 8.17 5H4.83A4.83 4.83 0 0 0 0 9.83v.34A4.83 4.83 0 0 0 4.83 15"/><path fill="currentColor" d="M15.17 5h-2.91a4.88 4.88 0 0 1 1.55 2H15a3 3 0 1 1 0 6h-3a3 3 0 0 1-2.82-4h-2.1a4.82 4.82 0 0 0-.08.83v.34A4.83 4.83 0 0 0 11.83 15h3.34A4.83 4.83 0 0 0 20 10.17v-.34A4.83 4.83 0 0 0 15.17 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkExternalLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 17H3V3h5V1H3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5h-2z"/><path fill="currentColor" d="m11 1l3.29 3.29l-5.73 5.73l1.42 1.42l5.73-5.73L19 9V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkExternalRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 12H1v5c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2h-5v2h5v14H3z"/><path fill="currentColor" d="m1 9l3.3-3.3l5.7 5.7l1.4-1.4l-5.7-5.7L9 1H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkSecure(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.07 8H15V5s0-5-5-5s-5 5-5 5v3H3.93A1.93 1.93 0 0 0 2 9.93v8.15A1.93 1.93 0 0 0 3.93 20h12.14A1.93 1.93 0 0 0 18 18.07V9.93A1.93 1.93 0 0 0 16.07 8M7 5.5C7 4 7 2 10 2s3 2 3 3.5V8H7zM10 16a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListBulletLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 15h12v2H7zm0-6h12v2H7zm0-6h12v2H7z"/><circle cx="3" cy="4" r="2" fill="currentColor"/><circle cx="3" cy="10" r="2" fill="currentColor"/><circle cx="3" cy="16" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListBulletRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15h12v2H1zm0-6h12v2H1zm0-6h12v2H1z"/><circle cx="17" cy="4" r="2" fill="currentColor"/><circle cx="17" cy="10" r="2" fill="currentColor"/><circle cx="17" cy="16" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListNumberedLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 15h12v2H7zm0-6h12v2H7zm0-6h12v2H7zM2 6h1V1H1v1h1zm1 9v1H2v1h1v1H1v1h3v-5H1v1zM1 8v1h2v1H1.5a.5.5 0 0 0-.5.5V13h3v-1H2v-1h1.5a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListNumberedRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 15h11v2H2zm0-6h11v2H2zm0-6h11v2H2zm15-2h-1v1h1v4h1V1zm-2 12v1h2v1h-1v1h1v1h-2v1h3v-5zm0-6v1h2v1h-1.5c-.3 0-.5.2-.5.5V12h3v-1h-2v-1h1.5c.3 0 .5-.2.5-.5v-2c0-.3-.2-.5-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LiteralLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3H5V1h1c.768 0 1.47.289 2 .764A2.989 2.989 0 0 1 10 1h1v2h-1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h1v2h-1c-.768 0-1.47-.289-2-.764A2.989 2.989 0 0 1 6 19H5v-2h1a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1m6 12h6a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-6v2h6v6h-6zm-8-2v2H2a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h2v2H2v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LiteralRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 17h1v2h-1c-.768 0-1.47-.289-2-.764A2.989 2.989 0 0 1 10 19H9v-2h1a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H9V1h1c.768 0 1.47.289 2 .764A2.989 2.989 0 0 1 14 1h1v2h-1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1M8 5H2a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6v-2H2V7h6zm8 2V5h2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2v-2h2V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.07 8H15V5s0-5-5-5s-5 5-5 5v3H3.93A1.93 1.93 0 0 0 2 9.93v8.15A1.93 1.93 0 0 0 3.93 20h12.14A1.93 1.93 0 0 0 18 18.07V9.93A1.93 1.93 0 0 0 16.07 8M10 16a2 2 0 1 1 2-2a2 2 0 0 1-2 2m3-8H7V5.5C7 4 7 2 10 2s3 2 3 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogInLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 11v6c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2H3c-1.1 0-2 .9-2 2v6h8V5l4.75 5L9 15v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogInRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 11v6c0 1.1-.9 2-2 2H3c-1.1 0-2-.9-2-2V3c0-1.1.9-2 2-2h14c1.1 0 2 .9 2 2v6h-8V5l-4.75 5L11 15v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOutLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h8V1H3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8v-2H3z"/><path fill="currentColor" d="M13 5v4H5v2h8v4l6-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOutRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 17H9v2h8c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2H9v2h8z"/><path fill="currentColor" d="M7 15v-4h8V9H7V5l-6 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoCc(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m0-18A9.94 9.94 0 0 0 0 10a9.94 9.94 0 0 0 10 10a9.94 9.94 0 0 0 10-10A9.94 9.94 0 0 0 10 0"/><path fill="currentColor" d="M13.49 11.67c-1 0-1.43-.57-1.43-1.71s.43-1.71 1.43-1.71c.57 0 .86.29 1.14.86l1.29-.71A2.8 2.8 0 0 0 13.2 7a2.91 2.91 0 0 0-2.14.86A2.7 2.7 0 0 0 10.2 10a3 3 0 0 0 .86 2.29a2.91 2.91 0 0 0 2.14.86a3.24 3.24 0 0 0 2.71-1.57L14.63 11a1.46 1.46 0 0 1-1.14.71zm-6 0c-1 0-1.43-.57-1.43-1.71s.43-1.71 1.43-1.71c.57 0 .86.29 1.14.86l1.29-.71A2.8 2.8 0 0 0 7.2 7a2.91 2.91 0 0 0-2.14.86A2.7 2.7 0 0 0 4.2 10a3 3 0 0 0 .86 2.29a2.91 2.91 0 0 0 2.14.86a3.24 3.24 0 0 0 2.71-1.57L8.63 11a1.46 1.46 0 0 1-1.14.71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoMediaWiki(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g transform="translate(10 10)"><g id="oouiLogoMediaWiki0"><path id="oouiLogoMediaWiki1" fill="currentColor" d="M0 10c-2.9-3.3-.8-5.9 0-5.9S2.9 6.7 0 10"/><use href="#oouiLogoMediaWiki1" transform="rotate(15)"/><use href="#oouiLogoMediaWiki1" transform="rotate(30)"/><use href="#oouiLogoMediaWiki1" transform="rotate(45)"/><use href="#oouiLogoMediaWiki1" transform="rotate(60)"/><use href="#oouiLogoMediaWiki1" transform="rotate(75)"/></g><use href="#oouiLogoMediaWiki0" transform="rotate(90)"/><use href="#oouiLogoMediaWiki0" transform="rotate(180)"/><use href="#oouiLogoMediaWiki0" transform="rotate(270)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoMetaWiki(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.346 1.15A9.966 9.966 0 0 1 9.991 0a9.963 9.963 0 0 1 4.652 1.146a9.997 9.997 0 0 1 3.585 3.187l-1.932 1.108a7.778 7.778 0 0 0-2.762-2.375a7.751 7.751 0 0 0-9.835 2.375L1.767 4.333A10 10 0 0 1 5.346 1.15m-.197 4.294A6.701 6.701 0 0 1 7.034 4.04c-.419.41-.786.869-1.093 1.368c-.263.03-.528.042-.792.036m2.56-1.683l-.048.002v.017l.048-.02Zm0 0a7.156 7.156 0 0 1 1.622.118a6.47 6.47 0 0 0 .527-.554a6.608 6.608 0 0 0-2.15.436Zm-1.18 1.544c.234-.332.497-.643.787-.927l.26.603c-.34.14-.69.248-1.047.324m1.197-1.1L7.724 4.2v.006h.002Zm0 0l.255.594c.312-.16.61-.345.892-.554a6.622 6.622 0 0 0-1.147-.04m2.127-.207a6.88 6.88 0 0 0 .57-.648v.006a6.721 6.721 0 0 1 2.999.922a9.49 9.49 0 0 1-.809 1.144a7.197 7.197 0 0 0-2.76-1.424M3.462 8.711a6.746 6.746 0 0 1 1.317-2.844v.006c.297.02.596.02.894 0a7.241 7.241 0 0 0-.809 3.004a9.89 9.89 0 0 1-1.402-.166m6.002-4.35l.008-.006l-.011.006zm0 0a6.524 6.524 0 0 1-1.295.845l1.19 2.768a9.451 9.451 0 0 0 2.964-2.215a6.742 6.742 0 0 0-2.859-1.397Zm-3.252 1.46l.002-.004l-.01.006l.008-.001Zm0 0a6.796 6.796 0 0 0-.914 3.07a9.49 9.49 0 0 0 3.635-.745l-1.19-2.752a6.557 6.557 0 0 1-1.531.428Zm7.577-1.29l.005-.01l-.01.006zm0 0a9.696 9.696 0 0 1-.845 1.187a7.089 7.089 0 0 1 1.839 2.877c.448-.46.862-.954 1.237-1.476a6.617 6.617 0 0 0-2.231-2.588M3.4 9.156l.001-.007l-.01.005l.01.002Zm0 0a6.664 6.664 0 0 0 .44 3.398c.637.07 1.278.094 1.919.072a6.779 6.779 0 0 1-.304-.62a7.023 7.023 0 0 1-.585-2.683a10.03 10.03 0 0 1-1.47-.167m9.247-3.095l.007-.008l-.01.005zm0 0a9.94 9.94 0 0 1-3.11 2.307l1.31 3.032a13.248 13.248 0 0 0 3.588-2.44a7.07 7.07 0 0 0-.276-.742a6.614 6.614 0 0 0-1.512-2.157M5.85 11.826a6.663 6.663 0 0 1-.552-2.492a9.912 9.912 0 0 0 3.82-.8l1.312 3.046c-1.331.567-2.746.91-4.188 1.016a6.488 6.488 0 0 1-.392-.77M14.913 9.1c.466-.466.899-.965 1.294-1.493h.003a6.62 6.62 0 0 1 .337 3.691c-.515.487-1.06.94-1.635 1.354a7.101 7.101 0 0 0 0-3.552Zm-8.365 6.65a6.637 6.637 0 0 1-2.51-2.721c.655.06 1.313.073 1.97.038a7.092 7.092 0 0 0 2.602 2.385c-.68.144-1.369.244-2.062.299Zm4.467-3.94a13.633 13.633 0 0 0 3.53-2.353a6.712 6.712 0 0 1-.215 3.602c-.714.473-1.463.888-2.242 1.244zm-1.737 3.484a6.65 6.65 0 0 1-2.76-2.265a13.494 13.494 0 0 0 4.077-1.041l1.074 2.492c-.775.332-1.574.604-2.39.814Zm5.405-1.938a16.44 16.44 0 0 0 1.657-1.269a6.663 6.663 0 0 1-3.393 3.943a7.184 7.184 0 0 0 1.736-2.675Zm-4.906 3.368a6.627 6.627 0 0 1-2.528-.591l.006-.009a16.868 16.868 0 0 0 2.001-.357a7.177 7.177 0 0 0 3.081.532a6.626 6.626 0 0 1-2.56.425m2.483-2.006c.61-.277 1.203-.592 1.775-.942l.005-.008a6.777 6.777 0 0 1-1.363 1.916zm.005 1.138a6.678 6.678 0 0 1-2.274-.296c.63-.189 1.25-.412 1.855-.67zM.715 6.302a10.062 10.062 0 0 0 .704 8.879a10.02 10.02 0 0 0 3.172 3.286A9.984 9.984 0 0 0 8.886 20v-2.248a7.771 7.771 0 0 1-5.542-3.686a7.825 7.825 0 0 1-.683-6.634zm16.619 1.13l1.951-1.13a10.063 10.063 0 0 1-.704 8.88a10.021 10.021 0 0 1-3.173 3.286A9.985 9.985 0 0 1 11.112 20v-2.248a7.763 7.763 0 0 0 5.544-3.684a7.817 7.817 0 0 0 .678-6.636"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikibooks(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.222.03c-1.812 1.753-3.604 3.49-3.604 3.49l.644 1.31c.873-.864 1.75-1.625 2.587-2.472c1.762 3.869 2.963 6.987 4.44 11.109c.483 1.339.384 1.399-.771 2.122c-1.857.923-2.445 1.53-4.658 3.884h1.515c1.85-1.673 2.536-2.263 4.853-3.484c.82-.473.932-1.07.596-2C18.74 10.977 16.39 4.484 14.222.03"/><path fill="currentColor" d="M9.498.071C8.172.966 7.149 2.254 5.996 3.358l.473 1.508c.903-.71 1.688-1.569 2.5-2.38l4.376 10.977c.252.743-.145 1.405-.944 2.125c-.956.769-2.925 2.56-3.818 3.926h1.358c1.426-1.464 1.683-1.663 3.306-3.068c1.09-1.112 2.074-1.681 1.56-3.106z"/><path fill="currentColor" d="M5.023 0C4.67.12 4.376.347 4.1.588a879.5 879.5 0 0 0-3.889 3.75c-.196.182-.213.45-.21.718L.017 6.57S1.743 4.892 4.32 2.41c1.39 3.615 2.837 7.21 4.263 10.813c.195.544-.272.99-.864 1.585c-.818.798-2.53 2.44-4.281 4.635h1.756l4.22-4.629c.402-.552.9-1.091.69-1.591L5.024 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikidata(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 4v12.258h.742V4zm1.482 0v12.258h2.223V4zm2.96 0v12.258H6.67V4zm2.964 0v12.258h.744V4zm1.48 0v12.258h.745V4zm1.483 0v12.258h2.224V4zm2.962 0v12.258h.742V4zm1.482 0v12.258h2.223V4zm2.96 0v12.258h.744V4zm1.484 0v12.258H20V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikifunctions(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.752 7.06c-.165 0-.352.01-.56.03c-.175.02-.352.03-.528.032c-.145 0-.325-.01-.54-.032a7.631 7.631 0 0 0-.6-.03c-.063 0-.094.047-.094.144c0 .098.02.145.062.145c.464.021.697.2.697.538c-.004.16-.04.318-.104.465l-1.415 3.284l.23.515l1.672-3.758c.175-.393.34-.66.503-.798c.178-.146.397-.233.627-.246a.11.11 0 0 0 .087-.052a.188.188 0 0 0 .04-.114c-.003-.083-.028-.124-.077-.124Zm-5.038.289c.465.022.697.19.697.506a1.117 1.117 0 0 1-.156.497l-.028.053l.23.516l.286-.528c.245-.427.448-.701.61-.822c.193-.13.418-.207.65-.222c.076 0 .115-.056.115-.166c0-.087-.021-.124-.062-.124c-.174 0-.396.012-.659.036a6.246 6.246 0 0 1-.47.027c-.15 0-.34-.01-.571-.032a8.075 8.075 0 0 0-.61-.03c-.02 0-.041.015-.062.046a.173.173 0 0 0-.031.098c-.002.095.019.145.061.145M8.07 11.625L6.56 8.228a.803.803 0 0 1-.087-.31a.47.47 0 0 1 .2-.394c.163-.108.353-.169.548-.174c.034 0 .051-.047.051-.144c0-.096-.028-.145-.087-.145c-.23 0-.497.01-.798.03a7.648 7.648 0 0 1-1.41 0a9.66 9.66 0 0 0-.727-.03c-.054 0-.087.048-.087.145c0 .037.013.072.037.1a.103.103 0 0 0 .078.044c.272.017.532.111.751.272c.18.146.351.404.513.773l1.91 4.368l.616-1.138Z"/><path fill="currentColor" d="M11.979 14.312a.138.138 0 0 0 .125.087h1.137a.134.134 0 0 0 .122-.19l-2.777-6.215a1.67 1.67 0 0 0-.358-.506a1.413 1.413 0 0 0-.552-.314a2.968 2.968 0 0 0-.878-.11c-.347 0-.692.05-1.025.148a.135.135 0 0 0-.092.15l.16.963a.132.132 0 0 0 .13.112a.15.15 0 0 0 .043-.008c.18-.06.37-.089.56-.087c.14-.004.28.014.414.055a.68.68 0 0 1 .269.162a.877.877 0 0 1 .174.26c.048.102.098.22.151.348l.113.284l-2.536 4.765a.135.135 0 0 0 .126.182H8.63a.133.133 0 0 0 .122-.08l.008-.015l1.707-3.148l1.513 3.157ZM3.752 5.208a7.855 7.855 0 0 1 6.225-3.085A7.875 7.875 0 0 1 16.22 5.17l1.77-1.186a10.009 10.009 0 0 0-16.027.064z"/><path fill="currentColor" d="M17.288 7.01a7.8 7.8 0 0 1 .585 2.965c0 4.331-3.531 7.903-7.873 7.903c-4.341 0-7.871-3.572-7.871-7.903c-.001-1 .191-1.993.567-2.92L.883 5.88a9.886 9.886 0 0 0-.882 4.095C0 15.477 4.487 20 10 20c5.513 0 10-4.523 10-10.025a9.9 9.9 0 0 0-.917-4.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikimedia(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.08 5.555a6.048 6.048 0 0 0 3.055 10.593v-7.54L6.08 5.556zm7.828.004l-3.05 3.05v7.536a6.048 6.048 0 0 0 3.05-10.587z"/><path fill="currentColor" d="M3.414 2.89C1.424 4.69.164 7.287.168 10.173c.007 5.406 4.42 9.806 9.828 9.806c5.407 0 9.82-4.4 9.828-9.806c.004-2.886-1.255-5.482-3.246-7.285L14.865 4.6a7.355 7.355 0 0 1 2.524 5.568c-.007 4.09-3.3 7.375-7.394 7.375S2.61 14.26 2.604 10.17a7.355 7.355 0 0 1 2.523-5.568z"/><circle cx="10" cy="3.32" r="3.32" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikimediaCommons(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.09 6.18a3.68 3.68 0 0 1-2.18-2.55c.09.09 1.82.91 1.82.91L10 0L7.27 4.55l1.82-.91a5.08 5.08 0 0 0 .55 1.91a5.13 5.13 0 0 0 2 2a8.86 8.86 0 0 1 2 1.18l-.64.63l-.45-.45l-.26 1.54l1.54-.26l-.45-.45l.62-.65a5.69 5.69 0 0 1 1.45 3.45h-.91v-.73l-1.26.91l1.26.91v-.73h.91A5.21 5.21 0 0 1 14 16.36l-.64-.64l.45-.45l-1.53-.27l.26 1.54l.45-.45l.64.64a5.69 5.69 0 0 1-3.45 1.45v-.91h.73L10 16l-.91 1.27h.73v.91a5.21 5.21 0 0 1-3.45-1.45l.63-.64l.45.45l.27-1.54l-1.54.26l.45.45l-.63.65a5.69 5.69 0 0 1-1.45-3.45h.91v.73l1.26-.91l-1.26-.91v.73h-.91A5.21 5.21 0 0 1 6 9.09l.64.64l-.45.45l1.54.26l-.28-1.53l-.45.45l-.64-.64L5 7.45a7.29 7.29 0 1 0 8.09-1.27"/><circle cx="10" cy="12.7" r="2.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikimediaDiscovery(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17c0 1.1-2 2-2 2s-2-.9-2-2m2-10a1.54 1.54 0 0 1-1.5-1.5a1.5 1.5 0 0 1 3 0A1.54 1.54 0 0 1 10 7m3.3 4.7C14.1 7.9 12.7 1 10 1S5.8 7.7 6.6 11.5L5 15h2.7l.3 1h4c.2-.3.1-.5.3-1H15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikinews(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10a6 6 0 0 1-6 6a6 6 0 0 1-6-6a6 6 0 0 1 6-6a6 6 0 0 1 6 6M7.55 8.94a9.97 9.97 0 0 1 1.33-3.81a5 5 0 0 0-3.86 4.45H7.5c0-.22.03-.43.05-.64m1 .12a7.96 7.96 0 0 0-.05.52h3.06a7.99 7.99 0 0 0-.05-.52a9.09 9.09 0 0 0-1.48-3.89a9.07 9.07 0 0 0-1.48 3.9Zm4.01.52h2.42a5 5 0 0 0-3.79-4.44a9.99 9.99 0 0 1 1.37 4.44m-7.53 1a5 5 0 0 0 3.87 4.3a9.37 9.37 0 0 1-1.4-4.3zm5 4.21c1-1.47 1.43-2.9 1.52-4.21H8.52a8.53 8.53 0 0 0 1.52 4.21zm1.14.07a5 5 0 0 0 3.8-4.28h-2.41a9.35 9.35 0 0 1-1.39 4.28M3.07 6a7.97 7.97 0 0 0 0 8h1.18A6.96 6.96 0 0 1 3 10c0-1.49.46-2.87 1.25-4Zm12.68 0h1.18a7.97 7.97 0 0 1 0 8h-1.18A6.96 6.96 0 0 0 17 10c0-1.49-.46-2.87-1.25-4m2.74 7h1.05a9.9 9.9 0 0 0 0-6h-1.05a9.04 9.04 0 0 1 0 6M.46 13h1.05a9.04 9.04 0 0 1 0-6H.46a9.9 9.9 0 0 0 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikipedia(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.14 4H14a.69.69 0 0 1 0 .65c-1 .16-1.36.91-1.81 1.83l-1.4 2.75l2.35 5.21h.07l3.52-8.1c.44-1.07.4-1.59-.79-1.7a.68.68 0 0 1 0-.65h3.45a.68.68 0 0 1 0 .65c-1.21.16-1.42.91-1.81 1.83l-4.37 10.08c-.13.3-.24.45-.44.45s-.33-.16-.42-.45l-2.48-5.73l-2.72 5.73c-.11.3-.24.45-.44.45s-.31-.16-.42-.45l-4-10.09c-.57-1.4-.6-1.7-1.65-1.8A.68.68 0 0 1 .62 4h3.91a.68.68 0 0 1 0 .65c-1.16.13-1.21.45-.74 1.58l3.41 8.19h.05L9.3 10L7.78 6.45C7.17 5.05 7 4.77 6.24 4.66a.69.69 0 0 1 0-.65h3.32a.68.68 0 0 1 0 .65c-.74.12-.7.45-.19 1.58l.87 2l.08.09l1-2c.57-1.14.64-1.58-.15-1.7a.69.69 0 0 1-.03-.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikiquote(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.133 4.643a6.521 6.521 0 0 1 2.79 5.354a6.52 6.52 0 0 1-2.672 5.269l-.12.086l-.923-1.322a4.908 4.908 0 0 0 2.103-4.033c0-1.578-.748-3.03-1.988-3.95l-.114-.082z"/><path fill="currentColor" d="M5.347 13.275a3.347 3.347 0 1 0-.001-6.695a3.347 3.347 0 0 0 .001 6.695M13.874 0a13.082 13.082 0 0 1 4.627 9.997c0 3.86-1.678 7.448-4.538 9.92l-.096.083l-1.15-1.357a11.303 11.303 0 0 0 4.006-8.646c0-3.33-1.444-6.424-3.908-8.562l-.091-.079L13.873 0Z"/><path fill="currentColor" d="M11.453 2.386a9.71 9.71 0 0 1 3.668 7.611a9.71 9.71 0 0 1-3.576 7.537l-.09.073l-1.074-1.347A7.988 7.988 0 0 0 13.4 9.997a7.987 7.987 0 0 0-2.932-6.194l-.087-.07l1.073-1.347Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikisource(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.94 0a10 10 0 0 0-8.08 4.2h4.21c.06-.35.16-.72.49-1.04c.2-.23.67-.2 1-.26c.33-.07.87-.92 1.4-1.17c.64.33 1.34.83 1.7 1.08c.62-.49 1.16-1.33 1.4-.65c.2-.33.53-.54.9-.63c.17-.04.4-.03.52-.05c.14-.06.42-.27.66-.34c-.38.58-.43 2.25-.44 3.06v.35c.55.12 1.25.03 1.49-.35h2.95A9.99 9.99 0 0 0 10 0zm-.36 2.5c.17.47.31.94.26 1.4c-.06.3.07.28.31.16c.24-.12.36-.26.72-.37c.27-.07.08-.26-.05-.4c-.4-.34-.82-.56-1.24-.8zm-8.3 2.6A9.95 9.95 0 0 0 0 10a10 10 0 0 0 20 0a9.94 9.94 0 0 0-1.28-4.9H2.8c.52 1.1.81 1.7 2.6 1.46c.62-.1 1-.16 1 .48c.08.39.38.52 1.15.38c.3-.11.57-.28.79-.14h.02a.1.1 0 0 1 0 .01a.04.04 0 0 1 .01.02c0 .1-.19.31.04.57c.57.55 1.12 1.04 1.7 1.6c.13.64.7.55 1.15.42c.68-.2.38.1.63.41c.14.22.46.81.84.6c.01 1.08.28 1.85.84 2.27c-.17.63-.33 1.5-.03 2.02c-.5-.16-.3.81-.45 1.22c-.26.68-.48.65-.72.71c.22.24.5.49.92.64v.01l.02.01c0 .23.04.43.18.57c-.26-.06-.57-.42-1.08-.19c-.56.25-.68.08-.48-.29c.26-.57-.3-.78-.88-.92v-.23c.12-.04.26-.08.42-.1c.25-.05.42-.14.4-.5c-.01-.5-.03-1 0-1.48a6.08 6.08 0 0 1 .18-1.21c.07-.2.44-.38.13-.64A42.07 42.07 0 0 0 8.3 9.3a.47.47 0 0 0-.06-.02l1.65 2.93c.2.49.58 1-.1 1.13v.01c-.03.47.06 1.28-.26 1.44c-.36.25-.39.48-.42.86c-.06.4-.06.86.06 1.18c.12.05.24.08.37.1l.32.04c-.43.56-.96.39-1.27-.13l-1.67-2.58c.08-2.51.13-4.2-.05-6.3c-.92.08-1.85.01-2.83-.56c-.22-.4-.47-.36-.89-.26c.05-.37-.1-.73-.32-1.1l-.1-.93H1.29zm14.51 1.03c1.6-.04 1.14 2.78 1.14 4.15h.01l-.13.7l-.2 1.21c.07.6-.09.68-.2 1.02c0 1.32-.18 1.7-1.29 2.7a2.16 2.16 0 0 1-1.17-2.25c.04-.24.2-.51-.01-.9a7.39 7.39 0 0 1 .14-3.35a.64.64 0 0 1 .2-.02a.63.63 0 0 0 .57-.74a3.34 3.34 0 0 0 .6-.2c-.81-.76-.52-2.3.34-2.32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikispecies(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="9.5" cy="2.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="M14.434 7.33a6.407 6.407 0 0 1 1.677 4.322c0 3.582-2.951 6.488-6.59 6.488c-3.64 0-6.595-2.906-6.595-6.488c0-1.643.624-3.141 1.648-4.284L3.138 5.967a8.374 8.374 0 0 0-2.137 5.592C1 16.219 4.81 20 9.503 20s8.498-3.78 8.498-8.441a8.383 8.383 0 0 0-2.152-5.615Z"/><path fill="currentColor" d="M4.368 3.754L4.34 5.525s.22.93 1.75 1.411c1.487.468 3.581.3 5.048.964c1.273.575 1.71 1.53 1.71 1.53l.03-1.66s-.11-.876-1.285-1.239c-1.176-.362-3.141-.48-5.428-1.265c-1.025-.352-1.812-1.485-1.798-1.512Zm1.665 3.798L5.97 9.5s.16.415 1.578.826c.848.246 3.364.467 4.666.998c1.139.464 1.665 1.554 1.665 1.554l.12-1.623c-.376-.92-.81-1.148-2.342-1.544c-1.044-.27-2.24-.318-3.163-.625c-1.811-.603-2.46-1.533-2.46-1.533Z"/><path fill="currentColor" d="M14.492 3.74v1.738s-.244.807-1.556 1.295c-1.69.628-4.273.583-6.004 1.556a2.033 2.033 0 0 0-.962 1.167l.027-1.999s.3-.743 1.83-1.22c1.964-.614 3.21-.38 4.962-1.128c1.14-.487 1.716-1.395 1.703-1.409m-9.6 7.474l-.069 1.918s.301.61 1.192.99c1.703.728 2.904.193 5.405 1.332c1.129.513 1.395 1.362 1.395 1.362l-.006-1.926s-.147-.59-1.242-.953c-1.095-.362-3.488-.609-5.113-1.22c-1.117-.42-1.576-1.517-1.562-1.503"/><path fill="currentColor" d="m5.984 16.62l.027-1.837s.126-.53 1.83-1.074c1.606-.512 3.602-.502 4.537-.89c.935-.39 1.62-1.566 1.62-1.566l-.007 1.95s.049.696-1.378 1.15c-1.856.59-3.03.38-4.859 1.065c-1.036.389-1.743 1.243-1.77 1.203Zm6.896-8.851l.022 1.674s-.262.492-1.417.96c-1.573.64-3.003.278-5.007 1.133c-.748.318-1.656 1.596-1.656 1.596l.053-2.026s.539-.87 1.51-1.22c2.083-.752 2.858-.202 5.035-1.1c.736-.305 1.438-.934 1.46-1.017"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikiversity(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.982 3c-.2.04-.44.196-.641.47c-.2.234-.361.585-.521.976h2.284c-.16-.43-.32-.742-.52-.977c-.161-.273-.402-.43-.602-.469m.722.079c.08.078.12.117.16.195c.24.312.44.703.601 1.172h1.604c-.642-.703-1.443-1.211-2.365-1.367m-1.444 0a4.365 4.365 0 0 0-2.365 1.406H8.5c.16-.469.36-.899.601-1.172a.6.6 0 0 1 .16-.234Zm-1.162.039a6.346 6.346 0 0 0-2.285 1.328h.682a5.163 5.163 0 0 1 1.603-1.328m3.848.039c.601.312 1.123.742 1.604 1.328h.681a6.124 6.124 0 0 0-2.285-1.328m-6.494 1.64A6.08 6.08 0 0 0 4.29 6.596h1.042a7.5 7.5 0 0 1 .882-1.797h-.762Zm1.163 0a7.08 7.08 0 0 0-.962 1.798h2.365c.08-.586.2-1.211.4-1.797zm2.124 0c-.2.587-.32 1.173-.4 1.798h3.327a9.649 9.649 0 0 0-.401-1.797zm2.887 0c.16.587.32 1.173.4 1.798h2.325a5.61 5.61 0 0 0-.962-1.797zm2.164 0c.401.548.682 1.134.882 1.798h1.042c-.28-.664-.681-1.25-1.162-1.797zm-9.62 2.11a6.091 6.091 0 0 0-.361 2.072H4.97c0-.704.08-1.407.28-2.071zm1.403 0a7.18 7.18 0 0 0-.281 2.072h2.566c0-.704.04-1.368.12-2.071zm2.725 0c-.08.704-.12 1.368-.12 2.072h3.648c0-.704-.04-1.368-.12-2.071zm3.728 0c.08.704.12 1.368.12 2.072h2.566a7.18 7.18 0 0 0-.28-2.071zm2.726 0c.2.665.28 1.368.28 2.072h1.163a6.1 6.1 0 0 0-.36-2.071zM20 9.332H0v.646h20zm-16.645 1H.685v6.707h2.67zm5.321 0H6.005v6.707h2.67zm5.32 0h-2.67v6.707h2.67zm5.321 0h-2.671v6.707h2.67zM20 17.355H0V18h20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWikivoyage(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.874 8.19L20 0l-8.254 4.095c.99.263 1.892.78 2.617 1.498a5.78 5.78 0 0 1 1.51 2.596Zm-6.349 4.043L14.678 20l.56-9.291a5.916 5.916 0 0 1-2.65 1.523a5.95 5.95 0 0 1-3.063 0Zm-.159-7.508L.001 5.28l7.829 5.113a5.813 5.813 0 0 1 0-3.038c.27-.995.8-1.902 1.536-2.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoWiktionary(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.22 18.6v.01c-.35-.21-.61-.5-.71-.84l-.07-.37L.21 3.36c-.03-.45.17-.9.57-1.25c.39-.36.97-.6 1.62-.66L15.35.22a2.81 2.81 0 0 1 1.7.35a1.5 1.5 0 0 1 .77 1.13l1.23 14.12c.03.45-.17.9-.57 1.25a2.85 2.85 0 0 1-1.62.67L3.92 18.95a2.75 2.75 0 0 1-1.7-.35m-1-1.1c.02.18.07.35.15.5l.02.25c.05.56.4 1.03.9 1.34c.51.3 1.19.46 1.9.4l13.34-1.27a3.15 3.15 0 0 0 1.8-.74c.45-.4.71-.93.66-1.49L18.73 1.87a1.77 1.77 0 0 0-.9-1.33a2.92 2.92 0 0 0-1.24-.4a3.23 3.23 0 0 0-1.27-.12L2.4 1.23a3.1 3.1 0 0 0-1.74.72c-.44.39-.7.9-.64 1.44l1.22 14.1zm1.2 1.9a1.6 1.6 0 0 1-.78-1a2 2 0 0 0 .47.39c.49.3 1.14.44 1.84.38l12.93-1.22c.7-.06 1.31-.33 1.74-.72c.44-.38.7-.9.64-1.43L18.04 1.69a1.62 1.62 0 0 0-.62-1.11c.1.04.2.09.29.15c.46.28.76.7.8 1.16l1.26 14.62c.04.48-.18.94-.59 1.3c-.4.37-1 .63-1.67.7L4.17 19.75a2.9 2.9 0 0 1-1.76-.36ZM1.21 5.3l4.34-.5l.06.47l-.28.04c-.28.03-.48.12-.6.26a.57.57 0 0 0-.15.46a11.8 11.8 0 0 0 .53 1.33l2.91 6.12l1.15-5.56l-.8-1.68c-.16-.27-.31-.5-.48-.7a.97.97 0 0 0-.28-.23a1.42 1.42 0 0 0-.42-.17c-.1-.02-.25-.02-.48 0l-.08.02l-.06-.48l4.56-.53l.06.48l-.38.04c-.3.04-.5.13-.6.26a.67.67 0 0 0-.13.53c0 .02 0 .06.03.15l.15.4l3.2 6.84l1.32-6.48c.16-.75.22-1.25.18-1.51a.57.57 0 0 0-.14-.32a.57.57 0 0 0-.3-.18c-.2-.05-.47-.06-.8-.02h-.08l-.06-.48l3.53-.4l.06.48h-.08c-.29.04-.5.12-.66.24c-.15.12-.3.33-.42.64c-.08.2-.2.7-.35 1.5l-2.01 9.9l-.45.05l-3.4-7.06l-1.61 7.64l-.42.04l-4.56-9.42c-.34-.7-.55-1.11-.63-1.23a.99.99 0 0 0-.47-.4a1.6 1.6 0 0 0-.76-.07h-.08Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3L7 1L1 3v16l6-2l6 2l6-2V1zM7 14.89l-4 1.36V4.35L7 3zm10 .75L13 17V5.1l4-1.36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPin(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0a7.65 7.65 0 0 0-8 8c0 2.52 2 5 3 6s5 6 5 6s4-5 5-6s3-3.48 3-6a7.65 7.65 0 0 0-8-8m0 11.25A3.25 3.25 0 1 1 13.25 8A3.25 3.25 0 0 1 10 11.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinAdd(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0a7.65 7.65 0 0 0-8 8c0 2.52 2 5 3 6s5 6 5 6s4-5 5-6s3-3.48 3-6a7.65 7.65 0 0 0-8-8m5 9h-4v4H9V9H5V7h4V3h2v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 1v16l6 2l6-2l6 2V3l-6-2l-6 2zm12 2l4 1.36v11.9l-4-1.36zM3 3.74L7 5.1V17l-4-1.36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapTrail(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20 6l-1-1l-1.5 1.5L16 5l-1 1l1.5 1.5L15 9l1 1l1.5-1.5L19 10l1-1l-1.5-1.5z"/><circle cx="7.5" cy="14.5" r="3.5" fill="currentColor"/><circle cx="7" cy="3" r="2" fill="currentColor"/><circle cx="13" cy="7" r="1" fill="currentColor"/><circle cx="10" cy="6" r="1" fill="currentColor"/><circle cx="3" cy="3" r="1" fill="currentColor"/><circle cx="1" cy="6" r="1" fill="currentColor"/><circle cx="1" cy="9" r="1" fill="currentColor"/><circle cx="3" cy="12" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Markup(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 3.5L0 10l1.5 1.5l5 5L8 15l-5-5l5-5zm7 0L12 5l5 5l-5 5l1.5 1.5L20 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mathematics(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H4l5 8l-5 8h12v-4h-2v2H8.25L12 10L8.25 4H14v2h2V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathematicsDisplayBlock(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 5H5l3 5l-3 5h10v-3h-2v1H9.2l1.8-3l-1.8-3H13v1h2V5zM2 1h16v2H2zm0 16h16v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathematicsDisplayDefault(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5H4l3 5l-3 5h10v-3h-2v1H8.2l1.8-3l-1.8-3H12v1h2V5zM1 9h3v2H1zm15 0h3v2h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathematicsDisplayInline(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 13H0V7h4zm12-6h4v6h-4zM6 6l3 4l-3 4h8v-3h-2v1H9.5l1.5-2l-1.5-2H12v1h2V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3v2h18V3zm0 8h18V9H1zm0 6h18v-2H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 8v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8l-10 4z"/><path fill="currentColor" d="M2 2a2 2 0 0 0-2 2v2l10 4l10-4V4a2 2 0 0 0-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.39 15.14A7.33 7.33 0 0 1 11.75 1.6c.23-.11.56-.23.79-.34a8.19 8.19 0 0 0-5.41.45a9 9 0 1 0 7 16.58a8.42 8.42 0 0 0 4.29-3.84a5.3 5.3 0 0 1-1.03.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19 10l-4-3v2h-4V5h2l-3-4l-3 4h2v4H5V7l-4 3l4 3v-2h4v4H7l3 4l3-4h-2v-4h4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveFirstLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1h2v18H3zm13.5 1.5L15 1l-9 9l9 9l1.5-1.5L9 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveFirstRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1h2v18h-2zM3.5 2.5L11 10l-7.5 7.5L5 19l9-9l-9-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveLastLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1h2v18h-2zM3.5 2.5L11 10l-7.5 7.5L5 19l9-9l-9-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveLastRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1h2v18H3zm13.5 1.5L15 1l-9 9l9 9l1.5-1.5L9 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicalScore(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15V5h8v10h2V2H6v13"/><circle cx="15" cy="15" r="3" fill="currentColor"/><circle cx="5" cy="15" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Network(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="10" cy="15" r="2" fill="currentColor"/><path fill="currentColor" d="M1 7.4a12 13 0 0 1 18 0l-1.5 1.4a10 11.1 0 0 0-15 0zm3.7 3.2a7 7.3 0 0 1 10.7 0L14 12a5 5.3 0 0 0-7.8 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkOff(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="10" cy="16" r="2" fill="currentColor"/><path fill="currentColor" d="M16.4 11.6A7.1 7.1 0 0 0 12 9.1l3.4 3.4zM19 8.4A12.2 14 0 0 0 8.2 4.2L10 6a9.9 9.9 0 0 1 7.4 3.7zM3.5 2L2 3.4l2.2 2.2A13.1 13.1 0 0 0 1 8.4l1.5 1.3a10.7 10.7 0 0 1 3.2-2.6L8 9.3a7.3 7.3 0 0 0-3.3 2.3L6.1 13a5.2 5.2 0 0 1 3.6-2l6.8 7l1.5-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewWindowLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 17H3V3h5V1H3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5h-2z"/><path fill="currentColor" d="m11 1l3.3 3.3L8.6 10l1.4 1.4l5.7-5.7L19 9V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewWindowRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 12H1v5c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2h-5v2h5v14H3z"/><path fill="currentColor" d="m1 9l3.3-3.3l5.7 5.7l1.4-1.4l-5.7-5.7L9 1H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewlineLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4v6H7V6l-6 5l6 5v-4h12V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewlineRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 4v8h12v4l6-5l-6-5v4H3V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewspaperLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2a2 2 0 0 0-2 2v12a1 1 0 0 1-1-1V5h-.5A1.5 1.5 0 0 0 0 6.5v10A1.5 1.5 0 0 0 1.5 18H18a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm1 2h11v4H6zm0 6h6v1H6zm0 2h6v1H6zm0 2h6v1H6zm7-4h4v5h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewspaperRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 16c0 1.1.9 2 2 2h16.5c.8 0 1.5-.7 1.5-1.5v-10c0-.8-.7-1.5-1.5-1.5H18v10c0 .6-.4 1-1 1V4c0-1.1-.9-2-2-2H2C.9 2 0 2.9 0 4zM3 4h11v4H3zm0 6h4v5H3zm5 0h6v1H8zm0 2h6v1H8zm0 2h6v1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1L5.6 2.5L13 10l-7.4 7.5L7 19l9-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4 10l9 9l1.4-1.5L7 10l7.4-7.5L13 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoWikiText(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3v2h1v10l2 2V3zM9 5V3H5l2 2zM1 1L0 2l1 1v14h3v-2H3V5l2 2v10h4v-2H7V9l6 6h-2v2h4l3 3l1-1zm12 10l2 2V3h-4v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotBright(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="9.85" cy="10" r="9" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notice(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0a10 10 0 1 0 10 10A10 10 0 0 0 10 0m1 16H9v-2h2zm0-4H9V4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ocr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1a2 2 0 0 0-2 2v4.5h2V3h5V1zM1 17v-5.5h2V17h5v2H3a2 2 0 0 1-2-2m11 0v2h5a2 2 0 0 0 2-2v-5.5h-2V17zm5-9.5h2V3a2 2 0 0 0-2-2h-5v2h5z"/><path fill="currentColor" d="M6 5h8v1H6zM5 8h10v1H5zm1 3h8v1H6zm-1 3h10v1H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OngoingConversationLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0a2 2 0 0 0-2 2v18l4-4h14a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm3 9.06a1.39 1.39 0 1 1 1.37-1.39A1.39 1.39 0 0 1 5 9.06m5.16 0a1.39 1.39 0 1 1 1.39-1.39a1.39 1.39 0 0 1-1.42 1.39zm5.16 0a1.39 1.39 0 1 1 1.39-1.39a1.39 1.39 0 0 1-1.42 1.39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OngoingConversationRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 14c0 1.1.9 2 2 2h14l4 4V2c0-1.1-.9-2-2-2H2C.9 0 0 .9 0 2zm13.6-6.3c0-.8.6-1.4 1.4-1.4c.8 0 1.4.6 1.4 1.4s-.6 1.4-1.4 1.4c-.8-.1-1.4-.7-1.4-1.4M9.9 9.1s-.1 0 0 0c-.8 0-1.4-.6-1.4-1.4c0-.8.6-1.4 1.4-1.4c.8 0 1.4.6 1.4 1.4s-.7 1.4-1.4 1.4m-5.2 0c-.8 0-1.4-.6-1.4-1.4c0-.8.6-1.4 1.4-1.4c.8 0 1.4.6 1.4 1.4c0 .7-.7 1.4-1.4 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OutdentLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 16h18v2H1zm8-9h10v2H9zm0 4h10v2H9zM1 2h18v2H1zm0 8l5 4V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OutdentRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 16h18v2H1zm0-9h10v2H1zm0 4h10v2H1zm0-9h18v2H1zm13 4v8l5-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OutlineLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 12h18v7H1zM1 1v8h8V1zm6 6H3V3h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OutlineRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 12h18v7H1zM11 1v8h8V1zm2 2h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageSettings(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="10" cy="10" r="1.75" fill="currentColor"/><path fill="currentColor" d="M15 1H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2m0 9.75l-1.37.25a3.73 3.73 0 0 1-.38.93l.82 1.07L13 14.07l-1.12-.82a3.73 3.73 0 0 1-.93.38l-.2 1.37h-1.5L9 13.63a3.73 3.73 0 0 1-.93-.38L7 14.07L5.93 13l.82-1.12a3.73 3.73 0 0 1-.38-.88L5 10.75v-1.5L6.37 9a3.72 3.72 0 0 1 .38-.93L5.93 7L7 5.93l1.12.82A3.73 3.73 0 0 1 9 6.37L9.25 5h1.5L11 6.37a3.74 3.74 0 0 1 .93.38L13 5.93L14.07 7l-.82 1.12a3.73 3.73 0 0 1 .38.93l1.37.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaletteLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 15v2a3 3 0 0 1-3 3a10 10 0 1 1 10-10a5 5 0 0 1-5 5ZM3 8.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m3-4a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m5 0a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m3 4a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaletteRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 15v2a3 3 0 0 0 3 3A10 10 0 1 0 0 10a5 5 0 0 0 5 5ZM3 8.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 1 1-3 0m3-4a1.5 1.5 0 1 1 3 0a1.5 1.5 0 1 1-3 0m5 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 1 1-3 0m3 4a1.5 1.5 0 1 1 3 0a1.5 1.5 0 1 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PasteLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h8v2h2V3c0-1.1-.895-2-2-2H3c-1.1 0-2 .895-2 2v8c0 1.1.895 2 2 2h2v-2H3zm4 12v2c0 1.1.895 2 2 2h8c1.1 0 2-.895 2-2V9c0-1.1-.895-2-2-2h-2v2h2v8H9v-2z"/><path fill="currentColor" d="M10 5H8v3H5v2h3v3h2v-3h3V8h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PasteRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3H9v2H7V3c0-1.1.895-2 2-2h8c1.1 0 2 .895 2 2v8c0 1.1-.895 2-2 2h-2v-2h2zm-4 12v2c0 1.1-.895 2-2 2H3c-1.1 0-2-.895-2-2V9c0-1.1.895-2 2-2h2v2H3v8h8v-2z"/><path fill="currentColor" d="M10 5h2v3h3v2h-3v3h-2v-3H7V8h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="6" height="16" x="3" y="2" fill="currentColor" rx="1"/><rect width="6" height="16" x="11" y="2" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.55 19A1 1 0 0 1 3 18.13V1.87A1 1 0 0 1 4.55 1l12.2 8.13a1 1 0 0 1 0 1.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PreviousLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4 10l9 9l1.4-1.5L7 10l7.4-7.5L13 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PreviousRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1L5.6 2.5L13 10l-7.4 7.5L7 19l9-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1h10v4H5zM3 6a2 2 0 0 0-2 2v7h4v4h10v-4h4V8a2 2 0 0 0-2-2zm11 12H6v-6h8zm2-8a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushPin(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 8V2a2 2 0 0 0 2-2H5a2 2 0 0 0 2 2v6H6a2 2 0 0 0-2 2v1h5v5l1 4l1-4v-5h5v-1a2 2 0 0 0-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzleLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="17" cy="10" r="3" fill="currentColor"/><path fill="currentColor" d="M10.58 3A3 3 0 0 1 11 4.5a3 3 0 0 1-6 0A3 3 0 0 1 5.42 3H1v12a2 2 0 0 0 2 2h12V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzleRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="3" cy="10" r="3" fill="currentColor"/><path fill="currentColor" d="M9.42 3A2.94 2.94 0 0 0 9 4.5a3 3 0 0 0 6 0a2.94 2.94 0 0 0-.42-1.5H19v12a2 2 0 0 1-2 2H5V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 3H3v4h4zm6 0v4h4V3zM7 13H3v4h4zm8 0v-2h4v2h-2v4h-2v2h4v-4h-4v2h-4v2h2v-4h-2v-4h2v2zm-4-4V1h8v8zM1 9V1h8v8zm0 2h8v8H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuotesLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7 6l1-2H6C3.79 4 2 6.79 2 9v7h7V9H5c0-3 2-3 2-3m7 3c0-3 2-3 2-3l1-2h-2c-2.21 0-4 2.79-4 5v7h7V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuotesRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 9v7h7V9c0-2.2-1.8-5-4-5h-2l1 2s2 0 2 3zM2 9v7h7V9c0-2.2-1.8-5-4-5H3l1 2s2 0 2 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecentChangesLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3h16v2H1zm0 6h11v2H1zm0 6h7v2H1zm17.8-3.1l1-1.1a.6.6 0 0 0 0-.8L18 8.2a.6.6 0 0 0-.8 0l-1 1zm-3.3-2L10 15.3V18h2.6l5.6-5.5l-2.7-2.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecentChangesRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H3v2h16zm0 6h-7v2h7zm0 6H8v2h11zM8.8 11.9l1-1.1a.6.6 0 0 0 0-.8L8 8.2a.6.6 0 0 0-.8 0l-1 1L8.7 12zm-3.3-2L0 15.3V18h2.6l5.6-5.5l-2.7-2.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedoLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 8.5L12 3v11zM12 7v3h-1c-4 0-7 2-7 6v1H1v-1c0-6 5-9 10-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedoRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8.5L8 3v11zM8 7v3h1c4 0 7 2 7 6v1h3v-1c0-6-5-9-10-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reference(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15 10l-2.78-2.78L9.44 10V1H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReferenceExistingLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0a2 2 0 0 0-2 2h9a2 2 0 0 1 2 2v12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2z"/><path fill="currentColor" d="m13 12l-2.8-2.8L7.4 12V3H4c-1.1 0-2 .9-2 2v13c0 1.1.9 2 2 2h9c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReferenceExistingRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0a2 2 0 0 1 2 2H6a2 2 0 0 0-2 2v12a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M5 18c0 1.1.9 2 2 2h9c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2v9l-2.8-2.8l-2.8 2.8V3H7C5.4 3 5 4.6 5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReferencesLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3v16h5V3zm4 12H1v-1h3zm0-3H1v-1h3zm2-9v16h5V3zm4 12H7v-1h3zm0-3H7v-1h3zm1-8.5l4.1 15.4l4.8-1.3l-4-15.3zm7 10.6l-2.9.8l-.3-1l2.9-.8zm-.8-2.9l-2.9.8l-.2-1l2.9-.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReferencesRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3v16h5V3zm1 11h3v1h-3zm0-3h3v1h-3zM9 3v16h5V3zm1 11h3v1h-3zm0-3h3v1h-3zM4.1 2.3l-4 15.3l4.8 1.3L9 3.5zM2.3 13.1l2.9.8l-.3 1l-2.9-.8zm.7-2.9l2.9.8l-.2 1l-2.9-.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reload(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.65 4.35A8 8 0 1 0 17.4 13h-2.22a6 6 0 1 1-1-7.22L11 9h7V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Restore(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.22 0L0 1.22l4 4V17a2 2 0 0 0 2 2h8a2 2 0 0 0 2-1.8l2.8 2.8l1.2-1.22zM17 4V2h-3.5l-1-1h-5l-1 1h-.84l2 2zM8.66 5H16v7.34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Robot(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 5h6.505C18.107 5 19 5.896 19 6.997V14h-7v2h5.005c1.102 0 1.995.888 1.995 2v2H1v-2c0-1.105.893-2 1.995-2H8v-2H1V6.997C1 5.894 1.893 5 2.995 5H9.5V2.915a1.5 1.5 0 1 1 1 0zm-4 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m7 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sandbox(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12V9l6-6l3 3l-6 6zm10-7l-3-3l2-2l3 3zM8 2h2v2H8zM4 2h2v2H4zM0 3a1 1 0 0 1 1-1h1v2H0zm0 3h2v2H0zm0 4h2v2H0zm0 4h2v2H0zm0 4h2v2H1a1 1 0 0 1-1-1zm4 0h2v2H4zm4 0h2v2H8zm4 0h2v1a1 1 0 0 1-1 1h-1zm0-4h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.2 13.6a7 7 0 1 1 1.4-1.4l5.4 5.4l-1.4 1.4zM3 8a5 5 0 1 0 10 0A5 5 0 0 0 3 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchCaseSensitive(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.59 15.87h-1.52a.64.64 0 0 1-.42-.13a.84.84 0 0 1-.24-.32l-1-2.67H3.18l-1 2.67a.8.8 0 0 1-.23.31a.63.63 0 0 1-.42.14H0L4.8 3.76h2zm-3.72-4.54L6.2 6.91a12.12 12.12 0 0 1-.41-1.3q-.09.4-.2.73c-.07.22-.14.42-.2.58l-1.67 4.41zm5.58-2.84a4.91 4.91 0 0 1 3.46-1.35a3.41 3.41 0 0 1 1.32.24a2.62 2.62 0 0 1 1 .68a3 3 0 0 1 .6 1a4.08 4.08 0 0 1 .17 1.36v5.45h-.81a.78.78 0 0 1-.39-.08a.61.61 0 0 1-.23-.32l-.18-.7a7.87 7.87 0 0 1-.65.53a4.12 4.12 0 0 1-.66.39a3.3 3.3 0 0 1-.73.24a4.3 4.3 0 0 1-.86.08a3.18 3.18 0 0 1-1-.14a2.12 2.12 0 0 1-.78-.43a2 2 0 0 1-.52-.72a2.48 2.48 0 0 1-.19-1a2 2 0 0 1 .26-1a2.42 2.42 0 0 1 .87-.85a5.66 5.66 0 0 1 1.6-.62a11.7 11.7 0 0 1 2.51-.25v-.57A2.06 2.06 0 0 0 17.85 9a1.46 1.46 0 0 0-1.16-.45a2.53 2.53 0 0 0-.87.13a3.9 3.9 0 0 0-.62.32l-.46.28a.77.77 0 0 1-.43.13a.52.52 0 0 1-.32-.1a.81.81 0 0 1-.21-.24zm4.79 3.63a11.49 11.49 0 0 0-1.63.15a4.61 4.61 0 0 0-1.08.31a1.42 1.42 0 0 0-.59.45a1 1 0 0 0-.18.57a1.25 1.25 0 0 0 .1.52a.94.94 0 0 0 .27.35a1.08 1.08 0 0 0 .4.2a1.93 1.93 0 0 0 .51.06a2.59 2.59 0 0 0 1.21-.27a3.79 3.79 0 0 0 1-.77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchDiacritics(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.3 7.9a7.3 7.3 0 0 1 5.1-2a5 5 0 0 1 2 .3a4 4 0 0 1 1.4 1a4.4 4.4 0 0 1 .9 1.6a6 6 0 0 1 .3 2v8h-1.2a1.2 1.2 0 0 1-.6-.1a1 1 0 0 1-.3-.5l-.3-1a11.5 11.5 0 0 1-1 .8a6 6 0 0 1-1 .6a4.8 4.8 0 0 1-1 .3a6.4 6.4 0 0 1-1.3.1a4.7 4.7 0 0 1-1.4-.2a3.1 3.1 0 0 1-1.2-.7a3 3 0 0 1-.7-1a3.7 3.7 0 0 1-.3-1.5a2.9 2.9 0 0 1 .4-1.4a3.6 3.6 0 0 1 1.3-1.3a8.4 8.4 0 0 1 2.4-.9a17.2 17.2 0 0 1 3.6-.4v-.9a3 3 0 0 0-.6-2a2.1 2.1 0 0 0-1.7-.7a3.8 3.8 0 0 0-1.3.2a5.9 5.9 0 0 0-.9.4l-.7.4a1.1 1.1 0 0 1-.6.2A.8.8 0 0 1 6 9a1.2 1.2 0 0 1-.3-.4zm6.2-5.8a.9.9 0 0 0 .9-1.1H14a3.8 3.8 0 0 1-.2 1.2a2.7 2.7 0 0 1-.5.9a2.2 2.2 0 0 1-.7.5a2.3 2.3 0 0 1-1 .2a2 2 0 0 1-.8-.1a6.5 6.5 0 0 1-.8-.4L9.4 3a1.2 1.2 0 0 0-.5-.2a.8.8 0 0 0-.7.3a1.1 1.1 0 0 0-.2.8H6.4a3.7 3.7 0 0 1 .1-1.2a2.8 2.8 0 0 1 .5-.9a2.3 2.3 0 0 1 .8-.6a2.2 2.2 0 0 1 1-.2a2 2 0 0 1 .8.2a6.4 6.4 0 0 1 .7.3zm1 11.2a17 17 0 0 0-2.5.2a6.9 6.9 0 0 0-1.6.4a2.1 2.1 0 0 0-.9.7a1.4 1.4 0 0 0-.2.9a1.9 1.9 0 0 0 0 .7a1.4 1.4 0 0 0 .5.6a1.6 1.6 0 0 0 .6.2a2.9 2.9 0 0 0 .7.1a3.8 3.8 0 0 0 1.8-.4a5.6 5.6 0 0 0 1.5-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchRegularExpression(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.6 2a14.5 14.5 0 0 1 0 16l-.9-.56c-.28-.18-.17-.49-.05-.7a13.97 13.97 0 0 0-.09-13.6c-.08-.18-.16-.44.17-.65zM2.4 2c-2.9 3.87-3.38 10.9.03 15.94l.91-.56c.28-.17.16-.5.04-.7C.57 11.65 1.49 6.55 3.43 3.1c.08-.18.12-.38-.2-.59zM12 4h1v2.41l-.1.35l.35-.34l1.98-1.15l.54.94l-2.02 1.15l-.43.13l.43.12L15.8 8.8l-.54.94l-1.98-1.14l-.38-.35l.1.43V11h-1V8.76l.12-.52l-.34.34L9.8 9.72l-.54-.94l2-1.15l.48-.14l-.47-.13L9.2 6.2l.54-.94l1.97 1.14l.38.37l-.09-.41z"/><circle cx="6.5" cy="13.5" r="1.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g transform="translate(10 10)"><path id="oouiSettings0" fill="currentColor" d="M1.5-10h-3l-1 6.5h5m0 7h-5l1 6.5h3"/><use href="#oouiSettings0" transform="rotate(45)"/><use href="#oouiSettings0" transform="rotate(90)"/><use href="#oouiSettings0" transform="rotate(135)"/></g><path fill="currentColor" d="M10 2.5a7.5 7.5 0 0 0 0 15a7.5 7.5 0 0 0 0-15v4a3.5 3.5 0 0 1 0 7a3.5 3.5 0 0 1 0-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6V2l7 7l-7 7v-4c-5 0-8.5 1.5-11 5l.8-3l.2-.4A12 12 0 0 1 12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignatureLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 18h20v1H0zm-.003-6.155l1.06-1.06l4.363 4.362l-1.06 1.06z"/><path fill="currentColor" d="m.004 15.147l4.363-4.363l1.06 1.061l-4.362 4.363zM17 5c0 9-11 9-11 9v-1.5s8 .5 9.5-6.5C16 4 15 2.5 14 2.5S11 4 10.75 10c-.08 2 .75 4.5 3.25 4.5c1.5 0 2-1 3.5-1a2.07 2.07 0 0 1 2.25 2.5h-1.5s.13-1-.5-1C16 15 16 16 14 16c0 0-4.75 0-4.75-6S12 1 14 1c.5 0 3 0 3 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignatureRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 18h20v1H0zm14.542-2.883l4.384-4.384l1.06 1.06l-4.384 4.384z"/><path fill="currentColor" d="m14.54 11.86l1.06-1.062l4.384 4.384l-1.06 1.061zM6 1c2 0 4.8 3 4.8 9S6 16 6 16c-2 0-2-1-3.8-1c-.6 0-.5 1-.5 1H.2c0-.2-.1-.4 0-.7c.1-1.1 1.1-2 2.3-1.8c1.5 0 2 1 3.5 1c2.5 0 3.3-2.5 3.3-4.5C9 4 7 2.5 6 2.5S4 4 4.5 6C6 13 14 12.5 14 12.5V14S3 14 3 5c0-4 2.5-4 3-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smaller(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h-1.05a.44.44 0 0 1-.29-.09a.58.58 0 0 1-.17-.22l-.7-1.84H6.2l-.7 1.84a.56.56 0 0 1-.16.21a.43.43 0 0 1-.29.1H4l3.31-8.35h1.38zm-2.57-3.13L8.28 9.82a8.5 8.5 0 0 1-.28-.9q-.06.27-.14.5l-.14.4l-1.15 3zM15 6l3-4h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallerText(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.75 18h-1.51a.64.64 0 0 1-.42-.13a.83.83 0 0 1-.24-.32l-1-2.65H7.41l-1 2.65a.79.79 0 0 1-.23.31a.62.62 0 0 1-.42.14H4.25L9 6h2zm-3.69-4.5L10.4 9.12a12.13 12.13 0 0 1-.4-1.3q-.09.39-.2.72t-.2.58L7.95 13.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpecialCharacter(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 15.9v1.29a.77.77 0 0 1-.23.58a.86.86 0 0 1-.63.23h-6.76v-2.87a4.41 4.41 0 0 0 1.74-.71a5.51 5.51 0 0 0 1.4-1.42a6.92 6.92 0 0 0 .93-1.91a7.47 7.47 0 0 0 .34-2.28a6.15 6.15 0 0 0-.47-2.48a5.1 5.1 0 0 0-1.26-1.78a5.2 5.2 0 0 0-1.85-1.07a7.15 7.15 0 0 0-4.43 0a5.08 5.08 0 0 0-3.11 2.87a6.08 6.08 0 0 0-.47 2.48a7.47 7.47 0 0 0 .34 2.28A6.81 6.81 0 0 0 5.47 13a5.59 5.59 0 0 0 1.41 1.39a4.41 4.41 0 0 0 1.74.71V18H1.86a.86.86 0 0 1-.63-.23a.77.77 0 0 1-.23-.58V15.9h4.76l1 .12a6.94 6.94 0 0 1-2-1.05a7.39 7.39 0 0 1-1.58-1.63a7.75 7.75 0 0 1-1-2.1a8 8 0 0 1-.38-2.47a7.61 7.61 0 0 1 .65-3.17A7.48 7.48 0 0 1 4.1 3.17a8.14 8.14 0 0 1 2.65-1.6A9.19 9.19 0 0 1 10 1a9.18 9.18 0 0 1 3.25.57a8.14 8.14 0 0 1 2.65 1.6a7.48 7.48 0 0 1 1.78 2.47a7.61 7.61 0 0 1 .65 3.17a8 8 0 0 1-.33 2.48a7.74 7.74 0 0 1-1 2.1A7.37 7.37 0 0 1 15.33 15a7 7 0 0 1-2 1.05l1-.12h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpecialPagesLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0a2 2 0 0 0-2 2h9a2 2 0 0 1 2 2v12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2z"/><path fill="currentColor" d="M13 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2m-6.5-3.5l.41-1.09L8 15l-1.09-.41l-.41-1.09l-.41 1.09L5 15l1.09.41zm2.982-.949l.952-2.561l2.53-.964l-2.53-.964L9.482 8.5l-.952 2.562l-2.53.964l2.53.964zM6 10.5l.547-1.453L8 8.5l-1.453-.547L6 6.5l-.547 1.453L4 8.5l1.453.547z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpecialPagesRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0a2 2 0 0 1 2 2H6a2 2 0 0 0-2 2v12a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M7 20a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2zm6.5-3.5l.41-1.09L15 15l-1.09-.41l-.41-1.09l-.41 1.09L12 15l1.09.41zm-2.982-.949l.952-2.561l2.53-.964l-2.53-.964l-.952-2.562l-.952 2.562l-2.53.964l2.53.964zM14 10.5l.547-1.453L16 8.5l-1.453-.547L14 6.5l-.547 1.453L12 8.5l1.453.547z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeechBubbleAddLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1a2 2 0 0 0-2 2v16l4-4h12a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm12 8h-4v4H9V9H5V7h4V3h2v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeechBubbleAddRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3v10c0 1.1.9 2 2 2h12l4 4V3c0-1.1-.9-2-2-2H3c-1.1 0-2 .9-2 2m4 4h4V3h2v4h4v2h-4v4H9V9H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeechBubbleLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 14H0v6z"/><rect width="20" height="16" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeechBubbleRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 14h6v6z"/><rect width="20" height="16" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeechBubblesLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4v7a2 2 0 0 1-2 2H4v1a2 2 0 0 0 2 2h10l4 4V6a2 2 0 0 0-2-2zM6 10H0v6z"/><rect width="16" height="12" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeechBubblesRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 20l4-4h10c1.1 0 2-.9 2-2v-1H5c-1.1 0-2-.9-2-2V4H2C.9 4 0 4.9 0 6zm14-10h6v6z"/><rect width="16" height="12" x="4" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 7h-7L10 .5L7 7H0l5.46 5.47l-1.64 7l6.18-3.7l6.18 3.73l-1.63-7zm-10 6.9l-3.76 2.27l1-4.28L3.5 8.5h4.61L10 4.6l1.9 3.9h4.6l-3.73 3.4l1 4.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="16" height="16" x="2" y="2" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrikethroughA(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 8h2.6L10 4l1.4 4h2.7l-2.4-6H8.3zm-5 2v2h3.4L2 18h3l1.7-4.6h6.6L15 18h3l-2.4-6H19v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrikethroughS(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.6 9h5.9l-.4-.2l-1.4-.4a7 7 0 0 1-1.3-.6a2.9 2.9 0 0 1-.9-.7a1.6 1.6 0 0 1-.3-1a2.3 2.3 0 0 1 .2-1a2 2 0 0 1 .5-.7a2.5 2.5 0 0 1 .9-.5a3.9 3.9 0 0 1 1.2-.1a3.9 3.9 0 0 1 1.3.2a5.8 5.8 0 0 1 1 .4l.6.4a.9.9 0 0 0 .5.2a.6.6 0 0 0 .3-.1a1 1 0 0 0 .3-.3l.7-1.3a6 6 0 0 0-2-1.3a7.4 7.4 0 0 0-2.7-.5a6.1 6.1 0 0 0-2.3.4A5 5 0 0 0 6.1 3a4.5 4.5 0 0 0-1.4 3.2A4.7 4.7 0 0 0 5 8.1a4 4 0 0 0 .6.9M19 11H1v2h11.4a2.3 2.3 0 0 1 0 .6a2.5 2.5 0 0 1-.7 2a3.3 3.3 0 0 1-2.3.7A4.3 4.3 0 0 1 8 16a6 6 0 0 1-1.1-.5l-.8-.6a1 1 0 0 0-.6-.2a.7.7 0 0 0-.4 0a.8.8 0 0 0-.2.3L4 16.3a6.4 6.4 0 0 0 1 1a7.2 7.2 0 0 0 1.4.6a8 8 0 0 0 1.4.5a7.7 7.7 0 0 0 1.5.1a6.5 6.5 0 0 0 2.4-.4a5.1 5.1 0 0 0 1.8-1.1a4.9 4.9 0 0 0 1.1-1.7a5.6 5.6 0 0 0 .4-2.1V13h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrikethroughY(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.3 8h2.3l-4-6H3.3zm3.1 0h2.3l4-6h-2.3zM1 10v2h8v6h2v-6h8v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubscriptLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.68 16h-2.42a.67.67 0 0 1-.46-.15a1.33 1.33 0 0 1-.28-.34l-2.77-4.44a2.65 2.65 0 0 1-.28.69L5 15.51a2.22 2.22 0 0 1-.29.34a.58.58 0 0 1-.42.15H2l4.15-6.19L2.17 4h2.42a.81.81 0 0 1 .41.09a.8.8 0 0 1 .24.26L8 8.59a2.71 2.71 0 0 1 .33-.74L10.6 4.4a.69.69 0 0 1 .6-.4h2.32l-4 5.71zm3.82-4h.5v-1h-.5a1.49 1.49 0 0 0-1 .39a1.49 1.49 0 0 0-1-.39H15v1h.5a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5H15v1h.5a1.49 1.49 0 0 0 1-.39a1.49 1.49 0 0 0 1 .39h.5v-1h-.5a.5.5 0 0 1-.5-.5v-6a.5.5 0 0 1 .5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubscriptRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 9.7L6.5 4h2.3c.3 0 .5.2.6.4l2.3 3.5c.1.2.3.5.3.7l2.8-4.2c.1-.1.1-.2.2-.3c.1-.1.3-.1.4-.1h2.4l-4 5.8L18 16h-2.3c-.2 0-.3 0-.4-.1l-.3-.3l-2.5-3.8c-.1-.2-.2-.4-.3-.7l-2.8 4.4l-.3.3c0 .2-.2.2-.4.2H6.3zM3 18.5c0 .3-.2.5-.5.5H2v1h.5c.4 0 .7-.1 1-.4c.3.2.6.4 1 .4H5v-1h-.5c-.3 0-.5-.2-.5-.5v-6c0-.3.2-.5.5-.5H5v-1h-.5c-.4 0-.7.1-1 .4c-.3-.2-.6-.4-1-.4H2v1h.5c.3 0 .5.2.5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subtract(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 9h12v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Success(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 20a10 10 0 0 1 0-20a10 10 0 1 1 0 20m-2-5l9-8.5L15.5 5L8 12L4.5 8.5L3 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SuperscriptLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 1V0h-.5a1.49 1.49 0 0 0-1 .39a1.49 1.49 0 0 0-1-.39H15v1h.5a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5H15v1h.5a1.49 1.49 0 0 0 1-.39a1.49 1.49 0 0 0 1 .39h.5V8h-.5a.5.5 0 0 1-.5-.5v-6a.5.5 0 0 1 .5-.5zm-4.32 15h-2.42a.67.67 0 0 1-.46-.15a1.33 1.33 0 0 1-.28-.34l-2.77-4.44a2.65 2.65 0 0 1-.28.69L5 15.51a2.22 2.22 0 0 1-.29.34a.58.58 0 0 1-.42.15H2l4.15-6.19L2.17 4h2.42a.81.81 0 0 1 .41.09a.8.8 0 0 1 .24.26L8 8.59a2.71 2.71 0 0 1 .33-.74L10.6 4.4a.69.69 0 0 1 .6-.4h2.32l-4 5.71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SuperscriptRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7.5c0 .3-.2.5-.5.5H2v1h.5c.4 0 .7-.1 1-.4c.3.3.6.4 1 .4H5V8h-.5c-.3 0-.5-.2-.5-.5v-6c0-.3.2-.5.5-.5H5V0h-.5c-.4 0-.7.1-1 .4c-.3-.3-.6-.4-1-.4H2v1h.5c.3 0 .5.2.5.5zm7.5 2.2L6.5 4h2.3c.3 0 .5.2.6.4l2.3 3.5c.1.2.3.5.3.7l2.8-4.2c.1-.1.1-.2.2-.3c.1-.1.3-.1.4-.1h2.4l-4 5.8L18 16h-2.3c-.2 0-.3 0-.4-.1l-.3-.3l-2.5-3.8c-.1-.2-.2-.4-.3-.7l-2.8 4.4l-.3.3c0 .2-.2.2-.4.2H6.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm0 4h7v4H2zm0 10v-4h7v4zm16 0h-7v-4h7zm0-6h-7V6h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableAddColumnAfterLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3v14h8v1h12V2H8v1zm10 6h3V6h2v3h3v2h-3v3h-2v-3h-3zM6 5h2v10H6zM2 5h2v10H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableAddColumnAfterRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3h-6V2H0v16h12v-1h8V3zm-8 8H7v3H5v-3H2V9h3V6h2v3h3zm4 4h-2V5h2zm4 0h-2V5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableAddColumnBeforeLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3h-6V2H0v16h12v-1h8V3zm-8 8H7v3H5v-3H2V9h3V6h2v3h3zm4 4h-2V5h2zm4 0h-2V5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableAddColumnBeforeRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3v14h8v1h12V2H8v1zm10 6h3V6h2v3h3v2h-3v3h-2v-3h-3zM6 5h2v10H6zM2 5h2v10H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableAddRowAfter(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0v8H2v12h16V8h-1V0zm8 10v3h3v2h-3v3H9v-3H6v-2h3v-3zm4-4v2H5V6zm0-4v2H5V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableAddRowBefore(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 20v-8h1V0H2v12h1v8zM9 10V7H6V5h3V2h2v3h3v2h-3v3zm-4 4v-2h10v2zm0 4v-2h10v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableCaption(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2zm0 2h7v2H2zm0 6v-2h7v2zm16 0h-7v-2h7zm0-4h-7v-2h7zM2 2h16v4H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableMergeCells(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path id="oouiTableMergeCells0" fill="currentColor" d="M9 10L4 6v3H0v2h4v3zm-7 3H0v5h8v-2H2zM0 2v5h2V4h6V2z"/><use href="#oouiTableMergeCells0" transform="matrix(-1 0 0 1 20 0)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableMoveColumnAfterLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 10l-5-4v3H6v2h5v3z"/><path fill="currentColor" d="M0 2h20v16H0zm5 6v4h5v4h8V4h-8v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableMoveColumnAfterRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4 10l5-4v3h5v2H9v3z"/><path fill="currentColor" d="M0 2v16h20V2zm2 2h8v4h5v4h-5v4H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableMoveColumnBeforeLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4 10l5-4v3h5v2H9v3z"/><path fill="currentColor" d="M0 2v16h20V2zm2 2h8v4h5v4h-5v4H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableMoveColumnBeforeRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 10l-5-4v3H6v2h5v3z"/><path fill="currentColor" d="M0 2h20v16H0zm5 6v4h5v4h8V4h-8v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableMoveRowAfter(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10 16l-4-5h3V6h2v5h3z"/><path fill="currentColor" d="M2 0v20h16V0zm2 10h4V5h4v5h4v8H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableMoveRowBefore(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9H6l4-5l4 5h-3v5H9z"/><path fill="currentColor" d="M2 0h16v20H2zm2 2v8h4v5h4v-5h4V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1.28A1 1 0 0 0 8.35 1H2a1 1 0 0 0-1 1v6.35a1 1 0 0 0 .28.65L11 18.72a1 1 0 0 0 1.37 0l6.38-6.38a1 1 0 0 0-.03-1.34zM5 7a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.6 18.7c.4.4 1 .4 1.4 0L18.7 9c.2-.2.3-.4.3-.6V2c0-.6-.4-1-1-1h-6.4c-.2 0-.5.1-.6.3L1.3 11c-.4.4-.4 1-.1 1.3zM13 5c0-1.1.9-2 2-2s2 .9 2 2s-.9 2-2 2s-2-.9-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemplateAddLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5V1h-2v4h-4v2h4v4h2V7h4V5z"/><path fill="currentColor" d="M0 17V5h8v2H2v8h12v-2h2v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemplateAddRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7h4v4h2V7h4V5H6V1H4v4H0z"/><path fill="currentColor" d="M4 13h2v2h12V7h-6V5h8v12H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextDirLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19 10l-6-5v4H6v2h7v4zM6 2V1H4.5a1.49 1.49 0 0 0-1 .39a1.49 1.49 0 0 0-1-.39H1v1h1.5a.5.5 0 0 1 .5.5v15a.5.5 0 0 1-.5.5H1v1h1.5a1.49 1.49 0 0 0 1-.39a1.49 1.49 0 0 0 1 .39H6v-1H4.5a.5.5 0 0 1-.5-.5v-15a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextDirRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1 10l6-5v4h7v2H7v4zm13-8V1h1.5a1.49 1.49 0 0 1 1 .39a1.49 1.49 0 0 1 1-.39H19v1h-1.5a.5.5 0 0 0-.5.5v15a.5.5 0 0 0 .5.5H19v1h-1.5a1.49 1.49 0 0 1-1-.39a1.49 1.49 0 0 1-1 .39H14v-1h1.5a.5.5 0 0 0 .5-.5v-15a.5.5 0 0 0-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextFlowLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3h18v2H1zm0 4h14v2H1zm0 4h10v2H1zm0 4h18v2H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextFlowRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3h18v2H1zm4 4h14v2H5zm4 4h10v2H9zm-8 4h18v2H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextStyle(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 17h16v2H2zm9.34-15h3.31l2 14h-3.19l-.29-2.88H8L6.43 16H3.37zm-2 8.71h3.55l-.61-5.51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSummaryLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 7h18v2H1zm0 4h14v2H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSummaryRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 7h18v2H1zm4 4h14v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-3.5l-1-1h-5l-1 1H3v2h14zM4 17a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V5H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tray(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm14 12h-4l-1 2H8l-1-2H3V3h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnBlock(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.22 0L0 1.22l3.06 3.06a9 9 0 0 0 12.66 12.66L18.78 20L20 18.78zM5 11V9h2.78l2 2zm5-10a9 9 0 0 0-4.26 1.08L12.66 9H15v2h-.34l3.26 3.26A9 9 0 0 0 10 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnFlagLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.14 8.48L17 6L5.58 1.92zM1.22 0L0 1.22l3 3V19h2v-6.87l3.91-2L18.78 20L20 18.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnFlagRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 1.2l4.27 4.27L0 7l11.84 6.04l.16.16V20h2v-4.8l4.74 4.74l1.198-1.198L1.198.002zM14 2L7.809 4.209L14 10.399z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnLink(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.83 5A4.83 4.83 0 0 0 0 9.83v.34A4.83 4.83 0 0 0 4.83 15h2.91a4.88 4.88 0 0 1-1.55-2H5c-4 0-4-6 0-6h3c.075.001.15.005.225.012L6.215 5zm7.43 0a4.88 4.88 0 0 1 1.55 2H15c3.179.003 4.17 4.3 1.314 5.695l1.508 1.508A4.83 4.83 0 0 0 20 10.17v-.34A4.83 4.83 0 0 0 15.17 5zm-3.612.03l4.329 4.327A4.83 4.83 0 0 0 8.648 5.03M7.227 8.411C7.17 8.595 7.08 9 7.08 9c-.045.273-.08.584-.08.83v.34A4.83 4.83 0 0 0 11.83 15h3.34c.316 0 .631-.032.941-.094L14.205 13H12c-2.067-.006-3.51-2.051-2.82-4zm3.755 1.36A3 3 0 0 1 10.82 11h1.389z"/><path fill="currentColor" d="M1.22 0L0 1.22L18.8 20l1.2-1.22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnLock(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 8V5s0-5-5-5a4.63 4.63 0 0 0-4.88 4h2C7.31 2.93 8 2 10 2c3 0 3 2 3 3.5V8H3.93A1.93 1.93 0 0 0 2 9.93v8.15A1.93 1.93 0 0 0 3.93 20h12.14A1.93 1.93 0 0 0 18 18.07V9.93A1.93 1.93 0 0 0 16.07 8zm-5 8a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnStar(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 7h-7L10 .5L7 7H0l5.46 5.47l-1.64 7l6.18-3.7l6.18 3.73l-1.63-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnderlineA(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 17h14v2H3zm4.7-6.7L10 3.7l2.3 6.6zm6.6 5.7H17L11.5 2h-3L3 16h2.7L7 12h5.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnderlineU(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 17h14v2H3zm2.6-2.7a5.5 5.5 0 0 0 1.9 1.2a6.6 6.6 0 0 0 2.5.5a6.6 6.6 0 0 0 2.5-.5a5.4 5.4 0 0 0 3-3.1A6.8 6.8 0 0 0 16 10V2h-2.2v8a5 5 0 0 1-.3 1.7a3.7 3.7 0 0 1-.7 1.3a3.3 3.3 0 0 1-1.2.8a4 4 0 0 1-1.6.3a4 4 0 0 1-1.6-.3a3.3 3.3 0 0 1-1.2-.9a3.6 3.6 0 0 1-.7-1.3a5.2 5.2 0 0 1-.3-1.6V2H4v8a6.8 6.8 0 0 0 .4 2.4a5.5 5.5 0 0 0 1.2 1.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UndoLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8.5L8 14v-4h1c4 0 7 2 7 6v1h3v-1c0-6-5-9-10-9H8V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UndoRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3v4h-1C6 7 1 10 1 16v1h3v-1c0-4 3-6 7-6h1v4l7-5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpTriangle(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10 5l8 10H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 12v5H3v-5H1v5a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5z"/><path fill="currentColor" d="M10 1L5 7h4v8h2V7h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserActive(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12.5c-5.92 0-9 3.5-9 5.5v1h18v-1c0-2-3.08-5.5-9-5.5"/><circle cx="10" cy="6" r="5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAddLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="8.5" cy="10.5" r="3.5" fill="currentColor"/><path fill="currentColor" d="M14 0v4h-4v2h4v4h2V6h4V4h-4V0zM8 15c-4.6 0-7 2.69-7 4.23V20h14v-.77C15 17.69 12.6 15 8 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAddRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="11.5" cy="10.5" r="3.5" fill="currentColor"/><path fill="currentColor" d="M6 0v4h4v2H6v4H4V6H0V4h4V0zm6 15c4.6 0 7 2.69 7 4.23V20H5v-.77C5 17.69 7.4 15 12 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAnonymous(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2H5L4 8h12zM0 10s2 1 10 1s10-1 10-1l-4-2H4zm8 4h4v1H8z"/><circle cx="6" cy="15" r="3" fill="currentColor"/><circle cx="14" cy="15" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAvatar(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 11c-5.92 0-8 3-8 5v3h16v-3c0-2-2.08-5-8-5"/><circle cx="10" cy="5.5" r="4.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAvatarOutline(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8c1.7 0 3.06-1.35 3.06-3S11.7 2 10 2S6.94 3.35 6.94 5S8.3 8 10 8m0 2c-2.8 0-5.06-2.24-5.06-5S7.2 0 10 0s5.06 2.24 5.06 5s-2.26 5-5.06 5m-7 8h14v-1.33c0-1.75-2.31-3.56-7-3.56s-7 1.81-7 3.56zm7-6.89c6.66 0 9 3.33 9 5.56V20H1v-3.33c0-2.23 2.34-5.56 9-5.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserContributionsLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="15.5" cy="10.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="M1 15h8v2H1Zm0-6h10v2H1Zm0-6h16v2H1Zm14.5 10.6c-3.3 0-4.5 1.6-4.5 2.7V18h9v-1.7c0-1-1.2-2.7-4.5-2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserContributionsRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="4.5" cy="10.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="M19 15h-8v2h8zm0-6H9v2h10Zm0-6H3v2h16ZM4.5 13.6c3.3 0 4.5 1.6 4.5 2.7V18H0v-1.7c0-1 1.2-2.7 4.5-2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserGroupLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="6" cy="6" r="3" fill="currentColor"/><circle cx="14" cy="6" r="3" fill="currentColor"/><path fill="currentColor" d="M14 10c3.31 0 6 1.79 6 4v2h-6v-2c0-1.48-1.21-2.77-3-3.46c.88-.35 1.91-.54 3-.54m-8 0c3.31 0 6 1.79 6 4v2H0v-2c0-2.21 2.69-4 6-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserGroupRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="6" cy="6" r="3" fill="currentColor"/><circle cx="14" cy="6" r="3" fill="currentColor"/><path fill="currentColor" d="M6 10c-3.31 0-6 1.79-6 4v2h6v-2c0-1.48 1.21-2.77 3-3.46c-.88-.35-1.91-.54-3-.54m8 0c-3.31 0-6 1.79-6 4v2h12v-2c0-2.21-2.69-4-6-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserRightsLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="10" cy="5.5" r="4.5" fill="currentColor"/><path fill="currentColor" d="M20 16.288v-1.632l-1.012-.168c-.055-.22-.164-.437-.324-.76l-.014-.028l.619-.844l-1.181-1.18l-.844.618a3.003 3.003 0 0 0-.788-.338l-.112-1.012h-1.632l-.141.851l-.027.161c-.282.056-.507.17-.788.338l-.844-.619l-1.18 1.181l.562.844c-.106.176-.168.33-.227.49l-.025.07l-.01.027c-.024.064-.048.13-.076.2L11 14.6v1.631l1.012.17c.057.28.17.505.338.786l-.563.844l.969.969l.213.213l.32-.213l.524-.35c.199.1.443.2.691.3l.05.02l.046.018l.002.012l.167 1h1.687l.167-1l.002-.012c.281-.057.506-.17.787-.338l.477.35l.367.27l1.182-1.182l-.62-.844c.17-.28.282-.563.338-.788Zm-4.5.843a1.659 1.659 0 0 1-1.688-1.687c0-.956.732-1.688 1.688-1.688s1.688.732 1.688 1.688s-.732 1.687-1.688 1.687M10 11c.739 0 1.418.047 2.04.133l-1.594 1.596l.574.862l-1.02.12v3.367l1.098.183l-.598.897l.842.842H2v-3c0-2 2.083-5 8-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserRightsRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="10" cy="5.5" r="4.5" fill="currentColor"/><path fill="currentColor" d="M0 16.288v-1.632l1.012-.168c.055-.22.164-.437.324-.76l.014-.028l-.619-.844l1.181-1.18l.844.618c.245-.15.51-.264.788-.338l.112-1.012h1.632l.141.851l.027.161c.282.056.507.17.788.338l.844-.619l1.18 1.181l-.562.844c.106.176.168.33.227.49l.025.07l.01.027c.024.064.048.13.076.2L9 14.6v1.631l-1.012.17c-.057.28-.17.505-.338.786l.563.844l-.969.969l-.213.213l-.32-.213l-.524-.35c-.199.1-.443.2-.691.3l-.05.02l-.046.018l-.002.012l-.167 1H3.544l-.167-1l-.002-.012c-.281-.057-.506-.17-.787-.338l-.477.35l-.367.27l-1.182-1.182l.62-.844a2.958 2.958 0 0 1-.338-.788Zm4.5.843a1.658 1.658 0 0 0 1.688-1.687c0-.956-.732-1.688-1.688-1.688s-1.688.732-1.688 1.688s.732 1.687 1.688 1.687M10 11c-.739 0-1.418.047-2.04.133l1.594 1.596l-.574.862l1.02.12v3.367l-1.098.183l.598.897l-.842.842H18v-3c0-2-2.083-5-8-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTalkLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 0H2a2 2 0 0 0-2 2v18l4-4h14a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2m-4 4a1.5 1.5 0 1 1-1.5 1.5A1.5 1.5 0 0 1 14 4M6 4a1.5 1.5 0 1 1-1.5 1.5A1.5 1.5 0 0 1 6 4m4 8c-2.61 0-4.83-.67-5.65-3h11.3c-.82 2.33-3.04 3-5.65 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTalkRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v12c0 1.1.9 2 2 2h14l4 4V2c0-1.1-.9-2-2-2H2C.9 0 0 .9 0 2m7.5 3.5C7.5 6.3 6.8 7 6 7s-1.5-.7-1.5-1.5S5.2 4 6 4s1.5.7 1.5 1.5m8 0c0 .8-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5S13.2 4 14 4s1.5.7 1.5 1.5M4.4 9h11.3c-.8 2.3-3 3-5.6 3s-4.9-.7-5.7-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTemporaryLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 16c0-2-2.083-5-8-5c-.356 0-.698.01-1.026.032l.276 1.986c.239-.012.489-.018.75-.018v4H6v2h2v-2h2v2h8zm-8-6a4.5 4.5 0 1 0-.878-8.914l.388 1.962C9.667 3.017 9.83 3 10 3v5c-.17 0-.333-.017-.49-.048l-.388 1.962c.284.057.577.086.878.086m-2.933 1.29c-.69.149-1.3.35-1.834.59l1.051 1.713c.351-.14.755-.263 1.218-.36zM4 19v-2H2v2zm-1.84-4h2.228c.124-.18.284-.37.49-.557L3.64 12.866c-.78.66-1.254 1.416-1.48 2.134M7.5 1.758L8.613 3.42c-.274.183-.51.419-.693.693L6.258 3c.329-.49.751-.913 1.242-1.242M5.586 4.622l1.962.388a2.521 2.521 0 0 0 0 .98l-1.962.388a4.52 4.52 0 0 1 0-1.756M6.258 8L7.92 6.887c.183.274.419.51.693.693L7.5 9.242A4.525 4.525 0 0 1 6.258 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTemporaryRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 16c0-2 2.083-5 8-5c.356 0 .698.01 1.026.032l-.276 1.986A15.24 15.24 0 0 0 10 13v4h4v2h-2v-2h-2v2H2zm8-6a4.5 4.5 0 1 1 .878-8.914l-.388 1.962A2.521 2.521 0 0 0 10 3v5c.17 0 .333-.017.49-.048l.388 1.962A4.52 4.52 0 0 1 10 10m2.933 1.29c.69.149 1.3.35 1.834.59l-1.051 1.713a7.336 7.336 0 0 0-1.218-.36zM16 19v-2h2v2zm1.84-4h-2.228a3.204 3.204 0 0 0-.49-.557l1.238-1.577c.78.66 1.254 1.416 1.48 2.134M12.5 1.758L11.387 3.42c.274.183.51.419.693.693L13.742 3A4.525 4.525 0 0 0 12.5 1.758m1.914 2.864l-1.962.388a2.524 2.524 0 0 1 0 .98l1.962.388a4.521 4.521 0 0 0 0-1.756M13.742 8L12.08 6.887c-.183.274-.419.51-.693.693L12.5 9.242c.49-.329.913-.751 1.242-1.242"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalEllipsis(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="10" cy="10" r="2" fill="currentColor"/><circle cx="10" cy="3" r="2" fill="currentColor"/><circle cx="10" cy="17" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewCompact(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h4v4H2zm12 0h4v4h-4zM8 2h4v4H8zM2 14h4v4H2zm12 0h4v4h-4zm-6 0h4v4H8zM2 8h4v4H2zm12 0h4v4h-4zM8 8h4v4H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewDetailsLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6h9v2H8zm0-3h11v2H8zM1 3h6v6H1zm7 11h9v2H8zm0-3h11v2H8zm-7 0h6v6H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewDetailsRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8V6h9v2zm9-3H1V3h11zm1-2h6v6h-6zm-1 13H3v-2h9zm0-3H1v-2h11zm1-2h6v6h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisionSimulator(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 11.83a.79.79 0 0 1-.83.83h-3.34A1.49 1.49 0 0 1 11.67 11V9.33a.79.79 0 0 1 .83-.83h4.17a.79.79 0 0 1 .83.83zM8.33 11a1.49 1.49 0 0 1-1.67 1.67H3.33a.79.79 0 0 1-.83-.83V9.33a.79.79 0 0 1 .83-.83H7.5a.79.79 0 0 1 .83.83zM0 6.2v6.28a.2.2 0 0 0 .2.2h1.72a1.61 1.61 0 0 0 1.42.83h3.33A2.46 2.46 0 0 0 9.13 12a.19.19 0 0 1 .18-.13h1.39a.19.19 0 0 1 .18.13a2.46 2.46 0 0 0 2.46 1.53h3.33c.55 0 1.1 0 1.37-.7a.2.2 0 0 1 .18-.13h1.58a.2.2 0 0 0 .2-.2V6.2a.2.2 0 0 0-.2-.2H.2a.2.2 0 0 0-.2.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDownLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4 14l5.2 3.9c.3.3.8 0 .8-.5V2.6c0-.5-.5-.8-.8-.5L4 6m0 8H1c-.6 0-1-.4-1-1V7c0-.6.4-1 1-1h3m9.5 8.5c-.3 0-.5-.1-.7-.3c-.4-.4-.4-1 0-1.4c1.6-1.6 1.6-4.1 0-5.7c-.4-.4-.3-1.1.1-1.4c.4-.3.9-.3 1.3 0c2.3 2.3 2.3 6.1 0 8.5c-.1.2-.4.3-.7.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDownRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.925 14l-5.2 3.9c-.3.3-.8 0-.8-.5V2.6c0-.5.5-.8.8-.5l5.2 3.9m0 8h3c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1h-3m-9.5 8.5c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4c-1.6-1.6-1.6-4.1 0-5.7c.4-.4.3-1.1-.1-1.4c-.4-.3-.9-.3-1.3 0c-2.3 2.3-2.3 6.1 0 8.5c.1.2.4.3.7.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOffLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20 18.6l-2.3-2.3l-1.4-1.4l-1.4-1.4l-1.5-1.5L10 8.6l-4-4L1.4 0L0 1.4l4.3 4.3L4 6H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h3l5.2 3.9c.3.3.8 0 .8-.5v-6l5.6 5.6l2.9 2.9zm-10-16c0-.5-.5-.8-.8-.5L7.6 3.3L10 5.8zm5.7 1.7c2.4 2.4 2.9 5.9 1.7 8.8l1.5 1.5c2-3.8 1.4-8.5-1.8-11.7c-.4-.3-.9-.3-1.3 0c-.4.4-.5 1-.1 1.4m-2.9 2.9c.7.7 1.1 1.6 1.1 2.6l1.8 1.8c.5-2 0-4.2-1.5-5.8c-.4-.3-.9-.3-1.3 0s-.4 1-.1 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOffRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 18.6L18.6 20l-4.5-4.5l-3.3 2.4c-.3.3-.8 0-.8-.5v-6L6.6 8c-.9 1.5-.7 3.3.5 4.7l.2.2c.4.4.4 1 0 1.4c-.2.2-.4.3-.7.3c-.3 0-.6-.1-.7-.3c-2-2.1-2.3-5.4-.7-7.7L3.7 5.1c-2.5 3.1-2.2 7.7.6 10.6c.2.2.3.4.3.7c0 .5-.4 1-1 1c-.2 0-.5-.1-.7-.3C-.7 13.4-1 7.6 2.3 3.7L0 1.4L1.4 0m9.4 2.1L16 6h3c.6 0 .9.3 1 .9V13c0 .6-.3.9-.9 1h-.8L10 5.8V2.7c0-.6.5-.9.8-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUpLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6v8l5.2 3.9c.3.3.8 0 .8-.5V2.6c0-.5-.5-.8-.8-.5zm0 8H1a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h3m12.4 11.4a1 1 0 0 1-.7-1.7a8 8 0 0 0 0-11.4A1 1 0 0 1 17 3a10 10 0 0 1 0 14.2a1 1 0 0 1-.7.3z"/><path fill="currentColor" d="M13.5 14.5a1 1 0 0 1-.7-.3a1 1 0 0 1 0-1.4a4 4 0 0 0 0-5.6a1 1 0 0 1 1.4-1.4a6 6 0 0 1 0 8.4a1 1 0 0 1-.7.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUpRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6v8l-5.2 3.9c-.3.3-.8 0-.8-.5V2.6c0-.5.5-.8.8-.5zm0 8h3a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-3M3.6 17.4a1 1 0 0 0 .7-1.7a8 8 0 0 1 0-11.4A1 1 0 0 0 3 3a10 10 0 0 0 0 14.2a1 1 0 0 0 .7.3z"/><path fill="currentColor" d="M6.5 14.5a1 1 0 0 0 .7-.3a1 1 0 0 0 0-1.4a4 4 0 0 1 0-5.6a1 1 0 0 0-1.4-1.4a6 6 0 0 0 0 8.4a1 1 0 0 0 .7.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchlistLtr(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3h16v2H1Zm0 6h6v2H1Zm0 6h8v2H1Zm8-4.24h3.85L14.5 7l1.65 3.76H20l-3 3.17l.9 4.05l-3.4-2.14L11.1 18l.9-4.05Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchlistRtl(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H3v2h16zm0 6h-6v2h6zm0 6h-8v2h8zm-8-4.24H7.15L5.5 7l-1.65 3.76H0l3 3.17l-.9 4.05l3.4-2.14L8.9 18L8 13.95Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WikiText(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3v14h3v-2H3V5h1V3zm4 0v14h4v-2H7V5h2V3zm11 0v2h1v10h-1v2h3V3zm-5 0v2h2v10h-2v2h4V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Window(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm0 2h16v12H2z"/><path fill="currentColor" d="M4 6h12v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15a7 7 0 0 0 4.2-1.4l5.4 5.4l1.4-1.4l-5.4-5.4A7 7 0 1 0 8 15m0-2A5 5 0 1 1 8 3a5 5 0 0 1 0 10m1-6h2v2H9v2H7V9H5V7h2V5h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *OouiIcon {
	return &OouiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15a7 7 0 0 0 4.2-1.4l5.4 5.4l1.4-1.4l-5.4-5.4A7 7 0 1 0 8 15m0-2A5 5 0 1 1 8 3a5 5 0 0 1 0 10M5 7h6v2H5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
