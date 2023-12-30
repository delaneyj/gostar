package teenyicons

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "0.4.1"
	hAttr          = "1em"
	viewbox        = "0 0 15 15"
	fill           = "currentColor"
)

type TeenyiconsIcon struct {
	*SVGSVGElement
}

func AbTestingOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 11V6.5a2 2 0 1 1 4 0V11m-4-2.5h4m6.5-1H9.5m2.5 0a1.5 1.5 0 0 0 0-3H9.5v3m2.5 0a1.5 1.5 0 0 1 0 3H9.5v-3M7.5 1v13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbTestingSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 5A1.5 1.5 0 0 0 2 6.5V8h3V6.5A1.5 1.5 0 0 0 3.5 5"/><path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4A1.5 1.5 0 0 1 7 2.5v10A1.5 1.5 0 0 1 5.5 14h-4A1.5 1.5 0 0 1 0 12.5zM2 11V9h3v2h1V6.5a2.5 2.5 0 0 0-5 0V11z" clip-rule="evenodd"/><path fill="currentColor" d="M12 7h-2V5h2a1 1 0 1 1 0 2m0 3h-2V8h2a1 1 0 1 1 0 2"/><path fill="currentColor" fill-rule="evenodd" d="M8 2.5A1.5 1.5 0 0 1 9.5 1h4A1.5 1.5 0 0 1 15 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-4A1.5 1.5 0 0 1 8 12.5zM12 4H9v7h3a2 2 0 0 0 1.323-3.5A2 2 0 0 0 12 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 1v13M1 7.5h13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 4v7M4 7.5h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 7V4h1v3h3v1H8v3H7V8H4V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 7V1h1v6h6v1H8v6H7V8H1V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressBookOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 11.5H5v.5h.5zm5 0v.5h.5v-.5zm-4.5 0V11H5v.5zm4-.5v.5h1V11zm.5 0h-5v1h5zM8 9a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3zm-2 2a2 2 0 0 1 2-2V8a3 3 0 0 0-3 3zm2-8a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1zm2 2a2 2 0 0 0-2-2v1a1 1 0 0 1 1 1zM8 7a2 2 0 0 0 2-2H9a1 1 0 0 1-1 1zm0-1a1 1 0 0 1-1-1H6a2 2 0 0 0 2 2zM3.5 1h9V0h-9zm9.5.5v12h1v-12zM12.5 14h-9v1h9zM3 13.5v-12H2v12zm.5.5a.5.5 0 0 1-.5-.5H2A1.5 1.5 0 0 0 3.5 15zm9.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM12.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 12.5 0zm-9-1A1.5 1.5 0 0 0 2 1.5h1a.5.5 0 0 1 .5-.5zM4 4H1v1h3zm0 6H1v1h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressBookSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 0A1.5 1.5 0 0 0 2 1.5V4H1v1h1v5H1v1h1v2.5A1.5 1.5 0 0 0 3.5 15h9a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 12.5 0zM6 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0m-1 6a3 3 0 1 1 6 0v1H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustHorizontalAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 3.5H6.5m0 0a2 2 0 1 0-4 0m4 0a2 2 0 1 1-4 0m0 0H0m15 8h-2.5m0 0a2 2 0 1 0-4 0m4 0a2 2 0 1 1-4 0m0 0H0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustHorizontalAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 1a2.5 2.5 0 0 0-2.45 2H0v1h2.05a2.5 2.5 0 0 0 4.9 0H15V3H6.95A2.5 2.5 0 0 0 4.5 1m6 8a2.5 2.5 0 0 0-2.45 2H0v1h8.05a2.5 2.5 0 0 0 4.9 0H15v-1h-2.05a2.5 2.5 0 0 0-2.45-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustHorizontalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 7.5H0m15 5h-2.5m2.5-10H8.5m-2 0H0m4.5 5H15m-4.5 5H0m10.5-2v4h2v-4zm-8-5v4h2v-4zm4-5v4h2v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustHorizontalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 0H6v2H0v1h6v2h3V3h6V2H9zM5 5H2v2H0v1h2v2h3V8h10V7H5zm8 5h-3v2H0v1h10v2h3v-2h2v-1h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustVerticalAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v8.5m0 0a2 2 0 1 0 0 4m0-4a2 2 0 1 1 0 4m0 0V15m8-15v2.5m0 0a2 2 0 1 0 0 4m0-4a2 2 0 1 1 0 4m0 0V15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustVerticalAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0v8.05a2.5 2.5 0 0 0 0 4.9V15h1v-2.05a2.5 2.5 0 0 0 0-4.9V0zm8 0v2.05a2.5 2.5 0 0 0 0 4.9V15h1V6.95a2.5 2.5 0 0 0 0-4.9V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustVerticalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 12.5V15m5-15v2.5M2.5 0v6.5m0 2V15m5-4.5V0m5 4.5V15m-2-10.5h4v-2h-4zm-5 8h4v-2h-4zm-5-4h4v-2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustVerticalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0v6H0v3h2v6h1V9h2V6H3V0zm3 10h2V0h1v10h2v3H8v2H7v-2H5zm7-10v2h-2v3h2v10h1V5h2V2h-2V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplayOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M3 11.5H.5v-10h14v10H12m-4.5-2l-4 4h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 1a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5H3v-1H1V2h13v9h-2v1h2.5a.5.5 0 0 0 .5-.5v-10a.5.5 0 0 0-.5-.5z"/><path fill="currentColor" d="M7.854 9.146a.5.5 0 0 0-.708 0l-4 4A.5.5 0 0 0 3.5 14h8a.5.5 0 0 0 .354-.854z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirpodsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 3.5a2.648 2.648 0 0 1-2.977 2.628l-.32-.04a2.667 2.667 0 0 1-1.27-.513L.5 4.5v-2l1.433-1.075a2.667 2.667 0 0 1 1.27-.513l.32-.04A2.648 2.648 0 0 1 6.5 3.5Zm0 0v11h-2V6m4-2.5a2.648 2.648 0 0 0 2.977 2.628l.32-.04c.46-.058.898-.234 1.27-.513L14.5 4.5v-2l-1.433-1.075a2.667 2.667 0 0 0-1.27-.513l-.32-.04A2.648 2.648 0 0 0 8.5 3.5Zm0 0v11h2V6M2 3.5h2m7 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirpodsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 3.5A3.148 3.148 0 0 0 3.461.376l-.32.04a3.167 3.167 0 0 0-1.508.609L0 2.25v2.5l1.633 1.225c.441.33.96.54 1.508.609l.32.04c.182.023.362.03.539.022V15h3zM4 4H2V3h2zm4-.5A3.148 3.148 0 0 1 11.539.376l.32.04a3.167 3.167 0 0 1 1.508.609L15 2.25v2.5l-1.633 1.225c-.441.33-.96.54-1.508.609l-.32.04a3.18 3.18 0 0 1-.539.022V15H8zm3 .5h2V3h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5 3.5l3-3m8 0l3 3M7.5 5v3.5H10m-2.5-6a6 6 0 1 0 0 12a6 6 0 0 0 0-12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.854.854l-3 3l-.708-.708l3-3zm10.293 3l-3-3l.707-.708l3 3z"/><path fill="currentColor" fill-rule="evenodd" d="M1 8.5a6.5 6.5 0 1 1 13 0a6.5 6.5 0 0 1-13 0M8 8V5H7v4h3V8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlienOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" d="M6.52 11.435c.24.107.719.497.981.497c.263 0 .741-.39.982-.497m-3.926-4.97l1.472 1.49m4.417-1.49l-1.472 1.49M7.5.5C3.94.5 1.967 2.45 1.612 4.974c-.358 2.33.136 4.53 1.472 6.461c.643.953 1.486 1.876 2.454 2.486c1.271.772 2.654.772 3.925 0c.967-.61 1.81-1.533 2.454-2.486c1.33-1.934 1.824-4.13 1.472-6.461C13.034 2.449 11.062.5 7.501.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlienSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.039 1.403C4.127.513 5.629 0 7.501 0c1.873 0 3.374.514 4.463 1.403c1.089.89 1.727 2.125 1.921 3.498c.37 2.453-.151 4.776-1.554 6.816c-.668.99-1.558 1.97-2.6 2.627l-.007.004c-1.431.87-3.014.87-4.445 0l-.007-.004c-1.042-.657-1.931-1.637-2.6-2.626C1.264 9.68.742 7.353 1.118 4.9c.194-1.373.832-2.608 1.921-3.498M6.736 7.96l-.711.703L3.85 6.46l.712-.702zm1.53 0l.712.703l2.175-2.203l-.712-.702zm-2.001 2.816l.457.202c.097.043.2.105.28.156a15.37 15.37 0 0 1 .26.163a2.036 2.036 0 0 0 .24.127a2.036 2.036 0 0 0 .24-.128A8.542 8.542 0 0 0 8 11.134c.082-.051.184-.113.28-.156l.458-.202l.404.915l-.457.202c-.024.01-.071.037-.156.09l-.086.054l-.183.114c-.097.059-.21.124-.324.175a1.084 1.084 0 0 1-.435.106c-.173 0-.33-.06-.434-.106a3.003 3.003 0 0 1-.324-.175c-.06-.036-.125-.078-.183-.114l-.087-.055a1.543 1.543 0 0 0-.155-.09l-.457-.201z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignBottomOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 14.5H0m11.5-3h-3v-7h3zm-5 0h-3V.5h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignBottomSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0H3v12h4zm5 4H8v8h4zm3 10H0v1h15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterHorizontalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 0v3.5m0 8V15m0-8.5v2M4 3.5v3h7v-3zm-2.5 5v3h12v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterHorizontalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0v3H3.5v4H7v1H1v4h6v3h1v-3h6V8H8V7h3.5V3H8V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterVerticalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 7.5h-3.5m-8 0H0m8.5 0h-2m5 3.5h-3V4h3zm-5 2.5h-3v-12h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterVerticalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1H3v6H0v1h3v6h4V8h1v3.5h4V8h3V7h-3V3.5H8V7H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 0v15m3-11.5v3h7v-3zm0 5v3h11v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v15h1V0zm11 3H3v4h8zm4 5H3v4h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 0v15m-3-11.5v3h-7v-3zm0 5v3H.5v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 0v15h1V0zm-2 3H4v4h8zm0 5H0v4h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTextCenterOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 3.5H0m10 4H5m7 4H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTextCenterSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 3h15v1H0zm5 4h5v1H5zm-2 4h9v1H3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTextJustifyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 3.5h15m-15 8h15m-15-4h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTextJustifySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 4H0V3h15zm0 4H0V7h15zm0 4H0v-1h15z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTextLeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 3.5h15m-15 8h9m-9-4h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTextLeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 4H0V3h15zM6 8H0V7h6zm3 4H0v-1h9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTextRightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 3.5H0m15 8H6m9-4H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTextRightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 3h15v1H0zm9 4h6v1H9zm-3 4h9v1H6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTopOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 .5H0m11.5 3h-3v7h3zm-5 0h-3v11h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTopSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 0H0v1h15zM7 3H3v12h4zm5 0H8v8h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnchorOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5 4.497a2 2 0 0 0 2-1.998a2 2 0 0 0-4 0a2 2 0 0 0 2 1.998Zm0 0v9.994m0 0c-3.866 0-7-3.132-7-6.996h2m5 6.996c3.866 0 7-3.132 7-6.996h-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnchorSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 1A1.5 1.5 0 0 0 6 2.499a1.5 1.5 0 0 0 3 0A1.5 1.5 0 0 0 7.5 1M5 2.499a2.5 2.5 0 1 1 3 2.448v9.025a6.499 6.499 0 0 0 5.981-5.977H12v-1h3v.5a7.498 7.498 0 0 1-7.5 7.496A7.498 7.498 0 0 1 0 7.495v-.5h3v1H1.019A6.499 6.499 0 0 0 7 13.972V4.947A2.5 2.5 0 0 1 5 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m2.5 2.5l2 3m8-3l-2 3M4 9.5h1m5 0h1m-9.5 3v-2a6 6 0 1 1 12 0v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 4a6.473 6.473 0 0 0-2.934.698l-1.65-2.475l-.832.554l1.627 2.44A6.492 6.492 0 0 0 1 10.5V13h13v-2.5a6.492 6.492 0 0 0-2.711-5.282l1.627-2.44l-.832-.555l-1.65 2.475A6.473 6.473 0 0 0 7.5 4M5 10H4V9h1zm5 0h1V9h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngularOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.137-.48a.5.5 0 0 0-.274 0zm-7 2l-.137-.48a.5.5 0 0 0-.36.535zm1 9l-.497.055a.5.5 0 0 0 .273.392zm6 3l-.224.447a.5.5 0 0 0 .448 0zm6-3l.224.447a.5.5 0 0 0 .273-.392zm1-9l.497.055a.5.5 0 0 0-.36-.536zm-7 .5l.458-.2L7.5 1.753L7.042 2.8zM7.363.02l-7 2l.274.96l7-2zM.003 2.554l1 9l.994-.11l-1-9zm1.273 9.392l6 3l.448-.894l-6-3zm6.448 3l6-3l-.448-.894l-6 3zm6.273-3.392l1-9l-.994-.11l-1 9zm.64-9.536l-7-2l-.274.962l7 2zM4.458 11.2l3.5-8l-.916-.4l-3.5 8zm2.584-8l3.5 8l.916-.4l-3.5-8zM5 9h5V8H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngularSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 4.247L9.142 8H5.858z"/><path fill="currentColor" fill-rule="evenodd" d="M7.363.02a.5.5 0 0 1 .274 0l7 2a.5.5 0 0 1 .36.535l-1 9a.5.5 0 0 1-.273.392l-6 3a.5.5 0 0 1-.448 0l-6-3a.5.5 0 0 1-.273-.392l-1-9a.5.5 0 0 1 .36-.536zM7.5 1.752l3.958 9.047l-.916.4L9.579 9H5.421l-.963 2.2l-.916-.4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnjaOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.007 4.4l.034-.498zM2.955 8.139l-.461-.194zM7.5 14c-1.627 0-3.1-.958-3.762-2.444l-.913.406A5.116 5.116 0 0 0 7.5 15zm3.762-2.444A4.116 4.116 0 0 1 7.5 14v1a5.116 5.116 0 0 0 4.675-3.038zM2 15V6H1v9zM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0zM2 6.5A5.5 5.5 0 0 1 7.5 1V0A6.5 6.5 0 0 0 1 6.5zm11 0V15h1V6.5zM8.974 4.9l1.493.099l.066-.998l-1.492-.1zM11 4.5V3h-1v1.5zM3.416 8.331A5.624 5.624 0 0 1 8.974 4.9l.067-.997a6.624 6.624 0 0 0-6.547 4.042zm6.73-3.477c.47.47 1.33 1.49 1.69 3.342l.98-.19c-.406-2.1-1.394-3.291-1.962-3.86zM6.5 13H8v-1H6.5zM8 13a2 2 0 0 0 2-2H9a1 1 0 0 1-1 1zm4.5-4h-2v1h2zm-4-1H12V7H8.5zm-5 0h3V7h-3zm3 0h2V7h-2zm-2 1h-2v1h2zM6 7.5A1.5 1.5 0 0 1 4.5 9v1A2.5 2.5 0 0 0 7 7.5zM10.5 9A1.5 1.5 0 0 1 9 7.5H8a2.5 2.5 0 0 0 2.5 2.5zm-6.762 2.556c-.495-1.116-.73-2.255-.322-3.225l-.922-.387c-.57 1.36-.2 2.826.33 4.018zm8.437.406c.53-1.19.91-2.574.642-3.957l-.982.19c.212 1.09-.08 2.25-.574 3.36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnjaSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 13a2 2 0 0 0 2-2H9a1 1 0 0 1-1 1H6.5v1z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 15a5.14 5.14 0 0 1-.543-.029L7 15H1V6h.019A6.5 6.5 0 0 1 14 6.5V15zM8.974 4.9l1.3.086A5.951 5.951 0 0 1 11.503 7H4.207a5.616 5.616 0 0 1 4.767-2.1M9.085 8h2.71l.04.196a4 4 0 0 1 .07.804H10.5a1.5 1.5 0 0 1-1.415-1m1.415 2h1.286a8.016 8.016 0 0 1-.524 1.556a4.116 4.116 0 0 1-7.524 0A6.23 6.23 0 0 1 3.254 10H4.5a2.5 2.5 0 0 0 2.45-2h1.1a2.5 2.5 0 0 0 2.45 2m-6-1H3.235c.032-.23.09-.453.18-.669c.048-.112.099-.223.153-.331h2.347A1.5 1.5 0 0 1 4.5 9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AntiClockwiseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.5 2.499l-.354-.354l-.353.354l.353.353zm1-.5H7v1h.5zM2 8.495v-.5H1v.5zM8.145.146l-1.999 2l.708.706L8.853.854zM6.146 2.852l2 1.999l.707-.707l-2-1.999zM7.5 3C10.537 3 13 5.461 13 8.496h1A6.499 6.499 0 0 0 7.5 2zM13 8.495a5.499 5.499 0 0 1-5.5 5.496v1c3.589 0 6.5-2.909 6.5-6.496zM7.5 13.99A5.499 5.499 0 0 1 2 8.495H1a6.499 6.499 0 0 0 6.5 6.496z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AntiClockwiseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.145.146l.708.708l-1.149 1.148A6.499 6.499 0 0 1 14 8.495a6.499 6.499 0 0 1-6.5 6.496A6.499 6.499 0 0 1 1 8.495v-.5h1v.5a5.499 5.499 0 0 0 5.5 5.496A5.5 5.5 0 0 0 13 8.495a5.499 5.499 0 0 0-5.289-5.492l1.142 1.14l-.708.708l-2.352-2.352z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.825 3.241a.343.343 0 0 1-.343-.342A2.399 2.399 0 0 1 9.881.5c.19 0 .342.154.342.343A2.399 2.399 0 0 1 7.826 3.24Zm5.003 7.216c.132.099.175.28.1.427c-1.205 2.414-2.168 3.616-3.047 3.616c-.409 0-.811-.132-1.203-.39a1.782 1.782 0 0 0-1.895-.041c-.474.284-.927.431-1.356.431C4.133 14.5 2 10.518 2 8.332C2 6 3.223 4.22 5.084 4.22c.875 0 1.631.13 2.266.39c.269.112.573.104.836-.022c.515-.248 1.194-.368 2.038-.368c1.03 0 1.926.513 2.672 1.508a.343.343 0 0 1-.068.48c-.833.624-1.234 1.326-1.234 2.124c0 .799.401 1.5 1.234 2.125Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.875 2.937A.371.371 0 0 1 7.5 2.57C7.5 1.15 8.676 0 10.124 0a.37.37 0 0 1 .375.367c0 1.42-1.175 2.57-2.624 2.57"/><path fill="currentColor" d="M7.875 2.937A.371.371 0 0 1 7.5 2.57C7.5 1.15 8.676 0 10.124 0a.37.37 0 0 1 .375.367c0 1.42-1.175 2.57-2.624 2.57m5.475 7.731c.145.106.192.3.11.458C12.14 13.712 11.087 15 10.125 15c-.448 0-.888-.142-1.317-.418a1.985 1.985 0 0 0-2.073-.044c-.52.305-1.015.462-1.485.462c-1.415 0-3.75-4.267-3.75-6.608c0-2.5 1.339-4.406 3.375-4.406c.958 0 1.785.138 2.48.419c.294.118.627.11.914-.025c.564-.266 1.308-.394 2.23-.394c1.127 0 2.11.55 2.926 1.615a.362.362 0 0 1-.075.514c-.911.67-1.35 1.421-1.35 2.277c0 .855.439 1.607 1.35 2.276"/><path fill="currentColor" d="M13.35 10.668c.145.106.192.3.11.458C12.14 13.712 11.087 15 10.125 15c-.448 0-.888-.142-1.317-.418a1.985 1.985 0 0 0-2.073-.044c-.52.305-1.015.462-1.485.462c-1.415 0-3.75-4.267-3.75-6.608c0-2.5 1.339-4.406 3.375-4.406c.958 0 1.785.138 2.48.419c.294.118.627.11.914-.025c.564-.266 1.308-.394 2.23-.394c1.127 0 2.11.55 2.926 1.615a.362.362 0 0 1-.075.514c-.911.67-1.35 1.421-1.35 2.277c0 .855.439 1.607 1.35 2.276"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppointmentsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5M3 7.5h3m6 0H9m-6 3h3m3 0h3m-10.5-8h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppointmentsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2h1.5A1.5 1.5 0 0 1 15 3.5v10a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-10A1.5 1.5 0 0 1 1.5 2H3V0h1v2h7V0h1zM6 8H3V7h3zm6-1H9v1h3zm-6 4H3v-1h3zm3 0h3v-1H9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 8.5h5M.5.5h14v4H.5zm1 4v9a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h15v5H0z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6v7.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V6zm9 3H5V8h5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChartAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 14.5H0v.5h.5zm1.5-3v.5h1v-.5zm13-9V2h-1v.5zM0 0v14.5h1V0zm.5 15H15v-1H.5zM3 11.5c0-1.454.244-2.88.707-3.922C4.177 6.52 4.798 6 5.5 6V5c-1.298 0-2.178.98-2.707 2.172C2.256 8.38 2 9.954 2 11.5zM5.5 6c.32 0 .642.158 1.005.492c.366.338.713.798 1.095 1.308c.368.49.77 1.03 1.217 1.442c.45.416 1.004.758 1.683.758V9c-.32 0-.642-.158-1.005-.492C9.13 8.17 8.782 7.71 8.4 7.2c-.368-.49-.77-1.03-1.217-1.442C6.733 5.342 6.179 5 5.5 5zm5 4c1.223 0 2.363-.763 3.173-2.045C14.485 6.668 15 4.819 15 2.5h-1c0 2.18-.485 3.832-1.173 4.92C12.137 8.514 11.277 9 10.5 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChartAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0H0v15h15V2.5a.5.5 0 1 0-1 0c0 2.18-.485 3.832-1.173 4.92C12.137 8.514 11.277 9 10.5 9c-.32 0-.642-.158-1.005-.492C9.13 8.17 8.782 7.71 8.4 7.2l-.016-.021c-.363-.485-.761-1.015-1.201-1.421C6.733 5.342 6.179 5 5.5 5c-1.298 0-2.178.98-2.707 2.172C2.256 8.38 2 9.954 2 11.5v.197c0 .842-.37 1.636-1 2.177z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChartOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 14.5H0v.5h.5zm6-9l.4-.3a.5.5 0 0 0-.816.023zm3 4l-.4.3a.5.5 0 0 0 .807-.01zM0 0v14.5h1V0zm.5 15H15v-1H.5zm2.416-3.223l4-6l-.832-.554l-4 6zM6.1 5.8l3 4l.8-.6l-3-4zm3.807 3.99l5-7l-.814-.58l-5 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChartSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0H0v15h15V2.5a.5.5 0 0 0-.907-.29L9.49 8.653L6.9 5.2a.5.5 0 0 0-.816.023L1 12.849z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.854 8.146L5.5 7.793l-.707.707l.353.354zM7.5 10.5l-.354.354l.354.353l.354-.353zm2.354-1.646l.353-.354l-.707-.707l-.354.353zM.5 7.5H0zm7-7V0zm0 14V14zm7-7H14zM5.146 8.854l2 2l.708-.708l-2-2zm2.708 2l2-2l-.708-.708l-2 2zM8 10.5V4H7v6.5zm-7-3A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zm1 0A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 1 0 15a7.5 7.5 0 0 1 0-15m2.707 8.5L7.5 11.207L4.793 8.5l.707-.707l1.5 1.5V4h1v5.293l1.5-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m7.5 13.5l4-4m-4 4l-4-4m4 4V1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.854 8.854l.353-.354l-.707-.707l-.354.353zM7.5 10.5l-.354.354l.354.353l.354-.353zM5.854 8.146L5.5 7.793l-.707.707l.353.354zm3.292 0l-2 2l.708.708l2-2zm-1.292 2l-2-2l-.708.708l2 2zM8 10.5V4H7v6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 4v5.293l1.5-1.5l.707.707L7.5 11.207L4.793 8.5l.707-.707l1.5 1.5V4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1v11.293l3.146-3.147l.708.708L7.5 14.207L3.146 9.854l.708-.708L7 12.293V1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.854 5.854l.353-.354l-.707-.707l-.354.353zM4.5 7.5l-.354-.354l-.353.354l.353.354zm1.646 2.354l.354.353l.707-.707l-.353-.354zM7.5.5V0zm7 7h.5zm-14 0H1zm7 7V14zM6.146 5.146l-2 2l.708.708l2-2zm-2 2.708l2 2l.708-.708l-2-2zM4.5 8H11V7H4.5zm3-7A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zm0 1A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0m-8.5 2.707L3.793 7.5L6.5 4.793l.707.707l-1.5 1.5H11v1H5.707l1.5 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m1.5 7.5l4-4m-4 4l4 4m-4-4H14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.146 9.854l.354.353l.707-.707l-.353-.354zM4.5 7.5l-.354-.354l-.353.354l.353.354zm2.354-1.646l.353-.354l-.707-.707l-.354.353zm0 3.292l-2-2l-.708.708l2 2zm-2-1.292l2-2l-.708-.708l-2 2zM4.5 8H11V7H4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.207 5.5L5.707 7H11v1H5.707l1.5 1.5l-.707.707L3.793 7.5L6.5 4.793z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2.707 8l3.147 3.146l-.708.707L.793 7.5l4.353-4.354l.708.708L2.707 7H14v1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.5.5zm0 4H0a.5.5 0 0 0 .854.354zm4-4l.354.354A.5.5 0 0 0 4.5 0zM2.146 2.854l12 12l.708-.708l-12-12zM0 .5v4h1v-4zm.854 4.354l4-4l-.708-.708l-4 4zM4.5 0h-4v1h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.146 9.146l-.353.354l.707.707l.354-.353zM10.5 7.5l.354.354l.353-.354l-.353-.354zM8.854 5.146L8.5 4.793l-.707.707l.353.354zm0 4.708l2-2l-.708-.708l-2 2zm2-2.708l-2-2l-.708.708l2 2zM10.5 7H4v1h6.5zm-3 7A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m8.5-2.707L11.207 7.5L8.5 10.207L7.793 9.5l1.5-1.5H4V7h5.293l-1.5-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m13.5 7.5l-4-4m4 4l-4 4m4-4H1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.146 9.146l-.353.354l.707.707l.354-.353zM10.5 7.5l.354.354l.353-.354l-.353-.354zM8.854 5.146L8.5 4.793l-.707.707l.353.354zm0 4.708l2-2l-.708-.708l-2 2zm2-2.708l-2-2l-.708.708l2 2zM10.5 7H4v1h6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 4.793L11.207 7.5L8.5 10.207L7.793 9.5l1.5-1.5H4V7h5.293l-1.5-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.854 3.146L14.207 7.5l-4.353 4.354l-.708-.708L12.293 8H1V7h11.293L9.146 3.854z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 0h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .854.354L2.5 3.207l11.646 11.647l.708-.708L3.207 2.5L4.854.854A.5.5 0 0 0 4.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.146 6.854l.354.353l.707-.707l-.353-.354zM7.5 4.5l.354-.354l-.354-.353l-.354.353zM5.146 6.146l-.353.354l.707.707l.354-.353zM14.5 7.5H14zm-7 7V14zm0-14V0zm-7 7H0zm9.354-1.354l-2-2l-.708.708l2 2zm-2.708-2l-2 2l.708.708l2-2zM7 4.5V11h1V4.5zm7 3A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zm-1 0A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 15a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15M4.793 6.5L7.5 3.793L10.207 6.5l-.707.707l-1.5-1.5V11H7V5.707l-1.5 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 1.5l.354-.354L7.5.793l-.354.353zm-.354.354l4 4l.708-.708l-4-4zm0-.708l-4 4l.708.708l4-4zM7 1.5V14h1V1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.146 6.146l-.353.354l.707.707l.354-.353zM7.5 4.5l.354-.354l-.354-.353l-.354.353zm1.646 2.354l.354.353l.707-.707l-.353-.354zm-3.292 0l2-2l-.708-.708l-2 2zm1.292-2l2 2l.708-.708l-2-2zM7 4.5V11h1V4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.793L10.207 6.5l-.707.707l-1.5-1.5V11H7V5.707l-1.5 1.5l-.707-.707z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.5.793l4.354 4.353l-.707.708L8 2.707V14H7V2.707L3.854 5.854l-.708-.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArtboardOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 4.5V4a.5.5 0 0 0-.5.5zm6 0h.5a.5.5 0 0 0-.5-.5zm0 6v.5a.5.5 0 0 0 .5-.5zm-6 0H4a.5.5 0 0 0 .5.5zM4 0v2h1V0zm6 0v2h1V0zM0 5h2V4H0zm0 6h2v-1H0zm13-6h2V4h-2zm0 6h2v-1h-2zm-9 2v2h1v-2zm6 0v2h1v-2zM4.5 5h6V4h-6zm5.5-.5v6h1v-6zm.5 5.5h-6v1h6zm-5.5.5v-6H4v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArtboardSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0v2h1V0zm6 0v2h1V0zM2 5H0V4h2zm-2 6h2v-1H0zm15-6h-2V4h2zm-2 6h2v-1h-2zm-9 4v-2h1v2zm6-2v2h1v-2zM4.5 4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.5 7.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm0 0v1a2 2 0 1 0 4 0v-1a7 7 0 1 0-3 5.745"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0v1a2.499 2.499 0 0 1-4.727 1.136A3.5 3.5 0 1 1 11 7.5v1a1.5 1.5 0 0 0 3 0v-1a6.5 6.5 0 1 0-2.786 5.335l.572.82A7.5 7.5 0 0 1 0 7.5m10 0a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttachOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m13.5 7.5l-5.757 5.757a4.243 4.243 0 0 1-6-6l5.929-5.929a2.828 2.828 0 0 1 4 4l-5.758 5.758a1.414 1.414 0 0 1-2-2L9.5 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttachSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.318.975a3.328 3.328 0 1 1 4.707 4.707l-5.757 5.757A1.914 1.914 0 1 1 3.56 8.732l5.585-5.586l.708.708l-5.586 5.585a.914.914 0 1 0 1.293 1.293l5.757-5.757a2.328 2.328 0 1 0-3.293-3.293L2.096 7.611a3.743 3.743 0 0 0 5.293 5.293l5.757-5.758l.708.708l-5.758 5.757A4.743 4.743 0 0 1 1.39 6.904z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttachmentOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 0v4.5a2 2 0 1 0 4 0v-3a1 1 0 0 0-2 0V5M6 .5h6.5a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1V8M11 4.5H7m4 3H7m4 3H4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttachmentSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 4.5V0h1v4.5a1.5 1.5 0 1 0 3 0v-3a.5.5 0 0 0-1 0V5H2V1.5a1.5 1.5 0 1 1 3 0v3a2.5 2.5 0 0 1-5 0"/><path fill="currentColor" fill-rule="evenodd" d="M12.5 0H6v4.5A3.5 3.5 0 0 1 2.5 8H1v5.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 12.5 0M11 4H7v1h4zm0 3H7v1h4zm-7 3h7v1H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioCableOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.5 1.5l-.354-.354A.5.5 0 0 0 6 1.5zm1-1l.354-.354a.5.5 0 0 0-.708 0zm1 1H9a.5.5 0 0 0-.146-.354zM6.5 9h2V8h-2zm2.5.5v3h1v-3zM8.5 13h-2v1h2zM6 12.5v-3H5v3zm.5.5a.5.5 0 0 1-.5-.5H5A1.5 1.5 0 0 0 6.5 14zm2.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM8.5 9a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 8.5 8zm-2-1A1.5 1.5 0 0 0 5 9.5h1a.5.5 0 0 1 .5-.5zm.5.5v-5H6v5zM6.5 4h2V3h-2zM8 3.5v5h1v-5zM7 13v2h1v-2zm0-9.5v-2H6v2zm-.146-1.646l1-1l-.708-.708l-1 1zm.292-1l1 1l.708-.708l-1-1zM8 1.5v2h1v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioCableSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.854.146a.5.5 0 0 0-.708 0l-1 1A.5.5 0 0 0 6 1.5V3h3V1.5a.5.5 0 0 0-.146-.354zM9 4v4H6V4zm1 5v3.5A1.5 1.5 0 0 1 8.5 14H8v1H7v-1h-.5A1.5 1.5 0 0 1 5 12.5V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioDocumentOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-3 4l.4-.3a.5.5 0 0 0-.9.3zm.3.4l.4-.3zm4.7 9.1h-10v1h10zM2 13.5v-12H1v12zm11-10v10h1v-10zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM6 11a1 1 0 0 1-1-1H4a2 2 0 0 0 2 2zm1-1a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2zM6 9a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2zm0-1a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1zm1-1.5V10h1V6.5zM8 7V4.5H7V7zm-.9-2.2l.3.4l.8-.6l-.3-.4zm.3.4A4.5 4.5 0 0 0 11 7V6a3.5 3.5 0 0 1-2.8-1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioDocumentSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 10a1 1 0 1 0-2 0a1 1 0 0 0 2 0"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zm6.342 2.526A.5.5 0 0 1 7.9 4.2l.3.4A3.5 3.5 0 0 0 11 6v1a4.5 4.5 0 0 1-3-1.146V10a2 2 0 1 1-1-1.732V4.5a.5.5 0 0 1 .342-.474" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AzureOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m3.5 4.5l-3 7h3l4-11zm11 9l-5-11l-2 5l3 4l-6 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AzureSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.97.67a.5.5 0 0 0-.824-.524l-4 4a.5.5 0 0 0-.106.157l-3 7A.5.5 0 0 0 .5 12h3a.5.5 0 0 0 .47-.33zm1.985 1.623a.5.5 0 0 0-.92.021l-2 5A.5.5 0 0 0 7.1 7.8l2.584 3.445l-5.342 1.78A.5.5 0 0 0 4.5 14h10a.5.5 0 0 0 .455-.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 12.5l-.39.312A.5.5 0 0 0 4.5 13zm-4-5l-.39-.312a.5.5 0 0 0 0 .624zm4-5V2a.5.5 0 0 0-.39.188zm9.5 1v8h1v-8zm-.5 8.5h-9v1h9zm-8.61.188l-4-5l-.78.624l4 5zm-4-4.376l4-5l-.78-.624l-4 5zM4.5 3h9V2h-9zm9.5 8.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zm1-8A1.5 1.5 0 0 0 13.5 2v1a.5.5 0 0 1 .5.5zM6.146 5.854l4 4l.708-.708l-4-4zm4-.708l-4 4l.708.708l4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.11 2.188A.5.5 0 0 1 4.5 2h9A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-9a.5.5 0 0 1-.39-.188l-4-5a.5.5 0 0 1 0-.624zm6.036 7.666L8.5 8.207L6.854 9.854l-.708-.708L7.793 7.5L6.146 5.854l.708-.708L8.5 6.793l1.646-1.647l.708.708L9.207 7.5l1.647 1.646z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.5 2v2.5a3 3 0 0 1-6 0V2m-3-.5v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h10A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zm9 3a2.5 2.5 0 0 1-5 0V2H4v2.5a3.5 3.5 0 1 0 7 0V2h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagMinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 4v-.5a3 3 0 0 1 6 0V4M5 9.5h5M2.401 6.39l-.778 7a1 1 0 0 0 .994 1.11h9.766a1 1 0 0 0 .994-1.11l-.778-7a1 1 0 0 0-.994-.89h-8.21a1 1 0 0 0-.994.89Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagMinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1A2.5 2.5 0 0 0 5 3.5V4H4v-.5a3.5 3.5 0 1 1 7 0V4h-1v-.5A2.5 2.5 0 0 0 7.5 1"/><path fill="currentColor" fill-rule="evenodd" d="M3.395 5a1.5 1.5 0 0 0-1.49 1.334l-.778 7A1.5 1.5 0 0 0 2.617 15h9.766a1.5 1.5 0 0 0 1.49-1.666l-.777-7A1.5 1.5 0 0 0 11.606 5zM5 9v1h5V9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.401 6.39l-.497-.056zm-.778 7l.497.055zm11.754 0l-.497.055zm-.778-7l.497-.056zM1.904 6.334l-.778 7l.994.11l.778-7zM2.617 15h9.766v-1H2.617zm11.257-1.666l-.778-7l-.994.11l.778 7zM11.604 5H3.396v1h8.21zm1.492 1.334A1.5 1.5 0 0 0 11.605 5v1a.5.5 0 0 1 .497.445zM12.383 15a1.5 1.5 0 0 0 1.49-1.666l-.993.11a.5.5 0 0 1-.497.556zM1.126 13.334A1.5 1.5 0 0 0 2.617 15v-1a.5.5 0 0 1-.497-.555zm1.772-6.89A.5.5 0 0 1 3.395 6V5a1.5 1.5 0 0 0-1.49 1.334zM5 4v-.5H4V4zm5-.5V4h1v-.5zM7.5 1A2.5 2.5 0 0 1 10 3.5h1A3.5 3.5 0 0 0 7.5 0zM5 3.5A2.5 2.5 0 0 1 7.5 1V0A3.5 3.5 0 0 0 4 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagPlusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 4v-.5a3 3 0 0 1 6 0V4m-3 3v5M5 9.5h5M2.401 6.39l-.778 7a1 1 0 0 0 .994 1.11h9.766a1 1 0 0 0 .994-1.11l-.778-7a1 1 0 0 0-.994-.89h-8.21a1 1 0 0 0-.994.89Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagPlusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3.5a2.5 2.5 0 0 1 5 0V4h1v-.5a3.5 3.5 0 1 0-7 0V4h1z"/><path fill="currentColor" fill-rule="evenodd" d="M1.904 6.334A1.5 1.5 0 0 1 3.395 5h8.21a1.5 1.5 0 0 1 1.49 1.334l.779 7A1.5 1.5 0 0 1 12.383 15H2.617a1.5 1.5 0 0 1-1.49-1.666zM7 9V7h1v2h2v1H8v2H7v-2H5V9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3.5a2.5 2.5 0 0 1 5 0V4h1v-.5a3.5 3.5 0 1 0-7 0V4h1zM1.904 6.334A1.5 1.5 0 0 1 3.395 5h8.21a1.5 1.5 0 0 1 1.49 1.334l.779 7A1.5 1.5 0 0 1 12.383 15H2.617a1.5 1.5 0 0 1-1.49-1.666z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.224-.447a.5.5 0 0 0-.448 0zM0 15h15v-1H0zM7.276.053l-6 3l.448.894l6-3zM0 6h15V5H0zm13.724-2.947l-6-3l-.448.894l6 3zM5 8v4h1V8zm4 0v4h1V8zM1 5.5v9h1v-9zm12 0v9h1v-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.724.053a.5.5 0 0 0-.448 0l-6 3l.448.894L7.5 1.06l5.776 2.888l.448-.894z"/><path fill="currentColor" fill-rule="evenodd" d="M14 6h1V5H0v1h1v8H0v1h15v-1h-1zm-9 6V8h1v4zm4 0V8h1v4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 14.5h15M5.5 12V6m4 6V3m4 9V0m-12 9v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 0h1v12h-1zm-3 3v9H9V3zM6 6v6H5V6zm-5 6V9h1v3zm14 3H0v-1h15z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarcodeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 2v11m5-11v11m2-11v11m7-11v11m-4-11v11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarcodeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 13V2h1v11zm5 0V2h1v11zm2 0V2h1v11zm3 0V2h1v11zm4 0V2h1v11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketMinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.449 14.398l.447-.224zm10.102 0l.447.224zM.703 6h13.594V5H.703zM14 5.703v.439h1v-.439zM12.386 14H2.614v1h9.772zM1 6.142v-.439H0v.439zm1.896 8.032A17.96 17.96 0 0 1 1 6.142H0c0 2.944.685 5.847 2.002 8.48zM2.614 14c.12 0 .229.068.282.174l-.894.448a.685.685 0 0 0 .612.378zm9.49.174a.315.315 0 0 1 .282-.174v1c.26 0 .496-.146.612-.378zM14 6.142c0 2.788-.65 5.538-1.896 8.032l.894.448A18.96 18.96 0 0 0 15 6.142zM14.297 6A.297.297 0 0 1 14 5.703h1A.703.703 0 0 0 14.297 5zM.703 5A.703.703 0 0 0 0 5.703h1A.297.297 0 0 1 .703 6zm3.226.757l3-5L6.07.243l-3 5zm4.142-5l3 5l.858-.514l-3-5zM5 10h5V9H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketMinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.383 5L6.93.757L6.07.243L3.217 5H.703A.703.703 0 0 0 0 5.703v.439c0 2.944.685 5.847 2.002 8.48a.685.685 0 0 0 .612.378h9.772c.26 0 .496-.146.612-.379A18.96 18.96 0 0 0 15 6.141v-.438A.703.703 0 0 0 14.297 5h-2.514L8.93.243l-.86.514L10.617 5zM5 10h5V9H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.449 14.398l.447-.224zm10.102 0l.447.224zM.703 6h13.594V5H.703zM14 5.703v.439h1v-.439zM12.386 14H2.614v1h9.772zM1 6.142v-.439H0v.439zm1.896 8.032A17.96 17.96 0 0 1 1 6.142H0c0 2.944.685 5.847 2.002 8.48zM2.614 14c.12 0 .229.068.282.174l-.894.448a.685.685 0 0 0 .612.378zm9.49.174a.315.315 0 0 1 .282-.174v1c.26 0 .496-.146.612-.378zM14 6.142c0 2.788-.65 5.538-1.896 8.032l.894.448A18.96 18.96 0 0 0 15 6.142zM14.297 6A.297.297 0 0 1 14 5.703h1A.703.703 0 0 0 14.297 5zM.703 5A.703.703 0 0 0 0 5.703h1A.297.297 0 0 1 .703 6zm3.226.757l3-5L6.07.243l-3 5zm4.142-5l3 5l.858-.514l-3-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketPlusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.449 14.398l.447-.224zm10.102 0l.447.224zM.703 6h13.594V5H.703zM14 5.703v.439h1v-.439zM12.386 14H2.614v1h9.772zM1 6.142v-.439H0v.439zm1.896 8.032A17.96 17.96 0 0 1 1 6.142H0c0 2.944.685 5.847 2.002 8.48zM2.614 14c.12 0 .229.068.282.174l-.894.448a.685.685 0 0 0 .612.378zm9.49.174a.315.315 0 0 1 .282-.174v1c.26 0 .496-.146.612-.378zM14 6.142c0 2.788-.65 5.538-1.896 8.032l.894.448A18.96 18.96 0 0 0 15 6.142zM14.297 6A.297.297 0 0 1 14 5.703h1A.703.703 0 0 0 14.297 5zM.703 5A.703.703 0 0 0 0 5.703h1A.297.297 0 0 1 .703 6zm3.226.757l3-5L6.07.243l-3 5zm4.142-5l3 5l.858-.514l-3-5zM7 7v5h1V7zm-2 3h5V9H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketPlusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.383 5L6.93.757L6.07.243L3.217 5H.703A.703.703 0 0 0 0 5.703v.439c0 2.944.685 5.847 2.002 8.48a.685.685 0 0 0 .612.378h9.772c.26 0 .496-.146.612-.379A18.96 18.96 0 0 0 15 6.141v-.438A.703.703 0 0 0 14.297 5h-2.514L8.93.243l-.86.514L10.617 5zM7 12v-2H5V9h2V7h1v2h2v1H8v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.929.757L4.383 5h6.234L8.07.757l.86-.514L11.783 5h2.514c.388 0 .703.315.703.703v.439a18.96 18.96 0 0 1-2.002 8.48a.684.684 0 0 1-.612.378H2.614a.685.685 0 0 1-.612-.379A18.96 18.96 0 0 1 0 6.141v-.438C0 5.315.315 5 .703 5h2.514L6.07.243z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BathOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 7.5h15m-10.5 5h6m-6 0a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3h2V2m-2 10.5V15m6-2.5a3 3 0 0 0 3-3v-2m-3 5V15M5 3.5h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BathSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3.5A2.5 2.5 0 0 1 4.5 1H6v1h1V0H4.5A3.5 3.5 0 0 0 1 3.5V7H0v1h1v1.5a3.5 3.5 0 0 0 3 3.465V15h1v-2h5v2h1v-2.035A3.501 3.501 0 0 0 14 9.5V8h1V7H2z"/><path fill="currentColor" d="M8 4H5V3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryChargeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 11.5H12zm-1 1v.5zm0-10V2zm1 1h.5zm-12 0H0zm1-1V3zm-1 9H1zm1 1V12zm5-3l-.224.447A.5.5 0 0 0 7 9.5zm0-4l.224-.447A.5.5 0 0 0 6 5.5zm-5.5 6v-8H0v8zM1.5 3h10V2h-10zm10.5.5v8h1v-8zm-.5 8.5h-10v1h10zm.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM11.5 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 2zM1 3.5a.5.5 0 0 1 .5-.5V2A1.5 1.5 0 0 0 0 3.5zm-1 8A1.5 1.5 0 0 0 1.5 13v-1a.5.5 0 0 1-.5-.5zM15 10V5h-1v5zM2.276 7.947l4 2l.448-.894l-4-2zM7 9.5v-4H6v4zm-.724-3.553l4 2l.448-.894l-4-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryChargeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5zm6.724-6.447A.5.5 0 0 0 6 5.5v3.191L2.724 7.053l-.448.894l4 2A.5.5 0 0 0 7 9.5V6.309l3.276 1.638l.448-.894z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFiveOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m2 7V4m2 7V4m2 7V4m2 7V4m2 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFiveSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5zm3-.5V4H2v7zm2 0V4H4v7zm2-7v7H6V4zm2 7V4H8v7zm2-7v7h-1V4z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFourOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m2 7V4m2 7V4m2 7V4m4 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFourSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5zM3 4v7H2V4zm2 0v7H4V4zm2 7V4H6v7zm2-7v7H8V4z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryOneOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m10 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryOneSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5zm3-.5V4H2v7z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryThreeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m2 7V4m2 7V4m6 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryThreeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5zm3-.5V4H2v7zm2 0V4H4v7zm2-7v7H6V4z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryTwoOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 10V5m-12 6V4m2 7V4m8 7.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryTwoSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 11.5A1.5 1.5 0 0 0 1.5 13h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 11.5 2h-10A1.5 1.5 0 0 0 0 3.5zM3 4v7H2V4zm2 0v7H4V4z" clip-rule="evenodd"/><path fill="currentColor" d="M15 5v5h-1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryZeroOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 10V5m-2 6.5v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryZeroSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 13A1.5 1.5 0 0 1 0 11.5v-8A1.5 1.5 0 0 1 1.5 2h10A1.5 1.5 0 0 1 13 3.5v8a1.5 1.5 0 0 1-1.5 1.5zM15 10V5h-1v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BedDoubleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 7v8M.5 7v8M0 10.5h15m-15-3h15m-13-2h4m3 0h4M.5 6V.5h14V6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BedDoubleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h15v6h-1V1H1v5H0z"/><path fill="currentColor" d="M6 6H2V5h4zm-6 9h1v-4h13v4h1V7H0zm9-9h4V5H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BedSingleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 4v11M.5 0v15M0 10.5h15m-15-3h15m-13-2h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BedSingleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 15h1v-4h13v4h1V4h-1v3H1V0H0z"/><path fill="currentColor" d="M6 6H2V5h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 9.5v.5h.5v-.5zm-14-7V2H0v.5zm0 10H0v.5h.5zm7.5-3a3.5 3.5 0 0 0 3.5 3.5v-1A2.5 2.5 0 0 1 9 9.5zM11.5 6A3.5 3.5 0 0 0 8 9.5h1A2.5 2.5 0 0 1 11.5 7zM15 9.5A3.5 3.5 0 0 0 11.5 6v1A2.5 2.5 0 0 1 14 9.5zm-1.473 1.464A2.496 2.496 0 0 1 11.5 12v1a3.496 3.496 0 0 0 2.837-1.45zM.5 3H3V2H.5zM3 7H.5v1H3zm-2 .5v-5H0v5zM5 5a2 2 0 0 1-2 2v1a3 3 0 0 0 3-3zM3 3a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3zM.5 8H4V7H.5zM4 12H.5v1H4zm-3 .5v-5H0v5zM6 10a2 2 0 0 1-2 2v1a3 3 0 0 0 3-3zM4 8a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3zm4.5 2h6V9h-6zM9 4h5V3H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2h3a3 3 0 0 1 2.051 5.19A3.001 3.001 0 0 1 4 13H0zm1 6v4h3a2 2 0 1 0 0-4zm0-1h2a2 2 0 1 0 0-4H1zm13-3H9V3h5zM8 9.5a3.5 3.5 0 1 1 7 0v.5H9.05a2.5 2.5 0 0 0 4.477.964l.81.586A3.5 3.5 0 0 1 8 9.5M9.05 9h4.9a2.5 2.5 0 0 0-4.9 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 10.5h13m-11.5 0v-5a5 5 0 0 1 10 0v5m-7 1.5v.5a2 2 0 1 0 4 0V12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0A5.5 5.5 0 0 0 2 5.5V10H1v1h13v-1h-1V5.5A5.5 5.5 0 0 0 7.5 0M5 12.5V12h5v.5a2.5 2.5 0 0 1-5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 3V1.5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1V3M0 3.5h15m-13.5 0v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-10M7.5 7v5m-3-3v3m6-3v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 3V1.5A1.5 1.5 0 0 0 9.5 0h-4A1.5 1.5 0 0 0 4 1.5V3H0v1h1v9.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V4h1V3zM5 1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3H5zM7 7v5h1V7zm-3 5V9h1v3zm6-3v3h1V9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitbucketOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.495.57zm14 0l.495.07A.5.5 0 0 0 14.5 0zm-2 14v.5a.5.5 0 0 0 .495-.43zm-10 0l-.495.07A.5.5 0 0 0 2.5 15zM5 4.5V4a.5.5 0 0 0-.498.542zm4.5 6v.5a.5.5 0 0 0 .498-.459zm-4 0l-.498.041A.5.5 0 0 0 5.5 11zM.5 1h14V0H.5zM14.005.43l-2 14l.99.14l2-14zM12.5 14h-10v1h10zm-9.505.43l-2-14l-.99.14l2 14zM5 5h5V4H5zm4.502-.542l-.5 6l.996.083l.5-6zM9.5 10h-4v1h4zm-3.502.459l-.5-6l-.996.083l.5 6zM10 5h4V4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitbucketSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 0a.5.5 0 0 0-.495.57l2 14A.5.5 0 0 0 2.5 15h10a.5.5 0 0 0 .495-.43L14.219 6H9.373l-.333 4H5.96l-.417-5h8.82l.632-4.43A.5.5 0 0 0 14.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 1.5h5a3 3 0 1 1 0 6h-5m0-6v6m0-6H2m1.5 0V0m0 7.5h6a3 3 0 1 1 0 6h-6m0-6v6m0-6H2m1.5 6H2m1.5 0V15m4-15v1.5m0 12V15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1V0h1v1h3V0h1v1h.5a3.5 3.5 0 0 1 2.21 6.215A3.501 3.501 0 0 1 9.5 14H8v1H7v-1H4v1H3v-1H2v-1h1V8H2V7h1V2H2V1zm1 1v5h4.5a2.5 2.5 0 0 0 0-5zm0 6h5.5a2.5 2.5 0 0 1 0 5H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 14.5H7a.5.5 0 0 0 .787.41zm0-14l.287-.41A.5.5 0 0 0 7 .5zm5 3.5l.282.413a.5.5 0 0 0 .005-.823zm0 7l.287.41a.5.5 0 0 0-.005-.823zM8 14.5V7.41H7v7.09zm0-7.09V.5H7v6.91zM7.213.91l5 3.5l.574-.82l-5-3.5zm5.005 2.677l-5 3.409l.564.826l5-3.409zm-5 3.409l-6 4.09l.564.827l6-4.09zm.569 7.914l5-3.5l-.574-.82l-5 3.5zm4.995-4.323l-11-7.5l-.564.826l11 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.27.057a.5.5 0 0 1 .517.033l5 3.5a.5.5 0 0 1-.005.823L8.254 7.5l4.528 3.087a.5.5 0 0 1 .005.823l-5 3.5A.5.5 0 0 1 7 14.5V8.355l-5.218 3.558l-.564-.826L6.48 7.5L1.22 3.913l.563-.826L7 6.645V.5a.5.5 0 0 1 .27-.443M8 8.537l3.62 2.468L8 13.54zm0-2.074V1.46l3.62 2.535z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 7.5h5a3 3 0 1 0 0-6H3.58a.08.08 0 0 0-.08.08zm0 0h6a3 3 0 1 1 0 6H3.59a.09.09 0 0 1-.09-.09z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoldSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1.58c0-.32.26-.58.58-.58H8.5a3.5 3.5 0 0 1 2.21 6.215A3.501 3.501 0 0 1 9.5 14H3.59a.59.59 0 0 1-.59-.59zM4 8v5h5.5a2.5 2.5 0 0 0 0-5zm0-1h4.5a2.5 2.5 0 0 0 0-5H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5.5V0a.5.5 0 0 0-.5.5zm0 13H1a.5.5 0 0 0 .5.5zM4 0v15h1V0zM1.5 1h10V0h-10zM13 2.5v9h1v-9zM11.5 13h-10v1h10zm-9.5.5V.5H1v13zm11-2a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5zM11.5 1A1.5 1.5 0 0 1 13 2.5h1A2.5 2.5 0 0 0 11.5 0zM7 5h4V4H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0a.5.5 0 0 0-.5.5v13a.5.5 0 0 0 .5.5H3V0z"/><path fill="currentColor" fill-rule="evenodd" d="M4 15h1v-1h6.5a2.5 2.5 0 0 0 2.5-2.5v-9A2.5 2.5 0 0 0 11.5 0H4zm7-10H7V4h4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.5 14.5l-.312.39A.5.5 0 0 0 13 14.5zm0-14h.5V0h-.5zm-10 0V0H2v.5zm0 14H2a.5.5 0 0 0 .812.39zm5-4l.312-.39a.5.5 0 0 0-.624 0zm5.5 4V.5h-1v14zM2 .5v14h1V.5zm.812 14.39l5-4l-.624-.78l-5 4zm4.376-4l5 4l.624-.78l-5-4zM12.5 0h-10v1h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0H2v14.5a.5.5 0 0 0 .812.39L7.5 11.14l4.688 3.75A.5.5 0 0 0 13 14.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderAllOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 1.5v12m-6-6h12m-12-6h12v12h-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderAllSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v13H1zm1 1v5h5V2zm6 0v5h5V2zm5 6H8v5h5zm-6 5V8H2v5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottomOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 1.5h1m5 0h1m2 0h1m2 0h1m-1 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m5 3h1m-7 0h1m-7 0h1m-1-3h1m2 0h1m-4-3h1m2-3h1m-4 12h13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottomSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1zm3 0H4V1h1zm3 0H7V1h1zm3 0h-1V1h1zm3 0h-1V1h1zM2 5H1V4h1zm6 0H7V4h1zm6 0h-1V4h1zM2 8H1V7h1zm3 0H4V7h1zm3 0H7V7h1zm3 0h-1V7h1zm3 0h-1V7h1zM2 11H1v-1h1zm6 0H7v-1h1zm6 0h-1v-1h1zm0 3H1v-1h13z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderHorizontalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 1.5h1m5 0h1m2 0h1m2 0h1m-1 3h1m-7 0h1m5 6h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1m-1-6h1m2-3h1m-4 6h13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderHorizontalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1zm3 0H4V1h1zm3 0H7V1h1zm3 0h-1V1h1zm3 0h-1V1h1zM2 5H1V4h1zm6 0H7V4h1zm6 0h-1V4h1zm0 3H1V7h13zM2 11H1v-1h1zm6 0H7v-1h1zm6 0h-1v-1h1zM2 14H1v-1h1zm3 0H4v-1h1zm3 0H7v-1h1zm3 0h-1v-1h1zm3 0h-1v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInnerOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 1.5h1m8 0h1m2 0h1m-1 3h1m-1 6h1m-1 3h1m-4 0h1m-7 0h1m-4 0h1m-1-3h1m-1-6h1m2-3h1m-4 6h13M7.5 1v13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInnerSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1zm3 0H4V1h1zm2 5H1v1h6v6h1V8h6V7H8V1H7zm4-5h-1V1h1zm3 0h-1V1h1zM2 5H1V4h1zm12 0h-1V4h1zM2 11H1v-1h1zm12 0h-1v-1h1zM2 14H1v-1h1zm3 0H4v-1h1zm6 0h-1v-1h1zm3 0h-1v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7 1.5h1m2 0h1m2 0h1m-1 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m5 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m-4 0h1m-1-6h1m-1-6h1M1.5 1v13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 14V1h1v13zM5 2H4V1h1zm3 0H7V1h1zm3 0h-1V1h1zm3 0h-1V1h1zM8 5H7V4h1zm6 0h-1V4h1zM5 8H4V7h1zm3 0H7V7h1zm3 0h-1V7h1zm3 0h-1V7h1zm-6 3H7v-1h1zm6 0h-1v-1h1zm-9 3H4v-1h1zm3 0H7v-1h1zm3 0h-1v-1h1zm3 0h-1v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderNoneOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 1.5h1m5 0h1m2 0h1m2 0h1m-1 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m5 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1m-1-3h1m2 0h1m-4-3h1m2-3h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderNoneSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1zm3 0H4V1h1zm3 0H7V1h1zm3 0h-1V1h1zm3 0h-1V1h1zM2 5H1V4h1zm6 0H7V4h1zm6 0h-1V4h1zM2 8H1V7h1zm3 0H4V7h1zm3 0H7V7h1zm3 0h-1V7h1zm3 0h-1V7h1zM2 11H1v-1h1zm6 0H7v-1h1zm6 0h-1v-1h1zM2 14H1v-1h1zm3 0H4v-1h1zm3 0H7v-1h1zm3 0h-1v-1h1zm3 0h-1v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOuterOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7 4.5h1m2 3h1m-4 0h1m-1 3h1m-4-3h1m-3.5-6h12v12h-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOuterSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v13H1zm1 1v11h11V2zm6 3H7V4h1zM5 8H4V7h1zm3 0H7V7h1zm3 0h-1V7h1zm-3 3H7v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRadiusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 1.5h1m2 0h1m-1 3h1m-1 3h1m-1 3h1m-1 3h1m-4 0h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1M1.5 8V5.5a4 4 0 0 1 4-4H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRadiusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 2A3.5 3.5 0 0 0 2 5.5V8H1V5.5A4.5 4.5 0 0 1 5.5 1H8v1zM11 2h-1V1h1zm3 0h-1V1h1zm0 3h-1V4h1zm0 3h-1V7h1zM2 11H1v-1h1zm12 0h-1v-1h1zM2 14H1v-1h1zm3 0H4v-1h1zm3 0H7v-1h1zm3 0h-1v-1h1zm3 0h-1v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 1.5h1m5 0h1m2 0h1m-4 3h1m2 3h1m-4 0h1m-1 3h1m2 3h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1m-1-3h1m2 0h1m-4-3h1m2-3h1m8.5-.5v13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1zm3 0H4V1h1zm3 0H7V1h1zm3 0h-1V1h1zm2 12V1h1v13zM2 5H1V4h1zm6 0H7V4h1zM2 8H1V7h1zm3 0H4V7h1zm3 0H7V7h1zm3 0h-1V7h1zm-9 3H1v-1h1zm6 0H7v-1h1zm-6 3H1v-1h1zm3 0H4v-1h1zm3 0H7v-1h1zm3 0h-1v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTopOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13 4.5h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m5 3h1m-7 0h1m5 3h1m-4 0h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1m-1-3h1m2 0h1m-4-3h1m-1-3h13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTopSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 2H1V1h13zM2 5H1V4h1zm6 0H7V4h1zm6 0h-1V4h1zM2 8H1V7h1zm3 0H4V7h1zm3 0H7V7h1zm3 0h-1V7h1zm3 0h-1V7h1zM2 11H1v-1h1zm6 0H7v-1h1zm6 0h-1v-1h1zM2 14H1v-1h1zm3 0H4v-1h1zm3 0H7v-1h1zm3 0h-1v-1h1zm3 0h-1v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderVerticalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 1.5h1m8 0h1m2 0h1m-1 3h1m-1 3h1m-4 0h1m2 3h1m-1 3h1m-4 0h1m-7 0h1m-4 0h1m-1-3h1m-1-3h1m2 0h1m-4-3h1m2-3h1M7.5 1v13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderVerticalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1zm3 0H4V1h1zm2 12V1h1v13zm4-12h-1V1h1zm3 0h-1V1h1zM2 5H1V4h1zm12 0h-1V4h1zM2 8H1V7h1zm3 0H4V7h1zm6 0h-1V7h1zm3 0h-1V7h1zM2 11H1v-1h1zm12 0h-1v-1h1zM2 14H1v-1h1zm3 0H4v-1h1zm6 0h-1v-1h1zm3 0h-1v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BottomLeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 13.5H1a.5.5 0 0 0 .5.5zm0 .5H7v-1H1.5zm.5-.5V8H1v5.5zm-.146.354l12-12l-.708-.708l-12 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BottomLeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.854 1.854L2.707 13H7v1H1V8h1v4.293L13.146 1.146z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BottomRightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 13.5v.5a.5.5 0 0 0 .5-.5zm0-.5H8v1h5.5zm.5.5V8h-1v5.5zm-.146-.354l-12-12l-.708.708l12 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BottomRightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.854 1.146L13 12.293V8h1v6H8v-1h4.293L1.146 1.854z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M.5 3.498L7.5.5l7 2.998m-14 0l7 2.998m-7-2.998V3.5m14-.002l-7 2.998m7-2.998V11.5l-7 3m7-11.002L7.5 6.5v8m0-8.004V14.5m0-8.004L.5 3.5m7 11l-7-3v-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.303.04a.5.5 0 0 1 .394 0L14.5 2.956l-7 3l-7-3zM0 3.83v7.67c0 .2.12.38.303.46L7 14.83v-8zm8 3l7-3v7.67a.5.5 0 0 1-.303.46L8 14.83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BracketOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 7.5h2m13 0h-2m-8 7H3.5a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2H5m5 14h1.5a2 2 0 0 0 2-2v-10a2 2 0 0 0-2-2H10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BracketSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 1A1.5 1.5 0 0 0 2 2.5v10A1.5 1.5 0 0 0 3.5 14H5v1H3.5A2.5 2.5 0 0 1 1 12.5V8H0V7h1V2.5A2.5 2.5 0 0 1 3.5 0H5v1zM10 0h1.5A2.5 2.5 0 0 1 14 2.5V7h1v1h-1v4.5a2.5 2.5 0 0 1-2.5 2.5H10v-1h1.5a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 11.5 1H10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 3v-.5a2 2 0 1 1 4 0V3m-9 3.5c3.706 4.235 10.294 4.235 14 0m-13-3h12a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 2.5V3H1.5A1.5 1.5 0 0 0 0 4.5v9A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-9A1.5 1.5 0 0 0 13.5 3H10v-.5a2.5 2.5 0 0 0-5 0M7.5 1A1.5 1.5 0 0 0 6 2.5V3h3v-.5A1.5 1.5 0 0 0 7.5 1M.66 7.367a10.083 10.083 0 0 0 13.68 0l-.68-.734a9.083 9.083 0 0 1-12.32 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 3v-.5a2 2 0 1 1 4 0V3m-9 8.5h14m-13-8h12a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 2.5V3H1.5A1.5 1.5 0 0 0 0 4.5V11h15V4.5A1.5 1.5 0 0 0 13.5 3H10v-.5a2.5 2.5 0 0 0-5 0M7.5 1A1.5 1.5 0 0 0 6 2.5V3h3v-.5A1.5 1.5 0 0 0 7.5 1" clip-rule="evenodd"/><path fill="currentColor" d="M15 12H0v1.5A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrushOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.5 7.5v-5a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v5m12 0h-12m12 0v2a2 2 0 0 1-2 2h-3v2a1 1 0 1 1-2 0v-2h-3a2 2 0 0 1-2-2v-2m4-7V5m2-4.5V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrushSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 2.5A2.5 2.5 0 0 1 3.5 0H5v5h1V0h1v3h1V0h3.5A2.5 2.5 0 0 1 14 2.5v7a2.5 2.5 0 0 1-2.5 2.5H9v1.5a1.5 1.5 0 0 1-3 0V12H3.5A2.5 2.5 0 0 1 1 9.5zM2 8v1.5A1.5 1.5 0 0 0 3.5 11H7v2.5a.5.5 0 0 0 1 0V11h3.5A1.5 1.5 0 0 0 13 9.5V8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 4.5h6m-6 0l-.367.733A6 6 0 0 0 3.5 7.916V10.5a4 4 0 0 0 8 0V7.916a6 6 0 0 0-.633-2.683L10.5 4.5m-6 0v-1a3 3 0 0 1 6 0v1M0 8.5h3.5m11.5 0h-3.5M1 14l3-1.5M14 14l-3-1.5M1 3l3 1.5M14 3l-3 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.941V3.5a3.5 3.5 0 1 1 7 0v.441l2.776-1.388l.448.894l-2.953 1.477l.043.086A6.5 6.5 0 0 1 12 7.916V8h3v1h-3v1.5c0 .625-.127 1.22-.358 1.762l2.582 1.29l-.448.895l-2.627-1.313A4.494 4.494 0 0 1 7.5 15a4.494 4.494 0 0 1-3.65-1.866l-2.626 1.313l-.448-.894l2.582-1.291A4.486 4.486 0 0 1 3 10.5V9H0V8h3v-.084a6.5 6.5 0 0 1 .686-2.906l.043-.086L.776 3.447l.448-.894zM5 3.5a2.5 2.5 0 0 1 5 0V4H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.224-.447a.5.5 0 0 0-.448 0zm-3 8V8H4v.5zm4 0H9V8h-.5zM0 15h15v-1H0zM7.276.053l-6 3l.448.894l6-3zM0 6h15V5H0zm13.724-2.947l-6-3l-.448.894l6 3zM1 5.5v9h1v-9zm12 0v9h1v-9zm-8 9v-6H4v6zM4.5 9h4V8h-4zM8 8.5v6h1v-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.724.053a.5.5 0 0 0-.448 0l-6 3l.448.894L7.5 1.06l5.776 2.888l.448-.894zM14 6h1V5H0v1h1v8H0v1h4V8h5v7h6v-1h-1z"/><path fill="currentColor" d="M8 15V9H5v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulbOffOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.076 6.47L3.58 6.4zm-.01.07l.495.07zm6.858-.07l-.495.07zm.01.07l.495-.07zM9.5 12.5v.5a.5.5 0 0 0 .5-.5zm-4 0H5a.5.5 0 0 0 .5.5zm-.745-3.347l.396-.306zm5.49 0l-.396-.306zM6 15h3v-1H6zM3.58 6.4l-.01.07l.99.14l.01-.07zM7.5 3a3.959 3.959 0 0 0-3.92 3.4l.99.14A2.959 2.959 0 0 1 7.5 4zm3.92 3.4A3.959 3.959 0 0 0 7.5 3v1a2.96 2.96 0 0 1 2.93 2.54zm.01.07l-.01-.07l-.99.14l.01.07zm-.79 2.989c.63-.814.948-1.875.79-2.99l-.99.142a2.951 2.951 0 0 1-.59 2.236zM9 10.9v1.6h1v-1.599zm.5 1.1h-4v1h4zm-3.5.5v-1.599H5V12.5zM3.57 6.47a3.951 3.951 0 0 0 .79 2.989l.79-.612a2.951 2.951 0 0 1-.59-2.236zM6 10.9c0-.823-.438-1.523-.85-2.054l-.79.612c.383.495.64.968.64 1.442zm3.85-2.054C9.437 9.378 9 10.077 9 10.9h1c0-.474.257-.947.64-1.442z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulbOffSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 3a3.959 3.959 0 0 0-3.92 3.4l-.01.07a3.951 3.951 0 0 0 .79 2.989c.383.495.64.968.64 1.442V12.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-1.599c0-.474.256-.947.64-1.442c.63-.814.948-1.875.79-2.99l-.01-.07A3.959 3.959 0 0 0 7.5 3M6 15h3v-1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulbOnOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.076 6.47l.495.07zm-.01.07l-.495-.07zm6.858-.07l.495-.07zm.01.07l-.495.07zM9.5 12.5v.5a.5.5 0 0 0 .5-.5zm-4 0H5a.5.5 0 0 0 .5.5zm-.745-3.347l.396-.306zm5.49 0l-.396-.306zM6 15h3v-1H6zM3.58 6.4l-.01.07l.99.14l.01-.07zM7.5 3a3.959 3.959 0 0 0-3.92 3.4l.99.14A2.959 2.959 0 0 1 7.5 4zm3.92 3.4A3.959 3.959 0 0 0 7.5 3v1a2.96 2.96 0 0 1 2.93 2.54zm.01.07l-.01-.07l-.99.14l.01.07zm-.79 2.989c.63-.814.948-1.875.79-2.99l-.99.142a2.951 2.951 0 0 1-.59 2.236zM9 10.9v1.6h1v-1.599zm.5 1.1h-4v1h4zm-3.5.5v-1.599H5V12.5zM3.57 6.47a3.951 3.951 0 0 0 .79 2.989l.79-.612a2.951 2.951 0 0 1-.59-2.236zM6 10.9c0-.823-.438-1.523-.85-2.054l-.79.612c.383.495.64.968.64 1.442zm3.85-2.054C9.437 9.378 9 10.077 9 10.9h1c0-.474.257-.947.64-1.442zM7 0v2h1V0zM0 8h2V7H0zm13 0h2V7h-2zM3.354 3.646l-1.5-1.5l-.708.708l1.5 1.5zm9 .708l1.5-1.5l-.708-.708l-1.5 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulbOnSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0v2h1V0zM3.354 3.646l-1.5-1.5l-.708.708l1.5 1.5zm9 .708l1.5-1.5l-.708-.708l-1.5 1.5zM7.5 3a3.959 3.959 0 0 0-3.92 3.4l-.01.07a3.951 3.951 0 0 0 .79 2.989c.383.495.64.968.64 1.442V12.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-1.599c0-.474.257-.947.64-1.442c.63-.814.948-1.875.79-2.99l-.01-.07A3.959 3.959 0 0 0 7.5 3M0 8h2V7H0zm13 0h2V7h-2zm-7 7h3v-1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 10V8.5m0 0v-5a1 1 0 0 1 2 0v4h3.353c.91 0 1.647.737 1.647 1.647V10A4.5 4.5 0 0 1 8 14.5h-.5a4 4 0 0 1-4-4a2 2 0 0 1 2-2Zm3.5-3h2a2.5 2.5 0 0 0 0-5H4a2.5 2.5 0 0 0 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6a3 3 0 0 1 0-6h7a3 3 0 1 1 0 6H9V3.5a2.5 2.5 0 0 0-5 0z"/><path fill="currentColor" d="M6.5 2A1.5 1.5 0 0 0 5 3.5v4.55a2.5 2.5 0 0 0-2 2.45A4.5 4.5 0 0 0 7.5 15H8a5 5 0 0 0 5-5v-.853A2.147 2.147 0 0 0 10.853 7H8V3.5A1.5 1.5 0 0 0 6.5 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m10 5.5l-.068-.068a3.182 3.182 0 0 0-2.25-.932H7.5a3 3 0 0 0 0 6h.182c.844 0 1.653-.335 2.25-.932L10 9.5m-8.5 1v-6l6-3.5l6 3.5v6l-6 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CsharpOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 5v5m2-5v5M8 6.5h5m-5 2h5M8 10l-.402.201A2.831 2.831 0 0 1 3.5 7.668v-.336a2.832 2.832 0 0 1 4.098-2.533L8 5m-6.5 5.5v-6l6-3.5l6 3.5v6l-6 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CsharpSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8V7h1v1z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5.421L14 4.213v6.574L7.5 14.58L1 10.787V4.213zM6.332 4A3.332 3.332 0 0 0 3 7.332v.336a3.332 3.332 0 0 0 4.821 2.98l.403-.2l-.448-.895l-.402.2A2.331 2.331 0 0 1 4 7.669v-.336a2.332 2.332 0 0 1 3.374-2.086l.402.201l.448-.894l-.403-.201A3.332 3.332 0 0 0 6.331 4M9 5v1H8v1h1v1H8v1h1v1h1V9h1v1h1V9h1V8h-1V7h1V6h-1V5h-1v1h-1V5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Csolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.421L14 4.213v6.574L7.5 14.58L1 10.787V4.213zM7.5 4a3.5 3.5 0 1 0 0 7h.182c.976 0 1.913-.388 2.604-1.078l.068-.068l-.708-.708l-.068.068A2.682 2.682 0 0 1 7.682 10H7.5a2.5 2.5 0 0 1 0-5h.182c.711 0 1.393.283 1.896.786l.068.068l.708-.708l-.068-.068A3.682 3.682 0 0 0 7.682 4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 4.5h7m-7 4h1m2 0h1m2 0h1m-7 3h1m2 0h1m2 0h1m-8.5 3h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h10A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM4 5h7V4H4zm0 4h1V8H4zm4 0H7V8h1zm2 0h1V8h-1zm-5 3H4v-1h1zm2 0h1v-1H7zm4 0h-1v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5M5 8.5h5m-8.5-6h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2M5 9h5V8H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarNoAccessOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5m-2 1.5l-4 4m-4-8h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Zm6 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarNoAccessSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 6a2.5 2.5 0 0 0-2.086 3.879L8.88 6.414A2.488 2.488 0 0 0 7.5 6m0 5c-.51 0-.983-.152-1.379-.414L9.586 7.12A2.5 2.5 0 0 1 7.5 11"/><path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2M4 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5m-10-2.5h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5m-4 1v5M5 8.5h5m-8.5-6h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2M7 11V9H5V8h2V6h1v2h2v1H8v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h1.5A1.5 1.5 0 0 1 15 3.5v10a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5v-10A1.5 1.5 0 0 1 1.5 2H3V0h1v2h7V0h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTickOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5M5 8.5l2 2l3.5-4m-9-4h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTickSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2m-6.476 9.232l3.852-4.403l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarXoutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5m-6 1.5l4 4m-4 0l4-4m-8-4h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarXsolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12V0h-1v2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v10A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-10A1.5 1.5 0 0 0 13.5 2m-4.354 8.854L7.5 9.207l-1.646 1.646l-.708-.707L6.793 8.5L5.146 6.854l.708-.708L7.5 7.793l1.646-1.647l.708.708L8.207 8.5l1.647 1.646z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7 1.5H2m12.5 11v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1Zm-5-2a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1h5v1H2zm6 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M0 12.5A1.5 1.5 0 0 0 1.5 14h12a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 3h-12A1.5 1.5 0 0 0 0 4.5zM9.5 6a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CandleChartOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 0v14.5H15M8.5 0v3.5m-5 6V12m0-8v1.5m10-1.5v2.5m0 4V13m-11-7.5h2v4h-2zm5-2h2v4h-2zm5 3h2v4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CandleChartSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0H0v15h15v-1H1z"/><path fill="currentColor" d="M8 0v3H7v5h3V3H9V0zM3 4v1H2v5h1v2h1v-2h1V5H4V4zm9 2h1V4h1v2h1v5h-1v2h-1v-2h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 6.497h.5v-.139l-.071-.119zm-14 0l-.429-.258L0 6.36v.138zm2.126-3.541l-.429-.258zm9.748 0l.429-.258zM3.5 11.5V11H3v.5zm8 0h.5V11h-.5zM14 6.497V12.5h1V6.497zM.929 6.754l2.126-3.54l-.858-.516L.071 6.24zM5.198 2h4.604V1H5.198zm6.747 1.213l2.126 3.541l.858-.515l-2.126-3.54zM2.5 13h-1v1h1zm.5-1.5v1h1v-1zM13.5 13h-1v1h1zm-1.5-.5v-1h-1v1zm-.5-1.5h-8v1h8zM1 12.5V6.497H0V12.5zm11.5.5a.5.5 0 0 1-.5-.5h-1a1.5 1.5 0 0 0 1.5 1.5zm-10 1A1.5 1.5 0 0 0 4 12.5H3a.5.5 0 0 1-.5.5zm-1-1a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 14zM9.802 2a2.5 2.5 0 0 1 2.143 1.213l.858-.515A3.5 3.5 0 0 0 9.802 1zM3.055 3.213A2.5 2.5 0 0 1 5.198 2V1a3.5 3.5 0 0 0-3 1.698zM14 12.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM2 10h3V9H2zm11-1h-3v1h3zM3 7h9V6H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.197 2.698A3.5 3.5 0 0 1 5.198 1h4.604a3.5 3.5 0 0 1 3 1.698L15 6.358V12.5a1.5 1.5 0 0 1-1.5 1.5h-1a1.5 1.5 0 0 1-1.5-1.5V12H4v.5A1.5 1.5 0 0 1 2.5 14h-1A1.5 1.5 0 0 1 0 12.5V6.358zM12 7H3V6h9zM2 10h3V9H2zm11-1h-3v1h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretVerticalCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.312 9.39l.39-.312l-.624-.78l-.39.312zM7.5 11l-.312.39l.312.25l.312-.25zM5.312 8.61l-.39-.313l-.625.781l.39.312zm-.624-3l-.39.312l.624.78l.39-.312zM7.5 4l.312-.39l-.312-.25l-.312.25zm2.188 2.39l.39.313l.625-.781l-.39-.312zm0 2.22l-2.5 2l.624.78l2.5-2zm-1.876 2l-2.5-2l-.624.78l2.5 2zm-2.5-4.22l2.5-2l-.624-.78l-2.5 2zm1.876-2l2.5 2l.624-.78l-2.5-2zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretVerticalCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M11 6L7.5 3.375L4 6zM4 9l3.5 2.625L11 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretVerticalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 13l-.332.374l.332.295l.332-.295zm0-11l.34-.367l-.333-.308l-.34.301zm.332 11.374l4.5-4l-.664-.748l-4.5 4zm0-.748l-4.5-4l-.664.748l4.5 4zm-.664-11l-4.5 4l.664.748l4.5-4zm-.008.74l4.313 4l.68-.733l-4.313-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretVerticalSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m10 9l-2.5 2L5 9m0-3l2.5-2L10 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretVerticalSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 3.375L11 6H4zm0 8.25L4 9h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretVerticalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1.336L2.17 6h10.66zm0 12.328L12.83 9H2.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartMinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l.6 2m0 0l2.4 8h11v-6a2 2 0 0 0-2-2zm4.9 4h5m1.5 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-8-1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartMinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.021.644L.979.356L1.472 2H12.5A2.5 2.5 0 0 1 15 4.5V11H3.128zM6 7h5V6H6zm-2 6.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m5.5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l.6 2m0 0l2.4 8h11v-6a2 2 0 0 0-2-2zm11.4 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-8-1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartPlusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l.6 2m0 0l2.4 8h11v-6a2 2 0 0 0-2-2zM8.5 4v5M6 6.5h5m1.5 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-8-1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartPlusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.021.644L.979.356L1.472 2H12.5A2.5 2.5 0 0 1 15 4.5V11H3.128zM8 9V7H6V6h2V4h1v2h2v1H9v2zm-4 4.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m5.5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.979.356L.02.644L3.128 11H15V4.5A2.5 2.5 0 0 0 12.5 2H1.472z"/><path fill="currentColor" fill-rule="evenodd" d="M5.5 12a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M5 13.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0m7.5-1.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m-.5 1.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CertificateOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 14.5H9a.5.5 0 0 0 .8.4zm2-1.5l.3-.4a.5.5 0 0 0-.6 0zm2 1.5l-.3.4a.5.5 0 0 0 .8-.4zm-2-3.5A2.5 2.5 0 0 1 9 8.5H8a3.5 3.5 0 0 0 3.5 3.5zM14 8.5a2.5 2.5 0 0 1-2.5 2.5v1A3.5 3.5 0 0 0 15 8.5zM11.5 6A2.5 2.5 0 0 1 14 8.5h1A3.5 3.5 0 0 0 11.5 5zm0-1A3.5 3.5 0 0 0 8 8.5h1A2.5 2.5 0 0 1 11.5 6zM9 10.5v4h1v-4zm.8 4.4l2-1.5l-.6-.8l-2 1.5zm1.4-1.5l2 1.5l.6-.8l-2-1.5zm2.8 1.1v-4h-1v4zM15 5V1.5h-1V5zm-1.5-5h-12v1h12zM0 1.5v12h1v-12zM1.5 15H8v-1H1.5zM0 13.5A1.5 1.5 0 0 0 1.5 15v-1a.5.5 0 0 1-.5-.5zM1.5 0A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5zM15 1.5A1.5 1.5 0 0 0 13.5 0v1a.5.5 0 0 1 .5.5zM3 5h5V4H3zm0 3h3V7H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CertificateSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v4l-.8-.601a4.5 4.5 0 0 0-6.3 6.3l.1.134V15H1.5A1.5 1.5 0 0 1 0 13.5zM8 5H3V4h5zM3 8h3V7H3z"/><path d="M11.5 5A3.5 3.5 0 0 0 9 10.95v3.55a.5.5 0 0 0 .8.4l1.7-1.275l1.7 1.275a.5.5 0 0 0 .8-.4v-3.55A3.5 3.5 0 0 0 11.5 5M10 13.5v-1.837c.455.216.963.337 1.5.337s1.045-.12 1.5-.337V13.5l-1.2-.9a.5.5 0 0 0-.6 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.5 13.5l.157-.475l-.218-.072l-.197.119zm2-2l-.421-.27l-.129.202l.076.226zm1 2.99l-.157.476a.5.5 0 0 0 .631-.634zm-3.258-1.418c-.956.575-2.485.919-3.742.919v1c1.385 0 3.106-.37 4.258-1.063zM7.5 13.99c-3.59 0-6.5-2.908-6.5-6.496H0a7.498 7.498 0 0 0 7.5 7.496zM1 7.495A6.498 6.498 0 0 1 7.5 1V0A7.498 7.498 0 0 0 0 7.495zM7.5 1C11.09 1 14 3.908 14 7.495h1A7.498 7.498 0 0 0 7.5 0zM14 7.495c0 1.331-.296 2.758-.921 3.735l.842.54C14.686 10.575 15 8.937 15 7.495zm-2.657 6.48l3 .99l.314-.949l-3-.99zm3.631.357l-1-2.99l-.948.316l1 2.991z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0A7.498 7.498 0 0 0 0 7.495a7.498 7.498 0 0 0 7.5 7.496c1.306 0 2.91-.328 4.054-.947l2.79.922a.5.5 0 0 0 .63-.634l-.926-2.771c.672-1.173.952-2.706.952-4.066A7.498 7.498 0 0 0 7.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatTypingAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7 7.5h1m-4 0h1m5 0h1m3.5 7h-7a7 7 0 1 1 7-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatTypingAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0V15H7.5A7.5 7.5 0 0 1 0 7.5M4 8h1V7H4zm7 0h-1V7h1zM7 8h1V7H7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatTypingOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222zm2 2.998l-.416.277a.5.5 0 0 0 .832 0zm2-2.998v-.5a.5.5 0 0 0-.416.222zm-4.416.277l2 2.998l.832-.555l-2-2.998zm2.832 2.998l2-2.998l-.832-.555l-2 2.998zM9.5 11.993h4v-1h-4zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5zm1.5-1.5V1.5h-1v8.994zM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5zM13.5 0h-12v1h12zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5zM0 1.5v8.994h1V1.499zm0 8.994c0 .829.671 1.499 1.5 1.499v-1a.499.499 0 0 1-.5-.5zm1.5 1.499h4v-1h-4zM7 7h1V6H7zM4 7h1V6H4zm6 0h1V6h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatTypingSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.994c0 .829-.671 1.499-1.5 1.499H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5zM4 7h1V6H4zm3 0h1V6H7zm4 0h-1V6h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatbotOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2.5V2zm-3 0V3zm6.856 9.422l-.35-.356l-.205.2l.07.277zM13.5 14.5l-.121.485a.5.5 0 0 0 .606-.606zm-4-1l-.354-.354l-.624.625l.857.214zm.025-.025l.353.354a.5.5 0 0 0-.4-.852zM.5 8H0zM7 0v2.5h1V0zm2 2H6v1h3zm6 6a6 6 0 0 0-6-6v1a5 5 0 0 1 5 5zm-1.794 4.279A5.983 5.983 0 0 0 15 7.999h-1a4.983 4.983 0 0 1-1.495 3.567zm.78 2.1L13.34 11.8l-.97.242l.644 2.578zm-4.607-.394l4 1l.242-.97l-4-1zm-.208-.863l-.025.024l.708.707l.024-.024zM9 14c.193 0 .384-.01.572-.027l-.094-.996A5.058 5.058 0 0 1 9 13zm-3 0h3v-1H6zM0 8a6 6 0 0 0 6 6v-1a5 5 0 0 1-5-5zm6-6a6 6 0 0 0-6 6h1a5 5 0 0 1 5-5zm1.5 6A1.5 1.5 0 0 1 6 6.5H5A2.5 2.5 0 0 0 7.5 9zM9 6.5A1.5 1.5 0 0 1 7.5 8v1A2.5 2.5 0 0 0 10 6.5zM7.5 5A1.5 1.5 0 0 1 9 6.5h1A2.5 2.5 0 0 0 7.5 4zm0-1A2.5 2.5 0 0 0 5 6.5h1A1.5 1.5 0 0 1 7.5 5zm0 8c1.064 0 2.042-.37 2.813-.987l-.626-.78c-.6.48-1.359.767-2.187.767zm-2.813-.987c.77.617 1.75.987 2.813.987v-1a3.483 3.483 0 0 1-2.187-.767z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatbotSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/><path fill="currentColor" fill-rule="evenodd" d="M9 2H8V0H7v2H6a6 6 0 0 0 0 12h3c.13 0 .26-.004.389-.013l3.99.998a.5.5 0 0 0 .606-.606l-.577-2.309A6 6 0 0 0 9 2M5 6.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0M7.5 12a4.483 4.483 0 0 1-2.813-.987l.626-.78c.599.48 1.359.767 2.187.767c.828 0 1.588-.287 2.187-.767l.626.78A4.483 4.483 0 0 1 7.5 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 5a.5.5 0 0 0 0-1zM2.964 2.814a.5.5 0 1 0-.928.372zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0 4h6V4h-6zM4.964 7.814l-2-5l-.928.372l2 5zM9.25 9.301l-3.65 4.9l.802.598l3.65-4.9zM7.5 10A2.5 2.5 0 0 1 5 7.5H4A3.5 3.5 0 0 0 7.5 11zM10 7.5A2.5 2.5 0 0 1 7.5 10v1A3.5 3.5 0 0 0 11 7.5zM7.5 5A2.5 2.5 0 0 1 10 7.5h1A3.5 3.5 0 0 0 7.5 4zm0-1A3.5 3.5 0 0 0 4 7.5h1A2.5 2.5 0 0 1 7.5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.503 1.907A7.472 7.472 0 0 1 7.5 0a7.498 7.498 0 0 1 6.635 4H7.5a3.501 3.501 0 0 0-3.23 2.149z"/><path fill="currentColor" d="M1.745 2.69a7.503 7.503 0 0 0 3.41 11.937l2.812-3.658a3.501 3.501 0 0 1-3.88-2.685a.502.502 0 0 1-.049-.092z"/><path fill="currentColor" d="M6.215 14.89a7.5 7.5 0 0 0 8.357-9.895A.503.503 0 0 1 14.5 5H9.95A3.49 3.49 0 0 1 11 7.5a3.487 3.487 0 0 1-.953 2.405z"/><path fill="currentColor" d="M5 7.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChurchOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8.5a.5.5 0 0 0-1 0zm10 0a.5.5 0 0 0-1 0zm-7.5 2V10H5v.5zm4 0h.5V10h-.5zM0 15h15v-1H0zm7.252-9.934l-7 4l.496.868l7-4zm7.496 4l-7-4l-.496.868l7 4zM7 0v2.5h1V0zm0 2.5v3h1v-3zM5 3h2.5V2H5zm2.5 0H10V2H7.5zM2 8.5v6h1v-6zm10 0v6h1v-6zm-6 6v-4H5v4zM5.5 11h4v-1h-4zm3.5-.5v4h1v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChurchSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2V0h1v2h2v1H8v2.21l6.748 3.856l-.496.868L13 9.22V14h2v1h-5v-5H5v5H0v-1h2V9.219l-1.252.715l-.496-.868L7 5.21V3H5V2z"/><path fill="currentColor" d="M6 15h3v-4H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 7.5a7 7 0 1 0 14 0a7 7 0 0 0-14 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardMinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m1 7h5M4.5.5h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardMinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5zm0 8h5V8H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardNoAccessOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m5.5 5l-4 4m-1-10h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zm3 11a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardNoAccessSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 6a2.5 2.5 0 0 0-2.086 3.879L8.88 6.414A2.488 2.488 0 0 0 7.5 6m0 5c-.51 0-.983-.152-1.379-.414L9.586 7.12A2.5 2.5 0 0 1 7.5 11"/><path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5zM4 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m.5-1h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardPlusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4M7.5 6v5M5 8.5h5M4.5.5h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardPlusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5zm2 10V9H5V8h2V6h1v2h2v1H8v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 0H4v1H1v12.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V1h-3zm-1 1H5v1.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardTickOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m1 7l2 2l3.5-4m-6-6h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardTickSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5zm2.024 10.232l3.852-4.403l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardXoutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m1.5 5l4 4m-4 0l4-4m-5-6h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardXsolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5zm4.146 9.854L7.5 9.207l-1.646 1.646l-.708-.707L6.793 8.5L5.146 6.854l.708-.708L7.5 7.793l1.646-1.647l.708.708L8.207 8.5l1.647 1.646z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 7.5H7a.5.5 0 0 0 .146.354zm0 6.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM7 3v4.5h1V3zm.146 4.854l3 3l.708-.708l-3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m7 0V3h1v4.293l2.854 2.853l-.708.708l-3-3A.499.499 0 0 1 7 7.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockwiseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 8.495v-.5h-1v.5zM7.5 2.999H8v-1h-.5zm1-.5l.353.353l.354-.353l-.354-.354zM13 8.495a5.499 5.499 0 0 1-5.5 5.496v1c3.589 0 6.5-2.909 6.5-6.496zM7.5 13.99A5.499 5.499 0 0 1 2 8.495H1a6.499 6.499 0 0 0 6.5 6.496zM2 8.495a5.499 5.499 0 0 1 5.5-5.496v-1A6.499 6.499 0 0 0 1 8.495zM6.147.854l2 1.998l.706-.707l-2-1.999zm2 1.291l-2 1.999l.706.707l2-1.999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockwiseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.295 2.002L6.147.854l.706-.708L9.207 2.5L6.853 4.85l-.706-.707l1.141-1.141A5.499 5.499 0 0 0 2 8.495a5.499 5.499 0 0 0 5.5 5.496A5.5 5.5 0 0 0 13 8.495v-.5h1v.5a6.499 6.499 0 0 1-6.5 6.496A6.499 6.499 0 0 1 1 8.495a6.499 6.499 0 0 1 6.295-6.493" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.146 10.146l-.353.354l.707.707l.354-.353zM13.5 7.5l.354.354l.353-.354l-.353-.354zm-2.646-3.354l-.354-.353l-.707.707l.353.354zm-6.708 6.708l.354.353l.707-.707l-.353-.354zM1.5 7.5l-.354-.354l-.353.354l.353.354zm3.354-2.646l.353-.354l-.707-.707l-.354.353zm6 6l3-3l-.708-.708l-3 3zm3-3.708l-3-3l-.708.708l3 3zm-9 3l-3-3l-.708.708l3 3zm-3-2.292l3-3l-.708-.708l-3 3zm6.153-6.436l-2 12l.986.164l2-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.007 13.418l2-12l.986.164l-2 12zm-.8-8.918l-3 3l3 3l-.707.707L.793 7.5L4.5 3.793zm5.293-.707L14.207 7.5L10.5 11.207l-.707-.707l3-3l-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodepenOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.29-.407a.5.5 0 0 0-.58 0zm7 5h.5a.5.5 0 0 0-.21-.407zm0 4l.29.407A.5.5 0 0 0 15 9.5zm-7 5l-.29.407a.5.5 0 0 0 .58 0zm-7-5H0a.5.5 0 0 0 .21.407zm0-4l-.29-.407A.5.5 0 0 0 0 5.5zM7.21.907l7 5l.58-.814l-7-5zM14 5.5v4h1v-4zm.21 3.593l-7 5l.58.814l7-5zm-6.42 5l-7-5l-.58.814l7 5zM1 9.5v-4H0v4zM.79 5.907l7-5l-.58-.814l-7 5zm0 4l7-5l-.58-.814l-7 5zm6.42-5l7 5l.58-.814l-7-5zm-7 1l7 5l.58-.814l-7-5zm7.58 5l7-5l-.58-.814l-7 5zM7 .5v4h1v-4zm0 10v4h1v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodepenSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.21.093a.5.5 0 0 1 .58 0l7 5A.5.5 0 0 1 15 5.5v4a.5.5 0 0 1-.21.407l-7 5a.5.5 0 0 1-.58 0l-7-5A.5.5 0 0 1 0 9.5v-4a.5.5 0 0 1 .21-.407zM1 6.472L2.44 7.5L1 8.528zM1.36 9.5L7 13.528v-2.77L3.3 8.113zm2.8-2L7.5 9.886L10.84 7.5L7.5 5.114zM8 4.243l3.7 2.643L13.64 5.5L8 1.472zM7 1.472v2.77L3.3 6.887L1.36 5.5zm7 5L12.56 7.5L14 8.528zM13.64 9.5L11.7 8.114L8 10.757v2.771z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" clip-rule="evenodd"><path d="m5.944.5l-.086.437l-.329 1.598a5.52 5.52 0 0 0-1.434.823L2.487 2.82l-.432-.133l-.224.385L.724 4.923L.5 5.31l.328.287l1.244 1.058c-.045.277-.103.55-.103.841c0 .291.058.565.103.842L.828 9.395L.5 9.682l.224.386l1.107 1.85l.224.387l.432-.135l1.608-.537c.431.338.908.622 1.434.823l.329 1.598l.086.437h3.111l.087-.437l.328-1.598a5.524 5.524 0 0 0 1.434-.823l1.608.537l.432.135l.225-.386l1.106-1.851l.225-.386l-.329-.287l-1.244-1.058c.046-.277.103-.55.103-.842c0-.29-.057-.564-.103-.841l1.244-1.058l.329-.287l-.225-.386l-1.106-1.85l-.225-.386l-.432.134l-1.608.537a5.52 5.52 0 0 0-1.434-.823L9.142.937L9.055.5z"/><path d="M9.5 7.495a2 2 0 0 1-4 0a2 2 0 0 1 4 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.026 7.5c0-.827.66-1.5 1.473-1.5c.813 0 1.473.673 1.473 1.5S8.312 9 7.499 9c-.813 0-1.473-.673-1.473-1.5"/><path fill="currentColor" fill-rule="evenodd" d="M5.568 0h3.86l.164.837V.84l.27 1.335c.383.172.74.386 1.068.627l1.345-.458l.796-.251l.417.727l1.087 1.854l.425.743l-.633.563l-1.009.874c.032.197.061.418.061.646c0 .228-.03.45-.06.646l1.012.878l.629.559l-.428.748l-1.084 1.849l-.417.73l-.807-.258l-1.334-.454a5.877 5.877 0 0 1-1.068.627l-.27 1.335v.003L9.43 15H5.57l-.163-.839l-.27-1.336a5.877 5.877 0 0 1-1.068-.627l-1.343.457l-.799.255l-.415-.73l-1.088-1.855L0 9.583l.632-.563l1.008-.875a4.075 4.075 0 0 1-.06-.645c0-.227.03-.45.061-.645l-1.014-.88L0 5.418l.426-.748l1.085-1.85l.415-.728l.808.256l1.334.454a5.875 5.875 0 0 1 1.068-.627l.27-1.337zM7.5 5C6.144 5 5.045 6.12 5.045 7.5s1.1 2.5 2.454 2.5c1.355 0 2.454-1.12 2.454-2.5S8.853 5 7.5 5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 10.5l-.447-.224a.5.5 0 0 0 .67.671zm2-4l-.224-.447a.5.5 0 0 0-.223.223zm4-2l.447.224a.5.5 0 0 0-.67-.671zm-2 4l.224.447a.5.5 0 0 0 .223-.223zm-1 5.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM4.947 10.724l2-4l-.894-.448l-2 4zm1.777-3.777l4-2l-.448-.894l-4 2zm3.329-2.67l-2 4l.894.447l2-4zM8.276 8.052l-4 2l.448.894l4-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.618 9.382l1.255-2.51l2.509-1.254l-1.255 2.51z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m10.947-2.776a.5.5 0 0 0-.67-.671l-4 2a.5.5 0 0 0-.224.223l-2 4a.5.5 0 0 0 .67.671l4-2a.5.5 0 0 0 .224-.223z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComputerOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 14.5h4a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Zm0 0H4m2.5-3v3m2-9h6m-4.5 6h3m-11.5-8h7v8h-7a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComputerSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12h3v-1h-3z"/><path fill="currentColor" fill-rule="evenodd" d="M9.5 0A1.5 1.5 0 0 0 8 1.5V3H1.5A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12H6v2H4v1h9.5a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0zM8.085 14H7v-2h1v1.5c0 .175.03.344.085.5M9.5 14h4a.5.5 0 0 0 .5-.5V6H9v7.5a.5.5 0 0 0 .5.5M9 5h5V1.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0-.5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12.5v.5h1v-.5zm5 0v.5h1v-.5zm-4 0V12H2v.5zm4-.5v.5h1V12zm-2-2a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3zm-2 2a2 2 0 0 1 2-2V9a3 3 0 0 0-3 3zm2-8a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1zm2 2a2 2 0 0 0-2-2v1a1 1 0 0 1 1 1zM5 8a2 2 0 0 0 2-2H6a1 1 0 0 1-1 1zm0-1a1 1 0 0 1-1-1H3a2 2 0 0 0 2 2zM1.5 3h12V2h-12zm12.5.5v8h1v-8zm-.5 8.5h-12v1h12zM1 11.5v-8H0v8zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 13zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM13.5 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 2zm-12-1A1.5 1.5 0 0 0 0 3.5h1a.5.5 0 0 1 .5-.5zM9 6h3V5H9zm0 3h3V8H9zm-9 6h15v-1H0zM3 0v2.5h1V0zm8 0v2.5h1V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 2H4V0H3v2H1.5A1.5 1.5 0 0 0 0 3.5v8A1.5 1.5 0 0 0 1.5 13h12a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 2H12V0h-1zM3 6a2 2 0 1 1 4 0a2 2 0 0 1-4 0m-.618 4.618a2.927 2.927 0 0 1 5.236 0l.33.658A.5.5 0 0 1 7.5 12h-5a.5.5 0 0 1-.447-.724zM9 6h3V5H9zm0 3h3V8H9z" clip-rule="evenodd"/><path fill="currentColor" d="M15 14v1H0v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContractOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 7H4v1h.5zm6 1h.5V7h-.5zm-6-4H4v1h.5zm2 1H7V4h-.5zm4-4.5l.354-.354L10.707 0H10.5zm3 3h.5v-.207l-.146-.147zM8 11l-.354.354zm-.5.5l.224.447l.04-.02l.036-.027zM4.5 8h6V7h-6zm0-3h2V4h-2zm8 9h-10v1h10zM2 13.5v-12H1v12zM2.5 1h8V0h-8zM13 3.5v10h1v-10zM10.146.854l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zm3.474 10.158c.111-.333.427-.642.83-.75c.377-.101.862-.035 1.342.446l.708-.708c-.72-.72-1.569-.903-2.309-.704c-.713.192-1.297.733-1.52 1.4zm2.172-.304a.933.933 0 0 1 .079.087l.79-.614a1.945 1.945 0 0 0-.161-.18zm.079.087c.078.1.06.132.063.11c.002-.014.006.009-.054.063a1 1 0 0 1-.29.169a1.781 1.781 0 0 1-.394.108a.848.848 0 0 1-.25.01c-.017-.004.018 0 .07.037a.388.388 0 0 1 .131.43a.209.209 0 0 1-.023.047c-.002.002.015-.02.072-.067c.114-.092.324-.226.674-.4l-.448-.895c-.377.188-.66.36-.854.517a1.375 1.375 0 0 0-.26.267a.705.705 0 0 0-.14.438a.61.61 0 0 0 .255.468c.113.084.238.12.33.139c.187.037.4.027.593-.002c.38-.058.872-.222 1.207-.526c.174-.159.339-.387.374-.686c.036-.306-.074-.593-.267-.84zm.075.459a2.56 2.56 0 0 1 .518-.307l-.397-.918c-.24.104-.48.245-.721.425zm.518-.307c.65-.281 1.231-.133 1.826.15c.15.072.296.15.444.23c.144.078.296.161.44.235c.276.139.618.292.972.292v-1c-.094 0-.248-.047-.52-.184c-.128-.066-.262-.14-.416-.223c-.15-.081-.316-.17-.49-.252c-.698-.333-1.611-.616-2.653-.166z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContractSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.796 11.9H6.8z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM7 4H4v1h3zm4 3H4v1h7zm-4.695 3.908c-.404.108-.72.417-.83.75l-.95-.316c.223-.667.807-1.208 1.52-1.4c.707-.19 1.514-.03 2.212.611a2.75 2.75 0 0 1 .622-.107c.54-.029 1.023.107 1.438.28c.305.127.6.287.85.422c.078.044.153.084.222.12c.323.17.5.232.611.232v1c-.39 0-.774-.188-1.076-.346a21.802 21.802 0 0 1-.272-.146a7.689 7.689 0 0 0-.72-.359c-.334-.14-.663-.222-.999-.204a1.686 1.686 0 0 0-.15.014l.001.014c.027.324-.107.591-.28.783c-.318.354-.837.54-1.227.61a1.962 1.962 0 0 1-.614.025a.9.9 0 0 1-.33-.11a.623.623 0 0 1-.303-.433a.677.677 0 0 1 .111-.48a1.28 1.28 0 0 1 .262-.282c.19-.157.465-.327.834-.513l.027-.02a1.23 1.23 0 0 0-.96-.145" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CostEstimateOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 7H4v1h.5zm6 1h.5V7h-.5zm-2 2H8v1h.5zm2 1h.5v-1h-.5zm-6-7H4v1h.5zm2 1H7V4h-.5zm4-4.5l.354-.354L10.707 0H10.5zm3 3h.5v-.207l-.146-.147zM4.5 8h6V7h-6zm4 3h2v-1h-2zm-4-6h2V4h-2zm8 9h-10v1h10zM2 13.5v-12H1v12zM2.5 1h8V0h-8zM13 3.5v10h1v-10zM10.146.854l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CostEstimateSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM4 4h3v1H4zm7 3H4v1h7zm0 3H8v1h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CplusplusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 9.5c-.64.64-1.509 1-2.414 1H6.5a3 3 0 0 1 0-6h.586c.905 0 1.774.36 2.414 1m-2 .5v3M6 7.5h6M10.5 6v3m-9 1.5v-6l6-3.5l6 3.5v6l-6 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CplusplusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.421L14 4.213v6.574L7.5 14.58L1 10.787V4.213zM6.5 4a3.5 3.5 0 1 0 0 7h.586a3.914 3.914 0 0 0 2.768-1.146l-.708-.708a2.914 2.914 0 0 1-2.06.854H6.5a2.5 2.5 0 0 1 0-5h.586a2.91 2.91 0 0 1 2.06.854l.708-.708A3.914 3.914 0 0 0 7.086 4zM7 7V6h1v1h2V6h1v1h1v1h-1v1h-1V8H8v1H7V8H6V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 5.5h14M2 9.5h6m2 0h3M.5 3.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 2A1.5 1.5 0 0 1 15 3.5V5H0V3.5A1.5 1.5 0 0 1 1.5 2z"/><path fill="currentColor" fill-rule="evenodd" d="M0 6v5.5A1.5 1.5 0 0 0 1.5 13h12a1.5 1.5 0 0 0 1.5-1.5V6zm2 4h6V9H2zm11 0h-3V9h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 3.5h8v8m-8-8V0m0 3.5H0m11.5 8h-8V6m8 5.5V15m0-3.5H15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 3V0h1v3h8v8h3v1h-3v3h-1v-3H3V6h1v5h7V4H0V3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThreeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.498.542zm14 0l.498.042A.5.5 0 0 0 14.5 0zm-1 12l.158.474a.5.5 0 0 0 .34-.433zm-6 2l-.158.474a.499.499 0 0 0 .316 0zm-6-2l-.498.041a.5.5 0 0 0 .34.433zm9-9h.5V3h-.5zm0 6l.158.474L11 9.86V9.5zm-3 1l-.158.474l.158.053l.158-.053zm-3-1H4v.36l.342.114zM.5 1h14V0H.5zM14.002.458l-1 12l.996.083l1-12zm-.66 11.568l-6 2l.316.948l6-2zm-5.684 2l-6-2l-.316.948l6 2zm-5.66-1.567l-1-12l-.996.083l1 12zM10.5 3H4v1h6.5zM6 7h4.5V6H6zm4-.5v3h1v-3zm.342 2.526l-3 1l.316.948l3-1zm-2.684 1l-3-1l-.316.948l3 1zM5 9.5V8H4v1.5zm5-6v3h1v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThreeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.132.162A.5.5 0 0 1 .5 0h14a.5.5 0 0 1 .498.542l-1 11.916a.5.5 0 0 1-.34.432l-6 2a.5.5 0 0 1-.316 0l-6-2a.5.5 0 0 1-.34-.432L.002.542a.5.5 0 0 1 .13-.38M11 3H4v1h6v2H6v1h4v2.14l-2.5.833L5 9.14V8H4v1.86l3.5 1.167L11 9.86z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CsvOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-4 6V6H6v.5zm0 2H6V9h.5zm2 0H9V8h-.5zm0 2v.5H9v-.5zm2-1H10v.207l.146.147zm1 1l-.354.354l.354.353l.354-.353zm1-1l.354.354l.146-.147V9.5zm-10-3V6H2v.5zm0 4H2v.5h.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zM9 6H6.5v1H9zm-3 .5v2h1v-2zM6.5 9h2V8h-2zM8 8.5v2h1v-2zm.5 1.5H6v1h2.5zM10 6v3.5h1V6zm.146 3.854l1 1l.708-.708l-1-1zm1.708 1l1-1l-.708-.708l-1 1zM13 9.5V6h-1v3.5zM5 6H2.5v1H5zm-3 .5v4h1v-4zm.5 4.5H5v-1H2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CsvSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM2 6h3v1H3v3h2v1H2zm7 0H6v3h2v1H6v1h3V8H7V7h2zm2 0h-1v3.707l1.5 1.5l1.5-1.5V6h-1v3.293l-.5.5l-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CupOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.5 6.5v5a3 3 0 0 1-3 3h-5a3 3 0 0 1-3-3v-5a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1Zm0 0h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2M4.5 4V2m3 2V0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CupSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 4h1V0H7zM5 2v2H4V2z"/><path fill="currentColor" fill-rule="evenodd" d="M0 6.5A1.5 1.5 0 0 1 1.5 5h9a1.5 1.5 0 0 1 1.415 1H13.5A1.5 1.5 0 0 1 15 7.5v2a1.5 1.5 0 0 1-1.5 1.5H12v.5A3.5 3.5 0 0 1 8.5 15h-5A3.5 3.5 0 0 1 0 11.5zM12 10h1.5a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5H12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurvedConnectorOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 1.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 0h2a3 3 0 0 1 3 3v6a3 3 0 0 0 3 3h2m0 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurvedConnectorSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0a1.5 1.5 0 1 0 1.415 2H4.5A2.5 2.5 0 0 1 7 4.5v6a3.5 3.5 0 0 0 3.5 3.5h1.585a1.5 1.5 0 1 0 0-1H10.5A2.5 2.5 0 0 1 8 10.5v-6A3.5 3.5 0 0 0 4.5 1H2.915A1.5 1.5 0 0 0 1.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DthreeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 1.5h1.5a6 6 0 1 1 0 12H0m7-12h4.5a3 3 0 1 1 0 6m0 0H9m2.5 0h-2m2 0a3 3 0 1 1 0 6H7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DthreeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2H0V1h1.5a6.5 6.5 0 0 1 0 13H0v-1h1.5a5.5 5.5 0 1 0 0-11m10 0H7V1h4.5a3.5 3.5 0 0 1 1.804 6.5A3.5 3.5 0 0 1 11.5 14H7v-1h4.5a2.5 2.5 0 0 0 0-5H9V7h2.5a2.5 2.5 0 0 0 0-5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M14.5 2.499c0 1.103-3.134 1.998-7 1.998S.5 3.602.5 2.5m14 0c0-1.105-3.134-2-7-2s-7 .895-7 1.999m14 0v9.993c0 1.103-3.134 1.999-7 1.999s-7-.895-7-1.999V2.5m14 4.996c0 1.104-3.134 2-7 2s-7-.896-7-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0C5.534 0 3.736.227 2.413.605c-.658.188-1.227.42-1.643.701C.372 1.576 0 1.97 0 2.5v9.993c0 .53.372.924.77 1.192c.416.281.985.514 1.643.702c1.323.378 3.121.605 5.087.605c1.966 0 3.764-.227 5.087-.605c.658-.188 1.227-.421 1.643-.702c.398-.268.77-.662.77-1.192V2.5c0-.53-.372-.924-.77-1.193c-.416-.28-.985-.513-1.643-.701C11.264.227 9.466 0 7.5 0M1.262 2.864l.452.214c1.127.534 3.27.92 5.786.92c2.517 0 4.659-.386 5.786-.92l.452-.214l.428.904l-.452.214c-1.323.627-3.638 1.017-6.214 1.017c-2.576 0-4.891-.39-6.214-1.017l-.452-.215zm.452 5.184l-.452-.214l-.428.904l.452.214c1.323.627 3.638 1.017 6.214 1.017c2.576 0 4.891-.39 6.214-1.017l.452-.214l-.428-.904l-.452.214c-1.127.534-3.27.92-5.786.92c-2.517 0-4.659-.386-5.786-.92" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeniedOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m2.5 2.5l10 10m-5 2a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeniedSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m2.564-4.23a6.5 6.5 0 0 0 9.165 9.165zm.707-.706l9.165 9.165a6.5 6.5 0 0 0-9.165-9.165" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DenoOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 14.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm0 0v-5H6a2.5 2.5 0 0 1 0-5h1.694a3.495 3.495 0 0 1 3.49 3.301L11.5 13.5M7 9.5h.382c.685 0 1.312-.387 1.618-1m-5-2h1m1 0h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DenoSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7H6V6h1z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15M8 13.981a6.462 6.462 0 0 0 2.971-.985l-.287-5.167A2.995 2.995 0 0 0 7.694 5H6a2 2 0 0 0-1.732 1H5v1H4a2 2 0 0 0 2 2h.882c.496 0 .95-.28 1.17-.724l.895.448A2.308 2.308 0 0 1 8 9.71z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DepthChartOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 0v14.5h14V0M.5 4.5h2v1h2v3h2v3h1v3v-2h2v-2h2v-3h1v-2h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DepthChartSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0H0v15h15V0h-1v5h-2v2h-1v3H9v2H8v-1H7V8H5V5H3V4H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesklampOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.5 8.5l-.395-.307a.5.5 0 0 0 .041.66zm5-7l-.354.354zm2 2l.354.354l.353-.354l-.353-.354zm-2 2l-.354.354l.354.353l.354-.353zm-2-2l-.354.354zm3.136.864l.354.353zm-.272.272l-.354-.353zm.272 5.728l-.353-.354zm5.728-5.728l-.354-.353zm0-.272l-.354.353zm-6 6l.353-.354zM2 15h11v-1H2zm5.854-.854l-6-6l-.708.708l6 6zm-5.96-5.339l3.5-4.5l-.789-.614l-3.5 4.5zm4.252-6.953l2 2l.708-.708l-2-2zm2 1.292l-2 2l.708.708l2-2zm-1.292 2l-2-2l-.708.708l2 2zm-2-3.292a.914.914 0 0 1 1.292 0l.708-.708a1.914 1.914 0 0 0-2.708 0zm-.708-.708a1.914 1.914 0 0 0 0 2.708l.708-.708a.914.914 0 0 1 0-1.292zM7.283 4.01l-.273.273l.707.707l.273-.273zm.707 6.707l5.727-5.727l-.707-.707l-5.727 5.727zm5.727-5.727c.27-.27.27-.71 0-.98l-.707.707a.307.307 0 0 1 0-.434zM7.01 10.717c.27.27.71.27.98 0l-.707-.707c.12-.12.314-.12.434 0zm0-6.434a4.55 4.55 0 0 0 0 6.434l.707-.707a3.55 3.55 0 0 1 0-5.02zm.98.434a3.55 3.55 0 0 1 5.02 0l.707-.707a4.55 4.55 0 0 0-6.434 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesklampSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.854 1.146l1.884 1.885a4.551 4.551 0 0 1 4.98.98c.27.27.27.708 0 .979L7.99 10.717c-.27.27-.71.27-.98 0a4.551 4.551 0 0 1-.979-4.979l-.984-.984L2.166 8.46L7.707 14H13v1H2v-1h4.293L1.146 8.854a.5.5 0 0 1-.04-.66L4.333 4.04l-.188-.187a1.914 1.914 0 1 1 2.708-2.708"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiamondOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 14.5l-.395.307a.5.5 0 0 0 .79 0zm-7-9l-.429-.257a.5.5 0 0 0 .034.564zm3-5V0h-.283L3.07.243zm8 0l.429-.257L11.783 0H11.5zm3 5l.395.307a.5.5 0 0 0 .034-.564zm-6.605 8.693l-7-9l-.79.614l7 9zM.929 5.757l3-5L3.07.243l-3 5zM3.5 1h8V0h-8zm7.571-.243l3 5l.858-.514l-3-5zm3.034 4.436l-7 9l.79.614l7-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiamondSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.783 0H3.217L.07 5.243a.5.5 0 0 0 .034.564l7 9a.5.5 0 0 0 .79 0l7-9a.5.5 0 0 0 .034-.564z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m.5.5l6 14l2-6l6-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.697.04A.5.5 0 0 0 .04.697l6 14a.5.5 0 0 0 .934-.039l1.921-5.763l5.763-1.92a.5.5 0 0 0 .039-.935z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscordOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.5 13.5l-.326.379a.5.5 0 0 0 .342.12zm-1.066-1.712a.5.5 0 0 0-.785.62zm.398-.41l-.174-.468a.672.672 0 0 0-.02.007zm-1.738.516L9.01 11.4l-.008.001zm-3.104-.012l-.095.49l.003.001zm-1.762-.515l-.182.465zm-.875-.408l-.278.415a.46.46 0 0 0 .033.021zm-.108-.06l.277-.416a.491.491 0 0 0-.054-.031zm-.048-.036l.353-.354a.502.502 0 0 0-.11-.083zm2.154 1.52a.5.5 0 0 0-.785-.62zM3.5 13.5l-.016.5a.5.5 0 0 0 .347-.125zm-3-2.253H0a.5.5 0 0 0 .006.08zm1.726-8.488l-.3-.4a.5.5 0 0 0-.168.225zM5.594 1.5l.498-.047A.5.5 0 0 0 5.605 1zm-.378 1.306a.5.5 0 0 0 .996-.095zm3.526-.063a.5.5 0 0 0 .992.127zM9.406 1.5L9.395 1a.5.5 0 0 0-.485.436zm3.368 1.259l.468-.175a.5.5 0 0 0-.168-.225zm1.726 8.488l.494.08a.497.497 0 0 0 .006-.08zM6.481 8.8l-.5-.008V8.8zm5.019 4.7l.326-.379l-.002-.002a.794.794 0 0 1-.044-.038a21.355 21.355 0 0 1-.536-.48c-.325-.298-.66-.622-.81-.813l-.785.62c.208.264.603.64.918.93a29.109 29.109 0 0 0 .593.53l.01.008l.003.002zm.436-3.246c-.46.303-.894.513-1.278.656l.348.937a7.352 7.352 0 0 0 1.48-.758zm-1.297.663a7.387 7.387 0 0 1-1.629.484l.168.986a8.39 8.39 0 0 0 1.848-.548zm-1.637.485a7.895 7.895 0 0 1-2.92-.012l-.184.983a8.896 8.896 0 0 0 3.288.012zm-2.917-.011a9.57 9.57 0 0 1-1.675-.49l-.364.931c.512.2 1.13.402 1.849.54zm-1.675-.49a6.536 6.536 0 0 1-.813-.378l-.489.872c.326.183.648.324.938.437zm-.78-.358a.802.802 0 0 0-.108-.061c-.02-.01-.011-.007 0 .001l-.555.832a.87.87 0 0 0 .108.061c.021.01.012.007 0-.002zm-.162-.091a.332.332 0 0 1 .082.058l-.707.707c.023.023.081.08.178.13zm-.028-.026a4.697 4.697 0 0 1-.28-.168l-.011-.008a.025.025 0 0 0-.001 0l-.287.41l-.286.409l.001.001l.002.002l.007.004l.021.014l.075.049c.064.04.156.096.273.161zm1.126 1.338c-.152.193-.489.525-.813.829a30.38 30.38 0 0 1-.538.491l-.034.031l-.01.008l-.001.002h-.001l.331.375l.331.375l.001-.001l.003-.002l.01-.009l.036-.032a38.039 38.039 0 0 0 .555-.508c.315-.296.708-.677.915-.94zM3.516 13c-1.166-.037-1.778-.521-2.11-.96a2.394 2.394 0 0 1-.4-.82a1.1 1.1 0 0 1-.013-.056v.002l-.493.08c-.494.08-.494.08-.493.081v.006a1.367 1.367 0 0 0 .028.127a3.394 3.394 0 0 0 .573 1.183c.505.667 1.393 1.31 2.876 1.357zM1 11.247c0-1.867.42-3.94.847-5.564a35.45 35.45 0 0 1 .776-2.552a16.43 16.43 0 0 1 .067-.186l.004-.01v-.001l-.468-.175l-.469-.175v.001l-.001.003l-.004.011a9.393 9.393 0 0 0-.072.2a36.445 36.445 0 0 0-.8 2.629C.443 7.083 0 9.253 0 11.247zm1.526-8.088c.8-.6 1.577-.89 2.15-1.03a4.764 4.764 0 0 1 .86-.128A1.48 1.48 0 0 1 5.585 2h-.001l.01-.5l.01-.5H5.6a.848.848 0 0 0-.028 0h-.068a3.973 3.973 0 0 0-.24.016a5.763 5.763 0 0 0-.825.141a6.938 6.938 0 0 0-2.513 1.2zm2.57-1.612l.12 1.259l.996-.095l-.12-1.258zM9.734 2.87l.168-1.306l-.992-.128l-.168 1.307zM9.406 1.5l.01.5h-.001a.497.497 0 0 1 .049 0c.038.002.1.005.179.013c.16.014.394.047.681.117a5.94 5.94 0 0 1 2.15 1.029l.6-.8a6.937 6.937 0 0 0-2.513-1.2a5.76 5.76 0 0 0-.825-.142A3.98 3.98 0 0 0 9.399 1h-.003zm3.368 1.259l-.469.174l.001.003l.004.009l.013.037l.053.149a35.482 35.482 0 0 1 .777 2.552c.428 1.624.847 3.697.847 5.564h1c0-1.994-.444-4.164-.88-5.819a36.512 36.512 0 0 0-.8-2.629a15.246 15.246 0 0 0-.057-.158l-.015-.042l-.004-.01l-.001-.004zm1.726 8.488l-.493-.08v-.003l-.002.008l-.01.047c-.012.045-.03.113-.061.197c-.062.17-.167.396-.34.624c-.332.439-.944.923-2.11.96l.032 1c1.483-.047 2.37-.69 2.876-1.356a3.395 3.395 0 0 0 .573-1.184a2.05 2.05 0 0 0 .026-.118l.002-.01v-.004c0-.001 0-.002-.493-.081M5.259 6.97c-1.002 0-1.723.867-1.723 1.83h1c0-.498.357-.83.723-.83zM3.536 8.8c0 .967.736 1.83 1.723 1.83v-1c-.357 0-.723-.334-.723-.83zm1.723 1.83c1 0 1.722-.866 1.722-1.83h-1c0 .5-.357.83-.722.83zM6.98 8.81c.016-.978-.728-1.84-1.722-1.84v1.001c.372 0 .73.338.722.822zm2.653-1.84c-1.002.001-1.723.868-1.723 1.831h1c0-.498.357-.83.723-.83zM7.91 8.802c0 .967.736 1.83 1.723 1.83v-1c-.357 0-.723-.334-.723-.83zm1.723 1.83c1 0 1.722-.866 1.722-1.83h-1c0 .5-.357.83-.722.83zm1.722-1.83c0-.963-.721-1.83-1.722-1.83v1c.365 0 .722.332.722.83zM3.74 4.44c1.443-.787 2.619-1.154 3.763-1.155c1.145 0 2.318.365 3.758 1.154l.48-.876c-1.522-.835-2.865-1.279-4.238-1.278c-1.373 0-2.717.445-4.241 1.277z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscordSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.42 7.727c0-.534.382-.89.774-.89c.399 0 .783.363.774.882v.008c0 .535-.383.89-.774.89c-.382 0-.774-.359-.774-.89m4.687 0c0-.282.086-.48.194-.599a.622.622 0 0 1 .485-.195c.391 0 .774.355.774.89c0 .534-.383.89-.774.89c-.16 0-.316-.078-.45-.255a1.233 1.233 0 0 1-.229-.73"/><path fill="currentColor" fill-rule="evenodd" d="M3.8 12.142c.277.108.583.217.918.317c-.59.612-1.504 1.47-1.504 1.47C.376 13.839 0 11.514 0 11.514C0 7.38 1.85 2.42 1.85 2.42c1.848-1.387 3.607-1.35 3.607-1.35l.163 1.096c-.852.227-1.726.601-2.663 1.113l.513.94c1.546-.843 2.805-1.236 4.032-1.237c1.226 0 2.483.391 4.025 1.237l.515-.94c-.983-.539-1.896-.925-2.788-1.147l.287-1.062s1.76-.038 3.609 1.349c0 0 1.849 4.959 1.849 9.094c0 0-.376 2.325-3.214 2.415c0 0-.854-.842-1.418-1.449c.296-.09.596-.2.9-.326a7.88 7.88 0 0 0 1.574-.809l-.59-.895c-.492.325-.957.55-1.368.703l-.007.003l-.005.002l-.01.004a7.916 7.916 0 0 1-1.744.518l-.009.001a8.462 8.462 0 0 1-3.127-.012a10.255 10.255 0 0 1-1.793-.525c-.504-.197-.935-.406-1.352-.698l-.614.878a7.3 7.3 0 0 0 1.576.818zm1.394-6.376c-1.073 0-1.846.93-1.846 1.961c0 1.036.79 1.962 1.846 1.962c1.071 0 1.843-.927 1.845-1.957c.015-1.046-.782-1.966-1.845-1.966m4.592.096c-.507 0-.959.193-1.28.547c-.315.35-.47.818-.47 1.318c0 .5.156.995.446 1.378c.293.388.744.68 1.304.68c1.073 0 1.845-.93 1.845-1.962s-.772-1.961-1.845-1.961" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscountOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 5.5h1m3 4h1M10 5l-5 5M6.801.79L5.672 1.917a.988.988 0 0 1-.698.29H3.196a.988.988 0 0 0-.988.988v1.778a.988.988 0 0 1-.29.698L.79 6.802a.988.988 0 0 0 0 1.397l1.13 1.129a.987.987 0 0 1 .289.698v1.778c0 .546.442.988.988.988h1.778c.262 0 .513.104.698.29l1.13 1.129a.988.988 0 0 0 1.397 0l1.129-1.13a.988.988 0 0 1 .698-.289h1.778a.988.988 0 0 0 .988-.988v-1.778c0-.262.104-.513.29-.698l1.129-1.13a.988.988 0 0 0 0-1.397l-1.13-1.129a.988.988 0 0 1-.289-.698V3.196a.988.988 0 0 0-.988-.988h-1.778a.988.988 0 0 1-.698-.29L8.198.79a.988.988 0 0 0-1.397 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscountSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.448.436l-1.13 1.129a.488.488 0 0 1-.344.143H3.196c-.822 0-1.488.666-1.488 1.488v1.778a.49.49 0 0 1-.143.345L.435 6.448a1.488 1.488 0 0 0 0 2.104l1.13 1.13a.488.488 0 0 1 .143.344v1.778c0 .822.666 1.488 1.488 1.488h1.778a.49.49 0 0 1 .345.143l1.129 1.13a1.488 1.488 0 0 0 2.104 0l1.13-1.13a.488.488 0 0 1 .344-.143h1.778c.822 0 1.488-.666 1.488-1.488v-1.778a.49.49 0 0 1 .143-.345l1.13-1.129a1.488 1.488 0 0 0 0-2.104l-1.13-1.13a.488.488 0 0 1-.143-.344V3.196c0-.822-.666-1.488-1.488-1.488h-1.778a.488.488 0 0 1-.345-.143L8.552.435a1.488 1.488 0 0 0-2.104 0m-1.802 9.21l5-5l.708.708l-5 5zM5 5v1h1V5zm4 5h1V9H9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DistributeHorizontalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 15V0m14 15V0m-9 13.5v-12h4v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DistributeHorizontalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15V0H0v15zm14 0V0h-1v15zM10 1H5v13h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DistributeVerticalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 14.5H0m15-14H0m13.5 9h-12v-4h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DistributeVerticalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 0H0v1h15zm-1 5H1v5h13zm1 9H0v1h15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DividerLineOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 7.5h15m-12-3h7m-7-3h9m-9 9h9m-9 3h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DividerLineSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2H3V1h9zm-2 3H3V4h7zm5 3H0V7h15zm-3 3H3v-1h9zm-2 3H3v-1h7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 6.5V6H2v.5zm0 4H2v.5h.5zm10-4h.5V6h-.5zm0 4v.5h.5v-.5zm1-7h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zM2 6.5v4h1v-4zm.5 4.5h1v-1h-1zM5 9.5v-2H4v2zM3.5 6h-1v1h1zM5 7.5A1.5 1.5 0 0 0 3.5 6v1a.5.5 0 0 1 .5.5zM3.5 11A1.5 1.5 0 0 0 5 9.5H4a.5.5 0 0 1-.5.5zM6 7.5v2h1v-2zm3 2v-2H8v2zm0-2A1.5 1.5 0 0 0 7.5 6v1a.5.5 0 0 1 .5.5zM7.5 11A1.5 1.5 0 0 0 9 9.5H8a.5.5 0 0 1-.5.5zM6 9.5A1.5 1.5 0 0 0 7.5 11v-1a.5.5 0 0 1-.5-.5zm1-2a.5.5 0 0 1 .5-.5V6A1.5 1.5 0 0 0 6 7.5zM10 6v5h1V6zm.5 1h2V6h-2zm1.5-.5V8h1V6.5zM10.5 11h2v-1h-2zm2.5-.5V9h-1v1.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 10V7h.5a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5zm4-2.5a.5.5 0 0 1 1 0v2a.5.5 0 0 1-1 0z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM3.5 6H2v5h1.5A1.5 1.5 0 0 0 5 9.5v-2A1.5 1.5 0 0 0 3.5 6m4 0A1.5 1.5 0 0 0 6 7.5v2a1.5 1.5 0 0 0 3 0v-2A1.5 1.5 0 0 0 7.5 6m2.5 5V6h3v2h-1V7h-1v3h1V9h1v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockerOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 5.5V5H0v.5zm2-2V3H2v.5zm4-2V1H6v.5zm2 0H9V1h-.5zm4 6H12V8h.5zM1 7.5v-2H0v2zm2 0v-4H2v4zM2.5 4h6V3h-6zM8 3.5v4h1v-4zm-3 4v-4H4v4zm2 0v-6H6v6zM6.5 2h2V1h-2zM8 1.5v2h1v-2zm5.736 8.5H15V9h-1.264zM10 5v.5h1V5zm2 1.5v1h1v-1zm.5 1.5h1V7h-1zm1.5.5v1h1v-1zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 7zm-2-2a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 5zM3 10h1V9H3zm5.5-3h-8v1h8zM0 7.5v1h1v-1zM5.5 14h.528v-1H5.5zm.528 0a7.736 7.736 0 0 0 6.23-3.15l-.805-.593A6.737 6.737 0 0 1 6.028 13zM0 8.5A5.5 5.5 0 0 0 5.5 14v-1A4.5 4.5 0 0 1 1 8.5zM.5 6h11V5H.5zm9.5-.5A1.5 1.5 0 0 1 8.5 7v1A2.5 2.5 0 0 0 11 5.5zM13.736 9c-.96 0-1.769.558-2.283 1.257l.806.593c.383-.522.922-.85 1.477-.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockerSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 1H6v2H2v2H0v3.5A5.5 5.5 0 0 0 5.5 14h.528a7.736 7.736 0 0 0 6.774-4H15V8.5A1.5 1.5 0 0 0 13.5 7H13v-.5A1.5 1.5 0 0 0 11.5 5H9zM1 7h1V6H1zm2 0h1V6H3zm2 0h1V6H5zm2 0h1V6H7zm2 0h1V6H9zM8 3V2H7v1zM6 4H5v1h1zm1 1V4h1v1zM4 5V4H3v1zm-1 5h1V9H3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.5.5l.277-.416L10.651 0H10.5zm3 2h.5v-.268l-.223-.148zm-1 9.5h-8v1h8zM4 11.5v-10H3v10zM4.5 1h6V0h-6zM13 2.5v9h1v-9zM10.223.916l3 2l.554-.832l-3-2zM4.5 12a.5.5 0 0 1-.5-.5H3A1.5 1.5 0 0 0 4.5 13zm8 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM4 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 3 1.5zm-3 2v10h1v-10zM2.5 15h8v-1h-8zm0-12h1V2h-1zM12 13.5v-1h-1v1zM10.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zm1-10a.5.5 0 0 1 .5-.5V2A1.5 1.5 0 0 0 1 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1.5A1.5 1.5 0 0 1 4.5 0h6.151L14 2.232V11.5a1.5 1.5 0 0 1-1.5 1.5H12v.5a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 1 13.5v-10A1.5 1.5 0 0 1 2.5 2H3zM3 3h-.5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V13H4.5A1.5 1.5 0 0 1 3 11.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 10.5a3 3 0 0 0 3 3h4a3 3 0 1 0 0-6h-4a3 3 0 0 1 0-6h4a3 3 0 0 1 3 3M7.5 0v1.5m0 13.5v-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 1V0h1v1h1.5A3.5 3.5 0 0 1 13 4.5h-1A2.5 2.5 0 0 0 9.5 2h-4a2.5 2.5 0 0 0 0 5h4a3.5 3.5 0 1 1 0 7H8v1H7v-1H5.5A3.5 3.5 0 0 1 2 10.5h1A2.5 2.5 0 0 0 5.5 13h4a2.5 2.5 0 0 0 0-5h-4a3.5 3.5 0 1 1 0-7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DonutChartOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zm2.697 6.46l3.5-1.5l-.394-.92l-3.5 1.5zM7 0v4.5h1V0zm2.146 9.854l3 3l.708-.708l-3-3zM7.5 10A2.5 2.5 0 0 1 5 7.5H4A3.5 3.5 0 0 0 7.5 11zM10 7.5A2.5 2.5 0 0 1 7.5 10v1A3.5 3.5 0 0 0 11 7.5zM7.5 5A2.5 2.5 0 0 1 10 7.5h1A3.5 3.5 0 0 0 7.5 4zm0-1A3.5 3.5 0 0 0 4 7.5h1A2.5 2.5 0 0 1 7.5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DonutChartSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7.5A7.5 7.5 0 0 1 7 .016v4.02a3.5 3.5 0 1 0 2.596 6.267l2.842 2.842A7.5 7.5 0 0 1 0 7.5"/><path fill="currentColor" d="M13.145 12.438A7.471 7.471 0 0 0 15 7.5c0-1.034-.21-2.018-.587-2.914L10.755 6.21a3.498 3.498 0 0 1-.452 3.385zM8 4.035V.016a7.499 7.499 0 0 1 5.963 3.676L10.254 5.34A3.497 3.497 0 0 0 8 4.035"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretDownCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.854 8.854l.353-.354l-.707-.707l-.354.353zM7.5 10.5l-.354.354l.354.353l.354-.353zM5.854 8.146L5.5 7.793l-.707.707l.353.354zm4-2.292l.353-.354l-.707-.707l-.354.353zM7.5 7.5l-.354.354l.354.353l.354-.353zM5.854 5.146L5.5 4.793l-.707.707l.353.354zM14.5 7.5H14zm-7-7V1zm0 14v.5zm-7-7H0zm8.646.646l-2 2l.708.708l2-2zm-1.292 2l-2-2l-.708.708l2 2zm1.292-5l-2 2l.708.708l2-2zm-1.292 2l-2-2l-.708.708l2 2zM15 7.5A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5zM7.5 15A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14zM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5zm1 0A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretDownCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15M4.793 5.5L7.5 8.207L10.207 5.5L9.5 4.793l-2 2l-2-2zm0 3L7.5 11.207L10.207 8.5L9.5 7.793l-2 2l-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretDownOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m1 7l6.5 7L14 7M1 1.5l6.5 7l6.5-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretDownSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m9.5 8.5l-2 2l-2-2m4-3l-2 2l-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretDownSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.5 4.793l2 2l2-2l.707.707L7.5 8.207L4.793 5.5zm0 3l2 2l2-2l.707.707L7.5 11.207L4.793 8.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretDownSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.707 1.474L7.5 9.234L.293 1.475l.733-.68L7.5 7.764L13.974.793zm-13.68 4.82l6.473 6.97l6.474-6.972l.733.68L7.5 14.736L.293 6.974z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretLeftCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.146 9.854l.354.353l.707-.707l-.353-.354zM4.5 7.5l-.354-.354l-.353.354l.353.354zm2.354-1.646l.353-.354l-.707-.707l-.354.353zm2.292 4l.354.353l.707-.707l-.353-.354zM7.5 7.5l-.354-.354l-.353.354l.353.354zm2.354-1.646l.353-.354l-.707-.707l-.354.353zm-3 3.292l-2-2l-.708.708l2 2zm-2-1.292l2-2l-.708-.708l-2 2zm5 1.292l-2-2l-.708.708l2 2zm-2-1.292l2-2l-.708-.708l-2 2zM7.5 15A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14zM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5zM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zm0 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretLeftCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 0-15 0a7.5 7.5 0 0 0 15 0M9.5 4.793L6.793 7.5L9.5 10.207l.707-.707l-2-2l2-2zm-3 0L3.793 7.5L6.5 10.207l.707-.707l-2-2l2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretLeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M8 1L1 7.5L8 14m5.5-13l-7 6.5l7 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretLeftSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m6.5 9.5l-2-2l2-2m3 4l-2-2l2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretLeftSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.207 5.5l-2 2l2 2l-.707.707L3.793 7.5L6.5 4.793zm3 0l-2 2l2 2l-.707.707L6.793 7.5L9.5 4.793z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretLeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.707 1.026L1.735 7.5l6.972 6.474l-.68.733L.264 7.5L8.026.293zm5.5 0L7.235 7.5l6.972 6.474l-.68.733L5.764 7.5L13.526.293z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretRightCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.146 9.146l-.353.354l.707.707l.354-.353zM10.5 7.5l.354.354l.353-.354l-.353-.354zM8.854 5.146L8.5 4.793l-.707.707l.353.354zm-3.708 4l-.353.354l.707.707l.354-.353zM7.5 7.5l.354.354l.353-.354l-.353-.354zM5.854 5.146L5.5 4.793l-.707.707l.353.354zm3 4.708l2-2l-.708-.708l-2 2zm2-2.708l-2-2l-.708.708l2 2zm-5 2.708l2-2l-.708-.708l-2 2zm2-2.708l-2-2l-.708.708l2 2zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretRightCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m5.5-2.707L8.207 7.5L5.5 10.207L4.793 9.5l2-2l-2-2zm3 0L11.207 7.5L8.5 10.207L7.793 9.5l2-2l-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretRightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m7 14l7-6.5L7 1M1.5 14l7-6.5l-7-6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretRightSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m8.5 9.5l2-2l-2-2m-3 4l2-2l-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretRightSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 4.793L8.207 7.5L5.5 10.207L4.793 9.5l2-2l-2-2zm3 0L11.207 7.5L8.5 10.207L7.793 9.5l2-2l-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretRightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.474.293L9.234 7.5l-7.76 7.207l-.68-.733L7.764 7.5L.793 1.026zm5.5 0l7.76 7.207l-7.76 7.207l-.68-.733l6.97-6.474l-6.971-6.474z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretUpCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.146 6.146l-.353.354l.707.707l.354-.353zM7.5 4.5l.354-.354l-.354-.353l-.354.353zm1.646 2.354l.354.353l.707-.707l-.353-.354zm-4 2.292l-.353.354l.707.707l.354-.353zM7.5 7.5l.354-.354l-.354-.353l-.354.353zm1.646 2.354l.354.353l.707-.707l-.353-.354zM.5 7.5H0zm7 7v.5zm0-14V1zm7 7H14zm-8.646-.646l2-2l-.708-.708l-2 2zm1.292-2l2 2l.708-.708l-2-2zm-1.292 5l2-2l-.708-.708l-2 2zm1.292-2l2 2l.708-.708l-2-2zM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5zM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM15 7.5A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5zm-1 0A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretUpCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 15a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15m2.707-5.5L7.5 6.793L4.793 9.5l.707.707l2-2l2 2zm0-3L7.5 3.793L4.793 6.5l.707.707l2-2l2 2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretUpOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M14 8L7.5 1L1 8m13 5.5l-6.5-7l-6.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretUpSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m5.5 6.5l2-2l2 2m-4 3l2-2l2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretUpSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.793L10.207 6.5l-.707.707l-2-2l-2 2l-.707-.707zm0 3L10.207 9.5l-.707.707l-2-2l-2 2l-.707-.707z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCaretUpSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.5.265l7.207 7.761l-.733.68L7.5 1.736L1.026 8.707l-.733-.68zm0 5.5l7.207 7.761l-.733.68L7.5 7.236l-6.474 6.972l-.733-.68z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.854 6.146L4.5 5.793l-.707.707l.353.354zM7.5 9.5l-.354.354l.354.353l.354-.353zm3.354-2.646l.353-.354l-.707-.707l-.354.353zm-6.708 0l3 3l.708-.708l-3-3zm3.708 3l3-3l-.708-.708l-3 3zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zm-1 0A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 15a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15m4.207-9L7.5 10.207L3.293 6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m14 5l-6.5 7L1 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m4.5 6.5l3 3l3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 10.207L11.707 6H3.293z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 12L0 4h15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m7.5 10.5l-3.25-3m3.25 3l3-3m-3 3V1m6 6v6.5h-12V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 9.358V1h1v8.293l2.146-2.147l.708.708l-3.34 3.34L3.91 7.866l.678-.734zM2 13V7H1v7h13V7h-1v6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragHorizontalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-10 4a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragHorizontalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 5.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m-10 4a1 1 0 1 1 2 0a1 1 0 0 1-2 0m5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 7.5l.137-.48a.5.5 0 0 0-.618.617zm2 7l-.48.137a.5.5 0 0 0 .94.06zm5-5l.197.46a.5.5 0 0 0-.06-.94zM11 11l-.197-.46a.5.5 0 0 0-.263.263zM3.5 3.5V3H3v.5zm10 0h.5V3h-.5zm-10 10H3v.5h.5zM0 1h1V0H0zm4 0h1V0H4zm4 0h1V0H8zM0 5h1V4H0zm0 4h1V8H0zm7.02-1.363l2 7l.96-.274l-2-7zm7.617 1.382l-7-2l-.274.962l7 2zM9.96 14.697l1.5-3.5l-.92-.394l-1.5 3.5zm1.237-3.237l3.5-1.5l-.394-.92l-3.5 1.5zM3.5 4h10V3h-10zm9.5-.5V7h1V3.5zm-10 0v10h1v-10zM3.5 14H7v-1H3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1H0V0h1zm4 0H4V0h1zm4 0H8V0h1zM3 3h11v4h-1V4H4v9h3v1H3zM1 5H0V4h1zm6.146 2.146a.5.5 0 0 1 .491-.127l7 2a.5.5 0 0 1 .06.94l-3.316 1.422l-1.421 3.316a.5.5 0 0 1-.94-.06l-2-7a.5.5 0 0 1 .126-.49m1.082 1.082l1.366 4.782l.946-2.207a.5.5 0 0 1 .263-.263l2.207-.946zM1 9H0V8h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragVerticalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 3a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm-4-10a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragVerticalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m4 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m-4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m4 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m-4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m4 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.839 1.024c3.346 4.041 5.096 7.922 5.704 12.782M.533 6.82c5.985-.138 9.402-1.083 11.97-4.216M2.7 12.594c3.221-4.902 7.171-5.65 11.755-4.293M14.5 7.5a7 7 0 1 0-14 0a7 7 0 0 0 14 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.044.842A7.508 7.508 0 0 0 .092 6.32h.435c2.805-.065 5.004-.308 6.8-.858c-.78-1.383-1.732-2.74-2.874-4.12l-.001-.001zM.002 7.32L0 7.5c0 2.017.796 3.848 2.091 5.196l.161-.324a.498.498 0 0 1 .03-.052C3.94 9.798 5.816 8.298 7.914 7.625c.142-.046.286-.088.43-.126a21.803 21.803 0 0 0-.537-1.14c-1.965.633-4.327.893-7.263.96H.533z"/><path fill="currentColor" d="M2.86 13.393A7.468 7.468 0 0 0 7.5 15c.893 0 1.75-.156 2.543-.442v-.72c-.244-1.935-.673-3.71-1.318-5.404c-.17.042-.339.09-.506.143c-1.822.585-3.525 1.903-5.085 4.268zm8.183.719a7.512 7.512 0 0 0 3.802-5.086l-.565-.255c-1.626-.478-3.141-.674-4.553-.515c.638 1.72 1.067 3.526 1.312 5.488c.003.02.004.04.004.062zm3.941-6.12a7.471 7.471 0 0 0-1.805-5.39l-.297.329c-1.17 1.423-2.506 2.41-4.13 3.082c.212.424.41.851.593 1.284c1.672-.24 3.43-.014 5.251.525a.497.497 0 0 1 .065.024zm-2.508-6.104A7.472 7.472 0 0 0 7.5 0c-.878 0-1.72.15-2.503.428l.228.278c1.22 1.473 2.232 2.929 3.058 4.418c1.543-.623 2.766-1.534 3.834-2.837a.48.48 0 0 1 .015-.018z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5.5l4.286 4.142a5.737 5.737 0 0 1 1.608 2.971a5.62 5.62 0 0 1-.363 3.334a5.85 5.85 0 0 1-2.21 2.584A6.15 6.15 0 0 1 7.5 14.5a6.15 6.15 0 0 1-3.32-.97a5.85 5.85 0 0 1-2.211-2.583a5.62 5.62 0 0 1-.363-3.334a5.737 5.737 0 0 1 1.608-2.97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.847.14a.5.5 0 0 0-.694 0L2.867 4.283l-.004.003a6.237 6.237 0 0 0-1.747 3.23a6.12 6.12 0 0 0 .394 3.63a6.35 6.35 0 0 0 2.4 2.806A6.65 6.65 0 0 0 7.5 15a6.65 6.65 0 0 0 3.59-1.048a6.348 6.348 0 0 0 2.4-2.805a6.12 6.12 0 0 0 .394-3.63a6.238 6.238 0 0 0-1.747-3.23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropperOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 12.5l-.354-.354l-.146.147v.207zM8 5l.354-.354a.5.5 0 0 0-.708 0zm2 2l.354.354a.5.5 0 0 0 0-.708zm-7.5 7.5v.5h.207l.147-.146zm-2 0H0a.5.5 0 0 0 .5.5zm7-10l-.354-.354l-.353.354l.353.354zm3.086-3.086l-.354-.353zm2.828 0l-.353.354zm.172.172l.353-.354zm0 2.828l-.354-.353zM10.5 7.5l-.354.354l.354.353l.354-.353zM.854 12.854l7.5-7.5l-.708-.708l-7.5 7.5zm6.792-7.5l2 2l.708-.708l-2-2zm2 1.292l-7.5 7.5l.708.708l7.5-7.5zM2.5 14h-2v1h2zm-1.5.5v-2H0v2zM6.146 3.854l5 5l.708-.708l-5-5zm1.708 1l3.085-3.086l-.707-.707l-3.086 3.085zm5.207-3.086l.171.171l.707-.707l-.171-.171zm.171 2.293l-3.086 3.085l.708.708l3.085-3.086zm-2.378 3.085l-3-3l-.708.708l3 3zm2.378-5.207a1.5 1.5 0 0 1 0 2.122l.707.707a2.5 2.5 0 0 0 0-3.536zm-2.293-.171a1.5 1.5 0 0 1 2.122 0l.707-.707a2.5 2.5 0 0 0-3.536 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropperSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.768 1.06a2.5 2.5 0 0 0-3.536 0L7.5 3.794l-.646-.647l-.708.708l5 5l.708-.708l-.647-.646l2.732-2.732a2.5 2.5 0 0 0 0-3.536zM6.146 6.146a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1 0 .708L2.707 15H.5a.5.5 0 0 1-.5-.5v-2.207z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EdgeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.875 3.79C2.512 3.79.526 5.402.5 7.4a7 7 0 0 1 7-6.9c2.72 0 5.22 1.304 6.38 3.57c.182.361.634 1.176.62 2.401a3.357 3.357 0 0 1-1.664 2.892a3.26 3.26 0 0 1-1.656.457c-.003 0-1.486.072-2.272-.475c-.165-.115-.258-.252-.258-.395c0-.173.136-.252.18-.31c.277-.36.427-.795.427-1.088c0-.294-.06-1.031-.399-1.615M4.875 3.79c.191 0 1.283.019 2.296.55c.894.47 1.362 1.037 1.687 1.598M4.875 3.79C2.69 3.79.827 5.17.538 6.958a3.228 3.228 0 0 0-.032.451a7 7 0 0 0 9.214 6.732c-.57.18-1.172.234-1.765.159m.903-8.363c.292.503.377 1.12.395 1.47a1.769 1.769 0 0 0-1.75-1.652c-.373-.008-.657.127-.83.208l-.02.01a4.414 4.414 0 0 0-1.49 1.194M7.956 14.3a4.131 4.131 0 0 1-1.668-.596l-.005-.002c-.488-.303-.91-.7-1.244-1.167a4.414 4.414 0 0 1 .126-5.368m2.79 7.133a4.13 4.13 0 0 0 1.275-.037c.179-.038.347-.08.483-.124l.072-.024a7.018 7.018 0 0 0 3.642-2.887a.219.219 0 0 0-.291-.309a5.13 5.13 0 0 1-.577.258a5.571 5.571 0 0 1-1.963.353c-2.587 0-4.841-1.78-4.841-4.064a1.721 1.721 0 0 1 .895-1.492a4.414 4.414 0 0 0-1.486 1.193m4.091.3v.018m0 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EdgeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0A7.5 7.5 0 0 0 0 7.382v.013h.006v.012c-.01 1.13.232 2.25.709 3.275v.002a7.5 7.5 0 0 0 9.163 3.932l-.001-.004l.067-.023l.004-.001a7.518 7.518 0 0 0 3.902-3.093a.718.718 0 0 0-.95-1.017c-.167.088-.34.164-.516.23a5.07 5.07 0 0 1-1.787.322c-2.405 0-4.34-1.64-4.342-3.561a1.221 1.221 0 0 1 .621-1.048l.011-.006c.168-.079.36-.165.607-.16h.005a1.269 1.269 0 0 1 1.254 1.17a2.4 2.4 0 0 1 .004.127c0 .157-.095.479-.31.767l-.01.009c-.02.02-.06.058-.102.108a.803.803 0 0 0-.185.514c0 .375.241.644.471.805h.001c.502.35 1.176.477 1.662.53c.504.054.915.035.896.035h.002a3.757 3.757 0 0 0 1.907-.526A3.856 3.856 0 0 0 15 6.475c.015-1.286-.438-2.17-.642-2.569l-.032-.063C13.065 1.377 10.369 0 7.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 8.5l-.354-.354L4 8.293V8.5zm4-4l.354-.354a.5.5 0 0 0-.708 0zm2 2l.354.354a.5.5 0 0 0 0-.708zm-4 4v.5h.207l.147-.146zm-2 0H4a.5.5 0 0 0 .5.5zm3 3.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM4.854 8.854l4-4l-.708-.708l-4 4zm3.292-4l2 2l.708-.708l-2-2zm2 1.292l-4 4l.708.708l4-4zM6.5 10h-2v1h2zm-1.5.5v-2H4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m8.146-3.354a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1 0 .708L6.707 11H4.5a.5.5 0 0 1-.5-.5V8.293z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditOneOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 9.5l-.354-.354L0 9.293V9.5zm9-9l.354-.354a.5.5 0 0 0-.708 0zm5 5l.354.354a.5.5 0 0 0 0-.708zm-9 9v.5h.207l.147-.146zm-5 0H0a.5.5 0 0 0 .5.5zm.354-4.646l9-9l-.708-.708l-9 9zm8.292-9l5 5l.708-.708l-5-5zm5 4.292l-9 9l.708.708l9-9zM5.5 14h-5v1h5zm-4.5.5v-5H0v5zM6.146 3.854l5 5l.708-.708l-5-5zM8 15h7v-1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditOneSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.854.146a.5.5 0 0 0-.708 0L6.5 2.793L12.207 8.5l2.647-2.646a.5.5 0 0 0 0-.708zM0 9.293L5.793 3.5L11.5 9.207L5.707 15H.5a.5.5 0 0 1-.5-.5zM8 15h7v-1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 10.5l-.354-.354l-.146.147v.207zm10-10l.354-.354a.5.5 0 0 0-.708 0zm4 4l.354.354a.5.5 0 0 0 0-.708zm-10 10v.5h.207l.147-.146zm-4 0H0a.5.5 0 0 0 .5.5zm.354-3.646l10-10l-.708-.708l-10 10zm9.292-10l4 4l.708-.708l-4-4zm4 3.292l-10 10l.708.708l10-10zM4.5 14h-4v1h4zm-3.5.5v-4H0v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 8.5l-.354-.354L4 8.293V8.5zm4-4l.354-.354a.5.5 0 0 0-.708 0zm2 2l.354.354a.5.5 0 0 0 0-.708zm-4 4v.5h.207l.147-.146zm-2 0H4a.5.5 0 0 0 .5.5zm.354-1.646l4-4l-.708-.708l-4 4zm3.292-4l2 2l.708-.708l-2-2zm2 1.292l-4 4l.708.708l4-4zM6.5 10h-2v1h2zm-1.5.5v-2H4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.854 4.146a.5.5 0 0 0-.708 0L4 8.293V10.5a.5.5 0 0 0 .5.5h2.207l4.147-4.146a.5.5 0 0 0 0-.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.854.146a.5.5 0 0 0-.708 0L0 10.293V14.5a.5.5 0 0 0 .5.5h4.207L14.854 4.854a.5.5 0 0 0 0-.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElbowConnectorOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 1.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 0h5v12h5m0 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElbowConnectorSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0a1.5 1.5 0 1 0 1.415 2H7v12h5.085a1.5 1.5 0 1 0 0-1H8V1H2.915A1.5 1.5 0 0 0 1.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5 5l7 3.5l7-3.5m0 .08v8.42a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1V5.08a1 1 0 0 1 .504-.868l6-3.428a1 1 0 0 1 .992 0l6 3.428a1 1 0 0 1 .504.868Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.756.35a1.5 1.5 0 0 1 1.488 0l6 3.428a1.5 1.5 0 0 1 .408.341L7.5 7.933L.348 4.12c.113-.135.25-.251.408-.341z"/><path fill="currentColor" d="M0 5.067V13.5A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5V5.067l-7.5 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5 4.5l7 4l7-4m-13-3h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 2.5A1.5 1.5 0 0 1 2 1h12a1.5 1.5 0 0 1 1.5 1.5v1.208L8 7.926L.5 3.708z"/><path fill="currentColor" d="M.5 4.855V12.5A1.5 1.5 0 0 0 2 14h12a1.5 1.5 0 0 0 1.5-1.5V4.855L8 9.074z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EpsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm0 6V6H10v.5zm0 2H10V9h.5zm2 0h.5V8h-.5zm0 2v.5h.5v-.5zm-6-4V6H6v.5zm2 0H9V6h-.5zm0 2V9H9v-.5zm-6-2V6H2v.5zm0 4H2v.5h.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zM13 6h-2.5v1H13zm-3 .5v2h1v-2zm.5 2.5h2V8h-2zm1.5-.5v2h1v-2zm.5 1.5H10v1h2.5zm-6-3h2V6h-2zM8 6.5v2h1v-2zM7 11V8.5H6V11zm0-2.5v-2H6v2zM8.5 8h-2v1h2zM5 6H2.5v1H5zm-3 .5v4h1v-4zm.5 4.5H5v-1H2.5zm0-2H5V8H2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EpsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8h1V7H7z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM5 6H2v5h3v-1H3V9h2V8H3V7h2zm1 0h3v3H7v2H6zm4 0h3v1h-2v1h2v3h-3v-1h2V9h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EslintOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 7.5l-.447-.224a.5.5 0 0 0 0 .448zm3-6V1a.5.5 0 0 0-.447.276zm8 0l.447-.224A.5.5 0 0 0 11.5 1zm3 6l.447.224a.5.5 0 0 0 0-.448zm-3 6v.5a.5.5 0 0 0 .447-.276zm-8 0l-.447.224A.5.5 0 0 0 3.5 14zm4-9.5l.277-.416l-.277-.185l-.277.185zm-3 2l-.277-.416L4 5.732V6zm0 3H4v.268l.223.148zm3 2l-.277.416l.277.185l.277-.185zm3-2l.277.416l.223-.148V9zm0-3h.5v-.268l-.223-.148zM.947 7.724l3-6l-.894-.448l-3 6zM3.5 2h8V1h-8zm7.553-.276l3 6l.894-.448l-3-6zm3 5.552l-3 6l.894.448l3-6zM11.5 13h-8v1h8zm-7.553.276l-3-6l-.894.448l3 6zm3.276-9.692l-3 2l.554.832l3-2zM4 6v3h1V6zm.223 3.416l3 2l.554-.832l-3-2zm3.554 2l3-2l-.554-.832l-3 2zM11 9V6h-1v3zm-.223-3.416l-3-2l-.554.832l3 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EslintSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8.732V6.268L7.5 4.6L10 6.268v2.464L7.5 10.4z"/><path fill="currentColor" fill-rule="evenodd" d="M3.053 1.276A.5.5 0 0 1 3.5 1h8a.5.5 0 0 1 .447.276l3 6a.5.5 0 0 1 0 .448l-3 6A.5.5 0 0 1 11.5 14h-8a.5.5 0 0 1-.447-.276l-3-6a.5.5 0 0 1 0-.448zM11 5.732L7.5 3.4L4 5.732v3.536L7.5 11.6L11 9.268z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EthereumOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.384-.32a.5.5 0 0 0-.768 0zm-5 6l-.384-.32a.5.5 0 0 0-.04.585zm5 8l-.424.265a.5.5 0 0 0 .848 0zm5-8l.424.265a.5.5 0 0 0-.04-.585zm-5-2l.186-.464L7.5 3.96l-.186.075zM7.116.18l-5 6l.768.64l5-6zm-5.04 6.585l5 8l.848-.53l-5-8zm5.848 8l5-8l-.848-.53l-5 8zm4.96-8.585l-5-6l-.768.64l5 6zm-10.198.784l5-2l-.372-.928l-5 2zm4.628-2l5 2l.372-.928l-5-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EthereumSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a.5.5 0 0 1 .384.18l5 6a.5.5 0 0 1 .04.585l-5 8a.5.5 0 0 1-.848 0l-5-8a.5.5 0 0 1 .04-.585l5-6A.5.5 0 0 1 7.5 0M3.241 6.742L7.5 5.04l4.259 1.703L7.5 13.557zm7.61-1.44L7.5 3.962l-3.35 1.34L7.5 1.28z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.374 3A6 6 0 0 0 2.5 6.5v2A6 6 0 0 0 13.374 12M0 5.5h7m-7 4h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.174 5A6.503 6.503 0 0 1 13.78 2.708l-.812.584A5.502 5.502 0 0 0 3.207 5H7v1H3.022A5.57 5.57 0 0 0 3 6.5v2c0 .169.008.335.022.5H7v1H3.207a5.502 5.502 0 0 0 9.761 1.708l.812.584A6.503 6.503 0 0 1 2.174 10H0V9h2.019A6.593 6.593 0 0 1 2 8.5v-2c0-.168.006-.335.019-.5H0V5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 10.5V10H7v.5zm-1 .01v.5h1v-.5zM7 4v4h1V4zm0 6.5v.01h1v-.01zm.5 3.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M7 8V4h1v4zm1 2v1.01H7V10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 13.5V13H7v.5zm-1 .01v.5h1v-.5zM7 1v10h1V1zm0 12.5v.01h1v-.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 10.5V10H7v.5zm-1 .01v.5h1v-.5zM7 4v4h1V4zm0 6.5v.01h1v-.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 8V4h1v4zm1 2v1.01H7V10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 11V1h1v10zm1 2v1.01H7V13z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 1.5h3.5m0 0V5m0-3.5l-4 4m-8 4.5v3.5m0 0H5m-3.5 0l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.293 2H10V1h4v4h-1V2.707L9.854 5.854l-.708-.708zm-6.44 7.854L2.708 13H5v1H1v-4h1v2.293l3.146-3.147z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.5 13.5H10m3.5 0V10m0 3.5l-4-4m.5-8h3.5m0 0V5m0-3.5l-4 4M5 1.5H1.5m0 0V5m0-3.5l4 4m-4 4.5v3.5m0 0H5m-3.5 0l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h4v1H2.707l3.147 3.146l-.708.708L2 2.707V5H1zm11.293 1H10V1h4v4h-1V2.707L9.854 5.854l-.708-.708zm-6.44 7.854L2.708 13H5v1H1v-4h1v2.293l3.146-3.147zm4-.708L13 12.293V10h1v4h-4v-1h2.293L9.146 9.854z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosedOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 9C5.186 9 3.561 7.848 2.497 6.666a9.368 9.368 0 0 1-1.449-2.164a5.065 5.065 0 0 1-.08-.18l-.004-.007v-.001L.5 4.5l-.464.186v.002l.003.004a2.107 2.107 0 0 0 .026.063l.078.173a10.367 10.367 0 0 0 1.61 2.406C2.94 8.652 4.814 10 7.5 10zm7-4.5a68.887 68.887 0 0 1-.464-.186l-.003.008l-.015.035l-.066.145a9.37 9.37 0 0 1-1.449 2.164C11.44 7.848 9.814 9 7.5 9v1c2.686 0 4.561-1.348 5.747-2.666a10.365 10.365 0 0 0 1.61-2.406a6.164 6.164 0 0 0 .104-.236l.002-.004v-.001h.001zM8 12V9.5H7V12zm-6.646-1.646l2-2l-.708-.708l-2 2zm10.292-2l2 2l.708-.708l-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosedSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.497 6.666C3.56 7.848 5.186 9 7.5 9s3.939-1.152 5.003-2.334a9.37 9.37 0 0 0 1.449-2.164a4.967 4.967 0 0 0 .08-.18l.004-.007v-.001l.464.186l.464.186v.002l-.003.004l-.005.014a3.334 3.334 0 0 1-.1.222a10.37 10.37 0 0 1-1.61 2.406a9.204 9.204 0 0 1-.598.607l1.706 1.705l-.708.708l-1.774-1.775A7.304 7.304 0 0 1 8 9.984V12H7V9.984A7.304 7.304 0 0 1 3.128 8.58l-1.774 1.775l-.708-.708l1.706-1.705a9.237 9.237 0 0 1-.599-.607a10.367 10.367 0 0 1-1.61-2.406a6.064 6.064 0 0 1-.099-.222L.04 4.692l-.002-.004v-.001H.035L.5 4.5l.464-.186l.004.008a2.633 2.633 0 0 0 .08.18a9.368 9.368 0 0 0 1.449 2.164M.964 4.314" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 7.5l-.464-.186a.5.5 0 0 0 0 .372zm14 0l.464.186a.5.5 0 0 0 0-.372zm-7 4.5c-2.314 0-3.939-1.152-5.003-2.334a9.368 9.368 0 0 1-1.449-2.164a5.065 5.065 0 0 1-.08-.18l-.004-.007v-.001L.5 7.5l-.464.186v.002l.003.004a2.107 2.107 0 0 0 .026.063l.078.173a10.368 10.368 0 0 0 1.61 2.406C2.94 11.653 4.814 13 7.5 13zm-7-4.5l.464.186l.004-.008a2.62 2.62 0 0 1 .08-.18a9.368 9.368 0 0 1 1.449-2.164C3.56 4.152 5.186 3 7.5 3V2C4.814 2 2.939 3.348 1.753 4.666a10.367 10.367 0 0 0-1.61 2.406a6.05 6.05 0 0 0-.104.236l-.002.004v.001H.035zm7-4.5c2.314 0 3.939 1.152 5.003 2.334a9.37 9.37 0 0 1 1.449 2.164a4.705 4.705 0 0 1 .08.18l.004.007v.001L14.5 7.5l.464-.186v-.002l-.003-.004a.656.656 0 0 0-.026-.063a9.094 9.094 0 0 0-.39-.773a10.365 10.365 0 0 0-1.298-1.806C12.06 3.348 10.186 2 7.5 2zm7 4.5a68.887 68.887 0 0 1-.464-.186l-.003.008l-.015.035l-.066.145a9.37 9.37 0 0 1-1.449 2.164C11.44 10.848 9.814 12 7.5 12v1c2.686 0 4.561-1.348 5.747-2.665a10.366 10.366 0 0 0 1.61-2.407a6.164 6.164 0 0 0 .104-.236l.002-.004v-.001h.001zM7.5 9A1.5 1.5 0 0 1 6 7.5H5A2.5 2.5 0 0 0 7.5 10zM9 7.5A1.5 1.5 0 0 1 7.5 9v1A2.5 2.5 0 0 0 10 7.5zM7.5 6A1.5 1.5 0 0 1 9 7.5h1A2.5 2.5 0 0 0 7.5 5zm0-1A2.5 2.5 0 0 0 5 7.5h1A1.5 1.5 0 0 1 7.5 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 7.686V9a1.5 1.5 0 0 1 0-3zM7.5 5a2.5 2.5 0 0 0 0 5v3c-2.686 0-4.561-1.348-5.747-2.665a10.368 10.368 0 0 1-1.61-2.407a6.05 6.05 0 0 1-.099-.222l-.006-.014l-.001-.004l-.001-.002L.5 7.5l-.464.186a.5.5 0 0 1 0-.372l.066.027a1.304 1.304 0 0 1-.066-.028v-.001l.002-.004l.006-.014a3.62 3.62 0 0 1 .1-.222a10.367 10.367 0 0 1 1.61-2.406C2.938 3.348 4.813 2 7.5 2zm0 1v3a1.5 1.5 0 1 0 0-3m0 4a2.5 2.5 0 0 0 0-5V2c2.686 0 4.561 1.348 5.747 2.666a10.365 10.365 0 0 1 1.61 2.406a6.164 6.164 0 0 1 .099.222l.005.014l.002.004l.001.002l-.364.146l.364-.146a.5.5 0 0 1 0 .372L14.5 7.5l.464.187v.001l-.003.004l-.005.014a3.334 3.334 0 0 1-.1.222a10.366 10.366 0 0 1-1.61 2.406C12.062 11.653 10.187 13 7.5 13z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceIdOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h1V5H4zm6 0h1V5h-1zm.1 2.7a3.25 3.25 0 0 1-5.2 0l-.8.6c1.7 2.267 5.1 2.267 6.8 0zM1 5V2.5H0V5zm1.5-4H5V0H2.5zM1 2.5A1.5 1.5 0 0 1 2.5 1V0A2.5 2.5 0 0 0 0 2.5zM0 10v2.5h1V10zm2.5 5H5v-1H2.5zM0 12.5A2.5 2.5 0 0 0 2.5 15v-1A1.5 1.5 0 0 1 1 12.5zM10 1h2.5V0H10zm4 1.5V5h1V2.5zM12.5 1A1.5 1.5 0 0 1 14 2.5h1A2.5 2.5 0 0 0 12.5 0zM10 15h2.5v-1H10zm5-2.5V10h-1v2.5zM12.5 15a2.5 2.5 0 0 0 2.5-2.5h-1a1.5 1.5 0 0 1-1.5 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceIdSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 1A1.5 1.5 0 0 0 1 2.5V5H0V2.5A2.5 2.5 0 0 1 2.5 0H5v1zm10 0H10V0h2.5A2.5 2.5 0 0 1 15 2.5V5h-1V2.5A1.5 1.5 0 0 0 12.5 1M5 6H4V5h1zm6 0h-1V5h1zM4.9 8.7a3.25 3.25 0 0 0 5.2 0l.8.6c-1.7 2.267-5.1 2.267-6.8 0zM0 12.5V10h1v2.5A1.5 1.5 0 0 0 2.5 14H5v1H2.5A2.5 2.5 0 0 1 0 12.5M15 10v2.5a2.5 2.5 0 0 1-2.5 2.5H10v-1h2.5a1.5 1.5 0 0 0 1.5-1.5V10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 14.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm0 0v-8a2 2 0 0 1 2-2h.5m-5 4h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7.5a7.5 7.5 0 1 1 8 7.484V9h2V8H8V6.5A1.5 1.5 0 0 1 9.5 5h.5V4h-.5A2.5 2.5 0 0 0 7 6.5V8H5v1h2v5.984A7.5 7.5 0 0 1 0 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FigmaOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 1v8.5m0 0v2a2 2 0 1 1-2-2m2 0h-2m0 0a2 2 0 1 1 0-4m0 0h2m-2 0h4m-4 0a2 2 0 1 1 0-4h4a2 2 0 1 1 0 4m0 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FigmaSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 9.5a2.496 2.496 0 0 1-1-2c0-.818.393-1.544 1-2A2.5 2.5 0 0 1 5.5 1h4A2.5 2.5 0 0 1 11 5.5a2.5 2.5 0 0 1-3 4v2a2.5 2.5 0 1 1-4-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5zm3 3h.5v-.207l-.146-.147zm-1 10.5h-10v1h10zM2 13.5v-12H1v12zM2.5 1h8V0h-8zM13 3.5v10h1v-10zM10.146.854l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM5 8h5V7H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM5 8h5V7H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileNoAccessOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m9.5 5.5l-4 4m5-9h-8a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-10zm-3 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileNoAccessSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 5a2.5 2.5 0 0 0-2.086 3.879L8.88 5.414A2.488 2.488 0 0 0 7.5 5m0 5c-.51 0-.983-.152-1.379-.414L9.586 6.12A2.5 2.5 0 0 1 7.5 10"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zm3 6a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5zm3 3h.5v-.207l-.146-.147zm-1 10.5h-10v1h10zM2 13.5v-12H1v12zM2.5 1h8V0h-8zM13 3.5v10h1v-10zM10.146.854l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5zm3 3h.5v-.207l-.146-.147zm-1 10.5h-10v1h10zM2 13.5v-12H1v12zM2.5 1h8V0h-8zM13 3.5v10h1v-10zM10.146.854l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM7 5v5h1V5zM5 8h5V7H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM7 10V8H5V7h2V5h1v2h2v1H8v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 0A1.5 1.5 0 0 0 1 1.5v12A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V3.293L10.707 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTickOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m5 7.5l2 2l3.5-4m0-5h-8a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTickSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zm6.024 8.732l3.852-4.403l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileXoutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5zm3 3h.5v-.207l-.146-.147zm-1 10.5h-10v1h10zM2 13.5v-12H1v12zM2.5 1h8V0h-8zM13 3.5v10h1v-10zM10.146.854l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zm3.146 4.354l4 4l.708-.708l-4-4zm.708 4l4-4l-.708-.708l-4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileXsolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zm8.146 8.354L7.5 8.207L5.854 9.854l-.708-.708L6.793 7.5L5.146 5.854l.708-.708L7.5 6.793l1.646-1.647l.708.708L8.207 7.5l1.647 1.646z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 2.5h15m-12 5h9m-7 5h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 3H0V2h15zm-3 5H3V7h9zm-2 5H5v-1h5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" d="M12.587 3.513a6.03 6.03 0 0 1 .818 3.745v.75c0 .788.205 1.563.595 2.247M4.483 6.509c0-.795.313-1.557.871-2.119a2.963 2.963 0 0 1 2.103-.877c.789 0 1.545.315 2.103.877c.558.562.871 1.324.871 2.12v.748c0 1.621.522 3.198 1.487 4.495m-4.46-5.244v1.498A10.542 10.542 0 0 0 9.315 14M4.483 9.505A13.559 13.559 0 0 0 5.821 14m-3.643-1.498a16.63 16.63 0 0 1-.669-5.244V6.51a6.028 6.028 0 0 1 .79-3.002a5.97 5.97 0 0 1 2.177-2.2a5.914 5.914 0 0 1 5.955-.004"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.18 1.735a5.414 5.414 0 0 0-5.452.004a5.471 5.471 0 0 0-1.994 2.016a5.528 5.528 0 0 0-.725 2.753v.767a16.07 16.07 0 0 0 .649 5.085l.141.48l-.959.283l-.141-.48a17.096 17.096 0 0 1-.69-5.393v-.74a6.528 6.528 0 0 1 .857-3.25A6.47 6.47 0 0 1 4.224.874A6.414 6.414 0 0 1 10.683.87l.432.251l-.503.864zm2.58 1.092l.257.43a6.53 6.53 0 0 1 .888 4.028v.723c0 .701.182 1.39.53 1.999l.247.434l-.868.496l-.248-.434a5.02 5.02 0 0 1-.66-2.496v-.749c0-.018 0-.036.002-.054a5.53 5.53 0 0 0-.75-3.435l-.256-.429zM7.457 4.013c-.655 0-1.284.262-1.748.73a2.508 2.508 0 0 0-.726 1.766v.5h-1v-.5c0-.926.365-1.815 1.016-2.47a3.463 3.463 0 0 1 2.458-1.026c.923 0 1.807.369 2.458 1.025A3.508 3.508 0 0 1 10.93 6.51v.75c0 1.513.487 2.985 1.388 4.195l.299.401l-.802.597l-.299-.4A8.028 8.028 0 0 1 9.93 7.259V6.51c0-.663-.261-1.299-.726-1.766a2.463 2.463 0 0 0-1.748-.73m.5 3.995a10.042 10.042 0 0 0 1.77 5.708l.284.412l-.823.567l-.284-.411a11.042 11.042 0 0 1-1.947-6.277m-2.035.944l.058.497a13.058 13.058 0 0 0 1.289 4.329l.223.447l-.894.447l-.224-.447a14.06 14.06 0 0 1-1.388-4.66l-.057-.497z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirebaseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m2.5 11.5l9-8l1 9l-5 2zm0 0l5-9l2 3m-7 6l1-11l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirebaseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.33.03a.5.5 0 0 1 .524.116l2.078 2.08a.505.505 0 0 0-.032.056L2.175 9.988zM2.262 11.94l4.98 2.989a.5.5 0 0 0 .444.035l5-2a.5.5 0 0 0 .31-.52l-1-9a.5.5 0 0 0-.828-.318L9.513 4.597zm6.676-8.184L7.916 2.223a.5.5 0 0 0-.853.034l-.31.558l-2.986 6.177z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.5 6.5l.224.447a.5.5 0 0 0 .033-.876zm-10-6l.257-.429A.5.5 0 0 0 2 .5zm10.257 5.571l-10-6l-.514.858l10 6zM2 .5v11h1V.5zm.724 11.447l10-5l-.448-.894l-10 5zM3 15v-3.5H2V15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.254.065a.5.5 0 0 1 .503.006l10 6a.5.5 0 0 1-.033.876L3 11.81V15H2V.5a.5.5 0 0 1 .254-.435"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.5.5l.457.203A.5.5 0 0 0 14.5 0zM.5.5V0a.5.5 0 0 0-.5.5zm14 9v.5a.5.5 0 0 0 .457-.703zm-2-4.5l-.457-.203a.5.5 0 0 0 0 .406zm2-5H.5v1h14zM0 .5v9h1v-9zM.5 10h14V9H.5zm14.457-.703l-2-4.5l-.914.406l2 4.5zm-2-4.094l2-4.5l-.914-.406l-2 4.5zM1 15V9.5H0V15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 .5A.5.5 0 0 1 .5 0h14a.5.5 0 0 1 .457.703L13.047 5l1.91 4.297A.5.5 0 0 1 14.5 10H1v5H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipHorizontalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5.5l.224-.447A.5.5 0 0 0 3 .5zm8 4V5a.5.5 0 0 0 .224-.947zm-8 0H3a.5.5 0 0 0 .5.5zm0 6V10a.5.5 0 0 0-.5.5zm8 0l.224.447A.5.5 0 0 0 11.5 10zm-8 4H3a.5.5 0 0 0 .724.447zM3.276.947l8 4l.448-.894l-8-4zM11.5 4h-8v1h8zM4 4.5v-4H3v4zM0 8h15V7H0zm3.5 3h8v-1h-8zm7.776-.947l-8 4l.448.894l8-4zM4 14.5v-4H3v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipHorizontalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.237.075a.5.5 0 0 1 .487-.022l8 4A.5.5 0 0 1 11.5 5h-8a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .237-.425M0 8h15V7H0zm3.5 2a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .724.447l8-4A.5.5 0 0 0 11.5 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipVerticalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 3.5H5a.5.5 0 0 0-.947-.224zm0 8v.5a.5.5 0 0 0 .5-.5zm-4 0l-.447-.224A.5.5 0 0 0 .5 12zm10-8l.447-.224A.5.5 0 0 0 10 3.5zm0 8H10a.5.5 0 0 0 .5.5zm4 0v.5a.5.5 0 0 0 .447-.724zM4 3.5v8h1v-8zm.5 7.5h-4v1h4zm-3.553.724l4-8l-.894-.448l-4 8zM10 3.5v8h1v-8zm.5 8.5h4v-1h-4zm4.447-.724l-4-8l-.894.448l4 8zM7 0v15h1V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipVerticalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0v15h1V0zM4.615 3.013A.5.5 0 0 1 5 3.5v8a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.447-.724l4-8a.5.5 0 0 1 .562-.263m5.77 0a.5.5 0 0 1 .562.263l4 8A.5.5 0 0 1 14.5 12h-4a.5.5 0 0 1-.5-.5v-8a.5.5 0 0 1 .385-.487"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloatCenterOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 5.5h2m11 0h2m-15-4h2m11 0h2m-15 8h15m-15 4h15M5.5.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloatCenterSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 0A1.5 1.5 0 0 0 4 1.5v4A1.5 1.5 0 0 0 5.5 7h4A1.5 1.5 0 0 0 11 5.5v-4A1.5 1.5 0 0 0 9.5 0zM0 2h2V1H0zm13 0h2V1h-2zM0 6h2V5H0zm13 0h2V5h-2zM0 10h15V9H0zm0 4h15v-1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloatLeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9 5.5h6m-6-4h6m-15 8h15m-15 4h15M1.5.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloatLeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v4A1.5 1.5 0 0 0 1.5 7h4A1.5 1.5 0 0 0 7 5.5v-4A1.5 1.5 0 0 0 5.5 0zM9 2h6V1H9zm0 4h6V5H9zm-9 4h15V9H0zm0 4h15v-1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloatRightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 5.5h6m-6-4h6m-6 8h15m-15 4h15M9.5.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloatRightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 0A1.5 1.5 0 0 0 8 1.5v4A1.5 1.5 0 0 0 9.5 7h4A1.5 1.5 0 0 0 15 5.5v-4A1.5 1.5 0 0 0 13.5 0zM0 2h6V1H0zm0 4h6V5H0zm0 4h15V9H0zm0 4h15v-1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloorplanOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 .5h4.5v14H.5V.5h4l3 2m-1 12v-7M4 7.5h5m3 0h2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloorplanSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 0h4.651l3.126 2.084l-.554.832L4.349 1H1v13h5V8H4V7h5v1H7v6h7V8h-2V7h2V1h-4V0h5v15H0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 8.5h5m-9.5-6v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5zM5 9h5V8H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderNoAccessOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m9.5 6.5l-4 4m-5-8v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Zm7 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderNoAccessSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 6a2.5 2.5 0 0 0-2.086 3.879L8.88 6.414A2.488 2.488 0 0 0 7.5 6m0 5c-.51 0-.983-.152-1.379-.414L9.586 7.12A2.5 2.5 0 0 1 7.5 11"/><path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5zm4 6a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 12.5v-10a1 1 0 0 1 1-1h4l2 2h6a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 6v5M5 8.5h5m-9.5-6v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5zM7 11V9H5V8h2V6h1v2h2v1H8v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 1A1.5 1.5 0 0 0 0 2.5v10A1.5 1.5 0 0 0 1.5 14h12a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 3H7.707l-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTickOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m5 8.5l2 2l3.5-4m-10-4v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTickSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5zm7.024 8.732l3.852-4.403l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderXoutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m5.5 6.5l4 4m-4 0l4-4m-9-4v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-6l-2-2h-4a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderXsolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h4.207l2 2H13.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5zm9.146 8.354L7.5 9.207l-1.646 1.647l-.708-.707L6.793 8.5L5.146 6.854l.708-.708L7.5 7.793l1.646-1.647l.708.708L8.207 8.5l1.647 1.646z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FoldersOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 8.5v-7a1 1 0 0 1 1-1h3l2 2h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-9m-1-1a1 1 0 0 0 1 1m-1-1v-3h-2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-4h-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FoldersSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 0A1.5 1.5 0 0 0 3 1.5v7A1.5 1.5 0 0 0 4.5 10h9A1.5 1.5 0 0 0 15 8.5v-5A1.5 1.5 0 0 0 13.5 2H9.707l-2-2z"/><path fill="currentColor" d="M12 11H4.5A2.5 2.5 0 0 1 2 8.5V5h-.5A1.5 1.5 0 0 0 0 6.5v7A1.5 1.5 0 0 0 1.5 15h9a1.5 1.5 0 0 0 1.5-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.5 5.5l.248-.434A.5.5 0 0 0 8 5.5zm0 4H8a.5.5 0 0 0 .748.434zm3.5-2l.248.434a.5.5 0 0 0 0-.868zm-7.5-2l.248-.434A.5.5 0 0 0 4 5.5zm0 4H4a.5.5 0 0 0 .748.434zm3.5-2l.248.434a.5.5 0 0 0 0-.868zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM8 5.5v4h1v-4zm.748 4.434l3.5-2l-.496-.868l-3.5 2zm3.5-2.868l-3.5-2l-.496.868l3.5 2zM4 5.5v4h1v-4zm.748 4.434l3.5-2l-.496-.868l-3.5 2zm3.5-2.868l-3.5-2l-.496.868l3.5 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m4.249-2.432a.5.5 0 0 1 .5-.002L8 6.924V5.5a.5.5 0 0 1 .748-.434l3.5 2a.5.5 0 0 1 0 .868l-3.5 2A.5.5 0 0 1 8 9.5V8.076L4.748 9.934A.5.5 0 0 1 4 9.5v-4a.5.5 0 0 1 .249-.432" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M.5 12.5v-10l7 5zm7 0v-10l7 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M8.5 9.5v-4l3.5 2zm-4 0v-4l3.5 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.748 5.066A.5.5 0 0 0 4 5.5v4a.5.5 0 0 0 .748.434L8 8.076V9.5a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868l-3.5-2A.5.5 0 0 0 8 5.5v1.424z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.79 2.093A.5.5 0 0 0 0 2.5v10a.5.5 0 0 0 .79.407L7 8.472V12.5a.5.5 0 0 0 .79.407l7-5a.5.5 0 0 0 0-.814l-7-5A.5.5 0 0 0 7 2.5v4.028z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v15m8-15v15M0 3.5h15m-15 8h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v15m8-15v15M0 3.5h15m-15 8h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FramerOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M7.5 5.5h5v-5h-10zm0 0h-5v4l5 5v-4h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FramerSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.038.309A.5.5 0 0 1 2.5 0h10a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-.5.5H8.707l4.147 4.146A.5.5 0 0 1 12.5 11H8v3.5a.5.5 0 0 1-.854.354l-5-5A.5.5 0 0 1 2 9.5v-4a.5.5 0 0 1 .5-.5h3.793L2.146.854a.5.5 0 0 1-.108-.545"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GameControllerOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.817 11.133l-.447.224zM9.5 10.5l.447-.224A.5.5 0 0 0 9.5 10zm-4 0V10a.5.5 0 0 0-.447.276zm8.5-5v4.528h1V5.5zm-3.736 5.41l-.317-.634l-.894.448l.316.633zM9.5 10h-4v1h4zm-4.447.276l-.317.634l.894.447l.317-.633zM1 10.028V5.5H0v4.528zM3.5 3h8V2h-8zm-.528 9A1.972 1.972 0 0 1 1 10.028H0A2.972 2.972 0 0 0 2.972 13zm9.056 0c-.747 0-1.43-.422-1.764-1.09l-.894.447A2.972 2.972 0 0 0 12.028 13zM14 10.028A1.972 1.972 0 0 1 12.028 12v1A2.972 2.972 0 0 0 15 10.028zm-9.264.882A1.972 1.972 0 0 1 2.972 12v1a2.972 2.972 0 0 0 2.658-1.643zM15 5.5A3.5 3.5 0 0 0 11.5 2v1A2.5 2.5 0 0 1 14 5.5zm-14 0A2.5 2.5 0 0 1 3.5 3V2A3.5 3.5 0 0 0 0 5.5zM3 7h3V6H3zm1-2v3h1V5zm7 1h1V5h-1zM9 8h1V7H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GameControllerRetroOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8 5.5h2M4.5 6v3M3 7.5h3m4 2h2M.5 3.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GameControllerRetroSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5v-8A1.5 1.5 0 0 1 1.5 2zM10 6H8V5h2zM4 7V6h1v1h1v1H5v1H4V8H3V7zm8 3h-2V9h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GameControllerSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 5.5A3.5 3.5 0 0 1 3.5 2h8A3.5 3.5 0 0 1 15 5.5v4.528a2.972 2.972 0 0 1-5.63 1.329L9.19 11H5.809l-.179.357A2.972 2.972 0 0 1 0 10.027zM4 8V7H3V6h1V5h1v1h1v1H5v1zm6 0H9V7h1zm1-2h1V5h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GanttChartOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 0v14.5H15M5 2.5H2m6 3H3m5 3H5m10 3H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GanttChartSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 0h1v14h14v1H0zm2 2h3v1H2zm1 3h5v1H3zm2 3h3v1H5zm3 3h7v1H8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GarageOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 5.5l-.29-.407A.5.5 0 0 0 0 5.5zm7-5l.29-.407a.5.5 0 0 0-.58 0zm7 5h.5a.5.5 0 0 0-.21-.407zm-12 2V7H2v.5zm10 0h.5V7h-.5zM1 15V5.5H0V15zM.79 5.907l7-5l-.58-.814l-7 5zm6.42-5l7 5l.58-.814l-7-5zM14 5.5V15h1V5.5zM3 15V7.5H2V15zm-.5-7h10V7h-10zm9.5-.5V15h1V7.5zM2.5 11h10v-1h-10zM6 13h3v-1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GarageSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.21.093a.5.5 0 0 1 .58 0l7 5A.5.5 0 0 1 15 5.5v9a.5.5 0 0 1-.5.5H13V7H2v8H.5a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .21-.407z"/><path fill="currentColor" fill-rule="evenodd" d="M3 15h9v-4H3zm6-2H6v-1h3z" clip-rule="evenodd"/><path fill="currentColor" d="M12 10V8H3v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GatsbyjsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.07 4a5.002 5.002 0 0 0-8.342 2L9 12.271A5.004 5.004 0 0 0 12.475 8H9m-6.5.5l4 4m1 2a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GatsbyjsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m3.305-1.631a4.502 4.502 0 0 1 7.409-1.519l.714-.7a5.502 5.502 0 0 0-9.176 2.2l-.09.29l6.699 6.699l.289-.09a5.504 5.504 0 0 0 3.823-4.7l.054-.549H9v1h2.889a4.51 4.51 0 0 1-2.758 3.195zm2.841 6.985l-4-4l.708-.708l4 4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GbaOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 6.5h3m7 0h1m-2 2h1M3.5 5v3m1.606-5.5h4.788a2 2 0 0 1 1.11.336L12 3.5h.138a2 2 0 0 1 1.995 1.858l.32 4.475a1 1 0 0 1-.55.966l-2.091 1.045a3.175 3.175 0 0 1-.65.24a15.097 15.097 0 0 1-7.324 0a3.176 3.176 0 0 1-.65-.24l-2.09-1.045a1 1 0 0 1-.55-.966l.32-4.475A2 2 0 0 1 2.861 3.5H3l.996-.664a2 2 0 0 1 1.11-.336Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GbaSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.719 2.42A2.5 2.5 0 0 1 5.106 2h4.789a2.5 2.5 0 0 1 1.386.42l.87.58a2.5 2.5 0 0 1 2.48 2.322l.32 4.476a1.5 1.5 0 0 1-.825 1.448l-2.09 1.045a3.68 3.68 0 0 1-.753.279a15.62 15.62 0 0 1-7.566 0a3.68 3.68 0 0 1-.753-.279l-2.09-1.045A1.5 1.5 0 0 1 .05 9.798l.32-4.476A2.5 2.5 0 0 1 2.849 3zM3 8V7H2V6h1V5h1v1h1v1H4v1zm10-1h-1V6h1zm-2 2h1V8h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GbcOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 2.5V2a.5.5 0 0 0-.5.5zm6 0h.5a.5.5 0 0 0-.5-.5zm-6 4H4a.5.5 0 0 0 .5.5zM3.5 1h8V0h-8zm8.5.5v10h1v-10zM9.5 14h-6v1h6zM3 13.5v-12H2v12zm.5.5a.5.5 0 0 1-.5-.5H2A1.5 1.5 0 0 0 3.5 15zm8.5-2.5A2.5 2.5 0 0 1 9.5 14v1a3.5 3.5 0 0 0 3.5-3.5zM11.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 0zm-8-1A1.5 1.5 0 0 0 2 1.5h1a.5.5 0 0 1 .5-.5zm1 3h6V2h-6zm5.5-.5v3h1v-3zM9.5 6h-5v1h5zM5 6.5v-4H4v4zm5-1a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 11 5.5zM5 8v3h1V8zm-1 2h3V9H4zm6-1h1V8h-1zm-1 2h1v-1H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GbcSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6V3h5v2.5a.5.5 0 0 1-.5.5z"/><path fill="currentColor" fill-rule="evenodd" d="M2 1.5A1.5 1.5 0 0 1 3.5 0h8A1.5 1.5 0 0 1 13 1.5v10A3.5 3.5 0 0 1 9.5 15h-6A1.5 1.5 0 0 1 2 13.5zm2.5.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h5A1.5 1.5 0 0 0 11 5.5v-3a.5.5 0 0 0-.5-.5zM5 8v1H4v1h1v1h1v-1h1V9H6V8zm5 0v1h1V8zm-1 3v-1h1v1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GhostOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.914 14.086l.354.353zM4 14l-.354-.354zm2.5 0l.354-.354zm.146.146l-.353.354zm1.708 0L8 13.793zM8.5 14l.354.354zm2.5 0l-.354.354zm.086.086l-.354.353zM5 6h1V5H5zm4 0h1V5H9zm1.1 1.7a3.25 3.25 0 0 1-5.2 0l-.8.6c1.7 2.267 5.1 2.267 6.8 0zm-5.832 6.74l.086-.086l-.708-.708l-.085.086zm.94-.44h.085v-1h-.086zm.938.354l.147.146l.707-.707l-.146-.147zm2.561.146l.147-.146l-.708-.708l-.146.147zm1-.5h.086v-1h-.086zm.94.354l.085.085l.707-.707l-.085-.086zm1.439.646A1.914 1.914 0 0 0 14 13.086h-1c0 .505-.41.914-.914.914zm-1.354-.56c.36.358.846.56 1.354.56v-1a.914.914 0 0 1-.647-.268zm-.94-.44c.321 0 .628.127.854.354l.708-.708A2.207 2.207 0 0 0 9.793 13zm-.938.354c.226-.227.533-.354.853-.354v-1c-.585 0-1.147.232-1.56.646zM7.5 15c.453 0 .887-.18 1.207-.5L8 13.793a.707.707 0 0 1-.5.207zm-1.207-.5c.32.32.754.5 1.207.5v-1a.707.707 0 0 1-.5-.207zm-1-.5c.32 0 .627.127.853.354l.708-.708A2.207 2.207 0 0 0 5.293 13zm-.94.354c.227-.227.534-.354.854-.354v-1c-.585 0-1.147.232-1.56.646zM2.915 15c.508 0 .995-.202 1.354-.56l-.707-.708a.914.914 0 0 1-.647.268zM1 13.086C1 14.143 1.857 15 2.914 15v-1A.914.914 0 0 1 2 13.086zM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0zM2 6.5A5.5 5.5 0 0 1 7.5 1V0A6.5 6.5 0 0 0 1 6.5zm-1 0v6.586h1V6.5zm13 6.586V6.5h-1v6.586z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GhostSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 0 1 13 0v6.586a1.914 1.914 0 0 1-3.268 1.353l-.086-.085A1.207 1.207 0 0 0 9.793 14h-.086c-.32 0-.627.127-.853.354l-.147.146a1.707 1.707 0 0 1-2.414 0l-.147-.146A1.207 1.207 0 0 0 5.293 14h-.086c-.32 0-.627.127-.853.354l-.086.085A1.914 1.914 0 0 1 1 13.086zM5 6h1V5H5zm4 0h1V5H9zM4.9 7.7a3.25 3.25 0 0 0 5.2 0l.8.6c-1.7 2.267-5.1 2.267-6.8 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GifOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 10.5H2v.5h.5zm2 0v.5H5v-.5zm9-7h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zM2 6v4.5h1V6zm.5 5h2v-1h-2zm2.5-.5v-2H4v2zM3 7h2V6H3zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zM6 7h3V6H6zm0 4h3v-1H6zm1-4.5v4h1v-4zm3.5.5H13V6h-2.5zM10 6v5h1V6zm.5 3H12V8h-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GifSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM2 11V6h3v1H3v3h1V8.5h1V11zm6-4h1V6H6v1h1v3H6v1h3v-1H8zm5-1v1h-2v1h1v1h-1v2h-1V6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.5 7.5h-12m12 0a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1m12 0v5a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-5m6-3v-1m0 1H4.214A1.714 1.714 0 0 1 2.5 2.786V2.5a2 2 0 0 1 2-2a3 3 0 0 1 3 3m0 1h3.286c.947 0 1.714-.768 1.714-1.714V2.5a2 2 0 0 0-2-2a3 3 0 0 0-3 3m0 1v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 0A2.5 2.5 0 0 0 2 2.5v.286c0 .448.133.865.362 1.214H1.5A1.5 1.5 0 0 0 0 5.5v1A1.5 1.5 0 0 0 1.5 8H7V4h1v4h5.5A1.5 1.5 0 0 0 15 6.5v-1A1.5 1.5 0 0 0 13.5 4h-.862c.229-.349.362-.766.362-1.214V2.5A2.5 2.5 0 0 0 10.5 0c-1.273 0-2.388.68-3 1.696A3.498 3.498 0 0 0 4.5 0M8 4h2.786C11.456 4 12 3.456 12 2.786V2.5A1.5 1.5 0 0 0 10.5 1A2.5 2.5 0 0 0 8 3.5zM7 4H4.214C3.544 4 3 3.456 3 2.786V2.5A1.5 1.5 0 0 1 4.5 1A2.5 2.5 0 0 1 7 3.5z" clip-rule="evenodd"/><path fill="currentColor" d="M7 9H1v3.5A2.5 2.5 0 0 0 3.5 15H7zm1 6h3.5a2.5 2.5 0 0 0 2.5-2.5V9H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitBranchOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 4.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 0v6m2 2a2 2 0 1 1-2-2m2 2a2 2 0 0 0-2-2m2 2h5a3 3 0 0 0 3-3v-2m0 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitBranchSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95v5.1A2.5 2.5 0 1 0 4.95 13H9.5A3.5 3.5 0 0 0 13 9.5V7.95a2.5 2.5 0 1 0-1 0V9.5A2.5 2.5 0 0 1 9.5 12H4.95A2.503 2.503 0 0 0 3 10.05v-5.1A2.5 2.5 0 0 0 2.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCommitOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 10.5a3 3 0 0 1 0-6m0 6a3 3 0 0 0 0-6m0 6V15m0-10.5V0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCommitSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 7.5a3.5 3.5 0 0 1 3-3.465V0h1v4.035a3.5 3.5 0 0 1 0 6.93V15H7v-4.035A3.5 3.5 0 0 1 4 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCompareOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m8.5.5l-2 2m0 0l2 2m-2-2h3a3 3 0 0 1 3 3v5m-10-6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 0v5a3 3 0 0 0 3 3H8m-1.5 2l2-2l-2-2m6 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCompareSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95V9.5A3.5 3.5 0 0 0 5.5 13h1.793l-1.147 1.146l.708.708L9.207 12.5l-2.353-2.354l-.708.708L7.293 12H5.5A2.5 2.5 0 0 1 3 9.5V4.95A2.5 2.5 0 0 0 2.5 0m6.354.854L8.146.146L5.793 2.5l2.353 2.354l.708-.708L7.707 3H9.5A2.5 2.5 0 0 1 12 5.5v4.55a2.5 2.5 0 1 0 1 0V5.5A3.5 3.5 0 0 0 9.5 2H7.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitForkOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 4.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 0v6m0 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm10-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 0v1a3 3 0 0 1-3 3h-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitForkSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95v5.1a2.5 2.5 0 1 0 1 0V9h6.5A3.5 3.5 0 0 0 13 5.5v-.55a2.5 2.5 0 1 0-1 0v.55A2.5 2.5 0 0 1 9.5 8H3V4.95A2.5 2.5 0 0 0 2.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitMergeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 10.5a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0v-6m2-2a2 2 0 1 0-2 2m2-2a2 2 0 0 1-2 2m2-2h5a3 3 0 0 1 3 3v2m0 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitMergeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95v5.1a2.5 2.5 0 1 0 1 0v-5.1A2.503 2.503 0 0 0 4.95 3H9.5A2.5 2.5 0 0 1 12 5.5v1.55a2.5 2.5 0 1 0 1 0V5.5A3.5 3.5 0 0 0 9.5 2H4.95A2.5 2.5 0 0 0 2.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.793 1.207l.353.354zM1.207 6.793l-.353-.354zm0 1.414l.354-.353zm5.586 5.586l-.354.353zm1.414 0l-.353-.354zm5.586-5.586l.353.354zm0-1.414l-.354.353zM8.207 1.207l.354-.353zM6.44.854L.854 6.439l.707.707l5.585-5.585zM.854 8.56l5.585 5.585l.707-.707l-5.585-5.585zm7.707 5.585l5.585-5.585l-.707-.707l-5.585 5.585zm5.585-7.707L8.561.854l-.707.707l5.585 5.585zm0 2.122a1.5 1.5 0 0 0 0-2.122l-.707.707a.5.5 0 0 1 0 .708zM6.44 14.146a1.5 1.5 0 0 0 2.122 0l-.707-.707a.5.5 0 0 1-.708 0zM.854 6.44a1.5 1.5 0 0 0 0 2.122l.707-.707a.5.5 0 0 1 0-.708zm6.292-4.878a.5.5 0 0 1 .708 0L8.56.854a1.5 1.5 0 0 0-2.122 0zm-2 1.293l1 1l.708-.708l-1-1zM7.5 5a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 6zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 4.5zM7.5 4a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 3zm0-1A1.5 1.5 0 0 0 6 4.5h1a.5.5 0 0 1 .5-.5zm.646 2.854l1.5 1.5l.707-.708l-1.5-1.5zM10.5 8a.5.5 0 0 1-.5-.5H9A1.5 1.5 0 0 0 10.5 9zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 12 7.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 6zm0-1A1.5 1.5 0 0 0 9 7.5h1a.5.5 0 0 1 .5-.5zM7 5.5v4h1v-4zm.5 5.5a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 12zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 10.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 9zm0-1A1.5 1.5 0 0 0 6 10.5h1a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m8.5.5l-2 2m0 0l2 2m-2-2h3a3 3 0 0 1 3 3v2m-10 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0v-6m0 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm10 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 2 4.95v5.1a2.5 2.5 0 1 0 1 0v-5.1A2.5 2.5 0 0 0 2.5 0m6.354.854L8.146.146L5.793 2.5l2.353 2.354l.708-.708L7.707 3H9.5A2.5 2.5 0 0 1 12 5.5v1.55a2.5 2.5 0 1 0 1 0V5.5A3.5 3.5 0 0 0 9.5 2H7.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.44.854a1.5 1.5 0 0 1 2.12 0l5.586 5.585a1.5 1.5 0 0 1 0 2.122l-5.585 5.585a1.5 1.5 0 0 1-2.122 0L.854 8.561a1.5 1.5 0 0 1 0-2.122L4.793 2.5l1.353 1.353A1.5 1.5 0 0 0 7 5.914v3.171a1.5 1.5 0 1 0 1 0v-3.17c.05-.018.1-.038.147-.061l1 1a1.5 1.5 0 1 0 .707-.707l-1-1a1.5 1.5 0 0 0-2-2L5.5 1.792z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.65 12.477a.5.5 0 1 0-.3-.954zm-3.648-2.96l-.484-.128l-.254.968l.484.127zM9 14.5v.5h1v-.5zm.063-4.812l-.054-.498a.5.5 0 0 0-.299.852zM12.5 5.912h.5v-.001zm-.833-2.006l-.466-.18a.5.5 0 0 0 .112.533zm-.05-2.017l.456-.204a.5.5 0 0 0-.319-.276zm-2.173.792l-.126.484a.5.5 0 0 0 .398-.064zm-3.888 0l-.272.42a.5.5 0 0 0 .398.064zM3.383 1.89l-.137-.48a.5.5 0 0 0-.32.276zm-.05 2.017l.354.353a.5.5 0 0 0 .112-.534zM2.5 5.93H3v-.002zm3.438 3.758l.352.355a.5.5 0 0 0-.293-.851zM5.5 11H6l-.001-.037zM5 14.5v.5h1v-.5zm.35-2.977c-.603.19-.986.169-1.24.085c-.251-.083-.444-.25-.629-.49a4.8 4.8 0 0 1-.27-.402c-.085-.139-.182-.302-.28-.447c-.191-.281-.473-.633-.929-.753l-.254.968c.08.02.184.095.355.346c.082.122.16.252.258.412c.094.152.202.32.327.484c.253.33.598.663 1.11.832c.51.168 1.116.15 1.852-.081zm4.65-.585c0-.318-.014-.608-.104-.878c-.096-.288-.262-.51-.481-.727l-.705.71c.155.153.208.245.237.333c.035.105.053.254.053.562zm-.884-.753c.903-.097 1.888-.325 2.647-.982c.78-.675 1.237-1.729 1.237-3.29h-1c0 1.359-.39 2.1-.892 2.534c-.524.454-1.258.653-2.099.743zM13 5.91a3.354 3.354 0 0 0-.98-2.358l-.707.706c.438.44.685 1.034.687 1.655zm-.867-1.824c.15-.384.22-.794.21-1.207l-1 .025a2.12 2.12 0 0 1-.142.82zm.21-1.207a3.119 3.119 0 0 0-.27-1.195l-.913.408c.115.256.177.532.184.812zm-.726-.99c.137-.481.137-.482.136-.482h-.003l-.004-.002a.462.462 0 0 0-.03-.007a1.261 1.261 0 0 0-.212-.024a2.172 2.172 0 0 0-.51.054c-.425.091-1.024.317-1.82.832l.542.84c.719-.464 1.206-.634 1.488-.694a1.2 1.2 0 0 1 .306-.03l-.008-.001a.278.278 0 0 1-.01-.002l-.006-.002h-.003l-.002-.001c-.001 0-.002 0 .136-.482m-2.047.307a8.209 8.209 0 0 0-4.14 0l.252.968a7.209 7.209 0 0 1 3.636 0zm-3.743.064C5.03 1.746 4.43 1.52 4.005 1.43a2.17 2.17 0 0 0-.51-.053a1.259 1.259 0 0 0-.241.03l-.004.002h-.003l.136.481l.137.481h-.001l-.002.001l-.003.001a.327.327 0 0 1-.016.004l-.008.001h.008a1.19 1.19 0 0 1 .298.03c.282.06.769.23 1.488.694zm-2.9-.576a3.12 3.12 0 0 0-.27 1.195l1 .025a2.09 2.09 0 0 1 .183-.812zm-.27 1.195c-.01.413.06.823.21 1.207l.932-.362a2.12 2.12 0 0 1-.143-.82zm.322.673a3.354 3.354 0 0 0-.726 1.091l.924.38c.118-.285.292-.545.51-.765zm-.726 1.091A3.354 3.354 0 0 0 2 5.93l1-.003c0-.31.06-.616.177-.902zM2 5.93c0 1.553.458 2.597 1.239 3.268c.757.65 1.74.88 2.64.987l.118-.993C5.15 9.09 4.416 8.89 3.89 8.438C3.388 8.007 3 7.276 3 5.928zm3.585 3.404c-.5.498-.629 1.09-.584 1.704L6 10.963c-.03-.408.052-.683.291-.921zM5 11v3.5h1V11zm5 3.5V13H9v1.5zm0-1.5v-2.062H9V13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.358 2.145a8.209 8.209 0 0 0-3.716 0c-.706-.433-1.245-.632-1.637-.716a2.17 2.17 0 0 0-.51-.053a1.258 1.258 0 0 0-.232.028l-.01.002l-.004.002h-.003l.137.481l-.137-.48a.5.5 0 0 0-.32.276a3.12 3.12 0 0 0-.159 2.101A3.354 3.354 0 0 0 2 5.93c0 1.553.458 2.597 1.239 3.268c.547.47 1.211.72 1.877.863a2.34 2.34 0 0 0-.116.958v.598c-.407.085-.689.058-.89-.008c-.251-.083-.444-.25-.629-.49a4.798 4.798 0 0 1-.27-.402l-.057-.093a9.216 9.216 0 0 0-.224-.354c-.19-.281-.472-.633-.928-.753l-.484-.127l-.254.968l.484.127c.08.02.184.095.355.346a7.2 7.2 0 0 1 .19.302l.068.11c.094.152.202.32.327.484c.253.33.598.663 1.11.832c.35.116.748.144 1.202.074V14.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-3.562c0-.316-.014-.605-.103-.874c.663-.14 1.322-.39 1.866-.86c.78-.676 1.237-1.73 1.237-3.292v-.001a3.354 3.354 0 0 0-.768-2.125a3.12 3.12 0 0 0-.159-2.1a.5.5 0 0 0-.319-.277l-.137.48c.137-.48.136-.48.135-.48l-.002-.001l-.004-.002l-.009-.002a.671.671 0 0 0-.075-.015a1.261 1.261 0 0 0-.158-.013a2.172 2.172 0 0 0-.51.053c-.391.084-.93.283-1.636.716"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitlabOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m.5 8.5l7 6l7-6l-2-8l-2 6h-6l-2-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitlabSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.974.342a.5.5 0 0 0-.96.037l-2 8a.5.5 0 0 0 .16.5l7 6a.5.5 0 0 0 .651 0l7-6a.5.5 0 0 0 .16-.5l-2-8a.5.5 0 0 0-.96-.037L10.14 6H4.86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAfricaOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1H8a.5.5 0 0 0-.276-.447zm0 1l.224.447A.5.5 0 0 0 8 2zm-1 .5l-.224-.447a.5.5 0 0 0-.053.863zm1.5 1l-.277.416A.5.5 0 0 0 8 4zm.5 0V4a.5.5 0 0 0 .5-.5zm0-1V2a.5.5 0 0 0-.5.5zm1.5 0h.5A.5.5 0 0 0 10 2zm0 1h-.5a.5.5 0 0 0 .146.354zm.5.5h.5a.5.5 0 0 0-.146-.354zm0 1l.354.354A.5.5 0 0 0 11 5zm-.5.5V6a.5.5 0 0 0 .354-.146zm-1 0l-.224.447A.5.5 0 0 0 9 6zm-2-1l.224-.447a.5.5 0 0 0-.448 0zM6 5v.5a.5.5 0 0 0 .224-.053zM4.5 5v-.5a.5.5 0 0 0-.485.379zM4 7l-.485-.121a.5.5 0 0 0 .131.475zm1.5 1.5l-.354.354a.5.5 0 0 0 .13.093zm1 .5H7a.5.5 0 0 0-.276-.447zm0 1.5H6a.5.5 0 0 0 .146.354zm.5.5h.5a.5.5 0 0 0-.146-.354zm0 1h-.5a.5.5 0 0 0 .053.224zm.5 1l-.447.224a.5.5 0 0 0 .447.276zm1 0v.5a.5.5 0 0 0 .416-.223zm1-1.5l.416.277a.5.5 0 0 0 .031-.053zm.5-1l.447.224a.5.5 0 0 0 .053-.224zm0-1l-.4-.3a.5.5 0 0 0-.1.3zm1.5-2l.4.3a.5.5 0 0 0 .047-.524zm-.5-1V6a.5.5 0 0 0-.447.724zm1.5 0V7a.5.5 0 0 0 .5-.5zm0-1V5a.5.5 0 0 0-.5.5zm-5 8.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM6.276.947l1 .5l.448-.894l-1-.5zM7 1v1h1V1zm.276.553l-1 .5l.448.894l1-.5zM6.223 2.916l1.5 1l.554-.832l-1.5-1zM8 4h.5V3H8zm1-.5v-1H8v1zM8.5 3H10V2H8.5zm1-.5v1h1v-1zm.146 1.354l.5.5l.708-.708l-.5-.5zM10 4v1h1V4zm.146.646l-.5.5l.708.708l.5-.5zM10 5H9v1h1zm-.776.053l-2-1l-.448.894l2 1zm-2.448-1l-1 .5l.448.894l1-.5zM6 4.5H4.5v1H6zm-1.985.379l-.5 2l.97.242l.5-2zm-.369 2.475l1.5 1.5l.708-.708l-1.5-1.5zm1.63 1.593l1 .5l.448-.894l-1-.5zM6 9v1.5h1V9zm.146 1.854l.5.5l.708-.708l-.5-.5zM6.5 11v1h1v-1zm.053 1.224l.5 1l.894-.448l-.5-1zM7.5 13.5h1v-1h-1zm1.416-.223l1-1.5l-.832-.554l-1 1.5zm1.031-1.553l.5-1l-.894-.448l-.5 1zM10.5 10.5v-1h-1v1zm-.1-.7l1.5-2l-.8-.6l-1.5 2zm1.547-2.524l-.5-1l-.894.448l.5 1zM11 7h1.5V6H11zm2-.5v-1h-1v1zm-.5-.5h2V5h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAfricaSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m1 0a6.502 6.502 0 0 1 5.527-6.428L7 1.31v.382l-.724.362a.5.5 0 0 0-.053.863l1.5 1A.5.5 0 0 0 8 4h.5a.5.5 0 0 0 .5-.5V3h.5v.5a.5.5 0 0 0 .146.354l.354.353v.586L9.793 5h-.675l-1.894-.947a.5.5 0 0 0-.448 0l-.894.447H4.5a.5.5 0 0 0-.485.379l-.5 2a.5.5 0 0 0 .131.475l1.5 1.5a.5.5 0 0 0 .13.093L6 9.31v1.19a.5.5 0 0 0 .146.354l.354.353V12a.5.5 0 0 0 .053.224l.5 1a.5.5 0 0 0 .447.276h1a.5.5 0 0 0 .416-.223l1-1.5a.5.5 0 0 0 .031-.053l.5-1a.5.5 0 0 0 .053-.224v-.833L11.9 7.8a.5.5 0 0 0 .047-.524L11.81 7h.691a.5.5 0 0 0 .5-.5V6h.826A6.5 6.5 0 1 1 1 7.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAmericasOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 5.5V5a.5.5 0 0 0-.5.5zm0 1l.354.354L5 6.707V6.5zm-1.707.293l-.354.353zM6.5 13H7v-.207l-.146-.147zm-1-1H5v.207l.146.147zm0-1.5H6v-.207l-.146-.147zm-1-1H4v.207l.146.147zm0-1V8a.5.5 0 0 0-.5.5zM9 .5v2h1v-2zM8.5 3h-1v1h1zm-3 2h-1v1h1zM4 5.5v1h1v-1zm.146.646l-.292.293l.707.707l.293-.292zm-1 .293L1.354 4.646l-.708.708L2.44 7.146zM6 4.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 7 4.5zM7.5 3A1.5 1.5 0 0 0 6 4.5h1a.5.5 0 0 1 .5-.5zM3.854 6.44a.5.5 0 0 1-.708 0l-.707.706a1.5 1.5 0 0 0 2.122 0zM9 2.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 10 2.5zm-2 12V13H6v1.5zm-.146-1.854l-1-1l-.708.708l1 1zM6 12v-1.5H5V12zm-.146-1.854l-1-1l-.708.708l1 1zM5 9.5v-1H4v1zM4.5 9h4V8h-4zm4.5.5v.667h1V9.5zm1.833 2.5H13.5v-1h-2.667zM10 11.167c0 .46.373.833.833.833v-1c.092 0 .167.075.167.167zM9.833 11c.092 0 .167.075.167.167h1C11 10.522 10.478 10 9.833 10zM9 10.167c0 .46.373.833.833.833v-1c.092 0 .167.075.167.167zM8.5 9a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 8.5 8zm-1 5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAmericasSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15M1.197 5.904A6.503 6.503 0 0 0 6 13.826v-.619l-1-1v-1.5l-1-1V8.5a.5.5 0 0 1 .5-.5h4A1.5 1.5 0 0 1 10 9.5v.512c.51.073.915.477.988.988h1.99A6.502 6.502 0 0 0 10 1.498V2.5A1.5 1.5 0 0 1 8.5 4h-1a.5.5 0 0 0-.5.5A1.5 1.5 0 0 1 5.5 6H5v.707l-.44.44a1.5 1.5 0 0 1-2.12 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 15v-3M4 14.5h7M11.469 1A6 6 0 1 1 3.5 9.972m4-7.472a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 5.5A6.486 6.486 0 0 0 11.8.625l-.662.75A5.5 5.5 0 1 1 3.834 9.6l-.667.745A6.476 6.476 0 0 0 7 11.98V14H4v1h7v-1H8v-2.019A6.5 6.5 0 0 0 14 5.5"/><path fill="currentColor" d="M7.5 2a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleAdOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 10V7a1.5 1.5 0 1 1 3 0v3m5-5v5m0-2.5h-2a1 1 0 0 0 0 2h2m-8-1h3m-4-6h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleAdSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6a1 1 0 0 0-1 1v1h2V7a1 1 0 0 0-1-1m6 2H9.5a.5.5 0 0 0 0 1H11z"/><path fill="currentColor" fill-rule="evenodd" d="M0 4.5A2.5 2.5 0 0 1 2.5 2h10A2.5 2.5 0 0 1 15 4.5v6a2.5 2.5 0 0 1-2.5 2.5h-10A2.5 2.5 0 0 1 0 10.5zM4 10V9h2v1h1V7a2 2 0 1 0-4 0v3zm7-3H9.5a1.5 1.5 0 1 0 0 3H12V5h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDriveOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m5 1.5l-4.5 8l2 4M5 1.5l2.5 4l-5 8M5 1.5h5l4.5 8M5 1.5l5 8h4.5m-12 4l2.5-4h9.5m-12 4h10l2-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDriveSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.563 1.258A.5.5 0 0 1 5 1h5a.5.5 0 0 1 .436.255L14.23 8H8.473L4.576 1.765a.5.5 0 0 1-.013-.507M3.91 14h8.59a.5.5 0 0 0 .447-.276l2-4A.5.5 0 0 0 14.5 9H7.092zM.064 9.255l3.792-6.742l3.05 4.88l-3.982 6.372a.5.5 0 0 1-.871-.041l-2-4a.5.5 0 0 1 .011-.47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 7.5h.5V7h-.5zm-7 6.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zm0 1c1.819 0 3.463.746 4.643 1.951l.714-.7A7.479 7.479 0 0 0 7.5 0zM8 8h6.5V7H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlayStoreOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.5.5l.252-.432A.5.5 0 0 0 1 .5zm0 14H1a.5.5 0 0 0 .752.432zm12-7l.252.432a.5.5 0 0 0 0-.864zM1 .5v14h1V.5zm.752 14.432l12-7l-.504-.864l-12 7zm12-7.864l-12-7l-.504.864l12 7zM1.829 12.876l8-7l-.658-.752l-8 7zm-.658-10l8 7l.658-.752l-8-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlayStoreSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.251.066a.5.5 0 0 1 .5.002l7.843 4.575l-2.427 2.184L1 1.277V.5a.5.5 0 0 1 .251-.434M1 2.623v9.754L6.42 7.5zm0 11.1v.777a.5.5 0 0 0 .752.432l7.842-4.575l-2.427-2.184zm9.501-3.895l3.25-1.896a.5.5 0 0 0 0-.864l-3.25-1.896L7.914 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 0 1 12.857-5.249l-.714.7A6.5 6.5 0 1 0 13.98 8H8V7h7v.5a7.5 7.5 0 0 1-15 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleStreetviewOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 11.5H4a.5.5 0 0 0 .5.5zm6 0v.5a.5.5 0 0 0 .5-.5zm-6 .5h6v-1h-6zm6.5-.5v-2h-1v2zm-7-2v2h1v-2zM7.5 6A3.5 3.5 0 0 0 4 9.5h1A2.5 2.5 0 0 1 7.5 7zM11 9.5A3.5 3.5 0 0 0 7.5 6v1A2.5 2.5 0 0 1 10 9.5zm3 2c0 .245-.114.52-.406.816c-.294.299-.745.59-1.341.846c-1.191.51-2.871.838-4.753.838v1c1.984 0 3.804-.344 5.147-.92c.67-.287 1.245-.642 1.659-1.061c.416-.422.694-.936.694-1.519zM7.5 14c-1.882 0-3.562-.328-4.753-.838c-.596-.256-1.047-.547-1.341-.846C1.114 12.02 1 11.747 1 11.5H0c0 .583.278 1.097.694 1.519c.414.42.989.774 1.66 1.062C3.695 14.656 5.515 15 7.5 15zM1 11.5c0-.242.11-.513.394-.805c.286-.294.725-.582 1.306-.837l-.4-.916c-.656.287-1.218.64-1.622 1.056C.27 10.416 0 10.925 0 11.5zm11.3-1.642c.581.255 1.02.543 1.305.837c.284.292.395.563.395.805h1c0-.575-.27-1.084-.678-1.502c-.404-.416-.966-.769-1.622-1.056zM7.5 4A1.5 1.5 0 0 1 6 2.5H5A2.5 2.5 0 0 0 7.5 5zM9 2.5A1.5 1.5 0 0 1 7.5 4v1A2.5 2.5 0 0 0 10 2.5zM7.5 1A1.5 1.5 0 0 1 9 2.5h1A2.5 2.5 0 0 0 7.5 0zm0-1A2.5 2.5 0 0 0 5 2.5h1A1.5 1.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleStreetviewSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0M7.5 6A3.5 3.5 0 0 0 4 9.5v2a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-2A3.5 3.5 0 0 0 7.5 6"/><path fill="currentColor" d="M1.394 10.695c-.283.292-.394.563-.394.805c0 .245.114.52.406.816c.294.299.745.59 1.341.846c1.191.51 2.871.838 4.753.838c1.882 0 3.562-.328 4.753-.838c.596-.256 1.047-.547 1.341-.846c.292-.296.406-.57.406-.816c0-.242-.11-.513-.395-.805c-.285-.294-.724-.582-1.305-.837l.4-.916c.656.287 1.218.64 1.622 1.056c.407.418.678.927.678 1.502c0 .583-.278 1.097-.694 1.519c-.414.42-.989.774-1.66 1.062c-1.342.575-3.162.919-5.146.919s-3.804-.344-5.147-.92c-.67-.287-1.245-.642-1.659-1.061c-.416-.422-.694-.936-.694-1.52c0-.575.27-1.083.678-1.501c.404-.416.966-.769 1.622-1.056l.4.916c-.581.255-1.02.543-1.306.837"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphqlOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.61 4.432l4.142-2.417l-.504-.864l-4.143 2.417zM2 9.5v-4H1v4zm6.248-7.485l4.143 2.417l.504-.864l-4.143-2.417zM13 5.5v4h1v-4zm-.252 4.86l-4.5 2.625l.504.864l4.5-2.625zm-5.996 2.625l-4.143-2.417l-.504.864l4.143 2.417zM6.584 1.973l-5 7.5l.832.554l5-7.5zm6.832 7.5l-5-7.5l-.832.554l5 7.5zM2.5 11h10v-1h-10zm5-9a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 3zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 1.5zM7.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 0zm0-1A1.5 1.5 0 0 0 6 1.5h1a.5.5 0 0 1 .5-.5zm6 5a.5.5 0 0 1-.5-.5h-1A1.5 1.5 0 0 0 13.5 6zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 15 4.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 3zm0-1A1.5 1.5 0 0 0 12 4.5h1a.5.5 0 0 1 .5-.5zm0 8a.5.5 0 0 1-.5-.5h-1a1.5 1.5 0 0 0 1.5 1.5zm.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 9zm0-1a1.5 1.5 0 0 0-1.5 1.5h1a.5.5 0 0 1 .5-.5zm-6 5a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 15zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 13.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 12zm0-1A1.5 1.5 0 0 0 6 13.5h1a.5.5 0 0 1 .5-.5zm-6-1a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 12zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 3 10.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 1.5 9zm0-1A1.5 1.5 0 0 0 0 10.5h1a.5.5 0 0 1 .5-.5zm0-4a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 6zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 3 4.5zM1.5 4a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 1.5 3zm0-1A1.5 1.5 0 0 0 0 4.5h1a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphqlSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.015 1.287a1.5 1.5 0 0 1 2.97 0l3.545 2.069A1.5 1.5 0 1 1 14 5.916v3.17a1.5 1.5 0 1 1-1.47 2.559l-3.545 2.068a1.5 1.5 0 0 1-2.97 0L2.47 11.645A1.5 1.5 0 1 1 1 9.085v-3.17a1.5 1.5 0 1 1 1.47-2.56zm.225 1.027L2.974 4.219A1.5 1.5 0 0 1 2 5.914V8.85L6.3 2.4a1.5 1.5 0 0 1-.06-.085m.891.64l-4.43 6.647c.09.12.163.254.214.399h9.17a1.51 1.51 0 0 1 .214-.4L7.87 2.955a1.503 1.503 0 0 1-.738 0m1.57-.555L13 8.85V5.915a1.5 1.5 0 0 1-.974-1.696L8.76 2.314a1.5 1.5 0 0 1-.06.085M11.65 11H3.349l2.89 1.686a1.499 1.499 0 0 1 2.521 0zM7.5 1a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m-6 3a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m12 0a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m-12 6a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m12 0a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m-6 3a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridLayoutOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5.5h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Zm8 0h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Zm0 8h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Zm-8 0h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridLayoutSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v4A1.5 1.5 0 0 0 1.5 7h4A1.5 1.5 0 0 0 7 5.5v-4A1.5 1.5 0 0 0 5.5 0zm8 0A1.5 1.5 0 0 0 8 1.5v4A1.5 1.5 0 0 0 9.5 7h4A1.5 1.5 0 0 0 15 5.5v-4A1.5 1.5 0 0 0 13.5 0zm-8 8A1.5 1.5 0 0 0 0 9.5v4A1.5 1.5 0 0 0 1.5 15h4A1.5 1.5 0 0 0 7 13.5v-4A1.5 1.5 0 0 0 5.5 8zm8 0A1.5 1.5 0 0 0 8 9.5v4A1.5 1.5 0 0 0 9.5 15h4a1.5 1.5 0 0 0 1.5-1.5v-4A1.5 1.5 0 0 0 13.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashtagOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 5.5h11m-11 4h11m-6.5-8l-2 12m6-12l-2 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashtagSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.41 5l.597-3.582l.986.164L6.423 5H9.41l.597-3.582l.986.164L10.423 5H13v1h-2.743l-.5 3H13v1H9.59l-.597 3.582l-.986-.164l.57-3.418H5.59l-.597 3.582l-.986-.164l.57-3.418H2V9h2.743l.5-3H2V5zm.847 1l-.5 3h2.986l.5-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HdScreenOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 13.5h11m-5.5-3V14m6.69-2.589a24.35 24.35 0 0 0-13.38 0a.243.243 0 0 1-.31-.234V1.823c0-.162.155-.279.31-.234a24.35 24.35 0 0 0 13.38 0a.243.243 0 0 1 .31.234v9.354a.243.243 0 0 1-.31.234Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HdScreenSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.948 1.108A.744.744 0 0 0 0 1.823v9.354c0 .494.473.85.948.715A23.85 23.85 0 0 1 7 10.98V13H2v1h11v-1H8v-2.02c2.039.042 4.073.346 6.052.912a.744.744 0 0 0 .948-.715V1.823a.744.744 0 0 0-.948-.715a23.85 23.85 0 0 1-13.104 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HdmiCableOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 4V.5h8V4M5 2.5h1m3 0h1M5.5 13v2m4-2v2m-7-10.5h10v5l-2 1v2h-6v-2l-2-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HdmiCableSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 0h9v5H3zm3 3H5V2h1zm4 0H9V2h1z" clip-rule="evenodd"/><path fill="currentColor" d="M2 6h11v3.809l-2 1V13h-1v2H9v-2H6v2H5v-2H4v-2.191l-2-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadphonesOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 9.5h1a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1Zm0 0v-3a6 6 0 1 1 12 0v3m0 0h-1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadphonesSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6.5a5.5 5.5 0 1 1 11 0V9h-.5a1.5 1.5 0 0 0-1.5 1.5v3a1.5 1.5 0 0 0 1.5 1.5h1a1.5 1.5 0 0 0 1.5-1.5v-3a1.5 1.5 0 0 0-1-1.415V6.5a6.5 6.5 0 1 0-13 0v2.585A1.5 1.5 0 0 0 0 10.5v3A1.5 1.5 0 0 0 1.5 15h1A1.5 1.5 0 0 0 4 13.5v-3A1.5 1.5 0 0 0 2.5 9H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 12.5a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1Zm0 0a2 2 0 0 1-2 2H8m6.5-4.5V7.5a7 7 0 1 0-14 0V10m2 2.5a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadsetSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 6a2.49 2.49 0 0 0-1.414.438a6.502 6.502 0 0 1 12.828 0A2.488 2.488 0 0 0 12.5 6A1.5 1.5 0 0 0 11 7.5v4a1.5 1.5 0 0 0 .947 1.395A1.5 1.5 0 0 1 10.5 14H8v1h2.5a2.5 2.5 0 0 0 2.458-2.042A2.5 2.5 0 0 0 15 10.5v-3a7.5 7.5 0 0 0-15 0v3A2.5 2.5 0 0 0 2.5 13A1.5 1.5 0 0 0 4 11.5v-4A1.5 1.5 0 0 0 2.5 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="m6.5 5.5l1 1l1-1a1.414 1.414 0 0 1 2 2l-3 3l-3-3a1.414 1.414 0 0 1 2-2Z"/><path d="M.5 7.5a7 7 0 1 0 14 0a7 7 0 0 0-14 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m4.146-2.354a1.914 1.914 0 0 1 2.707 0l.647.647l.646-.647a1.914 1.914 0 0 1 2.707 2.708L7.5 11.207L4.146 7.854a1.914 1.914 0 0 1 0-2.708" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 13.5l-.354.354a.5.5 0 0 0 .708 0zM1.536 7.536l-.354.353zm5-5l-.354.353zM7.5 3.5l-.354.354a.5.5 0 0 0 .708 0zm.964-.964l-.353-.354zm-.61 10.61L1.889 7.182l-.707.707l5.964 5.965zm5.257-5.964l-5.965 5.964l.708.708l5.964-5.965zM6.182 2.889l.964.965l.708-.708l-.965-.964zm1.672.965l.964-.965l-.707-.707l-.965.964zM10.964 1c-1.07 0-2.096.425-2.853 1.182l.707.707A3.037 3.037 0 0 1 10.964 2zM14 5.036c0 .805-.32 1.577-.89 2.146l.708.707A4.036 4.036 0 0 0 15 5.036zm1 0A4.036 4.036 0 0 0 10.964 1v1A3.036 3.036 0 0 1 14 5.036zM4.036 2c.805 0 1.577.32 2.146.89l.707-.708A4.036 4.036 0 0 0 4.036 1zM1 5.036A3.036 3.036 0 0 1 4.036 2V1A4.036 4.036 0 0 0 0 5.036zm.89 2.146A3.035 3.035 0 0 1 1 5.036H0c0 1.07.425 2.096 1.182 2.853z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m6.5 5.5l1 1l1-1a1.414 1.414 0 1 1 2 2l-3 3l-3-3a1.414 1.414 0 1 1 2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.765 5.235a1.79 1.79 0 1 0-2.53 2.53L7.5 11.03l3.265-3.265a1.79 1.79 0 0 0-2.53-2.53L7.5 5.97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.036 1a4.036 4.036 0 0 0-2.854 6.89l5.964 5.964a.5.5 0 0 0 .708 0l5.964-5.965a4.036 4.036 0 0 0-5.707-5.707l-.611.61l-.61-.61A4.036 4.036 0 0 0 4.035 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HexagonOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 4.5v6l6 3.5l6-3.5v-6L7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HexagonSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4.213L7.5.42L1 4.213v6.574l6.5 3.792l6.5-3.792z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistoryOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.5 12.399l.37-.336l-.006-.007l-.007-.007zm1 1.101v.5H4v-.5zm3.5.982l.018-.5l-.053 1zM7.5 7.5H7a.5.5 0 0 0 .146.354zm6.5 0A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM2.857 12.049A6.477 6.477 0 0 1 1 7.5H0c0 2.043.818 3.897 2.143 5.249zm-.727.686l1 1.101l.74-.672l-1-1.101zM7.5 14a6.62 6.62 0 0 1-.465-.016l-.07.997c.177.013.355.019.535.019zm.018 0l-.5-.017l-.036 1l.5.017zM7 3v4.5h1V3zm.146 4.854l3 3l.708-.708l-3-3zM0 14h3.5v-1H0zm4-.5V10H3v3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistorySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 6.965 7.481l.053-.998l.49.017a6.5 6.5 0 1 0-4.65-1.951l.006.007l.136.15V10h1v4H0v-1h2.37l-.234-.258A7.477 7.477 0 0 1 0 7.5m7 0V3h1v4.293l2.854 2.853l-.708.708l-3-3A.5.5 0 0 1 7 7.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.354-.354a.5.5 0 0 0-.708 0zm-6 6l-.354-.354L1 6.293V6.5zm12 0h.5v-.207l-.146-.147zm.354-.354l-6-6l-.708.708l6 6zm-6.708-6l-6 6l.708.708l6-6zM14 13.5v-7h-1v7zm-13-7v7h1v-7zM2.5 15h10v-1h-10zM13 13.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zm-12 0A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.854.146a.5.5 0 0 0-.708 0L1 6.293V13.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V6.293z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.325-.38a.5.5 0 0 0-.65 0zm-7 6l-.325-.38L0 6.27v.23zm5 8v.5a.5.5 0 0 0 .5-.5zm4 0H9a.5.5 0 0 0 .5.5zm5-8h.5v-.23l-.175-.15zM1.5 15h4v-1h-4zm13.325-8.88l-7-6l-.65.76l7 6zm-7.65-6l-7 6l.65.76l7-6zM6 14.5v-3H5v3zm3-3v3h1v-3zm.5 3.5h4v-1h-4zm5.5-1.5v-7h-1v7zm-15-7v7h1v-7zM7.5 10A1.5 1.5 0 0 1 9 11.5h1A2.5 2.5 0 0 0 7.5 9zm0-1A2.5 2.5 0 0 0 5 11.5h1A1.5 1.5 0 0 1 7.5 10zm6 6a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zm-12-1a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.825.12a.5.5 0 0 0-.65 0L0 6.27v7.23A1.5 1.5 0 0 0 1.5 15h4a.5.5 0 0 0 .5-.5v-3a1.5 1.5 0 0 1 3 0v3a.5.5 0 0 0 .5.5h4a1.5 1.5 0 0 0 1.5-1.5V6.27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.224-.447a.5.5 0 0 0-.448 0zm-6 3l-.224-.447A.5.5 0 0 0 1 3.5zm12 0h.5a.5.5 0 0 0-.276-.447zm-8 7V10H5v.5zm4 0h.5V10h-.5zM0 15h15v-1H0zM7.276.053l-6 3l.448.894l6-3zm6.448 3l-6-3l-.448.894l6 3zM7 3v2.5h1V3zm0 2.5V8h1V5.5zM5 6h2.5V5H5zm2.5 0H10V5H7.5zM1 3.5v11h1v-11zm12 0v11h1v-11zm-7 11v-4H5v4zM5.5 11h4v-1h-4zm3.5-.5v4h1v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.724.053a.5.5 0 0 0-.448 0l-5.99 2.995A.5.5 0 0 0 1 3.5V14H0v1h5v-5h5v5h5v-1h-1V3.5a.5.5 0 0 0-.286-.452zM7 5V3h1v2h2v1H8v2H7V6H5V5z" clip-rule="evenodd"/><path fill="currentColor" d="M9 15v-4H6v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 0v4.672a2 2 0 0 0 .586 1.414l.707.707a1 1 0 0 1 0 1.414l-1 1a1 1 0 0 0-.293.707V15m8-15v5.086a1 1 0 0 1-.293.707l-1 1a1 1 0 0 0 0 1.414l1 1a1 1 0 0 1 .293.707V15M1 .5h13m-13 14h13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1h2V0H1v1h2v3.672a2.5 2.5 0 0 0 .732 1.767l.707.707a.5.5 0 0 1 0 .708l-1 1A1.5 1.5 0 0 0 3 9.914V14H1v1h13v-1h-2V9.914a1.5 1.5 0 0 0-.44-1.06l-1-1a.5.5 0 0 1 0-.708l1-1a1.5 1.5 0 0 0 .44-1.06zM4.25 5.5h6.543l.06-.06A.5.5 0 0 0 11 5.085V1H4v3.672c0 .296.088.584.25.828" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 5.5l-.29-.407l-.21.15V5.5zm7-5l.29-.407a.5.5 0 0 0-.58 0zm7 5h.5v-.257l-.21-.15zm-12 3V8H2v.5zm4 0H7V8h-.5zM1 15V5.5H0V15zM.79 5.907l7-5l-.58-.814l-7 5zm6.42-5l7 5l.58-.814l-7-5zM14 5.5V15h1V5.5zM3 15V8.5H2V15zm-.5-6h4V8h-4zM6 8.5V15h1V8.5zM5 12h1.5v-1H5zm6-4v4h1V8zm2-6v2.5h1V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.79.093a.5.5 0 0 0-.58 0l-7 5A.5.5 0 0 0 0 5.5v9a.5.5 0 0 0 .5.5H2V8h5v7h7.5a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.21-.407L14 4.528V2h-1v1.814zM11 12V8h1v4z" clip-rule="evenodd"/><path fill="currentColor" d="M6 15v-3H5v-1h1V9H3v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFiveOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.498.542zm14 0l.498.042A.5.5 0 0 0 14.5 0zm-1 12l.158.474a.5.5 0 0 0 .34-.433zm-6 2l-.158.474a.499.499 0 0 0 .316 0zm-6-2l-.498.041a.5.5 0 0 0 .34.433zm3-9V3H4v.5zm0 3H4V7h.5zm6 0h.5V6h-.5zm0 3l.158.474L11 9.86V9.5zm-3 1l-.158.474l.158.053l.158-.053zm-3-1H4v.36l.342.114zM.5 1h14V0H.5zM14.002.458l-1 12l.996.083l1-12zm-.66 11.568l-6 2l.316.948l6-2zm-5.684 2l-6-2l-.316.948l6 2zm-5.66-1.567l-1-12l-.996.083l1 12zM11 3H4.5v1H11zm-7 .5v3h1v-3zM4.5 7h6V6h-6zm5.5-.5v3h1v-3zm.342 2.526l-3 1l.316.948l3-1zm-2.684 1l-3-1l-.316.948l3 1zM5 9.5V8H4v1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFiveSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.132.161A.5.5 0 0 1 .5 0h14a.5.5 0 0 1 .498.542l-1 12a.5.5 0 0 1-.34.432l-6 2a.499.499 0 0 1-.316 0l-6-2a.5.5 0 0 1-.34-.433l-1-12a.5.5 0 0 1 .13-.38M11 3H4v4h6v2.14l-2.5.833L5 9.14V8H4v1.86l3.5 1.167L11 9.86V6H5V4h6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12.5v.5h1v-.5zm5 0v.5h1v-.5zm-4 0V12H2v.5zm4-.5v.5h1V12zm-2-2a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3zm-2 2a2 2 0 0 1 2-2V9a3 3 0 0 0-3 3zm2-8a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1zm2 2a2 2 0 0 0-2-2v1a1 1 0 0 1 1 1zM5 8a2 2 0 0 0 2-2H6a1 1 0 0 1-1 1zm0-1a1 1 0 0 1-1-1H3a2 2 0 0 0 2 2zM1.5 3h12V2h-12zm12.5.5v8h1v-8zm-.5 8.5h-12v1h12zM1 11.5v-8H0v8zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 13zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM13.5 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 2zm-12-1A1.5 1.5 0 0 0 0 3.5h1a.5.5 0 0 1 .5-.5zM9 6h3V5H9zm0 3h3V8H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 3.5A1.5 1.5 0 0 1 1.5 2h12A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5zM3 6a2 2 0 1 1 4 0a2 2 0 0 1-4 0m9 0H9V5h3zm0 3H9V8h3zM5 9a2.927 2.927 0 0 0-2.618 1.618l-.33.658A.5.5 0 0 0 2.5 12h5a.5.5 0 0 0 .447-.724l-.329-.658A2.927 2.927 0 0 0 5 9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImacOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 14.5h11m-5.5-4v4M0 7.5h15M.5 1.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImacSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 0A1.5 1.5 0 0 1 15 1.5V7H0V1.5A1.5 1.5 0 0 1 1.5 0zM0 8v1.5A1.5 1.5 0 0 0 1.5 11H7v3H2v1h11v-1H8v-3h5.5A1.5 1.5 0 0 0 15 9.5V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 3.5h1m3.5 5.993l-3-2.998l-3 2.998l-4-4.996L.5 9.5m1-9h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4h-1V3h1z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12.01A1.5 1.5 0 0 1 13.5 15h-12A1.5 1.5 0 0 1 0 13.5zm14 6.787l-2.5-2.498l-2.959 2.956L4.5 3.696L1 8.074V1.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageDocumentOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm1 7.995l.354-.353l-.353-.354l-.354.354zm-3 2.999l-.39.312l.349.436l.394-.395zM4.5 6.5l.39-.313l-.403-.503L4.1 6.2zm8 7.5h-10v1h10zM2 13.5v-12H1v12zm11-10v10h1v-10zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM10 6h1V5h-1zm3.854 4.147l-2-2.005l-.708.707l2 2.004zm-2.707-2.005l-3 2.998l.706.707l3-2.998zM8.89 11.18l-4-4.994l-.78.626l4 4.993zM4.1 6.2l-3 4l.8.6l3-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageDocumentSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293v6.999l-2.5-2.504l-2.959 2.957L4.5 5.7L1 10.075zM11 6h-1V5h1z" clip-rule="evenodd"/><path fill="currentColor" d="M1 11.676V13.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-1.793l-2.5-2.504l-3.041 3.039L4.5 7.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 3.5l.354-.354a.5.5 0 0 0-.708 0zM1.5 1h12V0h-12zm12.5.5v12h1v-12zM13.5 14h-12v1h12zM1 13.5v-12H0v12zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5zm-1 11h14v-1H.5zm.354-3.146l4-4l-.708-.708l-4 4zm3.292-4l7 7l.708-.708l-7-7zM10.5 5a.5.5 0 0 1-.5-.5H9A1.5 1.5 0 0 0 10.5 6zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 12 4.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 3zm0-1A1.5 1.5 0 0 0 9 4.5h1a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.505 1.505 0 0 1-.178.71A1.5 1.5 0 0 1 13.5 15h-12A1.5 1.5 0 0 1 0 13.5zm4.85 1.642l-.35-.35l-3.5 3.5V1.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5V10h-2.293L4.854 3.146z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InEarHeadphonesOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 12.5v.5h.5v-.5zm-2 0H12v.5h.5zm-12 0H0v.5h.5zm2 0v.5H3v-.5zM10 5V1.654H9V5zm4-1.5v9h1v-9zm.5 8.5h-2v1h2zm-2-6H11v1h1.5zm-1.846-5h.846V0h-.846zM11 6a1 1 0 0 1-1-1H9a2 2 0 0 0 2 2zm4-2.5A3.5 3.5 0 0 0 11.5 0v1A2.5 2.5 0 0 1 14 3.5zm-5-1.846c0-.361.293-.654.654-.654V0C9.74 0 9 .74 9 1.654zM13 12.5V15h1v-2.5zM6 5V1.654H5V5zM0 3.5v9h1v-9zM.5 13h2v-1h-2zm2-6H4V6H2.5zm1.846-7H3.5v1h.846zM4 7a2 2 0 0 0 2-2H5a1 1 0 0 1-1 1zM1 3.5A2.5 2.5 0 0 1 3.5 1V0A3.5 3.5 0 0 0 0 3.5zm5-1.846C6 .74 5.26 0 4.346 0v1c.361 0 .654.293.654.654zM1 12.5V15h1v-2.5zM12 4v2.5h1V4zM2 4v2.5h1V4zm11 8.5v-6h-1v6zm-10 0v-6H2v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InEarHeadphonesSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 1.654C6 .74 5.26 0 4.346 0H3.5A3.5 3.5 0 0 0 0 3.5V13h1v2h1v-2h1V7h1a2 2 0 0 0 2-2zM10.654 0C9.74 0 9 .74 9 1.654V5a2 2 0 0 0 2 2h1v6h1v2h1v-2h1V3.5A3.5 3.5 0 0 0 11.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.713 11.493l-.035-.5zM1.5 1h12V0h-12zm12.5.5v12h1v-12zM13.5 14h-12v1h12zM1 13.5v-12H0v12zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5zm6 12c.083 0 .166-.003.248-.009l-.07-.997A2.546 2.546 0 0 1 7.5 11zm-.823-.098c.264.064.54.098.823.098v-1c-.203 0-.4-.024-.589-.07zm.234-.972c-.969-.233-1.9-.895-2.97-1.586C2.924 8.687 1.771 8 .5 8v1c.938 0 1.858.512 2.899 1.184c.987.638 2.099 1.434 3.278 1.719zm.837 1.061c1.386-.097 2.7-.927 3.865-1.632C12.843 9.616 13.922 9 15 9V8c-1.407 0-2.732.794-3.905 1.503c-1.237.749-2.324 1.414-3.417 1.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5V8h-4.5a.5.5 0 0 0-.5.5a2.5 2.5 0 0 1-5 0a.5.5 0 0 0-.5-.5H0z"/><path fill="currentColor" d="M0 9v4.5A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5V9h-4.035a3.5 3.5 0 0 1-6.93 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentDecreaseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 7.5l-.354-.354a.5.5 0 0 0 0 .708zM3 4h12V3H3zm4 4h8V7H7zm-4 4h12v-1H3zm-.146-2.854l-2-2l-.708.708l2 2zm-2-1.292l2-2l-.708-.708l-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentDecreaseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 4H3V3h12zM1.207 7.5l1.647-1.646l-.708-.708l-2 2a.5.5 0 0 0 0 .708l2 2l.708-.708zM15 8H7V7h8zm0 4H3v-1h12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentIncreaseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.5 7.5l.354.354a.5.5 0 0 0 0-.708zM3 4h12V3H3zm4 4h8V7H7zm-4 4h12v-1H3zM.854 9.854l2-2l-.708-.708l-2 2zm2-2.708l-2-2l-.708.708l2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentIncreaseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 4H3V3h12zM.854 5.146l2 2a.5.5 0 0 1 0 .708l-2 2l-.708-.708L1.793 7.5L.146 5.854zM15 8H7V7h8zm0 4H3v-1h12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 4.5V5h1v-.5zm1-.01v-.5H7v.5zM8 11V7H7v4zm0-6.5v-.01H7v.01zM6 8h1.5V7H6zm0 3h3v-1H6zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zm0 1A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0M7 5V3.99h1V5zm1 2v3h1v1H6v-1h1V8H6V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1.5V2h1v-.5zm1-.01v-.5H7v.5zM8 13.5V4H7v9.5zm0-12v-.01H7v.01zM4 5h3.5V4H4zm-2 9h11v-1H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 4.5V5h1v-.5zm1-.01v-.5H7v.5zM8 11V7H7v4zm0-6.5v-.01H7v.01zM6 8h1.5V7H6zm0 3h3v-1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3.99V5H7V3.99zM6 11v-1h1V8H6V7h2v3h1v1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 .99V2H7V.99zM7 13H2v1h11v-1H8V4H4v1h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 3.5h1M4.5.5h6a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4h-6a4 4 0 0 1-4-4v-6a4 4 0 0 1 4-4Zm3 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5"/><path fill="currentColor" fill-rule="evenodd" d="M4.5 0A4.5 4.5 0 0 0 0 4.5v6A4.5 4.5 0 0 0 4.5 15h6a4.5 4.5 0 0 0 4.5-4.5v-6A4.5 4.5 0 0 0 10.5 0zM4 7.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0M11 4h1V3h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InvoiceOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 6.995H4v1h.5zm6 1h.5v-1h-.5zM4.5 9.5H4v.5h.5zm6 0v.5h.5v-.5zm-6-4V5H4v.5zm6 0h.5V5h-.5zm3-2h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-6 7.495h6v-1h-6zM4.5 10h6V9h-6zm0-4h6V5h-6zm8 8h-10v1h10zM2 13.5v-12H1v12zm11-10v10h1v-10zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zm2 4v4h1v-4zm3 0v4h1v-4zm3 0v4h1v-4zM4 4h3V3H4zm4 8h3v-1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InvoiceSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 7.995V9H8V7.995zM10 6v.995H8V6zM7 6H5v.995h2zm0 1.995H5V9h2z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM4 4h3V3H4zm7 1H4v5h7zm0 7H8v-1h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 1.5h9m-11 12h9m-2.5-12l-2 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ItalicSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.91 2H4V1h9v1H8.924L7.09 13H11v1H2v-1h4.076z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JavascriptOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 8v-.167c0-.736-.597-1.333-1.333-1.333H10a1.5 1.5 0 1 0 0 3h1a1.5 1.5 0 0 1 0 3h-1A1.5 1.5 0 0 1 8.5 11m-2-5v5a1.5 1.5 0 0 1-3 0M.5.5h14v14H.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JavascriptSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 0H0v15h15zM8 8a2 2 0 0 1 2-2h1.167C12.179 6 13 6.82 13 7.833V8h-1v-.167A.833.833 0 0 0 11.167 7H10a1 1 0 0 0 0 2h1a2 2 0 1 1 0 4h-1a2 2 0 0 1-2-2h1a1 1 0 0 0 1 1h1a1 1 0 1 0 0-2h-1a2 2 0 0 1-2-2M6 6v5a1 1 0 1 1-2 0H3a2 2 0 1 0 4 0V6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoystickOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 10.5V10a.5.5 0 0 0-.5.5zm12 0h.5a.5.5 0 0 0-.5-.5zm0 4v.5a.5.5 0 0 0 .5-.5zm-12 0H1a.5.5 0 0 0 .5.5zm0-3.5h12v-1h-12zm11.5-.5v4h1v-4zm.5 3.5h-12v1h12zM2 14.5v-4H1v4zm6-4v-4H7v4zM7.5 0A3.5 3.5 0 0 0 4 3.5h1A2.5 2.5 0 0 1 7.5 1zM11 3.5A3.5 3.5 0 0 0 7.5 0v1A2.5 2.5 0 0 1 10 3.5zM7.5 7A3.5 3.5 0 0 0 11 3.5h-1A2.5 2.5 0 0 1 7.5 6zm0-1A2.5 2.5 0 0 1 5 3.5H4A3.5 3.5 0 0 0 7.5 7zM3 9h2V8H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoystickSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3.5a3.5 3.5 0 1 1 4 3.465V10h5.5a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-12a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5H7V6.965A3.5 3.5 0 0 1 4 3.5"/><path fill="currentColor" d="M3 8v1h2V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JpgOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 6.5V6H6v.5zm4 4H10v.5h.5zm2 0v.5h.5v-.5zm1-7h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-6 6H5V6h-.5zm0 4v.5H5v-.5zm-2 0H2v.5h.5zm4-3.5h1V6h-1zm.5 4V8.5H6V11zm0-2.5v-2H6v2zm.5-.5h-1v1h1zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 7.5zM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6zM10 6v4.5h1V6zm.5 5h2v-1h-2zm2.5-.5v-2h-1v2zM10.5 7H13V6h-2.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zM2 7h2.5V6H2zm2-.5v4h1v-4zm.5 3.5h-2v1h2zm-1.5.5V9H2v1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JpgSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8h.5a.5.5 0 0 0 0-1H7z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM4 7H2V6h3v5H2V9h1v1h1zm2-1h1.5a1.5 1.5 0 1 1 0 3H7v2H6zm4 0h3v1h-2v3h1V8.5h1V11h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KanbanOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 3.5V3a.5.5 0 0 0-.5.5zm6 0H7a.5.5 0 0 0-.5-.5zm0 11v.5a.5.5 0 0 0 .5-.5zm-6 0H0a.5.5 0 0 0 .5.5zm8-11V3a.5.5 0 0 0-.5.5zm6 0h.5a.5.5 0 0 0-.5-.5zm0 6v.5a.5.5 0 0 0 .5-.5zm-6 0H8a.5.5 0 0 0 .5.5zM0 1h7V0H0zm8 0h7V0H8zM.5 4h6V3h-6zM6 3.5v11h1v-11zM6.5 14h-6v1h6zm-5.5.5v-11H0v11zM8.5 4h6V3h-6zm5.5-.5v6h1v-6zm.5 5.5h-6v1h6zM9 9.5v-6H8v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KanbanSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h7V0H0zm8 0h7V0H8zM.5 3a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-11a.5.5 0 0 0-.5-.5zm8 0a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5 14.5l8-8m-6 6l2 2m0-4l2 2m4.5-5a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 4a4 4 0 1 1 1.547 3.16l-3.34 3.34l1.647 1.646l-.708.708L4.5 11.207L3.207 12.5l1.647 1.646l-.708.708L2.5 13.207L.854 14.854l-.708-.708L7.84 6.453A3.983 3.983 0 0 1 7 4m4-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 11.5H4m7-3h-1m-2 0H7m-2 0H4m7-2h-1m-2 0H7m-2 0H4m3.5-2V0m6 4.5h-12a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 0v4h5.5A1.5 1.5 0 0 1 15 5.5v7a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-7A1.5 1.5 0 0 1 1.5 4H7V0zm2 6h1v1h-1zm1 2h-1v1h1zm0 3H4v1h7zM7 8h1v1H7zM5 8H4v1h1zm3-2H7v1h1zM4 6h1v1H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanCableOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 11.5V15m2-3.5V15M6 9.5h3M6.5.5v2h2v-2m-4 0h6v4h1v3l-2 4h-4l-2-4v-3h1zm2 4v2a1 1 0 0 0 2 0v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LanCableSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 6.5V5h1v1.5a.5.5 0 0 1-1 0"/><path fill="currentColor" fill-rule="evenodd" d="M9 0h2v4h1v3.618L9.809 12H9v3H8v-3H7v3H6v-3h-.809L3 7.618V4h1V0h2v3h3zm0 4H6v2.5a1.5 1.5 0 1 0 3 0zM6 9v1h3V9z" clip-rule="evenodd"/><path fill="currentColor" d="M8 0H7v2h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 13.5h15M1.5 2.5v8a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 1A1.5 1.5 0 0 0 1 2.5v8A1.5 1.5 0 0 0 2.5 12h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 12.5 1zM0 14h15v-1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaravelOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2 3.571l-.158-.474a.5.5 0 0 0-.327.596zm2.357-.785l.447-.224a.5.5 0 0 0-.605-.25zm4.714 9.428l-.447.224a.5.5 0 0 0 .705.205zM13 9.857l.257.429a.5.5 0 0 0 .165-.697zm-2.75-4.321l-.121-.485a.5.5 0 0 0-.3.753zm1.571-.393l.4-.3a.5.5 0 0 0-.52-.185zM13 6.714l.158.475a.5.5 0 0 0 .242-.775zM3.571 9.857l-.485.121a.5.5 0 0 0 .644.354zM2.158 4.046l2.357-.786l-.316-.949l-2.357.786zM3.91 3.009l2.918 5.837l.895-.447l-2.919-5.837zm2.918 5.837l1.796 3.592l.895-.447l-1.796-3.592zm2.5 3.797l3.93-2.357l-.515-.858l-3.929 2.358zm4.094-3.054l-1.65-2.593l-.844.537l1.65 2.593zm-1.65-2.593l-1.1-1.729l-.844.537l1.1 1.729zm-1.4-.975l1.57-.393l-.242-.97l-1.571.393zm1.05-.578L12.6 7.014l.8-.6l-1.179-1.571zm1.42.797l-1.65.55l.316.949l1.65-.55zm-1.65.55L7.117 8.148l.317.949l4.074-1.358zM7.117 8.148L3.413 9.383l.317.949l3.704-1.235zm-3.06 1.588L2.484 3.45l-.97.243l1.571 6.285zM2.5 1h10V0h-10zM14 2.5v10h1v-10zM12.5 14h-10v1h10zM1 12.5v-10H0v10zM2.5 14A1.5 1.5 0 0 1 1 12.5H0A2.5 2.5 0 0 0 2.5 15zM14 12.5a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5zM12.5 1A1.5 1.5 0 0 1 14 2.5h1A2.5 2.5 0 0 0 12.5 0zm-10-1A2.5 2.5 0 0 0 0 2.5h1A1.5 1.5 0 0 1 2.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaravelSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.104 3.397L6.57 8.33l-2.645.882L2.597 3.9zm5.173 8.111L7.981 8.915l3.157-1.053l1.165 1.83zm2.907-5.048l-.622.207l-.518-.814l.577-.145z"/><path fill="currentColor" fill-rule="evenodd" d="M0 2.5A2.5 2.5 0 0 1 2.5 0h10A2.5 2.5 0 0 1 15 2.5v10a2.5 2.5 0 0 1-2.5 2.5h-10A2.5 2.5 0 0 1 0 12.5zm4.804.062a.5.5 0 0 0-.605-.25l-2.357.785a.5.5 0 0 0-.327.596l1.571 6.285a.5.5 0 0 0 .644.354l3.292-1.098l1.602 3.204a.5.5 0 0 0 .705.205l3.928-2.357a.5.5 0 0 0 .165-.697l-1.306-2.053l1.042-.347a.5.5 0 0 0 .242-.775l-1.178-1.571a.5.5 0 0 0-.522-.185l-1.571.393a.5.5 0 0 0-.3.753l.755 1.188L7.53 8.011z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersDifferenceOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 4.5v-3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-3m-6-6h-3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-3m-6-6V7m0-2.5H7m3.5 6H8m2.5 0V8M8 4.5h2.5V7m-6 1v2.5H7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersDifferenceSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 4V1.5A1.5 1.5 0 0 1 5.5 0h8A1.5 1.5 0 0 1 15 1.5v8a1.5 1.5 0 0 1-1.5 1.5H11v2.5A1.5 1.5 0 0 1 9.5 15h-8A1.5 1.5 0 0 1 0 13.5v-8A1.5 1.5 0 0 1 1.5 4zm6 1v5H5V5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersIntersectOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 3V1.5a1 1 0 0 1 1-1H7m5 0h1.5a1 1 0 0 1 1 1V3M8 .5h3m1 10h1.5a1 1 0 0 0 1-1V8m0-4v3M3 4.5H1.5a1 1 0 0 0-1 1V7m0 5v1.5a1 1 0 0 0 1 1H3M.5 8v3M8 14.5h1.5a1 1 0 0 0 1-1V12M4 14.5h3m-2.5-10v6h6v-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersIntersectSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 1.5V4H1.5A1.5 1.5 0 0 0 0 5.5v8A1.5 1.5 0 0 0 1.5 15h8a1.5 1.5 0 0 0 1.5-1.5V11h2.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-8A1.5 1.5 0 0 0 4 1.5M5.5 1a.5.5 0 0 0-.5.5V4h4.5A1.5 1.5 0 0 1 11 5.5V10h2.5a.5.5 0 0 0 .5-.5v-8a.5.5 0 0 0-.5-.5zm0 10A1.5 1.5 0 0 1 4 9.5V5H1.5a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 1.5l.197-.46a.5.5 0 0 0-.394 0zm-7 3l-.197-.46a.5.5 0 0 0 0 .92zm7 3l-.197.46a.5.5 0 0 0 .394 0zm7-3l.197.46a.5.5 0 0 0 0-.92zm-7 6l-.197.46l.197.084l.197-.084zm0 3l-.197.46l.197.084l.197-.084zM7.303 1.04l-7 3l.394.92l7-3zm-7 3.92l7 3l.394-.92l-7-3zm7.394 3l7-3l-.394-.92l-7 3zm7-3.92l-7-3l-.394.92l7 3zM.303 7.96l7 3l.394-.92l-7-3zm7.394 3l7-3l-.394-.92l-7 3zm-7.394 0l7 3l.394-.92l-7-3zm7.394 3l7-3l-.394-.92l-7 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.697 1.04a.5.5 0 0 0-.394 0l-7 3a.5.5 0 0 0 0 .92l7 3a.5.5 0 0 0 .394 0l7-3a.5.5 0 0 0 0-.92z"/><path fill="currentColor" d="M7.5 9.956L.697 7.04l-.394.92L7.5 11.044l7.197-3.084l-.394-.92z"/><path fill="currentColor" d="m.697 10.04l-.394.92L7.5 14.044l7.197-3.084l-.394-.92L7.5 12.956z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersSubtractOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.5 10.5v3a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h3m0-3v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-8a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersSubtractSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 4H1.5A1.5 1.5 0 0 0 0 5.5v8A1.5 1.5 0 0 0 1.5 15h8a1.5 1.5 0 0 0 1.5-1.5V11h2.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0h-8A1.5 1.5 0 0 0 4 1.5zm1-2.5a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersUnionOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 1.5v3h-3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-3h3a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-8a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersUnionSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 0A1.5 1.5 0 0 0 4 1.5V4H1.5A1.5 1.5 0 0 0 0 5.5v8A1.5 1.5 0 0 0 1.5 15h8a1.5 1.5 0 0 0 1.5-1.5V11h2.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.854 4.854l.353-.354l-.707-.707l-.354.353zM5.5 7.5l-.354-.354l-.353.354l.353.354zm2.646 3.354l.354.353l.707-.707l-.353-.354zm0-6.708l-3 3l.708.708l3-3zm-3 3.708l3 3l.708-.708l-3-3zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m9 4.207L4.793 7.5L9 3.293z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M10 14L3 7.5L10 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m8.5 4.5l-3 3l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.793 7.5L9 11.707V3.293z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7.5L11 0v15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LegoOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 3h8V2h-8zM13 4.5v8h1v-8zM11.5 14h-8v1h8zM2 12.5v-8H1v8zM3.5 14A1.5 1.5 0 0 1 2 12.5H1A2.5 2.5 0 0 0 3.5 15zm9.5-1.5a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5zM11.5 3A1.5 1.5 0 0 1 13 4.5h1A2.5 2.5 0 0 0 11.5 2zm-8-1A2.5 2.5 0 0 0 1 4.5h1A1.5 1.5 0 0 1 3.5 3zM5 8h1V7H5zm4 0h1V7H9zm1.1 1.7a3.25 3.25 0 0 1-5.2 0l-.8.6c1.7 2.267 5.1 2.267 6.8 0zM5 2.5v-1H4v1zM5.5 1h4V0h-4zm4.5.5v1h1v-1zM9.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 0zM5 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 4 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LegoSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 1.5A1.5 1.5 0 0 1 5.5 0h4A1.5 1.5 0 0 1 11 1.5V2h.5A2.5 2.5 0 0 1 14 4.5v8a2.5 2.5 0 0 1-2.5 2.5h-8A2.5 2.5 0 0 1 1 12.5v-8A2.5 2.5 0 0 1 3.5 2H4zM5 7v1h1V7zm4 0v1h1V7zM4.9 9.7a3.25 3.25 0 0 0 5.2 0l.8.6c-1.7 2.267-5.1 2.267-6.8 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifebuoyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.329 10.328a4 4 0 0 1-5.657 0m5.657 0a4 4 0 0 0 0-5.656m0 5.656l2.12 2.122m-7.777-2.122a4 4 0 0 1 0-5.656m0 5.656L2.55 12.45m7.779-7.778a4 4 0 0 0-5.657 0m5.657 0l2.12-2.122M4.673 4.672L2.55 2.55m9.9 9.9a7 7 0 0 1-9.9 0m9.9 0a7 7 0 0 0 0-9.9m-9.9 9.9a7 7 0 0 1 0-9.9m9.9 0a7 7 0 0 0-9.9 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifebuoySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.404 4.697L2.562 1.855a7.501 7.501 0 0 1 9.876 0L9.596 4.697a3.502 3.502 0 0 0-4.192 0m4.899.707a3.502 3.502 0 0 1 0 4.192l2.842 2.842a7.501 7.501 0 0 0 0-9.876zm-.707 4.899a3.502 3.502 0 0 1-4.192 0l-2.842 2.842a7.501 7.501 0 0 0 9.876 0zM4.697 5.404a3.502 3.502 0 0 0 0 4.192l-2.842 2.842a7.501 7.501 0 0 1 0-9.876z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightningCableOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 5.5h6m-6 0a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1m-6 0v-4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4M9 2.5H6M5.5 13v2m4-2v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightningCableSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 0A1.5 1.5 0 0 0 4 1.5V5h7V1.5A1.5 1.5 0 0 0 9.5 0zM6 2h3v1H6z" clip-rule="evenodd"/><path fill="currentColor" d="M3 6h9v5.5a1.5 1.5 0 0 1-1.5 1.5H10v2H9v-2H6v2H5v-2h-.5A1.5 1.5 0 0 1 3 11.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m2 2l11 11M1.5 2.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm12 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0a1.5 1.5 0 1 0 .647 2.854l10 10a1.5 1.5 0 1 0 .707-.707l-10-10A1.5 1.5 0 0 0 1.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 6.5L1.328 9.672a2.828 2.828 0 1 0 4 4L8.5 10.5m2-2l3.172-3.172a2.829 2.829 0 0 0-4-4L6.5 4.5m-2 6l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkRemoveOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 6.5L1.328 9.672a2.828 2.828 0 1 0 4 4L8.5 10.5m2-2l3.172-3.172a2.829 2.829 0 0 0-4-4L6.5 4.5m-2 6l6-6M3 4.5H0M4.5 0v3m6 9v3m1.5-4.5h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkRemoveSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3V0h1v3zM9.318.975a3.328 3.328 0 1 1 4.707 4.707l-3.171 3.172l-.708-.708l3.172-3.171a2.328 2.328 0 1 0-3.293-3.293L6.854 4.854l-.708-.708zM0 4h3v1H0zm10.854.854l-6 6l-.708-.708l6-6zm-6 2l-3.172 3.171a2.329 2.329 0 0 0 3.293 3.293l3.171-3.172l.708.708l-3.172 3.171A3.328 3.328 0 1 1 .975 9.318l3.171-3.172zM15 11h-3v-1h3zm-5 4v-3h1v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.318.975a3.328 3.328 0 1 1 4.707 4.707l-3.171 3.172l-.708-.708l3.172-3.171a2.328 2.328 0 1 0-3.293-3.293L6.854 4.854l-.708-.708zm1.536 3.879l-6 6l-.708-.708l6-6zm-6 2l-3.172 3.171a2.329 2.329 0 0 0 3.293 3.293l3.171-3.172l.708.708l-3.172 3.171A3.328 3.328 0 1 1 .975 9.318l3.171-3.172z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 6v5m6 0V8.5a2 2 0 1 0-4 0V11V6M4 4.5h1M1.5.5h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5zM5 5H4V4h1zm-1 6V6h1v5zm4.5-4A1.5 1.5 0 0 0 7 8.5V11H6V6h1v.5a2.5 2.5 0 0 1 4 2V11h-1V8.5A1.5 1.5 0 0 0 8.5 7" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinuxAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 9.662c0-.758.224-1.498.645-2.129l.565-.847a7.203 7.203 0 0 0 1.07-2.583l.328-1.642a2.44 2.44 0 0 1 4.784 0l.329 1.642a7.18 7.18 0 0 0 1.07 2.583l.564.847c.42.63.645 1.371.645 2.129m-7.392 3.637c.386.13.8.201 1.23.201h2.324c.43 0 .844-.07 1.23-.201M6.5 5.5L4.947 8.606a2 2 0 0 0 1.79 2.894h1.527a2 2 0 0 0 1.789-2.894L8.5 5.5M6.5 3v1.5m2-1.5v1.5m-1.894-.053L5.5 5l.586.586a2 2 0 0 0 2.828 0L9.5 5l-1.106-.553a2 2 0 0 0-1.788 0ZM.77 11.326l.479-1.197a1 1 0 0 1 .928-.629h1.164a1 1 0 0 1 .919.606l.93 2.172a1 1 0 0 1-.319 1.194l-.738.553a1 1 0 0 1-1.24-.031l-1.835-1.529a1 1 0 0 1-.288-1.14Zm13.46 0l-.479-1.197a1 1 0 0 0-.928-.629h-1.164a1 1 0 0 0-.919.606l-.93 2.172a1 1 0 0 0 .319 1.194l.738.553a1 1 0 0 0 1.24-.031l1.835-1.529a1 1 0 0 0 .288-1.14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinuxAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.54 2.407A3.013 3.013 0 0 1 7.5 0a3.013 3.013 0 0 1 2.96 2.407l.338 1.672a6.796 6.796 0 0 0 1.022 2.449l.58.863c.363.539.6 1.148.698 1.782a1.54 1.54 0 0 1 1.3.955l.492 1.22a1.52 1.52 0 0 1-.445 1.74l-1.884 1.558a1.55 1.55 0 0 1-1.91.048l-.76-.564a1.256 1.256 0 0 1-.029-.023c-.372.1-.764.154-1.168.154H6.306c-.404 0-.796-.053-1.168-.154l-.03.023l-.759.564a1.55 1.55 0 0 1-1.91-.048L.554 13.089a1.52 1.52 0 0 1-.444-1.742l.492-1.219a1.54 1.54 0 0 1 1.3-.955a4.397 4.397 0 0 1 .697-1.782l.58-.863a6.797 6.797 0 0 0 1.023-2.449zm6.305 7.069a1.53 1.53 0 0 0-.489.618l-.768 1.777a2.581 2.581 0 0 1-1.303.353h-1.57c-.467 0-.915-.126-1.303-.353l-.768-1.777a1.528 1.528 0 0 0-.489-.618c.026-.322.114-.641.264-.938l1.258-2.495l.007.007a2.583 2.583 0 0 0 3.632 0l.007-.007l1.258 2.495c.15.297.238.616.264.938M9.04 4.269V3.056H8.014v.801c.218.044.431.117.634.218zm-2.055-.412v-.801H5.96v1.213l.393-.194c.203-.101.417-.174.634-.218m-4.953 6.33a.514.514 0 0 0-.477.32l-.492 1.219a.507.507 0 0 0 .148.58l1.884 1.557a.517.517 0 0 0 .637.017l.759-.565a.507.507 0 0 0 .164-.608L3.7 10.495a.514.514 0 0 0-.472-.309zm11.411.32a.514.514 0 0 0-.477-.32h-1.195a.514.514 0 0 0-.472.308l-.956 2.212a.507.507 0 0 0 .164.608l.759.565c.19.141.454.134.637-.017l1.884-1.557a.507.507 0 0 0 .148-.58zM6.811 4.986a1.552 1.552 0 0 1 1.378 0l.498.247l-.097.097a1.55 1.55 0 0 1-2.18 0l-.097-.097z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinuxOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 14.5l-.354-.354A.5.5 0 0 0 .5 15zm.656-.656l-.353-.354zM14.5 14.5v.5a.5.5 0 0 0 .354-.854zm-11-4l-.224-.447a.5.5 0 0 0-.13.8zm.87-.435l.223.447zm6.26 0l-.223.447zm.87.435l.354.354a.5.5 0 0 0-.13-.801zm-7.414.586l.353-.354zM.854 14.854l.656-.657l-.707-.707l-.657.656zM2 13.014V6.5H1v6.514zM13 6.5v6.514h1V6.5zm.49 7.697l.656.657l.708-.708l-.657-.656zM14.5 14H.5v1h14zm-1.5-.986c0 .444.176.87.49 1.183l.707-.707a.673.673 0 0 1-.197-.476zM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0zM2 6.5A5.5 5.5 0 0 1 7.5 1V0A6.5 6.5 0 0 0 1 6.5zm-.49 7.697c.314-.314.49-.74.49-1.183H1c0 .179-.07.35-.197.476zm2.214-3.25l.87-.434l-.448-.895l-.87.435zm6.683-.434l.87.434l.447-.894l-.87-.435zm.74-.367l-.586.586l.707.707l.586-.585zm-6.708.586l-.585-.586l-.708.708l.586.585zM7.5 12a4.328 4.328 0 0 1-3.06-1.268l-.708.707c1 1 2.355 1.561 3.768 1.561zm3.06-1.268A4.328 4.328 0 0 1 7.5 12v1a5.328 5.328 0 0 0 3.768-1.56zm-5.967-.22a6.5 6.5 0 0 1 5.814 0l.447-.894a7.5 7.5 0 0 0-6.708 0zM7 6.5v1h1v-1zm-3 1v-1H3v1zM5.5 9A1.5 1.5 0 0 1 4 7.5H3A2.5 2.5 0 0 0 5.5 10zM7 7.5A1.5 1.5 0 0 1 5.5 9v1A2.5 2.5 0 0 0 8 7.5zM5.5 5A1.5 1.5 0 0 1 7 6.5h1A2.5 2.5 0 0 0 5.5 4zm0-1A2.5 2.5 0 0 0 3 6.5h1A1.5 1.5 0 0 1 5.5 5zM11 6.5v1h1v-1zM9.5 9A1.5 1.5 0 0 1 8 7.5H7A2.5 2.5 0 0 0 9.5 10zM11 7.5A1.5 1.5 0 0 1 9.5 9v1A2.5 2.5 0 0 0 12 7.5zM9.5 5A1.5 1.5 0 0 1 11 6.5h1A2.5 2.5 0 0 0 9.5 4zm0-1A2.5 2.5 0 0 0 7 6.5h1A1.5 1.5 0 0 1 9.5 5zM5 7v1h1V7zm4 0v1h1V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinuxSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8V7h1v1zm4 0V7h1v1z"/><path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 0 1 13 0v6.514c0 .179.07.35.197.476l.657.656A.5.5 0 0 1 14.5 15H.5a.5.5 0 0 1-.354-.854l.657-.656A.673.673 0 0 0 1 13.014zm3 0a1.5 1.5 0 1 1 3 0v1a1.5 1.5 0 1 1-3 0zm4 0a1.5 1.5 0 1 1 3 0v1a1.5 1.5 0 0 1-3 0zm-3.407 4.012a6.5 6.5 0 0 1 5.814 0l.249.125l-.095.095a4.329 4.329 0 0 1-6.122 0l-.095-.095z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListLayoutOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7 3.5h8m-8 8h8M1.5 1.5h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1Zm0 8h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListLayoutSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 1A1.5 1.5 0 0 0 0 2.5v2A1.5 1.5 0 0 0 1.5 6h2A1.5 1.5 0 0 0 5 4.5v-2A1.5 1.5 0 0 0 3.5 1zM7 4h8V3H7zM1.5 9A1.5 1.5 0 0 0 0 10.5v2A1.5 1.5 0 0 0 1.5 14h2A1.5 1.5 0 0 0 5 12.5v-2A1.5 1.5 0 0 0 3.5 9zM7 12h8v-1H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOrderedOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 13.5l-.354-.354A.5.5 0 0 0 .5 14zm1-11H2V2h-.5zM5 8h10V7H5zm0-4h10V3H5zm0 8h10v-1H5zm-2 1H.5v1H3zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792zM1.793 10H1.5v1h.293zM1.5 10A1.5 1.5 0 0 0 0 11.5h1a.5.5 0 0 1 .5-.5zM3 11.207C3 10.54 2.46 10 1.793 10v1c.114 0 .207.093.207.207zm-.354.854c.227-.227.354-.534.354-.854H2a.207.207 0 0 1-.06.147zM0 6h3V5H0zm2-.5v-3H1v3zM1.5 2H0v1h1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOrderedSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2h2v3h1v1H0V5h1V3H0zm15 2H5V3h10zm0 4H5V7h10zM0 11.5A1.5 1.5 0 0 1 1.5 10h.293a1.207 1.207 0 0 1 .853 2.06l-.939.94H3v1H.5a.5.5 0 0 1-.354-.854l1.793-1.792A.207.207 0 0 0 1.793 11H1.5a.5.5 0 0 0-.5.5zm15 .5H5v-1h10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUnorderedOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 7.5h11m-15 0h2m2-4h11m-15 0h2m2 8h11m-15 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUnorderedSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4H0V3h2zm13 0H4V3h11zM2 8H0V7h2zm13 0H4V7h11zM2 12H0v-1h2zm13 0H4v-1h11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LitecoinOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m5.5 1.5l-3 12H13m-12-5l6-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LitecoinSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.714 6.584l1.3-5.205l.971.242l-1.093 4.374l1.884-.942l.448.894l-2.652 1.326L3.14 13H13v1H1.86l1.534-6.138l-2.17 1.085l-.448-.894z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoaderOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1V.5H7V1zM7 4.5V5h1v-.5zm1 6V10H7v.5zM7 14v.5h1V14zM4.5 7.995H5v-1h-.5zm-3.5-1H.5v1H1zM14 8h.5V7H14zm-3.5-1.005H10v1h.5zM7 1v3.5h1V1zm0 9.5V14h1v-3.5zM4.5 6.995H1v1h3.5zM14 7l-3.5-.005v1L14 8zM2.147 2.854l3 3l.708-.708l-3-3zm10-.708l-3 3l.708.708l3-3zM2.854 12.854l3-3l-.708-.708l-3 3zm6.292-3l3 3l.708-.708l-3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoaderSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 .5V5H7V.5zM5.146 5.854l-3-3l.708-.708l3 3zm4-.708l3-3l.708.708l-3 3zm.855 1.849L14.5 7l-.002 1l-4.5-.006zm-9.501 0H5v1H.5zm5.354 2.859l-3 3l-.708-.708l3-3zm6.292 3l-3-3l.708-.708l3 3zM8 10v4.5H7V10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5.5v14m7-7.005H.5m13 0a6.006 6.006 0 0 1-6 6.005c-3.313 0-6-2.694-6-6.005a5.999 5.999 0 0 1 6-5.996a6 6 0 0 1 6 5.996Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.018V0H7v1.018a6.5 6.5 0 0 0-5.981 5.977H0v1h1.019A6.508 6.508 0 0 0 7 13.981V15h1v-1.019a6.508 6.508 0 0 0 5.981-5.986H15v-1h-1.019A6.5 6.5 0 0 0 8 1.018M8 3v3.995h4v1H8V12H7V7.995H3v-1h4V3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 7h4V6h-4zm4.5.5v3h1v-3zM9.5 11h-4v1h4zM5 10.5v-3H4v3zm.5.5a.5.5 0 0 1-.5-.5H4A1.5 1.5 0 0 0 5.5 12zm4.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM9.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 6zm-4-1A1.5 1.5 0 0 0 4 7.5h1a.5.5 0 0 1 .5-.5zm.5.5v-1H5v1zm3-1v1h1v-1zM7.5 4A1.5 1.5 0 0 1 9 5.5h1A2.5 2.5 0 0 0 7.5 3zM6 5.5A1.5 1.5 0 0 1 7.5 4V3A2.5 2.5 0 0 0 5 5.5zm-5 2A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zm1 0A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 4A1.5 1.5 0 0 0 6 5.5V6h3v-.5A1.5 1.5 0 0 0 7.5 4"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 1 0 15a7.5 7.5 0 0 1 0-15M5 5.5v.585A1.5 1.5 0 0 0 4 7.5v3A1.5 1.5 0 0 0 5.5 12h4a1.5 1.5 0 0 0 1.5-1.5v-3a1.5 1.5 0 0 0-1-1.415V5.5a2.5 2.5 0 0 0-5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 6.5v-3a3 3 0 0 1 6 0v3m-8 0h10a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 7h4V6h-4zm4.5.5v3h1v-3zM9.5 11h-4v1h4zM5 10.5v-3H4v3zm.5.5a.5.5 0 0 1-.5-.5H4A1.5 1.5 0 0 0 5.5 12zm4.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM9.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 6zm-4-1A1.5 1.5 0 0 0 4 7.5h1a.5.5 0 0 1 .5-.5zm.5.5v-1H5v1zm3-1v1h1v-1zM7.5 4A1.5 1.5 0 0 1 9 5.5h1A2.5 2.5 0 0 0 7.5 3zM6 5.5A1.5 1.5 0 0 1 7.5 4V3A2.5 2.5 0 0 0 5 5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 5.5v.585A1.5 1.5 0 0 1 11 7.5v3A1.5 1.5 0 0 1 9.5 12h-4A1.5 1.5 0 0 1 4 10.5v-3a1.5 1.5 0 0 1 1-1.415V5.5a2.5 2.5 0 0 1 5 0m-4 0a1.5 1.5 0 1 1 3 0V6H6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 3.5V6h1.5A1.5 1.5 0 0 1 14 7.5v6a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-6A1.5 1.5 0 0 1 2.5 6H4V3.5a3.5 3.5 0 1 1 7 0m-6 0a2.5 2.5 0 0 1 5 0V6H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoutOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m13.5 7.5l-3 3.25m3-3.25l-3-3m3 3H4m4 6H1.5v-12H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogoutSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h7v1H2v11h6v1H1zm9.854 3.146l3.34 3.34l-3.327 3.603l-.734-.678L12.358 8H4V7h8.293l-2.147-2.146z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoopOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.5 3.5l.354.354a.5.5 0 0 0 0-.708zm-14 8l-.354-.354a.5.5 0 0 0 0 .708zM11.146.854l3 3l.708-.708l-3-3zm3 2.292l-3 3l.708.708l3-3zm-10.292 11l-3-3l-.708.708l3 3zm-3-2.292l3-3l-.708-.708l-3 3zM.5 12h11v-1H.5zM15 8.5V7h-1v1.5zM11.5 12A3.5 3.5 0 0 0 15 8.5h-1a2.5 2.5 0 0 1-2.5 2.5zm3-9h-11v1h11zM0 6.5V8h1V6.5zM3.5 3A3.5 3.5 0 0 0 0 6.5h1A2.5 2.5 0 0 1 3.5 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoopSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.293 3L11.146.854l.708-.708l3 3a.5.5 0 0 1 0 .708l-3 3l-.708-.708L13.293 4H3.5A2.5 2.5 0 0 0 1 6.5V8H0V6.5A3.5 3.5 0 0 1 3.5 3zM15 7v1.5a3.5 3.5 0 0 1-3.5 3.5H1.707l2.147 2.146l-.708.708l-3-3a.5.5 0 0 1 0-.707l3-3l.708.707L1.707 11H11.5A2.5 2.5 0 0 0 14 8.5V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagsafeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 2.5v-2h8v2m-6 9V15m4-3.5V15m-5-5.5v2h6v-2m-9-7h12v7h-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagsafeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0H3v3h9zm2 4H1v6h3v2h1v3h1v-3h3v3h1v-3h1v-2h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkdownOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.5 5.5l.354-.354A.5.5 0 0 0 2 5.5zm2 2l-.354.354l.354.353l.354-.353zm2-2H7a.5.5 0 0 0-.854-.354zm4 4l-.354.354l.354.353l.354-.353zM1.5 3h12V2h-12zm12.5.5v8h1v-8zm-.5 8.5h-12v1h12zM1 11.5v-8H0v8zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 13zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM13.5 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 2zm-12-1A1.5 1.5 0 0 0 0 3.5h1a.5.5 0 0 1 .5-.5zM3 10V5.5H2V10zm-.854-4.146l2 2l.708-.708l-2-2zm2.708 2l2-2l-.708-.708l-2 2zM6 5.5V10h1V5.5zm4-.5v4.5h1V5zM8.146 7.854l2 2l.708-.708l-2-2zm2.708 2l2-2l-.708-.708l-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkdownSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 3.5A1.5 1.5 0 0 1 1.5 2h12A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5zM10 5v3.293L8.854 7.146l-.708.708l2 2a.5.5 0 0 0 .708 0l2-2l-.707-.708L11 8.293V5zm-7.146.146A.5.5 0 0 0 2 5.5V10h1V6.707l1.5 1.5l1.5-1.5V10h1V5.5a.5.5 0 0 0-.854-.354L4.5 6.793z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m4.5 4.5l3 4l3-4m-6 0H3m1.5 0V11m6-6.5H12m-1.5 0V11M3 10.5h3m3 0h3M1.5.5h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5zM4 5H3V4h1.5a.5.5 0 0 1 .4.2l2.6 3.467l2.593-3.458A.5.5 0 0 1 10.5 4H12v1h-1v5h1v1H9v-1h1V6L7.5 9.333L5 6v4h1v1H3v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 5.5h15m-15-4h15m-15 8h15m-15 4h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 2H0V1h15zm0 4H0V5h15zm0 4H0V9h15zm0 4H0v-1h15z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222zm2 2.998l-.416.277a.5.5 0 0 0 .832 0zm2-2.998v-.5a.5.5 0 0 0-.416.222zm-4.416.277l2 2.998l.832-.555l-2-2.998zm2.832 2.998l2-2.998l-.832-.555l-2 2.998zM9.5 11.993h4v-1h-4zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5zm1.5-1.5V1.5h-1v8.994zM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5zM13.5 0h-12v1h12zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5zM0 1.5v8.994h1V1.499zm0 8.994c0 .829.671 1.499 1.5 1.499v-1a.499.499 0 0 1-.5-.5zm1.5 1.499h4v-1h-4zM5 7h5V6H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.994c0 .829-.671 1.499-1.5 1.499H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5zM5 7h5V6H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageNoAccessOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222zm2 2.998l-.416.277a.5.5 0 0 0 .832 0zm2-2.998v-.5a.5.5 0 0 0-.416.222zm-4.416.277l2 2.998l.832-.555l-2-2.998zm2.832 2.998l2-2.998l-.832-.555l-2 2.998zM9.5 11.993h4v-1h-4zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5zm1.5-1.5V1.5h-1v8.994zM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5zM13.5 0h-12v1h12zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5zM0 1.5v8.994h1V1.499zm0 8.994c0 .829.671 1.499 1.5 1.499v-1a.499.499 0 0 1-.5-.5zm1.5 1.499h4v-1h-4zM7.5 9A2.5 2.5 0 0 1 5 6.5H4A3.5 3.5 0 0 0 7.5 10zM10 6.5A2.5 2.5 0 0 1 7.5 9v1A3.5 3.5 0 0 0 11 6.5zM7.5 4A2.5 2.5 0 0 1 10 6.5h1A3.5 3.5 0 0 0 7.5 3zm0-1A3.5 3.5 0 0 0 4 6.5h1A2.5 2.5 0 0 1 7.5 4zm1.646 1.146l-4 4l.708.708l4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageNoAccessSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 4a2.5 2.5 0 0 0-2.086 3.879L8.88 4.414A2.488 2.488 0 0 0 7.5 4m0 5c-.51 0-.983-.152-1.379-.414L9.586 5.12A2.5 2.5 0 0 1 7.5 9"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.994c0 .829-.671 1.499-1.5 1.499H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5zm4 5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" d="m5.5 11.493l2 2.998l2-2.998h4a1 1 0 0 0 1-1V1.5a.999.999 0 0 0-1-.999h-12a1 1 0 0 0-1 1v8.994c0 .552.447.999 1 .999z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222zm2 2.998l-.416.277a.5.5 0 0 0 .832 0zm2-2.998v-.5a.5.5 0 0 0-.416.222zm-4.416.277l2 2.998l.832-.555l-2-2.998zm2.832 2.998l2-2.998l-.832-.555l-2 2.998zM9.5 11.993h4v-1h-4zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5zm1.5-1.5V1.5h-1v8.994zM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5zM13.5 0h-12v1h12zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5zM0 1.5v8.994h1V1.499zm0 8.994c0 .829.671 1.499 1.5 1.499v-1a.499.499 0 0 1-.5-.5zm1.5 1.499h4v-1h-4zM7 4v5h1V4zM5 7h5V6H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.994c0 .829-.671 1.499-1.5 1.499H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5zM7 9V7H5V6h2V4h1v2h2v1H8v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0C.671 0 0 .67 0 1.5v8.994c0 .829.671 1.499 1.5 1.499h3.733l1.851 2.775a.5.5 0 0 0 .832 0l1.851-2.775H13.5c.829 0 1.5-.67 1.5-1.5V1.5c0-.83-.671-1.5-1.5-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageTextAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 11.493H4v-.5h-.5zm0 2.998H3a.5.5 0 0 0 .8.4zm4-2.998v-.5h-.167l-.133.1zm-3-7.496H4v1h.5zm6 1h.5v-1h-.5zm-6 1.998H4v1h.5zm4 1H9v-1h-.5zM3 11.493v2.998h1v-2.998zm.8 3.398l4-2.998l-.6-.8l-4 2.998zm3.7-2.898h6v-1h-6zm6 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5zm1.5-1.5V1.5h-1v8.994zM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5zM13.5 0h-12v1h12zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5zM0 1.5v8.994h1V1.499zm0 8.994c0 .829.671 1.499 1.5 1.499v-1a.499.499 0 0 1-.5-.5zm1.5 1.499h2v-1h-2zm3-6.996h6v-1h-6zm0 2.998h4v-1h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageTextAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.994c0 .829-.671 1.499-1.5 1.499H7.667L3.8 14.89a.5.5 0 0 1-.8-.4v-2.498H1.5c-.829 0-1.5-.67-1.5-1.5zm4 2.497h7v1H4zm0 2.998h5v1H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageTextOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222zm2 2.998l-.416.277a.5.5 0 0 0 .832 0zm2-2.998v-.5a.5.5 0 0 0-.416.222zm-4.416.277l2 2.998l.832-.555l-2-2.998zm2.832 2.998l2-2.998l-.832-.555l-2 2.998zM9.5 11.993h4v-1h-4zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5zm1.5-1.5V1.5h-1v8.994zM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5zM13.5 0h-12v1h12zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5zM0 1.5v8.994h1V1.499zm0 8.994c0 .829.671 1.499 1.5 1.499v-1a.499.499 0 0 1-.5-.5zm1.5 1.499h4v-1h-4zM5 8h5V7H5zM4 5h7V4H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageTextSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.994c0 .829-.671 1.499-1.5 1.499H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5zM11 5H4V4h7zm-1 3H5V7h5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageTickOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222zm2 2.998l-.416.277a.5.5 0 0 0 .832 0zm2-2.998v-.5a.5.5 0 0 0-.416.222zM7 8l-.354.354l.378.377l.352-.402zm-1.916 3.77l2 2.998l.832-.555l-2-2.998zm2.832 2.998l2-2.998l-.832-.555l-2 2.998zM9.5 11.993h4v-1h-4zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5zm1.5-1.5V1.5h-1v8.994zM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5zM13.5 0h-12v1h12zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5zM0 1.5v8.994h1V1.499zm0 8.994c0 .829.671 1.499 1.5 1.499v-1a.499.499 0 0 1-.5-.5zm1.5 1.499h4v-1h-4zm3.146-5.64l2 2l.708-.707l-2-2zm2.73 1.976l3.5-4l-.752-.658l-3.5 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageTickSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.994c0 .829-.671 1.499-1.5 1.499H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5zm7.024 7.23l3.852-4.402l-.752-.658l-3.148 3.598l-1.622-1.623l-.708.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageXoutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.5 11.493l.416-.278a.5.5 0 0 0-.416-.222zm2 2.998l-.416.277a.5.5 0 0 0 .832 0zm2-2.998v-.5a.5.5 0 0 0-.416.222zm-4.416.277l2 2.998l.832-.555l-2-2.998zm2.832 2.998l2-2.998l-.832-.555l-2 2.998zM9.5 11.993h4v-1h-4zm4 0c.829 0 1.5-.67 1.5-1.5h-1c0 .277-.223.5-.5.5zm1.5-1.5V1.5h-1v8.994zM15 1.5c0-.83-.671-1.5-1.5-1.5v1c.277 0 .5.223.5.5zM13.5 0h-12v1h12zm-12 0C.671 0 0 .67 0 1.5h1c0-.277.223-.5.5-.5zM0 1.5v8.994h1V1.499zm0 8.994c0 .829.671 1.499 1.5 1.499v-1a.499.499 0 0 1-.5-.5zm1.5 1.499h4v-1h-4zm3.646-7.14l4 4l.708-.707l-4-4zm.708 4l4-4l-.708-.707l-4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageXsolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5C0 .67.671 0 1.5 0h12c.829 0 1.5.67 1.5 1.5v8.994c0 .829-.671 1.499-1.5 1.499H9.768l-1.852 2.775a.5.5 0 0 1-.832 0l-1.851-2.775H1.5c-.829 0-1.5-.67-1.5-1.5zm9.146 7.354L7.5 7.207L5.854 8.854l-.708-.708L6.793 6.5L5.146 4.854l.708-.708L7.5 5.793l1.646-1.647l.708.708L8.207 6.5l1.647 1.646z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessengerOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.935 12.264l.482.132a.5.5 0 0 0-.164-.517zM2.326 14.5l-.482-.131a.5.5 0 0 0 .694.584zm2.435-1.141l.188-.463a.5.5 0 0 0-.4.01zM6.5 6.5l.3-.4l-.293-.22l-.298.213zm2 1.5l-.3.4l.316.237l.304-.253zm-1-8C3.379 0 0 3.201 0 7.196h1C1 3.795 3.889 1 7.5 1zM0 7.196c0 2.188 1.023 4.139 2.617 5.454l.636-.771C1.87 10.739 1 9.062 1 7.196zm2.452 4.937l-.608 2.236l.965.262l.608-2.235zm.086 2.82l2.435-1.142l-.424-.905l-2.435 1.141zm2.035-1.131c.9.366 1.89.57 2.927.57v-1a6.764 6.764 0 0 1-2.55-.496zm2.927.57c4.121 0 7.5-3.202 7.5-7.196h-1c0 3.4-2.889 6.195-6.5 6.195zM15 7.195C15 3.2 11.621 0 7.5 0v1C11.111 1 14 3.795 14 7.196zM3.29 9.406l3.5-2.5l-.58-.813l-3.5 2.5zM6.2 6.9l2 1.5l.6-.8l-2-1.5zm2.62 1.484l3-2.5l-.64-.768l-3 2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessengerSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.196C0 3.2 3.379 0 7.5 0S15 3.201 15 7.196c0 3.994-3.379 7.195-7.5 7.195a7.77 7.77 0 0 1-2.72-.489l-2.242 1.05a.5.5 0 0 1-.694-.583l.526-1.932C.918 11.129 0 9.269 0 7.196m8.516 1.441l3.304-2.753l-.64-.768l-2.696 2.247L6.507 5.88L2.71 8.593l.582.814L6.493 7.12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicroSdCardOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 3v3m2-3v3m2-3v3m-8 8.5h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-8a1 1 0 0 0-1 1v5l-2 2v5a1 1 0 0 0 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicroSdCardSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1.5A1.5 1.5 0 0 1 4.5 0h8A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V8.293l2-2zM6 3v3h1V3zm2 0v3h1V3zm2 3V3h1v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 4v2.5a5 5 0 0 0 5 5m5-7.5v2.5a5 5 0 0 1-5 5m0 0V15M5 14.5h5m-.5-12v4a2 2 0 1 1-4 0v-4a2 2 0 1 1 4 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2.5a2.5 2.5 0 0 1 5 0v4a2.5 2.5 0 0 1-5 0z"/><path fill="currentColor" d="M2 4v2.5a5.5 5.5 0 0 0 5 5.478V14H5v1h5v-1H8v-2.022A5.5 5.5 0 0 0 13 6.5V4h-1v2.5a4.5 4.5 0 0 1-9 0V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinimiseAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13 5.5H9.5m0 0V2m0 3.5l4-4M5.5 13V9.5m0 0H2m3.5 0l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinimiseAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.854 1.854L10.707 5H13v1H9V2h1v2.293l3.146-3.147zM2 9h4v4H5v-2.293l-3.146 3.147l-.708-.707L4.293 10H2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinimiseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 9.5H13m-3.5 0V13m0-3.5l4 4m-.5-8H9.5m0 0V2m0 3.5l4-4M2 5.5h3.5m0 0V2m0 3.5l-4-4m4 11.5V9.5m0 0H2m3.5 0l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinimiseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.854 1.854L10.707 5H13v1H9V2h1v2.293l3.146-3.147zM4.293 5L1.146 1.854l.708-.708L5 4.293V2h1v4H2V5zM2 9h4v4H5v-2.293l-3.146 3.147l-.708-.707L4.293 10H2zm7 0h4v1h-2.293l3.147 3.146l-.708.708L10 10.707V13H9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 7.5h7m-3.5 7a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M4 8h7V7H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 7.5h13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 7.5h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 8H4V7h7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 8H1V7h13z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 11.5h3m-5.5 3h8a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-8a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1.5A1.5 1.5 0 0 1 3.5 0h8A1.5 1.5 0 0 1 13 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 13.5zM6 12h3v-1H6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 10.5h1.5V9M11 4.5h1.5V6M4 4.5H2.5V6m0 3v1.5H4m3.5-1a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6-7h12a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M0 3.5A1.5 1.5 0 0 1 1.5 2h12A1.5 1.5 0 0 1 15 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 11.5zM4 4H2v2h1V5h1zm8 1h-1V4h2v2h-1zM7.5 5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5m3.5 5v1h2V9h-1v1zM2 9h1v1h1v1H2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyStackOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 12.5h15m-15 2h15M2.5 4V2.5H4m7 0h1.5V4m-10 3v1.5H4m7 0h1.5V7m-5 .5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6-7h12a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyStackSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v8A1.5 1.5 0 0 0 1.5 11h12A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0zM4 2H2v2h1V3h1zm3.5 1a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5M12 3h-1V2h2v2h-1zM3 7H2v2h2V8H3zm8 2V8h1V7h1v2z" clip-rule="evenodd"/><path fill="currentColor" d="M0 12v1h15v-1zm0 2v1h15v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MongodbOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.369-.338a.5.5 0 0 0-.738 0zm0 13l-.393.309a.5.5 0 0 0 .786 0zM4.623 9.838l-.393.31zm.246-6.467L4.5 3.032zm5.262 0l.369-.338zm.246 6.467l.393.31zM8 15V.5H7V15zm-.107-1.809L5.016 9.53l-.786.618l2.877 3.662zM5.237 3.708L7.87.838L7.13.162L4.5 3.032zM7.131.838l2.632 2.87l.737-.675L7.869.163zm2.853 8.691l-2.877 3.662l.786.618l2.877-3.662zm-.221-5.82a4.5 4.5 0 0 1 .22 5.82l.787.618a5.5 5.5 0 0 0-.27-7.114zm-4.747 5.82a4.5 4.5 0 0 1 .221-5.82L4.5 3.032a5.5 5.5 0 0 0-.27 7.114z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MongodbSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.869.162a.5.5 0 0 0-.738 0l-2.63 2.87a5.5 5.5 0 0 0-.271 7.115L7 13.673V15h1v-1.327l2.77-3.526a5.5 5.5 0 0 0-.27-7.114zM8 3a.5.5 0 0 0-1 0v7.5a.5.5 0 0 0 1 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodFlatOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 5.5h1m5 0h1m-7 4h7m-3.5 5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodFlatSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M4 6h1V5H4zm6 0h1V5h-1zm1 3v1H4V9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodFrownOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.29 10.187l-.424.266l.531.847l.424-.265zm4.898-.244l.498.035l.07-.998l-.5-.034zM5.82 11.035a7.278 7.278 0 0 1 4.367-1.092l.069-.997a8.278 8.278 0 0 0-4.967 1.241zM4 6h1V5H4zm6 0h1V5h-1zm-2.5 8A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodFrownSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M4 6h1V5H4zm1.82 5.035a7.278 7.278 0 0 1 4.368-1.092l.498.035l.07-.998l-.5-.034a8.278 8.278 0 0 0-4.966 1.241l-.424.266l.531.847zM11 6h-1V5h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodLaughOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 7.5V7a.5.5 0 0 0-.5.5zm6 0h.5a.5.5 0 0 0-.5-.5zm0-.5h-6v1h6zm-3 4C9.47 11 11 9.47 11 7.5h-1C10 8.918 8.918 10 7.5 10zM4 7.5C4 9.47 5.53 11 7.5 11v-1C6.082 10 5 8.918 5 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM4 6h1V5H4zm6 0h1V5h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodLaughSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 10c-1.246 0-2.233-.835-2.454-2h4.908c-.221 1.165-1.208 2-2.454 2"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M4 6h1V5H4zm.5 1a.5.5 0 0 0-.5.5C4 9.47 5.53 11 7.5 11S11 9.47 11 7.5a.5.5 0 0 0-.5-.5zM11 6h-1V5h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodSadOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.1 9.7l-.3.4l.8.6l.3-.4zm6 .6l.3.4l.8-.6l-.3-.4zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM4 6h1V5H4zm6 0h1V5h-1zm.9 3.7c-1.7-2.267-5.1-2.267-6.8 0l.8.6a3.25 3.25 0 0 1 5.2 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodSadSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M4 6h1V5H4zm6 0h1V5h-1zm-5.1 4.3a3.25 3.25 0 0 1 5.2 0l.8-.6c-1.7-2.267-5.1-2.267-6.8 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodSmileOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.9 8.7l-.3-.4l-.8.6l.3.4zm6 .6l.3-.4l-.8-.6l-.3.4zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM4 6h1V5H4zm6 0h1V5h-1zm.1 2.7a3.25 3.25 0 0 1-5.2 0l-.8.6c1.7 2.267 5.1 2.267 6.8 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodSmileSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M4 6h1V5H4zm6 0h1V5h-1zM4.9 8.7a3.25 3.25 0 0 0 5.2 0l.8.6c-1.7 2.267-5.1 2.267-6.8 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodSurprisedOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 5.5h1m5 0h1m-3.5 9a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm-.5-7h1a1.5 1.5 0 1 1 0 3H7a1.5 1.5 0 1 1 0-3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodSurprisedSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M4 6h1V5H4zm6 0h1V5h-1zM5 9a2 2 0 0 1 2-2h1a2 2 0 1 1 0 4H7a2 2 0 0 1-2-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodTongueOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 5.5h1m5 0h1m-7 3h7m-5.5 0v2a2 2 0 1 0 4 0v-2m-2 6a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodTongueSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 10.5V9h3v1.5a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M5 6H4V5h1zm6 0h-1V5h1zM4 9h1v1.5a2.5 2.5 0 0 0 5 0V9h1V8H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M1.66 11.362A6.5 6.5 0 0 0 7.693.502a7 7 0 1 1-6.031 10.86Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.707.003a.5.5 0 0 0-.375.846a6 6 0 0 1-5.569 10.024a.5.5 0 0 0-.519.765A7.5 7.5 0 1 0 7.707.003"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreHorizontalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3 7.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreHorizontalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 7.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m5 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVerticalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 3a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVerticalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 2.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m0 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m0 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 4v3m0 7.5a5 5 0 0 1-5-5v-4a5 5 0 0 1 10 0v4a5 5 0 0 1-5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 5.5a5.5 5.5 0 1 1 11 0v4a5.5 5.5 0 1 1-11 0zM7 4v3h1V4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MovOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-8 6l.354-.354A.5.5 0 0 0 2 6.5zm1 1l-.354.354l.354.353l.354-.353zm1-1H5a.5.5 0 0 0-.854-.354zm6 3H10v.207l.146.147zm1 1l-.354.354a.5.5 0 0 0 .708 0zm1-1l.354.354l.146-.147V9.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zM3 11V6.5H2V11zm-.854-4.146l1 1l.708-.708l-1-1zm1.708 1l1-1l-.708-.708l-1 1zM4 6.5V11h1V6.5zm4 1v2h1v-2zm-1 2v-2H6v2zm.5.5a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 11zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 9.5zM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6zm0-1A1.5 1.5 0 0 0 6 7.5h1a.5.5 0 0 1 .5-.5zM10 6v3.5h1V6zm.146 3.854l1 1l.708-.708l-1-1zm1.708 1l1-1l-.708-.708l-1 1zM13 9.5V6h-1v3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MovSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7.5a.5.5 0 0 1 1 0v2a.5.5 0 0 1-1 0z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM7.5 6A1.5 1.5 0 0 0 6 7.5v2a1.5 1.5 0 0 0 3 0v-2A1.5 1.5 0 0 0 7.5 6m-4.646.146A.5.5 0 0 0 2 6.5V11h1V7.707l.5.5l.5-.5V11h1V6.5a.5.5 0 0 0-.854-.354l-.646.647zM10 6h1v3.293l.5.5l.5-.5V6h1v3.707l-1.146 1.147a.5.5 0 0 1-.708 0L10 9.707z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MpFourOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-8 6l.354-.354A.5.5 0 0 0 2 6.5zm1 1l-.354.354l.354.353l.354-.353zm1-1H5a.5.5 0 0 0-.854-.354zm2 0V6H6v.5zm4 2H10V9h.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zM3 11V6.5H2V11zm-.854-4.146l1 1l.708-.708l-1-1zm1.708 1l1-1l-.708-.708l-1 1zM4 6.5V11h1V6.5zm2.5.5h1V6h-1zm.5 4V8.5H6V11zm0-2.5v-2H6v2zm.5-.5h-1v1h1zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 7.5zM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6zM10 6v2.5h1V6zm.5 3h2V8h-2zM12 6v5h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MpFourSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8h.5a.5.5 0 0 0 0-1H7z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM6 6h1.5a1.5 1.5 0 1 1 0 3H7v2H6zm-3.691.038a.5.5 0 0 1 .545.108l.646.647l.646-.647A.5.5 0 0 1 5 6.5V11H4V7.707l-.5.5l-.5-.5V11H2V6.5a.5.5 0 0 1 .309-.462M11 6h-1v3h2v2h1V6h-1v2h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MpThreeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-8 6l.354-.354A.5.5 0 0 0 2 6.5zm1 1l-.354.354l.354.353l.354-.353zm1-1H5a.5.5 0 0 0-.854-.354zm2 0V6H6v.5zm6 0l.4.3a.5.5 0 0 0-.4-.8zm-1.5 2l-.4-.3a.5.5 0 0 0 .4.8zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zM3 11V6.5H2V11zm-.854-4.146l1 1l.708-.708l-1-1zm1.708 1l1-1l-.708-.708l-1 1zM4 6.5V11h1V6.5zm2.5.5h1V6h-1zm.5 4V8.5H6V11zm0-2.5v-2H6v2zm.5-.5h-1v1h1zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 7.5zM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6zM10 7h2.5V6H10zm2.1-.8l-1.5 2l.8.6l1.5-2zM11 9h.5V8H11zm.5 1H10v1h1.5zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 13 9.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MpThreeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 8H7V7h.5a.5.5 0 0 1 0 1"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM7.5 6H6v5h1V9h.5a1.5 1.5 0 1 0 0-3m-4.646.146A.5.5 0 0 0 2 6.5V11h1V7.707l.5.5l.5-.5V11h1V6.5a.5.5 0 0 0-.854-.354l-.646.647zM11.5 7H10V6h2.5a.5.5 0 0 1 .4.8l-.951 1.268A1.5 1.5 0 0 1 11.5 11H10v-1h1.5a.5.5 0 0 0 0-1H11a.5.5 0 0 1-.4-.8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MsExcelOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 3.5v-2a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-2m0-6l4 4m-4 0l4-4m-5-2h6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MsExcelSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.793 7.5L2.146 5.854l.708-.708L4.5 6.793l1.646-1.647l.708.708L5.207 7.5l1.647 1.646l-.708.708L4.5 8.207L2.854 9.854l-.708-.708z"/><path fill="currentColor" fill-rule="evenodd" d="M3.5 0A1.5 1.5 0 0 0 2 1.5V3h-.5A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12H2v1.5A1.5 1.5 0 0 0 3.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0zm-2 4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MsPowerpointOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.755 3.5a7 7 0 1 1 0 8M2.5 10V8.5m0 0v-3H5a1.5 1.5 0 1 1 0 3zm-1-5h6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MsPowerpointSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 5H2v5h1V9h2a2 2 0 1 0 0-4m0 3H3V6h2a1 1 0 0 1 0 2"/><path d="M7.5 0a7.49 7.49 0 0 0-6 3A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12a7.5 7.5 0 1 0 6-12M1 4.5a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 .5.487v6.026a.5.5 0 0 1-.5.487h-6a.5.5 0 0 1-.5-.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MsWordOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 9.5l-.485.121a.5.5 0 0 0 .901.156zm1-1.5l.416-.277a.5.5 0 0 0-.832 0zm1 1.5l-.416.277a.5.5 0 0 0 .901-.156zM1.5 4h6V3h-6zm6.5.5v6h1v-6zM7.5 11h-6v1h6zM1 10.5v-6H0v6zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 12zm6.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 10.5zM7.5 4a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 3zm-6-1A1.5 1.5 0 0 0 0 4.5h1a.5.5 0 0 1 .5-.5zm.515 2.621l1 4l.97-.242l-1-4zm1.901 4.156l1-1.5l-.832-.554l-1 1.5zm.168-1.5l1 1.5l.832-.554l-1-1.5zm1.901 1.344l1-4l-.97-.242l-1 4zM3 3.5v-2H2v2zM3.5 1h10V0h-10zm10.5.5v12h1v-12zM13.5 14h-10v1h10zM3 13.5v-2H2v2zm.5.5a.5.5 0 0 1-.5-.5H2A1.5 1.5 0 0 0 3.5 15zm10.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0zM3 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 2 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MsWordSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.015 5.621l1 4a.5.5 0 0 0 .901.156l.584-.876l.584.876a.5.5 0 0 0 .901-.156l1-4l-.97-.242l-.726 2.903l-.373-.56a.5.5 0 0 0-.832 0l-.373.56l-.726-2.903z"/><path fill="currentColor" fill-rule="evenodd" d="M3.5 0A1.5 1.5 0 0 0 2 1.5V3h-.5A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12H2v1.5A1.5 1.5 0 0 0 3.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0zm-2 4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NsixtyFourOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 4v3M3 5.5h3m3 0h1m1 1h1m-7.5-4h-1a3 3 0 0 0-3 3v4.838a1.162 1.162 0 0 0 2.265.367L3.5 8.5h1.095a1 1 0 0 1 .995.9l.26 2.607a1.657 1.657 0 0 0 3.3 0L9.41 9.4a1 1 0 0 1 .995-.9H11.5l.735 2.205a1.162 1.162 0 0 0 2.265-.367V5.5a3 3 0 0 0-3-3h-1l-1-1h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NsixtyFourSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.293 1h4.414l1 1h.793A3.5 3.5 0 0 1 15 5.5v4.838a1.662 1.662 0 0 1-3.24.525L11.14 9h-.735a.5.5 0 0 0-.498.45l-.26 2.607a2.157 2.157 0 0 1-4.294 0l-.26-2.607A.5.5 0 0 0 4.595 9H3.86l-.62 1.863A1.662 1.662 0 0 1 0 10.338V5.5A3.5 3.5 0 0 1 3.5 2h.793zM4 7V6H3V5h1V4h1v1h1v1H5v1zm5-1h1V5H9zm3 0v1h-1V6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NesOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 7v3M2 8.5h3m6 1h1m-3 0h1m-8.5-6h12a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NesSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 4.5A1.5 1.5 0 0 1 1.5 3h12A1.5 1.5 0 0 1 15 4.5v6a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 10.5zM3 10V9H2V8h1V7h1v1h1v1H4v1zm8 0h1V9h-1zm-1 0H9V9h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetlifyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 7.5l-.354-.354a.5.5 0 0 0 0 .708zm7-7l.354-.354a.5.5 0 0 0-.708 0zm7 7l.354.354a.5.5 0 0 0 0-.708zm-7 7l-.354.354a.5.5 0 0 0 .708 0zM.854 7.854l7-7l-.708-.708l-7 7zm6.292-7l7 7l.708-.708l-7-7zm7 6.292l-7 7l.708.708l7-7zm-6.292 7l-7-7l-.708.708l7 7zM4.314 3.964l10 4l.372-.928l-10-4zM8.58 1.728l-5.5 8.5l.84.544l5.5-8.5zM2.1 5.8l6 8l.8-.6l-6-8zM.394 7.989l11.5 2.5l.212-.978l-11.5-2.5zm2.32 1.963l9.5-4.5l-.428-.904l-9.5 4.5zm7.292-6.53l-1.5 9.5l.988.156l1.5-9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetlifySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.146.146a.5.5 0 0 1 .708 0l.98.98l-1.82 2.94l-2.706-1.081zM3.539 3.754L2.426 4.867L4.72 7.772L6.479 4.93zM1.714 5.579L.146 7.146a.502.502 0 0 0-.061.075l3.522.755zM.66 8.367l1.007 1.007l1.175-.54zm1.76 1.761l.52.52l.654-1.058zm1.246 1.246l3.48 3.48a.5.5 0 0 0 .708 0l.323-.324l.119-.615L4.819 9.51zm5.772 1.895l2.488-2.488l-1.93-.413zm3.33-3.33l1.781-1.78l-3.833-1.534l-.531 2.76zM14.92 7.23a.51.51 0 0 0-.066-.084L12.99 5.283l-1.368.628zm-2.684-2.7l-.937-.938l-.288 1.499zm-1.791-1.792l-.885-.885l-1.604 2.59l2.007.804zm-4.447 5.75L9.662 6.81l-.455 2.367zm.114 1.047l2.906.623l-.473 2.46zm.075-2.233l2.9-1.329l-1.665-.666z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 5.5l.248-.434A.5.5 0 0 0 4 5.5zm0 4H4a.5.5 0 0 0 .748.434zm3.5-2l.248.434a.5.5 0 0 0 0-.868zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM4 5.5v4h1v-4zm.748 4.434l3.5-2l-.496-.868l-3.5 2zm3.5-2.868l-3.5-2l-.496.868l3.5 2zM10 5v5h1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M10 5h1v5h-1zm-5.252.066A.5.5 0 0 0 4 5.5v4a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.5 2.5l.29-.407A.5.5 0 0 0 1 2.5zm0 10H1a.5.5 0 0 0 .79.407zm7-5l.29.407a.5.5 0 0 0 0-.814zM1 2.5v10h1v-10zm.79 10.407l7-5l-.58-.814l-7 5zm7-5.814l-7-5l-.58.814l7 5zM13 2v11h1V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 5.5l.248-.434A.5.5 0 0 0 4 5.5zm0 4H4a.5.5 0 0 0 .748.434zm3.5-2l.248.434a.5.5 0 0 0 0-.868zm-4-2v4h1v-4zm.748 4.434l3.5-2l-.496-.868l-3.5 2zm3.5-2.868l-3.5-2l-.496.868l3.5 2zM10 5v5h1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.748 5.066A.5.5 0 0 0 4 5.5v4a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868zM10 10h1V5h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.79 2.093A.5.5 0 0 0 1 2.5v10a.5.5 0 0 0 .79.407l7-5a.5.5 0 0 0 0-.814zM13 13h1V2h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextjsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 4.5l.405-.293A.5.5 0 0 0 4 4.5zm3 9.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM5 12V4.5H4V12zm-.905-7.207l6.5 9l.81-.586l-6.5-9zM10 4v6h1V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextjsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 11.698 6.216L4.906 4.21A.5.5 0 0 0 4 4.5V12h1V6.06l5.83 8.162A7.5 7.5 0 0 1 0 7.5M10 10V4h1v6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NgcOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3 3.5a11.08 11.08 0 0 1 9 0M.5 6.5v5.764a1.236 1.236 0 0 0 2.342.553L3.5 11.5m11-5v5.764a1.236 1.236 0 0 1-2.342.553L11.5 11.5M6 7.5h3m-5.7 3.499L1.148 8.31A2.961 2.961 0 0 1 .5 6.461v-.396a2.565 2.565 0 0 1 5.032-.705l1.117 3.91a1.922 1.922 0 0 1-3.35 1.729Zm8.4 0l2.151-2.688a2.96 2.96 0 0 0 .649-1.85v-.396a2.565 2.565 0 0 0-5.032-.705L8.351 9.27a1.922 1.922 0 0 0 3.35 1.729Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NgcSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.117 3.005a11.58 11.58 0 0 0-9.234 0A3.065 3.065 0 0 0 0 6.065v6.199a1.736 1.736 0 0 0 3.289.776l.516-1.033A2.422 2.422 0 0 0 7.13 9.133L6.806 8h1.388l-.323 1.133a2.422 2.422 0 0 0 3.324 2.874l.516 1.033A1.736 1.736 0 0 0 15 12.264V6.065a3.065 3.065 0 0 0-2.883-3.06m-7.473.433c.65.39 1.15 1.018 1.368 1.785L6.52 7h1.96l.508-1.777a3.063 3.063 0 0 1 1.368-1.785a10.582 10.582 0 0 0-5.712 0M14 8.925l-1.909 2.386a2.416 2.416 0 0 1-.08.095l.595 1.187a.736.736 0 0 0 1.394-.33zm-11.012 2.48a2.516 2.516 0 0 1-.08-.094L1 8.925v3.339a.736.736 0 0 0 1.394.33z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NintendoSwitchOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 14.5H8a.5.5 0 0 0 .5.5zm0-12V2a.5.5 0 0 0-.5.5zm6 9h.5zm-3 3V14zm3-9H14zm-3-3V3zm-5-2H7a.5.5 0 0 0-.5-.5zm0 12v.5a.5.5 0 0 0 .5-.5zm-6-9H1zm3-3V1zm0 12v.5zm-3-3H0zm13.5-4v6h1v-6zM11.5 14h-3v1h3zm-2.5.5v-12H8v12zM8.5 3h3V2h-3zm5.5 8.5a2.5 2.5 0 0 1-2.5 2.5v1a3.5 3.5 0 0 0 3.5-3.5zm1-6A3.5 3.5 0 0 0 11.5 2v1A2.5 2.5 0 0 1 14 5.5zm-9-5v12h1V.5zM6.5 12h-3v1h3zM1 9.5v-6H0v6zM3.5 1h3V0h-3zM1 3.5A2.5 2.5 0 0 1 3.5 1V0A3.5 3.5 0 0 0 0 3.5zM3.5 12A2.5 2.5 0 0 1 1 9.5H0A3.5 3.5 0 0 0 3.5 13zm9.5-1.5A1.5 1.5 0 0 0 11.5 9v1a.5.5 0 0 1 .5.5zM11.5 12a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM10 10.5a1.5 1.5 0 0 0 1.5 1.5v-1a.5.5 0 0 1-.5-.5zm1 0a.5.5 0 0 1 .5-.5V9a1.5 1.5 0 0 0-1.5 1.5zM3.5 5a.5.5 0 0 1-.5-.5H2A1.5 1.5 0 0 0 3.5 6zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 4.5zM3.5 4a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 3zm0-1A1.5 1.5 0 0 0 2 4.5h1a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NintendoSwitchSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 4a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1"/><path fill="currentColor" fill-rule="evenodd" d="M7 .5v12a.5.5 0 0 1-.5.5h-3A3.5 3.5 0 0 1 0 9.5v-6A3.5 3.5 0 0 1 3.5 0h3a.5.5 0 0 1 .5.5m-5 4a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/><path fill="currentColor" d="M11.5 10a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1"/><path fill="currentColor" fill-rule="evenodd" d="M15 5.5v6a3.5 3.5 0 0 1-3.5 3.5h-3a.5.5 0 0 1-.5-.5v-12a.5.5 0 0 1 .5-.5h3A3.5 3.5 0 0 1 15 5.5m-5 5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodejsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.5 6v-.167c0-.736-.597-1.333-1.333-1.333H9a1.5 1.5 0 1 0 0 3h1a1.5 1.5 0 0 1 0 3H9A1.5 1.5 0 0 1 7.5 9m-2-5v5.264a2 2 0 0 1-1.106 1.789L3.5 11.5m-2-1v-6l6-3.5l6 3.5v6l-6 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodejsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 4.213L7.5.42L1 4.213v6.574l1.006.587l2.057-.832A1.5 1.5 0 0 0 5 9.152V4h1v5.152a2.5 2.5 0 0 1-1.562 2.317l-1.34.542L7.5 14.58l6.5-3.792zM7 6a2 2 0 0 1 2-2h1.167C11.179 4 12 4.82 12 5.833V6h-1v-.167A.833.833 0 0 0 10.167 5H9a1 1 0 0 0 0 2h1a2 2 0 1 1 0 4H9a2 2 0 0 1-2-2h1a1 1 0 0 0 1 1h1a1 1 0 1 0 0-2H9a2 2 0 0 1-2-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 14.5H10a.5.5 0 0 0 .854.354zm0-4V10a.5.5 0 0 0-.5.5zm4 0l.354.354A.5.5 0 0 0 14.5 10zM1.5 1h12V0h-12zM1 13.5v-12H0v12zm13-12v8.586h1V1.5zM10.086 14H1.5v1h8.586zm3.768-3.56l-3.415 3.414l.707.707l3.415-3.415zM10.086 15a1.5 1.5 0 0 0 1.06-.44l-.707-.706a.5.5 0 0 1-.353.146zM14 10.086a.5.5 0 0 1-.146.353l.707.707a1.5 1.5 0 0 0 .439-1.06zM0 13.5A1.5 1.5 0 0 0 1.5 15v-1a.5.5 0 0 1-.5-.5zM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5zM11 14.5v-4h-1v4zm-.5-3.5h4v-1h-4zm3.646-.854l-4 4l.708.708l4-4zM3 4h9V3H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5V9H9.5a.5.5 0 0 0-.5.5V15H1.5A1.5 1.5 0 0 1 0 13.5zM12 4H3V3h9z" clip-rule="evenodd"/><path fill="currentColor" d="M10 15h.086a1.5 1.5 0 0 0 1.06-.44l3.415-3.414a1.5 1.5 0 0 0 .439-1.06V10h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NpmOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 10.5v2h2v-2h8v-6H.5v6zm0 0v-6m4 0v6M6.5 6v3m-4-3v4.5m8-4.5v4.5m2-4.5v4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NpmSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 4h15v7H7v2H4v-2H0zm4 6V5H1v5h1V6h1v4zm1-5v7h1v-2h2V5zm4 0v5h1V6h1v4h1V6h1v4h1V5zM6 9V6h1v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NuxtjsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="m.5 12.5l6-10l6 10z"/><path d="m4.5 12.5l5-8.5l5 8.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NuxtjsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 2a.5.5 0 0 1 .429.243l1.527 2.545l.613-1.042a.5.5 0 0 1 .862 0l5 8.5A.5.5 0 0 1 14.5 13H.5a.5.5 0 0 1-.429-.757l6-10A.5.5 0 0 1 6.5 2M5.374 12h6.243L8.465 6.746zM7.88 5.77L4.214 12h-2.83L6.5 3.472zm1.163-.005L12.783 12h.843L9.5 4.986z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OmegaOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.333 11.687a.5.5 0 1 0 .334.943zm-4 .943a.5.5 0 1 0 .334-.943zM5.5 14.5v.5H6v-.5zm4 0H9v.5h.5zM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0zm0-1A6.5 6.5 0 0 0 1 6.5h1A5.5 5.5 0 0 1 7.5 1zM13 6.5a5.503 5.503 0 0 1-3.667 5.187l.334.943A6.503 6.503 0 0 0 14 6.5zm-7.333 5.187A5.503 5.503 0 0 1 2 6.5H1c0 2.83 1.81 5.238 4.333 6.13zM0 15h5.5v-1H0zm6-.5V12H5v2.5zm9-.5H9.5v1H15zm-5 .5V12H9v2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OmegaSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 1 1 9 6.002V14h5v1H9v-3h.026a.499.499 0 0 1 .307-.313a5.5 5.5 0 1 0-3.667 0a.496.496 0 0 1 .308.313H6v3H0v-1h5v-1.498A6.502 6.502 0 0 1 1 6.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OperaOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 12.5a3 3 0 0 1-3-3v-4a3 3 0 0 1 3-3m0 10a3 3 0 0 0 3-3v-4a3 3 0 0 0-3-3m0 10h.585c1.907 0 3.78-.518 5.415-1.5m-6-8.5h.585c1.907 0 3.78.518 5.415 1.5m-6 10.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OperaSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7.5a7.5 7.5 0 0 1 13.147-4.936A17.5 17.5 0 0 0 8.74 2H7.5A3.5 3.5 0 0 0 4 5.5v4A3.5 3.5 0 0 0 7.5 13h1.241c1.488 0 2.969-.19 4.406-.563A7.5 7.5 0 0 1 0 7.5"/><path fill="currentColor" d="M14.073 11.115A7.466 7.466 0 0 0 15 7.5c0-1.31-.336-2.543-.927-3.615l-.114-.038a16.5 16.5 0 0 0-3.962-.8A3.489 3.489 0 0 1 11 5.5v4a3.49 3.49 0 0 1-1.003 2.452a16.499 16.499 0 0 0 3.962-.799z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OtpOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 5.5h3m-1.5 0V10m3 0V7.5m0 0v-2h1a1 1 0 1 1 0 2zm-6-1v2a1 1 0 0 1-2 0v-2a1 1 0 0 1 2 0Zm-3-6h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OtpSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 6a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 1 0v-2a.5.5 0 0 0-.5-.5M11 7h.5a.5.5 0 0 0 0-1H11z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5zm2 5a1.5 1.5 0 1 1 3 0v2a1.5 1.5 0 1 1-3 0zM7 6H6V5h3v1H8v4H7zm3-1h1.5a1.5 1.5 0 0 1 0 3H11v2h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageBreakOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 8.993H0v1h.5zm2 1H3v-1h-.5zm2-1H4v1h.5zm2 1H7v-1h-.5zm2-1H8v1h.5zm2 1h.5v-1h-.5zm2-1H12v1h.5zm2 1h.5v-1h-.5zM10.5.5l.354-.354L10.707 0H10.5zm3 3h.5v-.207l-.146-.147zM.5 9.993h2v-1h-2zm4 0h2v-1h-2zm4 0h2v-1h-2zm4 0h2v-1h-2zm0 4.007h-10v1h10zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM13 3.5V8h1V3.5zm0 8.25v1.75h1v-1.75zM2 8V1.5H1V8zm0 5.5V11H1v2.5zm.5.5a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2.5 0A1.5 1.5 0 0 0 1 1.5h1a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageBreakSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V8H1zM1 11h13v2.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM0 8.993h3v1H0zm4 0h3v1H4zm7 0H8v1h3zm1 0h3v1h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageNumberOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.5.5l.354-.354L10.707 0H10.5zm3 3h.5v-.207l-.146-.147zm-4 9l-.354-.354A.5.5 0 0 0 9.5 13zm3 1.5h-10v1h10zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2.5 0A1.5 1.5 0 0 0 1 1.5h1a.5.5 0 0 1 .5-.5zM12 12H9.5v1H12zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792zM10.793 9H10.5v1h.293zM10.5 9A1.5 1.5 0 0 0 9 10.5h1a.5.5 0 0 1 .5-.5zm1.5 1.207C12 9.54 11.46 9 10.793 9v1c.114 0 .207.093.207.207zm-.354.854c.227-.227.354-.534.354-.854h-1a.207.207 0 0 1-.06.147zM13 3.5v10h1v-10zm-11 10v-12H1v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageNumberSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM10.5 9A1.5 1.5 0 0 0 9 10.5h1a.5.5 0 0 1 .5-.5h.293a.207.207 0 0 1 .146.354l-1.793 1.792A.5.5 0 0 0 9.5 13H12v-1h-1.293l.94-.94A1.207 1.207 0 0 0 10.793 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintbrushOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 14l-.481-.136a.5.5 0 0 0 .758.552zm.971-3.422l.481.136zm-.46 3.081l-.277-.416zm2.062-.148l-.215.451zM5 6l-.25-.433a.5.5 0 0 0-.104.787zm4 4l-.354.354a.5.5 0 0 0 .787-.103zM14.5.5l.433.25a.5.5 0 0 0-.684-.683zM.981 14.137l.971-3.423l-.962-.273l-.971 3.422zm2.956-4.923H4v-1h-.063zm2.063 2v.053h1v-.053zm-1.947 2h-.08v1h.08zm-3.32.03l-.51.34l.554.832l.512-.34zm2.555-.185a2.594 2.594 0 0 0-2.554.184l.555.832a1.594 1.594 0 0 1 1.569-.113zm.685.155c-.237 0-.471-.053-.685-.155l-.43.903c.348.166.73.252 1.115.252zM6 11.267a1.947 1.947 0 0 1-1.947 1.947v1A2.947 2.947 0 0 0 7 11.267zM4 9.214a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3zm-2.048 1.5a2.063 2.063 0 0 1 1.985-1.5v-1c-1.37 0-2.573.91-2.947 2.227zm2.694-4.36l4 4l.708-.708l-4-4zM14.25.067l-9.5 5.5l.5.866l9.5-5.5zM9.433 10.251l5.5-9.5l-.866-.502l-5.5 9.5zM7.146 4.854l3 3l.708-.708l-3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintbrushSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.854.146a.5.5 0 0 1 .079.605l-3.841 6.634l-3.477-3.477L14.25.068a.5.5 0 0 1 .605.078M6.72 4.427l-1.97 1.14a.5.5 0 0 0-.104.787l4 4a.5.5 0 0 0 .787-.103l1.14-1.97zM.99 10.441a3.063 3.063 0 0 1 2.947-2.227H4a3 3 0 0 1 3 3v.053a2.947 2.947 0 0 1-2.947 2.947h-.08a2.59 2.59 0 0 1-1.115-.252a1.594 1.594 0 0 0-1.57.113l-.51.341a.5.5 0 0 1-.759-.553z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintbucketOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.086 2.914l-.354-.353zM1.914 7.086l-.353-.354zm0 2.828l.354-.353zm3.172 3.172l.353-.354zm2.828 0l.354.353zm4.172-4.172l-.354-.353zm0-2.828l.353-.354zM8.914 2.914l-.353.354zm3.767 7.586l.353-.354l-.353-.353l-.354.353zm1.18 1.18l-.353.354zM11.639 14l.312-.39zm2.086 0l-.312-.39zM5.732 2.56L1.561 6.733l.707.707l4.171-4.171zm-4.171 7.708l3.171 3.171l.707-.707l-3.171-3.171zm6.707 3.171l4.171-4.171l-.707-.707l-4.171 4.171zm4.171-7.707L9.268 2.561l-.707.707l3.171 3.171zm0 3.536a2.5 2.5 0 0 0 0-3.536l-.707.707a1.5 1.5 0 0 1 0 2.122zm-7.707 4.171a2.5 2.5 0 0 0 3.536 0l-.707-.707a1.5 1.5 0 0 1-2.122 0zM1.561 6.732a2.5 2.5 0 0 0 0 3.536l.707-.707a1.5 1.5 0 0 1 0-2.122zm4.878-3.464a1.5 1.5 0 0 1 2.122 0l.707-.707a2.5 2.5 0 0 0-3.536 0zM7 7V2.5H6V7zM2 2.5v4h1v-4zM4.5 0A2.5 2.5 0 0 0 2 2.5h1A1.5 1.5 0 0 1 4.5 1zM7 2.5A2.5 2.5 0 0 0 4.5 0v1A1.5 1.5 0 0 1 6 2.5zM1.5 10h10V9h-10zm10.827.854l1.181 1.18l.707-.707l-1.18-1.18zm-.473 1.18l1.18-1.18l-.707-.708l-1.18 1.181zm.096 1.576c-.29-.232-.422-.51-.437-.77c-.016-.257.08-.545.34-.806l-.707-.707c-.443.444-.666 1.004-.632 1.575c.035.569.324 1.099.811 1.488zm1.462 0a1.17 1.17 0 0 1-1.462 0l-.625.78a2.17 2.17 0 0 0 2.711 0zm.096-1.576c.26.26.357.549.341.807c-.016.259-.147.537-.437.769l.624.78c.487-.39.777-.919.811-1.489c.035-.57-.188-1.13-.632-1.574z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintbucketSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 0A2.5 2.5 0 0 0 2 2.5v3.793l-.44.44a2.5 2.5 0 0 0 0 3.535l3.172 3.171a2.5 2.5 0 0 0 3.536 0l4.171-4.171a2.5 2.5 0 0 0 0-3.536L9.268 2.561a2.498 2.498 0 0 0-2.342-.666A2.501 2.501 0 0 0 4.5 0M6 3.707V7h1V2.914a1.5 1.5 0 0 1 1.56.354l3.172 3.171a1.5 1.5 0 0 1 0 2.122l-.44.439H1.915a1.5 1.5 0 0 1 .354-1.56zm-.009-1.372A1.5 1.5 0 0 0 3 2.5v2.793L5.732 2.56c.082-.083.169-.158.259-.226" clip-rule="evenodd"/><path fill="currentColor" d="m12.645 9.737l1.534 1.534a2.17 2.17 0 1 1-3.069 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13 1.5H6.5a4 4 0 1 0 0 8h1m3 4.5V1.5M7.5 14V1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 5.5A4.5 4.5 0 0 1 6.5 1H13v1h-2v12h-1V2H8v12H7v-4h-.5A4.5 4.5 0 0 1 2 5.5M7 9V2h-.5a3.5 3.5 0 1 0 0 7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PasswordOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 8.5v-1a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-1m0-4h-4a2 2 0 1 0 0 4h4m0-4a2 2 0 1 1 0 4m-9-6v-3a3 3 0 0 1 6 0v3m2.5 4h1m-3 0h1m-3 0h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PasswordSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 11h-1v-1h1zm-3 0h1v-1H8zm5 0h-1v-1h1z"/><path fill="currentColor" fill-rule="evenodd" d="M3 6V3.5a3.5 3.5 0 1 1 7 0V6h1.5A1.5 1.5 0 0 1 13 7.5v.55a2.5 2.5 0 0 1 0 4.9v.55a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 0 13.5v-6A1.5 1.5 0 0 1 1.5 6zm1-2.5a2.5 2.5 0 0 1 5 0V6H4zM8.5 9a1.5 1.5 0 1 0 0 3h4a1.5 1.5 0 0 0 0-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PatreonOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5.5h-2v14h2zm2 5a5 5 0 1 0 10 0a5 5 0 0 0-10 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PatreonSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0H0v15h3zm6.5 0a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 5v5m2-5v5m-1 4.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M7 10H6V5h1zm2 0H8V5h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 3v9m4-9v9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 5v5m2-5v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 10V5h1v5zm2 0V5h1v5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 12V3h1v9zm4 0V3h1v9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PawOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 3V2a1.5 1.5 0 1 0-3 0v1a1.5 1.5 0 1 0 3 0Zm5 0V2a1.5 1.5 0 0 0-3 0v1a1.5 1.5 0 1 0 3 0Zm3 4.5V7a1.5 1.5 0 0 0-3 0v.5a1.5 1.5 0 0 0 3 0Zm-11 0V7a1.5 1.5 0 1 0-3 0v.5a1.5 1.5 0 1 0 3 0Zm-.645 4.14l2.918-3.543a2.237 2.237 0 0 1 3.454 0l2.918 3.543c.939 1.14.128 2.86-1.35 2.86c-.194 0-.385-.045-.559-.132l-.36-.18a5.315 5.315 0 0 0-4.753 0l-.36.18a1.248 1.248 0 0 1-.558.132c-1.478 0-2.289-1.72-1.35-2.86Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PawSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 0a2 2 0 0 0-2 2v1a2 2 0 1 0 4 0V2a2 2 0 0 0-2-2m5 0a2 2 0 0 0-2 2v1a2 2 0 1 0 4 0V2a2 2 0 0 0-2-2M2 5a2 2 0 0 0-2 2v.5a2 2 0 1 0 4 0V7a2 2 0 0 0-2-2m11 0a2 2 0 0 0-2 2v.5a2 2 0 1 0 4 0V7a2 2 0 0 0-2-2M9.613 7.779a2.737 2.737 0 0 0-4.226 0L2.47 11.322C1.261 12.789 2.305 15 4.205 15c.272 0 .54-.063.782-.185l.36-.18a4.814 4.814 0 0 1 4.306 0l.36.18c.242.122.51.185.782.185c1.9 0 2.944-2.211 1.736-3.678z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PawsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-miterlimit="10" d="m1.425 2.118l.307.647M.75 4.647l.307.647m5.277-3.176l.307.706M3.88 1l.306.647m6.934 5.118l-.306.647m-2.148.47l-.307.706m5.891 1.824l-.307.647m-.368-3.177l-.307.647M3.941 4.235c.43-.176.92-.176 1.35.06l1.657.823c.552.294.736 1 .307 1.47c-.185.236-.491.353-.798.353H5.72c-.306 0-.613.118-.797.353l-.43.53c-.184.235-.49.352-.798.352c-.613 0-1.104-.588-.981-1.176l.368-1.765c.123-.411.43-.823.859-1ZM11.059 10c.43.177.737.588.86 1.059l.367 1.765A.976.976 0 0 1 11.305 14c-.307 0-.614-.118-.798-.353l-.43-.53c-.184-.235-.49-.352-.797-.352h-.737c-.307 0-.613-.118-.798-.353a.978.978 0 0 1 .307-1.47l1.657-.824a1.25 1.25 0 0 1 1.35-.118Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PawsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.852 1.885L4.117.334l-.903.428l.735 1.551zM2.398 3.002l-.735-1.55l-.904.428l.735 1.55zm4.901.081L6.593 1.46l-.917.398l.706 1.623zm-3.549.69a2.057 2.057 0 0 1 1.772.077l1.66.826c.786.419 1.124 1.484.454 2.236c-.3.373-.764.53-1.18.53H5.72c-.192 0-.336.074-.403.16l-.006.007l-.427.527c-.301.381-.769.54-1.189.54c-.917 0-1.661-.866-1.47-1.778l.367-1.765a.5.5 0 0 1 .01-.04c.157-.524.55-1.074 1.149-1.32M1.723 5.532L.988 3.98l-.904.428l.735 1.551zm9.328 2.546l.735-1.55l-.903-.43l-.735 1.551zm2.455 1.117l.735-1.55l-.904-.429l-.735 1.551zm-4.888.051l.706-1.623l-.917-.399L7.7 8.847zm.825.445a1.751 1.751 0 0 1 1.832-.143c.591.254.976.806 1.127 1.385a.4.4 0 0 1 .006.024l.368 1.764c.186.89-.47 1.779-1.471 1.779c-.42 0-.888-.16-1.189-.54l-.433-.534c-.067-.086-.21-.161-.403-.161h-.737c-.42 0-.89-.16-1.191-.545a1.479 1.479 0 0 1 .465-2.22l.013-.007zm4.737 2.034l.736-1.55l-.904-.43l-.735 1.551z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaypalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M.5 12.5h4l1-4h1.795a4.625 4.625 0 0 0 4.33-3.001C12.532 3.08 10.745.5 8.161.5H3.5z"/><path d="M4 14.5h4L9 11h1.577c1.477 0 2.82-.859 3.438-2.2c.927-2.008-.54-4.3-2.75-4.3H6.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaypalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.015.379A.5.5 0 0 1 3.5 0h4.661c2.397 0 4.191 1.957 4.204 4.172c1.928.626 3.021 2.851 2.105 4.837a4.287 4.287 0 0 1-3.893 2.491h-1.2l-.896 3.137A.5.5 0 0 1 8 15H4a.5.5 0 0 1-.485-.621L3.86 13H.5a.5.5 0 0 1-.485-.621zM8.16 1c1.762 0 3.097 1.388 3.197 3.001L11.264 4H6.5a.5.5 0 0 0-.485.379L4.11 12H1.14L3.89 1zm-.866 8H5.89l-1.25 5h2.983l.896-3.137A.5.5 0 0 1 9 10.5h1.577a3.287 3.287 0 0 0 2.985-1.91c.626-1.357-.057-2.87-1.32-3.396c-.039.16-.089.32-.149.48A5.125 5.125 0 0 1 7.296 9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PdfOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 6.5V6H2v.5zm4 0V6H6v.5zm0 4H6v.5h.5zm7-7h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zM2.5 7h1V6h-1zm.5 4V8.5H2V11zm0-2.5v-2H2v2zm.5-.5h-1v1h1zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 7.5zM3.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 6zM6 6.5v4h1v-4zm.5 4.5h1v-1h-1zM9 9.5v-2H8v2zM7.5 6h-1v1h1zM9 7.5A1.5 1.5 0 0 0 7.5 6v1a.5.5 0 0 1 .5.5zM7.5 11A1.5 1.5 0 0 0 9 9.5H8a.5.5 0 0 1-.5.5zM10 6v5h1V6zm.5 1H13V6h-2.5zm0 2H12V8h-1.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PdfSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 8H3V7h.5a.5.5 0 0 1 0 1M7 10V7h.5a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM3.5 6H2v5h1V9h.5a1.5 1.5 0 1 0 0-3m4 0H6v5h1.5A1.5 1.5 0 0 0 9 9.5v-2A1.5 1.5 0 0 0 7.5 6m2.5 5V6h3v1h-2v1h1v1h-1v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5.5V0H2v.5zm10 0h.5V0h-.5zM4.947 4.724a.5.5 0 0 0-.894-.448zM2.5 8.494l-.447-.223l-.146.293l.21.251zm5 5.997l-.384.32a.5.5 0 0 0 .769 0zm5-5.996l.384.32l.21-.251l-.146-.293zm-1.553-4.219a.5.5 0 0 0-.894.448zM8 9.494v-.5H7v.5zm-.5-4.497A4.498 4.498 0 0 1 3 .5H2a5.498 5.498 0 0 0 5.5 5.497zM2.5 1h10V0h-10zM12 .5a4.498 4.498 0 0 1-4.5 4.497v1c3.038 0 5.5-2.46 5.5-5.497zM4.053 4.276l-2 3.995l.895.448l2-3.995zM2.116 8.815l5 5.996l.769-.64l-5-5.996zm5.768 5.996l5-5.996l-.768-.64l-5 5.996zm5.064-6.54l-2-3.995l-.895.448l2 3.995zM8 14.49V9.494H7v4.997z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0h11v.5a5.482 5.482 0 0 1-1.874 4.134l1.968 3.93L8 14.672V8.994H7v5.678L1.907 8.564l1.967-3.93A5.482 5.482 0 0 1 2 .5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.72.5H2.5a2 2 0 0 0-2 2v2c0 5.523 4.477 10 10 10h2a2 2 0 0 0 2-2v-1.382a1 1 0 0 0-.553-.894l-2.416-1.208a1 1 0 0 0-1.396.578l-.297.893a1.21 1.21 0 0 1-1.385.804A6.047 6.047 0 0 1 3.71 6.547a1.21 1.21 0 0 1 .804-1.385l1.108-.37a1 1 0 0 0 .654-1.19L5.69 1.257A1 1 0 0 0 4.72.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 0 2.5v2C0 10.299 4.701 15 10.5 15h2a2.5 2.5 0 0 0 2.5-2.5v-1.382a1.5 1.5 0 0 0-.83-1.342l-2.415-1.208a1.5 1.5 0 0 0-2.094.868l-.298.893a.71.71 0 0 1-.812.471A5.547 5.547 0 0 1 4.2 6.45a.71.71 0 0 1 .471-.812l1.109-.37a1.5 1.5 0 0 0 .98-1.787l-.586-2.344A1.5 1.5 0 0 0 4.72 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonecallBlockedOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m13.5 1.5l-4 4m-7-5h2.22a1 1 0 0 1 .97.757l.585 2.345a1 1 0 0 1-.654 1.19l-1.108.37a1.21 1.21 0 0 0-.804 1.385a6.047 6.047 0 0 0 4.744 4.744a1.21 1.21 0 0 0 1.385-.804l.297-.893a1 1 0 0 1 1.396-.578l2.416 1.208a1 1 0 0 1 .553.894V12.5a2 2 0 0 1-2 2h-2c-5.523 0-10-4.477-10-10v-2a2 2 0 0 1 2-2Zm9 6a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonecallBlockedSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2.5A2.5 2.5 0 0 1 2.5 0h2.22a1.5 1.5 0 0 1 1.454 1.136L6.76 3.48a1.5 1.5 0 0 1-.98 1.787l-1.109.37a.71.71 0 0 0-.471.812A5.547 5.547 0 0 0 8.55 10.8a.71.71 0 0 0 .812-.471l.298-.893a1.5 1.5 0 0 1 2.094-.868l2.416 1.208a1.5 1.5 0 0 1 .83 1.342V12.5a2.5 2.5 0 0 1-2.5 2.5h-2C4.701 15 0 10.299 0 4.5z"/><path fill="currentColor" fill-rule="evenodd" d="M8 3.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0M11.5 1a2.5 2.5 0 0 0-2.086 3.879l3.465-3.465A2.488 2.488 0 0 0 11.5 1m0 5c-.51 0-.983-.152-1.379-.414l3.465-3.465A2.5 2.5 0 0 1 11.5 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonecallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m14.5.5l-6 6m6-6V4m0-3.5H11M2.5.5h2.22a1 1 0 0 1 .97.757l.585 2.345a1 1 0 0 1-.654 1.19l-1.108.37a1.21 1.21 0 0 0-.804 1.385a6.047 6.047 0 0 0 4.744 4.744a1.21 1.21 0 0 0 1.385-.804l.297-.893a1 1 0 0 1 1.396-.578l2.416 1.208a1 1 0 0 1 .553.894V12.5a2 2 0 0 1-2 2h-2c-5.523 0-10-4.477-10-10v-2a2 2 0 0 1 2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonecallReceiveOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m14.5.5l-6 6m0 0V3m0 3.5H12M2.5.5h2.22a1 1 0 0 1 .97.757l.585 2.345a1 1 0 0 1-.654 1.19l-1.108.37a1.21 1.21 0 0 0-.804 1.385a6.047 6.047 0 0 0 4.744 4.744a1.21 1.21 0 0 0 1.385-.804l.297-.893a1 1 0 0 1 1.396-.578l2.416 1.208a1 1 0 0 1 .553.894V12.5a2 2 0 0 1-2 2h-2c-5.523 0-10-4.477-10-10v-2a2 2 0 0 1 2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonecallReceiveSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 0 2.5v2C0 10.299 4.701 15 10.5 15h2a2.5 2.5 0 0 0 2.5-2.5v-1.382a1.5 1.5 0 0 0-.83-1.342l-2.415-1.208a1.5 1.5 0 0 0-2.094.868l-.298.893a.71.71 0 0 1-.812.471A5.547 5.547 0 0 1 4.2 6.45a.71.71 0 0 1 .471-.812l1.109-.37a1.5 1.5 0 0 0 .98-1.787l-.586-2.344A1.5 1.5 0 0 0 4.72 0z"/><path fill="currentColor" d="M14.146.146L9 5.293V3H8v4h4V6H9.707L14.854.854z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonecallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 0h-4v1h2.293L8.146 6.146l.708.708L14 1.707V4h1z"/><path fill="currentColor" d="M2.5 0A2.5 2.5 0 0 0 0 2.5v2C0 10.299 4.701 15 10.5 15h2a2.5 2.5 0 0 0 2.5-2.5v-1.382a1.5 1.5 0 0 0-.83-1.342l-2.415-1.208a1.5 1.5 0 0 0-2.094.868l-.298.893a.71.71 0 0 1-.812.471A5.547 5.547 0 0 1 4.2 6.45a.71.71 0 0 1 .471-.812l1.109-.37a1.5 1.5 0 0 0 .98-1.787l-.586-2.344A1.5 1.5 0 0 0 4.72 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5.5V0a.5.5 0 0 0-.5.5zm8 8V9a.5.5 0 0 0 .5-.5zm-8 0H6V9h.5zm0 6.5A6.5 6.5 0 0 0 13 8.5h-1A5.5 5.5 0 0 1 6.5 14zM0 8.5A6.5 6.5 0 0 0 6.5 15v-1A5.5 5.5 0 0 1 1 8.5zm1 0A5.5 5.5 0 0 1 6.5 3V2A6.5 6.5 0 0 0 0 8.5zM6.5 1A7.5 7.5 0 0 1 14 8.5h1A8.5 8.5 0 0 0 6.5 0zM6 .5v8h1v-8zM6.5 9h8V8h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 0H6v9h9v-.5A8.5 8.5 0 0 0 6.5 0"/><path fill="currentColor" d="M12.826 10H5V2.174A6.5 6.5 0 1 0 12.826 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 7.5H7a.5.5 0 0 0 .146.354zm0 6.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM7 0v7.5h1V0zm.724 7.947l6-3l-.448-.894l-6 3zm-.578-.093l5 5l.708-.708l-5-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 .016a7.5 7.5 0 1 0 5.438 13.13L7.15 7.857A.5.5 0 0 1 7 7.5z"/><path fill="currentColor" d="M13.145 12.438A7.47 7.47 0 0 0 15 7.5a7.476 7.476 0 0 0-.581-2.9L8.344 7.637zm.825-8.732A7.499 7.499 0 0 0 8 .016v6.675z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 15V8.5m0 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0A4.5 4.5 0 0 0 7 8.973V15h1V8.973A4.5 4.5 0 0 0 7.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="square" clip-rule="evenodd"><path d="M7.5 8.495a2 2 0 0 0 2-1.999a2 2 0 0 0-4 0a2 2 0 0 0 2 1.999Z"/><path d="M13.5 6.496c0 4.997-5 7.995-6 7.995s-6-2.998-6-7.995A5.999 5.999 0 0 1 7.5.5c3.313 0 6 2.685 6 5.996Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 6.496a1.5 1.5 0 0 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M1 6.496A6.499 6.499 0 0 1 7.5 0C11.089 0 14 2.909 14 6.496c0 2.674-1.338 4.793-2.772 6.225a10.865 10.865 0 0 1-2.115 1.654c-.322.19-.623.34-.885.442c-.247.098-.506.174-.728.174c-.222 0-.481-.076-.728-.174a6.453 6.453 0 0 1-.885-.442a10.868 10.868 0 0 1-2.115-1.654C2.338 11.289 1 9.17 1 6.496m6.5-2.499a2.5 2.5 0 0 0-2.5 2.5a2.5 2.5 0 0 0 5 0a2.5 2.5 0 0 0-2.5-2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m4.5 13.5l3-7m-3.236 3a2.989 2.989 0 0 1-.764-2V7A3.5 3.5 0 0 1 7 3.5h1A3.5 3.5 0 0 1 11.5 7v.5a3 3 0 0 1-3 3a2.081 2.081 0 0 1-1.974-1.423L6.5 9m1 5.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7.5a7.5 7.5 0 1 1 4.584 6.912l1.911-4.367A2.58 2.58 0 0 0 8.5 11A3.5 3.5 0 0 0 12 7.5V7a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v.5c0 .896.337 1.715.891 2.333l.745-.666A2.489 2.489 0 0 1 4 7.5V7a3 3 0 0 1 3-3h1a3 3 0 0 1 3 3v.5A2.5 2.5 0 0 1 8.5 10A1.58 1.58 0 0 1 7 8.919l-.005-.016l.963-2.203l-.916-.4l-3.352 7.66A7.496 7.496 0 0 1 0 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlantOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 15V7m0 .5v3m0-3a4 4 0 0 0-4-4h-3v3a4 4 0 0 0 4 4h3m0-3h3a4 4 0 0 0 4-4v-3h-3a4 4 0 0 0-4 4zm0 0l4-4m-4 7l-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlantSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 4.5A4.5 4.5 0 0 1 11.5 0H15v3.5A4.5 4.5 0 0 1 10.5 8H8v7H7v-4H4.5A4.5 4.5 0 0 1 0 6.5V3h3.5c1.414 0 2.675.652 3.5 1.671zm1.146 1.646l3-3l.708.708l-3 3zm-2 3.708l-3-3l.708-.708l3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.5 5.5l.248-.434A.5.5 0 0 0 6 5.5zm0 4H6a.5.5 0 0 0 .748.434zm3.5-2l.248.434a.5.5 0 0 0 0-.868zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM6 5.5v4h1v-4zm.748 4.434l3.5-2l-.496-.868l-3.5 2zm3.5-2.868l-3.5-2l-.496.868l3.5 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m6.249-2.432a.5.5 0 0 1 .5-.002l3.5 2a.5.5 0 0 1 0 .868l-3.5 2A.5.5 0 0 1 6 9.5v-4a.5.5 0 0 1 .249-.432" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M4.5 12.5v-10l7 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M6.5 9.5v-4l3.5 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.748 5.066A.5.5 0 0 0 6 5.5v4a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.79 2.093A.5.5 0 0 0 4 2.5v10a.5.5 0 0 0 .79.407l7-5a.5.5 0 0 0 0-.814z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlugOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 11.5V15m2-3.5V15m-4-15v4.5m6-4.5v4.5m-8 0h10v3h-1v2a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-2h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlugSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4H5V0H4v4H2v4h1v1.5A2.5 2.5 0 0 0 5.5 12H6v3h1v-3h1v3h1v-3h.5A2.5 2.5 0 0 0 12 9.5V8h1V4h-2V0h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 4v7M4 7.5h7m-3.5 7a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M7 11V8H4V7h3V4h1v3h3v1H8v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PngOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 6.5V6H2v.5zm8 4H10v.5h.5zm2 0v.5h.5v-.5zm1-7h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-4 6l.447-.224L6 6.5zm-.5 4v.5h1v-.5zm2.5 0l-.447.224A.5.5 0 0 0 9 10.5zm.5-4V6H8v.5zM2.5 7h1V6h-1zm.5 4V8.5H2V11zm0-2.5v-2H2v2zm.5-.5h-1v1h1zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 7.5zM3.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 6zM10 6v4.5h1V6zm.5 5h2v-1h-2zm2.5-.5v-2h-1v2zM10.5 7H13V6h-2.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zm5-7v4h1v-4zm.053.224l2 4l.894-.448l-2-4zM8 6.5v4h1v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PngSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8h.5a.5.5 0 0 0 0-1H3z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM2 6h1.5a1.5 1.5 0 1 1 0 3H3v2H2zm8 0h3v1h-2v3h1V8.5h1V11h-3zM7 8.618V11H6V6h1v.382l1 2V6h1v5H8v-.382z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoolOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="10" d="M1 12.6c.65.733 1.655 1.4 2.955 1.4c2.954 0 3.545-2 6.5-2c1.359 0 2.6.733 3.545 1.467M2.5 12V3.727C2.5 1.945 3.855.5 5.636.5M9.5 10V3.636C9.5 1.855 10.9.5 12.682.5M2.5 4.5h7m-7 4h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoolSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.636 1C4.15 1 3 2.203 3 3.727V4h6v-.364C9 1.57 10.633 0 12.682 0v1C11.167 1 10 2.14 10 3.636V10H9V9H3v3H2V3.727C2 1.688 3.56 0 5.636 0zM3 8h6V5H3z" clip-rule="evenodd"/><path fill="currentColor" d="M7.44 13.442c-.895.504-1.877 1.058-3.485 1.058c-1.483 0-2.614-.762-3.33-1.568l.75-.664c.584.66 1.462 1.232 2.58 1.232c1.339 0 2.128-.442 3.004-.935l.01-.007c.895-.504 1.877-1.058 3.485-1.058c1.531 0 2.884.82 3.852 1.572l-.612.79c-.923-.716-2.052-1.362-3.24-1.362c-1.339 0-2.128.442-3.004.935z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoundOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5 14.5l1.103-.367A2.775 2.775 0 0 0 3.5 11.5V4.442A3.942 3.942 0 0 1 7.442.5h.865C9.934.5 11.396 1.49 12 3M3 13h1.084a6 6 0 0 1 2.683.633l.05.025a6 6 0 0 0 5.366 0L13.5 13M1 7.5h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoundSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.442A4.442 4.442 0 0 1 7.442 0h.865a4.477 4.477 0 0 1 4.157 2.814l-.928.372A3.477 3.477 0 0 0 8.307 1h-.865A3.44 3.44 0 0 0 4 4.442V7h5v1H4v3.5c0 .346-.054.683-.156 1h.24a6.5 6.5 0 0 1 2.906.686l.05.025l-.223.447l.223-.447a5.5 5.5 0 0 0 4.92 0l1.316-.658l.448.894l-1.317.659a6.5 6.5 0 0 1-5.814 0l-.05-.025l.224-.448l-.224.448a5.5 5.5 0 0 0-2.46-.581h-.765a3.27 3.27 0 0 1-1.557 1.107l-1.103.367l-.316-.948l1.103-.368A2.275 2.275 0 0 0 3 11.5V8H1V7h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5 8.5v-8M2.618 2.499A6.963 6.963 0 0 0 .5 7.495c0 3.864 3.135 7.005 7 7.005c3.867 0 7-3.141 7-7.005A6.968 6.968 0 0 0 12.395 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 0v9H7V0zm4.387 1.792l.358.35A7.468 7.468 0 0 1 15 7.494C15 11.635 11.644 15 7.5 15C3.358 15 0 11.635 0 7.495a7.46 7.46 0 0 1 2.269-5.354l.357-.35l.7.716l-.359.35A6.463 6.463 0 0 0 1 7.494A6.506 6.506 0 0 0 7.5 14c3.59 0 6.5-2.917 6.5-6.505a6.464 6.464 0 0 0-1.955-4.639l-.357-.35z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PptOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 6.5V6H2v.5zm4 0V6H6v.5zm7-3h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zM2.5 7h1V6h-1zm.5 4V8.5H2V11zm0-2.5v-2H2v2zm.5-.5h-1v1h1zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 7.5zM3.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 6zm3 0h1V6h-1zm.5 4V8.5H6V11zm0-2.5v-2H6v2zm.5-.5h-1v1h1zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 7.5zM7.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 6zM11 6v5h1V6zm-1 1h3V6h-3zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PptSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8h.5a.5.5 0 0 0 0-1H3zm4 0h.5a.5.5 0 0 0 0-1H7z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM2 6h1.5a1.5 1.5 0 1 1 0 3H3v2H2zm4 0h1.5a1.5 1.5 0 1 1 0 3H7v2H6zm5 5h1V7h1V6h-3v1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrintOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 12.5h-2a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-2m-8-6v-5a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v5m-8 4h8v4h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrintSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1.5A1.5 1.5 0 0 1 4.5 0h6A1.5 1.5 0 0 1 12 1.5V5H3zM1.5 6A1.5 1.5 0 0 0 0 7.5v4A1.5 1.5 0 0 0 1.5 13H2V9h11v4h.5a1.5 1.5 0 0 0 1.5-1.5v-4A1.5 1.5 0 0 0 13.5 6z"/><path fill="currentColor" d="M3 10h9v5H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PythonOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 2.5h1M4.5 4V1.5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-4a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V11M8 4.5H1.5a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h3m2.5-1h6.5a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1h-3m-2.5 9h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PythonSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 12H1.5A1.5 1.5 0 0 1 0 10.5v-5A1.5 1.5 0 0 1 1.5 4H8V3H7V2H6v1H4V1.5A1.5 1.5 0 0 1 5.5 0h4A1.5 1.5 0 0 1 11 1.5v5a.5.5 0 0 1-.5.5h-6A1.5 1.5 0 0 0 3 8.5z"/><path fill="currentColor" d="M12 3v3.5A1.5 1.5 0 0 1 10.5 8h-6a.5.5 0 0 0-.5.5v5A1.5 1.5 0 0 0 5.5 15h4a1.5 1.5 0 0 0 1.5-1.5V12H9v1H8v-1H7v-1h6.5A1.5 1.5 0 0 0 15 9.5v-5A1.5 1.5 0 0 0 13.5 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCodeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 8.5H8.5V12M14 8.5h1m-3 6H8m3-3h3.5V15M3 3.5h1m7 0h1m-9 8h1M1.5.5h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Zm8 0h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Zm-8 8h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCodeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h4A1.5 1.5 0 0 1 7 1.5v4A1.5 1.5 0 0 1 5.5 7h-4A1.5 1.5 0 0 1 0 5.5zM1.5 1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5zm6.5.5A1.5 1.5 0 0 1 9.5 0h4A1.5 1.5 0 0 1 15 1.5v4A1.5 1.5 0 0 1 13.5 7h-4A1.5 1.5 0 0 1 8 5.5zM9.5 1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5zM4 4H3V3h1zm8 0h-1V3h1zM0 9.5A1.5 1.5 0 0 1 1.5 8h4A1.5 1.5 0 0 1 7 9.5v4A1.5 1.5 0 0 1 5.5 15h-4A1.5 1.5 0 0 1 0 13.5zM1.5 9a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5zM8 8h4v1H9v3H8zm7 1h-1V8h1zM4 12H3v-1h1zm10 0h-3v-1h4v4h-1zm-6 2h4v1H8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 9V7.5H8A1.5 1.5 0 0 0 9.5 6v-.1a1.4 1.4 0 0 0-1.4-1.4h-.6A1.5 1.5 0 0 0 6 6m1 4.5h1m-.5 4a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M5.5 6a2 2 0 0 1 2-2h.6c1.05 0 1.9.85 1.9 1.9V6a2 2 0 0 1-2 2v1H7V7h1a1 1 0 0 0 1-1v-.1a.9.9 0 0 0-.9-.9h-.6a1 1 0 0 0-1 1zM7 11v-1h1v1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 12v-1.264c0-1.37.774-2.623 2-3.236a3.65 3.65 0 0 0 2-3.257C11.5 2.195 9.84.5 7.792.5h-.207c-1.335 0-2.615.53-3.56 1.474L3.5 2.5m3.5 12h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 9V7.5H8A1.5 1.5 0 0 0 9.5 6v-.1a1.4 1.4 0 0 0-1.4-1.4h-.6A1.5 1.5 0 0 0 6 6m1 4.5h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 6a2 2 0 0 1 2-2h.6c1.05 0 1.9.85 1.9 1.9V6a2 2 0 0 1-2 2v1H7V7h1a1 1 0 0 0 1-1v-.1a.9.9 0 0 0-.9-.9h-.6a1 1 0 0 0-1 1zM8 10v1H7v-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.672 1.62A5.534 5.534 0 0 1 7.585 0h.207C10.122 0 12 1.925 12 4.243a4.15 4.15 0 0 1-2.276 3.704A3.118 3.118 0 0 0 8 10.737V12H7v-1.264c0-1.56.881-2.986 2.276-3.683A3.15 3.15 0 0 0 11 4.243C11 2.465 9.558 1 7.792 1h-.207a4.534 4.534 0 0 0-3.206 1.328l-.525.526l-.708-.708zM8 15H7v-1h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 6.5h4a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1zm0 0V10A3.5 3.5 0 0 0 5 13.5m3.5-7h4a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1zm0 0V10a3.5 3.5 0 0 0 3.5 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 1A1.5 1.5 0 0 0 1 2.5V10a4 4 0 0 0 4 4v-1a3 3 0 0 1-3-3V7h3.5A1.5 1.5 0 0 0 7 5.5v-3A1.5 1.5 0 0 0 5.5 1zm7 0A1.5 1.5 0 0 0 8 2.5V10a4 4 0 0 0 4 4v-1a3 3 0 0 1-3-3V7h3.5A1.5 1.5 0 0 0 14 5.5v-3A1.5 1.5 0 0 0 12.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RandOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 14V8.5m0 0v-7H8a3.5 3.5 0 1 1 0 7zm0 0h5a3 3 0 0 1 3 3V14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RandSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1h5a4 4 0 0 1 2.117 7.395A3.5 3.5 0 0 1 12 11.5V14h-1v-2.5A2.5 2.5 0 0 0 8.5 9H4v5H3zm1 7h4a3 3 0 0 0 0-6H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReactOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M14.5 7.584c0 1.657-3.134 3-7 3s-7-1.343-7-3s3.134-3 7-3s7 1.343 7 3Z"/><path d="M4.166 13.739c1.457.79 4.13-1.327 5.972-4.726c1.841-3.4 2.153-6.795.696-7.584c-1.457-.79-4.13 1.327-5.972 4.726c-1.841 3.4-2.153 6.795-.696 7.584Z"/><path d="M10.834 13.739c-1.457.79-4.13-1.327-5.972-4.726c-1.841-3.4-2.153-6.795-.696-7.584c1.457-.79 4.13 1.327 5.972 4.726c1.841 3.4 2.153 6.795.696 7.584Z"/><path d="M6.5 7.584a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReactSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.315 1.837c-.4-.116-.695-.085-.91.032c-.216.116-.404.347-.526.745c-.122.401-.163.936-.104 1.582c.01.105.022.212.037.321a13.654 13.654 0 0 1 1.676-.311a13.28 13.28 0 0 1 1.275-1.54a7.111 7.111 0 0 0-.066-.053c-.508-.402-.98-.66-1.382-.776m2.185.14c-.06-.05-.121-.1-.182-.148c-.572-.452-1.158-.789-1.724-.953C5.024.711 4.441.711 3.928.99c-.513.278-.833.767-1.005 1.334c-.172.564-.21 1.238-.144 1.965c.016.17.037.344.065.523c-.17.06-.334.125-.49.192c-.671.287-1.246.642-1.66 1.062C.278 6.487 0 7 0 7.584c0 .583.278 1.097.694 1.519c.414.42.989.774 1.66 1.062c.156.067.32.131.49.192a8.672 8.672 0 0 0-.065.523c-.066.726-.028 1.4.144 1.965c.172.567.492 1.056 1.005 1.333c.513.278 1.097.279 1.666.114c.566-.165 1.152-.5 1.724-.953l.182-.149c.06.051.121.1.182.149c.572.452 1.158.788 1.724.953c.569.165 1.153.164 1.666-.114c.513-.277.833-.766 1.005-1.333c.172-.564.21-1.239.144-1.965a8.616 8.616 0 0 0-.065-.523a8.2 8.2 0 0 0 .49-.192c.671-.288 1.246-.643 1.66-1.062c.416-.422.694-.936.694-1.52c0-.582-.278-1.096-.694-1.518c-.414-.42-.989-.775-1.66-1.062a8.706 8.706 0 0 0-.49-.192c.027-.179.049-.353.065-.523c.066-.727.028-1.4-.144-1.965c-.172-.567-.492-1.056-1.005-1.334C10.56.711 9.975.711 9.406.876c-.566.164-1.152.5-1.724.953zm0 1.365c-.225.23-.45.483-.672.755a16.99 16.99 0 0 1 1.344 0a11.385 11.385 0 0 0-.672-.755m2.012.864c-.41-.574-.84-1.092-1.275-1.54l.065-.053c.51-.402.98-.66 1.383-.776c.399-.116.695-.085.91.032c.216.116.404.347.525.745c.122.401.164.936.105 1.582c-.01.105-.022.212-.037.32a13.655 13.655 0 0 0-1.676-.31m-.563.944a15.628 15.628 0 0 0-2.898 0A15.627 15.627 0 0 0 4.72 7.584a15.693 15.693 0 0 0 1.33 2.433a15.634 15.634 0 0 0 2.9 0a15.63 15.63 0 0 0 1.33-2.433A15.691 15.691 0 0 0 8.95 5.15m1.824 1.138a17.244 17.244 0 0 0-.527-.956c.26.05.511.106.752.168c-.063.256-.138.52-.225.788m0 2.591a16.837 16.837 0 0 1-.527.957c.26-.05.511-.106.752-.169a11.69 11.69 0 0 0-.225-.788m1.18.487a13.805 13.805 0 0 0-.588-1.782c.246-.61.443-1.209.588-1.782c.103.038.203.079.3.12c.596.256 1.047.547 1.341.845c.292.296.406.572.406.817s-.114.52-.406.816c-.294.299-.745.59-1.341.846a7.467 7.467 0 0 1-.3.12m-.765 1.285a13.57 13.57 0 0 1-1.676.311c-.41.574-.84 1.091-1.275 1.54l.066.052c.508.403.98.66 1.382.777c.399.116.695.085.91-.032c.216-.117.404-.348.525-.746c.123-.4.164-.936.105-1.582a7.422 7.422 0 0 0-.037-.32M7.5 11.826c.225-.23.45-.483.672-.755a16.945 16.945 0 0 1-1.344 0c.222.272.447.524.672.755m-2.746-1.99a16.972 16.972 0 0 1-.527-.957c-.087.27-.162.532-.225.788c.24.063.492.119.752.169m-.942.815a13.57 13.57 0 0 0 1.676.311c.41.574.839 1.091 1.275 1.54a6.761 6.761 0 0 1-.066.052c-.508.403-.98.66-1.382.777c-.4.116-.695.085-.911-.032c-.216-.117-.403-.348-.525-.746c-.122-.4-.163-.936-.104-1.582a7.47 7.47 0 0 1 .037-.32m-.765-1.285c.145-.574.341-1.172.588-1.782a13.806 13.806 0 0 1-.588-1.782a8.518 8.518 0 0 0-.3.12c-.596.256-1.047.547-1.341.845c-.292.296-.406.572-.406.817s.114.52.406.816c.294.299.745.59 1.341.846c.097.041.197.081.3.12m.955-3.865c.063.255.138.519.225.787a17.116 17.116 0 0 1 .527-.956c-.26.05-.511.106-.752.169M6 7.584a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m1.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5.5V0a.5.5 0 0 0-.5.5zm12 0h.5a.5.5 0 0 0-.5-.5zm0 14l-.224.447A.5.5 0 0 0 14 14.5zm-2-1l.224-.447a.5.5 0 0 0-.448 0zm-2 1l-.224.447a.5.5 0 0 0 .448 0zm-2-1l.224-.447a.5.5 0 0 0-.448 0zm-2 1l-.224.447a.5.5 0 0 0 .448 0zm-4 0H1a.5.5 0 0 0 .724.447zm2-1l.224-.447a.5.5 0 0 0-.448 0zM1.5 1h12V0h-12zM13 .5v14h1V.5zm.724 13.553l-2-1l-.448.894l2 1zm-2.448-1l-2 1l.448.894l2-1zm-1.552 1l-2-1l-.448.894l2 1zm-2.448-1l-2 1l.448.894l2-1zM2 14.5V.5H1v14zm3.724-.447l-2-1l-.448.894l2 1zm-2.448-1l-2 1l.448.894l2-1zM4 5h2V4H4zm4 0h3V4H8zM4 8h2V7H4zm4 0h3V7H8zm0 3h3v-1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 .5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-.724.447L11.5 14.06l-1.776.888a.5.5 0 0 1-.448 0L7.5 14.06l-1.776.888a.5.5 0 0 1-.448 0L3.5 14.06l-1.776.888A.5.5 0 0 1 1 14.5zM4 5h2V4H4zm4 0h3V4H8zM6 8H4V7h2zm2 0h3V7H8zm3 3H8v-1h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 1.5l.121-.485A.5.5 0 0 0 7 1.5zm5.5 8c0 .774-.55 1.641-1.583 2.343C10.4 12.533 8.998 13 7.5 13v1c1.696 0 3.294-.525 4.479-1.33C13.148 11.876 14 10.743 14 9.5zM7.5 13c-1.498 0-2.9-.466-3.917-1.157C2.551 11.14 2 10.273 2 9.5H1c0 1.243.852 2.376 2.021 3.17C4.206 13.475 5.804 14 7.5 14zM2 9.5c0-.774.55-1.641 1.583-2.343C4.6 6.467 6.002 6 7.5 6V5c-1.696 0-3.294.525-4.479 1.33C1.852 7.124 1 8.257 1 9.5zM7.5 6c1.498 0 2.9.467 3.917 1.157C12.449 7.86 13 8.727 13 9.5h1c0-1.243-.852-2.376-2.021-3.17C10.794 5.525 9.196 5 7.5 5zm2.306 4.54c-.69.29-1.32.46-2.306.46v1c1.136 0 1.898-.204 2.694-.54zM7.5 11c-.987 0-1.617-.17-2.306-.46l-.388.92c.796.336 1.558.54 2.694.54zM8 5.5v-4H7v4zm-.621-3.515l4 1l.242-.97l-4-1zM3.974 6.841c-.286-.855-1.12-1.297-1.952-1.297v1c.51 0 .886.261 1.004.615zM2.022 5.544A2.022 2.022 0 0 0 0 7.566h1a1.02 1.02 0 0 1 1.022-1.022zM0 7.566C0 8.589.76 9.424 1.74 9.56l.139-.99A1.016 1.016 0 0 1 1 7.565zm11.974-.407c.118-.354.493-.615 1.004-.615v-1c-.832 0-1.666.442-1.952 1.297zm1.004-.615A1.02 1.02 0 0 1 14 7.566h1a2.022 2.022 0 0 0-2.022-2.022zM14 7.566c0 .511-.38.934-.879 1.004l.139.99A2.016 2.016 0 0 0 15 7.567zM12.5 3a.5.5 0 0 1-.5-.5h-1A1.5 1.5 0 0 0 12.5 4zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 14 2.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 12.5 1zm0-1A1.5 1.5 0 0 0 11 2.5h1a.5.5 0 0 1 .5-.5zM5 9h1V8H5zm4 0h1V8H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.621 1.015A.5.5 0 0 0 7 1.5v3.515A8.16 8.16 0 0 0 3.455 6.06c-.388-.341-.911-.515-1.433-.515A2.022 2.022 0 0 0 0 7.566c0 .747.406 1.394 1.007 1.742A2.725 2.725 0 0 0 1 9.5c0 1.243.852 2.376 2.021 3.17C4.206 13.475 5.804 14 7.5 14c1.696 0 3.294-.525 4.479-1.33C13.148 11.876 14 10.743 14 9.5c0-.064-.002-.128-.007-.192A2.008 2.008 0 0 0 15 7.566a2.022 2.022 0 0 0-2.022-2.022c-.522 0-1.045.174-1.433.515A8.16 8.16 0 0 0 8 5.015V2.14l3.055.764a1.5 1.5 0 1 0 .074-1.012zM12.5 3a.5.5 0 0 1-.5-.492v-.016a.5.5 0 1 1 .5.508M5 9h1V8H5zm2.5 2c-.987 0-1.617-.17-2.306-.46l-.388.92c.796.336 1.558.54 2.694.54s1.898-.204 2.694-.54l-.388-.92c-.69.29-1.32.46-2.306.46M10 9H9V8h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedwoodjsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5.5l.248-.434a.5.5 0 0 0-.496 0zM4 2.5l-.248-.434a.5.5 0 0 0-.142.122zM2 5l-.39-.312a.5.5 0 0 0-.09.175zM1 8.5l-.48-.137a.5.5 0 0 0 .033.36zm2 4l-.447.224a.5.5 0 0 0 .244.233zm4.5 2l-.203.457a.5.5 0 0 0 .406 0zm4.5-2l.203.457a.5.5 0 0 0 .244-.233zm2-4l.447.224a.5.5 0 0 0 .034-.361zM13 5l.48-.137a.5.5 0 0 0-.09-.175zm-2-2.5l.39-.312a.499.499 0 0 0-.142-.122zM7.252.066l-3.5 2l.496.868l3.5-2zM3.61 2.188l-2 2.5l.78.624l2-2.5zM1.52 4.863l-1 3.5l.96.274l1-3.5zm-.967 3.86l2 4l.894-.447l-2-4zm2.244 4.234l4.5 2l.406-.914l-4.5-2zm4.906 2l4.5-2l-.406-.914l-4.5 2zm4.744-2.233l2-4l-.894-.448l-2 4zm2.034-4.361l-1-3.5l-.962.274l1 3.5zm-1.09-3.675l-2-2.5l-.781.624l2 2.5zm-2.143-2.622l-3.5-2l-.496.868l3.5 2zm-8.011.86l10.5 6.5l.526-.851l-10.5-6.5zm8-.851l-10.5 6.5l.526.85l10.5-6.5zm-9.59 3.279l5.5 5.5l.707-.708l-5.5-5.5zm5.63 5.593l4 2l.447-.894l-4-2zm5.37-6.3l-5.5 5.5l.707.707l5.5-5.5zm-5.37 5.406l-4 2l.447.894l4-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedwoodjsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.252.066a.5.5 0 0 1 .496 0l3.36 1.92L7.5 4.337L3.893 1.985zM3.185 2.718L1.96 4.25L4.2 6.49l2.386-1.556zM1.438 5.146L.52 8.363a.5.5 0 0 0 .034.36l.054.11l2.735-1.784zm-.38 4.587l1.371 2.743l4.227-2.113l-2.591-2.591zm2.219 3.437l4.02 1.787a.5.5 0 0 0 .406 0l4.02-1.787L7.5 11.06zm9.293-.694l1.372-2.743l-3.007-1.961l-2.59 2.591zm1.823-3.643l.054-.11a.5.5 0 0 0 .034-.36l-.92-3.217l-1.903 1.903zm-1.351-4.582l-1.227-1.533l-3.4 2.217l2.387 1.556zM7.5 5.532l2.58 1.682L7.5 9.793l-2.58-2.58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 7.5A7 7 0 0 1 13 3.17m1.5 4.33A7 7 0 0 1 2 11.83m3-.33H1.5V15m12-15v3.5H10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 1A6.5 6.5 0 0 0 1 7.5H0a7.5 7.5 0 0 1 13-5.1V0h1v4h-4V3h2.19A6.48 6.48 0 0 0 7.5 1m0 13A6.5 6.5 0 0 0 14 7.5h1a7.5 7.5 0 0 1-13 5.099V15H1v-4h4v1H2.81a6.48 6.48 0 0 0 4.69 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 14.5A7 7 0 0 1 3.17 2M7.5.5A7 7 0 0 1 11.83 13m-.33-3v3.5H15M0 1.5h3.5V5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 7.5A6.5 6.5 0 0 0 7.5 1V0a7.5 7.5 0 0 1 5.099 13H15v1h-4v-4h1v2.19a6.48 6.48 0 0 0 2-4.69M2.4 2H0V1h4v4H3V2.81A6.5 6.5 0 0 0 7.5 14v1A7.5 7.5 0 0 1 2.4 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 5.5H7a.5.5 0 0 0-.748-.434zm0 4l-.248.434A.5.5 0 0 0 7 9.5zM3 7.5l-.248-.434a.5.5 0 0 0 0 .868zm7.5-2h.5a.5.5 0 0 0-.748-.434zm0 4l-.248.434A.5.5 0 0 0 11 9.5zM7 7.5l-.248-.434a.5.5 0 0 0 0 .868zm.5 7.5A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14zM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5zM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zm0 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zM6 5.5v4h1v-4zm.748 3.566l-3.5-2l-.496.868l3.5 2zm-3.5-1.132l3.5-2l-.496-.868l-3.5 2zM10 5.5v4h1v-4zm.748 3.566l-3.5-2l-.496.868l3.5 2zm-3.5-1.132l3.5-2l-.496-.868l-3.5 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 0-15 0a7.5 7.5 0 0 0 15 0m-4.249-2.432a.5.5 0 0 0-.5-.002L7 6.924V5.5a.5.5 0 0 0-.748-.434l-3.5 2a.5.5 0 0 0 0 .868l3.5 2A.5.5 0 0 0 7 9.5V8.076l3.252 1.858A.5.5 0 0 0 11 9.5v-4a.5.5 0 0 0-.249-.432" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M14.5 12.5v-10l-7 5zm-7 0v-10l-7 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M6.5 9.5v-4L3 7.5zm4 0v-4L7 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5.5a.5.5 0 0 0-.748-.434l-3.5 2a.5.5 0 0 0 0 .868l3.5 2A.5.5 0 0 0 7 9.5V8.076l3.252 1.858A.5.5 0 0 0 11 9.5v-4a.5.5 0 0 0-.748-.434L7 6.924z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2.5a.5.5 0 0 0-.79-.407l-7 5a.5.5 0 0 0 0 .814l7 5A.5.5 0 0 0 8 12.5V8.472l6.21 4.435A.5.5 0 0 0 15 12.5v-10a.5.5 0 0 0-.79-.407L8 6.528z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.146 10.146l-.353.354l.707.707l.354-.353zM9.5 7.5l.354.354l.353-.354l-.353-.354zM6.854 4.146L6.5 3.793l-.707.707l.353.354zm0 6.708l3-3l-.708-.708l-3 3zm3-3.708l-3-3l-.708.708l3 3zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zm0 1A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0M6 3.293L10.207 7.5L6 11.707z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m5 14l7-6.5L5 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m6.5 10.5l3-3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.207 7.5L6 3.293v8.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7.5L4 0v15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RippleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m1.5 1.5l3.06 3.316a4 4 0 0 0 5.88 0L13.5 1.5m-12 12l3.06-3.316a4 4 0 0 1 5.88 0L13.5 13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RippleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.193 5.155l-3.06-3.316l.734-.678l3.061 3.316a3.5 3.5 0 0 0 5.144 0l3.06-3.316l.735.678l-3.06 3.316a4.5 4.5 0 0 1-6.614 0m5.879 5.368a3.5 3.5 0 0 0-5.144 0l-3.06 3.316l-.735-.678l3.06-3.316a4.5 4.5 0 0 1 6.614 0l3.06 3.316l-.734.678z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 2.5a5 5 0 0 1 5 5v6a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1v-6a5 5 0 0 1 5-5Zm0 0V0M4 11.5h7M.5 8v4m14-4v4m-9-2.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm4 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0m4 0a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0"/><path fill="currentColor" fill-rule="evenodd" d="M8 2.022A5.5 5.5 0 0 1 13 7.5v6a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 13.5v-6a5.5 5.5 0 0 1 5-5.478V0h1zM5.5 7a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m4 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m1.5 5H4v-1h7z" clip-rule="evenodd"/><path fill="currentColor" d="M0 8v4h1V8zm15 0h-1v4h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RollerOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.05 7.914l-.138-.48zM8.224 9.293l.138.48zM5.5 11.5V11a.5.5 0 0 0-.5.5zm4 0h.5a.5.5 0 0 0-.5-.5zm0 3v.5a.5.5 0 0 0 .5-.5zm-4 0H5a.5.5 0 0 0 .5.5zM1.5 1h10V0h-10zm10.5.5v2h1v-2zM11.5 4h-10v1h10zM1 3.5v-2H0v2zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 5zM12 3.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 13 3.5zM11.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 0zm-10-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5zM14 4.5v1.491h1V4.5zm-1.088 2.934L8.088 8.812l.275.962l4.824-1.379zM7 10.254V11.5h1v-1.246zm1.088-1.442A1.5 1.5 0 0 0 7 10.254h1a.5.5 0 0 1 .363-.48zM14 5.992a1.5 1.5 0 0 1-1.088 1.442l.275.961A2.5 2.5 0 0 0 15 5.991zM12.5 3A1.5 1.5 0 0 1 14 4.5h1A2.5 2.5 0 0 0 12.5 2zm-7 9h4v-1h-4zm3.5-.5v3h1v-3zm.5 2.5h-4v1h4zm-3.5.5v-3H5v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RollerSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h10A1.5 1.5 0 0 1 13 1.5v.55a2.5 2.5 0 0 1 2 2.45v1.491a2.5 2.5 0 0 1-1.813 2.404L8.363 9.774a.5.5 0 0 0-.363.48V11h1.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5H7v-.746a1.5 1.5 0 0 1 1.088-1.442l4.824-1.378A1.5 1.5 0 0 0 14 5.99V4.5a1.5 1.5 0 0 0-1-1.415V3.5A1.5 1.5 0 0 1 11.5 5h-10A1.5 1.5 0 0 1 0 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RollupjsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5.5V0a.5.5 0 0 0-.5.5zm12 14v.5a.5.5 0 0 0 .455-.707zm-2.488-5.474l-.224-.447a.5.5 0 0 0-.231.654zM7.36 2.78l.447.224zM4.765 14.058l-.428-.258zM8.772 3.317l.475-.158zm.457 1.371l-.474.158zm2.235-.295l-.474.158zm-2.487-2.61l.417.278zm1.362 3.885l.098-.49zM1 .5v12h1V.5zM3.5 15h10v-1h-10zm10.455-.707L11.467 8.82l-.91.414l2.488 5.474zM1.5 1H9V0H1.5zM13 5a4 4 0 0 1-2.212 3.58l.448.893A5 5 0 0 0 14 5zM9 1a4 4 0 0 1 4 4h1a5 5 0 0 0-5-5zM1 12.5A2.5 2.5 0 0 0 3.5 15v-1A1.5 1.5 0 0 1 2 12.5zm1.447 1.224l5.36-10.72l-.894-.447l-5.36 10.72zm5.36-10.72a2.838 2.838 0 0 1 2.624-1.568l.03-1a3.838 3.838 0 0 0-3.548 2.12zM4.929 14.757l.265-.442l-.857-.515l-.266.443zm.265-.442a58.922 58.922 0 0 1 6.182-8.486l-.752-.658a59.924 59.924 0 0 0-6.287 8.63zm3.104-10.84l.457 1.371l.949-.316l-.457-1.371zm3.64.76l-.464-1.393l-.948.316l.464 1.393zM10.443.5a2.26 2.26 0 0 0-1.88 1.006l.832.555a1.26 1.26 0 0 1 1.048-.561zm-.201 5.658a1.483 1.483 0 0 0 1.698-1.923l-.949.316c.117.352-.19.7-.553.627zM8.755 4.846a1.974 1.974 0 0 0 1.486 1.312l.196-.98a.974.974 0 0 1-.733-.648zm.492-1.687a1.26 1.26 0 0 1 .147-1.098l-.833-.555a2.26 2.26 0 0 0-.263 1.969zm1.184-1.723a.438.438 0 0 1-.25-.09a.475.475 0 0 1-.186-.41c.026-.385.382-.436.447-.436v1c.132 0 .523-.082.551-.497a.525.525 0 0 0-.207-.454a.562.562 0 0 0-.325-.112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RollupjsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 .5a.5.5 0 0 1 .5-.5H9c.782 0 1.523.18 2.182.5h-.135a3.838 3.838 0 0 0-4.134 2.057L1.325 13.733A2.489 2.489 0 0 1 1 12.5z"/><path fill="currentColor" d="M4.87 15h8.63a.5.5 0 0 0 .455-.707l-2.298-5.057A4.997 4.997 0 0 0 14 5a4.984 4.984 0 0 0-1.43-3.5h-2.128a1.26 1.26 0 0 0-1.195 1.659l.457 1.371c.11.332.39.579.733.648a.485.485 0 0 0 .178.002l.009-.01l.007.007a.484.484 0 0 0 .359-.626l-.464-1.393l.948-.316l.465 1.393a1.483 1.483 0 0 1-.736 1.793a55.058 55.058 0 0 0-5.95 8.315z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RouterOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 9.5V9H1v.5zm12 0h.5V9h-.5zm0 5v.5h.5v-.5zm-12 0H1v.5h.5zm1.72-8.339C4.373 4.761 5.916 4 7.5 4V3c-1.917 0-3.732.924-5.052 2.525zM7.5 4c1.583 0 3.126.762 4.281 2.161l.771-.636C11.232 3.924 9.417 3 7.5 3zm-6.614.318C2.658 2.17 5.04 1 7.5 1V0C4.71 0 2.055 1.33.114 3.682zM7.5 1c2.46 0 4.842 1.17 6.614 3.318l.772-.636C12.945 1.329 10.29 0 7.5 0zM7 6v3h1V6zm-5.5 4h12V9h-12zM13 9.5v5h1v-5zm.5 4.5h-12v1h12zM2 14.5v-5H1v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RouterSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 1C5.04 1 2.658 2.17.886 4.318l-.772-.636C2.055 1.329 4.71 0 7.5 0s5.445 1.33 7.386 3.682l-.772.636C12.342 2.17 9.96 1 7.5 1m0 3c-1.583 0-3.126.762-4.28 2.161l-.772-.636C3.768 3.924 5.583 3 7.5 3c1.917 0 3.732.924 5.052 2.525l-.771.636C10.626 4.761 9.083 4 7.5 4M7 9V6h1v3h6v6H1V9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 13.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm14 1.5C14.5 6.992 8.008.5 0 .5m0 6A8.5 8.5 0 0 1 8.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 15C14 7.268 7.732 1 0 1V0c8.284 0 15 6.716 15 15z"/><path fill="currentColor" fill-rule="evenodd" d="M0 13.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/><path fill="currentColor" d="M9 15a9 9 0 0 0-9-9v1a8 8 0 0 1 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RubyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 14.5v.5h.5v-.5zm0-14h.5V0h-.5zm-6 0V0h-.207l-.147.146zm-8 8l-.354-.354L0 8.293V8.5zm0 6H0v.5h.5zm4-4l-.224.447l.019.01l.02.007zm0-6V4a.5.5 0 0 0-.5.5zm6 0l.464-.186l-.008-.019l-.009-.019zm4.5 10V.5h-1v14zM14.5 0h-6v1h6zM8.146.146l-8 8l.708.708l8-8zM0 8.5v6h1v-6zM.5 15h14v-1H.5zM14.146.146l-14 14l.708.708l14-14zM5 10.5v-6H4v6zM4.5 5h6V4h-6zm-.186 5.964l10 4l.372-.928l-10-4zm5.722-6.278l4 10l.928-.372l-4-10zM8.053.724l2 4l.894-.448l-2-4zM.276 8.947l4 2l.448-.894l-4-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RubySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.293 4L8.058.236L9.73 4zm10-4l-3.632 3.632L9.047 0zM.236 8.058L4 9.73V4.293zm3.396 2.603L0 9.047v5.246zM5 9.293L9.293 5H5zm10 4.438l-3.907-9.117L15 .707zm-.952.317l-3.717-8.672l-4.955 4.955zm-9.434-2.955L13.731 15H.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RupeeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 8.5V8a.5.5 0 0 0-.325.88zM2 1h11V0H2zm.5 8h3V8h-3zm3-9h-3v1h3zM2.175 8.88l7 6l.65-.76l-7-6zM10 4.5A4.5 4.5 0 0 0 5.5 0v1A3.5 3.5 0 0 1 9 4.5zM5.5 9A4.5 4.5 0 0 0 10 4.5H9A3.5 3.5 0 0 1 5.5 8zM2 5h11V4H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RupeeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1V0h11v1H8.329a4.494 4.494 0 0 1 1.644 3H13v1H9.973A4.5 4.5 0 0 1 5.5 9H3.852l5.973 5.12l-.65.76l-7-6A.5.5 0 0 1 2.5 8h3a3.5 3.5 0 0 0 3.465-3H2V4h6.965A3.5 3.5 0 0 0 5.5 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RustOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M2 10.5h5m-4.5-6H9a1.5 1.5 0 1 1 0 3H4.5m0-3v6m3-3h1a2 2 0 0 1 2 2a1 1 0 0 0 1 1H13M7.5.5l1.344 1.11l1.693-.417l.73 1.584l1.706.359l-.03 1.743l1.382 1.063L13.54 7.5l.784 1.558l-1.382 1.063l.03 1.743l-1.707.359l-.729 1.584l-1.693-.418L7.5 14.5l-1.344-1.11l-1.693.417l-.73-1.584l-1.706-.359l.03-1.743L.675 9.058L1.46 7.5L.675 5.942L2.057 4.88l-.03-1.743l1.706-.359l.73-1.584l1.693.417z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RustSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.174.115a.521.521 0 0 1 .652 0l1.178.95l1.483-.357a.516.516 0 0 1 .588.276l.64 1.355l1.494.307c.24.05.411.258.407.498l-.027 1.492l1.211.91a.492.492 0 0 1 .145.621L14.26 7.5l.686 1.333a.49.49 0 0 1-.145.62l-.726.547h-2.478a.506.506 0 0 1-.512-.5a2.47 2.47 0 0 0-.861-1.87c.521-.363.86-.958.86-1.63c0-1.105-.916-2-2.047-2H1.399l-.015-.856a.503.503 0 0 1 .407-.498l1.495-.307l.639-1.355a.516.516 0 0 1 .588-.276l1.483.357z"/><path fill="currentColor" d="M.926 5L.2 5.546a.492.492 0 0 0-.145.621L.74 7.5L.055 8.833a.492.492 0 0 0 .145.62L.926 10h2.99V5z"/><path fill="currentColor" d="m1.4 11l-.016.856a.503.503 0 0 0 .407.498l1.495.307l.639 1.355c.103.218.35.334.588.276l1.483-.357l1.178.95a.52.52 0 0 0 .652 0l1.178-.95l1.483.357a.516.516 0 0 0 .588-.276l.64-1.355l1.494-.307a.504.504 0 0 0 .407-.498L13.6 11h-2.005c-.848 0-1.536-.672-1.536-1.5S9.372 8 8.524 8H4.94v2h2.048v1zm7.636-4H4.94V5h4.096c.565 0 1.024.448 1.024 1S9.6 7 9.036 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafariOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 11.5l-.429-.257a.5.5 0 0 0 .686.686zm3-5l-.257-.429l-.107.065l-.065.107zm5-3l.429.257a.5.5 0 0 0-.686-.686zm-3 5l.257.429l.107-.065l.065-.107zm5.5-1A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zM3.929 11.757l3-5l-.858-.514l-3 5zM6.757 6.93l5-3l-.514-.858l-5 3zm4.314-3.686l-3 5l.858.514l3-5zM8.243 8.07l-5 3l.514.858l5-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafariSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.958 10.042l1.906-3.178l3.178-1.906l-1.906 3.178z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5A7.5 7.5 0 0 1 7.5 0A7.5 7.5 0 0 1 15 7.5A7.5 7.5 0 0 1 7.5 15A7.5 7.5 0 0 1 0 7.5m11.929-3.743a.5.5 0 0 0-.686-.686L6.136 6.136L3.07 11.243a.5.5 0 0 0 .686.686l5.107-3.065z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 3v3m0 2v3m0 2.5V15m10-1.5V15m-3-5.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-8-9h12a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 7.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" fill-rule="evenodd" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v11A1.5 1.5 0 0 0 1.5 14H2v1h1v-1h9v1h1v-1h.5a1.5 1.5 0 0 0 1.5-1.5v-11A1.5 1.5 0 0 0 13.5 0zM2 3v3h1V3zm7.5 2a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5M2 11V8h1v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 14.5v-3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v3m3 0h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1h8.586a1 1 0 0 1 .707.293l3.414 3.414a1 1 0 0 1 .293.707V13.5a1 1 0 0 1-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h8.586a1.5 1.5 0 0 1 1.06.44l3.415 3.414A1.5 1.5 0 0 1 15 4.914V13.5a1.5 1.5 0 0 1-1.5 1.5H11v-3.5A1.5 1.5 0 0 0 9.5 10h-4A1.5 1.5 0 0 0 4 11.5V15H1.5A1.5 1.5 0 0 1 0 13.5z"/><path fill="currentColor" d="M5 15h5v-3.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0-.5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScanOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 5V2.5a2 2 0 0 1 2-2H5m5 0h2.5a2 2 0 0 1 2 2V5m-14 5v2.5a2 2 0 0 0 2 2H5m9.5-4.5v2.5a2 2 0 0 1-2 2H10m-8-7h11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScanSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 1A1.5 1.5 0 0 0 1 2.5V5H0V2.5A2.5 2.5 0 0 1 2.5 0H5v1zm10 0H10V0h2.5A2.5 2.5 0 0 1 15 2.5V5h-1V2.5A1.5 1.5 0 0 0 12.5 1m.5 7H2V7h11zM0 12.5V10h1v2.5A1.5 1.5 0 0 0 2.5 14H5v1H2.5A2.5 2.5 0 0 1 0 12.5m14 0V10h1v2.5a2.5 2.5 0 0 1-2.5 2.5H10v-1h2.5a1.5 1.5 0 0 0 1.5-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SchoolOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m7.5 4.5l4 2v8h-8v-8zm0 0V0M0 14.5h15m-13.5 0v-6h2m10 6v-6h-2m-5 6v-3h2v3m-1-14h3v2h-3m0 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SchoolSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 8a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1"/><path fill="currentColor" fill-rule="evenodd" d="m12 6.191l-4-2V3h3V0H7v4.191l-4 2V8H1v6H0v1h6v-4h3v4h6v-1h-1V8h-2zM13 14V9h-1v5zM3 14H2V9h1zm3-5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/><path fill="currentColor" d="M8 15v-3H7v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 14.5h11m-5.5-4v4m-7-13v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v8A1.5 1.5 0 0 0 1.5 11H7v3H2v1h11v-1H8v-3h5.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenAltTwoOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 14.5h11m-7.5-4v4m4-4v4m-9-13v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenAltTwoSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v8A1.5 1.5 0 0 0 1.5 11H5v3H2v1h11v-1h-3v-3h3.5A1.5 1.5 0 0 0 15 9.5v-8A1.5 1.5 0 0 0 13.5 0zM6 14v-3h3v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 14.5h7M.5 2.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 1A1.5 1.5 0 0 0 0 2.5v8A1.5 1.5 0 0 0 1.5 12h12a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 1zM4 15h7v-1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScribbleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 5C3 2 7.3.5 6.5 2.5C5.5 5-.5 9.5 3 11c1.343.576 3.055.45 4.654-.05m0 0C10.222 10.145 12.5 8.377 12.5 7C12.5 4.5 9 5.5 8 9c-.206.722-.328 1.381-.346 1.95Zm0 0C7.584 13.133 9.032 13.983 13 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScribbleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.06 2.334c-1.077.463-2.426 1.515-3.113 2.89l-.894-.448c.813-1.625 2.364-2.823 3.612-3.36c.316-.136.626-.236.912-.287c.278-.05.571-.061.833.017a.923.923 0 0 1 .645.619c.091.3.028.623-.09.92c-.284.708-.897 1.514-1.538 2.302c-.196.241-.396.482-.596.723a35.682 35.682 0 0 0-1.365 1.71c-.593.811-.972 1.501-1.04 2.034c-.032.247.007.434.103.588c.1.16.293.338.668.498c1.104.474 2.543.426 3.98.028a9.65 9.65 0 0 1 .342-1.705c.536-1.876 1.757-3.141 2.93-3.581c.583-.219 1.223-.254 1.743.053c.542.32.808.924.808 1.665c0 .48-.196.947-.483 1.367c-.29.424-.692.834-1.164 1.213c-.86.69-1.99 1.308-3.195 1.73c.02.326.086.581.186.77c.123.234.31.394.603.476c.313.088.77.092 1.418-.062c.643-.154 1.44-.456 2.411-.941l.448.894c-1.013.507-1.885.842-2.627 1.02c-.738.175-1.382.202-1.92.052a1.917 1.917 0 0 1-1.218-.972a2.711 2.711 0 0 1-.279-.946c-1.483.37-3.064.421-4.377-.142c-.5-.214-.885-.505-1.123-.888c-.242-.389-.3-.819-.246-1.244c.103-.81.63-1.683 1.225-2.497c.43-.59.939-1.201 1.425-1.786c.195-.235.386-.465.567-.688c.656-.805 1.168-1.499 1.385-2.042a.832.832 0 0 0 .06-.216a1.008 1.008 0 0 0-.343.015a3.27 3.27 0 0 0-.693.221m3.172 7.88c.952-.375 1.825-.876 2.495-1.414c.419-.336.745-.676.964-.996C11.91 7.48 12 7.209 12 7c0-.509-.171-.718-.317-.804c-.168-.099-.465-.134-.882.022c-.827.31-1.856 1.295-2.32 2.92a9.81 9.81 0 0 0-.249 1.077" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SdCardOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 3v3m2-3v3m2-3v3m-8 8.5h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-7l-4 4v9a1 1 0 0 0 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SdCardSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.293 0H12.5A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V4.293zM6 3v3h1V3zm2 0v3h1V3zm2 3V3h1v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m8.5 8.5l2 2M7 9.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Zm.5 5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M7 4a3 3 0 1 0 1.738 5.445l1.409 1.409l.707-.707l-1.409-1.409A3 3 0 0 0 7 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m14.5 14.5l-4-4m-4 2a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchPropertyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 8.5H4a.5.5 0 0 0 .5.5zm4 0V9a.5.5 0 0 0 .5-.5zm0-2.5H9a.5.5 0 0 0-.146-.354zm-2-2l.354-.354a.5.5 0 0 0-.708 0zm-2 2l-.354-.354A.5.5 0 0 0 4 6zm10.354 8.146l-4-4l-.708.708l4 4zM6.5 12A5.5 5.5 0 0 1 1 6.5H0A6.5 6.5 0 0 0 6.5 13zM12 6.5A5.5 5.5 0 0 1 6.5 12v1A6.5 6.5 0 0 0 13 6.5zM6.5 1A5.5 5.5 0 0 1 12 6.5h1A6.5 6.5 0 0 0 6.5 0zm0-1A6.5 6.5 0 0 0 0 6.5h1A5.5 5.5 0 0 1 6.5 1zm-2 9h4V8h-4zM9 8.5V6H8v2.5zm-.146-2.854l-2-2l-.708.708l2 2zm-2.708-2l-2 2l.708.708l2-2zM4 6v2.5h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchPropertySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8V6.207l1.5-1.5l1.5 1.5V8z"/><path fill="currentColor" fill-rule="evenodd" d="M0 6.5a6.5 6.5 0 1 1 11.436 4.23l3.418 3.416l-.707.707l-3.418-3.417A6.5 6.5 0 0 1 0 6.5m8.854-.854l-2-2a.5.5 0 0 0-.708 0l-2 2A.5.5 0 0 0 4 6v2.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.146-.354" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m8.5 8.5l2 2M7 9.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 4a3 3 0 1 0 1.738 5.445l1.408 1.409l.708-.708l-1.409-1.408A3 3 0 0 0 7 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 0a6.5 6.5 0 1 0 4.23 11.436l3.416 3.418l.708-.708l-3.418-3.417A6.5 6.5 0 0 0 6.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SectionAddOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 .5H1.5a1 1 0 0 0-1 1V4M6 .5h3m2 0h2.5a1 1 0 0 1 1 1V4M.5 6v3m14-3v3m-14 2v2.5a1 1 0 0 0 1 1H4M14.5 11v2.5a1 1 0 0 1-1 1H11M7.5 4v7M4 7.5h7m-5 7h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SectionAddSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0H4v1H1.5a.5.5 0 0 0-.5.5V4H0zM9 1H6V0h3zm4.5 0H11V0h2.5A1.5 1.5 0 0 1 15 1.5V4h-1V1.5a.5.5 0 0 0-.5-.5M7 7V4h1v3h3v1H8v3H7V8H4V7zM0 9V6h1v3zm14 0V6h1v3zM0 13.5V11h1v2.5a.5.5 0 0 0 .5.5H4v1H1.5A1.5 1.5 0 0 1 0 13.5m14 0V11h1v2.5a1.5 1.5 0 0 1-1.5 1.5H11v-1h2.5a.5.5 0 0 0 .5-.5M9 15H6v-1h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SectionRemoveOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 .5H1.5a1 1 0 0 0-1 1V4M6 .5h3m2 0h2.5a1 1 0 0 1 1 1V4M.5 6v3m14-3v3m-14 2v2.5a1 1 0 0 0 1 1H4M14.5 11v2.5a1 1 0 0 1-1 1H11m-7-7h7m-5 7h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SectionRemoveSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0H4v1H1.5a.5.5 0 0 0-.5.5V4H0zM9 1H6V0h3zm4.5 0H11V0h2.5A1.5 1.5 0 0 1 15 1.5V4h-1V1.5a.5.5 0 0 0-.5-.5M0 9V6h1v3zm14 0V6h1v3zm-3-1H4V7h7zM0 13.5V11h1v2.5a.5.5 0 0 0 .5.5H4v1H1.5A1.5 1.5 0 0 1 0 13.5m14 0V11h1v2.5a1.5 1.5 0 0 1-1.5 1.5H11v-1h2.5a.5.5 0 0 0 .5-.5M9 15H6v-1h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendDownOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m7.5 13.5l-4-4m4 4l4-4m-4 4V3M14 1.5H1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendDownSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v1H1zm7 2v9.293l3.146-3.147l.708.708L7.5 14.207L3.146 9.854l.708-.708L7 12.293V3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendLeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m1.5 7.5l4-4m-4 4l4 4m-4-4H12m1.5 6.5V1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendLeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 1h1v13h-1zM2.707 8l3.147 3.146l-.708.708L.793 7.5l4.353-4.354l.708.708L2.707 7H12v1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.5.5l.46.197a.5.5 0 0 0-.657-.657zm-14 6l-.197-.46a.5.5 0 0 0-.06.889zm8 8l-.429.257a.5.5 0 0 0 .889-.06zM14.303.04l-14 6l.394.92l14-6zM.243 6.93l5 3l.514-.858l-5-3zM5.07 9.757l3 5l.858-.514l-3-5zm3.889 4.94l6-14l-.92-.394l-6 14zM14.146.147l-9 9l.708.707l9-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendRightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m13.5 7.5l-4 4m4-4l-4-4m4 4H3M1.5 1v13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendRightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 14V1h1v13zM9.854 3.146L14.207 7.5l-4.353 4.354l-.708-.708L12.293 8H3V7h9.293L9.146 3.854z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.954.71a.5.5 0 0 1-.1.144L5.4 10.306l2.67 4.451a.5.5 0 0 0 .889-.06zM4.694 9.6L.243 6.928a.5.5 0 0 1 .06-.889L14.293.045a.5.5 0 0 0-.146.101z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendUpOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m7.5 1.5l4 4m-4-4l-4 4m4-4V12M1 13.5h13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SendUpSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.5.793l4.354 4.353l-.708.708L8 2.707V12H7V2.707L3.854 5.854l-.708-.708zM14 13v1H1v-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServersOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.5 5.5h-12m12 0a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1m12 0a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1m-12-4a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1m12 0h-12m12 0a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1m.5-6h3m-3 4h3m-3 4h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServersSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h12A1.5 1.5 0 0 1 15 2.5v2c0 .384-.144.735-.382 1c.238.265.382.616.382 1v2c0 .384-.144.735-.382 1c.238.265.382.616.382 1v2a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-2c0-.384.144-.735.382-1A1.494 1.494 0 0 1 0 8.5v-2c0-.384.144-.735.382-1A1.494 1.494 0 0 1 0 4.5zM2 4h3V3H2zm3 4H2V7h3zm-3 4h3v-1H2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M4.5 7.5h4M11 4L8.5 7.495L11 11m3.5-8.501a2 2 0 0 1-4 0a2 2 0 0 1 4 0Zm0 9.993a2 2 0 0 1-4 0a2 2 0 0 1 4 0Zm-10-4.997a2 2 0 0 1-4 0a2 2 0 0 1 4 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2.499a2.5 2.5 0 0 1 5 0a2.5 2.5 0 0 1-3.565 2.26L9.13 7.52l2.038 2.858A2.5 2.5 0 0 1 15 12.493a2.5 2.5 0 1 1-4.559-1.417L8.246 8H4.949A2.501 2.501 0 0 1 0 7.495A2.5 2.5 0 0 1 4.95 7h3.312l2.37-2.84A2.488 2.488 0 0 1 10 2.499"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5 4.5l7-4l7 4v.72a9.651 9.651 0 0 1-7 9.28a9.651 9.651 0 0 1-7-9.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.748.066a.5.5 0 0 0-.496 0l-7 4A.5.5 0 0 0 0 4.5v.72a10.15 10.15 0 0 0 7.363 9.76a.5.5 0 0 0 .274 0A10.152 10.152 0 0 0 15 5.22V4.5a.5.5 0 0 0-.252-.434z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldTickOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M4 7.5L7 10l4-5M7.5.5l-7 4v.72a9.651 9.651 0 0 0 7 9.28a9.651 9.651 0 0 0 7-9.28V4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldTickSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.252.066a.5.5 0 0 1 .496 0l7 4A.5.5 0 0 1 15 4.5v.72a10.15 10.15 0 0 1-7.363 9.76a.5.5 0 0 1-.274 0A10.152 10.152 0 0 1 0 5.22V4.5a.5.5 0 0 1 .252-.434zm-.18 10.645l4.318-5.399l-.78-.624l-3.682 4.601L4.32 7.116l-.64.768z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldXoutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m4.5 4.5l6 6m-6 0l6-6m-3-4l-7 4v.72a9.651 9.651 0 0 0 7 9.28a9.651 9.651 0 0 0 7-9.28V4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldXsolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.252.066a.5.5 0 0 1 .496 0l7 4A.5.5 0 0 1 15 4.5v.72a10.15 10.15 0 0 1-7.363 9.76a.5.5 0 0 1-.274 0A10.152 10.152 0 0 1 0 5.22V4.5a.5.5 0 0 1 .252-.434zm2.895 10.788L7.5 8.207l-2.646 2.647l-.708-.707L6.793 7.5L4.146 4.854l.708-.708L7.5 6.793l2.646-2.647l.708.708L8.207 7.5l2.647 2.646z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 2.5V2a.5.5 0 0 0-.49.402zm12 0l.49-.098A.5.5 0 0 0 13.5 2zm1 5V8a.5.5 0 0 0 .49-.598zm-14 0l-.49-.098A.5.5 0 0 0 .5 8zm3 3H3v.5h.5zm8 0v.5h.5v-.5zM0 15h15v-1H0zm1-7.5v7h1v-7zm12 0v7h1v-7zM1.5 3h12V2h-12zm11.51-.402l1 5l.98-.196l-1-5zM14.5 7H.5v1h14zM.99 7.598l1-5l-.98-.196l-1 5zM1 1h13V0H1zm2 6.5v3h1v-3zm.5 3.5h8v-1h-8zm8.5-.5v-3h-1v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1H1V0h13zM1.01 2.402A.5.5 0 0 1 1.5 2h12a.5.5 0 0 1 .49.402l1 5A.5.5 0 0 1 14.5 8H.5a.5.5 0 0 1-.49-.598zM1 9v5H0v1h15v-1h-1V9h-2v2H3V9z"/><path fill="currentColor" d="M4 9h7v1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 6.5V6a.5.5 0 0 0-.5.5zm10 0h.5a.5.5 0 0 0-.5-.5zm0 6v.5a.5.5 0 0 0 .5-.5zm-10 0H4a.5.5 0 0 0 .5.5zM1 1v14h1V1zM0 4h15V3H0zm4.5 3h10V6h-10zm9.5-.5v6h1v-6zm.5 5.5h-10v1h10zm-9.5.5v-6H4v6zm1-9v3h1v-3zm6 0v3h1v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1v2H0v1h1v11h1V4h4v2H4.5a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5H13V4h2V3H2V1zm6 5V4h5v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SigninOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m10.5 7.5l-3 3.25m3-3.25l-3-3m3 3H1m6-6h6.5v12H7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SigninSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 1h7v13H7v-1h6V2H7zm.854 3.146l3.34 3.34l-3.327 3.603l-.734-.678L9.358 8H1V7h8.293L7.146 4.854z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 5.5h3.5V12M4 8.5h2m-2 3h2m3-6h2m-2 3h2m-2 3h2m1.5 3h-10a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1h7l4 4v9a1 1 0 0 1-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.707 0H2.5A1.5 1.5 0 0 0 1 1.5v12A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V4.293zM7 6H4V5h4v7H7zM6 9H4V8h2zm-2 3h2v-1H4zm7-6H9V5h2zM9 9h2V8H9zm2 3H9v-1h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimohamedOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 5.5v3a6 6 0 0 1-1.2 3.6l-.6.8a4 4 0 0 1-6.4 0l-.6-.8a6 6 0 0 1-1.2-3.6v-3m10 0a5 5 0 0 0-10 0m10 0l-1.121-1.121A3 3 0 0 0 9.257 3.5H5.743a3 3 0 0 0-2.122.879L2.5 5.5m2.5 3h1m3 0h1m-4-2H2.5A1.562 1.562 0 0 0 .985 8.44l.151.605A1.921 1.921 0 0 0 3 10.5m6-4h3.5c1.016 0 1.761.955 1.515 1.94l-.151.605A1.921 1.921 0 0 1 12 10.5M3 11l.837.502a7 7 0 0 0 3.602.998h.122a7 7 0 0 0 3.602-.998L12 11m-6.5 1l1-1h2l1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimohamedSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 9H5V8h1zm4 0H9V8h1z"/><path fill="currentColor" fill-rule="evenodd" d="M2 5.5a5.5 5.5 0 1 1 11 0v.56a2.063 2.063 0 0 1 1.5 2.502l-.151.604a2.422 2.422 0 0 1-1.825 1.777A6.5 6.5 0 0 1 11.7 12.4l-.6.8a4.5 4.5 0 0 1-7.2 0l-.6-.8a6.5 6.5 0 0 1-.823-1.457A2.422 2.422 0 0 1 .65 9.166L.5 8.562A2.063 2.063 0 0 1 2 6.06zM3 7h3V6H3v-.293l.975-.975A2.5 2.5 0 0 1 5.743 4h3.514a2.5 2.5 0 0 1 1.768.732l.975.975V6H9v1h3v1.5a5.5 5.5 0 0 1-.455 2.19l-.64.384c-.35.21-.718.386-1.098.526l-1.1-1.1H6.293l-1.1 1.1a6.5 6.5 0 0 1-1.098-.526l-.64-.384A5.5 5.5 0 0 1 3 8.5zm-1 .124c-.411.22-.653.702-.53 1.195l.151.605c.078.309.253.573.488.762A6.499 6.499 0 0 1 2 8.5zm4.307 4.777c.372.066.75.099 1.132.099h.122c.381 0 .76-.034 1.133-.1l-.401-.4H6.707zm6.584-2.215c.235-.19.41-.453.488-.762l.15-.605A1.062 1.062 0 0 0 13 7.124V8.5c0 .4-.037.797-.11 1.186" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkullOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 11.5H4v-.309l-.276-.138zm-1.447-.724l.223-.447zm10.894 0l-.223-.447zM11.5 11.5l-.224-.447l-.276.138v.309zm-4-2l.354-.354a.5.5 0 0 0-.708 0zM2 9.882V6.5H1v3.382zM13 6.5v3.382h1V6.5zm-9.276 4.553l-1.448-.724l-.447.895l1.447.723zm9-.724l-1.448.724l.448.894l1.447-.723zM4 12.5v-1H3v1zm7-1v1h1v-1zM9.5 14h-4v1h4zm1.5-1.5A1.5 1.5 0 0 1 9.5 14v1a2.5 2.5 0 0 0 2.5-2.5zm2-2.618a.5.5 0 0 1-.276.447l.447.895A1.5 1.5 0 0 0 14 9.882zM7.5 1A5.5 5.5 0 0 1 13 6.5h1A6.5 6.5 0 0 0 7.5 0zM3 12.5A2.5 2.5 0 0 0 5.5 15v-1A1.5 1.5 0 0 1 4 12.5zm-1-6A5.5 5.5 0 0 1 7.5 1V0A6.5 6.5 0 0 0 1 6.5zM1 9.882a1.5 1.5 0 0 0 .83 1.342l.446-.895A.5.5 0 0 1 2 9.882zM4.5 8a.5.5 0 0 1-.5-.5H3A1.5 1.5 0 0 0 4.5 9zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 6 7.5zM4.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 4.5 6zm0-1A1.5 1.5 0 0 0 3 7.5h1a.5.5 0 0 1 .5-.5zm6 2a.5.5 0 0 1-.5-.5H9A1.5 1.5 0 0 0 10.5 9zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 12 7.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 6zm0-1A1.5 1.5 0 0 0 9 7.5h1a.5.5 0 0 1 .5-.5zm-3.646 4.854l1-1l-.708-.708l-1 1zm.292-1l1 1l.708-.708l-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkullSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 7.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0m6 0a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0"/><path fill="currentColor" fill-rule="evenodd" d="M1 6.5a6.5 6.5 0 0 1 13 0v3.382a1.5 1.5 0 0 1-.83 1.342l-1.17.585v.691A2.5 2.5 0 0 1 9.5 15h-4A2.5 2.5 0 0 1 3 12.5v-.691l-1.17-.585A1.5 1.5 0 0 1 1 9.882zM4.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m6 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M7.146 9.146a.5.5 0 0 1 .708 0l1 1l-.708.708l-.646-.647l-.646.647l-.708-.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkypeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.83 8.34l-.496-.066a.5.5 0 0 0 .08.346zM6.66 1.17l-.28.415a.5.5 0 0 0 .344.081zM1.17 6.66l.496.065a.5.5 0 0 0-.08-.345zm7.17 7.17l.279-.415a.5.5 0 0 0-.345-.08zm5.985-5.423c.039-.29.066-.593.066-.907h-1c0 .257-.022.513-.057.774zm.066-.907A6.892 6.892 0 0 0 7.5.609v1A5.892 5.892 0 0 1 13.391 7.5zM7.5.609a7.07 7.07 0 0 0-.905.065l.129.992c.264-.034.52-.057.776-.057zm-.562.146A4.437 4.437 0 0 0 4.457 0v1c.712 0 1.374.216 1.923.585zM4.457 0A4.456 4.456 0 0 0 0 4.457h1A3.456 3.456 0 0 1 4.457 1zM0 4.457c0 .918.279 1.772.755 2.481l.83-.558A3.436 3.436 0 0 1 1 4.457zm.675 2.137a6.88 6.88 0 0 0-.066.906h1c0-.257.022-.513.057-.775zM.609 7.5A6.891 6.891 0 0 0 7.5 14.391v-1A5.891 5.891 0 0 1 1.609 7.5zM7.5 14.391c.314 0 .616-.027.906-.066l-.132-.99a5.888 5.888 0 0 1-.774.056zm.561-.146c.71.477 1.564.755 2.483.755v-1a3.439 3.439 0 0 1-1.925-.585zm2.483.755A4.457 4.457 0 0 0 15 10.544h-1A3.457 3.457 0 0 1 10.544 14zM15 10.544a4.44 4.44 0 0 0-.756-2.482l-.83.558c.37.55.586 1.21.586 1.924zM7 8h1V7H7zm2.947-2.724A2.309 2.309 0 0 0 7.882 4v1c.496 0 .949.28 1.17.724zM9 9a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2zM8 8a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2zm-.882 2c-.496 0-.95-.28-1.17-.724l-.895.448A2.309 2.309 0 0 0 7.118 11zM5 6a2 2 0 0 0 2 2V7a1 1 0 0 1-1-1zm1 0a1 1 0 0 1 1-1V4a2 2 0 0 0-2 2zm2 4h-.882v1H8zm-.118-6H7v1h.882z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkypeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 4.457A4.456 4.456 0 0 1 6.778.652A6.46 6.46 0 0 1 7.5.61a6.892 6.892 0 0 1 6.891 6.89c0 .248-.017.49-.043.723a4.457 4.457 0 0 1-6.126 6.125a6.477 6.477 0 0 1-.722.043A6.891 6.891 0 0 1 .609 7.5c0-.248.017-.49.043-.723A4.435 4.435 0 0 1 0 4.457M6 6a1 1 0 0 1 1-1h.882c.496 0 .949.28 1.17.724l.895-.448A2.309 2.309 0 0 0 7.882 4H7a2 2 0 1 0 0 4h1a1 1 0 0 1 0 2h-.882c-.496 0-.95-.28-1.17-.724l-.895.448A2.309 2.309 0 0 0 7.118 11H8a2 2 0 1 0 0-4H7a1 1 0 0 1-1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.5 7.5V6A1.5 1.5 0 1 1 12 7.5zm0 0h-3m3 0V2a1.5 1.5 0 1 0-3 0v5.5m0 0v-3m0 3H2a1.5 1.5 0 1 1 0-3h5.5m0 3H13a1.5 1.5 0 0 1 0 3H7.5m0-3v3m0-3h-3m3 0V13a1.5 1.5 0 0 1-3 0V7.5m3-3V3A1.5 1.5 0 1 0 6 4.5zm0 6H9A1.5 1.5 0 1 1 7.5 12zm-3-3V9A1.5 1.5 0 1 1 3 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.385 6.923H8.077v-5.77a1.154 1.154 0 0 1 2.308 0zm2.307 0H11.54V5.77a1.154 1.154 0 1 1 1.153 1.154m1.153 1.153h-5.77v2.308h5.77a1.154 1.154 0 0 0 0-2.308m-5.769 4.615V11.54H9.23a1.154 1.154 0 1 1-1.154 1.153M1.154 4.615h5.77v2.308h-5.77a1.154 1.154 0 0 1 0-2.308m5.769-2.307v1.154H5.77a1.154 1.154 0 1 1 1.154-1.154M1.154 9.23c0-.636.516-1.153 1.154-1.153h1.154V9.23a1.154 1.154 0 0 1-2.308 0m3.461 4.616v-5.77h2.308v5.77a1.154 1.154 0 0 1-2.308 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 1c-1.155 0-2.174.412-2.894 1.281c-.642.775-1.006 2.35-1.066 3.666l-.073.01l-.022.004a8.68 8.68 0 0 0-.368.059c-.465.089-1.346.326-1.543 1.277c-.093.445.011.833.247 1.134c.211.269.497.429.708.53c.09.041.181.08.274.117c-.21.584-.579 1.184-.987 1.728c-.382.508-.28 1.153-.083 1.573c.197.421.57.402 1.192.43c.352.015.722.051 1.09.12c.166.03.362.098.606.2c.142.06.283.123.423.187c.113.052.235.106.374.167c.573.25 1.276.517 2.056.517s1.483-.267 2.055-.517c.14-.06.26-.115.375-.167l.025-.012c.135-.06.26-.117.398-.174c.243-.103.44-.17.606-.201a7.951 7.951 0 0 1 1.09-.12c.622-.028 1.104-.009 1.303-.43c.197-.42.298-1.065-.084-1.573c-.406-.54-.772-1.136-.983-1.716a5.24 5.24 0 0 0 .305-.127c.216-.098.518-.261.73-.543c.245-.326.315-.739.175-1.184c-.28-.886-1.092-1.122-1.568-1.216a6.857 6.857 0 0 0-.355-.058l-.012-.002l-.056-.009c-.065-1.234-.41-2.795-1.036-3.581C9.695 1.485 8.682 1 7.5 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0C6.249 0 5.144.476 4.365 1.479c-.696.893-1.09 2.712-1.155 4.23l-.08.012l-.022.003c-.117.018-.256.04-.4.068c-.504.103-1.458.376-1.672 1.474c-.1.513.013.961.268 1.308c.23.31.54.495.768.61c.097.05.196.095.296.137c-.228.673-.627 1.366-1.069 1.993c-.414.587-.304 1.33-.09 1.816c.213.486.617.464 1.291.495c.382.018.783.06 1.181.139c.18.035.393.114.657.232c.153.07.306.14.458.215c.123.06.254.123.405.193c.62.288 1.383.596 2.227.596c.845 0 1.607-.308 2.227-.596c.151-.07.282-.133.406-.193l.027-.014a16.8 16.8 0 0 1 .43-.201c.265-.118.477-.196.657-.232a8.09 8.09 0 0 1 1.18-.139c.675-.031 1.198-.01 1.413-.495c.213-.485.323-1.229-.09-1.816c-.44-.623-.837-1.31-1.066-1.98a4.72 4.72 0 0 0 .33-.146c.234-.113.562-.301.79-.626c.267-.377.342-.853.19-1.366c-.302-1.023-1.182-1.295-1.697-1.403a6.404 6.404 0 0 0-.385-.068l-.014-.002l-.06-.01c-.07-1.424-.443-3.225-1.123-4.132C9.878.56 8.78 0 7.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnesOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 7v3M2 8.5h3m6 0h1m-1-2h1m-3 2h1m-1-2h1m-8.5-3h12a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnesSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 4.5A1.5 1.5 0 0 1 1.5 3h12A1.5 1.5 0 0 1 15 4.5v6a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 10.5zM12 7h-1V6h1zm-9 3V9H2V8h1V7h1v1h1v1H4v1zm8-1h1V8h-1zM9 9h1V8H9zm1-2H9V6h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphabeticallyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 14.5l-.354.354a.5.5 0 0 0 .708 0zm6-6V8H9v.5zm0 6H9v.5h.5zm-5.646.354l3-3l-.708-.708l-3 3zm0-.708l-3-3l-.708.708l3 3zM3 0v14.5h1V0zm6.5 9H12V8H9.5zm2.5 2H9.5v1H12zm-2 .5v-3H9v3zm3-1.5a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2zm-1-1a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2zm0 5H9.5v1H12zm-2 .5v-3H9v3zm3-1.5a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2zm-1-1a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2zm-2-5V2.5H9V7zm3-4.5V7h1V2.5zM11.5 1A1.5 1.5 0 0 1 13 2.5h1A2.5 2.5 0 0 0 11.5 0zM10 2.5A1.5 1.5 0 0 1 11.5 1V0A2.5 2.5 0 0 0 9 2.5zM9.5 5h4V4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphabeticallySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 13.293V0h1v13.293l2.146-2.147l.708.708l-3 3a.5.5 0 0 1-.708 0l-3-3l.708-.708zM11.5 1A1.5 1.5 0 0 0 10 2.5V4h3V2.5A1.5 1.5 0 0 0 11.5 1M13 5v2h1V2.5a2.5 2.5 0 0 0-5 0V7h1V5zM9 8h3a2 2 0 0 1 1.323 3.5A2 2 0 0 1 12 15H9zm3 3a1 1 0 1 0 0-2h-2v2zm-2 1h2a1 1 0 1 1 0 2h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDownOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 14.5l-.354.354a.5.5 0 0 0 .708 0zm.354.354l3-3l-.708-.708l-3 3zm0-.708l-3-3l-.708.708l3 3zM3 0v14.5h1V0zm6 4h6V3H9zm0 4h4V7H9zm0 4h2v-1H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDownSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 13.293V0h1v13.293l2.146-2.147l.708.708l-3 3a.5.5 0 0 1-.708 0l-3-3l.708-.708zM15 4H9V3h6zm-2 4H9V7h4zm-2 4H9v-1h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortHighToLowOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5.5h.5a.5.5 0 0 0-.5-.5zm1 11l.474.158zm.482-1.446l-.474-.158zM4.5 14.5l-.354.354a.5.5 0 0 0 .708 0zM10 1h1.5V0H10zm1-.5v6h1v-6zM10 7h3V6h-3zm1.862 1H11v1h.862zm1.112 3.658l.482-1.446l-.948-.316l-.482 1.446zM11 12h1.5v-1H11zm.974 2.658l1-3l-.948-.316l-1 3zM9 10a2 2 0 0 0 2 2v-1a1 1 0 0 1-1-1zm2-2a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1zm.862 1a.68.68 0 0 1 .646.896l.948.316A1.68 1.68 0 0 0 11.862 8zm-7.008 5.854l3-3l-.708-.708l-3 3zm0-.708l-3-3l-.708.708l3 3zM4 0v14.5h1V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortHighToLowSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 13.293V0h1v13.293l2.146-2.147l.708.708l-3 3a.5.5 0 0 1-.708 0l-3-3l.708-.708zM11 1h-1V0h1.5a.5.5 0 0 1 .5.5V6h1v1h-3V6h1zm-2 9a2 2 0 0 1 2-2h.862a1.68 1.68 0 0 1 1.594 2.212l-1.482 4.446l-.948-.316l.78-2.342H11a2 2 0 0 1-2-2m3.14 1l.368-1.104A.68.68 0 0 0 11.862 9H11a1 1 0 1 0 0 2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortLowToHighOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5.5h.5a.5.5 0 0 0-.5-.5zm-7 0l.354-.354a.5.5 0 0 0-.708 0zM10 1h1.5V0H10zm1-.5v6h1v-6zM10 7h3V6h-3zm1.578 1H11v1h.578zm1.396 3.658l.393-1.176l-.95-.317l-.391 1.177zM11 12h1.5v-1H11zm.974 2.658l1-3l-.948-.316l-1 3zM9 10a2 2 0 0 0 2 2v-1a1 1 0 0 1-1-1zm2-2a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1zm.578 1a.885.885 0 0 1 .84 1.165l.949.317A1.885 1.885 0 0 0 11.578 8zM4.146.146l-3 3l.708.708l3-3zm0 .708l3 3l.708-.708l-3-3zM4 .5V15h1V.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortLowToHighSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.146.146a.5.5 0 0 1 .708 0l3 3l-.708.708L5 1.707V15H4V1.707L1.854 3.854l-.708-.708zM11 1h-1V0h1.5a.5.5 0 0 1 .5.5V6h1v1h-3V6h1zm-2 9a2 2 0 0 1 2-2h.578a1.885 1.885 0 0 1 1.789 2.482l-1.393 4.176l-.948-.316l.78-2.342H11a2 2 0 0 1-2-2m3.14 1l.278-.835A.885.885 0 0 0 11.578 9H11a1 1 0 1 0 0 2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortReverseAlphabeticallyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5.5l.354-.354a.5.5 0 0 0-.708 0zm6 8V8H9v.5zm0 6H9v.5h.5zM3.146.146l-3 3l.708.708l3-3zm0 .708l3 3l.708-.708l-3-3zM3 .5V15h1V.5zM9.5 9H12V8H9.5zm2.5 2H9.5v1H12zm-2 .5v-3H9v3zm3-1.5a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2zm-1-1a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2zm0 5H9.5v1H12zm-2 .5v-3H9v3zm3-1.5a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2zm-1-1a1 1 0 0 1 1 1h1a2 2 0 0 0-2-2zm-2-5V2.5H9V7zm3-4.5V7h1V2.5zM11.5 1A1.5 1.5 0 0 1 13 2.5h1A2.5 2.5 0 0 0 11.5 0zM10 2.5A1.5 1.5 0 0 1 11.5 1V0A2.5 2.5 0 0 0 9 2.5zM9.5 5h4V4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortReverseAlphabeticallySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.146.146a.5.5 0 0 1 .708 0l3 3l-.708.708L4 1.707V15H3V1.707L.854 3.854l-.708-.708zM11.5 1A1.5 1.5 0 0 0 10 2.5V4h3V2.5A1.5 1.5 0 0 0 11.5 1M13 5v2h1V2.5a2.5 2.5 0 0 0-5 0V7h1V5zM9 8h3a2 2 0 0 1 1.323 3.5A2 2 0 0 1 12 15H9zm3 3a1 1 0 1 0 0-2h-2v2zm-2 1h2a1 1 0 1 1 0 2h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortUpOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5.5l.354-.354a.5.5 0 0 0-.708 0zM3.146.146l-3 3l.708.708l3-3zm0 .708l3 3l.708-.708l-3-3zM3 .5V15h1V.5zM9 4h6V3H9zm0 4h4V7H9zm0 4h2v-1H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortUpSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.146.146a.5.5 0 0 1 .708 0l3 3l-.708.708L4 1.707V15H3V1.707L.854 3.854l-.708-.708zM15 4H9V3h6zm-2 4H9V7h4zm-2 4H9v-1h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundOffOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5zm0-5.996v.5h.138l.12-.071zm5-2.998H9a.5.5 0 0 0-.757-.429zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492zm2.94-7.763l-.354-.353l-.707.707l.354.354zm2.12 3.534l.354.353l.707-.707l-.354-.353zm.708-2.826l.353-.354l-.707-.707l-.353.353zm-3.535 2.119l-.354.353l.707.707l.354-.353zM3.5 9.994h-2v1h2zm-2 0a.499.499 0 0 1-.5-.5H0c0 .83.671 1.5 1.5 1.5zm-.5-.5V5.498H0v3.998zm0-3.997c0-.276.223-.499.5-.499v-1c-.829 0-1.5.67-1.5 1.5zm.5-.499h2v-1h-2zm2.257-.071l5-2.998l-.514-.858l-5 2.998zM8 1.5v11.992h1V1.5zm.757 11.563l-5-2.998l-.514.858l5 2.998zm1.976-6.626l2.827 2.826l.707-.707l-2.828-2.827zm2.828-.708l-2.828 2.827l.707.707l2.828-2.826z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundOffSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5c-.829 0-1.5.67-1.5 1.5V9.5c0 .829.67 1.5 1.5 1.5h1.862l4.88 2.929A.5.5 0 0 0 9 13.5zm4.208 5.996L14.62 8.91l-.707.707L12.5 8.203l-1.414 1.413l-.707-.707l1.414-1.413l-1.414-1.413l.707-.707L12.5 6.789l1.415-1.413l.706.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundOnOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5zm0-5.996v.5h.138l.12-.071zm5-2.998H9a.5.5 0 0 0-.757-.429zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492zm5.353-9.348l-.353-.353l-.707.707l.354.354zm-.706 5.996l-.354.354l.707.707l.353-.353zM3.5 9.994h-2v1h2zm-2 0a.5.5 0 0 1-.5-.5H0c0 .83.672 1.5 1.5 1.5zm-.5-.5V5.498H0v3.998zm0-3.997a.5.5 0 0 1 .5-.499v-1a1.5 1.5 0 0 0-1.5 1.5zm.5-.499h2v-1h-2zm2.257-.071l5-2.998l-.514-.858l-5 2.998zM8 1.5v11.992h1V1.5zm.757 11.563l-5-2.998l-.514.858l5 2.998zM13.5 4.498c-.353.354-.354.354-.354.353a.01.01 0 0 1-.002-.002l.003.003l.02.022a3.186 3.186 0 0 1 .386.597c.22.439.447 1.112.447 2.025h1c0-1.086-.272-1.911-.553-2.472a4.198 4.198 0 0 0-.39-.639a2.932 2.932 0 0 0-.181-.217l-.014-.015l-.005-.005l-.002-.002l-.001-.001zm.5 2.998c0 .913-.228 1.586-.447 2.025a3.184 3.184 0 0 1-.386.597a.83.83 0 0 1-.02.022l-.003.003l.001-.001l.001-.001l.354.353c.353.354.354.354.354.353h.001l.002-.003l.005-.005l.014-.014l.043-.048c.035-.04.082-.097.137-.17c.11-.146.251-.36.391-.639c.28-.56.553-1.386.553-2.472z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundOnSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.746 1.065A.5.5 0 0 1 9 1.5v12a.5.5 0 0 1-.757.429L3.362 11H1.5A1.5 1.5 0 0 1 0 9.5V5.497a1.5 1.5 0 0 1 1.5-1.499h1.862l4.88-2.927a.5.5 0 0 1 .504-.006m5.108 3.079l-.354-.353l-.707.707l.351.352l.003.002l.02.022a3.194 3.194 0 0 1 .386.597c.22.439.447 1.112.447 2.025c0 .913-.228 1.586-.447 2.025a3.19 3.19 0 0 1-.297.486a1.988 1.988 0 0 1-.11.133l-.002.003l-.351.35l.707.708l.354-.353l-.354-.354l.354.354l.001-.002l.002-.002l.005-.005l.014-.014l.043-.048c.035-.04.082-.097.137-.17c.11-.146.251-.36.391-.639c.28-.56.553-1.386.553-2.472s-.272-1.911-.553-2.472a4.19 4.19 0 0 0-.39-.639a2.89 2.89 0 0 0-.181-.217l-.014-.015l-.005-.005l-.002-.002c0-.001-.002-.002-.355.352z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpotifyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm-3.838 9.116c.986-.24 2.2-.345 3.392-.228c1.196.117 2.334.454 3.202 1.065l.575-.819c-1.053-.74-2.375-1.113-3.679-1.241a11.446 11.446 0 0 0-3.727.252zm-.336-2.124c3.446-.61 5.848-.297 7.839 1.132l.583-.813C9.45 6.661 6.732 6.374 3.152 7.008zm-.225-2.151c1.353-.478 3.003-.676 4.624-.544c1.623.133 3.18.595 4.364 1.4l.561-.828c-1.364-.927-3.099-1.426-4.843-1.568c-1.746-.143-3.539.067-5.039.597z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpotifySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m7.726-2.203c-1.621-.132-3.272.066-4.625.544l-.333-.943c1.5-.53 3.293-.74 5.04-.597c1.743.142 3.478.641 4.843 1.568l-.562.827c-1.185-.804-2.74-1.266-4.363-1.399m-4.4 2.695c3.446-.61 5.848-.297 7.839 1.132l.583-.813C9.45 6.661 6.732 6.374 3.152 7.008zm.336 2.124c.986-.24 2.2-.345 3.392-.228c1.196.117 2.334.454 3.202 1.065l.575-.819c-1.053-.74-2.375-1.113-3.68-1.241a11.446 11.446 0 0 0-3.726.252z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpreadsheetOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 6.995H4v1h.5zm6 1h.5v-1h-.5zm-6 2.505H4v.5h.5zm6 0v.5h.5v-.5zm-6-6.503H4v1h.5zm6 1h.5v-1h-.5zm3-1.497h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-6 7.495h6v-1h-6zM4.5 11h6v-1h-6zm0-6.003h6v-1h-6zm8 9.003h-10v1h10zM2 13.5v-12H1v12zm11-10v10h1v-10zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zm2 3v6h1v-6zm3 0v6h1v-6zm3 0v6h1v-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpreadsheetSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 7.995V10H8V7.995zm0-2.998v1.998H8V4.997zm-3 0H5v1.998h2zm0 2.998H5V10h2z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zm10 2.497H4V11h7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.5.5h-12a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v12A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackoverflowOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 9v5.5h10V9M4 12.5h7M4 10l7 .7M4.3 7.5L11 9M5.3 4.5l6.1 3.1M7 2l5 4.5M9.5.5l3.5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackoverflowSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12.59 5.787l-3.5-5l.82-.574l3.5 5zm-.925 1.085l-5-4.5l.67-.744l5 4.5zm-.491 1.174l-6.1-3.1l.453-.892l6.1 3.1zm-.283 1.442l-6.7-1.5l.218-.976l6.7 1.5zM2 9h1v5h9V9h1v6H2zm8.95 2.198l-7-.7l.1-.996l7 .7zM11 13H4v-1h7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StampOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 14.5h15m-8.5-8v3m2-3v3m-1-3a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm-7 6v-1a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StampSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3.5a3.5 3.5 0 1 1 5 3.163V9h3.5a2.5 2.5 0 0 1 2.5 2.5V13H0v-1.5A2.5 2.5 0 0 1 2.5 9H6V6.663A3.5 3.5 0 0 1 4 3.5M0 14v1h15v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 9.804l.242-.437a.5.5 0 0 0-.484 0zM5.337 11l-.494-.08a.5.5 0 0 0 .736.518zm.413-2.533l.493.08a.5.5 0 0 0-.135-.429zM4 6.674l-.075-.495a.5.5 0 0 0-.283.844zm2.418-.37l.076.495a.5.5 0 0 0 .377-.282zM7.5 4l.453-.212a.5.5 0 0 0-.906 0zm1.082 2.304l-.453.213a.5.5 0 0 0 .377.282zm2.418.37l.358.349a.5.5 0 0 0-.283-.844zM9.25 8.467l-.358-.349a.5.5 0 0 0-.135.43zM9.663 11l-.242.438a.5.5 0 0 0 .736-.519zM7.258 9.367l-2.163 1.195l.484.876l2.163-1.196zM5.83 11.08l.413-2.532l-.986-.161l-.414 2.532zm.278-2.962l-1.75-1.794l-.716.699l1.75 1.793zm-2.033-.95l2.419-.37l-.151-.988l-2.418.37zm2.796-.651l1.082-2.305l-.906-.424l-1.081 2.304zm.176-2.305L8.13 6.517l.905-.425l-1.081-2.304zM8.507 6.8l2.418.369l.15-.989l-2.418-.369zm2.135-.475l-1.75 1.794l.716.698l1.75-1.793zM8.757 8.548l.413 2.533l.987-.162l-.414-2.532zm1.148 2.014L7.742 9.367l-.484.875l2.163 1.196zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m7.5-4a.5.5 0 0 1 .453.288L8.92 5.85l2.155.33a.5.5 0 0 1 .282.843L9.784 8.636l.373 2.284a.5.5 0 0 1-.736.518L7.5 10.376l-1.921 1.062a.5.5 0 0 1-.736-.519l.373-2.283l-1.574-1.613a.5.5 0 0 1 .283-.844l2.154-.329l.968-2.062A.5.5 0 0 1 7.5 3.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 12.04l-4.326 2.275L4 9.497L.5 6.086l4.837-.703L7.5 1l2.163 4.383l4.837.703L11 9.497l.826 4.818z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 9.804L5.337 11l.413-2.533L4 6.674l2.418-.37L7.5 4l1.082 2.304l2.418.37l-1.75 1.793L9.663 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.953 3.788a.5.5 0 0 0-.906 0L6.08 5.85l-2.154.33a.5.5 0 0 0-.283.843l1.574 1.613l-.373 2.284a.5.5 0 0 0 .736.518l1.92-1.062l1.921 1.062a.5.5 0 0 0 .736-.519l-.373-2.283l1.574-1.613a.5.5 0 0 0-.283-.844L8.922 5.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.948.779a.5.5 0 0 0-.896 0L5.005 4.926l-4.577.665a.5.5 0 0 0-.277.853l3.312 3.228l-.782 4.559a.5.5 0 0 0 .725.527L7.5 12.605l4.094 2.153a.5.5 0 0 0 .725-.527l-.782-4.56l3.312-3.227a.5.5 0 0 0-.277-.853l-4.577-.665z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M.5 7.5a7 7 0 1 1 14 0a7 7 0 0 1-14 0Z"/><path d="M9.5 5.5h-4v4h4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M5 5h5v5H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.5 3.5h-8v8h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 5.5h-4v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5H5v5h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3H3v9h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopwatchOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 5v3.5H10m-4-8h3m-1.5 2a6 6 0 1 0 0 12a6 6 0 0 0 0-12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopwatchSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1H6V0h3z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 2a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13M8 8V5H7v4h3V8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrikethroughOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 10A3.5 3.5 0 0 0 7 13.5h1.487a3.013 3.013 0 0 0 3.013-3.013c0-1.205-.892-2.512-2-2.987M6.08 6.177A2.607 2.607 0 0 1 4.5 3.781A2.281 2.281 0 0 1 6.781 1.5H8A2.5 2.5 0 0 1 10.5 4M2 7.5h11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrikethroughSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.781A2.781 2.781 0 0 1 6.781 1H8a3 3 0 0 1 3 3h-1a2 2 0 0 0-2-2H6.781A1.78 1.78 0 0 0 5 3.781c0 .843.502 1.605 1.277 1.937l-.394.919A3.107 3.107 0 0 1 4 3.78M9.392 8H2V7h11v1h-2.01c.123.14.237.287.34.44c.406.602.67 1.326.67 2.047A3.513 3.513 0 0 1 8.487 14H7a4 4 0 0 1-4-4h1a3 3 0 0 0 3 3h1.487A2.513 2.513 0 0 0 11 10.487c0-.484-.182-1.017-.5-1.488c-.296-.44-.69-.797-1.108-.999" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubscriptOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.5 14.5l-.354-.354A.5.5 0 0 0 12.5 15zM15 14h-2.5v1H15zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792zM13.793 11H13.5v1h.293zm-.293 0a1.5 1.5 0 0 0-1.5 1.5h1a.5.5 0 0 1 .5-.5zm1.5 1.207C15 11.54 14.46 11 13.793 11v1c.114 0 .207.093.207.207zm-.354.854c.227-.227.354-.534.354-.854h-1a.207.207 0 0 1-.06.147zM1.9 13.8l9-12l-.8-.6l-9 12zm-.8-12l9 12l.8-.6l-9-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubscriptSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 8.333L1.9 13.8l-.8-.6l4.275-5.7L1.1 1.8l.8-.6L6 6.667L10.1 1.2l.8.6l-4.275 5.7l4.275 5.7l-.8.6zm6 4.167a1.5 1.5 0 0 1 1.5-1.5h.293a1.207 1.207 0 0 1 .854 2.06l-.94.94H15v1h-2.5a.5.5 0 0 1-.354-.854l1.793-1.792a.207.207 0 0 0-.146-.354H13.5a.5.5 0 0 0-.5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M7.5 1.5v-1m0 13.99v-.998m6-5.997h1m-13 0h-1m2-4.996l-1-1m12 0l-1 1m-10 9.993l-1 1m12 0l-1-1m-2-4.997a2.999 2.999 0 0 1-3 2.998a2.999 2.999 0 1 1 3-2.998Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2V0H7v2zm-4.793.498L1.5.792L.793 1.5L2.5 3.206zm9.293.708L14.207 1.5L13.5.792l-1.707 1.706zm-5 .791a3.499 3.499 0 1 0 0 6.996a3.499 3.499 0 1 0 0-6.996M2 6.995H0v1h2zm13 0h-2v1h2zM1.5 14.199l1.707-1.707l-.707-.707l-1.707 1.706zm12.707-.708L12.5 11.785l-.707.707L13.5 14.2zM8 14.99v-1.998H7v1.999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SuperscriptOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.5 3.5l-.354-.354A.5.5 0 0 0 12.5 4zm1.793-1.793l-.354-.353zM15 3h-2.5v1H15zm-2.146.854l1.792-1.793l-.707-.707l-1.793 1.792zM13.793 0H13.5v1h.293zM13.5 0A1.5 1.5 0 0 0 12 1.5h1a.5.5 0 0 1 .5-.5zM15 1.207C15 .54 14.46 0 13.793 0v1c.114 0 .207.093.207.207zm-.354.854c.227-.227.354-.534.354-.854h-1a.207.207 0 0 1-.06.147zM1.1 1.8l9 12l.8-.6l-9-12zm9-.6l-9 12l.8.6l9-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SuperscriptSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.5A1.5 1.5 0 0 1 13.5 0h.293a1.207 1.207 0 0 1 .854 2.06l-.94.94H15v1h-2.5a.5.5 0 0 1-.354-.854l1.793-1.792A.207.207 0 0 0 13.793 1H13.5a.5.5 0 0 0-.5.5zm-6.625 6L1.1 1.8l.8-.6L6 6.667L10.1 1.2l.8.6l-4.275 5.7l4.275 5.7l-.8.6L6 8.333L1.9 13.8l-.8-.6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SvelteOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m9.625 8.357l-5.088 3.18m2.968-1.855a3.5 3.5 0 0 1-3.71-5.937l4.241-2.65A3.5 3.5 0 0 1 12.405 6.5M7.536 5.296a3.5 3.5 0 0 1 3.71 5.936l-4.24 2.65A3.5 3.5 0 0 1 2.614 8.5m2.8-1.88l5.09-3.179"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SvelteSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.283 1.944a4.001 4.001 0 0 1-.27 4.622a4 4 0 0 1-1.503 5.09l-4.24 2.65A4 4 0 0 1 2.028 8.41a4 4 0 0 1 1.503-5.09L7.77.671a4 4 0 0 1 5.512 1.273M8.3 1.52a3 3 0 0 1 4.13 4.143a3.993 3.993 0 0 0-2.386-1.345l.724-.453l-.53-.848l-5.088 3.18l.53.847L7.8 5.72a3 3 0 1 1 3.18 5.088l-4.24 2.65a3 3 0 0 1-4.13-4.143a3.993 3.993 0 0 0 2.386 1.345l-.725.452l.53.848L9.89 8.78l-.53-.847l-2.12 1.325a3 3 0 0 1-3.18-5.089z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SvgOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-8 6V6H2v.5zm0 2H2V9h.5zm2 0H5V8h-.5zm0 2v.5H5v-.5zm2-1H6v.207l.146.147zm1 1l-.354.354l.354.353l.354-.353zm1-1l.354.354L9 9.707V9.5zm2-3V6H10v.5zm0 4H10v.5h.5zm2 0v.5h.5v-.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zM5 6H2.5v1H5zm-3 .5v2h1v-2zM2.5 9h2V8h-2zM4 8.5v2h1v-2zm.5 1.5H2v1h2.5zM6 6v3.5h1V6zm.146 3.854l1 1l.708-.708l-1-1zm1.708 1l1-1l-.708-.708l-1 1zM9 9.5V6H8v3.5zM13 6h-2.5v1H13zm-3 .5v4h1v-4zm.5 4.5h2v-1h-2zm2.5-.5v-2h-1v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SvgSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM5 6H2v3h2v1H2v1h3V8H3V7h2zm2 0H6v3.707l1.5 1.5l1.5-1.5V6H8v3.293l-.5.5l-.5-.5zm3 0h3v1h-2v3h1V8.5h1V11h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 4.5h14m-10-4v14m-3-14h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-12a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1.5A1.5 1.5 0 0 1 1.5 0H4v4H0zM0 5v8.5A1.5 1.5 0 0 0 1.5 15H4V5zm5 10h8.5a1.5 1.5 0 0 0 1.5-1.5V5H5zM15 4V1.5A1.5 1.5 0 0 0 13.5 0H5v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 11.5h3m-6.5 3h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h10A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM6 12h3v-1H6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 7.5l-.354-.354a.5.5 0 0 0 0 .708zm7 7l-.354.354a.5.5 0 0 0 .708 0zm7-7l.354.354A.5.5 0 0 0 15 7.5zm-7-7V0a.5.5 0 0 0-.354.146zM.146 7.854l7 7l.708-.708l-7-7zm7.708 7l7-7l-.708-.708l-7 7zM15 7.5v-6h-1v6zM13.5 0h-6v1h6zM7.146.146l-7 7l.708.708l7-7zM15 1.5A1.5 1.5 0 0 0 13.5 0v1a.5.5 0 0 1 .5.5zM10.5 5a.5.5 0 0 1-.5-.5H9A1.5 1.5 0 0 0 10.5 6zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 12 4.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 3zm0-1A1.5 1.5 0 0 0 9 4.5h1a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0"/><path fill="currentColor" fill-rule="evenodd" d="M7.146.146A.5.5 0 0 1 7.5 0h6A1.5 1.5 0 0 1 15 1.5v6a.5.5 0 0 1-.146.354l-7 7a.5.5 0 0 1-.708 0l-7-7a.5.5 0 0 1 0-.708zM10.5 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TailwindOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M7.5 3C5.633 3 4.467 4 4 6c.7-1 1.517-1.375 2.45-1.125c.533.143.913.557 1.334 1.015C8.471 6.636 9.265 7.5 11 7.5c1.867 0 3.033-1 3.5-3c-.7 1-1.517 1.375-2.45 1.125c-.533-.143-.913-.557-1.334-1.015C10.029 3.864 9.235 3 7.5 3ZM4 7.5c-1.867 0-3.033 1-3.5 3c.7-1 1.517-1.375 2.45-1.125c.533.143.913.557 1.334 1.015C4.971 11.136 5.765 12 7.5 12c1.867 0 3.033-1 3.5-3c-.7 1-1.517 1.375-2.45 1.125c-.533-.143-.913-.557-1.334-1.015C6.529 8.364 5.735 7.5 4 7.5Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TailwindSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 2.5c-1.026 0-1.908.277-2.6.87c-.688.59-1.137 1.447-1.387 2.516a.5.5 0 0 0 .897.4c.316-.452.632-.723.936-.863c.294-.135.611-.162.975-.065c.366.098.65.386 1.095.87l.015.016c.336.365.745.81 1.305 1.156c.582.359 1.305.6 2.264.6c1.026 0 1.908-.277 2.6-.87c.688-.59 1.138-1.447 1.387-2.516a.5.5 0 0 0-.897-.4c-.316.452-.632.723-.936.863c-.294.135-.611.162-.975.065c-.366-.098-.65-.386-1.095-.87l-.015-.016c-.336-.365-.745-.81-1.305-1.156c-.582-.359-1.305-.6-2.264-.6M4 7c-1.026 0-1.908.277-2.6.87C.712 8.46.263 9.317.013 10.386a.5.5 0 0 0 .897.4c.316-.452.632-.723.936-.863c.294-.135.611-.162.975-.065c.366.098.65.386 1.095.87l.015.016c.336.365.745.81 1.305 1.156c.582.359 1.305.6 2.264.6c1.026 0 1.908-.277 2.6-.87c.688-.59 1.138-1.447 1.387-2.516a.5.5 0 0 0-.897-.4c-.316.452-.632.723-.936.863c-.294.135-.611.162-.975.065c-.366-.098-.65-.386-1.095-.87l-.015-.016c-.335-.365-.745-.81-1.305-1.156C5.682 7.24 4.959 7 4 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TargetOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M.5 7.5a7 7 0 1 0 14 0a7 7 0 0 0-14 0Z"/><path d="M3.5 7.5a4 4 0 1 0 8 0a4 4 0 0 0-8 0Z"/><path d="M6.5 7.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TargetSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7m0 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15M3 7.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelegramOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m14.5 1.5l-14 5l4 2l6-4l-4 5l6 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelegramSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.993 1.582a.5.5 0 0 0-.661-.553l-14 5a.5.5 0 0 0-.056.918l4 2a.5.5 0 0 0 .501-.031l3.32-2.214L6.11 9.188a.5.5 0 0 0 .113.728l6 4a.5.5 0 0 0 .77-.334z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m3.5 4.5l3 3l-3 3m4.5 0h4m-10.5-9h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h12A1.5 1.5 0 0 1 15 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5zm5.793 5L3.146 4.854l.708-.708L7.207 7.5l-3.353 3.354l-.708-.707zM12 11H8v-1h4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextDocumentAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 6.995H4v1h.5zm6 1h.5v-1h-.5zM4.5 10H4v1h.5zm6 1h.5v-1h-.5zm-6-7.003H4v1h.5zm6 1h.5v-1h-.5zm3-1.497h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-6 7.495h6v-1h-6zM4.5 11h6v-1h-6zm0-6.003h6v-1h-6zm8 9.003h-10v1h10zM2 13.5v-12H1v12zm11-10v10h1v-10zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextDocumentAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zm3 2.497h7v1H4zm7 2.998H4v1h7zM11 10H4v1h7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextDocumentOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 6.995H4v1h.5zm6 1h.5v-1h-.5zm-6 1.998H4v1h.5zm6 1.007h.5v-1h-.5zm-6-7.003H4v1h.5zM8.5 5H9V4h-.5zm2-4.5l.354-.354L10.707 0H10.5zm3 3h.5v-.207l-.146-.147zm-9 4.495h6v-1h-6zm0 2.998l6 .007v-1l-6-.007zm0-5.996L8.5 5V4l-4-.003zm8 9.003h-10v1h10zM2 13.5v-12H1v12zM2.5 1h8V0h-8zM13 3.5v10h1v-10zM10.146.854l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextDocumentSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zm3 2.497L9 4v1l-5-.003zm7 2.998H4v1h7zm0 3.006l-7-.008v1L11 11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 14V1.5M1.5 5V1.5h12V5m-10 8.5H11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v4h-1V2H8v11h3v1H3.5v-1H7V2H2v3H1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeHundredSixtyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 5.5l.4.3a.5.5 0 0 0-.4-.8zM3 7.5l-.4-.3A.5.5 0 0 0 3 8zm9.736-4.646l.382.323zM2 6h2.5V5H2zm2.1-.8l-1.5 2l.8.6l1.5-2zM3 8h.5V7H3zm.5 1H2v1h1.5zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 5 8.5zM3.5 8a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 3.5 7zM8 5h-.5v1H8zM6 6.5v1h1v-1zm0 1v1h1v-1zM7.5 7h-1v1h1zM9 8.5A1.5 1.5 0 0 0 7.5 7v1a.5.5 0 0 1 .5.5zM7.5 10A1.5 1.5 0 0 0 9 8.5H8a.5.5 0 0 1-.5.5zm0-1a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 10zm0-4A1.5 1.5 0 0 0 6 6.5h1a.5.5 0 0 1 .5-.5zM12 6.5v2h1v-2zm-1 2v-2h-1v2zm.5.5a.5.5 0 0 1-.5-.5h-1a1.5 1.5 0 0 0 1.5 1.5zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 13 8.5zM11.5 6a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 11.5 5zm0-1A1.5 1.5 0 0 0 10 6.5h1a.5.5 0 0 1 .5-.5zm-4 9A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zm6 3a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 2zm0 1a.5.5 0 0 1-.5-.5h-1A1.5 1.5 0 0 0 13.5 5zm-.5-.5c0-.123.044-.235.118-.323l-.763-.646c-.221.261-.355.6-.355.969zm.118-.323A.498.498 0 0 1 13.5 3V2c-.46 0-.871.207-1.145.531zM7.5 1c1.934 0 3.671.844 4.862 2.186l.748-.664A7.483 7.483 0 0 0 7.5 0zm5.854 3.67c.414.856.646 1.815.646 2.83h1c0-1.17-.268-2.277-.746-3.265zM14 3.5a.5.5 0 0 1-.348.477l.304.952A1.5 1.5 0 0 0 15 3.5zm-.348.477A.5.5 0 0 1 13.5 4v1a1.5 1.5 0 0 0 .456-.07z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeHundredSixtySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8.5V8h.5a.5.5 0 1 1-.5.5M11.5 6a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 1 0v-2a.5.5 0 0 0-.5-.5"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 0 1 12.787-5.32a1.5 1.5 0 0 1 1.659 2.484A7.52 7.52 0 0 1 15 7.5a7.5 7.5 0 0 1-15 0M13.5 3a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1M2 6h1.5l-.9 1.2A.5.5 0 0 0 3 8h.5a.5.5 0 0 1 0 1H2v1h1.5a1.5 1.5 0 0 0 .449-2.932L4.9 5.8a.5.5 0 0 0-.4-.8H2zm5.5-1A1.5 1.5 0 0 0 6 6.5v2A1.5 1.5 0 1 0 7.5 7H7v-.5a.5.5 0 0 1 .5-.5H8V5zM10 6.5a1.5 1.5 0 0 1 3 0v2a1.5 1.5 0 0 1-3 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbDownOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 9.5H3v.146l.078.122zm2.698 4.24l.421-.27zm2.667-1.51l-.448.223zM7.5 9.5V9a.5.5 0 0 0-.447.724zm7-5h.5v-.167l-.1-.133zm-2.4-3.2l.4-.3zM8.282 14.231l.257.429zM1 10V0H0v10zm2.078-.232l2.698 4.24l.843-.537l-2.697-4.24zm6.234 2.237L7.947 9.276l-.894.448l1.364 2.729zM7.5 10h5V9h-5zM15 7.5v-3h-1v3zm-.1-3.3L12.5 1l-.8.6l2.4 3.2zM10.5 0h-5v1h5zM3 2.5v7h1v-7zm9.5 7.5A2.5 2.5 0 0 0 15 7.5h-1A1.5 1.5 0 0 1 12.5 9zm-7-10A2.5 2.5 0 0 0 3 2.5h1A1.5 1.5 0 0 1 5.5 1zm3.039 14.66a2.034 2.034 0 0 0 .773-2.655l-.895.448c.242.483.07 1.071-.393 1.35zM12.5 1a2.5 2.5 0 0 0-2-1v1a1.5 1.5 0 0 1 1.2.6zM5.776 14.008a2.034 2.034 0 0 0 2.763.652l-.515-.858a1.034 1.034 0 0 1-1.405-.331z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbDownSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 10V0H0v10zM5.5 0A2.5 2.5 0 0 0 3 2.5v7.146l2.776 4.361a2.034 2.034 0 0 0 3.536-2.001L8.309 10H12.5A2.5 2.5 0 0 0 15 7.5V4.333L12.5 1a2.5 2.5 0 0 0-2-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbUpOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 5.5l-.422-.268L3 5.354V5.5zm2.698-4.24l.421.27zm2.667 1.51l-.448-.223zM7.5 5.5l-.447-.224A.5.5 0 0 0 7.5 6zm7 5l.4.3l.1-.133V10.5zm-2.4 3.2l.4.3zM8.282.769L8.539.34zM0 5v10h1V5zm3.922.768L6.619 1.53L5.776.992l-2.698 4.24zm4.495-3.22L7.053 5.275l.894.448l1.365-2.73zM7.5 6h5V5h-5zM14 7.5v3h1v-3zm.1 2.7l-2.4 3.2l.8.6l2.4-3.2zM10.5 14h-5v1h5zM4 12.5v-7H3v7zM12.5 6A1.5 1.5 0 0 1 14 7.5h1A2.5 2.5 0 0 0 12.5 5zm-7 8A1.5 1.5 0 0 1 4 12.5H3A2.5 2.5 0 0 0 5.5 15zM8.024 1.198c.464.278.635.866.393 1.35l.895.446A2.034 2.034 0 0 0 8.539.34zM11.7 13.4a1.5 1.5 0 0 1-1.2.6v1a2.5 2.5 0 0 0 2-1zM6.62 1.53c.3-.474.924-.62 1.404-.332L8.54.34a2.034 2.034 0 0 0-2.763.652z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbUpSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.312 2.995A2.034 2.034 0 0 0 5.776.992L3 5.354V12.5A2.5 2.5 0 0 0 5.5 15h5a2.5 2.5 0 0 0 2-1l2.5-3.333V7.5A2.5 2.5 0 0 0 12.5 5H8.309zM0 5v10h1V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbtackOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 14.5L5 10M.5 5.5l9 9m-1-14l6 6m-13 0l8-5m-1 12l5-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbtackSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.702 1.41L8.146.853l.708-.708l6 6l-.707.708l-.556-.556l-4.456 7.13l.719.718l-.708.708L5 10.707L.854 14.854l-.708-.707L4.293 10L.146 5.854l.708-.708l.718.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TickCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 7.5L7 10l4-5m-3.5 9.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TickCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m7.072 3.21l4.318-5.398l-.78-.624l-3.682 4.601L4.32 7.116l-.64.768z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TickOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m1 7l4.5 4.5L14 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TickSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 7.5L7 10l4-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TickSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.39 5.312l-4.318 5.399L3.68 7.884l.64-.768l2.608 2.173l3.682-4.601z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TickSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.707 3L5.5 12.207L.293 7L1 6.293l4.5 4.5l8.5-8.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TiktokOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 0v11A3.5 3.5 0 1 1 6 7.5m8-2A4.5 4.5 0 0 1 9.5 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TiktokSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 0h1v1a4 4 0 0 0 4 4v1a4.992 4.992 0 0 1-4-2v7a4 4 0 1 1-4-4v1a3 3 0 1 0 3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M3.5 2.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/><path d="M11.5.5h-8a3 3 0 0 0 0 6h8a3 3 0 1 0 0-6Zm0 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/><path d="M3.5 14.5h8a3 3 0 1 0 0-6h-8a3 3 0 0 0 0 6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 3.5A3.5 3.5 0 0 1 3.5 0h8a3.5 3.5 0 1 1 0 7h-8A3.5 3.5 0 0 1 0 3.5M3.5 2a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3M15 11.5a3.5 3.5 0 0 1-3.5 3.5h-8a3.5 3.5 0 1 1 0-7h8a3.5 3.5 0 0 1 3.5 3.5M11.5 13a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TopLeftOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 1.5V1a.5.5 0 0 0-.5.5zm0 .5H7V1H1.5zM1 1.5V7h1V1.5zm.146.354l12 12l.707-.708l-12-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TopLeftSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h6v1H2.707l11.147 11.146l-.708.708L2 2.707V7H1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TopRightOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 1.5h.5a.5.5 0 0 0-.5-.5zm0-.5H8v1h5.5zm-.5.5V7h1V1.5zm.146-.354l-12 12l.708.708l12-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TopRightSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1h6v6h-1V2.707L1.854 13.854l-.708-.708L12.293 2H8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendDownOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 3.5L5 8l3-3l6.5 6.5m0 0V5m0 6.5H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendDownSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m.146 3.854l.708-.708L5 7.293l3-3l6 6V5h1v7H8v-1h5.293L8 5.707l-3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendUpOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 11.5L5 7l3 3l6.5-6.5m0 0V10m0-6.5H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendUpSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3h7v7h-1V4.707l-6 6l-3-3l-4.146 4.147l-.708-.708L5 6.293l3 3L13.293 4H8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m7.5 1.5l-7 12h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.932 1.248a.5.5 0 0 0-.864 0l-7 12A.5.5 0 0 0 .5 14h14a.5.5 0 0 0 .432-.752z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 14.5h7m-3.5 0v-5m0 0a4 4 0 0 0 4-4v-4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v4a4 4 0 0 0 4 4Zm-4-7h-1a2 2 0 1 0 0 4h1m8-4h1a2 2 0 1 1 0 4h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophySolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 0A1.5 1.5 0 0 0 3 1.5V2h-.5a2.5 2.5 0 0 0 0 5h.756A4.504 4.504 0 0 0 7 9.973V14H4v1h7v-1H8V9.973A4.504 4.504 0 0 0 11.744 7h.756a2.5 2.5 0 0 0 0-5H12v-.5A1.5 1.5 0 0 0 10.5 0zM12 3v2.5c0 .169-.01.336-.027.5h.527a1.5 1.5 0 0 0 0-3zM2.5 3H3v2.5c0 .169.01.336.027.5H2.5a1.5 1.5 0 1 1 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 11.5v2m10-2v2M1 13.5h3m7 0h3M.5 2.5v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-12a1 1 0 0 0-1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 1A1.5 1.5 0 0 0 0 2.5v8A1.5 1.5 0 0 0 1.5 12H2v1H1v1h3v-1H3v-1h9v1h-1v1h3v-1h-1v-1h.5a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 13.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitchOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.5.5zm14 0h.5a.5.5 0 0 0-.5-.5zm0 8l.354.354A.5.5 0 0 0 15 8.5zm-3 3v.5a.5.5 0 0 0 .354-.146zm-5 0V11a.5.5 0 0 0-.325.12zm-3.5 3h-.5a.5.5 0 0 0 .825.38zm0-3h.5A.5.5 0 0 0 3 11zm-2.5 0H0a.5.5 0 0 0 .5.5zM.5 1h14V0H.5zM14 .5v8h1v-8zm.146 7.646l-3 3l.708.708l3-3zM11.5 11h-5v1h5zm-5.325.12l-3.5 3l.65.76l3.5-3zM3.5 14.5v-3h-1v3zM3 11H.5v1H3zm-2 .5V.5H0v11zM10 3v5h1V3zM7 3v5h1V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitchSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.5 0a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h2v2.5a.5.5 0 0 0 .825.38L6.685 12H11.5a.5.5 0 0 0 .354-.146l3-3A.5.5 0 0 0 15 8.5v-8a.5.5 0 0 0-.5-.5zM10 8V3h1v5zM7 3v5h1V3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.478 1.5l.5-.033a.5.5 0 0 0-.871-.301zm-.498 2.959a.5.5 0 1 0-1 0zm-6.49.082h-.5zm0 .959h.5zm-6.99 7V12a.5.5 0 0 0-.278.916zm.998-11l.469-.175a.5.5 0 0 0-.916-.048zm3.994 9l.354.353a.5.5 0 0 0-.195-.827zm7.224-8.027l-.37.336l.18.199l.265-.04zm1.264-.94c.051.778.003 1.25-.123 1.606c-.122.345-.336.629-.723 1l.692.722c.438-.42.776-.832.974-1.388c.193-.546.232-1.178.177-2.006zm0 3.654V4.46h-1v.728zm-6.99-.646V5.5h1v-.959zm0 .959V6h1v-.5zM10.525 1a3.539 3.539 0 0 0-3.537 3.541h1A2.539 2.539 0 0 1 10.526 2zm2.454 4.187C12.98 9.503 9.487 13 5.18 13v1c4.86 0 8.8-3.946 8.8-8.813zM1.03 1.675C1.574 3.127 3.614 6 7.49 6V5C4.174 5 2.421 2.54 1.966 1.325zm.021-.398C.004 3.373-.157 5.407.604 7.139c.759 1.727 2.392 3.055 4.73 3.835l.317-.948c-2.155-.72-3.518-1.892-4.132-3.29c-.612-1.393-.523-3.11.427-5.013zm4.087 8.87C4.536 10.75 2.726 12 .5 12v1c2.566 0 4.617-1.416 5.346-2.147zm7.949-8.009A3.445 3.445 0 0 0 10.526 1v1c.721 0 1.37.311 1.82.809zm-.296.83a3.513 3.513 0 0 0 2.06-1.134l-.744-.668a2.514 2.514 0 0 1-1.466.813zM.222 12.916C1.863 14.01 3.583 14 5.18 14v-1c-1.63 0-3.048-.011-4.402-.916z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.977 1.467a.5.5 0 0 0-.87-.301a2.559 2.559 0 0 1-1.226.763A3.441 3.441 0 0 0 10.526 1a3.539 3.539 0 0 0-3.537 3.541v.44C3.998 4.75 2.4 2.477 1.967 1.325a.5.5 0 0 0-.916-.048C.004 3.373-.157 5.407.604 7.139C1.27 8.656 2.61 9.864 4.51 10.665C3.647 11.276 2.194 12 .5 12a.5.5 0 0 0-.278.916C1.847 14 3.55 14 5.132 14h.048c4.861 0 8.8-3.946 8.8-8.812v-.479c.363-.37.646-.747.82-1.236c.193-.546.232-1.178.177-2.006"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypescriptOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 8v-.167c0-.736-.597-1.333-1.333-1.333H10a1.5 1.5 0 1 0 0 3h1a1.5 1.5 0 0 1 0 3h-1A1.5 1.5 0 0 1 8.5 11M8 6.5H3m2.5 0V13M.5.5h14v14H.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypescriptSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 0h15v15H0zm10 6a2 2 0 1 0 0 4h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1H8a2 2 0 0 0 2 2h1a2 2 0 1 0 0-4h-1a1 1 0 0 1 0-2h1.167c.46 0 .833.373.833.833V8h1v-.167C13 6.821 12.18 6 11.167 6zM3 6h5v1H6v6H5V7H3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnderlineOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 13.5h11M3.5 1v6.5a4 4 0 1 0 8 0V1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnderlineSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 7.5V1h1v6.5a3.5 3.5 0 1 0 7 0V1h1v6.5a4.5 4.5 0 0 1-9 0M13 13v1H2v-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.121 4.121l.354-.353zM5.5 7h4V6h-4zm4.5.5v3h1v-3zM9.5 11h-4v1h4zM5 10.5v-3H4v3zm.5.5a.5.5 0 0 1-.5-.5H4A1.5 1.5 0 0 0 5.5 12zm4.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM9.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 6zm-4-1A1.5 1.5 0 0 0 4 7.5h1a.5.5 0 0 1 .5-.5zm.5.5v-.879H5V6.5zm2.768-2.025l.378.379l.708-.708l-.38-.378zM7.62 4c.43 0 .843.17 1.147.475l.707-.707A2.621 2.621 0 0 0 7.62 3zM6 5.621C6 4.726 6.726 4 7.621 4V3A2.621 2.621 0 0 0 5 5.621zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0M7.621 4C6.726 4 6 4.726 6 5.621V6h3.5A1.5 1.5 0 0 1 11 7.5v3A1.5 1.5 0 0 1 9.5 12h-4A1.5 1.5 0 0 1 4 10.5v-3a1.5 1.5 0 0 1 1-1.415v-.464a2.621 2.621 0 0 1 4.475-1.853l.379.378l-.708.708l-.378-.38A1.621 1.621 0 0 0 7.62 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 6.5v-3a3 3 0 0 1 6 0V4m-8 2.5h10a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.121 4.121l.354-.353zM5.5 7h4V6h-4zm4.5.5v3h1v-3zM9.5 11h-4v1h4zM5 10.5v-3H4v3zm.5.5a.5.5 0 0 1-.5-.5H4A1.5 1.5 0 0 0 5.5 12zm4.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM9.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 6zm-4-1A1.5 1.5 0 0 0 4 7.5h1a.5.5 0 0 1 .5-.5zm.5.5v-.879H5V6.5zm2.768-2.025l.378.379l.708-.708l-.38-.378zM7.62 4c.43 0 .843.17 1.147.475l.707-.707A2.621 2.621 0 0 0 7.62 3zM6 5.621C6 4.726 6.726 4 7.621 4V3A2.621 2.621 0 0 0 5 5.621z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 5.621a1.621 1.621 0 0 1 2.768-1.146l.378.379l.708-.708l-.38-.378A2.621 2.621 0 0 0 5 5.62v.464A1.5 1.5 0 0 0 4 7.5v3A1.5 1.5 0 0 0 5.5 12h4a1.5 1.5 0 0 0 1.5-1.5v-3A1.5 1.5 0 0 0 9.5 6H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3.5a2.5 2.5 0 0 1 5 0V4h1v-.5a3.5 3.5 0 1 0-7 0V6H2.5A1.5 1.5 0 0 0 1 7.5v6A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-6A1.5 1.5 0 0 0 12.5 6H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.146 8.854l.354.353l.707-.707l-.353-.354zM7.5 5.5l.354-.354l-.354-.353l-.354.353zM4.146 8.146l-.353.354l.707.707l.354-.353zm6.708 0l-3-3l-.708.708l3 3zm-3.708-3l-3 3l.708.708l3-3zM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zm1 0A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 1 0 15a7.5 7.5 0 0 1 0-15M3.293 9L7.5 4.793L11.707 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m1 10l6.5-7l6.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpSmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m10.5 8.5l-3-3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpSmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 4.793L3.293 9h8.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 3l7.5 8H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m7.5 1.5l3.25 3m-3.25-3l-3 3m3-3V11m6-4v6.5h-12V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.486.807l3.603 3.326l-.678.734L8 2.642V11H7V2.707L4.854 4.854l-.708-.708zM2 13V7H1v7h13V7h-1v6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsbCableOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 5.5h6m-6 0a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1m-6 0v-5h6v5M6.5 2v2m2-2v2m-3 7.5v2h4v-2m-3 2V15m2-1.5V15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsbCableSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 0H4v6h7zM6 4V2h1v2zm2 0V2h1v2z" clip-rule="evenodd"/><path fill="currentColor" d="M12 7H3v3.5A1.5 1.5 0 0 0 4.5 12H5v2h1v1h1v-1h1v1h1v-1h1v-2h.5a1.5 1.5 0 0 0 1.5-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 13v.5h1V13zm8 0v.5h1V13zm-7 0v-.5H3v.5zm2.5-3h2V9h-2zm4.5 2.5v.5h1v-.5zM8.5 10a2.5 2.5 0 0 1 2.5 2.5h1A3.5 3.5 0 0 0 8.5 9zM4 12.5A2.5 2.5 0 0 1 6.5 10V9A3.5 3.5 0 0 0 3 12.5zM7.5 3A2.5 2.5 0 0 0 5 5.5h1A1.5 1.5 0 0 1 7.5 4zM10 5.5A2.5 2.5 0 0 0 7.5 3v1A1.5 1.5 0 0 1 9 5.5zM7.5 8A2.5 2.5 0 0 0 10 5.5H9A1.5 1.5 0 0 1 7.5 7zm0-1A1.5 1.5 0 0 1 6 5.5H5A2.5 2.5 0 0 0 7.5 8zm0 7A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15M1 7.5a6.5 6.5 0 1 1 10.988 4.702A3.5 3.5 0 0 0 8.5 9h-2a3.5 3.5 0 0 0-3.488 3.202A6.482 6.482 0 0 1 1 7.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 14.49v.5h.5v-.5zm-10 0H0v.5h.5zm7-4.996v.5zm-4 0v-.5zM8 3.498a2.499 2.499 0 0 1-2.5 2.498v1C7.433 6.996 9 5.43 9 3.498zM5.5 5.996A2.499 2.499 0 0 1 3 3.498H2a3.499 3.499 0 0 0 3.5 3.498zM3 3.498A2.499 2.499 0 0 1 5.5 1V0A3.499 3.499 0 0 0 2 3.498zM5.5 1A2.5 2.5 0 0 1 8 3.498h1A3.499 3.499 0 0 0 5.5 0zm5 12.99H.5v1h10zm-9.5.5v-1.995H0v1.995zm2.5-4.496h4v-1h-4zm6.5 2.5v1.996h1v-1.996zm-2.5-2.5a2.5 2.5 0 0 1 2.5 2.5h1a3.5 3.5 0 0 0-3.5-3.5zm-6.5 2.5a2.5 2.5 0 0 1 2.5-2.5v-1a3.5 3.5 0 0 0-3.5 3.5zM10 8h5V7h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 0a3.499 3.499 0 1 0 0 6.996A3.499 3.499 0 1 0 5.5 0M10 8h5V7h-5zm-2.5.994h-4a3.5 3.5 0 0 0-3.5 3.5v2.497h11v-2.497a3.5 3.5 0 0 0-3.5-3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M10.5 3.498a2.999 2.999 0 0 1-3 2.998a2.999 2.999 0 1 1 3-2.998Zm2 10.992h-10v-1.996a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlusOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 14.49v.5h.5v-.5zm-10 0H0v.5h.5zm7-4.996v.5zm-4 0v-.5zM8 3.498a2.499 2.499 0 0 1-2.5 2.498v1C7.433 6.996 9 5.43 9 3.498zM5.5 5.996A2.499 2.499 0 0 1 3 3.498H2a3.499 3.499 0 0 0 3.5 3.498zM3 3.498A2.499 2.499 0 0 1 5.5 1V0A3.499 3.499 0 0 0 2 3.498zM5.5 1A2.5 2.5 0 0 1 8 3.498h1A3.499 3.499 0 0 0 5.5 0zm5 12.99H.5v1h10zm-9.5.5v-1.995H0v1.995zm2.5-4.496h4v-1h-4zm6.5 2.5v1.996h1v-1.996zm-2.5-2.5a2.5 2.5 0 0 1 2.5 2.5h1a3.5 3.5 0 0 0-3.5-3.5zm-6.5 2.5a2.5 2.5 0 0 1 2.5-2.5v-1a3.5 3.5 0 0 0-3.5 3.5zM12 5v5h1V5zm-2 3h5V7h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlusSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 0a3.499 3.499 0 1 0 0 6.996A3.499 3.499 0 1 0 5.5 0M12 5v2h-2v1h2v2h1V8h2V7h-2V5zM7.5 8.994h-4a3.5 3.5 0 0 0-3.5 3.5v2.497h11v-2.497a3.5 3.5 0 0 0-3.5-3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0a3.499 3.499 0 1 0 0 6.996A3.499 3.499 0 1 0 7.5 0m-2 8.994a3.5 3.5 0 0 0-3.5 3.5v2.497h11v-2.497a3.5 3.5 0 0 0-3.5-3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSquareOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 14.5v.5h1v-.5zm8 0v.5h1v-.5zm-7 0v-2H3v2zM6.5 10h2V9h-2zm4.5 2.5v2h1v-2zM8.5 10a2.5 2.5 0 0 1 2.5 2.5h1A3.5 3.5 0 0 0 8.5 9zM4 12.5A2.5 2.5 0 0 1 6.5 10V9A3.5 3.5 0 0 0 3 12.5zM7.5 3A2.5 2.5 0 0 0 5 5.5h1A1.5 1.5 0 0 1 7.5 4zM10 5.5A2.5 2.5 0 0 0 7.5 3v1A1.5 1.5 0 0 1 9 5.5zM7.5 8A2.5 2.5 0 0 0 10 5.5H9A1.5 1.5 0 0 1 7.5 7zm0-1A1.5 1.5 0 0 1 6 5.5H5A2.5 2.5 0 0 0 7.5 8zm-6-6h12V0h-12zm12.5.5v12h1v-12zM13.5 14h-12v1h12zM1 13.5v-12H0v12zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSquareSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 0A1.5 1.5 0 0 0 0 1.5v12A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0zm5 9A3.5 3.5 0 0 0 3 12.5v1a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5v-1A3.5 3.5 0 0 0 8.5 9zM5 5.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsersOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 14.49v.5h.5v-.5zm-10 0H0v.5h.5zm14 .01v.5h.5v-.5zM8 3.498a2.499 2.499 0 0 1-2.5 2.498v1C7.433 6.996 9 5.43 9 3.498zM5.5 5.996A2.499 2.499 0 0 1 3 3.498H2a3.499 3.499 0 0 0 3.5 3.498zM3 3.498A2.499 2.499 0 0 1 5.5 1V0A3.499 3.499 0 0 0 2 3.498zM5.5 1A2.5 2.5 0 0 1 8 3.498h1A3.499 3.499 0 0 0 5.5 0zm5 12.99H.5v1h10zm-9.5.5v-1.995H0v1.995zm2.5-4.496h4v-1h-4zm6.5 2.5v1.996h1v-1.996zm-2.5-2.5a2.5 2.5 0 0 1 2.5 2.5h1a3.5 3.5 0 0 0-3.5-3.5zm-6.5 2.5a2.5 2.5 0 0 1 2.5-2.5v-1a3.5 3.5 0 0 0-3.5 3.5zM14 13v1.5h1V13zm.5 1H12v1h2.5zM12 11a2 2 0 0 1 2 2h1a3 3 0 0 0-3-3zm-.5-3A1.5 1.5 0 0 1 10 6.5H9A2.5 2.5 0 0 0 11.5 9zM13 6.5A1.5 1.5 0 0 1 11.5 8v1A2.5 2.5 0 0 0 14 6.5zM11.5 5A1.5 1.5 0 0 1 13 6.5h1A2.5 2.5 0 0 0 11.5 4zm0-1A2.5 2.5 0 0 0 9 6.5h1A1.5 1.5 0 0 1 11.5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsersSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 0a3.499 3.499 0 1 0 0 6.996A3.499 3.499 0 1 0 5.5 0m-2 8.994a3.5 3.5 0 0 0-3.5 3.5v2.497h11v-2.497a3.5 3.5 0 0 0-3.5-3.5zm9 1.006H12v5h3v-2.5a2.5 2.5 0 0 0-2.5-2.5"/><path fill="currentColor" d="M11.5 4a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorDocumentOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-7 4V4H3v.5zm2 0H6V4h-.5zm0 2V7H6v-.5zm-2 0H3V7h.5zm8-2h.5V4h-.5zm0 2V7h.5v-.5zm-2 0H9V7h.5zm0-2V4H9v.5zm-6 6V10H3v.5zm2 0H6V10h-.5zm0 2v.5H6v-.5zm-2 0H3v.5h.5zm6-2V10H9v.5zm0 2H9v.5h.5zm2 0v.5h.5v-.5zm0-2h.5V10h-.5zm1 3.5h-10v1h10zM2 13.5v-12H1v12zm11-10v10h1v-10zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM3.5 5h2V4h-2zM5 4.5v2h1v-2zM5.5 6h-2v1h2zM4 6.5v-2H3v2zm7-2v2h1v-2zm.5 1.5h-2v1h2zm-1.5.5v-2H9v2zM9.5 5h2V4h-2zm-6 6h2v-1h-2zm1.5-.5v2h1v-2zm.5 1.5h-2v1h2zm-1.5.5v-2H3v2zm5-2v2h1v-2zm.5 2.5h2v-1h-2zm2.5-.5v-2h-1v2zm-.5-2.5h-2v1h2zm-6-4h4V5h-4zM4 6.5v4h1v-4zm6 0v4h1v-4zM5.5 12h4v-1h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorDocumentSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 10v1H6v-1H5V7h1V6h3v1h1v3zM4 5v1h1V5zm6 0v1h1V5zm-6 7v-1h1v1zm6-1v1h1v-1z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zM3 4h3v1h3V4h3v3h-1v3h1v3H9v-1H6v1H3v-3h1V7H3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VennDiagramOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M2.5 9.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/><path d="M.5 5.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/><path d="M4.5 5.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VennDiagramSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 15a5.502 5.502 0 0 1-5.255-3.873A6.47 6.47 0 0 0 5.5 12c.698 0 1.37-.11 2-.313a6.51 6.51 0 0 0 2 .313a6.47 6.47 0 0 0 3.255-.873A5.503 5.503 0 0 1 7.5 15m6.455-6.273A5.5 5.5 0 0 0 9 .023a6.5 6.5 0 0 1 2.96 4.747a6.484 6.484 0 0 1 1.994 3.956M6 .022a5.5 5.5 0 0 0-4.954 8.704A6.484 6.484 0 0 1 3.04 4.772A6.5 6.5 0 0 1 6 .022M2.005 9.747A5.477 5.477 0 0 0 6 10.977a6.5 6.5 0 0 1-2.953-4.704a5.475 5.475 0 0 0-1.04 3.474M13 9.5c0-1.205-.388-2.32-1.046-3.227a6.5 6.5 0 0 1-2.953 4.705a5.477 5.477 0 0 0 3.994-1.23c.003-.083.005-.165.005-.248M7.5.375a5.515 5.515 0 0 1 3.255 3.498A6.47 6.47 0 0 0 7.5 3a6.47 6.47 0 0 0-3.255.873A5.515 5.515 0 0 1 7.5.375M7.5 4c-1.327 0-2.544.47-3.495 1.253A5.502 5.502 0 0 0 7.5 10.625a5.502 5.502 0 0 0 3.495-5.372A5.477 5.477 0 0 0 7.5 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumnOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 0v15m4-15v15m-8-15v15m-4-15v15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumnSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 15V0h1v15zm4 0V0h1v15zm4 0V0h1v15zm4 0V0h1v15z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewGridOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 5.5h15m-15-4h15m-15 8h15m-15 4h15M9.5 0v15m4-15v15m-8-15v15m-4-15v15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewGridSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1V0h1v1h3V0h1v1h3V0h1v1h3V0h1v1h1v1h-1v3h1v1h-1v3h1v1h-1v3h1v1h-1v1h-1v-1h-3v1H9v-1H6v1H5v-1H2v1H1v-1H0v-1h1v-3H0V9h1V6H0V5h1V2H0V1zm1 1v3h3V2zm4 0v3h3V2zm4 0v3h3V2zm3 4h-3v3h3zm0 4h-3v3h3zm-4 3v-3H6v3zm-4 0v-3H2v3zM2 9h3V6H2zm4-3v3h3V6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VimOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 3.5h1v10h3l8-10v-2h-5v2h2l-5 6v-6h1v-2h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VimSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1H1v3h1v10h3.74L14 3.675V1H8v3h1.432L6 8.119V4h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOneOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5zm0-5.996v.5h.138l.12-.071zm5-2.998H9a.5.5 0 0 0-.757-.429zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492zm-5-3.498h-2v1h2zm-2 0a.5.5 0 0 1-.5-.5H0c0 .83.672 1.5 1.5 1.5zm-.5-.5V5.498H0v3.998zm0-3.997a.5.5 0 0 1 .5-.499v-1a1.5 1.5 0 0 0-1.5 1.5zm.5-.499h2v-1h-2zm2.257-.071l5-2.998l-.514-.858l-5 2.998zM8 1.5v11.992h1V1.5zm.757 11.563l-5-2.998l-.514.858l5 2.998zM10 6v3h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOneSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5a1.5 1.5 0 0 0-1.5 1.5v3.997c0 .83.672 1.5 1.5 1.5h1.862l4.88 2.926A.5.5 0 0 0 9 13.492zM10 6v3h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeThreeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5zm0-5.996v.5h.138l.12-.071zm5-2.998H9a.5.5 0 0 0-.757-.429zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492zm-5-3.498h-2v1h2zm-2 0a.5.5 0 0 1-.5-.5H0c0 .83.672 1.5 1.5 1.5zm-.5-.5V5.498H0v3.998zm0-3.997a.5.5 0 0 1 .5-.499v-1a1.5 1.5 0 0 0-1.5 1.5zm.5-.499h2v-1h-2zm2.257-.071l5-2.998l-.514-.858l-5 2.998zM8 1.5v11.992h1V1.5zm.757 11.563l-5-2.998l-.514.858l5 2.998zM10 6v3h1V6zm2-2v7h1V4zm2-2v11h1V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeThreeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5a1.5 1.5 0 0 0-1.5 1.5v3.997c0 .83.672 1.5 1.5 1.5h1.862l4.88 2.926A.5.5 0 0 0 9 13.492zm5 .5v11h1V2zm-2 2v7h1V4zm-2 2v3h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeTwoOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5zm0-5.996v.5h.138l.12-.071zm5-2.998H9a.5.5 0 0 0-.757-.429zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492zm-5-3.498h-2v1h2zm-2 0a.5.5 0 0 1-.5-.5H0c0 .83.672 1.5 1.5 1.5zm-.5-.5V5.498H0v3.998zm0-3.997a.5.5 0 0 1 .5-.499v-1a1.5 1.5 0 0 0-1.5 1.5zm.5-.499h2v-1h-2zm2.257-.071l5-2.998l-.514-.858l-5 2.998zM8 1.5v11.992h1V1.5zm.757 11.563l-5-2.998l-.514.858l5 2.998zM10 6v3h1V6zm2-2v7h1V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeTwoSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1.5a.5.5 0 0 0-.757-.429L3.362 3.998H1.5a1.5 1.5 0 0 0-1.5 1.5v3.997c0 .83.672 1.5 1.5 1.5h1.862l4.88 2.926A.5.5 0 0 0 9 13.492zM12 4v7h1V4zm-2 2v3h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VrHeadsetOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.851 5.4l.137.48zm13.298 0l.137-.481zM4.58 12.352l.44.24zm5.84 0l.438-.239zM2.996 3.757l-.464-.185zm-.961 1.057a.5.5 0 0 0 .928.372zm9.967-1.057l.464-.185zm.033 1.429a.5.5 0 1 0 .928-.372zm1.964.68V9.21h1V5.865zM11.21 12h-.542v1h.542zm-6.878 0H3.79v1h.542zM1 9.21V5.865H0V9.21zM.988 5.88a23.703 23.703 0 0 1 13.024 0l.274-.961a24.703 24.703 0 0 0-13.572 0zM3.79 12A2.79 2.79 0 0 1 1 9.21H0A3.79 3.79 0 0 0 3.79 13zm.352.113a.217.217 0 0 1 .19-.113v1a.78.78 0 0 0 .687-.408zm.877.479c1.071-1.963 3.89-1.963 4.962 0l.877-.479c-1.45-2.658-5.267-2.658-6.716 0zM10.668 12c.08 0 .152.043.19.113l-.877.479a.78.78 0 0 0 .687.408zM14 9.21A2.79 2.79 0 0 1 11.21 12v1A3.79 3.79 0 0 0 15 9.21zm1-3.345a.984.984 0 0 0-.714-.946l-.274.961A.016.016 0 0 1 14 5.865zm-14 0a.016.016 0 0 1-.012.015l-.274-.96A.984.984 0 0 0 0 5.865zm1.533-2.293l-.497 1.242l.928.372l.497-1.243zm9.006.37l.497 1.244l.928-.372l-.497-1.242zM4.854 3h5.292V2H4.854zm7.613.572A2.5 2.5 0 0 0 10.146 2v1a1.5 1.5 0 0 1 1.393.943zm-9.006.37A1.5 1.5 0 0 1 4.854 3V2a2.5 2.5 0 0 0-2.321 1.572z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VrHeadsetSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12.467 3.572l.394.985c.478.106.953.227 1.425.362c.423.12.714.507.714.946V9.21A3.79 3.79 0 0 1 11.21 13h-.542a.783.783 0 0 1-.687-.408c-1.071-1.963-3.89-1.963-4.962 0a.783.783 0 0 1-.687.408H3.79A3.79 3.79 0 0 1 0 9.21V5.865c0-.44.291-.825.714-.946c.472-.135.947-.256 1.425-.362l.394-.985A2.5 2.5 0 0 1 4.854 2h5.292a2.5 2.5 0 0 1 2.321 1.572m-9.006.37A1.5 1.5 0 0 1 4.854 3h5.292a1.5 1.5 0 0 1 1.393.943l.153.384a24.703 24.703 0 0 0-8.384 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VueOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 13.5l-.432.252a.5.5 0 0 0 .864 0zm4-12l.434.248A.5.5 0 0 0 11.5 1zm-4 7l-.434.248a.5.5 0 0 0 .868 0zm-4-7V1a.5.5 0 0 0-.434.748zm3 0l.447-.224L6.81 1H6.5zm1 2l-.447.224a.5.5 0 0 0 .894 0zm1-2V1h-.309l-.138.276zm-8.432.252l7 12l.864-.504l-7-12zm7.864 12l7-12l-.864-.504l-7 12zm3.134-12.5l-4 7l.868.496l4-7zm-3.132 7l-4-7l-.868.496l4 7zM3.5 2h3V1h-3zm2.553-.276l1 2l.894-.448l-1-2zm1.894 2l1-2l-.894-.448l-1 2zM8.5 2h3V1h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VueSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.717 1H.5a.5.5 0 0 0-.432.752l7 12a.5.5 0 0 0 .864 0l7-12A.5.5 0 0 0 14.5 1h-2.217L7.5 8.972z"/><path fill="currentColor" d="M11.117 1H8.19L7.5 2.382L6.809 1H3.883L7.5 7.028z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletAltOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.5 6V2.5a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V9m.93-3.5H9.5a2 2 0 1 0 0 4h4.93a.07.07 0 0 0 .07-.07V5.57a.07.07 0 0 0-.07-.07Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletAltSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 12.5V10h.43a.57.57 0 0 0 .57-.57V5.57a.57.57 0 0 0-.57-.57H14V2.5A1.5 1.5 0 0 0 12.5 1h-11A1.5 1.5 0 0 0 0 2.5v10A1.5 1.5 0 0 0 1.5 14h11a1.5 1.5 0 0 0 1.5-1.5M14 6v3H9.5a1.5 1.5 0 1 1 0-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 3.5v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1H3m-2.5 0v-1a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v1H3m-2.5 0H3m6 6h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h8A1.5 1.5 0 0 1 11 2.5V3h2.5A1.5 1.5 0 0 1 15 4.5v8a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5zM9 10h3V9H9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WanOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 14.5l-.489.105a.5.5 0 0 0 .967.042zm4-13l.478-.147a.5.5 0 0 0-.956 0zm4 13l-.478.147a.5.5 0 0 0 .967-.042zM.011.605l3 14l.978-.21l-3-14zm3.967 14.042l4-13l-.956-.294l-4 13zm3.044-13l4 13l.956-.294l-4-13zm4.967 12.958l3-14l-.978-.21l-3 14zM0 6h15V5H0zm0 4h15V9H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WanSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.953 5L.01.605l.98-.21L1.976 5H5.9l1.122-3.647a.5.5 0 0 1 .956 0L9.1 5h3.924l.987-4.605l.978.21L14.047 5H15v1h-1.167l-.643 3H15v1h-2.024l-.987 4.605a.5.5 0 0 1-.967.042L9.592 10H5.408l-1.43 4.647a.5.5 0 0 1-.967-.042L2.024 10H0V9h1.81l-.643-3H0V5zM2.19 6l.643 3h1.836l.923-3zm4.449 0l-.924 3h3.57L8.36 6zm1.415-1L7.5 3.2L6.946 5zm1.354 1l.923 3h1.836l.643-3zm2.545 4h-1.314l.774 2.518zM4.36 10H3.047l.54 2.518z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WandOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.853 8.848l.354-.353l-.707-.707l-.353.353zM9 10.493v.5h1v-.5zm1-.999v-.5H9v.5zM9 1.5V2h1v-.5zm1-1V0H9v.5zM4.5 4.997H4v1h.5zm1 1H6v-1h-.5zm8-1H13v1h.5zm1 1h.5v-1h-.5zm-2.353-3.852l-.354.354l.707.707l.353-.354zm1.706-.292l.354-.353L13.5.792l-.353.354zm-8-.707L5.5.792l-.707.708l.354.353zm.294 1.706l.353.354l.707-.707l-.354-.354zm6.706 5.29l-.353-.354l-.707.707l.354.353zm.294 1.706l.353.353l.707-.707l-.354-.354zM.853 14.844l6-5.996l-.706-.707l-6 5.996zM10 10.495v-1H9v1zM10 1.5v-1H9v1zM4.5 5.997h1v-1h-1zm9 0h1v-1h-1zm-.647-3.145l1-.999l-.706-.707l-1 1zm-7.706-.999l1 1l.706-.708l-1-1zm7 6.995l1 1l.706-.708l-1-.999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WandSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 0v2H9V0zM5.5.792L7.207 2.5l-.707.707L4.793 1.5zm8.707.708L12.5 3.206l-.707-.707L13.5.792zM4 4.997h2v1H4zm9 0h2v1h-2zM7.207 8.495l-6.354 6.35l-.706-.708L6.5 7.787zm5.293-.707l1.707 1.706l-.707.707l-1.707-1.706zM10 8.994v2H9v-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 3.5h6m-6 0a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m0-8v-2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2m0 0a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1m0 0h-6m6 0v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2m9-5.5v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.085V1.5A1.5 1.5 0 0 1 5.5 0h4A1.5 1.5 0 0 1 11 1.5v1.585A1.5 1.5 0 0 1 12 4.5v6a1.5 1.5 0 0 1-1 1.415V13.5A1.5 1.5 0 0 1 9.5 15h-4A1.5 1.5 0 0 1 4 13.5v-1.585A1.5 1.5 0 0 1 3 10.5v-6a1.5 1.5 0 0 1 1-1.415M5 1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3H5zM5 12h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5z" clip-rule="evenodd"/><path fill="currentColor" d="M13 6v3h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebpackOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 10.5v-6m0 6l6 3.5m-6-3.5L5 9M1.5 4.5l6-3.5m-6 3.5l6 3m0-6.5l6 3.5M7.5 1v3.5m6 0v6m0-6l-6 3m6 3l-6 3.5m6-3.5L10 9m-2.5 5V7.5m0-3L5 6v3m2.5-4.5L10 6v3M5 9l2.5 1.5L10 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebpackSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 4.217l2.25 1.35l3.268-1.635L8 .712zM7 .713L1.482 3.932L4.75 5.566L7 4.216zM1 4.809v5.422l3.5-1.556V6.56zm.523 6.283L7 14.287v-3.504l-2.034-1.22zM8 14.287l5.477-3.195l-3.443-1.53L8 10.783zm6-4.057V4.81l-3.5 1.75v2.116zm-6-.613l1.5-.9V7.059l-1.5.75zM7 7.809v1.808l-1.5-.9V7.059zM5.811 6.096l1.689.845l1.689-.845L7.5 5.083z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhatsappOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.9 11.7l.447.224l.138-.277L2.3 11.4zm1.4 1.4l.3-.4l-.247-.185l-.277.138zM.5 14.5l-.447-.224a.5.5 0 0 0 .67.671zm4-10l-.277-.416A.5.5 0 0 0 4 4.5zm6 6v.5a.5.5 0 0 0 .416-.223zM6.254 5.026l.493-.083zm.14.836l-.493.082zm-.432.997l.277.416zm4.68 3.428l.416.277zm-.668-1.541l.083-.493zm-.836-.14l-.082.493zm-.997.432l-.416-.277zM0 7.5c0 1.688.558 3.247 1.5 4.5l.8-.6A6.47 6.47 0 0 1 1 7.5zM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM15 7.5A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5zM7.5 15A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14zM3 13.5A7.47 7.47 0 0 0 7.5 15v-1a6.469 6.469 0 0 1-3.9-1.3zM.723 14.947l2.8-1.4l-.448-.894l-2.8 1.4zm.729-3.47l-1.4 2.8l.894.447l1.4-2.8zM4 4.5v1h1v-1zM9.5 11h1v-1h-1zM4 5.5A5.5 5.5 0 0 0 9.5 11v-1A4.5 4.5 0 0 1 5 5.5zm.777-.584l.214-.142l-.555-.832l-.213.142zm.984.192l.14.836l.986-.164l-.14-.837zm-.076 1.335l-.962.641l.554.832l.962-.641zm.216-.499a.5.5 0 0 1-.216.499l.554.832a1.5 1.5 0 0 0 .648-1.495zm-.91-1.17a.5.5 0 0 1 .77.334l.986-.165a1.5 1.5 0 0 0-2.311-1.001zm5.925 6.003l.142-.213l-.832-.555l-.142.214zm-.86-2.524l-.836-.14l-.164.987l.836.139zm-2.33.508l-.642.962l.832.554l.641-.962zm1.494-.648a1.5 1.5 0 0 0-1.495.648l.832.554a.5.5 0 0 1 .499-.216zm1.838 2.451a1.5 1.5 0 0 0-1.001-2.311l-.165.986a.5.5 0 0 1 .334.77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhatsappSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 4.768a.5.5 0 0 1 .761.34l.14.836a.5.5 0 0 1-.216.499l-.501.334A4.501 4.501 0 0 1 5 5.5zM9.5 10a4.5 4.5 0 0 1-1.277-.184l.334-.5a.5.5 0 0 1 .499-.217l.836.14a.5.5 0 0 1 .34.761z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 3.253 6.182l-2.53 1.265a.5.5 0 0 1-.67-.67l1.265-2.53A7.467 7.467 0 0 1 0 7.5m4.23-3.42l.206-.138a1.5 1.5 0 0 1 2.311 1.001l.14.837a1.5 1.5 0 0 1-.648 1.495l-.658.438A4.522 4.522 0 0 0 7.286 9.42l.44-.658a1.5 1.5 0 0 1 1.494-.648l.837.14a1.5 1.5 0 0 1 1.001 2.311l-.138.207a.5.5 0 0 1-.42.229h-1A5.5 5.5 0 0 1 4 5.5v-1a.5.5 0 0 1 .23-.42" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiFullOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.219 9.318c1.155-1.4 2.698-2.161 4.281-2.161v-1c-1.917 0-3.732.924-5.052 2.525zM7.5 7.157c1.583 0 3.126.762 4.281 2.161l.771-.636C11.232 7.08 9.417 6.157 7.5 6.157zM.886 6.318C2.659 4.168 5.042 2.985 7.5 2.985v-1c-2.793 0-5.446 1.346-7.386 3.697zM7.5 2.985c2.458 0 4.84 1.183 6.614 3.333l.772-.636C12.946 3.33 10.293 1.985 7.5 1.985zM7.5 12a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 13zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 11.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 10zm0-1A1.5 1.5 0 0 0 6 11.5h1a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiFullSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 2.985c-2.458 0-4.84 1.183-6.614 3.333l-.772-.636C2.054 3.33 4.707 1.985 7.5 1.985s5.446 1.346 7.386 3.697l-.772.636C12.341 4.168 9.958 2.985 7.5 2.985"/><path fill="currentColor" d="M7.5 7.157c-1.583 0-3.126.762-4.28 2.161l-.772-.636C3.768 7.08 5.583 6.157 7.5 6.157c1.918 0 3.732.924 5.053 2.525l-.772.636c-1.155-1.4-2.698-2.161-4.28-2.161M6 11.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiLowOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.22 9.318c1.154-1.4 2.697-2.161 4.28-2.161v-1c-1.917 0-3.732.924-5.052 2.525zM7.5 7.157c1.583 0 3.126.762 4.281 2.161l.771-.636C11.232 7.08 9.417 6.157 7.5 6.157zM7.5 12a.5.5 0 0 1-.5-.5H6A1.5 1.5 0 0 0 7.5 13zm.5-.5a.5.5 0 0 1-.5.5v1A1.5 1.5 0 0 0 9 11.5zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 7.5 10zm0-1A1.5 1.5 0 0 0 6 11.5h1a.5.5 0 0 1 .5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiLowSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.219 9.318c1.155-1.4 2.698-2.161 4.28-2.161c1.584 0 3.127.762 4.282 2.161l.771-.636C11.232 7.08 9.417 6.157 7.5 6.157c-1.918 0-3.732.924-5.052 2.525zM6 11.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiNoneOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 11.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiNoneSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 10a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.5 3.5l-.105-.489A.5.5 0 0 0 0 3.5zm14-3h.5a.5.5 0 0 0-.605-.489zm0 14l-.07.495A.5.5 0 0 0 15 14.5zm-14-2H0a.5.5 0 0 0 .43.495zm.105-8.511l14-3l-.21-.978l-14 3zM14 .5v14h1V.5zm.57 13.505l-14-2l-.14.99l14 2zM1 12.5v-9H0v9zM.5 8h14V7H.5zM6 2v11h1V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.814.111A.5.5 0 0 1 15 .5V7H7V1.596L14.395.01a.5.5 0 0 1 .42.1M6 1.81L.395 3.011A.5.5 0 0 0 0 3.5V7h6zM0 8v4.5a.5.5 0 0 0 .43.495l5.57.796V8zm7 5.934l7.43 1.061A.5.5 0 0 0 15 14.5V8H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WordpressOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.5 13.5l-.479.144a.5.5 0 0 0 .957.003zm4 0l-.479.144a.5.5 0 0 0 .943.042zm3.53-8.827l.465.186zm-.54-1.931l-.3.4zm-.133-.1l.3-.4zm-2 2l.4-.3zm.286-2l-.3-.4zM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15zM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5zM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1zM2.021 3.644l3 10l.958-.288l-3-10zm3.957 10.003l2-6.5l-.956-.294l-2 6.5zM6.02 3.644l3 10l.958-.288l-3-10zm6.768-1.302l-.132-.1l-.6.8l.132.1zm-2.832 2.6l.643.858l.8-.6l-.643-.857zm.386-2.7a1.929 1.929 0 0 0-.386 2.7l.8-.6a.929.929 0 0 1 .186-1.3zm2.314 0a1.928 1.928 0 0 0-2.314 0l.6.8a.928.928 0 0 1 1.114 0zm.838 2.617a2.149 2.149 0 0 0-.706-2.517l-.6.8c.416.312.57.863.377 1.345zM2 4h2V3H2zm3 0h3V3H5zm4.964 9.686l3.531-8.827l-.928-.372l-3.531 8.827z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WordpressSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 3a7.5 7.5 0 1 1-.634 1h1.262l2.893 9.644a.5.5 0 0 0 .957.003l1.541-5.01l1.502 5.007a.5.5 0 0 0 .943.042l3.53-8.827a2.149 2.149 0 0 0-.705-2.517l-.132-.1a1.929 1.929 0 0 0-2.7 2.7l.643.858l.8-.6l-.643-.857a.929.929 0 0 1 1.3-1.3l.132.099c.416.312.57.863.377 1.345l-2.999 7.498L7.172 4H8V3H5v1h1.128l.875 2.916l-1.497 4.864L3.172 4H4V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XcircleOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m4.5 4.5l6 6m-6 0l6-6m-3 10a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XcircleSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0m10.147 3.354L7.5 8.207l-2.646 2.647l-.708-.707L6.793 7.5L4.146 4.854l.708-.708L7.5 6.793l2.646-2.647l.708.708L8.207 7.5l2.647 2.646z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xoutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m1.5 1.5l12 12m-12 0l12-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XsmallOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m4.5 4.5l6 6m-6 0l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XsmallSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.793 7.5L4.146 4.854l.708-.708L7.5 6.793l2.646-2.647l.708.708L8.207 7.5l2.647 2.646l-.708.707L7.5 8.208l-2.646 2.646l-.708-.707z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xsolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.793 7.5L1.146 1.854l.708-.708L7.5 6.793l5.647-5.647l.707.708L8.207 7.5l5.647 5.646l-.707.707L7.5 8.208l-5.646 5.647l-.708-.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XlsOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147zm-3-3l.354-.354L10.707 0H10.5zm-4 10H6v.5h.5zm-2-1H5v-.207l-.146-.147zm-2-2H2v.207l.146.147zm8-1V6H10v.5zm0 2H10V9h.5zm2 0h.5V8h-.5zm0 2v.5h.5v-.5zm-10-1l-.354-.354L2 9.293V9.5zm2-2l.354.354L5 7.707V7.5zM2 5V1.5H1V5zm11-1.5V5h1V3.5zM2.5 1h8V0h-8zm7.646-.146l3 3l.708-.708l-3-3zM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5zM1 12v1.5h1V12zm1.5 3h10v-1h-10zM14 13.5V12h-1v1.5zM12.5 15a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5zM1 13.5A1.5 1.5 0 0 0 2.5 15v-1a.5.5 0 0 1-.5-.5zM6 6v4.5h1V6zm.5 5H9v-1H6.5zM4 9.5V11h1V9.5zm.854-.354l-2-2l-.708.708l2 2zM3 7.5V6H2v1.5zM13 6h-2.5v1H13zm-3 .5v2h1v-2zm.5 2.5h2V8h-2zm1.5-.5v2h1v-2zm.5 1.5H10v1h2.5zM3 11V9.5H2V11zm-.146-1.146l2-2l-.708-.708l-2 2zM5 7.5V6H4v1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XlsSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5zm2 5.793V6H2v1.707l.793.793L2 9.293V11h1V9.707l.5-.5l.5.5V11h1V9.293L4.207 8.5L5 7.707V6H4v1.293l-.5.5zM6 6h1v4h2v1H6zm7 0h-3v3h2v1h-2v1h3V8h-2V7h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YenOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 15V7.5m0 0l-5-7m5 7l5-7M3 7.5h9m-9 4h9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YenSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.528 7H3v1h4v3H3v1h4v3h1v-3h4v-1H8V8h4V7H8.472L12.907.79l-.814-.58L7.5 6.64L2.907.21l-.814.58z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.61 12.738l-.104.489zm11.78 0l.104.489zm0-10.476l.104-.489zm-11.78 0l.106.489zM6.5 5.5l.277-.416A.5.5 0 0 0 6 5.5zm0 4H6a.5.5 0 0 0 .777.416zm3-2l.277.416a.5.5 0 0 0 0-.832zM0 3.636v7.728h1V3.636zm15 7.728V3.636h-1v7.728zM1.506 13.227c3.951.847 8.037.847 11.988 0l-.21-.978a27.605 27.605 0 0 1-11.568 0zM13.494 1.773a28.606 28.606 0 0 0-11.988 0l.21.978a27.607 27.607 0 0 1 11.568 0zM15 3.636c0-.898-.628-1.675-1.506-1.863l-.21.978c.418.09.716.458.716.885zm-1 7.728a.905.905 0 0 1-.716.885l.21.978A1.905 1.905 0 0 0 15 11.364zm-14 0c0 .898.628 1.675 1.506 1.863l.21-.978A.905.905 0 0 1 1 11.364zm1-7.728c0-.427.298-.796.716-.885l-.21-.978A1.905 1.905 0 0 0 0 3.636zM6 5.5v4h1v-4zm.777 4.416l3-2l-.554-.832l-3 2zm3-2.832l-3-2l-.554.832l3 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.599 7.5L7 8.566V6.434z"/><path fill="currentColor" fill-rule="evenodd" d="M1.506 1.773a28.606 28.606 0 0 1 11.988 0A1.905 1.905 0 0 1 15 3.636v7.728c0 .898-.628 1.675-1.506 1.863a28.605 28.605 0 0 1-11.988 0A1.905 1.905 0 0 1 0 11.364V3.636c0-.898.628-1.675 1.506-1.863m5.271 3.311A.5.5 0 0 0 6 5.5v4a.5.5 0 0 0 .777.416l3-2a.5.5 0 0 0 0-.832z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZipOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 8.5V8h-.39l-.095.379zm2 0l.485-.121L5.89 8H5.5zm.69 2.758l.484-.122zm-3.38 0l.486.12zM9.5 10.5l.277-.416A.5.5 0 0 0 9 10.5zm3 2l-.277.416A.5.5 0 0 0 13 12.5zm-3-8H9V5h.5zM1.5 1h12V0h-12zm12.5.5v12h1v-12zM13.5 14h-12v1h12zM1 13.5v-12H0v12zm.5.5a.5.5 0 0 1-.5-.5H0A1.5 1.5 0 0 0 1.5 15zm12.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5zM13.5 1a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 13.5 0zm-12-1A1.5 1.5 0 0 0 0 1.5h1a.5.5 0 0 1 .5-.5zM3 3h3V2H3zm0 2h3V4H3zm0 2h3V6H3zm.5 2h2V8h-2zm1.515-.379l.69 2.758l.97-.243l-.69-2.757zM5.219 12H3.781v1h1.438zm-1.923-.621l.69-2.758l-.971-.242l-.69 2.757zM3.78 12a.5.5 0 0 1-.485-.621l-.97-.243A1.5 1.5 0 0 0 3.78 13zm1.923-.621A.5.5 0 0 1 5.22 12v1a1.5 1.5 0 0 0 1.455-1.864zM10 13v-2.5H9V13zm-.777-2.084l3 2l.554-.832l-3-2zM13 12.5V10h-1v2.5zM9 6v3h1V6zm3 0v3h1V6zM9.5 8h3V7h-3zm.5-3.5v-1H9v1zm3-.5h-1.5v1H13zm-1.5 0h-2v1h2zm-.5-.5v1h1v-1zm-.5-.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 10.5 2zm-.5.5a.5.5 0 0 1 .5-.5V2A1.5 1.5 0 0 0 9 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZipSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.296 11.379L3.89 9h1.22l.594 2.379A.5.5 0 0 1 5.22 12H3.781a.5.5 0 0 1-.485-.621M10.5 3a.5.5 0 0 0-.5.5V4h1v-.5a.5.5 0 0 0-.5-.5"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 13.5zM3 3h3V2H3zm0 2h3V4H3zm3 2H3V6h3zm-.11 1H3.11l-.784 3.136A1.5 1.5 0 0 0 3.78 13h1.438a1.5 1.5 0 0 0 1.455-1.864zm3.374 2.06a.5.5 0 0 1 .513.024L12 11.566V10h1v2.5a.5.5 0 0 1-.777.416L10 11.434V13H9v-2.5a.5.5 0 0 1 .264-.44M9 6v3h1V8h2v1h1V6h-1v1h-2V6zm3-2v-.5a1.5 1.5 0 0 0-3 0V5h4V4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomInOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m14.5 14.5l-4-4M6.5 4v5M4 6.5h5m-2.5 6a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomInSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 6.5a6.5 6.5 0 1 1 11.436 4.23l3.418 3.416l-.707.708l-3.418-3.418A6.5 6.5 0 0 1 0 6.5M6 9V7H4V6h2V4h1v2h2v1H7v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOutOutline(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m14.5 14.5l-4-4M4 6.5h5m-2.5 6a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOutSolid(children ...ElementRenderer) *TeenyiconsIcon {
	return &TeenyiconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 6.5a6.5 6.5 0 1 1 11.436 4.23l3.418 3.416l-.707.708l-3.418-3.418A6.5 6.5 0 0 1 0 6.5M4 7h5V6H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
