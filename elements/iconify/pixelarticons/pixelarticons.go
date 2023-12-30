package pixelarticons

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "1.7.0"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type PixelarticonsIcon struct {
	*SVGSVGElement
}

func AbTesting(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h6v2H5v2h4v2H5v2h4v2H3zm6 8h2V9H9zm0-4h2V5H9zm4 4h8v10h-2v-4h-4v4h-2zm2 4h4v-2h-4zm0-12h6v6h-2V5h-4zM3 15h2v4h4v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ac(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2h-2v4H9V4H7v2h2v2h2v3H8V9H6V7H4v2h2v2H2v2h4v2H4v2h2v-2h2v-2h3v3H9v2H7v2h2v-2h2v4h2v-4h2v2h2v-2h-2v-2h-2v-3h3v2h2v2h2v-2h-2v-2h4v-2h-4V9h2V7h-2v2h-2v2h-3V8h2V6h2V4h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm16 16V5H5v14zm-6-8h4v2h-4v4h-2v-4H7v-2h4V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddBoxMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h14v14H3zm12 12V5H5v10zm-8 6v-2h12V7h2v14zm4-12h2v2h-2v2H9v-2H7V9h2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddCol(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h10v20H2v-2h8v-4H2v-2h8v-4H2V8h8V4H2zm17 9h3v2h-3v3h-2v-3h-3v-2h3V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddGrid(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h8v8H3zm6 6V5H5v4zm9 4h-2v3h-3v2h3v3h2v-3h3v-2h-3zM15 3h6v8h-8V3zm4 6V5h-4v4zM5 13h6v8H3v-8zm4 6v-4H5v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddRow(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 10V2H2v10h20V2h-2v8h-4V2h-2v8h-4V2H8v8zm9 9v3h-2v-3H8v-2h3v-3h2v3h3v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alert(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 1h-2v2H9v2H7v2H5v2H3v2H1v2h2v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v-2h2v-2h2v-2h-2V9h-2V7h-2V5h-2V3h-2zm0 2v2h2v2h2v2h2v2h2v2h-2v2h-2v2h-2v2h-2v2h-2v-2H9v-2H7v-2H5v-2H3v-2h2V9h2V7h2V5h2V3zm0 4h-2v6h2zm0 8h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5H4v2h16zm-4 4H8v2h8zM4 13h16v2H4zm12 4H8v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5H4v2h16zm0 4H4v2h16zM4 13h16v2H4zm16 4H4v2h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5H4v2h16zm-8 4H4v2h8zm8 4v2H4v-2zm-8 4H4v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5h16v2H4zm8 4h8v2h-8zm-8 4v2h16v-2zm8 4h8v2h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Analytics(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm16 2H5v14h14zM7 12h2v5H7zm10-5h-2v10h2zm-6 3h2v2h-2zm2 4h-2v3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 3h-4v2H8v4h2v2h1v8H6v-4h2v-2H4v6h2v2h12v-2h2v-6h-4v2h2v4h-5v-8h1V9h2V5h-2zm0 2v4h-4V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h2v2H2zm4 4H4V7h2zm2 0H6v2H4v2H2v6h20v-6h-2v-2h-2V9h2V7h2V5h-2v2h-2v2h-2V7H8zm0 0h8v2h2v2h2v4H4v-4h2v-2h2zm2 4H8v2h2zm4 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Animation(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2H2v12h2V4h10V2zm2 4h12v2H8v10H6zm4 4h12v12H10zm10 10v-8h-8v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 4H2v6h2v10h16V10h2zM6 10h12v8H6zm14-4v2H4V6zm-5 6H9v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBarDown(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4h2v8h2v2h-2v2h-2v-2H9v-2h2zm-2 8H7v-2h2zm6 0v-2h2v2zM4 18h16v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBarLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4v16H4V4zm14 7v2h-8v2h-2v-2H8v-2h2V9h2v2zm-8-2V7h2v2zm0 6h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBarRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4v16h2V4zM4 11v2h8v2h-2v2h2v-2h2v-2h2v-2h-2V9h-2V7h-2v2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBarUp(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h16V4H4zm7 14h2v-8h2v2h2v-2h-2v-2h-2V8h-2v2H9v2H7v2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4h2v12h2v2h-2v2h-2v-2H9v-2h2zM7 14v2h2v-2zm0 0v-2H5v2zm10 0v2h-2v-2zm0 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm16 16V5H5v14zM11 7h2v6h2v2h-2v2h-2v-2H9v-2h2zm-2 4v2H7v-2zm8 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 11v2H8v2H6v-2H4v-2h2V9h2v2zM10 7H8v2h2zm0 0h2V5h-2zm0 10H8v-2h2zm0 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 3v18H3V3zM5 19h14V5H5zm12-8v2h-6v2H9v-2H7v-2h2V9h2v2zm-4-2h-2V7h2zm0 8v-2h-2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 11v2h12v2h2v-2h2v-2h-2V9h-2v2zm10-4h2v2h-2zm0 0h-2V5h2zm0 10h2v-2h-2zm0 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 21V3h18v18zM19 5H5v14h14zM7 13v-2h6V9h2v2h2v2h-2v2h-2v-2zm4 2h2v2h-2zm0-8v2h2V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 20h2V8h2V6h-2V4h-2v2H9v2h2zM7 10V8h2v2zm0 0v2H5v-2zm10 0V8h-2v2zm0 0v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 21h18V3H3zM19 5v14H5V5zm-8 12h2v-6h2V9h-2V7h-2v2H9v2h2zm-2-4v-2H7v2zm8 0h-2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsHorizontal(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 9V7h2v2zm2 6v-2h-4v-2h4V9h2v2h2v2h-2v2zm0 0v2h-2v-2zm-6-4v2H7v2H5v-2H3v-2h2V9h2v2zm-4 4h2v2H7zm2-8v2H7V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsVertical(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 11h2V7h2v2h2V7h-2V5h-2V3h-2v2H9v2H7v2h2V7h2zm0 2h2v4h2v2h-2v2h-2v-2H9v-2h2zm-2 4v-2H7v2zm6 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArtText(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 7h10v10H2zm8 8V9H4v6zm12-8h-8v2h8zm-8 4h8v2h-8zm8 4h-8v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Article(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v18h18V3zm14 2v14H5V5zm-2 2H7v2h10zM7 11h10v2H7zm7 4H7v2h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArticleMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 1H1v18h18V1zm14 2v14H3V3zm4 18H5v2h18V5h-2zM15 5H5v2h10zM5 9h10v2H5zm7 4H5v2h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AspectRatio(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h20v16H2zm2 14h16V6H4zM8 8h2v2H8v2H6V8zm8 8h-2v-2h2v-2h2v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h16v12H8V8h8v6h2V6H6v12h14v2H4zm10 10v-4h-4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5v14H5V3h14v18H9V7h6v10h-2V9h-2v10h6V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioDevice(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h4v2H4v8h4v2H2V4zm6 0h10v2h-8v12h8v2H10zm12 0h-2v16h2zm-7 4h2v2h-2zm3 4h-4v4h4zM8 18H4v2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Avatar(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm16 16V5H5v14zM14 7h-4v4h4zm1 6H9v2H7v2h2v-2h6v2h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backburger(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 7h10v2H11zm-8 4h2V9h2v2h14v2H7v2H5v-2H3zm4 4v2h2v-2zm0-6V7h2v2zm14 6H11v2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Battery(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5H2v14h18v-4h2V9h-2V5zm14 2v10H4V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharging(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5H2v14h6v-2H4V7h4V5zm10 0h6v4h2v6h-2v4h-6v-2h4V7h-4zm-4 2h2v4h4v2h-2v2h-2v2h-2v-4H6v-2h2V9h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5H2v14h18v-4h2V9h-2V5zm0 2v10H4V7zM8 9H6v6h2zm2 0h2v6h-2zm6 0h-2v6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryOne(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5H2v14h18v-4h2V9h-2V5zm14 2v10H4V7zM8 9H6v6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryTwo(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5H2v14h18v-4h2V9h-2V5zm14 2v10H4V7zM6 9h2v6H6zm6 0h-2v6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 4h2v12h10V8h10v2h-8v6h8v-6h2v10h-2v-2H2v2H0zm3 5h2v4H3zm6 4v2H5v-2zm0-4h2v4H9zm0 0H5V7h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3h2v2h2v2H9v4h8v2H9v4h8v2h-2v2h-2v-2h-2v2H9v-2H5v-2h2v-4H5v-2h2V7H5V5h4V3h2v2h2zm4 14v-4h2v4zm0-6V7h2v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3h-2v2h2v2h2v2h-2v2h2V9h2V7h-2V5h-2zm-2 0h-2v6H9V7H7V5H5v2h2v2h2v2h2v2H9v2H7v2H5v2h2v-2h2v-2h2v6h2zm2 8h-2v2h2v2h2v2h-2v2h-2v2h2v-2h2v-2h2v-2h-2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2h12v20H4V2zm4 8h-2v2H8V4H6v16h12V4h-4v8h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpen(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h8v2H3v12h8V5h2v12h8V5h-8V3h10v16H13v2h-2v-2H1V3zm16 7h-4v2h4zm-4-3h4v2h-4zm2 6h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H6v2h12v16h-2v-2h-2v-2h-4v2H8v2H6V2H4v20h4v-2h2v-2h4v2h2v2h4V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmarks(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18V2H7v2h12v14zM5 6H3v16h4v-2h2v-2h2v2h2v2h4V6zm8 14v-2h-2v-2H9v2H7v2H5V8h10v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3h8v4h6v14H2V7h6zm2 4h4V5h-4zM4 9v10h16V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseAccount(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3H8v4H2v14h20V7h-6zm-2 4h-4V5h4zM4 19V9h16v10zm6-8h4v3h-4zm-2 4h8v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseCheck(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3H8v4H2v14h20V7h-6zm-2 4h-4V5h4zM4 19V9h16v10zm10-8h2v2h-2zm-2 4v-2h2v2zm-2 0h2v2h-2zm0 0H8v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3H8v4H2v14h12v-2H4V9h16v4h2V7h-6zm-2 4h-4V5h4zm4 8h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2h2v-2h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseDownload(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3h8v4h6v14h-5v-2h3V9H4v10h3v2H2V7h6zm6 2h-4v2h4zm-3 6h2v6h2v2h-2v2h-2v-2H9v-2h2zm-2 6H7v-2h2zm6 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3h8v4h6v6h-2V9H4v10h10v2H2V7h6zm6 2h-4v2h4zm2 12h6v2h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcasePlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3h8v4h6v4h-2V9H4v10h8v2H2V7h6zm2 4h4V5h-4zm7 14h2v-3h3v-2h-3v-3h-2v3h-3v2h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseSearch(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3H8v4H2v14h10v-2H4V9h16v2h2V7h-6zm-2 4h-4V5h4zm6 6h-6v6h6v2h2v-2h-2zm-4 4v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseSearchOne(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3H8v4H2v14h7v-2H4V9h18V7h-6zm-2 4h-4V5h4zm0 4h8v2h-8zm0 10h-2v-8h2zm8 0v2h-8v-2zm0 0h2v-8h-2zm-6-6h2v2h2v2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseUpload(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3h8v4h6v14h-5v-2h3V9H4v10h3v2H2V7h6zm6 2h-4v2h4zm-3 16h2v-6h2v2h2v-2h-2v-2h-2v-2h-2v2H9v2H7v2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2h2v4h4V2h2v4h2v3h2v2h-2v2h4v2h-4v2h2v2h-2v3H6v-3H4v-2h2v-2H2v-2h4v-2H4V9h2V6h2zm8 6H8v3h8zm-5 5H8v7h3zm2 7h3v-7h-3zM4 9H2V7h2zm0 10v2H2v-2zm16 0h2v2h-2zm0-10V7h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h18v20H3zm12 16v2h4V4H5v16h4v-2zM7 6h2v2H7zm6 0h-2v2h2zm2 0h2v2h-2zm-6 4H7v2h2zm2 0h2v2h-2zm6 0h-2v2h2zM7 14h2v2H7zm6 0h-2v2h2zm4 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingCommunity(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2h2v20H2v-8h2v6h4v-4h2v4h4v-6h2v6h4V4H10v2H8V2zm-8 10h2v2h-2zm-2-2h2v2h-2zm-2 0V8h2v2zm-2 2v-2h2v2zm0 0H4v2h2zm10-6h2v2h-2zm-2 0h-2v2h2zm2 4h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingSkyscraper(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2h4v5h2v2h-2v11h4v-9h2v9h2v2H2v-2h2V8h2v12h6V4h-2zM8 6V4h2v2zm0 0H6v2h2zm10 5h-2V9h2zm-8-1H8v2h2zm-2 4h2v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buildings(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h14v4h6v16H2zm18 6h-4v2h2v2h-2v2h2v2h-2v2h2v2h2zm-6-4H4v16h2v-2h6v2h2zM6 6h2v2H6zm6 0h-2v2h2zm-6 4h2v2H6zm6 0h-2v2h2zm-6 4h2v2H6zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bulletlist(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 11V5h6v6zm4-2V7H4v2zm16-4H10v2h12zm0 4H10v2h12zm-12 4h12v2H10zm12 4H10v2h12zM2 13v6h6v-6zm4 2v2H4v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullseye(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H6v2H4v2H2v12h2v2h2v2h12v-2h2v-2h2V6h-2V4h-2zm0 2v2h2v12h-2v2H6v-2H4V6h2V4zm-8 6h4v4h-4zM8 6h8v2H8zm0 10H6V8h2zm8 0v2H8v-2zm0 0h2V8h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BullseyeArrow(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h10v2H6zM4 6V4h2v2zm0 12H2V6h2zm2 2H4v-2h2zm12 0H6v2h12zm2-2v2h-2v-2zm0 0h2V8h-2zM12 6H8v2H6v8h2v2h8v-2h2v-4h-2v4H8V8h4zm2 8v-4h2V8h2V6h4V4h-2V2h-2v4h-2v2h-2v2h-4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2h14v2H5zm0 2v6h14V4h2v16h-2v2h-4v-2H9v2H5v-2H3V4zm0 14h14v-6H5zm2-4h2v2H7zm10 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cake(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h2v2H6zm2 3H6v3H2v9h6v-2h2v2h4v-2h2v2h6V8h-4V5h-2v3h-3V5h-2v3H8zm12 10h-4v-3h-2v3h-4v-3H8v3H4v-5h16zM2 20h20v2H2zM13 2h-2v2h2zm3 0h2v2h-2zM2 17h2v3H2zm18 0h2v3h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2H3v20h18V2zm14 18H5V4h14zM17 6H7v4h10zM7 12h2v2H7zm6 0h-2v2h2zm2 0h2v2h-2zm-6 4H7v2h2zm2 0h2v2h-2zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6zM5 8h14V6H5zm0 2v10h14V10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarAlert(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5V4H5v2H3v14h14V6h-2V4h-2v2H7zm-2 5V8h10v2zm0 2h10v6H5zm16-3V8h-2v6h2zm0 6h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarArrowLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v8h2v-2h14v10h-8v2h10V4h-4zm2 6H5V6h14zm-6 8H7v-2h2v-2H7v2H5v2H3v2h2v2h2v2h2v-2H7v-2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarArrowRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h10v-2H5V10h14v2h2V4h-4zM7 6h12v2H5V6zm14 10h-2v-2h-2v-2h-2v2h2v2h-6v2h6v2h-2v2h2v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheck(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6zm4 6V6H5v2zm0 2H5v10h14zm-3 2v2h-2v-2zm-4 4v-2h2v2zm-2 0h2v2h-2zm0 0H8v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarExport(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h4v-2H5V10h14v10h-2v2h4V4h-4zM7 6h12v2H5V6zm6 6h-2v6H9v-2H7v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarGrid(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm2 2v2h14V5zm14 4h-6v2h6zm0 4h-6v2h6zm0 4h-6v2h6zm-8 2v-2H5v2zm-6-4h6v-2H5zm0-4h6V9H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarImport(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h4v-2H5V10h14v10h-2v2h4V4h-4zM7 6h12v2H5V6zm6 16h-2v-6H9v-2h2v-2h2v2h2v2h-2zm2-6v2h2v-2zm-6 0v2H7v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4zM7 6h12v2H5V6zM5 20V10h14v10zm10-6H9v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMonth(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6zM9 6H5v2h14V6zm-4 4v10h14V10zm2 2h2v2H7zm6 0h-2v2h2zm2 0h2v2h-2zm-6 4H7v2h2zm2 0h2v2h-2zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h2v2h4v14H5V4h4V2h2v2h6zm-6 4H7v2h14V6zm-4 4v6h14v-6zM3 20h16v2H1V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMultipleCheck(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h2v2h4v10h-2v-4H7v6h6v2H5V4h4V2h2v2h6zm-6 4H7v2h14V6zm2 14v2H1V8h2v12zm2-2h2v2h-2zm4 2v2h-2v-2zm2-2h-2v2h2zm0 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6zM9 6H5v2h14V6zm-4 4v10h14V10zm6 2h2v2h2v2h-2v2h-2v-2H9v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarRange(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4zM7 6h12v2H5V6zM5 20V10h14v10zm4-8H7v2h2zm2 0h2v2h-2zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarRemove(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4zM7 6h12v2H5V6zM5 20V10h14v10zm6-4H9v2h2zm0-2v-2H9v2zm2 0h-2v2h2v2h2v-2h-2zm0 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSearch(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2h2v2h4v8h-2v-2H5v10h6v2H3V4h4V2h2v2h6zM9 6H5v2h14V6zm8 6v2h-4v-2zm-4 6h-2v-4h2zm4 0h-4v2h6v2h2v-2h-2v-6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSortAscending(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5H8v2H4V5H2v2H0v12h12V7h-2zM2 9h8v2H2zm0 8v-4h8v4zM20 7h-2v8h-2v-2h-2v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSortDescending(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5H8v2H4V5H2v2H0v12h12V7h-2zM2 9h8v2H2zm0 8v-4h8v4zm18 2h-2v-8h-2V9h2V7h2v2h2v2h-2zm2-8v2h2v-2zm-6 0v2h-2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarText(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2h2v2h4v18H3V4h4V2h2v2h6zM9 6H5v2h14V6zm-4 4v10h14V10zm2 2h8v2H7zm4 6v-2H7v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarToday(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4zM7 6h12v2H5V6zM5 20V10h14v10zm6-4v-4H7v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTomorrow(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4zM7 6h12v2H5V6zM5 20V10h14v10zm12-2v-4h-4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarWeek(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4zM7 6h12v2H5V6zM5 20V10h14v10zm12-8H7v2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarWeekBegin(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4zM7 6h12v2H5V6zM5 20V10h14v10zm4-8H7v6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarWeekend(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4zM7 6h12v2H5V6zM5 20V10h14v10zm12-8h-2v6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3H7v2H2v16h20V5h-5V3zm8 4h3v12H4V7h5V5h6v2zm-7 2h4v2h-4zm4 6h-4v2h4zh2v-4h-2zm-6-4h2v4H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraAdd(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2H3v3H0v2h3v3h2V7h3V5H5zm12 1h-7v2h5v2h5v12H5v-7H3v9h19V5h-5zm-7 6h4v2h2v4h-2v2h-4v-2h4v-4h-4zm-2 2h2v4H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4H2v16h20V4zm16 2v12H4V6zM8 8H6v2h2zm4 0h4v2h-4zm-2 2h2v4h-2zm6 4h2v-4h-2zm0 0h-4v2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraFace(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 3h10v2h5v16H2V7h2v12h16V7h-5V5H9v2H2V5h5zm7 12h-4v2h4zm-4-2v2H8v-2zm0-2V9H8v2zm6 2v2h-2v-2zm0-2V9h-2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4H7v2H5v2H3v12h4v-2h10v2h4V8h-2V6h-2zm0 2v2h2v2H5V8h2V6zm2 10H5v-4h14zm-2-3h-2v2h2zM7 13h2v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Card(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h20v16H2zm18 14V6H4v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardId(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h20v16H2zm2 2v4h16V6zm16 6H10v2h10zm0 4h-4v2h4zm-6 2v-2H4v2zM4 14h4v-2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardPlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 4H2v16h10v-2H4V6h16v4h2zm-3 13h3v-2h-3v-3h-2v3h-3v2h3v3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardStack(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h18v12H2V4zm16 10V6H4v8zm2 4H2v2h20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardText(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4H2v16h20V4zm0 2h16v12H4zm2 2h12v2H6zm0 4h10v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h4v4h16v11H4V4H2zm4 13h14V8H6zm0 4h3v3H6zm14 0h-3v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cast(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3h18v18h-8v-2h6V5H4v4H2V3zm0 16H2v2h2zm-2-4h4v2H2zm8-4H2v2h8v8h2V11zm-4 4h2v6H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CellularSignalOff(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2H2v2h2v2H2v2h2V6h2v2h2V6H6V4h2V2H6v2H4zm12 2v16h6V4zm2 2h2v12h-2zm-9 4v10h6V10zm2 8v-6h2v6zm-3-4v6H2v-6zm-2 4v-2H4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CellularSignalOne(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4v16h6V4zm2 2h2v12h-2zm-9 4v10h6V10zm2 8v-6h2v6zm-3-4H2v6h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CellularSignalThree(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4h6v16h-6zM2 14h6v6H2zm13-4H9v10h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CellularSignalTwo(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4v16h6V4zm4 2v12h-2V6zM2 14h6v6H2zm13-4H9v10h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CellularSignalZero(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 4v16h-6V4zm-2 2h-2v12h2zm-5 4v10H9V10zm-2 8v-6h-2v6zm-5-4v6H2v-6zm-2 4v-2H4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v18h18V3zm14 2v14H5V5zM9 11H7v6h2zm2-4h2v10h-2zm6 6h-2v4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartAdd(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h10v2H5v14h14v-8h2v10H3zm6 8H7v6h2zm2-4h2v10h-2zm6 6h-2v4h2zm0-10h2v2h2v2h-2v2h-2V7h-2V5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBar(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 5h2v14h-2zm-2 4H9v10h2zm-4 4H5v6h2zm12 0h-2v6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3H3v18h18V11h-2v8H5V5h8zm-6 8h2v6H7zm6-4h-2v10h2zm2 6h2v4h-2zm2-6h-2v2h2zm0-2V3h-2v2zm2 0h-2v2h2v2h2V7h-2zm0 0V3h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3H3v18h18V11h-2v8H5V5h8zm-6 8h2v6H7zm6-4h-2v10h2zm2 6h2v4h-2zm6-8h-6v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2H1v16h18V2zm0 2h14v12H3zm18 2v14H5v2h18V6zM7 8H5v6h2zm2-2h2v8H9zm6 4h-2v4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H2v20h2V4h16v12H6v2H4v2h2v-2h16V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6h2v2h-2zm-2 4V8h2v2zm-2 2v-2h2v2zm-2 2h2v-2h-2zm-2 2h2v-2h-2zm-2 0v2h2v-2zm-2-2h2v2H6zm0 0H4v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckDouble(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6h2v2h-2zm-2 4V8h2v2zm-2 2v-2h2v2zm-2 2v-2h2v2zm-2 2v-2h2v2zm-2 0h2v2H5zm-2-2h2v2H3zm0 0H1v-2h2zm8 2h2v2h-2zm4-2v2h-2v-2zm2-2v2h-2v-2zm2-2v2h-2v-2zm2-2h-2v2h2zm0 0h2V6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkbox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v18h18V3zm0 2h14v14H5zm4 7H7v2h2v2h2v-2h2v-2h2v-2h2V8h-2v2h-2v2h-2v2H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxOn(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm16 16V5H5v14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checklist(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h2v2h-2zm-2 4V6h2v2zm-2 0h2v2h-2zm0 0h-2V6h2zM3 6h8v2H3zm8 10H3v2h8zm7 2v-2h2v-2h-2v2h-2v-2h-2v2h2v2h-2v2h2v-2zm0 0v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chess(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v4h4v4H4v4h4v4h4v-4h4v4h4v-4h-4v-4h4V8h-4V4h-4v4H8V4zm8 8H8v4h4zm0-4v4h4V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8H5v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2V8h-2v2h-2v2h-2v2h-2v-2H9v-2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5v2h-2V5zm-4 4V7h2v2zm-2 2V9h2v2zm0 2H8v-2h2zm2 2v-2h-2v2zm0 0h2v2h-2zm4 4v-2h-2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5v2h2V5zm4 4V7h-2v2zm2 2V9h-2v2zm0 2h2v-2h-2zm-2 2v-2h2v2zm0 0h-2v2h2zm-4 4v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 16H5v-2h2v-2h2v-2h2V8h2v2h2v2h2v2h2v2h-2v-2h-2v-2h-2v-2h-2v2H9v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsHorizontal(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 9V7h2v2zm-2 2V9h2v2zm0 2H4v-2h2zm2 2v-2H6v2zm0 0h2v2H8zm8-6V7h-2v2zm2 2V9h-2v2zm0 2v-2h2v2zm-2 2v-2h2v2zm0 0v2h-2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsVertical(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4h2v2h-2zM9 8V6h2v2zm0 0v2H7V8zm6 0h-2V6h2zm0 0h2v2h-2zm-6 8H7v-2h2zm2 2H9v-2h2zm2 0v2h-2v-2zm2-2h-2v2h2zm0 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3H7v2H5v2H3v10h2v2h2v2h10v-2h2v-2h2V7h-2V5h-2zm0 2v2h2v10h-2v2H7v-2H5V7h2V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2h6v2h4v18H4V4h4V2zm6 4v2H8V6H6v14h12V6zm-2 0V4h-4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5v2H3v14h2v2h14v-2h2V5h-2zm0 2v14H5V5zm-8 2h2v6h4v2h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5h2v2H5zm4 4H7V7h2zm2 2H9V9h2zm2 0h-2v2H9v2H7v2H5v2h2v-2h2v-2h2v-2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2zm2-2v2h-2V9zm2-2v2h-2V7zm0 0V5h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v18h18V3zm14 2v14H5V5zm-8 4H9V7H7v2h2v2h2v2H9v2H7v2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2v-2h2V9h2V7h-2v2h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4h-6v2H8v2H4v2H2v2H0v6h2v2h20v-2h2v-6h-2v-2h-2V8h-2V6h-2zm2 8h4v6H2v-6h2v-2h4v2h2v-2H8V8h2V6h6v2h2zm0 0v2h-2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDone(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4h-6v2H8v2H4v2H2v2H0v6h2v2h20v-2h2v-6h-2v-2h-2V8h-2V6h-2zm0 2v2h2v4h4v6H2v-6h2v-2h4V8h2V6zm-6 6H8v2h2v2h2v-2h2v-2h2v-2h-2v2h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4h6v2h-6zM8 8V6h2v2zm-4 2V8h4v2zm-2 2v-2h2v2zm0 6H0v-6h2zm0 0h5v2H2zM18 8h-2V6h2zm4 4h-4V8h2v2h2zm0 6v-6h2v6zm0 0v2h-5v-2zm-11 2h2v-2h2v-2h2v-2h-4V9h-2v5H7v2h2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudMoon(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2h-8v2H8v2H6v4h2V6h2V4h4v2h-2v4h2v2h4v-2h2v4h-2v2h2v-2h2V6h-2v2h-2v2h-4V6h2V4h2zM8 14v-2h4v2zm0 2v-2H4v2H2v4h2v2h10v-2h2v-4h-2v-2h-2v2h2v4H4v-4zm0 0h2v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSun(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 0h2v4h-2zm1 12H8v2H4v2H2v4h2v2h10v-2h2v-4h-2v-2h-2zm0 2v2h2v4H4v-4h4v2h2v-2H8v-2zM8 6h6v2H8zm0 2v2H6V8zm8 2h-2V8h2zm0 0h2v2h-2zm4-8h2v2h-2zm0 2v2h-2V4zM2 2h2v2H2zm2 2h2v2H4zm20 7h-4v2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4h6v2h-6zM8 8V6h2v2zm-4 2V8h4v2zm-2 2v-2h2v2zm0 6H0v-6h2zm0 0h7v2H2zM18 8h-2V6h2zm4 4h-4V8h2v2h2zm0 6v-6h2v6zm0 0v2h-7v-2zM11 9h2v2h2v2h2v2h-4v5h-2v-5H7v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cocktail(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H3v4h2v2h2v2h2v2h2v6H7v2h10v-2h-4v-6h2v-2h2V9h2V7h2V3zm0 4H5V5h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5h2v2H8zM6 7h2v2H6zM4 9h2v2H4zm-2 2h2v2H2zm2 2h2v2H4zm2 2h2v2H6zm2 2h2v2H8zm8-12h-2v2h2zm2 2h-2v2h2zm2 2h-2v2h2zm2 2h-2v2h2zm-2 2h-2v2h2zm-2 2h-2v2h2zm-2 2h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h18v7h-4v5H4zm14 5h2V6h-2zm-2-3H6v8h10zm3 14H3v-2h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 3H5v4h2zm4 0H9v4h2zm2 0h2v4h-2zm8 6H3v12h14v-5h4zm-2 5h-2v-3h2zM5 11h10v8H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coin(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h12v2H6zM4 6V4h2v2zm0 12V6H2v12zm2 2v-2H4v2zm12 0v2H6v-2zm2-2v2h-2v-2zm0-12h2v12h-2zm0 0V4h-2v2zm-9-1h2v2h3v2h-6v2h6v6h-3v2h-2v-2H8v-2h6v-2H8V7h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collapse(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3h-2v2h-2v2h-2V5H9V3H7v2h2v2h2v2h2V7h2V5h2zM4 13h16v-2H4zm9 4h-2v-2h2zm2 2h-2v-2h2zm0 0h2v2h-2zm-6 0h2v-2H9zm0 0H7v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorsSwatch(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2h8v20H12V2zm6 2h-6v16h6zM10 20H4v-6h6v-2H6v-2H4V8h2V6h2V4h2V2H8v2H6v2H4v2H2v2h2v2H2v10h8zm8-4h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Command(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2H2v8h2zm16 0h2v8h-2zm-6 6h-4V2H4v2h4v4H4v2h4v4H4v2h4v4H4v2h6v-6h4v6h2v-6h4v-2h-4v-4h4V8h-4V2h-2zm-4 6v-4h4v4zM20 2h-4v2h4zM2 14h2v8H2zm14 6h4v2h-4zm6-6h-2v8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2H2v14h2V4h16v12h-8v2h-2v2H8v-4H2v2h4v4h4v-2h2v-2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contact(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3H0v18h24V3zm20 2v14H2V5zM10 7H6v4h4zm-6 6h8v4H4zm16-6h-6v2h6zm-6 4h6v2h-6zm6 4h-6v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 3H0v18h16v-2H2V5h20v10h2V3zM6 7h4v4H6zm0 8H4v2h2zm4 0H6v-2h4zm0 0v2h2v-2zm4-8h6v2h-6zm6 4h-6v2h6zm-6 4h2v2h-2zm8 4h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2zm0 0h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3h20v14H4zm18 12V5H6v10zm-2 4H2V7H0v14h20zM9 7h2v2H9zm3 4H8v2h4zm2-4h6v2h-6zm6 4h-6v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactPlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h22v11h-2V5H2v14h12v2H0V3zm8 4H6v4h4zm-6 6h8v4H4zm16-6h-6v2h6zm-6 4h6v2h-6zm3 4h-3v2h3zm4 6v3h-2v-3h-3v-2h3v-3h2v3h3v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h11v2H6v13H4zm4 4h12v16H8zm2 2v12h8V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 16H8v2H6v-2H4v-2h2v-2h2v2h10V4h2v12zM8 12v-2h2v2zm0 6v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 16h10v2h2v-2h2v-2h-2v-2h-2v2H6V4H4v12zm10-4v-2h-2v2zm0 6v2h-2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLeftDown(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6v10H6v2h2v2h2v-2h2v-2h-2V6h10V4H8zm4 10h2v-2h-2zm-6 0H4v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLeftUp(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 18V8H6V6h2V4h2v2h2v2h-2v10h10v2H8zm4-10h2v2h-2zM6 8H4v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerRightDown(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6v10h2v2h-2v2h-2v-2h-2v-2h2V6H4V4h12zm-4 10h-2v-2h2zm6 0h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerRightUp(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 18V8h2V6h-2V4h-2v2h-2v2h2v10H4v2h12zM12 8h-2v2h2zm6 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 8H8V6H6v2H4v2h2v2h2v-2h10v10h2V8zM8 12v2h2v-2zm0-6V4h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 8h10V6h2v2h2v2h-2v2h-2v-2H6v10H4V8zm10 4v2h-2v-2zm0-6V4h-2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h16v2H4v2h16v4H4v6h16v2H2V4zm18 0h-2v16h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H2v16h12v-2H4v-6h16V8H4V6h16zm0 0h2v8h-2zm2 14h-2v-2h2v-2h-2v2h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H2v16h12v-2H4v-6h16V8H4V6h16zm0 0h2v8h-2zm2 12h-6v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3h16v2H3v2h14v4H3v4h14v2H1zm18 0h-2v14h2zM5 19h16v2H5zM23 7h-2v14h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardPlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h18v2H4v2h16v4H4v6h10v2H2zm20 0h-2v8h2zm-4 10h2v2h2v2h-2v2h-2v-2h-2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardSettings(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H2v16h18v-2H4v-6h16V8H4V6h16zm0 0h2v16h-2zm-7 18h-2v2h2zm2 0h2v2h-2zm-6 0H7v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardWireless(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8v2H6v2h2V4h8v2h2V4h-2zM8 8h2v2H8zm6 0V6h-4v2zm0 0h2v2h-2zM4 11h16v12H4zm14 10v-3H6v3zm0-6v-2H6v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2H6v4H2v2h14v14h2v-4h4v-2h-4V6H8zm0 8H6v8h8v-2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h2v2H2zm4 4H4V4h2zm2 2H6V6h2zm2 2V8H8v2zm4 0h-4v4H2v8h8v-8h4v8h8v-8h-8zm2-2v2h-2V8zm2-2v2h-2V6zm2-2h-2v2h2zm0 0V2h2v2zM4 20v-4h4v4zm12 0v-4h4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashbaord(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h8v10H3zm2 2v6h4V5zm8-2h8v6h-8zm2 2v2h4V5zm-2 6h8v10h-8zm2 2v6h4v-6zM3 15h8v6H3zm2 2v2h4v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Debug(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 2h2v2H6zm4 9h4v2h-4zm4 4h-4v2h4z"/><path d="M16 4h-2v2h-4V4H8v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h12v-3h2v2h2v-2h-2v-2h-2v-2h4v-2h-4v-2h2V9h2V7h-2v2h-2V6h-2zM8 20V8h8v12zm8-16V2h2v2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugCheck(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2H6v2h2v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h6v-2H8V8h8v6h2v-3h2V9h2V7h-2v2h-2V6h-2V4h2V2h-2v2h-2v2h-4V4H8zm6 9h-4v2h4zm-4 4h2v2h-2zm4 3h2v2h-2zm4 2v2h-2v-2zm2-2h-2v2h2zm0 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugOff(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2h2v2h-2zm4 7h-2V6h-2V4h-2v2h-2v2h4v5h2v2h4v-2h-4v-2h2zm0 0V7h2v2zM8 20v-9H6V9H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h10v-2zm2-5h2v2h-2zM2 2h2v2H2zm4 4H4V4h2zm2 2H6V6h2zm2 2H8V8h2zm0 0v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2v-2h-2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugPause(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2H6v2h2v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h8v-2H8V8h8v6h2v-3h2V9h2V7h-2v2h-2V6h-2V4h2V2h-2v2h-2v2h-4V4H8zm6 9h-4v2h4zm-4 4h4v2h-4zm6 1h2v6h-2zm6 0h-2v6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugPlay(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h2v2H6zm10 2h-2v2h-4V4H8v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h8v-2H8V8h8v3h4V9h2V7h-2v2h-2V6h-2zm0 0V2h2v2zm-6 7h4v2h-4zm4 4h-4v2h4zm4-2h-2v10h2v-2h2v-2h2v-2h-2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugStop(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h2v2H6zm10 2h-2v2h-4V4H8v2H6v3H4V7H2v2h2v2h2v2H2v2h4v2H4v2H2v2h2v-2h2v3h8v-2H8V8h8v6h2v-3h2V9h2V7h-2v2h-2V6h-2zm0 0V2h2v2zm-6 7h4v2h-4zm4 4h-4v2h4zm8 1h-6v6h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 5H7v2H5v2H3v2H1v2h2v2h2v2h2v2h16V5zM7 17v-2H5v-2H3v-2h2V9h2V7h14v10zm8-6h-2V9h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2zm0 0V9h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deskphone(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm2 2v6h8V5zm10 0v14h4V5zm-2 14v-2h-3v2zm-5 0v-2H5v2zm-3-4h3v-2H5zm5-2v2h3v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceLaptop(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4H4v12h16V4zm12 2v8H6V6zm4 12H2v2h20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DevicePhone(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3h12v18H6zm10 16V5h-2v2h-4V5H8v14zm-5-4h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceTablet(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2H4v20h16V2zm12 2v16H6V4zm-5 12h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceTv(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 20h20V6h-7V4h-2v2h-2V4H9v2H2zM9 4V2H7v2zm6 0h2V2h-2zm5 4v10H4V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceTvSmart(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h18v14h-6v2H8v-2H2V4zm16 12V6H4v10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceVibrate(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3H6v18h12V3zm8 2v14H8V5zm-3 10h-2v2h2zm7-8h2v2h-2zm2 4V9h2v2zm0 2h-2v-2h2zm0 2v-2h2v2zm0 0v2h-2v-2zM2 17h2v-2H2v-2h2v-2H2V9h2V7H2v2H0v2h2v2H0v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceWatch(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2h8v4h5v12h-5v4H8v-4H3V6h5zM5 16h14V8H5zm6-6h2v2h2v2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Devices(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h16v6h4v14H12v-6H2zm14 6V4H4v10h8V8zm-6-2H6v2h4zm10 14V10h-6v10zm-4-4h2v2h-2zM6 10h4v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dice(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v18h18V3zm14 2v14H5V5zM9 7H7v2h2zm6 0h2v2h-2zm-6 8H7v2h2zm6 0h2v2h-2zm-2-4h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2h2v4h6v2H7v3H5V6h6zM5 18h6v4h2v-4h6v-2H5zm14-7H5v2h12v3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Downasaur(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4h14v2h2v6h-8v2h6v2h-4v2h-2v2H2V8h2V6h2zm2 6h2V8H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 17V3h-2v10H9v-2H7v2h2v2h2v2zm8 2v-4h-2v4H5v-4H3v6h18zm-8-6v2h2v-2h2v-2h-2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Draft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2h-4v2H8v2H6v2H4v2H2v12h20V10h-2V8h-2V6h-2V4h-2zm0 2v2h2v2h2v4h-2v2h-2v2h-4v-2H8v-2H6V8h2V6h2V4zm-8 8v2h2v2h2v2h4v-2h2v-2h2v-2h2v8H4v-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragAndDrop(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v2h2zm14 4h2v6h-2V9H9v10h4v2H7V7zM7 3h2v2H7zM5 7H3v2h2zm-2 4h2v2H3zm2 4H3v2h2zm6-12h2v2h-2zm6 0h-2v2h2zm-2 14v-2h6v2h-2v2h-2v2h-2zm4 2v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drop(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2h-2v2H9v4H7v4H5v6h2v2h2v2h6v-2h2v-2h2v-6h-2V8h-2V4h-2zm0 2v4h2v4h2v6h-2v2H9v-2H7v-6h2V8h2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropArea(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v2h2zm2 0h2v2H7zm6 0h-2v2h2zm2 0h2v2h-2zm4 0h2v2h-2zM3 7h2v2H3zm2 4H3v2h2zm-2 4h2v2H3zm2 4H3v2h2zm2 0h2v2H7zm6 0h-2v2h2zm6-8h2v2h-2zm2-4h-2v2h2zm-6 10v-2h6v2h-2v2h-2v2h-2zm4 2v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropFull(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2h2v2h2v4h2v4h2v6h-2v2h-2v2H9v-2H7v-2H5v-6h2V8h2V4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropHalf(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2h-2v2H9v4H7v4H5v6h2v2h2v2h6v-2h2v-2h2v-6h-2V8h-2V4h-2zm0 2v4h2v4h2v3H7v-3h2V8h2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Duplicate(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3h12v4h4v14H7v-4H3V3zm10 4V5H5v10h2V7zM9 17v2h10V9H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DuplicateAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1H3v14h10v2h-2v2h2v-2h2v-2h2v-2h-2v-2h-2V9h-2v2h2v2H5V3h12V1zm4 4H7v6h2V7h10v14H9v-4H7v6h14V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2h-2v2h-2v2h-2v2h-2v2H8v2H6v2H4v2H2v6h6v-2h2v-2h2v-2h2v-2h2v-2h2v-2h2V8h2V6h-2V4h-2zm0 8h-2v2h-2v2h-2v2h-2v2H8v-2H6v-2h2v-2h2v-2h2V8h2V6h2v2h2zM6 16H4v4h4v-2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2h-2v2h2zM4 4h6v2H4v14h14v-6h2v8H2V4zm4 8H6v6h6v-2h2v-2h-2v2H8zm4-2h-2v2H8v-2h2V8h2V6h2v2h-2zm2-6h2v2h-2zm4 0h2v2h2v2h-2v2h-2v2h-2v-2h2V8h2V6h-2zm-4 8h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 4h10v2H9v3h7v2H9v2h7v2H9v3h10v2H7v-5H5v-2h2v-2H5V9h2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 5h2v2h2v2h2V7h-2V5h-2V3h-2zM9 7V5h2v2zm0 0v2H7V7zm-5 6h16v-2H4zm9 6h-2v-2H9v-2H7v2h2v2h2v2h2zm2-2h-2v2h2zm0 0h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11V3h-8v2h4v2h-2v2h-2v2h-2v2H9v2h2v-2h2v-2h2V9h2V7h2v4zM11 5H3v16h16v-8h-2v6H5V7h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6h8v2H8zm-4 4V8h4v2zm-2 2v-2h2v2zm0 2v-2H0v2zm2 2H2v-2h2zm4 2H4v-2h4zm8 0v2H8v-2zm4-2v2h-4v-2zm2-2v2h-2v-2zm0-2h2v2h-2zm-2-2h2v2h-2zm0 0V8h-4v2zm-10 1h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosed(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7h2v2H0zm4 4H2V9h2zm4 2v-2H4v2H2v2h2v-2zm8 0H8v2H6v2h2v-2h8v2h2v-2h-2zm4-2h-4v2h4v2h2v-2h-2zm2-2v2h-2V9zm0 0V7h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 22h18V8h-2V6h-2v2h-2V6h2V4h-2V2H3zm2-2V4h8v6h6v10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 22H3V2h12v2h2v2h2v2h2zM17 6h-2v2h2zM5 4v16h14V10h-6V4zm8 12H7v2h6zm-6-4h10v2H7zm4-4H7v2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 22h10V8h-2V6h-2v2h-2V6h2V4h-2V2H3v12h2V4h8v6h6v10h-8zm-4-2H5v2H3v-2h2v-2H3v-2h2v2h2v-2h2v2H7zm0 0h2v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileFlash(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22h-6v-2h6V10h-6V4H5v8H3V2h12v2h2v2h2v2h2v14zM17 6h-2v2h2zM7 12h2v4h4v2h-2v2H9v2H7v-4H3v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 22h8V8h-2V6h-2v2h-2V6h2V4h-2V2H3v13h2V4h8v6h6v10h-6zm-2-3H3v-2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18H7V2h8v2h2v2h-2v2h2V6h2v2h2zM9 4v12h10v-6h-6V4zM3 6h2v14h12v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileOff(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2H3v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2v-2h-2v-2h6v4h2V8h-2V6h-2V4h-2V2H9v2h4v6h-2V8H9V6H7V4H5zm12 4v2h-2V6zM3 8h2v12h12v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22h-7v-2h7V10h-6V4H5v8H3V2h12v2h2v2h2v2h2v14zM17 6h-2v2h2zM8 19h3v-2H8v-3H6v3H3v2h3v3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fill(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2h2v2H9zm4 4V4h-2v2H9v2H7v2H5v2H3v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v6h2V12h-2v-2h-2V8h-2V6zm0 0v2h2v2h2v2h2v2h-2v2h-2v2h-2v2h-2v-2H9v-2H7v-2H5v-2h2v-2h2V8h2V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FillHalf(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2h2v2H9zm4 4V4h-2v2H9v2H7v2H5v2H3v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v6h2V12h-2v-2h-2V8h-2V6zm0 0v2h2v2h2v2h2v2H5v-2h2v-2h2V8h2V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveG(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 7H3v6h5v2H3v2h7v-6H5V9h5zm11 0h-9v10h9v-6h-4v2h2v2h-5V9h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h10v2h8v14H11v-2H5v6H3zm2 12h8v2h6V6h-8V4H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flatten(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2h2v8h2v2h-2v2h-2v-2H9v-2h2zm-2 8H7V8h2zm6 0V8h2v2zm5 6H4v2h16zm-4 4H8v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipToBack(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3H7v2h2zm0 12H7v2h2zm2-12h2v2h-2zm2 12h-2v2h2zm2-12h2v2h-2zm2 12h-2v2h2zm2-12h2v2h-2zm2 4h-2v2h2zM7 7h2v2H7zm14 4h-2v2h2zM7 11h2v2H7zm14 4h-2v2h2zM3 7h2v12h12v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipToFront(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 3H7v14h14zm-2 12H9V5h10zM5 7H3v2h2zm-2 4h2v2H3zm2 4H3v2h2zm-2 4h2v2H3zm6 0H7v2h2zm2 0h2v2h-2zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloatCenter(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4h6v8H8V4zm4 6V6h-4v4zM2 6h4v2H2zm20 0h-4v2h4zm0 4h-4v2h4zM6 10H2v2h4zm-4 4h20v2H2zm20 4H2v2h20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloatLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h6v8H2V4zm4 6V6H4v4zm14-4H12v2h10zm0 4H12v2h10zm0 4v2H2v-2zm0 6v-2H2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloatRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4h6v8h-8V4zm4 6V6h-4v4zm-8-4H2v2h10zm0 4H2v2h10zm10 4v2H2v-2zm0 6v-2H2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h8v2h10v14H2V4zm16 4H10V6H4v12h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4H2v16h20V6H12zm-2 4h10v10H4V6h6zm8 6v-2h-6v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h8v2h10v14H2V4zm16 4H10V6H4v12h16zm-6 2h2v2h2v2h-2v2h-2v-2h-2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderX(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4H2v16h20V6H12zm-2 4h10v10H4V6h6zm6 4h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2zm0 0h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 5h-2v4H6v2H4v6h2v-2h6v4h2v-2h2v-2h2v-2h2v-2h-2V9h-2V7h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forwardburger(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 7H3v2h10zm8 4h-2V9h-2V7h-2v2h2v2H3v2h14v2h-2v2h2v-2h2v-2h2zM3 15h10v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourG(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 7H3v6h5v4h2V7H8v4H5zm16 0h-9v10h9v-6h-4v2h2v2h-5V9h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourK(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h2v4h4V7h2v10H9v-4H3zm10 0h2v4h2v2h-2v4h-2zm6 8h-2v-2h2zm0 0h2v2h-2zm0-6h-2v2h2zm0 0V7h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourKbox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4H1v16h22V4zm18 2v12H3V6zM7 8H5v5h4v3h2V8H9v3H7zm8 0h-2v8h2v-3h2v3h2v-3h-2v-2h2V8h-2v3h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frame(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v18H2zm18 16V7H4v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameAdd(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v18H2zm18 16V7H4v12zm-7-7h3v2h-3v3h-2v-3H8v-2h3V9h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameCheck(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v18H2zm18 16V7H4v12zm-4-9h-2v2h-2v2h-2v-2H8v2h2v2h2v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v18H2zm18 16V7H4v12zM9 10h2v2H9zm4 2h-2v2H9v2h2v-2h2v2h2v-2h-2zm0 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrameMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v18H2zm18 16V7H4v12zM8 12h8v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gamepad(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h20v14H2zm18 12V7H4v10zM8 9h2v2h2v2h-2v2H8v-2H6v-2h2zm6 0h2v2h-2zm4 4h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gif(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h6v2H3v6h4v-2H5v-2h4v6H1V7zm14 0h6v2h-6v2h4v2h-4v4h-2V7zm-4 0h-2v10h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M19 12v8h-7m7-8h2V8h-3m1 4H5m13-4V4h-6m6 4H6m0 0V4h6M6 8H3v4h2m0 0v8h7m0 0V4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitBranch(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2h2v12h3v3h7v-7h-3V2h8v8h-3v9h-9v3H2v-8h3zm15 6V4h-4v4zM8 19v-3H4v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCommit(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7h10v4h5v2h-5v4H7v-4H2v-2h5zm2 2v6h6V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitMerge(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2H2v8h3v12h2V10h3v2h2v2h2v8h8v-8h-8v-2h-2v-2h-2zM4 8V4h4v4zm12 12v-4h4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequest(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h8v8H7v12H5V10H2zm2 2v4h4V4zm8 1h7.09v9H22v8h-8v-8h3.09V7H12zm4 11v4h4v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gps(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2v4h5v5h4v2h-4v5h-5v4h-2v-4H6v-5H2v-2h4V6h5V2zM8 8v8h8V8zm2 2h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v4h4V4zm6 0v4h4V4zm6 0v4h4V4zm4 6h-4v4h4zm0 6h-4v4h4zm-6 4v-4h-4v4zm-6 0v-4H4v4zm-4-6h4v-4H4zm6-4v4h4v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm2 2v14h14V5zm2 2h4v4H7zm6 0h4v4h-4zm-6 6h4v4H7zm6 0h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hd(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h2v4h4V7h2v10H9v-4H5v4H3zm10 8V7h6v2h-4v6h4v2h-6zm6 0V9h2v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphone(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4H5v2H3v14h7v-8H5V6h14v6h-5v8h7V6h-2zm-3 10h3v4h-3zm-8 0v4H5v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headset(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5v2H3v14h7v-8H5V4h14v6h-5v8h3v2h-6v2h8v-4h2V4h-2zm-3 10h3v4h-3zm-8 0v4H5v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2H5v2H3v2H1v6h2v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v-2h2v-2h2V6h-2V4h-2V2h-4v2h-2v2h-2V4H9zm0 2v2h2v2h2V6h2V4h4v2h2v6h-2v2h-2v2h-2v2h-2v2h-2v-2H9v-2H7v-2H5v-2H3V6h2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hidden(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6h8v2H8zm-4 4V8h4v2zm-2 2v-2h2v2zm0 2v-2H0v2zm2 2H2v-2h2zm4 2H4v-2h4zm8 0v2H8v-2zm4-2v2h-4v-2zm2-2v2h-2v-2zm0-2h2v2h-2zm-2-2h2v2h-2zm0 0V8h-4v2zM9 10h2v2H9zm4 2h-2v2H9v2h2v-2h2v2h2v-2h-2zm0 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2h-4v2H8v2H6v2H4v2H2v2h2v10h7v-6h2v6h7V12h2v-2h-2V8h-2V6h-2V4h-2zm0 2v2h2v2h2v2h2v2h-2v8h-3v-6H9v6H6v-8H4v-2h2V8h2V6h2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H6v6h2v2h2v4H8v2H6v6h12v-6h-2v-2h-2v-4h2V8h2zm-2 6h-2v2h-4V8H8V4h8zm-2 6v2h2v4H8v-4h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hq(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h2v4h4V7h2v10H9v-4H5v4H3zm10 2h2v6h-2zm6 6h-4v2h8v-2h-2V9h-2V7h-4v2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Human(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2h4v4h-4zM3 7h18v2h-6v13h-2v-6h-2v6H9V9H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HumanHandsdown(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2h4v4h-4zM7 7h10v2h-2v13h-2v-6h-2v6H9V9H7zm-2 4h2V9H5zm0 0v2H3v-2zm14 0h-2V9h2zm0 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HumanHandsup(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2h4v4h-4zM7 7h10v2h-2v13h-2v-6h-2v6H9V9H7zM5 5v2h2V5zm0 0H3V3h2zm14 0v2h-2V5zm0 0V3h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HumanHeight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h4v4H6zM3 7h10v9h-2v6H9v-6H7v6H5v-6H3zm18-4h-6v2h6zm-4 4h4v2h-4zm4 4h-6v2h6zm-6 8h6v2h-6zm6-4h-4v2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HumanHeightAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h4v4H4zM1 7h10v9H9v6H7v-6H5v6H3v-6H1zm18-5h-2v2h-2v2h-2v2h2V6h2v12h-2v-2h-2v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2V6h2v2h2V6h-2V4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HumanRun(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3H8v2H6v2h2V5h2v2h2v2h-2v2H8v2H6v2H4v-2H2v2h2v2h2v-2h4v2h2v2h-2v2h2v-2h2v-2h-2v-4h2v-2h2v2h2v2h2v-2h2v-2h-2v2h-2v-2h-2V9h2V5h-4v2h-2V5h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3H2v18h20V3zm16 2v14H4V5zm-6 4h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2zM8 7H6v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageArrowRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 1h-2v2h2v2h-6v2h6v2h-2v2h2V9h2V7h2V5h-2V3h-2zm-8 2H2v18h20v-8h-2v6H4V5h7zm1 8V9h2v2zm-2 2v-2h2v2zm-2 2v-2h2v2zm0 0v2H6v-2zm8-2h-2v-2h2zm0 0h2v2h-2zM6 7h2v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageBroken(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 3H2v18h20v-2h-2v-2h2v-2h-2v-2h2v-2h-2V9h2V7h-2V5h2zm-2 4v2h-2v2h2v2h-2v2h2v2h-2v2H4V5h14v2zm-6 2h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v-2h-2zM6 7h2v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 3H2v18h20V11h-2v8H4V5h10zM6 7h2v2H6zm14-2h-2V3h-2v2h2v2h-2v2h2V7h2v2h2V7h-2zm0 0V3h2v2zm-8 4h2v2h-2zm-2 4v-2h2v2zm-2 2h2v-2H8zm0 0v2H6v-2zm8-2h-2v-2h2zm0 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageFlash(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 0h2v4h4v2h-2v2h-2v2h-2V6h-4V4h2V2h2zM4 3h8v2H4v14h16v-7h2v9H2V3zm10 6h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2zM8 7H6v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageFrame(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 1h-2v2H9v2H7v2H2v16h20V7h-5V5h-2V3h-2zm2 6H9V5h2V3h2v2h2zM4 9h16v12H4zm10 6v-2h-2v2h-2v2H8v2h2v-2h2v-2zm2 2v-2h-2v2zm0 0v2h2v-2zM6 13v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageGallery(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h20v16h-5v2h-2v-2H9v2H7v-2H2zm5 18v2H5v-2zm10 0v2h2v-2zm3-16H4v12h16zm-8 4h2v2h-2zm-2 4v-2h2v2zm0 0v2H8v-2zm6 0h-2v-2h2zm0 0h2v2h-2zM8 6H6v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 2H4v16h20zM6 16V4h16v12zM2 4H0v18h20v-2H2zm12 2h2v2h-2zm-2 4V8h2v2zm-2 2v-2h2v2zm0 0v2H8v-2zm8-2h-2V8h2zm0 0h2v2h-2zM8 6h2v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageNew(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0h12v2H6zM4 3H2v18h20V3zm16 2v14H4V5zm-6 4h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2zM8 7H6v2h2zm10 17v-2H6v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImagePlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3h10v2H4v14h16v-8h2v10H2V3zm10 6h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2zM8 7H6v2h2zm10-4h2v2h2v2h-2v2h-2V7h-2V5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h18v20H3zm2 2v10h4v2h6v-2h4V4zm14 12h-2v2H7v-2H5v4h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxAll(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h18v20H3zm2 2v4h4v2h6V8h4V4zm14 6h-2v2H7v-2H5v4h14zm0 6h-2v2H7v-2H5v4h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxFull(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h18v20H3zm2 2v10h4v2h6v-2h4V4zm14 12h-2v2H7v-2H5v4h14zM7 6h10v2H7zm0 4h10v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h2v18H3zm16 0H5v2h14v14H5v2h16V3zm-8 6h2V7h-2zm2 8h-2v-6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invert(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm16 4h-2v2h-2v2h-2v2h-2v2H9v2H7v2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iso(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3H6v3H3v2h3v3h2V8h3V6H8zm11 2h-2v2h-2v2h-2v2h-2v2H9v2H7v2H5v2h2v-2h2v-2h2v-2h2v-2h2V9h2V7h2zm-6 13v-2h8v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kanban(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 3H3v18h18zM5 19V5h14v14zM9 7H7v8h2zm2 0h2v4h-2zm6 0h-2v10h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 3H3v18h18zM5 19V5h14v14zM9 7H7v2h2zm8 8H7v2h10zm-2-8h2v2h-2zm-2 0h-2v2h2zm-6 4h2v2H7zm10 0h-2v2h2zm-6 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2H2v10h2v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v-2h2v-2h-2v-2h-2V8h-2V6h-2V4h-2zm0 2v2h2v2h2v2h2v2h2v2h-2v2h-2v2h-2v2h-2v-2h-2v-2H8v-2H6v-2H4V4zM6 6h2v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5H2v14h14v-2h2v-2h2v-2h2v-2h-2V9h-2V7h-2zm0 2v2h2v2h2v2h-2v2h-2v2H4V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelAltMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5H6v10h12v-2h2v-2h2V9h-2V7h-2V5zm10 2v2h2v2h-2v2H8V7zM4 9H2v10h12v-2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelSharp(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5H2v4h2v2h2v2H4v2H2v4h14v-2h2v-2h2v-2h2v-2h-2V9h-2V7h-2zm0 2v2h2v2h2v2h-2v2h-2v2H4v-2h2v-2h2v-2H6V9H4V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layout(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h20v14H2zm2 2v4h16V7zm16 6H10v4h10zM8 17v-4H4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutAlignBottom(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4H8v12h8zm-6 10V6h4v8zm10 6v-2H4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutAlignLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 16V8H8v8zm-10-6h8v4h-8zM4 20h2V4H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutAlignRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 8v8h12V8zm10 6H6v-4h8zm6-10h-2v16h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutAlignTop(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 20H8V8h8zm-6-10v8h4v-8zm10-6v2H4V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutColumns(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h20v14H2zm2 2v10h7V7zm9 0v10h7V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutDistributeHorizontal(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4H4v16h2zm14 0h-2v16h2zM10 7h6v10H8V7zm4 8V9h-4v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutDistributeVertical(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6V4H4v2zm0 14v-2H4v2zM17 8v8h-2V8zm-8 6v-4h6V8H7v8h8v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutFooter(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h20v14H2zm2 2v6h16V7zm16 8H4v2h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutHeader(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 19h20V5H2zm2-2v-6h16v6zm16-8H4V7h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutRows(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h20v14H2zm2 2v4h16V7zm16 6H4v4h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSidebarLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h20v14H2zm2 2v10h2V7zm4 0v10h12V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSidebarRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 5H2v14h20zm-2 2v10h-2V7zm-4 0v10H4V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h7v2H4v8h7v2H2V6zm16 0h-7v2h7v8h-7v2h9V6zm-3 5H7v2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 6H4v2h2zm14 0H8v2h12zM4 11h2v2H4zm16 0H8v2h12zM4 16h2v2H4zm16 0H8v2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v18H2zm18 16V5H4v14zM8 7H6v2h2zm2 0h8v2h-8zm-2 4H6v2h2zm2 0h8v2h-8zm-2 4H6v2h2zm2 0h8v2h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loader(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2h-2v6h2zm0 14h-2v6h2zm9-5v2h-6v-2zM8 13v-2H2v2zm7-6h2v2h-2zm4-2h-2v2h2zM9 7H7v2h2zM5 5h2v2H5zm10 12h2v2h2v-2h-2v-2h-2zm-8 0v-2h2v2zv2H5v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2H9v2H7v4H4v14h16V8h-3V4h-2zm0 2v4H9V4zm-6 6h9v10H6V10zm4 3h-2v4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2H9v2H7v2h2V4h6v4H4v14h16V8h-3V4h-2zm0 8h3v10H6V10zm-2 3h-2v4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Login(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v4h2V5h14v14H5v-2H3v4h18V3zm12 8h-2V9h-2V7h-2v2h2v2H3v2h10v2h-2v2h2v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Logout(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3h16v4h-2V5H5v14h14v-2h2v4H3V3zm16 8h-2V9h-2V7h-2v2h2v2H7v2h10v2h-2v2h2v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Luggage(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2h6v4h4v14h-2v2h-2v-2H9v2H7v-2H5V6h4zm2 4h2V4h-2zM7 18h10V8H7zm4-8v6H9v-6zm4 0v6h-2v-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 4H2v16h20zM4 18V6h16v12zM8 8H6v2h2v2h2v2h4v-2h2v-2h2V8h-2v2h-2v2h-4v-2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailArrowRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H2v16h10v-2H4V6h16v6h2V4zM6 8h2v2H6zm4 4H8v-2h2zm4 0v2h-4v-2zm2-2v2h-2v-2zm0 0V8h2v2zm8 8h-2v-2h-2v-2h-2v2h2v2h-6v2h6v2h-2v2h2v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailCheck(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h18v10h-2V6H4v12h8v2H2V4zm4 4H6v2h2v2h2v2h4v-2h2v-2h2V8h-2v2h-2v2h-4v-2H8zm6 10h2v2h-2zm4 2v2h-2v-2zm2-2h-2v2h2zm0 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H2v16h12v-2H4V6h16v8h2V4zM6 8h2v2H6zm4 4H8v-2h2zm4 0v2h-4v-2zm2-2v2h-2v-2zm0 0V8h2v2zm2 6h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2h2v-2h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailFlash(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h18v8h-2V6H4v12h8v2H2V4zm4 4H6v2h2v2h2v2h4v-2h2v-2h2V8h-2v2h-2v2h-4v-2H8zm10 6h2v4h4v2h-2v2h-2v2h-2v-4h-4v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 2H4v16h20zM6 16V4h16v12zM2 7H0v15h19v-2H2zm8-1H8v2h2v2h2v2h4v-2h2V8h2V6h-2v2h-2v2h-4V8h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailOff(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h2v2H2zm4 4H4V4h2zm2 2H6V6h2zm2 2H8V8h2zm2 2h-2v-2h2zm2 0h-2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2zm2-2h-2v2h2zm0 0V8h2v2zm-6-6h12v12h-2V6H10zm4 14v2H2V8h2v10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailUnread(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2h-6v6h6zM4 4h10v2H4v12h16v-8h2v10H2V4zm4 4H6v2h2v2h2v2h4v-2h2v-2h-2v2h-4v-2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2h2v2h2v2h-2v10H8V6H6V4h2zM4 8V6h2v2zm2 10v2H4v2H2V8h2v10zm0 0h2v-2H6zm6 0h-2v-2h2zm2-10V6h-2v2zm2 0h-2v10h-2v2h2v2h2v-2h2v-2h2v-2h2V2h-2v2h-2v2h-2zm0 0h2V6h2v10h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Membercard(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v14h-7v3h-2v-2h-2v2H9v-3H2zm2 2v4h16V5zm16 8H4v2h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h16v2H4zm0 5h16v2H4zm16 5H4v2h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H2v20h2V4h16v12H6v2H4v2h2v-2h16V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageArrowLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h18v12h-2V4H4v18H2V2zm2 14h4v2H6v2H4v-2h2zm16 0h-6v-2h2v-2h-2v2h-2v2h-2v2h2v2h2v2h2v-2h-2v-2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageArrowRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h18v10h-2V4H4v18H2V2zm2 14h4v2H6v2H4v-2h2zm16 0h-2v-2h-2v-2h-2v2h2v2h-6v2h6v2h-2v2h2v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageBookmark(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2zm14 4h-6v8h2v-2h2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageClock(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H2v20h2V4h16v4h2V2zM8 16H6v2H4v2h2v-2h2zm6-2h2v2h2v2h-4zm6-4h-8v2h-2v8h2v2h8v-2h2v-8h-2zm0 2v8h-8v-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2zm9 7h-2V7H9v2h2v2H9v2h2v-2h2v2h2v-2h-2zm0 0V7h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageFlash(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H2v20h2V4h16v10h2V2zM10 16H6v2H4v2h2v-2h4zm6-4h2v4h4v2h-2v2h-2v2h-2v-4h-4v-2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageImage(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2zm10 4h-2v2h-2v2H8v2H6v2h2v-2h2v-2h2V8h2v2h2v2h2v-2h-2V8h-2zM6 6h2v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2zm12 7H8v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H2v20h2V4h16v12H6v2H4v2h2v-2h16V2zm-7 7h3v2h-3v3h-2v-3H8V9h3V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageProcessing(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h18v16H6v2H4v-2h2v-2h14V4H4v18H2V2zm5 7H7v2h2zm2 0h2v2h-2zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageReply(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h18v20h-2V4H4v12h14v2h2v2h-2v-2H2V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageText(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H2v20h2V4h16v12H6v2H4v2h2v-2h16V2zM6 7h12v2H6zm8 4H6v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 11h16v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MissedCall(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6h-4v2h2v2h-2v2h-2v2h-2v2h-2v-2H8v-2H6v-2H4V8H2v2h2v2h2v2h2v2h2v2h2v-2h2v-2h2v-2h2v-2h2v2h2V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modem(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2h-8v2H9v2h2V4h8v2h2V4h-2zm-8 6h2v2h-2zm6 0V6h-4v2zm0 0h2v2h-2zm-1 2h-2v2H2v10h20V12h-6zm4 4v6H4v-6zm-2 2h-2v2h2zm-6 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4H2v12h4v4h16V8h-4V4zm0 2v2H6v6H4V6zm-8 4h12v8H8zm8 2h-4v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitor(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H2v14h8v2H8v2h8v-2h-2v-2h8V3zm-6 12H4V5h16v10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodHappy(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3h14v2H5zm0 16H3V5h2zm14 0v2H5v-2zm0 0h2V5h-2zM10 8H8v2h2zm4 0h2v2h-2zm-5 6v-2H7v2zm6 0v2H9v-2zm0 0h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodNeutral(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3h14v2H5zm0 16H3V5h2zm14 0v2H5v-2zm0 0h2V5h-2zM10 8H8v2h2zm4 0h2v2h-2zm1 5H9v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodSad(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3h14v2H5zm0 16H3V5h2zm14 0v2H5v-2zm0 0h2V5h-2zM10 8H8v2h2zm4 0h2v2h-2zm-5 8v-2h6v2h2v-2h-2v-2H9v2H7v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h8v2h-2v2h-2V4H6zM4 6V4h2v2zm0 10H2V6h2zm2 2H4v-2h2zm2 2H6v-2h2zm10 0v2H8v-2zm2-2v2h-2v-2zm-2-4h2v4h2v-8h-2v2h-2zm-6 0v2h6v-2zm-2-2h2v2h-2zm0 0V6H8v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonStar(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h8v2h-2v2h-2V4H6zM4 6V4h2v2zm0 10H2V6h2zm2 2H4v-2h2zm2 2H6v-2h2zm10 0v2H8v-2zm2-2v2h-2v-2zm-2-4v-2h2v-2h2v8h-2v-4zm-6 0h6v2h-6zm-2-2h2v2h-2zm0 0V6H8v6zm8-10h2v2h2v2h-2v2h-2V6h-2V4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonStars(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 0h2v2h2v2h-2v2h-2V4h-2V2h2zM8 4h8v2h-2v2h-2V6H8zM6 8V6h2v2zm0 8H4V8h2zm2 2H6v-2h2zm8 0v2H8v-2zm2-2v2h-2v-2zm-2-4v-2h2V8h2v8h-2v-4zm-4 0h4v2h-4zm0 0V8h-2v4zm-8 6H2v2H0v2h2v2h2v-2h2v-2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreHorizontal(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 9h6v6H1zm2 2v2h2v-2zm6-2h6v6H9zm2 2v2h2v-2zm6-2h6v6h-6zm2 2v2h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVertical(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1v6H9V1zm-2 2h-2v2h2zm2 6v6H9V9zm-2 2h-2v2h2zm2 6v6H9v-6zm-2 2h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3h12v18H6zm2 2v4h3V5zm5 0v4h3V5zm3 6H8v8h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0h-2v2H9v2H7v2h2V4h2v7H4V9h2V7H4v2H2v2H0v2h2v2h2v2h2v-2H4v-2h7v7H9v-2H7v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2v-7h7v2h-2v2h2v-2h2v-2h2v-2h-2V9h-2V7h-2v2h2v2h-7V4h2v2h2V4h-2V2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Movie(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm2 2v2h2V5zm4 0v6h6V5zm8 0v2h2V5zm2 4h-2v2h2zm0 4h-2v2h2zm0 4h-2v2h2zm-4 2v-6H9v6zm-8 0v-2H5v2zm-2-4h2v-2H5zm0-4h2V9H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4h12v16h-8v-8h6V8h-8v12H2v-8h6zm0 10H4v4h4zm10 0h-4v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4h2v2h2v2h2v2h2v4h-2v2h-2v2H8v2H6zm12 0h-2v16h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Note(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h18v14h-2v2h-2v-2h-2v2h2v2h-2v2H3zm2 2v16h8v-6h6V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2h10v14h-2v2h-2v-2h-2v2h2v2h-2v2H3V10h2v10h8v-6h6V4h-8zM7 4H5V2H3v2h2v2H3v2h2V6h2v2h2V6H7zm0 0h2V2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6H7v16h8v-2h2v-2h-2v-2h2v2h2v-2h2zM9 20V8h10v6h-6v6zm-6-2h2V4h12V2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotePlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1H5v3H2v2h3v3h2V6h3V4H7zm12 1h-7v2h7v10h-6v6H5v-9H3v11h12v-2h2v-2h2v-2h2V2zm-2 16h-2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notes(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2h16v20H3V2zm14 18V4H5v16zM7 6h10v2H7zm10 4H7v2h10zM7 14h7v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotesDelete(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H3v20h10v-2H5V4h14v10h2V2zm-2 4H7v2h10zM7 10h10v2H7zm6 4H7v2h6zm6 4h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2zm0 0h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotesMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0h16v20H5V0zm14 18V2H7v16zM9 4h10v2H9zm10 4H9v2h10zM9 12h7v2H9zm10 10H3V4H1v20h18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotesPlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2h16v12h-2V4H5v16h8v2H3V2zm2 4h10v2H7zm10 4H7v2h10zM7 14h7v2H7zm13 5h3v2h-3v3h-2v-3h-3v-2h3v-3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notification(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4V2h-4v2H5v2h14V4zm5 12H5v-4H3v6h5v4h2v-4h4v2h-4v2h6v-4h5v-6h-2V6h-2v8h2zM5 6v8h2V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOff(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2v2h5v2h-8V2zM5 16h9v2h2v4h-6v-2h4v-2h-4v4H8v-4H3v-6h2v-2h2v4H5zm16-2h-2v-2h-2V6h2v6h2zM5 2H3v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h-2v-2h-2V8H9V6H7V4H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Open(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3h6v2H5v14h14v-6h2v8H3V3zm8 0h8v8h-2V7h-2V5h-4zm0 8h-2v2H9v2h2v-2h2zm4-4h-2v2h-2v2h2V9h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintBucket(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3h8v2H8zm0 2H6v4H4v12h16V9h-2V5h-2v4H8zm8 6h2v8H6v-8h2v6h2v-4h2v2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5h16v10H7V9h10v2H9v2h10V7H5v10h14v2H3V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4H5v16h5zm9 0h-5v16h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percent(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4h-2v2h-2v2h-2v2h-2v2h-2v2H8v2H6v2H4v2h2v-2h2v-2h2v-2h2v-2h2v-2h2V8h2V6h2zm-4 10h4v6h-6v-6zm2 4v-2h-2v2zM6 4h4v6H4V4zm2 4V6H6v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureInPicture(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h20v16H2zm2 2v12h16V6zm6 2h8v6h-8zm2 2v2h4v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureInPictureAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h20v16H2zm2 2v12h16V6zm6 4h8v6h-8zm2 2v2h4v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2h10v2H7zM5 6V4h2v2zm0 8H3V6h2zm2 2H5v-2h2zm2 2H7v-2h2zm2 2H9v-2h2zm2 0v2h-2v-2zm2-2v2h-2v-2zm2-2v2h-2v-2zm2-2v2h-2v-2zm0-8h2v8h-2zm0 0V4h-2v2zm-5 2h-4v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pixelarticons(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3v18h18V3zm16 2v14H5V5zM7 7h6v6H9v2H7zm8 6h-2v2h-2v2h2v-2h2v2h2v-2h-2zm0 0h2v-2h-2zM9 9v2h2V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 20H8V4h2v2h2v3h2v2h2v2h-2v2h-2v3h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playlist(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 13h6V5h6v4h-4v10h-8zm2 2v2h4v-2zM2 17h6v2H2zm6-4H2v2h6zM2 9h12v2H2zm12-4H2v2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4h2v7h7v2h-7v7h-2v-7H4v-2h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2h-2v4H6v2H4v8h2v2h2v4h8v-2h4v-2h-4v-2h4v-2h-4v-2H8v4H6V8h12V6h2zm-6 18h-4v-6h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Prev(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4h2v16H6zm12 0h-2v2h-2v3h-2v2h-2v2h2v3h2v2h2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h12v6h4v10h-4v4H6v-4H2V8h4zm2 6h8V4H8zm-2 8v-4h12v4h2v-6H4v6zm2-2v6h8v-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioHandheld(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2v5h8v15H7V2zm0 7v4h6V9zm6 6H9v5h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioOn(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3H7v2H5v2H3v10h2v2h2v2h10v-2h2v-2h2V7h-2V5h-2zm0 2v2h2v10h-2v2H7v-2H5V7h2V5zm-9 6h2v2h2v2h-2v-2H8zm8-2h-2v2h-2v2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioSignal(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2h2v2h-2zm2 14V4h2v12zm0 0v2h-2v-2zM1 4h2v12H1zm2 12h2v2H3zM3 4h2V2H3zm2 2h2v8H5zm2 8h2v2H7zm0-8h2V4H7zm10 0h2v8h-2zm0 0h-2V4h2zm0 8v2h-2v-2zm-6-7h4v6h-2v9h-2v-9H9V7zm0 4h2V9h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioTower(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2h-2v2h2v12h-2v2h2v-2h2V4h-2zM2 4H0v12h2v2h2v-2H2zm0 0V2h2v2zm4 2H4v8h2zm0 0V4h2v2zm4 0h4v2h-4zm0 6H8V8h2zm4 0h-4v2H8v4H6v4h2v-4h2v-4h4v4h2v4h2v-4h-2v-4h-2zm0 0h2V8h-2zm6-6h-2V4h-2v2h2v8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reciept(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h2v2h2v2H5v14h14V6h-2V4h2V2h2v20H3zm12 2V2h2v2zm-2 0h2v2h-2zm-2 0V2h2v2zM9 4h2v2H9zm0 0V2H7v2zm8 4H7v2h10zM7 12h10v2H7zm10 6v-2h-4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecieptAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2H3v20h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h2v2h2V2h-2v2h-2V2h-2v2h-2V2h-2v2H9V2H7v2H5zm2 2h2v2h2V4h2v2h2V4h2v2h2v12h-2v2h-2v-2h-2v2h-2v-2H9v2H7v-2H5V6h2zm0 4h10v2H7zm10 4H7v2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4h-2v2h2v2H6v2H4v8h2v2h6v-2H6v-8h10v2h-2v2h2v-2h2v-2h2V8h-2V6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reload(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2h-2v2h2v2H4v2H2v5h2V8h12v2h-2v2h2v-2h2V8h2V6h-2V4h-2zM6 20h2v2h2v-2H8v-2h12v-2h2v-5h-2v5H8v-2h2v-2H8v2H6v2H4v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v18h18V3zm14 2v14H5V5zm-3 6H8v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveBoxMultiple(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v14h14V3zm10 2v10H5V5zm4 2v12H7v2h14V7zm-6 2H7v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1H9v2h2v2H5v2H3v10h2v2h2v-2H5V7h6v2H9v2h2V9h2V7h2V5h-2V3h-2zm8 4h-2v2h2v10h-6v-2h2v-2h-2v2h-2v2H9v2h2v2h2v2h2v-2h-2v-2h6v-2h2V7h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19h-2v-2H8v-2H6v-2H4v-2h2V9h2V7h2V5h2v4h8v6h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplyAll(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 19h2v-4h7V9h-7V5h-2v2h-2v2H9v2H7v2h2v2h2v2h2zM8 7H6v2H4v2H2v2h2v2h2v2h2v2h2v-2H8v-2H6v-2H4v-2h2V9h2zm0 0h2V5H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundedCorner(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h2v2H3zm0 4h2v2H3zm2 4H3v2h2zm-2 4h2v2H3zm2 4H3v2h2zm2 0h2v2H7zm6 0h-2v2h2zm2 0h2v2h-2zm6 0h-2v2h2zm-2-4h2v2h-2zM17 5h-2V3h-4v2h4v2h2v2h2v4h2V9h-2V7h-2zM7 3h2v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h14v2H4v16h2v-6h12v6h2V6h2v16H2V2zm4 18h8v-4H8zM20 6h-2V4h2zM6 6h9v4H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scale(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 3h-8v2h4v2h2v4h2zm-4 4h-2v2h-2v2h2V9h2zm-8 8h2v-2H9zH7v2h2zm-4-2v4h2v2H5h6v2H3v-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Script(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3h14v2h2v6h-2v8h-2V5H6zm8 14v-2H6V5H4v10H2v4h2v2h14v-2h-2v-2zm0 0v2H4v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScriptText(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3h14v2h2v6h-2v8h-2V5H6zm8 14v-2H6V5H4v10H2v4h2v2h14v-2h-2v-2zm0 0v2H4v-2zM8 7h8v2H8zm8 4H8v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScrollHorizontal(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2v2H2V2zm0 18v2H2v-2zm-6-5v-2H8v2H6v-2H4v-2h2V9h2v2h8V9h2v2h2v2h-2v2zm0 0v2h-2v-2zm0-6h-2V7h2zM8 9V7h2v2zm0 6h2v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScrollVertical(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h2v20H2zm9 18h2v-2h2v-2h2v-2h-2v2h-2V8h2v2h2V8h-2V6h-2V4h-2v2H9v2H7v2h2V8h2v8H9v-2H7v2h2v2h2zM22 2h-2v20h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sd(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2h2v20H4V6h2v14h12V4H8V2zM8 4H6v2h2zm6 2h2v4h-2zm-2 0h-2v4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2h8v2H6zM4 6V4h2v2zm0 8H2V6h2zm2 2H4v-2h2zm8 0v2H6v-2zm2-2h-2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2zm0-8h2v8h-2zm0 0V4h-2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Section(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v2h2zm4 0H7v2h2zM7 19h2v2H7zM5 7H3v2h2zm14 0h2v2h-2zM5 11H3v2h2zm14 0h2v2h-2zM5 15H3v2h2zm14 0h2v2h-2zM5 19H3v2h2zm6-16h2v2h-2zm2 16h-2v2h2zm2-16h2v2h-2zm2 16h-2v2h2zm2-16h2v2h-2zm2 16h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SectionCopy(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v2h2zm2 4h2v2H7zm4 0h2v2h-2zm2 12h-2v2h2zm2 0h2v2h-2zm6 0h-2v2h2zM7 11h2v2H7zm14 0h-2v2h2zm-2 4h2v2h-2zM7 19h2v2H7zM19 7h2v2h-2zM7 3h2v2H7zm2 12H7v2h2zM3 7h2v2H3zm14 0h-2v2h2zM3 11h2v2H3zm2 4H3v2h2zm6-12h2v2h-2zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SectionMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v2h2zm4 0H7v2h2zM7 19h2v2H7zm6 0h-2v2h2zM3 7h2v2H3zm18 0h-2v2h2zm-2 4h2v2h-2zM5 11H3v2h2zm-2 4h2v2H3zm2 4H3v2h2zm6-16h2v2h-2zm6 0h-2v2h2zm2 0h2v2h-2zm2 14h-6v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SectionPlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h2v2H3zm4 0h2v2H7zm2 16H7v2h2zm2 0h2v2h-2zM5 7H3v2h2zm14 0h2v2h-2zm2 4h-2v2h2zM3 11h2v2H3zm2 4H3v2h2zm12 0h2v2h2v2h-2v2h-2v-2h-2v-2h2zM5 19H3v2h2zm6-16h2v2h-2zm6 0h-2v2h2zm4 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SectionX(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3H3v2h2zm4 0H7v2h2zM7 19h2v2H7zm6 0h-2v2h2zM3 7h2v2H3zm18 0h-2v2h2zm-2 4h2v2h-2zm2 8h-2v-2h2v-2h-2v2h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2zM3 11h2v2H3zm2 4H3v2h2zm-2 4h2v2H3zM13 3h-2v2h2zm2 0h2v2h-2zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h18v18H3zm2 2v6h14V5zm14 8H5v6h14zM7 7h2v2H7zm2 8H7v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SharpCorner(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h2v2H3zm0 4h2v2H3zm2 4H3v2h2zm-2 4h2v2H3zm2 4H3v2h2zm2 0h2v2H7zm6 0h-2v2h2zm2 0h2v2h-2zm6 0h-2v2h2zm-2-4h2v2h-2zm2-2V3H11v2h8v8zM7 3h2v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2H2v12h2V4h16v10h2zM6 14H4v2h2zm0 2h2v2h2v2H8v-2H6zm4 4v2h4v-2h2v-2h-2v2zm10-6h-2v2h-2v2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldOff(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2h14v12h-2V4H8zM2 8h2v6H2zm2 6h2v2H4zm4 2H6v2h2v2h2v2h4v-2h-4v-2H8zm10 0h-2v2h2v2h2v2h2v-2h-2v-2h-2zM4 2H2v2h2v2h2v2h2v2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2V8H8V6H6V4H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ship(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4v2h4v2H6v2h6V8h2v2h8v6h-2v-4H4v6h14v-2h2v2h4v2H0v-2h2v-8h2V6h2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2h6v2H9zm6 4V4h2v2h4v16H3V6h4V4h2v2zm0 2H9v2H7V8H5v12h14V8h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5h-2v2h2v2h-6v2h-2v6H2v2h8v-2h2v-6h6v2h-2v2h2v-2h2v-2h2V9h-2V7h-2zM2 9h6v2H2zm20 10v-2h-8v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sliders(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4h2v10h-2zm0 12h-2v2h2v2h2v-2h2v-2zm-4-6h-2v10h2zm-8 2H3v2h2v6h2v-6h2v-2zm8-8h-2v2H9v2h6V6h-2zM5 4h2v6H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlidersTwo(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-width="2" d="M3 8h4m0 0V6h4v2M7 8v2h4V8m0 0h10M3 16h10m0 0v-2h4v2m-4 0v2h4v-2m0 0h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 20H6V8H4V6h2V4h2v2h2v2H8zm2-12v2h2V8zM4 8v2H2V8zm14-4h-2v12h-2v-2h-2v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlpabetic(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2h2v2h-2zm0 2v2H9V4zm2 0h2v2h-2zM9 18v2h2v2h2v-2h2v-2h-2v2h-2v-2zM8 8H2v8h2v-2h2v2h2zm-2 4H4v-2h2zm6-1v-1h2v1zm4-3h-6v8h6zm-4 6v-1h2v1zm10-6h-4v8h4v-2h-2v-4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumeric(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2h-2v2H9v2h2V4h2v2h2V4h-2zM2 8h4v8H4v-6H2zm6 0h6v5h-4v1h4v2H8v-5h4v-1H8zm12 0h-4v2h4v1h-4v2h4v1h-4v2h6V8zm-9 10v2H9v-2zm2 2h-2v2h2zm0 0v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2H3v20h18V2zm15 2v16H5V4zm-6 2h-2v2h2zm-5 4h8v6h-2v-4h-4v4H8zm8 6H8v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeedFast(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 5H9v2H5v2H3v2H1v6h2v2h2v-2H3v-6h2V9h4V7h6zm8 6h-2v6h-2v2h2v-2h2zm-13 2h4v4h-4zm6-2h-2v2h2zm2-2v2h-2V9zm0 0V7h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeedMedium(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 5h-2v8h-1v4h4v-4h-1zM9 7H5v2H3v2H1v6h2v2h2v-2H3v-6h2V9h4zm12 4h2v6h-2zm-2-2h2v2h-2zm0 0h-4V7h4zm2 8v2h-2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpeedSlow(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 5h6v2H9zm10 4h-4V7h4zm2 2h-2V9h2zm0 6v-6h2v6zm0 0v2h-2v-2zM1 11h2v6H1zm2 6h2v2H3zm11-4h-4v-2H8V9H6V7H4v2h2v2h2v2h2v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotlight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2h16v20H3V2zm14 18V4H5v16zM13 6H7v2h6zm-6 4h10v8H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Store(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3h16v2H4zm0 4h18v8h-2v6h-2v-6h-4v6H4v-6H2V7zm8 12v-4H6v4zm0-6h8V9H4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subscriptions(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H6v2h12zM4 6h16v2H4zm-2 4h20v12H2zm18 10v-8H4v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subtitles(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7h-8v10h8v-2h-6V9h6zM3 15V7h8v2H5v6h6v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3h8v4h6v14H2V7h6zm2 4h4V5h-4zM4 9v10h16V9zm4 2v6H6v-6zm10 0v6h-2v-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3h-2v2h2zm4 2h2v2h-2zm-6 6h2v2h-2zm-8 0h2v2H3zm18 0h-2v2h2zM5 5h2v2H5zm14 14h-2v-2h2zm-8 2h2v-2h-2zm-4-2H5v-2h2zM9 7h6v2H9zm0 8H7V9h2zm0 0v2h6v-2h2V9h-2v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0h-2v4h2zM0 11v2h4v-2zm24 0v2h-4v-2zM13 24h-2v-4h2zM8 6h8v2H8zM6 8h2v8H6zm2 10v-2h8v2zm10-2h-2V8h2zm2-14h2v2h-2zm0 2v2h-2V4zm2 18h-2v-2h2zm-2-2h-2v-2h2zM4 2H2v2h2v2h2V4H4zM2 22h2v-2h2v-2H4v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Switch(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5V3h2v2zm4 2H5V5h2zm2 2H7V7h2zm2 2H9V9h2zm2 0h-2v2h2v2h2v2h2v2h-2v2h6v-6h-2v2h-2v-2h-2v-2h-2zm2-2v2h-2V9zm2-2v2h-2V7zm0-2v2h2v2h2V3h-6v2zM5 19v-2h2v2zm0 0v2H3v-2zm2-2v-2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sync(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 9V7h12V5h2v2h2v2h-2v2h-2V9zm12 2h-2v2h2zm0-6h-2V3h2zm4 12v-2H8v-2h2v-2H8v2H6v2H4v2h2v2h2v2h2v-2H8v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tab(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v18H2zm2 2v14h16V9h-8V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v18H2zm2 4v5h7V7zm9 0v5h7V7zm7 7h-7v5h7zm-9 5v-5H4v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tea(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h18v7h-4v5H4zm14 5h2V6h-2zm-2-3h-4v2h2v4H8V8h2V6H6v8h10zm3 12v2H3v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Teach(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2H5v4h4zm7 7V7H2v9h2v6h2v-6h2v6h2V9zm-5-7h11v14H11v-2h9V4h-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAdd(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4H3v2h16zm0 4H3v2h16zM3 12h8v2H3zm8 4H3v2h8zm7-1h3v2h-3v3h-2v-3h-3v-2h3v-3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextColums(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 5H3v2h8zm10 0h-8v2h8zM3 9h8v2H3zm18 0h-8v2h8zM3 13h8v2H3zm18 0h-8v2h8zM3 17h8v2H3zm18 0h-8v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSearch(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4v2h16zm0 4H4v2h16zm-8 4H4v2h8zm8 0h-6v6h6v2h2v-2h-2zm-4 4v-2h2v2zm-4 0H4v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextWrap(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 5H3v2h16v6h-6v-2h2V9h-2v2h-2v2H9v2h2v2h2v2h2v-2h-2v-2h6v-2h2V7h-2zM7 13H3v2h4zM3 9h6v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeline(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7h4v4H7zm-2 6v-2h2v2zm0 0v4H1v-4zm8 0h-2v-2h2zm4 0h-4v4h4zm2-2v2h-2v-2zm0 0h4V7h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleLeft(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5h16v2H4zm0 12H2V7h2zm16 0v2H4v-2zm0 0h2V7h-2zM10 9H6v6h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleRight(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5h16v2H4zm0 12H2V7h2zm16 0v2H4v-2zm0 0h2V7h-2zm-2-8h-4v6h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tournament(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2H2v2h5v4H2v2h7V7h5v10H9v-3H2v2h5v4H2v2h7v-3h7v-6h6v-2h-6V5H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrackChanges(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2H2v20h20V4h-2v16H4V4h7v2H6v12h12V8h-2v8H8V8h3v2h-1v4h4v-4h-1V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2v4h6v2h-2v14H4V8H2V6h6V2zm-2 2h-4v2h4zm0 4H6v12h12V8zm-5 2h2v8H9zm6 0h-2v8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashAlt(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2v4h6v2h-2v14H4V8H2V6h6V2zm-2 2h-4v2h4zm0 4H6v12h12V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trending(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4h2v14h16v2H3zm6 10H7v2h2zm2-2v2H9v-2zm2 0v-2h-2v2zm2 0h-2v2h2zm2-2h-2v2h2zm2-2v2h-2V8zm0 0V6h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingDown(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8h2v2h2v2h2v2h2v-2h2v-2h2v2h2v2h2v2h-4v2h8v-8h-2v4h-2v-2h-2v-2h-2V8h-2v2h-2v2H8v-2H6V8H4V6H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingUp(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6h8v8h-2v-4h-2V8h-4zm2 6v-2h2v2zm-2 2v-2h2v2zm-2 0h2v2h-2zm-2-2h2v2h-2zm-2 0v-2h2v2zm-2 2v-2h2v2zm-2 2v-2h2v2zm0 0v2H2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3H6v2H2v10h6V5h8v10h6V5h-4V3zm4 4v6h-2V7zM6 13H4V7h2zm12 2H6v2h12zm-7 2h2v2h3v2H8v-2h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h14v4h4v2h-4v6h6v-4h2v6h-4v2h-4v-2H8v2H4v-2H0V4zm20 8h-2v-2h2zm-8-2V6H2v10h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4h2v2H8zm10 6V8H8V6H6v2H4v2h2v2h2v2h2v-2H8v-2zm0 8v-8h2v8zm0 0v2h-6v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ungroup(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 3H3v4h4zm0 14H3v4h4zM17 3h4v4h-4zm4 14h-4v4h4zM8 8h2v2H8zm4 2h-2v4H8v2h2v-2h4v2h2v-2h-2v-4h2V8h-2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlink(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4h-2v16h2zM4 6h5v2H4v8h5v2H2V6zm11 0h7v12h-7v-2h5V8h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 5V3h2v2h2v2h2v2h-2V7h-2v10h-2V7H9v2H7V7h2V5zM3 15v6h18v-6h-2v4H5v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2H9v2H7v6h2V4h6zm0 8H9v2h6zm0-6h2v6h-2zM4 16h2v-2h12v2H6v4h12v-4h2v6H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h6v2h-6v6h-2V4h2zm0 8h6v2h-6zm8-6h-2v6h2zM9 16H7v6h16v-6h-2v4H9zh12v-2H9zm-2-6H1v2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2h-6v2h-2v6h2V4h6zm0 8h-6v2h6zm0-6h2v6h-2zM7 16h2v-2h12v2H9v4h12v-4h2v6H7zM3 8h2v2h2v2H5v2H3v-2H1v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserX(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h6v2h-6v6h-2V4h2zm0 8h6v2h-6zm8-6h-2v6h2zM7 16v6h16v-6h-2v4H9v-4h12v-2H9v2zm-1-6H4V8H2v2h2v2H2v2h2v-2h2v2h2v-2H6zm0 0h2V8H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 0H5v2H3v6h2v2h6V8H5V2h6zm0 2h2v6h-2zM0 14h2v4h12v2H0zm2 0h12v-2H2zm14 0h-2v6h2zM15 0h4v2h-4zm4 8h-4v2h4zm0-6h2v6h-2zm5 12h-2v4h-4v2h6zm-6-2h4v2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h14v4h2V7h2V5h2v14h-2v-2h-2v-2h-2v4H2zm2 12h10V7H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoOff(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5H2v14h14v-4h2v2h2v2h2V5h-2v2h-2v2h-2V5zm10 12H4V7h10zm-4-6H8V9H6v2h2v2H6v2h2v-2h2v2h2v-2h-2zm0 0V9h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewCol(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h20v14H2zm2 2v10h4V7zm6 0v10h4V7zm6 0v10h4V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewList(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5h20v14H2zm2 2v2h16V7zm16 4H4v2h16zm0 4H4v2h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewportNarrow(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2H8v4h2V4h4v2h2V2zM8 20v-2h2v2h4v-2h2v4H8zm9-9h5v2h-5v2h-2v-2h-2v-2h2V9h2zm0-2V7h2v2zm0 6h2v2h-2zM2 11h5V9h2v2h2v2H9v2H7v-2H2zm5 4v2H5v-2zm0-6V7H5v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewportWide(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2H2v4h2V4h16v2h2V2zM2 20v-2h2v2h16v-2h2v4H2zm16-9h-5v2h5v2h-2v2h2v-2h2v-2h2v-2h-2V9h-2V7h-2v2h2zm-7 0H6V9h2V7H6v2H4v2H2v2h2v2h2v2h2v-2H6v-2h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Visible(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6h8v2H8zm-4 4V8h4v2zm-2 2v-2h2v2zm0 2v-2H0v2zm2 2H2v-2h2zm4 2H4v-2h4zm8 0v2H8v-2zm4-2v2h-4v-2zm2-2v2h-2v-2zm0-2h2v2h-2zm-2-2h2v2h-2zm0 0V8h-4v2zm-10 1h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2h2v20h-2v-2h-2v-2h2V6h-2V4h2zm-4 6V6h2v2zm-2 2h2V8H7v8h4v2h2v-2h-2v-2H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMinus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-2v2H8v2H6v2H2v8h4v2h2v2h2v2h2zM8 18v-2H6v-2H4v-4h2V8h2V6h2v12zm14-7h-8v2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOne(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2h-2v2h-2v2H9v2H5v8h4v2h2v2h2v2h2zm-4 16v-2H9v-2H7v-4h2V8h2V6h2v12zm6-8h2v4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumePlus(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2h2v20h-2v-2H8v-2h2V6H8V4h2zM6 8V6h2v2zm0 8H2V8h4v2H4v4h2zm0 0v2h2v-2zm13-5h3v2h-3v3h-2v-3h-3v-2h3V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeThree(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2H9v2H7v2H5v2H1v8h4v2h2v2h2v2h2zM7 18v-2H5v-2H3v-4h2V8h2V6h2v12zm6-8h2v4h-2zm8-6h-2V2h-6v2h6v2h2v12h-2v2h-6v2h6v-2h2v-2h2V6h-2zm-2 4h-2V6h-4v2h4v8h-4v2h4v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeTwo(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2h2v20h-2v-2H9v-2h2V6H9V4h2zM7 8V6h2v2zm0 8H3V8h4v2H5v4h2zm0 0v2h2v-2zm10-6h-2v4h2zm2-2h2v8h-2zm0 8v2h-4v-2zm0-10v2h-4V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeVibrate(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2h-2v2h-2v2H8v2H4v8h4v2h2v2h2v2h2zm-4 16v-2H8v-2H6v-4h2V8h2V6h2v12zm8-15h-2v2h2v2h-2v2h2v2h-2v2h2v2h-2v2h2v2h-2v2h2v-2h2v-2h-2v-2h2v-2h-2v-2h2V9h-2V7h2V5h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeX(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2h-2v2H9v2H7v2H3v8h4v2h2v2h2v2h2zM9 18v-2H7v-2H5v-4h2V8h2V6h2v12zm10-6.777h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2zm0 0h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3H2v18h18v-4h2V7h-2V3zm0 14v2H4V5h14v2h-8v10zm2-2h-8V9h8zm-4-4h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningBox(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h16v2H5v14h14v2H3zm18 0h-2v18h2zM11 15h2v2h-2zm2-8h-2v6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wind(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3H8v2h4v2H2v2h12V3zm10 8V7h-6v2h4v2H2v2h20zM2 17v-2h14v6h-6v-2h4v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zap(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1h2v8h8v4h-2v-2h-8V5h-2V3h2zM8 7V5h2v2zM6 9V7h2v2zm-2 2V9h2v2zm10 8v2h-2v2h-2v-8H2v-4h2v2h8v6zm2-2v2h-2v-2zm2-2v2h-2v-2zm0 0h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H6v2H4v2H2v8h2v2h2v2h8v-2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h2V6h-2V4h-2zm0 2v2h2v8h-2v2H6v-2H4V6h2V4zM9 6h2v3h3v2h-3v3H9v-3H6V9h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *PixelarticonsIcon {
	return &PixelarticonsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H6v2H4v2H2v8h2v2h2v2h8v-2h2v2h2v2h2v2h2v-2h-2v-2h-2v-2h-2v-2h2V6h-2V4h-2zm0 2v2h2v8h-2v2H6v-2H4V6h2V4zm0 5v2H6V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
