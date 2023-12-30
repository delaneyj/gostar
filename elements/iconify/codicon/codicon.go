package codicon

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "0.0.35"
	hAttr          = "1em"
	viewbox        = "0 0 0 0"
	fill           = "currentColor"
)

type CodiconIcon struct {
	*SVGSVGElement
}

func Account(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7.992C16 3.58 12.416 0 8 0S0 3.58 0 7.992c0 2.43 1.104 4.62 2.832 6.09c.016.016.032.016.032.032c.144.112.288.224.448.336c.08.048.144.111.224.175A7.98 7.98 0 0 0 8.016 16a7.98 7.98 0 0 0 4.48-1.375c.08-.048.144-.111.224-.16c.144-.111.304-.223.448-.335c.016-.016.032-.016.032-.032c1.696-1.487 2.8-3.676 2.8-6.106m-8 7.001c-1.504 0-2.88-.48-4.016-1.279c.016-.128.048-.255.08-.383a4.17 4.17 0 0 1 .416-.991c.176-.304.384-.576.64-.816c.24-.24.528-.463.816-.639c.304-.176.624-.304.976-.4A4.15 4.15 0 0 1 8 10.342a4.185 4.185 0 0 1 2.928 1.166c.368.368.656.8.864 1.295c.112.288.192.592.24.911A7.03 7.03 0 0 1 8 14.993m-2.448-7.4a2.49 2.49 0 0 1-.208-1.024c0-.351.064-.703.208-1.023c.144-.32.336-.607.576-.847c.24-.24.528-.431.848-.575c.32-.144.672-.208 1.024-.208c.368 0 .704.064 1.024.208c.32.144.608.336.848.575c.24.24.432.528.576.847c.144.32.208.672.208 1.023c0 .368-.064.704-.208 1.023a2.84 2.84 0 0 1-.576.848a2.84 2.84 0 0 1-.848.575a2.715 2.715 0 0 1-2.064 0a2.84 2.84 0 0 1-.848-.575a2.526 2.526 0 0 1-.56-.848zm7.424 5.306c0-.032-.016-.048-.016-.08a5.22 5.22 0 0 0-.688-1.406a4.883 4.883 0 0 0-1.088-1.135a5.207 5.207 0 0 0-1.04-.608a2.82 2.82 0 0 0 .464-.383a4.2 4.2 0 0 0 .624-.784a3.624 3.624 0 0 0 .528-1.934a3.71 3.71 0 0 0-.288-1.47a3.799 3.799 0 0 0-.816-1.199a3.845 3.845 0 0 0-1.2-.8a3.72 3.72 0 0 0-1.472-.287a3.72 3.72 0 0 0-1.472.288a3.631 3.631 0 0 0-1.2.815a3.84 3.84 0 0 0-.8 1.199a3.71 3.71 0 0 0-.288 1.47c0 .352.048.688.144 1.007c.096.336.224.64.4.927c.16.288.384.544.624.784c.144.144.304.271.48.383a5.12 5.12 0 0 0-1.04.624c-.416.32-.784.703-1.088 1.119a4.999 4.999 0 0 0-.688 1.406c-.016.032-.016.064-.016.08C1.776 11.636.992 9.91.992 7.992C.992 4.14 4.144.991 8 .991s7.008 3.149 7.008 7.001a6.96 6.96 0 0 1-2.032 4.907"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ActivateBreakpoints(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 5.5a4.394 4.394 0 0 1-4 4.5a2.955 2.955 0 0 0-.2-1A3.565 3.565 0 0 0 14 5.5a3.507 3.507 0 0 0-7-.3A3.552 3.552 0 0 0 6 5a4.622 4.622 0 0 1 4.5-4A4.481 4.481 0 0 1 15 5.5M5.5 6a4.5 4.5 0 1 0 0 9.001a4.5 4.5 0 0 0 0-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Add(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7v1H8v6H7V8H1V7h6V1h1v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 1h-13l-.5.5v3l.5.5H2v8.5l.5.5h11l.5-.5V5h.5l.5-.5v-3zm-1 3H2V2h12v2zM3 13V5h10v8zm8-6H5v1h6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBoth(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3 9l2.146 2.146l-.707.708l-3-3v-.708l3-3l.707.708L3 8h10l-2.146-2.146l.707-.708l3 3v.708l-3 3l-.707-.707L13 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDown(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m5.369 8.08l2.14 2.14V4.468h1v5.68l2.066-2.067l.707.707l-2.957 2.956h-.707L4.662 8.788z"/><path d="M14 8A6 6 0 1 0 2 8a6 6 0 0 0 12 0m-1 0A5 5 0 1 1 3 8a5 5 0 0 1 10 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m7.92 10.631l-2.14-2.14h5.752v-1h-5.68L7.92 5.426l-.707-.707l-2.956 2.957v.707l2.956 2.956z"/><path d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2m0 1a5 5 0 1 0 0 10A5 5 0 0 0 8 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRight(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m8.08 5.369l2.14 2.14H4.468v1h5.68L8.08 10.574l.707.707l2.956-2.957v-.707L8.788 4.662z"/><path d="M8 14A6 6 0 1 1 8 2a6 6 0 0 1 0 12m0-1A5 5 0 1 0 8 3a5 5 0 0 0 0 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUp(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m5.369 7.92l2.14-2.14v5.752h1v-5.68l2.066 2.067l.707-.707l-2.957-2.956h-.707L4.662 7.212z"/><path d="M14 8A6 6 0 1 1 2 8a6 6 0 0 1 12 0m-1 0A5 5 0 1 0 3 8a5 5 0 0 0 10 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.147 9l5 5h.707l5-5l-.707-.707L9 12.439V2H8v10.44L3.854 8.292z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7 3.093l-5 5V8.8l5 5l.707-.707l-4.146-4.147H14v-1H3.56L7.708 3.8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9 13.887l5-5V8.18l-5-5l-.707.707l4.146 4.147H2v1h10.44L8.292 13.18z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallDown(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.7 8.64l-2.5 2.5h-.7L5 8.64l.7-.71l1.65 1.64V4h1v5.57L10 7.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallLeft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 10.7L4 8.2v-.7L6.5 5l.71.7l-1.64 1.65h5.57v1H5.57L7.22 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallRight(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.64 5l2.5 2.5v.7l-2.5 2.5l-.71-.7l1.64-1.65H4v-1h5.57L7.92 5.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSmallUp(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 6.5L7.5 4h.7l2.5 2.5l-.7.71l-1.65-1.64v5.57h-1V5.57L5.7 7.22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSwap(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.207 15.061L1 11.854v-.707L4.207 7.94l.707.707l-2.353 2.354H15v1H2.56l2.354 2.353zm7.586-7L15 4.854v-.707L11.793.94l-.707.707L13.439 4H1v1h12.44l-2.354 2.354z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.854 7l-5-5h-.707l-5 5l.707.707L8 3.561V14h1V3.56l4.146 4.147z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Azure(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.37 13.68l-4-12a1 1 0 0 0-1-.68H5.63a1 1 0 0 0-.95.68l-4.05 12a1 1 0 0 0 1 1.32h2.93a1 1 0 0 0 .94-.68l.61-1.78l3 2.27a1 1 0 0 0 .6.19h4.68a1 1 0 0 0 .98-1.32m-5.62.66a.32.32 0 0 1-.2-.07L3.9 10.08l-.09-.07h3l.08-.21l1-2.53l2.24 6.63a.34.34 0 0 1-.38.44m4.67 0H10.7a1 1 0 0 0 0-.66l-4.05-12h3.72a.34.34 0 0 1 .32.23l4.05 12a.34.34 0 0 1-.32.43" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AzureDevops(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3.622v8.512L11.5 15l-5.425-1.975v1.957l-3.071-4.01l8.951.698V4.006zm-2.984.428L6.994 1v2.001L2.383 4.356L1 6.13v4.029l1.978.874V5.868z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beaker(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.893 13.558L10 6.006v-4h1v-1H9.994V1l-.456.005H5V2h1v3.952l-3.894 7.609A1 1 0 0 0 3 15.006h10a1 1 0 0 0 .893-1.448m-7-7.15L7 6.193V2.036l2-.024v4.237l.11.215l1.827 3.542H5.049zM3 14.017l1.54-3.011h6.916l1.547 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeakerStop(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 15.006h5.007a5.03 5.03 0 0 1-.998-.993L3 14.017l1.542-3.011H6V11c0-.34.034-.673.099-.994h-1.05l1.844-3.598L7 6.193V2.036l2-.024v4.237l.07.137c.31-.13.636-.23.974-.295L10 6.006v-4h1v-1H9.994V1l-.456.005H5V2h1v3.952l-3.894 7.609A1 1 0 0 0 3 15.006m5.778-7.332a4 4 0 1 1 4.444 6.652a4 4 0 0 1-4.444-6.652m.1 5.447A3 3 0 0 0 11 14a3 3 0 0 0 1.74-.55L8.55 9.26A3 3 0 0 0 8 11a3 3 0 0 0 .879 2.121m.382-4.57l4.19 4.189A3 3 0 0 0 14 11a3 3 0 0 0-3-3a3 3 0 0 0-1.74.55" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.377 10.573a7.63 7.63 0 0 1-.383-2.38V6.195a5.115 5.115 0 0 0-1.268-3.446a5.138 5.138 0 0 0-3.242-1.722c-.694-.072-1.4 0-2.07.227c-.67.215-1.28.574-1.794 1.053a4.923 4.923 0 0 0-1.208 1.675a5.067 5.067 0 0 0-.431 2.022v2.2a7.61 7.61 0 0 1-.383 2.37L2 12.343l.479.658h3.505c0 .526.215 1.04.586 1.412c.37.37.885.586 1.412.586c.526 0 1.04-.215 1.411-.586s.587-.886.587-1.412h3.505l.478-.658zm-4.69 3.147a.997.997 0 0 1-.705.299a.997.997 0 0 1-.706-.3a.997.997 0 0 1-.3-.705h1.999a.939.939 0 0 1-.287.706zm-5.515-1.71l.371-1.114a8.633 8.633 0 0 0 .443-2.691V6.004c0-.563.12-1.113.347-1.616c.227-.514.55-.969.969-1.34c.419-.382.91-.67 1.436-.837c.538-.18 1.1-.24 1.65-.18a4.147 4.147 0 0 1 2.597 1.4a4.133 4.133 0 0 1 1.004 2.776v2.01c0 .909.144 1.818.443 2.691l.371 1.113h-9.63v-.012z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellDot(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12.994 7.875A4.008 4.008 0 0 1 12 8h-.01v.217c0 .909.143 1.818.442 2.691l.371 1.113h-9.63v-.012l.37-1.113a8.633 8.633 0 0 0 .443-2.691V6.004c0-.563.12-1.113.347-1.616c.227-.514.55-.969.969-1.34c.419-.382.91-.67 1.436-.837c.538-.18 1.1-.24 1.65-.18l.12.018a4 4 0 0 1 .673-.887a5.15 5.15 0 0 0-.697-.135c-.694-.072-1.4 0-2.07.227c-.67.215-1.28.574-1.794 1.053a4.923 4.923 0 0 0-1.208 1.675a5.067 5.067 0 0 0-.431 2.022v2.2a7.61 7.61 0 0 1-.383 2.37L2 12.343l.479.658h3.505c0 .526.215 1.04.586 1.412c.37.37.885.586 1.412.586c.526 0 1.04-.215 1.411-.586s.587-.886.587-1.412h3.505l.478-.658l-.586-1.77a7.63 7.63 0 0 1-.383-2.381zM7.982 14.02a.997.997 0 0 0 .706-.3a.939.939 0 0 0 .287-.705H6.977c0 .263.107.515.299.706a.999.999 0 0 0 .706.299" clip-rule="evenodd"/><path d="M12 7a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlash(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M11.027 2.086a5.128 5.128 0 0 0-2.543-1.06c-.694-.071-1.4 0-2.07.228c-.67.215-1.28.574-1.794 1.053a4.923 4.923 0 0 0-1.208 1.675a5.067 5.067 0 0 0-.431 2.022v2.2a7.6 7.6 0 0 1-.355 2.282l1.3-1.3c.04-.326.06-.654.06-.981V6.004c0-.563.12-1.113.347-1.616c.227-.514.55-.969.969-1.34c.419-.382.91-.67 1.436-.837c.538-.18 1.1-.24 1.65-.18a4.17 4.17 0 0 1 1.92.774zM7.024 12.02h5.779l-.37-1.113a8.298 8.298 0 0 1-.444-2.691V7.055l1.005-1.004v2.142c0 .813.132 1.615.383 2.38l.586 1.771l-.478.658H9.98c0 .526-.216 1.04-.587 1.412c-.37.37-.885.586-1.411.586a2.011 2.011 0 0 1-1.412-.586a2.014 2.014 0 0 1-.585-1.354zm.958 1.998a.997.997 0 0 0 .706-.3a.939.939 0 0 0 .287-.705H6.977c0 .263.107.515.299.706a.997.997 0 0 0 .706.299" clip-rule="evenodd"/><path d="M1 14.5L15.142.358l.707.707L1.707 15.207z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlashDot(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M8.617 1.045c-.335.28-.628.607-.87.97c-.34.02-.68.086-1.01.196a3.877 3.877 0 0 0-1.435.838a3.85 3.85 0 0 0-.97 1.34a3.918 3.918 0 0 0-.346 1.615v2.2c0 .616-.071 1.23-.206 1.834L2.268 11.55l.33-.977a7.61 7.61 0 0 0 .383-2.368V6.004c0-.706.156-1.388.43-2.022A4.923 4.923 0 0 1 4.62 2.307a4.777 4.777 0 0 1 1.795-1.053a4.874 4.874 0 0 1 2.202-.21m4.397 7.694a4.475 4.475 0 0 1-.991.231c.058.656.193 1.307.41 1.938l.37 1.113H6.318l-.98.981h.646c0 .526.215 1.04.586 1.412c.37.37.885.586 1.412.586c.526 0 1.04-.215 1.411-.586s.587-.886.587-1.412h3.505l.478-.658l-.586-1.77a7.655 7.655 0 0 1-.363-1.835m-4.326 4.98a.997.997 0 0 1-.706.3a.997.997 0 0 1-.706-.3a.997.997 0 0 1-.3-.705h1.999a.939.939 0 0 1-.287.706m6.561-12.053a4.024 4.024 0 0 0-.682-.733l.575-.575l.707.707zM8.933 6.567L1 14.5l.707.707L9.666 7.25a4.023 4.023 0 0 1-.733-.682" clip-rule="evenodd"/><path d="M12 7a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blank(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(``),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 13V3h3.362c1.116 0 1.954.224 2.515.673c.565.449.848 1.113.848 1.992c0 .467-.137.881-.41 1.243c-.273.357-.645.634-1.116.831c.556.151.993.44 1.314.865c.325.422.487.925.487 1.511c0 .898-.299 1.603-.897 2.116c-.598.513-1.443.769-2.536.769zm1.356-4.677v3.599h2.24c.63 0 1.127-.158 1.49-.474c.367-.32.55-.76.55-1.319c0-1.204-.673-1.806-2.02-1.806zm0-1.058h2.049c.593 0 1.066-.144 1.42-.433c.357-.288.536-.68.536-1.174c0-.55-.165-.948-.494-1.195c-.33-.252-.831-.378-1.505-.378H6.356z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 2H9l-.35.15l-.65.64l-.65-.64L7 2H1.5l-.5.5v10l.5.5h5.29l.86.85h.7l.86-.85h5.29l.5-.5v-10zm-7 10.32l-.18-.17L7 12H2V3h4.79l.74.74zM14 12H9l-.35.15l-.14.13V3.7l.7-.7H14zM6 5H3v1h3zm0 4H3v1h3zM3 7h3v1H3zm10-2h-3v1h3zm-3 2h3v1h-3zm0 2h3v1h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 1h-9l-.5.5v13l.872.335L8 10.247l4.128 4.588L13 14.5v-13zM12 13.2L8.372 9.165h-.744L4 13.2V2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BracketDot(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M6 2.972v.012h-.09c-.199 0-.386.04-.562.121a1.468 1.468 0 0 0-.457.32a1.47 1.47 0 0 0-.42 1.035c0 .227.004.453.012.677c.008.224.008.446 0 .665a4.751 4.751 0 0 1-.055.64c-.033.21-.089.41-.168.603A2.284 2.284 0 0 1 3.522 8a2.284 2.284 0 0 1 .738.955c.08.193.135.395.168.608c.033.206.051.42.055.64c.008.216.008.438 0 .666c-.008.22-.012.444-.012.672c0 .203.037.394.111.573c.079.177.182.333.309.468c.13.13.283.237.457.319c.175.077.363.115.563.115H6V14h-.09c-.313 0-.616-.062-.909-.185a2.33 2.33 0 0 1-.775-.53a2.23 2.23 0 0 1-.493-.753v-.001a3.54 3.54 0 0 1-.198-.824v-.002a6.186 6.186 0 0 1-.024-.87c.012-.289.018-.578.018-.868c0-.198-.04-.387-.117-.566V9.4a1.414 1.414 0 0 0-.308-.465l-.002-.002a1.377 1.377 0 0 0-.455-.318h-.001a1.316 1.316 0 0 0-.557-.122H2v-.984h.09c.195 0 .38-.038.557-.115a1.504 1.504 0 0 0 .765-.786v-.002c.078-.18.117-.37.117-.572c0-.29-.006-.58-.018-.869a6.08 6.08 0 0 1 .024-.863V4.3c.033-.287.099-.564.197-.83v-.001a2.23 2.23 0 0 1 .494-.753A2.33 2.33 0 0 1 5 2.185c.293-.123.596-.185.91-.185H6zm7.923 5.52H14v-.984h-.09c-.195 0-.38-.04-.556-.121l-.001-.001a1.376 1.376 0 0 1-.455-.318l-.002-.002a1.414 1.414 0 0 1-.307-.465l-.001-.002a1.405 1.405 0 0 1-.117-.566c0-.29.006-.58.018-.869a6.19 6.19 0 0 0-.024-.87v-.001a3.542 3.542 0 0 0-.197-.824v-.001a2.23 2.23 0 0 0-.494-.753a2.33 2.33 0 0 0-.775-.53a2.325 2.325 0 0 0-.91-.185H10v.984h.09c.2 0 .386.038.562.115c.174.082.326.188.457.32c.127.134.23.29.309.467c.074.18.11.37.11.573c0 .228-.003.452-.011.672c-.008.228-.008.45 0 .665c.004.222.022.435.055.64c.033.214.089.416.168.609a2.282 2.282 0 0 0 .738.955l-.032.025c.53.058 1.03.221 1.477.467" clip-rule="evenodd"/><path d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BracketError(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2.972v.012h-.09c-.199 0-.386.04-.562.121a1.468 1.468 0 0 0-.457.32a1.47 1.47 0 0 0-.42 1.035c0 .227.004.453.012.677c.008.224.008.446 0 .665a4.751 4.751 0 0 1-.055.64c-.033.21-.089.41-.168.603A2.284 2.284 0 0 1 3.522 8a2.284 2.284 0 0 1 .738.955c.08.193.135.395.168.608c.033.206.051.42.055.64c.008.216.008.438 0 .666c-.008.22-.012.444-.012.672c0 .203.037.394.111.573c.079.177.182.333.309.468c.13.13.283.237.457.319c.175.077.363.115.563.115H6V14h-.09c-.313 0-.616-.062-.909-.185a2.33 2.33 0 0 1-.775-.53a2.23 2.23 0 0 1-.493-.753v-.001a3.54 3.54 0 0 1-.198-.824v-.002a6.186 6.186 0 0 1-.024-.87c.012-.289.018-.578.018-.868c0-.198-.04-.387-.117-.566V9.4a1.414 1.414 0 0 0-.308-.465l-.002-.002a1.377 1.377 0 0 0-.455-.318h-.001a1.316 1.316 0 0 0-.557-.122H2v-.984h.09c.195 0 .38-.038.557-.115a1.504 1.504 0 0 0 .765-.786v-.002c.078-.18.117-.37.117-.572c0-.29-.006-.58-.018-.869a6.08 6.08 0 0 1 .024-.863V4.3c.033-.287.099-.564.197-.83v-.001a2.23 2.23 0 0 1 .494-.753A2.33 2.33 0 0 1 5 2.185c.293-.123.596-.185.91-.185H6zm7.923 5.52H14v-.984h-.09c-.195 0-.38-.04-.556-.121l-.001-.001a1.376 1.376 0 0 1-.455-.318l-.002-.002a1.414 1.414 0 0 1-.307-.465l-.001-.002a1.405 1.405 0 0 1-.117-.566c0-.29.006-.58.018-.869a6.19 6.19 0 0 0-.024-.87v-.001a3.542 3.542 0 0 0-.197-.824v-.001a2.23 2.23 0 0 0-.494-.753a2.33 2.33 0 0 0-.775-.53a2.325 2.325 0 0 0-.91-.185H10v.984h.09c.2 0 .386.038.562.115c.174.082.326.188.457.32c.127.134.23.29.309.467c.074.18.11.37.11.573c0 .228-.003.452-.011.672c-.008.228-.008.45 0 .665c.004.222.022.435.055.64c.033.214.089.416.168.609a2.282 2.282 0 0 0 .738.955l-.032.025c.53.058 1.03.221 1.477.467m-3.59 1.014a3 3 0 1 1 3.334 4.988a3 3 0 0 1-3.334-4.988m2.813.64L12 11.293l-1.146-1.147l-.707.707L11.293 12l-1.147 1.146l.708.708L12 12.707l1.146 1.146l.708-.707L12.707 12l1.147-1.146z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 4H11V2.5l-.5-.5h-5l-.5.5V4H1.5l-.5.5v8l.5.5h13l.5-.5v-8zM6 3h4v1H6zm8 2v.76L10 8v-.5L9.51 7h-3L6 7.5V8L2 5.71V5zM9 8v1H7V8zm-7 4V6.86l4 2.29v.35l.5.5h3l.5-.5v-.31l4-2.28V12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Broadcast(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.667 2.011A6 6 0 0 1 8 1a6.007 6.007 0 0 1 6 6a6 6 0 0 1-3.996 5.655v-.044c.016-.014.031-.03.046-.045a1.48 1.48 0 0 0 .434-1.046v-.137A5.042 5.042 0 0 0 12.19 4.2a5.04 5.04 0 1 0-6.69 7.176v.144a1.48 1.48 0 0 0 .48 1.09v.04A5.999 5.999 0 0 1 4.667 2.01z"/><path d="M9.343 11.86a.48.48 0 0 1-.34.14v2.52a.48.48 0 0 1-.48.48H7.46c.011 0-.004-.004-.034-.012c-.075-.02-.241-.064-.305-.129a.48.48 0 0 1-.141-.34V12a.48.48 0 0 1-.48-.48V9.5a1 1 0 0 1 1-1h.984a1 1 0 0 1 1 1v2.02a.48.48 0 0 1-.137.335z"/><path d="M10.64 7c0 .525-.157 1.034-.445 1.465c.183.302.289.656.289 1.035v.106a3.596 3.596 0 0 0 .06-5.15A3.6 3.6 0 1 0 5.5 9.59V9.5c0-.384.108-.743.296-1.047A2.64 2.64 0 1 1 10.64 7"/><path d="M9 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Browser(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 1h13l.5.5v12l-.5.5h-13l-.5-.5v-12zM2 5v8h12V5zm0-1h12V2H2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.877 4.5v-.582a2.918 2.918 0 1 0-5.836 0V4.5h-.833L2.545 2.829l-.593.59l1.611 1.619l-.019.049a8.03 8.03 0 0 0-.503 2.831c0 .196.007.39.02.58l.003.045H1v.836h2.169l.006.034c.172.941.504 1.802.954 2.531l.034.055L2.2 13.962l.592.592l1.871-1.872l.058.066c.868.992 2.002 1.589 3.238 1.589c1.218 0 2.336-.579 3.199-1.544l.057-.064l1.91 1.92l.593-.591l-1.996-2.006l.035-.056c.467-.74.81-1.619.986-2.583l.006-.034h2.171v-.836h-2.065l.003-.044a8.43 8.43 0 0 0 .02-.58a8.02 8.02 0 0 0-.517-2.866l-.019-.05l1.57-1.57l-.592-.59L11.662 4.5zm-5 0v-.582a2.082 2.082 0 1 1 4.164 0V4.5H5.878zm5.697.837l.02.053c.283.753.447 1.61.447 2.528c0 1.61-.503 3.034-1.274 4.037c-.77 1.001-1.771 1.545-2.808 1.545c-1.036 0-2.037-.544-2.807-1.545c-.772-1.003-1.275-2.427-1.275-4.037c0-.918.164-1.775.448-2.528l.02-.053z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 2H13V1h-1v1H4V1H3v1H1.5l-.5.5v12l.5.5h13l.5-.5v-12zM14 14H2V5h12zm0-10H2V3h12zM4 8H3v1h1zm-1 2h1v1H3zm1 2H3v1h1zm2-4h1v1H6zm1 2H6v1h1zm-1 2h1v1H6zm1-6H6v1h1zm2 2h1v1H9zm1 2H9v1h1zm-1 2h1v1H9zm1-6H9v1h1zm2 2h1v1h-1zm1 2h-1v1h1zm-1-4h1v1h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallIncoming(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.547 9.328a1.567 1.567 0 0 0-.594-.117a1.202 1.202 0 0 0-.555.101a2.762 2.762 0 0 0-.43.258a2.166 2.166 0 0 0-.359.328c-.104.12-.205.23-.304.329a2.409 2.409 0 0 1-.29.25a.534.534 0 0 1-.695-.063a32.17 32.17 0 0 1-.328-.312c-.14-.136-.312-.3-.515-.493A61.776 61.776 0 0 1 7.844 9l-.68-.664a25.847 25.847 0 0 1-1.21-1.266a5.312 5.312 0 0 1-.391-.484c-.094-.135-.141-.234-.141-.297a.46.46 0 0 1 .101-.312c.073-.094.16-.19.258-.29c.1-.098.209-.203.328-.312c.12-.11.23-.227.329-.352c.098-.125.182-.268.25-.43c.067-.16.101-.343.101-.546a1.567 1.567 0 0 0-.453-1.102a7.604 7.604 0 0 1-.531-.578a6.487 6.487 0 0 0-.617-.64a4.207 4.207 0 0 0-.696-.516A1.46 1.46 0 0 0 3.742 1a1.567 1.567 0 0 0-1.101.453c-.271.271-.508.513-.711.727a4.006 4.006 0 0 0-.516.664a2.63 2.63 0 0 0-.312.765A4.39 4.39 0 0 0 1 4.625c0 .552.089 1.125.266 1.719c.177.593.416 1.185.718 1.773c.302.589.67 1.167 1.102 1.735c.432.567.901 1.106 1.406 1.617c.505.51 1.042.982 1.61 1.414c.567.432 1.148.805 1.742 1.117c.593.313 1.19.557 1.789.734a6.157 6.157 0 0 0 1.75.266a4.696 4.696 0 0 0 1.008-.11a2.59 2.59 0 0 0 .773-.312c.23-.14.45-.312.664-.515c.214-.204.453-.438.719-.704A1.568 1.568 0 0 0 15 12.257a2.009 2.009 0 0 0-.102-.515a1.674 1.674 0 0 0-.257-.484a7.24 7.24 0 0 0-.368-.445a5.381 5.381 0 0 0-.421-.422a91.549 91.549 0 0 0-.43-.383a8.277 8.277 0 0 1-.367-.344a1.516 1.516 0 0 0-.508-.336m-.367 4.586a3.13 3.13 0 0 1-.797.086a5.526 5.526 0 0 1-1.516-.242a8.362 8.362 0 0 1-1.586-.664a13.205 13.205 0 0 1-3.047-2.297a17.15 17.15 0 0 1-1.289-1.461a10.502 10.502 0 0 1-1.03-1.578a9.12 9.12 0 0 1-.673-1.61A5.308 5.308 0 0 1 2 4.602a3.34 3.34 0 0 1 .094-.79c.057-.218.143-.414.258-.585c.114-.172.255-.339.421-.5c.167-.162.357-.35.57-.563a.542.542 0 0 1 .4-.164c.062-.005.158.036.288.125c.13.089.271.195.422.32a7.058 7.058 0 0 1 .899.899c.125.15.229.289.312.414c.083.125.125.221.125.289a.429.429 0 0 1-.101.312c-.073.084-.16.18-.258.29c-.1.109-.209.213-.328.312c-.12.099-.23.216-.329.351a2.266 2.266 0 0 0-.25.438a1.345 1.345 0 0 0-.101.54c.005.213.047.413.125.6c.078.188.19.355.336.5l3.726 3.727a1.527 1.527 0 0 0 1.102.46a1.2 1.2 0 0 0 .547-.1a2.414 2.414 0 0 0 .789-.586c.11-.12.21-.23.305-.329c.093-.098.19-.182.289-.25a.545.545 0 0 1 .312-.101c.073 0 .172.042.297.125c.125.083.263.19.414.32c.151.13.307.274.469.43c.161.156.305.312.43.469c.124.156.229.297.312.422c.083.125.125.22.125.289a.533.533 0 0 1-.164.39c-.224.219-.414.41-.57.57a3.159 3.159 0 0 1-.5.422a1.93 1.93 0 0 1-.586.266M15 1.704l-4.64 4.648h3.288v1h-5v-5h1V5.64L14.297 1l.703.703z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallOutgoing(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.648 6.648L13.29 2H10V1h5v5h-1V2.71L9.352 7.353l-.704-.704zm3.305 2.563a1.567 1.567 0 0 1 1.102.453c.11.11.232.224.367.344l.43.383c.15.135.291.276.421.421c.13.146.253.295.368.446c.114.15.2.312.257.484c.058.172.092.344.102.516a1.568 1.568 0 0 1-.453 1.101c-.266.266-.505.5-.719.704a4.006 4.006 0 0 1-.664.515c-.23.14-.487.245-.773.313a4.696 4.696 0 0 1-1.008.109a6.157 6.157 0 0 1-1.75-.266A9.819 9.819 0 0 1 7.843 14a12.445 12.445 0 0 1-1.741-1.117a15.329 15.329 0 0 1-1.61-1.414c-.505-.51-.974-1.05-1.406-1.617a11.64 11.64 0 0 1-1.102-1.735a10.38 10.38 0 0 1-.718-1.773A6.005 6.005 0 0 1 1 4.625c0-.396.034-.734.102-1.016a2.63 2.63 0 0 1 .312-.765c.14-.23.313-.45.516-.664c.203-.214.44-.456.71-.727A1.567 1.567 0 0 1 3.743 1c.26 0 .51.07.75.21c.24.142.472.313.696.517c.223.203.43.416.617.64c.187.224.364.417.53.578a1.567 1.567 0 0 1 .453 1.102a1.4 1.4 0 0 1-.1.547a1.824 1.824 0 0 1-.25.43a2.983 2.983 0 0 1-.329.351c-.12.11-.229.214-.328.313a3.128 3.128 0 0 0-.258.289a.46.46 0 0 0-.101.312c0 .063.047.162.14.297a5.3 5.3 0 0 0 .391.484a24.386 24.386 0 0 0 1.211 1.266c.234.23.461.45.68.664c.218.214.43.417.633.61c.203.192.375.356.515.492c.14.135.25.24.328.312a.534.534 0 0 0 .696.063c.093-.068.19-.152.289-.25c.099-.1.2-.209.304-.329c.104-.12.224-.229.36-.328c.135-.099.278-.185.43-.258a1.21 1.21 0 0 1 .554-.101zM11.383 14c.318 0 .583-.029.797-.086a1.93 1.93 0 0 0 .586-.266c.177-.12.343-.26.5-.421c.156-.162.346-.352.57-.57c.11-.11.164-.24.164-.391c0-.068-.042-.164-.125-.29a6.122 6.122 0 0 0-.313-.421a5.01 5.01 0 0 0-.43-.47c-.16-.155-.317-.299-.468-.429a4.322 4.322 0 0 0-.414-.32c-.125-.083-.224-.125-.297-.125a.545.545 0 0 0-.312.101a1.801 1.801 0 0 0-.29.25c-.093.1-.195.209-.304.329c-.11.12-.23.229-.36.328c-.13.099-.273.185-.43.258a1.208 1.208 0 0 1-.546.101a1.527 1.527 0 0 1-1.102-.46L4.883 7.39a1.537 1.537 0 0 1-.336-.5a1.655 1.655 0 0 1-.125-.602c0-.203.034-.383.101-.539c.068-.156.151-.302.25-.438c.1-.135.209-.252.329-.351c.12-.099.229-.203.328-.313c.099-.109.185-.205.258-.289a.429.429 0 0 0 .101-.312c0-.068-.042-.164-.125-.29a5.085 5.085 0 0 0-.312-.413a6.791 6.791 0 0 0-.43-.469a6.787 6.787 0 0 0-.469-.43a5.674 5.674 0 0 0-.422-.32c-.13-.089-.226-.13-.289-.125a.542.542 0 0 0-.398.164a65.24 65.24 0 0 1-.57.563a3.073 3.073 0 0 0-.422.5a1.9 1.9 0 0 0-.258.586A3.377 3.377 0 0 0 2 4.601c0 .5.08 1.015.242 1.546a9.12 9.12 0 0 0 .672 1.61c.287.541.63 1.068 1.031 1.578c.401.51.831.997 1.29 1.46a13.205 13.205 0 0 0 3.046 2.298a8.37 8.37 0 0 0 1.586.664a5.526 5.526 0 0 0 1.516.242z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaseSensitive(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.854 11.702h-1l-.816-2.159H3.772l-.768 2.16H2L4.954 4h.935zm-2.111-2.97L5.534 5.45a3.142 3.142 0 0 1-.118-.515h-.021c-.036.218-.077.39-.124.515L4.073 8.732zm7.013 2.97h-.88v-.86h-.022c-.383.66-.947.99-1.692.99c-.548 0-.978-.146-1.29-.436c-.307-.29-.461-.675-.461-1.155c0-1.027.605-1.625 1.815-1.794l1.65-.23c0-.935-.379-1.403-1.134-1.403c-.663 0-1.26.226-1.794.677V6.59c.54-.344 1.164-.516 1.87-.516c1.292 0 1.938.684 1.938 2.052zm-.88-2.782l-1.327.183c-.409.057-.717.159-.924.306c-.208.143-.312.399-.312.768c0 .268.095.489.285.66c.193.169.45.253.768.253a1.41 1.41 0 0 0 1.08-.457c.286-.308.43-.696.43-1.165z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.431 3.323l-8.47 10l-.79-.036l-3.35-4.77l.818-.574l2.978 4.24l8.051-9.506z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckAll(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15.62 3.596L7.815 12.81l-.728-.033L4 8.382l.754-.53l2.744 3.907L14.917 3z"/><path d="m7.234 8.774l4.386-5.178L10.917 3l-4.23 4.994zm-1.55.403l.548.78l-.547-.78zm-1.617 1.91l.547.78l-.799.943l-.728-.033L0 8.382l.754-.53l2.744 3.907l.57-.672z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checklist(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 4.48h-.71L2 3.43l.71-.7l.69.68L4.81 2l.71.71zM6.99 3h8v1h-8zm0 3h8v1h-8zm8 3h-8v1h8zm-8 3h8v1h-8zM3.04 7.48h.71l1.77-1.77l-.71-.7L3.4 6.42l-.69-.69l-.71.71zm.71 3.01h-.71L2 9.45l.71-.71l.69.69l1.41-1.42l.71.71zm-.71 3.01h.71l1.77-1.77l-.71-.71l-1.41 1.42l-.69-.69l-.71.7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.976 10.072l4.357-4.357l.62.618L8.284 11h-.618L3 6.333l.619-.618z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.928 7.976l4.357 4.357l-.618.62L5 8.284v-.618L9.667 3l.618.619z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.072 8.024L5.715 3.667l.618-.62L11 7.716v.618L6.333 13l-.618-.619z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.024 5.928l-4.357 4.357l-.62-.618L7.716 5h.618L13 9.667l-.619.618z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chip(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 1h1v2h1V1h1v2h1V1h1v2h1l1 1v1h2v1h-2v1h2v1h-2v1h2v1h-2v1l-1 1h-1v2H9v-2H8v2H7v-2H6v2H5v-2H4l-1-1v-1H1V9h2V8H1V7h2V6H1V5h2V4l1-1h1zM4 11h7V4H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeClose(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.116 8l-4.558 4.558l.884.884L8 8.884l4.558 4.558l.884-.884L8.884 8l4.558-4.558l-.884-.884L8 7.116L3.442 2.558l-.884.884z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeMaximize(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3v10h10V3zm9 9H4V4h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeMinimize(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 8v1H3V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChromeRestore(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3 5v9h9V5zm8 8H4V6h7z"/><path fill-rule="evenodd" d="M5 5h1V4h7v7h-1v1h2V3H5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8m2.61-4a2.61 2.61 0 1 1-5.22 0a2.61 2.61 0 0 1 5.22 0M8 5.246" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4c.367 0 .721.048 1.063.145a3.943 3.943 0 0 1 1.762 1.031a3.944 3.944 0 0 1 1.03 1.762c.097.34.145.695.145 1.062c0 .367-.048.721-.145 1.063a3.94 3.94 0 0 1-1.03 1.765a4.017 4.017 0 0 1-1.762 1.031C8.72 11.953 8.367 12 8 12s-.721-.047-1.063-.14a4.056 4.056 0 0 1-1.765-1.032A4.055 4.055 0 0 1 4.14 9.062A3.992 3.992 0 0 1 4 8c0-.367.047-.721.14-1.063a4.02 4.02 0 0 1 .407-.953A4.089 4.089 0 0 1 5.98 4.546a3.94 3.94 0 0 1 .957-.401A3.89 3.89 0 0 1 8 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLarge(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.588 2.215A5.808 5.808 0 0 0 8 2c-.554 0-1.082.073-1.588.215l-.006.002c-.514.141-.99.342-1.432.601A6.156 6.156 0 0 0 2.82 4.98l-.002.004A5.967 5.967 0 0 0 2.21 6.41A5.986 5.986 0 0 0 2 8c0 .555.07 1.085.21 1.591a6.05 6.05 0 0 0 1.548 2.651c.37.365.774.677 1.216.94a6.1 6.1 0 0 0 1.435.609A6.02 6.02 0 0 0 8 14c.555 0 1.085-.07 1.591-.21c.515-.145.99-.348 1.426-.607l.004-.002a6.16 6.16 0 0 0 2.161-2.155a5.85 5.85 0 0 0 .6-1.432l.003-.006A5.807 5.807 0 0 0 14 8c0-.554-.072-1.082-.215-1.588l-.002-.006a5.772 5.772 0 0 0-.6-1.423l-.002-.004a5.9 5.9 0 0 0-.942-1.21l-.008-.008a5.902 5.902 0 0 0-1.21-.942l-.004-.002a5.772 5.772 0 0 0-1.423-.6zm4.455 9.32a7.157 7.157 0 0 1-2.516 2.508a6.966 6.966 0 0 1-1.668.71A6.984 6.984 0 0 1 8 15a6.984 6.984 0 0 1-1.86-.246a7.098 7.098 0 0 1-1.674-.711a7.3 7.3 0 0 1-1.415-1.094a7.295 7.295 0 0 1-1.094-1.415a7.098 7.098 0 0 1-.71-1.675A6.985 6.985 0 0 1 1 8c0-.643.082-1.262.246-1.86a6.968 6.968 0 0 1 .711-1.667a7.156 7.156 0 0 1 2.509-2.516a6.895 6.895 0 0 1 1.675-.704A6.808 6.808 0 0 1 8 1a6.8 6.8 0 0 1 1.86.253a6.899 6.899 0 0 1 3.083 1.805a6.903 6.903 0 0 1 1.804 3.083C14.916 6.738 15 7.357 15 8s-.084 1.262-.253 1.86a6.9 6.9 0 0 1-.704 1.674z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLargeFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1a6.8 6.8 0 0 1 1.86.253a6.899 6.899 0 0 1 3.083 1.805a6.903 6.903 0 0 1 1.804 3.083C14.916 6.738 15 7.357 15 8s-.084 1.262-.253 1.86a6.9 6.9 0 0 1-.704 1.674a7.157 7.157 0 0 1-2.516 2.509a6.966 6.966 0 0 1-1.668.71A6.984 6.984 0 0 1 8 15a6.984 6.984 0 0 1-1.86-.246a7.098 7.098 0 0 1-1.674-.711a7.3 7.3 0 0 1-1.415-1.094a7.295 7.295 0 0 1-1.094-1.415a7.098 7.098 0 0 1-.71-1.675A6.985 6.985 0 0 1 1 8c0-.643.082-1.262.246-1.86a6.968 6.968 0 0 1 .711-1.667a7.156 7.156 0 0 1 2.509-2.516a6.895 6.895 0 0 1 1.675-.704A6.808 6.808 0 0 1 8 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleSlash(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1a7 7 0 1 1-7 7a7.008 7.008 0 0 1 7-7M2 8c0 1.418.504 2.79 1.423 3.87l8.447-8.447A5.993 5.993 0 0 0 2 8m12 0c0-1.418-.504-2.79-1.423-3.87L4.13 12.577A5.993 5.993 0 0 0 14 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleSmall(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.832 8.556a1 1 0 1 1-1.663-1.112a1 1 0 0 1 1.663 1.112m.831.555A2 2 0 1 0 6.338 6.89a2 2 0 0 0 3.325 2.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleSmallFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircuitBoard(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 1h-13l-.5.5v13l.5.5h13l.5-.5v-13zM14 14H5v-2h2.3c.3.6 1 1 1.7 1c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2H4v3H2V2h2v2.3c-.6.3-1 1-1 1.7c0 1.1.9 2 2 2s2-.9 2-2h2c0 1.1.9 2 2 2s2-.9 2-2s-.9-2-2-2c-.7 0-1.4.4-1.7 1H6.7c-.3-.6-1-1-1.7-1V2h9zm-6-3c0-.6.4-1 1-1s1 .4 1 1s-.4 1-1 1s-1-.4-1-1M5 5c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1m6 0c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClearAll(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10 12.6l.7.7l1.6-1.6l1.6 1.6l.8-.7L13 11l1.7-1.6l-.8-.8l-1.6 1.7l-1.6-1.7l-.7.8l1.6 1.6zM1 4h14V3H1zm0 3h14V6H1zm8 2.5V9H1v1h8zM9 13v-1H1v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clippy(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 13.992H4v-9h8v2h1v-2.5l-.5-.5H11v-1h-1a2 2 0 0 0-4 0H4.94v1H3.5l-.5.5v10l.5.5H7zm0-11.2a1 1 0 0 1 .8-.8a1 1 0 0 1 .58.06a.94.94 0 0 1 .45.36a1 1 0 1 1-1.75.94a1 1 0 0 1-.08-.56m7.08 9.46L13 13.342v-5.35h-1v5.34l-1.08-1.08l-.71.71l1.94 1.93h.71l1.93-1.93zm-5.92-4.16h.71l1.93 1.93l-.71.71l-1.08-1.08v5.34h-1v-5.35l-1.08 1.09l-.71-.71z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8 8.707l3.646 3.647l.708-.707L8.707 8l3.647-3.646l-.707-.708L8 7.293L4.354 3.646l-.707.708L7.293 8l-3.646 3.646l.707.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseAll(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="m8.621 8.086l-.707-.707L6.5 8.793L5.086 7.379l-.707.707L5.793 9.5l-1.414 1.414l.707.707L6.5 10.207l1.414 1.414l.707-.707L7.207 9.5z"/><path d="m5 3l1-1h7l1 1v7l-1 1h-2v2l-1 1H3l-1-1V6l1-1h2zm1 2h4l1 1v4h2V3H6zm4 1H3v7h7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.957 6h.05a2.99 2.99 0 0 1 2.116.879a3.003 3.003 0 0 1 0 4.242a2.99 2.99 0 0 1-2.117.879v-.013L12 12H4.523a3.486 3.486 0 0 1-2.628-1.16a3.502 3.502 0 0 1 1.958-5.78a3.462 3.462 0 0 1 1.468.04a3.486 3.486 0 0 1 3.657-2.06A3.479 3.479 0 0 1 11.957 6M5 11h7.01a1.994 1.994 0 0 0 1.992-2a2.002 2.002 0 0 0-1.996-2h-.914l-.123-.857a2.49 2.49 0 0 0-2.126-2.122A2.478 2.478 0 0 0 6.231 5.5l-.333.762l-.809-.189A2.49 2.49 0 0 0 4.523 6c-.662 0-1.297.263-1.764.732A2.503 2.503 0 0 0 4.523 11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.957 6h.05a2.99 2.99 0 0 1 2.116.879a3.003 3.003 0 0 1 0 4.242a2.99 2.99 0 0 1-2.117.879v-1a2.002 2.002 0 0 0 0-4h-.914l-.123-.857a2.49 2.49 0 0 0-2.126-2.122A2.478 2.478 0 0 0 6.231 5.5l-.333.762l-.809-.189A2.49 2.49 0 0 0 4.523 6c-.662 0-1.297.263-1.764.732A2.503 2.503 0 0 0 4.523 11h.498v1h-.498a3.486 3.486 0 0 1-2.628-1.16a3.502 3.502 0 0 1 1.958-5.78a3.462 3.462 0 0 1 1.468.04a3.486 3.486 0 0 1 3.657-2.06A3.479 3.479 0 0 1 11.957 6m-5.25 5.121l1.314 1.314V7h.994v5.4l1.278-1.279l.707.707l-2.146 2.147h-.708L6 11.829z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.956 6h.05a2.99 2.99 0 0 1 2.117.879a3.003 3.003 0 0 1 0 4.242a2.99 2.99 0 0 1-2.117.879h-1.995v-1h1.995a2.002 2.002 0 0 0 0-4h-.914l-.123-.857a2.49 2.49 0 0 0-2.126-2.122A2.478 2.478 0 0 0 6.23 5.5l-.333.762l-.809-.189A2.49 2.49 0 0 0 4.523 6c-.662 0-1.297.263-1.764.732A2.503 2.503 0 0 0 4.523 11h2.494v1H4.523a3.486 3.486 0 0 1-2.628-1.16a3.502 3.502 0 0 1-.4-4.137A3.497 3.497 0 0 1 3.853 5.06c.486-.09.987-.077 1.468.041a3.486 3.486 0 0 1 3.657-2.06A3.479 3.479 0 0 1 11.956 6m-1.663 3.853L8.979 8.54v5.436h-.994v-5.4L6.707 9.854L6 9.146L8.146 7h.708L11 9.146z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.708 5.578L2.061 8.224l2.647 2.646l-.708.708l-3-3V7.87l3-3zm7-.708L11 5.578l2.647 2.646L11 10.87l.708.708l3-3V7.87zM4.908 13l.894.448l5-10L9.908 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3 1v.5c0 .47.274.706.8 1.1l.04.03C4.314 2.985 5 3.498 5 4.5V5H4v-.5c0-.47-.274-.706-.8-1.1l-.04-.03C2.686 3.015 2 2.502 2 1.5V1zm3 0v.5c0 .47.274.706.8 1.1l.04.03C7.314 2.985 8 3.498 8 4.5V5H7v-.5c0-.47-.274-.706-.8-1.1l-.04-.03C5.686 3.015 5 2.502 5 1.5V1zm3 0v.5c0 .47.274.706.8 1.1l.04.03c.474.355 1.16.868 1.16 1.87V5h-1v-.5c0-.47-.274-.706-.8-1.1l-.04-.03C8.686 3.015 8 2.502 8 1.5V1z"/><path fill-rule="evenodd" d="m2 7l1-1h10.5a2.5 2.5 0 0 1 0 5h-.626A4.002 4.002 0 0 1 9 14H6a4 4 0 0 1-4-4zm10 3V7H3v3a3 3 0 0 0 3 3h3a3 3 0 0 0 3-3m1-3v3h.5a1.5 1.5 0 0 0 0-3z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CollapseAll(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 9H4v1h5z"/><path fill-rule="evenodd" d="m5 3l1-1h7l1 1v7l-1 1h-2v2l-1 1H3l-1-1V6l1-1h2zm1 2h4l1 1v4h2V3H6zm4 1H3v7h7z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorMode(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1m0 13V2a6 6 0 1 1 0 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Combine(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.5 1l-.5.5v3l.5.5h3l.5-.5v-3L4.5 1zM2 4V2h2v2zm-.5 2l-.5.5v3l.5.5h3l.5-.5v-3L4.5 6zM2 9V7h2v2zm-1 2.5l.5-.5h3l.5.5v3l-.5.5h-3l-.5-.5zm1 .5v2h2v-2zm10.5-7l-.5.5v6l.5.5h3l.5-.5v-6l-.5-.5zM15 8h-2V6h2zm0 3h-2V9h2zM9.1 8H6v1h3.1l-1 1l.7.6l1.8-1.8v-.7L8.8 6.3l-.7.7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 2h-13l-.5.5v9l.5.5H4v2.5l.854.354L7.707 12H14.5l.5-.5v-9zm-.5 9H7.5l-.354.146L5 13.293V11.5l-.5-.5H2V3h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDiscussion(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4 11.29l1-1v1.42l-1.15 1.14L3 12.5V10H1.5L1 9.5v-8l.5-.5h12l.5.5V6h-1V2H2v7h1.5l.5.5zM10.29 13l1.86 1.85l.85-.35V13h1.5l.5-.5v-5l-.5-.5h-8l-.5.5v5l.5.5zm.21-1H7V8h7v4h-1.5l-.5.5v.79l-1.15-1.14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDraft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 2.001H13v1h1v1h1v-1.5zm-5.5 0h2v1H9zm-4 0h2v1H5zm9 8v2h.5l.5-.5v-1.5zm-2 2v-1h-2v1zm-4-1h-.5l-.354.146L5 13.294v-1.793l-.5-.5H4v3.5l.854.354l2.853-2.854H8zm7-3v-2h-1v2zm-13 3v-1H1v1.5l.5.5H2zm0-3v-2H1v2zm0-5v1H1v-1.5l.5-.5H3v1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentUnresolved(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 2h13l.5.5v6.854a4.019 4.019 0 0 0-1-.819V3H2v8h2.5l.5.5v1.793l2.146-2.147L7.5 11h.626c-.082.32-.126.655-.126 1h-.293l-2.853 2.854L4 14.5V12H1.5l-.5-.5v-9z" clip-rule="evenodd"/><path d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m9.009 6.991l2.024 4.042L6.991 9.01L4.967 4.967zm.426 2.444L8.481 7.52l-1.916-.955l.954 1.917z"/><path fill-rule="evenodd" d="M13.98 8.5a6.002 6.002 0 0 1-5.48 5.48V13h-1v.98a6.001 6.001 0 0 1-5.482-5.518H3v-1h-.976A6.001 6.001 0 0 1 7.5 2.02V3h1v-.98a6.001 6.001 0 0 1 5.48 5.48H13v1zM8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassActive(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9.101 13.9a6.004 6.004 0 0 1-.601.08V13h-1v.98a6.001 6.001 0 0 1-5.482-5.518H3v-1h-.976A6.001 6.001 0 0 1 7.5 2.02V3h1v-.98a6.001 6.001 0 0 1 5.48 5.48H13v1h.98a6.004 6.004 0 0 1-.08.601c.334.077.652.196.95.351a7 7 0 1 0-5.397 5.397a3.973 3.973 0 0 1-.352-.95m.803-3.433L6.99 9.01L4.967 4.967L9.009 6.99l1.459 2.913a4.02 4.02 0 0 0-.564.563m-.469-1.032L8.481 7.52l-1.916-.955l.954 1.917z"/><path d="M11.333 10.506a3 3 0 1 1 3.333 4.987a3 3 0 0 1-3.333-4.987m1.698 3.817l1.79-2.387l-.8-.6l-1.48 1.974l-.876-.7l-.624.78l1.278 1.023z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassDot(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M9.101 13.9a6.004 6.004 0 0 1-.601.08V13h-1v.98a6.001 6.001 0 0 1-5.482-5.518H3v-1h-.976A6.001 6.001 0 0 1 7.5 2.02V3h1v-.98a6.001 6.001 0 0 1 5.48 5.48H13v1h.98a6.004 6.004 0 0 1-.08.601c.334.077.652.196.95.351a7 7 0 1 0-5.397 5.397a3.973 3.973 0 0 1-.352-.95m.803-3.433L6.99 9.01L4.967 4.967L9.009 6.99l1.459 2.913a4.02 4.02 0 0 0-.564.563m-.469-1.032L8.481 7.52l-1.916-.955l.954 1.917z" clip-rule="evenodd"/><circle cx="13" cy="13" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copilot(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6.25 9a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-1.5 0v-1.5A.75.75 0 0 1 6.25 9m4.25.75a.75.75 0 0 0-1.5 0v1.5a.75.75 0 0 0 1.5 0z"/><path d="M7.86 1.77c.05.052.097.107.14.164a1.97 1.97 0 0 1 .14-.165c.681-.73 1.737-.899 2.943-.765c1.23.137 2.145.528 2.724 1.261c.566.716.693 1.614.693 2.485c0 .572-.053 1.147-.254 1.655l.168.838l.066.033A2.75 2.75 0 0 1 16 9.736V11c0 .24-.086.438-.156.567a2.177 2.177 0 0 1-.259.366a5.18 5.18 0 0 1-.605.58a9.985 9.985 0 0 1-.556.43c-.24.174-.485.333-.741.48c-.307.177-.749.41-1.296.642C11.296 14.528 9.756 15 8 15c-1.756 0-3.296-.472-4.387-.935a12.079 12.079 0 0 1-1.296-.642a8.503 8.503 0 0 1-.74-.48a9.608 9.608 0 0 1-.557-.43a5.18 5.18 0 0 1-.605-.58a2.17 2.17 0 0 1-.259-.366A1.19 1.19 0 0 1 0 11V9.736a2.75 2.75 0 0 1 1.52-2.46l.067-.033l.167-.838C1.553 5.897 1.5 5.322 1.5 4.75c0-.87.127-1.77.693-2.485c.579-.733 1.494-1.124 2.724-1.26c1.206-.135 2.262.034 2.944.764M3.024 7.708L3 7.824v4.261l.065.038c.264.152.65.356 1.134.561c.972.413 2.307.816 3.801.816c1.494 0 2.83-.403 3.8-.816a10.58 10.58 0 0 0 1.2-.599v-4.26l-.023-.116c-.49.21-1.075.29-1.727.29c-1.146 0-2.06-.327-2.71-.99A3.223 3.223 0 0 1 8 6.266a3.244 3.244 0 0 1-.54.743c-.65.663-1.564.99-2.71.99c-.652 0-1.237-.08-1.727-.29m3.741-4.916c-.193-.207-.637-.414-1.681-.298c-1.02.114-1.48.404-1.713.7c-.247.313-.37.79-.37 1.555c0 .792.129 1.17.308 1.37c.162.181.52.38 1.442.38c.854 0 1.339-.236 1.638-.541c.315-.322.527-.826.618-1.552c.117-.936-.038-1.396-.242-1.614m2.472 0c-.204.218-.359.678-.242 1.614c.09.726.303 1.23.618 1.552c.299.305.784.54 1.638.54c.922 0 1.28-.198 1.442-.379c.179-.2.308-.578.308-1.37c0-.766-.123-1.242-.37-1.555c-.233-.296-.693-.586-1.713-.7c-1.044-.116-1.488.091-1.681.298"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="m4 4l1-1h5.414L14 6.586V14l-1 1H5l-1-1zm9 3l-3-3H5v10h8z"/><path d="M3 1L2 2v10l1 1V2h6.414l-1-1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 5v1H2V5zM2 7h12v5H2zm12-3H2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1m-3 6h2v1h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dash(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8h6v1H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.889 2.095a6.5 6.5 0 1 1 7.222 10.81A6.5 6.5 0 0 1 3.89 2.094zm.555 9.978A5.5 5.5 0 0 0 7.5 13A5.506 5.506 0 0 0 13 7.5a5.5 5.5 0 1 0-8.556 4.573M10.294 4l.706.707l-2.15 2.15a1.514 1.514 0 1 1-.707-.707L10.293 4zM7.221 7.916a.5.5 0 1 0 .556-.832a.5.5 0 0 0-.556.832m4.286-2.449l-.763.763c.166.403.253.834.255 1.27a3.463 3.463 0 0 1-.5 1.777l.735.735a4.477 4.477 0 0 0 .274-4.545zM8.733 4.242A3.373 3.373 0 0 0 7.5 4A3.5 3.5 0 0 0 4 7.5a3.46 3.46 0 0 0 .5 1.777l-.734.735A4.5 4.5 0 0 1 9.5 3.473z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3.5C13 2.119 10.761 1 8 1S3 2.119 3 3.5c0 .04.02.077.024.117H3v8.872l.056.357C3.336 14.056 5.429 15 8 15c2.571 0 4.664-.944 4.944-2.154l.056-.357V3.617h-.024c.004-.04.024-.077.024-.117M8 2.032c2.442 0 4 .964 4 1.468s-1.558 1.468-4 1.468S4 4 4 3.5s1.558-1.468 4-1.468m4 10.458l-.03.131C11.855 13.116 10.431 14 8 14s-3.855-.884-3.97-1.379L4 12.49v-7.5A7.414 7.414 0 0 0 8 6a7.414 7.414 0 0 0 4-1.014z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Debug(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m3.463 12.86l-.005-.07zm7.264.69l-3.034-3.049l1.014-1.014l3.209 3.225l3.163-3.163l1.014 1.014l-3.034 3.034l3.034 3.05l-1.014 1.014l-3.209-3.225L8.707 17.6l-1.014-1.014l3.034-3.034z"/><path fill-rule="evenodd" d="M16.933 5.003V6h1.345l2.843-2.842l1.014 1.014l-2.692 2.691l.033.085a13.75 13.75 0 0 1 .885 4.912c0 .335-.011.667-.034.995l-.005.075h3.54v1.434h-3.72l-.01.058c-.303 1.653-.891 3.16-1.692 4.429l-.06.094l3.423 3.44l-1.017 1.012l-3.274-3.29l-.099.11c-1.479 1.654-3.395 2.646-5.483 2.646c-2.12 0-4.063-1.023-5.552-2.723l-.098-.113l-3.209 3.208l-1.014-1.014l3.366-3.365l-.059-.095c-.772-1.25-1.34-2.725-1.636-4.34l-.01-.057H0V12.93h3.538l-.005-.075a14.23 14.23 0 0 1-.034-.995c0-1.743.31-3.39.863-4.854l.032-.084l-2.762-2.776L2.65 3.135L5.5 6h1.427v-.997a5.003 5.003 0 0 1 10.006 0m-8.572 0V6H15.5v-.997a3.569 3.569 0 0 0-7.138 0zm9.8 2.522l-.034-.09H5.733l-.034.09a12.328 12.328 0 0 0-.766 4.335c0 2.76.862 5.201 2.184 6.92c1.32 1.716 3.036 2.649 4.813 2.649c1.777 0 3.492-.933 4.813-2.65c1.322-1.718 2.184-4.16 2.184-6.919c0-1.574-.28-3.044-.766-4.335" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugAll(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m7.293 9.006l-.88.88A2.484 2.484 0 0 0 4 8a2.488 2.488 0 0 0-2.413 1.886l-.88-.88L0 9.712l1.147 1.146l-.147.146v1H0v.999h1v.053c.051.326.143.643.273.946L0 15.294L.707 16l1.1-1.099A2.873 2.873 0 0 0 4 16a2.875 2.875 0 0 0 2.193-1.099L7.293 16L8 15.294l-1.273-1.292A3.92 3.92 0 0 0 7 13.036v-.067h1v-.965H7v-1l-.147-.146L8 9.712zM4 9.006a1.5 1.5 0 0 1 1.5 1.499h-3A1.498 1.498 0 0 1 4 9.006m2 3.997A2.217 2.217 0 0 1 4 15a2.22 2.22 0 0 1-2-1.998v-1.499h4z"/><path fill-rule="evenodd" d="M3.78 2L3 2.41V7h1V3.35l7.6 5.07L9 10.15v1.2l3.78-2.52V8zM9 13.35v-1.202l5.6-3.728L7 3.35V2.147L15.78 8v.83z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugAlt(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.94 13.5l-1.32 1.32a3.73 3.73 0 0 0-7.24 0L1.06 13.5L0 14.56l1.72 1.72l-.22.22V18H0v1.5h1.5v.08c.077.489.214.966.41 1.42L0 22.94L1.06 24l1.65-1.65A4.308 4.308 0 0 0 6 24a4.31 4.31 0 0 0 3.29-1.65L10.94 24L12 22.94L10.09 21c.198-.464.336-.951.41-1.45v-.1H12V18h-1.5v-1.5l-.22-.22L12 14.56zM6 13.5a2.25 2.25 0 0 1 2.25 2.25h-4.5A2.25 2.25 0 0 1 6 13.5m3 6a3.33 3.33 0 0 1-3 3a3.33 3.33 0 0 1-3-3v-2.25h6zm14.76-9.9v1.26L13.5 17.37V15.6l8.5-5.37L9 2v9.46a5.07 5.07 0 0 0-1.5-.72V.63L8.64 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugAltSmall(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m7.293 9.006l-.88.88A2.484 2.484 0 0 0 4 8a2.488 2.488 0 0 0-2.413 1.886l-.88-.88L0 9.712l1.147 1.146l-.147.146v1H0v.999h1v.053c.051.326.143.643.273.946L0 15.294L.707 16l1.1-1.099A2.873 2.873 0 0 0 4 16a2.875 2.875 0 0 0 2.193-1.099L7.293 16L8 15.294l-1.273-1.292A3.92 3.92 0 0 0 7 13.036v-.067h1v-.965H7v-1l-.147-.146L8 9.712zM4 9.006a1.5 1.5 0 0 1 1.5 1.499h-3A1.498 1.498 0 0 1 4 9.006m2 3.997A2.217 2.217 0 0 1 4 15a2.22 2.22 0 0 1-2-1.998v-1.499h4z"/><path fill-rule="evenodd" d="M5 2.41L5.78 2l9 6v.83L9 12.683v-1.2l4.6-3.063L6 3.35V7H5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugBreakpointConditional(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 4a4 4 0 1 0 0 8a4 4 0 0 0 0-8m2 5v1H6V9zm0-3v1H6V6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugBreakpointConditionalUnverified(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.778 4.674a4 4 0 1 1 4.444 6.652a4 4 0 0 1-4.444-6.652m.694 5.612a2.75 2.75 0 1 0 3.056-4.572a2.75 2.75 0 0 0-3.056 4.572M9.5 6.5h-3v1h3zm0 2h-3v1h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugBreakpointData(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.238 8l-2.31 4H5.31L3 8l2.31-4h4.618z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugBreakpointDataUnverified(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.931 4h-4.62l-2.31 4l2.31 4h4.62l2.31-4zm-.75 6.7h-3.12L4.501 8l1.56-2.7h3.12l1.56 2.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugBreakpointFunction(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 4l4 6.905H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugBreakpointFunctionUnverified(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 11h8L8 4zm2.154-1.25h3.692L8 6.52z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugBreakpointLog(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 3l5 5l-5 5l-5-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugBreakpointLogUnverified(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.02 7.98L8 3l4.98 4.98L8 12.96zM8 10.77l2.79-2.79L8 5.19L5.21 7.98z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugBreakpointUnsupported(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.326 10.222a4 4 0 1 0-6.653-4.444a4 4 0 0 0 6.653 4.444M8.65 10H7.4v1h1.25zM7.4 9V5h1.25v4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugConsole(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.04 1.361l.139-.057H21.32l.14.057l1.178 1.179l.057.139V16.82l-.057.14l-1.179 1.178l-.139.057H14V18a1.99 1.99 0 0 0-.548-1.375h7.673V2.875H7.375v7.282a5.73 5.73 0 0 0-1.571-.164V2.679l.057-.14L7.04 1.362zm9.531 9.452l-2.809 2.8a2 2 0 0 0-.348-.467l-.419-.42l2.236-2.235l-3.606-3.694l.813-.833l4.133 4.133zM9.62 14.82l1.32-1.32L12 14.56l-1.72 1.72l.22.22V18H12v1.45h-1.5v.1a5.888 5.888 0 0 1-.41 1.45L12 22.94L10.94 24l-1.65-1.65A4.308 4.308 0 0 1 6 24a4.31 4.31 0 0 1-3.29-1.65L1.06 24L0 22.94L1.91 21a5.889 5.889 0 0 1-.41-1.42v-.08H0V18h1.5v-1.5l.22-.22L0 14.56l1.06-1.06l1.32 1.32a3.73 3.73 0 0 1 7.24 0m-2.029-.661A2.25 2.25 0 0 0 3.75 15.75h4.5a2.25 2.25 0 0 0-.659-1.591m.449 7.38A3.33 3.33 0 0 0 9 19.5v-2.25H3v2.25a3.33 3.33 0 0 0 3 3a3.33 3.33 0 0 0 2.04-.96z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugContinue(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 2H4v12H2.5zm4.936.39L6.25 3v10l1.186.61l7-5V7.39zM12.71 8l-4.96 3.543V4.457z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugContinueSmall(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2H3v12h1zm3.29.593L6.5 3v10l.79.407l7-5v-.814zM13.14 8L7.5 12.028V3.972z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugCoverage(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 2.41L5.78 2l9 6v.83L9 12.683v-1.2l4.6-3.063L6 3.35V7H5z"/><path d="M6.13 12.124c-.182.18-.322.379-.42.6a1.744 1.744 0 0 0-.145.715v.862a.685.685 0 0 1-.205.495a.703.703 0 0 1-.496.204h-.865a.691.691 0 0 1-.497-.204a.701.701 0 0 1-.205-.495v-.862c0-.258-.049-.496-.147-.716a1.913 1.913 0 0 0-.418-.6a2.525 2.525 0 0 1-.542-.773a2.255 2.255 0 0 1-.19-.927A2.386 2.386 0 0 1 2.332 9.2a2.404 2.404 0 0 1 .87-.87A2.473 2.473 0 0 1 4.432 8a2.41 2.41 0 0 1 1.225.33a2.41 2.41 0 0 1 1.205 2.093c0 .332-.063.641-.19.927a2.525 2.525 0 0 1-.542.774m-1.103.991H3.835v1.186c0 .043.016.08.049.114c.033.033.07.048.115.048h.865a.156.156 0 0 0 .114-.048a.154.154 0 0 0 .049-.114z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugDisconnect(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.617 3.844a2.87 2.87 0 0 0-.451-.868l1.354-1.36L13.904 1l-1.36 1.354a2.877 2.877 0 0 0-.868-.452a3.073 3.073 0 0 0-2.14.075a3.03 3.03 0 0 0-.991.664L7 4.192l4.327 4.328l1.552-1.545c.287-.287.508-.618.663-.992a3.074 3.074 0 0 0 .075-2.14zm-.889 1.804a2.15 2.15 0 0 1-.471.705l-.93.93l-3.09-3.09l.93-.93a2.15 2.15 0 0 1 .704-.472a2.134 2.134 0 0 1 1.689.007c.264.114.494.271.69.472c.2.195.358.426.472.69a2.134 2.134 0 0 1 .007 1.688zm-4.824 4.994l1.484-1.545l-.616-.622l-1.49 1.551l-1.86-1.859l1.491-1.552L6.291 6L4.808 7.545l-.616-.615l-1.551 1.545a3 3 0 0 0-.663.998a3.023 3.023 0 0 0-.233 1.169c0 .332.05.656.15.97c.105.31.258.597.459.862L1 13.834l.615.615l1.36-1.353c.265.2.552.353.862.458c.314.1.638.15.97.15c.406 0 .796-.077 1.17-.232c.378-.155.71-.376.998-.663l1.545-1.552zm-2.262 2.023a2.16 2.16 0 0 1-.834.164c-.301 0-.586-.057-.855-.17a2.278 2.278 0 0 1-.697-.466a2.28 2.28 0 0 1-.465-.697a2.167 2.167 0 0 1-.17-.854a2.16 2.16 0 0 1 .642-1.545l.93-.93l3.09 3.09l-.93.93a2.22 2.22 0 0 1-.711.478" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugLineByLine(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 10V9h9v1zm4-4h5v1h-5zm5-3v1H6V3zm-9 9v1h9v-1z"/><path fill-rule="evenodd" d="m1 2.795l.783-.419l5.371 3.581v.838l-5.371 3.581L1 9.957zm1.007.94v5.281l3.96-2.64z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugPause(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 3H6v10H4.5zm7 0v10H10V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugRerun(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.167 12a3 3 0 0 1-5.74 1.223l-.928.376A4.001 4.001 0 1 0 1 9.556V8.333H0V11l.5.5h2.333v-1H1.568A3 3 0 0 1 7.167 12"/><path d="M5 2.41L5.78 2l9 6v.83L10 12.017v-1.2l3.6-2.397L6 3.35V7H5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugRestart(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.75 8a4.5 4.5 0 0 1-8.61 1.834l-1.391.565A6.001 6.001 0 0 0 14.25 8A6 6 0 0 0 3.5 4.334V2.5H2v4l.75.75h3.5v-1.5H4.352A4.5 4.5 0 0 1 12.75 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugRestartFrame(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 10V9h5.207a5.48 5.48 0 0 0-.185 1zm6.257-3a5.54 5.54 0 0 1 1.08-1H1v1zM6.6 13a5.465 5.465 0 0 1-.393-1H1v1zM15 3v1H1V3zm-3.36 10.031a2.531 2.531 0 1 0-2.192-3.797h1.068v.844h-1.97l-.421-.422v-2.25h.844v1.032a3.375 3.375 0 1 1-.423 3.412l.782-.318a2.532 2.532 0 0 0 2.313 1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugReverseContinue(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 2H12v12h1.5zm-4.936.39L9.75 3v10l-1.186.61l-7-5V7.39zM3.29 8l4.96 3.543V4.457z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugStackframe(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.5 7.15l-4.26-4.74L9.31 2H4.25L3 3.25v9.48l1.25 1.25h5.06l.93-.42l4.26-4.74zm-5.19 5.58H4.25V3.25h5.06l4.26 4.73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugStackframeActive(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path d="m14.5 7.15l-4.26-4.74L9.31 2H4.25L3 3.25v9.48l1.25 1.25h5.06l.93-.42l4.26-4.74zm-5.19 5.58H4.25V3.25h5.06l4.26 4.73z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugStart(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.25 3l1.166-.624l8 5.333v1.248l-8 5.334l-1.166-.624zm1.5 1.401v7.864l5.898-3.932z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugStepBack(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 5.75v-4h1.5v2.542c1.145-1.359 2.911-2.209 4.84-2.209c3.177 0 5.92 2.307 6.16 5.398l.02.269h-1.5l-.022-.226c-.212-2.195-2.202-3.94-4.656-3.94c-1.736 0-3.244.875-4.05 2.166h2.83v1.5H2.707l-.961-.975V5.75h.003zM8 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugStepInto(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 9.532h.542l3.905-3.905l-1.061-1.06l-2.637 2.61V1H7.251v6.177l-2.637-2.61l-1.061 1.06l3.905 3.905zm1.956 3.481a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugStepOut(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1h-.542L3.553 4.905l1.061 1.06l2.637-2.61v6.177h1.498V3.355l2.637 2.61l1.061-1.06L8.542 1zm1.956 12.013a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugStepOver(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.25 5.75v-4h-1.5v2.542c-1.145-1.359-2.911-2.209-4.84-2.209c-3.177 0-5.92 2.307-6.16 5.398l-.02.269h1.501l.022-.226c.212-2.195 2.202-3.94 4.656-3.94c1.736 0 3.244.875 4.05 2.166h-2.83v1.5h4.163l.962-.975V5.75zM8 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DebugStop(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13 2l1 1v10l-1 1H3l-1-1V3l1-1zm-.254 1.25H3.255v9.5h9.491z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopDownload(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 15v-1c2 0 2-.6 2-1H1.5l-.5-.5v-10l.5-.5h13l.5.5v9.24l-1-1V3H2v9h5.73l-.5.5l2.5 2.5zm7.86 0l2.5-2.5l-.71-.7L12 13.45V7h-1v6.44l-1.64-1.65l-.71.71l2.5 2.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceCamera(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.707 3H14.5l.5.5v9l-.5.5h-13l-.5-.5v-9l.5-.5h3.793l.853-.854L6.5 2h3l.354.146zM2 12h12V4h-3.5l-.354-.146L9.293 3H6.707l-.853.854L5.5 4H2zm1.5-7a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1M8 6a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0-1a3 3 0 1 0 0 6a3 3 0 0 0 0-6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceCameraVideo(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 4.74L11 6.62V4.5l-.5-.5h-9l-.5.5v7l.5.5h9l.5-.5v-2l3.25 1.87l.75-.47V5.18zM10 11H2V5h8zm4-1l-3-1.7v-.52L14 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceMobile(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 1h8l.5.5v13l-.5.5h-8l-.5-.5v-13zM5 14h7V2H5zm2.5-2h2v1h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diff(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2 3.5l.5-.5h5l.5.5v9l-.5.5h-5l-.5-.5zM3 12h4V6H3zm0-7h4V4H3zm6.5-2h5l.5.5v9l-.5.5h-5l-.5-.5v-9zm.5 9h4v-2h-4zm0-4h4V4h-4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffAdded(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M1.5 1h12l.5.5v12l-.5.5h-12l-.5-.5v-12zM2 13h11V2H2z"/><path d="M8 4H7v3H4v1h3v3h1V8h3V7H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffIgnored(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 1h13l.5.5v13l-.5.5h-13l-.5-.5v-13zM2 14h12V2H2zm8-10h2v2l-6 6H4v-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffModified(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 1h13l.5.5v13l-.5.5h-13l-.5-.5v-13zM2 2v12h12V2zm6 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffMultiple(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="m9.71 1.29l3 3L13 5v9l-1 1H3l-1-1V2l1-1h6zM3 14h9V5L9 2H3zm4-8H5v1h2v2h1V7h2V6H8V4H7zm-2 5h5v1H5z" clip-rule="evenodd"/><path d="m12.42 1l2.29 2.29L15 4v10l-1 1V4l-3-3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffRemoved(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 7v1H5V7z"/><path fill-rule="evenodd" d="M1.5 1h12l.5.5v12l-.5.5h-12l-.5-.5v-12zM2 13h11V2H2z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffRenamed(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 1h13l.5.5v13l-.5.5h-13l-.5-.5v-13zM2 14h12V2H2zm2-5h3v3l5-4l-5-4v3H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffSingle(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.71 1.289l3 3l.29.71v9l-1 1H4l-1-1v-12l1-1h6zM4 13.999h9v-9l-3-3H4zm4-8H6v1h2v2h1v-2h2v-1H9v-2H8zm-2 5h5v1H6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discard(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2v3.5L4 6h3.5V5H4.979l.941-.941a3.552 3.552 0 1 1 5.023 5.023L5.746 14.28l.72.72l5.198-5.198A4.57 4.57 0 0 0 5.2 3.339l-.7.7V2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.23 1h-1.46L3.52 9.25l-.16.22L1 13.59L2.41 15l4.12-2.36l.22-.16L15 4.23V2.77zM2.41 13.59l1.51-3l1.45 1.45zm3.83-2.06L4.47 9.76l8-8l1.77 1.77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditorLayout(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15 6.5l-.47-.5H7V1.47L6.53 1H1.47L1 1.47v8.06l.47.47H4v4.53l.47.47h10.06l.47-.47zM2 9V3h4v6zm12 5H5v-4h1.53L7 9.53V8.013h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ellipsis(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmptyWindow(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 4h3v1H4v3H3V5H0V4h3V1h1zM1 14.5V9h1v5h12V7H8V6h6V4H8V3h6.5l.5.5v11l-.5.5h-13z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorIcon(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.6 1c1.6.1 3.1.9 4.2 2c1.3 1.4 2 3.1 2 5.1c0 1.6-.6 3.1-1.6 4.4c-1 1.2-2.4 2.1-4 2.4c-1.6.3-3.2.1-4.6-.7c-1.4-.8-2.5-2-3.1-3.5C.9 9.2.8 7.5 1.3 6c.5-1.6 1.4-2.9 2.8-3.8C5.4 1.3 7 .9 8.6 1m.5 12.9c1.3-.3 2.5-1 3.4-2.1c.8-1.1 1.3-2.4 1.2-3.8c0-1.6-.6-3.2-1.7-4.3c-1-1-2.2-1.6-3.6-1.7c-1.3-.1-2.7.2-3.8 1c-1.1.8-1.9 1.9-2.3 3.3c-.4 1.3-.4 2.7.2 4c.6 1.3 1.5 2.3 2.7 3c1.2.7 2.6.9 3.9.6M7.9 7.5L10.3 5l.7.7l-2.4 2.5l2.4 2.5l-.7.7l-2.4-2.5l-2.4 2.5l-.7-.7l2.4-2.5l-2.4-2.5l.7-.7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorSmall(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9.177 10.105L8 8.928l-1.177 1.177l-.928-.928L7.072 8L5.895 6.823l.928-.928L8 7.072l1.177-1.177l.928.928L8.928 8l1.177 1.177z"/><path d="M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-1 0a3 3 0 1 0-6 0a3 3 0 0 0 6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exclude(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.412 1H6.588l-.484 2.423l-2.056-1.37l-1.996 1.995l1.37 2.056L1 6.588v2.824l2.423.484l-1.37 2.056l1.995 1.996l2.056-1.37L6.588 15h2.083a4.526 4.526 0 0 1-.917-1.005h-.342l-.288-1.441a4.473 4.473 0 0 1-.067-.334l-.116-.583l-.764-.316l-2 1.334l-.832-.831L4.68 9.823l-.316-.764l-2.358-.471V7.412l2.358-.471l.316-.764l-1.334-2l.831-.832l2 1.335l.764-.316l.471-2.358h1.176l.471 2.358l.764.316l2-1.334l.832.831l-1.334 2.001l.316.764l.582.116c.113.018.225.04.335.067l1.441.288v.342c.38.254.719.563 1.005.917V6.588l-2.422-.484l1.37-2.056l-1.996-1.996l-2.056 1.37zM8 6a2 2 0 0 1 1.875 1.302a4.46 4.46 0 0 0-.9.473a1 1 0 1 0-1.2 1.2a4.46 4.46 0 0 0-.473.9A2 2 0 0 1 8 6m1.28 2.795a3.5 3.5 0 1 1 4.44 5.41a3.5 3.5 0 0 1-4.44-5.41M9 11v1h5v-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandAll(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 9H4v1h5z"/><path d="M7 12V7H6v5z"/><path fill-rule="evenodd" d="m5 3l1-1h7l1 1v7l-1 1h-2v2l-1 1H3l-1-1V6l1-1h2zm1 2h4l1 1v4h2V3H6zm4 1H3v7h7z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.086 7l-2.39-2.398l.702-.704L15 7.5l-3.602 3.602l-.703-.704l2.383-2.382V8H3V7zM1 4h1v7H1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Extensions(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 1.5L15 0h7.5L24 1.5V9l-1.5 1.5H15L13.5 9zm1.5 0V9h7.5V1.5zM0 15V6l1.5-1.5H9L10.5 6v7.5H18l1.5 1.5v7.5L18 24H1.5L0 22.5zm9-1.5V6H1.5v7.5zM9 15H1.5v7.5H9zm1.5 7.5H18V15h-7.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6.003a2.667 2.667 0 1 1 0 5.334a2.667 2.667 0 0 1 0-5.334m0 1a1.667 1.667 0 1 0 0 3.334a1.667 1.667 0 0 0 0-3.334m0-3.336c3.076 0 5.73 2.1 6.467 5.043a.5.5 0 1 1-.97.242a5.67 5.67 0 0 0-10.995.004a.5.5 0 0 1-.97-.243A6.669 6.669 0 0 1 8 3.667"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosed(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.48 1.48a.5.5 0 0 0-.049.65l.049.057l2.69 2.69A6.657 6.657 0 0 0 1.533 8.71a.5.5 0 0 0 .97.242a5.66 5.66 0 0 1 2.386-3.356l1.207 1.207a2.667 2.667 0 0 0 3.771 3.771l3.946 3.946a.5.5 0 0 0 .756-.65l-.049-.057l-4.075-4.076v-.001l-.8-.799l-1.913-1.913h.001l-1.92-1.919v-.001l-.755-.754l-2.871-2.87a.5.5 0 0 0-.707 0m5.323 6.03l2.356 2.357A1.667 1.667 0 0 1 6.802 7.51M8 3.667c-.667 0-1.314.098-1.926.283l.825.824a5.669 5.669 0 0 1 6.6 4.181a.5.5 0 0 0 .97-.242A6.669 6.669 0 0 0 8 3.667m.13 2.34l2.534 2.533A2.668 2.668 0 0 0 8.13 6.006"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feedback(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m4.5 1l-.5.5v1.527a4.551 4.551 0 0 1 1 0V2h9v5h-1.707L11 8.293V7H8.973a4.551 4.551 0 0 1 0 1H10v1.5l.854.354L12.707 8H14.5l.5-.5v-6l-.5-.5z"/><path fill-rule="evenodd" d="M6.417 10.429a3.5 3.5 0 1 0-3.834 0A4.501 4.501 0 0 0 0 14.5v.5h1v-.5a3.502 3.502 0 0 1 7 0v.5h1v-.5a4.501 4.501 0 0 0-2.583-4.071M4.5 10a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.71 4.29l-3-3L10 1H4L3 2v12l1 1h9l1-1V5zM13 14H4V2h5v4h4zm-3-9V2l3 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileBinary(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.57 1.14l3.28 3.3l.15.36v9.7l-.5.5h-11l-.5-.5v-13l.5-.5h7.72zM3 2v12h10V5l-3-3zm1.46 4.052c0 1.287.458 1.93 1.374 1.93c.457 0 .807-.173 1.05-.52c.246-.348.368-.847.368-1.499C7.252 4.654 6.805 4 5.91 4c-.471 0-.831.175-1.08.526c-.247.35-.37.858-.37 1.526m.862-.022c0-.922.183-1.383.55-1.383c.344 0 .516.448.516 1.343s-.176 1.343-.527 1.343c-.36 0-.54-.434-.54-1.303zm3.187 1.886h2.435v-.672h-.792V4l-1.665.336v.687l.82-.177v2.398h-.798zm-1.337 5H4.736v-.672h.798V9.846l-.82.177v-.687L6.38 9v3.244h.792v.671zm1.035-1.931c0 1.287.458 1.93 1.375 1.93c.457 0 .807-.173 1.05-.52c.245-.348.368-.847.368-1.499c0-1.309-.448-1.963-1.343-1.963c-.47 0-.83.175-1.08.526c-.246.35-.37.858-.37 1.526m.862-.022c0-.922.184-1.383.55-1.383c.344 0 .516.448.516 1.343s-.175 1.343-.526 1.343c-.36 0-.54-.434-.54-1.303" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCode(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.57 1.14l3.28 3.3l.15.36v9.7l-.5.5h-11l-.5-.5v-13l.5-.5h7.72zM10 5h3l-3-3zM3 2v12h10V6H9.5L9 5.5V2zm2.062 7.533l1.817-1.828L6.17 7L4 9.179v.707l2.171 2.174l.707-.707zM8.8 7.714l.7-.709l2.189 2.175v.709L9.5 12.062l-.705-.709l1.831-1.82z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMedia(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2h6v3.5l.5.5H12v1h1V4.8l-.15-.36l-3.28-3.3L9.22 1H1.5l-.5.5v13l.5.5H5v-1H2zm7 0l3 3H9zm5.5 6h-8l-.5.5v6l.5.5h8l.5-.5v-6zM14 9v4l-1.63-1.6h-.71l-1.16 1.17l-2.13-2.13h-.71L7 11.1V9zm-2.8 4.27l.81-.81L13.55 14h-1.62zM7 14v-1.49l1-1L10.52 14zm5.5-3.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdf(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.85 4.44l-3.28-3.3l-.35-.14H2.5l-.5.5V7h1V2h6v3.5l.5.5H13v1h1V4.8zM10 5V2l3 3zM2.5 8l-.5.5v6l.5.5h11l.5-.5v-6l-.5-.5zM13 13v1H3V9h10zm-8-1h-.32v1H4v-3h1.06c.75 0 1.13.36 1.13 1a.94.94 0 0 1-.32.72A1.33 1.33 0 0 1 5 12m-.06-1.45h-.26v.93h.26c.36 0 .54-.16.54-.47c0-.31-.18-.46-.54-.46M9 12.58a1.48 1.48 0 0 0 .44-1.12c0-1-.53-1.46-1.6-1.46H6.78v3h1.06A1.6 1.6 0 0 0 9 12.58m-1.55-.13v-1.9h.33a.94.94 0 0 1 .7.25a.91.91 0 0 1 .25.67a1 1 0 0 1-.25.72a.94.94 0 0 1-.69.26zm4.45-.61h-.97V13h-.68v-3h1.74v.55h-1.06v.74h.97z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSubmodule(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 11h1V6.99H2zm1-5.01V5.5l.5-.5h4.43l.43.25l.43.75h5.71l.5.5v8l-.5.5h-11l-.5-.5V12H1.5l-.5-.5v-9l.5-.5h4.42l.44.25l.43.75h5.71l.5.5V6l-1-.03V4H6.5l-.43-.25L5.64 3H2v2.99zm5.07.76L7.64 6H4v3h3.15l.41-.74L8 8h6V7H8.5zM7.45 10H4v4h10V9H8.3l-.41.74z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSymlinkDirectory(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.71 3h6.79l.51.5v10l-.5.5h-13l-.5-.5v-11l.5-.5h5l.35.15zm6.28 10v-1.51l.01-4v-1.5H7.7l-.86.86l-.35.15H2v6zm-6.5-8h6.5l.01-.99H7.5l-.36-.15l-.85-.85H2v3h4.28l.86-.86zm2.29 4.07L8.42 7.7l.74-.69l2.22 2.22v.71l-2.29 2.21l-.7-.72l1.4-1.35H8.42a2 2 0 0 0-1.35.61A1.8 1.8 0 0 0 6.54 12h-1a2.76 2.76 0 0 1 .81-2a3 3 0 0 1 2-.93z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSymlinkFile(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.57 1.14l3.28 3.3l.15.36v9.7l-.5.5H10v-1h3V6H9.5L9 5.5V2H3v4H2V1.5l.5-.5h7.72zM10 5h3l-3-3zM8.5 7h-7l-.5.5v7l.5.5h7l.5-.5v-7zM8 14H2V8h6zM7 9.5v3H6v-1.793l-2.646 2.647l-.708-.708L5.293 10H3.53V9H6.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileZip(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 1h11l.5.5v5l-.15.35l-.85.86v6.79l-.5.5h-10l-.5-.5v-13zM6 2H5v2h1zm0 12h4V7.68l-.85-.85L9 6.47V2H7v2.5l-.5.5H6v1H5V5h-.5L4 4.5V2H3v12h2v-1h1zm0-2v1h1v-1zm0-1v1H5v-1zm0-1h1v1H6zm0-1v1H5V9zm0-1h1v1H6zm0-1v1H5V7zm0 0h1V6H6zm6.15.15l.85-.86V2h-3v4.27l.85.85l.15.35V14h1V7.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Files(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 0h-9L7 1.5V6H2.5L1 7.5v15.07L2.5 24h12.07L16 22.57V18h4.7l1.3-1.43V4.5zm0 2.12l2.38 2.38H17.5zm-3 20.38h-12v-15H7v9.07L8.5 18h6zm6-6h-12v-15H16V6h4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 2v1.67l-5 4.759V14H6V8.429l-5-4.76V2zM7 8v5h2V8l5-4.76V3H2v.24z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 2v1.67l-5 4.759V14H6V8.429l-5-4.76V2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flame(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.13 15l-.53-.77a1.85 1.85 0 0 0-.28-2.54a3.51 3.51 0 0 1-1.19-2c-1.56 2.23-.75 3.46 0 4.55l-.55.76A4.4 4.4 0 0 1 3 10.46S2.79 8.3 5.28 6.19c0 0 2.82-2.61 1.84-4.54L7.83 1a6.57 6.57 0 0 1 2.61 6.94a2.57 2.57 0 0 0 .56-.81l.87-.07c.07.12 1.84 2.93.89 5.3A4.72 4.72 0 0 1 9.13 15m-2-6.95l.87.39a3 3 0 0 0 .92 2.48a2.64 2.64 0 0 1 1 2.8A3.241 3.241 0 0 0 11.8 12a4.87 4.87 0 0 0-.41-3.63a1.85 1.85 0 0 1-1.84.86l-.35-.68a5.31 5.31 0 0 0-.89-5.8C8.17 4.87 6 6.83 5.93 6.94C3.86 8.7 4 10.33 4 10.4a3.47 3.47 0 0 0 1.59 3.14C5 12.14 5 10.46 7.16 8.05z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fold(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.952 2.381L7.976 6.357L4 2.381L3.38 3l4.286 4.285h.619L12.57 3zM3.904 14l4.072-4.072L12.047 14l.62-.619L8.284 9h-.619l-4.381 4.381z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FoldDown(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.207 1.707L13.5 1l-6 6l-6-6l-.707.707l6.353 6.354h.708zm0 6L13.5 7l-6 6l-6-6l-.707.707l6.353 6.354h.708z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FoldUp(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1 7.4l.7.7l6-6l6 6l.7-.7L8.1 1h-.7zm0 6l.7.7l6-6l6 6l.7-.7L8.1 7h-.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 3H7.71l-.85-.85L6.51 2h-5l-.5.5v11l.5.5h13l.5-.5v-10zm-.51 8.49V13h-12V7h4.49l.35-.15l.86-.86H14v1.5zm0-6.49h-6.5l-.35.15l-.86.86H2v-3h4.29l.85.85l.36.15H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderActive(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.71 3h6.79l.51.5v4.507A4.997 4.997 0 0 0 14 7.416V5.99H7.69l-.86.86l-.35.15H1.99v6H7.1c.07.348.177.682.316 1H1.51l-.5-.5v-11l.5-.5h5l.35.15zm-.22 2h6.5l.01-.99H7.5l-.36-.15l-.85-.85H2v3h4.28l.86-.86z"/><path d="M9.778 8.674a4 4 0 1 1 4.444 6.652a4 4 0 0 1-4.444-6.652m2.13 4.99l2.387-3.182l-.8-.6l-2.077 2.769l-1.301-1.041l-.625.78l1.704 1.364l.713-.09z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderLibrary(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.71 3h6.79l.51.5V7H14V5.99H7.69l-.86.86l-.35.15H1.99v6H7v1H1.51l-.5-.5v-11l.5-.5h5l.35.15zm-.22 2h6.5l.01-.99H7.5l-.36-.15l-.85-.85H2v3h4.28l.86-.86z" clip-rule="evenodd"/><path d="M8 8h1v6H8zm2 0h1v6h-1zm2.004.352l.94-.342l2.052 5.638l-.94.342z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpened(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 14h11l.48-.37l2.63-7l-.48-.63H14V3.5l-.5-.5H7.71l-.86-.85L6.5 2h-5l-.5.5v11zM2 3h4.29l.86.85l.35.15H13v2H8.5l-.35.15l-.86.85H3.5l-.47.34l-1 3.08zm10.13 10H2.19l1.67-5H7.5l.35-.15l.86-.85h5.79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Game(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M4 3h8a4 4 0 0 1 4 4v3a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4m0 1a3 3 0 0 0-3 3v3a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3z" clip-rule="evenodd"/><path d="M5.5 6a.5.5 0 0 0-.5.5V8H3.5a.5.5 0 0 0 0 1H5v1.5a.5.5 0 0 0 1 0V9h1.5a.5.5 0 0 0 0-1H6V6.5a.5.5 0 0 0-.5-.5M13 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gear(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.1 4.4L8.6 2H7.4l-.5 2.4l-.7.3l-2-1.3l-.9.8l1.3 2l-.2.7l-2.4.5v1.2l2.4.5l.3.8l-1.3 2l.8.8l2-1.3l.8.3l.4 2.3h1.2l.5-2.4l.8-.3l2 1.3l.8-.8l-1.3-2l.3-.8l2.3-.4V7.4l-2.4-.5l-.3-.8l1.3-2l-.8-.8l-2 1.3zM9.4 1l.5 2.4L12 2.1l2 2l-1.4 2.1l2.4.4v2.8l-2.4.5L14 12l-2 2l-2.1-1.4l-.5 2.4H6.6l-.5-2.4L4 13.9l-2-2l1.4-2.1L1 9.4V6.6l2.4-.5L2.1 4l2-2l2.1 1.4l.4-2.4zm.6 7c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2M8 9c.6 0 1-.4 1-1s-.4-1-1-1s-1 .4-1 1s.4 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 4h-1.6c.1-.4.1-.8.1-1.2c-.1-.3-.2-.6-.4-.9c-.2-.3-.4-.5-.7-.6c-.3-.1-.6-.3-.9-.3c-.3 0-.6 0-.9.2c-.7.2-1.2.7-1.6 1.3c-.4-.6-.9-1.1-1.6-1.3c-.3-.1-.6-.2-.9-.2c-.3 0-.6.1-.9.3c-.3.1-.5.3-.7.6c-.2.2-.3.6-.4.9c0 .4 0 .8.1 1.2H1.5l-.5.5v9l.5.5h12l.5-.5v-9zM7 13H2V5h5zm0-9H4v-.2c-.1-.3-.1-.5-.1-.8c.1-.2.1-.4.3-.5c.1-.2.3-.3.5-.4c.1-.1.3-.1.5-.1s.4 0 .6.1c.3.1.6.3.8.6c.2.3.4.6.4 1zm1-.3c0-.4.2-.7.4-1c.2-.3.5-.5.8-.6c.2-.1.4-.1.6-.1c.2 0 .4 0 .6.1c.2.1.3.2.5.4c.1.1.1.3.2.5c0 .3 0 .5-.1.8c0 .1 0 .1-.1.2H8zm5 9.3H8V5h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gist(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.57 1.14l3.28 3.3l.15.36v9.7l-.5.5h-11l-.5-.5v-13l.5-.5h7.72zM10 5h3l-3-3zM3 2v12h10V6H9.5L9 5.5V2zm2.062 7.533l1.817-1.828L6.17 7L4 9.179v.707l2.171 2.174l.707-.707zM8.8 7.714l.7-.709l2.189 2.175v.709L9.5 12.062l-.705-.709l1.831-1.82z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GistSecret(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 14h4v.91l.09.09H2.5l-.5-.5v-13l.5-.5h7.72l.35.14l3.28 3.3l.15.36v2.54a3.1 3.1 0 0 0-1-.94V6H9.5L9 5.5V2H3zm10-9l-3-3v3zm.5 4v1h1l.5.5v4l-.5.5h-6l-.5-.5v-4l.5-.5h1V9a2 2 0 0 1 4 0m-2.707-.707A1 1 0 0 0 10.5 9v1h2V9a1 1 0 0 0-1.707-.707M9 11v3h5v-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCommit(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0h1v4.03a4 4 0 0 1 0 7.94V16h-1v-4.03a4 4 0 0 1 0-7.94zM8 10.6a2.6 2.6 0 1 0 0-5.2a2.6 2.6 0 0 0 0 5.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCompare(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.389 12.99l-1.27-1.27l.67-.7l2.13 2.13v.7l-2.13 2.13l-.71-.71L7.349 14h-1.85a2.49 2.49 0 0 1-2.5-2.5V5.95a2.59 2.59 0 0 1-1.27-.68a2.52 2.52 0 0 1-.54-2.73A2.5 2.5 0 0 1 3.499 1a2.45 2.45 0 0 1 1 .19a2.48 2.48 0 0 1 1.35 1.35c.133.317.197.658.19 1a2.5 2.5 0 0 1-2 2.45v5.5a1.5 1.5 0 0 0 1.5 1.5zm-4.68-8.25a1.5 1.5 0 0 0 2.08-2.08a1.55 1.55 0 0 0-.68-.56a1.49 1.49 0 0 0-.86-.08a1.49 1.49 0 0 0-1.18 1.18a1.49 1.49 0 0 0 .08.86c.117.277.311.513.56.68m10.33 6.3c.48.098.922.335 1.27.68a2.51 2.51 0 0 1 .31 3.159a2.5 2.5 0 1 1-3.47-3.468c.269-.182.571-.308.89-.37V5.49a1.5 1.5 0 0 0-1.5-1.5h-1.85l1.27 1.27l-.71.71l-2.13-2.13v-.7l2.13-2.13l.71.71l-1.27 1.27h1.85a2.49 2.49 0 0 1 2.5 2.5zm-.351 3.943a1.5 1.5 0 0 0 1.1-2.322a1.55 1.55 0 0 0-.68-.56a1.49 1.49 0 0 0-.859-.08a1.49 1.49 0 0 0-1.18 1.18a1.49 1.49 0 0 0 .08.86a1.5 1.5 0 0 0 1.539.922" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitFetch(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 2H8v1h1zm-.854 12l-5-5l.708-.707L8 12.439V11h1v1.44l4.146-4.147l.707.707l-5 5zM8 5h1v1H8z"/><path d="M9 8H8v1h1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitMerge(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.273 7.73a2.51 2.51 0 0 0-3.159-.31a2.5 2.5 0 0 0-.921 1.12a2.23 2.23 0 0 0-.13.44a4.52 4.52 0 0 1-4-4a2.23 2.23 0 0 0 .44-.13a2.5 2.5 0 0 0 1.54-2.31a2.45 2.45 0 0 0-.19-1A2.48 2.48 0 0 0 5.503.19a2.45 2.45 0 0 0-1-.19a2.5 2.5 0 0 0-2.31 1.54a2.52 2.52 0 0 0 .54 2.73c.35.343.79.579 1.27.68v5.1a2.411 2.411 0 0 0-.89.37a2.5 2.5 0 1 0 3.47 3.468a2.5 2.5 0 0 0 .42-1.387a2.45 2.45 0 0 0-.19-1a2.48 2.48 0 0 0-1.81-1.49v-2.4a5.52 5.52 0 0 0 2 1.73a5.65 5.65 0 0 0 2.09.6a2.5 2.5 0 0 0 4.95-.49a2.51 2.51 0 0 0-.77-1.72zm-8.2 3.38c.276.117.512.312.68.56a1.5 1.5 0 0 1-2.08 2.08a1.55 1.55 0 0 1-.56-.68a1.49 1.49 0 0 1-.08-.86a1.49 1.49 0 0 1 1.18-1.18a1.49 1.49 0 0 1 .86.08M4.503 4a1.5 1.5 0 0 1-1.39-.93a1.49 1.49 0 0 1-.08-.86a1.49 1.49 0 0 1 1.18-1.18a1.49 1.49 0 0 1 .86.08A1.5 1.5 0 0 1 4.503 4m8.06 6.56a1.5 1.5 0 0 1-2.45-.49a1.49 1.49 0 0 1-.08-.86a1.49 1.49 0 0 1 1.18-1.18a1.49 1.49 0 0 1 .86.08a1.499 1.499 0 0 1 .49 2.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequest(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.616 4.928a2.487 2.487 0 0 1-1.119.922c-.148.06-.458.138-.458.138v5.008a2.51 2.51 0 0 1 1.579 1.062c.273.412.419.895.419 1.388c.008.343-.057.684-.19 1A2.485 2.485 0 0 1 3.5 15.984a2.482 2.482 0 0 1-1.388-.419A2.487 2.487 0 0 1 1.05 13c.095-.486.331-.932.68-1.283c.349-.343.79-.579 1.269-.68V5.949a2.6 2.6 0 0 1-1.269-.68a2.503 2.503 0 0 1-.68-1.283a2.487 2.487 0 0 1 1.06-2.565A2.49 2.49 0 0 1 3.5 1a2.504 2.504 0 0 1 1.807.729a2.493 2.493 0 0 1 .729 1.81c.002.494-.144.978-.42 1.389m-.756 7.861a1.5 1.5 0 0 0-.552-.579a1.45 1.45 0 0 0-.77-.21a1.495 1.495 0 0 0-1.47 1.79a1.493 1.493 0 0 0 1.18 1.179c.288.058.586.03.86-.08c.276-.117.512-.312.68-.56c.15-.226.235-.49.249-.76a1.51 1.51 0 0 0-.177-.78M2.708 4.741c.247.161.536.25.83.25c.271 0 .538-.075.77-.211a1.514 1.514 0 0 0 .729-1.359a1.513 1.513 0 0 0-.25-.76a1.551 1.551 0 0 0-.68-.56a1.49 1.49 0 0 0-.86-.08a1.494 1.494 0 0 0-1.179 1.18c-.058.288-.03.586.08.86c.117.276.312.512.56.68m10.329 6.296c.48.097.922.335 1.269.68c.466.47.729 1.107.725 1.766c.002.493-.144.977-.42 1.388a2.499 2.499 0 0 1-4.532-.899a2.5 2.5 0 0 1 1.067-2.565c.267-.183.571-.308.889-.37V5.489a1.5 1.5 0 0 0-1.5-1.499H8.687l1.269 1.27l-.71.709L7.117 3.84v-.7l2.13-2.13l.71.711l-1.269 1.27h1.85a2.484 2.484 0 0 1 2.312 1.541c.125.302.189.628.187.957zm.557 3.509a1.493 1.493 0 0 0 .191-1.89a1.552 1.552 0 0 0-.68-.559a1.49 1.49 0 0 0-.86-.08a1.493 1.493 0 0 0-1.179 1.18a1.49 1.49 0 0 0 .08.86a1.496 1.496 0 0 0 2.448.49z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequestClosed(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.616 4.928a2.487 2.487 0 0 1-1.119.922c-.148.06-.458.138-.458.138v5.008a2.51 2.51 0 0 1 1.579 1.062c.273.412.419.895.419 1.388c.008.343-.057.684-.19 1A2.485 2.485 0 0 1 3.5 15.984a2.48 2.48 0 0 1-1.388-.42A2.486 2.486 0 0 1 1.05 13c.095-.486.331-.932.68-1.283c.349-.343.79-.579 1.269-.68V5.949a2.6 2.6 0 0 1-1.269-.68a2.503 2.503 0 0 1-.68-1.283a2.487 2.487 0 0 1 1.06-2.565A2.49 2.49 0 0 1 3.5 1a2.504 2.504 0 0 1 1.807.729a2.493 2.493 0 0 1 .729 1.81c.002.494-.144.977-.42 1.389m-.756 7.861a1.5 1.5 0 0 0-.552-.579a1.45 1.45 0 0 0-.77-.21a1.495 1.495 0 0 0-1.47 1.79a1.493 1.493 0 0 0 1.18 1.179c.288.058.586.03.86-.08c.276-.117.512-.312.68-.56c.15-.226.235-.49.249-.76a1.51 1.51 0 0 0-.177-.78M2.708 4.741c.247.161.536.25.83.25c.271 0 .538-.075.77-.211a1.514 1.514 0 0 0 .729-1.36a1.513 1.513 0 0 0-.25-.76a1.551 1.551 0 0 0-.68-.559a1.49 1.49 0 0 0-.86-.08a1.494 1.494 0 0 0-1.179 1.18c-.058.288-.03.586.08.86c.117.276.312.512.56.68m10.329 6.296c.48.097.922.335 1.269.68c.466.47.729 1.107.725 1.766c.002.493-.144.977-.42 1.388a2.5 2.5 0 0 1-3.848.384a2.5 2.5 0 0 1 .382-3.848a2.39 2.39 0 0 1 .89-.37V7.489h1.002zm.557 3.508a1.493 1.493 0 0 0 .191-1.888a1.551 1.551 0 0 0-.68-.56a1.49 1.49 0 0 0-.86-.08a1.493 1.493 0 0 0-1.179 1.18a1.49 1.49 0 0 0 .08.86a1.497 1.497 0 0 0 2.448.49M11.688 3.4L10 5.088l.707.707l1.688-1.688l1.687 1.688l.707-.707L13.103 3.4l1.688-1.687l-.708-.707l-1.687 1.687l-1.688-1.687l-.707.707z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequestCreate(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.616 4.928a2.487 2.487 0 0 1-1.119.922c-.148.06-.458.138-.458.138v5.008a2.51 2.51 0 0 1 1.579 1.062c.273.412.419.895.419 1.388c.008.343-.057.684-.19 1A2.485 2.485 0 0 1 3.5 15.984a2.482 2.482 0 0 1-1.388-.419A2.487 2.487 0 0 1 1.05 13c.095-.486.331-.932.68-1.283c.349-.343.79-.579 1.269-.68V5.949a2.6 2.6 0 0 1-1.269-.68a2.503 2.503 0 0 1-.68-1.283a2.487 2.487 0 0 1 1.06-2.565A2.49 2.49 0 0 1 3.5 1a2.504 2.504 0 0 1 1.807.729a2.493 2.493 0 0 1 .729 1.81c.002.494-.144.978-.42 1.389m-.756 7.861a1.5 1.5 0 0 0-.552-.579a1.45 1.45 0 0 0-.77-.21a1.495 1.495 0 0 0-1.47 1.79a1.493 1.493 0 0 0 1.18 1.179c.288.058.586.03.86-.08c.276-.117.512-.312.68-.56c.15-.226.235-.49.249-.76a1.51 1.51 0 0 0-.177-.78M2.708 4.741c.247.161.536.25.83.25c.271 0 .538-.075.77-.211a1.514 1.514 0 0 0 .729-1.359a1.513 1.513 0 0 0-.25-.76a1.551 1.551 0 0 0-.68-.56a1.49 1.49 0 0 0-.86-.08a1.494 1.494 0 0 0-1.179 1.18c-.058.288-.03.586.08.86c.117.276.312.512.56.68M13.037 7h-1.002V5.49a1.5 1.5 0 0 0-1.5-1.5H8.687l1.269 1.27l-.71.709L7.117 3.84v-.7l2.13-2.13l.71.711l-1.269 1.27h1.85a2.484 2.484 0 0 1 2.312 1.541c.125.302.189.628.187.957zM13 16h-1v-3H9v-1h3V9h1v3h3v1h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequestDraft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4.497 5.85a2.496 2.496 0 0 0 1.538-2.31a2.493 2.493 0 0 0-1.538-2.35A2.504 2.504 0 0 0 3.499 1a2.49 2.49 0 0 0-1.388.42A2.487 2.487 0 0 0 1.05 3.987c.095.486.331.932.68 1.283a2.6 2.6 0 0 0 1.269.68v5.088c-.48.101-.92.337-1.269.68c-.349.35-.585.797-.68 1.283a2.486 2.486 0 0 0 1.062 2.565a2.48 2.48 0 0 0 1.388.419a2.44 2.44 0 0 0 1-.19a2.485 2.485 0 0 0 1.538-2.349a2.51 2.51 0 0 0-1.998-2.45V5.989s.31-.078.458-.138m-.189 6.36a1.5 1.5 0 0 1 .48 2.12a1.551 1.551 0 0 1-.68.559a1.492 1.492 0 0 1-.86.08a1.487 1.487 0 0 1-1.18-1.18a1.49 1.49 0 0 1 .08-.86c.117-.276.312-.512.56-.68a1.49 1.49 0 0 1 .83-.25c.271-.003.538.07.77.211m-.77-7.22a1.52 1.52 0 0 1-.83-.25a1.551 1.551 0 0 1-.56-.68a1.491 1.491 0 0 1-.08-.86a1.486 1.486 0 0 1 1.18-1.179a1.49 1.49 0 0 1 .86.08c.276.117.512.312.68.56A1.49 1.49 0 0 1 4.86 4.2a1.52 1.52 0 0 1-.552.579c-.232.136-.499.21-.77.21"/><path fill-rule="evenodd" d="M15.054 13.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m-2.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3" clip-rule="evenodd"/><circle cx="12.554" cy="7.751" r="1"/><circle cx="12.554" cy="3.501" r="1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequestGoToChanges(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3 10v4l1 1h9l1-1V5l-.29-.71l-3-3L10 1H8v1h2l3 3v9H4v-4zm8-4H9V4H8v2H6v1h2v2h1V7h2zm-5 5h5v1H6z"/><path d="M7.06 3.854L4.915 6l-.707-.707L5.5 4h-3a1.5 1.5 0 0 0 0 3H3v1h-.5a2.5 2.5 0 1 1 0-5h3L4.207 1.707L4.914 1l2.147 2.146z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequestNewChanges(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="m13.71 4.29l-3-3L10 1H4L3 2v12l1 1h5.354a4.019 4.019 0 0 1-.819-1H4V2h6l3 3v3.126c.355.091.69.23 1 .41V5zM8.126 11H6v1h2c0-.345.044-.68.126-1M6 6h2V4h1v2h2v1H9v2H8V7H6z" clip-rule="evenodd"/><path d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0a12 12 0 1 0 0 24a12 12 0 0 0 0-24m3.163 21.783h-.093a.513.513 0 0 1-.382-.14a.513.513 0 0 1-.14-.372v-1.406c.006-.467.01-.94.01-1.416a3.693 3.693 0 0 0-.151-1.028a1.832 1.832 0 0 0-.542-.875a8.014 8.014 0 0 0 2.038-.471a4.051 4.051 0 0 0 1.466-.964c.407-.427.71-.943.885-1.506a6.77 6.77 0 0 0 .3-2.13a4.138 4.138 0 0 0-.26-1.476a3.892 3.892 0 0 0-.795-1.284a2.81 2.81 0 0 0 .162-.582c.033-.2.05-.402.05-.604c0-.26-.03-.52-.09-.773a5.309 5.309 0 0 0-.221-.763a.293.293 0 0 0-.111-.02h-.11c-.23.002-.456.04-.674.111a5.34 5.34 0 0 0-.703.26a6.503 6.503 0 0 0-.661.343c-.215.127-.405.249-.573.362a9.578 9.578 0 0 0-5.143 0a13.507 13.507 0 0 0-.572-.362a6.022 6.022 0 0 0-.672-.342a4.516 4.516 0 0 0-.705-.261a2.203 2.203 0 0 0-.662-.111h-.11a.29.29 0 0 0-.11.02a5.844 5.844 0 0 0-.23.763c-.054.254-.08.513-.081.773c0 .202.017.404.051.604c.033.199.086.394.16.582A3.888 3.888 0 0 0 5.702 10a4.142 4.142 0 0 0-.263 1.476a6.871 6.871 0 0 0 .292 2.12c.181.563.483 1.08.884 1.516c.415.422.915.75 1.466.964c.653.25 1.337.41 2.033.476a1.828 1.828 0 0 0-.452.633a2.99 2.99 0 0 0-.2.744a2.754 2.754 0 0 1-1.175.27a1.788 1.788 0 0 1-1.065-.3a2.904 2.904 0 0 1-.752-.824a3.1 3.1 0 0 0-.292-.382a2.693 2.693 0 0 0-.372-.343a1.841 1.841 0 0 0-.432-.24a1.2 1.2 0 0 0-.481-.101c-.04.001-.08.005-.12.01a.649.649 0 0 0-.162.02a.408.408 0 0 0-.13.06a.116.116 0 0 0-.06.1a.33.33 0 0 0 .14.242c.093.074.17.131.232.171l.03.021c.133.103.261.214.382.333c.112.098.213.209.3.33c.09.119.168.246.231.381c.073.134.15.288.231.463c.188.474.522.875.954 1.145c.453.243.961.364 1.476.351c.174 0 .349-.01.522-.03c.172-.028.343-.057.515-.091v1.743a.5.5 0 0 1-.533.521h-.062a10.286 10.286 0 1 1 6.324 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubAction(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.04 10h2.58l.65 1H2.54l-.5-.5v-9l.5-.5h12l.5.5v4.77l-1-1.75V2h-11zm5.54 1l-1.41 3.47h2.2L15 8.7L14.27 7h-1.63l.82-1.46L12.63 4H9.76l-.92.59l-2.28 5L7.47 11zm1.18-6h2.87l-1.87 3h3.51l-5.76 5.84L10.2 10H7.47zM6.95 7H4.04V6H7.4zm-.9 2H4.04V8H6.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubAlt(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.035 5.257c.91 1.092 1.364 2.366 1.364 3.822c0 5.277-3.002 6.824-5.823 7.279c.364.637.455 1.365.455 2.093v3.73c0 .455-.273.728-.637.728a.718.718 0 0 1-.728-.728v-3.73a2.497 2.497 0 0 0-.728-2.093l.455-1.183c2.821-.364 5.733-1.274 5.733-6.187c0-1.183-.455-2.275-1.274-3.185l-.182-.727a4.04 4.04 0 0 0 .09-2.73c-.454.09-1.364.273-2.91 1.365l-.547.09a13.307 13.307 0 0 0-6.55 0l-.547-.09C7.57 2.71 6.66 2.437 6.204 2.437c-.273.91-.273 1.91.09 2.73l-.181.727c-.91.91-1.365 2.093-1.365 3.185c0 4.822 2.73 5.823 5.732 6.187l.364 1.183c-.546.546-.819 1.274-.728 2.002v3.821a.718.718 0 0 1-.728.728a.718.718 0 0 1-.728-.728V20.18c-3.002.637-4.185-.91-5.095-2.092c-.455-.546-.819-1.001-1.274-1.092c-.09-.091-.364-.455-.273-.819c.091-.364.455-.637.82-.455c.91.182 1.455.91 2 1.547c.82 1.092 1.639 2.092 4.095 1.547v-.364c-.09-.728.091-1.456.455-2.093c-2.73-.546-5.914-2.093-5.914-7.279c0-1.456.455-2.73 1.365-3.822c-.273-1.273-.182-2.638.273-3.73l.455-.364C5.749 1.073 7.023.8 9.66 2.437a13.673 13.673 0 0 1 6.642 0C18.851.708 20.216.98 20.398 1.072l.455.364c.455 1.274.546 2.548.182 3.821"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubInverted(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.976 0A7.977 7.977 0 0 0 0 7.976c0 3.522 2.3 6.507 5.431 7.584c.392.049.538-.196.538-.392v-1.37c-2.201.49-2.69-1.076-2.69-1.076c-.343-.93-.881-1.175-.881-1.175c-.734-.489.048-.489.048-.489c.783.049 1.224.832 1.224.832c.734 1.223 1.859.88 2.3.685c.048-.538.293-.88.489-1.076c-1.762-.196-3.621-.881-3.621-3.964c0-.88.293-1.566.832-2.153c-.05-.147-.343-.978.098-2.055c0 0 .685-.196 2.201.832c.636-.196 1.322-.245 2.007-.245s1.37.098 2.006.245c1.517-1.027 2.202-.832 2.202-.832c.44 1.077.146 1.908.097 2.104a3.16 3.16 0 0 1 .832 2.153c0 3.083-1.86 3.719-3.62 3.915c.293.244.538.733.538 1.467v2.202c0 .196.146.44.538.392A7.984 7.984 0 0 0 16 7.976C15.951 3.572 12.38 0 7.976 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 1a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13m4.894 4a5.527 5.527 0 0 0-3.053-2.676c.444.84.765 1.74.953 2.676zm.582 2.995A5.11 5.11 0 0 0 14 7.5a5.464 5.464 0 0 0-.213-1.5h-2.342c.032.331.055.664.055 1a10.114 10.114 0 0 1-.206 2h2.493c.095-.329.158-.665.19-1.005zm-3.535 0l.006-.051A9.04 9.04 0 0 0 10.5 7a8.994 8.994 0 0 0-.076-1H6.576A8.82 8.82 0 0 0 6.5 7a8.98 8.98 0 0 0 .233 2h3.534c.077-.332.135-.667.174-1.005M10.249 5a8.974 8.974 0 0 0-1.255-2.97C8.83 2.016 8.666 2 8.5 2a3.62 3.62 0 0 0-.312.015l-.182.015L8 2.04A8.97 8.97 0 0 0 6.751 5zM5.706 5a9.959 9.959 0 0 1 .966-2.681A5.527 5.527 0 0 0 3.606 5zM3.213 6A5.48 5.48 0 0 0 3 7.5A5.48 5.48 0 0 0 3.213 9h2.493A10.016 10.016 0 0 1 5.5 7c0-.336.023-.669.055-1zm2.754 4h-2.36a5.515 5.515 0 0 0 3.819 2.893A10.023 10.023 0 0 1 5.967 10M8.5 12.644A8.942 8.942 0 0 0 9.978 10H7.022A8.943 8.943 0 0 0 8.5 12.644M11.033 10a10.024 10.024 0 0 1-1.459 2.893A5.517 5.517 0 0 0 13.393 10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoToFile(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6 5.914l2.06-2.06v-.708L5.915 1l-.707.707l.043.043l.25.25l1 1h-3a2.5 2.5 0 0 0 0 5H4V7h-.5a1.5 1.5 0 1 1 0-3h3L5.207 5.293L5.914 6zM11 2H8.328l-1-1H12l.71.29l3 3L16 5v9l-1 1H6l-1-1V6.5l1 .847V14h9V6h-4zm1 0v3h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grabber(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 6H1v1h14zm0 3H1v1h14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Graph(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 14H15v-1H2V0H1v13.5zM3 11.5v-8l.5-.5h2l.5.5v8l-.5.5h-2zm2-.5V4H4v7zm6-9.5v10l.5.5h2l.5-.5v-10l-.5-.5h-2zm2 .5v9h-1V2zm-6 9.5v-6l.5-.5h2l.5.5v6l-.5.5h-2zm2-.5V6H8v5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphLeft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.119 4L3 4.881l-.619.619L.715 3.833v-.618L2.38 1.548l.62.619L2.167 3H15v1zM4 14.546V5.455L4.5 5h2l.5.455v9.09L6.5 15h-2zm2-.455V5.909H5v8.182zm2-1.535V5.444L8.5 5h2l.5.444v7.112l-.5.444h-2zm2-.445V5.89H9v6.222zm2-6.682v5.143l.5.428h2l.5-.428V5.429L14.5 5h-2zm2 .428v4.286h-1V5.857z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphLine(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15 13v1H1.5l-.5-.5V0h1v13z"/><path d="M13 3.207L7.854 8.354h-.708L5.5 6.707l-3.646 3.647l-.708-.708l4-4h.708L7.5 7.293l5.146-5.147h.707l2 2l-.707.708z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphScatter(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15 13v1H1.5l-.5-.5V0h1v13z"/><path d="M5 2h2v2H5zm7-1h2v2h-2zM8 5h2v2H8zM5 9h2v2H5zm7-1h2v2h-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gripper(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3h2v2H5zm0 4h2v2H5zm0 4h2v2H5zm4-8h2v2H9zm0 4h2v2H9zm0 4h2v2H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupByRefType(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 1h2v1H2v12h1.5v1h-2l-.5-.5v-13zm6 6h-2L5 6.5v-2l.5-.5h2l.5.5v2zM6 6h1V5H6zm7.5 1h-3l-.5-.5v-3l.5-.5h3l.5.5v3zM11 6h2V4h-2zm-3.5 6h-2l-.5-.5v-2l.5-.5h2l.5.5v2zM6 11h1v-1H6zm7.5 2h-3l-.5-.5v-3l.5-.5h3l.5.5v3zM11 12h2v-2h-2zm-1-2H8v1h2zm0-5H8v1h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.88 4.78a3.489 3.489 0 0 0-.37-.9a3.24 3.24 0 0 0-.6-.79a3.78 3.78 0 0 0-1.21-.81a3.74 3.74 0 0 0-2.84 0a4 4 0 0 0-1.16.75l-.05.06l-.65.65l-.65-.65l-.05-.06a4 4 0 0 0-1.16-.75a3.74 3.74 0 0 0-2.84 0a3.78 3.78 0 0 0-1.21.81a3.55 3.55 0 0 0-.97 1.69a3.75 3.75 0 0 0-.12 1c0 .317.04.633.12.94a4 4 0 0 0 .36.89a3.8 3.8 0 0 0 .61.79L8 14.31l5.91-5.91c.237-.233.44-.5.6-.79A3.578 3.578 0 0 0 15 5.78a3.747 3.747 0 0 0-.12-1m-1 1.63a2.69 2.69 0 0 1-.69 1.21l-5.21 5.2l-5.21-5.2a2.9 2.9 0 0 1-.44-.57a3 3 0 0 1-.27-.65a3.25 3.25 0 0 1-.08-.69A3.36 3.36 0 0 1 2.06 5a2.8 2.8 0 0 1 .27-.65c.12-.21.268-.4.44-.57a2.91 2.91 0 0 1 .89-.6a2.8 2.8 0 0 1 2.08 0c.33.137.628.338.88.59l1.36 1.37l1.36-1.37a2.72 2.72 0 0 1 .88-.59a2.8 2.8 0 0 1 2.08 0c.331.143.633.347.89.6c.174.165.32.357.43.57a2.69 2.69 0 0 1 .35 1.34a2.6 2.6 0 0 1-.06.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.88 4.78a3.489 3.489 0 0 0-.37-.9a3.24 3.24 0 0 0-.6-.79a3.78 3.78 0 0 0-1.21-.81a3.74 3.74 0 0 0-2.84 0a4 4 0 0 0-1.16.75l-.05.06l-.65.65l-.65-.65l-.05-.06a4 4 0 0 0-1.16-.75a3.74 3.74 0 0 0-2.84 0a3.78 3.78 0 0 0-1.21.81a3.55 3.55 0 0 0-.97 1.69a3.75 3.75 0 0 0-.12 1c0 .318.04.634.12.94a4 4 0 0 0 .36.89a3.8 3.8 0 0 0 .61.79L8 14.31l5.91-5.91c.237-.232.44-.498.6-.79A3.578 3.578 0 0 0 15 5.78a3.747 3.747 0 0 0-.12-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.507 12.324a7 7 0 0 0 .065-8.56A7 7 0 0 0 2 4.393V2H1v3.5l.5.5H5V5H2.811a6.008 6.008 0 1 1-.135 5.77l-.887.462a7 7 0 0 0 11.718 1.092m-3.361-.97l.708-.707L8 7.792V4H7v4l.146.354z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.36 1.37l6.36 5.8l-.71.71L13 6.964v6.526l-.5.5h-3l-.5-.5v-3.5H7v3.5l-.5.5h-3l-.5-.5V6.972L2 7.88l-.71-.71l6.35-5.8zM4 6.063v6.927h2v-3.5l.5-.5h3l.5.5v3.5h2V6.057L8 2.43z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalRule(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.432 10h.823V4h-.823v2.61h-2.61V4H3v6h.823V7.394h2.61zm5.668 0h.9l-1.28-2.63c.131-.058.26-.134.389-.23a1.666 1.666 0 0 0 .585-.797c.064-.171.096-.364.096-.58a1.77 1.77 0 0 0-.082-.557a1.644 1.644 0 0 0-.22-.446a1.504 1.504 0 0 0-.31-.341a1.864 1.864 0 0 0-.737-.373A1.446 1.446 0 0 0 11.1 4H8.64v6h.824V7.518h1.467zm-.681-3.32a.874.874 0 0 1-.293.055H9.463V4.787h1.663a.87.87 0 0 1 .576.24a.956.956 0 0 1 .306.737c0 .168-.029.314-.087.437a.91.91 0 0 1-.503.479zM13 12H3v1h10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hubot(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.48 4h4l.5.5v2.03h.52l.5.5V8l-.5.5h-.52v3l-.5.5H9.36l-2.5 2.76L6 14.4V12H3.5l-.5-.64V8.5h-.5L2 8v-.97l.5-.5H3V4.36L3.53 4h4V2.86A1 1 0 0 1 7 2a1 1 0 0 1 2 0a1 1 0 0 1-.52.83zM12 8V5H4v5.86l2.5.14H7v2.19l1.8-2.04l.35-.15H12zm-2.12.51a2.71 2.71 0 0 1-1.37.74v-.01a2.71 2.71 0 0 1-2.42-.74l-.7.71c.34.34.745.608 1.19.79c.45.188.932.286 1.42.29a3.7 3.7 0 0 0 2.58-1.07zM6.49 6.5h-1v1h1zm3 0h1v1h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 14h13l.5-.5V9l-2.77-7.66l-.47-.34H4.27l-.47.33L1 8.74v4.76zM14 13H2v-2.98h2.55l.74 1.25l.43.24h4.57l.44-.26l.69-1.23H14zm-.022-3.98H11.12l-.43.26l-.69 1.23H6.01l-.75-1.25l-.43-.24H2V9l2.62-7h6.78z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Indent(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 5v1.984a.5.5 0 0 0 .5.5h6.882L9.749 5.851l.707-.707l2.121 2.121l.423.423v.568L10.456 10.8l-.707-.707l1.61-1.609H4.5a1.5 1.5 0 0 1-1.5-1.5V5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.568 1.031A6.8 6.8 0 0 1 12.76 3.05a7.06 7.06 0 0 1 .46 9.39a6.85 6.85 0 0 1-8.58 1.74a7 7 0 0 1-3.12-3.5a7.12 7.12 0 0 1-.23-4.71a7 7 0 0 1 2.77-3.79a6.8 6.8 0 0 1 4.508-1.149M9.04 13.88a5.89 5.89 0 0 0 3.41-2.07a6.07 6.07 0 0 0-.4-8.06a5.82 5.82 0 0 0-7.43-.74a6.06 6.06 0 0 0 .5 10.29a5.81 5.81 0 0 0 3.92.58M7.375 6h1.25V5h-1.25zm1.25 1v4h-1.25V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Insert(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="m14 1l1 1v4l-1 1H6L5 6V2l1-1zm0 1H6v4h8zm0 7l1 1v4l-1 1H6l-1-1v-4l1-1zm0 1H6v4h8z" clip-rule="evenodd"/><path d="m1 6.393l1.614 1.614L1 9.62l.694.693L4 8.007L1.694 5.7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inspect(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1 3l1-1h12l1 1v6h-1V3H2v8h5v1H2l-1-1zm14.707 9.707L9 6v9.414l2.707-2.707zM10 13V8.414l3.293 3.293h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueClosed(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" clip-rule="evenodd" d="M7.5 13a5.5 5.5 0 0 0 5.498-5.354l.953-.952a6.5 6.5 0 1 1-1.208-3.038l-.717.718A5.5 5.5 0 1 0 7.5 13zm6.197-7.467L15 4.23l-.71-.71l-3.41 3.42L9.5 5.56l-.71.71L10.52 8h.71l1.65-1.65l.817-.817zM7 7v2h1V4H7v3zm0 4v-1h1v1H7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueDraft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12.732 9.2l.952.31A6.494 6.494 0 0 0 14 7.5c0-.701-.111-1.377-.316-2.01l-.952.31c.174.534.268 1.105.268 1.7s-.094 1.166-.268 1.7m-.33-4.197l.89-.455a6.529 6.529 0 0 0-2.84-2.84l-.455.89a5.528 5.528 0 0 1 2.405 2.405M9.2 2.268l.31-.951A6.495 6.495 0 0 0 7.5 1c-.701 0-1.377.111-2.01.317l.31.95A5.495 5.495 0 0 1 7.5 2c.595 0 1.166.094 1.7.268m-4.197.33l-.455-.89a6.528 6.528 0 0 0-2.84 2.84l.89.455a5.528 5.528 0 0 1 2.405-2.405M1 7.5c0-.701.111-1.377.317-2.01l.95.31A5.495 5.495 0 0 0 2 7.5c0 .595.094 1.166.268 1.7l-.951.31A6.495 6.495 0 0 1 1 7.5m1.598 2.497l-.89.455a6.529 6.529 0 0 0 2.84 2.84l.455-.89a5.528 5.528 0 0 1-2.405-2.405M5.8 12.732l-.31.952A6.494 6.494 0 0 0 7.5 14c.701 0 1.377-.111 2.01-.316l-.31-.952A5.497 5.497 0 0 1 7.5 13a5.497 5.497 0 0 1-1.7-.268m4.197-.33l.455.89a6.53 6.53 0 0 0 2.84-2.84l-.89-.455a5.528 5.528 0 0 1-2.405 2.405M7.5 8.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IssueReopened(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.28 5.656L2 7.006l-.66-.26L0 3.506l.92-.38l.81 1.95a6.48 6.48 0 0 1 12.48 1.93h-1a5.48 5.48 0 0 0-10.64-1.28l2.32-1zm8.86 2.68l1.34 3.23l-.92.44l-.82-2a6.49 6.49 0 0 1-12.5-2h1v-.5a5.49 5.49 0 0 0 10.64 1.89l-2.25.93l-.39-.92l3.25-1.35z" clip-rule="evenodd"/><circle cx="7.74" cy="7.54" r="1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Issues(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.5 1a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13m0 12a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11"/><circle cx="7.5" cy="7.5" r="1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.001 13.593l-.097.325H4l.123-.325c.492-.012.817-.053.976-.123c.257-.1.448-.238.57-.413c.194-.276.394-.768.599-1.477l2.074-7.19c.176-.597.263-1.048.263-1.353a.643.643 0 0 0-.114-.387a.683.683 0 0 0-.351-.237c-.153-.059-.454-.088-.906-.088L7.34 2h4.605l-.096.325c-.375-.006-.654.035-.835.123a1.358 1.358 0 0 0-.607.501c-.134.217-.31.697-.527 1.442l-2.066 7.19c-.187.661-.28 1.083-.28 1.265c0 .146.034.272.105.378c.076.1.193.178.351.237c.164.053.501.097 1.011.132"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jersey(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.91 14.22H4.06l-.5-.5V7.06H2.15l-.48-.38L1 4l.33-.6L5.59 2l.64.32a2.7 2.7 0 0 0 .21.44c.071.103.152.2.24.29c.168.169.369.302.59.39a1.82 1.82 0 0 0 1.43 0a1.74 1.74 0 0 0 .59-.39c.09-.095.173-.195.25-.3l.15-.29a1.21 1.21 0 0 0 .05-.14l.64-.32l4.26 1.42L15 4l-.66 2.66l-.49.38h-1.44v6.66zm-7.35-1h6.85V6.56l.5-.5h1.52l.46-1.83l-3.4-1.14a1.132 1.132 0 0 1-.12.21c-.11.161-.233.312-.37.45a2.75 2.75 0 0 1-.91.61a2.85 2.85 0 0 1-2.22 0A2.92 2.92 0 0 1 6 3.75a2.17 2.17 0 0 1-.36-.44l-.13-.22l-3.43 1.14l.46 1.83h1.52l.5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Json(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2.984V2h-.09c-.313 0-.616.062-.909.185a2.33 2.33 0 0 0-.775.53a2.23 2.23 0 0 0-.493.753v.001a3.542 3.542 0 0 0-.198.83v.002a6.08 6.08 0 0 0-.024.863c.012.29.018.58.018.869c0 .203-.04.393-.117.572v.001a1.504 1.504 0 0 1-.765.787a1.376 1.376 0 0 1-.558.115H2v.984h.09c.195 0 .38.04.556.121l.001.001c.178.078.329.184.455.318l.002.002c.13.13.233.285.307.465l.001.002c.078.18.117.368.117.566c0 .29-.006.58-.018.869c-.012.296-.004.585.024.87v.001c.033.283.099.558.197.824v.001c.106.273.271.524.494.753c.223.23.482.407.775.53c.293.123.596.185.91.185H6v-.984h-.09c-.2 0-.387-.038-.563-.115a1.613 1.613 0 0 1-.457-.32a1.659 1.659 0 0 1-.309-.467c-.074-.18-.11-.37-.11-.573c0-.228.003-.453.011-.672c.008-.228.008-.45 0-.665a4.639 4.639 0 0 0-.055-.64a2.682 2.682 0 0 0-.168-.609A2.284 2.284 0 0 0 3.522 8a2.284 2.284 0 0 0 .738-.955c.08-.192.135-.393.168-.602c.033-.21.051-.423.055-.64c.008-.22.008-.442 0-.666c-.008-.224-.012-.45-.012-.678a1.47 1.47 0 0 1 .877-1.354a1.33 1.33 0 0 1 .563-.121zm4 10.032V14h.09c.313 0 .616-.062.909-.185c.293-.123.552-.3.775-.53c.223-.23.388-.48.493-.753v-.001c.1-.266.165-.543.198-.83v-.002c.028-.28.036-.567.024-.863c-.012-.29-.018-.58-.018-.869c0-.203.04-.393.117-.572v-.001a1.502 1.502 0 0 1 .765-.787a1.38 1.38 0 0 1 .558-.115H14v-.984h-.09c-.196 0-.381-.04-.557-.121l-.001-.001a1.376 1.376 0 0 1-.455-.318l-.002-.002a1.415 1.415 0 0 1-.307-.465v-.002a1.405 1.405 0 0 1-.118-.566c0-.29.006-.58.018-.869a6.174 6.174 0 0 0-.024-.87v-.001a3.537 3.537 0 0 0-.197-.824v-.001a2.23 2.23 0 0 0-.494-.753a2.331 2.331 0 0 0-.775-.53a2.325 2.325 0 0 0-.91-.185H10v.984h.09c.2 0 .387.038.562.115c.174.082.326.188.457.32c.127.134.23.29.309.467c.074.18.11.37.11.573c0 .228-.003.452-.011.672c-.008.228-.008.45 0 .665c.004.222.022.435.055.64c.033.214.089.416.168.609a2.285 2.285 0 0 0 .738.955a2.285 2.285 0 0 0-.738.955a2.689 2.689 0 0 0-.168.602c-.033.21-.051.423-.055.64a9.15 9.15 0 0 0 0 .666c.008.224.012.45.012.678a1.471 1.471 0 0 1-.877 1.354a1.33 1.33 0 0 1-.563.121z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KebabVertical(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.444 13.832a1 1 0 1 0 1.111-1.663a1 1 0 0 0-1.11 1.662zM8 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0-5a1 1 0 1 1 0-2a1 1 0 0 1 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.351 1.091a4.528 4.528 0 0 1 3.44 3.16c.215.724.247 1.49.093 2.23a4.583 4.583 0 0 1-4.437 3.6c-.438 0-.874-.063-1.293-.19l-.8.938l-.379.175H7v1.5l-.5.5H5v1.5l-.5.5h-3l-.5-.5v-2.307l.146-.353L6.12 6.87a4.464 4.464 0 0 1-.2-1.405a4.528 4.528 0 0 1 5.431-4.375zm1.318 7.2a3.568 3.568 0 0 0 1.239-2.005l.004.005A3.543 3.543 0 0 0 9.72 2.08a3.576 3.576 0 0 0-2.8 3.4c-.01.456.07.908.239 1.33l-.11.543L2 12.404v1.6h2v-1.5l.5-.5H6v-1.5l.5-.5h1.245l.876-1.016l.561-.14a3.47 3.47 0 0 0 1.269.238a3.568 3.568 0 0 0 2.218-.795m-.838-2.732a1 1 0 1 0-1.662-1.11a1 1 0 0 0 1.662 1.11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Law(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.63 7L13 3h1V2H9V1H8v1H3v1h1L2.38 7H2v1h.15c.156.498.473.93.9 1.23a2.47 2.47 0 0 0 2.9 0A2.44 2.44 0 0 0 6.86 8H7V7h-.45L4.88 3H8v8H6l-.39.18l-2 2.51l.39.81h9l.39-.81l-2-2.51L11 11H9V3h3.13l-1.67 4H10v1h.15a2.48 2.48 0 0 0 4.71 0H15V7zM5.22 8.51a1.52 1.52 0 0 1-.72.19a1.45 1.45 0 0 1-.71-.19A1.47 1.47 0 0 1 3.25 8h2.5a1.52 1.52 0 0 1-.53.51M5.47 7h-2l1-2.4zm5.29 5L12 13.5H5L6.24 12zm1.78-7.38l1 2.4h-2zm.68 3.91a1.41 1.41 0 0 1-.72.19a1.35 1.35 0 0 1-.71-.19a1.55 1.55 0 0 1-.54-.53h2.5a1.37 1.37 0 0 1-.53.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layers(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.627 1.087l.558-.004l6.091 4.037l-.003.836L8.182 9.92l-.551-.004l-5.91-3.963l-.003-.828zm.286 1.016l-5.021 3.43l5.021 3.368l5.176-3.368zM1.793 8.5l5.838 3.915l.55.004L14.206 8.5h-1.833l-4.459 2.9l-4.325-2.9zm5.838 6.415L1.793 11h1.795l4.325 2.9l4.459-2.9h1.833l-6.023 3.92z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersActive(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.185 1.083l-.558.004l-5.909 4.037l.004.828L7.63 9.915l.55.004l6.092-3.963l.003-.836zm-5.293 4.45l5.021-3.43l5.176 3.43l-5.176 3.368zm4.739 6.882L1.793 8.5h1.795l4.325 2.9l4.459-2.9h1.833l-.8.52a4.002 4.002 0 0 0-4.21 2.739l-1.013.66zm1.373.776l-1.09.71L3.587 11H1.793l5.838 3.915l.55.004l1.02-.663a3.988 3.988 0 0 1-.197-1.065m2.329-2.685a3 3 0 1 1 3.333 4.987a3 3 0 0 1-3.333-4.987m1.698 3.817l1.79-2.387l-.8-.6l-1.48 1.974l-.876-.7l-.624.78l1.278 1.023z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersDot(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="m8.185 1.083l-.558.004l-5.909 4.037l.004.828L7.63 9.915l.55.004l6.092-3.963l.003-.836zm-5.293 4.45l5.021-3.43l5.176 3.43l-5.176 3.368zm4.739 6.882L1.793 8.5h1.795l4.325 2.9l4.459-2.9h1.833l-.8.52a4.002 4.002 0 0 0-4.21 2.739l-1.013.66zm1.373.776l-1.09.71L3.587 11H1.793l5.838 3.915l.55.004l1.02-.663a3.988 3.988 0 0 1-.197-1.065" clip-rule="evenodd"/><circle cx="13" cy="13" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layout(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2L2 3v10l1 1h4l1-1V3L7 2zm0 11V3h4v10zm7-10l1-1h3l1 1v3l-1 1h-3l-1-1zm1 0v3h3V3zm-1 7l1-1h3l1 1v3l-1 1h-3l-1-1zm1 0v3h3v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutActivitybarLeft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm12 13H4V2h10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutActivitybarRight(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 13V2h10v12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutCentered(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 13V2h4v12zm8 0V2h4v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutMenubar(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1 2l1-1h12l1 1v12l-1 1H2l-1-1zm1 0v12h12V2zm1 1h2v1H3zm3 0h2v1H6zm5 0H9v1h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutPanel(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 9V2h12v8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutPanelCenter(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 13V2h2v12zm3-4V2h6v8zm7-8h2v12h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutPanelJustify(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 9V2h2v8zm3 0V2h6v8zm7 0V2h2v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutPanelLeft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1 2l1-1h12l1 1v12l-1 1H2l-1-1zm1 0v8h8V2zm9 0v12h3V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutPanelOff(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 9V2h12v8zm0 1h12v3H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutPanelRight(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1 2l1-1h12l1 1v12l-1 1H2l-1-1zm1 0v12h3V2zm4 0v8h8V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSidebarLeft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm12 13H7V2h7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSidebarLeftOff(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 13V2h4v12zm5 0V2h7v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSidebarRight(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 13V2h7v12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutSidebarRightOff(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 13V2h7v12zm8 0V2h4v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutStatusbar(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1L1 2v12l1 1h12l1-1V2l-1-1zm0 12V2h12v11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Library(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5 2.5l.5-.5h2l.5.5v11l-.5.5h-2l-.5-.5zM6 3v10h1V3zm3.171.345l.299-.641l1.88-.684l.64.299l3.762 10.336l-.299.641l-1.879.684l-.64-.299zm1.11.128l3.42 9.396l.94-.341l-3.42-9.397zM1 2.5l.5-.5h2l.5.5v11l-.5.5h-2l-.5-.5zM2 3v10h1V3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.67 8.658a3.661 3.661 0 0 0-.781 1.114a3.28 3.28 0 0 0-.268 1.329v1.6a1.304 1.304 0 0 1-.794 1.197a1.282 1.282 0 0 1-.509.102H7.712a1.285 1.285 0 0 1-.922-.379a1.303 1.303 0 0 1-.38-.92v-1.6c0-.479-.092-.921-.274-1.329a3.556 3.556 0 0 0-.776-1.114a4.689 4.689 0 0 1-1.006-1.437A4.187 4.187 0 0 1 4 5.5a4.432 4.432 0 0 1 .616-2.27c.197-.336.432-.64.705-.914a4.6 4.6 0 0 1 .911-.702c.338-.196.7-.348 1.084-.454a4.45 4.45 0 0 1 1.2-.16a4.476 4.476 0 0 1 2.276.614a4.475 4.475 0 0 1 1.622 1.616a4.438 4.438 0 0 1 .616 2.27c0 .617-.117 1.191-.353 1.721a4.69 4.69 0 0 1-1.006 1.437zM9.623 10.5H7.409v2.201c0 .081.028.15.09.212a.29.29 0 0 0 .213.09h1.606a.289.289 0 0 0 .213-.09a.286.286 0 0 0 .09-.212V10.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbAutofix(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6m1.31 5L12 13l-1.3 1l.5-1.53l-1.2-.83h1.47L12 10l.54 1.64H14l-1.2.83z"/><path fill-rule="evenodd" d="M11.17 8.085A3.979 3.979 0 0 0 8.288 10.5H6.409v2.201c0 .081.028.15.09.212a.29.29 0 0 0 .213.09h1.413c.089.348.223.678.396.982c-.066.01-.134.015-.203.015H6.712a1.285 1.285 0 0 1-.922-.379a1.303 1.303 0 0 1-.38-.92v-1.6c0-.479-.092-.921-.274-1.329a3.556 3.556 0 0 0-.776-1.114a4.689 4.689 0 0 1-1.006-1.437A4.187 4.187 0 0 1 3 5.5a4.432 4.432 0 0 1 .616-2.27c.197-.336.432-.64.705-.914a4.6 4.6 0 0 1 .911-.702c.338-.196.7-.348 1.084-.454a4.45 4.45 0 0 1 1.2-.16a4.476 4.476 0 0 1 2.276.614a4.475 4.475 0 0 1 1.622 1.616a4.438 4.438 0 0 1 .616 2.27c0 .617-.117 1.191-.353 1.721a4.537 4.537 0 0 1-.506.864z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.4 3h3.085a3.4 3.4 0 0 1 3.4 3.4v.205A3.4 3.4 0 0 1 7.485 10H7V9h.485A2.4 2.4 0 0 0 9.88 6.61V6.4A2.4 2.4 0 0 0 7.49 4H4.4A2.4 2.4 0 0 0 2 6.4v.205A2.394 2.394 0 0 0 4 8.96v1a3.4 3.4 0 0 1-3-3.35V6.4A3.405 3.405 0 0 1 4.4 3M12 7.04v-1a3.4 3.4 0 0 1 3 3.36v.205A3.405 3.405 0 0 1 11.605 13h-3.09A3.4 3.4 0 0 1 5.12 9.61V9.4A3.4 3.4 0 0 1 8.515 6H9v1h-.485A2.4 2.4 0 0 0 6.12 9.4v.205A2.4 2.4 0 0 0 8.515 12h3.09A2.4 2.4 0 0 0 14 9.61V9.4a2.394 2.394 0 0 0-2-2.36" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkExternal(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M1.5 1H6v1H2v12h12v-4h1v4.5l-.5.5h-13l-.5-.5v-13z"/><path d="M15 1.5V8h-1V2.707L7.243 9.465l-.707-.708L13.293 2H8V1h6.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListFilter(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12v-1h4v1zM4 7h8v1H4zm10-4v1H2V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListFlat(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 9h9v1H2zm0 3h8v1H2zm0-6h12v1H2zm0-3h11v1H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOrdered(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.287 2.326L2.692 2h.677v3h-.708V2.792l-.374.281zM5 3h10v1H5zm0 4h10v1H5zm10 4H5v1h10zM3.742 7.626l.029-.039l.065-.09a.84.84 0 0 0 .156-.507c0-.12-.02-.24-.057-.354a.848.848 0 0 0-.492-.529a1.123 1.123 0 0 0-.452-.082a1.094 1.094 0 0 0-.458.087a.867.867 0 0 0-.479.522A1.038 1.038 0 0 0 2 6.957v.05h.81v-.05a.3.3 0 0 1 .046-.157a.174.174 0 0 1 .057-.054a.19.19 0 0 1 .172 0a.188.188 0 0 1 .056.06a.24.24 0 0 1 .031.081a.445.445 0 0 1-.036.29a1.309 1.309 0 0 1-.12.182l-1 1.138l-.012.013v.54h1.988v-.7h-.9zm-.037 3.817c.046.032.086.07.12.114a.841.841 0 0 1 .167.55c0 .107-.017.213-.05.314a.792.792 0 0 1-.487.5a1.288 1.288 0 0 1-.48.079c-.115 0-.23-.016-.341-.049a.94.94 0 0 1-.258-.123a.751.751 0 0 1-.182-.177a1.063 1.063 0 0 1-.116-.2A1.038 1.038 0 0 1 2 12.078v-.049h.814v.049c0 .027.003.055.009.082a.207.207 0 0 0 .03.074a.14.14 0 0 0 .053.052a.2.2 0 0 0 .157.008a.159.159 0 0 0 .056-.039a.22.22 0 0 0 .042-.075a.417.417 0 0 0 .017-.126a.483.483 0 0 0-.022-.163a.2.2 0 0 0-.051-.08a.138.138 0 0 0-.06-.029a.537.537 0 0 0-.077-.007h-.161v-.645h.168a.241.241 0 0 0 .069-.011a.164.164 0 0 0 .065-.034a.175.175 0 0 0 .048-.067a.286.286 0 0 0 .021-.121a.28.28 0 0 0-.016-.1a.166.166 0 0 0-.097-.099a.2.2 0 0 0-.156.007a.164.164 0 0 0-.055.053a.344.344 0 0 0-.04.156v.049H2v-.049a.987.987 0 0 1 .18-.544a.8.8 0 0 1 .179-.186a.87.87 0 0 1 .262-.133c.114-.036.234-.053.354-.051c.116-.001.231.01.344.036c.092.021.18.055.263.1a.757.757 0 0 1 .32.318a.73.73 0 0 1 .09.347a.81.81 0 0 1-.167.528a.562.562 0 0 1-.12.114" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListSelection(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 12v-1h9v1zm0-5h14v1H1zm11-4v1H1V3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListTree(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4 9h9v1H4zm0 3h7v1H4zm0-6h10v1H4zM1 3h11v1H1z"/><path d="M4 4h1v9H4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUnordered(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 3H1v1h1zm0 3H1v1h1zM1 9h1v1H1zm1 3H1v1h1zm2-9h11v1H4zm11 3H4v1h11zM4 9h11v1H4zm11 3H4v1h11z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LiveShare(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.735 1.694L15.178 1l8.029 6.328v1.388l-8.029 6.072l-1.443-.694v-2.776h-.59c-4.06-.02-6.71.104-10.61 5.163l-1.534-.493a8.23 8.23 0 0 1 .271-2.255a11.026 11.026 0 0 1 3.92-6.793a11.339 11.339 0 0 1 7.502-2.547h1.04zm1.804 7.917v2.776l5.676-4.281l-5.648-4.545v2.664h-2.86A9.299 9.299 0 0 0 5.77 8.848a10.444 10.444 0 0 0-2.401 4.122c3.351-3.213 6.19-3.359 9.798-3.359zm-7.647 5.896a4.31 4.31 0 1 1 4.788 7.166a4.31 4.31 0 0 1-4.788-7.166m.955 5.728a2.588 2.588 0 1 0 2.878-4.302a2.588 2.588 0 0 0-2.878 4.302" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loading(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.917 7A6.002 6.002 0 0 0 2.083 7H1.071a7.002 7.002 0 0 1 13.858 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.832 2.688A4.056 4.056 0 0 0 8.02 1.5h-.04a4.056 4.056 0 0 0-4 4c-.013.75.198 1.487.606 2.117L7.734 14h.533l3.147-6.383c.409-.63.62-1.367.606-2.117a4.056 4.056 0 0 0-1.188-2.812M7.925 2.5l.082.01l.074-.01a3.075 3.075 0 0 1 2.941 3.037a2.74 2.74 0 0 1-.467 1.568l-.02.034l-.017.035L8 12.279l-2.517-5.1l-.017-.039l-.02-.034a2.74 2.74 0 0 1-.467-1.568A3.074 3.074 0 0 1 7.924 2.5zm.612 2.169a1 1 0 1 0-1.112 1.663a1 1 0 0 0 1.112-1.663M6.87 3.837a2 2 0 1 1 2.22 3.326a2 2 0 0 1-2.22-3.326" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 7h-1V5a4 4 0 1 0-8 0v2H3L2 8v6l1 1h10l1-1V8zM5 5a3 3 0 1 1 6 0v2H5zm8 9H3V8h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSmall(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3 8l1-1h8l1 1v5l-1 1H4l-1-1zm1 0v5h8V8zm7-1V5a3 3 0 0 0-6 0v2h1V5a2 2 0 1 1 4 0v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1.5c-3.9 0-7 3.1-7 7v5l1 1h3l1-1v-5c0-1.1.9-2 2-2s2 .9 2 2v5l1 1h3l1-1v-5c0-3.9-3.1-7-7-7m-3 12H2v-3h3zm9 0h-3v-3h3zm-3-4v-1c0-1.7-1.3-3-3-3c-1.6 0-2.9 1.3-3 2.8v1.2H2v-1c0-3.3 2.7-6 6-6s6 2.7 6 6v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1 3.5l.5-.5h13l.5.5v9l-.5.5h-13l-.5-.5zm1 1.035V12h12V4.536L8.31 8.9H7.7zM13.03 4H2.97L8 7.869z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailRead(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.25 1.57h-.51L1 5.56v7.94l.5.5h13l.5-.5V5.56zM8 2.58l5.63 3.32l-1.37 1.59H3.74L2.43 5.9zM14 13H2V6.92L3.11 8.3l.39.19h9l.39-.19L14 6.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 5.777v6.32l3-1.874V3.902zm4-1.875v6.32l3 1.876V5.777zM6 11.09l-3.735 2.334L1.5 13V5.5l.235-.424l4-2.5h.53L10 4.91l3.735-2.334L14.5 3v7.5l-.235.424l-4 2.5h-.53zm4.5-5.313v6.32l3-1.874V3.902z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5.5V13l3.5-2.187v-7.5zm7.5 7.188v-7.5l-3-1.875v7.5zm1 0v-7.5L14 3v7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Markdown(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.345 5h2.1v6.533H6.993l.055-5.31l-1.774 5.31H4.072l-1.805-5.31c.04.644.06 5.31.06 5.31H1V5h2.156s1.528 4.493 1.577 4.807zm6.71 3.617v-3.5H11.11v3.5H9.166l2.917 2.916L15 8.617z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2 6.77l12.33-3.43l.67.53v8.6l-.67.53l-6.089-1.595a2.16 2.16 0 1 1-4.178-1.095L2 9.77l-.42-.53V7.3zm3.006 3.787a1.13 1.13 0 0 0-.04.242a1.17 1.17 0 0 0 2.288.347zM2.58 8.82L14 11.83V4.5L2.58 7.72z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mention(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.465 12.36a5.516 5.516 0 0 1-3.053.736a4.262 4.262 0 0 1-4.57-4.543a5.381 5.381 0 0 1 5.391-5.571c2.377 0 4.413 1.375 4.413 4.006c0 2.182-1.292 3.66-2.9 3.66c-.676 0-1.1-.274-1.126-.917a2.012 2.012 0 0 1-1.756.913c-.969 0-1.629-.645-1.629-1.923c0-1.763 1.148-3.4 2.62-3.4a1.314 1.314 0 0 1 1.427.93l.211-.809h.9L9.6 8.646c-.226.916-.13 1.215.342 1.215c.984 0 1.833-1.21 1.833-2.825c0-2.068-1.445-3.265-3.61-3.265c-2.643 0-4.374 2.132-4.382 4.786a3.443 3.443 0 0 0 3.686 3.717c.973.04 1.94-.179 2.8-.634zM6.217 8.639c0 .788.307 1.206.913 1.206c.758 0 1.38-.6 1.683-1.831C9.136 6.746 8.85 6.1 7.94 6.1c-1.04 0-1.723 1.339-1.723 2.539"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5H0V4h16zm0 8H0v-1h16zm0-4.008H0V8h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Merge(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 4.646L8.354 2.5h-.707L5.5 4.646l.707.707L7.3 4.261V5.28h-.02v.456l.025.001l.006.319c.004.187.02.379.05.574c.03.195.069.39.117.586c.048.195.114.404.2.627c.155.379.343.722.565 1.031c.221.309.46.598.715.867c.255.27.508.535.76.797c.25.262.478.541.681.838c.203.297.368.621.494.973c.125.351.188.755.188 1.213v.884H12.5v-.884a5.991 5.991 0 0 0-.166-1.39a4.638 4.638 0 0 0-.427-1.1a5.875 5.875 0 0 0-.604-.897c-.222-.27-.453-.527-.693-.774c-.24-.246-.471-.492-.693-.738a6.39 6.39 0 0 1-.604-.785a3.794 3.794 0 0 1-.433-.914a3.676 3.676 0 0 1-.16-1.13V5.28h-.001v-1l1.074 1.074zM7.042 9.741a8.19 8.19 0 0 0 .329-.369a6.06 6.06 0 0 1-.62-1.15L6.744 8.2a7.26 7.26 0 0 1-.095-.263c-.17.256-.359.498-.565.726c-.222.246-.453.492-.693.738c-.24.247-.47.504-.693.774c-.221.27-.423.568-.604.896a4.643 4.643 0 0 0-.427 1.102a5.995 5.995 0 0 0-.166 1.389v.884h1.42v-.884c0-.457.062-.862.188-1.213c.125-.352.29-.676.493-.973c.203-.297.43-.576.682-.838c.251-.262.504-.527.76-.797z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 1a3 3 0 0 0-3 3v4a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3m2 7a2 2 0 1 1-4 0V4a2 2 0 1 1 4 0z"/><path d="M11.696 9.53A4 4 0 0 0 12 8h1a5.001 5.001 0 0 1-4.5 4.975V15h-1v-2.025A4.997 4.997 0 0 1 3 8h1a4.002 4.002 0 0 0 4 4a4 4 0 0 0 3.695-2.47"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 1a3 3 0 0 0-3 3v4a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3"/><path d="M11.696 9.53A4 4 0 0 0 12 8h1a5.001 5.001 0 0 1-4.5 4.975V15h-1v-2.025A4.997 4.997 0 0 1 3 8h1a4.002 4.002 0 0 0 4 4a4 4 0 0 0 3.695-2.47"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Milestone(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1H7v2H1.5l-.5.5v4l.5.5H7v7h1V8h4.49l.34-.13l2.18-2v-.74l-2.18-2L12.5 3H8zm4.29 6H2V4h10.29l1.63 1.5zM5 5h5v1H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mirror(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.57 1l6.2 4l.23.38v9.2l-.76.42L8 11l-6.24 4l-.76-.42v-9.2L1.23 5l6.2-4zm-.06 9.13L14 13.67V5.65l-5.49-3.5V5h-1V2.13L2 5.67v8l5.51-3.56v.02zm.9-4.78l.71-.7l2.47 2.48v.71l-2.46 2.46l-.7-.7L11.02 8h-6L6.6 9.6l-.7.7l-2.46-2.46v-.71l2.48-2.48l.7.7L4.98 7h6.08z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MortarBoard(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 5.66L8.18 3h-.36L1 5.66V12h1V7l2.31.9a4.35 4.35 0 0 0-.79 2.48c-.01.11-.01.22 0 .33v.11l.28.4L7.78 13h.41l3.94-1.81l.28-.4v-.44a4.39 4.39 0 0 0-.78-2.47L15 6.57zm-3.52 4.68v.07L8 12l-3.5-1.6a.13.13 0 0 1 0-.06a3.44 3.44 0 0 1 .75-2.12l2.58 1h.36l2.56-1a3.4 3.4 0 0 1 .73 2.12M8 8.25L2.52 6.12L8 4l5.48 2.14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.232 10.707L8.5 12.44V9h-1v3.44l-1.732-1.733l-.707.707L7.646 14h.708l2.585-2.586zM5.061 3.586l.707.707L7.5 2.56V6h1V2.56l1.732 1.733l.707-.707L8.354 1h-.708zm-.268 1.682L3.06 7H6.5v1H3.06l1.733 1.732l-.707.707L1.5 7.854v-.708l2.586-2.585zM9.5 7h3.44l-1.733-1.732l.707-.707L14.5 7.146v.708l-2.586 2.585l-.707-.707L12.94 8H9.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultipleWindows(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6 1.5l.5-.5h8l.5.5v7l-.5.5H12V8h2V4H7v1H6zM7 2v1h7V2zM1.5 7l-.5.5v7l.5.5h8l.5-.5v-7L9.5 7zM2 9V8h7v1zm0 1h7v4H2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M14 7h-1v2.5a2.5 2.5 0 1 0 1 2zm-2.5 3a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3"/><path d="m13.469 2.001l-8 .5L5 3v7.5a2.5 2.5 0 1 0 1 2V6.47l7-.438V7h1V2.5zM13 3.032V5.03l-7 .438V3.47zM3.5 11a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mute(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 5h2.79l3.86-3.83l.85.35v13l-.85.33L4.29 11H1.5l-.5-.5v-5zm3.35 5.17L8 13.31V2.73L4.85 5.85L4.5 6H2v4h2.5zm9.381-4.108l.707.707L13.207 8.5l1.731 1.732l-.707.707L12.5 9.207l-1.732 1.732l-.707-.707L11.793 8.5L10.06 6.77l.707-.707l1.733 1.73z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewFile(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.5 1.1l3.4 3.5l.1.4v2h-1V6H8V2H3v11h4v1H2.5l-.5-.5v-12l.5-.5h6.7zM9 2v3h2.9zm4 14h-1v-3H9v-1h3V9h1v3h3v1h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NewFolder(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 2H7.71l-.85-.85L6.51 1h-5l-.5.5v11l.5.5H7v-1H1.99V6h4.49l.35-.15l.86-.86H14v1.5l-.001.51h1.011V2.5zm-.51 2h-6.5l-.35.15l-.86.86H2v-3h4.29l.85.85l.36.15H14zM13 16h-1v-3H9v-1h3V9h1v3h3v1h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newline(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 5.5v1.984a.5.5 0 0 1-.5.5H4.618l1.633-1.633l-.707-.707l-2.121 2.121L3 8.188v.568L5.544 11.3l.707-.707l-1.61-1.609H11.5a1.5 1.5 0 0 0 1.5-1.5V5.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoNewline(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.333 5.506a3 3 0 1 1 3.334 4.989a3 3 0 0 1-3.334-4.99zm2.677.777A1.986 1.986 0 0 0 2 8.009c.004.353.102.698.283 1.001zM2.99 9.717A1.986 1.986 0 0 0 6 7.991a1.988 1.988 0 0 0-.283-1.001zM14 5v1.984a.5.5 0 0 1-.5.5H9.367L11 5.851l-.707-.707l-2.121 2.121l-.423.423v.568l2.544 2.544l.707-.707l-1.61-1.609h4.11a1.5 1.5 0 0 0 1.5-1.5V5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Note(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2h13l.5.5v10l-.5.5h-13l-.5-.5v-10zM2 3v9h12V3zm2 2h8v1H4zm6 2H4v1h6zM4 9h4v1H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notebook(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2 2l1-1h9l1 1v12l-1 1H3l-1-1zm1 0v12h9V2zm1 2l1-1h5l1 1v1l-1 1H5L4 5zm1 0v1h5V4zm10 1h-1v2h1zm-1 3h1v2h-1zm1 3h-1v2h1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotebookTemplate(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M1 5H0V4h1zm0 2H0V6h1zm0 2H0V8h1zm0 2H0v-1h1zm0 2H0v-1h1zm0 1v1H0v-1zm0 1h1v1H1zm2 0h1v1H3zM1 1H0V0h1zm2 0H2V0h1zm1-1h1v1H4zm3 1H6V0h1zm2 0H8V0h1zm2 0h-1V0h1zm0 1V1h1v1zm1 2h-1V3h1zM1 3H0V2h1z"/><path fill-rule="evenodd" d="m5 6l1-1h7l1 1v9l-1 1H6l-1-1zm1 0v9h7V6z" clip-rule="evenodd"/><path d="M15 7h1v2h-1zm0 3h1v2h-1zm0 3h1v2h-1zM7 8h5v1H7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Octoface(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.863 5.673c.113-.28.48-1.392-.114-2.897c0 0-.919-.288-3.01 1.138c-.875-.245-1.812-.28-2.739-.28c-.928 0-1.864.035-2.739.28c-2.091-1.435-3.01-1.138-3.01-1.138c-.595 1.505-.227 2.617-.113 2.897C1.428 6.433 1 7.413 1 8.603c0 4.507 2.914 5.522 6.982 5.522c4.07 0 7.018-1.015 7.018-5.521c0-1.19-.429-2.17-1.137-2.931M8 13.268c-2.888 0-5.232-.132-5.232-2.932c0-.665.332-1.295.892-1.811c.936-.857 2.537-.402 4.34-.402c1.811 0 3.395-.455 4.34.402c.569.516.893 1.138.893 1.811c0 2.791-2.346 2.931-5.233 2.931zM5.804 8.883c-.578 0-1.05.7-1.05 1.557c0 .858.472 1.566 1.05 1.566c.577 0 1.05-.7 1.05-1.566c0-.866-.473-1.557-1.05-1.557m4.392 0c-.577 0-1.05.691-1.05 1.557s.473 1.566 1.05 1.566c.578 0 1.05-.7 1.05-1.566c0-.866-.463-1.557-1.05-1.557" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenPreview(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 1h11l1 1v5.3a3.21 3.21 0 0 0-1-.3V2H9v10.88L7.88 14H3l-1-1V2zm0 12h5V2H3zm10.379-4.998a2.53 2.53 0 0 0-1.19.348h-.03a2.51 2.51 0 0 0-.799 3.53L9 14.23l.71.71l2.35-2.36c.325.22.7.358 1.09.4a2.47 2.47 0 0 0 1.14-.13a2.51 2.51 0 0 0 1-.63a2.46 2.46 0 0 0 .58-1a2.63 2.63 0 0 0 .07-1.15a2.53 2.53 0 0 0-1.35-1.81a2.53 2.53 0 0 0-1.211-.258m.24 3.992a1.5 1.5 0 0 1-.979-.244a1.55 1.55 0 0 1-.56-.68a1.49 1.49 0 0 1-.08-.86a1.49 1.49 0 0 1 1.18-1.18a1.49 1.49 0 0 1 .86.08c.276.117.512.311.68.56a1.5 1.5 0 0 1-1.1 2.324z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Organization(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.111 4.663A2 2 0 1 1 6.89 1.337a2 2 0 0 1 2.222 3.326zm-.555-2.494A1 1 0 1 0 7.444 3.83a1 1 0 0 0 1.112-1.66zm2.61.03a1.494 1.494 0 0 1 1.895.188a1.513 1.513 0 0 1-.487 2.46a1.492 1.492 0 0 1-1.635-.326a1.512 1.512 0 0 1 .228-2.321zm.48 1.61a.499.499 0 1 0 .705-.708a.509.509 0 0 0-.351-.15a.499.499 0 0 0-.5.503a.51.51 0 0 0 .146.356zM3.19 12.487H5v1.005H3.19a1.197 1.197 0 0 1-.842-.357a1.21 1.21 0 0 1-.348-.85v-1.81a.997.997 0 0 1-.71-.332A1.007 1.007 0 0 1 1 9.408V7.226c.003-.472.19-.923.52-1.258c.329-.331.774-.52 1.24-.523H4.6a2.912 2.912 0 0 0-.55 1.006H2.76a.798.798 0 0 0-.54.232a.777.777 0 0 0-.22.543v2.232h1v2.826a.202.202 0 0 0 .05.151a.24.24 0 0 0 .14.05zm7.3-6.518a1.765 1.765 0 0 0-1.25-.523H6.76a1.765 1.765 0 0 0-1.24.523c-.33.335-.517.786-.52 1.258v3.178a1.06 1.06 0 0 0 .29.734a1 1 0 0 0 .71.332v2.323a1.202 1.202 0 0 0 .35.855c.18.168.407.277.65.312h2a1.15 1.15 0 0 0 1-1.167V11.47a.997.997 0 0 0 .71-.332a1.006 1.006 0 0 0 .29-.734V7.226a1.8 1.8 0 0 0-.51-1.258zM10 10.454H9v3.34a.202.202 0 0 1-.06.14a.17.17 0 0 1-.14.06H7.19a.21.21 0 0 1-.2-.2v-3.34H6V7.226c0-.203.079-.398.22-.543a.798.798 0 0 1 .54-.232h2.48a.778.778 0 0 1 .705.48a.748.748 0 0 1 .055.295zm2.81 3.037H11v-1.005h1.8a.24.24 0 0 0 .14-.05a.2.2 0 0 0 .06-.152V9.458h1V7.226a.777.777 0 0 0-.22-.543a.798.798 0 0 0-.54-.232h-1.29a2.91 2.91 0 0 0-.55-1.006h1.84a1.77 1.77 0 0 1 1.24.523c.33.335.517.786.52 1.258v2.182c0 .273-.103.535-.289.733c-.186.199-.44.318-.711.333v1.81c0 .319-.125.624-.348.85a1.197 1.197 0 0 1-.842.357M4 1.945a1.494 1.494 0 0 0-1.386.932A1.517 1.517 0 0 0 2.94 4.52A1.497 1.497 0 0 0 5.5 3.454c0-.4-.158-.784-.44-1.067A1.496 1.496 0 0 0 4 1.945m0 2.012a.499.499 0 0 1-.5-.503a.504.504 0 0 1 .5-.503a.509.509 0 0 1 .5.503a.504.504 0 0 1-.5.503" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Output(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.5 0v1.5L21 3v19.5L19.5 24h-15L3 22.5V3l1.5-1.5V0H6v1.5h3V0h1.5v1.5h3V0H15v1.5h3V0zm-15 22.5h15V3h-15zM7.5 6h9v1.5h-9zm9 6h-9v1.5h9zm-9 6h9v1.5h-9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.61 3l5.74 1.53L15 5v6.74l-.37.48l-6.13 1.69l-6.14-1.69l-.36-.48V5l.61-.47L8.34 3zm-.09 1l-4 1l.55.2l3.43.9l3-.81l.95-.29zM3 11.36l5 1.37V7L3 5.66zM9 7v5.73l5-1.37V5.63l-2.02.553V8.75l-1 .26V6.457z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paintcan(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.54 11.811l-1.14-3.12v-.06l-4.91-4.91v-1.24a1.66 1.66 0 0 0-.11-.58a1.48 1.48 0 0 0-.83-.8a1.42 1.42 0 0 0-.58-.1a1.47 1.47 0 0 0-1.48 1.48v3.26l-3.06 3a1.52 1.52 0 0 0 0 2.12l3.63 3.63c.14.141.307.253.49.33a1.53 1.53 0 0 0 1.14 0a1.51 1.51 0 0 0 .49-.33l4.93-4.92l-.66 2.2a1.19 1.19 0 0 0 0 .46c.033.152.098.296.19.42c.098.121.216.223.35.3c.14.07.294.11.45.12a1 1 0 0 0 .48-.09a1.14 1.14 0 0 0 .39-.29a.98.98 0 0 0 .22-.44c.032-.145.035-.294.01-.44m-8-9.33a.46.46 0 0 1 0-.2a.52.52 0 0 1 .12-.17a.64.64 0 0 1 .18-.1a.5.5 0 0 1 .21 0a.5.5 0 0 1 .32.15a.5.5 0 0 1 .12.33v1.26l-1 1zm1 11.35a.36.36 0 0 1-.16.11a.47.47 0 0 1-.38 0a.361.361 0 0 1-.16-.11l-3.63-3.62a.5.5 0 0 1 0-.71l4.35-4.35v2.85a.74.74 0 0 0-.24.55a.75.75 0 1 0 1.17-.55v-2.83l3.85 3.87z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pass(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6.27 10.87h.71l4.56-4.56l-.71-.71l-4.2 4.21l-1.92-1.92L4 8.6z"/><path fill-rule="evenodd" d="M8.6 1c1.6.1 3.1.9 4.2 2c1.3 1.4 2 3.1 2 5.1c0 1.6-.6 3.1-1.6 4.4c-1 1.2-2.4 2.1-4 2.4c-1.6.3-3.2.1-4.6-.7c-1.4-.8-2.5-2-3.1-3.5C.9 9.2.8 7.5 1.3 6c.5-1.6 1.4-2.9 2.8-3.8C5.4 1.3 7 .9 8.6 1m.5 12.9c1.3-.3 2.5-1 3.4-2.1c.8-1.1 1.3-2.4 1.2-3.8c0-1.6-.6-3.2-1.7-4.3c-1-1-2.2-1.6-3.6-1.7c-1.3-.1-2.7.2-3.8 1c-1.1.8-1.9 1.9-2.3 3.3c-.4 1.3-.4 2.7.2 4c.6 1.3 1.5 2.3 2.7 3c1.2.7 2.6.9 3.9.6" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PassFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14m-1.02-4.13h-.71L4 8.6l.71-.71l1.92 1.92l4.2-4.21l.71.71z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Person(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2a1 1 0 1 1 0 2a1 1 0 0 1 0-2m0-1a2 2 0 1 0 0 4a2 2 0 0 0 0-4m1.23 4.49H6.77A1.77 1.77 0 0 0 5 7.26V9.9A1.06 1.06 0 0 0 6 11v2.33a1.2 1.2 0 0 0 1.2 1.2h1.6a1.2 1.2 0 0 0 1.2-1.24V11a1.06 1.06 0 0 0 1-1.1V7.26a1.77 1.77 0 0 0-1.77-1.77M6 10V7.26a.76.76 0 0 1 .77-.77h2.46a.76.76 0 0 1 .77.77V10H9v3.31a.2.2 0 0 1-.2.2H7.2a.2.2 0 0 1-.2-.2V10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonAdd(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 10h-1v2h-2v1h2v2h1v-2h2v-1h-2zM8.556 2.169a1 1 0 1 0-1.112 1.663a1 1 0 0 0 1.112-1.663m-1.667-.832A2 2 0 1 1 9.11 4.663a2 2 0 0 1-2.22-3.326zM6.77 5.49h2.46A1.77 1.77 0 0 1 11 7.26V8h-1v-.74a.76.76 0 0 0-.77-.77H6.77a.76.76 0 0 0-.77.77V10h1v3.31a.2.2 0 0 0 .2.2H8v1.02h-.8a1.2 1.2 0 0 1-1.2-1.2V11a1.06 1.06 0 0 1-1-1.1V7.26a1.77 1.77 0 0 1 1.77-1.77" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Piano(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 2L0 3v6h1V3h2v5.5l.5.5h1l.5-.5V3h2v5.5l.5.5h1l.5-.5V3h2v5.5l.5.5h1l.5-.5V3h2v10h-3v-3h-1v3H8.5v-3h-1v3H5v-3H4v3H1V9H0v4l1 1h14l1-1V3l-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChart(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 6h3.9A5.006 5.006 0 0 0 10 2.1zm0-4.917A6.005 6.005 0 0 1 15 7H9V1c.34 0 .675.028 1 .083M7 8l1 1h4.9A5.002 5.002 0 0 1 3 8a5.002 5.002 0 0 1 4-4.9zm1 6a6.002 6.002 0 0 0 6-6H8V2a6.002 6.002 0 0 0 0 12" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 5v7h-.278c-.406 0-.778-.086-1.117-.258A2.528 2.528 0 0 1 11.73 11H8.87a3.463 3.463 0 0 1-.546.828a3.685 3.685 0 0 1-.735.633c-.27.177-.565.31-.882.398a3.875 3.875 0 0 1-.985.141h-.5V9H2l-1-.5L2 8h3.222V4h.5c.339 0 .664.047.977.14c.312.094.607.227.883.4A3.404 3.404 0 0 1 8.87 6h2.859a2.56 2.56 0 0 1 .875-.734c.338-.172.71-.26 1.117-.266zm-.778 1.086a1.222 1.222 0 0 0-.32.156a1.491 1.491 0 0 0-.43.461L12.285 7H8.183l-.117-.336a2.457 2.457 0 0 0-.711-1.047C7.027 5.331 6.427 5.09 6 5v7c.427-.088 1.027-.33 1.355-.617c.328-.287.565-.636.71-1.047L8.184 10h4.102l.18.297c.057.094.122.177.195.25c.073.073.153.143.242.21c.088.069.195.12.32.157z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinned(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2h7v.278c0 .406-.086.778-.258 1.117c-.172.339-.42.63-.742.875v2.86c.307.145.583.328.828.546c.245.219.456.464.633.735c.177.27.31.565.398.882c.089.318.136.646.141.985v.5H8V14l-.5 1l-.5-1v-3.222H3v-.5c0-.339.047-.664.14-.977c.094-.312.227-.607.4-.883A3.404 3.404 0 0 1 5 7.13V4.27a2.561 2.561 0 0 1-.734-.875A2.505 2.505 0 0 1 4 2.278zm1.086.778c.042.125.094.232.156.32a1.494 1.494 0 0 0 .461.43L6 3.715v4.102l-.336.117c-.411.146-.76.383-1.047.711C4.331 8.973 4.09 9.573 4 10h7c-.088-.427-.33-1.027-.617-1.355a2.456 2.456 0 0 0-1.047-.71L9 7.816V3.715l.297-.18c.094-.057.177-.122.25-.195a2.28 2.28 0 0 0 .21-.242a.968.968 0 0 0 .157-.32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinnedDirty(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2h7v.278c0 .406-.086.778-.258 1.117c-.172.339-.42.63-.742.875v2.86c.307.145.583.328.828.546a3.7 3.7 0 0 1 .54.598a4.92 4.92 0 0 0-.896.412l-.007.004l-.03.018a2.456 2.456 0 0 0-1.099-.774L9 7.817V3.715l.297-.18c.094-.057.177-.122.25-.195a2.28 2.28 0 0 0 .21-.242a.968.968 0 0 0 .157-.32H5.086c.042.125.094.232.156.32a1.494 1.494 0 0 0 .461.43L6 3.715v4.102l-.336.117c-.411.146-.76.383-1.047.711C4.331 8.973 4.09 9.573 4 10h5.002a5.025 5.025 0 0 0-.481.778H8V14l-.5 1l-.5-1v-3.222H3v-.5c0-.339.047-.664.14-.977c.094-.312.227-.607.4-.883A3.404 3.404 0 0 1 5 7.13V4.27a2.561 2.561 0 0 1-.734-.875A2.505 2.505 0 0 1 4 2.278zm7.485 8.41a2.924 2.924 0 0 1 .718-.302c.256-.072.522-.108.797-.108s.541.036.797.108a2.956 2.956 0 0 1 1.321.773a2.956 2.956 0 0 1 .774 1.322c.072.256.108.522.108.797s-.036.541-.108.797a2.953 2.953 0 0 1-.774 1.324a3.013 3.013 0 0 1-1.321.774c-.256.07-.522.105-.797.105s-.541-.035-.797-.105a3.037 3.037 0 0 1-1.324-.774a3.037 3.037 0 0 1-.773-1.324A2.994 2.994 0 0 1 10 13c0-.275.035-.541.105-.797a3.013 3.013 0 0 1 .883-1.425c.154-.14.32-.262.497-.368" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinnedDirtySmall(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 3h5v.278c0 .406-.086.778-.258 1.117c-.172.339-.42.63-.742.875v1.86c.307.145.583.328.828.546a3.686 3.686 0 0 1 .759.945a4.829 4.829 0 0 0-.115.065l-.007.004a5.05 5.05 0 0 0-.73.526a2.471 2.471 0 0 0-.352-.571a2.457 2.457 0 0 0-1.047-.71L8 7.816V4.715l.297-.18c.094-.057.177-.122.25-.195c.073-.073.143-.153.21-.242a.968.968 0 0 0 .157-.32H6.086c.042.125.094.232.156.32a1.494 1.494 0 0 0 .461.43L7 4.715v3.102l-.336.117c-.411.146-.76.383-1.047.711C5.331 8.973 5.09 9.573 5 10h4.002a5.025 5.025 0 0 0-.481.778H8V13l-.5 1l-.5-1v-2.222H4v-.5c0-.339.047-.664.14-.977c.094-.312.227-.607.4-.883A3.404 3.404 0 0 1 6 7.13V5.27a2.561 2.561 0 0 1-.734-.875A2.505 2.505 0 0 1 5 3.278V3zm8.797 7.108A2.917 2.917 0 0 0 13 10c-.275 0-.541.036-.797.108a2.953 2.953 0 0 0-1.324.774a3.013 3.013 0 0 0-.774 1.321c-.07.256-.105.522-.105.797s.035.541.105.797a3.037 3.037 0 0 0 .774 1.324a3.037 3.037 0 0 0 1.324.773c.256.07.522.106.797.106s.541-.035.797-.105a3.013 3.013 0 0 0 1.321-.773a2.953 2.953 0 0 0 .773-1.324c.073-.257.109-.523.109-.798s-.036-.541-.108-.797a2.956 2.956 0 0 0-.773-1.321a2.956 2.956 0 0 0-1.322-.774z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinnedSmall(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 3h5v.278c0 .406-.086.778-.258 1.117c-.172.339-.42.63-.742.875v1.86c.307.145.583.328.828.546c.245.219.456.464.633.735c.177.27.31.565.398.882c.089.318.136.646.141.985v.5H8V13l-.5 1l-.5-1v-2.222H4v-.5c0-.339.047-.664.14-.977c.094-.312.227-.607.4-.883A3.404 3.404 0 0 1 6 7.13V5.27a2.561 2.561 0 0 1-.734-.875A2.505 2.505 0 0 1 5 3.278V3zm1.086.778c.042.125.094.232.156.32a1.494 1.494 0 0 0 .461.43L7 4.715v3.102l-.336.117c-.411.146-.76.383-1.047.711C5.331 8.973 5.09 9.573 5 10h5c-.089-.427-.33-1.027-.617-1.355a2.457 2.457 0 0 0-1.047-.71L8 7.816V4.715l.297-.18c.094-.057.177-.122.25-.195c.073-.073.143-.153.21-.242a.968.968 0 0 0 .157-.32H6.086z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.78 2L3 2.41v12l.78.42l9-6V8zM4 13.48V3.35l7.6 5.07z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.6 1c1.6.1 3.1.9 4.2 2c1.3 1.4 2 3.1 2 5.1c0 1.6-.6 3.1-1.6 4.4c-1 1.2-2.4 2.1-4 2.4c-1.6.3-3.2.1-4.6-.7c-1.4-.8-2.5-2-3.1-3.5C.9 9.2.8 7.5 1.3 6c.5-1.6 1.4-2.9 2.8-3.8C5.4 1.3 7 .9 8.6 1m.5 12.9c1.3-.3 2.5-1 3.4-2.1c.8-1.1 1.3-2.4 1.2-3.8c0-1.6-.6-3.2-1.7-4.3c-1-1-2.2-1.6-3.6-1.7c-1.3-.1-2.7.2-3.8 1c-1.1.8-1.9 1.9-2.3 3.3c-.4 1.3-.4 2.7.2 4c.6 1.3 1.5 2.3 2.7 3c1.2.7 2.6.9 3.9.6"/><path d="m6 5l.777-.416l4.5 3v.832l-4.5 3L6 11zm1 .934v4.132L10.099 8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plug(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 1H6v3H4.5l-.5.5V8a4 4 0 0 0 3.5 3.969V15h1v-3.031A4 4 0 0 0 12 8V4.5l-.5-.5H10V1H9v3H7zm3.121 9.121A3 3 0 0 1 5 8V5h6v3a3 3 0 0 1-.879 2.121" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PreserveCase(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.514 11h-1l-.816-2.16H3.433L2.664 11H1.66l2.954-7.702h.935zM6.403 8.03L5.194 4.748a3.144 3.144 0 0 1-.118-.516h-.021c-.036.219-.077.39-.124.516L3.733 8.03zM9.597 11V3.298h2.192c.666 0 1.194.163 1.584.489c.39.325.586.75.586 1.273c0 .436-.119.816-.355 1.138a1.911 1.911 0 0 1-.977.688v.021c.519.061.934.258 1.246.591c.311.33.467.76.467 1.29c0 .658-.236 1.191-.71 1.6c-.472.408-1.068.612-1.788.612zm.903-6.886v2.487h.923c.495 0 .883-.118 1.166-.354c.283-.24.424-.577.424-1.01c0-.749-.492-1.123-1.477-1.123zm0 3.298v2.772h1.224c.53 0 .94-.126 1.23-.376c.294-.251.44-.595.44-1.032c0-.91-.619-1.364-1.858-1.364z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Preview(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2h12l1 1v10l-1 1H2l-1-1V3zm0 11h12V3H2zm11-9H3v3h10zm-1 2H4V5h8zm-3 6h4V8H9zm1-3h2v2h-2zM7 8H3v1h4zm-4 3h4v1H3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrimitiveSquare(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.5 4l.5-.5h8l.5.5v8l-.5.5H4l-.5-.5zm1 .5v7h7v-7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Project(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 1h13l.5.5v13l-.5.5h-13l-.5-.5v-13zM2 14h12V2H2zM3 3h2v10H3zm6 0H7v6h2zm2 0h2v8h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pulse(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.8 9L10 3H9L7.158 9.64L5.99 4.69h-.97L3.85 9H1v.99h3.23l.49-.37l.74-2.7L6.59 12h1.03l1.87-7.04l1.46 4.68l.48.36H15V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 1a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13m0 12a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11m1.55-8.42a1.84 1.84 0 0 0-.61-.42A2.25 2.25 0 0 0 7.53 4a2.16 2.16 0 0 0-.88.17c-.239.1-.45.254-.62.45a1.89 1.89 0 0 0-.38.62a3 3 0 0 0-.15.72h1.23a.84.84 0 0 1 .506-.741a.72.72 0 0 1 .304-.049a.86.86 0 0 1 .27 0a.64.64 0 0 1 .22.14a.6.6 0 0 1 .16.22a.73.73 0 0 1 .06.3c0 .173-.037.343-.11.5a2.4 2.4 0 0 1-.27.46l-.35.42c-.12.13-.24.27-.35.41a2.33 2.33 0 0 0-.27.45a1.18 1.18 0 0 0-.1.5v.66H8v-.49a.94.94 0 0 1 .11-.42a3.09 3.09 0 0 1 .28-.41l.36-.44a4.29 4.29 0 0 0 .36-.48a2.59 2.59 0 0 0 .28-.55a1.91 1.91 0 0 0 .11-.64a2.18 2.18 0 0 0-.1-.67a1.52 1.52 0 0 0-.35-.55M6.8 9.83h1.17V11H6.8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.16 3.5C4.73 5.06 3.55 6.67 3.55 9.36c.16-.05.3-.05.44-.05c1.27 0 2.5.86 2.5 2.41c0 1.61-1.03 2.61-2.5 2.61c-1.9 0-2.99-1.52-2.99-4.25c0-3.8 1.75-6.53 5.02-8.42zm7 0c-2.43 1.56-3.61 3.17-3.61 5.86c.16-.05.3-.05.44-.05c1.27 0 2.5.86 2.5 2.41c0 1.61-1.03 2.61-2.5 2.61c-1.89 0-2.98-1.52-2.98-4.25c0-3.8 1.75-6.53 5.02-8.42l1.14 1.84z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioTower(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.998 5.58a5.55 5.55 0 0 1 1.62-3.88l-.71-.7a6.45 6.45 0 0 0 0 9.16l.71-.7a5.55 5.55 0 0 1-1.62-3.88m1.06 0a4.42 4.42 0 0 0 1.32 3.17l.71-.71a3.27 3.27 0 0 1-.76-1.12a3.45 3.45 0 0 1 0-2.67a3.22 3.22 0 0 1 .76-1.13l-.71-.71a4.46 4.46 0 0 0-1.32 3.17m7.65 3.21l-.71-.71c.33-.32.59-.704.76-1.13a3.449 3.449 0 0 0 0-2.67a3.22 3.22 0 0 0-.76-1.13l.71-.7a4.468 4.468 0 0 1 0 6.34M13.068 1l-.71.71a5.43 5.43 0 0 1 0 7.74l.71.71a6.45 6.45 0 0 0 0-9.16M9.993 5.43a1.5 1.5 0 0 1-.245.98a2 2 0 0 1-.27.23l3.44 7.73l-.92.4l-.77-1.73h-5.54l-.77 1.73l-.92-.4l3.44-7.73a1.52 1.52 0 0 1-.33-1.63a1.55 1.55 0 0 1 .56-.68a1.5 1.5 0 0 1 2.325 1.1m-1.595-.34a.52.52 0 0 0-.25.14a.52.52 0 0 0-.11.22a.48.48 0 0 0 0 .29c.04.09.102.17.18.23a.54.54 0 0 0 .28.08a.51.51 0 0 0 .5-.5a.54.54 0 0 0-.08-.28a.58.58 0 0 0-.23-.18a.48.48 0 0 0-.29 0m.23 2.05h-.27l-.87 1.94h2zm2.2 4.94l-.89-2h-2.88l-.89 2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reactions(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 7.5c0 .169-.01.336-.027.5h1.005A5.5 5.5 0 1 0 8 12.978v-1.005A4.5 4.5 0 1 1 12 7.5M5.5 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2m2 2.5c.712 0 1.355-.298 1.81-.776l.707.708A3.49 3.49 0 0 1 7.5 10.5a3.49 3.49 0 0 1-2.555-1.108l.707-.708A2.494 2.494 0 0 0 7.5 9.5m2-2.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m2.5 3h1v2h2v1h-2v2h-1v-2h-2v-1h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M8.6 1c1.6.1 3.1.9 4.2 2c1.3 1.4 2 3.1 2 5.1c0 1.6-.6 3.1-1.6 4.4c-1 1.2-2.4 2.1-4 2.4c-1.6.3-3.2.1-4.6-.7c-1.4-.8-2.5-2-3.1-3.5C.9 9.2.8 7.5 1.3 6c.5-1.6 1.4-2.9 2.8-3.8C5.4 1.3 7 .9 8.6 1m.5 12.9c1.3-.3 2.5-1 3.4-2.1c.8-1.1 1.3-2.4 1.2-3.8c0-1.6-.6-3.2-1.7-4.3c-1-1-2.2-1.6-3.6-1.7c-1.3-.1-2.7.2-3.8 1c-1.1.8-1.9 1.9-2.3 3.3c-.4 1.3-.4 2.7.2 4c.6 1.3 1.5 2.3 2.7 3c1.2.7 2.6.9 3.9.6" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecordKeys(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 3H3a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1m0 8H3V4h11zm-3-6h-1v1h1zm-1 2H9v1h1zm2-2h1v1h-1zm1 4h-1v1h1zM6 9h5v1H6zm7-2h-2v1h2zM8 5h1v1H8zm0 2H7v1h1zM4 9h1v1H4zm0-4h1v1H4zm3 0H6v1h1zM4 7h2v1H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecordSmall(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-1 0a3 3 0 1 0-6 0a3 3 0 0 0 6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.5 2v3.5L12 6H8.5V5h2.521l-.941-.941a3.552 3.552 0 1 0-5.023 5.023l5.197 5.198l-.72.72l-5.198-5.198A4.57 4.57 0 0 1 10.8 3.339l.7.7V2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func References(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.105 4.561l-3.43 3.427l-1.134-1.12l2.07-2.08h-4.8a2.4 2.4 0 1 0 0 4.8h.89v1.6h-.88a4 4 0 0 1 0-7.991h4.8L6.54 1.13L7.675 0l3.43 3.432zM16.62 24h-9.6l-.8-.8V10.412l.8-.8h9.6l.8.8V23.2zm-8.8-1.6h8V11.212h-8zm5.6-20.798h9.6l.8.8v12.786l-.8.8h-4v-1.6h3.2V3.2h-8v4.787h-1.6V2.401l.8-.8zm.8 11.186h-4.8v1.6h4.8zm-4.8 3.2h4.8v1.6h-4.8zm4.8 3.2h-4.8v1.6h4.8zm1.6-14.4h4.8v1.6h-4.8zm4.8 6.4h-1.6v1.6h1.6zm-3.338-3.176v-.024h3.338v1.6h-1.762z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.681 3H2V2h3.5l.5.5V6H5V4a5 5 0 1 0 4.53-.761l.302-.954A6 6 0 1 1 4.681 3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Regex(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.012 2h.976v3.113l2.56-1.557l.486.885L11.47 6l2.564 1.559l-.485.885l-2.561-1.557V10h-.976V6.887l-2.56 1.557l-.486-.885L9.53 6L6.966 4.441l.485-.885l2.561 1.557zM2 10h4v4H2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remote(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.904 9.57L8.928 5.596l3.976-3.976l-.619-.62L8 5.286v.619l4.285 4.285l.62-.618zM3 5.62l4.072 4.07L3 13.763l.619.618L8 10v-.619L3.619 5L3 5.619z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoteExplorer(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.344 2.125h20.312l.782.781v8.599a7.825 7.825 0 0 0-1.563-.912V3.688H2.125V17.75h7.813a7.813 7.813 0 0 0 1.562 4.688H5.25v-1.563h4.688v-1.563H1.344l-.782-.78V2.905l.782-.781zM17.75 11.5a6.25 6.25 0 1 0 0 12.5a6.25 6.25 0 0 0 0-12.5m0 10.938a4.688 4.688 0 1 1 0-9.377a4.688 4.688 0 0 1 0 9.377m2.603-3.132L18.2 17.153L20.353 15l.647.646l-1.506 1.507L21 18.659zM15 17.246l1.506 1.507L15 20.259l.647.647l2.153-2.153l-2.153-2.153z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remove(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 8H1V7h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Replace(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.221 3.739l2.261 2.269L7.7 3.784l-.7-.7l-1.012 1.007l-.008-1.6a.523.523 0 0 1 .5-.526H8V1H6.48A1.482 1.482 0 0 0 5 2.489V4.1L3.927 3.033zm6.67 1.794h.01c.183.311.451.467.806.467c.393 0 .706-.168.94-.503c.236-.335.353-.78.353-1.333c0-.511-.1-.913-.301-1.207c-.201-.295-.488-.442-.86-.442c-.405 0-.718.194-.938.581h-.01V1H9v4.919h.89zm-.015-1.061v-.34c0-.248.058-.448.175-.601a.54.54 0 0 1 .445-.23a.49.49 0 0 1 .436.233c.104.154.155.368.155.643c0 .33-.056.587-.169.768a.524.524 0 0 1-.47.27a.495.495 0 0 1-.411-.211a.853.853 0 0 1-.16-.532zM9 12.769c-.256.154-.625.231-1.108.231c-.563 0-1.02-.178-1.369-.533c-.349-.355-.523-.813-.523-1.374c0-.648.186-1.158.56-1.53c.374-.376.875-.563 1.5-.563c.433 0 .746.06.94.179v.998a1.26 1.26 0 0 0-.792-.276c-.325 0-.583.1-.774.298c-.19.196-.283.468-.283.816c0 .338.09.603.272.797c.182.191.431.287.749.287c.282 0 .558-.092.828-.276zM4 7L3 8v6l1 1h7l1-1V8l-1-1zm0 1h7v6H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplaceAll(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.6 2.677c.147-.31.356-.465.626-.465c.248 0 .44.118.573.353c.134.236.201.557.201.966c0 .443-.078.798-.235 1.067c-.156.268-.365.402-.627.402c-.237 0-.416-.125-.537-.374h-.008v.31H11V1h.593v1.677zm-.016 1.1a.78.78 0 0 0 .107.426c.071.113.163.169.274.169c.136 0 .24-.072.314-.216c.075-.145.113-.35.113-.615c0-.22-.035-.39-.104-.514c-.067-.124-.164-.187-.29-.187c-.12 0-.219.062-.297.185a.886.886 0 0 0-.117.48zM4.12 7.695L2 5.568l.662-.662l1.006 1v-1.51A1.39 1.39 0 0 1 5.055 3H7.4v.905H5.055a.49.49 0 0 0-.468.493l.007 1.5l.949-.944l.656.656l-2.08 2.085zM9.356 4.93H10V3.22C10 2.408 9.685 2 9.056 2c-.135 0-.285.024-.45.073a1.444 1.444 0 0 0-.388.167v.665c.237-.203.487-.304.75-.304c.261 0 .392.156.392.469l-.6.103c-.506.086-.76.406-.76.961c0 .263.061.473.183.631A.61.61 0 0 0 8.69 5c.29 0 .509-.16.657-.48h.009zm.004-1.355v.193a.75.75 0 0 1-.12.436a.368.368 0 0 1-.313.17a.276.276 0 0 1-.22-.095a.38.38 0 0 1-.08-.248c0-.222.11-.351.332-.389zM7 12.93h-.644v-.41h-.009c-.148.32-.367.48-.657.48a.61.61 0 0 1-.507-.235c-.122-.158-.183-.368-.183-.63c0-.556.254-.876.76-.962l.6-.103c0-.313-.13-.47-.392-.47c-.263 0-.513.102-.75.305v-.665c.095-.063.224-.119.388-.167c.165-.049.315-.073.45-.073c.63 0 .944.407.944 1.22zm-.64-1.162v-.193l-.4.068c-.222.037-.333.166-.333.388c0 .1.027.183.08.248a.276.276 0 0 0 .22.095a.368.368 0 0 0 .312-.17c.08-.116.12-.26.12-.436zM9.262 13c.321 0 .568-.058.738-.173v-.71a.9.9 0 0 1-.552.207a.619.619 0 0 1-.5-.215c-.12-.145-.181-.345-.181-.598c0-.26.063-.464.189-.612a.644.644 0 0 1 .516-.223c.194 0 .37.069.528.207v-.749c-.129-.09-.338-.134-.626-.134c-.417 0-.751.14-1.001.422c-.249.28-.373.662-.373 1.148c0 .42.116.764.349 1.03c.232.267.537.4.913.4M2 9l1-1h9l1 1v5l-1 1H3l-1-1zm1 0v5h9V9zm3-2l1-1h7l1 1v5l-1 1V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.306 2.146l-4.02 4.02v.708l4.02 4.02l.708-.707L3.807 6.98H5.69c2.813 0 4.605.605 5.705 1.729c1.102 1.125 1.615 2.877 1.615 5.421v.35h1v-.35c0-2.646-.527-4.72-1.9-6.121C10.735 6.605 8.617 5.98 5.69 5.98H3.887l3.127-3.126z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repo(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 10V1.5l-.5-.5H3.74a1.9 1.9 0 0 0-.67.13a1.77 1.77 0 0 0-.94 1a1.7 1.7 0 0 0-.13.62v9.5a1.7 1.7 0 0 0 .13.67c.177.427.515.768.94.95a1.9 1.9 0 0 0 .67.13H4v-1h-.26a.72.72 0 0 1-.29-.06a.74.74 0 0 1-.4-.4a.93.93 0 0 1-.05-.29v-.5a.93.93 0 0 1 .05-.29a.74.74 0 0 1 .4-.4a.72.72 0 0 1 .286-.06H13v2H9v1h4.5l.5-.5zM4 10V2h9v8zm1-7h1v1H5zm0 2h1v1H5zm1 2H5v1h1zm.5 6.49L5.28 15H5v-3h3v3h-.28z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RepoClone(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 10H4V2h4V1H3.74a1.9 1.9 0 0 0-.67.13a1.66 1.66 0 0 0-.57.41a1.73 1.73 0 0 0-.37.59a1.68 1.68 0 0 0-.13.62v9.5a1.75 1.75 0 0 0 1.07 1.62a1.9 1.9 0 0 0 .67.13H4v-1h-.26a.72.72 0 0 1-.29-.06a.78.78 0 0 1-.4-.4a.93.93 0 0 1 0-.29v-.5a.93.93 0 0 1 0-.29a.78.78 0 0 1 .4-.4a.72.72 0 0 1 .29-.06H13v2H9v1h4.5l.5-.5V9h-1zM6 3H5v1h1zM5 5h1v1H5zm0 2h1v1H5zm.28 8H5v-3h3v3h-.28L6.5 13.49zM10 1h4.5l.5.5v6l-.5.5H12v1h-1V8h-1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1m.5 6h.5V6h-.5a.5.5 0 0 0 0 1M12 7h2V6h-2zm-1-2h3V2h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RepoForcePush(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.74 1h9.76l.5.5v12l-.5.5H10v-1h3v-2h-3v-1h3V2H4v8h3v1H3.74a.74.74 0 0 0-.74.75v.5a.74.74 0 0 0 .74.75H7v1H3.74A1.74 1.74 0 0 1 2 12.25v-9.5A1.74 1.74 0 0 1 3.74 1m1.6 4.83l.71.7L8 4.58v1.45L5.38 8.65l.71.7l1.92-1.92V15h1V7.328l2.03 2.022l.7-.7L9 5.9V4.538l2 1.992l.7-.7L8.88 3h-.71z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RepoForked(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4a2 2 0 1 0-2.47 1.94V7a.48.48 0 0 1-.27.44L8.49 8.88l-2.76-1.4A.49.49 0 0 1 5.46 7V5.94a2 2 0 1 0-1 0V7a1.51 1.51 0 0 0 .82 1.34L8 9.74v1.32a2 2 0 1 0 1 0V9.74l2.7-1.36A1.49 1.49 0 0 0 12.52 7V5.92A2 2 0 0 0 14 4M4 4a1 1 0 1 1 2 0a1 1 0 0 1-2 0m5.47 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0M12 5a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RepoPull(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 1.5V3h-1V2H3v8h10v3.5l-.5.5H8v-1h4v-2H2.735a.72.72 0 0 0-.285.06a.74.74 0 0 0-.4.4a.93.93 0 0 0-.05.29v.5a.93.93 0 0 0 .05.29a.74.74 0 0 0 .4.4c.091.04.19.06.29.06H3v1h-.26a1.9 1.9 0 0 1-.67-.13a1.77 1.77 0 0 1-.94-.95a1.7 1.7 0 0 1-.13-.67v-9.5a1.7 1.7 0 0 1 .13-.62a1.77 1.77 0 0 1 .94-1A1.9 1.9 0 0 1 2.74 1h9.76zM2 10.17V2.748zM5 3H4v1h1zm0 2H4v1h1zM4 7h1v1H4zm8.07-3.61l-.7.71l1.92 1.92H7v1h6.39l-2.02 2.03l.7.7l2.83-2.82v-.71zM5.5 13.49L4.28 15H4v-3h3v3h-.28z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RepoPush(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 1H3.74A1.74 1.74 0 0 0 2 2.75v9.5A1.74 1.74 0 0 0 3.74 14H7v-1H3.74a.74.74 0 0 1-.74-.75v-.5a.74.74 0 0 1 .74-.75H7v-1H4V2h9v8h-3v1h3v2h-3v1h3.5l.5-.5v-12zM3 2.73a.75.75 0 0 0 0 .02v7.42zM6 3H5v1h1zm-.62 5.65l.71.7l1.92-1.92V15h1V7.328l2.03 2.022l.7-.7l-2.82-2.83h-.71zM5 5h1v1H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Report(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 1h13l.5.5v10l-.5.5H7.707l-2.853 2.854L4 14.5V12H1.5l-.5-.5v-10zm6 10H14V2H2v9h2.5l.5.5v1.793l2.146-2.147zm0-8h1v5h-1zm0 7h1V9h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RequestChanges(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.71 1.29l3 3L14 5v9l-1 1H4l-1-1V2l1-1h6zM4 14h9V5l-3-3H4zm4-8H6v1h2v2h1V7h2V6H9V4H8zm-2 5h5v1H6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Robot(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0M5 8a1 1 0 1 0 2 0a1 1 0 0 0-2 0m3 0a2 2 0 1 0 4 0a2 2 0 0 0-4 0m3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-2.49 3.251a2.71 2.71 0 0 0 1.37-.74l.7.71A3.7 3.7 0 0 1 8 12.291a3.761 3.761 0 0 1-1.42-.29a3.641 3.641 0 0 1-1.19-.79l.7-.71a2.71 2.71 0 0 0 2.42.74z"/><path d="M9.5 1.5a1.5 1.5 0 0 1-1 1.415V3H11a3 3 0 0 1 3 3v1l1 1v2l-1 1v1a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-1l-1-1V8l1-1V6a3 3 0 0 1 3-3h2.5v-.085a1.5 1.5 0 1 1 2-1.415M5 4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.491 1c-3.598.004-6.654 1.983-8.835 4H1.5l-.5.5v3l.147.354l.991.991l.001.009l4 4l.009.001l.999.999L7.5 15h3l.5-.5v-4.154c2.019-2.178 3.996-5.233 3.992-8.846zM2 6h2.643a23.828 23.828 0 0 0-2.225 2.71L2 8.294zm5.7 8l-.42-.423a23.59 23.59 0 0 0 2.715-2.216V14zm-1.143-1.144L3.136 9.437C4.128 8 8.379 2.355 13.978 2.016c-.326 5.612-5.987 9.853-7.421 10.84M4 15v-1H2v-2H1v3zm6.748-7.667a1.5 1.5 0 1 0-2.496-1.666a1.5 1.5 0 0 0 2.495 1.666z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RootFolder(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.71 3h6.79l.51.5v10l-.5.5H8.743a5.48 5.48 0 0 0 .657-1h4.59v-1.51l.01-4v-1.5H7.69l-.017.017a5.494 5.494 0 0 0-.881-.508l.348-.349l.35-.15h6.5l.01-.99H7.5l-.36-.15l-.85-.85H2V5.6a5.45 5.45 0 0 0-.99.649V2.5l.5-.5h5l.35.15z" clip-rule="evenodd"/><path d="M6 10.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/><path fill-rule="evenodd" d="M8 10.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0M4.5 13a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RootFolderOpened(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1 6.257V2.5l.5-.5h5l.35.15l.86.85h5.79l.5.5V6h1.13l.48.63l-2.63 7l-.48.37H8.743a5.48 5.48 0 0 0 .657-1h2.73l2.37-6H8.743a5.534 5.534 0 0 0-.72-.724l.127-.126L8.5 6H13V4H7.5l-.35-.15L6.29 3H2l.01 2.594c-.361.184-.7.407-1.01.663" clip-rule="evenodd"/><path d="M6 10.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/><path fill-rule="evenodd" d="M8 10.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0M4.5 13a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 13H3v-2c1.11 0 2 .89 2 2M3 3v1a9 9 0 0 1 9 9h1C13 7.48 8.52 3 3 3m0 4v1c2.75 0 5 2.25 5 5h1c0-3.31-2.69-6-6-6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruby(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1 7.19l6.64 6.64h.72L15 7.19v-.72l-3.32-3.32l-.36-.15H4.68l-.36.15L1 6.47zm7 5.56L2.08 6.83L4.89 4h6.22l2.81 2.83zm0-7.73h2.69l1.81 1.81l-4.5 4.4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RunAbove(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.77 1.01L1 1.42v12l.78.42l9-6v-.83zM2 12.49V2.36l7.6 5.07zM12.15 8h.71l2.5 2.5l-.71.71L13 9.56V15h-1V9.55l-1.65 1.65l-.7-.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RunAll(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2.78 2L2 2.41v12l.78.42l9-6V8zM3 13.48V3.35l7.6 5.07z"/><path fill-rule="evenodd" d="m6 14.683l8.78-5.853V8L6 2.147V3.35l7.6 5.07L6 13.48z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RunBelow(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.8 1.01l-.78.41v12l.78.42l9-6v-.83zm.22 11.48V2.36l7.6 5.07zM12.85 15h-.71l-2.5-2.5l.71-.71L12 13.44V8h1v5.45l1.65-1.65l.71.71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RunErrors(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 2.41L5.78 2l9 6v.83L9 12.683v-1.2l4.6-3.063L6 3.35V7H5z"/><path d="M4.872 7.808c-.85-.053-1.705.159-2.403.641c-.753.485-1.233 1.184-1.499 2.03c-.269.81-.213 1.72.106 2.518a3.774 3.774 0 0 0 1.658 1.873c.756.432 1.616.537 2.467.378c.861-.162 1.61-.645 2.143-1.285l.005-.006c.529-.687.852-1.489.852-2.352a3.882 3.882 0 0 0-1.066-2.72l-.006-.005c-.585-.585-1.388-1.018-2.257-1.072M2.951 9.183c.512-.373 1.172-.517 1.792-.47h.001c.656.048 1.22.328 1.697.804c.516.516.803 1.274.803 2.038v.014c.047.649-.183 1.26-.566 1.789c-.426.518-.993.85-1.61.991a2.512 2.512 0 0 1-1.83-.282c-.572-.333-1-.808-1.288-1.43c-.28-.607-.282-1.265-.091-1.885v-.004a2.865 2.865 0 0 1 1.092-1.565m1.554 1.83L3.292 9.748l-.638.637l1.22 1.27l-1.22 1.27l.638.637L4.505 12.3l1.213 1.264l.638-.637l-1.219-1.27l1.219-1.27l-.638-.637z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.353 1.146l1.5 1.5L15 3v11.5l-.5.5h-13l-.5-.5v-13l.5-.5H13zM2 2v12h12V3.208L12.793 2H11v4H4V2zm6 0v3h2V2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveAll(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.85 2.65l-1.5-1.5L13 1H4.48l-.5.5V4H1.5l-.5.5v10l.5.5h10l.5-.5V12h2.5l.5-.5V3zM11 14H2V5h1v3.07h6V5h.79L11 6.21zM6 7V5h2v2zm8 4h-2V6l-.15-.35l-1.5-1.5L10 4H5V2h7.81l1.21 1.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveAs(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.04 1.33L12.71 3l.29.71v.33h-.5l-.5.5v-.83l-1.67-1.67H10v4H4v-4H2v10h3l-.5 1H2l-1-1v-10l1-1h8.33zM7 5h2V2H7zm6.5 0L15 6.5l-.02.69l-5.5 5.5l-.13.12l-.37.37l-.1.09l-3 1.5l-.67-.67l1.5-3l.09-.1l.37-.37l.12-.13l5.5-5.5zm-6.22 7.24l-.52 1l1.04-.48zm.69-1.03l.79.79l5.15-5.15l-.79-.79z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenFull(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 12h10V4H3zm2-6h6v4H5zM2 6H1V2.5l.5-.5H5v1H2zm13-3.5V6h-1V3h-3V2h3.5zM14 10h1v3.5l-.5.5H11v-1h3zM2 13h3v1H1.5l-.5-.5V10h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenNormal(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 4H1V3h2V1h1v2.5zM13 3V1h-1v2.5l.5.5H15V3zm-1 9.5V15h1v-2h2v-1h-2.5zM1 12v1h2v2h1v-2.5l-.5-.5zm11-1.5l-.5.5h-7l-.5-.5v-5l.5-.5h7l.5.5zM10 7H6v2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 0a8.25 8.25 0 0 0-6.18 13.72L1 22.88l1.12 1l8.05-9.12A8.251 8.251 0 1 0 15.25.01zm0 15a6.75 6.75 0 1 1 0-13.5a6.75 6.75 0 0 1 0 13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchFuzzy(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1.5a4.5 4.5 0 0 0-3.263 7.6l-3.694 3.694l.707.707l3.755-3.755A4.5 4.5 0 1 0 8 1.5M4.5 6a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0M9 13.855L6.854 16h-.708l-1.5-1.5l.708-.707L6.5 14.94l2.146-2.146h.708L11.5 14.94l2.146-2.146h.707L16 14.44v1.415l-2-2L11.854 16h-.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchStop(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.738 3.318a4.5 4.5 0 0 0-.877 5.123A4.48 4.48 0 0 0 6.1 10a4.62 4.62 0 0 0-.1 1v.17c-.16-.11-.32-.22-.47-.34L1.75 14.5L1 13.84l3.8-3.69a5.5 5.5 0 1 1 9.62-3.65c0 .268-.02.535-.06.8a5.232 5.232 0 0 0-.94-.68V6.5a4.5 4.5 0 0 0-7.682-3.182m3.04 4.356a4 4 0 1 1 4.444 6.652a4 4 0 0 1-4.444-6.652m.1 5.447A3 3 0 0 0 11 14a3 3 0 0 0 1.74-.55L8.55 9.26A3 3 0 0 0 8 11a3 3 0 0 0 .879 2.121zm.382-4.57l4.19 4.189A3 3 0 0 0 14 11a3 3 0 0 0-3-3a3 3 0 0 0-1.74.55z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Send(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1 1.91l.78-.41L15 7.449v.95L1.78 14.33L1 13.91L2.583 8zM3.612 8.5L2.33 13.13L13.5 7.9L2.33 2.839l1.282 4.6L9 7.5v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 5L2 4.5v-3l.5-.5h11l.5.5v3l-.5.5zM10 2H9v1H8V2H7v1H6V2H5v1H4V2H3v2h10V2h-2v1h-1zm-7.5 8L2 9.5v-3l.5-.5h11l.5.5v3l-.5.5zM6 7H5v1H4V7H3v2h10V7h-2v1h-1V7H9v1H8V7H7v1H6zm7.5 8l.5-.5v-3l-.5-.5h-11l-.5.5v3l.5.5zM3 14v-2h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h2v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerEnvironment(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3h4v1H6zm0 6h4v1H6zm0 2h4v1H6zm9.14 5H.86l1.25-5H4V2a.95.95 0 0 1 .078-.383c.052-.12.123-.226.211-.32a.922.922 0 0 1 .32-.219A1.01 1.01 0 0 1 5 1h6a.95.95 0 0 1 .383.078c.12.052.226.123.32.211a.922.922 0 0 1 .219.32c.052.125.078.256.078.391v9h1.89zM5 13h6V2H5zm8.86 2l-.75-3H12v2H4v-2H2.89l-.75 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerProcess(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2h13l.5.5V9h-1V6H2v7h7v1H1.5l-.5-.5v-11zM2 5h12V3H2zm5 7v-1.094a1.633 1.633 0 0 1-.469-.265l-.945.539l-.5-.86l.937-.547a1.57 1.57 0 0 1 0-.547l-.937-.546l.5-.86l.945.54c.151-.12.308-.209.469-.266V7h1v1.094a1.48 1.48 0 0 1 .469.265l.945-.539l.5.86l-.937.547a1.57 1.57 0 0 1 0 .546l.937.547l-.5.86l-.945-.54a1.807 1.807 0 0 1-.469.266V12zm-.25-2.5c0 .208.073.385.219.531a.723.723 0 0 0 .531.219a.723.723 0 0 0 .531-.219a.723.723 0 0 0 .219-.531a.723.723 0 0 0-.219-.531a.723.723 0 0 0-.531-.219a.723.723 0 0 0-.531.219a.723.723 0 0 0-.219.531m5.334 5.5v-1.094a1.634 1.634 0 0 1-.469-.265l-.945.539l-.5-.86l.938-.547a1.572 1.572 0 0 1 0-.547l-.938-.546l.5-.86l.945.54c.151-.12.308-.209.47-.266V10h1v1.094a1.486 1.486 0 0 1 .468.265l.945-.539l.5.86l-.937.547a1.562 1.562 0 0 1 0 .546l.937.547l-.5.86l-.945-.54a1.806 1.806 0 0 1-.469.266V15zm-.25-2.5c0 .208.073.385.219.531a.723.723 0 0 0 .531.219a.723.723 0 0 0 .531-.219a.723.723 0 0 0 .22-.531a.723.723 0 0 0-.22-.531a.723.723 0 0 0-.53-.219a.723.723 0 0 0-.532.219a.723.723 0 0 0-.219.531" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2h-1v5h1zm6.1 5H6.4L6 6.45v-1L6.4 5h3.2l.4.5v1zm-5 3H1.4L1 9.5v-1l.4-.5h3.2l.4.5v1zm3.9-8h-1v2h1zm-1 6h1v6h-1zm-4 3h-1v3h1zm7.9 0h3.19l.4-.5v-.95l-.4-.5H11.4l-.4.5v.95zm2.1-9h-1v6h1zm-1 10h1v2h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsGear(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m19.85 8.75l4.15.83v4.84l-4.15.83l2.35 3.52l-3.43 3.43l-3.52-2.35l-.83 4.15H9.58l-.83-4.15l-3.52 2.35l-3.43-3.43l2.35-3.52L0 14.42V9.58l4.15-.83L1.8 5.23L5.23 1.8l3.52 2.35L9.58 0h4.84l.83 4.15l3.52-2.35l3.43 3.43zm-1.57 5.07l4-.81v-2l-4-.81l-.54-1.3l2.29-3.43l-1.43-1.43l-3.43 2.29l-1.3-.54l-.81-4h-2l-.81 4l-1.3.54l-3.43-2.29l-1.43 1.43L6.38 8.9l-.54 1.3l-4 .81v2l4 .81l.54 1.3l-2.29 3.43l1.43 1.43l3.43-2.29l1.3.54l.81 4h2l.81-4l1.3-.54l3.43 2.29l1.43-1.43l-2.29-3.43zm-8.186-4.672A3.43 3.43 0 0 1 12 8.57A3.44 3.44 0 0 1 15.43 12a3.43 3.43 0 1 1-5.336-2.852m.956 4.274c.281.188.612.288.95.288A1.7 1.7 0 0 0 13.71 12a1.71 1.71 0 1 0-2.66 1.422" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 4.001H1.5l-.5.5v10l.5.5h11l.5-.5V11.5h-1v2.501H2v-9h3z"/><path fill-rule="evenodd" d="M6 11H5V8.5a5.002 5.002 0 0 1 4-4.9V2.349a1.349 1.349 0 0 1 2.333-.923l3.646 3.89v1.368l-3.646 3.89A1.349 1.349 0 0 1 9 9.65V8.6c-.933.211-1.384.711-1.632 1.145A2.925 2.925 0 0 0 7 11zm4-6.5a4 4 0 0 0-4 4V10h.176c.138-.443.384-.965.824-1.417c.442-.455 1.08-.839 2-1a5.762 5.762 0 0 1 1-.083v2.151a.349.349 0 0 0 .603.239L14.25 6l-3.647-3.89a.349.349 0 0 0-.603.24z" clip-rule="evenodd"/><path d="M7 11c.052 0 .06 0 0 0m0 0v.008v-.006z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.246 14.713a27.792 27.792 0 0 1-1.505-.953c-.501-.34-.983-.707-1.444-1.1c-.458-.395-.888-.82-1.288-1.274c-.4-.455-.753-.95-1.05-1.478a7.8 7.8 0 0 1-.7-1.69A7.041 7.041 0 0 1 2 6.3V3.1l.5-.5c.333 0 .656-.011.97-.036c.296-.023.591-.066.882-.128c.284-.062.562-.148.832-.256c.284-.118.557-.261.816-.427a4.83 4.83 0 0 1 1.184-.565a4.8 4.8 0 0 1 2-.142a4.018 4.018 0 0 1 1.237.383c.199.097.392.204.58.322c.26.167.535.31.821.428c.27.109.547.194.831.256c.291.062.587.106.884.129c.311.024.634.035.967.035l.5.5v3.2a7.043 7.043 0 0 1-.256 1.919a7.804 7.804 0 0 1-.7 1.69a8.751 8.751 0 0 1-1.05 1.478c-.4.452-.829.877-1.286 1.27a15.94 15.94 0 0 1-1.448 1.1a28.71 28.71 0 0 1-1.51.956h-.508zM3 3.59V6.3c-.004.555.07 1.11.22 1.645a6.7 6.7 0 0 0 .61 1.473c.263.467.575.905.93 1.308c.37.417.766.81 1.188 1.174c.432.368.883.712 1.352 1.03c.4.267.8.523 1.2.769c.4-.242.8-.498 1.2-.768c.47-.319.923-.663 1.355-1.031c.421-.364.817-.756 1.186-1.172a7.8 7.8 0 0 0 .93-1.308c.261-.465.466-.96.61-1.473c.15-.537.223-1.09.22-1.647V3.59c-.159 0-.313-.012-.465-.023l-.079-.006a7.95 7.95 0 0 1-1.018-.147a6.112 6.112 0 0 1-1.976-.814a5.166 5.166 0 0 0-.482-.27a3.123 3.123 0 0 0-.943-.29a3.686 3.686 0 0 0-1.558.106c-.332.108-.649.26-.94.452c-.312.2-.64.372-.983.513a6.4 6.4 0 0 1-1 .307c-.335.07-.675.12-1.017.146c-.174.01-.355.02-.54.026m6.065 4.3a1.5 1.5 0 1 0-1.13 0L7.5 10.5h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignIn(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.02 3.77l.01-.01l.99.99V2.5l-.5-.5h-9l-.51.5v.493L2 3v10.29l.36.46l5 1.72L8 15v-1h3.52l.5-.5v-2.25l-1 1V13H8V4.71l-.33-.46L4.036 3h6.984zM7 14.28l-4-1.34V3.72l4 1.34zm3.09-6.75h4.97v1h-4.93l1.59 1.6l-.71.7l-2.47-2.46v-.71l2.49-2.48l.7.7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignOut(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.02 3.77v1.56l1-.99V2.5l-.5-.5h-9l-.5.5v.486L2 3v10.29l.36.46l5 1.72L8 15v-1h3.52l.5-.5v-1.81l-1-1V13H8V4.71l-.33-.46L4.036 3h6.984zM7 14.28l-4-1.34V3.72l4 1.34zm6.52-5.8H8.55v-1h4.93l-1.6-1.6l.71-.7l2.47 2.46v.71l-2.49 2.48l-.7-.7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smiley(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.111 2.18a7 7 0 1 1 7.778 11.64A7 7 0 0 1 4.11 2.18zm.556 10.809a6 6 0 1 0 6.666-9.978a6 6 0 0 0-6.666 9.978M6.5 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0m5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0M8 11a3 3 0 0 1-2.65-1.58l-.87.48a4 4 0 0 0 7.12-.16l-.9-.43A3 3 0 0 1 8 11" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snake(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0"/><path fill-rule="evenodd" d="M5.5 1a2.5 2.5 0 0 0-2.086 3.879L2.293 6H.5v1H2v1.5h1V6.707l1.121-1.121c.396.262.87.414 1.379.414H7v2H6a2 2 0 0 0-2 2a2 2 0 1 0 0 4h9a2 2 0 1 0 0-4a2 2 0 0 0-2-2V4.5A3.5 3.5 0 0 0 7.5 1zM4 3.5A1.5 1.5 0 0 1 5.5 2h2A2.5 2.5 0 0 1 10 4.5v4l.5.5h.5a1 1 0 0 1 1 1v1h1a1 1 0 1 1 0 2H4a1 1 0 1 1 0-2h1v-1a1 1 0 0 1 1-1h1.5l.5-.5v-3L7.5 5h-2A1.5 1.5 0 0 1 4 3.5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortPrecedence(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 2L6 3v3h1V3h7v2.453l.207-.16l.793.793V3l-1-1zm1 2h2v2H8zM5 9H3v2h2zM2 7L1 8v5l1 1h7l1-1V8L9 7zm0 6V8h7v5zm6-3H6v2h2zm5-6h-1v3.864l-1.182-1.182l-.707.707l2.035 2.036h.708l2.035-2.036l-.707-.707L13 7.864z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SourceControl(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.007 8.222A3.738 3.738 0 0 0 15.045 5.2a3.737 3.737 0 0 0 1.156 6.583a2.988 2.988 0 0 1-2.668 1.67h-2.99a4.456 4.456 0 0 0-2.989 1.165V7.4a3.737 3.737 0 1 0-1.494 0v9.117a3.776 3.776 0 1 0 1.816.099a2.99 2.99 0 0 1 2.668-1.667h2.99a4.484 4.484 0 0 0 4.223-3.039a3.736 3.736 0 0 0 3.25-3.687zM4.565 3.738a2.242 2.242 0 1 1 4.484 0a2.242 2.242 0 0 1-4.484 0m4.484 16.441a2.242 2.242 0 1 1-4.484 0a2.242 2.242 0 0 1 4.484 0m8.221-9.715a2.242 2.242 0 1 1 0-4.485a2.242 2.242 0 0 1 0 4.485"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sparkle(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.398 10.807a1.042 1.042 0 0 0 1.204-.003c.178-.13.313-.31.387-.518l.447-1.373a2.336 2.336 0 0 1 1.477-1.479l1.391-.45a1.045 1.045 0 0 0-.044-1.98l-1.375-.448a2.335 2.335 0 0 1-1.48-1.477l-.452-1.388a1.044 1.044 0 0 0-1.973.017l-.457 1.4a2.336 2.336 0 0 1-1.44 1.45l-1.39.447a1.045 1.045 0 0 0 .016 1.974l1.374.445a2.333 2.333 0 0 1 1.481 1.488l.452 1.391c.072.204.206.38.382.504m.085-7.415l.527-1.377l.44 1.377a3.331 3.331 0 0 0 2.117 2.114l1.406.53l-1.382.447A3.343 3.343 0 0 0 6.476 8.6l-.523 1.378L5.504 8.6a3.336 3.336 0 0 0-.8-1.31a3.373 3.373 0 0 0-1.312-.812l-1.378-.522l1.386-.45a3.358 3.358 0 0 0 1.29-.813a3.4 3.4 0 0 0 .793-1.3m6.052 11.457a.806.806 0 0 0 1.226-.398l.248-.762c.053-.158.143-.302.26-.42c.118-.12.262-.208.42-.26l.772-.252a.8.8 0 0 0-.023-1.52l-.764-.25a1.075 1.075 0 0 1-.68-.678l-.252-.773a.8.8 0 0 0-1.518.01l-.247.762a1.068 1.068 0 0 1-.665.679l-.773.252a.8.8 0 0 0 .008 1.518l.763.247c.16.054.304.143.422.261c.119.119.207.263.258.422l.253.774a.8.8 0 0 0 .292.388m-.913-2.793l-.179-.059l.184-.064a2.091 2.091 0 0 0 1.3-1.317l.058-.178l.06.181a2.078 2.078 0 0 0 1.316 1.316l.195.063l-.18.06a2.076 2.076 0 0 0-1.317 1.32L12 13.56l-.058-.18a2.075 2.075 0 0 0-1.32-1.323"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SparkleFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.398 10.807a1.039 1.039 0 0 0 1.204-.003c.178-.13.313-.31.387-.518l.447-1.373c.115-.344.308-.657.564-.914a2.35 2.35 0 0 1 .913-.565l1.391-.451a1.047 1.047 0 0 0 .645-.67a1.038 1.038 0 0 0-.689-1.31l-1.375-.447a2.338 2.338 0 0 1-1.48-1.477l-.452-1.388a1.043 1.043 0 0 0-1.717-.39a1.054 1.054 0 0 0-.256.407l-.457 1.4a2.323 2.323 0 0 1-1.44 1.449l-1.391.448a1.061 1.061 0 0 0-.644.67a1.051 1.051 0 0 0 .144.918c.128.18.309.315.517.386l1.374.445a2.331 2.331 0 0 1 1.481 1.488l.452 1.391c.072.204.206.38.382.504m6.137 4.042a.8.8 0 0 0 .926.002a.803.803 0 0 0 .3-.4l.248-.762a1.066 1.066 0 0 1 .68-.68l.772-.252a.793.793 0 0 0 .531-.64a.796.796 0 0 0-.554-.881l-.764-.249a1.075 1.075 0 0 1-.68-.678l-.252-.773a.784.784 0 0 0-.293-.39a.796.796 0 0 0-1.03.085a.8.8 0 0 0-.195.315l-.247.762a1.071 1.071 0 0 1-.665.679l-.773.252a.797.797 0 0 0-.543.762a.799.799 0 0 0 .551.756l.763.247c.159.054.304.143.422.261c.119.119.207.263.258.422l.253.774a.8.8 0 0 0 .292.388"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SplitHorizontal(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1H3L2 2v11l1 1h11l1-1V2zM8 13H3V2h5zm6 0H9V2h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SplitVertical(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1H3L2 2v11l1 1h11l1-1V2zm0 12H3V8h11zm0-6H3V2h11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Squirrel(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.558 2.642a3.698 3.698 0 0 0-.123-.01A1.47 1.47 0 0 0 3.999 1.52v1.307a4.898 4.898 0 0 0-2.993 3.587v.39c.459.836 1.906 1.13 2.154 1.18c.027.006.04.009.035.009c-2.419.32-2.19 2.249-2.19 2.249a1 1 0 0 0 1 .93c.272-.019.538-.08.79-.18h2.06a3 3 0 0 0-.36 1h-.32a2.55 2.55 0 0 0-2.17 2.528a.42.42 0 0 0 .39.48h6.677a3.76 3.76 0 0 0 3.929-4.158a3.649 3.649 0 0 0-.75-2.09l-.11-.14c-.43-.55-.68-.909-.55-1.289c.13-.38.365-.4.365-.4s.185-.03.455.09c.22.128.46.22.71.27a1.58 1.58 0 0 0 1.736-.905c.095-.208.143-.435.143-.664c.006-.718-.33-1.455-.725-2.088a4.998 4.998 0 0 0-1.554-1.57a3.998 3.998 0 0 0-2.639-.4a3.049 3.049 0 0 0-1.67.89a3.56 3.56 0 0 0-.779 1.359a4.358 4.358 0 0 0-.636-.747v-.159A1.47 1.47 0 0 0 5.558 1.52zm5.304 8.739c.111.741.22 1.821-.867 2.442c-.296.103-.608.16-.923.167H3.215a1 1 0 0 1 .92-1h1.279v-.499a1.79 1.79 0 0 1 1.653-1.825l-.626-.887c-.236.067-.463.153-.577.233H2.655a.754.754 0 0 0-.264.07c-.138.055-.274.109-.396.03c-.2-.13.11-1.12 1.01-1.12h1c.49 0 .57-.54.57-.54l.28-1.129a3.389 3.389 0 0 1-2.85-.93a3.389 3.389 0 0 1 3.14-2.658l.083.002c.26.008.435.014.776.168c.93.42 2.149 2.469 2.149 2.469l.06.09h.17v-.07c-.06-.443-.02-1.464.116-1.89c.137-.424.367-.814.673-1.14a2.349 2.349 0 0 1 1.3-.659a2.639 2.639 0 0 1 1.86.29c.46.284.85.67 1.139 1.127c.289.457.476.836.535 1.374c-.001.02 0 .047.002.081c.007.143.02.39-.128.547c-.127.135-.448.23-.67.18a1.57 1.57 0 0 1-.45-.18a1.33 1.33 0 0 0-1.139-.13a1.42 1.42 0 0 0-.94 1a2.318 2.318 0 0 0 .64 2.238l.11.14c.347.434.546.966.57 1.52a2.999 2.999 0 0 1-.306 1.425a2.708 2.708 0 0 0-.464-1.304zM4.24 5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarEmpty(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.595 6.252L8 1L6.405 6.252H1l4.373 3.4L3.75 15L8 11.695L12.25 15l-1.623-5.348L15 6.252zm-7.247.47H6.72L8 2.507L6.72 6.722zm3.537 2.75l-1.307 4.305zm7.767-2.75H9.28zm-8.75.9h2.366L8 5.214l.732 2.41h2.367l-1.915 1.49l.731 2.409L8 10.032l-1.915 1.49l.731-2.41l-1.915-1.49z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFull(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.595 6.252L8 1L6.405 6.252H1l4.373 3.4L3.75 15L8 11.695L12.25 15l-1.623-5.348L15 6.252z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.405 6.252L8 1l1.595 5.252H15l-4.373 3.4L12.25 15L8 11.695L3.75 15l1.623-5.348L1 6.252zM8 10.032l1.915 1.49l-.731-2.41l1.915-1.49H8.732L8 5.214v4.82zm5.652-3.31H9.28z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircle(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 6h4v4H6z"/><path fill-rule="evenodd" d="M8.6 1c1.6.1 3.1.9 4.2 2c1.3 1.4 2 3.1 2 5.1c0 1.6-.6 3.1-1.6 4.4c-1 1.2-2.4 2.1-4 2.4c-1.6.3-3.2.1-4.6-.7c-1.4-.8-2.5-2-3.1-3.5C.9 9.2.8 7.5 1.3 6c.5-1.6 1.4-2.9 2.8-3.8C5.4 1.3 7 .9 8.6 1m.5 12.9c1.3-.3 2.5-1 3.4-2.1c.8-1.1 1.3-2.4 1.2-3.8c0-1.6-.6-3.2-1.7-4.3c-1-1-2.2-1.6-3.6-1.7c-1.3-.1-2.7.2-3.8 1c-1.1.8-1.9 1.9-2.3 3.3c-.4 1.3-.4 2.7.2 4c.6 1.3 1.5 2.3 2.7 3c1.2.7 2.6.9 3.9.6" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SurroundWith(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4h1V3H1.5l-.5.5v9l.5.5H3v-1H2zm12.5-1H13v1h1v8h-1v1h1.5l.5-.5v-9zM5 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolArray(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.5 2l-.5.5v11l.5.5H4v-1H2V3h2V2zm13 12l.5-.5v-11l-.5-.5H12v1h2v10h-2v1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolBoolean(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1 3.5l.5-.5h13l.5.5v9l-.5.5h-13l-.5-.5zM14 4H8v3.493h-.5l-3.574-.005l2.09-2.09l-.707-.707l-2.955 2.955v.708l2.955 2.955l.707-.707l-2.114-2.114l3.996.005H8v-.986l3.907.005l-2.114-2.114l.707-.707l2.956 2.955v.708L10.5 11.309l-.707-.707l2.09-2.09L8 8.507V12h6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolClass(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.34 9.71h.71l2.67-2.67v-.71L13.38 5h-.7l-1.82 1.81h-5V5.56l1.86-1.85V3l-2-2H5L1 5v.71l2 2h.71l1.14-1.15v5.79l.5.5H10v.52l1.33 1.34h.71l2.67-2.67v-.71L13.37 10h-.7l-1.86 1.85h-5v-4H10v.48zm1.69-3.65l.63.63l-2 2l-.63-.63zm0 5l.63.63l-2 2l-.63-.63zM3.35 6.65l-1.29-1.3l3.29-3.29l1.3 1.29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolColor(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.003a7 7 0 0 0-7 7v.43c.09 1.51 1.91 1.79 3 .7a1.87 1.87 0 0 1 2.64 2.64c-1.1 1.16-.79 3.07.8 3.2h.6a7 7 0 1 0 0-14zm0 13h-.52a.58.58 0 0 1-.36-.14a.56.56 0 0 1-.15-.3a1.24 1.24 0 0 1 .35-1.08a2.87 2.87 0 0 0 0-4a2.87 2.87 0 0 0-4.06 0a1 1 0 0 1-.9.34a.41.41 0 0 1-.22-.12a.42.42 0 0 1-.1-.29v-.37a6 6 0 1 1 6 6zM9 3.997a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 7.007a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-7-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m7-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0M13 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolConstant(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4 6h8v1H4zm8 3H4v1h8z"/><path d="m1 4l1-1h12l1 1v8l-1 1H2l-1-1zm1 0v8h12V4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolEnum(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 2H8L7 3v3h1V3h6v5h-4v1h4l1-1V3zM9 6h4v1H9.41L9 6.59zM7 7H2L1 8v5l1 1h6l1-1V8L8 7zm1 6H2V8h6zM3 9h4v1H3zm0 2h4v1H3zm6-7h4v1H9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolEnumMember(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7 3l1-1h6l1 1v5l-1 1h-4V8h4V3H8v3H7zm2 6V8L8 7H2L1 8v5l1 1h6l1-1zM8 8v5H2V8zm1.414-1L9 6.586V6h4v1zM9 4h4v1H9zm-2 6H3v1h4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolEvent(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.414 1.56L8.312 1h3.294l.818 1.575L10.236 6h1.781l.72 1.695L5.618 15l-1.602-1.163L6.119 10H4.898L4 8.56zM7.78 9L4.9 14.305L12.018 7H8.312l3.294-5H8.312L4.898 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolField(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.45 4.5l-5-2.5h-.9l-7 3.5l-.55.89v4.5l.55.9l5 2.5h.9l7-3.5l.55-.9v-4.5zm-8 8.64l-4.5-2.25V7.17l4.5 2zm.5-4.8L2.29 6.23l6.66-3.34l4.67 2.34zm7 1.55l-6.5 3.25V9.21l6.5-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolFile(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.85 4.44l-3.28-3.3l-.35-.14H2.5l-.5.5v13l.5.5h11l.5-.5V4.8zM13 5h-3V2zM3 14V2h6v3.5l.5.5H13v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolInterface(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.496 4a3.49 3.49 0 0 0-3.46 3h-3.1a2 2 0 1 0 0 1h3.1a3.5 3.5 0 1 0 3.46-4m0 6a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolKey(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.223 10.933c.326.192.699.29 1.077.282a2.159 2.159 0 0 0 1.754-.842a3.291 3.291 0 0 0 .654-2.113a2.886 2.886 0 0 0-.576-1.877a1.99 1.99 0 0 0-1.634-.733a2.294 2.294 0 0 0-1.523.567V3.475h-.991V11.1h.995v-.344c.076.066.158.125.244.177M7.85 6.7c.186-.079.388-.113.59-.1a1.08 1.08 0 0 1 .896.428c.257.363.382.802.357 1.245a2.485 2.485 0 0 1-.4 1.484a1.133 1.133 0 0 1-.96.508a1.224 1.224 0 0 1-.976-.417A1.522 1.522 0 0 1 6.975 8.8v-.6a1.722 1.722 0 0 1 .393-1.145c.13-.154.296-.276.482-.355M3.289 5.675a3.03 3.03 0 0 0-.937.162a2.59 2.59 0 0 0-.8.4l-.1.077v1.2l.423-.359a2.1 2.1 0 0 1 1.366-.572a.758.758 0 0 1 .661.282c.15.232.23.503.231.779L2.9 7.825a2.6 2.6 0 0 0-1.378.575a1.65 1.65 0 0 0-.022 2.336a1.737 1.737 0 0 0 1.253.454a1.96 1.96 0 0 0 1.107-.332c.102-.068.197-.145.286-.229v.444h.941V7.715a2.193 2.193 0 0 0-.469-1.5a1.687 1.687 0 0 0-1.329-.54m.857 3.041c.02.418-.12.829-.391 1.148a1.221 1.221 0 0 1-.955.422a.832.832 0 0 1-.608-.2a.833.833 0 0 1 0-1.091c.281-.174.6-.277.93-.3l1.02-.148zm8.313 2.317c.307.13.64.193.973.182c.495.012.983-.114 1.41-.365l.123-.075l.013-.007V9.615l-.446.32c-.316.224-.696.34-1.084.329A1.3 1.3 0 0 1 12.4 9.8a1.975 1.975 0 0 1-.4-1.312a2.01 2.01 0 0 1 .453-1.381A1.432 1.432 0 0 1 13.6 6.6a1.8 1.8 0 0 1 .971.279l.43.265V5.97l-.17-.073a2.9 2.9 0 0 0-1.17-.247a2.52 2.52 0 0 0-1.929.817a2.9 2.9 0 0 0-.747 2.049c-.028.707.21 1.4.67 1.939c.222.249.497.446.804.578" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolKeyword(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4h-5V3h5zm-1 3h-2v1h2zm-4 0H1v1h9zm2 6H1v1h11zm-5-3H1v1h6zm8 0h-5v1h5zM8 2v3H1V2zM7 3H2v1h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolMethod(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.51 4l-5-3h-1l-5 3l-.49.86v6l.49.85l5 3h1l5-3l.49-.85v-6zm-6 9.56l-4.5-2.7V5.7l4.5 2.45zM3.27 4.7l4.74-2.84l4.74 2.84l-4.74 2.59zm9.74 6.16l-4.5 2.7V8.15l4.5-2.45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolMisc(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2h8v4c.341.035.677.112 1 .23V1H3v8.48l1-1.75zm2.14 8L5 8L4 9.75L3.29 11L1 15h8l-2.29-4zm-3.42 4l1.72-3L5 10l.56 1l1.72 3zm6.836-6.41a3.5 3.5 0 1 1 3.888 5.82a3.5 3.5 0 0 1-3.888-5.82m.555 4.989a2.5 2.5 0 1 0 2.778-4.157a2.5 2.5 0 0 0-2.778 4.157" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolNamespace(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 2.984V2h-.09c-.313 0-.616.062-.909.185a2.33 2.33 0 0 0-.775.53a2.23 2.23 0 0 0-.493.753v.001a3.542 3.542 0 0 0-.198.83v.002a6.08 6.08 0 0 0-.024.863c.012.29.018.58.018.869c0 .203-.04.393-.117.572v.001a1.504 1.504 0 0 1-.765.787a1.376 1.376 0 0 1-.558.115H2v.984h.09c.195 0 .38.04.556.121l.001.001c.178.078.329.184.455.318l.002.002c.13.13.233.285.307.465l.001.002c.078.18.117.368.117.566c0 .29-.006.58-.018.869c-.012.296-.004.585.024.87v.001c.033.283.099.558.197.824v.001c.106.273.271.524.494.753c.223.23.482.407.775.53c.293.123.596.185.91.185H6v-.984h-.09c-.199 0-.387-.038-.562-.115a1.613 1.613 0 0 1-.457-.32a1.659 1.659 0 0 1-.309-.467c-.074-.18-.11-.37-.11-.573c0-.228.003-.453.011-.672c.008-.228.008-.45 0-.665a4.639 4.639 0 0 0-.055-.64a2.682 2.682 0 0 0-.168-.609A2.284 2.284 0 0 0 3.522 8a2.284 2.284 0 0 0 .738-.955c.08-.192.135-.393.168-.602c.033-.21.051-.423.055-.64c.008-.22.008-.442 0-.666c-.008-.224-.012-.45-.012-.678a1.47 1.47 0 0 1 .877-1.354a1.33 1.33 0 0 1 .563-.121zm4 10.032V14h.09c.313 0 .616-.062.909-.185c.293-.123.552-.3.775-.53c.223-.23.388-.48.493-.753v-.001c.1-.266.165-.543.198-.83v-.002c.028-.28.036-.567.024-.863c-.012-.29-.018-.58-.018-.869c0-.203.04-.393.117-.572v-.001a1.504 1.504 0 0 1 .765-.787c.176-.077.362-.115.558-.115H14v-.984h-.09c-.195 0-.38-.04-.556-.121l-.001-.001a1.376 1.376 0 0 1-.455-.318l-.002-.002a1.414 1.414 0 0 1-.307-.465l-.001-.002a1.405 1.405 0 0 1-.117-.566c0-.29.006-.58.018-.869a6.19 6.19 0 0 0-.024-.87v-.001a3.542 3.542 0 0 0-.197-.824v-.001a2.23 2.23 0 0 0-.494-.753a2.33 2.33 0 0 0-.775-.53a2.325 2.325 0 0 0-.91-.185H10v.984h.09c.2 0 .386.038.562.115c.174.082.326.188.457.32c.127.134.23.29.309.467c.074.18.11.37.11.573c0 .228-.003.452-.011.672c-.008.228-.008.45 0 .665c.004.222.022.435.055.64c.033.214.089.416.168.609a2.282 2.282 0 0 0 .738.955a2.282 2.282 0 0 0-.738.955a2.7 2.7 0 0 0-.168.602c-.033.21-.051.423-.055.64c-.008.22-.008.442 0 .666c.008.224.012.45.012.678a1.47 1.47 0 0 1-.42 1.035a1.466 1.466 0 0 1-.457.319a1.33 1.33 0 0 1-.563.121z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolNumeric(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 1v4h4v1h-4v4h4v1h-4v4h-1v-4H6v4H5v-4H1v-1h4V6H1V5h4V1h1v4h4V1zM6 6v4h4V6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolOperator(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.873 1.1c.335.136.602.398.745.73c.072.17.109.352.107.537a1.34 1.34 0 0 1-.61 1.135a1.359 1.359 0 0 1-.753.223A1.355 1.355 0 0 1 1 2.362a1.355 1.355 0 0 1 .83-1.256A1.37 1.37 0 0 1 2.873 1.1m-.298 1.765a.551.551 0 0 0 .332-.5a.548.548 0 1 0-.332.5M6.43 1.109L1.11 6.43l.686.687l5.32-5.32zM11.5 9h1v2.5H15v1h-2.5V15h-1v-2.5H9v-1h2.5zm-5.732.525l.707.707L4.707 12l1.768 1.768l-.707.707L4 12.707l-1.768 1.768l-.707-.707L3.293 12l-1.768-1.768l.707-.707L4 11.293zm1.35-4.195a1.353 1.353 0 0 0-1.256-.83a1.355 1.355 0 0 0-1.256.83a1.362 1.362 0 0 0 1.257 1.895A1.358 1.358 0 0 0 7.118 5.33m-.753.745a.553.553 0 0 1-.289.29a.547.547 0 0 1-.599-.117a.529.529 0 0 1-.117-.173a.544.544 0 0 1 .716-.715a.565.565 0 0 1 .173.116a.549.549 0 0 1 .116.599M14 3h-4v1h4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolParameter(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 6h-1v-.5a.5.5 0 0 0-.5-.5H8.479v5.5a.5.5 0 0 0 .5.5h.5v1h-3v-1h.5a.5.5 0 0 0 .5-.5V5H6.5a.5.5 0 0 0-.5.5V6H5V4h6zm2.914 2.048l-1.462-1.462l.707-.707l1.816 1.816v.707l-1.768 1.767l-.707-.707zM3.548 9.462L2.086 8L3.5 6.586l-.707-.707l-1.768 1.767v.708l1.816 1.815z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolProperty(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.807 14.975a1.75 1.75 0 0 1-1.255-.556a1.684 1.684 0 0 1-.544-1.1A1.72 1.72 0 0 1 1.36 12.1c1.208-1.27 3.587-3.65 5.318-5.345a4.257 4.257 0 0 1 .048-3.078a4.095 4.095 0 0 1 1.665-1.969a4.259 4.259 0 0 1 4.04-.36l.617.268l-2.866 2.951l1.255 1.259l2.944-2.877l.267.619a4.295 4.295 0 0 1 .04 3.311a4.198 4.198 0 0 1-.923 1.392a4.27 4.27 0 0 1-.743.581a4.217 4.217 0 0 1-3.812.446c-1.098 1.112-3.84 3.872-5.32 5.254a1.63 1.63 0 0 1-1.084.423zm7.938-13.047a3.32 3.32 0 0 0-1.849.557c-.213.13-.412.284-.591.458a3.321 3.321 0 0 0-.657 3.733l.135.297l-.233.227c-1.738 1.697-4.269 4.22-5.485 5.504a.805.805 0 0 0 .132 1.05a.911.911 0 0 0 .298.22c.1.044.209.069.319.072a.694.694 0 0 0 .45-.181c1.573-1.469 4.612-4.539 5.504-5.44l.23-.232l.294.135a3.286 3.286 0 0 0 3.225-.254a3.33 3.33 0 0 0 .591-.464a3.28 3.28 0 0 0 .964-2.358c0-.215-.021-.43-.064-.642L11.43 7.125L8.879 4.578l2.515-2.59a3.286 3.286 0 0 0-.65-.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolRuler(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 1L3 2v12l1 1h8l1-1V2l-1-1zm0 2V2h8v12H4v-1h2v-1H4v-2h4V9H4V7h2V6H4V4h4V3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolSnippet(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2.5 1l-.5.5V13h1V2h11v11h1V1.5l-.5-.5zM2 15v-1h1v1zm3-1H4v1h1zm1 0h1v1H6zm3 0H8v1h1zm1 0h1v1h-1zm5 1v-1h-1v1zm-3-1h1v1h-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolString(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2L1 3v9l1 1h12l1-1V3l-1-1zm0 10V3h12v9zm3.356-3.07H6V7.22C6 6.408 5.685 6 5.056 6c-.135 0-.285.024-.45.073a1.444 1.444 0 0 0-.388.167v.665c.237-.203.487-.304.75-.304c.261 0 .392.156.392.469l-.6.103c-.506.086-.76.406-.76.961c0 .263.061.473.183.631A.61.61 0 0 0 4.69 9c.29 0 .509-.16.657-.48h.009zm.004-1.355v.193a.75.75 0 0 1-.12.436a.368.368 0 0 1-.313.17a.276.276 0 0 1-.22-.095a.38.38 0 0 1-.08-.248c0-.222.11-.351.332-.389l.4-.067zM7.6 8.626h-.007v.31H7V5h.593v1.677h.008c.146-.31.355-.465.625-.465c.248 0 .44.118.573.353c.134.236.201.557.201.966c0 .443-.078.798-.235 1.067C8.61 8.866 8.4 9 8.138 9c-.237 0-.416-.125-.537-.374zm-.016-1.121v.272a.78.78 0 0 0 .107.426c.071.113.163.169.274.169c.135 0 .24-.072.314-.216c.075-.145.113-.35.113-.615c0-.22-.035-.39-.104-.514c-.067-.124-.164-.187-.29-.187c-.12 0-.219.062-.298.185a.887.887 0 0 0-.116.48M11.262 9c.321 0 .567-.058.738-.173v-.71a.9.9 0 0 1-.552.207a.619.619 0 0 1-.5-.215c-.12-.145-.181-.345-.181-.598c0-.26.063-.464.189-.612a.644.644 0 0 1 .516-.223c.194 0 .37.069.528.207v-.749c-.129-.09-.338-.134-.626-.134c-.417 0-.751.14-1.001.422c-.249.28-.373.662-.373 1.148c0 .42.116.764.349 1.03c.232.267.537.4.913.4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolStructure(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 2L1 3v3l1 1h12l1-1V3l-1-1zm0 1h12v3H2zm-1 7l1-1h3l1 1v3l-1 1H2l-1-1zm2 0H2v3h3v-3zm7 0l1-1h3l1 1v3l-1 1h-3l-1-1zm2 0h-1v3h3v-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymbolVariable(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 5h2V4H1.5l-.5.5v8l.5.5H4v-1H2zm12.5-1H12v1h2v7h-2v1h2.5l.5-.5v-8zm-2.74 2.57L12 7v2.51l-.3.45l-4.5 2h-.46l-2.5-1.5l-.24-.43v-2.5l.3-.46l4.5-2h.46zM5 9.71l1.5.9V9.28L5 8.38zm.58-2.15l1.45.87l3.39-1.5l-1.45-.87zm1.95 3.17l3.5-1.56v-1.4l-3.5 1.55z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sync(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.006 8.267L.78 9.5L0 8.73l2.09-2.07l.76.01l2.09 2.12l-.76.76l-1.167-1.18a5 5 0 0 0 9.4 1.983l.813.597a6 6 0 0 1-11.22-2.683m10.99-.466L11.76 6.55l-.76.76l2.09 2.11l.76.01l2.09-2.07l-.75-.76l-1.194 1.18a6 6 0 0 0-11.11-2.92l.81.594a5 5 0 0 1 9.3 2.346z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SyncIgnored(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="m5.468 3.687l-.757-.706a6 6 0 0 1 9.285 4.799L15.19 6.6l.75.76l-2.09 2.07l-.76-.01L11 7.31l.76-.76l1.236 1.25a5 5 0 0 0-7.528-4.113m4.55 8.889l.784.73a6 6 0 0 1-8.796-5.04L.78 9.5L0 8.73l2.09-2.07l.76.01l2.09 2.12l-.76.76l-1.167-1.18a5 5 0 0 0 7.005 4.206" clip-rule="evenodd"/><path d="m1.123 2.949l.682-.732L13.72 13.328l-.682.732z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 2h-12l-.5.5v11l.5.5h12l.5-.5v-11zM2 3h11v1H2zm7 4H6V5h3zm0 1v2H6V8zM2 5h3v2H2zm0 3h3v2H2zm0 5v-2h3v2zm4 0v-2h3v2zm7 0h-3v-2h3zm0-3h-3V8h3zm-3-3V5h3v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.2 2H8.017l-.353.146L1 8.81v.707L6.183 14.7h.707l2.215-2.215A4.48 4.48 0 0 0 15.65 9c.027-.166.044-.332.051-.5a4.505 4.505 0 0 0-2-3.74V2.5l-.5-.5zm-.5 2.259A4.504 4.504 0 0 0 11.2 4a.5.5 0 1 0 0 1a3.5 3.5 0 0 1 1.5.338v2.138L8.775 11.4a.506.506 0 0 0-.217.217l-2.022 2.022l-4.475-4.476L8.224 3H12.7zm1 1.792a3.5 3.5 0 0 1 1 2.449a3.438 3.438 0 0 1-.051.5a3.487 3.487 0 0 1-4.793 2.735l3.698-3.698l.146-.354z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Target(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-4 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0m-7 6A6 6 0 1 0 8 2a6 6 0 0 0 0 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tasklist(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.57 6.699l5.693-4.936L8.585 1L3.273 5.596l-1.51-1.832L1 4.442l1.85 2.214zM15 5H6.824l2.307-2H15zM6 7h9v2H6zm9 4H6v2h9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telescope(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.24 1l.59.24l2.11 4.93l-.23.59l-3.29 1.41l-.59-.24l-.17-.41L6.1 9l-.58-.19l-.16-.38L2.8 9.49l-.58-.24l-.72-1.67l.28-.59l2.5-1.06l-.18-.41l.24-.58L7.9 3.41L7.72 3L8 2.42zM2.5 7.64l.35.85l2.22-.91l-.37-.85zm2.74-2.12l1.11 2.45l3-1.28l-1.11-2.44zM8.79 3l1.86 4.11l2.29-1.01L11.18 2L8.72 3zM8.5 9.1l3.02 4.9h-1.17l-1.88-3.03v4h-1V9.82L5.58 14h-1.1l1.7-3.9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 3L3 1.5h18L22.5 3v18L21 22.5H3L1.5 21zM3 3v18h18V3z" clip-rule="evenodd"/><path d="M7.06 7.5L6 8.56l4.243 4.243L6 17.046l1.06 1.06L12 13.168v-.728zm4.94 9h6V18h-6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalBash(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.655 3.56L8.918.75a1.785 1.785 0 0 0-1.82 0L2.363 3.56a1.889 1.889 0 0 0-.921 1.628v5.624a1.889 1.889 0 0 0 .913 1.627l4.736 2.812a1.785 1.785 0 0 0 1.82 0l4.736-2.812a1.888 1.888 0 0 0 .913-1.627V5.188a1.889 1.889 0 0 0-.904-1.627zm-3.669 8.781v.404a.149.149 0 0 1-.07.124l-.239.137c-.038.02-.07 0-.07-.053v-.396a.78.78 0 0 1-.545.053a.073.073 0 0 1-.027-.09l.086-.365a.153.153 0 0 1 .071-.096a.048.048 0 0 1 .038 0a.662.662 0 0 0 .497-.063a.662.662 0 0 0 .37-.567c0-.206-.112-.292-.384-.293c-.344 0-.661-.066-.67-.574A1.47 1.47 0 0 1 9.6 9.437V9.03a.147.147 0 0 1 .07-.126l.231-.147c.038-.02.07 0 .07.054v.409a.754.754 0 0 1 .453-.055a.073.073 0 0 1 .03.095l-.081.362a.156.156 0 0 1-.065.09a.055.055 0 0 1-.035 0a.6.6 0 0 0-.436.072a.549.549 0 0 0-.331.486c0 .185.098.242.425.248c.438 0 .627.199.632.639a1.568 1.568 0 0 1-.576 1.185zm2.481-.68a.094.094 0 0 1-.036.092l-1.198.727a.034.034 0 0 1-.04.003a.035.035 0 0 1-.016-.037v-.31a.086.086 0 0 1 .055-.076l1.179-.706a.035.035 0 0 1 .056.035v.273zm.827-6.914L8.812 7.515c-.559.331-.97.693-.97 1.367v5.52c0 .404.165.662.413.741a1.465 1.465 0 0 1-.248.025c-.264 0-.522-.072-.748-.207L2.522 12.15a1.558 1.558 0 0 1-.75-1.338V5.188a1.558 1.558 0 0 1 .75-1.34l4.738-2.81a1.46 1.46 0 0 1 1.489 0l4.736 2.812a1.548 1.548 0 0 1 .728 1.083c-.154-.334-.508-.427-.92-.185h.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalCmd(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m10.875 7l2.008 5h-.711l-2.008-5zm-5.125.594c-.276 0-.526.041-.75.125a1.542 1.542 0 0 0-.578.375c-.162.166-.287.37-.375.61a2.364 2.364 0 0 0-.133.827c0 .287.04.547.117.781c.078.235.196.433.352.594c.156.162.346.29.57.383c.224.094.48.138.766.133a2.63 2.63 0 0 0 .992-.195l.125.484a1.998 1.998 0 0 1-.492.148a4.381 4.381 0 0 1-.75.07a2.61 2.61 0 0 1-.914-.156a2.207 2.207 0 0 1-.742-.453a1.878 1.878 0 0 1-.485-.742a3.204 3.204 0 0 1-.18-1.023c0-.365.06-.698.18-1c.12-.302.287-.563.5-.782c.214-.218.471-.388.774-.507a2.69 2.69 0 0 1 1-.18c.296 0 .536.023.718.07c.183.047.315.094.399.14l-.149.493a1.85 1.85 0 0 0-.406-.14a2.386 2.386 0 0 0-.539-.055M8 8h1v1H8zm0 2h1v1H8z"/><path d="M15.5 1H.5l-.5.5v13l.5.5h15l.5-.5v-13zM15 14H1V5h14zm0-10H1V2h14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalDebian(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.084.029a1.276 1.276 0 0 0-.355.05L6.622.065a9.46 9.46 0 0 1 .514-.048c.075-.005.15-.01.224-.017a1.67 1.67 0 0 1-.276.029m4.127 7.646c.094-.238.172-.436.16-.762l-.133.282c.135-.41.123-.847.112-1.262c-.005-.187-.01-.37-.002-.543l-.054-.015c-.048-1.411-1.268-2.911-2.354-3.419c-.936-.432-2.376-.506-3.042-.18a.657.657 0 0 1 .212-.085c.107-.031.197-.058.135-.093c-.6.06-.778.171-.973.294a1.92 1.92 0 0 1-.635.273c-.11.106.051.063.181.029c.129-.035.226-.06-.004.076a1.7 1.7 0 0 1-.303.05c-.26.025-.492.048-.96.532c.026.041.11-.009.168-.044c.072-.043.106-.063-.054.137C3.07 2.871 1.78 4.31 1.507 4.787l.143.025c-.1.25-.213.461-.313.649c-.136.254-.249.464-.273.667a16.97 16.97 0 0 1-.062.635C.907 7.619.79 8.679 1.12 9.06l-.04.406l.052.11c.036.079.071.157.12.23l-.093.008c.22.692.338.704.473.717c.137.013.291.028.585.757c-.084-.028-.17-.06-.293-.226c-.015.127.18.508.41.806l-.097.112a.89.89 0 0 0 .27.311c.023.019.045.036.066.055c-.372-.203.1.428.371.79c.078.104.14.186.159.218l.073-.132c-.01.19.136.433.41.772l.229-.009c.094.186.438.522.647.538l-.139.181c.254.08.321.135.397.195c.08.064.17.136.502.253l-.13-.23c.108.095.192.186.273.272c.162.176.31.335.62.481c.352.123.536.152.74.184c.168.026.35.055.649.14a33.82 33.82 0 0 0-.217-.005c-.506-.012-1.056-.025-1.443-.163c-3.016-.817-5.776-4.356-5.574-8c-.02-.311-.01-.655 0-.961c.012-.422.022-.776-.049-.882l.032-.105c.166-.54.365-1.191.742-1.957L.861 3.92v-.002v.001c.012.012.106.107.275-.18c.04-.09.079-.182.117-.276c.08-.19.16-.383.264-.56l.08-.02c.054-.315.533-.744.93-1.1c.19-.171.362-.326.46-.443l.02.138C3.541.977 4.414.611 5.074.334c.152-.063.291-.122.414-.176c-.107.118.067.082.311.032c.15-.03.325-.067.478-.076c-.04.023-.082.044-.122.065c-.085.045-.17.088-.25.145c.26-.062.373-.044.499-.024c.109.018.227.036.456.006c-.174.025-.384.094-.35.12c.245.029.398-.002.537-.03c.174-.034.327-.065.61.03L7.625.275c.235.085.409.137.564.183c.313.094.55.165 1.067.439a.58.58 0 0 0 .23-.037c.112-.035.218-.069.477.037c.014.025.022.046.03.066c.03.08.054.143.456.383c.056-.022-.097-.162-.22-.274l-.003-.004c1.01.54 2.108 1.692 2.443 2.924c-.188-.347-.162-.171-.134.015c.018.124.037.253-.006.235c.14.377.255.766.325 1.168l-.023-.085c-.102-.368-.3-1.081-.626-1.555c-.012.137-.092.122-.165.108c-.105-.019-.196-.036-.058.393c.081.119.096.074.109.034c.015-.047.027-.086.147.164c.002.133.034.266.07.414c.022.094.046.195.065.306c-.034-.006-.07-.07-.106-.13c-.045-.076-.087-.147-.117-.101c.076.358.201.545.25.572c-.009.02-.021.02-.034.021c-.027.002-.056.003-.059.167c.022.428.102.39.166.361c.02-.009.037-.017.051-.01a1.724 1.724 0 0 1-.083.245c-.086.221-.188.48-.106.816a2.356 2.356 0 0 0-.106-.295a5.896 5.896 0 0 1-.046-.117c-.018.151-.01.256-.003.355c.013.166.023.312-.094.62c.135-.442.12-.841-.007-.649c.03.343-.12.642-.254.908c-.111.222-.211.42-.184.602l-.161-.222c-.238.344-.22.417-.202.489c.015.06.03.12-.105.339c.051-.09.041-.112.031-.133c-.01-.024-.021-.046.053-.158c-.05.003-.17.12-.316.265c-.123.121-.265.261-.402.368c-1.172.94-2.571 1.062-3.926.556c.006-.031-.006-.066-.097-.128c-1.148-.88-1.827-1.628-1.591-3.36c.068-.051.117-.193.175-.362c.09-.263.203-.59.448-.745c.245-.541.979-1.04 1.764-1.052c.8-.044 1.476.427 1.816.872c-.618-.576-1.63-.751-2.493-.324c-.882.396-1.405 1.368-1.329 2.336c.01-.016.021-.023.03-.03c.02-.015.037-.027.048-.108c-.027 1.88 2.026 3.258 3.504 2.563l.018.039c.397-.109.497-.205.633-.335c.07-.067.148-.142.28-.233a.441.441 0 0 1-.075.085c-.068.067-.143.14-.05.142c.166-.043.634-.465.947-.746l.133-.119c.062-.134.051-.177.04-.221c-.012-.052-.025-.104.076-.3l.229-.114c.03-.088.062-.168.092-.243M6.612 10.06a.018.018 0 0 0-.005.016a.114.114 0 0 0 .005-.016m-.005.016c.008.069.269.268.465.369c.516.19 1.1.198 1.559.181c-.993.415-2.889-.422-3.509-1.532c.057.012.168.14.303.297c.204.234.462.532.678.605c-.213-.17-.377-.387-.53-.61c.288.33.637.6 1.019.779a.102.102 0 0 1 .01-.077zM6.752.219a6.612 6.612 0 0 1-.075-.013c.472.014.437.045.283.08c.018-.029-.09-.047-.208-.067M9.63 6.732c.032-.477-.094-.326-.136-.144c.019.01.036.059.052.107c.028.08.054.158.084.037m-.211.664a1.68 1.68 0 0 1-.314.703c.006-.061-.038-.074-.083-.086c-.092-.026-.183-.052.176-.504a1.113 1.113 0 0 1-.126.242c-.112.184-.21.344.126.133l.033-.06a1.43 1.43 0 0 0 .188-.428m-1.34 1.247c-.347-.053-.662-.186-.397-.19c.221.02.44.02.656-.033a3.544 3.544 0 0 1-.26.223zM6.958.285l-.1.02l.094-.008zM4.79 8.818l-.038.186c.047.064.092.13.136.195c.12.175.237.348.4.483a4.73 4.73 0 0 0-.214-.368c-.08-.13-.169-.272-.285-.496zm.226-.319c.052.108.104.213.185.302l.082.24l-.038-.063c-.1-.166-.2-.333-.252-.524zm7.474-1.282l-.039.098a4.717 4.717 0 0 1-.462 1.474c.261-.49.43-1.028.501-1.572M.438 3.448c.008.037.043.028.075.02c.06-.015.114-.03-.004.236c-.074.052-.119.087-.144.106l-.027.02a.05.05 0 0 1 .008-.017a.597.597 0 0 0 .092-.365M.118 4.76a2.92 2.92 0 0 1-.106.436a.588.588 0 0 0-.005-.154c-.013-.105-.025-.197.135-.402a4.009 4.009 0 0 0-.023.12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalLinux(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.281 11.156a.84.84 0 0 1 .375.297c.084.125.143.276.18.453c.02.104.044.2.07.29a1.772 1.772 0 0 0 .219.476c.047.073.11.153.188.242c.067.073.127.167.18.281a.793.793 0 0 1 .077.328a.49.49 0 0 1-.093.305a.944.944 0 0 1-.235.219c-.12.083-.245.156-.375.219c-.13.062-.26.127-.39.195a3.624 3.624 0 0 0-.555.328c-.156.115-.313.26-.469.438a2.815 2.815 0 0 1-.625.523a1.471 1.471 0 0 1-.383.172c-.13.036-.26.06-.39.07c-.302 0-.552-.052-.75-.156c-.198-.104-.37-.294-.516-.57c-.042-.079-.083-.128-.125-.149a.774.774 0 0 0-.203-.055L8.67 15c-.26-.02-.525-.031-.796-.031a4.28 4.28 0 0 0-.672.054c-.229.037-.456.081-.68.133c-.046.01-.093.05-.14.117a1.7 1.7 0 0 1-.196.227a1.106 1.106 0 0 1-.335.219a1.475 1.475 0 0 1-.555.101c-.172 0-.357-.018-.555-.054a1.82 1.82 0 0 1-.531-.18a3.578 3.578 0 0 0-.953-.328c-.313-.057-.643-.11-.992-.156a3.392 3.392 0 0 1-.344-.063a.774.774 0 0 1-.29-.133a.705.705 0 0 1-.194-.219a.78.78 0 0 1-.079-.351c0-.162.021-.318.063-.469c.042-.15.065-.31.07-.476c0-.115-.008-.227-.023-.336a3.53 3.53 0 0 1-.032-.352c0-.265.063-.46.188-.586c.125-.125.307-.224.547-.297a.99.99 0 0 0 .297-.148a2.27 2.27 0 0 0 .234-.203a1.86 1.86 0 0 0 .203-.242c.063-.089.133-.178.211-.266a.114.114 0 0 0 .024-.07c0-.063-.003-.123-.008-.18l-.016-.188c0-.354.055-.71.164-1.07c.11-.36.253-.71.43-1.055a9.08 9.08 0 0 1 .594-.992c.218-.317.435-.612.648-.883a4.35 4.35 0 0 0 .68-1.203c.15-.416.229-.87.234-1.36c0-.207-.01-.413-.031-.616a6.122 6.122 0 0 1-.031-.625c0-.417.047-.792.14-1.125c.094-.334.24-.62.438-.86s.456-.419.773-.539C7.474.075 7.854.01 8.296 0c.527 0 .946.104 1.259.313c.312.208.552.481.718.82c.167.338.274.716.32 1.133c.048.416.074.838.079 1.265v.133c0 .214.002.404.008.57a2.527 2.527 0 0 0 .226.977c.073.161.182.336.328.523c.25.329.506.66.766.993c.26.333.497.677.71 1.03c.214.355.389.725.524 1.11c.136.386.206.802.211 1.25a3.3 3.3 0 0 1-.164 1.04zm-6.554-8.14c.072 0 .132.018.18.054a.357.357 0 0 1 .109.149a.85.85 0 0 1 .054.187c.01.063.016.128.016.196a.282.282 0 0 1-.024.125a.27.27 0 0 1-.07.086l-.094.078a.796.796 0 0 0-.093.093a.428.428 0 0 1-.149.141a2.129 2.129 0 0 0-.18.117a1.31 1.31 0 0 0-.156.133a.264.264 0 0 0-.07.195c0 .047.023.086.07.117a.704.704 0 0 1 .266.305c.052.12.11.237.172.352c.062.114.143.21.242.289c.099.078.253.117.46.117h.048c.208-.01.406-.065.594-.164c.187-.099.375-.203.562-.313a.633.633 0 0 1 .102-.046a.37.37 0 0 0 .101-.055l.57-.445a.926.926 0 0 0 .024-.102a2.75 2.75 0 0 0 .016-.11a.236.236 0 0 0-.04-.14a.4.4 0 0 0-.093-.094a.34.34 0 0 0-.133-.054a.909.909 0 0 1-.14-.04a1.083 1.083 0 0 1-.352-.14a1.457 1.457 0 0 0-.344-.156c-.02-.006-.036-.021-.047-.047a.983.983 0 0 1-.031-.094a.23.23 0 0 1-.008-.102a.126.126 0 0 0-.008-.078c0-.062.005-.127.016-.195a.551.551 0 0 1 .07-.195a.417.417 0 0 1 .125-.14a.411.411 0 0 1 .203-.056c.162 0 .279.06.352.18c.073.12.112.25.117.39a.397.397 0 0 1-.039.18a.379.379 0 0 0-.04.172c0 .042.014.07.04.086a.26.26 0 0 0 .102.031c.12 0 .197-.028.234-.085a.533.533 0 0 0 .062-.258c0-.12-.01-.253-.03-.399a1.32 1.32 0 0 0-.126-.406a.969.969 0 0 0-.242-.313a.574.574 0 0 0-.383-.124c-.27 0-.466.067-.586.203c-.12.135-.182.338-.187.609c0 .078.005.156.015.234c.01.079.016.157.016.235c0 .026-.003.039-.008.039a.218.218 0 0 1-.047-.016a4.263 4.263 0 0 1-.093-.039a.774.774 0 0 0-.118-.039a.514.514 0 0 0-.203-.008a1.007 1.007 0 0 1-.125.008c-.073 0-.11-.013-.11-.039c0-.078-.004-.177-.015-.297c-.01-.12-.036-.24-.078-.36a.995.995 0 0 0-.156-.296c-.063-.078-.156-.12-.281-.125a.323.323 0 0 0-.227.086a.905.905 0 0 0-.164.203a.64.64 0 0 0-.086.266a5.4 5.4 0 0 1-.031.25a1.459 1.459 0 0 0 .07.406c.026.083.055.156.086.219c.031.062.068.093.11.093c.025 0 .06-.018.101-.054c.042-.037.063-.07.063-.102c0-.016-.008-.026-.024-.031a.147.147 0 0 0-.047-.008c-.036 0-.068-.018-.094-.055a.468.468 0 0 1-.062-.125a5.144 5.144 0 0 1-.047-.148a.564.564 0 0 1 .055-.398c.047-.084.133-.128.258-.133M5.023 15.18c.125 0 .248-.01.368-.032a.97.97 0 0 0 .336-.125a.614.614 0 0 0 .234-.242a.943.943 0 0 0 .094-.375a.816.816 0 0 0-.047-.273a.963.963 0 0 0-.133-.25a2.763 2.763 0 0 0-.203-.281a2.763 2.763 0 0 1-.203-.282a62.93 62.93 0 0 1-.29-.43c-.093-.14-.187-.288-.28-.445a8.124 8.124 0 0 1-.235-.406a2.646 2.646 0 0 0-.266-.398a1.203 1.203 0 0 0-.218-.211a.469.469 0 0 0-.29-.094a.436.436 0 0 0-.296.11a2.26 2.26 0 0 0-.258.265a3.241 3.241 0 0 1-.297.305c-.11.099-.25.177-.422.234a.744.744 0 0 0-.312.172c-.073.073-.11.185-.11.336c0 .104.008.208.024.312c.015.104.026.209.031.313c0 .14-.02.273-.063.398a1.157 1.157 0 0 0-.062.367c0 .141.05.24.148.297c.1.058.211.097.336.117c.157.027.305.047.446.063c.14.016.278.04.414.07c.135.032.27.065.406.102c.135.036.279.094.43.172c.03.015.078.034.14.054l.211.07c.078.027.151.048.219.063a.741.741 0 0 0 .148.024m2.86-.938c.146 0 .302-.015.469-.047a3.54 3.54 0 0 0 .976-.336a2.59 2.59 0 0 0 .406-.257a.222.222 0 0 0 .032-.047a.305.305 0 0 0 .023-.063v-.008c.031-.114.057-.24.078-.375a8.63 8.63 0 0 0 .055-.414a8.98 8.98 0 0 1 .055-.414c.02-.135.039-.268.054-.398c.021-.14.047-.276.078-.406c.032-.13.073-.253.125-.368a1.03 1.03 0 0 1 .211-.304a1.54 1.54 0 0 1 .344-.25v-.016l-.008-.023a.29.29 0 0 1 .047-.149a1.4 1.4 0 0 1 .117-.164a.582.582 0 0 1 .149-.133a.946.946 0 0 1 .164-.078a9.837 9.837 0 0 0-.102-.375a4.938 4.938 0 0 1-.094-.375a7.126 7.126 0 0 0-.093-.476a2.954 2.954 0 0 0-.11-.36a1.317 1.317 0 0 0-.18-.32c-.077-.104-.174-.23-.288-.375a1.189 1.189 0 0 1-.118-.156a.555.555 0 0 1-.046-.196a2.206 2.206 0 0 0-.047-.203a9.48 9.48 0 0 0-.242-.75a2.91 2.91 0 0 0-.172-.383a3.87 3.87 0 0 0-.172-.289c-.052-.078-.107-.117-.164-.117c-.125 0-.274.05-.446.149c-.171.099-.354.208-.546.328c-.193.12-.38.232-.563.336c-.182.104-.346.153-.492.148a.7.7 0 0 1-.43-.148a2.236 2.236 0 0 1-.36-.344c-.109-.13-.2-.242-.273-.336c-.073-.094-.127-.146-.164-.156c-.041 0-.065.031-.07.093a2.56 2.56 0 0 0-.008.211v.133c0 .032-.005.052-.016.063c-.057.12-.12.237-.187.351c-.068.115-.135.232-.203.352a1.611 1.611 0 0 0-.219.758c0 .078.005.156.016.234c.01.078.036.154.078.227l-.016.03a1.31 1.31 0 0 1-.133.157a1.072 1.072 0 0 0-.132.164a2.796 2.796 0 0 0-.407.93c-.078.333-.12.672-.125 1.015c0 .089.006.178.016.266c.01.089.016.177.016.266a.526.526 0 0 1-.008.086a.525.525 0 0 0-.008.086a.75.75 0 0 1 .313.109c.12.068.25.154.39.258c.14.104.274.224.399.36c.125.135.244.267.359.398c.115.13.198.26.25.39c.052.13.086.237.101.32a.444.444 0 0 1-.125.329a.955.955 0 0 1-.312.203c.089.156.198.289.328.398c.13.11.271.198.422.266c.151.068.315.117.492.148c.177.032.35.047.516.047m3.133 1.11c.109 0 .216-.016.32-.047a1.65 1.65 0 0 0 .445-.203c.136-.089.26-.198.375-.329a3.07 3.07 0 0 1 .977-.75l.258-.117a2.18 2.18 0 0 0 .257-.133a.962.962 0 0 0 .165-.132a.256.256 0 0 0 .078-.188a.295.295 0 0 0-.024-.117a.58.58 0 0 0-.07-.117a5.136 5.136 0 0 1-.203-.305a1.978 1.978 0 0 1-.149-.297l-.125-.312a2.558 2.558 0 0 1-.11-.352a.28.28 0 0 0-.054-.101a.53.53 0 0 0-.46-.235a.533.533 0 0 0-.266.07l-.266.149a7.335 7.335 0 0 1-.281.148a.656.656 0 0 1-.297.07a.411.411 0 0 1-.258-.077a.636.636 0 0 1-.172-.211a2.218 2.218 0 0 1-.117-.258l-.094-.258a1.26 1.26 0 0 1-.14.188a.666.666 0 0 0-.125.203c-.068.156-.11.33-.125.523c-.026.302-.06.596-.102.883a4.7 4.7 0 0 1-.21.86a1.914 1.914 0 0 0-.063.273a2.88 2.88 0 0 0-.032.289c0 .255.079.466.235.633c.156.166.367.25.633.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalPowershell(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2.5 1.007l12.999.17l.43.501l-1.82 12.872l-.57.489l-13-.17l-.43-.502L1.93 1.495zM1.18 13.885l11.998.157l1.68-11.882L2.86 2.003zm5.791-3.49l-.14.991l5 .066l.14-.99zm1.71-2.457l-3.663-2.93l-.692.796l2.636 2.112L3.739 9.95l.465.812z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalTmux(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.5 1h-12l-.5.5v13l.5.5h12l.5-.5v-13zM7 7.5V13H2V2h5zm6 5.5H8V8h5zm0-6H8V2h5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalUbuntu(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.137 3.065c0 1.14-.953 2.065-2.128 2.065c-1.174 0-2.127-.924-2.127-2.065C8.882 1.925 9.835 1 11.01 1c1.175 0 2.127.925 2.127 2.065M4.254 7.6c0 1.14-.952 2.065-2.127 2.065C.952 9.665 0 8.74 0 7.6s.952-2.065 2.127-2.065c1.175 0 2.127.924 2.127 2.065m2.683 5.327c-1.537-.32-2.822-1.273-3.542-2.622a3.14 3.14 0 0 1-1.823.22c.876 2.081 2.726 3.59 4.992 4.062a7.593 7.593 0 0 0 1.51.152a2.93 2.93 0 0 1-.624-1.728L7.382 13a10.96 10.96 0 0 1-.445-.074m5.713.009c0 1.14-.952 2.065-2.127 2.065c-1.175 0-2.127-.925-2.127-2.065s.952-2.065 2.127-2.065c1.175 0 2.127.925 2.127 2.065m.842-.759a6.733 6.733 0 0 0 1.355-2.79a6.605 6.605 0 0 0-1.007-5.15a2.984 2.984 0 0 1-1.198 1.366a4.985 4.985 0 0 1 .495 3.439a5.02 5.02 0 0 1-.66 1.601c.495.388.85.928 1.015 1.534M2.038 4.629a.98.98 0 0 1-.076.004C3.291 2.357 5.895 1.017 8.57 1.219a2.884 2.884 0 0 0-.581 1.214a3.98 3.98 0 0 0-.061.463a5.263 5.263 0 0 0-4.185 2.175a3.23 3.23 0 0 0-.972-.38a3.256 3.256 0 0 0-.643-.067c-.031 0-.06.003-.089.005"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSize(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.36 7L1 13h1.34l.51-1.47h2.26L5.64 13H7L4.65 7zm-.15 3.53l.78-2.14l.78 2.14zM11.82 4h-1.6L7 13h1.56l.75-2.29h3.36l.77 2.29H15zM9.67 9.5l1.18-3.59c.059-.185.1-.376.12-.57c.027.192.064.382.11.57l1.25 3.59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeBars(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 5H2V3h12zm0 4H2V7h12zM2 13h12v-2H2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thumbsdown(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.46 14.11a1.46 1.46 0 0 1-.81-.25a1.38 1.38 0 0 1-.45-1.69L5.17 10H2.38a1.36 1.36 0 0 1-1.16-.61a1.35 1.35 0 0 1-.09-1.32C1.8 6.62 3 4 3.4 2.9A1.38 1.38 0 0 1 4.69 2h8.93A1.4 1.4 0 0 1 15 3.4v3.51a1.38 1.38 0 0 1-1.38 1.38h-1.38L6.4 13.75a1.41 1.41 0 0 1-.94.36M4.69 3a.39.39 0 0 0-.36.25C3.93 4.34 2.86 6.7 2 8.49a.39.39 0 0 0 0 .36a.37.37 0 0 0 .38.15h3.3l.52.68v.46l-1.09 2.44a.37.37 0 0 0 .13.46a.38.38 0 0 0 .48 0l6.06-5.59l.47-.13h1.37a.38.38 0 0 0 .38-.41V3.4a.4.4 0 0 0-.38-.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsdownFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.46 14.11a1.46 1.46 0 0 1-.81-.25a1.38 1.38 0 0 1-.45-1.69L5.17 10H2.38a1.36 1.36 0 0 1-1.16-.61a1.35 1.35 0 0 1-.09-1.32C1.8 6.62 3 4 3.4 2.9A1.38 1.38 0 0 1 4.69 2h8.93A1.4 1.4 0 0 1 15 3.4v3.51a1.38 1.38 0 0 1-1.38 1.38h-1.38L6.4 13.75a1.41 1.41 0 0 1-.94.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thumbsup(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.54 2c.289.001.57.088.81.25a1.38 1.38 0 0 1 .45 1.69l-.97 2.17h2.79a1.36 1.36 0 0 1 1.16.61a1.35 1.35 0 0 1 .09 1.32c-.67 1.45-1.87 4.07-2.27 5.17a1.38 1.38 0 0 1-1.29.9H2.38A1.4 1.4 0 0 1 1 12.71V9.2a1.38 1.38 0 0 1 1.38-1.38h1.38L9.6 2.36a1.41 1.41 0 0 1 .94-.36m.77 11.11a.39.39 0 0 0 .36-.25c.4-1.09 1.47-3.45 2.33-5.24a.39.39 0 0 0 0-.36a.37.37 0 0 0-.38-.15h-3.3l-.52-.68v-.46l1.09-2.44a.37.37 0 0 0-.13-.46a.38.38 0 0 0-.48 0L4.22 8.66l-.47.13H2.38A.38.38 0 0 0 2 9.2v3.51a.4.4 0 0 0 .38.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsupFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.54 2c.289.001.57.088.81.25a1.38 1.38 0 0 1 .45 1.69l-.97 2.17h2.79a1.36 1.36 0 0 1 1.16.61a1.35 1.35 0 0 1 .09 1.32c-.67 1.45-1.87 4.07-2.27 5.17a1.38 1.38 0 0 1-1.29.9H2.38A1.4 1.4 0 0 1 1 12.71V9.2a1.38 1.38 0 0 1 1.38-1.38h1.38L9.6 2.36a1.41 1.41 0 0 1 .94-.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tools(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.773 3.485l-.78-.184l-2.108 2.096l-1.194-1.216l2.056-2.157l-.18-.792a4.42 4.42 0 0 0-1.347-.228a3.64 3.64 0 0 0-1.457.28a3.824 3.824 0 0 0-1.186.84a3.736 3.736 0 0 0-.875 1.265a3.938 3.938 0 0 0 0 2.966a335.341 335.341 0 0 0-6.173 6.234c-.21.275-.31.618-.284.963a1.403 1.403 0 0 0 .464.967c.124.135.272.247.437.328c.17.075.353.118.538.127c.316-.006.619-.126.854-.337c1.548-1.457 4.514-4.45 6.199-6.204c.457.194.948.294 1.444.293a3.736 3.736 0 0 0 2.677-1.133a3.885 3.885 0 0 0 1.111-2.73a4.211 4.211 0 0 0-.196-1.378M2.933 13.928a.31.31 0 0 1-.135.07a.437.437 0 0 1-.149 0a.346.346 0 0 1-.144-.057a.336.336 0 0 1-.114-.11c-.14-.143-.271-.415-.14-.568c1.37-1.457 4.191-4.305 5.955-6.046c.1.132.21.258.328.376c.118.123.245.237.38.341c-1.706 1.75-4.488 4.564-5.98 5.994zm11.118-9.065c.002.765-.296 1.5-.832 2.048a2.861 2.861 0 0 1-4.007 0a2.992 2.992 0 0 1-.635-3.137A2.748 2.748 0 0 1 10.14 2.18a2.76 2.76 0 0 1 1.072-.214h.254L9.649 3.839v.696l1.895 1.886h.66l1.847-1.816zM3.24 6.688h1.531l.705.717l.678-.674l-.665-.678V6.01l.057-1.649l-.22-.437l-2.86-1.882l-.591.066l-.831.849l-.066.599l1.838 2.918l.424.215zm-.945-3.632L4.609 4.58L4.57 5.703H3.494L2.002 3.341zm7.105 6.96l.674-.673l3.106 3.185a1.479 1.479 0 0 1 0 2.039a1.404 1.404 0 0 1-1.549.315a1.31 1.31 0 0 1-.437-.315l-3.142-3.203l.679-.678l3.132 3.194a.402.402 0 0 0 .153.105a.477.477 0 0 0 .359 0a.403.403 0 0 0 .153-.105a.436.436 0 0 0 .1-.153a.525.525 0 0 0 .036-.184a.547.547 0 0 0-.035-.184a.436.436 0 0 0-.1-.153z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 3h3v1h-1v9l-1 1H4l-1-1V4H2V3h3V2a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1zM9 2H6v1h3zM4 13h7V4H4zm2-8H5v7h1zm1 0h1v7H7zm2 0h1v7H9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDown(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5.56L2.413 5h11.194l.393.54L8.373 11h-.827z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleLeft(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.44 2l.56.413v11.194l-.54.393L5 8.373v-.827z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleRight(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.56 14L5 13.587V2.393L5.54 2L11 7.627v.827z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleUp(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14 10.44l-.413.56H2.393L2 10.46L7.627 5h.827z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3.784a5.63 5.63 0 0 1-.65.803a6.058 6.058 0 0 1-.786.68a5.442 5.442 0 0 1 .014.377c0 .574-.061 1.141-.184 1.702a8.467 8.467 0 0 1-.534 1.627a8.444 8.444 0 0 1-1.264 2.04a7.768 7.768 0 0 1-1.72 1.521a7.835 7.835 0 0 1-2.095.95a8.524 8.524 0 0 1-2.379.329a8.178 8.178 0 0 1-2.293-.325A7.921 7.921 0 0 1 1 12.52a5.762 5.762 0 0 0 4.252-1.19a2.842 2.842 0 0 1-2.273-1.19a2.878 2.878 0 0 1-.407-.8c.091.014.181.026.27.035a2.797 2.797 0 0 0 1.022-.089a2.808 2.808 0 0 1-.926-.362a2.942 2.942 0 0 1-.728-.633a2.839 2.839 0 0 1-.65-1.822v-.033c.402.227.837.348 1.306.362a2.943 2.943 0 0 1-.936-1.04a2.955 2.955 0 0 1-.253-.649a2.945 2.945 0 0 1 .007-1.453c.063-.243.161-.474.294-.693c.364.451.77.856 1.216 1.213a8.215 8.215 0 0 0 3.008 1.525a7.965 7.965 0 0 0 1.695.263a2.15 2.15 0 0 1-.058-.325a3.265 3.265 0 0 1-.017-.331c0-.397.075-.77.226-1.118a2.892 2.892 0 0 1 1.528-1.528a2.79 2.79 0 0 1 1.117-.225a2.846 2.846 0 0 1 2.099.909a5.7 5.7 0 0 0 1.818-.698a2.815 2.815 0 0 1-1.258 1.586A5.704 5.704 0 0 0 15 3.785z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypeHierarchy(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 12h-1.793L10 10.293V6.5L9.5 6H8V4h.5l.5-.5v-2L8.5 1h-2l-.5.5v2l.5.5H7v2H5.5l-.5.5v3.793L3.293 12H1.5l-.5.5v2l.5.5h2l.5-.5v-1.793L5.707 11h3.586L11 12.707V14.5l.5.5h2l.5-.5v-2zM7 2h1v1H7zM6 7h3v3H6zm-3 7H2v-1h1zm10 0h-1v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypeHierarchySub(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 11h-1.729L8.438 6H9.5l.5-.5v-4L9.5 1h-4l-.5.5v4l.5.5h1.062l-3.333 5H1.5l-.5.5v3l.5.5h3l.5-.5v-3l-.5-.5h-.068L7.5 6.4l3.068 4.6H10.5l-.5.5v3l.5.5h3l.5-.5v-3zM6 5V2h3v3zm-2 7v2H2v-2zm9 2h-2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypeHierarchySuper(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 1h-3l-.5.5v3l.5.5h.068L7.5 9.6L4.432 5H4.5l.5-.5v-3L4.5 1h-3l-.5.5v3l.5.5h1.729l3.333 5H5.5l-.5.5v4l.5.5h4l.5-.5v-4l-.5-.5H8.438l3.333-5H13.5l.5-.5v-3zM2 4V2h2v2zm7 7v3H6v-3zm4-7h-2V2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unfold(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.53 6.51v-4l-1 1l-.71-.71L7.65 1h.71l1.84 1.83l-.71.7l-1-1v3.98zm0 2.98v4l-1-1l-.71.71L7.65 15h.71l1.84-1.83l-.71-.7l-1 1V9.49zM13.73 4L14 5.02l-3.68 2.93L14 10.98L13.73 12h-4.2v-1h3L9.55 8.57H6.54L3.45 11h3.08v1H2.27L2 10.98l3.68-2.92L2 5.02L2.27 4h4.26v1H3.45l3 2.42h3.01L12.53 5h-3V4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UngroupByRefType(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.9 1L5 3.1l-.8.7L3 2.6V7H2V2.5L.8 3.8l-.7-.7L2.2 1zM3 13.4V9H2v4.4L.8 12.2l-.7.7L2.2 15h.7L5 12.9l-.7-.7zM8.5 7h-2L6 6.5v-2l.5-.5h2l.5.5v2zM7 6h1V5H7zm7.5 1h-3l-.5-.5v-3l.5-.5h3l.5.5v3zM12 6h2V4h-2zm-3.5 6h-2l-.5-.5v-2l.5-.5h2l.5.5v2zM7 11h1v-1H7zm7.5 2h-3l-.5-.5v-3l.5-.5h3l.5.5v3zM12 12h2v-2h-2zm-1-2H9v1h2zm0-5H9v1h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 7V5a3 3 0 0 1 5.83-1h1.044A4.002 4.002 0 0 0 4 5v2H3L2 8v6l1 1h10l1-1V8l-1-1zm6 1h2v6H3V8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unmute(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4.83h2.79L8.15 1l.85.35v13l-.85.33l-3.86-3.85H1.5l-.5-.5v-5zM4.85 10L8 13.14V2.56L4.85 5.68l-.35.15H2v4h2.5zM15 7.83a6.97 6.97 0 0 1-1.578 4.428l-.712-.71A5.975 5.975 0 0 0 14 7.83c0-1.4-.48-2.689-1.284-3.71l.712-.71A6.971 6.971 0 0 1 15 7.83m-2 0a4.978 4.978 0 0 1-1.002 3.004l-.716-.716A3.982 3.982 0 0 0 12 7.83a3.98 3.98 0 0 0-.713-2.28l.716-.716c.626.835.997 1.872.997 2.996m-2 0c0 .574-.16 1.11-.44 1.566l-.739-.738a1.993 1.993 0 0 0 .005-1.647l.739-.739c.276.454.435.988.435 1.558" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unverified(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.67 14.72h.71L10.1 13h2.4l.5-.5v-2.42l1.74-1.72v-.71l-1.71-1.72V3.49l-.5-.49H10.1L8.38 1.29h-.71L6 3H3.53L3 3.5v2.43L1.31 7.65v.71L3 10.08v2.42l.53.5H6zM6.16 12H4V9.87l-.12-.35L2.37 8l1.48-1.51l.15-.35V4h2.16l.36-.14L8 2.35l1.54 1.51l.35.14H12v2.14l.17.35L13.69 8l-1.55 1.52l-.14.35V12H9.89l-.38.15L8 13.66l-1.48-1.52zm1.443-5.859a.962.962 0 0 0-.128.291c-.03.109-.05.215-.062.317l-.005.043h-.895l.003-.051c.018-.326.089-.615.212-.864c.052-.108.117-.214.193-.318c.081-.106.18-.2.294-.28c.119-.084.255-.15.409-.2A1.71 1.71 0 0 1 8.165 5c.28 0 .523.046.726.14c.2.089.366.21.494.363c.127.152.22.326.279.52c.058.194.087.394.087.599c0 .191-.032.371-.098.54c-.064.164-.143.32-.238.466c-.094.143-.197.28-.31.41c-.11.129-.211.252-.304.372a2.47 2.47 0 0 0-.23.34a.653.653 0 0 0-.088.318v.48h-.888v-.539c0-.168.031-.323.094-.464a2.15 2.15 0 0 1 .24-.401c.096-.127.2-.25.308-.368a4.74 4.74 0 0 0 .299-.356c.093-.12.17-.246.228-.377a.984.984 0 0 0 .09-.421a1.04 1.04 0 0 0-.047-.318v-.001a.638.638 0 0 0-.13-.243a.558.558 0 0 0-.216-.158H8.46a.689.689 0 0 0-.294-.059a.643.643 0 0 0-.339.083a.742.742 0 0 0-.223.215zM8.5 11h-.888v-.888H8.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VariableGroup(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.387 11.523a.402.402 0 0 1 .593-.367c.058.031.11.065.157.102c.047.036.088.07.125.101a.177.177 0 0 0 .117.047c.052 0 .12-.04.203-.117c.083-.078.175-.182.273-.313c.1-.13.201-.268.305-.414c.104-.146.2-.294.29-.445l.226-.39c.062-.11.107-.199.133-.266a15.33 15.33 0 0 0-.133-.524a15.384 15.384 0 0 1-.133-.523a3.72 3.72 0 0 0-.133-.422a1.04 1.04 0 0 0-.187-.313a.656.656 0 0 0-.266-.187a1.374 1.374 0 0 0-.375-.07a1.628 1.628 0 0 0-.328.031v-.195L7.69 7a2.345 2.345 0 0 1 .461.734c.052.13.097.263.133.399c.037.135.076.283.117.445c.078-.115.175-.26.29-.438a4.49 4.49 0 0 1 .398-.523c.15-.172.31-.315.476-.43A1.02 1.02 0 0 1 10.089 7c.13 0 .247.034.351.101c.105.068.157.175.157.32c0 .282-.141.423-.422.423a.608.608 0 0 1-.29-.07a.608.608 0 0 0-.288-.071c-.1 0-.203.05-.313.148a2.3 2.3 0 0 0-.312.352a9.5 9.5 0 0 0-.485.734l.446 1.852a1.56 1.56 0 0 0 .093.344a.669.669 0 0 0 .094.171a.184.184 0 0 0 .125.079a.37.37 0 0 0 .211-.086a2.14 2.14 0 0 0 .43-.47c.052-.077.093-.15.125-.218l.187.094a2.025 2.025 0 0 1-.219.367a3.775 3.775 0 0 1-.351.422a3.38 3.38 0 0 1-.406.36c-.141.104-.269.153-.383.148a.397.397 0 0 1-.281-.102a1.491 1.491 0 0 1-.204-.234a1.599 1.599 0 0 1-.132-.36a8.263 8.263 0 0 1-.118-.507a34.16 34.16 0 0 1-.101-.532a2.212 2.212 0 0 0-.11-.414l-.203.375a4.489 4.489 0 0 1-.28.453c-.11.157-.222.316-.337.477a2.46 2.46 0 0 1-.375.422c-.135.12-.265.221-.39.305A.66.66 0 0 1 5.91 12a.539.539 0 0 1-.36-.133a.454.454 0 0 1-.163-.344m6.11.477c.28-.36.496-.748.648-1.164a3.87 3.87 0 0 0 .226-1.32c0-.47-.075-.912-.226-1.329A4.57 4.57 0 0 0 11.495 7h.734a3.77 3.77 0 0 1 .922 2.515c0 .474-.073.917-.218 1.329c-.146.411-.38.796-.704 1.156h-.734zM3.77 12a3.373 3.373 0 0 1-.704-1.149a3.97 3.97 0 0 1-.218-1.336c0-.953.307-1.791.922-2.515h.726a4.132 4.132 0 0 0-.64 1.18a4.205 4.205 0 0 0-.227 1.335A3.929 3.929 0 0 0 4.496 12z"/><path d="M15.5 1H.5l-.5.5v13l.5.5h15l.5-.5v-13zM15 14H1V5h14zm0-10H1V2h14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Verified(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.67 14.72h.71L10.1 13h2.4l.5-.5v-2.42l1.74-1.72v-.71l-1.71-1.72V3.49l-.5-.49H10.1L8.38 1.29h-.71L6 3H3.53L3 3.5v2.43L1.31 7.65v.71L3 10.08v2.42l.53.5H6zM6.16 12H4V9.87l-.12-.35L2.37 8l1.48-1.51l.15-.35V4h2.16l.36-.14L8 2.35l1.54 1.51l.35.14H12v2.14l.17.35L13.69 8l-1.55 1.52l-.14.35V12H9.89l-.38.15L8 13.66l-1.48-1.52zm.57-1.52h.71l3.77-3.77L10.5 6L7.09 9.42L5.71 8.04L5 8.75z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerifiedFilled(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.38 14.72h-.71L6 13H3.53L3 12.5v-2.42L1.31 8.36v-.71L3 5.93V3.5l.53-.5H6l1.67-1.71h.71L10.1 3h2.43l.5.49v2.44l1.71 1.72v.71L13 10.08v2.42l-.5.5h-2.4zm-1.65-4.24h.71l3.77-3.77L10.5 6L7.09 9.42L5.71 8.04L5 8.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Versions(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3L7 4v8l1 1h6l1-1V4l-1-1zm6 9H8V4h6zM5 9V5h1V4H4.5l-.5.5v7l.5.5H6v-1H5zM2 8V6h1V5H1.5l-.5.5v5l.5.5H3v-1H2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vm(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 2h-13l-.5.5v10l.5.5H7v1H4v1h8v-1H9v-1h5.5l.5-.5v-10zM14 12H2V3h12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VmActive(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M1.5 2h13l.5.5v5.503a5.006 5.006 0 0 0-1-.583V3H2v9h5a5 5 0 0 0 1 3H4v-1h3v-1H1.5l-.5-.5v-10z"/><path d="M9.778 8.674a4 4 0 1 1 4.444 6.652a4 4 0 0 1-4.444-6.652m2.13 4.99l2.387-3.182l-.8-.6l-2.077 2.769l-1.301-1.041l-.625.78l1.704 1.364l.713-.09z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VmConnect(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 2h13l.5.5v5.503a5.006 5.006 0 0 0-1-.583V3H2v9h5a5 5 0 0 0 1 3H4v-1h3v-1H1.5l-.5-.5v-10z" clip-rule="evenodd"/><path d="M12 8a4 4 0 1 0 0 8a4 4 0 0 0 0-8m0 7a3 3 0 1 1 0-6.001A3 3 0 0 1 12 15"/><path fill-rule="evenodd" d="m12.133 11.435l1.436 1.436l.431-.431l-1.004-1.005L14 10.431l-.431-.43zm-1.129 1.067L10 11.498l.431-.431l1.436 1.435l-1.436 1.436l-.431-.431z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VmOutline(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 2h13l.5.5v5.503a5.006 5.006 0 0 0-1-.583V3H2v9h5a5 5 0 0 0 1 3H4v-1h3v-1H1.5l-.5-.5v-10z" clip-rule="evenodd"/><path d="M12 8a4 4 0 1 0 0 8a4 4 0 0 0 0-8m0 7a3 3 0 1 1 0-6.001A3 3 0 0 1 12 15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VmRunning(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 2h13l.5.5v5.503a5.006 5.006 0 0 0-1-.583V3H2v9h5a5 5 0 0 0 1 3H4v-1h3v-1H1.5l-.5-.5v-10z" clip-rule="evenodd"/><path d="M12 8c.367 0 .721.047 1.063.14c.34.094.658.23.953.407c.294.177.563.385.808.625c.245.24.455.509.63.808a4.03 4.03 0 0 1 .405 3.082c-.093.342-.229.66-.406.954a4.382 4.382 0 0 1-.625.808c-.24.245-.509.455-.808.63a4.029 4.029 0 0 1-3.082.405a3.784 3.784 0 0 1-.954-.406a4.382 4.382 0 0 1-.808-.625a3.808 3.808 0 0 1-.63-.808a4.027 4.027 0 0 1-.405-3.082c.093-.342.229-.66.406-.954c.177-.294.385-.563.625-.808c.24-.245.509-.455.808-.63A4.028 4.028 0 0 1 12 8m2 3.988L11 10v4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vr(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M4 3h8a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3h-.394a3 3 0 0 1-1.665-.504l-.832-.555a2 2 0 0 0-2.218 0l-.832.555A3 3 0 0 1 4.394 13H4a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3m0 1a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h.394a2 2 0 0 0 1.11-.336l.832-.555a3 3 0 0 1 3.328 0l.832.555a2 2 0 0 0 1.11.336H12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2z" clip-rule="evenodd"/><path d="M0 7h1v3H0zm15 0h1v3h-1zM6.5 8a.5.5 0 0 1 0 1H4a.5.5 0 0 1 0-1zM12 8a.5.5 0 0 1 0 1H9.5a.5.5 0 0 1 0-1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wand(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.38 5h1V4h1V3h-1V2h-1v1h-1v1h1zm8 4h-1v1h-1v1h1v1h1v-1h1v-1h-1zM14 2V1h-1v1h-1v1h1v1h1V3h1V2zm-2.947 2.442a1.49 1.49 0 0 0-2.12 0l-7.49 7.49a1.49 1.49 0 0 0 0 2.12c.59.59 1.54.59 2.12 0l7.49-7.49c.58-.58.58-1.53 0-2.12m-8.2 8.9c-.2.2-.51.2-.71 0c-.2-.2-.2-.51 0-.71l6.46-6.46l.71.71zm7.49-7.49l-.32.32l-.71-.71l.32-.32c.2-.2.51-.2.71 0c.19.2.19.52 0 .71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.56 1h.88l6.54 12.26l-.44.74H1.44L1 13.26zM8 2.28L2.28 13H13.7zM8.625 12v-1h-1.25v1zm-1.25-2V6h1.25v4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.5 9h2V8H8V5.5H7v3z"/><path fill-rule="evenodd" d="M5.5 3.669A4.998 4.998 0 0 0 3 8a4.998 4.998 0 0 0 2.5 4.331V14.5l.5.5h4l.5-.5v-2.169A4.998 4.998 0 0 0 13 8a4.998 4.998 0 0 0-2.5-4.331V1.5L10 1H6l-.5.5zM12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whitespace(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2V1H6.5a3.5 3.5 0 0 0 0 7H8v5H7v1h5v-1h-1V2zM8 7H6.5a2.5 2.5 0 1 1 0-5H8zm2 6H9V2h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WholeWord(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M0 11h1v2h14v-2h1v3H0z" clip-rule="evenodd"/><path d="M6.84 11h-.88v-.86h-.022c-.383.66-.947.989-1.692.989c-.548 0-.977-.145-1.289-.435c-.308-.29-.462-.675-.462-1.155c0-1.028.605-1.626 1.816-1.794l1.649-.23c0-.935-.378-1.403-1.134-1.403c-.662 0-1.26.226-1.794.677v-.902c.541-.344 1.164-.516 1.87-.516c1.292 0 1.938.684 1.938 2.052zm-.88-2.782L4.633 8.4c-.408.058-.716.16-.924.307c-.208.143-.311.399-.311.768c0 .268.095.488.284.66c.194.168.45.253.768.253a1.41 1.41 0 0 0 1.08-.457c.286-.308.43-.696.43-1.165zm3.388 1.987h-.022V11h-.88V2.857h.88v3.61h.021c.434-.73 1.068-1.096 1.902-1.096c.705 0 1.257.247 1.654.741c.401.49.602 1.15.602 1.977c0 .92-.224 1.658-.672 2.213c-.447.551-1.06.827-1.837.827c-.726 0-1.276-.308-1.649-.924m-.022-2.218v.768c0 .455.147.841.44 1.16c.298.315.674.473 1.128.473c.534 0 .951-.204 1.252-.613c.304-.408.456-.975.456-1.702c0-.613-.141-1.092-.424-1.44c-.283-.347-.666-.52-1.15-.52c-.511 0-.923.178-1.235.536c-.311.355-.467.8-.467 1.338"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Window(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 2h-13l-.5.5v11l.5.5h13l.5-.5v-11zM14 13H2V6h12zm0-8H2V3h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WordWrap(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.868 3.449a1.21 1.21 0 0 0-.473-.329c-.274-.111-.623-.15-1.055-.076a3.5 3.5 0 0 0-.71.208c-.082.035-.16.077-.235.125l-.043.03v1.056l.168-.139c.15-.124.326-.225.527-.303c.196-.074.4-.113.604-.113c.188 0 .33.051.431.157c.087.095.137.248.147.456l-.962.144c-.219.03-.41.086-.57.166a1.245 1.245 0 0 0-.398.311c-.103.125-.181.27-.229.426c-.097.33-.093.68.011 1.008a1.096 1.096 0 0 0 .638.67c.155.063.328.093.528.093a1.25 1.25 0 0 0 .978-.441v.345h1.007V4.65c0-.255-.03-.484-.089-.681a1.423 1.423 0 0 0-.275-.52m-.636 1.896v.236c0 .119-.018.231-.055.341a.745.745 0 0 1-.377.447a.694.694 0 0 1-.512.027a.454.454 0 0 1-.156-.094a.389.389 0 0 1-.094-.139a.474.474 0 0 1-.035-.186c0-.077.01-.147.024-.212a.33.33 0 0 1 .078-.141a.436.436 0 0 1 .161-.109a1.3 1.3 0 0 1 .305-.073zm5.051-1.067a2.253 2.253 0 0 0-.244-.656a1.354 1.354 0 0 0-.436-.459a1.165 1.165 0 0 0-.642-.173a1.136 1.136 0 0 0-.69.223a1.33 1.33 0 0 0-.264.266V1H5.09v6.224h.918v-.281c.123.152.287.266.472.328c.098.032.208.047.33.047c.255 0 .483-.06.677-.177c.192-.115.355-.278.486-.486a2.29 2.29 0 0 0 .293-.718a3.87 3.87 0 0 0 .096-.886a3.714 3.714 0 0 0-.078-.773zm-.86.758c0 .232-.02.439-.06.613c-.036.172-.09.315-.159.424a.639.639 0 0 1-.233.237a.582.582 0 0 1-.565.014a.683.683 0 0 1-.21-.183a.925.925 0 0 1-.142-.283A1.187 1.187 0 0 1 6 5.5v-.517c0-.164.02-.314.06-.447c.036-.132.087-.242.156-.336a.668.668 0 0 1 .228-.208a.584.584 0 0 1 .29-.071a.554.554 0 0 1 .496.279c.063.099.108.214.143.354c.031.143.05.306.05.482M2.407 9.9a.913.913 0 0 1 .316-.239c.218-.1.547-.105.766-.018c.104.042.204.1.32.184l.33.26V8.945l-.097-.062a1.932 1.932 0 0 0-.905-.215c-.308 0-.593.057-.846.168c-.25.11-.467.27-.647.475c-.18.21-.318.453-.403.717c-.09.272-.137.57-.137.895c0 .289.043.561.13.808c.086.249.211.471.373.652c.161.185.361.333.597.441c.232.104.493.155.778.155c.233 0 .434-.028.613-.084c.165-.05.322-.123.466-.217l.078-.061v-.889l-.2.095a.4.4 0 0 1-.076.026c-.05.017-.099.035-.128.049c-.036.023-.227.09-.227.09c-.06.024-.14.043-.218.059a.977.977 0 0 1-.599-.057a.827.827 0 0 1-.306-.225a1.088 1.088 0 0 1-.205-.376a1.728 1.728 0 0 1-.076-.529c0-.21.028-.399.083-.56c.054-.158.13-.294.22-.4M14 6h-4V5h4.5l.5.5v6l-.5.5H7.879l2.07 2.071l-.706.707l-2.89-2.889v-.707l2.89-2.89L9.95 9l-2 2H14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WorkspaceTrusted(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.069 0c.262 0 .52.017.76.057a4.1 4.1 0 0 1 .697.154c.228.069.451.155.674.263c.217.103.44.229.663.366c.377.24.748.434 1.126.589a7.537 7.537 0 0 0 2.331.525c.406.029.823.046 1.257.046v4c0 .76-.097 1.48-.291 2.166a8.996 8.996 0 0 1-.789 1.943a10.312 10.312 0 0 1-1.188 1.725a15.091 15.091 0 0 1-1.492 1.532a17.57 17.57 0 0 1-1.703 1.325c-.594.412-1.194.795-1.794 1.143l-.24.143l-.24-.143a27.093 27.093 0 0 1-1.806-1.143a15.58 15.58 0 0 1-1.703-1.325a15.082 15.082 0 0 1-1.491-1.532a10.947 10.947 0 0 1-1.194-1.725a9.753 9.753 0 0 1-.789-1.943A7.897 7.897 0 0 1 .571 6V2c.435 0 .852-.017 1.258-.046a8.16 8.16 0 0 0 1.188-.171c.383-.086.766-.2 1.143-.354A6.563 6.563 0 0 0 5.28.846C5.72.56 6.166.349 6.606.21A4.79 4.79 0 0 1 8.069 0m6.502 2.983a9.566 9.566 0 0 1-2.234-.377a7.96 7.96 0 0 1-2.046-.943A4.263 4.263 0 0 0 9.23 1.16A3.885 3.885 0 0 0 8.074.994a3.99 3.99 0 0 0-1.165.166a3.946 3.946 0 0 0-1.058.503A7.926 7.926 0 0 1 3.8 2.61c-.709.206-1.451.332-2.229.378v3.017c0 .663.086 1.297.258 1.908a8.58 8.58 0 0 0 .72 1.743a9.604 9.604 0 0 0 1.08 1.572c.417.491.862.948 1.342 1.382c.48.435.983.835 1.509 1.206c.531.372 1.063.709 1.594 1.017a22.397 22.397 0 0 0 1.589-1.017a15.389 15.389 0 0 0 1.514-1.206c.48-.434.926-.891 1.343-1.382a9.596 9.596 0 0 0 1.08-1.572a8.258 8.258 0 0 0 .709-1.743a6.814 6.814 0 0 0 .262-1.908z"/><path fill-rule="evenodd" d="m11.797 4.709l-.44-.378l-.406.035l-4.36 5.148l-1.485-2.12l-.4-.068l-.463.331l-.069.4l1.909 2.726l.217.12l.457.028l.234-.102l4.835-5.715z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WorkspaceUnknown(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.067 0c.263 0 .52.017.76.057a4.1 4.1 0 0 1 .697.154c.229.069.452.155.675.263c.217.103.44.229.662.366a7.2 7.2 0 0 0 1.126.589a7.534 7.534 0 0 0 2.332.525c.405.029.822.046 1.257.046v4c0 .76-.097 1.48-.292 2.166a8.996 8.996 0 0 1-.788 1.943a10.306 10.306 0 0 1-1.189 1.725a15.082 15.082 0 0 1-1.491 1.532a17.57 17.57 0 0 1-1.703 1.325c-.594.412-1.194.795-1.794 1.143l-.24.143l-.24-.143a27.088 27.088 0 0 1-1.806-1.143a15.579 15.579 0 0 1-1.703-1.325a15.08 15.08 0 0 1-1.491-1.532a10.948 10.948 0 0 1-1.195-1.725a9.753 9.753 0 0 1-.788-1.943A7.897 7.897 0 0 1 .57 6V2c.434 0 .851-.017 1.257-.046a8.16 8.16 0 0 0 1.189-.171c.383-.086.765-.2 1.143-.354a6.563 6.563 0 0 0 1.12-.583C5.719.56 6.164.349 6.604.21A4.79 4.79 0 0 1 8.067 0m6.503 2.983a9.567 9.567 0 0 1-2.234-.377a7.96 7.96 0 0 1-2.046-.943a4.264 4.264 0 0 0-1.063-.503A3.885 3.885 0 0 0 8.073.994a3.99 3.99 0 0 0-1.166.166a3.946 3.946 0 0 0-1.057.503a7.927 7.927 0 0 1-2.051.948c-.709.206-1.452.332-2.229.378v3.017c0 .663.086 1.297.257 1.908a8.58 8.58 0 0 0 .72 1.743a9.604 9.604 0 0 0 1.08 1.572c.417.491.863.948 1.343 1.382c.48.435.983.835 1.509 1.206c.531.372 1.062.709 1.594 1.017a22.4 22.4 0 0 0 1.588-1.017a15.384 15.384 0 0 0 1.515-1.206c.48-.434.925-.891 1.343-1.382a9.609 9.609 0 0 0 1.08-1.572a8.269 8.269 0 0 0 .708-1.743a6.814 6.814 0 0 0 .263-1.908z"/><path fill-rule="evenodd" d="M9.433 4.72c.171.171.314.377.411.606c.103.228.155.48.149.754a1.6 1.6 0 0 1-.114.64a2.24 2.24 0 0 1-.292.48a2.787 2.787 0 0 1-.354.383a4.52 4.52 0 0 0-.337.32a1.421 1.421 0 0 0-.24.32a.7.7 0 0 0-.086.348v.36l-.131.138h-.715l-.143-.143V8.57c0-.24.04-.45.12-.634c.075-.177.166-.343.28-.486a3.42 3.42 0 0 1 .366-.382c.12-.109.229-.212.332-.32c.097-.103.182-.212.245-.326a.707.707 0 0 0 .086-.354a.966.966 0 0 0-.074-.36a.972.972 0 0 0-.2-.298a.94.94 0 0 0-1.32 0a.88.88 0 0 0-.2.298a.829.829 0 0 0-.075.36L7 6.21h-.715l-.131-.137c0-.263.046-.514.148-.748c.103-.229.24-.435.412-.606c.177-.177.383-.32.611-.417a1.883 1.883 0 0 1 1.503 0c.229.103.434.24.606.417zM7.57 9.646l.143-.143h.714l.143.143v.714l-.143.143h-.714l-.143-.143z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WorkspaceUntrusted(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.067 0c.263 0 .52.017.76.057a4.1 4.1 0 0 1 .697.154c.229.069.452.155.675.263c.217.103.44.229.662.366a7.2 7.2 0 0 0 1.126.589a7.534 7.534 0 0 0 2.332.525c.405.029.822.046 1.257.046v4c0 .76-.097 1.48-.292 2.166a8.996 8.996 0 0 1-.788 1.943a10.306 10.306 0 0 1-1.189 1.725a15.082 15.082 0 0 1-1.491 1.532a17.57 17.57 0 0 1-1.703 1.325c-.594.412-1.194.795-1.794 1.143l-.24.143l-.24-.143a27.088 27.088 0 0 1-1.806-1.143a15.579 15.579 0 0 1-1.703-1.325a15.08 15.08 0 0 1-1.491-1.532a10.948 10.948 0 0 1-1.195-1.725a9.753 9.753 0 0 1-.788-1.943A7.897 7.897 0 0 1 .57 6V2c.434 0 .851-.017 1.257-.046a8.16 8.16 0 0 0 1.189-.171c.383-.086.765-.2 1.143-.354a6.563 6.563 0 0 0 1.12-.583C5.719.56 6.164.349 6.604.21A4.79 4.79 0 0 1 8.067 0m6.503 2.983a9.567 9.567 0 0 1-2.234-.377a7.96 7.96 0 0 1-2.046-.943a4.264 4.264 0 0 0-1.063-.503A3.885 3.885 0 0 0 8.073.994a3.99 3.99 0 0 0-1.166.166a3.946 3.946 0 0 0-1.057.503a7.927 7.927 0 0 1-2.051.948c-.709.206-1.452.332-2.229.378v3.017c0 .663.086 1.297.257 1.908a8.58 8.58 0 0 0 .72 1.743a9.604 9.604 0 0 0 1.08 1.572c.417.491.863.948 1.343 1.382c.48.435.983.835 1.509 1.206c.531.372 1.062.709 1.594 1.017a22.4 22.4 0 0 0 1.588-1.017a15.384 15.384 0 0 0 1.515-1.206c.48-.434.925-.891 1.343-1.382a9.609 9.609 0 0 0 1.08-1.572a8.269 8.269 0 0 0 .708-1.743a6.814 6.814 0 0 0 .263-1.908z"/><path d="m10.787 5.446l-.4-.406h-.206L8.2 7.023L6.216 5.04h-.2l-.406.406v.2l1.983 1.983L5.61 9.61v.206l.406.4h.2l1.983-1.983l1.982 1.983h.206l.4-.4V9.61L8.804 7.63l1.983-1.983v-.2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.027 6.149a5.52 5.52 0 0 1-1.27 3.908l4.26 4.26l-.7.71l-4.26-4.27a5.52 5.52 0 1 1 1.97-4.608m-5.45 4.888a4.51 4.51 0 0 0 3.18-1.32l-.04.02a4.51 4.51 0 0 0 1.36-3.2a4.5 4.5 0 1 0-4.5 4.5m2.44-4v-1h-2v-2h-1v2h-2v1h2v2h1v-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *CodiconIcon {
	return &CodiconIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.027 6.149a5.52 5.52 0 0 1-1.27 3.908l4.26 4.26l-.7.71l-4.26-4.27a5.52 5.52 0 1 1 1.97-4.608m-5.45 4.888a4.51 4.51 0 0 0 3.18-1.32l-.04.02a4.51 4.51 0 0 0 1.36-3.2a4.5 4.5 0 1 0-4.5 4.5m-2.54-4.98h5v1h-5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
