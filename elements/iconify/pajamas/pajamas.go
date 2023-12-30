package pajamas

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "3.74.0"
	hAttr          = "1em"
	viewbox        = "0 0 0 0"
	fill           = "currentColor"
)

type PajamasIcon struct {
	*SVGSVGElement
}

func Abuse(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.308 1.5L1.5 5.308v5.384L5.308 14.5h5.384l3.808-3.808V5.308L10.692 1.5zM11.314 0H4.686L0 4.686v6.628L4.686 16h6.628L16 11.314V4.686zM9 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25-6.25a.75.75 0 0 0-1.5 0v3.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Accessibility(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.434 4.187a.874.874 0 0 1-.67 1.04l-2.718.587a1 1 0 0 0-.79.977v.88a8 8 0 0 0 .507 2.8l1.624 4.349a.874.874 0 1 1-1.638.611l-1.554-4.16a.5.5 0 0 0-.937 0l-1.554 4.16a.874.874 0 0 1-1.638-.611l1.625-4.35a8 8 0 0 0 .506-2.8v-.879a1 1 0 0 0-.789-.977L2.69 5.226a.874.874 0 0 1 .37-1.709l3.822.826a4 4 0 0 0 1.69 0l3.822-.826a.874.874 0 0 1 1.04.67M7.684 0a1.749 1.749 0 1 1 0 3.497a1.749 1.749 0 0 1 0-3.497"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Account(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M7 5a2.99 2.99 0 0 1-.87 2.113A3.997 3.997 0 0 1 8 10.5V12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-1.5c0-1.427.747-2.679 1.87-3.387A3 3 0 1 1 7 5m-5.5 5.5a2.5 2.5 0 0 1 5 0V12a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5zm12.251-7.165c.21.122.462.163.685.07l.313-.133a.5.5 0 0 1 .628.21l.308.535a.5.5 0 0 1-.13.649l-.272.205c-.193.146-.283.387-.283.629s.09.483.283.629l.271.205a.5.5 0 0 1 .131.649l-.308.534a.5.5 0 0 1-.628.21l-.313-.131c-.223-.094-.475-.053-.685.069c-.209.121-.374.32-.404.56l-.042.337a.5.5 0 0 1-.496.438h-.618a.5.5 0 0 1-.496-.438l-.042-.337c-.03-.24-.195-.439-.404-.56c-.21-.122-.462-.163-.685-.07l-.313.133a.5.5 0 0 1-.628-.21l-.308-.535a.5.5 0 0 1 .13-.649l.272-.205C9.91 5.983 10 5.742 10 5.5s-.09-.483-.283-.629l-.271-.205a.5.5 0 0 1-.131-.649l.308-.534a.5.5 0 0 1 .628-.21l.313.131c.223.094.475.053.685-.069c.209-.121.374-.32.404-.56l.042-.337A.5.5 0 0 1 12.191 2h.618a.5.5 0 0 1 .496.438l.042.337c.03.24.195.439.404.56M13.5 5.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Admin(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.879 5l1.06-1.06l1.421-1.422a3.5 3.5 0 0 0-3.653 4.674l.326.897l-.675.674l-3.797 3.798a.621.621 0 1 0 .878.878l3.798-3.797l.674-.675l.897.325a3.5 3.5 0 0 0 4.674-3.653L12.06 7.062L11 8.12L9.94 7.06l-1-1zm6.173-1.93A4.987 4.987 0 0 1 15 6a5 5 0 0 1-6.703 4.703L4.5 14.5a2.121 2.121 0 0 1-3-3l3.797-3.797A5 5 0 0 1 13 2l-1.076 1.076l-.863.863L10 5l1 1l1.06-1.06l.864-.864L14 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Api(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.876.165a.25.25 0 0 1 .296-.146l1.022.279a.25.25 0 0 1 .181.276l-.127.893c.551.317.995.774 1.299 1.31l.895-.12a.25.25 0 0 1 .275.185l.27 1.024a.25.25 0 0 1-.15.295l-.836.336a3.501 3.501 0 0 1-.485 1.781l.55.715a.25.25 0 0 1-.022.33l-.751.745a.25.25 0 0 1-.33.02l-.711-.558a3.496 3.496 0 0 1-1.784.47l-.344.835a.25.25 0 0 1-.296.146l-1.022-.278a.25.25 0 0 1-.182-.277l.128-.893a3.496 3.496 0 0 1-1.299-1.31l-.895.12a.25.25 0 0 1-.275-.185l-.27-1.024a.25.25 0 0 1 .15-.295l.836-.336a3.504 3.504 0 0 1 .485-1.781l-.55-.715a.25.25 0 0 1 .022-.33l.751-.745a.25.25 0 0 1 .33-.02l.711.558A3.496 3.496 0 0 1 11.532 1zm1.554 4.86a2 2 0 1 1-3.86-1.05a2 2 0 0 1 3.86 1.05M5.777 6.22A.25.25 0 0 0 5.53 6H4.471a.25.25 0 0 0-.248.219l-.11.88a3.977 3.977 0 0 0-1.244.515l-.7-.544a.25.25 0 0 0-.33.02l-.749.749a.25.25 0 0 0-.02.33l.544.7a3.977 3.977 0 0 0-.515 1.244l-.88.11A.25.25 0 0 0 0 10.47v1.058a.25.25 0 0 0 .219.248l.88.11c.101.448.278.867.515 1.244l-.544.7a.25.25 0 0 0 .02.33l.749.749a.25.25 0 0 0 .33.02l.7-.544c.377.237.796.414 1.244.515l.11.88a.25.25 0 0 0 .247.22h1.058a.25.25 0 0 0 .248-.219l.11-.88a3.974 3.974 0 0 0 1.244-.515l.7.544a.25.25 0 0 0 .33-.02l.749-.749a.25.25 0 0 0 .02-.33l-.544-.7c.237-.377.414-.796.515-1.244l.88-.11a.25.25 0 0 0 .22-.247v-1.058a.25.25 0 0 0-.219-.248l-.88-.11a3.977 3.977 0 0 0-.515-1.244l.544-.7a.25.25 0 0 0-.02-.33l-.75-.75a.25.25 0 0 0-.33-.02l-.7.544A3.977 3.977 0 0 0 5.887 7.1zM7.5 11a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Appearance(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.489 8.388l-.001.006a.115.115 0 0 1-.027.028a.428.428 0 0 1-.264.082h-3.186c-3.118 0-4.68 3.77-2.476 5.974a6.5 6.5 0 1 1 5.953-6.09Zm-.292 1.616c.913 0 1.736-.618 1.79-1.529a8 8 0 1 0-7.032 7.468c1.243-.147 1.527-1.639.641-2.525c-1.26-1.26-.367-3.414 1.415-3.414zM10 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M6 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Applications(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5M1 3a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2zm2 7.5h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5M1 11a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2zm12-8.5h-2a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V3a.5.5 0 0 0-.5-.5M11 1a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm0 9.5h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5M9 11a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Approval(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M7 5a2.99 2.99 0 0 1-.87 2.113A3.997 3.997 0 0 1 8 10.5V12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-1.5c0-1.427.747-2.679 1.87-3.387A3 3 0 1 1 7 5m-5.5 5.5a2.5 2.5 0 0 1 5 0V12a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5zm14.28-5.22a.75.75 0 0 0-1.06-1.06L12 6.94l-1.22-1.22a.75.75 0 1 0-1.06 1.06l1.75 1.75a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApprovalSolid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M7 5a2.99 2.99 0 0 1-.87 2.113A3.997 3.997 0 0 1 8 10.5V12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-1.5c0-1.427.747-2.679 1.87-3.387A3 3 0 1 1 7 5m8.78.28a.75.75 0 0 0-1.06-1.06L12 6.94l-1.22-1.22a.75.75 0 1 0-1.06 1.06l1.75 1.75a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4.5v-2h13v2zM1 6a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1v8a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1zm1.5 0v7.5h11V6zM4 8.75A.75.75 0 0 1 4.75 8h6.5a.75.75 0 0 1 0 1.5h-6.5A.75.75 0 0 1 4 8.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.72 10.159a.75.75 0 1 1 1.06 1.06l-3.25 3.25L8 15l-.53-.53l-3.25-3.25a.75.75 0 0 1 1.06-1.061l1.97 1.97V1.75a.75.75 0 1 1 1.5 0v10.379z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.841 5.28a.75.75 0 0 0-1.06-1.06L1.53 7.47L1 8l.53.53l3.25 3.25a.75.75 0 0 0 1.061-1.06l-1.97-1.97H14.25a.75.75 0 0 0 0-1.5H3.871z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.159 10.72a.75.75 0 1 0 1.06 1.06l3.25-3.25L15 8l-.53-.53l-3.25-3.25a.75.75 0 0 0-1.061 1.06l1.97 1.97H1.75a.75.75 0 1 0 0 1.5h10.379z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.72 5.841a.75.75 0 1 0 1.06-1.06L8.53 1.53L8 1l-.53.53l-3.25 3.25a.75.75 0 0 0 1.06 1.061l1.97-1.97V14.25a.75.75 0 0 0 1.5 0V3.871z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Assignee(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M7 5a2.99 2.99 0 0 1-.87 2.113A3.997 3.997 0 0 1 8 10.5V12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-1.5c0-1.427.747-2.679 1.87-3.387A3 3 0 1 1 7 5m-5.5 5.5a2.5 2.5 0 0 1 5 0V12a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5zM13 10a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 13 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.774 1.747a6.5 6.5 0 1 0 1.142 12.062a.75.75 0 0 1 .673 1.34A8 8 0 1 1 16 8v1.25a2.75 2.75 0 0 1-5.072 1.475A4 4 0 1 1 12 8v1.25a1.25 1.25 0 0 0 2.5 0V8a6.5 6.5 0 0 0-4.726-6.253M10.5 8a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attention(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.04 8l-3.174 5.5a1 1 0 0 0 .866 1.5h2.173a1 1 0 0 0 .724-.31L13 8L6.629 1.31A1 1 0 0 0 5.905 1H3.732a1 1 0 0 0-.866 1.5zm4.889 0L5.69 13.5H4.598L7.34 8.75L7.773 8l-.433-.75L4.598 2.5H5.69z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttentionSolid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 8L6.629 1.31A1 1 0 0 0 5.905 1H3.732a1 1 0 0 0-.866 1.5L6.04 8l-3.175 5.5a1 1 0 0 0 .866 1.5h2.173a1 1 0 0 0 .724-.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttentionSolidSm(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.004 8L7.538 3.31A1 1 0 0 0 6.813 3H5.732a1 1 0 0 0-.866 1.5L6.886 8l-2.02 3.5a1 1 0 0 0 .866 1.5h1.081a1 1 0 0 0 .725-.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Autoplay(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.194.665a8 8 0 1 0 4.194 10.404a.75.75 0 0 0-1.385-.575a6.5 6.5 0 1 1-.526-5.994H11.75a.75.75 0 0 0 0 1.5H16V1.75a.75.75 0 0 0-1.5 0v1.586a8 8 0 0 0-3.306-2.67Zm-.423 6.913a.5.5 0 0 1 0 .844l-4.003 2.54A.5.5 0 0 1 6 10.538V5.461a.5.5 0 0 1 .768-.422z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitbucket(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.454 1.816a.452.452 0 0 0-.345.153a.435.435 0 0 0-.103.358L2.91 13.684c.049.287.3.498.596.5h9.135a.447.447 0 0 0 .449-.37L14.994 2.33a.435.435 0 0 0-.1-.358a.452.452 0 0 0-.342-.156zm8.018 8.208H6.557l-.79-4.05h4.412z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h5.5a3.5 3.5 0 0 0 1.852-6.47A3.5 3.5 0 0 0 8.5 2zm4.5 5a1.5 1.5 0 1 0 0-3H5v3zM5 9v3h4.5a1.5 1.5 0 0 0 0-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 2a2 2 0 0 1 2-2H13a.75.75 0 0 1 .75.75v10.5A.75.75 0 0 1 13 12H7.5v1.5H13a.75.75 0 0 1 0 1.5H7.5v1H6v-1H4.5a2.25 2.25 0 0 1-2.25-2.25zM6 13.5V12H4.5a.75.75 0 0 0 0 1.5zm6.25-12v9H4.5c-.263 0-.515.045-.75.128V2a.5.5 0 0 1 .5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8 9.293l1.06 1.06l3.44 3.44V2a.5.5 0 0 0-.5-.5H4a.5.5 0 0 0-.5.5v11.793l3.44-3.44zm1.06 3.182l3.233 3.232c.63.63 1.707.184 1.707-.707V2a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v13c0 .89 1.077 1.337 1.707.707l3.232-3.232L8 11.415z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Branch(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m2.5-1a2.501 2.501 0 0 1-1.872 2.42A3.502 3.502 0 0 1 8.75 8.5h-1.5a2 2 0 0 0-1.965 1.626a2.501 2.501 0 1 1-1.535-.011v-4.23a2.501 2.501 0 1 1 1.5 0v1.742a3.484 3.484 0 0 1 2-.627h1.5a2 2 0 0 0 1.823-1.177A2.5 2.5 0 1 1 14 3.5m-8.5 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0m0-9a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BranchDeleted(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.28 1.22a.75.75 0 1 0-1.06 1.06l1.22 1.22l-1.22 1.22a.75.75 0 0 0 1.06 1.06l1.22-1.22l1.22 1.22a.75.75 0 1 0 1.06-1.06L12.56 3.5l1.22-1.22a.75.75 0 0 0-1.06-1.06L11.5 2.44zm-2.5 6a.75.75 0 0 1 0 1.06l-2.045 2.046A2.499 2.499 0 0 1 4.5 15a2.5 2.5 0 0 1-.75-4.886V5.886a2.501 2.501 0 1 1 1.5 0v2.803l1.47-1.47a.75.75 0 0 1 1.06 0ZM5.5 12.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m0-9a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandZoom(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.45 3.334C.648 3.334 0 3.982 0 4.783v4.986c0 1.6 1.298 2.898 2.898 2.898h6.986c.8 0 1.45-.649 1.45-1.45V6.233a2.9 2.9 0 0 0-2.899-2.899zM16 4.643v6.715c0 .544-.618.86-1.059.539l-2.059-1.498a1.333 1.333 0 0 1-.549-1.078V6.679c0-.427.204-.827.55-1.078l2.058-1.498a.667.667 0 0 1 1.059.54" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 1.75A.75.75 0 0 1 4.75 1h.5C6.216 1 7 1.784 7 2.75v.35a5.023 5.023 0 0 1 2 0v-.35C9 1.784 9.784 1 10.75 1h.5a.75.75 0 0 1 0 1.5h-.5a.25.25 0 0 0-.25.25v.919A5.02 5.02 0 0 1 12.584 6h1.666a.75.75 0 0 1 0 1.5h-1.275c.017.164.025.331.025.5v.5h1.25a.75.75 0 0 1 0 1.5H13c0 .342-.034.677-.1 1h1.35a.75.75 0 0 1 0 1.5h-1.919A4.998 4.998 0 0 1 8 15a4.998 4.998 0 0 1-4.331-2.5H1.75a.75.75 0 0 1 0-1.5H3.1a5.022 5.022 0 0 1-.1-1H1.75a.75.75 0 0 1 0-1.5H3V8c0-.169.008-.336.025-.5H1.75a.75.75 0 0 1 0-1.5h1.666A5.02 5.02 0 0 1 5.5 3.669V2.75a.25.25 0 0 0-.25-.25h-.5A.75.75 0 0 1 4 1.75M11.5 8v2a3.501 3.501 0 0 1-2.75 3.42V7.5h2.715c.023.163.035.33.035.5m-4.25-.5v5.92A3.501 3.501 0 0 1 4.5 10V8c0-.17.012-.337.035-.5zM10.873 6H5.127A3.496 3.496 0 0 1 8 4.5c1.19 0 2.24.593 2.873 1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 14.5v-13h9v13H11V11a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v3.5zm3 0h3v-3h-3zM2 1a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1zm3 2.75A.75.75 0 0 1 5.75 3h1a.75.75 0 0 1 0 1.5h-1A.75.75 0 0 1 5 3.75M9.25 3a.75.75 0 0 0 0 1.5h1a.75.75 0 0 0 0-1.5zM5 7.25a.75.75 0 0 1 .75-.75h1a.75.75 0 0 1 0 1.5h-1A.75.75 0 0 1 5 7.25m4.25-.75a.75.75 0 0 0 0 1.5h1a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bulb(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.17 8.533C2.21 5.5 4.388 1.5 8 1.5s5.79 4 3.83 7.033L9.592 12H6.408zM5 12.584L2.91 9.347C.305 5.315 3.2 0 8 0s7.694 5.315 5.09 9.347L11 12.584V14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm1.5.916v.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullhorn(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.5 8.974l6 2.333V2.693l-6 2.333zm6.138-7.944L7 4H3a3 3 0 0 0 0 6h.969l1.066 4.445A2.029 2.029 0 0 0 7.008 16c1.524 0 2.504-1.618 1.797-2.969l-1.53-2.924l7.363 2.863A1 1 0 0 0 16 12.038V1.962a1 1 0 0 0-1.362-.932M5.51 10h.016l1.95 3.726a.529.529 0 1 1-.983.369L5.511 10ZM3 5.5h4v3H3a1.5 1.5 0 1 1 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.75 1a.75.75 0 0 1 .75.75V3h5V1.75a.75.75 0 0 1 1.5 0V3h2a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h2V1.75A.75.75 0 0 1 4.75 1M2.5 4.5V6h11V4.5zm0 9v-6h11v6zM11 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.035 13.096a6.5 6.5 0 0 1-9.131-9.131zm1.061-1.06L3.965 2.903a6.5 6.5 0 0 1 9.131 9.131ZM16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CanceledCircle(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M5.28 4.22a.75.75 0 0 0-1.06 1.06l6.5 6.5a.75.75 0 1 0 1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 2.5H5a.5.5 0 0 0-.5.5v2h7V3a.5.5 0 0 0-.5-.5M1.75 4.25H3V3a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1.25h1.25a.75.75 0 0 1 0 1.5h-.918A3.995 3.995 0 0 1 15 9v5a1 1 0 1 1-2 0v-1H3v1a1 1 0 1 1-2 0V9c0-1.339.658-2.524 1.668-3.25H1.75a.75.75 0 0 1 0-1.5M12 11.5h1.5V9A2.5 2.5 0 0 0 11 6.5H5A2.5 2.5 0 0 0 2.5 9v2.5H4V11a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CatalogCheckmark(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.75 1a.75.75 0 1 0 0 1.5h4.5a.75.75 0 1 0 0-1.5zm-2 2.5a.75.75 0 1 0 0 1.5h8.5a.75.75 0 1 0 0-1.5zm-1.25 4v6h11v-6zM2 6a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1zm8.28 3.83a.75.75 0 1 0-1.06-1.06l-1.97 1.97l-.47-.47a.75.75 0 1 0-1.06 1.06l1 1a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 1a.75.75 0 0 1 .75.75v9.75h9.75a.75.75 0 0 1 0 1.5H4.5v1.25a.75.75 0 0 1-1.5 0V13H1.75a.75.75 0 0 1 0-1.5H3V1.75A.75.75 0 0 1 3.75 1m9.75 4.75a.75.75 0 0 0-1.5 0v3.5a.75.75 0 0 0 1.5 0zM9.75 3a.75.75 0 0 1 .75.75v5.5a.75.75 0 0 1-1.5 0v-5.5A.75.75 0 0 1 9.75 3M7.5 7.75a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.78 4.62a.75.75 0 0 1 0 1.06l-6.097 6.097a.75.75 0 0 1-1.069-.009L3.211 9.284a.75.75 0 1 1 1.078-1.043l1.873 1.936L11.72 4.62a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0m-4.22-1.72a.75.75 0 0 0-1.06-1.06L6.75 9.19L5.53 7.97a.75.75 0 0 0-1.06 1.06l1.75 1.75a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleDashed(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.701 1.08a.75.75 0 0 0-.934-.5A7.75 7.75 0 0 0 .573 5.785a.75.75 0 1 0 1.438.429a6.25 6.25 0 0 1 4.188-4.199a.75.75 0 0 0 .502-.934Zm2.598 0a.75.75 0 0 1 .934-.501a7.75 7.75 0 0 1 5.194 5.206a.75.75 0 1 1-1.438.429a6.25 6.25 0 0 0-4.188-4.199a.75.75 0 0 1-.502-.934Zm2.481 4.14a.75.75 0 0 1 0 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0L4.47 9.03a.75.75 0 0 1 1.06-1.06l1.22 1.22l3.97-3.97a.75.75 0 0 1 1.06 0m-1.547 10.2a.75.75 0 1 1-.432-1.436a6.251 6.251 0 0 0 4.188-4.199a.75.75 0 1 1 1.438.429a7.75 7.75 0 0 1-5.194 5.206m-4.466 0a.75.75 0 1 0 .432-1.436a6.25 6.25 0 0 1-4.188-4.199a.75.75 0 1 0-1.438.429a7.75 7.75 0 0 0 5.194 5.206" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleFilled(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16m3.78-9.72a.75.75 0 0 0-1.06-1.06L6.75 9.19L5.53 7.97a.75.75 0 0 0-1.06 1.06l1.75 1.75a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSm(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.78 2.62a.75.75 0 0 1 0 1.06L4.683 9.777a.75.75 0 0 1-1.069-.009L1.211 7.284a.75.75 0 0 1 1.078-1.043l1.873 1.936L9.72 2.62a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckXs(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.53 3.22a.75.75 0 0 1 0 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0l-1.5-1.5a.75.75 0 0 1 1.06-1.06l.97.97l3.97-3.97a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CherryPickCommit(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.75.75a.75.75 0 0 0-1.5 0v2.528a10.555 10.555 0 0 0-2.55.51a9.11 9.11 0 0 0-.952.375a6.415 6.415 0 0 0-.337.167l-.022.012l-.007.004l-.002.002h-.001L3.75 5l-.372-.651A.75.75 0 0 0 3 5v2.25a.75.75 0 0 0 1.5 0V5.47A9.002 9.002 0 0 1 8 4.75a9.001 9.001 0 0 1 3.5.72v1.78a.75.75 0 0 0 1.5 0V5a.75.75 0 0 0-.378-.651L12.25 5l.372-.651l-.002-.001l-.002-.002l-.007-.004l-.022-.012a6.42 6.42 0 0 0-.337-.168a9.315 9.315 0 0 0-.952-.374a10.555 10.555 0 0 0-2.55-.51zM10.5 12a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m1.43.75a4.001 4.001 0 0 1-7.86 0H.75a.75.75 0 0 1 0-1.5h3.32a4.001 4.001 0 0 1 7.86 0h3.32a.75.75 0 0 1 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLgLeft(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.721 2.22a.75.75 0 0 1 1.061 1.06L4.061 8.002l4.721 4.721a.75.75 0 0 1-1.06 1.061L2.47 8.532a.75.75 0 0 1 0-1.06L7.722 2.22Zm5 0a.75.75 0 0 1 1.061 1.06L9.061 8.002l4.721 4.721a.75.75 0 0 1-1.06 1.061L7.47 8.532a.75.75 0 0 1 0-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLgRight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.53 2.22a.75.75 0 0 0-1.06 1.06l4.72 4.722l-4.72 4.721a.75.75 0 0 0 1.06 1.061l5.252-5.252a.75.75 0 0 0 0-1.06zm5 0a.75.75 0 0 0-1.06 1.06l4.721 4.722l-4.721 4.721a.75.75 0 0 0 1.06 1.061l5.252-5.252a.75.75 0 0 0 0-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.22 6.22a.75.75 0 0 1 1.06 0L8 8.94l2.72-2.72a.75.75 0 1 1 1.06 1.06l-3.25 3.25a.75.75 0 0 1-1.06 0L4.22 7.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.78 4.22a.75.75 0 0 1 0 1.06L7.06 8l2.72 2.72a.75.75 0 1 1-1.06 1.06L5.47 8.53a.75.75 0 0 1 0-1.06l3.25-3.25a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLgDown(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.22 5.22a.75.75 0 0 0 0 1.06l5.252 5.252a.75.75 0 0 0 1.06 0l5.252-5.252a.75.75 0 1 0-1.06-1.06L8.001 9.94L3.28 5.22a.75.75 0 0 0-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLgLeft(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.78 2.22a.75.75 0 0 0-1.06 0L4.468 7.472a.75.75 0 0 0 0 1.06l5.252 5.252a.75.75 0 1 0 1.06-1.06L6.06 8.001l4.72-4.721a.75.75 0 0 0 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLgRight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.22 2.22a.75.75 0 0 1 1.06 0l5.252 5.252a.75.75 0 0 1 0 1.06L6.28 13.784a.75.75 0 1 1-1.06-1.06l4.72-4.723L5.22 3.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLgUp(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.22 10.78a.75.75 0 0 1 0-1.06l5.252-5.252a.75.75 0 0 1 1.06 0l5.252 5.252a.75.75 0 1 1-1.06 1.06L8.001 6.06L3.28 10.78a.75.75 0 0 1-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.22 4.22a.75.75 0 0 0 0 1.06L8.94 8l-2.72 2.72a.75.75 0 1 0 1.06 1.06l3.25-3.25a.75.75 0 0 0 0-1.06L7.28 4.22a.75.75 0 0 0-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.22 9.78a.75.75 0 0 0 1.06 0L8 7.06l2.72 2.72a.75.75 0 1 0 1.06-1.06L8.53 5.47a.75.75 0 0 0-1.06 0L4.22 8.72a.75.75 0 0 0 0 1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clear(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M4.22 4.22a.75.75 0 0 1 1.06 0L8 6.94l2.72-2.72a.75.75 0 1 1 1.06 1.06L9.06 8l2.72 2.72a.75.75 0 1 1-1.06 1.06L8 9.06l-2.72 2.72a.75.75 0 0 1-1.06-1.06L6.94 8L4.22 5.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClearAll(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.963 7.23A8 8 0 0 1 .044 8.841a.75.75 0 0 1 1.492-.158a6.5 6.5 0 1 0 9.964-6.16V4.25a.75.75 0 0 1-1.5 0V0h4.25a.75.75 0 0 1 0 1.5h-1.586a8.001 8.001 0 0 1 3.299 5.73M7 2a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-2.25.25a1 1 0 1 1-2 0a1 1 0 0 1 2 0M1.5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M8.75 3.75a.75.75 0 0 0-1.5 0v4.56l.22.22l2.254 2.254a.75.75 0 1 0 1.06-1.06L8.75 7.689z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.28 3.22a.75.75 0 0 0-1.06 1.06L6.94 8l-3.72 3.72a.75.75 0 1 0 1.06 1.06L8 9.06l3.72 3.72a.75.75 0 1 0 1.06-1.06L9.06 8l3.72-3.72a.75.75 0 0 0-1.06-1.06L8 6.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseXs(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.28 3.22a.75.75 0 0 0-1.06 1.06L4.94 6L3.22 7.72a.75.75 0 0 0 1.06 1.06L6 7.06l1.72 1.72a.75.75 0 0 0 1.06-1.06L7.06 6l1.72-1.72a.75.75 0 0 0-1.06-1.06L6 4.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudGear(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5a3.75 3.75 0 0 0-3.745 3.552l-.04.787l-.785-.08a1.75 1.75 0 0 0-1.055 3.257a.75.75 0 0 1-.75 1.299a3.249 3.249 0 0 1 1.215-6.04a5.251 5.251 0 0 1 10.037-.973a3.75 3.75 0 0 1 1.561 6.744a.75.75 0 0 1-.876-1.218a2.25 2.25 0 0 0-1.256-4.077l-.55-.014l-.152-.528A3.752 3.752 0 0 0 8 2.5m-.803 3.924c-.05.405-.355.724-.731.88c-.323.135-.624.31-.898.52c-.324.248-.752.352-1.129.193a.485.485 0 0 0-.608.204l-.322.558a.485.485 0 0 0 .127.629c.325.246.45.668.397 1.073a4.04 4.04 0 0 0 0 1.038c.053.405-.072.827-.397 1.073a.485.485 0 0 0-.127.629l.322.558c.122.212.383.3.608.204c.377-.159.805-.055 1.129.193c.274.21.575.385.898.52c.376.156.68.475.731.88c.03.242.236.424.48.424h.645c.245 0 .45-.182.481-.424c.05-.405.355-.724.731-.88c.323-.135.624-.31.898-.52c.324-.248.752-.352 1.128-.193a.485.485 0 0 0 .61-.204l.321-.558a.485.485 0 0 0-.127-.629c-.325-.246-.45-.668-.397-1.073a4.046 4.046 0 0 0 0-1.038c-.053-.405.072-.827.397-1.073a.485.485 0 0 0 .127-.629l-.322-.558a.485.485 0 0 0-.609-.204c-.376.159-.804.055-1.128-.193a3.994 3.994 0 0 0-.898-.52c-.376-.156-.68-.475-.731-.88A.484.484 0 0 0 8.323 6h-.645a.484.484 0 0 0-.481.424M10.5 11a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0M8 10a1 1 0 1 1 0 2a1 1 0 0 1 0-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudPod(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.255 6.052a3.75 3.75 0 0 1 7.349-.843l.152.528l.55.014a2.25 2.25 0 0 1 1.256 4.077a.75.75 0 0 0 .876 1.218a3.75 3.75 0 0 0-1.561-6.744a5.251 5.251 0 0 0-10.037.974a3.25 3.25 0 0 0-1.216 6.039a.75.75 0 1 0 .752-1.299A1.749 1.749 0 0 1 3.43 6.76l.784.08l.041-.787ZM8 9.838l-1.732-.99L8 7.858l1.732.99zm.75 1.299l1.75-1v1.98l-1.75 1zm-1.5 0l-1.75-1v1.98l1.75 1zm.254-4.723a1 1 0 0 1 .992 0l3 1.714a1 1 0 0 1 .504.868v3.41a1 1 0 0 1-.504.87l-3 1.713a1 1 0 0 1-.992 0l-3-1.714A1 1 0 0 1 4 12.407v-3.41a1 1 0 0 1 .504-.87z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudTerminal(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.256 6.041a3.75 3.75 0 0 1 7.348-.832l.152.528l.55.014a2.25 2.25 0 0 1 1.069 4.198a.75.75 0 1 0 .75 1.299a3.75 3.75 0 0 0-1.25-6.946a5.251 5.251 0 0 0-10.035.974a3.25 3.25 0 0 0-.896 6.2a.75.75 0 1 0 .603-1.373A1.75 1.75 0 0 1 3.25 6.75h.967zM6.22 7.22a.75.75 0 0 1 1.06 0l1.75 1.75l.53.53l-.53.53l-1.75 1.75a.75.75 0 0 1-1.06-1.06L7.44 9.5L6.22 8.28a.75.75 0 0 1 0-1.06M8 13.25a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.424 2.023a.75.75 0 0 1 .556.904L7.48 13.42a.75.75 0 0 1-1.46-.348L8.52 2.58a.75.75 0 0 1 .904-.556ZM11.16 4.22a.75.75 0 0 1 1.06 0l3.25 3.25L16 8l-.53.53l-3.25 3.25a.75.75 0 1 1-1.06-1.06L13.88 8l-2.72-2.72a.75.75 0 0 1 0-1.06M4.84 5.28a.75.75 0 1 0-1.06-1.06L.53 7.47L0 8l.53.53l3.25 3.25a.75.75 0 0 0 1.06-1.06L2.12 8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collapse(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 3.16a.75.75 0 0 0-1.06 0L8.75 5.13V.75a.75.75 0 0 0-1.5 0v4.38L5.28 3.16a.75.75 0 0 0-1.06 1.06l3.25 3.25L8 8l-.53.53l-3.25 3.25a.75.75 0 0 0 1.06 1.061l1.97-1.97v4.379a.75.75 0 0 0 1.5 0v-4.379l1.97 1.97a.75.75 0 1 0 1.06-1.06L8.53 8.53L8 8l.53-.53l3.25-3.25a.75.75 0 0 0 0-1.061Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollapseLeft(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 12.25a.75.75 0 0 0 1.5 0v-8.5a.75.75 0 0 0-1.5 0zm7.841-8.03a.75.75 0 0 1 0 1.06l-1.97 1.97h9.379a.75.75 0 0 1 0 1.5H5.871l1.97 1.97a.75.75 0 1 1-1.06 1.06L3.53 8.53L3 8l.53-.53l3.25-3.25a.75.75 0 0 1 1.061 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollapseRight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 3.75a.75.75 0 0 1 1.5 0v8.5a.75.75 0 0 1-1.5 0zm-6.341.47a.75.75 0 0 0 0 1.06l1.97 1.97H.75a.75.75 0 0 0 0 1.5h9.379l-1.97 1.97a.75.75 0 1 0 1.06 1.06l3.25-3.25L13 8l-.53-.53l-3.25-3.25a.75.75 0 0 0-1.061 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollapseSolid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M8 8l-.53-.53l-2.25-2.25a.75.75 0 0 1 1.06-1.061l.97.97v-2.38a.75.75 0 0 1 1.5 0v2.38l.97-.97a.75.75 0 1 1 1.06 1.06L8.53 7.47zl.53.53l2.25 2.25a.75.75 0 0 1-1.06 1.061l-.97-.97v2.38a.75.75 0 0 1-1.5 0v-2.38l-.97.97a.75.75 0 1 1-1.06-1.06l2.25-2.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5A1.5 1.5 0 0 0 1.5 4v6A1.5 1.5 0 0 0 3 11.5h1.5v1.6l2.067-1.462l.194-.138H13a1.5 1.5 0 0 0 1.5-1.5V4A1.5 1.5 0 0 0 13 2.5zM0 4a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H7.239l-3.056 2.162L3 16v-3a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDots(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4A1.5 1.5 0 0 1 3 2.5h10A1.5 1.5 0 0 1 14.5 4v6a1.5 1.5 0 0 1-1.5 1.5H6.761l-.194.138L4.5 13.1v-1.6H3A1.5 1.5 0 0 1 1.5 10zM3 1a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3v3l1.183-.838L7.24 13H13a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3zm8 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2M9 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0M5 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentLines(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5A1.5 1.5 0 0 0 1.5 4v6A1.5 1.5 0 0 0 3 11.5h1.5v1.6l2.067-1.462l.194-.138H13a1.5 1.5 0 0 0 1.5-1.5V4A1.5 1.5 0 0 0 13 2.5zM0 4a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H7.239l-3.056 2.162L3 16v-3a3 3 0 0 1-3-3zm3 4.25a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 0 1.5h-5.5A.75.75 0 0 1 3 8.25M3.75 5a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentNext(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5A1.5 1.5 0 0 0 1.5 4v6A1.5 1.5 0 0 0 3 11.5h1.5v1.6l2.067-1.462l.194-.138H13a1.5 1.5 0 0 0 1.5-1.5V4A1.5 1.5 0 0 0 13 2.5zM0 4a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H7.239l-3.056 2.162L3 16v-3a3 3 0 0 1-3-3zm9.191 3.75H4.75a.75.75 0 0 1 0-1.5h4.441a.75.75 0 0 1 1.09-1.03l1.249 1.25l.53.53l-.53.53l-1.25 1.25a.75.75 0 0 1-1.089-1.03" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 0a2 2 0 0 0-2 2v10.06l1.28-1.28l1.53-1.53H4V11a2 2 0 0 0 2 2h7l1.5 1.5L16 16V6a2 2 0 0 0-2-2h-2V2a2 2 0 0 0-2-2zm8.5 4V2a.5.5 0 0 0-.5-.5H2a.5.5 0 0 0-.5.5v6.44l.47-.47l.22-.22H4V6a2 2 0 0 1 2-2zm3.56 7.94l.44.439V6a.5.5 0 0 0-.5-.5H6a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h7.621z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Commit(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 10.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5M8 12c1.953 0 3.579-1.4 3.93-3.25h3.32a.75.75 0 0 0 0-1.5h-3.32a4.001 4.001 0 0 0-7.86 0H.75a.75.75 0 0 0 0 1.5h3.32A4.001 4.001 0 0 0 8 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comparison(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25 2.386a2.501 2.501 0 1 0-1.5 0V10.5a2.75 2.75 0 0 0 2.75 2.75h.629l-.47.47a.75.75 0 1 0 1.061 1.06l1.75-1.75l.53-.53l-.53-.53l-1.75-1.75a.747.747 0 0 0-1.061 0a.75.75 0 0 0 0 1.06l.47.47H5.5c-.69 0-1.25-.56-1.25-1.25zM9.28 1.22a.75.75 0 1 1 1.06 1.06l-.47.47h.63a2.75 2.75 0 0 1 2.75 2.75v4.614a2.501 2.501 0 1 1-1.5 0V5.5c0-.69-.56-1.25-1.25-1.25h-.63l.47.47a.75.75 0 1 1-1.06 1.06L7.53 4.03L7 3.5l.53-.53zM12.5 13.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M7.186 5.605L12 4l-1.605 4.814a2.5 2.5 0 0 1-1.58 1.581L4 12l1.605-4.814a2.5 2.5 0 0 1 1.58-1.581ZM9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Connected(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.78 2.28a.75.75 0 0 0-1.06-1.06l-1.617 1.616a4 4 0 0 0-5.275.336L5.53 4.47l-.25-.25a.75.75 0 0 0-1.06 1.06l.25.25l-1.302 1.302a4 4 0 0 0-.336 5.275L1.22 13.72a.75.75 0 1 0 1.06 1.06l1.613-1.612a4 4 0 0 0 5.275-.336l1.302-1.302l.25.25a.75.75 0 1 0 1.06-1.06l-.25-.25l1.298-1.298a4 4 0 0 0 .336-5.275zm-4.31 7.13l1.297-1.297a2.5 2.5 0 0 0 0-3.536l-.343-.343l.404-.405l-.404.405a2.5 2.5 0 0 0-3.536 0L6.591 5.53zM5.53 6.591l-1.3 1.301a2.5 2.5 0 0 0 0 3.536l-.4.4l.4-.4l.343.343a2.5 2.5 0 0 0 3.536 0L9.41 10.47L5.53 6.59Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContainerImage(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.38 1.103a.75.75 0 0 0-.76 0L.37 5.353a.75.75 0 0 0 0 1.294l7.25 4.25a.75.75 0 0 0 .76 0l7.25-4.25a.75.75 0 0 0 0-1.294zM8 9.381L2.233 6L8 2.62L13.767 6zm-6.87-.028a.75.75 0 0 0-.76 1.294l7.25 4.25a.75.75 0 0 0 .76 0l7.25-4.25a.75.75 0 0 0-.76-1.294L8 13.381z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyToClipboard(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M9.45 2a2.5 2.5 0 0 0-4.9 0H3a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h2v-1.5H3.5v-9h1V5h5V3.5h1V7H12V3a1 1 0 0 0-1-1zM7.5 9.5h1.25a.75.75 0 0 0 0-1.5h-1.5C6.56 8 6 8.56 6 9.25v1.5a.75.75 0 0 0 1.5 0zm1.25 5H7.5v-1.25a.75.75 0 0 0-1.5 0v1.5c0 .69.56 1.25 1.25 1.25h1.5a.75.75 0 0 0 0-1.5m3.75-5h-1.25a.75.75 0 0 1 0-1.5h1.5c.69 0 1.25.56 1.25 1.25v1.5a.75.75 0 0 1-1.5 0zm-1.25 5h1.25v-1.25a.75.75 0 0 1 1.5 0v1.5c0 .69-.56 1.25-1.25 1.25h-1.5a.75.75 0 0 1 0-1.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 3.5H2a.5.5 0 0 0-.5.5v1h13V4a.5.5 0 0 0-.5-.5M1.5 12V6.5h13V12a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5M2 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm1.75 7.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dash(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.545 2.1a.75.75 0 0 1 .274 1.025l-3.472 6.007a3 3 0 1 1-1.208-.908l1.759-3.042a6.5 6.5 0 0 0-2.148-.639V5a.75.75 0 1 1-1.5 0v-.457a6.5 6.5 0 0 0-1.829.49l.229.396a.75.75 0 1 1-1.3.75l-.228-.396a6.5 6.5 0 0 0-1.339 1.34l.396.227a.75.75 0 0 1-.75 1.3l-.396-.229a6.5 6.5 0 0 0-.498 1.905a.75.75 0 0 1-1.492-.155A8 8 0 0 1 11.65 3.88l.87-1.506a.75.75 0 0 1 1.025-.274m-.107 4.047a.75.75 0 0 1 1.047.169a8 8 0 0 1 1.51 4.963a.75.75 0 1 1-1.499-.052a6.5 6.5 0 0 0-1.227-4.033a.75.75 0 0 1 .17-1.047ZM9.5 11a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deployments(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 1a.75.75 0 0 0 0 1.5h6.5a3.25 3.25 0 0 1 3.25 3.25V9.5h-.75a.75.75 0 0 0 0 1.5H15V5.75A4.75 4.75 0 0 0 10.25 1zM13 14.25a.75.75 0 0 1-.75.75h-6.5A4.75 4.75 0 0 1 1 10.25V5h2.25a.75.75 0 0 1 0 1.5H2.5v3.75a3.25 3.25 0 0 0 3.25 3.25h6.5a.75.75 0 0 1 .75.75M6.22 5.22a.75.75 0 0 1 1.06 0l2.25 2.25l.53.53l-.53.53l-2.25 2.25a.75.75 0 1 1-1.06-1.06L7.94 8L6.22 6.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DetailsBlock(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 4.13v1.428a.5.5 0 0 0 .725.446l.886-.446l.377-.19L2 5.362l1.404-.708l.07-.036l.662-.333l.603-.304a.5.5 0 0 0 0-.893l-.603-.305l-.662-.333l-.07-.036L2 1.706l-.012-.005l-.377-.19l-.886-.447A.5.5 0 0 0 0 1.51zM7.25 2a.75.75 0 0 0 0 1.5h7a.25.25 0 0 1 .25.25v8.5a.25.25 0 0 1-.25.25h-9.5a.25.25 0 0 1-.25-.25V6.754a.75.75 0 0 0-1.5 0v5.496c0 .966.784 1.75 1.75 1.75h9.5A1.75 1.75 0 0 0 16 12.25v-8.5A1.75 1.75 0 0 0 14.25 2zm-.5 4a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5zM6 9.25a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5A.75.75 0 0 1 6 9.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diagram(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 5.5v-3h5v3zM4 2a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-.103l1.535 2H14a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h.541L9.007 7H6.993L5.46 9H6a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h1.568l1.535-2H5a1 1 0 0 1-1-1zm-1.5 8.5v3h3v-3zm8 3v-3h3v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discord(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.554 3.016A13.233 13.233 0 0 0 10.253 2a9.068 9.068 0 0 0-.423.86a12.293 12.293 0 0 0-3.664 0A9.112 9.112 0 0 0 5.744 2A13.358 13.358 0 0 0 2.44 3.018C.351 6.108-.215 9.123.068 12.094a13.306 13.306 0 0 0 4.048 2.033a9.78 9.78 0 0 0 .867-1.399a8.605 8.605 0 0 1-1.365-.652c.115-.083.227-.168.335-.251a9.51 9.51 0 0 0 8.094 0c.11.09.222.175.335.251a8.648 8.648 0 0 1-1.368.654a9.7 9.7 0 0 0 .867 1.396a13.248 13.248 0 0 0 4.051-2.03c.332-3.446-.568-6.433-2.379-9.08m-8.212 7.25c-.789 0-1.44-.715-1.44-1.596c0-.881.628-1.603 1.438-1.603c.809 0 1.456.722 1.442 1.603c-.014.88-.636 1.597-1.44 1.597m5.316 0c-.79 0-1.44-.715-1.44-1.596c0-.881.63-1.603 1.44-1.603c.81 0 1.452.722 1.438 1.603c-.014.88-.634 1.597-1.438 1.597"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disk(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5h10a.5.5 0 0 1 .5.5v5.063A2.004 2.004 0 0 0 13 8H3c-.173 0-.34.022-.5.063V3a.5.5 0 0 1 .5-.5M2.5 10v3a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5H3a.5.5 0 0 0-.5.5M1 10V3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2zm11 1.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-4 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocChanges(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 13.5v-12H8v2.75C8 5.216 8.784 6 9.75 6h3.375a.76.76 0 0 0 .063-.003A.75.75 0 0 0 14 5.25v-.774a1 1 0 0 0-.282-.695L10.363.305A1 1 0 0 0 9.643 0H3a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h4.25a.75.75 0 0 0 0-1.5zm8.828-9L9.5 1.57v2.68c0 .138.112.25.25.25zM10 15.25a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75m3-2a.75.75 0 0 1-.75-.75V11h-1.5a.75.75 0 0 1 0-1.5h1.5V8a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5a.75.75 0 0 1-.75.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocChart(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8v4.5h-13v-9H10v2.75c0 .966.784 1.75 1.75 1.75zm-3-4.379L14.379 6.5H11.75a.25.25 0 0 1-.25-.25zM1 2a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V6.414a1 1 0 0 0-.293-.707l-3.414-3.414A1 1 0 0 0 11.586 2zm2.75 5a.75.75 0 0 0-.75.75v2.5a.75.75 0 0 0 1.5 0v-2.5A.75.75 0 0 0 3.75 7M6 5.75a.75.75 0 0 1 1.5 0v4.5a.75.75 0 0 1-1.5 0zM9.75 9a.75.75 0 0 0-.75.75v.5a.75.75 0 0 0 1.5 0v-.5A.75.75 0 0 0 9.75 9" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocCode(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 14.5V6H9.75A1.75 1.75 0 0 1 8 4.25V1.5H3.5v13zm-.121-10L9.5 1.621V4.25c0 .138.112.25.25.25zM2 1a1 1 0 0 1 1-1h6.586a1 1 0 0 1 .707.293l3.414 3.414a1 1 0 0 1 .293.707V15a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1zm5.75 10.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5zM6.28 6.22a.75.75 0 0 0-1.06 1.06L6.44 8.5L5.22 9.72a.75.75 0 1 0 1.06 1.06l1.75-1.75l.53-.53l-.53-.53z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocCompressed(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 14.5V6H9.75A1.75 1.75 0 0 1 8 4.25V1.5H3.5v13H5V13h2v1.5zM7 11h2v2H7zm0 0V9h2V7H7V5H5v2h2v2H5v2zm2.5-9.379L12.379 4.5H9.75a.25.25 0 0 1-.25-.25zM3 0a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V4.414a1 1 0 0 0-.293-.707L10.293.293A1 1 0 0 0 9.586 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocExpand(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 1.5v12h5.75a.75.75 0 0 1 0 1.5H3a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1h6.644a1 1 0 0 1 .72.305l3.355 3.476a1 1 0 0 1 .281.695v.774a.75.75 0 0 1-.813.747a.76.76 0 0 1-.062.003H9.75A1.75 1.75 0 0 1 8 4.25V1.5zm6 .07l2.828 2.93H9.75a.25.25 0 0 1-.25-.25zm6.353 12.228L13.5 16l-2.353-2.202c-.314-.295-.091-.798.353-.798h1.25v-3H11.5c-.444 0-.667-.503-.353-.798L13.5 7l2.353 2.202c.314.295.091.798-.353.798h-1.25v3h1.25c.444 0 .667.503.353.798" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocImage(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 6v8.5h-9v-13H8v2.75C8 5.216 8.784 6 9.75 6zm-.121-1.5L9.5 1.621V4.25c0 .138.112.25.25.25zM2 1a1 1 0 0 1 1-1h6.586a1 1 0 0 1 .707.293l3.414 3.414a1 1 0 0 1 .293.707V15a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1zm7.125 9L11 12.143V13H5v-.9l1.125-1.35l1.143 1.372zM8 8.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocNew(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 1.5v13h5.75a.75.75 0 0 1 0 1.5H3a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1h6.644a1 1 0 0 1 .72.305l3.355 3.476a1 1 0 0 1 .281.695V6.25a.75.75 0 0 1-1.5 0V6H9.75A1.75 1.75 0 0 1 8 4.25V1.5zm6 .07l2.828 2.93H9.75a.25.25 0 0 1-.25-.25zM13 15a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 13 15" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocSymlink(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 0a1 1 0 0 0-1 1v6a.75.75 0 0 0 1.5 0V1.5H8v2.75C8 5.216 8.784 6 9.75 6h2.75v8.5H4.75a.75.75 0 0 0 0 1.5H13a1 1 0 0 0 1-1V4.476a1 1 0 0 0-.28-.695L10.362.305A1 1 0 0 0 9.643 0zm9.328 4.5L9.5 1.57v2.68c0 .138.112.25.25.25zM1 13.5a3.75 3.75 0 0 1 3.75-3.75H7V8.5c0-.445.503-.667.798-.353L10 10.5l-2.202 2.353c-.295.314-.798.091-.798-.353v-1.25H4.75A2.25 2.25 0 0 0 2.5 13.5v.75a.75.75 0 0 1-1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocText(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 6v8.5h-9v-13H8v2.75C8 5.216 8.784 6 9.75 6zm-.121-1.5L9.5 1.621V4.25c0 .138.112.25.25.25zM2 1a1 1 0 0 1 1-1h6.586a1 1 0 0 1 .707.293l3.414 3.414a1 1 0 0 1 .293.707V15a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1zm3.75 7a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5zM5 11.25a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocVersions(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 7v5.5h-7v-9H11V6a1 1 0 0 0 1 1zm-2-3.379L14.379 5.5H12.5zM7 2a1 1 0 0 0-1 1H4a1 1 0 0 0-1 1H1a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2a1 1 0 0 0 1 1h2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V5.414a1 1 0 0 0-.293-.707l-2.414-2.414A1 1 0 0 0 12.586 2zm-1 9.5v-7H4.5v7zm-3-1v-5H1.5v5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 6v8.5h-9v-13H8v2.75C8 5.216 8.784 6 9.75 6zm-.121-1.5L9.5 1.621V4.25c0 .138.112.25.25.25zM2 1a1 1 0 0 1 1-1h6.586a1 1 0 0 1 .707.293l3.414 3.414a1 1 0 0 1 .293.707V15a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Documents(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 6v5.5h-8v-10H9v2.75C9 5.216 9.784 6 10.75 6zm-3-4.379L13.379 4.5H10.75a.25.25 0 0 1-.25-.25zM5 0a1 1 0 0 0-1 1v2H2.25C1.56 3 1 3.56 1 4.25v10.5c0 .69.56 1.25 1.25 1.25h8.25c.69 0 1.25-.56 1.25-1.25v-1.67a.756.756 0 0 0-.004-.08H14a1 1 0 0 0 1-1V4.414a1 1 0 0 0-.293-.707L11.293.293A1 1 0 0 0 10.586 0zM4 12V4.5H2.5v10h7.75v-1.42c0-.027.001-.054.004-.08H5a1 1 0 0 1-1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotGrid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0m11.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0m-1.25 7a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5m1.25 4.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0M2.25 9.25a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5m7-1.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0M8 15a1.25 1.25 0 1 0 0-2.5A1.25 1.25 0 0 0 8 15M9.25 2.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0M2.25 15a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 7.159a.75.75 0 0 0-1.06 0l-1.97 1.97V1.75a.75.75 0 0 0-1.5 0v7.379l-1.97-1.97a.75.75 0 0 0-1.06 1.06l3.25 3.25L8 12l.53-.53l3.25-3.25a.75.75 0 0 0 0-1.061M2.5 9.75a.75.75 0 0 0-1.5 0V13a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drag(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.78 3.78a.75.75 0 0 0 0-1.06L8.53.47L8-.06l-.53.53l-2.25 2.25a.75.75 0 0 0 1.06 1.06L8 2.06l1.72 1.72a.75.75 0 0 0 1.06 0ZM3.84 5.22a.75.75 0 0 0-1.06 0L.53 7.47L0 8l.53.53l2.25 2.25a.75.75 0 1 0 1.06-1.06L2.12 8l1.72-1.72a.75.75 0 0 0 0-1.06Zm6.94 6.88a.75.75 0 0 1 0 1.06l-2.25 2.25l-.53.53l-.53-.53l-2.25-2.25a.75.75 0 1 1 1.06-1.06L8 13.82l1.72-1.72a.75.75 0 0 1 1.06 0Zm2.44-6.88a.75.75 0 0 0-1.06 1.06L13.88 8l-1.72 1.72a.75.75 0 1 0 1.06 1.06l2.25-2.25L16 8l-.53-.53l-2.25-2.25ZM8 6a.75.75 0 0 1 .75.75v.5h.5a.75.75 0 0 1 0 1.5h-.5v.5a.75.75 0 1 1-1.5 0v-.5h-.5a.75.75 0 1 1 0-1.5h.5v-.5A.75.75 0 0 1 8 6Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragHorizontal(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 4.25a.75.75 0 0 0 1.5 0v-2.5a.75.75 0 0 0-1.5 0v2.5Zm4.91.97a.75.75 0 0 1 1.06 0l2.25 2.25L16 8l-.53.53l-2.25 2.25a.75.75 0 0 1-1.06-1.06L13.88 8l-1.72-1.72a.75.75 0 0 1 0-1.06ZM8 10a.75.75 0 0 1-.75-.75v-2.5a.75.75 0 0 1 1.5 0v2.5A.75.75 0 0 1 8 10Zm0 5a.75.75 0 0 1-.75-.75v-2.5a.75.75 0 0 1 1.5 0v2.5A.75.75 0 0 1 8 15ZM2.78 5.22a.75.75 0 0 1 1.06 1.06L2.12 8l1.72 1.72a.75.75 0 1 1-1.06 1.06L.53 8.53L0 8l.53-.53l2.25-2.25Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragVertical(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.75 7.25a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5Zm-.97 4.85a.75.75 0 0 1 0 1.06l-2.25 2.25l-.53.53l-.53-.53l-2.25-2.25a.75.75 0 1 1 1.06-1.06L8 13.82l1.72-1.72a.75.75 0 0 1 1.06 0ZM6 8a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 6 8ZM1 8a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 1 8Zm9.78-5.28a.75.75 0 1 1-1.06 1.06L8 2.06L6.28 3.78a.75.75 0 0 1-1.06-1.06L7.47.47L8-.06l.53.53l2.25 2.25Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dumbbell(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 4.5v7h1v-7zM3.5 3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V8.75h3V12a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3.25h-3V4a1 1 0 0 0-1-1zM11 4.5v7h1v-7zm4.25.75A.75.75 0 0 1 16 6v4.5a.75.75 0 0 1-1.5 0V6a.75.75 0 0 1 .75-.75M1.5 6A.75.75 0 0 0 0 6v4.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Duplicate(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 9.5H2.5v-7h7V4H11V2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h2zm9.5 4h-7v-7h7zM5 6a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Earth(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.5c.23 0 .843-.226 1.487-1.514c.306-.612.563-1.37.742-2.236H5.771c.179.866.436 1.624.742 2.236C7.157 14.274 7.77 14.5 8 14.5M5.554 9.25a14.444 14.444 0 0 1 0-2.5h4.892a14.452 14.452 0 0 1 0 2.5zm6.203 1.5c-.224 1.224-.593 2.308-1.066 3.168a6.525 6.525 0 0 0 3.2-3.168zm2.623-1.5h-2.43a16.019 16.019 0 0 0 0-2.5h2.429a6.533 6.533 0 0 1 0 2.5Zm-10.331 0H1.62a6.533 6.533 0 0 1 0-2.5h2.43a15.994 15.994 0 0 0 0 2.5Zm-1.94 1.5h2.134c.224 1.224.593 2.308 1.066 3.168a6.525 6.525 0 0 1-3.2-3.168m3.662-5.5h4.458c-.179-.866-.436-1.624-.742-2.236C8.843 1.726 8.23 1.5 8 1.5c-.23 0-.843.226-1.487 1.514c-.306.612-.563 1.37-.742 2.236m5.986 0h2.134a6.526 6.526 0 0 0-3.2-3.168c.473.86.842 1.944 1.066 3.168M5.31 2.082c-.473.86-.842 1.944-1.066 3.168H2.109a6.525 6.525 0 0 1 3.2-3.168ZM8 0a8 8 0 1 1 0 16A8 8 0 0 1 8 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisH(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0m6 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0m4 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisV(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0-6a2 2 0 1 1 0 4a2 2 0 0 1 0-4m2-4a2 2 0 1 0-4 0a2 2 0 0 0 4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EntityBlocked(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.5H4A1.5 1.5 0 0 0 2.5 4v8A1.5 1.5 0 0 0 4 13.5h8a1.5 1.5 0 0 0 1.5-1.5V4A1.5 1.5 0 0 0 12 2.5M4 1a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3zm6.217 8.156a2.5 2.5 0 0 0-3.373-3.373zm-1.06 1.061a2.5 2.5 0 0 1-3.373-3.373l3.372 3.373ZM8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Environment(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.712 5.997l3.291-3.29l3.29 3.29l-3.29 3.29l-3.29-3.29Zm2.584-4.705a1 1 0 0 1 1.414 0l3.998 3.998a1 1 0 0 1 0 1.414l-3.998 3.998a1 1 0 0 1-1.414 0L7.827 9.233l-1.242 1.243a3 3 0 1 1-1.06-1.06l1.242-1.243l-1.469-1.469a1 1 0 0 1 0-1.414zM4 13.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Epic(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.5 8.5l2.5-6h8l2.5 6zM0 8.569v2.181c0 .966.784 1.75 1.75 1.75H2v.75c0 .966.784 1.75 1.75 1.75h8.5a.75.75 0 0 0 0-1.5h-8.5a.25.25 0 0 1-.25-.25v-.75h10.25a.75.75 0 0 0 0-1.5h-12a.25.25 0 0 1-.25-.25V10h13a1.5 1.5 0 0 0 1.385-2.077l-2.629-6.308A1 1 0 0 0 12.333 1H3.667a1 1 0 0 0-.923.615L.115 7.923A1.498 1.498 0 0 0 0 8.569" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EpicClosed(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4 2.5l-2.5 6h13l-2.5-6zm-4 8.25V8.569c-.01-.212.026-.431.115-.646l2.629-6.308A1 1 0 0 1 3.667 1h8.666a1 1 0 0 1 .923.615l2.629 6.308A1.5 1.5 0 0 1 14.5 10h-13v.75c0 .138.112.25.25.25h12a.75.75 0 0 1 0 1.5H3.5v.75c0 .138.112.25.25.25h8.5a.75.75 0 0 1 0 1.5h-8.5A1.75 1.75 0 0 1 2 13.25v-.75h-.25A1.75 1.75 0 0 1 0 10.75m10.28-5.72a.75.75 0 1 0-1.06-1.06L7.5 5.69l-.72-.72a.75.75 0 1 0-1.06 1.06l1.25 1.25a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorIcon(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16m1-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25-6.25a.75.75 0 0 0-1.5 0v3.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 11.16a.75.75 0 0 0-1.06 0l-1.97 1.97V2.87l1.97 1.97a.75.75 0 1 0 1.06-1.06L8.53.53L8 0l-.53.53l-3.25 3.25a.75.75 0 0 0 1.06 1.06l1.97-1.97v10.26l-1.97-1.97a.75.75 0 0 0-1.06 1.06l3.25 3.25L8 16l.53-.53l3.25-3.25a.75.75 0 0 0 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandDown(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 1a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zm8.03 9.159a.75.75 0 0 0-1.06 0l-1.97 1.97V4.75a.75.75 0 0 0-1.5 0v7.379l-1.97-1.97a.75.75 0 0 0-1.06 1.06l3.25 3.25L8 15l.53-.53l3.25-3.25a.75.75 0 0 0 0-1.061" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandLeft(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 12.25a.75.75 0 0 0 1.5 0v-8.5a.75.75 0 0 0-1.5 0zm11.159-8.03a.75.75 0 0 0 0 1.06l1.97 1.97H3.75a.75.75 0 1 0 0 1.5h9.379l-1.97 1.97a.75.75 0 1 0 1.06 1.06l3.25-3.25L16 8l-.53-.53l-3.25-3.25a.75.75 0 0 0-1.061 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandRight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 12.25a.75.75 0 0 0 1.5 0v-8.5a.75.75 0 0 0-1.5 0zM4.841 4.22a.75.75 0 0 1 0 1.06l-1.97 1.97h9.379a.75.75 0 0 1 0 1.5H2.871l1.97 1.97a.75.75 0 1 1-1.06 1.06L.53 8.53L0 8l.53-.53l3.25-3.25a.75.75 0 0 1 1.061 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandUp(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 5.841a.75.75 0 0 1-1.06 0l-1.97-1.97v7.379a.75.75 0 0 1-1.5 0V3.871l-1.97 1.97a.75.75 0 0 1-1.06-1.06l3.25-3.25L8 1l.53.53l3.25 3.25a.75.75 0 0 1 0 1.061M3.75 13.5a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expire(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.175.002a8 8 0 1 0 2.309 15.603a.75.75 0 0 0-.466-1.426a6.5 6.5 0 1 1 3.996-8.646a.75.75 0 0 0 1.388-.569A8 8 0 0 0 8.175.002M8.75 3.75a.75.75 0 0 0-1.5 0v3.94L5.216 9.723a.75.75 0 1 0 1.06 1.06L8.53 8.53l.22-.22zM15 15a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25-6.25a.75.75 0 0 0-1.5 0v3.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 5.841a.75.75 0 0 1-1.06 0l-1.97-1.97v7.379a.75.75 0 0 1-1.5 0V3.871l-1.97 1.97a.75.75 0 0 1-1.06-1.06l3.25-3.25L8 1l.53.53l3.25 3.25a.75.75 0 0 1 0 1.061M2.5 9.75a.75.75 0 0 0-1.5 0V13a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.75 1a.75.75 0 0 0 0 1.5h1.69L8.22 6.72a.75.75 0 0 0 1.06 1.06l4.22-4.22v1.69a.75.75 0 0 0 1.5 0V1zM2.5 4v9a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V8.75a.75.75 0 0 1 1.5 0V13a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4.25a.75.75 0 0 1 0 1.5H3a.5.5 0 0 0-.5.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 8s3-6 8-6s8 6 8 6s-3 6-8 6s-8-6-8-6m1.81.13A13.593 13.593 0 0 1 1.73 8l.082-.13c.326-.51.806-1.187 1.42-1.856C4.494 4.635 6.12 3.5 8 3.5c1.878 0 3.506 1.135 4.77 2.514A13.705 13.705 0 0 1 14.27 8a14.021 14.021 0 0 1-1.502 1.986C11.506 11.365 9.88 12.5 8 12.5c-1.878 0-3.506-1.135-4.77-2.514A13.703 13.703 0 0 1 1.81 8.13M11 8a3 3 0 1 1-2.117-2.868a1.5 1.5 0 1 0 1.985 1.985A3 3 0 0 1 11 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlash(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.776 2.284a.75.75 0 0 0-1.06-1.06l-1.953 1.953C10.707 2.499 9.44 2 8 2C5.506 2 3.533 3.49 2.24 4.86A14.615 14.615 0 0 0 .226 7.576a5.39 5.39 0 0 0-.028.052l-.008.014l-.003.005v.002L.85 8l-.663-.35L.002 8l.185.35L.85 8l-.663.35v.002l.002.002l.004.007l.012.023c.01.02.026.047.046.082a14.417 14.417 0 0 0 .82 1.262c.47.647 1.13 1.447 1.96 2.18L1.22 13.72a.75.75 0 1 0 1.06 1.06zm-10.681 8.56l1.32-1.32a3 3 0 0 1 4.109-4.109l1.148-1.147C9.864 3.8 8.969 3.5 8 3.5c-1.88 0-3.483 1.134-4.67 2.39A13.114 13.114 0 0 0 1.716 8c.13.213.32.508.567.846a11.98 11.98 0 0 0 1.811 1.998Zm9.42-5.166a.75.75 0 0 1 1.053.122c.447.564.901 1.2 1.245 1.85l.185.35l-.185.35L15.15 8l.663.351l-.001.002l-.003.005l-.008.014a9.81 9.81 0 0 1-.53.865a14.62 14.62 0 0 1-1.51 1.903C12.467 12.51 10.494 14 8 14a6.939 6.939 0 0 1-1.021-.08l-.02-.002l-.006-.001h-.002l.12-.741l-.12.74a.75.75 0 0 1 .239-1.48h.002l.011.001a6.024 6.024 0 0 0 .235.03c.158.017.362.033.562.033c1.88 0 3.483-1.134 4.67-2.39a13.11 13.11 0 0 0 1.616-2.115a12.33 12.33 0 0 0-.893-1.263a.75.75 0 0 1 .121-1.054Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceNeutral(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16m-2.5-5.25a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75M7 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceUnhappy(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16m0-7a3 3 0 0 0-3 3h6a3 3 0 0 0-3-3M6.002 5a1.5 1.5 0 0 0-1.3.747a.5.5 0 1 0 .866.502a.5.5 0 0 1 .865 0a.5.5 0 0 0 .866-.5A1.5 1.5 0 0 0 6.002 5m3.25.2a1.5 1.5 0 0 1 2.047.55a.5.5 0 1 1-.866.5a.5.5 0 0 0-.865-.001a.5.5 0 1 1-.865-.502a1.5 1.5 0 0 1 .55-.547Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FalsePositive(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.229.199a8 8 0 0 1 9.727 6.964a.75.75 0 0 1-1.492.157a6.5 6.5 0 1 0-7.132 7.146a.75.75 0 1 1-.154 1.492a8 8 0 0 1-.95-15.76Zm5.051 10.02a.75.75 0 1 0-1.06 1.061L11.94 13l-1.72 1.72a.75.75 0 1 0 1.06 1.06L13 14.06l1.72 1.72a.75.75 0 1 0 1.06-1.06L14.06 13l1.72-1.72a.75.75 0 1 0-1.06-1.06L13 11.94zm.25-3.939a.75.75 0 0 0-1.06-1.06L6.5 9.19L5.28 7.97a.75.75 0 0 0-1.06 1.06l1.75 1.75a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FeatureFlag(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 5v6a3.5 3.5 0 1 1-7 0V5a3.5 3.5 0 1 1 7 0M3 5a5 5 0 0 1 10 0v6a5 5 0 0 1-10 0zm5 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FeatureFlagDisabled(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 5v6a3.5 3.5 0 1 1-7 0V5a3.5 3.5 0 1 1 7 0M3 5a5 5 0 0 1 10 0v6a5 5 0 0 1-10 0zm5 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAddition(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H4a.5.5 0 0 1-.5-.5V4a.5.5 0 0 1 .5-.5M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm6 7a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 8 11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAdditionSolid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm4 9a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 8 11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDeletion(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H4a.5.5 0 0 1-.5-.5V4a.5.5 0 0 1 .5-.5M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm3.75 3.25a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDeletionSolid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm1.75 5.25a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileModified(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H4a.5.5 0 0 1-.5-.5V4a.5.5 0 0 1 .5-.5M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm6 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileModifiedSolid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm4 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTree(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 1.75a.75.75 0 0 0-1.5 0v8.5a3 3 0 0 0 3 3h3V14a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v.75H4a1.5 1.5 0 0 1-1.5-1.5V6.849c.441.255.954.401 1.5.401h3V8a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v.75H4a1.5 1.5 0 0 1-1.5-1.5zm11 11.75h-5v-2h5zm-5-6v-2h5v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.5 8.379l.44-.44l4.56-4.56V2.5h-11v.879l4.56 4.56l.44.44v4l1-1zM10 12l-2.5 2.5L6 16V9L1.293 4.293A1 1 0 0 1 1 3.586V2a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v1.586a1 1 0 0 1-.293.707L10 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.728 6.475c.052-.265.07-.534.054-.806a3.645 3.645 0 0 0-.328-1.272c-.425-.947-1.263-1.941-2.558-3.005c-.146-.12-.298-.241-.455-.363A25.558 25.558 0 0 0 6 0a7.302 7.302 0 0 1-.336 2.2c-.419 1.346-1.167 2.472-1.878 3.543C2.862 7.133 2 8.43 2 10a6 6 0 0 0 12 0c0-.593-.093-1.175-.26-1.741a7.39 7.39 0 0 0-.226-.654a9.204 9.204 0 0 0-.732-1.436a10.736 10.736 0 0 0-.328-.5c0 .261-.037.518-.105.769c-.05.182-.116.36-.195.535c-.09.195-.196.385-.317.567a5.03 5.03 0 0 1-.335.45a5.856 5.856 0 0 1-.686.696a6.588 6.588 0 0 1-1.196.821a2.239 2.239 0 0 1-.01.005a5.878 5.878 0 0 1-.177.09c-.432.21-.63-.236-.285-.57l.006-.006a7.474 7.474 0 0 0 .66-.719a5.82 5.82 0 0 0 .298-.405l.011-.018c.158-.238.288-.477.39-.718a3.5 3.5 0 0 0 .215-.69ZM8.11 7.948c.325-.315.572-.608.753-.88c.341-.51.444-.93.42-1.307a2.152 2.152 0 0 0-.199-.75c-.28-.623-.882-1.402-1.997-2.34c-.484 1.542-1.328 2.811-2.013 3.842l-.04.06C4.056 8.046 3.5 8.963 3.5 10a4.5 4.5 0 1 0 9 0c0-.262-.025-.525-.073-.79a8.026 8.026 0 0 1-2.34 1.741l-.654-1.35l.655 1.35a1.91 1.91 0 0 1-1.274.157a1.73 1.73 0 0 1-1.169-.94c-.425-.89-.014-1.756.46-2.214l.006-.006Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstContribution(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.716.315a1 1 0 0 0-1.432 0L6.646.97a1 1 0 0 1-.988.265l-.88-.249a1 1 0 0 0-1.24.717l-.226.886a1 1 0 0 1-.723.723l-.886.225a1 1 0 0 0-.717 1.24l.249.881a1 1 0 0 1-.265.987l-.655.639a1 1 0 0 0 0 1.432l.655.639a1 1 0 0 1 .265.987l-.249.88a1 1 0 0 0 .717 1.24l.886.226a1 1 0 0 1 .723.723l.225.886a1 1 0 0 0 1.24.716l.881-.248a1 1 0 0 1 .988.265l.638.655a1 1 0 0 0 1.432 0l.639-.655a1 1 0 0 1 .987-.265l.88.248a1 1 0 0 0 1.24-.716l.226-.886a1 1 0 0 1 .723-.723l.886-.225a1 1 0 0 0 .716-1.24l-.248-.881a1 1 0 0 1 .265-.988l.655-.638a1 1 0 0 0 0-1.432l-.655-.638a1 1 0 0 1-.265-.988l.248-.88a1 1 0 0 0-.716-1.24l-.886-.226a1 1 0 0 1-.723-.723l-.225-.886a1 1 0 0 0-1.24-.717l-.881.249A1 1 0 0 1 9.354.97zM7 4h2v8H7V6H6V5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 10.174v.012c-1.393.72-2.81.247-4.959-.585l-.17-.066C6.763 8.91 4.685 8.105 2.5 8.63V2.814c1.393-.72 2.81-.247 4.959.585l.17.066c1.607.624 3.685 1.43 5.871.904zM8 11c-1.83-.708-3.659-1.416-5.5-.81v4.06a.75.75 0 0 1-1.5 0V2C3.35.2 5.675 1.1 8 2c1.83.708 3.659 1.416 5.5.81A5.068 5.068 0 0 0 15 2v9c-2.35 1.8-4.675.9-7 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4V2.5h4.697l1 1.5zM0 4V2a1 1 0 0 1 1-1h5.465a1 1 0 0 1 .832.445l1.667 2.5l.034.055H15a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1zm1.5 1.5v7h13v-7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderNew(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2.5V4h5.697l-1-1.5zM0 2v10.75C0 13.44.56 14 1.25 14h6a.75.75 0 0 0 0-1.5H1.5v-7h13v1.75a.75.75 0 0 0 1.5 0v-2C16 4.56 15.44 4 14.75 4H8.998a1.054 1.054 0 0 0-.034-.055l-1.667-2.5A1 1 0 0 0 6.465 1H1a1 1 0 0 0-1 1m13 13a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 13 15" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderO(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4V2.5h4.697l1 1.5zM0 4V2a1 1 0 0 1 1-1h5.465a1 1 0 0 1 .832.445l1.667 2.5l.034.055H15a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1zm1.5 1.5v7h13v-7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 7.33V2.5h4.618l1.453 2h-4.54a1 1 0 0 0-.965.739L1.5 7.329ZM.005 13.087A.758.758 0 0 1 0 13V2.12C0 1.501.501 1 1.12 1h5.191c.359 0 .696.172.907.462L9.062 4h3.954c.409 0 .772.196 1 .5h.95a1 1 0 0 1 .966 1.261l-2.029 7.5a1 1 0 0 1-.965.739H1.002a1 1 0 0 1-.997-.913m1.65-.587L3.414 6h10.9l-1.759 6.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Food(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.75 1A5.75 5.75 0 0 0 1 6.75v.518a2 2 0 0 0 0 3.464v1.518A2.75 2.75 0 0 0 3.75 15h8.5A2.75 2.75 0 0 0 15 12.25v-1.518a2 2 0 0 0 0-3.464V6.75A5.75 5.75 0 0 0 9.25 1zM14 8.5H2a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1M13.5 7v-.25A4.25 4.25 0 0 0 9.25 2.5h-2.5A4.25 4.25 0 0 0 2.5 6.75V7zM11 11h2.5v1.25c0 .69-.56 1.25-1.25 1.25h-8.5c-.69 0-1.25-.56-1.25-1.25V11H9l1 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fork(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25 2.386a2.501 2.501 0 1 0-1.5 0v.364a2.5 2.5 0 0 0 2.5 2.5a1 1 0 0 1 1 1v.364a2.501 2.501 0 1 0 1.5 0V9.75a1 1 0 0 1 1-1a2.5 2.5 0 0 0 2.5-2.5v-.364a2.501 2.501 0 1 0-1.5 0v.364a1 1 0 0 1-1 1c-.681 0-1.3.273-1.75.715a2.492 2.492 0 0 0-1.75-.715a1 1 0 0 1-1-1zM11.5 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-3.5 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Formula(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.581 4.49A2.75 2.75 0 0 1 8.319 2h.931a.75.75 0 0 1 0 1.5h-.931a1.25 1.25 0 0 0-1.245 1.131l-.083.869H9.25a.75.75 0 0 1 0 1.5H6.849l-.43 4.51A2.75 2.75 0 0 1 3.681 14H2.75a.75.75 0 0 1 0-1.5h.931a1.25 1.25 0 0 0 1.245-1.132L5.342 7H3.75a.75.75 0 0 1 0-1.5h1.735zM9.22 9.22a.75.75 0 0 1 1.06 0l1.22 1.22l1.22-1.22a.75.75 0 1 1 1.06 1.06l-1.22 1.22l1.22 1.22a.75.75 0 1 1-1.06 1.06l-1.22-1.22l-1.22 1.22a.75.75 0 1 1-1.06-1.06l1.22-1.22l-1.22-1.22a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Git(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.698 7.287L8.712.302a1.03 1.03 0 0 0-1.457 0l-1.45 1.45l1.84 1.84a1.223 1.223 0 0 1 1.55 1.56l1.773 1.774a1.224 1.224 0 0 1 1.267 2.025a1.226 1.226 0 0 1-2.001-1.334L8.579 5.963v4.353a1.226 1.226 0 1 1-1.008-.036V5.887a1.226 1.226 0 0 1-.666-1.608L5.093 2.465l-4.79 4.79a1.031 1.031 0 0 0 0 1.457l6.986 6.986a1.03 1.03 0 0 0 1.457 0l6.953-6.953a1.03 1.03 0 0 0 0-1.458"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitMerge(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.34 1.22a.75.75 0 0 0-1.06 0L7.53 2.97L7 3.5l.53.53l1.75 1.75a.75.75 0 1 0 1.06-1.06l-.47-.47h.63c.69 0 1.25.56 1.25 1.25v4.614a2.501 2.501 0 1 0 1.5 0V5.5a2.75 2.75 0 0 0-2.75-2.75h-.63l.47-.47a.75.75 0 0 0 0-1.06M13.5 12.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-9 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1.5 0a2.5 2.5 0 1 1-3.25-2.386V5.886a2.501 2.501 0 1 1 1.5 0v4.228A2.501 2.501 0 0 1 6 12.5m-1.5-9a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gitea(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.46 3.206s.14-.003.245.102l.01.01c.064.06.258.244.285 1.074c0 2.902-1.405 5.882-1.405 5.882a8.486 8.486 0 0 1-.359.713c-.458.802-.786 1.153-.786 1.153s-.318.379-.675.595c-.415.265-.72.263-.72.263L7.247 13c-.636-.079-1.29-.736-1.927-1.578c-.47-.677-.779-1.413-.779-1.413s-2.51.034-3.675-1.394C.235 7.895.103 7.067.06 6.769c0-.008-.002-.018-.004-.029c-.05-.324-.285-1.873.821-2.86c.517-.496 1.148-.638 1.37-.684c.371-.081.667-.06.903-.044l.09.006c.391.035 3.99.216 3.99.216s1.532.066 2.27.056c0 0 .003 1.853.003 2.78c.07.032.14.066.211.1l.212.1V3.427c.33-.002.661-.01.996-.017h.011c1.545-.036 4.528-.204 4.528-.204ZM2.113 8.026s.28.26.94.477c.43.152 1.094.231 1.094.231S3.699 7.5 3.516 6.757c-.22-.886-.4-2.398-.4-2.398s-.438-.015-.789.079c-.766.19-.98.763-.98.763s-.384.688.036 1.813c.244.672.73 1.013.73 1.013Zm8.084 3.607c.344-.023.499-.392.499-.392s1.24-2.486 1.4-2.878a.7.7 0 0 0 .046-.438c-.07-.267-.39-.412-.39-.412l-1.926-.935l-.165.339l-.18.369a.458.458 0 0 1 .128.341s.433.186.743.387c0 0 .257.135.32.425c.075.273-.04.488-.066.539l-.002.003s-.216.51-.343.774l-.004.007c-.047.097-.092.189-.139.28a.454.454 0 1 1-.32-.15s.41-.84.468-1.033c0 0 .096-.24.048-.38a.474.474 0 0 0-.19-.188a5.905 5.905 0 0 0-.678-.34s-.076.068-.18.09a.454.454 0 0 1-.158.014l-.611 1.25a.459.459 0 0 1 .046.587a.457.457 0 0 1-.578.138a.458.458 0 0 1-.232-.51a.46.46 0 0 1 .44-.35L8.8 7.886a.457.457 0 0 1 .361-.744l.185-.375l.167-.341l-.579-.281s-.251-.125-.458-.072a.568.568 0 0 0-.114.039c-.189.084-.31.33-.31.33L6.668 9.293s-.124.254-.068.46c.048.252.325.397.325.397l2.874 1.4l.135.054s.114.04.262.03Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.976 0A7.977 7.977 0 0 0 0 7.976c0 3.522 2.3 6.507 5.431 7.584c.392.049.538-.196.538-.392v-1.37c-2.201.49-2.69-1.076-2.69-1.076c-.343-.93-.881-1.175-.881-1.175c-.734-.489.048-.489.048-.489c.783.049 1.224.832 1.224.832c.734 1.223 1.859.88 2.3.685c.048-.538.293-.88.489-1.076c-1.762-.196-3.621-.881-3.621-3.964c0-.88.293-1.566.832-2.153c-.05-.147-.343-.978.098-2.055c0 0 .685-.196 2.201.832c.636-.196 1.322-.245 2.007-.245s1.37.098 2.006.245c1.517-1.027 2.202-.832 2.202-.832c.44 1.077.146 1.908.097 2.104a3.16 3.16 0 0 1 .832 2.153c0 3.083-1.86 3.719-3.62 3.915c.293.244.538.733.538 1.467v2.202c0 .196.146.44.538.392A7.984 7.984 0 0 0 16 7.976C15.951 3.572 12.38 0 7.976 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoBack(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2.87 7.75l1.97 1.97a.75.75 0 1 1-1.06 1.06L.53 7.53L0 7l.53-.53l3.25-3.25a.75.75 0 0 1 1.06 1.06L2.87 6.25h9.88a3.25 3.25 0 0 1 0 6.5h-2a.75.75 0 0 1 0-1.5h2a1.75 1.75 0 1 0 0-3.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 7v2.4h3.97c-.16 1.03-1.2 3.02-3.97 3.02c-2.39 0-4.34-1.98-4.34-4.42S5.61 3.58 8 3.58c1.36 0 2.27.58 2.79 1.08l1.9-1.83C11.47 1.69 9.89 1 8 1C4.13 1 1 4.13 1 8s3.13 7 7 7c4.04 0 6.72-2.84 6.72-6.84c0-.46-.05-.81-.11-1.16z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grip(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 3.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0m0 4.75A1.25 1.25 0 1 1 9 8a1.25 1.25 0 0 1 2.5 0m-5.75 6a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5m4.5 0a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5m-4.5-4.75a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5m0-4.75a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m1.5 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0m4 5.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0 1.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-7 2.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0 1.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hamburger(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 3.75A.75.75 0 0 1 .75 3h14.5a.75.75 0 0 1 0 1.5H.75A.75.75 0 0 1 0 3.75M0 8a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H.75A.75.75 0 0 1 0 8m.75 3.5a.75.75 0 0 0 0 1.5h14.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heading(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 2.75a.75.75 0 0 0-1.5 0v10.5a.75.75 0 0 0 1.5 0v-4.5h7v4.5a.75.75 0 0 0 1.5 0V2.75a.75.75 0 0 0-1.5 0v4.5h-7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.753 2.247L8 3l-.753-.753A4.243 4.243 0 0 0 1.25 8.25l5.69 5.69L8 15l1.06-1.06l5.69-5.69a4.243 4.243 0 0 0-5.997-6.003M8 12.879l5.69-5.69a2.743 2.743 0 0 0-3.877-3.881l-.752.753L8 5.12L6.94 4.06l-.753-.752v-.001A2.743 2.743 0 0 0 2.31 7.189L8 12.88Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Highlight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 14.25a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5H2.75a.75.75 0 0 1-.75-.75m1.5-3.252V9.112l1.29-1.258L6.6 9.688l-1.268 1.31zM7.645 8.61l4.715-4.866a.5.5 0 0 0-.028-.722l-1.01-.895l.995-1.123l1.01.895a2 2 0 0 1 .11 2.889l-7.47 7.71H2V8.48l7.594-7.41a2 2 0 0 1 2.723-.066l-.995 1.123a.5.5 0 0 0-.68.016L5.863 6.806z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.806.665A8 8 0 1 1 .612 11.07a.75.75 0 1 1 1.385-.575A6.5 6.5 0 1 0 2.523 4.5H4.25a.75.75 0 0 1 0 1.5H0V1.75a.75.75 0 0 1 1.5 0v1.586A8 8 0 0 1 4.806.666ZM8 3a.75.75 0 0 1 .75.75v3.94l2.034 2.034a.75.75 0 1 1-1.06 1.06L7.47 8.53l-.22-.22V3.75A.75.75 0 0 1 8 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.38 1.353L8 1.131l-.38.222l-7.25 4.25a.75.75 0 0 0 .76 1.294l.87-.51V14h12V6.387l.87.51a.75.75 0 1 0 .76-1.294zm4.12 4.154L8 2.87L3.5 5.507V12.5H6V8h4v4.5h2.5zM8.5 9.5v3h-1v-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hook(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 11.125a3.875 3.875 0 0 0 7 2.292a3.875 3.875 0 0 0 7-2.292V7.002l-1.28 1.28l-1.49 1.488a.75.75 0 0 0 1.061 1.061l.208-.208v.502a2.375 2.375 0 1 1-4.75 0v-5.24a2.501 2.501 0 1 0-1.5 0v5.24a2.375 2.375 0 1 1-4.75 0v-.502l.208.208a.75.75 0 1 0 1.06-1.06L2.28 8.281L1 7.002zM9 3.5a1 1 0 1 0-2 0a1 1 0 0 0 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 0a.75.75 0 0 0 0 1.5H3v.593c0 1.26.5 2.468 1.391 3.359L6.94 8l-2.548 2.548A4.75 4.75 0 0 0 3 13.907v.593h-.25a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5H13v-.593c0-1.26-.5-2.468-1.391-3.359L9.06 8l2.548-2.548A4.75 4.75 0 0 0 13 2.093V1.5h.25a.75.75 0 0 0 0-1.5zm8.75 1.5h-7v.593c0 .69.219 1.356.618 1.907h5.764a3.25 3.25 0 0 0 .618-1.907zM8 6.94L6.56 5.5h2.88zm3.5 7.56v-.593a3.25 3.25 0 0 0-.952-2.298L8 9.06l-2.548 2.548a3.25 3.25 0 0 0-.952 2.298v.593z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageCommentDark(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16m2-5H5l-1 1V7a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2m0-1.5H5.5V7a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageCommentLight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0m-6 3H5l-1 1V7a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2m0-1.5H5.5V7a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 7.159a.75.75 0 0 0-1.06 0l-1.97 1.97V1.75a.75.75 0 0 0-1.5 0v7.379l-1.97-1.97a.75.75 0 0 0-1.06 1.06l3.25 3.25L8 12l.53-.53l3.25-3.25a.75.75 0 0 0 0-1.061M2.5 9.75a.75.75 0 1 0-1.5 0V13a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Incognito(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.248 2.042a.75.75 0 0 0-.933.403L2.513 6.5H.75a.75.75 0 1 0 0 1.5h2.233a.76.76 0 0 0 .033 0H15.25a.75.75 0 0 0 0-1.5h-1.763l-1.802-4.055a.75.75 0 0 0-.933-.403L8 3.005zM11.846 6.5l-1.25-2.814l-2.348.822a.75.75 0 0 1-.496 0l-2.347-.822L4.155 6.5zm-6.346 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 1.5A2.5 2.5 0 0 0 8 11.5a2.5 2.5 0 0 0 5 0h.25a.75.75 0 0 0 0-1.5h-.75a2.496 2.496 0 0 0-2-1c-.818 0-1.544.393-2 1h-1a2.496 2.496 0 0 0-2-1c-.818 0-1.544.393-2 1h-.75a.75.75 0 0 0 0 1.5H3A2.5 2.5 0 0 0 5.5 14m5-1.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Information(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M9 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M7 7a.75.75 0 0 0 0 1.5h.25v2h-1a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-1V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationO(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0m-9.75 2.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-1V7H7a.75.75 0 0 0 0 1.5h.25v2zM8 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfrastructureRegistry(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8 2.732l-2.945 1.7L8 6.135l2.945-1.701zm4.445.834L8 1L3.555 3.566l-1.43-.825a.75.75 0 1 0-.75 1.298l1.429.826V10l4.446 2.567v1.683a.75.75 0 0 0 1.5 0v-1.683L13.196 10V4.865l1.43-.826a.75.75 0 0 0-.751-1.298zm-.749 2.165L8.75 7.433v3.402l2.946-1.701zM4.304 9.134l2.946 1.7v-3.4L4.304 5.73z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueBlock(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a2 2 0 0 0-2 2v3.25a.75.75 0 0 0 1.5 0V3a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 .5.5v10.234a.753.753 0 0 0 .57.744a.752.752 0 0 0 .851-.393l3.7-7.047a1.75 1.75 0 0 0-.74-2.366l-2.91-1.514A2 2 0 0 0 9 1zm8 3.364v5.844l2.293-4.367a.25.25 0 0 0-.106-.338zm-4.783 8.792a2.5 2.5 0 0 0-3.373-3.373zm-1.06 1.061a2.5 2.5 0 0 1-3.373-3.373l3.372 3.373ZM4 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueClose(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M3.75 7.25a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueClosed(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.5H3a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5V3a.5.5 0 0 0-.5-.5M3 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h6a2 2 0 0 0 1.651-.87a.753.753 0 0 0 .098-.145l3.878-7.45a1.75 1.75 0 0 0-.744-2.361l-2.912-1.516A2 2 0 0 0 9 1zm10.297 4.841L11 10.254v-5.89l2.19 1.14a.25.25 0 0 1 .107.337M8.28 7.281A.75.75 0 0 0 7.22 6.22L5.25 8.19l-.47-.47a.75.75 0 0 0-1.06 1.06l1 1a.75.75 0 0 0 1.06 0l2.5-2.5Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueNew(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3.25a.75.75 0 0 0 0-1.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 .5.5v5.25a.75.75 0 0 0 1.5 0V4.364l2.19 1.14a.25.25 0 0 1 .107.338l-1.072 2.062a.75.75 0 0 0 1.33.692l1.073-2.062a1.75 1.75 0 0 0-.745-2.36l-2.912-1.516A2 2 0 0 0 9 1zm6 12a.75.75 0 0 1 .75-.75h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5a.75.75 0 0 1-1.5 0v-1.5h-1.5A.75.75 0 0 1 9 13" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueOpenM(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeEnhancement(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.75.75a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0zm4.03 2.53a.75.75 0 1 0-1.06-1.06l-1.5 1.5a.75.75 0 1 0 1.06 1.06zM8.293 4.646a1 1 0 0 1 1.414 0l1.647 1.647a1 1 0 0 1 0 1.414l-7.99 7.99a1 1 0 0 1-1.415 0L.303 14.051a1 1 0 0 1 0-1.414l7.99-7.99ZM9 6.061L7.06 8l.94.94L9.94 7L9 6.06Zm-7.283 7.283L6 9.06l.94.939l-4.284 4.283l-.94-.94ZM13 7a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5A.75.75 0 0 1 13 7m.78 4.78a.75.75 0 0 0 0-1.06l-1.5-1.5a.75.75 0 0 0-1.06 1.06l1.5 1.5a.75.75 0 0 0 1.06 0m-9.56-8.5a.75.75 0 0 1 1.06-1.06l1.5 1.5a.75.75 0 0 1-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeFeature(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M8.688 3.202a.75.75 0 0 0-1.376 0L6.266 5.613l-2.617.25a.75.75 0 0 0-.425 1.309l1.97 1.74l-.571 2.565a.75.75 0 0 0 1.113.81L8 10.95l2.264 1.336a.75.75 0 0 0 1.113-.809l-.571-2.565l1.97-1.74a.75.75 0 0 0-.425-1.31l-2.617-.249zM7.466 6.616L8 5.386l.534 1.23a.75.75 0 0 0 .617.449l1.336.127l-1.006.888a.75.75 0 0 0-.236.726l.292 1.31l-1.156-.683a.75.75 0 0 0-.762 0l-1.156.682l.292-1.31a.75.75 0 0 0-.236-.725l-1.006-.888l1.336-.127a.75.75 0 0 0 .617-.449" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeFeatureFlag(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 10.174v.012c-1.393.72-2.81.247-4.959-.585l-.17-.066C6.763 8.91 4.685 8.105 2.5 8.63V2.814c1.393-.72 2.81-.247 4.959.585l.17.066c1.607.624 3.685 1.43 5.871.904zM8 11c-1.83-.708-3.659-1.416-5.5-.81v4.06a.75.75 0 0 1-1.5 0V2C3.35.2 5.675 1.1 8 2c1.83.708 3.659 1.416 5.5.81A5.068 5.068 0 0 0 15 2v9c-2.35 1.8-4.675.9-7 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeIncident(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 12.25a1.5 1.5 0 0 0 1.5 1.5h8a1.5 1.5 0 0 0 1.5-1.5v-8a1.5 1.5 0 0 0-1.5-1.5H4a1.5 1.5 0 0 0-1.5 1.5zm1.5 3a3 3 0 0 1-3-3v-8a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3zM9 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25-5.75a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeIssue(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5h6a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5M1 3a2 2 0 0 1 2-2h6a2 2 0 0 1 1.97 1.658l2.913 1.516a1.75 1.75 0 0 1 .744 2.36l-3.878 7.45a.753.753 0 0 1-.098.145c-.36.526-.965.871-1.651.871H3a2 2 0 0 1-2-2zm10 7.254l2.297-4.413a.25.25 0 0 0-.106-.337L11 4.364z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeKeyresult(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 5.25a.75.75 0 0 0 1.5 0V1h-4.25a.75.75 0 0 0 0 1.5h1.69L7.94 7H4.75a.75.75 0 0 0-.53.22l-3 3a.75.75 0 0 0 .53 1.28H4.5v2.75a.75.75 0 0 0 1.28.53l3-3a.75.75 0 0 0 .22-.53V8.06l4.5-4.5zM7.47 8.53a.756.756 0 0 1-.029-.03h-2.38L3.56 10H6v2.44l1.5-1.5V8.558a.756.756 0 0 1-.03-.029Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeMaintenance(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.25 2.5a2.25 2.25 0 0 0-2.154 2.904l.13.43l-.317.318l-6.254 6.253l-.53-.53l.53.53a.664.664 0 0 0 .94.94L9.848 7.09l.318-.318l.43.13a2.25 2.25 0 0 0 2.685-3.124l-1.5 1.501a.75.75 0 1 1-1.061-1.06l1.5-1.5a2.241 2.241 0 0 0-.97-.22ZM7.5 4.75a3.75 3.75 0 1 1 3.114 3.696L10.061 9l.939.94l.47-.47l.53-.53l.53.53l1.875 1.875a2.164 2.164 0 1 1-3.06 3.06L9.47 12.53L8.94 12l.53-.53l.47-.47l-.94-.94l-4.345 4.345l-.53-.53l.53.53a2.164 2.164 0 1 1-3.06-3.06L5.939 7L3.5 4.56l-.617.617l-.507-.761l-1-1.5l-.341-.512l.435-.434l.5-.5l.434-.435l.512.341l1.5 1l.761.507l-.616.617L7 5.94l.554-.554A3.769 3.769 0 0 1 7.5 4.75m4.5 6.31l1.345 1.345a.664.664 0 0 1-.94.94L11.061 12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeObjective(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.037.057A.75.75 0 0 1 13.5.75V2.5h1.75a.75.75 0 0 1 .53 1.28l-3 3a.75.75 0 0 1-.53.22h-.377a4 4 0 1 1-4.797-2.892a.75.75 0 0 1 .347 1.46A2.5 2.5 0 1 0 10.29 7h-.23L8.53 8.53a.75.75 0 1 1-1.06-1.06L9 5.94V3.75a.75.75 0 0 1 .22-.53l3-3a.75.75 0 0 1 .817-.163M10.5 4.061V5.5h1.44l1.5-1.5H12V2.56zM4.82 2.33a6.5 6.5 0 0 1 3.853-.796a.75.75 0 1 0 .155-1.492a8 8 0 1 0 7.129 7.128a.75.75 0 1 0-1.492.155A6.5 6.5 0 1 1 4.82 2.331Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeRequirements(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h5.25a.75.75 0 0 1 0 1.5H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6.25a.75.75 0 0 1-1.5 0V3a.5.5 0 0 0-.5-.5zm12.28 8.72a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06 0l-1-1a.75.75 0 1 1 1.06-1.06l.47.47l1.97-1.97a.75.75 0 0 1 1.06 0M4.75 4a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5zM4 8a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5A.75.75 0 0 1 4 8m.75 2.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeTask(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.78 3.43a.75.75 0 0 0-1.06-1.06L7.162 7.927L5.289 5.991a.75.75 0 1 0-1.078 1.043l2.403 2.484a.75.75 0 0 0 1.07.01zM2.5 9.75a.75.75 0 0 0-1.5 0V12a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V9.75a.75.75 0 0 0-1.5 0V12a1.5 1.5 0 0 1-1.5 1.5H4A1.5 1.5 0 0 1 2.5 12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeTestCase(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 0a.75.75 0 0 0 0 1.5H4V12a4 4 0 0 0 8 0V1.5h.25a.75.75 0 0 0 0-1.5zm6.75 1.5h-5V4h5zm0 8.5H8.75a.75.75 0 0 0 0 1.5h1.75v.5a2.5 2.5 0 0 1-5 0V5.5h5V7H8.75a.75.75 0 0 0 0 1.5h1.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueTypeTicket(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m.879 10.879l.182.182l.878.878a1.5 1.5 0 0 1 2.121 2.122l.88.878l.181.182L6 16l1.06-1.06l7.88-7.88L16 6l-.879-.879l-.182-.182l-.878-.878a1.5 1.5 0 0 1-2.122-2.121l-.878-.88l-.182-.181L10 0L8.94 1.06L1.06 8.94L0 10zM3 10c-.268 0-.529.036-.777.102L2.121 10l4.41-4.409l3.878 3.879L6 13.879l-.102-.102A3 3 0 0 0 3 10m8.47-1.591L13.879 6l-.102-.102a3 3 0 0 1-3.675-3.675L10 2.121l-2.409 2.41z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueUpdate(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3.25a.75.75 0 0 0 0-1.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 .5.5v5.25a.75.75 0 0 0 1.5 0V4.364l2.19 1.14a.25.25 0 0 1 .107.338l-1.072 2.062a.75.75 0 0 0 1.33.692l1.073-2.062a1.75 1.75 0 0 0-.745-2.36l-2.912-1.516A2 2 0 0 0 9 1zm10 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Issues(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5h6a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5M1 3a2 2 0 0 1 2-2h6a2 2 0 0 1 1.97 1.658l2.913 1.516a1.75 1.75 0 0 1 .744 2.36l-3.878 7.45a.753.753 0 0 1-.098.145c-.36.526-.965.871-1.651.871H3a2 2 0 0 1-2-2zm10 7.254l2.297-4.413a.25.25 0 0 0-.106-.337L11 4.364z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 2a.75.75 0 0 0 0 1.5h1.93l-2.412 9H4A.75.75 0 0 0 4 14h5.5a.75.75 0 0 0 0-1.5H7.57l2.412-9H12A.75.75 0 0 0 12 2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iteration(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.391 4.844a3.375 3.375 0 1 1 1.485 6.406h-.002a3.865 3.865 0 0 1-.06 0H.75a.75.75 0 0 0 0 1.5h12.379l-.97.97a.75.75 0 1 0 1.061 1.06l2.25-2.25L16 12l-.53-.53l-2.25-2.25a.75.75 0 1 0-1.06 1.06l.97.97H8.392A4.874 4.874 0 1 0 .018 8.298a.75.75 0 1 0 1.495-.13A3.375 3.375 0 0 1 3.39 4.844Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.358 8.763l.675-.674l-.325-.897A3.5 3.5 0 1 1 10 9.5H7.5v1H6.379l-.44.44l-1 1l-.439.439V13.5h-2v-.879zM6 15v-2l1-1h2v-1h1a5 5 0 1 0-4.703-3.297L1 12v3zm5-9a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 3.5H2a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5V4a.5.5 0 0 0-.5-.5M2 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm2 8.75a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1-.75-.75M11 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2M9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0M5 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m8.5-3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-4 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-2-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-4 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kubernetes(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.449 9.15a.283.283 0 0 1 .067.006l.003-.004l1.734.293a3.435 3.435 0 0 1-1.388 1.744l-.673-1.626l.002-.002a.294.294 0 0 1 .255-.41Zm-2.732.073a.294.294 0 0 1 .085.344l.005.006l-.666 1.61A3.448 3.448 0 0 1 4.757 9.45l1.72-.292l.002.004a.288.288 0 0 1 .238.06Zm1.433.681a.293.293 0 0 1 .108.113h.006l.848 1.532a3.443 3.443 0 0 1-2.218-.002l.845-1.529h.002a.293.293 0 0 1 .409-.114M9.633 7.63l1.3-1.163c.4.647.575 1.408.502 2.164l-1.684-.485l-.002-.007a.294.294 0 0 1-.115-.504V7.63Zm-1.067-.884a.294.294 0 0 1-.048-.149h-.002l-.099-1.744a3.46 3.46 0 0 1 1.995.961L8.99 6.824l-.005-.003a.294.294 0 0 1-.418-.075Zm-1.409.137a.294.294 0 0 1-.146-.055l-.003.001l-1.433-1.015a3.43 3.43 0 0 1 2.008-.961l-.1 1.746l-.007.004a.293.293 0 0 1-.319.28m-.775 1.185a.294.294 0 0 1-.135.08l-.001.005l-1.677.484a3.429 3.429 0 0 1 .487-2.17l1.308 1.17l-.002.006a.293.293 0 0 1 .02.425M8 8.893l-.483-.232l-.12-.519l.334-.416h.536l.333.417l-.12.519z"/><path fill="currentColor" fill-rule="evenodd" d="m14.6 3.714l1.375 5.97a1.049 1.049 0 0 1-.04.595a1.041 1.041 0 0 1-.166.298l-3.85 4.789a1.05 1.05 0 0 1-.372.29c-.026.013-.054.024-.082.035a1.086 1.086 0 0 1-.377.071l-6.175.002a1.096 1.096 0 0 1-.778-.338l-.007-.007c-.016-.017-.032-.034-.047-.053L.231 10.58a1.101 1.101 0 0 1-.167-.3a1.067 1.067 0 0 1-.039-.593l1.373-5.971A1.073 1.073 0 0 1 1.973 3L7.537.341a1.072 1.072 0 0 1 .923 0l5.564 2.657a1.074 1.074 0 0 1 .576.716m-4.341 8.427a1.123 1.123 0 0 1-.056-.134a4.334 4.334 0 0 0 1.92-2.415l.142.025a.263.263 0 0 1 .186-.07c.2.039.394.096.581.172c.098.044.197.082.299.115l.053.012l.033.006l.007.002l.004.001a.334.334 0 1 0 .148-.649a1.946 1.946 0 0 1-.034-.008l-.063-.015a2.81 2.81 0 0 0-.319-.025a3.11 3.11 0 0 1-.597-.098a.374.374 0 0 1-.144-.145l-.134-.039a4.327 4.327 0 0 0-.694-2.997a10.898 10.898 0 0 1 .118-.106a.263.263 0 0 1 .063-.188c.153-.132.32-.249.495-.348c.095-.049.188-.103.277-.162a1.449 1.449 0 0 0 .073-.06a.334.334 0 1 0-.415-.52l-.02.016c-.018.014-.039.03-.053.043a2.746 2.746 0 0 0-.22.233c-.137.15-.287.286-.45.406a.37.37 0 0 1-.201.022l-.127.09A4.364 4.364 0 0 0 8.367 3.97a11.302 11.302 0 0 1-.008-.148a.265.265 0 0 1-.11-.167a3.12 3.12 0 0 1 .039-.604c.021-.105.036-.21.045-.317V2.64a.334.334 0 1 0-.666.01v.085c.009.107.024.212.045.317c.032.2.044.402.037.604a.363.363 0 0 1-.108.172l-.008.141a4.283 4.283 0 0 0-2.778 1.336a6.015 6.015 0 0 1-.12-.086a.263.263 0 0 1-.198-.019a3.107 3.107 0 0 1-.449-.406a2.78 2.78 0 0 0-.219-.232a1.792 1.792 0 0 0-.059-.048l-.015-.011a.397.397 0 0 0-.232-.088a.32.32 0 0 0-.265.118a.353.353 0 0 0 .082.49l.005.004l.018.015l.05.04c.09.06.182.114.277.162c.176.1.342.216.495.348a.37.37 0 0 1 .067.192l.107.095a4.31 4.31 0 0 0-.68 3.004l-.14.041a.46.46 0 0 1-.143.145a3.108 3.108 0 0 1-.596.098a2.78 2.78 0 0 0-.32.025a1.45 1.45 0 0 0-.062.015l-.027.006h-.003l-.005.002a.334.334 0 1 0 .148.649h.005l.005-.002h.002c.008-.003.016-.004.025-.006l.061-.014a2.71 2.71 0 0 0 .299-.115a3.12 3.12 0 0 1 .58-.17a.37.37 0 0 1 .192.067l.145-.025a4.334 4.334 0 0 0 1.92 2.398l-.06.145a.33.33 0 0 1 .03.188c-.084.19-.185.372-.302.542a2.773 2.773 0 0 0-.179.266l-.03.063l-.012.027a.334.334 0 1 0 .601.285v-.001a3.145 3.145 0 0 1 .042-.085c.038-.1.07-.201.096-.305c.057-.21.144-.412.258-.598a.27.27 0 0 1 .143-.07l.076-.137a4.31 4.31 0 0 0 3.073.008a10.147 10.147 0 0 1 .07.128a.26.26 0 0 1 .17.103a3.1 3.1 0 0 1 .23.56c.026.104.058.206.096.306l.03.063l.011.022a.333.333 0 0 0 .568.106a.335.335 0 0 0 .034-.39l-.012-.026l-.031-.064a2.774 2.774 0 0 0-.179-.266a3.098 3.098 0 0 1-.295-.528a.26.26 0 0 1 .026-.197Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KubernetesAgent(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.223.183a1.75 1.75 0 0 1 1.518 0l4.996 2.406c.48.231.829.668.947 1.187l1.234 5.407a1.75 1.75 0 0 1-.338 1.48l-3.233 4.054a.75.75 0 1 1-1.173-.935l3.233-4.054a.25.25 0 0 0 .049-.212L13.222 4.11a.25.25 0 0 0-.136-.17L8.09 1.536a.25.25 0 0 0-.217 0L2.878 3.94a.25.25 0 0 0-.136.17L1.508 9.515a.25.25 0 0 0 .049.212l3.233 4.054a.75.75 0 1 1-1.173.935L.384 10.663a1.75 1.75 0 0 1-.338-1.48L1.28 3.776a1.75 1.75 0 0 1 .947-1.187zM5.22 11.16a.75.75 0 0 0 0 1.06l2.25 2.25L8 15l.53-.53l2.25-2.25a.75.75 0 1 0-1.06-1.06l-.97.97V5.87l.97.97a.75.75 0 1 0 1.06-1.06L8.53 3.53L8 3l-.53.53l-2.25 2.25a.75.75 0 0 0 1.06 1.06l.97-.97v6.26l-.97-.97a.75.75 0 0 0-1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.664 1a1.75 1.75 0 0 0-1.237.512L1.514 8.419a1.75 1.75 0 0 0-.001 2.475L5.1 14.48a1.75 1.75 0 0 0 2.474 0l6.914-6.906A1.75 1.75 0 0 0 15 6.335V1zm-.177 1.573a.25.25 0 0 1 .177-.073H13.5v3.835a.25.25 0 0 1-.073.177L6.513 13.42a.25.25 0 0 1-.353 0L2.574 9.833a.25.25 0 0 1 0-.353zM11 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Labels(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.427.512A1.75 1.75 0 0 1 7.664 0H13v3h3v5.335c0 .465-.185.91-.513 1.239L9.573 15.48a1.75 1.75 0 0 1-2.473 0l-2.293-2.293l-1.293-1.293l-3-3a1.75 1.75 0 0 1 0-2.475L6.428.512ZM11.5 1.5V3h-.836a1.75 1.75 0 0 0-1.237.512L3.514 9.419c-.06.06-.115.123-.165.19L1.574 7.833a.25.25 0 0 1 0-.353l5.913-5.907a.25.25 0 0 1 .177-.073zM5.866 12.126l-1.292-1.293a.25.25 0 0 1 0-.353l5.913-5.907a.25.25 0 0 1 .177-.073H14.5v3.835a.25.25 0 0 1-.073.177L8.513 14.42a.25.25 0 0 1-.353 0zM12 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leave(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.16 4.28a.75.75 0 1 1 1.06-1.06l3.25 3.25L16 7l-.53.53l-3.25 3.25a.75.75 0 0 1-1.06-1.06l1.97-1.97H3.25a1.75 1.75 0 1 0 0 3.5h2a.75.75 0 0 1 0 1.5h-2a3.25 3.25 0 0 1 0-6.5h9.88z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelUp(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.28 5.84a.75.75 0 0 1-1.06-1.06l3.25-3.25L8 1l.53.53l3.25 3.25a.75.75 0 0 1-1.06 1.06L8.75 3.87V12a1.5 1.5 0 0 0 1.5 1.5h4a.75.75 0 0 1 0 1.5h-4a3 3 0 0 1-3-3V3.87z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func License(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 10.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9M14 6a5.997 5.997 0 0 1-2.886 5.13l.58 3.185L12 16l-1.623-.544L8 14.66l-2.377.796L4 16l.306-1.684l.58-3.187A6 6 0 1 1 14 6m-7.748 6h3.496l.322 1.772l-1.594-.534l-.476-.16l-.476.16l-1.594.534zM9.5 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M11 6a3 3 0 1 1-6 0a3 3 0 0 1 6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LicenseSm(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 6.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5M10 4a3.997 3.997 0 0 1-1.842 3.369l.517 2.843L9 12l-1.66-.741L6 10.66l-1.34.599L3 12l.325-1.788l.517-2.843A4 4 0 1 1 10 4M5.252 8h1.496l.268 1.47l-.404-.18L6 9.017l-.612.273l-.404.18zM6 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.929 3.132a2.078 2.078 0 1 1 2.94 2.94l-.65.648a.75.75 0 0 0 1.061 1.06l.649-.648a3.579 3.579 0 0 0-5.06-5.06L6.218 4.72a3.578 3.578 0 0 0 0 5.06a.75.75 0 0 0 1.061-1.06a2.078 2.078 0 0 1 0-2.94zm-.15 3.086a.75.75 0 0 0-1.057 1.064c.816.81.818 2.13.004 2.942l-2.654 2.647a2.08 2.08 0 0 1-2.94-2.944l.647-.647a.75.75 0 0 0-1.06-1.06l-.648.647a3.58 3.58 0 0 0 5.06 5.066l2.654-2.647a3.575 3.575 0 0 0-.007-5.068Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm1.102 4.297a1.195 1.195 0 1 0 0-2.39a1.195 1.195 0 0 0 0 2.39m1 7.516V6.234h-2v6.579zM6.43 6.234h2v.881c.295-.462.943-1.084 2.148-1.084c1.438 0 2.219.953 2.219 2.766c0 .087.008.484.008.484v3.531h-2v-3.53c0-.485-.102-1.438-1.18-1.438c-1.079 0-1.17 1.198-1.195 1.982v2.986h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListBulleted(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.75a1 1 0 1 0 0-2a1 1 0 0 0 0 2M5.75 3a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zm0 4.25a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zm-.75 5a.75.75 0 0 1 .75-.75h8.5a.75.75 0 0 1 0 1.5h-8.5a.75.75 0 0 1-.75-.75M3 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1 5.25a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListIndent(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.75 3a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zm2 4.25a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zM.22 5.72a.75.75 0 0 1 1.06 0l1.75 1.75l.53.53l-.53.53l-1.75 1.75A.75.75 0 1 1 .22 9.22L1.44 8L.22 6.78a.75.75 0 0 1 0-1.06m4.53 5.78a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListNumbered(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2h2v4H1V3H0zm1.637 9.008H0v-1h1.637a1.382 1.382 0 0 1 .803 2.506L1.76 13H3v1H0v-.972L1.859 11.7a.382.382 0 0 0-.222-.693ZM4.75 3a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 0-1.5zm0 4.25a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 0-1.5zm-.75 5a.75.75 0 0 1 .75-.75h9.5a.75.75 0 0 1 0 1.5h-9.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOutdent(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.75 3a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zm2 4.25a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zM3.34 5.72a.75.75 0 0 0-1.06 0L.53 7.47L0 8l.53.53l1.75 1.75a.75.75 0 1 0 1.06-1.06L2.12 8l1.22-1.22a.75.75 0 0 0 0-1.06m1.41 5.78a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListTask(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.78 3.28a.75.75 0 0 0-1.06-1.06L1.75 4.19l-.47-.47A.75.75 0 0 0 .22 4.78l1 1a.75.75 0 0 0 1.06 0zM6 3.75A.75.75 0 0 1 6.75 3h8.5a.75.75 0 0 1 0 1.5h-8.5A.75.75 0 0 1 6 3.75M6 8a.75.75 0 0 1 .75-.75h8.5a.75.75 0 0 1 0 1.5h-8.5A.75.75 0 0 1 6 8m.75 3.5a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zm-1.97-1.28a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06 0l-1-1a.75.75 0 1 1 1.06-1.06l.47.47l1.97-1.97a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LivePreview(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5h10a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5M1 3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2zm9.534 4.172a1 1 0 0 1 0 1.656L7.56 10.841A1 1 0 0 1 6 10.012V5.988a1 1 0 0 1 1.56-.829z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LiveStream(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.596 1.164a.75.75 0 0 1 1.054.118a10.75 10.75 0 0 1-.017 13.438a.75.75 0 1 1-1.17-.94a9.25 9.25 0 0 0 .015-11.562a.75.75 0 0 1 .118-1.054m-3.84 2.97a.75.75 0 0 1 1.048.166a6.281 6.281 0 0 1 0 7.381a.75.75 0 1 1-1.214-.881a4.781 4.781 0 0 0 0-5.619a.75.75 0 0 1 .167-1.047ZM3 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 6a4.47 4.47 0 0 1-.883 2.677L8 13.5L4.383 8.677A4.5 4.5 0 1 1 12.5 6M14 6c0 1.34-.439 2.576-1.18 3.574L8.937 14.75L8 16l-.938-1.25L3.18 9.574A6 6 0 1 1 14 6M8 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationDot(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.617 8.677a4.5 4.5 0 1 0-7.235 0L8 13.5zm1.203.897a6 6 0 1 0-9.64 0L6.875 14.5H4.75a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5H9.125zM8 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 5a4 4 0 1 1 8 0v1h1a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1zm6.5 0v1h-5V5a2.5 2.5 0 0 1 5 0m-7 2.5v6h9v-6zM9 12V9H7v3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.523.029a4 4 0 0 1 4.37 3.05a.75.75 0 0 1-1.46.345a2.5 2.5 0 0 0-4.933.592V7H13a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h1.75H4V4.024A4 4 0 0 1 7.523.029M3.5 8.5v6h9v-6zM9 10H7v3h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Log(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2.5v11h9v-11zM3 1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm5 10a.75.75 0 0 1 .75-.75h1.75a.75.75 0 0 1 0 1.5H8.75A.75.75 0 0 1 8 11m-2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m2-4a.75.75 0 0 1 .75-.75h1.75a.75.75 0 0 1 0 1.5H8.75A.75.75 0 0 1 8 8M6 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m2-4a.75.75 0 0 1 .75-.75h1.75a.75.75 0 0 1 0 1.5H8.75A.75.75 0 0 1 8 5M6 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LongArrow(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.159 10.72a.75.75 0 1 0 1.06 1.06l3.25-3.25L15 8l-.53-.53l-3.25-3.25a.75.75 0 0 0-1.061 1.06l1.97 1.97H1.75a.75.75 0 1 0 0 1.5h10.379z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MachineLearning(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.75 3.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5M9.5 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M8 11a3.001 3.001 0 0 0 2.905-2.25h1.845c.071.095.155.179.25.25v3.75a1.25 1.25 0 1 0 1.5 0V9a1.25 1.25 0 1 0-1.75-1.75h-1.845A3.005 3.005 0 0 0 8.75 5.095V3.25A1.25 1.25 0 1 0 7 1.5H3.25a1.25 1.25 0 1 0 0 1.5H7c.071.095.155.179.25.25v1.845A3.001 3.001 0 0 0 8 11m-5.75 4a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5m7-1.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0M3.5 8A1.25 1.25 0 1 1 1 8a1.25 1.25 0 0 1 2.5 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.5h12a.5.5 0 0 1 .5.5v.572L8 8.286L1.5 4.572V4a.5.5 0 0 1 .5-.5m-.5 2.8V12a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5V6.3L8.372 9.8L8 10.014L7.628 9.8zM0 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkdownMark(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.308 5.308v5.23h1.538v-3l1.539 1.924l1.538-1.924v3h1.539v-5.23H6.923L5.385 7.23L3.846 5.308zM9.615 8l2.308 2.539L14.231 8h-1.539V5.308h-1.538V8z"/><path fill="currentColor" fill-rule="evenodd" d="M1.154 3C.517 3 0 3.517 0 4.154v7.538c0 .637.517 1.154 1.154 1.154h13.692c.637 0 1.154-.517 1.154-1.154V4.154C16 3.517 15.483 3 14.846 3zM.769 4.154c0-.213.172-.385.385-.385h13.692c.213 0 .385.172.385.385v7.538a.385.385 0 0 1-.385.385H1.154a.385.385 0 0 1-.385-.385z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkdownMarkSolid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.846 12.923H1.154A1.154 1.154 0 0 1 0 11.77V4.231a1.154 1.154 0 0 1 1.154-1.154h13.692A1.154 1.154 0 0 1 16 4.23v7.538a1.154 1.154 0 0 1-1.154 1.154Zm-11-2.308v-3l1.539 1.923l1.538-1.923v3h1.539v-5.23H6.923L5.385 7.308L3.846 5.385H2.308v5.23zM14.154 8h-1.539V5.385h-1.538V8H9.538l2.308 2.692z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarqueeSelection(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 1H1v2.25a.75.75 0 0 0 1.5 0V2.5h.75a.75.75 0 0 0 0-1.5zm0 5a.75.75 0 0 1 .75.75v2.5a.75.75 0 0 1-1.5 0v-2.5A.75.75 0 0 1 1.75 6m7.106 2.856l1.507 3.517l2.01-2.01zm-.558-1.871l5.344 2.29a1 1 0 0 1 .314 1.627l-.997.996l1.821 1.822a.75.75 0 1 1-1.06 1.06l-1.822-1.821l-.996.997a1 1 0 0 1-1.627-.314l-2.29-5.344c-.356-.83.483-1.669 1.313-1.313M1.75 15H1v-2.25a.75.75 0 0 1 1.5 0v.75h.75a.75.75 0 0 1 0 1.5zM6 1.75A.75.75 0 0 1 6.75 1h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 6 1.75M14.25 1H15v2.25a.75.75 0 0 1-1.5 0V2.5h-.75a.75.75 0 0 1 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mastodon(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.498 3.706C15.264 1.986 13.749.632 11.954.37C11.65.325 10.503.164 7.844.164h-.02c-2.66 0-3.23.161-3.533.206C2.546.625.951 1.842.565 3.582C.379 4.438.359 5.388.393 6.259C.443 7.51.453 8.756.567 10c.08.827.218 1.647.414 2.454c.368 1.49 1.856 2.731 3.314 3.237a8.98 8.98 0 0 0 4.848.253a7.25 7.25 0 0 0 .525-.141c.39-.123.849-.26 1.186-.502a.039.039 0 0 0 .01-.013a.036.036 0 0 0 .005-.016v-1.206a.035.035 0 0 0-.028-.034a.038.038 0 0 0-.016 0c-1.03.243-2.087.365-3.146.363c-1.824 0-2.314-.856-2.455-1.212a3.722 3.722 0 0 1-.213-.955a.035.035 0 0 1 .028-.036a.035.035 0 0 1 .016 0a13.32 13.32 0 0 0 3.095.364c.25 0 .5 0 .751-.007c1.049-.03 2.154-.082 3.186-.281c.025-.006.051-.01.073-.016c1.627-.31 3.176-1.28 3.333-3.736c.006-.097.02-1.013.02-1.113c.002-.342.112-2.42-.015-3.697m-2.505 6.13h-1.71V5.69c0-.873-.368-1.318-1.116-1.318c-.822 0-1.234.526-1.234 1.566v2.27h-1.7v-2.27c0-1.04-.413-1.566-1.235-1.566c-.744 0-1.115.445-1.116 1.318v4.145h-1.71v-4.27c0-.873.226-1.566.677-2.08c.464-.513 1.074-.776 1.83-.776c.876 0 1.538.333 1.979.999l.426.706l.426-.706c.441-.666 1.103-.999 1.977-.999c.756 0 1.366.263 1.832.776c.45.513.676 1.207.676 2.08z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maximize(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 5.25a.75.75 0 0 1-1.5 0V3.56l-3.22 3.22a.75.75 0 1 1-1.06-1.06l3.22-3.22h-1.69a.75.75 0 0 1 0-1.5H15zM3.81 13.5l2.97-2.97a.75.75 0 1 0-1.06-1.06L2.5 12.69v-1.94a.75.75 0 0 0-1.5 0V15h4.25a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Media(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 2.5H3a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V3a.5.5 0 0 0-.5-.5M3 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2zm9 9.857L9.5 8l-2.476 2.83L5.5 9L4 10.8V12h8zM6.5 8a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaBroken(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h4.25a.75.75 0 0 0 0-1.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v4.25a.75.75 0 0 0 1.5 0V3a2 2 0 0 0-2-2zm12 9.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5h-2.25a.75.75 0 0 0 0 1.5H13a2 2 0 0 0 2-2zm-3.867-.883L9.5 8l-2.476 2.83L5.5 9L4 10.8V12h5zM6.5 8a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Merge(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.044 2.31a2.5 2.5 0 1 0-1.706.076v4.228a2.501 2.501 0 1 0 1.5 0V8.373a5.735 5.735 0 0 0 3.86 1.864a2.501 2.501 0 1 0 .01-1.504a4.254 4.254 0 0 1-3.664-2.922ZM11.5 10.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-6 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MergeRequest(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.34 1.22a.75.75 0 0 0-1.06 0L7.53 2.97L7 3.5l.53.53l1.75 1.75a.75.75 0 1 0 1.06-1.06l-.47-.47h.63c.69 0 1.25.56 1.25 1.25v4.614a2.501 2.501 0 1 0 1.5 0V5.5a2.75 2.75 0 0 0-2.75-2.75h-.63l.47-.47a.75.75 0 0 0 0-1.06M13.5 12.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-9 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1.5 0a2.5 2.5 0 1 1-3.25-2.386V5.886a2.501 2.501 0 1 1 1.5 0v4.228A2.501 2.501 0 0 1 6 12.5m-1.5-9a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MergeRequestClose(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.22 1.22a.75.75 0 0 1 1.06 0L3.5 2.44l1.22-1.22a.75.75 0 0 1 1.06 1.06L4.56 3.5l1.22 1.22a.75.75 0 0 1-1.06 1.06L3.5 4.56L2.28 5.78a.75.75 0 0 1-1.06-1.06L2.44 3.5L1.22 2.28a.75.75 0 0 1 0-1.06M7.5 3.5a.75.75 0 0 1 .75-.75h2.25a2.75 2.75 0 0 1 2.75 2.75v4.614a2.501 2.501 0 1 1-1.5 0V5.5c0-.69-.56-1.25-1.25-1.25H8.25a.75.75 0 0 1-.75-.75m5 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2m-8-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1.5 0a2.5 2.5 0 1 1-3.25-2.386V7.75a.75.75 0 0 1 1.5 0v2.364A2.501 2.501 0 0 1 6 12.5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MergeRequestCloseM(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.28 4.22a.75.75 0 0 0-1.06 1.06L6.94 8l-2.72 2.72a.75.75 0 1 0 1.06 1.06L8 9.06l2.72 2.72a.75.75 0 1 0 1.06-1.06L9.06 8l2.72-2.72a.75.75 0 0 0-1.06-1.06L8 6.94z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MergeRequestOpen(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.34 1.22a.75.75 0 0 0-1.06 0L7.53 2.97L7 3.5l.53.53l1.75 1.75a.75.75 0 1 0 1.06-1.06l-.47-.47h.63c.69 0 1.25.56 1.25 1.25v4.614a2.501 2.501 0 1 0 1.5 0V5.5a2.75 2.75 0 0 0-2.75-2.75h-.63l.47-.47a.75.75 0 0 0 0-1.06M13.5 12.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-9 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1.5 0a2.5 2.5 0 1 1-3.25-2.386V5.886a2.501 2.501 0 1 1 1.5 0v4.228A2.501 2.501 0 0 1 6 12.5m-1.5-9a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Messages(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.5 8.974l6 2.333V2.693l-6 2.333zm6.138-7.944L7 4H3a3 3 0 0 0 0 6h.969l1.066 4.445A2.029 2.029 0 0 0 7.008 16c1.524 0 2.504-1.618 1.797-2.969l-1.53-2.924l7.363 2.863A1 1 0 0 0 16 12.038V1.962a1 1 0 0 0-1.362-.932M5.51 10h.016l1.95 3.726a.529.529 0 1 1-.983.369L5.511 10ZM3 5.5h4v3H3a1.5 1.5 0 1 1 0-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minimize(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.78 2.28a.75.75 0 0 0-1.06-1.06L10.5 4.44V2.75a.75.75 0 0 0-1.5 0V7h4.25a.75.75 0 0 0 0-1.5h-1.69zM5.5 11.56v1.69a.75.75 0 0 0 1.5 0V9H2.75a.75.75 0 0 0 0 1.5h1.69l-3.22 3.22a.75.75 0 1 0 1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 1.5h6a.5.5 0 0 1 .5.5v12a.5.5 0 0 1-.5.5H5a.5.5 0 0 1-.5-.5V2a.5.5 0 0 1 .5-.5M3 2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm5 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileIssueClose(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.78 4.62a.75.75 0 0 1 0 1.06l-6.097 6.097a.75.75 0 0 1-1.069-.009L3.211 9.284a.75.75 0 1 1 1.078-1.042l1.873 1.935L11.72 4.62a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitor(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2.5v8h13v-8zM1 12h4v1.5H3.75a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5H11V12h4a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H1a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1m5.5 1.5V12h3v1.5zm5-7.75a.75.75 0 0 0-1.5 0v2.5a.75.75 0 0 0 1.5 0zM7.75 4a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-1.5 0v-3.5A.75.75 0 0 1 7.75 4M5.5 6.75a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MonitorLines(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 10.5v-8h13v8zM5 12H1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-4v1.5h1.25a.75.75 0 0 1 0 1.5h-8.5a.75.75 0 0 1 0-1.5H5zm1.5 0v1.5h3V12zM3 7.75A.75.75 0 0 1 3.75 7h5.5a.75.75 0 0 1 0 1.5h-5.5A.75.75 0 0 1 3 7.75m.75-3.25a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MonitorO(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 10.5v-8h13v8zM5 12H1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-4v1.5h1.25a.75.75 0 0 1 0 1.5h-8.5a.75.75 0 0 1 0-1.5H5zm1.5 0v1.5h3V12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Namespace(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.3 4.125a.75.75 0 0 0-1.3-.75l-4.9 8.5a.75.75 0 1 0 1.3.75zM8 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2.5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m3.5 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nature(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.98 2.5A6.37 6.37 0 0 0 15 2V1h-1.75a6.003 6.003 0 0 0-5.761 4.32A5.99 5.99 0 0 0 2.75 3H1v1a6 6 0 0 0 6 6h.25v4.25a.75.75 0 0 0 1.5 0V8H9a6 6 0 0 0 5.98-5.5m-6.203 4a4.5 4.5 0 0 1 4.473-4h.223A4.5 4.5 0 0 1 9 6.5zm-6.027-2a4.5 4.5 0 0 1 4.473 4H7a4.5 4.5 0 0 1-4.473-4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notifications(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a1 1 0 0 0-1 1v.1A5.002 5.002 0 0 0 3 7v4l-1.205 1.328c-.583.643-.127 1.672.74 1.672h3.733a2 2 0 0 0 3.464 0h3.733c.867 0 1.323-1.03.74-1.672L13 11V7a5.002 5.002 0 0 0-4-4.9V2a1 1 0 0 0-1-1M4.5 11.58l-.39.428l-.446.492h8.672l-.447-.492l-.389-.429V7a3.5 3.5 0 1 0-7 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationsOff(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a1 1 0 0 0-1 1v.1A5.002 5.002 0 0 0 3 7v4.94l-1.78 1.78a.75.75 0 1 0 1.06 1.06L14.776 2.284a.75.75 0 0 0-1.06-1.06l-2.211 2.21A4.991 4.991 0 0 0 9 2.1V2a1 1 0 0 0-1-1m0 2.5c.95 0 1.813.379 2.444.995L4.5 10.439V7A3.5 3.5 0 0 1 8 3.5m5 4.25a.75.75 0 0 0-1.5 0v3.817l.194.214l.65.719H6.75a.75.75 0 0 0-.728.932l.011.043A2.02 2.02 0 0 0 7.993 15c.737 0 1.389-.4 1.738-1h3.74c.868 0 1.324-1.028.742-1.671L13 10.989z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Object(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.5c.8 0 1.566-.145 2.274-.409l.665-2.046h2.15a6.472 6.472 0 0 0 1.405-4.327l-1.739-1.263l.665-2.045a6.512 6.512 0 0 0-3.68-2.674L8 3L6.26 1.736A6.512 6.512 0 0 0 2.58 4.41l.665 2.045l-1.739 1.263a6.472 6.472 0 0 0 1.406 4.327h2.15l.664 2.046A6.486 6.486 0 0 0 8 14.5M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M8 5.286l2.853 2.073l-1.09 3.354H6.237L5.147 7.36z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Organization(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0m-8 2.25a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0M4.516 8a3 3 0 1 0 4.932 2.81a3 3 0 1 0 0-5.62A3 3 0 1 0 4.516 8M6.5 7.25a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m4 2.25a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Overview(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 5.5v-3h3v3zM1 2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1zm8 .25a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 9 2.25M9.75 5a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5zM2.5 10.5v3h3v-3zM2 9a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1zm7.75.5a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5zm0 3.5a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.72 2.5L2.92 6h4.33V2.5zm3.03 0V6h4.33l-2.8-3.5zm-6.25 11v-6h11v6zM5.48 1a1 1 0 0 0-.78.375L1.22 5.726a1 1 0 0 0-.22.625V14a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6.35a1 1 0 0 0-.22-.624l-3.48-4.35A1 1 0 0 0 10.52 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperAirplane(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.968 1.966a.75.75 0 0 0-.934-.934l-12.5 3.75a.75.75 0 0 0-.18 1.355L5.952 8.99l-1.731 1.73a.75.75 0 1 0 1.06 1.061l1.731-1.73l2.852 4.595a.75.75 0 0 0 1.355-.18l3.75-12.5ZM8.101 8.96l2.159 3.48l2.417-8.056L8.1 8.96Zm3.515-5.637L3.56 5.74L7.04 7.9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.414 3.05L4.085 8.38a3 3 0 0 0 4.243 4.242l2.403-2.403a.75.75 0 1 1 1.06 1.06l-2.403 2.404a4.5 4.5 0 0 1-6.364-6.364l5.33-5.33a3.25 3.25 0 0 1 4.596 4.597l-5.33 5.329a2 2 0 0 1-2.828-2.828l5.33-5.33a.75.75 0 0 1 1.06 1.061l-5.33 5.33a.5.5 0 1 0 .708.706l5.33-5.329A1.75 1.75 0 0 0 9.413 3.05Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zm5 0a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 13.5v-1.879l7.28-7.28l1.88 1.879l-7.28 7.28zm10.22-8.341l.805-.805a.5.5 0 0 0 0-.708l-1.171-1.171a.5.5 0 0 0-.708 0l-.805.805l1.879 1.88ZM1 13.5V11l9.586-9.586a2 2 0 0 1 2.828 0l1.172 1.172a2 2 0 0 1 0 2.828L5 15H1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilSquare(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.287.303a1 1 0 1 1 1.415 1.414l-.707.708L13.58 1.01zm0 2.829l-6.873 6.873H6V8.59l6.873-6.874zM3 13.5a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h6.25a.75.75 0 0 0 0-1.5H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pipeline(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.272 3.864a4.5 4.5 0 1 0-.1 8.314a.75.75 0 1 0-.557-1.393a3 3 0 1 1 .066-5.543a.75.75 0 1 0 .59-1.378ZM11.5 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6m0 1.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9M6 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Planning(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.25 1a.75.75 0 0 1 .75.75V3h2a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h2V1.75a.75.75 0 0 1 1.5 0V3h5V1.75a.75.75 0 0 1 .75-.75M2.5 7.5v6h11v-6zm0-1.5h11V4.5h-11zm7.78 2.97a.75.75 0 0 1 0 1.06l-2.25 2.25a.75.75 0 0 1-1.06 0l-1.25-1.25a.75.75 0 1 1 1.06-1.06l.72.72l1.72-1.72a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.629 7.306a.835.835 0 0 1 0 1.388l-6.401 4.177C4.695 13.218 4 12.825 4 12.176V3.824c0-.649.695-1.042 1.228-.695z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.75 2.75a.75.75 0 0 0-1.5 0v4.5h-4.5a.75.75 0 0 0 0 1.5h4.5v4.5a.75.75 0 0 0 1.5 0v-4.5h4.5a.75.75 0 0 0 0-1.5h-4.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm4 9a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 8 11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareO(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 3.5h8a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H4a.5.5 0 0 1-.5-.5V4a.5.5 0 0 1 .5-.5M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm6 7a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 8 11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pod(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.248 1.87l4.485 2.562L8 7.136L3.268 4.432l4.484-2.563a.5.5 0 0 1 .496 0ZM2.5 10.84V5.72l4.75 2.714v5.409l-4.498-2.57a.5.5 0 0 1-.252-.435Zm6.25 3.004l4.498-2.57a.5.5 0 0 0 .252-.435V5.721L8.75 8.435zM8.992.567a2 2 0 0 0-1.984 0l-5 2.857A2 2 0 0 0 1 5.161v5.678a2 2 0 0 0 1.008 1.737l5 2.857a2 2 0 0 0 1.984 0l5-2.857A2 2 0 0 0 15 10.839V5.161a2 2 0 0 0-1.008-1.737z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Podcast(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.397 2.59a6.5 6.5 0 0 1 8.227 9.978a.75.75 0 0 0 1.067 1.054a8 8 0 1 0-11.382 0a.75.75 0 1 0 1.067-1.054A6.5 6.5 0 0 1 4.397 2.59M8 4.5a3.5 3.5 0 0 1 2.5 5.949a.75.75 0 1 0 1.072 1.05a5 5 0 1 0-7.144 0A.75.75 0 0 0 5.5 10.45A3.5 3.5 0 0 1 8 4.5M10 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-1.25 4.25a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 7.25a.75.75 0 0 0 1.5 0V.75a.75.75 0 0 0-1.5 0zm4.04 5.157A5.5 5.5 0 1 1 5 3.39a.75.75 0 1 0-.818-1.257a7 7 0 1 0 7.635 0A.75.75 0 0 0 11 3.39a5.5 5.5 0 0 1 .291 9.017Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Preferences(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M12 8a3 3 0 1 0-2.905-3.75H1.75a.75.75 0 0 0 0 1.5h7.345A3.001 3.001 0 0 0 12 8m-6.5 3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m1.405.75A3.001 3.001 0 0 1 1 11a3 3 0 0 1 5.905-.75h7.345a.75.75 0 0 1 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Profile(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.5a6.47 6.47 0 0 0 3.25-.87V11.5A2.25 2.25 0 0 0 9 9.25H7a2.25 2.25 0 0 0-2.25 2.25v2.13A6.47 6.47 0 0 0 8 14.5m4.75-3v.937a6.5 6.5 0 1 0-9.5 0V11.5a3.752 3.752 0 0 1 2.486-3.532a3 3 0 1 1 4.528 0A3.752 3.752 0 0 1 12.75 11.5M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M9.5 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Progress(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 5.5h8a2.5 2.5 0 0 1 0 5H4a2.5 2.5 0 0 1 0-5M0 8a4 4 0 0 1 4-4h8a4 4 0 0 1 0 8H4a4 4 0 0 1-4-4m4-1a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Project(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.5 14.5l-6-2.5V4l6-2.5zm-6.885-1.244A1 1 0 0 1 2 12.333V3.667a1 1 0 0 1 .615-.923L8.923.115A1.5 1.5 0 0 1 11 1.5V2h1.25c.966 0 1.75.783 1.75 1.75v8.5A1.75 1.75 0 0 1 12.25 14H11v.5a1.5 1.5 0 0 1-2.077 1.385zM11 12.5h1.25a.25.25 0 0 0 .25-.25v-8.5a.25.25 0 0 0-.25-.25H11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushRules(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.25 4.07V1.75a.75.75 0 0 1 1.5 0v2.32a4.001 4.001 0 0 1 0 7.86v2.32a.75.75 0 0 1-1.5 0v-2.32a4.001 4.001 0 0 1 0-7.86M4 10.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5M9.75 4a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5zm0 3.25a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5zm-.75 4a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 0 1.5h-5.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M4.927 4.99c-.285.429-.427.853-.427 1.27c0 .203.09.392.27.566c.18.174.4.26.661.26c.443 0 .744-.248.903-.746c.168-.475.373-.835.616-1.08c.243-.244.62-.366 1.134-.366c.439 0 .797.12 1.075.363c.277.242.416.54.416.892a.97.97 0 0 1-.136.502a1.91 1.91 0 0 1-.336.419a14.35 14.35 0 0 1-.648.558c-.34.282-.611.525-.812.73c-.2.205-.362.443-.483.713c-.322 1.245 1.35 1.345 1.736.456c.047-.086.118-.18.213-.284c.096-.103.223-.223.382-.36a41.14 41.14 0 0 0 1.194-1.034c.221-.204.412-.448.573-.73a1.95 1.95 0 0 0 .242-.984c0-.475-.141-.915-.424-1.32c-.282-.406-.682-.726-1.2-.962c-.518-.235-1.115-.353-1.792-.353c-.728 0-1.365.14-1.911.423c-.546.282-.961.637-1.246 1.066Zm2.14 7.08a1 1 0 1 0 2 0a1 1 0 0 0-2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionO(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M4.927 4.99c-.285.429-.427.853-.427 1.27c0 .203.09.392.27.566c.18.174.4.26.661.26c.443 0 .744-.248.903-.746c.168-.475.373-.835.616-1.08c.243-.244.62-.366 1.134-.366c.439 0 .797.12 1.075.363c.277.242.416.54.416.892a.97.97 0 0 1-.136.502a1.91 1.91 0 0 1-.336.419a14.35 14.35 0 0 1-.648.558c-.34.282-.611.525-.812.73c-.2.205-.362.443-.483.713c-.322 1.245 1.35 1.345 1.736.456c.047-.086.118-.18.213-.284c.096-.103.223-.223.382-.36a41.14 41.14 0 0 0 1.194-1.034c.221-.204.412-.448.573-.73a1.95 1.95 0 0 0 .242-.984c0-.475-.141-.915-.424-1.32c-.282-.406-.682-.726-1.2-.962c-.518-.235-1.115-.353-1.792-.353c-.728 0-1.365.14-1.911.423c-.546.282-.961.637-1.246 1.066Zm2.14 7.08a1 1 0 1 0 2 0a1 1 0 0 0-2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuickActions(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5h10a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5M1 3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2zm8.717 1.97a.75.75 0 1 0-1.434-.44l-2 6.5a.75.75 0 0 0 1.434.44z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quota(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.457 8.75a6.501 6.501 0 0 1-12.914 0h.707a.75.75 0 0 0 0-1.5h-.707a6.469 6.469 0 0 1 1.361-3.285l.5.5a.75.75 0 0 0 1.06-1.061l-.5-.5A6.469 6.469 0 0 1 7.25 1.543v.707a.75.75 0 0 0 1.5 0v-.707a6.47 6.47 0 0 1 3.285 1.361l-.5.5a.75.75 0 1 0 1.061 1.06l.5-.5a6.469 6.469 0 0 1 1.361 3.286h-.707a.75.75 0 0 0 0 1.5zM16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0m-8 4a2 2 0 0 0 .75-3.855V4.75a.75.75 0 0 0-1.5 0v3.395A2 2 0 0 0 8 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 3.75a.75.75 0 0 0-1.5 0v8.5a.75.75 0 0 0 1.5 0zM4.75 3a.75.75 0 0 0 0 1.5h7.5a.75.75 0 0 0 0-1.5zm0 4.25a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5zm-.75 5a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.095.28A8 8 0 0 0 1.5 3.335V1.75a.75.75 0 0 0-1.5 0V6h4.25a.75.75 0 1 0 0-1.5H2.523a6.5 6.5 0 1 1-.526 5.994a.75.75 0 0 0-1.385.575A8 8 0 1 0 10.095.279Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remove(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.75 3V1.5h4.5V3zm-1.5 0V1a1 1 0 0 1 1-1h5.5a1 1 0 0 1 1 1v2h2.5a.75.75 0 0 1 0 1.5h-.365l-.743 9.653A2 2 0 0 1 11.148 16H4.852a2 2 0 0 1-1.994-1.847L2.115 4.5H1.75a.75.75 0 0 1 0-1.5zm-.63 1.5h8.76l-.734 9.538a.5.5 0 0 1-.498.462H4.852a.5.5 0 0 1-.498-.462z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveAll(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.75 3V1.5h4.5V3zm-1.5 0V1a1 1 0 0 1 1-1h5.5a1 1 0 0 1 1 1v2h2.5a.75.75 0 0 1 0 1.5h-.365l-.743 9.653A2 2 0 0 1 11.148 16H4.852a2 2 0 0 1-1.994-1.847L2.115 4.5H1.75a.75.75 0 0 1 0-1.5zm-.63 1.5h8.76l-.734 9.538a.5.5 0 0 1-.498.462H4.852a.5.5 0 0 1-.498-.462z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.905.28A8 8 0 0 1 14.5 3.335V1.75a.75.75 0 0 1 1.5 0V6h-4.25a.75.75 0 0 1 0-1.5h1.727a6.5 6.5 0 1 0 .526 5.994a.75.75 0 1 1 1.385.575A8 8 0 1 1 5.905.279Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.84 9.72a.75.75 0 1 1-1.06 1.06L2.53 7.53L2 7l.53-.53l3.25-3.25a.75.75 0 0 1 1.06 1.06L4.87 6.25h4.378a4.75 4.75 0 0 1 4.75 4.75v3.25a.75.75 0 0 1-1.5 0V11a3.25 3.25 0 0 0-3.25-3.25H4.871l1.97 1.97Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Requirements(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 2.5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h5.25a.75.75 0 0 1 0 1.5H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6.25a.75.75 0 0 1-1.5 0V3a.5.5 0 0 0-.5-.5zm12.28 8.72a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06 0l-1-1a.75.75 0 1 1 1.06-1.06l.47.47l1.97-1.97a.75.75 0 0 1 1.06 0M4.75 4a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5zM4 8a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5A.75.75 0 0 1 4 8m.75 2.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Retry(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.32.029a8 8 0 0 1 7.18 3.307V1.75a.75.75 0 0 1 1.5 0V6h-4.25a.75.75 0 0 1 0-1.5h1.727A6.5 6.5 0 0 0 1.694 6.424A.75.75 0 1 1 .239 6.06A8 8 0 0 1 7.319.03Zm-3.4 14.852A8 8 0 0 0 15.76 9.94a.75.75 0 0 0-1.455-.364A6.5 6.5 0 0 1 2.523 11.5H4.25a.75.75 0 0 0 0-1.5H0v4.25a.75.75 0 0 0 1.5 0v-1.586a8 8 0 0 0 2.42 2.217" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReviewCheckmark(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2M8 0a2.5 2.5 0 0 1 2.45 2H13a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h2.55A2.5 2.5 0 0 1 8 0M7 5h3.5V3.5h2v11h-9v-11h2V5zm3.53 3.28a.75.75 0 1 0-1.06-1.06L7.5 9.19l-.47-.47a.75.75 0 0 0-1.06 1.06l1 1a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReviewList(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1.45-.5a2.5 2.5 0 0 0-4.9 0H3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1zM8 5H5.5V3.5h-2v11h9v-11h-2V5zM5 7.75A.75.75 0 0 1 5.75 7h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 5 7.75m.75 1.75a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReviewWarning(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1.45-.5a2.5 2.5 0 0 0-4.9 0H3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1zM8 5H5.5V3.5h-2v11h9v-11h-2V5zm1 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25-4.75a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m16 .776l.027-.803l-.803.027l-1.309.046A10.75 10.75 0 0 0 5.753 4.25H3.667A2.75 2.75 0 0 0 .962 6.504l-.8 4.36L0 11.75h2.69l1.56 1.56V16l.885-.162l4.36-.8a2.75 2.75 0 0 0 2.255-2.705v-2.086a10.75 10.75 0 0 0 4.204-8.162L16 .775ZM9.348 9.988l-4.2 2.1l-1.235-1.236l2.1-4.2a9.25 9.25 0 0 1 7.954-5.107l.506-.018l-.018.506a9.25 9.25 0 0 1-5.107 7.955M5.75 14.2v-.736l4.268-2.135l.232-.116v1.12a1.25 1.25 0 0 1-1.025 1.23zm-3.214-3.95l2.135-4.268l.115-.232h-1.12a1.25 1.25 0 0 0-1.229 1.025L1.8 10.25zM10.5 6a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0M12 6a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketLaunch(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.808 0a9.999 9.999 0 0 0-7.142 3H6.75A2.75 2.75 0 0 0 4 5.75V8h2l2 2v2h2.25A2.75 2.75 0 0 0 13 9.25V7.334a9.998 9.998 0 0 0 3-7.142V0zM6.44 6.5a9.964 9.964 0 0 1 1.015-2H6.75c-.69 0-1.25.56-1.25 1.25v.75zm3.06 4v-.94a9.966 9.966 0 0 0 2-1.015v.705c0 .69-.56 1.25-1.25 1.25zm4.88-8.88a8.502 8.502 0 0 0-6.71 5.928l.782.783a8.502 8.502 0 0 0 5.928-6.71Zm-11.6 8.66a.75.75 0 1 0-1.06-1.06l-1.5 1.5a.75.75 0 1 0 1.06 1.06zm3 1a.75.75 0 1 0-1.06-1.06l-4.5 4.5a.75.75 0 1 0 1.06 1.06zm1 1.94a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 0 1-1.06-1.06l1.5-1.5a.75.75 0 0 1 1.06 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.529 4.18a9.5 9.5 0 0 0-3.76-.677a.75.75 0 1 1-.037-1.5A11 11 0 0 1 13.997 13.27a.75.75 0 1 1-1.5-.036A9.5 9.5 0 0 0 6.53 4.18ZM4.624 8.803a4.5 4.5 0 0 0-1.838-.298a.75.75 0 1 1-.07-1.498a6 6 0 0 1 6.277 6.282a.75.75 0 0 1-1.498-.072a4.5 4.5 0 0 0-2.871-4.414M3 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scale(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a.75.75 0 0 1 .75.75V3h5.5a.75.75 0 0 1 0 1.5h-1.393l.565.79l1.086 1.521A2.7 2.7 0 0 1 15 8.375C15 9.825 13.88 11 12.5 11S10 9.825 10 8.375a2.7 2.7 0 0 1 .492-1.564l1.086-1.52l.565-.791H8.75v9.75a.749.749 0 0 1-.043.25h2.543a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1 0-1.5h2.543a.749.749 0 0 1-.043-.25V4.5H3.857l.565.79l1.086 1.521A2.7 2.7 0 0 1 6 8.375C6 9.825 4.88 11 3.5 11S1 9.825 1 8.375a2.7 2.7 0 0 1 .492-1.564l1.086-1.52l.565-.791H1.75a.75.75 0 0 1 0-1.5h5.5V1.75A.75.75 0 0 1 8 1m5.5 7.375c0-.27-.083-.508-.215-.695l-.785-1.1l-.785 1.1a1.201 1.201 0 0 0-.215.695c0 .691.516 1.125 1 1.125s1-.434 1-1.125M4.285 7.68c.132.187.215.425.215.695c0 .691-.516 1.125-1 1.125s-1-.434-1-1.125c0-.27.083-.508.215-.695l.785-1.1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScrollDown(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 14a2 2 0 1 1-4 0a2 2 0 0 1 4 0m1.78-8.841a.75.75 0 0 0-1.06 0l-1.97 1.97V.75a.75.75 0 0 0-1.5 0v6.379l-1.97-1.97a.75.75 0 0 0-1.06 1.06l3.25 3.25L8 10l.53-.53l3.25-3.25a.75.75 0 0 0 0-1.061" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScrollHandle(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16M6 3a1 1 0 0 1 1 1v8a1 1 0 1 1-2 0V4a1 1 0 0 1 1-1m4 0a1 1 0 0 1 1 1v8a1 1 0 1 1-2 0V4a1 1 0 0 1 1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScrollUp(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2a2 2 0 1 0-4 0a2 2 0 0 0 4 0m1.78 8.841a.75.75 0 0 1-1.06 0l-1.97-1.97v6.379a.75.75 0 0 1-1.5 0V8.871l-1.97 1.97a.75.75 0 1 1-1.06-1.06l3.25-3.25L8 6l.53.53l3.25 3.25a.75.75 0 0 1 0 1.061" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 7a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-.82 4.74a6 6 0 1 1 1.06-1.06l3.04 3.04a.75.75 0 1 1-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchDot(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 7a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-.82 4.74a6 6 0 1 1 1.06-1.06l3.04 3.04a.75.75 0 1 1-1.06 1.06zM9 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchMinus(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 7a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-.82 4.74a6 6 0 1 1 1.06-1.06l3.04 3.04a.75.75 0 1 1-1.06 1.06zM4 7a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 4 7" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchPlus(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 11.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9M7 13c1.387 0 2.663-.47 3.68-1.26l3.04 3.04a.75.75 0 1 0 1.06-1.06l-3.04-3.04A6 6 0 1 0 7 13m0-3a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 7 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchResults(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 1a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5zM1 4.75A.75.75 0 0 1 1.75 4H7a.75.75 0 0 1 0 1.5H1.75A.75.75 0 0 1 1 4.75m9 7.75a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m0 1.5c.834 0 1.607-.255 2.248-.691l1.472 1.471a.75.75 0 1 0 1.06-1.06l-1.471-1.472A4 4 0 1 0 10 14M1.75 7a.75.75 0 0 0 0 1.5H4A.75.75 0 0 0 4 7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSm(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 5.5a3 3 0 1 1-6 0a3 3 0 0 1 6 0m-.393 3.668a4.5 4.5 0 1 1 1.06-1.06l2.613 2.612a.75.75 0 1 1-1.06 1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.43 1.168l-.74-.123zm-.156.939l.74.123zm-.173.2l.237.711zM4.02 3.509l-.498-.56zm-.26.05l.263-.702zm-.893-.334l.263-.703zm-.608.218l-.65-.375zM1.183 5.307l-.65-.375zm.115.636l.477-.579l-.477.58Zm.736.606l-.477.579l.477-.58Zm.086.25l.735.149l-.735-.15Zm0 2.403l.735-.15zm-.086.25l.476.578l-.476-.579Zm-.736.605l-.476-.58zm-.115.636l-.65.375zm1.077 1.864l.65-.375zm.608.218l.263.703zm.893-.334l-.263-.703zm.26.05l-.498.56zm2.08 1.202l.237-.711zm.173.2l.74-.123zm.156.94l.74-.124l-.74.123Zm3.14 0l-.74-.124l.74.123Zm.156-.94l.74.123zm.173-.2l.238.712zm2.08-1.202l-.497-.561zm.26-.05l.263-.703zm.893.334l-.263.703zm.609-.218l-.65-.375zm1.076-1.864l.65.375zm-.115-.636l-.477.579zm-.736-.606l-.476.58zm-.086-.25l.735.15zm0-2.403l.735-.15zm.086-.25l-.476-.578zm.736-.605l-.477-.579l.477.58Zm.115-.636l.65-.375zM13.74 3.443l-.65.375zm-.609-.218l-.263-.703zm-.893.334l-.263-.702zm-.26-.05l-.497.561l.497-.56ZM9.9 2.307l-.237.711zm-.173-.2l.74-.123zm-.156-.94l-.74.124zM6.924 1.5h2.152V0H6.924zm.246-.209a.25.25 0 0 1-.246.209V0A1.25 1.25 0 0 0 5.69 1.044zm-.156.94l.156-.94l-1.48-.246l-.156.939zm-.676.787c.343-.114.613-.41.676-.788l-1.48-.246a.494.494 0 0 1 .33-.389zM4.518 4.07a5.244 5.244 0 0 1 1.82-1.052l-.474-1.423a6.744 6.744 0 0 0-2.34 1.353zm-1.02.192c.36.134.75.048 1.02-.192l-.995-1.122a.494.494 0 0 1 .501-.091zm-.893-.335l.893.335l.526-1.405l-.893-.335zm.304-.11a.25.25 0 0 1-.304.11l.526-1.405a1.25 1.25 0 0 0-1.521.546l1.3.75ZM1.833 5.683l1.076-1.864l-1.299-.75L.534 4.932zm-.058-.318a.25.25 0 0 1 .058.318l-1.3-.75a1.25 1.25 0 0 0 .289 1.59zm.735.606l-.735-.606l-.953 1.158l.735.606zm.345.978a1.006 1.006 0 0 0-.345-.978l-.953 1.158a.494.494 0 0 1-.172-.48zM2.75 8c0-.361.036-.713.105-1.052l-1.47-.3c-.088.438-.135.89-.135 1.352zm.105 1.052A5.277 5.277 0 0 1 2.75 8h-1.5c0 .462.047.914.135 1.351zm-.345.978c.296-.243.417-.624.345-.978l-1.47.3a.494.494 0 0 1 .172-.48zm-.735.606l.735-.606l-.953-1.158l-.735.606zm.058-.318a.25.25 0 0 1-.058.318L.822 9.478a1.25 1.25 0 0 0-.288 1.59zm1.076 1.864l-1.076-1.864l-1.3.75l1.077 1.864zm-.304-.109a.25.25 0 0 1 .304.11l-1.299.75c.306.528.949.76 1.521.545zm.893-.335l-.893.335l.526 1.405l.893-.335zm1.02.192a1.006 1.006 0 0 0-1.02-.191l.526 1.404a.494.494 0 0 1-.5-.091zm1.82 1.052a5.243 5.243 0 0 1-1.82-1.052l-.995 1.122a6.745 6.745 0 0 0 2.34 1.353zm.676.788a1.006 1.006 0 0 0-.676-.788l-.474 1.423a.494.494 0 0 1-.33-.388zm.156.939l-.156-.94l-1.48.248l.157.939l1.48-.247Zm-.246-.209a.25.25 0 0 1 .246.209l-1.48.247A1.25 1.25 0 0 0 6.925 16zm2.152 0H6.924V16h2.152zm-.246.209a.25.25 0 0 1 .246-.209V16a1.25 1.25 0 0 0 1.233-1.044zm.156-.94l-.156.94l1.48.247l.156-.94zm.676-.787c-.343.114-.613.41-.676.788l1.48.246a.494.494 0 0 1-.33.389zm1.82-1.052a5.244 5.244 0 0 1-1.82 1.052l.474 1.423a6.745 6.745 0 0 0 2.34-1.353zm1.02-.191a1.006 1.006 0 0 0-1.02.19l.995 1.123a.494.494 0 0 1-.501.091l.526-1.405Zm.893.334l-.893-.335l-.526 1.405l.893.335zm-.304.11a.25.25 0 0 1 .304-.11l-.526 1.405a1.25 1.25 0 0 0 1.521-.546l-1.299-.75Zm1.076-1.865l-1.076 1.864l1.299.75l1.076-1.864zm.058.318a.25.25 0 0 1-.058-.318l1.3.75a1.25 1.25 0 0 0-.289-1.59zm-.735-.606l.735.606l.953-1.158l-.735-.606zm-.345-.978c-.072.354.049.735.345.978l.953-1.158c.15.123.206.311.172.48zM13.25 8c0 .361-.036.713-.105 1.052l1.47.3c.088-.438.135-.89.135-1.352zm-.105-1.052c.069.339.105.69.105 1.052h1.5c0-.462-.046-.914-.135-1.351zm.345-.978a1.007 1.007 0 0 0-.345.978l1.47-.3a.494.494 0 0 1-.172.48zm.735-.606l-.735.606l.953 1.158l.735-.606zm-.058.318a.25.25 0 0 1 .058-.318l.953 1.158a1.25 1.25 0 0 0 .288-1.59zm-1.076-1.864l1.076 1.864l1.3-.75l-1.077-1.864zm.304.109a.25.25 0 0 1-.304-.11l1.299-.75a1.25 1.25 0 0 0-1.521-.545zm-.893.334l.893-.334l-.526-1.405l-.893.335zm-1.02-.19c.27.24.66.325 1.02.19l-.526-1.404a.494.494 0 0 1 .5.091zm-1.82-1.053a5.244 5.244 0 0 1 1.82 1.052l.995-1.122a6.744 6.744 0 0 0-2.34-1.353zm-.676-.788c.063.379.333.674.676.788l.474-1.423a.494.494 0 0 1 .33.389zm-.156-.939l.156.94l1.48-.247l-.157-.94l-1.48.247Zm.246.209a.25.25 0 0 1-.246-.209l1.48-.246A1.25 1.25 0 0 0 9.075 0v1.5ZM10.25 8A2.25 2.25 0 0 1 8 10.25v1.5A3.75 3.75 0 0 0 11.75 8zM8 5.75A2.25 2.25 0 0 1 10.25 8h1.5A3.75 3.75 0 0 0 8 4.25zM5.75 8A2.25 2.25 0 0 1 8 5.75v-1.5A3.75 3.75 0 0 0 4.25 8zM8 10.25A2.25 2.25 0 0 1 5.75 8h-1.5A3.75 3.75 0 0 0 8 11.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeverityCritical(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.706.185l4.088 2.31c.437.246.706.702.706 1.195v4.62c0 .493-.269.949-.706 1.195l-4.088 2.31a1.438 1.438 0 0 1-1.412 0l-4.088-2.31A1.376 1.376 0 0 1 .5 8.31V3.69c0-.493.269-.949.706-1.195L5.294.185a1.438 1.438 0 0 1 1.412 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeverityHigh(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.713.295l4.992 4.992a1.008 1.008 0 0 1 0 1.426l-4.992 4.992a1.008 1.008 0 0 1-1.426 0L.295 6.713a1.008 1.008 0 0 1 0-1.426L5.287.295a1.008 1.008 0 0 1 1.426 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeverityInfo(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 6A6 6 0 1 1 0 6a6 6 0 0 1 12 0M6.75 3.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0M6 5.25a.75.75 0 0 0-.75.75v2.25a.75.75 0 1 0 1.5 0V6A.75.75 0 0 0 6 5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeverityLow(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="6" cy="6" r="6" fill="currentColor" fill-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeverityMedium(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.04 10.454L.24 3.182c-.398-.603-.29-1.457.24-1.91C.688 1.097.94 1 1.2 1h9.6c.663 0 1.2.61 1.2 1.364c0 .295-.084.582-.24.818l-4.8 7.272c-.398.603-1.15.725-1.68.273a1.295 1.295 0 0 1-.24-.273"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeverityUnknown(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 12A6 6 0 1 1 6 0a6 6 0 0 1 0 12M3.324 4.643c0-.313.107-.631.32-.953c.214-.322.526-.589.935-.8c.41-.212.887-.317 1.433-.317c.508 0 .956.088 1.344.265c.389.176.689.417.9.72c.212.304.318.635.318.991c0 .281-.06.527-.18.738a2.287 2.287 0 0 1-.431.548c-.167.153-.465.412-.895.775a3.549 3.549 0 0 0-.287.27a1.1 1.1 0 0 0-.16.213c-.289.667-1.543.592-1.302-.342a1.82 1.82 0 0 1 .363-.535a7.98 7.98 0 0 1 .609-.547c.224-.185.385-.325.485-.419c.1-.094.184-.199.252-.314a.727.727 0 0 0 .103-.377a.854.854 0 0 0-.313-.669c-.208-.181-.477-.272-.806-.272c-.385 0-.668.092-.85.275c-.182.183-.336.453-.462.81c-.12.373-.345.56-.677.56a.686.686 0 0 1-.496-.196c-.135-.13-.203-.272-.203-.424M6 9.75a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 1.5a2.5 2.5 0 1 0-2.469-2.104L5.298 6.263a2.5 2.5 0 1 0 0 3.475l4.733 2.366a2.5 2.5 0 1 0 .671-1.341L5.97 8.396a2.519 2.519 0 0 0 0-.792l4.733-2.367c.455.47 1.092.763 1.798.763Zm1 6.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M4.5 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 5.4V3.132l4.75-1.357v11.608l-1.782-1.528A8.5 8.5 0 0 1 2.5 5.401Zm6.25 7.982l1.782-1.528A8.5 8.5 0 0 0 13.5 5.401V3.13L8.75 1.774zM1 2l7-2l7 2v3.4a10 10 0 0 1-3.492 7.593L8 16l-3.508-3.007A10 10 0 0 1 1 5.401z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sidebar(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2.5h7a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5H6zm-1.5 0H3a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h1.5zM1 3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SidebarRight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 2.5H13a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5h-1.5zm-1.5 0H3a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h7zM1 3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.028 12.213c-2.325 0-3.388-1.184-3.388-2.053a.754.754 0 0 1 .795-.754c.996 0 .736 1.494 2.593 1.494c.95 0 1.508-.572 1.508-1.11c0-.322-.185-.69-.812-.838L6.65 8.429c-1.666-.424-1.957-1.351-1.957-2.21c0-1.785 1.632-2.43 3.186-2.43c1.43 0 3.132.79 3.132 1.86c0 .461-.385.708-.832.708c-.85 0-.708-1.194-2.418-1.194c-.85 0-1.298.398-1.298.956c0 .558.662.747 1.243.874l1.53.345c1.678.378 2.125 1.363 2.125 2.305c0 1.45-1.122 2.57-3.335 2.57zm6.422-3.067c.066-.381.1-.767.099-1.153a6.55 6.55 0 0 0-7.685-6.479A3.854 3.854 0 0 0 4.933 1a3.918 3.918 0 0 0-3.381 5.857a6.555 6.555 0 0 0 7.587 7.629a3.852 3.852 0 0 0 1.93.514a3.918 3.918 0 0 0 3.38-5.854Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlightFrown(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M6 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0M6.63 9.616a2.75 2.75 0 0 1 3.751 1.009a.75.75 0 0 1-1.299.75a1.25 1.25 0 0 0-2.163-.003a.75.75 0 1 1-1.297-.753a2.75 2.75 0 0 1 1.007-1.003Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlightSmile(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M6 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-4.37 4.384a2.749 2.749 0 0 0 3.751-1.009a.75.75 0 0 0-1.299-.75a1.25 1.25 0 0 1-2.163.003a.75.75 0 0 0-1.297.753c.242.417.59.763 1.007 1.003Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartCard(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 3.5H2a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5V4a.5.5 0 0 0-.5-.5M2 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm4 7h3v1a1 1 0 0 1-1 1H6zm-2 2h1V9H3v1a1 1 0 0 0 1 1m3-6h1a1 1 0 0 1 1 1v2H7zM4 5a1 1 0 0 0-1 1v2h3V5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smile(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M5.252 4.7a1.5 1.5 0 0 1 2.047.55a.5.5 0 0 1-.866.5a.5.5 0 0 0-.865-.001a.5.5 0 1 1-.865-.502a1.5 1.5 0 0 1 .55-.547ZM8 12a4 4 0 0 0 4-4H4a4 4 0 0 0 4 4m1.252-7.3a1.5 1.5 0 0 1 2.047.55a.5.5 0 1 1-.866.5a.5.5 0 0 0-.865-.001a.5.5 0 1 1-.865-.502a1.5 1.5 0 0 1 .55-.547Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smiley(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 8a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0M7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1 6a3 3 0 0 0 3-3H5a3 3 0 0 0 3 3m3-6a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snippet(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.625.1A.75.75 0 0 1 4.65.375L8 6.177L11.35.375a.75.75 0 1 1 1.3.75L8.864 7.677l1.97 3.412A2.503 2.503 0 0 1 14 13.5a2.5 2.5 0 1 1-4.425-1.595L7.999 9.176l-.26.45a.75.75 0 0 1-1.298-.751l.692-1.199L3.35 1.125A.75.75 0 0 1 3.625.1M5.5 13.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m1.5 0a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m5.5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoftUnwrap(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.75 4.25a.75.75 0 0 0 0 1.5h14.5a.75.75 0 0 0 0-1.5zm9 6a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5zm-4.56 0l-.97-.97a.75.75 0 0 1 1.06-1.06l2.25 2.25l.53.53l-.53.53l-2.25 2.25a.75.75 0 0 1-1.06-1.06l.97-.97H.75a.75.75 0 0 1 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoftWrap(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 5a.75.75 0 0 1 .75-.75h11.5a3.75 3.75 0 1 1 0 7.5H9.87l.97.97a.75.75 0 1 1-1.06 1.06l-2.25-2.25L7 11l.53-.53l2.25-2.25a.75.75 0 1 1 1.06 1.06l-.97.97h2.38a2.25 2.25 0 0 0 0-4.5H.75A.75.75 0 0 1 0 5m6 6a.75.75 0 0 0-.75-.75H.75a.75.75 0 0 0 0 1.5h4.5A.75.75 0 0 0 6 11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortHighest(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3 14l.53-.53l2.25-2.25a.75.75 0 0 0-1.06-1.061l-.97.97v-8.38a.75.75 0 1 0-1.5 0v8.38l-.97-.97a.75.75 0 0 0-1.06 1.06l2.25 2.25zM8.75 2a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5zm0 3.25a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5zm-.75 4a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5A.75.75 0 0 1 8 9.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortLowest(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3 2l.53.53l2.25 2.25a.75.75 0 0 1-1.06 1.061l-.97-.97v8.38a.75.75 0 0 1-1.5 0V4.87l-.97.97A.75.75 0 0 1 .22 4.78l2.25-2.25zm5 4.75c0 .414.336.75.75.75h1.5a.75.75 0 0 0 0-1.5h-1.5a.75.75 0 0 0-.75.75m.75 4a.75.75 0 0 1 0-1.5h4.5a.75.75 0 0 1 0 1.5zm0 3.25a.75.75 0 0 1 0-1.5h6.5a.75.75 0 0 1 0 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spam(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.103 1.529L8 0L5.897 1.529l-2.6-.001L2.496 4L.392 5.528L1.196 8l-.804 2.472L2.495 12l.803 2.473l2.6-.001L8 16l2.103-1.529l2.6.001l.802-2.473l2.103-1.527L14.804 8l.804-2.472L13.505 4l-.803-2.473zm1.51 1.5H9.614l-.395-.287L8 1.855l-1.22.887l-.395.287H4.387l-.465 1.435l-.15.464l-.395.287l-1.222.886l.467 1.435l.151.464l-.15.464L2.154 9.9l1.222.886l.395.287l.15.464l.466 1.436l1.509-.001h.488l.395.287l1.22.887l1.22-.887l.395-.287h1.998l.465-1.435l.15-.464l.395-.287l1.222-.886l-.467-1.435l-.15-.465l.15-.464l.467-1.435l-1.22-.886l-.396-.287l-.15-.464l-.466-1.436ZM9 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25-6.25a.75.75 0 0 0-1.5 0v3.5a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StageAll(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.75 3a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5zm0 4.25a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5zm-.75 5a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1-.75-.75M6.28 7.28a.75.75 0 0 0-1.06-1.06L2.5 8.94L1.28 7.72A.75.75 0 0 0 .22 8.78l1.75 1.75a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.454 1.694a.591.591 0 0 1 1.092 0l1.585 3.81a.25.25 0 0 0 .21.154l4.114.33a.591.591 0 0 1 .338 1.038L11.658 9.71a.25.25 0 0 0-.08.247l.957 4.015a.591.591 0 0 1-.883.641l-3.522-2.15a.25.25 0 0 0-.26 0l-3.522 2.15a.591.591 0 0 1-.883-.641l.957-4.015a.25.25 0 0 0-.08-.247L1.207 7.026a.591.591 0 0 1 .338-1.038l4.113-.33a.25.25 0 0 0 .211-.153z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarO(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.189 2.332l-.001.003l-1.319 3.17a.25.25 0 0 1-.21.153l-3.423.274h-.003l-.095.008l-.593.048a.591.591 0 0 0-.338 1.038l.452.387l.073.062l.002.002l2.608 2.234a.25.25 0 0 1 .08.247l-.796 3.34l-.001.003l-.022.093l-.138.579a.591.591 0 0 0 .883.641l.507-.31l.082-.05l.003-.001l2.93-1.79a.25.25 0 0 1 .26 0l2.93 1.79l.003.002l.082.05l.507.31a.591.591 0 0 0 .883-.642l-.138-.579l-.022-.093v-.003l-.797-3.34a.25.25 0 0 1 .08-.247l2.608-2.234l.002-.002l.073-.062l.452-.387a.591.591 0 0 0-.338-1.038l-.593-.048l-.095-.008h-.003l-3.422-.274a.25.25 0 0 1-.211-.153l-1.319-3.17l-.001-.003l-.037-.089l-.228-.549a.591.591 0 0 0-1.092 0l-.228.55zM8 4.288L7.254 6.08a1.75 1.75 0 0 1-1.476 1.072l-1.935.155L5.317 8.57a1.75 1.75 0 0 1 .564 1.736l-.45 1.888l1.657-1.012a1.75 1.75 0 0 1 1.824 0l1.657 1.012l-.45-1.889a1.75 1.75 0 0 1 .564-1.735l1.474-1.263l-1.935-.155A1.75 1.75 0 0 1 8.746 6.08z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Status(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.941 7.033a8 8 0 0 1-14.784 5.112a.75.75 0 1 1 1.283-.778a6.5 6.5 0 1 0 8.922-8.93a.75.75 0 0 1 .776-1.284a8 8 0 0 1 3.803 5.88M9 1a1 1 0 1 1-2 0a1 1 0 0 1 2 0M2.804 5a1 1 0 1 0-1.732-1a1 1 0 0 0 1.732 1M1 7a1 1 0 1 1 0 2a1 1 0 0 1 0-2m4-4.196a1 1 0 1 0-1-1.732a1 1 0 0 0 1 1.732" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusActive(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="6" cy="6" r="6" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusAlert(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 6A6 6 0 1 1 0 6a6 6 0 0 1 12 0M5 3a1 1 0 0 1 2 0v3a1 1 0 0 1-2 0zm1 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusCancelled(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 12A6 6 0 1 0 6 0a6 6 0 0 0 0 12M4.28 3.22a.75.75 0 0 0-1.06 1.06l4.5 4.5a.75.75 0 0 0 1.06-1.06z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusClosed(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.536 8.657l2.828-2.83a1 1 0 0 1 1.414 1.416l-3.535 3.535a1 1 0 0 1-1.415.001l-2.12-2.12a1 1 0 1 1 1.413-1.415zM8 16A8 8 0 1 1 8 0a8 8 0 0 1 0 16m0-2A6 6 0 1 0 8 2a6 6 0 0 0 0 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusHealth(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.999 1a.75.75 0 0 1 .715.521L12 11.79l1.286-4.018A.75.75 0 0 1 14 7.25h1.25a.75.75 0 0 1 0 1.5h-.703l-1.833 5.729a.75.75 0 0 1-1.428 0L8.005 4.226l-2.29 7.25a.75.75 0 0 1-1.42.03L3.031 8.03l-.07.208a.75.75 0 0 1-.711.513H.75a.75.75 0 0 1 0-1.5h.96l.578-1.737a.75.75 0 0 1 1.417-.02L4.95 8.919l2.335-7.394A.75.75 0 0 1 7.999 1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusNeutral(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 12A6 6 0 1 0 6 0a6 6 0 0 0 0 12M3 4.75A.75.75 0 0 1 3.75 4h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 3 4.75m0 2.5a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5A.75.75 0 0 1 3 7.25" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusPaused(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 6A6 6 0 1 1 0 6a6 6 0 0 1 12 0M3 3.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5zM7.5 3a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusPreparingBorderless(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><circle cx="8" cy="8" r="2"/><circle cx="14" cy="8" r="2"/><circle cx="2" cy="8" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusStopped(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 12A6 6 0 1 0 6 0a6 6 0 0 0 0 12M3.5 3a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatusWaiting(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 10a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0 2A6 6 0 1 0 6 0a6 6 0 0 0 0 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="10" height="10" x="3" y="3" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.886 1a3.136 3.136 0 0 0-2.41 5.144L6.4 7.25H2.75a.75.75 0 0 0 0 1.5h4.899l1.722 2.066A1.636 1.636 0 0 1 8.114 13.5H8a1.75 1.75 0 0 1-1.75-1.75a.75.75 0 0 0-1.5 0A3.25 3.25 0 0 0 8 15h.114a3.136 3.136 0 0 0 2.41-5.144L9.6 8.75h3.649a.75.75 0 0 0 0-1.5H8.351L6.63 5.184A1.636 1.636 0 0 1 7.886 2.5H8c.966 0 1.75.784 1.75 1.75a.75.75 0 0 0 1.5 0A3.25 3.25 0 0 0 8 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subgroup(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m1.5 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0m7 4a3 3 0 1 1-6 0a3 3 0 0 1 6 0m-10 5.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0 1.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subscript(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.22 4.22a.75.75 0 0 1 1.06 0L6 6.94l2.72-2.72a.75.75 0 0 1 1.06 1.06L7.06 8l2.72 2.72a.75.75 0 1 1-1.06 1.06L6 9.06l-2.72 2.72a.75.75 0 0 1-1.06-1.06L4.94 8L2.22 5.28a.75.75 0 0 1 0-1.06m10.407 6.782H11V10h1.627c.362 0 .71.144.967.402a1.392 1.392 0 0 1-.17 2.11l-.675.486h1.232V14H11v-.974l1.847-1.33a.385.385 0 0 0 .047-.583a.378.378 0 0 0-.267-.111" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Substitute(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m0 5l.53-.53l2.25-2.25a.75.75 0 0 1 1.06 1.06l-.969.97h10.38a.75.75 0 0 1 0 1.5H2.87l.97.97a.75.75 0 0 1-1.06 1.06L.53 5.53zm16 6l-.53-.53l-2.25-2.25a.75.75 0 1 0-1.061 1.06l.97.97H2.748a.75.75 0 0 0 0 1.5h10.38l-.97.97a.75.75 0 1 0 1.06 1.06l2.25-2.25z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Superscript(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.627 3.002H11V2h1.627c.362 0 .71.144.967.402a1.392 1.392 0 0 1-.17 2.11l-.675.486h1.232V6H11v-.974l1.847-1.33a.385.385 0 0 0 .047-.583a.378.378 0 0 0-.267-.111M2.22 4.22a.75.75 0 0 1 1.06 0L6 6.94l2.72-2.72a.75.75 0 0 1 1.06 1.06L7.06 8l2.72 2.72a.75.75 0 1 1-1.06 1.06L6 9.06l-2.72 2.72a.75.75 0 0 1-1.06-1.06L4.94 8L2.22 5.28a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Symlink(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12V9H6a3 3 0 0 0-3 3v2a5 5 0 0 1 3-9h4V2l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4.5v-2h11v2zm0 1.5v3h4.75V6zm6.25 0v3h4.75V6zM2.5 10.5h4.75v3H2.5zm6.25 0v3h4.75v-3zM1 2a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.5H4a.5.5 0 0 0-.5.5v1h9V2a.5.5 0 0 0-.5-.5m.5 3h-9V14a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5zM4 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm1 12.25a.75.75 0 0 1 .75-.75h4.75a.75.75 0 0 1 0 1.5H5.75a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tachometer(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.545 2.1a.75.75 0 0 1 .274 1.025l-3.472 6.007a3 3 0 1 1-1.208-.908l1.759-3.042a6.5 6.5 0 0 0-2.148-.639V5a.75.75 0 1 1-1.5 0v-.457a6.5 6.5 0 0 0-1.829.49l.229.396a.75.75 0 1 1-1.3.75l-.228-.396a6.5 6.5 0 0 0-1.339 1.34l.396.227a.75.75 0 0 1-.75 1.3l-.396-.229a6.5 6.5 0 0 0-.498 1.905a.75.75 0 0 1-1.492-.155A8 8 0 0 1 11.65 3.88l.87-1.506a.75.75 0 0 1 1.025-.274m-.107 4.047a.75.75 0 0 1 1.047.169a8 8 0 0 1 1.51 4.963a.75.75 0 1 1-1.499-.052a6.5 6.5 0 0 0-1.227-4.033a.75.75 0 0 1 .17-1.047ZM9.5 11a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.172 5.5a.5.5 0 0 1 .353.146l1.06-1.06l-1.06 1.06L13.88 8l-2.354 2.354a.5.5 0 0 1-.353.146H2a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 .5-.5zm3.767 1.44l-2.353-2.354A2 2 0 0 0 11.172 4H2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h9.172a2 2 0 0 0 1.414-.586l2.353-2.353L16 8zm-8.189.31a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tanuki(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.329 6.922l-.02-.053l-1.239-3.232l-.668 2.044l-.337 1.035H3.936L3.598 5.68l-.667-2.043l-1.24 3.237l-.005.01l-.004.01l-.018.045a2.543 2.543 0 0 0 .848 2.93l.005.004l.01.008l.01.008l.009.006l3.302 2.474l.004.004l1.644 1.242l.5.379l.501-.379l1.644-1.242l.006-.004l3.325-2.489l.007-.005l.009-.007a2.543 2.543 0 0 0 .844-2.939l-.003-.008Zm.053 4.152l-3.337 2.497l-1.643 1.242l-.999.755a.675.675 0 0 1-.813 0l-.999-.755l-1.643-1.242l-3.313-2.483l-.018-.014l-.008-.006A4.043 4.043 0 0 1 .267 6.395l.023-.057l1.9-4.963l.013-.033l.081-.212l.023-.06l.158-.413A.565.565 0 0 1 2.941.3a.576.576 0 0 1 .613.415l.134.41l.005.017l.078.239l.011.033l1.242 3.802h5.953l1.241-3.802l.011-.033l.078-.24l.006-.016l.133-.41a.575.575 0 0 1 1.088-.058l.157.409l.025.064l.08.211l.013.033l1.903 4.963l.022.058a4.043 4.043 0 0 1-1.343 4.671z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TanukiAi(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.75.75a.75.75 0 0 0-1.5 0v.5h-.5a.75.75 0 0 0 0 1.5h.5v.5a.75.75 0 0 0 1.5 0v-.5h.5a.75.75 0 0 0 0-1.5h-.5zm11.645 6.118l-.014-.038l-.55-1.44l-.172.53l-.337 1.035H6.678L6.342 5.92l-.172-.528l-.551 1.445l-.004.01l-.004.01l-.012.03a1.557 1.557 0 0 0 .512 1.78l.001.001l.003.003l.01.008l.01.008l.004.002l2.474 1.862l.005.004l1.232.935l.148.112l.146-.112l1.233-.935l.006-.004l2.49-1.873l.006-.004l.009-.006a1.558 1.558 0 0 0 .51-1.79zm.391 3l-2.502 1.88l-1.233.936l-.748.569a.505.505 0 0 1-.61 0l-.75-.569l-1.232-.935l-2.485-1.87l-.013-.011l-.007-.005A3.057 3.057 0 0 1 4.2 6.343l.017-.042L5.431 3.12l.068-.178l.043-.114l.173-.453l.017-.045l.118-.308a.425.425 0 0 1 .356-.27a.43.43 0 0 1 .46.314l.099.305l.004.013l.193.594l.06.181l.746 2.296h4.464l.747-2.296l.06-.18l.193-.595l.004-.013l.099-.305a.433.433 0 0 1 .648-.248a.43.43 0 0 1 .168.204l.116.305l.019.048l.171.45l.045.117l.068.177L15.784 6.3l.016.044a3.057 3.057 0 0 1-1.007 3.518zM3.5 11a.75.75 0 0 1 .75.75v1h1a.75.75 0 0 1 0 1.5h-1v1a.75.75 0 0 1-1.5 0v-1h-1a.75.75 0 0 1 0-1.5h1v-1A.75.75 0 0 1 3.5 11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TanukiVerified(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.712 6.337l.022.058a4.043 4.043 0 0 1-1.343 4.671l-.01.008l-3.336 2.497l-1.643 1.242l-.999.755a.675.675 0 0 1-.813 0l-.999-.755l-1.643-1.242l-3.313-2.483l-.018-.014l-.008-.007A4.043 4.043 0 0 1 .267 6.395l.023-.057L2.466.657a.565.565 0 0 1 .475-.358a.576.576 0 0 1 .613.416l1.47 4.5h5.953l1.47-4.5a.575.575 0 0 1 1.087-.058zM11.28 8.53a.75.75 0 0 0-1.06-1.06L7.5 10.19L6.28 8.97a.75.75 0 1 0-1.06 1.06l1.75 1.75a.75.75 0 0 0 1.06 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TaskDone(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 13.5a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h9.25a.75.75 0 0 0 0-1.5H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5zm12.78-8.82a.75.75 0 0 0-1.06-1.06L9.162 9.177L7.289 7.241a.75.75 0 1 0-1.078 1.043l2.403 2.484a.75.75 0 0 0 1.07.01z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Template(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2.5h-11V5h11zm0 4h-7v7h7zM5 6.5H2.5v7H5zM2.5 1H1v14h14V1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 3.5H2a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5V4a.5.5 0 0 0-.5-.5M2 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm5 8.25a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75M4.28 5.22a.75.75 0 0 0-1.06 1.06L4.94 8L3.22 9.72a.75.75 0 1 0 1.06 1.06l2.25-2.25l.53-.53l-.53-.53z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terraform(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.321 7.576L1 5.051V0l4.321 2.525zM10.117 16l-4.321-2.525V8.424l4.321 2.525zM5.796 2.819l4.321 2.528v5.048L5.796 7.869V2.82Zm4.796 7.576l4.321-2.523V2.819l-4.321 2.528z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextDescription(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 3.75A.75.75 0 0 1 .75 3h14.5a.75.75 0 0 1 0 1.5H.75A.75.75 0 0 1 0 3.75M0 8a.75.75 0 0 1 .75-.75h14.5a.75.75 0 0 1 0 1.5H.75A.75.75 0 0 1 0 8m.75 3.5a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbDown(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.72 13.402l-.27-.902L10 11h3.989a2 2 0 0 0 1.932-2.517l-1.19-4.448A4 4 0 0 0 9.852 1.2L3 3H0v8h3l4.39 4.39c1.47 1.47 3.928.002 3.33-1.988M3 9.5H1.5v-5H3zm7.232-6.85L4.5 4.157v6.222l3.951 3.951c.12.119.22.149.296.155a.533.533 0 0 0 .314-.08a.533.533 0 0 0 .22-.238a.466.466 0 0 0 .003-.334l-.72-2.402l-.58-1.931h6.005a.5.5 0 0 0 .483-.63l-1.19-4.447a2.5 2.5 0 0 0-3.05-1.772Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbUp(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.72 2.598l-.27.902L10 5h3.989a2 2 0 0 1 1.932 2.517l-1.19 4.448a4 4 0 0 1-4.88 2.835L3 13H0V5h3L7.39.61c1.47-1.47 3.928-.002 3.33 1.988M3 6.5H1.5v5H3zm7.232 6.85L4.5 11.843V5.621L8.451 1.67a.466.466 0 0 1 .296-.155a.533.533 0 0 1 .314.08c.11.065.183.154.22.238c.03.07.051.173.003.334l-.72 2.402l-.58 1.931h6.005a.5.5 0 0 1 .483.63l-1.19 4.447a2.5 2.5 0 0 1-3.05 1.773" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thumbtack(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.28 1.22a.75.75 0 0 0-1.26.7L6.69 5.25H4.206c-1.114 0-1.671 1.346-.884 2.134l1.911 1.911l-3.72 4.135A2 2 0 0 0 1 14.768V15h.233a2 2 0 0 0 1.337-.513l4.135-3.721l1.911 1.91c.788.788 2.134.23 2.134-.883V9.31l3.33-3.33a.75.75 0 0 0 .7-1.261l-.603-.604l-2.293-2.293zM12.94 5L11 3.06L8.06 6L10 7.94zM6.69 6.75l2.56 2.56v1.88L4.81 6.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbtackSolid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.28 1.22a.75.75 0 0 0-1.26.7L6.69 5.25H4.206c-1.114 0-1.671 1.346-.884 2.134l1.911 1.911l-3.72 4.135A2 2 0 0 0 1 14.768V15h.233a2 2 0 0 0 1.337-.513l4.135-3.721l1.911 1.91c.788.788 2.134.23 2.134-.883V9.31l3.33-3.33a.75.75 0 0 0 .7-1.261l-.603-.604l-2.293-2.293z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeOut(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.229.199a8 8 0 0 1 9.727 6.964a.75.75 0 0 1-1.492.157a6.5 6.5 0 1 0-7.132 7.146a.75.75 0 1 1-.154 1.492a8 8 0 0 1-.95-15.76ZM8 3a.75.75 0 0 1 .75.75V9h-4a.75.75 0 0 1 0-1.5h2.5V3.75A.75.75 0 0 1 8 3m2.22 7.22a.75.75 0 0 1 1.06 0L13 11.94l1.72-1.72a.75.75 0 1 1 1.06 1.06L14.06 13l1.72 1.72a.75.75 0 1 1-1.06 1.06L13 14.06l-1.72 1.72a.75.75 0 1 1-1.06-1.06L11.94 13l-1.72-1.72a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 1a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2v.07a6.958 6.958 0 0 1 2.812 1.058l.908-.908a.75.75 0 1 1 1.06 1.06l-.8.8A7 7 0 1 1 7 2.07V2a1 1 0 0 1-1-1m7.5 8a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0M8.75 5.75a.75.75 0 0 0-1.5 0v3.56l.22.22l2.254 2.254a.75.75 0 1 0 1.06-1.06L8.75 8.689z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Title(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.75 2a.75.75 0 0 0-.75.75v1.5a.75.75 0 0 0 1.5 0V3.5h2.75v9h-.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-.5v-9H8.5v.75a.75.75 0 0 0 1.5 0v-1.5A.75.75 0 0 0 9.25 2zM8 7.75A.75.75 0 0 1 8.75 7h6.5a.75.75 0 0 1 0 1.5h-6.5A.75.75 0 0 1 8 7.75m0 3.5a.75.75 0 0 1 .75-.75h6.5a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1-.75-.75" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TodoAdd(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 6a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5V.75a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 13 6M3 13.5a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h4.25a.75.75 0 0 0 0-1.5H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V8.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TodoDone(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 13.5a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h9.25a.75.75 0 0 0 0-1.5H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5zm12.78-8.82a.75.75 0 0 0-1.06-1.06L9.162 9.177L7.289 7.241a.75.75 0 1 0-1.078 1.043l2.403 2.484a.75.75 0 0 0 1.07.01z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Token(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 4.5h6a3.5 3.5 0 1 1 0 7H5a3.5 3.5 0 1 1 0-7M0 8a5 5 0 0 1 5-5h6a5 5 0 0 1 0 10H5a5 5 0 0 1-5-5m11 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2M9 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0M5 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendDown(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.143 9.847a1 1 0 0 0 1.715 0l3.999-6.665a1 1 0 0 0-.858-1.515H2.001a1 1 0 0 0-.858 1.515z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendStatic(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.847 5.142a1 1 0 0 1 0 1.715l-6.665 4a1 1 0 0 1-1.515-.858V2.001a1 1 0 0 1 1.515-.858z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendUp(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.143 2.153a1 1 0 0 1 1.715 0l3.999 6.665a1 1 0 0 1-.858 1.515H2.001a1 1 0 0 1-.858-1.515z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriggerSource(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.25.75a.75.75 0 0 1 1.5 0V2c0 2.9 2.35 5.25 5.25 5.25h1.25a.75.75 0 0 1 0 1.5H14A5.25 5.25 0 0 0 8.75 14v1.25a.75.75 0 0 1-1.5 0V14c0-2.9-2.35-5.25-5.25-5.25H.75a.75.75 0 0 1 0-1.5H2C4.9 7.25 7.25 4.9 7.25 2zM5.095 8A6.78 6.78 0 0 1 8 10.905A6.78 6.78 0 0 1 10.905 8A6.78 6.78 0 0 1 8 5.095A6.78 6.78 0 0 1 5.095 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unapproval(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M7 5a2.99 2.99 0 0 1-.87 2.113A3.997 3.997 0 0 1 8 10.5V12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-1.5c0-1.427.747-2.679 1.87-3.387A3 3 0 1 1 7 5m-5.5 5.5a2.5 2.5 0 0 1 5 0V12a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5zm9.72-6.28a.75.75 0 0 1 1.06 0l1.22 1.22l1.22-1.22a.75.75 0 1 1 1.06 1.06L14.56 6.5l1.22 1.22a.75.75 0 0 1-1.06 1.06L13.5 7.56l-1.22 1.22a.75.75 0 1 1-1.06-1.06l1.22-1.22l-1.22-1.22a.75.75 0 0 1 0-1.06" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unassignee(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M7 5a2.99 2.99 0 0 1-.87 2.113A3.997 3.997 0 0 1 8 10.5V12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-1.5c0-1.427.747-2.679 1.87-3.387A3 3 0 1 1 7 5m-5.5 5.5a2.5 2.5 0 0 1 5 0V12a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5zm9.25-4.25a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 2.75a.75.75 0 0 0-1.5 0V7a4 4 0 1 0 8 0V2.75a.75.75 0 0 0-1.5 0V7a2.5 2.5 0 0 1-5 0zM3.75 12.5a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlink(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 1.75a.75.75 0 0 0-1.5 0V4.5H1.75a.75.75 0 0 0 0 1.5H6zm5.151 1.16a1.371 1.371 0 0 1 1.94 1.939l-2.871 2.87a.75.75 0 1 0 1.06 1.061l2.871-2.87a2.871 2.871 0 0 0-4.06-4.061L7.22 4.719A.75.75 0 0 0 8.28 5.78zM10 14.25a.75.75 0 0 0 1.5 0V11.5h2.75a.75.75 0 0 0 0-1.5H10zM5.781 7.22a.75.75 0 0 1 0 1.06l-2.87 2.871a1.371 1.371 0 0 0 1.939 1.94l2.87-2.871a.75.75 0 1 1 1.061 1.06l-2.87 2.871a2.871 2.871 0 0 1-4.061-4.06L4.72 7.22a.75.75 0 0 1 1.061 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnstageAll(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 3.75a.75.75 0 0 1 1.5 0v.943a5 5 0 1 1 .67 7.246a.75.75 0 1 1 .924-1.182A3.5 3.5 0 1 0 2.088 6.5H4.25a.75.75 0 0 1 0 1.5H0zm12 0a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75M12 8a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 12 8m.75 3.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upgrade(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14 6l-1.5-1.5l-3.44-3.44L8 0L6.94 1.06L3.5 4.5L2 6h3v7h6V6zM8 2.121L10.379 4.5H9.5v7h-3v-7h-.879zM5.75 14.5a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 5.841a.75.75 0 0 1-1.06 0l-1.97-1.97v7.379a.75.75 0 0 1-1.5 0V3.871l-1.97 1.97a.75.75 0 0 1-1.06-1.06l3.25-3.25L8 1l.53.53l3.25 3.25a.75.75 0 0 1 0 1.061M2.5 9.75a.75.75 0 0 0-1.5 0V13a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m.514 2.63a4 4 0 1 0-6.028 0A4.002 4.002 0 0 0 2 11.5V13a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-1.5a4.002 4.002 0 0 0-2.986-3.87M8 9H6a2.5 2.5 0 0 0-2.5 2.5V13a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5v-1.5A2.5 2.5 0 0 0 10 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m.63 2.113a3 3 0 1 0-4.259 0A3.997 3.997 0 0 0 1 9.5V13a2 2 0 0 0 2 2h4c.597 0 1.134-.262 1.5-.677c.366.415.903.677 1.5.677h3a2 2 0 0 0 2-2v-2c0-1.218-.622-2.29-1.565-2.917a2.5 2.5 0 1 0-3.87 0c-.241.16-.462.35-.656.564a4.005 4.005 0 0 0-1.78-2.534ZM5 7a2.5 2.5 0 0 0-2.5 2.5V13a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V9.5A2.5 2.5 0 0 0 5 7m7.5-.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1 2.5a2 2 0 0 0-2 2v2a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-2a2 2 0 0 0-2-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 5H3l3.364-2.77A1 1 0 0 1 8 3.002v9.996a1 1 0 0 1-1.636.772L3 11H0V5zM3 9.5H1.5v-3h2.038l.416-.342L6.5 4.061v7.878L3.954 9.842L3.538 9.5zm9.457-7.267a.75.75 0 0 1 1.06-.026a7.999 7.999 0 0 1 0 11.586a.75.75 0 0 1-1.034-1.086a6.5 6.5 0 0 0 0-9.414a.75.75 0 0 1-.026-1.06m-2.292 2.034a.75.75 0 0 1 1.057-.09a5 5 0 0 1 .001 7.645a.75.75 0 1 1-.967-1.146a3.5 3.5 0 0 0 0-5.353a.75.75 0 0 1-.091-1.056" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.429 2.746a.5.5 0 0 0-.858 0L1.58 12.743a.5.5 0 0 0 .429.757h11.984a.5.5 0 0 0 .43-.757zm-2.144-.77C7.06.68 8.939.68 9.715 1.975l5.993 9.996c.799 1.333-.161 3.028-1.716 3.028H2.008C.453 15-.507 13.305.292 11.972l5.993-9.997ZM9 11.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-.25-5.75a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningSolid(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.285 1.975C7.06.68 8.939.68 9.715 1.975l5.993 9.997c.799 1.333-.161 3.028-1.716 3.028H2.008C.453 15-.507 13.305.292 11.972zM8 5a.75.75 0 0 1 .75.75v3a.75.75 0 0 1-1.5 0v-3A.75.75 0 0 1 8 5m1 6.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weight(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.25 3.75a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0M10.45 5a2.75 2.75 0 1 0-4.9 0H3.82a1 1 0 0 0-.98.804l-1.6 8A1 1 0 0 0 2.22 15h11.56a1 1 0 0 0 .98-1.196l-1.6-8A1 1 0 0 0 12.18 5zM8 6.5H4.23l-1.4 7h10.34l-1.4-7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Work(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 1a1.75 1.75 0 0 0-1.75 1.75V4H3a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-1.25V2.75A1.75 1.75 0 0 0 10 1zm4.25 3V2.75A.25.25 0 0 0 10 2.5H6a.25.25 0 0 0-.25.25V4zM3 5.5h10a.5.5 0 0 1 .5.5v1h-11V6a.5.5 0 0 1 .5-.5m-.5 3V13a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V8.5H9V10H7V8.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func X(children ...ElementRenderer) *PajamasIcon {
	return &PajamasIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.294 6.928L14.357 1h-1.2L8.762 6.147L5.25 1H1.2l5.31 7.784L1.2 15h1.2l4.642-5.436L10.751 15h4.05zM7.651 8.852l-.538-.775L2.832 1.91h1.843l3.454 4.977l.538.775l4.491 6.47h-1.843z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
