package icomoon_free

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 0 0"
	fill           = "currentColor"
)

type IcomoonFreeIcon struct {
	*SVGSVGElement
}

func Accessibility(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 1.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 6.5 1.5"/><path fill="currentColor" d="m10 5l5.15-2.221l-.371-.929L8.5 4h-1L1.221 1.85l-.371.929L6 5v4l-2.051 6.634l.935.355L7.786 9.5h.429l2.902 6.489l.935-.355L10.001 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressBook(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0v16h12V0zm6 4.005a1.995 1.995 0 1 1 0 3.99a1.995 1.995 0 0 1 0-3.99M12 12H6v-1a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2zM1 1h1.5v3H1zm0 4h1.5v3H1zm0 4h1.5v3H1zm0 4h1.5v3H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AidKit(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4h-3V2c0-.55-.45-1-1-1H6c-.55 0-1 .45-1 1v2H2C.9 4 0 4.9 0 6v8c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2M6 2h4v2H6zm6 9H9v3H7v-3H4V9h3V6h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplane(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9.999L9.143 7.142L16 1.999l-2-2l-8.571 3.429L2.731.729C1.953-.049.867-.235.317.315s-.364 1.636.414 2.414l2.698 2.698L0 13.999l2 2l5.144-6.857l2.857 2.857v4h2l1-3l3-1v-2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2a7 7 0 1 0 0 14A7 7 0 0 0 8 2m0 12.625a5.624 5.624 0 1 1 0-11.25a5.624 5.624 0 1 1 0 11.25m6.606-10.138a3 3 0 0 0-4.98-3.321a8.008 8.008 0 0 1 4.98 3.322zM6.374 1.166a3 3 0 0 0-4.98 3.321a8.006 8.006 0 0 1 4.98-3.322z"/><path fill="currentColor" d="M8 9V5H7v5h4V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amazon(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.463 13.831c-1.753 1.294-4.291 1.981-6.478 1.981c-3.066 0-5.825-1.131-7.912-3.019c-.163-.147-.019-.35.178-.234c2.253 1.313 5.041 2.1 7.919 2.1c1.941 0 4.075-.403 6.041-1.238c.294-.125.544.197.253.409z"/><path fill="currentColor" d="M15.191 13c-.225-.287-1.481-.137-2.047-.069c-.172.019-.197-.128-.044-.238c1.003-.703 2.647-.5 2.838-.266c.194.238-.05 1.884-.991 2.672c-.144.122-.281.056-.219-.103c.216-.528.688-1.709.463-1.997zm-4.138-1.162l.003.003c.387-.341 1.084-.95 1.478-1.278c.156-.125.128-.334.006-.509c-.353-.488-.728-.884-.728-1.784v-3c0-1.272.088-2.438-.847-3.313c-.738-.706-1.963-.956-2.9-.956c-1.831 0-3.875.684-4.303 2.947c-.047.241.131.369.287.403l1.866.203c.175-.009.3-.181.334-.356c.159-.778.813-1.156 1.547-1.156c.397 0 .847.144 1.081.5c.269.397.234.938.234 1.397v.25c-1.116.125-2.575.206-3.619.666c-1.206.522-2.053 1.584-2.053 3.147c0 2 1.259 3 2.881 3c1.369 0 2.116-.322 3.172-1.403c.35.506.463.753 1.103 1.284a.395.395 0 0 0 .456-.044zm-1.94-4.694c0 .75.019 1.375-.359 2.041c-.306.544-.791.875-1.331.875c-.737 0-1.169-.563-1.169-1.394c0-1.641 1.472-1.938 2.863-1.938v.416z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6c-.55 0-1 .45-1 1v4c0 .55.45 1 1 1s1-.45 1-1V7c0-.55-.45-1-1-1M2 6c-.55 0-1 .45-1 1v4c0 .55.45 1 1 1s1-.45 1-1V7c0-.55-.45-1-1-1m1.5 5.5A1.5 1.5 0 0 0 5 13v2c0 .55.45 1 1 1s1-.45 1-1v-2h2v2c0 .55.45 1 1 1s1-.45 1-1v-2a1.5 1.5 0 0 0 1.5-1.5V6h-9zM12.472 5a4.5 4.5 0 0 0-2.025-3.276l.5-1.001a.5.5 0 0 0-.895-.447L9.55 1.28l-.13-.052a4.504 4.504 0 0 0-2.84 0l-.13.052L5.948.276a.5.5 0 0 0-.895.447l.5 1.001A4.499 4.499 0 0 0 3.528 5v.5H12.5V5zM6.5 4a.5.5 0 0 1-.001-1h.002A.5.5 0 0 1 6.5 4m3 0a.5.5 0 0 1-.001-1h.003a.5.5 0 0 1-.001 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Angry(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13m3.002 10.699C10.39 11.181 9.275 10.5 8 10.5s-2.389.681-3.002 1.699l-1.286-.772C4.586 9.973 6.179 9 8 9s3.414.973 4.288 2.427zm.983-7.82a.5.5 0 0 1-.364.606c-.275.07-.602.189-.89.334A.998.998 0 0 1 9.999 7a1 1 0 0 1-1-1l.002-.054c.032-.741.706-1.234 1.275-1.518c.543-.271 1.08-.407 1.102-.413a.5.5 0 0 1 .606.364zm-7.97 0a.5.5 0 0 1 .606-.364c.023.006.559.141 1.102.413c.568.284 1.243.776 1.275 1.518L7 6a1 1 0 1 1-1.732-.681a4.638 4.638 0 0 0-.89-.334a.5.5 0 0 1-.364-.606z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngryTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m1.001 5.946c.032-.741.706-1.234 1.275-1.518c.543-.271 1.08-.407 1.102-.413a.5.5 0 1 1 .242.97c-.275.07-.602.189-.89.334A.998.998 0 0 1 9.998 7a1 1 0 0 1-1-1zM4.015 4.379a.5.5 0 0 1 .606-.364c.023.006.559.141 1.102.413c.568.284 1.243.776 1.275 1.518L7 6a1 1 0 1 1-1.732-.681a4.638 4.638 0 0 0-.89-.334a.5.5 0 0 1-.364-.606zm6.987 7.82C10.39 11.181 9.275 10.5 8 10.5s-2.389.681-3.002 1.699l-1.286-.772C4.586 9.973 6.179 9 8 9s3.414.973 4.288 2.427z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Appleinc(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.367 8.501c-.02-2.026 1.652-2.998 1.727-3.046c-.94-1.375-2.404-1.564-2.926-1.585c-1.246-.126-2.431.734-3.064.734c-.631 0-1.607-.715-2.64-.696c-1.358.02-2.61.79-3.31 2.006c-1.411 2.448-.361 6.076 1.014 8.061c.672.972 1.473 2.064 2.525 2.025c1.013-.04 1.396-.656 2.621-.656s1.569.656 2.641.635c1.09-.02 1.781-.991 2.448-1.966c.772-1.128 1.089-2.219 1.108-2.275c-.024-.011-2.126-.816-2.147-3.236zm-2.014-5.946c.558-.677.935-1.617.832-2.555c-.804.033-1.779.536-2.356 1.212c-.518.6-.971 1.557-.85 2.476c.898.07 1.815-.456 2.373-1.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15.5L15.5 8H11V0H5v8H.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.5 16l-4-4L16 3.5L12.5 0L4 8.5l-4-4V16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.293 2.293L4 10.586V7a1 1 0 0 0-2 0v6a1.002 1.002 0 0 0 1 1v.001h6a1 1 0 0 0 0-2H5.414l8.293-8.293a.997.997 0 0 0 0-1.414a.999.999 0 0 0-1.414 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 4.5l-4 4L3.5 0L0 3.5L8.5 12l-4 4H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.293 3.707L10.586 12H7a1 1 0 0 0 0 2h6a1.002 1.002 0 0 0 1-1h.001V7a1 1 0 0 0-2 0v3.586L3.708 2.293a.997.997 0 0 0-1.414 0a.999.999 0 0 0 0 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.707 9.707l-5 5a.999.999 0 0 1-1.414 0l-5-5a.999.999 0 1 1 1.414-1.414L7 11.586V2a1 1 0 0 1 2 0v9.586l3.293-3.293a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 8L8 15.5V11h8V5H8V.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.293 13.707l-5-5a.999.999 0 0 1 0-1.414l5-5a.999.999 0 1 1 1.414 1.414L4.414 7H14a1 1 0 0 1 0 2H4.414l3.293 3.293a.997.997 0 0 1 0 1.414a.999.999 0 0 1-1.414 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 8L8 .5V5H0v6h8v4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.707 13.707l5-5a.999.999 0 0 0 0-1.414l-5-5a.999.999 0 1 0-1.414 1.414L11.586 7H2a1 1 0 0 0 0 2h9.586l-3.293 3.293a.997.997 0 0 0 0 1.414a.999.999 0 0 0 1.414 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .5L.5 8H5v8h6V8h4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 11.5l4-4l8.5 8.5l3.5-3.5L7.5 4l4-4H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.707 12.293L5.414 4H9a1 1 0 0 0 0-2H3a1.002 1.002 0 0 0-1 1h-.001v6a1 1 0 0 0 2 0V5.414l8.293 8.293a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.5 0l4 4L0 12.5L3.5 16L12 7.5l4 4V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.707 13.707L12 5.414V9a1 1 0 0 0 2 0V3a1.002 1.002 0 0 0-1-1v-.001H7a1 1 0 0 0 0 2h3.586l-8.293 8.293a.997.997 0 0 0 0 1.414a.999.999 0 0 0 1.414 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.707 6.293l-5-5a.999.999 0 0 0-1.414 0l-5 5a.999.999 0 1 0 1.414 1.414L7 4.414V14a1 1 0 0 0 2 0V4.414l3.293 3.293a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.404 5.11L9.389 4.096L4.314 9.17a2.152 2.152 0 1 0 3.045 3.044l6.09-6.089a3.588 3.588 0 0 0-5.075-5.074L1.98 7.444l-.014.013c-1.955 1.955-1.955 5.123 0 7.077s5.123 1.954 7.078 0l.013-.014l.001.001l4.365-4.364l-1.015-1.014l-4.365 4.363l-.013.013a3.573 3.573 0 0 1-5.048 0a3.572 3.572 0 0 1 .014-5.06l-.001-.001L9.39 2.065c.839-.84 2.205-.84 3.045 0s.839 2.205 0 3.044l-6.09 6.089a.718.718 0 0 1-1.015-1.014l5.075-5.075z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13m3 9L7.5 8L11 5.5zm-4 0L3.5 8L7 5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackwardTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2.5v5l5-5v11l-5-5v5L3.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baffled(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13"/><path fill="currentColor" d="M6 6.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0"/><path fill="currentColor" d="M5.5 5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 5.5 5m0-1C4.122 4 3 5.122 3 6.5S4.122 9 5.5 9S8 7.878 8 6.5S6.878 4 5.5 4M11 6.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0"/><path fill="currentColor" d="M10.5 5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 10.5 5m0-1C9.121 4 8 5.122 8 6.5S9.121 9 10.5 9S13 7.878 13 6.5S11.879 4 10.5 4M6 11h4v1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BaffledTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 6.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0m5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0"/><path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M4 6.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 4 6.5m6 5.5H6v-1h4zm.5-4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 10.5 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2h2v10H0zm3 0h1v10H3zm2 0h1v10H5zm3 0h1v10H8zm4 0h1v10h-1zm3 0h1v10h-1zm-5 0h.5v10H10zM7 2h.5v10H7zm6.5 0h.5v10h-.5zM0 13h1v1H0zm3 0h1v1H3zm2 0h1v1H5zm5 0h1v1h-1zm5 0h1v1h-1zm-3 0h2v1h-2zm-5 0h2v1H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basecamp(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1.666c-2.919 0-5.169 2.444-6.444 4.838a13.625 13.625 0 0 0-1.453 4.331c-.025.172-.05.344-.069.519c-.009.094-.019.188-.025.281c-.009.119-.003.156.059.256c.187.303.409.584.659.838c.512.525 1.134.928 1.794 1.241c1.503.709 3.2.966 4.85 1.022c1.703.056 3.453-.084 5.081-.616c1.391-.453 2.731-1.244 3.503-2.522c.084-.137.025-.341.009-.5a11.71 11.71 0 0 0-.297-1.66a13.737 13.737 0 0 0-.728-2.159c-1.088-2.525-3.1-5.219-5.963-5.775a5.096 5.096 0 0 0-.978-.094zm.1 12.243c-1.784 0-3.728-.159-5.334-1.019c-.625-.334-1.262-.819-1.563-1.484c-.087-.194-.056-.269-.016-.497c.028-.147.041-.291.106-.428c.091-.191.184-.378.281-.566c.328-.634.681-1.262 1.091-1.853c.203-.291.419-.578.669-.828c.175-.175.388-.362.634-.422c.756-.181 1.334.694 1.794 1.134c.222.213.519.453.85.412c.228-.028.431-.206.594-.353c.553-.497.997-1.112 1.456-1.691c.228-.284.453-.572.7-.844c.166-.184.347-.394.569-.513c.397-.216.903.228 1.178.456c.469.391.884.847 1.281 1.309c.378.441.744.888 1.066 1.372c.497.75.928 1.55 1.322 2.359c.084.175.113.294.144.488c.019.106.059.228.044.338c-.022.153-.128.319-.206.444a3.074 3.074 0 0 1-.719.769c-1.166.903-2.744 1.203-4.178 1.338a18.56 18.56 0 0 1-1.762.078z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.641 3.206c.472 0 .897.041 1.284.125c.388.081.716.219.994.406c.275.188.487.438.644.75c.15.309.225.697.225 1.156c0 .497-.112.909-.338 1.241c-.228.331-.559.6-1.003.813c.606.175 1.053.481 1.353.916c.3.438.444.963.444 1.581c0 .5-.097.928-.287 1.291a2.535 2.535 0 0 1-.778.891c-.325.231-.7.4-1.119.509a5.052 5.052 0 0 1-1.287.166H.001V3.207h4.641zm-.282 3.975c.384 0 .703-.091.953-.275c.25-.181.369-.481.369-.894c0-.228-.041-.419-.122-.566a.905.905 0 0 0-.334-.344a1.325 1.325 0 0 0-.478-.172a2.98 2.98 0 0 0-.556-.05H2.166v2.3H4.36zm.119 4.191c.213 0 .416-.019.606-.063a1.43 1.43 0 0 0 .509-.209c.144-.097.266-.225.353-.394c.088-.166.128-.378.128-.637c0-.506-.144-.869-.428-1.088c-.284-.216-.662-.322-1.131-.322h-2.35v2.709h2.313zm6.853-.034c.294.287.716.431 1.266.431c.394 0 .738-.1 1.022-.3s.456-.412.522-.631h1.725c-.278.859-.697 1.469-1.272 1.838c-.566.369-1.259.556-2.063.556a4.11 4.11 0 0 1-1.519-.269a3.22 3.22 0 0 1-1.15-.766a3.51 3.51 0 0 1-.725-1.188a4.382 4.382 0 0 1-.256-1.519c0-.534.088-1.031.262-1.491c.178-.463.422-.859.747-1.194s.706-.6 1.156-.794a3.692 3.692 0 0 1 1.488-.291c.603 0 1.131.116 1.584.353c.45.234.822.55 1.113.944s.497.847.625 1.353s.172 1.034.137 1.588h-5.147c0 .559.188 1.094.484 1.378zm2.247-3.744c-.231-.256-.628-.397-1.106-.397c-.313 0-.572.053-.778.159a1.564 1.564 0 0 0-.497.394a1.396 1.396 0 0 0-.262.503c-.05.172-.081.331-.091.469h3.188c-.047-.5-.219-.869-.453-1.128zM10.444 4h3.991v.972h-3.991z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.316 7.009c.203-.147.3-.391.3-.728a.936.936 0 0 0-.097-.459a.72.72 0 0 0-.272-.278a1.075 1.075 0 0 0-.388-.141a2.386 2.386 0 0 0-.453-.041H3.759v1.869H5.54c.313.003.572-.072.775-.222zm.278 1.688c-.231-.175-.537-.262-.919-.262H3.759v2.203h1.878c.175 0 .338-.016.494-.05s.297-.088.416-.169c.119-.078.216-.184.287-.319s.106-.309.106-.519c0-.412-.116-.706-.347-.884z"/><path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5m-4.012 4.209h3.241V5h-3.241zm-2.025 6.516a2.06 2.06 0 0 1-.631.725a2.82 2.82 0 0 1-.909.416A4.007 4.007 0 0 1 5.879 12H2.001V4.003H5.77c.381 0 .728.034 1.044.1c.313.066.581.178.806.331c.222.153.397.356.522.609c.122.25.184.566.184.938c0 .403-.091.737-.275 1.006s-.453.487-.816.659c.494.141.856.391 1.097.744c.244.356.363.784.363 1.284c.003.409-.075.759-.231 1.05zm6.528-1.237h-4.178c0 .456.156.891.394 1.125c.238.231.581.35 1.028.35c.322 0 .597-.081.831-.244c.231-.162.372-.334.425-.512h1.4c-.225.697-.566 1.194-1.031 1.494c-.459.3-1.022.45-1.675.45c-.456 0-.866-.075-1.234-.219a2.613 2.613 0 0 1-.934-.622a2.854 2.854 0 0 1-.588-.966A3.562 3.562 0 0 1 9.22 9.11c0-.434.072-.838.213-1.213a2.815 2.815 0 0 1 1.544-1.616a3.013 3.013 0 0 1 1.206-.234c.491 0 .919.094 1.287.287c.366.191.666.447.903.769s.403.688.509 1.1c.103.406.137.834.109 1.284z"/><path fill="currentColor" d="M12.134 7.247c-.253 0-.466.044-.631.131s-.3.194-.403.319a1.131 1.131 0 0 0-.213.409a1.725 1.725 0 0 0-.072.381h2.588c-.037-.406-.178-.706-.366-.916c-.194-.213-.512-.325-.903-.325z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.023 12.5c0-4.5-4-3.5-4-7c0-.29-.028-.538-.079-.749c-.263-1.766-1.44-3.183-2.965-3.615A.86.86 0 0 0 9 .945C9 .425 8.55 0 8 0S7 .425 7 .945c0 .065.007.129.021.191c-1.71.484-2.983 2.208-3.02 4.273L4 5.5C4 9 0 8 0 12.5c0 1.191 2.665 2.187 6.234 2.439a2 2 0 0 0 3.532 0C13.334 14.688 16 13.691 16 12.5v-.011l.024.011zm-3.113.845c-.847.226-1.846.389-2.918.479a2 2 0 0 0-3.984 0c-1.072-.09-2.071-.253-2.918-.479c-1.166-.311-1.724-.659-1.928-.845c.204-.186.762-.534 1.928-.845c1.356-.362 3.1-.561 4.91-.561s3.554.199 4.91.561c1.166.311 1.724.659 1.928.845c-.204.186-.762.534-1.928.845"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bin(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5v10c0 .55.45 1 1 1h9c.55 0 1-.45 1-1V5zm3 9H4V7h1zm2 0H6V7h1zm2 0H8V7h1zm2 0h-1V7h1zm2.25-12H10V.75A.753.753 0 0 0 9.25 0h-3.5A.753.753 0 0 0 5 .75V2H1.75a.752.752 0 0 0-.75.75V4h13V2.75a.752.752 0 0 0-.75-.75M9 2H6v-.987h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 16h10l1-11H2zm7-14V0H6v2H1v3l1-1h12l1 1V2zM9 2H7V1h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Binoculars(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0h6v1H1zm8 0h6v1H9zm5.875 5H14V1h-4v4H6V1H2v4h-.875C.506 5 0 5.506 0 6.125v8.75C0 15.494.506 16 1.125 16h4.75C6.494 16 7 15.494 7 14.875V9h2v5.875C9 15.494 9.506 16 10.125 16h4.75c.619 0 1.125-.506 1.125-1.125v-8.75C16 5.506 15.494 5 14.875 5M5.438 15H1.563C1.254 15 1 14.775 1 14.5s.253-.5.563-.5h3.875c.309 0 .563.225.563.5s-.253.5-.563.5M8.5 8h-1c-.275 0-.5-.225-.5-.5s.225-.5.5-.5h1c.275 0 .5.225.5.5s-.225.5-.5.5m5.938 7h-3.875c-.309 0-.563-.225-.563-.5s.253-.5.563-.5h3.875c.309 0 .563.225.563.5s-.253.5-.563.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blocked(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.657 2.343C12.146.832 10.137 0 8 0S3.854.832 2.343 2.343C.832 3.854 0 5.863 0 8s.832 4.146 2.343 5.657C3.854 15.168 5.863 16 8 16s4.146-.832 5.657-2.343C15.168 12.146 16 10.137 16 8s-.832-4.146-2.343-5.657M14 8a5.97 5.97 0 0 1-1.111 3.475L4.525 3.111A5.97 5.97 0 0 1 8 2c3.308 0 6 2.692 6 6M2 8a5.97 5.97 0 0 1 1.111-3.475l8.364 8.364A5.97 5.97 0 0 1 8 14c-3.308 0-6-2.692-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blog(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0v1.5a8.46 8.46 0 0 1 6.01 2.489a8.472 8.472 0 0 1 2.489 6.01h1.5c0-5.523-4.477-10-10-10z"/><path fill="currentColor" d="M6 3v1.5c1.469 0 2.85.572 3.889 1.611S11.5 8.531 11.5 10H13a7 7 0 0 0-7-7m1.5 3l-1 1L3 8l-3 6.5l.396.396l3.638-3.638a1 1 0 1 1 .707.707l-3.638 3.638l.396.396l6.5-3l1-3.5l1-1l-2.5-2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blogger(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.966 6h-.897C13.52 6 13.038 5.535 13 5c0-2.854-2.301-5-5.175-5H5.203C2.331 0 .002 2.313 0 5.167v5.669c0 2.854 2.331 5.165 5.203 5.165h5.6c2.874 0 5.197-2.311 5.197-5.165V7.174c0-.57-.46-1.173-1.034-1.173zM5 4h3c.55 0 1 .45 1 1s-.45 1-1 1H5c-.55 0-1-.45-1-1s.45-1 1-1m6 8H5c-.55 0-1-.45-1-1s.45-1 1-1h6c.55 0 1 .45 1 1s-.45 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BloggerTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M14 10.125C14 12.266 12.259 14 10.103 14h-4.2C3.747 14 2 12.266 2 10.125v-4.25C2 3.734 3.747 2 5.903 2h1.966c2.156 0 3.881 1.609 3.881 3.75c.028.4.391.75.8.75h.672c.431 0 .775.453.775.881v2.744z"/><path fill="currentColor" d="M11 10c0 .55-.45 1-1 1H6c-.55 0-1-.45-1-1s.45-1 1-1h4c.55 0 1 .45 1 1M9 6c0 .55-.45 1-1 1H6c-.55 0-1-.45-1-1s.45-1 1-1h2c.55 0 1 .45 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.061 7.573A3.982 3.982 0 0 0 12 5c0-2.206-1.794-4-4-4H3v14h6c2.206 0 4-1.794 4-4a4.002 4.002 0 0 0-1.939-3.427M6 3h1.586c.874 0 1.586.897 1.586 2s-.711 2-1.586 2H6zm2.484 10H6V9h2.484c.913 0 1.656.897 1.656 2s-.743 2-1.656 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2v13H3.5a1.5 1.5 0 1 1 0-3H13V0H3C1.9 0 1 .9 1 2v12c0 1.1.9 2 2 2h12V2z"/><path fill="currentColor" d="M3.501 13a.5.5 0 0 0 0 1H13v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0v16l5-5l5 5V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmarks(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2v14l5-5l5 5V2zm8-2H2v14l1-1V1h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Books(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 2h-3c-.275 0-.5.225-.5.5v11c0 .275.225.5.5.5h3c.275 0 .5-.225.5-.5v-11c0-.275-.225-.5-.5-.5M3 5H1V4h2zm5.5-3h-3c-.275 0-.5.225-.5.5v11c0 .275.225.5.5.5h3c.275 0 .5-.225.5-.5v-11c0-.275-.225-.5-.5-.5M8 5H6V4h2z"/><path fill="currentColor" d="m11.954 2.773l-2.679 1.35a.502.502 0 0 0-.222.671l4.5 8.93a.502.502 0 0 0 .671.222l2.679-1.35a.502.502 0 0 0 .222-.671l-4.5-8.93a.502.502 0 0 0-.671-.222"/><path fill="currentColor" d="M14.5 13.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxAdd(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 1H3L0 4v10.5a.5.5 0 0 0 .5.5h15a.5.5 0 0 0 .5-.5V4zM8 13L3 9h3V6h4v3h3zM2.414 3l1-1h9.172l1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxRemove(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 1H3L0 4v10.5a.5.5 0 0 0 .5.5h15a.5.5 0 0 0 .5-.5V4zm-3 9v3H6v-3H3l5-4l5 4zM2.414 3l1-1h9.171l1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4h-4V3c0-.55-.45-1-1-1H6c-.55 0-1 .45-1 1v1H1c-.55 0-1 .45-1 1v9c0 .55.45 1 1 1h14c.55 0 1-.45 1-1V5c0-.55-.45-1-1-1m-9-.998L6.002 3h3.996l.002.002V4H6zM15 8h-2v1.5c0 .275-.225.5-.5.5h-1a.501.501 0 0 1-.5-.5V8H5v1.5c0 .275-.225.5-.5.5h-1a.501.501 0 0 1-.5-.5V8H1V7h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessContrast(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4a4 4 0 1 0 0 8a4 4 0 0 0 0-8m0 6.5v-5c1.379 0 2.5 1.122 2.5 2.5S9.379 10.5 8 10.5M8 13a1 1 0 0 1 1 1v1a1 1 0 0 1-2 0v-1a1 1 0 0 1 1-1M8 3a1 1 0 0 1-1-1V1a1 1 0 0 1 2 0v1a1 1 0 0 1-1 1m7 4a1 1 0 0 1 0 2h-1a1 1 0 0 1 0-2zM3 8a1 1 0 0 1-1 1H1a1 1 0 0 1 0-2h1a1 1 0 0 1 1 1m9.95 3.536l.707.707a1 1 0 0 1-1.414 1.414l-.707-.707a1 1 0 0 1 1.414-1.414m-9.9-7.072l-.707-.707a.999.999 0 1 1 1.414-1.414l.707.707A.999.999 0 1 1 3.05 4.464m9.9 0a.999.999 0 1 1-1.414-1.414l.707-.707a.999.999 0 1 1 1.414 1.414zm-9.9 7.072a1 1 0 0 1 1.414 1.414l-.707.707a1 1 0 0 1-1.414-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bubble(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c4.418 0 8 2.91 8 6.5S12.418 14 8 14c-.424 0-.841-.027-1.247-.079c-1.718 1.718-3.77 2.027-5.753 2.072v-.421c1.071-.525 2-1.48 2-2.572a3.01 3.01 0 0 0-.034-.448C1.157 11.36 0 9.54 0 7.5C0 3.91 3.582 1 8 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3c-.858 0-1.687.135-2.464.402c-.73.251-1.38.605-1.932 1.054C2.569 5.297 2 6.378 2 7.5c0 .63.175 1.24.52 1.815c.356.592.89 1.134 1.547 1.566c.474.312.793.812.878 1.373c.028.187.046.376.053.564c.117-.097.23-.201.342-.312a2 2 0 0 1 1.666-.57c.328.042.662.063.995.063c.858 0 1.687-.135 2.464-.402a6.278 6.278 0 0 0 1.932-1.054c1.035-.841 1.604-1.922 1.604-3.044s-.57-2.203-1.604-3.044a6.299 6.299 0 0 0-1.932-1.054a7.563 7.563 0 0 0-2.464-.402zm0-2c4.418 0 8 2.91 8 6.5S12.418 14 8 14c-.424 0-.841-.027-1.247-.079c-1.718 1.718-3.77 2.027-5.753 2.072v-.421c1.071-.525 2-1.48 2-2.572a3.01 3.01 0 0 0-.034-.448C1.157 11.36 0 9.54 0 7.5C0 3.91 3.582 1 8 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bubbles(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 14.081c0 .711.407 1.327 1 1.628v.249a3.71 3.71 0 0 1-3.207-1.125a5.03 5.03 0 0 1-1.293.168c-2.485 0-4.5-1.791-4.5-4s2.015-4 4.5-4s4.5 1.791 4.5 4c0 .865-.309 1.665-.834 2.32a1.81 1.81 0 0 0-.166.761zM8 0c4.351 0 7.89 2.822 7.997 6.336a6.123 6.123 0 0 0-2.497-.524c-1.493 0-2.903.523-3.971 1.472C8.422 8.268 7.812 9.588 7.812 11c0 .698.149 1.373.433 1.997a10.089 10.089 0 0 1-1.493-.076c-1.718 1.718-3.77 2.027-5.753 2.072v-.421c1.071-.525 2-1.48 2-2.572a3.01 3.01 0 0 0-.034-.448C1.156 10.36-.001 8.54-.001 6.5c0-3.59 3.582-6.5 8-6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubblesFour(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 2a7 7 0 0 0-2.272.369a5.79 5.79 0 0 0-1.773.963C2.517 4.091 2 5.063 2 6.069c0 .562.157 1.109.467 1.623c.323.537.811 1.028 1.41 1.421a2 2 0 0 1 .881 1.374c.014.094.025.188.034.282a3.62 3.62 0 0 0 .127-.122a1.999 1.999 0 0 1 1.662-.567c.303.038.611.058.918.058a7 7 0 0 0 2.272-.369a5.807 5.807 0 0 0 1.774-.963C12.483 8.047 13 7.075 13 6.069s-.517-1.978-1.455-2.737a5.753 5.753 0 0 0-1.774-.963A7 7 0 0 0 7.499 2zm0-2C11.642 0 15 2.717 15 6.069s-3.358 6.069-7.5 6.069a9.16 9.16 0 0 1-1.169-.074C4.72 13.669 2.86 13.956 1 13.999v-.393c1.004-.49 1.813-1.382 1.813-2.402c0-.142-.011-.282-.032-.419C1.085 9.672 0 7.973 0 6.068C0 2.716 3.358-.001 7.5-.001zm8.063 13.604c0 .874.567 1.639 1.438 2.059V16c-1.611-.036-3.09-.283-4.487-1.658c-.33.041-.669.063-1.013.063c-1.492 0-2.866-.402-3.963-1.079c2.261-.008 4.395-.732 6.013-2.042a7.346 7.346 0 0 0 1.913-2.302a6.23 6.23 0 0 0 .704-3.4C17.302 6.518 18 7.795 18 9.202c0 1.633-.94 3.089-2.41 4.043a2.361 2.361 0 0 0-.027.359"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubblesThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 14.081c0 .711.407 1.327 1 1.628v.249a3.71 3.71 0 0 1-3.207-1.125a5.03 5.03 0 0 1-1.293.168c-2.485 0-4.5-1.791-4.5-4s2.015-4 4.5-4s4.5 1.791 4.5 4c0 .865-.309 1.665-.834 2.32a1.81 1.81 0 0 0-.166.761zM3.604 3.456C2.569 4.297 2 5.378 2 6.5c0 .63.175 1.24.52 1.815c.356.592.89 1.134 1.547 1.566c.474.312.793.812.878 1.373c.028.187.046.376.053.564c.117-.097.23-.201.342-.312a2 2 0 0 1 1.666-.57c.327.042.662.063.994.063v2c-.424 0-.84-.027-1.246-.079c-1.718 1.718-3.77 2.027-5.753 2.072v-.421c1.071-.525 2-1.48 2-2.572a3.01 3.01 0 0 0-.034-.448C1.158 10.359.001 8.539.001 6.499c0-3.59 3.582-6.5 8-6.5c4.351 0 7.89 2.822 7.997 6.336a6.083 6.083 0 0 0-2.067-.509c-.18-.876-.709-1.7-1.535-2.371a6.299 6.299 0 0 0-1.932-1.054C9.687 2.134 8.858 1.999 8 1.999s-1.687.135-2.464.402c-.73.251-1.38.605-1.932 1.054z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubblesTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0C11.642 0 15 2.717 15 6.069s-3.358 6.069-7.5 6.069a9.16 9.16 0 0 1-1.169-.074C4.72 13.669 2.86 13.956 1 13.999v-.393c1.004-.49 1.813-1.382 1.813-2.402c0-.142-.011-.282-.032-.419C1.085 9.672 0 7.973 0 6.068C0 2.716 3.358-.001 7.5-.001zm8.063 13.604c0 .874.567 1.639 1.438 2.059V16c-1.611-.036-3.09-.283-4.487-1.658c-.33.041-.669.063-1.013.063c-1.492 0-2.866-.402-3.963-1.079c2.261-.008 4.395-.732 6.013-2.042a7.346 7.346 0 0 0 1.913-2.302a6.23 6.23 0 0 0 .704-3.4C17.302 6.518 18 7.795 18 9.202c0 1.633-.94 3.089-2.41 4.043a2.361 2.361 0 0 0-.027.359"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 9V8h-3.02a5.785 5.785 0 0 0-1.12-3.004h2.53L15.485.617l-.97-.243l-.905 3.621h-2.729l-.042-.032A2.996 2.996 0 0 0 8.001-.001a2.996 2.996 0 0 0-2.838 3.964l-.042.032H2.392L1.487.374l-.97.243l1.095 4.379h2.53A5.785 5.785 0 0 0 3.022 8H.002v1h3.021c.059.713.242 1.388.526 1.996H1.612L.517 15.375l.97.243l.905-3.621h1.756c.917 1.219 2.303 1.996 3.854 1.996s2.937-.777 3.854-1.996h1.756l.905 3.621l.97-.243l-1.095-4.379h-1.937A5.903 5.903 0 0 0 12.981 9h3.021z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullhorn(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6.707C16 3.568 15.081 1.02 13.946 1h.014h-1.296S9.62 3.287 5.239 4.184c-.134.708-.219 1.551-.219 2.523s.085 1.816.219 2.523c4.382.897 7.425 3.184 7.425 3.184h1.296l-.014-.001C15.082 12.393 16 9.846 16 6.706zm-2.487 4.844c-.147 0-.305-.152-.387-.243c-.197-.22-.387-.562-.55-.989c-.363-.957-.564-2.239-.564-3.611s.2-2.655.564-3.611c.162-.428.353-.77.55-.99c.081-.091.24-.243.387-.243s.305.152.387.243c.197.22.387.562.55.99c.363.957.564 2.239.564 3.611s-.2 2.655-.564 3.611c-.162.428-.353.77-.55.989c-.081.091-.24.243-.387.243M3.935 6.707c0-.812.06-1.6.173-2.33c-.74.102-1.39.161-2.193.161H.867L0 6.017v1.378l.867 1.479h1.048c.803 0 1.453.059 2.193.161a15.326 15.326 0 0 1-.173-2.33zm1.817 3.327l-2-.383l1.279 5.024c.066.26.324.391.573.291l1.852-.741a.427.427 0 0 0 .222-.611zm7.761-1.46c-.057 0-.118-.059-.149-.094a1.268 1.268 0 0 1-.212-.381c-.14-.369-.217-.863-.217-1.392s.077-1.023.217-1.392c.063-.165.136-.297.212-.381c.031-.035.092-.094.149-.094s.118.059.149.094c.076.085.149.217.212.381c.14.369.217.863.217 1.392s-.077 1.023-.217 1.392a1.243 1.243 0 0 1-.212.381c-.031.035-.092.094-.149.094"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 1H1c-.55 0-1 .45-1 1v5c0 .55.45 1 1 1h5c.55 0 1-.45 1-1V2c0-.55-.45-1-1-1m0 4H1V4h5zm8-4H9c-.55 0-1 .45-1 1v13c0 .55.45 1 1 1h5c.55 0 1-.45 1-1V2c0-.55-.45-1-1-1m0 9H9V9h5zm0-3H9V6h5zM6 9H1c-.55 0-1 .45-1 1v5c0 .55.45 1 1 1h5c.55 0 1-.45 1-1v-5c0-.55-.45-1-1-1m0 4H4v2H3v-2H1v-1h2v-2h1v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6h2v2H5zm3 0h2v2H8zm3 0h2v2h-2zm-9 6h2v2H2zm3 0h2v2H5zm3 0h2v2H8zM5 9h2v2H5zm3 0h2v2H8zm3 0h2v2h-2zM2 9h2v2H2zm11-9v1h-2V0H4v1H2V0H0v16h15V0zm1 15H1V4h13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.75 9.5a3.25 3.25 0 1 0 6.5 0a3.25 3.25 0 0 0-6.5 0M15 4h-3.5c-.25-1-.5-2-1.5-2H6C5 2 4.75 3 4.5 4H1c-.55 0-1 .45-1 1v9c0 .55.45 1 1 1h14c.55 0 1-.45 1-1V5c0-.55-.45-1-1-1m-7 9.938a4.438 4.438 0 1 1 0-8.876a4.438 4.438 0 0 1 0 8.876M15 7h-2V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CancelCircle(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13"/><path fill="currentColor" d="M10.5 4L8 6.5L5.5 4L4 5.5L6.5 8L4 10.5L5.5 12L8 9.5l2.5 2.5l1.5-1.5L9.5 8L12 5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 14.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 6 14.5m10 0a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 16 14.5M16 8V2H4a1 1 0 0 0-1-1H0v1h2l.751 6.438A2 2 0 0 0 4 12h12v-1H4a1 1 0 0 1-1-1v-.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxChecked(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 0H2C.9 0 0 .9 0 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V2c0-1.1-.9-2-2-2M7 12.414L3.293 8.707l1.414-1.414L7 9.586l4.793-4.793l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxUnchecked(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 0H2C.9 0 0 .9 0 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V2c0-1.1-.9-2-2-2m0 14H2V2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkmark(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 2L6 9.5L2.5 6L0 8.5l6 6l10-10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckmarkTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.21 14.339L-.007 8.22l3.084-3.035L6.21 8.268l6.713-6.607l3.084 3.035zM1.686 8.22l4.524 4.453l8.104-7.976l-1.391-1.369L6.21 9.935L3.077 6.852z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chrome(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.036 6.977l-2.29-3.966A7.986 7.986 0 0 1 8-.001a7.994 7.994 0 0 1 6.883 3.922H8.355a4.1 4.1 0 0 0-4.32 3.055zm6.828-1.899h4.585a8 8 0 0 1-7.358 10.921l3.272-5.667a4.08 4.08 0 0 0-.499-5.254M5.094 8c0-1.603 1.304-2.906 2.906-2.906S10.906 6.398 10.906 8S9.602 10.906 8 10.906S5.094 9.602 5.094 8m4.003 3.944l-2.29 3.967a8.001 8.001 0 0 1-5.78-11.833l3.266 5.657a4.1 4.1 0 0 0 4.804 2.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleDown(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8A8 8 0 1 0 0 8a8 8 0 0 0 16 0M1.5 8a6.5 6.5 0 1 1 13 0a6.5 6.5 0 0 1-13 0"/><path fill="currentColor" d="M4.957 5.543L3.543 6.957L8 11.414l4.457-4.457l-1.414-1.414L8 8.586z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLeft(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13"/><path fill="currentColor" d="M10.457 4.957L9.043 3.543L4.586 8l4.457 4.457l1.414-1.414L7.414 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleRight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13"/><path fill="currentColor" d="m5.543 11.043l1.414 1.414L11.414 8L6.957 3.543L5.543 4.957L8.586 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleUp(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 8a8 8 0 1 0 16 0A8 8 0 0 0 0 8m14.5 0a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0"/><path fill="currentColor" d="m11.043 10.457l1.414-1.414L8 4.586L3.543 9.043l1.414 1.414L8 7.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClearFormatting(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 14h9v2H0zM14 2H9.273L6.402 13H4.335L7.206 2H3.001V0h11zm.528 14L12.5 13.972L10.472 16l-.972-.972L11.528 13L9.5 10.972l.972-.972l2.028 2.028L14.528 10l.972.972L13.472 13l2.028 2.028z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 2H10a2 2 0 1 0-4 0H1.5a.5.5 0 0 0-.5.5v13a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-13a.5.5 0 0 0-.5-.5M8 1a1 1 0 1 1 0 2a1 1 0 0 1 0-2m6 14H2V3h2v1.5a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5V3h2z"/><path fill="currentColor" d="M7 13.414L3.793 9.707l.914-.914L7 10.586l4.293-3.793l.914.914z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.293 11.707L7 8.414V4h2v3.586l2.707 2.707zM8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14A6 6 0 1 1 8 2a6 6 0 0 1 0 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m2.293 11.707L7 8.414V4h2v3.586l2.707 2.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10.274a2.726 2.726 0 0 0-2.078-2.648A3.72 3.72 0 0 0 10.205 4a3.712 3.712 0 0 0-2.92 1.418a2.09 2.09 0 0 0-3.719 1.573a3.028 3.028 0 0 0-3.567 2.98A3.028 3.028 0 0 0 3.026 13H13.28a2.725 2.725 0 0 0 2.719-2.726z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudCheck(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.942 8.039a2.5 2.5 0 0 0-3.085-2.955a3 3 0 0 0-5.737.075A4 4 0 1 0 4 13h9.5a2.5 2.5 0 0 0 .442-4.961M6.5 12L4 9.5l1-1L6.5 10L10 6.5l1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.922 5.626A3.72 3.72 0 0 0 10.205 2a3.712 3.712 0 0 0-2.92 1.418a2.09 2.09 0 0 0-3.719 1.573a3.028 3.028 0 0 0-3.567 2.98A3.028 3.028 0 0 0 3.026 11H4.46l3.539 3.664L11.538 11h1.742a2.725 2.725 0 0 0 .641-5.374zM8 13l-3-3h2V7h2v3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.942 6.039a2.5 2.5 0 0 0-3.085-2.955a3 3 0 0 0-5.737.075A4 4 0 1 0 4 11h2v3h4v-3h3.5a2.5 2.5 0 0 0 .442-4.961M9 10v3H7v-3H4.5L8 6.5l3.5 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clubs(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.294 6.137c-.922 0-1.751.384-2.341 1.011c-.25.265-.684.58-1.153.856c.22-.842.917-1.902 1.4-2.367a3.27 3.27 0 0 0 1-2.367C11.2 1.475 9.771.018 8-.001C6.229.018 4.8 1.474 4.8 3.27c0 .932.38 1.771 1 2.367c.484.465 1.18 1.525 1.4 2.367c-.469-.277-.903-.591-1.153-.856a3.197 3.197 0 0 0-2.341-1.011C1.919 6.137.47 7.6.47 9.408s1.448 3.271 3.236 3.271c.923 0 1.751-.396 2.341-1.023c.263-.279.726-.627 1.223-.916c-.047 2.308-1.149 4.003-2.271 4.67V16h6v-.59c-1.122-.668-2.224-2.363-2.271-4.67c.498.289.961.637 1.223.916a3.21 3.21 0 0 0 2.341 1.023c1.787 0 3.236-1.464 3.236-3.271s-1.448-3.271-3.236-3.271z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Codepen(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.777 5.751l-7-4.667a.5.5 0 0 0-.555 0l-7 4.667a.501.501 0 0 0-.223.416v4.667c0 .167.084.323.223.416l7 4.667a.5.5 0 0 0 .554 0l7-4.667a.501.501 0 0 0 .223-.416V6.167a.501.501 0 0 0-.223-.416zM7.5 10.232L4.901 8.5L7.5 6.768L10.099 8.5zM8 5.899V2.434l5.599 3.732L11 7.898l-3-2zm-1 0l-3 2l-2.599-1.732L7 2.435V5.9zM3.099 8.5L1 9.899V7.101zM4 9.101l3 2v3.465l-5.599-3.732zm4 2l3-2l2.599 1.732L8 14.565V11.1zM11.901 8.5L14 7.101v2.798z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.59 9.535a3.053 3.053 0 0 1 1.127-4.164l-1.572-2.723a3.017 3.017 0 0 1-1.529.414A3.052 3.052 0 0 1 9.574 0H6.429a3.009 3.009 0 0 1-.406 1.535c-.839 1.454-2.706 1.948-4.17 1.106L.281 5.364a3 3 0 0 1 1.123 1.117a3.053 3.053 0 0 1-1.12 4.16l1.572 2.723c.448-.261.967-.41 1.522-.41A3.052 3.052 0 0 1 6.42 16h3.145a3.012 3.012 0 0 1 .406-1.519a3.053 3.053 0 0 1 4.163-1.11l1.572-2.723a3.008 3.008 0 0 1-1.116-1.113M8 11.24a3.24 3.24 0 1 1 0-6.48a3.24 3.24 0 0 1 0 6.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cogs(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.683 11.282l.645-.903l-.707-.707l-.903.645a2.515 2.515 0 0 0-.535-.222L4 9H3l-.183 1.095a2.444 2.444 0 0 0-.535.222l-.903-.645l-.707.707l.645.903a2.515 2.515 0 0 0-.222.535L0 12v1l1.095.183c.053.188.128.368.222.535l-.645.903l.707.707l.903-.645c.168.094.347.168.535.222L3 16h1l.183-1.095c.188-.053.368-.128.535-.222l.903.645l.707-.707l-.645-.903c.094-.168.168-.347.222-.535L7 13.001v-1l-1.095-.183a2.444 2.444 0 0 0-.222-.535zM3.5 13.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2M16 6V5l-1.053-.191a4.147 4.147 0 0 0-.074-.372l.899-.58l-.383-.924l-1.046.226a4.432 4.432 0 0 0-.211-.315l.609-.88l-.707-.707l-.88.609a4.8 4.8 0 0 0-.315-.211l.226-1.046l-.924-.383l-.58.899a4.53 4.53 0 0 0-.372-.074l-.191-1.053h-1l-.191 1.053c-.126.019-.25.044-.372.074l-.58-.899l-.924.383l.226 1.046a4.432 4.432 0 0 0-.315.211l-.88-.609l-.707.707l.609.88a4.8 4.8 0 0 0-.211.315l-1.046-.226l-.383.924l.899.58a4.53 4.53 0 0 0-.074.372L4.996 5v1l1.053.191c.019.126.044.25.074.372l-.899.58l.383.924l1.046-.226c.066.108.136.213.211.315l-.609.88l.707.707l.88-.609a4.8 4.8 0 0 0 .315.211l-.226 1.046l.924.383l.58-.899c.122.03.246.054.372.074l.191 1.053h1l.191-1.053c.126-.019.25-.044.372-.074l.58.899l.924-.383l-.226-1.046c.108-.066.213-.136.315-.211l.88.609l.707-.707l-.609-.88a4.8 4.8 0 0 0 .211-.315l1.046.226l.383-.924l-.899-.58a4.53 4.53 0 0 0 .074-.372zm-5.5 1.675a2.175 2.175 0 1 1 0-4.35a2.175 2.175 0 0 1 0 4.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinDollar(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15m0 13.5a6 6 0 1 1 0-12a6 6 0 0 1 0 12M8 8V6h2V5H8V4H7v1H5v4h2v2H5v1h2v1h1v-1h2V8zM7 8H6V6h1zm2 3H8V9h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinEuro(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15m0 13.5a6 6 0 1 1 0-12a6 6 0 0 1 0 12"/><path fill="currentColor" d="M10.482 10.068a.5.5 0 0 0-.684.181A1.51 1.51 0 0 1 8.5 11h-2a1.502 1.502 0 0 1-1.414-1H7.5a.5.5 0 0 0 0-1H5V8h2.5a.5.5 0 0 0 0-1H5.086c.206-.582.762-1 1.414-1h2a1.51 1.51 0 0 1 1.298.751a.5.5 0 1 0 .865-.503a2.513 2.513 0 0 0-2.162-1.249h-2c-1.207 0-2.217.86-2.45 2h-.55a.5.5 0 0 0 0 1h.5v1h-.5a.5.5 0 0 0 0 1h.55c.232 1.14 1.242 2 2.45 2h2a2.51 2.51 0 0 0 2.162-1.249a.5.5 0 0 0-.181-.684z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinPound(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15m0 13.5a6 6 0 1 1 0-12a6 6 0 0 1 0 12"/><path fill="currentColor" d="M9.5 11H6V9h1.5a.5.5 0 0 0 0-1H6v-.5a1.502 1.502 0 0 1 2.8-.75a.499.499 0 1 0 .865-.501A2.51 2.51 0 0 0 7.5 4.999a2.503 2.503 0 0 0-2.5 2.5v.5h-.5a.5.5 0 0 0 0 1H5v3h4.5a.5.5 0 0 0 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinYen(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15m0 13.5a6 6 0 1 1 0-12a6 6 0 0 1 0 12"/><path fill="currentColor" d="M9.5 9a.5.5 0 0 0 0-1H8.434l1.482-2.223a.5.5 0 1 0-.832-.554L7.5 7.599L5.916 5.223a.5.5 0 1 0-.832.554L6.566 8H5.5a.5.5 0 0 0 0 1H7v1H5.5a.5.5 0 0 0 0 1H7v1.5a.5.5 0 0 0 1 0V11h1.5a.5.5 0 0 0 0-1H8V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Command(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 14A2.502 2.502 0 0 1 9 11.5V10H7v1.5C7 12.879 5.878 14 4.5 14S2 12.879 2 11.5S3.122 9 4.5 9H6V7H4.5C3.122 7 2 5.878 2 4.5S3.122 2 4.5 2S7 3.122 7 4.5V6h2V4.5C9 3.122 10.121 2 11.5 2S14 3.122 14 4.5S12.879 7 11.5 7H10v2h1.5c1.379 0 2.5 1.121 2.5 2.5S12.879 14 11.5 14M10 10v1.5c0 .827.673 1.5 1.5 1.5s1.5-.673 1.5-1.5s-.673-1.5-1.5-1.5zm-5.5 0c-.827 0-1.5.673-1.5 1.5S3.673 13 4.5 13S6 12.327 6 11.5V10zM7 9h2V7H7zm3-3h1.5c.827 0 1.5-.673 1.5-1.5S12.327 3 11.5 3S10 3.673 10 4.5zM4.5 3C3.673 3 3 3.673 3 4.5S3.673 6 4.5 6H6V4.5C6 3.673 5.327 3 4.5 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 16a.5.5 0 0 1-.5-.5V8H.5a.5.5 0 0 1-.211-.953l15-7a.5.5 0 0 1 .665.665l-7 15a.5.5 0 0 1-.453.289zM2.754 7H8.5a.5.5 0 0 1 .5.5v5.746l5.465-11.712L2.753 6.999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M1.5 8a6.5 6.5 0 0 1 10.93-4.756L6 6l-2.756 6.43A6.476 6.476 0 0 1 1.5 8m7.643 1.143l-4.001 1.715l1.715-4.001zM8 14.5a6.476 6.476 0 0 1-4.43-1.744L10 10l2.756-6.43A6.5 6.5 0 0 1 8 14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Confused(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13M4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m1.345 5h1.014a2.505 2.505 0 0 1-1.624 2.665a2.502 2.502 0 0 1-3.204-1.494a1.502 1.502 0 0 0-1.923-.896A1.502 1.502 0 0 0 4.655 12H3.642a2.505 2.505 0 0 1 1.624-2.665a2.502 2.502 0 0 1 3.204 1.494a1.502 1.502 0 0 0 1.923.896A1.502 1.502 0 0 0 11.346 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfusedTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m3 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2M5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2m5.735 8.665a2.502 2.502 0 0 1-3.204-1.494a1.502 1.502 0 0 0-1.923-.896A1.502 1.502 0 0 0 4.655 12H3.642a2.505 2.505 0 0 1 1.624-2.665a2.502 2.502 0 0 1 3.204 1.494a1.502 1.502 0 0 0 1.923.896A1.502 1.502 0 0 0 11.346 10h1.014a2.505 2.505 0 0 1-1.624 2.665z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Connection(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 9c1.654 0 3.154.673 4.241 1.759l-1.414 1.414C12.103 11.449 11.103 11 10 11s-2.103.449-2.827 1.173l-1.414-1.414A5.982 5.982 0 0 1 10 9M2.929 7.929C4.818 6.04 7.329 5 10 5s5.182 1.04 7.071 2.929l-1.414 1.414C14.146 7.832 12.137 7 10 7s-4.146.832-5.657 2.343zM15.45 2.101a13.966 13.966 0 0 1 4.45 3l-1.414 1.414C16.219 4.249 13.206 3 10.001 3S3.782 4.248 1.516 6.515L.102 5.101A13.955 13.955 0 0 1 10.002 1c1.89 0 3.723.37 5.45 1.101zM9 14a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contrast(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M2 8a6 6 0 0 1 6-6v12a6 6 0 0 1-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cool(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13M12.5 4c.275 0 .5.225.5.5V6c0 .55-.45 1-1 1h-2c-.55 0-1-.45-1-1H7c0 .55-.45 1-1 1H4c-.55 0-1-.45-1-1V4.5c0-.275.225-.5.5-.5h3c.275 0 .5.225.5.5V5h2v-.5c0-.275.225-.5.5-.5zM8 12a3.996 3.996 0 0 0 3.43-1.942l.857.515a4.996 4.996 0 0 1-6.406 1.957l.518-.864c.49.214 1.031.334 1.6.334z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoolTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 13a4.984 4.984 0 0 1-2.118-.47l.518-.864a3.996 3.996 0 0 0 5.03-1.608l.858.515A4.996 4.996 0 0 1 8 13m5-7c0 .55-.45 1-1 1h-2c-.55 0-1-.45-1-1H7c0 .55-.45 1-1 1H4c-.55 0-1-.45-1-1V4.5c0-.275.225-.5.5-.5h3c.275 0 .5.225.5.5V5h2v-.5c0-.275.225-.5.5-.5h3c.275 0 .5.225.5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4V0H3L0 3v9h6v4h10V4zM3 1.414V3H1.414zM1 11V4h3V1h5v3L6 7v4zm8-5.586V7H7.414zM15 15H7V8h3V5h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 2h-13C.675 2 0 2.675 0 3.5v9c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-9c0-.825-.675-1.5-1.5-1.5m-13 1h13c.271 0 .5.229.5.5V5H1V3.5c0-.271.229-.5.5-.5m13 10h-13a.507.507 0 0 1-.5-.5V8h14v4.5c0 .271-.229.5-.5.5M2 10h1v2H2zm2 0h1v2H4zm2 0h1v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 4l3-3l-1-1l-3 3H5V0H3v3H0v2h3v8h8v3h2v-3h3v-2h-3zM5 5h5l-5 5zm1 6l5-5v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cross(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.854 12.854L11 8l4.854-4.854a.503.503 0 0 0 0-.707L13.561.146a.499.499 0 0 0-.707 0L8 5L3.146.146a.5.5 0 0 0-.707 0L.146 2.439a.499.499 0 0 0 0 .707L5 8L.146 12.854a.5.5 0 0 0 0 .707l2.293 2.293a.499.499 0 0 0 .707 0L8 11l4.854 4.854a.5.5 0 0 0 .707 0l2.293-2.293a.499.499 0 0 0 0-.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crying(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13"/><path fill="currentColor" d="M12.5 6h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 0 1m-7 0h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 0 1m4 7.375a.502.502 0 0 1-.354-.146C9.074 13.157 8.686 13 8 13s-1.075.157-1.146.229a.5.5 0 0 1-.707-.707c.471-.471 1.453-.521 1.854-.521s1.383.051 1.854.521a.5.5 0 0 1-.354.853zM11.5 9a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m0 3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m-7-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m0 3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CryingTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M5 11.5a.5.5 0 0 1-1 0v-1a.5.5 0 0 1 1 0zm0-3a.5.5 0 0 1-1 0v-1a.5.5 0 0 1 1 0zM5.5 6h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 0 1m4.354 7.229a.498.498 0 0 1-.708 0C9.074 13.157 8.686 13 8 13s-1.075.157-1.146.229a.5.5 0 0 1-.707-.707c.471-.471 1.453-.521 1.854-.521s1.383.051 1.854.521a.5.5 0 0 1 0 .707zM12 11.5a.5.5 0 0 1-1 0v-1a.5.5 0 0 1 1 0zm0-3a.5.5 0 0 1-1 0v-1a.5.5 0 0 1 1 0zm.5-2.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.381.758l-.537 2.686h10.934l-.342 1.735H1.496l-.53 2.686h10.933l-.61 3.063l-4.406 1.46l-3.819-1.46l.261-1.329H.639L0 12.823l6.316 2.417l7.281-2.417L16 .757z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ctrl(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 7a.5.5 0 0 1-.377-.171l-3.124-3.57l-3.124 3.57a.5.5 0 1 1-.753-.659l3.5-4a.502.502 0 0 1 .752 0l3.5 4a.5.5 0 0 1-.376.83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.582 0 0 1.119 0 2.5v2C0 5.881 3.582 7 8 7s8-1.119 8-2.5v-2C16 1.119 12.418 0 8 0"/><path fill="currentColor" d="M8 8.5C3.582 8.5 0 7.381 0 6v3c0 1.381 3.582 2.5 8 2.5s8-1.119 8-2.5V6c0 1.381-3.582 2.5-8 2.5"/><path fill="currentColor" d="M8 13c-4.418 0-8-1.119-8-2.5v3C0 14.881 3.582 16 8 16s8-1.119 8-2.5v-3c0 1.381-3.582 2.5-8 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delicious(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm8 15V8H1V1h7v7h7v7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deviantart(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.953 2.909V0h-2.909l-.291.294L8.378 2.91l-.431.291h-4.9v3.994h2.694l.241.291l-2.934 5.606v2.909h2.909l.291-.294l1.375-2.616l.431-.291h4.9V8.806H10.26l-.241-.294z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamonds(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0L3 8l5 8l5-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dice(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3h-8A2.507 2.507 0 0 0 3 5.5v8C3 14.875 4.125 16 5.5 16h8c1.375 0 2.5-1.125 2.5-2.5v-8C16 4.125 14.875 3 13.5 3m-7 11a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 6.5 14m0-6a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 6.5 8m3 3a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 9.5 11m3 3a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 12.5 14m0-6a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 12.5 8m.449-6A2.509 2.509 0 0 0 10.5 0h-8A2.507 2.507 0 0 0 0 2.5v8c0 1.204.862 2.216 2 2.449V3c0-.55.45-1 1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Display(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1v10h16V1zm15 9H1V2h14zm-4.5 2h-5L5 14l-1 1h8l-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 9l4-4H9V1H7v4H4zm3.636-1.636l-1.121 1.121L14.579 10L8 12.453L1.421 10l4.064-1.515l-1.121-1.121L0 9v4l8 3l8-3V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.5 7l-4 4l-4-4H6V1h3v6zm-4 4H0v4h15v-4zm6.5 2h-2v-1h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 8h-2.5L8 11.5L4.5 8H2l-2 4v1h16v-1zM0 14h16v1H0zm9-9V1H7v4H3.5L8 9.5L12.5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drawer(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.89 10.188l-4-5A.5.5 0 0 0 11.5 5h-7a.497.497 0 0 0-.39.188l-4 5A.5.5 0 0 0 0 10.5V15a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-4.5a.497.497 0 0 0-.11-.312M15 11h-3.5l-2 2h-3l-2-2H1v-.325L4.74 6h6.519l3.74 4.675z"/><path fill="currentColor" d="M11.5 8h-7a.5.5 0 0 1 0-1h7a.5.5 0 0 1 0 1m1 2h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DrawerTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.89 10.188l-4-5A.5.5 0 0 0 11.5 5h-7a.497.497 0 0 0-.39.188l-4 5A.5.5 0 0 0 0 10.5V15a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-4.5a.497.497 0 0 0-.11-.312M15 11h-3.5l-2 2h-3l-2-2H1v-.325L4.74 6h6.519l3.74 4.675z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16c-4.412 0-8-3.588-8-8s3.587-8 8-8c4.412 0 8 3.587 8 8s-3.588 8-8 8m6.747-6.906c-.234-.075-2.116-.634-4.256-.291a29.7 29.7 0 0 1 1.328 4.872a6.845 6.845 0 0 0 2.928-4.581M10.669 14.3c-.103-.6-.497-2.688-1.456-5.181c-.016.006-.031.009-.044.016c-3.856 1.344-5.241 4.016-5.362 4.266a6.807 6.807 0 0 0 6.863.9zm-7.747-1.722c.156-.266 2.031-3.369 5.553-4.509a7.04 7.04 0 0 1 .269-.081a24.04 24.04 0 0 0-.553-1.159c-3.409 1.022-6.722.978-7.022.975c-.003.069-.003.138-.003.209c0 1.753.666 3.356 1.756 4.566zM1.313 6.609c.306.003 3.122.016 6.319-.831a43.092 43.092 0 0 0-2.534-3.953a6.854 6.854 0 0 0-3.784 4.784zM6.4 1.366a36.612 36.612 0 0 1 2.55 4c2.431-.909 3.459-2.294 3.581-2.469A6.799 6.799 0 0 0 6.4 1.366m6.891 2.325c-.144.194-1.291 1.663-3.816 2.694c.159.325.313.656.453.991c.05.119.1.234.147.353c2.275-.284 4.534.172 4.759.219a6.816 6.816 0 0 0-1.544-4.256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drive(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 14h10a3 3 0 0 0 3-3H0a3 3 0 0 0 3 3m10-2h1v1h-1zm2-10H1l-1 8h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5.5L8 3.5l4.5 3l3.5-3zM8 3.5L4.5.5L0 3.5l3.5 3zm4.5 3l3.5 3l-4.5 2.5L8 9zM8 9L3.5 6.5L0 9.5L4.5 12z"/><path fill="currentColor" d="M11.377 13.212L8 10.317l-3.377 2.895L2.5 12.033V13.5L8 16l5.5-2.5v-1.467z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Droplet(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.51 7.393C12.483 4.527 10.305 1.953 8 0C5.695 1.953 3.518 4.527 2.49 7.393c-.635 1.772-.698 3.696.197 5.397C3.716 14.745 5.791 16 8 16s4.284-1.255 5.313-3.21c.895-1.701.832-3.624.197-5.397m-1.967 4.466A4.045 4.045 0 0 1 8 14a4.03 4.03 0 0 1-2.377-.791c.207.027.416.041.627.041a5.056 5.056 0 0 0 4.428-2.676c.701-1.333.64-2.716.373-3.818c.227.44.42.878.576 1.311c.353.985.625 2.443-.084 3.791z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Earth(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 15a6.96 6.96 0 0 1-2.769-.57l3.643-4.098A.503.503 0 0 0 9 10V8.5a.5.5 0 0 0-.5-.5C6.735 8 4.872 6.165 4.854 6.146A.5.5 0 0 0 4.5 6h-2a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .276.447L4 10.809v2.936A6.992 6.992 0 0 1 1 8a6.97 6.97 0 0 1 .674-3H3.5c.133 0 .26-.053.354-.146l2-2A.5.5 0 0 0 6 2.5V1.29A6.989 6.989 0 0 1 8 1c1.1 0 2.141.254 3.067.706A2.98 2.98 0 0 0 10 3.999c0 .801.312 1.555.879 2.121a2.994 2.994 0 0 0 2.268.875c.216.809.605 2.917-.131 5.818a.466.466 0 0 0-.013.082a6.979 6.979 0 0 1-5.002 2.104z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edge(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.241 7.103C.71 3.403 3.235.047 7.76 0c2.731.053 4.978 1.291 6.316 3.65c.672 1.231.881 2.525.925 3.953v1.678H4.96c.047 4.141 6.094 4 8.697 2.175v3.372c-1.525.916-4.984 1.734-7.662.681c-2.281-.856-3.906-3.244-3.897-5.541c-.075-2.978 1.481-4.95 3.897-6.072c-.513.634-.903 1.334-1.106 2.547h5.669s.331-3.388-3.209-3.388C4.011 3.171 1.605 5.111.243 7.102z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 12h16v2H0zM8 2l8 8H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ello(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m4.885 9.212C12.31 11.442 10.301 13 8 13s-4.31-1.558-4.885-3.788a.708.708 0 0 1 .685-.884c.322 0 .604.218.684.531a3.631 3.631 0 0 0 7.032 0a.707.707 0 1 1 1.369.354z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Embed(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 11.5l1.5 1.5l5-5l-5-5L9 4.5L12.5 8zm-2-7L5.5 3l-5 5l5 5L7 11.5L3.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmbedTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 11.5l1.5 1.5l5-5l-5-5L13 4.5L16.5 8zm-6-7L5.5 3l-5 5l5 5L7 11.5L3.5 8zm3.958-2.148l1.085.296l-3 11l-1.085-.296z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enlarge(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0H9.5L12 2.5l-3 3L10.5 7l3-3L16 6.5zm0 16V9.5L13.5 12l-3-3L9 10.5l3 3L9.5 16zM0 16h6.5L4 13.5l3-3L5.5 9l-3 3L0 9.5zM0 0v6.5L2.5 4l3 3L7 5.5l-3-3L6.5 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnlargeTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0v6.5L13.5 4l-3 3L9 5.5l3-3L9.5 0zM7 10.5l-3 3L6.5 16H0V9.5L2.5 12l3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enter(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 8H1V6h5V4l3 3l-3 3zm10-8v13l-6 3v-3H4V9h1v3h5V3l4-2H5v4H4V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelop(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 2h-13C.675 2 0 2.675 0 3.5v10c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-10c0-.825-.675-1.5-1.5-1.5M6.23 8.6L2 11.895V4.057zM2.756 4h10.488L8 7.938zm3.639 4.777L8 10.5l1.605-1.723L12.895 13h-9.79zM9.77 8.6L14 4.057v7.838z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equalizer(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2v-.25A.752.752 0 0 0 6.25 1h-2.5a.752.752 0 0 0-.75.75V2H0v2h3v.25c0 .412.337.75.75.75h2.5c.412 0 .75-.338.75-.75V4h9V2zM4 4V2h2v2zm9 2.75a.753.753 0 0 0-.75-.75h-2.5a.753.753 0 0 0-.75.75V7H0v2h9v.25c0 .412.338.75.75.75h2.5c.412 0 .75-.338.75-.75V9h3V7h-3zM10 9V7h2v2zm-3 2.75a.753.753 0 0 0-.75-.75h-2.5a.752.752 0 0 0-.75.75V12H0v2h3v.25c0 .412.337.75.75.75h2.5c.412 0 .75-.338.75-.75V14h9v-2H7zM4 14v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EqualizerTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7h.25c.412 0 .75-.338.75-.75v-2.5a.752.752 0 0 0-.75-.75H14V0h-2v3h-.25a.752.752 0 0 0-.75.75v2.5c0 .412.338.75.75.75H12v9h2zm-2-3h2v2h-2zm-2.75 9c.412 0 .75-.338.75-.75v-2.5A.753.753 0 0 0 9.25 9H9V0H7v9h-.25a.753.753 0 0 0-.75.75v2.5c0 .412.338.75.75.75H7v3h2v-3zM7 10h2v2H7zM4.25 7c.412 0 .75-.338.75-.75v-2.5A.752.752 0 0 0 4.25 3H4V0H2v3h-.25a.752.752 0 0 0-.75.75v2.5c0 .412.337.75.75.75H2v9h2V7zM2 4h2v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Evil(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 7a1 1 0 0 1-1-1l.002-.054c.032-.741.706-1.234 1.275-1.518c.543-.271 1.08-.407 1.102-.413a.5.5 0 1 1 .242.97c-.275.07-.602.189-.89.334A.998.998 0 0 1 9.999 7zM4.379 4.985a.5.5 0 1 1 .242-.97c.023.006.559.141 1.102.413c.568.284 1.243.776 1.275 1.518L7 6a1 1 0 1 1-1.732-.681a4.638 4.638 0 0 0-.89-.334zM8 11.5a3.5 3.5 0 0 0 3.002-1.699l1.286.772C11.414 12.027 9.821 13 8 13s-3.413-.973-4.288-2.427l1.286-.772A3.498 3.498 0 0 0 8 11.5M16 1c0-.711-.149-1.387-.416-2a5.016 5.016 0 0 1-2.726 2.643C11.511.612 9.828 0 8.001 0s-3.51.613-4.857 1.643A5.013 5.013 0 0 1 .418-1a4.979 4.979 0 0 0-.416 2c0 1.15.388 2.208 1.04 3.053a8 8 0 1 0 13.92 0A4.979 4.979 0 0 0 16.002 1zM8 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EvilTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1c0-.711-.149-1.387-.416-2a5.016 5.016 0 0 1-2.726 2.643C11.511.612 9.828 0 8.001 0s-3.51.613-4.857 1.643A5.013 5.013 0 0 1 .418-1a4.979 4.979 0 0 0-.416 2c0 1.15.388 2.208 1.04 3.053a8 8 0 1 0 13.92 0A4.979 4.979 0 0 0 16.002 1zM9.001 5.946c.032-.741.706-1.234 1.275-1.518c.543-.271 1.08-.407 1.102-.413a.5.5 0 1 1 .242.97c-.275.07-.602.189-.89.334A.998.998 0 0 1 9.998 7a1 1 0 0 1-1-1zM4.015 4.379a.5.5 0 0 1 .606-.364c.023.006.559.141 1.102.413c.568.284 1.243.776 1.275 1.518L7 6a1 1 0 1 1-1.732-.681a4.638 4.638 0 0 0-.89-.334a.5.5 0 0 1-.364-.606zM8 13a4.999 4.999 0 0 1-4.288-2.427l1.286-.772C5.61 10.819 6.725 11.5 8 11.5s2.389-.681 3.002-1.699l1.286.772A4.996 4.996 0 0 1 8 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exit(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10V8H7V6h5V4l3 3zm-1-1v4H6v3l-6-3V0h11v5h-1V1H2l4 2v9h4V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3C4.511 3 1.486 5.032 0 8c1.486 2.968 4.511 5 8 5s6.514-2.032 8-5c-1.486-2.968-4.511-5-8-5m3.945 2.652c.94.6 1.737 1.403 2.335 2.348a7.594 7.594 0 0 1-2.335 2.348a7.326 7.326 0 0 1-7.889 0A7.615 7.615 0 0 1 1.721 8a7.594 7.594 0 0 1 2.52-2.462a4 4 0 1 0 7.518 0c.062.037.124.075.185.114zM8 6.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 8 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeBlocked(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.78.22a.75.75 0 0 0-1.061 0L10.56 3.379a8.815 8.815 0 0 0-2.561-.378c-3.489 0-6.514 2.032-8 5a9.176 9.176 0 0 0 2.703 3.236L.218 13.721a.75.75 0 0 0 1.06 1.061l13.5-13.5a.75.75 0 0 0 0-1.061zM6.5 5a1.5 1.5 0 0 1 1.421 1.019L6.019 7.921A1.5 1.5 0 0 1 6.5 5M1.721 8a7.594 7.594 0 0 1 2.52-2.462A3.981 3.981 0 0 0 4 6.907c0 .858.27 1.652.73 2.303l-.952.952A7.625 7.625 0 0 1 1.721 8M12 6.906c0-.424-.066-.833-.189-1.217l-5.028 5.028A4 4 0 0 0 12 6.906"/><path fill="currentColor" d="m12.969 4.531l-1.084 1.084l.059.037c.94.6 1.737 1.403 2.335 2.348a7.594 7.594 0 0 1-2.335 2.348a7.326 7.326 0 0 1-5.725.933l-1.201 1.201A8.808 8.808 0 0 0 8 13c3.489 0 6.514-2.032 8-5a9.142 9.142 0 0 0-3.031-3.469"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeMinus(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2h6v2h-6z"/><path fill="currentColor" d="M13.599 5H9V3.056A8.923 8.923 0 0 0 8 3C4.511 3 1.486 5.032 0 8c1.486 2.968 4.511 5 8 5s6.514-2.032 8-5a9.173 9.173 0 0 0-2.401-3M6.5 5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 6.5 5m5.444 5.348a7.326 7.326 0 0 1-7.889 0A7.626 7.626 0 0 1 1.72 8a7.594 7.594 0 0 1 2.52-2.462a4 4 0 1 0 7.518 0A7.615 7.615 0 0 1 14.278 8a7.594 7.594 0 0 1-2.335 2.348z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyePlus(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2h-2V0h-2v2h-2v2h2v2h2V4h2z"/><path fill="currentColor" d="M13.498 6.969c.288.32.55.665.782 1.031a7.594 7.594 0 0 1-2.335 2.348a7.326 7.326 0 0 1-7.889 0A7.626 7.626 0 0 1 1.721 8a7.594 7.594 0 0 1 2.52-2.462A4 4 0 1 0 12 6.907v-.032a4.002 4.002 0 0 1-2.999-3.817A8.94 8.94 0 0 0 8 3.001c-3.489 0-6.514 2.032-8 5c1.486 2.968 4.511 5 8 5s6.514-2.032 8-5a9.217 9.217 0 0 0-.979-1.548a3.973 3.973 0 0 1-1.523.517zM6.5 5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 6.5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eyedropper(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.414.586a2 2 0 0 0-2.828 0L9.897 3.275L8.001 1.379L5.88 3.5l1.663 1.663L.166 12.54a.56.56 0 0 0-.161.46H.001v2.5a.5.5 0 0 0 .5.5h2.563a.561.561 0 0 0 .398-.165l7.377-7.377l1.663 1.663L14.623 8l-1.896-1.896l2.689-2.689a2 2 0 0 0 0-2.828zM2.705 15H1v-1.705l7.337-7.337l1.704 1.704l-7.337 7.337z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 3H12V0H9.5C7.57 0 6 1.57 6 3.5V5H4v3h2v8h3V8h2.5l.5-3H9V3.5c0-.271.229-.5.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5H8V9H6V7h2V6c0-1.653 1.347-3 3-3h2v2h-2c-.55 0-1 .45-1 1v1h3l-.5 2H10v7h4.5c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feed(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 8a2 2 0 1 1 3.999-.001A2 2 0 0 1 6 8m4.38-4.398a4.999 4.999 0 0 1 0 8.796c.689-1.096 1.12-2.66 1.12-4.398s-.431-3.302-1.12-4.398M4.5 8c0 1.738.431 3.302 1.12 4.398a4.999 4.999 0 0 1 0-8.796C4.931 4.698 4.5 6.262 4.5 8m-3 0c0 2.686.85 5.097 2.198 6.746C1.475 13.325 0 10.835 0 8s1.474-5.325 3.698-6.746C2.35 2.903 1.5 5.314 1.5 8m10.802-6.746C14.525 2.675 16 5.165 16 8s-1.474 5.325-3.698 6.746C13.65 13.097 14.5 10.686 14.5 8s-.85-5.097-2.198-6.746"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileEmpty(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExcel(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.61 6H9.496L8 8.204L6.504 6H4.39l2.534 3.788L4.065 14H8v-1.431h-.784L8 11.397L9.741 14h2.194L9.076 9.788z"/><path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMusic(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/><path fill="currentColor" d="M11.817 6.113a.501.501 0 0 0-.415-.104l-5 1a.5.5 0 0 0-.402.49V11.2a2.514 2.514 0 0 0-1-.201c-1.105 0-2 .672-2 1.5s.895 1.5 2 1.5s2-.672 2-1.5v-3.59l4-.8V10.2a2.514 2.514 0 0 0-1-.201c-1.105 0-2 .672-2 1.5s.895 1.5 2 1.5s2-.672 2-1.5v-5a.5.5 0 0 0-.183-.387z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileOpenoffice(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.785 7.37c-.948-.448-2.156-.538-3.044.095c1.08-.103 2.265.076 3.049.893c.75-.861 1.939-1.022 3.015-.933c-.898-.596-2.082-.516-3.019-.054zm-.384 2.095c-1.068-.025-2.101.362-2.986.939c-1.675-.712-3.793-.58-5.219.609c.411-.015.813-.116 1.22-.169c1.487-.148 3.072.221 4.196 1.247a4.52 4.52 0 0 1 1.87-1.561c.986-.477 2.096-.526 3.169-.539c-.651-.448-1.478-.531-2.249-.526z"/><path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdf(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.156 9.211c-.213-.21-.686-.321-1.406-.331a11.754 11.754 0 0 0-1.69.124c-.276-.159-.561-.333-.784-.542c-.601-.561-1.103-1.34-1.415-2.197c.02-.08.038-.15.054-.222c0 0 .339-1.923.249-2.573a.73.73 0 0 0-.044-.184l-.029-.076c-.092-.212-.273-.437-.556-.425l-.171-.005c-.316 0-.573.161-.64.403c-.205.757.007 1.889.39 3.355l-.098.239c-.275.67-.619 1.345-.923 1.94l-.04.077c-.32.626-.61 1.157-.873 1.607l-.271.144c-.02.01-.485.257-.594.323c-.926.553-1.539 1.18-1.641 1.678c-.032.159-.008.362.156.456l.263.132a.792.792 0 0 0 .357.086c.659 0 1.425-.821 2.48-2.662a24.79 24.79 0 0 1 3.819-.908c.926.521 2.065.883 2.783.883c.128 0 .238-.012.327-.036a.558.558 0 0 0 .325-.222c.139-.21.168-.499.13-.795a.531.531 0 0 0-.157-.271zM3.307 12.72c.12-.329.596-.979 1.3-1.556c.044-.036.153-.138.253-.233c-.736 1.174-1.229 1.642-1.553 1.788zm4.169-9.6c.212 0 .333.534.343 1.035s-.107.853-.252 1.113c-.12-.385-.179-.992-.179-1.389c0 0-.009-.759.088-.759M6.232 9.961c.148-.264.301-.543.458-.839c.383-.724.624-1.29.804-1.755a5.813 5.813 0 0 0 1.328 1.649c.065.055.135.111.207.166c-1.066.211-1.987.467-2.798.779zm6.72-.06c-.065.041-.251.064-.37.064c-.386 0-.864-.176-1.533-.464c.257-.019.493-.029.705-.029c.387 0 .502-.002.88.095s.383.293.318.333z"/><path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePicture(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 14H3v-2l3-5l4.109 5L13 10zm0-6.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 13 7.5"/><path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlay(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6 6l5 3.5L6 13z"/><path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileText(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 0h-12C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h12c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M13 14H2V2h11zM4 7h7v1H4zm0 2h7v1H4zm0 2h7v1H4zm0-6h7v1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTextTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/><path fill="currentColor" d="M11.5 13h-7a.5.5 0 0 1 0-1h7a.5.5 0 0 1 0 1m0-2h-7a.5.5 0 0 1 0-1h7a.5.5 0 0 1 0 1m0-2h-7a.5.5 0 0 1 0-1h7a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileVideo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/><path fill="currentColor" d="M4 8h5v5H4zm5 2l3-2v5l-3-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileWord(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.997 7.436h.691l-.797 3.534l-1.036-4.969H7.19L5.985 10.97l-.903-4.969H3.341l1.767 7.998h1.701l1.192-4.73l1.066 4.73h1.568l2.025-7.998H9.997z"/><path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileZip(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.341 3.579c-.347-.473-.831-1.027-1.362-1.558S11.894 1.006 11.421.659C10.615.068 10.224 0 10 0H2.25C1.561 0 1 .561 1 1.25v13.5c0 .689.561 1.25 1.25 1.25h11.5c.689 0 1.25-.561 1.25-1.25V5c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V1.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25H2.25a.253.253 0 0 1-.25-.25V1.25c0-.135.115-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/><path fill="currentColor" d="M4 1h2v1H4zm2 1h2v1H6zM4 3h2v1H4zm2 1h2v1H6zM4 5h2v1H4zm2 1h2v1H6zM4 7h2v1H4zm2 1h2v1H6zm-2 5.25c0 .412.338.75.75.75h2.5c.412 0 .75-.338.75-.75v-2.5a.753.753 0 0 0-.75-.75H6V9H4zM7 12v1H5v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilesEmpty(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.341 5.579c-.347-.473-.831-1.027-1.362-1.558s-1.085-1.015-1.558-1.362C10.615 2.068 10.224 2 10 2H4.25C3.561 2 3 2.561 3 3.25v11.5c0 .689.561 1.25 1.25 1.25h9.5c.689 0 1.25-.561 1.25-1.25V7c0-.224-.068-.615-.659-1.421m-2.07-.85c.48.48.856.912 1.134 1.271h-2.406V3.595c.359.278.792.654 1.271 1.134zM14 14.75c0 .136-.114.25-.25.25h-9.5a.253.253 0 0 1-.25-.25V3.25c0-.135.114-.25.25-.25H10v3.5a.5.5 0 0 0 .5.5H14z"/><path fill="currentColor" d="M9.421.659C8.615.068 8.224 0 8 0H2.25C1.561 0 1 .561 1 1.25v11.5c0 .604.43 1.109 1 1.225V1.25c0-.135.115-.25.25-.25h7.607A10.455 10.455 0 0 0 9.42.659z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v12h16V2zm3 11H1v-2h2zm0-4H1V7h2zm0-4H1V3h2zm9 8H4V3h8zm3 0h-2v-2h2zm0-4h-2V7h2zm0-4h-2V3h2zM6 5v6l4-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.582 0 0 1.119 0 2.5V4l6 6v5c0 .552.895 1 2 1s2-.448 2-1v-5l6-6V2.5C16 1.119 12.418 0 8 0M1.475 2.169c.374-.213.9-.416 1.52-.586C4.369 1.207 6.147 1 8 1s3.631.207 5.005.583c.62.17 1.146.372 1.52.586c.247.141.38.26.442.331c-.062.071-.195.19-.442.331c-.374.213-.9.416-1.52.586C11.631 3.793 9.853 4 8 4s-3.631-.207-5.005-.583c-.62-.17-1.146-.372-1.52-.586a1.741 1.741 0 0 1-.442-.331c.062-.071.195-.19.442-.331"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Finder(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.894 12.16v-.002zm.022.567l-.001-.011zm-.013-.285v-.008zM15 0H1C.45 0 0 .45 0 1v14c0 .55.45 1 1 1h14c.55 0 1-.45 1-1V1c0-.55-.45-1-1-1M3 3.5a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0zM15 15H9.14a20.425 20.425 0 0 1-.209-1.994v.005a8.551 8.551 0 0 1-6.595-2.09a.563.563 0 0 1 .744-.844a7.426 7.426 0 0 0 5.807 1.806a31.808 31.808 0 0 1 .11-3.334A.499.499 0 0 0 8.499 8H7.012c.022-.541.079-1.466.234-2.503c.295-1.981.812-3.528 1.502-4.497h6.251v14z"/><path fill="currentColor" d="M12.5 5a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5m-4.077 6.925l.037-.002zm-.219.009a.803.803 0 0 0 .05-.002zm5.509-1.806a.563.563 0 0 0-.794-.05a7.428 7.428 0 0 1-4.032 1.806c.007.364.02.742.043 1.127a8.552 8.552 0 0 0 4.733-2.09a.563.563 0 0 0 .05-.794zM8.68 13.035l-.183.013z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.016 16c-1.066-2.219-.498-3.49.321-4.688c.897-1.312 1.129-2.61 1.129-2.61s.706.917.423 2.352c1.246-1.387 1.482-3.598 1.293-4.445c2.817 1.969 4.021 6.232 2.399 9.392c8.631-4.883 2.147-12.19 1.018-13.013c.376.823.448 2.216-.313 2.893C9.999 1.002 6.818.002 6.818.002c.376 2.516-1.364 5.268-3.042 7.324c-.059-1.003-.122-1.696-.649-2.656c-.118 1.823-1.511 3.309-1.889 5.135c-.511 2.473.383 4.284 3.777 6.197z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Firefox(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.977 5.221l-.185 1.189s-.265-2.201-.59-3.024c-.498-1.261-.719-1.251-.72-1.249c.333.847.273 1.302.273 1.302s-.591-1.609-2.152-2.121C10.874.751 9.938.906 9.83.935h-.047l.038.003l-.001.001c.007.009 1.911.333 2.249.797c0 0-.809 0-1.614.232c-.036.01 2.961.374 3.574 3.37c0 0-.329-.686-.735-.802c.267.813.199 2.356-.056 3.123c-.033.099-.066-.426-.568-.652c.161 1.151-.01 2.976-.808 3.479c-.062.039.5-1.802.113-1.09c-2.23 3.419-4.866 1.578-6.051.767c.607.132 1.76-.021 2.271-.4l.002-.001c.554-.379.882-.656 1.177-.59s.491-.23.262-.493a1.527 1.527 0 0 0-1.539-.428c-.531.139-1.19.727-2.194.132c-.771-.457-.844-.837-.851-1.1c.019-.093.043-.18.072-.26c.089-.248.358-.323.508-.382c.254.044.473.123.703.241a4.049 4.049 0 0 0 0-.293c.022-.044.008-.176-.027-.337a2.364 2.364 0 0 0-.106-.48l.002-.001l.003-.003l.001-.001l.003-.007c.016-.072.188-.211.402-.361c.192-.134.417-.277.595-.387c.157-.098.277-.17.302-.189l.034-.026l.007-.006l.004-.004a.623.623 0 0 0 .237-.462v-.002l.002-.024l.001-.017l.001-.013l.001-.032v-.002c0-.026 0-.053-.002-.081a.268.268 0 0 0-.005-.043v-.002l-.001-.004a.025.025 0 0 0-.002-.007v-.001a.049.049 0 0 0-.003-.007c-.027-.064-.13-.088-.554-.096h-.001a36.33 36.33 0 0 0-.695-.002c-.52.002-.807-.508-.898-.705c.126-.695.489-1.19 1.085-1.525c.011-.006.009-.012-.004-.015c.117-.071-1.41-.002-2.112.891c-.623-.155-1.166-.144-1.635-.035a2.024 2.024 0 0 1-.335-.041c-.311-.282-.757-.803-.781-1.425l-.004.003l-.001-.018s-.949.729-.807 2.717l-.002.092c-.257.348-.384.641-.394.706C.418 5.1.188 5.797 0 6.855c0 0 .131-.417.395-.889C.201 6.56.049 7.484.138 8.87c0 0 .024-.307.107-.75c.065.86.352 1.921 1.076 3.169c1.39 2.396 3.526 3.605 5.887 3.791a8.06 8.06 0 0 0 1.272.003l.118-.009a8.877 8.877 0 0 0 1.457-.224c6.643-1.606 5.921-9.628 5.921-9.628z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func First(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 14V2h2v5.5l5-5v5l5-5v11l-5-5v5l-5-5V14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveHundredPx(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.953 10.512a5.24 5.24 0 0 0 6.996 3.141c.625-.262 1.184-.641 1.666-1.122s.859-1.041 1.122-1.666c.272-.647.412-1.331.412-2.037s-.137-1.394-.412-2.037c-.262-.625-.641-1.184-1.122-1.666s-1.041-.859-1.666-1.122a5.226 5.226 0 0 0-2.037-.413c-.716 0-1.431.144-2.066.413c-.509.216-1.372.769-1.875 1.291l-.003.003V.984h7.241c.262-.003.262-.372.262-.491c0-.122 0-.487-.266-.491H4.377a.343.343 0 0 0-.344.341v6.066c0 .197.244.338.472.384c.444.094.544-.047.653-.197l.016-.019c.166-.247.681-.766.688-.772a4.262 4.262 0 0 1 3.037-1.25c1.147 0 2.222.444 3.028 1.25a4.245 4.245 0 0 1 1.256 3.019a4.236 4.236 0 0 1-1.25 3.019a4.336 4.336 0 0 1-3.047 1.25a4.136 4.136 0 0 1-2.159-.597l.003-3.688c0-.491.213-1.028.572-1.431a2.09 2.09 0 0 1 1.588-.716c.594 0 1.15.225 1.566.634c.409.406.637.95.637 1.528a2.179 2.179 0 0 1-2.206 2.197c-.238 0-.672-.106-.691-.109c-.25-.075-.356.272-.391.387c-.134.441.069.528.109.541c.397.125.659.147 1.003.147a3.173 3.173 0 0 0 3.169-3.169c0-1.734-1.422-3.144-3.166-3.144c-.856 0-1.659.328-2.263.919c-.575.566-.903 1.319-.903 2.069v.019c-.003.094-.003 2.306-.006 3.031l-.003-.003c-.328-.363-.653-.919-.869-1.488c-.084-.222-.275-.184-.534-.103c-.125.034-.469.141-.391.394zm3.722-.865c0 .106.097.2.156.253l.019.019c.1.097.194.147.281.147a.181.181 0 0 0 .131-.05c.044-.041.537-.544.588-.591l.553.55c.05.056.106.088.172.088c.088 0 .184-.053.284-.156c.238-.244.119-.375.063-.438l-.559-.559l.584-.588c.128-.137.016-.284-.097-.397c-.162-.162-.322-.206-.422-.112l-.581.581l-.588-.588a.16.16 0 0 0-.113-.047c-.078 0-.172.053-.275.156c-.181.181-.219.306-.125.406l.588.584l-.584.584c-.053.05-.078.103-.075.156zm1.278-7.931c-.938 0-1.938.191-2.669.506a.207.207 0 0 0-.134.181a.753.753 0 0 0 .069.337c.047.116.166.425.4.334a6.689 6.689 0 0 1 2.334-.444a6.35 6.35 0 0 1 2.469.497c.622.263 1.206.644 1.844 1.194a.22.22 0 0 0 .147.059c.125 0 .244-.122.347-.237c.169-.191.287-.35.119-.509a6.858 6.858 0 0 0-2.1-1.356a7.326 7.326 0 0 0-2.825-.563zM14.006 13.3c-.113-.113-.209-.178-.294-.203s-.162-.006-.222.053l-.056.056a6.32 6.32 0 0 1-6.938 1.356a6.336 6.336 0 0 1-2.013-1.356a6.046 6.046 0 0 1-1.356-2.012c-.288-.713-.381-1.247-.413-1.422c-.003-.016-.006-.028-.006-.037c-.041-.206-.231-.222-.503-.178c-.112.019-.459.072-.428.319v.006a7.261 7.261 0 0 0 2.04 3.994a7.266 7.266 0 0 0 10.288 0l.059-.059c.069-.084.134-.225-.159-.516z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h2v16H0zm13 10.047c1.291 0 2.415-.312 3-.773v-8c-.585.461-1.709.773-3 .773s-2.415-.312-3-.773v8c.585.461 1.709.773 3 .773M9.5.508C8.767.196 7.695 0 6.5 0C4.994 0 3.682.312 3 .773v8C3.682 8.312 4.994 8 6.5 8c1.195 0 2.267.197 3 .508z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flattr(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.743 0C1.941 0 0 2.19 0 6.279v8.579l3.725-3.729V6.771c0-1.694.449-2.772 1.955-3.014c.526-.103 1.621-.067 2.317-.067v2.587a.247.247 0 0 0 .245.269c.063 0 .123-.033.184-.093L14.881 0L5.742-.001zm6.532 4.871v4.358c0 1.694-.449 2.772-1.955 3.014c-.526.103-1.621.067-2.317.067V9.723a.414.414 0 0 0-.009-.087a.246.246 0 0 0-.236-.182c-.064 0-.123.033-.184.093L1.119 16l9.139.001c3.802 0 5.743-2.19 5.743-6.279V1.143l-3.725 3.729z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0m9 0a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlickrFour(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.582 0 0 3.606 0 8.055s3.582 8.055 8 8.055s8-3.606 8-8.055S12.418 0 8 0M4.5 10.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5m7 0a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlickrThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5m-10 10.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5m7 0a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlickrTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 6.5c-1.103 0-2 .897-2 2s.897 2 2 2s2-.897 2-2s-.897-2-2-2m0-1.5a3.5 3.5 0 1 1 0 7a3.5 3.5 0 0 1 0-7M0 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyDisk(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 0H0v16h16V2zM8 2h2v4H8zm6 12H2V2h1v5h9V2h1.172l.828.828z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7 2l2 2h7v11H0V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderDownload(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 4L7 2H0v13h16V4zm-1 9.5L4.5 10H7V6h2v4h2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinus(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 4L7 2H0v13h16V4zm2 7H5V9h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 15l3-8H3l-3 8zM2 6l-2 9V2h4.5l2 2H13v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 4L7 2H0v13h16V4zm2 7H9v2H7v-2H5V9h2V7h2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderUpload(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 4L7 2H0v13h16V4zM8 7.5l3.5 3.5H9v4H7v-4H4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.494.253C11.08.253 10.172 0 8.715 0C4.007 0 1.812 2.681 1.812 5.404c0 1.604.76 2.132 2.259 2.132c-.106-.232-.296-.486-.296-1.626c0-3.188 1.203-4.117 2.744-4.18c0 0-1.264 12.396-4.934 13.883v.385h4.947l1.688-8h3.091l.689-2H8.642l.812-3.847c.929.19 1.837.38 2.618.38c.971 0 1.858-.296 2.343-2.533c-.591.19-1.224.253-1.921.253z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontSize(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8h6v2H5v6H3v-6H1zm14-4h-3.934v12H8.933V4H4.999V2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.096 0C2.319 3.219 2.02 8.13 9 7.966V4l6 6l-6 6v-3.881C.641 12.337-.29 4.741 4.096 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 13.5v-5l-5 5v-11l5 5v-5L13.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13m-3-9L8.5 8L5 10.5zm4 0L12.5 8L9 10.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foursquare(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.306 1.408A1.001 1.001 0 0 0 12.5 1H3a1 1 0 0 0-1 1v12a1.002 1.002 0 0 0 .999 1a1 1 0 0 0 .707-.293L7.413 11h2.586a.999.999 0 0 0 .954-.702l2.5-8a.999.999 0 0 0-.149-.891zM10.515 5H7a1 1 0 0 0 0 2h2.89l-.625 2H7a.997.997 0 0 0-.707.293L4 11.586V3h7.14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frustrated(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.724 4.428c-.543-.271-1.08-.407-1.102-.413a.5.5 0 1 0-.242.97c.275.07.602.189.89.334A.998.998 0 0 0 6.002 7a1 1 0 0 0 1-1L7 5.946c-.032-.741-.706-1.234-1.275-1.518z"/><path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M3.695 12.87c.167.083.356.13.555.13h7.5c.199 0 .387-.047.555-.13C11.158 13.884 9.651 14.5 8 14.5s-3.158-.616-4.305-1.63M4 11.75v-1.5c0-.136.114-.25.25-.25H6v2H4.25a.253.253 0 0 1-.25-.25M7 12v-2h2v2zm3 0v-2h1.75c.136 0 .25.114.25.25v1.5c0 .136-.114.25-.25.25zm2.87.305c.083-.167.13-.356.13-.555v-1.5C13 9.561 12.439 9 11.75 9h-7.5C3.561 9 3 9.561 3 10.25v1.5c0 .199.047.387.13.555a6.5 6.5 0 1 1 9.74 0"/><path fill="currentColor" d="M11.379 4.015c-.023.006-.559.141-1.102.413c-.568.284-1.243.776-1.275 1.518L9 6a1 1 0 1 0 1.732-.681c.288-.144.614-.264.89-.334a.5.5 0 1 0-.242-.97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrustratedTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 10.25v1.5c0 .136.114.25.25.25H6v-2H4.25a.253.253 0 0 0-.25.25M7 10h2v2H7zm4.75 0H10v2h1.75c.136 0 .25-.114.25-.25v-1.5a.253.253 0 0 0-.25-.25"/><path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m1.002 5.946c.032-.741.706-1.234 1.275-1.518c.543-.271 1.08-.407 1.102-.413a.5.5 0 1 1 .242.97c-.275.07-.602.189-.89.334A.998.998 0 0 1 9.999 7a1 1 0 0 1-1-1zM4.015 4.379a.5.5 0 0 1 .606-.364c.023.006.559.141 1.102.413c.568.284 1.243.776 1.275 1.518L7 6a1 1 0 1 1-1.732-.681a4.638 4.638 0 0 0-.89-.334a.5.5 0 0 1-.364-.606zM13 11.75c0 .689-.561 1.25-1.25 1.25h-7.5C3.561 13 3 12.439 3 11.75v-1.5C3 9.561 3.561 9 4.25 9h7.5c.689 0 1.25.561 1.25 1.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.055 5a6.27 6.27 0 0 0 .804-.674c.522-.522.871-1.164.983-1.808c.123-.706-.057-1.362-.494-1.798c-.348-.348-.82-.533-1.365-.533c-.775 0-1.593.372-2.242 1.021c-1.039 1.039-1.644 2.472-1.97 3.496c-.241-1.028-.722-2.416-1.657-3.351C5.613.852 4.972.594 4.366.594c-.495 0-.965.172-1.317.523c-.781.781-.675 2.153.236 3.064c.325.325.705.595 1.105.819H.999v4h1v7h12V9h1V5zm-1.519-2.997c.433-.433.974-.692 1.446-.692c.167 0 .402.035.57.203c.407.407.178 1.349-.489 2.016c-.687.687-1.61 1.159-2.413 1.47h-.792c.29-.899.813-2.132 1.678-2.997m-6.881.511c-.011-.143-.001-.41.191-.601a.727.727 0 0 1 .521-.194c.332 0 .679.157.952.429c.529.529.965 1.371 1.26 2.436l.023.086l-.086-.023c-1.064-.295-1.906-.731-2.436-1.26a1.42 1.42 0 0 1-.426-.872zM7 15H3V8.5h4zm0-7H2V6h5zm6 7H9V8.5h4zm1-7H9V6h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Git(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.698 7.287L8.712.301a1.03 1.03 0 0 0-1.457 0L5.632 1.924l1.221 1.221a1.5 1.5 0 0 1 2.001 2.001l2 2a1.5 1.5 0 1 1-.707.707l-2-2A1.569 1.569 0 0 1 8 5.914v4.171a1.5 1.5 0 1 1-1 0V5.914a1.5 1.5 0 0 1-.854-2.061L4.925 2.632L.302 7.255a1.032 1.032 0 0 0 0 1.458l6.986 6.986a1.03 1.03 0 0 0 1.457 0l6.953-6.953a1.032 1.032 0 0 0 0-1.458z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .198a8 8 0 0 0-2.529 15.591c.4.074.547-.174.547-.385c0-.191-.008-.821-.011-1.489c-2.226.484-2.695-.944-2.695-.944c-.364-.925-.888-1.171-.888-1.171c-.726-.497.055-.486.055-.486c.803.056 1.226.824 1.226.824c.714 1.223 1.872.869 2.328.665c.072-.517.279-.87.508-1.07c-1.777-.202-3.645-.888-3.645-3.954c0-.873.313-1.587.824-2.147c-.083-.202-.357-1.015.077-2.117c0 0 .672-.215 2.201.82A7.672 7.672 0 0 1 8 4.066c.68.003 1.365.092 2.004.269c1.527-1.035 2.198-.82 2.198-.82c.435 1.102.162 1.916.079 2.117c.513.56.823 1.274.823 2.147c0 3.073-1.872 3.749-3.653 3.947c.287.248.543.735.543 1.481c0 1.07-.009 1.932-.009 2.195c0 .213.144.462.55.384A8 8 0 0 0 8.001.196z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glass(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.153.263a.5.5 0 0 0-.44-.263H4.288a.5.5 0 0 0-.44.263C3.294 1.295 3.001 2.5 3.001 3.75c0 1.647.506 3.2 1.424 4.374c.71.907 1.601 1.508 2.576 1.753V15h-1.5a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-1.5V9.877c.975-.244 1.866-.846 2.576-1.753c.918-1.174 1.424-2.727 1.424-4.374c0-1.249-.293-2.455-.847-3.487zM4.595 1h6.809a6.46 6.46 0 0 1 .59 3H4.003a6.46 6.46 0 0 1 .59-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.893 2.809a.499.499 0 0 0-.393-.808h-11a.501.501 0 0 0-.393.808L7 9.037V15H5.5a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1H9V9.037zM12.471 3L10.9 5H5.1L3.529 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.159 6.856V9.6h4.537c-.184 1.178-1.372 3.45-4.537 3.45C5.428 13.05 3.2 10.788 3.2 8s2.228-5.05 4.959-5.05c1.553 0 2.594.663 3.188 1.234l2.172-2.091C12.125.787 10.319-.001 8.16-.001c-4.422 0-8 3.578-8 8s3.578 8 8 8c4.616 0 7.681-3.247 7.681-7.816c0-.525-.056-.925-.125-1.325z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDrive(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.844 10L3.96 15h9.072l2.884-5zm8.662-1l-4.619-8H5.112l4.619 8zM4.534 2L0 9.856l2.888 5L7.422 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlus(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.091 7.147v1.747h2.888c-.116.75-.872 2.197-2.888 2.197c-1.737 0-3.156-1.441-3.156-3.216s1.419-3.216 3.156-3.216c.991 0 1.65.422 2.028.784L8.5 4.112c-.888-.828-2.037-1.331-3.409-1.331C2.275 2.784 0 5.059 0 7.875s2.275 5.091 5.091 5.091c2.937 0 4.888-2.066 4.888-4.975c0-.334-.037-.591-.081-.844zM16 7h-1.5V5.5H13V7h-1.5v1.5H13V10h1.5V8.5H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.581 0 0 3.581 0 8s3.581 8 8 8s8-3.581 8-8s-3.581-8-8-8M6 12c-2.212 0-4-1.787-4-4s1.788-4 4-4c1.081 0 1.984.394 2.681 1.047L7.593 6.091c-.297-.284-.816-.616-1.594-.616c-1.366 0-2.481 1.131-2.481 2.525s1.116 2.525 2.481 2.525c1.584 0 2.178-1.137 2.269-1.725H5.999V7.428h3.778c.034.2.063.4.063.663C9.84 10.378 8.309 12 5.999 12zm7-4v1h-1V8h-1V7h1V6h1v1h1v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M6 12c-2.212 0-4-1.787-4-4s1.788-4 4-4c1.081 0 1.984.394 2.681 1.047L7.593 6.091c-.297-.284-.816-.616-1.594-.616c-1.366 0-2.481 1.131-2.481 2.525s1.116 2.525 2.481 2.525c1.584 0 2.178-1.137 2.269-1.725H5.999V7.428h3.778c.034.2.063.4.063.663C9.84 10.378 8.309 12 5.999 12zm8-4h-1v1h-1V8h-1V7h1V6h1v1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.581 0 0 3.581 0 8s3.581 8 8 8s8-3.581 8-8s-3.581-8-8-8m.119 14c-3.316 0-6-2.684-6-6s2.684-6 6-6c1.619 0 2.975.591 4.019 1.569L10.51 5.138c-.447-.428-1.225-.925-2.391-.925C6.069 4.213 4.4 5.91 4.4 8s1.672 3.787 3.719 3.787c2.375 0 3.266-1.706 3.403-2.588H8.119V7.143h5.666c.05.3.094.6.094.994C13.882 11.565 11.585 14 8.12 14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M8.119 14c-3.316 0-6-2.684-6-6s2.684-6 6-6c1.619 0 2.975.591 4.019 1.569L10.51 5.138c-.447-.428-1.225-.925-2.391-.925C6.069 4.213 4.4 5.91 4.4 8s1.672 3.787 3.719 3.787c2.375 0 3.266-1.706 3.403-2.588H8.119V7.143h5.666c.05.3.094.6.094.994C13.882 11.565 11.585 14 8.12 14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grin(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13M3 8v1c0 2.2 1.8 4 4 4h2c2.2 0 4-1.8 4-4V8zm3 3.828a3.008 3.008 0 0 1-1.118-.71A2.978 2.978 0 0 1 4 9h2zM9 12H7V9h2zm2.118-.882a3.008 3.008 0 0 1-1.118.71V9h2c0 .797-.313 1.549-.882 2.118M3.521 6c.153 0 .283-.11.308-.261c.096-.573.589-.989 1.171-.989s1.074.416 1.171.989a.312.312 0 0 0 .616 0a1.815 1.815 0 0 0-1.788-2.115a1.815 1.815 0 0 0-1.788 2.115a.312.312 0 0 0 .308.261zm6 0c.153 0 .283-.11.308-.261c.096-.573.589-.989 1.171-.989s1.074.416 1.171.989a.312.312 0 0 0 .616 0a1.815 1.815 0 0 0-1.788-2.115a1.815 1.815 0 0 0-1.788 2.115a.312.312 0 0 0 .308.261z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GrinTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m3 3.688a1.815 1.815 0 0 1 1.788 2.115a.312.312 0 0 1-.616 0c-.096-.573-.589-.833-1.171-.833s-1.074.26-1.171.833a.312.312 0 0 1-.616 0a1.815 1.815 0 0 1 1.788-2.115zm-6 0a1.815 1.815 0 0 1 1.788 2.115a.312.312 0 0 1-.616 0c-.096-.573-.589-.833-1.171-.833s-1.074.26-1.171.833a.312.312 0 0 1-.616 0a1.815 1.815 0 0 1 1.788-2.115zM3 9h3v3.873A4.017 4.017 0 0 1 3 9m4 4V9h2v4zm3-.127V9h3a4.017 4.017 0 0 1-3 3.873"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hackernews(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm8.5 9.125V12.5h-1V9.125L4.766 4H5.9L8 7.938L10.1 4h1.134z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hammer(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.781 12.953l-4.712-4.712a.752.752 0 0 0-1.061 0l-.354.354L6.779 5.72L11.499 1h-5l-2.22 2.22l-.22-.22H2.998v1.061l.22.22l-3.22 3.22l2.5 2.5l3.22-3.22l2.875 2.875l-.354.354a.752.752 0 0 0 0 1.061l4.712 4.712a.752.752 0 0 0 1.061 0l1.768-1.768a.752.752 0 0 0 0-1.061z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HammerTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.784 14.309L7.212 6.505l.399-.4a1.83 1.83 0 0 0 .53-1.181a.395.395 0 0 0 .046-.023l1.609-1.006a.676.676 0 0 0-.036-.898L6.961.191a.673.673 0 0 0-.896-.036L5.061 1.769l-.022.046c-.43.027-.852.204-1.178.531L2.339 3.873c-.327.327-.503.75-.53 1.181a.395.395 0 0 0-.046.023L.154 6.083a.676.676 0 0 0 .036.898l2.799 2.806a.673.673 0 0 0 .896.036l1.004-1.614l.023-.046c.43-.027.852-.204 1.178-.531l.442-.443l7.783 8.596c.226.249.573.289.773.089l.787-.789c.199-.2.159-.549-.089-.775z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hangouts(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.997 0a6.91 6.91 0 0 0-6.909 6.909c0 3.616 3.294 6.547 6.909 6.547V16c4.197-2.128 6.916-5.556 6.916-9.091c0-3.816-3.1-6.909-6.916-6.909M7 8c0 .828-.447 1.5-1 1.5V8H4V5h3zm5 0c0 .828-.447 1.5-1 1.5V8H9V5h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Happy(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13m0 7.856c1.812 0 3.535-.481 5-1.327C12.772 10.817 10.607 13 8 13s-4.772-2.186-5-4.973a10.017 10.017 0 0 0 5 1.329M4 5.5C4 4.672 4.448 4 5 4s1 .672 1 1.5S5.552 7 5 7s-1-.672-1-1.5m6 0c0-.828.448-1.5 1-1.5s1 .672 1 1.5S11.552 7 11 7s-1-.672-1-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HappyTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m3 4c.552 0 1 .672 1 1.5S11.552 7 11 7s-1-.672-1-1.5s.448-1.5 1-1.5M5 4c.552 0 1 .672 1 1.5S5.552 7 5 7s-1-.672-1-1.5S4.448 4 5 4m3 10c-2.607 0-4.772-2.186-5-4.973c1.465.846 3.188 1.329 5 1.329s3.535-.481 5-1.327C12.772 11.817 10.607 14 8 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 9h-1v7h1c.275 0 .5-.225.5-.5v-6c0-.275-.225-.5-.5-.5m7 0c-.275 0-.5.225-.5.5v6c0 .275.225.5.5.5h1V9z"/><path fill="currentColor" d="M16 8A8 8 0 1 0 .479 10.732A3.5 3.5 0 0 0 3 15.964V9.036a3.478 3.478 0 0 0-1.371.506a6.5 6.5 0 1 1 12.743 0A3.484 3.484 0 0 0 13 9.035v6.929a3.5 3.5 0 0 0 2.521-5.232C15.831 9.879 16 8.959 16 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.8 1c-1.682 0-3.129 1.368-3.799 2.797C7.33 2.368 5.883 1 4.201 1a4.202 4.202 0 0 0-4.2 4.2c0 4.716 4.758 5.953 8 10.616c3.065-4.634 8-6.05 8-10.616c0-2.319-1.882-4.2-4.2-4.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartBroken(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.8 1C14.118 1 16 2.882 16 5.2c0 4.566-4.935 5.982-8 10.616c-3.243-4.663-8-5.9-8-10.616C0 2.881 1.882 1 4.2 1c.943 0 1.812.43 2.512 1.06L5.499 4l3.5 2l-2 5l5.5-6l-3.5-2l.967-1.451c.553-.34 1.175-.549 1.833-.549z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hipster(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13M4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/><path fill="currentColor" d="M10.561 8.439a1.5 1.5 0 1 0-2.063 2.176c1.352 1.227 4.503-.029 4.503-1.615c-.969.625-1.726.153-2.439-.561z"/><path fill="currentColor" d="M5.439 8.439a1.5 1.5 0 1 1 2.063 2.176C6.15 11.842 2.999 10.586 2.999 9c.969.625 1.726.153 2.439-.561z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HipsterTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m3 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2M5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2m3.497 6.615A1.507 1.507 0 0 1 8 9.5a1.491 1.491 0 0 1-.497 1.115C6.151 11.842 3 10.586 3 9c.969.625 1.726.153 2.439-.561a1.5 1.5 0 0 1 2.56 1.06a1.5 1.5 0 0 1 2.56-1.06c.713.714 1.471 1.186 2.439.561c0 1.586-3.151 2.842-4.503 1.615z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 1a7 7 0 1 1 0 14v-1.5c1.469 0 2.85-.572 3.889-1.611S15.5 9.469 15.5 8c0-1.469-.572-2.85-1.611-3.889S11.469 2.5 10 2.5c-1.469 0-2.85.572-3.889 1.611A5.455 5.455 0 0 0 4.591 7H7.5L4 11L.5 7h2.571A7.001 7.001 0 0 1 10 1m3 6v2H9V4h2v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 9.226l-8-6.21l-8 6.21V6.694l8-6.21l8 6.21zM14 9v6h-4v-4H6v4H2V9l6-4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 9.5l-3-3V2h-2v2.5l-3-3l-8 8v.5h2v5h5v-3h2v3h5v-5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 .5l-8 8L1.5 10L3 8.5V15h4v-3h2v3h4V8.5l1.5 1.5L16 8.5zM8 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourGlass(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.39 8C13.542 6.635 15 4.012 15 1c0-.339-.019-.672-.054-1H1.055a9.155 9.155 0 0 0-.054 1c0 3.012 1.457 5.635 3.609 7c-2.152 1.365-3.609 3.988-3.609 7c0 .339.019.672.054 1h13.891c.036-.328.054-.661.054-1c0-3.012-1.457-5.635-3.609-7zM2.5 15c0-2.921 1.253-5.397 3.5-6.214V7.214C3.753 6.397 2.5 3.92 2.5 1h11c0 2.921-1.253 5.397-3.5 6.214v1.572c2.247.817 3.5 3.294 3.5 6.214zm7.182-4.538C8.562 9.827 8.501 9.003 8.5 8.503V7.499c0-.5.059-1.327 1.184-1.963c.602-.349 1.122-.88 1.516-1.537H4.8c.395.657.916 1.188 1.518 1.538c1.12.635 1.181 1.459 1.182 1.959V8.5c0 .5-.059 1.327-1.184 1.963c-1.135.659-1.98 1.964-2.236 3.537h7.839c-.256-1.574-1.102-2.879-2.238-3.538z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.946 0L2.23 14.4L7.992 16l5.777-1.602L15.055 0zM12.26 4.71H5.502l.161 1.809H12.1l-.485 5.422l-3.623 1.004l-3.618-1.004l-.247-2.774H5.9l.126 1.41l1.967.53l.004-.001l1.968-.531l.204-2.29H4.048l-.476-5.341h8.847l-.158 1.766z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFiveTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.946 0L2.23 14.4L7.992 16l5.777-1.602L15.055 0zm11.722 13.482l-4.644 1.287v.007l-.012-.004l-.012.004v-.007l-4.644-1.287L2.258 1.178h11.508zm-2.5-5.198l-.204 2.29l-1.972.532l-1.967-.53l-.126-1.41H4.126l.247 2.774l3.626 1.003l3.615-1.003l.485-5.422H5.662l-.161-1.809h6.758l.158-1.766H3.57l.477 5.341z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Icomoon(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.055 8a1.851 1.851 0 1 1 3.703 0a1.851 1.851 0 0 1-3.703 0M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M5.928 14.989C3.522 13.589 1.905 10.984 1.905 8s1.617-5.589 4.023-6.989C8.334 2.41 9.953 5.016 9.953 8s-1.618 5.589-4.025 6.989"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ie(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.472 9.825h3.688c.028-.256.04-.517.04-.784a6.835 6.835 0 0 0-.924-3.442c.607-1.614.586-2.984-.227-3.803c-.773-.77-2.848-.645-5.194.394a6.876 6.876 0 0 0-7.193 5.181c1.01-1.293 2.072-2.231 3.492-2.913c-.129.121-.882.87-1.009.996c-3.743 3.742-4.923 8.63-3.653 9.9c.965.965 2.715.802 4.725-.182a6.836 6.836 0 0 0 3.113.744a6.869 6.869 0 0 0 6.501-4.648h-3.717a3.026 3.026 0 0 1-5.32 0a3.042 3.042 0 0 1-.358-1.432v-.011zm-6.03-1.812c.085-1.517 1.347-2.728 2.887-2.728s2.802 1.21 2.887 2.728zm8.573-5.454c.524.529.511 1.503.063 2.719a6.895 6.895 0 0 0-3.2-2.619c1.408-.604 2.553-.684 3.137-.1M1.461 15.113c-.668-.669-.467-2.072.394-3.763a6.902 6.902 0 0 0 2.927 3.581c-1.491.677-2.712.792-3.321.182"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.998 2l.002.002v11.996l-.002.002H1.002L1 13.998V2.002L1.002 2zM15 1H1c-.55 0-1 .45-1 1v12c0 .55.45 1 1 1h14c.55 0 1-.45 1-1V2c0-.55-.45-1-1-1"/><path fill="currentColor" d="M13 4.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 13 4.5m1 8.5H2v-2l3.5-6l4 5h1L14 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Images(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-1V1c0-.55-.45-1-1-1H1C.45 0 0 .45 0 1v12c0 .55.45 1 1 1h1v1c0 .55.45 1 1 1h14c.55 0 1-.45 1-1V3c0-.55-.45-1-1-1M2 3v10h-.998L1 12.998V1.002L1.002 1h13.996l.002.002V2H3c-.55 0-1 .45-1 1m15 11.998l-.002.002H3.002L3 14.998V3.002L3.002 3h13.996l.002.002z"/><path fill="currentColor" d="M15 5.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 15 5.5m1 8.5H4v-2l3.5-6l4 5h1L16 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentDecrease(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h16v2H0zm6 3h10v2H6zm0 3h10v2H6zm0 3h10v2H6zm-6 3h16v2H0zm4-8v6L0 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentIncrease(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h16v2H0zm6 3h10v2H6zm0 3h10v2H6zm0 3h10v2H6zm-6 3h16v2H0zm0-2V5l4 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinite(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.25 11.75a3.728 3.728 0 0 1-2.652-1.098L8 9.054l-1.598 1.598c-.708.708-1.65 1.098-2.652 1.098s-1.944-.39-2.652-1.098C.39 9.944 0 9.002 0 8s.39-1.943 1.098-2.652C1.806 4.64 2.748 4.25 3.75 4.25s1.943.39 2.652 1.098L8 6.946l1.598-1.598c.708-.708 1.65-1.098 2.652-1.098s1.944.39 2.652 1.098C15.61 6.056 16 6.998 16 8s-.39 1.943-1.098 2.652a3.726 3.726 0 0 1-2.652 1.098m-1.598-2.152c.427.427.994.662 1.598.662s1.171-.235 1.598-.662c.427-.427.662-.994.662-1.598s-.235-1.171-.662-1.598c-.427-.427-.994-.662-1.598-.662s-1.171.235-1.598.662L9.054 8zM3.75 5.74c-.604 0-1.171.235-1.598.662S1.49 7.396 1.49 8c0 .604.235 1.171.662 1.598s.994.662 1.598.662c.604 0 1.171-.235 1.598-.662L6.946 8L5.348 6.402A2.244 2.244 0 0 0 3.75 5.74"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 4.75c0-.412.338-.75.75-.75h.5c.412 0 .75.338.75.75v.5c0 .412-.338.75-.75.75h-.5A.753.753 0 0 1 7 5.25zM10 12H6v-1h1V8H6V7h3v4h1z"/><path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertTemplate(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3h2v1H6zm3 0h2v1H9zm5 0v4h-3V6h2V4h-1V3zM5 6h2v1H5zm3 0h2v1H8zM3 4v2h1v1H2V3h3v1zm3 5h2v1H6zm3 0h2v1H9zm5 0v4h-3v-1h2v-2h-1V9zm-9 3h2v1H5zm3 0h2v1H8zm-5-2v2h1v1H2V9h3v1zm12-9H1v14h14zm1-1v16H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M11 2.5c0-.275.225-.5.5-.5h2c.275 0 .5.225.5.5v2c0 .275-.225.5-.5.5h-2a.501.501 0 0 1-.5-.5zM8 5a3.001 3.001 0 0 1 0 6a3.001 3.001 0 0 1 0-6m6 8.5c0 .275-.225.5-.5.5h-11a.501.501 0 0 1-.5-.5V7h1.1A4.999 4.999 0 0 0 8 13a4.999 4.999 0 0 0 4.9-6H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1v1h-2L7 14h2v1H2v-1h2L9 2H7V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Joomla(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.156 4.323a1.312 1.312 0 0 1 1.856 0l.122.123l1.58-1.581l-.123-.123a3.538 3.538 0 0 0-3.319-.946A2.142 2.142 0 0 0 2.16 0a2.143 2.143 0 0 0-.473 4.234a3.552 3.552 0 0 0 .888 3.531l3.56 3.561l1.578-1.581l-3.56-3.561a1.318 1.318 0 0 1 .003-1.861m11.824-2.18a2.141 2.141 0 0 0-4.26-.296a3.543 3.543 0 0 0-3.51.897L4.65 6.305l1.58 1.581l3.559-3.56a1.31 1.31 0 0 1 1.854-.003a1.316 1.316 0 0 1-.001 1.859l-.122.122l1.578 1.582l.123-.124a3.55 3.55 0 0 0 .9-3.494a2.142 2.142 0 0 0 1.858-2.125zm-1.82 9.592a3.55 3.55 0 0 0-.939-3.352L9.666 4.821l-1.58 1.58l3.555 3.563a1.313 1.313 0 1 1-1.854 1.857l-.121-.122l-1.578 1.582l.121.121a3.544 3.544 0 0 0 3.553.883A2.14 2.14 0 0 0 16 13.858c0-1.081-.8-1.976-1.84-2.122zM9.568 8.261l-3.555 3.562a1.316 1.316 0 0 1-1.86-1.861l.122-.121l-1.579-1.58l-.121.12a3.55 3.55 0 0 0-.929 3.39A2.142 2.142 0 0 0 0 13.857a2.141 2.141 0 0 0 4.227.481a3.539 3.539 0 0 0 3.365-.934l3.555-3.562L9.569 8.26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 0a5 5 0 0 0-4.916 5.916L0 12v3a1 1 0 0 0 1 1h1v-1h2v-2h2v-2h2l1.298-1.298A5 5 0 1 0 11 0m1.498 5.002a1.5 1.5 0 1 1 .001-3.001a1.5 1.5 0 0 1-.001 3.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.658 4.91l-1.58-1.58l-1.409-1.409l-1.58-1.58c-.387-.387-1.077-.456-1.533-.152l-4.319 2.88a1.162 1.162 0 0 0-.383 1.444l1.101 2.203c.034.067.073.139.115.213L.499 12.5l-.5 3.5h3v-1h2v-2h2v-2h2V9.888c.1.06.196.113.284.157l2.203 1.101c.49.245 1.14.072 1.444-.383l2.88-4.319c.304-.456.236-1.146-.152-1.533zM2.354 13.354l-.707-.707l4.868-4.868l.707.707zm11.974-6.733l-.707.707a.502.502 0 0 1-.707 0L8.671 3.085a.502.502 0 0 1 0-.707l.707-.707a.502.502 0 0 1 .707 0l4.243 4.243a.502.502 0 0 1 0 .707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H1c-.55 0-1 .45-1 1v10c0 .55.45 1 1 1h16c.55 0 1-.45 1-1V3c0-.55-.45-1-1-1m-7 2h2v2h-2zm3 3v2h-2V7zM7 4h2v2H7zm3 3v2H8V7zM4 4h2v2H4zm3 3v2H5V7zM2 4h1v2H2zm0 3h2v2H2zm1 5H2v-2h1zm9 0H4v-2h8zm4 0h-3v-2h3zm0-3h-2V7h2zm0-3h-3V4h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lab(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.942 12.57L10 4.335V1h.5c.275 0 .5-.225.5-.5s-.225-.5-.5-.5h-5c-.275 0-.5.225-.5.5s.225.5.5.5H6v3.335L1.058 12.57C-.074 14.456.8 16 3 16h10c2.2 0 3.074-1.543 1.942-3.43M3.766 10L7 4.61V1h2v3.61L12.234 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lanyrd(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5m-1.65 12.012l-5.444 1.781c-1.244.406-1.369.341-1.931-1.4L4.1 8.134c-.328-1.009-1.328-3.728-1.497-4.25c-.313-.969-.313-1.022 1.516-1.616c1.431-.469 1.491-.453 2.009 1.163c.419 1.3.688 2.35 1.119 3.678l1.172 3.625l3.744-1.225c.738-.244.984-.231 1.194.678l.15.688c.175.797-.228 1-.656 1.137z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11V3c0-.55-.45-1-1-1H3c-.55 0-1 .45-1 1v8H0v3h16v-3zm-4 2H6v-1h4zm3-2H3V3.002L3.002 3h9.996l.002.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Last(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2v12h-2V8.5l-5 5v-5l-5 5v-11l5 5v-5l5 5V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lastfm(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.056 11.972l-.588-1.594s-.953 1.063-2.381 1.063c-1.266 0-2.163-1.1-2.163-2.859c0-2.253 1.137-3.059 2.253-3.059c1.612 0 2.125 1.044 2.566 2.381l.588 1.831c.588 1.778 1.688 3.206 4.856 3.206c2.272 0 3.813-.697 3.813-2.528c0-1.484-.844-2.253-2.419-2.622l-1.172-.256c-.806-.184-1.044-.513-1.044-1.063c0-.622.494-.991 1.3-.991c.881 0 1.356.331 1.428 1.119l1.831-.219c-.147-1.65-1.284-2.328-3.153-2.328c-1.65 0-3.262.622-3.262 2.622c0 1.247.606 2.034 2.125 2.4l1.247.294c.934.219 1.247.606 1.247 1.137c0 .678-.659.953-1.906.953c-1.85 0-2.622-.972-3.059-2.309l-.606-1.831c-.766-2.384-1.994-3.263-4.431-3.263c-2.694 0-4.125 1.703-4.125 4.6c0 2.784 1.428 4.287 3.997 4.287c2.069 0 3.059-.972 3.059-.972z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LastfmTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5m-2.834 11.887c-2.775 0-3.737-1.25-4.25-2.806l-.513-1.603c-.384-1.172-.834-2.084-2.244-2.084c-.978 0-1.972.706-1.972 2.678c0 1.541.784 2.503 1.894 2.503c1.25 0 2.084-.931 2.084-.931l.513 1.394s-.866.85-2.678.85c-2.25 0-3.5-1.313-3.5-3.75c0-2.534 1.25-4.025 3.609-4.025c2.134 0 3.206.769 3.881 2.853l.528 1.603c.384 1.172 1.059 2.022 2.678 2.022c1.091 0 1.669-.241 1.669-.834c0-.466-.272-.803-1.091-.994l-1.091-.256c-1.331-.322-1.859-1.009-1.859-2.1c0-1.747 1.412-2.294 2.853-2.294c1.634 0 2.631.594 2.759 2.038l-1.603.194c-.066-.691-.481-.978-1.25-.978c-.706 0-1.137.322-1.137.866c0 .481.209.769.912.931l1.025.225c1.378.322 2.116.994 2.116 2.294c0 1.597-1.347 2.206-3.334 2.206z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.802 2.102C14.072.791 11.409.008 8.678.008c-3.377 0-6.129 1.179-7.549 3.235C.462 4.208.093 5.352.032 6.641c-.054 1.148.139 2.418.573 3.784C2.087 5.981 6.227 2.502 11 2.502c0 0-4.466 1.175-7.274 4.816A10.535 10.535 0 0 0 2.2 10.037A15.387 15.387 0 0 0 1 16.002h2s-.304-1.91.224-4.106a17.71 17.71 0 0 0 2.357.177c1.839 0 3.146-.398 4.115-1.252c.868-.765 1.347-1.794 1.854-2.882c.774-1.663 1.651-3.547 4.198-5.002a.5.5 0 0 0 .054-.833z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Library(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 15v-1h-1V8h1V7h-3v1h1v6h-3V8h1V7H9v1h1v6H7V8h1V7H5v1h1v6H3V8h1V7H1v1h1v6H1v1H0v1h17v-1zM8 0h1l8 5v1H0V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Libreoffice(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.354.354A1.394 1.394 0 0 0 7.5 0h-6c-.275 0-.5.225-.5.5v15c0 .275.225.5.5.5h12c.275 0 .5-.225.5-.5v-9c0-.275-.159-.659-.354-.854L8.353.353zM13 15H2V1h5.487a.545.545 0 0 1 .169.07l5.274 5.274a.545.545 0 0 1 .07.169zm.5-15h-3c-.275 0-.341.159-.146.354l3.293 3.293c.194.194.354.129.354-.146v-3c0-.275-.225-.5-.5-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lifebuoy(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M5 8a3 3 0 1 1 6 0a3 3 0 0 1-6 0m9.468 2.679l-2.772-1.148C11.892 9.059 12 8.542 12 8s-.108-1.059-.304-1.531l2.772-1.148C14.81 6.146 15 7.051 15 8s-.189 1.854-.532 2.679m-3.789-9.147L9.531 4.304C9.059 4.108 8.542 4 8 4s-1.059.108-1.531.304L5.321 1.532C6.146 1.19 7.051 1 8 1s1.854.189 2.679.532M1.532 5.321l2.772 1.148C4.108 6.941 4 7.458 4 8s.108 1.059.304 1.531l-2.772 1.148C1.19 9.854 1 8.949 1 8s.189-1.854.532-2.679m3.789 9.147l1.148-2.772c.472.196.989.304 1.531.304s1.059-.108 1.531-.304l1.148 2.772C9.854 14.81 8.949 15 8 15s-1.854-.189-2.679-.532"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ligature(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13.622v-.002l-.005-6.821l-1.992.097H6.067V6.56c0-1.274.091-2.546.269-3.042c.123-.343.353-.652.683-.919c.322-.261.643-.393.955-.393c.262 0 .48.045.647.134c.235.134.464.359.682.669c.577.82.812 1.038.939 1.131c.216.158.477.238.776.238c.292 0 .546-.109.757-.324a1.09 1.09 0 0 0 .315-.792c0-.335-.139-.691-.414-1.057c-.268-.358-.683-.652-1.232-.875a4.743 4.743 0 0 0-1.793-.329c-.949 0-1.813.228-2.568.678a4.405 4.405 0 0 0-1.725 1.863c-.359.728-.333 2.105-.355 3.355H2.038v1.116H4v5.073c0 1.12-.342 1.422-.472 1.583a1.225 1.225 0 0 1-.944.455H1.98v.878h6.021v-.878h-.105c-1.424 0-1.828-.154-1.828-1.888v-.001l-.001-5.222h2.191c1.163 0 1.43.054 1.491.077c.074.028.169.075.204.143c.014.026.081.391.081 1.296v3.917c0 .913-.111 1.217-.179 1.319c-.145.222-.319.345-.854.358v.879h4.588v-.873c-1.431 0-1.588-.153-1.588-1.505z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LigatureTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.364 14.335c-.183 0-1.307-.206-1.375-.458c-.161-.619-.183-1.284-.183-2.04V3.384c0-1.261.252-1.994.252-1.994c-.023-.115-.138-.367-.275-.367h-.069c-.069 0-.871.504-1.605.504c-.596 0-.967-.527-1.655-.527c-2.892 0-4.249 2.349-4.249 5.672v.173c0 .069-.046.138-.115.138h-.94c-.115 0-.344.642-.344.94c0 .092.023.137.069.137H4.09c.069 0 .115.092.115.16c0 2.04-.023 4.052-.023 4.052c0 .321-.023 1.031-.16 1.605c-.069.252-1.123.458-1.398.458c-.115 0-.115.55 0 .665c.94-.046 1.559-.115 2.499-.115c.871 0 1.536.069 2.453.115c.046-.138.046-.665-.069-.665c-.183 0-1.307-.206-1.375-.458c-.16-.619-.16-1.284-.183-2.04V8.198c0-.069.069-.138.138-.138h2.361c.16-.321.275-.711.275-.917c0-.138 0-.16-.115-.16H6.064c-.046 0-.115-.069-.115-.115v-.825c0-2.04.836-3.837 2.234-3.837c.99 0 1.854.642 1.854 3.093c.003.063.005.114.005.148v6.825c0 .321-.023 1.031-.16 1.605c-.069.252-1.123.458-1.398.458c-.115 0-.115.55 0 .665c.94-.046 1.559-.115 2.499-.115c.871 0 1.536.069 2.453.115c.046-.137.046-.665-.069-.665z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.879 9.934a.81.81 0 0 1-.575-.238a3.818 3.818 0 0 1 0-5.392l3-3C10.024.584 10.982.187 12 .187s1.976.397 2.696 1.117a3.818 3.818 0 0 1 0 5.392l-1.371 1.371a.813.813 0 0 1-1.149-1.149l1.371-1.371A2.19 2.19 0 0 0 12 1.812c-.584 0-1.134.228-1.547.641l-3 3a2.19 2.19 0 0 0 0 3.094a.813.813 0 0 1-.575 1.387z"/><path fill="currentColor" d="M4 15.813a3.789 3.789 0 0 1-2.696-1.117a3.818 3.818 0 0 1 0-5.392l1.371-1.371a.813.813 0 0 1 1.149 1.149l-1.371 1.371A2.19 2.19 0 0 0 4 14.188c.585 0 1.134-.228 1.547-.641l3-3a2.19 2.19 0 0 0 0-3.094a.813.813 0 0 1 1.149-1.149a3.818 3.818 0 0 1 0 5.392l-3 3A3.789 3.789 0 0 1 4 15.813"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M6 13H4V6h2zM5 5a1 1 0 1 1 0-2a1 1 0 1 1 0 2m8 8h-2V9a1 1 0 1 0-2 0v4H7V6h2v1.241C9.412 6.675 10.044 6 10.75 6C11.994 6 13 7.119 13 8.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 6h2.767v1.418h.04C9.192 6.727 10.134 6 11.539 6C14.46 6 15 7.818 15 10.183V15h-2.885v-4.27c0-1.018-.021-2.329-1.5-2.329c-1.502 0-1.732 1.109-1.732 2.255V15H6zM1 6h3v9H1zm3-2.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 4 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h4v4H0zm6 1h10v2H6zM0 6h4v4H0zm6 1h10v2H6zm-6 5h4v4H0zm6 1h10v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListNumbered(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 13h10v2H6zm0-6h10v2H6zm0-6h10v2H6zM3 0v4H2V1H1V0zM2 8.219V9h2v1H1V7.719l2-.938V6H1V5h3v2.281zM4 11v5H1v-1h2v-1H1v-1h2v-1H1v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 1h10v2H6zm0 6h10v2H6zm0 6h10v2H6zM0 2a2 2 0 1 1 3.999-.001A2 2 0 0 1 0 2m0 6a2 2 0 1 1 3.999-.001A2 2 0 0 1 0 8m0 6a2 2 0 1 1 3.999-.001A2 2 0 0 1 0 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a5 5 0 0 0-5 5c0 5 5 11 5 11s5-6 5-11a5 5 0 0 0-5-5m0 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a5 5 0 0 0-5 5c0 5 5 11 5 11s5-6 5-11a5 5 0 0 0-5-5m0 8.063a3.063 3.063 0 1 1 0-6.126a3.063 3.063 0 0 1 0 6.126M6.063 5a1.938 1.938 0 1 1 3.876 0a1.938 1.938 0 0 1-3.876 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.25 7H9V4c0-1.654-1.346-3-3-3H4C2.346 1 1 2.346 1 4v3H.75a.753.753 0 0 0-.75.75v7.5c0 .412.338.75.75.75h8.5c.412 0 .75-.338.75-.75v-7.5A.753.753 0 0 0 9.25 7M3 4c0-.551.449-1 1-1h2c.551 0 1 .449 1 1v3H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loop(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h10v3l4-4l-4-4v3H0v6h2zm12 6H4V8l-4 4l4 4v-3h12V7h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoopTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.901 2.599A8 8 0 0 0 0 8h1.5a6.5 6.5 0 0 1 11.339-4.339L10.5 6H16V.5zM14.5 8a6.5 6.5 0 0 1-11.339 4.339L5.5 10H0v5.5l2.099-2.099A8 8 0 0 0 16 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ltr(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a4 4 0 0 0 0 8v8h2V2h2v14h2V2h2V0zM0 11l4-4l-4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagicWand(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3L2 1H1v1l2 2zm1-3h1v2H5zm4 5h2v1H9zm1-3V1H9L7 3l1 1zM0 5h2v1H0zm5 4h1v2H5zM1 9v1h1l2-2l-1-1zm14.781 4.781L5.842 3.842a.752.752 0 0 0-1.061 0l-.939.939a.752.752 0 0 0 0 1.061l9.939 9.939a.752.752 0 0 0 1.061 0l.939-.939a.752.752 0 0 0 0-1.061M7.5 8.5l-3-3l1-1l3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 0h-4l1 9a3 3 0 1 1-6 0l1-9H2L1 9a7 7 0 1 0 14 0zm-1.846 13.154c-1.11 1.11-2.585 1.721-4.154 1.721s-3.045-.611-4.154-1.721a5.834 5.834 0 0 1-1.72-4.095l.564-5.075h1.736l-.55 4.953v.062c0 1.102.429 2.138 1.208 2.917s1.815 1.208 2.917 1.208s2.138-.429 2.917-1.208a4.098 4.098 0 0 0 1.208-2.917v-.062l-.007-.062l-.543-4.891h1.736l.564 5.075a5.838 5.838 0 0 1-1.72 4.095z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.333 0H2.667A2.675 2.675 0 0 0 0 2.667v10.666C0 14.8 1.2 16 2.667 16h10.666C14.801 16 16 14.8 16 13.333V2.667A2.674 2.674 0 0 0 13.333 0M4 4h8c.143 0 .281.031.409.088L8 9.231L3.591 4.088A.982.982 0 0 1 4 4m-1 7V5l.002-.063l2.932 3.421l-2.9 2.9A.967.967 0 0 1 3 11m9 1H4c-.088 0-.175-.012-.258-.034L6.588 9.12l1.413 1.648L9.414 9.12l2.846 2.846a.967.967 0 0 1-.258.034zm1-1c0 .088-.012.175-.034.258l-2.9-2.9l2.932-3.421L13 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailFour(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M4 4h8c.143 0 .281.031.409.088L8 9.231L3.591 4.088A.982.982 0 0 1 4 4m-1 7V5l.002-.063l2.932 3.421l-2.9 2.9A.967.967 0 0 1 3 11m9 1H4c-.088 0-.175-.012-.258-.034L6.588 9.12l1.413 1.648L9.414 9.12l2.846 2.846a.967.967 0 0 1-.258.034zm1-1c0 .088-.012.175-.034.258l-2.9-2.9l2.932-3.421L13 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.333 0H2.667A2.675 2.675 0 0 0 0 2.667v10.666C0 14.801 1.2 16 2.667 16h10.666A2.674 2.674 0 0 0 16 13.333V2.667C16 1.2 14.8 0 13.333 0M2.854 13.854l-1.207-1.207l4-4l.457.457zm-.458-10.75l.457-.457l5.146 4.146l5.146-4.146l.457.457l-5.604 6.604l-5.604-6.604zm10.75 10.75l-3.25-4.75l.457-.457l4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.333 0H2.667A2.675 2.675 0 0 0 0 2.667v10.666C0 14.801 1.2 16 2.667 16h10.666A2.674 2.674 0 0 0 16 13.333V2.667C16 1.2 14.8 0 13.333 0m0 2c.125 0 .243.036.344.099L7.999 6.793L2.322 2.099A.648.648 0 0 1 2.666 2zM2.667 14a.654.654 0 0 1-.089-.006l3.525-4.89l-.457-.457L2 12.293V2.744L8 10l6-7.256v9.549l-3.646-3.646l-.457.457l3.525 4.89a.65.65 0 0 1-.088.006z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MakeGroup(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2H3c-.55 0-1 .45-1 1v2c0 .55.45 1 1 1h2c.55 0 1-.45 1-1V3c0-.55-.45-1-1-1m6 4h2c.55 0 1-.45 1-1V3c0-.55-.45-1-1-1h-2c-.55 0-1 .45-1 1v2c0 .55.45 1 1 1m0-3h2v2h-2zm-6 7H3c-.55 0-1 .45-1 1v2c0 .55.45 1 1 1h2c.55 0 1-.45 1-1v-2c0-.55-.45-1-1-1m0 3H3v-2h2zm8-3h-2c-.55 0-1 .45-1 1v2c0 .55.45 1 1 1h2c.55 0 1-.45 1-1v-2c0-.55-.45-1-1-1"/><path fill="currentColor" d="M14 8h-1c-1.336 0-2.591-.52-3.536-1.464S8 4.336 8 3V2c0-1.1-.9-2-2-2H2C.9 0 0 .9 0 2v4c0 1.1.9 2 2 2h1c1.336 0 2.591.52 3.536 1.464S8 11.664 8 13v1c0 1.1.9 2 2 2h4c1.1 0 2-.9 2-2v-4c0-1.1-.9-2-2-2m1 6c0 .265-.105.515-.295.705S14.265 15 14 15h-4c-.265 0-.515-.105-.705-.295S9 14.265 9 14v-1a6 6 0 0 0-6-6H2c-.265 0-.515-.105-.705-.295S1 6.264 1 6V2c0-.265.105-.515.295-.705S1.735 1 2 1h4c.265 0 .515.105.705.295S7 1.735 7 2v1a6 6 0 0 0 6 6h1c.265 0 .515.105.705.295S15 9.735 15 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Man(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 9 1.5M9 4H6a1 1 0 0 0-1 1v5h1v6h1.25v-6h.5v6H9v-6h1V5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ManWoman(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 1.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 4 1.5m9 0a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 13 1.5M4 4H1a1 1 0 0 0-1 1v5h1v6h1.25v-6h.5v6H4v-6h1V5a1 1 0 0 0-1-1m11.234 4L16 7.445l-2.083-3.221a.5.5 0 0 0-.417-.225h-4a.497.497 0 0 0-.417.225L7 7.445L7.766 8l1.729-2.244l.601 1.402l-2.095 3.841h1.917l.333 5h1v-5h.5v5h1l.333-5h1.917l-2.095-3.842l.601-1.402l1.729 2.244z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 3l5-2v12l-5 2zM6 .5l5 3V15l-5-2.5zm6 3l4-3v12l-4 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.5 3l-5-2L0 3v12l5.5-2l5 2l5.5-2V1zM6 2.277l4 1.6v9.846l-4-1.6zM1 3.7l4-1.455v9.872l-4 1.454V3.699zm14 8.6l-4 1.455V3.883l4-1.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3h14v3H1zm0 4h14v3H1zm0 4h14v3H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuFour(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h14v3H0zm0 4h14v3H0zm0 4h14v3H0zm15.5-1l3-3l3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h14v3H0zm0 4h14v3H0zm0 4h14v3H0zm15.5-4l3 3l3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h14v3H0zm0 4h14v3H0zm0 4h14v3H0zm15.5-2l3 3l3-3zm6-1l-3-3l-3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meter(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1a8 8 0 0 1 3.875 15h-7.75A8 8 0 0 1 8 1m4.53 12.53A6.364 6.364 0 0 0 14.406 9H13V8h1.329a6.346 6.346 0 0 0-.665-2H12V5h1.004a6.372 6.372 0 0 0-3.005-2.089V4h-1V2.671a6.506 6.506 0 0 0-2 0V4h-1V2.911A6.384 6.384 0 0 0 2.994 5h1.004v1H2.334a6.346 6.346 0 0 0-.665 2h1.329v1H1.592c0 1.711.666 3.32 1.876 4.53c.167.167.343.324.524.47h3.006l.571-8h.857l.571 8h3.006c.182-.146.357-.303.524-.47z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MeterTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M4.732 13.034a4.501 4.501 0 0 0-2.688-5.762a5.953 5.953 0 0 1 1.714-3.514c1.133-1.133 2.64-1.757 4.243-1.757s3.109.624 4.243 1.757a5.958 5.958 0 0 1 1.714 3.514a4.501 4.501 0 0 0-2.688 5.762c-.964.629-2.09.966-3.268.966s-2.304-.338-3.268-.966zm3.889-3.018A.502.502 0 0 1 9 10.5v1c0 .275-.225.5-.5.5h-1a.501.501 0 0 1-.5-.5v-1a.5.5 0 0 1 .379-.484L7.75 3h.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 11A2.5 2.5 0 0 0 10 8.5v-6a2.5 2.5 0 1 0-5 0v6A2.5 2.5 0 0 0 7.5 11M11 7v1.5a3.5 3.5 0 1 1-7 0V7H3v1.5a4.5 4.5 0 0 0 4 4.472V15H5v1h5v-1H8v-2.028A4.5 4.5 0 0 0 12 8.5V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 6.5v3a.5.5 0 0 0 .5.5h15a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5H.5a.5.5 0 0 0-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 0h-7C3.675 0 3 .675 3 1.5v13c0 .825.675 1.5 1.5 1.5h7c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M6 .75h4v.5H6zM8 15a1 1 0 1 1 0-2a1 1 0 0 1 0 2m4-3H4V2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0H3c-.55 0-1 .45-1 1v14c0 .55.45 1 1 1h9c.55 0 1-.45 1-1V1c0-.55-.45-1-1-1M7.5 15.278a.778.778 0 1 1 0-1.555a.778.778 0 0 1 0 1.555M12 13H3V2h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveDown(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11V5h-1v6H8.5l3 3l3-3zM5 4v3H2V4zm1-1H1v5h5zm-5 7h1.5v1H1zm2 0h1.5v1H3zm2 0h1v1.5H5zm-4 3.5h1V15H1zm1.5.5H4v1H2.5zm2 0H6v1H4.5zM1 11.5h1V13H1zm4 .5h1v1.5H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveUp(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 8v6h1V8h2.5l-3-3l-3 3zM1 3h1.5v1H1zm2 0h1.5v1H3zm2 0h1v1.5H5zM1 6.5h1V8H1zm1.5.5H4v1H2.5zm2 0H6v1H4.5zM1 4.5h1V6H1zM5 5h1v1.5H5zm0 6v3H2v-3zm1-1H1v5h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mug(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 5h-3V3.5C12 2.119 9.314 1 6 1S0 2.119 0 3.5v10C0 14.881 2.686 16 6 16s6-1.119 6-2.5V12h3a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1M2.751 4.037c-.578-.19-.928-.394-1.116-.537c.188-.143.538-.347 1.116-.537C3.656 2.665 4.81 2.5 6 2.5s2.344.164 3.249.463c.578.19.928.394 1.116.537c-.188.143-.538.347-1.116.537C8.344 4.335 7.19 4.5 6 4.5s-2.344-.164-3.249-.463M14 10h-2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 0h1v11.5c0 1.381-1.567 2.5-3.5 2.5S9 12.881 9 11.5S10.567 9 12.5 9c.979 0 1.865.287 2.5.751V4L7 5.778V13.5C7 14.881 5.433 16 3.5 16S0 14.881 0 13.5S1.567 11 3.5 11c.979 0 1.865.287 2.5.751V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Neutral(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13M4 5a1 1 0 1 0 2 0a1 1 0 0 0-2 0m6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-4 6h4v1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NeutralTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m2 12H6v-1h4zm1-8a1 1 0 1 1 0 2a1 1 0 0 1 0-2M5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewTab(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1v12h12V1zm11 11H4V2h10zM2 14V3.5l-1-1V15h12.5l-1-1z"/><path fill="currentColor" d="M5.5 4L8 6.5l-3 3L6.5 11l3-3l2.5 2.5V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newspaper(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4V2H0v11a1 1 0 0 0 1 1h13.5a1.5 1.5 0 0 0 1.5-1.5V4zm-1 9H1V3h12zM2 5h10v1H2zm6 2h4v1H8zm0 2h4v1H8zm0 2h3v1H8zM2 7h5v5H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 1 0 16A8 8 0 0 1 8 0m0 14.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13"/><path fill="currentColor" d="M9 8L5 5v6zm2-3H9v6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2v12h-2V8.5l-5 5v-11l5 5V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notification(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1.5c-1.736 0-3.369.676-4.596 1.904S1.5 6.264 1.5 8c0 1.736.676 3.369 1.904 4.596S6.264 14.5 8 14.5c1.736 0 3.369-.676 4.596-1.904S14.5 9.736 14.5 8c0-1.736-.676-3.369-1.904-4.596S9.736 1.5 8 1.5M8 0a8 8 0 1 1 0 16A8 8 0 0 1 8 0M7 11h2v2H7zm0-8h2v6H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Npm(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm13 13h-2V5H8v8H3V3h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Office(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 16h8V0H0zM5 2h2v2H5zm0 4h2v2H5zm0 4h2v2H5zM1 2h2v2H1zm0 4h2v2H1zm0 4h2v2H1zm8-5h7v1H9zm0 11h2v-4h3v4h2V7H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Omega(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 14h4l1-2v4h-6v-3.347c2.049-.883 3.5-3.081 3.5-5.653c0-3.35-2.462-5.973-5.5-5.973S2.5 3.649 2.5 7c0 2.572 1.451 4.77 3.5 5.653V16H0v-4l1 2h4v-.509C2.068 12.453 0 9.938 0 7c0-3.866 3.582-7 8-7s8 3.134 8 7c0 2.938-2.068 5.452-5 6.491z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Onedrive(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.482 12.944c-.942-.235-1.466-.984-1.468-2.095c0-.355.025-.525.114-.754c.217-.56.793-.982 1.55-1.138c.377-.077.493-.16.493-.353c0-.06.045-.24.1-.399c.249-.724.71-1.327 1.202-1.573c.515-.258.776-.316 1.399-.313c.886.005 1.327.197 1.945.846l.34.357l.304-.105c1.473-.51 2.942.358 3.061 1.809l.032.397l.29.104c.829.297 1.218.92 1.148 1.837c-.046.599-.326 1.078-.77 1.315l-.209.112l-4.638.009c-3.564.007-4.697-.006-4.893-.055zm-3.869-.663c-.565-.142-1.164-.67-1.445-1.273C.009 10.666 0 10.615 0 10.01c0-.576.014-.668.14-.954a2.141 2.141 0 0 1 1.422-1.21c.136-.036.263-.094.283-.128s.043-.221.05-.415c.045-1.206.794-2.269 1.839-2.61c.565-.184 1.306-.202 1.92.058c.195.082.173.1.585-.471c.244-.338.705-.695 1.108-.909c.435-.231.887-.337 1.428-.336c1.512.004 2.815 1.003 3.297 2.529c.154.487.146.624-.035.628a3.144 3.144 0 0 0-.505.102l-.361.099l-.329-.348c-.928-.98-2.441-1.192-3.728-.522a3.156 3.156 0 0 0-1.239 1.153c-.222.357-.506 1.024-.506 1.189c0 .117-.09.176-.474.309c-1.189.412-1.883 1.364-1.882 2.582c0 .443.108.986.258 1.296c.057.117.088.228.07.247c-.046.049-1.525.032-1.73-.019z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opera(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8c0 2.369-1.031 4.5-2.669 5.963c-2.053 1-3.966.3-4.597-.137c2.016-.441 3.537-2.878 3.537-5.825s-1.522-5.384-3.537-5.828c.634-.438 2.547-1.137 4.597-.138A7.99 7.99 0 0 1 16 8.001z"/><path fill="currentColor" d="M5.366 3.491C4.482 4.535 3.91 6.078 3.872 7.813v.378c.038 1.731.613 3.275 1.497 4.319c1.147 1.491 2.853 2.434 4.759 2.434a5.768 5.768 0 0 0 3.206-.978a7.984 7.984 0 0 1-5.715 2.025A8 8 0 0 1 8 0h.031a7.952 7.952 0 0 1 5.303 2.038a5.773 5.773 0 0 0-3.206-.981c-1.906 0-3.612.944-4.763 2.434z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opt(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 13h-4a.499.499 0 0 1-.457-.297L6.175 4H1.5a.5.5 0 0 1 0-1h5c.198 0 .377.116.457.297L10.825 12H14.5a.5.5 0 0 1 0 1m0-9h-5a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pacman(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.074 2.794a8 8 0 1 0 0 10.412L10 8zM11 1.884a1.116 1.116 0 1 1 0 2.232a1.116 1.116 0 0 1 0-2.232"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageBreak(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 8h2v1H0zm3 0h3v1H3zm4 0h2v1H7zm3 0h3v1h-3zm4 0h2v1h-2zm-.25-8L14 7H2l.25-7h.5L3 6h10l.25-6zM2.25 16L2 10h12l-.25 6h-.5L13 11H3l-.25 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pagebreak(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6V0h12v6h-1V1H5v5zm12 3v7H4V9h1v6h10V9zM8 7h2v1H8zM5 7h2v1H5zm6 0h2v1h-2zm3 0h2v1h-2zM0 4.5l3 3l-3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintFormat(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 9V3h-3V2c0-.55-.45-1-1-1H1c-.55 0-1 .45-1 1v3c0 .55.45 1 1 1h11c.55 0 1-.45 1-1V4h2v4H6v2h-.5a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5H7V9zm-4-6H1V2h11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphCenter(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h16v2H0zm3 3h10v2H3zm0 6h10v2H3zM0 7h16v2H0zm0 6h16v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphJustify(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h16v2H0zm0 3h16v2H0zm0 3h16v2H0zm0 3h16v2H0zm0 3h16v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphLeft(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h16v2H0zm0 3h10v2H0zm0 6h10v2H0zm0-3h16v2H0zm0 6h16v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParagraphRight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h16v2H0zm6 3h10v2H6zm0 6h10v2H6zM0 7h16v2H0zm0 6h16v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paste(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2H9V1c0-.55-.45-1-1-1H6c-.55 0-1 .45-1 1v1H3v2h8zM8 2H6v-.998L6.002 1h1.996L8 1.002zm5 3V2.5c0-.275-.225-.5-.5-.5h-1v1h.5v2H9L6 8v4H2V3h.5V2h-1c-.275 0-.5.225-.5.5v10c0 .275.225.5.5.5H6v3h10V5zM9 6.414V8H7.414zM15 15H7V9h3V6h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13M5 5h2v6H5zm4 0h2v6H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h5v12H2zm7 0h5v12H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.531 4.822c-.747 3.316-3.053 5.066-6.688 5.066H6.634l-.841 5.338H4.78l-.053.344a.374.374 0 0 0 .369.431h2.588a.625.625 0 0 0 .616-.525l.025-.131l.488-3.091l.031-.169a.622.622 0 0 1 .616-.525h.384c2.506 0 4.469-1.019 5.044-3.963c.216-1.119.134-2.069-.356-2.775z"/><path fill="currentColor" d="M12.984 1.206C12.243.362 10.903 0 9.19 0H4.218c-.35 0-.65.253-.703.6L1.443 13.731a.427.427 0 0 0 .422.494h3.072l.772-4.891l-.025.153c.053-.347.35-.6.7-.6h1.459c2.866 0 5.109-1.162 5.766-4.531c.019-.1.037-.197.05-.291c.194-1.244 0-2.094-.675-2.859"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.909 4.561l-4.47-4.47c-.146-.146-.338-.113-.427.073l-.599 1.248l4.175 4.175l1.248-.599c.186-.089.219-.282.073-.427M9.615 2.115L5.5 2.458c-.273.034-.501.092-.58.449v.001C3.804 8.268 0 13.499 0 13.499l.896.896l4.25-4.25a1.5 1.5 0 1 1 .707.707l-4.25 4.25l.896.896s5.231-3.804 10.591-4.92h.001c.357-.078.415-.306.449-.58l.343-4.115l-4.269-4.269z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 0a2.5 2.5 0 0 1 2 4l-1 1L11 1.5l1-1c.418-.314.937-.5 1.5-.5M1 11.5L0 16l4.5-1l9.25-9.25l-3.5-3.5zm10.181-5.819l-7 7l-.862-.862l7-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6 10l2-1l7-7l-1-1l-7 7zm-1.48 3.548c-.494-1.043-1.026-1.574-2.069-2.069l1.548-4.262l2-1.217l6-6h-3l-6 6l-3 10l10-3l6-6V4l-6 6l-1.217 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 10c-1 1-1 2-2 2s-2-1-3-2s-2-2-2-3s1-1 2-2s-2-4-3-4s-3 3-3 3c0 2 2.055 6.055 4 8s6 4 8 4c0 0 3-2 3-3s-3-4-4-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneHangUp(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.897 9c.125.867.207 2.053-.182 2.507c-.643.751-4.714.751-4.714-.751c0-.756.67-1.252.027-2.003c-.632-.738-1.766-.75-3.027-.751s-2.394.012-3.027.751c-.643.751.027 1.247.027 2.003c0 1.501-4.071 1.501-4.714.751C-.102 11.053-.02 9.867.105 9c.096-.579.339-1.203 1.118-2c1.168-1.09 2.935-1.98 6.716-2h.126c3.781.019 5.548.91 6.716 2c.778.797 1.022 1.421 1.118 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChart(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 9V2a7 7 0 1 0 6.262 3.869zm7.262-5.131A6.999 6.999 0 0 0 8 0v7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pilcrow(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0h8v2h-2v14h-2V2H8v14H6V8a4 4 0 0 1 0-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1.069a6.93 6.93 0 0 0-2.525 13.384c-.059-.547-.116-1.391.025-1.988c.125-.541.813-3.444.813-3.444s-.206-.416-.206-1.028c0-.963.559-1.684 1.253-1.684c.591 0 .878.444.878.975c0 .594-.378 1.484-.575 2.306c-.166.691.344 1.253 1.025 1.253c1.231 0 2.178-1.3 2.178-3.175c0-1.659-1.194-2.819-2.894-2.819c-1.972 0-3.128 1.478-3.128 3.009c0 .597.228 1.234.516 1.581c.056.069.066.128.047.2a95.89 95.89 0 0 1-.194.787c-.031.128-.1.153-.231.094c-.866-.403-1.406-1.669-1.406-2.684c0-2.188 1.587-4.194 4.578-4.194c2.403 0 4.272 1.712 4.272 4.003c0 2.388-1.506 4.313-3.597 4.313c-.703 0-1.362-.366-1.588-.797c0 0-.347 1.322-.431 1.647c-.156.603-.578 1.356-.862 1.816a6.93 6.93 0 0 0 8.984-6.622a6.931 6.931 0 0 0-6.931-6.934z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.588 0 0 3.587 0 8s3.587 8 8 8s8-3.588 8-8s-3.588-8-8-8m0 14.931a6.959 6.959 0 0 1-2.053-.309c.281-.459.706-1.216.862-1.816c.084-.325.431-1.647.431-1.647c.225.431.888.797 1.587.797c2.091 0 3.597-1.922 3.597-4.313c0-2.291-1.869-4.003-4.272-4.003c-2.991 0-4.578 2.009-4.578 4.194c0 1.016.541 2.281 1.406 2.684c.131.063.2.034.231-.094c.022-.097.141-.566.194-.787a.213.213 0 0 0-.047-.2c-.287-.347-.516-.988-.516-1.581c0-1.528 1.156-3.009 3.128-3.009c1.703 0 2.894 1.159 2.894 2.819c0 1.875-.947 3.175-2.178 3.175c-.681 0-1.191-.563-1.025-1.253c.197-.825.575-1.713.575-2.306c0-.531-.284-.975-.878-.975c-.697 0-1.253.719-1.253 1.684c0 .612.206 1.028.206 1.028s-.688 2.903-.813 3.444c-.141.6-.084 1.441-.025 1.988a6.922 6.922 0 0 1-4.406-6.45a6.93 6.93 0 1 1 6.931 6.931z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.331 2.502C13.087 2.179 10.607 2 8 2s-5.087.179-7.331.502C.239 4.185 0 6.045 0 8s.239 3.815.669 5.498C2.913 13.821 5.393 14 8 14s5.087-.179 7.331-.502C15.761 11.815 16 9.955 16 8s-.239-3.815-.669-5.498M6 11V5l5 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3 2l10 6l-10 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13m-2-10L12 8l-6 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 6H10V.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5V6H.5a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5H6v5.5a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5V10h5.5a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Podcast(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8a8 8 0 1 0-10.786 7.501L5 16h6l-.214-.499A8.002 8.002 0 0 0 16 8M7.606 9.919a1 1 0 1 1 .788 0L8 9zm.804.039a2 2 0 1 0-.82-.001l-1.166 2.721A4.002 4.002 0 0 1 4 9.001c0-2.209 1.791-4.188 4-4.188s4 1.978 4 4.188a4 4 0 0 1-2.424 3.677zm2.347 5.475l-1.155-2.695A5.002 5.002 0 0 0 8 3a5 5 0 0 0-1.602 9.738l-1.155 2.695A6.997 6.997 0 0 1 1.003 9c0-3.865 3.133-7.185 6.997-7.185S14.997 5.135 14.997 9a7 7 0 0 1-4.24 6.433"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointDown(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6.5V9a1.502 1.502 0 0 1-2.236 1.307a1.5 1.5 0 0 1-2.264.31a1.494 1.494 0 0 1-1.5.297V14.5c0 .827-.673 1.5-1.5 1.5S6 15.327 6 14.5V8.333L3.25 9.799a1.502 1.502 0 0 1-1.789-2.381l.012-.011L5.21 4H4.5a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-.691l1.138 2.276A.496.496 0 0 1 15 6.5m-1-4a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0m0 4.118L12.691 4H6.694L2.15 8.143a.5.5 0 0 0 .614.782l3.5-1.866a.502.502 0 0 1 .735.442v7a.5.5 0 0 0 1 0v-5a.5.5 0 0 1 1 0a.5.5 0 0 0 1 0a.5.5 0 0 1 1 0a.5.5 0 0 0 1 0v-.5a.5.5 0 0 1 1 0a.5.5 0 0 0 1 0V6.619z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointLeft(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 15H7a1.502 1.502 0 0 1-1.307-2.236a1.5 1.5 0 0 1-.31-2.264A1.494 1.494 0 0 1 5.086 9H1.5C.673 9 0 8.327 0 7.5S.673 6 1.5 6h6.167L6.201 3.25a1.502 1.502 0 0 1 2.381-1.789l.011.012L12 5.21V4.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-.691l-2.276 1.138A.496.496 0 0 1 9.5 15m4-1a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m-4.118 0L12 12.691V6.694L7.857 2.15a.5.5 0 0 0-.782.614l1.866 3.5a.499.499 0 0 1-.441.735h-7a.5.5 0 0 0 0 1h5a.5.5 0 0 1 0 1a.5.5 0 0 0 0 1a.5.5 0 0 1 0 1a.5.5 0 0 0 0 1H7a.5.5 0 0 1 0 1a.5.5 0 0 0 0 1h2.382z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointRight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 15H9a1.502 1.502 0 0 0 1.307-2.236a1.5 1.5 0 0 0 .31-2.264a1.494 1.494 0 0 0 .297-1.5H14.5c.827 0 1.5-.673 1.5-1.5S15.327 6 14.5 6H8.333l1.466-2.75a1.502 1.502 0 0 0-2.381-1.789l-.011.012L4 5.21V4.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-.691l2.276 1.138A.496.496 0 0 0 6.5 15m-4-1a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m4.118 0L4 12.691V6.694L8.143 2.15a.5.5 0 0 1 .782.614l-1.866 3.5a.502.502 0 0 0 .442.735h7a.5.5 0 0 1 0 1h-5a.5.5 0 0 0 0 1a.5.5 0 0 1 0 1a.5.5 0 0 0 0 1a.5.5 0 0 1 0 1h-.5a.5.5 0 0 0 0 1a.5.5 0 0 1 0 1H6.619z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointUp(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 9.5V7a1.502 1.502 0 0 0-2.236-1.307a1.5 1.5 0 0 0-2.264-.31A1.494 1.494 0 0 0 9 5.086V1.5C9 .673 8.327 0 7.5 0S6 .673 6 1.5v6.167L3.25 6.201a1.502 1.502 0 0 0-1.789 2.381l.012.011L5.21 12H4.5a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5h-.691l1.138-2.276A.496.496 0 0 0 15 9.5m-1 4a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0m0-4.118L12.691 12H6.694L2.15 7.857a.5.5 0 0 1 .614-.782l3.5 1.866a.499.499 0 0 0 .735-.441v-7a.5.5 0 0 1 1 0v5a.5.5 0 0 0 1 0a.5.5 0 0 1 1 0a.5.5 0 0 0 1 0a.5.5 0 0 1 1 0V7a.5.5 0 0 0 1 0a.5.5 0 0 1 1 0v2.382z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0L0 8h6l-4 8L16 6H8l6-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerCord(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4.414L14.586 3l-2.793 2.793l-1.586-1.586L13 1.414L11.586 0L8.793 2.793L7 1L5.646 2.353l8 8L15 9l-1.793-1.793zm-3.593 6.114L5.472 3.593C3.975 5.388 2.276 8.163 3.45 10.55l-2.066 2.066a1.254 1.254 0 0 0 0 1.768l.232.232a1.254 1.254 0 0 0 1.768 0L5.45 12.55c2.387 1.174 5.161-.524 6.957-2.022"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13"/><path fill="currentColor" d="m7 8l4-3v6zM5 5h2v6H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PreviousTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 14V2h2v5.5l5-5v11l-5-5V14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriceTag(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 0h-6c-.412 0-.989.239-1.28.53L.531 7.969a.752.752 0 0 0 0 1.061l6.439 6.439a.752.752 0 0 0 1.061 0L15.47 8.03c.292-.292.53-.868.53-1.28v-6a.753.753 0 0 0-.75-.75M11.5 6a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 11.5 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriceTags(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.25 0h-6c-.412 0-.989.239-1.28.53L4.531 7.969a.752.752 0 0 0 0 1.061l6.439 6.439a.752.752 0 0 0 1.061 0L19.47 8.03c.292-.292.53-.868.53-1.28v-6a.752.752 0 0 0-.75-.75M15.5 6a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 15.5 6"/><path fill="currentColor" d="M2 8.5L10.5 0H9.25c-.412 0-.989.239-1.28.53L.531 7.969a.752.752 0 0 0 0 1.061l6.439 6.439a.752.752 0 0 0 1.061 0l.47-.47z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 1h8v2H4zm11 3H1c-.55 0-1 .45-1 1v5c0 .55.45 1 1 1h3v4h8v-4h3c.55 0 1-.45 1-1V5c0-.55-.45-1-1-1M2 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2m9 7H5V9h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Profile(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 0h-12C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h12c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M13 14H2V2h11zM4 9h7v1H4zm0 2h7v1H4zm1-6.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 5 4.5M7.5 6h-2C4.675 6 4 6.45 4 7v1h5V7c0-.55-.675-1-1.5-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pushpin(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 0L7 1.5L8.5 3L5 7H1.5l2.75 2.75L0 15.385V16h.615l5.635-4.25L9 14.5V11l4-3.5L14.5 9L16 7.5zM7 8.5l-1-1L9.5 4l1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qrcode(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1H1v4h4zm1-1v6H0V0zM2 2h2v2H2zm13-1h-4v4h4zm1-1v6h-6V0zm-4 2h2v2h-2zm-7 9H1v4h4zm1-1v6H0v-6zm-4 2h2v2H2zM7 0h1v1H7zm1 1h1v1H8zM7 2h1v1H7zm1 1h1v1H8zM7 4h1v1H7zm1 1h1v1H8zM7 6h1v1H7zm0 2h1v1H7zm1 1h1v1H8zm-1 1h1v1H7zm1 1h1v1H8zm-1 1h1v1H7zm1 1h1v1H8zm-1 1h1v1H7zm1 1h1v1H8zm7-7h1v1h-1zM1 8h1v1H1zm1-1h1v1H2zM0 7h1v1H0zm4 0h1v1H4zm1 1h1v1H5zm1-1h1v1H6zm3 1h1v1H9zm1-1h1v1h-1zm1 1h1v1h-1zm1-1h1v1h-1zm1 1h1v1h-1zm1-1h1v1h-1zm1 3h1v1h-1zm-6 0h1v1H9zm1-1h1v1h-1zm1 1h1v1h-1zm2 0h1v1h-1zm1-1h1v1h-1zm1 3h1v1h-1zm-6 0h1v1H9zm1-1h1v1h-1zm2 0h1v1h-1zm1 1h1v1h-1zm1-1h1v1h-1zm1 3h1v1h-1zm-5-1h1v1h-1zm1 1h1v1h-1zm1-1h1v1h-1zm1 1h1v1h-1zm-3 1h1v1h-1zm2 0h1v1h-1zm2 0h1v1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 11h2v2H7zm4-7a1 1 0 0 1 1 1v3l-3 2H7V9l3-2V6H5V4zM8 1.5c-1.736 0-3.369.676-4.596 1.904S1.5 6.264 1.5 8c0 1.736.676 3.369 1.904 4.596S6.264 14.5 8 14.5c1.736 0 3.369-.676 4.596-1.904S14.5 9.736 14.5 8c0-1.736-.676-3.369-1.904-4.596S9.736 1.5 8 1.5M8 0a8 8 0 1 1 0 16A8 8 0 0 1 8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quill(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 16C2 10 7.234 0 16 0c-4.109 3.297-6 11-9 11H4l-3 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuotesLeft(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.516 7a3.5 3.5 0 1 1-3.5 3.5L0 10a7 7 0 0 1 7-7v2a4.97 4.97 0 0 0-3.536 1.464a5.01 5.01 0 0 0-.497.578c.179-.028.362-.043.548-.043zm9 0a3.5 3.5 0 1 1-3.5 3.5L9 10a7 7 0 0 1 7-7v2a4.97 4.97 0 0 0-3.536 1.464a5.01 5.01 0 0 0-.497.578c.179-.028.362-.043.549-.043z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuotesRight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 10A3.5 3.5 0 1 1 16 6.5l.016.5a7 7 0 0 1-7 7v-2a4.97 4.97 0 0 0 3.536-1.464a5.01 5.01 0 0 0 .497-.578a3.547 3.547 0 0 1-.549.043zm-9 0A3.5 3.5 0 1 1 7 6.5l.016.5a7 7 0 0 1-7 7v-2a4.97 4.97 0 0 0 3.536-1.464a5.01 5.01 0 0 0 .497-.578a3.547 3.547 0 0 1-.549.043z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioChecked(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14A6 6 0 1 1 8 2a6 6 0 0 1 0 12M5 8a3 3 0 1 1 6 0a3 3 0 0 1-6 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioCheckedTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 10a2 2 0 1 1-.001-3.999A2 2 0 0 1 8 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioUnchecked(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14A6 6 0 1 1 8 2a6 6 0 0 1 0 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 10a1 1 0 1 1 2 0a1 1 0 0 1-2 0m6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m.049 2.137a.593.593 0 1 1 .735.933c-.717.565-1.81.93-2.783.93s-2.066-.365-2.784-.93a.593.593 0 1 1 .735-.933c.413.325 1.23.675 2.049.675s1.636-.35 2.049-.675zM16 8a2 2 0 0 0-3.748-.972c-1.028-.562-2.28-.926-3.645-1.01L9.8 3.338l2.284.659a1.5 1.5 0 1 0 .094-1.209l-2.545-.735a.593.593 0 0 0-.707.329L7.305 6.023c-1.33.094-2.551.453-3.557 1.004a2 2 0 1 0-2.555 2.802A3.661 3.661 0 0 0 1 10.999c0 2.761 3.134 5 7 5s7-2.239 7-5c0-.403-.067-.795-.193-1.17A2 2 0 0 0 16 7.999zm-2.5-5.062a.563.563 0 1 1 0 1.126a.563.563 0 0 1 0-1.126M1 8a1 1 0 0 1 1.904-.427a5.292 5.292 0 0 0-1.276 1.355A1.001 1.001 0 0 1 1 8m7 6.813c-3.21 0-5.813-1.707-5.813-3.813S4.789 7.187 8 7.187c3.21 0 5.813 1.707 5.813 3.813S11.211 14.813 8 14.813m6.372-5.885a5.276 5.276 0 0 0-1.276-1.355C13.257 7.235 13.601 7 14 7a1.001 1.001 0 0 1 .372 1.928"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 9a7.982 7.982 0 0 0 2.709 6l1.323-1.5a6 6 0 1 1 8.212-8.743L10.001 7h6V1l-2.343 2.343A8 8 0 0 0 .001 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedoTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3.881V0l6 6l-6 6V8.034C2.02 7.87 2.319 12.781 4.096 16C-.29 11.259.641 3.663 9 3.881"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Renren(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.644.166C2.875.8 0 4.079 0 8.028a7.96 7.96 0 0 0 1.887 5.15c2.791-1.35 4.744-4.406 4.756-7.966zm2.712 0C13.125.8 16 4.079 16 8.028a7.96 7.96 0 0 1-1.887 5.15c-2.791-1.35-4.744-4.406-4.756-7.966zm-1.384 9.875c-.497 2.056-1.981 3.813-3.828 4.981c1.138.622 2.441.978 3.828.978s2.691-.356 3.828-.978c-1.847-1.169-3.331-2.925-3.828-4.981"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 12.119V16l-6-6l6-6v3.966c6.98.164 6.681-4.747 4.904-7.966C16.29 4.741 15.359 12.337 7 12.119"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Road(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 16h5L12 0H9l.5 4h-3L7 0H4L0 16h5l.5-4h5zm-5.25-6l.5-4h3.5l.5 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1L6 6H3l-3 4s3.178-.885 5.032-.47L0 16l6.592-5.127C7.511 12.977 6 16 6 16l4-3v-3l5-5l1-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.13 11.733c-1.175 0-2.13.958-2.13 2.126c0 1.174.955 2.122 2.13 2.122a2.126 2.126 0 0 0 2.133-2.122a2.133 2.133 0 0 0-2.133-2.126M.002 5.436v3.067c1.997 0 3.874.781 5.288 2.196a7.45 7.45 0 0 1 2.192 5.302h3.08c0-5.825-4.739-10.564-10.56-10.564zM.006 0v3.068C7.128 3.068 12.924 8.87 12.924 16H16C16 7.18 8.824 0 .006 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M4.359 12.988c-.75 0-1.359-.603-1.359-1.353a1.36 1.36 0 0 1 2.718 0c0 .75-.609 1.353-1.359 1.353M7.772 13a4.753 4.753 0 0 0-1.397-3.381A4.74 4.74 0 0 0 3 8.219V6.263c3.713 0 6.738 3.022 6.738 6.737zm3.472 0c0-4.547-3.697-8.25-8.241-8.25V2.794c5.625 0 10.203 4.581 10.203 10.206h-1.963z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rtl(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0a4 4 0 0 0 0 8v8h2V2h2v14h2V2h2V0zm12 3l-4 4l4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sad(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13M4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m-5.002 7.199l-1.286-.772C4.586 9.973 6.179 9 8 9s3.413.973 4.288 2.427l-1.286.772C10.39 11.181 9.275 10.5 8 10.5s-2.389.681-3.002 1.699"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SadTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m3 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2M5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2m6.002 8.199C10.39 11.181 9.275 10.5 8 10.5s-2.389.681-3.002 1.699l-1.286-.772C4.586 9.973 6.179 9 8 9s3.414.973 4.288 2.427z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safari(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.581 0 0 3.581 0 8s3.581 8 8 8s8-3.581 8-8s-3.581-8-8-8m6.975 7.388l-.016-.166c.003.056.009.109.016.166M13.881 4.2l-.113-.169zm-.434-.597l-.069-.084c.025.028.047.056.069.084m-.969-.984l-.084-.069zm-.509-.388l-.169-.112zM8.778 1.044l-.169-.016c.056.003.113.009.169.016m-1.39-.019l-.169.016c.056-.003.112-.009.169-.016M4.2 2.119l-.169.112zm-.597.434l-.081.066zm-.984.969l-.069.084c.025-.028.047-.056.069-.084m-.388.509l-.112.169zM1.044 7.222l-.016.169c.003-.056.009-.112.016-.169m-.019 1.391l.016.169a2.413 2.413 0 0 0-.016-.169m1.094 3.184l.112.169a2.915 2.915 0 0 1-.112-.169m.131.197l1.247-.834l-.138-.209l-1.247.834A6.936 6.936 0 0 1 1.049 8.81l.747-.075l-.025-.25l-.747.075a7.016 7.016 0 0 1-.022-.438h1.5v-.25h-1.5a6.42 6.42 0 0 1 .022-.438l.747.072l.025-.25l-.747-.072a6.94 6.94 0 0 1 1.066-2.975l1.247.834l.138-.209l-1.25-.828c.084-.119.169-.237.259-.35l.578.475l.159-.194l-.578-.475c.094-.112.194-.219.294-.325l1.059 1.059l.178-.178L3.14 2.959c.106-.1.212-.2.322-.294l.475.581l.194-.159l-.475-.578c.116-.091.231-.178.35-.263l.834 1.247l.209-.138l-.834-1.247A6.941 6.941 0 0 1 7.19 1.045l.075.747l.25-.025l-.075-.747c.144-.012.291-.019.438-.022v1.5h.25v-1.5c.147.003.291.009.438.022l-.072.747l.25.025l.072-.747a6.94 6.94 0 0 1 2.975 1.066l-.834 1.247l.209.138L12 2.249c.119.084.238.169.35.259l-.475.578l.194.159l.475-.578c.113.094.219.194.325.294l-.4.391L7 6.999l-3.647 5.469l-.391.391a9.82 9.82 0 0 1-.294-.322l.578-.475l-.159-.194l-.578.475a6.874 6.874 0 0 1-.259-.35zm.369.484l-.066-.081zm.903.903l.081.066zm.509.385l.169.113a2.618 2.618 0 0 1-.169-.113m3.191 1.19l.169.016a2.471 2.471 0 0 1-.169-.016m1.391.019l.166-.016a2.387 2.387 0 0 0-.166.016m3.187-1.094l.169-.113zm.597-.434l.084-.069c-.028.025-.056.047-.084.069m.547-.491l.012-.012zm.437-.478l.069-.084c-.025.028-.047.056-.069.084m.11-.134l-.578-.475l-.159.194l.578.475c-.094.113-.194.219-.294.325l-1.059-1.059l-.178.178l1.059 1.059c-.106.1-.213.2-.322.294l-.475-.581l-.194.159l.475.578a8.61 8.61 0 0 1-.35.262l-.834-1.247l-.209.137l.834 1.247a6.936 6.936 0 0 1-2.975 1.063l-.075-.747l-.25.025l.075.747a7.016 7.016 0 0 1-.438.022v-1.5h-.25V15a6.42 6.42 0 0 1-.438-.022l.072-.747l-.25-.025l-.072.747a6.94 6.94 0 0 1-2.975-1.066l.834-1.247l-.209-.137l-.828 1.247a7.532 7.532 0 0 1-.35-.259l.475-.578l-.194-.159l-.475.578a7.556 7.556 0 0 1-.325-.294l.394-.391L9 9l3.647-5.469l.391-.391c.1.106.2.212.294.322l-.578.475l.159.194l.578-.475c.091.116.178.231.262.35l-1.247.834l.137.209l1.247-.834a6.941 6.941 0 0 1 1.063 2.975l-.747.075l.025.25l.747-.075c.012.144.019.291.022.438h-1.5v.25H15a6.42 6.42 0 0 1-.022.438l-.747-.072l-.025.25l.747.072a6.94 6.94 0 0 1-1.066 2.975l-1.247-.834l-.137.209L13.75 12c-.081.113-.169.228-.259.344m1.484-3.735c-.006.056-.009.113-.016.169zM13.881 11.8l-.113.169z"/><path fill="currentColor" d="m6.758 1.111l.293 1.471l-.245.049l-.293-1.471zM9.245 14.89l-.293-1.471l.245-.049l.293 1.471zM6.088 1.264l.218.718l-.239.073l-.218-.718zm3.825 13.469l-.218-.718l.239-.073l.218.718zM5.438 1.486l.574 1.386l-.231.096l-.574-1.386zm5.126 13.029l-.574-1.386l.231-.096l.574 1.386zM4.588 1.885l.22-.118l.354.661l-.22.118zm6.82 12.229l-.22.118l-.354-.661l.22-.118zM1.884 4.591l.662.353l-.118.221l-.661-.353zm12.229 6.818l-.662-.353l.118-.22l.662.353zM2.872 6.01l-1.386-.574l.096-.231l1.386.574zM13.13 9.989l1.386.574l-.096.231l-1.386-.574zM1.337 5.85l.718.218l-.073.239l-.718-.218zm13.324 4.301l-.718-.218l.073-.239l.718.218zM1.157 6.512l1.471.293l-.049.245l-1.471-.293zM14.84 9.488l-1.471-.293l.049-.245l1.471.293zM1.109 9.243L2.58 8.95l.049.245l-1.471.293zm13.779-2.486l-1.471.293l-.049-.245l1.471-.293zM1.265 9.914l.718-.218l.073.239l-.718.218zm13.469-3.825l-.718.218l-.073-.239l.718-.218zM1.58 10.796l-.096-.231l1.386-.574l.096.231zm12.839-5.589l.096.231l-1.386.574l-.096-.231zM1.888 11.41l-.118-.22l.661-.354l.118.22zm12.228-6.82l.118.22l-.661.354l-.118-.22zm-9.305 9.642l-.22-.118l.354-.661l.22.118zm6.378-12.465l.22.118l-.353.661l-.22-.118zM5.207 14.419l.574-1.386l.231.096l-.574 1.386zM10.795 1.58l-.574 1.386l-.231-.096l.574-1.386zM6.088 14.735l-.239-.073l.218-.718l.239.073zM9.912 1.264l.239.073l-.218.718l-.239-.073zM6.757 14.888l-.245-.049l.293-1.471l.245.049zM9.243 1.109l.245.049l-.293 1.471l-.245-.049z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.279 10.62c-1.042-1.628-2.829-2.345-3.992-1.601c-.1.064-.193.138-.277.218L8.769 7.295l2.867-4.5c.235-.433.321-.949.207-1.468A2.12 2.12 0 0 0 11.091.12l-.192-.122l-3.398 5.314L4.103-.002L3.911.12a2.12 2.12 0 0 0-.545 2.675l2.867 4.5l-1.241 1.942a1.834 1.834 0 0 0-.277-.218c-1.163-.744-2.95-.028-3.992 1.601s-.944 3.551.219 4.296c1.163.744 2.95.028 3.992-1.601l2.567-4.029l2.567 4.029c1.042 1.628 2.829 2.345 3.992 1.601s1.261-2.667.219-4.296M3.67 12.507c-.469.733-1.071 1.089-1.478 1.179c-.133.029-.317.047-.443-.033c-.139-.089-.231-.324-.247-.629c-.025-.494.151-1.076.483-1.594c.469-.733 1.071-1.089 1.478-1.179c.133-.029.317-.047.443.033c.139.089.231.324.247.629c.025.495-.151 1.076-.483 1.594M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m5.998 5.023c-.016.305-.108.54-.247.629c-.125.08-.31.062-.443.033c-.407-.089-1.009-.446-1.478-1.179c-.332-.519-.508-1.1-.483-1.594c.016-.305.108-.54.247-.629c.125-.08.31-.062.443-.033c.407.089 1.009.446 1.478 1.179c.332.519.508 1.1.483 1.594"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.504 13.616l-3.79-3.223c-.392-.353-.811-.514-1.149-.499a6 6 0 1 0-.672.672c-.016.338.146.757.499 1.149l3.223 3.79c.552.613 1.453.665 2.003.115s.498-1.452-.115-2.003zM6 10a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Section(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.749 16c-.771 0-1.424-.225-1.939-.669c-.519-.447-.782-.969-.782-1.552a.98.98 0 0 1 .307-.726c.207-.202.465-.309.748-.309c.281 0 .534.1.732.29c.195.187.294.435.294.736c0 .177-.029.372-.086.58a1.53 1.53 0 0 0-.068.364c0 .058.014.126.121.199c.199.138.439.204.732.204c.353 0 .667-.123.962-.375c.29-.249.431-.505.431-.782c0-.308-.082-.575-.252-.816c-.287-.402-.826-.874-1.603-1.401c-1.248-.835-2.079-1.559-2.54-2.211c-.358-.511-.539-1.061-.539-1.636c0-.579.19-1.155.564-1.713c.32-.477.794-.908 1.41-1.283c-.33-.355-.577-.689-.736-.995a2.565 2.565 0 0 1-.303-1.189c0-.747.295-1.393.878-1.92S7.39.001 8.241.001c.783 0 1.441.22 1.956.654c.521.439.785.952.785 1.524c0 .292-.109.553-.324.776l-.004.004c-.125.124-.353.271-.735.271c-.299 0-.561-.098-.758-.283a.876.876 0 0 1-.296-.656c0-.108.027-.272.084-.515c.028-.115.042-.221.042-.316a.476.476 0 0 0-.183-.39C8.679.962 8.494.909 8.243.909c-.389 0-.708.118-.975.361s-.399.533-.399.883c0 .315.071.574.212.771c.267.374.731.778 1.378 1.201c1.315.853 2.233 1.636 2.727 2.325c.365.518.549 1.068.549 1.637c0 .572-.186 1.148-.552 1.714c-.316.487-.793.926-1.42 1.308c.347.367.591.688.743.977c.189.359.284.751.284 1.165c0 .776-.296 1.435-.879 1.96s-1.31.79-2.161.79zM6.975 5.568c-.753.452-1.12.972-1.12 1.583c0 .356.102.674.31.973c.311.436.926.97 1.825 1.583c.381.259.724.511 1.025.751c.767-.461 1.14-.974 1.14-1.565c0-.322-.127-.668-.378-1.03c-.263-.378-.826-.872-1.674-1.467a21.316 21.316 0 0 1-1.128-.827z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 10s.919-3 6-3v3l6-4l-6-4v3c-4 0-6 2.495-6 5m7 2H2V6h1.967c.158-.186.327-.365.508-.534A6.933 6.933 0 0 1 6.914 4H0v10h13V9.803l-2 1.333z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 11c-.706 0-1.342.293-1.797.763L4.969 8.396a2.46 2.46 0 0 0 0-.792l6.734-3.367a2.5 2.5 0 1 0-.672-1.341L4.297 6.263a2.5 2.5 0 1 0 0 3.474l6.734 3.367A2.5 2.5 0 1 0 13.5 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 0L8 2L1 0S.93.808 1 2l7 2.189L15 2c.07-1.192 0-2 0-2M1.128 3.049C1.503 6.966 2.901 13.553 8 16c5.099-2.448 6.497-9.034 6.872-12.951L8 5.633z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shift(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 14h-5a.5.5 0 0 1-.5-.5V8H3a.5.5 0 0 1-.354-.854l5-5a.5.5 0 0 1 .707 0l5 5a.499.499 0 0 1-.354.854h-2v5.5a.5.5 0 0 1-.5.5zM6 13h4V7.5a.5.5 0 0 1 .5-.5h1.293L8 3.207L4.207 7H5.5a.5.5 0 0 1 .5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shocked(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13M6 11a2 2 0 1 1 3.999-.001A2 2 0 0 1 6 11m4-5.5c0-.828.448-1.5 1-1.5s1 .672 1 1.5S11.552 7 11 7s-1-.672-1-1.5m-6 0C4 4.672 4.448 4 5 4s1 .672 1 1.5S5.552 7 5 7s-1-.672-1-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShockedTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M5 7c-.552 0-1-.672-1-1.5S4.448 4 5 4s1 .672 1 1.5S5.552 7 5 7m3 6a2 2 0 1 1-.001-3.999A2 2 0 0 1 8 13m3-6c-.552 0-1-.672-1-1.5S10.448 4 11 4s1 .672 1 1.5S11.552 7 11 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shrink(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 7h6.5L13 4.5l3-3L14.5 0l-3 3L9 .5zm0 2v6.5l2.5-2.5l3 3l1.5-1.5l-3-3L15.5 9zM7 9H.5L3 11.5l-3 3L1.5 16l3-3L7 15.5zm0-2V.5L4.5 3l-3-3L0 1.5l3 3L.5 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShrinkTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 9v6.5L4.5 13l-3 3L0 14.5l3-3L.5 9zm9-7.5l-3 3L15.5 7H9V.5L11.5 3l3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11h-1.586l-2.5-2.5l2.5-2.5H12v2.5L15.5 5L12 1.5V4h-2a.997.997 0 0 0-.707.293L6.5 7.086L3.707 4.293A1 1 0 0 0 3 4H0v2h2.586l2.5 2.5l-2.5 2.5H0v2h3c.265 0 .52-.105.707-.293L6.5 9.914l2.793 2.793A1 1 0 0 0 10 13h2v2.5l3.5-3.5L12 8.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sigma(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.713 11.48l.694-1.48h.594l-1 6h-15v-1.16l5.18-6.113l-5.18-5.18V0h15.313l.688 4h-.537l-.293-.607C14.62 2.247 14.205 2 13.002 2H2.658l5.517 5.516l-4.647 5.483h8.474c1.813 0 2.291-.65 2.713-1.52z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SinaWeibo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.722 14.031c-2.65.262-4.938-.938-5.109-2.675c-.172-1.741 1.837-3.359 4.484-3.622c2.65-.263 4.938.938 5.106 2.675c.175 1.741-1.834 3.362-4.481 3.622m5.297-5.772c-.225-.069-.381-.113-.262-.409c.256-.644.281-1.197.003-1.594c-.519-.741-1.941-.703-3.569-.019c0 0-.513.222-.381-.181c.25-.806.213-1.478-.178-1.869c-.884-.884-3.234.034-5.25 2.05C.876 7.74.001 9.343.001 10.728c0 2.644 3.394 4.253 6.713 4.253c4.35 0 7.247-2.528 7.247-4.534c0-1.216-1.022-1.903-1.941-2.188zm2.89-4.843a4.234 4.234 0 0 0-4.031-1.306a.61.61 0 1 0 .256 1.194a3.01 3.01 0 0 1 3.494 3.872a.615.615 0 0 0 .394.772a.615.615 0 0 0 .772-.394v-.003a4.222 4.222 0 0 0-.884-4.134z"/><path fill="currentColor" d="M13.294 4.875a2.055 2.055 0 0 0-1.963-.634a.527.527 0 1 0 .219 1.031c.341-.072.709.034.959.309c.25.278.319.656.209.987a.532.532 0 0 0 .341.666a.531.531 0 0 0 .666-.341a2.062 2.062 0 0 0-.431-2.019zm-6.425 6.009c-.094.159-.297.234-.456.169c-.159-.063-.206-.244-.116-.397c.094-.153.291-.228.447-.169c.156.056.213.234.125.397m-.847 1.082c-.256.409-.806.588-1.219.4c-.406-.184-.528-.659-.272-1.059c.253-.397.784-.575 1.194-.403c.416.178.55.65.297 1.063zm.962-2.894c-1.259-.328-2.684.3-3.231 1.409c-.559 1.131-.019 2.391 1.253 2.803c1.319.425 2.875-.228 3.416-1.447c.534-1.197-.131-2.425-1.438-2.766z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.65.584C6.625.568 6.597.556 6.572.543c-.028.006-.053.009-.081.016zM.575 6.578c-.006.028-.009.056-.012.081c.016.025.025.05.041.075l-.028-.156zm14.844 2.838c.006-.028.009-.056.016-.084c-.016-.025-.025-.05-.041-.075zM9.25 15.359c.025.016.053.028.078.041c.028-.006.056-.009.084-.012z"/><path fill="currentColor" d="M15.434 9.331c-.006.028-.009.056-.016.084l-.028-.162c.016.028.028.053.044.078c.081-.45.125-.909.125-1.369a7.54 7.54 0 0 0-2.213-5.344a7.56 7.56 0 0 0-2.4-1.619A7.51 7.51 0 0 0 8.005.405c-.481 0-.963.044-1.431.134h-.003c.025.012.053.025.078.041L6.49.555c.028-.006.053-.009.081-.016A4.492 4.492 0 0 0 4.474.014C3.28.014 2.155.48 1.311 1.323S.002 3.292.002 4.486c0 .759.197 1.509.563 2.169c.006-.028.009-.056.012-.081l.028.159c-.016-.025-.028-.05-.041-.075a7.566 7.566 0 0 0-.112 1.303a7.514 7.514 0 0 0 2.213 5.341a7.54 7.54 0 0 0 6.666 2.094a.64.64 0 0 0-.078-.041l.162.028c-.028.006-.056.009-.084.012a4.479 4.479 0 0 0 2.2.581c1.194 0 2.319-.466 3.162-1.309s1.309-1.969 1.309-3.162a4.532 4.532 0 0 0-.569-2.175zm-7.4 3.26c-2.684 0-3.884-1.319-3.884-2.309c0-.506.375-.863.891-.863c1.15 0 .85 1.65 2.994 1.65c1.097 0 1.703-.597 1.703-1.206c0-.366-.181-.772-.903-.95l-2.388-.597c-1.922-.481-2.272-1.522-2.272-2.5c0-2.028 1.909-2.791 3.703-2.791c1.653 0 3.6.913 3.6 2.131c0 .522-.453.825-.969.825c-.981 0-.8-1.356-2.775-1.356c-.981 0-1.522.444-1.522 1.078s.775.838 1.447.991l1.769.394c1.934.431 2.425 1.563 2.425 2.625c0 1.647-1.266 2.878-3.819 2.878"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sleepy(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13"/><path fill="currentColor" d="M10 10.5c0 1.381-.895 2.5-2 2.5s-2-1.119-2-2.5S6.895 8 8 8s2 1.119 2 2.5M6.5 5.313a.502.502 0 0 1-.354-.146c-.302-.302-.991-.302-1.293 0a.5.5 0 0 1-.707-.707c.696-.696 2.011-.696 2.707 0a.5.5 0 0 1-.354.853zm5 0a.502.502 0 0 1-.354-.146c-.302-.302-.991-.302-1.293 0a.5.5 0 0 1-.707-.707c.696-.696 2.011-.696 2.707 0a.5.5 0 0 1-.354.853z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SleepyTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M4.854 5.166a.5.5 0 0 1-.707-.707c.696-.696 2.011-.696 2.707 0a.5.5 0 0 1-.708.707c-.302-.302-.991-.302-1.293 0zM8 13c-1.105 0-2-1.119-2-2.5S6.895 8 8 8s2 1.119 2 2.5S9.105 13 8 13m3.854-7.834a.498.498 0 0 1-.708 0c-.302-.302-.991-.302-1.293 0a.5.5 0 0 1-.707-.707c.696-.696 2.011-.696 2.707 0a.5.5 0 0 1 0 .707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smile(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13M4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m1.002 4.801l1.286.772C11.414 12.027 9.821 13 8 13s-3.413-.973-4.288-2.427l1.286-.772C5.61 10.819 6.725 11.5 8 11.5s2.389-.681 3.002-1.699"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m3 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2M5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2m3 9a4.999 4.999 0 0 1-4.288-2.427l1.286-.772C5.61 10.819 6.725 11.5 8 11.5s2.389-.681 3.002-1.699l1.286.772A4.996 4.996 0 0 1 8 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaAsc(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12zm9.5 4h-4a.5.5 0 0 1-.416-.777L13.566 10H10.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 .416.777L11.434 15H14.5a.5.5 0 0 1 0 1m1.447-9.724l-3-6a.5.5 0 0 0-.894 0l-3 6a.5.5 0 0 0 .895.447l.862-1.724h3.382l.862 1.724a.5.5 0 0 0 .895-.447zM11.309 4L12.5 1.618L13.691 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaDesc(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12zm9.5-5h-4a.5.5 0 0 1-.416-.777L13.566 1H10.5a.5.5 0 0 1 0-1h4a.502.502 0 0 1 .416.777L11.434 6H14.5a.5.5 0 0 1 0 1m1.447 8.276l-3-6a.5.5 0 0 0-.894 0l-3 6a.5.5 0 0 0 .895.447l.862-1.724h3.382l.862 1.724a.5.5 0 0 0 .895-.447zM11.309 13l1.191-2.382L13.691 13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountAsc(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12z"/><path fill="currentColor" d="M7 9h9v2H7zm0-3h7v2H7zm0-3h5v2H7zm0-3h3v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountDesc(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12zM7 0h9v2H7zm0 3h7v2H7zm0 3h5v2H7z"/><path fill="currentColor" d="M7 9h3v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumbericDesc(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12zm8.5 4a.5.5 0 0 1-.5-.5V10h-.5a.5.5 0 0 1 0-1h1a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5m1-16h-3a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5H14v2h-2.5a.5.5 0 0 0 0 1h3a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5M12 1h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericAsc(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12V0H3v12H.5L4 15.5L7.5 12zm8.5-5a.5.5 0 0 1-.5-.5V1h-.5a.5.5 0 0 1 0-1h1a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5m1 2h-3a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5H14v2h-2.5a.5.5 0 0 0 0 1h3a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5M12 10h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soundcloud(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.937 8.034c-.283 0-.552.055-.798.154C12.975 6.401 11.416 5 9.514 5c-.465 0-.917.088-1.317.237c-.156.058-.197.117-.197.233v6.292c0 .121.098.222.221.234l5.717.003c1.139 0 2.062-.888 2.062-1.983s-.924-1.983-2.063-1.983zM6.25 12h.5L7 8.497L6.75 5h-.5L6 8.497zm-1.5 0h-.5L4 9.457L4.25 7h.5L5 9.5zm-2.5 0h.5L3 10l-.25-2h-.5L2 10zm-2-1h.5L1 10L.75 9h-.5L0 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundcloudTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M2.75 11h-.5L2 9.5L2.25 8h.5L3 9.5zm2 0h-.5L4 9l.25-2h.5L5 9zm2 0h-.5L6 8l.25-3h.5L7 8zm6.144 0l-4.709-.003a.208.208 0 0 1-.184-.2V5.403c0-.1.034-.15.162-.2a3 3 0 0 1 4.075 2.531a1.701 1.701 0 0 1 2.356 1.569c0 .938-.762 1.697-1.7 1.697"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spades(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.775 5.44C9.751 3.192 8.708 1.393 8.001 0c-.708 1.393-1.75 3.192-4.774 5.44c-5.157 3.833-.303 9.182 3.965 6.238c-.278 1.827-1.227 3.159-2.191 3.733v.59h6v-.59c-.964-.574-1.913-1.906-2.191-3.733c4.268 2.944 9.122-2.405 3.965-6.238"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpellCheck(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h2v3h1V1c0-.55-.45-1-1-1H2c-.55 0-1 .45-1 1v6h1zm0-3h2v2H2zm13 0V0h-3c-.55 0-1 .45-1 1v5c0 .55.45 1 1 1h3V6h-3V1zm-5 1.5V1c0-.55-.45-1-1-1H6v7h3c.55 0 1-.45 1-1V4.5c0-.55-.137-1-.688-1c.55 0 .688-.45.688-1M9 6H7V4h2zm0-3H7V1h2zm4 6l-6.5 7L3 11.5l1.281-1.094L6.5 12.719L12 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sphere(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 1a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15m4.244 10c.134-.632.219-1.303.246-2h1.991a6.443 6.443 0 0 1-.479 2zM3.256 6a11.876 11.876 0 0 0-.246 2H1.019a6.443 6.443 0 0 1 .479-2zm7.463 0c.15.64.241 1.31.27 2H8V6zM8 5V2.073c.228.066.454.178.675.334c.415.293.813.744 1.149 1.304c.233.388.434.819.601 1.289H7.999zM5.176 3.711c.336-.561.734-1.012 1.149-1.304c.222-.156.447-.268.675-.334V5H4.574a7.29 7.29 0 0 1 .601-1.289zM7 6v2H4.011c.029-.69.12-1.36.27-2zm-5.502 5a6.443 6.443 0 0 1-.479-2H3.01c.028.697.112 1.368.246 2zm2.513-2H7v2H4.281c-.15-.64-.241-1.31-.27-2M7 12v2.927a2.27 2.27 0 0 1-.675-.334c-.415-.293-.813-.744-1.149-1.304A7.221 7.221 0 0 1 4.574 12zm2.825 1.289c-.336.561-.734 1.012-1.149 1.304a2.282 2.282 0 0 1-.675.334V12h2.426c-.168.47-.369.901-.602 1.289M8 11V9h2.989c-.029.69-.12 1.36-.27 2zm3.99-3a11.98 11.98 0 0 0-.246-2h1.758c.267.639.427 1.309.479 2zm.989-3h-1.498c-.291-.918-.693-1.723-1.177-2.366a6.462 6.462 0 0 1 1.792 1.27c.336.336.631.702.883 1.096M2.904 3.904a6.492 6.492 0 0 1 1.792-1.27C4.213 3.277 3.81 4.082 3.519 5H2.021c.252-.394.547-.761.883-1.096M2.021 12h1.498c.291.918.693 1.723 1.177 2.366a6.462 6.462 0 0 1-1.792-1.27A6.505 6.505 0 0 1 2.021 12m10.075 1.096a6.492 6.492 0 0 1-1.792 1.27c.483-.643.886-1.448 1.177-2.366h1.498a6.466 6.466 0 0 1-.883 1.096"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2a2 2 0 1 1 3.999-.001A2 2 0 0 1 6 2m4.243 1.757a2 2 0 1 1 3.999-.001a2 2 0 0 1-3.999.001M13 8a1 1 0 1 1 2 0a1 1 0 0 1-2 0m-1.757 4.243a1 1 0 1 1 2 0a1 1 0 0 1-2 0M7 14a1 1 0 0 1 2 0a1 1 0 0 1-2 0m-4.243-1.757a1 1 0 0 1 2 0a1 1 0 0 1-2 0m-.5-8.486a1.5 1.5 0 0 1 3 0a1.5 1.5 0 0 1-3 0M.875 8a1.125 1.125 0 1 1 2.25 0a1.125 1.125 0 0 1-2.25 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerEight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16c-2.137 0-4.146-.832-5.657-2.343S0 10.137 0 8c0-1.513.425-2.986 1.228-4.261A8.02 8.02 0 0 1 4.421.844l.672 1.341a6.53 6.53 0 0 0-2.596 2.354A6.48 6.48 0 0 0 1.5 8c0 3.584 2.916 6.5 6.5 6.5s6.5-2.916 6.5-6.5c0-1.23-.345-2.426-.997-3.461a6.515 6.515 0 0 0-2.596-2.354l.672-1.341a8.02 8.02 0 0 1 3.193 2.895A7.981 7.981 0 0 1 16 8c0 2.137-.832 4.146-2.343 5.657S10.137 16 8 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerEleven(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6h-6l2.243-2.243C11.11 2.624 9.603 2 8 2s-3.109.624-4.243 1.757C2.624 4.89 2 6.397 2 8s.624 3.109 1.757 4.243A5.961 5.961 0 0 0 8 14a5.963 5.963 0 0 0 4.516-2.049l1.505 1.317a8 8 0 1 1-.364-10.924L16 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerFive(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 4a4 4 0 1 1 0 8a4 4 0 0 1 0-8m4.773 8.773C11.498 14.048 9.803 14.75 8 14.75s-3.498-.702-4.773-1.977S1.25 9.803 1.25 8c0-1.803.702-3.498 1.977-4.773l1.061 1.061a5.256 5.256 0 0 0 0 7.425c.992.992 2.31 1.538 3.712 1.538s2.721-.546 3.712-1.538a5.256 5.256 0 0 0 0-7.425l1.061-1.061C14.048 4.502 14.75 6.197 14.75 8s-.702 3.498-1.977 4.773"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerFour(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8c0-.19.011-.378.032-.563l-2.89-.939A8.085 8.085 0 0 0 .001 8c0 2.3.971 4.374 2.526 5.833l1.786-2.458A4.982 4.982 0 0 1 3.001 8zm10 0a4.978 4.978 0 0 1-1.312 3.375l1.786 2.458A7.975 7.975 0 0 0 16 8c0-.513-.049-1.015-.141-1.502l-2.89.939c.021.185.032.373.032.563zM9 3.1a5.01 5.01 0 0 1 3.351 2.435l2.89-.939A8.01 8.01 0 0 0 9 .062zM3.649 5.535A5.007 5.007 0 0 1 7 3.1V.062A8.005 8.005 0 0 0 .759 4.596zm6.422 7.017C9.44 12.84 8.739 13 8 13s-1.44-.16-2.071-.448L4.143 15.01C5.287 15.641 6.601 16 8 16s2.713-.359 3.857-.99z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerNine(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0A8 8 0 0 0 .002 7.812C.094 4.033 2.968 1 6.5 1C10.09 1 13 4.134 13 8a1.5 1.5 0 0 0 3 0a8 8 0 0 0-8-8m0 16a8 8 0 0 0 7.998-7.812C15.906 11.967 13.032 15 9.5 15C5.91 15 3 11.866 3 8a1.5 1.5 0 0 0-3 0a8 8 0 0 0 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerSeven(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 14.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 6.5 14.5M0 8a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 0 8m13 0a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 13 8M1.904 3.404a1.5 1.5 0 1 1 3.001.001a1.5 1.5 0 0 1-3.001-.001m9.192 9.192a1.5 1.5 0 1 1 3.001.001a1.5 1.5 0 0 1-3.001-.001m-9.192 0a1.5 1.5 0 1 1 3.001.001a1.5 1.5 0 0 1-3.001-.001m9.192-9.192a1.5 1.5 0 1 1 3.001.001a1.5 1.5 0 0 1-3.001-.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerSix(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2a2 2 0 1 1 3.999-.001A2 2 0 0 1 6 2m6.359 6a1.641 1.641 0 0 1 3.282 0a1.641 1.641 0 0 1-3.282 0m-1.602 4.243a1.486 1.486 0 1 1 2.971 0a1.486 1.486 0 0 1-2.971 0M6.654 14a1.346 1.346 0 1 1 2.693.001A1.346 1.346 0 0 1 6.654 14m-4.116-1.757a1.22 1.22 0 1 1 2.439 0a1.22 1.22 0 0 1-2.439 0M.896 8a1.104 1.104 0 1 1 2.207 0A1.104 1.104 0 0 1 .896 8m1.861-4.243a1 1 0 0 1 2 0a1 1 0 0 1-2 0m11.297 0a1.811 1.811 0 1 1-3.623 0a1.811 1.811 0 0 1 3.623 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerTen(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.001 8.025v.003c.002.061.009.12.021.177l.011.09l.011.095l.019.128l.045.296l.068.326c.011.058.028.117.044.178l.049.188c.009.032.016.065.027.097l.031.098l.065.203l.017.052l.019.052l.039.105l.081.215l.094.218l.048.111c.016.037.035.073.053.111l.11.224c.039.075.081.149.123.224l.063.113c.021.038.045.074.068.112l.14.224c.198.295.417.587.66.864c.245.275.511.535.792.775c.284.236.582.452.886.642c.306.188.619.349.928.487l.232.095c.038.015.076.032.115.046l.115.041c.038.014.151.054.226.078l.224.066c.037.011.073.022.109.031l.109.027l.213.052l.207.04l.101.019l.049.009l.049.007l.192.027l.093.013l.091.009l.174.015c.056.005.111.011.164.012h.011a1 1 0 0 0 1 .982l.025-.001h.004a1.02 1.02 0 0 0 .177-.021l.09-.011l.095-.011l.128-.019l.296-.045l.326-.068c.058-.011.117-.028.178-.044l.188-.049c.032-.009.065-.016.097-.027l.098-.031l.203-.065l.052-.017l.052-.019l.105-.039l.215-.081l.218-.094l.111-.048c.037-.016.073-.035.111-.053l.224-.11c.075-.039.149-.081.224-.123l.113-.063c.038-.021.074-.045.112-.068l.224-.14c.295-.197.587-.417.864-.66c.275-.245.535-.511.775-.792c.236-.284.452-.582.642-.886c.188-.306.349-.619.487-.928l.095-.232c.015-.038.032-.076.046-.115l.04-.115c.013-.038.054-.151.078-.226l.066-.224c.011-.037.022-.073.031-.109l.027-.109l.052-.213l.04-.207l.019-.101l.009-.049a.753.753 0 0 0 .007-.05l.027-.192l.013-.093l.009-.091l.015-.174c.005-.056.011-.111.012-.165l.001-.025a1 1 0 0 0 .996-1l-.001-.025v-.003a1.02 1.02 0 0 0-.021-.177l-.011-.09l-.011-.095l-.019-.128l-.045-.296l-.068-.326c-.011-.058-.028-.117-.044-.178l-.049-.188c-.009-.032-.016-.065-.027-.097l-.031-.098l-.065-.203l-.017-.052l-.019-.052l-.039-.105l-.081-.215l-.094-.218l-.048-.111c-.016-.037-.035-.073-.053-.111l-.11-.224c-.039-.075-.081-.149-.123-.224l-.063-.113c-.021-.038-.045-.074-.068-.112l-.14-.224a8.307 8.307 0 0 0-.66-.864a8.164 8.164 0 0 0-.792-.775a8.155 8.155 0 0 0-.886-.642a8.024 8.024 0 0 0-.928-.487l-.232-.095c-.038-.015-.076-.032-.115-.046l-.115-.04c-.038-.013-.151-.054-.226-.078l-.224-.066c-.037-.01-.073-.022-.109-.031l-.109-.027a54.114 54.114 0 0 0-.213-.052l-.207-.04l-.101-.019l-.049-.009l-.049-.007l-.192-.027l-.093-.013l-.091-.009l-.174-.015C9.136.987 9.081.981 9.028.98L8.989.979a1 1 0 0 0-.999-.981l-.025.001h-.003a1.02 1.02 0 0 0-.177.021l-.09.011L7.6.042l-.128.019l-.296.045l-.326.068c-.058.011-.117.028-.178.044l-.188.049c-.032.009-.065.016-.097.027l-.098.031l-.203.065l-.052.017l-.052.019l-.105.039l-.215.081l-.218.094l-.111.048c-.037.016-.073.035-.111.053l-.224.11c-.075.039-.149.081-.224.123l-.113.063c-.038.021-.074.045-.112.068l-.224.14a8.448 8.448 0 0 0-.864.66a8.164 8.164 0 0 0-.775.792a8.155 8.155 0 0 0-.642.886a8.024 8.024 0 0 0-.487.928l-.095.232c-.015.038-.032.076-.046.115l-.04.115c-.013.038-.054.151-.078.226l-.066.224l-.032.109c-.01.036-.018.073-.027.109l-.052.213l-.04.207l-.019.101l-.009.049a.753.753 0 0 0-.007.05l-.027.192l-.013.093l-.009.091l-.015.174c-.005.056-.011.111-.012.165l-.001.025A1 1 0 0 0 .002 8l.001.025zm1.148-1.014l.002-.009c.01-.051.026-.102.04-.155l.045-.163l.024-.084l.028-.086l.058-.176l.015-.045l.017-.045l.035-.091l.073-.186l.084-.189l.043-.096l.048-.096l.098-.194c.034-.065.073-.128.109-.194l.056-.098c.019-.033.04-.064.061-.096l.124-.194a7.38 7.38 0 0 1 .583-.744c.217-.236.451-.459.697-.665c.25-.202.511-.385.776-.547c.268-.159.541-.294.808-.41l.202-.079c.033-.013.066-.027.099-.038l.1-.033c.033-.011.131-.045.196-.065l.194-.054c.032-.009.063-.019.095-.026l.094-.021l.184-.042l.179-.032l.087-.015l.043-.008c.014-.003.029-.003.043-.005l.166-.02l.08-.01l.078-.006l.15-.011c.049-.003.095-.008.142-.008l.256-.006l.2.007l.085.002a11.23 11.23 0 0 1 .34.024l.022-.001h.004a1 1 0 0 0 .962-.84l.025.006c.051.01.102.026.155.04l.163.045l.084.024l.086.028l.176.058l.045.015l.045.017l.091.035l.186.073l.189.084l.096.043l.096.048l.194.098c.065.034.129.073.194.109l.098.056c.033.019.064.04.096.061l.194.124c.255.176.506.369.744.583c.236.217.459.451.665.697c.202.25.385.511.547.776c.159.268.294.541.41.808l.079.202c.013.033.027.066.038.099l.033.1c.011.033.045.131.065.196l.054.194c.009.032.019.063.026.095l.021.094l.042.184l.032.179l.015.087l.008.043c.003.014.003.029.005.043l.02.166l.01.08l.006.078l.011.15c.003.049.008.095.008.142l.006.256l-.007.2l-.002.085a11.23 11.23 0 0 1-.024.34l.001.022v.004a1 1 0 0 0 .823.959l-.003.014c-.01.051-.025.102-.04.155l-.045.163l-.024.084l-.028.086l-.058.176l-.015.045l-.017.045l-.035.091l-.073.186l-.084.189l-.043.096l-.048.096l-.098.194c-.034.065-.073.129-.109.194l-.056.098c-.019.033-.04.064-.061.096l-.124.194a7.38 7.38 0 0 1-.583.744a7.422 7.422 0 0 1-.697.665c-.25.202-.511.385-.776.547a7.227 7.227 0 0 1-.808.41l-.202.079c-.033.013-.066.027-.099.038l-.1.033c-.033.011-.131.045-.196.065l-.194.054c-.032.009-.063.019-.095.026l-.094.021l-.184.042l-.179.032l-.087.015l-.043.008c-.015.003-.029.003-.043.005l-.166.02l-.08.01l-.078.006l-.15.011c-.049.003-.095.008-.142.008a77.64 77.64 0 0 0-.256.006l-.2-.007l-.085-.002a11.23 11.23 0 0 1-.34-.024l-.022.001H7.96a1.001 1.001 0 0 0-.961.834c-.05-.01-.101-.025-.153-.039l-.163-.045l-.084-.024l-.086-.028l-.176-.058l-.045-.015l-.045-.017l-.091-.035a31.884 31.884 0 0 1-.186-.073a17.554 17.554 0 0 0-.189-.084l-.096-.043l-.096-.048l-.194-.098c-.065-.034-.129-.073-.194-.109l-.098-.056c-.033-.019-.064-.04-.096-.061l-.194-.124a7.238 7.238 0 0 1-.744-.583a7.422 7.422 0 0 1-.665-.697a7.276 7.276 0 0 1-.547-.776a7.227 7.227 0 0 1-.41-.808l-.079-.202c-.013-.033-.027-.066-.038-.099l-.033-.1c-.011-.033-.045-.131-.065-.196l-.054-.194c-.009-.032-.019-.063-.026-.095l-.021-.094a78.832 78.832 0 0 0-.042-.184l-.032-.179l-.015-.087l-.008-.043c-.003-.015-.003-.029-.005-.043a12.365 12.365 0 0 1-.02-.166l-.01-.08l-.006-.078l-.011-.15c-.003-.049-.008-.095-.008-.142a77.64 77.64 0 0 0-.006-.256l.007-.2l.002-.085a11.23 11.23 0 0 1 .024-.34L2 7.98v-.003a1 1 0 0 0-.851-.964z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4.736a.933.933 0 0 1-.933-.933V1.005a.933.933 0 0 1 1.866 0v2.798A.933.933 0 0 1 8 4.736m0 10.841a.583.583 0 0 1-.583-.583v-2.798a.583.583 0 1 1 1.166 0v2.798a.583.583 0 0 1-.583.583M5.902 5.24a.875.875 0 0 1-.758-.437L3.745 2.38a.874.874 0 0 1 1.514-.874l1.399 2.423a.874.874 0 0 1-.756 1.311m5.596 9.342a.525.525 0 0 1-.455-.262l-1.399-2.423a.525.525 0 1 1 .909-.525l1.399 2.423a.525.525 0 0 1-.454.787M4.365 6.718a.813.813 0 0 1-.407-.109L1.535 5.21a.817.817 0 0 1 .816-1.414l2.423 1.399a.817.817 0 0 1-.408 1.523zm9.692 5.246a.461.461 0 0 1-.233-.063l-2.423-1.399a.467.467 0 0 1 .466-.808l2.423 1.399a.467.467 0 0 1-.233.87zM3.803 8.758H1.005a.758.758 0 0 1 0-1.516h2.798a.758.758 0 1 1 0 1.516m11.192-.292h-2.798a.466.466 0 1 1 0-.932h2.798a.466.466 0 1 1 0 .932M1.943 12.197a.698.698 0 0 1-.35-1.305l2.423-1.399a.698.698 0 1 1 .699 1.211l-2.423 1.399a.69.69 0 0 1-.349.094m9.692-5.829a.466.466 0 0 1-.233-.87l2.423-1.399a.466.466 0 1 1 .466.808l-2.423 1.399a.469.469 0 0 1-.233.063zm-7.133 8.331a.64.64 0 0 1-.555-.962l1.399-2.423a.641.641 0 1 1 1.111.641l-1.399 2.423a.64.64 0 0 1-.556.321m5.596-9.867a.467.467 0 0 1-.404-.7l1.399-2.423a.467.467 0 0 1 .808.466l-1.399 2.423a.464.464 0 0 1-.404.233z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8a7.917 7.917 0 0 0-2.431-5.568A7.776 7.776 0 0 0 5.057.896c-.923.405-1.758.992-2.449 1.712S1.371 4.182 1.011 5.105a7.531 7.531 0 0 0 .115 5.742c.392.892.961 1.7 1.658 2.368S4.307 14.41 5.2 14.758a7.286 7.286 0 0 0 2.799.493a7.157 7.157 0 0 0 6.526-4.547a6.98 6.98 0 0 0 .415-1.622l.059.002a1 1 0 0 0 .996-1.083h.004zm-1.589 2.655c-.367.831-.898 1.584-1.55 2.206s-1.422 1.112-2.254 1.434a6.759 6.759 0 0 1-2.608.454a6.676 6.676 0 0 1-4.685-2.065a6.597 6.597 0 0 1-1.38-2.173a6.514 6.514 0 0 1 .116-4.976c.342-.77.836-1.468 1.441-2.044s1.321-1.029 2.092-1.326c.771-.298 1.596-.438 2.416-.416s1.629.202 2.368.532c.74.329 1.41.805 1.963 1.387s.988 1.27 1.272 2.011a6.02 6.02 0 0 1 .397 2.32h.004a1 1 0 0 0 .888 1.077a6.872 6.872 0 0 1-.481 1.578z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpoonKnife(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 0c-1.657 0-3 1.567-3 3.5c0 1.655.985 3.042 2.308 3.406l-.497 8.096A.931.931 0 0 0 3.25 16h.5c.55 0 .972-.449.939-.998l-.497-8.096C5.515 6.541 6.5 5.155 6.5 3.5c0-1.933-1.343-3.5-3-3.5m10.083 0l-.833 5h-.625l-.417-5h-.417l-.417 5h-.625l-.833-5h-.417v6.5a.5.5 0 0 0 .5.5h1.302l-.491 8.002a.931.931 0 0 0 .939.998h.5c.55 0 .972-.449.939-.998L12.197 7h1.302a.5.5 0 0 0 .5-.5V0h-.417z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotify(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.559-8-8-8m3.681 11.559c-.159.241-.441.319-.681.159c-1.881-1.159-4.241-1.4-7.041-.759c-.281.081-.519-.119-.6-.359c-.081-.281.119-.519.359-.6c3.041-.681 5.681-.4 7.759.881c.281.119.322.438.203.678zm.96-2.2c-.2.281-.559.4-.841.2C9.641 8.24 6.359 7.84 3.841 8.64c-.319.081-.681-.081-.759-.4c-.081-.319.081-.681.4-.759c2.919-.881 6.519-.441 9 1.081c.238.119.359.519.159.797m.078-2.24C10.16 5.6 5.878 5.438 3.438 6.2a.749.749 0 0 1-.919-.481c-.119-.4.119-.8.481-.919c2.841-.841 7.519-.681 10.481 1.081c.359.2.481.681.281 1.041c-.203.278-.681.397-1.044.197z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stack(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5L8 1L0 5l8 4zM8 2.328L13.345 5L8 7.672L2.655 5zm6.398 4.871L16 8l-8 4l-8-4l1.602-.801L8 10.398zm0 3L16 11l-8 4l-8-4l1.602-.801L8 13.398z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stackoverflow(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10v6H0v-6h2v4h12v-4zM3 11h10v2H3zm.237-2.165l.433-1.953l9.763 2.164L13 10.999zM4.37 4.821l.845-1.813l9.063 4.226l-.845 1.813zm11.126.827l-1.218 1.587l-7.934-6.088L7.224 0h.91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarEmpty(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 6.204l-5.528-.803L8 .392L5.528 5.401L0 6.204l4 3.899l-.944 5.505L8 13.009l4.944 2.599L12 10.103zm-8 5.569l-3.492 1.836l.667-3.888L2.35 6.968l3.904-.567L8 2.864l1.746 3.537l3.904.567l-2.825 2.753l.667 3.888z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFull(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 6.204l-5.528-.803L8 .392L5.528 5.401L0 6.204l4 3.899l-.944 5.505L8 13.009l4.944 2.599L12 10.103z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 6.204l-5.528-.803L8 .392L5.528 5.401L0 6.204l4 3.899l-.944 5.505L8 13.009l4.944 2.599L12 10.103zm-8 5.569l-.015.008L8 2.863L9.746 6.4l3.904.567l-2.825 2.753l.667 3.888z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatsBars(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 13h16v2H0zm2-4h2v3H2zm3-4h2v7H5zm3 3h2v4H8zm3-6h2v10h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatsBarsTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 6h-3c-.275 0-.5.225-.5.5v9c0 .275.225.5.5.5h3c.275 0 .5-.225.5-.5v-9c0-.275-.225-.5-.5-.5m0 9h-3v-4h3zm5-11h-3c-.275 0-.5.225-.5.5v11c0 .275.225.5.5.5h3c.275 0 .5-.225.5-.5v-11c0-.275-.225-.5-.5-.5m0 11h-3v-5h3zm5-13h-3c-.275 0-.5.225-.5.5v13c0 .275.225.5.5.5h3c.275 0 .5-.225.5-.5v-13c0-.275-.225-.5-.5-.5m0 13h-3V9h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatsDots(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 14h14v2H0V0h2zm2.5-1a1.5 1.5 0 1 1 .131-2.994l1.612-2.687a1.5 1.5 0 1 1 2.514 0l1.612 2.687a1.42 1.42 0 0 1 .23-.002l2.662-4.658a1.5 1.5 0 1 1 1.14.651l-2.662 4.658a1.5 1.5 0 1 1-2.496.026L7.631 7.994a1.42 1.42 0 0 1-.262 0l-1.612 2.687A1.5 1.5 0 0 1 4.5 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Steam(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 11 4.5m3.975-2.475A3.5 3.5 0 0 0 9.036 4L5.838 8.847c-.43.022-.856.132-1.249.328L2.122 7.247A1.313 1.313 0 0 0 .506 9.315l2.436 1.905a3.157 3.157 0 1 0 6.168.228l3.891-3.484a3.5 3.5 0 0 0 1.975-5.939zM6 14.105A2.105 2.105 0 0 1 3.895 12l.001-.033l1.046.817a1.31 1.31 0 1 0 1.616-2.068l-.992-.776A2.105 2.105 0 1 1 6 14.105M12.5 7a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SteamTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.749 13.063a1.313 1.313 0 0 0 .729-2.405L4.446 9.97a2.105 2.105 0 1 1-1.548 2.124l1.124.749c.224.149.477.221.727.221zM13.333 0C14.8 0 16 1.2 16 2.667v10.666C16 14.801 14.8 16 13.333 16H2.667A2.674 2.674 0 0 1 0 13.333v-3.172l1.896 1.264a3.157 3.157 0 1 0 6.213.024l3.892-3.484a3.5 3.5 0 1 0-3.964-3.964L4.839 8.848a3.14 3.14 0 0 0-1.433.428L.002 7.007V2.668c0-1.467 1.2-2.667 2.667-2.667h10.666zM14 4.5a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0m-4 0a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 10 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 14.5a6.5 6.5 0 1 1 0-13a6.5 6.5 0 0 1 0 13M5 5h6v6H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h12v12H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3.019V2h2V1a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v1h2v1.019a6.5 6.5 0 1 0 1 0m3.036 10.017c-.944.944-2.2 1.464-3.536 1.464s-2.591-.52-3.536-1.464C3.02 12.092 2.5 10.836 2.5 9.5s.52-2.591 1.464-3.536a4.967 4.967 0 0 1 3.377-1.462l-.339 4.907c-.029.411.195.591.497.591s.527-.18.497-.591l-.339-4.907c1.276.04 2.47.555 3.377 1.462c.944.944 1.464 2.2 1.464 3.536s-.52 2.591-1.464 3.536z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8v1h-3.664c.43.602.664 1.292.664 2c0 1.107-.573 2.172-1.572 2.921C10.501 14.617 9.283 15 8 15s-2.501-.383-3.428-1.079C3.573 13.172 3 12.107 3 11h2c0 1.084 1.374 2 3 2s3-.916 3-2s-1.374-2-3-2H0V8h4.68a3.003 3.003 0 0 1-.108-.079C3.573 7.172 3 6.107 3 5s.573-2.172 1.572-2.921C5.499 1.383 6.717 1 8 1s2.501.383 3.428 1.079C12.427 2.828 13 3.893 13 5h-2c0-1.084-1.374-2-3-2s-3 .916-3 2s1.374 2 3 2c1.234 0 2.407.354 3.32 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stumbleupon(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5c-.55 0-1 .45-1 1v4c0 1.653-1.347 3-3 3s-3-1.347-3-3V8h2v2c0 .55.45 1 1 1s1-.45 1-1V6c0-1.653 1.347-3 3-3s3 1.347 3 2.781v.969l-1.281.375L9 6.75v-.969C9 5.45 8.55 5 8 5"/><path fill="currentColor" d="M15 10c0 1.653-1.347 3-3 3s-3-1.347-3-3.219V7.843l.719.375L11 7.843v1.938c0 .769.45 1.219 1 1.219s1-.45 1-1V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StumbleuponTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.313 0H2.688A2.696 2.696 0 0 0 0 2.688v10.625a2.696 2.696 0 0 0 2.688 2.688h10.625a2.696 2.696 0 0 0 2.688-2.688V2.688A2.696 2.696 0 0 0 13.313 0M8 5c-.551 0-1 .449-1 1v4c0 1.654-1.346 3-3 3s-3-1.346-3-3V8h2v2c0 .551.449 1 1 1s1-.449 1-1V6c0-1.654 1.346-3 3-3s3 1.346 3 2.781v.969l-1.281.375L9 6.75v-.969C9 5.448 8.551 5 8 5m7 5c0 1.654-1.346 3-3 3s-3-1.346-3-3.219V7.843l.719.375L11 7.843v1.938c0 .77.449 1.219 1 1.219s1-.449 1-1V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subscript(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14.219V15h2v1h-3v-2.281l2-.938V12h-2v-1h3v2.281zM10.563 4H8.438L5.5 6.938L2.562 4H.437l4 4l-4 4h2.125L5.5 9.062L8.438 12h2.125l-4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubscriptTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.032 13l.9-3h4.137l.9 3h1.775l-3-10H4.256l-3 10zm2.4-8h1.137l.9 3H4.532zM16 3l-2.5 4L11 3zm-2.5 10h-1a.5.5 0 0 1 0-1h2a.5.5 0 0 0 0-1h-2a1.502 1.502 0 0 0-1.117 2.5c.275.307.674.5 1.117.5h1a.5.5 0 0 1 0 1h-2a.5.5 0 0 0 0 1h2a1.502 1.502 0 0 0 1.117-2.5A1.496 1.496 0 0 0 13.5 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 13a1 1 0 0 1 1 1v1a1 1 0 0 1-2 0v-1a1 1 0 0 1 1-1M8 3a1 1 0 0 1-1-1V1a1 1 0 0 1 2 0v1a1 1 0 0 1-1 1m7 4a1 1 0 0 1 0 2h-1a1 1 0 0 1 0-2zM3 8a1 1 0 0 1-1 1H1a1 1 0 0 1 0-2h1a1 1 0 0 1 1 1m9.95 3.536l.707.707a1 1 0 0 1-1.414 1.414l-.707-.707a1 1 0 0 1 1.414-1.414m-9.9-7.072l-.707-.707a.999.999 0 1 1 1.414-1.414l.707.707A.999.999 0 1 1 3.05 4.464m9.9 0a.999.999 0 1 1-1.414-1.414l.707-.707a.999.999 0 1 1 1.414 1.414zm-9.9 7.072a1 1 0 0 1 1.414 1.414l-.707.707a1 1 0 0 1-1.414-1.414zM8 4a4 4 0 1 0 0 8a4 4 0 0 0 0-8m0 6.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Superscript(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.219V4h2v1h-3V2.719l2-.938V1h-2V0h3v2.281zM10.563 4H8.438L5.5 6.938L2.562 4H.437l4 4l-4 4h2.125L5.5 9.062L8.438 12h2.125l-4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SuperscriptTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.032 13l.9-3h4.137l.9 3h1.775l-3-10H4.256l-3 10zm2.4-8h1.137l.9 3H4.532zM11 13l2.5-4l2.5 4zm2.5-11h-1a.5.5 0 0 1 0-1h2a.5.5 0 0 0 0-1h-2a1.502 1.502 0 0 0-1.117 2.5c.275.307.674.5 1.117.5h1a.5.5 0 0 1 0 1h-2a.5.5 0 0 0 0 1h2a1.502 1.502 0 0 0 1.117-2.5A1.496 1.496 0 0 0 13.5 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Svg(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 6.5c-.444 0-.843.193-1.118.5h-2.968l2.099-2.099a1.5 1.5 0 1 0-1.414-1.414L9 5.586V2.618a1.5 1.5 0 1 0-2 0v2.968L4.901 3.487a1.5 1.5 0 1 0-1.414 1.414L5.586 7H2.618a1.5 1.5 0 1 0 0 2h2.968l-2.099 2.099a1.5 1.5 0 1 0 1.414 1.414L7 10.414v2.968a1.5 1.5 0 1 0 2 0v-2.968l2.099 2.099a1.5 1.5 0 1 0 1.414-1.414L10.414 9h2.968A1.5 1.5 0 1 0 14.5 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Switch(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2.29v2.124c.566.247 1.086.6 1.536 1.05C12.48 6.408 13 7.664 13 9s-.52 2.591-1.464 3.536C10.592 13.48 9.336 14 8 14s-2.591-.52-3.536-1.464C3.52 11.592 3 10.336 3 9s.52-2.591 1.464-3.536c.45-.45.97-.803 1.536-1.05V2.29a7 7 0 1 0 4 0M7 0h2v8H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tab(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 0h1v8h-1zM0 8h1v8H0zm5 3h11v2H5v2.5L1.5 12L5 8.5zm6-6H0V3h11V.5L14.5 4L11 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3v11h16V3zm6 7V8h4v2zm4 1v2H6v-2zm0-6v2H6V5zM5 5v2H1V5zM1 8h4v2H1zm10 0h4v2h-4zm0-1V5h4v2zM1 11h4v2H1zm10 2v-2h4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1v14h16V1zm6 9V7h4v3zm4 1v3H6v-3zm0-8v3H6V3zM5 3v3H1V3zM1 7h4v3H1zm10 0h4v3h-4zm0-1V3h4v3zM1 11h4v3H1zm10 3v-3h4v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 0h-10C1.675 0 1 .675 1 1.5v13c0 .825.675 1.5 1.5 1.5h10c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5m-5 15.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M12 14H3V2h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Target(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7h-1.577A6.516 6.516 0 0 0 9 1.577V0H7v1.577A6.516 6.516 0 0 0 1.577 7H0v2h1.577A6.516 6.516 0 0 0 7 14.423V16h2v-1.577A6.516 6.516 0 0 0 14.423 9H16zm-3.612 0h-1.559A3.008 3.008 0 0 0 9 5.171V3.612A4.516 4.516 0 0 1 12.388 7M8 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2M7 3.612v1.559A3.008 3.008 0 0 0 5.171 7H3.612A4.516 4.516 0 0 1 7 3.612M3.612 9h1.559A3.008 3.008 0 0 0 7 10.829v1.559A4.516 4.516 0 0 1 3.612 9M9 12.388v-1.559A3.008 3.008 0 0 0 10.829 9h1.559A4.516 4.516 0 0 1 9 12.388"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telegram(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.581 0 0 3.581 0 8s3.581 8 8 8s8-3.581 8-8s-3.581-8-8-8m3.931 5.484l-1.313 6.184c-.091.441-.356.544-.725.341l-2-1.478l-.959.934c-.112.109-.2.2-.4.2c-.259 0-.216-.097-.303-.344L5.55 9.084l-1.978-.616c-.428-.131-.431-.425.097-.634l7.706-2.975c.35-.159.691.084.556.625"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1v14h16V1zm15 13H1V2h14zM14 3H2v10h12zM7 8H6v1H5v1H4V9h1V8h1V7H5V6H4V5h1v1h1v1h1zm4 2H8V9h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextColor(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.032 13l.9-3h4.137l.9 3h1.775l-3-10H6.256l-3 10zm2.4-8h1.137l.9 3H6.532z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextHeight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 12h2l-2.5 3l-2.5-3h2V4h-2l2.5-3L16 4h-2zM10 1v4L9 3H6v11h2v1H2v-1h2V3H1L0 5V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextWidth(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 14v2l-3-2.5L4 11v2h8v-2l3 2.5l-3 2.5v-2zm9-13v4l-1-2H9v7h2v1H5v-1h2V3H4L3 5V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 5l2 2l-4 4l-2-2zm6.649-.351L14.5 3.5L14 4a1.414 1.414 0 1 1-1.999-2l.5-.5L11.352.351a1.208 1.208 0 0 0-1.703 0L.352 9.648a1.208 1.208 0 0 0 0 1.703L1.501 12.5L2 12.001A1.414 1.414 0 1 1 4 14l-.5.5l1.149 1.149a1.208 1.208 0 0 0 1.703 0l9.297-9.297a1.208 1.208 0 0 0 0-1.703M7 13L3 9l6-6l4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tongue(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13M4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m2 4v1h-1v1.5a1.5 1.5 0 0 1-3 0V10H4V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TongueTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2m7 6h-1v1.5a1.5 1.5 0 0 1-3 0V10H4V9h8zm-1-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 12H15V8.75C15 7.785 14.215 7 13.25 7H9V5h.25c.412 0 .75-.338.75-.75v-2.5A.752.752 0 0 0 9.25 1h-2.5a.752.752 0 0 0-.75.75v2.5c0 .412.338.75.75.75H7v2H2.75C1.785 7 1 7.785 1 8.75V12H.75a.753.753 0 0 0-.75.75v2.5c0 .412.338.75.75.75h2.5c.413 0 .75-.338.75-.75v-2.5a.752.752 0 0 0-.75-.75H3V9h4v3h-.25a.753.753 0 0 0-.75.75v2.5c0 .412.338.75.75.75h2.5c.412 0 .75-.338.75-.75v-2.5a.753.753 0 0 0-.75-.75H9V9h4v3h-.25a.753.753 0 0 0-.75.75v2.5c0 .412.338.75.75.75h2.5c.412 0 .75-.338.75-.75v-2.5a.753.753 0 0 0-.75-.75M3 15H1v-2h2zm6 0H7v-2h2zM7 4V2h2v2zm8 11h-2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trello(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M7 12c0 .55-.45 1-1 1H4c-.55 0-1-.45-1-1V4c0-.55.45-1 1-1h2c.55 0 1 .45 1 1zm6-3c0 .55-.45 1-1 1h-2c-.55 0-1-.45-1-1V4c0-.55.45-1 1-1h2c.55 0 1 .45 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3V1H3v2H0v2a3 3 0 0 0 3.9 2.862A5 5 0 0 0 7 9.899v3.1H6a2 2 0 0 0-2 2h8a2 2 0 0 0-2-2H9v-3.1a5.003 5.003 0 0 0 3.1-2.037A3 3 0 0 0 16 5V3zM3 6.813A1.815 1.815 0 0 1 1.187 5V4H3v1c0 .628.116 1.229.327 1.782c-.106.019-.216.03-.327.03zM14.813 5a1.815 1.815 0 0 1-2.14 1.783A4.994 4.994 0 0 0 13 5.001v-1h1.813v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 9l-2-4h-3V3c0-.55-.45-1-1-1H1c-.55 0-1 .45-1 1v8l1 1h1.268a2 2 0 1 0 3.464 0h5.536a2 2 0 1 0 3.464 0H16zm-5 0V6h2.073l1.5 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.001 7v3.659c0 .928-.012 1.463.086 1.727c.098.262.342.534.609.691c.354.212.758.318 1.214.318c.81 0 1.289-.107 2.09-.633v2.405a9.089 9.089 0 0 1-1.833.639A7.93 7.93 0 0 1 9.369 16a4.9 4.9 0 0 1-1.725-.276a4.195 4.195 0 0 1-1.438-.79c-.398-.343-.672-.706-.826-1.091s-.23-.944-.23-1.676V6.556H3.003V4.29c.628-.204 1.331-.497 1.778-.877a4.386 4.386 0 0 0 1.08-1.374C6.133 1.505 6.32.825 6.422 0h2.579v4H13v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumblrTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5m-3.066 12.884c-.472.222-.9.378-1.281.469a5.522 5.522 0 0 1-1.241.134c-.506 0-.803-.063-1.191-.191s-.719-.309-.994-.544c-.275-.238-.463-.488-.569-.753s-.159-.65-.159-1.156V6.971h-1.5V5.408c.434-.141.938-.344 1.244-.606c.309-.263.559-.578.744-.947c.188-.369.316-.837.388-1.406h1.569v2.55H11v1.972H8.447v2.831c0 .641-.009 1.009.059 1.191s.238.369.422.475c.244.147.525.219.838.219c.559 0 1.116-.181 1.669-.544z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tux(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.87 11.514c-1.28.596-2.471.589-3.271.532c-.954-.069-1.721-.33-2.058-.558a.454.454 0 1 0-.509.753c.542.367 1.501.64 2.503.712c.17.012.356.02.557.02c.872 0 1.979-.083 3.163-.634a.454.454 0 1 0-.384-.824zm5.051-.673C13.965 6.893 14.37-.568 6.812.035C-.651.636 1.328 8.519 1.217 11.159C1.118 12.556.655 14.263 0 16h2.017a12.3 12.3 0 0 0 .425-2.159c.122.085.252.167.391.245c.226.133.42.31.626.497c.48.438 1.025.934 2.089.996c.071.004.143.006.214.006c1.077 0 1.813-.471 2.404-.85c.283-.181.528-.338.759-.413c.655-.205 1.227-.536 1.655-.957c.067-.066.129-.133.187-.202c.238.873.564 1.856.926 2.836H16c-1.034-1.597-2.101-3.162-2.079-5.159zM1.939 8.693c-.074-1.288.542-2.372 1.377-2.421s1.571.957 1.645 2.245v.001c.004.069.006.138.006.206a3.239 3.239 0 0 0-.717.275l-.002-.029c-.071-.731-.462-1.284-.873-1.234s-.686.684-.614 1.415a2 2 0 0 0 .251.819a91.083 91.083 0 0 0-.511.376c-.311-.408-.524-.993-.562-1.655zm8.456 3.185c-.03.681-.92 1.322-1.743 1.579l-.005.002c-.342.111-.647.306-.97.513c-.543.347-1.104.706-1.914.706c-.053 0-.108-.002-.161-.005c-.742-.043-1.09-.36-1.529-.761c-.232-.211-.472-.43-.781-.611l-.007-.004c-.667-.377-1.081-.845-1.108-1.253c-.013-.203.077-.378.268-.522c.416-.312.695-.516.879-.651c.205-.15.267-.195.313-.239l.106-.103c.382-.371 1.021-.993 2.002-.993c.6 0 1.264.231 1.971.686c.333.217.623.317.99.444c.252.087.539.186.922.35l.006.003c.357.147.78.415.76.858zm-.197-1.6a2.645 2.645 0 0 0-.215-.098a12.404 12.404 0 0 0-.852-.328c.127-.248.206-.558.213-.894c.018-.818-.395-1.483-.922-1.484s-.968.661-.986 1.479a1.067 1.067 0 0 0 0 .08a4.534 4.534 0 0 0-.956-.324l-.004-.092v-.001c-.03-1.491.884-2.725 2.043-2.756s2.122 1.152 2.153 2.642v.001a3.184 3.184 0 0 1-.475 1.776z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.331 4.502a49.408 49.408 0 0 0-4.407-.425L13.5 1.501l-1-1L8.991 4.01A54.24 54.24 0 0 0 8 4.001l-4-4l-1 1l3.034 3.034a50.304 50.304 0 0 0-5.365.467C.239 6.185 0 8.045 0 10s.239 3.815.669 5.498C2.913 15.821 5.393 16 8 16s5.087-.179 7.331-.502c.43-1.683.669-3.543.669-5.498s-.239-3.815-.669-5.498m-1.833 9.164C11.815 13.881 9.955 14 8 14s-3.815-.119-5.498-.334C2.179 12.544 2 11.304 2 10s.179-2.543.502-3.666C4.185 6.119 6.045 6 8 6s3.815.119 5.498.334C13.821 7.456 14 8.696 14 10s-.179 2.543-.502 3.666"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitch(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 0L0 2.5V14h4v2h2l2-2h2.5L15 9.5V0zM13 8.5L10.5 11H8l-2 2v-2H3V2h10z"/><path fill="currentColor" d="M9.5 4H11v4H9.5zm-3 0H8v4H6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3.538a6.461 6.461 0 0 1-1.884.516a3.301 3.301 0 0 0 1.444-1.816a6.607 6.607 0 0 1-2.084.797a3.28 3.28 0 0 0-2.397-1.034a3.28 3.28 0 0 0-3.197 4.028a9.321 9.321 0 0 1-6.766-3.431a3.284 3.284 0 0 0 1.015 4.381A3.301 3.301 0 0 1 .643 6.57v.041A3.283 3.283 0 0 0 3.277 9.83a3.291 3.291 0 0 1-1.485.057a3.293 3.293 0 0 0 3.066 2.281a6.586 6.586 0 0 1-4.862 1.359a9.286 9.286 0 0 0 5.034 1.475c6.037 0 9.341-5.003 9.341-9.341c0-.144-.003-.284-.009-.425a6.59 6.59 0 0 0 1.637-1.697z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1h2v6.5c0 2.485-2.239 4.5-5 4.5S3 9.985 3 7.5V1h2v6.5c0 .628.285 1.23.802 1.695C6.379 9.714 7.159 10 8 10s1.621-.286 2.198-.805C10.715 8.729 11 8.127 11 7.5zM3 13h10v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1a7.979 7.979 0 0 0-5.657 2.343L0 1v6h6L3.757 4.757a6 6 0 1 1 8.211 8.743l1.323 1.5A8 8 0 0 0 8 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UndoTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.904 16C13.681 12.781 13.98 7.87 7 8.034V12L1 6l6-6v3.881c8.359-.218 9.29 7.378 4.904 12.119"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ungroup(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7.25c0 .412-.338.75-.75.75h-1.5A.752.752 0 0 1 3 7.25v-1.5c0-.412.337-.75.75-.75h1.5c.412 0 .75.338.75.75zm5 0c0 .412-.338.75-.75.75h-1.5A.753.753 0 0 1 8 7.25v-1.5c0-.412.338-.75.75-.75h1.5c.412 0 .75.338.75.75zm-5 5c0 .412-.338.75-.75.75h-1.5a.752.752 0 0 1-.75-.75v-1.5c0-.412.337-.75.75-.75h1.5c.412 0 .75.338.75.75zm5 0c0 .412-.338.75-.75.75h-1.5a.753.753 0 0 1-.75-.75v-1.5c0-.412.338-.75.75-.75h1.5c.412 0 .75.338.75.75zm3.251-9.75L16 .751V0h-.751L13.5 1.749L11.751 0H11v.751L12.749 2.5L11 4.249V5h.751L13.5 3.251L15.249 5H16v-.751zM0 12h1v2H0zm0-3h1v2H0zm13-2h1v2h-1zm0 6h1v2h-1zm0-3h1v2h-1zM0 6h1v2H0zm0-3h1v2H0zm8-1h2v1H8zM5 2h2v1H5zM2 2h2v1H2zm5 13h2v1H7zm3 0h2v1h-2zm-6 0h2v1H4zm-3 0h2v1H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlocked(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1c1.654 0 3 1.346 3 3v3h-2V4c0-.551-.449-1-1-1h-2c-.551 0-1 .449-1 1v3h.25c.412 0 .75.338.75.75v7.5c0 .412-.338.75-.75.75H.75a.753.753 0 0 1-.75-.75v-7.5C0 7.338.338 7 .75 7H7V4c0-1.654 1.346-3 3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 9h2V5h3L8 1L4 5h3zm3-2.25v1.542L14.579 10L8 12.453L1.421 10L6 8.292V6.75L0 9v4l8 3l8-3V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadThree(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 11H0v4h15v-4zm6.5 2h-2v-1h2zM3.5 5l4-4l4 4H9v5H6V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 14h16v1H0zm16-2v1H0v-1l2-4h4v2h4V8h4zM3.5 5L8 .5L12.5 5H9v4H7V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11.041v-.825c1.102-.621 2-2.168 2-3.716C11 4.015 11 2 8 2S5 4.015 5 6.5c0 1.548.898 3.095 2 3.716v.825c-3.392.277-6 1.944-6 3.959h14c0-2.015-2.608-3.682-6-3.959"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCheck(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 9.5L10.5 14L9 12.5l-1 1l2.5 2.5l5.5-5.5z"/><path fill="currentColor" d="M7 12h5v-1.799c-1.05-.613-2.442-1.033-4-1.16v-.825c1.102-.621 2-2.168 2-3.716C10 2.015 10 0 7 0S4 2.015 4 4.5c0 1.548.898 3.095 2 3.716v.825C2.608 9.318 0 10.985 0 13h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinus(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 11.5a5.508 5.508 0 0 1 3.594-5.159A4.574 4.574 0 0 0 10 4.5C10 2.015 10 0 7 0S4 2.015 4 4.5c0 1.548.898 3.095 2 3.716v.825C2.608 9.318 0 10.985 0 13h6.208A5.5 5.5 0 0 1 6 11.5"/><path fill="currentColor" d="M11.5 7a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9m2.5 5H9v-1h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 11.5a5.508 5.508 0 0 1 3.594-5.159A4.574 4.574 0 0 0 10 4.5C10 2.015 10 0 7 0S4 2.015 4 4.5c0 1.548.898 3.095 2 3.716v.825C2.608 9.318 0 10.985 0 13h6.208A5.5 5.5 0 0 1 6 11.5"/><path fill="currentColor" d="M11.5 7a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9m2.5 5h-2v2h-1v-2H9v-1h2V9h1v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTie(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3a3 3 0 1 1 6 0a3 3 0 0 1-6 0m7.001 4h-.553l-3.111 6.316L9.5 7.5L8 6L6.5 7.5l1.163 5.816L4.552 7h-.554c-1.999 0-1.999 1.344-1.999 3v5h12v-5c0-1.656 0-3-1.999-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12.041v-.825c1.102-.621 2-2.168 2-3.716C14 5.015 14 3 11 3S8 5.015 8 7.5c0 1.548.898 3.095 2 3.716v.825c-3.392.277-6 1.944-6 3.959h14c0-2.015-2.608-3.682-6-3.959"/><path fill="currentColor" d="M5.112 12.427c.864-.565 1.939-.994 3.122-1.256a5.667 5.667 0 0 1-.633-.922a5.726 5.726 0 0 1-.726-2.748c0-1.344 0-2.614.478-3.653c.464-1.008 1.299-1.633 2.488-1.867C9.577.786 8.873.001 7 .001c-3 0-3 2.015-3 4.5c0 1.548.898 3.095 2 3.716v.825c-3.392.277-6 1.944-6 3.959h4.359c.227-.202.478-.393.753-.573z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCamera(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0m-6 0a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0m12 5V8c0-.55-.45-1-1-1H1c-.55 0-1 .45-1 1v5c0 .55.45 1 1 1h10c.55 0 1-.45 1-1v-1.5l4 2.5V7zM10 12H2V9h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.994 4.281c-.072 1.556-1.159 3.691-3.263 6.397c-2.175 2.825-4.016 4.241-5.522 4.241c-.931 0-1.722-.859-2.366-2.581c-.431-1.578-.859-3.156-1.291-4.734c-.478-1.722-.991-2.581-1.541-2.581c-.119 0-.538.253-1.256.753l-.753-.969c.791-.694 1.569-1.388 2.334-2.081c1.053-.909 1.844-1.387 2.372-1.438c1.244-.119 2.013.731 2.3 2.553c.309 1.966.525 3.188.647 3.666c.359 1.631.753 2.447 1.184 2.447c.334 0 .838-.528 1.509-1.588c.669-1.056 1.028-1.862 1.078-2.416c.097-.912-.262-1.372-1.078-1.372a2.98 2.98 0 0 0-1.184.263c.787-2.575 2.287-3.825 4.506-3.753c1.641.044 2.416 1.109 2.322 3.194z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VimeoTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5m-1.037 5.313c-.05 1.125-.838 2.666-2.359 4.622c-1.572 2.044-2.903 3.066-3.991 3.066c-.675 0-1.244-.622-1.709-1.866c-.313-1.141-.622-2.281-.934-3.422c-.344-1.244-.716-1.866-1.112-1.866c-.087 0-.391.181-.906.544l-.544-.7c.572-.5 1.134-1.003 1.687-1.503c.763-.656 1.331-1.003 1.712-1.038c.9-.087 1.453.528 1.662 1.844c.225 1.422.381 2.303.469 2.65c.259 1.178.544 1.766.856 1.766c.241 0 .606-.381 1.091-1.147s.744-1.347.778-1.747c.069-.659-.191-.991-.778-.991c-.278 0-.563.063-.856.191c.569-1.859 1.653-2.766 3.256-2.712c1.188.034 1.747.803 1.678 2.309"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vine(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.012 7.953a5.297 5.297 0 0 1-1.169.137c-2.019 0-3.572-1.409-3.572-3.862c0-1.203.466-1.825 1.122-1.825c.625 0 1.041.559 1.041 1.697c0 .647-.172 1.356-.3 1.775c0 0 .622 1.084 2.322.753c.363-.803.556-1.841.556-2.75c0-2.45-1.25-3.878-3.541-3.878c-2.356 0-3.734 1.809-3.734 4.197c0 2.366 1.106 4.394 2.928 5.319c-.766 1.534-1.741 2.884-2.759 3.903c-1.844-2.231-3.513-5.206-4.197-11.016H.987c1.259 9.675 5.006 12.756 6 13.347c.559.337 1.044.322 1.556.031c.806-.456 3.222-2.875 4.563-5.703c.563 0 1.238-.066 1.909-.219V7.953z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vk(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5m-1.541 11.2l-1.463.022s-.316.063-.728-.222c-.547-.375-1.063-1.353-1.466-1.225c-.406.128-.394 1.006-.394 1.006s.003.188-.091.287c-.1.109-.3.131-.3.131h-.653s-1.444.088-2.716-1.238C3.76 8.517 2.536 5.652 2.536 5.652s-.072-.188.006-.278c.087-.103.322-.109.322-.109l1.566-.009s.147.025.253.103c.088.063.134.184.134.184s.253.641.588 1.219c.653 1.128.959 1.375 1.181 1.256c.322-.175.225-1.597.225-1.597s.006-.516-.162-.744c-.131-.178-.378-.231-.484-.244c-.088-.013.056-.216.244-.309c.281-.138.778-.147 1.366-.141c.456.003.591.034.769.075c.541.131.356.634.356 1.841c0 .388-.069.931.209 1.109c.119.078.412.012 1.147-1.234c.347-.591.609-1.284.609-1.284s.056-.125.144-.178c.091-.053.213-.037.213-.037l1.647-.009s.494-.059.575.166c.084.234-.184.781-.856 1.678c-1.103 1.472-1.228 1.334-.309 2.184c.875.813 1.056 1.209 1.088 1.259c.356.6-.406.647-.406.647z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDecrease(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 7h8v2H8zm-1.5 8a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeHigh(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.907 14.407a.75.75 0 0 1-.53-1.281c1.369-1.369 2.123-3.19 2.123-5.127s-.754-3.757-2.123-5.127a.75.75 0 1 1 1.061-1.061c1.653 1.653 2.563 3.85 2.563 6.187s-.91 4.534-2.563 6.187a.748.748 0 0 1-.53.22zm-2.664-1.414a.75.75 0 0 1-.53-1.281a5.256 5.256 0 0 0 0-7.425a.75.75 0 1 1 1.061-1.061c1.275 1.275 1.977 2.97 1.977 4.773s-.702 3.498-1.977 4.773a.748.748 0 0 1-.53.22zm-2.665-1.415a.75.75 0 0 1-.53-1.281a3.254 3.254 0 0 0 0-4.596a.75.75 0 1 1 1.061-1.061a4.756 4.756 0 0 1 0 6.718a.748.748 0 0 1-.53.22zM6.5 15a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeIncrease(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 9h-3v3h-2V9H8V7h3V4h2v3h3zm-9.5 6a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeLow(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.578 11.578a.75.75 0 0 1-.53-1.281a3.254 3.254 0 0 0 0-4.596a.75.75 0 1 1 1.061-1.061a4.756 4.756 0 0 1 0 6.718a.748.748 0 0 1-.53.22zM6.5 15a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMedium(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.243 12.993a.75.75 0 0 1-.53-1.281a5.256 5.256 0 0 0 0-7.425a.75.75 0 1 1 1.061-1.061c1.275 1.275 1.977 2.97 1.977 4.773s-.702 3.498-1.977 4.773a.748.748 0 0 1-.53.22zm-2.665-1.415a.75.75 0 0 1-.53-1.281a3.254 3.254 0 0 0 0-4.596a.75.75 0 1 1 1.061-1.061a4.756 4.756 0 0 1 0 6.718a.748.748 0 0 1-.53.22zM6.5 15a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMute(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 15a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMuteTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 9.674V11h-1.326L12 9.326L10.326 11H9V9.674L10.674 8L9 6.326V5h1.326L12 6.674L13.674 5H15v1.326L13.326 8zM6.5 15a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 1.45l6.705 13.363H1.296zM8 0c-.345 0-.69.233-.951.698L.22 14.309C-.303 15.239.142 16 1.209 16h13.583c1.067 0 1.512-.761.989-1.691L8.952.698C8.69.233 8.346 0 8.001 0z"/><path fill="currentColor" d="M9 13a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1-2a1 1 0 0 1-1-1V7a1 1 0 0 1 2 0v3a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.641 2.325A7.869 7.869 0 0 0 8.032 0C3.663 0 .107 3.556.107 7.928c0 1.397.366 2.763 1.059 3.963L.041 16l4.203-1.103a7.929 7.929 0 0 0 3.787.966h.003c4.369 0 7.928-3.556 7.928-7.928a7.887 7.887 0 0 0-2.322-5.609zm-5.607 12.2a6.583 6.583 0 0 1-3.356-.919l-.241-.144l-2.494.653l.666-2.431l-.156-.25a6.537 6.537 0 0 1-1.009-3.506a6.599 6.599 0 0 1 6.594-6.591c1.759 0 3.416.688 4.659 1.931a6.554 6.554 0 0 1 1.928 4.662c-.003 3.637-2.959 6.594-6.591 6.594zm3.613-4.937c-.197-.1-1.172-.578-1.353-.644s-.313-.1-.447.1c-.131.197-.512.644-.628.778c-.116.131-.231.15-.428.05s-.838-.309-1.594-.984c-.588-.525-.987-1.175-1.103-1.372s-.013-.306.088-.403c.091-.088.197-.231.297-.347s.131-.197.197-.331c.066-.131.034-.247-.016-.347s-.447-1.075-.609-1.472c-.159-.388-.325-.334-.447-.341c-.116-.006-.247-.006-.378-.006s-.347.05-.528.247c-.181.197-.694.678-.694 1.653s.709 1.916.809 2.05c.1.131 1.397 2.134 3.384 2.991c.472.203.841.325 1.128.419c.475.15.906.128 1.247.078c.381-.056 1.172-.478 1.338-.941s.166-.859.116-.941c-.047-.088-.178-.137-.378-.238z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wikipedia(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.106 3.65c0 .05-.016.097-.047.141c-.031.041-.066.063-.106.063c-.313.031-.569.131-.766.3c-.2.169-.403.497-.613.975l-3.225 7.272c-.022.069-.081.1-.178.1a.205.205 0 0 1-.178-.1L8.184 8.62l-2.081 3.781c-.044.069-.1.1-.178.1c-.094 0-.153-.034-.184-.1L2.575 5.132c-.197-.45-.406-.766-.625-.944s-.525-.291-.916-.331C1 3.857.968 3.838.94 3.804a.173.173 0 0 1-.044-.122c0-.119.034-.178.1-.178c.281 0 .578.013.888.038c.288.025.556.038.809.038c.256 0 .563-.013.913-.038c.366-.025.691-.038.975-.038c.069 0 .1.059.1.178s-.022.175-.063.175c-.281.022-.506.094-.669.216s-.244.281-.244.481c0 .1.034.228.1.378l2.616 5.912l1.487-2.806l-1.384-2.903c-.25-.519-.453-.853-.612-1.003s-.403-.241-.728-.275c-.031 0-.056-.019-.084-.053s-.041-.075-.041-.122c0-.119.028-.178.088-.178c.281 0 .541.013.778.038c.228.025.469.038.728.038c.253 0 .519-.013.803-.038c.291-.025.578-.038.859-.038c.069 0 .1.059.1.178s-.019.175-.063.175c-.566.038-.847.2-.847.481c0 .125.066.322.197.588l.916 1.859l.912-1.7c.125-.241.191-.444.191-.606c0-.388-.281-.594-.847-.619c-.05 0-.075-.059-.075-.175c0-.044.012-.081.037-.119s.05-.056.075-.056c.203 0 .45.013.747.038c.281.025.516.038.697.038c.131 0 .322-.013.575-.031c.319-.028.588-.044.803-.044c.05 0 .075.05.075.15c0 .134-.047.203-.137.203c-.328.034-.594.125-.794.272s-.45.481-.75 1.006L8.905 7.379l1.644 3.35l2.428-5.647c.084-.206.125-.397.125-.569c0-.412-.281-.631-.847-.659c-.05 0-.075-.059-.075-.175c0-.119.037-.178.113-.178c.206 0 .45.013.734.038c.262.025.481.038.656.038c.188 0 .4-.013.644-.038c.253-.025.481-.038.684-.038c.063 0 .094.05.094.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.441 7.999c-.745-.383-1.47-.577-2.154-.577c-.093 0-.187.003-.28.011a9.082 9.082 0 0 0-2.603.643L.001 12.942c.964-.357 1.817-.53 2.598-.53c1.263 0 2.18.472 2.937.958c.359-1.217 1.219-4.158 1.476-5.036a12.34 12.34 0 0 0-.571-.333zm1.814 1.236l-1.413 4.909c.419.24 1.83 1.001 2.91 1.001c.872 0 1.848-.223 2.982-.684l1.349-4.718c-.916.296-1.795.446-2.617.446c-1.499 0-2.549-.486-3.211-.952zm-3.68-3.473c1.205.012 2.096.472 2.835.945l1.449-4.958c-.305-.175-1.106-.611-1.685-.759A5.28 5.28 0 0 0 5.968.855c-.809.015-1.694.218-2.701.622L1.885 6.33c1.013-.382 1.885-.568 2.689-.568zM16 3.096c-.919.357-1.816.539-2.672.539c-1.433 0-2.489-.497-3.173-.974L8.718 7.633c.965.62 2.005.936 3.096.936c.89 0 1.812-.214 2.742-.636l-.003-.035l.058-.014l1.39-4.788z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowsEight(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.005 8L0 3.124l6-.815V8zM7 2.164L14.998 1v7H7zM15 9l-.002 7L7 14.875V9zm-9 5.747l-5.995-.822V8.999H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wink(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13m.48 9.61c2.191-.433 3.892-1.43 4.507-2.759C12.649 10.975 10.463 13 7.817 13c-1.863 0-3.498-1.004-4.42-2.515c1.1.86 3.04 1.028 5.083.625M10 5.5c0-.828.448-1.5 1-1.5s1 .672 1 1.5S11.552 7 11 7s-1-.672-1-1.5m-4.5.305c-.653 0-1.208.245-1.414.586C4.031 6.299 4 5.888 4 5.786c0-.485.672-.879 1.5-.879s1.5.394 1.5.879c0 .103-.03.514-.086.605c-.206-.341-.761-.586-1.414-.586"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WinkTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m3 4c.552 0 1 .672 1 1.5S11.552 7 11 7s-1-.672-1-1.5s.448-1.5 1-1.5m-5.5.876c.932 0 1.594.349 1.594.895c0 .116.06.672-.003.775c-.232-.384-.856-.659-1.591-.659s-1.359.275-1.591.659c-.062-.103-.003-.659-.003-.775c0-.546.662-.895 1.594-.895M7.818 13c-1.863 0-3.498-1.004-4.42-2.515c1.1.86 3.04 1.028 5.083.625c2.191-.433 3.892-1.43 4.507-2.759C12.65 10.975 10.464 13 7.818 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Woman(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 9 1.5M11.234 8L12 7.445L9.917 4.224a.5.5 0 0 0-.417-.225h-4a.497.497 0 0 0-.417.225L3 7.445L3.766 8l1.729-2.244l.601 1.402l-2.095 3.841h1.917l.333 5h1v-5h.5v5h1l.333-5h1.917L8.906 7.157l.601-1.402l1.729 2.244z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wondering(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13m3.652 7.9l.351 1.2l-6.828 2l-.351-1.2zM4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WonderingTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m3 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2M4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0m1.176 7.6l-.351-1.2l6.828-2l.351 1.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8c0 2.313 1.38 4.312 3.382 5.259L2.52 5.622A5.693 5.693 0 0 0 2 8m10.05-.295c0-.722-.266-1.222-.495-1.612c-.304-.482-.589-.889-.589-1.371c0-.537.418-1.037 1.008-1.037c.027 0 .052.003.078.005A6.064 6.064 0 0 0 8 2.156A6.036 6.036 0 0 0 2.987 4.79c.141.004.274.007.386.007c.627 0 1.599-.074 1.599-.074c.323-.018.361.444.038.482c0 0-.325.037-.687.055l2.185 6.33l1.313-3.835l-.935-2.495a12.304 12.304 0 0 1-.629-.055c-.323-.019-.285-.5.038-.482c0 0 .991.074 1.58.074c.627 0 1.599-.074 1.599-.074c.323-.018.362.444.038.482c0 0-.326.037-.687.055l2.168 6.282l.599-1.947c.259-.809.457-1.389.457-1.889zm-3.945.806l-1.8 5.095a6.148 6.148 0 0 0 3.687-.093a.52.52 0 0 1-.043-.081zm5.16-3.315c.026.186.04.386.04.601c0 .593-.114 1.259-.456 2.093l-1.833 5.16c1.784-1.013 2.983-2.895 2.983-5.051a5.697 5.697 0 0 0-.735-2.803zM8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m0 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.671 12.779L8.475 6.611A4.5 4.5 0 0 0 3.193.193l2.6 2.6a1.003 1.003 0 0 1 0 1.414L4.207 5.793a1.003 1.003 0 0 1-1.414 0l-2.6-2.6a4.5 4.5 0 0 0 6.418 5.282l6.168 7.196a.914.914 0 0 0 1.358.052l1.586-1.586a.914.914 0 0 0-.052-1.358"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xing(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 0h-13C.675 0 0 .675 0 1.5v13c0 .825.675 1.5 1.5 1.5h13c.825 0 1.5-.675 1.5-1.5v-13c0-.825-.675-1.5-1.5-1.5M4.884 10.406H3.156a.253.253 0 0 1-.225-.119c-.047-.075-.047-.169 0-.266l1.838-3.244c.003-.003.003-.006 0-.009L3.6 4.743c-.047-.097-.056-.191-.009-.266c.044-.072.131-.109.237-.109h1.731c.266 0 .397.172.481.325c0 0 1.181 2.063 1.191 2.075c-.069.125-1.869 3.303-1.869 3.303c-.094.162-.219.334-.478.334zm8.185-8.028L9.238 9.153a.01.01 0 0 0 0 .012l2.441 4.456c.047.097.05.194.003.269c-.044.072-.125.109-.231.109H9.723c-.266 0-.397-.175-.484-.328L6.78 9.159l3.85-6.828c.094-.166.206-.328.463-.328h1.753c.103 0 .188.041.231.109c.044.072.044.169-.006.266z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func XingTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.431 3.159c-.138 0-.256.05-.316.144c-.059.1-.05.225.013.353l1.559 2.7c.003.006.003.009 0 .013L1.237 10.7c-.063.128-.059.256 0 .353a.34.34 0 0 0 .3.156h2.306c.344 0 .513-.234.628-.447l2.491-4.406L5.374 3.59c-.116-.203-.287-.431-.644-.431h-2.3zM12.125 0c-.344 0-.494.216-.619.441c0 0-4.972 8.816-5.134 9.106l3.278 6.016c.116.203.291.441.644.441H12.6c.137 0 .247-.053.306-.147c.063-.1.059-.228-.006-.356L9.65 9.554c-.003-.006-.003-.009 0-.016L14.759.504c.063-.128.066-.256.006-.356c-.059-.094-.169-.147-.306-.147z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yahoo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.878 9.203C10.637 6.115 13.544 1.078 14.341 0c-.35.234-.887.353-1.381.466L12.213 0c-.6 1.119-2.813 4.734-4.222 7.05C6.563 4.684 4.872 1.953 3.769 0C2.894.188 2.532.197 1.66 0c1.731 2.606 4.503 7.572 5.447 9.203L6.979 16l1.013-.466v-.012L9.004 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YahooTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.019 1.087C5.191 1.087 2.519.715 0 0v16c2.522-.716 5.194-1.088 8.019-1.088c2.794 0 5.459.363 7.981 1.088V0c-2.522.725-5.184 1.087-7.981 1.087m4.431 1.366l-.097.153c-.091.144-.172.266-.284.438c-.15.225-.431.672-.769 1.247c-.094.159-.209.35-.328.556l-.688 1.162l-.256.447l-.678 1.181c-.228.403-.453.8-.678 1.194v.397c0 .55.012 1.15.031 1.684c.009.244.019.678.031 1.137c.012.547.025 1.113.041 1.4l.003.088v.009l-.094-.025l-.109-.028a3.093 3.093 0 0 0-.353-.056a2.691 2.691 0 0 0-.444 0a3.058 3.058 0 0 0-.462.084l-.094.025v-.009l.003-.088c.013-.284.028-.853.041-1.4c.009-.459.022-.894.031-1.137a42.53 42.53 0 0 0 .031-1.684v-.397L6.65 7.637c-.222-.391-.453-.791-.675-1.181c-.088-.15-.172-.3-.256-.447c-.2-.347-.459-.781-.688-1.162a31.856 31.856 0 0 1-.328-.556a31.033 31.033 0 0 0-.769-1.247c-.112-.172-.194-.294-.284-.438l-.097-.153l.175.05c.222.063.45.094.694.094s.478-.031.697-.094l.053-.016l.028.047c.431.778 1.591 2.684 2.284 3.825c.237.394.428.703.522.862v-.003v.003l.522-.862c.694-1.138 1.853-3.044 2.284-3.825l.028-.047l.053.016c.219.063.453.094.697.094s.472-.031.694-.094z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yelp(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.514 10.21c-.27.272-.042.768-.042.768l2.033 3.394s.334.447.623.447c.29 0 .577-.239.577-.239l1.607-2.297s.162-.29.166-.544c.006-.361-.538-.46-.538-.46l-3.805-1.222s-.373-.099-.621.152zM9.321 8.5c.195.33.732.234.732.234l3.796-1.109s.517-.21.591-.491c.072-.281-.085-.619-.085-.619l-1.814-2.137s-.157-.27-.483-.297c-.36-.031-.581.405-.581.405L9.332 7.861s-.19.336-.01.64zM7.527 7.184c.447-.11.518-.759.518-.759l-.03-5.404S7.948.354 7.648.174c-.47-.285-.609-.136-.744-.116L3.753 1.229s-.309.102-.469.36c-.23.365.233.899.233.899l3.276 4.465s.323.334.735.233zm-.778 2.187c.011-.417-.5-.667-.5-.667L2.862 6.993s-.502-.207-.746-.063c-.187.11-.352.31-.368.486l-.221 2.716s-.033.471.089.685c.173.304.741.092.741.092l3.955-.874c.154-.103.423-.113.438-.664zm.983 1.466c-.339-.174-.746.187-.746.187l-2.648 2.915s-.33.446-.246.72c.079.257.21.384.396.474l2.659.839s.322.067.567-.004c.347-.1.283-.643.283-.643l.06-3.947s-.014-.38-.324-.541z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.841 4.8s-.156-1.103-.637-1.587c-.609-.637-1.291-.641-1.603-.678c-2.237-.163-5.597-.163-5.597-.163h-.006s-3.359 0-5.597.163c-.313.038-.994.041-1.603.678C.317 3.697.164 4.8.164 4.8S.005 6.094.005 7.391v1.213c0 1.294.159 2.591.159 2.591s.156 1.103.634 1.588c.609.637 1.409.616 1.766.684c1.281.122 5.441.159 5.441.159s3.363-.006 5.6-.166c.313-.037.994-.041 1.603-.678c.481-.484.637-1.588.637-1.588s.159-1.294.159-2.591V7.39c-.003-1.294-.162-2.591-.162-2.591zm-9.494 5.275V5.578l4.322 2.256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeTwo(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.375 2.647l.006-.028l.016-.118l-.74-.004c-.668-.004-.873 0-.891.017c-.009.008-.24.885-.651 2.473c-.196.758-.361 1.363-.367 1.345s-.24-.883-.522-1.922a107.288 107.288 0 0 0-.524-1.901c-.01-.01-.906-.014-1.632-.008c-.105.001-.164-.205.938 3.299c.152.485.381 1.172.507 1.526c.146.408.25.724.321.987c.126.501.13.815.103 1.182c-.032.423-.036 3.413-.005 3.463c.024.038 1.425.056 1.558.02c.021-.006.035-.026.045-.139c.033-.097.036-.484.036-2.09V8.698l.09-.283c.059-.185.206-.672.328-1.082l.327-1.09c.529-1.724 1.033-3.419 1.047-3.516l.011-.079zm7.846 2.488v.107h-.017l-.009 2.953l-.009 2.863l-.229.233c-.257.261-.462.361-.648.314c-.203-.051-.197.028-.214-3.356l-.016-3.115h-1.474v.107h-.017v3.38c0 3.621 0 3.619.184 3.982c.146.29.36.431.725.479c.481.064 1-.154 1.481-.622l.209-.203v.351c0 .303.009.353.064.368c.09.025 1.206.027 1.326.002l.1-.021v-.104l.017-.003V5.114l-1.472.02zM9.483 6.661c-.14-.599-.401-1.002-.832-1.28c-.676-.437-1.449-.484-2.165-.13c-.522.258-.859.686-1.032 1.314a1.383 1.383 0 0 0-.047.231c-.044.222-.049.552-.061 2.093c-.018 2.374.01 2.656.307 3.195c.292.529.897.917 1.556.997c.198.024.6-.013.832-.078c.525-.146 1.029-.561 1.252-1.032a1.8 1.8 0 0 0 .189-.604c.065-.353.07-.925.07-2.381c0-1.857-.006-2.06-.068-2.326zM7.802 11.5a.688.688 0 0 1-.515.098c-.135-.029-.318-.241-.374-.434c-.07-.241-.075-3.594-.015-4.251c.1-.329.378-.501.682-.419c.237.064.358.212.427.523c.051.231.057.518.046 2.207c-.007 1.12-.011 1.668-.048 1.962c-.037.185-.099.235-.203.315zm28.142-3.154h.712l-.011-.645c-.011-.592-.02-.659-.099-.82c-.125-.253-.309-.366-.601-.366c-.351 0-.573.17-.678.518c-.045.148-.092 1.167-.058 1.255c.019.049.121.058.735.058m-4.76-1.467a.49.49 0 0 0-.477-.278a.914.914 0 0 0-.508.203l-.127.097v4.634l.127.097c.288.22.604.266.822.12a.482.482 0 0 0 .186-.263c.057-.164.062-.375.055-2.325c-.008-2.032-.012-2.152-.078-2.285"/><path fill="currentColor" d="M40.014 4.791c-.142-1.701-.255-2.253-.605-2.962C38.944.89 38.273.395 37.317.286c-.739-.084-3.521-.203-6.094-.26c-4.456-.099-11.782.092-12.718.331a2.252 2.252 0 0 0-1.094.634c-.591.588-.944 1.432-1.085 2.6c-.323 2.666-.33 5.886-.019 8.649c.134 1.188.41 1.96.928 2.596c.323.397.881.734 1.379.835c.35.071 2.1.169 4.65.26c.38.014 1.385.037 2.235.052c1.77.031 5.025.013 6.886-.039c1.252-.035 3.534-.128 3.961-.161c.12-.009.398-.027.618-.039c.739-.042 1.209-.196 1.65-.543c.571-.449 1.013-1.278 1.2-2.251c.177-.92.295-2.559.319-4.42c.02-1.555-.007-2.393-.119-3.741zM22.27 4.175l-.828.01l-.036 8.83l-.718.009c-.555.008-.724-.001-.737-.036c-.01-.025-.021-2.016-.026-4.424l-.009-4.379l-1.617-.02v-1.38l4.779.019l.02 1.36zm5.077 5.061v3.797h-1.308v-.4c0-.301-.011-.4-.047-.4c-.026 0-.144.099-.263.22c-.259.263-.565.474-.827.572c-.542.203-1.056.084-1.275-.293c-.201-.345-.204-.423-.204-4.005v-3.29h1.307l.01 3.098c.01 3.044.011 3.1.084 3.224c.097.164.244.209.478.144c.138-.038.232-.105.455-.327l.282-.28V5.437h1.308v3.797zm5.102 3.255c-.115.257-.372.508-.583.57c-.549.162-.99.03-1.499-.449c-.158-.149-.305-.269-.327-.269c-.027 0-.041.116-.041.345v.345h-1.308V2.785h1.308v1.672c0 .919.012 1.672.027 1.672s.153-.122.307-.27c.354-.341.649-.491 1.024-.519c.669-.051 1.068.294 1.25 1.08c.057.245.062.525.062 2.798c0 2.768 0 2.78-.221 3.273zm5.535-1.52a4.706 4.706 0 0 1-.077.727c-.182.674-.666 1.152-1.366 1.348c-.942.264-1.98-.168-2.394-.997c-.232-.465-.241-.558-.241-2.831c0-1.853.007-2.081.066-2.334c.168-.715.584-1.178 1.289-1.435c.204-.074.417-.113.63-.117c.761-.016 1.515.393 1.832 1.059c.213.449.24.642.261 1.908l.019 1.136l-2.789.019l-.01.763c-.015 1.077.058 1.408.349 1.603c.244.165.62.152.824-.027c.192-.168.246-.349.265-.877l.017-.463h1.347z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.504 13.616l-3.79-3.223c-.392-.353-.811-.514-1.149-.499a6 6 0 1 0-.672.672c-.016.338.146.757.499 1.149l3.223 3.79c.552.613 1.453.665 2.003.115s.498-1.452-.115-2.003zM6 10a4 4 0 1 1 0-8a4 4 0 0 1 0 8m1-7H5v2H3v2h2v2h2V7h2V5H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *IcomoonFreeIcon {
	return &IcomoonFreeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.504 13.616l-3.79-3.223c-.392-.353-.811-.514-1.149-.499a6 6 0 1 0-.672.672c-.016.338.146.757.499 1.149l3.223 3.79c.552.613 1.453.665 2.003.115s.498-1.452-.115-2.003zM6 10a4 4 0 1 1 0-8a4 4 0 0 1 0 8M3 5h6v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
