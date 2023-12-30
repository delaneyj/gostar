package guidance

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type GuidanceIcon struct {
	*SVGSVGElement
}

func AccesibleRestroom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 13V6.5h-1.172a3 3 0 0 0-2.906 2.255l-.963 3.764M20 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.207M9 23.5a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11Zm5.35-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccessForHearingLoss(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 9a6.5 6.5 0 0 1 13 0c0 .662-.111 1.32-.328 1.944l-2.85 8.195A3.516 3.516 0 0 1 12 21.5c-.98 0-1.865-.352-2.5-1m6-11.5a3.5 3.5 0 1 0-7 0m15-8.5L18 6m-6 6L.5 23.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccessToLowVision(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m16 13.5l6 10m-12-15c0 2.4 3.546 4 5 4m-5-4v5.25m0-5.25a2 2 0 0 0-2-2H6.5A3.5 3.5 0 0 0 3 10v3m2.76-6.421V14L10 16.5v.5c0 1.5 0 3.5.75 5c0 0 .75 1.5 1.75 1.5m-5.75-6.627C6.56 20.03 5.415 22.853 3 23.5m4.359-19s-1.6-1-1.6-2.25a1.746 1.746 0 1 1 3.495 0C9.254 3.5 7.66 4.5 7.66 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccessibleExit(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 23.5h5V.5h-15V8m8 15.5c-.727 0-1.273-1.071-1.273-1.071c-.545-1.072-.545-1.786-.545-2.858V18.5H8.374M7.5 12l-1.836 3.672M3 14l1.5-3s.8-.5 2.6-.5c2.4 0 3.9 2 3.9 2l-1.875 3.75m-3.461-.578a4 4 0 1 0-2.328 7.656a4 4 0 0 0 2.328-7.656Zm6.141-5.172s1.81-.557 2.135-1.776a1.768 1.768 0 0 0-1.242-2.163a1.75 1.75 0 0 0-2.146 1.25c-.324 1.219.962 2.61.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccessibleExitTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M21 23.5h2.5V.5h-15V4m-8 11.79a4 4 0 1 1 0 7.418m9.5.292c-.727 0-1.273-1.071-1.273-1.071c-.545-1.072-.545-1.786-.545-2.858V18.5H6M5 12l-2 4m3.5.5l2-4s-1.5-2-3.9-2c-1.8 0-2.6.5-2.6.5L.5 14M19 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-.5L12.75 15v-.25l1.908-3.642A3 3 0 0 1 17.315 9.5H18v2.264c0 .758.44 1.442 1.184 1.58A9.955 9.955 0 0 0 21 13.5M15 19c-1.65-.057-3.365.506-4 .895M9.305 10.5s1.81-.557 2.135-1.776a1.768 1.768 0 0 0-1.242-2.163a1.75 1.75 0 0 0-2.146 1.25c-.324 1.219.962 2.61.962 2.61zm8.955-2.972s-1.287-1.38-.963-2.588a1.75 1.75 0 0 1 2.143-1.237a1.746 1.746 0 0 1 1.234 2.142c-.324 1.208-2.124 1.76-2.124 1.76z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccessibleMenRestroom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24M4.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C6.246 3.5 4.65 4.5 4.65 4.5zM19 14V8.5h-1.158a3 3 0 0 0-2.91 2.272L14 14.527M24 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.258m-3.742 7a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm4.35-17s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccessibleWomenRestroom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.5 15.646a4.5 4.5 0 1 1-1.244 4.854M19 14V8.5h-1.158a3 3 0 0 0-2.91 2.272L14 14.527M24 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.258m.608-10s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25zM3.5 24v-5.5h-3v-.25l.072-.15C2.17 14.742 2.5 11.07 2.5 7.352v-.329S4 6.5 5.5 6.5s3 .523 3 .523v.329c0 2.78.184 5.532.945 8.148c.257.884.58 1.752.983 2.6l.072.15v.25h-3V24M5.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.246 3.5 5.65 4.5 5.65 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneMode(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 15.5v-2.25S6 10.5 9.75 9.5V4.211c0-.79.23-1.565.707-2.194c.364-.48.828-1.052 1.293-1.517h.5c.465.465.93 1.037 1.293 1.517c.478.629.707 1.404.707 2.194V9.5c3.75 1 9.25 3.75 9.25 3.75v2.25s-5.5-1-9.25-1c0 0-.5 2.28-.5 6.5c0 0 1.25.75 2.25 1.75v.5S13.75 23 12 23s-4 .25-4 .25v-.5c1-1 2.25-1.75 2.25-1.75c0-4.22-.5-6.5-.5-6.5c-3.75 0-9.25 1-9.25 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneModeOff(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l23 23m-8.999-8.999c3.74.043 8.999.999 8.999.999v-2.25S18 10.5 14.25 9.5V4.211c0-.79-.23-1.565-.707-2.194A15.368 15.368 0 0 0 12.25.5h-.5a15.17 15.17 0 0 0-1.293 1.517c-.478.629-.707 1.404-.707 2.194V9.5l-.196.054m4.33 7.83A46.418 46.418 0 0 0 13.75 21s1.25.75 2.25 1.75v.5S13.75 23 12 23s-4 .25-4 .25v-.5c1-1 2.25-1.75 2.25-1.75c0-4.22-.5-6.5-.5-6.5c-3.75 0-9.25 1-9.25 1v-2.25s3.232-1.616 6.43-2.82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m20.5 23l-2.664-3.148M3.5 23l2.664-3.148M22.912 8A12.027 12.027 0 0 0 15 1.378M1.088 8A12.027 12.027 0 0 1 9 1.378M12 6s.5 3.5 0 7c2.75 1.5 5 4 5 4M6.164 19.852a9 9 0 1 1 11.672 0m-11.672 0A8.964 8.964 0 0 0 12 22a8.965 8.965 0 0 0 5.836-2.148"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmBell(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M21.753 11.48A12.56 12.56 0 0 1 15 20.228V23.5H4V20h9m-3.5-9.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm12.5 1a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm-12.5 6a8.5 8.5 0 1 1 0-17a8.5 8.5 0 0 1 0 17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertOctagon(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 6v8.5m0 1.5v2M7.444 1s1.774.25 4.556.25S16.556 1 16.556 1s1.078 1.431 3.045 3.399C21.57 6.366 23 7.444 23 7.444s-.25 1.774-.25 4.556s.25 4.556.25 4.556s-1.431 1.078-3.399 3.045C17.634 21.57 16.556 23 16.556 23s-1.774-.25-4.556-.25s-4.556.25-4.556.25s-1.078-1.431-3.045-3.399C2.43 17.634 1 16.556 1 16.556s.25-1.774.25-4.556S1 7.444 1 7.444s1.431-1.078 3.399-3.045C6.366 2.43 7.444 1 7.444 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlertTriangle(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 7v8.5m0 1.5v2m11 1.5L22 22s-5.7-.25-10-.25S2 22 2 22l-1-1.5C6 13 11 2 11 2h2s5 11 10 18.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmbulanceEntrance(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 6.5v-3H.5V20h3m11-13.5h4a7.138 7.138 0 0 0 4.184 5.637l.816.363V17a3 3 0 0 1-3 3m-6-13.5V13h9m-3 7a2.5 2.5 0 0 1-5 0m5 0a2.5 2.5 0 0 0-5 0m0 0h-7m-5 0a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0M6 6.5V9H3.5v3H6v2.5h3V12h2.5V9H9V6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmusementPark(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M16.127 20a9.032 9.032 0 0 0 2.559-1.975M8 20.065a9.033 9.033 0 0 1-2.686-2.04M5.5 23.5s5.122 0 5.697-9.667M18.5 23.5s-5.122 0-5.698-9.667m-7.488 4.192a2 2 0 1 1-1.877-3.247m1.877 3.247a2 2 0 0 0-1.877-3.247m0 0A8.993 8.993 0 0 1 3 12c0-.97.153-1.903.437-2.778m0 0a2 2 0 0 0 1.877-3.247M3.437 9.222a2 2 0 1 1 1.877-3.247m0 0a8.99 8.99 0 0 1 4.81-2.78m0 0a2 2 0 0 0 3.752 0m-3.752 0a2 2 0 1 1 3.751 0m0 0a8.991 8.991 0 0 1 4.811 2.78m0 0a2 2 0 1 1 1.877 3.247m-1.877-3.247a2 2 0 0 0 1.877 3.247m0 0c.284.875.437 1.809.437 2.778a8.98 8.98 0 0 1-.437 2.778m0 0a2 2 0 0 0-1.877 3.247m1.877-3.247a2 2 0 1 1-1.877 3.247m-7.489-4.192a1.994 1.994 0 0 0 1.605 0m-1.605 0a2 2 0 1 1 1.605 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aquarium(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 19.5c-4.97 0-9-4.253-9-9.5c0-1.294.245-2.527.689-3.65c0 0-.995-.85-3.189-.85L.75 5S2.5 1.5 7.705 1.959C9.093 1.035 11.238.5 13 .5c3.61 0 5.263 1.353 6.571 3.384c.664 1.03 1.446 1.989 2.441 2.706l.488.351V8.5H14c0 3-.877 5.393-2.5 6.33l-.5-.33s.5-3-.92-3.934c-.968.866-1.58 2.128-1.58 3.934c0 2.761 1.79 5 4 5Zm0 0c1.198-2.207 2.5-4 5-4v.286c-.74 1.025-1.5 2.37-1.5 3.714c0 1.344.76 2.689 1.5 3.714v.286c-2.5 0-3.802-1.793-5-4Zm3.5-16v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arcade(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 15.5v-8m0 0a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm-9.5 11v.219a6 6 0 0 1-1.516 3.986L.5 23.25v.25h23v-.25l-.485-.545a6 6 0 0 1-1.515-3.986V18.5zm3 0v-3h13v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrival(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 20.5h24m-13.5-9.76L3 13.99c-.5-3-2.25-6.5-2.25-6.5L2 6.99s1.123.775 2.5 2.5l11.502-6.03a4.074 4.074 0 0 1 3.707-.025c1.054.525 2.346 1.152 3.291 1.556v.5l-7.85 3.59C14 13.49 12 16.49 12 16.49H9.8s.7-2.75.7-5.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssistiveListeningDevice(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 22L.5 23.5M5 19l-1.5 1.5m4.5-2L5.5 16m4.5-2l-1.5 1.5M13 11l-1.5 1.5M19 .5A4.5 4.5 0 0 1 23.5 5M19 3.5A1.5 1.5 0 0 1 20.5 5m-15 4a6.5 6.5 0 0 1 13 0c0 .662-.111 1.32-.328 1.944l-2.85 8.195A3.516 3.516 0 0 1 12 21.5c-.98 0-1.865-.352-2.5-1m6-11.5a3.5 3.5 0 1 0-7 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Atm(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17.5 6.5c0 4.038.444 8.062 1.322 12M6.5 6.5c0 4.038-.444 8.062-1.322 12M4 6.5h16m-16 4H.5v-8h23v8H20m-12 11V21a2.5 2.5 0 0 0-2.5-2.5h-.322M16 21.5V21a2.5 2.5 0 0 1 2.5-2.5h.322m-13.644 0a55.143 55.143 0 0 1-.427 1.778l-.251.972v.25h15v-.25l-.25-.972a55.09 55.09 0 0 1-.428-1.778M12 14.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioDescription(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m2.75 16.25l1.605-2.841m0 0L7.408 8H8.75v5.409m-4.395 0H4.25m.105 0H8.75m0 0V16.5m9.25-9s.5 1.5.5 4.5s-.5 4.5-.5 4.5m2.5-9S21 9 21 12s-.5 4.5-.5 4.5m-19 4S.5 17 .5 12s1-8.5 1-8.5h21s1 3.5 1 8.5s-1 8.5-1 8.5zM11 16V8h3.75c.644 0 1.25.522 1.25 1.167v5.666c0 .645-.606 1.167-1.25 1.167z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Auditorium(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 19.5h7m-.5-5.5V6.5h-1.172a3 3 0 0 0-2.906 2.255L11.5 16.25v.25h9V18c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5m-21-4h7M8.5 14V6.5H7.328a3 3 0 0 0-2.906 2.255L2.5 16.25v.25H9m8.35-12s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25zm-9 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M21.5 9v10M5.5 9v10m-3-10v10m16-10v10M2 21h20M0 23.5h24M12 11h-1a1 1 0 0 0-1 1v.375a1 1 0 0 0 .72.96l2.56.747a1 1 0 0 1 .72.96V16a1 1 0 0 1-1 1h-1m0-6h1a1 1 0 0 1 1 1v.5M12 11V9m0 8h-1a1 1 0 0 1-1-1v-.5m2 1.5v2M23.5 6.25V7H.5v-.75C5.5 4.5 8.5 3 11.75.5h.5C15.5 3 18.5 4.5 23.5 6.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bar(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 19v4.5m0-4.5c0-4 3.167-9.806 7.063-13.053L18.5 4.75V4.5h-4.972M10 19c0-4-3.167-9.806-7.063-13.053L1.5 4.75V4.5h12.028M10 23.5H5m5 0h5m-9.5-15h9m3.5 1a4.5 4.5 0 1 0-4.472-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 20a2.5 2.5 0 0 1-5 0m13-2.5h-21v-.25l.848-.908a12 12 0 0 0 3.134-6.7l.336-2.684a6.23 6.23 0 0 1 12.364 0l.336 2.685a12 12 0 0 0 3.134 6.699l.848.908z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bicycle(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 6h4M7 8.5h9.5m3.5-5h-3.5v4c0 2.808.322 4.329 1.224 6.111c.469.927 1.237 1.889 2.276 1.889m-12.5-7c-.098 2.45-.41 3.617-1.208 5.12c-.487.918-1.253 1.88-2.292 1.88m1.5 5a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm13 0a5 5 0 1 1 0-10a5 5 0 0 1 0 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BillingDepartment(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 12.5h14.464m7.072 0H24m-20 4H3a1 1 0 0 0-1 1v.375a1 1 0 0 0 .72.96l2.56.747a1 1 0 0 1 .72.96v.958a1 1 0 0 1-1 1H4m0-6h1a1 1 0 0 1 1 1v.5m-2-1.5V15m0 7.5H3a1 1 0 0 1-1-1V21m2 1.5V24m6-14V7.5s-1.5-1-4-1s-4 1-4 1V10m14 14v-5.5h-3.5v-.25l.072-.15A24.999 24.999 0 0 0 15 7.35v-.328a8.58 8.58 0 0 1 3-.523c1.288 0 2.311.266 3 .523v.328c0 3.72.83 7.391 2.428 10.749l.072.15v.25H20V24M5.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.746 3.5 6.15 4.5 6.15 4.5zm12 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bin(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 9v9m5-9v9m-6-13.5h-6v.25l.24 1.05A70 70 0 0 1 4.5 21.398V22.5h15v-1.102c0-5.249.59-10.48 1.76-15.598l.24-1.05V4.5h-6m-7 0V4a3.5 3.5 0 1 1 7 0v.5m-7 0h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinPerson(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 11.5h3.146a2 2 0 0 0 1.857-1.257L11.5 6.5h.59a2.5 2.5 0 0 1 2.474 2.136L15.5 15v9m-4-9v9M.5.5l23 23m-11.35-19s1.6-1 1.6-2.25A1.75 1.75 0 0 0 12 .5c-.966 0-1.746.784-1.746 1.75c0 1.25 1.596 2.25 1.596 2.25zm-7.4 12.75a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5ZM3.125 22.5a1.625 1.625 0 1 0 0-3.25a1.625 1.625 0 0 0 0 3.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BinThrowPerson(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M22 15c-1.5 3-1.5 5.833-1.5 8.5h-6c0-2.667 0-5.5-1.5-8.5m0-3.5H9.854a2 2 0 0 1-1.857-1.257L6.5 6.5h-.59a2.5 2.5 0 0 0-2.474 2.136L2.5 15v9m4-9v9m9.75-10a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5ZM5.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.746 3.5 6.15 4.5 6.15 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bowling(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 7.5h3m-4.125 3h5.25M10 9.69a8 8 0 1 1 .5 12.055M5 .5A2.5 2.5 0 0 1 7.5 3c0 1.46-1 2.859-1 4.319v.08c0 1.04.337 2.05.96 2.882l.224.298a9.08 9.08 0 0 1 1.35 8.319L7.5 23.5h-5L.966 18.898a9.08 9.08 0 0 1 1.35-8.32l.224-.297c.623-.832.96-1.843.96-2.882v-.08c0-1.46-1-2.86-1-4.319A2.5 2.5 0 0 1 5 .5ZM17.5 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-2 3.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm4 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Boxing(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1l-6-2.75V14l1.018-5.088A3 3 0 0 1 5.459 6.5H5.5s2 2 3.5 2S11.5 7 12.5 3M8 11.5c1 .5 2.5 1 4 1M5.8 14l1.255-6.246M0 23.5c1 0 3.5-2.5 4.5-5.5m15-.5c2 0 3-1 3-1v-15s-1-1-3-1s-3 1-3 1v15s1 1 3 1Zm0 0v3m0 0c2.5 0 4 1 4 1v2h-8v-2s1.5-1 4-1ZM5.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.246 3.5 5.65 4.5 5.65 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Braille(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 15A2.5 2.5 0 0 0 6 12.5h-.5v8H6A2.5 2.5 0 0 1 8.5 23m0-6V5.5H9A2.5 2.5 0 0 1 11.5 8v2.5m6 12.5v-8a2.5 2.5 0 0 0-2.5-2.5h-.55m.05 2.5v-2a2.5 2.5 0 0 0-2.5-2.5h-.5m0 0V15M1.5 1v2m6-2v2m6-2v2m3-2v2m6-2v2m-21 2v2m3-2v2m12-2v2m3-2v2m3-2v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 17.5h2m2 0h6m2 0h2M2.5 4.5h19m0 9s-3.969 1-9.5 1c-5.531 0-9.5-1-9.5-1m0-12h19v21h-2.25l-.075-.12a4 4 0 0 0-3.392-1.88H8.217a4 4 0 0 0-3.392 1.88l-.075.12H2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cablecar(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 8.5V4.275M23.5 2s-9.609 3-23 3m21 10.863s-4.398-.75-9.5-.75s-9.5.75-9.5.75m6-7.363v6.722m7-6.722v6.722M14.5 2a92.58 92.58 0 0 1-2 .285m-1 .126c-.646.078-1.313.151-2 .219m12 5.87h-19v14h19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cafeteria(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 20.5h12.5v-7a3 3 0 1 1 3 3H20m-20 7h24m-9-13H2.5v.25l.063.1A19.008 19.008 0 0 1 5.5 21m0-21v1A6.5 6.5 0 0 0 12 7.5m2.5.5V4.5H12A3.5 3.5 0 0 1 8.5 1V0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 6V1m10 5V1m4 16v4.5h-18v-3m17.863-10H3.352M.5 18.25v.25h17.9l.15-.25l.234-.491A28 28 0 0 0 21.5 5.729V3.5h-18v2.128A28 28 0 0 1 .743 17.744z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3 9.5h3m8 8a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm9.5-11h-3a3 3 0 0 1-3-3h-7a3 3 0 0 1-3 3h-7v14h20a3 3 0 0 0 3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Campfires(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 17L9 19.402m6 1.696l8.5 2.402m-23 0l23-6.5M12 17c-3.314 0-6-2.15-6-5.425v-.242a5.11 5.11 0 0 1 1.524-3.636L10.02 5.23C11.288 3.977 12 1.772 12 0c0 1.772.712 3.977 1.98 5.23l2.496 2.467A5.113 5.113 0 0 1 18 11.333v.242C18 14.85 15.313 17 12 17Zm0 0c-1.381 0-2.5-.914-2.5-2.262c0-.561.229-1.17.635-1.567l1.04-1.015c.528-.516.825-1.423.825-2.152c0 .73.297 1.636.825 2.152l1.04 1.015c.407.397.635 1.006.635 1.567C14.5 16.086 13.38 17 12 17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Campground(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18 4.5h1.5c0 4.92 1.247 9.76 3.623 14.067l.377.683v.25H14m5 0V14m-.5-9.5h-14c0 4.92-1.246 9.76-3.623 14.067L.5 19.25v.25h13.85l.15-.25l.377-.683A29.12 29.12 0 0 0 18.5 4.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 15.5h2m2 0h8m2 0h2M3.127 10.64c1.326.275 4.643.86 8.873.86c4.23 0 7.547-.585 8.873-.86m-17.746 0A16 16 0 0 0 4.5 4.155V3.5h15v.656a16 16 0 0 0 1.373 6.484m-17.746 0a16 16 0 0 1-1.314 2.39l-.313.47v7h2.25l.075-.12a4 4 0 0 1 3.392-1.88h9.566a4 4 0 0 1 3.392 1.88l.075.12h2.25v-7l-.313-.47a15.993 15.993 0 0 1-1.314-2.39"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarRental(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 18.5h2m2 0h6m2 0h2M4.25 14.894c1.691.28 4.591.606 7.75.606c3.16 0 6.059-.326 7.75-.606m-15.5 0A12 12 0 0 0 5.5 9.561V9.5h13v.06c0 1.865.434 3.688 1.25 5.334m-15.5 0a12 12 0 0 1-1.64 2.476l-.111.13v6h2.25l.075-.12a4 4 0 0 1 3.392-1.88h7.566a4 4 0 0 1 3.392 1.88l.075.12h2.25v-6l-.111-.13a12.03 12.03 0 0 1-1.639-2.476M8.5 3.5h1M15 5V3.5M13 5V3.5M11.6 5a3 3 0 1 1 0-3h4.813c.347.492.815.885 1.36 1.142L18 3.25v.5l-.229.108A3.486 3.486 0 0 0 16.411 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Card(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M23.5 7.25c-3-.5-8.5-.75-11.5-.75s-8.5.25-11.5.75m23 4c-3-.5-8.5-.75-11.5-.75s-8.5.25-11.5.75m12 5.75c0-1.652.607-2.5 1.618-2.5c1.012 0 1.382 1.273 1.382 2.121h.047s.286-.848 1.203-.848s.991.943.991.943c.897.05 1.784.11 2.759.185m3-12.651V20.5H23c-3-.5-8-.75-11-.75S4 20 1 20.5H.5V4.25c3-.5 8.5-.75 11.5-.75s8.5.25 11.5.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CareStaffArea(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 22.5h24M13.5 20v-6S11 12.5 7 12.5S.5 14 .5 14v6m17-18.5v3h-3v3h3v3h3v-3h3v-3h-3v-3zm-10.756 9S3.999 8.752 3.999 6.566c0-1.691 1.344-3.062 3.002-3.062c1.658 0 2.995 1.371 2.995 3.062C9.996 8.75 7.26 10.5 7.26 10.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Casino(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 7.5C10 4.857 9.5.75 9.5.75V.5h14v.25S23 4.857 23 7.5s.5 6.75.5 6.75v.25H16m-8.5 1v2m-3.5 1v2m7-8v2m2.5-11v2m6-2v2m0 4v2m-19-2v.25s.5 4.107.5 6.75s-.5 6.75-.5 6.75v.25h14v-.25s-.5-4.107-.5-6.75s.5-6.75.5-6.75V9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChanginRoomDisabled(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 11.5v-5H8.328a3 3 0 0 0-2.906 2.255l-1.02 3.98M17 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.207M18.5 3.29c.136-.1.522-.423.83-.683A1.198 1.198 0 0 0 18.552.5H18.5c-.69 0-1.25.56-1.25 1.25V2m1.25 1.29c-1.318.961-2.877 1.542-4.56 1.905l-.44.095V7l.179.072a12.982 12.982 0 0 0 9.642 0L23.5 7V5.29l-.44-.095c-1.683-.363-3.242-.944-4.56-1.905ZM6 23.5a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11Zm3.35-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChanginRoomFamily(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 24v-6.5a2 2 0 0 0-2-2v-8s1.5-1 4-1a8.891 8.891 0 0 1 2.251.286M10.501 24v-5.5H7v-.25l.072-.15A25 25 0 0 0 9.5 7.352v-.329a8.58 8.58 0 0 1 3-.523c1.288 0 2.311.266 3 .523v.329a25.004 25.004 0 0 0 .116 2.406M20.75 24v-3.426c0-1.146.784-2.074 1.75-2.074v-6.463S21.188 11 19 11c-2.187 0-3.5 1.037-3.5 1.037V18.5c.967 0 1.75.929 1.75 2.074V24M5.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.246 3.5 5.65 4.5 5.65 4.5zm7 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25zM18.873 9S17.5 8.125 17.5 7.031c0-.845.672-1.531 1.502-1.531s1.498.686 1.498 1.531C20.5 8.125 19.13 9 19.13 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChanginRoomMan(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18.5 3.29c.136-.1.522-.423.83-.683A1.198 1.198 0 0 0 18.552.5H18.5c-.69 0-1.25.56-1.25 1.25V2m1.25 1.29c-1.318.961-2.877 1.542-4.56 1.905l-.44.095V7l.179.072a12.982 12.982 0 0 0 9.642 0L23.5 7V5.29l-.44-.095c-1.683-.363-3.242-.944-4.56-1.905ZM8 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24M5.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.746 3.5 6.15 4.5 6.15 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChanginRoomWoman(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18.5 3.29c.136-.1.522-.423.83-.683A1.198 1.198 0 0 0 18.552.5H18.5c-.69 0-1.25.56-1.25 1.25V2m1.25 1.29c-1.318.961-2.877 1.542-4.56 1.905l-.44.095V7l.179.072a12.982 12.982 0 0 0 9.642 0L23.5 7V5.29l-.44-.095c-1.683-.363-3.242-.944-4.56-1.905ZM4 24v-5.5H.5v-.25l.072-.15A25 25 0 0 0 3 7.35v-.328A8.58 8.58 0 0 1 6 6.5a8.58 8.58 0 0 1 3 .523v.328A25 25 0 0 0 11.428 18.1l.072.15v.25H8V24M5.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.746 3.5 6.15 4.5 6.15 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChangingRoomHanger(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 6V5a2.5 2.5 0 0 1 5 0v.216a3.1 3.1 0 0 1-.908 2.192a26 26 0 0 1-12.079 6.839L.5 14.5v3l.782.26a33.892 33.892 0 0 0 21.436 0l.782-.26v-3l-1.013-.253A26 26 0 0 1 14 10.405"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chapel(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.5 22v-7.5l3-3v-6l.272-.102A12 12 0 0 0 11.652.65L11.75.5h.5l.097.151a12 12 0 0 0 5.88 4.747l.273.102v6l3 3V22M9 8.5h6m-3-3v8m4.5 8.5v-6h-9v6m4.5-6v6m12 1.5H0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChargingStation(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 21V3.5l.088-.059a11.56 11.56 0 0 1 12.824 0l.088.059V21M6 22.5h18m-15.5-6h-1a4 4 0 0 1-4-4V6m0 0c1 0 1.5-.5 1.5-.5V3M3.5 6C2.5 6 2 5.5 2 5.5V3m14 9c-.5 2-2 3-2 3v.5h2v.5c-.5 2-2 3-2 3M2 3h3M2 3V1m3 2V1m13.5 8.5h-7V5.043a11.559 11.559 0 0 1 7 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 8.5h12M6 13h6M23.5 2H23c-3 .5-8 .75-11 .75S4 2.5 1 2H.5v21.5h.25l.154-.154A15.692 15.692 0 0 1 12 18.75c3 0 8 .25 11 .75h.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChildrenMustBeSupervised(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M19.25 24v-3.667c0-1.145.784-2.074 1.75-2.074v-6.222S19.688 11 17.5 11c-2.187 0-3.5 1.037-3.5 1.037v6.222c.967 0 1.75.929 1.75 2.075V24M9 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24M17.373 9S16 8.125 16 7.031c0-.845.672-1.531 1.502-1.531S19 6.186 19 7.031C19 8.125 17.63 9 17.63 9zM6.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C8.746 3.5 7.15 4.5 7.15 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CleaningRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18 20.5V20m0 .5a3 3 0 0 1-3 3h-2.5v-.25l.22-.357a12 12 0 0 0 1.78-6.29V16.5H18m0 4a3 3 0 0 0 3 3h2.5v-.25l-.22-.357a12 12 0 0 1-1.78-6.29V16.5H18m0 0V0M3.5 13.5V11a2.5 2.5 0 0 1 5 0v2.5m-6 10v-1.03a20 20 0 0 0-1.904-8.515L.5 13.75v-.25h11v.25l-.096.205A20 20 0 0 0 9.5 22.47v1.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CleaningRoomTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 7.5a4 4 0 0 1 4-4h1v-3h-11v3a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2v10h11v-7l-.145-.087C14 14 12.5 11.5 12.5 7.5Zm0 0h-4m7-3.874V5A2.5 2.5 0 0 0 18 7.5M20.5 0v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClimbingWall(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m16 4l-3.621 3.621a3 3 0 0 1-2.122.879H6.5l-2.121 2.121a3 3 0 0 0-.879 2.122V16m3-5v4.556a1.945 1.945 0 0 0 3.32 1.374l3.43-3.43h.25l.296.592A7.082 7.082 0 0 0 17.5 17.5m-5.5.498a8.859 8.859 0 0 1-.056-.11l-.194-.388zm0 0A10.623 10.623 0 0 0 17.5 23m2 1V8L17 5.5m-5-5L14.5 3m-4 5.5v5m-2.805-7s-1.81-.557-2.135-1.776a1.768 1.768 0 0 1 1.242-2.163a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m17.5 17.5l-.404-.566A19.67 19.67 0 0 0 12 12l.055-.493A34.15 34.15 0 0 0 12 3.5M22.498 12c0-5.798-4.7-10.498-10.498-10.498c-5.798 0-10.498 4.7-10.498 10.498c0 5.798 4.7 10.498 10.498 10.498c5.798 0 10.498-4.7 10.498-10.498Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClosedCaptioning(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.333 9.5v-1a1.5 1.5 0 0 0-1.5-1.5H6a1.5 1.5 0 0 0-1.5 1.5v7A1.5 1.5 0 0 0 6 17h2.833a1.5 1.5 0 0 0 1.5-1.5v-1m9-5v-1a1.5 1.5 0 0 0-1.5-1.5H15a1.5 1.5 0 0 0-1.5 1.5v7A1.5 1.5 0 0 0 15 17h2.833a1.5 1.5 0 0 0 1.5-1.5v-1M1.5 20.5S.5 17 .5 12s1-8.5 1-8.5h21s1 3.5 1 8.5s-1 8.5-1 8.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 9.362C11 6.5 13.5 4.5 16.536 4.5c3.964 0 6.964 3.154 6.964 7a6 6 0 0 1-6 6H.5v-.536l.367-.446c1.471-1.79 1.984-4.325 3.558-6.025A4.867 4.867 0 0 1 8 8.93a4.821 4.821 0 0 1 4.821 4.821"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeShop(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18.608 8.5h.892V9c-.577 2.831-1 5.714-1 8.604v.896h-.911m1.019-10c.24-1.508.52-3.01.84-4.506c.034-.162.052-.328.052-.494h-15c0 .166.018.332.052.494c.32 1.496.6 2.998.84 4.506m13.216 0H5.392m12.197 10a88.458 88.458 0 0 0-.089 3.964V23.5h-11v-1.036c0-1.323-.03-2.644-.089-3.964m11.178 0H6.41m-1.019-10H4.5V9c.577 2.831 1 5.714 1 8.604v.896h.911M2.5 3.5V3l.586-.293A3.817 3.817 0 0 0 5 .5h14a3.817 3.817 0 0 0 1.914 2.207L21.5 3v.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComputerRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13 16.5a11.003 11.003 0 0 0 4.041 3.93l1.459.82v.25h-13v-.25l1.459-.82A11.003 11.003 0 0 0 11 16.5m-9.5-3h21m0-11h-21v14h21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConferenceRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 14.5H1.5v-13h21v13H18m-.5 9.5a10.6 10.6 0 0 1 1.572-5.555l.428-.695v-.25h-15v.25l.428.695A10.6 10.6 0 0 1 6.5 24m9.5-9v-3.5s-1.5-1-4-1s-4 1-4 1V15m3.85-6.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoworkingSpace(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 21a7.526 7.526 0 0 0-.847-3.469A3.227 3.227 0 0 0 5.3 17H.5v.25c.655 1.147 1 2.43 1 3.75H10m5 2.5H0m15.5-14a7.527 7.527 0 0 0-.847-3.469A3.224 3.224 0 0 0 14.3 5.5H9.5v.25c.655 1.147 1 2.43 1 3.75H19m5 2.5h-9m-.5 9.5V19s-1.5-1-4-1c-1.006 0-1.85.162-2.5.355M23.5 10V7.5s-1.5-1-4-1c-1.006 0-1.85.162-2.5.356M10.35 16s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C12.246 15 10.65 16 10.65 16zm9-11.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crutch(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 13.5h6m-4.5 0v4c.96.96 1.5 2.263 1.5 3.621m0 0V23.5m0-2.379c0-1.358.54-2.66 1.5-3.621v-4m-3 10h3m4.5.5v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1V12m3.85-7.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyExchange(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 9h-1a1 1 0 0 0-1 1v.375a1 1 0 0 0 .72.96l2.56.747a1 1 0 0 1 .72.96V14a1 1 0 0 1-1 1h-1m0-6h1a1 1 0 0 1 1 1v.5M12 9V7.5m0 7.5h-1a1 1 0 0 1-1-1v-.5m2 1.5v1.5M1 3h12a8 8 0 0 1 8 8M1 3c.41 0 .99-.247 1.41-.503a5.633 5.633 0 0 0 1.458-1.305C4.186.792 4.5.318 4.5 0M1 3c.41 0 .99.247 1.41.503a5.64 5.64 0 0 1 1.458 1.305c.318.4.632.874.632 1.192M23 21H11a8 8 0 0 1-8-8m20 8c-.41 0-.99.247-1.41.503a5.636 5.636 0 0 0-1.458 1.305c-.318.4-.632.874-.632 1.192m3.5-3c-.41 0-.99-.247-1.41-.503a5.636 5.636 0 0 1-1.458-1.305c-.318-.4-.632-.874-.632-1.192M18 12a6 6 0 1 0-12 0a6 6 0 0 0 12 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Customs(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m18.5 18.5l5 5M8 7.5v8m6-8v8m-5.5-8a2.5 2.5 0 1 1 5 0m-2.5 14C5.201 21.5.5 16.799.5 11S5.201.5 11 .5S21.5 5.201 21.5 11S16.799 21.5 11 21.5Zm6-6H5v-.25l.128-.77a18.124 18.124 0 0 0 0-5.96L5 7.75V7.5h12v.25l-.128.77a18.126 18.126 0 0 0 0 5.96l.128.77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DangerPoison(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14 18.501c-.185 0-.712-.229-.945-.46a3.952 3.952 0 0 1-.762-1.064C12.143 16.67 12 16.299 12 16c0 .299-.144.67-.293.977c-.2.409-.462.765-.762 1.063c-.233.232-.76.461-.945.461M2.5 18.5V10a9.5 9.5 0 1 1 19 0v8.5a5 5 0 0 0-5 5h-9a5 5 0 0 0-5-5Zm8-8V14c-1 1-4 1.5-5 1.5v-1.015c0-1.591.69-3.117 2.225-3.538c1.087-.298 2.23-.447 2.775-.447Zm3 0V14c1 1 4 1.5 5 1.5v-1.015c0-1.591-.69-3.117-2.225-3.538c-1.087-.298-2.23-.447-2.775-.447Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Departure(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 20.5h24M9.019 10.193L1.417 7.188C3.185 4.713 4.422 1 4.422 1l1.238.53s.246 1.343 0 3.536l12.398 3.869a4.074 4.074 0 0 1 2.64 2.604c.373 1.116.843 2.473 1.225 3.427l-.353.353l-8.09-3.011c-3.931 2.304-7.467 3.011-7.467 3.011l-1.555-1.555s2.44-1.45 4.56-3.571Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeskForLaptop(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18.5 10.5h-13m13 0v-8h-13v8m13 0l.674 6.5H19c-1.95-.217-4.405-.5-7-.5c-2.595 0-5.05.283-7 .5h-.174l.674-6.5m5.5 4h2M3.5 7.06c-.403 0-.968.072-1.5.157L.608 21.5H.75s5.25-1 11.25-1s11.25 1 11.25 1h.087L22 7.217c-.532-.085-1.097-.157-1.5-.157"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discotheque(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-3.5-.75-5v-.5l-5-2.5V6.5h5V14m-3 3.5c-.21 0-.466.133-.737.344C5.958 19.24 5.5 21.718 5.5 24m-5-14l2.328-2.328A4 4 0 0 1 5.657 6.5h7.186a4 4 0 0 0 2.829-1.172L18 3m1.5 14.5a2 2 0 1 0-4 0a2 2 0 0 0 4 0Zm0 0v-8h.75s.75 4 3.75 4m-14.805-9s-1.81-.557-2.135-1.776A1.768 1.768 0 0 1 8.302.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dislike(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 15V3m17 .5v.16a16 16 0 0 0 3.761 10.307l.239.283v.25h-9V18a3.5 3.5 0 0 1-3.5 3.5h-.5V19A4.5 4.5 0 0 0 5 14.5h-.5v-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotSit(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 19.5h7m-.5-7v-6h-1.172a3 3 0 0 0-2.829 2m5.501 8H6.5v-.25L8 10.4m10 13.1c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4M.5.5l23 23m-11.15-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotSkateboardRollerboard(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.5 13.5H7v-.25l1-2.7M13.5 20c0-1.851.538-3.464 1.54-4.96M14.5 6.5l-2.162 5.838M10.764 15.5C9.94 16.94 9.5 18.326 9.5 20m9.5 1.5l-1.09.272a24.371 24.371 0 0 1-11.82 0L5 21.5m2 .482V24m10-2.018V24M.5.5l11.838 11.838M23.5 23.5l-8.46-8.46M2.5 5.5s4 1 9.5 1s9.5-1 9.5-1m-9.162 6.838l2.701 2.701M11.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotStandHere(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l2.65 2.65M23.5 23.5l-8.707-8.707m4.07 6.213c-.374 1.3-1.582 2.232-2.911 1.944a11.457 11.457 0 0 1-.151-.033c-1.225-.283-1.906-1.473-1.802-2.718l.196-2.336L16 18.18M9.894 12.5l-.346 2.233l-5.916 1.037l-1.163-4.047A12.039 12.039 0 0 1 2.489 5m.66-1.85l.018-.04l.443-1.172S6.207 1 8.916 1l.956 2.85c.612 1.826.778 3.77.483 5.673l-.111.72M3.15 3.15l7.094 7.094M14.4 14.4l-.755-4.877a12.04 12.04 0 0 1 .483-5.672L15.084 1c2.71 0 5.306.938 5.306.938l.832 2.2a12.04 12.04 0 0 1 .31 7.585l-1.164 4.047l-5.575-.977M14.4 14.4l.393.393M14.4 14.4l-4.156-4.156M8.199 22.917l-.15.033c-1.33.288-2.538-.644-2.912-1.944l-.636-2.213l5.304-.93L10 20.2c.104 1.245-.577 2.435-1.802 2.718Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownAngleArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M24 16c-.741 0-1.85.733-2.78 1.475c-1.2.954-2.247 2.094-3.046 3.401C17.575 21.856 17 23.044 17 24m0 0c0-.956-.575-2.145-1.174-3.124c-.8-1.307-1.847-2.447-3.045-3.401C11.85 16.733 10.74 16 10 16m7 8V12.5c0-6.627-5.373-12-12-12H0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 16c.742 0 1.85.733 2.78 1.475c1.2.954 2.247 2.094 3.046 3.401C11.425 21.856 12 23.044 12 24m0 0c0-.956.575-2.145 1.174-3.124c.8-1.307 1.847-2.447 3.045-3.401C17.15 16.733 18.26 16 19 16m-7 8V0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownLeftArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3 21L21 3M3 21c.676-.676 1.923-1.11 3.039-1.379c1.49-.359 3.036-.424 4.559-.252c1.182.134 2.484.4 3.009.924M3 21c.676-.676 1.11-1.923 1.379-3.039c.359-1.49.424-3.036.252-4.559c-.134-1.182-.4-2.484-.924-3.009"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownLeftTwoShortArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.01 20.991L21 3.001M4.07 5.08c.785.786 1.182 2.737 1.381 4.51c.258 2.282.159 4.6-.381 6.834c-.404 1.673-1.056 3.543-2.07 4.557m15.92-1.051c-.785-.785-2.737-1.181-4.51-1.381c-2.282-.258-4.6-.159-6.834.381c-1.673.404-3.543 1.056-4.557 2.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownRightArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M21 21L3 3m18 18c-.676-.676-1.923-1.11-3.039-1.379c-1.49-.359-3.036-.424-4.559-.252c-1.182.134-2.484.4-3.009.924M21 21c-.676-.676-1.11-1.923-1.379-3.039c-.359-1.49-.424-3.036-.252-4.559c.134-1.182.4-2.484.924-3.009"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownRightTwoShortArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M20.991 20.991L3.001 3.001m16.93 2.08c-.786.786-1.183 2.737-1.382 4.51c-.258 2.282-.159 4.6.381 6.834c.404 1.673 1.056 3.543 2.07 4.557M5.081 19.93c.786-.785 2.737-1.181 4.51-1.381c2.282-.258 4.6-.159 6.834.381c1.673.404 3.543 1.056 4.557 2.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownTwoShortArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 12c1.11 0 2.771 1.1 4.166 2.212c1.796 1.432 3.365 3.141 4.563 5.102c.897 1.469 1.758 3.253 1.758 4.686M22.5 12c-1.11 0-2.771 1.1-4.166 2.212c-1.796 1.432-3.365 3.141-4.563 5.102c-.897 1.469-1.758 3.253-1.758 4.686M12 24V0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DrinkingFountain(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path stroke="currentColor" d="M6.5 15v9m11 0v-.179c0-4.527-1.799-8.87-5-12.071v-.25H22M2.5 24v-9.5l2.22-5.182A3 3 0 0 1 7.479 7.5H8.5v3.757a3 3 0 0 0 .879 2.122L12 16M9.805 5.902s1.812.414 2.686-.46A1.726 1.726 0 0 0 12.498 3a1.722 1.722 0 0 0-2.438.01c-.874.873-.464 2.683-.464 2.683z"/><path fill="currentColor" d="m15.275 7.421l.137.481zm4.278 2.303l.223.447l.895-.447l-.224-.448zm-4.376-1.755l.48-.137l-.274-.961l-.481.137zm-.079-1.017l-.48.137l.274.962l.481-.138zm.656.87l.492-.09l-.18-.983l-.491.09zm2.04-.832l-.475-.155l-.31.95l.475.156zm.33 1.232l.436.246l.491-.871l-.435-.246zm2.103.663l-.267-.423l-.846.533l.267.423zm-5.09-.904l.04-.012l-.275-.961l-.04.011zm.236-.068l.04-.01l-.275-.962l-.04.01zm.04-.01a3.57 3.57 0 0 1 .34-.08l-.178-.984c-.146.026-.292.06-.437.102zm2.07.038c.223.072.438.167.642.281l.49-.87a4.555 4.555 0 0 0-.82-.362zm1.898 1.477c.061.098.119.2.172.306l.894-.448a4.551 4.551 0 0 0-.22-.391z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DrinkingWater(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7 6.5H3.5v.25l.24 1.05C4.91 12.917 5.5 18 5.5 23.5h13c0-5.5.59-10.583 1.76-15.7l.24-1.05V6.5H17M4.723 13H6a3 3 0 0 0 3-3a3 3 0 1 0 6 0a3 3 0 0 0 3 3h1.277M9.5 5c0 1.5 1 2.5 2.5 2.5s2.5-1 2.5-2.5a2.32 2.32 0 0 0-.635-1.604l-1.04-1.089C12.297 1.754 12 .782 12 0c0 .782-.297 1.754-.825 2.307l-1.04 1.089A2.324 2.324 0 0 0 9.5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DroneZone(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m10.5 3l-1.03.348A12.41 12.41 0 0 1 5.5 4m-5-1l1.03.348C2.81 3.78 4.15 4 5.5 4m0 0v3.5m18-4.5l-1.03.348A12.41 12.41 0 0 1 18.5 4m-5-1l1.03.348C15.81 3.78 17.15 4 18.5 4m0 0v3.5M5 15.5C3 15.5.5 18 .5 21M19 15.5c2 0 4.5 2.5 4.5 5.5M.5 7.5v4c5.5 0 8.5 6 8.5 6h6s3-6 8.5-6v-4zm11.5 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EbikeOne(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M16.5 5.5h-9V6c0 5-3.5 6.5-3.5 6.5m7.5-.5a5 5 0 0 0-5-5a6 6 0 0 0-6 6a9 9 0 0 0 9 9h.5m0 0c0 1 .5 1.5.5 1.5H13M10 22c0-1 .5-1.5.5-1.5H13m0 0v3m0-3h2m-2 3h2M11.498 12a5 5 0 0 1-5 5m0 0a4.978 4.978 0 0 1-3-1m3 1a4.98 4.98 0 0 1-2.098-.46M5 3h4M20 .5h-3.5V6c0 5 3.5 6.5 3.5 6.5M18.5 17a5 5 0 1 1 0-10a5 5 0 0 1 0 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EbikeTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M16.5 5.5h-9V6c0 5-3.5 6.5-3.5 6.5m6 9.5h-.5a9 9 0 0 1-9-9a6.003 6.003 0 0 1 4-5.659M10 22c0 1 .5 1.5.5 1.5H13M10 22c0-1 .5-1.5.5-1.5H13m0 0v3m0-3h2m-2 3h2M6.498 17a4.978 4.978 0 0 1-3-1m3 1a4.98 4.98 0 0 1-2.098-.46m2.098.46a5.001 5.001 0 0 0 4.584-3M5 3h4M20 .5h-3.5V6c0 5 3.5 6.5 3.5 6.5M.5.5l13.356 13.356M23.5 23.5l-6.856-6.856m0 0a5 5 0 1 0-2.789-2.789m2.79 2.79l-2.79-2.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElderlyPeoplePrioritySeating(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.004 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-3.5-.75-5v-.5l-5-2.5v-.25l1.803-5.228A3 3 0 0 1 9.143 6.5h1.36l.519 2.588a3 3 0 0 0 2.941 2.412H15m-7.496 6c-.21 0-.466.133-.737.344C4.962 19.24 4.504 21.718 4.504 24M16 14.5V13a1.5 1.5 0 0 1 3 0v10.5m-8.646-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectricScooter(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3 14a5 5 0 0 1 5 5v.5h.25S10 19 12 19s3.75.5 3.75.5H16V19a5 5 0 0 1 2.86-4.52M13 2.5h3.5v.825c0 3.852.808 7.65 2.36 11.155m0 0c.337.761.709 1.508 1.115 2.239m0 0a2.5 2.5 0 1 0 2.05 4.561a2.5 2.5 0 0 0-2.05-4.561ZM5.5 19a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectricWheelchair(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7 18.5a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm0 0L5 6m2 12.5h5.5a2 2 0 0 1 2 2m0 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm5.5 3c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-2H9l-.24-1.5M17.5 10h-3.05a2.5 2.5 0 0 1-2.17-1.26L11 6.5h-.588a2.5 2.5 0 0 0-2.469 2.895L8.52 13m-2.4 0H16.5a3 3 0 0 0 3-3v-.5m-8.9-5S9 3.5 9 2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Elevator(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 5.5v18M14 6h6.5v17.5h-17V6H10M5 3c.265 0 .66-.275.993-.553a4.857 4.857 0 0 0 1.088-1.276C7.295.804 7.5.358 7.5 0c0 .358.205.804.42 1.171c.285.49.659.918 1.087 1.276c.332.278.728.553.993.553m9-2.5c-.265 0-.66.275-.993.553a4.857 4.857 0 0 0-1.088 1.276c-.214.367-.419.813-.419 1.171c0-.358-.205-.804-.42-1.171a4.857 4.857 0 0 0-1.087-1.276C14.661.775 14.265.5 14 .5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElevatorAccesible(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.5 13V6.5h-1.172a3 3 0 0 0-2.906 2.255l-.963 3.764M17 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.207M23 14c-.265 0-.66.275-.993.553a4.857 4.857 0 0 0-1.088 1.276c-.214.367-.419.813-.419 1.171c0-.358-.205-.804-.42-1.171a4.857 4.857 0 0 0-1.087-1.276C18.661 14.275 18.265 14 18 14m5-4c-.265 0-.66-.275-.993-.553a4.857 4.857 0 0 1-1.088-1.276C20.705 7.804 20.5 7.358 20.5 7c0 .358-.205.804-.42 1.171c-.285.49-.659.918-1.087 1.276c-.332.278-.728.553-.993.553M6 23.5a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11Zm5.35-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElevatorPerson(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.5 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24M21 14c-.265 0-.66.275-.993.553a4.857 4.857 0 0 0-1.088 1.276c-.214.367-.419.813-.419 1.171c0-.358-.205-.804-.42-1.171a4.857 4.857 0 0 0-1.087-1.276C16.661 14.275 16.265 14 16 14m5-4c-.265 0-.66-.275-.993-.553a4.857 4.857 0 0 1-1.088-1.276C18.705 7.804 18.5 7.358 18.5 7c0 .358-.205.804-.42 1.171c-.285.49-.659.918-1.087 1.276c-.332.278-.728.553-.993.553M8.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmergencyExit(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-.5l-4.5-3.25V14l2.6-3.342A3 3 0 0 1 10.466 9.5H12v2a2 2 0 0 0 2 2h3M3.5 21c0-1 1.5-1.75 1.5-1.75c1.327-.664 2.263-.74 3.5-.749m6 4.999h5V.5h-15V10m7.35-2.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Entry(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17 12H1m10.5-5c0 .577.665 1.562 1.228 2.294a7.494 7.494 0 0 0 1.745 1.662C15.2 11.445 16.2 12 16.99 12c-.79 0-1.79.556-2.517 1.044a7.494 7.494 0 0 0-1.745 1.662c-.563.732-1.228 1.717-1.228 2.294m-3-10V2.5h.329A46 46 0 0 0 21.897.605L22.25.5h.25v23h-.25l-.353-.105A45.998 45.998 0 0 0 8.829 21.5H8.5V17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Escalator(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 22.5h5.03a8 8 0 0 0 6.813-3.807L18 9.5h5.5v-8h-5.03a8 8 0 0 0-6.813 3.807L6 14.5H.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EscalatorDownArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m23.5 10.5l-10-10m10 10c-.398-.398-.654-1.133-.812-1.79a7.746 7.746 0 0 1-.149-2.687c.08-.697.235-1.464.544-1.773m.417 6.25c-.398-.398-1.133-.654-1.79-.812a7.745 7.745 0 0 0-2.687-.149c-.697.08-1.464.235-1.773.544M18.485 23.5H23.5v-6H18L9.757 9.257A6 6 0 0 0 5.515 7.5H.5v6H6l8.243 8.243a6 6 0 0 0 4.242 1.757Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EscalatorDownPerson(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5 5.5l9 9h14v9H10.642a10 10 0 0 1-7.07-2.929L.5 17.5m11-3l1.018-5.588A3 3 0 0 1 15.459 6.5H16l.598 2.99a2.5 2.5 0 0 0 2.451 2.01H22m-5.65-7s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EscalatorUpArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m10.5.5l-10 10m10-10c-.398.398-1.133.654-1.79.812c-.878.212-1.79.25-2.687.149c-.697-.08-1.464-.235-1.773-.544M10.5.5c-.398.398-.654 1.133-.812 1.79a7.745 7.745 0 0 0-.149 2.687c.08.697.235 1.464.544 1.773M5.515 23.5H.5v-6H6l8.243-8.243A6 6 0 0 1 18.485 7.5H23.5v6H18l-8.243 8.243A6 6 0 0 1 5.515 23.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EscalatorUpPerson(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m23.5 5.5l-9 9H.5v9h12.858a10 10 0 0 0 7.07-2.929L23.5 17.5m-20-3l1.018-5.588A3 3 0 0 1 7.459 6.5H8l.598 2.99a2.5 2.5 0 0 0 2.451 2.01H14m-5.65-7s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationMark(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 19.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm2-3h-4v-.129A62 62 0 0 0 8.033.88L8 .75V.5h8v.25l-.033.129A61.999 61.999 0 0 0 14 16.37z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exit(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M24 12H8m10.5-5c0 .577.665 1.562 1.228 2.294a7.494 7.494 0 0 0 1.745 1.662C22.2 11.445 23.2 12 23.99 12c-.79 0-1.79.556-2.517 1.044a7.494 7.494 0 0 0-1.745 1.662c-.563.732-1.228 1.717-1.228 2.294m-4-10V2.5h-.329A46 46 0 0 1 1.103.605L.75.5H.5v23h.25l.353-.105A45.998 45.998 0 0 1 14.171 21.5h.329V17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M6.106 7.053a8.079 8.079 0 0 1 11.788 0L20.5 10c1 1 2.224 2 3.5 2c-1.276 0-2.5 1-3.5 2l-2.606 2.947a8.079 8.079 0 0 1-11.788 0L3.5 14c-1-1-2.224-2-3.5-2c1.276 0 2.5-1 3.5-2z"/><path d="M9.5 12a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceScan(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17 6v3M7 6v3m5-1c0 2.5.5 4.5.5 4.5l-2 1m-3.5 3c1.444.984 3 1.5 5 1.5s3.556-.516 5-1.5M7 1l-5.7.3L1 7m16-6l5.7.3L23 7m-6 16l5.7-.3l.3-5.7M7 23l-5.7-.3L1 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FallingRocks(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 1c3 1.5 5 4 5 4s-.5 3 1 6c5.5 2 8.5 5.5 8.5 5.5s.5 4 3.5 6v.5H1M20.252 8.129s-.559-.46-2.002-1.294c-1.443-.833-2.122-1.087-2.122-1.087s-.46.559-1.293 2.002c-.833 1.443-1.088 2.122-1.088 2.122s.56.46 2.003 1.293s2.122 1.088 2.122 1.088s.46-.56 1.293-2.003s1.087-2.121 1.087-2.121Zm2.32 8.45s-.153-.407-.655-1.276c-.501-.869-.778-1.205-.778-1.205s-.408.153-1.277.654c-.868.502-1.205.779-1.205.779s.153.408.654 1.276c.502.869.779 1.205.779 1.205s.408-.152 1.276-.654c.87-.501 1.206-.778 1.206-.778Zm-9.56-12s-.152-.407-.654-1.276c-.501-.869-.778-1.205-.778-1.205s-.408.153-1.277.654c-.869.502-1.205.779-1.205.779s.153.408.654 1.276c.502.869.779 1.205.779 1.205s.408-.152 1.276-.654c.869-.501 1.205-.778 1.205-.778Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FemaleSign(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 15.5a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15Zm0 0V24m-4-4h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintScan(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 10v4m-3.92-3c.37-2.282 1.985-4 3.92-4c2.21 0 4 2.239 4 5s-1.79 5-4 5c-1.935 0-3.55-1.718-3.92-4M11 20.425C7.33 19.871 4.5 16.31 4.5 12c0-4.694 3.358-8.5 7.5-8.5c1.902 0 3.639.802 4.96 2.125M13 20.425c3.67-.554 6.5-4.115 6.5-8.425c0-1.775-.48-3.422-1.3-4.785M7 1l-5.7.3L1 7m16-6l5.7.3L23 7m-6 16l5.7-.3l.3-5.7M7 23l-5.7-.3L1 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireAlarm(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 21.5H1A2.5 2.5 0 0 1 3.5 24m-3-2.5v.5m0-.5v-8H1A2.5 2.5 0 0 1 3.5 16v2m9 6v-8a2.5 2.5 0 0 0-2.5-2.5h-.5V15m0 1v-2A2.5 2.5 0 0 0 7 11.5h-.5V16m0-5V9A2.5 2.5 0 0 0 4 6.5h-.5V17M24 2.5h-3.719A6 6 0 0 1 16.295.984L15.75.5h-.25v1a4 4 0 0 0 4 4H24m0 13h-1.719a6 6 0 0 1-3.986-1.515l-.545-.485h-.25v1a4 4 0 0 0 4 4H24M5 3.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm13.281 6a6 6 0 0 1-3.986-1.516L13.75 7.5h-.25v1a4 4 0 0 0 4 4h1.219a6 6 0 0 1 3.986 1.515l.545.485h.25v-1a4 4 0 0 0-4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireExtinguisher(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15.5 3A1.5 1.5 0 0 1 17 1.5h3.5v3H17A1.5 1.5 0 0 1 15.5 3Zm0 0h-3a8 8 0 0 0-8 8v4A2.5 2.5 0 0 1 2 17.5M12 3V0m4.5 12.5H12v5h4.5m0-7v13h-9v-13a4.5 4.5 0 0 1 9 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireHazard(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 23.5c-1.933 0-3.5-1.24-3.5-3.07c0-.762.32-1.588.89-2.127l1.455-1.378c.74-.7 1.155-1.93 1.155-2.92c0 .99.416 2.22 1.156 2.92l1.456 1.378c.569.54.889 1.365.889 2.127c0 1.83-1.567 3.07-3.5 3.07Zm0 0c-4.556 0-8.25-2.972-8.25-7.5c0-1.885.754-4.027 2.096-5.36l3.432-3.41C11.02 5.496 12 2.45 12 0c0 2.45.98 5.497 2.723 7.23l3.431 3.41c1.342 1.333 2.096 3.475 2.096 5.36c0 4.528-3.693 7.5-8.25 7.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireHose(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.27 10.568a2.31 2.31 0 0 1-2.308 2.312a2.31 2.31 0 0 1-2.308-2.312A3.85 3.85 0 0 1 11.5 6.714a3.85 3.85 0 0 1 3.846 3.854a5.39 5.39 0 0 1-5.384 5.396a5.39 5.39 0 0 1-5.385-5.396c0-3.832 3.1-6.938 6.923-6.938s6.923 3.106 6.923 6.938c0 4.683-3.788 8.479-8.461 8.479c-4.674 0-8.462-3.796-8.462-8.48c0-5.534 4.477-10.02 10-10.02s10 4.486 10 10.02v10.35l-.333.347c-.658.684-.985 1.65-1.167 2.6c-.183-.95-.509-1.916-1.167-2.6l-.333-.347v-4.87m0 2.5h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstAid(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.527 4.5H.5v.25l.055.31a45.684 45.684 0 0 1 0 15.88l-.055.31v.25h23v-.25l-.055-.31a45.686 45.686 0 0 1 0-15.88l.055-.31V4.5h-7.027m-8.946 0a4.5 4.5 0 0 1 8.946 0m-8.946 0h8.946m-5.973 4v3h-3v3h3v3h3v-3h3v-3h-3v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FishingArea(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2.093 2.093C2.324 1.351 2.5 1 2.5 1H3s1.825 1.5 4 1.5h.035m-2.88 1.656A8.99 8.99 0 0 1 7.036 2.5m1.027 9.062C6.832 12.33 6 13.806 6 15.5c0 2.485 1.79 4.5 4 4.5m0 0A9 9 0 0 1 2.51 6.01M10 20c1.078 1.931 3.093 3.5 5.4 3.5v-.25S13.5 22 13.5 20c0-.672.215-1.26.5-1.744M10 20c.68-1.22 1.735-2.294 3-2.921M21.25 0v6.25l.25.25A1.5 1.5 0 1 1 20 8m-8-2V4M.5.5l10.499 10.499M23.5 23.5L10.999 10.999m0 0c2.41-.027 4.054-.552 6.001-2.499v-.25c-2-.25-3-.75-3-.75v-.25S15.5 6 17.25 5.5v-.25C15.724 3.223 13.401 2 10.8 2H10a8.987 8.987 0 0 0-2.965.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FishingAreaOne(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M21.6 12.5s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75m-.15 4h.3m-.3 0l-1.227 2m1.377-6c.966 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25m-.15-4v-8h-.25C16.373 2 11 7.373 9.5 12.5m12.4 0l1.223 2M0 19.5h7m.5-5.5V6.5H6.328a3 3 0 0 0-2.906 2.255L1.5 16.25v.25h8V18c0 5.5 2.5 5.5 2.5 5.5m11.496-5.25s-2.464-1.687-5.748 0c-3.285 1.688-5.748 0-5.748 0m11.496 2.5s-2.464-1.687-5.748 0c-3.285 1.688-5.748 0-5.748 0M7.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C9.246 3.5 7.65 4.5 7.65 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FishingAreaTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M21.25 6V0M20 8a1.5 1.5 0 1 0 1.5-1.5l-.25-.25V5.5M10 20c-2.21 0-4-2.015-4-4.5S7.79 11 10 11h.8c2.519 0 4.2-.5 6.2-2.5v-.25c-2-.25-3-.75-3-.75v-.25S15.5 6 17.25 5.5v-.25C15.724 3.223 13.401 2 10.8 2H10a8.987 8.987 0 0 0-2.965.5M10 20a9 9 0 0 1-9-9m9 9c1.078-1.931 3.093-3.5 5.4-3.5v.25S13.5 18 13.5 20s1.9 3.25 1.9 3.25v.25c-2.307 0-4.322-1.569-5.4-3.5Zm-9-9C1 4 2.5 1 2.5 1H3s1.825 1.5 4 1.5h.035M1 11a9.004 9.004 0 0 1 6.035-8.5M12 6V4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlashAllowed(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.05 2.05a2.801 2.801 0 0 0 3.136.571A2.801 2.801 0 0 0 17 0a2.8 2.8 0 0 0 1.814 2.621a2.801 2.801 0 0 0 3.136-.57a2.8 2.8 0 0 0-.571 3.135A2.801 2.801 0 0 0 24 7M3 11.5h2m16.5-3h-3a3 3 0 0 1-3-3h-6a3 3 0 0 1-3 3h-6v13h18a3 3 0 0 0 3-3zm-5 6a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlashNotAllowed(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l7.569 7.569M23.5 23.5l-2.882-2.882M12.05 2.05a2.801 2.801 0 0 0 3.136.571A2.801 2.801 0 0 0 17 0a2.8 2.8 0 0 0 1.814 2.621a2.801 2.801 0 0 0 3.136-.57a2.801 2.801 0 0 0-.571 3.135A2.801 2.801 0 0 0 24 7m-6 14.5H.5v-13H5m3.069-.431L8.159 8A2.997 2.997 0 0 0 9.5 5.5h6a3 3 0 0 0 3 3h3v10c0 .768-.289 1.47-.764 2l-.118.118M8.068 8.068l2.786 2.786m9.764 9.764l-4.471-4.471m0 0a4 4 0 0 0-5.293-5.293m5.293 5.293l-5.293-5.293M9 12.562A4 4 0 0 0 14.438 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 10V2.5h5l3 3h11v3m3 .25V8.5H4.6l-.15.25l-.234.492A28 28 0 0 0 1.5 21.272v.228h19v-.128a28 28 0 0 1 2.757-12.116z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forbidden(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M.5 12C.5 5.649 5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5S.5 18.351.5 12Z"/><path d="M4.5 14h15v-.25l-.068-.272a6.093 6.093 0 0 1 0-2.956l.068-.272V10h-15v.25l.068.272a6.092 6.092 0 0 1 0 2.956l-.068.272z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForbiddenTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 12C.5 5.649 5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5S.5 18.351.5 12Zm7.342-6.836a8 8 0 0 1 10.994 10.994a67.213 67.213 0 0 1-2.216-1.829c-.883-.77-2.912-2.742-3.56-3.39c-.647-.647-2.62-2.676-3.39-3.56a67.247 67.247 0 0 1-1.828-2.215ZM5.164 7.842a8 8 0 0 0 10.994 10.994a67.213 67.213 0 0 0-1.829-2.216c-.77-.883-2.742-2.912-3.39-3.56c-.647-.647-2.676-2.62-3.56-3.39a67.226 67.226 0 0 0-2.215-1.828Z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fountain(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 14.5v-.25a3.75 3.75 0 1 1 7.5 0m0 0V17V7.75m0 6.5a3.75 3.75 0 1 1 7.5 0v.25M6.5 8v-.25a2.75 2.75 0 0 1 5.5 0m0 0a2.75 2.75 0 1 1 5.5 0V8m-9-5.5v-.25a1.75 1.75 0 1 1 3.5 0m0 0V4m0-1.75a1.75 1.75 0 1 1 3.5 0v.25m-13 21v-.219a6 6 0 0 0-1.516-3.986L.5 18.75v-.25h23v.25l-.485.545a6 6 0 0 0-1.515 3.986v.219z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fragile(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m12 18.939l.002.086v4.451M12 18.94c-.055-.985-1.055-2.187-2.998-2.969c-2.786-1.112-4.506-3.772-4.504-6.876c0-.6.073-1.197.217-1.78L6.502.524h11l1.783 6.79a7.42 7.42 0 0 1 .217 1.78c.002 3.104-1.718 5.765-4.504 6.876c-1.943.782-2.943 1.984-2.998 2.969Zm.002 4.537h6m-6 0h-6m7.5-13.45l.232-.233a6.035 6.035 0 0 0 1.768-4.268h-3.5v-.25l.232-.23a6.04 6.04 0 0 0 1.768-4.27v-.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunicularRailway(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.478 18.149a1.5 1.5 0 0 1-2.954.52m11.999-2.25a1.5 1.5 0 0 0 2.954-.52M8 11.758V4.636m8 5.648V3.182m6.97 6.23c.019-.477.03-.98.03-1.503C23 4.41 22.5 2 22.5 2l-21 3.818S1 8.41 1 11.91c0 .523.011 1.022.03 1.492m21.94-3.99C22.862 12.127 22.5 14 22.5 14l-21 3.818s-.362-1.743-.47-4.417m21.94-3.99c-10.656.973-21.302 3.818-21.94 3.99M23 19L1 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gallery(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 19.5c3.314 0 6-2.5 6-4.5m0 0c0 2 2.686 4.5 6 4.5s6-2.5 6-4.5M6 15v8.5M18 15c0 2 2.686 4.5 6 4.5M18 15v8.5m-14 0h4m8 0h4m1.5-11h-19V.5h19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Garden(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 6.5V7m0-.5a6 6 0 0 1 6-6h.5V7a6.5 6.5 0 0 1-6.5 6.5m0-7a6 6 0 0 0-6-6h-.5V7a6.5 6.5 0 0 0 6.5 6.5m0 0V23m0-9.5v9m0 1H7.5a6 6 0 0 1-6-6H6a6 6 0 0 1 6 6Zm0 0h4.5a6 6 0 0 0 6-6H18a6 6 0 0 0-6 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GasStation(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 21V3.5l.088-.059a11.56 11.56 0 0 1 12.824 0l.088.059V21M6 22.5h18M0 4.5h4.5v8a4 4 0 0 0 4 4m-7-12a3 3 0 0 0 3 3m14 2h-7V5.043a11.559 11.559 0 0 1 7 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftShop(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 6.5v15m0-15V4m0 2.5H9.5A2.5 2.5 0 1 1 12 4m0 2.5h2.5A2.5 2.5 0 1 0 12 4M3.25 14h17.5m-17.5 0c0-2.328-.23-4.65-.686-6.932L2.5 6.75V6.5h19v.25l-.064.318A35.346 35.346 0 0 0 20.75 14m-17.5 0c0 2.328-.23 4.65-.686 6.932l-.064.318v.25h19v-.25l-.064-.318A35.345 35.345 0 0 1 20.75 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glass(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.5 7.015V.5h3v6.515a6 6 0 0 0 1.757 4.242L17.5 13.5v10h-11v-10l2.243-2.243A6 6 0 0 0 10.5 7.015Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M23 12c0-6.075-4.925-11-11-11m11 11c0 6.075-4.925 11-11 11m11-11c0 2.21-4.925 4-11 4S1 14.21 1 12M12 1C5.925 1 1 5.925 1 12M12 1c2.761 0 5 4.925 5 11s-2.239 11-5 11m0-22C9.239 1 7 5.925 7 12s2.239 11 5 11M1 12c0 6.075 4.925 11 11 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GolfCourse(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 23.5h5m-2.5 0V0m0 .5h1.879C9.737.5 11.039 1.04 12 2c.96.96 2.263 1.5 3.621 1.5h.879v7h-1.879c-1.358 0-2.66-.54-3.621-1.5a5.121 5.121 0 0 0-3.621-1.5H6.5m17 11s-3.5-4-11.5-4c-1.28 0-2.446.102-3.5.275M.5 18.5s1.261-1.441 4-2.594m11 5.594a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupTraining(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 18v5m2-6v7m15-6v5m-2-6v7m-13-3.5h13m-2 0l-1-4s-1.5-1-3.5-1s-3.5 1-3.5 1l-1 4M23.5 9v5m-2-6v7M13 11.5h8.5m-2 0l-1-4s-1.5-1-3.5-1s-3.5 1-3.5 1l-.25 1m-2.4 5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25zm6-9s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuestHeightLimit(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.524 22v-4.5a2 2 0 0 1 2-2v-6s-1.5-1-4-1s-4 1-4 1v6a2 2 0 0 1 2 2V22M23.936.5H.064m23.873 23H.063M18.524 3v18m0-18c0 .358.205.804.419 1.171c.285.49.66.918 1.088 1.276c.332.278.728.553.993.553m-2.5-3c0 .358-.206.804-.42 1.171c-.285.49-.659.918-1.087 1.276c-.333.278-.728.553-.993.553m2.5 15c0-.358.205-.804.419-1.171c.285-.49.66-.918 1.088-1.276c.332-.278.728-.553.993-.553m-2.5 3c0-.358-.206-.804-.42-1.171a4.856 4.856 0 0 0-1.087-1.276c-.333-.278-.728-.553-.993-.553M7.374 6.5s-1.6-1-1.6-2.25a1.746 1.746 0 1 1 3.495 0c0 1.25-1.595 2.25-1.595 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuestWithinHeightLimitMustBeSupervised(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.578 21.86v-4.437a1.97 1.97 0 0 1 1.971-1.972V9.535S7.07 8.55 4.606 8.55c-2.465 0-3.944.986-3.944.986v5.916c1.089 0 1.972.883 1.972 1.972v4.436m12.817 0v-3.615c0-.887.662-1.607 1.479-1.607v-4.82s-1.11-.803-2.958-.803c-1.849 0-2.958.803-2.958.803v4.82c.817 0 1.48.72 1.48 1.607v3.615M23.83.662H.17m23.66 22.676H.17M21.366 3.127v17.746m0-17.746c0 .353.203.793.414 1.155c.281.483.65.904 1.072 1.257c.328.274.718.545.98.545m-2.466-2.957c0 .353-.202.793-.413 1.155c-.281.483-.65.904-1.072 1.257c-.328.274-.718.545-.98.545m2.465 14.79c0-.354.203-.794.414-1.156c.281-.483.65-.904 1.072-1.257c.328-.274.718-.545.98-.545m-2.466 2.957c0-.353-.202-.793-.413-1.155a4.787 4.787 0 0 0-1.072-1.257c-.328-.274-.718-.545-.98-.545M4.458 6.577S2.88 5.591 2.88 4.36c0-.953.773-1.725 1.726-1.725s1.72.772 1.72 1.725c0 1.232-1.573 2.218-1.573 2.218zm9.389 2.712s-1.354-.863-1.354-1.941c0-.834.663-1.51 1.48-1.51c.818 0 1.478.676 1.478 1.51c0 1.078-1.35 1.94-1.35 1.94z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gym(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m8.11 15.89l7.779-7.78M2 22l1.848-1.848M22 2l-1.848 1.848M2.1 9.878L14.121 21.9l-.177.176l-1.737-.208c-1.75-.21-3.522.138-5.098.926L6.697 23L1 17.303l.206-.412c.788-1.576 1.136-3.348.926-5.098l-.208-1.737zm7.78-7.776l12.02 12.02l.177-.177l-.208-1.737c-.21-1.75.138-3.522.926-5.098L23 6.697L17.303 1l-.413.206c-1.575.788-3.348 1.136-5.097.926l-1.738-.208z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hairdresser(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m12 16.61l-4.276 4.589C7 22 6.102 22.5 5.002 22.5a3.5 3.5 0 1 1 3.162-5M12 16.61l9.641-10.35a3.2 3.2 0 0 0-.078-4.444l-.316-.316h-.25l-.106.133A55.371 55.371 0 0 1 8.6 12.961L2.36 6.26a3.2 3.2 0 0 1 .078-4.444l.316-.316M12 16.61l4.276 4.589C17 22 17.898 22.5 18.998 22.5a3.5 3.5 0 1 0-3.162-5M2.753 1.5h.079m-.08 0h.25l.107.133A55.374 55.374 0 0 0 12 10.477"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M22.5 13.5a6.575 6.575 0 0 1-4.369-1.66l-.381-.34h-.25v11h.25l.381-.34A6.575 6.575 0 0 1 22.5 20.5V12c0-5.799-4.701-10.5-10.5-10.5S1.5 6.201 1.5 12v8.5c1.61 0 3.165.591 4.369 1.66l.381.34h.25v-11h-.25l-.381.34A6.576 6.576 0 0 1 1.5 13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HealthServices(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 8.5v-8h.25l.27.054c.652.13 2.315.196 2.98.196s2.329-.066 2.98-.196L15.25.5h.25v8h8v.25l-.054.27c-.13.652-.196 2.315-.196 2.98s.066 2.329.196 2.98l.054.27v.25h-8v8h-.25l-.27-.054c-.652-.13-2.315-.196-2.98-.196s-2.328.066-2.98.196l-.27.054H8.5v-8h-8v-.25l.054-.27c.13-.652.196-2.315.196-2.98S.684 9.672.554 9.02L.5 8.75V8.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 8.5V7.486m0 .264c0-2.9 2.35-5.25 5.25-5.25s5.25 2.382 5.25 5.282c0 1.56-.687 3.055-1.88 4.062l-5.657 4.776A8.35 8.35 0 0 0 12 23a8.35 8.35 0 0 0-2.963-6.38L3.38 11.844A5.327 5.327 0 0 1 1.5 7.782c0-2.9 2.35-5.282 5.25-5.282S12 4.85 12 7.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Helicopter(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15.5 7.5H5.158a5.857 5.857 0 0 1-4.408-2H.5v5h4a7 7 0 0 1 7 7h12v-2a8 8 0 0 0-8-8Zm0 0v3a4 4 0 0 0 4 4h3.938m-7.938-7V4m0 0c-2.275 0-4.544-.245-6.767-.73L7.5 3m8 1c2.275 0 4.544-.245 6.767-.73L23.5 3m-9 14.5v3.754m6-3.754v3.754m-9 .746l1.096-.313a17.851 17.851 0 0 1 9.808 0L23.5 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HideEyeCrossbar(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17.416 17.416a8.13 8.13 0 0 0 .478-.47L20.5 14c1-1 2.224-2 3.5-2c-1.276 0-2.5-1-3.5-2l-2.606-2.947a8.079 8.079 0 0 0-11.31-.469m8.713 12.213a8.077 8.077 0 0 1-9.19-1.85L3.5 14c-1-1-2.224-2-3.5-2c1.276 0 2.5-1 3.5-2l1.408-1.592M.5.5l23 23m-9.732-9.732a2.5 2.5 0 0 0-3.536-3.536"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HighVoltage(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 14.25v.25h7.974A30.864 30.864 0 0 1 7.5 23v.5c1.32 0 3.184-.544 5.089-1.273c5.044-1.931 7.912-7.076 7.915-12.477L20.5 9.5h-5.586A25.22 25.22 0 0 0 16.5.699V.5h-9v.08a25.36 25.36 0 0 1-4 13.67Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hiking(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.004 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-3.5-.75-5v-.5l-5-2.5v-.25l1.803-5.228A3 3 0 0 1 11.143 6.5h1.361l.518 2.588a3 3 0 0 0 2.942 2.412h3.04m-9.5 6c-.21 0-.466.133-.737.344C6.963 19.24 6.504 21.718 6.504 24m11.932-11l-1.932 10.9M19.5 7l-.532 3M6.504 6L4.78 11m7.574-6.5s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75c.967 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HoldTight(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 8.5h4.5A3.5 3.5 0 0 0 10 5h4.5m3 0h3v.5a3 3 0 0 1-3 3h-4h5.25a1.75 1.75 0 1 1 0 3.5H16.5h1.25a1.75 1.75 0 1 1 0 3.5H16.5h.25a1.75 1.75 0 1 1 0 3.5H2M16 8.5V0m0 24v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M22.5 21V10.25l-.247-.113a20 20 0 0 1-8.942-8.104L13 1.5h-2l-.311.533a20 20 0 0 1-8.942 8.104l-.247.113V21M23 22.5H1M15 21v-4a3 3 0 1 0-6 0v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17.5 2.5h-11l-.299.986a14 14 0 0 1-4.366 6.636L.5 11.25v.25m17-9l-.302 1.005a14 14 0 0 1-4.387 6.683L11.55 11.25l-.3.25H.5m17-9l.302 1.005a14 14 0 0 0 4.388 6.683l1.26 1.062l.05.05v10.2H.5v-10m11-.208V21.5m6 0V17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 22v-7.961a20 20 0 0 0-1.567-7.761L1.5 5.25V5h6m13 17v-7.961a20 20 0 0 1 1.567-7.761L22.5 5.25V5h-6m0 17v-6.5h-9V22M24 23.5H0m12-21V5m0 0v2.5M12 5h2.5M12 5H9.5M12 15.5V22M7.5.5h9v9h-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HotelRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 18.5V14l.676-.184a37.335 37.335 0 0 1 19.648 0L22.5 14v4.5m-21 0s0 3-1.5 3m1.5-3h21m0 0s0 3 1.5 3M3.5 11c0-1.989-.297-3.966-.882-5.867L2.5 4.75V4.5h19v.25l-.118.383A19.95 19.95 0 0 0 20.5 11M12 7.5H6.5V11M12 7.5V11m0-3.5h5.5V11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HotelRoomTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 18.5V4m0 14.5s0 3-1.5 3m1.5-3h21m0 0s0 3 1.5 3m-1.5-3V13A3.5 3.5 0 0 0 19 9.5h-8.5v4m-9 2h21m-14-4.1S7.5 13 6.25 13a1.75 1.75 0 0 1-1.75-1.75c0-.966.784-1.746 1.75-1.746C7.5 9.504 8.5 11.1 8.5 11.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hydrant(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 7.5h11m-11 0A5.5 5.5 0 0 1 12 2M6.5 7.5v2a1 1 0 0 1-1 1H4v5h1.5a1 1 0 0 1 1 1v1.828a6 6 0 0 1-1.986 4.46L4 23.25v.25h16v-.25l-.514-.462a6 6 0 0 1-1.986-4.46V16.5a1 1 0 0 1 1-1H20v-5h-1.5a1 1 0 0 1-1-1v-2m0 0A5.5 5.5 0 0 0 12 2m0 0V0M4.5 7.5h15M9.5 21v2.5m5-2.5v2.5m0-10.5a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M22.5 18.5h-.304c-.94 0-1.877-.274-2.564-.917a10.931 10.931 0 0 1-2.061-2.624l-.82-1.459h-.5l-.822 1.459a10.922 10.922 0 0 1-1.75 2.32a4.604 4.604 0 0 1-1.919-1.062a15.292 15.292 0 0 1-2.868-3.674L7.75 10.5h-.498l-1.141 2.043a15.3 15.3 0 0 1-3.286 4.053a3.54 3.54 0 0 1-1.325.722m0 5.182v-21h21v21zm15-13a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InPatient(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 20.5V6m0 14.5s0 3-1.5 3m1.5-3h21m0 0s0 3 1.5 3m-1.5-3V15a3.5 3.5 0 0 0-3.5-3.5h-8.5v4m-9 2h21m-6.5-8v-2s-1.5-1-4-1s-4 1-4 1v2m.5 3.9S7.5 15 6.25 15a1.75 1.75 0 0 1-1.75-1.75c0-.966.784-1.746 1.75-1.746c1.25 0 2.25 1.596 2.25 1.596zm3.352-8.9s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75c.967 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationDeskPeople(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 12.5h14.464m7.072 0H24M10 10V7.5s-1.5-1-4-1s-4 1-4 1V10m14 14v-5.5h-3.5v-.25l.072-.15A24.999 24.999 0 0 0 15 7.35v-.328a8.58 8.58 0 0 1 3-.523c1.288 0 2.311.266 3 .523v.328c0 3.72.83 7.391 2.428 10.749l.072.15v.25H20V24M5.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.746 3.5 6.15 4.5 6.15 4.5zm12 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationDeskSymbol(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 7V5M9 9.5h3v8m0 0H8m4 0h4m-4 6C5.649 23.5.5 18.351.5 12S5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntensiveCare(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 20.5V6m0 14.5s0 3-1.5 3m1.5-3h21m0 0s0 3 1.5 3m-1.5-3V15a3.5 3.5 0 0 0-3.5-3.5h-8.5v4m-9 2h21M10.5 4h1.573a1.5 1.5 0 0 0 1.342-.83L14.75.5H15l3.5 7h.25l1.335-2.67A1.5 1.5 0 0 1 21.427 4H23M8.5 13.4S7.5 15 6.25 15a1.75 1.75 0 0 1-1.75-1.75c0-.966.784-1.746 1.75-1.746c1.25 0 2.25 1.596 2.25 1.596z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IrisScan(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m7 1l-5.7.3L1 7m16-6l5.7.3L23 7m-6 16l5.7-.3l.3-5.7M7 23l-5.7-.3L1 17m-1-5c1.276 0 2.5 1 3.5 2l2.606 2.947a8.079 8.079 0 0 0 11.788 0L20.5 14c1-1 2.224-2 3.5-2c-1.276 0-2.5-1-3.5-2l-2.606-2.947a8.079 8.079 0 0 0-11.788 0L3.5 10c-1 1-2.224 2-3.5 2Zm12 2.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kitchen(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 22.5H1.5v-20s2-1 5.25-1s5.25 1 5.25 1zm0 0v-10h.25s1.75 1 5 1s5-1 5-1h.25v10zM1.5 9.5s2-1 5.25-1s5.25 1 5.25 1m-8-5v2M4 11v2m14.5.449V9.75a1.25 1.25 0 1 0-2.5 0V10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laboratory(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 12.5h.146c2.206 0 4.381.514 6.354 1.5c1.973.986 4.148 1.5 6.354 1.5H20.5m-13-14h9v.25l-.707.972a6.762 6.762 0 0 0 .688 8.759L22.5 17.5c0 2-1 3.5-3 5h-15c-2-1.5-3-3-3-5l6.02-6.02a6.762 6.762 0 0 0 .687-8.758L7.5 1.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laundry(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M16.9 13H16a2 2 0 0 1-2-2a2 2 0 1 1-4 0a2 2 0 0 1-2 2h-.9M5.5 4h3m.217 17h6.566a4 4 0 0 1 3.392 1.88l.075.12H21V1H3v22h2.25l.075-.12A4 4 0 0 1 8.717 21ZM12 17a5 5 0 1 1 0-10a5 5 0 0 1 0 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaundryTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9 14.5h6M17 4H7m10 0a3 3 0 0 0 3-3V.5H7.5a3.5 3.5 0 1 0 0 7H20V7a3 3 0 0 0-3-3ZM3.5 23.5h17c0-4 .934-7.79 2.722-11.217l.278-.533v-.25H.5v.25l.278.533C2.566 15.71 3.5 19.5 3.5 23.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftAngleArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M16 0c0 .742.733 1.85 1.475 2.78c.954 1.2 2.094 2.247 3.401 3.046C21.856 6.425 23.044 7 24 7m0 0c-.956 0-2.145.575-3.124 1.174c-1.307.8-2.447 1.847-3.401 3.045C16.733 12.15 16 13.26 16 14m8-7H12.5C5.873 7 .5 12.373.5 19v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M16 5c0 .742.733 1.85 1.475 2.78c.954 1.2 2.094 2.247 3.401 3.046C21.856 11.425 23.044 12 24 12m0 0c-.956 0-2.145.575-3.124 1.174c-1.307.8-2.447 1.847-3.401 3.045C16.733 17.15 16 18.26 16 19m8-7H0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftLuggage(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 7.5v11m8-11v11m-1-11a3 3 0 1 0-6 0M1 1.5h21.5v21H1m18.5-4h-16v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L3.5 7.676V7.5h16v.176l-.203 1.346a26.748 26.748 0 0 0 0 7.956l.203 1.345z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftTwoShortArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 1.5c0 1.11 1.1 2.771 2.212 4.166c1.432 1.796 3.141 3.365 5.102 4.563c1.469.897 3.253 1.758 4.686 1.758M12 22.5c0-1.11 1.1-2.771 2.212-4.166c1.432-1.796 3.141-3.365 5.102-4.563c1.469-.897 3.253-1.758 4.686-1.758M24 12H0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LessMinus(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 14.5h23v-.25l-.054-.27a10.101 10.101 0 0 1 0-3.96l.054-.27V9.5H.5v.25l.054.27a10.1 10.1 0 0 1 0 3.96l-.054.27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LgbtFriendly(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.507 7.5h10.487m10.5 0H12.005m-8.996 4h17.98m-13.28 4h8.58m-4.296-8a5.25 5.25 0 0 0-5.244-5c-2.9 0-5.25 2.382-5.25 5.282c0 1.56.688 3.055 1.88 4.062l5.657 4.776A8.35 8.35 0 0 1 12 23a8.35 8.35 0 0 1 2.964-6.38l5.657-4.776a5.327 5.327 0 0 0 1.88-4.062c0-2.9-2.351-5.282-5.25-5.282a5.25 5.25 0 0 0-5.245 5m-.012 0h.012"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Library(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 7.5V17m0-9.5a3 3 0 0 0-3-3H2v1.085A62.99 62.99 0 0 1 .5 19.25v.25H9a3 3 0 0 1 3 3m0-15a3 3 0 0 1 3-3h7v1.085c0 4.596.503 9.179 1.5 13.665v.25H15a3 3 0 0 0-3 3m0 0v.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 9v12m17-.5v-.16a16 16 0 0 1 3.761-10.307l.239-.283V9.5h-9V6A3.5 3.5 0 0 0 10 2.5h-.5V5A4.5 4.5 0 0 1 5 9.5h-.5v11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationOff(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m16.19 16.19l1.337-1.215A8.365 8.365 0 0 0 20.25 8.8c0-4.558-3.694-8.3-8.25-8.3c-3.104 0-5.807 1.737-7.216 4.284M12 23.92a9.041 9.041 0 0 1 2.383-6.037M12 23.92a9.04 9.04 0 0 0-2.96-6.61l-2.567-2.334a8.365 8.365 0 0 1-2.6-7.603M12 23.92V24v-.057M.5.5l23 23M9.564 9.564a2.5 2.5 0 1 1 1.872 1.872"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationPin(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 23.92a9.04 9.04 0 0 0-2.96-6.61l-2.567-2.334A8.365 8.365 0 0 1 3.75 8.799C3.75 4.242 7.444.5 12 .5s8.25 3.741 8.25 8.298c0 2.343-.989 4.6-2.723 6.177l-2.568 2.334a9.041 9.041 0 0 0-2.96 6.61Zm0 0V24v-.057M12 11.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 10.5V6a5.5 5.5 0 1 1 11 0v4.5M12 15v4m8.5 4.5h-17v-.25l.075-.327a26.342 26.342 0 0 0 0-11.846L3.5 10.75v-.25h17v.25l-.075.327a26.34 26.34 0 0 0 0 11.846l.075.327z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lockers(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M.5 9.5v.25l4.75 4.75h.25a3.89 3.89 0 0 1 5.5 0l9 9h3.5V20l-1-1h-2v-2h-2v-2l-4-4a3.89 3.89 0 0 1 0-5.5v-.25L9.75.5H9.5a9 9 0 0 0-9 9Z"/><path d="M4.5 6a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LostAndFound(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 11v-1a2.5 2.5 0 0 1 5 0v.039a2 2 0 0 1-.75 1.562l-1 .798a2 2 0 0 0-.75 1.562V15.5m0 1.5v2M5.5 4.5v17m13 0v-17h-2.027m-8.946 0H.5v.25l.055.31a45.684 45.684 0 0 1 0 15.88l-.055.31v.25h23v-.25l-.055-.31a45.686 45.686 0 0 1 0-15.88l.055-.31V4.5h-7.027m-8.946 0a4.5 4.5 0 0 1 8.946 0m-8.946 0h8.946"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lounge(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M22 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-9L8.621 8.985a2 2 0 0 1 1.94-2.485h.939l4 8m1 5.5h-5a4.643 4.643 0 0 1-4.504-3.517L3 .5m6.195 4s-1.81-.557-2.135-1.776A1.768 1.768 0 0 1 8.302.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Luggage(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 4.5v17m13-17v17m-2.027-17H23.5v.25l-.055.31a45.687 45.687 0 0 0 0 15.88l.055.31v.25H.5v-.25l.055-.31a45.682 45.682 0 0 0 0-15.88L.5 4.75V4.5h7.027m8.946 0a4.5 4.5 0 0 0-8.946 0m8.946 0H7.527"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LuggageCheckIn(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 23.5h20.5v-19m0 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-16 4v11m8-11v11m-1-11a3 3 0 1 0-6 0m11 11h-16v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L1.5 8.676V8.5h16v.176l-.203 1.346a26.747 26.747 0 0 0 0 7.956l.203 1.345z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LuggageConveyorBelt(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8 3.5v11m8-11v11m-1-11a3 3 0 1 0-6 0M4 21h1m4 0h1m4 0h1m4 0h1m0-6.5H4v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L4 3.676V3.5h16v.176l-.203 1.346a26.749 26.749 0 0 0 0 7.956L20 14.323zm1 4H3a2.5 2.5 0 0 0 0 5h18a2.5 2.5 0 0 0 0-5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LuggageLockers(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 3.5h1M15 5V3.5M13 5V3.5m-5 9v11m8-11v11m-1-11a3 3 0 1 0-6 0M11.599 5A2.999 2.999 0 0 1 6 3.5A3 3 0 0 1 11.599 2h4.813c.347.492.815.885 1.36 1.142L18 3.25v.5l-.229.108A3.486 3.486 0 0 0 16.411 5zM20 23.5H4v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L4 12.677V12.5h16v.177l-.203 1.345a26.747 26.747 0 0 0 0 7.956L20 23.323z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LuggageTrolley(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.5 5v11m8-11v11m-1-11a3 3 0 1 0-6 0m-8 16.5a2 2 0 1 0 4 0m11 0a2 2 0 1 0 4 0M24 19H4.5v-3.249c0-3.165-.058-6.353-.795-9.432C3.035 3.52 1.877.5 0 .5M23.5 16h-16v-.177l.202-1.345a26.737 26.737 0 0 0 0-7.956L7.5 5.176V5h16v.176l-.203 1.346a26.749 26.749 0 0 0 0 7.956l.203 1.345z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5 6.5l11.375 7h.25l11.375-7m0-2.5v16.5H23c-3-.5-8-.75-11-.75S4 20 1 20.5H.5V4c3-.5 8.5-.75 11.5-.75s8.5.25 11.5.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mailbox(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 21.5v-7.174C11 11.27 9.058 8.5 6 8.5s-5 2.769-5 5.826V21.5zm0 0h.5l11.5-2v-7.541c0-2.845-2.35-4.816-5.163-4.934a3.554 3.554 0 0 0-.534.023L16 7.192m-2 .227c-2.382.275-6.484.8-8.6 1.118M14 16.5V4.75m0 0v-3L19 1s.5 1.5 0 3zM3 14s1.5.5 3 .5s3-.5 3-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaleSign(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M23.497.5L13.302 10.695M23.497.5c-.676.676-1.924 1.11-3.04 1.379c-1.489.359-3.035.424-4.558.252c-1.183-.134-2.485-.4-3.009-.924M23.497.5c-.676.676-1.11 1.923-1.38 3.039c-.358 1.49-.424 3.036-.251 4.559c.134 1.182.4 2.484.924 3.009M8 23.5a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15.5 21.5v-17m0 17h-.333l-.358-.22A12 12 0 0 0 8.52 19.5H8.5m7 2h.177a12 12 0 0 0 6.173-1.71l.65-.39V2.5h-.25l-.357.22a12 12 0 0 1-6.29 1.78h-.353l-.483-.29A12 12 0 0 0 8.593 2.5H8.5m0 0h-.176A12 12 0 0 0 2.15 4.21l-.65.39v16.9h.25l.357-.22a12 12 0 0 1 6.29-1.78m.103-17v17m0 0h-.104m0 0H8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Massage(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17.5 24v-5.5a3 3 0 0 0-3-3C13.5 15.5.1 17 .1 17M15 21.5H0m24 0h-4m-10.5-8l2.918-2.085a1.39 1.39 0 0 0 .47-1.677l-.985-2.297C11.534 7.226 10.13 6.5 8 6.5s-3.534.726-3.903.94l-.985 2.298a1.39 1.39 0 0 0 .47 1.678L6.5 13.5m13 3.9s1 1.6 2.25 1.6a1.75 1.75 0 0 0 1.75-1.75c0-.966-.784-1.746-1.75-1.746c-1.25 0-2.25 1.596-2.25 1.596zM7.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C9.746 3.5 8.15 4.5 8.15 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalLaboratory(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 23.5h20M8.646 15a4.5 4.5 0 1 0 4.262-7.408M5.124 15a7.519 7.519 0 0 0 4.84 4.22M6.5 9.5L8 11m-6 2.5h9m4.258-8.258l.712-.712a1.81 1.81 0 0 0 0-2.56l-.94-.94a1.81 1.81 0 0 0-2.56 0L7 6.5l3.5 3.5l2.408-2.408m2.35-2.35a7.503 7.503 0 0 1-1.222 13.978m1.222-13.978l-2.35 2.35M9.964 19.22a3.235 3.235 0 0 1 4.072 0m-4.072 0a3.233 3.233 0 0 0-.858 1.069L8.5 21.5s-1 2-3 2h13c-2 0-3-2-3-2l-.606-1.211a3.234 3.234 0 0 0-.858-1.069M15.5 1.5L16.6.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MeetingPoint(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24m-6.5 0v-6.5a2 2 0 0 0-2-2v-8s1.5-1 4-1m15 17.5v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1m-6.65-2s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25zm-6.5 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.246 3.5 5.65 4.5 5.65 4.5zm13 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MeetingPointTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8 8L1 1m7 7c-.279-.279-.458-.793-.569-1.253a5.422 5.422 0 0 1-.104-1.88c.056-.489.165-1.026.381-1.242M8 8c-.279-.279-.793-.458-1.253-.569a5.421 5.421 0 0 0-1.88-.104c-.489.056-1.026.165-1.242.381M8 16l-7 7m7-7c-.279.279-.458.793-.569 1.253a5.422 5.422 0 0 0-.104 1.88c.056.489.165 1.026.381 1.242M8 16c-.279.279-.793.458-1.253.569a5.422 5.422 0 0 1-1.88.104c-.489-.056-1.026-.165-1.242-.381M16 8l7-7m-7 7c.279-.279.458-.793.569-1.253a5.422 5.422 0 0 0 .104-1.88c-.056-.489-.165-1.026-.381-1.242M16 8c.279-.279.793-.458 1.253-.569a5.42 5.42 0 0 1 1.88-.104c.489.056 1.026.165 1.242.381M16 16l7 7m-7-7c.279.279.458.793.569 1.253a5.43 5.43 0 0 1 .104 1.88c-.056.489-.165 1.026-.381 1.242M16 16c.279.279.793.458 1.253.569a5.422 5.422 0 0 0 1.88.104c.489-.056 1.026-.165 1.242-.381M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MeetingRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 13.5h14m-7 0V24M6.5 11V6.5H5.328a3 3 0 0 0-2.906 2.255L.5 16.25v.25h7V18c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5M17.5 11V6.5h1.172a3 3 0 0 1 2.906 2.255L23.5 16.25v.25h-7V18c0 1.5 0 2.5-.75 4c0 0-.75 1.5-1.75 1.5m-7.65-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C8.246 3.5 6.65 4.5 6.65 4.5zm11.3 0s1.6-1 1.6-2.25A1.75 1.75 0 0 0 17.5.5c-.966 0-1.746.784-1.746 1.75c0 1.25 1.596 2.25 1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenRestroom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24m1.85-19.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 12h2m12 0h2M.5 16.5V20H1a43.131 43.131 0 0 1 3-.37M.5 16.5H1a3 3 0 0 1 3 3v.13M.5 16.5v-9M4 19.63a94.725 94.725 0 0 1 8-.38c2.134 0 5.281.127 8 .38m0 0a43.03 43.03 0 0 1 3 .37h.5v-3.5M20 19.63v-.13a3 3 0 0 1 3-3h.5m0 0v-9m0 0V4H23a43.03 43.03 0 0 1-3 .37m3.5 3.13H23a3 3 0 0 1-3-3v-.13m0 0a94.72 94.72 0 0 1-8 .38a94.72 94.72 0 0 1-8-.38m0 0A43.178 43.178 0 0 1 1 4H.5v3.5M4 4.37v.13a3 3 0 0 1-3 3H.5m11.5 7a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreAddPlus(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 9.5v-9h.25l.27.054a10.1 10.1 0 0 0 3.96 0L14.25.5h.25v9h9v.25l-.054.27a10.101 10.101 0 0 0 0 3.96l.054.27v.25h-9v9h-.25l-.27-.054a10.101 10.101 0 0 0-3.96 0l-.27.054H9.5v-9h-9v-.25l.054-.27a10.1 10.1 0 0 0 0-3.96L.5 9.75V9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Motorcycle(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 .5h3.5V1A2.5 2.5 0 0 0 10 3.5m10-3h-3.5V1A2.5 2.5 0 0 1 14 3.5m-4 3H4.5v.25l.365.638A20 20 0 0 1 7.5 17.311V19.5h3m3.5-13h5.5v.25l-.365.639A20 20 0 0 0 16.5 17.31v2.19h-3m-3-5V22a1.5 1.5 0 0 0 3 0v-7.5zm1.5-7a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MriPet(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.765 7A10.498 10.498 0 0 1 13 1.5c5.799 0 10.5 4.701 10.5 10.5S18.799 22.5 13 22.5A10.503 10.503 0 0 1 3.289 16M18 16H3.289M0 16h3.289M2 13h9.5v-1a3 3 0 0 0-3-3c-.54 0-4.692.437-8.4.84m13.4 1.56s1 1.6 2.25 1.6a1.75 1.75 0 0 0 1.75-1.75c0-.966-.784-1.746-1.75-1.746c-1.25 0-2.25 1.596-2.25 1.596z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Museum(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M19.5 10.5V19m-10-8.5V19m-5-8.5V19m10-8.5V19M2 21h20M0 23.5h24m-.5-15.75v.75H.5v-.75C5 6 9.186 3.577 11.438.875L11.75.5h.5l.312.375C14.814 3.577 19 6 23.5 7.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 19a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm0 0V5.5a25.49 25.49 0 0 0 12.817-3.457l.933-.543h.25v3.85m0 11.65a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm0 0V5.35M8.5 9.5a25.49 25.49 0 0 0 12.817-3.457l.933-.543l.25-.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MuteNotification(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 20a2.5 2.5 0 0 1-5 0M.5.5l23 23m-6-6h5v-.25l-.848-.908a12 12 0 0 1-3.134-6.7l-.336-2.684A6.23 6.23 0 0 0 6.011 6.01M14 17.5H1.5v-.25l.848-.908a12 12 0 0 0 3.134-6.7l.074-.586"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newspapers(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18.91 20c.276-1.823.59-4.637.59-8c0-6-1-10.25-1-10.25l-.05-.25H1.5v.25S2.5 6 2.5 12s-1 10.25-1 10.25v.25h19.95l.05-.25s1-4.25 1-10.25c0-3.07-.262-5.681-.517-7.5H19m-2.685 12h-11m-.332-12a54.458 54.458 0 0 1 .496 9h11a54.453 54.453 0 0 0-.496-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoAccessForServiceAnimalFour(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l15.848 15.848M23.5 23.5l-7.152-7.152M11.5 18a5.5 5.5 0 1 0-2 4.243m2-10.743v-5h-1.172a3 3 0 0 0-2.906 2.255l-.966 3.764M21 21v-6.357a2.5 2.5 0 0 0 2.5-2.5v-.357h-3.786c-.986 0-1.785-.8-1.785-1.786v2.782a4.815 4.815 0 0 1-1.581 3.566M11.5 20v1.75c0 .966.784 1.75 1.75 1.75h.393v-2.71c0-1.264.496-2.462 1.357-3.35M16 20h1.5v.292c0 .875 0 1.458.45 2.333c.001.002.45.875 1.05.875m-7.65-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoAccessForServiceAnimalThree(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l23 23M7.052 17s-2.708 0-4.393 2.973C2.125 20.915 2 22.02 2 23.101V24v-4.995a7.464 7.464 0 0 1 6.444-7.394L9 11.534M7.052 17l3.657.671A20 20 0 0 0 14.32 18h.93l.25.25v2.042c0 .875 0 1.458.45 2.333c0 0 .45.875 1.05.875M7.052 17s-1.266 1.224-1.88 3.648c-.137.54-.172 1.099-.172 1.656V24M.5 8v3.5a3 3 0 0 0 3 3h.013M19.5 14V9.5h1a3 3 0 0 0 3-3v-1H19A2.5 2.5 0 0 1 16.5 3a8.597 8.597 0 0 1-2.992 6.523M19.5 14v5.5m0-5.5a6.243 6.243 0 0 1-5.923-4.27l-.07-.207M8.5 0c0 5.5 2.5 8.5 5 9.5l.008.023m0 0A8.583 8.583 0 0 1 11.068 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoAccessForServiceAnimalTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l14.749 14.749M23.5 23.5l-8.251-8.251M19 24v-9.357a2.5 2.5 0 0 0 2.5-2.5v-.357h-3.786c-.986 0-1.785-.8-1.785-1.786v2.782c0 .882-.242 1.732-.68 2.467M9.5 20v1.75c0 .966.784 1.75 1.75 1.75h.393v-2.71a4.81 4.81 0 0 1 2.143-4.004c.073-.049.144-.1.214-.152M14 20h1.5v.292c0 .875 0 1.458.45 2.333c.001.002.45.875 1.05.875M5.5 24v-6.5a2 2 0 0 0-2-2v-8s1.5-1 4-1s4 1 4 1v4m-4.15-7s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C9.246 3.5 7.65 4.5 7.65 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoAlcohol(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 16v7.5H20m1.5-2v-8l-1.8-2.4a6 6 0 0 1-1.2-3.6v-7h-3v7a6 6 0 0 1-1.2 3.6L13 12.833l-.083.084M5 19.5a3.5 3.5 0 0 1-3.394-4.364L2.5 11.5h5l.894 3.636A3.5 3.5 0 0 1 5 19.5Zm0 0v4m0 0h3m-3 0H2M.5.5l12.417 12.417M23.5 23.5L12.917 12.917"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoDrinkingWater(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 23.5H5v.5h.5zM3.74 7.8l.488-.111zm.247-1.161l-.111-.488l-.975.223l.112.487zM18.5 23.5v.5h.5v-.5zm1.76-15.7l-.488-.111zm.24-1.05l.487.111l.013-.055V6.75zm0-.25h.5V6h-.5zM10.135 3.396l-.362-.346zm1.04-1.089l.361.346zm1.65 0l-.362.346zm1.04 1.089l.361-.346zM18 13v-.5zM6 13v.5zm7-.17v.5h.086l.08-.03zM6 23.5v-.102H5v.102zM4.228 7.689l-.24-1.05l-.975.222l.24 1.05zM5.5 24h13v-1h-13zm13.5-.5v-.102h-1v.102zm1.747-15.588l.24-1.05l-.974-.223l-.24 1.05zM21 6.75V6.5h-1v.25zM17 7h3.5V6H17zM6 23.398A70.5 70.5 0 0 0 4.228 7.69l-.975.223A69.5 69.5 0 0 1 5 23.398zM9 5c0 .864.29 1.622.834 2.166C10.378 7.71 11.136 8 12 8V7c-.636 0-1.128-.21-1.46-.541C10.21 6.128 10 5.636 10 5zm.773-1.95A2.824 2.824 0 0 0 9 5h1c0-.476.18-.929.496-1.26zm1.04-1.088L9.773 3.05l.723.69l1.04-1.087zM11.5 0c0 .693-.271 1.527-.687 1.962l.723.69c.642-.67.964-1.781.964-2.652zm1.686 1.962C12.771 1.527 12.5.692 12.5 0h-1c0 .87.322 1.982.963 2.653zm1.04 1.088l-1.04-1.088l-.723.69l1.04 1.089zM15 5c0-.727-.276-1.43-.774-1.95l-.723.69c.316.331.497.784.497 1.26zm-3 3c.863 0 1.622-.29 2.166-.834C14.71 6.622 15 5.864 15 5h-1c0 .636-.21 1.128-.541 1.459c-.331.331-.823.541-1.46.541zm6 4.5a2.5 2.5 0 0 1-2.5-2.5h-1a3.5 3.5 0 0 0 3.5 3.5zm-12 0H4.723v1H6zm12 1h1.277v-1H18zm-4.833-.2A3.501 3.501 0 0 0 15.5 10h-1a2.501 2.501 0 0 1-1.667 2.358zM6 13.5a3.499 3.499 0 0 0 3.031-1.75l-.865-.5A2.499 2.499 0 0 1 6 12.5zm12.307 5.514l4.84 4.84l.707-.708l-4.84-4.839zm.852-.32a69.501 69.501 0 0 1 1.588-10.782l-.974-.223a70.491 70.491 0 0 0-1.611 10.938zM19 23.398c0-.628.009-1.257.026-1.884l-1-.028A70.406 70.406 0 0 0 18 23.398zM.146.854l12.33 12.329l.707-.707L.853.146zm12.33 12.329l5.831 5.831l.707-.707l-5.831-5.831zm.524-.854h-.17v1H13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoDroneAllowed(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 17.5H9s-3-6-8.5-6v-4H5m11 8.42c1.302-1.779 3.85-4.42 7.5-4.42v-4h-16m3-4.5l-1.03.348A12.409 12.409 0 0 1 4 3.909M23.5 3l-1.03.348A12.41 12.41 0 0 1 18.5 4m-5-1l1.03.348C15.81 3.78 17.15 4 18.5 4m0 0v3.5M5 15.5C3 15.5.5 18 .5 21M19 15.5c2 0 4.5 2.5 4.5 5.5M.5.5l10.358 10.358M23.5 23.5l-9.858-9.858m-2.784-2.784a2 2 0 0 1 2.784 2.784m-2.784-2.784l2.784 2.784"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoDrugOrSubstance(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m13.251 16.5l-4.125 4.125s-1.438-.479-3.355-2.395m0 0c-1.916-1.917-2.395-3.355-2.395-3.355L7.5 10.75m-1.73 7.48L.5 23.5m14.376-8.624l5.75-5.75l-5.75-5.75l-5.75 5.75m13.416 1.916l-9.583-9.584M17.75 6.25l3.833-3.833M23.5 4.333L19.667.5m-7.666 5.75l1.916 1.917M6.25 12l1.917 1.917M.5.5l23 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoElectricScooter(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l16.169 16.169M23.5 23.5l-2.05-2.05M3 14a5 5 0 0 1 5 5v.5h.25S10 19 12 19s3.75.5 3.75.5M19 14.416a5.02 5.02 0 0 0-2.331 2.084v.169M13 2.5h3.5v.825c0 4.705 1.205 9.331 3.5 13.438m1.45 4.687h.05a2.5 2.5 0 1 0-2.95-2.95v.05m2.9 2.9l-2.9-2.9m0 0l-1.881-1.881M5.5 19a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoEntryForPedestrians(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l8.903 8.903M23.5 23.5L13.118 13.118M16 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-3.5-.75-5v-.5l-5-2.5v-.25l.88-4.25l.023-.097m6.097.785C17 11.5 19 12 21 12m-10.5 5.5c-1 3-3 5-5.5 5m.155-9.5l.535-2.583A4.919 4.919 0 0 1 6.22 9M7.5 7.523a4.902 4.902 0 0 1 3-1.023h1.544a2 2 0 0 1 1.958 2.405L13.155 13l-.037.118M9.403 9.403l3.715 3.715M13.35 4.5s-1.6-1-1.6-2.25a1.746 1.746 0 1 1 3.495 0c0 1.25-1.595 2.25-1.595 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoFirearmWeapon(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l13.356 13.356M23.5 23.5l-9.644-9.644M14 14H8.75M4 4h19s.5 1 .5 2.5S23 9 23 9h-8v.5c0 2.463-.843 3.915-1.144 4.356m-11.46-7.96C1.638 7.465.5 8.572.5 8.572v.374c1 .553 2 1.553 2 1.553v.25S1 16.161 1 20h7c0-4.988 1.75-9.5 1.75-9.5h5.202"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoFishingOne(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M21.6 12.5s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75m-.15 4h.3m-.3 0l-1.227 2m1.377-6c.966 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25m-.15-4v-8h-.25c-4.41 1.29-9.002 5.445-11.155 9.845M21.9 12.5l1.223 2M0 19.5h7m.5-5.5V7.5m4.5 16s-2.5 0-2.5-5.5v-1.5h-8v-.25l1.922-7.495c.11-.43.31-.82.578-1.148M.5.5l9.845 9.845M23.5 23.5l-3.484-3.484m3.234-1.766s-2.2-1.506-5.212-.25l-.026.012m-6.262.238s1.485 1.017 3.686.683m7.814 1.817s-1.284-.88-3.234-.734m-8.266.734s2.365 1.62 5.55.099m-6.955-10.504l7.667 7.667m2.004 2.004l-2.004-2.004M7.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C9.246 3.5 7.65 4.5 7.65 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoFood(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 6.5h16v.25l-.168.434c-1.425 3.681-2.322 8.12-2.668 12.316l-.01.155M14 23.5h5.5v-.155c0-.442.007-.89.021-1.345M17.5 6.5V6A5.5 5.5 0 0 1 23 .5M.5.5l19.155 19.155M23.5 23.5l-3.845-3.845M1.5 17.5v-1a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v1zm.5 0h9a1.5 1.5 0 0 1 0 3H2a1.5 1.5 0 0 1 0-3Zm-.5 3v3h10v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoLitterTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m5.18 13.86l.416.416a2.5 2.5 0 0 0 3.536 0l7.313-7.314l.245.245l-.048 3.256a2.5 2.5 0 0 0 2.5 2.537h.462V0M4.876 10.046a2.5 2.5 0 0 1-3.522-.012L1 9.68l2.895-2.895m.981 3.261l.013-.012l1.127-1.127zm0 0l-3.169 3.17l.326.325a2.5 2.5 0 0 0 3.535 0L8.11 11m1.446-1.445l1.67-1.67M7.461 7.46l1.67-1.67M5.34 5.34L10.18.5M.5.5l23 23m-7.25-2.75a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoLuggageTrolleysBeyondThisPoint(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l23 23M11 7v4m7-4v9m-6-9v-.5a2.5 2.5 0 0 1 5 0V7M7 7h14.5v.176l-.08.369a18.801 18.801 0 0 0 0 7.91l.08.368V16h-14v-.177l.08-.368c.279-1.3.42-2.625.42-3.955v-1m-3.5 11a2 2 0 1 0 4 0m9 0a2 2 0 0 0 3.323 1.5M22 19H4.5v-.497c0-4.646.211-9.938-3.047-13.25C.998 4.79.509 4.5 0 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoNoise(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17 6.5a5.5 5.5 0 0 1 .5 10.978M17 9.5a2.5 2.5 0 0 1 0 5m-2.5 0v-12h-.25l-.523.548a15.999 15.999 0 0 1-6.483 4.12l-.06.015M14.5 18v3.5h-.25l-.523-.548A16 16 0 0 0 2.154 16H1.5V8h.654c.79 0 1.574-.058 2.346-.173M.5.5l6.683 6.683M23.5 23.5L7.183 7.183"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoParking(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 12.5h1a3.5 3.5 0 1 0 0-7H9V9m0 10v-7M3.868 3.868l16.264 16.264M12 23.5C5.649 23.5.5 18.351.5 12S5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoPetsAllowed(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l9.082 9.082M23.5 23.5L9.582 9.582M16 22v-3m0-3V7.5h1A2.5 2.5 0 0 0 19.5 5V3.5H14A2.5 2.5 0 0 1 11.5 1v3.894A6.74 6.74 0 0 1 9.68 9.5l-.098.082M7.5 11.315a6.737 6.737 0 0 0-2 4.79V21.5H5A2.5 2.5 0 0 1 2.5 19v-4m5.5.5h2.5v.5c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoPictures(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l9.857 9.857M23.5 23.5l-6.857-6.857M9.5 13a4.5 4.5 0 0 0 4.5 4.5m-11-8h3m-3-3H.5v14H17M6.5 6.5h1a3 3 0 0 0 3-3h7a3 3 0 0 0 3 3h3v11a3 3 0 0 1-3 3M10.357 10.357a4.5 4.5 0 1 1 6.285 6.285m-6.285-6.285l6.286 6.286"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoRunning(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5l-5-3.25V13l1.077-2M5 20c0-1 1.5-1.75 1.5-1.75c1.327-.664 2.263-.74 3.5-.749M20 11.5h-3a3 3 0 0 1-3-3v-2h-1.208a3 3 0 0 0-2.642 1.578l-.725 1.347M.5.5l8.925 8.925M23.5 23.5L9.425 9.425M13.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSelfieStickAllowed(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l23 23m-23 .4l.649-4.802A3 3 0 0 1 4.122 16.5H5.5l2.353 4.706A2.342 2.342 0 0 0 9.947 22.5c.364 0 .722-.084 1.024-.287c.709-.477 2.119-1.598 3.529-3.713M16 16l4-10.5m3.5-3.143V6.5h-.25L16.5 4.643V.5h.25zM4.547 14.507s-2.434-.091-3.242-1.527c-.626-1.11-.258-2.527.821-3.166c1.08-.638 2.457-.253 3.082.857c.809 1.435-.326 3.638-.326 3.638z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSmoking(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 11.5h-8v5h13m-2-5h9v5h-4m7-5.5v6m-9-16v1A6.5 6.5 0 0 0 21 8.5m2.5.5V5.5H21A3.5 3.5 0 0 1 17.5 2V1M.5.5l23 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoTouch(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 9v15M9 9V3.5a3 3 0 0 1 3-3h.5v12v-7a3 3 0 0 1 3-3h.5V16v-3.5a3 3 0 0 1 3-3h.5v10M9.17 2.5H8.5a3 3 0 0 0-3 3m-5-5l23 23m-5.924-2.424A3.497 3.497 0 0 0 16 24m-3.5-4v-1c0-.84.295-1.61.788-2.212"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoVideo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17.5 17.5v-3h.054c.94 0 1.877.274 2.564.917c.814.761 1.51 1.644 2.061 2.624L23 19.5h.5v-15H23l-.82 1.459a10.929 10.929 0 0 1-2.062 2.624c-.687.643-1.624.917-2.564.917H17.5v-5h-13m11.5 15H.5V4M.5.5l23 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoWheelchairAccess(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l9.582 9.582M23.5 23.5L10.082 10.082M14.5 18A5.5 5.5 0 1 1 9 12.5m5.5-.5V6.5h-1.172a3 3 0 0 0-2.906 2.255l-.34 1.327M17.5 17.5v.5c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5m-5.65-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nursery(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9 14.5h6m-6 0L5.5 18v.25c1.341 1.341 2.283 2.483 2.825 3.782c.127.306.175.637.175.968m.5-8.5v.5a3 3 0 1 0 6 0v-.5m-6 0v-7m6 7l3.5 3.5v.25c-1.341 1.341-2.283 2.483-2.825 3.782A2.508 2.508 0 0 0 15.5 23m-.5-8.5v-7M3.5 3C3.5 4 5 8 12 8s8.5-4 8.5-5m-8.666 3S10 4.906 10 3.5c0-1.087.915-2 2.002-2A2.02 2.02 0 0 1 14 3.5C14 4.906 12.171 6 12.171 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Office(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M24 6.5c-.329 0-.659.07-.937.245C21.5 7.733 21.5 9.08 21.5 10v.5H17m-8 3h15m-7.5 0V24m-4.5-.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-8v-.25l1.922-7.495A3 3 0 0 1 6.328 6.5H7.5l1.447 2.894a2 2 0 0 0 1.79 1.106H15m-7.65-6s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C9.246 3.5 7.65 4.5 7.65 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OfficePod(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 15V9.5h-.676a3 3 0 0 0-2.905 2.25L3.5 17.25v.25H11v1.456c0 1.24 0 2.044.554 3.305c0 0 .62 1.239 1.446 1.239m-4 0H3.5a3 3 0 0 1-3-3v-17a3 3 0 0 1 3-3h17a3 3 0 0 1 3 3v17a3 3 0 0 1-3 3h-4v-9m-4.5 0h9m-12.613-7s-1.201-.75-1.201-1.687a1.313 1.313 0 0 1 2.625 0C9.81 6.75 8.613 7.5 8.613 7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paper(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.5 1.5v7h7m-16-7v21h16V8l-.282-.126a12 12 0 0 1-6.092-6.092L14 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParentAndInfantPrioritySeating(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m16 14.5l-2.232-2.511A7 7 0 0 1 12 7.339v-.84h-1.515A2.735 2.735 0 0 0 7.75 9.236a9.118 9.118 0 0 1-1.06 4.266L5.5 15.75V16l6.5 1s0 4.5 2 6.39M13.5 8h4.75v2.117c0 1.113-.46 2.145-1.227 2.883M8 24c0-1.613.229-3.323 1-4.68M10.6 4.5S9 3.5 9 2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25zm5.273 1.567S14.5 5.192 14.5 4.098c0-.845.672-1.53 1.502-1.53s1.498.685 1.498 1.53c0 1.094-1.37 1.97-1.37 1.97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Park(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 19.56v4.38M4.5 12.56h3.724A33.264 33.264 0 0 1 1.5 19.31v.25h21v-.25a33.263 33.263 0 0 1-6.724-6.75H19.5v-.25l-1.386-1.04A15.383 15.383 0 0 1 12 .06a15.384 15.384 0 0 1-6.114 11.21L4.5 12.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parking(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9 19v-6.5m0 0v-7h4.5a3.5 3.5 0 1 1 0 7zm3 11C5.649 23.5.5 18.351.5 12S5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Passports(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18.5 5.5h2v18h-17v-20c8.5-1 14.75-3 14.75-3h.25zm0 0H6m10.5 9A4.5 4.5 0 0 0 12 10m4.5 4.5A4.5 4.5 0 0 1 12 19m4.5-4.5c-.5.5-2 1-4.5 1s-4-.5-4.5-1M12 10a4.5 4.5 0 0 0-4.5 4.5M12 10c.966 0 1.75 2.015 1.75 4.5S12.966 19 12 19m0-9c-.966 0-1.75 2.015-1.75 4.5S11.034 19 12 19m-4.5-4.5A4.5 4.5 0 0 0 12 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Patio(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 19.5h10m.5-5l-2-8H7.067a3 3 0 0 0-2.803 1.932L1 17m14 6.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-7v-6M15.069 0c0 .33-.05.664-.206.955A3.811 3.811 0 0 1 11.5 2.969a3.814 3.814 0 0 1 3.569 2.47a3.814 3.814 0 0 1-.778 4.27a3.814 3.814 0 0 1 4.27-.778a3.814 3.814 0 0 1 2.47 3.569c0-1.455.815-2.72 2.014-3.363c.291-.156.624-.206.955-.206M19 .526A3.177 3.177 0 1 0 23.474 5m-16.28-.5s-1.81-.557-2.134-1.776A1.768 1.768 0 0 1 6.302.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pedestrians(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 6.525L8.5 13.75V14l5 2.5v.5c0 1.5 0 3.5.75 5c0 0 .75 1.5 1.75 1.5m-3-9.75l1.002-4.845A2 2 0 0 0 12.044 6.5H10.5a4.912 4.912 0 0 0-4.81 3.917L5.154 13M15.5 10.188C17 11.5 19 12 21 12m-10.5 5.5c-1 3-3 5-5.5 5m8.35-18s-1.6-1-1.6-2.25a1.746 1.746 0 1 1 3.495 0c0 1.25-1.595 2.25-1.595 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 22.5h13M3.5 16L18 1.5A4.5 4.5 0 0 1 22.5 6L8 20.5H7c-1.974 0-3.377.584-5.02 1.68l-.48.32l.32-.48C2.917 20.376 3.5 18.973 3.5 17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonalTraining(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17.5 13.599a13.28 13.28 0 0 0-1-.099a3 3 0 0 0-3 3v1h4.677a1.323 1.323 0 0 0 1.276-1.67L18 10.436M7 20.5h17m-17 3c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1l-4-3.25v-.25l1.789-5.009A3 3 0 0 1 5.114 6.5H6.5v3.446a1.554 1.554 0 0 0 2.378 1.318L12.5 9M.5 23.75S2 21 2 17.5m21.85-3.118a399.79 399.79 0 0 1-3.612-.439zM11.5 15.9s-1 1.6-2.25 1.6a1.75 1.75 0 0 1-1.75-1.75c0-.966.783-1.746 1.75-1.746c1.25 0 2.25 1.596 2.25 1.596zm6-5.4a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-10.4-6s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C8.996 3.5 7.4 4.5 7.4 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PetsAllowed(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M16 22V7.5h1A2.5 2.5 0 0 0 19.5 5V3.5H14A2.5 2.5 0 0 1 11.5 1v3.894a6.737 6.737 0 0 1-3 5.606a6.737 6.737 0 0 0-3 5.606V21.5H5A2.5 2.5 0 0 1 2.5 19v-4m5.5.5h2.5v.5c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pharmacy(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M.5 6.5V10c0 4.56 2.5 8.5 7 10.36v.14l-.123.123a12 12 0 0 1-3.119 2.248l-.758.379v.25h17v-.25l-.758-.38a12 12 0 0 1-3.119-2.247L16.5 20.5v-.14c4.5-1.86 7-5.8 7-10.36V6.5z"/><path d="M10.5 12.5v-3h3v3h3v3h-3v3h-3v-3h-3v-3zm9-12l-6 6h4l4-4V2A1.5 1.5 0 0 0 20 .5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M2 4.706C2 14.257 9.743 22 19.294 22L23 18.294l-4.323-4.323l-3.089 3.088l-8.647-8.647l3.088-3.088L5.706 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhysicalTherapy(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 22.5h24m-13-10l-1.036-1.036A5 5 0 0 1 8.5 7.93v-.697c.635-.296 1.85-.732 3.5-.732s2.865.436 3.5.732v.697a5 5 0 0 0 1.465 3.535L18 12.5m-2 7H6.5v-1a3 3 0 0 1 3-3c.472 0 3.713.335 7 .688c3.67.395 7-4.188 7-4.188m-19 5.9s-1 1.6-2.25 1.6A1.75 1.75 0 0 1 .5 17.75c0-.966.784-1.746 1.75-1.746c1.25 0 2.25 1.596 2.25 1.596zm7.352-13.4s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75c.967 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pilates(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 20.5h24m-10-11H3.5a2 2 0 0 0-1.831 1.195c-.214.485.042 1.016.417 1.39l4.53 4.531a3.018 3.018 0 0 0 4.268 0L15 12.5h2s3.036.174 6.5-1.826M10 14L5.5 9.5m-2.805-2S.885 6.943.56 5.724a1.768 1.768 0 0 1 1.242-2.163a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m1.204 4.056l1.833-1.834S9.759 4.463 13.63 6.704l4.31-4.31c.643-.643 1.461-1.089 2.363-1.212c.688-.094 1.532-.182 2.29-.182l.407.407c0 .757-.088 1.602-.182 2.29c-.123.902-.569 1.72-1.212 2.364l-4.31 4.31c2.241 3.87 4.482 10.592 4.482 10.592l-1.834 1.833s-3.666-5.296-6.722-8.352c0 0-2.264 1.45-5.703 4.89c0 0 .407 1.629.407 3.259L7.519 23s-1.63-2.037-3.056-3.463C3.037 18.111 1 16.482 1 16.482l.407-.408c1.63 0 3.26.407 3.26.407c3.439-3.439 4.889-5.703 4.889-5.703C6.5 7.722 1.204 4.056 1.204 4.056Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plastic(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 .5h4M6.5 10H7l1.027.171a24.168 24.168 0 0 0 7.946 0L17 10h.5m-11 7H7l1.027-.171a24.168 24.168 0 0 1 7.946 0L17 17h.5m-4-13.5h-3a7.236 7.236 0 0 0-4 6.472V21.5a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V9.972a7.236 7.236 0 0 0-4-6.472Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m2 22.5l1 .5c7.5-5 12.5-7.75 19-10.25v-1.5C15.5 8.75 10.5 6 3 1l-1 .5s.5 6.2.5 10.5S2 22.5 2 22.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 8v5m2.5-2.5H5m-.909 10H5.5s1.5-5 6.5-5s6.5 5 6.5 5h1.409a3.591 3.591 0 0 0 3.499-4.398L20.5 3.5h-.25S17 4.5 12 4.5s-8.25-1-8.25-1H3.5L.592 16.102A3.59 3.59 0 0 0 4.091 20.5ZM16.5 8.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playground(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 23.5c3-5 3-14.5 3-20.578V.5h17v2.422c0 6.078 0 15.578 3 20.578M9.5.5v15m5-15v15m-8 0v3h11v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Police(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M11.75 23.5C4 19 2.5 16 2.5 5.5c3.15 0 6.356-1.238 8.276-3.357c.422-.465.687-1.044.874-1.643h.7c.187.599.452 1.178.874 1.643C15.144 4.262 18.35 5.5 21.5 5.5c0 10.5-1.5 13.5-9.25 18z"/><path d="M11.898 7.5h.204l.15.542a4 4 0 0 0 3.856 2.935h.392v.166l-.365.252a4 4 0 0 0-1.55 4.473l.196.632h-.21a4.066 4.066 0 0 0-5.142 0h-.21l.195-.632a4 4 0 0 0-1.55-4.473l-.364-.252v-.166h.392a4 4 0 0 0 3.856-2.935z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Port(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 4.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 0v19M7 15c-2 0-3.5-1-5-3h-.5c0 4.68 3.06 8.643 7.29 10c1.21.4 2.21.5 3.085 1.5h.25C13 22.5 14 22.4 15.211 22c4.228-1.357 7.289-5.32 7.289-10H22c-1.5 2-3 3-5 3M7 7.5s2.5 1 5 1s5-1 5-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrayingRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m17 9l-4.25 4.5h-.25l-1.5-6h-.5a3 3 0 0 0-3 3V17c5.5 0 9 2.5 9 3.5H7m6-10.882L15 7.5m-15 16h24m-12.445-18s1.81-.557 2.135-1.776a1.768 1.768 0 0 0-1.242-2.163a1.75 1.75 0 0 0-2.145 1.25c-.325 1.219.961 2.61.961 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PregnantWomanPrioritySeating(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 23.5c-2.5-2-2-5.5-2-5.5l-6.5-1v-.25l1.343-2.552A10.053 10.053 0 0 0 9 9.516A3.016 3.016 0 0 1 12.016 6.5H13.5v3l.546.137a3.24 3.24 0 0 1 2.454 3.142c0 1.13.283 2.24.824 3.232L18 17.25v.25l-3 .5m-5 6c0-1.613.229-3.323 1-4.68m.85-14.82s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 13.5c0 4.5-2 7.5-2 7.5v.5h17V21s-2-3-2-7.5m-13-6v-5h13v5M3 13.5h18m-18 4H.5v-10h23v10H21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProjectionRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 20.5h5.5M5 15v-5h-.557a3 3 0 0 0-2.938 2.392L.5 17.25v.25h5m1.5 3h5.5M12 15v-5h-.557a3 3 0 0 0-2.938 2.392L7.5 17.25v.25H14v.5c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5m-2-9h9v-.25S23 12 23 7.5s.5-6.75.5-6.75V.5h-19V3m.382 5.249S3.623 7.436 3.623 6.42C3.623 5.636 4.239 5 5 5c.76 0 1.374.636 1.374 1.421c0 1.015-1.256 1.828-1.256 1.828zm7 0s-1.26-.813-1.26-1.828C10.623 5.636 11.24 5 12 5c.76 0 1.374.636 1.374 1.421c0 1.015-1.256 1.828-1.256 1.828z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProstheticArm(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 18v6m4 0v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1V11m1.5 4.5H8m0 0H6.5m1.5 0V12m3.85-7.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProstheticLeg(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 24v-6.5a2 2 0 0 0-2-2v-8s1.5-1 4-1s4 1 4 1v8a2 2 0 0 0-2 2V19m-1.5 4.5H14m0 0h1.5m-1.5 0V20M11.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pull(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 5h3v.5a3 3 0 0 1-3 3H7h5.25a1.75 1.75 0 1 1 0 3.5H10h1.25a1.75 1.75 0 1 1 0 3.5H10h.25a1.75 1.75 0 1 1 0 3.5H0M8 5H3.5A3.5 3.5 0 0 1 0 8.5m9.5 0v-8h5M9.5 19v4.5h5M16 12h8m-3.429-4c0 .423-.419 1.056-.842 1.587a7.506 7.506 0 0 1-1.944 1.738c-.56.342-1.239.67-1.785.67M20.572 16c0-.423-.42-1.056-.843-1.587a7.509 7.509 0 0 0-1.944-1.738c-.56-.342-1.239-.67-1.785-.67"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Push(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M24 12h-8m3.429-4c0 .423.419 1.056.842 1.587a7.506 7.506 0 0 0 1.944 1.738c.56.342 1.239.67 1.785.67M19.43 16c0-.423.418-1.056.842-1.587a7.509 7.509 0 0 1 1.944-1.738c.56-.342 1.239-.67 1.785-.67M0 22.5h6a6.5 6.5 0 0 0 6.5-6.5V9.5H12A2.5 2.5 0 0 0 9.5 12v4V6A2.5 2.5 0 0 0 7 3.5h-.59M.5 16V3.5H1A2.5 2.5 0 0 1 3.5 6v6V1.5h.46A2.5 2.5 0 0 1 6.46 4v8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMark(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 8.5v-1a7 7 0 0 1 14 0a4.83 4.83 0 0 1-1.414 3.414l-2.414 2.414A4 4 0 0 0 14 16.157v.343h-4v-1.172c0-1.81.72-3.547 2-4.828l1.293-1.293A2.414 2.414 0 0 0 14 7.5a2 2 0 1 0-4 0v1zm7 11a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuietArea(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 24v-4a2.5 2.5 0 0 0-2.5-2.5h-.5V24m9-4v-1.5A2.5 2.5 0 0 0 18 16h-.5v4v-9.5H17a2.5 2.5 0 0 0-2.5 2.5v8m9 3v-4a2.5 2.5 0 0 0-2.5-2.5h-.708M.5 24v-2.5a4 4 0 0 1 4-4h4v-7h3v-.25l-.063-.1A19.008 19.008 0 0 1 8.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radiation(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M13.5 12a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm-13 0A11.56 11.56 0 0 1 6.249 2l3.75 6.535A3.998 3.998 0 0 0 8 12zm9.5 3.465L6.25 21.96A11.447 11.447 0 0 0 12 23.5a11.44 11.44 0 0 0 5.75-1.54L14 15.465A3.982 3.982 0 0 1 12 16a3.98 3.98 0 0 1-2-.535ZM16 12h7.5a11.56 11.56 0 0 0-5.749-10l-3.75 6.535A4.002 4.002 0 0 1 16 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radiology(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.5 16s2-1 4.5-1m0 0c2.5 0 4.5 1 4.5 1M12 15v7m-4.5-1s2-1 4.5-1s4.5 1 4.5 1m-10-2.5s2.5-1 5.5-1s5.5 1 5.5 1m1-7.5v-1S16 8.5 12 8.5S5.5 10 5.5 10v1m6.286-4.498s-2.29-1.5-2.29-3.374c0-1.45 1.121-2.625 2.505-2.625c1.383 0 2.499 1.175 2.499 2.625c0 1.875-2.284 3.374-2.284 3.374zM21.5 12.5h-19v.25S3 15 3 18s-.5 5.25-.5 5.25v.25h19v-.25S21 21 21 18s.5-5.25.5-5.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RampDown(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 23.5v-7h.5S9 21 23 23M13.57 9.782a4.424 4.424 0 1 0 3.256 4.734M13.57 9.782l1.078-1.729a3 3 0 0 1 3.322-1.31l1.042.28l-1.027 3.831M13.57 9.782a4.425 4.425 0 0 1 3.256 4.734m3.852 6.33c-1.235-1.32-1.74-2.618-1.276-4.348l.335-1.25l-2.733-.732h-.178m2.539-9.352s-1.213-1.302-.908-2.44a1.65 1.65 0 0 1 2.021-1.167a1.647 1.647 0 0 1 1.163 2.02c-.305 1.139-2.003 1.66-2.003 1.66z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RampDownArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1 23h22v-1C9 18.772 1.5 13 1.5 13H1zm22.5-8.746L11 7m12.5 7.254c-.498-.29-.944-.95-1.274-1.559a7.972 7.972 0 0 1-.856-2.622c-.106-.712-.155-1.514.068-1.902m2.062 6.083c-.498-.29-1.29-.348-1.98-.33a7.869 7.869 0 0 0-2.687.566c-.667.263-1.382.622-1.605 1.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RampUp(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M23 23.5v-7h-.5S15 21 1 23m3.28-11.366a4.424 4.424 0 1 0 5.145 2.382M4.28 11.634l.069-2.036a3 3 0 0 1 2.222-2.796l1.042-.28l1.027 3.833m-4.36 1.28a4.425 4.425 0 0 1 5.145 2.381m6.543 3.647c-1.73-.527-2.816-1.4-3.28-3.13l-.334-1.249l-2.733.732h-.196m-2.436-9.28s-1.701-.52-2.006-1.658A1.65 1.65 0 0 1 6.15 1.057a1.647 1.647 0 0 1 2.017 1.168c.305 1.138-.904 2.439-.904 2.439z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RampUpArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M23 23H1v-1c14-3.228 21.5-9 21.5-9h.5zM13 8.247L.5 15.5M13 8.246c-.498.29-.944.95-1.274 1.559a7.972 7.972 0 0 0-.856 2.622c-.106.712-.155 1.514.068 1.902M13 8.246c-.498.29-1.29.348-1.98.33a7.87 7.87 0 0 1-2.687-.566C7.666 7.747 6.95 7.388 6.728 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceptionHotelBell(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 5.5c5.344 0 8.705 2.8 9.376 7.63c.163 1.176 1.124 2.37 2.124 3.12V17H.5v-.75c1-.75 1.96-1.944 2.124-3.12C3.295 8.3 6.656 5.5 12 5.5Zm0 0v-3m-12 17h24m-14-17h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecordingStudio(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 19.5c5 0 7.5-2 7.5-2V8M12 19.5c-5 0-7.5-2-7.5-2V8M12 19.5v4m0 0H7m5 0h5m-9.5-19H10m4 0h2.5m-9 4H10m4 0h2.5m-9 4H10m4 0h2.5M12 .5c-1.838 0-3.338.27-4.5.612v14.776c1.162.342 2.662.612 4.5.612s3.338-.27 4.5-.612V1.112C15.338.77 13.838.5 12 .5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Recycling(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11 20h11.5v-.5l-3.736-6.743M11 20c1.933 0 3.5-3.5 3.5-3.5M11 20c1.933 0 3.5 3.5 3.5 3.5M17.79 11L12.25 1h-.5L8.025 7.73M17.789 11c-.966-1.674-4.781-1.281-4.781-1.281M17.789 11c-.88-1.657 1.281-4.781 1.281-4.781M7.05 9.484L1.5 19.5v.5H9M7.05 9.484c-.88 1.657 1.282 4.781 1.282 4.781M7.05 9.485c-.967 1.673-4.781 1.28-4.781 1.28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefillForWaterBottle(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 9h2v3c2.5 2 2.5 3.731 2.5 6.472V21.5a2 2 0 0 1-2 2H0M24 9h-2v3c-2.5 2-2.5 3.731-2.5 6.472V21.5a2 2 0 0 0 2 2H24M7.5 18.472V21.5a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-3.028c0-2.74 0-4.472-2.5-6.472V9h-4v3c-2.5 2-2.5 3.731-2.5 6.472ZM14 4.446c0 1.12-.895 2.054-2 2.054c-1.104 0-2-.934-2-2.054c0-.109.018-.222.048-.336c.19-.731.697-1.326 1.166-1.91l.12-.148C11.993 1.119 12 .342 12 0c0 .342.007 1.12.667 2.052l.12.148c.468.584.975 1.179 1.165 1.91c.03.114.048.227.048.336Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refrigeration(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 5.357c0-.857-.246-1.348-.503-1.767a5.634 5.634 0 0 0-1.305-1.458C9.792 1.814 9.318 1.5 9 1.5m3 3.857c0-.857.246-1.348.503-1.767a5.633 5.633 0 0 1 1.305-1.458c.4-.318.874-.632 1.192-.632m-3 3.857v3.815m0 9.47c0 .858.246 1.349.503 1.768c.343.56.792 1.049 1.305 1.458c.4.318.874.632 1.192.632m-3-3.857c0 .857-.246 1.348-.503 1.767a5.633 5.633 0 0 1-1.305 1.458c-.4.318-.874.632-1.192.632m3-3.857v-3.815m5.753-6.15c.742-.428 1.044-.887 1.279-1.319c.314-.577.512-1.21.61-1.86c.076-.503.11-1.072-.049-1.347m-1.84 4.527c.742-.429 1.29-.46 1.782-.448a5.632 5.632 0 0 1 1.915.401c.475.187.984.44 1.143.716m-4.84-.67l-3.304 1.908M6.247 15.32c-.742.429-1.044.888-1.279 1.32a5.634 5.634 0 0 0-.61 1.86c-.075.503-.11 1.072.049 1.347m1.84-4.527c-.742.429-1.29.46-1.782.448a5.632 5.632 0 0 1-1.915-.401c-.475-.187-.984-.44-1.143-.716m4.84.67l3.303-1.908m8.203 1.907c.742.429 1.29.46 1.782.448a5.632 5.632 0 0 0 1.915-.401c.475-.187.984-.44 1.143-.716m-4.84.67c.742.428 1.044.887 1.279 1.319c.314.577.512 1.21.61 1.859c.076.504.11 1.073-.049 1.348m-1.84-4.527l-3.304-1.907M6.247 8.68c-.742-.43-1.29-.46-1.782-.449a5.633 5.633 0 0 0-1.915.402c-.475.187-.984.44-1.143.716m4.84-.67c-.742-.428-1.044-.887-1.279-1.319a5.633 5.633 0 0 1-.61-1.86c-.075-.504-.11-1.072.049-1.347m1.84 4.526l3.303 1.908m0 2.828a5.465 5.465 0 0 0 0-2.828m0 2.828A5.464 5.464 0 0 1 12 14.828m0 0a5.465 5.465 0 0 1 2.45-1.414m0 0a5.466 5.466 0 0 1 0-2.828m0 0A5.464 5.464 0 0 1 12 9.172m0 0a5.465 5.465 0 0 1-2.45 1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveXcross(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 8.465L18.965 1.5l.177.177l.152.228a10.097 10.097 0 0 0 2.8 2.801l.23.153l.176.177L15.536 12l6.964 6.965l-.177.176l-.228.153a10.096 10.096 0 0 0-2.801 2.8l-.153.23l-.176.176L12 15.536L5.036 22.5l-.177-.177l-.153-.228a10.098 10.098 0 0 0-2.8-2.801l-.23-.153l-.176-.176L8.465 12L1.5 5.036l.177-.177l.229-.153a10.098 10.098 0 0 0 2.8-2.8l.153-.23l.177-.176z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Restaurant(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15.5 24V.5h.5A4.5 4.5 0 0 1 20.5 5v9.5H20a2.5 2.5 0 0 0-2.5 2.5M7 24V9.5m0 0A3.5 3.5 0 0 0 10.5 6V0M7 9.5A3.5 3.5 0 0 1 3.5 6V0M7 0v7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightAngleArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8 24c0-.741-.733-1.85-1.475-2.78c-.954-1.2-2.094-2.247-3.401-3.046C2.144 17.575.956 17 0 17m0 0c.956 0 2.145-.575 3.124-1.174c1.307-.8 2.447-1.847 3.401-3.045C7.267 11.85 8 10.74 8 10m-8 7h11.5c6.627 0 12-5.373 12-12V0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8 5c0 .742-.733 1.85-1.475 2.78c-.954 1.2-2.094 2.247-3.401 3.046C2.144 11.425.956 12 0 12m0 0c.956 0 2.145.575 3.124 1.174c1.307.8 2.447 1.847 3.401 3.045C7.267 17.15 8 18.26 8 19m-8-7h24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightTwoShortArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 1.5c0 1.11-1.1 2.771-2.212 4.166c-1.432 1.796-3.141 3.365-5.102 4.563c-1.469.897-3.253 1.758-4.686 1.758M12 22.5c0-1.11-1.1-2.771-2.212-4.166c-1.432-1.796-3.141-3.365-5.102-4.563c-1.469-.897-3.253-1.758-4.686-1.758M0 12h24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RvTrailer(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M19.5 7.5v.429c0 3.335 1.288 4.215 4 5.571v5H21m-1.5-11h2V7s-1.5-1-2-2.5H.5v14h2m17-11h-5v4h6.019M16 18.5a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0m-5 0H7.5m-5 0a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0m4-7v-4h-8v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafeSocialDistancing(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M24 12h-6m2.571-3c0 .317.315.792.632 1.19c.41.513.898.962 1.458 1.304c.42.256.93.502 1.339.502M20.571 15c0-.317.315-.792.632-1.19a5.63 5.63 0 0 1 1.458-1.304c.42-.256.93-.502 1.339-.502M0 12h6M3.428 9c0 .317-.314.792-.632 1.19a5.63 5.63 0 0 1-1.457 1.304c-.42.256-.93.502-1.34.502M3.43 15c0-.317-.315-.792-.633-1.19a5.631 5.631 0 0 0-1.457-1.304c-.42-.256-.93-.502-1.34-.502M14 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24m1.85-19.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SanitizeHands(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 15.5H3m.5 0A3.5 3.5 0 0 1 0 19m3.5-3.5h11v.5a3 3 0 0 1-3 3H8h7l3.824-1.593A5.285 5.285 0 0 1 20.857 17H21a2.5 2.5 0 0 1 2.5 2.5l-9.5 4H0m10-15h2m0 0h2m-2 0v-2m0 2v2M7.875 8.511c0 2.408 1.847 3.989 4.125 3.989s4.125-1.58 4.125-3.989c0-1.003-.377-2.143-1.048-2.852l-1.716-1.814C12.49 2.924 12 1.303 12 0c0 1.303-.49 2.924-1.361 3.845L8.923 5.66c-.671.71-1.048 1.849-1.048 2.852Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sauna(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15.5 1s-2.25 2.571 0 6s0 6 0 6m3-12s-2.25 2.571 0 6s0 6 0 6m3-12s-2.25 2.571 0 6s0 6 0 6M0 19.5h10m.5-5l-2-8H7.067a3 3 0 0 0-2.803 1.932L1 17m14 6.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-7v-6m1.695-6s-1.81-.557-2.135-1.776A1.768 1.768 0 0 1 6.302.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SchoolZone(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.5 23.5c-1 0-1.75-1.5-1.75-1.5C9 20.5 9 18.5 9 17v-.5L4 14v-.25l1.495-7.225M8.5 13.75l1.003-4.845A2 2 0 0 0 7.544 6.5H6c-.17 0-.338.009-.505.025M.5 13l.6-2.603a5.029 5.029 0 0 1 4.395-3.872M6 17.5c-1 3-3 5-5.5 5m19.5 1c-.813 0-1.422-1.211-1.422-1.211c-.61-1.219-.61-2.844-.61-4.063v-.406l-4.124-2.031v-.203l1.214-5.87m2.442 5.87l.846-4.087a1.5 1.5 0 0 0-1.47-1.804H15.47c-.139 0-.276.008-.41.021m-3.933 5.26l.435-2.098a3.991 3.991 0 0 1 3.497-3.162M19.92 13c1.165.837 2.623 1.164 4.08 1.164m-8.53 4.469c-.6 1.801-1.645 3.159-2.97 3.745M8.85 4.5s-1.6-1-1.6-2.25C7.25 1.284 8.034.5 9 .5c.967 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25zm8.5 3s-1.6-1-1.6-2.25a1.746 1.746 0 1 1 3.495 0c0 1.25-1.595 2.25-1.595 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18.425 18.425L23.5 23.5m-12.5-2C5.201 21.5.5 16.799.5 11S5.201.5 11 .5S21.5 5.201 21.5 11S16.799 21.5 11 21.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Send(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 13L18 6m-1.75 17.5h.25a72.694 72.694 0 0 1 6.504-21.962L23.26 1L23 .74l-.538.256A72.692 72.692 0 0 1 .5 7.5v.25l5 5v7.75h.25l1.774-1.69a11.997 11.997 0 0 1 2.313-1.723z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServiceAnimalOne(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.5 18a5.5 5.5 0 1 0-2 4.243m2-9.243V6.5h-1.172a3 3 0 0 0-2.906 2.255l-.966 3.764M21 24v-9.357a2.5 2.5 0 0 0 2.5-2.5v-.357h-3.786c-.986 0-1.785-.8-1.785-1.786v2.782a4.812 4.812 0 0 1-2.143 4.004a4.812 4.812 0 0 0-2.143 4.004v2.71h-.393a1.75 1.75 0 0 1-1.75-1.75V20m4.5 0h1.5v.292c0 .875 0 1.458.45 2.333c.001.002.45.875 1.05.875m-7.65-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServiceAnimalThree(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M7.052 17s-2.708 0-4.393 2.973C2.125 20.915 2 22.02 2 23.101V24v-4.995a7.464 7.464 0 0 1 6.444-7.394l.626-.086A8.606 8.606 0 0 0 16.5 3A2.5 2.5 0 0 0 19 5.5h4.5v1a3 3 0 0 1-3 3h-1v10.792c0 .875 0 1.458.45 2.333c0 0 .45.875 1.05.875M7.052 17l3.657.671A20 20 0 0 0 14.32 18h1.18m-8.448-1s-1.266 1.224-1.88 3.648c-.137.54-.172 1.099-.172 1.656V24m10.5-6v2.292c0 .875 0 1.458.45 2.333c0 0 .45.875 1.05.875M15.5 18l-.234-.312C13.613 15.483 12.443 12.702 12.9 10m2.599 8c2.564-.56 4-2.5 4-4a6.243 6.243 0 0 1-5.923-4.27L13.5 9.5c-2.5-1-5-4-5-9.5m-8 8v3.5a3 3 0 0 0 3 3h.013"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServiceAnimalTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M19 24v-9.357a2.5 2.5 0 0 0 2.5-2.5v-.357h-3.786c-.986 0-1.785-.8-1.785-1.786v2.782a4.812 4.812 0 0 1-2.143 4.004a4.812 4.812 0 0 0-2.143 4.004v2.71h-.393a1.75 1.75 0 0 1-1.75-1.75V20m4.5 0h1.5v.292c0 .875 0 1.458.45 2.333c.001.002.45.875 1.05.875M5.5 24v-6.5a2 2 0 0 0-2-2v-8s1.5-1 4-1s4 1 4 1V15M7.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C9.246 3.5 7.65 4.5 7.65 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="m10.5 1.5l-.181.543a7 7 0 0 1-.716 1.514a4.632 4.632 0 0 1-3.717 2.146a6.998 6.998 0 0 1-1.668-.137l-.561-.115l-1.5 2.598l.38.429c.374.422.693.884.953 1.376a4.632 4.632 0 0 1 0 4.292a7 7 0 0 1-.953 1.376l-.38.429l1.5 2.598l.56-.115a6.997 6.997 0 0 1 1.67-.137a4.632 4.632 0 0 1 3.716 2.146c.296.47.537.979.716 1.514l.181.543h3l.181-.543c.179-.536.42-1.043.716-1.514a4.632 4.632 0 0 1 3.717-2.146a6.996 6.996 0 0 1 1.668.137l.561.115l1.5-2.598l-.38-.429a7.007 7.007 0 0 1-.953-1.376a4.632 4.632 0 0 1 0-4.292c.26-.492.579-.954.953-1.376l.38-.429l-1.5-2.598l-.56.115a6.999 6.999 0 0 1-1.67.137a4.632 4.632 0 0 1-3.716-2.146a6.997 6.997 0 0 1-.716-1.514L13.5 1.5z"/><path d="M15.502 12a3.502 3.502 0 1 1-7.004 0a3.502 3.502 0 0 1 7.004 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15.567 16.946a4 4 0 1 1 6.867 4.107a4 4 0 0 1-6.867-4.107Zm0 0A24.001 24.001 0 0 0 9.06 13.77l-.41-.129m6.917-6.586a4 4 0 1 1 6.867-4.107a4 4 0 0 1-6.867 4.107Zm0 0A24 24 0 0 1 9.06 10.23l-.41.129m0 0c.225.5.35 1.055.35 1.64s-.125 1.14-.35 1.64m0-3.28a4 4 0 1 0 0 3.28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shelter(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 20v-2.5s-1.5-1-4-1s-4 1-4 1V20m15 0v-2.5s-1.5-1-4-1a8.97 8.97 0 0 0-1.5.124m8.5-6.374l-.247-.113a20 20 0 0 1-8.942-8.104L13 1.5h-2l-.311.533a20 20 0 0 1-8.942 8.104l-.247.113V22.5h21zM8.35 14.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25zm7 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ship(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 23.5a3 3 0 0 0 3-3a3 3 0 1 0 6 0a3 3 0 1 0 6 0a3 3 0 1 0 6 0a3 3 0 0 0 3 3M5.5 10.237V3.5h13v6.737M12 0v3.5M5.5 19c0-.174-.12-.77-.279-1.476c-.448-2.004-1.342-3.878-2.55-5.538L2.5 11.75v-.25l8.118-3.418A2.259 2.259 0 0 0 12 6m0 0c0 .909.545 1.73 1.383 2.082L21.5 11.5v.25l-.172.236c-1.207 1.66-2.101 3.534-2.55 5.538c-.157.705-.278 1.302-.278 1.476M12 6v13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shop(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 21a2 2 0 1 0 4 0m3 0a2 2 0 1 0 4 0M7 7.5h15.5v.25l-.239.283A16 16 0 0 0 18.5 18.34v.16h-13v-1.88c0-2.08-.066-4.158-.386-6.212C4.56 6.852 3.337 1.5 1 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Showers(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m13.228 4.772l3.03-3.03a4.242 4.242 0 0 1 7.242 3V24M5.803 7.954L4.39 9.368m-.353 2.475L2.62 13.257m-.707-3.535L.5 11.136m5.657 2.829l-1.414 1.414m3.535.707L6.864 17.5m3.182-5.303L8.632 13.61m-.707-3.536L6.51 11.49m4.243-9.193l-.155.155a6 6 0 0 1-3.89 1.747l-.728.043l-.177.177l7.779 7.778l.176-.177l.043-.728a6 6 0 0 1 1.747-3.89l.155-.155z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignLanguage(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 24V6H1a3 3 0 0 1 3 3v4V3h.5a3 3 0 0 1 3 3v6.5h2a3 3 0 0 1 3 3v6h-2A2.5 2.5 0 0 0 8 24m1-8.5V19m0-1.75h3.5M23.5 0v18H23a3 3 0 0 1-3-3v-4v10h-.5a3 3 0 0 1-3-3v-6.5h-2a3 3 0 0 1-3-3v-6h2A2.5 2.5 0 0 0 16 0m-1 8.5V5m0 1.75h-3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkatePark(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 6.5L7 13.25v.25h9v.25l-.518.674C14.197 16.094 13.5 17.893 13.5 20m1-13.5l-1.852 5m-1.884 4C9.94 16.94 9.5 18.326 9.5 20m-7-14.5s4 1 9.5 1s9.5-1 9.5-1m-2.5 16l-1.09.272a24.371 24.371 0 0 1-11.82 0L5 21.5m2 .482V24m10-2.018V24M11.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smiley(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 7v3m7-3v3M12 22.5C6.201 22.5 1.5 17.799 1.5 12S6.201 1.5 12 1.5S22.5 6.201 22.5 12S17.799 22.5 12 22.5Zm.367-9.5h-.735c-2.055 0-4.078-.516-5.882-1.5H5.5v.5a6.5 6.5 0 1 0 13 0v-.5h-.25a12.284 12.284 0 0 1-5.882 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmokingArea(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M23.5 11v6m-9-16v1A6.5 6.5 0 0 0 21 8.5m2.5.5V5.5H21A3.5 3.5 0 0 1 17.5 2V1M.5 16.5v-5h20v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Socket(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5 9.5v-4m5 4v-4m-2.5 12h3.5v-.492a10 10 0 0 1 2.191-6.247L18.5 9.75V9.5h-13v.25l.809 1.01A10 10 0 0 1 8.5 17.009v.492zm0 0v6c6.351 0 11.5-5.149 11.5-11.5S18.351.5 12 .5S.5 5.649.5 12c0 5.493 3.85 10.086 9 11.228"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinning(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 17.52a7.558 7.558 0 0 0-.5-.02a3 3 0 1 0 0 6c2 0 9.913-3.152 12-3.75v-.5c-1.017-.17-5.254-.914-8.5-1.376m8.5-2.374V11m0 4.5a4 4 0 1 1-3.71 5.5m3.71-5.5a4.002 4.002 0 0 0-3.874 3M19 8l3-3m-3 4.5h-2.969a3 3 0 0 1-2.785-1.886L12 4.5h-.914a3 3 0 0 0-1.92.695L5.5 8.25v.25l3 4v.25S7.5 15 7.5 19M10 9.5l3 2v.25s-2 1.25-2.5 3M14.805 4.5s1.81-.557 2.135-1.776A1.768 1.768 0 0 0 15.698.561a1.75 1.75 0 0 0-2.145 1.25c-.325 1.219.961 2.61.961 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stadium(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 3.5c-2.663 0-5.323-.189-7.96-.566L1 2.5H.5v19H1l3.04-.434A56.277 56.277 0 0 1 12 20.5m0-17c2.663 0 5.323-.189 7.96-.566L23 2.5h.5v19H23l-3.04-.434A56.276 56.276 0 0 0 12 20.5m0-17V9m0 11.5V15M.5 7.5H1a4.5 4.5 0 0 1 0 9H.5m23 0H23a4.5 4.5 0 1 1 0-9h.5M12 9a3 3 0 1 1 0 6m0-6a3 3 0 1 0 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StaffOnly(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M19.5 20.5v.5m0-.5A3.5 3.5 0 0 0 16 24m3.5-3.5v-11H19a3 3 0 0 0-3 3v3m0 0a3.5 3.5 0 0 0-3.5 3.5v1m3.5-4.5v-13h-.5a3 3 0 0 0-3 3V12V.5H12a3 3 0 0 0-3 3V12V2.5h-.5a3 3 0 0 0-3 3V24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stairs(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.5 24v-.149a7.5 7.5 0 0 0-1.894-4.982l-.106-.12v-.25h7v-.648a7.5 7.5 0 0 0-1.894-4.982l-.106-.12v-.25h7v-.648a7.5 7.5 0 0 0-1.894-4.982l-.106-.12V6.5h7v-.648A7.5 7.5 0 0 0 16.606.87L16.5.749V.5H24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StairsDownArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15.5 24v-.149a7.5 7.5 0 0 1 1.894-4.982l.106-.12v-.25h-7v-.648a7.5 7.5 0 0 1 1.894-4.982l.106-.12v-.25h-7v-.648A7.5 7.5 0 0 1 7.394 6.87l.106-.12V6.5H0m23.5 4l-10-10m10 10c-.398-.398-.654-1.133-.812-1.79a7.746 7.746 0 0 1-.149-2.687c.08-.697.235-1.464.544-1.773m.417 6.25c-.398-.398-1.133-.654-1.79-.812a7.745 7.745 0 0 0-2.687-.149c-.697.08-1.464.235-1.773.544"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StairsDownPerson(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14 23.5H8.5v-.008a7.5 7.5 0 0 1 .988-3.721l.012-.021v-.25h-5v-.008a7.5 7.5 0 0 1 .988-3.721l.012-.021v-.25H0M21.5 14V6.5h-1.057a3 3 0 0 0-2.938 2.392l-.998 4.823a.347.347 0 0 0 .185.38l4.36 2.181A.809.809 0 0 1 21.5 17c0 1.5 0 3.5.75 5c0 0 .75 1.5 1.75 1.5m-4.5-6c-.21 0-.466.133-.737.344C16.958 19.24 16.5 21.718 16.5 24m4.85-19.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StairsUpArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.5 24v-.149a7.5 7.5 0 0 0-1.894-4.982l-.106-.12v-.25h7v-.648a7.5 7.5 0 0 0-1.894-4.982l-.106-.12v-.25h7v-.648a7.5 7.5 0 0 0-1.894-4.982l-.106-.12V6.5H24M10.5.5l-10 10m10-10c-.398.398-1.133.654-1.79.812c-.878.212-1.79.25-2.687.149c-.697-.08-1.464-.235-1.773-.544M10.5.5c-.398.398-.654 1.133-.812 1.79a7.745 7.745 0 0 0-.149 2.687c.08.697.235 1.464.545 1.773"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StairsUpPerson(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 14V6.5H4.443a3 3 0 0 0-2.938 2.392L.5 13.75V14l5 2.5v.5c0 1.5 0 3.5.75 5c0 0 .75 1.5 1.75 1.5m-4.5-6c-.21 0-.466.133-.737.344C.958 19.24.5 21.718.5 24m9.5-.5h5.5v-.008a7.5 7.5 0 0 0-.988-3.721a.09.09 0 0 1-.012-.045V19.5h5v-.008a7.5 7.5 0 0 0-.988-3.721a.09.09 0 0 1-.012-.045V15.5H24M5.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.246 3.5 5.65 4.5 5.65 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StandHere(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M8.048 22.95c.051-.01.102-.022.151-.033c1.225-.283 1.906-1.473 1.802-2.718l-.196-2.336l-5.304.93l.636 2.213c.374 1.3 1.582 2.232 2.911 1.944Zm-4.416-7.18l-1.163-4.047a12.04 12.04 0 0 1 .31-7.585l.83-2.2S6.208 1 8.917 1l.956 2.85c.612 1.826.778 3.77.483 5.673l-.807 5.21zm12.32 7.18a11.457 11.457 0 0 1-.151-.033c-1.225-.283-1.906-1.473-1.802-2.718l.196-2.336l5.304.93l-.636 2.213c-.374 1.3-1.582 2.232-2.911 1.944Zm4.416-7.18l1.163-4.047a12.04 12.04 0 0 0-.31-7.585l-.83-2.2S17.792 1 15.083 1l-.956 2.85a12.04 12.04 0 0 0-.483 5.673l.807 5.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M11.761 1.5h.478l.619 2.243a8 8 0 0 0 7.711 5.87H22.5V10l-1.656 1.143a8 8 0 0 0-3.1 8.946l.745 2.411H18l-.904-.747a8 8 0 0 0-10.192 0L6 22.5h-.49l.746-2.41a8 8 0 0 0-3.1-8.947L1.5 10v-.386h1.93a8 8 0 0 0 7.712-5.871z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrollerParking(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3 18.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm0 0v-2l1.693-1.395M23.5 6V5A2.5 2.5 0 0 0 21 2.5h-1L4.693 15.106M16 5.75L13.539 15.8a2.886 2.886 0 0 1-4.185 1.848l-4.66-2.543M0 14c.197 0 .384-.005.561-.014c1.642-.08 2.953-1.17 3.939-2.486h.25l2.608 1.422M17.75 20.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM10.85 7s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C12.746 6 11.15 7 11.15 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StudyRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m12.165 7.835l1.3-1.3a3.536 3.536 0 0 0-5-5l-1.3 1.3c-.437.437-.97.767-1.558.963L3.5 4.5v.25l6.75 6.75h.25l.703-2.107c.195-.587.525-1.12.962-1.558Zm0 0L12.27 8a9.724 9.724 0 0 0 5.365 4.38M6.5 10.5a1.414 1.414 0 1 1-2-2m8.965-6.964L14.6.4M12 20.5c0-.66.113-1.322.415-1.91a10.533 10.533 0 0 1 5.264-4.88M12 20.5c5.5 0 8.5 2 8.5 2v1h-17v-1s3-2 8.5-2Zm5.679-6.79a1.5 1.5 0 1 0-.045-1.33m.044 1.33a1.493 1.493 0 0 1-.044-1.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M12 24a4.802 4.802 0 0 1 3.11-4.494a4.802 4.802 0 0 1 5.375.98a4.802 4.802 0 0 1-.979-5.377A4.802 4.802 0 0 1 24 12a4.802 4.802 0 0 1-4.494-3.11a4.802 4.802 0 0 1 .98-5.375a4.802 4.802 0 0 1-5.377.979A4.802 4.802 0 0 1 12 0a4.802 4.802 0 0 1-3.11 4.494a4.802 4.802 0 0 1-5.375-.98a4.802 4.802 0 0 1 .979 5.377A4.802 4.802 0 0 1 0 12a4.802 4.802 0 0 1 4.494 3.11a4.802 4.802 0 0 1-.98 5.375a4.802 4.802 0 0 1 5.377-.979A4.802 4.802 0 0 1 12 24Z"/><path d="M16 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Surgery(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M16 19.5H6.5v-1a3 3 0 0 1 3-3c1 0 14.4 1.5 14.4 1.5M0 22.5h24m-9.5-9l-2.918-2.085a1.39 1.39 0 0 1-.47-1.677l.985-2.297C12.466 7.226 13.87 6.5 16 6.5s3.534.726 3.903.94l.985 2.298a1.39 1.39 0 0 1-.47 1.678L17.5 13.5M.5.5l1.666 1.666M4.5 17.9s-1 1.6-2.25 1.6A1.75 1.75 0 0 1 .5 17.75c0-.966.783-1.746 1.75-1.746c1.25 0 2.25 1.596 2.25 1.596zM15.852 4.5s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75c.967 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25zM4.221 1.544L6.5 2v.25L2.25 6.5H2l-.456-2.28a2.275 2.275 0 0 1 2.677-2.676Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SurveillanceCamera(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M23.5 9v5m-13-.5v.649a7.5 7.5 0 0 0 1.894 4.982l.106.119v.25H0m20.5-6v-2.7a7.5 7.5 0 0 1 2.432-5.53l.568-.52V4.5H5a4.5 4.5 0 0 0 0 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwimmingPool(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 23.5a3 3 0 0 0 3-3a3 3 0 1 0 6 0a3 3 0 1 0 6 0a3 3 0 1 0 6 0a3 3 0 0 0 3 3M3 3v-.5a2 2 0 1 1 4 0v10c0 1.787.4 3.974 1.555 4.776c.272.189.614.224.945.224M13 3v-.5a2 2 0 1 1 4 0v10c0 1.787.4 3.974 1.555 4.776c.272.189.614.224.945.224M7 9.5h10m-10-4h10m-9.959 8h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Taxi(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4 16.5h2m2 0h8m2 0h2m.873-4.86c-1.326.275-4.643.86-8.873.86c-4.23 0-7.547-.585-8.873-.86m17.746 0A16 16 0 0 1 19.5 5.156V4.5h-15v.655a16 16 0 0 1-1.373 6.485m17.746 0c.367.83.807 1.63 1.314 2.39l.313.47v7h-2.25l-.075-.12a4 4 0 0 0-3.392-1.88H7.217a4 4 0 0 0-3.392 1.88l-.075.12H1.5v-7l.313-.47c.507-.76.947-1.56 1.314-2.39M9.5 2v2.5h5V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telecoil(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M19 16.5h5m-2.5.5v5m2-21.5L18 6m-6 6L.5 23.5M5.5 9a6.5 6.5 0 0 1 13 0c0 .662-.111 1.32-.328 1.944l-2.85 8.195A3.516 3.516 0 0 1 12 21.5c-.98 0-1.865-.352-2.5-1m6-11.5a3.5 3.5 0 1 0-7 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terrace(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 14v3.5a2 2 0 0 0 2 2h2V23m17-9v3.5a2 2 0 0 1-2 2h-2V23M12 8.5V15m-7 1.5h14m-7 0V23M1.5 8.5h21v-2h-.086a5 5 0 0 1-.822-.068l-.396-.066a16 16 0 0 1-8.861-4.65l-.21-.216h-.25l-.21.216a16 16 0 0 1-8.861 4.65l-.396.066a5 5 0 0 1-.822.068H1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tickets(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17.5 4.23V8m0 2v4m0 2v3.77M.5 3.5v6a2.5 2.5 0 0 1 0 5v6H1l3.04-.434a56.277 56.277 0 0 1 15.92 0L23 20.5h.5v-6a2.5 2.5 0 0 1 0-5v-6H23l-3.04.434a56.285 56.285 0 0 1-15.92 0L1 3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3 .5h18m-18 23h18m-15.5 0v-6l2.856-1.714a4.415 4.415 0 0 0 0-7.572L5.5 6.5v-6m13 0v6l-2.856 1.714a4.416 4.416 0 0 0 0 7.572L18.5 17.5v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tools(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m6.75 6.75l3.75 3.75M6.75 6.75L8.5 5v-.25l-1.44-.6a12 12 0 0 1-2.88-1.706L3 1.5h-.25L1.5 2.75V3l.944 1.18A12 12 0 0 1 4.15 7.06l.6 1.44H5zM13.734 17a102.1 102.1 0 0 1 4.516 5.5h.25l4-4v-.25s-2.47-1.852-5.5-4.516m.5-2.234a5 5 0 0 0 4.584-7H21.5l-3 3h-2v-2l3-3v-.584a5 5 0 0 0-6.703 6.287L11 10c-4.5 4.5-9.5 8.25-9.5 8.25v.25l4 4h.25S9.5 17.5 14 13l1.797-1.797a4.989 4.989 0 0 0 1.703.297Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Touch(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9 16v-3.5a3 3 0 0 0-3-3h-.5v11A3.5 3.5 0 0 1 9 24m10.5 0V12.5a3 3 0 0 0-3-3h-.901M16 12v-1a3 3 0 0 0-3-3h-.5v4V.5H12a3 3 0 0 0-3 3V13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Track(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15 21.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4L8 13.5v-.25L8.462 12m5.038-2V9A2.5 2.5 0 0 0 11 6.5h-1L6.909 8.266a.811.811 0 0 0 .075 1.447l5.766 2.537M5.5 20c0-1 1.5-1.75 1.5-1.75c1.327-.664 2.263-.74 3.5-.749M0 23.5h23.5v-.25s-5-9.75-5-13.25A3.5 3.5 0 0 1 22 6.5m-10.148-2s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75c.967 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrailerSites(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18.5 18.5H24m-5.5 0v-14h-15S.5 9 .5 16.616V18.5H7m11.5 0H12m9 2h3m-17-2a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0m-2-11v4m-6 0a23.99 23.99 0 0 1 1.172-4H15.5v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 15.5h3m6 0h3M3.5 11.715c1.513.294 4.623.785 8.5.785c3.878 0 6.987-.491 8.5-.785M8.5 18.5l-.586 1.172A7.63 7.63 0 0 1 6.636 21.5m8.864-3l.586 1.171c.337.675.77 1.29 1.279 1.829m3.635 2a2.56 2.56 0 0 1-.969-.174a7.636 7.636 0 0 1-2.666-1.826M3 23.5c.331 0 .662-.049.969-.174A7.635 7.635 0 0 0 6.636 21.5M3.5 3.5h17m-13.864 18h10.729M3.5 18.5V.5h17v18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransgenderSign(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 16.5a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm0 0V24m-3.5-3.5h7m7.997-20l-7.722 7.721M23.497.5c-.412.411-1.171.676-1.85.84a8 8 0 0 1-2.775.153c-.72-.082-1.512-.243-1.832-.563M23.497.5c-.412.411-.676 1.17-.84 1.85a8 8 0 0 0-.153 2.775c.082.72.243 1.512.562 1.831M.502.5l7.722 7.722M.502.5c.411.411 1.17.676 1.85.84a8 8 0 0 0 2.775.153c.72-.082 1.512-.243 1.831-.563M.502.5c.411.411.675 1.17.84 1.85a8 8 0 0 1 .152 2.775c-.081.72-.243 1.512-.562 1.831M3.5 7.5l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tunnel(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M22.5 21V9.625M1.5 21V9.625M12 1.5v4m0-4c-2.939 0-5.366.703-7.144 1.474M12 1.5c2.939 0 5.366.703 7.144 1.474M0 22.5h24M5.5 21v-9c0-.46.048-.908.138-1.34M18.5 21v-9c0-.46-.048-.908-.138-1.34M12 5.5a6.473 6.473 0 0 0-4.136 1.485M12 5.5c1.571 0 3.012.557 4.136 1.485M22.5 16.5h-4m-17 0h4M4.856 2.974C2.7 3.91 1.5 4.942 1.5 4.942v4.683m3.356-6.65l3.008 4.01m0 0a6.498 6.498 0 0 0-2.226 3.675M1.5 9.625l4.138 1.035m12.724 0L22.5 9.625m-4.138 1.035a6.497 6.497 0 0 0-2.226-3.675m6.364 2.64V4.942s-1.2-1.033-3.356-1.968m-3.008 4.011l3.008-4.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwentyFourHours(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M4.5 10.5v-.25a1.75 1.75 0 1 1 3.5 0v.09c0 .427-.148.84-.42 1.17L4.5 15.25v.25h4m5-2h-4v-.25A8.931 8.931 0 0 0 11 8.296V8m2.5 5.5V10m0 3.5V16M16 8v4.75m0 0a1.75 1.75 0 1 1 3.5 0V16M16 12.75V16m-4 7.5C5.649 23.5.5 18.351.5 12S5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UiPhone(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M9.5.5V2A1.5 1.5 0 0 0 11 3.5h2A1.5 1.5 0 0 0 14.5 2V.5m-4.5 20h4M4.5.75S5 7.492 5 12s-.5 11.25-.5 11.25v.25h15v-.25S19 16.508 19 12S19.5.75 19.5.75V.5h-15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnisexRestroom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M16 24v-5.5h-3.5v-.25l.072-.15A24.999 24.999 0 0 0 15 7.35v-.328a8.58 8.58 0 0 1 3-.523c1.288 0 2.311.266 3 .523v.328c0 3.72.83 7.391 2.428 10.749l.072.15v.25H20V24M8 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24M17.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25zm-12 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.746 3.5 6.15 4.5 6.15 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 10.5V6a5.5 5.5 0 1 1 11 0M12 15v4m8.5 4.5h-17v-.25l.075-.327a26.342 26.342 0 0 0 0-11.846L3.5 10.75v-.25h17v.25l-.075.327a26.34 26.34 0 0 0 0 11.846l.075.327z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpAngleArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 8c.742 0 1.85-.733 2.78-1.475c1.2-.954 2.247-2.094 3.046-3.401C6.425 2.144 7 .956 7 0m0 0c0 .956.575 2.145 1.174 3.124c.8 1.307 1.847 2.447 3.045 3.401C12.15 7.267 13.26 8 14 8M7 0v11.5c0 6.627 5.373 12 12 12h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 8c.742 0 1.85-.733 2.78-1.475c1.2-.954 2.247-2.094 3.046-3.401C11.425 2.144 12 .956 12 0m0 0c0 .956.575 2.145 1.174 3.124c.8 1.307 1.847 2.447 3.045 3.401C17.15 7.267 18.26 8 19 8m-7-8v24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpLeftArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m3 3l18 18M3 3c.676.676 1.923 1.11 3.039 1.379c1.49.359 3.036.424 4.559.252c1.182-.134 2.484-.4 3.009-.924M3 3c.676.676 1.11 1.923 1.379 3.039c.359 1.49.424 3.036.252 4.559c-.134 1.182-.4 2.484-.924 3.009"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpLeftTwoShortArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.01 3.009L21 21M4.07 18.919c.785-.786 1.182-2.737 1.381-4.51c.258-2.282.159-4.6-.381-6.834C4.666 5.902 4.014 4.032 3 3.018M18.92 4.07c-.785.785-2.737 1.182-4.51 1.381c-2.282.258-4.6.159-6.834-.381C5.902 4.666 4.032 4.014 3.018 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpRightArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M21 3L3 21M21 3c-.676.676-1.923 1.11-3.039 1.379c-1.49.359-3.036.424-4.559.252c-1.182-.134-2.484-.4-3.009-.924M21 3c-.676.676-1.11 1.923-1.379 3.039c-.359 1.49-.424 3.036-.252 4.559c.134 1.182.4 2.484.924 3.009"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpRightTwoShortArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M20.991 3.009L3.001 21m16.93-2.081c-.786-.786-1.183-2.737-1.382-4.51c-.258-2.282-.159-4.6.381-6.834c.404-1.673 1.056-3.543 2.07-4.557M5.081 4.07c.786.785 2.737 1.182 4.51 1.381c2.282.258 4.6.159 6.834-.381c1.673-.404 3.543-1.056 4.557-2.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpTwoShortArrow(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M1.5 12c1.11 0 2.771-1.1 4.166-2.212c1.796-1.432 3.365-3.141 4.563-5.102c.897-1.469 1.758-3.253 1.758-4.686M22.5 12c-1.11 0-2.771-1.1-4.166-2.212c-1.796-1.432-3.365-3.141-4.563-5.102c-.897-1.469-1.758-3.253-1.758-4.686M12 0v24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserOne(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18.5 20.247V16S16 14.5 12 14.5S5.5 16 5.5 16v4.247M1.5 12C1.5 6.201 6.201 1.5 12 1.5S22.5 6.201 22.5 12S17.799 22.5 12 22.5S1.5 17.799 1.5 12Zm10.426.5S8.5 10.68 8.5 8c0-1.933 1.569-3.5 3.504-3.5A3.495 3.495 0 0 1 15.5 8c0 2.68-3.426 4.5-3.426 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M21.435 16.677V23.5H2.565v-6.823S6.192 14.5 12 14.5c5.806 0 9.435 2.177 9.435 2.177ZM6.92 5.58A5.084 5.084 0 0 1 12.005.5a5.073 5.073 0 0 1 5.075 5.08c0 3.89-4.974 6.533-4.974 6.533h-.214S6.919 9.47 6.919 5.58Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VendingMachine(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.5 2V.5h-8v23h8V22M14 9h-1a1 1 0 0 0-1 1v.375a1 1 0 0 0 .72.96l2.56.747a1 1 0 0 1 .72.96V14a1 1 0 0 1-1 1h-1m0-6h1a1 1 0 0 1 1 1v.5M14 9V7.5m0 7.5h-1a1 1 0 0 1-1-1v-.5m2 1.5v1.5M6.5 3v4.996M6.5 21v-4.996m0-8.008a8.5 8.5 0 1 1 0 8.007m0-8.007v8.008"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 4.5h17v5h.054c.94 0 1.877-.274 2.564-.917a10.929 10.929 0 0 0 2.061-2.624L23 4.5h.5v15H23l-.82-1.459a10.928 10.928 0 0 0-2.062-2.624c-.687-.643-1.624-.917-2.564-.917H17.5v5H.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisualImpairment(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.278 14.278a3.222 3.222 0 0 0-4.556-4.556M6.584 6.584l10.832 10.832M4.204 9.204l9.992 9.992M0 12c1.276 0 2.5 1 3.5 2l2.606 2.947a8.079 8.079 0 0 0 11.788 0L20.5 14c1-1 2.224-2 3.5-2c-1.276 0-2.5-1-3.5-2l-2.606-2.947a8.079 8.079 0 0 0-11.788 0L3.5 10c-1 1-2.224 2-3.5 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoiceScan(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12 4v16m4.5-13s-.5 2.038-.5 5s.5 5 .5 5m-9-10s.5 2.038.5 5s-.5 5-.5 5m12.75-7s-.25.75-.25 2s.25 2 .25 2m-16.5-4s.25.75.25 2s-.25 2-.25 2M17 1l5.7.3L23 7M7 1l-5.7.3L1 7m6 16l-5.7-.3L1 17m16 6l5.7-.3l.3-5.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17 6.5a5.5 5.5 0 1 1 0 11m0-8a2.5 2.5 0 0 1 0 5m-2.5 7h-.25l-.523-.548A16 16 0 0 0 2.154 16H1.5V8h.654a16 16 0 0 0 11.573-4.952l.523-.548h.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeControlTelephone(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M15.553 10.972c2.08-2.194 2.08-5.75 0-7.944M17.95 13.5c3.402-3.59 3.402-9.41 0-13M13.5 8.806c.945-.998.945-2.614 0-3.612M6.94 23.5c-4.587-5.247-4.587-13.753 0-19h3.56v4.75H7.534v9.5H10.5v4.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaitingRoom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M18 3v2h2M1 19.5h7M7.5 14V6.5H6.328a3 3 0 0 0-2.906 2.255L1.5 16.25v.25h9V18c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5m5-14a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm-10.65-5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C9.246 3.5 7.65 4.5 7.65 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Walker(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10.004 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-3.5-.75-5v-.5l-5-2.5v-.25l1.803-5.228A3 3 0 0 1 7.143 6.5h1.36l.519 4.588a3 3 0 0 0 2.941 2.412H12m-6.496 4c-.21 0-.466.133-.737.344C2.962 19.24 2.504 21.718 2.504 24m8.996-.2a26 26 0 0 0 2-10v-.3h6v.3a26 26 0 0 0 2 10m-8.878-3.3h7.756M8.354 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WashHands(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M0 23.5h21.5V23a3 3 0 0 0-3-3H12h11.5v-.5a3 3 0 0 0-3-3H12M0 13a3.5 3.5 0 0 0 3.5-3.5h11v.5a3 3 0 0 1-3 3H8h10.5a3 3 0 0 1 3 3v.67m2-10.17c0 1.638-1.343 3.004-3 3.004s-3-1.366-3-3.004c0-.158.027-.324.072-.49c.284-1.07 1.045-1.94 1.749-2.793L19.5 3c.99-1.363 1-2.5 1-3c0 .5.01 1.637 1 3l.18.217c.703.853 1.464 1.723 1.749 2.792c.044.167.071.333.071.491Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Waste(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M12.5 4C12.5 2 14 .5 16 .5m-4 6.08l-.31-.196c-1.441-.912-3.266-1.258-4.717-.36C6.243 6.473 5.079 7.375 4.5 8v.293c2.08 1.073 4 3.226 4 5.707s-1.92 4.634-4 5.706V20c.58.624 1.743 1.525 2.473 1.977c1.45.897 3.276.551 4.717-.36l.31-.197l.31.196c1.441.912 3.266 1.258 4.717.36c.73-.45 1.894-1.352 2.473-1.976v-.294c-2.08-1.072-4-3.225-4-5.706c0-2.481 1.92-4.634 4-5.706V8c-.58-.624-1.743-1.525-2.473-1.977c-1.45-.897-3.276-.551-4.717.36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wc(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5.5 8v7.5h.25s1.016-1.152 1.5-2c.62-1.087 1.125-2.5 1.125-2.5h.25s.505 1.413 1.125 2.5c.484.848 1.5 2 1.5 2h.25V8m7 2.5v-1a1 1 0 0 0-1-1H15a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h2.5a1 1 0 0 0 1-1v-1m-6.5 10C5.649 23.5.5 18.351.5 12S5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WearGasMask(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M20.257 15c.181-.91.243-1.894.243-3V4.5s-2.55-4-8.5-4s-8.5 4-8.5 4V12c0 1.107.062 2.09.243 3M5 5s2 1.5 5.5 1.5L9.5 4c2 2 7.5 2.286 9.5 2.286M7.856 21.018c-.071.784-.417 1.446-1.021 1.794c-1.196.69-2.949-.107-3.915-1.78c-.967-1.675-.78-3.592.415-4.282c.885-.51 2.075-.207 3.026.664m1.495 3.604c.069-.77-.127-1.658-.606-2.487a4.706 4.706 0 0 0-.89-1.117m1.495 3.604c1.31 1.045 2.755 1.861 4.145 2.232c1.39-.37 2.834-1.187 4.145-2.232m0 0c.07.785.417 1.446 1.02 1.794c1.196.69 2.949-.107 3.915-1.78c.967-1.675.78-3.592-.415-4.282c-.885-.51-2.075-.207-3.026.664m-1.494 3.604c-.07-.77.126-1.658.605-2.487c.251-.435.556-.81.89-1.117m-11.28 0C8.466 15.138 11.5 14 11.5 14h1s3.035 1.138 5.14 3.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WearGloves(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M5 15.5v-3a3 3 0 0 0-3-3h-.5V20A3.5 3.5 0 0 1 5 23.5h4M12 2V.5h-.5a3 3 0 0 0-3 3v4m.17-5H8a3 3 0 0 0-3 3V16m10.5-4V3.5a3 3 0 0 1 3-3h.5V12V5.5a3 3 0 0 1 3-3h.5v21H12A3.5 3.5 0 0 0 8.5 20V9.5H9a3 3 0 0 1 3 3v3m3.67-13H15a3 3 0 0 0-3 3m0 0V12m0-6.5V16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WearGoggles(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M20.5 7V4.5s-2.551-4-8.5-4s-8.5 4-8.5 4V7m.787 9.75c1.321 3.01 4.59 5.667 7.713 6.5c3.124-.833 6.394-3.489 7.714-6.5M5 5s2 1.5 5.5 1.5L9.5 4c2 2 7.5 2.286 9.5 2.286M12 10c0 .505.028.982.075 1.425c.24 2.26 2.36 3.575 4.633 3.575H21s.5-3 .5-6.5h-19C2.5 12 3 15 3 15h4.292c2.272 0 4.392-1.315 4.633-3.575c.047-.443.075-.92.075-1.425Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WearHardHat(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M3.52 13c.105 2.458.66 4.304 2.438 6.244C7.658 21.1 9.89 22.688 12 23.25c2.112-.563 4.35-2.155 6.052-4.015c1.77-1.935 2.324-3.783 2.428-6.235M8.5 6l.799-4.17M15.5 6l-.799-4.17m0 0L14.5.78A10.903 10.903 0 0 0 12 .5c-.864 0-1.704.093-2.5.28l-.201 1.05m5.402 0c2.804.725 5.033 2.65 5.799 5.867c.261 1.096.563 2.178 1.5 2.803v.5s-2.257 1-10 1s-10-1-10-1v-.5c.937-.625 1.239-1.707 1.5-2.803c.766-3.216 2.995-5.142 5.799-5.867"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WearHeadset(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M19.819 16.5c-.375.94-.937 1.828-1.767 2.735c-1.702 1.86-3.94 3.452-6.052 4.015c-2.109-.562-4.342-2.15-6.042-4.006c-.835-.911-1.4-1.801-1.776-2.744M3.5 7.944L1 7.5s-.5 1.436-.5 4s.5 4 .5 4l2.5-.444zm0 0V4.5s2.551-4 8.5-4s8.5 4 8.5 4v3.444m0 0L23 7.5s.5 1.436.5 4s-.5 4-.5 4l-2.5-.444zM5 5s2 1.5 5.5 1.5L9.5 4c2 2 7.5 2.286 9.5 2.286"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WearMask(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6.5 14.018v5.79m11-5.79v5.792M5 5s2 1.5 5.5 1.5L9.5 4c2 2 7.5 2.286 9.5 2.286M20.5 10V4.5s-2.551-4-8.5-4s-8.5 4-8.5 4V10m0 2h.25S5.6 15 12 15s8.25-3 8.25-3h.25c0 2.968-.445 5.046-2.448 7.235c-1.702 1.86-3.94 3.452-6.052 4.015c-2.109-.562-4.342-2.15-6.042-4.006C3.948 17.05 3.5 14.976 3.5 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weightlifting(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M.5 11v5m2-6v7m21-6v5m-2-6v7m-19-3.5h19m-7.5 2l2 8.35m-6-8.35l-2 8.35m9.5-10.35l-1.5-6s-1.5-1-4-1s-4 1-4 1l-1.5 6m5.35-9s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WetFloor(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m17.5 1l-4.621 4.621a3 3 0 0 1-2.122.879H7.243a3 3 0 0 0-1.132.222M.5 12l4.621-4.621a3 3 0 0 1 .99-.657m5.346-.305L14 11.5h5c.936 1.123 1.591 1.965 2.596 2.528c.723.405 1.575.472 2.404.472M15.5 24c0-1.5-1-2.5-2.5-2.5m6.5 2.5c0-1.5 1-2.5 2.5-2.5m-4-1c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-.5h-3.646a3 3 0 0 1-2.683-1.658l-3.06-6.12M8.195 4.5s-1.81-.557-2.135-1.776A1.768 1.768 0 0 1 7.302.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelchairBasketball(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.207M6.5 12.522l2.136-4.35a3 3 0 0 1 2.69-1.672H12.5l2.47 2.47a1.81 1.81 0 0 0 2.56 0L20.5 6m-4.75 1.25L18.5 4.5M6 23.5a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11Zm15.5-19a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-9.805 0s-1.81-.557-2.135-1.776A1.768 1.768 0 0 1 10.802.561a1.75 1.75 0 0 1 2.146 1.25c.324 1.219-.962 2.61-.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelchairOne(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M20 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.207M15 13l1.192-3.576c.204-.612.363-1.269.172-1.885C15.943 6.18 14.561 5.5 12.5 5.5c-2.5 0-3.667 1-3.667 1l-1.333 4m4.833-3l-1.747 5.232M9 23.5a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11Zm7.805-19s1.81-.557 2.135-1.776A1.768 1.768 0 0 0 17.698.561a1.75 1.75 0 0 0-2.146 1.25c-.324 1.219.962 2.61.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelchairTennis(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M17 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.207M8 6.5c-2.5 0-3.5 1.5-4.5 4m3 2.022l2.136-4.35a3 3 0 0 1 2.69-1.672H12.5l2.463 4.105a1.842 1.842 0 0 0 2.44.682L21.26 9.25l-.38-1.447m0 0c-.229-.867-.84-1.63-2.453-3.244c-.38-.38-.7-1.02-.838-1.544c-.28-1.066.34-2.16 1.387-2.446c1.046-.286 2.122.346 2.402 1.412c.08.305.127.675.123 1.034c-.003.258-.03.51-.09.73c-.605 2.22-.76 3.191-.531 4.058ZM6 23.5a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11Zm6.305-19s1.81-.557 2.135-1.776A1.768 1.768 0 0 0 13.198.561a1.75 1.75 0 0 0-2.146 1.25c-.324 1.219.962 2.61.962 2.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelchairTwo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M14.5 6.439V12M20 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.207M11.027 6.394L9.46 12.52M3.5 4s4 2.5 9.5 2.5S22.5 4 22.5 4M9 23.5a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11Zm3.85-19s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WiFi(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path d="M8.646 17A4.489 4.489 0 0 1 12 15.5c1.333 0 2.53.58 3.354 1.5h.265l2.841-3.966A9.466 9.466 0 0 0 12 10.5a9.466 9.466 0 0 0-6.46 2.534L8.381 17z"/><path d="M20.218 10.58L23 6.5v-.25C20.09 4.051 16.5 2.5 12 2.5S3.91 4.05 1 6.25v.25l2.782 4.08A12.452 12.452 0 0 1 12 7.5c3.146 0 6.02 1.162 8.218 3.08ZM10.5 20a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiOff(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m.5.5l23 23m-7.305-7.305l2.265-3.16A9.466 9.466 0 0 0 12 10.5a9.59 9.59 0 0 0-1.398.102m-2.683.817a9.517 9.517 0 0 0-2.379 1.615L8.381 17h.265A4.489 4.489 0 0 1 12 15.5M8.115 8.115A12.49 12.49 0 0 1 12 7.5c3.146 0 6.02 1.162 8.218 3.08L23 6.5v-.25C20.09 4.051 16.5 2.5 12 2.5c-2.954 0-5.516.668-7.75 1.75M2.022 5.522c-.35.233-.69.476-1.023.727v.25l2.782 4.08A12.6 12.6 0 0 1 5.7 9.199M10.5 20a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WomenRestroom(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M10 24v-5.5H6.5v-.25l.072-.15A25 25 0 0 0 9 7.35v-.328a8.58 8.58 0 0 1 3-.523c1.288 0 2.311.266 3 .523v.328c0 3.72.83 7.391 2.428 10.749l.072.15v.25H14V24M11.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yoga(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="m2 14.5l1.654-.551a6 6 0 0 0 3.795-3.795L8 8.5s1.5-1 4-1s4 1 4 1l.551 1.654a6 6 0 0 0 3.795 3.795L22 14.5m-12.5-4v5l-4.25.85a.937.937 0 0 0-.042 1.825c1.305.332 4.198 1.095 6.792 1.99c2.657.916 5 1.97 5 2.835m-2.5-12.5v5l4.25.85a.937.937 0 0 1 .042 1.825c-.974.248-2.833.736-4.792 1.34m-4 1.397c-1.73.712-3 1.45-3 2.088m4.85-17.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoo(children ...ElementRenderer) *GuidanceIcon {
	return &GuidanceIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" d="M6 3h-.5a5 5 0 0 0-5 5v13h3l1.072-2.679a3.693 3.693 0 0 1 6.857 0L12.5 21h3v-6.5a1.5 1.5 0 1 1 3 0v4A2.5 2.5 0 0 0 21 21h2.5v-3h-2V8a5 5 0 0 0-5-5h-8v6a2.5 2.5 0 0 0 2.5 2.5h.5l.867-2.603A2.775 2.775 0 0 1 15 7m3.5 2.5v-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
