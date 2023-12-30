package ci

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type CiIcon struct {
	*SVGSVGElement
}

func AddColumn(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h3m0 0h3m-3 0v-3m0 3v3m6 1h1a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddMinusSquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h8M4 16.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C20 5.52 20 6.08 20 7.2v9.6c0 1.12 0 1.68-.218 2.108a2.001 2.001 0 0 1-.874.874c-.428.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C4 18.48 4 17.92 4 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddPlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12h6m0 0h6m-6 0v6m0-6V6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddPlusCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h4m0 0h4m-4 0v4m0-4V8m0 13a9 9 0 1 1 0-18a9 9 0 0 1 0 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddPlusSquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h4m0 0h4m-4 0v4m0-4V8m-8 8.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C20 5.52 20 6.08 20 7.2v9.6c0 1.12 0 1.68-.218 2.108a2.001 2.001 0 0 1-.874.874c-.428.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C4 18.48 4 17.92 4 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddRow(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 14v1a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-6m-3-3H7m0 0H4m3 0V5m0 3v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddToQueue(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9v10.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C3.76 21 4.04 21 4.598 21H15m-1-8v-3m0 0V7m0 3h-3m3 0h3M7 13.8V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C8.52 3 9.08 3 10.2 3h7.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C21 4.52 21 5.08 21 6.2v7.6c0 1.12 0 1.68-.218 2.108a2.001 2.001 0 0 1-.874.874c-.428.218-.986.218-2.104.218h-7.607c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C7 15.48 7 14.92 7 13.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeXd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4ZM6 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H6Zm9.762 12.776a3.188 3.188 0 0 1-2.351-.9a3.762 3.762 0 0 1-.928-2.69a3.8 3.8 0 0 1 .955-2.6a3.422 3.422 0 0 1 2.587-1.069h.233v-2c0-.156.127-.283.283-.284h1.317a.29.29 0 0 1 .2.07a.263.263 0 0 1 .082.2v7.679c-.005.265.009.53.04.793V16a.286.286 0 0 1-.056.177a.34.34 0 0 1-.121.094a5.188 5.188 0 0 1-2.222.506l-.019-.001Zm-.916-5.116l-.005.006a2.222 2.222 0 0 0-.45 1.455a2.264 2.264 0 0 0 .443 1.524c.267.294.652.454 1.049.434a1.63 1.63 0 0 0 .375-.038v-3.849a1.31 1.31 0 0 0-.284-.029a1.436 1.436 0 0 0-1.128.5v-.003Zm-2.621 5h-1.451a.358.358 0 0 1-.354-.22l-.21-.44l-.005-.011L10 15.58l-.064-.133c-.319-.658-.648-1.336-.978-2.034a125.419 125.419 0 0 1-1.417 3.031v.01l-.005.009a.4.4 0 0 1-.136.137a.354.354 0 0 1-.188.05H5.841a.264.264 0 0 1-.263-.179a.282.282 0 0 1 .043-.245l2.229-4.261l-2.156-4.314a.27.27 0 0 1-.012-.288a.273.273 0 0 1 .234-.131H7.35a.372.372 0 0 1 .2.046a.33.33 0 0 1 .136.163c.536 1.125.988 2.1 1.383 2.965c.491-1.1 1.013-2.218 1.363-2.959l.005-.011l.006-.01v-.011a.495.495 0 0 1 .093-.115a.332.332 0 0 1 .221-.077H12.1a.266.266 0 0 1 .247.151a.285.285 0 0 1-.023.263l-2.157 4.187l2.282 4.4a.3.3 0 0 1 .029.259a.266.266 0 0 1-.253.175v.002Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplay(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.25 17h-.053c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C3 15.48 3 14.92 3 13.8V8.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v5.606c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.874.875C19.48 17 18.92 17 17.803 17h-.05M16 20H8l4-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7v5h5m4.004-7.429L17.939 2M6.064 2L3 4.571M12 20a8 8 0 1 1 0-16a8 8 0 0 1 0 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a9.121 9.121 0 0 1-9-9a9.121 9.121 0 0 1 9-9a9.121 9.121 0 0 1 9 9a9.121 9.121 0 0 1-9 9Zm0-16a7.094 7.094 0 0 0-7 7a7.094 7.094 0 0 0 7 7a7.094 7.094 0 0 0 7-7a7.094 7.094 0 0 0-7-7Zm1 12h-2v-4H7v-2h4V8h2v4h4v2h-4v4Zm7.292-11.292l-3.009-3l1.409-1.417l3.01 3l-1.41 1.416v.001Zm-16.583 0L2.292 5.291l2.991-3l1.415 1.416l-2.989 3v.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppStore(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-4.273-7.213c-.129 0-.257.016-.383.046l-.15.041l-.594 1.017a.823.823 0 0 0 1.365.916l.062-.093l.79-1.371a1.386 1.386 0 0 0-1.1-.556h.01Zm5.484-6.127a1.944 1.944 0 0 0-.587 1.1a2.4 2.4 0 0 0 .3 1.689l3.04 5.266a.827.827 0 0 0 1.071.322a.823.823 0 0 0 .4-1.044l-.049-.1l-.8-1.391h1.19a.823.823 0 0 0 .1-1.639l-.1-.006h-2.14l-2.194-3.8l-.229-.4l-.002.003ZM6.26 12.85a.817.817 0 0 0-.815.821c0 .414.308.764.719.817l.1.005h7.48a1.16 1.16 0 0 0-.017-1.067a1.034 1.034 0 0 0-.793-.571l-.121-.006h-2.551L13.79 6.74a.828.828 0 0 0 .082-.625a.824.824 0 0 0-1.447-.29l-.062.092l-.363.634l-.359-.634a.827.827 0 0 0-1.072-.322a.825.825 0 0 0-.4 1.045l.049.1l.83 1.46l-2.685 4.65H6.26Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.094 21.006a2.457 2.457 0 0 1-1.653-.766a9.899 9.899 0 0 1-1.377-1.648a11.406 11.406 0 0 1-1.459-2.893a10.653 10.653 0 0 1-.61-3.445a6.317 6.317 0 0 1 .825-3.292a4.88 4.88 0 0 1 1.73-1.751a4.663 4.663 0 0 1 2.344-.662a5.51 5.51 0 0 1 1.81.421c.459.202.941.344 1.436.422a8.564 8.564 0 0 0 1.593-.5a5.79 5.79 0 0 1 1.903-.393c.087 0 .173 0 .256.01c1.43.041 2.76.743 3.6 1.9a4.02 4.02 0 0 0-2.123 3.637a4.022 4.022 0 0 0 1.317 3.023a4.38 4.38 0 0 0 1.316.863c-.1.3-.215.59-.337.882a10.38 10.38 0 0 1-1.02 1.837a9.437 9.437 0 0 1-1.317 1.592a2.57 2.57 0 0 1-1.692.743a4.249 4.249 0 0 1-1.562-.373a4.501 4.501 0 0 0-1.68-.373a4.647 4.647 0 0 0-1.73.371a4.656 4.656 0 0 1-1.495.393l-.075.002Zm3.15-14.507c-.075 0-.15 0-.225-.009a2.893 2.893 0 0 1-.025-.359a4.17 4.17 0 0 1 1.1-2.688a4.277 4.277 0 0 1 1.35-1.009a4.03 4.03 0 0 1 1.535-.435c.015.131.015.259.015.381a4.071 4.071 0 0 1-1.04 2.66a3.643 3.643 0 0 1-2.71 1.459Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.6 9h10.8M6.6 9c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437c-.11.214-.11.494-.11 1.054v5.2c0 1.12 0 1.68.219 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.117 0 1.676 0 2.104-.218c.376-.192.683-.498.875-.874c.218-.427.218-.986.218-2.104V10.59c0-.554 0-.832-.109-1.045a1.002 1.002 0 0 0-.437-.437C18.24 9 17.96 9 17.4 9M6.6 9H4.975c-.849 0-1.273 0-1.514-.148a1 1 0 0 1-.472-.938c.024-.282.277-.623.783-1.307c.147-.197.22-.296.31-.372c.119-.1.26-.172.413-.208C4.609 6 4.73 6 4.977 6h14.045c.246 0 .368 0 .482.027c.152.036.294.108.414.208c.09.076.163.174.31.372c.506.683.759 1.025.783 1.307a1 1 0 0 1-.472.938C20.298 9 19.872 9 19.023 9H17.4M10 14h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 13l3 3m0 0l3-3m-3 3V8m9 4a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDownLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 11v4m0 0h4m-4 0l6-6m6 3a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDownRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 15h4m0 0v-4m0 4L9 9m12 3a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 9l-3 3m0 0l3 3m-3-3h8m5 0a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 15l3-3m0 0l-3-3m3 3H8m13 0a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l-3-3m0 0l-3 3m3-3v8m9-4a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUpLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 9H9m0 0v4m0-4l6 6m6-3a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUpRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13V9m0 0h-4m4 0l-6 6m12-3a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.636 18.364h7.07m-7.07 0v-7.071m0 7.07L18.364 5.637"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7L7 17m0 0h8m-8 0V9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 8l-8 8m0 0h6m-6 0v-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 21l5-5m-5 5l-5-5m5 5V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v14m0 0l6-6m-6 6l-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364v-7.071m0 7.071h-7.071m7.07 0L5.637 5.636"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10m0 0V9m0 8H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 8l8 8m0 0v-6m0 6h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7v10m0 0l4-4m-4 4l-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 16l-3 3m0 0l-3-3m3 3V5m5 3l3-3m0 0l3 3m-3-3v14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 12l5 5m-5-5l5-5m-5 5h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 12H5m0 0l6 6m-6-6l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 13l3 3m0 0l-3 3m3-3H5m3-5L5 8m0 0l3-3M5 8h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12H7m0 0l4 4m-4-4l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowReloadTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 16h5v5M10 8H5V3m14.418 6.003A8 8 0 0 0 5.086 7.976m-.504 7.021a8 8 0 0 0 14.331 1.027"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 12l-5-5m5 5l-5 5m5-5H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m0 0l-6-6m6 6l-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12h10m0 0l-4-4m4 4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSubDownLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 11l-5 5m0 0l5 5m-5-5h7.803c1.118 0 1.677 0 2.105-.218a2 2 0 0 0 .874-.874c.218-.428.218-.987.218-2.105V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSubDownRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 11l5 5m0 0l-5 5m5-5h-7.803c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C7 14.48 7 13.92 7 12.8V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSubLeftDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.5 12.5l-5 5m0 0l-5-5m5 5V9.7c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874c.428-.218.988-.218 2.108-.218h9.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSubLeftUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12.5 11.5l-5-5m0 0l-5 5m5-5v7.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218H20.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSubRightDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 13l5 5m0 0l5-5m-5 5v-7.803c0-1.118 0-1.678-.218-2.105a2 2 0 0 0-.874-.874C14.48 7 13.92 7 12.8 7H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSubRightUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 11l5-5m0 0l5 5m-5-5v7.803c0 1.118 0 1.677-.218 2.105a2.002 2.002 0 0 1-.874.874C14.48 17 13.92 17 12.803 17H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSubUpLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 13L6 8m0 0l5-5M6 8h7.8c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105V21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSubUpRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 13l5-5m0 0l-5-5m5 5h-7.8c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C7 9.52 7 10.08 7 11.2V21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUndoDownLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 11l-4 4m0 0l4 4m-4-4h13a5 5 0 0 0 0-10h-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUndoDownRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 11l4 4m0 0l-4 4m4-4H8A5 5 0 0 1 8 5h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUndoUpLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 13L3 9m0 0l4-4M3 9h13a5 5 0 0 1 0 10h-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUndoUpRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 13l4-4m0 0l-4-4m4 4H8a5 5 0 0 0 0 10h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.636 5.636v7.07m0-7.07h7.07m-7.07 0l12.728 12.728"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17L7 7m0 0v8m0-8h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 16L8 8m0 0v6m0-6h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3L7 8m5-5l5 5m-5-5v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19V5m0 0l-6 6m6-6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.365 5.636h-7.071m7.07 0v7.071m0-7.07L5.638 18.363"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 17L17 7m0 0H9m8 0v8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 16l8-8m0 0h-6m6 0v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17V7m0 0l-4 4m4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsReloadOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 16H5v5m9-13h5V3M4.583 9.003a8 8 0 0 1 14.331-1.027m.504 7.021a8 8 0 0 1-14.332 1.027"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarBottom(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 15H4m16 0V7.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 5.52 4 6.08 4 7.2V15m16 0v1.803c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C4 18.48 4 17.92 4 16.8V15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChart(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 21h-3V11h3v10Zm-5 0h-3V8h3v13Zm-5 0H8V5h3v16Zm-5 0H3v-8h3v8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 21H2V11a2 2 0 0 1 2-2h4V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v3h4a2 2 0 0 1 2 2v12ZM16 9v10h4V9h-4Zm-6-5v15h4V4h-4Zm-6 7v8h4v-8H4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 20a8 8 0 1 1 0-16a8 8 0 0 1 0 16ZM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm13 4h2V9h-2v7Zm-8 0h2v-5H7v5Zm4 0h2V6h-2v10Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartHorizontal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 22H2V2h2v20Zm11-1H5v-3h10v3Zm3-5H5v-3h13v3Zm3-5H5V8h16v3Zm-8-5H5V3h8v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartSquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM5 5v14h14V5H5Zm12 12h-2v-7h2v7Zm-4 0h-2V7h2v10Zm-4 0H7v-5h2v5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20V4m0 16h7.803c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104V7.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4H9m0 16H7.197c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C4 18.48 4 17.92 4 16.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 4v16m0-16H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 5.52 4 6.08 4 7.2v9.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218H15m0-16h1.8c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v9.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarTop(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9h16M4 9v7.8c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104V9M4 9V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105V9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h3v16H2V4Zm4 0h1v16H6V4Zm5 0H9v16h2V4Zm1 0h2v16h-2V4Zm3 0h1v16h-1V4Zm5 0h-3v16h3V4Zm1 0h1v16h-1V4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.669 18.268a5.163 5.163 0 0 1-1.9-.337a4.024 4.024 0 0 1-1.439-.958a4.348 4.348 0 0 1-.905-1.485a5.461 5.461 0 0 1-.319-1.9a5.244 5.244 0 0 1 .328-1.864a4.358 4.358 0 0 1 .934-1.493a4.46 4.46 0 0 1 1.444-.994a4.6 4.6 0 0 1 1.858-.363a4.262 4.262 0 0 1 1.979.44a3.983 3.983 0 0 1 1.39 1.183c.369.507.635 1.082.783 1.691c.162.648.22 1.317.171 1.983h-6.431c-.026.633.191 1.25.606 1.729a2.186 2.186 0 0 0 1.584.538c.454.01.9-.12 1.278-.373c.3-.183.53-.46.653-.789h2.155a4.108 4.108 0 0 1-1.588 2.3a4.642 4.642 0 0 1-2.581.692Zm-.069-7.544a2.092 2.092 0 0 0-.974.2c-.238.12-.449.287-.62.49c-.15.185-.263.399-.328.628a2.812 2.812 0 0 0-.111.587h3.982a2.377 2.377 0 0 0-.563-1.409a1.834 1.834 0 0 0-1.386-.496Zm-9.63 7.312H2V5.731h5.8a7.645 7.645 0 0 1 1.6.155c.446.087.87.26 1.251.508a2.4 2.4 0 0 1 .8.939c.2.454.295.947.28 1.443a2.71 2.71 0 0 1-.417 1.555c-.32.443-.755.792-1.256 1.011c.681.17 1.28.577 1.69 1.147c.38.587.573 1.276.552 1.975a3.418 3.418 0 0 1-.358 1.614c-.23.444-.564.826-.974 1.113a4.32 4.32 0 0 1-1.4.64a6.28 6.28 0 0 1-1.603.206l.005-.001Zm-3.26-5.483v3.39H7.6c.255.003.51-.023.76-.077c.229-.044.446-.133.64-.262a1.29 1.29 0 0 0 .44-.491c.118-.25.173-.524.162-.8a1.614 1.614 0 0 0-.534-1.358a2.286 2.286 0 0 0-1.414-.4l-2.944-.002Zm0-4.725v2.875h2.742c.424.017.842-.104 1.192-.345a1.3 1.3 0 0 0 .464-1.119a1.487 1.487 0 0 0-.152-.708a1.122 1.122 0 0 0-.417-.428a1.677 1.677 0 0 0-.6-.215a3.575 3.575 0 0 0-.7-.061l-2.529.001Zm15.341-.056h-4.99V6.557h4.99v1.215Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17v1a3 3 0 1 1-6 0v-1m6 0H9m6 0h3.59c.383 0 .575 0 .73-.052a1 1 0 0 0 .627-.628c.053-.156.053-.348.053-.734c0-.169 0-.253-.014-.334a.998.998 0 0 0-.173-.42c-.048-.067-.108-.127-.227-.246l-.39-.39a.67.67 0 0 1-.196-.474V10a7 7 0 1 0-14 0v3.722a.67.67 0 0 1-.196.474l-.39.39c-.12.12-.179.179-.226.245a1 1 0 0 0-.175.421c-.013.08-.013.165-.013.334c0 .386 0 .578.052.734a1 1 0 0 0 .629.628c.155.052.346.052.729.052H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17v1a3 3 0 1 1-6 0v-1m6 0H9m6 0h3.59c.383 0 .575 0 .73-.052a1 1 0 0 0 .627-.628c.053-.156.053-.348.053-.734c0-.169 0-.253-.014-.334a.998.998 0 0 0-.173-.42c-.048-.067-.108-.127-.227-.246l-.39-.39a.67.67 0 0 1-.196-.474V10a7 7 0 1 0-14 0v3.722a.67.67 0 0 1-.196.474l-.39.39c-.12.12-.179.179-.226.245a1 1 0 0 0-.175.421c-.013.08-.013.165-.013.334c0 .386 0 .578.052.734a1 1 0 0 0 .629.628c.155.052.346.052.729.052H9m3-4v-3m0 0V7m0 3H9m3 0h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17v1a3 3 0 1 1-6 0v-1m6 0H9m6 0h3.59c.383 0 .575 0 .73-.052a1 1 0 0 0 .628-.628c.052-.156.052-.348.052-.734c0-.169 0-.253-.013-.334a.999.999 0 0 0-.174-.42c-.048-.067-.108-.127-.227-.246l-.39-.39a.67.67 0 0 1-.196-.474V10a7 7 0 1 0-14 0v3.722a.67.67 0 0 1-.196.474l-.39.39c-.12.12-.179.179-.226.245a1 1 0 0 0-.175.421c-.013.08-.013.165-.013.334c0 .386 0 .578.052.734a1 1 0 0 0 .628.628c.155.052.347.052.73.052H9m1-5l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellNotification(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17v1a3 3 0 1 1-6 0v-1m6 0H9m6 0h4a1 1 0 0 0 1-1v-.586a1 1 0 0 0-.293-.707l-.51-.51a.67.67 0 0 1-.197-.475V10c0-.176-.006-.351-.02-.524M9 17H5a1 1 0 0 1-1-1v-.586a1 1 0 0 1 .293-.707l.51-.511A.669.669 0 0 0 5 13.722V10a7 7 0 0 1 9.045-6.696m4.935 6.172a4 4 0 1 0-4.935-6.173m4.935 6.173a4 4 0 0 1-4.935-6.173m4.935 6.173s0 0 0 0m-4.935-6.172h.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellOff(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v1a3 3 0 1 0 6 0v-1m-6 0h6m-6 0H5a1 1 0 0 1-.993-.883L4 16v-.586a1 1 0 0 1 .293-.707l.51-.51A.67.67 0 0 0 5 13.722V10c0-1.842.711-3.517 1.874-4.767M15 17h4M5 3l16 16m-2-6.001V10a7 7 0 0 0-7-7l-.24.004a6.972 6.972 0 0 0-2.29.467"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17v1a3 3 0 1 1-6 0v-1m6 0H9m6 0h3.59c.383 0 .575 0 .73-.052a1 1 0 0 0 .627-.628c.053-.156.053-.348.053-.734c0-.169 0-.253-.014-.334a.998.998 0 0 0-.173-.42c-.048-.067-.108-.127-.227-.246l-.39-.39a.67.67 0 0 1-.196-.474V10a7 7 0 1 0-14 0v3.722a.67.67 0 0 1-.196.474l-.39.39c-.12.12-.179.179-.226.245a1 1 0 0 0-.175.421c-.013.08-.013.165-.013.334c0 .386 0 .578.052.734a1 1 0 0 0 .629.628c.155.052.346.052.729.052H9m0-7h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellRing(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17v1a3 3 0 1 1-6 0v-1m6 0H9m6 0h3.59c.383 0 .575 0 .73-.052a1 1 0 0 0 .627-.628c.053-.156.053-.348.053-.734c0-.169 0-.253-.014-.334a.998.998 0 0 0-.173-.42c-.048-.067-.108-.127-.227-.246l-.39-.39a.67.67 0 0 1-.196-.474V10a7 7 0 1 0-14 0v3.722a.67.67 0 0 1-.196.474l-.39.39c-.12.12-.179.179-.226.245a1 1 0 0 0-.175.421c-.013.08-.013.165-.013.334c0 .386 0 .578.052.734a1 1 0 0 0 .629.628c.155.052.346.052.729.052H9m9.019-14.986a10 10 0 0 1 3.153 4.002M5.982 2.014a10 10 0 0 0-3.154 4.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlackLivesMatter(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.172 3.703a.6.6 0 0 0-.162-.817l-1.18-.792a.551.551 0 0 0-.775.167l-1.46 2.273L9.88 5.719l1.293-2.016Zm2.005 3.729l1.068-1.664a.6.6 0 0 0-.163-.817l-1.178-.792a.551.551 0 0 0-.776.167l-1.232 1.921l2.281 1.185Zm1.289 3.426l-1.033-.477a.597.597 0 0 1-.24-.86L15.2 6.39a.551.551 0 0 1 .776-.167l.935.63a.6.6 0 0 1 .163.816l-1.91 2.977a.552.552 0 0 1-.699.21Zm5.276-2.102l-.935-.63a.551.551 0 0 0-.776.168l-1.713 2.67a.597.597 0 0 0 .24.86l1.033.476a.552.552 0 0 0 .699-.21l1.615-2.517a.6.6 0 0 0-.163-.817Z"/><path fill="currentColor" d="M17.133 13.371c.218.101.449.152.687.152a1.66 1.66 0 0 0 1.236-.558v1.715a3.488 3.488 0 0 1-1.482 2.877l-.351.243a.593.593 0 0 0-.25.487v3.127a.575.575 0 0 1-.564.586H7.878c-.367 0-.636-.358-.547-.728l1.031-4.286a.604.604 0 0 0-.094-.49l-4.157-5.83a.605.605 0 0 1-.016-.675L6.69 5.947a.552.552 0 0 1 .722-.199l4.563 2.37a.598.598 0 0 1 .217.85l-.345.537a.559.559 0 0 1-.47.261H7.816a.577.577 0 0 0-.564.614a.584.584 0 0 0 .58.558h2.413l2.615 5.432c.135.282.461.417.739.289c.296-.136.42-.5.277-.798l-2.37-4.923h.357l5.27 2.433Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h4.5M8 12V5h4.5a3.5 3.5 0 1 1 0 7M8 12v7h5.5a3.5 3.5 0 1 0 0-7h-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19.5V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C6.52 3 7.08 3 8.2 3h9.2c.56 0 .84 0 1.055.109a1 1 0 0 1 .436.437C19 3.76 19 4.04 19 4.6v11.8c0 .56 0 .84-.11 1.054a.998.998 0 0 1-.435.437C18.24 18 17.96 18 17.402 18H7.25A2.25 2.25 0 0 0 5 20.25c0 .414.336.75.75.75h10.652c.559 0 .84 0 1.053-.109a.998.998 0 0 0 .436-.437C18 20.24 18 19.96 18 19.4V18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpen(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9.8V20m0-10.2c0-1.68 0-2.52.327-3.162a3 3 0 0 1 1.31-1.311C14.28 5 15.12 5 16.8 5h2.6c.56 0 .84 0 1.054.109c.188.096.34.249.437.437C21 5.76 21 6.04 21 6.6v8.8c0 .56 0 .84-.11 1.054a.999.999 0 0 1-.436.437c-.213.109-.493.109-1.052.109h-2.833c-.939 0-1.41 0-1.836.13a2.996 2.996 0 0 0-1.032.552c-.344.283-.605.674-1.126 1.455L12 20m0-10.2c0-1.68 0-2.52-.327-3.162a3 3 0 0 0-1.311-1.311C9.72 5 8.88 5 7.2 5H4.6c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C3 5.76 3 6.04 3 6.6v8.8c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C3.76 17 4.039 17 4.598 17h2.833c.939 0 1.408 0 1.835.13c.377.114.73.302 1.034.552c.343.282.602.67 1.118 1.446L12 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7.2v9.485c0 1.361 0 2.042.204 2.458a2 2 0 0 0 2.06 1.102c.46-.06 1.026-.438 2.158-1.193l.003-.002c.449-.3.673-.449.908-.532a2 2 0 0 1 1.333 0c.235.083.46.233.911.534c1.133.755 1.7 1.132 2.16 1.193a2 2 0 0 0 2.059-1.102c.204-.416.204-1.097.204-2.458V7.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C16.48 4 15.92 4 14.8 4H9.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C6 5.52 6 6.08 6 7.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21H4a1 1 0 0 1-1-1v-7.3a1 1 0 0 1 .341-.752L5 10.5V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1ZM9 7.329a1 1 0 0 1 .658.248l5 4.375A1 1 0 0 1 15 12.7V19h4V5H7v3.75l1.341-1.174A1 1 0 0 1 9 7.329ZM8 16h2v3h3v-5.843l-4-3.5l-4 3.5V19h3v-3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingFour(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h2m0 0h10M4 20V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 3 6.08 3 7.2 3h3.6c1.12 0 1.68 0 2.107.218c.377.192.684.497.875.874c.218.427.218.987.218 2.105V12m0 8h6m-6 0v-8m6 8h2m-2 0v-8c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.082-1.083C18.398 9 17.932 9 17 9s-1.398 0-1.766.152a2 2 0 0 0-1.082 1.083C14 10.602 14 11.068 14 12m-7-2h4M7 7h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h2m0 0h5m-5 0V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h1.6c1.12 0 1.68 0 2.107.218c.377.192.684.497.875.874c.218.427.218.987.218 2.105V10M9 20h11M9 20v-5.632c0-.525 0-.788.063-1.033c.056-.217.148-.423.272-.61c.14-.21.336-.387.726-.737l2.302-2.068c.755-.678 1.133-1.017 1.56-1.146a2 2 0 0 1 1.154 0c.428.129.806.468 1.562 1.147l2.3 2.067c.39.35.585.527.726.737c.124.187.216.393.271.61c.064.245.064.508.064 1.033V20m0 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingThree(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h2m0 0h11M4 20v-5.632c0-.525 0-.788.063-1.033c.056-.217.148-.423.272-.61c.14-.21.336-.387.727-.737L7.363 9.92c.755-.678 1.132-1.017 1.56-1.146a2 2 0 0 1 1.154 0c.428.129.806.468 1.562 1.147l2.3 2.067c.39.35.585.527.726.737c.124.187.216.393.272.61c.063.245.063.508.063 1.033V20m0 0h5m0 0h2m-2 0V7.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4h-6.6c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C7 5.52 7 6.08 7 7.2V10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h2m0 0h8m-8 0V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h1.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105v3.05M12 20h8m-8 0v-9.754M20 20h2m-2 0v-5.632c0-.525 0-.788-.063-1.033a1.998 1.998 0 0 0-.272-.61c-.14-.21-.335-.386-.726-.737l-2.3-2.067c-.756-.679-1.134-1.018-1.562-1.147a2 2 0 0 0-1.154 0c-.428.129-.806.468-1.562 1.147l-.361.325"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bulb(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 21h6M12 3a6 6 0 0 0-5.019 9.29c.954 1.452 1.43 2.178 1.493 2.286c.55.965.449.625.518 1.734c.008.124.008.313.008.69a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1c0-.377 0-.566.008-.69c.07-1.11-.033-.769.518-1.734c.062-.108.54-.834 1.493-2.287A6 6 0 0 0 12 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 8v8.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V8M4 8v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4H8m12 4v-.803c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4H16m0-2v2m0 0H8m0-2v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 8v8.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.607c1.118 0 1.677 0 2.104-.218c.376-.192.682-.498.874-.874c.218-.428.218-.986.218-2.104V8M4 8v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4H8m12 4v-.803c0-1.118 0-1.678-.218-2.105a2 2 0 0 0-.874-.874C18.48 4 17.92 4 16.8 4H16M8 4h8M8 4V2m8 2V2m-1.5 12H12m0 0H9.5m2.5 0v-2.5m0 2.5v2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCalendar(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.999 22h-14a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2V2h2v2h6V2h2v2h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2Zm-14-12v10h14V10h-14Zm0-4v2h14V6h-14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 8v8.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V8M4 8v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4H8m12 4v-.803c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4H16M8 4h8M8 4V2m8 2V2m-1 10l-4 4l-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 8v8.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.607c1.118 0 1.677 0 2.104-.218c.376-.192.682-.498.874-.874c.218-.428.218-.986.218-2.104V8M4 8v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4H8m12 4v-.803c0-1.118 0-1.678-.218-2.105a2 2 0 0 0-.874-.874C18.48 4 17.92 4 16.8 4H16M8 4h8M8 4V2m8 2V2m-2 14l-2-2m0 0l-2-2m2 2l2-2m-2 2l-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDays(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h-.8c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 5.52 4 6.08 4 7.2V8m4-4h8M8 4V2m8 2h.8c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105V8m-4-4V2M4 8v8.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V8M4 8h16m-4 8h.002v.002H16zm-4 0h.002v.002H12zm-4 0h.002v.002H8zm8.002-4v.002H16V12zM12 12h.002v.002H12zm-4 0h.002v.002H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarEdit(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2V2h2v2h6V2h2v2h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM5 10v10h14V10H5Zm0-4v2h14V6H5Zm4.8 13H8v-1.8l4.2-4.19l1.8 1.8L9.8 19Zm4.825-4.818l-1.8-1.8l1.375-1.369l1.8 1.8l-1.37 1.37l-.005-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarEvent(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 8v8.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V8M4 8v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4H8m12 4v-.803c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4H16M8 4h8M8 4V2m8 2V2m-4.25 14a.25.25 0 0 0 .25-.25v-3.5a.25.25 0 0 0-.25-.25h-3.5a.25.25 0 0 0-.25.25v3.5c0 .138.112.25.25.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2V2h2v2h6V2h2v2h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM5 10v10h14V10H5Zm0-4v2h14V6H5Zm10 10H9v-2h6v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2V2h2v2h6V2h2v2h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM5 10v10h14V10H5Zm0-4v2h14V6H5Zm8 12h-2v-2H9v-2h2v-2h2v2h2v2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 8v8.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V8M4 8v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4H8m12 4v-.803c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4H16M8 4h8M8 4V2m8 2V2m-1.5 12h-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarWeek(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 8v8.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V8M4 8v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4H8m12 4v-.803c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4H16M8 4h8M8 4V2m8 2V2m0 10H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarX(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2V2h2v2h6V2h2v2h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM5 10v10h14V10H5Zm0-4v2h14V6H5Zm4.706 12.707l-1.413-1.414L10.586 15l-2.293-2.293l1.414-1.414L12 13.586l2.293-2.292l1.414 1.414L13.414 15l2.293 2.293l-1.413 1.413L12 16.414l-2.293 2.293h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.489 7H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 8.52 3 9.08 3 10.2v5.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104v-5.606c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C19.48 7 18.92 7 17.8 7h-3.29M9.49 7h.062M9.49 7h.062M9.49 7c-.106 0-.166 0-.213-.006a1.001 1.001 0 0 1-.867-1.203a2.71 2.71 0 0 1 .08-.257l.001-.006c.052-.154.077-.23.106-.299a2 2 0 0 1 1.699-1.224C10.368 4 10.449 4 10.61 4h2.778c.162 0 .243 0 .317.005a2 2 0 0 1 1.698 1.224c.029.068.054.145.106.3c.046.138.07.207.08.262a1 1 0 0 1-.866 1.203A2.125 2.125 0 0 1 14.51 7M9.55 7h4.898m0 0h.062m-.062 0h.062M12 16a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarAuto(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 11h1.045m0 0h15.91m-15.91 0c.012-.052.025-.102.04-.153a4.04 4.04 0 0 1 .19-.467L5.822 6.9c.305-.687.458-1.032.7-1.284c.214-.223.476-.393.766-.498C7.617 5 7.993 5 8.746 5h6.508c.752 0 1.13 0 1.458.118c.29.105.552.275.765.498c.242.252.395.596.7 1.283l1.553 3.493c.099.223.15.337.185.455c.015.05.028.101.04.153m-15.91 0a2.021 2.021 0 0 0-.03.174C4 11.3 4 11.426 4 11.68V17m15.955-6H21m-1.046 0c.013.058.023.116.03.174c.016.124.016.25.016.5V17m0 0h-4m4 0v1a2 2 0 1 1-4 0v-1m0 0H8m0 0H4m4 0v1a2 2 0 1 1-4 0v-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretCircleDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l-3 3l-3-3m12 1a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretCircleLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 15l-3-3l3-3m8 3a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretCircleRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 9l3 3l-3 3m10-3a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretCircleUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 13l3-3l3 3m6-1a9 9 0 1 0-18 0a9 9 0 0 0 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 14.5l5-5H7l5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 10l-4 4l-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l-3 3l-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.5 12l5 5V7l-5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeftSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 15l-3-3l3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.5 12l-5-5v10l5-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRightSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 9l3 3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 9.5l-5 5h10l-5-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 14l4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 13l3-3l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cast(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 21h-7v-2h7V5H3v3H1V5a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2Zm-9 0h-2a8.94 8.94 0 0 0-2.639-6.361A8.939 8.939 0 0 0 1 12v-2a10.929 10.929 0 0 1 7.776 3.22A10.927 10.927 0 0 1 12 21Zm-4 0H6a4.968 4.968 0 0 0-1.466-3.534A4.966 4.966 0 0 0 1 16v-2a6.955 6.955 0 0 1 4.951 2.049A6.956 6.956 0 0 1 8 21Zm-4 0H1v-3a3 3 0 0 1 3 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarHorizontalOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 9H4m9 0V4.6c0-.56 0-.84-.11-1.054a.997.997 0 0 0-.435-.437C12.24 3 11.96 3 11.4 3H4v6m9 0h5.4c.56 0 .84 0 1.055.109a1 1 0 0 1 .436.437C20 9.76 20 10.04 20 10.6v2.8c0 .56 0 .84-.11 1.054a.998.998 0 0 1-.435.437C19.24 15 18.96 15 18.402 15H16M4 9v6m0 0v6h10.402c.559 0 .84 0 1.053-.109a.998.998 0 0 0 .436-.437C16 20.24 16 19.96 16 19.4V15M4 15h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarVerticalOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 11v9m0-9H4.6c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C3 11.76 3 12.04 3 12.6V20h6m0-9V5.6c0-.56 0-.84.109-1.054a1 1 0 0 1 .437-.437C9.76 4 10.04 4 10.6 4h2.8c.56 0 .84 0 1.054.109c.188.096.34.249.437.437C15 4.76 15 5.04 15 5.6V8M9 20h6m0 0h6V9.6c0-.56 0-.84-.11-1.054a.997.997 0 0 0-.435-.437C20.24 8 19.96 8 19.4 8H15m0 12V8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLine(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15v1.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218H21M3 15V5m0 10l3.853-3.21l.004-.003c.697-.581 1.046-.872 1.425-.99a2 2 0 0 1 1.362.061c.367.153.688.474 1.332 1.118l.006.006c.654.654.981.982 1.354 1.133a2 2 0 0 0 1.385.046c.383-.128.733-.434 1.433-1.046L21 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a9 9 0 1 0 9 9m-9-9a9 9 0 0 1 9 9m-9-9v9m9 0h-9m6 6.5L12 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.6 19.92l1.524-1.219l.01-.008c.318-.255.479-.383.658-.474c.16-.082.331-.142.508-.178c.199-.041.406-.041.822-.041h8.681c1.118 0 1.678 0 2.105-.218a2 2 0 0 0 .874-.874C21 16.48 21 15.92 21 14.804V7.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C19.48 4 18.92 4 17.8 4H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 5.52 3 6.08 3 7.2v11.471c0 1.066 0 1.599.218 1.872a1 1 0 0 0 .783.377c.35 0 .766-.334 1.599-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14v-3m0 0V8m0 3H9m3 0h3m-7.876 7.701L5.6 19.921c-.833.665-1.249.998-1.599.999a1 1 0 0 1-.783-.377C3 20.27 3 19.737 3 18.671V7.201c0-1.12 0-1.681.218-2.11c.192-.376.497-.681.874-.873C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105v7.607c0 1.117 0 1.676-.218 2.104a2 2 0 0 1-.874.874c-.427.218-.986.218-2.104.218H9.123c-.416 0-.625 0-.824.04a2 2 0 0 0-.507.179c-.18.092-.342.221-.665.48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2a2 2 0 0 1 2 2H4v11.177a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h12Z"/><path fill="currentColor" d="m14 22l-2.667-2.823H8a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v9.177a2 2 0 0 1-2 2h-3.333L14 22Zm1.805-4.823H20V8H8v9.177h4.195L14 19.087l1.805-1.91Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 9l-4 4l-2-2m-1.876 7.701L5.6 19.921c-.833.665-1.249.998-1.599.999a1 1 0 0 1-.783-.377C3 20.27 3 19.737 3 18.671V7.201c0-1.12 0-1.681.218-2.11c.192-.376.497-.681.874-.873C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105v7.607c0 1.117 0 1.676-.218 2.104a2 2 0 0 1-.874.874c-.427.218-.986.218-2.104.218H9.123c-.416 0-.625 0-.824.04a2 2 0 0 0-.507.179c-.18.092-.342.221-.665.48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.51 19.802a9 9 0 1 0-3.312-3.312l.003.005c.073.127.11.191.127.252c.016.057.02.108.016.168a1.058 1.058 0 0 1-.07.26l-.768 2.307l-.001.003c-.162.487-.243.73-.186.892c.05.142.163.253.304.304c.162.057.404-.023.889-.185l.006-.002l2.306-.769c.131-.044.198-.066.262-.07a.471.471 0 0 1 .167.017a1.3 1.3 0 0 1 .253.127z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatCircleAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v-3m0 0V9m0 3H9m3 0h3m-3 9a8.96 8.96 0 0 1-4.49-1.198a1.276 1.276 0 0 0-.257-.13a.475.475 0 0 0-.167-.017a1.12 1.12 0 0 0-.258.07l-2.31.769h-.002c-.487.163-.731.245-.893.187a.5.5 0 0 1-.304-.304c-.057-.162.024-.405.186-.892v-.003l.77-2.306l.002-.005c.042-.129.064-.194.068-.256a.478.478 0 0 0-.017-.168a1.228 1.228 0 0 0-.127-.252l-.003-.005A9 9 0 1 1 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatCircleCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 10l-4 4l-2-2m3 9a8.96 8.96 0 0 1-4.49-1.198a1.276 1.276 0 0 0-.257-.13a.475.475 0 0 0-.167-.017a1.12 1.12 0 0 0-.258.07l-2.31.769h-.002c-.487.163-.731.245-.893.187a.5.5 0 0 1-.304-.304c-.057-.162.024-.405.186-.892v-.003l.77-2.306l.002-.005c.042-.129.064-.194.068-.256a.478.478 0 0 0-.017-.168a1.228 1.228 0 0 0-.127-.252l-.003-.005A9 9 0 1 1 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatCircleClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m-2 7a8.959 8.959 0 0 1-4.49-1.198l-.004-.003c-.128-.073-.192-.11-.253-.127a.475.475 0 0 0-.167-.017c-.064.004-.13.026-.262.07l-2.306.769l-.006.002c-.485.162-.727.242-.889.185a.502.502 0 0 1-.304-.304c-.057-.162.024-.405.186-.892v-.003l.77-2.306c.044-.132.065-.198.07-.261a.498.498 0 0 0-.017-.168a1.228 1.228 0 0 0-.127-.252l-.003-.005A9 9 0 1 1 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatCircleDots(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.51 19.802a9 9 0 1 0-3.312-3.312l.003.005c.073.127.11.191.127.252c.016.057.02.108.016.168a1.058 1.058 0 0 1-.07.26l-.768 2.307l-.001.003c-.162.487-.243.73-.186.892c.05.142.163.253.304.304c.162.057.404-.023.889-.185l.006-.002l2.306-.769c.131-.044.198-.066.262-.07a.471.471 0 0 1 .167.017a1.3 1.3 0 0 1 .253.127z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatCircleRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-3 9a8.959 8.959 0 0 1-4.49-1.198l-.004-.003c-.128-.073-.192-.11-.253-.127a.475.475 0 0 0-.167-.017c-.064.004-.13.026-.262.07l-2.306.769l-.006.002c-.485.162-.727.242-.889.185a.502.502 0 0 1-.304-.304c-.057-.162.024-.405.186-.892v-.003l.77-2.306c.044-.132.065-.198.07-.261a.498.498 0 0 0-.017-.168a1.228 1.228 0 0 0-.127-.252l-.003-.005A9 9 0 1 1 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 13l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m-6.876 5.701L5.6 19.921c-.833.665-1.249.998-1.599.999a1 1 0 0 1-.783-.377C3 20.27 3 19.737 3 18.671V7.201c0-1.12 0-1.681.218-2.11c.192-.376.497-.681.874-.873C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v7.607c0 1.117 0 1.676-.218 2.104a2 2 0 0 1-.874.874c-.427.218-.987.218-2.105.218h-8.68c-.417 0-.624 0-.823.04a2.004 2.004 0 0 0-.508.179c-.18.091-.34.22-.657.474z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatConversation(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8h4a1 1 0 0 1 1 1v11l-3.333-2.769a1.002 1.002 0 0 0-.64-.231H9a1 1 0 0 1-1-1v-3m8-5V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v11l3.333-2.77c.18-.148.406-.23.64-.23H8m8-5v4a1 1 0 0 1-1 1H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatConversationCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.338 15.99a6 6 0 1 0-5.408-2.78l-.424 1.272v.002c-.163.487-.244.73-.187.893c.05.141.163.253.304.303c.162.058.404-.023.888-.184l.007-.002l1.272-.424a5.971 5.971 0 0 0 3.548.92m0 0s0 0 0 0m0 0a6.003 6.003 0 0 0 8.872 3.08l1.272.424h.003c.487.163.73.244.893.186a.5.5 0 0 0 .302-.303c.058-.162-.023-.406-.186-.895l-.424-1.272l.142-.235A6 6 0 0 0 15 8l-.225.004l-.113.006"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatDots(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.6 19.92l1.524-1.219l.01-.008c.318-.255.479-.383.658-.474c.16-.082.331-.142.508-.178c.199-.041.406-.041.822-.041h8.681c1.118 0 1.678 0 2.105-.218a2 2 0 0 0 .874-.874C21 16.48 21 15.92 21 14.804V7.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C19.48 4 18.92 4 17.8 4H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 5.52 3 6.08 3 7.2v11.471c0 1.066 0 1.599.218 1.872a1 1 0 0 0 .783.377c.35 0 .766-.334 1.599-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 11h6m-7.876 7.701L5.6 19.921c-.833.665-1.249.998-1.599.999a1 1 0 0 1-.783-.377C3 20.27 3 19.737 3 18.671V7.201c0-1.12 0-1.681.218-2.11c.192-.376.497-.681.874-.873C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v7.607c0 1.117 0 1.676-.218 2.104a2 2 0 0 1-.874.874c-.427.218-.987.218-2.105.218h-8.68c-.417 0-.624 0-.823.04a2.004 2.004 0 0 0-.508.179c-.18.091-.34.22-.657.474z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 12l4.243 4.243l8.484-8.486"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckAll(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 12.485l4.243 4.243l8.484-8.485M3 12.485l4.243 4.243m8.485-8.485L12.5 11.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckAllBig(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 12l4.95 4.95L22.557 6.343M2.05 12.05L7 17M17.606 6.394l-5.303 5.303"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBig(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 12l4.95 4.95L19.557 6.343"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBold(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.818 19.779l-6.364-6.364l2.83-2.83l3.534 3.544l9.898-9.908l2.83 2.83L8.818 19.779Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkbox(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 19H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2ZM7 7v10h10V7H7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 12l3 3l5-6M4 16.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v9.607c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C4 18.48 4 17.92 4 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxChecked(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 19H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2ZM7 7v10h10V7H7Zm4 8.362l-2.7-2.647l1.4-1.43l1.3 1.271l3.3-3.267l1.4 1.422l-4.7 4.65v.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxFill(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 7.2v9.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V7.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 5.52 4 6.08 4 7.2"/><path d="M15 13.4v-2.8c0-.56 0-.84-.11-1.054a.998.998 0 0 0-.436-.437C14.24 9 13.96 9 13.4 9h-2.8c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C9 9.76 9 10.04 9 10.6v2.8c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437c.214.109.493.109 1.052.109h2.803c.559 0 .84 0 1.053-.109a1 1 0 0 0 .437-.437C15 14.24 15 13.96 15 13.4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxSquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 19H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2ZM7 7v10h10V7H7Zm8 8H9V9h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxUnchecked(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7.2v9.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V7.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 4 17.92 4 16.8 4H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 5.52 4 6.08 4 7.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronBigDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.515 8.465L12 16.95l8.485-8.485L19.07 7.05L12 14.122L4.929 7.05L3.515 8.465Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronBigLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.535 3.515L7.05 12l8.485 8.485l1.415-1.414L9.878 12l7.072-7.071l-1.415-1.414Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronBigRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.465 20.485L16.95 12L8.465 3.515L7.05 4.929L14.122 12L7.05 19.071l1.415 1.414Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronBigUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.485 15.535L12 7.05l-8.485 8.485L4.93 16.95L12 9.878l7.071 7.072l1.414-1.415Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 9l-7 7l-7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownDuo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 13l-4 4l-4-4m8-6l-4 4l-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDuoDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 18.424l-6.01-6.01l1.414-1.415l4.6 4.6l4.6-4.6l1.406 1.415l-6.009 6.01H12ZM12 13L5.99 6.989l1.414-1.414l4.6 4.6l4.6-4.6l1.406 1.414l-6.009 6.01L12 13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDuoLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.585 18.01L5.575 12l6.01-6.01L13 7.404l-4.6 4.6l4.6 4.6l-1.414 1.406h-.001Zm5.425 0L10.999 12l6.011-6.01l1.414 1.414l-4.6 4.6l4.6 4.6l-1.413 1.406h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDuoRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.414 18.01l-1.415-1.413l4.6-4.6l-4.6-4.6l1.415-1.407l6.01 6.01l-6.009 6.01h-.001Zm-5.425 0l-1.414-1.413l4.6-4.6l-4.6-4.593L6.989 5.99L13 12l-6.01 6.01h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDuoUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.403 18.425L5.99 17.01L12 10.999l6.01 6.011l-1.413 1.413l-4.6-4.6l-4.6 4.6l.006.002Zm0-5.424L5.99 11.585L12 5.575l6.01 6.01l-1.413 1.414l-4.6-4.6l-4.6 4.6l.006.002Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 19l-7-7l7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftDuo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 16l-4-4l4-4m-6 8l-4-4l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 16l-4-4l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 5l7 7l-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightDuo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 8l4 4l-4 4M7 8l4 4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 8l4 4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 16l7-7l7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpDuo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 17l4-4l4 4m-8-6l4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chromecast(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8.25V8.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v7.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.874.875C19.48 19 18.92 19 17.803 19H14m-9 0a2 2 0 0 0-2-2m5 2a5 5 0 0 0-5-5m8 5a8 8 0 0 0-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 10l-4 4l-2-2m3 9a9 9 0 1 1 0-18a9 9 0 0 1 0 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheckOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-.016-2H12a8 8 0 1 0-.016 0ZM10 17l-4-4l1.41-1.41L10 14.17l6.59-6.59L18 9l-8 8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm0-18a8 8 0 1 0 8 8a8.009 8.009 0 0 0-8-8Zm-.05 11.5L7 10.55l1.414-1.414l3.536 3.535l3.536-3.535L16.9 10.55l-4.95 4.95Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm0-18a8 8 0 1 0 8 8.18v1.783V12a8.009 8.009 0 0 0-8-8Zm1.45 12.9L8.5 11.95L13.45 7l1.414 1.414l-3.536 3.536l3.535 3.536l-1.412 1.414h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c-5.52-.006-9.994-4.48-10-10v-.2C2.11 6.305 6.635 1.928 12.13 2c5.497.074 9.904 4.569 9.868 10.065C21.962 17.562 17.497 22 12 22Zm0-18a8 8 0 1 0 8 8a8.009 8.009 0 0 0-8-8Zm-1.45 13l-1.414-1.415l3.535-3.535l-3.535-3.535L10.55 7.1l4.95 4.95L10.551 17h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-.016-2H12a8 8 0 1 0-.016 0Zm-3.47-5.136L7.1 13.45l4.95-4.95L17 13.45l-1.414 1.413l-3.536-3.535l-3.535 3.536h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm0-18a8 8 0 1 0 8 8a8.009 8.009 0 0 0-8-8Zm0 13l-5-5l1.41-1.41L11 13.17V7h2v6.17l2.59-2.58L17 12l-5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleHelp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.146 9.074a2.998 2.998 0 0 1 5.28-.838A3 3 0 0 1 12 13v1m0 7a9 9 0 1 1 0-18a9 9 0 0 1 0 18m.05-4v.1h-.1V17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm0-18a8 8 0 1 0 8 8.18v1.783V12a8.009 8.009 0 0 0-8-8Zm0 13l-5-5l5-5l1.41 1.41L10.83 11H17v2h-6.17l2.58 2.59L12 17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c-5.52-.006-9.994-4.48-10-10v-.2C2.11 6.305 6.635 1.928 12.13 2c5.497.074 9.904 4.569 9.868 10.065C21.962 17.562 17.497 22 12 22Zm0-18a8 8 0 1 0 8 8a8.009 8.009 0 0 0-8-8Zm0 13l-1.41-1.41L13.17 13H7v-2h6.17l-2.58-2.59L12 7l5 5l-5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-.016-2H12a8 8 0 1 0-.016 0ZM13 17h-2v-6.17l-2.59 2.58L7 12l5-5l5 5l-1.41 1.41L13 10.83V17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleWarning(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8.45v4M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18m.05-5.55v.1h-.1v-.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7v5h5m-5 9a9 9 0 1 1 0-18a9 9 0 0 1 0 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseBig(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.59 5L12 10.59L6.41 5L5 6.41L10.59 12L5 17.59L6.41 19L12 13.41L17.59 19L19 17.59L13.41 12L19 6.41L17.59 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 9l3 3m0 0l3 3m-3-3l-3 3m3-3l3-3m-3 12a9 9 0 1 1 0-18a9 9 0 0 1 0 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21l-9-9m0 0L3 3m9 9l9-9m-9 9l-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 18l-6-6m0 0L6 6m6 6l6-6m-6 6l-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSm(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 16l-4-4m0 0L8 8m4 4l4-4m-4 4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSmall(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.59 7L12 10.59L8.41 7L7 8.41L10.59 12L7 15.59L8.41 17L12 13.41L15.59 17L17 15.59L13.41 12L17 8.41L15.59 7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 9l3 3m0 0l3 3m-3-3l-3 3m3-3l3-3M4 16.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C20 5.52 20 6.08 20 7.2v9.6c0 1.12 0 1.68-.218 2.108a2.001 2.001 0 0 1-.874.874c-.428.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C4 18.48 4 17.92 4 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11a4 4 0 0 1 0 8H6a5 5 0 0 1-.331-9.99A7.002 7.002 0 0 1 18.929 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v-3m0 0v-3m0 3H9m3 0h3m8 2a4 4 0 0 0-4.07-4A7.002 7.002 0 0 0 5.669 9.01A5 5 0 0 0 6 19h13a4 4 0 0 0 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l-4 4l-2-2m14 2a4 4 0 0 0-4.07-4A7.002 7.002 0 0 0 5.669 9.01A5 5 0 0 0 6 19h13a4 4 0 0 0 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 15l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m9 0a4 4 0 0 0-4.07-4A7.002 7.002 0 0 0 5.669 9.01A5 5 0 0 0 6 19h13a4 4 0 0 0 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#ciCloudDown0)"><path fill="currentColor" d="M19 20H6a6 6 0 0 1-.974-11.921A8.018 8.018 0 0 1 12 3.999a7.916 7.916 0 0 1 4.962 1.725a8.041 8.041 0 0 1 2.8 4.334A5 5 0 0 1 19 20ZM12 6a6.014 6.014 0 0 0-5.232 3.061L6.3 9.9l-.95.155A4 4 0 0 0 6 18h13a3 3 0 0 0 .46-5.965l-1.316-.2l-.322-1.292A5.988 5.988 0 0 0 12 6Zm0 10l-4-4h2.55V9h2.9v3H16l-4 4Z"/></g><defs><clipPath id="ciCloudDown0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v6m0 0l3-2m-3 2l-3-2m14 2a4 4 0 0 0-4.07-4A7.002 7.002 0 0 0 5.669 9.01A5 5 0 0 0 6 19h13a4 4 0 0 0 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOff(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 19H6a5 5 0 0 1-.332-9.99a7.018 7.018 0 0 1 1.333-1.909M19 19L5 5m14 14l2 2M10 5.29A7.002 7.002 0 0 1 18.93 11H19a3.999 3.999 0 0 1 3.122 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#ciCloudOutline0)"><path fill="currentColor" d="M19 20H6a6 6 0 0 1-.974-11.921A8.018 8.018 0 0 1 12 3.999a7.916 7.916 0 0 1 4.96 1.725a8.041 8.041 0 0 1 2.8 4.333A5 5 0 0 1 19 20ZM12 6a6.014 6.014 0 0 0-5.232 3.061L6.3 9.9l-.95.155A4 4 0 0 0 6 18h13a3 3 0 0 0 .46-5.965l-1.316-.2l-.322-1.292A5.988 5.988 0 0 0 12 6Z"/></g><defs><clipPath id="ciCloudOutline0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 13h6m8 2a4 4 0 0 0-4.07-4A7.002 7.002 0 0 0 5.669 9.01A5 5 0 0 0 6 19h13a4 4 0 0 0 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#ciCloudUp0)"><path fill="currentColor" d="M19 20H6a6 6 0 0 1-.974-11.921A8.018 8.018 0 0 1 12 3.999a7.916 7.916 0 0 1 4.962 1.725a8.041 8.041 0 0 1 2.8 4.334A5 5 0 0 1 19 20ZM12 6a6.014 6.014 0 0 0-5.232 3.061L6.3 9.9l-.95.155A4 4 0 0 0 6 18h13a3 3 0 0 0 .46-5.965l-1.316-.2l-.322-1.292A5.988 5.988 0 0 0 12 6Zm1.45 10h-2.9v-3H8l4-4l4 4h-2.55v3Z"/></g><defs><clipPath id="ciCloudUp0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v-6m0 0l-3 2m3-2l3 2m8 3a4 4 0 0 0-4.07-4A7.002 7.002 0 0 0 5.669 9.01A5 5 0 0 0 6 19h13a4 4 0 0 0 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 7l5 5l-5 5m-6 0l-5-5l5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeToGo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.686 15h10.629m-10.8-2h10.971M7.743 7h8.514M7.743 7c-.605 0-.907 0-1.13.12a1.006 1.006 0 0 0-.44.48c-.101.233-.076.534-.024 1.136l.8 9.337c.089 1.032.133 1.55.362 1.94a2 2 0 0 0 .863.792c.41.195.927.195 1.964.195h3.725c1.036 0 1.554 0 1.963-.195a2 2 0 0 0 .863-.792c.23-.39.274-.907.362-1.94l.8-9.338c.052-.602.078-.902-.023-1.134a1 1 0 0 0-.44-.48C17.165 7 16.862 7 16.258 7M7.742 7H6.75c-.9 0-1.35 0-1.611-.188a1 1 0 0 1-.41-.707c-.034-.32.19-.712.637-1.494V4.61c.334-.585.501-.877.736-1.09a2 2 0 0 1 .727-.422C7.129 3 7.467 3 8.14 3h7.72c.674 0 1.01 0 1.312.098c.27.087.517.231.726.422c.235.213.402.505.736 1.09c.447.782.671 1.174.637 1.495a1 1 0 0 1-.41.707C18.602 7 18.15 7 17.252 7h-.994"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h6.943m0 0h.114m-.114 0h.114m-.114 0A7 7 0 0 1 4 13V8.923c0-.51.413-.923.923-.923h12.154c.51 0 .923.413.923.923V9m-6.943 11H18m-6.943 0A7 7 0 0 0 18 13m0-4h1.5a2.5 2.5 0 0 1 0 5H18v-1m0-4v4M15 3l-1 2m-2-2l-1 2M9 3L8 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeTogo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.13 3.008A2 2 0 0 1 6.866 2h10.268a2 2 0 0 1 1.736 1.008l1.143 2c.648 1.134.049 2.503-1.09 2.887c.005.082.004.164-.002.247l-.857 12A2 2 0 0 1 16.069 22H7.93a2 2 0 0 1-1.995-1.858l-.857-12a2.023 2.023 0 0 1-.003-.247c-1.138-.384-1.737-1.753-1.09-2.887l1.144-2ZM7.074 8l.214 3h9.424l.214-3H7.074Zm.643 9l.214 3h8.138l.214-3H7.717Zm10.56-11l-1.143-2H6.866L5.723 6h12.554Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Color(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.75 22H4.917a2.086 2.086 0 0 1-2.084-2.084v-9.7a1.977 1.977 0 0 1-.592-2.558c.9-1.608 1.818-3.193 2.721-4.712A1.976 1.976 0 0 1 6.666 2c.424.004.836.139 1.18.385l.024.015c.412.24 4.392 2.609 7.591 4.512l2.84 1.688l.13.077a372.255 372.255 0 0 1 2.744 1.628a1.963 1.963 0 0 1 .583 2.553c-.883 1.577-1.8 3.162-2.72 4.712a1.983 1.983 0 0 1-1.68.951c-.136 0-.27-.015-.4-.045c-.117.069-.757.48-1.568 1c-1.443.925-3.415 2.19-3.643 2.307a1.976 1.976 0 0 1-.997.217ZM4.5 13.03v6.886c0 .23.187.417.417.417H8.74A485.729 485.729 0 0 1 4.5 13.03Zm1.358-1c.4.687 4.716 8.132 4.752 8.182a.3.3 0 0 0 .38.082c.134-.076 4.193-2.684 4.233-2.71c-.314-.185-2.586-1.539-4.992-2.972L5.838 12l.007.011l.01.017v.006l.003-.004Zm.474-8.13c-.808 1.3-2.616 4.528-2.64 4.587a.312.312 0 0 0 .124.386c.133.079 13.354 7.956 13.48 7.976a.32.32 0 0 0 .048 0a.31.31 0 0 0 .266-.149a414.198 414.198 0 0 0 2.694-4.667a.316.316 0 0 0-.125-.387c-.127-.076-13.323-7.918-13.408-7.955a.314.314 0 0 0-.383.117l-.007.012l-.008.012l-.012.02v.005l-.015.024l-.015.019h.001Zm1.353 5.49a1.666 1.666 0 1 1 1.667-1.67a1.669 1.669 0 0 1-1.667 1.667v.003Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Columns(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 6.5v11c0 .465 0 .697.038.89a2 2 0 0 0 1.572 1.571c.193.039.425.039.89.039s.697 0 .89-.039a2 2 0 0 0 1.571-1.57c.039-.194.039-.426.039-.891v-11c0-.465 0-.697-.039-.89a2 2 0 0 0-1.57-1.572C17.196 4 16.964 4 16.5 4s-.697 0-.89.038a2 2 0 0 0-1.572 1.572C14 5.803 14 6.035 14 6.5m-9 0v11c0 .465 0 .697.038.89a2 2 0 0 0 1.572 1.571c.193.039.425.039.89.039s.697 0 .89-.039a2 2 0 0 0 1.571-1.57c.039-.194.039-.426.039-.891v-11c0-.465 0-.697-.039-.89a2 2 0 0 0-1.57-1.572C8.196 4 7.964 4 7.5 4s-.697 0-.89.038A2 2 0 0 0 5.038 5.61C5 5.803 5 6.035 5 6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CombineCells(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v1m0-6v2m0-6v1m-8 8.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105v9.607c0 1.118 0 1.677-.218 2.104a2.001 2.001 0 0 1-.874.874c-.428.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C4 18.48 4 17.92 4 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Command(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 15v3a3 3 0 1 1-3-3zm0 0h6m-6 0V9m6 6v3a3 3 0 1 0 3-3zm0 0V9m0 0H9m6 0V6a3 3 0 1 1 3 3zM9 9V6a3 3 0 1 0-3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0"/><path d="M10.5 10.5L16 8l-2.5 5.5L8 16z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Confused(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm0-18a8 8 0 1 0 8 8a8.009 8.009 0 0 0-8-8Zm4 13H8v-2h8v2Zm-7.5-5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm6.993-.014a1.493 1.493 0 1 1 0-2.986a1.493 1.493 0 0 1 0 2.986Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cookie(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.152 4.087c0-.367-.052-.733-.152-1.087c4.968.005 8.994 4.04 9 9c.016 4.962-4.03 8.984-9 9c-4.97.016-8.984-4.037-9-9c1.112.236 2.27-.002 3.15-.72a3.831 3.831 0 0 0 1.42-3.006a4.5 4.5 0 0 0-.07-.781a3.277 3.277 0 0 0 3.07-.351a3.69 3.69 0 0 0 1.582-3.055m-9.15 2.915V7H3v.002zm5-4V3H8v.002zm-4 0V3H4v.002z"/><path d="M10.002 17.002V17H10v.002zm5-2V15H15v.002zm-4-3V12H11v.002zm5-2V10H16v.002zm-13-3V7H3v.002zm5-4V3H8v.002zm-4 0V3H4v.002z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coolicons(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.286 17.714A5.72 5.72 0 0 1 12 15.78a5.714 5.714 0 1 1-1.7-8.875a7.812 7.812 0 0 0-1.427 2.481a2.857 2.857 0 1 0 1.7 2.612a5.714 5.714 0 1 1 5.714 5.714v.002ZM13.43 12.1a2.857 2.857 0 1 0 0-.1v.1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 9V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C10.52 3 11.08 3 12.2 3h5.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C21 4.52 21 5.08 21 6.2v5.6c0 1.12 0 1.68-.218 2.108a2.002 2.002 0 0 1-.874.874C19.48 15 18.92 15 17.803 15H15M9 9H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 10.52 3 11.08 3 12.2v5.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h5.607c1.117 0 1.676 0 2.104-.218a2 2 0 0 0 .874-.874c.218-.428.218-.987.218-2.105V15M9 9h2.8c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105V15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2ZM4 12v6h16v-6H4Zm0-6v2h16V6H4Zm9 10H6v-2h7v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2ZM4 6v12h16V6H4Zm11.5 10a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Zm-4 0a2.5 2.5 0 0 1 0-5c.543 0 1.07.18 1.5.512a2.476 2.476 0 0 0 0 3.975a2.46 2.46 0 0 1-1.5.513Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 11v4.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V11M3 11V9m0 2h18M3 9v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105V9M3 9h18M7 15h4m10-4V9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 15a1.5 1.5 0 0 1 0-3M3 15.8V8.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v7.607c0 1.118 0 1.676-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218H6.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C3 17.48 3 16.92 3 15.8m14-2.3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 21v-3m0 0H9.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C6 16.48 6 15.92 6 14.8V6m12 12h3M6 6H3m3 0V3m4 3h4.8c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105V14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThree(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.983 20.989l-6.37-1.813l-1.42-16.033h15.615l-1.42 16.034l-6.4 1.812h-.004Zm-4.261-7.637l.216 2.867L12 17.483l3.99-1.14l.906-9.923h-9.8l.158 1.949h7.529l-.186 2.024H9.66l.178 1.912h4.6l-.272 2.62l-2.164.6l-2.2-.6l-.14-1.57h-1.94v-.003Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cupcake(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 14l-.804 5.626c-.07.487-.104.731-.222.915a1 1 0 0 1-.426.369c-.197.09-.443.09-.933.09H14m5-7h-5m5 0c1.303-.604 2-2.236 2-3.666c0-1.536-1.03-2.85-2.49-3.397a.787.787 0 0 1-.51-.729c0-1.265-1.12-2.291-2.5-2.291c-.226 0-.445.027-.654.079c-.432.107-.915.083-1.287-.145A5.833 5.833 0 0 0 10.5 3C7.462 3 5 5.257 5 8.042c0 .352-.23.674-.557.857C3.578 9.38 3 10.254 3 11.25c0 1.277.712 2.44 2 2.75m0 0l.804 5.626v.002c.07.486.104.73.222.913a1 1 0 0 0 .426.369c.197.09.443.09.933.09H10m-5-7h5m0 0h4m-4 0v7m4-7v7m0 0h-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cylinder(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 7v10c0 1.657-2.686 3-6 3s-6-1.343-6-3V7m12 0c0-1.657-2.686-3-6-3S6 5.343 6 7m12 0c0 1.657-2.686 3-6 3S6 8.657 6 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 21h-8v-6h8v6Zm-10 0H3V11h8v10Zm10-8h-8V3h8v10ZM11 9H3V3h8v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 20H11v-7h11v7ZM9 20H2v-7h7v7Zm13-9H2V4h20v7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Data(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 12v5c0 1.657-2.686 3-6 3s-6-1.343-6-3v-5m12 0V7m0 5c0 1.657-2.686 3-6 3s-6-1.343-6-3m12-5c0-1.657-2.686-3-6-3S6 5.343 6 7m12 0c0 1.657-2.686 3-6 3S6 8.657 6 7m0 5V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteColumn(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 21H9a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v6m6 5h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteRow(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 16h6m1-6V9a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20H9m-1.803-3c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875c-.121-.237-.175-.515-.199-.907m3.178 2h9.606m-9.606 0H5.6c-.56 0-.84 0-1.054-.11a.997.997 0 0 1-.437-.435C4 16.24 4 15.96 4 15.4V15h.02m0 0C4 14.686 4 14.299 4 13.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v6.606c0 .497 0 .883-.02 1.197M4.02 15h15.96m0 0c-.023.392-.077.67-.198.907a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218m3.178-2H20v.4c0 .56 0 .84-.11 1.055a.996.996 0 0 1-.436.436C19.24 17 18.96 17 18.4 17h-1.597"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopTower(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20H7m11-3h-2m-3-8H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 10.52 3 11.08 3 12.2v1.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218H13m0-8v8m0-8V7.2c0-1.12 0-1.68.218-2.108a2 2 0 0 1 .874-.874C14.52 4 15.08 4 16.2 4h1.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v9.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218h-1.606c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875c-.205-.401-.217-.919-.218-1.907m4.002-10v.002H17V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Devices(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7.5v-.3c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C7.52 4 8.08 4 9.2 4h8.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v6.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.874.875C19.48 17 18.92 17 17.803 17H10.5M3 16.8v-5.6c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 8 5.08 8 6.2 8h.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v5.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875C8.48 20 7.921 20 6.803 20h-.606c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C3 18.48 3 17.92 3 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discord(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.625 22l-2.042-1.8l-1.151-1.07L15.215 18l.5 1.759H4.947A1.954 1.954 0 0 1 3 17.8V4.957A1.954 1.954 0 0 1 4.947 3h12.731a1.954 1.954 0 0 1 1.947 1.957V22ZM15.1 13.508a3.007 3.007 0 0 1-1.715 1.128c.282.36.622.767.625.771a3.5 3.5 0 0 0 2.907-1.444a12.826 12.826 0 0 0-1.368-5.538a4.721 4.721 0 0 0-2.658-1h-.011l-.133.152a6.4 6.4 0 0 1 2.366 1.206a7.788 7.788 0 0 0-2.86-.912a7.977 7.977 0 0 0-1.918.019a.8.8 0 0 0-.134.014h-.028a7.166 7.166 0 0 0-2.156.6c-.342.157-.552.271-.561.275c-.009.004 0 0 0-.007c.738-.58 1.59-1 2.499-1.231l-.1-.114h-.006a4.723 4.723 0 0 0-2.658 1a12.729 12.729 0 0 0-1.368 5.538a3.44 3.44 0 0 0 2.9 1.444s.358-.436.637-.789a2.966 2.966 0 0 1-1.667-1.12c.007.005.106.072.266.16c.01.013.024.023.039.03a.27.27 0 0 0 .04.023a5.366 5.366 0 0 0 .738.346c.453.18.922.317 1.4.41c.419.079.844.12 1.27.12c.4 0 .798-.037 1.191-.11a6.3 6.3 0 0 0 1.377-.408a5.516 5.516 0 0 0 1.086-.563Zm-2.038-.427a1.06 1.06 0 1 1 .967-1.055c.022.558-.41 1.03-.968 1.055h.001Zm-3.468 0a1.059 1.059 0 0 1 0-2.11a.93.93 0 0 1 .668.285c.199.206.307.483.3.769a1.015 1.015 0 0 1-.969 1.056h.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotFiveXl(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotFourL(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17a5 5 0 1 1 0-10a5 5 0 0 1 0 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotOneXs(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotThreeM(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16a4 4 0 1 1 0-8a4 4 0 0 1 0 8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotTwoS(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleQuotesL(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 12v3.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437c.214.109.494.109 1.053.109H17.4c.56 0 .839 0 1.053-.109c.188-.096.341-.25.437-.437c.11-.214.11-.494.11-1.054v-1.803c0-.559 0-.838-.109-1.052a1 1 0 0 0-.437-.437C18.24 12 17.96 12 17.4 12zm0 0v-2a3 3 0 0 1 3-3M5 12v3.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C5.76 17 6.04 17 6.598 17h1.804c.559 0 .838 0 1.052-.109c.188-.096.341-.25.437-.437C10 16.24 10 15.96 10 15.4v-1.803c0-.559 0-.838-.109-1.052a1 1 0 0 0-.437-.437C9.24 12 8.96 12 8.4 12zm0 0v-2a3 3 0 0 1 3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleQuotesR(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 17a3 3 0 0 0 3-3v-2m0 0V8.598c0-.558 0-.838-.109-1.052a1 1 0 0 0-.437-.437C18.24 7 17.96 7 17.4 7h-1.8c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C14 7.76 14 8.04 14 8.6v1.8c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437c.214.109.494.109 1.053.109zM7 17a3 3 0 0 0 3-3v-2m0 0V8.598c0-.559 0-.838-.109-1.052a1 1 0 0 0-.437-.437C9.24 7 8.96 7 8.4 7H6.6c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C5 7.76 5 8.04 5 8.6v1.8c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C5.76 12 6.04 12 6.598 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoughnutChart(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.051 21.949a10 10 0 0 1-1-19.949v4.04a5.994 5.994 0 1 0 6.91 6.909H22a10.015 10.015 0 0 1-9.95 9Zm9.95-11h-4.04a5.993 5.993 0 0 0-4.91-4.909V2a10.022 10.022 0 0 1 8.95 8.948v.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 21h12M12 3v14m0 0l5-5m-5 5l-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadDone(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 19H5v-2h14v2Zm-9-4.58l-4-4l1.41-1.41L10 11.59L16.59 5L18 6.42l-8 8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadPackage(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8v8.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.987.218-2.105V8M4 8h16M4 8l1.365-2.39c.335-.585.503-.878.738-1.092c.209-.189.456-.332.723-.42C7.13 4 7.466 4 8.143 4h7.714c.676 0 1.015 0 1.318.099c.267.087.513.23.721.42c.236.213.404.506.74 1.093L20 8m-8 3v6m0 0l3-2m-3 2l-3-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragHorizontal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 14a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-6 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-6 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2m12-6a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-6 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2M6 8a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragVertical(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 18a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0m6-6a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0m6-6a1 1 0 1 0 2 0a1 1 0 0 0-2 0M8 6a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-4.88-3.455l-.005.01a8.148 8.148 0 0 0 7.939 1.033a34.534 34.534 0 0 0-1.679-5.977a12.638 12.638 0 0 0-4.546 2.773a9.68 9.68 0 0 0-1.65 2.057l-.007.011v-.01l-.018.035l-.012.02l-.007.013l-.017.036l.002-.001Zm9.659-5.5c-.492 0-.984.034-1.472.1a37.456 37.456 0 0 1 1.5 5.463a8.052 8.052 0 0 0 3.22-5.127a12.187 12.187 0 0 0-3.248-.439v.003Zm-12.962-1.09v.056a8.181 8.181 0 0 0 1.925 5.264a13.522 13.522 0 0 1 6.772-5.375l.119-.034c-.177-.391-.343-.74-.509-1.066a31.124 31.124 0 0 1-8.195 1.155h-.112Zm12.649-.606a21.36 21.36 0 0 1 3.7.338a8.105 8.105 0 0 0-1.692-4.659a11.791 11.791 0 0 1-4.4 3.084l.029-.012l-.031.013a23.484 23.484 0 0 1 .548 1.233l.041.1c.6-.066 1.202-.099 1.805-.097ZM8.377 4.677a8.178 8.178 0 0 0-4.33 5.414c2.33-.036 4.65-.33 6.915-.878h.028L11 9.2l.054-.013h.006l.04-.008h.016l.021-.006l.04-.01l.023-.01l.022-.006h.02a51.796 51.796 0 0 0-2.866-4.47h.001ZM12 3.839a8.437 8.437 0 0 0-1.631.155a45.662 45.662 0 0 1 2.879 4.528a9.666 9.666 0 0 0 4.036-2.746A8.192 8.192 0 0 0 12 3.839Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.006 20.5L7 17.311l5-3.186l5 3.186l-5 3.188l.006.001Zm5-4.251l-5-3.187l5-3.187l-5-3.186l5-3.189l5 3.189l-5 3.186l5 3.187l-5 3.187ZM7 16.249l-5-3.187l5-3.187l-5-3.186L7 3.5l5 3.189l-5 3.185l5 3.187l-5 3.188Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DummyCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6a6 6 0 1 0 0 12a6 6 0 0 0 0-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DummyCircleSmall(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8a4 4 0 1 0 0 8a4 4 0 0 0 0-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DummySquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9.2v5.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h5.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.987.218-2.104V9.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C16.48 6 15.92 6 14.8 6H9.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C6 7.52 6 8.08 6 9.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DummySquareSmall(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 11.2v1.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h1.606c1.118 0 1.677 0 2.104-.218c.377-.191.684-.498.875-.874c.218-.428.218-.986.218-2.104v-1.607c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C14.48 8 13.92 8 12.8 8h-1.6c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C8 9.52 8 10.08 8 11.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.42 20.579a1 1 0 0 1-.737-.326a.988.988 0 0 1-.263-.764l.245-2.694L14.983 5.481l3.537 3.536L7.205 20.33l-2.694.245a.95.95 0 0 1-.091.004ZM19.226 8.31L15.69 4.774l2.121-2.121a1 1 0 0 1 1.415 0l2.121 2.121a1 1 0 0 1 0 1.415l-2.12 2.12l-.001.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditPencilLineOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16M4 20v-4l8-8M4 20h4l8-8m-4-4l2.869-2.869l.001-.001c.395-.395.593-.593.821-.667a1 1 0 0 1 .618 0c.228.074.425.272.82.666l1.74 1.74c.396.396.594.594.668.822a1 1 0 0 1 0 .618c-.074.228-.272.426-.668.822h0L16 12.001m-4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditPencilLineTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16M4 20v-4L14.869 5.131l.001-.001c.395-.395.593-.593.821-.667a1 1 0 0 1 .618 0c.228.074.425.272.82.666l1.74 1.74c.396.396.594.594.668.822a1 1 0 0 1 0 .618c-.074.228-.272.426-.668.822h0L8 20.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditPencilOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 8l-8 8v4h4l8-8m-4-4l2.869-2.869l.001-.001c.395-.395.593-.593.821-.667a1 1 0 0 1 .618 0c.228.074.425.272.82.666l1.74 1.74c.396.396.594.594.668.822a1 1 0 0 1 0 .618c-.074.228-.272.426-.668.822h0L16 12.001m-4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditPencilTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v4h4L18.869 9.131h0c.396-.396.594-.594.668-.822a1 1 0 0 0 0-.618c-.074-.228-.272-.426-.668-.822l-1.74-1.74c-.395-.394-.592-.592-.82-.666a1 1 0 0 0-.618 0c-.228.074-.426.272-.82.667l-.002.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorIcon(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-1-7v2h2v-2h-2Zm0-8v6h2V7h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.99 22C6.468 21.994 1.996 17.515 2 11.993C2.004 6.472 6.482 1.998 12.003 2C17.525 2.002 22 6.478 22 12c-.003 5.525-4.485 10.002-10.01 10ZM4 12.172A8 8 0 1 0 4 12v.172ZM13 17h-2v-2h2v2Zm0-4h-2V7h2v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exit(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 15l3-3m0 0l-3-3m3 3H4m0-4.752V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v9.607c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C4 18.48 4 17.92 4 16.8v-.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19H5v-5m9-9h5v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 5H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 6.52 5 7.08 5 8.2v7.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.987.218-2.105V14m1-5V4m0 0h-5m5 0l-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.002 12.002a10.005 10.005 0 0 0 8.437 9.879v-6.989H7.902v-2.89h2.54v-2.2a3.528 3.528 0 0 1 3.773-3.9c.75.012 1.5.079 2.24.2v2.459h-1.264a1.446 1.446 0 0 0-1.628 1.563v1.878h2.771l-.443 2.891h-2.328v6.988a10 10 0 1 0-11.561-9.879Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForward(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 18V6l8.5 6l-8.5 6Zm-9 0V6l8.5 6L4 18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastRewind(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20 18l-8.5-6L20 6v12Zm-9 0l-8.5-6L11 6v12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Figma(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 15h3m-3 0a3 3 0 1 0 3 3v-3m-3 0a3 3 0 1 1 0-6m3 6v-3M9 9h3M9 9a3 3 0 0 1 0-6h3m0 6v3m0-3V3m0 6h3m-3 3a3 3 0 1 0 3-3m-3 3a3 3 0 0 1 3-3m-3-6h3a3 3 0 1 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18v-3m0 0v-3m0 3H9m3 0h3M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9m-6-6c.286.003.466.014.639.055c.204.05.399.13.578.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802m0 0H19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileArchive(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h1v2h2v2H7v2h2v2H7v7h4v-5H9v-2h2V8H9V6h2V4H9V2h4a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2ZM13 4v5h5l-5-5Zm-3 12H8v-2h2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileBlank(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9m-6-6c.286.003.466.014.639.055c.204.05.399.13.578.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802m0 0H19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileBlankFill(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2ZM13 4v5h5l-5-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileBlankOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.09.09 0 0 1 .032.006a.15.15 0 0 0 .03.006c.088.006.175.023.259.051l.028.009a.988.988 0 0 1 .359.228l6 6a.987.987 0 0 1 .2.293c.01.022.017.045.025.068l.009.026c.028.083.044.17.049.258l.007.027C20 8.982 20 8.991 20 9v11a2 2 0 0 1-2 2ZM6 4v16h12V10h-5a1 1 0 0 1-1-1V4H6Zm8 1.414V8h2.586L14 5.414Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 15l2 2l4-4M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9m-6-6c.286.003.466.014.639.055c.204.05.399.13.578.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802m0 0H19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 17l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9m-6-6c.286.003.466.014.639.055c.204.05.399.13.578.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802m0 0H19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCode(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9m-6-6c.286.003.466.014.639.055c.204.05.399.13.578.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802m0 0H19m-5 4l2 2l-2 2m-4 0l-2-2l2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCss(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2Zm-3.973-4.6c-.008.297.061.59.2.852c.132.235.328.428.566.555c.247.13.521.196.8.193c.375.02.742-.109 1.024-.357A1.31 1.31 0 0 0 17 17.65a1.478 1.478 0 0 0-.3-.953a2.259 2.259 0 0 0-.962-.632a1.723 1.723 0 0 1-.529-.334a.528.528 0 0 1-.149-.364a.6.6 0 0 1 .134-.408a.47.47 0 0 1 .373-.154a.455.455 0 0 1 .389.181a.84.84 0 0 1 .137.517H17a1.684 1.684 0 0 0-.177-.782a1.253 1.253 0 0 0-.5-.531a1.586 1.586 0 0 0-1.778.184a1.28 1.28 0 0 0-.4.986c-.005.32.1.632.3.882c.269.294.605.517.981.649c.19.072.364.18.513.317a.632.632 0 0 1 .143.448c0 .361-.161.544-.479.544a.6.6 0 0 1-.5-.19a.984.984 0 0 1-.153-.612l-.923.002Zm-3.443 0c-.008.297.061.59.2.852c.132.235.328.428.566.555c.247.13.521.196.8.193c.375.02.742-.109 1.024-.357c.26-.263.397-.624.376-.993a1.478 1.478 0 0 0-.3-.953a2.259 2.259 0 0 0-.962-.632a1.723 1.723 0 0 1-.529-.334a.532.532 0 0 1-.149-.364a.6.6 0 0 1 .134-.408a.47.47 0 0 1 .374-.154a.455.455 0 0 1 .389.181c.101.153.15.334.137.517h.913a1.7 1.7 0 0 0-.177-.782a1.261 1.261 0 0 0-.5-.531a1.457 1.457 0 0 0-.745-.19c-.38-.013-.75.121-1.033.374a1.28 1.28 0 0 0-.4.986c-.006.32.1.632.3.882c.268.294.605.517.98.649c.19.072.365.18.514.317a.632.632 0 0 1 .143.448c0 .361-.161.544-.48.544a.6.6 0 0 1-.5-.19a.984.984 0 0 1-.159-.61h-.916ZM8.6 14a1.41 1.41 0 0 0-1.176.545c-.31.453-.458.996-.424 1.544v.819a2.56 2.56 0 0 0 .408 1.551A1.4 1.4 0 0 0 8.587 19c.421.025.833-.13 1.133-.427c.27-.32.421-.724.429-1.143v.053v-.173h-.908c.011.24-.043.478-.156.69a.582.582 0 0 1-.5.187a.55.55 0 0 1-.52-.271a2.3 2.3 0 0 1-.149-1V16c-.015-.31.04-.62.162-.905a.546.546 0 0 1 .52-.278a.557.557 0 0 1 .495.2c.112.219.163.463.149.708h.918a2.115 2.115 0 0 0-.439-1.277A1.406 1.406 0 0 0 8.6 14ZM13 4v5h5l-5-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDocument(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17h6m-6-3h6M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9.326L18.999 9M13 3c.286.003.466.014.638.055c.204.05.4.13.579.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownload(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12v6m0 0l3-2m-3 2l-3-2m4-13H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9m-6-6c.286.003.466.014.639.055c.204.05.399.13.578.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802m0 0H19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileEdit(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 11V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C7.52 3 8.08 3 9.2 3H14m6 6v8.804c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.874.874C18.48 21 17.92 21 16.803 21H13m7-12c-.004-.285-.014-.466-.056-.639a1.993 1.993 0 0 0-.24-.578c-.123-.202-.295-.374-.641-.72l-3.125-3.125c-.346-.346-.52-.52-.721-.643a1.999 1.999 0 0 0-.578-.24c-.173-.041-.353-.052-.639-.054M20 9h-2.803c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C14 7.48 14 6.92 14 5.8V3M9 14l2 2m-7 5v-2.5l7.5-7.5l2.5 2.5L6.5 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileFind(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7c.265 0 .52.105.707.293l6 6A.991.991 0 0 1 20 9v11a2 2 0 0 1-2 2ZM6 4v16h10.586l-2.566-2.566A3.941 3.941 0 0 1 12 18a4.044 4.044 0 1 1 3.434-1.98L18 18.588V9.414L12.586 4H6Zm6 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileHtml(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2Zm-1.739-7v4H18.5v-.67h-1.426V15h-.813Zm-1.4 1.258l-.074 1.659V19h.815v-4h-1.057l-.757 2.894l-.76-2.894h-1.061v4h.813v-1.083l-.08-1.667l.8 2.75h.554l.807-2.742ZM8.712 15v.673h.978V19h.815v-3.327h.995V15H8.712Zm-2.4 2.289h1.2V19h.81v-4h-.81v1.618h-1.2V15H5.5v4h.812v-1.711ZM13 4v5h5l-5-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImage(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2Zm-9-6l-3 4h12l-5-7l-3 4l-1-1Zm-.5-5a1.5 1.5 0 1 0 1.5 1.578v.29v-.368A1.5 1.5 0 0 0 8.5 11ZM13 4v5h5l-5-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileJpg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2Zm-2.539-8a1.4 1.4 0 0 0-1.188.545a2.552 2.552 0 0 0-.406 1.547v.876c-.029.54.128 1.074.443 1.514c.3.35.746.542 1.207.518c.286.003.57-.05.836-.154h.008h-.006a.189.189 0 0 1 .029-.014h.005l.016-.008l.021-.01c.218-.1.412-.247.568-.43v-2.011H15.45v.742h.637v.893l-.08.06a.753.753 0 0 1-.436.117a.661.661 0 0 1-.607-.3a1.962 1.962 0 0 1-.183-.969v-.883a1.985 1.985 0 0 1 .173-.925a.574.574 0 0 1 .535-.292a.55.55 0 0 1 .446.171c.12.19.182.408.18.632H17a1.933 1.933 0 0 0-.433-1.227c-.297-.28-.7-.421-1.106-.392ZM7 17.5c-.021.396.1.787.344 1.1c.244.272.598.419.963.4c.362.009.709-.149.941-.427a1.7 1.7 0 0 0 .365-1.137v-3.369H8.7v3.322c0 .528-.133.8-.4.8s-.393-.232-.393-.689H7Zm3.327-3.429v4.866h.91v-1.715h.607c.405.022.802-.128 1.09-.414a1.58 1.58 0 0 0 .392-1.13a1.69 1.69 0 0 0-.4-1.164c-.27-.3-.661-.464-1.065-.447l-1.534.004ZM13 4v5h5l-5-5Zm-1.138 12.4h-.625v-1.515h.635c.16 0 .307.084.39.221c.105.174.156.375.146.578a.866.866 0 0 1-.142.535a.483.483 0 0 1-.404.181Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileJs(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2Zm-5.884-4.6c-.006.296.063.59.2.852a1.4 1.4 0 0 0 .55.555c.237.129.502.195.772.193c.366.019.724-.11.994-.357A1.33 1.33 0 0 0 15 17.65a1.506 1.506 0 0 0-.3-.95a2.178 2.178 0 0 0-.932-.632a1.67 1.67 0 0 1-.514-.334a.54.54 0 0 1-.145-.364a.6.6 0 0 1 .13-.408a.449.449 0 0 1 .362-.154a.434.434 0 0 1 .378.181a.859.859 0 0 1 .132.517H15a1.738 1.738 0 0 0-.172-.782a1.237 1.237 0 0 0-.486-.531A1.382 1.382 0 0 0 13.62 14a1.407 1.407 0 0 0-1 .374a1.3 1.3 0 0 0-.392.986a1.4 1.4 0 0 0 .29.882c.258.292.585.515.951.649c.186.072.356.18.5.317c.101.127.15.287.139.448c0 .361-.157.544-.465.544a.578.578 0 0 1-.483-.19a1.012 1.012 0 0 1-.148-.612l-.896.002ZM9 17.5c-.022.395.097.784.335 1.1c.235.27.581.418.939.4c.356.008.695-.15.918-.427c.25-.325.376-.728.356-1.137v-3.369h-.887v3.322c0 .528-.13.8-.387.8s-.383-.232-.383-.689H9ZM13 4v5h5l-5-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.09.09 0 0 1 .032.006a.15.15 0 0 0 .03.006c.088.006.175.023.259.051l.028.009a.988.988 0 0 1 .359.228l6 6a.987.987 0 0 1 .2.293c.01.022.017.045.025.068l.009.026c.028.083.044.17.049.258l.007.027C20 8.982 20 8.991 20 9v11a2 2 0 0 1-2 2ZM6 4v16h12V10h-5a1 1 0 0 1-1-1V4H6Zm8 1.414V8h2.586L14 5.414ZM15 16H9v-2h6v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileNew(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066v.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006c.005.006.01.013.013.02v11a2 2 0 0 1-2 2ZM6 4v16h12V10h-5a1 1 0 0 1-1-1V4H6Zm8 1.414V8h2.586L14 5.414ZM13 18h-2v-2H9v-2h2v-2h2v2h2v2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdf(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2Zm-3.576-8v5h.94v-2.04h1.46v-.838h-1.46v-1.281H17V14h-2.576Zm-3.7 0v5h1.206a1.67 1.67 0 0 0 1.332-.56a2.3 2.3 0 0 0 .486-1.549v-.81a2.287 2.287 0 0 0-.5-1.526c-.325-.37-.8-.574-1.293-.555h-1.231ZM7 14v5h.94v-1.759h.626c.418.023.826-.132 1.124-.426a1.62 1.62 0 0 0 .41-1.16a1.725 1.725 0 0 0-.412-1.194A1.4 1.4 0 0 0 8.585 14H7Zm6-10v5h5l-5-5Zm-1.054 14.162h-.282v-3.321h.342a.716.716 0 0 1 .62.292c.147.303.21.64.182.976v.869c.022.32-.047.64-.2.921a.765.765 0 0 1-.662.263ZM8.585 16.4h-.646v-1.559h.655a.475.475 0 0 1 .4.227c.108.179.16.385.15.594a.89.89 0 0 1-.147.55a.5.5 0 0 1-.412.188Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePng(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2Zm-2.458-8a1.3 1.3 0 0 0-1.126.545a2.66 2.66 0 0 0-.384 1.547v.876a2.479 2.479 0 0 0 .419 1.514c.277.348.705.54 1.149.518c.271.003.54-.05.791-.154l.021-.009a1.55 1.55 0 0 0 .588-.455v-2.009h-1.469v.742h.6v.893l-.076.06a.689.689 0 0 1-.414.117a.62.62 0 0 1-.575-.3a2.063 2.063 0 0 1-.173-.969v-.883c-.016-.317.04-.633.165-.925a.537.537 0 0 1 .507-.292c.16-.013.317.05.423.171c.113.191.171.41.17.632H17a2 2 0 0 0-.41-1.227A1.34 1.34 0 0 0 15.543 14h-.001Zm-4.258 1.746l1.259 3.187h.857v-4.866h-.855v3.195l-1.263-3.195h-.862v4.866h.862v-3.187h.002ZM7 14.067v4.866h.862v-1.711h.575c.388.02.766-.131 1.032-.414c.262-.317.395-.72.373-1.13a1.747 1.747 0 0 0-.379-1.164a1.253 1.253 0 0 0-1.009-.447H7ZM13 4v5h5l-5-5ZM8.454 16.4h-.592v-1.515h.6a.432.432 0 0 1 .37.221c.1.176.148.376.138.578a.9.9 0 0 1-.135.535a.448.448 0 0 1-.381.181Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 15h6M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9m-6-6c.286.003.466.014.639.055c.204.05.399.13.578.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802m0 0H19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSearch(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 18l-1.658-1.657M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9.326L18.999 9M13 3c.286.003.466.014.638.055c.204.05.4.13.579.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802M11.5 17a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSvg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h7a.104.104 0 0 1 .027 0h.006a.15.15 0 0 0 .029.006c.088.006.175.023.259.051h.042a.421.421 0 0 1 .052.043a.988.988 0 0 1 .293.2l6 6a.987.987 0 0 1 .2.293a.735.735 0 0 1 .023.066l.01.028c.028.083.044.17.049.258a.1.1 0 0 0 .007.029v.006A.112.112 0 0 1 20 9v11a2 2 0 0 1-2 2Zm-2.5-8c-.454-.024-.89.18-1.161.545a2.6 2.6 0 0 0-.4 1.547v.876a2.426 2.426 0 0 0 .432 1.514c.29.35.728.541 1.181.518c.28.003.557-.05.816-.154l.013-.005l-.016.007c.245-.101.46-.26.63-.464v-2.011h-1.509v.742h.623v.893l-.078.06a.723.723 0 0 1-.427.117a.641.641 0 0 1-.592-.3a2.01 2.01 0 0 1-.178-.969v-.883a2.018 2.018 0 0 1 .166-.926a.557.557 0 0 1 .523-.292a.531.531 0 0 1 .435.171c.117.19.177.41.175.632H17a1.96 1.96 0 0 0-.423-1.227A1.4 1.4 0 0 0 15.5 14ZM7 17.4c-.006.296.063.59.2.852c.126.234.317.426.55.555c.237.129.503.195.773.193c.366.02.724-.11.994-.357a1.33 1.33 0 0 0 .366-.993c.013-.34-.09-.676-.293-.95a2.186 2.186 0 0 0-.934-.632a1.652 1.652 0 0 1-.513-.334a.536.536 0 0 1-.143-.367a.608.608 0 0 1 .13-.408a.45.45 0 0 1 .363-.154c.148-.008.29.06.377.181A.853.853 0 0 1 9 15.5h.886a1.739 1.739 0 0 0-.172-.782a1.245 1.245 0 0 0-.487-.531A1.382 1.382 0 0 0 8.505 14a1.409 1.409 0 0 0-1 .374a1.3 1.3 0 0 0-.393.986a1.4 1.4 0 0 0 .291.882c.258.293.585.516.951.649c.186.072.356.18.5.317c.101.127.15.287.139.448c0 .361-.157.544-.465.544a.58.58 0 0 1-.484-.19a1.012 1.012 0 0 1-.148-.612L7 17.4Zm3.182-3.332l1.26 4.866h.925l1.269-4.866h-.991L11.9 17.6l-.737-3.529l-.981-.003ZM13 4v5h5l-5-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUpload(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18v-6m0 0l-3 2m3-2l3 2M13 3H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 4.52 5 5.08 5 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h7.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9m-6-6c.286.003.466.014.639.055c.204.05.399.13.578.24c.202.124.375.297.72.643l3.126 3.125c.346.346.518.518.642.72c.11.18.19.374.24.578c.04.173.051.354.054.639M13 3v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h2.802m0 0H19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Files(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6h.337c.244 0 .367 0 .482.028c.102.024.2.065.29.12c.1.061.187.148.36.32l3.063 3.063c.172.173.258.26.32.36c.055.09.096.188.12.29c.028.114.028.235.028.474V18M9 6H4.6c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C3 6.76 3 7.04 3 7.6v11.8c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C3.76 21 4.039 21 4.598 21h7.803c.559 0 .84 0 1.053-.109a.999.999 0 0 0 .437-.437C14 20.24 14 19.96 14 19.4V18M9 6v3.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437c.214.109.493.109 1.052.109H14m-4-5V4.6c0-.56 0-.84.109-1.054a1 1 0 0 1 .437-.437C10.76 3 11.04 3 11.6 3H16m0 0h.337c.244 0 .367 0 .482.028c.102.024.2.065.29.12c.1.061.187.148.36.32l3.063 3.063c.172.173.258.26.32.36c.055.09.096.188.12.29c.028.114.028.235.028.474V16.4c0 .56 0 .84-.11 1.054a.998.998 0 0 1-.435.437C20.24 18 19.96 18 19.402 18H14m2-15v3.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C16.76 8 17.039 8 17.598 8H21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 5.6c0-.56 0-.84-.11-1.054a.998.998 0 0 0-.436-.437C19.24 4 18.96 4 18.4 4H5.6c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C4 4.76 4 5.04 4 5.6v.737c0 .245 0 .367.028.482a1 1 0 0 0 .12.29c.061.1.148.187.32.36l5.063 5.062c.173.173.26.26.321.36c.055.09.096.188.12.29c.028.114.028.235.028.474v4.756c0 .857 0 1.286.18 1.544a1 1 0 0 0 .674.416c.311.046.695-.145 1.461-.529l.8-.4c.322-.16.482-.24.599-.36a1 1 0 0 0 .231-.374c.055-.158.055-.338.055-.697v-4.348c0-.245 0-.367.028-.482a.998.998 0 0 1 .12-.29c.06-.1.147-.186.317-.356l.004-.004l5.063-5.062c.172-.173.258-.26.32-.36a.994.994 0 0 0 .12-.29C20 6.706 20 6.584 20 6.345z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterOff(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 4h5.4c.56 0 .84 0 1.055.109a1 1 0 0 1 .436.437C20 4.76 20 5.04 20 5.6v.745c0 .24 0 .36-.027.474c-.025.102-.066.2-.12.29c-.062.1-.149.187-.322.36L18 9M7.5 4H5.6c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C4 4.76 4 5.04 4 5.6v.737c0 .245 0 .367.028.482a1 1 0 0 0 .12.29c.061.1.148.187.32.36l5.063 5.062c.173.173.26.26.321.36c.055.09.096.188.12.29c.028.114.028.235.028.474v4.756c0 .857 0 1.286.18 1.544a1 1 0 0 0 .674.416c.311.046.695-.145 1.461-.529l.8-.4c.322-.16.482-.24.599-.36a.999.999 0 0 0 .231-.374c.055-.158.055-.338.055-.697v-4.348c0-.245 0-.367.028-.482a.998.998 0 0 1 .12-.29c.061-.1.147-.186.319-.358l.002-.002l1.031-1.03m0 0L5 1m10.5 10.5L19 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterOffOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.414.414L7 1.828L8.172 3H4a1 1 0 0 0-1 1v2.59c0 .523.213 1.037.583 1.407L9 13.414V21a1.001 1.001 0 0 0 1.447.895l4-2c.339-.17.553-.516.553-.895v-5.586l1.793-1.793l2.935 2.935l1.414-1.414L8.414.414Zm6.965 9.793l-2.086 2.086A.996.996 0 0 0 13 13v5.382l-2 1V13a.996.996 0 0 0-.293-.707L5 6.59V5h5.172l5.207 5.207ZM20 3h-6.172l2 2h3.173l.002 1.583l-.796.796l1.414 1.414l.796-.796c.37-.37.583-.884.583-1.407V4a1 1 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4a1 1 0 0 0-1 1v2.59c0 .523.213 1.037.583 1.407L9 13.414V21a1.001 1.001 0 0 0 1.447.895l4-2c.339-.17.553-.516.553-.895v-5.586l5.417-5.417c.37-.37.583-.884.583-1.407V4a1 1 0 0 0-1-1Zm-6.707 9.293A.996.996 0 0 0 13 13v5.382l-2 1V13a.996.996 0 0 0-.293-.707L5 6.59V5h14.001l.002 1.583l-5.71 5.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstAid(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-5h5a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-5V4a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v5H4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstPage(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.444 18.01L10.432 12l6.012-6.01l1.415 1.414l-4.6 4.6l4.6 4.6l-1.415 1.406Zm-8.3-.01h-2V6h2v12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 21v-5.313m0 0c5.818-4.55 10.182 4.55 16 0V4.313c-5.818 4.55-10.182-4.55-16 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagFill(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.66 4.3a1 1 0 0 0-.98-.8H5.5a1 1 0 0 0-1 1v15a1 1 0 1 0 2 0v-6h5.6l.24 1.2c.09.468.503.805.98.8h5.18a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-4.6l-.24-1.2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 20.5a1 1 0 0 1-1-1v-15a1 1 0 0 1 1-1h6.38a1 1 0 0 1 .9.55l.72 1.45h5a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-5.39a1 1 0 0 1-.89-.55l-.72-1.45h-5v6a1 1 0 0 1-1 1Zm1-15v6h6l1 2h4v-6h-5l-1-2h-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v-3m0 0v-3m0 3H9m3 0h3M3 6v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l-4 4l-2-2M3 6v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 15l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2M3 6v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCode(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 11l2 2l-2 2m-4 0l-2-2l2-2M3 6v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderDocument(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 15h6m-6-3h6M3 6v10.8c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.607c1.118 0 1.677 0 2.104-.218c.376-.192.682-.498.874-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2.001 2.001 0 0 0-.875-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderDownload(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l3-2m-3 2l-3-2M3 6v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderEdit(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 9.25V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H3m0 0v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874C4.52 20 5.08 20 6.2 20H7M3 6v-.4c0-.56 0-.84.109-1.054a1 1 0 0 1 .437-.437C3.76 4 4.04 4 4.6 4h4.737c.245 0 .367 0 .482.028a1 1 0 0 1 .29.12c.1.061.187.148.36.32L12 6m4 8l2 2m-7 5v-2.5l7.5-7.5l2.5 2.5l-7.5 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6a1 1 0 0 1 .707.293L12.414 5H20a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2ZM4 7v12h16V7H4Zm12 7H8v-2h8v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.803 20h1.99c.433 0 .649 0 .83-.074c.161-.066.302-.172.41-.308c.12-.155.18-.362.299-.778l1.086-3.8c.198-.693.296-1.04.218-1.313a1 1 0 0 0-.435-.577c-.228-.141-.561-.15-1.2-.15m-3.198 7H6.197m11.606 0c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874C21 18.48 21 17.92 21 16.8V13M6.197 20H5.12c-.72 0-1.08 0-1.322-.15a1 1 0 0 1-.436-.577a.71.71 0 0 1-.025-.16m2.86.887c-1.118 0-1.678 0-2.105-.218a1.999 1.999 0 0 1-.754-.67M21 13H5l-1.417 4.96l-.002.007c-.16.56-.255.893-.243 1.145M21 13V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H12M3.338 19.112a1.998 1.998 0 0 1-.12-.204C3 18.48 3 17.92 3 16.8V6m0 0h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6a1 1 0 0 1 .707.293L12.414 5H20a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2ZM4 7v12h16V7H4Zm9 10h-2v-3H8v-2h3V9h2v3h3v2h-3v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 13h6M3 6v10.8c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.607c1.118 0 1.677 0 2.104-.218c.376-.192.682-.498.874-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2.001 2.001 0 0 0-.875-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSearch(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6m3 10l-1.658-1.657M11.5 15a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderUpload(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v-6m0 0l-3 2m3-2l3 2M3 6v10.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874C21 18.48 21 17.92 21 16.8V9.2c0-1.12 0-1.68-.218-2.108a2 2 0 0 0-.874-.874C19.48 6 18.92 6 17.8 6H12M3 6h9M3 6a2 2 0 0 1 2-2h3.675c.489 0 .734 0 .964.055c.204.05.399.13.578.24c.202.124.375.297.72.643L12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folders(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 11v8.4c0 .56 0 .84.109 1.053a1 1 0 0 0 .437.438C3.76 21 4.039 21 4.598 21h10.803c.56 0 .84 0 1.053-.109a.998.998 0 0 0 .437-.437C17 20.24 17 19.96 17 19.4V15M3 11h7m-7 0v-.4c0-.56 0-.84.109-1.054a1 1 0 0 1 .437-.437C3.76 9 4.04 9 4.6 9H7m3 2h5.4c.56 0 .842 0 1.056.109a.994.994 0 0 1 .435.437C17 11.76 17 12.04 17 12.6V15m-7-4L8.925 9.618c-.176-.227-.265-.34-.375-.422a.994.994 0 0 0-.324-.159C8.094 9 7.949 9 7.662 9H7m0-4h12.4c.56 0 .842 0 1.056.109a.994.994 0 0 1 .435.437C21 5.76 21 6.04 21 6.6v6.8c0 .56 0 .84-.11 1.054a.997.997 0 0 1-.435.437c-.214.109-.494.109-1.053.109h-2.4M7 5v4m0-4v-.4c0-.56 0-.84.109-1.054a1 1 0 0 1 .437-.437C7.76 3 8.04 3 8.6 3h3.062c.287 0 .43 0 .562.037a.999.999 0 0 1 .326.159c.11.081.199.195.375.422L14 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 18L8 6L3 18m8-4H5m16 4v-3m0 0v-3m0 3a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12V7l9 5l-9 5zm0 0l-9 5V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5.5V8m0-2.5A2.5 2.5 0 1 1 14.5 8M12 5.5A2.5 2.5 0 1 0 9.5 8M12 8h2.5M12 8H9.5M12 8v6m2.5-6h3.3c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105V14M9.5 8H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 9.52 3 10.08 3 11.2V14m0 0v2.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218H12m-9-6h9m0 0v6m0-6h9m-9 6h5.803c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.026 2a9.975 9.975 0 0 0-3.153 19.439c.5.09.679-.217.679-.481c0-.237-.008-.865-.011-1.7c-2.775.6-3.361-1.338-3.361-1.338a2.635 2.635 0 0 0-1.107-1.459c-.9-.619.069-.605.069-.605c.64.088 1.205.467 1.527 1.028a2.124 2.124 0 0 0 2.9.829c.046-.506.272-.979.635-1.334c-2.214-.251-4.542-1.107-4.542-4.93a3.865 3.865 0 0 1 1.026-2.671a3.588 3.588 0 0 1 .1-2.64s.837-.269 2.742 1.021a9.439 9.439 0 0 1 4.992 0c1.906-1.291 2.742-1.021 2.742-1.021c.37.835.405 1.78.1 2.64a3.84 3.84 0 0 1 1.024 2.675c0 3.833-2.33 4.675-4.552 4.922c.48.49.725 1.162.675 1.846c0 1.334-.012 2.41-.012 2.737c0 .267.178.577.687.479A9.975 9.975 0 0 0 12.026 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h5m-5 0a9 9 0 0 0 9 9m-9-9a9 9 0 0 1 9-9m-4 9h8m-8 0c0 4.97 1.79 9 4 9m-4-9c0-4.97 1.79-9 4-9m4 9h5m-5 0c0-4.97-1.79-9-4-9m4 9c0 4.97-1.79 9-4 9m9-9a9 9 0 0 0-9-9m9 9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.956 10.356v3.451h4.792c-.446 2.193-2.313 3.453-4.792 3.453a5.28 5.28 0 0 1 0-10.559a5.166 5.166 0 0 1 3.29 1.178l2.6-2.6a8.93 8.93 0 1 0-5.89 15.635c4.467 0 8.529-3.249 8.529-8.934a7.468 7.468 0 0 0-.2-1.625l-8.329.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 20h-4v-4h4v4Zm-6 0h-4v-4h4v4Zm-6 0H4v-4h4v4Zm12-6h-4v-4h4v4Zm-6 0h-4v-4h4v4Zm-6 0H4v-4h4v4Zm12-6h-4V4h4v4Zm-6 0h-4V4h4v4ZM8 8H4V4h4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridBig(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 19h-6v-6h6v6Zm-8 0H5v-6h6v6Zm8-8h-6V5h6v6Zm-8 0H5V5h6v6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridBigRound(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 19a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm-8 0a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm8-8a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm-8 0a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridHorizontal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 17h-4v-4h4v4Zm-6 0h-4v-4h4v4Zm-6 0H4v-4h4v4Zm12-6h-4V7h4v4Zm-6 0h-4V7h4v4Zm-6 0H4V7h4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridHorizontalRound(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 17h-4v-4h4v4Zm-6 0h-4v-4h4v4Zm-6 0H4v-4h4v4Zm12-6h-4V7h4v4Zm-8 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-4 0H4V7h4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridRound(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm12-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm12-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM6 8a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridSmall(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 17h-4v-4h4v4Zm-6 0H7v-4h4v4Zm6-6h-4V7h4v4Zm-6 0H7V7h4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridSmallRound(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 17h-4v-4h4v4Zm-6 0H7v-4h4v4Zm6-6h-4V7h4v4Zm-8 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridVertical(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 20h-4v-4h4v4Zm-6 0H7v-4h4v4Zm6-6h-4v-4h4v4Zm-6 0H7v-4h4v4Zm6-6h-4V4h4v4Zm-6 0H7V4h4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridVerticalRound(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 20h-4v-4h4v4Zm-6 0H7v-4h4v4Zm6-6h-4v-4h4v4Zm-8 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm8-6h-4V4h4v4Zm-6 0H7V4h4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 7a4 4 0 1 1-8 0a4 4 0 0 1 8 0ZM9 7a2 2 0 1 0-4 0a2 2 0 0 0 4 0Zm12 4.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm-2 0a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0ZM10 21v-4a3 3 0 1 0-6 0v4H2v-4a5 5 0 0 1 10 0v4h-2Zm10-.5v.5h2v-.5a4.5 4.5 0 1 0-9 0v.5h2v-.5a2.5 2.5 0 0 1 5 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3a5 5 0 1 0 0 10A5 5 0 0 0 9 3ZM6 8a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm10.908.218A2 2 0 0 0 16 8V6a4 4 0 1 1-2.357 7.232l1.178-1.616a2 2 0 1 0 2.087-3.398ZM19.998 21A3.999 3.999 0 0 0 16 17.002V15a6.001 6.001 0 0 1 6 6h-2.002ZM16 21h-2a5 5 0 0 0-10 0H2a7 7 0 1 1 14 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hamburger(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18H3v-2h18v2Zm0-5H3v-2h18v2Zm0-5H3V6h18v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HamburgerLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17h18M3 12h18M3 7h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HamburgerMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h14M5 12h14M5 7h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Handbag(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8h-.29c-.962 0-1.443 0-1.834.175a2 2 0 0 0-.847.718c-.237.356-.316.83-.475 1.78l-.933 5.6v.003c-.214 1.28-.32 1.921-.135 2.42c.163.439.475.806.88 1.039C4.828 20 5.478 20 6.777 20h10.445c1.3 0 1.95 0 2.413-.265a2 2 0 0 0 .88-1.038c.184-.5.078-1.14-.136-2.42v-.003l-.934-5.6c-.158-.95-.237-1.425-.474-1.781a2 2 0 0 0-.847-.718C17.734 8 17.252 8 16.289 8H16M8 8h8M8 8a4 4 0 1 1 8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Happy(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm0-18a8 8 0 1 0 8 8a8.009 8.009 0 0 0-8-8Zm0 14a4.837 4.837 0 0 1-4-2a6.3 6.3 0 0 1-1-2h10v.008A6.422 6.422 0 0 1 16 16a4.838 4.838 0 0 1-4 2Zm-3.5-6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm6.993-.014a1.493 1.493 0 1 1 0-2.986a1.493 1.493 0 0 1 0 2.986Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heading(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 5v7m0 0v7m0-7h10m0-7v7m0 0v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingHfive(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 9h-4l-1.25 5.016a2.998 2.998 0 0 1 3.555-.717a3 3 0 1 1-3.555 4.685M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingHfour(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 9l-2.5 8H20m0 0h1m-1 0v-3m0 3v2M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingHone(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 10l3-1v10M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingHsix(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.402 14.525A2.974 2.974 0 0 0 16.5 18.6a3.01 3.01 0 0 0 4.099-1.092a2.973 2.973 0 0 0-1.099-4.075a3.01 3.01 0 0 0-4.098 1.092m0 0L19 8M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingHthree(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 9h6l-4 4h1a3 3 0 1 1-2.83 3.999M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingHtwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12.5V12a3 3 0 0 1 3-3h.172a2.828 2.828 0 0 1 2 4.829L15 19h6M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11a7 7 0 1 0-14 0m11 3.5v2c0 .465 0 .697.038.89a2 2 0 0 0 1.572 1.572c.193.038.425.038.89.038s.697 0 .89-.038a2 2 0 0 0 1.571-1.572c.039-.193.039-.425.039-.89v-2c0-.465 0-.697-.039-.89a2 2 0 0 0-1.57-1.572C19.196 12 18.964 12 18.5 12s-.697 0-.89.038a2 2 0 0 0-1.572 1.571c-.038.194-.038.426-.038.891m-8 0v2c0 .465 0 .697-.039.89a2 2 0 0 1-1.57 1.572C6.196 19 5.964 19 5.5 19s-.697 0-.89-.038a2 2 0 0 1-1.572-1.572C3 17.197 3 16.965 3 16.5v-2c0-.465 0-.697.038-.89a2 2 0 0 1 1.572-1.572C4.803 12 5.035 12 5.5 12s.697 0 .89.038a2 2 0 0 1 1.571 1.571c.039.194.039.426.039.891"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartFill(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8.4A5.4 5.4 0 0 1 7.5 3A5.991 5.991 0 0 1 12 5a5.991 5.991 0 0 1 4.5-2A5.4 5.4 0 0 1 22 8.4c0 5.356-6.379 9.4-10 12.6C8.387 17.773 2 13.76 2 8.4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7.694C10 3 3 3.5 3 9.5s9 11 9 11s9-5 9-11s-7-6.5-9-1.806"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21c-.645-.572-1.374-1.167-2.145-1.8h-.01c-2.715-2.22-5.792-4.732-7.151-7.742c-.446-.958-.683-2-.694-3.058A5.39 5.39 0 0 1 7.5 3a6.158 6.158 0 0 1 3.328.983A5.6 5.6 0 0 1 12 5c.344-.39.738-.732 1.173-1.017A6.152 6.152 0 0 1 16.5 3A5.39 5.39 0 0 1 22 8.4a7.422 7.422 0 0 1-.694 3.063c-1.359 3.01-4.435 5.521-7.15 7.737l-.01.008c-.772.629-1.5 1.224-2.145 1.8L12 21ZM7.5 5a3.535 3.535 0 0 0-2.5.992A3.342 3.342 0 0 0 4 8.4c.011.77.186 1.53.512 2.228A12.316 12.316 0 0 0 7.069 14.1c.991 1 2.131 1.968 3.117 2.782c.273.225.551.452.829.679l.175.143c.267.218.543.444.81.666l.013-.012l.006-.005h.006l.009-.007h.01l.018-.015l.041-.033l.007-.006l.011-.008h.006l.009-.008l.664-.545l.174-.143c.281-.229.559-.456.832-.681c.986-.814 2.127-1.781 3.118-2.786a12.298 12.298 0 0 0 2.557-3.471c.332-.704.51-1.472.52-2.25A3.343 3.343 0 0 0 19 6a3.535 3.535 0 0 0-2.5-1a3.988 3.988 0 0 0-2.99 1.311L12 8.051l-1.51-1.74A3.988 3.988 0 0 0 7.5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.237 6.237a4.098 4.098 0 0 1 .135 5.654l-7.373 8.11l-7.37-8.11a4.098 4.098 0 1 1 6.23-5.316L12 8l1.14-1.425a4.098 4.098 0 0 1 6.097-.338"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Help(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.195 8.765A4 4 0 1 1 12 14v1m.05 4v.1h-.1V19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-1-5v2h2v-2h-2Zm1-10a2 2 0 0 1 2 2a1.95 1.95 0 0 1-.59 1.41l-1.24 1.26A4.015 4.015 0 0 0 11 14.5v.5h2a3.42 3.42 0 0 1 1.17-2.83l.9-.92A3.16 3.16 0 0 0 16 9a4 4 0 0 0-8 0h2a2 2 0 0 1 2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpCircleOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c-5.52-.006-9.994-4.48-10-10v-.2C2.11 6.305 6.635 1.928 12.13 2c5.497.074 9.904 4.569 9.868 10.065C21.962 17.562 17.497 22 12 22Zm-.016-2H12a8 8 0 1 0-.016 0ZM13 18h-2v-2h2v2Zm0-3h-2a3.583 3.583 0 0 1 1.77-3.178C13.43 11.316 14 10.88 14 10a2 2 0 1 0-4 0H8v-.09a4 4 0 1 1 8 .09a3.413 3.413 0 0 1-1.56 2.645A3.1 3.1 0 0 0 13 15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpQuestionmark(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 22h-3v-3h3v3Zm0-5h-3v-.007c0-1.65 0-3.075.672-4.073a6.304 6.304 0 0 1 1.913-1.62c.334-.214.649-.417.914-.628a3.712 3.712 0 0 0 1.332-3.824A3.033 3.033 0 0 0 9 8H6a6 6 0 0 1 6-6a6.04 6.04 0 0 1 5.434 3.366a6.017 6.017 0 0 1-.934 6.3c-.453.502-.96.95-1.514 1.337a7.248 7.248 0 0 0-1.316 1.167A4.23 4.23 0 0 0 13 17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hide(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l16 16m-3.5-3.244C15.147 17.485 13.618 18 12 18c-3.53 0-6.634-2.452-8.413-4.221c-.47-.467-.705-.7-.854-1.159c-.107-.327-.107-.913 0-1.24c.15-.459.385-.693.855-1.16c.897-.892 2.13-1.956 3.584-2.793M19.5 14.634c.333-.293.638-.582.912-.854l.003-.003c.468-.466.703-.7.852-1.156c.107-.327.107-.914 0-1.241c-.15-.458-.384-.692-.854-1.159C18.633 8.452 15.531 6 12 6c-.338 0-.671.022-1 .064m2.323 7.436a2 2 0 0 1-2.762-2.889"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1ZM12 5.828l-6 6V20h12v-8.172l-6-6Zm-1.5 12.731l-2.706-2.7L9.2 14.441l1.3 1.292l4.3-4.292l1.412 1.416l-5.712 5.701v.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltFill(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2 13l9.293-9.293a1 1 0 0 1 1.414 0L22 13h-2v8a1 1 0 0 1-1 1h-5v-7h-4v7H5a1 1 0 0 1-1-1v-8H2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1ZM12 5.828l-6 6V20h12v-8.172l-6-6ZM16 16H8v-2h8v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-8H2l9.292-9.293a1 1 0 0 1 1.415 0L22 13h-2v8a1 1 0 0 1-1 1Zm-9-7h4v5h4v-8.172l-6-6l-6 6V20h4v-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltPlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1ZM12 5.828l-6 6V20h12v-8.172l-6-6Zm1 12.171h-2v-3H8v-2h3v-3h2v3h3v2h-3v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltX(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1ZM12 5.828l-6 6V20h12v-8.172l-6-6ZM9.878 18.071l-1.413-1.414l2.121-2.121l-2.121-2.122L9.879 11L12 13.121L14.121 11l1.414 1.414l-2.121 2.121l2.121 2.121l-1.413 1.413L12 15.949l-2.121 2.122h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 20H5a1 1 0 0 1-1-1v-7.86a1 1 0 0 1 .281-.695l5.5-5.7a1 1 0 0 1 1.439 0l2.8 2.9l-1.43 1.402L10.5 6.88L6 11.54v6.455h11v1A1 1 0 0 1 16 20Zm-1.712-4l-2.706-2.7L13 11.882l1.292 1.291l4.3-4.292L20 10.298l-5.712 5.7V16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeFill(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.293 10.707l7-7a1 1 0 0 1 1.414 0l7 7a1 1 0 0 1 .293.707V21a1 1 0 0 1-1 1h-5v-7h-4v7H5a1 1 0 0 1-1-1v-9.586a1 1 0 0 1 .293-.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeHeart(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1ZM12 5.828l-6 6V20h12v-8.172l-6-6Zm0 12.171a22.972 22.972 0 0 0-.692-.582l-.047-.038c-1.157-.944-2.6-2.119-2.6-3.58A1.8 1.8 0 0 1 10.5 12a2.008 2.008 0 0 1 1.5.667A2.009 2.009 0 0 1 13.5 12a1.8 1.8 0 0 1 1.835 1.8c0 1.466-1.452 2.649-2.618 3.6l-.057.047c-.237.194-.461.377-.661.554l.001-.002Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeHeartOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1Zm-7.622-4.525c.22.18.433.356.622.525c.2-.177.424-.36.662-.554l.057-.047c1.166-.951 2.618-2.134 2.618-3.6A1.8 1.8 0 0 0 13.5 12a2.008 2.008 0 0 0-1.5.667A2.007 2.007 0 0 0 10.5 12a1.8 1.8 0 0 0-1.835 1.8c0 1.461 1.44 2.636 2.6 3.58h-.001h.005l.018.015h.005l.008.019l.052.043l.013.01h.006l.007.008Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 20H5a1 1 0 0 1-1-1v-7.86a1 1 0 0 1 .281-.695l5.5-5.7a1 1 0 0 1 1.439 0l2.8 2.9l-1.43 1.402L10.5 6.88L6 11.54v6.455h11v1A1 1 0 0 1 16 20Zm4-6h-8v-2h8v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.586a1 1 0 0 1 .293-.707l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1Zm-9-7h4v5h4v-8.172l-6-6l-6 6V20h4v-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomePlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 20H5a1 1 0 0 1-1-1v-7.86a1 1 0 0 1 .281-.695l5.5-5.7a1 1 0 0 1 1.439 0l1.651 1.713l-1.433 1.394l-.938-.972L6 11.54v6.455h11v1A1 1 0 0 1 16 20Zm1-4h-2v-3h-3v-2h3V8h2v3h3v2h-3v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeX(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 20H5a1 1 0 0 1-1-1v-7.86a1 1 0 0 1 .281-.695l5.5-5.7a1 1 0 0 1 1.439 0l2.136 2.215l-1.434 1.394L10.5 6.88L6 11.54v6.455h11v1A1 1 0 0 1 16 20Zm-1.658-4l-1.413-1.414l2.121-2.122l-2.121-2.124l1.414-1.414l2.121 2.121l2.121-2.121L20 10.34l-2.121 2.124L20 14.586l-1.414 1.413l-2.122-2.121L14.343 16h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v-3m0 0v-3m0 3H9m3 0h3M4 16.8v-5.348c0-.534 0-.801.065-1.05c.058-.22.152-.429.28-.617c.145-.213.346-.39.748-.741l4.801-4.202c.746-.652 1.119-.978 1.538-1.102c.37-.11.765-.11 1.135 0c.42.124.794.45 1.54 1.104l4.8 4.2c.403.352.603.528.748.74a2 2 0 0 1 .28.618c.065.248.065.516.065 1.05v5.352c0 1.118 0 1.677-.218 2.105a2 2 0 0 1-.874.873c-.428.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a1.999 1.999 0 0 1-.874-.873C4 18.48 4 17.92 4 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l-4 4l-2-2m-5 3.8v-5.348c0-.534 0-.801.065-1.05a2 2 0 0 1 .28-.617c.145-.213.346-.39.748-.741l4.801-4.202c.746-.652 1.119-.978 1.538-1.102c.37-.11.765-.11 1.135 0c.42.124.794.45 1.54 1.104l4.8 4.2c.403.352.603.528.748.74c.127.19.222.398.28.618c.065.249.065.516.065 1.05v5.352c0 1.118 0 1.677-.218 2.105a2 2 0 0 1-.875.873c-.427.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a1.999 1.999 0 0 1-.874-.873C4 18.48 4 17.92 4 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 15l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2M4 16.8v-5.348c0-.534 0-.801.065-1.05c.058-.22.152-.429.28-.617c.145-.213.346-.39.748-.741l4.801-4.202c.746-.652 1.119-.978 1.538-1.102c.37-.11.765-.11 1.135 0c.42.124.794.45 1.54 1.104l4.8 4.2c.403.352.603.528.748.74a2 2 0 0 1 .28.618c.065.248.065.516.065 1.05v5.352c0 1.118 0 1.677-.218 2.105a2 2 0 0 1-.874.873c-.428.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a1.999 1.999 0 0 1-.874-.873C4 18.48 4 17.92 4 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 17v-5.548c0-.534 0-.801-.065-1.05a1.998 1.998 0 0 0-.28-.617c-.145-.213-.345-.39-.748-.741l-4.8-4.2c-.746-.653-1.12-.98-1.54-1.104c-.37-.11-.764-.11-1.135 0c-.42.124-.792.45-1.538 1.102L5.093 9.044c-.402.352-.603.528-.747.74a2 2 0 0 0-.281.618C4 10.65 4 10.918 4 11.452V17c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.082 1.083C5.602 20 6.068 20 7 20s1.398 0 1.766-.152a2 2 0 0 0 1.082-1.083C10 18.398 10 17.932 10 17v-1a2 2 0 1 1 4 0v1c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.082 1.083C15.602 20 16.068 20 17 20s1.398 0 1.766-.152a2 2 0 0 0 1.082-1.083C20 18.398 20 17.932 20 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 13h6M4 16.8v-5.348c0-.534 0-.801.065-1.05a2 2 0 0 1 .28-.617c.145-.213.346-.39.748-.741l4.801-4.202c.746-.652 1.119-.978 1.538-1.102c.37-.11.765-.11 1.135 0c.42.124.794.45 1.54 1.104l4.8 4.2c.403.352.603.528.748.74c.127.19.222.398.28.618c.065.249.065.516.065 1.05v5.352c0 1.118 0 1.677-.218 2.105a2 2 0 0 1-.875.873c-.427.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a1.999 1.999 0 0 1-.874-.873C4 18.48 4 17.92 4 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseThree(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h2m0 0h6m-6 0v-8.548c0-.534 0-.801.065-1.05a2 2 0 0 1 .28-.617c.145-.213.346-.39.748-.741l4.801-4.202c.746-.652 1.119-.978 1.538-1.102c.37-.11.765-.11 1.135 0c.42.124.794.45 1.54 1.104l4.8 4.2c.402.352.603.528.748.74c.127.19.222.398.28.618c.065.249.065.516.065 1.05V20m-10 0h4m-4 0v-4a2 2 0 1 1 4 0v4m0 0h6m0 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 11.452V16.8c0 1.12 0 1.68.218 2.109c.192.376.497.682.874.873c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.873c.218-.428.218-.987.218-2.105v-5.352c0-.534 0-.801-.065-1.05a1.998 1.998 0 0 0-.28-.617c-.145-.213-.345-.39-.748-.741l-4.8-4.2c-.746-.653-1.12-.98-1.54-1.104c-.37-.11-.764-.11-1.135 0c-.42.124-.792.45-1.538 1.102L5.093 9.044c-.402.352-.603.528-.747.74a2 2 0 0 0-.281.618C4 10.65 4 10.918 4 11.452"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.983 20.988l-6.414-1.826l-1.433-16.15h15.729l-1.429 16.15l-6.453 1.826Zm-4.292-7.691l.245 3.123l4.063 1.085l4.028-1.08l.559-6.113H9.402l-.173-2.033h7.533l.174-1.961h-9.87l.522 6h6.836l-.243 2.566l-2.179.6l-2.216-.607l-.141-1.58H7.691Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCard(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 20H4a1.943 1.943 0 0 1-2-1.876V5.875A1.942 1.942 0 0 1 4 4h16a1.942 1.942 0 0 1 2 1.875v12.25A1.943 1.943 0 0 1 20 20ZM4 6v11.989L20 18V6.011L4 6Zm9.43 10H6a3.21 3.21 0 0 1 1.093-2.14a3.829 3.829 0 0 1 2.622-1.11c.984.02 1.923.417 2.622 1.11A3.212 3.212 0 0 1 13.43 16ZM18 15h-3v-2h3v2Zm-8.285-3a1.934 1.934 0 0 1-2-2a1.935 1.935 0 0 1 2-2a1.935 1.935 0 0 1 2 2a1.934 1.934 0 0 1-2 2ZM18 11h-4V9h4v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM5 5v14h14V5H5Zm13 12H6l3-4l1 1l3-4l5 7Zm-9.5-6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 18h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2Zm0-2V4h12v12H8Z"/><path fill="currentColor" d="M2 8h2v12h12v2H4a2 2 0 0 1-2-2V8Z"/><path fill="currentColor" d="M10 14h9l-3.75-5.444l-2.25 3.11l-.75-.777L10 14Zm.75-5.833c0 .644.504 1.166 1.125 1.166S13 8.811 13 8.167C13 7.522 12.496 7 11.875 7s-1.125.522-1.125 1.167Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105v9.606c0 .485 0 .865-.018 1.174M3 17c0 .988.013 1.506.218 1.907c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.607c1.118 0 1.677 0 2.104-.218c.376-.192.682-.498.874-.875c.123-.242.177-.526.2-.93M3 17l4.768-5.562l.001-.002c.423-.493.635-.74.886-.83a1 1 0 0 1 .681.005c.25.093.46.343.876.843l2.671 3.205c.386.464.58.696.816.79a1 1 0 0 0 .651.028c.244-.072.46-.287.889-.716l.497-.497c.437-.438.656-.656.904-.728a.995.995 0 0 1 .659.037c.238.099.431.34.818.822l2.865 3.582m0 0L21 18m-6-8a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 18V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 3 5.08 3 6.2 3h11.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105v11.606c0 .485 0 .865-.018 1.174M3 18c0 .988.013 1.506.218 1.907c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.607c1.118 0 1.677 0 2.104-.218c.376-.192.682-.498.874-.875c.123-.242.177-.526.2-.93M3 18l4.768-5.562l.001-.002c.423-.493.635-.74.886-.83a1 1 0 0 1 .681.005c.25.093.46.343.876.843l2.671 3.205c.386.464.58.696.816.79a1 1 0 0 0 .651.028c.244-.072.46-.287.889-.716l.497-.497c.437-.438.656-.656.904-.728a.995.995 0 0 1 .659.037c.238.099.431.34.818.822l2.865 3.582m0 0L21 19M15 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11v5m0 5a9 9 0 1 1 0-18a9 9 0 0 1 0 18m.05-13v.1h-.1V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10ZM9.99 10.99V13h1v4h3.02v-2H13l.01-4.009l-3.02-.001Zm1-3.99v2.019h2.02V7h-2.02Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-8-9.828A8 8 0 1 0 4 12v.172ZM14 17h-3v-4h-1v-2h3v4h1v2Zm-1-8h-2V7h2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoSquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1ZM10 11v2h1v4h3v-2h-1v-4h-3Zm1-4v2h2V7h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoSquareOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1ZM5 5v14h14V5H5Zm9 12h-3v-4h-1v-2h3v4h1v2Zm-1-8h-2V7h2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.002 21.041c-2.46 0-2.75-.013-3.71-.055a6.583 6.583 0 0 1-2.185-.45a4.619 4.619 0 0 1-2.63-2.631a6.591 6.591 0 0 1-.419-2.187c-.056-.958-.056-1.272-.056-3.713c0-2.467.013-2.755.056-3.71a6.59 6.59 0 0 1 .419-2.184A4.61 4.61 0 0 1 6.11 3.479a6.533 6.533 0 0 1 2.184-.42c.955-.054 1.269-.054 3.708-.054c2.48 0 2.765.013 3.71.054c.748.014 1.49.156 2.19.42a4.615 4.615 0 0 1 2.633 2.632c.267.71.41 1.46.421 2.217c.056.958.056 1.271.056 3.711s-.014 2.76-.056 3.707a6.603 6.603 0 0 1-.42 2.189a4.624 4.624 0 0 1-2.634 2.632c-.7.262-1.439.404-2.186.419c-.955.055-1.268.055-3.714.055Zm-.034-16.453c-2.446 0-2.7.012-3.655.055a4.99 4.99 0 0 0-1.67.311a2.99 2.99 0 0 0-1.718 1.712c-.2.54-.305 1.111-.311 1.687c-.053.969-.053 1.223-.053 3.652c0 2.4.009 2.691.053 3.654c.009.57.114 1.135.311 1.67c.306.787.93 1.409 1.719 1.711a4.92 4.92 0 0 0 1.669.311c.968.056 1.223.056 3.655.056c2.453 0 2.707-.012 3.654-.056a4.962 4.962 0 0 0 1.67-.311a3 3 0 0 0 1.71-1.709c.2-.54.305-1.112.311-1.688h.011c.043-.956.043-1.211.043-3.654c0-2.443-.011-2.7-.054-3.655a5.06 5.06 0 0 0-.311-1.668a3 3 0 0 0-1.71-1.712a4.94 4.94 0 0 0-1.67-.311c-.967-.055-1.22-.055-3.654-.055Zm.034 12.036A4.623 4.623 0 1 1 16.622 12a4.63 4.63 0 0 1-4.62 4.624Zm0-7.626a3 3 0 1 0 3 3a3.007 3.007 0 0 0-3-3Zm4.8-.713a1.078 1.078 0 1 1 .008-2.156a1.078 1.078 0 0 1-.008 2.156Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instance(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.263 20.344l6.081-6.081c.792-.792 1.188-1.188 1.336-1.645c.13-.402.13-.834 0-1.235c-.148-.457-.544-.854-1.336-1.646l-6.082-6.082c-.792-.791-1.187-1.187-1.644-1.335a2 2 0 0 0-1.236 0c-.456.148-.853.545-1.644 1.336L3.656 9.737v.001c-.792.792-1.188 1.188-1.336 1.645c-.13.401-.13.833 0 1.235c.148.456.544.852 1.335 1.644l6.086 6.086c.79.79 1.185 1.185 1.641 1.333a2 2 0 0 0 1.236 0c.457-.149.853-.544 1.644-1.336z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invision(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.362 20.993H4.638a1.63 1.63 0 0 1-1.63-1.63V4.637c.001-.9.73-1.629 1.63-1.63h14.723c.9.001 1.63.73 1.632 1.63v14.724c-.001.9-.73 1.63-1.63 1.631ZM6.673 9.67l-.388 1.43h1.278l-.826 3.366c-.06.266-.094.538-.1.811a1.444 1.444 0 0 0 1.621 1.584c.93-.04 1.765-.58 2.183-1.412l-.327 1.3h1.811l1.033-4.14c.26-1.06.772-1.6 1.52-1.6a.907.907 0 0 1 .97.983c.003.197-.026.393-.085.581l-.528 1.906a2.904 2.904 0 0 0-.114.812a1.48 1.48 0 0 0 1.651 1.566c1.03 0 1.748-.7 2.137-2.08l-.707-.274c-.276.79-.57 1.173-.9 1.173c-.24 0-.368-.167-.368-.484c.007-.166.032-.331.075-.492l.52-1.858c.118-.382.179-.78.18-1.18a1.9 1.9 0 0 0-1.855-2.11a3.13 3.13 0 0 0-2.406 1.761l.36-1.617H10.62l-.39 1.4h1.293l-.8 3.188a2.09 2.09 0 0 1-1.826 1.383a.515.515 0 0 1-.092-.007c-.243-.057-.384-.15-.384-.452c.006-.25.043-.498.11-.738l1.216-4.8H6.673Zm2.345-3.162a1.09 1.09 0 1 0 0 2.18v-.025c.292.002.572-.11.781-.314a1.06 1.06 0 0 0 .317-.766a1.089 1.089 0 0 0-1.098-1.075Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19h2m0 0h2m-2 0l4-14m-2 0h2m0 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Javascript(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 21H3V3h18v18ZM9.1 17.232l-1.368.844c.213.468.559.863.994 1.137c.473.271 1.01.41 1.555.4c.286-.002.57-.035.849-.1a2.108 2.108 0 0 0 1.358-1.058c.28-.66.384-1.382.3-2.094v-.415c.005-1.027 0-2.072 0-3.084v-1.593H11.1v4.897c.05.5.014 1.003-.107 1.49a.786.786 0 0 1-.759.422a1.43 1.43 0 0 1-.416-.063a1.43 1.43 0 0 1-.623-.641l-.033-.056c-.031-.053-.052-.09-.062-.09v.004Zm5.949-.2l-1.373.787c.112.246.268.47.458.662c.047.052.1.108.149.17c.671.6 1.549.918 2.449.887a2.63 2.63 0 0 0 2.7-1.641v-.008c.114-.4.132-.82.053-1.228l.034.049c-.148-.929-.821-1.574-2.252-2.155c-.111-.052-.226-.1-.338-.15c-.445-.192-.865-.372-1.011-.7a.886.886 0 0 1-.034-.529a.762.762 0 0 1 .808-.543c.111 0 .222.016.329.048c.33.11.595.357.732.676c.775-.507.775-.507 1.316-.844a3.247 3.247 0 0 0-.439-.586a2.466 2.466 0 0 0-2-.776h-.12l-.528.067a2.435 2.435 0 0 0-1.283.754a2.354 2.354 0 0 0 .427 3.354c.438.3.91.546 1.407.733c.641.27 1.194.5 1.306.921a.811.811 0 0 1-.135.707a1.3 1.3 0 0 1-1.027.37a2.58 2.58 0 0 1-.313-.02a1.977 1.977 0 0 1-1.317-1.008l.002.003Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 15h1M9 15h6m-9 0H5m0-3h14M5 9h14M2 14.8V9.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C3.52 6 4.08 6 5.2 6h13.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v5.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H5.197c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C2 16.48 2 15.92 2 14.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19.293 9.951l-2.333-2.8c-.353-.423-.53-.635-.746-.787a2.001 2.001 0 0 0-.632-.295C15.327 6 15.052 6 14.502 6H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 7.52 4 8.08 4 9.2v5.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218H14.5c.551 0 .826 0 1.081-.069c.226-.06.44-.16.632-.296c.216-.152.393-.363.746-.786l2.333-2.8c.608-.729.91-1.093 1.027-1.5c.102-.359.102-.74 0-1.098c-.116-.407-.42-.77-1.027-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 17h-.5a1.5 1.5 0 0 0 0 3h17a1.5 1.5 0 0 0 0-3H20M4 17h16M4 17V8.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 5 6.08 5 7.2 5h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105V17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LastPage(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.558 18.01l-1.414-1.413l4.6-4.6l-4.6-4.593L7.558 5.99L13.569 12l-6.01 6.01h-.001Zm10.3-.01h-2V6h2v12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layer(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 14l-9 6l-9-6m18-4l-9 6l-9-6l9-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layers(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 12l-9 6l-9-6m18 4l-9 6l-9-6m18-8l-9 6l-9-6l9-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 23.27l-9-7l1.62-1.26l7.37 5.73l7.38-5.739L21 16.27l-9 7ZM12 19l-9-7l1.62-1.26l7.37 5.73l7.38-5.74L21 12l-9 7Zm0-4.27L4.63 9L3 7.73l9-7l9 7L19.36 9L12 14.73Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.83 17.08c7.07 4.243 12.727-1.414 12.02-12.02C8.244 4.353 2.587 10.01 6.83 17.08m0 0c-.001 0 0 0 0 0m0 0L5 18.91m1.83-1.828l3.827-3.829"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineBreak(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6v6H7.83l2.58-2.59L9 8l-5 5l5 5l1.41-1.41L7.83 14H20V6h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChartDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 21H4a1 1 0 0 1-1-1V3h2v16h16v2Zm-1.373-5l-4.17-4.082l-2.228 2.182a.985.985 0 0 1-1.373 0L7 9.344L8.373 8l4.17 4.082L14.77 9.9a.985.985 0 0 1 1.373 0L21 14.656L19.627 16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChartUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 21H4a1 1 0 0 1-1-1V3h2v16h16v2ZM8.373 16L7 14.656L11.856 9.9a.985.985 0 0 1 1.373 0l2.227 2.181L19.627 8L21 9.344L16.144 14.1a.985.985 0 0 1-1.373 0l-2.228-2.182L8.374 16h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineL(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19V5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineM(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineS(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15V9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineSx(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 10h-2v4h2v-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineXl(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.172 14.829l5.657-5.657M7.05 11.293l-1.414 1.414a4 4 0 1 0 5.657 5.657l1.412-1.414m-1.413-9.9l1.414-1.414a4 4 0 1 1 5.657 5.657l-1.414 1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkBreak(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 20v-2m2-2h2M7.05 11.293l-1.414 1.414a4 4 0 0 0 5.657 5.657l1.415-1.414M6 8H4m4-4v2m3.293 1.05l1.414-1.414a4 4 0 1 1 5.656 5.657l-1.414 1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkHorizontal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h8m-1-4h2a4 4 0 0 1 0 8h-2M9 8H7a4 4 0 1 0 0 8h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkHorizontalOff(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h4m0 0L8 8m4 4l8 8M15 8h2a4 4 0 0 1 2.645 7M9 16H7a4 4 0 0 1 0-8h1m0 0L4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.465 20.535A5 5 0 0 1 4.929 12L7.05 9.878l1.414 1.414l-2.121 2.121a3 3 0 1 0 4.243 4.243l2.12-2.121l1.415 1.415L12 19.071a4.969 4.969 0 0 1-3.536 1.464Zm.707-4.293l-1.414-1.414l7.07-7.071l1.415 1.414l-7.07 7.07l-.001.001Zm7.779-2.121l-1.415-1.414l2.12-2.121a3 3 0 1 0-4.241-4.243l-2.122 2.121L9.879 7.05L12 4.928a5 5 0 0 1 7.07 7.071l-2.12 2.121v.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkVertical(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v8m4-1v2a4 4 0 0 1-8 0v-2m8-6V7a4 4 0 0 0-8 0v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 21H9V9h4v2a4.618 4.618 0 0 1 3.525-1.763A4.5 4.5 0 0 1 21 13.75V21h-4v-6.75a2.265 2.265 0 0 0-2.247-1.944A1.815 1.815 0 0 0 13 14.25V21Zm-6 0H3V9h4v12ZM5 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkpath(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.396 22a1.84 1.84 0 0 1-.918-.246l-4.52-2.61a1.84 1.84 0 0 1-.917-1.588v-5.22a1.84 1.84 0 0 1 .917-1.589l1.373-.793v1.71a2.76 2.76 0 0 0 1.376 2.383l4.521 2.61a2.75 2.75 0 0 0 2.752 0l1.772-1.022v1.921a1.84 1.84 0 0 1-.917 1.589l-4.521 2.61a1.84 1.84 0 0 1-.918.245Zm3.207-5.891c-.322 0-.638-.085-.917-.246l-4.521-2.61a1.84 1.84 0 0 1-.918-1.589v-5.22c.001-.655.35-1.26.918-1.59l4.521-2.608a1.834 1.834 0 0 1 1.834 0l4.521 2.61c.567.328.916.934.917 1.589v5.22c0 .655-.35 1.26-.917 1.589l-4.521 2.61c-.279.16-.595.245-.917.245Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17h7m5-1h3m0 0h3m-3 0v3m0-3v-3M3 12h11M3 7h11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 17h7m9-3l-4 4l-2-2M4 12h11M4 7h11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListChecklist(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 17h9M8 15l-2.5 3L4 17m7-5h9M8 10l-2.5 3L4 12m7-5h9M8 5L5.5 8L4 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListChecklistAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4h10v2H11V4Zm0 4h6v2h-6V8Zm0 6h10v2H11v-2Zm0 4h6v2h-6v-2ZM3 4h6v6H3V4Zm2 2v2h2V6H5Zm-2 8h6v6H3v-6Zm2 2v2h2v-2H5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 17h-9v-2h9v2Zm-11 0H2v-2h9v2Zm4-4H2v-2h13v2Zm0-4H2V7h13v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOl(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.983 19H4v-1h1.989v-.5h-.994v-1h.995V16H4v-1h2.983v4ZM21 18H9.069v-2H21v2ZM6.983 14H4v-.9L5.79 11H4v-1h2.983v.9L5.193 13h1.79v1ZM21 13H9.069v-2H21v2ZM6.486 9h-.995V6H4.5V5h1.986v4ZM21 8H9.069V6H21v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOrdered(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 17h10M4 15.685V15.5A1.5 1.5 0 0 1 5.5 14h.04c.807 0 1.46.653 1.46 1.46c0 .35-.114.692-.324.973L4 20h3m3-8h10M10 7h10M4 5l2-1v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListPlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 20h-2v-3h-3v-2h3v-3h2v3h3v2h-3v3Zm-7-3H2v-2h10v2Zm3-4H2v-2h13v2Zm0-4H2V7h13v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17h7m5-1h6M3 12h11M3 7h11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUl(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 18H8v-2h12v2ZM6 18H4v-2h2v2Zm14-5H8v-2h12v2ZM6 13H4v-2h2v2Zm14-5H8.023V6H20v2ZM6 8H4V6h2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUnordered(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17h10M9 12h10M9 7h10M5.002 17v.002H5V17zm0-5v.002H5V12zm0-5v.002H5V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loading(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 14a1 1 0 1 0 2 0a1 1 0 0 0-2 0m6-2a1 1 0 1 0 2 0a1 1 0 0 0-2 0m6-2a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a29.776 29.776 0 0 1-3.5-3.531C6.9 15.558 5 12.712 5 10a7 7 0 0 1 11.952-4.951A6.955 6.955 0 0 1 19 10c0 2.712-1.9 5.558-3.5 7.469A29.777 29.777 0 0 1 12 21Zm0-14a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a29.776 29.776 0 0 1-3.5-3.531C6.9 15.558 5 12.712 5 10a7 7 0 0 1 11.952-4.951A6.955 6.955 0 0 1 19 10c0 2.712-1.9 5.558-3.5 7.469A29.777 29.777 0 0 1 12 21Zm0-16a5.006 5.006 0 0 0-5 5c0 1.166.527 3.185 3.035 6.186A27.93 27.93 0 0 0 12 18.3a28.121 28.121 0 0 0 1.966-2.111C16.473 13.184 17 11.165 17 10a5.006 5.006 0 0 0-5-5Zm0 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.23 9H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 10.52 4 11.08 4 12.2v5.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104v-5.607c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 9 17.92 9 16.8 9h-2.03M9.23 9h5.539M9.23 9A.23.23 0 0 1 9 8.77V6a3 3 0 1 1 6 0v2.77a.23.23 0 0 1-.231.23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 9H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 10.52 4 11.08 4 12.2v5.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104v-5.607c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 9 17.92 9 16.8 9zm0 0V6.12C9 4.397 10.3 3 11.904 3c.824 0 1.567.369 2.095.961"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOut(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 15l3-3m0 0l-3-3m3 3H4m5-4.751V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C10.52 4 11.08 4 12.2 4h4.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v9.607c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218h-4.606c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C9 18.48 9 17.92 9 16.8v-.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongBottomDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.021 15.55L7.014 12H5v7h7v-2.014l-3.55-.007L19 6.429L17.571 5L7.021 15.55Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongBottomUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.979 8.45l.007 3.55H19V5h-7v2.014l3.55.007L5 17.571L6.429 19l10.55-10.55Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11 18.17l-2.59-2.58L7 17l5 5l5-5l-1.41-1.41L13 18.17V2h-2v16.17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.83 11l2.58-2.59L7 7l-5 5l5 5l1.41-1.41L5.83 13H22v-2H5.83Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.17 13l-2.58 2.59L17 17l5-5l-5-5l-1.41 1.41L18.17 11H2v2h16.17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 5.83l2.59 2.58L17 7l-5-5l-5 5l1.41 1.41L11 5.83V22h2V5.83Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongUpLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.45 7.021L12 7.014V5H5v7h2.014l.007-3.55L17.571 19L19 17.571L8.45 7.021Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 10h6m2 5l6 6m-11-4a7 7 0 1 1 0-14a7 7 0 0 1 0 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlassPlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 10h3m0 0h3m-3 0V7m0 3v3m5 2l6 6m-11-4a7 7 0 1 1 0-14a7 7 0 0 1 0 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 6l6.108 4.612l.002.002c.678.497 1.017.746 1.389.842a2 2 0 0 0 1.002 0c.372-.096.712-.345 1.392-.844c0 0 3.917-3.006 6.107-4.612M3 15.8V8.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v7.607c0 1.118 0 1.676-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218H6.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C3 17.48 3 16.92 3 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailOpen(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 10l6.108 4.612l.002.002c.678.497 1.017.746 1.389.842a2 2 0 0 0 1.002 0c.372-.096.712-.346 1.392-.844L20 10m-.2-.96l-5.599-4.483c-.695-.557-1.043-.835-1.43-.946a2.001 2.001 0 0 0-1.046-.014c-.389.1-.744.368-1.454.905L4.27 9.04c-.466.352-.699.528-.867.75a2 2 0 0 0-.327.658C3 10.716 3 11.008 3 11.592V17.8c0 1.12 0 1.68.218 2.108c.192.377.497.682.874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.427.218-.987.218-2.105v-6.276c0-.558 0-.838-.071-1.097a2.003 2.003 0 0 0-.31-.642c-.159-.22-.378-.396-.819-.749"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MainComponent(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.235 2.374c-.368.152-.697.482-1.356 1.14c-.659.66-.989.989-1.14 1.356a2 2 0 0 0 0 1.531c.151.368.48.697 1.14 1.356c.658.659.988.989 1.356 1.14a2 2 0 0 0 1.53 0c.368-.151.697-.48 1.356-1.14c.66-.659.988-.988 1.14-1.356a2 2 0 0 0 0-1.53c-.152-.368-.48-.697-1.14-1.356c-.659-.66-.988-.989-1.356-1.141a2 2 0 0 0-1.53 0M4.87 8.738c-.367.152-.697.481-1.355 1.14c-.66.66-.989.989-1.141 1.356a2 2 0 0 0 0 1.531c.152.368.482.697 1.14 1.356c.66.659.989.988 1.356 1.14a2 2 0 0 0 1.531 0c.368-.152.697-.481 1.356-1.14c.66-.659.988-.988 1.14-1.356a2 2 0 0 0 0-1.53c-.152-.368-.48-.698-1.14-1.357c-.659-.659-.988-.988-1.356-1.14a2 2 0 0 0-1.53 0m11.372 1.14c-.659.66-.988.989-1.14 1.356a2 2 0 0 0 0 1.531c.152.368.481.697 1.14 1.356c.659.659.989.988 1.356 1.14a2 2 0 0 0 1.53 0c.368-.152.698-.481 1.357-1.14c.659-.659.987-.988 1.14-1.356a2 2 0 0 0 0-1.53c-.153-.368-.481-.698-1.14-1.357c-.66-.659-.989-.988-1.356-1.14a2 2 0 0 0-1.531 0c-.367.152-.697.481-1.356 1.14m-5.008 5.224c-.368.152-.697.482-1.356 1.14c-.659.66-.989.989-1.14 1.357a2 2 0 0 0 0 1.53c.151.368.48.697 1.14 1.356c.658.659.988.989 1.356 1.14a2 2 0 0 0 1.53 0c.368-.151.697-.48 1.356-1.14c.66-.659.988-.988 1.14-1.356c.203-.49.203-1.04 0-1.53c-.152-.368-.48-.698-1.14-1.356c-.659-.66-.988-.989-1.356-1.141a2 2 0 0 0-1.53 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 6v15m0-15l6-3v15l-6 3m0-15L9 3m6 18l-6-3m0 0l-6 3V6l6-3m0 15V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPin(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 9.923c0 4.852 4.244 8.864 6.123 10.402c.27.22.405.332.606.388c.156.044.386.044.542 0c.201-.056.336-.167.606-.388C14.756 18.787 19 14.775 19 9.923a6.885 6.885 0 0 0-2.05-4.895A7.04 7.04 0 0 0 12 3a7.04 7.04 0 0 0-4.95 2.028A6.884 6.884 0 0 0 5 9.923"/><path d="M10 9a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mention(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12.002V13a2 2 0 1 0 4 0v-1a7 7 0 1 0-4.406 6.502m.406-6.5a3 3 0 1 1 0-.004m0 .004v-.004m0 0V9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuAltFive(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h8m-8-5h14m-8-5h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuAltFour(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h14M5 12h14M5 7h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuAltOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17h7M5 12h14M5 7h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuAltThree(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h8m-8-5h14M5 7h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuAltTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 17h8M5 12h14m-8-5h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuDuo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 16H3v-2h18v2Zm0-6H3V8h18v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuDuoLg(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15h18M3 9h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuDuoMd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15h14M5 9h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 21V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H9c-.433 0-.854.14-1.2.4L3 21ZM5 5v12l2.134-1.6a1.99 1.99 0 0 1 1.2-.4H19V5H5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 21V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H9c-.433 0-.854.14-1.2.4L3 21ZM5 5v12l2.134-1.6a1.99 1.99 0 0 1 1.2-.4H19V5H5Zm6 8.561L7.293 9.853L8.707 8.44L11 10.733l4.293-4.293l1.414 1.414L11 13.56v.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.051 17.828l.654.35a6.96 6.96 0 0 0 3.292.822H12a7 7 0 1 0-7-7v.003a6.96 6.96 0 0 0 .822 3.292l.35.654l-.538 2.417l2.417-.538ZM3 21l1.058-4.762A9 9 0 0 1 12 3a9 9 0 0 1 9 9a9 9 0 0 1-13.238 7.942L3 21Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.879 13.536L8.464 12.12L10.586 10L8.464 7.879L9.88 6.464L12 8.586l2.121-2.122l1.415 1.415L13.414 10l2.122 2.121l-1.415 1.415L12 11.414l-2.121 2.122Z"/><path fill="currentColor" d="M3 5v16l4.8-3.6c.346-.26.767-.4 1.2-.4h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2Zm2 12V5h14v10H8.334a1.984 1.984 0 0 0-1.2.4L5 17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 21V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H9c-.433 0-.854.14-1.2.4L3 21ZM5 5v12l2.134-1.6a1.99 1.99 0 0 1 1.2-.4H19V5H5Zm11 6H8V9h8v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 21V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H9c-.433 0-.854.14-1.2.4L3 21ZM5 5v12l2.134-1.6a1.99 1.99 0 0 1 1.2-.4H19V5H5Zm8 9h-2v-3H8V9h3V6h2v3h3v2h-3v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlusAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 10h-2V7h-3V5h3V2h2v3h3v2h-3v3Z"/><path fill="currentColor" d="M21 12h-2v3H8.334a1.984 1.984 0 0 0-1.2.4L5 17V5h7V3H5a2 2 0 0 0-2 2v16l4.8-3.6c.346-.26.767-.4 1.2-.4h10a2 2 0 0 0 2-2v-3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRound(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.124 12.114l.003.012a5.909 5.909 0 0 0 3.336 4.16L12 17.894V16h2a5 5 0 0 0 5-5v-1a5 5 0 0 0-5-5h-4a5 5 0 0 0-5 5v1c0 .38.042.748.121 1.1l.003.014ZM14 21l-6.364-2.893A7.909 7.909 0 0 1 3.17 12.54A7.024 7.024 0 0 1 3 11v-1a7 7 0 0 1 7-7h4a7 7 0 0 1 7 7v1a7 7 0 0 1-7 7v3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageWriting(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9H7v2h2V9Zm2 0h2v2h-2V9Zm6 0h-2v2h2V9Z"/><path fill="currentColor" d="M3 5v16l4.8-3.6c.346-.26.767-.4 1.2-.4h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2Zm2 12V5h14v10H8.334a1.984 1.984 0 0 0-1.2.4L5 17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Messenger(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.401 21v-3.189a8.1 8.1 0 0 1-3.31-6.479C3.09 6.738 7.09 3 12 3s8.91 3.738 8.91 8.332c.001 4.594-4 8.33-8.91 8.33a9.463 9.463 0 0 1-2.559-.349L6.403 21h-.002Zm4.66-11.931l-4.866 5.163l4.438-2.454l2.3 2.394L17.8 9.01l-4.434 2.454l-2.305-2.395Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 13v-2h14v2H5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c-5.52-.006-9.994-4.48-10-10v-.2C2.11 6.305 6.635 1.928 12.13 2c5.497.074 9.904 4.569 9.868 10.065C21.962 17.562 17.497 22 12 22ZM7 11v2h10v-2H7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-8-9.828A8 8 0 1 0 4 12v.172ZM17 13H7v-2h10v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM5 5v14h14V5H5Zm12 8H7v-2h10v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6.2v11.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h3.606c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104V6.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C15.48 3 14.92 3 13.8 3h-3.6c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C7 4.52 7 5.08 7 6.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.75 23h-10a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v18a2 2 0 0 1-2 2Zm-10-20v18h10V3h-10Zm5 17a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileButton(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h3.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V6.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C15.48 3 14.92 3 13.8 3h-3.6c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C7 4.52 7 5.08 7 6.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitor(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20H9m-5-6.2V8.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 5 6.08 5 7.2 5h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v5.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C4 15.48 4 14.92 4 13.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MonitorPlay(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 20H9m-5-6.2V8.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 5 6.08 5 7.2 5h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v5.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H7.197c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C4 15.48 4 14.92 4 13.8M14.5 11L10 8v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6a9 9 0 0 0 9 9c.91 0 1.787-.134 2.614-.385A9.004 9.004 0 0 1 12 21A9 9 0 0 1 9.386 3.386A8.999 8.999 0 0 0 9 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreGridBig(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 18a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0m12-6a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0m12-6a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0M5 6a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreGridSmall(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 15a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0m6-6a1 1 0 1 0 2 0a1 1 0 0 0-2 0M8 9a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreHorizontal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVertical(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 18a1 1 0 1 0 2 0a1 1 0 0 0-2 0m0-6a1 1 0 1 0 2 0a1 1 0 0 0-2 0m0-6a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10V7m6 2v6a6 6 0 0 1-12 0V9a6 6 0 1 1 12 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21v-9m0 9l3-3m-3 3l-3-3m3-6V3m0 9H3m9 0h9m-9-9L9 6m3-3l3 3M3 12l3 3m-3-3l3-3m15 3l-3-3m3 3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveHorizontal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h18M3 12l3 3m-3-3l3-3m15 3l-3-3m3 3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveVertical(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21V3m0 18l3-3m-3 3l-3-3m3-15L9 6m3-3l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MovingDesk(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12v3m0-3h1.803c1.118 0 1.677 0 2.104-.218c.377-.191.683-.498.875-.875C21 10.48 21 9.921 21 8.803v-.606c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C19.48 5 18.92 5 17.8 5h-3.6c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C11 6.52 11 7.08 11 8.2v.6c0 1.12 0 1.68.218 2.107c.192.377.498.684.874.875c.427.218.987.218 2.105.218zm-9-2l-2 2l-2-2m0-3l2-2l2 2M3 21v-1.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 16 5.08 16 6.2 16h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105V21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Navigation(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4.805 3.555l15.81 4.865c.862.265.96 1.446.153 1.85l-6.7 3.35a1 1 0 0 0-.448.447l-3.35 6.7c-.403.807-1.584.71-1.85-.153L3.555 4.804a1 1 0 0 1 1.25-1.249"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Note(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 20H7.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C4 18.48 4 17.92 4 16.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105V13m-7 7c.286-.003.466-.014.639-.055c.204-.05.399-.13.578-.24c.202-.124.375-.296.72-.642l4.126-4.125c.346-.346.518-.52.642-.721c.11-.18.19-.375.24-.579c.04-.172.051-.352.054-.638M13 20v-5.4c0-.56 0-.84.109-1.054a1 1 0 0 1 .437-.437C13.76 13 14.04 13 14.6 13H20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteEdit(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 4H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 5.52 4 6.08 4 7.2v9.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.987.218-2.105V14m-4-9l-6 6v3h3l6-6m-3-3l3-3l3 3l-3 3m-3-3l3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteSearch(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19l-3-3m-6 4H7.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C4 18.48 4 17.92 4 16.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105V10m-6.5 7a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notebook(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h-.8c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 5.52 4 6.08 4 7.2v9.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218H8M8 4h8.8c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v9.607c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218H8M8 4v16m4-9h4m-4-3h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notification(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2.02 2.02 0 0 1-2.01-2h4a2.02 2.02 0 0 1-.15.78a2.042 2.042 0 0 1-1.44 1.18h-.047A1.922 1.922 0 0 1 12 22Zm8-3H4v-2l2-1v-5.5a8.065 8.065 0 0 1 .924-4.06A4.654 4.654 0 0 1 10 4.18V2h4v2.18c2.579.614 4 2.858 4 6.32V16l2 1v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationActive(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2.02 2.02 0 0 1-2.01-2h4a2.02 2.02 0 0 1-.149.78a2.043 2.043 0 0 1-1.44 1.18h-.046A1.944 1.944 0 0 1 12 22Zm8-3H4v-2l2-1v-5.5a8.065 8.065 0 0 1 .924-4.06A4.654 4.654 0 0 1 10 4.18V2h4v2.18c2.579.614 4 2.858 4 6.32V16l2 1v2Zm1.97-9h-2a8.672 8.672 0 0 0-3.39-6.57L18 2a9.9 9.9 0 0 1 2.825 3.486A11.52 11.52 0 0 1 21.97 10ZM4 10H2c.07-1.567.46-3.103 1.145-4.514A9.9 9.9 0 0 1 5.97 2l1.42 1.43A8.67 8.67 0 0 0 4 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationDeactivated(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2 2 0 0 1-2-2h4a2 2 0 0 1-2 2Zm7.585-.1l-2.9-2.9H4v-2l2-1v-5.5a10.57 10.57 0 0 1 .182-2L2.615 4.929L4.03 3.515L21 20.487L19.585 21.9ZM18 14.659L8.28 4.938A4.946 4.946 0 0 1 10 4.18V2h4v2.18c2.58.613 4 2.858 4 6.32v4.16v-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationDot(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2 2 0 0 1-2-2h4a2 2 0 0 1-2 2Zm8-3H4v-2l2-1v-5.5a8.065 8.065 0 0 1 .924-4.06A4.654 4.654 0 0 1 10 4.18V2h3.646a4.5 4.5 0 0 0 4.3 7.4c.035.357.052.727.052 1.1V16l2 1v2H20ZM17 8a3 3 0 1 1 .01-6A3 3 0 0 1 17 8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2.02 2.02 0 0 1-2.01-2h4a2.02 2.02 0 0 1-.15.78a2.042 2.042 0 0 1-1.44 1.18a1.758 1.758 0 0 1-.4.04Zm8-3H4v-2l2-1v-5.5a8.065 8.065 0 0 1 .924-4.06A4.654 4.654 0 0 1 10 4.18V2h4v2.18h.011c.025 0 .049.01.073.016C16.611 4.845 18 7.082 18 10.5V16l2 1v2ZM9 11v2h6v-2H9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2 2 0 0 1-2-2h4a2 2 0 0 1-2 2Zm8-3H4v-2l2-1v-5.5c0-3.462 1.421-5.707 4-6.32V2h4v2.18c2.579.612 4 2.856 4 6.32V16l2 1v2ZM12 5.75A3.6 3.6 0 0 0 8.875 7.2A5.692 5.692 0 0 0 8 10.5V17h8v-6.5a5.693 5.693 0 0 0-.875-3.3A3.6 3.6 0 0 0 12 5.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOutlineDot(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2 2 0 0 1-2-2h4a2 2 0 0 1-2 2Zm8-3H4v-2l2-1v-5.5c0-3.462 1.421-5.707 4-6.32V2h3a4.955 4.955 0 0 0-1 3c0 .251.019.502.056.751H12A3.6 3.6 0 0 0 8.875 7.2A5.692 5.692 0 0 0 8 10.5V17h8v-6.5c0-.211-.007-.414-.021-.6a5.044 5.044 0 0 0 2.006.007c.011.211.015.412.015.6V16l2 1v2ZM17 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOutlineMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2 2 0 0 1-2-2h4a2 2 0 0 1-2 2Zm8-3H4v-2l2-1v-5.5c0-3.462 1.421-5.707 4-6.32V2h4v2.18h.041l.042.01C16.611 4.843 18 7.08 18 10.5V16l2 1v2ZM12 5.75A3.6 3.6 0 0 0 8.875 7.2A5.692 5.692 0 0 0 8 10.5V17h8v-6.5a5.693 5.693 0 0 0-.875-3.3A3.6 3.6 0 0 0 12 5.75ZM15 13H9v-2h6v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOutlinePlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2 2 0 0 1-2-2h4a2 2 0 0 1-2 2Zm8-3H4v-2l2-1v-5.5c0-3.462 1.421-5.707 4-6.32V2h4v2.18c2.579.612 4 2.856 4 6.32V16l2 1v2ZM12 5.75A3.6 3.6 0 0 0 8.875 7.2A5.692 5.692 0 0 0 8 10.5V17h8v-6.5a5.693 5.693 0 0 0-.875-3.3A3.6 3.6 0 0 0 12 5.75ZM13 15h-2v-2H9v-2h2V9h2v2h2v2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationPlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2.02 2.02 0 0 1-2.01-2h4a2.02 2.02 0 0 1-.15.78a2.042 2.042 0 0 1-1.44 1.18a1.758 1.758 0 0 1-.4.04Zm8-3H4v-2l2-1v-5.5a8.065 8.065 0 0 1 .924-4.06A4.654 4.654 0 0 1 10 4.18V2h4v2.18h.011c.025 0 .049.01.073.016C16.611 4.845 18 7.082 18 10.5V16l2 1v2ZM9 11v2h2v2h2v-2h2v-2h-2V9h-2v2H9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Octagon(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.405 21H9.594c-.49 0-.734 0-.964-.056a1.996 1.996 0 0 1-.578-.239c-.202-.123-.374-.296-.72-.642L3.938 16.67c-.346-.346-.52-.519-.643-.72a1.999 1.999 0 0 1-.24-.579C3 15.14 3 14.895 3 14.405V9.594c0-.49 0-.734.055-.964c.05-.204.13-.399.24-.578c.122-.2.293-.37.633-.71l.01-.01l3.394-3.394l.01-.01c.34-.34.51-.51.71-.633c.179-.11.374-.19.578-.24C8.86 3 9.105 3 9.594 3h4.811c.49 0 .735 0 .965.055c.204.05.4.13.578.24c.202.124.375.297.72.643l3.395 3.394c.346.346.519.518.642.72c.11.179.19.374.24.578c.055.23.055.475.055.964v4.811m0 .001c0 .49 0 .734-.056.964a1.98 1.98 0 0 1-.239.578c-.123.202-.296.375-.642.72l-3.394 3.395c-.346.346-.519.519-.72.642c-.18.11-.375.19-.579.24c-.23.055-.475.055-.964.055"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 10l-4 4l-2-2M7.332 3.938L3.938 7.332l-.004.003c-.343.343-.516.516-.64.717c-.109.179-.19.374-.239.578C3 8.86 3 9.105 3 9.594v4.812c0 .489 0 .733.055.964a2 2 0 0 0 .24.578c.124.201.297.375.643.72l3.394 3.395c.346.346.518.518.72.642c.179.11.374.19.578.24c.23.055.474.055.962.055h4.816c.487 0 .732 0 .962-.055c.204-.05.4-.13.578-.24c.202-.124.375-.296.72-.642l3.395-3.394c.346-.346.519-.52.642-.721c.11-.18.19-.374.24-.578c.055-.23.055-.475.055-.964V9.594c0-.489 0-.733-.056-.964a2 2 0 0 0-.239-.578c-.123-.202-.296-.374-.642-.72L16.67 3.938c-.346-.346-.519-.52-.72-.643a1.999 1.999 0 0 0-.579-.24C15.14 3 14.895 3 14.405 3H9.594c-.49 0-.734 0-.964.055a2.02 2.02 0 0 0-.578.24c-.2.122-.37.293-.71.633z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonHelp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.146 9.074a2.998 2.998 0 0 1 5.28-.838A3 3 0 0 1 12 13v1m.05 3v.1h-.1V17zM7.332 3.938L3.938 7.332l-.004.003c-.343.343-.516.516-.64.717c-.109.179-.19.374-.239.578C3 8.86 3 9.105 3 9.594v4.812c0 .489 0 .733.055.964a2 2 0 0 0 .24.578c.124.201.297.375.643.72l3.394 3.395c.346.346.518.518.72.642c.179.11.374.19.578.24c.23.055.474.055.962.055h4.816c.487 0 .732 0 .962-.055c.204-.05.4-.13.578-.24c.202-.124.375-.296.72-.642l3.395-3.394c.346-.346.519-.52.642-.721c.11-.18.19-.374.24-.578c.055-.23.055-.475.055-.964V9.594c0-.489 0-.733-.056-.964a2 2 0 0 0-.239-.578c-.123-.202-.296-.374-.642-.72L16.67 3.938c-.346-.346-.519-.52-.72-.643a1.999 1.999 0 0 0-.579-.24C15.14 3 14.895 3 14.405 3H9.594c-.49 0-.734 0-.964.055a2.02 2.02 0 0 0-.578.24c-.2.122-.37.293-.71.633z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonWarning(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8.45v4M7.332 3.938L3.938 7.332l-.004.003c-.343.343-.516.516-.64.717c-.109.179-.19.374-.239.578C3 8.86 3 9.105 3 9.594v4.812c0 .489 0 .733.055.964a2 2 0 0 0 .24.578c.124.201.297.375.643.72l3.394 3.395c.346.346.518.518.72.642c.179.11.374.19.578.24c.23.055.474.055.962.055h4.816c.487 0 .732 0 .962-.055c.204-.05.4-.13.578-.24c.202-.124.375-.296.72-.642l3.395-3.394c.346-.346.519-.52.642-.721c.11-.18.19-.374.24-.578c.055-.23.055-.475.055-.964V9.594c0-.489 0-.733-.056-.964a2 2 0 0 0-.239-.578c-.123-.202-.296-.374-.642-.72L16.67 3.938c-.346-.346-.519-.52-.72-.643a1.999 1.999 0 0 0-.579-.24C15.14 3 14.895 3 14.405 3H9.594c-.49 0-.734 0-.964.055a2.02 2.02 0 0 0-.578.24c-.2.122-.37.293-.71.633zM12.05 15.45v.1h-.1v-.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OffClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22A9.99 9.99 0 0 1 2 12v-.2a10.005 10.005 0 0 1 17.074-6.874A10 10 0 0 1 12 22Zm0-8.59L14.59 16L16 14.59L13.41 12L16 9.41L14.59 8L12 10.59L9.41 8L8 9.41L10.59 12L8 14.59L9.41 16L12 13.411v-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OffOutlineClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22A10 10 0 0 1 4.926 4.926a10.004 10.004 0 1 1 14.148 14.148A9.936 9.936 0 0 1 12 22Zm-8-9.828A8 8 0 1 0 4 12v.172ZM9.409 16l-1.41-1.41L10.59 12L8 9.41L9.41 8L12 10.59L14.59 8L16 9.41L13.41 12L16 14.59L14.592 16L12 13.41L9.41 16h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Option(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7h5.094c.33 0 .495 0 .643.047c.132.042.253.111.357.202c.117.103.202.245.372.528l5.068 8.446c.17.284.255.425.372.528c.103.09.224.16.356.202c.148.047.314.047.644.047H21M15 7h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlane(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.308 13.692l4.846-4.846M20.11 5.89l-4.09 13.294c-.367 1.192-.55 1.788-.867 1.985a.999.999 0 0 1-.912.076c-.344-.143-.624-.7-1.182-1.816l-2.59-5.182a2.104 2.104 0 0 0-.193-.342a1.002 1.002 0 0 0-.18-.181a2.036 2.036 0 0 0-.331-.186L4.572 10.94c-1.115-.558-1.673-.837-1.816-1.181a1 1 0 0 1 .076-.913c.197-.316.793-.5 1.985-.867l13.295-4.09c.937-.289 1.405-.433 1.722-.316a1 1 0 0 1 .594.594c.116.316-.028.784-.316 1.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperclipAttechmentHorizontal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6h9.75a5.25 5.25 0 1 1 0 10.5H5.5a3.5 3.5 0 1 1 0-7h11.25a1.75 1.75 0 1 1 0 3.5H7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperclipAttechmentTilt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.536 11.465L11.43 4.57a5.25 5.25 0 1 1 7.424 7.425L10.9 19.95A3.5 3.5 0 0 1 5.95 15l7.956-7.954A1.75 1.75 0 0 1 16.38 9.52l-6.895 6.894"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v8m0-8h4m-4 0h-1a4 4 0 1 0 0 8h1m0 0v6m4-14v14m0-14h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Path(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.121 15.879a3 3 0 1 0-4.243 4.243a3 3 0 0 0 4.243-4.243m0 0L15.88 8.12m0 0a3 3 0 1 0 4.243-4.243A3 3 0 0 0 15.88 8.12m0 0l.004-.004"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 5.5v13c0 .465 0 .697.038.89a2 2 0 0 0 1.571 1.572c.194.038.426.038.89.038c.465 0 .698 0 .892-.038a2 2 0 0 0 1.57-1.572c.039-.19.039-.42.039-.878V5.488c0-.457 0-.687-.038-.879a2 2 0 0 0-1.572-1.57C18.197 3 17.965 3 17.5 3s-.697 0-.89.038a1.999 1.999 0 0 0-1.572 1.571C15 4.803 15 5.035 15 5.5m-11 0v13c0 .465 0 .697.038.89a2 2 0 0 0 1.571 1.572c.194.038.426.038.89.038c.465 0 .698 0 .892-.038a2 2 0 0 0 1.57-1.572C9 19.2 9 18.97 9 18.512V5.488c0-.457 0-.687-.038-.879A2 2 0 0 0 7.39 3.04C7.197 3 6.965 3 6.5 3s-.697 0-.89.038A1.999 1.999 0 0 0 4.037 4.61C4 4.803 4 5.035 4 5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 9v6m-4-6v6m2 6a9 9 0 1 1 0-18a9 9 0 0 1 0 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleFilled(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm1-14v8h2V8h-2ZM9 8v8h2V8H9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircleOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c-5.52-.006-9.994-4.48-10-10v-.2C2.11 6.305 6.635 1.928 12.13 2c5.497.074 9.904 4.569 9.868 10.065C21.962 17.562 17.497 22 12 22Zm-.016-2H12a8 8 0 1 0-.016 0ZM15 16h-2V8h2v8Zm-4 0H9V8h2v8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.398 21.009H8.874a.392.392 0 0 1-.33-.15a.459.459 0 0 1-.09-.363c.04-.241.09-.559.152-.963l.114-.727a313.336 313.336 0 0 1 .348-2.188v-.02l.188-1.164l.1-.619v-.005c.072-.445.129-.798.17-1.061a.39.39 0 0 1 .433-.371h1.52a10.08 10.08 0 0 0 2.171-.212a6.09 6.09 0 0 0 2.886-1.449a6.084 6.084 0 0 0 1.56-2.473c.152-.436.27-.884.352-1.338c.007-.042.015-.066.025-.074a.03.03 0 0 1 .025-.012h.01a.31.31 0 0 1 .062.035c.525.397.876.982.98 1.632c.118.68.105 1.376-.04 2.051a5.406 5.406 0 0 1-1.857 3.35a6.053 6.053 0 0 1-3.825 1.116h-.439a.662.662 0 0 0-.444.166a.722.722 0 0 0-.239.427l-.04.194l-.554 3.478l-.02.151a.706.706 0 0 1-.249.427a.67.67 0 0 1-.445.162Zm-3.578-2H4.855a.473.473 0 0 1-.369-.165a.466.466 0 0 1-.115-.39l2.331-14.79a.775.775 0 0 1 .277-.483a.783.783 0 0 1 .518-.19h6.014c.33.013.658.057.98.131c.382.073.758.18 1.122.32a3.425 3.425 0 0 1 1.645 1.238a3.41 3.41 0 0 1 .568 1.972a7.11 7.11 0 0 1-.46 2.374a4.909 4.909 0 0 1-3.042 3.165a8.046 8.046 0 0 1-2.535.425c-.01.006-.44.007-.9.007l-.9-.007a1.083 1.083 0 0 0-1.187.964c-.013.054-.317 1.947-.855 5.329a.107.107 0 0 1-.128.105l.001-.005Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.502 4.257A2 2 0 0 0 7.646 3H4.895A1.895 1.895 0 0 0 3 4.895C3 13.789 10.21 21 19.106 21A1.895 1.895 0 0 0 21 19.105v-2.751a2 2 0 0 0-1.257-1.857l-2.636-1.054a2 2 0 0 0-2.023.32l-.68.568a2.001 2.001 0 0 1-2.696-.122L9.792 12.29a2 2 0 0 1-.123-2.694l.567-.68a2 2 0 0 0 .322-2.024z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.742 8.682c-1.073.912-1.466 2.575-.54 3.909a12.818 12.818 0 0 0 3.208 3.207c1.334.926 2.997.533 3.909-.54l.01.005A13.38 13.38 0 0 0 19 16.36V19h-.004C11.022 19.011 4.991 12.911 5 5.004V5h2.64v.001c.188 1.27.558 2.506 1.097 3.67l.005.01ZM19 21h1a1 1 0 0 0 1-1v-4.502a1 1 0 0 0-.853-.99l-.854-.126a11.416 11.416 0 0 1-3.123-.934l-.753-.349a1 1 0 0 0-1.234.326l-.341.477c-.299.419-.87.546-1.291.253a10.819 10.819 0 0 1-2.706-2.705c-.293-.422-.165-.993.253-1.291l.477-.34a1 1 0 0 0 .326-1.235l-.35-.754a11.424 11.424 0 0 1-.933-3.123l-.127-.854A1 1 0 0 0 8.501 3H4a1 1 0 0 0-1 1v1.001C2.99 14.008 9.91 21.013 18.999 21Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartFifty(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-.016-2H12V4a8 8 0 1 0-.016 16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10ZM11 4.062a8 8 0 1 0 5.419 14.608l-.1.071l.094-.065l.059-.041l.064-.045l.016-.011l.009-.007l-5.128-5.13A1.51 1.51 0 0 1 11 12.379V4.062ZM13.829 13l4.227 4.227l.007-.008l.005-.006l-.01.011A7.944 7.944 0 0 0 19.938 13h-6.109ZM13 4.062V11h6.938A8 8 0 0 0 13 4.062Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartOutlineTwentyFive(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10ZM11 4.062a8 8 0 1 0 8.915 9.1V13H11V4.062Zm2 0V11h6.938A8 8 0 0 0 13 4.062Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartSeventyFive(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c-5.52-.006-9.994-4.48-10-10v-.2C2.11 6.305 6.635 1.928 12.13 2c5.497.074 9.904 4.569 9.868 10.065C21.962 17.562 17.497 22 12 22Zm0-18a8.008 8.008 0 0 0-8 8h8V4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartTwentyFive(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c-5.52-.006-9.994-4.48-10-10v-.2C2.11 6.305 6.635 1.928 12.13 2c5.497.074 9.904 4.569 9.868 10.065C21.962 17.562 17.497 22 12 22Zm-8-9.828A8 8 0 1 0 20 12h-8V4a8.01 8.01 0 0 0-8 8v.172Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Planet(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.898 18.162a8 8 0 0 0 12.93-7.821m-12.93 7.82a8 8 0 1 1 12.93-7.821m-12.93 7.822c1.955-.447 4.282-1.38 6.628-2.734c2.73-1.576 4.951-3.416 6.302-5.088m-12.93 7.822c-2.43.555-4.286.362-4.898-.698c-.63-1.091.196-2.867 2-4.755m15.828-2.369c1.252-1.55 1.756-2.956 1.224-3.876c-.539-.933-2.043-1.195-4.052-.865"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17.334V6.667c0-.88 0-1.32.185-1.58a1 1 0 0 1 .687-.412c.317-.04.705.166 1.48.58l10 5.333l.004.002c.857.457 1.286.686 1.427.99c.122.266.122.573 0 .839c-.141.305-.571.535-1.43.993l-10 5.333c-.777.414-1.164.62-1.48.58a1 1 0 0 1-.688-.412C5 18.653 5 18.213 5 17.333"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayArrow(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 19l11-7L8 5v14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0"/><path d="M10 15V9l5 3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleFilled(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c-5.52-.006-9.994-4.48-10-10v-.2C2.11 6.305 6.635 1.928 12.13 2c5.497.074 9.904 4.569 9.868 10.065C21.962 17.562 17.497 22 12 22ZM10 7.5v9l6-4.5l-6-4.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-8-9.828A8 8 0 1 0 4 12v.172Zm6 4.328v-9l6 4.5l-6 4.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayStore(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.4 21c-.17.003-.34-.026-.5-.086l8.054-8.057l2.666 2.669l-9.255 5.2A1.998 1.998 0 0 1 5.4 21Zm-1.164-.665a1.9 1.9 0 0 1-.236-.97V4.66a2.13 2.13 0 0 1 .1-.658l8.233 8.235l-8.1 8.1l.003-.002Zm12.179-5.258l-2.841-2.839l3.133-3.132l2.783 1.563c.534.24.892.755.928 1.339a1.574 1.574 0 0 1-.929 1.34l-3.074 1.729Zm-3.461-3.463l-8.34-8.339c.229-.17.506-.26.791-.261c.336.012.664.107.955.277l9.551 5.368l-2.956 2.955h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 13v6h-2v-6H5v-2h6V5h2v6h6v2h-6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c-5.52-.006-9.994-4.48-10-10v-.2C2.11 6.305 6.635 1.928 12.13 2c5.497.074 9.904 4.569 9.868 10.065C21.962 17.562 17.497 22 12 22ZM7 11v2h4v4h2v-4h4v-2h-4V7h-2v4H7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-8-9.828A8 8 0 1 0 4 12v.172ZM13 17h-2v-4H7v-2h4V7h2v4h4v2h-4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2ZM5 5v14h14V5H5Zm8 12h-2v-4H7v-2h4V7h2v4h4v2h-4v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 17c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C21 15.398 21 14.932 21 14v-3.803c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C19.506 7.013 18.988 7.001 18 7m0 0H6m12 0H6m12 0v-.803c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C16.48 3 15.92 3 14.8 3H9.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C6 4.52 6 5.08 6 6.2V7m0 0c-.988 0-1.507.013-1.908.218a1.999 1.999 0 0 0-.874.874C3 8.52 3 9.08 3 10.2V14c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.082 1.083C4.602 17 5.068 17 6 17m4-6h4m-5 4h6c.932 0 1.398 0 1.765.152a2 2 0 0 1 1.083 1.082C18 16.602 18 17.068 18 18s0 1.398-.152 1.765a2 2 0 0 1-1.083 1.083C16.398 21 15.932 21 15 21H9c-.932 0-1.398 0-1.766-.152a1.999 1.999 0 0 1-1.082-1.082C6 19.398 6 18.932 6 18s0-1.398.152-1.766a2 2 0 0 1 1.082-1.082C7.602 15 8.068 15 9 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7h-2.151C17.35 7 17 6.498 17 6a3 3 0 1 0-6 0c0 .498-.351 1-.849 1H8a1 1 0 0 0-1 1v2.151C7 10.65 6.498 11 6 11a3 3 0 1 0 0 6c.498 0 1 .351 1 .849V20a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2.151c0-.498-.503-.849-1-.849a3 3 0 1 1 0-6c.497 0 1-.351 1-.849V8a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20h1m-4 0h-2v-3m3 0h3v-3h-1m-5 0h2M4 17c0-.932 0-1.398.152-1.766a2 2 0 0 1 1.082-1.082C5.602 14 6.068 14 7 14s1.398 0 1.766.152c.49.203.879.592 1.082 1.082c.152.368.152.834.152 1.766s0 1.398-.152 1.765a2 2 0 0 1-1.082 1.083C8.398 20 7.932 20 7 20s-1.398 0-1.766-.152a1.999 1.999 0 0 1-1.082-1.082C4 18.398 4 17.932 4 17M14 7c0-.932 0-1.398.152-1.766a2 2 0 0 1 1.082-1.082C15.602 4 16.068 4 17 4s1.398 0 1.766.152c.49.203.879.592 1.082 1.082C20 5.602 20 6.068 20 7s0 1.398-.152 1.765a2 2 0 0 1-1.082 1.083C18.398 10 17.932 10 17 10s-1.398 0-1.766-.152a1.999 1.999 0 0 1-1.082-1.082C14 8.398 14 7.932 14 7M4 7c0-.932 0-1.398.152-1.766a2 2 0 0 1 1.082-1.082C5.602 4 6.068 4 7 4s1.398 0 1.766.152c.49.203.879.592 1.082 1.082C10 5.602 10 6.068 10 7s0 1.398-.152 1.765a2 2 0 0 1-1.082 1.083C8.398 10 7.932 10 7 10s-1.398 0-1.766-.152a1.999 1.999 0 0 1-1.082-1.082C4 8.398 4 7.932 4 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCodeOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 11V3h8v8H3Zm3-5v2h2V6H6Zm7 5V3h8v8h-8Zm3-5v2h2V6h-2ZM3 13v8h8v-8H3Zm5 3v2H6v-2h2Zm8-3h-3v2h3v4h-1v-2h-2v4h4v-2h2v2h2v-3h-2v-1h-1v-1h3v-3h-3v2h-2v-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19a7 7 0 1 1 7-7a7.008 7.008 0 0 1-7 7Zm0-12a5 5 0 1 0 0 10a5 5 0 0 0 0-10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioFill(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16"/><path d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioFilled(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19a7 7 0 1 1 7-7a7.008 7.008 0 0 1-7 7Zm0-12a5 5 0 1 0 0 10a5 5 0 0 0 0-10Zm0 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioUnchecked(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rainbow(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17v-2a9 9 0 1 1 18 0v2M6 17v-2a6 6 0 0 1 12 0v2m-9 0v-2a3 3 0 1 1 6 0v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10ZM6.807 10.543a1.46 1.46 0 0 0-.608 2.787a2.562 2.562 0 0 0 0 .439c0 2.24 2.615 4.062 5.829 4.062c3.214 0 5.83-1.822 5.83-4.062a2.773 2.773 0 0 0 0-.439a1.441 1.441 0 0 0-.648-2.737h-.053A1.456 1.456 0 0 0 16.2 11a7.116 7.116 0 0 0-3.85-1.23L13 6.65l2.138.45a1 1 0 0 0 1.102.886a1 1 0 0 0-.1-1.995a.847.847 0 0 0-.113.009a1 1 0 0 0-.759.492L12.82 6a.31.31 0 0 0-.366.237l-.748 3.473a7.123 7.123 0 0 0-3.9 1.229a1.45 1.45 0 0 0-.999-.396Zm5.373 5.981h-.338a3.852 3.852 0 0 1-2.3-.773a.269.269 0 0 1 .379-.38c.56.41 1.238.631 1.933.629h.324a3.268 3.268 0 0 0 1.912-.614a.28.28 0 0 1 .4 0a.284.284 0 0 1-.005.4v-.04a3.837 3.837 0 0 1-2.305.777v.001Zm2.127-2.444h-.016l.008-.039a.939.939 0 1 1 .736-.281a.988.988 0 0 1-.694.32h-.035h.001ZM9.67 14a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 8h5V3m-.29 13.357a8 8 0 1 1-.19-8.991"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.995 4a8 8 0 1 0 7.735 10h-2.081a6 6 0 1 1-5.654-8a5.92 5.92 0 0 1 4.223 1.78L13 11h7V4l-2.351 2.35A7.965 7.965 0 0 0 11.995 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 20v-7h7l-3.217 3.22A5.917 5.917 0 0 0 12 18a6 6 0 0 0 5.648-4h.018c.114-.325.201-.66.259-1h2.012A8 8 0 0 1 12 20h-.01a7.941 7.941 0 0 1-5.653-2.34L4 20Zm2.074-9H4.062a8 8 0 0 1 7.933-7H12a7.94 7.94 0 0 1 5.654 2.34L20 4v7h-7l3.222-3.22A5.916 5.916 0 0 0 12 6a6 6 0 0 0-5.648 4h-.018c-.115.325-.202.66-.259 1h-.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveMinusCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h8m-4 9a9 9 0 1 1 0-18a9 9 0 0 1 0 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7 22l-4-4l4-4v3h10v-4h2v5a1 1 0 0 1-1 1H7v3Zm0-11H5V6a1 1 0 0 1 1-1h11V2l4 4l-4 4V7H7v4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rewind(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 12l9 5V7zm0 0V7l-9 5l9 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rows(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.5 19h11c.465 0 .697 0 .89-.039a2 2 0 0 0 1.571-1.57c.039-.194.039-.426.039-.891s0-.697-.039-.89a2 2 0 0 0-1.57-1.572C18.196 14 17.964 14 17.5 14h-11c-.465 0-.697 0-.89.038a2 2 0 0 0-1.572 1.572C4 15.803 4 16.035 4 16.5s0 .697.038.89a2 2 0 0 0 1.572 1.571c.193.039.425.039.89.039m0-9h11c.465 0 .697 0 .89-.039a2 2 0 0 0 1.571-1.57C20 8.196 20 7.964 20 7.5s0-.697-.039-.89a2 2 0 0 0-1.57-1.572C18.196 5 17.964 5 17.5 5h-11c-.465 0-.697 0-.89.038A2 2 0 0 0 4.038 6.61C4 6.803 4 7.035 4 7.5s0 .697.038.89A2 2 0 0 0 5.61 9.961c.193.039.425.039.89.039"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.344 12l2.121 2.122m.707-4.95l2.121 2.12M12 6.344l2.122 2.122M5.07 13.273l8.202-8.203c.792-.792 1.188-1.188 1.645-1.336a2 2 0 0 1 1.236 0c.456.148.852.544 1.643 1.335L18.93 6.2c.792.792 1.188 1.19 1.337 1.646c.13.402.13.834 0 1.235c-.149.457-.545.853-1.337 1.645l-8.202 8.203c-.792.792-1.189 1.188-1.645 1.336c-.402.13-.834.13-1.235 0c-.457-.148-.854-.544-1.646-1.336l-1.133-1.133c-.79-.791-1.187-1.187-1.335-1.643a2 2 0 0 1 0-1.236c.148-.457.545-.853 1.337-1.645"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sad(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm0-18a8 8 0 1 0 8 8a8.009 8.009 0 0 0-8-8Zm-2 14c-.014 0-.14-.005-1-.005H8v-.066c0-.033 0-.078.007-.133a4.7 4.7 0 0 1 .472-1.743a3.6 3.6 0 0 1 1.23-1.414l.014-.009l.016-.012l.015-.008h.025l.011-.007A4.117 4.117 0 0 1 12 14a4.06 4.06 0 0 1 2.29.635c.527.355.951.843 1.23 1.414c.204.414.346.855.419 1.311c.032.188.046.339.052.432c0 .044.006.088.007.133v.062h-2v-.059c0-.055-.013-.14-.031-.246a2.698 2.698 0 0 0-.236-.747a1.638 1.638 0 0 0-.551-.645A2.11 2.11 0 0 0 12 16a2.11 2.11 0 0 0-1.18.3a1.647 1.647 0 0 0-.551.645a2.716 2.716 0 0 0-.266.993v.058H10V18Zm-1.5-6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm6.993-.014a1.493 1.493 0 1 1 0-2.986a1.493 1.493 0 0 1 0 2.986Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 21H7m10 0h.803c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.427.218-.987.218-2.105V9.22c0-.45 0-.675-.048-.889a1.994 1.994 0 0 0-.209-.545c-.106-.19-.256-.355-.55-.682l-2.755-3.062c-.341-.378-.514-.57-.721-.708a1.999 1.999 0 0 0-.61-.271C15.863 3 15.6 3 15.075 3H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 4.52 3 5.08 3 6.2v11.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218H7m10 0v-3.803c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C15.48 14 14.92 14 13.8 14h-3.6c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C7 15.52 7 16.08 7 17.2V21m8-14H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.677 19.607l-5.715-5.716a6 6 0 0 1-7.719-9.133a6 6 0 0 1 9.134 7.718l5.715 5.716l-1.414 1.414l-.001.001ZM9.485 5a4 4 0 1 0 2.917 1.264l.605.6l-.682-.68l-.012-.012A3.972 3.972 0 0 0 9.485 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchMagnifyingGlass(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 15l6 6m-11-4a7 7 0 1 1 0-14a7 7 0 0 1 0 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSmall(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.577 19l-4.767-4.766a5.035 5.035 0 0 1-6.336-7.76a5.035 5.035 0 0 1 7.761 6.335L19 17.577L17.577 19ZM10.034 7.014a3.02 3.02 0 1 0-.004 6.04a3.02 3.02 0 0 0 .004-6.04Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSmallMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.577 20l-5.767-5.766a5.035 5.035 0 0 1-6.336-7.76a5.035 5.035 0 0 1 7.761 6.335L18 18.576L16.577 20ZM8.034 7.014a3.02 3.02 0 1 0-.004 6.04a3.02 3.02 0 0 0 .004-6.04ZM21 9h-6V7h6v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSmallPlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.577 20l-5.767-5.766a5.035 5.035 0 0 1-6.336-7.76a5.035 5.035 0 0 1 7.761 6.335L18 18.576L16.577 20ZM8.034 7.014a3.02 3.02 0 1 0-.004 6.04a3.02 3.02 0 0 0 .004-6.04ZM19 11h-2V9h-2V7h2V5h2v2h2v2h-2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectMultiple(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9v10.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C3.76 21 4.04 21 4.598 21H15m2-13l-4 4l-2-2m-4 3.8V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C8.52 3 9.08 3 10.2 3h7.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C21 4.52 21 5.08 21 6.2v7.6c0 1.12 0 1.68-.218 2.108a2.001 2.001 0 0 1-.874.874c-.428.218-.986.218-2.104.218h-7.607c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C7 15.48 7 14.92 7 13.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m20.35 8.923l-.366-.204a2 2 0 0 1-.784-.724c-.017-.027-.033-.056-.065-.112a2.002 2.002 0 0 1-.3-1.157l.006-.425c.012-.68.018-1.022-.078-1.328a1.998 1.998 0 0 0-.417-.736c-.214-.24-.511-.412-1.106-.754l-.494-.285c-.592-.341-.889-.512-1.204-.577a1.999 1.999 0 0 0-.843.007c-.313.07-.606.246-1.191.596l-.003.002l-.354.211c-.056.034-.085.05-.113.066c-.278.155-.588.24-.907.25c-.032.002-.065.002-.13.002l-.13-.001a1.997 1.997 0 0 1-.91-.252c-.028-.015-.055-.032-.111-.066l-.357-.214c-.589-.354-.884-.53-1.199-.601a1.998 1.998 0 0 0-.846-.006c-.316.066-.612.238-1.205.582l-.003.001l-.488.283l-.005.004c-.588.34-.883.512-1.095.751a2 2 0 0 0-.415.734c-.095.307-.09.649-.078 1.333l.007.424c0 .065.003.097.002.128a2.002 2.002 0 0 1-.301 1.027c-.033.056-.048.084-.065.11a2 2 0 0 1-.675.664l-.112.063l-.361.2c-.602.333-.903.5-1.121.738a2 2 0 0 0-.43.73c-.1.307-.1.65-.099 1.338l.002.563c.001.683.003 1.024.104 1.329a2 2 0 0 0 .427.726c.218.236.516.402 1.113.734l.358.199c.061.034.092.05.121.068a2 2 0 0 1 .74.781l.067.12a2 2 0 0 1 .23 1.038l-.007.407c-.012.686-.017 1.03.079 1.337c.085.272.227.523.417.736c.214.24.512.411 1.106.754l.494.285c.593.341.889.512 1.204.577a2 2 0 0 0 .843-.007c.314-.07.607-.246 1.194-.598l.354-.212a1.997 1.997 0 0 1 1.02-.317h.26c.318.01.63.097.91.252l.092.055l.376.226c.59.354.884.53 1.199.6a2 2 0 0 0 .846.008c.315-.066.613-.239 1.206-.583l.495-.287c.588-.342.883-.513 1.095-.752c.19-.213.33-.463.415-.734c.095-.305.09-.644.078-1.318l-.008-.44a2 2 0 0 1 .3-1.155l.065-.11a2 2 0 0 1 .675-.664l.11-.061l.002-.001l.361-.2c.602-.334.903-.5 1.122-.738c.194-.21.34-.46.429-.73c.1-.305.1-.647.098-1.327l-.002-.574c-.001-.683-.002-1.025-.103-1.33a2.002 2.002 0 0 0-.428-.725c-.217-.236-.515-.402-1.111-.733z"/><path d="M8 12a4 4 0 1 0 8 0a4 4 0 0 0-8 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsFilled(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.82 22h-3.64a1 1 0 0 1-.977-.786l-.407-1.884a8.002 8.002 0 0 1-1.535-.887l-1.837.585a1 1 0 0 1-1.17-.453L2.43 15.424a1.006 1.006 0 0 1 .193-1.239l1.425-1.3a8.1 8.1 0 0 1 0-1.772L2.623 9.816a1.006 1.006 0 0 1-.193-1.24l1.82-3.153a1 1 0 0 1 1.17-.453l1.837.585c.244-.18.498-.348.76-.5c.253-.142.513-.271.779-.386l.408-1.882A1 1 0 0 1 10.18 2h3.64a1 1 0 0 1 .976.787l.412 1.883a8.192 8.192 0 0 1 1.535.887l1.838-.585a1 1 0 0 1 1.169.453l1.82 3.153c.232.407.152.922-.193 1.239l-1.425 1.3a8.1 8.1 0 0 1 0 1.772l1.425 1.3c.345.318.425.832.193 1.239l-1.82 3.153a1 1 0 0 1-1.17.453l-1.837-.585a7.98 7.98 0 0 1-1.534.886l-.413 1.879a1 1 0 0 1-.976.786ZM11.996 8a4 4 0 1 0 0 8a4 4 0 0 0 0-8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsFuture(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13.6 21.076l5.46-3.152c.584-.337.875-.505 1.087-.74a2 2 0 0 0 .416-.72c.097-.301.097-.637.097-1.307V8.843c0-.67 0-1.006-.098-1.307a2 2 0 0 0-.416-.72c-.21-.234-.5-.402-1.079-.736L13.6 2.924c-.583-.337-.874-.505-1.184-.57a2 2 0 0 0-.832 0c-.31.065-.601.233-1.184.57L4.938 6.077c-.582.336-.873.504-1.084.739a2 2 0 0 0-.416.72c-.098.302-.098.638-.098 1.311v6.305c0 .673 0 1.01.098 1.311a2 2 0 0 0 .416.72c.211.236.503.404 1.085.74l5.46 3.153c.584.337.875.505 1.185.57c.274.059.558.059.832 0c.31-.065.602-.233 1.185-.57"/><path d="M9 12a3 3 0 1 0 6 0a3 3 0 0 0-6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 15a3.478 3.478 0 0 0 2.357-.93l6.26 3.577a3.52 3.52 0 1 0 1.026-1.717l-6.26-3.577c.066-.25.102-.509.108-.768l6.15-3.515A3.489 3.489 0 1 0 14 5.5c.004.288.043.575.117.853L8.433 9.6A3.5 3.5 0 1 0 5.5 15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAndroid(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 13.5l6 3m0-9l-6 3M18 21a3 3 0 1 1 0-6a3 3 0 0 1 0 6M6 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6m12-6a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareIosExport(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 6l3-3m0 0l3 3m-3-3v10m-5-3c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C4 11.602 4 12.068 4 13v4.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.607c1.118 0 1.677 0 2.104-.218c.376-.192.682-.498.874-.874c.218-.428.218-.987.218-2.105V13c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C18.398 10 17.932 10 17 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 22a3.46 3.46 0 0 1-3.383-4.352l-6.26-3.578a3.494 3.494 0 1 1 .576-4.47l5.683-3.249A3.494 3.494 0 0 1 14 5.5a3.531 3.531 0 1 1 1.142 2.57l-6.15 3.515c-.007.26-.043.517-.109.768l6.26 3.577A3.495 3.495 0 1 1 17.5 22Zm0-5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-12-7a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm12-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.926 20.631C15.032 19.678 20 16.733 20 10.165V6.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 3 17.92 3 16.8 3H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 4.52 4 5.08 4 6.2v3.965c0 6.568 4.968 9.513 7.074 10.466c.223.102.335.152.588.195c.16.028.518.028.677 0c.252-.043.363-.093.585-.194z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 9l-4 4l-2-2m11-.835c0 6.568-4.968 9.513-7.074 10.466l-.003.002c-.221.1-.332.15-.584.193c-.16.028-.518.028-.677 0a2.01 2.01 0 0 1-.588-.195C8.968 19.678 4 16.733 4 10.165V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 3 6.08 3 7.2 3h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldWarning(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m8-1.835c0 6.568-4.968 9.513-7.074 10.466l-.003.002c-.221.1-.332.15-.584.193c-.16.028-.518.028-.677 0a2.01 2.01 0 0 1-.588-.195C8.968 19.678 4 16.733 4 10.165V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 3 6.08 3 7.2 3h9.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105zM12.05 15v.1h-.1V15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 8a3 3 0 1 0 6 0M3 16.8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v9.607c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218H6.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C3 18.48 3 17.92 3 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8H6.777c-1.3 0-1.949 0-2.41.265c-.406.233-.718.6-.881 1.039c-.185.499-.079 1.14.135 2.42v.002l.933 5.6c.159.95.238 1.425.475 1.782c.21.314.503.562.847.717c.39.175.872.175 1.835.175h8.578c.963 0 1.444 0 1.835-.175c.344-.155.638-.403.847-.717c.237-.357.316-.832.474-1.782l.934-5.6v-.004c.214-1.28.32-1.92.135-2.418a2 2 0 0 0-.88-1.039C19.173 8 18.523 8 17.223 8H16M8 8h8M8 8a4 4 0 1 1 8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 0H9.294c-.461 0-.692 0-.882-.082a1.002 1.002 0 0 1-.418-.337c-.12-.167-.167-.39-.261-.83L5.27 4.265c-.096-.451-.145-.677-.265-.845a1.003 1.003 0 0 0-.419-.338C4.397 3 4.167 3 3.707 3H3m3 3h12.873c.722 0 1.082 0 1.325.15a1 1 0 0 1 .435.579c.077.274-.022.621-.222 1.314l-1.385 4.8c-.12.415-.18.622-.3.776a1.004 1.004 0 0 1-.409.307c-.18.074-.396.074-.825.074H7.73M8 21a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h.268c.474 0 .711 0 .905.085c.17.076.316.197.42.351c.12.174.163.407.248.871L7 16h10.422c.453 0 .68 0 .868-.08a.998.998 0 0 0 .415-.331c.12-.165.171-.385.273-.826v-.003l1.57-6.8v-.001c.154-.669.232-1.004.147-1.267a1.001 1.001 0 0 0-.44-.55C20.019 6 19.676 6 18.99 6H5.5M18 21a1 1 0 1 1 0-2a1 1 0 0 1 0 2M8 21a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShortDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11 16.17l-3.59-3.58L6 14l6 6l6-6l-1.41-1.41L13 16.17V4h-2v12.17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShortLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.83 11l3.58-3.59L10 6l-6 6l6 6l1.41-1.41L7.83 13H20v-2H7.83Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShortRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.17 13l-3.58 3.59L14 18l6-6l-6-6l-1.41 1.41L16.17 11H4v2h12.17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShortUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 7.83l3.59 3.58L18 10l-6-6l-6 6l1.41 1.41L11 7.83V20h2V7.83Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Show(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.587 13.779c1.78 1.769 4.883 4.22 8.413 4.22c3.53 0 6.634-2.451 8.413-4.22c.47-.467.705-.7.854-1.159c.107-.327.107-.913 0-1.24c-.15-.458-.385-.692-.854-1.159C18.633 8.452 15.531 6 12 6c-3.53 0-6.634 2.452-8.413 4.221c-.47.467-.705.7-.854 1.159c-.107.327-.107.913 0 1.24c.15.458.384.692.854 1.159"/><path d="M10 12a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shrink(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 14h5v5m9-9h-5V5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 20l3-3m0 0l-3-3m3 3h-4a5 5 0 0 1-5-5a5 5 0 0 0-5-5H3m15-3l3 3m0 0l-3 3m3-3h-4a4.978 4.978 0 0 0-3 1M3 17h4a4.978 4.978 0 0 0 3-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SingleQuotesL(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 12v3.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437c.214.109.494.109 1.053.109h1.803c.559 0 .838 0 1.052-.109c.188-.096.341-.25.437-.437C15 16.24 15 15.96 15 15.4v-1.803c0-.559 0-.838-.109-1.052a1 1 0 0 0-.437-.437C14.24 12 13.96 12 13.4 12zm0 0v-2a3 3 0 0 1 3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SingleQuotesR(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17a3 3 0 0 0 3-3v-2m0 0V8.598c0-.559 0-.838-.109-1.052a1 1 0 0 0-.437-.437C14.24 7 13.96 7 13.4 7h-1.8c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C10 7.76 10 8.04 10 8.6v1.8c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437c.214.109.494.109 1.053.109z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sketch(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21h-.051a.712.712 0 0 1-.491-.253L2.168 9.673a.676.676 0 0 1-.149-.289c0-.013-.006-.026-.009-.039a.677.677 0 0 1 .1-.482l.019-.028L5.759 3.7a.7.7 0 0 1 .541-.3l5.65-.4h.099l5.651.4a.7.7 0 0 1 .535.294l3.637 5.144a.676.676 0 0 1 .106.577a.667.667 0 0 1-.142.259l-9.293 11.072a.716.716 0 0 1-.441.247A.74.74 0 0 1 12 21ZM7.946 9.911L12 18.663l4.054-8.752H7.946Zm9.64 0l-2.762 5.963l5-5.963h-2.238Zm-13.413 0l5 5.963l-2.759-5.963H4.173Zm13.995-3.917l-.26 2.551h2.062l-1.802-2.551ZM12 4.869L8.872 8.545h6.256L12 4.869ZM5.832 5.994L4.03 8.545h2.062l-.26-2.551Zm7.645-1.523l3.074 3.613l.349-3.373l-3.423-.24Zm-2.956 0l-3.421.24l.345 3.373l3.076-3.613Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipBack(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 5v14m11-8.429v2.858c0 1.827 0 2.74-.384 3.267a2 2 0 0 1-1.413.811c-.648.066-1.437-.394-3.015-1.315L10.73 14.76c-1.551-.905-2.328-1.358-2.59-1.949a2 2 0 0 1 0-1.62c.263-.591 1.041-1.045 2.598-1.954l2.45-1.428l.002-.002c1.576-.92 2.365-1.38 3.013-1.313a2 2 0 0 1 1.413.812C18 7.83 18 8.745 18 10.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipForward(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 5v14M6 10.571v2.858c0 1.827 0 2.74.384 3.267a2 2 0 0 0 1.413.811c.648.066 1.437-.394 3.016-1.315l2.449-1.428l.008-.005c1.552-.905 2.328-1.358 2.59-1.949a2 2 0 0 0 0-1.62c-.263-.591-1.041-1.045-2.598-1.954l-2.45-1.428c-1.578-.921-2.367-1.381-3.015-1.315a2 2 0 0 0-1.413.812C6 7.83 6 8.745 6 10.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipNext(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 18h-2V6h2v12ZM6 18V6l8.5 6L6 18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipPrevious(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18 18l-8.5-6L18 6v12ZM8 18H6V6h2v12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slack(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.641 22.024a2.107 2.107 0 0 1-2.1-2.106v-2.107h2.1a2.107 2.107 0 1 1 0 4.214v-.001Zm-5.282 0a2.107 2.107 0 0 1-2.1-2.106v-5.274a2.104 2.104 0 0 1 4.207 0v5.274a2.107 2.107 0 0 1-2.107 2.106Zm10.55-5.273h-5.268a2.106 2.106 0 1 1 0-4.213h5.268a2.106 2.106 0 1 1 0 4.213Zm-15.817 0a2.106 2.106 0 1 1 0-4.213h2.1v2.106a2.107 2.107 0 0 1-2.105 2.107h.005Zm15.817-5.289h-2.1V9.356a2.1 2.1 0 1 1 2.1 2.106Zm-5.268 0a2.107 2.107 0 0 1-2.1-2.106V4.082a2.103 2.103 0 1 1 4.207 0v5.274a2.107 2.107 0 0 1-2.107 2.106Zm-5.282 0H4.092a2.106 2.106 0 1 1 0-4.213H9.36a2.106 2.106 0 1 1 0 4.213Zm2.1-5.273h-2.1a2.106 2.106 0 1 1 2.1-2.107v2.107Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 15h7M3 15h2m0 0a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0m15-6h1M3 9h7m6.5 2.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderThree(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 18h11M3 18h3m0 0v2m0-2v-2m14-4h1M3 12h13m0 0v2m0-2v-2m-2-4h7M3 6h7m0 0v2m0-2V4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 18h7M3 18h2m0 0a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0m15-6h1M3 12h7m3-6h8m-8 0a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0M3 6h1m12.5 8.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallLongDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 18h-3V2h-2v16H8l4 4l4-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallLongLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 16v-3h16v-2H6V8l-4 4l4 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallLongRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 8v3H2v2h16v3l4-4l-4-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallLongUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6h3v16h2V6h3l-4-4l-4 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snapchat(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.935 6.073a9.646 9.646 0 0 1 .24 3.9v.047c-.011.146-.018.278-.024.41a.62.62 0 0 0 .322.072c.29-.026.572-.108.831-.241a.812.812 0 0 1 .373-.084a1.1 1.1 0 0 1 .409.072c.321.074.56.346.59.674c.013.361-.312.674-.975.939c-.071.023-.169.061-.276.1c-.363.109-.917.289-1.073.651a.836.836 0 0 0 .1.7l.012.013a5.678 5.678 0 0 0 3.854 3.229a.4.4 0 0 1 .337.409a.515.515 0 0 1-.036.181c-.193.457-1.023.8-2.53 1.021a2.23 2.23 0 0 0-.132.459c-.025.15-.06.299-.107.444a.42.42 0 0 1-.446.325h-.024a2.45 2.45 0 0 1-.433-.059a4.802 4.802 0 0 0-1.756-.048a3.613 3.613 0 0 0-1.386.71a4.392 4.392 0 0 1-2.651 1.027c-.048 0-.1-.012-.146-.012h-.121a4.28 4.28 0 0 1-2.633-1.036a3.565 3.565 0 0 0-1.373-.711a5.514 5.514 0 0 0-.745-.06a5.06 5.06 0 0 0-1.023.12a2.27 2.27 0 0 1-.435.059a.451.451 0 0 1-.47-.337c-.048-.154-.071-.313-.107-.456a2.268 2.268 0 0 0-.133-.458c-1.543-.179-2.373-.517-2.565-.986a.4.4 0 0 1 .288-.59a5.675 5.675 0 0 0 3.854-3.232l.009-.025a.8.8 0 0 0 .1-.7c-.156-.349-.711-.529-1.07-.65a1.759 1.759 0 0 1-.279-.1c-.89-.35-1.011-.748-.962-1.024a.95.95 0 0 1 .94-.626c.106 0 .21.02.308.06c.276.142.578.224.888.241c.13.004.26-.027.375-.089l-.041-.458a9.608 9.608 0 0 1 .241-3.889a5.057 5.057 0 0 1 4.728-3.054L12.117 3h.049a5.142 5.142 0 0 1 4.769 3.073Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAscending(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 17h6m-6-5h9m5-1v8m0 0l3-3m-3 3l-3-3M4 7h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDescending(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 17h12M4 12h9M4 7h6m8 6V5m0 0l3 3m-3-3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spectrum(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 22.001l-1.2-.006c-3.725-.04-5.982-.3-7.238-1.556c-1.256-1.256-1.515-3.515-1.556-7.238L2 11.691l.006-.89c-.028-1.568.09-3.135.35-4.681a4.715 4.715 0 0 1 1.3-2.648c1.317-1.218 3.637-1.451 7.732-1.471l1.812.005c3.961.043 6.2.325 7.415 1.748C21.793 5.128 22 7.508 22 12.001l-.005 1.2c-.043 3.961-.325 6.2-1.749 7.415c-1.374 1.177-3.754 1.385-8.246 1.385Zm-4.333-15a.669.669 0 0 0-.661.569L7 7.667v3.666a.67.67 0 0 0 .568.66l.1.007h.582A3.752 3.752 0 0 1 12 15.55v.782c.002.33.242.61.568.661l.1.007h3.667a.67.67 0 0 0 .66-.569l.007-.1v-.582a8.749 8.749 0 0 0-8.494-8.748h-.841Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotify(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.01 22A9.978 9.978 0 0 1 4.954 4.954a9.978 9.978 0 1 1 14.111 14.111A9.919 9.919 0 0 1 12.01 22Zm-1.482-6.728a9.794 9.794 0 0 1 5.227 1.384a.59.59 0 0 0 .33.1a.632.632 0 0 0 .521-.3a.7.7 0 0 0 .042-.545a.519.519 0 0 0-.293-.3a11.039 11.039 0 0 0-5.88-1.549c-1.283.01-2.56.16-3.81.449a.628.628 0 0 0-.45.749a.656.656 0 0 0 .591.472a.56.56 0 0 0 .157-.023a16.12 16.12 0 0 1 3.565-.437Zm-.19-2.976a12.648 12.648 0 0 1 6.416 1.661c.111.08.245.124.383.124c.27-.01.518-.15.668-.373a.71.71 0 0 0 .1-.586a.622.622 0 0 0-.3-.412a14.262 14.262 0 0 0-7.268-1.911a13.768 13.768 0 0 0-3.972.56a.773.773 0 0 0-.5.952a.738.738 0 0 0 .741.523a.85.85 0 0 0 .213-.025a11.66 11.66 0 0 1 3.52-.513Zm.19-3.088a14.987 14.987 0 0 1 7.376 1.7c.153.084.326.127.5.125a1 1 0 0 0 .8-.373a.985.985 0 0 0 .088-.724a.91.91 0 0 0-.436-.574a16.838 16.838 0 0 0-8.337-1.964a17.36 17.36 0 0 0-4.758.614a.915.915 0 0 0 .287 1.789a.916.916 0 0 0 .262-.038a14.7 14.7 0 0 1 4.217-.556v.001Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6.2v11.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104V6.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C19.48 3 18.92 3 17.8 3H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 4.52 3 5.08 3 6.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 9l-5.333 6L8 12m-5 5.8V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 3 5.08 3 6.2 3h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v11.607c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218H6.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C3 19.48 3 18.92 3 17.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareHelp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.146 9.074a2.998 2.998 0 0 1 5.28-.838A3 3 0 0 1 12 13v1m.05 3v.1h-.1V17zM3 17.8V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 3 5.08 3 6.2 3h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v11.607c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218H6.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C3 19.48 3 18.92 3 17.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareWarning(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8.45v4m.05 3v.1h-.1v-.1zM3 17.8V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 3 5.08 3 6.2 3h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v11.607c0 1.118 0 1.677-.218 2.104a2.002 2.002 0 0 1-.875.874c-.427.218-.986.218-2.104.218H6.197c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C3 19.48 3 18.92 3 17.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackOverflow(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.84 20.999H4.381v-6.4h1.6v4.8h11.26v-4.8h1.6v6.4Zm-3.2-3.2H7.581v-1.6h8.055v1.6h.004Zm0-2L7.781 14.16l.338-1.552l7.861 1.642l-.343 1.549h.003Zm.44-2.037l-7.28-3.4v-.006l.673-1.457l7.281 3.4l-.673 1.464l-.002-.001Zm.922-1.845l-6.17-5.14h.01l1.012-1.214l6.162 5.136L17 11.918l.003-.001Zm1.308-1.5l-4.807-6.449l1.31-.969L19.62 9.45l-1.312.971l.003-.004Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.335 10.337a.5.5 0 0 1 .28-.864l6.004-.712a.5.5 0 0 0 .396-.287l2.532-5.49a.5.5 0 0 1 .908 0l2.532 5.49a.499.499 0 0 0 .395.287l6.004.712a.5.5 0 0 1 .28.864l-4.438 4.105a.5.5 0 0 0-.15.464l1.177 5.93a.5.5 0 0 1-.735.534l-5.275-2.953a.499.499 0 0 0-.488 0l-5.276 2.952a.5.5 0 0 1-.735-.533l1.178-5.93a.5.5 0 0 0-.15-.464z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8.2v7.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h7.607c1.118 0 1.676 0 2.104-.218c.376-.192.682-.498.874-.875c.218-.427.218-.986.218-2.104V8.197c0-1.118 0-1.678-.218-2.105a2 2 0 0 0-.874-.874C17.48 5 16.92 5 15.8 5H8.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C5 6.52 5 7.08 5 8.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0"/><path d="M15 13.4v-2.8c0-.56 0-.84-.11-1.054a.998.998 0 0 0-.436-.437C14.24 9 13.96 9 13.4 9h-2.8c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C9 9.76 9 10.04 9 10.6v2.8c0 .56 0 .84.109 1.054c.096.188.249.34.437.437C9.76 15 10.04 15 10.6 15h2.8c.56 0 .84 0 1.054-.11a.997.997 0 0 0 .437-.436C15 14.24 15 13.96 15 13.4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSign(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.75 5.75l12.5 12.5M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8Zm0-14a6 6 0 1 0 6 6a6.007 6.007 0 0 0-6-6Zm1 7h-2V9h2v5Zm6.293-6.293l-2-2l1.414-1.414l2 2l-1.413 1.413l-.001.001ZM15 4H9V2h6v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12c.896 0 1.775.193 2.546.557c.348.165.669.362.955.587c.347.272.645.585.881.93c.432.627.644 1.336.616 2.052c-.028.716-.296 1.412-.776 2.017c-.48.605-1.154 1.096-1.952 1.42a6.075 6.075 0 0 1-2.583.43a5.865 5.865 0 0 1-2.497-.685c-.74-.402-1.332-.957-1.713-1.605M12 12H4m8 0h8m-3.476-5.703c-.381-.648-.973-1.202-1.714-1.605a5.866 5.866 0 0 0-2.496-.684a6.075 6.075 0 0 0-2.584.428c-.798.325-1.472.816-1.952 1.42c-.48.606-.747 1.303-.776 2.019c-.008.21.005.418.037.625"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5v8H7.83l2.58-2.59L9 9l-5 5l5 5l1.41-1.41L7.83 15H20V5h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 5v8h10.17l-2.58-2.59L15 9l5 5l-5 5l-1.41-1.41L16.17 15H4V5h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 9.52 3 10.08 3 11.2v5.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104v-5.607c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C19.48 8 18.92 8 17.8 8H16M8 8h8M8 8a4 4 0 1 1 8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4V2m0 18v2M6.414 6.414L5 5m12.728 12.728l1.414 1.414M4 12H2m18 0h2m-4.271-5.586L19.143 5M6.415 17.728L5 19.142M12 17a5 5 0 1 1 0-10a5 5 0 0 1 0 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwatchesPalette(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.5 21h13.97a.53.53 0 0 0 .53-.53v-5.94M7.98 20.67l12.662-5.904a.53.53 0 0 0 .256-.703l-2.51-5.385a.53.53 0 0 0-.704-.256l-5.654 2.636m-2.148 7.346a3.5 3.5 0 0 1-6.762-1.812L6.736 3.1a.53.53 0 0 1 .648-.375l5.74 1.538a.53.53 0 0 1 .373.648zM6.5 17.6h.002v.002H6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwichtLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 18h8a6 6 0 0 0 0-12H8a6 6 0 1 0 0 12"/><path d="M8 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwichtRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 18h8a6 6 0 0 0 0-12H8a6 6 0 1 0 0 12"/><path d="M16 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 15v1.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218H12m-8-5V9m0 6h8M4 9V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C5.52 4 6.08 4 7.2 4H12M4 9h8m0-5h4.8c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105V9m-8-5v5m0 0v6m0-6h8m-8 6v5m0-5h8m-8 5h4.804c1.118 0 1.677 0 2.104-.218c.376-.192.682-.498.874-.874c.218-.428.218-.986.218-2.104V15m0 0V9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4h4.8c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105V9h-8m0-5H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 5.52 3 6.08 3 7.2V9m8-5v5M3 9v6m0-6h8m-8 6v1.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218H11V9m-8 6h8m4 1h3m0 0h3m-3 0v3m0-3v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableDelete(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 10h-6v11H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8h-2v-3Zm-8 0H5v4h6v-4Zm0 9v-3H5v3h6Zm2-14v3h6V5h-6Zm-2 0H5v3h6V5Z"/><path fill="currentColor" d="M16 18v-2h8v2h-8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4h4.8c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.105V9h-8m0-5H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 5.52 3 6.08 3 7.2V9m8-5v5M3 9v6m0-6h8m-8 6v1.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218H11V9m-8 6h8m4 1h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h9.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V6.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C18.48 3 17.92 3 16.8 3H7.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C4 4.52 4 5.08 4 6.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletButton(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6.2v11.6c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V6.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C19.48 3 18.92 3 17.8 3H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 4.52 3 5.08 3 6.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m4.748 7.645l-.331 3.644c-.05.54-.074.811-.03 1.07a2 2 0 0 0 .238.655c.131.228.325.422.711.808l5.176 5.176c.787.787 1.18 1.18 1.636 1.328c.402.131.835.131 1.237 0c.456-.148.853-.544 1.645-1.336l3.96-3.96c.792-.792 1.187-1.188 1.336-1.644a2 2 0 0 0-.001-1.236c-.148-.457-.543-.853-1.335-1.645l-5.163-5.163c-.39-.39-.584-.584-.813-.716a2 2 0 0 0-.656-.238c-.26-.045-.535-.02-1.084.03l-3.63.33c-.944.086-1.417.129-1.787.335a2 2 0 0 0-.775.775c-.205.368-.248.838-.333 1.773z"/><path d="M9.713 9.713a1 1 0 1 0-1.415-1.414a1 1 0 0 0 1.415 1.414"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.897 21.968a1.987 1.987 0 0 1-1.415-.586l-7.834-7.835a1.994 1.994 0 0 1-.58-1.567l.5-6.566a1.989 1.989 0 0 1 1.846-1.841l6.566-.5c.051-.011.103-.011.154-.011a2.01 2.01 0 0 1 1.413.586l7.835 7.834a2 2 0 0 1 0 2.829l-7.07 7.071a1.987 1.987 0 0 1-1.415.586Zm-.764-16.906l-6.57.5l-.5 6.571l7.834 7.835l7.07-7.07l-7.834-7.836Zm-3.479 5.592a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tennis(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.847 15.827c2.643 1.113 6.116-.372 7.98-3.505c1.974-3.317 1.397-7.226-1.29-8.731c-2.685-1.505-6.463-.036-8.437 3.282c-1.636 2.749-1.52 5.905.098 7.762L5 20.009l1.769.99l3.078-5.172Zm6.236-4.482a6.647 6.647 0 0 1-1.346 1.633V5.04c.288.051.55.142.778.27c1.375.77 2.205 3.284.568 6.035ZM11.674 14.2V5.84c.678-.47 1.386-.732 2.042-.817v8.657c-.7.37-1.41.532-2.042.52Zm-1.021-.2a2.203 2.203 0 0 1-.24-.116c-1.376-.77-2.205-3.284-.568-6.035c.248-.418.52-.785.808-1.104V14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TennisMatch(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.85 6.594c.107-.226.224-.45.353-.672l.027-.046l.015-.026c.227-.384.478-.742.748-1.075c1.05-1.29 2.395-2.18 3.77-2.562a8.06 8.06 0 0 1 .014-.005c1.304-.36 2.634-.265 3.77.371l.03.018c2.646 1.51 3.212 5.417 1.265 8.733c-1.721 2.93-4.83 4.419-7.361 3.694L17 21.01l-1.744.99L12 16.462L8.748 22L7 21.008l4.326-7.364c-.894-1.04-1.324-2.486-1.246-4.04a8.05 8.05 0 0 1 .77-3.01Zm6.938-2.545v7.916a6.662 6.662 0 0 0 1.314-1.622c1.623-2.763.79-5.268-.548-6.027a2.388 2.388 0 0 0-.766-.267Zm-1.009 8.625V4.036c-.647.087-1.347.35-2.018.825v8.34c.623.01 1.324-.153 2.018-.527Zm-4.018-.477c.15.205.318.38.5.52a2.135 2.135 0 0 0 .479.282l.012.005v-7.23a7.03 7.03 0 0 0-.783 1.089l-.015.026l-.023.039c-.06.104-.117.209-.17.313c-1.044 2.038-.775 3.895 0 4.956Z"/><path fill="currentColor" d="M9.422 5.006c-.728-.567-1.502-.876-2.211-.97v8.607c.5.269 1.003.429 1.477.493a7.1 7.1 0 0 0 .277.579l-.799 1.36c-1.9-.362-3.799-1.709-5.01-3.769C1.21 7.995 1.78 4.092 4.428 2.59c1.842-1.045 4.205-.654 6.128.795a10.51 10.51 0 0 0-1.134 1.62ZM4.896 10.32a6.636 6.636 0 0 0 1.308 1.614V4.048c-.283.05-.54.14-.762.266c-1.334.756-2.165 3.251-.546 6.006Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TennisMatchAlt(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.828 11.322c-1.865 3.133-5.338 4.618-7.981 3.505L3.769 20L2 19.01l3.198-5.375C3.58 11.778 3.464 8.622 5.1 5.873c1.974-3.318 5.752-4.787 8.438-3.282a4.551 4.551 0 0 1 1.742 1.748a5.55 5.55 0 0 1 .64 1.87c.254 1.596-.091 3.43-1.092 5.113Zm-.841-4.003a4.64 4.64 0 0 0-.01-.31c-.083-1.274-.686-2.265-1.462-2.7a2.466 2.466 0 0 0-.778-.269v7.938c.07-.06.137-.122.205-.186c.41-.392.798-.872 1.14-1.447c.64-1.075.903-2.113.904-3.026Zm-3.98 5.667c.222-.075.445-.168.667-.283l.042-.022V4.024c-.655.085-1.364.346-2.042.816v8.36a3.97 3.97 0 0 0 1.334-.214Zm-2.594-.1c.076.042.156.08.24.115V5.746a6.93 6.93 0 0 0-.808 1.104c-1.637 2.751-.808 5.265.568 6.035Z"/><path fill="currentColor" d="M16.716 14.681c-.91.48-1.837.61-2.586.462c-1.122.912-2.412 1.57-3.754 1.873L8 21.01l1.769.99l3.078-5.172c2.643 1.113 6.116-.372 7.98-3.505c1.974-3.317 1.397-7.226-1.29-8.731a4.693 4.693 0 0 0-2.175-.589c.283.654.467 1.354.56 2.077c.215.053.415.13.593.23c1.375.77 2.205 3.285.568 6.035a6.646 6.646 0 0 1-1.346 1.634v-4.6a10.402 10.402 0 0 1-1.02 2.67v2.633Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 15h-5m-5-5l3 2.5L7 15m-4 .8V8.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v7.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H6.197c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C3 17.48 3 16.92 3 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignCenter(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 18H7m13-4H4m13-4H7m13-4H4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignJustify(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 18H4m16-4H4m16-4H4m16-4H4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 18h10M4 14h16M4 10h10M4 6h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 18H10m10-4H4m16-4H10m10-4H4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextIcon(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19h2m0 0h2m-2 0V5m0 0H6v1m6-1h6v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinBigDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 17.5L13 24l6.5-6.5l-.707-.707l-5.293 5.293V0h-1v22.086l-5.293-5.293l-.707.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinBigLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 5.5L0 12l6.5 6.5l.707-.707L1.914 12.5H24v-1H1.914l5.293-5.293L6.5 5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinBigRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 18.5L24 12l-6.5-6.5l-.707.707l5.293 5.293H0v1h22.086l-5.293 5.293l.707.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinBigUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#ciThinBigUp0)"><path fill="currentColor" d="M19.5 6.5L13 0L6.5 6.5l.707.707L12.5 1.914V24h1V1.914l5.293 5.293l.707-.707Z"/></g><defs><clipPath id="ciThinBigUp0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinLongDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 18.5L12 22l3.5-3.5l-.707-.707l-2.293 2.293V2h-1v18.086l-2.293-2.293l-.707.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinLongLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 8.5L2 12l3.5 3.5l.707-.707L3.914 12.5H22v-1H3.914l2.293-2.293L5.5 8.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinLongRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 15.5L22 12l-3.5-3.5l-.707.707l2.293 2.293H2v1h18.086l-2.293 2.293l.707.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinLongTwoDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 16l3 3l3-3l-.707-.707l-1.793 1.793V5h-1v12.086l-1.793-1.793L9 16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinLongTwoLeft(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 9l-3 3l3 3l.707-.707L6.914 12.5H19v-1H6.914l1.793-1.793L8 9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinLongTwoRight(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 15l3-3l-3-3l-.707.707l1.793 1.793H5v1h12.086l-1.793 1.793L16 15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinLongTwoUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15 8l-3-3l-3 3l.707.707L11.5 6.914V19h1V6.914l1.793 1.793L15 8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinLongUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 5.5L12 2L8.5 5.5l.707.707L11.5 3.914V22h1V3.914l2.293 2.293l.707-.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TicketVoucher(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 6H6c-.932 0-1.398 0-1.766.152a2 2 0 0 0-1.082 1.083C3 7.602 3 8.068 3 9a3 3 0 1 1 0 6c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.082 1.083C4.602 18 5.068 18 6 18h8m0-12h4c.932 0 1.398 0 1.765.152a2 2 0 0 1 1.083 1.083C21 7.602 21 8.068 21 9a3 3 0 1 0 0 6c0 .932 0 1.398-.152 1.765a2 2 0 0 1-1.083 1.083C19.398 18 18.932 18 18 18h-4m0-12v12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13V9m9-3l-2-2m-9-2h4m-2 19a8 8 0 1 1 0-16a8 8 0 0 1 0 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimerAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v-3m0 0v-3m0 3H9m3 0h3m6-7l-2-2m-9-2h4m-2 19a8 8 0 1 1 0-16a8 8 0 0 1 0 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimerClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 15l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-9l-2-2m-9-2h4m-2 19a8 8 0 1 1 0-16a8 8 0 0 1 0 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimerRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 13h6m6-7l-2-2m-9-2h4m-2 19a8 8 0 1 1 0-16a8 8 0 0 1 0 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Transfer(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 20l-5-4l5-4v3h13v2H9v3Zm6-8V9H2V7h13V4l5 4l-5 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashEmpty(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 6v11.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h5.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.987.218-2.105V6M6 6h2M6 6H4m4 0h8M8 6c0-.932 0-1.398.152-1.765a2 2 0 0 1 1.082-1.083C9.602 3 10.068 3 11 3h2c.932 0 1.398 0 1.765.152a2 2 0 0 1 1.083 1.083C16 4.602 16 5.068 16 6m0 0h2m0 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashFull(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10v7m-4-7v7M6 6v11.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h5.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.987.218-2.105V6M6 6h2M6 6H4m4 0h8M8 6c0-.932 0-1.398.152-1.765a2 2 0 0 1 1.082-1.083C9.602 3 10.068 3 11 3h2c.932 0 1.398 0 1.765.152a2 2 0 0 1 1.083 1.083C16 4.602 16 5.068 16 6m0 0h2m0 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trello(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.75 21H5.25A2.253 2.253 0 0 1 3 18.75V5.25A2.253 2.253 0 0 1 5.25 3h13.5A2.253 2.253 0 0 1 21 5.25v13.5A2.253 2.253 0 0 1 18.75 21ZM6.42 5.34a1.08 1.08 0 0 0-1.08 1.08v10.215c0 .596.484 1.08 1.08 1.08h3.33a1.08 1.08 0 0 0 1.08-1.08V6.42a1.08 1.08 0 0 0-1.08-1.08H6.42Zm7.83 0a1.08 1.08 0 0 0-1.08 1.08v5.715c0 .596.484 1.08 1.08 1.08h3.33a1.08 1.08 0 0 0 1.08-1.08V6.42a1.08 1.08 0 0 0-1.08-1.08h-3.33Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingDown(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 17l-5.846-5.937a4.945 4.945 0 0 0-.205-.203a1.999 1.999 0 0 0-2.667 0c-.047.042-.1.096-.205.203c-.105.106-.158.16-.205.202a2 2 0 0 1-2.667 0A4.898 4.898 0 0 1 8 11.062L4 7m16 10v-6m0 6h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingUp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 7l-5.846 5.938c-.105.106-.158.16-.205.202c-.76.68-1.907.68-2.667 0a4.814 4.814 0 0 1-.205-.203c-.105-.106-.158-.16-.205-.202a2 2 0 0 0-2.667 0a4.86 4.86 0 0 0-.204.202L4 17M20 7v6m0-6h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.379 15.2c-.91 1.575-1.364 2.363-1.296 3.01a2 2 0 0 0 .813 1.408C4.422 20 5.331 20 7.15 20h9.703c1.817 0 2.726 0 3.251-.382a2 2 0 0 0 .814-1.408c.068-.647-.386-1.435-1.296-3.01l-4.85-8.4c-.909-1.575-1.364-2.363-1.958-2.627a2 2 0 0 0-1.627 0c-.593.264-1.048 1.052-1.957 2.625z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 13l-4 4l-2-2m-4.621.2c-.91 1.575-1.364 2.363-1.296 3.01a2 2 0 0 0 .813 1.408C4.422 20 5.331 20 7.15 20h9.703c1.817 0 2.726 0 3.251-.382a2 2 0 0 0 .814-1.409c.068-.646-.386-1.434-1.296-3.01l-4.85-8.4c-.909-1.574-1.364-2.362-1.958-2.626a2 2 0 0 0-1.627 0c-.593.264-1.048 1.052-1.957 2.625z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleWarning(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v4m-7.621 2.2c-.91 1.575-1.364 2.363-1.296 3.01a2 2 0 0 0 .813 1.408C4.422 20 5.331 20 7.15 20h9.703c1.817 0 2.726 0 3.251-.382a2 2 0 0 0 .814-1.409c.068-.646-.386-1.434-1.296-3.01l-4.85-8.4c-.909-1.574-1.364-2.362-1.958-2.626a2 2 0 0 0-1.627 0c-.593.264-1.048 1.052-1.957 2.625zm7.672.8v.1h-.1V16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.995 6.688a4.3 4.3 0 0 0 1.887-2.374a8.59 8.59 0 0 1-2.725 1.041a4.3 4.3 0 0 0-7.316 3.915a12.184 12.184 0 0 1-8.844-4.484a4.3 4.3 0 0 0 1.328 5.73a4.276 4.276 0 0 1-1.943-.538v.054a4.294 4.294 0 0 0 3.443 4.208a4.3 4.3 0 0 1-1.938.074a4.3 4.3 0 0 0 4.009 2.98a8.61 8.61 0 0 1-5.33 1.837c-.343 0-.685-.02-1.025-.059A12.148 12.148 0 0 0 8.12 21A12.127 12.127 0 0 0 20.33 8.789c0-.186-.004-.371-.013-.555a8.718 8.718 0 0 0 2.142-2.222a8.58 8.58 0 0 1-2.464.676Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 19h12M8 5v6a4 4 0 0 0 8 0V5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 8H5V3m.291 13.357a8 8 0 1 0 .188-8.991"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnfoldLess(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 19l4-4l4 4m0-14l-4 4l-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnfoldMore(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 15l-4 4l-4-4m0-6l4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlink(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.292 21.707L11.586 13H7v-2h2.586l-2-2H7a3 3 0 1 0 0 6h3v2H7a5 5 0 0 1-1.255-9.841L2.292 3.707l1.415-1.414l18 18l-1.414 1.413v.001Zm-.156-5.814l-1.428-1.427A3 3 0 0 0 17 9h-3V7h3a5 5 0 0 1 3.137 8.893Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unsplash(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 21H3V10.875h5.625v5.063h6.75v-5.063H21V21ZM15.375 8.063h-6.75V3h6.75v5.063Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8a5 5 0 1 1 10 0A5 5 0 0 1 7 8Zm5 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm-5.657 5.343A8 8 0 0 0 4 22h2a6 6 0 1 1 12 0h2a8 8 0 0 0-13.657-5.657Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAdd(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19c0-2.21-2.686-4-6-4s-6 1.79-6 4m16-3v-3m0 0v-3m0 3h-3m3 0h3M9 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCardId(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18h6m-6 0c-.988 0-1.507-.013-1.908-.218a2.001 2.001 0 0 1-.874-.875C3 16.48 3 15.92 3 14.8V9.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 6 5.08 6 6.2 6h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v5.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H12m-6 0c0-1.105 1.343-2 3-2s3 .895 3 2m-6 0s0 0 0 0m12-4h-4m4-3h-3m-6 2a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19c0-2.21-2.686-4-6-4s-6 1.79-6 4m18-9l-4 4l-2-2m-6 0a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircle(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.217 19.332A6.982 6.982 0 0 0 12 17c-2.073 0-3.935.9-5.217 2.332M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18m0-7a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19c0-2.21-2.686-4-6-4s-6 1.79-6 4m14-5l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2M9 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserHeart(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 19h2a4 4 0 0 1 8 0h2a6 6 0 0 0-12 0ZM4 8a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm2.002-.029A2 2 0 0 0 10 8.09V8a2 2 0 0 0-3.998-.029ZM18.339 15a22.972 22.972 0 0 0-.692-.583l-.047-.038l-.006-.004C16.438 13.432 15 12.258 15 10.799A1.8 1.8 0 0 1 16.839 9a2.008 2.008 0 0 1 1.5.667a2.009 2.009 0 0 1 1.5-.667a1.8 1.8 0 0 1 1.835 1.8c0 1.465-1.45 2.647-2.615 3.598l-.003.002l-.057.047l-.018.015c-.23.189-.448.367-.643.54l.001-.003Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 19H2a6 6 0 0 1 12 0h-2a4 4 0 0 0-8 0Zm18-6h-8v-2h8v2ZM8 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-6a2 2 0 1 0 2 2.09v.4V8a2 2 0 0 0-2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserOne(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21a7 7 0 1 0-14 0m7-10a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPin(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 22l-3-3H5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2h-4l-3 3ZM5 4v13h5l2 2l2-2h5V4H5Zm11 10.986H8V14.6a3.82 3.82 0 0 1 4-3.6a3.82 3.82 0 0 1 4 3.6v.386ZM12 10a1.934 1.934 0 0 1-2-2a1.935 1.935 0 0 1 2-2a1.935 1.935 0 0 1 2 2a1.934 1.934 0 0 1-2 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 19H2a6 6 0 0 1 12 0h-2a4 4 0 0 0-8 0Zm15-3h-2v-3h-3v-2h3V8h2v3h3v2h-3v3ZM8 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-6a2 2 0 1 0 2 2.09v.4V8a2 2 0 0 0-2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserRemove(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19c0-2.21-2.686-4-6-4s-6 1.79-6 4m12-6h6M9 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSquare(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 21a5 5 0 0 0-10 0m10 0h.803c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104V6.197c0-1.118 0-1.678-.218-2.105a2.001 2.001 0 0 0-.875-.874C19.48 3 18.92 3 17.8 3H6.2c-1.12 0-1.68 0-2.108.218a1.999 1.999 0 0 0-.874.874C3 4.52 3 5.08 3 6.2v11.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218H7m10 0H7m5-8a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserThree(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 19c0-2.21-2.686-4-6-4s-6 1.79-6 4m6-7a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 21c0-2.761-3.582-5-8-5s-8 2.239-8 5m8-8a5 5 0 1 1 0-10a5 5 0 0 1 0 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserVoice(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19c0-2.21-2.686-4-6-4s-6 1.79-6 4M16.828 5.172a3.999 3.999 0 0 1 0 5.657M19 3a7.07 7.07 0 0 1 0 10M9 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 20c0-1.742-1.67-3.223-4-3.773M15 20c0-2.21-2.686-4-6-4s-6 1.79-6 4m12-7a4 4 0 0 0 0-8m-6 8a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsersGroup(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20c0-1.657-2.239-3-5-3s-5 1.343-5 3m14-3c0-1.23-1.234-2.287-3-2.75M3 17c0-1.23 1.234-2.287 3-2.75m12-4.014a3 3 0 1 0-4-4.472m-8 4.472a3 3 0 0 1 4-4.472M12 14a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMax(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.82 4.687a10 10 0 0 1 .05 14.579M16.093 7.612a6 6 0 0 1 .03 8.748m-8.642-.953l1.676 2.073c.873 1.08 1.31 1.621 1.692 1.68a1 1 0 0 0 .891-.315c.261-.286.261-.981.261-2.37v-8.95c0-1.389 0-2.083-.26-2.37a1 1 0 0 0-.892-.315c-.383.059-.82.6-1.692 1.68L7.48 8.593c-.176.219-.264.328-.373.406a1 1 0 0 1-.32.153c-.13.036-.27.036-.551.036H4.813c-.757 0-1.135 0-1.44.1A2 2 0 0 0 2.1 10.56c-.1.306-.1.684-.1 1.44c0 .756 0 1.134.1 1.44a2 2 0 0 0 1.273 1.272c.305.1.683.1 1.44.1h1.423c.28 0 .42 0 .55.036a1 1 0 0 1 .32.153c.11.078.198.188.374.406"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMin(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.41 8.343a5 5 0 0 1 .025 7.29m-8.955-.226l1.676 2.073c.873 1.08 1.31 1.62 1.692 1.68a1 1 0 0 0 .891-.316c.261-.286.261-.98.261-2.37V7.525c0-1.39 0-2.084-.26-2.37a1 1 0 0 0-.892-.315c-.383.059-.82.599-1.692 1.68L7.48 8.593c-.176.218-.264.327-.373.406a1 1 0 0 1-.32.153c-.13.035-.27.035-.551.035H4.813c-.757 0-1.135 0-1.44.101A2 2 0 0 0 2.1 10.56c-.1.305-.1.683-.1 1.44c0 .756 0 1.134.1 1.44a2 2 0 0 0 1.273 1.271c.305.101.683.101 1.44.101h1.423c.28 0 .42 0 .55.036a1 1 0 0 1 .32.152c.11.079.198.188.374.406"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMinus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12h6M7.48 15.406l1.676 2.074c.873 1.08 1.31 1.62 1.692 1.68a1 1 0 0 0 .891-.316c.261-.286.261-.98.261-2.37V7.525c0-1.39 0-2.084-.26-2.37a1 1 0 0 0-.892-.315c-.383.059-.82.599-1.692 1.68L7.48 8.593c-.176.218-.264.327-.373.406a1 1 0 0 1-.32.153c-.13.035-.27.035-.551.035H4.813c-.757 0-1.135 0-1.44.101A2 2 0 0 0 2.1 10.56c-.1.305-.1.683-.1 1.44c0 .756 0 1.134.1 1.44a2 2 0 0 0 1.273 1.271c.305.101.683.101 1.44.101h1.423c.28 0 .42 0 .55.036a1 1 0 0 1 .32.152c.11.079.198.188.374.406"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16.172 9.171l5.656 5.657m-5.656 0l5.656-5.657M7.48 15.407l1.676 2.073c.873 1.08 1.31 1.62 1.692 1.68a1 1 0 0 0 .891-.316c.261-.286.261-.98.261-2.37V7.525c0-1.39 0-2.084-.26-2.37a1 1 0 0 0-.892-.315c-.383.059-.82.599-1.692 1.68L7.48 8.593c-.176.218-.264.327-.373.406a1 1 0 0 1-.32.153c-.13.035-.27.035-.551.035H4.813c-.757 0-1.135 0-1.44.101A2 2 0 0 0 2.1 10.56c-.1.305-.1.683-.1 1.44c0 .756 0 1.134.1 1.44a2 2 0 0 0 1.273 1.271c.305.101.683.101 1.44.101h1.423c.28 0 .42 0 .55.036a1 1 0 0 1 .32.152c.11.079.198.188.374.406"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOffTwo(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9.187H8.813c-.757 0-1.135 0-1.44.101A2 2 0 0 0 6.1 10.56c-.1.305-.1.683-.1 1.44c0 .756 0 1.134.1 1.44a2 2 0 0 0 1.273 1.271c.305.101.683.101 1.44.101h1.423c.28 0 .42 0 .55.036a.998.998 0 0 1 .32.153c.11.078.198.187.374.406l1.675 2.073c.874 1.08 1.31 1.62 1.693 1.68a1 1 0 0 0 .891-.316c.261-.286.261-.98.261-2.37V15m0-4.5V6.977c0-.936 0-1.404-.122-1.628a1 1 0 0 0-1.26-.445c-.235.097-.53.461-1.118 1.19l-.625.773M6 5l14 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumePlus(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12h6m-3 3V9M7.48 15.406l1.676 2.074c.873 1.08 1.31 1.62 1.692 1.68a1 1 0 0 0 .891-.316c.261-.286.261-.98.261-2.37V7.525c0-1.39 0-2.084-.26-2.37a1 1 0 0 0-.892-.315c-.383.059-.82.599-1.692 1.68L7.48 8.593c-.176.218-.264.327-.373.406a1 1 0 0 1-.32.153c-.13.035-.27.035-.551.035H4.813c-.757 0-1.135 0-1.44.101A2 2 0 0 0 2.1 10.56c-.1.305-.1.683-.1 1.44c0 .756 0 1.134.1 1.44a2 2 0 0 0 1.273 1.271c.305.101.683.101 1.44.101h1.423c.28 0 .42 0 .55.036a1 1 0 0 1 .32.152c.11.079.198.188.374.406"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v8m.05 4v.1h-.1V18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningOutline(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.266 20.998H2.733a1 1 0 0 1-.866-1.5l9.266-16a1 1 0 0 1 1.73 0l9.267 16a1 1 0 0 1-.865 1.5ZM12 5.998l-7.531 13h15.064L12 5.998Zm.995 9.001h-2V9.998h2v5.001Z"/><path fill="currentColor" d="M11 16h2v2h-2v-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterDrop(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 13.385a4.676 4.676 0 0 1-1.318 3.263A4.473 4.473 0 0 1 13 17.736m6-4.044C19 7.115 12 2 12 2S5 7.115 5 13.692c0 1.938.737 3.797 2.05 5.168C8.363 20.23 10.144 21 12 21c1.857 0 3.637-.77 4.95-2.14c1.313-1.371 2.05-3.23 2.05-5.168"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wavy(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18.92 17.158l.128-1.599a1.92 1.92 0 0 1 .453-1.092l1.04-1.222a1.92 1.92 0 0 0 0-2.491L19.5 9.533a1.921 1.921 0 0 1-.454-1.093l-.127-1.6a1.921 1.921 0 0 0-1.762-1.761l-1.599-.128a1.918 1.918 0 0 1-1.092-.452l-1.221-1.04a1.921 1.921 0 0 0-2.492 0l-1.221 1.04c-.308.262-.69.42-1.093.453l-1.6.128m12.08 12.079a1.92 1.92 0 0 1-1.76 1.762m-.001-.001l-1.6.127a1.92 1.92 0 0 0-1.092.453l-1.221 1.04a1.921 1.921 0 0 1-2.492 0l-1.22-1.04a1.92 1.92 0 0 0-1.094-.452l-1.6-.128m.002 0a1.92 1.92 0 0 1-1.761-1.76l-.128-1.6a1.92 1.92 0 0 0-.453-1.092l-1.04-1.221a1.922 1.922 0 0 1 0-2.492l1.04-1.221c.263-.308.42-.69.452-1.093l.128-1.6m.001.002A1.92 1.92 0 0 1 6.842 5.08"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WavyCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 10l-4 4l-2-2m4.246-8.541l1.221 1.04c.308.262.69.42 1.092.453l1.6.128a1.92 1.92 0 0 1 1.761 1.76l.127 1.6c.033.403.192.786.454 1.093l1.04 1.22a1.92 1.92 0 0 1 0 2.492l-1.04 1.221c-.262.308-.421.69-.453 1.093l-.128 1.6a1.92 1.92 0 0 1-1.76 1.761l-1.6.128a1.92 1.92 0 0 0-1.093.452l-1.221 1.04a1.921 1.921 0 0 1-2.492 0l-1.22-1.04a1.92 1.92 0 0 0-1.094-.452l-1.6-.128a1.92 1.92 0 0 1-1.76-1.762l-.128-1.599a1.92 1.92 0 0 0-.453-1.092l-1.04-1.222a1.92 1.92 0 0 1 0-2.49l1.04-1.222c.263-.308.42-.69.452-1.093l.128-1.599A1.922 1.922 0 0 1 6.842 5.08l1.598-.127A1.92 1.92 0 0 0 9.533 4.5l1.221-1.04a1.921 1.921 0 0 1 2.492 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WavyHelp(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.146 9.073a2.999 2.999 0 0 1 5.844.691A3 3 0 0 1 12 13v1m.05 3v.1h-.1V17zm1.196-13.541l1.221 1.04c.308.262.69.42 1.092.453l1.6.128a1.92 1.92 0 0 1 1.761 1.76l.127 1.6c.033.403.192.786.454 1.093l1.04 1.22a1.92 1.92 0 0 1 0 2.492l-1.04 1.221c-.262.308-.421.69-.453 1.093l-.128 1.6a1.92 1.92 0 0 1-1.76 1.761l-1.6.128a1.92 1.92 0 0 0-1.093.452l-1.221 1.04a1.921 1.921 0 0 1-2.492 0l-1.22-1.04a1.92 1.92 0 0 0-1.094-.452l-1.6-.128a1.92 1.92 0 0 1-1.76-1.762l-.128-1.599a1.92 1.92 0 0 0-.453-1.092l-1.04-1.222a1.92 1.92 0 0 1 0-2.49l1.04-1.222c.263-.308.42-.69.452-1.093l.128-1.599A1.922 1.922 0 0 1 6.842 5.08l1.598-.127A1.92 1.92 0 0 0 9.533 4.5l1.221-1.04a1.921 1.921 0 0 1 2.492 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WavyWarning(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8.45v4m1.246-8.991l1.221 1.04c.308.262.69.42 1.092.453l1.6.128a1.92 1.92 0 0 1 1.761 1.76l.127 1.6c.033.403.192.786.454 1.093l1.04 1.22a1.92 1.92 0 0 1 0 2.492l-1.04 1.221c-.262.308-.421.69-.453 1.093l-.128 1.6a1.92 1.92 0 0 1-1.76 1.761l-1.6.128a1.92 1.92 0 0 0-1.093.452l-1.221 1.04a1.921 1.921 0 0 1-2.492 0l-1.22-1.04a1.92 1.92 0 0 0-1.094-.452l-1.6-.128a1.92 1.92 0 0 1-1.76-1.762l-.128-1.599a1.92 1.92 0 0 0-.453-1.092l-1.04-1.222a1.92 1.92 0 0 1 0-2.49l1.04-1.222c.263-.308.42-.69.452-1.093l.128-1.599A1.922 1.922 0 0 1 6.842 5.08l1.598-.127A1.92 1.92 0 0 0 9.533 4.5l1.221-1.04a1.921 1.921 0 0 1 2.492 0M12.05 15.45v.1h-.1v-.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiHigh(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.343 14.59a5 5 0 0 1 7.29-.025m-9.484-3.021a8 8 0 0 1 11.664-.04M3.223 8.816a12 12 0 0 1 17.497-.06M12 19a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiLow(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.343 14.59a5.001 5.001 0 0 1 7.29-.025M12 19a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiMedium(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.344 14.59a5.001 5.001 0 0 1 7.289-.025m-9.484-3.021a8 8 0 0 1 11.664-.042M12 19a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiNone(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 18a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiOff(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.85 11.544a8 8 0 0 0-2.88-1.972m5.806-.756a12 12 0 0 0-9.488-3.795m-2.945 9.57a5 5 0 0 1 4.902-1.434m-7.096-1.613A8 8 0 0 1 9.623 9.36m-6.4-.545a12 12 0 0 1 3.11-2.393M4.413 4l14.142 14.142M12 19a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiProblem(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.343 14.59a4.998 4.998 0 0 1 7.29-.025m-9.484-3.021A8 8 0 0 1 11.96 9m-8.735-.184A12 12 0 0 1 11.959 5M16 9l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m-8 10a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Window(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h.02m0 0h17.96M3.02 6C3 6.314 3 6.702 3 7.2v9.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104V7.197c0-.497 0-.883-.02-1.197M3.02 6c.023-.392.077-.67.198-.908c.192-.377.497-.682.874-.874C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.121.237.175.516.199.908m0 0H21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowCheck(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h.02m0 0h17.96M3.02 6C3 6.314 3 6.702 3 7.2v9.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104V7.197c0-.497 0-.883-.02-1.197M3.02 6c.023-.392.077-.67.198-.908c.192-.377.497-.682.874-.874C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.121.237.175.516.199.908m0 0H21m-6 5l-4 4l-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowClose(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h.02m0 0h17.96M3.02 6C3 6.314 3 6.702 3 7.2v9.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104V7.197c0-.497 0-.883-.02-1.197M3.02 6c.023-.392.077-.67.198-.908c.192-.377.497-.682.874-.874C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.121.237.175.516.199.908m0 0H21m-11 9l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowCodeBlock(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h.02m0 0h17.96M3.02 6C3 6.314 3 6.702 3 7.2v9.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218a2 2 0 0 0 .875-.875c.218-.427.218-.986.218-2.104V7.197c0-.497 0-.883-.02-1.197M3.02 6c.023-.392.077-.67.198-.908c.192-.377.497-.682.874-.874C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.121.237.175.516.199.908m0 0H21m-7 5l2 2l-2 2m-4 0l-2-2l2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowSidebar(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h.02m0 0H9M3.02 6C3 6.314 3 6.702 3 7.2v9.6c0 1.12 0 1.68.218 2.107c.192.377.497.683.874.875c.427.218.987.218 2.105.218H9M3.02 6c.023-.392.077-.67.198-.908c.192-.377.497-.682.874-.874C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.121.237.175.516.199.908M9 6h11.98M9 6v14M20.98 6H21m-.02 0c.02.314.02.7.02 1.197v9.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowTerminal(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16m-4 10h-3m-5-4l2 2l-2 2m-5 .8V7.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v9.606c0 1.118 0 1.677-.218 2.104a2.003 2.003 0 0 1-.875.875c-.427.218-.986.218-2.104.218H6.197c-1.118 0-1.678 0-2.105-.218a2.001 2.001 0 0 1-.874-.875C3 18.48 3 17.92 3 16.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *CiIcon {
	return &CiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.006 19.012h-.02c-.062 0-6.265-.012-7.83-.437a2.5 2.5 0 0 1-1.764-1.765A26.494 26.494 0 0 1 1.986 12a26.646 26.646 0 0 1 .417-4.817A2.564 2.564 0 0 1 4.169 5.4c1.522-.4 7.554-.4 7.81-.4H12c.063 0 6.282.012 7.831.437c.859.233 1.53.904 1.762 1.763c.29 1.594.427 3.211.407 4.831a26.568 26.568 0 0 1-.418 4.811a2.51 2.51 0 0 1-1.767 1.763c-1.52.403-7.553.407-7.809.407Zm-2-10.007l-.005 6l5.212-3l-5.207-3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
