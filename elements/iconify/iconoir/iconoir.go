package iconoir

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "7.2.0"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type IconoirIcon struct {
	*SVGSVGElement
}

func Accessibility(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10M7 9l5 1m5-1l-5 1m0 0v3m0 0l-2 5m2-5l2 5"/><path fill="currentColor" d="M12 7a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccessibilitySign(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11.5 12.5l7-.5l-1.5 6.5m-5.5-6l4.5-5L12.5 5L10 7.5m8.5-1a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/><path d="M5.5 12.5a5 5 0 0 1 7.584 6M3.729 15A5 5 0 0 0 11 20.831"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccessibilityTech(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m12.5 12.16l4-.16l-.5 4.5M11.833 12L13.5 9.538L10.833 8L9.5 9.846"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 7.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/><path stroke-linecap="round" stroke-linejoin="round" d="M10.5 18a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Activity(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h3l3-9l6 18l3-9h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeAfterEffects(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7v10a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4"/><path d="M14 13v-1a2 2 0 0 1 2-2v0a2 2 0 0 1 2 2v1zm0 0v1a2 2 0 0 0 2 2h1.5M6 16l1.125-3M12 16l-1.125-3m-3.75 0L9 8l1.875 5m-3.75 0h3.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeAfterEffectsSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.75 17A4.75 4.75 0 0 1 17 21.75H7A4.75 4.75 0 0 1 2.25 17V7A4.75 4.75 0 0 1 7 2.25h10A4.75 4.75 0 0 1 21.75 7zm-3.5-1a.75.75 0 0 1-.75.75H16A2.75 2.75 0 0 1 13.25 14v-2a2.75 2.75 0 1 1 5.5 0v1a.75.75 0 0 1-.75.75h-3.25V14c0 .69.56 1.25 1.25 1.25h1.5a.75.75 0 0 1 .75.75m-1-3.75V12a1.25 1.25 0 1 0-2.5 0v.25zM5.298 15.736a.75.75 0 1 0 1.404.527l.943-2.513h2.71l.943 2.513a.75.75 0 1 0 1.404-.527l-3-8a.75.75 0 0 0-1.404 0zm4.495-3.486H8.207L9 10.136z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeIllustrator(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7v10a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4m-5 5v4m0-7v.01"/><path d="m7 16l1.125-3M13 16l-1.125-3m-3.75 0L10 8l1.875 5m-3.75 0h3.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeIllustratorSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 21.75A4.75 4.75 0 0 0 21.75 17V7A4.75 4.75 0 0 0 17 2.25H7A4.75 4.75 0 0 0 2.25 7v10A4.75 4.75 0 0 0 7 21.75zM15.25 16a.75.75 0 0 0 1.5 0v-4a.75.75 0 0 0-1.5 0zM16 9.76a.75.75 0 0 1-.75-.75V9a.75.75 0 0 1 1.5 0v.01a.75.75 0 0 1-.75.75m-9.702 5.977a.75.75 0 1 0 1.404.526l.943-2.513h2.71l.943 2.513a.75.75 0 1 0 1.404-.527l-3-8a.75.75 0 0 0-1.404 0zM10 10.136l.793 2.114H9.207z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeIndesign(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7v10a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4M8.5 8v8"/><path d="M15.5 12v3.4a.6.6 0 0 1-.6.6h-1.4a2 2 0 0 1-2-2v0a2 2 0 0 1 2-2zm0 0V9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeIndesignSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.75 17A4.75 4.75 0 0 1 17 21.75H7A4.75 4.75 0 0 1 2.25 17V7A4.75 4.75 0 0 1 7 2.25h10A4.75 4.75 0 0 1 21.75 7zM8.5 16.75a.75.75 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v8a.75.75 0 0 1-.75.75m5-4h1.25v2.5H13.5a1.25 1.25 0 1 1 0-2.5m0 4a2.75 2.75 0 1 1 0-5.5h1.25V9a.75.75 0 0 1 1.5 0v6.4a1.35 1.35 0 0 1-1.35 1.35z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeLightroom(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7v10a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4"/><path d="M7 8v8h4m3-5.5V13m0 3v-3m0 0s0-2.5 3-2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeLightroomSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.75 17A4.75 4.75 0 0 1 17 21.75H7A4.75 4.75 0 0 1 2.25 17V7A4.75 4.75 0 0 1 7 2.25h10A4.75 4.75 0 0 1 21.75 7zM7 16.75a.75.75 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v7.25H11a.75.75 0 0 1 0 1.5zm6.25-.75a.75.75 0 0 0 1.5 0v-2.998l.003-.045a1.968 1.968 0 0 1 .265-.82c.235-.392.736-.887 1.982-.887a.75.75 0 0 0 0-1.5c-.973 0-1.713.232-2.268.586a.75.75 0 0 0-1.482.164z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobePhotoshop(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7v10a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4"/><path d="M7 16v-4m0 0V8h2a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2zm10-1c-.306-.613-.933-1-1.618-1H15a1.5 1.5 0 0 0-1.5 1.5v0A1.5 1.5 0 0 0 15 13h.5a1.5 1.5 0 0 1 1.5 1.5v0a1.5 1.5 0 0 1-1.5 1.5h-.382a1.809 1.809 0 0 1-1.618-1v0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobePhotoshopSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.75 17A4.75 4.75 0 0 1 17 21.75H7A4.75 4.75 0 0 1 2.25 17V7A4.75 4.75 0 0 1 7 2.25h10A4.75 4.75 0 0 1 21.75 7zM7 16.75a.75.75 0 0 1-.75-.75V8A.75.75 0 0 1 7 7.25h2a2.75 2.75 0 1 1 0 5.5H7.75V16a.75.75 0 0 1-.75.75m.75-5.5H9a1.25 1.25 0 1 0 0-2.5H7.75zm8.579.085a.75.75 0 1 0 1.342-.67a2.56 2.56 0 0 0-2.29-1.415H15a2.25 2.25 0 0 0 0 4.5h.5a.75.75 0 0 1 0 1.5h-.382a1.06 1.06 0 0 1-.947-.585a.75.75 0 0 0-1.342.67a2.559 2.559 0 0 0 2.289 1.415h.382a2.25 2.25 0 1 0 0-4.5H15a.75.75 0 0 1 0-1.5h.382c.4 0 .768.227.947.585" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeXd(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7v10a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4M7 8l4 8m-4 0l4-8"/><path d="M17 12v3.4a.6.6 0 0 1-.6.6H15a2 2 0 0 1-2-2v0a2 2 0 0 1 2-2zm0 0V9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeXdSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21.75 17A4.75 4.75 0 0 1 17 21.75H7A4.75 4.75 0 0 1 2.25 17V7A4.75 4.75 0 0 1 7 2.25h10A4.75 4.75 0 0 1 21.75 7zm-15.085-.33a.75.75 0 0 1-.336-1.006L8.162 12L6.329 8.335a.75.75 0 0 1 1.342-.67l1.33 2.658l1.328-2.659a.75.75 0 0 1 1.342.671L9.839 12l1.832 3.664a.75.75 0 0 1-1.342.671L9 13.677l-1.329 2.658a.75.75 0 0 1-1.006.336M15 12.75h1.25v2.5H15a1.25 1.25 0 1 1 0-2.5m0 4a2.75 2.75 0 1 1 0-5.5h1.25V9a.75.75 0 0 1 1.5 0v6.4a1.35 1.35 0 0 1-1.35 1.35z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AfricanTree(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22V12m0-4v4m0 0l3-3m-2.576 9.576l6.169-6.169a5.502 5.502 0 0 0-.513-8.234a9.904 9.904 0 0 0-12.16 0a5.502 5.502 0 0 0-.513 8.234l6.169 6.169a.6.6 0 0 0 .848 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Agile(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 19H22m0 0l-2.5-2.5M22 19l-2.5 2.5M12 2L9.5 4.5L12 7"/><path d="M10.5 4.5a7.5 7.5 0 0 1 0 15H2"/><path d="M6.756 5.5A7.497 7.497 0 0 0 3 12c0 1.688.558 3.246 1.5 4.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirConditioner(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 3.6V11H2V3.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6M18 7h1M2 11l.79 2.584A2 2 0 0 0 4.702 15H6m16-4l-.79 2.584A2 2 0 0 1 19.298 15H18m-8.5-.5s0 7-3.5 7m8.5-7s0 7 3.5 7m-6-7v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplane(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 4.5v4.667a.6.6 0 0 1-.282.51l-7.436 4.647a.6.6 0 0 0-.282.508v.9a.6.6 0 0 0 .746.582l6.508-1.628a.6.6 0 0 1 .746.582v2.96a.6.6 0 0 1-.205.451l-2.16 1.89c-.458.402-.097 1.151.502 1.042l3.256-.591a.6.6 0 0 1 .214 0l3.256.591c.599.11.96-.64.502-1.041l-2.16-1.89a.6.6 0 0 1-.205-.452v-2.96a.6.6 0 0 1 .745-.582l6.51 1.628a.6.6 0 0 0 .745-.582v-.9a.6.6 0 0 0-.282-.508l-7.436-4.648a.6.6 0 0 1-.282-.509V4.5a1.5 1.5 0 0 0-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneHelix(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path d="M12 9s-1.988-1.975-2-4c.001-1.993-.05-4.001 2-4c1.948.001 1.997 1.976 2 4c.003 1.985-2 4-2 4m3 3s1.975-1.988 4-2c1.993.001 4.001-.05 4 2c-.001 1.948-1.976 1.997-4 2c-1.985.003-4-2-4-2m-6 0s-1.975 1.988-4 2c-1.993-.001-4.001.05-4-2c.001-1.948 1.976-1.997 4-2c1.985-.003 4 2 4 2m3 3s1.988 1.975 2 4c-.001 1.993.05 4.001-2 4c-1.948-.001-1.997-1.976-2-4c-.003-1.985 2-4 2-4" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneHelixFortyFiveDeg(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M14.12 14.121A3 3 0 1 0 9.879 9.88a3 3 0 0 0 4.243 4.242"/><path d="M9.879 9.879s-2.803.009-4.243-1.415c-1.409-1.41-2.864-2.793-1.414-4.242c1.378-1.377 2.81-.015 4.242 1.414C9.87 7.037 9.88 9.879 9.88 9.879m4.241 0s-.009-2.803 1.415-4.243c1.41-1.409 2.793-2.864 4.242-1.414c1.377 1.378.015 2.81-1.414 4.242c-1.402 1.406-4.243 1.415-4.243 1.415m-4.242 4.242s.009 2.803-1.415 4.243c-1.41 1.409-2.793 2.864-4.242 1.414c-1.377-1.378-.015-2.81 1.414-4.242c1.401-1.406 4.243-1.415 4.243-1.415m4.242 0s2.803-.009 4.243 1.415c1.409 1.41 2.864 2.793 1.414 4.242c-1.378 1.377-2.81.015-4.242-1.414c-1.406-1.402-1.415-4.243-1.415-4.243" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.881 9.887l-7.099 4.437a.6.6 0 0 0-.282.508v.9a.6.6 0 0 0 .746.582l6.508-1.628a.6.6 0 0 1 .746.582v2.96a.6.6 0 0 1-.205.451l-2.16 1.89c-.458.402-.097 1.151.502 1.042l3.256-.591a.6.6 0 0 1 .214 0l3.256.591c.599.11.96-.64.502-1.041l-2.16-1.89a.6.6 0 0 1-.205-.452v-2.96a.6.6 0 0 1 .745-.582l.458.115M10.5 7.5v-3A1.5 1.5 0 0 1 12 3v0a1.5 1.5 0 0 1 1.5 1.5v4.667a.6.6 0 0 0 .282.51l7.436 4.647a.6.6 0 0 1 .282.508v.9a.6.6 0 0 1-.745.582l-2.006-.502M3 3l18 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneRotation(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.879 14.122a3 3 0 1 0 4.242-4.243a3 3 0 0 0-4.242 4.243M4.37 16.773A8.956 8.956 0 0 1 3.002 12c0-4.236 2.934-7.792 6.878-8.747A8.998 8.998 0 0 1 12 3.002m7.715 4.365A8.953 8.953 0 0 1 20.999 12c0 3.806-2.368 7.063-5.709 8.378c-1.02.4-2.13.621-3.29.621"/><path d="M14.121 9.88s-.009-2.803 1.415-4.243c1.41-1.409 2.793-2.865 4.242-1.415c1.377 1.378.015 2.81-1.414 4.243c-1.402 1.406-4.243 1.414-4.243 1.414M9.879 14.12s.009 2.803-1.415 4.243c-1.41 1.409-2.793 2.865-4.242 1.415c-1.377-1.378-.015-2.81 1.414-4.243c1.402-1.406 4.243-1.414 4.243-1.414" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplay(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 17H3V4h18v13h-3"/><path d="M8.622 19.067L11.5 14.75a.6.6 0 0 1 .998 0l2.88 4.318a.6.6 0 0 1-.5.933H9.12a.6.6 0 0 1-.5-.933Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaySolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 17H3V4h18v13h-3"/><path fill="currentColor" d="M8.622 19.067L11.5 14.75a.6.6 0 0 1 .998 0l2.88 4.318a.6.6 0 0 1-.5.933H9.12a.6.6 0 0 1-.5-.933Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 13h-5V8M5 3.5L7 2m12 1.5L17 2"/><path d="M12 22a9 9 0 1 0 0-18a9 9 0 0 0 0 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75s9.75-4.365 9.75-9.75S17.385 3.25 12 3.25m0 10.5a.75.75 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v4.25H17a.75.75 0 0 1 0 1.5zm-7.6-9.8a.75.75 0 0 0 1.05.15l2-1.5a.75.75 0 1 0-.9-1.2l-2 1.5a.75.75 0 0 0-.15 1.05m15.2 0a.75.75 0 0 1-1.05.15l-2-1.5a.75.75 0 1 1 .9-1.2l2 1.5a.75.75 0 0 1 .15 1.05" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Album(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" d="M12 15.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0V7.6a.6.6 0 0 1 .6-.6H15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlbumCarousel(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 19.4V4.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" d="M22 6v12m-11-3.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0V8.6a.6.6 0 0 1 .6-.6H13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlbumList(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 17.4V2.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" d="M8 22h13.4a.6.6 0 0 0 .6-.6V8m-11 4.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0V6.6a.6.6 0 0 1 .6-.6H13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlbumOpen(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M15 2.2c4.564.927 8 4.962 8 9.8c0 4.838-3.436 8.873-8 9.8"/><path stroke-linejoin="round" d="M15 9c1.141.284 2 1.519 2 3s-.859 2.716-2 3M1 2h10v20H1"/><path d="M4 15.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0V7.6a.6.6 0 0 1 .6-.6H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignBottomBox(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 8l.01.011M4 4l.01.011M8 4l.01.011M12 4l.01.011M16 4l.01.011M20 4l.01.011M20 8l.01.011M4 12v8h16v-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignBottomBoxSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4 8l.01.01M4 4l.01.01M8 4l.01.01M12 4l.01.01M16 4l.01.01M20 4l.01.01M20 8l.01.01"/><path fill="currentColor" d="M4 12v8h16v-8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h18M3 14h18M6 10h12M6 18h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontalCenters(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 22V2"/><path d="M19 16H5a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontalCentersSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 22V2"/><path fill="currentColor" d="M19 16H5a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontalSpacing(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3 22V2m18 20V2"/><path d="M15 16H9a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontalSpacingSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3 22V2m18 20V2"/><path fill="currentColor" d="M15 16H9a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h18M3 10h18M3 14h18M3 18h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 10h14M3 6h18M3 18h14M3 14h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeftBox(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.004 3.995l-.011.01m4.011-.01l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m-3.989-.01l-.011.01m-3.987-16.01h-8v16h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeftBoxSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m16.005 3.995l-.011.01m4.011-.01l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m.011 3.99l-.011.01m-3.989-.01l-.011.01"/><path fill="currentColor" d="M12.005 3.995h-8v16h8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 10h14M3 6h18M7 18h14M3 14h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRightBox(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.006 20.005l.01-.01m-4.01.01l.01-.01m-.01-3.99l.01-.01m-.01-3.99l.01-.01m-.01-3.99l.01-.01m-.01-3.99l.01-.01m3.99.01l.01-.01m3.99 16.01h8v-16h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRightBoxSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.005 20.005l.011-.01m-4.011.01l.011-.01m-.011-3.99l.011-.01m-.011-3.99l.011-.01m-.011-3.99l.011-.01m-.011-3.99l.011-.01m3.989.01l.011-.01"/><path fill="currentColor" d="M12.005 20.005h8v-16h-8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTopBox(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 16l.01-.011M4 20l.01-.011M8 20l.01-.011M12 20l.01-.011M16 20l.01-.011M20 20l.01-.011M20 16l.01-.011M4 12V4h16v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTopBoxSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4 16l.01-.01M4 20l.01-.01M8 20l.01-.01M12 20l.01-.01M16 20l.01-.01M20 20l.01-.01M20 16l.01-.01"/><path fill="currentColor" d="M4 12V4h16v8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVerticalCenters(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 12H2"/><path d="M8 19V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVerticalCentersSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 12H2"/><path fill="currentColor" d="M8 19V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVerticalSpacing(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 3H2m20 18H2"/><path d="M8 15V9a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVerticalSpacingSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 3H2m20 18H2"/><path fill="currentColor" d="M8 15V9a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleTool(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M3 21V3h6v12h12v6z"/><path d="M13 19v2m-4-2v2M3 7h2m-2 4h2m-2 4h2m12 4v2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Antenna(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M16 1s1.5 1 1.5 3S16 7 16 7M8 1S6.5 2 6.5 4S8 7 8 7M7 23l1.111-4M17 23l-1.111-4M14.5 14L12 5l-2.5 9m5 0h-5m5 0l1.389 5M9.5 14l-1.389 5m0 0h7.778"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AntennaOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="m7 23l1.111-4M17 23l-1.111-4M9.5 14l-1.389 5M9.5 14h4m-4 0l.8-2.88M8.11 19h7.778m0 0l-1.184-4.264M11.445 7L12 5l1.047 3.768M3 3l18 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AntennaSignal(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 8S19 9.5 19 12s-1.5 4-1.5 4m3-11S23 7.5 23 12s-2.5 7-2.5 7M6.5 8S5 9.5 5 12s1.5 4 1.5 4m-3-11S1 7.5 1 12s2.5 7 2.5 7"/><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AntennaSignalTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 9s1 1.125 1 3s-1 3-1 3m-3-2.99l.01-.011M17 7s2 1.786 2 5s-2 5-2 5M9 9s-1 1.125-1 3s1 3 1 3M7 7s-2 1.786-2 5s2 5 2 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppNotification(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6m2 4v3a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6V9a6 6 0 0 1 6-6h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppStore(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10M10.5 5.5l7 11m-4-11l-7 11m7-2.5h-7m11 0H16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppStoreSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12m8.847-7.133a.75.75 0 0 1 1.036.23L12 6.46l.867-1.363a.75.75 0 1 1 1.266.806l-1.244 1.954l3.432 5.393H17.5a.75.75 0 0 1 0 1.5h-.225l.858 1.347a.75.75 0 1 1-1.266.806L12 9.254L9.457 13.25H13.5a.75.75 0 0 1 0 1.5H8.503l-1.37 2.153a.75.75 0 0 1-1.266-.806l.858-1.347H6.5a.75.75 0 0 1 0-1.5h1.18l3.431-5.393l-1.244-1.954a.75.75 0 0 1 .23-1.036" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 19V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m12.147 21.265l-.147-.03l-.147.03c-2.377.475-4.62.21-6.26-1.1C3.964 18.86 2.75 16.373 2.75 12c0-4.473 1.008-6.29 2.335-6.954c.695-.347 1.593-.448 2.735-.317c1.141.132 2.458.488 3.943.983l.26.086l.255-.102c2.482-.992 4.713-1.373 6.28-.641c1.47.685 2.692 2.538 2.692 6.945c0 4.374-1.213 6.86-2.843 8.164c-1.64 1.312-3.883 1.576-6.26 1.1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 5.5C12 3 11 2 9 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleHalf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m12.147 21.265l-.147-.03l-.147.03c-2.377.475-4.62.21-6.26-1.1C3.964 18.86 2.75 16.373 2.75 12c0-4.473 1.008-6.29 2.335-6.954c.695-.347 1.593-.448 2.735-.317c1.141.132 2.458.488 3.943.983l.26.086l.255-.102c2.482-.992 4.713-1.373 6.28-.641c1.47.685 2.692 2.538 2.692 6.945c0 4.374-1.213 6.86-2.843 8.164c-1.64 1.312-3.883 1.576-6.26 1.1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 5.5C12 3 11 2 9 2"/><path d="M12 6v15"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 12v2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleHalfAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m12.147 21.265l-.147-.03l-.147.03c-2.377.475-4.62.21-6.26-1.1C3.964 18.86 2.75 16.373 2.75 12c0-4.473 1.008-6.29 2.335-6.954c.695-.347 1.593-.448 2.735-.317c1.141.132 2.458.488 3.943.983l.26.086l.255-.102c2.482-.992 4.713-1.373 6.28-.641c1.47.685 2.692 2.538 2.692 6.945c0 4.374-1.213 6.86-2.843 8.164c-1.64 1.312-3.883 1.576-6.26 1.1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 5.5C12 3 11 2 9 2"/><path d="M12 6v15"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 12v2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleImacTwoThousandTwentyOne(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 15.5V2.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.9m-20 0v1.9a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6v-1.9m-20 0h20M9 22h1.5m0 0v-4m0 4h3m0 0H15m-1.5 0v-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleImacTwoThousandTwentyOneSide(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 22h2m6 0H8m0 0l2-8.5m0 0L7 2m3 11.5l1.5 5.5m5.5 3h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleMac(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 2c.363 2.18-1.912 3.83-3.184 4.571c-.375.219-.799-.06-.734-.489C12.299 4.64 13.094 2 16 2Z"/><path d="M9 6.5c.897 0 1.69.2 2.294.42a3.58 3.58 0 0 0 2.412 0A6.73 6.73 0 0 1 16 6.5c1.085 0 2.465.589 3.5 1.767C16 11 17 15.5 20.269 16.692c-1.044 2.867-3.028 4.808-4.77 4.808c-1.5 0-1.499-.7-2.999-.7s-1.5.7-3 .7c-2.5 0-5.5-4-5.5-9c0-4 3-6 5-6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleShortcuts(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m9.853 14.633l-6.201-3.946a2 2 0 0 1 0-3.374l6.2-3.946a4 4 0 0 1 4.296 0l6.2 3.946a2 2 0 0 1 0 3.374l-6.2 3.946a4 4 0 0 1-4.296 0Z"/><path d="m18.286 12l2.063 1.313a2 2 0 0 1 0 3.374l-6.201 3.946a4 4 0 0 1-4.296 0l-6.2-3.946a2 2 0 0 1 0-3.374L5.714 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleShortcutsSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3.249 11.32c-1.698-1.08-1.698-3.56 0-4.64l6.2-3.946a4.75 4.75 0 0 1 5.101 0l6.201 3.946c1.698 1.08 1.698 3.56 0 4.64l-6.2 3.946a4.75 4.75 0 0 1-5.101 0z"/><path d="m19.66 11.986l-5.11-3.252a4.75 4.75 0 0 0-5.1 0l-5.11 3.252l5.915 3.765a3.25 3.25 0 0 0 3.49 0zm1.362.889c-.085.07-.175.134-.27.195l-6.202 3.946a4.75 4.75 0 0 1-5.1 0L3.249 13.07a2.85 2.85 0 0 1-.27-.195c-1.423 1.16-1.333 3.425.27 4.445l6.2 3.946a4.75 4.75 0 0 0 5.101 0l6.201-3.946c1.603-1.02 1.693-3.285.27-4.445"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleSwift(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.457 14.59c.446-1.437 1.451-6.75-5.93-11.49a.636.636 0 0 0-.808.1a.593.593 0 0 0-.022.79c.03.036 2.75 3.35 1.783 7.135c-1.673-1.151-8.324-6.423-8.324-6.423L11 11L3.862 6.4s5.046 6.195 8.134 8.526c-1.495.536-4.743 1.104-9.033-1.562a.637.637 0 0 0-.771.074a.593.593 0 0 0-.106.743C2.229 14.42 5.668 20 12.939 20c1.995 0 3.16-.568 4.098-1.024c.576-.279 1.031-.501 1.528-.501c1.236 0 2.047 1.227 2.054 1.238a.632.632 0 0 0 .583.285a.62.62 0 0 0 .526-.37c.893-2.074-.645-4.269-1.271-5.039"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleWallet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2Z"/><path d="M3 15h6.4c.331 0 .605.278.75.576c.206.423.694.924 1.85.924c1.156 0 1.644-.5 1.85-.924c.145-.298.419-.576.75-.576H21M3 7h18M3 11h18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 15.5v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H13v4.2m2.857 0H13m2.857 0L18 15.5m-7 0L9.929 13M5 15.5L6.071 13m0 0L8 8.5L9.929 13M6.07 13h3.86"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArcThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 16c0-5.523-4.477-10-10-10S2 10.477 2 16"/><path fill="currentColor" d="M2 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2m20 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArcThreeDcenterPoint(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-dasharray="3 3" d="M22 16c0-5.523-4.477-10-10-10c-4.1 0-7.625 2.468-9.168 6"/><path fill="currentColor" d="M2 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M2 16h10"/><path fill="currentColor" d="M12 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arcade(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11 8.5L9.8 9l-7.448 3.386a.6.6 0 0 0-.352.546v.136a.6.6 0 0 0 .352.546l8.82 4.01a2 2 0 0 0 1.656 0l8.82-4.01a.6.6 0 0 0 .352-.546v-.136a.6.6 0 0 0-.352-.546L14.2 9L13 8.5"/><path d="M22 13v4.112a.6.6 0 0 1-.354.547l-8.825 3.972a2 2 0 0 1-1.642 0l-8.825-3.972A.6.6 0 0 1 2 17.112V13"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/><path d="M11 8v5a1 1 0 1 0 2 0V8"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 13h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archery(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h9m-9 0l-2-2H2l2 2l-2 2h4zm9 0l-2-2m2 2l-2 2m1 8.5c2.761 0 5-4.701 5-10.5S18.761 1.5 16 1.5S11 6.201 11 12s2.239 10.5 5 10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArcheryMatch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.611 15.89l12.02-12.022M8.612 15.89H5.783l-2.829 2.829h2.829v2.828l2.828-2.828zm12.02-12.02h-2.828m2.829 0v2.828M15.39 15.89L3.367 3.867M15.39 15.89h2.829l2.828 2.829h-2.828v2.828l-2.829-2.828zM3.37 3.87h2.828m-2.829 0v2.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 6h10M7 9h10m-8 8h6"/><path d="M3 12h-.4a.6.6 0 0 0-.6.6v8.8a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6v-8.8a.6.6 0 0 0-.6-.6H21M3 12V2.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V12M3 12h18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaSearch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.124 20.119a3 3 0 1 0-4.248-4.237a3 3 0 0 0 4.248 4.237m0 0L22 22M7 2H4v3m0 6v2m7-11h2m-2 20h2m7-11v2M17 2h3v3M7 22H4v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowArchery(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.611 15.89l12.02-12.022M8.612 15.89H5.783l-2.829 2.829h2.829v2.828l2.828-2.828zm12.02-12.02h-2.828m2.829 0v2.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v18m0 0l8.5-8.5M12 21l-8.5-8.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v8m0 0l3.5-3.5M12 16l-3.5-3.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m4.03 11.78l-3.5 3.5a.75.75 0 0 1-1.06 0l-3.5-3.5a.75.75 0 1 1 1.06-1.06l2.22 2.22V8a.75.75 0 0 1 1.5 0v6.19l2.22-2.22a.75.75 0 1 1 1.06 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 6L6 19m0 0V6.52M6 19h12.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.828 9.172l-5.656 5.656m0 0h4.95m-4.95 0v-4.95M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m3.359 8.452a.75.75 0 0 0-1.06-1.061l-4.377 4.377v-3.14a.75.75 0 1 0-1.5 0v4.95c0 .414.335.75.75.75h4.95a.75.75 0 0 0 0-1.5h-3.14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.828 9.172l-5.656 5.656m0 0h4.95m-4.95 0v-4.95M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 6l13 13m0 0V6.52M19 19H6.52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.171 9.172l5.657 5.656m0 0h-4.95m4.95 0v-4.95M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25M9.702 8.641a.75.75 0 1 0-1.06 1.06l4.376 4.377h-3.14a.75.75 0 0 0 0 1.5h4.95a.747.747 0 0 0 .75-.75v-4.95a.75.75 0 0 0-1.5 0v3.14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.171 9.172l5.657 5.656m0 0h-4.95m4.95 0v-4.95M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm6.102 6.391a.75.75 0 1 0-1.06 1.06l4.376 4.377h-3.14a.75.75 0 0 0 0 1.5h4.95a.747.747 0 0 0 .75-.75v-4.95a.75.75 0 0 0-1.5 0v3.14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 5h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V9a4 4 0 0 1 4-4"/><path d="m14.5 10.75l-2.5 2.5l-2.5-2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowEmailForward(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 10H8c-8 0-8 11 0 11m14-11l-7-7m7 7l-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowEnlargeTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.5 9.5L6 12l2.5 2.5m7-5L18 12l-2.5 2.5"/><path d="M2 15V9a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12H3m0 0l8.5-8.5M3 12l8.5 8.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 12H8m0 0l3.5 3.5M8 12l3.5-3.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m-1.03 14.78l-3.5-3.5a.75.75 0 0 1 0-1.06l3.5-3.5a.75.75 0 1 1 1.06 1.06l-2.22 2.22H16a.75.75 0 0 1 0 1.5H9.81l2.22 2.22a.75.75 0 1 1-1.06 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.75 12h-10m0 0l2.75 2.75M6.75 12L9.5 9.25"/><path d="M2 15V9a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowReduceTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 9.5L9.5 12L7 14.5m9.5-5L14 12l2.5 2.5"/><path d="M6 5h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V9a4 4 0 0 1 4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h18m0 0l-8.5-8.5M21 12l-8.5 8.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h8m0 0l-3.5-3.5M16 12l-3.5 3.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m1.03 6.72l3.5 3.5a.75.75 0 0 1 0 1.06l-3.5 3.5a.75.75 0 1 1-1.06-1.06l2.22-2.22H8a.75.75 0 0 1 0-1.5h6.19l-2.22-2.22a.75.75 0 0 1 1.06-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6.75 12h10m0 0L14 14.75M16.75 12L14 9.25"/><path d="M2 15V9a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSeparate(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 7l-5 5l5 5m8-10l5 5l-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSeparateVertical(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 8l-5-5l-5 5m10 8l-5 5l-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUnion(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 7l-5 5l5 5M4 7l5 5l-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUnionVertical(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 4l-5 5l-5-5m10 16l-5-5l-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21V3m0 0l8.5 8.5M12 3l-8.5 8.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16V8m0 0l3.5 3.5M12 8l-3.5 3.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m4.03 9.72l-3.5-3.5a.75.75 0 0 0-1.06 0l-3.5 3.5a.75.75 0 1 0 1.06 1.06l2.22-2.22V16a.75.75 0 0 0 1.5 0V9.81l2.22 2.22a.75.75 0 1 0 1.06-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 19L6 6m0 0v12.48M6 6h12.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.828 14.828L9.172 9.172m0 0h4.95m-4.95 0v4.95M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m2.121 7.171h-4.95a.748.748 0 0 0-.75.75v4.95a.75.75 0 0 0 1.5 0v-3.139l4.377 4.377a.75.75 0 1 0 1.06-1.061L10.983 9.92h3.14a.75.75 0 0 0 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.828 14.828L9.172 9.172m0 0h4.95m-4.95 0v4.95M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm10.521 6.171h-4.95a.748.748 0 0 0-.75.75v4.95a.75.75 0 0 0 1.5 0v-3.139l4.377 4.377a.75.75 0 1 0 1.06-1.061L10.983 9.92h3.14a.75.75 0 0 0 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 19L19 6m0 0v12.48M19 6H6.52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.171 14.828l5.657-5.656m0 0h-4.95m4.95 0v4.95M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25M9.879 8.421h4.95a.747.747 0 0 1 .75.75v4.95a.75.75 0 0 1-1.5 0v-3.139l-4.377 4.377a.75.75 0 1 1-1.06-1.061l4.376-4.377h-3.14a.75.75 0 0 1 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.171 14.828l5.657-5.656m0 0h-4.95m4.95 0v4.95M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm6.279 6.171h4.95a.747.747 0 0 1 .75.75v4.95a.75.75 0 0 1-1.5 0v-3.139l-4.377 4.377a.75.75 0 1 1-1.06-1.061l4.376-4.377h-3.14a.75.75 0 0 1 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.5 13.25l-2.5-2.5l-2.5 2.5"/><path d="M6 5h12a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V9a4 4 0 0 1 4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsUpFromLine(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20h18M6 17V4m0 0L2 8m4-4l4 4m8 9V4m0 0l-4 4m4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asana(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 11.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8m-5 9a4 4 0 1 0 0-8a4 4 0 0 0 0 8m10 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22V2m8.572 15.55L3.428 7.249m0 10.301L20.572 7.249"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSign(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 19.125A9 9 0 1 1 21 12c0 5.5-6 5.5-6 2V8"/><path d="M15 12v-1.5C15 9.12 13.657 8 12 8s-3 1.12-3 2.5V12m6 0v1.5c0 1.38-1.343 2.5-3 2.5s-3-1.12-3-2.5V12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSignCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.278 17.541A7 7 0 1 1 19 12c0 4.278-5 3.722-5 1V8.5"/><path d="M14 12v-.5a2.5 2.5 0 0 0-5 0v.5m5 0v.5a2.5 2.5 0 0 1-5 0V12"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Atom(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.404 13.61C3.515 13.145 3 12.592 3 12c0-1.657 4.03-3 9-3s9 1.343 9 3c0 .714-.75 1.37-2 1.886m-7-2.876l.01-.011"/><path d="M16.883 6c-.005-1.023-.263-1.747-.797-2.02c-1.477-.751-4.503 2.23-6.76 6.658c-2.256 4.429-2.889 8.629-1.412 9.381c.527.269 1.252.061 2.07-.519"/><path d="M9.6 4.252c-.66-.386-1.243-.497-1.686-.271c-1.477.752-.844 4.952 1.413 9.38c2.256 4.43 5.282 7.41 6.759 6.658c1.312-.669.958-4.061-.722-7.917"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21.438 11.662l-9.19 9.19a6.003 6.003 0 1 1-8.49-8.49l9.19-9.19a4.002 4.002 0 0 1 5.66 5.66l-9.2 9.19a2.001 2.001 0 1 1-2.83-2.83l8.49-8.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AugmentedReality(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m5.5 15.5l.614-1.718M10.5 15.5l-.614-1.718m-3.772 0L8 8.5l1.886 5.282m-3.772 0h3.772M13 15.5v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H13v4.2m2.857 0H13m2.857 0L18 15.5"/><path d="M2 18.4V5.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AutoFlash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16 9.5l.692-1.5M22 9.5L21.308 8m0 0L19 3l-2.308 5m4.616 0h-4.616M13 10h-3V3L2 14h6v7l6-8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AviFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m12 9l1.5 6L15 9m3 6V9"/><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M6 15v-3m0 0v-1.5A1.5 1.5 0 0 1 7.5 9v0A1.5 1.5 0 0 1 9 10.5V12m-3 0h3m0 0v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Axes(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 19.452l-9-6.61m0 0V3m0 9.843l-9 6.609m17.438-2.742L21 19.452L18.188 20M9.75 5.194L12 3l2.25 2.194M5.813 20L3 19.452l.563-2.742"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackwardFifteenSeconds(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 13a9 9 0 1 0 9-9M9 9v7"/><path d="M15 9h-2a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h1a1 1 0 0 1 1 1V15a1 1 0 0 1-1 1h-2m0-12H4.5m0 0l2-2m-2 2l2 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BadgeCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M10.521 2.624a2 2 0 0 1 2.958 0l1.02 1.12a2 2 0 0 0 1.572.651l1.513-.07a2 2 0 0 1 2.092 2.09l-.071 1.514a2 2 0 0 0 .651 1.572l1.12 1.02a2 2 0 0 1 0 2.958l-1.12 1.02a2 2 0 0 0-.651 1.572l.07 1.513a2 2 0 0 1-2.09 2.092l-1.514-.071a2 2 0 0 0-1.572.651l-1.02 1.12a2 2 0 0 1-2.958 0l-1.02-1.12a2 2 0 0 0-1.572-.651l-1.513.07a2 2 0 0 1-2.092-2.09l.071-1.514a2 2 0 0 0-.651-1.572l-1.12-1.02a2 2 0 0 1 0-2.958l1.12-1.02a2 2 0 0 0 .651-1.572l-.07-1.513a2 2 0 0 1 2.09-2.092l1.514.071a2 2 0 0 0 1.572-.651z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 12l2 2l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M4.508 20h14.984a.6.6 0 0 0 .592-.501l1.8-10.8A.6.6 0 0 0 21.292 8H2.708a.6.6 0 0 0-.592.699l1.8 10.8a.6.6 0 0 0 .592.501Z"/><path d="M7 8V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Balcony(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 13v8m4-8v8m8-8v8m-4-8v8m8-8v8M2 21h20M2 13h20m-4-3V3.6a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 0-.6.6V10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9.5L12 4l9 5.5M5 20h14M10 9h4m-8 8v-5m4 5v-5m4 5v-5m4 5v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 19V5h1m6 14V5h1M9 5v14m7-14v14m3-14v14M6 5v14H5m8-14v14h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basketball(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.736 20.192c4.524-3.168 5.623-9.404 2.455-13.928C17.024 1.74 10.788.641 6.264 3.81C1.74 6.976.641 13.212 3.808 17.736c3.168 4.524 9.404 5.623 13.928 2.456m0 0L6.264 3.809"/><path d="M19.577 5.473c-3.77 5.896-8.508 9.214-16.302 11.415"/><path d="M13.06 2.056c.413 5.24 3.392 9.494 8.646 12.35M2.293 9.595c4.783 2.18 7.761 6.434 8.647 12.349"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketballField(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 5h9.4a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H12m0-14H2.6a.6.6 0 0 0-.6.6v12.8a.6.6 0 0 0 .6.6H12m0-14v14"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6M2 17A5 5 0 0 0 2 7m20 10a5 5 0 0 1 0-10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bathroom(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 13v3a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4v-2.4a.6.6 0 0 1 .6-.6zm-5 7l1 2m-9-2l-1 2m14-9V7a4 4 0 0 0-4-4h-5"/><path d="M15.4 8H8.6c-.331 0-.596-.268-.56-.598C8.186 6.075 8.863 3 12 3s3.814 3.075 3.96 4.402c.036.33-.229.598-.56.598"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BathroomSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M21 13v3a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4v-2.4a.6.6 0 0 1 .6-.6z"/><path d="m16 20l1 2m-9-2l-1 2m14-9V7a4 4 0 0 0-4-4h-5"/><path fill="currentColor" d="M15.4 8H8.6c-.331 0-.596-.268-.56-.598C8.186 6.075 8.863 3 12 3s3.814 3.075 3.96 4.402c.036.33-.229.598-.56.598"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharging(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10.167 9L8.5 12h4l-1.667 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFifty(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path d="M4 14.4V9.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path d="M4 14.4V9.6a.6.6 0 0 1 .6-.6h11.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryIndicator(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 13h4M6 13h2m2 0H8m0 0v-2m0 2v2M6 7H2.6a.6.6 0 0 0-.6.6v10.8a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H18M6 7V5h2v2M6 7h2m0 0h8m0 0V5h2v2m-2 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatterySeventyFive(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path d="M4 14.4V9.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatterySlash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 3l18 18m2-11v4M5.5 6H3a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h14.5m2.5-3.5V8a2 2 0 0 0-2-2h-6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryTwentyFive(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path d="M4 14.4V9.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M23 10v4"/><path d="M1 16V8a2 2 0 0 1 2-2h15a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10.5 9v2m0 4.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bbq(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.5 14.5L5 22M8 6s1-1.061 1-2c0-1.333-1-2-1-2m4 4s1-1.061 1-2c0-1.333-1-2-1-2m4 4s1-1.061 1-2c0-1.333-1-2-1-2"/><path stroke-linejoin="round" d="M16.5 17.5h-9"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.5 14.5l2.1 4.5m.9 3a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/><path d="M12 15a7 7 0 0 0 6.975-6.4a.563.563 0 0 0-.575-.6H5.6a.563.563 0 0 0-.575.6A7 7 0 0 0 12 15Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeachBag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m2.77 13l-.633-3.287A.6.6 0 0 1 2.727 9h18.547a.6.6 0 0 1 .589.713L21.23 13M2.769 13h18.462M2.769 13l.616 4m17.846-4l-.616 4m0 0l-.537 3.491a.6.6 0 0 1-.593.509H4.515a.6.6 0 0 1-.593-.509L3.385 17m17.23 0H3.385"/><path d="M8 9V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeachBagBig(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="m2.77 12l-.633-3.287A.6.6 0 0 1 2.727 8h18.547a.6.6 0 0 1 .589.713L21.23 12M2.769 12h18.462M2.769 12l.616 4m17.846-4l-.616 4m0 0l-.537 3.491a.6.6 0 0 1-.593.509H4.515a.6.6 0 0 1-.593-.509L3.385 16m17.23 0H3.385M5 8V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 4v16a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2"/><path d="M3 8h8V6m10 2h-8V6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BedReady(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 16V4a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h9"/><path d="M3 8h8V6m10 2h-8V6m3 14l2 2l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.197 11.217c5.07 0 5.07 6.783 0 6.783H2v-6.783m6.197 0H2m6.197 0c5.07 0 5.07-6.217 0-6.217H2v6.217M18 9c-2.21 0-4 2.015-4 4.5h8c0-2.485-1.79-4.5-4-4.5m-4 4.5c0 2.485 1.79 4.5 4 4.5c2.755 0 3.5-2 3.5-2m-1-10h-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5"/><path d="M9.099 11.826c2.535 0 2.535 4.174 0 4.174H6v-4.174m3.099 0H6m3.099 0c2.535 0 2.535-3.826 0-3.826H6v3.826M15.5 11a2.5 2.5 0 0 0-2.5 2.5h5a2.5 2.5 0 0 0-2.5-2.5M13 13.5a2.5 2.5 0 0 0 2.5 2.5c.928 0 1.49-.322 1.813-.62M17 8.5h-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 8.4c0-1.697-.632-3.325-1.757-4.525C15.117 2.675 13.59 2 12 2c-1.591 0-3.117.674-4.243 1.875C6.632 5.075 6 6.703 6 8.4C6 15.867 3 18 3 18h18s-3-2.133-3-9.6M13.73 21a1.999 1.999 0 0 1-3.46 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellNotification(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.134 11C18.715 16.375 21 18 21 18H3s3-2.133 3-9.6c0-1.697.632-3.325 1.757-4.525C8.883 2.675 10.41 2 12 2c.337 0 .672.03 1 .09M19 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-5.27 13a1.999 1.999 0 0 1-3.46 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.27 6.5C6.093 7.11 6 7.75 6 8.4C6 15.867 3 18 3 18h15M7.757 3.875C8.883 2.675 10.41 2 12 2c1.591 0 3.117.674 4.243 1.875C17.368 5.075 18 6.703 18 8.4c0 7.467 3 9.6 3 9.6m-7.27 3a1.999 1.999 0 0 1-3.46 0M3 3l18 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bicycle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8M8.5 7.5h6M19 15l-4-7.5h-.5m0 0l2-3m0 0H14m2.5 0h2"/><path d="m5 15l3.5-7.5L12 14h3M8.5 7.5c-.333-1-1.5-3-3.5-3"/><path d="M19 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bin(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinFull(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m19.262 17.038l1.676-12.575a.6.6 0 0 0-.372-.636L16 2h-5.5l-.682 1.5L5 2L3.21 3.79a.6.6 0 0 0-.17.504l1.698 12.744a4 4 0 0 0 1.98 2.944l.32.183a10 10 0 0 0 9.923 0l.32-.183a4 4 0 0 0 1.98-2.944ZM16 2l-2 5m-5-.5l.818-3"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinHalf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0M11 18l3-3.5m0 0l5 2.5m-5-2.5l6-3M4.5 16l3.236-.462a.6.6 0 0 1 .469.133L11 18l3 3m-6-5.5l2.615-3.05a.6.6 0 0 1 .84-.071L14 14.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinMinusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.992 13h6"/><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinPlusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.992 14h3m3 0h-3m0 0v-3m0 3v3"/><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Binocular(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.5 14L20 9s-.5-2-2.5-2c0 0 0-2-2-2s-2 2-2 2h-3s0-2-2-2s-2 2-2 2C4.5 7 4 9 4 9l-1.5 5"/><path d="M6 20a4 4 0 1 0 0-8a4 4 0 0 0 0 8m12 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path d="M12 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BirthdayCake(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 16.5V20a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-3.5M3 14v-1a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v1m-9-6v3m0-3c1.262 0 2-.968 2-2.625S12 2 12 2s-2 1.718-2 3.375S10.738 8 12 8"/><path d="M9 14a3 3 0 1 1-6 0m12 0a3 3 0 1 1-6 0m12 0a3 3 0 1 1-6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bishop(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 17h10m-8-5h6m-3-9V2m-.446 1.582c-.921 1.251-2.916 4.243-2.497 6.168C9.451 11.558 11.02 12 12 12c.981 0 2.549-.442 2.943-2.25c.42-1.925-1.576-4.917-2.497-6.168a.548.548 0 0 0-.892 0ZM17.8 22H6.2a.617.617 0 0 1-.5-.97c1.316-1.866 4.063-5.986 4.493-8.434c.057-.326.326-.596.657-.596h2.3c.331 0 .6.27.657.596c.43 2.448 3.177 6.568 4.492 8.434a.617.617 0 0 1-.499.97Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitbucket(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20.916 4.674l-1.85 14.8a.6.6 0 0 1-.596.526H5.53a.6.6 0 0 1-.596-.526l-1.85-14.8A.6.6 0 0 1 3.68 4h16.64a.6.6 0 0 1 .596.674"/><path d="m16.75 7.75l-.937 7.97a.6.6 0 0 1-.596.53H8.784a.6.6 0 0 1-.596-.53l-.859-7.3a.6.6 0 0 1 .596-.67zm0 0h3.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M9 12v4.394c0 .332.269.6.6.602c2.966.018 5.4.076 5.4-2.496c0-2.744-3-2.5-6-2.5Zm0 0V7.606c0-.331.269-.6.6-.602C12.566 6.986 15 6.928 15 9.5c0 2.744-3 2.5-6 2.5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 7V5.5m0 13V17m0 5C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m3.75 8.25c0-.763-.182-1.416-.583-1.936c-.397-.515-.947-.817-1.532-1a5.837 5.837 0 0 0-.885-.193V5.5a.75.75 0 0 0-1.5 0v.758a45.994 45.994 0 0 0-1.44-.006l-.214.002A1.353 1.353 0 0 0 8.25 7.606v8.788c0 .746.604 1.348 1.346 1.352l.214.002c.48.003.966.006 1.44-.006v.758a.75.75 0 0 0 1.5 0v-.87c.313-.047.61-.11.885-.195c.585-.182 1.135-.484 1.532-1c.4-.519.583-1.172.583-1.935c0-.83-.231-1.522-.71-2.051a2.73 2.73 0 0 0-.525-.449a2.73 2.73 0 0 0 .526-.449c.478-.529.709-1.22.709-2.051m-2.807 1.551c-.55.138-1.205.185-1.942.199a47.475 47.475 0 0 0-1.251-.004V7.753c1.467-.009 2.625-.01 3.44.244c.394.123.638.287.79.483c.148.193.27.497.27 1.02c0 .541-.144.848-.322 1.045c-.192.212-.502.385-.985.506M11 12.751c-.416.008-.838.006-1.251.003v3.493c1.467.009 2.625.01 3.44-.244c.394-.123.638-.287.79-.483c.148-.193.27-.497.27-1.02c0-.541-.144-.848-.322-1.045c-.192-.212-.502-.385-.985-.506c-.55-.138-1.205-.185-1.942-.199" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinRotateOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21.168 8A10.003 10.003 0 0 0 12 2c-5.185 0-9.45 3.947-9.95 9"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 8h3.4a.6.6 0 0 0 .6-.6V4M2.881 16c1.544 3.532 5.068 6 9.168 6c5.186 0 9.45-3.947 9.951-9"/><path stroke-linecap="round" stroke-linejoin="round" d="M6.05 16h-3.4a.6.6 0 0 0-.6.6V20"/><path d="M9 12v4.394c0 .332.269.6.6.602c2.966.018 5.4.076 5.4-2.496c0-2.744-3-2.5-6-2.5Zm0 0V7.606c0-.331.269-.6.6-.602C12.566 6.986 15 6.928 15 9.5c0 2.744-3 2.5-6 2.5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 7V5.5m0 13V17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.75 8l10.5 8.5l-5.5 5.5V2l5.5 5.5L6.75 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m9 9.6l6 5.1l-3.143 3.3V6L15 9.3l-6 5.1"/><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothTagSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.25A6.75 6.75 0 0 0 1.25 9v6A6.75 6.75 0 0 0 8 21.75h8A6.75 6.75 0 0 0 22.75 15V9A6.75 6.75 0 0 0 16 2.25zm4.4 3.233A.75.75 0 0 0 11.107 6v4.407L9.486 9.029a.75.75 0 1 0-.972 1.143L10.666 12l-2.152 1.829a.75.75 0 1 0 .972 1.143l1.621-1.379V18a.75.75 0 0 0 1.293.517l3.143-3.3a.75.75 0 0 0-.057-1.088L12.982 12l2.504-2.129a.75.75 0 0 0 .057-1.088zm1.51 9.275l-1.303-1.108v2.475zm-1.303-6.883v2.475l1.303-1.107z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M12 11.667H8m4 0s3.333 0 3.333-3.334C15.333 5 12 5 12 5H8.6a.6.6 0 0 0-.6.6v6.067m4 0s4 0 4 3.666C16 19 12 19 12 19H8.6a.6.6 0 0 1-.6-.6v-6.733"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path d="M12 12H9m3 0s2.5 0 2.5-2.5S12 7 12 7H9.6a.6.6 0 0 0-.6.6V12m3 0s3 0 3 2.75s-3 2.75-3 2.75H9.6a.6.6 0 0 1-.6-.6V12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm6 4A1.35 1.35 0 0 0 8.25 7.6v9.3c0 .746.604 1.35 1.35 1.35H12v-.75v.75h.02a2.51 2.51 0 0 0 .158-.007c.1-.006.24-.02.404-.045c.326-.05.773-.15 1.23-.36c.459-.21.95-.54 1.326-1.057c.382-.525.612-1.198.612-2.031c0-.833-.23-1.506-.612-2.031a3.182 3.182 0 0 0-.877-.815a2.8 2.8 0 0 0 .472-.543c.326-.488.517-1.105.517-1.861s-.191-1.373-.517-1.861a2.943 2.943 0 0 0-1.148-.997A3.85 3.85 0 0 0 12 6.25zm.15 6.5h2.261l.069.004a3.184 3.184 0 0 1 1.108.272c.291.133.55.319.737.575c.18.249.325.607.325 1.149s-.145.9-.325 1.149a1.837 1.837 0 0 1-.738.575a3.184 3.184 0 0 1-1.107.272a1.938 1.938 0 0 1-.07.004H9.75zm2.247 4H12zm-.001-5.5H9.75v-3.5h2.249h-.001h.001c.03.001.48.016.916.233c.223.112.424.268.57.488c.144.215.265.535.265 1.029c0 .494-.121.814-.265 1.03a1.45 1.45 0 0 1-.57.487a2.322 2.322 0 0 1-.912.233z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bonfire(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M9 14c0 1.61 1.377 2 3.076 2c2.89 0 3.845-1.667 1.922-5c-2.691 3-3.076-1.667-2.691-3C10.153 10 9 11.879 9 14"/><path stroke-linejoin="round" d="M12 16c3.156 0 5-2.098 5-5.687C17 6.723 12 3 12 3s-5 3.723-5 7.313S8.844 16 12 16"/><path d="m4.273 21.07l15.454-4.14m-15.454 0L12 19m7.727 2.07l-3.863-1.035"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/><path d="M9 7h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookLock(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M14 10h.4a.6.6 0 0 1 .6.6v2.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6v-2.8a.6.6 0 0 1 .6-.6h.4m4 0V8c0-.667-.4-2-2-2s-2 1.333-2 2v2m4 0h-4"/><path d="M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.75 16.714a.755.755 0 0 1-.014.143a.75.75 0 0 1-.736.893H6a1.25 1.25 0 1 0 0 2.5h14a.75.75 0 0 1 0 1.5H6A2.75 2.75 0 0 1 3.25 19V5A2.75 2.75 0 0 1 6 2.25h13.4c.746 0 1.35.604 1.35 1.35zM9 6.25a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookStack(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M5 19.5V5a2 2 0 0 1 2-2h11.4a.6.6 0 0 1 .6.6V21M9 7h6m-8.5 8H19M6.5 18H19M6.5 21H19"/><path stroke-linejoin="round" d="M6.5 18c-1 0-1.5-.672-1.5-1.5S5.5 15 6.5 15m0 6c-1 0-1.5-.672-1.5-1.5S5.5 18 6.5 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 21V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16l-5.918-3.805a2 2 0 0 0-2.164 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkBook(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M8 3v8l2.5-1.6L13 11V3"/><path d="M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16v-6a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v6l-1.89-1.26a2 2 0 0 0-2.22 0z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m-1 6A2.75 2.75 0 0 0 8.25 10v6a.75.75 0 0 0 1.166.624l1.89-1.26c.42-.28.968-.28 1.387 0l1.891 1.26A.75.75 0 0 0 15.75 16v-6A2.75 2.75 0 0 0 13 7.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 21V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16l-5.918-3.805a2 2 0 0 0-2.164 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBl(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4v16h16m.01-4l-.011.01M20.01 12l-.011.01M20.01 8l-.011.01M20.01 4l-.011.01M8.01 4l-.011.01M12.01 4l-.011.01M16.01 4l-.011.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottom(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 4.01l.01-.011M16 4.01l.01-.011M12 4.01l.01-.011M8 4.01l.01-.011M4 4.01l.01-.011M4 8.01l.01-.011M4 12.01l.01-.011m7.99.011l.01-.011M4 16.01l.01-.011M20 8.01l.01-.011M20 12.01l.01-.011M20 16.01l.01-.011M4 20h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBr(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.01 4v16h-16M4 16l.011.01M4 12l.011.01M4 8l.011.01M4 4l.011.01M16 4l.011.01M12 4l.011.01M8 4l.011.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInner(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 20.01l.01-.011M4 16.01l.01-.011M4 8.01l.01-.011M4 4.01l.01-.011M8 4.01l.01-.011M16 4.01l.01-.011M20 4.01l.01-.011M20 8.01l.01-.011M8 20.01l.01-.011m7.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M4 12h8m8 0h-8m0 0V4m0 8v8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.01 20l-.011.01M20.01 16l-.011.01M20.01 12l-.011.01M20.01 8l-.011.01M20.01 4l-.011.01M8.01 4l-.011.01M12.01 4l-.011.01M12.01 12l-.011.01M16.01 4l-.011.01M8.01 20l-.011.01M12.01 20l-.011.01M16.01 20l-.011.01M4 4v16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12.01 16l-.01.011M12.01 12l-.01.011M12.01 8l-.01.011M8.01 12l-.01.011M16.01 12l-.01.011M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 20l.01.01M4 16l.01.01M4 12l.01.01M4 8l.01.01M4 4l.01.01M16 4l.01.01M12 4l.01.01M12 12l.01.01M8 4l.01.01M16 20l.01.01M12 20l.01.01M8 20l.01.01M20.01 4v16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTl(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 20.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M20 12.01l.01-.011M20 8.01l.01-.011M4 20V4h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTop(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 20.01l.01-.011M16 20.01l.01-.011M12 20.01l.01-.011M8 20.01l.01-.011M4 20.01l.01-.011M4 8.01l.01-.011M4 12.01l.01-.011m7.99.011l.01-.011M4 16.01l.01-.011M20 8.01l.01-.011M20 12.01l.01-.011M20 16.01l.01-.011M4 4h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTr(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.01 20V4h-16M4 8l.01-.01M4 12l.01-.01M4 16l.01-.01M4 20l.01-.01M16 20l.01-.01M12 20l.01-.01M8 20l.01-.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BounceLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4m15 8.5c-3-1-5.5-.5-8 4.5c-.5-3-2-7.5-3.5-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BounceRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4M4 15.5c3-1 5.5-.5 8 4.5c.5-3 2-7.5 3.5-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowlingBall(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path fill="currentColor" d="M11.5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-4 3a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m4 2a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12h4M3 3h18m0 4v13.4a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxIso(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke-width="1.5"><path fill="currentColor" d="m2.695 7.185l9 4l.61-1.37l-9-4zM12.75 21.5v-11h-1.5v11zm-.445-10.315l9-4l-.61-1.37l-9 4z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 17.11V6.89a.6.6 0 0 1 .356-.548l8.4-3.734a.6.6 0 0 1 .488 0l8.4 3.734A.6.6 0 0 1 21 6.89v10.22a.6.6 0 0 1-.356.548l-8.4 3.734a.6.6 0 0 1-.488 0l-8.4-3.734A.6.6 0 0 1 3 17.11"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 4.5l8.644 3.842a.6.6 0 0 1 .356.548v3.61"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDcenter(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524"/><path d="m20.5 16.722l-8.209-4.56a.6.6 0 0 0-.582 0L3.5 16.722m.028-9.428l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 3v9m0 7.5V22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDpoint(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M3 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524"/><path d="m3.528 7.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21v-9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDthreePoints(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2M3 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524"/><path d="m3.528 7.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21v-9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxingGlove(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.489 17.727h9.867m-9.867 0V21h9.867v-3.273m-9.867 0C5.2 15.546 3.556 10.091 4.104 8.455c.438-1.31 2.375-.91 3.289-.546C7.393 4.091 9.037 3 13.423 3C17.806 3 20 4.09 20 9.545c0 4.364-1.096 7.273-1.644 8.182"/><path d="M7.393 7.91C7.758 8.272 8.818 9 10.133 9h4.934M7.393 7.91c0 3.817 1.644 4.363 2.74 4.363"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brain(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a3 3 0 1 0 1 5.83"/><path d="M4.264 15.605a4 4 0 0 1-.874-6.636m.03-.081A2.5 2.5 0 0 1 7 5.5m.238.065A2.5 2.5 0 1 1 12 4.5V20m-4 0a2 2 0 1 0 4 0m0-13a3 3 0 0 0 3 3m2 4a3 3 0 1 1-1 5.83"/><path d="M19.736 15.605a4 4 0 0 0 .874-6.636m-.03-.081A2.5 2.5 0 0 0 17 5.5m-5-1a2.5 2.5 0 1 1 4.762 1.065M16 20a2 2 0 1 1-4 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrainElectricity(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a3 3 0 1 0 1 5.83"/><path d="M4.264 15.605a4 4 0 0 1-.874-6.636m.03-.081A2.5 2.5 0 0 1 7 5.5m.238.065A2.5 2.5 0 1 1 12 4.5V20m-4 0a2 2 0 1 0 4 0m0-13a3 3 0 0 0 3 3m5.61-1.031A3.99 3.99 0 0 1 22 12c0 .703-.181 1.364-.5 1.938m-.92-5.05A2.5 2.5 0 0 0 17 5.5m-5-1a2.5 2.5 0 1 1 4.762 1.065M14 22a2 2 0 0 1-2-2m6.667-4L17 19h4l-1.667 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrainResearch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a3 3 0 1 0 1 5.83"/><path d="M4.264 15.605a4 4 0 0 1-.874-6.636m.03-.081A2.5 2.5 0 0 1 7 5.5m.238.065A2.5 2.5 0 1 1 12 4.5V20m-4 0a2 2 0 1 0 4 0m0-13a3 3 0 0 0 3 3m5.61-1.031A3.99 3.99 0 0 1 22 12c0 .703-.181 1.364-.5 1.938m-.92-5.05A2.5 2.5 0 0 0 17 5.5m-5-1a2.5 2.5 0 1 1 4.762 1.065M14 22a2 2 0 0 1-2-2m8.5.5L22 22m-6-3.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrainWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a3 3 0 1 0 1 5.83"/><path d="M4.264 15.605a4 4 0 0 1-.874-6.636m.03-.081A2.5 2.5 0 0 1 7 5.5m5-1a2.5 2.5 0 1 0-4.762 1.065M8 20a2 2 0 1 0 4 0m5-6a3 3 0 1 1-1 5.83"/><path d="M19.736 15.605a4 4 0 0 0 .874-6.636m-.03-.081A2.5 2.5 0 0 0 17 5.5m-5-1a2.5 2.5 0 1 1 4.762 1.065M16 20a2 2 0 1 1-4 0m0-12v4m0 4.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BreadSlice(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 20V9S3 4 9.5 4H17c7 0 3 5 3 5v9a2 2 0 0 1-2 2z"/><path d="M7 20H6a2 2 0 0 1-2-2V9S0 4 6.5 4H10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BridgeSurface(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0 4 2 9 8 9M10 6c0 4 2 9 8 9M3 8.5v-2M10 3V1M3 12l7-6m1 15l7-6m-3.5 6h2m4.5-6h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BridgeThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 4h3"/><path fill="currentColor" d="M10 21a1 1 0 1 0 0-2a1 1 0 0 0 0 2m4-16a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M10 20s6.5-2.5 2-8s2-8 2-8M3 20h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightCrown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 12h1M12 2V1m0 22v-1m8-2l-1-1m1-15l-1 1M4 20l1-1M4 4l1 1m-4 7h1m14.8 3.5l1.2-7l-4.2 2.1L12 8.5l-1.8 2.1L6 8.5l1.2 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightStar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.952 9.623l1.559-3.305a.535.535 0 0 1 .978 0l1.559 3.305l3.485.533c.447.068.625.644.302.974l-2.522 2.57l.595 3.631c.077.467-.391.822-.791.602L12 16.218l-3.117 1.715c-.4.22-.868-.135-.791-.602l.595-3.63l-2.522-2.571c-.323-.33-.145-.906.302-.974zM22 12h1M12 2V1m0 22v-1m8-2l-1-1m1-15l-1 1M4 20l1-1M4 4l1 1m-4 7h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brightness(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke-width="1.5"><path stroke="currentColor" d="m12 7l1.53 1.304l2.006.16l.16 2.005L17 12l-1.305 1.53l-.16 2.006l-2.004.16L12 17l-1.53-1.305l-2.006-.16l-.16-2.004L7 12l1.304-1.53l.16-2.006l2.005-.16z"/><path fill="currentColor" d="M10.47 15.696L12 17V7l-1.53 1.304l-2.006.16l-.16 2.005L7 12l1.304 1.53l.16 2.006z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" d="M12 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v7"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/><path stroke="currentColor" d="m18 14l1.225 1.044l1.603.128l.128 1.603L22 18l-1.044 1.225l-.128 1.603l-1.603.128L18 22l-1.225-1.044l-1.603-.128l-.128-1.603L14 18l1.044-1.225l.128-1.603l1.603-.128z"/><path fill="currentColor" d="M16.775 20.956L18 22v-8l-1.225 1.044l-1.603.128l-.128 1.603L14 18l1.044 1.225l.128 1.603z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleDownload(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 2v6m0 0l3-3m-3 3l-3-3m-3-2.95c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleIncome(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 5h-6m0 0l3-3m-3 3l3 3m-6-5.95c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleOutcome(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 5h6m0 0l-3 3m3-3l-3-3m-6 .05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleSearch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 6.5L22 8m-6-3.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0"/><path d="M13 2.05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleSearchSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M22.586 10.173a2.249 2.249 0 0 1-2.177-.582l-.541-.541a4.75 4.75 0 0 1-5.19-7.371a7.652 7.652 0 0 0-1.604-.376A10.871 10.871 0 0 0 12 1.25C6.063 1.25 1.25 6.063 1.25 12c0 1.856.471 3.605 1.3 5.13l-.787 4.233a.75.75 0 0 0 .874.874l4.233-.788A10.702 10.702 0 0 0 12 22.75c5.937 0 10.75-4.813 10.75-10.75c0-.362-.018-.72-.053-1.074c0 0-.045-.325-.111-.753M19.97 5.97a.75.75 0 0 1 1.06 0l1.5 1.5a.75.75 0 0 1-1.06 1.06l-1.5-1.5a.75.75 0 0 1 0-1.06"/><path d="M18.5 2.75a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5M15.25 4.5a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleStar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.306 4.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.233.128-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568z"/><path d="M13 2.05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleUpload(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 8V2m0 0l3 3m-3-3l-3 3m-3-2.95c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 2v3m0 4.01l.01-.011M13 2.05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 2.05c-.329-.033-.662-.05-1-.05C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10c0-.338-.017-.671-.05-1m-4.829-3.636l2.121-2.121m0 0l2.122-2.122m-2.122 2.122l-2.12-2.122m2.12 2.122l2.122 2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10 9.01l.01-.011M14 9.01l.01-.011M10 13.01l.01-.011m3.99.011l.01-.011M10 17.01l.01-.011m3.99.011l.01-.011M6 20.4V5.6a.6.6 0 0 1 .6-.6H12V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 16.01l.01-.011m9.99.011l.01-.011M3 12h18v7a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zm18-4V6a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2m4 0h10"/><path d="M4.5 20v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20m7 0v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusGreen(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 16.01l.01-.011M21 12H3v7a1 1 0 0 0 1 1h9m4 3s.9-3.118 3-5"/><path stroke-linejoin="round" d="m19.802 21.425l-.134.012a3.094 3.094 0 0 1-3.366-2.774a3.06 3.06 0 0 1 2.761-3.35l2.986-.28a.35.35 0 0 1 .381.314l.255 2.58a3.194 3.194 0 0 1-2.883 3.497M21 8V6a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2m4 0h10"/><path d="M4.5 20v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusStop(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m16 16.01l.01-.011M6 16.01l.01-.011M20 22V8m0 0h-2V2h4v6zm-4 12H2.6a.6.6 0 0 1-.6-.6v-6.8a.6.6 0 0 1 .6-.6H16m-2-4H6m8-6H6a4 4 0 0 0-4 4v2"/><path d="M3.5 20v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20m7 0v1.9a.6.6 0 0 0 .6.6h.9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Csquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M14 10v-.2A1.8 1.8 0 0 0 12.2 8h-.4A1.8 1.8 0 0 0 10 9.8v4.4a1.8 1.8 0 0 0 1.8 1.8h.4a1.8 1.8 0 0 0 1.8-1.8V14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CableTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.667 8L10 12h4l-1.667 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CableTagSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.25A6.75 6.75 0 0 0 1.25 9v6A6.75 6.75 0 0 0 8 21.75h8A6.75 6.75 0 0 0 22.75 15V9A6.75 6.75 0 0 0 16 2.25zm4.359 6.039a.75.75 0 0 0-1.385-.577l-1.666 4A.75.75 0 0 0 10 12.75h2.875l-1.234 2.962a.75.75 0 0 0 1.385.577l1.666-4A.75.75 0 0 0 14 11.25h-2.875z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 21V3a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v18a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 7h4m-4 8.5h4m-4 3h4M5 7h2m2 0H7m0 0V5m0 2v2m-1.414 9.414L7 17m1.415-1.414L7 17m0 0l-1.414-1.414M7 17l1.415 1.414"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 4V2m0 2v2m0-2h-4.5M3 10v9a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-9zm0 0V6a2 2 0 0 1 2-2h2m0-2v4m14 4V6a2 2 0 0 0-2-2h-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 21H5a2 2 0 0 1-2-2v-9h18v3m-6-9V2m0 2v2m0-2h-4.5M3 10V6a2 2 0 0 1 2-2h2m0-2v4m14 4V6a2 2 0 0 0-2-2h-.5m-3.508 14H21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 21H5a2 2 0 0 1-2-2v-9h18v3m-6-9V2m0 2v2m0-2h-4.5M3 10V6a2 2 0 0 1 2-2h2m0-2v4m14 4V6a2 2 0 0 0-2-2h-.5m-3.508 14h3M21 18h-3.008m0 0v-3m0 3v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 19V9a2 2 0 0 1 2-2h.5a2 2 0 0 0 1.6-.8l2.22-2.96A.6.6 0 0 1 8.8 3h6.4a.6.6 0 0 1 .48.24L17.9 6.2a2 2 0 0 0 1.6.8h.5a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2"/><path d="M12 17a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CandlestickChart(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 16v-2m7 7v-2m7-6v-2M5 8V6m7 7v-2m7-6V3M7 8.6v4.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V8.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6m7 5v4.8a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6v-4.8a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6m7-8v4.8a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8 10h8m-9 4h1m8 0h1"/><path d="M3 18v-6.59a2 2 0 0 1 .162-.787l2.319-5.41A2 2 0 0 1 7.319 4h9.362a2 2 0 0 1 1.838 1.212l2.32 5.41a2 2 0 0 1 .161.789V18M3 18v2.4a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V18m-4 0h4m14 0v2.4a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6V18m4 0h-4M7 18h10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardLock(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 9V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10m8-10H6m16 0v4m-.833 5.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardNoAccess(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 9V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8M22 9H6m16 0v4m-1 3.05a3.5 3.5 0 1 0-5 4.9m5-4.9a3.5 3.5 0 0 1-5 4.9m5-4.9l-5 4.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardReader(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5 19V3h14v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2Z"/><path d="M5 6H3.5a1.5 1.5 0 1 1 0-3h17a1.5 1.5 0 0 1 0 3H19"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 3v18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardShield(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 9V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8M22 9H6m16 0v2"/><path d="m18.992 14.125l2.556.649c.266.068.453.31.445.584C21.821 21.116 18.5 22 18.5 22s-3.321-.884-3.493-6.642a.588.588 0 0 1 .445-.584l2.556-.649c.323-.082.661-.082.984 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardWallet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 20H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2ZM7 7V3.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6V7m-7-4v4m2-4v4"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 14a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M19.5 22a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-10 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/><path d="M5 4h17l-2 11H7zm0 0c-.167-.667-1-2-3-2m18 13H5.23c-1.784 0-2.73.781-2.73 2c0 1.219.946 2 2.73 2H19.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M19.5 22a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-10 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/><path d="M16.5 4H22l-2 11h-4.5m1-11l-1 11m1-11h-5.75m4.75 11h-4m-.75-11H5l2 11h4.5m-.75-11l.75 11M5 4c-.167-.667-1-2-3-2m18 13H5.23c-1.784 0-2.73.781-2.73 2c0 1.219.946 2 2.73 2H19.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h19l-3 10H6zm0 0l-.75-2.5M9.992 11h4M11 19.5a1.5 1.5 0 0 1-3 0m9 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h19l-3 10H6zm0 0l-.75-2.5M9.992 11h2m2 0h-2m0 0V9m0 2v2M11 19.5a1.5 1.5 0 0 1-3 0m9 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 17V7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2"/><path d="M12 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6m6.5-2.99l.01-.011M5.5 12.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CellTwoXtwo(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M21 3.6V12h-9V3h8.4a.6.6 0 0 1 .6.6Zm0 16.8V12h-9v9h8.4a.6.6 0 0 0 .6-.6ZM3 12V3.6a.6.6 0 0 1 .6-.6H12v9zm0 0v8.4a.6.6 0 0 0 .6.6H12v-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cellar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 21h18v-9a9 9 0 1 0-18 0zm0-4h18"/><path d="M9 17v-4h12m-8 0V9h7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CenterAlign(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 16.01l.01-.011M4 20.01l.01-.011M4 8.01l.01-.011M4 4.01l.01-.011M4 12.01l.01-.011M8 20.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M20 12.01l.01-.011M20 8.01l.01-.011M20 4.01l.01-.011M16 4.01l.01-.011M12 4.01l.01-.011M8 4.01l.01-.011M8 16V8h8v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CenterAlignSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4 16.01l.01-.011M4 20.01l.01-.011M4 8.01l.01-.011M4 4.01l.01-.011M4 12.01l.01-.011M8 20.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M20 12.01l.01-.011M20 8.01l.01-.011M20 4.01l.01-.011M16 4.01l.01-.011M12 4.01l.01-.011M8 4.01l.01-.011"/><path fill="currentColor" d="M8 16V8h8v8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubble(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M17 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 12l3 3l5-5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleEmpty(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleQuestion(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 9c0-3.5 5.5-3.5 5.5 0c0 2.5-2.5 2-2.5 5m0 4.01l.01-.011"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleTranslate(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22"/><path d="M7 8.517h5m5 0h-1.786m-3.214 0h3.214m-3.214 0V7m3.214 1.517c-.586 2.075-1.813 4.037-3.214 5.76M8.429 18C9.56 16.97 10.84 15.705 12 14.277m0 0c-.714-.829-1.714-2.17-2-2.777m2 2.777l2.143 2.206"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v4m0 4.01l.01-.011M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.5 14.5l2.493-2.5M14.5 9.5L11.993 12m0 0L9.5 9.5m2.493 2.5l2.507 2.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatLines(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 10h8m-8 4h4m0 8c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatMinusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22M9 12h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatPlusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h3m3 0h-3m0 0V9m0 3v3m0 7c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5L2.5 21.5l4.5-.838A9.955 9.955 0 0 0 12 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 13l4 4L19 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7 12.5l3 3l7-7"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25M7.53 11.97a.75.75 0 0 0-1.06 1.06l3 3a.75.75 0 0 0 1.06 0l7-7a.75.75 0 0 0-1.06-1.06L10 14.44z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chocolate(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 6.5c-3 0-4.5-.5-4.5-3.5H5v18h14zm0 8.5H5m0-6h14m-7 12V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chromecast(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2 20.01l.01-.011M15 20h5a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v1m0 9c2 .5 3.5 2 4 4m-4-8c4 .5 7.5 4 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromecastActive(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2 20.01l.01-.011M15 20h5a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v1m0 9c2 .5 3.5 2 4 4m-4-8c4 .5 7.5 4 8 8"/><path fill="currentColor" fill-rule="evenodd" d="M5.002 7.63a.6.6 0 0 1 .6-.6h12.804a.6.6 0 0 1 .6.6v8.832a.6.6 0 0 1-.6.6H13.44a.617.617 0 0 1-.556-.355c-.422-.892-1.622-3.26-3.07-4.707c-1.42-1.419-3.572-2.444-4.435-2.82a.624.624 0 0 1-.378-.569z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Church(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 6l-7.718 4.824a.6.6 0 0 0-.282.508V21.4a.6.6 0 0 0 .6.6H12m0-16l7.718 4.824a.6.6 0 0 1 .282.508V21.4a.6.6 0 0 1-.6.6H12m0-16V4m0-2v2m-2 0h2m0 0h2m-2 18v-5m4 .01l.01-.011M16 13.01l.01-.011M12 13.01l.01-.011M8 13.01l.01-.011M8 17.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChurchSide(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.576 7.424a.6.6 0 0 1 .848 0l3.4 3.4a.6.6 0 0 1 .176.425V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V11.249a.6.6 0 0 1 .176-.425zM8 7V4m0-2v2m0 0H6m2 0h2"/><path d="M12 22h7.4a.6.6 0 0 0 .6-.6V10.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 7H8m0 15v-5m0-3.99l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CigaretteSlash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 15v3m0-7c0-1-1-2-3-2h-1a3 3 0 0 1-3-3V2m11 9c0-6-4-6-4-6s4 1 4-3m0 13v3"/><path d="M2.6 18H18l-3-3H2.6a.6.6 0 0 0-.6.6v1.8a.6.6 0 0 0 .6.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m3 3l18 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CinemaOld(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m10 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-5-5a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12m0 0v10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleSpark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2"/><path d="M13 6.5c3.134 0 4.5-1.318 4.5-4.5c0 3.182 1.357 4.5 4.5 4.5c-3.143 0-4.5 1.357-4.5 4.5c0-3.143-1.366-4.5-4.5-4.5Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func City(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 9.01l.01-.011M11 9.01l.01-.011M7 13.01l.01-.011m3.99.011l.01-.011M7 17.01l.01-.011m3.99.011l.01-.011M15 21H3.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6H9V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6V9m0 12h5.4a.6.6 0 0 0 .6-.6V9.6a.6.6 0 0 0-.6-.6H15m0 12v-4m0-8v4m0 0h2m-2 0v4m0 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M8.5 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6m3.5-18H18a2 2 0 0 1 2 2v9"/><path d="M8 6.4V4.5a.5.5 0 0 1 .5-.5c.276 0 .504-.224.552-.496C9.2 2.652 9.774 1 12 1s2.8 1.652 2.948 2.504c.048.272.276.496.552.496a.5.5 0 0 1 .5.5v1.9a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linejoin="round" d="m15.5 20.5l2 2l5-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6h6"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockRotateRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6h6"/><path d="M21.888 10.5C21.164 5.689 17.013 2 12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c4.1 0 7.625-2.468 9.168-6"/><path d="M17 16h4.4a.6.6 0 0 1 .6.6V21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25M12.75 6a.75.75 0 0 0-1.5 0v6c0 .414.336.75.75.75h6a.75.75 0 0 0 0-1.5h-5.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClosedCaptionsTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 15V9a6 6 0 0 1 6-6h10a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H7a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" d="m10.5 10l-.172-.172a2.828 2.828 0 0 0-2-.828v0A2.828 2.828 0 0 0 5.5 11.828v.344A2.828 2.828 0 0 0 8.328 15v0c.75 0 1.47-.298 2-.828L10.5 14m8-4l-.172-.172a2.828 2.828 0 0 0-2-.828v0a2.828 2.828 0 0 0-2.828 2.828v.344A2.828 2.828 0 0 0 16.328 15v0c.75 0 1.47-.298 2-.828L18.5 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Closet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M9 14H8m8 0h-1"/><path d="M12 2h8.4a.6.6 0 0 1 .6.6v18.8a.6.6 0 0 1-.6.6H12m0-20H3.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H12m0-20v20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M12 4c-6 0-6 4-6 6c-1.667 0-5 1-5 5s3.333 5 5 5h12c1.667 0 5-1 5-5s-3.333-5-5-5c0-2 0-6-6-6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudBookmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.5 12h7v10L12 20l-3.5 2z"/><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 18l3 3l5-5"/><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDesync(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607m12.42 1.88l-1.768 1.768a4 4 0 0 1-5.656 0l-.354-.353"/><path d="m16.067 21.962l.353-2.475l-2.475.354zm-8.487-5.06l1.768-1.768a4 4 0 0 1 5.657 0l.354.353"/><path d="m7.934 14.427l-.353 2.475l2.474-.354z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 13v9m0 0l3.5-3.5M12 22l-3.5-3.5m11.5-.893c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linejoin="round" d="M12 8c-3.273 0-3.273 2-3.273 3C7.818 11 6 11.5 6 13.5S7.818 16 8.727 16h6.546c.909 0 2.727-.5 2.727-2.5S16.182 11 15.273 11c0-1 0-3-3.273-3Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm8.4 5c-1.792 0-2.897.557-3.491 1.464c-.35.536-.47 1.128-.511 1.609c-.445.085-.94.255-1.388.551c-.76.502-1.36 1.353-1.36 2.626s.6 2.124 1.36 2.626c.72.475 1.555.624 2.117.624h6.546c.562 0 1.398-.149 2.117-.624c.76-.502 1.36-1.353 1.36-2.626s-.6-2.124-1.36-2.626a3.856 3.856 0 0 0-1.388-.551c-.04-.481-.16-1.073-.51-1.609c-.594-.907-1.7-1.464-3.492-1.464" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSunny(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 13c-1.667 0-5 1-5 5s3.333 5 5 5h12c1.667 0 5-1 5-5s-3.333-5-5-5m-6-1a3 3 0 1 0 0-6a3 3 0 0 0 0 6m7-3h1m-8-7V1m6.5 2.5l-1 1m-12-1l1 1M4 9h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSync(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607m3.58 1.88l1.768 1.768a4 4 0 0 0 5.657 0l.354-.353"/><path d="m7.934 21.962l-.353-2.475l2.474.354zm8.364-5.06l-1.768-1.768a4 4 0 0 0-5.657 0l-.353.353"/><path d="m15.945 14.427l.353 2.475l-2.475-.354z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22v-9m0 0l3.5 3.5M12 13l-3.5 3.5M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 22l3-3m3-3l-3 3m0 0l-3-3m3 3l3 3m5-4.393c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.5 6L10 18.5m-3.5-10L3 12l3.5 3.5m11-7L21 12l-3.5 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBrackets(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 21H8c-1.105 0-2-.894-2-1.999V14c0-1-1.5-2-1.5-2S6 11 6 10V5a2 2 0 0 1 2-2h1m6 18h1c1.105 0 2-.894 2-1.999V14c0-1 1.5-2 1.5-2S18 11 18 10V5a2 2 0 0 0-2-2h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBracketsSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 17h-.667a2 2 0 0 1-2-2v-1.889C7.333 12.556 6 12 6 12s1.333-.556 1.333-1.111V9a2 2 0 0 1 2-2H10m4 10h.667a2 2 0 0 0 2-2v-1.889C16.667 12.556 18 12 18 12s-1.333-.556-1.333-1.111V9a2 2 0 0 0-2-2H14"/><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Codepen(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 9v6M3 15V9m9 12v-6m0-12v6m0 6L3 9l9-6l9 6z"/><path d="m12 21l-9-6l9-6l9 6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeCup(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 11.6V15a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6v-3.4a.6.6 0 0 1 .6-.6h12.8a.6.6 0 0 1 .6.6M12 9c0-1 .714-2 2.143-2v0A2.857 2.857 0 0 0 17 4.143V3.5M8 9v-.5a3 3 0 0 1 3-3v0a2 2 0 0 0 2-2V3"/><path d="M16 11h2.5a2.5 2.5 0 0 1 0 5H17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinSlash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.623 5.248A9.964 9.964 0 0 0 2 12c0 5.523 4.477 10 10 10a9.962 9.962 0 0 0 6.615-2.5m2.687-3.822A9.974 9.974 0 0 0 22 12c0-5.523-4.477-10-10-10c-1.231 0-2.41.223-3.5.63"/><path d="M9 15c.644.86 1.843 1.35 3 1.391c1.114.04 2.19-.336 2.697-1.198M12 16.391V18.5m-2.5-9c0 1.181.852 1.665 1.886 2M15 8.5c-.685-.685-1.891-1.161-3-1.191V5.5M3 3l18 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coins(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 13c-2.761 0-5-1.12-5-2.5S13.239 8 16 8s5 1.12 5 2.5s-2.239 2.5-5 2.5m-5 1.5c0 1.38 2.239 2.5 5 2.5s5-1.12 5-2.5m-18-5C3 10.88 5.239 12 8 12c1.126 0 2.165-.186 3-.5M3 13c0 1.38 2.239 2.5 5 2.5c1.126 0 2.164-.186 3-.5"/><path d="M3 5.5v11C3 17.88 5.239 19 8 19c1.126 0 2.164-.186 3-.5m2-10v-3m-2 5v8c0 1.38 2.239 2.5 5 2.5s5-1.12 5-2.5v-8"/><path d="M8 8C5.239 8 3 6.88 3 5.5S5.239 3 8 3s5 1.12 5 2.5S10.761 8 8 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinsSwap(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.019 9A6.5 6.5 0 1 1 15 14.981"/><path d="M8.5 22a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13M22 17a3 3 0 0 1-3 3h-2m0 0l2-2m-2 2l2 2M2 7a3 3 0 0 1 3-3h2m0 0L5 6m2-2L5 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollageFrame(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M19.4 20H4.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6ZM11 12V4m-7 8h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collapse(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 20l-5-5m0 0v4m0-4h4M4 20l5-5m0 0v4m0-4H5M20 4l-5 5m0 0V5m0 4h4M4 4l5 5m0 0V5m0 4H5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorFilter(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 14.5a6 6 0 1 0 0-12a6 6 0 0 0 0 12"/><path d="M16 21.5a6 6 0 1 0 0-12a6 6 0 0 0 0 12"/><path d="M8 21.5a6 6 0 1 0 0-12a6 6 0 0 0 0 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorPicker(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 13.161l5.464-5.464a1 1 0 0 1 1.415 0l2.12 2.12a1 1 0 0 1 0 1.415l-1.928 1.929m-7.071 0l-2.172 2.172a.999.999 0 0 0-.218.327l-1.028 2.496c-.508 1.233.725 2.466 1.958 1.959l2.497-1.028a.998.998 0 0 0 .326-.218l5.708-5.708m-7.071 0h7.071m-.193-9.707l2.121 2.121m4.243 4.243l-2.121-2.121m-2.122-2.122l1.414-1.414a1 1 0 0 1 1.415 0l.707.707a1 1 0 0 1 0 1.414L18.12 7.697m-2.122-2.122l2.122 2.122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorPickerEmpty(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.879 7.697L16 9.817a1 1 0 0 1 0 1.415L8.363 18.87a1.001 1.001 0 0 1-.326.218L5.54 20.114c-1.233.508-2.466-.725-1.958-1.958L4.61 15.66a.999.999 0 0 1 .218-.327l7.636-7.636a1 1 0 0 1 1.415 0Zm0-4.243L16 5.575m4.243 4.243L18.12 7.697M16 5.575l1.413-1.414a1 1 0 0 1 1.414 0l.708.707a1 1 0 0 1 0 1.414L18.12 7.697M16 5.575l2.12 2.122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorWheel(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/><path d="M12 16a4 4 0 1 1 0-8a4 4 0 0 1 0 8m0-14v6m0 8v6M2 12h6m8 0h6M4.929 4.929L9.172 9.17m5.656 5.659l4.243 4.242m-14.142 0l4.243-4.242m5.656-5.658l4.243-4.242"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Combine(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 9.6v10.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6V9.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6"/><path d="M15 3.6v10.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Commodity(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m12.147 18.28l1.184-5.8a.6.6 0 0 1 .588-.48h6.162a.6.6 0 0 1 .588.48l1.184 5.8a.6.6 0 0 1-.588.72h-8.53a.6.6 0 0 1-.588-.72Z"/><path d="m7.147 11.28l1.184-5.8A.6.6 0 0 1 8.918 5h6.164a.6.6 0 0 1 .587.48l1.184 5.8a.6.6 0 0 1-.588.72h-8.53a.6.6 0 0 1-.588-.72Z"/><path d="m2.147 18.28l1.184-5.8a.6.6 0 0 1 .587-.48h6.163a.6.6 0 0 1 .588.48l1.184 5.8a.6.6 0 0 1-.588.72h-8.53a.6.6 0 0 1-.588-.72Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Community(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 18v-1a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v1M1 18v-1a3 3 0 0 1 3-3v0m19 4v-1a3 3 0 0 0-3-3v0m-8-2a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-8 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4m16 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompAlignBottom(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 21H2"/><path d="M8 15V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompAlignBottomSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 21H2"/><path fill="currentColor" d="M8 15V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompAlignLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3 22V2"/><path d="M19 16H9a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompAlignLeftSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3 22V2"/><path fill="currentColor" d="M19 16H9a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompAlignRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 22V2"/><path d="M15 16H5a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompAlignRightSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 22V2"/><path fill="currentColor" d="M15 16H5a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompAlignTop(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 3H2"/><path d="M8 19V9a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompAlignTopSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 3H2"/><path fill="currentColor" d="M8 19V9a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompactDisc(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.586 10.586L16.95 7.05l-3.536 6.364L7.05 16.95z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Component(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="m5.212 15.111l-2.687-2.687a.6.6 0 0 1 0-.848l2.687-2.687a.6.6 0 0 1 .848 0l2.687 2.687a.6.6 0 0 1 0 .848L6.06 15.111a.6.6 0 0 1-.848 0Zm6.364 6.364l-2.687-2.687a.6.6 0 0 1 0-.849l2.687-2.687a.6.6 0 0 1 .848 0l2.687 2.687a.6.6 0 0 1 0 .849l-2.687 2.687a.6.6 0 0 1-.848 0Zm0-12.727L8.889 6.06a.6.6 0 0 1 0-.848l2.687-2.687a.6.6 0 0 1 .848 0l2.687 2.687a.6.6 0 0 1 0 .848l-2.687 2.688a.6.6 0 0 1-.848 0Zm6.364 6.363l-2.687-2.687a.6.6 0 0 1 0-.848l2.687-2.687a.6.6 0 0 1 .848 0l2.687 2.687a.6.6 0 0 1 0 .848l-2.687 2.687a.6.6 0 0 1-.848 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComponentSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-width="1.5" d="m5.212 15.111l-2.687-2.687a.6.6 0 0 1 0-.848l2.687-2.687a.6.6 0 0 1 .848 0l2.687 2.687a.6.6 0 0 1 0 .848L6.06 15.111a.6.6 0 0 1-.848 0Zm6.364 6.365l-2.687-2.687a.6.6 0 0 1 0-.849l2.687-2.687a.6.6 0 0 1 .848 0l2.687 2.687a.6.6 0 0 1 0 .848l-2.687 2.688a.6.6 0 0 1-.848 0Zm0-12.729L8.889 6.06a.6.6 0 0 1 0-.849l2.687-2.687a.6.6 0 0 1 .848 0l2.687 2.687a.6.6 0 0 1 0 .849l-2.687 2.687a.6.6 0 0 1-.848 0Zm6.364 6.364l-2.687-2.687a.6.6 0 0 1 0-.848l2.687-2.687a.6.6 0 0 1 .848 0l2.687 2.687a.6.6 0 0 1 0 .848l-2.687 2.687a.6.6 0 0 1-.848 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 12H6m6 10v-6m0 0l3 3m-3-3l-3 3m3-17v6m0 0l3-3m-3 3L9 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressLines(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 2H6m12 20H6m6-17v5m0 0l3-3m-3 3L9 7m3 12v-5m0 0l3 3m-3-3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Computer(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2 21h15m4 0h1"/><path d="M2 16.4V3.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConstrainedSurface(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 23a2 2 0 1 1 0-4a2 2 0 0 1 0 4m18 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4M3 5a2 2 0 1 1 0-4a2 2 0 0 1 0 4m18 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4m-8 7c0-1.246-1.158-2.5-2.572-2.5h-.857C8.152 9.5 7 10.62 7 12c0 1.19.855 2.185 2 2.438c.188.041.38.062.572.062"/><path d="M11 12c0 1.246 1.159 2.5 2.572 2.5h.857C15.849 14.5 17 13.38 17 12c0-1.19-.855-2.186-2-2.438a2.651 2.651 0 0 0-.572-.062M21 19V5M3 19V5m2-2h14M5 21h14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Consumable(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22.003 3v4.497A.503.503 0 0 1 21.5 8v0a.52.52 0 0 1-.466-.3A10 10 0 0 0 12.003 2c-5.185 0-9.449 3.947-9.95 9"/><path d="M17 10v5a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2m-5 1V8"/><path d="M2.05 21v-4.497c0-.278.226-.503.504-.503v0c.2 0 .38.119.466.3a10.001 10.001 0 0 0 9.03 5.7c5.186 0 9.45-3.947 9.952-9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contactless(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 21.5c4-5.5 4-13.5 0-19M11.5 20c3.5-5 3.5-11 0-16m-3 14c2.667-3.75 2.667-8.25 0-12m-3 10c1.5-2.5 1.5-5.5 0-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ControlSlider(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m6.755 17.283l-1.429-10A2 2 0 0 1 7.306 5h3.388a2 2 0 0 1 1.98 2.283l-1.429 10A2 2 0 0 1 9.265 19h-.53a2 2 0 0 1-1.98-1.717Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 12h4m16 0H12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cookie(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M7.5 11a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m5.5 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2M11 7.01l.01-.011M8 16.01l.01-.011M16 9.01l.01-.011M17 14.01l.01-.011M13 12.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoolingSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6M12 7v5m0 5v-5m0 0L7.5 9.5M12 12l4.5 2.5M12 12l4.5-2.5M12 12l-4.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.4 20H9.6a.6.6 0 0 1-.6-.6V9.6a.6.6 0 0 1 .6-.6h9.8a.6.6 0 0 1 .6.6v9.8a.6.6 0 0 1-.6.6"/><path d="M15 9V4.6a.6.6 0 0 0-.6-.6H4.6a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6H9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M13.5 9.17a3 3 0 1 0 0 5.659"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerBottomLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 8l.01.011M4 4l.01.011M8 4l.01.011M12 4l.01.011M16 4l.01.011M20 4l.01.011M20 8l.01.011M20 12l.01.011M20 16l.01.011M20 20l.01.011M16 20l.01.011M4 12.01v8h8v-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerBottomRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.01 8l-.01.011M20.01 4l-.01.011M16.01 4l-.01.011M12.01 4l-.01.011M8.01 4L8 4.011M4.01 4L4 4.011M4.01 8L4 8.011M4.01 12l-.01.011M4.01 16l-.01.011M4.01 20l-.01.011M8.01 20l-.01.011m12.01-8.001v8h-8v-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerTopLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 16.01l.01-.011M4 20.01l.01-.011M8 20.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M20 12.01l.01-.011M20 8.01l.01-.011M20 4.01l.01-.011M16 4.01l.01-.011M4 12V4h8v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerTopRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.01 16.01l-.01-.011m.01 4.011l-.01-.011m-3.99.011l-.01-.011m-3.99.011l-.01-.011m-3.99.011L8 19.999m-3.99.011L4 19.999m.01-3.989L4 15.999m.01-3.989L4 11.999m.01-3.989L4 7.999m.01-3.989L4 3.999m4.01.011L8 3.999M20.01 12V4h-8v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cpu(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 15.4V8.6a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6"/><path d="M20 4.6v14.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6M17 4V2m-5 2V2M7 4V2m0 18v2m5-2v2m5-2v2m3-5h2m-2-5h2m-2-5h2M4 17H2m2-5H2m2-5H2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CpuWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 2v4m0 4.01l.01-.011M16 14v1.4a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6V8.6a.6.6 0 0 1 .6-.6H14"/><path d="M20 14v5.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6H14m6 13h2m-5 3v2m-5-2v2m-5-2v2m-3-5H2m2-5H2m2-5H2m10-3V2M7 4V2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrackedEgg(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22a8 8 0 0 0 8-8c0-4.418-3.582-12-8-12S4 9.582 4 14a8 8 0 0 0 8 8"/><path d="M9.5 3.5L12 8l-2.5 3l2.5 3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommons(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M10.5 9.17a3 3 0 1 0 0 5.659m6.25-5.659a3 3 0 1 0 0 5.659"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 9v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2zm0 0H6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardSlash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 9h3M3 3l18 18m1-12v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h1m17 4V7a2 2 0 0 0-2-2H10m12 4h-8M9 9H6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardTwo(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 9V5.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V9Zm0 0h14"/><rect width="4" height="4" x="15" y="12" fill="currentColor" rx=".6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCards(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 11.429V18a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-1.5m17-5.071V10a2 2 0 0 0-2-2h-1m3 3.429h-3"/><path d="M19 8v6.5a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2zm0 0H5.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crib(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M3 5v16"/><path d="M3 16h18M3 7h18m-3 9V7m-4 9V7m-4 9V7m-4 9V7M3 19h18"/><path stroke-linecap="round" d="M21 5v16m0-16a1 1 0 1 0 0-2a1 1 0 0 0 0 2M3 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 18H6V3"/><path d="M3 6h15v15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropRotateBl(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 14v3a4 4 0 0 0 4 4h4"/><path d="M1.5 16.5L4 14l2.5 2.5M20 11V5a1 1 0 0 0-1-1h-6M8 4h2m10 12v-2M10 2v11a1 1 0 0 0 1 1h11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropRotateBr(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 20h3a4 4 0 0 0 4-4v-4"/><path d="M16.5 22.5L14 20l2.5-2.5M14 11V5a1 1 0 0 0-1-1H7M2 4h2m10 12v-2M4 2v11a1 1 0 0 0 1 1h11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropRotateTl(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 4H7a4 4 0 0 0-4 4v4"/><path d="M7.5 1.5L10 4L7.5 6.5M20 17v-6a1 1 0 0 0-1-1h-6m-5 0h2m10 12v-2M10 8v11a1 1 0 0 0 1 1h11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropRotateTr(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 10V7a4 4 0 0 0-4-4h-4"/><path d="M22.5 7.5L20 10l-2.5-2.5M14 17v-6a1 1 0 0 0-1-1H7m-5 0h2m10 12v-2M4 8v11a1 1 0 0 0 1 1h11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.2 17L21 7l-6.3 3L12 7l-2.7 3L3 7l1.8 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrownCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/><path d="m16.8 15.5l1.2-7l-4.2 2.1L12 8.5l-1.8 2.1L6 8.5l1.2 7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThree(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4 3l1.778 17.09L12 22l6.222-1.91L20 3z"/><path d="M7 7h9.5l-1 10l-3.5 1l-3.5-1l-.25-2.5m7.75-3H7.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cube(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524"/><path d="m3.528 7.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21v-9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeBandage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12 22l-8.691-4.828A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524V11"/><path d="m3.528 7.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 12v5.5M12 2v7m6.657 8.243l.707.707m-2.121.707l.707.707m.707-4.95l-4.243 4.243a2 2 0 0 0 0 2.828l.707.708a2 2 0 0 0 2.829 0l4.242-4.243a2 2 0 0 0 0-2.828l-.707-.708a2 2 0 0 0-2.828 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeCutWithCurve(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.528 7.293L9 10.333M22 2h-2m-8 10v-2a8.004 8.004 0 0 1 5.5-7.602M22 12h-2m-8 10v-2a8.004 8.004 0 0 1 5.5-7.602"/><path d="m12 22l-8.691-4.828A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0L15 3.667M12 12L3.528 7.293M12 21v-9m3 1.5V4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeHole(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524M3.528 7.294L8.4 10m12.1-2.722L15.6 10M12 21v-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeReplaceFace(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 13.5v5.152a.6.6 0 0 1-.302.52l-6.4 3.658a.6.6 0 0 1-.596 0l-6.4-3.657A.6.6 0 0 1 5 18.652V13m7 9.5V17m11-9L11 1m2 14L1 8m0 0c3-5 7-2 10-7"/><path d="M13 15c3-5 7-2 10-7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorPointer(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M19.503 9.97c1.204.489 1.112 2.224-.137 2.583l-6.305 1.813l-2.88 5.895c-.571 1.168-2.296.957-2.569-.314L4.677 6.257A1.369 1.369 0 0 1 6.53 4.7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurveArray(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.164 17a17.47 17.47 0 0 1 1.132-3M11.5 7.794A16.838 16.838 0 0 1 14 6.296M4.5 22a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/><path d="M9.5 12a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5m10-5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 12h1m4 0h1M6.236 7a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4m0 0L19 18M6.236 17a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4m0 0L19 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CutAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.236 8a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4Zm0 0L16 16m1-4h1m4 0h1M6.236 16a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4Zm0 0L16 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cutlery(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 20h3m3 0H9m0 0v-5m8 5v-8s2.5-1 2.5-3V4.5m-2.5 4v-4M4.5 11c1 2.128 4.5 4 4.5 4s3.5-1.872 4.5-4c1.08-2.297 0-6.5 0-6.5h-9s-1.08 4.203 0 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cycling(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4m4 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6M6 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6m5.5-3l1.5-4l-4.882-2l3-3.5l3 2.5h3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cylinder(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 2c8 0 8 3 8 3s0 3-8 3s-8-3-8-3s0-3 8-3Zm0 14c8 0 8 3 8 3s0 3-8 3s-8-3-8-3s0-3 8-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 5v14M4 5v14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashFlag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 15l.95-10.454A.6.6 0 0 1 6.548 4h13.795a.6.6 0 0 1 .598.654l-.891 9.8a.6.6 0 0 1-.598.546zm0 0l-.6 6M9 7.5l7 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M15 15.8c0-1.767-3-4.8-3-4.8s-3 3.033-3 4.8s1.343 3.2 3 3.2s3-1.433 3-3.2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 4v4m-8.5-.5l3 3m11 0l3-3M2 17h4m12 0h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardDots(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12 7.01l.01-.011M16 9.01l.01-.011M8 9.01l.01-.011M18 13.01l.01-.011M6 13.01l.01-.011M17 17.01l.01-.011M7 17.01l.01-.011M12 17l1-6m-4.5 9.001H4A9.956 9.956 0 0 1 2 14C2 8.477 6.477 4 12 4s10 4.477 10 10c0 2.252-.744 4.33-2 6.001L15.5 20"/><path d="M12 23a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardSpeed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 4v4M4 8l2.5 2.5m11 0L20 8M3 17h3m6 0l1-6m5 6h3M8.5 20.001H4A9.956 9.956 0 0 1 2 14C2 8.477 6.477 4 12 4s10 4.477 10 10c0 2.252-.744 4.33-2 6.001L15.5 20"/><path d="M12 23a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataTransferBoth(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20V4m0 0l3 3m-3-3l-3 3M7 4v16m0 0l3-3m-3 3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataTransferCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 19l3 3l5-5m-5-3V4m0 0l3 3m-3-3l-3 3M7 4v16m0 0l3-3m-3 3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataTransferDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20v-1m0-4v-1M7 4v16m0 0l-3-3m3 3l3-3m7-7V4m0 0l-3 3m3-3l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataTransferUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4v1m0 4v1m10 10V4m0 0l3 3m-3-3l-3 3m-7 7v6m0 0l3-3m-3 3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataTransferWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4v1m0 4v1m10 2V4m0 0l3 3m-3-3l-3 3m6 9v2m0 4.01l.01-.011M7 14v6m0 0l3-3m-3 3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5 12v6s0 3 7 3s7-3 7-3v-6"/><path d="M5 6v6s0 3 7 3s7-3 7-3V6"/><path d="M12 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseBackup(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3c.592 0 1.135-.021 1.631-.06M18 6v6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6m18.666 5.667C22.048 16.097 20.634 15 18.99 15c-1.758 0-3.252 1.255-3.793 3"/><path d="M20.995 17.667h1.671v0c.185 0 .334-.15.334-.334v-1.888m-7.666 4.888C15.952 21.903 17.366 23 19.01 23c1.758 0 3.252-1.255 3.793-3"/><path d="M17.005 20.333h-1.671v0a.334.334 0 0 0-.334.334v1.888"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14 19l3 3l5-5M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseCheckSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.47 18.47a.75.75 0 0 1 1.06 0L17 20.94l4.47-4.47a.75.75 0 1 1 1.06 1.06l-5 5a.75.75 0 0 1-1.06 0l-3-3a.75.75 0 0 1 0-1.06"/><path d="M15.591 17.409a2.25 2.25 0 1 0-3.182 3.182l1.021 1.02c-.713.09-1.519.139-2.43.139c-3.59 0-5.547-.767-6.613-1.68c-.536-.46-.825-.939-.977-1.33a2.478 2.478 0 0 1-.156-.648a1.377 1.377 0 0 1-.003-.055v-.02l-.001-.01v-.005S3.25 18 4 18h-.75v-2.25A.75.75 0 0 1 4 15c.414-.001.954.266 1.237.568a.63.63 0 0 0 .074.072c.69.493 2.256 1.11 5.689 1.11c3.433 0 4.999-.617 5.69-1.11c.03-.022.059-.044.087-.067c.36-.288.761-.574 1.223-.573a.75.75 0 0 1 .75.75v1.318L17 18.818z"/><path d="M4 9a.75.75 0 0 0-.75.75V12H4c-.75 0-.75.002-.75.002v.035a1.384 1.384 0 0 0 .024.215c.021.128.061.296.136.489c.152.39.441.87.977 1.329c1.066.913 3.023 1.68 6.613 1.68c3.59 0 5.547-.767 6.613-1.68c.536-.46.825-.938.977-1.33a2.472 2.472 0 0 0 .156-.648a1.36 1.36 0 0 0 .003-.055v-.02l.001-.01v-.005S18.75 12 18 12h.75V9.75A.75.75 0 0 0 18 9c-.462-.001-.863.285-1.223.573a2.012 2.012 0 0 1-.088.067c-.69.493-2.256 1.11-5.689 1.11c-3.433 0-4.999-.617-5.69-1.11a.634.634 0 0 1-.073-.071C4.954 9.266 4.414 8.999 4 9"/><path d="M4.387 3.93C5.453 3.018 7.41 2.25 11 2.25c3.59 0 5.547.767 6.613 1.68c.536.46.825.939.977 1.33a2.474 2.474 0 0 1 .156.648a1.178 1.178 0 0 1-.02.344a2.474 2.474 0 0 1-.136.489c-.152.39-.441.87-.977 1.328C16.547 8.983 14.59 9.75 11 9.75c-3.59 0-5.547-.767-6.613-1.68c-.536-.46-.825-.939-.977-1.33a2.478 2.478 0 0 1-.136-.488a1.436 1.436 0 0 1-.024-.256c0-.083.01-.166.024-.248c.021-.128.061-.295.136-.489c.152-.39.441-.87.977-1.328"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseExport(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 16v6m0 0l3-3m-3 3l-3-3M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseMonitor(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6m14 10h3m-1.5-2.571h2.333V16h-4.666v3.429z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseRestore(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6m15 10v-6m0 0l3 3m-3-3l-3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseScript(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 14V6a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v7"/><path d="M12 21H6a4 4 0 0 1 0-8h12a4 4 0 1 0 4 4v-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseScriptMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 14V8.5M6 13V6a3 3 0 0 1 3-3h5m4 1h4M12 21H6a4 4 0 0 1 0-8h12a4 4 0 1 0 4 4v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseScriptPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 14V8.5M6 13V6a3 3 0 0 1 3-3h5m2.992 1h3m3 0h-3m0 0V1m0 3v3M12 21H6a4 4 0 0 1 0-8h12a4 4 0 1 0 4 4v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseSearch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 20.5L22 22m-6-3.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseSettings(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6m15 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path stroke-dasharray=".3 2" d="M19 22a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 15a.75.75 0 0 0-.75.75V18H5c-.75 0-.75.002-.75.002v.035a1.384 1.384 0 0 0 .024.215c.021.128.061.296.136.489c.152.39.441.87.977 1.329c1.066.913 3.023 1.68 6.613 1.68c3.59 0 5.547-.767 6.613-1.68c.536-.46.825-.939.977-1.33a2.472 2.472 0 0 0 .156-.648a1.36 1.36 0 0 0 .003-.055v-.02l.001-.01v-.005S19.75 18 19 18h.75v-2.25A.75.75 0 0 0 19 15c-.462-.001-.863.285-1.223.573a1.965 1.965 0 0 1-.088.067c-.69.493-2.256 1.11-5.689 1.11c-3.433 0-4.999-.617-5.69-1.11a.635.635 0 0 1-.073-.071c-.283-.303-.823-.57-1.237-.569"/><path d="M5 9a.75.75 0 0 0-.75.75V12H5c-.75 0-.75.002-.75.002v.035a1.384 1.384 0 0 0 .024.215c.021.128.061.296.136.489c.152.39.441.87.977 1.329c1.066.913 3.023 1.68 6.613 1.68c3.59 0 5.547-.767 6.613-1.68c.536-.46.825-.938.977-1.33a2.472 2.472 0 0 0 .156-.648a1.36 1.36 0 0 0 .003-.055v-.02l.001-.01v-.005S19.75 12 19 12h.75V9.75A.75.75 0 0 0 19 9c-.462-.001-.863.285-1.223.573a2.012 2.012 0 0 1-.088.067c-.69.493-2.256 1.11-5.689 1.11c-3.433 0-4.999-.617-5.69-1.11a.634.634 0 0 1-.073-.071C5.954 9.266 5.414 8.999 5 9"/><path d="M5.387 3.93C6.453 3.018 8.41 2.25 12 2.25c3.59 0 5.547.767 6.613 1.68c.536.46.825.939.977 1.33a2.474 2.474 0 0 1 .156.648a1.178 1.178 0 0 1-.02.344a2.474 2.474 0 0 1-.136.489c-.152.39-.441.87-.977 1.328C17.547 8.983 15.59 9.75 12 9.75c-3.59 0-5.547-.767-6.613-1.68c-.536-.46-.825-.939-.977-1.33a2.478 2.478 0 0 1-.136-.488a1.436 1.436 0 0 1-.024-.256c0-.083.01-.166.024-.248c.021-.128.061-.295.136-.489c.152-.39.441-.87.977-1.328"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseStar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6m13.306 5.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.233.129-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseStats(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4 6v6s0 3 7 3s7-3 7-3V6"/><path stroke-linejoin="round" d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6"/><path d="M15 21v-2m3 2v-4m3 4v-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.357 12c.714 0 2.143 0 2.143-2s-1.429-2-2.143-2H13.5v4m2.857 0H13.5m2.857 0c.714 0 2.143 0 2.143 2s-1.429 2-2.143 2H13.5v-4M8.357 8H5.5v8h2.857c.714 0 2.143 0 2.143-2v-4c0-2-1.429-2-2.143-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseTagSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 9A6.75 6.75 0 0 1 8 2.25h8A6.75 6.75 0 0 1 22.75 9v6A6.75 6.75 0 0 1 16 21.75H8A6.75 6.75 0 0 1 1.25 15zm16.08-.096c-.25-.14-.588-.154-.973-.154H14.25v2.5h2.128c.377 0 .706-.017.952-.154c.158-.089.42-.302.42-1.096s-.262-1.007-.42-1.096M18.582 12c.423-.45.668-1.112.668-2c0-1.206-.452-1.993-1.187-2.404c-.62-.348-1.325-.346-1.67-.346H13.5a.75.75 0 0 0-.75.75v8c0 .414.336.75.75.75h2.894c.344 0 1.049.002 1.669-.346c.735-.411 1.187-1.198 1.187-2.404c0-.888-.245-1.549-.668-2m-2.203.75H14.25v2.5h2.107c.385 0 .723-.014.973-.154c.158-.089.42-.302.42-1.096s-.262-1.007-.42-1.095c-.245-.138-.575-.154-.951-.155m-8.022-4c.385 0 .723.014.973.154c.158.089.42.302.42 1.096v4c0 .794-.262 1.007-.42 1.096c-.25.14-.588.154-.973.154H6.25v-6.5zM11.25 10c0-1.206-.452-1.993-1.187-2.404c-.62-.348-1.325-.346-1.67-.346H5.5a.75.75 0 0 0-.75.75v8c0 .414.336.75.75.75h2.894c.344 0 1.049.002 1.669-.346c.735-.411 1.187-1.198 1.187-2.404z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 16v2m0 4.01l.01-.011M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.121 21.364l2.122-2.121m2.121-2.122l-2.121 2.122m0 0L17.12 17.12m2.122 2.122l2.121 2.121M4 6v6s0 3 7 3s7-3 7-3V6"/><path d="M11 3c7 0 7 3 7 3s0 3-7 3s-7-3-7-3s0-3 7-3m0 18c-7 0-7-3-7-3v-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseXmarkSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M15.008 16.346c-.916.234-2.204.404-4.008.404c-3.433 0-4.999-.617-5.69-1.11a.63.63 0 0 1-.073-.072c-.283-.302-.823-.57-1.237-.568a.75.75 0 0 0-.75.75V18H4c-.75 0-.75.002-.75.002v.035a1.377 1.377 0 0 0 .024.215c.021.128.061.295.136.489c.152.39.441.87.977 1.328C5.453 20.983 7.41 21.75 11 21.75c1.589 0 2.858-.15 3.871-.397a2.243 2.243 0 0 1 .66-1.58l.53-.53l-.53-.531a2.25 2.25 0 0 1-.523-2.366" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M4 9a.75.75 0 0 0-.75.75V12H4c-.75 0-.75.002-.75.002v.035a1.384 1.384 0 0 0 .024.215c.021.128.061.296.136.489c.152.39.441.87.977 1.329c1.066.913 3.023 1.68 6.613 1.68c3.59 0 5.547-.767 6.613-1.68c.536-.46.825-.938.977-1.33a2.472 2.472 0 0 0 .156-.648a1.36 1.36 0 0 0 .003-.055v-.02l.001-.01v-.005S18.75 12 18 12h.75V9.75A.75.75 0 0 0 18 9c-.462-.001-.863.285-1.223.573a2.012 2.012 0 0 1-.088.067c-.69.493-2.256 1.11-5.689 1.11c-3.433 0-4.999-.617-5.69-1.11a.634.634 0 0 1-.073-.071C4.954 9.266 4.414 8.999 4 9" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M4.387 3.93C5.453 3.018 7.41 2.25 11 2.25c3.59 0 5.547.767 6.613 1.68c.536.46.825.939.977 1.33a2.474 2.474 0 0 1 .156.648a1.178 1.178 0 0 1-.02.344a2.474 2.474 0 0 1-.136.489c-.152.39-.441.87-.977 1.328C16.547 8.983 14.59 9.75 11 9.75c-3.59 0-5.547-.767-6.613-1.68c-.536-.46-.825-.939-.977-1.33a2.478 2.478 0 0 1-.136-.488a1.436 1.436 0 0 1-.024-.256c0-.083.01-.166.024-.248c.021-.128.061-.295.136-.489c.152-.39.441-.87.977-1.328" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17.121 21.364l2.122-2.122m2.121-2.121l-2.121 2.121m0 0l-2.122-2.121m2.122 2.121l2.121 2.121"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DbStar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6s0-3 7-3s7 3 7 3M4 6s0 3 7 3s7-3 7-3M4 6v6m14-6v6s0 3-7 3s-7-3-7-3m7 9c-7 0-7-3-7-3v-6m13.306 5.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.233.129-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568l2.033-.31Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeCompress(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 12H6m6 4v6m0 0l3-3m-3 3l-3-3m3-11V2m0 0l3 3m-3-3L9 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delivery(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-6m8-3V4M8 8H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeliveryTruck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M8 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4m10 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M10.05 17H15V6.6a.6.6 0 0 0-.6-.6H1m4.65 11H3.6a.6.6 0 0 1-.6-.6v-4.9"/><path stroke-linejoin="round" d="M2 9h4"/><path d="M15 9h5.61a.6.6 0 0 1 .548.356l1.79 4.028a.6.6 0 0 1 .052.243V16.4a.6.6 0 0 1-.6.6h-1.9M15 17h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Depth(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 20h20M5 4h14M3 16.01l.01-.011m18 .011l-.01-.011M4 12.01l.01-.011m16 .011l-.01-.011M5 8.01l.01-.011m14 .011L19 7.999M12 7v10m0-10l-1.5 1.5M12 7l1.5 1.5M12 17l-3-3m3 3l3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesignNib(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.674 11.408l-1.905 5.715a.6.6 0 0 1-.398.386L3.693 20.98a.6.6 0 0 1-.74-.765L6.745 8.841a.6.6 0 0 1 .34-.365l5.387-2.218a.6.6 0 0 1 .653.13l4.404 4.406a.6.6 0 0 1 .145.614M3.296 20.602l6.364-6.364"/><path d="m17.792 11.056l2.828-2.829a2 2 0 0 0 0-2.828L18.5 3.277a2 2 0 0 0-2.829 0l-2.828 2.829m-1.062 6.01a1.5 1.5 0 1 0-2.121 2.122a1.5 1.5 0 0 0 2.121-2.122"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesignNibSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.674 11.408l-1.905 5.716a.6.6 0 0 1-.398.385L3.693 20.981a.6.6 0 0 1-.74-.765L6.745 8.842a.6.6 0 0 1 .34-.365l5.387-2.218a.6.6 0 0 1 .653.13l4.404 4.405a.6.6 0 0 1 .145.614M3.296 20.602l6.364-6.364"/><path fill="currentColor" d="m18.403 3.182l2.364 2.364a1.846 1.846 0 1 1-2.61 2.61l-2.365-2.364a1.846 1.846 0 0 1 2.61-2.61"/><path d="M11.781 12.116a1.5 1.5 0 1 0-2.121 2.121a1.5 1.5 0 0 0 2.121-2.121"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesignPencil(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2"/><path d="M8 21.168V14l4-7l4 7v7.168"/><path d="M8 14s1.127 1 2 1s2-1 2-1s1.127 1 2 1s2-1 2-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desk(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 7v10M1 7h22M4 10h16m-6 4h6m0-7v10M14 7v10m3-7v1m0 3v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Developer(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.818 22v-2.857C6.662 17.592 5.633 16.416 4.682 15m9.772 7v-1.714c4.91 0 4.364-5.714 4.364-5.714s2.182 0 2.182-2.286l-2.182-3.428c0-4.572-3.709-6.816-7.636-6.857c-2.2-.023-3.957.53-5.27 1.499"/><path d="m13 7l2 2.5l-2 2.5M5 7L3 9.5L5 12m5-6l-2 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DewPoint(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 12a5 5 0 1 0 6 0m-6 0V3h6v9m0-9h2m-2 3h2m-2 3h2"/><path d="M8 14a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 0V9"/><path stroke-miterlimit="1.5" d="M19 3s3 2.993 3 4.887c0 1.655-1.345 3-3 3c-1.656 0-2.988-1.345-3-3C16.01 5.992 19 3 19 3" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dialpad(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.5 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M12 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m6.5-15a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diameter(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10m7-10l-3-3m3 3l-3 3m3-3H5m0 0l3-3m-3 3l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFive(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M12 12.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M7.5 17a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFour(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-9 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceOne(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 12.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceSix(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-9 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-9 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m9 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceThree(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4.5 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4.5 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceTwo(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m9 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DimmerSwitch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12m3.5-10.5L12 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectorChair(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M19 12L5 21M5 3v9m14-9v9M5 12l14 9M4 12h16"/><path d="M5 4h14M5 7h14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discord(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.5 16c5 2.5 8 2.5 13 0"/><path d="m15.5 17.5l1 2s4.171-1.328 5.5-3.5c0-1 .53-8.147-3-10.5c-1.5-1-4-1.5-4-1.5l-1 2h-2"/><path d="m8.528 17.5l-1 2s-4.171-1.328-5.5-3.5c0-1-.53-8.147 3-10.5c1.5-1 4-1.5 4-1.5l1 2h2"/><path d="M8.5 14c-.828 0-1.5-.895-1.5-2s.672-2 1.5-2s1.5.895 1.5 2s-.672 2-1.5 2m7 0c-.828 0-1.5-.895-1.5-2s.672-2 1.5-2s1.5.895 1.5 2s-.672 2-1.5 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dishwasher(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-2.5M21 7H3m9 9v5m0 0h-2m2 0h2"/><path d="M12 16c1.657 0 3-1.492 3-3.333V10H9v2.667C9 14.507 10.343 16 12 16m6-10.99l.01-.011M15 5.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DisplayFourK(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M13.5 9v4m0 2v-2m0 0l1.37-1.566M17 9l-2.13 2.434m0 0L17 15M9.5 9l-3 4.5H10V15"/><path d="M2 18.4V5.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Divide(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21v-9a5 5 0 0 0-5-5H3m9 14v-9a5 5 0 0 1 5-5h4"/><path d="M7 3L3 7l4 4m10-8l4 4l-4 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DivideThree(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21v-4a5 5 0 0 0-5-5H3m9 9v-4a5 5 0 0 1 5-5h4M12 2v20"/><path d="m6 8l-4 4l4 4M16 6l-4-4l-4 4m10 2l4 4l-4 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dna(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 3c0 5.625 8 9 8 9s8 3.375 8 9"/><path d="M20 3c0 5.625-8 9-8 9s-8 3.375-8 9M8 6h11M8 18h11m-8-9h5.5M11 15h5.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dns(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 13v-1c0-5.523-4.477-10-10-10S2 6.477 2 12c0 2.251.744 4.329 2 6"/><path d="M13 2.049s3 3.95 3 9.95v1m-5-10.95s-3 3.95-3 9.95v1M2.63 15.5H4m-1.37-7h18.74M7 22v-6h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2zm6 0v-6l3 6v-6m3 6h2a1.5 1.5 0 0 0 1.5-1.5v0A1.5 1.5 0 0 0 21 19h-.5a1.5 1.5 0 0 1-1.5-1.5v0a1.5 1.5 0 0 1 1.5-1.5H22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocMagnifyingGlass(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 20.5L22 22m-7-4a3 3 0 1 0 6 0a3 3 0 0 0-6 0"/><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocMagnifyingGlassIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14 15l1.5 1.5m-7-4a3 3 0 1 0 6 0a3 3 0 0 0-6 0"/><path d="M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocStar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m16.306 17.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.234.129-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568z"/><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocStarIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6"/><path d="m10.635 10.415l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L12 14.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.217-.22-.098-.604.2-.65zM16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DogecoinCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M10 16.402V7.598c0-.331.268-.599.6-.604c2.49-.035 5.9-.072 5.9 5.006s-3.41 5.042-5.9 5.006a.606.606 0 0 1-.6-.604Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 12h4m0 10C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DogecoinCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12m15.348-3.268c.426.845.652 1.919.652 3.268c0 1.35-.226 2.424-.652 3.268a4.102 4.102 0 0 1-1.761 1.814c-1.341.716-2.988.692-4.184.675l-.065-.001a1.356 1.356 0 0 1-1.338-1.354V12.75H8a.75.75 0 0 1 0-1.5h1.25V7.598c0-.747.603-1.343 1.338-1.354h.065c1.196-.018 2.843-.042 4.184.674c.71.38 1.329.959 1.76 1.814M10.75 11.25V7.742c1.242-.015 2.453.004 3.38.5c.454.241.844.602 1.129 1.167c.29.576.491 1.401.491 2.591c0 1.19-.2 2.015-.491 2.592a2.603 2.603 0 0 1-1.129 1.167c-.927.495-2.138.515-3.38.5V12.75H12a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DogecoinRotateOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21.168 8A10.003 10.003 0 0 0 12 2c-5.185 0-9.45 3.947-9.95 9"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 8h3.4a.6.6 0 0 0 .6-.6V4M2.881 16c1.544 3.532 5.068 6 9.168 6c5.186 0 9.45-3.947 9.951-9"/><path stroke-linecap="round" stroke-linejoin="round" d="M6.05 16h-3.4a.6.6 0 0 0-.6.6V20"/><path d="M10 16.402V7.598c0-.331.268-.599.6-.604c2.49-.035 5.9-.072 5.9 5.006s-3.41 5.042-5.9 5.006a.606.606 0 0 1-.6-.604Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 12h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.154 7.154c-.949-.949-2.619-1.608-4.154-1.65m-4.154 10.65c.893 1.19 2.552 1.868 4.154 1.926m0-12.576c-1.826-.049-3.461.778-3.461 3.034c0 4.154 7.615 2.077 7.615 6.231c0 2.37-2.027 3.387-4.154 3.31m0-12.575V3m0 15.08V21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M15 8.5c-.685-.685-1.891-1.161-3-1.191M9 15c.644.86 1.843 1.35 3 1.391m0-9.082c-1.32-.036-2.5.561-2.5 2.191c0 3 5.5 1.5 5.5 4.5c0 1.711-1.464 2.446-3 2.391m0-9.082V5.5m0 10.891V18.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DomoticWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v3m0 4.01l.01-.011M2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8m-2 3v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Donate(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M16 6.28a2.28 2.28 0 0 1-.662 1.606c-.976.984-1.923 2.01-2.936 2.958a.597.597 0 0 1-.822-.017l-2.918-2.94a2.281 2.281 0 0 1 0-3.214a2.277 2.277 0 0 1 3.232 0L12 4.78l.106-.107A2.276 2.276 0 0 1 16 6.28Z"/><path stroke-linecap="round" d="m18 20l3.824-3.824a.6.6 0 0 0 .176-.424V10.5A1.5 1.5 0 0 0 20.5 9v0a1.5 1.5 0 0 0-1.5 1.5V15"/><path stroke-linecap="round" d="m18 16l.858-.858a.484.484 0 0 0 .142-.343v0a.485.485 0 0 0-.268-.433l-.443-.221a2 2 0 0 0-2.308.374l-.895.895a2 2 0 0 0-.586 1.414V20M6 20l-3.824-3.824A.6.6 0 0 1 2 15.752V10.5A1.5 1.5 0 0 1 3.5 9v0A1.5 1.5 0 0 1 5 10.5V15"/><path stroke-linecap="round" d="m6 16l-.858-.858A.485.485 0 0 1 5 14.799v0c0-.183.104-.35.268-.433l.443-.221a2 2 0 0 1 2.308.374l.895.895a2 2 0 0 1 .586 1.414V20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotArrowDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6a2 2 0 1 1 0-4a2 2 0 0 1 0 4m0 3v13m0 0l3-3m-3 3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotArrowLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4m-5-2H2m0 0l3-3m-3 3l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotArrowRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4m5-2h13m0 0l-3-3m3 3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotArrowUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22a2 2 0 1 1 0-4a2 2 0 0 1 0 4m0-7V2m0 0l3 3m-3-3L9 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m1.5 12.5l4.076 4.076a.6.6 0 0 0 .848 0L9 14m7-7l-4 4"/><path d="m7 12l4.576 4.576a.6.6 0 0 0 .848 0L22 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 20h12M12 4v12m0 0l3.5-3.5M12 16l-3.5-3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 17h6M12 6v7m0 0l3.5-3.5M12 13L8.5 9.5M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12m7 5a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5H9a.75.75 0 0 1-.75-.75m7.78-6.97l-3.5 3.5a.75.75 0 0 1-1.06 0l-3.5-3.5a.75.75 0 1 1 1.06-1.06l2.22 2.22V6a.75.75 0 0 1 1.5 0v5.19l2.22-2.22a.75.75 0 1 1 1.06 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadDataWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v9"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M19.5 16v6m0 0L17 19.5m2.5 2.5l2.5-2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18h12M12 6v8m0 0l3.5-3.5M12 14l-3.5-3.5"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zm3 14.4a.75.75 0 0 1 .75-.75h12a.75.75 0 0 1 0 1.5H6a.75.75 0 0 1-.75-.75m10.78-6.97l-3.5 3.5a.75.75 0 0 1-1.06 0l-3.5-3.5a.75.75 0 1 1 1.06-1.06l2.22 2.22V6a.75.75 0 0 1 1.5 0v6.19l2.22-2.22a.75.75 0 1 1 1.06 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12L4 4m0 0v4m0-4h4m4 8l8-8m0 0v4m0-4h-4m-4 8l-8 8m0 0v-4m0 4h4m4-8l8 8m0 0v-4m0 4h-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragHandGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7 10.5l-2.004 2.672a2 2 0 0 0 .126 2.552l3.784 4.128c.378.413.912.648 1.473.648H15c2.4 0 4-1.5 4-4c0 0 0 0 0 0V7.929M16 8.5v-.571c0-2.286 3-2.286 3 0"/><path d="M13 8.5V7.027m0-.527v.527M16 8.5V7.027c0-2.286-3-2.286-3 0"/><path d="M13 8.5V7.027c0-2.286 3-2.286 3 0V8.5m-6 0v-2c0-2.286 3-2.286 3 0c0 0 0 0 0 0v2m-6 5v-7A1.5 1.5 0 0 1 8.5 5v0c.828 0 1.5.555 1.5 1.384V8.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drawer(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 14H3m0-6h18m-10 9h2m-2-6h2m-2-6h2m8-2.4v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V2.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6M17.5 20v2m-11-2v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12"/><path d="M16.673 20.844C15.5 14 12.5 8 8.5 2.63"/><path d="M2.067 10.84C6 11 15.283 10.5 19.142 5m2.826 7.81C15.344 10.84 7.5 14 5.23 19.361"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drone(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m4.5 11.5l-5-4"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m0 15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DroneChargeFull(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="m14.25 14.75l.25-2.25l.426-3.834A.6.6 0 0 0 14.33 8H9.67a.6.6 0 0 0-.596.666l.867 7.8a.6.6 0 0 0 .596.534H11"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 19v2m-8-2v2m2-2v2m2-2v2"/><path d="M13 22.4v-4.8a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-6.8a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DroneChargeHalf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="m14.25 14.75l.25-2.25l.426-3.834A.6.6 0 0 0 14.33 8H9.67a.6.6 0 0 0-.596.666l.867 7.8a.6.6 0 0 0 .596.534H11"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 19v2m-8-2v2m2-2v2"/><path d="M13 22.4v-4.8a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-6.8a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DroneChargeLow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="m14.25 14.75l.25-2.25l.426-3.834A.6.6 0 0 0 14.33 8H9.67a.6.6 0 0 0-.596.666l.867 7.8a.6.6 0 0 0 .596.534H11"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linecap="round" stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 19v2m-8-2v2"/><path d="M13 22.4v-4.8a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-6.8a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DroneCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m-.5 7.5l1.25 1"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="m16 20l2 2l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DroneLanding(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m-.5 7.5l1.25 1"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="M19.5 16v6m0 0L17 19.5m2.5 2.5l2.5-2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DroneRefresh(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m14.5 12.5l.426-3.834A.6.6 0 0 0 14.33 8H9.67a.6.6 0 0 0-.596.666l.867 7.8a.6.6 0 0 0 .596.534H11"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="M21.666 16.667C21.048 15.097 19.634 14 17.99 14c-1.758 0-3.252 1.255-3.793 3"/><path stroke-linejoin="round" d="M19.995 16.772H21.4a.6.6 0 0 0 .6-.6V14.55m-7.666 4.783C14.952 20.903 16.366 22 18.01 22c1.758 0 3.252-1.255 3.793-3"/><path stroke-linejoin="round" d="M16.005 19.228H14.6a.6.6 0 0 0-.6.6v1.622"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DroneTakeOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m-.5 7.5l1.25 1"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="M19.5 22v-6m0 0L17 18.5m2.5-2.5l2.5 2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DroneXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13.463 17h-2.926a.6.6 0 0 1-.596-.534l-.867-7.8A.6.6 0 0 1 9.67 8h4.66a.6.6 0 0 1 .596.666l-.867 7.8a.6.6 0 0 1-.596.534Z"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="M4.5 4.5L9 8M4.5 19.5l5-4m10-11L15 8m-.5 7.5l1.25 1"/><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M4.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m15-15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/><path stroke-linejoin="round" d="m18 22.243l2.121-2.122m0 0L22.243 18m-2.122 2.121L18 18m2.121 2.121l2.122 2.122"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Droplet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M20 14c0-4.418-8-12-8-12S4 9.582 4 14a8 8 0 1 0 16 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropletCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 21.57A8.132 8.132 0 0 1 6.25 7.75l5.326-5.326a.6.6 0 0 1 .848 0L17.75 7.75A8.131 8.131 0 0 1 19.74 16M16 20l2 2l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropletHalf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="m4.5 16.5l14-6.5m1.5 4c0-4.418-8-12-8-12S4 9.582 4 14a8 8 0 1 0 16 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropletSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-width="1.5" d="M20 14c0-4.418-8-12-8-12S4 9.582 4 14a8 8 0 1 0 16 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EaseCurveControlPoints(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20a2 2 0 1 0 4 0a2 2 0 0 0-4 0m0 0h-2M7 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0m0 0h2m5 0h-2m0 16h-2m-7 0c8 0 10-16 18-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EaseIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20c8 0 18-16 18-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EaseInControlPoint(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20c8 0 18-16 18-16m-4 16a2 2 0 1 0 4 0a2 2 0 0 0-4 0m0 0h-2m-3 0h-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EaseInOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20c8 0 10-16 18-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EaseOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20S13 4 21 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EaseOutControlPoint(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20S13 4 21 4M7 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0m0 0h2m5 0h-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EcologyBook(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M10 14s.9-3.118 3-5"/><path stroke-linejoin="round" d="m12.802 12.425l-.134.012a3.094 3.094 0 0 1-3.366-2.774a3.06 3.06 0 0 1 2.761-3.35l2.986-.28a.35.35 0 0 1 .381.314l.255 2.58a3.194 3.194 0 0 1-2.883 3.498"/><path d="M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21h18M12.222 5.828L15.05 3L20 7.95l-2.828 2.828m-4.95-4.95l-5.607 5.607a1 1 0 0 0-.293.707v4.536h4.536a1 1 0 0 0 .707-.293l5.607-5.607m-4.95-4.95l4.95 4.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditPencil(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.363 5.652l1.48-1.48a2 2 0 0 1 2.829 0l1.414 1.414a2 2 0 0 1 0 2.828l-1.48 1.48m-4.243-4.242l-9.616 9.615a2 2 0 0 0-.578 1.238l-.242 2.74a1 1 0 0 0 1.084 1.085l2.74-.242a2 2 0 0 0 1.24-.578l9.615-9.616m-4.243-4.242l4.243 4.242"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Egg(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22a8 8 0 0 0 8-8c0-4.418-3.582-12-8-12S4 9.582 4 14a8 8 0 0 0 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5 14l-.592-.46A.75.75 0 0 0 5 14.75zm14 0v.75a.75.75 0 0 0 .592-1.21zm-14 .75h14v-1.5H5zm5.619-9.196L4.408 13.54l1.184.92l6.21-7.985zm8.973 7.986l-6.21-7.986l-1.185.921l6.211 7.986zm-7.79-7.065a.25.25 0 0 1 .395 0l1.184-.92a1.749 1.749 0 0 0-2.762 0zM5 17.25a.75.75 0 0 0 0 1.5zm14 1.5a.75.75 0 0 0 0-1.5zm-14 0h14v-1.5H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectronicsChip(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 19.4V4.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6m7 .6v2.5M10 20v2.5M14 4V1.5M10 4V1.5M7 12H4.5m15 0H17M7 6.5H4.5m15 0H17m-10 11H4.5m15 0H17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectronicsTransistor(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16V3.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6V16M7 16h2m-2 0H5m12 0h-2m2 0h2m-7 0v6m0-6H9m3 0h3m-6 0v6m6-6v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Elevator(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 3v18m9-17.4v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="m6 12l1.5-2L9 12m6 0l1.5 2l1.5-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipseThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M12 22c4.418 0 8-4.477 8-10S16.418 2 12 2S4 6.477 4 12s3.582 10 8 10"/><path fill="currentColor" d="M12 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipseThreeDthreePoints(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M5 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M5 22h8m-8 0V2"/><path fill="currentColor" d="M5 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path stroke-dasharray="3 3" d="M8 4.193C9.37 2.821 11.108 2 13 2c4.418 0 8 4.477 8 10c0 3.271-1.256 6.176-3.2 8"/><path d="M8.2 20A9.098 9.098 0 0 1 7 18.615"/><path fill="currentColor" d="M13 23a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Emoji(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/><path d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2"/><path fill="currentColor" d="M15.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiBall(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9H8m8 0h-2m4 6H6m-4-3c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiBlinkLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 9H8m-6 3c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12"/><path d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2"/><path fill="currentColor" d="M15.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiBlinkRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/><path d="M14 9h2m6 3c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10"/><path d="M7.5 14.5s1.5 2 4.5 2s4.5-2 4.5-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiLookDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 14a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/><path d="M10 18h4m8-6c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiLookLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/><path d="M2.458 15A9.996 9.996 0 0 1 2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10c-4.478 0-8.268-2.943-9.542-7m0 0H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiLookRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M15.5 9a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/><path d="M21.542 15A9.997 9.997 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10c4.478 0 8.268-2.943 9.542-7m0 0H17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiLookUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 7a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/><path d="M10 11h4m8 1c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiPuzzled(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2"/><path d="M11.5 15.5s1.5-2 4.5-2s4.5 2 4.5 2M3 4c0-2.754 4-2.754 4 0c0 1.967-2 1.64-2 4m0 3.01l.01-.011"/><path fill="currentColor" d="M17.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m-7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiQuite(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 15h6m7-3c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiReally(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9H8m8 0h-2m1 6H9m-7-3c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSad(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M7.5 15.5s1.5-2 4.5-2s4.5 2 4.5 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSatisfied(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 9H8m8 0h-2M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12"/><path d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSingLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSingLeftNote(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke-width="1.5"><path fill="currentColor" d="M2.8 8.1a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0"/><path stroke="currentColor" stroke-linecap="round" d="M2.8 8.1a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm0 0V3.6a.6.6 0 0 1 .6-.6H5"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 17a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.05 13c.501 5.053 4.765 9 9.95 9c5.523 0 10-4.477 10-10S17.523 2 12 2a9.966 9.966 0 0 0-4 .832"/><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 9a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-7 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSingRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="iconoirEmojiSingRight0" fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></defs><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><use href="#iconoirEmojiSingRight0"/><use href="#iconoirEmojiSingRight0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSingRightNote(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke-width="1.5"><path fill="currentColor" d="M20.8 8.1a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0"/><path stroke="currentColor" stroke-linecap="round" d="M20.8 8.1a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm0 0V3.6a.6.6 0 0 1 .6-.6H23"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.95 13c-.501 5.053-4.765 9-9.95 9c-5.523 0-10-4.477-10-10S6.477 2 12 2a9.97 9.97 0 0 1 4 .832"/><path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSurprise(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m3.5 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3.5-8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiSurpriseAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiTalkingAngry(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9H8m8 0h-2M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12m12 6h-4v-3c0-.667.4-2 2-2s2 1.333 2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiTalkingHappy(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9H8m8 0h-2M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12m12 1h-4v3c0 .667.4 2 2 2s2-1.333 2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiThinkLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 15H7m-5-3c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmojiThinkRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 15h3m5-3c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10"/><path fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmptyPage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnergyUsageWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 19V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M11.667 11L10 14h4l-1.667 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enlarge(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 9l5-5m0 0v4m0-4h-4M9 15l-5 5m0 0v-4m0 4h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Erase(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21H9m6.889-6.11L8.464 7.463m-5.571 5.144l9.193-9.193a2 2 0 0 1 2.828 0l4.95 4.95a2 2 0 0 1 0 2.828l-9.243 9.243a1.929 1.929 0 0 1-2.728 0l-5-5a2 2 0 0 1 0-2.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EraseSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 21H9"/><path fill="currentColor" d="m14.914 3.414l4.95 4.95a2 2 0 0 1 0 2.828l-9.243 9.243a1.929 1.929 0 0 1-2.728 0l-5-5a2 2 0 0 1 0-2.828L7 8.5l4.75 4.75a1.768 1.768 0 1 0 2.5-2.5L9.5 6l2.586-2.586a2 2 0 0 1 2.828 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EthereumCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7 12l5 7l5-7M7 12l5-7m-5 7l5 1m0-8l5 7m-5-7v8m5-1l-5 1"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EthereumCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m.61 3.314a.75.75 0 0 0-1.22 0l-5 7a.75.75 0 0 0 0 .872l5 7a.75.75 0 0 0 1.22 0l5-7a.75.75 0 0 0 0-.872zM12 17.71l-3.287-4.603l3.14.628c.097.02.197.02.294 0l3.14-.628zm.75-5.625l2.966-.593L12.75 7.34zm-1.5 0V7.34l-2.966 4.152z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EthereumRotateOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.168 8A10.003 10.003 0 0 0 12 2c-5.185 0-9.45 3.947-9.95 9"/><path d="M18 8h3.4a.6.6 0 0 0 .6-.6V4M2.881 16c1.544 3.532 5.068 6 9.168 6c5.186 0 9.45-3.947 9.951-9"/><path d="M6.05 16h-3.4a.6.6 0 0 0-.6.6V20M7 12l5 7l5-7M7 12l5-7m-5 7l5 1m0-8l5 7m-5-7v8m5-1l-5 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.5 4.804a8 8 0 1 0 0 14.392M5 10h11M5 14h11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 7.503A4.746 4.746 0 0 0 13.87 7C11.18 7 9 9.239 9 12s2.18 5 4.87 5a4.73 4.73 0 0 0 2.13-.503M8 11h6m-6 2h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EvCharge(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M6 9v10a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-5M9 5.6V7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6ZM4 5V3m4 2V3"/><path stroke-linejoin="round" d="M18.167 4L16.5 7h4l-1.667 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EvChargeAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m22 5l-2 4l-2-4m-2 0h-2v4h2m-2-2h1.5"/><path d="M6 9v10a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-5M9 5.6V7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6ZM4 5V3m4 2V3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EvPlug(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 13.154V21m5-12.615v2.769a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-2.77a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2Zm-1.667-2V3M8.667 6.385V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EvPlugCharging(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M10 13.154V21m5-12.615v2.769a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-2.77a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2Zm-1.667-2V3M6.667 6.385V3"/><path stroke-linejoin="round" d="M16.667 16L15 19h4l-1.667 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EvPlugXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M10 13.154V21m5-12.615v2.769a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-2.77a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2Zm-1.667-2V3M6.667 6.385V3"/><path stroke-linejoin="round" d="m15.121 21.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L15.12 17.12m2.122 2.122l2.121 2.121"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EvStation(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 5v4"/><path d="M5 19V9a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2Z"/><path d="M5 10V5a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v5"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.167 11L9.5 14h4l-1.667 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EvTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m18.5 9l-3 6l-3-6M10 9H6v6h4m-4-3h3"/><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exclude(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 15h-.9a.6.6 0 0 0-.6.6v4.8a.6.6 0 0 0 .6.6h10.8a.6.6 0 0 0 .6-.6V9.6a.6.6 0 0 0-.6-.6h-4.8a.6.6 0 0 0-.6.6v.9"/><path d="M13.5 15h.9a.6.6 0 0 0 .6-.6v-.9m-6 0v.9a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-.9"/><path d="M9 10.5v-.9a.6.6 0 0 1 .6-.6h.9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 9L4 4m0 0v4m0-4h4m7 5l5-5m0 0v4m0-4h-4M9 15l-5 5m0 0v-4m0 4h4m7-5l5 5m0 0v-4m0 4h-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandLines(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 2H6m12 20H6m6-8v5m0 0l3-3m-3 3l-3-3m3-6V5m0 0l3 3m-3-3L9 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Extrude(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 12.353v4.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647v-4.294a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524"/><path d="m3.528 12.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21.5V17m0-5V2m0 0l2.5 2.5M12 2L9.5 4.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 13c3.6-8 14.4-8 18 0"/><path d="M12 17a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.5 16l-2.475-3.396M12 17.5V14m-7.5 2l2.469-3.388M3 8c3.6 8 14.4 8 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeEmpty(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M21 12c-1.889 2.991-5.282 6-9 6s-7.111-3.009-9-6c2.299-2.842 4.992-6 9-6s6.701 3.158 9 6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 3l18 18M10.5 10.677a2 2 0 0 0 2.823 2.823"/><path d="M7.362 7.561C5.68 8.74 4.279 10.42 3 12c1.889 2.991 5.282 6 9 6c1.55 0 3.043-.523 4.395-1.35M12 6c4.008 0 6.701 3.158 9 6a15.66 15.66 0 0 1-1.078 1.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 13c3.6-8 14.4-8 18 0"/><path fill="currentColor" d="M12 17a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fsquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M10 16V8h4m-4 4h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceId(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 3H5a2 2 0 0 0-2 2v2m14-4h2a2 2 0 0 1 2 2v2m-5 1v2M8 8v2m1 6s1 1 3 1s3-1 3-1m-3-8v5h-1m-4 8H5a2 2 0 0 1-2-2v-2m14 4h2a2 2 0 0 0 2-2v-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceThreeDdraft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 7.78v7.796a.6.6 0 0 1-.27.502l-6.616 4.347a.6.6 0 0 1-.249.093l-10.184 1.39A.6.6 0 0 1 2 21.312v-12.3a.6.6 0 0 1 .298-.519l10.789-6.28a.6.6 0 0 1 .688.058l6.01 5.05A.6.6 0 0 1 20 7.78"/><path d="m2.5 9l10.227 2.922a.6.6 0 0 0 .506-.084L19.5 7.5m-6.5 13V12m3.5 2.01l.01-.011M22 17.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 2h-3a5 5 0 0 0-5 5v3H6v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5"/><path d="M11 21v-9c0-2.187.5-4 4-4m-6 5h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facetime(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 16V8a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v8a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path d="M6 13v-2a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17.04 9.22l-3.067 2.3a.6.6 0 0 0 0 .96l3.067 2.3a.6.6 0 0 0 .96-.48V9.7a.6.6 0 0 0-.96-.48"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Farm(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 20H2V8l6-3l6 3v12h-3m-6 0v-7h6v7m-6 0h6m7-6v6m-4-3h8m-8-3h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastArrowDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 13l6 6l6-6M6 5l6 6l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastArrowDownSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.5 7.5L12 11L8.5 7.5m7 6L12 17l-3.5-3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastArrowLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11 6l-6 6l6 6m8-12l-6 6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastArrowLeftSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.5 8.5L13 12l3.5 3.5m-6-7L7 12l3.5 3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastArrowRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13 6l6 6l-6 6M5 6l6 6l-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastArrowRightSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m8 8.5l3.5 3.5L8 15.5m6-7l3.5 3.5l-3.5 3.5"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastArrowUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 11l6-6l6 6M6 19l6-6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastArrowUpSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M15.5 16.5L12 13l-3.5 3.5m7-6L12 7l-3.5 3.5"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastDownCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.5 7.5L12 11l3.5-3.5m-7 6L12 17l3.5-3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastLeftCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M16.5 8.5L13 12l3.5 3.5m-6-7L7 12l3.5 3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastRightCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 8.5l3.5 3.5L8 15.5m6-7l3.5 3.5l-3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastUpCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.5 16.5L12 13l3.5 3.5m-7-6L12 7l3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FavouriteBook(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M16 8.78a2.28 2.28 0 0 1-.662 1.606c-.976.984-1.923 2.01-2.936 2.958a.597.597 0 0 1-.822-.017l-2.918-2.94a2.281 2.281 0 0 1 0-3.214a2.277 2.277 0 0 1 3.232 0L12 7.28l.106-.107A2.276 2.276 0 0 1 16 8.78Z"/><path stroke-linecap="round" d="M6 17h14M6 21h14"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FavouriteWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M13 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"/><path stroke-linejoin="round" d="M22 17.28a2.28 2.28 0 0 1-.662 1.606c-.976.984-1.923 2.01-2.936 2.958a.597.597 0 0 1-.823-.017l-2.918-2.94a2.281 2.281 0 0 1 0-3.214a2.277 2.277 0 0 1 3.233 0l.106.107l.106-.107A2.277 2.277 0 0 1 22 17.28Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 15a6 6 0 1 0 0-12a6 6 0 0 0 0 12m0 0v4m0 2v-2m0 0h-2m2 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Figma(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 6a3 3 0 0 1 3-3h3v6H9a3 3 0 0 1-3-3m6-3h3a3 3 0 0 1 0 6h-3z"/><path d="M12 12a3 3 0 1 1 6 0a3 3 0 0 1-6 0m-6 6a3 3 0 0 1 3-3h3v3a3 3 0 0 1-6 0m0-6a3 3 0 0 1 3-3h3v6H9a3 3 0 0 1-3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileNotFound(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v4m0 4.01l.01-.011M9 3H4v3m0 5v2m16-2v2M15 3h5v3M9 21H4v-3m11 3h5v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FillColor(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m2.636 10.293l7.071-7.071l8.485 8.485l-7.07 7.071a2 2 0 0 1-2.83 0l-5.656-5.657a2 2 0 0 1 0-2.828m5.657-8.485l1.414 1.414"/><path stroke-miterlimit="1.5" d="M20 15s3 2.993 3 4.887c0 1.655-1.345 3-3 3c-1.656 0-2.988-1.345-3-3C17.01 17.992 20 15 20 15" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FillColorSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="m2.636 10.293l7.071-7.071l8.485 8.485l-7.07 7.071a2 2 0 0 1-2.83 0l-5.656-5.657a2 2 0 0 1 0-2.828"/><path d="m8.293 1.808l1.414 1.414"/><path fill="currentColor" fill-rule="evenodd" stroke-miterlimit="1.5" d="M20 15s3 2.993 3 4.887c0 1.655-1.345 3-3 3c-1.656 0-2.988-1.345-3-3C17.01 17.992 20 15 20 15" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilletThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 20v-4C2 8.268 8.268 2 16 2h4m.839 18.84h-3.536m3.536 0v-3.536m0 3.535L18 18"/><path stroke-dasharray="2 3" d="m9 9l7 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 3h16a1 1 0 0 1 1 1v1.586a1 1 0 0 1-.293.707l-6.414 6.414a1 1 0 0 0-.293.707v6.305a1 1 0 0 1-1.243.97l-2-.5a1 1 0 0 1-.757-.97v-5.805a1 1 0 0 0-.293-.707L3.293 6.293A1 1 0 0 1 3 5.586V4a1 1 0 0 1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 7V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v3M3 7l6.65 5.7a1 1 0 0 1 .35.76v6.26a1 1 0 0 0 1.242.97l2-.5a1 1 0 0 0 .758-.97v-5.76a1 1 0 0 1 .35-.76L21 7M3 7h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterList(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h18M7 12h10m-6 6h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterListCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 9h12M8 13h8m-6 4h4m-2 5C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 3h16a1 1 0 0 1 1 1v1.586a1 1 0 0 1-.293.707l-6.414 6.414a1 1 0 0 0-.293.707v6.305a1 1 0 0 1-1.242.97l-2-.5a1 1 0 0 1-.758-.97v-5.805a1 1 0 0 0-.293-.707L3.293 6.293A1 1 0 0 1 3 5.586V4a1 1 0 0 1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Finder(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 15V9a6 6 0 0 1 6-6h6a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6"/><path d="M15 3s-4.5 0-4.5 9H13c0 9 2 9 2 9"/><path d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2M7 9v2m10-2v2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fingerprint(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 3.516A9.004 9.004 0 0 1 20.648 8.5M21 22v-8M3 22V11c0-1.052.18-2.062.512-3"/><path d="M18 22V11.3C18 7.82 15.314 5 12 5s-6 2.82-6 6.3V14m0 8v-4"/><path d="M9 22V11.15C9 9.41 10.343 8 12 8c.865 0 1.645.385 2.193 1M15 22v-8m-3 8v-3.5m0-7.5v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintCheckCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 16v-3.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10a9.98 9.98 0 0 1-.458 3M15.5 20.5l2 2l5-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 16v-3.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintLockCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 13.5v-.685m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M14 21.8c-.646.131-1.315.2-2 .2c-5.523 0-10-4.477-10-10S6.477 2 12 2s10 4.477 10 10c0 .254-.01.506-.028.755"/><path d="M21.167 18.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintScan(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 16v-3.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794M6 3H3v3m15-3h3v3M6 21H3v-3m15 3h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 16v-3.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M9 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v6"/><path stroke-linejoin="round" d="M12 21v-4.639c0-.51.1-.999.284-1.453M22 21v-3.185m-7.778-5.08A5.506 5.506 0 0 1 17 12c2.28 0 4.203 1.33 4.805 3.15M15 22v-2.177M19 22v-5.147C19 15.83 18.105 15 17 15s-2 .83-2 1.853v.794M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintXmarkCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 16v-4.639c0-.51.1-.999.285-1.453M17 14v-1.185m-7.778-5.08A5.506 5.506 0 0 1 12 7c2.28 0 4.203 1.33 4.805 3.15M10 17v-2.177M14 17v-5.147C14 10.83 13.105 10 12 10s-2 .83-2 1.853v.794"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10a9.98 9.98 0 0 1-.458 3m-4.421 7.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L17.12 18.12m2.122 2.122l2.121 2.121"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireFlame(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 18c0 2.415 1.79 3 4 3c3.759 0 5-2.5 2.5-7.5C11 18 10.5 11 11 9c-1.5 3-3 5.818-3 9"/><path d="M12 21c5.05 0 8-2.904 8-7.875C20 8.155 12 3 12 3S4 8.154 4 13.125C4 18.095 6.95 21 12 21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fish(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 9s0-2-1-4c4 0 6.5 2.5 6.5 2.5s3.5-.5 6 4.5c-1 5.5-6 6-6 6l-4 2.5v-3c-2.5-1-5-3.5-5-5S10.5 9 10.5 9m0 0s1-.5 2-.5M2 9.5l1 3l-1 3s5 0 5-3s-5-3-5-3m15 2.51l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fishing(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 0v10c0 6-10 6-10 0v-4l2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.456 2.665a.6.6 0 0 1 1.088 0l2.864 6.137a.6.6 0 0 0 .29.29l6.137 2.864a.6.6 0 0 1 0 1.088l-6.137 2.864a.6.6 0 0 0-.29.29l-2.864 6.137a.6.6 0 0 1-1.088 0l-2.864-6.137a.6.6 0 0 0-.29-.29l-6.137-2.864a.6.6 0 0 1 0-1.088l6.137-2.864a.6.6 0 0 0 .29-.29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 10V3L5 14h6v7l8-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlashOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.795 8.782L5 14h6v7l4-5.5m2.182-3L19 10h-6V3l-2.182 3M4 4l16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flask(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M18.5 15h-13"/><path stroke-linecap="round" d="M16 4H8m1 .5v5.76a2 2 0 0 1-.481 1.302L3.48 17.438A2 2 0 0 0 3 18.74V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-.26a2 2 0 0 0-.482-1.302l-5.036-5.876A2 2 0 0 1 15 10.26V4.5m-3 4.51l.01-.011M11 2.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flip(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 20H2L9.5 4zm10.625 0H22l-.937-2m-4.688 2H14.5v-2m0-6v2m3.75-2l.938 2m-2.813-6L14.5 4v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipReverse(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 12v2m-3.75-2l-.937 2m2.812 6H9.5v-2m-5.625 2H2l.938-2M7.625 8L9.5 4v4m5 12H22L14.5 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyDisk(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h11.172a2 2 0 0 1 1.414.586l2.828 2.828A2 2 0 0 1 21 7.828V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path d="M8.6 9h6.8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6H8.6a.6.6 0 0 0-.6.6v4.8a.6.6 0 0 0 .6.6ZM6 13.6V21h12v-7.4a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 0-.6.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyDiskArrowIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 7.5V5a2 2 0 0 1 2-2h11.172a2 2 0 0 1 1.414.586l2.828 2.828A2 2 0 0 1 21 7.828V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2.5M6 21v-4"/><path d="M18 21v-7.4a.6.6 0 0 0-.6-.6H15m1-10v5.4a.6.6 0 0 1-.6.6h-1.9M8 3v3m-7 6h11m0 0L9 9m3 3l-3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyDiskArrowOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 6.5V5a2 2 0 0 1 2-2h11.172a2 2 0 0 1 1.414.586l2.828 2.828A2 2 0 0 1 21 7.828V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-1.5"/><path d="M8 3h8v5.4a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6zm10 18v-7.4a.6.6 0 0 0-.6-.6H15m-9 8v-3.5m6-5.5H1m0 0l3-3m-3 3l3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flower(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6m1-6s1-2 1-4s-2-4-2-4s-2 2-2 4s1 4 1 4"/><path d="M9 11s-2-1-4-1s-4 2-4 2s2 2 4 2s4-1 4-1m4 2s1 2 1 4s-2 4-2 4s-2-2-2-4s1-4 1-4m4-4s2-1 4-1s4 2 4 2s-2 2-4 2s-4-1-4-1m-4.414-3.828S9.879 7.05 8.464 5.636C7.05 4.222 4.222 4.222 4.222 4.222s0 2.828 1.414 4.243c1.414 1.414 3.536 2.121 3.536 2.121m0 2.828s-2.122.707-3.536 2.122c-1.414 1.414-1.414 4.242-1.414 4.242s2.828 0 4.242-1.414c1.415-1.414 2.122-3.536 2.122-3.536m4.243-1.414s2.12.707 3.535 2.122c1.414 1.414 1.414 4.242 1.414 4.242s-2.828 0-4.242-1.414c-1.415-1.414-2.122-3.536-2.122-3.536m0-5.656s.707-2.122 2.122-3.536c1.414-1.414 4.242-1.414 4.242-1.414s0 2.828-1.414 4.243c-1.414 1.414-3.536 2.121-3.536 2.121"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fog(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 14h6m-6 8h6m-8-4h10m-13.5-.618C2.188 16.707 1 15.388 1 13c0-4 3.333-5 5-5c0-2 0-6 6-6s6 4 6 6c1.667 0 5 1 5 5c0 2.388-1.188 3.707-2.5 4.382"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 11V4.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H21.4a.6.6 0 0 1 .6.6V11M2 11v8.4a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6V11M2 11h20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 6h4m-.6 14H2.6a.6.6 0 0 1-.6-.6V11h19.4a.6.6 0 0 1 .6.6v7.8a.6.6 0 0 1-.6.6M2 11V4.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 6h2m2 0h-2m0 0V4m0 2v2m1.4 12H2.6a.6.6 0 0 1-.6-.6V11h19.4a.6.6 0 0 1 .6.6v7.8a.6.6 0 0 1-.6.6M2 11V4.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSettings(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.6 4h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H21.4a.6.6 0 0 1 .6.6v2.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6M22 10v4"/><path d="M2 10v9.4a.6.6 0 0 0 .6.6H13m6 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path stroke-dasharray=".3 2" d="M19 22a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 3v4m0 4.01l.01-.011M22 7v12.4a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V11"/><path d="M14 7h-1.278a.6.6 0 0 1-.39-.144L9.169 4.144A.6.6 0 0 0 8.778 4H2.6a.6.6 0 0 0-.6.6V11h12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontQuestion(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.469 18.374l1.064-2.341m9.58 2.341l-1.064-2.341m0 0L8.79 6.664l-4.258 9.368m8.516 0H4.533m10.645-7.238c0-3.725 5.854-3.725 5.854 0c0 2.661-2.66 2.13-2.66 5.322m-.001 4.269l.01-.012"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Football(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 13.828V21M5 3v10.828h14V3"/><path d="M13 6.732c1.071-.618 1.434-2.114 1.549-2.833a.505.505 0 0 0-.321-.556c-.68-.26-2.157-.693-3.228-.075C9.93 3.886 9.567 5.38 9.452 6.1a.505.505 0 0 0 .32.556c.681.26 2.158.693 3.228.075"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FootballBall(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.115 14.015a22.314 22.314 0 0 0-.103 3.665a2.413 2.413 0 0 0 2.309 2.308c1.007.052 2.294.055 3.664-.103m-5.87-5.87C4.394 11.604 5.17 8.93 7.05 7.05c1.88-1.88 4.554-2.656 6.965-2.935m-9.9 9.9l5.87 5.87m0 0c2.411-.279 5.084-1.055 6.965-2.935c1.88-1.88 2.656-4.554 2.935-6.965m-5.87-5.87a22.314 22.314 0 0 1 3.665-.103a2.413 2.413 0 0 1 2.308 2.309a22.312 22.312 0 0 1-.103 3.664m-5.87-5.87l5.87 5.87M9.172 14.828l1.414-1.414m0 0L9.172 12m1.414 1.414L12 14.828m-1.414-1.414l2.828-2.828m0 0l1.414-1.414m-1.414 1.414L12 9.172m1.414 1.414L14.828 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.956 5.704A.6.6 0 0 0 2 6.187v11.626a.6.6 0 0 0 .956.483l7.889-5.813a.6.6 0 0 0 0-.966zm11 0a.6.6 0 0 0-.956.483v11.626a.6.6 0 0 0 .956.483l7.889-5.813a.6.6 0 0 0 0-.966z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardFifteenSeconds(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 13a9 9 0 1 1-9-9m0 0h7.5m0 0l-2-2m2 2l-2 2M9 9v7"/><path d="M15 9h-2a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h1a1 1 0 0 1 1 1V15a1 1 0 0 1-1 1h-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardMessage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 8l5 3l5-3"/><path d="M10 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v6.857"/><path stroke-linejoin="round" d="M22 17.111h-6.3c-3.6 0-3.6 4.889 0 4.889m6.3-4.889L18.85 14M22 17.111l-3.15 3.111"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.956 5.704A.6.6 0 0 0 2 6.187v11.626a.6.6 0 0 0 .956.483l7.889-5.813a.6.6 0 0 0 0-.966zm11 0a.6.6 0 0 0-.956.483v11.626a.6.6 0 0 0 .956.483l7.889-5.813a.6.6 0 0 0 0-.966z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frame(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M4.998 2.001H2v2.998h2.998zm0 8.501H2V13.5h2.998zM20.498 5v5.503M3.5 5v5.503m16.998 2.999v5.502M3.5 13.502v5.502m1.499 1.498h5.5"/><path stroke-width="1.22" d="M4.999 3.503h5.5"/><path d="M13.498 20.499h5.5"/><path stroke-width="1.22" d="M13.498 3.501h5.5"/><path d="M4.998 19.001H2v2.998h2.998zM21.997 2.002H19V5h2.998zM13.497 2H10.5v2.998h2.998zm8.5 8.503H19V13.5h2.998zm0 8.499H19V22h2.998zm-8.5-.002H10.5v2.998h2.998z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke-width="1.5"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 3v18M18 3v18M3 6h18"/><path fill="currentColor" fill-rule="evenodd" d="M9.6 9h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6V9.6a.6.6 0 0 1 .6-.6" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 18h18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameAltEmpty(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 3v18M18 3v18M3 6h18M3 18h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameMinusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-miterlimit="1.5" d="M4.998 2H2v2.998h2.998zm.001 1.5h14M3.5 4.998V19M20.498 5v14.002M4.999 20.5h14M4.998 19H2v2.998h2.998zM21.997 2.001H19v2.998h2.998zm0 17H19v2.998h2.998z"/><path d="M9 12h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FramePlusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-miterlimit="1.5" d="M4.998 2H2v2.998h2.998zm.001 1.5h14M3.5 4.998V19M20.498 5v14.002M4.999 20.5h14M4.998 19H2v2.998h2.998zM21.997 2.001H19v2.998h2.998zm0 17H19v2.998h2.998z"/><path d="M9 12h3m3 0h-3m0 0V9m0 3v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameSelect(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M4.998 2H2v2.998h2.998zm.001 1.501h14M3.5 4.999V19M20.498 5v14.002M4.999 20.501h14M4.998 19H2v2.998h2.998zM21.997 2.002H19V5h2.998zm0 17H19V22h2.998z"/><path d="m10.997 15.002l-3-7l7 3l-2.998.999zm1.002-3l2.998 3z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameSimple(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5" d="M4.998 2H2v2.998h2.998zm0 1.501h14M3.499 4.998V19M20.497 5v14.002M4.998 20.501h14M4.998 19H2v2.998h2.998zM21.996 2.002h-2.998V5h2.998zm0 17h-2.998V22h2.998z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameTool(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 7h1M2 17h1M21 7h1m-1 10h1M17 3V2M7 3V2m10 20v-1M7 22v-1M18 6.6v10.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6V6.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameToolSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 7h1M2 17h1M21 7h1m-1 10h1M17 3V2M7 3V2m10 20v-1M7 22v-1"/><path fill="currentColor" d="M6 17.4V6.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6v10.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fridge(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10 14H9m1-8H9"/><path d="M5 10V2.6a.6.6 0 0 1 .6-.6h12.8a.6.6 0 0 1 .6.6V10M5 10v11.4a.6.6 0 0 0 .6.6h12.8a.6.6 0 0 0 .6-.6V10M5 10h14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fx(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 17V7h7m-7 5h5m5 5l4-5m0 0l4-5m-4 5l-4-5m4 5l4 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FxTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 15V9h5m-5 3h3.571M13 15l2.5-3m0 0L18 9m-2.5 3L13 9m2.5 3l2.5 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gamepad(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 17.5c2.5 3.5 6.449.915 5.5-2.5c-1.425-5.129-2.2-7.984-2.603-9.492A2.032 2.032 0 0 0 18.438 4H5.562c-.918 0-1.718.625-1.941 1.515C2.78 8.863 2.033 11.802 1.144 15c-.948 3.415 3 6 5.5 2.5M18 8.5l.011.01M16.49 7l.011.01M16.49 10l.011.01M15 8.5l.011.01M7 7v3M5.5 8.5h3"/><path d="M8 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4m8 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Garage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 20H3V6l9-2l9 2v14h-3M6 20h12M6 20v-4m12 4v-4M6 12V8h12v4M6 12h12M6 12v4m12-4v4M6 16h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gas(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9 8a3 3 0 0 1 3-3v0a3 3 0 0 1 3 3v13.4a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6zm0 3h6m-3-6V2m0 0h-1m1 0h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GasTank(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path stroke-width="1.493" d="M3 7.562A2.562 2.562 0 0 1 5.563 5H7V3h5v2h2.002A6.998 6.998 0 0 1 21 11.998v6.442a2.562 2.562 0 0 1-2.563 2.562H5.563A2.565 2.565 0 0 1 3 18.44z" clip-rule="evenodd"/><path stroke-width="1.502" d="m8 8.878l8 8.238l-4-4.121l-4 4.12l4-4.12l4-4.121"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GasTankDroplet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5" clip-rule="evenodd"><path stroke-width="1.493" d="M3 7.562A2.562 2.562 0 0 1 5.563 5H7V3h5v2h2.002A6.998 6.998 0 0 1 21 11.998v6.442a2.562 2.562 0 0 1-2.563 2.562H5.563A2.565 2.565 0 0 1 3 18.44z"/><path d="M12 9s3 2.993 3 4.887c0 1.655-1.345 3-3 3c-1.656 0-2.988-1.345-3-3C9.01 11.992 12 9 12 9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GifFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M15.5 15V9h3m-3 3h2M12 15V9M8.5 9h-3v6h3v-2.4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 12v9.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V12m17.4-5H2.6a.6.6 0 0 0-.6.6v3.8a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6M12 22V7m0 0H7.5a2.5 2.5 0 1 1 0-5C11 2 12 7 12 7m0 0h4.5a2.5 2.5 0 0 0 0-5C13 2 12 7 12 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitBranch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4M6 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0-4V3"/><path d="M8 18h1c3.5 0 9-2.1 9-8.5V8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCherryPickCommit(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 11v-1a2 2 0 0 0-2-2h-3m-5 3v-1a2 2 0 0 1 2-2h3m0 0V4m0 16a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-3-3H3m12 0h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCommit(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-3-3H3m12 0h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCompare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4M6 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4m12 10V7s0-2-2-2h-3M6 7v10s0 2 2 2h3"/><path d="M15 7.5L12.5 5L15 2.5m-6.5 14L11 19l-2.5 2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitFork(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4M7 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4M7 7v10M17 7v1c0 2.5-2 3-2 3l-6 2s-2 .5-2 3v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitMerge(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4M6 21V7"/><path d="M6 7v2c0 3.5 2.5 9 8.5 9H16M6 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequest(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4M6 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4M6 7v10m12 0V7s0-2-2-2h-3"/><path d="M15 7.5L12.5 5L15 2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequestClosed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4M6 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0-10v6m12 0V7s0-2-2-2h-4M4 7.243L6.121 5.12m0 0L8.243 3M6.12 5.121L4 3m2.121 2.121l2.122 2.122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 22.027v-2.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7a5.44 5.44 0 0 0-1.5-3.75a5.07 5.07 0 0 0-.09-3.77s-1.18-.35-3.91 1.48a13.38 13.38 0 0 0-7 0c-2.73-1.83-3.91-1.48-3.91-1.48A5.07 5.07 0 0 0 5 5.797a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7a3.37 3.37 0 0 0-.94 2.58v2.87"/><path d="M9 20.027c-3 .973-5.5 0-7-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M14.333 19v-1.863c.025-.31-.018-.62-.126-.913a2.18 2.18 0 0 0-.5-.781c2.093-.227 4.293-1 4.293-4.544a3.48 3.48 0 0 0-1-2.434a3.211 3.211 0 0 0-.06-2.448s-.787-.227-2.607.961a9.152 9.152 0 0 0-4.666 0c-1.82-1.188-2.607-.96-2.607-.96A3.211 3.211 0 0 0 7 8.464a3.482 3.482 0 0 0-1 2.453c0 3.519 2.2 4.291 4.293 4.544a2.18 2.18 0 0 0-.496.773a2.134 2.134 0 0 0-.13.902V19"/><path d="M9.667 17.702c-2 .631-3.667 0-4.667-1.948"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitlabFull(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M17.057 2.544a.2.2 0 0 1 .378-.008l3.114 8.31l1.398 3.73a.2.2 0 0 1-.07.232l-9.76 7.106a.2.2 0 0 1-.235 0l-9.76-7.106a.2.2 0 0 1-.069-.231l1.398-3.73l.167-.45l2.944-7.861a.2.2 0 0 1 .378.008l2.47 7.6a.2.2 0 0 0 .19.137h4.8a.2.2 0 0 0 .19-.138z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassEmpty(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m7.5 11l1 5"/><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassFragile(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 22h4m4 0h-4m0 0v-7m-5.422-4.952C7.783 12.682 12 15 12 15s4.217-2.318 5.422-4.952c1.3-2.845 0-8.048 0-8.048H6.578s-1.3 5.203 0 8.048"/><path d="m12.5 2l-2 4h3l-2 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassHalf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0M4 13c1.032 1.203 3.925 1.864 7 1.981a25.406 25.406 0 0 0 4-.158c2.266-.279 4.197-.886 5-1.823M4 13c2.286-2.667 13.714-2.667 16 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassHalfAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.04 4.294a.496.496 0 0 1 .191-.479C3.927 3.32 6.314 2 12 2s8.073 1.32 8.769 1.815a.496.496 0 0 1 .192.479l-1.7 12.744a4 4 0 0 1-1.98 2.944l-.32.183a10 10 0 0 1-9.922 0l-.32-.183a4 4 0 0 1-1.98-2.944z"/><path d="M3 5c2.571 2.667 15.429 2.667 18 0M4 13c1.032 1.203 3.925 1.864 7 1.981c3.739.143 7.746-.518 9-1.981m-5 1.823V20.5M4 13c2.286-2.667 13.714-2.667 16 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glasses(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 14a4 4 0 1 0 8 0a4 4 0 0 0-8 0m0 0V6m20 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0m0 0V6m-8 8h-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="m2.5 12.5l5.5 2L7 18l1 3m9-.5l-.5-2.5l-2.5-1v-3.5l3-1l4.5.5M19 5.5L18.5 7l-3.5.5v3l2.5-1h2l2 1m-19 0l2.5-2L7.5 8l2-3l-1-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Golf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 18v-6m0 0V3.41a.6.6 0 0 1 .836-.552l8.444 3.62a.6.6 0 0 1 .022 1.093zm0 10c3.866 0 7-1.567 7-3.5S15.866 15 12 15s-7 1.567-7 3.5S8.134 22 12 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M15.548 8.303A5.148 5.148 0 0 0 12.108 7C9.288 7 7 9.239 7 12s2.287 5 5.109 5c3.47 0 4.751-2.57 4.891-4.583h-4.159"/><path stroke-linecap="round" stroke-linejoin="round" d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M15.548 8.303A5.148 5.148 0 0 0 12.108 7C9.288 7 7 9.239 7 12s2.287 5 5.109 5c3.47 0 4.751-2.57 4.891-4.583h-4.159"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDocs(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2M7 7h10M7 12h10M7 17h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDrive(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.143 3.004L14.857 3m-5.714.004L2 15.004m7.143-12L18.433 21M14.856 3L22 15.004M14.857 3L5.575 21m12.857 0H5.575m12.857 0L22 15.004M5.575 21L2 15.004m20 0H2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDriveCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.143 3.004L14.857 3m-5.714.004L2 15.004m7.143-12l4.902 9.496m.812-9.5L5.575 21m9.282-18L21.5 14M5.575 21L2 15.004M5.575 21h6.429M2 15.004h10.5M15 19l3 3l5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDriveSync(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.143 3.004L14.857 3m-5.714.004L2 15.004m7.143-12l4.902 9.496m.812-9.5L5.575 21m9.282-18l5.356 9M5.575 21L2 15.004M5.575 21h6.429M2 15.004h10.5m10.166 2.663C22.048 16.097 20.634 15 18.99 15c-1.758 0-3.252 1.255-3.793 3"/><path d="M20.995 17.772H22.4a.6.6 0 0 0 .6-.6V15.55m-7.666 4.783C15.952 21.903 17.366 23 19.01 23c1.758 0 3.252-1.255 3.793-3"/><path d="M17.005 20.228H15.6a.6.6 0 0 0-.6.6v1.622"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDriveWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.143 3.004L14.857 3m-5.714.004L2 15.004m7.143-12l4.902 9.496m.812-9.5L5.575 21m9.282-18L21.5 14M5.575 21L2 15.004M5.575 21h6.429M2 15.004h10.5M18 16v2m0 4.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleHome(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M17.708 17A9 9 0 1 0 4.291 5a9 9 0 0 0 13.417 12Zm0 0H19.5a2.5 2.5 0 0 1 2.5 2.5v0a2.5 2.5 0 0 1-2.5 2.5H17"/><path stroke-linejoin="round" d="m11 11.01l.01-.011M8 11.01l.01-.011m5.99.011l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleOne(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11 5v14a2 2 0 1 0 4 0V5a2 2 0 1 0-4 0"/><path d="M11.64 3.53L6.747 8.171a2 2 0 0 0 2.754 2.901l4.892-4.642a2 2 0 0 0-2.753-2.902"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gps(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 8.5h-2.25A1.75 1.75 0 0 0 18 10.25v0c0 .966.784 1.75 1.75 1.75h1.5c.966 0 1.75.784 1.75 1.75v0a1.75 1.75 0 0 1-1.75 1.75H18m-7.5 0v-2.8m0 0h2.857c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H10.5zm-4-3.573a3.5 3.5 0 1 0-2 6.373C6.433 15.5 8 14 8 12H5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraduationCap(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m2.573 8.463l8.659-4.329a.6.6 0 0 1 .536 0l8.659 4.33a.6.6 0 0 1 0 1.073l-8.659 4.329a.6.6 0 0 1-.536 0l-8.659-4.33a.6.6 0 0 1 0-1.073"/><path d="M22.5 13V9.5l-2-1m-16 2v5.412a2 2 0 0 0 1.142 1.807l5 2.374a2 2 0 0 0 1.716 0l5-2.374a2 2 0 0 0 1.142-1.807V10.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 20H4V4"/><path d="m4 7l8 8l3-3l4.5 4.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 20H4V4"/><path d="M4 16.5L12 9l3 3l4.5-4.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M13.992 17h6"/><path d="M4 9.4V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm0 10v-4.8a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm10-10V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-4.8a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M13.992 17h3m3 0h-3m0 0v-3m0 3v3"/><path d="M4 9.4V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm0 10v-4.8a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm10-10V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-4.8a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M14.871 19.121L16.993 17m2.121-2.121L16.993 17m0 0l-2.122-2.121M16.993 17l2.121 2.121"/><path d="M4 9.4V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm0 10v-4.8a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6Zm10-10V4.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6h-4.8a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M1 20v-1a7 7 0 0 1 7-7v0a7 7 0 0 1 7 7v1"/><path d="M13 14v0a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v.5"/><path stroke-linejoin="round" d="M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8m10-3a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gym(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.4 7H4.6a.6.6 0 0 0-.6.6v8.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6m12 0h-2.8a.6.6 0 0 0-.6.6v8.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6"/><path d="M1 14.4V9.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v4.8a.6.6 0 0 1-.6.6H1.6a.6.6 0 0 1-.6-.6m22 0V9.6a.6.6 0 0 0-.6-.6h-1.8a.6.6 0 0 0-.6.6v4.8a.6.6 0 0 0 .6.6h1.8a.6.6 0 0 0 .6-.6M8 12h8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hsquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6M10 8v4m0 4v-4m0 0h4m0 0V8m0 4v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HalfCookie(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.8 14c-.927 4.564-4.962 8-9.8 8c-5.523 0-10-4.477-10-10c0-5.185 3.947-9.449 9-9.95"/><path d="M6.5 10a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m14-6a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M12 19a1 1 0 1 1 0-2a1 1 0 0 1 0 2m-5-3.99l.01-.011m9.99.011l.01-.011M11 12.01l.01-.011M21 9.01l.01-.011M17 6.01l.01-.011M11 2c-.5 1.5.5 3 2.085 3C11 8.5 13 12 18 11.5c0 2.5 2.5 3 3.7 2.514"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HalfMoon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 11.507a9.493 9.493 0 0 0 18 4.219c-8.507 0-12.726-4.22-12.726-12.726A9.494 9.494 0 0 0 3 11.507"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hammer(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.634 11.056L2.148 19.54l2.122 2.121l8.485-8.485"/><path d="m10.634 11.056l1.414-1.415s.354-3.182-3.182-6.717l1.06-1.06l8.486 5.656l-1.06 1.06l1.413 1.415l1.061-1.06l2.475 2.474l-4.95 4.95l-2.475-2.475l1.061-1.06l-1.414-1.415l-1.768 1.768z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandBrake(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 16v-4m0-3V8"/><circle cx="12" cy="12" r="8"/><path stroke-linecap="round" stroke-linejoin="round" d="M3.953 4.5A10.961 10.961 0 0 0 1 12c0 2.899 1.121 5.535 2.953 7.5m16.094-15A10.962 10.962 0 0 1 23 12c0 2.899-1.121 5.535-2.953 7.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandCard(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M11 9h11M2 11l2.807-3.157A4 4 0 0 1 7.797 6.5H8m-6 13h5.5l4-3s.81-.547 2-1.5c2.5-2 0-5.166-2.5-3.5C8.964 12.857 7 14 7 14"/><path d="M8 13.5V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-6.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandCash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m2 11l2.807-3.157A4 4 0 0 1 7.797 6.5H8m-6 13h5.5l4-3s.81-.547 2-1.5c2.5-2 0-5.166-2.5-3.5C8.964 12.857 7 14 7 14"/><path d="M8 13.5V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-6.5"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4m4.5-1.99l.01-.011m-9.01.011l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandContactless(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m2 11l2.807-3.157A4 4 0 0 1 7.797 6.5H8m-6 13h5.5l4-3s.81-.547 2-1.5c2.5-2 0-5.166-2.5-3.5C8.964 12.857 7 14 7 14"/><path d="M8 13.5V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-6.5"/><path stroke-linecap="round" stroke-linejoin="round" d="M18.25 12c.5-1.5.5-2.5 0-4M16 9c.227.5.227 1.5 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Handbag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 8H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-9a2 2 0 0 0-2-2h-5M9 8V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6V8M9 8h6M9 8v6m6-6v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardDrive(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m10 17.01l.01-.011M6 17.01l.01-.011"/><path d="M2 13v7.4a.6.6 0 0 0 .6.6h18.8a.6.6 0 0 0 .6-.6V13M2 13h20M2 13l2.872-9.572A.6.6 0 0 1 5.446 3h13.108a.6.6 0 0 1 .574.428L22 13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hashtag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M10 3L6 21m14.5-5h-18M22 7H4m14-4l-4 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 17v-2a7 7 0 1 1 14 0v2zm0 0H2M14 6.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hd(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 7v5m0 5v-5m0 0h7m0 0V7m0 5v5m3-5V7c4 0 8 0 8 5s-4 5-8 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HdDisplay(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 8.5V12m0 3.5V12m0 0h4.5m0 0V8.5m0 3.5v3.5M14 12V8.5c2.5 0 5 0 5 3.5s-2.5 3.5-5 3.5z"/><path d="M1 15V9a6 6 0 0 1 6-6h10a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H7a6 6 0 0 1-6-6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hdr(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.5 8.5V12m0 3.5V12m0 0H6m0 0V8.5M6 12v3.5m11.5 0v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H17.5v4.2m2.857 0H17.5m2.857 0l2.143 2.8M9.5 12V8.5c2.5 0 5 0 5 3.5s-2.5 3.5-5 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headset(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 13.5l-.485.121A2 2 0 0 0 2 15.561v1.877a2 2 0 0 0 1.515 1.94l1.74.436A.6.6 0 0 0 6 19.23v-5.463a.6.6 0 0 0-.746-.582zm0 0V13c0-4.97 3.582-9 8-9s8 4.03 8 9v.5m0 0l.485.121A2 2 0 0 1 22 15.561v1.877a2 2 0 0 1-1.515 1.94l-1.74.436A.6.6 0 0 1 18 19.23v-5.463a.6.6 0 0 1 .745-.582z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetBolt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.5 13L10 17h4l-2.5 4"/><path d="m4 13.5l-.485.121A2 2 0 0 0 2 15.561v1.877a2 2 0 0 0 1.515 1.94l1.74.436A.6.6 0 0 0 6 19.23v-5.463a.6.6 0 0 0-.746-.582zm0 0V13c0-4.97 3.582-9 8-9s8 4.03 8 9v.5m0 0l.485.121A2 2 0 0 1 22 15.561v1.877a2 2 0 0 1-1.515 1.94l-1.74.436A.6.6 0 0 1 18 19.23v-5.463a.6.6 0 0 1 .745-.582z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetBoltSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.5 13L10 17h4l-2.5 4"/><path d="M4 13.5V13c0-4.97 3.582-9 8-9s8 4.03 8 9v.5"/><path fill="currentColor" d="M2 17.438v-1.877a2 2 0 0 1 1.515-1.94L4 13.5l1.254-.314a.6.6 0 0 1 .746.582v5.463a.6.6 0 0 1-.746.582l-1.74-.434A2 2 0 0 1 2 17.438m20 0v-1.877a2 2 0 0 0-1.515-1.94L20 13.5l-1.255-.314a.6.6 0 0 0-.745.582v5.463a.6.6 0 0 0 .745.582l1.74-.434A2 2 0 0 0 22 17.438"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetHelp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="m4 11.5l-.485.121A2 2 0 0 0 2 13.561v1.877a2 2 0 0 0 1.515 1.94l1.74.435A.6.6 0 0 0 6 17.231v-5.463a.6.6 0 0 0-.746-.582zm0 0V11a8 8 0 1 1 16 0v.5m0 0l.485.121A2 2 0 0 1 22 13.561v1.877a2 2 0 0 1-1.515 1.94L20 17.5m0-6l-1.255-.314a.6.6 0 0 0-.745.582v5.463a.6.6 0 0 0 .745.582L20 17.5m-5 3h3a2 2 0 0 0 2-2v-1m-5 3a1.5 1.5 0 0 0-1.5-1.5h-3a1.5 1.5 0 0 0 0 3h3a1.5 1.5 0 0 0 1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 13.5V13c0-4.97 3.582-9 8-9s8 4.03 8 9v.5"/><path fill="currentColor" d="M2 17.438v-1.877a2 2 0 0 1 1.515-1.94L4 13.5l1.254-.314a.6.6 0 0 1 .746.582v5.463a.6.6 0 0 1-.746.582l-1.74-.434A2 2 0 0 1 2 17.438m20 0v-1.877a2 2 0 0 0-1.515-1.94L20 13.5l-1.255-.314a.6.6 0 0 0-.745.582v5.463a.6.6 0 0 0 .745.582l1.74-.434A2 2 0 0 0 22 17.438"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 13v5m0 3.01l.01-.011M4 13.5l-.485.121A2 2 0 0 0 2 15.561v1.877a2 2 0 0 0 1.515 1.94l1.74.436A.6.6 0 0 0 6 19.23v-5.463a.6.6 0 0 0-.746-.582zm0 0V13c0-4.97 3.582-9 8-9s8 4.03 8 9v.5m0 0l.485.121A2 2 0 0 1 22 15.561v1.877a2 2 0 0 1-1.515 1.94l-1.74.436A.6.6 0 0 1 18 19.23v-5.463a.6.6 0 0 1 .745-.582z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetWarningSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12v5m0 4.01l.01-.011M4 13.5V13c0-4.97 3.582-9 8-9s8 4.03 8 9v.5"/><path fill="currentColor" d="M2 17.438v-1.877a2 2 0 0 1 1.515-1.94L4 13.5l1.254-.314a.6.6 0 0 1 .746.582v5.463a.6.6 0 0 1-.746.582l-1.74-.434A2 2 0 0 1 2 17.438m20 0v-1.877a2 2 0 0 0-1.515-1.94L20 13.5l-1.255-.314a.6.6 0 0 0-.745.582v5.463a.6.6 0 0 0 .745.582l1.74-.434A2 2 0 0 0 22 17.438"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HealthShield(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.667 16h-3.334v-2.333H8v-3.334h2.333V8h3.334v2.333H16v3.334h-2.333z"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Healthcare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m18 20l3.824-3.824a.6.6 0 0 0 .176-.424V10.5A1.5 1.5 0 0 0 20.5 9v0a1.5 1.5 0 0 0-1.5 1.5V15"/><path d="m18 16l.858-.858a.484.484 0 0 0 .142-.343v0a.485.485 0 0 0-.268-.433l-.443-.221a2 2 0 0 0-2.308.374l-.895.895a2 2 0 0 0-.586 1.414V20M6 20l-3.824-3.824A.6.6 0 0 1 2 15.752V10.5A1.5 1.5 0 0 1 3.5 9v0A1.5 1.5 0 0 1 5 10.5V15"/><path d="m6 16l-.858-.858A.485.485 0 0 1 5 14.799v0c0-.183.104-.35.268-.433l.443-.221a2 2 0 0 1 2.308.374l.895.895a2 2 0 0 1 .586 1.414V20m4.167-8h-3.334V9.667H8V6.333h2.333V4h3.334v2.333H16v3.334h-2.333z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M22 8.862a5.95 5.95 0 0 1-1.654 4.13c-2.441 2.531-4.809 5.17-7.34 7.608c-.581.55-1.502.53-2.057-.045l-7.295-7.562c-2.205-2.286-2.205-5.976 0-8.261a5.58 5.58 0 0 1 8.08 0l.266.274l.265-.274A5.612 5.612 0 0 1 16.305 3c1.52 0 2.973.624 4.04 1.732A5.95 5.95 0 0 1 22 8.862Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartArrowDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.478 14.883l-1.824-1.89c-2.205-2.286-2.205-5.976 0-8.261a5.58 5.58 0 0 1 8.08 0l.266.274l.265-.274A5.612 5.612 0 0 1 16.305 3c1.52 0 2.973.624 4.04 1.732A5.95 5.95 0 0 1 22 8.862a5.95 5.95 0 0 1-1.654 4.13c-.603.626-1.202 1.258-1.8 1.891M12 21.5V11m4 6.5l-4 4l-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.942a6.351 6.351 0 0 1 4.305-1.692c1.726 0 3.374.71 4.58 1.96a6.7 6.7 0 0 1 1.865 4.652a6.7 6.7 0 0 1-1.865 4.652c-.796.825-1.591 1.67-2.39 2.518c-1.624 1.724-3.265 3.467-4.97 5.108l-.003.004a2.213 2.213 0 0 1-3.113-.069l-7.295-7.561c-2.485-2.577-2.485-6.727 0-9.303A6.328 6.328 0 0 1 12 3.942" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeatingSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M8 6s-2.5 3 0 6s0 6 0 6m4-12s-2.5 3 0 6s0 6 0 6m4-12s-2.5 3 0 6s0 6 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeavyRain(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 13v7m8-7v7m-4-5v7m8-4.393c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M9 9c0-3.5 5.5-3.5 5.5 0c0 2.5-2.5 2-2.5 5m0 4.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12m9.095-4.397C10 7.895 9.75 8.341 9.75 9a.75.75 0 0 1-1.5 0c0-1.091.437-1.958 1.124-2.54c.67-.57 1.538-.835 2.376-.835c.838 0 1.705.265 2.376.834c.687.583 1.124 1.45 1.124 2.541c0 .766-.196 1.35-.517 1.83c-.269.404-.619.716-.894.962l-.002.002l-.085.076c-.308.276-.539.504-.709.804c-.162.287-.293.688-.293 1.326a.75.75 0 1 1-1.5 0c0-.862.181-1.524.488-2.065c.299-.528.693-.894 1.01-1.18l.072-.065c.3-.27.508-.455.665-.692c.149-.222.265-.514.265-.998c0-.659-.25-1.105-.595-1.397c-.36-.306-.868-.478-1.405-.478s-1.045.172-1.405.478m2.223 10.898a.75.75 0 0 0-1.115-1.004l-.01.011a.75.75 0 0 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M9 9c0-3.5 5.5-3.5 5.5 0c0 2.5-2.5 2-2.5 5m0 4.01l.01-.011"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zm8.095 4.003C10 7.895 9.75 8.341 9.75 9a.75.75 0 0 1-1.5 0c0-1.091.437-1.958 1.124-2.54c.67-.57 1.538-.835 2.376-.835c.838 0 1.705.265 2.376.834c.687.583 1.124 1.45 1.124 2.541c0 .766-.196 1.35-.517 1.83c-.269.404-.619.716-.894.962l-.087.078c-.308.276-.539.504-.709.804c-.162.287-.293.688-.293 1.326a.75.75 0 1 1-1.5 0c0-.862.181-1.524.488-2.065c.299-.528.693-.894 1.01-1.18l.072-.065c.3-.27.508-.455.665-.692c.149-.222.265-.514.265-.998c0-.659-.25-1.105-.595-1.397c-.36-.306-.868-.478-1.405-.478s-1.045.172-1.405.478m2.223 10.898a.75.75 0 0 0-1.115-1.004l-.01.011a.75.75 0 0 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heptagon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.74 1.625a.6.6 0 0 1 .52 0l8.08 3.891a.6.6 0 0 1 .324.407l1.996 8.743a.6.6 0 0 1-.116.508l-5.591 7.01a.6.6 0 0 1-.47.227H7.517a.6.6 0 0 1-.469-.226l-5.591-7.011a.6.6 0 0 1-.116-.508l1.996-8.743a.6.6 0 0 1 .324-.407z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hexagon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.7 1.173a.6.6 0 0 1 .6 0l8.926 5.154a.6.6 0 0 1 .3.52v10.307a.6.6 0 0 1-.3.52L12.3 22.826a.6.6 0 0 1-.6 0l-8.926-5.154a.6.6 0 0 1-.3-.52V6.847a.6.6 0 0 1 .3-.52z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HexagonAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.327 2.774a.6.6 0 0 1 .52-.3h10.307a.6.6 0 0 1 .52.3l5.153 8.926a.6.6 0 0 1 0 .6l-5.154 8.926a.6.6 0 0 1-.52.3H6.847a.6.6 0 0 1-.52-.3L1.174 12.3a.6.6 0 0 1 0-.6l5.154-8.926Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HexagonDice(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M11.7 1.173a.6.6 0 0 1 .6 0l8.926 5.154a.6.6 0 0 1 .3.52v10.307a.6.6 0 0 1-.3.52L12.3 22.826a.6.6 0 0 1-.6 0l-8.926-5.154a.6.6 0 0 1-.3-.52V6.847a.6.6 0 0 1 .3-.52z"/><path stroke-linecap="round" d="M17 15H7l5-8z"/><path d="M2.5 6.5L12 7m-9.5-.5L7 15m14.5-8.5L17 15m4.5-8.5L12 7V1m9.5 16.5L17 15M2.5 17.5L7 15m0 0l5 8l5-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HexagonPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h3m3 0h-3m0 0V9m0 3v3m-.3-13.827a.6.6 0 0 1 .6 0l8.926 5.154a.6.6 0 0 1 .3.52v10.307a.6.6 0 0 1-.3.52L12.3 22.826a.6.6 0 0 1-.6 0l-8.926-5.154a.6.6 0 0 1-.3-.52V6.847a.6.6 0 0 1 .3-.52z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistoricShield(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 15.528V2.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v12.928a4 4 0 0 1-2.211 3.578l-5.52 2.76a.6.6 0 0 1-.537 0l-5.52-2.76A4 4 0 0 1 4 15.528"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistoricShieldAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.732 21.866l-5.52-2.76A4 4 0 0 1 4 15.528V2.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v12.928a4 4 0 0 1-2.211 3.578l-5.52 2.76a.6.6 0 0 1-.537 0M12 10V2m-8 8h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8m-2 3v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 18v-3a2 2 0 0 1 2-2v0a2 2 0 0 1 2 2v3M2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8"/><path d="M20 11v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltSlim(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17v-4M2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8m-2 3v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltSlimHoriz(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 16h4M2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8m-2 3v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeHospital(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 9.5L12 4l9 5.5M19 13v6.4a.6.6 0 0 1-.6.6H5.6a.6.6 0 0 1-.6-.6V13"/><path d="M13.667 17h-3.334v-2.333H8v-3.334h2.333V9h3.334v2.333H16v3.334h-2.333z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSale(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 9.846c-1-.923-3.667-1.23-3.667.616S14 11.385 14 13.23s-3 1.846-4 .615m2 .857V16m0-6.887V8M2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8"/><path d="M20 11v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSecure(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 12h.4a.6.6 0 0 1 .6.6v2.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6v-2.8a.6.6 0 0 1 .6-.6h.4m4 0v-2c0-.667-.4-2-2-2s-2 1.333-2 2v2m4 0h-4"/><path d="m2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8m-2 3v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeShield(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12.502 9.13l2.049.531c.264.069.45.309.441.582C14.826 15.232 12 16 12 16s-2.826-.768-2.992-5.757a.584.584 0 0 1 .441-.582l2.05-.53a2 2 0 0 1 1.003 0M2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8"/><path d="M20 11v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSimple(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 21H7a4 4 0 0 1-4-4v-6.292a4 4 0 0 1 1.927-3.421l5-3.03a4 4 0 0 1 4.146 0l5 3.03A4 4 0 0 1 21 10.707V17a4 4 0 0 1-4 4m-8-4h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSimpleDoor(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 21H7a4 4 0 0 1-4-4v-6.292a4 4 0 0 1 1.927-3.421l5-3.03a4 4 0 0 1 4.146 0l5 3.03A4 4 0 0 1 21 10.707V17a4 4 0 0 1-4 4h-2m-6 0v-4a3 3 0 0 1 3-3v0a3 3 0 0 1 3 3v4m-6 0h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTable(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 7v10M1 7h22M4 10h16m0-3v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTemperatureIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m2 8l9.732-4.866a.6.6 0 0 1 .536 0L22 8m-2 3v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/><path d="M12 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0-10.5V14m0-2h2m-2-3h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTemperatureOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 8L4.311 3.156a.6.6 0 0 0-.6.037L2.5 4m9.5 7v8a2 2 0 0 1-2 2H7m0 0H3.6a.6.6 0 0 1-.6-.6v-4.8a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6zm12-3a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0-10.5V14m0-2h2m-2-3h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeUser(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.5 9.5L12 4l9.5 5.5"/><path d="M7 21v-1a5 5 0 0 1 10 0v1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizDistributionLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 17V7m0 10h-5.4a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6H19m0 10v3m0-13V4M9 17V7m0 10H5.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6H9m0 10v3M9 7V4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizDistributionLeftSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path fill="currentColor" d="M19 7v10h-5.4a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 17V7m0 10h-5.4a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6H19m0 10v3m0-13V4"/><path fill="currentColor" d="M9 7v10H5.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 17V7m0 10H5.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6H9m0 10v3M9 7V4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizDistributionRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 17V7m0 10h5.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H5m0 10v3M5 7V4m10 13V7m0 10h3.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H15m0 10v3m0-13V4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizDistributionRightSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path fill="currentColor" d="M5 7v10h5.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 17V7m0 10h5.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H5m0 10v3M5 7V4"/><path fill="currentColor" d="M15 7v10h3.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 17V7m0 10h3.4a.6.6 0 0 0 .6-.6V7.6a.6.6 0 0 0-.6-.6H15m0 10v3m0-13V4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalMerge(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 12h-8m0 0l3.5-3.5M14 12l3.5 3.5M2 12h8m0 0L6.5 8.5M10 12l-3.5 3.5M10 21V3m4 18V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalSplit(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12H2m0 0l3.5-3.5M2 12l3.5 3.5M14 12h8m0 0l-3.5-3.5M22 12l-3.5 3.5M10 21V3m4 18V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M6.4 8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 0 .6.6h1.8a.6.6 0 0 1 .6.6v11.8a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V8.6a.6.6 0 0 1 .6-.6z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.992 8h2m2 0h-2m0 0V6m0 2v2M16 17.01l.01-.011M16 13.01l.01-.011M12 13.01l.01-.011M8 13.01l.01-.011M8 17.01l.01-.011m3.99.011l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10M8 12h8m-8 0V7m0 5v5m8-5v5m0-5V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25M8.75 7a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0v-4.25h6.5V17a.75.75 0 0 0 1.5 0V7a.75.75 0 0 0-1.5 0v4.25h-6.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HotAirBalloon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" d="M4 9.5c0 4.571 5.714 8 5.714 8h4.572S20 14.071 20 9.5s-3.582-8-8-8s-8 3.429-8 8"/><path stroke-linejoin="round" d="M9 2c-3 6 1 15.5 1 15.5M14.884 2c3 6-1 15.5-1 15.5"/><path stroke-linecap="round" d="M13.4 23h-2.8a.6.6 0 0 1-.6-.6v-1.8a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 1-.6.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12a7 7 0 0 0 7-7H5a7 7 0 0 0 7 7m0 0a7 7 0 0 1 7 7H5a7 7 0 0 1 7-7M5 2h14M5 22h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseRooms(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11 19v2m0-9v4m5-4v4h-2m7-4H8m-3 0H3m0-7l9-2l9 2"/><path d="M21 8.6v11.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V8.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4 3l1.778 17.09L12 22l6.222-1.91L20 3z"/><path d="M17 7H7.5l.5 4.5h8l-.5 5.5l-3.5 1l-3.5-1l-.25-2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IceCream(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 18v2a2 2 0 1 1-4 0v-2m-5-6h14"/><path d="M7 18a2 2 0 0 1-2-2V9a7 7 0 1 1 14 0v7a2 2 0 0 1-2 2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IceCreamSolidSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.75 20a1.25 1.25 0 1 0 2.5 0v-2h1.5v2a2.75 2.75 0 1 1-5.5 0v-2h1.5z"/><path d="M4.25 16a2.75 2.75 0 0 0 2.751 2.75H17A2.75 2.75 0 0 0 19.75 16V9a7.75 7.75 0 0 0-15.5 0zm13-3.25a.75.75 0 0 0 0-1.5H6.75a.75.75 0 0 0 0 1.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iconoir(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 13v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6M12 3v12m0 0l-3.5-3.5M12 15l3.5-3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inclination(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 19H3.41a.6.6 0 0 1-.431-1.016L16.444 4"/><path d="M20 16c-.5-3.5-1-5-3-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Industry(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 10c0-1-1-2-3-2h-1a3 3 0 0 1-3-3V2m7 19h3v-9h-3v4.5m0 4.5v-4.5m0 4.5H3v-4l3.5-3l4 2.5l4-2.5l3.5 2.5m3-6.5c0-6-4-6-4-6s4 .5 4-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinite(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 9l-.25.375M10 9a5 5 0 1 0 0 6l.334-.5M10 9l4 6a5 5 0 1 0 0-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 11.5v5m0-8.99l.01-.011M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12M12 10.75a.75.75 0 0 1 .75.75v5a.75.75 0 0 1-1.5 0v-5a.75.75 0 0 1 .75-.75M12.568 8a.75.75 0 1 0-1.115-1.003l-.01.011a.75.75 0 1 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputField(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2m1 2.5h1.5m1.5 0H6.5m0 0v7m0 0H5m1.5 0H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputOutput(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-miterlimit="1.5" d="M14 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14"/><path d="M3 19V5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputSearch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12v-2a5 5 0 0 0-5-5H8a5 5 0 0 0-5 5v0a5 5 0 0 0 5 5h4m8.124 4.119a3 3 0 1 0-4.248-4.237a3 3 0 0 0 4.248 4.237m0 0L22 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path d="M3 16V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17.5 6.51l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Internet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 2.05S16 6 16 12m-5 9.95S8 18 8 12c0-6 3-9.95 3-9.95M2.63 15.5H12m-9.37-7h18.74"/><path d="M21.879 17.917c.494.304.463 1.043-.045 1.101l-2.567.291l-1.151 2.312c-.228.459-.933.234-1.05-.334l-1.255-6.116c-.099-.48.333-.782.75-.525z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Intersect(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 13.5v3M13.5 21h3m0-12H9.6a.6.6 0 0 0-.6.6v6.9m1.5 4.5h-.9a.6.6 0 0 1-.6-.6v-.9m12 0v.9a.6.6 0 0 1-.6.6h-.9m0-12h.9a.6.6 0 0 1 .6.6v.9m-18 0v-3M7.5 3h3"/><path d="M7.5 15h6.9a.6.6 0 0 0 .6-.6V7.5M4.5 15h-.9a.6.6 0 0 1-.6-.6v-.9m0-9v-.9a.6.6 0 0 1 .6-.6h.9m9 0h.9a.6.6 0 0 1 .6.6v.9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntersectAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.01 3l-.01.011M11.01 3l-.01.011M7.01 3L7 3.011M3.01 3L3 3.011M3.01 7L3 7.011M3.01 11l-.01.011M3.01 15l-.01.011m6 5.999l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M21 17.01l.01-.011M21 13.01l.01-.011M21 9.01l.01-.011M9 17v-7a1 1 0 0 1 1-1h7"/><path d="M15 7v7a1 1 0 0 1-1 1H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IosSettings(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12m6-6h-6M9 6.803L12 12m0 0l-3 5.197"/><path stroke-dasharray="1 3" d="M12 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IpAddressTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v6M9 9v6m3-3h2.5a1.5 1.5 0 0 0 1.5-1.5v0A1.5 1.5 0 0 0 14.5 9H12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IrisScan(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 3H3v3m9 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M21 12c-1.889 2.991-5.282 6-9 6s-7.111-3.009-9-6c2.299-2.842 4.992-6 9-6s6.701 3.158 9 6m-3-9h3v3M6 21H3v-3m15 3h3v-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 5h3m3 0h-3m0 0l-4 14m0 0H7m3 0h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 6h2m2 0h-2m0 0l-4 12m0 0H8m2 0h2"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zM16 6.75h-1.46l-3.5 10.5H12a.75.75 0 0 1 0 1.5H8a.75.75 0 0 1 0-1.5h1.46l3.5-10.5H12a.75.75 0 0 1 0-1.5h4a.75.75 0 0 1 0 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jellyfish(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2c4.97 0 9 4.104 9 9.167c0 .068 0 .136-.002.204c-.02.954-.865 1.629-1.819 1.629H4.821c-.954 0-1.798-.675-1.819-1.629A9.52 9.52 0 0 1 3 11.167C3 6.104 7.03 2 12 2M6 13l1 1.125c.57.642.57 1.608 0 2.25v0a1.693 1.693 0 0 0 0 2.25v0c.57.642.57 1.608 0 2.25L6 22m5-9l1 1.125c.57.642.57 1.608 0 2.25v0a1.693 1.693 0 0 0 0 2.25v0c.57.642.57 1.608 0 2.25L11 22m5-9l1 1.125c.57.642.57 1.608 0 2.25v0a1.693 1.693 0 0 0 0 2.25v0c.57.642.57 1.608 0 2.25L16 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Journal(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 6h12M6 10h12m-5 4h5m-5 4h5M2 21.4V2.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v18.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6"/><path d="M6 18v-4h3v4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JournalPage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 6h8m-8 4h12m-5 4h5m-5 4h5M2 21.4V2.6a.6.6 0 0 1 .6-.6h15.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 22 5.75V21.4a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6"/><path d="M6 18v-4h3v4zM18 2v3.4a.6.6 0 0 0 .6.6H22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JpegFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M7 15v-3m0 0V9h3v3zm9-3h-3v6h3m6-6h-3v6h3v-2.4M4 9v4.2C4 15 2 15 2 15m11-3h2"/><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JpgFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M10 15v-3m0 0V9h3v3zm9-3h-3v6h3v-2.4M7 9v4.2C7 15 5 15 5 15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KanbanBoard(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3.6v16.8a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6M6 6v10m4-10v3m4-3v7m4-7v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0m0 0h12v3m-4-3v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyBack(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 12a4 4 0 1 0 8 0a4 4 0 0 0-8 0m0 0H2v3m4-3v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyCommand(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6v12m6-12v12M9 6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3M9 18a3 3 0 1 1-3-3h12a3 3 0 1 1-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.992 18h6m-8.58-7.657a4 4 0 1 0 5.657-5.657a4 4 0 0 0-5.657 5.657m0 0l-8.485 8.485l2.121 2.122M6.755 16l2.122 2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.992 18h3m3 0h-3m0 0v-3m0 3v3m-5.58-10.657a4 4 0 1 0 5.657-5.657a4 4 0 0 0-5.657 5.657m0 0l-8.485 8.485l2.121 2.122M6.755 16l2.122 2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.87 20.121L17.993 18m2.121-2.121L17.993 18m0 0l-2.122-2.121M17.992 18l2.121 2.121m-7.701-9.778a4 4 0 1 0 5.657-5.657a4 4 0 0 0-5.657 5.657m0 0l-8.485 8.485l2.121 2.122M6.755 16l2.122 2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyframe(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.777 13.345l-7.297 8.027a2 2 0 0 1-2.96 0l-7.297-8.027a2 2 0 0 1 0-2.69l7.297-8.027a2 2 0 0 1 2.96 0l7.297 8.027a2 2 0 0 1 0 2.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframeAlignCenter(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.68 12.384l-4.22 5.063a.6.6 0 0 1-.92 0l-4.22-5.063a.6.6 0 0 1 0-.768l4.22-5.063a.6.6 0 0 1 .92 0l4.22 5.063a.6.6 0 0 1 0 .768M12 22v-2m0-16V2M4 12H2m20 0h-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframeAlignHorizontal(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.68 12.384l-4.22 5.063a.6.6 0 0 1-.92 0l-4.22-5.063a.6.6 0 0 1 0-.768l4.22-5.063a.6.6 0 0 1 .92 0l4.22 5.063a.6.6 0 0 1 0 .768M4 12H2m20 0h-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframeAlignVertical(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.68 12.384l-4.22 5.063a.6.6 0 0 1-.92 0l-4.22-5.063a.6.6 0 0 1 0-.768l4.22-5.063a.6.6 0 0 1 .92 0l4.22 5.063a.6.6 0 0 1 0 .768M12 22v-2m0-16V2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframeMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 5h6m-5.181 9.329l-5.324 5.99a2 2 0 0 1-2.99 0l-5.324-5.99a2 2 0 0 1 0-2.658l5.324-5.99a2 2 0 0 1 2.99 0l5.324 5.99a2 2 0 0 1 0 2.658"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframeMinusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m5.777 1.345l-7.297 8.027a2 2 0 0 1-2.96 0l-7.297-8.027a2 2 0 0 1 0-2.69l7.297-8.027a2 2 0 0 1 2.96 0l7.297 8.027a2 2 0 0 1 0 2.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframePlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 5h3m3 0h-3m0 0V2m0 3v3m-2.181 6.329l-5.324 5.99a2 2 0 0 1-2.99 0l-5.324-5.99a2 2 0 0 1 0-2.658l5.324-5.99a2 2 0 0 1 2.99 0l5.324 5.99a2 2 0 0 1 0 2.658"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframePlusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.777 13.345l-7.297 8.027a2 2 0 0 1-2.96 0l-7.297-8.027a2 2 0 0 1 0-2.69l7.297-8.027a2 2 0 0 1 2.96 0l7.297 8.027a2 2 0 0 1 0 2.69M9 12h3m3 0h-3m0 0V9m0 3v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframePosition(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.68 9.384l-4.22 5.063a.6.6 0 0 1-.92 0L7.32 9.384a.6.6 0 0 1 0-.768l4.22-5.063a.6.6 0 0 1 .92 0l4.22 5.063a.6.6 0 0 1 0 .768M3 20h9m9 0h-9m0 0v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyframes(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.848 13.317L9.505 18.28a2 2 0 0 1-3.01 0l-4.343-4.963a2 2 0 0 1 0-2.634L6.495 5.72a2 2 0 0 1 3.01 0l4.343 4.963a2 2 0 0 1 0 2.634"/><path d="m13 19l4.884-5.698a2 2 0 0 0 0-2.604L13 5"/><path d="m17 19l4.884-5.698a2 2 0 0 0 0-2.604L17 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframesCouple(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.819 13.329l-5.324 5.99a2 2 0 0 1-2.99 0l-5.324-5.99a2 2 0 0 1 0-2.658l5.324-5.99a2 2 0 0 1 2.99 0l5.324 5.99a2 2 0 0 1 0 2.658"/><path d="m12 6.375l1.505-1.693a2 2 0 0 1 2.99 0l5.324 5.99a2 2 0 0 1 0 2.657l-5.324 5.99a2 2 0 0 1-2.99 0L12 17.624"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframesMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12h6M6.25 6l.245-.28a2 2 0 0 1 3.01 0l4.343 4.963a2 2 0 0 1 0 2.634L9.505 18.28a2 2 0 0 1-3.01 0L6.25 18"/><path d="m13 19l4.884-5.698a2 2 0 0 0 0-2.604L13 5"/><path d="m17 19l4.884-5.698a2 2 0 0 0 0-2.604L17 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyframesPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12h3m3 0H5m0 0V9m0 3v3m1.25-9l.245-.28a2 2 0 0 1 3.01 0l4.343 4.963a2 2 0 0 1 0 2.634L9.505 18.28a2 2 0 0 1-3.01 0L6.25 18"/><path d="m13 19l4.884-5.698a2 2 0 0 0 0-2.604L13 5"/><path d="m17 19l4.884-5.698a2 2 0 0 0 0-2.604L17 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 17.4V6.6a.6.6 0 0 1 .6-.6h13.079c.2 0 .388.1.5.267l3.6 5.4a.6.6 0 0 1 0 .666l-3.6 5.4a.6.6 0 0 1-.5.267H3.6a.6.6 0 0 1-.6-.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-width="1.5" d="M3 17.4V6.6a.6.6 0 0 1 .6-.6h13.079c.2 0 .388.1.5.267l3.6 5.4a.6.6 0 0 1 0 .666l-3.6 5.4a.6.6 0 0 1-.5.267H3.6a.6.6 0 0 1-.6-.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lamp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 21h3m3 0h-3m0 0V11m0-4v4m0 0H6l3-8h6l3 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12"/><path d="M13 2.05S16 6 16 12c0 6-3 9.95-3 9.95m-2 0S8 18 8 12c0-6 3-9.95 3-9.95M2.63 15.5h18.74m-18.74-7h18.74"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.2 14.222V4a2 2 0 0 1 2-2h13.6a2 2 0 0 1 2 2v10.222m-17.6 0h17.6m-17.6 0l-1.48 5.234A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544l-1.48-5.234"/><path stroke-linecap="round" stroke-linejoin="round" d="M11 19h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopCharging(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.2 14.222V4a2 2 0 0 1 2-2h13.6a2 2 0 0 1 2 2v10.222m-17.6 0h17.6m-17.6 0l-1.48 5.234A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544l-1.48-5.234"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.667 5L10 8h4l-1.667 3M11 19h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopDevMode(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.2 14.222V4a2 2 0 0 1 2-2h13.6a2 2 0 0 1 2 2v10.222m-17.6 0h17.6m-17.6 0l-1.48 5.234A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544l-1.48-5.234"/><path stroke-linecap="round" stroke-linejoin="round" d="M11 19h2m1-13l2 2l-2 2m-4-4L8 8l2 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopFix(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.8 14.222H3.654a.6.6 0 0 0-.578.437L1.72 19.456A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544zm0 0v-6.11m-17.6 6.11V4a2 2 0 0 1 2-2H12m-1 17h2m4.657-14.172l-2.829 2.829m5.657-2.829A2 2 0 1 1 17.657 2m-2.829 8.485A2 2 0 0 0 12 7.657"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.2 14.222V4a2 2 0 0 1 2-2h13.6a2 2 0 0 1 2 2v10.222m-17.6 0h17.6m-17.6 0l-1.48 5.234A2 2 0 0 0 3.644 22h16.712a2 2 0 0 0 1.924-2.544l-1.48-5.234"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 5v3m0 3.01l.01-.011M11 19h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3.6 3h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6Zm6.15 6.75V21M3 9.75h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M20.4 3H3.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6Zm-6.15 6.75V21M21 9.75H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaderboard(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 19H9V8.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v5.9zm0-14H9m11.4 14H15v-3.9a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v3.3a.6.6 0 0 1-.6.6M9 19v-5.9a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v5.3a.6.6 0 0 0 .6.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeaderboardStar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 21H9v-8.4a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6zm5.4 0H15v-2.9a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6M9 21v-4.9a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v4.3a.6.6 0 0 0 .6.6zm1.806-15.887l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.176.568l-1.47 1.5l.347 2.118c.044.272-.228.48-.462.351l-1.818-1l-1.818 1c-.233.128-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.175-.568z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 21s.5-4.5 4-8.5"/><path d="m19.13 4.242l.594 6.175c.374 3.886-2.54 7.346-6.425 7.72c-3.813.367-7.267-2.42-7.634-6.233a6.936 6.936 0 0 1 6.239-7.569l6.571-.632a.6.6 0 0 1 .655.54"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Learning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.818 22v-2.857C6.52 16.166 3 14.572 3 10c0-4.57 2.727-8.056 8.182-8c3.927.042 7.636 2.286 7.636 6.858L21 12.286c0 2.286-2.182 2.286-2.182 2.286s.546 5.714-4.364 5.714V22"/><path d="M11 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path stroke-dasharray=".3 2" d="M11 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lens(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M17.197 9c-.1-.172-.207-.34-.323-.5m.937 5a6.01 6.01 0 0 1-4.311 4.311"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LensPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.992 6h3m3 0h-3m0 0V3m0 3v3m-3.88 4.5C2.835 18.311 6.987 22 12 22c5.523 0 10-4.477 10-10c0-5.013-3.689-9.165-8.5-9.888"/><path d="M17.197 9c-.1-.172-.207-.34-.323-.5m.937 5a6.01 6.01 0 0 1-4.311 4.311"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lifebelt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M8 12a4 4 0 1 0 8 0a4 4 0 0 0-8 0m1.235 2.89L5 19m9.765-4.11L19 19m-4.235-9.89L19 5M9.235 9.11L5 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulb(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 18h6m-5 3h4m-5-6c.001-2-.499-2.5-1.5-3.5c-1-1-1.476-2.013-1.5-3.5c-.047-3.05 2-5 6-5c4.001 0 6.049 1.95 6 5c-.023 1.487-.5 2.5-1.5 3.5c-.999 1-1.499 1.5-1.5 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 18h6m-5 3h4m2.5-9.5c1-1 1.477-2.013 1.5-3.5c.048-3.05-2-5-6-5c-1.168 0-2.169.166-3 .477M9 15c0-2-.5-2.5-1.5-3.5S6.023 9.487 6 8a5.618 5.618 0 0 1 .168-1.5M3 3l18 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbOn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 2l-1 1M3 2l1 1m17 13l-1-1M3 16l1-1m5 3h6m-5 3h4M12 3C8 3 5.952 4.95 6 8c.023 1.487.5 2.5 1.5 3.5S9 13 9 15h6c0-2 .5-2.5 1.5-3.5h0c1-1 1.477-2.013 1.5-3.5c.048-3.05-2-5-6-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineSpace(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 6h10m-10 6h10m-10 6h10M5 19V5m0 14l-2-2.5M5 19l2-2.5M5 5L3 7m2-2l2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linear(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20L21 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 11.998C14 9.506 11.683 7 8.857 7H7.143C4.303 7 2 9.238 2 11.998c0 2.378 1.71 4.368 4 4.873a5.3 5.3 0 0 0 1.143.124"/><path d="M10 11.998c0 2.491 2.317 4.997 5.143 4.997h1.714c2.84 0 5.143-2.237 5.143-4.997c0-2.379-1.71-4.37-4-4.874A5.304 5.304 0 0 0 16.857 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkSlash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.143 16.995c-.393 0-.775-.043-1.143-.123c-2.29-.506-4-2.496-4-4.874c0-2.714 2.226-4.923 5-4.996m6.318 2.632A5.517 5.517 0 0 0 11 7.5"/><path d="M16.857 7c.393 0 .775.043 1.143.124c2.29.505 4 2.495 4 4.874c0 2.76-2.302 4.997-5.143 4.997h-1.714c-2.826 0-5.143-2.506-5.143-4.997c0 0 0-.998.5-1.498M3 3l18 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 11.998C14 9.506 11.683 7 8.857 7H7.143C4.303 7 2 9.238 2 11.998c0 2.378 1.71 4.368 4 4.873a5.3 5.3 0 0 0 1.143.124M16.857 7c.393 0 .775.043 1.143.124c2.29.505 4 2.495 4 4.874a4.92 4.92 0 0 1-1.634 3.653"/><path d="M10 11.998c0 2.491 2.317 4.997 5.143 4.997M18 22.243l2.121-2.122m0 0L22.243 18m-2.122 2.121L18 18m2.121 2.121l2.122 2.122"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5M7 17v-7"/><path d="M11 17v-3.25M11 10v3.75m0 0c0-3.75 6-3.75 6 0V17M7 7.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linux(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M2.5 20c1 0 2-.8 2-2v-7c0-3.5 3.1-7 7.5-7m9.5 16c-1 0-2-.8-2-2v-7c0-3.5-3.1-7-7.5-7"/><path stroke-linejoin="round" d="M12 19c2.761 0 5-1.12 5-2.5S14.761 14 12 14s-5 1.12-5 2.5S9.239 19 12 19"/><path stroke-linejoin="round" d="M7.75 15c-.463-.635-.75-1.52-.75-2.5C7 10.567 8.12 9 9.5 9s2.5 1.567 2.5 3.5c0 .455-.062.89-.175 1.29M16.25 15c.463-.635.75-1.52.75-2.5c0-1.933-1.12-3.5-2.5-3.5S12 10.567 12 12.5c0 .455.062.89.175 1.29M9.5 12v2m5-2v2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 6h12M4 6.01l.01-.011M4 12.01l.01-.011M4 18.01l.01-.011M8 12h12M8 18h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListSelect(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6h11M5 6.01l.01-.011M5 12.01l.01-.011M3.8 17.8l.8.8l2-2M9 12h11M9 18h11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LitecoinCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M10.5 7v9.4a.6.6 0 0 0 .6.6h4.4"/><path stroke-linejoin="round" d="m8.5 13l4.5-2m-1 11C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LitecoinCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12m9.25-5.75a.75.75 0 0 1 .75.75v3.957l1.445-.642a.75.75 0 0 1 .61 1.37l-2.055.913v3.652h4.25a.75.75 0 0 1 0 1.5h-4.4a1.35 1.35 0 0 1-1.35-1.35v-3.135l-.945.42a.75.75 0 0 1-.61-1.37l1.555-.691V7a.75.75 0 0 1 .75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LitecoinRotateOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.168 8A10.003 10.003 0 0 0 12 2c-5.185 0-9.45 3.947-9.95 9"/><path stroke-linejoin="round" d="M18 8h3.4a.6.6 0 0 0 .6-.6V4M2.881 16c1.544 3.532 5.068 6 9.168 6c5.186 0 9.45-3.947 9.951-9"/><path stroke-linejoin="round" d="M6.05 16h-3.4a.6.6 0 0 0-.6.6V20"/><path d="M10.5 7v9.4a.6.6 0 0 0 .6.6h4.4"/><path stroke-linejoin="round" d="m8.5 13l4.5-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 12h1.4a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6v-6.8a.6.6 0 0 1 .6-.6H8m8 0V8c0-1.333-.8-4-4-4S8 6.667 8 8v4m8 0H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSlash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.5 12H6.6a.6.6 0 0 0-.6.6v6.8a.6.6 0 0 0 .6.6h10.8a.6.6 0 0 0 .6-.6v-.9M16 12V8c0-1.333-.8-4-4-4c-.747 0-1.363.145-1.869.385M16 12h1.4a.6.6 0 0 1 .6.6v.4M8 8v4M3 3l18 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M14.667 12h.733a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6v-3.8a.6.6 0 0 1 .6-.6h.733m5.334 0V9.5c0-.833-.534-2.5-2.667-2.5S9.333 8.667 9.333 9.5V12m5.334 0H9.333"/><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoftThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 17c-9 0-11 6-20 6M22 1C13 1 11 7 2 7m10 9.5v-9m0 9l2.5-2.5M12 16.5L9.5 14M12 7.5l2.5 2.5M12 7.5L9.5 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 12h-7m0 0l3 3m-3-3l3-3m4-3V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogNoAccess(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.857 9.2a4 4 0 0 0-5.713 5.6m5.713-5.6a4 4 0 0 1-5.713 5.6m5.713-5.6l-5.714 5.6"/><path d="M19 6V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12h7m0 0l-3 3m3-3l-3-3m3-3V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowDownLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.25 19.25l-3.5-3.5l3.5-3.5"/><path d="M6.75 15.75h6a4 4 0 0 0 4-4v-7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowDownRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.25 19.25l3.5-3.5l-3.5-3.5"/><path d="M16.75 15.75h-6a4 4 0 0 1-4-4v-7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowLeftDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 13.5L8 17l3.5-3.5"/><path d="M8 17v-6a4 4 0 0 1 4-4h7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowLeftUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 10.5L8 7l3.5 3.5"/><path d="M8 7v6a4 4 0 0 0 4 4h7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowRightDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 13.5L15.5 17L12 13.5"/><path d="M15.5 17v-6a4 4 0 0 0-4-4h-7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowRightUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 10.5L15.5 7L12 10.5"/><path d="M15.5 7v6a4 4 0 0 1-4 4h-7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowRightUpOne(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 7v6a4 4 0 0 1-4 4h-7m11-10l3.5 3.5M15.5 7L12 10.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowUpLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.25 4.75l-3.5 3.5l3.5 3.5"/><path d="M6.75 8.25h6a4 4 0 0 1 4 4v7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrowUpRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.25 4.75l3.5 3.5l-3.5 3.5"/><path d="M16.75 8.25h-6a4 4 0 0 0-4 4v7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LotOfCash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 18v-8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2m14.5-3.99l.01-.011M7.5 14.01l.01-.011"/><path d="M4 16H3a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2m-6 8a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lullaby(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M3 11.507a9.493 9.493 0 0 0 18 4.219c-8.507 0-12.726-4.22-12.726-12.726A9.494 9.494 0 0 0 3 11.507"/><path d="M19 9.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0V3.6a.6.6 0 0 1 .6-.6H21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MacControlKey(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m8 14l4-4l4 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MacDock(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8 17a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/><path d="M21 21H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1M2 17.5l2-1m18 1l-2-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MacOptionKey(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 10h3m0 4h-5l-1.667-4H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MacOsWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 17.714V6.286C2 5.023 2.995 4 4.222 4h15.556C21.005 4 22 5.023 22 6.286v11.428C22 18.977 21.005 20 19.778 20H4.222C2.995 20 2 18.977 2 17.714Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 7.01L5.01 7M7 7.01L7.01 7M9 7.01L9.01 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagicWand(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="m3 21l10-10m5-5l-2.5 2.5"/><path d="m9.5 2l.945 2.555L13 5.5l-2.555.945L9.5 9l-.945-2.555L6 5.5l2.555-.945zm9.5 8l.54 1.46L21 12l-1.46.54L19 14l-.54-1.46L17 12l1.46-.54z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4v8.296C4 16.551 7.582 20 12 20s8-3.45 8-7.704V4"/><path d="M4 4h5.63v6.818C9.63 12.023 10.69 13 12 13s2.37-.977 2.37-2.182V4H20M9 8H4m16 0h-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetEnergy(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M5 9v6.74C5 19.199 8.134 22 12 22s7-2.802 7-6.26V9M5 9h3m8 0h3"/><path stroke-linecap="round" d="M14.074 11.5v3.56c0 1.072-.928 1.94-2.074 1.94s-2.074-.868-2.074-1.94V11.5"/><path d="M10 13H5m14 0h-5"/><path stroke-linecap="round" d="M11.667 2L10 5h4l-1.667 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m7 9l5 3.5L17 9"/><path d="M2 17V7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m5 9l4.5 3L14 9"/><path d="M17 19H3a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v2"/><path stroke-linejoin="round" d="M23 14h-6m0 0l3-3m-3 3l3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailOpen(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m7 12l5 3.5l5-3.5"/><path d="M2 20V9.132a2 2 0 0 1 .971-1.715l8-4.8a2 2 0 0 1 2.058 0l8 4.8A2 2 0 0 1 22 9.132V20a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m5 9l4.5 3L14 9"/><path d="M17 19H3a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v2"/><path stroke-linejoin="round" d="M17 14h6m0 0l-3-3m3 3l-3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.232 9.747a6 6 0 1 0-8.465 8.506a6 6 0 0 0 8.465-8.506m0 0L20 4m0 0h-4m4 0v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 19l-5.21 1.737a.6.6 0 0 1-.79-.57V5.433a.6.6 0 0 1 .41-.569L9 3m0 16l6 2m-6-2V3m6 18l5.59-1.863a.6.6 0 0 0 .41-.57V3.832a.6.6 0 0 0-.79-.569L15 5m0 16V5m0 0L9 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPin(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20 10c0 4.418-8 12-8 12s-8-7.582-8-12a8 8 0 1 1 16 0Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 9.2C16 13.177 9 20 9 20S2 13.177 2 9.2C2 5.224 5.134 2 9 2s7 3.224 7 7.2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 19h6"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 9.2C16 13.177 9 20 9 20S2 13.177 2 9.2C2 5.224 5.134 2 9 2s7 3.224 7 7.2Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 19h3m3 0h-3m0 0v-3m0 3v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 9.2C16 13.177 9 20 9 20S2 13.177 2 9.2C2 5.224 5.134 2 9 2s7 3.224 7 7.2Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.88 21.121L19 19m2.122-2.121L19 19m0 0l-2.12-2.121M19 19l2.122 2.121"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 19l-5.21 1.737a.6.6 0 0 1-.79-.57V5.433a.6.6 0 0 1 .41-.569L9 3m0 16l5.21 1.737a.6.6 0 0 0 .79-.57V5.433a.6.6 0 0 0-.41-.569L9 3m0 16V3m6 2l5.21-1.737a.6.6 0 0 1 .79.57V15m-3.879 7.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L17.12 18.12m2.122 2.122l2.121 2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapsArrow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.685 18.783l7.88-14.008a.5.5 0 0 1 .87 0l7.88 14.008a.5.5 0 0 1-.617.71l-7.517-2.922a.5.5 0 0 0-.362 0l-7.517 2.923a.5.5 0 0 1-.617-.711"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapsArrowDiagonal(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.031 8.917l15.477-4.334a.5.5 0 0 1 .616.617l-4.333 15.476a.5.5 0 0 1-.94.067l-3.248-7.382a.5.5 0 0 0-.256-.257L3.965 9.856a.5.5 0 0 1 .066-.94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapsArrowXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 17.278l-1.819-.707a.5.5 0 0 0-.362 0l-7.517 2.923a.5.5 0 0 1-.617-.711l7.88-14.008a.5.5 0 0 1 .87 0l6.065 10.78m-1.379 6.809l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L17.12 18.12m2.122 2.122l2.121 2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapsGoStraight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.365 19.787l7.303-6.492a.5.5 0 0 1 .664 0l7.303 6.492c.38.338.072.962-.427.864l-7.113-1.382a.498.498 0 0 0-.19 0l-7.113 1.383c-.499.097-.808-.527-.427-.865M12 10.5V4m0 0L8 6.5M12 4l4 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapsTurnBack(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.365 19.787l7.303-6.492a.5.5 0 0 1 .664 0l7.303 6.492c.38.338.072.962-.427.864l-7.113-1.382a.498.498 0 0 0-.19 0l-7.113 1.383c-.499.097-.808-.527-.427-.865M5.5 11V6v0s0-3.5 3-3.5C12 2.5 12 6 12 6v4.5"/><path d="M9 7.5L5.5 11L2 7.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapsTurnLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.365 19.787l7.303-6.492a.5.5 0 0 1 .664 0l7.303 6.492c.38.338.072.962-.427.864l-7.113-1.382a.498.498 0 0 0-.19 0l-7.113 1.383c-.499.097-.808-.527-.427-.865M5 6.5h3s0 0 0 0s4 0 4 4"/><path d="M8.5 9L5 6.5L8.5 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapsTurnRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.365 19.787l7.303-6.492a.5.5 0 0 1 .664 0l7.303 6.492c.38.338.072.962-.427.864l-7.113-1.382a.498.498 0 0 0-.19 0l-7.113 1.383c-.499.097-.808-.527-.427-.865M19 6.5h-3s0 0 0 0s-4 0-4 4"/><path d="M15.5 9L19 6.5L15.5 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaskSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M10 17.659a6 6 0 0 0 4-11.317m-4 11.317a6 6 0 1 1 4-11.317m-4 11.317L14 6.34"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MastercardCard(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 9v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2zm0 0H6"/><path d="M16.5 13.382a1.5 1.5 0 1 1 0 2.236m0-2.236a1.5 1.5 0 1 0 0 2.236"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mastodon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 13.5V9c0-3 5-3 5 0v3m5 1.5V9c0-3-5-3-5 0v3"/><path d="M8 17c7.5 1 13 0 13-4V9c0-5.5-4-6.5-6-6.5H9c-3 0-6.067 1-5.863 6.5c.074 1.987.036 4.385.363 7c1 8 10.5 5.5 12 5v-1.5S7.5 21 8 17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathBook(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114M6 17h14M6 21h14"/><path stroke-linejoin="round" d="M6 21a2 2 0 1 1 0-4"/><path d="M10 10h4"/><path stroke-linejoin="round" d="m12 13.01l.01-.011M12 7.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maximize(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4H4v3m13-3h3v3M7 20H4v-3m13 3h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medal(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.272 10.445L18 2m-8.684 8.632L5 2m7.762 8.048L8.835 2m5.525 0l-1.04 2.5M6 16a6 6 0 1 0 12 0a6 6 0 0 0-12 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedalOneSt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.272 10.445L18 2m-8.684 8.632L5 2m7.762 8.048L8.835 2m5.525 0l-1.04 2.5M6 16a6 6 0 1 0 12 0a6 6 0 0 0-12 0"/><path d="m10.5 15l2-1.5v5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedalOneStSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.667 1.712a.75.75 0 1 1 1.385.576l-1.04 2.5a.75.75 0 1 1-1.385-.576zm3.647-.015a.75.75 0 1 1 1.372.606l-3.435 7.78a6.75 6.75 0 1 1-6.923.252l-4-8a.75.75 0 1 1 1.342-.67l4 7.998a6.718 6.718 0 0 1 1.875-.398L8.16 2.33a.75.75 0 0 1 1.348-.658l3.755 7.697c.204.039.404.087.6.143zM12.05 12.9a.75.75 0 0 1 1.2.6v5a.75.75 0 0 1-1.5 0V15l-.8.6a.75.75 0 0 1-.9-1.2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedalSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.648 1.308a.75.75 0 0 0-.98.404l-1.04 2.5a.75.75 0 0 0 1.384.576l1.04-2.5a.75.75 0 0 0-.404-.98m3.655.006a.75.75 0 0 0-.99.383l-3.449 7.814a6.713 6.713 0 0 0-.6-.143L9.51 1.671a.75.75 0 0 0-1.348.658l3.384 6.936a6.718 6.718 0 0 0-1.875.398l-4-7.998a.75.75 0 1 0-1.341.67l4 8a6.75 6.75 0 1 0 6.922-.252l3.435-7.78a.75.75 0 0 0-.383-.989" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaImage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="m3 16l7-3l11 5m-5-8a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaImageFolder(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 12.6v7.8a.6.6 0 0 1-.6.6h-7.8a.6.6 0 0 1-.6-.6v-7.8a.6.6 0 0 1 .6-.6h7.8a.6.6 0 0 1 .6.6m-2.5 1.91l.01-.011"/><path d="m13 18.2l3.5-1.2l5.5 2M2 10V3.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H21.4a.6.6 0 0 1 .6.6V9M2 10v8.4a.6.6 0 0 0 .6.6H10m-8-9h8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaImageList(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.6v12.8a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6h12.8a.6.6 0 0 1 .6.6"/><path d="M18 4H4.6a.6.6 0 0 0-.6.6V18m3-1.2l5.444-1.8L21 18m-4.5-5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaImagePlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 21H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V13"/><path d="m3 16l7-3l5.5 2.5M16 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4m0 9h3m3 0h-3m0 0v-3m0 3v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaImageXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 16l7-3l4 1.818M16 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4m.879 11.121L19 19m2.121-2.121L19 19m0 0l-2.121-2.121M19 19l2.121 2.121"/><path d="M13 21H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaVideo(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M9.898 8.513a.6.6 0 0 0-.898.52v5.933a.6.6 0 0 0 .898.521l5.19-2.966a.6.6 0 0 0 0-1.042z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaVideoFolder(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 12.6v7.8a.6.6 0 0 1-.6.6h-7.8a.6.6 0 0 1-.6-.6v-7.8a.6.6 0 0 1 .6-.6h7.8a.6.6 0 0 1 .6.6"/><path d="M16.918 14.574a.6.6 0 0 0-.918.508v2.835a.6.6 0 0 0 .918.51l2.268-1.418a.6.6 0 0 0 0-1.018zM2 10V3.6a.6.6 0 0 1 .6-.6h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H21.4a.6.6 0 0 1 .6.6V9M2 10v8.4a.6.6 0 0 0 .6.6H10m-8-9h8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaVideoList(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.6v12.8a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6h12.8a.6.6 0 0 1 .6.6"/><path d="M18 4H4.6a.6.6 0 0 0-.6.6V18m8.909-6.455a.6.6 0 0 0-.909.515v3.88a.6.6 0 0 0 .909.515l3.233-1.94a.6.6 0 0 0 0-1.03z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaVideoPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 21H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V13m-5 6h3m3 0h-3m0 0v-3m0 3v3"/><path d="M9.898 8.513a.6.6 0 0 0-.898.52v5.933a.6.6 0 0 0 .898.521l5.19-2.966a.6.6 0 0 0 0-1.042z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaVideoXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.879 21.121L19 19m2.121-2.121L19 19m0 0l-2.121-2.121M19 19l2.121 2.121M13 21H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V13"/><path d="M9.898 8.513a.6.6 0 0 0-.898.52v5.933a.6.6 0 0 0 .898.521l5.19-2.966a.6.6 0 0 0 0-1.042z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medium(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8m8 0c1.105 0 2-1.79 2-4s-.895-4-2-4s-2 1.79-2 4s.895 4 2 4m5 0c.552 0 1-1.79 1-4s-.448-4-1-4s-1 1.79-1 4s.448 4 1 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M14 14V6m0 8l6.102 3.487a.6.6 0 0 0 .898-.52V3.033a.6.6 0 0 0-.898-.521L14 6m0 8H7a4 4 0 1 1 0-8h7M7.757 19.3L7 14h4l.677 4.74a1.98 1.98 0 0 1-3.92.56Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 5h18M3 12h18M3 19h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuScale(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 5h8m-8 7h13M3 19h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 20.29V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7.961a2 2 0 0 0-1.561.75l-2.331 2.914A.6.6 0 0 1 3 20.29Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAlert(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 7v2m0 4.01l.01-.011"/><path d="M3 20.29V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7.961a2 2 0 0 0-1.561.75l-2.331 2.914A.6.6 0 0 1 3 20.29Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageText(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 12h10M7 8h6"/><path d="M3 20.29V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7.961a2 2 0 0 0-1.561.75l-2.331 2.914A.6.6 0 0 1 3 20.29Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MeterArrowDownRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.5 3.5L7 8m0 0V4m0 4H3m12 8l-3.5-3.5"/><path d="M14.5 9C10.358 9 7 12.283 7 16.333a7.17 7.17 0 0 0 .733 3.165a.925.925 0 0 0 .84.502h11.853a.925.925 0 0 0 .841-.502A7.17 7.17 0 0 0 22 16.333C22 12.283 18.642 9 14.5 9Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Metro(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m14 16.01l.01-.011M10 16.01l.01-.011M22 12v3a5 5 0 0 1-5 5H7a5 5 0 0 1-5-5v-3C2 6.477 6.477 2 12 2s10 4.477 10 10"/><path stroke-linejoin="round" d="M18 12v3a5 5 0 0 1-5 5h-2a5 5 0 0 1-5-5v-3a5 5 0 0 1 5-5h2a5 5 0 0 1 5 5"/><path d="m10.5 20l-2 2.5m5-2.5l2 2.5m1-2.5l2 2.5M7.5 20l-2 2.5"/><path stroke-linejoin="round" d="M11.786 10h.428C13.2 10 14 10.8 14 11.786a.214.214 0 0 1-.214.214h-3.572a.214.214 0 0 1-.214-.214C10 10.8 10.8 10 11.786 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="6" height="12" x="9" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H9m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m15.5 20.5l2 2l5-5"/><rect width="6" height="12" x="5" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M1 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H5m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneCheckSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m15.5 20.5l2 2l5-5"/><rect width="6" height="12" x="5" y="2" fill="currentColor" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M1 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H5m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.992 19h6"/><rect width="6" height="12" x="5" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M1 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H5m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneMinusSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.992 19h6"/><rect width="6" height="12" x="5" y="2" fill="currentColor" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M1 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H5m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneMute(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 3l18 18M9 9a5 5 0 0 0 5 5v0m1-3.5V5a3 3 0 0 0-3-3v0a3 3 0 0 0-3 3v.5"/><path d="M5 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H9m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophonePlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.992 19h3m3 0h-3m0 0v-3m0 3v3"/><rect width="6" height="12" x="5" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M1 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H5m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophonePlusSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.992 19h3m3 0h-3m0 0v-3m0 3v3"/><rect width="6" height="12" x="5" y="2" fill="currentColor" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M1 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H5m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="6" height="12" x="9" y="2" fill="currentColor" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H9m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSpeaking(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="6" height="12" x="9" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 3v2M1 2v4m18-3v2m4-3v4M5 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H9m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSpeakingSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="6" height="12" x="9" y="2" fill="currentColor" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 3v2M1 2v4m18-3v2m4-3v4M5 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H9m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 14v4m0 4.01l.01-.011"/><rect width="6" height="12" x="7" y="2" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H7m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneWarningSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 14v4m0 4.01l.01-.011"/><rect width="6" height="12" x="7" y="2" fill="currentColor" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 10v1a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7v-1m-7 8v4m0 0H7m3 0h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microscope(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 22H7m-2 0h2m0 0v-3m12-3h-9m6-14h-4m0 5c-3 0-5 1-5 4v2m9-8.4v6.8a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6"/><path d="M7 19a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h8m-4 10c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m-4 10a.75.75 0 0 0 0 1.5h8a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusHexagon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6M11.7 1.173a.6.6 0 0 1 .6 0l8.926 5.154a.6.6 0 0 1 .3.52v10.307a.6.6 0 0 1-.3.52L12.3 22.826a.6.6 0 0 1-.6 0l-8.926-5.154a.6.6 0 0 1-.3-.52V6.847a.6.6 0 0 1 .3-.52z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m6-8.4v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareDashed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4H4v3m4 5h8M4 11v2m7-9h2m-2 16h2m7-9v2m-3-9h3v3M7 20H4v-3m13 3h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm5.4 9a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mirror(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 4v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2m0 1l-6 5m6-1l-7.5 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileDevMode(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m12 19.01l.01-.011"/><path d="M18 18v3.4a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6V18M18 6V2.6a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 0-.6.6V6"/><path stroke-linejoin="round" d="M15.5 8.5L19 12l-3.5 3.5m-7-7L5 12l3.5 3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileFingerprint(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 11V6.362c0-.51.1-1 .284-1.454M22 11V7.815m-7.778-5.08A5.507 5.507 0 0 1 17 2c2.28 0 4.203 1.33 4.805 3.15M15 12V9.824M19 12V6.853C19 5.83 18.105 5 17 5s-2 .83-2 1.853v.794M8 17.01l.01-.011M8 5H3.6a.6.6 0 0 0-.6.6v14.8a.6.6 0 0 0 .6.6h8.8a.6.6 0 0 0 .6-.6V16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileVoice(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 17.01l.01-.011M8 5H3.6a.6.6 0 0 0-.6.6v14.8a.6.6 0 0 0 .6.6h8.8a.6.6 0 0 0 .6-.6V16m3-13v10m-3-8v6m9-4v2M10 7v2m9-5v8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ModernTv(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10"/><path d="M2 16.4V3.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ModernTvFourK(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10M13.5 7v4m0 2v-2m0 0l1.37-1.566M17 7l-2.13 2.434m0 0L17 13M9.5 7l-3 4.5H10V13"/><path d="M2 16.4V3.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneySquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 8.5c-.685-.685-1.891-1.161-3-1.191M9 15c.644.86 1.843 1.35 3 1.391m0-9.082c-1.32-.036-2.5.561-2.5 2.191c0 3 5.5 1.5 5.5 4.5c0 1.711-1.464 2.446-3 2.391m0-9.082V5.5m0 10.891V18.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonSat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path d="M7.633 3.067A3.001 3.001 0 1 1 4.017 6.32M22 13.05a3.5 3.5 0 1 0-3 5.914"/><path stroke-linecap="round" stroke-linejoin="round" d="m14.5 8.51l.01-.011M10 17a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreHoriz(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-8 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-8 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreHorizCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M7 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVert(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m0 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m0-16a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVertCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 7.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m0 10a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m0-5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Motorcycle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8m14-4l-3-9l1-1"/><path d="M16 8.5h-4.5l-4.5 3m-1.5 4H12l1-2.5l3.5-4.5m-8 1.5c-2-1.5-5-1.5-7 0M19 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseButtonLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20 10v4a8 8 0 1 1-16 0V9a7 7 0 0 1 7-7h1a8 8 0 0 1 8 8Z"/><path d="M12 2v6.4a.6.6 0 0 1-.6.6H4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseButtonRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 10v4a8 8 0 1 0 16 0V9a7 7 0 0 0-7-7h-1a8 8 0 0 0-8 8Z"/><path d="M12 2v6.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseScrollWheel(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 5l.53-.53a.75.75 0 0 0-1.06 0zm0 8l-.53.53a.75.75 0 0 0 1.06 0zM9.47 6.47a.75.75 0 0 0 1.06 1.06zm4 1.06a.75.75 0 1 0 1.06-1.06zm-2.94 2.94a.75.75 0 1 0-1.06 1.06zm4 1.06a.75.75 0 1 0-1.06-1.06zM3.25 10v4h1.5v-4zm17.5 4v-4h-1.5v4zm-9.5-9v8h1.5V5zm.22-.53l-2 2l1.06 1.06l2-2zm0 1.06l2 2l1.06-1.06l-2-2zm1.06 6.94l-2-2l-1.06 1.06l2 2zm0 1.06l2-2l-1.06-1.06l-2 2zM20.75 10A8.75 8.75 0 0 0 12 1.25v1.5A7.25 7.25 0 0 1 19.25 10zM12 22.75A8.75 8.75 0 0 0 20.75 14h-1.5A7.25 7.25 0 0 1 12 21.25zM3.25 14A8.75 8.75 0 0 0 12 22.75v-1.5A7.25 7.25 0 0 1 4.75 14zm1.5-4A7.25 7.25 0 0 1 12 2.75v-1.5A8.75 8.75 0 0 0 3.25 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Movie(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 8.01l.01-.011M17 8.01l.01-.011M7 12.01l.01-.011m9.99.011l.01-.011M7 16.01l.01-.011m9.99.011l.01-.011M7 2H3.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H7M7 2v2m0-2h10m0 0h3.4a.6.6 0 0 1 .6.6v18.8a.6.6 0 0 1-.6.6H17m0-20v2m0 18v-2m0 2H7m0 0v-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MpegFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M7.5 15v-3m0 0V9h3v3zm-6 3V9L3 12l1.5-3v6m12-6h-3v6h3m6-6h-3v6h3v-2.4m-9-.6h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiBubble(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 22a5.5 5.5 0 1 0-4.764-2.75l-.461 2.475l2.475-.46A5.474 5.474 0 0 0 7.5 22"/><path d="M15.282 17.898A7.946 7.946 0 0 0 18 16.93l3.6.67l-.67-3.6A8 8 0 1 0 6.083 8.849"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiMacOsWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 19v-8a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m10 12.01l.01-.011m1.99.011l.01-.011m1.99.011l.01-.011"/><path d="M6.5 16H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v3"/><path stroke-linecap="round" stroke-linejoin="round" d="m5 7.01l.01-.011M7 7.01l.01-.011M9 7.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 19v-8a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2Z"/><path d="M6.5 16H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v3"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 12h1M5 7h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplePages(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 18h7m-7-4h1m-1-4h3M7 2h9.5L21 6.5V19"/><path d="M3 20.5v-14A1.5 1.5 0 0 1 4.5 5h9.752a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 18 8.75V20.5a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 3 20.5"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplePagesEmpty(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 2h9.5L21 6.5V19"/><path d="M3 20.5v-14A1.5 1.5 0 0 1 4.5 5h9.752a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 18 8.75V20.5a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 3 20.5"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplePagesMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 2h9.5L21 6.5V19"/><path d="M11 22h5.5a1.5 1.5 0 0 0 1.5-1.5V8.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 14.25 5H4.5A1.5 1.5 0 0 0 3 6.5V13m-1.008 6h6"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplePagesPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.992 19h3m3 0h-3m0 0v-3m0 3v3M7 2h9.5L21 6.5V19"/><path d="M11 22h5.5a1.5 1.5 0 0 0 1.5-1.5V8.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 14.25 5H4.5A1.5 1.5 0 0 0 3 6.5V13"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplePagesXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.87 21.121L4.993 19m2.121-2.121L4.993 19m0 0L2.87 16.879M4.992 19l2.121 2.121M7 2h9.5L21 6.5V19"/><path d="M11 22h5.5a1.5 1.5 0 0 0 1.5-1.5V8.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 14.25 5H4.5A1.5 1.5 0 0 0 3 6.5V13"/><path d="M14 5v3.4a.6.6 0 0 0 .6.6H18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicDoubleNote(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 14V3L9 5v11"/><path d="M17 19h1a2 2 0 0 0 2-2v-3h-3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2M6 21h1a2 2 0 0 0 2-2v-3H6a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicDoubleNotePlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 6.5h3m3 0h-3m0 0v-3m0 3v3M6 16V5l8-1m1 10v-4m-3 9h1a2 2 0 0 0 2-2v-3h-3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2m-9 2h1a2 2 0 0 0 2-2v-3H3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNote(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16v3a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2zm0 0V8m0 0V4l5-1v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNotePlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 10h3m3 0h-3m0 0V7m0 3v3M7 16v3a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2zm0 0V8m0 0V4l5-1v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNotePlusSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 10h3m3 0h-3m0 0V7m0 3v3"/><path fill="currentColor" d="M5 21H4a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2h3V4l5-1v4L7 8v11a2 2 0 0 1-2 2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16v3a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2zm0 0V8m0 0V4l5-1v4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNoteSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path fill="currentColor" d="M10 21H9a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2h3V4l5-1v4l-5 1v11a2 2 0 0 1-2 2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16v3a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2zm0 0V8m0 0V4l5-1v4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nsquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M10 16V8l4 8V8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavArrowDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 9l6 6l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavArrowLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 6l-6 6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavArrowRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 6l6 6l-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavArrowUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 15l6-6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Navigator(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path d="M17.873 15.475c.46.87-.437 1.831-1.336 1.432l-4.538-2.017l-4.537 2.017c-.9.4-1.797-.562-1.337-1.432l4.959-9.365a1.036 1.036 0 0 1 1.831 0z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavigatorAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path d="M13.93 17.869c-.322.93-1.637.929-1.958-.001l-1.62-4.694l-4.57-1.943c-.905-.385-.814-1.698.136-1.954L16.15 6.516a1.036 1.036 0 0 1 1.249 1.34z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Neighbourhood(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11 21H4a2 2 0 0 1-2-2v-4.54a2 2 0 0 1 .963-1.71l3.5-2.122a2 2 0 0 1 2.074 0l3.5 2.121A2 2 0 0 1 13 14.46V19a2 2 0 0 1-2 2M6.5 10V6.46a2 2 0 0 1 .963-1.71l3.5-2.122a2 2 0 0 1 2.074 0l3.5 2.121a2 2 0 0 1 .963 1.71V10M16 21h4a2 2 0 0 0 2-2v-4.54a2 2 0 0 0-.963-1.71l-3.506-2.125a2 2 0 0 0-2.065-.005l-.633.38"/><path d="M9 21v-3.4a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 0-.6.6V21m12 0v-3.4a.6.6 0 0 0-.6-.6H16"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Network(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" x="3" y="2" rx=".6"/><rect width="7" height="5" x="8.5" y="17" rx=".6"/><rect width="7" height="5" x="14" y="2" rx=".6"/><path d="M6.5 7v3.5a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V7M12 12.5V17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" x="2" y="21" rx=".6" transform="rotate(-90 2 21)"/><rect width="7" height="5" x="17" y="15.5" rx=".6" transform="rotate(-90 17 15.5)"/><rect width="7" height="5" x="2" y="10" rx=".6" transform="rotate(-90 2 10)"/><path d="M7 17.5h3.5a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2H7m5.5 5.5H17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkLeftSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" x="2" y="21" fill="currentColor" rx=".6" transform="rotate(-90 2 21)"/><rect width="7" height="5" x="17" y="15.5" fill="currentColor" rx=".6" transform="rotate(-90 17 15.5)"/><rect width="7" height="5" x="2" y="10" fill="currentColor" rx=".6" transform="rotate(-90 2 10)"/><path d="M7 17.5h3.5a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2H7m5.5 5.5H17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkReverse(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" rx=".6" transform="matrix(1 0 0 -1 3 22)"/><rect width="7" height="5" rx=".6" transform="matrix(1 0 0 -1 8.5 7)"/><rect width="7" height="5" rx=".6" transform="matrix(1 0 0 -1 14 22)"/><path d="M6.5 17v-3.5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2V17M12 11.5V7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkReverseSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" fill="currentColor" rx=".6" transform="matrix(1 0 0 -1 3 22)"/><rect width="7" height="5" fill="currentColor" rx=".6" transform="matrix(1 0 0 -1 8.5 7)"/><rect width="7" height="5" fill="currentColor" rx=".6" transform="matrix(1 0 0 -1 14 22)"/><path d="M6.5 17v-3.5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2V17M12 11.5V7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" rx=".6" transform="matrix(0 -1 -1 0 22 21)"/><rect width="7" height="5" rx=".6" transform="matrix(0 -1 -1 0 7 15.5)"/><rect width="7" height="5" rx=".6" transform="matrix(0 -1 -1 0 22 10)"/><path d="M17 17.5h-3.5a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2H17M11.5 12H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkRightSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" fill="currentColor" rx=".6" transform="matrix(0 -1 -1 0 22 21)"/><rect width="7" height="5" fill="currentColor" rx=".6" transform="matrix(0 -1 -1 0 7 15.5)"/><rect width="7" height="5" fill="currentColor" rx=".6" transform="matrix(0 -1 -1 0 22 10)"/><path d="M17 17.5h-3.5a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2H17M11.5 12H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" x="3" y="2" fill="currentColor" rx=".6"/><rect width="7" height="5" x="8.5" y="17" fill="currentColor" rx=".6"/><rect width="7" height="5" x="14" y="2" fill="currentColor" rx=".6"/><path d="M6.5 7v3.5a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V7M12 12.5V17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewTab(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 19V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M9 14h3m3 0h-3m0 0v-3m0 3v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NintendoSwitch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 17V7a4 4 0 0 1 4-4h3.9a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H6a4 4 0 0 1-4-4Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m11 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M22 17V7a4 4 0 0 0-4-4h-3.9a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6H18a4 4 0 0 0 4-4Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSmokingCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M15 12v3m0-6c0-1-.714-2-2.143-2v0A2.857 2.857 0 0 1 10 4.143V3m8 6V4m0 8v3"/><path d="M15 15H6.6a.6.6 0 0 1-.6-.6v-1.8a.6.6 0 0 1 .6-.6H12"/><path stroke-linecap="round" stroke-linejoin="round" d="m5 5l14 14m-7 3c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NonBinary(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9a6 6 0 1 1 0 12a6 6 0 0 1 0-12m0 0V3M9 4l6 3m0-3L9 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notes(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 14h8m-8-4h2m-2 8h4M10 3H6a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-3.5M10 3V1m0 2v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Npm(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M1 8h22v7H11v2H7.5v-2H1zm6.5 0v7m6-7v7"/><path d="M18 11v4M5 11v4m6-4v1m9.5-1v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NpmSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M8 16h8V8H8z"/><path d="M13 11v5"/><path stroke-linejoin="round" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberEightSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 16c-1.38 0-2.5-.5-2.5-2s1.12-2 2.5-2s2.5.5 2.5 2s-1.12 2-2.5 2m0-8c-1.38 0-2.5.5-2.5 2s1.12 2 2.5 2s2.5-.5 2.5-2s-1.12-2-2.5-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberEightSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zm11.092 9.455c.225.155.408.41.408.945c0 .535-.183.79-.409.945c-.272.189-.72.305-1.341.305s-1.069-.116-1.341-.305c-.226-.155-.409-.41-.409-.945c0-.535.183-.79.409-.945c.272-.189.72-.305 1.341-.305s1.069.116 1.342.305M9.576 12c-.534.47-.826 1.15-.826 2c0 .965.376 1.71 1.056 2.18c.632.436 1.435.57 2.194.57c.76 0 1.562-.134 2.194-.57c.68-.47 1.056-1.215 1.056-2.18c0-.85-.292-1.53-.826-2c.534-.47.826-1.15.826-2c0-.965-.376-1.71-1.056-2.18c-.632-.436-1.435-.57-2.194-.57c-.76 0-1.562.134-2.194.57c-.68.47-1.056 1.215-1.056 2.18c0 .85.292 1.53.826 2M12 11.25c.621 0 1.069-.116 1.342-.305c.225-.155.408-.41.408-.945c0-.535-.183-.79-.409-.945c-.272-.189-.72-.305-1.341-.305s-1.069.116-1.341.305c-.226.155-.409.41-.409.945c0 .535.183.79.409.945c.272.189.72.305 1.341.305" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberFiveSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.5 14c0 1.105 1.12 2 2.5 2s2.5-1 2.5-2.5s-.62-2.5-2-2.5h-3l.5-3h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberFiveSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zM13.75 13.5c0 1.062-.76 1.75-1.75 1.75c-1.131 0-1.75-.707-1.75-1.25a.75.75 0 0 0-1.5 0c0 1.666 1.62 2.75 3.25 2.75c1.771 0 3.25-1.312 3.25-3.25c0-.836-.17-1.645-.632-2.262c-.488-.652-1.225-.988-2.118-.988h-2.115l.25-1.5H14a.75.75 0 0 0 0-1.5h-4a.75.75 0 0 0-.74.627l-.5 3a.75.75 0 0 0 .74.873h3c.488 0 .75.164.917.387c.193.258.333.7.333 1.363" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberFourSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M13.5 16V8L9 13.5h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberFourSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zM14.25 8a.75.75 0 0 0-1.33-.475l-4.5 5.5A.75.75 0 0 0 9 14.25h3.75V16a.75.75 0 0 0 1.5 0v-1.75H15a.75.75 0 0 0 0-1.5h-.75zm-1.5 2.101v2.649h-2.167z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberNineSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 8c1.38 0 2.5.5 2.5 2.5c0 1.5-1.12 2-2.5 2s-2.5-.5-2.5-2c0-2 1.12-2.5 2.5-2.5"/><path stroke-linecap="round" stroke-linejoin="round" d="M14.5 11c0 3 0 5-4.5 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberNineSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zm8.409 7.845c-.226-.155-.409-.41-.409-.945c0-.861.236-1.237.46-1.423c.254-.21.666-.327 1.29-.327s1.036.118 1.29.327c.224.186.46.562.46 1.423c0 .535-.183.79-.409.945c-.272.189-.72.305-1.341.305s-1.069-.116-1.341-.305m-.853 1.235c.632.436 1.435.57 2.194.57c.567 0 1.159-.075 1.685-.294a.1.1 0 0 0 0 .005c-.064.554-.185.957-.378 1.258c-.348.54-1.126 1.031-3.307 1.031a.75.75 0 0 0 0 1.5c2.319 0 3.79-.51 4.568-1.72c.37-.573.53-1.232.607-1.897c.075-.653.075-1.376.075-2.105V11a.75.75 0 0 0-.017-.158a3.45 3.45 0 0 0 .017-.342c0-1.139-.323-2.013-1.004-2.577c-.651-.54-1.49-.673-2.246-.673c-.757 0-1.595.132-2.246.673c-.68.564-1.004 1.438-1.004 2.577c0 .965.376 1.71 1.056 2.18" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberOneSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 16V8l-3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberOneSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zM13.75 8a.75.75 0 0 0-1.28-.53l-3 3a.75.75 0 1 0 1.06 1.06l1.72-1.72V16a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberSevenSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10.5 16c0-4 4-8 4-8h-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberSevenSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm5.9 5a.75.75 0 0 0 0 1.5h3.35c-.3.378-.637.828-.974 1.334C10.851 11.621 9.75 13.774 9.75 16a.75.75 0 0 0 1.5 0c0-1.775.899-3.621 1.874-5.084a18.4 18.4 0 0 1 1.752-2.225l.118-.124l.03-.03l.006-.006a.75.75 0 0 0-.53-1.281z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberSixSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 16c-1.38 0-2.5-.5-2.5-2.5c0-1.5 1.12-2 2.5-2s2.5.5 2.5 2c0 2-1.12 2.5-2.5 2.5"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.5 13c0-3 0-5 4.5-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberSixSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zm11.092 8.955c.225.155.408.41.408.945c0 .861-.236 1.236-.46 1.423c-.254.21-.666.327-1.29.327s-1.036-.118-1.29-.327c-.224-.187-.46-.562-.46-1.423c0-.535.183-.79.409-.945c.272-.189.72-.305 1.341-.305s1.069.116 1.342.305m.852-1.235c-.632-.436-1.435-.57-2.194-.57c-.567 0-1.159.075-1.685.294a.1.1 0 0 0 0-.005c.064-.554.185-.957.378-1.258C11.041 9.24 11.82 8.75 14 8.75a.75.75 0 0 0 0-1.5c-2.319 0-3.79.51-4.568 1.72c-.37.573-.53 1.232-.607 1.897c-.075.653-.075 1.376-.075 2.105V13a.75.75 0 0 0 .017.158a3.45 3.45 0 0 0-.017.342c0 1.139.323 2.014 1.004 2.577c.651.54 1.49.673 2.246.673c.757 0 1.595-.132 2.246-.673c.68-.564 1.004-1.439 1.004-2.577c0-.965-.376-1.71-1.056-2.18" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberThreeSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.5 10c0-1.105 1.12-2 2.5-2s2.5.895 2.5 2s-.62 2-2 2m-3 2c0 1.105 1.12 2 2.5 2s2.5-.895 2.5-2s-.62-2-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberThreeSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zM12 8.75c1.131 0 1.75.707 1.75 1.25c0 .42-.117.722-.288.911c-.16.177-.442.339-.962.339a.75.75 0 0 0 0 1.5c.52 0 .802.162.962.339c.17.19.288.492.288.911c0 .543-.619 1.25-1.75 1.25s-1.75-.707-1.75-1.25a.75.75 0 0 0-1.5 0c0 1.666 1.62 2.75 3.25 2.75s3.25-1.084 3.25-2.75c0-.685-.193-1.383-.676-1.917a2.452 2.452 0 0 0-.079-.083a2.36 2.36 0 0 0 .079-.083c.483-.534.676-1.232.676-1.917c0-1.666-1.62-2.75-3.25-2.75S8.75 8.334 8.75 10a.75.75 0 0 0 1.5 0c0-.543.619-1.25 1.75-1.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberTwoSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.5 10.8v-.4c0-1.325 1.033-2.4 2.308-2.4c1.274 0 2.307 1.075 2.307 2.4c0 .457-.122.884-.336 1.248C12.73 13.44 9.5 16 9.5 16h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberTwoSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm6.65 8.15c0-.94.725-1.65 1.558-1.65c.832 0 1.557.71 1.557 1.65c0 .321-.086.618-.233.868c-.462.789-1.454 1.81-2.394 2.676a35.726 35.726 0 0 1-1.671 1.442l-.025.02l-.006.005l-.002.001A.75.75 0 0 0 9.5 16.75h5a.75.75 0 0 0 0-1.5h-2.968l.222-.202c.936-.863 2.084-2.018 2.673-3.021a3.21 3.21 0 0 0 .438-1.627c0-1.712-1.341-3.15-3.057-3.15S8.75 8.688 8.75 10.4v.4a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberZeroSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path d="M9.5 14v-4a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberZeroSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zM10.25 10c0-.69.56-1.25 1.25-1.25h1c.69 0 1.25.56 1.25 1.25v4c0 .69-.56 1.25-1.25 1.25h-1c-.69 0-1.25-.56-1.25-1.25zm1.25-2.75A2.75 2.75 0 0 0 8.75 10v4a2.75 2.75 0 0 0 2.75 2.75h1A2.75 2.75 0 0 0 15.25 14v-4a2.75 2.75 0 0 0-2.75-2.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberedListLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5h12M5 7V3L3.5 4.5m2 9.5h-2l1.905-2.963a.428.428 0 0 0 .072-.323C5.42 10.456 5.216 10 4.5 10c-1 0-1 .889-1 .889s0 0 0 0v.222M4 19h.5a1 1 0 0 1 1 1v0a1 1 0 0 1-1 1h-1m0-4h2L4 19m5-7h12M9 19h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumberedListRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 5h12m5.5 2V3L19 4.5m2 9.5h-2l1.905-2.963a.428.428 0 0 0 .072-.323C20.92 10.456 20.716 10 20 10c-1 0-1 .889-1 .889s0 0 0 0v.222M19.5 19h.5a1 1 0 0 1 1 1v0a1 1 0 0 1-1 1h-1m0-4h2l-1.5 2M3 12h12M3 19h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Osquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M12.2 8h-.4A1.8 1.8 0 0 0 10 9.8v4.4a1.8 1.8 0 0 0 1.8 1.8h.4a1.8 1.8 0 0 0 1.8-1.8V9.8A1.8 1.8 0 0 0 12.2 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Octagon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.77 1.095a.6.6 0 0 1 .46 0l7.319 3.032a.6.6 0 0 1 .324.324l3.032 7.32a.6.6 0 0 1 0 .459l-3.032 7.319a.6.6 0 0 1-.324.324l-7.32 3.032a.6.6 0 0 1-.459 0l-7.319-3.032a.6.6 0 0 1-.324-.324l-3.032-7.32a.6.6 0 0 1 0-.459l3.032-7.319a.6.6 0 0 1 .324-.324z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OffTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 15V9a6 6 0 0 1 6-6h10a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H7a6 6 0 0 1-6-6Z"/><path d="M7 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15V9h3m2 6V9h3m-8 3h2.572M17 12h2.572"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OilIndustry(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 10c0-1-1-2-3-2h-1a3 3 0 0 1-3-3V2"/><path d="M9 10.8C9 9.033 6 6 6 6s-3 3.033-3 4.8S4.343 14 6 14s3-1.433 3-3.2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 21h3v-9h-3v4.5m0 4.5v-4.5m0 4.5h-7.5v-4.5l4-2.5l3.5 2.5m3-6.5c0-6-4-6-4-6s4 .5 4-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Okrs(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0-7a2 2 0 1 0 0-4a2 2 0 0 0 0 4M3 5h10M3 12h10M3 19h10m3 2.243l2.121-2.122m0 0L20.243 17m-2.122 2.121L16 17m2.121 2.121l2.122 2.122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OnTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 15V9a6 6 0 0 1 6-6h10a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H7a6 6 0 0 1-6-6Z"/><path d="M9 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 15V9l4 6V9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OneFingerSelectHandGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7.5 12l-2.004 2.672a2 2 0 0 0 .126 2.552l3.784 4.128c.378.413.912.648 1.473.648H15.5c2.4 0 4-2 4-4c0 0 0 0 0 0V9.429m-3 .571v-.571c0-2.286 3-2.286 3 0"/><path d="M13.5 10V8.286c0-2.286 3-2.286 3 0V10m-6 0V7.5c0-2.286 3-2.286 3 0c0 0 0 0 0 0V10m-3 0V3.499A1.5 1.5 0 0 0 9 2v0a1.5 1.5 0 0 0-1.5 1.5V15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OnePointCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="m19 19l-1.5-1.5m-2-2l-1-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenBook(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 21V7a2 2 0 0 1 2-2h7.4a.6.6 0 0 1 .6.6v13.114M12 21V7a2 2 0 0 0-2-2H2.6a.6.6 0 0 0-.6.6v13.114M14 19h8m-12 0H2"/><path stroke-linejoin="round" d="M12 21a2 2 0 0 1 2-2m-2 2a2 2 0 0 0-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenInBrowser(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 21h12.4a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6V16m7-10h8M6 6h1M3.5 20.5L12 12m0 0v4m0-4H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenInWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 21h12.4a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6V16m.5 4.5L12 12m0 0v4m0-4H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenNewWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21 3h-6m6 0l-9 9m9-9v6"/><path d="M21 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenSelectHandGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 14.571l-1.823-1.736a1.558 1.558 0 0 0-2.247.103v0a1.558 1.558 0 0 0 .035 2.092l5.942 6.338c.379.403.906.632 1.459.632H16c2.4 0 4-2 4-4c0 0 0 0 0 0V9.429"/><path d="M17 10v-.571c0-2.286 3-2.286 3 0M14 10V8.286C14 6 17 6 17 8.286V10m-6 0V7.5c0-2.286 3-2.286 3 0c0 0 0 0 0 0V10m-6 4.571V3.5A1.5 1.5 0 0 1 9.5 2v0c.828 0 1.5.67 1.5 1.499V10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenVpn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.835 15.29l.738.136zm-.358-.708l.381-.646zm-.275 7.247l.138-.738zm-.452-.678l.738.136zm7.099-1.337l.737-.139zm.872.378l-.43-.615zm-9.85-4.208l-.736-.139zm-.139-.52l-.581.474zm5.791-.882l.382.646zm-.358.707l.738-.136zm3.103.175l-.581-.473zm-.14.52l.737-.139zm-1.878 5.167l-.737.137zm-.453.679l.138.737zM6.28 20.192l-.43.614zM2.75 12A9.25 9.25 0 0 1 12 2.75v-1.5C6.062 1.25 1.25 6.062 1.25 12zm3.959 7.577C4.315 17.902 2.75 15.137 2.75 12h-1.5c0 3.65 1.824 6.865 4.599 8.806zm.426-3.732l-.721 3.83l1.474.278l.72-3.83zM5.75 12c0 1.494.526 2.865 1.4 3.938l1.163-.947A4.713 4.713 0 0 1 7.25 12zM12 5.75A6.25 6.25 0 0 0 5.75 12h1.5A4.75 4.75 0 0 1 12 7.25zM18.25 12A6.25 6.25 0 0 0 12 5.75v1.5A4.75 4.75 0 0 1 16.75 12zm-1.4 3.938A6.213 6.213 0 0 0 18.25 12h-1.5a4.713 4.713 0 0 1-1.063 2.99zm.736 3.737l-.72-3.83l-1.475.278l.72 3.83zM21.25 12a9.225 9.225 0 0 1-3.959 7.577l.86 1.23C20.926 18.864 22.75 15.65 22.75 12zM12 2.75A9.25 9.25 0 0 1 21.25 12h1.5c0-5.938-4.812-10.75-10.75-10.75zM15.75 12A3.75 3.75 0 0 0 12 8.25v1.5A2.25 2.25 0 0 1 14.25 12zm-1.845 3.228A3.745 3.745 0 0 0 15.75 12h-1.5c0 .823-.443 1.544-1.108 1.936zm1.083 5.787l-1.085-5.862l-1.475.273l1.085 5.862zM12 22.75c.665 0 1.31-.067 1.935-.183l-.275-1.474a9.036 9.036 0 0 1-1.66.157zm-1.937-.184c.625.117 1.271.184 1.937.184v-1.5a8.958 8.958 0 0 1-1.66-.159zm.035-7.413l-1.085 5.861l1.475.273l1.085-5.861zM8.25 12c0 1.377.744 2.578 1.846 3.228l.762-1.292A2.245 2.245 0 0 1 9.75 12zM12 8.25A3.75 3.75 0 0 0 8.25 12h1.5A2.25 2.25 0 0 1 12 9.75zm-.427 7.176c.122-.662-.259-1.22-.715-1.49l-.762 1.292a.053.053 0 0 1 .01.008c.003.003.002.003 0-.001a.1.1 0 0 1-.009-.03a.153.153 0 0 1 0-.052zm-1.233 5.666c.119.022.16.128.148.195l-1.475-.273c-.129.694.305 1.412 1.05 1.552zm5.772-1.14c.168.892 1.212 1.433 2.04.854l-.86-1.229a.21.21 0 0 1 .197-.019c.052.023.088.07.097.117zm-7.503-3.83a1.383 1.383 0 0 0-.296-1.131l-1.162.947c-.007-.008-.026-.04-.016-.093zm4.533-2.186c-.456.27-.837.828-.714 1.49l1.475-.273a.153.153 0 0 1 0 .053a.1.1 0 0 1-.009.029c-.002.004-.003.004 0 .001a.05.05 0 0 1 .01-.008zm2.545 1.055c-.245.3-.375.709-.296 1.132l1.474-.278c.01.054-.009.085-.016.093zm-2.174 6.297a.174.174 0 0 1 .147-.195l.275 1.474c.745-.139 1.181-.857 1.053-1.552zm-7.664-.482c.827.579 1.871.038 2.04-.853l-1.475-.277a.166.166 0 0 1 .097-.118a.21.21 0 0 1 .198.02z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrangeHalf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.5 0 10-4.5 10-10S17.5 2 12 2m0 20C6.5 22 2 17.5 2 12S6.5 2 12 2m0 20V12m0-10v10m0 0l5 5.5M12 12l5-5m-5 5h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrangeSlice(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.613 10.11l7.778-7.777c4.295 4.296 4.295 11.26 0 15.556c-4.296 4.296-11.261 4.296-15.557 0zm0 0l-.354 8.133m.354-8.132h7.778m-7.778 0l5.303 5.303"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrangeSliceAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.39 10.11L5.61 2.334c-4.295 4.296-4.295 11.26 0 15.556c4.296 4.296 11.26 4.296 15.557 0zm0 0l.353 8.133m-.354-8.132H5.612m7.779 0l-5.304 5.303"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrganicFood(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 18s.9-3.741 3-6"/><path d="m16.186 7.241l.374 3.89c.243 2.523-1.649 4.77-4.172 5.012c-2.475.238-4.718-1.571-4.956-4.047a4.503 4.503 0 0 1 4.05-4.914l4.147-.4a.51.51 0 0 1 .557.46"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrganicFoodSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 18s.9-3.741 3-6"/><path d="m16.186 7.241l.374 3.89c.243 2.523-1.649 4.77-4.172 5.012c-2.475.238-4.718-1.571-4.956-4.047a4.503 4.503 0 0 1 4.05-4.914l4.147-.4a.51.51 0 0 1 .557.46"/><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrthogonalView(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21V3h18v18zm0-4.5h18M3 12h18M3 7.5h18M16.5 3v18M12 3v18M7.5 3v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 6v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2m-8 3V4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PackageLock(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 20H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v6m-8-3V4m9.167 14.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Packages(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 15v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2m6-10v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2m6 10v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2M6 16v-3m6-7V3m6 13v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pacman(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16 12l.011.01M19 12l.011.01M22 12l.011.01M2 12c0 5.523 4.477 10 10 10a9.985 9.985 0 0 0 8-3.999L12 12l8-6.001A9.985 9.985 0 0 0 12 2C6.477 2 2 6.477 2 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Page(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6M8 10h8m-8 8h8m-8-4h4"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.5 11l3.5 3.5l3.5-3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageEdit(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11M8 10h8M8 6h4m-4 8h3m6.954 2.94l1-1a1.121 1.121 0 0 1 1.586 0v0a1.121 1.121 0 0 1 0 1.585l-1 1m-1.586-1.586l-2.991 2.991a1 1 0 0 0-.28.553l-.244 1.557l1.557-.243a1 1 0 0 0 .553-.28l2.99-2.992m-1.585-1.586l1.586 1.586"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageFlip(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 11h5m-5-4h5m-9 8V3.6a.6.6 0 0 1 .6-.6h11.8a.6.6 0 0 1 .6.6V17a4 4 0 0 1-4 4v0"/><path d="M5 15h7.4c.331 0 .603.267.63.597C13.153 17.115 13.78 21 17 21H6a3 3 0 0 1-3-3v-1a2 2 0 0 1 2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 8.5L9.5 12l3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 12V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H11"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20M1.992 19h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageMinusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 12h6M4 21.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PagePlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 12V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H11"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20M1.992 19h3m3 0h-3m0 0v-3m0 3v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PagePlusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 12h3m3 0h-3m0 0V9m0 3v3m-8 6.4V2.6a.6.6 0 0 1 .6-.6h11.652a.6.6 0 0 1 .424.176l3.148 3.148A.6.6 0 0 1 20 5.75V21.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11 8.5l3.5 3.5l-3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageSearch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11M8 10h8M8 6h4m-4 8h3m9.5 6.5L22 22"/><path d="M15 18a3 3 0 1 0 6 0a3 3 0 0 0-6 0m1-16v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageStar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H11M8 10h8M8 6h4m-4 8h3"/><path d="m16.306 17.113l.909-1.927a.312.312 0 0 1 .57 0l.91 1.927l2.032.311c.261.04.365.376.177.568l-1.471 1.5l.347 2.118c.044.272-.229.48-.462.351l-1.818-1l-1.818 1c-.234.129-.506-.079-.462-.351l.347-2.118l-1.47-1.5c-.19-.192-.085-.528.176-.568zM16 2v3.4a.6.6 0 0 0 .6.6H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.5 13L12 9.5l3.5 3.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palette(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20.51 9.54a1.899 1.899 0 0 1-1 1.09A7 7 0 0 0 15.37 17c.001.47.048.939.14 1.4a2.16 2.16 0 0 1-.31 1.65a1.79 1.79 0 0 1-1.21.8a9 9 0 0 1-10.62-9.13A9.05 9.05 0 0 1 11.85 3h.51a9 9 0 0 1 8.06 5a2 2 0 0 1 .09 1.52z"/><path stroke-linecap="round" stroke-linejoin="round" d="m8 16.01l.01-.011M6 12.01l.01-.011M8 8.01l.01-.011M12 6.01l.01-.011M16 8.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanoramaEnlarge(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 5c2.995 0 7.235.692 8.576.925a.581.581 0 0 1 .48.503c.13 1.028.444 3.691.444 5.572c0 1.88-.313 4.544-.444 5.572a.581.581 0 0 1-.48.503c-1.34.233-5.58.925-8.576.925c-2.995 0-7.235-.692-8.576-.925a.582.582 0 0 1-.48-.503C2.814 16.544 2.5 13.881 2.5 12c0-1.88.313-4.544.444-5.572a.582.582 0 0 1 .48-.503C4.764 5.692 9.004 5 12 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanoramaReduce(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 6.862v10.276a.615.615 0 0 1-.811.58C18.546 17.165 14.749 16 12 16c-2.749 0-6.546 1.166-8.189 1.717a.615.615 0 0 1-.811-.58V6.863c0-.418.415-.712.811-.58C5.454 6.835 9.251 8 12 8c2.749 0 6.546-1.166 8.189-1.717a.615.615 0 0 1 .811.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pants(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19h4.436a.6.6 0 0 0 .6-.563l.924-14.8A.6.6 0 0 0 17.361 3H6.634a.6.6 0 0 0-.599.633l.934 16.8a.6.6 0 0 0 .599.567H11.4a.6.6 0 0 0 .6-.6V8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PantsPockets(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5.035 3.633a.6.6 0 0 1 .6-.633h12.73a.6.6 0 0 1 .6.633l-.933 16.8a.6.6 0 0 1-.6.567h-2.898a.6.6 0 0 1-.596-.53L12.596 9.065c-.083-.706-1.109-.706-1.192 0L10.062 20.47a.6.6 0 0 1-.596.53H6.568a.6.6 0 0 1-.6-.567z"/><path d="M5 7.5h1.5a2 2 0 0 0 2-2V3m10 4.5h-1a2 2 0 0 1-2-2V3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parking(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10 15.5v-2.8m0 0h2.857c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H10z"/><circle cx="12" cy="12" r="10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PasswordCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 13V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h7m2.5 2.5l2 2l4-4M12 11.01l.01-.011m3.99.011l.01-.011M8 11.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PasswordCursor(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 13V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h7"/><path d="M20.879 16.917c.494.304.463 1.043-.045 1.101l-2.567.291l-1.151 2.312c-.228.459-.933.234-1.05-.334l-1.255-6.116c-.099-.48.333-.782.75-.525z" clip-rule="evenodd"/><path stroke-linecap="round" stroke-linejoin="round" d="m12 11.01l.01-.011m3.99.011l.01-.011M8 11.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PasswordXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.121 20.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L15.12 16.12m2.122 2.122l2.121 2.121M21 13V8a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6m1-4.99l.01-.011m3.99.011l.01-.011M8 11.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PasteClipboard(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M8.5 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-2.5"/><path d="M8 6.4V4.5a.5.5 0 0 1 .5-.5c.276 0 .504-.224.552-.496C9.2 2.652 9.774 1 12 1s2.8 1.652 2.948 2.504c.048.272.276.496.552.496a.5.5 0 0 1 .5.5v1.9a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathArrow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 16.5V3m0 0l3.5 3.5M18 3l-3.5 3.5m3.5 10a3.5 3.5 0 1 1-7 0v-9m0 0a3.5 3.5 0 1 0-7 0v12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M6 18.4V5.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6Zm8 0V5.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-width="1.5" d="M6 18.4V5.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6Zm8 0V5.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v12.8a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M16 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v9"/><path stroke-linejoin="round" d="M19 17v5m3-5v5M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 17.5L6 3h7c6 0 6 9 0 9H8.7l-1.2 5.5z"/><path d="m6.8 21l3-14.5h7c6 0 6 9 0 9h-4.3L11.3 21z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PcCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 22h10"/><path d="M2 17V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 10.5l2 2l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PcFirewall(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 22h10"/><path d="M2 17V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m12.485 6.121l3.06.765a.59.59 0 0 1 .449.586C15.818 14 12 15 12 15s-3.818-1-3.994-7.528a.59.59 0 0 1 .448-.586l3.06-.765a2 2 0 0 1 .971 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PcMouse(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 2v0a8 8 0 0 1 8 8v4a8 8 0 0 1-8 8v0a8 8 0 0 1-8-8v-4a8 8 0 0 1 8-8zm0 0v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PcNoEntry(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 22h10"/><path d="M2 17V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14.857 7.7a4 4 0 1 0-5.713 5.6m5.713-5.6a4 4 0 0 1-5.713 5.6m5.713-5.6l-5.714 5.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PcWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 22h10"/><path d="M2 17V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 7v3m0 4.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PeaceHand(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14.149 9.472v-5.86c0-.89-.722-1.612-1.612-1.612v0c-.89 0-1.611.722-1.611 1.612v4.834"/><path d="m16.346 12.841l2.176-7.252a1.584 1.584 0 0 0-1.083-1.98v0a1.585 1.585 0 0 0-1.961 1.098l-1.33 4.764M7.62 9.25l1.055 2.341a1.612 1.612 0 0 1-2.938 1.325L4.68 10.575A1.612 1.612 0 0 1 7.62 9.25Z"/><path d="M11.72 12.262v0a2.322 2.322 0 0 0-.068-1.743l-1.073-2.38a1.584 1.584 0 0 0-2.101-.79v0a1.584 1.584 0 0 0-.764 2.14l.135.276"/><path d="m13.857 17.677l.492-.984a.176.176 0 0 0-.108-.248l-3.55-1.044a1.537 1.537 0 0 1-1.095-1.635v0a1.537 1.537 0 0 1 1.67-1.37l4.788.446s3.81.586 2.49 4.395c-1.318 3.81-1.757 5.128-4.687 5.128H8.876a4.249 4.249 0 0 1-4.249-4.249v0L4.48 9.912"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Peerlist(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.87 3h6.26a6 6 0 0 1 5.963 5.337l.21 1.896c.131 1.174.131 2.36 0 3.534l-.21 1.896A6 6 0 0 1 15.13 21H8.87a6 6 0 0 1-5.963-5.337l-.21-1.896a16 16 0 0 1 0-3.534l.21-1.896A6 6 0 0 1 8.87 3"/><path d="M9 17v-4m0 0V7h4a3 3 0 0 1 3 3v0a3 3 0 0 1-3 3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenConnectBluetooth(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.5 17.5l-1 4l3.731-.933a1 1 0 0 0 .465-.263L21.5 8.5a2.121 2.121 0 0 0 0-3v0a2.121 2.121 0 0 0-3 0l-4 4m3-3l3 3M5 6.6l7 5.1L8.333 15V3L12 6.3l-7 5.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenConnectWifi(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 9.76l.01-.011M3 6.25c2.5-3 7.5-3 10 0m-8 2c1.5-2 4.5-2 6 0m6.5-1.75l1-1a2.121 2.121 0 0 1 3 0v0a2.121 2.121 0 0 1 0 3l-1 1m-3-3L6.696 17.304a1 1 0 0 0-.263.465L5.5 21.5l3.731-.933a1 1 0 0 0 .465-.263L20.5 9.5m-3-3l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenTablet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M22 5v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2ZM2 12h4m0-9v18"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.5 11.5l-3.5 3m5-4.49l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenTabletConnectUsb(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M22 7V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2"/><path d="M2 12h4m0-9v18"/><path stroke-linecap="round" stroke-linejoin="round" d="M19.25 12H11m7.7 0l-.825 3h-1.65m1.375-3l-1.1-3h-1.925M22 12a1.37 1.37 0 0 0-1.375-1.364c-.76 0-1.375.61-1.375 1.364a1.37 1.37 0 0 0 1.375 1.364c.76 0 1.375-.61 1.375-1.364"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenTabletConnectWifi(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m17 15.51l.01-.011M12 12c2.5-3 7.5-3 10 0m-8 2c1.5-2 4.5-2 6 0"/><path stroke-linecap="round" stroke-linejoin="round" d="M22 7V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2"/><path d="M2 12h4m0-9v18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pentagon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.647 2.256a.6.6 0 0 1 .706 0l9.756 7.089a.6.6 0 0 1 .218.67L18.6 21.485a.6.6 0 0 1-.57.414H5.97a.6.6 0 0 1-.57-.414l-3.727-11.47a.6.6 0 0 1 .218-.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PeopleTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 16V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.5 14.5s-1.5 2-4.5 2s-4.5-2-4.5-2"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 10a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentRotateOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.168 8A10.002 10.002 0 0 0 12 2c-5.185 0-9.449 3.947-9.95 9"/><path d="M18 8h3.4a.6.6 0 0 0 .6-.6V4M2.881 16c1.544 3.532 5.068 6 9.168 6c5.186 0 9.45-3.947 9.951-9"/><path d="M6.05 16h-3.4a.6.6 0 0 0-.6.6V20"/><path fill="currentColor" d="M14.5 15a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-5-5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/><path d="m15 9l-6 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percentage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 19a2 2 0 1 1 0-4a2 2 0 0 1 0 4M7 9a2 2 0 1 1 0-4a2 2 0 0 1 0 4m12-4L5 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path fill="currentColor" d="M15.5 16a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-7-7a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/><path d="m16 8l-8 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25m3.5 13a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5M7.25 8.5a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0m8.22-1.03a.75.75 0 1 1 1.06 1.06l-8 8a.75.75 0 0 1-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 16a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-7-7a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/><path stroke-linecap="round" stroke-linejoin="round" d="m16 8l-8 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm11.9 12a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5M7.25 8.5a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0m8.22-1.03a.75.75 0 1 1 1.06 1.06l-8 8a.75.75 0 0 1-1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PerspectiveView(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1 21L4.143 3h15.714L23 21zm1-4.5h20M3 12h18M4 7.5h16M12 3v18M8 3.5l-1.5 17m9.5-17l1.5 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PharmacyCrossCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M13.9 18h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 1-.6-.6v-3.8a.6.6 0 0 1 .6-.6h2.3a.6.6 0 0 0 .6-.6V6.6a.6.6 0 0 1 .6-.6h3.8a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 0 .6.6h2.3a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6h-2.3a.6.6 0 0 0-.6.6v2.3a.6.6 0 0 1-.6.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PharmacyCrossTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5"/><path d="M13.9 18h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 0-.6-.6H6.6a.6.6 0 0 1-.6-.6v-3.8a.6.6 0 0 1 .6-.6h2.3a.6.6 0 0 0 .6-.6V6.6a.6.6 0 0 1 .6-.6h3.8a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 0 .6.6h2.3a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6h-2.3a.6.6 0 0 0-.6.6v2.3a.6.6 0 0 1-.6.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.118 14.702L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneDisabled(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.78 8.5l.49-2.63L7.815 2H4.064c-1.128 0-2.016.93-1.848 2.046c.288 1.902.957 4.861 2.51 7.7M10.94 13.5c.837.744 1.847 1.392 3.059 2l4.118-.798L22 16.182v3.584c0 1.192-1.032 2.1-2.197 1.847c-2.83-.616-7.83-2.107-11.58-5.432M21 3L3 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneIncome(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 5h-6m0 0l3-3m-3 3l3 3m-.882 6.702L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.242 5.243h6m-4.124 9.459L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutcome(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 5h6m0 0l-3-3m3 3l-3 3m-.882 6.702L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonePaused(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 2v5m4-5v5m-3.882 7.702L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonePlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.243 5.243h3m3 0h-3m0 0v-3m0 3v3m-1.125 6.459L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17.121 7.364l2.122-2.121m2.121-2.122l-2.121 2.122m0 0L17.12 3.12m2.122 2.122l2.121 2.121m-3.245 7.339L14 15.5c-2.782-1.396-4.5-3-5.5-5.5l.77-4.13L7.815 2H4.064c-1.128 0-2.016.932-1.847 2.047c.42 2.783 1.66 7.83 5.283 11.453c3.805 3.805 9.286 5.456 12.302 6.113c1.165.253 2.198-.655 2.198-1.848v-3.584z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiggyBank(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M14.5 8.5c-.78-.202-1.866-.5-2.735-.5C7.476 8 4 10.668 4 13.958c0 1.891 1.148 3.577 2.938 4.668l-.485 1.6a.6.6 0 0 0 .574.774h1.764a.6.6 0 0 0 .36-.12l1.395-1.047h2.437l1.395 1.047a.6.6 0 0 0 .36.12h1.764a.6.6 0 0 0 .574-.774l-.485-1.6c1.067-.65 1.905-1.511 2.409-2.501M14.5 8.5L19 7l-.084 3.628L21 11.5V15l-1.926 1"/><path fill="currentColor" stroke-linecap="round" d="M15.5 13a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/><path stroke-linecap="round" d="M2 10s0 2.4 2 3"/><path d="M12.8 7.753c.13-.372.2-.772.2-1.188C13 4.596 11.433 3 9.5 3S6 4.596 6 6.565c0 .941.358 1.798.944 2.435"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pillow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21.04 12.283l.599 4.19a2 2 0 0 1-2.179 2.273l-7.26-.726a2.005 2.005 0 0 0-.398 0l-7.261.726a2 2 0 0 1-2.179-2.273l.599-4.19a2 2 0 0 0 0-.566l-.599-4.19A2 2 0 0 1 4.54 5.254l7.261.726a2 2 0 0 0 .398 0l7.261-.726a2 2 0 0 1 2.179 2.273l-.599 4.19a2 2 0 0 0 0 .566M21 6l-4 3M7 15l-4 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 14.5L3 21M5 9.485l9.193 9.193l1.697-1.697l-.393-3.787l5.51-4.673l-5.85-5.85l-4.674 5.51l-3.786-.393z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinSlash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 14.5L3 21M7.676 7.89l-.979-.102L5 9.485l9.193 9.193l1.697-1.697l-.102-.981m-4.303-9l3.672-4.329l5.85 5.85l-4.308 3.654M3 3l18 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinSlashSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 14.5L3 21"/><path fill="currentColor" d="m5 9.485l9.193 9.193l1.697-1.698l-.102-.98l-8.112-8.11l-.979-.102zm10.157-6.813l5.85 5.85l-4.308 3.653L11.485 7z"/><path d="m3 3l18 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 14.5L3 21"/><path fill="currentColor" d="m5 9.485l9.193 9.193l1.697-1.697l-.393-3.787l5.51-4.673l-5.85-5.85l-4.674 5.51l-3.786-.393z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PineTree(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 2L7 6.643S10.042 7 12 7c1.958 0 5-.357 5-.357zM8.5 7L5 10.94S7.625 12 12 12s7-1.06 7-1.06L15.5 7"/><path d="M6.5 11.5L3 15.523S5.7 18 12 18s9-2.477 9-2.477L17.5 11.5M12 22v-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 14.5c-3-4.5 1.462-8 4.5-8c3.038 0 5.5 1.654 5.5 5.5c0 3.038-2 5-4 5s-3-2-2.5-5m.5-2L9 21.5"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PipeThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 20a6 6 0 1 1 0-12a6 6 0 0 1 0 12m.773-15.258a6 6 0 0 1 8.7 8.258M3 21l6.5-6.5M21 3l-1.5 1.5M6 9.5L10.5 5l.25-.25M14.5 18l4.719-4.719"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PizzaSlice(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m14 9.01l.01-.011M8 8.01l.01-.011M8 14.01l.01-.011"/><path d="M6 19L2.236 3.004a.6.6 0 0 1 .754-.713L19 7"/><path stroke-linecap="round" d="M22.198 8.425a1.75 1.75 0 0 0-3.396-.85c-.391 1.568-1.9 4.05-4.227 6.375c-2.3 2.301-5.148 4.194-7.968 4.845a1.75 1.75 0 1 0 .787 3.41c3.68-.849 7.082-3.206 9.656-5.78c2.549-2.549 4.54-5.568 5.148-8Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Planet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="8"/><path d="M17.5 6.348c2.297-.538 3.945-.476 4.338.312c.73 1.466-3.158 4.89-8.686 7.645c-5.529 2.757-10.603 3.802-11.334 2.336c-.392-.786.544-2.134 2.349-3.64"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlanetAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="8"/><path d="M19.812 12.99c1.813 1.51 2.755 2.864 2.362 3.651c-.731 1.467-5.805.42-11.333-2.336C5.312 11.55 1.423 8.126 2.154 6.66c.392-.786 2.033-.85 4.322-.315"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlanetSat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="8"/><path d="M17.5 6.348c2.297-.538 3.945-.476 4.338.312c.73 1.466-3.158 4.89-8.686 7.645c-5.529 2.757-10.603 3.802-11.334 2.336c-.392-.786.544-2.134 2.349-3.64"/><path stroke-linecap="round" stroke-linejoin="round" d="m9.5 10.51l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Planimetry(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11 16v-5H3m18-3h-6v2m0 8v3m-4-2v2m0-18v3m10 9h-6v-2m-4-4v2"/><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.906 4.537A.6.6 0 0 0 6 5.053v13.894a.6.6 0 0 0 .906.516l11.723-6.947a.6.6 0 0 0 0-1.032z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.906 4.537A.6.6 0 0 0 6 5.053v13.894a.6.6 0 0 0 .906.516l11.723-6.947a.6.6 0 0 0 0-1.032z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playlist(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2 11h14M2 17h11M2 5h18"/><path d="M20 18.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 0v-7.9a.6.6 0 0 1 .6-.6H22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaylistPlay(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 17.5L18.5 20v-5zM2 5h18M2 11h18M2 17h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaylistPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 18h2m2 0h-2m0 0v-2m0 2v2M2 11h18M2 17h12M2 5h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaystationGamepad(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.5 17.5c2.5 3.5 6.449.915 5.5-2.5c-1.425-5.129-2.2-7.984-2.603-9.492A2.032 2.032 0 0 0 18.438 4H5.562c-.918 0-1.718.625-1.941 1.515C2.78 8.863 2.033 11.802 1.144 15c-.948 3.415 3 6 5.5 2.5"/><path d="M16 4v2a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V4m0 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4m8 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlugTypeA(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10M9 10v4m6-4v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlugTypeC(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M8 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m8 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlugTypeG(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10m0-15v3m2 4h3M7 14h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlugTypeL(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 3H5.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6H10m1-14h2m-2 5h2m-2 5h2m1-14h4.4a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h6m6 0h-6m0 0V6m0 6v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h4m4 0h-4m0 0V8m0 4v4m0 6c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25M12.75 8a.75.75 0 0 0-1.5 0v3.25H8a.75.75 0 0 0 0 1.5h3.25V16a.75.75 0 0 0 1.5 0v-3.25H16a.75.75 0 0 0 0-1.5h-3.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h3m3 0h-3m0 0V9m0 3v3m9-11.4v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareDashed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h4m4 0h-4m0 0V8m0 4v4M7 4H4v3m0 4v2m7-9h2m-2 16h2m7-9v2m-3-9h3v3M7 20H4v-3m13 3h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zM12.75 9a.75.75 0 0 0-1.5 0v2.25H9a.75.75 0 0 0 0 1.5h2.25V15a.75.75 0 0 0 1.5 0v-2.25H15a.75.75 0 0 0 0-1.5h-2.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PngFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.5 15v-3m0 0V9h3v3zm6 3V9l3 6V9m6 0h-3v6h3v-2.4"/><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pocket(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 6v5a9 9 0 1 1-18 0V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2"/><path d="m8 10l4 4l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Podcast(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="iconoirPodcast0" d="M6 19a9.985 9.985 0 0 1-4-8C2 5.477 6.477 1 12 1s10 4.477 10 10a9.985 9.985 0 0 1-4 8"/></defs><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><use href="#iconoirPodcast0"/><use href="#iconoirPodcast0"/><path d="M7.528 15a6 6 0 1 1 8.944 0"/><path d="M12 13a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-1.924 3.283l.815-.543a2 2 0 0 1 2.218 0l.815.543a2 2 0 0 1 .863 1.993l-.509 3.053A2 2 0 0 1 12.307 23h-.612a2 2 0 0 1-1.973-1.671l-.508-3.053a2 2 0 0 1 .863-1.993"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pokeball(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path d="M2 12h7m6 0h7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PolarSh(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/><path d="M12 22c-2.21 0-4-4.477-4-10S9.79 2 12 2s4 4.477 4 10s-1.79 10-4 10"/><path d="M9 21c-3-1-4-5.389-4-8.5c0-3.111 1.5-7 5-9.5m5 0c3 1 4 5.389 4 8.5c0 3.111-1.5 7-5 9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Position(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14m0 0v2m-7-9H3m9-7V3m7 9h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PositionAlign(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 16.01l.01-.011M4 20.01l.01-.011M4 8.01l.01-.011M4 4.01l.01-.011M4 12.01l.01-.011m7.99.011l.01-.011M8 20.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M20 12.01l.01-.011M20 8.01l.01-.011M20 4.01l.01-.011M16 4.01l.01-.011M12 4.01l.01-.011M8 4.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Post(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2.906 17.505L5.337 3.718a2 2 0 0 1 2.317-1.623L19.472 4.18a2 2 0 0 1 1.622 2.317l-2.431 13.787a2 2 0 0 1-2.317 1.623L4.528 19.822a2 2 0 0 1-1.622-2.317Z"/><path stroke-linecap="round" d="m8.929 6.382l7.879 1.389m-8.574 2.55l7.879 1.39M7.54 14.26l4.924.869"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PostSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.784 1.356a2.75 2.75 0 0 0-3.186 2.231l-2.43 13.787a2.75 2.75 0 0 0 2.23 3.186l11.818 2.084a2.75 2.75 0 0 0 3.185-2.23l2.432-13.788a2.75 2.75 0 0 0-2.231-3.186zM9.06 5.643A.75.75 0 1 0 8.8 7.12l7.878 1.39a.75.75 0 0 0 .26-1.478zm-1.563 4.548a.75.75 0 0 1 .869-.608l7.878 1.389a.75.75 0 1 1-.26 1.477l-7.879-1.39a.75.75 0 0 1-.608-.868m.174 3.33a.75.75 0 1 0-.26 1.477l4.924.869a.75.75 0 1 0 .26-1.478z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Potion(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10 4h4v2.568c0 .258.17.487.412.579C22.938 10.37 20.908 22 15 22H9c-5.907 0-7.937-11.63.588-14.853a.629.629 0 0 0 .412-.58z"/><path d="M6 10h12"/><path stroke-linecap="round" d="M9 2h6"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.667 13L10 16h4l-1.667 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pound(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.667 13.6c-1.111 2.667-2.778 5.333-5 6.4h10.555S17.89 20 19 18.933M15.111 13.6H4m13.333-4.8c0-2.651-2.238-4.8-5-4.8c-2.761 0-5 2.149-5 4.8s2.239 4.8 5 4.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrecisionTool(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v2m0 8v2m-4-6H6m12 0h-2m-4 10c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Presentation(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 4.6v12.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6M8.5 21.5L12 18l3.5 3.5M12 2v2m-3 8v2m3-4v4m3-6v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12.53 17.47a.75.75 0 0 0-1.06 0l-3.5 3.5a.75.75 0 1 0 1.06 1.06L12 19.06l2.97 2.97a.75.75 0 1 0 1.06-1.06zM12 1.25a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0V2a.75.75 0 0 1 .75-.75"/><path d="M21.4 18.75a1.35 1.35 0 0 0 1.35-1.35V4.6a1.35 1.35 0 0 0-1.35-1.35H2.6A1.35 1.35 0 0 0 1.25 4.6v12.8c0 .746.604 1.35 1.35 1.35zM9.75 12a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0zM12 9.25a.75.75 0 0 1 .75.75v4a.75.75 0 0 1-1.5 0v-4a.75.75 0 0 1 .75-.75M15.75 8a.75.75 0 0 0-1.5 0v6a.75.75 0 0 0 1.5 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 13.01l.01-.011M7 17h10M6 10V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6V10m3 10.4V14a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v6.4a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrintingPage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M17.571 18H20.4a.6.6 0 0 0 .6-.6V11a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v6.4a.6.6 0 0 0 .6.6h2.829M8 7V3.6a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6V7"/><path d="M6.098 20.315L6.428 18l.498-3.485A.6.6 0 0 1 7.52 14h8.96a.6.6 0 0 1 .594.515L17.57 18l.331 2.315a.6.6 0 0 1-.594.685H6.692a.6.6 0 0 1-.594-.685Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17 10.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriorityDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848zM12 16l4-4m-4 4l-4-4.167M12 16V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriorityDownSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.955.893a1.35 1.35 0 0 0-1.91 0L.894 11.045a1.35 1.35 0 0 0 0 1.91l10.151 10.15a1.35 1.35 0 0 0 1.91 0l10.151-10.15a1.35 1.35 0 0 0 0-1.91zM12.53 16.53a.75.75 0 0 1-1.071-.01l-4-4.167a.75.75 0 1 1 1.082-1.04l2.709 2.823V7a.75.75 0 0 1 1.5 0v7.19l2.72-2.72a.75.75 0 1 1 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriorityHigh(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848zM12 8v4m0 4.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriorityHighSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.045.893a1.35 1.35 0 0 1 1.91 0l10.151 10.152a1.35 1.35 0 0 1 0 1.91l-10.151 10.15a1.35 1.35 0 0 1-1.91 0L.894 12.956a1.35 1.35 0 0 1 0-1.91zM12 7.25a.75.75 0 0 1 .75.75v4a.75.75 0 0 1-1.5 0V8a.75.75 0 0 1 .75-.75m.568 9.25a.75.75 0 0 0-1.115-1.003l-.01.011a.75.75 0 0 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriorityMedium(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848zM6 12h4m4 0h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriorityMediumSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.045.893a1.35 1.35 0 0 1 1.91 0l10.151 10.152a1.35 1.35 0 0 1 0 1.91l-10.151 10.15a1.35 1.35 0 0 1-1.91 0L.894 12.956a1.35 1.35 0 0 1 0-1.91zM5.25 12a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 0 1.5H6a.75.75 0 0 1-.75-.75m8.75-.75a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriorityUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848zM12 7l4 4m-4-4l-4 4.167M12 7v9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriorityUpSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.955.893a1.35 1.35 0 0 0-1.91 0L.894 11.045a1.35 1.35 0 0 0 0 1.91l10.151 10.15a1.35 1.35 0 0 0 1.91 0l10.151-10.15a1.35 1.35 0 0 0 0-1.91zM12.53 6.47a.75.75 0 0 0-1.071.01l-4 4.167a.75.75 0 1 0 1.082 1.04l2.709-2.823V16a.75.75 0 0 0 1.5 0V8.81l2.72 2.72a.75.75 0 1 0 1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrivacyPolicy(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 12V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H13M8 10h8M8 6h4m-4 8h3"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20m-.008 9.125l2.556.649c.266.068.453.31.445.584C22.821 22.116 19.5 23 19.5 23s-3.321-.884-3.493-6.642a.588.588 0 0 1 .445-.584l2.556-.649c.323-.082.661-.082.984 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrivateWifi(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 18.51l.01-.011M2 7c6-4.5 14-4.5 20 0M5 11c4-3 10-3 14 0M8.5 14.5c2.25-1.4 4.75-1.4 7 0m5.667 4h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProfileCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2"/><path d="M4.271 18.346S6.5 15.5 12 15.5s7.73 2.846 7.73 2.846M12 12a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Prohibition(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.141 5A9.97 9.97 0 0 0 12 2C6.477 2 2 6.477 2 12a9.968 9.968 0 0 0 2.859 7M19.14 5A9.967 9.967 0 0 1 22 12c0 5.523-4.477 10-10 10a9.97 9.97 0 0 1-7.141-3M19.14 5L4.86 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProjectCurveThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21c-4.97 0-9-1.79-9-4s4.03-4 9-4s9 1.79 9 4s-4.03 4-9 4m0-19a3 3 0 0 1 3 3v1H9V5a3 3 0 0 1 3-3M3.5 15.5l4-7m13 7l-4-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 14v4.4a.6.6 0 0 0 .6.6H10m9-5v4.4a.6.6 0 0 1-.6.6H14m0-14h4.4a.6.6 0 0 1 .6.6V10M4 10V5.6a.6.6 0 0 1 .6-.6H10m4 14v1a2 2 0 1 1-4 0v-1m-6-9h1a2 2 0 1 1 0 4H4m15-4h1a2 2 0 1 1 0 4h-1m-5-9V4a2 2 0 1 0-4 0v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 12v3M12 3v3m6 6v3m-6 3h9m-3 3h3M6 12h3M6 6.011L6.01 6M12 12.011l.01-.011M3 12.011L3.01 12M12 9.011L12.01 9M12 15.011l.01-.011M15 21.011l.01-.011m-3.01.011l.01-.011M21 12.011l.01-.011M21 15.011l.01-.011M18 6.011L18.01 6M9 3.6v4.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6m12 0v4.8a.6.6 0 0 1-.6.6h-4.8a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6M6 18.011L6.01 18M9 15.6v4.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6v-4.8a.6.6 0 0 1 .6-.6h4.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.9 8.08c0-4.773 7.5-4.773 7.5 0c0 3.409-3.409 2.727-3.409 6.818M12 19.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M10 12H5a1 1 0 0 1-1-1V7.5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1zm0 0c0 2.5-1 4-4 5.5M20 12h-5a1 1 0 0 1-1-1V7.5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1zm0 0c0 2.5-1 4-4 5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteMessage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.29V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7.961a2 2 0 0 0-1.561.75l-2.331 2.914A.6.6 0 0 1 3 20.29Z"/><path stroke-linecap="round" d="M10.5 10h-2a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1zm0 0c0 1-1 2-2 3m8-3h-2a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1zm0 0c0 1-1 2-2 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radiation(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20.662a9.955 9.955 0 0 1-5 1.337a9.954 9.954 0 0 1-5-1.337L10 16s1 .4 2 .4s2-.4 2-.4zm-.002-17.323A9.954 9.954 0 0 1 20.656 7a9.954 9.954 0 0 1 1.342 5l-5.537-.268s-.154-1.066-.654-1.932c-.5-.866-1.346-1.532-1.346-1.532zM1.998 12A9.954 9.954 0 0 1 3.34 7a9.954 9.954 0 0 1 3.658-3.66l2.537 4.928S8.69 8.934 8.19 9.8s-.654 1.932-.654 1.932zM12 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radius(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10m7-10l-3-3m3 3l-3 3m3-3h-7"/><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rain(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14v2m0 4v2m-4-4v2m8-2v2m4-2.393c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RawFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M16.5 9v6l1.5-3l1.5 3V9m-9 6v-3m0 0v-1.5A1.5 1.5 0 0 1 12 9v0a1.5 1.5 0 0 1 1.5 1.5V12m-3 0h3m0 0v3m-9 0V9h2.4a.6.6 0 0 1 .6.6v.9A1.5 1.5 0 0 1 6 12v0"/><path stroke-linejoin="round" d="M4.5 12H6v0a1.5 1.5 0 0 1 1.5 1.5V15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiveDollars(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 8.23c-.8-.737-2.207-1.25-3.5-1.282M3 15.23c.752.925 2.15 1.453 3.5 1.498m0-9.781c-1.539-.038-2.917.604-2.917 2.36c0 3.23 6.417 1.615 6.417 4.846c0 1.842-1.708 2.634-3.5 2.575m0-9.781V5m0 11.729V19M21 12h-8m0 0l3.84-4M13 12l3.84 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiveEuros(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12h-8m0 0l3.84-4M13 12l3.84 4M11 7.503A4.746 4.746 0 0 0 8.87 7C6.18 7 4 9.239 4 12s2.18 5 4.87 5a4.73 4.73 0 0 0 2.13-.503M3 11h6m-6 2h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceivePounds(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12h-8m0 0l3.84-4M13 12l3.84 4M7 13c-.667 1.667-1.667 3.333-3 4h6.333s1 0 1.667-.667M9.667 13H3m8-3a3 3 0 1 0-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiveYens(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12h-8m0 0l3.84-4M13 12l3.84 4M3 13h8M3 7l4 5.5M11 7l-4 5.5m0 0V18m-4-3h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.5 8H9s0 0 0 0s-5 0-5 4.706C4 18 9 18 9 18h8.714"/><path d="M16.5 11.5L20 8l-3.5-3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedoAction(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 5v6m-3.5-3H9s0 0 0 0s-5 0-5 4.706C4 18 9 18 9 18h8.714"/><path d="M12.5 11.5L16 8l-3.5-3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedoCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 10.625H9.8s0 0 0 0s-2.8 0-2.8 3C7 17 9.8 17 9.8 17h.8"/><path d="m13.5 14l3.5-3.375L13.5 7"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedoCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12m16.321-.89a.76.76 0 0 1-.05.055l-3.5 3.375a.75.75 0 1 1-1.042-1.08l2.163-2.085H9.789a2.454 2.454 0 0 0-.308.035a2.49 2.49 0 0 0-.737.25a1.744 1.744 0 0 0-.672.624c-.178.288-.322.71-.322 1.341c0 1.438.567 2.032 1.029 2.312a2.253 2.253 0 0 0 1.014.313h.807a.75.75 0 0 1 0 1.5h-.8a3.72 3.72 0 0 1-1.798-.53c-.933-.565-1.752-1.658-1.752-3.595c0-.87.202-1.572.545-2.128c.341-.554.796-.92 1.238-1.157A4.007 4.007 0 0 1 9.8 9.875h5.433L12.96 7.521a.75.75 0 0 1 1.08-1.042l3.498 3.623a.748.748 0 0 1 .033 1.009" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reduce(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 20l5-5m0 0v4m0-4H5M20 4l-5 5m0 0V5m0 4h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.888 13.5C21.164 18.311 17.013 22 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2c4.1 0 7.625 2.468 9.168 6"/><path d="M17 8h4.4a.6.6 0 0 0 .6-.6V3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.583 9.667C15.81 8.097 14.043 7 11.988 7C9.388 7 7.25 8.754 7 11"/><path stroke-linecap="round" stroke-linejoin="round" d="M14.494 9.722H16.4a.6.6 0 0 0 .6-.6V7.5m-9.583 6.167C8.191 15.629 9.957 17 12.012 17c2.6 0 4.736-2.193 4.988-5"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.506 13.622H7.6a.6.6 0 0 0-.6.6V16.4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12m10.738-4.25c-2.287 0-4.04 1.532-4.243 3.334a.75.75 0 0 1-1.49-.168c.301-2.69 2.821-4.666 5.733-4.666c1.67 0 3.198.644 4.262 1.697V7.5a.75.75 0 0 1 1.5 0v1.622a1.35 1.35 0 0 1-1.35 1.35h-1.906a.75.75 0 0 1 0-1.5h.658c-.77-.74-1.89-1.222-3.164-1.222m.024 8.5c2.146 0 4.018-1.828 4.24-4.317a.75.75 0 0 1 1.495.134c-.28 3.126-2.682 5.683-5.735 5.683c-1.708 0-3.219-.807-4.262-2.062v.712a.75.75 0 0 1-1.5 0v-2.177c0-.746.604-1.35 1.35-1.35h1.906a.75.75 0 0 1 0 1.5h-.873c.79 1.158 2.027 1.877 3.38 1.877" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshDouble(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.168 8A10.003 10.003 0 0 0 12 2c-5.185 0-9.45 3.947-9.95 9"/><path d="M17 8h4.4a.6.6 0 0 0 .6-.6V3M2.881 16c1.544 3.532 5.068 6 9.168 6c5.186 0 9.45-3.947 9.951-9"/><path d="M7.05 16h-4.4a.6.6 0 0 0-.6.6V21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReloadWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M11 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v7"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011m10.657 11.668C21.047 15.097 19.635 14 17.99 14c-1.758 0-3.252 1.255-3.793 3"/><path stroke-linejoin="round" d="M19.995 16.772H21.4a.6.6 0 0 0 .6-.6V14.55m-7.666 4.783C14.953 20.903 16.366 22 18.01 22c1.758 0 3.252-1.255 3.793-3"/><path stroke-linejoin="round" d="M16.005 19.228H14.6a.6.6 0 0 0-.6.6v1.622"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReminderHandGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m17.5 12l2.004 2.672a2 2 0 0 1-.126 2.552l-3.783 4.127A1.998 1.998 0 0 1 14.12 22H9.5c-2.358 0-3.622-2.575-3.982-3.93a.55.55 0 0 1-.018-.143V9.43c0-2.286 3-2.286 3 0V10"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.5 10V8.286c0-2.286-3-2.286-3 0V10m6 0V7.5c0-2.286-3-2.286-3 0c0 0 0 0 0 0V10m3 0V3.499A1.5 1.5 0 0 1 16 2v0a1.5 1.5 0 0 1 1.5 1.5V15"/><path d="M17.563 6.5h2.062C20.5 6.5 21 6.078 21 5.25C21 4.422 20.5 4 19.625 4H12.25C11.56 4 11 4.56 11 5.25v.25a1 1 0 0 0 1 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 17H8c-1.667 0-5-1-5-5s3.333-5 5-5h8c1.667 0 5 1 5 5c0 1.494-.465 2.57-1.135 3.331"/><path d="M14.5 14.5L17 17l-2.5 2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RepeatOnce(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 17H8c-1.667 0-5-1-5-5m5-5h8c1.667 0 5 1 5 5c0 1.494-.465 2.57-1.135 3.331"/><path d="M14.5 14.5L17 17l-2.5 2.5M4 8V3L2 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 10h14c8 0 8 11 0 11M2 10l7-7m-7 7l7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplyToMessage(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 8l5 3l5-3"/><path d="M10 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v6.857"/><path stroke-linejoin="round" d="M13 17.111h6.3c3.6 0 3.6 4.889 0 4.889M13 17.111L16.15 14M13 17.111l3.15 3.111"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReportColumns(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 7.4V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm11 13v-3.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm0-8V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v8.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 8v-8.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v8.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reports(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M9 21h6m-6 0v-5m0 5H3.6a.6.6 0 0 1-.6-.6v-3.8a.6.6 0 0 1 .6-.6H9m6 5V9m0 12h5.4a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6h-4.8a.6.6 0 0 0-.6.6V9m0 0H9.6a.6.6 0 0 0-.6.6V16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReportsSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><rect width="5" height="18" x="16" y="3" rx="2"/><rect width="5" height="12" x="9.5" y="9" rx="2"/><rect width="5" height="5" x="3" y="16" rx="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repository(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 19V5a2 2 0 0 1 2-2h13.4a.6.6 0 0 1 .6.6v13.114"/><path stroke-linejoin="round" d="M15 17v5l2.5-1.6L20 22v-5"/><path d="M6 17h14"/><path stroke-linejoin="round" d="M6 17a2 2 0 1 0 0 4h5.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Restart(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.677 20.567C2.531 18.021.758 12.758 2.717 8.144C4.875 3.06 10.745.688 15.829 2.846c5.084 2.158 7.456 8.029 5.298 13.113a9.954 9.954 0 0 1-3.962 4.608"/><path d="M17 16v4.4a.6.6 0 0 0 .6.6H22m-10 1.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rewind(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.044 5.704a.6.6 0 0 1 .956.483v11.626a.6.6 0 0 1-.956.483l-7.889-5.813a.6.6 0 0 1 0-.966zm-11 0a.6.6 0 0 1 .956.483v11.626a.6.6 0 0 1-.956.483l-7.888-5.813a.6.6 0 0 1 0-.966z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.044 5.704a.6.6 0 0 1 .956.483v11.626a.6.6 0 0 1-.956.483l-7.889-5.813a.6.6 0 0 1 0-.966zm-11 0a.6.6 0 0 1 .956.483v11.626a.6.6 0 0 1-.956.483l-7.888-5.813a.6.6 0 0 1 0-.966z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rhombus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RhombusArrowRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M16 11h-6c-1 0-2 1-2 2v1m8-3l-2-2m2 2l-2 2"/><path d="M1.424 11.576L11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RhombusArrowRightSolidSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.955.893a1.35 1.35 0 0 0-1.91 0L.894 11.045a1.35 1.35 0 0 0 0 1.91l10.151 10.15a1.35 1.35 0 0 0 1.91 0l10.151-10.15a1.35 1.35 0 0 0 0-1.91zM16.53 11.53a.75.75 0 0 0 0-1.06l-2-2a.75.75 0 1 0-1.06 1.06l.72.72H10c-.744 0-1.425.364-1.905.845c-.48.48-.845 1.161-.845 1.905v1a.75.75 0 0 0 1.5 0v-1c0-.256.136-.575.405-.845c.27-.27.589-.405.845-.405h4.19l-.72.72a.75.75 0 1 0 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rings(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 8a6 6 0 1 0 0 12A6 6 0 0 0 8 8m0 0V3"/><path d="M16 8a6 6 0 1 0 0 12a6 6 0 0 0 0-12m0 0V3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.061 10.404L14 17h-4l-2.061-6.596a6 6 0 0 1 .998-5.484l2.59-3.315a.6.6 0 0 1 .946 0l2.59 3.315a6 6 0 0 1 .998 5.484M10 20c0 2 2 3 2 3s2-1 2-3m-5.5-7.5C5 15 7 19 7 19l3-2m5.931-4.5c3.5 2.5 1.5 6.5 1.5 6.5l-3-2"/><path d="M12 11a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rook(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 16h10m-8-5h6m-5-7v2m4-2v2m3.4 3H6.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6v3.8a.6.6 0 0 1-.6.6Zm.501 12H6.099a.615.615 0 0 1-.521-.932C6.792 18.06 9.5 13.328 9.5 11V9.6a.6.6 0 0 1 .6-.6h3.8a.6.6 0 0 1 .6.6V11c0 2.327 2.708 7.061 3.922 9.068a.615.615 0 0 1-.521.932Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateCameraLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.05 3v4.497c0 .278.226.503.504.503v0c.2 0 .38-.119.466-.3A10.001 10.001 0 0 1 12.05 2c5.186 0 9.45 3.947 9.951 9m0 10v-4.497a.503.503 0 0 0-.503-.503v0a.52.52 0 0 0-.465.3A10.001 10.001 0 0 1 12 22c-5.185 0-9.448-3.947-9.95-9"/><path d="M6 16.4V9.394a.6.6 0 0 1 .6-.6h1.173a.6.6 0 0 0 .504-.275l1.446-2.244A.6.6 0 0 1 10.227 6h3.546a.6.6 0 0 1 .504.275l1.446 2.244a.6.6 0 0 0 .504.275H17.4a.6.6 0 0 1 .6.6V16.4a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6"/><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateCameraRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22.003 3v4.497A.503.503 0 0 1 21.5 8v0c-.2 0-.38-.119-.466-.3A10.001 10.001 0 0 0 12.003 2c-5.186 0-9.45 3.947-9.95 9"/><path d="M6 16.4V9.394a.6.6 0 0 1 .6-.6h1.173a.6.6 0 0 0 .504-.275l1.446-2.244A.6.6 0 0 1 10.227 6h3.546a.6.6 0 0 1 .504.275l1.446 2.244a.6.6 0 0 0 .504.275H17.4a.6.6 0 0 1 .6.6V16.4a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6"/><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-9.95 7v-4.497c0-.278.226-.503.504-.503v0c.2 0 .38.119.466.3a10.001 10.001 0 0 0 9.03 5.7c5.186 0 9.45-3.947 9.951-9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundFlask(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M19 15H5"/><path stroke-linecap="round" d="M16 4H8m7 .5v4.253c0 .763.445 1.445 1.078 1.871C18.287 12.11 20 14.617 20 17.462c0 .812-.114 1.596-.325 2.338c-.215.75-.945 1.2-1.726 1.2H6.051c-.78 0-1.511-.45-1.726-1.2A8.505 8.505 0 0 1 4 17.462c0-2.845 1.713-5.353 3.922-6.838C8.555 10.198 9 9.516 9 8.754V4.5m4 2.51l.01-.011M11 2.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundedMirror(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 10v4a8 8 0 1 1-16 0v-4a8 8 0 1 1 16 0m-2.5-5.5L13 8m6-1l-7.5 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssFeed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19c0-4.2-2.8-7-7-7m14 7c0-8.4-5.6-14-14-14m0 14.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssFeedTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17c0-3-2-5-5-5m10 5c0-6-4-10-10-10m0 10.01l.01-.011"/><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RubikCube(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6M9 3v18M3 9h18M3 15h18M15 3v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 7V2.6a.6.6 0 0 0-.6-.6H8.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6h6.8a.6.6 0 0 0 .6-.6V17m0-10h-3m3 0v5m0 0h-3m3 0v5m0 0h-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RulerArrows(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.4 22H8.6a.6.6 0 0 1-.6-.6V2.6a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6v18.8a.6.6 0 0 1-.6.6m.6-5h-3m3-10h-3m0 5h10m0 0l-2 2m2-2l-2-2M1 12l2-2m-2 2l2 2m-2-2h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RulerCombine(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 21.4V2.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H10.6a.6.6 0 0 0-.6.6v10.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6M16 10V7m-6 3V7m0 9H7m3-6H7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RulerMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 5h6M11 7V2.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6h6.8a.6.6 0 0 0 .6-.6V17m0-10H8m3 0v5m0 0H8m3 0v5m0 0H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RulerPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 5h3m3 0h-3m0 0V2m0 3v3m-7-1V2.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6h6.8a.6.6 0 0 0 .6-.6V17m0-10H8m3 0v5m0 0H8m3 0v5m0 0H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Running(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-2.387 1.267l-3.308 4.135l4.135 4.135l-2.067 4.55"/><path d="m6.41 9.508l3.387-3.309l2.816 2.068l2.895 3.308h3.722M8.892 15.71l-1.241.827H4.343"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safari(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.586 10.586L16.95 7.05l-3.536 6.364m-2.828-2.828L7.05 16.95l6.364-3.536m-2.828-2.828l2.828 2.828"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10m7-10h-1M6 12H5m7-7v1m0 12v1M7.05 7.05l.707.707m8.486 8.486l.707.707"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safe(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6m8-1v-4m-5.5-.5l1-1m-6 1l-1-1m0 7l1-1m6 1l-1-1M2 8h1M2 6h1m0 10H2m1 2H2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafeArrowLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 3h7a2 2 0 0 1 2 2v.5M12 21h7a2 2 0 0 0 2-2v-.5M8.5 15C7.672 15 7 13.657 7 12s.672-3 1.5-3s1.5 1.343 1.5 3s-.672 3-1.5 3m1-5.5l1-1m-3 1l-1-1m0 7l1-1m3 1l-1-1M2 8h1M2 6h1m0 10H2m1 2H2m20-6h-7m0 0l3.5-3.5M15 12l3.5 3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafeArrowRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 3h7a2 2 0 0 1 2 2v.5M12 21h7a2 2 0 0 0 2-2v-.5M8.5 15C7.672 15 7 13.657 7 12s.672-3 1.5-3s1.5 1.343 1.5 3s-.672 3-1.5 3m1-5.5l1-1m-3 1l-1-1m0 7l1-1m3 1l-1-1M2 8h1M2 6h1m0 10H2m1 2H2m13-6h7m0 0l-3.5-3.5M22 12l-3.5 3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafeOpen(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path d="M13 3h6a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-6"/><path stroke-linecap="round" stroke-linejoin="round" d="M7.5 15C6.672 15 6 13.657 6 12s.672-3 1.5-3S9 10.343 9 12s-.672 3-1.5 3m5.5-1v-4m-4.5-.5l1-1m-3 1l-1-1m0 7l1-1m3 1l-1-1M2 8h1M2 6h1m0 10H2m1 2H2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sandals(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 7s.5-4-4-4s-4 4-4 4m8 0h-8m8 0l-.214 3M14 7l.214 3m7.572 0l-.587 8.214A3 3 0 0 1 18.207 21h-.414a3 3 0 0 1-2.992-2.786L14.214 10m7.572 0h-7.572M10 7s.5-4-4-4s-4 4-4 4m8 0H2m8 0l-.214 3M2 7l.214 3m7.572 0l-.587 8.214A3 3 0 0 1 6.207 21h-.414a3 3 0 0 1-2.992-2.786L2.214 10m7.572 0H2.214"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScaleFrameEnlarge(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 13.6V21H3.6a.6.6 0 0 1-.6-.6V13h7.4a.6.6 0 0 1 .6.6m0 7.4h3M3 13v-3m3-7H3.6a.6.6 0 0 0-.6.6V6m11-3h-4m11 7v4M18 3h2.4a.6.6 0 0 1 .6.6V6m-3 15h2.4a.6.6 0 0 0 .6-.6V18m-10-8h3v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScaleFrameReduce(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11 15v-1.4a.6.6 0 0 0-.6-.6H9m-3 0H3m8 5v3"/><path stroke-miterlimit="1.5" stroke-width="1.499" d="M20.4 3H3.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6V3.6a.6.6 0 0 0-.6-.6"/><path d="M16 11h-3V8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScanBarcode(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12V6h1m-1 6h1V6m-1 12v-3h1m0 0v3h-1M7 6v6m0 3v3m7-12v6m0 3v3m3-12v6m0 3v3M6 3H3v3m-1 6h20m-4-9h3v3M6 21H3v-3m15 3h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScanQrCode(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6.6v1.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6V6.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6M6 12h3m6 0v3m-3 3h3m-3-5.989l.01-.011m5.99.011l.01-.011M12 15.011l.01-.011m5.99.011l.01-.011M18 18.011l.01-.011M12 9.011L12.01 9M12 6.011L12.01 6M9 15.6v1.8a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6v-1.8a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6m9-9v1.8a.6.6 0 0 1-.6.6h-1.8a.6.6 0 0 1-.6-.6V6.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6M18 3h3v3m-3 15h3v-3M6 3H3v3m3 15H3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scanning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 3H3v3m-1 6h20M9 19v-4m3 1v-1m3 2v-2m-3 6v-3m6-15h3v3M6 21H3v-3m15 3h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scarf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 11H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v16m-3 0v-2M15 3v18m0-14H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissor(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.236 7a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4m0 0L20 18M7.236 17a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4m0 0L20 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScissorAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.236 8a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4m0 0L20 16m-9.764 0a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4m0 0L20 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Screenshot(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 21.4v-7.006a.6.6 0 0 1 .6-.6h1.173a.6.6 0 0 0 .504-.275l1.446-2.244a.6.6 0 0 1 .504-.275h3.546a.6.6 0 0 1 .504.275l1.446 2.244a.6.6 0 0 0 .504.275H21.4a.6.6 0 0 1 .6.6V21.4a.6.6 0 0 1-.6.6H10.6a.6.6 0 0 1-.6-.6"/><path d="M16 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4M3 18v3h2.5M3 9.5v5M3 6V3h3m3.5 0h5M18 3h3v2.5m0 4.5V8.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeaAndSun(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 15c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3M3 20c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3m-2-10a7 7 0 1 0-14 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeaWaves(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 10c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3M3 17c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 17l4 4M3 11a8 8 0 1 0 16 0a8 8 0 0 0-16 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchEngine(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M13.856 13.85a3.429 3.429 0 1 0-4.855-4.842a3.429 3.429 0 0 0 4.855 4.842m0 0L16 16"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v9"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011m9.114 15.12a3 3 0 1 0-4.248-4.237a3 3 0 0 0 4.248 4.237m0 0L22 22"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecureWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v7"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011m7.982 9.126l2.556.649c.266.068.453.31.445.584C21.821 21.116 18.5 22 18.5 22s-3.321-.884-3.493-6.642a.588.588 0 0 1 .445-.584l2.556-.649c.323-.082.661-.082.984 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecurityPass(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 11l3 3l8-8"/><path d="M20 12a8 8 0 1 1-5.3-7.533"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectEdgeThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524M12 21v-9"/><path fill="currentColor" d="M12.5 11v10a.5.5 0 0 1-1 0V11a.5.5 0 0 1 1 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectFaceThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524"/><path stroke-linecap="round" d="m3.528 7.294l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21v-9"/><path fill="currentColor" d="m11.691 11.829l-7.8-4.334A.6.6 0 0 0 3 8.02v8.627a.6.6 0 0 0 .309.525l7.8 4.333A.6.6 0 0 0 12 20.98v-8.627a.6.6 0 0 0-.309-.524Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectPointThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M14 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10"/><path stroke-linecap="round" stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011"/><path d="M22.082 18.365c.494.304.464 1.043-.045 1.1l-2.566.292l-1.152 2.312c-.228.458-.933.234-1.05-.334l-1.255-6.116c-.098-.48.333-.782.75-.525z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectiveTool(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Send(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 12L3 20l3.563-8L3 4zM6.5 12H22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendDiagonal(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22.153 3.553L11.176 21.004l-1.67-8.596L2 7.898zM9.456 12.444l12.696-8.89"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendDollars(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 8.23c-.8-.737-2.207-1.25-3.5-1.282M3 15.23c.752.925 2.15 1.453 3.5 1.498m0-9.781c-1.539-.038-2.917.604-2.917 2.36c0 3.23 6.417 1.615 6.417 4.846c0 1.842-1.708 2.634-3.5 2.575m0-9.781V5m0 11.729V19m6.5-7h8m0 0l-3.84-4M21 12l-3.84 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendEuros(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 12h8m0 0l-3.84-4M21 12l-3.84 4M11 7.503A4.746 4.746 0 0 0 8.87 7C6.18 7 4 9.239 4 12s2.18 5 4.87 5a4.73 4.73 0 0 0 2.13-.503M3 11h6m-6 2h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendMail(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m9 9l4.5 3L18 9M3 13.5h2m-4-3h4"/><path d="M5 7.5V7a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendPounds(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 12h8m0 0l-3.84-4M21 12l-3.84 4M7 13c-.667 1.667-1.667 3.333-3 4h6.333s1 0 1.667-.667M9.667 13H3m8-3a3 3 0 1 0-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendYens(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 12h8m0 0l-3.84-4M21 12l-3.84 4M3 13h8M3 7l4 5.5M11 7l-4 5.5m0 0V18m-4-3h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m6 18.01l.01-.011M6 6.01l.01-.011"/><path d="M2 9.4V2.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Zm0 12v-6.8a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerConnection(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 19h9m9 0h-9m0 0v-6m0 0h6V5H6v8zM9 9.01l.01-.011M12 9.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerConnectionSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.25 5A.75.75 0 0 1 6 4.25h12a.75.75 0 0 1 .75.75v8a.75.75 0 0 1-.75.75h-5.25v4.5H21a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1 0-1.5h8.25v-4.5H6a.75.75 0 0 1-.75-.75zm4.262 3.442A.75.75 0 0 1 9.567 9.5l-.01.01a.75.75 0 1 1-1.114-1.003l.01-.01a.75.75 0 0 1 1.059-.056M12.568 9.5a.75.75 0 1 0-1.115-1.004l-.01.011a.75.75 0 1 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.6 13.25a1.35 1.35 0 0 0-1.35 1.35v6.8c0 .746.604 1.35 1.35 1.35h18.8a1.35 1.35 0 0 0 1.35-1.35v-6.8a1.35 1.35 0 0 0-1.35-1.35zm3.967 5.25a.75.75 0 0 0-1.114-1.003l-.01.011a.75.75 0 0 0 1.114 1.004zM2.6 1.25A1.35 1.35 0 0 0 1.25 2.6v6.8c0 .746.604 1.35 1.35 1.35h18.8a1.35 1.35 0 0 0 1.35-1.35V2.6a1.35 1.35 0 0 0-1.35-1.35zM6.567 6.5a.75.75 0 0 0-1.114-1.003l-.01.011a.75.75 0 1 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path d="m19.622 10.395l-1.097-2.65L20 6l-2-2l-1.735 1.483l-2.707-1.113L12.935 2h-1.954l-.632 2.401l-2.645 1.115L6 4L4 6l1.453 1.789l-1.08 2.657L2 11v2l2.401.656L5.516 16.3L4 18l2 2l1.791-1.46l2.606 1.072L11 22h2l.604-2.387l2.651-1.098C16.697 18.832 18 20 18 20l2-2l-1.484-1.75l1.098-2.652l2.386-.62V11z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsProfiles(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11.607 2.342a.6.6 0 0 1 .787 0l1.948 1.692a.6.6 0 0 0 .445.145l2.572-.224a.6.6 0 0 1 .636.463l.582 2.514a.6.6 0 0 0 .275.38l2.212 1.33a.6.6 0 0 1 .243.748l-1.008 2.376a.6.6 0 0 0 0 .468l1.008 2.376a.6.6 0 0 1-.243.749l-2.212 1.33a.6.6 0 0 0-.275.379l-.582 2.514a.6.6 0 0 1-.636.463l-2.572-.224a.6.6 0 0 0-.445.144l-1.949 1.693a.6.6 0 0 1-.787 0l-1.948-1.693a.6.6 0 0 0-.445-.144l-2.572.224a.6.6 0 0 1-.636-.463l-.582-2.514a.6.6 0 0 0-.275-.38l-2.212-1.33a.6.6 0 0 1-.243-.748l1.008-2.376a.6.6 0 0 0 0-.468L2.693 9.39a.6.6 0 0 1 .243-.749l2.212-1.33a.6.6 0 0 0 .275-.379l.582-2.514a.6.6 0 0 1 .636-.463l2.572.224a.6.6 0 0 0 .445-.145z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 13l2 2l5-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAndroid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 22a3 3 0 1 0 0-6a3 3 0 0 0 0 6m0-14a3 3 0 1 0 0-6a3 3 0 0 0 0 6M6 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path d="m15.5 6.5l-7 4m0 3l7 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAndroidSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18 22a3 3 0 1 0 0-6a3 3 0 0 0 0 6m0-14a3 3 0 1 0 0-6a3 3 0 0 0 0 6M6 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path d="m15.5 6.5l-7 4m0 3l7 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareIos(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13v6a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-6m8 2V3m0 0L8.5 6.5M12 3l3.5 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldAlert(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7v5m0 4.01l.01-.011M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.571 8l-.44-3.084A1 1 0 0 1 3.904 3.8l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8a1 1 0 0 1 .773 1.117L20.43 8M3.57 8h16.86M3.57 8c.309 2.16.69 4.822 1 7m15.86-7c-.309 2.16-.69 4.822-1 7m0 0L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18l-.429-3m14.858 0H4.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldBroken(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.5 7L9 12h6l-2.5 5"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.5 11.5l3 3l5-5"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldDownload(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v7m0 0l3-3m-3 3l-3-3m-4 6L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldEye(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 9s1-1 4-1s4 1 4 1"/><path fill="currentColor" d="M12 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldLoading(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 10.01l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldPlusIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h3m3 0h-3m0 0V9m0 3v3m-7 3L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldQuestion(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 9c0-3.5 5.5-3.5 5.5 0c0 2.5-2.5 2-2.5 5m0 4.01l.01-.011"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldSearch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.5 13l1.5 1.5M9 11a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0"/><path d="M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldUpload(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 15V8m0 0l3 3m-3-3l-3 3m-4 7L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.871 14.121L11.993 12m2.121-2.121L11.993 12m0 0L9.87 9.879M11.993 12l2.121 2.121M5 18L3.13 4.913a.996.996 0 0 1 .774-1.114l7.662-1.703a2 2 0 0 1 .868 0L20.096 3.8c.51.113.848.596.774 1.114L19 18c-.07.495-.5 3.5-7 3.5S5.07 18.495 5 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shirt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 4h3s0 3 3 3s3-3 3-3h3m0 7v8.4a.6.6 0 0 1-.6.6H6.6a.6.6 0 0 1-.6-.6V11m12-7l4.443 1.777a.6.6 0 0 1 .334.78l-1.626 4.066a.6.6 0 0 1-.557.377H18M6 4L1.557 5.777a.6.6 0 0 0-.334.78l1.626 4.066a.6.6 0 0 0 .557.377H6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShirtTankTop(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 21H6s1.66-4.825 1.5-8c-.1-1.989-1.524-3.079-1-5c.23-.842 1-2 1-2S9 7 12 7s4.5-1 4.5-1s.77 1.158 1 2c.524 1.921-.9 3.011-1 5c-.16 3.175 1.5 8 1.5 8M7.5 6V3m9 3V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shop(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 10v9a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-9"/><path stroke-miterlimit="16" d="M14.833 21v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v6"/><path d="m21.818 9.364l-1.694-5.929A.6.6 0 0 0 19.547 3H15.5l.475 5.704a.578.578 0 0 0 .278.45c.39.233 1.152.663 1.747.846c1.016.313 2.5.2 3.346.096a.57.57 0 0 0 .472-.732Z"/><path d="M14 10c.568-.175 1.288-.574 1.69-.812a.578.578 0 0 0 .28-.549L15.5 3h-7l-.47 5.639a.578.578 0 0 0 .28.55c.402.237 1.122.636 1.69.811c1.493.46 2.507.46 4 0Z"/><path d="m3.876 3.435l-1.694 5.93a.57.57 0 0 0 .472.73c.845.105 2.33.217 3.346-.095c.595-.183 1.358-.613 1.747-.845a.578.578 0 0 0 .278-.451L8.5 3H4.453a.6.6 0 0 0-.577.435Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopFourTiles(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20.485 3h-3.992l.5 5s1 1 2.5 1a3.23 3.23 0 0 0 2.139-.806a.503.503 0 0 0 .15-.465L21.076 3.5a.6.6 0 0 0-.591-.5Z"/><path d="m16.493 3l.5 5s-1 1-2.5 1s-2.5-1-2.5-1V3z"/><path d="M11.993 3v5s-1 1-2.5 1s-2.5-1-2.5-1l.5-5z"/><path d="M7.493 3H3.502a.6.6 0 0 0-.592.501L2.205 7.73a.504.504 0 0 0 .15.465c.328.29 1.061.806 2.138.806c1.5 0 2.5-1 2.5-1z"/><path d="M3 9v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9"/><path stroke-miterlimit="16" d="M14.833 21v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopFourTilesWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 9v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9"/><path d="M20.485 3h-3.992l.5 5s1 1 2.5 1a3.23 3.23 0 0 0 2.139-.806a.503.503 0 0 0 .15-.465L21.076 3.5a.6.6 0 0 0-.591-.5Z"/><path d="m16.493 3l.5 5s-1 1-2.5 1s-2.5-1-2.5-1V3z"/><path d="M11.993 3v5s-1 1-2.5 1s-2.5-1-2.5-1l.5-5z"/><path d="M7.493 3H3.502a.6.6 0 0 0-.592.501L2.205 7.73a.504.504 0 0 0 .15.465c.328.29 1.061.806 2.138.806c1.5 0 2.5-1 2.5-1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m21.818 9.364l-1.694-5.929A.6.6 0 0 0 19.547 3H15.5l.475 5.704a.578.578 0 0 0 .278.45c.39.233 1.152.663 1.747.846c1.016.313 2.5.2 3.346.096a.57.57 0 0 0 .472-.732Z"/><path d="M14 10c.568-.175 1.288-.574 1.69-.812a.578.578 0 0 0 .28-.549L15.5 3h-7l-.47 5.639a.578.578 0 0 0 .28.55c.402.237 1.122.636 1.69.811c1.493.46 2.507.46 4 0Z"/><path d="m3.876 3.435l-1.694 5.93a.57.57 0 0 0 .472.73c.845.105 2.33.217 3.346-.095c.595-.183 1.358-.613 1.747-.845a.578.578 0 0 0 .278-.451L8.5 3H4.453a.6.6 0 0 0-.577.435Z"/><path d="M3 10v9a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.26 9.696l1.385 9A2 2 0 0 1 18.67 21H5.33a2 2 0 0 1-1.977-2.304l1.385-9A2 2 0 0 1 6.716 8h10.568a2 2 0 0 1 1.977 1.696M14 5a2 2 0 1 0-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagArrowDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 21h2.169a2 2 0 0 0 1.976-2.304l-1.384-9A2 2 0 0 0 17.284 8H6.716a2 2 0 0 0-1.977 1.696l-1.385 9A2 2 0 0 0 5.331 21H7.5m4.5-9v7m0 0l3-3m-3 3l-3-3m5-11a2 2 0 1 0-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagArrowUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 21h2.169a2 2 0 0 0 1.976-2.304l-1.384-9A2 2 0 0 0 17.284 8H6.716a2 2 0 0 0-1.977 1.696l-1.385 9A2 2 0 0 0 5.331 21H7.5m4.5-2v-7m0 0l3 3m-3-3l-3 3m5-10a2 2 0 1 0-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20 14.5l-.74-4.804A2 2 0 0 0 17.285 8H6.716a2 2 0 0 0-1.977 1.696l-1.385 9A2 2 0 0 0 5.331 21H12"/><path d="m14 19l3 3l5-5M14 5a2 2 0 1 0-4 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagMinus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.26 9.696l1.385 9A2 2 0 0 1 18.67 21H5.33a2 2 0 0 1-1.977-2.304l1.385-9A2 2 0 0 1 6.716 8h10.568a2 2 0 0 1 1.977 1.696M14 5a2 2 0 1 0-4 0M8.992 15h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.26 9.696l1.385 9A2 2 0 0 1 18.67 21H5.33a2 2 0 0 1-1.977-2.304l1.385-9A2 2 0 0 1 6.716 8h10.568a2 2 0 0 1 1.977 1.696M14 5a2 2 0 1 0-4 0M8.992 15h3m3 0h-3m0 0v-3m0 3v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagPocket(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.26 9.696l1.385 9A2 2 0 0 1 18.67 21H5.33a2 2 0 0 1-1.977-2.304l1.385-9A2 2 0 0 1 6.716 8h10.568a2 2 0 0 1 1.977 1.696M9 11v7m6-7v7M14 5a2 2 0 1 0-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 14.5l-.74-4.804A2 2 0 0 0 17.285 8H6.716a2 2 0 0 0-1.977 1.696l-1.385 9A2 2 0 0 0 5.331 21H12m5.5-4v2m0 3.01l.01-.011M14 5a2 2 0 1 0-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCode(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 5v2m4-2v6m8-6v1M6 10v6m0 2.5v.5m4-.5v.5m4-.5v.5m4-.5v.5m-8-5v2m4-3v3m0-11v5m4-1v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCodeCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 5v2m4-2v6m8-6v1M6 10v6m0 2.5v.5m4-.5v.5m0-5v2m4-3v2m0-10v5m4-1v4m-3 6l2 2l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCodeXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 5v2m4-2v6m8-6v1M6 10v6m0 2.5v.5m4-.5v.5m0-5v2m4-3v2m0-10v5m4-1v4m-1.879 8.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L16.12 17.12m2.122 2.122l2.121 2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShortPants(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16.8h6.966a.6.6 0 0 0 .596-.53l1.36-11.6a.6.6 0 0 0-.596-.67H3.659a.6.6 0 0 0-.597.656l1.387 14.8a.6.6 0 0 0 .597.544H11.4a.6.6 0 0 0 .6-.6V12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShortPantsPockets(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.06 5.655A.6.6 0 0 1 3.658 5h16.684a.6.6 0 0 1 .598.655l-1.176 12.8a.6.6 0 0 1-.597.545h-4.152a.6.6 0 0 1-.574-.424l-1.867-6.102c-.174-.566-.974-.566-1.148 0l-1.868 6.102a.6.6 0 0 1-.573.424H4.833a.6.6 0 0 1-.597-.545L3.643 12z"/><path d="M4 9.5h1.5a2 2 0 0 0 2-2V5m13 4.5h-2a2 2 0 0 1-2-2V5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShortcutSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M15.025 8.025h-4.95m4.95 0v4.95m0-4.95l-3.535 3.536c-2.475 2.475 0 4.95 0 4.95"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 7c-3 0-8.5 0-10.5 5.5S5 18 2 18"/><path d="m20 5l2 2l-2 2m2 9c-3 0-8.5 0-10.5-5.5S5 7 2 7"/><path d="m20 20l2-2l-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SidebarCollapse(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2"/><path d="M7.25 10L5.5 12l1.75 2m2.25 7V3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SidebarExpand(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2m-9.5 0V3"/><path d="m5.5 10l1.75 2l-1.75 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SigmaFunction(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v3M4 20h16v-3M4 20l8-8l-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimpleCart(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h19l-3 10H6zm0 0l-.75-2.5m8.75 16a1.5 1.5 0 0 1-3 0m9 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SineWave(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-3.857 1.286-9 3.857-9c3.857 0 6.429 18 10.286 18C19.714 21 21 15.857 21 12M3 12h2m14 0h2m-5.5 0h1m-9 0h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SingleTapGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 20.5a7 7 0 1 0 0-14a7 7 0 0 0 0 14"/><path d="M4 7.29C5.496 5.039 8.517 3.5 12 3.5c3.483 0 6.504 1.539 8 3.79"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skateboard(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 16a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M2 9l3.333 1h13.334L22 9m-4.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skateboarding(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 19l2.333 1h9.334L19 19M8 22.01l.01-.011m7.99.011l.01-.011M7 7.834l3-1.5c2-1 4.27.567 4.27.567l-4.308 3.135L14 13.334v4m-4.451-3.989l-1.24.827H5M15.165 9.21h2.722M17 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipNext(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7v10M6.972 5.267A.6.6 0 0 0 6 5.738v12.524a.6.6 0 0 0 .972.47l7.931-6.261a.6.6 0 0 0 0-.942z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipNextSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 7v10"/><path fill="currentColor" d="M6.972 5.267A.6.6 0 0 0 6 5.738v12.524a.6.6 0 0 0 .972.47l7.931-6.261a.6.6 0 0 0 0-.942z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipPrev(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 7v10M17.028 5.267a.6.6 0 0 1 .972.471v12.524a.6.6 0 0 1-.972.47l-7.931-6.261a.6.6 0 0 1 0-.942z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipPrevSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 7v10"/><path fill="currentColor" d="M17.028 5.267a.6.6 0 0 1 .972.471v12.524a.6.6 0 0 1-.972.47l-7.931-6.261a.6.6 0 0 1 0-.942z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 4L8 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlashSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6M10 16l4-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SleeperChair(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 18v3m1-11V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v5"/><path d="M19.5 10a2.5 2.5 0 0 0-2.5 2.5V14H7v-1.5a2.5 2.5 0 1 0-3 2.45V18h16v-3.05a2.5 2.5 0 0 0-.5-4.95m.5 8v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slips(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M1 4.6a.6.6 0 0 1 .6-.6h20.8a.6.6 0 0 1 .6.6v3.912c0 .284-.199.53-.476.595c-1.052.247-3.635.914-5.524 1.893c-3.444 1.786-3.93 6.655-3.993 8.382a.637.637 0 0 1-.626.618h-.761a.637.637 0 0 1-.627-.618C10.931 17.655 10.443 12.786 7 11c-1.889-.98-4.472-1.646-5.524-1.893A.614.614 0 0 1 1 8.512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallLamp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.872 3.428l-2.64 8.8a.6.6 0 0 0 .574.772h14.388a.6.6 0 0 0 .574-.772l-2.64-8.8A.6.6 0 0 0 16.554 3H7.446a.6.6 0 0 0-.574.428M12 17v-2m-3.4 6h6.8c.331 0 .595-.268.542-.596C15.763 19.29 15.026 17 12 17c-3.026 0-3.763 2.29-3.942 3.404c-.053.328.21.596.542.596"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallLampAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.872 3.428l-2.64 8.8a.6.6 0 0 0 .574.772h14.388a.6.6 0 0 0 .574-.772l-2.64-8.8A.6.6 0 0 0 16.554 3H7.446a.6.6 0 0 0-.574.428M8 15v-2m0 8h8m-4-6v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneDevice(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m12 16.01l.01-.011"/><path d="M7 19.4V4.6a.6.6 0 0 1 .6-.6h8.8a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smoking(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18 19v3m0-6c0-1-1-2-3-2h-1a3 3 0 0 1-3-3V8.5A2.5 2.5 0 0 1 13.5 6v0h.5m8 10c0-4.5-2-5.5-4-6c2-.5 4-1 4-4s-2.5-4-4-4m4 17v3"/><rect width="12" height="3" x="2" y="19" rx=".6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snapchat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 8.75c0-9-12-9-12 0v1.5H4.372c-.583 0-.823.749-.348 1.088L6 12.75c-.333 1.167-1.7 3.7-4.5 4.5c.333.5 1.3 1.5 2.5 1.5l1 1.5l2.5-.5c.833.667 2.9 2 4.5 2s3.667-1.333 4.5-2l2.5.5l1-1.5c1.2 0 2.167-1 2.5-1.5c-2.8-.8-4.167-3.333-4.5-4.5l1.976-1.412c.475-.339.235-1.088-.348-1.088H18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12v5m0 5v-5m0 0l-4.5-2.5M12 17l4.5 2.5M12 17l4.5-2.5M12 17l-4.5 2.5M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowFlake(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 7l3.5 2M21 17l-3.5-2M12 12L6.5 9m5.5 3l-5.5 3m5.5-3V5m0 7v6.5m0-6.5l5.5 3M12 12l5.5-3M12 2v3m0 17v-3.5M21 7l-3.5 2M3 17l3.5-2m0-6L3 10m3.5-1L6 5.5m.5 9.5L3 14m3.5 1L6 18.5M12 5L9.5 4M12 5l2.5-1M12 18.5l2.5 1.5M12 18.5L9.5 20m8-5l.5 3.5m-.5-3.5l3.5-1m-3.5-5l3.5 1m-3.5-1l.5-3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soap(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 11a4 4 0 0 1 4-4h2a4 4 0 0 1 4 4v9.4a.6.6 0 0 1-.6.6H7.6a.6.6 0 0 1-.6-.6zm0 2h10m-5-6V3m0 0H9m3 0h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoccerBall(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 8l3.804 2.764M12 8l-3.804 2.764M12 8V5m3.804 5.764l-1.453 4.472m1.453-4.472L18.5 9.5m-4.149 5.736H9.65m4.702 0L16 17.5m-6.351-2.264l-1.453-4.472m1.453 4.472L8 17.5m.196-6.736L5.5 9.5m0 0L2.05 13M5.5 9.5l-1-4.115m14 4.115l3.45 3.5M18.5 9.5l1-4.115M12 5L8.624 2.584M12 5l3.376-2.416M8 17.5L3.338 17M8 17.5l2.5 4.388M16 17.5l4.662-.5M16 17.5l-2.5 4.388M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sofa(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 16v3M4 9V7a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2"/><path d="M20 9a2 2 0 0 0-2 2v2H6v-2a2 2 0 1 0-4 0v6h20v-6a2 2 0 0 0-2-2m2 7v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soil(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 4h20M3 8.01l.01-.011M3 16.01l.01-.011M6 12.01l.01-.011M6 20.01l.01-.011M9 8.01l.01-.011M9 16.01l.01-.011M12 12.01l.01-.011M12 20.01l.01-.011M15 8.01l.01-.011M15 16.01l.01-.011M18 12.01l.01-.011M18 20.01l.01-.011M21 8.01l.01-.011M21 16.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoilAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12h4m11 0h5M3 20.01l.01-.011M6 16.01l.01-.011M9 20.01l.01-.011M12 16.01l.01-.011M15 20.01l.01-.011M18 16.01l.01-.011M21 20.01l.01-.011M9 13s.9-3.741 3-6"/><path d="m16.186 2.241l.374 3.89c.243 2.523-1.649 4.77-4.172 5.012c-2.475.238-4.718-1.571-4.956-4.047a4.503 4.503 0 0 1 4.05-4.914l4.147-.4a.51.51 0 0 1 .557.46"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 14H2m6-4H2m4-4H2m10 12H2m17 2V4m0 16l3-3m-3 3l-3-3m3-13l3 3m-3-3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 10H2m8 4H2m4 4H2M18 6H2m17 4v10m0 0l3-3m-3 3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 14H2m8-4H2m4-4H2m16 12H2m17-4V4m0 0l3 3m-3-3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundHigh(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M1 13.857v-3.714a2 2 0 0 1 2-2h2.9a1 1 0 0 0 .55-.165l6-3.956a1 1 0 0 1 1.55.835v14.286a1 1 0 0 1-1.55.835l-6-3.956a1 1 0 0 0-.55-.165H3a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.5 7.5S19 9 19 11.5s-1.5 4-1.5 4m3-11S23 7 23 11.5s-2.5 7-2.5 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundLow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 7.5S21 9 21 11.5s-1.5 4-1.5 4"/><path d="M2 13.857v-3.714a2 2 0 0 1 2-2h2.9a1 1 0 0 0 .55-.165l6-3.956a1 1 0 0 1 1.55.835v14.286a1 1 0 0 1-1.55.835l-6-3.956a1 1 0 0 0-.55-.165H4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundMin(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.5 13.857v-3.714a2 2 0 0 1 2-2h2.9a1 1 0 0 0 .55-.165l6-3.956a1 1 0 0 1 1.55.835v14.286a1 1 0 0 1-1.55.835l-6-3.956a1 1 0 0 0-.55-.165H5.5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" d="M20.5 15V9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m18 14l2-2m2-2l-2 2m0 0l-2-2m2 2l2 2"/><path d="M2 13.857v-3.714a2 2 0 0 1 2-2h2.9a1 1 0 0 0 .55-.165l6-3.956a1 1 0 0 1 1.55.835v14.286a1 1 0 0 1-1.55.835l-6-3.956a1 1 0 0 0-.55-.165H4a2 2 0 0 1-2-2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spades(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 14.5c3 4.5 9 4.47 9-.5c0-4-4-7-9-12c-5 5-9 8-9 12c0 4.97 6 5 9 .5"/><path d="m11.47 15.493l-3 5.625A.6.6 0 0 0 9 22h6a.6.6 0 0 0 .53-.882l-3-5.625a.6.6 0 0 0-1.06 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M3 12c6.268 0 9-2.637 9-9c0 6.363 2.713 9 9 9c-6.287 0-9 2.713-9 9c0-6.287-2.732-9-9-9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SparkSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M3 12c6.268 0 9-2.637 9-9c0 6.363 2.713 9 9 9c-6.287 0-9 2.713-9 9c0-6.287-2.732-9-9-9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sparks(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M8 15c4.875 0 7-2.051 7-7c0 4.949 2.11 7 7 7c-4.89 0-7 2.11-7 7c0-4.89-2.125-7-7-7ZM2 6.5c3.134 0 4.5-1.318 4.5-4.5c0 3.182 1.357 4.5 4.5 4.5c-3.143 0-4.5 1.357-4.5 4.5c0-3.143-1.366-4.5-4.5-4.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SparksSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M8 15c4.875 0 7-2.051 7-7c0 4.949 2.11 7 7 7c-4.89 0-7 2.11-7 7c0-4.89-2.125-7-7-7ZM2 6.5c3.134 0 4.5-1.318 4.5-4.5c0 3.182 1.357 4.5 4.5 4.5c-3.143 0-4.5 1.357-4.5 4.5c0-3.143-1.366-4.5-4.5-4.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sphere(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M12 22c-3.314 0-6-4.477-6-10S8.686 2 12 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spiral(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.5 6.004C3.5 7.808 6.357 9 11.5 9c7 0 8-2.996 8-2.996S18.5 3 11.5 3c-5.143 0-8 1.2-8 3.004m0 6c0 1.804 2.857 2.996 8 2.996c7 0 8-2.996 8-2.996S18.5 9 11.5 9c-5.143 0-8 1.2-8 3.004m0 6c0 1.804 2.857 2.996 8 2.996c7 0 8-2.996 8-2.996S18.5 15 11.5 15c-5.143 0-8 1.2-8 3.004M19.5 12s1-.975 1-3s-1-3-1-3m1-2c0 1.35-1 2-1 2m0 12s1-.975 1-3s-1-3-1-3m1 8c0-1.35-1-2-1-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SplitArea(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 20H4v-4h16zM2 12h20M7 4H4v3m7-3h2m4 0h3v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SplitSquareDashed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 12h20M7 4H4v3m7-3h2m4 0h3v3m-9 13h2m-6 0H4v-3m13 3h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpockHandGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m18 7.5l.919.153a2 2 0 0 1 1.623 2.407l-.528 2.376a.602.602 0 0 0-.014.13V17.5s0 0 0 0c0 2-1.6 4-4 4H9.42a2 2 0 0 1-1.519-.698l-4.548-5.307a1.582 1.582 0 0 1-.034-2.018v0a1.582 1.582 0 0 1 2.426-.054L8 16v-3.5"/><path d="m9 5l-.79.132a2 2 0 0 0-1.595 2.522L8 12.5m3 0L8.923 4.606a1.514 1.514 0 0 1 1.215-1.879v0a1.514 1.514 0 0 1 1.713 1.108L14 12m3 .5l1-5l.247-1.485a1.536 1.536 0 0 0-1.262-1.768v0a1.536 1.536 0 0 0-1.762 1.233L14 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotify(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 15s4.5-1 9 1m-9.5-4s6-1.5 11 1.5M6 9c3-.5 8-1 13 2"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCursor(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M21 12V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7"/><path d="M20.879 16.917c.494.304.463 1.043-.045 1.101l-2.567.291l-1.151 2.312c-.228.46-.933.234-1.05-.334l-1.255-6.116c-.099-.48.333-.782.75-.525z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCursorSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M21 12V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7"/><path fill="currentColor" d="M20.879 16.918c.494.304.463 1.043-.045 1.1l-2.567.292l-1.151 2.312c-.228.459-.933.234-1.05-.334l-1.255-6.116c-.099-.48.333-.782.75-.525z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDashed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4H4v3m0 4v2m7-9h2m-2 16h2m7-9v2m-3-9h3v3M7 20H4v-3m13 3h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareThreeDcornerToCorner(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2m18 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareThreeDfromCenter(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareThreeDthreePoints(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 21V3.6a.6.6 0 0 1 .6-.6H21"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 21h3.4a.6.6 0 0 0 .6-.6V17m0-10v2m0 3v2M7 21h2m3 0h2"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2M21 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareWave(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h3V4h6v16h6v-8h3m-6.5 0h1m-7 0h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stackoverflow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 15v6H5v-6m11 2H8m7.913-2.337L8.087 13m8.626-.62L9.463 9m8.71 1.642L12.044 5.5m7.99 3.304L15.109 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.587 8.236l2.598-5.232a.911.911 0 0 1 1.63 0l2.598 5.232l5.808.844a.902.902 0 0 1 .503 1.542l-4.202 4.07l.992 5.75c.127.738-.653 1.3-1.32.952L12 18.678l-5.195 2.716c-.666.349-1.446-.214-1.319-.953l.992-5.75l-4.202-4.07a.902.902 0 0 1 .503-1.54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarDashed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13.806 5l-.99-1.996a.911.911 0 0 0-1.631 0l-.496.998m4.322 3.425l.402.809l1.452.211m2.905.423l1.451.21a.902.902 0 0 1 .503 1.542l-1.05 1.017m-2.102 2.035l-1.05 1.018l.248 1.437m.496 2.875l.248 1.437c.127.739-.653 1.302-1.32.953l-1.298-.679M10.428 19.5L12 18.678l1.299.679m-7.628.012l-.185 1.072c-.127.739.653 1.302 1.32.953l.847-.443M6.253 16l.225-1.308l-.695-.673M3.699 12l-1.423-1.378a.902.902 0 0 1 .503-1.542l1.11-.161M7 8.467l1.587-.231l.804-1.618"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfDashed(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.815 3.004a.911.911 0 0 0-1.63 0l-.496.998M12 18.678l-1.572.822m-4.757-.131l-.185 1.072c-.127.739.653 1.302 1.32.953l.847-.443M6.253 16l.225-1.308l-.695-.673M3.699 12l-1.423-1.378a.902.902 0 0 1 .503-1.542l1.11-.161M7 8.467l1.587-.231l.804-1.618"/><path d="m15.413 8.236l-2.598-5.232A.899.899 0 0 0 12 2.5v16.178l5.195 2.716c.666.349 1.446-.214 1.319-.953l-.992-5.75l4.202-4.07a.902.902 0 0 0-.503-1.54z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.587 8.236l2.598-5.232a.911.911 0 0 1 1.63 0l2.598 5.232l5.808.844a.902.902 0 0 1 .503 1.542l-4.202 4.07l.992 5.75c.127.738-.653 1.3-1.32.952L12 18.678l-5.195 2.716c-.666.349-1.446-.214-1.319-.953l.992-5.75l-4.202-4.07a.902.902 0 0 1 .503-1.54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 10l8 8l3-3l5 5M16 4v8m0 0l3-3m-3 3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 20v-8m0 0l3 3m-3-3l-3 3m-9-1l8-8l3 3l5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatsDownSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8 16V8m4 8v-5m4 5v-3"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatsDownSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zM8.75 8a.75.75 0 0 0-1.5 0v8a.75.75 0 0 0 1.5 0zM12 10.25a.75.75 0 0 1 .75.75v5a.75.75 0 0 1-1.5 0v-5a.75.75 0 0 1 .75-.75M16.75 13a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatsReport(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10 9H6m9.5 2a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5M6 6h3m9 12l-4.5-3l-2.5 2l-5-4"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatsUpSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16 16V8m-4 8v-5m-4 5v-3"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatsUpSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zM16.75 8a.75.75 0 0 0-1.5 0v8a.75.75 0 0 0 1.5 0zM12 10.25a.75.75 0 0 1 .75.75v5a.75.75 0 0 1-1.5 0v-5a.75.75 0 0 1 .75-.75M8.75 13a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strategy(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 20.5C7 11 11.5 8 20 6"/><path d="m15.909 3.81l4.486 2.09l-2.092 4.486M5 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4m11 13.243l2.121-2.122m0 0L20.243 16m-2.122 2.121L16 16m2.121 2.121l2.122 2.122"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stretching(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4M5 20l4.91-.524l2.726-5.238L13.727 9l-4.909 1.048l1.636 2.095m4.364 3.143H17V20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h18m-4.714-9h-6.218a4.068 4.068 0 0 0-1.286 7.927L12 12m-6 9h7.932a4.068 4.068 0 0 0 3.58-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stroller(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.5 3a8.5 8.5 0 0 0-7.212 13m14.425 0A8.46 8.46 0 0 0 20 11.5v-2h2.5M8 21a2 2 0 1 1 0-4a2 2 0 0 1 0 4m7 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4M11.5 3v9m-8 0h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StyleBorder(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.499"><path stroke-dasharray="2 2" d="M16 2H8a6 6 0 0 0-6 6v8a6 6 0 0 0 6 6h8a6 6 0 0 0 6-6V8a6 6 0 0 0-6-6"/><path d="M16 5H8a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StyleBorderSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.499"><path stroke-dasharray="2 2" d="M16 2H8a6 6 0 0 0-6 6v8a6 6 0 0 0 6 6h8a6 6 0 0 0 6-6V8a6 6 0 0 0-6-6"/><path fill="currentColor" d="M16 5H8a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubmitDocument(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 13V5.749a.6.6 0 0 0-.176-.425l-3.148-3.148A.6.6 0 0 0 16.252 2H4.6a.6.6 0 0 0-.6.6v18.8a.6.6 0 0 0 .6.6H14"/><path d="M16 2v3.4a.6.6 0 0 0 .6.6H20m-4 13h6m0 0l-3-3m3 3l-3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Substract(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 3.6v10.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6M13.5 21h3m4.5-7.5v3m0 3v.9a.6.6 0 0 1-.6.6h-.9m-9 0h-.9a.6.6 0 0 1-.6-.6v-.9M19.5 9h.9a.6.6 0 0 1 .6.6v.9"/><path d="M16.5 9H9.6a.6.6 0 0 0-.6.6v6.9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suggestion(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v13.8a.6.6 0 0 1-.6.6h-4.14a.6.6 0 0 0-.438.189l-3.385 3.597a.6.6 0 0 1-.874 0l-3.385-3.597A.6.6 0 0 0 7.74 18H3.6a.6.6 0 0 1-.6-.6z"/><path stroke-linecap="round" stroke-linejoin="round" d="m12 7l1.425 2.575L16 11l-2.575 1.425L12 15l-1.425-2.575L8 11l2.575-1.425z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M8 7H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-4M8 7V3.6a.6.6 0 0 1 .6-.6h6.8a.6.6 0 0 1 .6.6V7M8 7h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunLight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12m10-6h1M12 2V1m0 22v-1m8-2l-1-1m1-15l-1 1M4 20l1-1M4 4l1 1m-4 7h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SvgFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.5 15h2A1.5 1.5 0 0 0 8 13.5v0A1.5 1.5 0 0 0 6.5 12H6a1.5 1.5 0 0 1-1.5-1.5v0A1.5 1.5 0 0 1 6 9h1.5m3 0l1.5 6l1.5-6m6 0h-3v6h3v-2.4"/><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SweepThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 22l5.5-5.5M20 5l-2.5 2.5M7 21h8.5l-7-12l-5.833 10M14.5 3l7 12m-13-6l6-6m1 18l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swimming(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 15c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3M3 20c2.483 0 4.345-3 4.345-3s1.862 3 4.345 3c2.482 0 4.965-3 4.965-3s2.483 3 4.345 3M5 10.5L9 8L7.813 6.516a1.262 1.262 0 0 1 .228-1.797v0a1.261 1.261 0 0 1 1.726.202L14 10m2.5-2a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwipeDownGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14a6 6 0 1 0 0-12a6 6 0 0 0 0 12m0 0v8m0 0l-3-3m3 3l3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwipeLeftGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12a6 6 0 1 0 12 0a6 6 0 0 0-12 0m0 0H2m0 0l3-3m-3 3l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwipeRightGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 12a6 6 0 1 1-12 0a6 6 0 0 1 12 0m0 0h8m0 0l-3-3m3 3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwipeTwoFingersDownGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 12a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7m0 0v7m0 0L9 16.6M6.5 19L4 16.6M17.5 12a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7m0 0v7m0 0l2.5-2.4M17.5 19L15 16.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwipeTwoFingersLeftGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0m0 0H5m0 0L7.4 15M5 17.5L7.4 20M12 6.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0m0 0H5m0 0L7.4 4M5 6.5L7.4 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwipeTwoFingersRightGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0m0 0h7m0 0L16.6 15m2.4 2.5L16.6 20M12 6.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0m0 0h7m0 0L16.6 4M19 6.5L16.6 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwipeTwoFingersUpGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 12a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7m0 0V5m0 0L9 7.4M6.5 5L4 7.4M17.5 12a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7m0 0V5m0 0L20 7.4M17.5 5L15 7.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwipeUpGesture(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 10a6 6 0 1 0 0 12a6 6 0 0 0 0-12m0 0V2m0 0l3 3m-3-3L9 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M17 17H7A5 5 0 0 1 7 7h10a5 5 0 0 1 0 10Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchOn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M17 17H7A5 5 0 0 1 7 7h10a5 5 0 0 1 0 10Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SystemRestart(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2v4m0 12v4m10-10h-4M6 12H2m2.929-7.071l2.828 2.828m8.486 8.486l2.828 2.828m0-14.142l-2.828 2.828m-8.486 8.486L4.93 19.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SystemShut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7v10m0 5c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3v18H3V3zM3 16.5h18M3 12h18M3 7.5h18M16.5 3v18M12 3v18M7.5 3v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableRows(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 12h18M3 12v4.5M3 12V7.5M21 12v4.5m0-4.5V7.5m-18 9v3.9a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6v-3.9m-18 0h18m0-9V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v3.9m18 0H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTwoColumns(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm0-3.9h18M3 12h18m0-4.5H3M12 21V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TaskList(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6h11M3.8 5.8l.8.8l2-2m-2.8 7.2l.8.8l2-2m-2.8 7.2l.8.8l2-2M9 12h11M9 18h11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telegram(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 5L2 12.5l7 1M21 5l-2.5 15L9 13.5M21 5L9 13.5m0 0V19l3.249-3.277"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelegramCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 8L5 12.5L9.5 14M18 8l-8.5 6M18 8l-4 10.5L9.5 14"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperatureDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 12a5 5 0 1 0 6 0m-6 0V3h6v9m0-9h2m-2 3h2m-2 3h2"/><path d="M8 14a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 0V6m11 0v12m0 0l2.5-2.5M19 18l-2.5-2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperatureHigh(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 12a5 5 0 1 0 6 0m-6 0V3h6v9m0-9h2m-2 3h2m-2 3h2m5-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M9 14a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 0V6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperatureLow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 12a5 5 0 1 0 6 0m-6 0V3h6v9m0-9h2m-2 3h2m-2 3h2m5-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M9 14a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 0v-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperatureUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 12a5 5 0 1 0 6 0m-6 0V3h6v9m0-9h2m-2 3h2m-2 3h2"/><path d="M7 14a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 0V6m12 12V6m0 0l2.5 2.5M19 6l-2.5 2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TennisBall(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M18.572 4.462c-2.667 4.53-2.667 9.723 0 15.076M5.428 4.462c2.667 4.53 2.667 9.723 0 15.076"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TennisBallAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.66 7c2.762 4.783 1.123 10.899-3.66 13.66C12.217 23.422 6.101 21.783 3.34 17C.578 12.217 2.217 6.1 7 3.34C11.783.578 17.899 2.217 20.66 7"/><path d="M21.46 15.242c-4.986-3.302-7.582-7.8-7.538-13.056m-3.844 19.628C9.71 15.844 7.114 11.347 2.54 8.758"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 17h7M5 7l5 5l-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 16h5M6 8l4 4l-4 4"/><path d="M2 18V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TestTube(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.141 19.995c2.458 1.72 4.281-.012 5.318-1.492l7.3-10.426l1.967-1.065l-6.554-4.588l-8.447 12.064c-1.037 1.48-2.041 3.786.416 5.507"/><path d="M16.091 11.02c-2.876-.853-4.403.781-7.28-.071"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextArrowsUpDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 21V11m0 10l-2-2.5m2 2.5l2-2.5M18 11l-2 2m2-2l2 2M9 5v12m0 0H7m2 0h2m4-10V5H3v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextBox(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M12 8v8m0-8H8m4 0h4"/><path d="M21 13.5V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-5.5m18-3V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v5.5m16.5 3v-3h3v3zm-18 0v-3h3v3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextIcon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7V5H5v2m7-2v14m0 0h-2m2 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextMagnifyingGlass(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 19.5L21 21m-7-4a3 3 0 1 0 6 0a3 3 0 0 0-6 0M9 5v12m0 0H7m2 0h2m4-10V5H3v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSize(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 7V5h14v2m-7-2v14m0 0h2m-2 0H8m5-5v-2h8v2m-4-2v7m0 0h-1.5m1.5 0h1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M7 9V7h10v2m-5-2v10m0 0h-2m2 0h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zm4 3.4A.75.75 0 0 1 7 6.25h10a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0V7.75h-3.5v8.5H14a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1 0-1.5h1.25v-8.5h-3.5V9a.75.75 0 0 1-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Threads(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.914 8.128c2.505-2.014 6.11-.94 6.536 2.372c.452 3.514-.45 6.3-3.95 6.3c-3.25 0-3.15-2.8-3.15-2.8c0-3 5.15-3.4 8.15-1.9C23 15.6 19 22 13 22c-4.97 0-9-2.5-9-10S8.03 2 13 2c3.508 0 6.672 1.807 7.835 5.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreePointsCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path fill="currentColor" d="M5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M5 10.5V9m0 6v-1.5"/><path fill="currentColor" d="M5 20a1 1 0 1 0 0-2a1 1 0 0 0 0 2m14 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M10.5 19H9m6 0h-1.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeStars(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.635 14.415l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L6 18.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.217-.22-.098-.604.2-.65zm12 0l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L18 18.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.216-.22-.098-.604.2-.65zm-6-9l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L12 9.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.217-.22-.098-.604.2-.65z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeStarsSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.635 14.415l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L6 18.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.217-.22-.098-.604.2-.65zm12 0l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L18 18.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.216-.22-.098-.604.2-.65zm-6-9l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L12 9.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.217-.22-.098-.604.2-.65z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M16.472 3.5H4.1a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6h2.768a2 2 0 0 1 1.715.971l2.71 4.517a1.631 1.631 0 0 0 2.961-1.308l-1.022-3.408a.6.6 0 0 1 .574-.772h4.576a2 2 0 0 0 1.929-2.526l-1.91-7A2 2 0 0 0 16.473 3.5Z"/><path stroke-linejoin="round" d="M7 14.5v-11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M16.472 20H4.1a.6.6 0 0 1-.6-.6V9.6a.6.6 0 0 1 .6-.6h2.768a2 2 0 0 0 1.715-.971l2.71-4.517a1.631 1.631 0 0 1 2.961 1.308l-1.022 3.408a.6.6 0 0 0 .574.772h4.576a2 2 0 0 1 1.929 2.526l-1.91 7A2 2 0 0 1 16.473 20Z"/><path stroke-linejoin="round" d="M7 20V9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thunderstorm(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.5 12L9 17h6l-2.5 5"/><path d="M20 17.607c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TifFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6"/><path stroke-linejoin="round" d="M15 15V9h3M6.5 9H8m1.5 0H8m0 0v6m7-3h2.5M12 15V9"/><path d="M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TiffFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M12 15V9h3m2.5 6V9h3m-17 0H5m1.5 0H5m0 0v6m7-3h2.5m3 0H20M9 15V9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tiktok(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 8v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5"/><path d="M10 12a3 3 0 1 0 3 3V6c.333 1 1.6 3 4 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeZone(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10"/><path d="M13 2.05S16 6 16 12m-5 9.95S8 18 8 12c0-6 3-9.95 3-9.95M2.63 15.5H12m-9.37-7h18.74m-2.37 9V19h1.5"/><path d="M19 23a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 2h6m-3 8v4m0 8a8 8 0 1 0 0-16a8 8 0 0 0 0 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimerOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 2h6M5 7l14 14.5M12 10v4M6.19 8.5a8 8 0 0 0 11.05 11.544M19.42 17A8 8 0 0 0 9.21 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimerSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 5.25a8.75 8.75 0 1 0 0 17.5a8.75 8.75 0 0 0 0-17.5m.75 4.75a.75.75 0 0 0-1.5 0v4a.75.75 0 0 0 1.5 0zm-4.5-8A.75.75 0 0 1 9 1.25h6a.75.75 0 0 1 0 1.5H9A.75.75 0 0 1 8.25 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tools(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.05 10.607l-7.07 7.07a2 2 0 0 0 0 2.83v0a2 2 0 0 0 2.828 0l7.07-7.072m4.315.365l3.878 3.878a2 2 0 0 1 0 2.828v0a2 2 0 0 1-2.828 0l-6.209-6.208M6.733 5.904L4.61 6.61L2.49 3.075l1.414-1.414L7.44 3.782zm0 0l2.83 2.83"/><path d="M10.05 10.607c-.844-2.153-.679-4.978 1.061-6.718c1.74-1.74 4.95-2.121 6.717-1.06l-3.04 3.04l-.283 3.111l3.111-.282l3.04-3.041c1.062 1.768.68 4.978-1.06 6.717c-1.74 1.74-4.564 1.905-6.717 1.061"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tournament(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3h5v6H3m5-3h7v12H8m7-6h7M3 15h5v6H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tower(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M17 22H7a2 2 0 0 1-2-2v-8.818a.6.6 0 0 0-.1-.333L3.1 8.15a.6.6 0 0 1-.1-.333V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v5.218a.6.6 0 0 1-.1.333l-1.8 2.698a.6.6 0 0 0-.1.333V20a2 2 0 0 1-2 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TowerCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13 19l3 3l5-5M9 22H7a2 2 0 0 1-2-2v-8.818a.6.6 0 0 0-.1-.333L3.1 8.15a.6.6 0 0 1-.1-.333V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v5.218a.6.6 0 0 1-.1.333l-1.8 2.698a.6.6 0 0 0-.1.333V13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TowerNoAccess(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.857 15.2a4 4 0 0 0-5.713 5.6m5.713-5.6a4 4 0 0 1-5.713 5.6m5.713-5.6l-5.714 5.6"/><path d="M9 22H7a2 2 0 0 1-2-2v-8.818a.6.6 0 0 0-.1-.333L3.1 8.15a.6.6 0 0 1-.1-.333V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v5.218a.6.6 0 0 1-.1.333L20 9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TowerWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 10v3m0 4.01l.01-.011"/><path d="M17 22H7a2 2 0 0 1-2-2v-8.818a.6.6 0 0 0-.1-.333L3.1 8.15a.6.6 0 0 1-.1-.333V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v5.218a.6.6 0 0 1-.1.333l-1.8 2.698a.6.6 0 0 0-.1.333V20a2 2 0 0 1-2 2Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trademark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 15.5v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H9.5v4.2m2.857 0H9.5m2.857 0l2.143 2.8"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M9.609 7h4.782A2.609 2.609 0 0 1 17 9.609a.391.391 0 0 1-.391.391H7.39A.391.391 0 0 1 7 9.609A2.609 2.609 0 0 1 9.609 7"/><path stroke-linejoin="round" d="M9 3h6a6 6 0 0 1 6 6v4a6 6 0 0 1-6 6H9a6 6 0 0 1-6-6V9a6 6 0 0 1 6-6m7 12.01l.01-.011M8 15.01l.01-.011"/><path d="m10.5 19l-2 2.5m5-2.5l2 2.5m1-2.5l2 2.5M7.5 19l-2 2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tram(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m15 16.01l.01-.011M9 16.01l.01-.011M13 6h2a5 5 0 0 1 5 5v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-7a5 5 0 0 1 5-5zm0 0l1-4m0 0h3m-3 0H7"/><path d="m10.5 20l-2 2.5m5-2.5l2 2.5m1-2.5l2 2.5M7.5 20l-2 2.5"/><path stroke-linejoin="round" d="M9.609 9h4.782A2.609 2.609 0 0 1 17 11.609a.391.391 0 0 1-.391.391H7.39a.391.391 0 0 1-.39-.391A2.609 2.609 0 0 1 9.609 9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransitionDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M18 2H6a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 16v2a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-2m-9-6v8m0 0l-3-3m3 3l3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransitionLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M22 18V6a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 3H6a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4h2m6-9H6m0 0l3-3m-3 3l3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransitionRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 18V6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 3h2a4 4 0 0 1 4 4v10a4 4 0 0 1-4 4h-2m-6-9h8m0 0l-3-3m3 3l-3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransitionUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M18 22H6a3 3 0 0 1-3-3v-2a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M3 8V6a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v2m-9 6V6m0 0L9 9m3-3l3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Translate(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 5h7m7 0h-2.5M9 5h4.5M9 5V3m4.5 2c-.82 2.735-2.539 5.32-4.5 7.593M4 17.5c1.585-1.359 3.376-3.026 5-4.907m0 0C8 11.5 6.4 9.3 6 8.5m3 4.093l3 2.907m1.5 5.5l1.143-3m6.857 3l-1.143-3m-5.714 0l2.857-7.5l2.857 7.5m-5.714 0h5.714"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 9l-1.995 11.346A2 2 0 0 1 16.035 22h-8.07a2 2 0 0 1-1.97-1.654L4 9m17-3h-5.625M3 6h5.625m0 0V4a2 2 0 0 1 2-2h2.75a2 2 0 0 1 2 2v2m-6.75 0h6.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path fill="currentColor" d="m20 9l-1.995 11.346A2 2 0 0 1 16.035 22h-8.07a2 2 0 0 1-1.97-1.654L4 9"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 9l-1.995 11.346A2 2 0 0 1 16.035 22h-8.07a2 2 0 0 1-1.97-1.654L4 9zm1-3h-5.625M3 6h5.625m0 0V4a2 2 0 0 1 2-2h2.75a2 2 0 0 1 2 2v2m-6.75 0h6.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Treadmill(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-2.387 1.267l-3.308 4.135l4.135 4.135l-2.067 4.55"/><path d="m4.41 8.508l3.387-3.309l2.816 2.068l2.895 3.308h1.722M6.892 14.71l-1.241.827H2.343m1 6l15.308-2V8"/><path d="M20.892 6L18.65 8L17 9.5m3.891 12.21l-2.24-2.173"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tree(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22v-8m0-4v4m0 0l4-2m1-5A5 5 0 0 0 7 7m5 11H7.5a5.5 5.5 0 1 1 0-11H9m3 11h4.5A5.5 5.5 0 0 0 17 7.022"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trekking(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18 10l-3 1.5l-4-3l-1 5.5l3.5 3l.5 4.5m4-13v13M10 17l-2 4.5m.5-13C7 9.5 6 12 6 12l2 1m4-6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trello(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2"/><path d="M10.4 6H6.6a.6.6 0 0 0-.6.6v10.8a.6.6 0 0 0 .6.6h3.8a.6.6 0 0 0 .6-.6V6.6a.6.6 0 0 0-.6-.6m7 0h-3.8a.6.6 0 0 0-.6.6v6.8a.6.6 0 0 0 .6.6h3.8a.6.6 0 0 0 .6-.6V6.6a.6.6 0 0 0-.6-.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.475 2.947a.6.6 0 0 1 1.05 0l9.373 16.912a.6.6 0 0 1-.524.891H2.626a.6.6 0 0 1-.525-.89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleFlag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 21v-5m0 0V3.577a.6.6 0 0 1 .916-.51l8.79 5.442a.6.6 0 0 1 .017 1.009z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleFlagCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 21.5v-6m0 0V6.997a.6.6 0 0 1 .88-.53l6.67 3.53a.6.6 0 0 1 .024 1.048zM22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleFlagTwoStripes(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 21v-5m0 0l9.723-6.482a.6.6 0 0 0-.017-1.01l-8.79-5.441a.6.6 0 0 0-.916.51zm0-5l6.5-4.476"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.745 4h10.568s-.88 13.257-5.284 13.257c-2.15 0-3.461-3.164-4.239-6.4C6.976 7.468 6.745 4 6.745 4"/><path d="M17.313 4s.921-.983 1.687-1c1.5-.034 1.777 1 1.777 1c.294.61.529 2.194-.88 3.657c-1.41 1.463-2.987 2.743-3.629 3.2M6.745 4S5.785 3.006 5 3c-1.5-.012-1.777 1-1.777 1c-.294.61-.529 2.194.88 3.657a29.896 29.896 0 0 0 3.687 3.2M8.507 20c0-1.829 3.522-2.743 3.522-2.743s3.523.914 3.523 2.743z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M7 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4m10 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M14 17V6.6a.6.6 0 0 0-.6-.6H2.6a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6h2.05M14 17H9.05M14 9h5.61a.6.6 0 0 1 .548.356l1.79 4.028a.6.6 0 0 1 .052.243V16.4a.6.6 0 0 1-.6.6h-1.9M14 17h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TruckGreen(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M7 17a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M14 15V4.6a.6.6 0 0 0-.6-.6H2.6a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6h2.05M14 15H9.05M14 7h5.61a.6.6 0 0 1 .548.356l1.29 2.903a.6.6 0 0 1 .052.243V12"/><path stroke-linejoin="round" d="M17 23s.9-3.118 3-5"/><path stroke-linejoin="round" d="m19.802 21.425l-.134.012a3.094 3.094 0 0 1-3.366-2.774a3.06 3.06 0 0 1 2.761-3.35l2.986-.28a.35.35 0 0 1 .381.314l.255 2.58a3.194 3.194 0 0 1-2.883 3.497"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TruckLength(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" stroke-miterlimit="1.5" d="M7 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4m10 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M14 14V3.6a.6.6 0 0 0-.6-.6H2.6a.6.6 0 0 0-.6.6v9.8a.6.6 0 0 0 .6.6h2.05M14 14H9.05M14 6h5.61a.6.6 0 0 1 .548.356l1.79 4.028a.6.6 0 0 1 .052.243V13.4a.6.6 0 0 1-.6.6h-1.9M14 14h1"/><path stroke-linejoin="round" d="M3 20h17.75M3 20l1.75 1.75M3 20l1.75-1.75m16 1.75L19 21.75M20.75 20L19 18.25"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tunnel(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21 20L3 14"/><path d="M16 10v1m-4-2v1M8 8v1"/><path stroke-linejoin="round" d="M3 21h18v-9a9 9 0 1 0-18 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 20V9a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.5 2.5L12 6l3.5-3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvFix(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 20V9a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m13.657 12.828l-2.829 2.829m5.657-2.829A2 2 0 1 1 13.657 10m-2.829 8.485A2 2 0 0 0 8 15.657M8.5 2.5L12 6l3.5-3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 11v3m0 4.01l.01-.011"/><path d="M2 20V9a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.5 2.5L12 6l3.5-3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M23 3.01s-2.018 1.192-3.14 1.53a4.48 4.48 0 0 0-7.86 3v1a10.66 10.66 0 0 1-9-4.53s-4 9 5 13a11.64 11.64 0 0 1-7 2c9 5 20 0 20-11.5c0-.278-.028-.556-.08-.83C21.94 5.674 23 3.01 23 3.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoPointsCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path fill="currentColor" d="M5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2m14 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="m19 19l-1.5-1.5m-2-2l-1-1m-2-2l-1-1m-2-2l-1-1m-2-2L5 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoSeaterSofa(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 16v3m10-6V7a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2m-8 4V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v2"/><path d="M20 9a2 2 0 0 0-2 2v2H6v-2a2 2 0 1 0-4 0v6h20v-6a2 2 0 0 0-2-2m2 7v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Type(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 16.249a.6.6 0 0 0-.176-.425l-3.648-3.648A.6.6 0 0 1 3 11.75V4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v7.752a.6.6 0 0 1-.176.424l-3.648 3.648a.6.6 0 0 0-.176.425V20a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9.5 11.5l.5-1.1m4.5 1.1l-.5-1.1m0 0L12 6l-2 4.4m4 0h-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UturnArrowLeft(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 14V8a5 5 0 0 1 10 0v13"/><path d="m12 11l-4 4l-4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UturnArrowRight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 14V8A5 5 0 0 0 6 8v13"/><path d="m12 11l4 4l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19.778 4.043C17.701 2.081 14.938 1 12 1C9.062 1 6.3 2.08 4.222 4.043C2.144 6.006 1 8.616 1 11.391c0 .336.289.609.644.609c.356 0 .645-.273.645-.609c0-1.013.872-1.837 1.944-1.837C6.126 9.554 5.431 12 6.823 12c1.39 0 .696-2.446 2.588-2.446C11.304 9.554 12 12 12 12s.697-2.446 2.589-2.446S15.988 12 17.178 12s.696-2.446 2.589-2.446c1.072 0 1.944.824 1.944 1.837c0 .336.289.609.645.609c.355 0 .644-.273.644-.609c0-2.775-1.144-5.385-3.222-7.348Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12v8c0 4-6 4-6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 5v6a4 4 0 0 1-4 4v0a4 4 0 0 1-4-4V5M6 19h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnderlineSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 6v4a4 4 0 0 1-4 4v0a4 4 0 0 1-4-4V6M6 18h12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnderlineSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zM8 5.25a.75.75 0 0 1 .75.75v4a3.25 3.25 0 0 0 6.5 0V6a.75.75 0 0 1 1.5 0v4a4.75 4.75 0 1 1-9.5 0V6A.75.75 0 0 1 8 5.25m-2 12a.75.75 0 0 0 0 1.5h12a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 8H15s0 0 0 0s5 0 5 4.706C20 18 15 18 15 18H6.286"/><path d="M7.5 11.5L4 8l3.5-3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UndoAction(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 5v6m3.5-3H15s0 0 0 0s5 0 5 4.706C20 18 15 18 15 18H6.286"/><path d="M11.5 11.5L8 8l3.5-3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UndoCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 10.625h7.2s0 0 0 0s2.8 0 2.8 3C17 17 14.2 17 14.2 17h-.8"/><path d="M10.5 14L7 10.625L10.5 7"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UndoCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12m7.608-.625l2.163 2.085a.75.75 0 1 1-1.042 1.08L6.48 11.166a.748.748 0 0 1-.019-1.063l3.5-3.624a.75.75 0 0 1 1.079 1.042L8.767 9.875H14.2v.75v-.75h.022a2.02 2.02 0 0 1 .156.008a4.027 4.027 0 0 1 1.577.456c.443.237.9.602 1.245 1.155c.346.556.551 1.26.551 2.131c0 1.94-.829 3.033-1.763 3.596a3.788 3.788 0 0 1-1.764.529H14.2V17v.75h-.8a.75.75 0 0 1 0-1.5h.801l.055-.004a2.288 2.288 0 0 0 .957-.31c.466-.281 1.037-.876 1.037-2.311c0-.629-.145-1.05-.324-1.338a1.763 1.763 0 0 0-.68-.626a2.527 2.527 0 0 0-1.041-.286h-.01z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Union(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 9h5.4a.6.6 0 0 1 .6.6v10.8a.6.6 0 0 1-.6.6H9.6a.6.6 0 0 1-.6-.6V15"/><path d="M15 9V3.6a.6.6 0 0 0-.6-.6H3.6a.6.6 0 0 0-.6.6v10.8a.6.6 0 0 0 .6.6H9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnionAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 22A7 7 0 1 0 9 8a7 7 0 0 0 0 14"/><path d="M15 16a7 7 0 1 0 0-14a7 7 0 0 0 0 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnionHorizAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 19A7 7 0 1 0 8 5a7 7 0 0 0 0 14"/><path d="M16 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unity(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 13h9.5M4 13l4 4.5M4 13l4-4.5m5.5 4.5l5-9m-5 9l5 7m0-16l-6 1m6-1L20 9.5M18.5 20l1.5-5.5M18.5 20l-6-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnityFive(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.672 20.786a.6.6 0 0 0 .656 0l9.284-6.062a.6.6 0 0 0 .24-.694L18.285 3.41a.6.6 0 0 0-.569-.41H6.221a.6.6 0 0 0-.57.412l-3.506 10.62a.6.6 0 0 0 .241.69z"/><path d="M14.5 6H10l-.5 5a3 3 0 1 1 0 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnjoinThreeD(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.5 7L9 12h6l-2.5 5m8.339 3.84h-3.536m3.536 0v-3.536m0 3.535L17 17M2.768 2.768h3.535m-3.535 0v3.535m0-3.535l3.839 3.839"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 20h12m-6-4V4m0 0l3.5 3.5M12 4L8.5 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadDataWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v9"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M19.5 22v-6m0 0L17 18.5m2.5-2.5l2.5 2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 18h12m-6-4V6m0 0l3.5 3.5M12 6L8.5 9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zm3 14.4a.75.75 0 0 1 .75-.75h12a.75.75 0 0 1 0 1.5H6a.75.75 0 0 1-.75-.75m7.28-12.53a.75.75 0 0 0-1.06 0l-3.5 3.5a.75.75 0 1 0 1.06 1.06l2.22-2.22V14a.75.75 0 0 0 1.5 0V7.81l2.22 2.22a.75.75 0 1 0 1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.5 2v15m0-3l5.5-2V8.5M12.5 16L7 14.5v-3M12.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m4-16.5v3h3v-3zm-6-1.5l2-2l2 2M7 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsbSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.5 2v15m0-3l5.5-2V8.5M12.5 16L7 14.5v-3"/><path fill="currentColor" d="M12.5 22a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m4-16.5v3h3v-3z"/><path d="m10.5 4l2-2l2 2"/><path fill="currentColor" d="M7 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 20v-1a7 7 0 0 1 7-7v0a7 7 0 0 1 7 7v1m-7-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserBadgeCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2 20v-1a7 7 0 0 1 7-7v0"/><path d="M15.804 12.314a1.618 1.618 0 0 1 2.392 0c.325.356.79.549 1.272.526a1.618 1.618 0 0 1 1.692 1.692c-.023.481.17.947.526 1.272c.705.642.705 1.75 0 2.392c-.356.325-.549.79-.526 1.272a1.618 1.618 0 0 1-1.692 1.692a1.618 1.618 0 0 0-1.272.526a1.618 1.618 0 0 1-2.392 0a1.618 1.618 0 0 0-1.272-.526a1.618 1.618 0 0 1-1.692-1.692a1.618 1.618 0 0 0-.527-1.272a1.618 1.618 0 0 1 0-2.392c.357-.325.55-.79.527-1.272a1.618 1.618 0 0 1 1.692-1.692c.481.023.947-.17 1.272-.527z"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.364 17l1.09 1.09l2.182-2.18M9 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserBag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path d="M2 18a7 7 0 0 1 11.33-5.5m8.034 4.207l.296 2A2 2 0 0 1 19.682 21h-3.364a2 2 0 0 1-1.978-2.293l.296-2A2 2 0 0 1 16.614 15h2.772a2 2 0 0 1 1.978 1.707M17 13h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCart(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m22 12.5l-.833 2.5m0 0L20 18.5h-4.5l-1-3.5zM16.5 20.51l.01-.011m2.99.011l.01-.011M9 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path d="M2 18a7 7 0 0 1 11.33-5.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M7 18v-1a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><circle cx="12" cy="12" r="10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCrown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path d="M5 20v-1a7 7 0 0 1 10-6.326M21 22l1-6l-3.5 1.8L17 16l-1.5 1.8L12 16l1 6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserLove(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M12 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path d="M22 17.28a2.28 2.28 0 0 1-.662 1.606c-.976.984-1.923 2.01-2.936 2.958a.597.597 0 0 1-.823-.017l-2.918-2.94a2.281 2.281 0 0 1 0-3.214a2.277 2.277 0 0 1 3.233 0l.106.107l.106-.107A2.277 2.277 0 0 1 22 17.28Z"/><path stroke-linecap="round" d="M5 20v-1a7 7 0 0 1 10-6.326"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 10h3m3 0h-3m0 0V7m0 3v3M1 20v-1a7 7 0 0 1 7-7v0a7 7 0 0 1 7 7v1m-7-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserScan(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 3H3v3m15-3h3v3M6 21H3v-3m4 0v-1a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v1m-5-6a3 3 0 1 0 0-6a3 3 0 0 0 0 6m6 9h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M7 18v-1a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v1"/><path stroke-linejoin="round" d="M12 12a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path stroke-linejoin="round" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserStar(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path d="M5 20v-1a7 7 0 0 1 10-6.326m1.635 3.741l1.039-2.203a.357.357 0 0 1 .652 0l1.04 2.203l2.323.356c.298.045.416.429.2.649l-1.68 1.713l.396 2.421c.051.311-.26.548-.527.401L18 20.812l-2.078 1.143c-.267.147-.578-.09-.527-.4l.396-2.422l-1.68-1.713c-.216-.22-.098-.604.2-.65z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.621 12.121L20.743 10m2.121-2.121L20.743 10m0 0L18.62 7.879M20.743 10l2.121 2.121M1 20v-1a7 7 0 0 1 7-7v0a7 7 0 0 1 7 7v1m-7-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vegan(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 11.063C12.53 13.65 10.059 20 10.059 20S6.529 11.063 3 9"/><path d="m20.496 5.577l.426 4.424c.276 2.87-1.875 5.425-4.745 5.702c-2.816.27-5.367-1.788-5.638-4.604a5.122 5.122 0 0 1 4.608-5.59l4.716-.454a.58.58 0 0 1 .633.522"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VeganCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 11.5C12.75 13.382 11 18 11 18s-2.5-6.5-5-8"/><path d="m18.016 7.73l.296 3.08c.192 1.998-1.306 3.777-3.304 3.97c-1.96.188-3.736-1.245-3.925-3.205a3.566 3.566 0 0 1 3.208-3.892l3.284-.316a.404.404 0 0 1 .44.363"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VeganSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M14.5 10.5C12.75 12.382 11 17 11 17s-2.5-6.5-5-8"/><path d="m18.016 6.73l.296 3.08c.192 1.998-1.306 3.777-3.304 3.97c-1.96.188-3.736-1.245-3.925-3.205a3.566 3.566 0 0 1 3.208-3.892l3.284-.316a.404.404 0 0 1 .44.363"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VehicleGreen(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M7 10h8m-9 4h1m8 0h1"/><path d="M6 18H2v2.4a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6zm0 0h7M2 18v-6.59a2 2 0 0 1 .162-.787l2.319-5.41A2 2 0 0 1 6.319 4h9.362a2 2 0 0 1 1.839 1.212l2.318 5.41a2 2 0 0 1 .162.789V12.5"/><path stroke-linejoin="round" d="M17 23s.9-3.118 3-5"/><path stroke-linejoin="round" d="m19.802 21.425l-.134.012a3.094 3.094 0 0 1-3.366-2.774a3.06 3.06 0 0 1 2.761-3.35l2.986-.28a.35.35 0 0 1 .381.314l.255 2.58a3.194 3.194 0 0 1-2.883 3.497"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerifiedBadge(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11.528 1.6a.6.6 0 0 1 .944 0l1.809 2.3a.6.6 0 0 0 .635.207l2.815-.798a.6.6 0 0 1 .764.554l.11 2.925a.6.6 0 0 0 .393.54l2.747 1.01a.6.6 0 0 1 .292.897l-1.63 2.431a.6.6 0 0 0 0 .668l1.63 2.431a.6.6 0 0 1-.292.897l-2.747 1.01a.6.6 0 0 0-.392.54l-.111 2.925a.6.6 0 0 1-.764.554l-2.815-.798a.6.6 0 0 0-.636.206L12.473 22.4a.6.6 0 0 1-.944 0L9.72 20.1a.6.6 0 0 0-.635-.207l-2.815.798a.6.6 0 0 1-.764-.554l-.11-2.925a.6.6 0 0 0-.393-.54l-2.747-1.01a.6.6 0 0 1-.292-.897l1.63-2.431a.6.6 0 0 0 0-.668l-1.63-2.431a.6.6 0 0 1 .292-.897l2.747-1.01a.6.6 0 0 0 .392-.54l.111-2.925a.6.6 0 0 1 .764-.554l2.815.798A.6.6 0 0 0 9.72 3.9l1.81-2.3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 12l2 2l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalMerge(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2v8m0 0l3.5-3.5M12 10L8.5 6.5M12 22v-8m0 0l3.5 3.5M12 14l-3.5 3.5M3 14h18M3 10h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalSplit(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14v8m0 0l3.5-3.5M12 22l-3.5-3.5M12 10V2m0 0l3.5 3.5M12 2L8.5 5.5M3 14h18M3 10h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vials(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21H3m6-9H5m14 0h-4m-8 6a2 2 0 0 1-2-2V3h4v13a2 2 0 0 1-2 2m10 0a2 2 0 0 1-2-2V3h4v13a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCamera(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 12v4.4a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V7.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6zm0 0l5.016-4.18a.6.6 0 0 1 .984.461v7.438a.6.6 0 0 1-.984.46z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCameraOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 7H3.6a.6.6 0 0 0-.6.6v8.8a.6.6 0 0 0 .6.6h10.8a.6.6 0 0 0 .6-.6V15m-3.5-8h2.9a.6.6 0 0 1 .6.6v3.119a.6.6 0 0 0 .984.46l4.032-3.359a.6.6 0 0 1 .984.461V15.5M3 3l18 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoProjector(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M4 19h2m12 0h2"/><path d="M2 16.4V7.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v8.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m5 10.01l.01-.011M8 10.01l.01-.011m2.99.011l.01-.011M5 14.01l.01-.011M8 14.01l.01-.011m2.99.011l.01-.011M17 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumnsThree(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M9 3H3.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6H9M9 3v18M9 3h6M9 21h6m0-18h5.4a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H15m0-18v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumnsTwo(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M12 3h8.4a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H12m0-18H3.6a.6.6 0 0 0-.6.6v16.8a.6.6 0 0 0 .6.6H12m0-18v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewGrid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M14 20.4v-5.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 0v-5.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm11-11V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 0V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewStructureDown(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 20.4v-5.8a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm11-11V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 0V3.6a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewStructureUp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3 9.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Zm11 11v-5.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6h-5.8a.6.6 0 0 1-.6-.6Zm-11 0v-5.8a.6.6 0 0 1 .6-.6h5.8a.6.6 0 0 1 .6.6v5.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewThreeHundredSixty(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 8.5h1.75m0 0a1.75 1.75 0 1 1 0 3.5H3m2.75-3.5a1.75 1.75 0 1 0 0-3.5H3m18 10c0 3.314-4.03 6-9 6s-9-2.686-9-6M14 5h-1a3 3 0 0 0-3 3v2m4.5-.5v.5a2 2 0 0 1-2 2H12a2 2 0 0 1-2-2v-.5a2 2 0 0 1 2-2h.5a2 2 0 0 1 2 2m2.5-1V7a2 2 0 0 1 2-2h.5a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H19a2 2 0 0 1-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Voice(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4v16M8 9v6m12-5v4M4 10v4m12-7v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoiceCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v16M8 8v6m12-5v4M4 9v4m12-7v9m-.5 4.5l2 2l5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoiceCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v12M9 9v6m9-4v2M6 11v2m9-6v10m-3 5c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoiceLockCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 21.8c-.646.131-1.315.2-2 .2c-5.523 0-10-4.477-10-10S6.477 2 12 2s10 4.477 10 10c0 .254-.01.506-.028.755M12 6v12M9 9v6m9-4v2M6 11v2m9-6v10"/><path d="M21.167 18.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoiceScan(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v12M9 9v6m9-4v2M6 11v2m9-6v10M6 3H3v3m15-3h3v3M6 21H3v-3m15 3h3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoiceSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6M12 6v12M9 9v6m9-4v2M6 11v2m9-6v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoiceXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v16M8 8v6m12-5v4M4 9v4m12-7v8m.121 7.364l2.122-2.121m0 0l2.121-2.122m-2.121 2.122L16.12 17.12m2.122 2.122l2.121 2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VrTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 15.5v-2.8m2.857 0c.714 0 2.143 0 2.143-2.1s-1.429-2.1-2.143-2.1H13v4.2m2.857 0H13m2.857 0L18 15.5m-7-7l-3 7l-3-7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VueJs(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 20.5L22.5 4h-4L12 14L5.5 4h-4z"/><path d="M18.5 4h-4L12 7.5L9.5 4h-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Waist(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.4 4s-1.6 3.75-1.6 6.857c0 .995.34 1.827.8 2.656c.528.954 1.214 1.903 1.717 3.09A8.49 8.49 0 0 1 20 20M5.6 4s1.6 3.75 1.6 6.857c0 .995-.34 1.827-.8 2.656c-.528.954-1.214 1.903-1.717 3.09A8.483 8.483 0 0 0 4 20m2.4-6.487h11.2"/><path d="M4.683 16.604S10.4 17.714 12 20c1.6-2.286 7.317-3.396 7.317-3.396"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Walking(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12.44 9.127l-1.408 5.635l4.93 6.339m-5.634-2.817L8.215 21.1"/><path d="M8.215 13.353c0-3.944 2.817-4.226 4.226-4.226h1.408c.235 1.174 1.268 3.663 3.522 4.226M13 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 20H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 14a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/><path d="M18 7V5.603a2 2 0 0 0-2.515-1.932l-11 2.933A2 2 0 0 0 3 8.537V9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7v6m0 4.01l.01-.011M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12M12 6.25a.75.75 0 0 1 .75.75v6a.75.75 0 0 1-1.5 0V7a.75.75 0 0 1 .75-.75m.568 11.25a.75.75 0 0 0-1.115-1.003l-.01.011a.75.75 0 0 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningHexagon(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.173 12.3a.6.6 0 0 1 0-.6l5.154-8.926a.6.6 0 0 1 .52-.3h10.307a.6.6 0 0 1 .52.3l5.153 8.926a.6.6 0 0 1 0 .6l-5.154 8.926a.6.6 0 0 1-.52.3H6.847a.6.6 0 0 1-.52-.3zM12 8v4m0 4.01l.01-.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 7v6m0 4.01l.01-.011"/><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 3.6c0-.746.604-1.35 1.35-1.35h16.8c.746 0 1.35.604 1.35 1.35v16.8a1.35 1.35 0 0 1-1.35 1.35H3.6a1.35 1.35 0 0 1-1.35-1.35zM12 6.25a.75.75 0 0 1 .75.75v6a.75.75 0 0 1-1.5 0V7a.75.75 0 0 1 .75-.75m.568 11.25a.75.75 0 0 0-1.115-1.003l-.01.011a.75.75 0 0 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningTriangle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20.043 21H3.957c-1.538 0-2.5-1.664-1.734-2.997l8.043-13.988c.77-1.337 2.699-1.337 3.468 0l8.043 13.988C22.543 19.336 21.58 21 20.043 21ZM12 9v4"/><path stroke-linejoin="round" d="m12 17.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningTriangleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.616 3.642c1.058-1.839 3.71-1.839 4.768 0l8.043 13.988c1.054 1.833-.27 4.12-2.384 4.12H3.957c-2.115 0-3.438-2.287-2.384-4.12zM12 8.25a.75.75 0 0 1 .75.75v4a.75.75 0 0 1-1.5 0V9a.75.75 0 0 1 .75-.75m.568 9.25a.75.75 0 0 0-1.115-1.003l-.01.011a.75.75 0 0 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M18 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M21 16v2m0 4.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wash(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m22 5l-1.954 12.314A2 2 0 0 1 18.07 19H5.93a2 2 0 0 1-1.975-1.686L2 5"/><path d="M21 11c-2 0-4.5-3-4.5-3s-2.149 3-4.5 3s-4.5-3-4.5-3S5 11 3 11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WashingMachine(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 4v16a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2m-3 1.01l.01-.011"/><path d="M12 19a6 6 0 1 0 0-12a6 6 0 0 0 0 12"/><path d="M12 16a3 3 0 0 1-3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WateringSoil(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 12h2m16 0h2M3 20.01l.01-.011M6 16.01l.01-.011M9 20.01l.01-.011m5.99.011l.01-.011M18 16.01l.01-.011M21 20.01l.01-.011M12.396 3.396L15.5 6.5a4.95 4.95 0 1 1-7 0l3.104-3.104a.56.56 0 0 1 .792 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebWindow(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 17V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 8h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebWindowEnergyConsumption(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 17V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.667 9L10 12h4l-1.667 3M6 8h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebWindowXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 17V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m10 14.243l2.121-2.122m0 0L14.243 10m-2.122 2.121L10 10m2.121 2.121l2.122 2.122M6 8h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebpFormat(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M13.5 15V9h2.4a.6.6 0 0 1 .6.6v.9A1.5 1.5 0 0 1 15 12v0"/><path stroke-linejoin="round" d="M13.5 15h2.4a.6.6 0 0 0 .6-.6v-.9A1.5 1.5 0 0 0 15 12v0h-1.5m6 3v-3m0 0V9h3v3zm-18-3v6L3 12l1.5 3V9m6 0h-3v6h3m-3-3h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weight(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 5h3.9a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h3.9"/><path d="m16.279 6.329l.205-1.23a.605.605 0 0 0 0-.198l-.206-1.23A2 2 0 0 0 14.307 2H9.694a2 2 0 0 0-1.973 1.671l-.205 1.23a.6.6 0 0 0 0 .198l.205 1.23A2 2 0 0 0 9.694 8h4.612a2 2 0 0 0 1.973-1.671M12 8l-1-2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeightAlt(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 5h3.9a.6.6 0 0 1 .6.6v14.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V5.6a.6.6 0 0 1 .6-.6h3.9"/><path d="m16.279 6.329l.205-1.23a.605.605 0 0 0 0-.198l-.206-1.23A2 2 0 0 0 14.307 2H9.694a2 2 0 0 0-1.973 1.671l-.205 1.23a.6.6 0 0 0 0 .198l.205 1.23A2 2 0 0 0 9.694 8h4.612a2 2 0 0 0 1.973-1.671M12 8l-1-2.5M7 17h10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhiteFlag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 15l.95-10.454A.6.6 0 0 1 6.548 4h13.795a.6.6 0 0 1 .598.654l-.891 9.8a.6.6 0 0 1-.598.546zm0 0l-.6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhiteFlagSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path fill="currentColor" d="m20.94 4.654l-.89 9.8a.6.6 0 0 1-.598.546H5l.95-10.454A.6.6 0 0 1 6.548 4h13.795a.6.6 0 0 1 .598.654"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 15l.95-10.454A.6.6 0 0 1 6.548 4h13.795a.6.6 0 0 1 .598.654l-.891 9.8a.6.6 0 0 1-.598.546zm0 0l-.6 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 19.51l.01-.011M2 8c6-4.5 14-4.5 20 0M5 12c4-3 10-3 14 0M8.5 15.5c2.25-1.4 4.75-1.4 7 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiOff(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 19.51l.01-.011M3 3l18 18M2 8a17.053 17.053 0 0 1 3.757-2.14M22 8c-3.572-2.68-7.854-3.763-12-3.252M5 12c1.333-1 2.889-1.667 4.518-2M19 12a11.274 11.274 0 0 0-4.282-1.95M8.5 15.5c2.25-1.4 4.75-1.4 7 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSignalNone(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M2.126 8.324c-.2-.262-.155-.605.086-.79C5.29 5.179 8.552 4 11.999 4c3.447 0 6.71 1.178 9.788 3.535c.252.212.28.558.085.789l-9.455 11.173a.548.548 0 0 1-.836 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSignalNoneSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.756 6.94c3.192-2.444 6.612-3.69 10.243-3.69c3.632 0 7.051 1.246 10.244 3.69l.014.01l.013.011c.552.465.653 1.282.175 1.847L12.99 19.981a1.296 1.296 0 0 1-1.981 0L1.543 8.795l-.011-.014c-.44-.572-.364-1.392.224-1.842" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiTag(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 15V9a6 6 0 0 1 6-6h8a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6H8a6 6 0 0 1-6-6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m12 14.76l.01-.011M7 11.25c2.5-3 7.5-3 10 0m-8 2c1.5-2 4.5-2 6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiTagSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.25A6.75 6.75 0 0 0 1.25 9v6A6.75 6.75 0 0 0 8 21.75h8A6.75 6.75 0 0 0 22.75 15V9A6.75 6.75 0 0 0 16 2.25zm4.568 13a.75.75 0 0 0-1.115-1.003l-.01.011a.75.75 0 0 0 1.114 1.004zm3.856-3.52c-2.2-2.64-6.648-2.64-8.848 0a.75.75 0 1 1-1.152-.96c2.8-3.36 8.352-3.36 11.152 0a.75.75 0 1 1-1.152.96M9.6 13.7a3 3 0 0 1 4.8 0a.75.75 0 1 0 1.2-.9c-1.8-2.4-5.4-2.4-7.2 0a.75.75 0 0 0 1.2.9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiWarning(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2.126 8.324c-.2-.262-.155-.605.086-.79C5.29 5.179 8.552 4 11.999 4c3.447 0 6.71 1.178 9.788 3.535c.252.212.28.558.085.789l-9.455 11.173a.548.548 0 0 1-.836 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 8v2m0 4.01l.01-.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiWarningSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.756 6.94c3.192-2.444 6.612-3.69 10.243-3.69c3.632 0 7.051 1.246 10.244 3.69l.014.01l.013.011c.552.465.653 1.282.175 1.847L12.99 19.981a1.296 1.296 0 0 1-1.981 0L1.543 8.795l-.011-.014c-.44-.572-.364-1.392.224-1.842M12 7.25a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0V8a.75.75 0 0 1 .75-.75m.568 7.25a.75.75 0 0 0-1.115-1.003l-.01.011a.75.75 0 0 0 1.114 1.004z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 18.51l.01-.011M2 7c6-4.5 14-4.5 20 0M5 11c4-3 10-3 14 0M8.5 14.5c2.25-1.4 4.75-1.4 7 0m1.621 6.864l2.122-2.121m2.121-2.122l-2.121 2.122m0 0L17.12 17.12m2.122 2.122l2.121 2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wind(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.279 7C19.782 7 21 8.12 21 9.5S19.782 12 18.279 12H3m14.938 8c1.139 0 2.562-.5 2.562-2.5S19.077 15 17.938 15H3m7.412-11C11.842 4 13 5.12 13 6.5S11.841 9 10.412 9H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowCheck(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M13 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M16 20l2 2l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowLock(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M21.167 18.5h.233a.6.6 0 0 1 .6.6v2.3a.6.6 0 0 1-.6.6h-3.8a.6.6 0 0 1-.6-.6v-2.3a.6.6 0 0 1 .6-.6h.233m3.334 0v-1.75c0-.583-.334-1.75-1.667-1.75s-1.667 1.167-1.667 1.75v1.75m3.334 0h-3.334"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowNoAccess(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M21 16.05a3.5 3.5 0 1 0-5 4.9m5-4.9a3.5 3.5 0 0 1-5 4.9m5-4.9l-5 4.9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowXmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M15 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10"/><path stroke-linejoin="round" d="M2 7h20M5 5.01l.01-.011M8 5.01l.01-.011M11 5.01l.01-.011M18 22.243l2.121-2.122m0 0L22.243 18m-2.122 2.121L18 18m2.121 2.121l2.122 2.122"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M4 16.987V7.013a.6.6 0 0 1 .507-.593l14.8-2.313A.6.6 0 0 1 20 4.7v14.598a.6.6 0 0 1-.693.593l-14.8-2.313A.6.6 0 0 1 4 16.986ZM4 12h16m-9.5-6.5v13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wolf(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.812 7s-.453.628-.996 1.667M18.188 7s.453.628.997 1.667m-14.37 0C4.008 10.214 3 12.674 3 15.333C5.813 15.333 7.5 17 7.5 17s1.125 5 4.5 5s4.5-5 4.5-5s1.688-1.667 4.5-1.667c0-2.659-1.007-5.119-1.815-6.666m-14.37 0s-2.94-2.223 0-6.667c.997.556 3.81 2.778 3.81 2.778S10.313 3.667 12 3.667c1.688 0 3.375 1.11 3.375 1.11S18.188 2.557 19.313 2c2.812 4.445-.128 6.667-.128 6.667M11 18h1m1 0h-1m0 0v1m-3.5-6.5L10 14m5.5-1.5L14 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrapText(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 7h16M4 17h5m-5-5h13.5a2.5 2.5 0 0 1 2.5 2.5v0a2.5 2.5 0 0 1-2.5 2.5h-5"/><path d="M15 15.5L12.5 17l2.5 1.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.05 10.607l-7.07 7.07a2 2 0 0 0 0 2.83v0a2 2 0 0 0 2.828 0l7.07-7.072m-2.828-2.828c-.844-2.153-.679-4.978 1.06-6.718c1.74-1.74 4.95-2.121 6.718-1.06l-3.04 3.04l-.283 3.111l3.111-.282l3.04-3.041c1.062 1.768.68 4.978-1.06 6.717c-1.74 1.74-4.564 1.905-6.717 1.061"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wristwatch(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 16.472V20a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-3.528m0-8.944V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v3.528"/><path d="M18 12a6 6 0 1 0-12 0a6 6 0 0 0 12 0"/><path d="M14 12h-2v-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Www(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.338 17A9.996 9.996 0 0 0 12 22a9.996 9.996 0 0 0 8.662-5M3.338 7A9.996 9.996 0 0 1 12 2a9.996 9.996 0 0 1 8.662 5"/><path d="M13 21.95s1.408-1.853 2.295-4.95M13 2.05S14.408 3.902 15.295 7M11 21.95S9.592 20.098 8.705 17M11 2.05S9.592 3.902 8.705 7M9 10l1.5 5l1.5-5l1.5 5l1.5-5M1 10l1.5 5L4 10l1.5 5L7 10m10 0l1.5 5l1.5-5l1.5 5l1.5-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func X(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16.82 20.768L3.753 3.968A.6.6 0 0 1 4.227 3h2.48a.6.6 0 0 1 .473.232l13.067 16.8a.6.6 0 0 1-.474.968h-2.48a.6.6 0 0 1-.473-.232Z"/><path stroke-linecap="round" d="M20 3L4 21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xsquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6M10 8l4 8m0-8l-4 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XboxA(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="m15 16l-3-8l-3 8m5-2h-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XboxB(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/><path d="M12.599 11.826c2.535 0 2.535 4.174 0 4.174H9.5v-4.174m3.099 0H9.5m3.099 0c2.535 0 2.535-3.826 0-3.826H9.5v3.826"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XboxX(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10m3-6L9 8m0 8l6-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XboxY(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10M9 8l3 5"/><path d="M12 16v-3l3-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xmark(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.758 17.243L12.001 12m5.243-5.243L12 12m0 0L6.758 6.757M12.001 12l5.243 5.243"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XmarkCircle(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 14.828L12.001 12m2.828-2.828L12.001 12m0 0L9.172 9.172M12.001 12l2.828 2.828M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XmarkCircleSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25M9.702 8.641a.75.75 0 0 0-1.061 1.06L10.939 12l-2.298 2.298a.75.75 0 0 0 1.06 1.06L12 13.062l2.298 2.298a.75.75 0 0 0 1.06-1.06L13.06 12l2.298-2.298a.75.75 0 1 0-1.06-1.06L12 10.938z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XmarkSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.879 14.121L12 12m2.121-2.121L12 12m0 0L9.879 9.879M12 12l2.121 2.121M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XmarkSquareSolid(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.6 2.25A1.35 1.35 0 0 0 2.25 3.6v16.8c0 .746.604 1.35 1.35 1.35h16.8a1.35 1.35 0 0 0 1.35-1.35V3.6a1.35 1.35 0 0 0-1.35-1.35zm6.809 7.098a.75.75 0 0 0-1.06 1.061L10.938 12l-1.59 1.591a.75.75 0 0 0 1.06 1.06L12 13.062l1.591 1.59a.75.75 0 0 0 1.06-1.06L13.062 12l1.59-1.591a.75.75 0 0 0-1.06-1.06L12 10.938z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XrayView(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 7.353v9.294a.6.6 0 0 1-.309.525l-8.4 4.666a.6.6 0 0 1-.582 0l-8.4-4.666A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0l8.4 4.667a.6.6 0 0 1 .309.524"/><path d="m20.5 16.722l-8.209-4.56a.6.6 0 0 0-.582 0L3.5 16.722m.028-9.428l8.18 4.544a.6.6 0 0 0 .583 0l8.209-4.56M12 21V3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ysquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6M10 8l2 4"/><path d="m14 8l-2 4v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yelp(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.5 11l-.458-8.24a.6.6 0 0 0-.771-.541L6.814 3.256a.6.6 0 0 0-.311.93zm2.5 1.5l4.57-.83a.6.6 0 0 0 .38-.94l-1.445-2.023a.6.6 0 0 0-.987.016zm.5 3.5l2.066 4.132a.6.6 0 0 0 1.017.091l1.835-2.446a.6.6 0 0 0-.373-.95zm-3 .5l-3.341 3.341a.6.6 0 0 0 .213.986l2.317.869a.6.6 0 0 0 .811-.562zm-2-2.5l-4.132-2.066a.6.6 0 0 0-.868.537v2.643a.6.6 0 0 0 .823.557z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yen(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h12M6 4l6 8m6-8l-6 8m0 0v8m-6-4h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YenSquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 13h8M8 7l4 5.5M16 7l-4 5.5m0 0V18m-4-3h8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yoga(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.571 15.004l.858 1.845s3.857.819 3.857 2.767C19.286 21 17.57 21 17.57 21H13l-2.25-1.25"/><path d="m9.429 15.004l-.857 1.845s-3.858.819-3.858 2.767C4.714 21 6.43 21 6.43 21H8.5l2.25-1.25L13.5 18"/><path d="M3 15.926s2.143-.461 3.429-.922C7.714 8.546 11.57 9.007 12 9.007c.429 0 4.286-.461 5.571 5.997c1.286.46 3.429.922 3.429.922M12 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14 12l-3.5 2v-4z"/><path d="M2 12.708v-1.416c0-2.895 0-4.343.905-5.274c.906-.932 2.332-.972 5.183-1.053C9.438 4.927 10.818 4.9 12 4.9c1.181 0 2.561.027 3.912.065c2.851.081 4.277.121 5.182 1.053c.906.931.906 2.38.906 5.274v1.415c0 2.896 0 4.343-.905 5.275c-.906.931-2.331.972-5.183 1.052c-1.35.039-2.73.066-3.912.066a141.1 141.1 0 0 1-3.912-.066c-2.851-.08-4.277-.12-5.183-1.052C2 17.05 2 15.602 2 12.708Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zsquare(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 3.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6"/><path d="M10 8h4l-4 8h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 11h3m3 0h-3m0 0V8m0 3v3m6 3l4 4M3 11a8 8 0 1 0 16 0a8 8 0 0 0-16 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *IconoirIcon {
	return &IconoirIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 17l4 4M3 11a8 8 0 1 0 16 0a8 8 0 0 0-16 0m5 0h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
